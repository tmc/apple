// Command executable demonstrates the MPSGraphExecutable deployment path:
// compile a graph once, ship the compiled package, load it later without
// rebuilding the graph.
//
// An MPSGraph is a description. Running it with MPSGraph.run repeats the
// compile work the graph needs before the GPU sees anything. Compiling the
// graph yields an MPSGraphExecutable, which holds the specialized result and
// takes its inputs positionally instead of through a feeds dictionary. An
// executable can be serialized to an .mpsgraphpackage directory and read back
// on another process or machine, so the graph-building code does not have to
// ship with the program that runs it.
//
// The example builds a two-layer network (matmul, ReLU, matmul, bias), then
// runs it three ways and compares the results elementwise:
//
//	interpreted   MPSGraph.run with a feeds dictionary
//	compiled      MPSGraphExecutable.run from compile(with:feeds:...)
//	reloaded      MPSGraphExecutable.init(package:descriptor:) after serialize
//
// All three must agree exactly; the measured deltas are printed.
//
// Usage:
//
//	executable [-out dir]
//
// Without -out the package is written to a temporary directory and removed.
package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	mps "github.com/tmc/apple/metalperformanceshaders"
	mpsg "github.com/tmc/apple/metalperformanceshadersgraph"
	"github.com/tmc/apple/objectivec"
)

func init() {
	runtime.LockOSThread()
}

// The network shape. A batch of rows goes through hidden then out.
const (
	batch  = 4
	inDim  = 6
	hidDim = 8
	outDim = 3
)

func main() {
	out := flag.String("out", "", "directory to write the .mpsgraphpackage into; a temporary one is used and removed if empty")
	flag.Parse()
	if err := run(*out); err != nil {
		fmt.Fprintf(os.Stderr, "executable: %v\n", err)
		os.Exit(1)
	}
}

