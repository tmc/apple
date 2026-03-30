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

// The class instance for the [InternalCustomGatherTree] class.
var (
	_InternalCustomGatherTreeClass     InternalCustomGatherTreeClass
	_InternalCustomGatherTreeClassOnce sync.Once
)

func getInternalCustomGatherTreeClass() InternalCustomGatherTreeClass {
	_InternalCustomGatherTreeClassOnce.Do(func() {
		_InternalCustomGatherTreeClass = InternalCustomGatherTreeClass{class: objc.GetClass("InternalCustomGatherTree")}
	})
	return _InternalCustomGatherTreeClass
}

// GetInternalCustomGatherTreeClass returns the class object for InternalCustomGatherTree.
func GetInternalCustomGatherTreeClass() InternalCustomGatherTreeClass {
	return getInternalCustomGatherTreeClass()
}

type InternalCustomGatherTreeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic InternalCustomGatherTreeClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic InternalCustomGatherTreeClass) Alloc() InternalCustomGatherTree {
	rv := objc.Send[InternalCustomGatherTree](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [InternalCustomGatherTree.EvaluateOnCPUWithInputsOutputsError]
//   - [InternalCustomGatherTree.OutputShapesForInputShapesError]
//   - [InternalCustomGatherTree.SetWeightDataError]
//   - [InternalCustomGatherTree.Shape]
//   - [InternalCustomGatherTree.InitWithParameterDictionaryError]
//
// See: https://developer.apple.com/documentation/CoreML/InternalCustomGatherTree
type InternalCustomGatherTree struct {
	objectivec.Object
}

// InternalCustomGatherTreeFromID constructs a [InternalCustomGatherTree] from an objc.ID.
func InternalCustomGatherTreeFromID(id objc.ID) InternalCustomGatherTree {
	return InternalCustomGatherTree{objectivec.Object{ID: id}}
}

// Ensure InternalCustomGatherTree implements IInternalCustomGatherTree.
var _ IInternalCustomGatherTree = InternalCustomGatherTree{}

// An interface definition for the [InternalCustomGatherTree] class.
//
// # Methods
//
//   - [IInternalCustomGatherTree.EvaluateOnCPUWithInputsOutputsError]
//   - [IInternalCustomGatherTree.OutputShapesForInputShapesError]
//   - [IInternalCustomGatherTree.SetWeightDataError]
//   - [IInternalCustomGatherTree.Shape]
//   - [IInternalCustomGatherTree.InitWithParameterDictionaryError]
//
// See: https://developer.apple.com/documentation/CoreML/InternalCustomGatherTree
type IInternalCustomGatherTree interface {
	objectivec.IObject

	// Topic: Methods

	EvaluateOnCPUWithInputsOutputsError(inputs objectivec.IObject, outputs objectivec.IObject) (bool, error)
	OutputShapesForInputShapesError(shapes objectivec.IObject) (objectivec.IObject, error)
	SetWeightDataError(data objectivec.IObject) (bool, error)
	Shape() objectivec.IObject
	InitWithParameterDictionaryError(dictionary objectivec.IObject) (InternalCustomGatherTree, error)
}

// Init initializes the instance.
func (i InternalCustomGatherTree) Init() InternalCustomGatherTree {
	rv := objc.Send[InternalCustomGatherTree](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i InternalCustomGatherTree) Autorelease() InternalCustomGatherTree {
	rv := objc.Send[InternalCustomGatherTree](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewInternalCustomGatherTree creates a new InternalCustomGatherTree instance.
func NewInternalCustomGatherTree() InternalCustomGatherTree {
	class := getInternalCustomGatherTreeClass()
	rv := objc.Send[InternalCustomGatherTree](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/InternalCustomGatherTree/initWithParameterDictionary:error:
func NewInternalCustomGatherTreeWithParameterDictionaryError(dictionary objectivec.IObject) (InternalCustomGatherTree, error) {
	var errorPtr objc.ID
	instance := getInternalCustomGatherTreeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameterDictionary:error:"), dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return InternalCustomGatherTree{}, foundation.NSErrorFrom(errorPtr)
	}
	return InternalCustomGatherTreeFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/InternalCustomGatherTree/evaluateOnCPUWithInputs:outputs:error:
func (i InternalCustomGatherTree) EvaluateOnCPUWithInputsOutputsError(inputs objectivec.IObject, outputs objectivec.IObject) (bool, error) {
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

// See: https://developer.apple.com/documentation/CoreML/InternalCustomGatherTree/outputShapesForInputShapes:error:
func (i InternalCustomGatherTree) OutputShapesForInputShapesError(shapes objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](i.ID, objc.Sel("outputShapesForInputShapes:error:"), shapes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/InternalCustomGatherTree/setWeightData:error:
func (i InternalCustomGatherTree) SetWeightDataError(data objectivec.IObject) (bool, error) {
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

// See: https://developer.apple.com/documentation/CoreML/InternalCustomGatherTree/initWithParameterDictionary:error:
func (i InternalCustomGatherTree) InitWithParameterDictionaryError(dictionary objectivec.IObject) (InternalCustomGatherTree, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](i.ID, objc.Sel("initWithParameterDictionary:error:"), dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return InternalCustomGatherTree{}, foundation.NSErrorFrom(errorPtr)
	}
	return InternalCustomGatherTreeFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/InternalCustomGatherTree/shape
func (i InternalCustomGatherTree) Shape() objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("shape"))
	return objectivec.Object{ID: rv}
}
