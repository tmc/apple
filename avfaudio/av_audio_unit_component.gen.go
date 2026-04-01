// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioUnitComponent] class.
var (
	_AVAudioUnitComponentClass     AVAudioUnitComponentClass
	_AVAudioUnitComponentClassOnce sync.Once
)

func getAVAudioUnitComponentClass() AVAudioUnitComponentClass {
	_AVAudioUnitComponentClassOnce.Do(func() {
		_AVAudioUnitComponentClass = AVAudioUnitComponentClass{class: objc.GetClass("AVAudioUnitComponent")}
	})
	return _AVAudioUnitComponentClass
}

// GetAVAudioUnitComponentClass returns the class object for AVAudioUnitComponent.
func GetAVAudioUnitComponentClass() AVAudioUnitComponentClass {
	return getAVAudioUnitComponentClass()
}

type AVAudioUnitComponentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioUnitComponentClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioUnitComponentClass) Alloc() AVAudioUnitComponent {
	rv := objc.Send[AVAudioUnitComponent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that provides details about an audio unit.
//
// # Overview
//
// Details can include information such as type, subtype, manufacturer, and
// location. An [AVAudioUnitComponent] can include user tags, which you can
// query later for display.
//
// # Getting the audio unit component’s audio unit
//
//   - [AVAudioUnitComponent.AudioComponent]: The underlying audio component.
//
// # Getting audio unit component information
//
//   - [AVAudioUnitComponent.AudioComponentDescription]: The audio component description.
//   - [AVAudioUnitComponent.AvailableArchitectures]: An array of architectures that the audio unit supports.
//   - [AVAudioUnitComponent.ConfigurationDictionary]: The audio unit component’s configuration dictionary.
//   - [AVAudioUnitComponent.HasCustomView]: A Boolean value that indicates whether the audio unit component has a custom view.
//   - [AVAudioUnitComponent.HasMIDIInput]: A Boolean value that indicates whether the audio unit component has MIDI input.
//   - [AVAudioUnitComponent.HasMIDIOutput]: A Boolean value that indicates whether the audio unit component has MIDI output.
//   - [AVAudioUnitComponent.ManufacturerName]: The name of the manufacturer of the audio unit component.
//   - [AVAudioUnitComponent.Name]: The name of the audio unit component.
//   - [AVAudioUnitComponent.PassesAUVal]: A Boolean value that indicates whether the audio unit component passes the validation tests.
//   - [AVAudioUnitComponent.SandboxSafe]: A Boolean value that indicates whether the audio unit component is safe for sandboxing.
//   - [AVAudioUnitComponent.SupportsNumberInputChannelsOutputChannels]: Gets a Boolean value that indicates whether the audio unit component supports the specified number of input and output channels.
//   - [AVAudioUnitComponent.TypeName]: The audio unit component type.
//   - [AVAudioUnitComponent.Version]: The audio unit component version number.
//   - [AVAudioUnitComponent.VersionString]: A string that represents the audio unit component version number.
//
// # Getting audio unit component tags
//
//   - [AVAudioUnitComponent.IconURL]: The URL of an icon that represents the audio unit component.
//   - [AVAudioUnitComponent.Icon]: An icon that represents the component.
//   - [AVAudioUnitComponent.LocalizedTypeName]: The localized type name of the component.
//   - [AVAudioUnitComponent.AllTagNames]: An array of tag names for the audio unit component.
//   - [AVAudioUnitComponent.UserTagNames]: An array of tags the user creates.
//   - [AVAudioUnitComponent.SetUserTagNames]
//
// # Audio unit manufacturer names
//
//   - [AVAudioUnitComponent.AVAudioUnitManufacturerNameApple]: The audio unit manufacturer is Apple.
//
// # Audio unit types
//
//   - [AVAudioUnitComponent.AVAudioUnitTypeOutput]: An audio unit type that represents an output.
//   - [AVAudioUnitComponent.AVAudioUnitTypeMusicDevice]: An audio unit type that represents a music device.
//   - [AVAudioUnitComponent.AVAudioUnitTypeMusicEffect]: An audio unit type that represents a music effect.
//   - [AVAudioUnitComponent.AVAudioUnitTypeFormatConverter]: An audio unit type that represents a format converter.
//   - [AVAudioUnitComponent.AVAudioUnitTypeEffect]: An audio unit type that represents an effect.
//   - [AVAudioUnitComponent.AVAudioUnitTypeMixer]: An audio unit type that represents a mixer.
//   - [AVAudioUnitComponent.AVAudioUnitTypePanner]: An audio unit type that represents a panner.
//   - [AVAudioUnitComponent.AVAudioUnitTypeGenerator]: An audio unit type that represents a generator.
//   - [AVAudioUnitComponent.AVAudioUnitTypeOfflineEffect]: An audio unit type that represents an offline effect.
//   - [AVAudioUnitComponent.AVAudioUnitTypeMIDIProcessor]: An audio unit type that represents a MIDI processor.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent
type AVAudioUnitComponent struct {
	objectivec.Object
}

// AVAudioUnitComponentFromID constructs a [AVAudioUnitComponent] from an objc.ID.
//
// An object that provides details about an audio unit.
func AVAudioUnitComponentFromID(id objc.ID) AVAudioUnitComponent {
	return AVAudioUnitComponent{objectivec.Object{ID: id}}
}

// NOTE: AVAudioUnitComponent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioUnitComponent] class.
//
// # Getting the audio unit component’s audio unit
//
//   - [IAVAudioUnitComponent.AudioComponent]: The underlying audio component.
//
// # Getting audio unit component information
//
//   - [IAVAudioUnitComponent.AudioComponentDescription]: The audio component description.
//   - [IAVAudioUnitComponent.AvailableArchitectures]: An array of architectures that the audio unit supports.
//   - [IAVAudioUnitComponent.ConfigurationDictionary]: The audio unit component’s configuration dictionary.
//   - [IAVAudioUnitComponent.HasCustomView]: A Boolean value that indicates whether the audio unit component has a custom view.
//   - [IAVAudioUnitComponent.HasMIDIInput]: A Boolean value that indicates whether the audio unit component has MIDI input.
//   - [IAVAudioUnitComponent.HasMIDIOutput]: A Boolean value that indicates whether the audio unit component has MIDI output.
//   - [IAVAudioUnitComponent.ManufacturerName]: The name of the manufacturer of the audio unit component.
//   - [IAVAudioUnitComponent.Name]: The name of the audio unit component.
//   - [IAVAudioUnitComponent.PassesAUVal]: A Boolean value that indicates whether the audio unit component passes the validation tests.
//   - [IAVAudioUnitComponent.SandboxSafe]: A Boolean value that indicates whether the audio unit component is safe for sandboxing.
//   - [IAVAudioUnitComponent.SupportsNumberInputChannelsOutputChannels]: Gets a Boolean value that indicates whether the audio unit component supports the specified number of input and output channels.
//   - [IAVAudioUnitComponent.TypeName]: The audio unit component type.
//   - [IAVAudioUnitComponent.Version]: The audio unit component version number.
//   - [IAVAudioUnitComponent.VersionString]: A string that represents the audio unit component version number.
//
// # Getting audio unit component tags
//
//   - [IAVAudioUnitComponent.IconURL]: The URL of an icon that represents the audio unit component.
//   - [IAVAudioUnitComponent.Icon]: An icon that represents the component.
//   - [IAVAudioUnitComponent.LocalizedTypeName]: The localized type name of the component.
//   - [IAVAudioUnitComponent.AllTagNames]: An array of tag names for the audio unit component.
//   - [IAVAudioUnitComponent.UserTagNames]: An array of tags the user creates.
//   - [IAVAudioUnitComponent.SetUserTagNames]
//
// # Audio unit manufacturer names
//
//   - [IAVAudioUnitComponent.AVAudioUnitManufacturerNameApple]: The audio unit manufacturer is Apple.
//
// # Audio unit types
//
//   - [IAVAudioUnitComponent.AVAudioUnitTypeOutput]: An audio unit type that represents an output.
//   - [IAVAudioUnitComponent.AVAudioUnitTypeMusicDevice]: An audio unit type that represents a music device.
//   - [IAVAudioUnitComponent.AVAudioUnitTypeMusicEffect]: An audio unit type that represents a music effect.
//   - [IAVAudioUnitComponent.AVAudioUnitTypeFormatConverter]: An audio unit type that represents a format converter.
//   - [IAVAudioUnitComponent.AVAudioUnitTypeEffect]: An audio unit type that represents an effect.
//   - [IAVAudioUnitComponent.AVAudioUnitTypeMixer]: An audio unit type that represents a mixer.
//   - [IAVAudioUnitComponent.AVAudioUnitTypePanner]: An audio unit type that represents a panner.
//   - [IAVAudioUnitComponent.AVAudioUnitTypeGenerator]: An audio unit type that represents a generator.
//   - [IAVAudioUnitComponent.AVAudioUnitTypeOfflineEffect]: An audio unit type that represents an offline effect.
//   - [IAVAudioUnitComponent.AVAudioUnitTypeMIDIProcessor]: An audio unit type that represents a MIDI processor.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent
type IAVAudioUnitComponent interface {
	objectivec.IObject

	// Topic: Getting the audio unit component’s audio unit

	// The underlying audio component.
	AudioComponent() unsafe.Pointer

	// Topic: Getting audio unit component information

	// The audio component description.
	AudioComponentDescription() unsafe.Pointer
	// An array of architectures that the audio unit supports.
	AvailableArchitectures() []foundation.NSNumber
	// The audio unit component’s configuration dictionary.
	ConfigurationDictionary() foundation.INSDictionary
	// A Boolean value that indicates whether the audio unit component has a custom view.
	HasCustomView() bool
	// A Boolean value that indicates whether the audio unit component has MIDI input.
	HasMIDIInput() bool
	// A Boolean value that indicates whether the audio unit component has MIDI output.
	HasMIDIOutput() bool
	// The name of the manufacturer of the audio unit component.
	ManufacturerName() string
	// The name of the audio unit component.
	Name() string
	// A Boolean value that indicates whether the audio unit component passes the validation tests.
	PassesAUVal() bool
	// A Boolean value that indicates whether the audio unit component is safe for sandboxing.
	SandboxSafe() bool
	// Gets a Boolean value that indicates whether the audio unit component supports the specified number of input and output channels.
	SupportsNumberInputChannelsOutputChannels(numInputChannels int, numOutputChannels int) bool
	// The audio unit component type.
	TypeName() string
	// The audio unit component version number.
	Version() uint
	// A string that represents the audio unit component version number.
	VersionString() string

	// Topic: Getting audio unit component tags

	// The URL of an icon that represents the audio unit component.
	IconURL() foundation.INSURL
	// An icon that represents the component.
	Icon() objectivec.IObject
	// The localized type name of the component.
	LocalizedTypeName() string
	// An array of tag names for the audio unit component.
	AllTagNames() []string
	// An array of tags the user creates.
	UserTagNames() []string
	SetUserTagNames(value []string)

	// Topic: Audio unit manufacturer names

	// The audio unit manufacturer is Apple.
	AVAudioUnitManufacturerNameApple() string

	// Topic: Audio unit types

	// An audio unit type that represents an output.
	AVAudioUnitTypeOutput() string
	// An audio unit type that represents a music device.
	AVAudioUnitTypeMusicDevice() string
	// An audio unit type that represents a music effect.
	AVAudioUnitTypeMusicEffect() string
	// An audio unit type that represents a format converter.
	AVAudioUnitTypeFormatConverter() string
	// An audio unit type that represents an effect.
	AVAudioUnitTypeEffect() string
	// An audio unit type that represents a mixer.
	AVAudioUnitTypeMixer() string
	// An audio unit type that represents a panner.
	AVAudioUnitTypePanner() string
	// An audio unit type that represents a generator.
	AVAudioUnitTypeGenerator() string
	// An audio unit type that represents an offline effect.
	AVAudioUnitTypeOfflineEffect() string
	// An audio unit type that represents a MIDI processor.
	AVAudioUnitTypeMIDIProcessor() string
}

// Init initializes the instance.
func (a AVAudioUnitComponent) Init() AVAudioUnitComponent {
	rv := objc.Send[AVAudioUnitComponent](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioUnitComponent) Autorelease() AVAudioUnitComponent {
	rv := objc.Send[AVAudioUnitComponent](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioUnitComponent creates a new AVAudioUnitComponent instance.
func NewAVAudioUnitComponent() AVAudioUnitComponent {
	class := getAVAudioUnitComponentClass()
	rv := objc.Send[AVAudioUnitComponent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Gets a Boolean value that indicates whether the audio unit component
// supports the specified number of input and output channels.
//
// numInputChannels: The number of input channels.
//
// numOutputChannels: The number of output channels.
//
// # Return Value
//
// A value of true if the audio unit component supports the specified number
// of input and output channels; otherwise, false.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/supportsNumberInputChannels(_:outputChannels:)
func (a AVAudioUnitComponent) SupportsNumberInputChannelsOutputChannels(numInputChannels int, numOutputChannels int) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("supportsNumberInputChannels:outputChannels:"), numInputChannels, numOutputChannels)
	return rv
}

// The underlying audio component.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/audioComponent
func (a AVAudioUnitComponent) AudioComponent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a.ID, objc.Sel("audioComponent"))
	return rv
}

// The audio component description.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/audioComponentDescription
func (a AVAudioUnitComponent) AudioComponentDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a.ID, objc.Sel("audioComponentDescription"))
	return rv
}

// An array of architectures that the audio unit supports.
//
// # Discussion
//
// This is an [NSArray] of [NSNumbers] where each entry corresponds to one of
// the constants in Mach-O Architecture in [Bundle].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/availableArchitectures
//
// [Bundle]: https://developer.apple.com/documentation/Foundation/Bundle
func (a AVAudioUnitComponent) AvailableArchitectures() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("availableArchitectures"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// The audio unit component’s configuration dictionary.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/configurationDictionary
func (a AVAudioUnitComponent) ConfigurationDictionary() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("configurationDictionary"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the audio unit component has a
// custom view.
//
// # Discussion
//
// This property is true if the component has a custom view; otherwise, false.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/hasCustomView
func (a AVAudioUnitComponent) HasCustomView() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("hasCustomView"))
	return rv
}

// A Boolean value that indicates whether the audio unit component has MIDI
// input.
//
// # Discussion
//
// This property is true if the component has MIDI input; otherwise, false.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/hasMIDIInput
func (a AVAudioUnitComponent) HasMIDIInput() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("hasMIDIInput"))
	return rv
}

