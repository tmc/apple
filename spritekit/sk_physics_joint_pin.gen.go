// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [SKPhysicsJointPin] class.
var (
	_SKPhysicsJointPinClass     SKPhysicsJointPinClass
	_SKPhysicsJointPinClassOnce sync.Once
)

func getSKPhysicsJointPinClass() SKPhysicsJointPinClass {
	_SKPhysicsJointPinClassOnce.Do(func() {
		_SKPhysicsJointPinClass = SKPhysicsJointPinClass{class: objc.GetClass("SKPhysicsJointPin")}
	})
	return _SKPhysicsJointPinClass
}

// GetSKPhysicsJointPinClass returns the class object for SKPhysicsJointPin.
func GetSKPhysicsJointPinClass() SKPhysicsJointPinClass {
	return getSKPhysicsJointPinClass()
}

type SKPhysicsJointPinClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKPhysicsJointPinClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKPhysicsJointPinClass) Alloc() SKPhysicsJointPin {
	rv := objc.Send[SKPhysicsJointPin](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A joint that pins together two physics bodies, allowing independent
// rotation.
//
// # Overview
//
// An [SKPhysicsJointPin] object allows two physics bodies to independently
// rotate around the anchor point as if pinned together. You can configure how
// far the two objects may rotate and the resistance to rotation.
//
// # Configuring a Pin Joint
//
//   - [SKPhysicsJointPin.RotationSpeed]: The speed, in radians per second, at which the physics bodies are driven around the pin joint.
//   - [SKPhysicsJointPin.SetRotationSpeed]
//   - [SKPhysicsJointPin.ShouldEnableLimits]: A Boolean value that indicates whether the pin joint’s rotation is limited to a specific range of values.
//   - [SKPhysicsJointPin.SetShouldEnableLimits]
//   - [SKPhysicsJointPin.LowerAngleLimit]: The smallest angle allowed for the pin joint, in radians.
//   - [SKPhysicsJointPin.SetLowerAngleLimit]
//   - [SKPhysicsJointPin.UpperAngleLimit]: The largest angle allowed for the pin joint, in radians.
//   - [SKPhysicsJointPin.SetUpperAngleLimit]
//   - [SKPhysicsJointPin.FrictionTorque]: The resistance applied by the pin joint to spinning around the anchor point.
//   - [SKPhysicsJointPin.SetFrictionTorque]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointPin
type SKPhysicsJointPin struct {
	SKPhysicsJoint
}

// SKPhysicsJointPinFromID constructs a [SKPhysicsJointPin] from an objc.ID.
//
// A joint that pins together two physics bodies, allowing independent
// rotation.
func SKPhysicsJointPinFromID(id objc.ID) SKPhysicsJointPin {
	return SKPhysicsJointPin{SKPhysicsJoint: SKPhysicsJointFromID(id)}
}

// NOTE: SKPhysicsJointPin adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKPhysicsJointPin] class.
//
// # Configuring a Pin Joint
//
//   - [ISKPhysicsJointPin.RotationSpeed]: The speed, in radians per second, at which the physics bodies are driven around the pin joint.
//   - [ISKPhysicsJointPin.SetRotationSpeed]
//   - [ISKPhysicsJointPin.ShouldEnableLimits]: A Boolean value that indicates whether the pin joint’s rotation is limited to a specific range of values.
//   - [ISKPhysicsJointPin.SetShouldEnableLimits]
//   - [ISKPhysicsJointPin.LowerAngleLimit]: The smallest angle allowed for the pin joint, in radians.
//   - [ISKPhysicsJointPin.SetLowerAngleLimit]
//   - [ISKPhysicsJointPin.UpperAngleLimit]: The largest angle allowed for the pin joint, in radians.
//   - [ISKPhysicsJointPin.SetUpperAngleLimit]
//   - [ISKPhysicsJointPin.FrictionTorque]: The resistance applied by the pin joint to spinning around the anchor point.
//   - [ISKPhysicsJointPin.SetFrictionTorque]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointPin
type ISKPhysicsJointPin interface {
	ISKPhysicsJoint

	// Topic: Configuring a Pin Joint

	// The speed, in radians per second, at which the physics bodies are driven around the pin joint.
	RotationSpeed() float64
	SetRotationSpeed(value float64)
	// A Boolean value that indicates whether the pin joint’s rotation is limited to a specific range of values.
	ShouldEnableLimits() bool
	SetShouldEnableLimits(value bool)
	// The smallest angle allowed for the pin joint, in radians.
	LowerAngleLimit() float64
	SetLowerAngleLimit(value float64)
	// The largest angle allowed for the pin joint, in radians.
	UpperAngleLimit() float64
	SetUpperAngleLimit(value float64)
	// The resistance applied by the pin joint to spinning around the anchor point.
	FrictionTorque() float64
	SetFrictionTorque(value float64)
}

// Init initializes the instance.
func (p SKPhysicsJointPin) Init() SKPhysicsJointPin {
	rv := objc.Send[SKPhysicsJointPin](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p SKPhysicsJointPin) Autorelease() SKPhysicsJointPin {
	rv := objc.Send[SKPhysicsJointPin](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKPhysicsJointPin creates a new SKPhysicsJointPin instance.
func NewSKPhysicsJointPin() SKPhysicsJointPin {
	class := getSKPhysicsJointPinClass()
	rv := objc.Send[SKPhysicsJointPin](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new pin joint.
//
// bodyA: The first body to connect. The body must be connected to a node that is
// already part of the scene’s node tree.
//
// bodyB: The second body to connect. The body must be connected to a node that is
// already part of the scene’s node tree.
//
// anchor: The connection point between the two bodies in the scene’s coordinate
// system.
//
// # Return Value
//
// A new pin joint.
//
// # Discussion
//
// You must add the joint to a physics world associated with the scene before
// it takes effect.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointPin/joint(withBodyA:bodyB:anchor:)
func (_SKPhysicsJointPinClass SKPhysicsJointPinClass) JointWithBodyABodyBAnchor(bodyA ISKPhysicsBody, bodyB ISKPhysicsBody, anchor corefoundation.CGPoint) SKPhysicsJointPin {
	rv := objc.Send[objc.ID](objc.ID(_SKPhysicsJointPinClass.class), objc.Sel("jointWithBodyA:bodyB:anchor:"), bodyA, bodyB, anchor)
	return SKPhysicsJointPinFromID(rv)
}

// The speed, in radians per second, at which the physics bodies are driven
// around the pin joint.
//
// # Discussion
//
// The [SKPhysicsJointPin.FrictionTorque] property limits the maximum amount
// of torque that can be applied to the physics bodies.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointPin/rotationSpeed
func (p SKPhysicsJointPin) RotationSpeed() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("rotationSpeed"))
	return rv
}
func (p SKPhysicsJointPin) SetRotationSpeed(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setRotationSpeed:"), value)
}

// A Boolean value that indicates whether the pin joint’s rotation is
// limited to a specific range of values.
//
// # Discussion
//
// The default value is false. If true, the
// [SKPhysicsJointPin.LowerAngleLimit] and [SKPhysicsJointPin.UpperAngleLimit]
// properties are used to limit the angle of the pin joint.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointPin/shouldEnableLimits
func (p SKPhysicsJointPin) ShouldEnableLimits() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("shouldEnableLimits"))
	return rv
}
func (p SKPhysicsJointPin) SetShouldEnableLimits(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setShouldEnableLimits:"), value)
}

// The smallest angle allowed for the pin joint, in radians.
//
// # Discussion
//
// The default value is `0.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointPin/lowerAngleLimit
func (p SKPhysicsJointPin) LowerAngleLimit() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("lowerAngleLimit"))
	return rv
}
func (p SKPhysicsJointPin) SetLowerAngleLimit(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setLowerAngleLimit:"), value)
}

// The largest angle allowed for the pin joint, in radians.
//
// # Discussion
//
// The default value is `0.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointPin/upperAngleLimit
func (p SKPhysicsJointPin) UpperAngleLimit() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("upperAngleLimit"))
	return rv
}
func (p SKPhysicsJointPin) SetUpperAngleLimit(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setUpperAngleLimit:"), value)
}

// The resistance applied by the pin joint to spinning around the anchor
// point.
//
// # Discussion
//
// The range of values is from `0.0` to `1.0`. The default value is `0.0`. If
// a value greater than the default is specified, friction is applied to
// reduce the object’s angular velocity around the pin.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointPin/frictionTorque
func (p SKPhysicsJointPin) FrictionTorque() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("frictionTorque"))
	return rv
}
func (p SKPhysicsJointPin) SetFrictionTorque(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setFrictionTorque:"), value)
}
