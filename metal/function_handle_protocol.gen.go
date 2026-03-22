// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// An object representing a function that you can add to a visible function table.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionHandle
type MTLFunctionHandle interface {
	objectivec.IObject

	// The device object that created the shader function.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFunctionHandle/device
	Device() MTLDevice

	// The shader function’s type.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFunctionHandle/functionType
	FunctionType() MTLFunctionType

	// The function’s name.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFunctionHandle/name
	Name() string

	// GpuResourceID protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFunctionHandle/gpuResourceID
	GpuResourceID() MTLResourceID
}

// MTLFunctionHandleObject wraps an existing Objective-C object that conforms to the MTLFunctionHandle protocol.
type MTLFunctionHandleObject struct {
	objectivec.Object
}
func (o MTLFunctionHandleObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLFunctionHandleObjectFromID constructs a [MTLFunctionHandleObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLFunctionHandleObjectFromID(id objc.ID) MTLFunctionHandleObject {
	return MTLFunctionHandleObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The device object that created the shader function.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionHandle/device
func (o MTLFunctionHandleObject) Device() MTLDevice {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}
// The shader function’s type.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionHandle/functionType
func (o MTLFunctionHandleObject) FunctionType() MTLFunctionType {
	rv := objc.Send[MTLFunctionType](o.ID, objc.Sel("functionType"))
	return rv
	}
// The function’s name.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionHandle/name
func (o MTLFunctionHandleObject) Name() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
	}
// See: https://developer.apple.com/documentation/Metal/MTLFunctionHandle/gpuResourceID
func (o MTLFunctionHandleObject) GpuResourceID() MTLResourceID {
	rv := objc.Send[MTLResourceID](o.ID, objc.Sel("gpuResourceID"))
	return rv
	}

