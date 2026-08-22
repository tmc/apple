// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKConstraint] class.
var (
	_SKConstraintClass     SKConstraintClass
	_SKConstraintClassOnce sync.Once
)

func getSKConstraintClass() SKConstraintClass {
	_SKConstraintClassOnce.Do(func() {
		_SKConstraintClass = SKConstraintClass{class: objc.GetClass("SKConstraint")}
	})
	return _SKConstraintClass
}

// GetSKConstraintClass returns the class object for SKConstraint.
func GetSKConstraintClass() SKConstraintClass {
	return getSKConstraintClass()
}

type SKConstraintClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKConstraintClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKConstraintClass) Alloc() SKConstraint {
	rv := objc.Send[SKConstraint](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A specification for constraining a node’s position or rotation.
//
// # Overview
//
// An [SKConstraint] object describes a mathematical constraint on a node’s
// position or orientation. You attach constraints to nodes; after a scene
// processes any actions and physics interactions, it applies constraints
// attached to nodes in its node tree. Use constraints to ensure that certain
// relationships are true before the system renders a scene. For example, you
// might use a constraint to:
//
// - Change a node’s [SKNode.ZRotation] property so that it always points at
// another node or a position in the scene. - Keep a node within a specified
// distance of another node or a point in the scene. - Keep a node inside a
// specified rectangle. - Restrict the [SKNode.ZRotation] property of a node
// so that it has a more limited rotation range of motion.
//
// To use constraints, create an [NSArray] object that contains one or more
// constraint objects and assign the array to a node’s [SKNode.Constraints]
// property. When the system evaluates a scene, it executes the constraints on
// a node in the order they appear in the [SKNode.Constraints] array.
//
// You can’t change a constraint after you create it. However, you can
// selectively disable or enable a constraint by setting its
// [SKConstraint.Enabled] property. You can also use the
// [SKConstraint.ReferenceNode] property to convert positions to the
// referenced coordinate system before applying the constraint.
//
// # Controlling the Coordinate System Where a Constraint is Applied
//
//   - [SKConstraint.ReferenceNode]: The node whose coordinate system should be used to apply the constraint.
//   - [SKConstraint.SetReferenceNode]
//
// # Enabling and Disabling a Constraint
//
//   - [SKConstraint.Enabled]: A Boolean value that specifies whether the constraint is applied.
//   - [SKConstraint.SetEnabled]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKConstraint
//
// [NSArray]: https://developer.apple.com/documentation/Foundation/NSArray
type SKConstraint struct {
	objectivec.Object
}

// SKConstraintFromID constructs a [SKConstraint] from an objc.ID.
//
// A specification for constraining a node’s position or rotation.
func SKConstraintFromID(id objc.ID) SKConstraint {
	return SKConstraint{objectivec.Object{ID: id}}
}

// NOTE: SKConstraint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKConstraint] class.
//
// # Controlling the Coordinate System Where a Constraint is Applied
//
//   - [ISKConstraint.ReferenceNode]: The node whose coordinate system should be used to apply the constraint.
//   - [ISKConstraint.SetReferenceNode]
//
// # Enabling and Disabling a Constraint
//
//   - [ISKConstraint.Enabled]: A Boolean value that specifies whether the constraint is applied.
//   - [ISKConstraint.SetEnabled]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKConstraint
type ISKConstraint interface {
	objectivec.IObject

	// Topic: Controlling the Coordinate System Where a Constraint is Applied

	// The node whose coordinate system should be used to apply the constraint.
	ReferenceNode() ISKNode
	SetReferenceNode(value ISKNode)

	// Topic: Enabling and Disabling a Constraint

	// A Boolean value that specifies whether the constraint is applied.
	Enabled() bool
	SetEnabled(value bool)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c SKConstraint) Init() SKConstraint {
	rv := objc.Send[SKConstraint](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c SKConstraint) Autorelease() SKConstraint {
	rv := objc.Send[SKConstraint](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKConstraint creates a new SKConstraint instance.
func NewSKConstraint() SKConstraint {
	class := getSKConstraintClass()
	rv := objc.Send[SKConstraint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (c SKConstraint) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates a constraint that restricts both coordinates of a node’s
// position.
//
// xRange: The range to restrict the x-coordinate to.
//
// yRange: The range to restrict the y-coordinate to.
//
// # Return Value
//
// A new constraint.
//
// # Discussion
//
// Each time constraints are applied, the node’s [SKNode.Position] property
// is clamped so that both coordinates lie inside the specified ranges.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKConstraint/positionX(_:y:)
func (_SKConstraintClass SKConstraintClass) PositionXY(xRange ISKRange, yRange ISKRange) SKConstraint {
	rv := objc.Send[objc.ID](objc.ID(_SKConstraintClass.class), objc.Sel("positionX:Y:"), xRange, yRange)
	return SKConstraintFromID(rv)
}

// Creates a constraint that restricts the x-coordinate of a node’s
// position.
//
// range: The range to restrict the coordinate to.
//
// # Return Value
//
// A new constraint.
//
// # Discussion
//
// Each time constraints are applied, the x-coordinate of the node’s
// [SKNode.Position] property is clamped so that it lies inside the specified
// range.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKConstraint/positionX(_:)
func (_SKConstraintClass SKConstraintClass) PositionX(range_ ISKRange) SKConstraint {
	rv := objc.Send[objc.ID](objc.ID(_SKConstraintClass.class), objc.Sel("positionX:"), range_)
	return SKConstraintFromID(rv)
}

// Creates a constraint that restricts the y-coordinate of a node’s
// position.
//
// range: The range to restrict the coordinate to.
//
// # Return Value
//
// A new constraint.
//
// # Discussion
//
// Each time when constraints are applied, the y-coordinate of the node’s
// [SKNode.Position] property is clamped so that it lies inside the specified
// range.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKConstraint/positionY(_:)
func (_SKConstraintClass SKConstraintClass) PositionY(range_ ISKRange) SKConstraint {
	rv := objc.Send[objc.ID](objc.ID(_SKConstraintClass.class), objc.Sel("positionY:"), range_)
	return SKConstraintFromID(rv)
}

// Creates a constraint that forces a node to rotate to face another node.
//
// node: The node that should be used to orient the node that this constraint is
// attached to.
//
// radians: An offset that is added to the [SKNode.ZRotation] value after it is
// calculated.
//
// # Return Value
//
// A new constraint.
//
// # Discussion
//
// Each time when constraints are applied, a new angle is calculated so that a
// line projected at this angle would point at the other node’s origin. This
// angle is added to the values specified in the `radians` property to create
// a new range. Finally, the node’s [SKNode.ZRotation] value is clamped to
// fit inside this range.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKConstraint/orient(to:offset:)-1h1tw
func (_SKConstraintClass SKConstraintClass) OrientToNodeOffset(node ISKNode, radians ISKRange) SKConstraint {
	rv := objc.Send[objc.ID](objc.ID(_SKConstraintClass.class), objc.Sel("orientToNode:offset:"), node, radians)
	return SKConstraintFromID(rv)
}

// Creates a constraint that forces a node to rotate to face a fixed point.
//
// point: A point in the node’s parent’s coordinate system.
//
// radians: An offset that is added to the [SKNode.ZRotation] value after it is
// calculated.
//
// # Return Value
//
// A new constraint.
//
// # Discussion
//
// Each time when constraints are applied, a new angle is calculated so that a
// line projected at this angle would point at the target point. This angle is
// added to the values specified in the `radians` property to create a new
// range. Finally, the node’s [SKNode.ZRotation] value is clamped to fit
// inside this range.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKConstraint/orient(to:offset:)-9lq3h
func (_SKConstraintClass SKConstraintClass) OrientToPointOffset(point corefoundation.CGPoint, radians ISKRange) SKConstraint {
	rv := objc.Send[objc.ID](objc.ID(_SKConstraintClass.class), objc.Sel("orientToPoint:offset:"), point, radians)
	return SKConstraintFromID(rv)
}

// Creates a constraint that forces a node to rotate to face a point in
// another node’s coordinate system.
//
// point: A point in the `node` parameter’s coordinate system.
//
// node: The node whose coordinate system the point is specified in.
//
// radians: An offset that is added to the [SKNode.ZRotation] value after it is
// calculated.
//
// # Return Value
//
// A new constraint.
//
// # Discussion
//
// Each time when constraints are applied, a new angle is calculated so that a
// line projected at this angle would point at the target point. This angle is
// added to the values specified in the `radians` property to create a new
// range. Finally, the node’s [SKNode.ZRotation] value is clamped to fit
// inside this range.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKConstraint/orient(to:in:offset:)
func (_SKConstraintClass SKConstraintClass) OrientToPointInNodeOffset(point corefoundation.CGPoint, node ISKNode, radians ISKRange) SKConstraint {
	rv := objc.Send[objc.ID](objc.ID(_SKConstraintClass.class), objc.Sel("orientToPoint:inNode:offset:"), point, node, radians)
	return SKConstraintFromID(rv)
}

// Creates a constraint that limits the orientation of a node.
//
// zRange: A range value that specifies the minimum and maximum values of the node’s
// [SKNode.ZRotation] property.
//
// # Return Value
//
// A new constraint.
//
// # Discussion
//
// Each time when constraints are applied, the node’s [SKNode.ZRotation]
// property is clamped so that it is within the specified range.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKConstraint/zRotation(_:)
func (_SKConstraintClass SKConstraintClass) ZRotation(zRange ISKRange) SKConstraint {
	rv := objc.Send[objc.ID](objc.ID(_SKConstraintClass.class), objc.Sel("zRotation:"), zRange)
	return SKConstraintFromID(rv)
}

// Creates a constraint that keeps a node within a certain distance of another
// node.
//
// range: The range of allowed distances between the two nodes.
//
// node: The target node used to calculate the distance.
//
// # Return Value
//
// A new constraint.
//
// # Discussion
//
// Distance constraints constrain a node to a specified distance range of
// another node or a point and can be used for effects such a simulating
// flocking around a node, repulsive fields and trails. Supplying a distance
// constraint with a range with a lower limit, an upper limit or both results
// in very different behaviors:
//
// [Table data omitted]
//
// Each time when constraints are applied, a line is projected between the
// node’s position and the target node’s position. The distance between
// the two points is calculated, and if it lies outside the specified range,
// the node is pushed or pulled along this line until it lies within the
// range.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKConstraint/distance(_:to:)-6507j
func (_SKConstraintClass SKConstraintClass) DistanceToNode(range_ ISKRange, node ISKNode) SKConstraint {
	rv := objc.Send[objc.ID](objc.ID(_SKConstraintClass.class), objc.Sel("distance:toNode:"), range_, node)
	return SKConstraintFromID(rv)
}

// Creates a constraint that keeps a node within a certain distance of a
// point.
//
// range: The range of allowed distances between the node and the point.
//
// point: A point in the coordinate system of the node’s parent that is used to
// calculate the distance.
//
// # Return Value
//
// A new constraint.
//
// # Discussion
//
// Each time when constraints are applied, a line is projected between the
// node’s position and the target point. The distance between the two points
// is calculated, and if it lies outside the specified range, the node is
// pushed or pulled along this line until it lies within the range.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKConstraint/distance(_:to:)-7yk7n
func (_SKConstraintClass SKConstraintClass) DistanceToPoint(range_ ISKRange, point corefoundation.CGPoint) SKConstraint {
	rv := objc.Send[objc.ID](objc.ID(_SKConstraintClass.class), objc.Sel("distance:toPoint:"), range_, point)
	return SKConstraintFromID(rv)
}

// Creates a constraint that keeps a node within a certain distance of a point
// in another node’s coordinate system.
//
// range: The range of allowed distances.
//
// point: The point to use as the target point.
//
// node: The node whose coordinate system the point is specified in.
//
// # Return Value
//
// A new constraint.
//
// # Discussion
//
// Each time when constraints are applied, a line is projected between the
// node’s position and the target point. The distance between the two points
// is calculated, and if it lies outside the specified range, the node is
// pushed or pulled along this line until it lies within the range.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKConstraint/distance(_:to:in:)
func (_SKConstraintClass SKConstraintClass) DistanceToPointInNode(range_ ISKRange, point corefoundation.CGPoint, node ISKNode) SKConstraint {
	rv := objc.Send[objc.ID](objc.ID(_SKConstraintClass.class), objc.Sel("distance:toPoint:inNode:"), range_, point, node)
	return SKConstraintFromID(rv)
}

// The node whose coordinate system should be used to apply the constraint.
//
// # Discussion
//
// The default value is `nil`, meaning that the coordinate system of the
// node’s parent is used to apply the constraint. If another node is
// specified, all positions are converted into this node’s coordinate system
// before the constraint is applied.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKConstraint/referenceNode
func (c SKConstraint) ReferenceNode() ISKNode {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("referenceNode"))
	return SKNodeFromID(objc.ID(rv))
}
func (c SKConstraint) SetReferenceNode(value ISKNode) {
	objc.Send[struct{}](c.ID, objc.Sel("setReferenceNode:"), value)
}

// A Boolean value that specifies whether the constraint is applied.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKConstraint/enabled
func (c SKConstraint) Enabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("enabled"))
	return rv
}
func (c SKConstraint) SetEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setEnabled:"), value)
}
