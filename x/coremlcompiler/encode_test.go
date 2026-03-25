package coremlcompiler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// TestEncodeModelRoundTrip builds a rich Model programmatically, encodes it,
// decodes the bytes, and compares all fields.
func TestEncodeModelRoundTrip(t *testing.T) {
	original := &Model{
		SpecVersion: 8,
		Description: ModelDescription{
			Inputs: []FeatureDescription{
				{Name: "input_ids", Type: &FeatureType{MultiArrayType: &ArrayFeatureType{Shape: []int64{1, 128}, DataType: ArrayDataTypeFloat16}}},
				{Name: "mask", Type: &FeatureType{MultiArrayType: &ArrayFeatureType{Shape: []int64{1, 128}, DataType: ArrayDataTypeFloat32}}},
			},
			Outputs: []FeatureDescription{
				{Name: "logits", Type: &FeatureType{MultiArrayType: &ArrayFeatureType{Shape: []int64{1, 128, 32000}, DataType: ArrayDataTypeFloat16}}},
			},
			States: []FeatureDescription{
				{Name: "k_cache", Type: &FeatureType{MultiArrayType: &ArrayFeatureType{Shape: []int64{1, 32, 128, 64}, DataType: ArrayDataTypeFloat16}}},
			},
		},
		MLProgram: &program{
			Version: 1,
			Functions: map[string]*function{
				"main": {
					OpSet: "ios18",
					Inputs: []namedValueType{
						{
							Name: "x",
							Type: &valueType{TensorType: &tensorType{
								DataType:   DataTypeFloat16,
								Rank:       2,
								Dimensions: []dimension{{Constant: 1}, {Constant: 128}},
							}},
						},
						{
							Name: "k_state",
							Type: &valueType{StateType: &stateType{
								WrappedType: &valueType{TensorType: &tensorType{
									DataType:   DataTypeFloat16,
									Dimensions: []dimension{{Constant: 1}, {Constant: 32}, {Constant: 128}, {Constant: 64}},
								}},
							}},
						},
					},
					BlockSpecializations: map[string]*block{
						"CoreML8": {
							Outputs: []string{"y"},
							Operations: []*operation{
								{
									Type: "const",
									Inputs: map[string]*argument{
										"val": {Bindings: []binding{{
											Value: &value{
												Type: &valueType{TensorType: &tensorType{
													DataType:   DataTypeFloat16,
													Dimensions: []dimension{{Constant: 1}, {Constant: 64}},
												}},
												BlobFile: &blobFileValue{
													FileName: "@model_path/weights/weight.bin",
													Offset:   64,
												},
											},
										}}},
									},
									Outputs: []namedValueType{{
										Name: "y",
										Type: &valueType{TensorType: &tensorType{
											DataType:   DataTypeFloat16,
											Dimensions: []dimension{{Constant: 1}, {Constant: 64}},
										}},
									}},
									Attributes: map[string]*value{
										"name": stringVal("const_y"),
									},
								},
							},
						},
					},
					Attributes: map[string]*value{
						"version": stringVal("1.0"),
					},
				},
			},
			Attributes: map[string]*value{
				"buildInfo": stringVal("test"),
			},
		},
	}

	encoded := EncodeModel(original)
	decoded, err := decodeModel(encoded)
	if err != nil {
		t.Fatalf("decode: %v", err)
	}

	if decoded.SpecVersion != original.SpecVersion {
		t.Errorf("SpecVersion = %d, want %d", decoded.SpecVersion, original.SpecVersion)
	}
	compareDescriptions(t, "Description", decoded.Description, original.Description)
	if original.MLProgram == nil {
		t.Fatal("original MLProgram is nil")
	}
	if decoded.MLProgram == nil {
		t.Fatal("decoded MLProgram is nil")
	}
	comparePrograms(t, decoded.MLProgram, original.MLProgram)
}

