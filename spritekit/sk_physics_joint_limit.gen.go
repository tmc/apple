// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [SKPhysicsJointLimit] class.
var (
	_SKPhysicsJointLimitClass     SKPhysicsJointLimitClass
	_SKPhysicsJointLimitClassOnce sync.Once
)

func getSKPhysicsJointLimitClass() SKPhysicsJointLimitClass {
	_SKPhysicsJointLimitClassOnce.Do(func() {
		_SKPhysicsJointLimitClass = SKPhysicsJointLimitClass{class: objc.GetClass("SKPhysicsJointLimit")}
	})
	return _SKPhysicsJointLimitClass
}

// GetSKPhysicsJointLimitClass returns the class object for SKPhysicsJointLimit.
func GetSKPhysicsJointLimitClass() SKPhysicsJointLimitClass {
	return getSKPhysicsJointLimitClass()
}

type SKPhysicsJointLimitClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKPhysicsJointLimitClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKPhysicsJointLimitClass) Alloc() SKPhysicsJointLimit {
	rv := objc.Send[SKPhysicsJointLimit](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A joint that imposes a maximum distance between two physics bodies, as if
// they were connected by a rope.
//
// # Configuring a Limit Joint
//
//   - [SKPhysicsJointLimit.MaxLength]: The maximum distance allowed between the two physics bodies connected by the limit joint.
//   - [SKPhysicsJointLimit.SetMaxLength]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointLimit
type SKPhysicsJointLimit struct {
	SKPhysicsJoint
}

// SKPhysicsJointLimitFromID constructs a [SKPhysicsJointLimit] from an objc.ID.
//
// A joint that imposes a maximum distance between two physics bodies, as if
// they were connected by a rope.
func SKPhysicsJointLimitFromID(id objc.ID) SKPhysicsJointLimit {
	return SKPhysicsJointLimit{SKPhysicsJoint: SKPhysicsJointFromID(id)}
}

// NOTE: SKPhysicsJointLimit adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKPhysicsJointLimit] class.
//
// # Configuring a Limit Joint
//
//   - [ISKPhysicsJointLimit.MaxLength]: The maximum distance allowed between the two physics bodies connected by the limit joint.
//   - [ISKPhysicsJointLimit.SetMaxLength]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointLimit
type ISKPhysicsJointLimit interface {
	ISKPhysicsJoint

	// Topic: Configuring a Limit Joint

	// The maximum distance allowed between the two physics bodies connected by the limit joint.
	MaxLength() float64
	SetMaxLength(value float64)
}

// Init initializes the instance.
func (p SKPhysicsJointLimit) Init() SKPhysicsJointLimit {
	rv := objc.Send[SKPhysicsJointLimit](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p SKPhysicsJointLimit) Autorelease() SKPhysicsJointLimit {
	rv := objc.Send[SKPhysicsJointLimit](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKPhysicsJointLimit creates a new SKPhysicsJointLimit instance.
func NewSKPhysicsJointLimit() SKPhysicsJointLimit {
	class := getSKPhysicsJointLimitClass()
	rv := objc.Send[SKPhysicsJointLimit](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new limit joint.
//
// bodyA: The first body to connect. The body must be connected to a node that is
// already part of the scene’s node tree.
//
// bodyB: The second body to connect. The body must be connected to a node that is
// already part of the scene’s node tree.
//
// anchorA: A connection point on the first body in the scene’s coordinate system.
//
// anchorB: A connection point on the second body in the scene’s coordinate system.
//
// # Return Value
//
// A new limit joint.
//
// # Discussion
//
// You must add the joint to a physics world associated with the scene before
// it takes effect.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointLimit/joint(withBodyA:bodyB:anchorA:anchorB:)
func (_SKPhysicsJointLimitClass SKPhysicsJointLimitClass) JointWithBodyABodyBAnchorAAnchorB(bodyA ISKPhysicsBody, bodyB ISKPhysicsBody, anchorA corefoundation.CGPoint, anchorB corefoundation.CGPoint) SKPhysicsJointLimit {
	rv := objc.Send[objc.ID](objc.ID(_SKPhysicsJointLimitClass.class), objc.Sel("jointWithBodyA:bodyB:anchorA:anchorB:"), bodyA, bodyB, anchorA, anchorB)
	return SKPhysicsJointLimitFromID(rv)
}

// The maximum distance allowed between the two physics bodies connected by
// the limit joint.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointLimit/maxLength
func (p SKPhysicsJointLimit) MaxLength() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("maxLength"))
	return rv
}
func (p SKPhysicsJointLimit) SetMaxLength(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setMaxLength:"), value)
}
