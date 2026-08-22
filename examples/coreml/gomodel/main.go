// Command gomodel authors, packages, compiles, and runs a CoreML model
// entirely from Go — no Python, no coremltools, no Xcode.
//
// The pipeline, every step in this process:
//
//  1. Define an MIL program in Go (y = relu(x·Wᵀ + b)) and write its
//     weights as a MILBlob Storage v2 weight.bin.
//  2. Encode it to CoreML's protobuf with x/coremlcompiler.EncodeModel.
//  3. Write an Apple-compatible .mlpackage with WriteMLPackage.
//  4. Compile it to .mlmodelc — by default with x/coremlcompiler's own
//     pure-Go compiler; -applecompile uses [MLModel compileModelAtURL:]
//     instead, proving Apple's toolchain accepts the Go-authored package.
//  5. Load it with CoreML and run a prediction.
//
// The prediction is verified against the same math computed in Go.
//
//	go run ./examples/coreml/gomodel
//	go run ./examples/coreml/gomodel -applecompile
package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"time"
	"unsafe"

	"github.com/tmc/apple/coreml"
	cc "github.com/tmc/apple/x/coremlcompiler"
)

const (
	inDim  = 4
	outDim = 3
)

func main() {
	log.SetFlags(0)
	appleCompile := flag.Bool("applecompile", false, "compile with CoreML's compileModelAtURL instead of the pure-Go compiler")
	keep := flag.String("keep", "", "keep build products in this directory instead of a temp dir")
	flag.Parse()

	dir := *keep
	if dir == "" {
		var err error
		dir, err = os.MkdirTemp("", "gomodel-*")
		if err != nil {
			log.Fatalf("gomodel: %v", err)
		}
		defer os.RemoveAll(dir)
	} else if err := os.MkdirAll(dir, 0o755); err != nil {
		log.Fatalf("gomodel: %v", err)
	}

	// The model's parameters, authored right here.
	weight := []float32{ // [outDim, inDim], row-major
		0.5, -1.0, 0.25, 2.0,
		1.5, 0.0, -0.5, 1.0,
		-2.0, 0.75, 1.0, -0.25,
	}
	bias := []float32{0.1, -0.2, 0.3}

	// 1. Write the weights as a MILBlob Storage v2 weight.bin — the same
	// format coremltools emits — and define the MIL program as a plain Go
	// value referencing them.
	blob, offsets := cc.WriteMILBlob([]cc.BlobEntry{
		{DType: cc.BlobDataTypeFloat32, Data: floatBytes(weight)},
		{DType: cc.BlobDataTypeFloat32, Data: floatBytes(bias)},
	})
	weightFile := filepath.Join(dir, "weight.bin")
	if err := os.WriteFile(weightFile, blob, 0o644); err != nil {
		log.Fatalf("gomodel: %v", err)
	}
	model := buildModel(offsets[0], offsets[1])

	// 2. Encode to CoreML's protobuf.
	proto := cc.EncodeModel(model)
	log.Printf("1. authored + encoded: %d-byte mlprogram proto (relu(x·Wᵀ+b), %dx%d), %d-byte weight.bin",
		len(proto), inDim, outDim, len(blob))

	// 3. Write the .mlpackage.
	pkg := filepath.Join(dir, "gomodel.mlpackage")
	if err := cc.WriteMLPackage(pkg, proto, weightFile); err != nil {
		log.Fatalf("gomodel: write package: %v", err)
	}
	log.Printf("2. wrote %s", pkg)

	// 4. Compile.
	var modelc string
	start := time.Now()
	if *appleCompile {
		compiled, err := cc.CompileMLModelAtURL(pkg)
		if err != nil {
			log.Fatalf("gomodel: apple compile: %v", err)
		}
		modelc = compiled
		log.Printf("3. compiled with CoreML's compileModelAtURL in %v — Apple's toolchain accepts the Go-authored package", time.Since(start).Round(time.Millisecond))
	} else {
		modelc = filepath.Join(dir, "gomodel.mlmodelc")
		if err := cc.Compile(pkg, modelc); err != nil {
			log.Fatalf("gomodel: pure-Go compile: %v", err)
		}
		log.Printf("3. compiled with the pure-Go compiler in %v (no Apple tool invoked)", time.Since(start).Round(time.Millisecond))
	}

	// 5. Load and predict.
	m, err := cc.LoadCoreMLModel(modelc)
	if err != nil {
		log.Fatalf("gomodel: load: %v", err)
	}
	defer m.Close()

	x := []float32{1.0, 2.0, -3.0, 0.5}
	out, err := m.Predict([]cc.PredictInput{{
		Name:    "x",
		Data:    unsafe.Pointer(&x[0]),
		Shape:   []int{1, inDim},
		Strides: []int{inDim, 1},
		DType:   coreml.MLMultiArrayDataTypeFloat32,
	}}, "y")
	if err != nil {
		log.Fatalf("gomodel: predict: %v", err)
	}
	got := unsafe.Slice((*float32)(unsafe.Pointer(&out.Bytes[0])), outDim)
	log.Printf("4. CoreML prediction for x=%v: %v (output shape %v)", x, got, out.Shape)

	// Verify against the same math in Go.
	want := reference(x, weight, bias)
	for i := range want {
		if math.Abs(float64(got[i])-float64(want[i])) > 1e-4 {
			log.Fatalf("gomodel: mismatch at %d: CoreML %v, Go reference %v", i, got, want)
		}
	}
	log.Printf("5. verified: CoreML output matches the Go reference %v", want)
	fmt.Println("ok: model authored, packaged, compiled, and executed — zero Python, zero Xcode")
}

