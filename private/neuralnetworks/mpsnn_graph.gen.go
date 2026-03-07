// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNGraphClass) Alloc() MPSNNGraph {
	rv := objc.Send[MPSNNGraph](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNGraph.CopyWithZoneDevice]
//   - [MPSNNGraph.DestinationImageAllocator]
//   - [MPSNNGraph.SetDestinationImageAllocator]
//   - [MPSNNGraph.EncodeBatchToCommandBufferSourceImagesSourceStates]
//   - [MPSNNGraph.EncodeBatchToCommandBufferSourceImagesSourceStatesIntermediateImagesDestinationStates]
//   - [MPSNNGraph.EncodeToCommandBufferSourceImages]
//   - [MPSNNGraph.EncodeToCommandBufferSourceImagesSourceStatesIntermediateImagesDestinationStates]
//   - [MPSNNGraph.ExecuteAsyncWithSourceImagesCompletionHandler]
//   - [MPSNNGraph.Format]
//   - [MPSNNGraph.SetFormat]
//   - [MPSNNGraph.IntermediateImageHandles]
//   - [MPSNNGraph.OutputStateIsTemporary]
//   - [MPSNNGraph.SetOutputStateIsTemporary]
//   - [MPSNNGraph.ReadCountForSourceImageAtIndex]
//   - [MPSNNGraph.ReadCountForSourceStateAtIndex]
//   - [MPSNNGraph.ReloadFromDataSources]
//   - [MPSNNGraph.ResultHandle]
//   - [MPSNNGraph.ResultImageIsNeeded]
//   - [MPSNNGraph.ResultStateHandles]
//   - [MPSNNGraph.SetOptions]
//   - [MPSNNGraph.SourceImageHandles]
//   - [MPSNNGraph.SourceStateHandles]
//   - [MPSNNGraph.InitWithCoderDevice]
//   - [MPSNNGraph.InitWithDeviceResultImage]
//   - [MPSNNGraph.InitWithDeviceResultImageResultImageIsNeeded]
//   - [MPSNNGraph.InitWithDeviceResultImagesResultsAreNeeded]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph
type MPSNNGraph struct {
	MPSKernel
}

// MPSNNGraphFromID constructs a [MPSNNGraph] from an objc.ID.
func MPSNNGraphFromID(id objc.ID) MPSNNGraph {
	return MPSNNGraph{MPSKernel: MPSKernelFromID(id)}
}
// Ensure MPSNNGraph implements IMPSNNGraph.
var _ IMPSNNGraph = MPSNNGraph{}





