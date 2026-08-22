// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNGraph] class.
var (
	_MPSNNGraphClass     MPSNNGraphClass
	_MPSNNGraphClassOnce sync.Once
)

func getMPSNNGraphClass() MPSNNGraphClass {
	_MPSNNGraphClassOnce.Do(func() {
		_MPSNNGraphClass = MPSNNGraphClass{class: objc.GetClass("MPSNNGraph")}
	})
	return _MPSNNGraphClass
}

// GetMPSNNGraphClass returns the class object for MPSNNGraph.
func GetMPSNNGraphClass() MPSNNGraphClass {
	return getMPSNNGraphClass()
}

type MPSNNGraphClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNGraphClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNGraphClass) Alloc() MPSNNGraph {
	rv := objc.Send[MPSNNGraph](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An optimized representation of a graph of neural network image and filter
// nodes.
//
// # Overview
//
// Once you have prepared a graph of [MPSNNImageNode], [MPSNNFilterNode], and,
// if needed, [MPSNNStateNode] objects, you may initialize a [MPSNNGraph]
// using the image node that you wish to appear as the result. The graph
// object will introspect the graph representation and determine which nodes
// are needed for inputs, and which nodes are produced as output state (if
// any). Nodes which are not needed to calculate the result image node are
// ignored. Some nodes may be internally concatenated with other nodes for
// better performance.
//
// During [MPSNNGraph] construction, the graph attached to the result node
// will be parsed and reduced to an optimized representation. This
// representation may be saved using the [NSSecureCoding] protocol for later
// recall.
//
// When decoding a [MPSNNGraph] using a [NSCoder], it will be created against
// the system default [MTLDevice]. If you would like to set the device, your
// [NSCoder] should conform to the [MPSDeviceProvider] protocol.
//
// # Debugging Tips
//
// In typical usage, some refinement, especially of padding policies, may be
// required to get the expected answer from Metal Performance Shaders. If the
// result image is the wrong size, padding is typically the problem. When the
// answers are incorrect, the [MPSCNNKernel.Offset] or other property may be
// incorrectly configured at some stage. As the graph is generated starting
// from an output image node, you may create other graphs starting at any
// image node within the graph. This will give you a view into the result
// produced from each intermediate layer with a minimum of fuss. In addition,
// the usual [debugDescription()] method is available to inspect objects to
// make sure they conform to expectation.
//
// Note that certain operations such as neuron filters that follow convolution
// filters and image concatenation may be optimized away by the [MPSNNGraph]
// when it is constructed. The convolution can do neuron operations as part of
// its operation. Concatenation is best done by writing the result of earlier
// filter passes in the right place using
// [MPSCNNKernel.DestinationFeatureChannelOffset] rather than by adding an
// extra copy. Other optimizations may be added as framework capabilities
// improve.
//
// # Initializers
//
//   - [MPSNNGraph.InitWithDeviceResultImageResultImageIsNeeded]
//   - [MPSNNGraph.InitWithDeviceResultImagesResultsAreNeeded]
//
// # Instance Properties
//
//   - [MPSNNGraph.DestinationImageAllocator]
//   - [MPSNNGraph.SetDestinationImageAllocator]
//   - [MPSNNGraph.IntermediateImageHandles]
//   - [MPSNNGraph.OutputStateIsTemporary]
//   - [MPSNNGraph.SetOutputStateIsTemporary]
//   - [MPSNNGraph.ResultHandle]
//   - [MPSNNGraph.ResultStateHandles]
//   - [MPSNNGraph.SourceImageHandles]
//   - [MPSNNGraph.SourceStateHandles]
//   - [MPSNNGraph.Format]
//   - [MPSNNGraph.SetFormat]
//   - [MPSNNGraph.ResultImageIsNeeded]
//
// # Instance Methods
//
//   - [MPSNNGraph.EncodeToCommandBufferSourceImages]
//   - [MPSNNGraph.EncodeToCommandBufferSourceImagesSourceStatesIntermediateImagesDestinationStates]
//   - [MPSNNGraph.ExecuteAsyncWithSourceImagesCompletionHandler]
//   - [MPSNNGraph.EncodeBatchToCommandBufferSourceImagesSourceStates]
//   - [MPSNNGraph.EncodeBatchToCommandBufferSourceImagesSourceStatesIntermediateImagesDestinationStates]
//   - [MPSNNGraph.ReadCountForSourceImageAtIndex]
//   - [MPSNNGraph.ReadCountForSourceStateAtIndex]
//   - [MPSNNGraph.ReloadFromDataSources]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph
//
// [MTLDevice]: https://developer.apple.com/documentation/Metal/MTLDevice
// [NSCoder]: https://developer.apple.com/documentation/Foundation/NSCoder
// [NSSecureCoding]: https://developer.apple.com/documentation/Foundation/NSSecureCoding
// [debugDescription()]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/debugDescription()
type MPSNNGraph struct {
	MPSKernel
}

// MPSNNGraphFromID constructs a [MPSNNGraph] from an objc.ID.
//
// An optimized representation of a graph of neural network image and filter
// nodes.
func MPSNNGraphFromID(id objc.ID) MPSNNGraph {
	return MPSNNGraph{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSNNGraph adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNGraph] class.
//
// # Initializers
//
//   - [IMPSNNGraph.InitWithDeviceResultImageResultImageIsNeeded]
//   - [IMPSNNGraph.InitWithDeviceResultImagesResultsAreNeeded]
//
// # Instance Properties
//
//   - [IMPSNNGraph.DestinationImageAllocator]
//   - [IMPSNNGraph.SetDestinationImageAllocator]
//   - [IMPSNNGraph.IntermediateImageHandles]
//   - [IMPSNNGraph.OutputStateIsTemporary]
//   - [IMPSNNGraph.SetOutputStateIsTemporary]
//   - [IMPSNNGraph.ResultHandle]
//   - [IMPSNNGraph.ResultStateHandles]
//   - [IMPSNNGraph.SourceImageHandles]
//   - [IMPSNNGraph.SourceStateHandles]
//   - [IMPSNNGraph.Format]
//   - [IMPSNNGraph.SetFormat]
//   - [IMPSNNGraph.ResultImageIsNeeded]
//
// # Instance Methods
//
//   - [IMPSNNGraph.EncodeToCommandBufferSourceImages]
//   - [IMPSNNGraph.EncodeToCommandBufferSourceImagesSourceStatesIntermediateImagesDestinationStates]
//   - [IMPSNNGraph.ExecuteAsyncWithSourceImagesCompletionHandler]
//   - [IMPSNNGraph.EncodeBatchToCommandBufferSourceImagesSourceStates]
//   - [IMPSNNGraph.EncodeBatchToCommandBufferSourceImagesSourceStatesIntermediateImagesDestinationStates]
//   - [IMPSNNGraph.ReadCountForSourceImageAtIndex]
//   - [IMPSNNGraph.ReadCountForSourceStateAtIndex]
//   - [IMPSNNGraph.ReloadFromDataSources]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph
type IMPSNNGraph interface {
	IMPSKernel

	// Topic: Initializers

	InitWithDeviceResultImageResultImageIsNeeded(device metal.MTLDevice, resultImage IMPSNNImageNode, resultIsNeeded bool) MPSNNGraph
	InitWithDeviceResultImagesResultsAreNeeded(device metal.MTLDevice, resultImages []MPSNNImageNode, areResultsNeeded *bool) MPSNNGraph

	// Topic: Instance Properties

	DestinationImageAllocator() MPSImageAllocator
	SetDestinationImageAllocator(value MPSImageAllocator)
	IntermediateImageHandles() []objectivec.IObject
	OutputStateIsTemporary() bool
	SetOutputStateIsTemporary(value bool)
	ResultHandle() MPSHandle
	ResultStateHandles() []objectivec.IObject
	SourceImageHandles() []objectivec.IObject
	SourceStateHandles() []objectivec.IObject
	Format() MPSImageFeatureChannelFormat
	SetFormat(value MPSImageFeatureChannelFormat)
	ResultImageIsNeeded() bool

	// Topic: Instance Methods

	EncodeToCommandBufferSourceImages(commandBuffer metal.MTLCommandBuffer, sourceImages []MPSImage) IMPSImage
	EncodeToCommandBufferSourceImagesSourceStatesIntermediateImagesDestinationStates(commandBuffer metal.MTLCommandBuffer, sourceImages []MPSImage, sourceStates []MPSState, intermediateImages foundation.INSArray, destinationStates foundation.INSArray) IMPSImage
	ExecuteAsyncWithSourceImagesCompletionHandler(sourceImages []MPSImage, handler MPSImageErrorHandler) IMPSImage
	EncodeBatchToCommandBufferSourceImagesSourceStates(commandBuffer metal.MTLCommandBuffer, sourceImages []foundation.NSArray, sourceStates []foundation.NSArray) MPSImageBatch
	EncodeBatchToCommandBufferSourceImagesSourceStatesIntermediateImagesDestinationStates(commandBuffer metal.MTLCommandBuffer, sourceImages []foundation.NSArray, sourceStates []foundation.NSArray, intermediateImages foundation.INSArray, destinationStates foundation.INSArray) MPSImageBatch
	ReadCountForSourceImageAtIndex(index uint) uint
	ReadCountForSourceStateAtIndex(index uint) uint
	ReloadFromDataSources()
}

// Init initializes the instance.
func (g MPSNNGraph) Init() MPSNNGraph {
	rv := objc.Send[MPSNNGraph](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSNNGraph) Autorelease() MPSNNGraph {
	rv := objc.Send[MPSNNGraph](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNGraph creates a new MPSNNGraph instance.
func NewMPSNNGraph() MPSNNGraph {
	class := getMPSNNGraphClass()
	rv := objc.Send[MPSNNGraph](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewGraphWithCoder(aDecoder foundation.INSCoder) MPSNNGraph {
	instance := getMPSNNGraphClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNNGraphFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/init(coder:device:)
func NewGraphWithCoderDevice(aDecoder foundation.INSCoder, device metal.MTLDevice) MPSNNGraph {
	instance := getMPSNNGraphClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	return MPSNNGraphFromID(rv)
}

// Initializes a new kernel object.
//
// device: The Metal device on which the kernel will be used.
//
// # Return Value
//
// An initialized kernel object.
//
// # Discussion
//
// This method fails if the device is not supported. Query the
// [MPSSupportsMTLDevice] function to determine whether the device is
// supported.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(device:)
func NewGraphWithDevice(device metal.MTLDevice) MPSNNGraph {
	instance := getMPSNNGraphClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNNGraphFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/init(device:resultImage:resultImageIsNeeded:)
func NewGraphWithDeviceResultImageResultImageIsNeeded(device metal.MTLDevice, resultImage IMPSNNImageNode, resultIsNeeded bool) MPSNNGraph {
	instance := getMPSNNGraphClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resultImage:resultImageIsNeeded:"), device, resultImage, resultIsNeeded)
	return MPSNNGraphFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/init(device:resultImages:resultsAreNeeded:)
func NewGraphWithDeviceResultImagesResultsAreNeeded(device metal.MTLDevice, resultImages []MPSNNImageNode, areResultsNeeded *bool) MPSNNGraph {
	instance := getMPSNNGraphClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resultImages:resultsAreNeeded:"), device, objectivec.IObjectSliceToNSArray(resultImages), areResultsNeeded)
	return MPSNNGraphFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/init(device:resultImage:resultImageIsNeeded:)
func (g MPSNNGraph) InitWithDeviceResultImageResultImageIsNeeded(device metal.MTLDevice, resultImage IMPSNNImageNode, resultIsNeeded bool) MPSNNGraph {
	rv := objc.Send[MPSNNGraph](g.ID, objc.Sel("initWithDevice:resultImage:resultImageIsNeeded:"), device, resultImage, resultIsNeeded)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/init(device:resultImages:resultsAreNeeded:)
func (g MPSNNGraph) InitWithDeviceResultImagesResultsAreNeeded(device metal.MTLDevice, resultImages []MPSNNImageNode, areResultsNeeded *bool) MPSNNGraph {
	rv := objc.Send[MPSNNGraph](g.ID, objc.Sel("initWithDevice:resultImages:resultsAreNeeded:"), device, objectivec.IObjectSliceToNSArray(resultImages), areResultsNeeded)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/encode(to:sourceImages:)
func (g MPSNNGraph) EncodeToCommandBufferSourceImages(commandBuffer metal.MTLCommandBuffer, sourceImages []MPSImage) IMPSImage {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("encodeToCommandBuffer:sourceImages:"), commandBuffer, objectivec.IObjectSliceToNSArray(sourceImages))
	return MPSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/encode(to:sourceImages:sourceStates:intermediateImages:destinationStates:)
func (g MPSNNGraph) EncodeToCommandBufferSourceImagesSourceStatesIntermediateImagesDestinationStates(commandBuffer metal.MTLCommandBuffer, sourceImages []MPSImage, sourceStates []MPSState, intermediateImages foundation.INSArray, destinationStates foundation.INSArray) IMPSImage {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("encodeToCommandBuffer:sourceImages:sourceStates:intermediateImages:destinationStates:"), commandBuffer, objectivec.IObjectSliceToNSArray(sourceImages), objectivec.IObjectSliceToNSArray(sourceStates), intermediateImages, destinationStates)
	return MPSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/executeAsync(withSourceImages:completionHandler:)
func (g MPSNNGraph) ExecuteAsyncWithSourceImagesCompletionHandler(sourceImages []MPSImage, handler MPSImageErrorHandler) IMPSImage {
	_block1, _ := NewMPSImageErrorBlock(handler)
	rv := objc.Send[objc.ID](g.ID, objc.Sel("executeAsyncWithSourceImages:completionHandler:"), sourceImages, _block1)
	return MPSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/encodeBatch(to:sourceImages:sourceStates:)
func (g MPSNNGraph) EncodeBatchToCommandBufferSourceImagesSourceStates(commandBuffer metal.MTLCommandBuffer, sourceImages []foundation.NSArray, sourceStates []foundation.NSArray) MPSImageBatch {
	rv := objc.Send[MPSImageBatch](g.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:sourceStates:"), commandBuffer, objectivec.IObjectSliceToNSArray(sourceImages), objectivec.IObjectSliceToNSArray(sourceStates))
	return MPSImageBatch(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/encodeBatch(to:sourceImages:sourceStates:intermediateImages:destinationStates:)
func (g MPSNNGraph) EncodeBatchToCommandBufferSourceImagesSourceStatesIntermediateImagesDestinationStates(commandBuffer metal.MTLCommandBuffer, sourceImages []foundation.NSArray, sourceStates []foundation.NSArray, intermediateImages foundation.INSArray, destinationStates foundation.INSArray) MPSImageBatch {
	rv := objc.Send[MPSImageBatch](g.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:sourceStates:intermediateImages:destinationStates:"), commandBuffer, objectivec.IObjectSliceToNSArray(sourceImages), objectivec.IObjectSliceToNSArray(sourceStates), intermediateImages, destinationStates)
	return MPSImageBatch(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/readCountForSourceImage(at:)
func (g MPSNNGraph) ReadCountForSourceImageAtIndex(index uint) uint {
	rv := objc.Send[uint](g.ID, objc.Sel("readCountForSourceImageAtIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/readCountForSourceState(at:)
func (g MPSNNGraph) ReadCountForSourceStateAtIndex(index uint) uint {
	rv := objc.Send[uint](g.ID, objc.Sel("readCountForSourceStateAtIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/reloadFromDataSources()
func (g MPSNNGraph) ReloadFromDataSources() {
	objc.Send[objc.ID](g.ID, objc.Sel("reloadFromDataSources"))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/graphWithDevice:resultImage:resultImageIsNeeded:
func (_MPSNNGraphClass MPSNNGraphClass) GraphWithDeviceResultImageResultImageIsNeeded(device metal.MTLDevice, resultImage IMPSNNImageNode, resultIsNeeded bool) MPSNNGraph {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNGraphClass.class), objc.Sel("graphWithDevice:resultImage:resultImageIsNeeded:"), device, resultImage, resultIsNeeded)
	return MPSNNGraphFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/graphWithDevice:resultImages:resultsAreNeeded:
func (_MPSNNGraphClass MPSNNGraphClass) GraphWithDeviceResultImagesResultsAreNeeded(device metal.MTLDevice, resultImages []MPSNNImageNode, areResultsNeeded *bool) MPSNNGraph {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNGraphClass.class), objc.Sel("graphWithDevice:resultImages:resultsAreNeeded:"), device, objectivec.IObjectSliceToNSArray(resultImages), areResultsNeeded)
	return MPSNNGraphFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/destinationImageAllocator
func (g MPSNNGraph) DestinationImageAllocator() MPSImageAllocator {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("destinationImageAllocator"))
	return MPSImageAllocatorObjectFromID(rv)
}
func (g MPSNNGraph) SetDestinationImageAllocator(value MPSImageAllocator) {
	objc.Send[struct{}](g.ID, objc.Sel("setDestinationImageAllocator:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/intermediateImageHandles
func (g MPSNNGraph) IntermediateImageHandles() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("intermediateImageHandles"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/outputStateIsTemporary
func (g MPSNNGraph) OutputStateIsTemporary() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("outputStateIsTemporary"))
	return rv
}
func (g MPSNNGraph) SetOutputStateIsTemporary(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setOutputStateIsTemporary:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/resultHandle
func (g MPSNNGraph) ResultHandle() MPSHandle {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resultHandle"))
	return MPSHandleObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/resultStateHandles
func (g MPSNNGraph) ResultStateHandles() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("resultStateHandles"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/sourceImageHandles
func (g MPSNNGraph) SourceImageHandles() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("sourceImageHandles"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/sourceStateHandles
func (g MPSNNGraph) SourceStateHandles() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("sourceStateHandles"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/format
func (g MPSNNGraph) Format() MPSImageFeatureChannelFormat {
	rv := objc.Send[MPSImageFeatureChannelFormat](g.ID, objc.Sel("format"))
	return MPSImageFeatureChannelFormat(rv)
}
func (g MPSNNGraph) SetFormat(value MPSImageFeatureChannelFormat) {
	objc.Send[struct{}](g.ID, objc.Sel("setFormat:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNGraph/resultImageIsNeeded
func (g MPSNNGraph) ResultImageIsNeeded() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("resultImageIsNeeded"))
	return rv
}
