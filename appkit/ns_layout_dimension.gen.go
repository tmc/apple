// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSLayoutDimension] class.
var (
	_NSLayoutDimensionClass     NSLayoutDimensionClass
	_NSLayoutDimensionClassOnce sync.Once
)

func getNSLayoutDimensionClass() NSLayoutDimensionClass {
	_NSLayoutDimensionClassOnce.Do(func() {
		_NSLayoutDimensionClass = NSLayoutDimensionClass{class: objc.GetClass("NSLayoutDimension")}
	})
	return _NSLayoutDimensionClass
}

// GetNSLayoutDimensionClass returns the class object for NSLayoutDimension.
func GetNSLayoutDimensionClass() NSLayoutDimensionClass {
	return getNSLayoutDimensionClass()
}

type NSLayoutDimensionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSLayoutDimensionClass) Alloc() NSLayoutDimension {
	rv := objc.Send[NSLayoutDimension](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A factory class for creating size-based layout constraint objects using a
// fluent API.
//
// # Overview
// 
// Use these constraints to programmatically define your layout using Auto
// Layout. All sizes are measured in points. In addition to providing
// size-specific methods for creating constraints, this class adds type
// information to the methods inherited from [NSLayoutAnchor]. Specifically,
// the generic methods declared by [NSLayoutAnchor] must now take a matching
// [NSLayoutDimension] object.
// 
// For more information on using layout anchors, see [NSLayoutAnchor].
//
// # Building constraints
//
//   - [NSLayoutDimension.ConstraintEqualToAnchorMultiplier]: Returns a constraint that defines the anchor’s size attribute as equal to the specified anchor multiplied by the constant.
//   - [NSLayoutDimension.ConstraintEqualToAnchorMultiplierConstant]: Returns a constraint that defines the anchor’s size attribute as equal to the specified size attribute multiplied by a constant plus an offset.
//   - [NSLayoutDimension.ConstraintEqualToConstant]: Returns a constraint that defines a constant size for the anchor’s size attribute.
//   - [NSLayoutDimension.ConstraintGreaterThanOrEqualToAnchorMultiplier]: Returns a constraint that defines the anchor’s size attribute as greater than or equal to the specified anchor multiplied by the constant.
//   - [NSLayoutDimension.ConstraintGreaterThanOrEqualToAnchorMultiplierConstant]: Returns a constraint that defines the anchor’s size attribute as greater than or equal to the specified anchor multiplied by the constant plus an offset.
//   - [NSLayoutDimension.ConstraintGreaterThanOrEqualToConstant]: Returns a constraint that defines the minimum size for the anchor’s size attribute.
//   - [NSLayoutDimension.ConstraintLessThanOrEqualToAnchorMultiplier]: Returns a constraint that defines the anchor’s size attribute as less than or equal to the specified anchor multiplied by the constant.
//   - [NSLayoutDimension.ConstraintLessThanOrEqualToAnchorMultiplierConstant]: Returns a constraint that defines the anchor’s size attribute as greater than or equal to the specified anchor multiplied by the constant plus an offset.
//   - [NSLayoutDimension.ConstraintLessThanOrEqualToConstant]: Returns a constraint that defines the maximum size for the anchor’s size attribute.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutDimension
type NSLayoutDimension struct {
	NSLayoutAnchor
}

// NSLayoutDimensionFromID constructs a [NSLayoutDimension] from an objc.ID.
//
// A factory class for creating size-based layout constraint objects using a
// fluent API.
func NSLayoutDimensionFromID(id objc.ID) NSLayoutDimension {
	return NSLayoutDimension{NSLayoutAnchor: NSLayoutAnchorFromID(id)}
}
// NOTE: NSLayoutDimension adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSLayoutDimension] class.
//
// # Building constraints
//
//   - [INSLayoutDimension.ConstraintEqualToAnchorMultiplier]: Returns a constraint that defines the anchor’s size attribute as equal to the specified anchor multiplied by the constant.
//   - [INSLayoutDimension.ConstraintEqualToAnchorMultiplierConstant]: Returns a constraint that defines the anchor’s size attribute as equal to the specified size attribute multiplied by a constant plus an offset.
//   - [INSLayoutDimension.ConstraintEqualToConstant]: Returns a constraint that defines a constant size for the anchor’s size attribute.
//   - [INSLayoutDimension.ConstraintGreaterThanOrEqualToAnchorMultiplier]: Returns a constraint that defines the anchor’s size attribute as greater than or equal to the specified anchor multiplied by the constant.
//   - [INSLayoutDimension.ConstraintGreaterThanOrEqualToAnchorMultiplierConstant]: Returns a constraint that defines the anchor’s size attribute as greater than or equal to the specified anchor multiplied by the constant plus an offset.
//   - [INSLayoutDimension.ConstraintGreaterThanOrEqualToConstant]: Returns a constraint that defines the minimum size for the anchor’s size attribute.
//   - [INSLayoutDimension.ConstraintLessThanOrEqualToAnchorMultiplier]: Returns a constraint that defines the anchor’s size attribute as less than or equal to the specified anchor multiplied by the constant.
//   - [INSLayoutDimension.ConstraintLessThanOrEqualToAnchorMultiplierConstant]: Returns a constraint that defines the anchor’s size attribute as greater than or equal to the specified anchor multiplied by the constant plus an offset.
//   - [INSLayoutDimension.ConstraintLessThanOrEqualToConstant]: Returns a constraint that defines the maximum size for the anchor’s size attribute.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutDimension
type INSLayoutDimension interface {
	INSLayoutAnchor

	// Topic: Building constraints

	// Returns a constraint that defines the anchor’s size attribute as equal to the specified anchor multiplied by the constant.
	ConstraintEqualToAnchorMultiplier(anchor INSLayoutDimension, m float64) INSLayoutConstraint
	// Returns a constraint that defines the anchor’s size attribute as equal to the specified size attribute multiplied by a constant plus an offset.
	ConstraintEqualToAnchorMultiplierConstant(anchor INSLayoutDimension, m float64, c float64) INSLayoutConstraint
	// Returns a constraint that defines a constant size for the anchor’s size attribute.
	ConstraintEqualToConstant(c float64) INSLayoutConstraint
	// Returns a constraint that defines the anchor’s size attribute as greater than or equal to the specified anchor multiplied by the constant.
	ConstraintGreaterThanOrEqualToAnchorMultiplier(anchor INSLayoutDimension, m float64) INSLayoutConstraint
	// Returns a constraint that defines the anchor’s size attribute as greater than or equal to the specified anchor multiplied by the constant plus an offset.
	ConstraintGreaterThanOrEqualToAnchorMultiplierConstant(anchor INSLayoutDimension, m float64, c float64) INSLayoutConstraint
	// Returns a constraint that defines the minimum size for the anchor’s size attribute.
	ConstraintGreaterThanOrEqualToConstant(c float64) INSLayoutConstraint
	// Returns a constraint that defines the anchor’s size attribute as less than or equal to the specified anchor multiplied by the constant.
	ConstraintLessThanOrEqualToAnchorMultiplier(anchor INSLayoutDimension, m float64) INSLayoutConstraint
	// Returns a constraint that defines the anchor’s size attribute as greater than or equal to the specified anchor multiplied by the constant plus an offset.
	ConstraintLessThanOrEqualToAnchorMultiplierConstant(anchor INSLayoutDimension, m float64, c float64) INSLayoutConstraint
	// Returns a constraint that defines the maximum size for the anchor’s size attribute.
	ConstraintLessThanOrEqualToConstant(c float64) INSLayoutConstraint

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (l NSLayoutDimension) Init() NSLayoutDimension {
	rv := objc.Send[NSLayoutDimension](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l NSLayoutDimension) Autorelease() NSLayoutDimension {
	rv := objc.Send[NSLayoutDimension](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSLayoutDimension creates a new NSLayoutDimension instance.
func NewNSLayoutDimension() NSLayoutDimension {
	class := getNSLayoutDimensionClass()
	rv := objc.Send[NSLayoutDimension](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Returns a constraint that defines the anchor’s size attribute as equal to
// the specified anchor multiplied by the constant.
//
// anchor: A dimension anchor from an [NSView] or [NSLayoutGuide] object.
//
// m: The multiplier constant for the constraint.
//
// # Return Value
// 
// An [NSLayoutConstraint] object that defines the attribute represented by
// this layout anchor as equal to the attribute represented by the `anchor`
// parameter multiplied by the `m` constant.
//
// # Discussion
// 
// This method defines the relationship `first attribute = m * second
// attribute`. Where `first attribute` is the layout attribute represented by
// the anchor receiving this method call, and `second attribute` is the layout
// attribute represented by the `anchor` parameter.
// 
// The constraints produced by the following two examples are identical.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutDimension/constraint(equalTo:multiplier:)
func (l NSLayoutDimension) ConstraintEqualToAnchorMultiplier(anchor INSLayoutDimension, m float64) INSLayoutConstraint {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("constraintEqualToAnchor:multiplier:"), anchor, m)
	return NSLayoutConstraintFromID(rv)
}

// Returns a constraint that defines the anchor’s size attribute as equal to
// the specified size attribute multiplied by a constant plus an offset.
//
// anchor: A dimension anchor from an [NSView] or [NSLayoutGuide] object.
//
// m: The multiplier constant for the constraint.
//
// c: The offset constant for this relationship.
//
// # Return Value
// 
// An [NSLayoutConstraint] object that defines the attribute represented by
// this layout anchor as equal to the attribute represented by the `anchor`
// parameter multiplied by the `m` constant plus the constant `c`.
//
// # Discussion
// 
// This method defines the relationship `first attribute = (m * second
// attribute) + c`. Where `first attribute` is the layout attribute
// represented by the anchor receiving this method call, and `second
// attribute` is the layout attribute represented by the `anchor` parameter.
// 
// The constraints produced by the following two examples are identical.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutDimension/constraint(equalTo:multiplier:constant:)
func (l NSLayoutDimension) ConstraintEqualToAnchorMultiplierConstant(anchor INSLayoutDimension, m float64, c float64) INSLayoutConstraint {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("constraintEqualToAnchor:multiplier:constant:"), anchor, m, c)
	return NSLayoutConstraintFromID(rv)
}

// Returns a constraint that defines a constant size for the anchor’s size
// attribute.
//
// c: A constant representing the size of the attribute associated with this
// dimension anchor.
//
// # Return Value
// 
// An [NSLayoutConstraint] object that defines a constant size for the
// attribute associated with this dimension anchor.
//
// # Discussion
// 
// This method defines the relationship `first attribute = c`. Where `first
// attribute` is the layout attribute represented by the anchor receiving this
// method call.
// 
// The constraints produced by the following two examples are identical.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutDimension/constraint(equalToConstant:)
func (l NSLayoutDimension) ConstraintEqualToConstant(c float64) INSLayoutConstraint {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("constraintEqualToConstant:"), c)
	return NSLayoutConstraintFromID(rv)
}

// Returns a constraint that defines the anchor’s size attribute as greater
// than or equal to the specified anchor multiplied by the constant.
//
// anchor: A dimension anchor from an [NSView] or [NSLayoutGuide] object.
//
// m: The multiplier constant for the constraint.
//
// # Return Value
// 
// An [NSLayoutConstraint] object that defines the attribute represented by
// this layout anchor as greater than or equal to the attribute represented by
// the `anchor` parameter multiplied by the `m` constant.
//
// # Discussion
// 
// This method defines the relationship `first attribute >= m * second
// attribute`. Where `first attribute` is the layout attribute represented by
// the anchor receiving this method call, and `second attribute` is the layout
// attribute represented by the `anchor` parameter.
// 
// The constraints produced by the following two examples are identical.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutDimension/constraint(greaterThanOrEqualTo:multiplier:)
func (l NSLayoutDimension) ConstraintGreaterThanOrEqualToAnchorMultiplier(anchor INSLayoutDimension, m float64) INSLayoutConstraint {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("constraintGreaterThanOrEqualToAnchor:multiplier:"), anchor, m)
	return NSLayoutConstraintFromID(rv)
}

// Returns a constraint that defines the anchor’s size attribute as greater
// than or equal to the specified anchor multiplied by the constant plus an
// offset.
//
// anchor: A dimension anchor from an [NSView] or [NSLayoutGuide] object.
//
// m: The multiplier constant for the constraint.
//
// c: The constant offset for this relationship.
//
// # Return Value
// 
// An [NSLayoutConstraint] object that defines the attribute represented by
// this layout anchor as greater than or equal to the attribute represented by
// the `anchor` parameter multiplied by the `m` constant plus the constant
// `c`.
//
// # Discussion
// 
// This method defines the relationship `first attribute >= (m * second
// attribute) + c`. Where `first attribute` is the layout attribute
// represented by the anchor receiving this method call, and `second
// attribute` is the layout attribute represented by the `anchor` parameter.
// 
// The constraints produced by the following two examples are identical.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutDimension/constraint(greaterThanOrEqualTo:multiplier:constant:)
func (l NSLayoutDimension) ConstraintGreaterThanOrEqualToAnchorMultiplierConstant(anchor INSLayoutDimension, m float64, c float64) INSLayoutConstraint {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("constraintGreaterThanOrEqualToAnchor:multiplier:constant:"), anchor, m, c)
	return NSLayoutConstraintFromID(rv)
}

// Returns a constraint that defines the minimum size for the anchor’s size
// attribute.
//
// c: A constant representing the minimum size of the attribute associated with
// this dimension anchor.
//
// # Return Value
// 
// An [NSLayoutConstraint] object that defines a minimum size for the
// attribute associated with this dimension anchor.
//
// # Discussion
// 
// This method defines the relationship `first attribute >= c`. Where `first
// attribute` is the layout attribute represented by the anchor receiving this
// method call.
// 
// The constraints produced by the following two examples are identical.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutDimension/constraint(greaterThanOrEqualToConstant:)
func (l NSLayoutDimension) ConstraintGreaterThanOrEqualToConstant(c float64) INSLayoutConstraint {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("constraintGreaterThanOrEqualToConstant:"), c)
	return NSLayoutConstraintFromID(rv)
}

// Returns a constraint that defines the anchor’s size attribute as less
// than or equal to the specified anchor multiplied by the constant.
//
// anchor: A dimension anchor from an [NSView] or [NSLayoutGuide] object.
//
// m: The multiplier constant for the constraint.
//
// # Return Value
// 
// An [NSLayoutConstraint] object that defines the attribute represented by
// this layout anchor as less than or equal to the attribute represented by
// the `anchor` parameter multiplied by the `m` constant.
//
// # Discussion
// 
// This method defines the relationship `first attribute <= m * second
// attribute` . Where `first attribute` is the layout attribute represented by
// the anchor receiving this method call, and `second attribute` is the layout
// attribute represented by the `anchor` parameter.
// 
// The constraints produced by the following two examples are identical.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutDimension/constraint(lessThanOrEqualTo:multiplier:)
func (l NSLayoutDimension) ConstraintLessThanOrEqualToAnchorMultiplier(anchor INSLayoutDimension, m float64) INSLayoutConstraint {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("constraintLessThanOrEqualToAnchor:multiplier:"), anchor, m)
	return NSLayoutConstraintFromID(rv)
}

// Returns a constraint that defines the anchor’s size attribute as greater
// than or equal to the specified anchor multiplied by the constant plus an
// offset.
//
// anchor: A dimension anchor from an [NSView] or [NSLayoutGuide] object.
//
// m: The multiplier constant for the constraint.
//
// c: The constant offset for this relationship.
//
// # Return Value
// 
// An [NSLayoutConstraint] object that defines the attribute represented by
// this layout anchor as less than or equal to the attribute represented by
// the `anchor` parameter multiplied by the `m` constant plus the constant
// `c`.
//
// # Discussion
// 
// This method defines the relationship `first attribute <= (m * second
// attribute) + c`. Where `first attribute` is the layout attribute
// represented by the anchor receiving this method call, and `second
// attribute` is the layout attribute represented by the `anchor` parameter.
// 
// The constraints produced by the following two examples are identical.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutDimension/constraint(lessThanOrEqualTo:multiplier:constant:)
func (l NSLayoutDimension) ConstraintLessThanOrEqualToAnchorMultiplierConstant(anchor INSLayoutDimension, m float64, c float64) INSLayoutConstraint {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("constraintLessThanOrEqualToAnchor:multiplier:constant:"), anchor, m, c)
	return NSLayoutConstraintFromID(rv)
}

// Returns a constraint that defines the maximum size for the anchor’s size
// attribute.
//
// c: A constant representing the maximum size of the attribute associated with
// this dimension anchor.
//
// # Return Value
// 
// An [NSLayoutConstraint] object that defines a maximum size for the
// attribute associated with this dimension anchor.
//
// # Discussion
// 
// This method defines the relationship `first attribute <= c`. Where `first
// attribute` is the layout attribute represented by the anchor receiving this
// method call.
// 
// The constraints produced by the following two examples are identical.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutDimension/constraint(lessThanOrEqualToConstant:)
func (l NSLayoutDimension) ConstraintLessThanOrEqualToConstant(c float64) INSLayoutConstraint {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("constraintLessThanOrEqualToConstant:"), c)
	return NSLayoutConstraintFromID(rv)
}
func (l NSLayoutDimension) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](l.ID, objc.Sel("encodeWithCoder:"), coder)
}




































