// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/quartzcore"
)

// The class instance for the [AVSynchronizedLayer] class.
var (
	_AVSynchronizedLayerClass     AVSynchronizedLayerClass
	_AVSynchronizedLayerClassOnce sync.Once
)

func getAVSynchronizedLayerClass() AVSynchronizedLayerClass {
	_AVSynchronizedLayerClassOnce.Do(func() {
		_AVSynchronizedLayerClass = AVSynchronizedLayerClass{class: objc.GetClass("AVSynchronizedLayer")}
	})
	return _AVSynchronizedLayerClass
}

// GetAVSynchronizedLayerClass returns the class object for AVSynchronizedLayer.
func GetAVSynchronizedLayerClass() AVSynchronizedLayerClass {
	return getAVSynchronizedLayerClass()
}

type AVSynchronizedLayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVSynchronizedLayerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVSynchronizedLayerClass) Alloc() AVSynchronizedLayer {
	rv := objc.Send[AVSynchronizedLayer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A Core Animation layer that derives its timing from a player item so that
// you can synchronize layer animations with media playback.
//
// # Overview
//
// You can create an arbitrary number of synchronized layers from the same
// [AVPlayerItem] object.
//
// A synchronized layer is similar to a [CATransformLayer] object in that it
// doesn’t display anything itself, it just confers state upon its layer
// subtree. [AVSynchronizedLayer] confers its timing state, synchronizing the
// timing of layers in its subtree with that of a player item.
//
// Any [CoreAnimation] layer with animation property set that is added as a
// sublayer of [AVSynchronizedLayer] should set the animation’s [AVSynchronizedLayer.BeginTime]
// property to a non-zero positive value so animations will be interpreted on
// the player item’s timeline. [CoreAnimation] replaces the default
// `beginTime` of 0.0 with [CACurrentMediaTime()]. To start the animation from
// time 0, use a small positive value like [AVSynchronizedLayer.AVCoreAnimationBeginTimeAtZero].
//
// You might use a layer as shown in the following example:
//
// # Managing the player item
//
//   - [AVSynchronizedLayer.PlayerItem]: The player item to which the timing of the layer is synchronized.
//   - [AVSynchronizedLayer.SetPlayerItem]
//
// # Supporting types
//
//   - [AVSynchronizedLayer.AVCoreAnimationBeginTimeAtZero]: A value that sets an animation begin time to
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSynchronizedLayer
//
// [CACurrentMediaTime()]: https://developer.apple.com/documentation/QuartzCore/CACurrentMediaTime()
// [CATransformLayer]: https://developer.apple.com/documentation/QuartzCore/CATransformLayer
type AVSynchronizedLayer struct {
	quartzcore.CALayer
}

// AVSynchronizedLayerFromID constructs a [AVSynchronizedLayer] from an objc.ID.
//
// A Core Animation layer that derives its timing from a player item so that
// you can synchronize layer animations with media playback.
func AVSynchronizedLayerFromID(id objc.ID) AVSynchronizedLayer {
	return AVSynchronizedLayer{CALayer: quartzcore.CALayerFromID(id)}
}

// NOTE: AVSynchronizedLayer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVSynchronizedLayer] class.
//
// # Managing the player item
//
//   - [IAVSynchronizedLayer.PlayerItem]: The player item to which the timing of the layer is synchronized.
//   - [IAVSynchronizedLayer.SetPlayerItem]
//
// # Supporting types
//
//   - [IAVSynchronizedLayer.AVCoreAnimationBeginTimeAtZero]: A value that sets an animation begin time to
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSynchronizedLayer
type IAVSynchronizedLayer interface {
	quartzcore.ICALayer

	// Topic: Managing the player item

	// The player item to which the timing of the layer is synchronized.
	PlayerItem() IAVPlayerItem
	SetPlayerItem(value IAVPlayerItem)

	// Topic: Supporting types

	// A value that sets an animation begin time to
	AVCoreAnimationBeginTimeAtZero() float64
}

// Init initializes the instance.
func (s AVSynchronizedLayer) Init() AVSynchronizedLayer {
	rv := objc.Send[AVSynchronizedLayer](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s AVSynchronizedLayer) Autorelease() AVSynchronizedLayer {
	rv := objc.Send[AVSynchronizedLayer](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSynchronizedLayer creates a new AVSynchronizedLayer instance.
func NewAVSynchronizedLayer() AVSynchronizedLayer {
	class := getAVSynchronizedLayerClass()
	rv := objc.Send[AVSynchronizedLayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new synchronized layer with timing synchronized with a given
// player item.
//
// playerItem: A player item.
//
// # Return Value
//
// A new synchronized layer with timing synchronized with `playerItem`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSynchronizedLayer/init(playerItem:)
func NewSynchronizedLayerWithPlayerItem(playerItem IAVPlayerItem) AVSynchronizedLayer {
	rv := objc.Send[objc.ID](objc.ID(getAVSynchronizedLayerClass().class), objc.Sel("synchronizedLayerWithPlayerItem:"), playerItem)
	return AVSynchronizedLayerFromID(rv)
}

// The player item to which the timing of the layer is synchronized.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSynchronizedLayer/playerItem
func (s AVSynchronizedLayer) PlayerItem() IAVPlayerItem {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("playerItem"))
	return AVPlayerItemFromID(objc.ID(rv))
}
func (s AVSynchronizedLayer) SetPlayerItem(value IAVPlayerItem) {
	objc.Send[struct{}](s.ID, objc.Sel("setPlayerItem:"), value)
}

// A value that sets an animation begin time to
//
// See: https://developer.apple.com/documentation/avfoundation/avcoreanimationbegintimeatzero
func (s AVSynchronizedLayer) AVCoreAnimationBeginTimeAtZero() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("AVCoreAnimationBeginTimeAtZero"))
	return rv
}
