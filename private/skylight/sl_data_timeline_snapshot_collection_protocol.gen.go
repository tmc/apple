// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLDataTimelineSnapshotCollection protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSnapshotCollection
type SLDataTimelineSnapshotCollection interface {
	objectivec.IObject

	// SnapshotsApplyBlock protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSnapshotCollection/snapshotsApplyBlock:
	SnapshotsApplyBlock(block SLDataTimelineServerSnapshotHandler)
}

// SLDataTimelineSnapshotCollectionObject wraps an existing Objective-C object that conforms to the SLDataTimelineSnapshotCollection protocol.
type SLDataTimelineSnapshotCollectionObject struct {
	objectivec.Object
}

func (o SLDataTimelineSnapshotCollectionObject) BaseObject() objectivec.Object {
	return o.Object
}

// SLDataTimelineSnapshotCollectionObjectFromID constructs a [SLDataTimelineSnapshotCollectionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SLDataTimelineSnapshotCollectionObjectFromID(id objc.ID) SLDataTimelineSnapshotCollectionObject {
	return SLDataTimelineSnapshotCollectionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSnapshotCollection/snapshots
func (o SLDataTimelineSnapshotCollectionObject) Snapshots() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("snapshots"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineSnapshotCollection/snapshotsApplyBlock:
func (o SLDataTimelineSnapshotCollectionObject) SnapshotsApplyBlock(block SLDataTimelineServerSnapshotHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("snapshotsApplyBlock:"), block)
}
