// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSLayoutConstraint] class.
var (
	_NSLayoutConstraintClass     NSLayoutConstraintClass
	_NSLayoutConstraintClassOnce sync.Once
)

func getNSLayoutConstraintClass() NSLayoutConstraintClass {
	_NSLayoutConstraintClassOnce.Do(func() {
		_NSLayoutConstraintClass = NSLayoutConstraintClass{class: objc.GetClass("NSLayoutConstraint")}
	})
	return _NSLayoutConstraintClass
}

// GetNSLayoutConstraintClass returns the class object for NSLayoutConstraint.
func GetNSLayoutConstraintClass() NSLayoutConstraintClass {
	return getNSLayoutConstraintClass()
}

type NSLayoutConstraintClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSLayoutConstraintClass) Alloc() NSLayoutConstraint {
	rv := objc.Send[NSLayoutConstraint](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The relationship between two user interface objects that must be satisfied
// by the constraint-based layout system.
//
// # Overview
// 
// Each constraint is a linear equation with the following format:
// 
// In this equation, `attribute1` and `attribute2` are the variables that Auto
// Layout can adjust when solving these constraints. The other values are
// defined when you create the constraint. For example, If you’re defining
// the relative position of two buttons, you might say “the leading edge of
// the second button should be 8 points after the trailing edge of the first
// button.” The linear equation for this relationship is shown below:
// 
// Auto Layout then modifies the values of the specified leading and trailing
// edges until both sides of the equation are equal. Note that Auto Layout
// does not simply assign the value of the right side of this equation to the
// left side. Instead, the system can modify either attribute or both
// attributes as needed to solve for this constraint.
// 
// The fact that constraints are equations (and not assignment operators)
// means that you can switch the order of the items in the equation as needed
// to more clearly express the desired relationship. However, if you switch
// the order, you must also invert the multiplier and constant. For example,
// the following two equations produce identical constraints:
// 
// A valid layout is defined as a set constraints with one and only one
// possible solution. Valid layouts are also referred to as a nonambiguous,
// nonconflicting layouts. Constraints with more than one solution are
// ambiguous. Constraints with no valid solutions are conflicting. For more
// information on resolving ambiguous and conflicting constraints, see [Types
// of Errors] in [Auto Layout Guide].
// 
// Additionally, constraints are not limited to equality relationships. They
// can also use greater than or equal to (>=) or less than or equal to (<=) to
// describe the relationship between the two attributes. Constraints also have
// priorities between 1 and 1,000. Constraints with a priority of 1,000 are
// required. All priorities less than 1,000 are optional. By default, all
// constraints are required (priority = 1,000).
// 
// After solving for the required constraints, Auto Layout tries to solve all
// the optional constraints in priority order from highest to lowest. If it
// cannot solve for an optional constraint, it tries to come as close as
// possible to the desired result, and then moves on to the next constraint.
// 
// This combination of inequalities, equalities, and priorities gives you a
// great amount of flexibility and power. By combining multiple constraints,
// you can define layouts that dynamically adapt as the size and location of
// the elements in your user interface change. For some example layouts, see
// [Stack Views] in [Auto Layout Guide].
//
// [Auto Layout Guide]: https://developer.apple.com/library/archive/documentation/UserExperience/Conceptual/AutolayoutPG/index.html#//apple_ref/doc/uid/TP40010853
// [Stack Views]: https://developer.apple.com/library/archive/documentation/UserExperience/Conceptual/AutolayoutPG/LayoutUsingStackViews.html#//apple_ref/doc/uid/TP40010853-CH11
// [Types of Errors]: https://developer.apple.com/library/archive/documentation/UserExperience/Conceptual/AutolayoutPG/TypesofErrors.html#//apple_ref/doc/uid/TP40010853-CH17
//
// # Activating and deactivating constraints
//
//   - [NSLayoutConstraint.Active]: The active state of the constraint.
//   - [NSLayoutConstraint.SetActive]
//
// # Accessing constraint data
//
//   - [NSLayoutConstraint.FirstItem]: The first object participating in the constraint.
//   - [NSLayoutConstraint.FirstAttribute]: The attribute of the first object participating in the constraint.
//   - [NSLayoutConstraint.Relation]: The relation between the two attributes in the constraint.
//   - [NSLayoutConstraint.SecondItem]: The second object participating in the constraint.
//   - [NSLayoutConstraint.SecondAttribute]: The attribute of the second object participating in the constraint.
//   - [NSLayoutConstraint.Multiplier]: The multiplier applied to the second attribute participating in the constraint.
//   - [NSLayoutConstraint.Constant]: The constant added to the multiplied second attribute participating in the constraint.
//   - [NSLayoutConstraint.SetConstant]
//   - [NSLayoutConstraint.FirstAnchor]: The first anchor that defines the constraint.
//   - [NSLayoutConstraint.SecondAnchor]: The second anchor that defines the constraint.
//
// # Getting the layout priority
//
//   - [NSLayoutConstraint.Priority]: The priority of the constraint.
//   - [NSLayoutConstraint.SetPriority]
//
// # Identifying a constraint
//
//   - [NSLayoutConstraint.Identifier]: The name that identifies the constraint.
//   - [NSLayoutConstraint.SetIdentifier]
//
// # Controlling constraint archiving
//
//   - [NSLayoutConstraint.ShouldBeArchived]: A Boolean value that determines whether the constraint should be archived by its owning view.
//   - [NSLayoutConstraint.SetShouldBeArchived]
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint
type NSLayoutConstraint struct {
	objectivec.Object
}

// NSLayoutConstraintFromID constructs a [NSLayoutConstraint] from an objc.ID.
//
// The relationship between two user interface objects that must be satisfied
// by the constraint-based layout system.
func NSLayoutConstraintFromID(id objc.ID) NSLayoutConstraint {
	return NSLayoutConstraint{objectivec.Object{ID: id}}
}
// NOTE: NSLayoutConstraint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSLayoutConstraint] class.
//
// # Activating and deactivating constraints
//
//   - [INSLayoutConstraint.Active]: The active state of the constraint.
//   - [INSLayoutConstraint.SetActive]
//
// # Accessing constraint data
//
//   - [INSLayoutConstraint.FirstItem]: The first object participating in the constraint.
//   - [INSLayoutConstraint.FirstAttribute]: The attribute of the first object participating in the constraint.
//   - [INSLayoutConstraint.Relation]: The relation between the two attributes in the constraint.
//   - [INSLayoutConstraint.SecondItem]: The second object participating in the constraint.
//   - [INSLayoutConstraint.SecondAttribute]: The attribute of the second object participating in the constraint.
//   - [INSLayoutConstraint.Multiplier]: The multiplier applied to the second attribute participating in the constraint.
//   - [INSLayoutConstraint.Constant]: The constant added to the multiplied second attribute participating in the constraint.
//   - [INSLayoutConstraint.SetConstant]
//   - [INSLayoutConstraint.FirstAnchor]: The first anchor that defines the constraint.
//   - [INSLayoutConstraint.SecondAnchor]: The second anchor that defines the constraint.
//
// # Getting the layout priority
//
//   - [INSLayoutConstraint.Priority]: The priority of the constraint.
//   - [INSLayoutConstraint.SetPriority]
//
// # Identifying a constraint
//
//   - [INSLayoutConstraint.Identifier]: The name that identifies the constraint.
//   - [INSLayoutConstraint.SetIdentifier]
//
// # Controlling constraint archiving
//
//   - [INSLayoutConstraint.ShouldBeArchived]: A Boolean value that determines whether the constraint should be archived by its owning view.
//   - [INSLayoutConstraint.SetShouldBeArchived]
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint
type INSLayoutConstraint interface {
	objectivec.IObject

	// Topic: Activating and deactivating constraints

	// The active state of the constraint.
	Active() bool
	SetActive(value bool)

	// Topic: Accessing constraint data

	// The first object participating in the constraint.
	FirstItem() objectivec.IObject
	// The attribute of the first object participating in the constraint.
	FirstAttribute() NSLayoutAttribute
	// The relation between the two attributes in the constraint.
	Relation() NSLayoutRelation
	// The second object participating in the constraint.
	SecondItem() objectivec.IObject
	// The attribute of the second object participating in the constraint.
	SecondAttribute() NSLayoutAttribute
	// The multiplier applied to the second attribute participating in the constraint.
	Multiplier() float64
	// The constant added to the multiplied second attribute participating in the constraint.
	Constant() float64
	SetConstant(value float64)
	// The first anchor that defines the constraint.
	FirstAnchor() INSLayoutAnchor
	// The second anchor that defines the constraint.
	SecondAnchor() INSLayoutAnchor

	// Topic: Getting the layout priority

	// The priority of the constraint.
	Priority() NSLayoutPriority
	SetPriority(value NSLayoutPriority)

	// Topic: Identifying a constraint

	// The name that identifies the constraint.
	Identifier() string
	SetIdentifier(value string)

	// Topic: Controlling constraint archiving

	// A Boolean value that determines whether the constraint should be archived by its owning view.
	ShouldBeArchived() bool
	SetShouldBeArchived(value bool)

	// Returns the animation that should be performed for the specified key.
	AnimationForKey(key NSAnimatablePropertyKey) objectivec.IObject
	// Returns a proxy object for the receiver that can be used to initiate implied animation for property changes.
	Animator() INSLayoutConstraint
}

// Init initializes the instance.
func (l NSLayoutConstraint) Init() NSLayoutConstraint {
	rv := objc.Send[NSLayoutConstraint](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l NSLayoutConstraint) Autorelease() NSLayoutConstraint {
	rv := objc.Send[NSLayoutConstraint](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSLayoutConstraint creates a new NSLayoutConstraint instance.
func NewNSLayoutConstraint() NSLayoutConstraint {
	class := getNSLayoutConstraintClass()
	rv := objc.Send[NSLayoutConstraint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a constraint that defines the relationship between the specified
// attributes of the given views.
//
// view1: The view for the left side of the constraint.
//
// attr1: The attribute of the view for the left side of the constraint.
//
// relation: The relationship between the left side of the constraint and the right side
// of the constraint.
//
// view2: The view for the right side of the constraint.
//
// attr2: The attribute of the view for the right side of the constraint.
//
// multiplier: The constant multiplied with the attribute on the right side of the
// constraint as part of getting the modified attribute.
//
// c: The constant added to the multiplied attribute value on the right side of
// the constraint to yield the final modified attribute.
//
// # Return Value
// 
// A constraint object relating the two provided views with the specified
// relation, attributes, multiplier, and constant.
//
// # Discussion
// 
// Constraints represent linear equations of the form `view1.Attr1()
// multiplier × view2.Attr2() + c`. If the constraint you wish to express
// does not have a second view and attribute, use `nil` and
// [NSLayoutConstraint.Attribute.notAnAttribute].
//
// [NSLayoutConstraint.Attribute.notAnAttribute]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute/notAnAttribute
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/init(item:attribute:relatedBy:toItem:attribute:multiplier:constant:)
func NewLayoutConstraintWithItemAttributeRelatedByToItemAttributeMultiplierConstant(view1 objectivec.IObject, attr1 NSLayoutAttribute, relation NSLayoutRelation, view2 objectivec.IObject, attr2 NSLayoutAttribute, multiplier float64, c float64) NSLayoutConstraint {
	rv := objc.Send[objc.ID](objc.ID(getNSLayoutConstraintClass().class), objc.Sel("constraintWithItem:attribute:relatedBy:toItem:attribute:multiplier:constant:"), view1, attr1, relation, view2, attr2, multiplier, c)
	return NSLayoutConstraintFromID(rv)
}

// Returns the animation that should be performed for the specified key.
//
// key: The action name or property specified as a string.
//
// # Return Value
// 
// The animation to perform. A subclass of [CAAnimation].
//
// [CAAnimation]: https://developer.apple.com/documentation/QuartzCore/CAAnimation
//
// # Discussion
// 
// When the action specified by `key` is triggered for an object, this method
// is consulted to find the animation, if any, that should be performed in
// response.
// 
// Like its Core Animation [CALayer] counterpart, [action(forKey:)], this
// method is a funnel point that defines the order in which the search for an
// animation proceeds.It first checks the receiver’s Getting the Animator
// Proxy dictionary for a value matching `key`, then falls back to [Animator]
// for the receiver’s class.
// 
// Subclasses should not typically need to override this method.
//
// [CALayer]: https://developer.apple.com/documentation/QuartzCore/CALayer
// [action(forKey:)]: https://developer.apple.com/documentation/QuartzCore/CALayer/action(forKey:)
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/animation(forKey:)
func (l NSLayoutConstraint) AnimationForKey(key NSAnimatablePropertyKey) objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("animationForKey:"), objc.String(string(key)))
	return objectivec.Object{ID: rv}
}
// Returns a proxy object for the receiver that can be used to initiate
// implied animation for property changes.
//
// # Return Value
// 
// Returns a proxy object for the receiver that can initiate implied
// animations in response to property changes.
//
// # Discussion
// 
// The animator proxy object should be treated as if it was the receiver
// itself, and may be passed to any code that accepts the receiver as a
// parameter.
// 
// Sending key-value coding compliant “set” messages to the proxy will
// trigger animation for automatically animated properties of its target
// object, if the active [NSAnimationContext] in the current thread has a
// duration value greater than zero, and an animation for the property key is
// found by the [NSAnimatablePropertyContainer] search mechanism.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/animator()
func (l NSLayoutConstraint) Animator() INSLayoutConstraint {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("animator"))
	return NSLayoutConstraintFromID(rv)
}

// Creates constraints described by an ASCII art-like visual format string.
//
// format: The format specification for the constraints. For more information, see
// [Auto Layout Cookbook] in [Auto Layout Guide].
// //
// [Auto Layout Cookbook]: https://developer.apple.com/library/archive/documentation/UserExperience/Conceptual/AutolayoutPG/LayoutUsingStackViews.html#//apple_ref/doc/uid/TP40010853-CH3
// [Auto Layout Guide]: https://developer.apple.com/library/archive/documentation/UserExperience/Conceptual/AutolayoutPG/index.html#//apple_ref/doc/uid/TP40010853
//
// opts: Options describing the attribute and the direction of layout for all
// objects in the visual format string.
//
// metrics: A dictionary of constants that appear in the visual format string. The
// dictionary’s keys must be the string values used in the visual format
// string. Their values must be [NSNumber] objects.
// //
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
//
// views: A dictionary of views that appear in the visual format string. The keys
// must be the string values used in the visual format string, and the values
// must be the view objects.
//
// # Return Value
// 
// An array of constraints that, combined, express the constraints between the
// provided views and their parent view as described by the visual format
// string. The constraints are returned in the same order they were specified
// in the visual format string.
//
// # Discussion
// 
// The language used for the visual format string is described in [Auto Layout
// Cookbook] in [Auto Layout Guide].
//
// [Auto Layout Cookbook]: https://developer.apple.com/library/archive/documentation/UserExperience/Conceptual/AutolayoutPG/LayoutUsingStackViews.html#//apple_ref/doc/uid/TP40010853-CH3
// [Auto Layout Guide]: https://developer.apple.com/library/archive/documentation/UserExperience/Conceptual/AutolayoutPG/index.html#//apple_ref/doc/uid/TP40010853
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/constraints(withVisualFormat:options:metrics:views:)
func (_NSLayoutConstraintClass NSLayoutConstraintClass) ConstraintsWithVisualFormatOptionsMetricsViews(format string, opts NSLayoutFormatOptions, metrics foundation.INSDictionary, views foundation.INSDictionary) []NSLayoutConstraint {
	rv := objc.Send[[]objc.ID](objc.ID(_NSLayoutConstraintClass.class), objc.Sel("constraintsWithVisualFormat:options:metrics:views:"), objc.String(format), opts, metrics, views)
	return objc.ConvertSlice(rv, func(id objc.ID) NSLayoutConstraint {
		return NSLayoutConstraintFromID(id)
	})
}
// Activates each constraint in the specified array.
//
// constraints: An array of constraints to activate.
//
// # Discussion
// 
// This convenience method provides an easy way to activate a set of
// constraints with one call. The effect of this method is the same as setting
// the [Active] property of each constraint to [true]. Typically, using this
// method is more efficient than activating each constraint individually.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/activate(_:)
func (_NSLayoutConstraintClass NSLayoutConstraintClass) ActivateConstraints(constraints []NSLayoutConstraint) {
	objc.Send[objc.ID](objc.ID(_NSLayoutConstraintClass.class), objc.Sel("activateConstraints:"), objectivec.IObjectSliceToNSArray(constraints))
}
// Deactivates each constraint in the specified array.
//
// constraints: An array of constraints to deactivate.
//
// # Discussion
// 
// This is a convenience method that provides an easy way to deactivate a set
// of constraints with one call. The effect of this method is the same as
// setting the [Active] property of each constraint to [false]. Typically,
// using this method is more efficient than deactivating each constraint
// individually.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/deactivate(_:)
func (_NSLayoutConstraintClass NSLayoutConstraintClass) DeactivateConstraints(constraints []NSLayoutConstraint) {
	objc.Send[objc.ID](objc.ID(_NSLayoutConstraintClass.class), objc.Sel("deactivateConstraints:"), objectivec.IObjectSliceToNSArray(constraints))
}
// Returns the default animation that should be performed for the specified
// key.
//
// key: The action name or property specified as a string.
//
// # Return Value
// 
// The animation to perform. A subclass of [CAAnimation].
//
// [CAAnimation]: https://developer.apple.com/documentation/QuartzCore/CAAnimation
//
// # Discussion
// 
// The [NSAnimatablePropertyContainer] method consults this class method when
// its search of the receivers Getting the Animator Proxy dictionary fails to
// return an animation for `key`.
// 
// An animatable property container should implement this method to return a
// default animation to be performed for each key that it wants to make
// auto-animatable, where `key` usually references a property of the receiver,
// but can also specify a special animation trigger
// ([NSAnimationTriggerOrderIn] or [NSAnimationTriggerOrderOut]).
// 
// A developer implementing a custom view subclass, can enable automatic
// animation for properties by overriding this method, and having it return
// the desired default [CAAnimation] subclass to use for each of the property
// keys of interest. The override should defer to super for any keys it
// doesn’t specifically handle, facilitating inheritance of default
// animation specifications. The following is an example of such an
// implementation.
//
// [CAAnimation]: https://developer.apple.com/documentation/QuartzCore/CAAnimation
// [NSAnimationTriggerOrderIn]: https://developer.apple.com/documentation/AppKit/NSAnimationTriggerOrderIn
// [NSAnimationTriggerOrderOut]: https://developer.apple.com/documentation/AppKit/NSAnimationTriggerOrderOut
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/defaultAnimation(forKey:)
func (_NSLayoutConstraintClass NSLayoutConstraintClass) DefaultAnimationForKey(key NSAnimatablePropertyKey) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NSLayoutConstraintClass.class), objc.Sel("defaultAnimationForKey:"), objc.String(string(key)))
	return objectivec.Object{ID: rv}
}

// The active state of the constraint.
//
// # Discussion
// 
// You can activate or deactivate a constraint by changing this property. Note
// that only active constraints affect the calculated layout. If you try to
// activate a constraint whose items have no common ancestor, an exception is
// thrown. For newly created constraints, the [Active] property is [false] by
// default.
// 
// Activating or deactivating the constraint calls [AddConstraint] and
// [RemoveConstraint] on the view that is the closest common ancestor of the
// items managed by this constraint. Use this property instead of calling
// [AddConstraint] or [RemoveConstraint] directly.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/isActive
func (l NSLayoutConstraint) Active() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("isActive"))
	return rv
}
func (l NSLayoutConstraint) SetActive(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setActive:"), value)
}
// The first object participating in the constraint.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/firstItem
func (l NSLayoutConstraint) FirstItem() objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("firstItem"))
	return objectivec.Object{ID: rv}
}
// The attribute of the first object participating in the constraint.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/firstAttribute
func (l NSLayoutConstraint) FirstAttribute() NSLayoutAttribute {
	rv := objc.Send[NSLayoutAttribute](l.ID, objc.Sel("firstAttribute"))
	return NSLayoutAttribute(rv)
}
// The relation between the two attributes in the constraint.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/relation-swift.property
func (l NSLayoutConstraint) Relation() NSLayoutRelation {
	rv := objc.Send[NSLayoutRelation](l.ID, objc.Sel("relation"))
	return NSLayoutRelation(rv)
}
// The second object participating in the constraint.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/secondItem
func (l NSLayoutConstraint) SecondItem() objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("secondItem"))
	return objectivec.Object{ID: rv}
}
// The attribute of the second object participating in the constraint.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/secondAttribute
func (l NSLayoutConstraint) SecondAttribute() NSLayoutAttribute {
	rv := objc.Send[NSLayoutAttribute](l.ID, objc.Sel("secondAttribute"))
	return NSLayoutAttribute(rv)
}
// The multiplier applied to the second attribute participating in the
// constraint.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/multiplier
func (l NSLayoutConstraint) Multiplier() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("multiplier"))
	return rv
}
// The constant added to the multiplied second attribute participating in the
// constraint.
//
// # Discussion
// 
// Unlike the other properties, the constant can be modified after constraint
// creation. Setting the constant on an existing constraint performs much
// better than removing the constraint and adding a new one that’s exactly
// like the old except that it has a different constant.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/constant
func (l NSLayoutConstraint) Constant() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("constant"))
	return rv
}
func (l NSLayoutConstraint) SetConstant(value float64) {
	objc.Send[struct{}](l.ID, objc.Sel("setConstant:"), value)
}
// The first anchor that defines the constraint.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/firstAnchor
func (l NSLayoutConstraint) FirstAnchor() INSLayoutAnchor {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("firstAnchor"))
	return NSLayoutAnchorFromID(objc.ID(rv))
}
// The second anchor that defines the constraint.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/secondAnchor
func (l NSLayoutConstraint) SecondAnchor() INSLayoutAnchor {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("secondAnchor"))
	return NSLayoutAnchorFromID(objc.ID(rv))
}
// The priority of the constraint.
//
// # Discussion
// 
// By default, all constraints are required; this property is set to
// [required] in macOS or [UILayoutPriorityRequired] in iOS.
// 
// If a constraint’s priority level is less than [required] in macOS or
// [UILayoutPriorityRequired] in iOS, then it is optional. Higher priority
// constraints are satisfied before lower priority constraints; however,
// optional constraint satisfaction is not all or nothing. If a constraint `a
// == b` is optional, the constraint-based layout system will attempt to
// minimize `abs(a-b)`.
// 
// Priorities may not change from nonrequired to required, or from required to
// nonrequired. An exception will be thrown if a priority of [required] in
// macOS or [UILayoutPriorityRequired] in iOS is changed to a lower priority,
// or if a lower priority is changed to a required priority after the
// constraints is added to a view. Changing from one optional priority to
// another optional priority is allowed even after the constraint is installed
// on a view.
// 
// Priorities must be greater than 0 and less than or equal to [required] in
// macOS or [UILayoutPriorityRequired] in iOS.
//
// [required]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Priority-swift.struct/required
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/priority-swift.property
func (l NSLayoutConstraint) Priority() NSLayoutPriority {
	rv := objc.Send[NSLayoutPriority](l.ID, objc.Sel("priority"))
	return NSLayoutPriority(rv)
}
func (l NSLayoutConstraint) SetPriority(value NSLayoutPriority) {
	objc.Send[struct{}](l.ID, objc.Sel("setPriority:"), value)
}
// The name that identifies the constraint.
//
// # Discussion
// 
// A constraint’s identifier is available in its description. Identifiers
// that start with [NS] are reserved by the system.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/identifier
func (l NSLayoutConstraint) Identifier() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}
func (l NSLayoutConstraint) SetIdentifier(value string) {
	objc.Send[struct{}](l.ID, objc.Sel("setIdentifier:"), objc.String(value))
}
// A Boolean value that determines whether the constraint should be archived
// by its owning view.
//
// # Discussion
// 
// When a view is archived, it archives some but not all constraints in its
// [constraints] array. The value of [ShouldBeArchived] informs the view if a
// particular constraint should be archived by the view.
// 
// If a constraint is created at runtime in response to the state of the
// object, it isn’t appropriate to archive the constraint. Instead you
// archive the state that gives rise to the constraint. The default value for
// this property is [false].
//
// [constraints]: https://developer.apple.com/documentation/AppKit/NSView/constraints
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/shouldBeArchived
func (l NSLayoutConstraint) ShouldBeArchived() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("shouldBeArchived"))
	return rv
}
func (l NSLayoutConstraint) SetShouldBeArchived(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setShouldBeArchived:"), value)
}
// Sets the option dictionary that maps event trigger keys to animation
// objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/animations
func (l NSLayoutConstraint) Animations() foundation.INSDictionary {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("animations"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (l NSLayoutConstraint) SetAnimations(value foundation.INSDictionary) {
	objc.Send[struct{}](l.ID, objc.Sel("setAnimations:"), value)
}

