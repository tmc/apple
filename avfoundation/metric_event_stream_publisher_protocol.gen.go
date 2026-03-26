// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A type for objects that publish metric events to the event stream.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricEventStreamPublisher
type AVMetricEventStreamPublisher interface {
	objectivec.IObject
}

// AVMetricEventStreamPublisherObject wraps an existing Objective-C object that conforms to the AVMetricEventStreamPublisher protocol.
type AVMetricEventStreamPublisherObject struct {
	objectivec.Object
}
func (o AVMetricEventStreamPublisherObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVMetricEventStreamPublisherObjectFromID constructs a [AVMetricEventStreamPublisherObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVMetricEventStreamPublisherObjectFromID(id objc.ID) AVMetricEventStreamPublisherObject {
	return AVMetricEventStreamPublisherObject{
		Object: objectivec.ObjectFromID(id),
	}
}

