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

// The class instance for the [MLE5InputPort] class.
var (
	_MLE5InputPortClass     MLE5InputPortClass
	_MLE5InputPortClassOnce sync.Once
)

func getMLE5InputPortClass() MLE5InputPortClass {
	_MLE5InputPortClassOnce.Do(func() {
		_MLE5InputPortClass = MLE5InputPortClass{class: objc.GetClass("MLE5InputPort")}
	})
	return _MLE5InputPortClass
}

// GetMLE5InputPortClass returns the class object for MLE5InputPort.
func GetMLE5InputPortClass() MLE5InputPortClass {
	return getMLE5InputPortClass()
}

type MLE5InputPortClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLE5InputPortClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLE5InputPortClass) Alloc() MLE5InputPort {
	rv := objc.Send[MLE5InputPort](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLE5InputPort.Binder]
//   - [MLE5InputPort.SetBinder]
//   - [MLE5InputPort.BoundFeatureDirectly]
//   - [MLE5InputPort.CopyFeatureValueError]
//   - [MLE5InputPort.Name]
//   - [MLE5InputPort.PixelBufferPool]
//   - [MLE5InputPort.SetPixelBufferPool]
//   - [MLE5InputPort.PortHandle]
//   - [MLE5InputPort.PrepareForFeatureValueError]
//   - [MLE5InputPort.Reset]
//   - [MLE5InputPort.ReusableForFeatureValueWillBindDirectly]
//   - [MLE5InputPort.InitWithPortHandleNameFeatureDescription]
//   - [MLE5InputPort.DebugDescription]
//   - [MLE5InputPort.Description]
//   - [MLE5InputPort.Hash]
//   - [MLE5InputPort.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLE5InputPort
type MLE5InputPort struct {
	objectivec.Object
}

// MLE5InputPortFromID constructs a [MLE5InputPort] from an objc.ID.
func MLE5InputPortFromID(id objc.ID) MLE5InputPort {
	return MLE5InputPort{objectivec.Object{ID: id}}
}

// Ensure MLE5InputPort implements IMLE5InputPort.
var _ IMLE5InputPort = MLE5InputPort{}

// An interface definition for the [MLE5InputPort] class.
//
// # Methods
//
//   - [IMLE5InputPort.Binder]
//   - [IMLE5InputPort.SetBinder]
//   - [IMLE5InputPort.BoundFeatureDirectly]
//   - [IMLE5InputPort.CopyFeatureValueError]
//   - [IMLE5InputPort.Name]
//   - [IMLE5InputPort.PixelBufferPool]
//   - [IMLE5InputPort.SetPixelBufferPool]
//   - [IMLE5InputPort.PortHandle]
//   - [IMLE5InputPort.PrepareForFeatureValueError]
//   - [IMLE5InputPort.Reset]
//   - [IMLE5InputPort.ReusableForFeatureValueWillBindDirectly]
//   - [IMLE5InputPort.InitWithPortHandleNameFeatureDescription]
//   - [IMLE5InputPort.DebugDescription]
//   - [IMLE5InputPort.Description]
//   - [IMLE5InputPort.Hash]
//   - [IMLE5InputPort.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLE5InputPort
type IMLE5InputPort interface {
	objectivec.IObject

	// Topic: Methods

	Binder() IMLE5InputPortBinder
	SetBinder(value IMLE5InputPortBinder)
	BoundFeatureDirectly() bool
	CopyFeatureValueError(value objectivec.IObject) (bool, error)
	Name() string
	PixelBufferPool() IMLPixelBufferPool
	SetPixelBufferPool(value IMLPixelBufferPool)
	PortHandle() objectivec.IObject
	PrepareForFeatureValueError(value objectivec.IObject) (bool, error)
	Reset()
	ReusableForFeatureValueWillBindDirectly(value objectivec.IObject) (bool, bool)
	InitWithPortHandleNameFeatureDescription(handle objectivec.IObject, name objectivec.IObject, description objectivec.IObject) MLE5InputPort
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (e MLE5InputPort) Init() MLE5InputPort {
	rv := objc.Send[MLE5InputPort](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e MLE5InputPort) Autorelease() MLE5InputPort {
	rv := objc.Send[MLE5InputPort](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLE5InputPort creates a new MLE5InputPort instance.
func NewMLE5InputPort() MLE5InputPort {
	class := getMLE5InputPortClass()
	rv := objc.Send[MLE5InputPort](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPort/initWithPortHandle:name:featureDescription:
func NewE5InputPortWithPortHandleNameFeatureDescription(handle objectivec.IObject, name objectivec.IObject, description objectivec.IObject) MLE5InputPort {
	instance := getMLE5InputPortClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPortHandle:name:featureDescription:"), handle, name, description)
	return MLE5InputPortFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPort/copyFeatureValue:error:
func (e MLE5InputPort) CopyFeatureValueError(value objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("copyFeatureValue:error:"), value, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("copyFeatureValue:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPort/prepareForFeatureValue:error:
func (e MLE5InputPort) PrepareForFeatureValueError(value objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("prepareForFeatureValue:error:"), value, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("prepareForFeatureValue:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPort/reset
func (e MLE5InputPort) Reset() {
	objc.Send[objc.ID](e.ID, objc.Sel("reset"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPort/reusableForFeatureValue:willBindDirectly:
func (e MLE5InputPort) ReusableForFeatureValueWillBindDirectly(value objectivec.IObject) (bool, bool) {
	var directly bool
	rv := objc.Send[bool](e.ID, objc.Sel("reusableForFeatureValue:willBindDirectly:"), value, unsafe.Pointer(&directly))
	return directly, rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPort/initWithPortHandle:name:featureDescription:
func (e MLE5InputPort) InitWithPortHandleNameFeatureDescription(handle objectivec.IObject, name objectivec.IObject, description objectivec.IObject) MLE5InputPort {
	rv := objc.Send[MLE5InputPort](e.ID, objc.Sel("initWithPortHandle:name:featureDescription:"), handle, name, description)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPort/binder
func (e MLE5InputPort) Binder() IMLE5InputPortBinder {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("binder"))
	return MLE5InputPortBinderFromID(objc.ID(rv))
}
func (e MLE5InputPort) SetBinder(value IMLE5InputPortBinder) {
	objc.Send[struct{}](e.ID, objc.Sel("setBinder:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPort/boundFeatureDirectly
func (e MLE5InputPort) BoundFeatureDirectly() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("boundFeatureDirectly"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPort/debugDescription
func (e MLE5InputPort) DebugDescription() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPort/description
func (e MLE5InputPort) Description() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPort/hash
func (e MLE5InputPort) Hash() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPort/name
func (e MLE5InputPort) Name() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPort/pixelBufferPool
func (e MLE5InputPort) PixelBufferPool() IMLPixelBufferPool {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("pixelBufferPool"))
	return MLPixelBufferPoolFromID(objc.ID(rv))
}
func (e MLE5InputPort) SetPixelBufferPool(value IMLPixelBufferPool) {
	objc.Send[struct{}](e.ID, objc.Sel("setPixelBufferPool:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPort/portHandle
func (e MLE5InputPort) PortHandle() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("portHandle"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPort/superclass
func (e MLE5InputPort) Superclass() objc.Class {
	rv := objc.Send[objc.Class](e.ID, objc.Sel("superclass"))
	return rv
}
