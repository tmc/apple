// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLDataTimelineSessionProcessCollection protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionProcessCollection
type SLDataTimelineSessionProcessCollection interface {
	objectivec.IObject

	// ForegroundAppPID protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionProcessCollection/foregroundAppPID
	ForegroundAppPID() int

	// ProcessesApplyBlock protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionProcessCollection/processesApplyBlock:
	ProcessesApplyBlock(block SLDataTimelineProcessHandler)

	// SessionSnapshotIndex protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionProcessCollection/sessionSnapshotIndex
	SessionSnapshotIndex() uint64

	// SessionSnapshotTimestamp protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionProcessCollection/sessionSnapshotTimestamp
	SessionSnapshotTimestamp() float64
}

// SLDataTimelineSessionProcessCollectionObject wraps an existing Objective-C object that conforms to the SLDataTimelineSessionProcessCollection protocol.
type SLDataTimelineSessionProcessCollectionObject struct {
	objectivec.Object
}

func (o SLDataTimelineSessionProcessCollectionObject) BaseObject() objectivec.Object {
	return o.Object
}

// SLDataTimelineSessionProcessCollectionObjectFromID constructs a [SLDataTimelineSessionProcessCollectionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SLDataTimelineSessionProcessCollectionObjectFromID(id objc.ID) SLDataTimelineSessionProcessCollectionObject {
	return SLDataTimelineSessionProcessCollectionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionProcessCollection/foregroundAppPID
func (o SLDataTimelineSessionProcessCollectionObject) ForegroundAppPID() int {
	rv := objc.Send[int](o.ID, objc.Sel("foregroundAppPID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionProcessCollection/processes
func (o SLDataTimelineSessionProcessCollectionObject) Processes() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("processes"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionProcessCollection/processesApplyBlock:
func (o SLDataTimelineSessionProcessCollectionObject) ProcessesApplyBlock(block SLDataTimelineProcessHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("processesApplyBlock:"), block)
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionProcessCollection/sessionSnapshotIndex
func (o SLDataTimelineSessionProcessCollectionObject) SessionSnapshotIndex() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("sessionSnapshotIndex"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSessionProcessCollection/sessionSnapshotTimestamp
func (o SLDataTimelineSessionProcessCollectionObject) SessionSnapshotTimestamp() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("sessionSnapshotTimestamp"))
	return rv
}