// A Boolean value that indicates whether the audio unit component has MIDI
// output.
//
// # Discussion
//
// This property is true if the component has MIDI output; otherwise, false.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/hasMIDIOutput
func (a AVAudioUnitComponent) HasMIDIOutput() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("hasMIDIOutput"))
	return rv
}

// The name of the manufacturer of the audio unit component.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/manufacturerName
func (a AVAudioUnitComponent) ManufacturerName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("manufacturerName"))
	return foundation.NSStringFromID(rv).String()
}

// The name of the audio unit component.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/name
func (a AVAudioUnitComponent) Name() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// A Boolean value that indicates whether the audio unit component passes the
// validation tests.
//
// # Discussion
//
// This property is true if the component passes the validation tests;
// otherwise, false.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/passesAUVal
func (a AVAudioUnitComponent) PassesAUVal() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("passesAUVal"))
	return rv
}

// A Boolean value that indicates whether the audio unit component is safe for
// sandboxing.
//
// # Discussion
//
// This property is true if the component is safe for sandboxing; otherwise,
// false. This only applies to the current process.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/isSandboxSafe
func (a AVAudioUnitComponent) SandboxSafe() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isSandboxSafe"))
	return rv
}

// The audio unit component type.
//
// # Discussion
//
// For information about possible values, see the “Audio Unit Types” topic
// under [AVAudioUnitComponent].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/typeName
func (a AVAudioUnitComponent) TypeName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("typeName"))
	return foundation.NSStringFromID(rv).String()
}

