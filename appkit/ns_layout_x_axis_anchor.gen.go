// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSLayoutXAxisAnchor] class.
var (
	_NSLayoutXAxisAnchorClass     NSLayoutXAxisAnchorClass
	_NSLayoutXAxisAnchorClassOnce sync.Once
)

func getNSLayoutXAxisAnchorClass() NSLayoutXAxisAnchorClass {
	_NSLayoutXAxisAnchorClassOnce.Do(func() {
		_NSLayoutXAxisAnchorClass = NSLayoutXAxisAnchorClass{class: objc.GetClass("NSLayoutXAxisAnchor")}
	})
	return _NSLayoutXAxisAnchorClass
}

// GetNSLayoutXAxisAnchorClass returns the class object for NSLayoutXAxisAnchor.
func GetNSLayoutXAxisAnchorClass() NSLayoutXAxisAnchorClass {
	return getNSLayoutXAxisAnchorClass()
}

type NSLayoutXAxisAnchorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSLayoutXAxisAnchorClass) Alloc() NSLayoutXAxisAnchor {
	rv := objc.Send[NSLayoutXAxisAnchor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A factory class for creating horizontal layout constraint objects using a
// fluent API.
//
// # Overview
// 
// [NSLayoutXAxisAnchor] adds type information to the methods inherited from
// [NSLayoutAnchor]. Specifically, the generic methods declared by
// [NSLayoutAnchor] must now take a matching [NSLayoutXAxisAnchor] object.
// 
// For more information on using layout anchors, see [NSLayoutAnchor].
//
// # Building system spacing constraints
//
//   - [NSLayoutXAxisAnchor.ConstraintEqualToSystemSpacingAfterAnchorMultiplier]: Returns a constraint that defines by how much the current anchor trails the specified anchor.
//   - [NSLayoutXAxisAnchor.ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchorMultiplier]: Returns a constraint that defines the minimum amount by which the current anchor trails the specified anchor.
//   - [NSLayoutXAxisAnchor.ConstraintLessThanOrEqualToSystemSpacingAfterAnchorMultiplier]: Returns a constraint that defines the maximum amount by which the current anchor trails the specified anchor.
//
// # Creating a layout dimension
//
//   - [NSLayoutXAxisAnchor.AnchorWithOffsetToAnchor]: Creates a layout dimension object from two anchors.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutXAxisAnchor
type NSLayoutXAxisAnchor struct {
	NSLayoutAnchor
}

// NSLayoutXAxisAnchorFromID constructs a [NSLayoutXAxisAnchor] from an objc.ID.
//
// A factory class for creating horizontal layout constraint objects using a
// fluent API.
func NSLayoutXAxisAnchorFromID(id objc.ID) NSLayoutXAxisAnchor {
	return NSLayoutXAxisAnchor{NSLayoutAnchor: NSLayoutAnchorFromID(id)}
}
// NOTE: NSLayoutXAxisAnchor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSLayoutXAxisAnchor] class.
//
// # Building system spacing constraints
//
//   - [INSLayoutXAxisAnchor.ConstraintEqualToSystemSpacingAfterAnchorMultiplier]: Returns a constraint that defines by how much the current anchor trails the specified anchor.
//   - [INSLayoutXAxisAnchor.ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchorMultiplier]: Returns a constraint that defines the minimum amount by which the current anchor trails the specified anchor.
//   - [INSLayoutXAxisAnchor.ConstraintLessThanOrEqualToSystemSpacingAfterAnchorMultiplier]: Returns a constraint that defines the maximum amount by which the current anchor trails the specified anchor.
//
// # Creating a layout dimension
//
//   - [INSLayoutXAxisAnchor.AnchorWithOffsetToAnchor]: Creates a layout dimension object from two anchors.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutXAxisAnchor
type INSLayoutXAxisAnchor interface {
	INSLayoutAnchor

	// Topic: Building system spacing constraints

	// Returns a constraint that defines by how much the current anchor trails the specified anchor.
	ConstraintEqualToSystemSpacingAfterAnchorMultiplier(anchor INSLayoutXAxisAnchor, multiplier float64) INSLayoutConstraint
	// Returns a constraint that defines the minimum amount by which the current anchor trails the specified anchor.
	ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchorMultiplier(anchor INSLayoutXAxisAnchor, multiplier float64) INSLayoutConstraint
	// Returns a constraint that defines the maximum amount by which the current anchor trails the specified anchor.
	ConstraintLessThanOrEqualToSystemSpacingAfterAnchorMultiplier(anchor INSLayoutXAxisAnchor, multiplier float64) INSLayoutConstraint

	// Topic: Creating a layout dimension

	// Creates a layout dimension object from two anchors.
	AnchorWithOffsetToAnchor(otherAnchor INSLayoutXAxisAnchor) INSLayoutDimension
}

// Init initializes the instance.
func (l NSLayoutXAxisAnchor) Init() NSLayoutXAxisAnchor {
	rv := objc.Send[NSLayoutXAxisAnchor](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l NSLayoutXAxisAnchor) Autorelease() NSLayoutXAxisAnchor {
	rv := objc.Send[NSLayoutXAxisAnchor](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSLayoutXAxisAnchor creates a new NSLayoutXAxisAnchor instance.
func NewNSLayoutXAxisAnchor() NSLayoutXAxisAnchor {
	class := getNSLayoutXAxisAnchorClass()
	rv := objc.Send[NSLayoutXAxisAnchor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a constraint that defines by how much the current anchor trails the
// specified anchor.
//
// anchor: The anchor to use as the starting point for the constraint.
//
// multiplier: The multiple of the system spacing to use as the distance between the two
// anchors.
//
// # Return Value
// 
// An [NSLayoutConstraint] object that imposes a specific distance between the
// current anchor and the object in the `anchor` parameter.
//
// # Discussion
// 
// The constraint causes the current anchor to trail the object in the
// `anchor` parameter. For example, in a left-to-right layout, the current
// anchor is to the right of `anchor`, but in a right-to-left layout, it’s
// to the left of `anchor`.
// 
// The distance between the two anchors is determined by multiplying the
// system spacing by the value in the `multiplier` parameter. The value of the
// system space is determined from information available from the anchors.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutXAxisAnchor/constraint(equalToSystemSpacingAfter:multiplier:)
func (l NSLayoutXAxisAnchor) ConstraintEqualToSystemSpacingAfterAnchorMultiplier(anchor INSLayoutXAxisAnchor, multiplier float64) INSLayoutConstraint {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("constraintEqualToSystemSpacingAfterAnchor:multiplier:"), anchor, multiplier)
	return NSLayoutConstraintFromID(rv)
}
// Returns a constraint that defines the minimum amount by which the current
// anchor trails the specified anchor.
//
// anchor: The anchor to use as the starting point for the constraint.
//
// multiplier: The multiple of the system spacing to use as the minimum distance between
// the two anchors.
//
// # Return Value
// 
// An [NSLayoutConstraint] object that imposes a minimum distance between the
// current anchor and the object in the `anchor` parameter.
//
// # Discussion
// 
// The constraint causes the current anchor to trail the object in the
// `anchor` parameter. For example, in a left-to-right layout, the current
// anchor is to the right of `anchor`, but in a right-to-left layout, it’s
// to the left of `anchor`.
// 
// The minimum distance between the two anchors is determined by multiplying
// the system spacing by the value in the `multiplier` parameter. (The actual
// distance must be greater than or equal to that value.) The value of the
// system space is determined from information available from the anchors.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutXAxisAnchor/constraint(greaterThanOrEqualToSystemSpacingAfter:multiplier:)
func (l NSLayoutXAxisAnchor) ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchorMultiplier(anchor INSLayoutXAxisAnchor, multiplier float64) INSLayoutConstraint {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("constraintGreaterThanOrEqualToSystemSpacingAfterAnchor:multiplier:"), anchor, multiplier)
	return NSLayoutConstraintFromID(rv)
}
// Returns a constraint that defines the maximum amount by which the current
// anchor trails the specified anchor.
//
// anchor: The anchor to use as the starting point for the constraint.
//
// multiplier: The multiple of the system spacing to use as the maximum distance between
// the two anchors.
//
// # Return Value
// 
// An [NSLayoutConstraint] object that imposes a maximum distance between the
// current anchor and the object in the `anchor` parameter.
//
// # Discussion
// 
// The constraint causes the current anchor to trail the object in the
// `anchor` parameter. For example, in a left-to-right layout, the current
// anchor is to the right of `anchor`, but in a right-to-left layout, it’s
// to the left of `anchor`.
// 
// The maximum distance between the two anchors is determined by multiplying
// the system spacing by the value in the `multiplier` parameter. (The actual
// distance must be less than or equal to that value.) The value of the system
// space is determined from information available from the anchors.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutXAxisAnchor/constraint(lessThanOrEqualToSystemSpacingAfter:multiplier:)
func (l NSLayoutXAxisAnchor) ConstraintLessThanOrEqualToSystemSpacingAfterAnchorMultiplier(anchor INSLayoutXAxisAnchor, multiplier float64) INSLayoutConstraint {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("constraintLessThanOrEqualToSystemSpacingAfterAnchor:multiplier:"), anchor, multiplier)
	return NSLayoutConstraintFromID(rv)
}
// Creates a layout dimension object from two anchors.
//
// otherAnchor: The second anchor to use when creating the layout dimension.
//
// # Return Value
// 
// The [NSLayoutDimension] object represented by the two anchors.
//
// # Discussion
// 
// Use the returned object to define constraints relative to the space between
// the current anchor and the object in the `otherAnchor` parameter.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutXAxisAnchor/anchorWithOffset(to:)
func (l NSLayoutXAxisAnchor) AnchorWithOffsetToAnchor(otherAnchor INSLayoutXAxisAnchor) INSLayoutDimension {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("anchorWithOffsetToAnchor:"), otherAnchor)
	return NSLayoutDimensionFromID(rv)
}

