// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [SKPhysicsJointSliding] class.
var (
	_SKPhysicsJointSlidingClass     SKPhysicsJointSlidingClass
	_SKPhysicsJointSlidingClassOnce sync.Once
)

func getSKPhysicsJointSlidingClass() SKPhysicsJointSlidingClass {
	_SKPhysicsJointSlidingClassOnce.Do(func() {
		_SKPhysicsJointSlidingClass = SKPhysicsJointSlidingClass{class: objc.GetClass("SKPhysicsJointSliding")}
	})
	return _SKPhysicsJointSlidingClass
}

// GetSKPhysicsJointSlidingClass returns the class object for SKPhysicsJointSliding.
func GetSKPhysicsJointSlidingClass() SKPhysicsJointSlidingClass {
	return getSKPhysicsJointSlidingClass()
}

type SKPhysicsJointSlidingClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKPhysicsJointSlidingClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKPhysicsJointSlidingClass) Alloc() SKPhysicsJointSliding {
	rv := objc.Send[SKPhysicsJointSliding](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A joint that allows two physics bodies to slide along an axis.
//
// # Overview
//
// An [SKPhysicsJointSliding] object allows the anchor points of the two
// physics bodies to slide along a chosen axis. The joint can be configured to
// limit the distance that the two objects are allowed to slide along the
// axis.
//
// # Configuring a Sliding Joint
//
//   - [SKPhysicsJointSliding.ShouldEnableLimits]: A Boolean value that indicates whether the sliding joint is restricted so that the objects may only slide a finite distance from the initial anchor point.
//   - [SKPhysicsJointSliding.SetShouldEnableLimits]
//   - [SKPhysicsJointSliding.LowerDistanceLimit]: The smallest distance allowed for the sliding joint.
//   - [SKPhysicsJointSliding.SetLowerDistanceLimit]
//   - [SKPhysicsJointSliding.UpperDistanceLimit]: The largest distance allowed for the sliding joint.
//   - [SKPhysicsJointSliding.SetUpperDistanceLimit]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointSliding
type SKPhysicsJointSliding struct {
	SKPhysicsJoint
}

// SKPhysicsJointSlidingFromID constructs a [SKPhysicsJointSliding] from an objc.ID.
//
// A joint that allows two physics bodies to slide along an axis.
func SKPhysicsJointSlidingFromID(id objc.ID) SKPhysicsJointSliding {
	return SKPhysicsJointSliding{SKPhysicsJoint: SKPhysicsJointFromID(id)}
}

// NOTE: SKPhysicsJointSliding adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKPhysicsJointSliding] class.
//
// # Configuring a Sliding Joint
//
//   - [ISKPhysicsJointSliding.ShouldEnableLimits]: A Boolean value that indicates whether the sliding joint is restricted so that the objects may only slide a finite distance from the initial anchor point.
//   - [ISKPhysicsJointSliding.SetShouldEnableLimits]
//   - [ISKPhysicsJointSliding.LowerDistanceLimit]: The smallest distance allowed for the sliding joint.
//   - [ISKPhysicsJointSliding.SetLowerDistanceLimit]
//   - [ISKPhysicsJointSliding.UpperDistanceLimit]: The largest distance allowed for the sliding joint.
//   - [ISKPhysicsJointSliding.SetUpperDistanceLimit]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointSliding
type ISKPhysicsJointSliding interface {
	ISKPhysicsJoint

	// Topic: Configuring a Sliding Joint

	// A Boolean value that indicates whether the sliding joint is restricted so that the objects may only slide a finite distance from the initial anchor point.
	ShouldEnableLimits() bool
	SetShouldEnableLimits(value bool)
	// The smallest distance allowed for the sliding joint.
	LowerDistanceLimit() float64
	SetLowerDistanceLimit(value float64)
	// The largest distance allowed for the sliding joint.
	UpperDistanceLimit() float64
	SetUpperDistanceLimit(value float64)
}

// Init initializes the instance.
func (p SKPhysicsJointSliding) Init() SKPhysicsJointSliding {
	rv := objc.Send[SKPhysicsJointSliding](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p SKPhysicsJointSliding) Autorelease() SKPhysicsJointSliding {
	rv := objc.Send[SKPhysicsJointSliding](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKPhysicsJointSliding creates a new SKPhysicsJointSliding instance.
func NewSKPhysicsJointSliding() SKPhysicsJointSliding {
	class := getSKPhysicsJointSlidingClass()
	rv := objc.Send[SKPhysicsJointSliding](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new sliding joint.
//
// bodyA: The first body to connect. The body must be connected to a node that is
// already part of the scene’s node tree.
//
// bodyB: The second body to connect. The body must be connected to a node that is
// already part of the scene’s node tree.
//
// anchor: The anchor point for the connection in the scene’s coordinate system.
//
// axis: A vector that defines the direction that the joint is allowed to slide.
//
// # Return Value
//
// A new sliding joint.
//
// # Discussion
//
// You must add the joint to a physics world associated with the scene before
// it takes effect.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointSliding/joint(withBodyA:bodyB:anchor:axis:)
func (_SKPhysicsJointSlidingClass SKPhysicsJointSlidingClass) JointWithBodyABodyBAnchorAxis(bodyA ISKPhysicsBody, bodyB ISKPhysicsBody, anchor corefoundation.CGPoint, axis corefoundation.CGVector) SKPhysicsJointSliding {
	rv := objc.Send[objc.ID](objc.ID(_SKPhysicsJointSlidingClass.class), objc.Sel("jointWithBodyA:bodyB:anchor:axis:"), bodyA, bodyB, anchor, axis)
	return SKPhysicsJointSlidingFromID(rv)
}

// A Boolean value that indicates whether the sliding joint is restricted so
// that the objects may only slide a finite distance from the initial anchor
// point.
//
// # Discussion
//
// The default value is false. If true, then the
// [SKPhysicsJointSliding.LowerDistanceLimit] and
// [SKPhysicsJointSliding.UpperDistanceLimit] properties are used to limit the
// distance of the sliding joint.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointSliding/shouldEnableLimits
func (p SKPhysicsJointSliding) ShouldEnableLimits() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("shouldEnableLimits"))
	return rv
}
func (p SKPhysicsJointSliding) SetShouldEnableLimits(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setShouldEnableLimits:"), value)
}

// The smallest distance allowed for the sliding joint.
//
// # Discussion
//
// The default value is `0.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointSliding/lowerDistanceLimit
func (p SKPhysicsJointSliding) LowerDistanceLimit() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("lowerDistanceLimit"))
	return rv
}
func (p SKPhysicsJointSliding) SetLowerDistanceLimit(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setLowerDistanceLimit:"), value)
}

// The largest distance allowed for the sliding joint.
//
// # Discussion
//
// The default value is `0.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointSliding/upperDistanceLimit
func (p SKPhysicsJointSliding) UpperDistanceLimit() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("upperDistanceLimit"))
	return rv
}
func (p SKPhysicsJointSliding) SetUpperDistanceLimit(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setUpperDistanceLimit:"), value)
}
