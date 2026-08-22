// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [SKPhysicsJointSpring] class.
var (
	_SKPhysicsJointSpringClass     SKPhysicsJointSpringClass
	_SKPhysicsJointSpringClassOnce sync.Once
)

func getSKPhysicsJointSpringClass() SKPhysicsJointSpringClass {
	_SKPhysicsJointSpringClassOnce.Do(func() {
		_SKPhysicsJointSpringClass = SKPhysicsJointSpringClass{class: objc.GetClass("SKPhysicsJointSpring")}
	})
	return _SKPhysicsJointSpringClass
}

// GetSKPhysicsJointSpringClass returns the class object for SKPhysicsJointSpring.
func GetSKPhysicsJointSpringClass() SKPhysicsJointSpringClass {
	return getSKPhysicsJointSpringClass()
}

type SKPhysicsJointSpringClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKPhysicsJointSpringClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKPhysicsJointSpringClass) Alloc() SKPhysicsJointSpring {
	rv := objc.Send[SKPhysicsJointSpring](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A joint that simulates a spring connecting two physics bodies.
//
// # Overview
//
// An [SKPhysicsJointSpring] object simulates connecting two physics bodies
// together with a spring. The farther the two objects move from each other,
// the more force is applied to bring the two bodies back together.
//
// # Configuring a Spring Joint
//
//   - [SKPhysicsJointSpring.Damping]: Defines how the spring’s motion should be damped due to the forces of friction.
//   - [SKPhysicsJointSpring.SetDamping]
//   - [SKPhysicsJointSpring.Frequency]: Defines the frequency or stiffness characteristics of the spring.
//   - [SKPhysicsJointSpring.SetFrequency]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointSpring
type SKPhysicsJointSpring struct {
	SKPhysicsJoint
}

// SKPhysicsJointSpringFromID constructs a [SKPhysicsJointSpring] from an objc.ID.
//
// A joint that simulates a spring connecting two physics bodies.
func SKPhysicsJointSpringFromID(id objc.ID) SKPhysicsJointSpring {
	return SKPhysicsJointSpring{SKPhysicsJoint: SKPhysicsJointFromID(id)}
}

// NOTE: SKPhysicsJointSpring adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKPhysicsJointSpring] class.
//
// # Configuring a Spring Joint
//
//   - [ISKPhysicsJointSpring.Damping]: Defines how the spring’s motion should be damped due to the forces of friction.
//   - [ISKPhysicsJointSpring.SetDamping]
//   - [ISKPhysicsJointSpring.Frequency]: Defines the frequency or stiffness characteristics of the spring.
//   - [ISKPhysicsJointSpring.SetFrequency]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointSpring
type ISKPhysicsJointSpring interface {
	ISKPhysicsJoint

	// Topic: Configuring a Spring Joint

	// Defines how the spring’s motion should be damped due to the forces of friction.
	Damping() float64
	SetDamping(value float64)
	// Defines the frequency or stiffness characteristics of the spring.
	Frequency() float64
	SetFrequency(value float64)
}

// Init initializes the instance.
func (p SKPhysicsJointSpring) Init() SKPhysicsJointSpring {
	rv := objc.Send[SKPhysicsJointSpring](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p SKPhysicsJointSpring) Autorelease() SKPhysicsJointSpring {
	rv := objc.Send[SKPhysicsJointSpring](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKPhysicsJointSpring creates a new SKPhysicsJointSpring instance.
func NewSKPhysicsJointSpring() SKPhysicsJointSpring {
	class := getSKPhysicsJointSpringClass()
	rv := objc.Send[SKPhysicsJointSpring](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new spring joint.
//
// bodyA: The first body to connect. The body must be connected to a node that is
// already part of the scene’s node tree.
//
// bodyB: The second body to connect. The body must be connected to a node that is
// already part of the scene’s node tree.
//
// anchorA: The connection point on the first body in the scene’s coordinate system.
//
// anchorB: The connection point on the second body in the scene’s coordinate system.
//
// # Return Value
//
// A new spring joint.
//
// # Discussion
//
// You must add the joint to a physics world associated with the scene before
// it takes effect.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointSpring/joint(withBodyA:bodyB:anchorA:anchorB:)
func (_SKPhysicsJointSpringClass SKPhysicsJointSpringClass) JointWithBodyABodyBAnchorAAnchorB(bodyA ISKPhysicsBody, bodyB ISKPhysicsBody, anchorA corefoundation.CGPoint, anchorB corefoundation.CGPoint) SKPhysicsJointSpring {
	rv := objc.Send[objc.ID](objc.ID(_SKPhysicsJointSpringClass.class), objc.Sel("jointWithBodyA:bodyB:anchorA:anchorB:"), bodyA, bodyB, anchorA, anchorB)
	return SKPhysicsJointSpringFromID(rv)
}

// Defines how the spring’s motion should be damped due to the forces of
// friction.
//
// # Discussion
//
// The default value is `0.0`. Increasing the value increases the energy loss
// with each oscillation: there will be fewer and smaller oscillations and
// time taken for the spring to settle will decrease.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointSpring/damping
func (p SKPhysicsJointSpring) Damping() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("damping"))
	return rv
}
func (p SKPhysicsJointSpring) SetDamping(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setDamping:"), value)
}

// Defines the frequency or stiffness characteristics of the spring.
//
// # Discussion
//
// The default value is `0.0`, creating a rigid joint between the spring’s
// two bodies. Setting the frequency to a low value, for example `0.5`,
// creates a spring with slow oscillations that will settle slowly. Setting
// the frequency to a high value, for example `10.0`, creates a stiffer spring
// with faster and fewer oscillations.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointSpring/frequency
func (p SKPhysicsJointSpring) Frequency() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("frequency"))
	return rv
}
func (p SKPhysicsJointSpring) SetFrequency(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setFrequency:"), value)
}
