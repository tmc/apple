// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSState] class.
var (
	_MPSStateClass     MPSStateClass
	_MPSStateClassOnce sync.Once
)

func getMPSStateClass() MPSStateClass {
	_MPSStateClassOnce.Do(func() {
		_MPSStateClass = MPSStateClass{class: objc.GetClass("MPSState")}
	})
	return _MPSStateClass
}

// GetMPSStateClass returns the class object for MPSState.
func GetMPSStateClass() MPSStateClass {
	return getMPSStateClass()
}

type MPSStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSStateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSStateClass) Alloc() MPSState {
	rv := objc.Send[MPSState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An opaque data container for large storage in MPS CNN filters.
//
// # Overview
//
// Some MPS CNN kernels produce additional information beyond an [MPSImage].
// These may be pooling indices where the result came from, convolution
// weights, or other information not contained in the usual [MPSImage] result
// from a [MPSCNNKernel]. An [MPSState] object typically contains one or more
// expensive [MTLResource] objects such as textures or buffers to store this
// information. It provides a base class with interfaces for managing this
// storage. Child classes may add additional functionality specific to their
// contents.
//
// Some [MPSState] objects are temporary. Temporary state objects, for
// example, [MPSTemporaryImage] and [MPSTemporaryMatrix], are for very short
// lived storage, perhaps just a few lines of code within the scope of a
// single [MTLCommandBuffer]. They are very efficient for storage, as several
// temporary objects can share the same memory over the course of a command
// buffer. This can improve both memory usage and time spent in the kernel
// wiring down memory and such. You may find that some large CNN tasks can not
// be computed without them, as nontemporary storage would simply take up too
// much memory.
//
// In exchange, the lifetime of the underlying storage in temporary [MPSState]
// objects needs to be carefully managed. ARC often waits until the end of
// scope to release objects. Temporary storage often needs to be released
// sooner than that. Consequently the lifetime of the data in the underlying
// Metal resources is managed by a [MPSState.ReadCount] property. Each time a
// [MPSCNNKernel] reads a temporary [MPSState] object the [MPSState.ReadCount]
// is automatically decremented. When it reaches 0, the underlying storage is
// recycled for use by other MPS temporary objects, and the data is becomes
// undefined. If you need to consume the data multiple times, you should set
// the [MPSState.ReadCount] to a larger number to prevent the data from
// becoming undefined. You may set the [MPSState.ReadCount] to 0 yourself to
// return the storage to MPS, if for any reason, you realize that the
// [MPSState] object will no longer be used.
//
// The contents of a temporary [MPSState] object are only valid from creation
// to the time the [MPSState.ReadCount] reaches 0. The data is only valid for
// the [MTLCommandBuffer] on which it was created. Nontemporary [MPSState]
// objects are valid on any [MTLCommandBuffer] on the same device until they
// are released.
//
// # Instance Properties
//
//   - [MPSState.IsTemporary]
//   - [MPSState.Label]
//   - [MPSState.SetLabel]
//   - [MPSState.ReadCount]
//   - [MPSState.SetReadCount]
//   - [MPSState.ResourceCount]
//
// # Initializers
//
//   - [MPSState.InitWithDeviceBufferSize]
//   - [MPSState.InitWithDeviceResourceList]
//   - [MPSState.InitWithDeviceTextureDescriptor]
//   - [MPSState.InitWithResource]
//   - [MPSState.InitWithResources]
//
// # Instance Methods
//
//   - [MPSState.BufferSizeAtIndex]
//   - [MPSState.DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor]
//   - [MPSState.ResourceAtIndexAllocateMemory]
//   - [MPSState.ResourceSize]
//   - [MPSState.ResourceTypeAtIndex]
//   - [MPSState.SynchronizeOnCommandBuffer]
//   - [MPSState.TextureInfoAtIndex]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState
//
// [MTLCommandBuffer]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer
// [MTLResource]: https://developer.apple.com/documentation/Metal/MTLResource
type MPSState struct {
	objectivec.Object
}

// MPSStateFromID constructs a [MPSState] from an objc.ID.
//
// An opaque data container for large storage in MPS CNN filters.
func MPSStateFromID(id objc.ID) MPSState {
	return MPSState{objectivec.Object{ID: id}}
}

// NOTE: MPSState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSState] class.
//
// # Instance Properties
//
//   - [IMPSState.IsTemporary]
//   - [IMPSState.Label]
//   - [IMPSState.SetLabel]
//   - [IMPSState.ReadCount]
//   - [IMPSState.SetReadCount]
//   - [IMPSState.ResourceCount]
//
// # Initializers
//
//   - [IMPSState.InitWithDeviceBufferSize]
//   - [IMPSState.InitWithDeviceResourceList]
//   - [IMPSState.InitWithDeviceTextureDescriptor]
//   - [IMPSState.InitWithResource]
//   - [IMPSState.InitWithResources]
//
// # Instance Methods
//
//   - [IMPSState.BufferSizeAtIndex]
//   - [IMPSState.DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor]
//   - [IMPSState.ResourceAtIndexAllocateMemory]
//   - [IMPSState.ResourceSize]
//   - [IMPSState.ResourceTypeAtIndex]
//   - [IMPSState.SynchronizeOnCommandBuffer]
//   - [IMPSState.TextureInfoAtIndex]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState
type IMPSState interface {
	objectivec.IObject

	// Topic: Instance Properties

	IsTemporary() bool
	Label() string
	SetLabel(value string)
	ReadCount() uint
	SetReadCount(value uint)
	ResourceCount() uint

	// Topic: Initializers

	InitWithDeviceBufferSize(device metal.MTLDevice, bufferSize uintptr) MPSState
	InitWithDeviceResourceList(device metal.MTLDevice, resourceList IMPSStateResourceList) MPSState
	InitWithDeviceTextureDescriptor(device metal.MTLDevice, descriptor metal.MTLTextureDescriptor) MPSState
	InitWithResource(resource metal.MTLResource) MPSState
	InitWithResources(resources []objectivec.IObject) MPSState

	// Topic: Instance Methods

	BufferSizeAtIndex(index uint) uint
	DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor(sourceImages []MPSImage, sourceStates []MPSState, kernel IMPSKernel, inDescriptor IMPSImageDescriptor) IMPSImageDescriptor
	ResourceAtIndexAllocateMemory(index uint, allocateMemory bool) metal.MTLResource
	ResourceSize() uint
	ResourceTypeAtIndex(index uint) MPSStateResourceType
	SynchronizeOnCommandBuffer(commandBuffer metal.MTLCommandBuffer)
	TextureInfoAtIndex(index uint) MPSStateTextureInfo
}

// Init initializes the instance.
func (s MPSState) Init() MPSState {
	rv := objc.Send[MPSState](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MPSState) Autorelease() MPSState {
	rv := objc.Send[MPSState](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSState creates a new MPSState instance.
func NewMPSState() MPSState {
	class := getMPSStateClass()
	rv := objc.Send[MPSState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:bufferSize:)
func NewStateWithDeviceBufferSize(device metal.MTLDevice, bufferSize uintptr) MPSState {
	instance := getMPSStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:bufferSize:"), device, bufferSize)
	return MPSStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:resourceList:)
func NewStateWithDeviceResourceList(device metal.MTLDevice, resourceList IMPSStateResourceList) MPSState {
	instance := getMPSStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:resourceList:"), device, resourceList)
	return MPSStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:textureDescriptor:)
func NewStateWithDeviceTextureDescriptor(device metal.MTLDevice, descriptor metal.MTLTextureDescriptor) MPSState {
	instance := getMPSStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:textureDescriptor:"), device, descriptor)
	return MPSStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resource:)
func NewStateWithResource(resource metal.MTLResource) MPSState {
	instance := getMPSStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return MPSStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resources:)
func NewStateWithResources(resources []objectivec.IObject) MPSState {
	instance := getMPSStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithResources:"), objectivec.IObjectSliceToNSArray(resources))
	return MPSStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:bufferSize:)
func (s MPSState) InitWithDeviceBufferSize(device metal.MTLDevice, bufferSize uintptr) MPSState {
	rv := objc.Send[MPSState](s.ID, objc.Sel("initWithDevice:bufferSize:"), device, bufferSize)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:resourceList:)
func (s MPSState) InitWithDeviceResourceList(device metal.MTLDevice, resourceList IMPSStateResourceList) MPSState {
	rv := objc.Send[MPSState](s.ID, objc.Sel("initWithDevice:resourceList:"), device, resourceList)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(device:textureDescriptor:)
func (s MPSState) InitWithDeviceTextureDescriptor(device metal.MTLDevice, descriptor metal.MTLTextureDescriptor) MPSState {
	rv := objc.Send[MPSState](s.ID, objc.Sel("initWithDevice:textureDescriptor:"), device, descriptor)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resource:)
func (s MPSState) InitWithResource(resource metal.MTLResource) MPSState {
	rv := objc.Send[MPSState](s.ID, objc.Sel("initWithResource:"), resource)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/init(resources:)
func (s MPSState) InitWithResources(resources []objectivec.IObject) MPSState {
	rv := objc.Send[MPSState](s.ID, objc.Sel("initWithResources:"), objectivec.IObjectSliceToNSArray(resources))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/bufferSize(at:)
func (s MPSState) BufferSizeAtIndex(index uint) uint {
	rv := objc.Send[uint](s.ID, objc.Sel("bufferSizeAtIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/destinationImageDescriptor(forSourceImages:sourceStates:for:suggestedDescriptor:)
func (s MPSState) DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor(sourceImages []MPSImage, sourceStates []MPSState, kernel IMPSKernel, inDescriptor IMPSImageDescriptor) IMPSImageDescriptor {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:forKernel:suggestedDescriptor:"), objectivec.IObjectSliceToNSArray(sourceImages), objectivec.IObjectSliceToNSArray(sourceStates), kernel, inDescriptor)
	return MPSImageDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/resource(at:allocateMemory:)
func (s MPSState) ResourceAtIndexAllocateMemory(index uint, allocateMemory bool) metal.MTLResource {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("resourceAtIndex:allocateMemory:"), index, allocateMemory)
	return metal.MTLResourceObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/resourceSize()
func (s MPSState) ResourceSize() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("resourceSize"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/resourceType(at:)
func (s MPSState) ResourceTypeAtIndex(index uint) MPSStateResourceType {
	rv := objc.Send[MPSStateResourceType](s.ID, objc.Sel("resourceTypeAtIndex:"), index)
	return MPSStateResourceType(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/synchronize(on:)
func (s MPSState) SynchronizeOnCommandBuffer(commandBuffer metal.MTLCommandBuffer) {
	objc.Send[objc.ID](s.ID, objc.Sel("synchronizeOnCommandBuffer:"), commandBuffer)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/textureInfo(at:)
func (s MPSState) TextureInfoAtIndex(index uint) MPSStateTextureInfo {
	rv := objc.Send[MPSStateTextureInfo](s.ID, objc.Sel("textureInfoAtIndex:"), index)
	return MPSStateTextureInfo(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/temporaryState(with:)
func (_MPSStateClass MPSStateClass) TemporaryStateWithCommandBuffer(cmdBuf metal.MTLCommandBuffer) MPSState {
	rv := objc.Send[objc.ID](objc.ID(_MPSStateClass.class), objc.Sel("temporaryStateWithCommandBuffer:"), cmdBuf)
	return MPSStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/temporaryState(with:bufferSize:)
func (_MPSStateClass MPSStateClass) TemporaryStateWithCommandBufferBufferSize(cmdBuf metal.MTLCommandBuffer, bufferSize uintptr) MPSState {
	rv := objc.Send[objc.ID](objc.ID(_MPSStateClass.class), objc.Sel("temporaryStateWithCommandBuffer:bufferSize:"), cmdBuf, bufferSize)
	return MPSStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/temporaryState(with:resourceList:)
func (_MPSStateClass MPSStateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.MTLCommandBuffer, resourceList IMPSStateResourceList) MPSState {
	rv := objc.Send[objc.ID](objc.ID(_MPSStateClass.class), objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), commandBuffer, resourceList)
	return MPSStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/temporaryState(with:textureDescriptor:)
func (_MPSStateClass MPSStateClass) TemporaryStateWithCommandBufferTextureDescriptor(cmdBuf metal.MTLCommandBuffer, descriptor metal.MTLTextureDescriptor) MPSState {
	rv := objc.Send[objc.ID](objc.ID(_MPSStateClass.class), objc.Sel("temporaryStateWithCommandBuffer:textureDescriptor:"), cmdBuf, descriptor)
	return MPSStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/isTemporary
func (s MPSState) IsTemporary() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isTemporary"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/label
func (s MPSState) Label() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (s MPSState) SetLabel(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setLabel:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/readCount
func (s MPSState) ReadCount() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("readCount"))
	return rv
}
func (s MPSState) SetReadCount(value uint) {
	objc.Send[struct{}](s.ID, objc.Sel("setReadCount:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState/resourceCount
func (s MPSState) ResourceCount() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("resourceCount"))
	return rv
}
