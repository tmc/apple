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

	// FramebufferDidUpdateCursor protocol.
	FramebufferDidUpdateCursor(framebuffer objectivec.IObject, cursor unsafe.Pointer)

	// FramebufferDidUpdateFrame protocol.
	FramebufferDidUpdateFrame(framebuffer objectivec.IObject, frame unsafe.Pointer)

	// FramebufferDidUpdateGraphicsOrientation protocol.
	FramebufferDidUpdateGraphicsOrientation(framebuffer objectivec.IObject, orientation int64)

	// FramebufferDidUpdateColorSpace protocol.
	FramebufferDidUpdateColorSpace(space objectivec.IObject)

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
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("framebuffer:didUpdateCursor:"), framebuffer, cursor)
}
func (o VZFramebufferObserverObject) FramebufferDidUpdateFrame(framebuffer objectivec.IObject, frame unsafe.Pointer) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("framebuffer:didUpdateFrame:"), framebuffer, frame)
}
func (o VZFramebufferObserverObject) FramebufferDidUpdateGraphicsOrientation(framebuffer objectivec.IObject, orientation int64) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("framebuffer:didUpdateGraphicsOrientation:"), framebuffer, orientation)
}
func (o VZFramebufferObserverObject) FramebufferDidUpdateColorSpace(space objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("framebufferDidUpdateColorSpace:"), space)
}
func (o VZFramebufferObserverObject) GetDisplayProtectionOptions() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("getDisplayProtectionOptions"))
	return rv
}
