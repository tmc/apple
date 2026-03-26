// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/quartzcore"
)

// The class instance for the [AVPlayerLayer] class.
var (
	_AVPlayerLayerClass     AVPlayerLayerClass
	_AVPlayerLayerClassOnce sync.Once
)

func getAVPlayerLayerClass() AVPlayerLayerClass {
	_AVPlayerLayerClassOnce.Do(func() {
		_AVPlayerLayerClass = AVPlayerLayerClass{class: objc.GetClass("AVPlayerLayer")}
	})
	return _AVPlayerLayerClass
}

// GetAVPlayerLayerClass returns the class object for AVPlayerLayer.
func GetAVPlayerLayerClass() AVPlayerLayerClass {
	return getAVPlayerLayerClass()
}

type AVPlayerLayerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerLayerClass) Alloc() AVPlayerLayer {
	rv := objc.Send[AVPlayerLayer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that presents the visual contents of a player object.
//
// # Overview
// 
// A common way to use this object in iOS or tvOS is as the backing layer for
// a [UIView], as the following example shows:
//
// [UIView]: https://developer.apple.com/documentation/UIKit/UIView
//
// # Configuring the presentation
//
//   - [AVPlayerLayer.VideoRect]: The current size and position of the video image that displays within the layer’s bounds.
//   - [AVPlayerLayer.VideoGravity]: A value that specifies how the layer displays the player’s visual content within the layer’s bounds.
//   - [AVPlayerLayer.SetVideoGravity]
//
// # Determining display readiness
//
//   - [AVPlayerLayer.ReadyForDisplay]: A Boolean value that indicates whether the first video frame of the player’s current item is ready for display.
//
// # Accessing the player
//
//   - [AVPlayerLayer.Player]: The player whose visual content the layer displays.
//   - [AVPlayerLayer.SetPlayer]
//
// # Processing pixel buffers
//
//   - [AVPlayerLayer.PixelBufferAttributes]: The attributes of the visual output that displays in the player layer during playback.
//   - [AVPlayerLayer.SetPixelBufferAttributes]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerLayer
type AVPlayerLayer struct {
	quartzcore.CALayer
}

// AVPlayerLayerFromID constructs a [AVPlayerLayer] from an objc.ID.
//
// An object that presents the visual contents of a player object.
func AVPlayerLayerFromID(id objc.ID) AVPlayerLayer {
	return AVPlayerLayer{CALayer: quartzcore.CALayerFromID(id)}
}
// NOTE: AVPlayerLayer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerLayer] class.
//
// # Configuring the presentation
//
//   - [IAVPlayerLayer.VideoRect]: The current size and position of the video image that displays within the layer’s bounds.
//   - [IAVPlayerLayer.VideoGravity]: A value that specifies how the layer displays the player’s visual content within the layer’s bounds.
//   - [IAVPlayerLayer.SetVideoGravity]
//
// # Determining display readiness
//
//   - [IAVPlayerLayer.ReadyForDisplay]: A Boolean value that indicates whether the first video frame of the player’s current item is ready for display.
//
// # Accessing the player
//
//   - [IAVPlayerLayer.Player]: The player whose visual content the layer displays.
//   - [IAVPlayerLayer.SetPlayer]
//
// # Processing pixel buffers
//
//   - [IAVPlayerLayer.PixelBufferAttributes]: The attributes of the visual output that displays in the player layer during playback.
//   - [IAVPlayerLayer.SetPixelBufferAttributes]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerLayer
type IAVPlayerLayer interface {
	quartzcore.ICALayer

	// Topic: Configuring the presentation

	// The current size and position of the video image that displays within the layer’s bounds.
	VideoRect() corefoundation.CGRect
	// A value that specifies how the layer displays the player’s visual content within the layer’s bounds.
	VideoGravity() AVLayerVideoGravity
	SetVideoGravity(value AVLayerVideoGravity)

	// Topic: Determining display readiness

	// A Boolean value that indicates whether the first video frame of the player’s current item is ready for display.
	ReadyForDisplay() bool

	// Topic: Accessing the player

	// The player whose visual content the layer displays.
	Player() IAVPlayer
	SetPlayer(value IAVPlayer)

	// Topic: Processing pixel buffers

	// The attributes of the visual output that displays in the player layer during playback.
	PixelBufferAttributes() foundation.INSDictionary
	SetPixelBufferAttributes(value foundation.INSDictionary)
}

// Init initializes the instance.
func (p AVPlayerLayer) Init() AVPlayerLayer {
	rv := objc.Send[AVPlayerLayer](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerLayer) Autorelease() AVPlayerLayer {
	rv := objc.Send[AVPlayerLayer](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerLayer creates a new AVPlayerLayer instance.
func NewAVPlayerLayer() AVPlayerLayer {
	class := getAVPlayerLayerClass()
	rv := objc.Send[AVPlayerLayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a layer object to present the visual contents of a player’s
// current item.
//
// player: The player whose visual contents the layer presents.
//
// # Return Value
// 
// A layer that displays the visual output of the associated player.
//
// # Discussion
// 
// You may create an arbitrary number of layers for the same player object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerLayer/init(player:)
func NewPlayerLayerWithPlayer(player IAVPlayer) AVPlayerLayer {
	rv := objc.Send[objc.ID](objc.ID(getAVPlayerLayerClass().class), objc.Sel("playerLayerWithPlayer:"), player)
	return AVPlayerLayerFromID(rv)
}

// The current size and position of the video image that displays within the
// layer’s bounds.
//
// # Discussion
// 
// The size and position of a rectangle depends on the aspect ratio of the
// media (16:9 or 4:3), the layer’s [bounds], and the value of its
// [VideoGravity] property.
// 
// This property is key-value observable.
//
// [bounds]: https://developer.apple.com/documentation/QuartzCore/CALayer/bounds
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerLayer/videoRect
func (p AVPlayerLayer) VideoRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p.ID, objc.Sel("videoRect"))
	return corefoundation.CGRect(rv)
}
// A value that specifies how the layer displays the player’s visual content
// within the layer’s bounds.
//
// # Discussion
// 
// A player layer supports the following video gravity values:
// 
// - [resizeAspect]
// - [resizeAspectFill]
// - [resize]
// 
// The default value is [resizeAspect].
// 
// This property is animatable.
//
// [resizeAspectFill]: https://developer.apple.com/documentation/AVFoundation/AVLayerVideoGravity/resizeAspectFill
// [resizeAspect]: https://developer.apple.com/documentation/AVFoundation/AVLayerVideoGravity/resizeAspect
// [resize]: https://developer.apple.com/documentation/AVFoundation/AVLayerVideoGravity/resize
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerLayer/videoGravity
func (p AVPlayerLayer) VideoGravity() AVLayerVideoGravity {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("videoGravity"))
	return AVLayerVideoGravity(foundation.NSStringFromID(rv).String())
}
func (p AVPlayerLayer) SetVideoGravity(value AVLayerVideoGravity) {
	objc.Send[struct{}](p.ID, objc.Sel("setVideoGravity:"), objc.String(string(value)))
}
// A Boolean value that indicates whether the first video frame of the
// player’s current item is ready for display.
//
// # Discussion
// 
// Use this property to determine when to show or animate a player layer into
// view. You can display a player layer while this property value is [false],
// but the layer doesn’t present any content until the value becomes [true].
// 
// This property remains [false] for a player when its [CurrentItem] contains
// no enabled video tracks.
// 
// This property is key-value observable.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerLayer/isReadyForDisplay
func (p AVPlayerLayer) ReadyForDisplay() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isReadyForDisplay"))
	return rv
}
// The player whose visual content the layer displays.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerLayer/player
func (p AVPlayerLayer) Player() IAVPlayer {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("player"))
	return AVPlayerFromID(objc.ID(rv))
}
func (p AVPlayerLayer) SetPlayer(value IAVPlayer) {
	objc.Send[struct{}](p.ID, objc.Sel("setPlayer:"), value)
}
// The attributes of the visual output that displays in the player layer
// during playback.
//
// # Discussion
// 
// Use this property to customize the format of the pixel buffers that the
// player layer vends.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerLayer/pixelBufferAttributes
func (p AVPlayerLayer) PixelBufferAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("pixelBufferAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (p AVPlayerLayer) SetPixelBufferAttributes(value foundation.INSDictionary) {
	objc.Send[struct{}](p.ID, objc.Sel("setPixelBufferAttributes:"), value)
}

