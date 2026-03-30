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

// The class instance for the [InternalCustomTileLike] class.
var (
	_InternalCustomTileLikeClass     InternalCustomTileLikeClass
	_InternalCustomTileLikeClassOnce sync.Once
)

func getInternalCustomTileLikeClass() InternalCustomTileLikeClass {
	_InternalCustomTileLikeClassOnce.Do(func() {
		_InternalCustomTileLikeClass = InternalCustomTileLikeClass{class: objc.GetClass("InternalCustomTileLike")}
	})
	return _InternalCustomTileLikeClass
}

// GetInternalCustomTileLikeClass returns the class object for InternalCustomTileLike.
func GetInternalCustomTileLikeClass() InternalCustomTileLikeClass {
	return getInternalCustomTileLikeClass()
}

type InternalCustomTileLikeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic InternalCustomTileLikeClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic InternalCustomTileLikeClass) Alloc() InternalCustomTileLike {
	rv := objc.Send[InternalCustomTileLike](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [InternalCustomTileLike.EvaluateOnCPUWithInputsOutputsError]
//   - [InternalCustomTileLike.InputRank]
//   - [InternalCustomTileLike.InputShape]
//   - [InternalCustomTileLike.Multiples]
//   - [InternalCustomTileLike.OutputShape]
//   - [InternalCustomTileLike.OutputShapesForInputShapesError]
//   - [InternalCustomTileLike.SetWeightDataError]
//   - [InternalCustomTileLike.InitWithParameterDictionaryError]
//
// See: https://developer.apple.com/documentation/CoreML/InternalCustomTileLike
type InternalCustomTileLike struct {
	objectivec.Object
}

// InternalCustomTileLikeFromID constructs a [InternalCustomTileLike] from an objc.ID.
func InternalCustomTileLikeFromID(id objc.ID) InternalCustomTileLike {
	return InternalCustomTileLike{objectivec.Object{ID: id}}
}

// Ensure InternalCustomTileLike implements IInternalCustomTileLike.
var _ IInternalCustomTileLike = InternalCustomTileLike{}

// An interface definition for the [InternalCustomTileLike] class.
//
// # Methods
//
//   - [IInternalCustomTileLike.EvaluateOnCPUWithInputsOutputsError]
//   - [IInternalCustomTileLike.InputRank]
//   - [IInternalCustomTileLike.InputShape]
//   - [IInternalCustomTileLike.Multiples]
//   - [IInternalCustomTileLike.OutputShape]
//   - [IInternalCustomTileLike.OutputShapesForInputShapesError]
//   - [IInternalCustomTileLike.SetWeightDataError]
//   - [IInternalCustomTileLike.InitWithParameterDictionaryError]
//
// See: https://developer.apple.com/documentation/CoreML/InternalCustomTileLike
type IInternalCustomTileLike interface {
	objectivec.IObject

	// Topic: Methods

	EvaluateOnCPUWithInputsOutputsError(inputs objectivec.IObject, outputs objectivec.IObject) (bool, error)
	InputRank() uint64
	InputShape() objectivec.IObject
	Multiples() objectivec.IObject
	OutputShape() objectivec.IObject
	OutputShapesForInputShapesError(shapes objectivec.IObject) (objectivec.IObject, error)
	SetWeightDataError(data objectivec.IObject) (bool, error)
	InitWithParameterDictionaryError(dictionary objectivec.IObject) (InternalCustomTileLike, error)
}

// Init initializes the instance.
func (i InternalCustomTileLike) Init() InternalCustomTileLike {
	rv := objc.Send[InternalCustomTileLike](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i InternalCustomTileLike) Autorelease() InternalCustomTileLike {
	rv := objc.Send[InternalCustomTileLike](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewInternalCustomTileLike creates a new InternalCustomTileLike instance.
func NewInternalCustomTileLike() InternalCustomTileLike {
	class := getInternalCustomTileLikeClass()
	rv := objc.Send[InternalCustomTileLike](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/InternalCustomTileLike/initWithParameterDictionary:error:
func NewInternalCustomTileLikeWithParameterDictionaryError(dictionary objectivec.IObject) (InternalCustomTileLike, error) {
	var errorPtr objc.ID
	instance := getInternalCustomTileLikeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameterDictionary:error:"), dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return InternalCustomTileLike{}, foundation.NSErrorFrom(errorPtr)
	}
	return InternalCustomTileLikeFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/InternalCustomTileLike/evaluateOnCPUWithInputs:outputs:error:
func (i InternalCustomTileLike) EvaluateOnCPUWithInputsOutputsError(inputs objectivec.IObject, outputs objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](i.ID, objc.Sel("evaluateOnCPUWithInputs:outputs:error:"), inputs, outputs, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("evaluateOnCPUWithInputs:outputs:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/InternalCustomTileLike/outputShapesForInputShapes:error:
func (i InternalCustomTileLike) OutputShapesForInputShapesError(shapes objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](i.ID, objc.Sel("outputShapesForInputShapes:error:"), shapes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/InternalCustomTileLike/setWeightData:error:
func (i InternalCustomTileLike) SetWeightDataError(data objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](i.ID, objc.Sel("setWeightData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setWeightData:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/InternalCustomTileLike/initWithParameterDictionary:error:
func (i InternalCustomTileLike) InitWithParameterDictionaryError(dictionary objectivec.IObject) (InternalCustomTileLike, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](i.ID, objc.Sel("initWithParameterDictionary:error:"), dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return InternalCustomTileLike{}, foundation.NSErrorFrom(errorPtr)
	}
	return InternalCustomTileLikeFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/InternalCustomTileLike/inputRank
func (i InternalCustomTileLike) InputRank() uint64 {
	rv := objc.Send[uint64](i.ID, objc.Sel("inputRank"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/InternalCustomTileLike/inputShape
func (i InternalCustomTileLike) InputShape() objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("inputShape"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/InternalCustomTileLike/multiples
func (i InternalCustomTileLike) Multiples() objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("multiples"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/InternalCustomTileLike/outputShape
func (i InternalCustomTileLike) OutputShape() objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("outputShape"))
	return objectivec.Object{ID: rv}
}
