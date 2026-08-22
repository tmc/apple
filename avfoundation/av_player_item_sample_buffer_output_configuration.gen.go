// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlayerItemSampleBufferOutputConfiguration] class.
var (
	_AVPlayerItemSampleBufferOutputConfigurationClass     AVPlayerItemSampleBufferOutputConfigurationClass
	_AVPlayerItemSampleBufferOutputConfigurationClassOnce sync.Once
)

func getAVPlayerItemSampleBufferOutputConfigurationClass() AVPlayerItemSampleBufferOutputConfigurationClass {
	_AVPlayerItemSampleBufferOutputConfigurationClassOnce.Do(func() {
		_AVPlayerItemSampleBufferOutputConfigurationClass = AVPlayerItemSampleBufferOutputConfigurationClass{class: objc.GetClass("AVPlayerItemSampleBufferOutputConfiguration")}
	})
	return _AVPlayerItemSampleBufferOutputConfigurationClass
}

// GetAVPlayerItemSampleBufferOutputConfigurationClass returns the class object for AVPlayerItemSampleBufferOutputConfiguration.
func GetAVPlayerItemSampleBufferOutputConfigurationClass() AVPlayerItemSampleBufferOutputConfigurationClass {
	return getAVPlayerItemSampleBufferOutputConfigurationClass()
}

type AVPlayerItemSampleBufferOutputConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlayerItemSampleBufferOutputConfigurationClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerItemSampleBufferOutputConfigurationClass) Alloc() AVPlayerItemSampleBufferOutputConfiguration {
	rv := objc.Send[AVPlayerItemSampleBufferOutputConfiguration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// Configuration options specified when creating an
// [AVPlayerItemSampleBufferOutput].
//
// # Overview
//
// Mutating [AVPlayerItemSampleBufferOutputConfiguration] after using it to
// create an [AVPlayerItemSampleBufferOutput] object will not affect the
// [AVPlayerItemSampleBufferOutput] object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSampleBufferOutputConfiguration
type AVPlayerItemSampleBufferOutputConfiguration struct {
	objectivec.Object
}

// AVPlayerItemSampleBufferOutputConfigurationFromID constructs a [AVPlayerItemSampleBufferOutputConfiguration] from an objc.ID.
//
// Configuration options specified when creating an
// [AVPlayerItemSampleBufferOutput].
func AVPlayerItemSampleBufferOutputConfigurationFromID(id objc.ID) AVPlayerItemSampleBufferOutputConfiguration {
	return AVPlayerItemSampleBufferOutputConfiguration{objectivec.Object{ID: id}}
}

// NOTE: AVPlayerItemSampleBufferOutputConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerItemSampleBufferOutputConfiguration] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSampleBufferOutputConfiguration
type IAVPlayerItemSampleBufferOutputConfiguration interface {
	objectivec.IObject
}

// Init initializes the instance.
func (p AVPlayerItemSampleBufferOutputConfiguration) Init() AVPlayerItemSampleBufferOutputConfiguration {
	rv := objc.Send[AVPlayerItemSampleBufferOutputConfiguration](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerItemSampleBufferOutputConfiguration) Autorelease() AVPlayerItemSampleBufferOutputConfiguration {
	rv := objc.Send[AVPlayerItemSampleBufferOutputConfiguration](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerItemSampleBufferOutputConfiguration creates a new AVPlayerItemSampleBufferOutputConfiguration instance.
func NewAVPlayerItemSampleBufferOutputConfiguration() AVPlayerItemSampleBufferOutputConfiguration {
	class := getAVPlayerItemSampleBufferOutputConfigurationClass()
	rv := objc.Send[AVPlayerItemSampleBufferOutputConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}
