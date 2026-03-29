// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _VZFramebufferObserver protocol.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferObserver
type VZFramebufferObserver interface {
	objectivec.IObject

	// GetDisplayProtectionOptions protocol.
	//
	// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferObserver/getDisplayProtectionOptions
	GetDisplayProtectionOptions() objectivec.IObject
}

// VZFramebufferObserverObject wraps an existing Objective-C object that conforms to the VZFramebufferObserver protocol.
type VZFramebufferObserverObject struct {
	objectivec.Object
}
func (o VZFramebufferObserverObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZFramebufferObserverObjectFromID constructs a [VZFramebufferObserverObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZFramebufferObserverObjectFromID(id objc.ID) VZFramebufferObserverObject {
	return VZFramebufferObserverObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferObserver/framebuffer:didUpdateCursor:
func (o VZFramebufferObserverObject) FramebufferDidUpdateCursor(framebuffer objectivec.IObject, cursor objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("framebuffer:didUpdateCursor:"), framebuffer, cursor)
	}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferObserver/framebuffer:didUpdateFrame:
func (o VZFramebufferObserverObject) FramebufferDidUpdateFrame(framebuffer objectivec.IObject, frame objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("framebuffer:didUpdateFrame:"), framebuffer, frame)
	}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferObserver/framebuffer:didUpdateGraphicsOrientation:
func (o VZFramebufferObserverObject) FramebufferDidUpdateGraphicsOrientation(framebuffer objectivec.IObject, orientation int64) {
	objc.Send[struct{}](o.ID, objc.Sel("framebuffer:didUpdateGraphicsOrientation:"), framebuffer, orientation)
	}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferObserver/framebufferDidUpdateColorSpace:
func (o VZFramebufferObserverObject) FramebufferDidUpdateColorSpace(space objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("framebufferDidUpdateColorSpace:"), space)
	}
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferObserver/getDisplayProtectionOptions
func (o VZFramebufferObserverObject) GetDisplayProtectionOptions() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("getDisplayProtectionOptions"))
	return objectivec.Object{ID: rv}
	}

