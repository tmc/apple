// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLParameterUtils] class.
var (
	_MLParameterUtilsClass     MLParameterUtilsClass
	_MLParameterUtilsClassOnce sync.Once
)

func getMLParameterUtilsClass() MLParameterUtilsClass {
	_MLParameterUtilsClassOnce.Do(func() {
		_MLParameterUtilsClass = MLParameterUtilsClass{class: objc.GetClass("MLParameterUtils")}
	})
	return _MLParameterUtilsClass
}

// GetMLParameterUtilsClass returns the class object for MLParameterUtils.
func GetMLParameterUtilsClass() MLParameterUtilsClass {
	return getMLParameterUtilsClass()
}

type MLParameterUtilsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLParameterUtilsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLParameterUtilsClass) Alloc() MLParameterUtils {
	rv := objc.Send[MLParameterUtils](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLParameterUtils
type MLParameterUtils struct {
	objectivec.Object
}

// MLParameterUtilsFromID constructs a [MLParameterUtils] from an objc.ID.
func MLParameterUtilsFromID(id objc.ID) MLParameterUtils {
	return MLParameterUtils{objectivec.Object{ID: id}}
}
// Ensure MLParameterUtils implements IMLParameterUtils.
var _ IMLParameterUtils = MLParameterUtils{}

// An interface definition for the [MLParameterUtils] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterUtils
type IMLParameterUtils interface {
	objectivec.IObject
}

// Init initializes the instance.
func (p MLParameterUtils) Init() MLParameterUtils {
	rv := objc.Send[MLParameterUtils](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLParameterUtils) Autorelease() MLParameterUtils {
	rv := objc.Send[MLParameterUtils](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLParameterUtils creates a new MLParameterUtils instance.
func NewMLParameterUtils() MLParameterUtils {
	class := getMLParameterUtilsClass()
	rv := objc.Send[MLParameterUtils](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLParameterUtils/appendParameterDescriptions:toModelDescription:
func (_MLParameterUtilsClass MLParameterUtilsClass) AppendParameterDescriptionsToModelDescription(descriptions objectivec.IObject, description objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_MLParameterUtilsClass.class), objc.Sel("appendParameterDescriptions:toModelDescription:"), descriptions, description)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterUtils/deScopeParameters:byDeletingPrefixingScope:
func (_MLParameterUtilsClass MLParameterUtilsClass) DeScopeParametersByDeletingPrefixingScope(parameters objectivec.IObject, scope objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterUtilsClass.class), objc.Sel("deScopeParameters:byDeletingPrefixingScope:"), parameters, scope)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterUtils/numberForKey:inDictionary:
func (_MLParameterUtilsClass MLParameterUtilsClass) NumberForKeyInDictionary(key objectivec.IObject, dictionary objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterUtilsClass.class), objc.Sel("numberForKey:inDictionary:"), key, dictionary)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterUtils/objectForKey:class:dictionary:
func (_MLParameterUtilsClass MLParameterUtilsClass) ObjectForKeyClassDictionary(key objectivec.IObject, class objc.Class, dictionary objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterUtilsClass.class), objc.Sel("objectForKey:class:dictionary:"), key, class, dictionary)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLParameterUtils/stringForKey:inDictionary:
func (_MLParameterUtilsClass MLParameterUtilsClass) StringForKeyInDictionary(key objectivec.IObject, dictionary objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterUtilsClass.class), objc.Sel("stringForKey:inDictionary:"), key, dictionary)
	return objectivec.Object{ID: rv}
}

