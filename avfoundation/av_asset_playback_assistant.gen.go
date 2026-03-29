// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetPlaybackAssistant] class.
var (
	_AVAssetPlaybackAssistantClass     AVAssetPlaybackAssistantClass
	_AVAssetPlaybackAssistantClassOnce sync.Once
)

func getAVAssetPlaybackAssistantClass() AVAssetPlaybackAssistantClass {
	_AVAssetPlaybackAssistantClassOnce.Do(func() {
		_AVAssetPlaybackAssistantClass = AVAssetPlaybackAssistantClass{class: objc.GetClass("AVAssetPlaybackAssistant")}
	})
	return _AVAssetPlaybackAssistantClass
}

// GetAVAssetPlaybackAssistantClass returns the class object for AVAssetPlaybackAssistant.
func GetAVAssetPlaybackAssistantClass() AVAssetPlaybackAssistantClass {
	return getAVAssetPlaybackAssistantClass()
}

type AVAssetPlaybackAssistantClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetPlaybackAssistantClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetPlaybackAssistantClass) Alloc() AVAssetPlaybackAssistant {
	rv := objc.Send[AVAssetPlaybackAssistant](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that provides playback information for an asset.
//
// # Loading playback configuration options
//
//   - [AVAssetPlaybackAssistant.LoadPlaybackConfigurationOptionsWithCompletionHandler]: Loads playback configuration options for an asset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetPlaybackAssistant
type AVAssetPlaybackAssistant struct {
	objectivec.Object
}

// AVAssetPlaybackAssistantFromID constructs a [AVAssetPlaybackAssistant] from an objc.ID.
//
// An object that provides playback information for an asset.
func AVAssetPlaybackAssistantFromID(id objc.ID) AVAssetPlaybackAssistant {
	return AVAssetPlaybackAssistant{objectivec.Object{ID: id}}
}
// NOTE: AVAssetPlaybackAssistant adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetPlaybackAssistant] class.
//
// # Loading playback configuration options
//
//   - [IAVAssetPlaybackAssistant.LoadPlaybackConfigurationOptionsWithCompletionHandler]: Loads playback configuration options for an asset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetPlaybackAssistant
type IAVAssetPlaybackAssistant interface {
	objectivec.IObject

	// Topic: Loading playback configuration options

	// Loads playback configuration options for an asset.
	LoadPlaybackConfigurationOptionsWithCompletionHandler(completionHandler VoidHandler)
}

// Init initializes the instance.
func (a AVAssetPlaybackAssistant) Init() AVAssetPlaybackAssistant {
	rv := objc.Send[AVAssetPlaybackAssistant](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetPlaybackAssistant) Autorelease() AVAssetPlaybackAssistant {
	rv := objc.Send[AVAssetPlaybackAssistant](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetPlaybackAssistant creates a new AVAssetPlaybackAssistant instance.
func NewAVAssetPlaybackAssistant() AVAssetPlaybackAssistant {
	class := getAVAssetPlaybackAssistantClass()
	rv := objc.Send[AVAssetPlaybackAssistant](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a playback assistant to inspect the specified asset.
//
// asset: An asset instance to inspect.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetPlaybackAssistant/init(asset:)
func NewAssetPlaybackAssistantWithAsset(asset IAVAsset) AVAssetPlaybackAssistant {
	rv := objc.Send[objc.ID](objc.ID(getAVAssetPlaybackAssistantClass().class), objc.Sel("assetPlaybackAssistantWithAsset:"), asset)
	return AVAssetPlaybackAssistantFromID(rv)
}

// Loads playback configuration options for an asset.
//
// completionHandler: A callback the system invokes with an array of
// [AVAssetPlaybackConfigurationOption] values that describe capabilities of
// the asset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetPlaybackAssistant/loadPlaybackConfigurationOptions(completionHandler:)
func (a AVAssetPlaybackAssistant) LoadPlaybackConfigurationOptionsWithCompletionHandler(completionHandler VoidHandler) {
_block0, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("loadPlaybackConfigurationOptionsWithCompletionHandler:"), _block0)
}

// LoadPlaybackConfigurationOptions is a synchronous wrapper around [AVAssetPlaybackAssistant.LoadPlaybackConfigurationOptionsWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAssetPlaybackAssistant) LoadPlaybackConfigurationOptions(ctx context.Context) error {
	done := make(chan struct{}, 1)
	a.LoadPlaybackConfigurationOptionsWithCompletionHandler(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

