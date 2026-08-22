// Command mpsgraph runs a small MLP forward pass on the GPU with MPSGraph.
//
// It builds the graph y = relu(x*W1 + b1)*W2 + b2 for a batch of two input
// rows, executes it on the default Metal device, and prints the result tensor
// next to the same computation done on the CPU so the numbers can be checked.
//
// Usage:
//
//	mpsgraph
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

// Problem dimensions: batch x in -> hidden -> out.
const (
	batch  = 2
	inDim  = 3
	hidDim = 4
	outDim = 2
)

// Inputs and weights, small enough to verify by hand.
var (
	x = []float32{
		1, 2, 3,
		-1, 0, 1,
	}
	w1 = []float32{
		0.5, -0.5, 1, 0,
		0, 1, -1, 0.5,
		0.25, 0.25, 0.25, 0.25,
	}
	b1 = []float32{0.1, -0.1, 0.2, -0.2}
	w2 = []float32{
		1, 0,
		0, 1,
		-1, 1,
		0.5, -0.5,
	}
	b2 = []float32{0.01, -0.01}
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "mpsgraph: %v\n", err)
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

	xT := graph.PlaceholderWithShapeDataTypeName(shape(batch, inDim), dataTypeFloat32, "x")
	w1T := graph.PlaceholderWithShapeDataTypeName(shape(inDim, hidDim), dataTypeFloat32, "w1")
	b1T := graph.PlaceholderWithShapeDataTypeName(shape(1, hidDim), dataTypeFloat32, "b1")
	w2T := graph.PlaceholderWithShapeDataTypeName(shape(hidDim, outDim), dataTypeFloat32, "w2")
	b2T := graph.PlaceholderWithShapeDataTypeName(shape(1, outDim), dataTypeFloat32, "b2")

	h := graph.MatrixMultiplicationWithPrimaryTensorSecondaryTensorName(xT, w1T, "xw1")
	h = graph.AdditionWithPrimaryTensorSecondaryTensorName(h, b1T, "xw1+b1")
	h = graph.ReLUWithTensorName(h, "relu")
	y := graph.MatrixMultiplicationWithPrimaryTensorSecondaryTensorName(h, w2T, "hw2")
	y = graph.AdditionWithPrimaryTensorSecondaryTensorName(y, b2T, "hw2+b2")

	// Every feed is a shared-storage MTLBuffer so the CPU can write the inputs
	// and read the output back without an explicit blit.
	feedTensors := []mpsg.IMPSGraphTensor{xT, w1T, b1T, w2T, b2T}
	feedValues := [][]float32{x, w1, b1, w2, b2}
	feedShapes := []foundation.NSArray{
		shape(batch, inDim),
		shape(inDim, hidDim),
		shape(1, hidDim),
		shape(hidDim, outDim),
		shape(1, outDim),
	}
	keys := make([]objectivec.IObject, len(feedTensors))
	objects := make([]objectivec.IObject, len(feedTensors))
	for i, t := range feedTensors {
		buf, err := bufferWith(device, feedValues[i])
		if err != nil {
			return err
		}
		data := mpsg.NewGraphTensorDataWithMTLBufferShapeDataType(buf, feedShapes[i], dataTypeFloat32)
		if data.GetID() == 0 {
			return fmt.Errorf("could not create tensor data for feed %d", i)
		}
		keys[i] = mpsg.MPSGraphTensorFromID(t.GetID())
		objects[i] = data
	}
	feeds := foundation.NewDictionaryWithObjectsForKeys(objects, keys)
	if feeds.GetID() == 0 {
		return fmt.Errorf("could not create feeds dictionary")
	}

	outBuf := device.NewBufferWithLengthOptions(batch*outDim*4, metal.MTLResourceStorageModeShared)
	if outBuf.GetID() == 0 {
		return fmt.Errorf("could not allocate output buffer")
	}
	outData := mpsg.NewGraphTensorDataWithMTLBufferShapeDataType(outBuf, shape(batch, outDim), dataTypeFloat32)
	if outData.GetID() == 0 {
		return fmt.Errorf("could not create tensor data for the result")
	}
	results := foundation.NewDictionaryWithObjectsForKeys(
		[]objectivec.IObject{outData},
		[]objectivec.IObject{mpsg.MPSGraphTensorFromID(y.GetID())},
	)
	if results.GetID() == 0 {
		return fmt.Errorf("could not create results dictionary")
	}

	graph.RunWithMTLCommandQueueFeedsTargetOperationsResultsDictionary(queue, feeds, nil, results)

	got := readFloats(outBuf, batch*outDim)
	want := cpuForward()

	fmt.Printf("y = relu(x*W1 + b1)*W2 + b2, shape [%d %d]\n", batch, outDim)
	maxDiff := 0.0
	for r := 0; r < batch; r++ {
		for c := 0; c < outDim; c++ {
			g, w := got[r*outDim+c], want[r*outDim+c]
			fmt.Printf("  y[%d][%d] = % .6f (cpu % .6f)\n", r, c, g, w)
			if d := math.Abs(float64(g - w)); d > maxDiff {
				maxDiff = d
			}
		}
	}
	fmt.Printf("max |gpu-cpu| = %g\n", maxDiff)
	if maxDiff > 1e-5 {
		return fmt.Errorf("GPU result disagrees with the CPU reference by %g", maxDiff)
	}
	return nil
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

// cpuForward computes the same MLP in Go, as a reference.
func cpuForward() []float32 {
	h := make([]float32, batch*hidDim)
	for r := 0; r < batch; r++ {
		for c := 0; c < hidDim; c++ {
			sum := b1[c]
			for k := 0; k < inDim; k++ {
				sum += x[r*inDim+k] * w1[k*hidDim+c]
			}
			if sum < 0 {
				sum = 0
			}
			h[r*hidDim+c] = sum
		}
	}
	y := make([]float32, batch*outDim)
	for r := 0; r < batch; r++ {
		for c := 0; c < outDim; c++ {
			sum := b2[c]
			for k := 0; k < hidDim; k++ {
				sum += h[r*hidDim+k] * w2[k*outDim+c]
			}
			y[r*outDim+c] = sum
		}
	}
	return y
}
