package coremlcompiler_test

import (
	"fmt"
	"os"

	"github.com/tmc/apple/x/coremlcompiler"
)

func ExampleParseDataType() {
	dt, err := coremlcompiler.ParseDataType("fp16")
	fmt.Println(dt, err)
	// Output: fp16 <nil>
}

func ExampleDataType_ArrayDataType() {
	dt, err := coremlcompiler.DataTypeFloat16.ArrayDataType()
	fmt.Println(dt == coremlcompiler.ArrayDataTypeFloat16, err)
	// Output: true <nil>
}

func ExampleStateFeatureDescription() {
	state, err := coremlcompiler.StateFeatureDescription("cache", "fp16", []int64{1, 8})
	fmt.Println(state.Name, state.Type.StateArrayType.Shape, err)
	// Output: cache [1 8] <nil>
}

func ExampleFloat16bits() {
	fmt.Printf("%04x\n", coremlcompiler.Float16bits(1.5))
	// Output: 3e00
}

func ExampleFloat16frombits() {
	fmt.Println(coremlcompiler.Float16frombits(0x3e00))
	// Output: 1.5
}

func ExampleFloat16Bytes() {
	fmt.Printf("%x\n", coremlcompiler.Float16Bytes([]float32{1, 1.5}))
	// Output: 003c003e
}

func ExampleBlobLayout_String() {
	fmt.Println(coremlcompiler.BlobLayoutSingleFile)
	fmt.Println(coremlcompiler.BlobLayoutFilePerConst)
	// Output:
	// single-file
	// file-per-const
}

func ExampleBuildWeights() {
	files, refs, err := coremlcompiler.BuildWeights([]coremlcompiler.WeightTensor{{
		Name: "w", DType: coremlcompiler.BlobDataTypeFloat16,
		Data: coremlcompiler.Float16Bytes([]float32{1, 2}),
	}}, coremlcompiler.BlobLayoutSingleFile)
	fmt.Println(len(files), refs["w"].Path, refs["w"].Offset, err)
	// Output: 1 @model_path/weights/weight.bin 64 <nil>
}

func ExampleWriteWeightRoot() {
	root, err := os.MkdirTemp("", "weight-root-example-")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(root)
	err = coremlcompiler.WriteWeightRoot(root, []coremlcompiler.WeightFile{{
		Path: "@model_path/weights/w.bin", Blob: []byte{1, 2},
	}})
	fmt.Println(err)
	// Output: <nil>
}

func ExampleAffineDequantization_Validate() {
	q := coremlcompiler.AffineDequantization{
		Shape: []int64{2}, QuantizedType: coremlcompiler.DataTypeInt8,
		OutputType: coremlcompiler.DataTypeFloat32, Scale: []float32{0.5}, ZeroPoint: []int64{0},
	}
	fmt.Println(q.Validate())
	// Output: <nil>
}

func ExampleAffineDequantization_NumElements() {
	q := coremlcompiler.AffineDequantization{Shape: []int64{2, 3}}
	fmt.Println(q.NumElements())
	// Output: 6
}

func ExampleAffineDequantization_Dequantize() {
	q := coremlcompiler.AffineDequantization{
		Shape: []int64{2}, QuantizedType: coremlcompiler.DataTypeInt8,
		OutputType: coremlcompiler.DataTypeFloat32, Scale: []float32{0.5}, ZeroPoint: []int64{0},
	}
	values, err := q.Dequantize([]byte{2, 4})
	fmt.Println(values, err)
	// Output: [1 2] <nil>
}

func ExampleAffineDequantization_Operation() {
	q := coremlcompiler.AffineDequantization{
		Shape: []int64{2}, QuantizedType: coremlcompiler.DataTypeInt8,
		OutputType: coremlcompiler.DataTypeFloat32, Scale: []float32{0.5}, ZeroPoint: []int64{0},
	}
	op, err := q.Operation("w", coremlcompiler.BlobRef{Path: "@model_path/weights/w.bin", Offset: 64})
	fmt.Println(op.Type, op.Outputs[0].Name, err)
	// Output: constexpr_affine_dequantize w <nil>
}

func ExampleQuantizeAffinePerChannel() {
	q, packed, err := coremlcompiler.QuantizeAffinePerChannel([]float32{0, 1}, []int64{1, 2}, 0, coremlcompiler.DataTypeInt8)
	fmt.Println(q.NumElements(), len(packed), err)
	// Output: 2 2 <nil>
}

func ExampleCompileProgram() {
	err := coremlcompiler.CompileProgram(nil, 0, coremlcompiler.ModelDescription{}, "", "unused.mlmodelc")
	fmt.Println(err)
	// Output: coremlcompiler: program is nil
}
