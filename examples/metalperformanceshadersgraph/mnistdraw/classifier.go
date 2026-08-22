package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	mps "github.com/tmc/apple/metalperformanceshaders"
	mpsg "github.com/tmc/apple/metalperformanceshadersgraph"
	"github.com/tmc/apple/objectivec"
)

// Network dimensions. imageSize and numClasses come from the dataset.
const (
	conv0Out = 32
	conv1Out = 64
	// Two 2x2 pools take 28x28 down to 7x7.
	flatDim = 7 * 7 * conv1Out
	fcDim   = 1024
)

// classifier holds the graph and the tensors a caller needs to drive it.
type classifier struct {
	graph         mpsg.MPSGraph
	images        mpsg.IMPSGraphTensor // placeholder, [-1, 28*28]
	labels        mpsg.IMPSGraphTensor // placeholder, one-hot [-1, 10]
	loss          mpsg.IMPSGraphTensor // mean cross entropy, [1]
	probabilities mpsg.IMPSGraphTensor // softmax of the logits, [batch, 10]
	updates       []mpsg.MPSGraphOperation
}

// newClassifier builds the graph and its SGD update operations. Weights are
// drawn uniformly from [-0.2, 0.2) and biases start at 0.1, matching the
// original sample.
//
// The placeholders use -1 for the batch dimension, so one graph -- and one set
// of trained weights -- serves both the batch-sized training steps and the
// single-image inference the canvas needs. batch is still used for the loss
// mean, which only training reads.
func newClassifier(rng *rand.Rand, batch int, lr float64) (*classifier, error) {
	graph := mpsg.NewMPSGraph()
	if graph.GetID() == 0 {
		return nil, fmt.Errorf("could not create MPSGraph")
	}
	c := &classifier{graph: graph}

	convDesc := mpsg.NewGraphConvolution2DOpDescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYGroupsPaddingStyleDataLayoutWeightsLayout(
		1, 1, 1, 1, 1,
		mpsg.MPSGraphPaddingStyleTF_SAME,
		mpsg.MPSGraphTensorNamedDataLayoutNHWC,
		mpsg.MPSGraphTensorNamedDataLayoutHWIO,
	)
	if convDesc.GetID() == 0 {
		return nil, fmt.Errorf("could not create the convolution descriptor")
	}
	poolDesc := mpsg.NewGraphPooling2DOpDescriptorWithKernelWidthKernelHeightStrideInXStrideInYPaddingStyleDataLayout(
		2, 2, 2, 2,
		mpsg.MPSGraphPaddingStyleTF_SAME,
		mpsg.MPSGraphTensorNamedDataLayoutNHWC,
	)
	if poolDesc.GetID() == 0 {
		return nil, fmt.Errorf("could not create the pooling descriptor")
	}

	c.images = graph.PlaceholderWithShapeDataTypeName(shape(-1, imageSize*imageSize), dataTypeFloat32, "images")
	c.labels = graph.PlaceholderWithShapeDataTypeName(shape(-1, numClasses), dataTypeFloat32, "labels")

	var variables []mpsg.MPSGraphTensor
	t := graph.ReshapeTensorWithShapeName(c.images, shape(-1, imageSize, imageSize, 1), "nhwc")

	t = c.convLayer(rng, t, convDesc, 5, 5, 1, conv0Out, "conv0", &variables)
	t = graph.MaxPooling2DWithSourceTensorDescriptorName(t, poolDesc, "pool0")
	t = c.convLayer(rng, t, convDesc, 5, 5, conv0Out, conv1Out, "conv1", &variables)
	t = graph.MaxPooling2DWithSourceTensorDescriptorName(t, poolDesc, "pool1")

	t = graph.ReshapeTensorWithShapeName(t, shape(-1, flatDim), "flatten")
	t = c.fullyConnected(rng, t, flatDim, fcDim, true, "fc0", &variables)
	logits := c.fullyConnected(rng, t, fcDim, numClasses, false, "fc1", &variables)

	c.probabilities = graph.SoftMaxWithTensorAxisName(logits, -1, "softmax")

	sum := graph.SoftMaxCrossEntropyWithSourceTensorLabelsTensorAxisReductionTypeName(
		logits, c.labels, -1, mpsg.MPSGraphLossReductionTypeSum, "crossEntropy")
	batchTensor := graph.ConstantWithScalarShapeDataType(float64(batch), shape(1), dataTypeFloat32)
	c.loss = graph.DivisionWithPrimaryTensorSecondaryTensorName(sum, batchTensor, "meanLoss")

	updates, err := c.sgdUpdates(variables, lr)
	if err != nil {
		return nil, err
	}
	c.updates = updates
	return c, nil
}

