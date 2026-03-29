// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlayerVideoOutputConfiguration] class.
var (
	_AVPlayerVideoOutputConfigurationClass     AVPlayerVideoOutputConfigurationClass
	_AVPlayerVideoOutputConfigurationClassOnce sync.Once
)

func getAVPlayerVideoOutputConfigurationClass() AVPlayerVideoOutputConfigurationClass {
	_AVPlayerVideoOutputConfigurationClassOnce.Do(func() {
		_AVPlayerVideoOutputConfigurationClass = AVPlayerVideoOutputConfigurationClass{class: objc.GetClass("AVPlayerVideoOutputConfiguration")}
	})
	return _AVPlayerVideoOutputConfigurationClass
}

// GetAVPlayerVideoOutputConfigurationClass returns the class object for AVPlayerVideoOutputConfiguration.
func GetAVPlayerVideoOutputConfigurationClass() AVPlayerVideoOutputConfigurationClass {
	return getAVPlayerVideoOutputConfigurationClass()
}

type AVPlayerVideoOutputConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlayerVideoOutputConfigurationClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerVideoOutputConfigurationClass) Alloc() AVPlayerVideoOutputConfiguration {
	rv := objc.Send[AVPlayerVideoOutputConfiguration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that provides configuration information for the related player
// item.
//
// # Inspecting the configuration
//
//   - [AVPlayerVideoOutputConfiguration.SourcePlayerItem]: The player item that’s the source of this configuration.
//   - [AVPlayerVideoOutputConfiguration.DataChannelDescription]: An array of data channels selected for this configuration.
//   - [AVPlayerVideoOutputConfiguration.SetDataChannelDescription]
//   - [AVPlayerVideoOutputConfiguration.ActivationTime]: The host time this configuration became active on its associated player object.
//   - [AVPlayerVideoOutputConfiguration.PreferredTransform]: The preferred transform of the visual media.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutput/Configuration
type AVPlayerVideoOutputConfiguration struct {
	objectivec.Object
}

// AVPlayerVideoOutputConfigurationFromID constructs a [AVPlayerVideoOutputConfiguration] from an objc.ID.
//
// An object that provides configuration information for the related player
// item.
func AVPlayerVideoOutputConfigurationFromID(id objc.ID) AVPlayerVideoOutputConfiguration {
	return AVPlayerVideoOutputConfiguration{objectivec.Object{ID: id}}
}
// NOTE: AVPlayerVideoOutputConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerVideoOutputConfiguration] class.
//
// # Inspecting the configuration
//
//   - [IAVPlayerVideoOutputConfiguration.SourcePlayerItem]: The player item that’s the source of this configuration.
//   - [IAVPlayerVideoOutputConfiguration.DataChannelDescription]: An array of data channels selected for this configuration.
//   - [IAVPlayerVideoOutputConfiguration.SetDataChannelDescription]
//   - [IAVPlayerVideoOutputConfiguration.ActivationTime]: The host time this configuration became active on its associated player object.
//   - [IAVPlayerVideoOutputConfiguration.PreferredTransform]: The preferred transform of the visual media.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutput/Configuration
type IAVPlayerVideoOutputConfiguration interface {
	objectivec.IObject

	// Topic: Inspecting the configuration

	// The player item that’s the source of this configuration.
	SourcePlayerItem() IAVPlayerItem
	// An array of data channels selected for this configuration.
	DataChannelDescription() objectivec.IObject
	SetDataChannelDescription(value objectivec.IObject)
	// The host time this configuration became active on its associated player object.
	ActivationTime() uintptr
	// The preferred transform of the visual media.
	PreferredTransform() corefoundation.CGAffineTransform

	// An array of data channels selected for this configuration.
	DataChannelDescriptions() foundation.INSArray
}

// Init initializes the instance.
func (p AVPlayerVideoOutputConfiguration) Init() AVPlayerVideoOutputConfiguration {
	rv := objc.Send[AVPlayerVideoOutputConfiguration](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerVideoOutputConfiguration) Autorelease() AVPlayerVideoOutputConfiguration {
	rv := objc.Send[AVPlayerVideoOutputConfiguration](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerVideoOutputConfiguration creates a new AVPlayerVideoOutputConfiguration instance.
func NewAVPlayerVideoOutputConfiguration() AVPlayerVideoOutputConfiguration {
	class := getAVPlayerVideoOutputConfigurationClass()
	rv := objc.Send[AVPlayerVideoOutputConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The player item that’s the source of this configuration.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutput/Configuration/sourcePlayerItem
func (p AVPlayerVideoOutputConfiguration) SourcePlayerItem() IAVPlayerItem {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("sourcePlayerItem"))
	return AVPlayerItemFromID(objc.ID(rv))
}
// An array of data channels selected for this configuration.
//
// See: https://developer.apple.com/documentation/avfoundation/avplayervideooutput/configuration/datachanneldescription
func (p AVPlayerVideoOutputConfiguration) DataChannelDescription() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("dataChannelDescription"))
	return objectivec.Object{ID: rv}
}
func (p AVPlayerVideoOutputConfiguration) SetDataChannelDescription(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setDataChannelDescription:"), value)
}
// The host time this configuration became active on its associated player
// object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutput/Configuration/activationTime
func (p AVPlayerVideoOutputConfiguration) ActivationTime() uintptr {
	rv := objc.Send[uintptr](p.ID, objc.Sel("activationTime"))
	return rv
}
// The preferred transform of the visual media.
//
// # Discussion
// 
// The system retrieves the transform from the [AVAssetTrack] that provides
// the media data. If the source track doesn’t specify a transform, the
// value of this property is [CGAffineTransformIdentity].
//
// [CGAffineTransformIdentity]: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformIdentity
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutput/Configuration/preferredTransform
func (p AVPlayerVideoOutputConfiguration) PreferredTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](p.ID, objc.Sel("preferredTransform"))
	return corefoundation.CGAffineTransform(rv)
}
// An array of data channels selected for this configuration.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerVideoOutputConfiguration/dataChannelDescriptions
func (p AVPlayerVideoOutputConfiguration) DataChannelDescriptions() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("dataChannelDescriptions"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

