// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol you implement to receive metadata callbacks from a player item metadata collector.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataCollectorPushDelegate
type AVPlayerItemMetadataCollectorPushDelegate interface {
	objectivec.IObject

	// Tells the delegate the collected metadata group information has changed and needs to be updated.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataCollectorPushDelegate/metadataCollector(_:didCollect:indexesOfNewGroups:indexesOfModifiedGroups:)
	MetadataCollectorDidCollectDateRangeMetadataGroupsIndexesOfNewGroupsIndexesOfModifiedGroups(metadataCollector IAVPlayerItemMetadataCollector, metadataGroups []AVDateRangeMetadataGroup, indexesOfNewGroups foundation.NSIndexSet, indexesOfModifiedGroups foundation.NSIndexSet)
}

// AVPlayerItemMetadataCollectorPushDelegateObject wraps an existing Objective-C object that conforms to the AVPlayerItemMetadataCollectorPushDelegate protocol.
type AVPlayerItemMetadataCollectorPushDelegateObject struct {
	objectivec.Object
}

func (o AVPlayerItemMetadataCollectorPushDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVPlayerItemMetadataCollectorPushDelegateObjectFromID constructs a [AVPlayerItemMetadataCollectorPushDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVPlayerItemMetadataCollectorPushDelegateObjectFromID(id objc.ID) AVPlayerItemMetadataCollectorPushDelegateObject {
	return AVPlayerItemMetadataCollectorPushDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate the collected metadata group information has changed and
// needs to be updated.
//
// metadataCollector: The [AVPlayerItemMetadataCollector] on which this delegate is set.
//
// metadataGroups: The complete array of all metadata groups meeting the criteria of the
// output.
//
// indexesOfNewGroups: The indexes of the `metadataGroups` added since the last delegate
// invocation of this method.
//
// indexesOfModifiedGroups: The indexes of the `metadataGroups` modified since the last delegate
// invocation of this method.
//
// # Discussion
//
// This method is called when additions or modifications are made to the array
// of collected metadata groups. The initial invocation will have
// `indexesOfNewGroup` referring to every index in `metadataGroups`.
// Subsequent invocations may not contain all previously collected metadata
// groups if they no longer refer to a region in the player item’s
// [SeekableTimeRanges].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataCollectorPushDelegate/metadataCollector(_:didCollect:indexesOfNewGroups:indexesOfModifiedGroups:)
func (o AVPlayerItemMetadataCollectorPushDelegateObject) MetadataCollectorDidCollectDateRangeMetadataGroupsIndexesOfNewGroupsIndexesOfModifiedGroups(metadataCollector IAVPlayerItemMetadataCollector, metadataGroups []AVDateRangeMetadataGroup, indexesOfNewGroups foundation.NSIndexSet, indexesOfModifiedGroups foundation.NSIndexSet) {
	objc.Send[struct{}](o.ID, objc.Sel("metadataCollector:didCollectDateRangeMetadataGroups:indexesOfNewGroups:indexesOfModifiedGroups:"), metadataCollector, objectivec.IObjectSliceToNSArray(metadataGroups), indexesOfNewGroups, indexesOfModifiedGroups)
}