// TestEncodeMLProgramRoundTrip focuses on a complex program with multiple
// functions, blob references, immediate values, nested blocks, and attributes.
func TestEncodeMLProgramRoundTrip(t *testing.T) {
	original := &program{
		Version: 1,
		Functions: map[string]*function{
			"main": {
				OpSet: "ios18",
				Inputs: []namedValueType{{
					Name: "x",
					Type: &valueType{TensorType: &tensorType{
						DataType:   DataTypeFloat32,
						Rank:       2,
						Dimensions: []dimension{{Constant: 1}, {Constant: 768}},
					}},
				}},
				BlockSpecializations: map[string]*block{
					"CoreML8": {
						Outputs: []string{"out"},
						Operations: []*operation{
							{
								Type: "const",
								Inputs: map[string]*argument{
									"val": {Bindings: []binding{{
										Value: &value{
											Type: &valueType{TensorType: &tensorType{
												DataType:   DataTypeFloat32,
												Dimensions: []dimension{{Constant: 768}, {Constant: 768}},
											}},
											BlobFile: &blobFileValue{
												FileName: "@model_path/weights/weight.bin",
												Offset:   0,
											},
										},
									}}},
								},
								Outputs: []namedValueType{{
									Name: "w",
									Type: &valueType{TensorType: &tensorType{
										DataType:   DataTypeFloat32,
										Dimensions: []dimension{{Constant: 768}, {Constant: 768}},
									}},
								}},
								Attributes: map[string]*value{
									"name": stringVal("weights_0"),
								},
							},
							{
								Type: "linear",
								Inputs: map[string]*argument{
									"x":      {Bindings: []binding{{Name: "x"}}},
									"weight": {Bindings: []binding{{Name: "w"}}},
									"bias": {Bindings: []binding{{
										Value: &value{
											Type: &valueType{TensorType: &tensorType{
												DataType:   DataTypeFloat32,
												Dimensions: []dimension{{Constant: 768}},
											}},
											Immediate: &immediateValue{Tensor: &tensorValue{
												Floats: []float32{0.1, 0.2, 0.3},
											}},
										},
									}}},
								},
								Outputs: []namedValueType{{
									Name: "out",
									Type: &valueType{TensorType: &tensorType{
										DataType:   DataTypeFloat32,
										Dimensions: []dimension{{Constant: 1}, {Constant: 768}},
									}},
								}},
								Attributes: map[string]*value{
									"name": stringVal("linear_0"),
								},
							},
						},
					},
				},
				Attributes: map[string]*value{
					"description": stringVal("test function"),
				},
			},
			"preprocess": {
				OpSet: "ios18",
				Inputs: []namedValueType{{
					Name: "raw",
					Type: &valueType{TensorType: &tensorType{
						DataType:   DataTypeInt32,
						Dimensions: []dimension{{Constant: 1}, {Constant: 256}},
					}},
				}},
				BlockSpecializations: map[string]*block{
					"CoreML8": {
						Outputs:    []string{"cast_out"},
						Operations: []*operation{},
					},
				},
			},
		},
		Attributes: map[string]*value{
			"buildVersion": stringVal("v1.0.0"),
		},
	}

	encoded := encodeProgram(original)
	decoded, err := decodeProgram(encoded)
	if err != nil {
		t.Fatalf("decode: %v", err)
	}
	comparePrograms(t, decoded, original)
}

// TestEncodeTensorValues tests each tensor value type round-trips correctly.
func TestEncodeTensorValues(t *testing.T) {
	tests := []struct {
		name string
		tv   *tensorValue
	}{
		{
			name: "float32",
			tv:   &tensorValue{Floats: []float32{1.5, -2.25, 0, 3.14}},
		},
		{
			name: "float64",
			tv:   &tensorValue{Doubles: []float64{1.5, -2.25, 0, 3.14159265358979}},
		},
		{
			name: "int32",
			tv:   &tensorValue{Ints: []int32{0, 1, -1, 42, 2147483647}},
		},
		{
			name: "int64",
			tv:   &tensorValue{Longs: []int64{0, 1, -1, 42, 9223372036854775807}},
		},
		{
			name: "bool",
			tv:   &tensorValue{Bools: []bool{true, false, true, true, false}},
		},
		{
			name: "string",
			tv:   &tensorValue{Strings: []string{"hello", "world", "", "test with spaces"}},
		},
		{
			name: "bytes",
			tv:   &tensorValue{Bytes: []byte{0x00, 0x01, 0xFE, 0xFF, 0xDE, 0xAD}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &value{
				Type: &valueType{TensorType: &tensorType{DataType: DataTypeFloat32}},
				Immediate: &immediateValue{
					Tensor: tt.tv,
				},
			}
			encoded := encodeValue(v)
			decoded, err := decodeValue(encoded)
			if err != nil {
				t.Fatalf("decode: %v", err)
			}
			if decoded.Immediate == nil || decoded.Immediate.Tensor == nil {
				t.Fatal("decoded value has no immediate tensor")
			}
			compareTensorValues(t, decoded.Immediate.Tensor, tt.tv)
		})
	}
}

