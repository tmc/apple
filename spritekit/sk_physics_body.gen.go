// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKPhysicsBody] class.
var (
	_SKPhysicsBodyClass     SKPhysicsBodyClass
	_SKPhysicsBodyClassOnce sync.Once
)

func getSKPhysicsBodyClass() SKPhysicsBodyClass {
	_SKPhysicsBodyClassOnce.Do(func() {
		_SKPhysicsBodyClass = SKPhysicsBodyClass{class: objc.GetClass("SKPhysicsBody")}
	})
	return _SKPhysicsBodyClass
}

// GetSKPhysicsBodyClass returns the class object for SKPhysicsBody.
func GetSKPhysicsBodyClass() SKPhysicsBodyClass {
	return getSKPhysicsBodyClass()
}

type SKPhysicsBodyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKPhysicsBodyClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKPhysicsBodyClass) Alloc() SKPhysicsBody {
	rv := objc.Send[SKPhysicsBody](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// An object that adds physics simulation to a node.
//
// # Overview
//
// Assign a [SKPhysicsBody] object to the [SKNode.PhysicsBody] property of the
// [SKNode] object to add physics simulation to the node. When a scene
// processes a new frame, it performs physics calculations on physics bodies
// attached to nodes in the scene. These calculations include gravity,
// friction, and collisions with other bodies. You can also apply your own
// forces and impulses to a body. After the scene completes these
// calculations, it updates the positions and orientations of the node
// objects.
//
// SpriteKit supports two kinds of physics bodies, volume-based bodies and
// edge-based bodies. When you create a physics body, its kind, size, and
// shape are determined by the constructor method you call. An edge-based body
// does not have mass or volume and is unaffected by forces or impulses in the
// system. Edge-based bodies are used to represent volumeless boundaries or
// hollow spaces in your physics simulation. In contrast, volume-based bodies
// are used to represent objects with mass and volume. The
// [SKPhysicsBody.Dynamic] property controls whether a volume-based body is
// affected by gravity, friction, collisions with other objects, and forces or
// impulses you directly apply to it.
//
// The [SKPhysicsBody] class defines the physical characteristics for the body
// when it is simulated by the scene. For volume-based bodies, the most
// important property is the [SKPhysicsBody.Mass] property. A volume-based
// body is assumed to have a uniform density. You can either set the
// [SKPhysicsBody.Mass] property directly, or you can set the body’s
// [SKPhysicsBody.Density] property and let the physics body calculate its own
// mass. All values in Sprite Kit are specified using the International System
// of Units (SI units). The actual forces and mass values are not important so
// long as your game uses consistent values.
//
// When you design a game that uses physics, you define the different
// categories of physics objects that appear in the scene. You define up to 32
// different categories of physics bodies, and a body can be assigned to as
// many of these categories as you want. In addition to declaring its own
// categories, a physics body also declares which categories of bodies it
// interacts with. See [SKPhysicsBody]. You use a similar mechanism to declare
// which physics field nodes ([SKFieldNode]) can affect the physics body.
//
// For a volume-based body, you can dynamically control how the body is
// affected by forces or collisions. See [SKPhysicsBody].
//
// # Defining How Forces Affect a Physics Body
//
//   - [SKPhysicsBody.AffectedByGravity]: A Boolean value that indicates whether this physics body is affected by the physics world’s gravity.
//   - [SKPhysicsBody.SetAffectedByGravity]
//   - [SKPhysicsBody.AllowsRotation]: A Boolean value that indicates whether the physics body is affected by angular forces and impulses applied to it.
//   - [SKPhysicsBody.SetAllowsRotation]
//   - [SKPhysicsBody.IsDynamic]: A Boolean value that indicates whether the physics body is moved by the physics simulation.
//   - [SKPhysicsBody.SetDynamic]
//
// # Defining a Physics Body’s Physical Properties
//
//   - [SKPhysicsBody.Mass]: The mass of the body, in kilograms.
//   - [SKPhysicsBody.SetMass]
//   - [SKPhysicsBody.Density]: The density of the object, in kilograms per square meter.
//   - [SKPhysicsBody.SetDensity]
//   - [SKPhysicsBody.Area]: The area covered by the body.
//   - [SKPhysicsBody.Friction]: The roughness of the surface of the physics body.
//   - [SKPhysicsBody.SetFriction]
//   - [SKPhysicsBody.Restitution]: The bounciness of the physics body.
//   - [SKPhysicsBody.SetRestitution]
//   - [SKPhysicsBody.LinearDamping]: A property that reduces the body’s linear velocity.
//   - [SKPhysicsBody.SetLinearDamping]
//   - [SKPhysicsBody.AngularDamping]: A property that reduces the body’s rotational velocity.
//   - [SKPhysicsBody.SetAngularDamping]
//
// # Working with Collisions and Contacts
//
//   - [SKPhysicsBody.CategoryBitMask]: A mask that defines which categories this physics body belongs to.
//   - [SKPhysicsBody.SetCategoryBitMask]
//   - [SKPhysicsBody.CollisionBitMask]: A mask that defines which categories of physics bodies can collide with this physics body.
//   - [SKPhysicsBody.SetCollisionBitMask]
//   - [SKPhysicsBody.UsesPreciseCollisionDetection]: A Boolean value that determines whether the physics world uses an iterative collision detection algorithm.
//   - [SKPhysicsBody.SetUsesPreciseCollisionDetection]
//   - [SKPhysicsBody.ContactTestBitMask]: A mask that defines which categories of physics bodies cause intersection notifications with this physics body.
//   - [SKPhysicsBody.SetContactTestBitMask]
//   - [SKPhysicsBody.AllContactedBodies]: The physics bodies that this physics body is in contact with.
//
// # Applying Forces and Impulses to a Physics Body
//
//   - [SKPhysicsBody.ApplyForce]: Applies a force to the center of gravity of a physics body.
//   - [SKPhysicsBody.ApplyTorque]: Applies torque to an object.
//   - [SKPhysicsBody.ApplyForceAtPoint]: Applies a force to a specific point of a physics body.
//   - [SKPhysicsBody.ApplyImpulse]: Applies an impulse to the center of gravity of a physics body.
//   - [SKPhysicsBody.ApplyAngularImpulse]: Applies an impulse that imparts angular momentum to an object.
//   - [SKPhysicsBody.ApplyImpulseAtPoint]: Applies an impulse to a specific point of a physics body.
//
// # Inspecting a Physics Body’s Position and Velocity
//
//   - [SKPhysicsBody.Velocity]: The physics body’s velocity vector, measured in meters per second.
//   - [SKPhysicsBody.SetVelocity]
//   - [SKPhysicsBody.AngularVelocity]: The physics body’s angular speed.
//   - [SKPhysicsBody.SetAngularVelocity]
//   - [SKPhysicsBody.IsResting]: A Boolean property that indicates whether the object is at rest within the physics simulation.
//   - [SKPhysicsBody.SetResting]
//
// # Reading a Physics Body’s Node
//
//   - [SKPhysicsBody.Node]: The node that this body is connected to.
//
// # Determining Which Joints Are Connected to a Physics Body
//
//   - [SKPhysicsBody.Joints]: The joints connected to this physics body.
//
// # Interacting with Physics Fields
//
//   - [SKPhysicsBody.FieldBitMask]: A mask that defines which categories of physics fields can exert forces on this physics body.
//   - [SKPhysicsBody.SetFieldBitMask]
//   - [SKPhysicsBody.Charge]: The electrical charge of the physics body.
//   - [SKPhysicsBody.SetCharge]
//
// # Pinning a Physics Body to a Node’s Parent
//
//   - [SKPhysicsBody.Pinned]: A Boolean value that indicates whether the physics body’s node is pinned to its parent node.
//   - [SKPhysicsBody.SetPinned]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody
type SKPhysicsBody struct {
	objectivec.Object
}

// SKPhysicsBodyFromID constructs a [SKPhysicsBody] from an objc.ID.
//
// An object that adds physics simulation to a node.
func SKPhysicsBodyFromID(id objc.ID) SKPhysicsBody {
	return SKPhysicsBody{objectivec.Object{ID: id}}
}

// NOTE: SKPhysicsBody adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKPhysicsBody] class.
//
// # Defining How Forces Affect a Physics Body
//
//   - [ISKPhysicsBody.AffectedByGravity]: A Boolean value that indicates whether this physics body is affected by the physics world’s gravity.
//   - [ISKPhysicsBody.SetAffectedByGravity]
//   - [ISKPhysicsBody.AllowsRotation]: A Boolean value that indicates whether the physics body is affected by angular forces and impulses applied to it.
//   - [ISKPhysicsBody.SetAllowsRotation]
//   - [ISKPhysicsBody.IsDynamic]: A Boolean value that indicates whether the physics body is moved by the physics simulation.
//   - [ISKPhysicsBody.SetDynamic]
//
// # Defining a Physics Body’s Physical Properties
//
//   - [ISKPhysicsBody.Mass]: The mass of the body, in kilograms.
//   - [ISKPhysicsBody.SetMass]
//   - [ISKPhysicsBody.Density]: The density of the object, in kilograms per square meter.
//   - [ISKPhysicsBody.SetDensity]
//   - [ISKPhysicsBody.Area]: The area covered by the body.
//   - [ISKPhysicsBody.Friction]: The roughness of the surface of the physics body.
//   - [ISKPhysicsBody.SetFriction]
//   - [ISKPhysicsBody.Restitution]: The bounciness of the physics body.
//   - [ISKPhysicsBody.SetRestitution]
//   - [ISKPhysicsBody.LinearDamping]: A property that reduces the body’s linear velocity.
//   - [ISKPhysicsBody.SetLinearDamping]
//   - [ISKPhysicsBody.AngularDamping]: A property that reduces the body’s rotational velocity.
//   - [ISKPhysicsBody.SetAngularDamping]
//
// # Working with Collisions and Contacts
//
//   - [ISKPhysicsBody.CategoryBitMask]: A mask that defines which categories this physics body belongs to.
//   - [ISKPhysicsBody.SetCategoryBitMask]
//   - [ISKPhysicsBody.CollisionBitMask]: A mask that defines which categories of physics bodies can collide with this physics body.
//   - [ISKPhysicsBody.SetCollisionBitMask]
//   - [ISKPhysicsBody.UsesPreciseCollisionDetection]: A Boolean value that determines whether the physics world uses an iterative collision detection algorithm.
//   - [ISKPhysicsBody.SetUsesPreciseCollisionDetection]
//   - [ISKPhysicsBody.ContactTestBitMask]: A mask that defines which categories of physics bodies cause intersection notifications with this physics body.
//   - [ISKPhysicsBody.SetContactTestBitMask]
//   - [ISKPhysicsBody.AllContactedBodies]: The physics bodies that this physics body is in contact with.
//
// # Applying Forces and Impulses to a Physics Body
//
//   - [ISKPhysicsBody.ApplyForce]: Applies a force to the center of gravity of a physics body.
//   - [ISKPhysicsBody.ApplyTorque]: Applies torque to an object.
//   - [ISKPhysicsBody.ApplyForceAtPoint]: Applies a force to a specific point of a physics body.
//   - [ISKPhysicsBody.ApplyImpulse]: Applies an impulse to the center of gravity of a physics body.
//   - [ISKPhysicsBody.ApplyAngularImpulse]: Applies an impulse that imparts angular momentum to an object.
//   - [ISKPhysicsBody.ApplyImpulseAtPoint]: Applies an impulse to a specific point of a physics body.
//
// # Inspecting a Physics Body’s Position and Velocity
//
//   - [ISKPhysicsBody.Velocity]: The physics body’s velocity vector, measured in meters per second.
//   - [ISKPhysicsBody.SetVelocity]
//   - [ISKPhysicsBody.AngularVelocity]: The physics body’s angular speed.
//   - [ISKPhysicsBody.SetAngularVelocity]
//   - [ISKPhysicsBody.IsResting]: A Boolean property that indicates whether the object is at rest within the physics simulation.
//   - [ISKPhysicsBody.SetResting]
//
// # Reading a Physics Body’s Node
//
//   - [ISKPhysicsBody.Node]: The node that this body is connected to.
//
// # Determining Which Joints Are Connected to a Physics Body
//
//   - [ISKPhysicsBody.Joints]: The joints connected to this physics body.
//
// # Interacting with Physics Fields
//
//   - [ISKPhysicsBody.FieldBitMask]: A mask that defines which categories of physics fields can exert forces on this physics body.
//   - [ISKPhysicsBody.SetFieldBitMask]
//   - [ISKPhysicsBody.Charge]: The electrical charge of the physics body.
//   - [ISKPhysicsBody.SetCharge]
//
// # Pinning a Physics Body to a Node’s Parent
//
//   - [ISKPhysicsBody.Pinned]: A Boolean value that indicates whether the physics body’s node is pinned to its parent node.
//   - [ISKPhysicsBody.SetPinned]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody
type ISKPhysicsBody interface {
	objectivec.IObject

	// Topic: Defining How Forces Affect a Physics Body

	// A Boolean value that indicates whether this physics body is affected by the physics world’s gravity.
	AffectedByGravity() bool
	SetAffectedByGravity(value bool)
	// A Boolean value that indicates whether the physics body is affected by angular forces and impulses applied to it.
	AllowsRotation() bool
	SetAllowsRotation(value bool)
	// A Boolean value that indicates whether the physics body is moved by the physics simulation.
	IsDynamic() bool
	SetDynamic(value bool)

	// Topic: Defining a Physics Body’s Physical Properties

	// The mass of the body, in kilograms.
	Mass() float64
	SetMass(value float64)
	// The density of the object, in kilograms per square meter.
	Density() float64
	SetDensity(value float64)
	// The area covered by the body.
	Area() float64
	// The roughness of the surface of the physics body.
	Friction() float64
	SetFriction(value float64)
	// The bounciness of the physics body.
	Restitution() float64
	SetRestitution(value float64)
	// A property that reduces the body’s linear velocity.
	LinearDamping() float64
	SetLinearDamping(value float64)
	// A property that reduces the body’s rotational velocity.
	AngularDamping() float64
	SetAngularDamping(value float64)

	// Topic: Working with Collisions and Contacts

	// A mask that defines which categories this physics body belongs to.
	CategoryBitMask() uint32
	SetCategoryBitMask(value uint32)
	// A mask that defines which categories of physics bodies can collide with this physics body.
	CollisionBitMask() uint32
	SetCollisionBitMask(value uint32)
	// A Boolean value that determines whether the physics world uses an iterative collision detection algorithm.
	UsesPreciseCollisionDetection() bool
	SetUsesPreciseCollisionDetection(value bool)
	// A mask that defines which categories of physics bodies cause intersection notifications with this physics body.
	ContactTestBitMask() uint32
	SetContactTestBitMask(value uint32)
	// The physics bodies that this physics body is in contact with.
	AllContactedBodies() []SKPhysicsBody

	// Topic: Applying Forces and Impulses to a Physics Body

	// Applies a force to the center of gravity of a physics body.
	ApplyForce(force corefoundation.CGVector)
	// Applies torque to an object.
	ApplyTorque(torque float64)
	// Applies a force to a specific point of a physics body.
	ApplyForceAtPoint(force corefoundation.CGVector, point corefoundation.CGPoint)
	// Applies an impulse to the center of gravity of a physics body.
	ApplyImpulse(impulse corefoundation.CGVector)
	// Applies an impulse that imparts angular momentum to an object.
	ApplyAngularImpulse(impulse float64)
	// Applies an impulse to a specific point of a physics body.
	ApplyImpulseAtPoint(impulse corefoundation.CGVector, point corefoundation.CGPoint)

	// Topic: Inspecting a Physics Body’s Position and Velocity

	// The physics body’s velocity vector, measured in meters per second.
	Velocity() corefoundation.CGVector
	SetVelocity(value corefoundation.CGVector)
	// The physics body’s angular speed.
	AngularVelocity() float64
	SetAngularVelocity(value float64)
	// A Boolean property that indicates whether the object is at rest within the physics simulation.
	IsResting() bool
	SetResting(value bool)

	// Topic: Reading a Physics Body’s Node

	// The node that this body is connected to.
	Node() ISKNode

	// Topic: Determining Which Joints Are Connected to a Physics Body

	// The joints connected to this physics body.
	Joints() []SKPhysicsJoint

	// Topic: Interacting with Physics Fields

	// A mask that defines which categories of physics fields can exert forces on this physics body.
	FieldBitMask() uint32
	SetFieldBitMask(value uint32)
	// The electrical charge of the physics body.
	Charge() float64
	SetCharge(value float64)

	// Topic: Pinning a Physics Body to a Node’s Parent

	// A Boolean value that indicates whether the physics body’s node is pinned to its parent node.
	Pinned() bool
	SetPinned(value bool)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (p SKPhysicsBody) Init() SKPhysicsBody {
	rv := objc.Send[SKPhysicsBody](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p SKPhysicsBody) Autorelease() SKPhysicsBody {
	rv := objc.Send[SKPhysicsBody](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKPhysicsBody creates a new SKPhysicsBody instance.
func NewSKPhysicsBody() SKPhysicsBody {
	class := getSKPhysicsBodyClass()
	rv := objc.Send[SKPhysicsBody](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a physics body that’s shaped like a union of the argument physics
// bodies.
//
// bodies: An array of [SKPhysicsBody] objects. The objects must be volume-based
// physics bodies. (You may not use a compound body created using this method
// in the array.)
//
// # Return Value
//
// A new compound-physics body.
//
// # Discussion
//
// The shapes of the physics bodies passed into this method are used to create
// a new physics body whose covered area is the union of the areas of its
// children. These areas do not need to be contiguous. If there is space
// between two parts, other bodies may be able to pass between these parts.
// However, the physics body is treated as a single connected body, meaning
// that a force or impulse applied to the body affects all of the pieces as if
// they are held together with an indestructible frame.
//
// The properties on the children, such as mass or friction, are ignored. Only
// the shapes of the child bodies are used.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(bodies:)
func NewPhysicsBodyWithBodies(bodies []SKPhysicsBody) SKPhysicsBody {
	rv := objc.Send[objc.ID](objc.ID(getSKPhysicsBodyClass().class), objc.Sel("bodyWithBodies:"), objectivec.IObjectSliceToNSArray(bodies))
	return SKPhysicsBodyFromID(rv)
}

// Creates a circular physics body centered on the owning node’s origin.
//
// r: The radius of the circle.
//
// # Return Value
//
// A new volume-based physics body.
//
// # Discussion
//
// The following code shows the code that creates the physics body for a
// spherical or circular object. Because the physics body is attached to a
// sprite object, it usually needs volume. In this case, the sprite image is
// assumed to closely approximate a circle centered on the anchor point, so
// the radius of the circle is calculated and used to create the physics body.
//
// Listing 1. A physics body for a circular sprite
//
// If the physics body were significantly smaller than the sprite’s image,
// the data used to create the physics body might need to be provided by some
// other source, such as a property list.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(circleOfRadius:)
func NewPhysicsBodyWithCircleOfRadius(r float64) SKPhysicsBody {
	rv := objc.Send[objc.ID](objc.ID(getSKPhysicsBodyClass().class), objc.Sel("bodyWithCircleOfRadius:"), r)
	return SKPhysicsBodyFromID(rv)
}

// Creates a circular physics body centered on an arbitrary point.
//
// r: The radius of the circle.
//
// center: The origin of the circle in the owning node’s coordinate system.
//
// # Return Value
//
// A new volume-based physics body.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(circleOfRadius:center:)
func NewPhysicsBodyWithCircleOfRadiusCenter(r float64, center corefoundation.CGPoint) SKPhysicsBody {
	rv := objc.Send[objc.ID](objc.ID(getSKPhysicsBodyClass().class), objc.Sel("bodyWithCircleOfRadius:center:"), r, center)
	return SKPhysicsBodyFromID(rv)
}

// Creates an edge chain from a path.
//
// path: A Core Graphics path. The points are specified relative to the owning
// node’s origin. The path must not intersect itself.
//
// # Return Value
//
// A new edge-based physics body.
//
// # Discussion
//
// An edge has no volume or mass and is always treated as if the
// [SKPhysicsBody.Dynamic] property is equal to false. Edges may only collide
// with volume-based physics bodies.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(edgeChainFrom:)
func NewPhysicsBodyWithEdgeChainFromPath(path coregraphics.CGPathRef) SKPhysicsBody {
	rv := objc.Send[objc.ID](objc.ID(getSKPhysicsBodyClass().class), objc.Sel("bodyWithEdgeChainFromPath:"), path)
	return SKPhysicsBodyFromID(rv)
}

// Creates an edge between two points.
//
// p1: The starting point for the edge, relative to the owning node’s origin.
//
// p2: The ending point for the edge, relative to the owning node’s origin.
//
// # Return Value
//
// A new edge-based physics body.
//
// # Discussion
//
// An edge has no volume or mass and is always treated as if the
// [SKPhysicsBody.Dynamic] property is equal to false. Edges may only collide
// with volume-based physics bodies.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(edgeFrom:to:)
func NewPhysicsBodyWithEdgeFromPointToPoint(p1 corefoundation.CGPoint, p2 corefoundation.CGPoint) SKPhysicsBody {
	rv := objc.Send[objc.ID](objc.ID(getSKPhysicsBodyClass().class), objc.Sel("bodyWithEdgeFromPoint:toPoint:"), p1, p2)
	return SKPhysicsBodyFromID(rv)
}

// Creates an edge loop from a path.
//
// path: A Core Graphics path. The points are specified relative to the owning
// node’s origin. The path must not intersect itself.
//
// # Return Value
//
// A new edge-based physics body.
//
// # Discussion
//
// If the path is not already closed, a loop is automatically created by
// joining the last point to the first.
//
// An edge has no volume or mass and is always treated as if the
// [SKPhysicsBody.Dynamic] property is equal to false. Edges may only collide
// with volume-based physics bodies.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(edgeLoopFrom:)-5grxu
func NewPhysicsBodyWithEdgeLoopFromPath(path coregraphics.CGPathRef) SKPhysicsBody {
	rv := objc.Send[objc.ID](objc.ID(getSKPhysicsBodyClass().class), objc.Sel("bodyWithEdgeLoopFromPath:"), path)
	return SKPhysicsBodyFromID(rv)
}

// Creates an edge loop from a rectangle.
//
// rect: The rectangle that defines the edges. The rectangle is specified relative
// to the owning node’s origin.
//
// # Return Value
//
// A new edge-based physics body.
//
// # Discussion
//
// An edge has no volume or mass and is always treated as if the
// [SKPhysicsBody.Dynamic] property is equal to false. Edges may only collide
// with volume-based physics bodies.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(edgeLoopFrom:)-8sqfy
func NewPhysicsBodyWithEdgeLoopFromRect(rect corefoundation.CGRect) SKPhysicsBody {
	rv := objc.Send[objc.ID](objc.ID(getSKPhysicsBodyClass().class), objc.Sel("bodyWithEdgeLoopFromRect:"), rect)
	return SKPhysicsBodyFromID(rv)
}

// Creates a polygonal physics body.
//
// path: A convex polygonal path with counterclockwise winding and no self
// intersections. The points are specified relative to the owning node’s
// origin.
//
// # Return Value
//
// A new volume-based physics body.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(polygonFrom:)
func NewPhysicsBodyWithPolygonFromPath(path coregraphics.CGPathRef) SKPhysicsBody {
	rv := objc.Send[objc.ID](objc.ID(getSKPhysicsBodyClass().class), objc.Sel("bodyWithPolygonFromPath:"), path)
	return SKPhysicsBodyFromID(rv)
}

// Creates a rectangular physics body centered on the owning node’s origin.
//
// s: The size of the rectangle.
//
// # Return Value
//
// A new volume-based physics body.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(rectangleOf:)
func NewPhysicsBodyWithRectangleOfSize(s corefoundation.CGSize) SKPhysicsBody {
	rv := objc.Send[objc.ID](objc.ID(getSKPhysicsBodyClass().class), objc.Sel("bodyWithRectangleOfSize:"), s)
	return SKPhysicsBodyFromID(rv)
}

// Creates a rectangular physics body centered on an arbitrary point.
//
// s: The size of the rectangle.
//
// center: The center of the square in the owning node’s coordinate system.
//
// # Return Value
//
// A new volume-based physics body.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(rectangleOf:center:)
func NewPhysicsBodyWithRectangleOfSizeCenter(s corefoundation.CGSize, center corefoundation.CGPoint) SKPhysicsBody {
	rv := objc.Send[objc.ID](objc.ID(getSKPhysicsBodyClass().class), objc.Sel("bodyWithRectangleOfSize:center:"), s, center)
	return SKPhysicsBodyFromID(rv)
}

// Creates a physics body from the contents of a texture, capturing only the
// texels that exceed a specified transparency value.
//
// texture: The texture to analyze.
//
// alphaThreshold: The minimum alpha value for texels that should be part of the new physics
// body.
//
// size: The size of the physics body to return.
//
// # Return Value
//
// A new volume-based physics body.
//
// # Discussion
//
// Use this method when your sprite has a shape that you want replicated in
// its physics body. The texture is scaled to the new size and then analyzed.
// A new physics body is created that includes all of the texels in the
// texture whose alpha values equal or exceed the `alphaThreshold` parameter.
// The shape of this body attempts to strike a good balance between
// performance and accuracy. For example, fine details may be ignored if
// keeping them would cause a significant performance penalty.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(texture:alphaThreshold:size:)
func NewPhysicsBodyWithTextureAlphaThresholdSize(texture ISKTexture, alphaThreshold float32, size corefoundation.CGSize) SKPhysicsBody {
	rv := objc.Send[objc.ID](objc.ID(getSKPhysicsBodyClass().class), objc.Sel("bodyWithTexture:alphaThreshold:size:"), texture, alphaThreshold, size)
	return SKPhysicsBodyFromID(rv)
}

// Creates a physics body from the contents of a texture.
//
// texture: The texture to convert into a physics body.
//
// size: The size of the physics body to return.
//
// # Return Value
//
// A new volume-based physics body.
//
// # Discussion
//
// Use this method when your sprite has a shape that you want replicated in
// its physics body. The texture is scaled to the new size and then analyzed.
// A new physics body is created that includes all of the texels in the
// texture that have a nonzero alpha value. The shape of this body attempts to
// strike a good balance between performance and accuracy. For example, fine
// details may be ignored if keeping them would cause a significant
// performance penalty.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(texture:size:)
func NewPhysicsBodyWithTextureSize(texture ISKTexture, size corefoundation.CGSize) SKPhysicsBody {
	rv := objc.Send[objc.ID](objc.ID(getSKPhysicsBodyClass().class), objc.Sel("bodyWithTexture:size:"), texture, size)
	return SKPhysicsBodyFromID(rv)
}

// The physics bodies that this physics body is in contact with.
//
// # Return Value
//
// An array of [SKPhysicsBody] objects that this body is in contact with.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/allContactedBodies()
func (p SKPhysicsBody) AllContactedBodies() []SKPhysicsBody {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("allContactedBodies"))
	return objc.ConvertSlice(rv, func(id objc.ID) SKPhysicsBody {
		return SKPhysicsBodyFromID(id)
	})
}

// Applies a force to the center of gravity of a physics body.
//
// force: A vector that describes how much force was applied in each dimension. The
// force is measured in Newtons.
//
// # Discussion
//
// This method accelerates the body without imparting any angular acceleration
// to it. The acceleration is applied for a single simulation step (one
// frame).
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/applyForce(_:)
func (p SKPhysicsBody) ApplyForce(force corefoundation.CGVector) {
	objc.Send[objc.ID](p.ID, objc.Sel("applyForce:"), force)
}

// Applies torque to an object.
//
// torque: The amount of torque, in Newton-meters.
//
// # Discussion
//
// This method generates an angular acceleration on the body without causing
// any linear acceleration. The force is applied for a single simulation step
// (one frame).
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/applyTorque(_:)
func (p SKPhysicsBody) ApplyTorque(torque float64) {
	objc.Send[objc.ID](p.ID, objc.Sel("applyTorque:"), torque)
}

// Applies a force to a specific point of a physics body.
//
// force: A vector that describes how much force was applied in each dimension. The
// force is measured in Newtons.
//
// point: A point in scene coordinates that defines where the force was applied to
// the physics body.
//
// # Discussion
//
// Because the force is applied to a specific point on the body, it may impart
// both linear acceleration and angular acceleration. The force is applied for
// a single simulation step (one frame).
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/applyForce(_:at:)
func (p SKPhysicsBody) ApplyForceAtPoint(force corefoundation.CGVector, point corefoundation.CGPoint) {
	objc.Send[objc.ID](p.ID, objc.Sel("applyForce:atPoint:"), force, point)
}

// Applies an impulse to the center of gravity of a physics body.
//
// impulse: A vector that describes how much momentum was imparted in each dimension.
// The impulse is measured in Newton-seconds.
//
// # Discussion
//
// This method affects the body’s linear velocity without changing the
// body’s angular velocity.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/applyImpulse(_:)
func (p SKPhysicsBody) ApplyImpulse(impulse corefoundation.CGVector) {
	objc.Send[objc.ID](p.ID, objc.Sel("applyImpulse:"), impulse)
}

// Applies an impulse that imparts angular momentum to an object.
//
// impulse: The magnitude of the impulse. The impulse is measured in Newton-seconds.
//
// # Discussion
//
// This method affects the body’s angular velocity without changing the
// body’s linear velocity.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/applyAngularImpulse(_:)
func (p SKPhysicsBody) ApplyAngularImpulse(impulse float64) {
	objc.Send[objc.ID](p.ID, objc.Sel("applyAngularImpulse:"), impulse)
}

// Applies an impulse to a specific point of a physics body.
//
// impulse: A vector that describes how much momentum to impart to the body. The
// impulse is measured in Newton-seconds.
//
// point: A point in scene coordinates that defines where the impulse was applied to
// the physics body.
//
// # Discussion
//
// Because this impulse is applied to a specific point on the object, it may
// change both the body’s velocity and angular velocity.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/applyImpulse(_:at:)
func (p SKPhysicsBody) ApplyImpulseAtPoint(impulse corefoundation.CGVector, point corefoundation.CGPoint) {
	objc.Send[objc.ID](p.ID, objc.Sel("applyImpulse:atPoint:"), impulse, point)
}
func (p SKPhysicsBody) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A Boolean value that indicates whether this physics body is affected by the
// physics world’s gravity.
//
// # Discussion
//
// The physics world’s [SKPhysicsWorld.Gravity] property defines the
// gravitational forces applied to volume-based bodies in the scene.
//
// The default value is `true`. This property is ignored on edge-based bodies,
// which are always unaffected by gravity.
//
// Physics bodies with `affectedByGravity` set to `false` are still affected
// by the gravity fields created by
// [SKFieldNodeClass.LinearGravityFieldWithVector] and
// [SKFieldNodeClass.RadialGravityField].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/affectedByGravity
func (p SKPhysicsBody) AffectedByGravity() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("affectedByGravity"))
	return rv
}
func (p SKPhysicsBody) SetAffectedByGravity(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAffectedByGravity:"), value)
}

// A Boolean value that indicates whether the physics body is affected by
// angular forces and impulses applied to it.
//
// # Discussion
//
// The default value is true. This property is ignored on edge-based bodies,
// which are unaffected by forces in the system.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/allowsRotation
func (p SKPhysicsBody) AllowsRotation() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("allowsRotation"))
	return rv
}
func (p SKPhysicsBody) SetAllowsRotation(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAllowsRotation:"), value)
}

// A Boolean value that indicates whether the physics body is moved by the
// physics simulation.
//
// # Discussion
//
// The default value is true. If the value is false, the physics body ignores
// all forces and impulses applied to it. This property is ignored on
// edge-based bodies; they are automatically static.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/isDynamic
func (p SKPhysicsBody) IsDynamic() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isDynamic"))
	return rv
}
func (p SKPhysicsBody) SetDynamic(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setDynamic:"), value)
}

// The mass of the body, in kilograms.
//
// # Discussion
//
// The actual unit is arbitrary as long as relative masses of objects are
// consistent throughout the game. The mass of the body affects its momentum
// as well as how forces are applied to the object.
//
// The [SKPhysicsBody.Mass] and [SKPhysicsBody.Density] properties are
// interrelated. When you change the value of either property, the other
// property’s value is automatically recalculated to be consistent. The
// default value is based on the size of the physics body and the body’s
// default density.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/mass
func (p SKPhysicsBody) Mass() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("mass"))
	return rv
}
func (p SKPhysicsBody) SetMass(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setMass:"), value)
}

