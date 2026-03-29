// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioChannelLayout] class.
var (
	_AVAudioChannelLayoutClass     AVAudioChannelLayoutClass
	_AVAudioChannelLayoutClassOnce sync.Once
)

func getAVAudioChannelLayoutClass() AVAudioChannelLayoutClass {
	_AVAudioChannelLayoutClassOnce.Do(func() {
		_AVAudioChannelLayoutClass = AVAudioChannelLayoutClass{class: objc.GetClass("AVAudioChannelLayout")}
	})
	return _AVAudioChannelLayoutClass
}

// GetAVAudioChannelLayoutClass returns the class object for AVAudioChannelLayout.
func GetAVAudioChannelLayoutClass() AVAudioChannelLayoutClass {
	return getAVAudioChannelLayoutClass()
}

type AVAudioChannelLayoutClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioChannelLayoutClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioChannelLayoutClass) Alloc() AVAudioChannelLayout {
	rv := objc.Send[AVAudioChannelLayout](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that describes the roles of a set of audio channels.
//
// # Overview
// 
// The [AVAudioChannelLayout] class is a thin wrapper for Core Audio’s
// [AudioChannelLayout].
//
// [AudioChannelLayout]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelLayout
//
// # Creating an Audio Channel Layout
//
//   - [AVAudioChannelLayout.InitWithLayout]: Creates an audio channel layout object from an existing one.
//   - [AVAudioChannelLayout.InitWithLayoutTag]: Creates an audio channel layout object from a layout tag.
//
// # Getting Audio Channel Layout Properties
//
//   - [AVAudioChannelLayout.ChannelCount]: The number of channels of audio data.
//   - [AVAudioChannelLayout.Layout]: The underlying audio channel layout.
//   - [AVAudioChannelLayout.LayoutTag]: The audio channel’s underlying layout tag.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout
type AVAudioChannelLayout struct {
	objectivec.Object
}

// AVAudioChannelLayoutFromID constructs a [AVAudioChannelLayout] from an objc.ID.
//
// An object that describes the roles of a set of audio channels.
func AVAudioChannelLayoutFromID(id objc.ID) AVAudioChannelLayout {
	return AVAudioChannelLayout{objectivec.Object{ID: id}}
}
// NOTE: AVAudioChannelLayout adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioChannelLayout] class.
//
// # Creating an Audio Channel Layout
//
//   - [IAVAudioChannelLayout.InitWithLayout]: Creates an audio channel layout object from an existing one.
//   - [IAVAudioChannelLayout.InitWithLayoutTag]: Creates an audio channel layout object from a layout tag.
//
// # Getting Audio Channel Layout Properties
//
//   - [IAVAudioChannelLayout.ChannelCount]: The number of channels of audio data.
//   - [IAVAudioChannelLayout.Layout]: The underlying audio channel layout.
//   - [IAVAudioChannelLayout.LayoutTag]: The audio channel’s underlying layout tag.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout
type IAVAudioChannelLayout interface {
	objectivec.IObject

	// Topic: Creating an Audio Channel Layout

	// Creates an audio channel layout object from an existing one.
	InitWithLayout(layout IAVAudioChannelLayout) AVAudioChannelLayout
	// Creates an audio channel layout object from a layout tag.
	InitWithLayoutTag(layoutTag objectivec.IObject) AVAudioChannelLayout

	// Topic: Getting Audio Channel Layout Properties

	// The number of channels of audio data.
	ChannelCount() AVAudioChannelCount
	// The underlying audio channel layout.
	Layout() IAVAudioChannelLayout
	// The audio channel’s underlying layout tag.
	LayoutTag() objectivec.IObject

	AVChannelLayoutKey() string
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (a AVAudioChannelLayout) Init() AVAudioChannelLayout {
	rv := objc.Send[AVAudioChannelLayout](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioChannelLayout) Autorelease() AVAudioChannelLayout {
	rv := objc.Send[AVAudioChannelLayout](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioChannelLayout creates a new AVAudioChannelLayout instance.
func NewAVAudioChannelLayout() AVAudioChannelLayout {
	class := getAVAudioChannelLayoutClass()
	rv := objc.Send[AVAudioChannelLayout](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an audio channel layout object from an existing one.
//
// layout: The existing audio channel layout object.
//
// # Return Value
// 
// A new [AVAudioChannelLayout] object.
//
// # Discussion
// 
// If the audio channel layout object’s tag is
// [kAudioChannelLayoutTag_UseChannelDescriptions], this initializer attempts
// to convert it to a more specific tag.
//
// [kAudioChannelLayoutTag_UseChannelDescriptions]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_UseChannelDescriptions
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/init(layout:)
func NewAudioChannelLayoutWithLayout(layout IAVAudioChannelLayout) AVAudioChannelLayout {
	instance := getAVAudioChannelLayoutClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLayout:"), layout)
	return AVAudioChannelLayoutFromID(rv)
}

// Creates an audio channel layout object from a layout tag.
//
// layoutTag: The audio channel layout tag.
//
// # Return Value
// 
// A new [AVAudioChannelLayout] object, or `nil` if `layoutTag` is
// [kAudioChannelLayoutTag_UseChannelDescriptions] or
// [kAudioChannelLayoutTag_UseChannelBitmap].
//
// [kAudioChannelLayoutTag_UseChannelBitmap]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_UseChannelBitmap
// [kAudioChannelLayoutTag_UseChannelDescriptions]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_UseChannelDescriptions
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/init(layoutTag:)
// layoutTag is a [coreaudiotypes.AudioChannelLayoutTag].
func NewAudioChannelLayoutWithLayoutTag(layoutTag objectivec.IObject) AVAudioChannelLayout {
	instance := getAVAudioChannelLayoutClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLayoutTag:"), layoutTag)
	return AVAudioChannelLayoutFromID(rv)
}

// Creates an audio channel layout object from an existing one.
//
// layout: The existing audio channel layout object.
//
// # Return Value
// 
// A new [AVAudioChannelLayout] object.
//
// # Discussion
// 
// If the audio channel layout object’s tag is
// [kAudioChannelLayoutTag_UseChannelDescriptions], this initializer attempts
// to convert it to a more specific tag.
//
// [kAudioChannelLayoutTag_UseChannelDescriptions]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_UseChannelDescriptions
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/init(layout:)
func (a AVAudioChannelLayout) InitWithLayout(layout IAVAudioChannelLayout) AVAudioChannelLayout {
	rv := objc.Send[AVAudioChannelLayout](a.ID, objc.Sel("initWithLayout:"), layout)
	return rv
}
// Creates an audio channel layout object from a layout tag.
//
// layoutTag: The audio channel layout tag.
//
// layoutTag is a [coreaudiotypes.AudioChannelLayoutTag].
//
// # Return Value
// 
// A new [AVAudioChannelLayout] object, or `nil` if `layoutTag` is
// [kAudioChannelLayoutTag_UseChannelDescriptions] or
// [kAudioChannelLayoutTag_UseChannelBitmap].
//
// [kAudioChannelLayoutTag_UseChannelBitmap]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_UseChannelBitmap
// [kAudioChannelLayoutTag_UseChannelDescriptions]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_UseChannelDescriptions
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/init(layoutTag:)
// layoutTag is a [coreaudiotypes.AudioChannelLayoutTag].
func (a AVAudioChannelLayout) InitWithLayoutTag(layoutTag objectivec.IObject) AVAudioChannelLayout {
	rv := objc.Send[AVAudioChannelLayout](a.ID, objc.Sel("initWithLayoutTag:"), layoutTag)
	return rv
}
func (a AVAudioChannelLayout) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates an audio channel layout object from an existing one.
//
// layout: The existing audio channel layout object.
//
// # Return Value
// 
// A new [AVAudioChannelLayout] object.
//
// # Discussion
// 
// If the layout’s tag is [kAudioChannelLayoutTag_UseChannelDescriptions],
// the method attempts to convert it to a more specific tag.
//
// [kAudioChannelLayoutTag_UseChannelDescriptions]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_UseChannelDescriptions
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/layoutWithLayout:
func (_AVAudioChannelLayoutClass AVAudioChannelLayoutClass) LayoutWithLayout(layout IAVAudioChannelLayout) AVAudioChannelLayout {
	rv := objc.Send[objc.ID](objc.ID(_AVAudioChannelLayoutClass.class), objc.Sel("layoutWithLayout:"), layout)
	return AVAudioChannelLayoutFromID(rv)
}
// Creates an audio channel layout object from an audio channel layout tag.
//
// layoutTag: The audio channel layout tag.
//
// layoutTag is a [coreaudiotypes.AudioChannelLayoutTag].
//
// # Return Value
// 
// A new [AVAudioChannelLayout] object.
//
// # Discussion
// 
// If the provided audio channel layout object’s tag is
// [kAudioChannelLayoutTag_UseChannelDescriptions], this initializer attempts
// to convert it to a more specific tag.
//
// [kAudioChannelLayoutTag_UseChannelDescriptions]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_UseChannelDescriptions
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/layoutWithLayoutTag:
// layoutTag is a [coreaudiotypes.AudioChannelLayoutTag].
func (_AVAudioChannelLayoutClass AVAudioChannelLayoutClass) LayoutWithLayoutTag(layoutTag objectivec.IObject) AVAudioChannelLayout {
	rv := objc.Send[objc.ID](objc.ID(_AVAudioChannelLayoutClass.class), objc.Sel("layoutWithLayoutTag:"), layoutTag)
	return AVAudioChannelLayoutFromID(rv)
}

// The number of channels of audio data.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/channelCount
func (a AVAudioChannelLayout) ChannelCount() AVAudioChannelCount {
	rv := objc.Send[AVAudioChannelCount](a.ID, objc.Sel("channelCount"))
	return AVAudioChannelCount(rv)
}
// The underlying audio channel layout.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/layout
func (a AVAudioChannelLayout) Layout() IAVAudioChannelLayout {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("layout"))
	return AVAudioChannelLayoutFromID(objc.ID(rv))
}
// The audio channel’s underlying layout tag.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/layoutTag
func (a AVAudioChannelLayout) LayoutTag() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("layoutTag"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/avfaudio/avchannellayoutkey
func (a AVAudioChannelLayout) AVChannelLayoutKey() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("AVChannelLayoutKey"))
	return foundation.NSStringFromID(rv).String()
}

