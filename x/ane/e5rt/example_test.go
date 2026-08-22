//go:build darwin

package e5rt_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/tmc/apple/x/ane/e5rt"
	"github.com/tmc/apple/x/ane/mil"
)

// Compile a MIL program and run it on the Neural Engine, from the model
// directory through to the output buffer. Handles are released in reverse.
func Example() {
	lib, err := e5rt.Open()
	if err != nil {
		log.Fatal(err)
	}

	// A 1x1 convolution over 4 channels whose weight matrix is the identity,
	// so the output equals the input and the answer needs no reference.
	const channels, spatial = 4, 4
	weights := make([]float32, channels*channels)
	for i := range channels {
		weights[i*channels+i] = 1
	}
	blob, err := mil.BuildWeightBlob(weights, channels, channels)
	if err != nil {
		log.Fatal(err)
	}

	dir, err := os.MkdirTemp("", "e5rt-example-")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)
	if err := os.MkdirAll(filepath.Join(dir, "weights"), 0o755); err != nil {
		log.Fatal(err)
	}
	milText := mil.GenConvFP16IO(channels, channels, spatial)
	if err := os.WriteFile(filepath.Join(dir, "model.mil"), []byte(milText), 0o644); err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "weights", "weight.bin"), blob, 0o644); err != nil {
		log.Fatal(err)
	}

	// Compile. The configuration is required: a zero one is rejected.
	config, err := lib.CompilerConfigOptionsCreate()
	if err != nil {
		log.Fatal(err)
	}
	defer lib.CompilerConfigOptionsRelease(config)
	if err := lib.CompilerConfigOptionsSetCacheBundleLocation(config, dir); err != nil {
		log.Fatal(err)
	}
	compiler, err := lib.CompilerCreateWithConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	defer lib.CompilerRelease(compiler)

	options, err := lib.CompilerOptionsCreate()
	if err != nil {
		log.Fatal(err)
	}
	defer lib.CompilerOptionsRelease(options)
	if err := lib.CompilerOptionsSetComputeDeviceTypesMask(options, e5rt.ComputeDeviceANE); err != nil {
		log.Fatal(err)
	}

	library, err := lib.CompilerCompile(compiler, filepath.Join(dir, "model.mil"), options)
	if err != nil {
		log.Fatal(err)
	}
	defer lib.ProgramLibraryRelease(library)

	function, err := lib.ProgramLibraryRetainProgramFunction(library, "main")
	if err != nil {
		log.Fatal(err)
	}
	defer lib.ProgramFunctionRelease(function)

	opOptions, err := lib.PrecompiledComputeOpOptionsCreate(function)
	if err != nil {
		log.Fatal(err)
	}
	defer lib.PrecompiledComputeOpOptionsRelease(opOptions)
	if err := lib.PrecompiledComputeOpOptionsSetOperationName(opOptions, "main"); err != nil {
		log.Fatal(err)
	}
	op, err := lib.OperationCreatePrecompiled(opOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer lib.OperationRelease(op)

	// Bind a buffer to each named port.
	in, inPtr := bindExamplePort(lib, op, "x", channels*spatial*2, true)
	defer lib.BufferObjectRelease(in)
	out, outPtr := bindExamplePort(lib, op, "y", channels*spatial*2, false)
	defer lib.BufferObjectRelease(out)

	// Encode once, then dispatch.
	stream, err := lib.ExecutionStreamCreate()
	if err != nil {
		log.Fatal(err)
	}
	defer lib.ExecutionStreamRelease(stream)
	if err := lib.EncodeOperation(stream, op); err != nil {
		log.Fatal(err)
	}

	writeExampleFP16(inPtr, []float32{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	})
	if err := lib.ExecuteSync(stream); err != nil {
		log.Fatal(err)
	}
	fmt.Println(readExampleFP16(outPtr, channels*spatial))

	// Output:
	// [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16]
}

// Resolve a symbol the package does not list. Lib.Sym reports only what Open
// resolved; Lib.Lookup goes to dlsym.
func ExampleLib_Lookup() {
	lib, err := e5rt.Open()
	if err != nil {
		log.Fatal(err)
	}
	if _, err := lib.Lookup("e5rt_buffer_object_get_size"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("resolved")

	// Output:
	// resolved
}

func ExampleLib_ErrorString() {
	lib, err := e5rt.Open()
	if err != nil {
		log.Fatal(err)
	}
	text, err := lib.ErrorString(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(text)

	// Output:
	// OK
}