// TestEncodeMapEntries verifies that map encoding works for functions,
// attributes, and inputs maps with multiple entries.
func TestEncodeMapEntries(t *testing.T) {
	prog := &program{
		Version: 1,
		Functions: map[string]*function{
			"main":       {OpSet: "ios18", BlockSpecializations: map[string]*block{}, Attributes: map[string]*value{}},
			"preprocess": {OpSet: "ios17", BlockSpecializations: map[string]*block{}, Attributes: map[string]*value{}},
			"postprocess": {
				OpSet: "ios18",
				BlockSpecializations: map[string]*block{
					"CoreML8": {
						Outputs: []string{"out"},
						Operations: []*operation{{
							Type: "relu",
							Inputs: map[string]*argument{
								"x":     {Bindings: []binding{{Name: "input_0"}}},
								"alpha": {Bindings: []binding{{Name: "alpha_0"}}},
								"beta":  {Bindings: []binding{{Name: "beta_0"}}},
							},
							Outputs: []namedValueType{{Name: "out", Type: &valueType{TensorType: &tensorType{DataType: DataTypeFloat32}}}},
							Attributes: map[string]*value{
								"name":  stringVal("relu_0"),
								"label": stringVal("activation"),
							},
						}},
					},
				},
				Attributes: map[string]*value{},
			},
		},
		Attributes: map[string]*value{
			"a": stringVal("1"),
			"b": stringVal("2"),
			"c": stringVal("3"),
		},
	}

	encoded := encodeProgram(prog)
	decoded, err := decodeProgram(encoded)
	if err != nil {
		t.Fatalf("decode: %v", err)
	}

	// Check all function keys present.
	for _, key := range []string{"main", "preprocess", "postprocess"} {
		if _, ok := decoded.Functions[key]; !ok {
			t.Errorf("missing function %q", key)
		}
	}

	// Check all attribute keys present.
	for _, key := range []string{"a", "b", "c"} {
		if _, ok := decoded.Attributes[key]; !ok {
			t.Errorf("missing attribute %q", key)
		}
	}

	// Check operation inputs map.
	pp := decoded.Functions["postprocess"]
	if pp == nil {
		t.Fatal("missing postprocess function")
	}
	blk, ok := pp.BlockSpecializations["CoreML8"]
	if !ok || len(blk.Operations) == 0 {
		t.Fatal("missing CoreML8 block or operations")
	}
	op := blk.Operations[0]
	for _, key := range []string{"x", "alpha", "beta"} {
		if _, ok := op.Inputs[key]; !ok {
			t.Errorf("missing operation input %q", key)
		}
	}
	for _, key := range []string{"name", "label"} {
		if _, ok := op.Attributes[key]; !ok {
			t.Errorf("missing operation attribute %q", key)
		}
	}
}

// TestWriteMLPackage writes an .mlpackage using WriteMLPackage, then reads it
// back using readMLPackage and verifies the model proto and weights.
func TestWriteMLPackage(t *testing.T) {
	dir := t.TempDir()
	pkgDir := filepath.Join(dir, "test.mlpackage")

	modelProto := buildTestModelProto()

	// Create weight directory with a test file.
	weightDir := filepath.Join(dir, "weights-src")
	if err := os.MkdirAll(filepath.Join(weightDir, "weights"), 0o755); err != nil {
		t.Fatal(err)
	}
	weightData := []byte("test-weight-data-1234567890")
	if err := os.WriteFile(filepath.Join(weightDir, "weights", "weight.bin"), weightData, 0o644); err != nil {
		t.Fatal(err)
	}

	if err := WriteMLPackage(pkgDir, modelProto, weightDir); err != nil {
		t.Fatalf("WriteMLPackage: %v", err)
	}

	// Verify the on-disk layout matches Apple's coremltools format.
	for _, path := range []string{
		"Manifest.json",
		"Data/com.apple.CoreML/model.mlmodel",
		"Data/com.apple.CoreML/weights/weight.bin",
	} {
		if _, err := os.Stat(filepath.Join(pkgDir, path)); err != nil {
			t.Errorf("missing %s: %v", path, err)
		}
	}

	// Verify Manifest.json has rootModelIdentifier and UUID keys.
	manifestData, err := os.ReadFile(filepath.Join(pkgDir, "Manifest.json"))
	if err != nil {
		t.Fatalf("read manifest: %v", err)
	}
	var manifest struct {
		RootModelIdentifier string                 `json:"rootModelIdentifier"`
		ItemInfoEntries     map[string]interface{} `json:"itemInfoEntries"`
	}
	if err := json.Unmarshal(manifestData, &manifest); err != nil {
		t.Fatalf("parse manifest: %v", err)
	}
	if manifest.RootModelIdentifier == "" {
		t.Error("manifest missing rootModelIdentifier")
	}
	if _, ok := manifest.ItemInfoEntries[manifest.RootModelIdentifier]; !ok {
		t.Error("rootModelIdentifier not found in itemInfoEntries")
	}

	// Read back via readMLPackage.
	data, wDir, err := readMLPackage(pkgDir)
	if err != nil {
		t.Fatalf("readMLPackage: %v", err)
	}

	if !bytes.Equal(data, modelProto) {
		t.Errorf("model proto mismatch: got %d bytes, want %d", len(data), len(modelProto))
	}

	// Verify weight file exists under the returned weight dir.
	gotWeights, err := os.ReadFile(filepath.Join(wDir, "weights", "weight.bin"))
	if err != nil {
		t.Fatalf("read weights: %v", err)
	}
	if !bytes.Equal(gotWeights, weightData) {
		t.Errorf("weight data mismatch: got %d bytes, want %d", len(gotWeights), len(weightData))
	}
}

