// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _VZDisplayPresenterObserver protocol.
type VZDisplayPresenterObserver interface {
	objectivec.IObject

	// PresenterDidUpdateContentHeadroom protocol.
	PresenterDidUpdateContentHeadroom(presenter objectivec.IObject, headroom float64)

	// PresenterDidUpdateCursor protocol.
	PresenterDidUpdateCursor(presenter objectivec.IObject, cursor unsafe.Pointer)

	// PresenterDidUpdateFrame protocol.
	PresenterDidUpdateFrame(presenter objectivec.IObject, frame unsafe.Pointer)

	// PresenterDidUpdateHostDisplay protocol.
	PresenterDidUpdateHostDisplay(presenter objectivec.IObject, display HostDisplayUpdate)
}

// VZDisplayPresenterObserverObject wraps an existing Objective-C object that conforms to the VZDisplayPresenterObserver protocol.
type VZDisplayPresenterObserverObject struct {
	objectivec.Object
}

func (o VZDisplayPresenterObserverObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZDisplayPresenterObserverObjectFromID constructs a [VZDisplayPresenterObserverObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZDisplayPresenterObserverObjectFromID(id objc.ID) VZDisplayPresenterObserverObject {
	return VZDisplayPresenterObserverObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o VZDisplayPresenterObserverObject) PresenterDidUpdateContentHeadroom(presenter objectivec.IObject, headroom float64) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("presenter:didUpdateContentHeadroom:"), presenter, headroom)
}
func (o VZDisplayPresenterObserverObject) PresenterDidUpdateCursor(presenter objectivec.IObject, cursor unsafe.Pointer) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("presenter:didUpdateCursor:"), presenter, cursor)
}
func (o VZDisplayPresenterObserverObject) PresenterDidUpdateFrame(presenter objectivec.IObject, frame unsafe.Pointer) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("presenter:didUpdateFrame:"), presenter, frame)
}
func (o VZDisplayPresenterObserverObject) PresenterDidUpdateHostDisplay(presenter objectivec.IObject, display HostDisplayUpdate) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("presenter:didUpdateHostDisplay:"), presenter, display)
}
