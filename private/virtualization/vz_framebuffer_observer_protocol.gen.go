// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _VZFramebufferObserver protocol.
type VZFramebufferObserver interface {
	objectivec.IObject

	// GetDisplayProtectionOptions protocol.
	GetDisplayProtectionOptions() unsafe.Pointer
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

func (o VZFramebufferObserverObject) FramebufferDidUpdateCursor(framebuffer objectivec.IObject, cursor unsafe.Pointer) {
	objc.Send[struct{}](o.ID, objc.Sel("framebuffer:didUpdateCursor:"), framebuffer, cursor)
}
func (o VZFramebufferObserverObject) FramebufferDidUpdateFrame(framebuffer objectivec.IObject, frame unsafe.Pointer) {
	objc.Send[struct{}](o.ID, objc.Sel("framebuffer:didUpdateFrame:"), framebuffer, frame)
}
func (o VZFramebufferObserverObject) FramebufferDidUpdateGraphicsOrientation(framebuffer objectivec.IObject, orientation int64) {
	objc.Send[struct{}](o.ID, objc.Sel("framebuffer:didUpdateGraphicsOrientation:"), framebuffer, orientation)
}
func (o VZFramebufferObserverObject) FramebufferDidUpdateColorSpace(space objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("framebufferDidUpdateColorSpace:"), space)
}
func (o VZFramebufferObserverObject) GetDisplayProtectionOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("getDisplayProtectionOptions"))
	return rv
}
