// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSAudioEffectConfiguration] class.
var (
	_TTSAudioEffectConfigurationClass     TTSAudioEffectConfigurationClass
	_TTSAudioEffectConfigurationClassOnce sync.Once
)

func getTTSAudioEffectConfigurationClass() TTSAudioEffectConfigurationClass {
	_TTSAudioEffectConfigurationClassOnce.Do(func() {
		_TTSAudioEffectConfigurationClass = TTSAudioEffectConfigurationClass{class: objc.GetClass("TTSAudioEffectConfiguration")}
	})
	return _TTSAudioEffectConfigurationClass
}

// GetTTSAudioEffectConfigurationClass returns the class object for TTSAudioEffectConfiguration.
func GetTTSAudioEffectConfigurationClass() TTSAudioEffectConfigurationClass {
	return getTTSAudioEffectConfigurationClass()
}

type TTSAudioEffectConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSAudioEffectConfigurationClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSAudioEffectConfigurationClass) Alloc() TTSAudioEffectConfiguration {
	rv := objc.Send[TTSAudioEffectConfiguration](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TTSAudioEffectConfiguration.EffectName]
//   - [TTSAudioEffectConfiguration.SetEffectName]
//   - [TTSAudioEffectConfiguration.Enabled]
//   - [TTSAudioEffectConfiguration.SetEnabled]
//   - [TTSAudioEffectConfiguration.Identifier]
//   - [TTSAudioEffectConfiguration.SetIdentifier]
//   - [TTSAudioEffectConfiguration.Parameters]
//   - [TTSAudioEffectConfiguration.SetParameters]
//   - [TTSAudioEffectConfiguration.Properties]
//   - [TTSAudioEffectConfiguration.SetProperties]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioEffectConfiguration
type TTSAudioEffectConfiguration struct {
	objectivec.Object
}

// TTSAudioEffectConfigurationFromID constructs a [TTSAudioEffectConfiguration] from an objc.ID.
func TTSAudioEffectConfigurationFromID(id objc.ID) TTSAudioEffectConfiguration {
	return TTSAudioEffectConfiguration{objectivec.Object{ID: id}}
}

// Ensure TTSAudioEffectConfiguration implements ITTSAudioEffectConfiguration.
var _ ITTSAudioEffectConfiguration = TTSAudioEffectConfiguration{}

// An interface definition for the [TTSAudioEffectConfiguration] class.
//
// # Methods
//
//   - [ITTSAudioEffectConfiguration.EffectName]
//   - [ITTSAudioEffectConfiguration.SetEffectName]
//   - [ITTSAudioEffectConfiguration.Enabled]
//   - [ITTSAudioEffectConfiguration.SetEnabled]
//   - [ITTSAudioEffectConfiguration.Identifier]
//   - [ITTSAudioEffectConfiguration.SetIdentifier]
//   - [ITTSAudioEffectConfiguration.Parameters]
//   - [ITTSAudioEffectConfiguration.SetParameters]
//   - [ITTSAudioEffectConfiguration.Properties]
//   - [ITTSAudioEffectConfiguration.SetProperties]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioEffectConfiguration
type ITTSAudioEffectConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	EffectName() string
	SetEffectName(value string)
	Enabled() bool
	SetEnabled(value bool)
	Identifier() foundation.NSUUID
	SetIdentifier(value foundation.NSUUID)
	Parameters() foundation.INSDictionary
	SetParameters(value foundation.INSDictionary)
	Properties() foundation.INSDictionary
	SetProperties(value foundation.INSDictionary)
}

// Init initializes the instance.
func (t TTSAudioEffectConfiguration) Init() TTSAudioEffectConfiguration {
	rv := objc.Send[TTSAudioEffectConfiguration](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSAudioEffectConfiguration) Autorelease() TTSAudioEffectConfiguration {
	rv := objc.Send[TTSAudioEffectConfiguration](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSAudioEffectConfiguration creates a new TTSAudioEffectConfiguration instance.
func NewTTSAudioEffectConfiguration() TTSAudioEffectConfiguration {
	class := getTTSAudioEffectConfigurationClass()
	rv := objc.Send[TTSAudioEffectConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioEffectConfiguration/effectName
func (t TTSAudioEffectConfiguration) EffectName() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("effectName"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSAudioEffectConfiguration) SetEffectName(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setEffectName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioEffectConfiguration/enabled
func (t TTSAudioEffectConfiguration) Enabled() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("enabled"))
	return rv
}
func (t TTSAudioEffectConfiguration) SetEnabled(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setEnabled:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioEffectConfiguration/identifier
func (t TTSAudioEffectConfiguration) Identifier() foundation.NSUUID {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("identifier"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
func (t TTSAudioEffectConfiguration) SetIdentifier(value foundation.NSUUID) {
	objc.Send[struct{}](t.ID, objc.Sel("setIdentifier:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioEffectConfiguration/parameters
func (t TTSAudioEffectConfiguration) Parameters() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("parameters"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t TTSAudioEffectConfiguration) SetParameters(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setParameters:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioEffectConfiguration/properties
func (t TTSAudioEffectConfiguration) Properties() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("properties"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t TTSAudioEffectConfiguration) SetProperties(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setProperties:"), value)
}
