// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANELog] class.
var (
	_ANELogClass     ANELogClass
	_ANELogClassOnce sync.Once
)

func getANELogClass() ANELogClass {
	_ANELogClassOnce.Do(func() {
		_ANELogClass = ANELogClass{class: objc.GetClass("_ANELog")}
	})
	return _ANELogClass
}

// GetANELogClass returns the class object for _ANELog.
func GetANELogClass() ANELogClass {
	return getANELogClass()
}

type ANELogClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANELogClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANELogClass) Alloc() ANELog {
	rv := objc.SendIfResponds[ANELog](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

type ANELog struct {
	objectivec.Object
}

// ANELogFromID constructs a [ANELog] from an objc.ID.
func ANELogFromID(id objc.ID) ANELog {
	return ANELog{objectivec.Object{ID: id}}
}

// Ensure ANELog implements IANELog.
var _ IANELog = ANELog{}

// An interface definition for the [ANELog] class.
type IANELog interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a ANELog) Init() ANELog {
	rv := objc.SendIfResponds[ANELog](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANELog) Autorelease() ANELog {
	rv := objc.SendIfResponds[ANELog](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANELog creates a new ANELog instance.
func NewANELog() ANELog {
	class := getANELogClass()
	rv := objc.SendIfResponds[ANELog](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_ANELogClass ANELogClass) Common() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANELogClass.class), objc.Sel("common"))
	return objectivec.Object{ID: rv}
}
func (_ANELogClass ANELogClass) Compiler() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANELogClass.class), objc.Sel("compiler"))
	return objectivec.Object{ID: rv}
}
func (_ANELogClass ANELogClass) Daemon() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANELogClass.class), objc.Sel("daemon"))
	return objectivec.Object{ID: rv}
}
func (_ANELogClass ANELogClass) Framework() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANELogClass.class), objc.Sel("framework"))
	return objectivec.Object{ID: rv}
}
func (_ANELogClass ANELogClass) Maintenance() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANELogClass.class), objc.Sel("maintenance"))
	return objectivec.Object{ID: rv}
}
func (_ANELogClass ANELogClass) Tests() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANELogClass.class), objc.Sel("tests"))
	return objectivec.Object{ID: rv}
}
func (_ANELogClass ANELogClass) Tool() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANELogClass.class), objc.Sel("tool"))
	return objectivec.Object{ID: rv}
}
