// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A log entry a Metal device generates when the it runs a command buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionLog
type MTLFunctionLog interface {
	objectivec.IObject

	// The type of message that was logged.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFunctionLog/type
	Type() MTLFunctionLogType

	// If known, the location of the logging command within a shader source file.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFunctionLog/debugLocation
	DebugLocation() MTLFunctionLogDebugLocation

	// The label for the encoder that logged the message.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFunctionLog/encoderLabel
	EncoderLabel() string

	// When known, the function object corresponding to the logged message.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFunctionLog/function
	Function() MTLFunction
}

// MTLFunctionLogObject wraps an existing Objective-C object that conforms to the MTLFunctionLog protocol.
type MTLFunctionLogObject struct {
	objectivec.Object
}

func (o MTLFunctionLogObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLFunctionLogObjectFromID constructs a [MTLFunctionLogObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLFunctionLogObjectFromID(id objc.ID) MTLFunctionLogObject {
	return MTLFunctionLogObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The type of message that was logged.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionLog/type
func (o MTLFunctionLogObject) Type() MTLFunctionLogType {
	rv := objc.Send[MTLFunctionLogType](o.ID, objc.Sel("type"))
	return rv
}

// If known, the location of the logging command within a shader source file.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionLog/debugLocation
func (o MTLFunctionLogObject) DebugLocation() MTLFunctionLogDebugLocation {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("debugLocation"))
	return MTLFunctionLogDebugLocationObjectFromID(rv)
}

// The label for the encoder that logged the message.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionLog/encoderLabel
func (o MTLFunctionLogObject) EncoderLabel() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("encoderLabel"))
	return foundation.NSStringFromID(rv).String()
}

// When known, the function object corresponding to the logged message.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionLog/function
func (o MTLFunctionLogObject) Function() MTLFunction {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("function"))
	return MTLFunctionObjectFromID(rv)
}
