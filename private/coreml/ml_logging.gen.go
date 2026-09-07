// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLLogging] class.
var (
	_MLLoggingClass     MLLoggingClass
	_MLLoggingClassOnce sync.Once
)

func getMLLoggingClass() MLLoggingClass {
	_MLLoggingClassOnce.Do(func() {
		_MLLoggingClass = MLLoggingClass{class: objc.GetClass("MLLogging")}
	})
	return _MLLoggingClass
}

// GetMLLoggingClass returns the class object for MLLogging.
func GetMLLoggingClass() MLLoggingClass {
	return getMLLoggingClass()
}

type MLLoggingClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLLoggingClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLLoggingClass) Alloc() MLLogging {
	rv := objc.SendIfResponds[MLLogging](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

type MLLogging struct {
	objectivec.Object
}

// MLLoggingFromID constructs a [MLLogging] from an objc.ID.
func MLLoggingFromID(id objc.ID) MLLogging {
	return MLLogging{objectivec.Object{ID: id}}
}

// Ensure MLLogging implements IMLLogging.
var _ IMLLogging = MLLogging{}

// An interface definition for the [MLLogging] class.
type IMLLogging interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MLLogging) Init() MLLogging {
	rv := objc.SendIfResponds[MLLogging](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLLogging) Autorelease() MLLogging {
	rv := objc.SendIfResponds[MLLogging](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLLogging creates a new MLLogging instance.
func NewMLLogging() MLLogging {
	class := getMLLoggingClass()
	rv := objc.SendIfResponds[MLLogging](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_MLLoggingClass MLLoggingClass) CoreChannel() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLLoggingClass.class), objc.Sel("coreChannel"))
	return objectivec.Object{ID: rv}
}
