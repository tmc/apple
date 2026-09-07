package coremlcompiler

import (
	"encoding/binary"
	"encoding/hex"
	"math"
	"strings"
	"testing"
)

// modelIRFinalNormBlob is the weight file github.com/tmc/modelir emitted for a
// const named final_norm holding modelIRFinalNormValues, captured from
// BuildMILTransformerWeightFiles (target/mil/ane/blob.go, buildBLOBFileFP16).
//
// modelir writes one blob per file with a 64-byte file header and a 64-byte
// chunk header, and describes that as its own format. It is the same format
// WriteMILBlob writes for a single entry, which is what makes
// BlobLayoutFilePerConst a layout rather than a second encoder.
const modelIRFinalNormBlob = "01000000020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000" +
	"efbeadde010000001000000000000000800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000" +
	"003c00c10030ff7b0000000048424056"

var modelIRFinalNormValues = []float32{1, -2.5, 0.125, 65504, 1e-8, 0, 3.14159, 100}

// TestBuildWeightsMatchesModelIR is the check that says the file-per-const
// layout is not a reimplementation: byte-for-byte, WriteMILBlob with one entry
// produces what modelir's own blob writer produced for the same tensor.
func TestBuildWeightsMatchesModelIR(t *testing.T) {
	want, err := hex.DecodeString(modelIRFinalNormBlob)
	if err != nil {
		t.Fatal(err)
	}
	tensors := []WeightTensor{{
		Name:  "final_norm",
		DType: BlobDataTypeFloat16,
		Data:  Float16Bytes(modelIRFinalNormValues),
	}}
	files, refs, err := BuildWeights(tensors, BlobLayoutFilePerConst)
	if err != nil {
		t.Fatalf("BuildWeights: %v", err)
	}
	if len(files) != 1 {
		t.Fatalf("BuildWeights returned %d files, want 1", len(files))
	}
	// modelir names this file @model_path/weights/final_norm.bin.
	if got, want := files[0].Path, "@model_path/weights/final_norm.bin"; got != want {
		t.Errorf("path = %q, want %q", got, want)
	}
	if got := refs["final_norm"].Offset; got != 64 {
		t.Errorf("offset = %d, want 64", got)
	}
	if got := files[0].Blob; !bytesEqual(got, want) {
		t.Errorf("blob differs from modelir's:\n got %s\nwant %s",
			hex.EncodeToString(got), hex.EncodeToString(want))
	}
}

func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// TestBuildWeightsSingleFile covers the layout coremltools produces and
// mlx-go's mlxcext reimplemented: one file, each tensor at its own offset.
func TestBuildWeightsSingleFile(t *testing.T) {
	tensors := []WeightTensor{
		{Name: "a", DType: BlobDataTypeFloat16, Data: Float16Bytes([]float32{1, 2, 3})},
		{Name: "b", DType: BlobDataTypeInt8, Data: []byte{1, 2, 3, 4, 5}},
		{Name: "c", DType: BlobDataTypeFloat16, Data: Float16Bytes([]float32{7})},
	}
	files, refs, err := BuildWeights(tensors, BlobLayoutSingleFile)
	if err != nil {
		t.Fatalf("BuildWeights: %v", err)
	}
	if len(files) != 1 {
		t.Fatalf("BuildWeights returned %d files, want 1", len(files))
	}
	if got, want := files[0].Path, "@model_path/weights/weight.bin"; got != want {
		t.Errorf("path = %q, want %q", got, want)
	}
	blob := files[0].Blob
	if n := binary.LittleEndian.Uint32(blob[0:]); n != 3 {
		t.Errorf("entry count = %d, want 3", n)
	}
	if v := binary.LittleEndian.Uint32(blob[4:]); v != 2 {
		t.Errorf("version = %d, want 2", v)
	}
	if len(refs) != 3 {
		t.Fatalf("got %d refs, want 3", len(refs))
	}
	seen := map[uint64]bool{}
	for _, name := range []string{"a", "b", "c"} {
		r, ok := refs[name]
		if !ok {
			t.Fatalf("no ref for %q", name)
		}
		if r.Path != files[0].Path {
			t.Errorf("%s path = %q, want the single weight file", name, r.Path)
		}
		if r.Offset%64 != 0 {
			t.Errorf("%s offset %d is not 64-byte aligned", name, r.Offset)
		}
		if seen[r.Offset] {
			t.Errorf("%s reuses offset %d", name, r.Offset)
		}
		seen[r.Offset] = true
		// The BLOBFILE offset must point at a metadata block.
		if got := binary.LittleEndian.Uint32(blob[r.Offset:]); got != milBlobSentinel {
			t.Errorf("%s offset %d does not point at the 0xDEADBEEF sentinel (got %#x)", name, r.Offset, got)
		}
	}
	// Every tensor's payload must be recoverable from its metadata alone.
	r := refs["b"]
	size := binary.LittleEndian.Uint64(blob[r.Offset+8:])
	off := binary.LittleEndian.Uint64(blob[r.Offset+16:])
	if got := blob[off : off+size]; !bytesEqual(got, []byte{1, 2, 3, 4, 5}) {
		t.Errorf("b payload = %v, want [1 2 3 4 5]", got)
	}
}

