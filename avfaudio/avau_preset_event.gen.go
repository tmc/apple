// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVAUPresetEvent] class.
var (
	_AVAUPresetEventClass     AVAUPresetEventClass
	_AVAUPresetEventClassOnce sync.Once
)

func getAVAUPresetEventClass() AVAUPresetEventClass {
	_AVAUPresetEventClassOnce.Do(func() {
		_AVAUPresetEventClass = AVAUPresetEventClass{class: objc.GetClass("AVAUPresetEvent")}
	})
	return _AVAUPresetEventClass
}

// GetAVAUPresetEventClass returns the class object for AVAUPresetEvent.
func GetAVAUPresetEventClass() AVAUPresetEventClass {
	return getAVAUPresetEventClass()
}

type AVAUPresetEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAUPresetEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAUPresetEventClass) Alloc() AVAUPresetEvent {
	rv := objc.Send[AVAUPresetEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a preset load and change on the music track’s
// destination audio unit.
//
// # Creating a Preset Event
//
//   - [AVAUPresetEvent.InitWithScopeElementDictionary]: Creates an event with the scope, element, and dictionary for the preset.
//
// # Configuring a Preset Event
//
//   - [AVAUPresetEvent.Scope]: The audio unit scope.
//   - [AVAUPresetEvent.SetScope]
//   - [AVAUPresetEvent.Element]: The element index in the scope.
//   - [AVAUPresetEvent.SetElement]
//   - [AVAUPresetEvent.PresetDictionary]: The dictionary that contains the preset.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAUPresetEvent
type AVAUPresetEvent struct {
	AVMusicEvent
}

// AVAUPresetEventFromID constructs a [AVAUPresetEvent] from an objc.ID.
//
// An object that represents a preset load and change on the music track’s
// destination audio unit.
func AVAUPresetEventFromID(id objc.ID) AVAUPresetEvent {
	return AVAUPresetEvent{AVMusicEvent: AVMusicEventFromID(id)}
}

// NOTE: AVAUPresetEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAUPresetEvent] class.
//
// # Creating a Preset Event
//
//   - [IAVAUPresetEvent.InitWithScopeElementDictionary]: Creates an event with the scope, element, and dictionary for the preset.
//
// # Configuring a Preset Event
//
//   - [IAVAUPresetEvent.Scope]: The audio unit scope.
//   - [IAVAUPresetEvent.SetScope]
//   - [IAVAUPresetEvent.Element]: The element index in the scope.
//   - [IAVAUPresetEvent.SetElement]
//   - [IAVAUPresetEvent.PresetDictionary]: The dictionary that contains the preset.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAUPresetEvent
type IAVAUPresetEvent interface {
	IAVMusicEvent

	// Topic: Creating a Preset Event

	// Creates an event with the scope, element, and dictionary for the preset.
	InitWithScopeElementDictionary(scope uint32, element uint32, presetDictionary foundation.INSDictionary) AVAUPresetEvent

	// Topic: Configuring a Preset Event

	// The audio unit scope.
	Scope() uint32
	SetScope(value uint32)
	// The element index in the scope.
	Element() uint32
	SetElement(value uint32)
	// The dictionary that contains the preset.
	PresetDictionary() foundation.INSDictionary
}

// Init initializes the instance.
func (a AVAUPresetEvent) Init() AVAUPresetEvent {
	rv := objc.Send[AVAUPresetEvent](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAUPresetEvent) Autorelease() AVAUPresetEvent {
	rv := objc.Send[AVAUPresetEvent](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAUPresetEvent creates a new AVAUPresetEvent instance.
func NewAVAUPresetEvent() AVAUPresetEvent {
	class := getAVAUPresetEventClass()
	rv := objc.Send[AVAUPresetEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an event with the scope, element, and dictionary for the preset.
//
// scope: The audio unit scope.
//
// element: The element index in the scope.
//
// presetDictionary: The dictionary that contains the preset.
//
// # Discussion
//
// The system copies the dictionary you specify and isn’t editable once it
// creates the event. The `scope` parameter must be [kAudioUnitScope_Global],
// and the element index should be `0`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAUPresetEvent/init(scope:element:dictionary:)
//
// [kAudioUnitScope_Global]: https://developer.apple.com/documentation/AudioToolbox/kAudioUnitScope_Global
func NewAUPresetEventWithScopeElementDictionary(scope uint32, element uint32, presetDictionary foundation.INSDictionary) AVAUPresetEvent {
	instance := getAVAUPresetEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithScope:element:dictionary:"), scope, element, presetDictionary)
	return AVAUPresetEventFromID(rv)
}

// Creates an event with the scope, element, and dictionary for the preset.
//
// scope: The audio unit scope.
//
// element: The element index in the scope.
//
// presetDictionary: The dictionary that contains the preset.
//
// # Discussion
//
// The system copies the dictionary you specify and isn’t editable once it
// creates the event. The `scope` parameter must be [kAudioUnitScope_Global],
// and the element index should be `0`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAUPresetEvent/init(scope:element:dictionary:)
//
// [kAudioUnitScope_Global]: https://developer.apple.com/documentation/AudioToolbox/kAudioUnitScope_Global
func (a AVAUPresetEvent) InitWithScopeElementDictionary(scope uint32, element uint32, presetDictionary foundation.INSDictionary) AVAUPresetEvent {
	rv := objc.Send[AVAUPresetEvent](a.ID, objc.Sel("initWithScope:element:dictionary:"), scope, element, presetDictionary)
	return rv
}

// The audio unit scope.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAUPresetEvent/scope
func (a AVAUPresetEvent) Scope() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("scope"))
	return rv
}
func (a AVAUPresetEvent) SetScope(value uint32) {
	objc.Send[struct{}](a.ID, objc.Sel("setScope:"), value)
}

// The element index in the scope.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAUPresetEvent/element
func (a AVAUPresetEvent) Element() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("element"))
	return rv
}
func (a AVAUPresetEvent) SetElement(value uint32) {
	objc.Send[struct{}](a.ID, objc.Sel("setElement:"), value)
}

// The dictionary that contains the preset.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAUPresetEvent/presetDictionary
func (a AVAUPresetEvent) PresetDictionary() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("presetDictionary"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
