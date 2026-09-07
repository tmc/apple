// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLPlatformUtils] class.
var (
	_MLPlatformUtilsClass     MLPlatformUtilsClass
	_MLPlatformUtilsClassOnce sync.Once
)

func getMLPlatformUtilsClass() MLPlatformUtilsClass {
	_MLPlatformUtilsClassOnce.Do(func() {
		_MLPlatformUtilsClass = MLPlatformUtilsClass{class: objc.GetClass("MLPlatformUtils")}
	})
	return _MLPlatformUtilsClass
}

// GetMLPlatformUtilsClass returns the class object for MLPlatformUtils.
func GetMLPlatformUtilsClass() MLPlatformUtilsClass {
	return getMLPlatformUtilsClass()
}

type MLPlatformUtilsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLPlatformUtilsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLPlatformUtilsClass) Alloc() MLPlatformUtils {
	rv := objc.SendIfResponds[MLPlatformUtils](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

type MLPlatformUtils struct {
	objectivec.Object
}

// MLPlatformUtilsFromID constructs a [MLPlatformUtils] from an objc.ID.
func MLPlatformUtilsFromID(id objc.ID) MLPlatformUtils {
	return MLPlatformUtils{objectivec.Object{ID: id}}
}

// Ensure MLPlatformUtils implements IMLPlatformUtils.
var _ IMLPlatformUtils = MLPlatformUtils{}

// An interface definition for the [MLPlatformUtils] class.
type IMLPlatformUtils interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MLPlatformUtils) Init() MLPlatformUtils {
	rv := objc.SendIfResponds[MLPlatformUtils](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLPlatformUtils) Autorelease() MLPlatformUtils {
	rv := objc.SendIfResponds[MLPlatformUtils](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLPlatformUtils creates a new MLPlatformUtils instance.
func NewMLPlatformUtils() MLPlatformUtils {
	class := getMLPlatformUtilsClass()
	rv := objc.SendIfResponds[MLPlatformUtils](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_MLPlatformUtilsClass MLPlatformUtilsClass) IsInternalBuild() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLPlatformUtilsClass.class), objc.Sel("isInternalBuild"))
	return rv
}
