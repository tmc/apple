//go:build darwin

package mil_test

import (
	"encoding/binary"
	"fmt"
	"math"
	"testing"

	"github.com/tmc/apple/x/ane"
	"github.com/tmc/apple/x/ane/mil"
)

// packedBlob encodes entries with each metadata block placed immediately after
// the previous entry's payload, rather than rounded up to the next 64-byte
// boundary. This is the layout ANEForge's _blob.py computes and the one
// x/coremlcompiler and x/ane/mil do not produce; the point of the test below is
// to find out whether the compiler rejects it or merely never sees it.
func packedBlob(entries [][]byte) ([]byte, []uint64) {
	buf := make([]byte, 64)
	binary.LittleEndian.PutUint32(buf[0:], uint32(len(entries)))
	binary.LittleEndian.PutUint32(buf[4:], 2)

	offsets := make([]uint64, len(entries))
	for i, data := range entries {
		metaOff := len(buf)
		offsets[i] = uint64(metaOff)

		meta := make([]byte, 64)
		binary.LittleEndian.PutUint32(meta[0:], 0xDEADBEEF)
		binary.LittleEndian.PutUint32(meta[4:], uint32(mil.BlobFloat16))
		binary.LittleEndian.PutUint64(meta[8:], uint64(len(data)))
		binary.LittleEndian.PutUint64(meta[16:], uint64(metaOff+64))
		buf = append(buf, meta...)
		buf = append(buf, data...)
	}
	return buf, offsets
}

func fp16LE(data []float32) []byte {
	raw := make([]byte, len(data)*2)
	for i, v := range data {
		binary.LittleEndian.PutUint16(raw[i*2:], ane.Float32ToFP16(v))
	}
	return raw
}

// scaledIdentity returns an n x n identity scaled by s, row-major, as a conv
// weight of shape [n, n, 1, 1].
func scaledIdentity(n int, s float32) []float32 {
	w := make([]float32, n*n)
	for i := range n {
		w[i*n+i] = s
	}
	return w
}

func twoConvOneFile(channels, spatial int, off1, off2 uint64) string {
	return fmt.Sprintf(`program(1.3)
[buildInfo = dict<string, string>({{"coremlc-component-MIL", "3400.4.1"}, {"coremlc-version", "3400.6.1"}})]
{
    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> x) {
        string pt = const()[name=string("pt"), val=string("valid")];
        tensor<int32, [2]> st = const()[name=string("st"), val=tensor<int32, [2]>([1,1])];
        tensor<int32, [4]> pd = const()[name=string("pd"), val=tensor<int32, [4]>([0,0,0,0])];
        tensor<int32, [2]> dl = const()[name=string("dl"), val=tensor<int32, [2]>([1,1])];
        int32 gr = const()[name=string("gr"), val=int32(1)];
        tensor<fp16, [%d,%d,1,1]> W1 = const()[name=string("W1"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string("@model_path/weights/w.bin"), offset=uint64(%d)))];
        tensor<fp16, [%d,%d,1,1]> W2 = const()[name=string("W2"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string("@model_path/weights/w.bin"), offset=uint64(%d)))];
        tensor<fp16, [1,%d,1,%d]> h = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=W1,x=x)[name=string("c1")];
        tensor<fp16, [1,%d,1,%d]> out = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=W2,x=h)[name=string("c2")];
    } -> (out);
}
`, channels, spatial,
		channels, channels, channels, channels, off1,
		channels, channels, channels, channels, off2,
		channels, spatial,
		channels, spatial,
	)
}

