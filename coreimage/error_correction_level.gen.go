// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ErrorCorrectionLevel] class.
var (
	_ErrorCorrectionLevelClass     ErrorCorrectionLevelClass
	_ErrorCorrectionLevelClassOnce sync.Once
)

func getErrorCorrectionLevelClass() ErrorCorrectionLevelClass {
	_ErrorCorrectionLevelClassOnce.Do(func() {
		_ErrorCorrectionLevelClass = ErrorCorrectionLevelClass{class: objc.GetClass("errorCorrectionLevel")}
	})
	return _ErrorCorrectionLevelClass
}

// GetErrorCorrectionLevelClass returns the class object for errorCorrectionLevel.
func GetErrorCorrectionLevelClass() ErrorCorrectionLevelClass {
	return getErrorCorrectionLevelClass()
}

type ErrorCorrectionLevelClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec ErrorCorrectionLevelClass) Alloc() ErrorCorrectionLevel {
	rv := objc.Send[ErrorCorrectionLevel](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/errorCorrectionLevel-c.ivar
type ErrorCorrectionLevel struct {
	objectivec.Object
}

// ErrorCorrectionLevelFromID constructs a [ErrorCorrectionLevel] from an objc.ID.
func ErrorCorrectionLevelFromID(id objc.ID) ErrorCorrectionLevel {
	return ErrorCorrectionLevel{objectivec.Object{ID: id}}
}
// Ensure ErrorCorrectionLevel implements IErrorCorrectionLevel.
var _ IErrorCorrectionLevel = ErrorCorrectionLevel{}

// An interface definition for the [ErrorCorrectionLevel] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/errorCorrectionLevel-c.ivar
type IErrorCorrectionLevel interface {
	objectivec.IObject
}

// Init initializes the instance.
func (e ErrorCorrectionLevel) Init() ErrorCorrectionLevel {
	rv := objc.Send[ErrorCorrectionLevel](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ErrorCorrectionLevel) Autorelease() ErrorCorrectionLevel {
	rv := objc.Send[ErrorCorrectionLevel](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewErrorCorrectionLevel creates a new ErrorCorrectionLevel instance.
func NewErrorCorrectionLevel() ErrorCorrectionLevel {
	class := getErrorCorrectionLevelClass()
	rv := objc.Send[ErrorCorrectionLevel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