// TestBuildWeightsLayoutsAgree checks the two layouts differ only in packing:
// the same tensor's metadata and payload must be identical under both.
func TestBuildWeightsLayoutsAgree(t *testing.T) {
	tensors := []WeightTensor{
		{Name: "w", DType: BlobDataTypeFloat16, Data: Float16Bytes([]float32{1, 2, 3, 4})},
	}
	single, singleRefs, err := BuildWeights(tensors, BlobLayoutSingleFile)
	if err != nil {
		t.Fatal(err)
	}
	perConst, perConstRefs, err := BuildWeights(tensors, BlobLayoutFilePerConst)
	if err != nil {
		t.Fatal(err)
	}
	if singleRefs["w"].Offset != perConstRefs["w"].Offset {
		t.Errorf("offsets differ: single=%d per-const=%d", singleRefs["w"].Offset, perConstRefs["w"].Offset)
	}
	if !bytesEqual(single[0].Blob, perConst[0].Blob) {
		t.Error("a lone tensor's file differs between layouts")
	}
}

func TestBuildWeightsRejects(t *testing.T) {
	for _, tt := range []struct {
		name    string
		tensors []WeightTensor
		layout  BlobLayout
		wantErr string
	}{
		{"unnamed", []WeightTensor{{DType: BlobDataTypeFloat16}}, BlobLayoutSingleFile, "has no name"},
		{
			"duplicate name",
			[]WeightTensor{{Name: "w", DType: BlobDataTypeFloat16}, {Name: "w", DType: BlobDataTypeFloat16}},
			BlobLayoutSingleFile, "appears twice",
		},
		{
			// The const name becomes a path under this layout.
			"name escapes the weights directory",
			[]WeightTensor{{Name: "../w", DType: BlobDataTypeFloat16}},
			BlobLayoutFilePerConst, "cannot name a file",
		},
		{
			"name contains a separator",
			[]WeightTensor{{Name: "a/b", DType: BlobDataTypeFloat16}},
			BlobLayoutFilePerConst, "cannot name a file",
		},
		{"unknown layout", []WeightTensor{{Name: "w"}}, BlobLayout(7), "unknown blob layout"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := BuildWeights(tt.tensors, tt.layout)
			if err == nil {
				t.Fatalf("BuildWeights = nil error, want %q", tt.wantErr)
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("BuildWeights = %v, want %q", err, tt.wantErr)
			}
		})
	}
}

// TestBuildWeightsFilePerConstIsWritable ties the builder to the writer: the
// paths one produces must be paths the other accepts.
func TestBuildWeightsFilePerConstIsWritable(t *testing.T) {
	tensors := []WeightTensor{
		{Name: "w1", DType: BlobDataTypeFloat16, Data: Float16Bytes([]float32{1, 2})},
		{Name: "w2", DType: BlobDataTypeFloat16, Data: Float16Bytes([]float32{3, 4})},
	}
	for _, layout := range []BlobLayout{BlobLayoutSingleFile, BlobLayoutFilePerConst} {
		files, _, err := BuildWeights(tensors, layout)
		if err != nil {
			t.Fatalf("%v: %v", layout, err)
		}
		if err := WriteWeightRoot(t.TempDir(), files); err != nil {
			t.Errorf("%v: WriteWeightRoot: %v", layout, err)
		}
	}
}

func TestFloat16bits(t *testing.T) {
	for _, tt := range []struct {
		in   float32
		want uint16
	}{
		{0, 0x0000},
		{float32(math.Copysign(0, -1)), 0x8000},
		{1, 0x3c00},
		{-1, 0xbc00},
		{-2.5, 0xc100},
		{0.125, 0x3000},
		{65504, 0x7bff},        // largest binary16
		{65520, 0x7c00},        // rounds up past it, to +Inf
		{1e-8, 0x0000},         // far below the smallest subnormal
		{6.0975552e-5, 0x03ff}, // largest subnormal
		{float32(math.Inf(1)), 0x7c00},
		{float32(math.Inf(-1)), 0xfc00},
	} {
		if got := Float16bits(tt.in); got != tt.want {
			t.Errorf("Float16bits(%v) = %#04x, want %#04x", tt.in, got, tt.want)
		}
	}
	if got := Float16bits(float32(math.NaN())); got&0x7c00 != 0x7c00 || got&0x03ff == 0 {
		t.Errorf("Float16bits(NaN) = %#04x, want a NaN pattern", got)
	}
}

// TestFloat16RoundTrip checks every binary16 value survives a round trip, so
// the two directions cannot drift apart.
func TestFloat16RoundTrip(t *testing.T) {
	for i := 0; i < 1<<16; i++ {
		h := uint16(i)
		exp := h & 0x7c00
		if exp == 0x7c00 && h&0x03ff != 0 {
			continue // NaN payloads are not preserved bit-for-bit
		}
		if got := Float16bits(Float16frombits(h)); got != h {
			t.Fatalf("Float16bits(Float16frombits(%#04x)) = %#04x", h, got)
		}
	}
}

// TestFloat16bitsTiesToEven pins the rounding mode. modelir's own converter
// rounds ties away from zero, so a value exactly between two binary16 numbers
// is where the two disagree; IEEE 754 and numpy round to even.
func TestFloat16bitsTiesToEven(t *testing.T) {
	// 1 + 2^-11 is exactly halfway between 1.0 (0x3c00) and the next
	// binary16 up (0x3c01). Ties to even picks 0x3c00.
	tie := math.Float32frombits(math.Float32bits(1) | 1<<12)
	if got := Float16bits(tie); got != 0x3c00 {
		t.Errorf("Float16bits(%v) = %#04x, want 0x3c00 (ties to even)", tie, got)
	}
	// The next tie up sits between 0x3c01 and 0x3c02 and rounds to 0x3c02.
	tie2 := math.Float32frombits(math.Float32bits(1) | 3<<12)
	if got := Float16bits(tie2); got != 0x3c02 {
		t.Errorf("Float16bits(%v) = %#04x, want 0x3c02 (ties to even)", tie2, got)
	}
}
