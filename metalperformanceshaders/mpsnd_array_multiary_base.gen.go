// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNDArrayMultiaryBase] class.
var (
	_MPSNDArrayMultiaryBaseClass     MPSNDArrayMultiaryBaseClass
	_MPSNDArrayMultiaryBaseClassOnce sync.Once
)

func getMPSNDArrayMultiaryBaseClass() MPSNDArrayMultiaryBaseClass {
	_MPSNDArrayMultiaryBaseClassOnce.Do(func() {
		_MPSNDArrayMultiaryBaseClass = MPSNDArrayMultiaryBaseClass{class: objc.GetClass("MPSNDArrayMultiaryBase")}
	})
	return _MPSNDArrayMultiaryBaseClass
}

// GetMPSNDArrayMultiaryBaseClass returns the class object for MPSNDArrayMultiaryBase.
func GetMPSNDArrayMultiaryBaseClass() MPSNDArrayMultiaryBaseClass {
	return getMPSNDArrayMultiaryBaseClass()
}

type MPSNDArrayMultiaryBaseClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSNDArrayMultiaryBaseClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNDArrayMultiaryBaseClass) Alloc() MPSNDArrayMultiaryBase {
	rv := objc.Send[MPSNDArrayMultiaryBase](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSNDArrayMultiaryBase.InitWithDeviceSourceCount]
//
// # Instance Properties
//
//   - [MPSNDArrayMultiaryBase.DestinationArrayAllocator]
//   - [MPSNDArrayMultiaryBase.SetDestinationArrayAllocator]
//
// # Instance Methods
//
//   - [MPSNDArrayMultiaryBase.DestinationArrayDescriptorForSourceArraysSourceState]
//   - [MPSNDArrayMultiaryBase.ResultStateForSourceArraysSourceStatesDestinationArray]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryBase
type MPSNDArrayMultiaryBase struct {
	MPSKernel
}

// MPSNDArrayMultiaryBaseFromID constructs a [MPSNDArrayMultiaryBase] from an objc.ID.
func MPSNDArrayMultiaryBaseFromID(id objc.ID) MPSNDArrayMultiaryBase {
	return MPSNDArrayMultiaryBase{MPSKernel: MPSKernelFromID(id)}
}

// NOTE: MPSNDArrayMultiaryBase adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNDArrayMultiaryBase] class.
//
// # Initializers
//
//   - [IMPSNDArrayMultiaryBase.InitWithDeviceSourceCount]
//
// # Instance Properties
//
//   - [IMPSNDArrayMultiaryBase.DestinationArrayAllocator]
//   - [IMPSNDArrayMultiaryBase.SetDestinationArrayAllocator]
//
// # Instance Methods
//
//   - [IMPSNDArrayMultiaryBase.DestinationArrayDescriptorForSourceArraysSourceState]
//   - [IMPSNDArrayMultiaryBase.ResultStateForSourceArraysSourceStatesDestinationArray]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryBase
type IMPSNDArrayMultiaryBase interface {
	IMPSKernel

	// Topic: Initializers

	InitWithDeviceSourceCount(device metal.MTLDevice, count uint) MPSNDArrayMultiaryBase

	// Topic: Instance Properties

	DestinationArrayAllocator() MPSNDArrayAllocator
	SetDestinationArrayAllocator(value MPSNDArrayAllocator)

	// Topic: Instance Methods

	DestinationArrayDescriptorForSourceArraysSourceState(sources []MPSNDArray, state IMPSState) IMPSNDArrayDescriptor
	ResultStateForSourceArraysSourceStatesDestinationArray(sourceArrays []MPSNDArray, sourceStates []MPSState, destinationArray IMPSNDArray) IMPSState
}

// Init initializes the instance.
func (n MPSNDArrayMultiaryBase) Init() MPSNDArrayMultiaryBase {
	rv := objc.Send[MPSNDArrayMultiaryBase](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MPSNDArrayMultiaryBase) Autorelease() MPSNDArrayMultiaryBase {
	rv := objc.Send[MPSNDArrayMultiaryBase](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNDArrayMultiaryBase creates a new MPSNDArrayMultiaryBase instance.
func NewMPSNDArrayMultiaryBase() MPSNDArrayMultiaryBase {
	class := getMPSNDArrayMultiaryBaseClass()
	rv := objc.Send[MPSNDArrayMultiaryBase](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel/init(coder:)
func NewNDArrayMultiaryBaseWithCoder(aDecoder foundation.INSCoder) MPSNDArrayMultiaryBase {
	instance := getMPSNDArrayMultiaryBaseClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return MPSNDArrayMultiaryBaseFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryBase/init(coder:device:)
func NewNDArrayMultiaryBaseWithCoderDevice(coder foundation.INSCoder, device metal.MTLDevice) MPSNDArrayMultiaryBase {
	instance := getMPSNDArrayMultiaryBaseClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	return MPSNDArrayMultiaryBaseFromID(rv)
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
func NewNDArrayMultiaryBaseWithDevice(device metal.MTLDevice) MPSNDArrayMultiaryBase {
	instance := getMPSNDArrayMultiaryBaseClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSNDArrayMultiaryBaseFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryBase/init(device:sourceCount:)
func NewNDArrayMultiaryBaseWithDeviceSourceCount(device metal.MTLDevice, count uint) MPSNDArrayMultiaryBase {
	instance := getMPSNDArrayMultiaryBaseClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:sourceCount:"), device, count)
	return MPSNDArrayMultiaryBaseFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryBase/init(device:sourceCount:)
func (n MPSNDArrayMultiaryBase) InitWithDeviceSourceCount(device metal.MTLDevice, count uint) MPSNDArrayMultiaryBase {
	rv := objc.Send[MPSNDArrayMultiaryBase](n.ID, objc.Sel("initWithDevice:sourceCount:"), device, count)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryBase/destinationArrayDescriptor(forSourceArrays:sourceState:)
func (n MPSNDArrayMultiaryBase) DestinationArrayDescriptorForSourceArraysSourceState(sources []MPSNDArray, state IMPSState) IMPSNDArrayDescriptor {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("destinationArrayDescriptorForSourceArrays:sourceState:"), objectivec.IObjectSliceToNSArray(sources), state)
	return MPSNDArrayDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryBase/resultState(forSourceArrays:sourceStates:destinationArray:)
func (n MPSNDArrayMultiaryBase) ResultStateForSourceArraysSourceStatesDestinationArray(sourceArrays []MPSNDArray, sourceStates []MPSState, destinationArray IMPSNDArray) IMPSState {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("resultStateForSourceArrays:sourceStates:destinationArray:"), objectivec.IObjectSliceToNSArray(sourceArrays), objectivec.IObjectSliceToNSArray(sourceStates), destinationArray)
	return MPSStateFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryBase/destinationArrayAllocator
func (n MPSNDArrayMultiaryBase) DestinationArrayAllocator() MPSNDArrayAllocator {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("destinationArrayAllocator"))
	return MPSNDArrayAllocatorObjectFromID(rv)
}
func (n MPSNDArrayMultiaryBase) SetDestinationArrayAllocator(value MPSNDArrayAllocator) {
	objc.Send[struct{}](n.ID, objc.Sel("setDestinationArrayAllocator:"), value)
}
