// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLDataTimelineUpdate protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineUpdate
type SLDataTimelineUpdate interface {
	objectivec.IObject

	// Action protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineUpdate/action
	Action() uint32
}

// SLDataTimelineUpdateObject wraps an existing Objective-C object that conforms to the SLDataTimelineUpdate protocol.
type SLDataTimelineUpdateObject struct {
	objectivec.Object
}

func (o SLDataTimelineUpdateObject) BaseObject() objectivec.Object {
	return o.Object
}

// SLDataTimelineUpdateObjectFromID constructs a [SLDataTimelineUpdateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SLDataTimelineUpdateObjectFromID(id objc.ID) SLDataTimelineUpdateObject {
	return SLDataTimelineUpdateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineUpdate/action
func (o SLDataTimelineUpdateObject) Action() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("action"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineUpdate/claimSnapshotCollection
func (o SLDataTimelineUpdateObject) ClaimSnapshotCollection() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("claimSnapshotCollection"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineUpdate/connection
func (o SLDataTimelineUpdateObject) Connection() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("connection"))
	return objectivec.Object{ID: rv}
}