// The audio unit component version number.
//
// # Discussion
//
// The version number is an [NSNumber] made of a hexadecimal number with a
// major, minor, and dot-release format, such as `0xMMMMmmDD`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/version
func (a AVAudioUnitComponent) Version() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("version"))
	return rv
}

// A string that represents the audio unit component version number.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/versionString
func (a AVAudioUnitComponent) VersionString() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("versionString"))
	return foundation.NSStringFromID(rv).String()
}

// The URL of an icon that represents the audio unit component.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/iconURL
func (a AVAudioUnitComponent) IconURL() foundation.INSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("iconURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// An icon that represents the component.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/icon
func (a AVAudioUnitComponent) Icon() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("icon"))
	return objectivec.Object{ID: rv}
}

// The localized type name of the component.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/localizedTypeName
func (a AVAudioUnitComponent) LocalizedTypeName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("localizedTypeName"))
	return foundation.NSStringFromID(rv).String()
}

// An array of tag names for the audio unit component.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/allTagNames
func (a AVAudioUnitComponent) AllTagNames() []string {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("allTagNames"))
	return objc.ConvertSliceToStrings(rv)
}

// An array of tags the user creates.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/userTagNames
func (a AVAudioUnitComponent) UserTagNames() []string {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("userTagNames"))
	return objc.ConvertSliceToStrings(rv)
}
func (a AVAudioUnitComponent) SetUserTagNames(value []string) {
	objc.Send[struct{}](a.ID, objc.Sel("setUserTagNames:"), objectivec.StringSliceToNSArray(value))
}