// sgdUpdates returns one assign operation per variable, writing back the value
// that a single gradient-descent step produces.
func (c *classifier) sgdUpdates(variables []mpsg.MPSGraphTensor, lr float64) ([]mpsg.MPSGraphOperation, error) {
	gradients := c.graph.GradientForPrimaryTensorWithTensorsName(c.loss, variables, "gradients")
	if gradients.GetID() == 0 {
		return nil, fmt.Errorf("could not differentiate the loss")
	}
	rateTensor := c.graph.ConstantWithScalarShapeDataType(lr, shape(1), dataTypeFloat32)

	ops := make([]mpsg.MPSGraphOperation, 0, len(variables))
	for i, v := range variables {
		grad := gradients.ObjectForKey(v)
		if grad == nil || grad.GetID() == 0 {
			return nil, fmt.Errorf("no gradient for variable %d", i)
		}
		gradTensor := mpsg.MPSGraphTensorFromID(grad.GetID())
		updated := c.graph.StochasticGradientDescentWithLearningRateTensorValuesTensorGradientTensorName(
			rateTensor, v, gradTensor, "sgd")
		assign := c.graph.AssignVariableWithValueOfTensorName(v, updated, "assign")
		if assign.GetID() == 0 {
			return nil, fmt.Errorf("could not create the assign operation for variable %d", i)
		}
		ops = append(ops, mpsg.MPSGraphOperationFromID(assign.GetID()))
	}
	return ops, nil
}

// convLayer appends relu(conv(source) + bias) and records its variables.
func (c *classifier) convLayer(rng *rand.Rand, source mpsg.IMPSGraphTensor, desc mpsg.IMPSGraphConvolution2DOpDescriptor, kh, kw, in, out int, name string, variables *[]mpsg.MPSGraphTensor) mpsg.IMPSGraphTensor {
	weights := c.variable(uniform(rng, kh*kw*in*out), shape(kh, kw, in, out), name+".w", variables)
	biases := c.variable(filled(out, 0.1), shape(out), name+".b", variables)

	t := c.graph.Convolution2DWithSourceTensorWeightsTensorDescriptorName(source, weights, desc, name)
	t = c.graph.AdditionWithPrimaryTensorSecondaryTensorName(t, biases, name+".bias")
	return c.graph.ReLUWithTensorName(t, name+".relu")
}

// fullyConnected appends source*W + b, with an optional relu.
func (c *classifier) fullyConnected(rng *rand.Rand, source mpsg.IMPSGraphTensor, in, out int, activation bool, name string, variables *[]mpsg.MPSGraphTensor) mpsg.IMPSGraphTensor {
	weights := c.variable(uniform(rng, in*out), shape(in, out), name+".w", variables)
	biases := c.variable(filled(out, 0.1), shape(out), name+".b", variables)

	t := c.graph.MatrixMultiplicationWithPrimaryTensorSecondaryTensorName(source, weights, name)
	t = c.graph.AdditionWithPrimaryTensorSecondaryTensorName(t, biases, name+".bias")
	if !activation {
		return t
	}
	return c.graph.ReLUWithTensorName(t, name+".relu")
}

