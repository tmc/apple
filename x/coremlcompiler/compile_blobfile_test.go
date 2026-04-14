//go:build darwin

package coremlcompiler

import (
	"encoding/binary"
	"math"
	"os"
	"path/filepath"
	"testing"
)

// TestCompileBlobFileLoadable builds a model with BLOBFILE weight refs,
// compiles it with the pure Go Compile path, and verifies the resulting
// .mlmodelc can be loaded by Apple's MLModel runtime.
func TestCompileBlobFileLoadable(t *testing.T) {
	fp16Data := makeFP16Bytes(1.0, 2.0, 3.0, 4.0)

	// Build weight.bin in MILBlob Storage v2 format.
	blobData, offsets := WriteMILBlob([]BlobEntry{{
		DType: BlobDataTypeFloat16,
		Data:  fp16Data,
	}})

	tmpDir := t.TempDir()
	weightDir := filepath.Join(tmpDir, "weights")
	if err := os.MkdirAll(weightDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(weightDir, "weight.bin"), blobData, 0o644); err != nil {
		t.Fatal(err)
	}

	// Build model proto with BLOBFILE ref at the correct offset.
	proto := buildBlobFileModelProto(t, offsets[0])

	outputDir := filepath.Join(tmpDir, "model.mlmodelc")
	if err := CompileMLProgram(proto, tmpDir, outputDir); err != nil {
		t.Fatalf("CompileMLProgram: %v", err)
	}

	// Verify output files.
	for _, name := range []string{"model.mil", "coremldata.bin", "metadata.json"} {
		if _, err := os.Stat(filepath.Join(outputDir, name)); err != nil {
			t.Errorf("missing %s: %v", name, err)
		}
	}

	// Verify weight.bin was copied.
	copiedWeights, err := os.ReadFile(filepath.Join(outputDir, "weights", "weight.bin"))
	if err != nil {
		t.Fatalf("weight.bin not copied: %v", err)
	}
	if len(copiedWeights) != len(blobData) {
		t.Fatalf("copied weight size = %d, want %d", len(copiedWeights), len(blobData))
	}

	// Load with Apple's MLModel runtime.
	model, err := LoadCoreMLModel(outputDir)
	if err != nil {
		milData, _ := os.ReadFile(filepath.Join(outputDir, "model.mil"))
		t.Fatalf("LoadCoreMLModel failed: %v\nmodel.mil:\n%s", err, milData)
	}
	model.Close()
}

// TestCompileBlobFileMultiTensor verifies BLOBFILE refs with multiple tensors
// at different offsets in a single weight.bin.
func TestCompileBlobFileMultiTensor(t *testing.T) {
	fp16A := makeFP16Bytes(1.0, 2.0, 3.0, 4.0)
	fp16B := makeFP16Bytes(10.0, 20.0, 30.0, 40.0)

	blobData, offsets := WriteMILBlob([]BlobEntry{
		{DType: BlobDataTypeFloat16, Data: fp16A},
		{DType: BlobDataTypeFloat16, Data: fp16B},
	})

	tmpDir := t.TempDir()
	weightDir := filepath.Join(tmpDir, "weights")
	os.MkdirAll(weightDir, 0o755)
	os.WriteFile(filepath.Join(weightDir, "weight.bin"), blobData, 0o644)

	// Build model: const_a (BLOBFILE @ offsets[0]), const_b (BLOBFILE @ offsets[1]),
	// y = add(const_a, const_b).
	proto := buildMultiTensorModelProto(t, offsets[0], offsets[1])

	outputDir := filepath.Join(tmpDir, "model.mlmodelc")
	if err := CompileMLProgram(proto, tmpDir, outputDir); err != nil {
		t.Fatalf("CompileMLProgram: %v", err)
	}

	model, err := LoadCoreMLModel(outputDir)
	if err != nil {
		milData, _ := os.ReadFile(filepath.Join(outputDir, "model.mil"))
		t.Fatalf("LoadCoreMLModel failed: %v\nmodel.mil:\n%s", err, milData)
	}
	model.Close()
}

