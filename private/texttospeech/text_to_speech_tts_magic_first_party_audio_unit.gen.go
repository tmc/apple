// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechTTSMagicFirstPartyAudioUnit] class.
var (
	_TextToSpeechTTSMagicFirstPartyAudioUnitClass     TextToSpeechTTSMagicFirstPartyAudioUnitClass
	_TextToSpeechTTSMagicFirstPartyAudioUnitClassOnce sync.Once
)

func getTextToSpeechTTSMagicFirstPartyAudioUnitClass() TextToSpeechTTSMagicFirstPartyAudioUnitClass {
	_TextToSpeechTTSMagicFirstPartyAudioUnitClassOnce.Do(func() {
		_TextToSpeechTTSMagicFirstPartyAudioUnitClass = TextToSpeechTTSMagicFirstPartyAudioUnitClass{class: objc.GetClass("TextToSpeech.TTSMagicFirstPartyAudioUnit")}
	})
	return _TextToSpeechTTSMagicFirstPartyAudioUnitClass
}

// GetTextToSpeechTTSMagicFirstPartyAudioUnitClass returns the class object for TextToSpeech.TTSMagicFirstPartyAudioUnit.
func GetTextToSpeechTTSMagicFirstPartyAudioUnitClass() TextToSpeechTTSMagicFirstPartyAudioUnitClass {
	return getTextToSpeechTTSMagicFirstPartyAudioUnitClass()
}

type TextToSpeechTTSMagicFirstPartyAudioUnitClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechTTSMagicFirstPartyAudioUnitClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechTTSMagicFirstPartyAudioUnitClass) Alloc() TextToSpeechTTSMagicFirstPartyAudioUnit {
	rv := objc.Send[TextToSpeechTTSMagicFirstPartyAudioUnit](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TextToSpeechTTSMagicFirstPartyAudioUnit.InitWithComponentDescriptionOptionsError]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSMagicFirstPartyAudioUnit
type TextToSpeechTTSMagicFirstPartyAudioUnit struct {
	TTSFirstPartyAudioUnit
}

// TextToSpeechTTSMagicFirstPartyAudioUnitFromID constructs a [TextToSpeechTTSMagicFirstPartyAudioUnit] from an objc.ID.
func TextToSpeechTTSMagicFirstPartyAudioUnitFromID(id objc.ID) TextToSpeechTTSMagicFirstPartyAudioUnit {
	return TextToSpeechTTSMagicFirstPartyAudioUnit{TTSFirstPartyAudioUnit: TTSFirstPartyAudioUnitFromID(id)}
}

// Ensure TextToSpeechTTSMagicFirstPartyAudioUnit implements ITextToSpeechTTSMagicFirstPartyAudioUnit.
var _ ITextToSpeechTTSMagicFirstPartyAudioUnit = TextToSpeechTTSMagicFirstPartyAudioUnit{}

// An interface definition for the [TextToSpeechTTSMagicFirstPartyAudioUnit] class.
//
// # Methods
//
//   - [ITextToSpeechTTSMagicFirstPartyAudioUnit.InitWithComponentDescriptionOptionsError]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSMagicFirstPartyAudioUnit
type ITextToSpeechTTSMagicFirstPartyAudioUnit interface {
	ITTSFirstPartyAudioUnit

	// Topic: Methods

	InitWithComponentDescriptionOptionsError(description objectivec.IObject, options uint32) (TextToSpeechTTSMagicFirstPartyAudioUnit, error)
}

// Init initializes the instance.
func (t TextToSpeechTTSMagicFirstPartyAudioUnit) Init() TextToSpeechTTSMagicFirstPartyAudioUnit {
	rv := objc.Send[TextToSpeechTTSMagicFirstPartyAudioUnit](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechTTSMagicFirstPartyAudioUnit) Autorelease() TextToSpeechTTSMagicFirstPartyAudioUnit {
	rv := objc.Send[TextToSpeechTTSMagicFirstPartyAudioUnit](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechTTSMagicFirstPartyAudioUnit creates a new TextToSpeechTTSMagicFirstPartyAudioUnit instance.
func NewTextToSpeechTTSMagicFirstPartyAudioUnit() TextToSpeechTTSMagicFirstPartyAudioUnit {
	class := getTextToSpeechTTSMagicFirstPartyAudioUnitClass()
	rv := objc.Send[TextToSpeechTTSMagicFirstPartyAudioUnit](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSMagicFirstPartyAudioUnit/initWithComponentDescription:options:error:
func NewTextToSpeechTTSMagicFirstPartyAudioUnitWithComponentDescriptionOptionsError(description objectivec.IObject, options uint32) (TextToSpeechTTSMagicFirstPartyAudioUnit, error) {
	var errorPtr objc.ID
	instance := getTextToSpeechTTSMagicFirstPartyAudioUnitClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithComponentDescription:options:error:"), description, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return TextToSpeechTTSMagicFirstPartyAudioUnit{}, foundation.NSErrorFrom(errorPtr)
	}
	return TextToSpeechTTSMagicFirstPartyAudioUnitFromID(rv), nil
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSMagicFirstPartyAudioUnit/initWithComponentDescription:options:error:
func (t TextToSpeechTTSMagicFirstPartyAudioUnit) InitWithComponentDescriptionOptionsError(description objectivec.IObject, options uint32) (TextToSpeechTTSMagicFirstPartyAudioUnit, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("initWithComponentDescription:options:error:"), description, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return TextToSpeechTTSMagicFirstPartyAudioUnit{}, foundation.NSErrorFrom(errorPtr)
	}
	return TextToSpeechTTSMagicFirstPartyAudioUnitFromID(rv), nil

}
