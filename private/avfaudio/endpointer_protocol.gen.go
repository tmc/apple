// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Endpointer protocol.
//
// See: https://developer.apple.com/documentation/AVFAudio/Endpointer
type Endpointer interface {
	objectivec.IObject

	// ConfigureWithASBDAndFrameRate protocol.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/Endpointer/configureWithASBD:andFrameRate:
	ConfigureWithASBDAndFrameRate(asbd unsafe.Pointer, rate uint32) bool

	// ConfigureWithSampleRateAndFrameRate protocol.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/Endpointer/configureWithSampleRate:andFrameRate:
	ConfigureWithSampleRateAndFrameRate(rate float64, rate2 uint32) bool

	// EndWaitTime protocol.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/Endpointer/endWaitTime
	EndWaitTime() float64

	// EndpointMode protocol.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/Endpointer/endpointMode
	EndpointMode() int

	// GetStatus protocol.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/Endpointer/getStatus:
	GetStatus(status unsafe.Pointer) int

	// InterspeechWaitTime protocol.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/Endpointer/interspeechWaitTime
	InterspeechWaitTime() float64

	// Reset protocol.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/Endpointer/reset
	Reset()

	// SetEndWaitTime protocol.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/Endpointer/setEndWaitTime:
	SetEndWaitTime(time float64)

	// SetEndpointMode protocol.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/Endpointer/setEndpointMode:
	SetEndpointMode(mode int)

	// SetInterspeechWaitTime protocol.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/Endpointer/setInterspeechWaitTime:
	SetInterspeechWaitTime(time float64)

	// SetStartWaitTime protocol.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/Endpointer/setStartWaitTime:
	SetStartWaitTime(time float64)

	// StartWaitTime protocol.
	//
	// See: https://developer.apple.com/documentation/AVFAudio/Endpointer/startWaitTime
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

// See: https://developer.apple.com/documentation/AVFAudio/Endpointer/configureWithASBD:andFrameRate:
func (o EndpointerObject) ConfigureWithASBDAndFrameRate(asbd unsafe.Pointer, rate uint32) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("configureWithASBD:andFrameRate:"), asbd, rate)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/Endpointer/configureWithSampleRate:andFrameRate:
func (o EndpointerObject) ConfigureWithSampleRateAndFrameRate(rate float64, rate2 uint32) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("configureWithSampleRate:andFrameRate:"), rate, rate2)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/Endpointer/endWaitTime
func (o EndpointerObject) EndWaitTime() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("endWaitTime"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/Endpointer/endpointMode
func (o EndpointerObject) EndpointMode() int {
	rv := objc.Send[int](o.ID, objc.Sel("endpointMode"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/Endpointer/getStatus:
func (o EndpointerObject) GetStatus(status unsafe.Pointer) int {
	rv := objc.Send[int](o.ID, objc.Sel("getStatus:"), status)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/Endpointer/interspeechWaitTime
func (o EndpointerObject) InterspeechWaitTime() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("interspeechWaitTime"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/Endpointer/reset
func (o EndpointerObject) Reset() {
	objc.Send[struct{}](o.ID, objc.Sel("reset"))
}

// See: https://developer.apple.com/documentation/AVFAudio/Endpointer/setEndWaitTime:
func (o EndpointerObject) SetEndWaitTime(time float64) {
	objc.Send[struct{}](o.ID, objc.Sel("setEndWaitTime:"), time)
}

// See: https://developer.apple.com/documentation/AVFAudio/Endpointer/setEndpointMode:
func (o EndpointerObject) SetEndpointMode(mode int) {
	objc.Send[struct{}](o.ID, objc.Sel("setEndpointMode:"), mode)
}

// See: https://developer.apple.com/documentation/AVFAudio/Endpointer/setInterspeechWaitTime:
func (o EndpointerObject) SetInterspeechWaitTime(time float64) {
	objc.Send[struct{}](o.ID, objc.Sel("setInterspeechWaitTime:"), time)
}

// See: https://developer.apple.com/documentation/AVFAudio/Endpointer/setStartWaitTime:
func (o EndpointerObject) SetStartWaitTime(time float64) {
	objc.Send[struct{}](o.ID, objc.Sel("setStartWaitTime:"), time)
}

// See: https://developer.apple.com/documentation/AVFAudio/Endpointer/startWaitTime
func (o EndpointerObject) StartWaitTime() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("startWaitTime"))
	return rv
}
