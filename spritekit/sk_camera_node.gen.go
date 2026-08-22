// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [SKCameraNode] class.
var (
	_SKCameraNodeClass     SKCameraNodeClass
	_SKCameraNodeClassOnce sync.Once
)

func getSKCameraNodeClass() SKCameraNodeClass {
	_SKCameraNodeClassOnce.Do(func() {
		_SKCameraNodeClass = SKCameraNodeClass{class: objc.GetClass("SKCameraNode")}
	})
	return _SKCameraNodeClass
}

// GetSKCameraNodeClass returns the class object for SKCameraNode.
func GetSKCameraNodeClass() SKCameraNodeClass {
	return getSKCameraNodeClass()
}

type SKCameraNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKCameraNodeClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKCameraNodeClass) Alloc() SKCameraNode {
	rv := objc.Send[SKCameraNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A node that determines which parts of the scene are visible within a view.
//
// # Overview
//
// If you don’t use a camera in your scene, you control the visible portion
// of a scene using its [SKScene.AnchorPoint] property.
//
// # Node Visibility
//
//   - [SKCameraNode.ContainedNodeSet]: Finds nodes that are visible in the camera’s viewport.
//   - [SKCameraNode.ContainsNode]: Checks to see if a node is visible in the camera’s viewport.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKCameraNode
type SKCameraNode struct {
	SKNode
}

// SKCameraNodeFromID constructs a [SKCameraNode] from an objc.ID.
//
// A node that determines which parts of the scene are visible within a view.
func SKCameraNodeFromID(id objc.ID) SKCameraNode {
	return SKCameraNode{SKNode: SKNodeFromID(id)}
}

// NOTE: SKCameraNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKCameraNode] class.
//
// # Node Visibility
//
//   - [ISKCameraNode.ContainedNodeSet]: Finds nodes that are visible in the camera’s viewport.
//   - [ISKCameraNode.ContainsNode]: Checks to see if a node is visible in the camera’s viewport.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKCameraNode
type ISKCameraNode interface {
	ISKNode

	// Topic: Node Visibility

	// Finds nodes that are visible in the camera’s viewport.
	ContainedNodeSet() foundation.INSSet
	// Checks to see if a node is visible in the camera’s viewport.
	ContainsNode(node ISKNode) bool
}

// Init initializes the instance.
func (c SKCameraNode) Init() SKCameraNode {
	rv := objc.Send[SKCameraNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c SKCameraNode) Autorelease() SKCameraNode {
	rv := objc.Send[SKCameraNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKCameraNode creates a new SKCameraNode instance.
func NewSKCameraNode() SKCameraNode {
	class := getSKCameraNodeClass()
	rv := objc.Send[SKCameraNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Called when a node is initialized from an .sks file.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(coder:)
func NewCameraNodeWithCoder(aDecoder foundation.INSCoder) SKCameraNode {
	instance := getSKCameraNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return SKCameraNodeFromID(rv)
}

// Creates a new node by loading an archive file from the game’s main
// bundle.
//
// filename: The name of the file, without a file extension. The file must be in the
// app’s main bundle and have a `XCUIElementTypeSks` filename extension.
//
// # Return Value
//
// The unarchived node object.
//
// # Discussion
//
// If you call this method on a subclass of the [SKScene] class and the object
// in the archive is an [SKScene] object, the returned object is initialized
// as if it is a member of the subclass. You use this behavior to create scene
// layouts in the Xcode Editor and provide custom behaviors in your subclass.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(fileNamed:)
func NewCameraNodeWithFileNamed(filename string) SKCameraNode {
	rv := objc.Send[objc.ID](objc.ID(getSKCameraNodeClass().class), objc.Sel("nodeWithFileNamed:"), objc.String(filename))
	return SKCameraNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(fileNamed:securelyWithClasses:)
func NewCameraNodeWithFileNamedSecurelyWithClassesAndError(filename string, classes foundation.INSSet) (SKCameraNode, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getSKCameraNodeClass().class), objc.Sel("nodeWithFileNamed:securelyWithClasses:andError:"), objc.String(filename), classes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SKCameraNode{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SKCameraNode{}, objc.ErrInitFailed
	}
	return SKCameraNodeFromID(rv), nil
}

// Finds nodes that are visible in the camera’s viewport.
//
// # Return Value
//
// The set of nodes that are in the same scene as the camera and contained in
// the camera’s viewport.
//
// # Discussion
//
// The camera must be part of a scene’s node hierarchy and the scene must be
// presented in an view.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKCameraNode/containedNodeSet()
func (c SKCameraNode) ContainedNodeSet() foundation.INSSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("containedNodeSet"))
	return foundation.NSSetFromID(rv)
}

// Checks to see if a node is visible in the camera’s viewport.
//
// node: An [SKNode] object.
//
// # Return Value
//
// true if the node in the same scene and inside the camera’s viewport;
// otherwise false.
//
// # Discussion
//
// The camera must be part of a scene’s node hierarchy and the scene must be
// presented in an view.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKCameraNode/contains(_:)
func (c SKCameraNode) ContainsNode(node ISKNode) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("containsNode:"), node)
	return rv
}
