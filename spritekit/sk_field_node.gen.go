// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [SKFieldNode] class.
var (
	_SKFieldNodeClass     SKFieldNodeClass
	_SKFieldNodeClassOnce sync.Once
)

func getSKFieldNodeClass() SKFieldNodeClass {
	_SKFieldNodeClassOnce.Do(func() {
		_SKFieldNodeClass = SKFieldNodeClass{class: objc.GetClass("SKFieldNode")}
	})
	return _SKFieldNodeClass
}

// GetSKFieldNodeClass returns the class object for SKFieldNode.
func GetSKFieldNodeClass() SKFieldNodeClass {
	return getSKFieldNodeClass()
}

type SKFieldNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKFieldNodeClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKFieldNodeClass) Alloc() SKFieldNode {
	rv := objc.Send[SKFieldNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A node that applies physics effects to nearby nodes.
//
// # Overview
//
// There are many different kinds of field nodes that can be created, each
// with different effects. The [SKFieldNode] section lists the field types you
// can create using SpriteKit, including a type that allows you to apply
// custom forces to physics bodies. Instantiate the appropriate kind of field
// node and then add it to the scene’s node tree.
//
// When the scene simulates physics effects, a field node applies its effect
// to a physics body so long as the following are true:
//
// - The field node is in the scene’s node tree. - The field node’s
// [SKFieldNode.Enabled] property is true. - The physics body is attached to a
// node that is in the scene’s node tree. - The physics body is located
// inside the field node’s region (see [SKFieldNode.Region]). - The physics
// body is not located inside the region of another field node whose
// [SKFieldNode.Exclusive] property is set to true. - A logical AND operation
// between the field node’s [SKFieldNode.CategoryBitMask] property and the
// physics body’s [SKPhysicsBody.FieldBitMask] property results in a nonzero
// value.
//
// # Determining Which Physics Bodies Are Affected by the Field
//
//   - [SKFieldNode.IsEnabled]: A Boolean value that indicates whether the field is active.
//   - [SKFieldNode.SetEnabled]
//   - [SKFieldNode.IsExclusive]: A Boolean value that indicates whether the field node should override all other field nodes that might otherwise affect physics bodies.
//   - [SKFieldNode.SetExclusive]
//   - [SKFieldNode.Region]: The area (relative to the node’s origin) that the field affects.
//   - [SKFieldNode.SetRegion]
//   - [SKFieldNode.MinimumRadius]: The minimum value for distance-based effects.
//   - [SKFieldNode.SetMinimumRadius]
//   - [SKFieldNode.CategoryBitMask]: A mask that defines which categories this field belongs to.
//   - [SKFieldNode.SetCategoryBitMask]
//
// # Configuring the Strength of the Field
//
//   - [SKFieldNode.Strength]: The strength of the field.
//   - [SKFieldNode.SetStrength]
//   - [SKFieldNode.Falloff]: The exponent that defines the rate of decay for the strength of the field as the distance increases between the node and the physics body being affected.
//   - [SKFieldNode.SetFalloff]
//
// # Configuring Other Field Properties
//
//   - [SKFieldNode.AnimationSpeed]: The rate at which a noise or turbulence field node changes.
//   - [SKFieldNode.SetAnimationSpeed]
//   - [SKFieldNode.Smoothness]: The smoothness of the noise used to generate the forces.
//   - [SKFieldNode.SetSmoothness]
//   - [SKFieldNode.Direction]: The direction of a velocity field node.
//   - [SKFieldNode.SetDirection]
//   - [SKFieldNode.Texture]: A normal texture that specifies the velocities at different points in a velocity field node.
//   - [SKFieldNode.SetTexture]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode
type SKFieldNode struct {
	SKNode
}

// SKFieldNodeFromID constructs a [SKFieldNode] from an objc.ID.
//
// A node that applies physics effects to nearby nodes.
func SKFieldNodeFromID(id objc.ID) SKFieldNode {
	return SKFieldNode{SKNode: SKNodeFromID(id)}
}

// NOTE: SKFieldNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKFieldNode] class.
//
// # Determining Which Physics Bodies Are Affected by the Field
//
//   - [ISKFieldNode.IsEnabled]: A Boolean value that indicates whether the field is active.
//   - [ISKFieldNode.SetEnabled]
//   - [ISKFieldNode.IsExclusive]: A Boolean value that indicates whether the field node should override all other field nodes that might otherwise affect physics bodies.
//   - [ISKFieldNode.SetExclusive]
//   - [ISKFieldNode.Region]: The area (relative to the node’s origin) that the field affects.
//   - [ISKFieldNode.SetRegion]
//   - [ISKFieldNode.MinimumRadius]: The minimum value for distance-based effects.
//   - [ISKFieldNode.SetMinimumRadius]
//   - [ISKFieldNode.CategoryBitMask]: A mask that defines which categories this field belongs to.
//   - [ISKFieldNode.SetCategoryBitMask]
//
// # Configuring the Strength of the Field
//
//   - [ISKFieldNode.Strength]: The strength of the field.
//   - [ISKFieldNode.SetStrength]
//   - [ISKFieldNode.Falloff]: The exponent that defines the rate of decay for the strength of the field as the distance increases between the node and the physics body being affected.
//   - [ISKFieldNode.SetFalloff]
//
// # Configuring Other Field Properties
//
//   - [ISKFieldNode.AnimationSpeed]: The rate at which a noise or turbulence field node changes.
//   - [ISKFieldNode.SetAnimationSpeed]
//   - [ISKFieldNode.Smoothness]: The smoothness of the noise used to generate the forces.
//   - [ISKFieldNode.SetSmoothness]
//   - [ISKFieldNode.Direction]: The direction of a velocity field node.
//   - [ISKFieldNode.SetDirection]
//   - [ISKFieldNode.Texture]: A normal texture that specifies the velocities at different points in a velocity field node.
//   - [ISKFieldNode.SetTexture]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode
type ISKFieldNode interface {
	ISKNode

	// Topic: Determining Which Physics Bodies Are Affected by the Field

	// A Boolean value that indicates whether the field is active.
	IsEnabled() bool
	SetEnabled(value bool)
	// A Boolean value that indicates whether the field node should override all other field nodes that might otherwise affect physics bodies.
	IsExclusive() bool
	SetExclusive(value bool)
	// The area (relative to the node’s origin) that the field affects.
	Region() ISKRegion
	SetRegion(value ISKRegion)
	// The minimum value for distance-based effects.
	MinimumRadius() float32
	SetMinimumRadius(value float32)
	// A mask that defines which categories this field belongs to.
	CategoryBitMask() uint32
	SetCategoryBitMask(value uint32)

	// Topic: Configuring the Strength of the Field

	// The strength of the field.
	Strength() float32
	SetStrength(value float32)
	// The exponent that defines the rate of decay for the strength of the field as the distance increases between the node and the physics body being affected.
	Falloff() float32
	SetFalloff(value float32)

	// Topic: Configuring Other Field Properties

	// The rate at which a noise or turbulence field node changes.
	AnimationSpeed() float32
	SetAnimationSpeed(value float32)
	// The smoothness of the noise used to generate the forces.
	Smoothness() float32
	SetSmoothness(value float32)
	// The direction of a velocity field node.
	Direction() Vector_float3
	SetDirection(value Vector_float3)
	// A normal texture that specifies the velocities at different points in a velocity field node.
	Texture() ISKTexture
	SetTexture(value ISKTexture)
}

// Init initializes the instance.
func (f SKFieldNode) Init() SKFieldNode {
	rv := objc.Send[SKFieldNode](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f SKFieldNode) Autorelease() SKFieldNode {
	rv := objc.Send[SKFieldNode](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKFieldNode creates a new SKFieldNode instance.
func NewSKFieldNode() SKFieldNode {
	class := getSKFieldNodeClass()
	rv := objc.Send[SKFieldNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Called when a node is initialized from an .sks file.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(coder:)
func NewFieldNodeWithCoder(aDecoder foundation.INSCoder) SKFieldNode {
	instance := getSKFieldNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return SKFieldNodeFromID(rv)
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
func NewFieldNodeWithFileNamed(filename string) SKFieldNode {
	rv := objc.Send[objc.ID](objc.ID(getSKFieldNodeClass().class), objc.Sel("nodeWithFileNamed:"), objc.String(filename))
	return SKFieldNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(fileNamed:securelyWithClasses:)
func NewFieldNodeWithFileNamedSecurelyWithClassesAndError(filename string, classes foundation.INSSet) (SKFieldNode, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getSKFieldNodeClass().class), objc.Sel("nodeWithFileNamed:securelyWithClasses:andError:"), objc.String(filename), classes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SKFieldNode{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SKFieldNode{}, objc.ErrInitFailed
	}
	return SKFieldNodeFromID(rv), nil
}

// Creates a field node that applies a force that resists the motion of
// physics bodies.
//
// # Return Value
//
// A new drag field node.
//
// # Discussion
//
// The force is applied in the opposite direction of the physics body’s
// [SKPhysicsBody.Velocity] property and has a magnitude proportional to the
// field’s [SKFieldNode.Strength] property and the physics body’s speed.
// This field models Stoke’s Law.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/dragField()
func (_SKFieldNodeClass SKFieldNodeClass) DragField() SKFieldNode {
	rv := objc.Send[objc.ID](objc.ID(_SKFieldNodeClass.class), objc.Sel("dragField"))
	return SKFieldNodeFromID(rv)
}

// Creates a field node that applies an electrical force proportional to the
// electrical charge of physics bodies.
//
// # Return Value
//
// A new electric field node.
//
// # Discussion
//
// The force points toward the field node’s position and has a magnitude
// proportional to the field’s [SKFieldNode.Strength] property and the
// physics body’s [SKPhysicsBody.Charge] property. This field models the
// first part of the Lorentz equation:
//
// F = qE
//
// Where F equals force, q equals charge, and E equals electric field.
//
// The [SKFieldNode.Falloff] property of an electrical field node is set by
// default to `2`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/electricField()
func (_SKFieldNodeClass SKFieldNodeClass) ElectricField() SKFieldNode {
	rv := objc.Send[objc.ID](objc.ID(_SKFieldNodeClass.class), objc.Sel("electricField"))
	return SKFieldNodeFromID(rv)
}

// Creates a field node that accelerates physics bodies in a specific
// direction.
//
// direction: The direction and magnitude of the gravitational force. The values
// represent the acceleration of the field in meters per second squared. For
// example, to simulate earth’s gravity, specify `(0, -9.8, 0)`. The `z`
// component on the vector is ignored.
//
// # Return Value
//
// A new linear gravity field node.
//
// # Discussion
//
// If the field node is rotated, the direction of its gravity field is also
// rotated. The calculated force is proportional to the physics body’s mass
// (meaning that the acceleration applied to all affected physics bodies is a
// constant).
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/linearGravityField(withVector:)
func (_SKFieldNodeClass SKFieldNodeClass) LinearGravityFieldWithVector(direction Vector_float3) SKFieldNode {
	rv := objc.Send[objc.ID](objc.ID(_SKFieldNodeClass.class), objc.Sel("linearGravityFieldWithVector:"), direction)
	return SKFieldNodeFromID(rv)
}

// Creates a field node that applies a magnetic force based on the velocity
// and electrical charge of the physics bodies.
//
// # Return Value
//
// A new magnetic field node.
//
// # Discussion
//
// The force generated by this field is directed on a line that is determined
// by calculating the cross-product between the direction of the the physics
// body’s [SKPhysicsBody.Velocity] property and a line traced between the
// field node and the physics body. This field models the second part of the
// Lorentz equation:
//
// F = qvB
//
// Where F equals force, q equals charge, v equals velocity, B equals magnetic
// field and E equals electric field.
//
// The force has a magnitude proportional to the field’s
// [SKFieldNode.Strength] property and the physics body’s
// [SKPhysicsBody.Charge] and [SKPhysicsBody.Velocity] properties. Therefore,
// physics bodies that are either stationary or with a charge of zero will not
// be affected by a magnetic field. Magnetic fields with a negative strength
// value impart a clockwise spin on the physics bodies they affect, while a
// positive strength give a clockwise spin.The [SKFieldNode.Falloff] property
// of a magnetic field node is set by default to `2`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/magneticField()
func (_SKFieldNodeClass SKFieldNodeClass) MagneticField() SKFieldNode {
	rv := objc.Send[objc.ID](objc.ID(_SKFieldNodeClass.class), objc.Sel("magneticField"))
	return SKFieldNodeFromID(rv)
}

// Creates a field node that applies a randomized acceleration to physics
// bodies.
//
// smoothness: The smoothness of the noise used to generate the forces. This parameter
// should be a value between `0.0` and `1.0`, where `1.0` represents a uniform
// smoothness.
//
// speed: The speed at which the noise field should change. A value of `0.0` means
// that the field should not animate at all.
//
// # Return Value
//
// A new noise field node.
//
// # Discussion
//
// Use a noise field to simulate effects such as fireflies or snow.
//
// The acceleration is proportional to the field strength in a pseudo-random
// direction. The calculated force is proportional to the physics body’s
// mass (meaning that the acceleration applied to all affected physics bodies
// is a constant).
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/noiseField(withSmoothness:animationSpeed:)
func (_SKFieldNodeClass SKFieldNodeClass) NoiseFieldWithSmoothnessAnimationSpeed(smoothness float64, speed float64) SKFieldNode {
	rv := objc.Send[objc.ID](objc.ID(_SKFieldNodeClass.class), objc.Sel("noiseFieldWithSmoothness:animationSpeed:"), smoothness, speed)
	return SKFieldNodeFromID(rv)
}

// Creates a field node that accelerates physics bodies toward the field node.
//
// # Return Value
//
// A new radial gravity field node.
//
// # Discussion
//
// The strength of the field measures the acceleration of the field in meters
// per second squared. A positive number indicates that the body is
// accelerating toward the field node. The calculated force is proportional to
// the physics body’s mass (meaning that the acceleration applied to all
// affected physics bodies is a constant). The [SKFieldNode.Falloff] property
// of a radial gravity field node is set by default to `2`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/radialGravityField()
func (_SKFieldNodeClass SKFieldNodeClass) RadialGravityField() SKFieldNode {
	rv := objc.Send[objc.ID](objc.ID(_SKFieldNodeClass.class), objc.Sel("radialGravityField"))
	return SKFieldNodeFromID(rv)
}

// Creates a field node that applies a spring-like force that pulls physics
// bodies toward the field node.
//
// # Return Value
//
// A new spring field node.
//
// # Discussion
//
// The strength of the field measures the strength of the spring. A positive
// number indicates that the body is accelerating toward the field node. The
// [SKFieldNode.Falloff] property of a spring field node is set by default to
// `-1`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/springField()
func (_SKFieldNodeClass SKFieldNodeClass) SpringField() SKFieldNode {
	rv := objc.Send[objc.ID](objc.ID(_SKFieldNodeClass.class), objc.Sel("springField"))
	return SKFieldNodeFromID(rv)
}

// Creates a field node that applies a randomized acceleration to physics
// bodies.
//
// smoothness: The smoothness of the noise used to generate the forces. This parameter
// should be a value between `0.0` and `1.0`, where `1.0` represents a uniform
// smoothness.
//
// speed: The speed at which the noise field should change. A value of `0.0` means
// that the field should not animate at all.
//
// # Return Value
//
// A new turbulence field node.
//
// # Discussion
//
// The acceleration’s magnitude is proportional to a body’s velocity.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/turbulenceField(withSmoothness:animationSpeed:)
func (_SKFieldNodeClass SKFieldNodeClass) TurbulenceFieldWithSmoothnessAnimationSpeed(smoothness float64, speed float64) SKFieldNode {
	rv := objc.Send[objc.ID](objc.ID(_SKFieldNodeClass.class), objc.Sel("turbulenceFieldWithSmoothness:animationSpeed:"), smoothness, speed)
	return SKFieldNodeFromID(rv)
}

// Creates a field node that sets the velocity of physics bodies that enter
// the node’s area based on the pixel values of a texture.
//
// velocityTexture: A normal texture used to specify the velocities at different points in the
// field.
//
// # Return Value
//
// A new velocity field node.
//
// # Discussion
//
// When a physics body is affected, its new velocity in each frame is
// calculated by performing a texture lookup (treating the value as a normal
// vector) and then multiplying that vector by the strength of the field. The
// field has an implicit size (region) equal to the size of the texture;
// physics bodies outside this area are unaffected by the field node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/velocityField(with:)
func (_SKFieldNodeClass SKFieldNodeClass) VelocityFieldWithTexture(velocityTexture ISKTexture) SKFieldNode {
	rv := objc.Send[objc.ID](objc.ID(_SKFieldNodeClass.class), objc.Sel("velocityFieldWithTexture:"), velocityTexture)
	return SKFieldNodeFromID(rv)
}

// Creates a field node that gives physics bodies a constant velocity.
//
// direction: The velocity that any affected physics bodies will have. The `z` component
// on the vector is ignored.
//
// # Return Value
//
// A new velocity field node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/velocityField(withVector:)
func (_SKFieldNodeClass SKFieldNodeClass) VelocityFieldWithVector(direction Vector_float3) SKFieldNode {
	rv := objc.Send[objc.ID](objc.ID(_SKFieldNodeClass.class), objc.Sel("velocityFieldWithVector:"), direction)
	return SKFieldNodeFromID(rv)
}

// Creates a field node that applies a perpendicular force to physics bodies.
//
// # Return Value
//
// A new vortex field node.
//
// # Discussion
//
// The strength of the field measures the acceleration of the field in meters
// per second squared (meaning that similar to gravity, all physics bodies are
// affected equally). The physics body is accelerated along the perpendicular
// of the line between the field node’s position and the position of the
// physics body. A positive field strength indicates the body is accelerated
// in a counter-clockwise direction. The [SKFieldNode.Falloff] property of a
// vortex field node is set by default to `2`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/vortexField()
func (_SKFieldNodeClass SKFieldNodeClass) VortexField() SKFieldNode {
	rv := objc.Send[objc.ID](objc.ID(_SKFieldNodeClass.class), objc.Sel("vortexField"))
	return SKFieldNodeFromID(rv)
}

// Creates a field node that calculates and applies a custom force to the
// physics body.
//
// block: A custom block to be executed when a physics body is affected by the field.
// Your block should calculate and return the force to be applied to the body.
//
// # Return Value
//
// A new custom field node.
//
// # Discussion
//
// The value returned by the custom block is a vector for an impulse force
// which is applied to the physics body being evaluated for that frame. Only
// the `x` and `y` components of the return value are used by SpriteKit, the
// `z` component is ignored.
//
// The values passed into the block by the `position` and `velocity` arguments
// measured in meters: if you need to convert them into points — as used by
// SpriteKit — multiply the values by 150.
//
// The following code shows how to create a custom field to emulate drag. The
// block returns the negative of the square root of the velocity of the
// physics body. This decelerates a physics body passing through the
// [SKFieldNode] object’s region.
//
// Listing 1. Creating a custom drag field
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/customField(evaluationBlock:)
func (_SKFieldNodeClass SKFieldNodeClass) CustomFieldWithEvaluationBlock(block Float32Float32Handler) SKFieldNode {
	_block0, _ := NewFloat32Float32Block(block)
	rv := objc.Send[objc.ID](objc.ID(_SKFieldNodeClass.class), objc.Sel("customFieldWithEvaluationBlock:"), _block0)
	return SKFieldNodeFromID(rv)
}

// A Boolean value that indicates whether the field is active.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/isEnabled
func (f SKFieldNode) IsEnabled() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isEnabled"))
	return rv
}
func (f SKFieldNode) SetEnabled(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setEnabled:"), value)
}

// A Boolean value that indicates whether the field node should override all
// other field nodes that might otherwise affect physics bodies.
//
// # Discussion
//
// If the value is set to true and a physics body is within this field’s
// region, all other field nodes that might otherwise affect this body are
// ignored. The default value is false.
//
// If you set this property to true on multiple field nodes within a scene,
// their regions should not overlap. If they do, the results are undefined.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/isExclusive
func (f SKFieldNode) IsExclusive() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isExclusive"))
	return rv
}
func (f SKFieldNode) SetExclusive(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setExclusive:"), value)
}

// The area (relative to the node’s origin) that the field affects.
//
// # Discussion
//
// A field node applies its effect to all physics bodies that are partially or
// completely inside its region. The default value is a region of infinite
// size.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/region
func (f SKFieldNode) Region() ISKRegion {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("region"))
	return SKRegionFromID(objc.ID(rv))
}
func (f SKFieldNode) SetRegion(value ISKRegion) {
	objc.Send[struct{}](f.ID, objc.Sel("setRegion:"), value)
}

// The minimum value for distance-based effects.
//
// # Discussion
//
// When the distance between the node and a physics body is calculated, any
// distance shorter than the value stored in the [SKFieldNode.MinimumRadius]
// property is treated as if it is equal to it. The default value is a very
// small (but nonzero) value.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/minimumRadius
func (f SKFieldNode) MinimumRadius() float32 {
	rv := objc.Send[float32](f.ID, objc.Sel("minimumRadius"))
	return rv
}
func (f SKFieldNode) SetMinimumRadius(value float32) {
	objc.Send[struct{}](f.ID, objc.Sel("setMinimumRadius:"), value)
}

// A mask that defines which categories this field belongs to.
//
// # Discussion
//
// Every field in a scene can be assigned to up to 32 different categories,
// each corresponding to a bit in the bit mask. The mask values are not
// predetermined by Sprite Kit. You define the mask values that are used in
// your game. The field node’s [SKFieldNode.CategoryBitMask] property is
// compared to a physics body’s [SKPhysicsBody.FieldBitMask] property using
// a logical AND operation. If the result is nonzero, the field is applied to
// the physics body.
//
// The default value is `0xFFFFFFFF` (all bits set).
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/categoryBitMask
func (f SKFieldNode) CategoryBitMask() uint32 {
	rv := objc.Send[uint32](f.ID, objc.Sel("categoryBitMask"))
	return rv
}
func (f SKFieldNode) SetCategoryBitMask(value uint32) {
	objc.Send[struct{}](f.ID, objc.Sel("setCategoryBitMask:"), value)
}

// The strength of the field.
//
// # Discussion
//
// The default value is `1.0`. There’s no specific unit of measurement for
// this property because the actual effect is dependent on the kind of field
// node being created. In practice, the best approach is to experiment with
// different field strengths and use them to determine the proper value
// empirically.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/strength
func (f SKFieldNode) Strength() float32 {
	rv := objc.Send[float32](f.ID, objc.Sel("strength"))
	return rv
}
func (f SKFieldNode) SetStrength(value float32) {
	objc.Send[struct{}](f.ID, objc.Sel("setStrength:"), value)
}

// The exponent that defines the rate of decay for the strength of the field
// as the distance increases between the node and the physics body being
// affected.
//
// # Discussion
//
// When the force of a field node is calculated, the force is multiplied by
// `pow(distance - minRadius, -falloff)`. The default falloff value is `0`,
// which indicates that no attenuation takes place. Some types of field nodes
// ignore the falloff parameter entirely, while others change the default
// value to something that is more logical for that type of field node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/falloff
func (f SKFieldNode) Falloff() float32 {
	rv := objc.Send[float32](f.ID, objc.Sel("falloff"))
	return rv
}
func (f SKFieldNode) SetFalloff(value float32) {
	objc.Send[struct{}](f.ID, objc.Sel("setFalloff:"), value)
}

// The rate at which a noise or turbulence field node changes.
//
// # Discussion
//
// A value of `0.0` means that the field should not animate at all.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/animationSpeed
func (f SKFieldNode) AnimationSpeed() float32 {
	rv := objc.Send[float32](f.ID, objc.Sel("animationSpeed"))
	return rv
}
func (f SKFieldNode) SetAnimationSpeed(value float32) {
	objc.Send[struct{}](f.ID, objc.Sel("setAnimationSpeed:"), value)
}

// The smoothness of the noise used to generate the forces.
//
// # Discussion
//
// This parameter should be a value between `0.0` and `1.0`, where `1.0`
// represents a uniform smoothness.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/smoothness
func (f SKFieldNode) Smoothness() float32 {
	rv := objc.Send[float32](f.ID, objc.Sel("smoothness"))
	return rv
}
func (f SKFieldNode) SetSmoothness(value float32) {
	objc.Send[struct{}](f.ID, objc.Sel("setSmoothness:"), value)
}

// The direction of a velocity field node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/direction
func (f SKFieldNode) Direction() Vector_float3 {
	rv := objc.Send[Vector_float3](f.ID, objc.Sel("direction"))
	return Vector_float3(rv)
}
func (f SKFieldNode) SetDirection(value Vector_float3) {
	objc.Send[struct{}](f.ID, objc.Sel("setDirection:"), value)
}

// A normal texture that specifies the velocities at different points in a
// velocity field node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKFieldNode/texture
func (f SKFieldNode) Texture() ISKTexture {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("texture"))
	return SKTextureFromID(objc.ID(rv))
}
func (f SKFieldNode) SetTexture(value ISKTexture) {
	objc.Send[struct{}](f.ID, objc.Sel("setTexture:"), value)
}