// The density of the object, in kilograms per square meter.
//
// # Discussion
//
// The actual unit is arbitrary as long as relative masses of objects are
// consistent throughout the game.
//
// The [SKPhysicsBody.Mass] and [SKPhysicsBody.Density] properties are
// interrelated. When you change the value of either property, the other
// property’s value is automatically recalculated to be consistent.
//
// The default value is `1.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/density
func (p SKPhysicsBody) Density() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("density"))
	return rv
}
func (p SKPhysicsBody) SetDensity(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setDensity:"), value)
}

// The area covered by the body.
//
// # Discussion
//
// This property is used in conjunction with the [SKPhysicsBody.Density]
// property to calculate the body’s mass.
//
// The value returned for the area is measured in meters: if you need to
// convert it into points — as used by SpriteKit — multiply the values by
// 150². The following listing shows how to calculate the area of a box which
// is ten points square.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/area
func (p SKPhysicsBody) Area() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("area"))
	return rv
}

// The roughness of the surface of the physics body.
//
// # Discussion
//
// This property is used to apply a frictional force to physics bodies in
// contact with this physics body. The property must be a value between `0.0`
// and `1.0`. The default value is `0.2`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/friction
func (p SKPhysicsBody) Friction() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("friction"))
	return rv
}
func (p SKPhysicsBody) SetFriction(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setFriction:"), value)
}

