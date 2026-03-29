// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANECloneHelper] class.
var (
	_ANECloneHelperClass     ANECloneHelperClass
	_ANECloneHelperClassOnce sync.Once
)

func getANECloneHelperClass() ANECloneHelperClass {
	_ANECloneHelperClassOnce.Do(func() {
		_ANECloneHelperClass = ANECloneHelperClass{class: objc.GetClass("_ANECloneHelper")}
	})
	return _ANECloneHelperClass
}

// GetANECloneHelperClass returns the class object for _ANECloneHelper.
func GetANECloneHelperClass() ANECloneHelperClass {
	return getANECloneHelperClass()
}

type ANECloneHelperClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANECloneHelperClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANECloneHelperClass) Alloc() ANECloneHelper {
	rv := objc.Send[ANECloneHelper](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECloneHelper
type ANECloneHelper struct {
	objectivec.Object
}

// ANECloneHelperFromID constructs a [ANECloneHelper] from an objc.ID.
func ANECloneHelperFromID(id objc.ID) ANECloneHelper {
	return ANECloneHelper{objectivec.Object{ID: id}}
}
// Ensure ANECloneHelper implements IANECloneHelper.
var _ IANECloneHelper = ANECloneHelper{}

// An interface definition for the [ANECloneHelper] class.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECloneHelper
type IANECloneHelper interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a ANECloneHelper) Init() ANECloneHelper {
	rv := objc.Send[ANECloneHelper](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANECloneHelper) Autorelease() ANECloneHelper {
	rv := objc.Send[ANECloneHelper](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANECloneHelper creates a new ANECloneHelper instance.
func NewANECloneHelper() ANECloneHelper {
	class := getANECloneHelperClass()
	rv := objc.Send[ANECloneHelper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECloneHelper/cloneIfWritable:isEncryptedModel:cloneDirectory:
func (_ANECloneHelperClass ANECloneHelperClass) CloneIfWritableIsEncryptedModelCloneDirectory(writable objectivec.IObject, model bool, directory objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANECloneHelperClass.class), objc.Sel("cloneIfWritable:isEncryptedModel:cloneDirectory:"), writable, model, directory)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECloneHelper/shouldSkipCloneFor:isEncryptedModel:
func (_ANECloneHelperClass ANECloneHelperClass) ShouldSkipCloneForIsEncryptedModel(for_ objectivec.IObject, model bool) bool {
	rv := objc.Send[bool](objc.ID(_ANECloneHelperClass.class), objc.Sel("shouldSkipCloneFor:isEncryptedModel:"), for_, model)
	return rv
}

