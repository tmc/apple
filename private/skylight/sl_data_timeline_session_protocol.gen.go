// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLDataTimelineSession protocol.
type SLDataTimelineSession interface {
	objectivec.IObject

	// AuditID protocol.
	AuditID() int

	// CgID protocol.
	CgID() uint32

	// CurrentSnapshotMember protocol.
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

func (o SLDataTimelineSessionObject) AuditID() int {
	rv := objc.Send[int](o.ID, objc.Sel("auditID"))
	return rv
}
func (o SLDataTimelineSessionObject) CgID() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("cgID"))
	return rv
}
func (o SLDataTimelineSessionObject) CurrentSnapshotMember() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("currentSnapshotMember"))
	return rv
}
func (o SLDataTimelineSessionObject) ProcessData() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("processData"))
	return objectivec.Object{ID: rv}
}
