// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"github.com/tmc/apple/coreaudiotypes"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Endpointer protocol.
type Endpointer interface {
	objectivec.IObject

	// ConfigureWithASBDAndFrameRate protocol.
	ConfigureWithASBDAndFrameRate(asbd coreaudiotypes.AudioStreamBasicDescription, rate uint32) bool

	// ConfigureWithSampleRateAndFrameRate protocol.
	ConfigureWithSampleRateAndFrameRate(rate float64, rate2 uint32) bool

	// EndWaitTime protocol.
	EndWaitTime() float64

	// EndpointMode protocol.
	EndpointMode() int

	// GetStatus protocol.
	GetStatus(status AudioQueueBuffer) int

	// InterspeechWaitTime protocol.
	InterspeechWaitTime() float64

	// Reset protocol.
	Reset()

	// SetEndWaitTime protocol.
	SetEndWaitTime(time float64)

	// SetEndpointMode protocol.
	SetEndpointMode(mode int)

	// SetInterspeechWaitTime protocol.
	SetInterspeechWaitTime(time float64)

	// SetStartWaitTime protocol.
	SetStartWaitTime(time float64)

	// StartWaitTime protocol.
	StartWaitTime() float64
}

// EndpointerObject wraps an existing Objective-C object that conforms to the Endpointer protocol.
type EndpointerObject struct {
	objectivec.Object
}

func (o EndpointerObject) BaseObject() objectivec.Object {
	return o.Object
}

// EndpointerObjectFromID constructs a [EndpointerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func EndpointerObjectFromID(id objc.ID) EndpointerObject {
	return EndpointerObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o EndpointerObject) ConfigureWithASBDAndFrameRate(asbd coreaudiotypes.AudioStreamBasicDescription, rate uint32) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("configureWithASBD:andFrameRate:"), asbd, rate)
	return rv
}
func (o EndpointerObject) ConfigureWithSampleRateAndFrameRate(rate float64, rate2 uint32) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("configureWithSampleRate:andFrameRate:"), rate, rate2)
	return rv
}
func (o EndpointerObject) EndWaitTime() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("endWaitTime"))
	return rv
}
func (o EndpointerObject) EndpointMode() int {
	rv := objc.Send[int](o.ID, objc.Sel("endpointMode"))
	return rv
}
func (o EndpointerObject) GetStatus(status AudioQueueBuffer) int {
	rv := objc.Send[int](o.ID, objc.Sel("getStatus:"), status)
	return rv
}
func (o EndpointerObject) InterspeechWaitTime() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("interspeechWaitTime"))
	return rv
}
func (o EndpointerObject) Reset() {
	objc.Send[struct{}](o.ID, objc.Sel("reset"))
}
func (o EndpointerObject) SetEndWaitTime(time float64) {
	objc.Send[struct{}](o.ID, objc.Sel("setEndWaitTime:"), time)
}
func (o EndpointerObject) SetEndpointMode(mode int) {
	objc.Send[struct{}](o.ID, objc.Sel("setEndpointMode:"), mode)
}
func (o EndpointerObject) SetInterspeechWaitTime(time float64) {
	objc.Send[struct{}](o.ID, objc.Sel("setInterspeechWaitTime:"), time)
}
func (o EndpointerObject) SetStartWaitTime(time float64) {
	objc.Send[struct{}](o.ID, objc.Sel("setStartWaitTime:"), time)
}
func (o EndpointerObject) StartWaitTime() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("startWaitTime"))
	return rv
}