func run(outDir string) error {
	device := metal.MTLCreateSystemDefaultDevice()
	if device.GetID() == 0 {
		return fmt.Errorf("no Metal device available; this example needs a Metal-capable GPU")
	}
	queue := device.NewCommandQueue()
	if queue.GetID() == 0 {
		return fmt.Errorf("could not create a Metal command queue")
	}
	fmt.Printf("device: %s\n", device.Name())

	if outDir == "" {
		dir, err := os.MkdirTemp("", "mpsgraph")
		if err != nil {
			return fmt.Errorf("create temp dir: %w", err)
		}
		defer os.RemoveAll(dir)
		outDir = dir
	} else if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("create %s: %w", outDir, err)
	}
	pkg := filepath.Join(outDir, "twolayer.mpsgraphpackage")

	graph := mpsg.NewMPSGraph()
	if graph.GetID() == 0 {
		return fmt.Errorf("could not create MPSGraph")
	}
	inShape := shape(batch, inDim)
	input := graph.PlaceholderWithShapeDataTypeName(inShape, dataTypeFloat32, "x")
	output := network(graph, input)

	x := ramp(batch*inDim, -1, 0.13)
	inputBuf, err := bufferWith(device, x)
	if err != nil {
		return err
	}
	inputData := mpsg.NewGraphTensorDataWithMTLBufferShapeDataType(inputBuf, inShape, dataTypeFloat32)
	if inputData.GetID() == 0 {
		return fmt.Errorf("could not create tensor data for the input")
	}

	// Interpreted: run the graph itself, through a feeds dictionary.
	interpreted, err := runGraph(device, queue, graph, input, output, inputData)
	if err != nil {
		return err
	}

	// Compiled: specialize the graph for these input shapes once.
	gdev := mpsg.NewGraphDeviceWithMTLDevice(device)
	if gdev.GetID() == 0 {
		return fmt.Errorf("could not wrap the Metal device in an MPSGraphDevice")
	}
	compileDesc := mpsg.NewMPSGraphCompilationDescriptor()
	if compileDesc.GetID() == 0 {
		return fmt.Errorf("could not create the compilation descriptor")
	}
	compileDesc.SetOptimizationLevel(mpsg.MPSGraphOptimizationLevel1)
	compileDesc.SetWaitForCompilationCompletion(true)

	inType := mpsg.NewGraphShapedTypeWithShapeDataType(inShape, dataTypeFloat32)
	if inType.GetID() == 0 {
		return fmt.Errorf("could not create the input shaped type")
	}
	compileFeeds := foundation.NewDictionaryWithObjectsForKeys(
		[]objectivec.IObject{inType},
		[]objectivec.IObject{mpsg.MPSGraphTensorFromID(input.GetID())},
	)
	if compileFeeds.GetID() == 0 {
		return fmt.Errorf("could not create the compilation feeds dictionary")
	}
	exec := graph.CompileWithDeviceFeedsTargetTensorsTargetOperationsCompilationDescriptor(
		gdev, compileFeeds,
		[]mpsg.MPSGraphTensor{mpsg.MPSGraphTensorFromID(output.GetID())},
		nil, compileDesc,
	)
	if exec.GetID() == 0 {
		return fmt.Errorf("compiling the graph produced no executable")
	}
	describe("compiled", mpsg.MPSGraphExecutableFromID(exec.GetID()), gdev, inType, compileDesc)

	compiled, err := runExecutable(device, queue, mpsg.MPSGraphExecutableFromID(exec.GetID()), inputData)
	if err != nil {
		return fmt.Errorf("run compiled executable: %w", err)
	}

	// Serialize the executable, then read it back as if in another process.
	serDesc := mpsg.NewMPSGraphExecutableSerializationDescriptor()
	if serDesc.GetID() == 0 {
		return fmt.Errorf("could not create the serialization descriptor")
	}
	serDesc.SetAppend(false)
	url := foundation.NewURLFileURLWithPath(pkg)
	if url.GetID() == 0 {
		return fmt.Errorf("could not build a file URL for %s", pkg)
	}
	mpsg.MPSGraphExecutableFromID(exec.GetID()).SerializeToMPSGraphPackageAtURLDescriptor(url, serDesc)
	size, err := treeSize(pkg)
	if err != nil {
		return fmt.Errorf("serialize to %s: %w", pkg, err)
	}
	fmt.Printf("serialized %s (%d bytes on disk)\n", filepath.Base(pkg), size)

	reloaded := mpsg.GetMPSGraphExecutableClass().Alloc().InitWithMPSGraphPackageAtURLCompilationDescriptor(url, compileDesc)
	if reloaded.GetID() == 0 {
		return fmt.Errorf("could not load the executable back from %s", pkg)
	}
	describe("reloaded", reloaded, gdev, inType, compileDesc)

	fromDisk, err := runExecutable(device, queue, reloaded, inputData)
	if err != nil {
		return fmt.Errorf("run reloaded executable: %w", err)
	}

	fmt.Printf("\nfirst row of %d x %d output:\n", batch, outDim)
	for i := range outDim {
		fmt.Printf("  interpreted % .6f  compiled % .6f  reloaded % .6f\n", interpreted[i], compiled[i], fromDisk[i])
	}

	dCompiled := maxAbsDiff(interpreted, compiled)
	dReloaded := maxAbsDiff(interpreted, fromDisk)
	fmt.Printf("\nmax |interpreted-compiled| = %g\n", dCompiled)
	fmt.Printf("max |interpreted-reloaded| = %g\n", dReloaded)
	if dCompiled != 0 || dReloaded != 0 {
		return fmt.Errorf("results disagree: compiled %g, reloaded %g", dCompiled, dReloaded)
	}
	fmt.Println("PASS: compiled and reloaded executables reproduce the graph exactly")
	return nil
}

