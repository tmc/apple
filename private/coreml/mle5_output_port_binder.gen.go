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

// The class instance for the [MLE5OutputPortBinder] class.
var (
	_MLE5OutputPortBinderClass     MLE5OutputPortBinderClass
	_MLE5OutputPortBinderClassOnce sync.Once
)

func getMLE5OutputPortBinderClass() MLE5OutputPortBinderClass {
	_MLE5OutputPortBinderClassOnce.Do(func() {
		_MLE5OutputPortBinderClass = MLE5OutputPortBinderClass{class: objc.GetClass("MLE5OutputPortBinder")}
	})
	return _MLE5OutputPortBinderClass
}

// GetMLE5OutputPortBinderClass returns the class object for MLE5OutputPortBinder.
func GetMLE5OutputPortBinderClass() MLE5OutputPortBinderClass {
	return getMLE5OutputPortBinderClass()
}

type MLE5OutputPortBinderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLE5OutputPortBinderClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLE5OutputPortBinderClass) Alloc() MLE5OutputPortBinder {
	rv := objc.Send[MLE5OutputPortBinder](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLE5OutputPortBinder._copyOutputFromPortToOutputBackingFeatureDescriptionError]
//   - [MLE5OutputPortBinder._directModeForOutputBackingError]
//   - [MLE5OutputPortBinder._directlyBindOutputBackingError]
//   - [MLE5OutputPortBinder._makeFeatureValueAndReturnError]
//   - [MLE5OutputPortBinder._makeFeatureValueFromOutputBackingError]
//   - [MLE5OutputPortBinder._makeFeatureValueFromPortFeatureDescriptionError]
//   - [MLE5OutputPortBinder._outputModeForOutputBackingError]
//   - [MLE5OutputPortBinder._reusableForCopyBoundOutputBacking]
//   - [MLE5OutputPortBinder._reusableForDirectlyBoundOutputBacking]
//   - [MLE5OutputPortBinder.BindAndReturnError]
//   - [MLE5OutputPortBinder.BoundFeatureDirectly]
//   - [MLE5OutputPortBinder.SetBoundFeatureDirectly]
//   - [MLE5OutputPortBinder.FeatureDescription]
//   - [MLE5OutputPortBinder.FeatureValue]
//   - [MLE5OutputPortBinder.OutputBacking]
//   - [MLE5OutputPortBinder.SetOutputBacking]
//   - [MLE5OutputPortBinder.OutputBackingWasDirectlyBound]
//   - [MLE5OutputPortBinder.SetOutputBackingWasDirectlyBound]
//   - [MLE5OutputPortBinder.PixelBufferPool]
//   - [MLE5OutputPortBinder.SetPixelBufferPool]
//   - [MLE5OutputPortBinder.PortHandle]
//   - [MLE5OutputPortBinder.Reset]
//   - [MLE5OutputPortBinder.ReusableForOutputBackingWillBindDirectly]
//   - [MLE5OutputPortBinder.SerialQueue]
//   - [MLE5OutputPortBinder.InitWithPortFeatureDescription]
//   - [MLE5OutputPortBinder.DebugDescription]
//   - [MLE5OutputPortBinder.Description]
//   - [MLE5OutputPortBinder.Hash]
//   - [MLE5OutputPortBinder.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder
type MLE5OutputPortBinder struct {
	objectivec.Object
}

// MLE5OutputPortBinderFromID constructs a [MLE5OutputPortBinder] from an objc.ID.
func MLE5OutputPortBinderFromID(id objc.ID) MLE5OutputPortBinder {
	return MLE5OutputPortBinder{objectivec.Object{ID: id}}
}

// Ensure MLE5OutputPortBinder implements IMLE5OutputPortBinder.
var _ IMLE5OutputPortBinder = MLE5OutputPortBinder{}

// An interface definition for the [MLE5OutputPortBinder] class.
//
// # Methods
//
//   - [IMLE5OutputPortBinder._copyOutputFromPortToOutputBackingFeatureDescriptionError]
//   - [IMLE5OutputPortBinder._directModeForOutputBackingError]
//   - [IMLE5OutputPortBinder._directlyBindOutputBackingError]
//   - [IMLE5OutputPortBinder._makeFeatureValueAndReturnError]
//   - [IMLE5OutputPortBinder._makeFeatureValueFromOutputBackingError]
//   - [IMLE5OutputPortBinder._makeFeatureValueFromPortFeatureDescriptionError]
//   - [IMLE5OutputPortBinder._outputModeForOutputBackingError]
//   - [IMLE5OutputPortBinder._reusableForCopyBoundOutputBacking]
//   - [IMLE5OutputPortBinder._reusableForDirectlyBoundOutputBacking]
//   - [IMLE5OutputPortBinder.BindAndReturnError]
//   - [IMLE5OutputPortBinder.BoundFeatureDirectly]
//   - [IMLE5OutputPortBinder.SetBoundFeatureDirectly]
//   - [IMLE5OutputPortBinder.FeatureDescription]
//   - [IMLE5OutputPortBinder.FeatureValue]
//   - [IMLE5OutputPortBinder.OutputBacking]
//   - [IMLE5OutputPortBinder.SetOutputBacking]
//   - [IMLE5OutputPortBinder.OutputBackingWasDirectlyBound]
//   - [IMLE5OutputPortBinder.SetOutputBackingWasDirectlyBound]
//   - [IMLE5OutputPortBinder.PixelBufferPool]
//   - [IMLE5OutputPortBinder.SetPixelBufferPool]
//   - [IMLE5OutputPortBinder.PortHandle]
//   - [IMLE5OutputPortBinder.Reset]
//   - [IMLE5OutputPortBinder.ReusableForOutputBackingWillBindDirectly]
//   - [IMLE5OutputPortBinder.SerialQueue]
//   - [IMLE5OutputPortBinder.InitWithPortFeatureDescription]
//   - [IMLE5OutputPortBinder.DebugDescription]
//   - [IMLE5OutputPortBinder.Description]
//   - [IMLE5OutputPortBinder.Hash]
//   - [IMLE5OutputPortBinder.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder
type IMLE5OutputPortBinder interface {
	objectivec.IObject

	// Topic: Methods

	_copyOutputFromPortToOutputBackingFeatureDescriptionError(port objectivec.IObject, backing objectivec.IObject, description objectivec.IObject) (bool, error)
	_directModeForOutputBackingError(backing objectivec.IObject) (byte, error)
	_directlyBindOutputBackingError(backing objectivec.IObject) (bool, error)
	_makeFeatureValueAndReturnError() (objectivec.IObject, error)
	_makeFeatureValueFromOutputBackingError(backing objectivec.IObject) (objectivec.IObject, error)
	_makeFeatureValueFromPortFeatureDescriptionError(port objectivec.IObject, description objectivec.IObject) (objectivec.IObject, error)
	_outputModeForOutputBackingError(backing objectivec.IObject) (int64, error)
	_reusableForCopyBoundOutputBacking() bool
	_reusableForDirectlyBoundOutputBacking(backing objectivec.IObject) bool
	BindAndReturnError() (bool, error)
	BoundFeatureDirectly() bool
	SetBoundFeatureDirectly(value bool)
	FeatureDescription() IMLFeatureDescription
	FeatureValue() IMLFeatureValue
	OutputBacking() objectivec.IObject
	SetOutputBacking(value objectivec.IObject)
	OutputBackingWasDirectlyBound() bool
	SetOutputBackingWasDirectlyBound(value bool)
	PixelBufferPool() IMLPixelBufferPool
	SetPixelBufferPool(value IMLPixelBufferPool)
	PortHandle() objectivec.IObject
	Reset()
	ReusableForOutputBackingWillBindDirectly(backing objectivec.IObject) (bool, bool)
	SerialQueue() objectivec.Object
	InitWithPortFeatureDescription(port objectivec.IObject, description objectivec.IObject) MLE5OutputPortBinder
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (e MLE5OutputPortBinder) Init() MLE5OutputPortBinder {
	rv := objc.Send[MLE5OutputPortBinder](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e MLE5OutputPortBinder) Autorelease() MLE5OutputPortBinder {
	rv := objc.Send[MLE5OutputPortBinder](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLE5OutputPortBinder creates a new MLE5OutputPortBinder instance.
func NewMLE5OutputPortBinder() MLE5OutputPortBinder {
	class := getMLE5OutputPortBinderClass()
	rv := objc.Send[MLE5OutputPortBinder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/initWithPort:featureDescription:
func NewE5OutputPortBinderWithPortFeatureDescription(port objectivec.IObject, description objectivec.IObject) MLE5OutputPortBinder {
	instance := getMLE5OutputPortBinderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPort:featureDescription:"), port, description)
	return MLE5OutputPortBinderFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/_copyOutputFromPort:toOutputBacking:featureDescription:error:
func (e MLE5OutputPortBinder) _copyOutputFromPortToOutputBackingFeatureDescriptionError(port objectivec.IObject, backing objectivec.IObject, description objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("_copyOutputFromPort:toOutputBacking:featureDescription:error:"), port, backing, description, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_copyOutputFromPort:toOutputBacking:featureDescription:error: returned NO with nil NSError")
	}
	return rv, nil

}

// CopyOutputFromPortToOutputBackingFeatureDescriptionError is an exported wrapper for the private method _copyOutputFromPortToOutputBackingFeatureDescriptionError.
func (e MLE5OutputPortBinder) CopyOutputFromPortToOutputBackingFeatureDescriptionError(port objectivec.IObject, backing objectivec.IObject, description objectivec.IObject) (bool, error) {
	return e._copyOutputFromPortToOutputBackingFeatureDescriptionError(port, backing, description)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/_directModeForOutputBacking:error:
func (e MLE5OutputPortBinder) _directModeForOutputBackingError(backing objectivec.IObject) (byte, error) {
	var errorPtr objc.ID
	rv := objc.Send[byte](e.ID, objc.Sel("_directModeForOutputBacking:error:"), backing, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// DirectModeForOutputBackingError is an exported wrapper for the private method _directModeForOutputBackingError.
func (e MLE5OutputPortBinder) DirectModeForOutputBackingError(backing objectivec.IObject) (byte, error) {
	return e._directModeForOutputBackingError(backing)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/_directlyBindOutputBacking:error:
func (e MLE5OutputPortBinder) _directlyBindOutputBackingError(backing objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("_directlyBindOutputBacking:error:"), backing, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_directlyBindOutputBacking:error: returned NO with nil NSError")
	}
	return rv, nil

}

// DirectlyBindOutputBackingError is an exported wrapper for the private method _directlyBindOutputBackingError.
func (e MLE5OutputPortBinder) DirectlyBindOutputBackingError(backing objectivec.IObject) (bool, error) {
	return e._directlyBindOutputBackingError(backing)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/_makeFeatureValueAndReturnError:
func (e MLE5OutputPortBinder) _makeFeatureValueAndReturnError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("_makeFeatureValueAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// MakeFeatureValueAndReturnError is an exported wrapper for the private method _makeFeatureValueAndReturnError.
func (e MLE5OutputPortBinder) MakeFeatureValueAndReturnError() (objectivec.IObject, error) {
	return e._makeFeatureValueAndReturnError()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/_makeFeatureValueFromOutputBacking:error:
func (e MLE5OutputPortBinder) _makeFeatureValueFromOutputBackingError(backing objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("_makeFeatureValueFromOutputBacking:error:"), backing, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// MakeFeatureValueFromOutputBackingError is an exported wrapper for the private method _makeFeatureValueFromOutputBackingError.
func (e MLE5OutputPortBinder) MakeFeatureValueFromOutputBackingError(backing objectivec.IObject) (objectivec.IObject, error) {
	return e._makeFeatureValueFromOutputBackingError(backing)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/_makeFeatureValueFromPort:featureDescription:error:
func (e MLE5OutputPortBinder) _makeFeatureValueFromPortFeatureDescriptionError(port objectivec.IObject, description objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("_makeFeatureValueFromPort:featureDescription:error:"), port, description, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// MakeFeatureValueFromPortFeatureDescriptionError is an exported wrapper for the private method _makeFeatureValueFromPortFeatureDescriptionError.
func (e MLE5OutputPortBinder) MakeFeatureValueFromPortFeatureDescriptionError(port objectivec.IObject, description objectivec.IObject) (objectivec.IObject, error) {
	return e._makeFeatureValueFromPortFeatureDescriptionError(port, description)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/_outputModeForOutputBacking:error:
func (e MLE5OutputPortBinder) _outputModeForOutputBackingError(backing objectivec.IObject) (int64, error) {
	var errorPtr objc.ID
	rv := objc.Send[int64](e.ID, objc.Sel("_outputModeForOutputBacking:error:"), backing, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// OutputModeForOutputBackingError is an exported wrapper for the private method _outputModeForOutputBackingError.
func (e MLE5OutputPortBinder) OutputModeForOutputBackingError(backing objectivec.IObject) (int64, error) {
	return e._outputModeForOutputBackingError(backing)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/_reusableForCopyBoundOutputBacking
func (e MLE5OutputPortBinder) _reusableForCopyBoundOutputBacking() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("_reusableForCopyBoundOutputBacking"))
	return rv
}

// ReusableForCopyBoundOutputBacking is an exported wrapper for the private method _reusableForCopyBoundOutputBacking.
func (e MLE5OutputPortBinder) ReusableForCopyBoundOutputBacking() bool {
	return e._reusableForCopyBoundOutputBacking()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/_reusableForDirectlyBoundOutputBacking:
func (e MLE5OutputPortBinder) _reusableForDirectlyBoundOutputBacking(backing objectivec.IObject) bool {
	rv := objc.Send[bool](e.ID, objc.Sel("_reusableForDirectlyBoundOutputBacking:"), backing)
	return rv
}

// ReusableForDirectlyBoundOutputBacking is an exported wrapper for the private method _reusableForDirectlyBoundOutputBacking.
func (e MLE5OutputPortBinder) ReusableForDirectlyBoundOutputBacking(backing objectivec.IObject) bool {
	return e._reusableForDirectlyBoundOutputBacking(backing)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/bindAndReturnError:
func (e MLE5OutputPortBinder) BindAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("bindAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("bindAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/reset
func (e MLE5OutputPortBinder) Reset() {
	objc.Send[objc.ID](e.ID, objc.Sel("reset"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/reusableForOutputBacking:willBindDirectly:
func (e MLE5OutputPortBinder) ReusableForOutputBackingWillBindDirectly(backing objectivec.IObject) (bool, bool) {
	var directly bool
	rv := objc.Send[bool](e.ID, objc.Sel("reusableForOutputBacking:willBindDirectly:"), backing, unsafe.Pointer(&directly))
	return directly, rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/initWithPort:featureDescription:
func (e MLE5OutputPortBinder) InitWithPortFeatureDescription(port objectivec.IObject, description objectivec.IObject) MLE5OutputPortBinder {
	rv := objc.Send[MLE5OutputPortBinder](e.ID, objc.Sel("initWithPort:featureDescription:"), port, description)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/boundFeatureDirectly
func (e MLE5OutputPortBinder) BoundFeatureDirectly() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("boundFeatureDirectly"))
	return rv
}
func (e MLE5OutputPortBinder) SetBoundFeatureDirectly(value bool) {
	objc.Send[struct{}](e.ID, objc.Sel("setBoundFeatureDirectly:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/debugDescription
func (e MLE5OutputPortBinder) DebugDescription() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/description
func (e MLE5OutputPortBinder) Description() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/featureDescription
func (e MLE5OutputPortBinder) FeatureDescription() IMLFeatureDescription {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("featureDescription"))
	return MLFeatureDescriptionFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/featureValue
func (e MLE5OutputPortBinder) FeatureValue() IMLFeatureValue {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("featureValue"))
	return MLFeatureValueFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/hash
func (e MLE5OutputPortBinder) Hash() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/outputBacking
func (e MLE5OutputPortBinder) OutputBacking() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("outputBacking"))
	return objectivec.Object{ID: rv}
}
func (e MLE5OutputPortBinder) SetOutputBacking(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setOutputBacking:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/outputBackingWasDirectlyBound
func (e MLE5OutputPortBinder) OutputBackingWasDirectlyBound() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("outputBackingWasDirectlyBound"))
	return rv
}
func (e MLE5OutputPortBinder) SetOutputBackingWasDirectlyBound(value bool) {
	objc.Send[struct{}](e.ID, objc.Sel("setOutputBackingWasDirectlyBound:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/pixelBufferPool
func (e MLE5OutputPortBinder) PixelBufferPool() IMLPixelBufferPool {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("pixelBufferPool"))
	return MLPixelBufferPoolFromID(objc.ID(rv))
}
func (e MLE5OutputPortBinder) SetPixelBufferPool(value IMLPixelBufferPool) {
	objc.Send[struct{}](e.ID, objc.Sel("setPixelBufferPool:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/portHandle
func (e MLE5OutputPortBinder) PortHandle() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("portHandle"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/serialQueue
func (e MLE5OutputPortBinder) SerialQueue() objectivec.Object {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("serialQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5OutputPortBinder/superclass
func (e MLE5OutputPortBinder) Superclass() objc.Class {
	rv := objc.Send[objc.Class](e.ID, objc.Sel("superclass"))
	return rv
}
