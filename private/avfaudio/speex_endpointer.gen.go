// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SpeexEndpointer] class.
var (
	_SpeexEndpointerClass     SpeexEndpointerClass
	_SpeexEndpointerClassOnce sync.Once
)

func getSpeexEndpointerClass() SpeexEndpointerClass {
	_SpeexEndpointerClassOnce.Do(func() {
		_SpeexEndpointerClass = SpeexEndpointerClass{class: objc.GetClass("SpeexEndpointer")}
	})
	return _SpeexEndpointerClass
}

// GetSpeexEndpointerClass returns the class object for SpeexEndpointer.
func GetSpeexEndpointerClass() SpeexEndpointerClass {
	return getSpeexEndpointerClass()
}

type SpeexEndpointerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SpeexEndpointerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SpeexEndpointerClass) Alloc() SpeexEndpointer {
	rv := objc.Send[SpeexEndpointer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [SpeexEndpointer.ConfigureWithASBDAndFrameRate]
//   - [SpeexEndpointer.ConfigureWithSampleRateAndFrameRate]
//   - [SpeexEndpointer.EndWaitTime]
//   - [SpeexEndpointer.SetEndWaitTime]
//   - [SpeexEndpointer.EndpointMode]
//   - [SpeexEndpointer.SetEndpointMode]
//   - [SpeexEndpointer.GetStatus]
//   - [SpeexEndpointer.GetStatusCount]
//   - [SpeexEndpointer.InterspeechWaitTime]
//   - [SpeexEndpointer.SetInterspeechWaitTime]
//   - [SpeexEndpointer.Reset]
//   - [SpeexEndpointer.StartWaitTime]
//   - [SpeexEndpointer.SetStartWaitTime]
//   - [SpeexEndpointer.DebugDescription]
//   - [SpeexEndpointer.Description]
//   - [SpeexEndpointer.Hash]
//   - [SpeexEndpointer.Superclass]
// See: https://developer.apple.com/documentation/AVFAudio/SpeexEndpointer
type SpeexEndpointer struct {
	objectivec.Object
}

// SpeexEndpointerFromID constructs a [SpeexEndpointer] from an objc.ID.
func SpeexEndpointerFromID(id objc.ID) SpeexEndpointer {
	return SpeexEndpointer{objectivec.Object{ID: id}}
}
// Ensure SpeexEndpointer implements ISpeexEndpointer.
var _ ISpeexEndpointer = SpeexEndpointer{}

// An interface definition for the [SpeexEndpointer] class.
//
// # Methods
//
//   - [ISpeexEndpointer.ConfigureWithASBDAndFrameRate]
//   - [ISpeexEndpointer.ConfigureWithSampleRateAndFrameRate]
//   - [ISpeexEndpointer.EndWaitTime]
//   - [ISpeexEndpointer.SetEndWaitTime]
//   - [ISpeexEndpointer.EndpointMode]
//   - [ISpeexEndpointer.SetEndpointMode]
//   - [ISpeexEndpointer.GetStatus]
//   - [ISpeexEndpointer.GetStatusCount]
//   - [ISpeexEndpointer.InterspeechWaitTime]
//   - [ISpeexEndpointer.SetInterspeechWaitTime]
//   - [ISpeexEndpointer.Reset]
//   - [ISpeexEndpointer.StartWaitTime]
//   - [ISpeexEndpointer.SetStartWaitTime]
//   - [ISpeexEndpointer.DebugDescription]
//   - [ISpeexEndpointer.Description]
//   - [ISpeexEndpointer.Hash]
//   - [ISpeexEndpointer.Superclass]
//
// See: https://developer.apple.com/documentation/AVFAudio/SpeexEndpointer
type ISpeexEndpointer interface {
	objectivec.IObject

	// Topic: Methods

	ConfigureWithASBDAndFrameRate(asbd unsafe.Pointer, rate uint32) bool
	ConfigureWithSampleRateAndFrameRate(rate float64, rate2 uint32) bool
	EndWaitTime() float64
	SetEndWaitTime(value float64)
	EndpointMode() int
	SetEndpointMode(value int)
	GetStatus(status unsafe.Pointer) int
	GetStatusCount(status []float32, count uint32) int
	InterspeechWaitTime() float64
	SetInterspeechWaitTime(value float64)
	Reset()
	StartWaitTime() float64
	SetStartWaitTime(value float64)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (s SpeexEndpointer) Init() SpeexEndpointer {
	rv := objc.Send[SpeexEndpointer](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SpeexEndpointer) Autorelease() SpeexEndpointer {
	rv := objc.Send[SpeexEndpointer](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSpeexEndpointer creates a new SpeexEndpointer instance.
func NewSpeexEndpointer() SpeexEndpointer {
	class := getSpeexEndpointerClass()
	rv := objc.Send[SpeexEndpointer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/SpeexEndpointer/configureWithASBD:andFrameRate:
func (s SpeexEndpointer) ConfigureWithASBDAndFrameRate(asbd unsafe.Pointer, rate uint32) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("configureWithASBD:andFrameRate:"), asbd, rate)
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/SpeexEndpointer/configureWithSampleRate:andFrameRate:
func (s SpeexEndpointer) ConfigureWithSampleRateAndFrameRate(rate float64, rate2 uint32) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("configureWithSampleRate:andFrameRate:"), rate, rate2)
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/SpeexEndpointer/getStatus:
func (s SpeexEndpointer) GetStatus(status unsafe.Pointer) int {
	rv := objc.Send[int](s.ID, objc.Sel("getStatus:"), status)
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/SpeexEndpointer/getStatus:count:
func (s SpeexEndpointer) GetStatusCount(status []float32, count uint32) int {
	rv := objc.Send[int](s.ID, objc.Sel("getStatus:count:"), objc.CArray(status), count)
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/SpeexEndpointer/reset
func (s SpeexEndpointer) Reset() {
	objc.Send[objc.ID](s.ID, objc.Sel("reset"))
}

// See: https://developer.apple.com/documentation/AVFAudio/SpeexEndpointer/debugDescription
func (s SpeexEndpointer) DebugDescription() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/AVFAudio/SpeexEndpointer/description
func (s SpeexEndpointer) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/AVFAudio/SpeexEndpointer/endWaitTime
func (s SpeexEndpointer) EndWaitTime() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("endWaitTime"))
	return rv
}
func (s SpeexEndpointer) SetEndWaitTime(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setEndWaitTime:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/SpeexEndpointer/endpointMode
func (s SpeexEndpointer) EndpointMode() int {
	rv := objc.Send[int](s.ID, objc.Sel("endpointMode"))
	return rv
}
func (s SpeexEndpointer) SetEndpointMode(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setEndpointMode:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/SpeexEndpointer/hash
func (s SpeexEndpointer) Hash() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/SpeexEndpointer/interspeechWaitTime
func (s SpeexEndpointer) InterspeechWaitTime() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("interspeechWaitTime"))
	return rv
}
func (s SpeexEndpointer) SetInterspeechWaitTime(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setInterspeechWaitTime:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/SpeexEndpointer/startWaitTime
func (s SpeexEndpointer) StartWaitTime() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("startWaitTime"))
	return rv
}
func (s SpeexEndpointer) SetStartWaitTime(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setStartWaitTime:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/SpeexEndpointer/superclass
func (s SpeexEndpointer) Superclass() objc.Class {
	rv := objc.Send[objc.Class](s.ID, objc.Sel("superclass"))
	return rv
}

