// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [SKReferenceNode] class.
var (
	_SKReferenceNodeClass     SKReferenceNodeClass
	_SKReferenceNodeClassOnce sync.Once
)

func getSKReferenceNodeClass() SKReferenceNodeClass {
	_SKReferenceNodeClassOnce.Do(func() {
		_SKReferenceNodeClass = SKReferenceNodeClass{class: objc.GetClass("SKReferenceNode")}
	})
	return _SKReferenceNodeClass
}

// GetSKReferenceNodeClass returns the class object for SKReferenceNode.
func GetSKReferenceNodeClass() SKReferenceNodeClass {
	return getSKReferenceNodeClass()
}

type SKReferenceNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKReferenceNodeClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKReferenceNodeClass) Alloc() SKReferenceNode {
	rv := objc.Send[SKReferenceNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A node that’s defined in an archived `XCUIElementTypeSks` file.
//
// # Overview
//
// [SKReferenceNode] is used within an archived `XCUIElementTypeSks` file to
// refer to node defined in another `XCUIElementTypeSks` file without
// duplicating its definition. This way, a change to the referenced node
// propagates to all the references in other files.
//
// As an example, you might want to share an enemy ship across two different
// levels, Scene1.sks and Scene2.sks, in a level-based game. Reference nodes
// allow you to do that without creating copies of the shared node and its
// properties.
//
// To use a reference node:
//
// - Create the shared content in a separate archive - Add references to the
// shared archive within your scene archives
//
// When each scene is loaded, the reference nodes are resolved dynamically,
// and therefore you only need to configure a shared object in one place.
//
// # Initializers
//
//   - [SKReferenceNode.InitWithURL]: Initializes a reference node from a URL.
//   - [SKReferenceNode.InitWithFileNamed]: Initializes a reference node from a file in the app’s main bundle.
//
// # Regenerating
//
//   - [SKReferenceNode.ResolveReferenceNode]: Loads the reference node’s content and adds it as a new child node.
//
// # Loading Callback
//
//   - [SKReferenceNode.DidLoadReferenceNode]: A method called by SpriteKit after the reference node’s contents are loaded.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKReferenceNode
type SKReferenceNode struct {
	SKNode
}

// SKReferenceNodeFromID constructs a [SKReferenceNode] from an objc.ID.
//
// A node that’s defined in an archived `XCUIElementTypeSks` file.
func SKReferenceNodeFromID(id objc.ID) SKReferenceNode {
	return SKReferenceNode{SKNode: SKNodeFromID(id)}
}

// NOTE: SKReferenceNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKReferenceNode] class.
//
// # Initializers
//
//   - [ISKReferenceNode.InitWithURL]: Initializes a reference node from a URL.
//   - [ISKReferenceNode.InitWithFileNamed]: Initializes a reference node from a file in the app’s main bundle.
//
// # Regenerating
//
//   - [ISKReferenceNode.ResolveReferenceNode]: Loads the reference node’s content and adds it as a new child node.
//
// # Loading Callback
//
//   - [ISKReferenceNode.DidLoadReferenceNode]: A method called by SpriteKit after the reference node’s contents are loaded.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKReferenceNode
type ISKReferenceNode interface {
	ISKNode

	// Topic: Initializers

	// Initializes a reference node from a URL.
	InitWithURL(url foundation.NSURL) SKReferenceNode
	// Initializes a reference node from a file in the app’s main bundle.
	InitWithFileNamed(fileName string) SKReferenceNode

	// Topic: Regenerating

	// Loads the reference node’s content and adds it as a new child node.
	ResolveReferenceNode()

	// Topic: Loading Callback

	// A method called by SpriteKit after the reference node’s contents are loaded.
	DidLoadReferenceNode(node ISKNode)
}

// Init initializes the instance.
func (r SKReferenceNode) Init() SKReferenceNode {
	rv := objc.Send[SKReferenceNode](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r SKReferenceNode) Autorelease() SKReferenceNode {
	rv := objc.Send[SKReferenceNode](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKReferenceNode creates a new SKReferenceNode instance.
func NewSKReferenceNode() SKReferenceNode {
	class := getSKReferenceNodeClass()
	rv := objc.Send[SKReferenceNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A method that initializes a reference node from an archive.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKReferenceNode/init(coder:)
func NewReferenceNodeWithCoder(aDecoder foundation.INSCoder) SKReferenceNode {
	instance := getSKReferenceNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return SKReferenceNodeFromID(rv)
}

// Initializes a reference node from a file in the app’s main bundle.
//
// fileName: The name of a file stored in the app’s main bundle.
//
// # Return Value
//
// A newly initialized reference node.
//
// # Discussion
//
// This initializer is for loading files that reside inside of the app bundle.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKReferenceNode/init(fileNamed:)-2yeh2
func NewReferenceNodeWithFileNamed(fileName string) SKReferenceNode {
	instance := getSKReferenceNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFileNamed:"), objc.String(fileName))
	return SKReferenceNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(fileNamed:securelyWithClasses:)
func NewReferenceNodeWithFileNamedSecurelyWithClassesAndError(filename string, classes foundation.INSSet) (SKReferenceNode, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getSKReferenceNodeClass().class), objc.Sel("nodeWithFileNamed:securelyWithClasses:andError:"), objc.String(filename), classes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SKReferenceNode{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SKReferenceNode{}, objc.ErrInitFailed
	}
	return SKReferenceNodeFromID(rv), nil
}

// Initializes a reference node from a URL.
//
// url: The URL of the reference node.
//
// # Return Value
//
// A newly initialized reference node.
//
// # Discussion
//
// This intializer is for loading archives that reside outside of the app
// bundle.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKReferenceNode/init(url:)-429mo
func NewReferenceNodeWithURL(url foundation.NSURL) SKReferenceNode {
	instance := getSKReferenceNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), url)
	return SKReferenceNodeFromID(rv)
}

// Initializes a reference node from a URL.
//
// url: The URL of the reference node.
//
// # Return Value
//
// A newly initialized reference node.
//
// # Discussion
//
// This intializer is for loading archives that reside outside of the app
// bundle.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKReferenceNode/init(url:)-429mo
func (r SKReferenceNode) InitWithURL(url foundation.NSURL) SKReferenceNode {
	rv := objc.Send[SKReferenceNode](r.ID, objc.Sel("initWithURL:"), url)
	return rv
}

// Initializes a reference node from a file in the app’s main bundle.
//
// fileName: The name of a file stored in the app’s main bundle.
//
// # Return Value
//
// A newly initialized reference node.
//
// # Discussion
//
// This initializer is for loading files that reside inside of the app bundle.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKReferenceNode/init(fileNamed:)-2yeh2
func (r SKReferenceNode) InitWithFileNamed(fileName string) SKReferenceNode {
	rv := objc.Send[SKReferenceNode](r.ID, objc.Sel("initWithFileNamed:"), objc.String(fileName))
	return rv
}

// Loads the reference node’s content and adds it as a new child node.
//
// # Discussion
//
// The archive is deserialized and the root node is added as a child of the
// reference node. If this method is called on a reference node whose content
// is already loaded, the existing node tree is discarded and replaced with a
// fresh copy of the archive’s data.
//
// SpriteKit calls this method automatically if the scene renders the
// reference node and the reference node has not previously been loaded.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKReferenceNode/resolve()
func (r SKReferenceNode) ResolveReferenceNode() {
	objc.Send[objc.ID](r.ID, objc.Sel("resolveReferenceNode"))
}

// A method called by SpriteKit after the reference node’s contents are
// loaded.
//
// node: The deserialized content’s root node.
//
// # Discussion
//
// This method is called after the referenced content is added as a child of
// the reference node. Override this method in a subclass to implement custom
// loading behavior.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKReferenceNode/didLoad(_:)
func (r SKReferenceNode) DidLoadReferenceNode(node ISKNode) {
	objc.Send[objc.ID](r.ID, objc.Sel("didLoadReferenceNode:"), node)
}

// Creates a reference node from a URL.
//
// referenceURL: The URL of the reference node.
//
// # Return Value
//
// A newly initialized reference node.
//
// # Discussion
//
// This intializer is for loading archives that reside outside of the app
// bundle.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKReferenceNode/init(url:)-3jryz
func (_SKReferenceNodeClass SKReferenceNodeClass) ReferenceNodeWithURL(referenceURL foundation.NSURL) SKReferenceNode {
	rv := objc.Send[objc.ID](objc.ID(_SKReferenceNodeClass.class), objc.Sel("referenceNodeWithURL:"), referenceURL)
	return SKReferenceNodeFromID(rv)
}

// Creates a reference node from a file in the app’s main bundle.
//
// fileName: The name of a file stored in the app’s main bundle.
//
// # Return Value
//
// A newly initialized reference node.
//
// # Discussion
//
// This initializer is for loading files that reside inside of the app bundle.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKReferenceNode/init(fileNamed:)-77gs0
func (_SKReferenceNodeClass SKReferenceNodeClass) ReferenceNodeWithFileNamed(fileName string) SKReferenceNode {
	rv := objc.Send[objc.ID](objc.ID(_SKReferenceNodeClass.class), objc.Sel("referenceNodeWithFileNamed:"), objc.String(fileName))
	return SKReferenceNodeFromID(rv)
}
