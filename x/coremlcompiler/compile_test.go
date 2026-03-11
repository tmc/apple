package coremlcompiler

import (
	"encoding/binary"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestProtoRoundTrip constructs a minimal mlprogram model proto by hand,
// decodes it, and verifies the parsed structure.
func TestProtoRoundTrip(t *testing.T) {
	// Build a minimal Model proto with an mlprogram.
	proto := buildTestModelProto()

	model, err := decodeModel(proto)
	if err != nil {
		t.Fatalf("decode: %v", err)
	}

	if model.SpecVersion != 8 {
		t.Errorf("spec version = %d, want 8", model.SpecVersion)
	}
	if len(model.Description.Inputs) != 1 {
		t.Fatalf("inputs = %d, want 1", len(model.Description.Inputs))
	}
	if model.Description.Inputs[0].Name != "x" {
		t.Errorf("input name = %q, want %q", model.Description.Inputs[0].Name, "x")
	}
	if model.MLProgram == nil {
		t.Fatal("mlprogram is nil")
	}
	if model.MLProgram.Version != 1 {
		t.Errorf("program version = %d, want 1", model.MLProgram.Version)
	}

	fn, ok := model.MLProgram.Functions["main"]
	if !ok {
		t.Fatal("no main function")
	}
	if len(fn.Inputs) != 1 {
		t.Fatalf("function inputs = %d, want 1", len(fn.Inputs))
	}
	if fn.Inputs[0].Name != "x" {
		t.Errorf("function input name = %q, want %q", fn.Inputs[0].Name, "x")
	}
	if fn.OpSet != "ios18" {
		t.Errorf("opset = %q, want %q", fn.OpSet, "ios18")
	}

	block, ok := fn.BlockSpecializations["CoreML8"]
	if !ok {
		t.Fatal("no CoreML8 block specialization")
	}
	if len(block.Operations) != 1 {
		t.Fatalf("operations = %d, want 1", len(block.Operations))
	}
	if block.Operations[0].Type != "const" {
		t.Errorf("op type = %q, want %q", block.Operations[0].Type, "const")
	}
	if len(block.Outputs) != 1 || block.Outputs[0] != "y" {
		t.Errorf("block outputs = %v, want [y]", block.Outputs)
	}
}

// TestMILTextEmit verifies MIL text generation from a parsed program.
func TestMILTextEmit(t *testing.T) {
	prog := &program{
		Version: 1,
		Functions: map[string]*function{
			"main": {
				OpSet: "ios18",
				Inputs: []namedValueType{
					{
						Name: "x",
						Type: &valueType{
							TensorType: &tensorType{
								DataType:   DataTypeFloat16,
								Dimensions: []dimension{{Constant: 1}, {Constant: 64}},
							},
						},
					},
				},
				BlockSpecializations: map[string]*block{
					"CoreML8": {
						Operations: []*operation{
							{
								Type: "const",
								Outputs: []namedValueType{
									{
										Name: "y",
										Type: &valueType{
											TensorType: &tensorType{
												DataType:   DataTypeFloat16,
												Dimensions: []dimension{{Constant: 1}, {Constant: 64}},
											},
										},
									},
								},
								Inputs: map[string]*argument{
									"val": {
										Bindings: []binding{
											{
												Value: &value{
													Type: &valueType{
														TensorType: &tensorType{
															DataType:   DataTypeFloat16,
															Dimensions: []dimension{{Constant: 1}, {Constant: 64}},
														},
													},
													BlobFile: &blobFileValue{
														FileName: "@model_path/weights/weight.bin",
														Offset:   64,
													},
												},
											},
										},
									},
								},
								Attributes: map[string]*value{
									"name": {
										Type: &valueType{
											TensorType: &tensorType{DataType: DataTypeString},
										},
										Immediate: &immediateValue{
											Tensor: &tensorValue{
												Strings: []string{"const_y"},
											},
										},
									},
								},
							},
						},
						Outputs: []string{"y"},
					},
				},
			},
		},
	}

	text := emitMILText(prog)

	// Verify key elements are present in the real Apple MIL format.
	checks := []string{
		"program(1.3)",
		"func main<ios18>(",
		"tensor<fp16, [1, 64]>",
		"const(",
		"BLOBFILE(",
		`tensor<string, []>("@model_path/weights/weight.bin")`,
		"tensor<uint64, []>(64)",
		"-> (y)",
	}
	for _, check := range checks {
		if !strings.Contains(text, check) {
			t.Errorf("MIL text missing %q\nfull text:\n%s", check, text)
		}
	}
}

// TestCoreMLDataBuild verifies coremldata.bin format matches real Apple output.
func TestCoreMLDataBuild(t *testing.T) {
	descProto := []byte{0x0A, 0x01, 0x78} // minimal proto: field 1, string "x"
	model := &Model{
		SpecVersion:    8,
		descriptionRaw: descProto,
	}

	data := buildCoreMLData(model)

	// Check magic at offset 0.
	magic := binary.LittleEndian.Uint32(data[0x00:])
	if magic != 502 {
		t.Errorf("magic = %d, want 502", magic)
	}

	// Check version at offset 4.
	version := binary.LittleEndian.Uint32(data[0x04:])
	if version != 8 {
		t.Errorf("version = %d, want 8", version)
	}

	// Check compute type string length at offset 0x1c.
	strLen := binary.LittleEndian.Uint32(data[0x1c:])
	if strLen != 7 {
		t.Errorf("compute type len = %d, want 7", strLen)
	}

	// Check compute type at offset 0x24.
	computeType := string(data[0x24:0x2b])
	if computeType != "generic" {
		t.Errorf("compute type = %q, want %q", computeType, "generic")
	}

	// Check proto length at offset 0x4b.
	protoLen := binary.LittleEndian.Uint16(data[0x4b:])
	if int(protoLen) != len(descProto) {
		t.Errorf("proto len = %d, want %d", protoLen, len(descProto))
	}

	// Check proto data starts at offset 0x53.
	for i, b := range descProto {
		if data[0x53+i] != b {
			t.Errorf("proto byte %d = 0x%02x, want 0x%02x", i, data[0x53+i], b)
		}
	}

	// Check trailer: magic repeated at last 16 bytes.
	trailerMagic := binary.LittleEndian.Uint32(data[len(data)-16:])
	if trailerMagic != 502 {
		t.Errorf("trailer magic = %d, want 502", trailerMagic)
	}

	// Check total size.
	wantSize := 0x53 + len(descProto) + 16
	if len(data) != wantSize {
		t.Errorf("total size = %d, want %d", len(data), wantSize)
	}
}

// TestCompileMLProgram does an end-to-end compile with a synthetic model.
func TestCompileMLProgram(t *testing.T) {
	// Create a temp directory with weight file.
	tmpDir := t.TempDir()
	weightDir := filepath.Join(tmpDir, "weights")
	if err := os.MkdirAll(weightDir, 0o755); err != nil {
		t.Fatal(err)
	}

	// Write a minimal weight file.
	weightData := make([]byte, 128+64*2) // header + 64 fp16 values
	weightData[0] = 0x01
	weightData[4] = 0x02
	binary.LittleEndian.PutUint32(weightData[64:], 0xDEADBEEF)
	if err := os.WriteFile(filepath.Join(weightDir, "weight.bin"), weightData, 0o644); err != nil {
		t.Fatal(err)
	}

	// Build model proto.
	proto := buildTestModelProto()

	// Compile.
	outputDir := filepath.Join(tmpDir, "model.mlmodelc")
	if err := CompileMLProgram(proto, tmpDir, outputDir); err != nil {
		t.Fatalf("compile: %v", err)
	}

	// Verify output files exist.
	for _, name := range []string{"model.mil", "coremldata.bin", filepath.Join("analytics", "coremldata.bin"), "metadata.json"} {
		path := filepath.Join(outputDir, name)
		if _, err := os.Stat(path); err != nil {
			t.Errorf("missing %s: %v", name, err)
		}
	}

	// Verify weight file was copied.
	weightPath := filepath.Join(outputDir, "weights", "weight.bin")
	if _, err := os.Stat(weightPath); err != nil {
		t.Errorf("missing weights/weight.bin: %v", err)
	}

	// Read and verify model.mil content.
	milData, err := os.ReadFile(filepath.Join(outputDir, "model.mil"))
	if err != nil {
		t.Fatal(err)
	}
	milText := string(milData)
	if !strings.Contains(milText, "program(") {
		t.Errorf("model.mil missing program header")
	}
	if !strings.Contains(milText, "func main<ios18>") {
		t.Errorf("model.mil missing main function")
	}
}

func TestCompileMILText(t *testing.T) {
	tmpDir := t.TempDir()
	weightRoot := filepath.Join(tmpDir, "weights-root")
	if err := os.MkdirAll(filepath.Join(weightRoot, "weights"), 0o755); err != nil {
		t.Fatal(err)
	}
	weightData := make([]byte, 128+64*2)
	weightData[0] = 0x01
	weightData[4] = 0x02
	binary.LittleEndian.PutUint32(weightData[64:], 0xDEADBEEF)
	if err := os.WriteFile(filepath.Join(weightRoot, "weights", "weight.bin"), weightData, 0o644); err != nil {
		t.Fatal(err)
	}

	milText := strings.TrimSpace(`
program(1.3) {
    func main<ios18>(tensor<fp16, [1, 64]> x, state<tensor<fp16, [1, 1, 4, 64]>> k_state) {
        tensor<fp16, [1, 1, 4, 64]> cached = read_state(state = k_state)[name = tensor<string, []>("read_k")];
    } -> (cached);
}
`)
	desc := ModelDescription{
		Inputs: []FeatureDescription{{
			Name: "x",
			Type: &FeatureType{MultiArrayType: &ArrayFeatureType{Shape: []int64{1, 64}, DataType: ArrayDataTypeFloat16}},
		}},
		Outputs: []FeatureDescription{{
			Name: "cached",
			Type: &FeatureType{MultiArrayType: &ArrayFeatureType{Shape: []int64{1, 1, 4, 64}, DataType: ArrayDataTypeFloat16}},
		}},
		States: []FeatureDescription{{
			Name: "k_state",
			Type: &FeatureType{MultiArrayType: &ArrayFeatureType{Shape: []int64{1, 1, 4, 64}, DataType: ArrayDataTypeFloat16}},
		}},
	}

	outputDir := filepath.Join(tmpDir, "model.mlmodelc")
	if err := CompileMILText(milText, 8, desc, weightRoot, outputDir); err != nil {
		t.Fatalf("CompileMILText: %v", err)
	}

	for _, name := range []string{"model.mil", "coremldata.bin", filepath.Join("analytics", "coremldata.bin"), "metadata.json", filepath.Join("weights", "weight.bin")} {
		path := filepath.Join(outputDir, name)
		if _, err := os.Stat(path); err != nil {
			t.Errorf("missing %s: %v", name, err)
		}
	}

	gotMIL, err := os.ReadFile(filepath.Join(outputDir, "model.mil"))
	if err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(string(gotMIL)) != milText {
		t.Fatalf("compiled model.mil mismatch\n got: %q\nwant: %q", string(gotMIL), milText)
	}

	model, err := decodeModelDescription(encodeModelDescription(desc))
	if err != nil {
		t.Fatalf("decodeModelDescription: %v", err)
	}
	if len(model.States) != 1 || model.States[0].Name != "k_state" {
		t.Fatalf("state description round trip = %+v", model.States)
	}
}

// buildTestModelProto constructs a minimal Model proto with an mlprogram.
func buildTestModelProto() []byte {
	// We build the proto bottom-up.

	// ConstantDimension { size = 1 }
	dim1 := encodeVarint(1<<3|wireVarint, encodeVarintVal(1))
	// ConstantDimension { size = 64 }
	dim64 := encodeVarint(1<<3|wireVarint, encodeVarintVal(64))

	// Dimension { constant = dim1 }  (field 1, bytes)
	dimMsg1 := encodeBytes(1, dim1)
	// Dimension { constant = dim64 }
	dimMsg64 := encodeBytes(1, dim64)

	// TensorType { dataType=FLOAT16(10), rank=2, dimensions=[1, 64] }
	tensorType := concatBytes(
		encodeVarint(1<<3|wireVarint, encodeVarintVal(10)), // dataType = FLOAT16
		encodeVarint(2<<3|wireVarint, encodeVarintVal(2)),  // rank = 2
		encodeBytes(3, dimMsg1),                            // dimensions[0]
		encodeBytes(3, dimMsg64),                           // dimensions[1]
	)

	// ValueType { tensorType = ... }
	valueType := encodeBytes(1, tensorType)

	// NamedValueType { name="x", type=valueType }
	nvtX := concatBytes(
		encodeBytes(1, []byte("x")),
		encodeBytes(2, valueType),
	)

	// BlobFileValue { fileName="@model_path/weights/weight.bin", offset=64 }
	blobFile := concatBytes(
		encodeBytes(1, []byte("@model_path/weights/weight.bin")),
		encodeVarint(2<<3|wireVarint, encodeVarintVal(64)),
	)

	// Value { type=valueType, blobFileValue=blobFile }
	val := concatBytes(
		encodeBytes(2, valueType),
		encodeBytes(5, blobFile),
	)

	// Binding { value = val }
	binding := encodeBytes(2, val)

	// Argument { arguments = [binding] }
	argument := encodeBytes(1, binding)

	// Map entry for inputs: key="val", value=argument
	inputMapEntry := concatBytes(
		encodeBytes(1, []byte("val")),
		encodeBytes(2, argument),
	)

	// String type for attribute.
	stringTensorType := encodeVarint(1<<3|wireVarint, encodeVarintVal(2)) // DataType STRING
	stringValueType := encodeBytes(1, stringTensorType)

	// TensorValue with string "const_y"
	repeatedStrings := encodeBytes(1, []byte("const_y"))
	tensorValue := encodeBytes(4, repeatedStrings)
	immediateValue := encodeBytes(1, tensorValue)
	nameVal := concatBytes(
		encodeBytes(2, stringValueType),
		encodeBytes(3, immediateValue),
	)
	attrMapEntry := concatBytes(
		encodeBytes(1, []byte("name")),
		encodeBytes(2, nameVal),
	)

	// NamedValueType for output "y"
	nvtY := concatBytes(
		encodeBytes(1, []byte("y")),
		encodeBytes(2, valueType),
	)

	// Operation { type="const", inputs={val: arg}, outputs=[y], attributes={name: ...} }
	operation := concatBytes(
		encodeBytes(1, []byte("const")), // type
		encodeBytes(2, inputMapEntry),   // inputs map entry
		encodeBytes(3, nvtY),            // outputs
		encodeBytes(5, attrMapEntry),    // attributes map entry
	)

	// Block { operations=[op], outputs=["y"] }
	block := concatBytes(
		encodeBytes(2, []byte("y")), // outputs
		encodeBytes(3, operation),   // operations
	)

	// Function block_specializations map entry: key="CoreML8", value=block
	blockMapEntry := concatBytes(
		encodeBytes(1, []byte("CoreML8")),
		encodeBytes(2, block),
	)

	// Function { inputs=[x], opset="ios18", block_specializations={CoreML8: block} }
	function := concatBytes(
		encodeBytes(1, nvtX),            // inputs
		encodeBytes(2, []byte("ios18")), // opset
		encodeBytes(3, blockMapEntry),   // block_specializations
	)

	// Program functions map entry: key="main", value=function
	funcMapEntry := concatBytes(
		encodeBytes(1, []byte("main")),
		encodeBytes(2, function),
	)

	// Program { version=1, functions={main: function} }
	program := concatBytes(
		encodeVarint(1<<3|wireVarint, encodeVarintVal(1)), // version
		encodeBytes(2, funcMapEntry),                      // functions
	)

	// FeatureDescription { name="x" }
	featureDesc := encodeBytes(1, []byte("x"))

	// ModelDescription { input=[featureDesc] }
	modelDesc := encodeBytes(1, featureDesc)

	// Model { specificationVersion=8, description=modelDesc, mlProgram=program }
	model := concatBytes(
		encodeVarint(1<<3|wireVarint, encodeVarintVal(8)), // specificationVersion
		encodeBytes(2, modelDesc),                         // description
		encodeBytes(502, program),                         // mlProgram (field 502)
	)

	return model
}

// TestMILTextWithState verifies state type formatting for iOS 18+ stateful inference.
func TestMILTextWithState(t *testing.T) {
	prog := &program{
		Version: 1,
		Functions: map[string]*function{
			"main": {
				OpSet: "ios18",
				Inputs: []namedValueType{
					{
						Name: "x",
						Type: &valueType{
							TensorType: &tensorType{
								DataType:   DataTypeFloat16,
								Dimensions: []dimension{{Constant: 1}, {Constant: 8}},
							},
						},
					},
					{
						Name: "state_k",
						Type: &valueType{
							StateType: &stateType{
								WrappedType: &valueType{
									TensorType: &tensorType{
										DataType:   DataTypeFloat16,
										Dimensions: []dimension{{Constant: 1}, {Constant: 1}, {Constant: 128}, {Constant: 64}},
									},
								},
							},
						},
					},
				},
				BlockSpecializations: map[string]*block{
					"CoreML8": {
						Operations: []*operation{
							{
								Type: "read_state",
								Outputs: []namedValueType{
									{
										Name: "cached",
										Type: &valueType{
											TensorType: &tensorType{
												DataType:   DataTypeFloat16,
												Dimensions: []dimension{{Constant: 1}, {Constant: 1}, {Constant: 128}, {Constant: 64}},
											},
										},
									},
								},
								Inputs: map[string]*argument{
									"state": {
										Bindings: []binding{{Name: "state_k"}},
									},
								},
								Attributes: map[string]*value{
									"name": stringVal("read_k"),
								},
							},
						},
						Outputs: []string{"cached"},
					},
				},
			},
		},
	}

	text := emitMILTextWithSpec(prog, 8)

	checks := []string{
		"program(1.3)",
		"state<tensor<fp16, [1, 1, 128, 64]>>",
		"read_state(",
		"state = state_k",
	}
	for _, check := range checks {
		if !strings.Contains(text, check) {
			t.Errorf("MIL text missing %q\nfull text:\n%s", check, text)
		}
	}
}

func stringVal(s string) *value {
	return &value{
		Type: &valueType{
			TensorType: &tensorType{DataType: DataTypeString},
		},
		Immediate: &immediateValue{
			Tensor: &tensorValue{Strings: []string{s}},
		},
	}
}

// TestDataTypeString verifies MIL data type names.
func TestDataTypeString(t *testing.T) {
	tests := []struct {
		dt   DataType
		want string
	}{
		{DataTypeFloat16, "fp16"},
		{DataTypeFloat32, "fp32"},
		{DataTypeInt32, "int32"},
		{DataTypeInt64, "int64"},
		{DataTypeBool, "bool"},
		{DataTypeString, "string"},
	}
	for _, tt := range tests {
		if got := tt.dt.String(); got != tt.want {
			t.Errorf("DataType(%d).String() = %q, want %q", int(tt.dt), got, tt.want)
		}
	}
}

// TestFormatTensorType verifies tensor type formatting.
func TestFormatTensorType(t *testing.T) {
	tests := []struct {
		name string
		tt   *tensorType
		want string
	}{
		{
			name: "scalar fp32",
			tt:   &tensorType{DataType: DataTypeFloat32},
			want: "fp32",
		},
		{
			name: "2D fp16",
			tt: &tensorType{
				DataType:   DataTypeFloat16,
				Dimensions: []dimension{{Constant: 1}, {Constant: 128}},
			},
			want: "tensor<fp16, [1, 128]>",
		},
		{
			name: "4D int32",
			tt: &tensorType{
				DataType:   DataTypeInt32,
				Dimensions: []dimension{{Constant: 1}, {Constant: 3}, {Constant: 224}, {Constant: 224}},
			},
			want: "tensor<int32, [1, 3, 224, 224]>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatTensorType(tt.tt)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
