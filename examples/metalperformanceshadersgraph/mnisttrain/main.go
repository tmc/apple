// Command mnisttrain trains a small MNIST digit classifier with MPSGraph.
//
// It is a Go port of Apple's "Training a Neural Network using MPS Graph"
// sample. The network is
//
//	conv 5x5x1x32 + relu -> max pool 2x2
//	conv 5x5x32x64 + relu -> max pool 2x2
//	fully connected 3136x1024 + relu
//	fully connected 1024x10
//
// trained with softmax cross entropy and plain stochastic gradient descent.
// The weights are graph variables, and each step assigns the descended value
// back into the variable, so training happens entirely inside the graph.
//
// The MNIST files are read from -data, and downloaded there if missing.
//
// Usage:
//
//	mnisttrain [-data dir] [-iterations n] [-batch n] [-lr rate]
package main

import (
	"flag"
	"fmt"
	"math/rand"
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

// Network dimensions. imageSize and numClasses come from the dataset.
const (
	conv0Out = 32
	conv1Out = 64
	// Two 2x2 pools take 28x28 down to 7x7.
	flatDim = 7 * 7 * conv1Out
	fcDim   = 1024
)

var (
	dataDir    = flag.String("data", defaultDataDir(), "directory holding the MNIST idx files")
	iterations = flag.Int("iterations", 300, "number of training iterations")
	batchSize  = flag.Int("batch", 40, "images per batch")
	rate       = flag.Float64("lr", 0.01, "learning rate")
	seed       = flag.Int64("seed", 1, "random seed for the weights and batch sampling")
)

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "mnisttrain: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	if *batchSize <= 0 || *iterations <= 0 {
		return fmt.Errorf("-batch and -iterations must be positive")
	}

	train, test, err := loadMNIST(*dataDir)
	if err != nil {
		return err
	}
	fmt.Printf("dataset: %d training images, %d test images\n", train.count, test.count)

	device := metal.MTLCreateSystemDefaultDevice()
	if device.GetID() == 0 {
		return fmt.Errorf("no Metal device available; this example needs a Metal-capable GPU")
	}
	queue := device.NewCommandQueue()
	if queue.GetID() == 0 {
		return fmt.Errorf("could not create a Metal command queue")
	}
	fmt.Printf("device: %s\n", device.Name())

	rng := rand.New(rand.NewSource(*seed))
	net, err := newClassifier(rng, *batchSize, *rate)
	if err != nil {
		return err
	}

	// The feeds are shared-storage buffers written in place before each run,
	// so a batch costs one memcpy and no blit.
	images, err := newFeed(device, net.images, *batchSize*imageSize*imageSize, *batchSize, imageSize*imageSize)
	if err != nil {
		return err
	}
	labels, err := newFeed(device, net.labels, *batchSize*numClasses, *batchSize, numClasses)
	if err != nil {
		return err
	}
	feeds := foundation.NewDictionaryWithObjectsForKeys(
		[]objectivec.IObject{images.data, labels.data},
		[]objectivec.IObject{images.tensor, labels.tensor},
	)
	if feeds.GetID() == 0 {
		return fmt.Errorf("could not create the feeds dictionary")
	}

	loss, err := newResult(device, net.loss, 1, 1)
	if err != nil {
		return err
	}
	probs, err := newResult(device, net.probabilities, *batchSize*numClasses, *batchSize, numClasses)
	if err != nil {
		return err
	}

	imageBatch := make([]float32, *batchSize*imageSize*imageSize)
	labelBatch := make([]float32, *batchSize*numClasses)

	for i := 1; i <= *iterations; i++ {
		train.randomBatch(rng, imageBatch, labelBatch)
		images.write(imageBatch)
		labels.write(labelBatch)
		net.graph.RunWithMTLCommandQueueFeedsTargetOperationsResultsDictionary(queue, feeds, net.updates, loss.dictionary)
		if i%25 == 0 || i == 1 || i == *iterations {
			fmt.Printf("iteration %3d/%d  loss %.4f\n", i, *iterations, loss.read(1)[0])
		}
	}

	// Inference runs the same graph without the assign operations, so the
	// weights stay where training left them.
	correct, seen := 0, 0
	for b := 0; b+*batchSize <= test.count; b += *batchSize {
		test.batchAt(b, imageBatch, labelBatch)
		images.write(imageBatch)
		labels.write(labelBatch)
		net.graph.RunWithMTLCommandQueueFeedsTargetOperationsResultsDictionary(queue, feeds, nil, probs.dictionary)
		correct += countCorrect(probs.read(*batchSize*numClasses), labelBatch)
		seen += *batchSize
	}
	if seen == 0 {
		return fmt.Errorf("test set has %d images, fewer than one batch of %d", test.count, *batchSize)
	}
	accuracy := 100 * float64(correct) / float64(seen)
	fmt.Printf("test accuracy: %d/%d = %.2f%%\n", correct, seen, accuracy)
	return nil
}

// countCorrect returns how many rows of probs have their largest value in the
// column that oneHot marks as the true class.
func countCorrect(probs, oneHot []float32) int {
	n := 0
	for row := 0; row+numClasses <= len(probs); row += numClasses {
		best, want := 0, 0
		for c := 1; c < numClasses; c++ {
			if probs[row+c] > probs[row+best] {
				best = c
			}
		}
		for c := 0; c < numClasses; c++ {
			if oneHot[row+c] == 1 {
				want = c
			}
		}
		if best == want {
			n++
		}
	}
	return n
}

// classifier holds the graph and the tensors a caller needs to drive it.
type classifier struct {
	graph         mpsg.MPSGraph
	images        mpsg.IMPSGraphTensor // placeholder, [batch, 28*28]
	labels        mpsg.IMPSGraphTensor // placeholder, one-hot [batch, 10]
	loss          mpsg.IMPSGraphTensor // mean cross entropy, [1]
	probabilities mpsg.IMPSGraphTensor // softmax of the logits, [batch, 10]
	updates       []mpsg.MPSGraphOperation
}

// newClassifier builds the graph and its SGD update operations. Weights are
// drawn uniformly from [-0.2, 0.2) and biases start at 0.1, matching the
// original sample.
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

	c.images = graph.PlaceholderWithShapeDataTypeName(shape(batch, imageSize*imageSize), dataTypeFloat32, "images")
	c.labels = graph.PlaceholderWithShapeDataTypeName(shape(batch, numClasses), dataTypeFloat32, "labels")

	var variables []mpsg.MPSGraphTensor
	t := graph.ReshapeTensorWithShapeName(c.images, shape(batch, imageSize, imageSize, 1), "nhwc")

	t = c.convLayer(rng, t, convDesc, 5, 5, 1, conv0Out, "conv0", &variables)
	t = graph.MaxPooling2DWithSourceTensorDescriptorName(t, poolDesc, "pool0")
	t = c.convLayer(rng, t, convDesc, 5, 5, conv0Out, conv1Out, "conv1", &variables)
	t = graph.MaxPooling2DWithSourceTensorDescriptorName(t, poolDesc, "pool1")

	t = graph.ReshapeTensorWithShapeName(t, shape(batch, flatDim), "flatten")
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
