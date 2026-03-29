// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLUpdateProgressHandlersUtils] class.
var (
	_MLUpdateProgressHandlersUtilsClass     MLUpdateProgressHandlersUtilsClass
	_MLUpdateProgressHandlersUtilsClassOnce sync.Once
)

func getMLUpdateProgressHandlersUtilsClass() MLUpdateProgressHandlersUtilsClass {
	_MLUpdateProgressHandlersUtilsClassOnce.Do(func() {
		_MLUpdateProgressHandlersUtilsClass = MLUpdateProgressHandlersUtilsClass{class: objc.GetClass("MLUpdateProgressHandlersUtils")}
	})
	return _MLUpdateProgressHandlersUtilsClass
}

// GetMLUpdateProgressHandlersUtilsClass returns the class object for MLUpdateProgressHandlersUtils.
func GetMLUpdateProgressHandlersUtilsClass() MLUpdateProgressHandlersUtilsClass {
	return getMLUpdateProgressHandlersUtilsClass()
}

type MLUpdateProgressHandlersUtilsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLUpdateProgressHandlersUtilsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLUpdateProgressHandlersUtilsClass) Alloc() MLUpdateProgressHandlersUtils {
	rv := objc.Send[MLUpdateProgressHandlersUtils](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLUpdateProgressHandlersUtils
type MLUpdateProgressHandlersUtils struct {
	objectivec.Object
}

// MLUpdateProgressHandlersUtilsFromID constructs a [MLUpdateProgressHandlersUtils] from an objc.ID.
func MLUpdateProgressHandlersUtilsFromID(id objc.ID) MLUpdateProgressHandlersUtils {
	return MLUpdateProgressHandlersUtils{objectivec.Object{ID: id}}
}
// Ensure MLUpdateProgressHandlersUtils implements IMLUpdateProgressHandlersUtils.
var _ IMLUpdateProgressHandlersUtils = MLUpdateProgressHandlersUtils{}

// An interface definition for the [MLUpdateProgressHandlersUtils] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLUpdateProgressHandlersUtils
type IMLUpdateProgressHandlersUtils interface {
	objectivec.IObject
}

// Init initializes the instance.
func (u MLUpdateProgressHandlersUtils) Init() MLUpdateProgressHandlersUtils {
	rv := objc.Send[MLUpdateProgressHandlersUtils](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u MLUpdateProgressHandlersUtils) Autorelease() MLUpdateProgressHandlersUtils {
	rv := objc.Send[MLUpdateProgressHandlersUtils](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLUpdateProgressHandlersUtils creates a new MLUpdateProgressHandlersUtils instance.
func NewMLUpdateProgressHandlersUtils() MLUpdateProgressHandlersUtils {
	class := getMLUpdateProgressHandlersUtilsClass()
	rv := objc.Send[MLUpdateProgressHandlersUtils](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLUpdateProgressHandlersUtils/progressEventsToString:
func (_MLUpdateProgressHandlersUtilsClass MLUpdateProgressHandlersUtilsClass) ProgressEventsToString(string_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLUpdateProgressHandlersUtilsClass.class), objc.Sel("progressEventsToString:"), string_)
	return objectivec.Object{ID: rv}
}

