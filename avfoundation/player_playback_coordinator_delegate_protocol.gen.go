// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines the methods to implement to participate in playback coordination.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerPlaybackCoordinatorDelegate
type AVPlayerPlaybackCoordinatorDelegate interface {
	objectivec.IObject
}

// AVPlayerPlaybackCoordinatorDelegateObject wraps an existing Objective-C object that conforms to the AVPlayerPlaybackCoordinatorDelegate protocol.
type AVPlayerPlaybackCoordinatorDelegateObject struct {
	objectivec.Object
}

func (o AVPlayerPlaybackCoordinatorDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVPlayerPlaybackCoordinatorDelegateObjectFromID constructs a [AVPlayerPlaybackCoordinatorDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVPlayerPlaybackCoordinatorDelegateObjectFromID(id objc.ID) AVPlayerPlaybackCoordinatorDelegateObject {
	return AVPlayerPlaybackCoordinatorDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns an identifier for a player item.
//
// coordinator: The object coordinating playback.
//
// playerItem: The player item to return an identifier for.
//
// # Return Value
//
// An identifier string.
//
// # Discussion
//
// A coordinator calls this method to identify the items that its player
// object plays.
//
// Implement this method to enable the coordinator to establish the identity
// of items that have different URLs. For example, two participants may play
// the same item, but one plays the item from a remote host and the other from
// a local version on a device.
//
// If you don’t implement this method, the coordinator derives an identifier
// from the item’s asset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerPlaybackCoordinatorDelegate/playbackCoordinator(_:identifierFor:)
func (o AVPlayerPlaybackCoordinatorDelegateObject) PlaybackCoordinatorIdentifierForPlayerItem(coordinator IAVPlayerPlaybackCoordinator, playerItem IAVPlayerItem) string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("playbackCoordinator:identifierForPlayerItem:"), coordinator, playerItem)
	return foundation.NSStringFromID(rv).String()
}

// Asks the delegate for time ranges in a player item that don’t correspond
// to the primary content.
//
// coordinator: The object coordinating playback.
//
// playerItem: The player item for which to retrieve interstitial time ranges.
//
// # Return Value
//
// An array of [NSValue] objects that contain the interstitial time ranges.
//
// # Discussion
//
// Implementing this method enables the coordinator to synchronize playback
// between participants that have different interstitials stitched into the
// primary content timeline.
//
// If you don’t implement this method, the coordinator assumes that the
// entire item corresponds to the primary content.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerPlaybackCoordinatorDelegate/playbackCoordinator(_:interstitialTimeRangesFor:)
//
// [NSValue]: https://developer.apple.com/documentation/Foundation/NSValue
func (o AVPlayerPlaybackCoordinatorDelegateObject) PlaybackCoordinatorInterstitialTimeRangesForPlayerItem(coordinator IAVPlayerPlaybackCoordinator, playerItem IAVPlayerItem) []foundation.NSValue {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("playbackCoordinator:interstitialTimeRangesForPlayerItem:"), coordinator, playerItem)
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}