// The bounciness of the physics body.
//
// # Discussion
//
// This property is used to determine how much energy the physics body loses
// when it bounces off another object. The property must be a value between
// `0.0` and `1.0`. The default value is `0.2`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/restitution
func (p SKPhysicsBody) Restitution() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("restitution"))
	return rv
}
func (p SKPhysicsBody) SetRestitution(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setRestitution:"), value)
}

// A property that reduces the body’s linear velocity.
//
// # Discussion
//
// This property is used to simulate fluid or air friction forces on the body.
// The property must be a value between `0.0` and `1.0`. The default value is
// `0.1`. If the value is `0.0`, no linear damping is applied to the object.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/linearDamping
func (p SKPhysicsBody) LinearDamping() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("linearDamping"))
	return rv
}
func (p SKPhysicsBody) SetLinearDamping(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setLinearDamping:"), value)
}

// A property that reduces the body’s rotational velocity.
//
// # Discussion
//
// This property is used to simulate fluid or air friction forces on the body.
// The property must be a value between `0.0` and `1.0`. The default value is
// `0.1`. If the value is `0.0`, no angular damping is applied to the object.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/angularDamping
func (p SKPhysicsBody) AngularDamping() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("angularDamping"))
	return rv
}
func (p SKPhysicsBody) SetAngularDamping(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setAngularDamping:"), value)
}

