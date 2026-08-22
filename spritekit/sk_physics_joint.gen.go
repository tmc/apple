// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKPhysicsJoint] class.
var (
	_SKPhysicsJointClass     SKPhysicsJointClass
	_SKPhysicsJointClassOnce sync.Once
)

func getSKPhysicsJointClass() SKPhysicsJointClass {
	_SKPhysicsJointClassOnce.Do(func() {
		_SKPhysicsJointClass = SKPhysicsJointClass{class: objc.GetClass("SKPhysicsJoint")}
	})
	return _SKPhysicsJointClass
}

// GetSKPhysicsJointClass returns the class object for SKPhysicsJoint.
func GetSKPhysicsJointClass() SKPhysicsJointClass {
	return getSKPhysicsJointClass()
}

type SKPhysicsJointClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKPhysicsJointClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKPhysicsJointClass) Alloc() SKPhysicsJoint {
	rv := objc.Send[SKPhysicsJoint](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// The abstract superclass for objects that connect physics bodies.
//
// # Overview
//
// An [SKPhysicsJoint] object connects two physics bodies so that they are
// simulated together by the physics world. You never instantiate objects of
// this class directly; instead, you instantiate one of the subclasses that
// defines the kind of joint you want to make.
//
// # Accessing or Setting a Joint’s Bodies
//
//   - [SKPhysicsJoint.BodyA]: The first body connected by the joint.
//   - [SKPhysicsJoint.SetBodyA]
//   - [SKPhysicsJoint.BodyB]: The second body connected by the joint.
//   - [SKPhysicsJoint.SetBodyB]
//
// # Reading the Stress and Speed that Are Currently Applied to a Joint
//
//   - [SKPhysicsJoint.ReactionForce]: The instantaneous reaction force, in newtons, currently being directed at the anchor point.
//   - [SKPhysicsJoint.ReactionTorque]: Instantaneous reaction torque, in newton-meters,  currently being directed at the anchor point.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJoint
type SKPhysicsJoint struct {
	objectivec.Object
}

// SKPhysicsJointFromID constructs a [SKPhysicsJoint] from an objc.ID.
//
// The abstract superclass for objects that connect physics bodies.
func SKPhysicsJointFromID(id objc.ID) SKPhysicsJoint {
	return SKPhysicsJoint{objectivec.Object{ID: id}}
}

// NOTE: SKPhysicsJoint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKPhysicsJoint] class.
//
// # Accessing or Setting a Joint’s Bodies
//
//   - [ISKPhysicsJoint.BodyA]: The first body connected by the joint.
//   - [ISKPhysicsJoint.SetBodyA]
//   - [ISKPhysicsJoint.BodyB]: The second body connected by the joint.
//   - [ISKPhysicsJoint.SetBodyB]
//
// # Reading the Stress and Speed that Are Currently Applied to a Joint
//
//   - [ISKPhysicsJoint.ReactionForce]: The instantaneous reaction force, in newtons, currently being directed at the anchor point.
//   - [ISKPhysicsJoint.ReactionTorque]: Instantaneous reaction torque, in newton-meters,  currently being directed at the anchor point.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJoint
type ISKPhysicsJoint interface {
	objectivec.IObject

	// Topic: Accessing or Setting a Joint’s Bodies

	// The first body connected by the joint.
	BodyA() ISKPhysicsBody
	SetBodyA(value ISKPhysicsBody)
	// The second body connected by the joint.
	BodyB() ISKPhysicsBody
	SetBodyB(value ISKPhysicsBody)

	// Topic: Reading the Stress and Speed that Are Currently Applied to a Joint

	// The instantaneous reaction force, in newtons, currently being directed at the anchor point.
	ReactionForce() corefoundation.CGVector
	// Instantaneous reaction torque, in newton-meters,  currently being directed at the anchor point.
	ReactionTorque() float64

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (p SKPhysicsJoint) Init() SKPhysicsJoint {
	rv := objc.Send[SKPhysicsJoint](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p SKPhysicsJoint) Autorelease() SKPhysicsJoint {
	rv := objc.Send[SKPhysicsJoint](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKPhysicsJoint creates a new SKPhysicsJoint instance.
func NewSKPhysicsJoint() SKPhysicsJoint {
	class := getSKPhysicsJointClass()
	rv := objc.Send[SKPhysicsJoint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (p SKPhysicsJoint) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The first body connected by the joint.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJoint/bodyA
func (p SKPhysicsJoint) BodyA() ISKPhysicsBody {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("bodyA"))
	return SKPhysicsBodyFromID(objc.ID(rv))
}
func (p SKPhysicsJoint) SetBodyA(value ISKPhysicsBody) {
	objc.Send[struct{}](p.ID, objc.Sel("setBodyA:"), value)
}

// The second body connected by the joint.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJoint/bodyB
func (p SKPhysicsJoint) BodyB() ISKPhysicsBody {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("bodyB"))
	return SKPhysicsBodyFromID(objc.ID(rv))
}
func (p SKPhysicsJoint) SetBodyB(value ISKPhysicsBody) {
	objc.Send[struct{}](p.ID, objc.Sel("setBodyB:"), value)
}

// The instantaneous reaction force, in newtons, currently being directed at
// the anchor point.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJoint/reactionForce
func (p SKPhysicsJoint) ReactionForce() corefoundation.CGVector {
	rv := objc.Send[corefoundation.CGVector](p.ID, objc.Sel("reactionForce"))
	return corefoundation.CGVector(rv)
}

// Instantaneous reaction torque, in newton-meters, currently being directed
// at the anchor point.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJoint/reactionTorque
func (p SKPhysicsJoint) ReactionTorque() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("reactionTorque"))
	return rv
}
