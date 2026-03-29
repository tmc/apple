// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioUnit] class.
var (
	_AVAudioUnitClass     AVAudioUnitClass
	_AVAudioUnitClassOnce sync.Once
)

func getAVAudioUnitClass() AVAudioUnitClass {
	_AVAudioUnitClassOnce.Do(func() {
		_AVAudioUnitClass = AVAudioUnitClass{class: objc.GetClass("AVAudioUnit")}
	})
	return _AVAudioUnitClass
}

// GetAVAudioUnitClass returns the class object for AVAudioUnit.
func GetAVAudioUnitClass() AVAudioUnitClass {
	return getAVAudioUnitClass()
}

type AVAudioUnitClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioUnitClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioUnitClass) Alloc() AVAudioUnit {
	rv := objc.Send[AVAudioUnit](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A subclass of the audio node class that, processes audio either in real
// time or nonreal time, depending on the type of the audio unit.
//
// # Getting the Core Audio audio unit
//
//   - [AVAudioUnit.AudioUnit]: The underlying Core Audio audio unit.
//
// # Loading an audio preset file
//
//   - [AVAudioUnit.LoadAudioUnitPresetAtURLError]: Loads an audio unit using a specified preset.
//
// # Getting audio unit values
//
//   - [AVAudioUnit.AudioComponentDescription]: The audio component description that represents the underlying Core Audio audio unit.
//   - [AVAudioUnit.ManufacturerName]: The name of the manufacturer of the audio unit.
//   - [AVAudioUnit.Name]: The name of the audio unit.
//   - [AVAudioUnit.Version]: The version number of the audio unit.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit
type AVAudioUnit struct {
	AVAudioNode
}

// AVAudioUnitFromID constructs a [AVAudioUnit] from an objc.ID.
//
// A subclass of the audio node class that, processes audio either in real
// time or nonreal time, depending on the type of the audio unit.
func AVAudioUnitFromID(id objc.ID) AVAudioUnit {
	return AVAudioUnit{AVAudioNode: AVAudioNodeFromID(id)}
}
// NOTE: AVAudioUnit adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioUnit] class.
//
// # Getting the Core Audio audio unit
//
//   - [IAVAudioUnit.AudioUnit]: The underlying Core Audio audio unit.
//
// # Loading an audio preset file
//
//   - [IAVAudioUnit.LoadAudioUnitPresetAtURLError]: Loads an audio unit using a specified preset.
//
// # Getting audio unit values
//
//   - [IAVAudioUnit.AudioComponentDescription]: The audio component description that represents the underlying Core Audio audio unit.
//   - [IAVAudioUnit.ManufacturerName]: The name of the manufacturer of the audio unit.
//   - [IAVAudioUnit.Name]: The name of the audio unit.
//   - [IAVAudioUnit.Version]: The version number of the audio unit.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit
type IAVAudioUnit interface {
	IAVAudioNode

	// Topic: Getting the Core Audio audio unit

	// The underlying Core Audio audio unit.
	AudioUnit() IAVAudioUnit

	// Topic: Loading an audio preset file

	// Loads an audio unit using a specified preset.
	LoadAudioUnitPresetAtURLError(url foundation.INSURL) (bool, error)

	// Topic: Getting audio unit values

	// The audio component description that represents the underlying Core Audio audio unit.
	AudioComponentDescription() objectivec.IObject
	// The name of the manufacturer of the audio unit.
	ManufacturerName() string
	// The name of the audio unit.
	Name() string
	// The version number of the audio unit.
	Version() uint
}

// Init initializes the instance.
func (a AVAudioUnit) Init() AVAudioUnit {
	rv := objc.Send[AVAudioUnit](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioUnit) Autorelease() AVAudioUnit {
	rv := objc.Send[AVAudioUnit](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioUnit creates a new AVAudioUnit instance.
func NewAVAudioUnit() AVAudioUnit {
	class := getAVAudioUnitClass()
	rv := objc.Send[AVAudioUnit](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Loads an audio unit using a specified preset.
//
// url: The URL of an audio unit preset file.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit/loadPreset(at:)
func (a AVAudioUnit) LoadAudioUnitPresetAtURLError(url foundation.INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("loadAudioUnitPresetAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("loadAudioUnitPresetAtURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Creates an instance of an audio unit component asynchronously and wraps it
// in an audio unit class.
//
// audioComponentDescription: The component to create.
//
// options: The options the method uses to create the component.
//
// completionHandler: A handler the framework calls in an arbitrary thread context when creation
// completes. Retain the [AVAudioUnit] this handler provides.
//
// audioComponentDescription is a [audiotoolbox.AudioComponentDescription].
//
// options is a [audiotoolbox.AudioComponentInstantiationOptions].
//
// # Discussion
// 
// You must create components with flags that include
// [requiresAsyncInstantiation] asynchronously through this method if
// they’re for use with [AVAudioEngine].
// 
// The [AVAudioUnit] instance is usually a subclass that the method selects
// according to the components type. For example, [AVAudioUnitEffect],
// [AVAudioUnitGenerator], [AVAudioUnitMIDIInstrument], or
// [AVAudioUnitTimeEffect].
//
// [requiresAsyncInstantiation]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentFlags/requiresAsyncInstantiation
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit/instantiate(with:options:completionHandler:)
// audioComponentDescription is a [audiotoolbox.AudioComponentDescription].
// options is a [audiotoolbox.AudioComponentInstantiationOptions].
func (_AVAudioUnitClass AVAudioUnitClass) InstantiateWithComponentDescriptionOptionsCompletionHandler(audioComponentDescription objectivec.IObject, options objectivec.IObject, completionHandler AVAudioUnitErrorHandler) {
_block2, _ := NewAVAudioUnitErrorBlock(completionHandler)
	objc.Send[objc.ID](objc.ID(_AVAudioUnitClass.class), objc.Sel("instantiateWithComponentDescription:options:completionHandler:"), audioComponentDescription, options, _block2)
}

// The underlying Core Audio audio unit.
//
// # Discussion
// 
// This property is a reference to the underlying audio unit. The
// [AVAudioUnit] exposes it here so that you can modify parameters, that you
// don’t see through [AVAudioUnit] subclasses, using the AudioUnit C API.
// For example, changing initialization state, stream formats, channel
// layouts, or connections to other audio units.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit/audioUnit
func (a AVAudioUnit) AudioUnit() IAVAudioUnit {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("audioUnit"))
	return AVAudioUnitFromID(objc.ID(rv))
}
// The audio component description that represents the underlying Core Audio
// audio unit.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit/audioComponentDescription
func (a AVAudioUnit) AudioComponentDescription() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("audioComponentDescription"))
	return objectivec.Object{ID: rv}
}
// The name of the manufacturer of the audio unit.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit/manufacturerName
func (a AVAudioUnit) ManufacturerName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("manufacturerName"))
	return foundation.NSStringFromID(rv).String()
}
// The name of the audio unit.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit/name
func (a AVAudioUnit) Name() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
// The version number of the audio unit.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit/version
func (a AVAudioUnit) Version() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("version"))
	return rv
}

// InstantiateWithComponentDescriptionOptions is a synchronous wrapper around [AVAudioUnit.InstantiateWithComponentDescriptionOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (ac AVAudioUnitClass) InstantiateWithComponentDescriptionOptions(ctx context.Context, audioComponentDescription objectivec.IObject, options objectivec.IObject) (*AVAudioUnit, error) {
	type result struct {
		val *AVAudioUnit
		err error
	}
	done := make(chan result, 1)
	ac.InstantiateWithComponentDescriptionOptionsCompletionHandler(audioComponentDescription, options, func(val *AVAudioUnit, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