// A mask that defines which categories this physics body belongs to.
//
// # Discussion
//
// Every physics body in a scene can be assigned to up to 32 different
// categories, each corresponding to a bit in the bit mask. You define the
// mask values used in your game. In conjunction with the
// [SKPhysicsBody.CollisionBitMask] and [SKPhysicsBody.ContactTestBitMask]
// properties, you define which physics bodies interact with each other and
// when your game is notified of these interactions.
//
// The default value is `0xFFFFFFFF` (all bits set).
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/categoryBitMask
func (p SKPhysicsBody) CategoryBitMask() uint32 {
	rv := objc.Send[uint32](p.ID, objc.Sel("categoryBitMask"))
	return rv
}
func (p SKPhysicsBody) SetCategoryBitMask(value uint32) {
	objc.Send[struct{}](p.ID, objc.Sel("setCategoryBitMask:"), value)
}

// A mask that defines which categories of physics bodies can collide with
// this physics body.
//
// # Discussion
//
// When two physics bodies contact each other, a collision may occur. This
// body’s collision mask is compared to the other body’s category mask by
// performing a logical AND operation. If the result is a nonzero value, this
// body is affected by the collision. Each body independently chooses whether
// it wants to be affected by the other body. For example, you might use this
// to avoid collision calculations that would make negligible changes to a
// body’s velocity.
//
// The default value is `0xFFFFFFFF` (all bits set).
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/collisionBitMask
func (p SKPhysicsBody) CollisionBitMask() uint32 {
	rv := objc.Send[uint32](p.ID, objc.Sel("collisionBitMask"))
	return rv
}
func (p SKPhysicsBody) SetCollisionBitMask(value uint32) {
	objc.Send[struct{}](p.ID, objc.Sel("setCollisionBitMask:"), value)
}