// network adds two dense layers with a ReLU between them and returns the
// output tensor. The weights are graph constants so the executable needs a
// single input.
func network(graph mpsg.MPSGraph, x mpsg.IMPSGraphTensor) mpsg.IMPSGraphTensor {
	w1 := constant(graph, ramp(inDim*hidDim, -0.5, 0.017), shape(inDim, hidDim))
	w2 := constant(graph, ramp(hidDim*outDim, -0.3, 0.031), shape(hidDim, outDim))
	b2 := constant(graph, ramp(outDim, 0.1, 0.05), shape(1, outDim))

	h := graph.MatrixMultiplicationWithPrimaryTensorSecondaryTensorName(x, w1, "x*w1")
	h = graph.ReLUWithTensorName(h, "relu")
	y := graph.MatrixMultiplicationWithPrimaryTensorSecondaryTensorName(h, w2, "h*w2")
	return graph.AdditionWithPrimaryTensorSecondaryTensorName(y, b2, "y+b2")
}

// runGraph evaluates the graph directly and returns the output values.
func runGraph(device metal.MTLDeviceObject, queue metal.MTLCommandQueue, graph mpsg.MPSGraph, input, output mpsg.IMPSGraphTensor, inputData mpsg.MPSGraphTensorData) ([]float32, error) {
	outBuf, outData, err := outputTensorData(device)
	if err != nil {
		return nil, err
	}
	feeds := foundation.NewDictionaryWithObjectsForKeys(
		[]objectivec.IObject{inputData},
		[]objectivec.IObject{mpsg.MPSGraphTensorFromID(input.GetID())},
	)
	results := foundation.NewDictionaryWithObjectsForKeys(
		[]objectivec.IObject{outData},
		[]objectivec.IObject{mpsg.MPSGraphTensorFromID(output.GetID())},
	)
	if feeds.GetID() == 0 || results.GetID() == 0 {
		return nil, fmt.Errorf("could not create the feeds and results dictionaries")
	}
	graph.RunWithMTLCommandQueueFeedsTargetOperationsResultsDictionary(queue, feeds, nil, results)
	return readFloats(outBuf, batch*outDim), nil
}

// runExecutable runs a compiled executable. Inputs and results are positional:
// they follow FeedTensors and TargetTensors, not a dictionary.
func runExecutable(device metal.MTLDeviceObject, queue metal.MTLCommandQueue, exec mpsg.MPSGraphExecutable, inputData mpsg.MPSGraphTensorData) ([]float32, error) {
	outBuf, outData, err := outputTensorData(device)
	if err != nil {
		return nil, err
	}
	execDesc := mpsg.NewMPSGraphExecutableExecutionDescriptor()
	if execDesc.GetID() == 0 {
		return nil, fmt.Errorf("could not create the execution descriptor")
	}
	execDesc.SetWaitUntilCompleted(true)
	got := exec.RunWithMTLCommandQueueInputsArrayResultsArrayExecutionDescriptor(
		queue,
		[]mpsg.MPSGraphTensorData{inputData},
		[]mpsg.MPSGraphTensorData{outData},
		execDesc,
	)
	if len(got) != 1 {
		return nil, fmt.Errorf("expected 1 result tensor, got %d", len(got))
	}
	return readFloats(outBuf, batch*outDim), nil
}

// describe prints what the executable reports about itself, including the
// output types it derives from the input types.
func describe(label string, exec mpsg.MPSGraphExecutable, gdev mpsg.MPSGraphDevice, inType mpsg.MPSGraphShapedType, desc mpsg.MPSGraphCompilationDescriptor) {
	inTypes := []mpsg.MPSGraphType{inType.MPSGraphType}
	exec.SpecializeWithDeviceInputTypesCompilationDescriptor(gdev, inTypes, desc)
	types := exec.GetOutputTypesWithDeviceInputTypesCompilationDescriptor(gdev, inTypes, desc)
	fmt.Printf("%s: %d feed tensor(s), %d target tensor(s), %d output type(s)",
		label, len(exec.FeedTensors()), len(exec.TargetTensors()), len(types))
	for _, t := range types {
		fmt.Printf(" %v", dims(t.Shape()))
	}
	fmt.Println()
}