// TestEncodeDecodeIdentity decodes buildTestModelProto(), re-encodes with
// EncodeModel, decodes again, and compares the two decoded structs.
func TestEncodeDecodeIdentity(t *testing.T) {
	original := buildTestModelProto()
	model1, err := decodeModel(original)
	if err != nil {
		t.Fatalf("first decode: %v", err)
	}

	reencoded := EncodeModel(model1)
	model2, err := decodeModel(reencoded)
	if err != nil {
		t.Fatalf("second decode: %v", err)
	}

	if model1.SpecVersion != model2.SpecVersion {
		t.Errorf("SpecVersion: %d vs %d", model1.SpecVersion, model2.SpecVersion)
	}
	compareDescriptions(t, "Description", model1.Description, model2.Description)
	if model1.MLProgram == nil || model2.MLProgram == nil {
		t.Fatal("MLProgram is nil in one of the models")
	}
	comparePrograms(t, model1.MLProgram, model2.MLProgram)
}

// Comparison helpers.

func compareDescriptions(t *testing.T, prefix string, got, want ModelDescription) {
	t.Helper()
	compareFeatureList(t, prefix+".Inputs", got.Inputs, want.Inputs)
	compareFeatureList(t, prefix+".Outputs", got.Outputs, want.Outputs)
	compareFeatureList(t, prefix+".States", got.States, want.States)
}

func compareFeatureList(t *testing.T, prefix string, got, want []FeatureDescription) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("%s: len = %d, want %d", prefix, len(got), len(want))
		return
	}
	for i := range got {
		if got[i].Name != want[i].Name {
			t.Errorf("%s[%d].Name = %q, want %q", prefix, i, got[i].Name, want[i].Name)
		}
	}
}

func comparePrograms(t *testing.T, got, want *program) {
	t.Helper()
	if got.Version != want.Version {
		t.Errorf("program.Version = %d, want %d", got.Version, want.Version)
	}
	if len(got.Functions) != len(want.Functions) {
		t.Errorf("program.Functions: len = %d, want %d", len(got.Functions), len(want.Functions))
	}
	for name, wantFn := range want.Functions {
		gotFn, ok := got.Functions[name]
		if !ok {
			t.Errorf("missing function %q", name)
			continue
		}
		compareFunctions(t, name, gotFn, wantFn)
	}
	compareValueMaps(t, "program.Attributes", got.Attributes, want.Attributes)
}

func compareFunctions(t *testing.T, name string, got, want *function) {
	t.Helper()
	prefix := fmt.Sprintf("function[%s]", name)
	if got.OpSet != want.OpSet {
		t.Errorf("%s.OpSet = %q, want %q", prefix, got.OpSet, want.OpSet)
	}
	if len(got.Inputs) != len(want.Inputs) {
		t.Errorf("%s.Inputs: len = %d, want %d", prefix, len(got.Inputs), len(want.Inputs))
	} else {
		for i := range got.Inputs {
			compareNamedValueTypes(t, fmt.Sprintf("%s.Inputs[%d]", prefix, i), got.Inputs[i], want.Inputs[i])
		}
	}
	if len(got.BlockSpecializations) != len(want.BlockSpecializations) {
		t.Errorf("%s.BlockSpecializations: len = %d, want %d", prefix, len(got.BlockSpecializations), len(want.BlockSpecializations))
	}
	for bname, wantBlk := range want.BlockSpecializations {
		gotBlk, ok := got.BlockSpecializations[bname]
		if !ok {
			t.Errorf("%s: missing block %q", prefix, bname)
			continue
		}
		compareBlocks(t, fmt.Sprintf("%s.Block[%s]", prefix, bname), gotBlk, wantBlk)
	}
	compareValueMaps(t, prefix+".Attributes", got.Attributes, want.Attributes)
}