// A Boolean value that determines whether the physics world uses an iterative
// collision detection algorithm.
//
// # Discussion
//
// When SpriteKit performs collision detection, it first determines the
// locations of all of the physics bodies in the scene. Then it determines
// whether collisions or contacts occurred. This computational method is fast,
// but can sometimes result in missed collisions. A small body might move so
// fast that it completely passes through another physics body without ever
// having a frame of animation where the two touch each other.
//
// If you have physics bodies that must collide, you can hint to SpriteKit to
// use a more precise collision model to check for interactions. This model is
// more expensive, so it should be used sparingly. When either body uses
// precise collisions, multiple contact iterations are evaluated to ensure
// that all contacts are detected.
//
// The default value is false. If two bodies in a collision do not perform
// precise collision detection, and one passes completely through the other in
// a single frame, no collision is detected. If this property is set to true
// on either body, the simulation performs a more precise and more expensive
// calculation to detect these collisions. This property should be set to true
// on small, fast moving bodies.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/usesPreciseCollisionDetection
func (p SKPhysicsBody) UsesPreciseCollisionDetection() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("usesPreciseCollisionDetection"))
	return rv
}
func (p SKPhysicsBody) SetUsesPreciseCollisionDetection(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setUsesPreciseCollisionDetection:"), value)
}

