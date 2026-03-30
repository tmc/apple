// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetWriterInputGroup] class.
var (
	_AVAssetWriterInputGroupClass     AVAssetWriterInputGroupClass
	_AVAssetWriterInputGroupClassOnce sync.Once
)

func getAVAssetWriterInputGroupClass() AVAssetWriterInputGroupClass {
	_AVAssetWriterInputGroupClassOnce.Do(func() {
		_AVAssetWriterInputGroupClass = AVAssetWriterInputGroupClass{class: objc.GetClass("AVAssetWriterInputGroup")}
	})
	return _AVAssetWriterInputGroupClass
}

// GetAVAssetWriterInputGroupClass returns the class object for AVAssetWriterInputGroup.
func GetAVAssetWriterInputGroupClass() AVAssetWriterInputGroupClass {
	return getAVAssetWriterInputGroupClass()
}

type AVAssetWriterInputGroupClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetWriterInputGroupClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetWriterInputGroupClass) Alloc() AVAssetWriterInputGroup {
	rv := objc.Send[AVAssetWriterInputGroup](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A group of inputs with tracks that are mutually exclusive to each other for
// playback or processing.
//
// # Overview
//
// Assets may contain multiple tracks of media that are mutually exclusive to
// each other when you play or process them. For example, an asset may contain
// multiple audio tracks for different spoken languages, but only one of them
// should play at a time. You use an input group to mark a collection of
// tracks as mutually exclusive to each other in the file the asset writer
// outputs.
//
// # Creating an input group
//
//   - [AVAssetWriterInputGroup.InitWithInputsDefaultInput]: Creates a group for the asset writer inputs.
//
// # Accessing the inputs
//
//   - [AVAssetWriterInputGroup.Inputs]: The inputs with tracks that are mutually exclusive to each other for playback or processing.
//   - [AVAssetWriterInputGroup.DefaultInput]: The default input for the group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputGroup
type AVAssetWriterInputGroup struct {
	AVMediaSelectionGroup
}

// AVAssetWriterInputGroupFromID constructs a [AVAssetWriterInputGroup] from an objc.ID.
//
// A group of inputs with tracks that are mutually exclusive to each other for
// playback or processing.
func AVAssetWriterInputGroupFromID(id objc.ID) AVAssetWriterInputGroup {
	return AVAssetWriterInputGroup{AVMediaSelectionGroup: AVMediaSelectionGroupFromID(id)}
}

// NOTE: AVAssetWriterInputGroup adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetWriterInputGroup] class.
//
// # Creating an input group
//
//   - [IAVAssetWriterInputGroup.InitWithInputsDefaultInput]: Creates a group for the asset writer inputs.
//
// # Accessing the inputs
//
//   - [IAVAssetWriterInputGroup.Inputs]: The inputs with tracks that are mutually exclusive to each other for playback or processing.
//   - [IAVAssetWriterInputGroup.DefaultInput]: The default input for the group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputGroup
type IAVAssetWriterInputGroup interface {
	IAVMediaSelectionGroup

	// Topic: Creating an input group

	// Creates a group for the asset writer inputs.
	InitWithInputsDefaultInput(inputs []AVAssetWriterInput, defaultInput IAVAssetWriterInput) AVAssetWriterInputGroup

	// Topic: Accessing the inputs

	// The inputs with tracks that are mutually exclusive to each other for playback or processing.
	Inputs() []AVAssetWriterInput
	// The default input for the group.
	DefaultInput() IAVAssetWriterInput
}

// Init initializes the instance.
func (a AVAssetWriterInputGroup) Init() AVAssetWriterInputGroup {
	rv := objc.Send[AVAssetWriterInputGroup](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetWriterInputGroup) Autorelease() AVAssetWriterInputGroup {
	rv := objc.Send[AVAssetWriterInputGroup](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetWriterInputGroup creates a new AVAssetWriterInputGroup instance.
func NewAVAssetWriterInputGroup() AVAssetWriterInputGroup {
	class := getAVAssetWriterInputGroupClass()
	rv := objc.Send[AVAssetWriterInputGroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a group for the asset writer inputs.
//
// inputs: The inputs with tracks to arrange into a mutually exclusive group.
//
// defaultInput: The group’s default input.
//
// # Discussion
//
// When you add an input group to an asset writer, the system sets the default
// input’s [MarksOutputTrackAsEnabled] property value to true, and the value
// of the other inputs in the group to false.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputGroup/init(inputs:defaultInput:)
func NewAssetWriterInputGroupWithInputsDefaultInput(inputs []AVAssetWriterInput, defaultInput IAVAssetWriterInput) AVAssetWriterInputGroup {
	instance := getAVAssetWriterInputGroupClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInputs:defaultInput:"), objectivec.IObjectSliceToNSArray(inputs), defaultInput)
	return AVAssetWriterInputGroupFromID(rv)
}

// Creates a group for the asset writer inputs.
//
// inputs: The inputs with tracks to arrange into a mutually exclusive group.
//
// defaultInput: The group’s default input.
//
// # Discussion
//
// When you add an input group to an asset writer, the system sets the default
// input’s [MarksOutputTrackAsEnabled] property value to true, and the value
// of the other inputs in the group to false.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputGroup/init(inputs:defaultInput:)
func (a AVAssetWriterInputGroup) InitWithInputsDefaultInput(inputs []AVAssetWriterInput, defaultInput IAVAssetWriterInput) AVAssetWriterInputGroup {
	rv := objc.Send[AVAssetWriterInputGroup](a.ID, objc.Sel("initWithInputs:defaultInput:"), objectivec.IObjectSliceToNSArray(inputs), defaultInput)
	return rv
}

// Returns a new group for the asset writer inputs.
//
// inputs: The inputs with tracks to arrange into a mutually exclusive group.
//
// defaultInput: The group’s default input.
//
// # Return Value
//
// An asset writer input group.
//
// # Discussion
//
// When you add an input group to an asset writer, the system sets the default
// input’s [MarksOutputTrackAsEnabled] property value to true, and the value
// of the other inputs in the group to false.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputGroup/assetWriterInputGroupWithInputs:defaultInput:
func (_AVAssetWriterInputGroupClass AVAssetWriterInputGroupClass) AssetWriterInputGroupWithInputsDefaultInput(inputs []AVAssetWriterInput, defaultInput IAVAssetWriterInput) AVAssetWriterInputGroup {
	rv := objc.Send[objc.ID](objc.ID(_AVAssetWriterInputGroupClass.class), objc.Sel("assetWriterInputGroupWithInputs:defaultInput:"), objectivec.IObjectSliceToNSArray(inputs), defaultInput)
	return AVAssetWriterInputGroupFromID(rv)
}

// The inputs with tracks that are mutually exclusive to each other for
// playback or processing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputGroup/inputs
func (a AVAssetWriterInputGroup) Inputs() []AVAssetWriterInput {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("inputs"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVAssetWriterInput {
		return AVAssetWriterInputFromID(id)
	})
}

// The default input for the group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputGroup/defaultInput
func (a AVAssetWriterInputGroup) DefaultInput() IAVAssetWriterInput {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("defaultInput"))
	return AVAssetWriterInputFromID(objc.ID(rv))
}