// An interface definition for the [MPSNNGraph] class.
//
// # Methods
//
//   - [IMPSNNGraph.CopyWithZoneDevice]
//   - [IMPSNNGraph.DestinationImageAllocator]
//   - [IMPSNNGraph.SetDestinationImageAllocator]
//   - [IMPSNNGraph.EncodeBatchToCommandBufferSourceImagesSourceStates]
//   - [IMPSNNGraph.EncodeBatchToCommandBufferSourceImagesSourceStatesIntermediateImagesDestinationStates]
//   - [IMPSNNGraph.EncodeToCommandBufferSourceImages]
//   - [IMPSNNGraph.EncodeToCommandBufferSourceImagesSourceStatesIntermediateImagesDestinationStates]
//   - [IMPSNNGraph.ExecuteAsyncWithSourceImagesCompletionHandler]
//   - [IMPSNNGraph.Format]
//   - [IMPSNNGraph.SetFormat]
//   - [IMPSNNGraph.IntermediateImageHandles]
//   - [IMPSNNGraph.OutputStateIsTemporary]
//   - [IMPSNNGraph.SetOutputStateIsTemporary]
//   - [IMPSNNGraph.ReadCountForSourceImageAtIndex]
//   - [IMPSNNGraph.ReadCountForSourceStateAtIndex]
//   - [IMPSNNGraph.ReloadFromDataSources]
//   - [IMPSNNGraph.ResultHandle]
//   - [IMPSNNGraph.ResultImageIsNeeded]
//   - [IMPSNNGraph.ResultStateHandles]
//   - [IMPSNNGraph.SetOptions]
//   - [IMPSNNGraph.SourceImageHandles]
//   - [IMPSNNGraph.SourceStateHandles]
//   - [IMPSNNGraph.InitWithCoderDevice]
//   - [IMPSNNGraph.InitWithDeviceResultImage]
//   - [IMPSNNGraph.InitWithDeviceResultImageResultImageIsNeeded]
//   - [IMPSNNGraph.InitWithDeviceResultImagesResultsAreNeeded]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph
type IMPSNNGraph interface {
	IMPSKernel

	// Topic: Methods

	CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject
	DestinationImageAllocator() unsafe.Pointer
	SetDestinationImageAllocator(value unsafe.Pointer)
	EncodeBatchToCommandBufferSourceImagesSourceStates(buffer objectivec.IObject, images objectivec.IObject, states objectivec.IObject) objectivec.IObject
	EncodeBatchToCommandBufferSourceImagesSourceStatesIntermediateImagesDestinationStates(buffer objectivec.IObject, images objectivec.IObject, states objectivec.IObject, images2 objectivec.IObject, states2 objectivec.IObject) objectivec.IObject
	EncodeToCommandBufferSourceImages(buffer objectivec.IObject, images objectivec.IObject) objectivec.IObject
	EncodeToCommandBufferSourceImagesSourceStatesIntermediateImagesDestinationStates(buffer objectivec.IObject, images objectivec.IObject, states objectivec.IObject, images2 objectivec.IObject, states2 objectivec.IObject) objectivec.IObject
	ExecuteAsyncWithSourceImagesCompletionHandler(images objectivec.IObject, handler ErrorHandler) objectivec.IObject
	Format() uint64
	SetFormat(value uint64)
	IntermediateImageHandles() foundation.INSArray
	OutputStateIsTemporary() bool
	SetOutputStateIsTemporary(value bool)
	ReadCountForSourceImageAtIndex(index uint64) uint64
	ReadCountForSourceStateAtIndex(index uint64) uint64
	ReloadFromDataSources()
	ResultHandle() unsafe.Pointer
	ResultImageIsNeeded() bool
	ResultStateHandles() foundation.INSArray
	SetOptions(options uint64)
	SourceImageHandles() foundation.INSArray
	SourceStateHandles() foundation.INSArray
	InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNGraph
	InitWithDeviceResultImage(device objectivec.IObject, image objectivec.IObject) MPSNNGraph
	InitWithDeviceResultImageResultImageIsNeeded(device objectivec.IObject, image objectivec.IObject, needed bool) MPSNNGraph
	InitWithDeviceResultImagesResultsAreNeeded(device objectivec.IObject, images objectivec.IObject, needed unsafe.Pointer) MPSNNGraph
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






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/initWithCoder:device:
func NewGraphWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNGraph {
	instance := getMPSNNGraphClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNNGraphFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/initWithDevice:resultImage:
func NewGraphWithDeviceResultImage(device objectivec.IObject, image objectivec.IObject) MPSNNGraph {
	instance := getMPSNNGraphClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resultImage:"), device, image)
	return MPSNNGraphFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/initWithDevice:resultImage:resultImageIsNeeded:
func NewGraphWithDeviceResultImageResultImageIsNeeded(device objectivec.IObject, image objectivec.IObject, needed bool) MPSNNGraph {
	instance := getMPSNNGraphClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resultImage:resultImageIsNeeded:"), device, image, needed)
	return MPSNNGraphFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/initWithDevice:resultImages:resultsAreNeeded:
func NewGraphWithDeviceResultImagesResultsAreNeeded(device objectivec.IObject, images objectivec.IObject, needed unsafe.Pointer) MPSNNGraph {
	instance := getMPSNNGraphClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resultImages:resultsAreNeeded:"), device, images, needed)
	return MPSNNGraphFromID(rv)
}







//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/copyWithZone:device:
func (g MPSNNGraph) CopyWithZoneDevice(zone foundation.NSZone, device objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/encodeBatchToCommandBuffer:sourceImages:sourceStates:
func (g MPSNNGraph) EncodeBatchToCommandBufferSourceImagesSourceStates(buffer objectivec.IObject, images objectivec.IObject, states objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:sourceStates:"), buffer, images, states)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/encodeBatchToCommandBuffer:sourceImages:sourceStates:intermediateImages:destinationStates:
func (g MPSNNGraph) EncodeBatchToCommandBufferSourceImagesSourceStatesIntermediateImagesDestinationStates(buffer objectivec.IObject, images objectivec.IObject, states objectivec.IObject, images2 objectivec.IObject, states2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:sourceStates:intermediateImages:destinationStates:"), buffer, images, states, images2, states2)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/encodeToCommandBuffer:sourceImages:
func (g MPSNNGraph) EncodeToCommandBufferSourceImages(buffer objectivec.IObject, images objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("encodeToCommandBuffer:sourceImages:"), buffer, images)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/encodeToCommandBuffer:sourceImages:sourceStates:intermediateImages:destinationStates:
func (g MPSNNGraph) EncodeToCommandBufferSourceImagesSourceStatesIntermediateImagesDestinationStates(buffer objectivec.IObject, images objectivec.IObject, states objectivec.IObject, images2 objectivec.IObject, states2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("encodeToCommandBuffer:sourceImages:sourceStates:intermediateImages:destinationStates:"), buffer, images, states, images2, states2)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/executeAsyncWithSourceImages:completionHandler:
func (g MPSNNGraph) ExecuteAsyncWithSourceImagesCompletionHandler(images objectivec.IObject, handler ErrorHandler) objectivec.IObject {
		_block1, _cleanup1 := NewErrorBlock(handler)
	defer _cleanup1()
		rv := objc.Send[objc.ID](g.ID, objc.Sel("executeAsyncWithSourceImages:completionHandler:"), images, _block1)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/readCountForSourceImageAtIndex:
func (g MPSNNGraph) ReadCountForSourceImageAtIndex(index uint64) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("readCountForSourceImageAtIndex:"), index)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/readCountForSourceStateAtIndex:
func (g MPSNNGraph) ReadCountForSourceStateAtIndex(index uint64) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("readCountForSourceStateAtIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/reloadFromDataSources
func (g MPSNNGraph) ReloadFromDataSources() {
	objc.Send[objc.ID](g.ID, objc.Sel("reloadFromDataSources"))
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/setOptions:
func (g MPSNNGraph) SetOptions(options uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("setOptions:"), options)
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/initWithCoder:device:
func (g MPSNNGraph) InitWithCoderDevice(coder objectivec.IObject, device objectivec.IObject) MPSNNGraph {
	rv := objc.Send[MPSNNGraph](g.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/initWithDevice:resultImage:
func (g MPSNNGraph) InitWithDeviceResultImage(device objectivec.IObject, image objectivec.IObject) MPSNNGraph {
	rv := objc.Send[MPSNNGraph](g.ID, objc.Sel("initWithDevice:resultImage:"), device, image)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/initWithDevice:resultImage:resultImageIsNeeded:
func (g MPSNNGraph) InitWithDeviceResultImageResultImageIsNeeded(device objectivec.IObject, image objectivec.IObject, needed bool) MPSNNGraph {
	rv := objc.Send[MPSNNGraph](g.ID, objc.Sel("initWithDevice:resultImage:resultImageIsNeeded:"), device, image, needed)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/initWithDevice:resultImages:resultsAreNeeded:
func (g MPSNNGraph) InitWithDeviceResultImagesResultsAreNeeded(device objectivec.IObject, images objectivec.IObject, needed unsafe.Pointer) MPSNNGraph {
	rv := objc.Send[MPSNNGraph](g.ID, objc.Sel("initWithDevice:resultImages:resultsAreNeeded:"), device, images, needed)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/graphWithDevice:resultImage:
func (_MPSNNGraphClass MPSNNGraphClass) GraphWithDeviceResultImage(device objectivec.IObject, image objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNGraphClass.class), objc.Sel("graphWithDevice:resultImage:"), device, image)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/graphWithDevice:resultImage:resultImageIsNeeded:
func (_MPSNNGraphClass MPSNNGraphClass) GraphWithDeviceResultImageResultImageIsNeeded(device objectivec.IObject, image objectivec.IObject, needed bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNGraphClass.class), objc.Sel("graphWithDevice:resultImage:resultImageIsNeeded:"), device, image, needed)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/graphWithDevice:resultImages:resultsAreNeeded:
func (_MPSNNGraphClass MPSNNGraphClass) GraphWithDeviceResultImagesResultsAreNeeded(device objectivec.IObject, images objectivec.IObject, needed unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNGraphClass.class), objc.Sel("graphWithDevice:resultImages:resultsAreNeeded:"), device, images, needed)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/libraryInfo:
func (_MPSNNGraphClass MPSNNGraphClass) LibraryInfo(info unsafe.Pointer) MPSLibraryInfo {
	rv := objc.Send[MPSLibraryInfo](objc.ID(_MPSNNGraphClass.class), objc.Sel("libraryInfo:"), info)
	return MPSLibraryInfo(rv)
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/destinationImageAllocator
func (g MPSNNGraph) DestinationImageAllocator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("destinationImageAllocator"))
	return rv
}
func (g MPSNNGraph) SetDestinationImageAllocator(value unsafe.Pointer) {
	objc.Send[struct{}](g.ID, objc.Sel("setDestinationImageAllocator:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/format
func (g MPSNNGraph) Format() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("format"))
	return rv
}
func (g MPSNNGraph) SetFormat(value uint64) {
	objc.Send[struct{}](g.ID, objc.Sel("setFormat:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/intermediateImageHandles
func (g MPSNNGraph) IntermediateImageHandles() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("intermediateImageHandles"))
	return foundation.NSArrayFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/outputStateIsTemporary
func (g MPSNNGraph) OutputStateIsTemporary() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("outputStateIsTemporary"))
	return rv
}
func (g MPSNNGraph) SetOutputStateIsTemporary(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setOutputStateIsTemporary:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/resultHandle
func (g MPSNNGraph) ResultHandle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("resultHandle"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/resultImageIsNeeded
func (g MPSNNGraph) ResultImageIsNeeded() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("resultImageIsNeeded"))
	return rv
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/resultStateHandles
func (g MPSNNGraph) ResultStateHandles() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("resultStateHandles"))
	return foundation.NSArrayFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/sourceImageHandles
func (g MPSNNGraph) SourceImageHandles() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sourceImageHandles"))
	return foundation.NSArrayFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNGraph/sourceStateHandles
func (g MPSNNGraph) SourceStateHandles() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sourceStateHandles"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

