// A mask that defines which categories of physics bodies cause intersection
// notifications with this physics body.
//
// # Discussion
//
// When two bodies share the same space, each body’s category mask is tested
// against the other body’s contact mask by performing a logical AND
// operation. If either comparison results in a nonzero value, an
// [SKPhysicsContact] object is created and passed to the physics world’s
// delegate. For best performance, only set bits in the contacts mask for
// interactions you are interested in.
//
// The default value is `0x00000000` (all bits cleared).
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/contactTestBitMask
func (p SKPhysicsBody) ContactTestBitMask() uint32 {
	rv := objc.Send[uint32](p.ID, objc.Sel("contactTestBitMask"))
	return rv
}
func (p SKPhysicsBody) SetContactTestBitMask(value uint32) {
	objc.Send[struct{}](p.ID, objc.Sel("setContactTestBitMask:"), value)
}

// The physics body’s velocity vector, measured in meters per second.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/velocity
func (p SKPhysicsBody) Velocity() corefoundation.CGVector {
	rv := objc.Send[corefoundation.CGVector](p.ID, objc.Sel("velocity"))
	return corefoundation.CGVector(rv)
}
func (p SKPhysicsBody) SetVelocity(value corefoundation.CGVector) {
	objc.Send[struct{}](p.ID, objc.Sel("setVelocity:"), value)
}

