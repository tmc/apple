// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVSpeechSynthesisProviderAudioUnit] class.
var (
	_AVSpeechSynthesisProviderAudioUnitClass     AVSpeechSynthesisProviderAudioUnitClass
	_AVSpeechSynthesisProviderAudioUnitClassOnce sync.Once
)

func getAVSpeechSynthesisProviderAudioUnitClass() AVSpeechSynthesisProviderAudioUnitClass {
	_AVSpeechSynthesisProviderAudioUnitClassOnce.Do(func() {
		_AVSpeechSynthesisProviderAudioUnitClass = AVSpeechSynthesisProviderAudioUnitClass{class: objc.GetClass("AVSpeechSynthesisProviderAudioUnit")}
	})
	return _AVSpeechSynthesisProviderAudioUnitClass
}

// GetAVSpeechSynthesisProviderAudioUnitClass returns the class object for AVSpeechSynthesisProviderAudioUnit.
func GetAVSpeechSynthesisProviderAudioUnitClass() AVSpeechSynthesisProviderAudioUnitClass {
	return getAVSpeechSynthesisProviderAudioUnitClass()
}

type AVSpeechSynthesisProviderAudioUnitClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVSpeechSynthesisProviderAudioUnitClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVSpeechSynthesisProviderAudioUnitClass) Alloc() AVSpeechSynthesisProviderAudioUnit {
	rv := objc.Send[AVSpeechSynthesisProviderAudioUnit](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A parent class referenced by other texttospeech classes. [Full Topic]
type AVSpeechSynthesisProviderAudioUnit struct {
	objectivec.Object
}

// AVSpeechSynthesisProviderAudioUnitFromID constructs a [AVSpeechSynthesisProviderAudioUnit] from an objc.ID.
//
// A parent class referenced by other texttospeech classes.
func AVSpeechSynthesisProviderAudioUnitFromID(id objc.ID) AVSpeechSynthesisProviderAudioUnit {
	return AVSpeechSynthesisProviderAudioUnit{objectivec.Object{ID: id}}
}

// Ensure AVSpeechSynthesisProviderAudioUnit implements IAVSpeechSynthesisProviderAudioUnit.
var _ IAVSpeechSynthesisProviderAudioUnit = AVSpeechSynthesisProviderAudioUnit{}

// An interface definition for the [AVSpeechSynthesisProviderAudioUnit] class.
type IAVSpeechSynthesisProviderAudioUnit interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s AVSpeechSynthesisProviderAudioUnit) Init() AVSpeechSynthesisProviderAudioUnit {
	rv := objc.Send[AVSpeechSynthesisProviderAudioUnit](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s AVSpeechSynthesisProviderAudioUnit) Autorelease() AVSpeechSynthesisProviderAudioUnit {
	rv := objc.Send[AVSpeechSynthesisProviderAudioUnit](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSpeechSynthesisProviderAudioUnit creates a new AVSpeechSynthesisProviderAudioUnit instance.
func NewAVSpeechSynthesisProviderAudioUnit() AVSpeechSynthesisProviderAudioUnit {
	class := getAVSpeechSynthesisProviderAudioUnitClass()
	rv := objc.Send[AVSpeechSynthesisProviderAudioUnit](objc.ID(class.class), objc.Sel("new"))
	return rv
}
