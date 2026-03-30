// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VoiceVerificationEndpointer] class.
var (
	_VoiceVerificationEndpointerClass     VoiceVerificationEndpointerClass
	_VoiceVerificationEndpointerClassOnce sync.Once
)

func getVoiceVerificationEndpointerClass() VoiceVerificationEndpointerClass {
	_VoiceVerificationEndpointerClassOnce.Do(func() {
		_VoiceVerificationEndpointerClass = VoiceVerificationEndpointerClass{class: objc.GetClass("VoiceVerificationEndpointer")}
	})
	return _VoiceVerificationEndpointerClass
}

// GetVoiceVerificationEndpointerClass returns the class object for VoiceVerificationEndpointer.
func GetVoiceVerificationEndpointerClass() VoiceVerificationEndpointerClass {
	return getVoiceVerificationEndpointerClass()
}

type VoiceVerificationEndpointerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VoiceVerificationEndpointerClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VoiceVerificationEndpointerClass) Alloc() VoiceVerificationEndpointer {
	rv := objc.Send[VoiceVerificationEndpointer](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VoiceVerificationEndpointer.ConfigureWithASBDAndFrameRate]
//   - [VoiceVerificationEndpointer.ConfigureWithSampleRateAndFrameRate]
//   - [VoiceVerificationEndpointer.EndWaitTime]
//   - [VoiceVerificationEndpointer.SetEndWaitTime]
//   - [VoiceVerificationEndpointer.EndpointMode]
//   - [VoiceVerificationEndpointer.SetEndpointMode]
//   - [VoiceVerificationEndpointer.GetStatus]
//   - [VoiceVerificationEndpointer.InterspeechWaitTime]
//   - [VoiceVerificationEndpointer.SetInterspeechWaitTime]
//   - [VoiceVerificationEndpointer.Reset]
//   - [VoiceVerificationEndpointer.StartWaitTime]
//   - [VoiceVerificationEndpointer.SetStartWaitTime]
//   - [VoiceVerificationEndpointer.DebugDescription]
//   - [VoiceVerificationEndpointer.Description]
//   - [VoiceVerificationEndpointer.Hash]
//   - [VoiceVerificationEndpointer.Superclass]
//
// See: https://developer.apple.com/documentation/AVFAudio/VoiceVerificationEndpointer
type VoiceVerificationEndpointer struct {
	objectivec.Object
}

// VoiceVerificationEndpointerFromID constructs a [VoiceVerificationEndpointer] from an objc.ID.
func VoiceVerificationEndpointerFromID(id objc.ID) VoiceVerificationEndpointer {
	return VoiceVerificationEndpointer{objectivec.Object{ID: id}}
}

// Ensure VoiceVerificationEndpointer implements IVoiceVerificationEndpointer.
var _ IVoiceVerificationEndpointer = VoiceVerificationEndpointer{}

// An interface definition for the [VoiceVerificationEndpointer] class.
//
// # Methods
//
//   - [IVoiceVerificationEndpointer.ConfigureWithASBDAndFrameRate]
//   - [IVoiceVerificationEndpointer.ConfigureWithSampleRateAndFrameRate]
//   - [IVoiceVerificationEndpointer.EndWaitTime]
//   - [IVoiceVerificationEndpointer.SetEndWaitTime]
//   - [IVoiceVerificationEndpointer.EndpointMode]
//   - [IVoiceVerificationEndpointer.SetEndpointMode]
//   - [IVoiceVerificationEndpointer.GetStatus]
//   - [IVoiceVerificationEndpointer.InterspeechWaitTime]
//   - [IVoiceVerificationEndpointer.SetInterspeechWaitTime]
//   - [IVoiceVerificationEndpointer.Reset]
//   - [IVoiceVerificationEndpointer.StartWaitTime]
//   - [IVoiceVerificationEndpointer.SetStartWaitTime]
//   - [IVoiceVerificationEndpointer.DebugDescription]
//   - [IVoiceVerificationEndpointer.Description]
//   - [IVoiceVerificationEndpointer.Hash]
//   - [IVoiceVerificationEndpointer.Superclass]
//
// See: https://developer.apple.com/documentation/AVFAudio/VoiceVerificationEndpointer
type IVoiceVerificationEndpointer interface {
	objectivec.IObject

	// Topic: Methods

	ConfigureWithASBDAndFrameRate(asbd unsafe.Pointer, rate uint32) bool
	ConfigureWithSampleRateAndFrameRate(rate float64, rate2 uint32) bool
	EndWaitTime() float64
	SetEndWaitTime(value float64)
	EndpointMode() int
	SetEndpointMode(value int)
	GetStatus(status unsafe.Pointer) int
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
func (v VoiceVerificationEndpointer) Init() VoiceVerificationEndpointer {
	rv := objc.Send[VoiceVerificationEndpointer](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VoiceVerificationEndpointer) Autorelease() VoiceVerificationEndpointer {
	rv := objc.Send[VoiceVerificationEndpointer](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVoiceVerificationEndpointer creates a new VoiceVerificationEndpointer instance.
func NewVoiceVerificationEndpointer() VoiceVerificationEndpointer {
	class := getVoiceVerificationEndpointerClass()
	rv := objc.Send[VoiceVerificationEndpointer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/VoiceVerificationEndpointer/configureWithASBD:andFrameRate:
func (v VoiceVerificationEndpointer) ConfigureWithASBDAndFrameRate(asbd unsafe.Pointer, rate uint32) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("configureWithASBD:andFrameRate:"), asbd, rate)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/VoiceVerificationEndpointer/configureWithSampleRate:andFrameRate:
func (v VoiceVerificationEndpointer) ConfigureWithSampleRateAndFrameRate(rate float64, rate2 uint32) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("configureWithSampleRate:andFrameRate:"), rate, rate2)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/VoiceVerificationEndpointer/getStatus:
func (v VoiceVerificationEndpointer) GetStatus(status unsafe.Pointer) int {
	rv := objc.Send[int](v.ID, objc.Sel("getStatus:"), status)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/VoiceVerificationEndpointer/reset
func (v VoiceVerificationEndpointer) Reset() {
	objc.Send[objc.ID](v.ID, objc.Sel("reset"))
}

// See: https://developer.apple.com/documentation/AVFAudio/VoiceVerificationEndpointer/debugDescription
func (v VoiceVerificationEndpointer) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/AVFAudio/VoiceVerificationEndpointer/description
func (v VoiceVerificationEndpointer) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/AVFAudio/VoiceVerificationEndpointer/endWaitTime
func (v VoiceVerificationEndpointer) EndWaitTime() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("endWaitTime"))
	return rv
}
func (v VoiceVerificationEndpointer) SetEndWaitTime(value float64) {
	objc.Send[struct{}](v.ID, objc.Sel("setEndWaitTime:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/VoiceVerificationEndpointer/endpointMode
func (v VoiceVerificationEndpointer) EndpointMode() int {
	rv := objc.Send[int](v.ID, objc.Sel("endpointMode"))
	return rv
}
func (v VoiceVerificationEndpointer) SetEndpointMode(value int) {
	objc.Send[struct{}](v.ID, objc.Sel("setEndpointMode:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/VoiceVerificationEndpointer/hash
func (v VoiceVerificationEndpointer) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/VoiceVerificationEndpointer/interspeechWaitTime
func (v VoiceVerificationEndpointer) InterspeechWaitTime() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("interspeechWaitTime"))
	return rv
}
func (v VoiceVerificationEndpointer) SetInterspeechWaitTime(value float64) {
	objc.Send[struct{}](v.ID, objc.Sel("setInterspeechWaitTime:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/VoiceVerificationEndpointer/startWaitTime
func (v VoiceVerificationEndpointer) StartWaitTime() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("startWaitTime"))
	return rv
}
func (v VoiceVerificationEndpointer) SetStartWaitTime(value float64) {
	objc.Send[struct{}](v.ID, objc.Sel("setStartWaitTime:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/VoiceVerificationEndpointer/superclass
func (v VoiceVerificationEndpointer) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}
