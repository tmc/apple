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

// The class instance for the [MLE5InputPortBinder] class.
var (
	_MLE5InputPortBinderClass     MLE5InputPortBinderClass
	_MLE5InputPortBinderClassOnce sync.Once
)

func getMLE5InputPortBinderClass() MLE5InputPortBinderClass {
	_MLE5InputPortBinderClassOnce.Do(func() {
		_MLE5InputPortBinderClass = MLE5InputPortBinderClass{class: objc.GetClass("MLE5InputPortBinder")}
	})
	return _MLE5InputPortBinderClass
}

// GetMLE5InputPortBinderClass returns the class object for MLE5InputPortBinder.
func GetMLE5InputPortBinderClass() MLE5InputPortBinderClass {
	return getMLE5InputPortBinderClass()
}

type MLE5InputPortBinderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLE5InputPortBinderClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLE5InputPortBinderClass) Alloc() MLE5InputPortBinder {
	rv := objc.Send[MLE5InputPortBinder](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLE5InputPortBinder._reusableForFeatureValueDirectMode]
//   - [MLE5InputPortBinder.BindMemoryObjectForFeatureValueError]
//   - [MLE5InputPortBinder.BindingMode]
//   - [MLE5InputPortBinder.SetBindingMode]
//   - [MLE5InputPortBinder.CopyFeatureValueError]
//   - [MLE5InputPortBinder.DirectlyBoundFeatureValue]
//   - [MLE5InputPortBinder.SetDirectlyBoundFeatureValue]
//   - [MLE5InputPortBinder.FeatureDescription]
//   - [MLE5InputPortBinder.PixelBufferPool]
//   - [MLE5InputPortBinder.SetPixelBufferPool]
//   - [MLE5InputPortBinder.PortHandle]
//   - [MLE5InputPortBinder.Reset]
//   - [MLE5InputPortBinder.ReusableForFeatureValueWillBindDirectly]
//   - [MLE5InputPortBinder.InitWithPortFeatureDescription]
//   - [MLE5InputPortBinder.DebugDescription]
//   - [MLE5InputPortBinder.Description]
//   - [MLE5InputPortBinder.Hash]
//   - [MLE5InputPortBinder.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLE5InputPortBinder
type MLE5InputPortBinder struct {
	objectivec.Object
}

// MLE5InputPortBinderFromID constructs a [MLE5InputPortBinder] from an objc.ID.
func MLE5InputPortBinderFromID(id objc.ID) MLE5InputPortBinder {
	return MLE5InputPortBinder{objectivec.Object{ID: id}}
}

// Ensure MLE5InputPortBinder implements IMLE5InputPortBinder.
var _ IMLE5InputPortBinder = MLE5InputPortBinder{}

// An interface definition for the [MLE5InputPortBinder] class.
//
// # Methods
//
//   - [IMLE5InputPortBinder._reusableForFeatureValueDirectMode]
//   - [IMLE5InputPortBinder.BindMemoryObjectForFeatureValueError]
//   - [IMLE5InputPortBinder.BindingMode]
//   - [IMLE5InputPortBinder.SetBindingMode]
//   - [IMLE5InputPortBinder.CopyFeatureValueError]
//   - [IMLE5InputPortBinder.DirectlyBoundFeatureValue]
//   - [IMLE5InputPortBinder.SetDirectlyBoundFeatureValue]
//   - [IMLE5InputPortBinder.FeatureDescription]
//   - [IMLE5InputPortBinder.PixelBufferPool]
//   - [IMLE5InputPortBinder.SetPixelBufferPool]
//   - [IMLE5InputPortBinder.PortHandle]
//   - [IMLE5InputPortBinder.Reset]
//   - [IMLE5InputPortBinder.ReusableForFeatureValueWillBindDirectly]
//   - [IMLE5InputPortBinder.InitWithPortFeatureDescription]
//   - [IMLE5InputPortBinder.DebugDescription]
//   - [IMLE5InputPortBinder.Description]
//   - [IMLE5InputPortBinder.Hash]
//   - [IMLE5InputPortBinder.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLE5InputPortBinder
type IMLE5InputPortBinder interface {
	objectivec.IObject

	// Topic: Methods

	_reusableForFeatureValueDirectMode(value objectivec.IObject, mode byte) bool
	BindMemoryObjectForFeatureValueError(value objectivec.IObject) (bool, error)
	BindingMode() byte
	SetBindingMode(value byte)
	CopyFeatureValueError(value objectivec.IObject) (bool, error)
	DirectlyBoundFeatureValue() IMLFeatureValue
	SetDirectlyBoundFeatureValue(value IMLFeatureValue)
	FeatureDescription() IMLFeatureDescription
	PixelBufferPool() IMLPixelBufferPool
	SetPixelBufferPool(value IMLPixelBufferPool)
	PortHandle() objectivec.IObject
	Reset()
	ReusableForFeatureValueWillBindDirectly(value objectivec.IObject) (bool, bool)
	InitWithPortFeatureDescription(port objectivec.IObject, description objectivec.IObject) MLE5InputPortBinder
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (e MLE5InputPortBinder) Init() MLE5InputPortBinder {
	rv := objc.Send[MLE5InputPortBinder](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e MLE5InputPortBinder) Autorelease() MLE5InputPortBinder {
	rv := objc.Send[MLE5InputPortBinder](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLE5InputPortBinder creates a new MLE5InputPortBinder instance.
func NewMLE5InputPortBinder() MLE5InputPortBinder {
	class := getMLE5InputPortBinderClass()
	rv := objc.Send[MLE5InputPortBinder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPortBinder/initWithPort:featureDescription:
func NewE5InputPortBinderWithPortFeatureDescription(port objectivec.IObject, description objectivec.IObject) MLE5InputPortBinder {
	instance := getMLE5InputPortBinderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPort:featureDescription:"), port, description)
	return MLE5InputPortBinderFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPortBinder/_reusableForFeatureValue:directMode:
func (e MLE5InputPortBinder) _reusableForFeatureValueDirectMode(value objectivec.IObject, mode byte) bool {
	rv := objc.Send[bool](e.ID, objc.Sel("_reusableForFeatureValue:directMode:"), value, mode)
	return rv
}

// ReusableForFeatureValueDirectMode is an exported wrapper for the private method _reusableForFeatureValueDirectMode.
func (e MLE5InputPortBinder) ReusableForFeatureValueDirectMode(value objectivec.IObject, mode byte) bool {
	return e._reusableForFeatureValueDirectMode(value, mode)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPortBinder/bindMemoryObjectForFeatureValue:error:
func (e MLE5InputPortBinder) BindMemoryObjectForFeatureValueError(value objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("bindMemoryObjectForFeatureValue:error:"), value, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("bindMemoryObjectForFeatureValue:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPortBinder/copyFeatureValue:error:
func (e MLE5InputPortBinder) CopyFeatureValueError(value objectivec.IObject) (bool, error) {
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

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPortBinder/reset
func (e MLE5InputPortBinder) Reset() {
	objc.Send[objc.ID](e.ID, objc.Sel("reset"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPortBinder/reusableForFeatureValue:willBindDirectly:
func (e MLE5InputPortBinder) ReusableForFeatureValueWillBindDirectly(value objectivec.IObject) (bool, bool) {
	var directly bool
	rv := objc.Send[bool](e.ID, objc.Sel("reusableForFeatureValue:willBindDirectly:"), value, unsafe.Pointer(&directly))
	return directly, rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPortBinder/initWithPort:featureDescription:
func (e MLE5InputPortBinder) InitWithPortFeatureDescription(port objectivec.IObject, description objectivec.IObject) MLE5InputPortBinder {
	rv := objc.Send[MLE5InputPortBinder](e.ID, objc.Sel("initWithPort:featureDescription:"), port, description)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPortBinder/bindingMode
func (e MLE5InputPortBinder) BindingMode() byte {
	rv := objc.Send[byte](e.ID, objc.Sel("bindingMode"))
	return rv
}
func (e MLE5InputPortBinder) SetBindingMode(value byte) {
	objc.Send[struct{}](e.ID, objc.Sel("setBindingMode:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPortBinder/debugDescription
func (e MLE5InputPortBinder) DebugDescription() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPortBinder/description
func (e MLE5InputPortBinder) Description() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPortBinder/directlyBoundFeatureValue
func (e MLE5InputPortBinder) DirectlyBoundFeatureValue() IMLFeatureValue {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("directlyBoundFeatureValue"))
	return MLFeatureValueFromID(objc.ID(rv))
}
func (e MLE5InputPortBinder) SetDirectlyBoundFeatureValue(value IMLFeatureValue) {
	objc.Send[struct{}](e.ID, objc.Sel("setDirectlyBoundFeatureValue:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPortBinder/featureDescription
func (e MLE5InputPortBinder) FeatureDescription() IMLFeatureDescription {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("featureDescription"))
	return MLFeatureDescriptionFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPortBinder/hash
func (e MLE5InputPortBinder) Hash() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPortBinder/pixelBufferPool
func (e MLE5InputPortBinder) PixelBufferPool() IMLPixelBufferPool {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("pixelBufferPool"))
	return MLPixelBufferPoolFromID(objc.ID(rv))
}
func (e MLE5InputPortBinder) SetPixelBufferPool(value IMLPixelBufferPool) {
	objc.Send[struct{}](e.ID, objc.Sel("setPixelBufferPool:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPortBinder/portHandle
func (e MLE5InputPortBinder) PortHandle() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("portHandle"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLE5InputPortBinder/superclass
func (e MLE5InputPortBinder) Superclass() objc.Class {
	rv := objc.Send[objc.Class](e.ID, objc.Sel("superclass"))
	return rv
}