func compareBlocks(t *testing.T, prefix string, got, want *block) {
	t.Helper()
	if !reflect.DeepEqual(got.Outputs, want.Outputs) {
		t.Errorf("%s.Outputs = %v, want %v", prefix, got.Outputs, want.Outputs)
	}
	if len(got.Inputs) != len(want.Inputs) {
		t.Errorf("%s.Inputs: len = %d, want %d", prefix, len(got.Inputs), len(want.Inputs))
	}
	if len(got.Operations) != len(want.Operations) {
		t.Errorf("%s.Operations: len = %d, want %d", prefix, len(got.Operations), len(want.Operations))
		return
	}
	for i := range got.Operations {
		compareOperations(t, fmt.Sprintf("%s.Op[%d]", prefix, i), got.Operations[i], want.Operations[i])
	}
}

func compareOperations(t *testing.T, prefix string, got, want *operation) {
	t.Helper()
	if got.Type != want.Type {
		t.Errorf("%s.Type = %q, want %q", prefix, got.Type, want.Type)
	}
	if len(got.Outputs) != len(want.Outputs) {
		t.Errorf("%s.Outputs: len = %d, want %d", prefix, len(got.Outputs), len(want.Outputs))
	} else {
		for i := range got.Outputs {
			compareNamedValueTypes(t, fmt.Sprintf("%s.Outputs[%d]", prefix, i), got.Outputs[i], want.Outputs[i])
		}
	}
	if len(got.Inputs) != len(want.Inputs) {
		t.Errorf("%s.Inputs: len = %d, want %d", prefix, len(got.Inputs), len(want.Inputs))
	}
	for key, wantArg := range want.Inputs {
		gotArg, ok := got.Inputs[key]
		if !ok {
			t.Errorf("%s: missing input %q", prefix, key)
			continue
		}
		compareArguments(t, fmt.Sprintf("%s.Inputs[%s]", prefix, key), gotArg, wantArg)
	}
	compareValueMaps(t, prefix+".Attributes", got.Attributes, want.Attributes)
	if len(got.Blocks) != len(want.Blocks) {
		t.Errorf("%s.Blocks: len = %d, want %d", prefix, len(got.Blocks), len(want.Blocks))
	} else {
		for i := range got.Blocks {
			compareBlocks(t, fmt.Sprintf("%s.Blocks[%d]", prefix, i), got.Blocks[i], want.Blocks[i])
		}
	}
}

func compareArguments(t *testing.T, prefix string, got, want *argument) {
	t.Helper()
	if len(got.Bindings) != len(want.Bindings) {
		t.Errorf("%s.Bindings: len = %d, want %d", prefix, len(got.Bindings), len(want.Bindings))
		return
	}
	for i := range got.Bindings {
		bp := fmt.Sprintf("%s.Bindings[%d]", prefix, i)
		if got.Bindings[i].Name != want.Bindings[i].Name {
			t.Errorf("%s.Name = %q, want %q", bp, got.Bindings[i].Name, want.Bindings[i].Name)
		}
		if (got.Bindings[i].Value == nil) != (want.Bindings[i].Value == nil) {
			t.Errorf("%s.Value: nil mismatch", bp)
		} else if got.Bindings[i].Value != nil {
			compareValues(t, bp+".Value", got.Bindings[i].Value, want.Bindings[i].Value)
		}
	}
}

