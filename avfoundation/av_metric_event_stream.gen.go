// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMetricEventStream] class.
var (
	_AVMetricEventStreamClass     AVMetricEventStreamClass
	_AVMetricEventStreamClassOnce sync.Once
)

func getAVMetricEventStreamClass() AVMetricEventStreamClass {
	_AVMetricEventStreamClassOnce.Do(func() {
		_AVMetricEventStreamClass = AVMetricEventStreamClass{class: objc.GetClass("AVMetricEventStream")}
	})
	return _AVMetricEventStreamClass
}

// GetAVMetricEventStreamClass returns the class object for AVMetricEventStream.
func GetAVMetricEventStreamClass() AVMetricEventStreamClass {
	return getAVMetricEventStreamClass()
}

type AVMetricEventStreamClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetricEventStreamClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetricEventStreamClass) Alloc() AVMetricEventStream {
	rv := objc.Send[AVMetricEventStream](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that allows clients to add publishers and then subscribe to
// specific metric event classes from those publishers.
//
// # Overview
// 
// Publishers are [AVFoundation] types that adopt
// [AVMetricEventStreamPublisher]. The protocol allows clients to receive
// metric events with a subscriber delegate which adopts the
// [AVMetricEventStreamSubscriber] protocol.
//
// [AVFoundation]: https://developer.apple.com/documentation/AVFoundation
//
// # Adding a publisher
//
//   - [AVMetricEventStream.AddPublisher]
//
// # Subscribing to events
//
//   - [AVMetricEventStream.SubscribeToAllMetricEvents]
//   - [AVMetricEventStream.SubscribeToMetricEvent]
//   - [AVMetricEventStream.SubscribeToMetricEvents]
//
// # Setting a subscriber
//
//   - [AVMetricEventStream.SetSubscriberQueue]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricEventStream
type AVMetricEventStream struct {
	objectivec.Object
}

// AVMetricEventStreamFromID constructs a [AVMetricEventStream] from an objc.ID.
//
// An object that allows clients to add publishers and then subscribe to
// specific metric event classes from those publishers.
func AVMetricEventStreamFromID(id objc.ID) AVMetricEventStream {
	return AVMetricEventStream{objectivec.Object{ID: id}}
}
// Ensure AVMetricEventStream implements IAVMetricEventStream.
var _ IAVMetricEventStream = AVMetricEventStream{}

// An interface definition for the [AVMetricEventStream] class.
//
// # Adding a publisher
//
//   - [IAVMetricEventStream.AddPublisher]
//
// # Subscribing to events
//
//   - [IAVMetricEventStream.SubscribeToAllMetricEvents]
//   - [IAVMetricEventStream.SubscribeToMetricEvent]
//   - [IAVMetricEventStream.SubscribeToMetricEvents]
//
// # Setting a subscriber
//
//   - [IAVMetricEventStream.SetSubscriberQueue]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricEventStream
type IAVMetricEventStream interface {
	objectivec.IObject

	// Topic: Adding a publisher

	AddPublisher(publisher AVMetricEventStreamPublisher) bool

	// Topic: Subscribing to events

	SubscribeToAllMetricEvents()
	SubscribeToMetricEvent(metricEventClass objc.Class)
	SubscribeToMetricEvents(metricEventClasses []objc.Class)

	// Topic: Setting a subscriber

	SetSubscriberQueue(subscriber AVMetricEventStreamSubscriber, queue dispatch.Queue) bool
}

// Init initializes the instance.
func (m AVMetricEventStream) Init() AVMetricEventStream {
	rv := objc.Send[AVMetricEventStream](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetricEventStream) Autorelease() AVMetricEventStream {
	rv := objc.Send[AVMetricEventStream](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetricEventStream creates a new AVMetricEventStream instance.
func NewAVMetricEventStream() AVMetricEventStream {
	class := getAVMetricEventStreamClass()
	rv := objc.Send[AVMetricEventStream](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricEventStream/addPublisher:
func (m AVMetricEventStream) AddPublisher(publisher AVMetricEventStreamPublisher) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("addPublisher:"), publisher)
	return rv
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricEventStream/subscribeToAllMetricEvents
func (m AVMetricEventStream) SubscribeToAllMetricEvents() {
	objc.Send[objc.ID](m.ID, objc.Sel("subscribeToAllMetricEvents"))
}
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricEventStream/subscribeToMetricEvent:
func (m AVMetricEventStream) SubscribeToMetricEvent(metricEventClass objc.Class) {
	objc.Send[objc.ID](m.ID, objc.Sel("subscribeToMetricEvent:"), metricEventClass)
}
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricEventStream/subscribeToMetricEvents:
func (m AVMetricEventStream) SubscribeToMetricEvents(metricEventClasses []objc.Class) {
	objc.Send[objc.ID](m.ID, objc.Sel("subscribeToMetricEvents:"), objectivec.ClassSliceToNSArray(metricEventClasses))
}
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricEventStream/setSubscriber:queue:
func (m AVMetricEventStream) SetSubscriberQueue(subscriber AVMetricEventStreamSubscriber, queue dispatch.Queue) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("setSubscriber:queue:"), subscriber, uintptr(queue.Handle()))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricEventStream/eventStream
func (_AVMetricEventStreamClass AVMetricEventStreamClass) EventStream() AVMetricEventStream {
	rv := objc.Send[objc.ID](objc.ID(_AVMetricEventStreamClass.class), objc.Sel("eventStream"))
	return AVMetricEventStreamFromID(rv)
}

