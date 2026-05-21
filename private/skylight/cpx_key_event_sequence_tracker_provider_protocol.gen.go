// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXKeyEventSequenceTrackerProvider protocol.
type CPXKeyEventSequenceTrackerProvider interface {
	objectivec.IObject

	// CurrentRegionID protocol.
	CurrentRegionID() uint64

	// EventLimit protocol.
	EventLimit() uint64

	// MainDisplayHeight protocol.
	MainDisplayHeight() uint16

	// StructuralRegionForID protocol.
	StructuralRegionForID(id uint64) WSStructuralRegionRef

	// WindowByID protocol.
	WindowByID(id uint32) unsafe.Pointer

	// WindowHeightForWindow protocol.
	WindowHeightForWindow(window unsafe.Pointer) uint16
}

// CPXKeyEventSequenceTrackerProviderObject wraps an existing Objective-C object that conforms to the CPXKeyEventSequenceTrackerProvider protocol.
type CPXKeyEventSequenceTrackerProviderObject struct {
	objectivec.Object
}

func (o CPXKeyEventSequenceTrackerProviderObject) BaseObject() objectivec.Object {
	return o.Object
}

// CPXKeyEventSequenceTrackerProviderObjectFromID constructs a [CPXKeyEventSequenceTrackerProviderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CPXKeyEventSequenceTrackerProviderObjectFromID(id objc.ID) CPXKeyEventSequenceTrackerProviderObject {
	return CPXKeyEventSequenceTrackerProviderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o CPXKeyEventSequenceTrackerProviderObject) CurrentRegionID() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("currentRegionID"))
	return rv
}
func (o CPXKeyEventSequenceTrackerProviderObject) EventLimit() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("eventLimit"))
	return rv
}
func (o CPXKeyEventSequenceTrackerProviderObject) MainDisplayHeight() uint16 {
	rv := objc.Send[uint16](o.ID, objc.Sel("mainDisplayHeight"))
	return rv
}
func (o CPXKeyEventSequenceTrackerProviderObject) StructuralRegionForID(id uint64) WSStructuralRegionRef {
	rv := objc.Send[WSStructuralRegionRef](o.ID, objc.Sel("structuralRegionForID:"), id)
	return rv
}
func (o CPXKeyEventSequenceTrackerProviderObject) WindowByID(id uint32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("windowByID:"), id)
	return rv
}
func (o CPXKeyEventSequenceTrackerProviderObject) WindowHeightForWindow(window unsafe.Pointer) uint16 {
	rv := objc.Send[uint16](o.ID, objc.Sel("windowHeightForWindow:"), window)
	return rv
}
