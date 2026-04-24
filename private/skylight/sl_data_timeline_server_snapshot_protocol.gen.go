// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLDataTimelineServerSnapshot protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineServerSnapshot
type SLDataTimelineServerSnapshot interface {
	objectivec.IObject

	// Index protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineServerSnapshot/index
	Index() uint64

	// SessionsApplyBlock protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineServerSnapshot/sessionsApplyBlock:
	SessionsApplyBlock(block SLDataTimelineSessionHandler)

	// Timestamp protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineServerSnapshot/timestamp
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

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineServerSnapshot/index
func (o SLDataTimelineServerSnapshotObject) Index() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("index"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineServerSnapshot/sessions
func (o SLDataTimelineServerSnapshotObject) Sessions() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sessions"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineServerSnapshot/sessionsApplyBlock:
func (o SLDataTimelineServerSnapshotObject) SessionsApplyBlock(block SLDataTimelineSessionHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("sessionsApplyBlock:"), block)
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineServerSnapshot/timestamp
func (o SLDataTimelineServerSnapshotObject) Timestamp() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("timestamp"))
	return rv
}