// TestBlobPackedVersusAlignedLayout compiles the same two-weight program twice
// from one weight file, once with each entry rounded up to a 64-byte boundary
// and once with the entries packed end to end. Twelve channels makes the two
// layouts disagree: the first payload is 288 bytes, so the aligned rule puts the
// second entry at 448 and the packed rule at 416. At a channel count whose
// payload is a multiple of 64 the two layouts coincide and the test would prove
// nothing.
//
// The two convolutions use different scale factors, so a run that read the wrong
// entry for either weight produces a different answer instead of passing.
func TestBlobPackedVersusAlignedLayout(t *testing.T) {
	c := openOrSkip(t)
	defer c.Close()

	const channels, spatial = 12, 4
	const s1, s2 = 2, 3

	w1 := fp16LE(scaledIdentity(channels, s1))
	w2 := fp16LE(scaledIdentity(channels, s2))
	if len(w1)%64 == 0 {
		t.Fatalf("payload is %d bytes, a multiple of 64: the aligned and packed layouts coincide and this test cannot discriminate", len(w1))
	}

	aw := mil.NewBlobWriter()
	i1 := aw.AddRaw(mil.BlobFloat16, w1)
	i2 := aw.AddRaw(mil.BlobFloat16, w2)
	alignedBlob, err := aw.Build()
	if err != nil {
		t.Fatal(err)
	}
	alignedOffs := []uint64{aw.Offset(i1), aw.Offset(i2)}

	packed, packedOffs := packedBlob([][]byte{w1, w2})

	if alignedOffs[1] == packedOffs[1] {
		t.Fatalf("aligned and packed second offsets are both %d; the layouts do not disagree", alignedOffs[1])
	}
	t.Logf("second entry: aligned offset %d, packed offset %d (payload %d bytes)",
		alignedOffs[1], packedOffs[1], len(w1))

	input := make([]float32, channels*spatial)
	for i := range input {
		input[i] = float32(i%7) + 1
	}
	want := make([]float32, len(input))
	for i, v := range input {
		want[i] = v * s1 * s2
	}

	run := func(name string, blob []byte, offs []uint64) (bool, []float32) {
		m, err := c.Compile(ane.CompileOptions{
			ModelType: ane.ModelTypeMIL,
			MILText:   []byte(twoConvOneFile(channels, spatial, offs[0], offs[1])),
			WeightFiles: []ane.WeightFile{
				{Path: "@model_path/weights/w.bin", Blob: blob},
			},
		})
		if err != nil {
			t.Logf("%s layout: compile REJECTED: %v", name, err)
			return false, nil
		}
		defer m.Close()
		if err := m.WriteInputFP16(0, input); err != nil {
			t.Fatal(err)
		}
		if err := m.Eval(); err != nil {
			t.Logf("%s layout: compiled but eval failed: %v", name, err)
			return false, nil
		}
		got := make([]float32, channels*spatial)
		if err := m.ReadOutputFP16(0, got); err != nil {
			t.Fatal(err)
		}
		return true, got
	}

	check := func(name string, got []float32) bool {
		for i := range want {
			if d := math.Abs(float64(got[i] - want[i])); d > 0.05 {
				t.Logf("%s layout: out[%d] = %v, want %v", name, i, got[i], want[i])
				return false
			}
		}
		return true
	}

	// Positive control: the layout this repository emits must compile and
	// compute correctly, or a rejection of the packed layout means nothing.
	alignedOK, alignedGot := run("aligned", alignedBlob, alignedOffs)
	if !alignedOK {
		t.Fatal("aligned layout failed; the packed result below would be uninterpretable")
	}
	if !check("aligned", alignedGot) {
		t.Fatal("aligned layout produced wrong values; the packed result below would be uninterpretable")
	}
	t.Log("aligned layout: compiled, ran, and matched the reference")

	packedOK, packedGot := run("packed", packed, packedOffs)
	switch {
	case !packedOK:
		t.Log("VERDICT: 64-byte entry alignment is REQUIRED; the packed layout was rejected")
	case check("packed", packedGot):
		t.Log("VERDICT: 64-byte entry alignment is ACCEPTED but not required; the packed layout compiled, ran, and matched the reference")
	default:
		t.Error("VERDICT: the packed layout compiled and ran but produced wrong values, which is worse than a rejection: the compiler followed the offsets without validating them")
	}
}
