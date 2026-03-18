// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A type that can programmatically customize a GPU frame capture.
//
// See: https://developer.apple.com/documentation/Metal/MTLCaptureScope
type MTLCaptureScope interface {
	objectivec.IObject

	// Tells Metal to begin recording command information.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCaptureScope/begin()
	BeginScope()

	// Tells Metal to stop recording command information.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCaptureScope/end()
	EndScope()

	// A string that helps you identify the capture scope.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCaptureScope/label
	Label() string

	// The device object from which you created the capture scope.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCaptureScope/device
	Device() MTLDevice

	// The command queue that this capture scope uses to limit which commands are recorded.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCaptureScope/commandQueue
	CommandQueue() MTLCommandQueue

	// If set, this scope will only capture Metal commands from the associated Metal 4 command queue. Defaults to nil (all command queues from the associated device are captured).
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCaptureScope/mtl4CommandQueue
	Mtl4CommandQueue() MTL4CommandQueue

	// A string that helps you identify the capture scope.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCaptureScope/label
	SetLabel(value string)
}

// MTLCaptureScopeObject wraps an existing Objective-C object that conforms to the MTLCaptureScope protocol.
type MTLCaptureScopeObject struct {
	objectivec.Object
}
func (o MTLCaptureScopeObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLCaptureScopeObjectFromID constructs a [MTLCaptureScopeObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLCaptureScopeObjectFromID(id objc.ID) MTLCaptureScopeObject {
	return MTLCaptureScopeObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells Metal to begin recording command information.
//
// See: https://developer.apple.com/documentation/Metal/MTLCaptureScope/begin()

func (o MTLCaptureScopeObject) BeginScope() {
	
	objc.Send[struct{}](o.ID, objc.Sel("beginScope"))
	}

// Tells Metal to stop recording command information.
//
// See: https://developer.apple.com/documentation/Metal/MTLCaptureScope/end()

func (o MTLCaptureScopeObject) EndScope() {
	
	objc.Send[struct{}](o.ID, objc.Sel("endScope"))
	}

// A string that helps you identify the capture scope.
//
// See: https://developer.apple.com/documentation/Metal/MTLCaptureScope/label

func (o MTLCaptureScopeObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}

// The device object from which you created the capture scope.
//
// See: https://developer.apple.com/documentation/Metal/MTLCaptureScope/device

func (o MTLCaptureScopeObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}

// The command queue that this capture scope uses to limit which commands are
// recorded.
//
// See: https://developer.apple.com/documentation/Metal/MTLCaptureScope/commandQueue

func (o MTLCaptureScopeObject) CommandQueue() MTLCommandQueue {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("commandQueue"))
	return MTLCommandQueueObjectFromID(rv)
	}

// If set, this scope will only capture Metal commands from the associated
// Metal 4 command queue. Defaults to nil (all command queues from the
// associated device are captured).
//
// See: https://developer.apple.com/documentation/Metal/MTLCaptureScope/mtl4CommandQueue

func (o MTLCaptureScopeObject) Mtl4CommandQueue() MTL4CommandQueue {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("mtl4CommandQueue"))
	return MTL4CommandQueueObjectFromID(rv)
	}

func (o MTLCaptureScopeObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}

