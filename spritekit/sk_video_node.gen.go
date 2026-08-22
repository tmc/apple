// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKVideoNode] class.
var (
	_SKVideoNodeClass     SKVideoNodeClass
	_SKVideoNodeClassOnce sync.Once
)

func getSKVideoNodeClass() SKVideoNodeClass {
	_SKVideoNodeClassOnce.Do(func() {
		_SKVideoNodeClass = SKVideoNodeClass{class: objc.GetClass("SKVideoNode")}
	})
	return _SKVideoNodeClass
}

// GetSKVideoNodeClass returns the class object for SKVideoNode.
func GetSKVideoNodeClass() SKVideoNodeClass {
	return getSKVideoNodeClass()
}

type SKVideoNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKVideoNodeClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKVideoNodeClass) Alloc() SKVideoNode {
	rv := objc.Send[SKVideoNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A graphical element that plays video content.
//
// # Overview
//
// This class renders a video at a given size and location in your scene with
// no exposed player controls.
//
// # Creating a Video Node
//
//   - [SKVideoNode.InitWithAVPlayer]: Initializes a video node using an existing [AVPlayer](<https://developer.apple.com/documentation/AVFoundation/AVPlayer>) object.
//   - [SKVideoNode.InitWithFileNamed]: Initializes a video node using a video file stored in the app bundle.
//   - [SKVideoNode.InitWithURL]: Initializes a video node using a URL.
//
// # Setting the Video Node’s Visual Properties
//
//   - [SKVideoNode.AnchorPoint]: The point in the sprite that corresponds to the node’s position.
//   - [SKVideoNode.SetAnchorPoint]
//   - [SKVideoNode.Size]: The dimensions of the video node, in points.
//   - [SKVideoNode.SetSize]
//
// # Controlling Video Playback
//
//   - [SKVideoNode.Play]: Starts video playback.
//   - [SKVideoNode.Pause]: Pauses video playback.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKVideoNode
type SKVideoNode struct {
	SKNode
}

// SKVideoNodeFromID constructs a [SKVideoNode] from an objc.ID.
//
// A graphical element that plays video content.
func SKVideoNodeFromID(id objc.ID) SKVideoNode {
	return SKVideoNode{SKNode: SKNodeFromID(id)}
}

// NOTE: SKVideoNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKVideoNode] class.
//
// # Creating a Video Node
//
//   - [ISKVideoNode.InitWithAVPlayer]: Initializes a video node using an existing [AVPlayer](<https://developer.apple.com/documentation/AVFoundation/AVPlayer>) object.
//   - [ISKVideoNode.InitWithFileNamed]: Initializes a video node using a video file stored in the app bundle.
//   - [ISKVideoNode.InitWithURL]: Initializes a video node using a URL.
//
// # Setting the Video Node’s Visual Properties
//
//   - [ISKVideoNode.AnchorPoint]: The point in the sprite that corresponds to the node’s position.
//   - [ISKVideoNode.SetAnchorPoint]
//   - [ISKVideoNode.Size]: The dimensions of the video node, in points.
//   - [ISKVideoNode.SetSize]
//
// # Controlling Video Playback
//
//   - [ISKVideoNode.Play]: Starts video playback.
//   - [ISKVideoNode.Pause]: Pauses video playback.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKVideoNode
type ISKVideoNode interface {
	ISKNode

	// Topic: Creating a Video Node

	// Initializes a video node using an existing [AVPlayer](<https://developer.apple.com/documentation/AVFoundation/AVPlayer>) object.
	InitWithAVPlayer(player objectivec.IObject) SKVideoNode
	// Initializes a video node using a video file stored in the app bundle.
	InitWithFileNamed(videoFile string) SKVideoNode
	// Initializes a video node using a URL.
	InitWithURL(url foundation.NSURL) SKVideoNode

	// Topic: Setting the Video Node’s Visual Properties

	// The point in the sprite that corresponds to the node’s position.
	AnchorPoint() corefoundation.CGPoint
	SetAnchorPoint(value corefoundation.CGPoint)
	// The dimensions of the video node, in points.
	Size() corefoundation.CGSize
	SetSize(value corefoundation.CGSize)

	// Topic: Controlling Video Playback

	// Starts video playback.
	Play()
	// Pauses video playback.
	Pause()
}

// Init initializes the instance.
func (v SKVideoNode) Init() SKVideoNode {
	rv := objc.Send[SKVideoNode](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v SKVideoNode) Autorelease() SKVideoNode {
	rv := objc.Send[SKVideoNode](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKVideoNode creates a new SKVideoNode instance.
func NewSKVideoNode() SKVideoNode {
	class := getSKVideoNodeClass()
	rv := objc.Send[SKVideoNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a video node using an existing [AVPlayer] object.
//
// player: A player object.
//
// # Return Value
//
// An initialized video node.
//
// # Discussion
//
// You can use the [AVPlayer] object to control playback.
//
// Listing 1 shows, in Swift, how you can create a video node using the
// [SKVideoNode.InitWithAVPlayer] initializer.
//
// Listing 1. Creating a video node with an AV Player
//
// See: https://developer.apple.com/documentation/SpriteKit/SKVideoNode/init(avPlayer:)
//
// [AVPlayer]: https://developer.apple.com/documentation/AVFoundation/AVPlayer
//
// [AVPlayer]: https://developer.apple.com/documentation/AVFoundation/AVPlayer
func NewVideoNodeWithAVPlayer(player objectivec.IObject) SKVideoNode {
	instance := getSKVideoNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAVPlayer:"), player)
	return SKVideoNodeFromID(rv)
}

// Tells you when to initialize a video node that was created from an archive.
//
// # Discussion
//
// Do not call this initializer yourself; it is called by the system when you
// should intialize a video node that was created from an archive.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKVideoNode/init(coder:)
func NewVideoNodeWithCoder(aDecoder foundation.INSCoder) SKVideoNode {
	instance := getSKVideoNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return SKVideoNodeFromID(rv)
}

// Initializes a video node using a video file stored in the app bundle.
//
// videoFile: The name of the video file.
//
// # Return Value
//
// An initialized video node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKVideoNode/init(fileNamed:)
func NewVideoNodeWithFileNamed(videoFile string) SKVideoNode {
	instance := getSKVideoNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFileNamed:"), objc.String(videoFile))
	return SKVideoNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(fileNamed:securelyWithClasses:)
func NewVideoNodeWithFileNamedSecurelyWithClassesAndError(filename string, classes foundation.INSSet) (SKVideoNode, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getSKVideoNodeClass().class), objc.Sel("nodeWithFileNamed:securelyWithClasses:andError:"), objc.String(filename), classes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SKVideoNode{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SKVideoNode{}, objc.ErrInitFailed
	}
	return SKVideoNodeFromID(rv), nil
}

// Initializes a video node using a URL.
//
// url: The URL for the video to play.
//
// # Return Value
//
// An initialized video node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKVideoNode/init(url:)
func NewVideoNodeWithURL(url foundation.NSURL) SKVideoNode {
	instance := getSKVideoNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), url)
	return SKVideoNodeFromID(rv)
}

// Initializes a video node using an existing [AVPlayer] object.
//
// player: A player object.
//
// player is a [*avfoundation.AVPlayer].
//
// # Return Value
//
// An initialized video node.
//
// # Discussion
//
// You can use the [AVPlayer] object to control playback.
//
// Listing 1 shows, in Swift, how you can create a video node using the
// [SKVideoNode.InitWithAVPlayer] initializer.
//
// Listing 1. Creating a video node with an AV Player
//
// See: https://developer.apple.com/documentation/SpriteKit/SKVideoNode/init(avPlayer:)
//
// [AVPlayer]: https://developer.apple.com/documentation/AVFoundation/AVPlayer
//
// [AVPlayer]: https://developer.apple.com/documentation/AVFoundation/AVPlayer
func (v SKVideoNode) InitWithAVPlayer(player objectivec.IObject) SKVideoNode {
	rv := objc.Send[SKVideoNode](v.ID, objc.Sel("initWithAVPlayer:"), player)
	return rv
}

// Initializes a video node using a video file stored in the app bundle.
//
// videoFile: The name of the video file.
//
// # Return Value
//
// An initialized video node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKVideoNode/init(fileNamed:)
func (v SKVideoNode) InitWithFileNamed(videoFile string) SKVideoNode {
	rv := objc.Send[SKVideoNode](v.ID, objc.Sel("initWithFileNamed:"), objc.String(videoFile))
	return rv
}

// Initializes a video node using a URL.
//
// url: The URL for the video to play.
//
// # Return Value
//
// An initialized video node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKVideoNode/init(url:)
func (v SKVideoNode) InitWithURL(url foundation.NSURL) SKVideoNode {
	rv := objc.Send[SKVideoNode](v.ID, objc.Sel("initWithURL:"), url)
	return rv
}

// Starts video playback.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKVideoNode/play()
func (v SKVideoNode) Play() {
	objc.Send[objc.ID](v.ID, objc.Sel("play"))
}

// Pauses video playback.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKVideoNode/pause()
func (v SKVideoNode) Pause() {
	objc.Send[objc.ID](v.ID, objc.Sel("pause"))
}

// The point in the sprite that corresponds to the node’s position.
//
// # Discussion
//
// You specify the anchor point using the unit coordinate space. The default
// value is `(0.5,0.5)`, which means that the video is centered on the
// node’s position.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKVideoNode/anchorPoint
func (v SKVideoNode) AnchorPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](v.ID, objc.Sel("anchorPoint"))
	return corefoundation.CGPoint(rv)
}
func (v SKVideoNode) SetAnchorPoint(value corefoundation.CGPoint) {
	objc.Send[struct{}](v.ID, objc.Sel("setAnchorPoint:"), value)
}

// The dimensions of the video node, in points.
//
// # Discussion
//
// The default value is the size of the video used to instantiate the node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKVideoNode/size
func (v SKVideoNode) Size() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v.ID, objc.Sel("size"))
	return corefoundation.CGSize(rv)
}
func (v SKVideoNode) SetSize(value corefoundation.CGSize) {
	objc.Send[struct{}](v.ID, objc.Sel("setSize:"), value)
}