// The audio unit manufacturer is Apple.
//
// See: https://developer.apple.com/documentation/avfaudio/avaudiounitmanufacturernameapple
func (a AVAudioUnitComponent) AVAudioUnitManufacturerNameApple() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("AVAudioUnitManufacturerNameApple"))
	return foundation.NSStringFromID(rv).String()
}

// An audio unit type that represents an output.
//
// See: https://developer.apple.com/documentation/avfaudio/avaudiounittypeoutput
func (a AVAudioUnitComponent) AVAudioUnitTypeOutput() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("AVAudioUnitTypeOutput"))
	return foundation.NSStringFromID(rv).String()
}

// An audio unit type that represents a music device.
//
// See: https://developer.apple.com/documentation/avfaudio/avaudiounittypemusicdevice
func (a AVAudioUnitComponent) AVAudioUnitTypeMusicDevice() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("AVAudioUnitTypeMusicDevice"))
	return foundation.NSStringFromID(rv).String()
}

// An audio unit type that represents a music effect.
//
// See: https://developer.apple.com/documentation/avfaudio/avaudiounittypemusiceffect
func (a AVAudioUnitComponent) AVAudioUnitTypeMusicEffect() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("AVAudioUnitTypeMusicEffect"))
	return foundation.NSStringFromID(rv).String()
}

