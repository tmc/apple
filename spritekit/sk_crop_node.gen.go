// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [SKCropNode] class.
var (
	_SKCropNodeClass     SKCropNodeClass
	_SKCropNodeClassOnce sync.Once
)

func getSKCropNodeClass() SKCropNodeClass {
	_SKCropNodeClassOnce.Do(func() {
		_SKCropNodeClass = SKCropNodeClass{class: objc.GetClass("SKCropNode")}
	})
	return _SKCropNodeClass
}

// GetSKCropNodeClass returns the class object for SKCropNode.
func GetSKCropNodeClass() SKCropNodeClass {
	return getSKCropNodeClass()
}

type SKCropNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKCropNodeClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKCropNodeClass) Alloc() SKCropNode {
	rv := objc.Send[SKCropNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A node that masks pixels drawn by its children so that only some pixels are
// seen.
//
// # Overview
//
// [SKCropNode] is a container node that you use to crop other nodes in the
// scene. You add other nodes to a crop node and set the crop node’s
// [SKCropNode.MaskNode] property. For example, here are some ways you might
// specify a mask:
//
// - An untextured sprite that limits content to a rectangular portion of the
// scene. - A textured sprite that works as a precise per-pixel mask. - A
// collection of child nodes that form a unique shape.
//
// You can animate the shape or contents of the mask to implement interesting
// effects such as hiding or revealing.
//
// # Setting the Mask Filter
//
//   - [SKCropNode.MaskNode]: The node used to determine the crop node’s mask.
//   - [SKCropNode.SetMaskNode]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKCropNode
type SKCropNode struct {
	SKNode
}

// SKCropNodeFromID constructs a [SKCropNode] from an objc.ID.
//
// A node that masks pixels drawn by its children so that only some pixels are
// seen.
func SKCropNodeFromID(id objc.ID) SKCropNode {
	return SKCropNode{SKNode: SKNodeFromID(id)}
}

// NOTE: SKCropNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKCropNode] class.
//
// # Setting the Mask Filter
//
//   - [ISKCropNode.MaskNode]: The node used to determine the crop node’s mask.
//   - [ISKCropNode.SetMaskNode]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKCropNode
type ISKCropNode interface {
	ISKNode

	// Topic: Setting the Mask Filter

	// The node used to determine the crop node’s mask.
	MaskNode() ISKNode
	SetMaskNode(value ISKNode)
}

// Init initializes the instance.
func (c SKCropNode) Init() SKCropNode {
	rv := objc.Send[SKCropNode](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c SKCropNode) Autorelease() SKCropNode {
	rv := objc.Send[SKCropNode](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKCropNode creates a new SKCropNode instance.
func NewSKCropNode() SKCropNode {
	class := getSKCropNodeClass()
	rv := objc.Send[SKCropNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Called when a node is initialized from an .sks file.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(coder:)
func NewCropNodeWithCoder(aDecoder foundation.INSCoder) SKCropNode {
	instance := getSKCropNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return SKCropNodeFromID(rv)
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
func NewCropNodeWithFileNamed(filename string) SKCropNode {
	rv := objc.Send[objc.ID](objc.ID(getSKCropNodeClass().class), objc.Sel("nodeWithFileNamed:"), objc.String(filename))
	return SKCropNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(fileNamed:securelyWithClasses:)
func NewCropNodeWithFileNamedSecurelyWithClassesAndError(filename string, classes foundation.INSSet) (SKCropNode, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getSKCropNodeClass().class), objc.Sel("nodeWithFileNamed:securelyWithClasses:andError:"), objc.String(filename), classes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SKCropNode{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SKCropNode{}, objc.ErrInitFailed
	}
	return SKCropNodeFromID(rv), nil
}

// The node used to determine the crop node’s mask.
//
// # Discussion
//
// The node supplied to the crop node must not be a child of another node;
// however, it may have children of its own.
//
// When the crop node’s contents are rendered, the crop node first draws its
// mask into a private buffer. Then, it renders its children. When rendering
// its children, each pixel is verified against the corresponding pixel in the
// mask. If the pixel in the mask has an alpha value of less than 0.05, the
// image pixel is masked out. Any pixel not rendered by the mask node is
// automatically masked out.
//
// The default value of this property is `nil`, which indicates that the child
// nodes should not be cropped.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKCropNode/maskNode
func (c SKCropNode) MaskNode() ISKNode {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("maskNode"))
	return SKNodeFromID(objc.ID(rv))
}
func (c SKCropNode) SetMaskNode(value ISKNode) {
	objc.Send[struct{}](c.ID, objc.Sel("setMaskNode:"), value)
}
