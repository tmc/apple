// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLE5OutputPort] class.
var (
	_MLE5OutputPortClass     MLE5OutputPortClass
	_MLE5OutputPortClassOnce sync.Once
)

func getMLE5OutputPortClass() MLE5OutputPortClass {
	_MLE5OutputPortClassOnce.Do(func() {
		_MLE5OutputPortClass = MLE5OutputPortClass{class: objc.GetClass("MLE5OutputPort")}
	})
	return _MLE5OutputPortClass
}

// GetMLE5OutputPortClass returns the class object for MLE5OutputPort.
func GetMLE5OutputPortClass() MLE5OutputPortClass {
	return getMLE5OutputPortClass()
}

type MLE5OutputPortClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLE5OutputPortClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLE5OutputPortClass) Alloc() MLE5OutputPort {
	rv := objc.Send[MLE5OutputPort](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLE5OutputPort.Binder]
//   - [MLE5OutputPort.SetBinder]
//   - [MLE5OutputPort.BoundFeatureDirectly]
//   - [MLE5OutputPort.FeatureDescription]
//   - [MLE5OutputPort.FeatureValue]
//   - [MLE5OutputPort.Name]
//   - [MLE5OutputPort.OutputBackingWasDirectlyBound]
//   - [MLE5OutputPort.PixelBufferPool]
//   - [MLE5OutputPort.SetPixelBufferPool]
//   - [MLE5OutputPort.PortHandle]
//   - [MLE5OutputPort.PrepareWithOptionsError]
//   - [MLE5OutputPort.Reset]
//   - [MLE5OutputPort.ReusableForOutputBackingWillBindDirectly]
//   - [MLE5OutputPort.InitWithPortHandleNameFeatureDescription]
//   - [MLE5OutputPort.DebugDescription]
//   - [MLE5OutputPort.Description]
//   - [MLE5OutputPort.Hash]
//   - [MLE5OutputPort.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPort
type MLE5OutputPort struct {
	objectivec.Object
}

// MLE5OutputPortFromID constructs a [MLE5OutputPort] from an objc.ID.
func MLE5OutputPortFromID(id objc.ID) MLE5OutputPort {
	return MLE5OutputPort{objectivec.Object{ID: id}}
}

// Ensure MLE5OutputPort implements IMLE5OutputPort.
var _ IMLE5OutputPort = MLE5OutputPort{}

// An interface definition for the [MLE5OutputPort] class.
//
// # Methods
//
//   - [IMLE5OutputPort.Binder]
//   - [IMLE5OutputPort.SetBinder]
//   - [IMLE5OutputPort.BoundFeatureDirectly]
//   - [IMLE5OutputPort.FeatureDescription]
//   - [IMLE5OutputPort.FeatureValue]
//   - [IMLE5OutputPort.Name]
//   - [IMLE5OutputPort.OutputBackingWasDirectlyBound]
//   - [IMLE5OutputPort.PixelBufferPool]
//   - [IMLE5OutputPort.SetPixelBufferPool]
//   - [IMLE5OutputPort.PortHandle]
//   - [IMLE5OutputPort.PrepareWithOptionsError]
//   - [IMLE5OutputPort.Reset]
//   - [IMLE5OutputPort.ReusableForOutputBackingWillBindDirectly]
//   - [IMLE5OutputPort.InitWithPortHandleNameFeatureDescription]
//   - [IMLE5OutputPort.DebugDescription]
//   - [IMLE5OutputPort.Description]
//   - [IMLE5OutputPort.Hash]
//   - [IMLE5OutputPort.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPort
type IMLE5OutputPort interface {
	objectivec.IObject

	// Topic: Methods

	Binder() IMLE5OutputPortBinder
	SetBinder(value IMLE5OutputPortBinder)
	BoundFeatureDirectly() bool
	FeatureDescription() IMLFeatureDescription
	FeatureValue() IMLFeatureValue
	Name() string
	OutputBackingWasDirectlyBound() bool
	PixelBufferPool() IMLPixelBufferPool
	SetPixelBufferPool(value IMLPixelBufferPool)
	PortHandle() E5rtIOPortRef
	PrepareWithOptionsError(options objectivec.IObject) (bool, error)
	Reset()
	ReusableForOutputBackingWillBindDirectly(backing objectivec.IObject) (bool, bool)
	InitWithPortHandleNameFeatureDescription(handle E5rtIOPortRef, name objectivec.IObject, description objectivec.IObject) MLE5OutputPort
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (m MLE5OutputPort) Init() MLE5OutputPort {
	rv := objc.Send[MLE5OutputPort](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLE5OutputPort) Autorelease() MLE5OutputPort {
	rv := objc.Send[MLE5OutputPort](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLE5OutputPort creates a new MLE5OutputPort instance.
func NewMLE5OutputPort() MLE5OutputPort {
	class := getMLE5OutputPortClass()
	rv := objc.Send[MLE5OutputPort](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPort/initWithPortHandle:name:featureDescription:
func NewE5OutputPortWithPortHandleNameFeatureDescription(handle E5rtIOPortRef, name objectivec.IObject, description objectivec.IObject) MLE5OutputPort {
	instance := getMLE5OutputPortClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPortHandle:name:featureDescription:"), handle, name, description)
	return MLE5OutputPortFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPort/prepareWithOptions:error:
func (m MLE5OutputPort) PrepareWithOptionsError(options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("prepareWithOptions:error:"), options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("prepareWithOptions:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPort/reset
func (m MLE5OutputPort) Reset() {
	objc.Send[objc.ID](m.ID, objc.Sel("reset"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPort/reusableForOutputBacking:willBindDirectly:
func (m MLE5OutputPort) ReusableForOutputBackingWillBindDirectly(backing objectivec.IObject) (bool, bool) {
	var directly bool
	rv := objc.Send[bool](m.ID, objc.Sel("reusableForOutputBacking:willBindDirectly:"), backing, unsafe.Pointer(&directly))
	return directly, rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPort/initWithPortHandle:name:featureDescription:
func (m MLE5OutputPort) InitWithPortHandleNameFeatureDescription(handle E5rtIOPortRef, name objectivec.IObject, description objectivec.IObject) MLE5OutputPort {
	rv := objc.Send[MLE5OutputPort](m.ID, objc.Sel("initWithPortHandle:name:featureDescription:"), handle, name, description)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPort/binder
func (m MLE5OutputPort) Binder() IMLE5OutputPortBinder {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("binder"))
	return MLE5OutputPortBinderFromID(objc.ID(rv))
}
func (m MLE5OutputPort) SetBinder(value IMLE5OutputPortBinder) {
	objc.Send[struct{}](m.ID, objc.Sel("setBinder:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPort/boundFeatureDirectly
func (m MLE5OutputPort) BoundFeatureDirectly() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("boundFeatureDirectly"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPort/debugDescription
func (m MLE5OutputPort) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPort/description
func (m MLE5OutputPort) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPort/featureDescription
func (m MLE5OutputPort) FeatureDescription() IMLFeatureDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("featureDescription"))
	return MLFeatureDescriptionFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPort/featureValue
func (m MLE5OutputPort) FeatureValue() IMLFeatureValue {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("featureValue"))
	return MLFeatureValueFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPort/hash
func (m MLE5OutputPort) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPort/name
func (m MLE5OutputPort) Name() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPort/outputBackingWasDirectlyBound
func (m MLE5OutputPort) OutputBackingWasDirectlyBound() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("outputBackingWasDirectlyBound"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPort/pixelBufferPool
func (m MLE5OutputPort) PixelBufferPool() IMLPixelBufferPool {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("pixelBufferPool"))
	return MLPixelBufferPoolFromID(objc.ID(rv))
}
func (m MLE5OutputPort) SetPixelBufferPool(value IMLPixelBufferPool) {
	objc.Send[struct{}](m.ID, objc.Sel("setPixelBufferPool:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPort/portHandle
func (m MLE5OutputPort) PortHandle() E5rtIOPortRef {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("portHandle"))
	return E5rtIOPortRef(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPort/superclass
func (m MLE5OutputPort) Superclass() objc.Class {
	rv := objc.Send[objc.Class](m.ID, objc.Sel("superclass"))
	return rv
}
