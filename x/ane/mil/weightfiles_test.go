//go:build darwin

package mil_test

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/tmc/apple/x/ane"
	"github.com/tmc/apple/x/ane/mil"
)

// chainedConvs builds a chain of n 1x1 convolutions. Each weight is read from
// paths[i] at offsets[i], so the same program can be fed either n separate
// one-entry files or a single n-entry file.
func chainedConvs(n, channels, spatial int, paths []string, offsets []uint64) string {
	var b strings.Builder
	fmt.Fprintf(&b, `program(1.3)
[buildInfo = dict<string, string>({{"coremlc-component-MIL", "3400.4.1"}, {"coremlc-version", "3400.6.1"}})]
{
    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> x) {
        string pt = const()[name=string("pt"), val=string("valid")];
        tensor<int32, [2]> st = const()[name=string("st"), val=tensor<int32, [2]>([1,1])];
        tensor<int32, [4]> pd = const()[name=string("pd"), val=tensor<int32, [4]>([0,0,0,0])];
        tensor<int32, [2]> dl = const()[name=string("dl"), val=tensor<int32, [2]>([1,1])];
        int32 gr = const()[name=string("gr"), val=int32(1)];
`, channels, spatial)
	for i := range n {
		fmt.Fprintf(&b, "        tensor<fp16, [%d,%d,1,1]> W%d = const()[name=string(\"W%d\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"%s\"), offset=uint64(%d)))];\n",
			channels, channels, i, i, channels, channels, paths[i], offsets[i])
	}
	prev := "x"
	for i := range n {
		fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> h%d = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=W%d,x=%s)[name=string(\"c%d\")];\n",
			channels, spatial, i, i, prev, i)
		prev = fmt.Sprintf("h%d", i)
	}
	fmt.Fprintf(&b, "    } -> (%s);\n}\n", prev)
	return b.String()
}

// chainWeights returns weight i of the chain as a [channels, channels, 1, 1]
// diagonal. Each odd weight is the elementwise reciprocal of the even one before
// it, so the product down an even-length chain is 1 per channel and the output
// stays in a range fp16 represents well at any depth. Every tensor is distinct,
// which matters: a chain built from a couple of repeated values could be folded
// by the compiler, and a large n would not really be a large-n model.
func chainWeights(i, channels int) []float32 {
	w := make([]float32, channels*channels)
	for c := range channels {
		v := 2 + float32(i/2)*0.01 + float32(c)*0.001
		if i%2 == 1 {
			v = 1 / v
		}
		w[c*channels+c] = v
	}
	return w
}

// chainScale is the exact per-channel gain of the first n weights, in float64.
func chainScale(n, channels int) []float64 {
	scale := make([]float64, channels)
	for c := range channels {
		scale[c] = 1
	}
	for i := range n {
		w := chainWeights(i, channels)
		for c := range channels {
			scale[c] *= float64(w[c*channels+c])
		}
	}
	return scale
}

// TestWeightsInOneFileExceedFileCeiling records that the ANE compiler's limit is
// on the number of weight FILES, not the number of weight tensors, and that
// packing every tensor into one multi-entry blob clears it.
//
// On macOS 26.x the n-files form compiles at 16 and fails at 17, while the
// one-file form compiles, runs and matches a float64 reference at every depth
// tried up to 128. That is the whole point of BlobWriter: it is what makes a
// model with more than sixteen weights reachable at all.
func TestWeightsInOneFileExceedFileCeiling(t *testing.T) {
	c := openOrSkip(t)
	defer c.Close()

	const channels, spatial = 8, 4
	input := make([]float32, channels*spatial)
	for i := range input {
		input[i] = float32(i%5) + 1
	}

	// Control: the tensors must be distinct, or a deep chain could be folded
	// into a few constants and prove nothing about depth.
	var seen [][]byte
	for i := range 32 {
		b, err := mil.BuildWeightBlob(chainWeights(i, channels), channels, channels)
		if err != nil {
			t.Fatal(err)
		}
		for j, prev := range seen {
			if bytes.Equal(prev, b) {
				t.Fatalf("weight %d is byte-identical to weight %d; the chain could be folded and this test would prove nothing", i, j)
			}
		}
		seen = append(seen, b)
	}

	// run compiles and evaluates the n-conv chain, returning the worst absolute
	// difference from the reference, or an error.
	run := func(n int, oneFile bool) (float64, error) {
		paths := make([]string, n)
		offsets := make([]uint64, n)
		var files []ane.WeightFile

		if oneFile {
			w := mil.NewBlobWriter()
			idx := make([]int, n)
			for i := range n {
				idx[i] = w.AddFloat16(chainWeights(i, channels))
			}
			blob, err := w.Build()
			if err != nil {
				return 0, err
			}
			const path = "@model_path/weights/w.bin"
			for i := range n {
				paths[i] = path
				offsets[i] = w.Offset(idx[i])
			}
			files = []ane.WeightFile{{Path: path, Blob: blob}}
		} else {
			for i := range n {
				blob, err := mil.BuildWeightBlob(chainWeights(i, channels), channels, channels)
				if err != nil {
					return 0, err
				}
				paths[i] = fmt.Sprintf("@model_path/weights/w%d.bin", i)
				offsets[i] = 64
				files = append(files, ane.WeightFile{Path: paths[i], Blob: blob})
			}
		}

		m, err := c.Compile(ane.CompileOptions{
			ModelType:   ane.ModelTypeMIL,
			MILText:     []byte(chainedConvs(n, channels, spatial, paths, offsets)),
			WeightFiles: files,
		})
		if err != nil {
			return 0, err
		}
		defer m.Close()
		if err := m.WriteInputFP16(0, input); err != nil {
			return 0, err
		}
		if err := m.Eval(); err != nil {
			return 0, err
		}
		got := make([]float32, channels*spatial)
		if err := m.ReadOutputFP16(0, got); err != nil {
			return 0, err
		}
		scale := chainScale(n, channels)
		worst := 0.0
		for i := range got {
			want := float64(input[i]) * scale[i/spatial]
			if d := math.Abs(float64(got[i]) - want); d > worst {
				worst = d
			}
		}
		return worst, nil
	}

	// Positive control: both forms must work at a depth below the ceiling, or a
	// failure of the n-files form higher up would not be attributable to depth.
	const small = 8
	if _, err := run(small, false); err != nil {
		t.Fatalf("%d separate files failed below the ceiling: %v", small, err)
	}
	if _, err := run(small, true); err != nil {
		t.Fatalf("one file with %d entries failed below the ceiling: %v", small, err)
	}

	// Locate the n-files ceiling rather than assuming the recorded value.
	ceiling := 0
	for n := small; n <= 24; n++ {
		if _, err := run(n, false); err != nil {
			break
		}
		ceiling = n
	}
	t.Logf("separate weight files: highest n that compiles is %d", ceiling)

	// The guarantee: one file carries far more tensors than that, and computes
	// the right answer rather than merely compiling.
	for _, n := range []int{ceiling + 1, 32, 64, 128} {
		worst, err := run(n, true)
		if err != nil {
			t.Errorf("one file with %d entries: %v", n, err)
			continue
		}
		if worst > 0.05 {
			t.Errorf("one file with %d entries: worst difference %.4f exceeds tolerance", n, worst)
			continue
		}
		t.Logf("one file with %-3d entries: ran, worst difference %.4f", n, worst)
	}

	if _, err := run(ceiling+1, false); err == nil {
		t.Logf("note: %d separate files now compiles; the file ceiling has moved and the comparison above is no longer discriminating", ceiling+1)
	}
}