// The physics body’s angular speed.
//
// # Discussion
//
// The angular velocity is a pseudo vector around an axis vector of
// `(0.0,0.0,1.0)` measured in radians per second.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/angularVelocity
func (p SKPhysicsBody) AngularVelocity() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("angularVelocity"))
	return rv
}
func (p SKPhysicsBody) SetAngularVelocity(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setAngularVelocity:"), value)
}

// A Boolean property that indicates whether the object is at rest within the
// physics simulation.
//
// # Discussion
//
// This property is automatically set to true by the physics simulation when
// it determines that the body is at rest. This means that the body is at rest
// on another body in the system. Resting bodies do not participate in the
// physics simulation until an impulse is applied to the object or another
// object collides with it. This improves the performance of the physics
// simulation. If all bodies in the world are resting, the entire simulation
// is at rest, reducing the number of calculations that are performed by the
// physics world.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/isResting
func (p SKPhysicsBody) IsResting() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isResting"))
	return rv
}
func (p SKPhysicsBody) SetResting(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setResting:"), value)
}

// The node that this body is connected to.
//
// # Discussion
//
// You associate the body with a node by assigning it to the
// [SKNode.PhysicsBody] property of the [SKNode] object. If the body is not
// associated with a node, the value is `nil`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/node
func (p SKPhysicsBody) Node() ISKNode {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("node"))
	return SKNodeFromID(objc.ID(rv))
}