// variable creates a trainable tensor initialized to values and appends it to
// variables, which is the list the gradients are taken with respect to.
func (c *classifier) variable(values []float32, shp foundation.NSArray, name string, variables *[]mpsg.MPSGraphTensor) mpsg.IMPSGraphTensor {
	data := foundation.NewDataWithBytesLength(floatBytes(values))
	runtime.KeepAlive(values)
	v := c.graph.VariableWithDataShapeDataTypeName(data, shp, dataTypeFloat32, name)
	*variables = append(*variables, mpsg.MPSGraphTensorFromID(v.GetID()))
	return v
}

// feed is a placeholder together with the shared buffer that backs it.
type feed struct {
	tensor objectivec.IObject
	data   mpsg.MPSGraphTensorData
	buffer metal.MTLBuffer
	length int
}

func newFeed(device metal.MTLDeviceObject, tensor mpsg.IMPSGraphTensor, length int, dims ...int) (*feed, error) {
	buf, data, err := newTensorData(device, length, dims...)
	if err != nil {
		return nil, err
	}
	return &feed{tensor: mpsg.MPSGraphTensorFromID(tensor.GetID()), data: data, buffer: buf, length: length}, nil
}

// write copies values into the buffer the graph reads from.
func (f *feed) write(values []float32) {
	copy(unsafe.Slice((*float32)(f.buffer.Contents()), f.length), values)
}

// result is an output tensor plus the one-entry dictionary that names it as a
// destination for a graph run.
type result struct {
	buffer     metal.MTLBuffer
	dictionary foundation.NSDictionary
}

func newResult(device metal.MTLDeviceObject, tensor mpsg.IMPSGraphTensor, length int, dims ...int) (*result, error) {
	buf, data, err := newTensorData(device, length, dims...)
	if err != nil {
		return nil, err
	}
	dict := foundation.NewDictionaryWithObjectsForKeys(
		[]objectivec.IObject{data},
		[]objectivec.IObject{mpsg.MPSGraphTensorFromID(tensor.GetID())},
	)
	if dict.GetID() == 0 {
		return nil, fmt.Errorf("could not create a results dictionary")
	}
	return &result{buffer: buf, dictionary: dict}, nil
}

// read copies n float32 values out of the result buffer.
func (r *result) read(n int) []float32 {
	out := make([]float32, n)
	copy(out, unsafe.Slice((*float32)(r.buffer.Contents()), n))
	return out
}

// newTensorData allocates a shared-storage buffer of length float32 values and
// wraps it as tensor data of the given shape.
func newTensorData(device metal.MTLDeviceObject, length int, dims ...int) (metal.MTLBuffer, mpsg.MPSGraphTensorData, error) {
	buf := device.NewBufferWithLengthOptions(uint(length*4), metal.MTLResourceStorageModeShared)
	if buf.GetID() == 0 {
		return buf, mpsg.MPSGraphTensorData{}, fmt.Errorf("could not allocate a %d-float buffer", length)
	}
	data := mpsg.NewGraphTensorDataWithMTLBufferShapeDataType(buf, shape(dims...), dataTypeFloat32)
	if data.GetID() == 0 {
		return buf, data, fmt.Errorf("could not create tensor data for a %d-float buffer", length)
	}
	return buf, data, nil
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

// uniform returns n values drawn uniformly from [-0.2, 0.2).
func uniform(rng *rand.Rand, n int) []float32 {
	v := make([]float32, n)
	for i := range v {
		v[i] = float32(rng.Float64()*0.4 - 0.2)
	}
	return v
}

// filled returns n copies of x.
func filled(n int, x float32) []float32 {
	v := make([]float32, n)
	for i := range v {
		v[i] = x
	}
	return v
}

// floatBytes returns the raw bytes of values, which stay valid only as long as
// values does.
func floatBytes(values []float32) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&values[0])), len(values)*4)
}
