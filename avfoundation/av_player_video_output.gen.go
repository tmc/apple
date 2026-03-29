// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlayerVideoOutput] class.
var (
	_AVPlayerVideoOutputClass     AVPlayerVideoOutputClass
	_AVPlayerVideoOutputClassOnce sync.Once
)

func getAVPlayerVideoOutputClass() AVPlayerVideoOutputClass {
	_AVPlayerVideoOutputClassOnce.Do(func() {
		_AVPlayerVideoOutputClass = AVPlayerVideoOutputClass{class: objc.GetClass("AVPlayerVideoOutput")}
	})
	return _AVPlayerVideoOutputClass
}

// GetAVPlayerVideoOutputClass returns the class object for AVPlayerVideoOutput.
func GetAVPlayerVideoOutputClass() AVPlayerVideoOutputClass {
	return getAVPlayerVideoOutputClass()
}

type AVPlayerVideoOutputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlayerVideoOutputClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerVideoOutputClass) Alloc() AVPlayerVideoOutput {
	rv := objc.Send[AVPlayerVideoOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that receives video data from a player object.
//
// # Overview
// 
// Attach a video output to an [AVPlayer] object to access the player’s
// video data as [CMTaggedBufferGroupRef] objects.
//
// [CMTaggedBufferGroupRef]: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupRef
//
// # Creating an output
//
//   - [AVPlayerVideoOutput.InitWithSpecification]: Creates a video output from a specification.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutput
type AVPlayerVideoOutput struct {
	objectivec.Object
}

// AVPlayerVideoOutputFromID constructs a [AVPlayerVideoOutput] from an objc.ID.
//
// An object that receives video data from a player object.
func AVPlayerVideoOutputFromID(id objc.ID) AVPlayerVideoOutput {
	return AVPlayerVideoOutput{objectivec.Object{ID: id}}
}
// NOTE: AVPlayerVideoOutput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerVideoOutput] class.
//
// # Creating an output
//
//   - [IAVPlayerVideoOutput.InitWithSpecification]: Creates a video output from a specification.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutput
type IAVPlayerVideoOutput interface {
	objectivec.IObject

	// Topic: Creating an output

	// Creates a video output from a specification.
	InitWithSpecification(specification IAVVideoOutputSpecification) AVPlayerVideoOutput

	CopyTaggedBufferGroupForHostTimePresentationTimeStampActiveConfiguration(hostTime coremedia.CMTime, presentationTimeStampOut *coremedia.CMTime, activeConfigurationOut *AVPlayerVideoOutputConfiguration) coremedia.CMTaggedBufferGroupRef
}

// Init initializes the instance.
func (p AVPlayerVideoOutput) Init() AVPlayerVideoOutput {
	rv := objc.Send[AVPlayerVideoOutput](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerVideoOutput) Autorelease() AVPlayerVideoOutput {
	rv := objc.Send[AVPlayerVideoOutput](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerVideoOutput creates a new AVPlayerVideoOutput instance.
func NewAVPlayerVideoOutput() AVPlayerVideoOutput {
	class := getAVPlayerVideoOutputClass()
	rv := objc.Send[AVPlayerVideoOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a video output from a specification.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutput/init(specification:)
func NewPlayerVideoOutputWithSpecification(specification IAVVideoOutputSpecification) AVPlayerVideoOutput {
	instance := getAVPlayerVideoOutputClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpecification:"), specification)
	return AVPlayerVideoOutputFromID(rv)
}

// Creates a video output from a specification.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutput/init(specification:)
func (p AVPlayerVideoOutput) InitWithSpecification(specification IAVVideoOutputSpecification) AVPlayerVideoOutput {
	rv := objc.Send[AVPlayerVideoOutput](p.ID, objc.Sel("initWithSpecification:"), specification)
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutput/copyTaggedBufferGroupForHostTime:presentationTimeStamp:activeConfiguration:
func (p AVPlayerVideoOutput) CopyTaggedBufferGroupForHostTimePresentationTimeStampActiveConfiguration(hostTime coremedia.CMTime, presentationTimeStampOut *coremedia.CMTime, activeConfigurationOut *AVPlayerVideoOutputConfiguration) coremedia.CMTaggedBufferGroupRef {
	rv := objc.Send[coremedia.CMTaggedBufferGroupRef](p.ID, objc.Sel("copyTaggedBufferGroupForHostTime:presentationTimeStamp:activeConfiguration:"), hostTime, presentationTimeStampOut, activeConfigurationOut)
	return coremedia.CMTaggedBufferGroupRef(rv)
}

