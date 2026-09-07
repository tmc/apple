// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANETensorDebugHelper] class.
var (
	_ANETensorDebugHelperClass     ANETensorDebugHelperClass
	_ANETensorDebugHelperClassOnce sync.Once
)

func getANETensorDebugHelperClass() ANETensorDebugHelperClass {
	_ANETensorDebugHelperClassOnce.Do(func() {
		_ANETensorDebugHelperClass = ANETensorDebugHelperClass{class: objc.GetClass("_ANETensorDebugHelper")}
	})
	return _ANETensorDebugHelperClass
}

// GetANETensorDebugHelperClass returns the class object for _ANETensorDebugHelper.
func GetANETensorDebugHelperClass() ANETensorDebugHelperClass {
	return getANETensorDebugHelperClass()
}

type ANETensorDebugHelperClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANETensorDebugHelperClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANETensorDebugHelperClass) Alloc() ANETensorDebugHelper {
	rv := objc.SendIfResponds[ANETensorDebugHelper](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

type ANETensorDebugHelper struct {
	objectivec.Object
}

// ANETensorDebugHelperFromID constructs a [ANETensorDebugHelper] from an objc.ID.
func ANETensorDebugHelperFromID(id objc.ID) ANETensorDebugHelper {
	return ANETensorDebugHelper{objectivec.Object{ID: id}}
}

// Ensure ANETensorDebugHelper implements IANETensorDebugHelper.
var _ IANETensorDebugHelper = ANETensorDebugHelper{}

// An interface definition for the [ANETensorDebugHelper] class.
type IANETensorDebugHelper interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a ANETensorDebugHelper) Init() ANETensorDebugHelper {
	rv := objc.SendIfResponds[ANETensorDebugHelper](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANETensorDebugHelper) Autorelease() ANETensorDebugHelper {
	rv := objc.SendIfResponds[ANETensorDebugHelper](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANETensorDebugHelper creates a new ANETensorDebugHelper instance.
func NewANETensorDebugHelper() ANETensorDebugHelper {
	class := getANETensorDebugHelperClass()
	rv := objc.SendIfResponds[ANETensorDebugHelper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_ANETensorDebugHelperClass ANETensorDebugHelperClass) NormalizeTensorDescriptorError(tensor iosurface.IOSurfaceRef, descriptor objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_ANETensorDebugHelperClass.class), objc.Sel("normalizeTensor:descriptor:error:"), tensor, descriptor, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("normalizeTensor:descriptor:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (_ANETensorDebugHelperClass ANETensorDebugHelperClass) QueryTensorsFromDebugFileError(file objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_ANETensorDebugHelperClass.class), objc.Sel("queryTensorsFromDebugFile:error:"), file, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_ANETensorDebugHelperClass ANETensorDebugHelperClass) QueryTensorsFromModelError(model objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_ANETensorDebugHelperClass.class), objc.Sel("queryTensorsFromModel:error:"), model, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_ANETensorDebugHelperClass ANETensorDebugHelperClass) ValidateDebugInfoFileError(file objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_ANETensorDebugHelperClass.class), objc.Sel("validateDebugInfoFile:error:"), file, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateDebugInfoFile:error: returned NO with nil NSError")
	}
	return rv, nil

}
