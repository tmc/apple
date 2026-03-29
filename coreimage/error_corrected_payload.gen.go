// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ErrorCorrectedPayload] class.
var (
	_ErrorCorrectedPayloadClass     ErrorCorrectedPayloadClass
	_ErrorCorrectedPayloadClassOnce sync.Once
)

func getErrorCorrectedPayloadClass() ErrorCorrectedPayloadClass {
	_ErrorCorrectedPayloadClassOnce.Do(func() {
		_ErrorCorrectedPayloadClass = ErrorCorrectedPayloadClass{class: objc.GetClass("errorCorrectedPayload")}
	})
	return _ErrorCorrectedPayloadClass
}

// GetErrorCorrectedPayloadClass returns the class object for errorCorrectedPayload.
func GetErrorCorrectedPayloadClass() ErrorCorrectedPayloadClass {
	return getErrorCorrectedPayloadClass()
}

type ErrorCorrectedPayloadClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ErrorCorrectedPayloadClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ErrorCorrectedPayloadClass) Alloc() ErrorCorrectedPayload {
	rv := objc.Send[ErrorCorrectedPayload](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/errorCorrectedPayload-c.ivar
type ErrorCorrectedPayload struct {
	objectivec.Object
}

// ErrorCorrectedPayloadFromID constructs a [ErrorCorrectedPayload] from an objc.ID.
func ErrorCorrectedPayloadFromID(id objc.ID) ErrorCorrectedPayload {
	return ErrorCorrectedPayload{objectivec.Object{ID: id}}
}
// Ensure ErrorCorrectedPayload implements IErrorCorrectedPayload.
var _ IErrorCorrectedPayload = ErrorCorrectedPayload{}

// An interface definition for the [ErrorCorrectedPayload] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/errorCorrectedPayload-c.ivar
type IErrorCorrectedPayload interface {
	objectivec.IObject
}

// Init initializes the instance.
func (e ErrorCorrectedPayload) Init() ErrorCorrectedPayload {
	rv := objc.Send[ErrorCorrectedPayload](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ErrorCorrectedPayload) Autorelease() ErrorCorrectedPayload {
	rv := objc.Send[ErrorCorrectedPayload](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewErrorCorrectedPayload creates a new ErrorCorrectedPayload instance.
func NewErrorCorrectedPayload() ErrorCorrectedPayload {
	class := getErrorCorrectedPayloadClass()
	rv := objc.Send[ErrorCorrectedPayload](objc.ID(class.class), objc.Sel("new"))
	return rv
}

