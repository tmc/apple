// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSLayoutYAxisAnchor] class.
var (
	_NSLayoutYAxisAnchorClass     NSLayoutYAxisAnchorClass
	_NSLayoutYAxisAnchorClassOnce sync.Once
)

func getNSLayoutYAxisAnchorClass() NSLayoutYAxisAnchorClass {
	_NSLayoutYAxisAnchorClassOnce.Do(func() {
		_NSLayoutYAxisAnchorClass = NSLayoutYAxisAnchorClass{class: objc.GetClass("NSLayoutYAxisAnchor")}
	})
	return _NSLayoutYAxisAnchorClass
}

// GetNSLayoutYAxisAnchorClass returns the class object for NSLayoutYAxisAnchor.
func GetNSLayoutYAxisAnchorClass() NSLayoutYAxisAnchorClass {
	return getNSLayoutYAxisAnchorClass()
}

type NSLayoutYAxisAnchorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSLayoutYAxisAnchorClass) Alloc() NSLayoutYAxisAnchor {
	rv := objc.Send[NSLayoutYAxisAnchor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A factory class for creating vertical layout constraint objects using a
// fluent API.
//
// # Overview
// 
// [NSLayoutYAxisAnchor] adds type information to the methods inherited from
// [NSLayoutAnchor]. Specifically, the generic methods declared by
// [NSLayoutAnchor] must now take a matching [NSLayoutYAxisAnchor] object.
// 
// For more information on using layout anchors, see [NSLayoutAnchor].
//
// # Building system spacing constraints
//
//   - [NSLayoutYAxisAnchor.ConstraintEqualToSystemSpacingBelowAnchorMultiplier]: Returns a constraint that defines the specific distance at which the current anchor is positioned below the specified anchor.
//   - [NSLayoutYAxisAnchor.ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchorMultiplier]: Returns a constraint that defines the minimum distance by which the current anchor is positioned below the specified anchor.
//   - [NSLayoutYAxisAnchor.ConstraintLessThanOrEqualToSystemSpacingBelowAnchorMultiplier]: Returns a constraint that defines the maximum distance by which the current anchor is positioned below the specified anchor.
//
// # Creating a layout dimension
//
//   - [NSLayoutYAxisAnchor.AnchorWithOffsetToAnchor]: Creates a layout dimension object from two anchors.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutYAxisAnchor
type NSLayoutYAxisAnchor struct {
	NSLayoutAnchor
}

// NSLayoutYAxisAnchorFromID constructs a [NSLayoutYAxisAnchor] from an objc.ID.
//
// A factory class for creating vertical layout constraint objects using a
// fluent API.
func NSLayoutYAxisAnchorFromID(id objc.ID) NSLayoutYAxisAnchor {
	return NSLayoutYAxisAnchor{NSLayoutAnchor: NSLayoutAnchorFromID(id)}
}
// NOTE: NSLayoutYAxisAnchor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSLayoutYAxisAnchor] class.
//
// # Building system spacing constraints
//
//   - [INSLayoutYAxisAnchor.ConstraintEqualToSystemSpacingBelowAnchorMultiplier]: Returns a constraint that defines the specific distance at which the current anchor is positioned below the specified anchor.
//   - [INSLayoutYAxisAnchor.ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchorMultiplier]: Returns a constraint that defines the minimum distance by which the current anchor is positioned below the specified anchor.
//   - [INSLayoutYAxisAnchor.ConstraintLessThanOrEqualToSystemSpacingBelowAnchorMultiplier]: Returns a constraint that defines the maximum distance by which the current anchor is positioned below the specified anchor.
//
// # Creating a layout dimension
//
//   - [INSLayoutYAxisAnchor.AnchorWithOffsetToAnchor]: Creates a layout dimension object from two anchors.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutYAxisAnchor
type INSLayoutYAxisAnchor interface {
	INSLayoutAnchor

	// Topic: Building system spacing constraints

	// Returns a constraint that defines the specific distance at which the current anchor is positioned below the specified anchor.
	ConstraintEqualToSystemSpacingBelowAnchorMultiplier(anchor INSLayoutYAxisAnchor, multiplier float64) INSLayoutConstraint
	// Returns a constraint that defines the minimum distance by which the current anchor is positioned below the specified anchor.
	ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchorMultiplier(anchor INSLayoutYAxisAnchor, multiplier float64) INSLayoutConstraint
	// Returns a constraint that defines the maximum distance by which the current anchor is positioned below the specified anchor.
	ConstraintLessThanOrEqualToSystemSpacingBelowAnchorMultiplier(anchor INSLayoutYAxisAnchor, multiplier float64) INSLayoutConstraint

	// Topic: Creating a layout dimension

	// Creates a layout dimension object from two anchors.
	AnchorWithOffsetToAnchor(otherAnchor INSLayoutYAxisAnchor) INSLayoutDimension

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (l NSLayoutYAxisAnchor) Init() NSLayoutYAxisAnchor {
	rv := objc.Send[NSLayoutYAxisAnchor](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l NSLayoutYAxisAnchor) Autorelease() NSLayoutYAxisAnchor {
	rv := objc.Send[NSLayoutYAxisAnchor](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSLayoutYAxisAnchor creates a new NSLayoutYAxisAnchor instance.
func NewNSLayoutYAxisAnchor() NSLayoutYAxisAnchor {
	class := getNSLayoutYAxisAnchorClass()
	rv := objc.Send[NSLayoutYAxisAnchor](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Returns a constraint that defines the specific distance at which the
// current anchor is positioned below the specified anchor.
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
// The constraint causes the current anchor to be positioned below the object
// in the `anchor` parameter. The distance between the two anchors is
// determined by multiplying the system spacing by the value in the
// `multiplier` parameter. The value of the system spacing is determined from
// information available from the anchors. For example, if the anchors
// represent text baselines, the spacing is determined by the fonts used at
// those baselines.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutYAxisAnchor/constraint(equalToSystemSpacingBelow:multiplier:)
func (l NSLayoutYAxisAnchor) ConstraintEqualToSystemSpacingBelowAnchorMultiplier(anchor INSLayoutYAxisAnchor, multiplier float64) INSLayoutConstraint {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("constraintEqualToSystemSpacingBelowAnchor:multiplier:"), anchor, multiplier)
	return NSLayoutConstraintFromID(rv)
}

// Returns a constraint that defines the minimum distance by which the current
// anchor is positioned below the specified anchor.
//
// anchor: The anchor to use as the starting point for the constraint.
//
// multiplier: The multiple of the system spacing to use as the distance between the two
// anchors.
//
// # Return Value
// 
// An [NSLayoutConstraint] object that imposes a minimum distance between the
// current anchor and the object in the `anchor` parameter.
//
// # Discussion
// 
// The constraint causes the current anchor to be positioned below the object
// in the `anchor` parameter. The minimum distance between the two anchors is
// determined by multiplying the system spacing by the value in the
// `multiplier` parameter. The value of the system spacing is determined from
// information available from the anchors. For example, if the anchors
// represent text baselines, the spacing is determined by the fonts used at
// those baselines.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutYAxisAnchor/constraint(greaterThanOrEqualToSystemSpacingBelow:multiplier:)
func (l NSLayoutYAxisAnchor) ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchorMultiplier(anchor INSLayoutYAxisAnchor, multiplier float64) INSLayoutConstraint {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("constraintGreaterThanOrEqualToSystemSpacingBelowAnchor:multiplier:"), anchor, multiplier)
	return NSLayoutConstraintFromID(rv)
}

// Returns a constraint that defines the maximum distance by which the current
// anchor is positioned below the specified anchor.
//
// anchor: The anchor to use as the starting point for the constraint.
//
// multiplier: The multiple of the system spacing to use as the distance between the two
// anchors.
//
// # Return Value
// 
// An [NSLayoutConstraint] object that imposes a minimum distance between the
// current anchor and the object in the `anchor` parameter.
//
// # Discussion
// 
// The constraint causes the current anchor to be positioned below the object
// in the `anchor` parameter. The maximum distance between the two anchors is
// determined by multiplying the system spacing by the value in the
// `multiplier` parameter. The value of the system spacing is determined from
// information available from the anchors. For example, if the anchors
// represent text baselines, the spacing is determined by the fonts used at
// those baselines.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutYAxisAnchor/constraint(lessThanOrEqualToSystemSpacingBelow:multiplier:)
func (l NSLayoutYAxisAnchor) ConstraintLessThanOrEqualToSystemSpacingBelowAnchorMultiplier(anchor INSLayoutYAxisAnchor, multiplier float64) INSLayoutConstraint {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("constraintLessThanOrEqualToSystemSpacingBelowAnchor:multiplier:"), anchor, multiplier)
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
// See: https://developer.apple.com/documentation/AppKit/NSLayoutYAxisAnchor/anchorWithOffset(to:)
func (l NSLayoutYAxisAnchor) AnchorWithOffsetToAnchor(otherAnchor INSLayoutYAxisAnchor) INSLayoutDimension {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("anchorWithOffsetToAnchor:"), otherAnchor)
	return NSLayoutDimensionFromID(rv)
}
func (l NSLayoutYAxisAnchor) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](l.ID, objc.Sel("encodeWithCoder:"), coder)
}




































