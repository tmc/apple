package main

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	mps "github.com/tmc/apple/metalperformanceshaders"
	mpsg "github.com/tmc/apple/metalperformanceshadersgraph"
	"github.com/tmc/apple/objectivec"
)

// newGPUArm builds an MPSGraph computing out = W2*relu(W1*x)^2, the same
// function the Neural Engine arm runs.
//
// The layouts differ and the transposition happens here rather than on either
// engine. The MIL program works in [1, channels, 1, positions], so an
// activation is addressed as x[c*seq+s], while MPSGraph is given the natural
// matrix form [positions, channels] and weights already transposed to
// [in, out]. Converting on the host at the boundary keeps both arms computing
// the same arithmetic instead of one of them also doing a transpose.
func newGPUArm(dim, hidden, seq int, w1, w2 []float32) (*arm, error) {
	device := metal.MTLCreateSystemDefaultDevice()
	if device.GetID() == 0 {
		return nil, fmt.Errorf("no Metal device available")
	}
	queue := device.NewCommandQueue()
	if queue.GetID() == 0 {
		return nil, fmt.Errorf("could not create a Metal command queue")
	}
	graph := mpsg.NewMPSGraph()
	if graph.GetID() == 0 {
		return nil, fmt.Errorf("could not create MPSGraph")
	}

	dt := uint32(mps.MPSDataTypeFloat32)
	xT := graph.PlaceholderWithShapeDataTypeName(shape(seq, dim), dt, "x")
	w1T := graph.PlaceholderWithShapeDataTypeName(shape(dim, hidden), dt, "w1")
	w2T := graph.PlaceholderWithShapeDataTypeName(shape(hidden, dim), dt, "w2")

	h := graph.MatrixMultiplicationWithPrimaryTensorSecondaryTensorName(xT, w1T, "xw1")
	h = graph.ReLUWithTensorName(h, "relu")
	h = graph.SquareWithTensorName(h, "square")
	y := graph.MatrixMultiplicationWithPrimaryTensorSecondaryTensorName(h, w2T, "hw2")

	// Weights, transposed from the OIHW [out][in] the MIL program uses into the
	// [in][out] MPSGraph multiplies by.
	w1t := transpose(w1, hidden, dim)
	w2t := transpose(w2, dim, hidden)

	xBuf := device.NewBufferWithLengthOptions(uint(seq*dim*4), metal.MTLResourceStorageModeShared)
	if xBuf.GetID() == 0 {
		return nil, fmt.Errorf("could not allocate the input buffer")
	}
	w1Buf, err := bufferWith(device, w1t)
	if err != nil {
		return nil, err
	}
	w2Buf, err := bufferWith(device, w2t)
	if err != nil {
		return nil, err
	}
	outBuf := device.NewBufferWithLengthOptions(uint(seq*dim*4), metal.MTLResourceStorageModeShared)
	if outBuf.GetID() == 0 {
		return nil, fmt.Errorf("could not allocate the output buffer")
	}

	tensorData := func(buf metal.MTLBuffer, dims ...int) (mpsg.MPSGraphTensorData, error) {
		data := mpsg.NewGraphTensorDataWithMTLBufferShapeDataType(buf, shape(dims...), dt)
		if data.GetID() == 0 {
			return data, fmt.Errorf("could not create tensor data")
		}
		return data, nil
	}
	xData, err := tensorData(xBuf, seq, dim)
	if err != nil {
		return nil, err
	}
	w1Data, err := tensorData(w1Buf, dim, hidden)
	if err != nil {
		return nil, err
	}
	w2Data, err := tensorData(w2Buf, hidden, dim)
	if err != nil {
		return nil, err
	}
	outData, err := tensorData(outBuf, seq, dim)
	if err != nil {
		return nil, err
	}

	feeds := foundation.NewDictionaryWithObjectsForKeys(
		[]objectivec.IObject{xData, w1Data, w2Data},
		[]objectivec.IObject{
			mpsg.MPSGraphTensorFromID(xT.GetID()),
			mpsg.MPSGraphTensorFromID(w1T.GetID()),
			mpsg.MPSGraphTensorFromID(w2T.GetID()),
		},
	)
	if feeds.GetID() == 0 {
		return nil, fmt.Errorf("could not create the feeds dictionary")
	}
	results := foundation.NewDictionaryWithObjectsForKeys(
		[]objectivec.IObject{outData},
		[]objectivec.IObject{mpsg.MPSGraphTensorFromID(y.GetID())},
	)
	if results.GetID() == 0 {
		return nil, fmt.Errorf("could not create the results dictionary")
	}

	return &arm{
		name: "gpu",
		once: func() error {
			graph.RunWithMTLCommandQueueFeedsTargetOperationsResultsDictionary(queue, feeds, nil, results)
			return nil
		},
		read: func() []float32 {
			// [seq][dim] back to the [dim][seq] the reference uses.
			src := unsafe.Slice((*float32)(outBuf.Contents()), seq*dim)
			out := make([]float32, dim*seq)
			for s := range seq {
				for c := range dim {
					out[c*seq+s] = src[s*dim+c]
				}
			}
			return out
		},
		write: func(data []float32) {
			dst := unsafe.Slice((*float32)(xBuf.Contents()), seq*dim)
			for s := range seq {
				for c := range dim {
					dst[s*dim+c] = data[c*seq+s]
				}
			}
		},
	}, nil
}

// transpose returns an [in][out] copy of a row-major [out][in] matrix.
func transpose(m []float32, rows, cols int) []float32 {
	out := make([]float32, len(m))
	for r := range rows {
		for c := range cols {
			out[c*rows+r] = m[r*cols+c]
		}
	}
	return out
}

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
