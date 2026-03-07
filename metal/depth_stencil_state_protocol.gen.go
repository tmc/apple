// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A depth and stencil state instance that specifies the depth and stencil configuration and operations used in a render pass.
//
// See: https://developer.apple.com/documentation/Metal/MTLDepthStencilState
type MTLDepthStencilState interface {
	objectivec.IObject

	// The device from which this state object was created.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDepthStencilState/device
	Device() MTLDevice

	// A string that identifies this object.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDepthStencilState/label
	Label() string

	// GpuResourceID protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDepthStencilState/gpuResourceID
	GpuResourceID() MTLResourceID
}



// MTLDepthStencilStateObject wraps an existing Objective-C object that conforms to the MTLDepthStencilState protocol.
type MTLDepthStencilStateObject struct {
	objectivec.Object
}
func (o MTLDepthStencilStateObject) BaseObject() objectivec.Object {
	return o.Object
}



// MTLDepthStencilStateObjectFromID constructs a [MTLDepthStencilStateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLDepthStencilStateObjectFromID(id objc.ID) MTLDepthStencilStateObject {
	return MTLDepthStencilStateObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// The device from which this state object was created.
//
// See: https://developer.apple.com/documentation/Metal/MTLDepthStencilState/device

func (o MTLDepthStencilStateObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}

// A string that identifies this object.
//
// See: https://developer.apple.com/documentation/Metal/MTLDepthStencilState/label

func (o MTLDepthStencilStateObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}

// See: https://developer.apple.com/documentation/Metal/MTLDepthStencilState/gpuResourceID

func (o MTLDepthStencilStateObject) GpuResourceID() MTLResourceID {
	
	rv := objc.Send[MTLResourceID](o.ID, objc.Sel("gpuResourceID"))
	return rv
	}