// The joints connected to this physics body.
//
// # Discussion
//
// This property holds an array of [SKPhysicsJoint] objects that define all
// joints added to the scene’s physics world that are connected to this
// physics body.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/joints
func (p SKPhysicsBody) Joints() []SKPhysicsJoint {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("joints"))
	return objc.ConvertSlice(rv, func(id objc.ID) SKPhysicsJoint {
		return SKPhysicsJointFromID(id)
	})
}

// A mask that defines which categories of physics fields can exert forces on
// this physics body.
//
// # Discussion
//
// When a physics body is inside the region of an [SKFieldNode] object, that
// field node’s [SKFieldNode.CategoryBitMask] property is compared to this
// physics body’s [SKPhysicsBody.FieldBitMask] property by performing a
// logical AND operation. If the result is a nonzero value, the field node’s
// effect is applied to the physics body.
//
// The default value is `0xFFFFFFFF` (all bits set).
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/fieldBitMask
func (p SKPhysicsBody) FieldBitMask() uint32 {
	rv := objc.Send[uint32](p.ID, objc.Sel("fieldBitMask"))
	return rv
}
func (p SKPhysicsBody) SetFieldBitMask(value uint32) {
	objc.Send[struct{}](p.ID, objc.Sel("setFieldBitMask:"), value)
}

// The electrical charge of the physics body.
//
// # Discussion
//
// The electrical charge is used by electromagnetic fields to calculate
// electromagnetic force effects on the physics body. See [SKFieldNode].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/charge
func (p SKPhysicsBody) Charge() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("charge"))
	return rv
}
func (p SKPhysicsBody) SetCharge(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setCharge:"), value)
}

// A Boolean value that indicates whether the physics body’s node is pinned
// to its parent node.
//
// # Discussion
//
// The default value is false. If true, the node’s position is fixed
// relative to its parent. The node’s position cannot be changed by actions
// or physics forces. The node can freely rotate around its position in
// response to collisions or other forces. If the parent node has a physics
// body, the two physics bodies are treated as if they are connected with a
// pin joint.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/pinned
func (p SKPhysicsBody) Pinned() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("pinned"))
	return rv
}
func (p SKPhysicsBody) SetPinned(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setPinned:"), value)
}
