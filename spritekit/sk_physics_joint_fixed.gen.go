// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [SKPhysicsJointFixed] class.
var (
	_SKPhysicsJointFixedClass     SKPhysicsJointFixedClass
	_SKPhysicsJointFixedClassOnce sync.Once
)

func getSKPhysicsJointFixedClass() SKPhysicsJointFixedClass {
	_SKPhysicsJointFixedClassOnce.Do(func() {
		_SKPhysicsJointFixedClass = SKPhysicsJointFixedClass{class: objc.GetClass("SKPhysicsJointFixed")}
	})
	return _SKPhysicsJointFixedClass
}

// GetSKPhysicsJointFixedClass returns the class object for SKPhysicsJointFixed.
func GetSKPhysicsJointFixedClass() SKPhysicsJointFixedClass {
	return getSKPhysicsJointFixedClass()
}

type SKPhysicsJointFixedClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKPhysicsJointFixedClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKPhysicsJointFixedClass) Alloc() SKPhysicsJointFixed {
	rv := objc.Send[SKPhysicsJointFixed](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A joint that fuses two physics bodies together at a reference point.
//
// # Overview
//
// An [SKPhysicsJointFixed] object fuses two physics bodies together at a
// reference point. Fixed joints are useful for creating complex shapes that
// can be broken apart later.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointFixed
type SKPhysicsJointFixed struct {
	SKPhysicsJoint
}

// SKPhysicsJointFixedFromID constructs a [SKPhysicsJointFixed] from an objc.ID.
//
// A joint that fuses two physics bodies together at a reference point.
func SKPhysicsJointFixedFromID(id objc.ID) SKPhysicsJointFixed {
	return SKPhysicsJointFixed{SKPhysicsJoint: SKPhysicsJointFromID(id)}
}

// NOTE: SKPhysicsJointFixed adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKPhysicsJointFixed] class.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointFixed
type ISKPhysicsJointFixed interface {
	ISKPhysicsJoint
}

// Init initializes the instance.
func (p SKPhysicsJointFixed) Init() SKPhysicsJointFixed {
	rv := objc.Send[SKPhysicsJointFixed](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p SKPhysicsJointFixed) Autorelease() SKPhysicsJointFixed {
	rv := objc.Send[SKPhysicsJointFixed](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKPhysicsJointFixed creates a new SKPhysicsJointFixed instance.
func NewSKPhysicsJointFixed() SKPhysicsJointFixed {
	class := getSKPhysicsJointFixedClass()
	rv := objc.Send[SKPhysicsJointFixed](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new fixed joint.
//
// bodyA: The first body to connect. The body must be connected to a node that is
// already part of the scene’s node tree.
//
// bodyB: The second body to connect. The body must be connected to a node that is
// already part of the scene’s node tree.
//
// anchor: The anchor point for the connection in the scene’s coordinate system.
//
// # Return Value
//
// A new fixed joint.
//
// # Discussion
//
// You must add the joint to a physics world associated with the scene before
// it takes effect.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointFixed/joint(withBodyA:bodyB:anchor:)
func (_SKPhysicsJointFixedClass SKPhysicsJointFixedClass) JointWithBodyABodyBAnchor(bodyA ISKPhysicsBody, bodyB ISKPhysicsBody, anchor corefoundation.CGPoint) SKPhysicsJointFixed {
	rv := objc.Send[objc.ID](objc.ID(_SKPhysicsJointFixedClass.class), objc.Sel("jointWithBodyA:bodyB:anchor:"), bodyA, bodyB, anchor)
	return SKPhysicsJointFixedFromID(rv)
}
