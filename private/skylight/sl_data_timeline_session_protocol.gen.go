// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLDataTimelineSession protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSession
type SLDataTimelineSession interface {
	objectivec.IObject

	// AuditID protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSession/auditID
	AuditID() int

	// CgID protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSession/cgID
	CgID() uint32

	// CurrentSnapshotMember protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSession/currentSnapshotMember
	CurrentSnapshotMember() bool
}

// SLDataTimelineSessionObject wraps an existing Objective-C object that conforms to the SLDataTimelineSession protocol.
type SLDataTimelineSessionObject struct {
	objectivec.Object
}

func (o SLDataTimelineSessionObject) BaseObject() objectivec.Object {
	return o.Object
}

// SLDataTimelineSessionObjectFromID constructs a [SLDataTimelineSessionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SLDataTimelineSessionObjectFromID(id objc.ID) SLDataTimelineSessionObject {
	return SLDataTimelineSessionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSession/auditID
func (o SLDataTimelineSessionObject) AuditID() int {
	rv := objc.Send[int](o.ID, objc.Sel("auditID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSession/cgID
func (o SLDataTimelineSessionObject) CgID() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("cgID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSession/currentSnapshotMember
func (o SLDataTimelineSessionObject) CurrentSnapshotMember() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("currentSnapshotMember"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSession/processData
func (o SLDataTimelineSessionObject) ProcessData() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("processData"))
	return objectivec.Object{ID: rv}
}
