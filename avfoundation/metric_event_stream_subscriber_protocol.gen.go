// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A type for objects that receive metric events.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricEventStreamSubscriber
type AVMetricEventStreamSubscriber interface {
	objectivec.IObject

	// Tells the subscriber that the publisher produced a metric event.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetricEventStreamSubscriber/publisher:didReceiveEvent:
	PublisherDidReceiveEvent(publisher AVMetricEventStreamPublisher, event IAVMetricEvent)
}

// AVMetricEventStreamSubscriberObject wraps an existing Objective-C object that conforms to the AVMetricEventStreamSubscriber protocol.
type AVMetricEventStreamSubscriberObject struct {
	objectivec.Object
}
func (o AVMetricEventStreamSubscriberObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVMetricEventStreamSubscriberObjectFromID constructs a [AVMetricEventStreamSubscriberObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVMetricEventStreamSubscriberObjectFromID(id objc.ID) AVMetricEventStreamSubscriberObject {
	return AVMetricEventStreamSubscriberObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the subscriber that the publisher produced a metric event.
//
// publisher: The publisher.
//
// event: The metric.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricEventStreamSubscriber/publisher:didReceiveEvent:
func (o AVMetricEventStreamSubscriberObject) PublisherDidReceiveEvent(publisher AVMetricEventStreamPublisher, event IAVMetricEvent) {
	objc.Send[struct{}](o.ID, objc.Sel("publisher:didReceiveEvent:"), publisher, event)
	}

