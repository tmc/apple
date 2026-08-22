// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLDataTimelineServerSnapshot protocol.
type SLDataTimelineServerSnapshot interface {
	objectivec.IObject

	// Index protocol.
	Index() uint64

	// Sessions protocol.
	Sessions() objectivec.IObject

	// Timestamp protocol.
	Timestamp() float64
}

// SLDataTimelineServerSnapshotObject wraps an existing Objective-C object that conforms to the SLDataTimelineServerSnapshot protocol.
type SLDataTimelineServerSnapshotObject struct {
	objectivec.Object
}

func (o SLDataTimelineServerSnapshotObject) BaseObject() objectivec.Object {
	return o.Object
}

// SLDataTimelineServerSnapshotObjectFromID constructs a [SLDataTimelineServerSnapshotObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SLDataTimelineServerSnapshotObjectFromID(id objc.ID) SLDataTimelineServerSnapshotObject {
	return SLDataTimelineServerSnapshotObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o SLDataTimelineServerSnapshotObject) Index() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("index"))
	return rv
}
func (o SLDataTimelineServerSnapshotObject) Sessions() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("sessions"))
	return objectivec.Object{ID: rv}
}
func (o SLDataTimelineServerSnapshotObject) Timestamp() float64 {
	rv := objc.SendIfResponds[float64](o.ID, objc.Sel("timestamp"))
	return rv
}
