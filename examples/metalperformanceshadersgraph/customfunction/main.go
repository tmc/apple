// Command customfunction adds a custom function to an MPSGraph.
//
// It is a Go port of Apple's "Adding Custom Functions to a Shader Graph"
// sample. MPSGraph has no GeLU primitive, so the sample writes one out of the
// primitives the graph does have:
//
//	gelu(x) = 0.5*x * (1 + erf(x * sqrt(0.5)))
//
// The Swift sample subclasses MPSGraph to hang the method off the graph. In Go
// the same composition is a plain function that takes the graph, which is why
// geLU below builds nodes rather than computing numbers: a custom function is
// a subgraph, not a kernel.
//
// The result is printed next to the same function evaluated on the CPU so the
// numbers can be checked.
//
// Usage:
//
//	customfunction
package main

import (
	"fmt"
	"math"
	"os"
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

// x holds the inputs to evaluate. The Swift sample feeds the single scalar 2,
// which is the third value here.
var x = []float32{-3, -1, 2, 0, 0.5, 1, 3}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "customfunction: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	device := metal.MTLCreateSystemDefaultDevice()
	if device.GetID() == 0 {
		return fmt.Errorf("no Metal device available; this example needs a Metal-capable GPU")
	}
	queue := device.NewCommandQueue()
	if queue.GetID() == 0 {
		return fmt.Errorf("could not create a Metal command queue")
	}
	fmt.Printf("device: %s\n", device.Name())

	graph := mpsg.NewMPSGraph()
	if graph.GetID() == 0 {
		return fmt.Errorf("could not create MPSGraph")
	}

	// The custom function is shape agnostic, so the placeholder shape only has
	// to match the buffers fed to it.
	shp := shape(len(x))
	input := graph.PlaceholderWithShapeDataTypeName(shp, dataTypeFloat32, "x")
	output := geLU(graph, input)

	inputBuf, err := bufferWith(device, x)
	if err != nil {
		return err
	}
	inputData := mpsg.NewGraphTensorDataWithMTLBufferShapeDataType(inputBuf, shp, dataTypeFloat32)
	if inputData.GetID() == 0 {
		return fmt.Errorf("could not create tensor data for the input")
	}
	feeds := foundation.NewDictionaryWithObjectsForKeys(
		[]objectivec.IObject{inputData},
		[]objectivec.IObject{mpsg.MPSGraphTensorFromID(input.GetID())},
	)
	if feeds.GetID() == 0 {
		return fmt.Errorf("could not create the feeds dictionary")
	}

	outputBuf := device.NewBufferWithLengthOptions(uint(len(x)*4), metal.MTLResourceStorageModeShared)
	if outputBuf.GetID() == 0 {
		return fmt.Errorf("could not allocate the output buffer")
	}
	outputData := mpsg.NewGraphTensorDataWithMTLBufferShapeDataType(outputBuf, shp, dataTypeFloat32)
	if outputData.GetID() == 0 {
		return fmt.Errorf("could not create tensor data for the result")
	}
	results := foundation.NewDictionaryWithObjectsForKeys(
		[]objectivec.IObject{outputData},
		[]objectivec.IObject{mpsg.MPSGraphTensorFromID(output.GetID())},
	)
	if results.GetID() == 0 {
		return fmt.Errorf("could not create the results dictionary")
	}

	graph.RunWithMTLCommandQueueFeedsTargetOperationsResultsDictionary(queue, feeds, nil, results)

	got := readFloats(outputBuf, len(x))
	fmt.Printf("gelu(x) = 0.5*x * (1 + erf(x * sqrt(0.5)))\n")
	maxDiff := 0.0
	for i, in := range x {
		want := cpuGeLU(in)
		fmt.Printf("  gelu(% .1f) = % .6f (cpu % .6f)\n", in, got[i], want)
		if d := math.Abs(float64(got[i] - want)); d > maxDiff {
			maxDiff = d
		}
	}
	fmt.Printf("max |gpu-cpu| = %g\n", maxDiff)
	if maxDiff > 1e-5 {
		return fmt.Errorf("GPU result disagrees with the CPU reference by %g", maxDiff)
	}
	return nil
}

// geLU adds the GeLU function to graph and returns its output tensor. It is
// built only from graph primitives, so it differentiates and fuses like any
// built-in operation.
func geLU(graph mpsg.MPSGraph, tensor mpsg.IMPSGraphTensor) mpsg.IMPSGraphTensor {
	ones := graph.ConstantWithScalarShapeDataType(1, shape(1), dataTypeFloat32)
	half := graph.ConstantWithScalarShapeDataType(0.5, shape(1), dataTypeFloat32)

	sqrtHalf := graph.SquareRootWithTensorName(half, "sqrt(0.5)")
	scaled := graph.MultiplicationWithPrimaryTensorSecondaryTensorName(sqrtHalf, tensor, "x*sqrt(0.5)")
	halved := graph.MultiplicationWithPrimaryTensorSecondaryTensorName(half, tensor, "0.5*x")

	erf := graph.ErfWithTensorName(scaled, "erf")
	sum := graph.AdditionWithPrimaryTensorSecondaryTensorName(erf, ones, "1+erf")

	return graph.MultiplicationWithPrimaryTensorSecondaryTensorName(halved, sum, "gelu")
}

// cpuGeLU computes the same function in Go, as a reference.
func cpuGeLU(v float32) float32 {
	x := float64(v)
	return float32(0.5 * x * (1 + math.Erf(x*math.Sqrt(0.5))))
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