func compareValues(t *testing.T, prefix string, got, want *value) {
	t.Helper()
	if (got.BlobFile == nil) != (want.BlobFile == nil) {
		t.Errorf("%s.BlobFile: nil mismatch", prefix)
	} else if got.BlobFile != nil {
		if got.BlobFile.FileName != want.BlobFile.FileName {
			t.Errorf("%s.BlobFile.FileName = %q, want %q", prefix, got.BlobFile.FileName, want.BlobFile.FileName)
		}
		if got.BlobFile.Offset != want.BlobFile.Offset {
			t.Errorf("%s.BlobFile.Offset = %d, want %d", prefix, got.BlobFile.Offset, want.BlobFile.Offset)
		}
	}
	if (got.Immediate == nil) != (want.Immediate == nil) {
		t.Errorf("%s.Immediate: nil mismatch", prefix)
	} else if got.Immediate != nil {
		if (got.Immediate.Tensor == nil) != (want.Immediate.Tensor == nil) {
			t.Errorf("%s.Immediate.Tensor: nil mismatch", prefix)
		} else if got.Immediate.Tensor != nil {
			compareTensorValues(t, got.Immediate.Tensor, want.Immediate.Tensor)
		}
	}
}

func compareTensorValues(t *testing.T, got, want *tensorValue) {
	t.Helper()
	if !reflect.DeepEqual(got.Floats, want.Floats) {
		t.Errorf("Floats = %v, want %v", got.Floats, want.Floats)
	}
	if !reflect.DeepEqual(got.Doubles, want.Doubles) {
		t.Errorf("Doubles = %v, want %v", got.Doubles, want.Doubles)
	}
	if !reflect.DeepEqual(got.Ints, want.Ints) {
		t.Errorf("Ints = %v, want %v", got.Ints, want.Ints)
	}
	if !reflect.DeepEqual(got.Longs, want.Longs) {
		t.Errorf("Longs = %v, want %v", got.Longs, want.Longs)
	}
	if !reflect.DeepEqual(got.Bools, want.Bools) {
		t.Errorf("Bools = %v, want %v", got.Bools, want.Bools)
	}
	if !reflect.DeepEqual(got.Strings, want.Strings) {
		t.Errorf("Strings = %v, want %v", got.Strings, want.Strings)
	}
	if !bytes.Equal(got.Bytes, want.Bytes) {
		t.Errorf("Bytes = %x, want %x", got.Bytes, want.Bytes)
	}
}

func compareNamedValueTypes(t *testing.T, prefix string, got, want namedValueType) {
	t.Helper()
	if got.Name != want.Name {
		t.Errorf("%s.Name = %q, want %q", prefix, got.Name, want.Name)
	}
	if (got.Type == nil) != (want.Type == nil) {
		t.Errorf("%s.Type: nil mismatch", prefix)
		return
	}
	if got.Type == nil {
		return
	}
	compareValueTypes(t, prefix+".Type", got.Type, want.Type)
}

func compareValueTypes(t *testing.T, prefix string, got, want *valueType) {
	t.Helper()
	if (got.TensorType == nil) != (want.TensorType == nil) {
		t.Errorf("%s.TensorType: nil mismatch", prefix)
	} else if got.TensorType != nil {
		if got.TensorType.DataType != want.TensorType.DataType {
			t.Errorf("%s.TensorType.DataType = %v, want %v", prefix, got.TensorType.DataType, want.TensorType.DataType)
		}
		if got.TensorType.Rank != want.TensorType.Rank {
			t.Errorf("%s.TensorType.Rank = %d, want %d", prefix, got.TensorType.Rank, want.TensorType.Rank)
		}
		if !reflect.DeepEqual(got.TensorType.Dimensions, want.TensorType.Dimensions) {
			t.Errorf("%s.TensorType.Dimensions = %v, want %v", prefix, got.TensorType.Dimensions, want.TensorType.Dimensions)
		}
	}
	if (got.StateType == nil) != (want.StateType == nil) {
		t.Errorf("%s.StateType: nil mismatch", prefix)
	} else if got.StateType != nil {
		if (got.StateType.WrappedType == nil) != (want.StateType.WrappedType == nil) {
			t.Errorf("%s.StateType.WrappedType: nil mismatch", prefix)
		} else if got.StateType.WrappedType != nil {
			compareValueTypes(t, prefix+".StateType.WrappedType", got.StateType.WrappedType, want.StateType.WrappedType)
		}
	}
}

func compareValueMaps(t *testing.T, prefix string, got, want map[string]*value) {
	t.Helper()
	// Treat nil and empty as equivalent.
	if len(got) != len(want) {
		t.Errorf("%s: len = %d, want %d", prefix, len(got), len(want))
		return
	}
	for key, wantVal := range want {
		gotVal, ok := got[key]
		if !ok {
			t.Errorf("%s: missing key %q", prefix, key)
			continue
		}
		compareValues(t, fmt.Sprintf("%s[%s]", prefix, key), gotVal, wantVal)
	}
}