// floatBytes returns the little-endian bytes of a float32 slice.
func floatBytes(f []float32) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&f[0])), len(f)*4)
}

// buildModel constructs the MIL program for y = relu(x·Wᵀ + b) with the
// weights referenced from weight.bin at the given blob offsets.
func buildModel(weightOff, biasOff uint64) *cc.Model {
	f32 := func(dims ...int64) *cc.ValueType {
		ds := make([]cc.Dimension, len(dims))
		for i, d := range dims {
			ds[i] = cc.Dimension{Constant: uint64(d)}
		}
		return &cc.ValueType{TensorType: &cc.TensorType{
			DataType:   cc.DataTypeFloat32,
			Rank:       int64(len(dims)),
			Dimensions: ds,
		}}
	}
	constOp := func(name string, vt *cc.ValueType, offset uint64) *cc.Operation {
		// In the protobuf spec (and Apple's parser) a const op carries its
		// value as the "val" attribute, not as an input; the MIL-text
		// emitter accepts either and prints both as attributes.
		return &cc.Operation{
			Type:    "const",
			Outputs: []cc.NamedValueType{{Name: name, Type: vt}},
			Attributes: map[string]*cc.Value{
				"name": stringVal("const_" + name),
				"val": {
					Type: vt,
					BlobFile: &cc.BlobFileValue{
						FileName: "@model_path/weights/weight.bin",
						Offset:   offset,
					},
				},
			},
		}
	}
	ref := func(name string) *cc.Argument {
		return &cc.Argument{Bindings: []cc.Binding{{Name: name}}}
	}

	// The protobuf layer names opsets "CoreMLn" (CoreML7 ≙ ios18 in MIL
	// text; the pure-Go compiler does that rewrite when emitting model.mil).
	// Apple's validator requires the spec version to match: 9 for CoreML7.
	return &cc.Model{
		SpecVersion: 9,
		Description: cc.ModelDescription{
			Inputs: []cc.FeatureDescription{{
				Name: "x",
				Type: &cc.FeatureType{MultiArrayType: &cc.ArrayFeatureType{
					Shape: []int64{1, inDim}, DataType: cc.ArrayDataTypeFloat32,
				}},
			}},
			Outputs: []cc.FeatureDescription{{
				Name: "y",
				Type: &cc.FeatureType{MultiArrayType: &cc.ArrayFeatureType{
					Shape: []int64{1, outDim}, DataType: cc.ArrayDataTypeFloat32,
				}},
			}},
		},
		MLProgram: &cc.Program{
			Version: 1,
			Functions: map[string]*cc.Function{
				"main": {
					OpSet:  "CoreML7",
					Inputs: []cc.NamedValueType{{Name: "x", Type: f32(1, inDim)}},
					// The specialization key must equal the function's opset
					// name; Apple's parser rejects anything else.
					BlockSpecializations: map[string]*cc.Block{
						"CoreML7": {
							Outputs: []string{"y"},
							Operations: []*cc.Operation{
								constOp("w", f32(outDim, inDim), weightOff),
								constOp("b", f32(outDim), biasOff),
								{
									Type: "linear",
									Inputs: map[string]*cc.Argument{
										"x":      ref("x"),
										"weight": ref("w"),
										"bias":   ref("b"),
									},
									Outputs:    []cc.NamedValueType{{Name: "pre", Type: f32(1, outDim)}},
									Attributes: map[string]*cc.Value{"name": stringVal("linear_0")},
								},
								{
									Type:       "relu",
									Inputs:     map[string]*cc.Argument{"x": ref("pre")},
									Outputs:    []cc.NamedValueType{{Name: "y", Type: f32(1, outDim)}},
									Attributes: map[string]*cc.Value{"name": stringVal("relu_0")},
								},
							},
						},
					},
				},
			},
		},
	}
}

func stringVal(s string) *cc.Value {
	return &cc.Value{
		Type:      &cc.ValueType{TensorType: &cc.TensorType{DataType: cc.DataTypeString}},
		Immediate: &cc.ImmediateValue{Tensor: &cc.TensorValue{Strings: []string{s}}},
	}
}

// reference computes relu(x·Wᵀ + b) in Go.
func reference(x, weight, bias []float32) []float32 {
	out := make([]float32, outDim)
	for o := range outDim {
		acc := bias[o]
		for i := range inDim {
			acc += x[i] * weight[o*inDim+i]
		}
		if acc < 0 {
			acc = 0
		}
		out[o] = acc
	}
	return out
}
