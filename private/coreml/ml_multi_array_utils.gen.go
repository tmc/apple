// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLMultiArrayUtils] class.
var (
	_MLMultiArrayUtilsClass     MLMultiArrayUtilsClass
	_MLMultiArrayUtilsClassOnce sync.Once
)

func getMLMultiArrayUtilsClass() MLMultiArrayUtilsClass {
	_MLMultiArrayUtilsClassOnce.Do(func() {
		_MLMultiArrayUtilsClass = MLMultiArrayUtilsClass{class: objc.GetClass("MLMultiArrayUtils")}
	})
	return _MLMultiArrayUtilsClass
}

// GetMLMultiArrayUtilsClass returns the class object for MLMultiArrayUtils.
func GetMLMultiArrayUtilsClass() MLMultiArrayUtilsClass {
	return getMLMultiArrayUtilsClass()
}

type MLMultiArrayUtilsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLMultiArrayUtilsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLMultiArrayUtilsClass) Alloc() MLMultiArrayUtils {
	rv := objc.Send[MLMultiArrayUtils](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayUtils
type MLMultiArrayUtils struct {
	objectivec.Object
}

// MLMultiArrayUtilsFromID constructs a [MLMultiArrayUtils] from an objc.ID.
func MLMultiArrayUtilsFromID(id objc.ID) MLMultiArrayUtils {
	return MLMultiArrayUtils{objectivec.Object{ID: id}}
}

// Ensure MLMultiArrayUtils implements IMLMultiArrayUtils.
var _ IMLMultiArrayUtils = MLMultiArrayUtils{}

// An interface definition for the [MLMultiArrayUtils] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayUtils
type IMLMultiArrayUtils interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MLMultiArrayUtils) Init() MLMultiArrayUtils {
	rv := objc.Send[MLMultiArrayUtils](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLMultiArrayUtils) Autorelease() MLMultiArrayUtils {
	rv := objc.Send[MLMultiArrayUtils](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLMultiArrayUtils creates a new MLMultiArrayUtils instance.
func NewMLMultiArrayUtils() MLMultiArrayUtils {
	class := getMLMultiArrayUtilsClass()
	rv := objc.Send[MLMultiArrayUtils](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLMultiArrayUtils/stringForDataType:
func (_MLMultiArrayUtilsClass MLMultiArrayUtilsClass) StringForDataType(type_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLMultiArrayUtilsClass.class), objc.Sel("stringForDataType:"), type_)
	return objectivec.Object{ID: rv}
}
