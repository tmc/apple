// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioUnitComponentManager] class.
var (
	_AVAudioUnitComponentManagerClass     AVAudioUnitComponentManagerClass
	_AVAudioUnitComponentManagerClassOnce sync.Once
)

func getAVAudioUnitComponentManagerClass() AVAudioUnitComponentManagerClass {
	_AVAudioUnitComponentManagerClassOnce.Do(func() {
		_AVAudioUnitComponentManagerClass = AVAudioUnitComponentManagerClass{class: objc.GetClass("AVAudioUnitComponentManager")}
	})
	return _AVAudioUnitComponentManagerClass
}

// GetAVAudioUnitComponentManagerClass returns the class object for AVAudioUnitComponentManager.
func GetAVAudioUnitComponentManagerClass() AVAudioUnitComponentManagerClass {
	return getAVAudioUnitComponentManagerClass()
}

type AVAudioUnitComponentManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioUnitComponentManagerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioUnitComponentManagerClass) Alloc() AVAudioUnitComponentManager {
	rv := objc.Send[AVAudioUnitComponentManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that provides a way to search and query audio components that the
// system registers.
//
// # Overview
//
// The component manager has methods to find various information about the
// audio components without opening them. Currently, you can only search audio
// components that are audio units.
//
// The class supports system tags and arbitrary user tags. You can tag each
// audio unit as part of its definition. Audio unit hosts, such as Logic or
// GarageBand, can present groupings of audio units according to the tags.
//
// You can search for audio units in the following ways:
//
// - Using a [NSPredicate] instance that contains search strings for tags or
// descriptions - Using a block to match on a custom criteria - Using an
// [AudioComponentDescription]
//
// # Getting matching audio components
//
//   - [AVAudioUnitComponentManager.ComponentsMatchingDescription]: Gets an array of audio component objects that match the description.
//   - [AVAudioUnitComponentManager.ComponentsMatchingPredicate]: Gets an array of audio component objects that match the search predicate.
//   - [AVAudioUnitComponentManager.ComponentsPassingTest]: Gets an array of audio components that pass the block method.
//
// # Getting audio unit tags
//
//   - [AVAudioUnitComponentManager.StandardLocalizedTagNames]: An array of the localized standard system tags the audio units define.
//   - [AVAudioUnitComponentManager.TagNames]: An array of all tags the audio unit associates with the current user, and the system tags the audio units define.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponentManager
type AVAudioUnitComponentManager struct {
	objectivec.Object
}

// AVAudioUnitComponentManagerFromID constructs a [AVAudioUnitComponentManager] from an objc.ID.
//
// An object that provides a way to search and query audio components that the
// system registers.
func AVAudioUnitComponentManagerFromID(id objc.ID) AVAudioUnitComponentManager {
	return AVAudioUnitComponentManager{objectivec.Object{ID: id}}
}

// NOTE: AVAudioUnitComponentManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioUnitComponentManager] class.
//
// # Getting matching audio components
//
//   - [IAVAudioUnitComponentManager.ComponentsMatchingDescription]: Gets an array of audio component objects that match the description.
//   - [IAVAudioUnitComponentManager.ComponentsMatchingPredicate]: Gets an array of audio component objects that match the search predicate.
//   - [IAVAudioUnitComponentManager.ComponentsPassingTest]: Gets an array of audio components that pass the block method.
//
// # Getting audio unit tags
//
//   - [IAVAudioUnitComponentManager.StandardLocalizedTagNames]: An array of the localized standard system tags the audio units define.
//   - [IAVAudioUnitComponentManager.TagNames]: An array of all tags the audio unit associates with the current user, and the system tags the audio units define.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponentManager
type IAVAudioUnitComponentManager interface {
	objectivec.IObject

	// Topic: Getting matching audio components

	// Gets an array of audio component objects that match the description.
	ComponentsMatchingDescription(desc unsafe.Pointer) []AVAudioUnitComponent
	// Gets an array of audio component objects that match the search predicate.
	ComponentsMatchingPredicate(predicate foundation.INSPredicate) []AVAudioUnitComponent
	// Gets an array of audio components that pass the block method.
	ComponentsPassingTest(testHandler AVAudioUnitComponentHandler) []AVAudioUnitComponent

	// Topic: Getting audio unit tags

	// An array of the localized standard system tags the audio units define.
	StandardLocalizedTagNames() []string
	// An array of all tags the audio unit associates with the current user, and the system tags the audio units define.
	TagNames() []string
}

// Init initializes the instance.
func (a AVAudioUnitComponentManager) Init() AVAudioUnitComponentManager {
	rv := objc.Send[AVAudioUnitComponentManager](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioUnitComponentManager) Autorelease() AVAudioUnitComponentManager {
	rv := objc.Send[AVAudioUnitComponentManager](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioUnitComponentManager creates a new AVAudioUnitComponentManager instance.
func NewAVAudioUnitComponentManager() AVAudioUnitComponentManager {
	class := getAVAudioUnitComponentManagerClass()
	rv := objc.Send[AVAudioUnitComponentManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Gets an array of audio component objects that match the description.
//
// desc is a [audiotoolbox.AudioComponentDescription].
//
// # Return Value
//
// An array of [AVAudioComponent] objects that match the `description`.
//
// # Discussion
//
// - desc: The [AudioComponentDescription] structure to match. The method uses
// the `type`, `subtype` and `manufacturer` fields to search for matching
// audio units. A value of `0` for any of these fields is a wildcard and
// returns the first match the method finds.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponentManager/components(matching:)-9qt94
//
// [AudioComponentDescription]: https://developer.apple.com/documentation/AudioToolbox/AudioComponentDescription
func (a AVAudioUnitComponentManager) ComponentsMatchingDescription(desc unsafe.Pointer) []AVAudioUnitComponent {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("componentsMatchingDescription:"), desc)
	return objc.ConvertSlice(rv, func(id objc.ID) AVAudioUnitComponent {
		return AVAudioUnitComponentFromID(id)
	})
}

// Gets an array of audio component objects that match the search predicate.
//
// predicate: The search predicate.
//
// # Return Value
//
// An array of [AVAudioComponent] objects that match the predicate.
//
// # Discussion
//
// You use the audio component’s information or tags to build search
// criteria, such as `“typeName CONTAINS 'Effect'"` or `“tags IN
// {'Sampler', 'MIDI'}"`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponentManager/components(matching:)-96l2c
func (a AVAudioUnitComponentManager) ComponentsMatchingPredicate(predicate foundation.INSPredicate) []AVAudioUnitComponent {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("componentsMatchingPredicate:"), predicate)
	return objc.ConvertSlice(rv, func(id objc.ID) AVAudioUnitComponent {
		return AVAudioUnitComponentFromID(id)
	})
}

// Gets an array of audio components that pass the block method.
//
// testHandler: The block to apply to the audio unit components.
//
// The block takes two parameters.
//
// comp: A block to test. stop: A reference to a Boolean value. To stop
// further processing of the search, the block sets the value to true. The
// stop argument is an out-only argument. Only set this Boolean to true within
// the block.
//
// The block returns a Boolean value that indicates whether `comp` passes the
// test. Returning true stops further processing of the audio components.
//
// # Return Value
//
// An array of audio components that pass the test.
//
// # Discussion
//
// For each audio component the manager finds, the system calls the block
// method. If the block returns true, the method adds [AVAudioComponent]
// instance to the array.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponentManager/components(passingTest:)
func (a AVAudioUnitComponentManager) ComponentsPassingTest(testHandler AVAudioUnitComponentHandler) []AVAudioUnitComponent {
	_block0, _ := NewAVAudioUnitComponentBlock(testHandler)
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("componentsPassingTest:"), _block0)
	return objc.ConvertSlice(rv, func(id objc.ID) AVAudioUnitComponent {
		return AVAudioUnitComponentFromID(id)
	})
}

// Gets the shared component manager instance.
//
// # Return Value
//
// The singleton instance of the [AVAudioUnitComponentManager] object.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponentManager/shared()
func (_AVAudioUnitComponentManagerClass AVAudioUnitComponentManagerClass) SharedAudioUnitComponentManager() AVAudioUnitComponentManager {
	rv := objc.Send[objc.ID](objc.ID(_AVAudioUnitComponentManagerClass.class), objc.Sel("sharedAudioUnitComponentManager"))
	return AVAudioUnitComponentManagerFromID(rv)
}

// An array of the localized standard system tags the audio units define.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponentManager/standardLocalizedTagNames
func (a AVAudioUnitComponentManager) StandardLocalizedTagNames() []string {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("standardLocalizedTagNames"))
	return objc.ConvertSliceToStrings(rv)
}

// An array of all tags the audio unit associates with the current user, and
// the system tags the audio units define.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponentManager/tagNames
func (a AVAudioUnitComponentManager) TagNames() []string {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("tagNames"))
	return objc.ConvertSliceToStrings(rv)
}
