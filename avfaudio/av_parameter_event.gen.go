// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVParameterEvent] class.
var (
	_AVParameterEventClass     AVParameterEventClass
	_AVParameterEventClassOnce sync.Once
)

func getAVParameterEventClass() AVParameterEventClass {
	_AVParameterEventClassOnce.Do(func() {
		_AVParameterEventClass = AVParameterEventClass{class: objc.GetClass("AVParameterEvent")}
	})
	return _AVParameterEventClass
}

// GetAVParameterEventClass returns the class object for AVParameterEvent.
func GetAVParameterEventClass() AVParameterEventClass {
	return getAVParameterEventClass()
}

type AVParameterEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVParameterEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVParameterEventClass) Alloc() AVParameterEvent {
	rv := objc.Send[AVParameterEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a parameter event on a music track’s
// destination.
//
// # Overview
//
// When you configure an audio unit as the destination for an [AVMusicTrack]
// that contains this event, you can schedule and automate parameter changes.
//
// When the track is playing as part of a sequence, the destination audio unit
// receives set-parameter messages whose values change smoothly along a linear
// ramp between each event’s beat location.
//
// If you add an event to an empty, non-automation track, the track becomes an
// automation track.
//
// # Creating a Parameter Event
//
//   - [AVParameterEvent.InitWithParameterIDScopeElementValue]: Creates an event with a parameter identifier, scope, element, and value for the parameter to set.
//
// # Configuring a Parameter Event
//
//   - [AVParameterEvent.ParameterID]: The identifier of the parameter.
//   - [AVParameterEvent.SetParameterID]
//   - [AVParameterEvent.Scope]: The audio unit scope for the parameter.
//   - [AVParameterEvent.SetScope]
//   - [AVParameterEvent.Element]: The element index in the scope.
//   - [AVParameterEvent.SetElement]
//   - [AVParameterEvent.Value]: The value of the parameter to set.
//   - [AVParameterEvent.SetValue]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVParameterEvent
type AVParameterEvent struct {
	AVMusicEvent
}

// AVParameterEventFromID constructs a [AVParameterEvent] from an objc.ID.
//
// An object that represents a parameter event on a music track’s
// destination.
func AVParameterEventFromID(id objc.ID) AVParameterEvent {
	return AVParameterEvent{AVMusicEvent: AVMusicEventFromID(id)}
}

// NOTE: AVParameterEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVParameterEvent] class.
//
// # Creating a Parameter Event
//
//   - [IAVParameterEvent.InitWithParameterIDScopeElementValue]: Creates an event with a parameter identifier, scope, element, and value for the parameter to set.
//
// # Configuring a Parameter Event
//
//   - [IAVParameterEvent.ParameterID]: The identifier of the parameter.
//   - [IAVParameterEvent.SetParameterID]
//   - [IAVParameterEvent.Scope]: The audio unit scope for the parameter.
//   - [IAVParameterEvent.SetScope]
//   - [IAVParameterEvent.Element]: The element index in the scope.
//   - [IAVParameterEvent.SetElement]
//   - [IAVParameterEvent.Value]: The value of the parameter to set.
//   - [IAVParameterEvent.SetValue]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVParameterEvent
type IAVParameterEvent interface {
	IAVMusicEvent

	// Topic: Creating a Parameter Event

	// Creates an event with a parameter identifier, scope, element, and value for the parameter to set.
	InitWithParameterIDScopeElementValue(parameterID uint32, scope uint32, element uint32, value float32) AVParameterEvent

	// Topic: Configuring a Parameter Event

	// The identifier of the parameter.
	ParameterID() uint32
	SetParameterID(value uint32)
	// The audio unit scope for the parameter.
	Scope() uint32
	SetScope(value uint32)
	// The element index in the scope.
	Element() uint32
	SetElement(value uint32)
	// The value of the parameter to set.
	Value() float32
	SetValue(value float32)
}

// Init initializes the instance.
func (p AVParameterEvent) Init() AVParameterEvent {
	rv := objc.Send[AVParameterEvent](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVParameterEvent) Autorelease() AVParameterEvent {
	rv := objc.Send[AVParameterEvent](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVParameterEvent creates a new AVParameterEvent instance.
func NewAVParameterEvent() AVParameterEvent {
	class := getAVParameterEventClass()
	rv := objc.Send[AVParameterEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an event with a parameter identifier, scope, element, and value for
// the parameter to set.
//
// parameterID: The identifier of the parameter.
//
// scope: The audio unit scope for the parameter.
//
// element: The element index in the scope.
//
// value: The value of the parameter to set.
//
// # Discussion
//
// For more information about the parameters, see [AudioUnitParameterID],
// [AudioUnitScope], and [AudioUnitElement]. The valid range of values depend
// on the parameter you set.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVParameterEvent/init(parameterID:scope:element:value:)
//
// [AudioUnitElement]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitElement
// [AudioUnitParameterID]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterID
// [AudioUnitScope]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitScope
func NewParameterEventWithParameterIDScopeElementValue(parameterID uint32, scope uint32, element uint32, value float32) AVParameterEvent {
	instance := getAVParameterEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameterID:scope:element:value:"), parameterID, scope, element, value)
	return AVParameterEventFromID(rv)
}

// Creates an event with a parameter identifier, scope, element, and value for
// the parameter to set.
//
// parameterID: The identifier of the parameter.
//
// scope: The audio unit scope for the parameter.
//
// element: The element index in the scope.
//
// value: The value of the parameter to set.
//
// # Discussion
//
// For more information about the parameters, see [AudioUnitParameterID],
// [AudioUnitScope], and [AudioUnitElement]. The valid range of values depend
// on the parameter you set.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVParameterEvent/init(parameterID:scope:element:value:)
//
// [AudioUnitElement]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitElement
// [AudioUnitParameterID]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitParameterID
// [AudioUnitScope]: https://developer.apple.com/documentation/AudioToolbox/AudioUnitScope
func (p AVParameterEvent) InitWithParameterIDScopeElementValue(parameterID uint32, scope uint32, element uint32, value float32) AVParameterEvent {
	rv := objc.Send[AVParameterEvent](p.ID, objc.Sel("initWithParameterID:scope:element:value:"), parameterID, scope, element, value)
	return rv
}

// The identifier of the parameter.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVParameterEvent/parameterID
func (p AVParameterEvent) ParameterID() uint32 {
	rv := objc.Send[uint32](p.ID, objc.Sel("parameterID"))
	return rv
}
func (p AVParameterEvent) SetParameterID(value uint32) {
	objc.Send[struct{}](p.ID, objc.Sel("setParameterID:"), value)
}

// The audio unit scope for the parameter.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVParameterEvent/scope
func (p AVParameterEvent) Scope() uint32 {
	rv := objc.Send[uint32](p.ID, objc.Sel("scope"))
	return rv
}
func (p AVParameterEvent) SetScope(value uint32) {
	objc.Send[struct{}](p.ID, objc.Sel("setScope:"), value)
}

// The element index in the scope.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVParameterEvent/element
func (p AVParameterEvent) Element() uint32 {
	rv := objc.Send[uint32](p.ID, objc.Sel("element"))
	return rv
}
func (p AVParameterEvent) SetElement(value uint32) {
	objc.Send[struct{}](p.ID, objc.Sel("setElement:"), value)
}

// The value of the parameter to set.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVParameterEvent/value
func (p AVParameterEvent) Value() float32 {
	rv := objc.Send[float32](p.ID, objc.Sel("value"))
	return rv
}
func (p AVParameterEvent) SetValue(value float32) {
	objc.Send[struct{}](p.ID, objc.Sel("setValue:"), value)
}
