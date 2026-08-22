// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKPhysicsContact] class.
var (
	_SKPhysicsContactClass     SKPhysicsContactClass
	_SKPhysicsContactClassOnce sync.Once
)

func getSKPhysicsContactClass() SKPhysicsContactClass {
	_SKPhysicsContactClassOnce.Do(func() {
		_SKPhysicsContactClass = SKPhysicsContactClass{class: objc.GetClass("SKPhysicsContact")}
	})
	return _SKPhysicsContactClass
}

// GetSKPhysicsContactClass returns the class object for SKPhysicsContact.
func GetSKPhysicsContactClass() SKPhysicsContactClass {
	return getSKPhysicsContactClass()
}

type SKPhysicsContactClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKPhysicsContactClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKPhysicsContactClass) Alloc() SKPhysicsContact {
	rv := objc.Send[SKPhysicsContact](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A description of the contact between two physics bodies.
//
// # Overview
//
// An [SKPhysicsContact] object is created automatically by SpriteKit to
// describe a contact between two physical bodies in a physics world.To
// receive contact messages, read the [SKScene.PhysicsWorld] property of an
// [SKScene] object you are interested in, and assign its
// [SKPhysicsWorld.ContactDelegate] property to point to an object that
// implements the [SKPhysicsContactDelegate] protocol. Then, for each physics
// body in your scene, set the [SKPhysicsBody.CategoryBitMask] and
// [SKPhysicsBody.ContactTestBitMask] properties to define which interactions
// should generate contact messages.
//
// # Inspecting the Contact Properties
//
//   - [SKPhysicsContact.BodyA]: The first body in the contact.
//   - [SKPhysicsContact.BodyB]: The second body in the contact.
//   - [SKPhysicsContact.ContactPoint]: The contact point between the two physics bodies, in scene coordinates.
//   - [SKPhysicsContact.CollisionImpulse]: The impulse that specifies how hard these two bodies struck each other in Newton-seconds.
//   - [SKPhysicsContact.ContactNormal]: The normal vector specifying the direction of the collision.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsContact
type SKPhysicsContact struct {
	objectivec.Object
}

// SKPhysicsContactFromID constructs a [SKPhysicsContact] from an objc.ID.
//
// A description of the contact between two physics bodies.
func SKPhysicsContactFromID(id objc.ID) SKPhysicsContact {
	return SKPhysicsContact{objectivec.Object{ID: id}}
}

// NOTE: SKPhysicsContact adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKPhysicsContact] class.
//
// # Inspecting the Contact Properties
//
//   - [ISKPhysicsContact.BodyA]: The first body in the contact.
//   - [ISKPhysicsContact.BodyB]: The second body in the contact.
//   - [ISKPhysicsContact.ContactPoint]: The contact point between the two physics bodies, in scene coordinates.
//   - [ISKPhysicsContact.CollisionImpulse]: The impulse that specifies how hard these two bodies struck each other in Newton-seconds.
//   - [ISKPhysicsContact.ContactNormal]: The normal vector specifying the direction of the collision.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsContact
type ISKPhysicsContact interface {
	objectivec.IObject

	// Topic: Inspecting the Contact Properties

	// The first body in the contact.
	BodyA() ISKPhysicsBody
	// The second body in the contact.
	BodyB() ISKPhysicsBody
	// The contact point between the two physics bodies, in scene coordinates.
	ContactPoint() corefoundation.CGPoint
	// The impulse that specifies how hard these two bodies struck each other in Newton-seconds.
	CollisionImpulse() float64
	// The normal vector specifying the direction of the collision.
	ContactNormal() corefoundation.CGVector
}

// Init initializes the instance.
func (p SKPhysicsContact) Init() SKPhysicsContact {
	rv := objc.Send[SKPhysicsContact](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p SKPhysicsContact) Autorelease() SKPhysicsContact {
	rv := objc.Send[SKPhysicsContact](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKPhysicsContact creates a new SKPhysicsContact instance.
func NewSKPhysicsContact() SKPhysicsContact {
	class := getSKPhysicsContactClass()
	rv := objc.Send[SKPhysicsContact](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The first body in the contact.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsContact/bodyA
func (p SKPhysicsContact) BodyA() ISKPhysicsBody {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("bodyA"))
	return SKPhysicsBodyFromID(objc.ID(rv))
}

// The second body in the contact.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsContact/bodyB
func (p SKPhysicsContact) BodyB() ISKPhysicsBody {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("bodyB"))
	return SKPhysicsBodyFromID(objc.ID(rv))
}

// The contact point between the two physics bodies, in scene coordinates.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsContact/contactPoint
func (p SKPhysicsContact) ContactPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](p.ID, objc.Sel("contactPoint"))
	return corefoundation.CGPoint(rv)
}

// The impulse that specifies how hard these two bodies struck each other in
// Newton-seconds.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsContact/collisionImpulse
func (p SKPhysicsContact) CollisionImpulse() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("collisionImpulse"))
	return rv
}

// The normal vector specifying the direction of the collision.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsContact/contactNormal
func (p SKPhysicsContact) ContactNormal() corefoundation.CGVector {
	rv := objc.Send[corefoundation.CGVector](p.ID, objc.Sel("contactNormal"))
	return corefoundation.CGVector(rv)
}
