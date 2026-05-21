// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoDataFrameExecutor] class.
var (
	_EspressoDataFrameExecutorClass     EspressoDataFrameExecutorClass
	_EspressoDataFrameExecutorClassOnce sync.Once
)

func getEspressoDataFrameExecutorClass() EspressoDataFrameExecutorClass {
	_EspressoDataFrameExecutorClassOnce.Do(func() {
		_EspressoDataFrameExecutorClass = EspressoDataFrameExecutorClass{class: objc.GetClass("EspressoDataFrameExecutor")}
	})
	return _EspressoDataFrameExecutorClass
}

// GetEspressoDataFrameExecutorClass returns the class object for EspressoDataFrameExecutor.
func GetEspressoDataFrameExecutorClass() EspressoDataFrameExecutorClass {
	return getEspressoDataFrameExecutorClass()
}

type EspressoDataFrameExecutorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoDataFrameExecutorClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoDataFrameExecutorClass) Alloc() EspressoDataFrameExecutor {
	rv := objc.Send[EspressoDataFrameExecutor](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [EspressoDataFrameExecutor.BindInputFromImageAttachmentToNetwork]
//   - [EspressoDataFrameExecutor.BindInputFromTensorAttachmentToNetwork]
//   - [EspressoDataFrameExecutor.BindInputsFromFrameToNetwork]
//   - [EspressoDataFrameExecutor.BindOutputsFromFrameToNetwork]
//   - [EspressoDataFrameExecutor.BindOutputsFromFrameToNetworkExecutionStatus]
//   - [EspressoDataFrameExecutor.BindOutputsFromFrameToNetworkReferenceNetwork]
//   - [EspressoDataFrameExecutor.FreeTemporaryResources]
//   - [EspressoDataFrameExecutor.OutputMatchingBuffers]
//   - [EspressoDataFrameExecutor.SetOutputMatchingBuffers]
//   - [EspressoDataFrameExecutor.UseCVPixelBuffers]
//   - [EspressoDataFrameExecutor.UseCVPixelBuffersForOutputs]
//   - [EspressoDataFrameExecutor.Use_cvpixelbuffer]
//   - [EspressoDataFrameExecutor.SetUse_cvpixelbuffer]
type EspressoDataFrameExecutor struct {
	objectivec.Object
}

// EspressoDataFrameExecutorFromID constructs a [EspressoDataFrameExecutor] from an objc.ID.
func EspressoDataFrameExecutorFromID(id objc.ID) EspressoDataFrameExecutor {
	return EspressoDataFrameExecutor{objectivec.Object{ID: id}}
}

// Ensure EspressoDataFrameExecutor implements IEspressoDataFrameExecutor.
var _ IEspressoDataFrameExecutor = EspressoDataFrameExecutor{}

// An interface definition for the [EspressoDataFrameExecutor] class.
//
// # Methods
//
//   - [IEspressoDataFrameExecutor.BindInputFromImageAttachmentToNetwork]
//   - [IEspressoDataFrameExecutor.BindInputFromTensorAttachmentToNetwork]
//   - [IEspressoDataFrameExecutor.BindInputsFromFrameToNetwork]
//   - [IEspressoDataFrameExecutor.BindOutputsFromFrameToNetwork]
//   - [IEspressoDataFrameExecutor.BindOutputsFromFrameToNetworkExecutionStatus]
//   - [IEspressoDataFrameExecutor.BindOutputsFromFrameToNetworkReferenceNetwork]
//   - [IEspressoDataFrameExecutor.FreeTemporaryResources]
//   - [IEspressoDataFrameExecutor.OutputMatchingBuffers]
//   - [IEspressoDataFrameExecutor.SetOutputMatchingBuffers]
//   - [IEspressoDataFrameExecutor.UseCVPixelBuffers]
//   - [IEspressoDataFrameExecutor.UseCVPixelBuffersForOutputs]
//   - [IEspressoDataFrameExecutor.Use_cvpixelbuffer]
//   - [IEspressoDataFrameExecutor.SetUse_cvpixelbuffer]
type IEspressoDataFrameExecutor interface {
	objectivec.IObject

	// Topic: Methods

	BindInputFromImageAttachmentToNetwork(input objectivec.IObject, attachment objectivec.IObject, network unsafe.Pointer) int
	BindInputFromTensorAttachmentToNetwork(input objectivec.IObject, attachment objectivec.IObject, network unsafe.Pointer) int
	BindInputsFromFrameToNetwork(frame objectivec.IObject, network unsafe.Pointer) int
	BindOutputsFromFrameToNetwork(frame objectivec.IObject, network unsafe.Pointer) int
	BindOutputsFromFrameToNetworkExecutionStatus(frame objectivec.IObject, network unsafe.Pointer, status int) int
	BindOutputsFromFrameToNetworkReferenceNetwork(frame objectivec.IObject, network unsafe.Pointer, network2 unsafe.Pointer) int
	FreeTemporaryResources()
	OutputMatchingBuffers() foundation.INSArray
	SetOutputMatchingBuffers(value foundation.INSArray)
	UseCVPixelBuffers() bool
	UseCVPixelBuffersForOutputs(outputs bool) bool
	Use_cvpixelbuffer() int
	SetUse_cvpixelbuffer(value int)
}

// Init initializes the instance.
func (e EspressoDataFrameExecutor) Init() EspressoDataFrameExecutor {
	rv := objc.Send[EspressoDataFrameExecutor](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoDataFrameExecutor) Autorelease() EspressoDataFrameExecutor {
	rv := objc.Send[EspressoDataFrameExecutor](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoDataFrameExecutor creates a new EspressoDataFrameExecutor instance.
func NewEspressoDataFrameExecutor() EspressoDataFrameExecutor {
	class := getEspressoDataFrameExecutorClass()
	rv := objc.Send[EspressoDataFrameExecutor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (e EspressoDataFrameExecutor) BindInputFromImageAttachmentToNetwork(input objectivec.IObject, attachment objectivec.IObject, network unsafe.Pointer) int {
	rv := objc.Send[int](e.ID, objc.Sel("bindInput:fromImageAttachment:toNetwork:"), input, attachment, network)
	return rv
}
func (e EspressoDataFrameExecutor) BindInputFromTensorAttachmentToNetwork(input objectivec.IObject, attachment objectivec.IObject, network unsafe.Pointer) int {
	rv := objc.Send[int](e.ID, objc.Sel("bindInput:fromTensorAttachment:toNetwork:"), input, attachment, network)
	return rv
}
func (e EspressoDataFrameExecutor) BindInputsFromFrameToNetwork(frame objectivec.IObject, network unsafe.Pointer) int {
	rv := objc.Send[int](e.ID, objc.Sel("bindInputsFromFrame:toNetwork:"), frame, network)
	return rv
}
func (e EspressoDataFrameExecutor) BindOutputsFromFrameToNetwork(frame objectivec.IObject, network unsafe.Pointer) int {
	rv := objc.Send[int](e.ID, objc.Sel("bindOutputsFromFrame:toNetwork:"), frame, network)
	return rv
}
func (e EspressoDataFrameExecutor) BindOutputsFromFrameToNetworkExecutionStatus(frame objectivec.IObject, network unsafe.Pointer, status int) int {
	rv := objc.Send[int](e.ID, objc.Sel("bindOutputsFromFrame:toNetwork:executionStatus:"), frame, network, status)
	return rv
}
func (e EspressoDataFrameExecutor) BindOutputsFromFrameToNetworkReferenceNetwork(frame objectivec.IObject, network unsafe.Pointer, network2 unsafe.Pointer) int {
	rv := objc.Send[int](e.ID, objc.Sel("bindOutputsFromFrame:toNetwork:referenceNetwork:"), frame, network, network2)
	return rv
}
func (e EspressoDataFrameExecutor) FreeTemporaryResources() {
	objc.Send[objc.ID](e.ID, objc.Sel("freeTemporaryResources"))
}
func (e EspressoDataFrameExecutor) UseCVPixelBuffers() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("useCVPixelBuffers"))
	return rv
}
func (e EspressoDataFrameExecutor) UseCVPixelBuffersForOutputs(outputs bool) bool {
	rv := objc.Send[bool](e.ID, objc.Sel("useCVPixelBuffersForOutputs:"), outputs)
	return rv
}

func (e EspressoDataFrameExecutor) OutputMatchingBuffers() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("outputMatchingBuffers"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e EspressoDataFrameExecutor) SetOutputMatchingBuffers(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setOutputMatchingBuffers:"), value)
}
func (e EspressoDataFrameExecutor) Use_cvpixelbuffer() int {
	rv := objc.Send[int](e.ID, objc.Sel("use_cvpixelbuffer"))
	return rv
}
func (e EspressoDataFrameExecutor) SetUse_cvpixelbuffer(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setUse_cvpixelbuffer:"), value)
}
