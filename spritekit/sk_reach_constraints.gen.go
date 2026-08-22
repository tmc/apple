// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKReachConstraints] class.
var (
	_SKReachConstraintsClass     SKReachConstraintsClass
	_SKReachConstraintsClassOnce sync.Once
)

func getSKReachConstraintsClass() SKReachConstraintsClass {
	_SKReachConstraintsClassOnce.Do(func() {
		_SKReachConstraintsClass = SKReachConstraintsClass{class: objc.GetClass("SKReachConstraints")}
	})
	return _SKReachConstraintsClass
}

// GetSKReachConstraintsClass returns the class object for SKReachConstraints.
func GetSKReachConstraintsClass() SKReachConstraintsClass {
	return getSKReachConstraintsClass()
}

type SKReachConstraintsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKReachConstraintsClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKReachConstraintsClass) Alloc() SKReachConstraints {
	rv := objc.Send[SKReachConstraints](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A specification of the degree of freedom when solving inverse kinematics.
//
// # Overview
//
// An [SKReachConstraints] object is used to describe the range of motion for
// an [SKNode] object whenever an inverse kinematic (IK) action is executed.
// To use reach constraints, create an [SKReachConstraints] object and assign
// it to a node’s [SKNode.ReachConstraints] property. For more information
// on using reach actions to perform IK animations, see the [SKAction] class.
//
// # Working with Reach Constraints
//
//   - [SKReachConstraints.InitWithLowerAngleLimitUpperAngleLimit]: Initializes a new reach constraint object.
//   - [SKReachConstraints.LowerAngleLimit]: The minimum angle that the node can have after it is rotated by a reach event.
//   - [SKReachConstraints.SetLowerAngleLimit]
//   - [SKReachConstraints.UpperAngleLimit]: The maximum angle that the node can have after it is rotated by a reach event.
//   - [SKReachConstraints.SetUpperAngleLimit]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKReachConstraints
type SKReachConstraints struct {
	objectivec.Object
}

// SKReachConstraintsFromID constructs a [SKReachConstraints] from an objc.ID.
//
// A specification of the degree of freedom when solving inverse kinematics.
func SKReachConstraintsFromID(id objc.ID) SKReachConstraints {
	return SKReachConstraints{objectivec.Object{ID: id}}
}

// NOTE: SKReachConstraints adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKReachConstraints] class.
//
// # Working with Reach Constraints
//
//   - [ISKReachConstraints.InitWithLowerAngleLimitUpperAngleLimit]: Initializes a new reach constraint object.
//   - [ISKReachConstraints.LowerAngleLimit]: The minimum angle that the node can have after it is rotated by a reach event.
//   - [ISKReachConstraints.SetLowerAngleLimit]
//   - [ISKReachConstraints.UpperAngleLimit]: The maximum angle that the node can have after it is rotated by a reach event.
//   - [ISKReachConstraints.SetUpperAngleLimit]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKReachConstraints
type ISKReachConstraints interface {
	objectivec.IObject

	// Topic: Working with Reach Constraints

	// Initializes a new reach constraint object.
	InitWithLowerAngleLimitUpperAngleLimit(lowerAngleLimit float64, upperAngleLimit float64) SKReachConstraints
	// The minimum angle that the node can have after it is rotated by a reach event.
	LowerAngleLimit() float64
	SetLowerAngleLimit(value float64)
	// The maximum angle that the node can have after it is rotated by a reach event.
	UpperAngleLimit() float64
	SetUpperAngleLimit(value float64)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (r SKReachConstraints) Init() SKReachConstraints {
	rv := objc.Send[SKReachConstraints](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r SKReachConstraints) Autorelease() SKReachConstraints {
	rv := objc.Send[SKReachConstraints](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKReachConstraints creates a new SKReachConstraints instance.
func NewSKReachConstraints() SKReachConstraints {
	class := getSKReachConstraintsClass()
	rv := objc.Send[SKReachConstraints](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a new reach constraint object.
//
// lowerAngleLimit: The minimum angle that the node can have when it is rotated by a reach
// event.
//
// upperAngleLimit: The maximum angle that the node can have when it is rotated by a reach
// event.
//
// # Return Value
//
// A newly initialized reach constraint.
//
// # Discussion
//
// When a reach action is executed, a node’s [SKNode.ZRotation] property may
// be changed by the action to satisfy the reach action. Any value calculated
// by the reach action for a node is always inside the range specified by the
// reach constraint attached to the node’s [SKNode.ReachConstraints]
// property.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKReachConstraints/init(lowerAngleLimit:upperAngleLimit:)
func NewReachConstraintsWithLowerAngleLimitUpperAngleLimit(lowerAngleLimit float64, upperAngleLimit float64) SKReachConstraints {
	instance := getSKReachConstraintsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLowerAngleLimit:upperAngleLimit:"), lowerAngleLimit, upperAngleLimit)
	return SKReachConstraintsFromID(rv)
}

// Initializes a new reach constraint object.
//
// lowerAngleLimit: The minimum angle that the node can have when it is rotated by a reach
// event.
//
// upperAngleLimit: The maximum angle that the node can have when it is rotated by a reach
// event.
//
// # Return Value
//
// A newly initialized reach constraint.
//
// # Discussion
//
// When a reach action is executed, a node’s [SKNode.ZRotation] property may
// be changed by the action to satisfy the reach action. Any value calculated
// by the reach action for a node is always inside the range specified by the
// reach constraint attached to the node’s [SKNode.ReachConstraints]
// property.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKReachConstraints/init(lowerAngleLimit:upperAngleLimit:)
func (r SKReachConstraints) InitWithLowerAngleLimitUpperAngleLimit(lowerAngleLimit float64, upperAngleLimit float64) SKReachConstraints {
	rv := objc.Send[SKReachConstraints](r.ID, objc.Sel("initWithLowerAngleLimit:upperAngleLimit:"), lowerAngleLimit, upperAngleLimit)
	return rv
}
func (r SKReachConstraints) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The minimum angle that the node can have after it is rotated by a reach
// event.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKReachConstraints/lowerAngleLimit
func (r SKReachConstraints) LowerAngleLimit() float64 {
	rv := objc.Send[float64](r.ID, objc.Sel("lowerAngleLimit"))
	return rv
}
func (r SKReachConstraints) SetLowerAngleLimit(value float64) {
	objc.Send[struct{}](r.ID, objc.Sel("setLowerAngleLimit:"), value)
}

// The maximum angle that the node can have after it is rotated by a reach
// event.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKReachConstraints/upperAngleLimit
func (r SKReachConstraints) UpperAngleLimit() float64 {
	rv := objc.Send[float64](r.ID, objc.Sel("upperAngleLimit"))
	return rv
}
func (r SKReachConstraints) SetUpperAngleLimit(value float64) {
	objc.Send[struct{}](r.ID, objc.Sel("setUpperAngleLimit:"), value)
}
