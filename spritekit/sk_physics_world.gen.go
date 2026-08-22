// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKPhysicsWorld] class.
var (
	_SKPhysicsWorldClass     SKPhysicsWorldClass
	_SKPhysicsWorldClassOnce sync.Once
)

func getSKPhysicsWorldClass() SKPhysicsWorldClass {
	_SKPhysicsWorldClassOnce.Do(func() {
		_SKPhysicsWorldClass = SKPhysicsWorldClass{class: objc.GetClass("SKPhysicsWorld")}
	})
	return _SKPhysicsWorldClass
}

// GetSKPhysicsWorldClass returns the class object for SKPhysicsWorld.
func GetSKPhysicsWorldClass() SKPhysicsWorldClass {
	return getSKPhysicsWorldClass()
}

type SKPhysicsWorldClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKPhysicsWorldClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKPhysicsWorldClass) Alloc() SKPhysicsWorld {
	rv := objc.Send[SKPhysicsWorld](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// The driver of the physics engine in a scene; it exposes the ability for you
// to configure and query the physics system.
//
// # Overview
//
// [SKPhysicsWorld] runs the physics engine of a scene and is the place that
// contact detection occurs. Do not create a [SKPhysicsWorld] directly; the
// system creates a physics world and adds it to the scene’s
// [SKScene.PhysicsWorld] property.
//
// The physics world allows you to:
//
// - Set important properties like [SKPhysicsWorld.Gravity] - Join two physics
// bodies using an [SKPhysicsJoint] - Respond to collision between two physics
// bodies using [SKPhysicsWorld.ContactDelegate] - Do custom collisions
// detection or hit testing
//
// # Configuring the Physics World
//
//   - [SKPhysicsWorld.Gravity]: A vector that specifies the gravitational acceleration applied to physics bodies in the physics world.
//   - [SKPhysicsWorld.SetGravity]
//   - [SKPhysicsWorld.Speed]: The rate at which the simulation executes.
//   - [SKPhysicsWorld.SetSpeed]
//
// # Joining Physics Bodies with Joints
//
//   - [SKPhysicsWorld.AddJoint]: Adds a joint to the physics world.
//   - [SKPhysicsWorld.RemoveAllJoints]: Removes all joints from the physics world.
//   - [SKPhysicsWorld.RemoveJoint]: Removes a specific joint from the physics world.
//
// # Detecting Collisions
//
//   - [SKPhysicsWorld.ContactDelegate]: A delegate that is called when two physics bodies come in contact with each other.
//   - [SKPhysicsWorld.SetContactDelegate]
//
// # Searching the Scene for Physics Bodies
//
//   - [SKPhysicsWorld.BodyAlongRayStartEnd]: Searches for the first physics body that intersects a ray.
//   - [SKPhysicsWorld.BodyAtPoint]: Searches for the first physics body that contains a point.
//   - [SKPhysicsWorld.BodyInRect]: Searches for the first physics body that intersects the specified rectangle.
//   - [SKPhysicsWorld.EnumerateBodiesAlongRayStartEndUsingBlock]: Enumerates all the physics bodies in the scene that intersect a ray.
//   - [SKPhysicsWorld.EnumerateBodiesAtPointUsingBlock]: Enumerates all the physics bodies in the scene that contain a point.
//   - [SKPhysicsWorld.EnumerateBodiesInRectUsingBlock]: Enumerates all the physics bodies in the scene that intersect the specified rectangle.
//
// # Sampling Physics Fields
//
//   - [SKPhysicsWorld.SampleFieldsAt]: Samples all of the field nodes in the scene and returns the summation of their forces at that point.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld
type SKPhysicsWorld struct {
	objectivec.Object
}

// SKPhysicsWorldFromID constructs a [SKPhysicsWorld] from an objc.ID.
//
// The driver of the physics engine in a scene; it exposes the ability for you
// to configure and query the physics system.
func SKPhysicsWorldFromID(id objc.ID) SKPhysicsWorld {
	return SKPhysicsWorld{objectivec.Object{ID: id}}
}

// NOTE: SKPhysicsWorld adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKPhysicsWorld] class.
//
// # Configuring the Physics World
//
//   - [ISKPhysicsWorld.Gravity]: A vector that specifies the gravitational acceleration applied to physics bodies in the physics world.
//   - [ISKPhysicsWorld.SetGravity]
//   - [ISKPhysicsWorld.Speed]: The rate at which the simulation executes.
//   - [ISKPhysicsWorld.SetSpeed]
//
// # Joining Physics Bodies with Joints
//
//   - [ISKPhysicsWorld.AddJoint]: Adds a joint to the physics world.
//   - [ISKPhysicsWorld.RemoveAllJoints]: Removes all joints from the physics world.
//   - [ISKPhysicsWorld.RemoveJoint]: Removes a specific joint from the physics world.
//
// # Detecting Collisions
//
//   - [ISKPhysicsWorld.ContactDelegate]: A delegate that is called when two physics bodies come in contact with each other.
//   - [ISKPhysicsWorld.SetContactDelegate]
//
// # Searching the Scene for Physics Bodies
//
//   - [ISKPhysicsWorld.BodyAlongRayStartEnd]: Searches for the first physics body that intersects a ray.
//   - [ISKPhysicsWorld.BodyAtPoint]: Searches for the first physics body that contains a point.
//   - [ISKPhysicsWorld.BodyInRect]: Searches for the first physics body that intersects the specified rectangle.
//   - [ISKPhysicsWorld.EnumerateBodiesAlongRayStartEndUsingBlock]: Enumerates all the physics bodies in the scene that intersect a ray.
//   - [ISKPhysicsWorld.EnumerateBodiesAtPointUsingBlock]: Enumerates all the physics bodies in the scene that contain a point.
//   - [ISKPhysicsWorld.EnumerateBodiesInRectUsingBlock]: Enumerates all the physics bodies in the scene that intersect the specified rectangle.
//
// # Sampling Physics Fields
//
//   - [ISKPhysicsWorld.SampleFieldsAt]: Samples all of the field nodes in the scene and returns the summation of their forces at that point.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld
type ISKPhysicsWorld interface {
	objectivec.IObject

	// Topic: Configuring the Physics World

	// A vector that specifies the gravitational acceleration applied to physics bodies in the physics world.
	Gravity() corefoundation.CGVector
	SetGravity(value corefoundation.CGVector)
	// The rate at which the simulation executes.
	Speed() float64
	SetSpeed(value float64)

	// Topic: Joining Physics Bodies with Joints

	// Adds a joint to the physics world.
	AddJoint(joint ISKPhysicsJoint)
	// Removes all joints from the physics world.
	RemoveAllJoints()
	// Removes a specific joint from the physics world.
	RemoveJoint(joint ISKPhysicsJoint)

	// Topic: Detecting Collisions

	// A delegate that is called when two physics bodies come in contact with each other.
	ContactDelegate() SKPhysicsContactDelegate
	SetContactDelegate(value SKPhysicsContactDelegate)

	// Topic: Searching the Scene for Physics Bodies

	// Searches for the first physics body that intersects a ray.
	BodyAlongRayStartEnd(start corefoundation.CGPoint, end corefoundation.CGPoint) ISKPhysicsBody
	// Searches for the first physics body that contains a point.
	BodyAtPoint(point corefoundation.CGPoint) ISKPhysicsBody
	// Searches for the first physics body that intersects the specified rectangle.
	BodyInRect(rect corefoundation.CGRect) ISKPhysicsBody
	// Enumerates all the physics bodies in the scene that intersect a ray.
	EnumerateBodiesAlongRayStartEndUsingBlock(start corefoundation.CGPoint, end corefoundation.CGPoint, block SKPhysicsBodyCGPointCGVectorBoolHandler)
	// Enumerates all the physics bodies in the scene that contain a point.
	EnumerateBodiesAtPointUsingBlock(point corefoundation.CGPoint, block SKPhysicsBodyBoolHandler)
	// Enumerates all the physics bodies in the scene that intersect the specified rectangle.
	EnumerateBodiesInRectUsingBlock(rect corefoundation.CGRect, block SKPhysicsBodyBoolHandler)

	// Topic: Sampling Physics Fields

	// Samples all of the field nodes in the scene and returns the summation of their forces at that point.
	SampleFieldsAt(position Vector_float3) Vector_float3

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (p SKPhysicsWorld) Init() SKPhysicsWorld {
	rv := objc.Send[SKPhysicsWorld](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p SKPhysicsWorld) Autorelease() SKPhysicsWorld {
	rv := objc.Send[SKPhysicsWorld](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKPhysicsWorld creates a new SKPhysicsWorld instance.
func NewSKPhysicsWorld() SKPhysicsWorld {
	class := getSKPhysicsWorldClass()
	rv := objc.Send[SKPhysicsWorld](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Adds a joint to the physics world.
//
// joint: The joint to add.
//
// # Discussion
//
// For a joint to take effect, it must be added to the physics world.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld/add(_:)
func (p SKPhysicsWorld) AddJoint(joint ISKPhysicsJoint) {
	objc.Send[objc.ID](p.ID, objc.Sel("addJoint:"), joint)
}

// Removes all joints from the physics world.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld/removeAllJoints()
func (p SKPhysicsWorld) RemoveAllJoints() {
	objc.Send[objc.ID](p.ID, objc.Sel("removeAllJoints"))
}

// Removes a specific joint from the physics world.
//
// joint: The joint to remove.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld/remove(_:)
func (p SKPhysicsWorld) RemoveJoint(joint ISKPhysicsJoint) {
	objc.Send[objc.ID](p.ID, objc.Sel("removeJoint:"), joint)
}

// Searches for the first physics body that intersects a ray.
//
// start: The starting point for the ray in scene coordinates.
//
// end: The ending point for the ray in scene coordinates.
//
// # Return Value
//
// The first physics body discovered that intersects the ray. This may be any
// body along the ray; it is not guaranteed to be the closest physics body. If
// no body intersects the ray, this method returns `nil`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld/body(alongRayStart:end:)
func (p SKPhysicsWorld) BodyAlongRayStartEnd(start corefoundation.CGPoint, end corefoundation.CGPoint) ISKPhysicsBody {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("bodyAlongRayStart:end:"), start, end)
	return SKPhysicsBodyFromID(rv)
}

// Searches for the first physics body that contains a point.
//
// point: A point in scene coordinates.
//
// # Return Value
//
// The first physics body discovered that contains the point. If no body
// contains the point, this method returns `nil`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld/body(at:)
func (p SKPhysicsWorld) BodyAtPoint(point corefoundation.CGPoint) ISKPhysicsBody {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("bodyAtPoint:"), point)
	return SKPhysicsBodyFromID(rv)
}

// Searches for the first physics body that intersects the specified
// rectangle.
//
// rect: A rectangle in scene coordinates.
//
// # Return Value
//
// The first physics body discovered that intersects the rectangle. If no body
// intersects the rectangle, this method returns `nil`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld/body(in:)
func (p SKPhysicsWorld) BodyInRect(rect corefoundation.CGRect) ISKPhysicsBody {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("bodyInRect:"), rect)
	return SKPhysicsBodyFromID(rv)
}

// Enumerates all the physics bodies in the scene that intersect a ray.
//
// start: The starting point for the ray in scene coordinates.
//
// end: The ending point for the ray in scene coordinates.
//
// block: A block to be called for each physics body that the ray touches. The block
// takes the following parameters:
//
// body: The physics body that the ray intersected. point: The point in scene
// coordinates where the ray contacted the physics body. normal: The normal
// vector for the physics body at the point of contact. stop: A pointer to a
// Boolean variable. Your block can set this to true to terminate the
// enumeration.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld/enumerateBodies(alongRayStart:end:using:)
func (p SKPhysicsWorld) EnumerateBodiesAlongRayStartEndUsingBlock(start corefoundation.CGPoint, end corefoundation.CGPoint, block SKPhysicsBodyCGPointCGVectorBoolHandler) {
	_block2, _ := NewSKPhysicsBodyCGPointCGVectorBoolBlock(block)
	objc.Send[objc.ID](p.ID, objc.Sel("enumerateBodiesAlongRayStart:end:usingBlock:"), start, end, _block2)
}

// Enumerates all the physics bodies in the scene that contain a point.
//
// point: A point in scene coordinates.
//
// block: A block to be called for each physics body that contains the point. The
// block takes the following parameters:
//
// body: The physics body that the ray intersected. stop: A pointer to a
// Boolean variable. Your block can set this to true to terminate the
// enumeration.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld/enumerateBodies(at:using:)
func (p SKPhysicsWorld) EnumerateBodiesAtPointUsingBlock(point corefoundation.CGPoint, block SKPhysicsBodyBoolHandler) {
	_block1, _ := NewSKPhysicsBodyBoolBlock(block)
	objc.Send[objc.ID](p.ID, objc.Sel("enumerateBodiesAtPoint:usingBlock:"), point, _block1)
}

// Enumerates all the physics bodies in the scene that intersect the specified
// rectangle.
//
// rect: A rectangle in scene coordinates.
//
// block: A block to be called for each physics body that contains the point. The
// block takes the following parameters:
//
// body: The physics body that intersected the rectangle. stop: A pointer to a
// Boolean variable. Your block can set this to true to terminate the
// enumeration.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld/enumerateBodies(in:using:)
func (p SKPhysicsWorld) EnumerateBodiesInRectUsingBlock(rect corefoundation.CGRect, block SKPhysicsBodyBoolHandler) {
	_block1, _ := NewSKPhysicsBodyBoolBlock(block)
	objc.Send[objc.ID](p.ID, objc.Sel("enumerateBodiesInRect:usingBlock:"), rect, _block1)
}

// Samples all of the field nodes in the scene and returns the summation of
// their forces at that point.
//
// position: A position in scene coordinates.
//
// # Return Value
//
// The summation of forces exerted on that point.
//
// # Discussion
//
// The sample is calculated as if a physics body is placed at that position in
// the scene. The body is assumed to have a mass of `1.0`, with no charge or
// velocity. The body is affected by all field nodes.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld/sampleFields(at:)
func (p SKPhysicsWorld) SampleFieldsAt(position Vector_float3) Vector_float3 {
	rv := objc.Send[Vector_float3](p.ID, objc.Sel("sampleFieldsAt:"), position)
	return Vector_float3(rv)
}
func (p SKPhysicsWorld) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A vector that specifies the gravitational acceleration applied to physics
// bodies in the physics world.
//
// # Discussion
//
// The components of this property are measured in meters per second. The
// default value is `(0.0,-9.8)`, which represent’s Earth’s gravity.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld/gravity
func (p SKPhysicsWorld) Gravity() corefoundation.CGVector {
	rv := objc.Send[corefoundation.CGVector](p.ID, objc.Sel("gravity"))
	return corefoundation.CGVector(rv)
}
func (p SKPhysicsWorld) SetGravity(value corefoundation.CGVector) {
	objc.Send[struct{}](p.ID, objc.Sel("setGravity:"), value)
}

// The rate at which the simulation executes.
//
// # Discussion
//
// The default value is `1.0`, which means the simulation runs at normal
// speed. A value other than the default changes the rate at which time passes
// in the physics simulation. For example, a speed value of `2.0` indicates
// that time in the physics simulation passes twice as fast as the scene’s
// simulation time. A value of `0.0` pauses the physics simulation.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld/speed
func (p SKPhysicsWorld) Speed() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("speed"))
	return rv
}
func (p SKPhysicsWorld) SetSpeed(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setSpeed:"), value)
}

// A delegate that is called when two physics bodies come in contact with each
// other.
//
// # Discussion
//
// A contact is created when two physics bodies overlap and one of the physics
// bodies has a [SKPhysicsBody.ContactTestBitMask] property that overlaps with
// the other body’s [SKPhysicsBody.CategoryBitMask] property. By default, a
// physics body’s [SKPhysicsBody.ContactTestBitMask] is set to all bits
// cleared.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld/contactDelegate
func (p SKPhysicsWorld) ContactDelegate() SKPhysicsContactDelegate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("contactDelegate"))
	return SKPhysicsContactDelegateObjectFromID(rv)
}
func (p SKPhysicsWorld) SetContactDelegate(value SKPhysicsContactDelegate) {
	objc.Send[struct{}](p.ID, objc.Sel("setContactDelegate:"), value)
}
