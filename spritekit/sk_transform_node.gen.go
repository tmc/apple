// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [SKTransformNode] class.
var (
	_SKTransformNodeClass     SKTransformNodeClass
	_SKTransformNodeClassOnce sync.Once
)

func getSKTransformNodeClass() SKTransformNodeClass {
	_SKTransformNodeClassOnce.Do(func() {
		_SKTransformNodeClass = SKTransformNodeClass{class: objc.GetClass("SKTransformNode")}
	})
	return _SKTransformNodeClass
}

// GetSKTransformNodeClass returns the class object for SKTransformNode.
func GetSKTransformNodeClass() SKTransformNodeClass {
	return getSKTransformNodeClass()
}

type SKTransformNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKTransformNodeClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKTransformNodeClass) Alloc() SKTransformNode {
	rv := objc.Send[SKTransformNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A node that allows its children to rotate in 3D.
//
// # Overview
//
// [SKTranformNode] adds the ability to rotate nodes across the x and y axes.
// When combined with [SKNode]’s [SKNode.ZRotation] property, nodes added as
// children to a transform node have the ability to rotate in 3D.
//
// # Rotating Child Nodes
//
//   - [SKTransformNode.XRotation]
//   - [SKTransformNode.SetXRotation]
//   - [SKTransformNode.YRotation]
//   - [SKTransformNode.SetYRotation]
//   - [SKTransformNode.SetEulerAngles]
//   - [SKTransformNode.SetRotationMatrix]
//
// # Reading the Current Rotation
//
//   - [SKTransformNode.EulerAngles]
//   - [SKTransformNode.RotationMatrix]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTransformNode
type SKTransformNode struct {
	SKNode
}

// SKTransformNodeFromID constructs a [SKTransformNode] from an objc.ID.
//
// A node that allows its children to rotate in 3D.
func SKTransformNodeFromID(id objc.ID) SKTransformNode {
	return SKTransformNode{SKNode: SKNodeFromID(id)}
}

// NOTE: SKTransformNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKTransformNode] class.
//
// # Rotating Child Nodes
//
//   - [ISKTransformNode.XRotation]
//   - [ISKTransformNode.SetXRotation]
//   - [ISKTransformNode.YRotation]
//   - [ISKTransformNode.SetYRotation]
//   - [ISKTransformNode.SetEulerAngles]
//   - [ISKTransformNode.SetRotationMatrix]
//
// # Reading the Current Rotation
//
//   - [ISKTransformNode.EulerAngles]
//   - [ISKTransformNode.RotationMatrix]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTransformNode
type ISKTransformNode interface {
	ISKNode

	// Topic: Rotating Child Nodes

	XRotation() float64
	SetXRotation(value float64)
	YRotation() float64
	SetYRotation(value float64)
	SetEulerAngles(euler Vector_float3)
	SetRotationMatrix(rotationMatrix [3][4]float32)

	// Topic: Reading the Current Rotation

	EulerAngles() Vector_float3
	RotationMatrix() [3][4]float32
}

// Init initializes the instance.
func (t SKTransformNode) Init() SKTransformNode {
	rv := objc.Send[SKTransformNode](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t SKTransformNode) Autorelease() SKTransformNode {
	rv := objc.Send[SKTransformNode](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKTransformNode creates a new SKTransformNode instance.
func NewSKTransformNode() SKTransformNode {
	class := getSKTransformNodeClass()
	rv := objc.Send[SKTransformNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Called when a node is initialized from an .sks file.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(coder:)
func NewTransformNodeWithCoder(aDecoder foundation.INSCoder) SKTransformNode {
	instance := getSKTransformNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return SKTransformNodeFromID(rv)
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
func NewTransformNodeWithFileNamed(filename string) SKTransformNode {
	rv := objc.Send[objc.ID](objc.ID(getSKTransformNodeClass().class), objc.Sel("nodeWithFileNamed:"), objc.String(filename))
	return SKTransformNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(fileNamed:securelyWithClasses:)
func NewTransformNodeWithFileNamedSecurelyWithClassesAndError(filename string, classes foundation.INSSet) (SKTransformNode, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getSKTransformNodeClass().class), objc.Sel("nodeWithFileNamed:securelyWithClasses:andError:"), objc.String(filename), classes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SKTransformNode{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SKTransformNode{}, objc.ErrInitFailed
	}
	return SKTransformNodeFromID(rv), nil
}

// See: https://developer.apple.com/documentation/SpriteKit/SKTransformNode/setEulerAngles(_:)
func (t SKTransformNode) SetEulerAngles(euler Vector_float3) {
	objc.Send[objc.ID](t.ID, objc.Sel("setEulerAngles:"), euler)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKTransformNode/setRotationMatrix(_:)
func (t SKTransformNode) SetRotationMatrix(rotationMatrix [3][4]float32) {
	objc.Send[objc.ID](t.ID, objc.Sel("setRotationMatrix:"), rotationMatrix)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKTransformNode/eulerAngles()
func (t SKTransformNode) EulerAngles() Vector_float3 {
	rv := objc.Send[Vector_float3](t.ID, objc.Sel("eulerAngles"))
	return Vector_float3(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKTransformNode/rotationMatrix()
func (t SKTransformNode) RotationMatrix() [3][4]float32 {
	rv := objc.Send[[3][4]float32](t.ID, objc.Sel("rotationMatrix"))
	return [3][4]float32(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKTransformNode/xRotation
func (t SKTransformNode) XRotation() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("xRotation"))
	return rv
}
func (t SKTransformNode) SetXRotation(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setXRotation:"), value)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKTransformNode/yRotation
func (t SKTransformNode) YRotation() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("yRotation"))
	return rv
}
func (t SKTransformNode) SetYRotation(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setYRotation:"), value)
}