// An audio unit type that represents a format converter.
//
// See: https://developer.apple.com/documentation/avfaudio/avaudiounittypeformatconverter
func (a AVAudioUnitComponent) AVAudioUnitTypeFormatConverter() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("AVAudioUnitTypeFormatConverter"))
	return foundation.NSStringFromID(rv).String()
}

// An audio unit type that represents an effect.
//
// See: https://developer.apple.com/documentation/avfaudio/avaudiounittypeeffect
func (a AVAudioUnitComponent) AVAudioUnitTypeEffect() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("AVAudioUnitTypeEffect"))
	return foundation.NSStringFromID(rv).String()
}

// An audio unit type that represents a mixer.
//
// See: https://developer.apple.com/documentation/avfaudio/avaudiounittypemixer
func (a AVAudioUnitComponent) AVAudioUnitTypeMixer() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("AVAudioUnitTypeMixer"))
	return foundation.NSStringFromID(rv).String()
}

// An audio unit type that represents a panner.
//
// See: https://developer.apple.com/documentation/avfaudio/avaudiounittypepanner
func (a AVAudioUnitComponent) AVAudioUnitTypePanner() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("AVAudioUnitTypePanner"))
	return foundation.NSStringFromID(rv).String()
}

// An audio unit type that represents a generator.
//
// See: https://developer.apple.com/documentation/avfaudio/avaudiounittypegenerator
func (a AVAudioUnitComponent) AVAudioUnitTypeGenerator() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("AVAudioUnitTypeGenerator"))
	return foundation.NSStringFromID(rv).String()
}

// An audio unit type that represents an offline effect.
//
// See: https://developer.apple.com/documentation/avfaudio/avaudiounittypeofflineeffect
func (a AVAudioUnitComponent) AVAudioUnitTypeOfflineEffect() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("AVAudioUnitTypeOfflineEffect"))
	return foundation.NSStringFromID(rv).String()
}

// An audio unit type that represents a MIDI processor.
//
// See: https://developer.apple.com/documentation/avfaudio/avaudiounittypemidiprocessor
func (a AVAudioUnitComponent) AVAudioUnitTypeMIDIProcessor() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("AVAudioUnitTypeMIDIProcessor"))
	return foundation.NSStringFromID(rv).String()
}
