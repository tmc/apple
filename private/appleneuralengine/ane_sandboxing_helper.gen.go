// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANESandboxingHelper] class.
var (
	_ANESandboxingHelperClass     ANESandboxingHelperClass
	_ANESandboxingHelperClassOnce sync.Once
)

func getANESandboxingHelperClass() ANESandboxingHelperClass {
	_ANESandboxingHelperClassOnce.Do(func() {
		_ANESandboxingHelperClass = ANESandboxingHelperClass{class: objc.GetClass("_ANESandboxingHelper")}
	})
	return _ANESandboxingHelperClass
}

// GetANESandboxingHelperClass returns the class object for _ANESandboxingHelper.
func GetANESandboxingHelperClass() ANESandboxingHelperClass {
	return getANESandboxingHelperClass()
}

type ANESandboxingHelperClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANESandboxingHelperClass) Alloc() ANESandboxingHelper {
	rv := objc.Send[ANESandboxingHelper](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESandboxingHelper
type ANESandboxingHelper struct {
	objectivec.Object
}

// ANESandboxingHelperFromID constructs a [ANESandboxingHelper] from an objc.ID.
func ANESandboxingHelperFromID(id objc.ID) ANESandboxingHelper {
	return ANESandboxingHelper{objectivec.Object{id}}
}
// Ensure ANESandboxingHelper implements IANESandboxingHelper.
var _ IANESandboxingHelper = ANESandboxingHelper{}





// An interface definition for the [ANESandboxingHelper] class.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESandboxingHelper
type IANESandboxingHelper interface {
	objectivec.IObject
}





// Init initializes the instance.
func (a ANESandboxingHelper) Init() ANESandboxingHelper {
	rv := objc.Send[ANESandboxingHelper](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANESandboxingHelper) Autorelease() ANESandboxingHelper {
	rv := objc.Send[ANESandboxingHelper](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANESandboxingHelper creates a new ANESandboxingHelper instance.
func NewANESandboxingHelper() ANESandboxingHelper {
	class := getANESandboxingHelperClass()
	rv := objc.Send[ANESandboxingHelper](objc.ID(class.class), objc.Sel("new"))
	return rv
}














//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESandboxingHelper/canAccessPathAt:methodName:error:
func (_ANESandboxingHelperClass ANESandboxingHelperClass) CanAccessPathAtMethodNameError(at objectivec.IObject, name objectivec.IObject) (bool, error) {
			var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_ANESandboxingHelperClass.class), objc.Sel("canAccessPathAt:methodName:error:"), at, name, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("canAccessPathAt:methodName:error: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESandboxingHelper/consumeSandboxExtension:forModel:error:
func (_ANESandboxingHelperClass ANESandboxingHelperClass) ConsumeSandboxExtensionForModelError(extension objectivec.IObject, model objectivec.IObject) (int64, error) {
			var errorPtr objc.ID
	rv := objc.Send[int64](objc.ID(_ANESandboxingHelperClass.class), objc.Sel("consumeSandboxExtension:forModel:error:"), extension, model, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESandboxingHelper/consumeSandboxExtension:forPath:error:
func (_ANESandboxingHelperClass ANESandboxingHelperClass) ConsumeSandboxExtensionForPathError(extension objectivec.IObject, path objectivec.IObject) (int64, error) {
			var errorPtr objc.ID
	rv := objc.Send[int64](objc.ID(_ANESandboxingHelperClass.class), objc.Sel("consumeSandboxExtension:forPath:error:"), extension, path, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESandboxingHelper/issueSandboxExtensionForModel:error:
func (_ANESandboxingHelperClass ANESandboxingHelperClass) IssueSandboxExtensionForModelError(model objectivec.IObject) (objectivec.IObject, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_ANESandboxingHelperClass.class), objc.Sel("issueSandboxExtensionForModel:error:"), model, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESandboxingHelper/issueSandboxExtensionForPath:error:
func (_ANESandboxingHelperClass ANESandboxingHelperClass) IssueSandboxExtensionForPathError(path objectivec.IObject) (objectivec.IObject, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_ANESandboxingHelperClass.class), objc.Sel("issueSandboxExtensionForPath:error:"), path, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESandboxingHelper/releaseSandboxExtension:handle:
func (_ANESandboxingHelperClass ANESandboxingHelperClass) ReleaseSandboxExtensionHandle(extension objectivec.IObject, handle int64) bool {
	rv := objc.Send[bool](objc.ID(_ANESandboxingHelperClass.class), objc.Sel("releaseSandboxExtension:handle:"), extension, handle)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESandboxingHelper/sandboxExtensionPathForModelURL:
func (_ANESandboxingHelperClass ANESandboxingHelperClass) SandboxExtensionPathForModelURL(url foundation.INSURL) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANESandboxingHelperClass.class), objc.Sel("sandboxExtensionPathForModelURL:"), url)
	return objectivec.Object{ID: rv}
}






