// outputTensorData allocates a shared-storage buffer for the network output
// and the tensor data that wraps it.
func outputTensorData(device metal.MTLDeviceObject) (metal.MTLBuffer, mpsg.MPSGraphTensorData, error) {
	buf := device.NewBufferWithLengthOptions(uint(batch*outDim*4), metal.MTLResourceStorageModeShared)
	if buf.GetID() == 0 {
		return nil, mpsg.MPSGraphTensorData{}, fmt.Errorf("could not allocate the output buffer")
	}
	data := mpsg.NewGraphTensorDataWithMTLBufferShapeDataType(buf, shape(batch, outDim), dataTypeFloat32)
	if data.GetID() == 0 {
		return nil, mpsg.MPSGraphTensorData{}, fmt.Errorf("could not create tensor data for the result")
	}
	return buf, data, nil
}

// constant adds vals to graph as a constant tensor of the given shape.
func constant(graph mpsg.MPSGraph, vals []float32, shp foundation.NSArray) mpsg.IMPSGraphTensor {
	b := unsafe.Slice((*byte)(unsafe.Pointer(&vals[0])), len(vals)*4)
	data := foundation.NewDataFromBytes(b)
	runtime.KeepAlive(vals)
	return graph.ConstantWithDataShapeDataType(data, shp, dataTypeFloat32)
}

// ramp returns n values starting at start and stepping by step.
func ramp(n int, start, step float32) []float32 {
	v := make([]float32, n)
	for i := range v {
		v[i] = start + step*float32(i)
	}
	return v
}

// maxAbsDiff returns the largest elementwise difference between a and b.
func maxAbsDiff(a, b []float32) float64 {
	best := 0.0
	for i := range a {
		if d := math.Abs(float64(a[i] - b[i])); d > best {
			best = d
		}
	}
	return best
}

// dims reads an MPSGraph shape back into Go ints.
func dims(shp foundation.NSArray) []int {
	n := int(shp.Count())
	out := make([]int, n)
	for i := range n {
		out[i] = foundation.NSNumberFromID(shp.ObjectAtIndex(uint(i)).GetID()).IntegerValue()
	}
	return out
}

// dataTypeFloat32 is MPSDataTypeFloat32 in the form the graph bindings take.
var dataTypeFloat32 = uint32(mps.MPSDataTypeFloat32)

// shape returns dims as the NSArray of NSNumber that MPSGraph expects.
func shape(dims ...int) foundation.NSArray {
	nums := make([]foundation.NSNumber, len(dims))
	for i, d := range dims {
		nums[i] = foundation.NewNumberWithInteger(d)
	}
	return foundation.NSArrayFromID(objectivec.IObjectSliceToNSArray(nums))
}

// bufferWith returns a shared-storage buffer holding vals.
func bufferWith(device metal.MTLDeviceObject, vals []float32) (metal.MTLBuffer, error) {
	buf := device.NewBufferWithBytesLengthOptions(unsafe.Pointer(&vals[0]), uint(len(vals))*4, metal.MTLResourceStorageModeShared)
	runtime.KeepAlive(vals)
	if buf.GetID() == 0 {
		return nil, fmt.Errorf("could not allocate a %d-float buffer", len(vals))
	}
	return buf, nil
}

// readFloats copies n float32 values out of a shared-storage buffer.
func readFloats(buf metal.MTLBuffer, n int) []float32 {
	out := make([]float32, n)
	copy(out, unsafe.Slice((*float32)(buf.Contents()), n))
	return out
}

// treeSize returns the number of bytes in the file tree rooted at path. An
// .mpsgraphpackage is a directory, so serialization is confirmed by walking it.
func treeSize(path string) (int64, error) {
	var total int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			total += info.Size()
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	if total == 0 {
		return 0, fmt.Errorf("package is empty")
	}
	return total, nil
}
