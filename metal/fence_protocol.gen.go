// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A synchronization mechanism that orders memory operations between GPU passes.
//
// See: https://developer.apple.com/documentation/Metal/MTLFence
type MTLFence interface {
	objectivec.IObject

	// The device object that created the fence.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFence/device
	Device() MTLDevice

	// A string that identifies the fence.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFence/label
	Label() string

	// A string that identifies the fence.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFence/label
	SetLabel(value string)
}

// MTLFenceObject wraps an existing Objective-C object that conforms to the MTLFence protocol.
type MTLFenceObject struct {
	objectivec.Object
}

func (o MTLFenceObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLFenceObjectFromID constructs a [MTLFenceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLFenceObjectFromID(id objc.ID) MTLFenceObject {
	return MTLFenceObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The device object that created the fence.
//
// See: https://developer.apple.com/documentation/Metal/MTLFence/device
func (o MTLFenceObject) Device() MTLDevice {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
}

// A string that identifies the fence.
//
// See: https://developer.apple.com/documentation/Metal/MTLFence/label
func (o MTLFenceObject) Label() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}

// A string that identifies the fence.
//
// # Discussion
//
// Object and command labels are useful identifiers at runtime or when
// profiling and debugging your app using any Metal tool. See [Naming
// resources and commands].
//
// See: https://developer.apple.com/documentation/Metal/MTLFence/label
//
// [Naming resources and commands]: https://developer.apple.com/documentation/Xcode/Naming-resources-and-commands
func (o MTLFenceObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}
