// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEServicesLog] class.
var (
	_ANEServicesLogClass     ANEServicesLogClass
	_ANEServicesLogClassOnce sync.Once
)

func getANEServicesLogClass() ANEServicesLogClass {
	_ANEServicesLogClassOnce.Do(func() {
		_ANEServicesLogClass = ANEServicesLogClass{class: objc.GetClass("ANEServicesLog")}
	})
	return _ANEServicesLogClass
}

// GetANEServicesLogClass returns the class object for ANEServicesLog.
func GetANEServicesLogClass() ANEServicesLogClass {
	return getANEServicesLogClass()
}

type ANEServicesLogClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANEServicesLogClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEServicesLogClass) Alloc() ANEServicesLog {
	rv := objc.Send[ANEServicesLog](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ANEServicesLog
type ANEServicesLog struct {
	objectivec.Object
}

// ANEServicesLogFromID constructs a [ANEServicesLog] from an objc.ID.
func ANEServicesLogFromID(id objc.ID) ANEServicesLog {
	return ANEServicesLog{objectivec.Object{ID: id}}
}

// Ensure ANEServicesLog implements IANEServicesLog.
var _ IANEServicesLog = ANEServicesLog{}

// An interface definition for the [ANEServicesLog] class.
//
// See: https://developer.apple.com/documentation/Espresso/ANEServicesLog
type IANEServicesLog interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a ANEServicesLog) Init() ANEServicesLog {
	rv := objc.Send[ANEServicesLog](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEServicesLog) Autorelease() ANEServicesLog {
	rv := objc.Send[ANEServicesLog](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEServicesLog creates a new ANEServicesLog instance.
func NewANEServicesLog() ANEServicesLog {
	class := getANEServicesLogClass()
	rv := objc.Send[ANEServicesLog](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ANEServicesLog/handle
func (_ANEServicesLogClass ANEServicesLogClass) Handle() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEServicesLogClass.class), objc.Sel("handle"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/ANEServicesLog/services
func (_ANEServicesLogClass ANEServicesLogClass) Services() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEServicesLogClass.class), objc.Sel("services"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/ANEServicesLog/test
func (_ANEServicesLogClass ANEServicesLogClass) Test() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEServicesLogClass.class), objc.Sel("test"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/ANEServicesLog/verbose
func (_ANEServicesLogClass ANEServicesLogClass) Verbose() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEServicesLogClass.class), objc.Sel("verbose"))
	return objectivec.Object{ID: rv}
}
