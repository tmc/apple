// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// VZGraphicsDisplayInternalObserver protocol.
type VZGraphicsDisplayInternalObserver interface {
	objectivec.IObject

	// DisplayDidUpdateColorSpace protocol.
	DisplayDidUpdateColorSpace(display objectivec.IObject, space unsafe.Pointer)

	// DisplayDidUpdateHeadroomRequest protocol.
	DisplayDidUpdateHeadroomRequest(display objectivec.IObject, request float64)

	// DisplayDidUpdateOrientation protocol.
	DisplayDidUpdateOrientation(display objectivec.IObject, orientation int64)
}

// VZGraphicsDisplayInternalObserverObject wraps an existing Objective-C object that conforms to the VZGraphicsDisplayInternalObserver protocol.
type VZGraphicsDisplayInternalObserverObject struct {
	objectivec.Object
}

func (o VZGraphicsDisplayInternalObserverObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZGraphicsDisplayInternalObserverObjectFromID constructs a [VZGraphicsDisplayInternalObserverObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZGraphicsDisplayInternalObserverObjectFromID(id objc.ID) VZGraphicsDisplayInternalObserverObject {
	return VZGraphicsDisplayInternalObserverObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o VZGraphicsDisplayInternalObserverObject) DisplayDidUpdateColorSpace(display objectivec.IObject, space unsafe.Pointer) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("display:didUpdateColorSpace:"), display, space)
}
func (o VZGraphicsDisplayInternalObserverObject) DisplayDidUpdateHeadroomRequest(display objectivec.IObject, request float64) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("display:didUpdateHeadroomRequest:"), display, request)
}
func (o VZGraphicsDisplayInternalObserverObject) DisplayDidUpdateOrientation(display objectivec.IObject, orientation int64) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("display:didUpdateOrientation:"), display, orientation)
}
