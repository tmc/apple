// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol for objects that perform timeline observations.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineObserver
type AVPlayerItemIntegratedTimelineObserver interface {
	objectivec.IObject
}

// AVPlayerItemIntegratedTimelineObserverObject wraps an existing Objective-C object that conforms to the AVPlayerItemIntegratedTimelineObserver protocol.
type AVPlayerItemIntegratedTimelineObserverObject struct {
	objectivec.Object
}
func (o AVPlayerItemIntegratedTimelineObserverObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVPlayerItemIntegratedTimelineObserverObjectFromID constructs a [AVPlayerItemIntegratedTimelineObserverObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVPlayerItemIntegratedTimelineObserverObjectFromID(id objc.ID) AVPlayerItemIntegratedTimelineObserverObject {
	return AVPlayerItemIntegratedTimelineObserverObject{
		Object: objectivec.ObjectFromID(id),
	}
}

