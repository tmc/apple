// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSLayoutAnchor] class.
var (
	_NSLayoutAnchorClass     NSLayoutAnchorClass
	_NSLayoutAnchorClassOnce sync.Once
)

func getNSLayoutAnchorClass() NSLayoutAnchorClass {
	_NSLayoutAnchorClassOnce.Do(func() {
		_NSLayoutAnchorClass = NSLayoutAnchorClass{class: objc.GetClass("NSLayoutAnchor")}
	})
	return _NSLayoutAnchorClass
}

// GetNSLayoutAnchorClass returns the class object for NSLayoutAnchor.
func GetNSLayoutAnchorClass() NSLayoutAnchorClass {
	return getNSLayoutAnchorClass()
}

type NSLayoutAnchorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSLayoutAnchorClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSLayoutAnchorClass) Alloc() NSLayoutAnchor {
	rv := objc.Send[NSLayoutAnchor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A factory class for creating layout constraint objects using a fluent API.
//
// # Overview
// 
// Use these constraints to programatically define your layout using Auto
// Layout. Instead of creating [NSLayoutConstraint] objects directly, start
// with an [NSView] or [NSLayoutGuide] object you wish to constrain, and
// select one of that object’s anchor properties. These properties
// correspond to the main [NSLayoutConstraint.Attribute] values used in Auto
// Layout, and provide an appropriate [NSLayoutAnchor] subclass for creating
// constraints to that attribute. Use the anchor’s methods to construct your
// constraint.
// 
// As you can see from these examples, the [NSLayoutAnchor] class provides
// several advantages over using the [NSLayoutConstraint] API directly.
// 
// - The code is cleaner, more concise, and easier to read. - The
// [NSLayoutConstraint.Attribute] subclasses provide additional type checking,
// preventing you from creating invalid constraints.
// 
// For more information on the anchor properties, see [bottomAnchor] in the
// [NSView] or [NSLayoutGuide].
//
// [NSLayoutConstraint.Attribute]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute
// [bottomAnchor]: https://developer.apple.com/documentation/AppKit/NSView/bottomAnchor
//
// # Building constraints
//
//   - [NSLayoutAnchor.ConstraintEqualToAnchor]: Returns a constraint that defines one item’s attribute as equal to another.
//   - [NSLayoutAnchor.ConstraintEqualToAnchorConstant]: Returns a constraint that defines one item’s attribute as equal to another item’s attribute plus a constant offset.
//   - [NSLayoutAnchor.ConstraintGreaterThanOrEqualToAnchor]: Returns a constraint that defines one item’s attribute as greater than or equal to another.
//   - [NSLayoutAnchor.ConstraintGreaterThanOrEqualToAnchorConstant]: Returns a constraint that defines one item’s attribute as greater than or equal to another item’s attribute plus a constant offset.
//   - [NSLayoutAnchor.ConstraintLessThanOrEqualToAnchor]: Returns a constraint that defines one item’s attribute as less than or equal to another.
//   - [NSLayoutAnchor.ConstraintLessThanOrEqualToAnchorConstant]: Returns a constraint that defines one item’s attribute as less than or equal to another item’s attribute plus a constant offset.
//
// # Debugging the anchor
//
//   - [NSLayoutAnchor.ConstraintsAffectingLayout]: The constraints that impact the layout of the anchor.
//   - [NSLayoutAnchor.HasAmbiguousLayout]: A Boolean value indicating whether the constraints impacting the anchor specify its location ambiguously.
//   - [NSLayoutAnchor.Name]: The name assigned to the anchor for debugging purposes.
//   - [NSLayoutAnchor.Item]: The layout item used to calculate the anchor’s position.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor
type NSLayoutAnchor struct {
	objectivec.Object
}

// NSLayoutAnchorFromID constructs a [NSLayoutAnchor] from an objc.ID.
//
// A factory class for creating layout constraint objects using a fluent API.
func NSLayoutAnchorFromID(id objc.ID) NSLayoutAnchor {
	return NSLayoutAnchor{objectivec.Object{ID: id}}
}
// NOTE: NSLayoutAnchor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSLayoutAnchor] class.
//
// # Building constraints
//
//   - [INSLayoutAnchor.ConstraintEqualToAnchor]: Returns a constraint that defines one item’s attribute as equal to another.
//   - [INSLayoutAnchor.ConstraintEqualToAnchorConstant]: Returns a constraint that defines one item’s attribute as equal to another item’s attribute plus a constant offset.
//   - [INSLayoutAnchor.ConstraintGreaterThanOrEqualToAnchor]: Returns a constraint that defines one item’s attribute as greater than or equal to another.
//   - [INSLayoutAnchor.ConstraintGreaterThanOrEqualToAnchorConstant]: Returns a constraint that defines one item’s attribute as greater than or equal to another item’s attribute plus a constant offset.
//   - [INSLayoutAnchor.ConstraintLessThanOrEqualToAnchor]: Returns a constraint that defines one item’s attribute as less than or equal to another.
//   - [INSLayoutAnchor.ConstraintLessThanOrEqualToAnchorConstant]: Returns a constraint that defines one item’s attribute as less than or equal to another item’s attribute plus a constant offset.
//
// # Debugging the anchor
//
//   - [INSLayoutAnchor.ConstraintsAffectingLayout]: The constraints that impact the layout of the anchor.
//   - [INSLayoutAnchor.HasAmbiguousLayout]: A Boolean value indicating whether the constraints impacting the anchor specify its location ambiguously.
//   - [INSLayoutAnchor.Name]: The name assigned to the anchor for debugging purposes.
//   - [INSLayoutAnchor.Item]: The layout item used to calculate the anchor’s position.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor
type INSLayoutAnchor interface {
	objectivec.IObject

	// Topic: Building constraints

	// Returns a constraint that defines one item’s attribute as equal to another.
	ConstraintEqualToAnchor(anchor INSLayoutAnchor) INSLayoutConstraint
	// Returns a constraint that defines one item’s attribute as equal to another item’s attribute plus a constant offset.
	ConstraintEqualToAnchorConstant(anchor INSLayoutAnchor, c float64) INSLayoutConstraint
	// Returns a constraint that defines one item’s attribute as greater than or equal to another.
	ConstraintGreaterThanOrEqualToAnchor(anchor INSLayoutAnchor) INSLayoutConstraint
	// Returns a constraint that defines one item’s attribute as greater than or equal to another item’s attribute plus a constant offset.
	ConstraintGreaterThanOrEqualToAnchorConstant(anchor INSLayoutAnchor, c float64) INSLayoutConstraint
	// Returns a constraint that defines one item’s attribute as less than or equal to another.
	ConstraintLessThanOrEqualToAnchor(anchor INSLayoutAnchor) INSLayoutConstraint
	// Returns a constraint that defines one item’s attribute as less than or equal to another item’s attribute plus a constant offset.
	ConstraintLessThanOrEqualToAnchorConstant(anchor INSLayoutAnchor, c float64) INSLayoutConstraint

	// Topic: Debugging the anchor

	// The constraints that impact the layout of the anchor.
	ConstraintsAffectingLayout() []NSLayoutConstraint
	// A Boolean value indicating whether the constraints impacting the anchor specify its location ambiguously.
	HasAmbiguousLayout() bool
	// The name assigned to the anchor for debugging purposes.
	Name() string
	// The layout item used to calculate the anchor’s position.
	Item() objectivec.IObject

	// A layout anchor representing the bottom edge of the view’s frame.
	BottomAnchor() INSLayoutYAxisAnchor
	SetBottomAnchor(value INSLayoutYAxisAnchor)
	// A layout anchor representing the leading edge of the view’s frame.
	LeadingAnchor() INSLayoutXAxisAnchor
	SetLeadingAnchor(value INSLayoutXAxisAnchor)
	// A layout anchor representing the left edge of the view’s frame.
	LeftAnchor() INSLayoutXAxisAnchor
	SetLeftAnchor(value INSLayoutXAxisAnchor)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (l NSLayoutAnchor) Init() NSLayoutAnchor {
	rv := objc.Send[NSLayoutAnchor](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l NSLayoutAnchor) Autorelease() NSLayoutAnchor {
	rv := objc.Send[NSLayoutAnchor](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSLayoutAnchor creates a new NSLayoutAnchor instance.
func NewNSLayoutAnchor() NSLayoutAnchor {
	class := getNSLayoutAnchorClass()
	rv := objc.Send[NSLayoutAnchor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a constraint that defines one item’s attribute as equal to
// another.
//
// anchor: A layout anchor from an [NSView] or [NSLayoutGuide] object. You must use a
// subclass of [NSLayoutAnchor] that matches the current anchor. For example,
// if you call this method on an [NSLayoutXAxisAnchor] object, this parameter
// must be another [NSLayoutXAxisAnchor].
//
// # Return Value
// 
// An [NSLayoutConstraint] object that defines an equal relationship between
// the attributes represented by the two layout anchors.
//
// # Discussion
// 
// This method defines the relationship `first attribute = second attribute`.
// Where `first attribute` is the layout attribute represented by the anchor
// receiving this method call, and `second attribute` is the layout attribute
// represented by the `anchor` parameter.
// 
// The constraints produced by the following two examples are identical.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor/constraint(equalTo:)
func (l NSLayoutAnchor) ConstraintEqualToAnchor(anchor INSLayoutAnchor) INSLayoutConstraint {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("constraintEqualToAnchor:"), anchor)
	return NSLayoutConstraintFromID(rv)
}
// Returns a constraint that defines one item’s attribute as equal to
// another item’s attribute plus a constant offset.
//
// anchor: A layout anchor from an [NSView] or [NSLayoutGuide] object. You must use a
// subclass of [NSLayoutAnchor] that matches the current anchor. For example,
// if you call this method on an [NSLayoutXAxisAnchor] object, this parameter
// must be another [NSLayoutXAxisAnchor].
//
// c: The constant offset for the constraint.
//
// # Return Value
// 
// An [NSLayoutConstraint] object that defines an equal relationship between
// the attributes represented by the two layout anchors plus a constant
// offset.
//
// # Discussion
// 
// This method defines the relationship `first attribute = second attribute +
// c`. Where `first attribute` is the layout attribute represented by the
// anchor receiving this method call, and `second attribute` is the layout
// attribute represented by the `anchor` parameter. The value `c`, represents
// a constant offset. All values are measured in points; however, these values
// can be interpreted in different ways, depending on the type of layout
// anchor.
// 
// - For [NSLayoutXAxisAnchor] objects, the first attribute is positioned `c`
// points after the second attribute. When using leading or trailing
// attributes, values increase as you move in the language’s reading
// direction. In English, for example, values increase as you move to the
// right. For left and right attributes, values always increase as you move
// right. - For [NSLayoutYAxisAnchor] objects, the first attribute is
// positioned `c` points below the second attribute. Values increase as you
// move down. - For [NSLayoutDimension] objects, the size of the first
// attribute is `c` points larger than the size of the second attribute.
// Values increase as items increase in size.
// 
// The constraints produced by the following two examples are identical.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor/constraint(equalTo:constant:)
func (l NSLayoutAnchor) ConstraintEqualToAnchorConstant(anchor INSLayoutAnchor, c float64) INSLayoutConstraint {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("constraintEqualToAnchor:constant:"), anchor, c)
	return NSLayoutConstraintFromID(rv)
}
// Returns a constraint that defines one item’s attribute as greater than or
// equal to another.
//
// anchor: A layout anchor from an [NSView] or [NSLayoutGuide] object. You must use a
// subclass of [NSLayoutAnchor] that matches the current anchor. For example,
// if you call this method on an [NSLayoutXAxisAnchor] object, this parameter
// must be another [NSLayoutXAxisAnchor].
//
// # Return Value
// 
// An [NSLayoutConstraint] object that defines the attribute represented by
// this layout anchor as greater than or equal to the attribute represented by
// the `anchor` parameter.
//
// # Discussion
// 
// This method creates a relationship where `first attribute >= second
// attribute`. Where `first attribute` is the layout attribute represented by
// the anchor receiving this method call, and `second attribute` is the layout
// attribute represented by the `anchor` parameter. All values are measured in
// points; however, these values can be interpreted in different ways,
// depending on the type of layout anchor.
// 
// - For leading or trailing anchors, the values increase as you move in the
// current language’s reading direction. In English, for example, values
// increase as you move to the right. - For left and right anchors, the values
// increase as you move to the right. - For [NSLayoutYAxisAnchor] objects, the
// values increase as you move down. - For [NSLayoutDimension] objects, the
// values increase as the items increase in size.
// 
// The constraints produced by the following two examples are identical.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor/constraint(greaterThanOrEqualTo:)
func (l NSLayoutAnchor) ConstraintGreaterThanOrEqualToAnchor(anchor INSLayoutAnchor) INSLayoutConstraint {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("constraintGreaterThanOrEqualToAnchor:"), anchor)
	return NSLayoutConstraintFromID(rv)
}
// Returns a constraint that defines one item’s attribute as greater than or
// equal to another item’s attribute plus a constant offset.
//
// anchor: A layout anchor from an [NSView] or [NSLayoutGuide] object. You must use a
// subclass of [NSLayoutAnchor] that matches the current anchor. For example,
// if you call this method on an [NSLayoutXAxisAnchor] object, this parameter
// must be another [NSLayoutXAxisAnchor].
//
// c: The constant offset for the constraint.
//
// # Return Value
// 
// An [NSLayoutConstraint] object that defines the attribute represented by
// this layout anchor as greater than or equal to the attribute represented by
// the `anchor` parameter plus a constant offset.
//
// # Discussion
// 
// This method defines the relationship `first attribute >= second attribute +
// c`. Where `first attribute` is the layout attribute represented by the
// anchor receiving this method call, and `second attribute` is the layout
// attribute represented by the `anchor` parameter. The value `c`, represents
// a constant offset. All values are measured in points; however, these values
// can be interpreted in different ways, depending on the type of layout
// anchor.
// 
// - For [NSLayoutXAxisAnchor] objects, the first attribute is positioned `c`
// points after the second attribute. When using leading or trailing
// attributes, values increase as you move in the language’s reading
// direction. In English, for example, values increase as you move to the
// right. For left and right attributes, values always increase as you move
// right. - For [NSLayoutYAxisAnchor] objects, the first attribute is
// positioned `c` points below the second attribute. Values increase as you
// move down. - For [NSLayoutDimension] objects, the size of the first
// attribute is `c` points larger than the size of the second attribute.
// Values increase as items increase in size.
// 
// The constraints produced by the following two examples are identical.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor/constraint(greaterThanOrEqualTo:constant:)
func (l NSLayoutAnchor) ConstraintGreaterThanOrEqualToAnchorConstant(anchor INSLayoutAnchor, c float64) INSLayoutConstraint {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("constraintGreaterThanOrEqualToAnchor:constant:"), anchor, c)
	return NSLayoutConstraintFromID(rv)
}
// Returns a constraint that defines one item’s attribute as less than or
// equal to another.
//
// anchor: A layout anchor from an [NSView] or [NSLayoutGuide] object. You must use a
// subclass of [NSLayoutAnchor] that matches the current anchor. For example,
// if you call this method on an [NSLayoutXAxisAnchor] object, this parameter
// must be another [NSLayoutXAxisAnchor].
//
// # Return Value
// 
// An [NSLayoutConstraint] object that defines the attribute represented by
// this layout anchor as less than or equal to the attribute represented by
// the `anchor` parameter.
//
// # Discussion
// 
// This method defines the relationship `first attribute <= second attribute`.
// Where `first attribute` is the layout attribute represented by the anchor
// receiving this method call, and `second attribute` is the layout attribute
// represented by the `anchor` parameter. All values are measured in points;
// however, these values can be interpreted in different ways, depending on
// the type of layout anchor.
// 
// - For leading or trailing anchors, the values increase as you move in the
// current language’s reading direction. For English, values increase as you
// move to the right. - For left and right anchors, the values increase as you
// move to the right. - For [NSLayoutYAxisAnchor] objects, the values increase
// as you move down. - For [NSLayoutDimension] objects, the values increase as
// the items increase in size.
// 
// The constraints produced by the following two examples are identical.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor/constraint(lessThanOrEqualTo:)
func (l NSLayoutAnchor) ConstraintLessThanOrEqualToAnchor(anchor INSLayoutAnchor) INSLayoutConstraint {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("constraintLessThanOrEqualToAnchor:"), anchor)
	return NSLayoutConstraintFromID(rv)
}
// Returns a constraint that defines one item’s attribute as less than or
// equal to another item’s attribute plus a constant offset.
//
// anchor: A layout anchor from an [NSView] or [NSLayoutGuide] object. You must use a
// subclass of [NSLayoutAnchor] that matches the current anchor. For example,
// if you call this method on an [NSLayoutXAxisAnchor] object, this parameter
// must be another [NSLayoutXAxisAnchor].
//
// c: The constant offset for the constraint.
//
// # Return Value
// 
// An [NSLayoutConstraint] object that defines the attribute represented by
// this layout anchor as less than or equal to the attribute represented by
// the `anchor` parameter plus a constant offset.
//
// # Discussion
// 
// This method defines the relationship `first attribute <= second attribute +
// c`. Where `first attribute` is the layout attribute represented by the
// anchor receiving this method call, and `second attribute` is the layout
// attribute represented by the `anchor` parameter. The value `c`, represents
// a constant offset. All values are measured in points; however, these values
// can be interpreted in different ways, depending on the type of layout
// anchor.
// 
// - For [NSLayoutXAxisAnchor] objects, the first attribute is positioned `c`
// points after the second attribute. When using leading or trailing
// attributes, values increase as you move in the language’s reading
// direction. In English, for example, values increase as you move to the
// right. For left and right attributes, values always increase as you move
// right. - For [NSLayoutYAxisAnchor] objects, the first attribute is
// positioned `c` points below the second attribute. Values increase as you
// move down. - For [NSLayoutDimension] objects, the size of the first
// attribute is `c` points larger than the size of the second attribute.
// Values increase as items increase in size.
// 
// The constraints produced by the following two examples are identical.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor/constraint(lessThanOrEqualTo:constant:)
func (l NSLayoutAnchor) ConstraintLessThanOrEqualToAnchorConstant(anchor INSLayoutAnchor, c float64) INSLayoutConstraint {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("constraintLessThanOrEqualToAnchor:constant:"), anchor, c)
	return NSLayoutConstraintFromID(rv)
}
func (l NSLayoutAnchor) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](l.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The constraints that impact the layout of the anchor.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor/constraintsAffectingLayout
func (l NSLayoutAnchor) ConstraintsAffectingLayout() []NSLayoutConstraint {
	rv := objc.Send[[]objc.ID](l.ID, objc.Sel("constraintsAffectingLayout"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSLayoutConstraint {
		return NSLayoutConstraintFromID(id)
	})
}
// A Boolean value indicating whether the constraints impacting the anchor
// specify its location ambiguously.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor/hasAmbiguousLayout
func (l NSLayoutAnchor) HasAmbiguousLayout() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("hasAmbiguousLayout"))
	return rv
}
// The name assigned to the anchor for debugging purposes.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor/name
func (l NSLayoutAnchor) Name() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
// The layout item used to calculate the anchor’s position.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor/item
func (l NSLayoutAnchor) Item() objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("item"))
	return objectivec.Object{ID: rv}
}
// A layout anchor representing the bottom edge of the view’s frame.
//
// See: https://developer.apple.com/documentation/appkit/nsview/bottomanchor
func (l NSLayoutAnchor) BottomAnchor() INSLayoutYAxisAnchor {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("bottomAnchor"))
	return NSLayoutYAxisAnchorFromID(objc.ID(rv))
}
func (l NSLayoutAnchor) SetBottomAnchor(value INSLayoutYAxisAnchor) {
	objc.Send[struct{}](l.ID, objc.Sel("setBottomAnchor:"), value)
}
// A layout anchor representing the leading edge of the view’s frame.
//
// See: https://developer.apple.com/documentation/appkit/nsview/leadinganchor
func (l NSLayoutAnchor) LeadingAnchor() INSLayoutXAxisAnchor {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("leadingAnchor"))
	return NSLayoutXAxisAnchorFromID(objc.ID(rv))
}
func (l NSLayoutAnchor) SetLeadingAnchor(value INSLayoutXAxisAnchor) {
	objc.Send[struct{}](l.ID, objc.Sel("setLeadingAnchor:"), value)
}
// A layout anchor representing the left edge of the view’s frame.
//
// See: https://developer.apple.com/documentation/appkit/nsview/leftanchor
func (l NSLayoutAnchor) LeftAnchor() INSLayoutXAxisAnchor {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("leftAnchor"))
	return NSLayoutXAxisAnchorFromID(objc.ID(rv))
}
func (l NSLayoutAnchor) SetLeftAnchor(value INSLayoutXAxisAnchor) {
	objc.Send[struct{}](l.ID, objc.Sel("setLeftAnchor:"), value)
}

