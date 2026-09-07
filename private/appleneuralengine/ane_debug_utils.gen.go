// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEDebugUtils] class.
var (
	_ANEDebugUtilsClass     ANEDebugUtilsClass
	_ANEDebugUtilsClassOnce sync.Once
)

func getANEDebugUtilsClass() ANEDebugUtilsClass {
	_ANEDebugUtilsClassOnce.Do(func() {
		_ANEDebugUtilsClass = ANEDebugUtilsClass{class: objc.GetClass("_ANEDebugUtils")}
	})
	return _ANEDebugUtilsClass
}

// GetANEDebugUtilsClass returns the class object for _ANEDebugUtils.
func GetANEDebugUtilsClass() ANEDebugUtilsClass {
	return getANEDebugUtilsClass()
}

type ANEDebugUtilsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANEDebugUtilsClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEDebugUtilsClass) Alloc() ANEDebugUtils {
	rv := objc.SendIfResponds[ANEDebugUtils](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

type ANEDebugUtils struct {
	objectivec.Object
}

// ANEDebugUtilsFromID constructs a [ANEDebugUtils] from an objc.ID.
func ANEDebugUtilsFromID(id objc.ID) ANEDebugUtils {
	return ANEDebugUtils{objectivec.Object{ID: id}}
}

// Ensure ANEDebugUtils implements IANEDebugUtils.
var _ IANEDebugUtils = ANEDebugUtils{}

// An interface definition for the [ANEDebugUtils] class.
type IANEDebugUtils interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a ANEDebugUtils) Init() ANEDebugUtils {
	rv := objc.SendIfResponds[ANEDebugUtils](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEDebugUtils) Autorelease() ANEDebugUtils {
	rv := objc.SendIfResponds[ANEDebugUtils](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEDebugUtils creates a new ANEDebugUtils instance.
func NewANEDebugUtils() ANEDebugUtils {
	class := getANEDebugUtilsClass()
	rv := objc.SendIfResponds[ANEDebugUtils](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_ANEDebugUtilsClass ANEDebugUtilsClass) ApplyDebugEnvToOptions(options objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEDebugUtilsClass.class), objc.Sel("applyDebugEnvToOptions:"), options)
	return objectivec.Object{ID: rv}
}
func (_ANEDebugUtilsClass ANEDebugUtilsClass) ParseDebugEnvVar() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEDebugUtilsClass.class), objc.Sel("parseDebugEnvVar"))
	return objectivec.Object{ID: rv}
}
func (_ANEDebugUtilsClass ANEDebugUtilsClass) ParseDebugString(string_ objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEDebugUtilsClass.class), objc.Sel("parseDebugString:"), string_)
	return objectivec.Object{ID: rv}
}
