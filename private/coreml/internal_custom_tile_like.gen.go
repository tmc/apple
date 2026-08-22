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
	rv := objc.SendIfResponds[InternalCustomTileLike](objc.ID(ic.class), objc.Sel("alloc"))
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
type IInternalCustomTileLike interface {
	objectivec.IObject

	// Topic: Methods

	EvaluateOnCPUWithInputsOutputsError(inputs objectivec.IObject, outputs objectivec.IObject) (bool, error)
	InputRank() uint64
	InputShape() unsafe.Pointer
	Multiples() unsafe.Pointer
	OutputShape() unsafe.Pointer
	OutputShapesForInputShapesError(shapes objectivec.IObject) (objectivec.IObject, error)
	SetWeightDataError(data objectivec.IObject) (bool, error)
	InitWithParameterDictionaryError(dictionary objectivec.IObject) (InternalCustomTileLike, error)
}

// Init initializes the instance.
func (i InternalCustomTileLike) Init() InternalCustomTileLike {
	rv := objc.SendIfResponds[InternalCustomTileLike](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i InternalCustomTileLike) Autorelease() InternalCustomTileLike {
	rv := objc.SendIfResponds[InternalCustomTileLike](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewInternalCustomTileLike creates a new InternalCustomTileLike instance.
func NewInternalCustomTileLike() InternalCustomTileLike {
	class := getInternalCustomTileLikeClass()
	rv := objc.SendIfResponds[InternalCustomTileLike](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewInternalCustomTileLikeWithParameterDictionaryError(dictionary objectivec.IObject) (InternalCustomTileLike, error) {
	var errorPtr objc.ID
	instance := getInternalCustomTileLikeClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParameterDictionary:error:"), dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return InternalCustomTileLike{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return InternalCustomTileLike{}, objc.ErrInitFailed
	}
	return InternalCustomTileLikeFromID(rv), nil
}

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
func (i InternalCustomTileLike) OutputShapesForInputShapesError(shapes objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](i.ID, objc.Sel("outputShapesForInputShapes:error:"), shapes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
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
func (i InternalCustomTileLike) InitWithParameterDictionaryError(dictionary objectivec.IObject) (InternalCustomTileLike, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](i.ID, objc.Sel("initWithParameterDictionary:error:"), dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return InternalCustomTileLike{}, foundation.NSErrorFrom(errorPtr)
	}
	return InternalCustomTileLikeFromID(rv), nil

}

func (i InternalCustomTileLike) InputRank() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("inputRank"))
	return rv
}
func (i InternalCustomTileLike) InputShape() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("inputShape"))
	return rv
}
func (i InternalCustomTileLike) Multiples() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("multiples"))
	return rv
}
func (i InternalCustomTileLike) OutputShape() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("outputShape"))
	return rv
}