// buildBlobFileModelProto builds a minimal Model proto with a const op
// referencing BLOBFILE at the given offset, added to the input.
func buildBlobFileModelProto(t *testing.T, blobOffset uint64) []byte {
	t.Helper()
	return EncodeModel(&Model{
		SpecVersion: 8,
		Description: ModelDescription{
			Inputs:  []FeatureDescription{{Name: "x", Type: fp16ArrayType(1, 4)}},
			Outputs: []FeatureDescription{{Name: "y", Type: fp16ArrayType(1, 4)}},
		},
		MLProgram: &Program{
			Version: 1,
			Functions: map[string]*Function{
				"main": {
					OpSet:  "CoreML8",
					Inputs: []NamedValueType{{Name: "x", Type: fp16TensorVT(1, 4)}},
					BlockSpecializations: map[string]*Block{
						"CoreML8": {
							Operations: []*Operation{
								{
									Type:    "const",
									Outputs: []NamedValueType{{Name: "w", Type: fp16TensorVT(1, 4)}},
									Attributes: map[string]*Value{
										"name": stringAttr("const_w"),
										"val":  blobFileAttr(fp16TensorVT(1, 4), "@model_path/weights/weight.bin", blobOffset),
									},
								},
								{
									Type:    "add",
									Outputs: []NamedValueType{{Name: "y", Type: fp16TensorVT(1, 4)}},
									Inputs: map[string]*Argument{
										"x": {Bindings: []Binding{{Name: "x"}}},
										"y": {Bindings: []Binding{{Name: "w"}}},
									},
									Attributes: map[string]*Value{
										"name": stringAttr("add_y"),
									},
								},
							},
							Outputs: []string{"y"},
						},
					},
				},
			},
		},
	})
}

// buildMultiTensorModelProto builds a model with two const ops from BLOBFILE
// and an add op combining them.
func buildMultiTensorModelProto(t *testing.T, offsetA, offsetB uint64) []byte {
	t.Helper()
	return EncodeModel(&Model{
		SpecVersion: 8,
		Description: ModelDescription{
			Inputs:  []FeatureDescription{{Name: "x", Type: fp16ArrayType(1, 4)}},
			Outputs: []FeatureDescription{{Name: "y", Type: fp16ArrayType(1, 4)}},
		},
		MLProgram: &Program{
			Version: 1,
			Functions: map[string]*Function{
				"main": {
					OpSet:  "CoreML8",
					Inputs: []NamedValueType{{Name: "x", Type: fp16TensorVT(1, 4)}},
					BlockSpecializations: map[string]*Block{
						"CoreML8": {
							Operations: []*Operation{
								{
									Type:    "const",
									Outputs: []NamedValueType{{Name: "a", Type: fp16TensorVT(1, 4)}},
									Attributes: map[string]*Value{
										"name": stringAttr("const_a"),
										"val":  blobFileAttr(fp16TensorVT(1, 4), "@model_path/weights/weight.bin", offsetA),
									},
								},
								{
									Type:    "const",
									Outputs: []NamedValueType{{Name: "b", Type: fp16TensorVT(1, 4)}},
									Attributes: map[string]*Value{
										"name": stringAttr("const_b"),
										"val":  blobFileAttr(fp16TensorVT(1, 4), "@model_path/weights/weight.bin", offsetB),
									},
								},
								{
									Type:    "add",
									Outputs: []NamedValueType{{Name: "y", Type: fp16TensorVT(1, 4)}},
									Inputs: map[string]*Argument{
										"x": {Bindings: []Binding{{Name: "a"}}},
										"y": {Bindings: []Binding{{Name: "b"}}},
									},
									Attributes: map[string]*Value{
										"name": stringAttr("add_y"),
									},
								},
							},
							Outputs: []string{"y"},
						},
					},
				},
			},
		},
	})
}

func fp16ArrayType(dims ...int64) *FeatureType {
	return &FeatureType{MultiArrayType: &ArrayFeatureType{Shape: dims, DataType: ArrayDataTypeFloat16}}
}

func fp16TensorVT(dims ...uint64) *ValueType {
	d := make([]Dimension, len(dims))
	for i, v := range dims {
		d[i] = Dimension{Constant: v}
	}
	return &ValueType{TensorType: &TensorType{
		DataType:   DataTypeFloat16,
		Rank:       int64(len(dims)),
		Dimensions: d,
	}}
}

func stringAttr(s string) *Value {
	return &Value{
		Type:      &ValueType{TensorType: &TensorType{DataType: DataTypeString}},
		Immediate: &ImmediateValue{Tensor: &TensorValue{Strings: []string{s}}},
	}
}

func blobFileAttr(vt *ValueType, path string, offset uint64) *Value {
	return &Value{
		Type:     vt,
		BlobFile: &BlobFileValue{FileName: path, Offset: offset},
	}
}

func makeFP16Bytes(vals ...float32) []byte {
	out := make([]byte, len(vals)*2)
	for i, v := range vals {
		binary.LittleEndian.PutUint16(out[i*2:], float32ToFloat16(v))
	}
	return out
}

func float32ToFloat16(f float32) uint16 {
	bits := math.Float32bits(f)
	sign := uint16((bits >> 31) & 1)
	exp := int((bits>>23)&0xFF) - 127
	frac := bits & 0x7FFFFF
	if exp > 15 {
		return (sign << 15) | 0x7C00
	}
	if exp < -14 {
		return sign << 15
	}
	return (sign << 15) | uint16(exp+15)<<10 | uint16(frac>>13)
}
