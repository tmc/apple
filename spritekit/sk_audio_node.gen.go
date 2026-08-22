// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/avfaudio"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [SKAudioNode] class.
var (
	_SKAudioNodeClass     SKAudioNodeClass
	_SKAudioNodeClassOnce sync.Once
)

func getSKAudioNodeClass() SKAudioNodeClass {
	_SKAudioNodeClassOnce.Do(func() {
		_SKAudioNodeClass = SKAudioNodeClass{class: objc.GetClass("SKAudioNode")}
	})
	return _SKAudioNodeClass
}

// GetSKAudioNodeClass returns the class object for SKAudioNode.
func GetSKAudioNodeClass() SKAudioNodeClass {
	return getSKAudioNodeClass()
}

type SKAudioNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKAudioNodeClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKAudioNodeClass) Alloc() SKAudioNode {
	rv := objc.Send[SKAudioNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A node that plays audio.
//
// # Overview
//
// A [SKAudioNode] object is used to add an audio to a scene. The sounds are
// played automatically using AVFoundation, and the node can optionally add 3D
// spatial audio effects to the audio when it is played.
//
// The currently presented [SKScene] object mixes the audio from nodes in the
// scene based on parameters defined in the [AVAudio3DMixing] protocol. A
// scene’s [SKScene.AudioEngine] property allows overall control of volume
// and playback.
//
// By default, [SKAudioNode] objects are positional, i.e. their
// [SKAudioNode.Positional] property is set to true. If you add an audio node
// to a scene with a [SKScene.Listener] set, SpriteKit will set the stereo
// balance and the volume based on the relative positions of the two nodes.
//
// You can explicitly set the volume or stereo balance to an audio node by
// running actions on it.
//
// SpriteKit includes actions that reduce an audio node’s volume by changing
// either its occlusion or obstruction. The difference between these actions
// is that occlusion affects both the direct and reverb paths of the sound
// while obstruction only affects the direct path. The change volume action
// offers absolute control over an audio node’s volume.
//
// You can manually set the stereo balance of an audio node with a stereo pan
// action.
//
// Special effects, such as speeding up or slowing down audio by changing the
// playback rate and adding reverb are also available as audio actions.
//
// To learn more about audio actions, see Controlling the Audio of a Node in
// [Action Initializers].
//
// # Initializing Audio Nodes
//
//   - [SKAudioNode.InitWithAVAudioNode]: Initializes an audio node from an AVFoundation audio node.
//   - [SKAudioNode.InitWithFileNamed]: Initializes an audio node from an audio asset with the specified filename.
//   - [SKAudioNode.InitWithURL]: Initializes an audio node from an audio asset with the specified URL.
//
// # Configuring Audio Nodes
//
//   - [SKAudioNode.AvAudioNode]: The audio node’s current audio asset.
//   - [SKAudioNode.SetAvAudioNode]
//   - [SKAudioNode.IsPositional]: A Boolean property that indicates whether the node’s audio is altered based on the position of the node.
//   - [SKAudioNode.SetPositional]
//   - [SKAudioNode.AutoplayLooped]: A Boolean value that indicates whether the audio should play in a loop when the node is added to the scene.
//   - [SKAudioNode.SetAutoplayLooped]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAudioNode
//
// [AVAudio3DMixing]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing
// [Action Initializers]: https://developer.apple.com/documentation/SpriteKit/action-initializers
type SKAudioNode struct {
	SKNode
}

// SKAudioNodeFromID constructs a [SKAudioNode] from an objc.ID.
//
// A node that plays audio.
func SKAudioNodeFromID(id objc.ID) SKAudioNode {
	return SKAudioNode{SKNode: SKNodeFromID(id)}
}

// NOTE: SKAudioNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKAudioNode] class.
//
// # Initializing Audio Nodes
//
//   - [ISKAudioNode.InitWithAVAudioNode]: Initializes an audio node from an AVFoundation audio node.
//   - [ISKAudioNode.InitWithFileNamed]: Initializes an audio node from an audio asset with the specified filename.
//   - [ISKAudioNode.InitWithURL]: Initializes an audio node from an audio asset with the specified URL.
//
// # Configuring Audio Nodes
//
//   - [ISKAudioNode.AvAudioNode]: The audio node’s current audio asset.
//   - [ISKAudioNode.SetAvAudioNode]
//   - [ISKAudioNode.IsPositional]: A Boolean property that indicates whether the node’s audio is altered based on the position of the node.
//   - [ISKAudioNode.SetPositional]
//   - [ISKAudioNode.AutoplayLooped]: A Boolean value that indicates whether the audio should play in a loop when the node is added to the scene.
//   - [ISKAudioNode.SetAutoplayLooped]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAudioNode
type ISKAudioNode interface {
	ISKNode

	// Topic: Initializing Audio Nodes

	// Initializes an audio node from an AVFoundation audio node.
	InitWithAVAudioNode(node avfaudio.AVAudioNode) SKAudioNode
	// Initializes an audio node from an audio asset with the specified filename.
	InitWithFileNamed(name string) SKAudioNode
	// Initializes an audio node from an audio asset with the specified URL.
	InitWithURL(url foundation.NSURL) SKAudioNode

	// Topic: Configuring Audio Nodes

	// The audio node’s current audio asset.
	AvAudioNode() avfaudio.AVAudioNode
	SetAvAudioNode(value avfaudio.AVAudioNode)
	// A Boolean property that indicates whether the node’s audio is altered based on the position of the node.
	IsPositional() bool
	SetPositional(value bool)
	// A Boolean value that indicates whether the audio should play in a loop when the node is added to the scene.
	AutoplayLooped() bool
	SetAutoplayLooped(value bool)
}

// Init initializes the instance.
func (a SKAudioNode) Init() SKAudioNode {
	rv := objc.Send[SKAudioNode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a SKAudioNode) Autorelease() SKAudioNode {
	rv := objc.Send[SKAudioNode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKAudioNode creates a new SKAudioNode instance.
func NewSKAudioNode() SKAudioNode {
	class := getSKAudioNodeClass()
	rv := objc.Send[SKAudioNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes an audio node from an AVFoundation audio node.
//
// node: An [AVAudioNode] object that holds an [AVAudioEngine] sound graph from a
// single sound source or URL.
//
// # Return Value
//
// A newly initialized audio node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAudioNode/init(avAudioNode:)
//
// [AVAudioEngine]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine
// [AVAudioNode]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode
func NewAudioNodeWithAVAudioNode(node avfaudio.AVAudioNode) SKAudioNode {
	instance := getSKAudioNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAVAudioNode:"), node)
	return SKAudioNodeFromID(rv)
}

// Tells you when to initialize an audio node that has been unarchived.
//
// # Discussion
//
// Do not call this initializer directly; it’s called by the system when you
// should initialize an audio node that has been unarchived.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAudioNode/init(coder:)
func NewAudioNodeWithCoder(aDecoder foundation.INSCoder) SKAudioNode {
	instance := getSKAudioNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return SKAudioNodeFromID(rv)
}

// Initializes an audio node from an audio asset with the specified filename.
//
// name: A file containing an [AVAudioNode].
//
// # Return Value
//
// A newly initialized audio node.
//
// # Discussion
//
// The named file containing the audio asset must reside within the main
// bundle.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAudioNode/init(fileNamed:)
//
// [AVAudioNode]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode
func NewAudioNodeWithFileNamed(name string) SKAudioNode {
	instance := getSKAudioNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFileNamed:"), objc.String(name))
	return SKAudioNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(fileNamed:securelyWithClasses:)
func NewAudioNodeWithFileNamedSecurelyWithClassesAndError(filename string, classes foundation.INSSet) (SKAudioNode, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getSKAudioNodeClass().class), objc.Sel("nodeWithFileNamed:securelyWithClasses:andError:"), objc.String(filename), classes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SKAudioNode{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SKAudioNode{}, objc.ErrInitFailed
	}
	return SKAudioNodeFromID(rv), nil
}

// Initializes an audio node from an audio asset with the specified URL.
//
// url: The URL of an audio file.
//
// # Return Value
//
// A newly initialized audio node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAudioNode/init(url:)
func NewAudioNodeWithURL(url foundation.NSURL) SKAudioNode {
	instance := getSKAudioNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), url)
	return SKAudioNodeFromID(rv)
}

// Initializes an audio node from an AVFoundation audio node.
//
// node: An [AVAudioNode] object that holds an [AVAudioEngine] sound graph from a
// single sound source or URL.
//
// # Return Value
//
// A newly initialized audio node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAudioNode/init(avAudioNode:)
//
// [AVAudioEngine]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine
// [AVAudioNode]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode
func (a SKAudioNode) InitWithAVAudioNode(node avfaudio.AVAudioNode) SKAudioNode {
	rv := objc.Send[SKAudioNode](a.ID, objc.Sel("initWithAVAudioNode:"), node)
	return rv
}

// Initializes an audio node from an audio asset with the specified filename.
//
// name: A file containing an [AVAudioNode].
//
// # Return Value
//
// A newly initialized audio node.
//
// # Discussion
//
// The named file containing the audio asset must reside within the main
// bundle.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAudioNode/init(fileNamed:)
//
// [AVAudioNode]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode
func (a SKAudioNode) InitWithFileNamed(name string) SKAudioNode {
	rv := objc.Send[SKAudioNode](a.ID, objc.Sel("initWithFileNamed:"), objc.String(name))
	return rv
}

// Initializes an audio node from an audio asset with the specified URL.
//
// url: The URL of an audio file.
//
// # Return Value
//
// A newly initialized audio node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAudioNode/init(url:)
func (a SKAudioNode) InitWithURL(url foundation.NSURL) SKAudioNode {
	rv := objc.Send[SKAudioNode](a.ID, objc.Sel("initWithURL:"), url)
	return rv
}

// The audio node’s current audio asset.
//
// # Discussion
//
// The AV audio node must refer to an [AVAudioEngine] sound graph from a
// single sound source or URL.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAudioNode/avAudioNode
//
// [AVAudioEngine]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine
func (a SKAudioNode) AvAudioNode() avfaudio.AVAudioNode {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("avAudioNode"))
	return avfaudio.AVAudioNodeFromID(objc.ID(rv))
}
func (a SKAudioNode) SetAvAudioNode(value avfaudio.AVAudioNode) {
	objc.Send[struct{}](a.ID, objc.Sel("setAvAudioNode:"), value)
}

// A Boolean property that indicates whether the node’s audio is altered
// based on the position of the node.
//
// # Discussion
//
// If true, the audio mixer considers the position and velocity of the
// [SKAudioNode] relative to scene’s current [SKScene.Listener] node. The
// mixer applies distance attenuation, doppler shift, and pan effects to the
// sound. If false, then the sound is played normally. The default value is
// true.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAudioNode/isPositional
func (a SKAudioNode) IsPositional() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isPositional"))
	return rv
}
func (a SKAudioNode) SetPositional(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setPositional:"), value)
}

// A Boolean value that indicates whether the audio should play in a loop when
// the node is added to the scene.
//
// # Discussion
//
// If the property value is true, then the audio starts playing as soon as the
// node is added to the scene, and repeats after it completes. If false, then
// the audio node’s content never plays automatically. It must be explicitly
// scheduled using the scene’s audio engine. The default value is true.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAudioNode/autoplayLooped
func (a SKAudioNode) AutoplayLooped() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("autoplayLooped"))
	return rv
}
func (a SKAudioNode) SetAutoplayLooped(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setAutoplayLooped:"), value)
}
