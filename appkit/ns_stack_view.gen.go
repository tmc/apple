// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSStackView] class.
var (
	_NSStackViewClass     NSStackViewClass
	_NSStackViewClassOnce sync.Once
)

func getNSStackViewClass() NSStackViewClass {
	_NSStackViewClassOnce.Do(func() {
		_NSStackViewClass = NSStackViewClass{class: objc.GetClass("NSStackView")}
	})
	return _NSStackViewClass
}

// GetNSStackViewClass returns the class object for NSStackView.
func GetNSStackViewClass() NSStackViewClass {
	return getNSStackViewClass()
}

type NSStackViewClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSStackViewClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSStackViewClass) Alloc() NSStackView {
	rv := objc.Send[NSStackView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A view that arranges an array of views horizontally or vertically and
// updates their placement and sizing when the window size changes.
//
// # Overview
//
// A stack view employs Auto Layout (the system’s constraint-based layout
// feature) to arrange and align an array of views according to your
// specification. To use a stack view effectively, you need to understand the
// basics of Auto Layout constraints as described in [Auto Layout Guide].
//
// # Basic Features of Stack Views
//
// A stack view supports vertical and horizontal layouts and interacts
// dynamically with window resizing and Cocoa animations. You can easily
// reconfigure the contents of a stack view at runtime. That is, after you
// create and configure a stack view in Interface Builder, you can add or
// remove views dynamically without explicitly working with layout
// constraints. For example, if you configure a stack view with three
// checkboxes and dynamically add a fourth, the stack view automatically adds
// constraints as needed, according to the stack view’s configuration. The
// new checkbox gains dynamic layout configuration from the stack view.
//
// Stack views are nestable: a stack view is a valid element in the [NSStackView.Views]
// array of another stack view.
//
// For more information on [NSStackView], see [Organize Your User Interface
// with a Stack View].
//
// # Layout Direction and Gravity Areas
//
// A stack view has three so-called that each identify a section of the stack
// view’s layout. A horizontal stack view, which is the default type, has a
// leading, a center, and a trailing gravity area. The ordering of these areas
// depends on the value of the stack view’s [NSStackView.UserInterfaceLayoutDirection]
// property (inherited from the [NSView] class). In a left to right language,
// the leading gravity area in a horizontal stack view is on the left. To
// enforce a left to right layout independently of language, explicitly set
// the layout direction by calling the inherited
// [NSStackView.UserInterfaceLayoutDirection] method on your stack view instance.
//
// To specify vertical layout, use the [NSStackView.Orientation] property and the
// [NSUserInterfaceLayoutOrientationVertical] constant from the
// [NSUserInterfaceLayoutOrientation] enumeration. In a vertical stack view,
// the gravity areas always are top, center, and bottom.
//
// # View Detachment and Hiding
//
// A stack view can automatically detach and reattach its views in response to
// layout changes, such as window resizing performed by the user, or
// resizing/repositioning of another view in the same view hierarchy. A view
// in a detached state is not present in the stack view’s view hierarchy,
// but it still consumes memory. A view that is hidden, but not detached,
// remains part of the view hierarchy and continues to participate in Auto
// Layout, but it is not visible and doesn’t receive input events.
//
// To allow views to detach, set the so-called for a stack view to a value
// lower than its default of [NSStackView.Required]. See the
// [NSStackView.SetClippingResistancePriorityForOrientation] method.
//
// You can influence which views detach first (and reattach last). Do this by
// setting the so-called for each view whose detachment order you want to
// specify. A view with a lower visibility priority detaches before one with a
// higher priority, and reattaches after it. See the
// [NSStackViewVisibilityPriority] enumeration and the
// [NSStackView.SetVisibilityPriorityForView] method.
//
// To explicitly detach a view from a stack view, call the
// [NSStackView.SetVisibilityPriorityForView] method with a value of [notVisible]. To
// explicitly reattach a view to a stack view, call the same method with a
// value of [mustHold]. If you hide a view that belongs to a stack view (by
// setting the view’s [Hidden] property to true), the view detaches from the
// stack view by default. Use the [NSStackView.DetachesHiddenViews] property to change the
// default behavior.
//
// The system calls a stack view delegate method when a view is about to be
// detached and when a view has been reattached, giving you the opportunity to
// run code at those times. See [NSStackViewDelegate].
//
// # Responding to Stack-Related Changes
//
//   - [NSStackView.Delegate]: The delegate object for the stack view.
//   - [NSStackView.SetDelegate]
//
// # Managing Views in Gravity Areas
//
//   - [NSStackView.AddViewInGravity]: Adds a view to the end of the stack view gravity area.
//   - [NSStackView.InsertViewAtIndexInGravity]: Adds a view to a stack view gravity area at a specified index position.
//   - [NSStackView.SetViewsInGravity]: Specifies an array of views for a specified gravity area in the stack view, replacing any previous views in that area.
//   - [NSStackView.RemoveView]: Removes a specified view from the stack view.
//
// # Managing the Arranged Subviews
//
//   - [NSStackView.AddArrangedSubview]: Adds the specified view to the end of the arranged subviews list.
//   - [NSStackView.InsertArrangedSubviewAtIndex]: Adds the provided view to the array of arranged subviews at the specified index.
//   - [NSStackView.RemoveArrangedSubview]: Removes the provided view from the stack’s array of arranged subviews.
//   - [NSStackView.ArrangedSubviews]: The array of views arranged by the stack view.
//
// # Inspecting a Stack View
//
//   - [NSStackView.Views]: The array of views owned by the stack view.
//   - [NSStackView.ViewsInGravity]: Returns the array of views in the specified gravity area in the stack view.
//   - [NSStackView.DetachedViews]: An array that contains the detached views from all the stack view’s gravity areas.
//   - [NSStackView.ClippingResistancePriorityForOrientation]: Returns the Auto Layout priority for resisting clipping of views in the stack view when Auto Layout attempts to reduce the stack view’s size.
//   - [NSStackView.HuggingPriorityForOrientation]: Returns the Auto Layout priority for the stack view to minimize its size to fit its contained views as closely as possible, for a specified user interface axis.
//
// # Configuring the Stack View Layout
//
//   - [NSStackView.Orientation]: The horizontal or vertical layout direction of the stack view.
//   - [NSStackView.SetOrientation]
//   - [NSStackView.Alignment]: The view alignment within the stack view.
//   - [NSStackView.SetAlignment]
//   - [NSStackView.Spacing]: The minimum spacing, in points, between adjacent views in the stack view.
//   - [NSStackView.SetSpacing]
//   - [NSStackView.EdgeInsets]: The geometric padding, in points, inside the stack view, surrounding its views.
//   - [NSStackView.SetEdgeInsets]
//   - [NSStackView.Distribution]: The spacing and sizing distribution of stacked views along the primary axis.
//   - [NSStackView.SetDistribution]
//
// # Configuring Views in a Stack View
//
//   - [NSStackView.CustomSpacingAfterView]: Returns the custom spacing, in points, between a specified view in the stack view and the view that follows it.
//   - [NSStackView.SetCustomSpacingAfterView]: Specifies the custom spacing, in points, between a specified view and the view that follows it in the stack view.
//   - [NSStackView.VisibilityPriorityForView]: Returns the visibility priority for a specified view in the stack view.
//   - [NSStackView.SetVisibilityPriorityForView]: Sets the Auto Layout priority for a view to remain attached to the stack view when Auto Layout reduces the stack view’s size.
//
// # Configuring Dynamic Behavior for a Stack View
//
//   - [NSStackView.DetachesHiddenViews]: A Boolean value that indicates whether the stack view removes hidden views from its view hierarchy.
//   - [NSStackView.SetDetachesHiddenViews]
//   - [NSStackView.SetClippingResistancePriorityForOrientation]: Sets the Auto Layout priority for resisting clipping of views in the stack view when Auto Layout attempts to reduce the stack view’s size.
//   - [NSStackView.SetHuggingPriorityForOrientation]: Sets the Auto Layout priority for the stack view to minimize its size, for a specified user interface axis.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView
//
// [Auto Layout Guide]: https://developer.apple.com/library/archive/documentation/UserExperience/Conceptual/AutolayoutPG/index.html#//apple_ref/doc/uid/TP40010853
// [NSUserInterfaceLayoutOrientation]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutOrientation
// [Organize Your User Interface with a Stack View]: https://developer.apple.com/documentation/AppKit/organize-your-user-interface-with-a-stack-view
// [mustHold]: https://developer.apple.com/documentation/AppKit/NSStackView/VisibilityPriority/mustHold
// [notVisible]: https://developer.apple.com/documentation/AppKit/NSStackView/VisibilityPriority/notVisible
type NSStackView struct {
	NSView
}

// NSStackViewFromID constructs a [NSStackView] from an objc.ID.
//
// A view that arranges an array of views horizontally or vertically and
// updates their placement and sizing when the window size changes.
func NSStackViewFromID(id objc.ID) NSStackView {
	return NSStackView{NSView: NSViewFromID(id)}
}

// NOTE: NSStackView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSStackView] class.
//
// # Responding to Stack-Related Changes
//
//   - [INSStackView.Delegate]: The delegate object for the stack view.
//   - [INSStackView.SetDelegate]
//
// # Managing Views in Gravity Areas
//
//   - [INSStackView.AddViewInGravity]: Adds a view to the end of the stack view gravity area.
//   - [INSStackView.InsertViewAtIndexInGravity]: Adds a view to a stack view gravity area at a specified index position.
//   - [INSStackView.SetViewsInGravity]: Specifies an array of views for a specified gravity area in the stack view, replacing any previous views in that area.
//   - [INSStackView.RemoveView]: Removes a specified view from the stack view.
//
// # Managing the Arranged Subviews
//
//   - [INSStackView.AddArrangedSubview]: Adds the specified view to the end of the arranged subviews list.
//   - [INSStackView.InsertArrangedSubviewAtIndex]: Adds the provided view to the array of arranged subviews at the specified index.
//   - [INSStackView.RemoveArrangedSubview]: Removes the provided view from the stack’s array of arranged subviews.
//   - [INSStackView.ArrangedSubviews]: The array of views arranged by the stack view.
//
// # Inspecting a Stack View
//
//   - [INSStackView.Views]: The array of views owned by the stack view.
//   - [INSStackView.ViewsInGravity]: Returns the array of views in the specified gravity area in the stack view.
//   - [INSStackView.DetachedViews]: An array that contains the detached views from all the stack view’s gravity areas.
//   - [INSStackView.ClippingResistancePriorityForOrientation]: Returns the Auto Layout priority for resisting clipping of views in the stack view when Auto Layout attempts to reduce the stack view’s size.
//   - [INSStackView.HuggingPriorityForOrientation]: Returns the Auto Layout priority for the stack view to minimize its size to fit its contained views as closely as possible, for a specified user interface axis.
//
// # Configuring the Stack View Layout
//
//   - [INSStackView.Orientation]: The horizontal or vertical layout direction of the stack view.
//   - [INSStackView.SetOrientation]
//   - [INSStackView.Alignment]: The view alignment within the stack view.
//   - [INSStackView.SetAlignment]
//   - [INSStackView.Spacing]: The minimum spacing, in points, between adjacent views in the stack view.
//   - [INSStackView.SetSpacing]
//   - [INSStackView.EdgeInsets]: The geometric padding, in points, inside the stack view, surrounding its views.
//   - [INSStackView.SetEdgeInsets]
//   - [INSStackView.Distribution]: The spacing and sizing distribution of stacked views along the primary axis.
//   - [INSStackView.SetDistribution]
//
// # Configuring Views in a Stack View
//
//   - [INSStackView.CustomSpacingAfterView]: Returns the custom spacing, in points, between a specified view in the stack view and the view that follows it.
//   - [INSStackView.SetCustomSpacingAfterView]: Specifies the custom spacing, in points, between a specified view and the view that follows it in the stack view.
//   - [INSStackView.VisibilityPriorityForView]: Returns the visibility priority for a specified view in the stack view.
//   - [INSStackView.SetVisibilityPriorityForView]: Sets the Auto Layout priority for a view to remain attached to the stack view when Auto Layout reduces the stack view’s size.
//
// # Configuring Dynamic Behavior for a Stack View
//
//   - [INSStackView.DetachesHiddenViews]: A Boolean value that indicates whether the stack view removes hidden views from its view hierarchy.
//   - [INSStackView.SetDetachesHiddenViews]
//   - [INSStackView.SetClippingResistancePriorityForOrientation]: Sets the Auto Layout priority for resisting clipping of views in the stack view when Auto Layout attempts to reduce the stack view’s size.
//   - [INSStackView.SetHuggingPriorityForOrientation]: Sets the Auto Layout priority for the stack view to minimize its size, for a specified user interface axis.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView
type INSStackView interface {
	INSView

	// Topic: Responding to Stack-Related Changes

	// The delegate object for the stack view.
	Delegate() NSStackViewDelegate
	SetDelegate(value NSStackViewDelegate)

	// Topic: Managing Views in Gravity Areas

	// Adds a view to the end of the stack view gravity area.
	AddViewInGravity(view INSView, gravity NSStackViewGravity)
	// Adds a view to a stack view gravity area at a specified index position.
	InsertViewAtIndexInGravity(view INSView, index uint, gravity NSStackViewGravity)
	// Specifies an array of views for a specified gravity area in the stack view, replacing any previous views in that area.
	SetViewsInGravity(views []NSView, gravity NSStackViewGravity)
	// Removes a specified view from the stack view.
	RemoveView(view INSView)

	// Topic: Managing the Arranged Subviews

	// Adds the specified view to the end of the arranged subviews list.
	AddArrangedSubview(view INSView)
	// Adds the provided view to the array of arranged subviews at the specified index.
	InsertArrangedSubviewAtIndex(view INSView, index int)
	// Removes the provided view from the stack’s array of arranged subviews.
	RemoveArrangedSubview(view INSView)
	// The array of views arranged by the stack view.
	ArrangedSubviews() []NSView

	// Topic: Inspecting a Stack View

	// The array of views owned by the stack view.
	Views() []NSView
	// Returns the array of views in the specified gravity area in the stack view.
	ViewsInGravity(gravity NSStackViewGravity) []NSView
	// An array that contains the detached views from all the stack view’s gravity areas.
	DetachedViews() []NSView
	// Returns the Auto Layout priority for resisting clipping of views in the stack view when Auto Layout attempts to reduce the stack view’s size.
	ClippingResistancePriorityForOrientation(orientation NSLayoutConstraintOrientation) NSLayoutPriority
	// Returns the Auto Layout priority for the stack view to minimize its size to fit its contained views as closely as possible, for a specified user interface axis.
	HuggingPriorityForOrientation(orientation NSLayoutConstraintOrientation) NSLayoutPriority

	// Topic: Configuring the Stack View Layout

	// The horizontal or vertical layout direction of the stack view.
	Orientation() NSUserInterfaceLayoutOrientation
	SetOrientation(value NSUserInterfaceLayoutOrientation)
	// The view alignment within the stack view.
	Alignment() NSLayoutAttribute
	SetAlignment(value NSLayoutAttribute)
	// The minimum spacing, in points, between adjacent views in the stack view.
	Spacing() float64
	SetSpacing(value float64)
	// The geometric padding, in points, inside the stack view, surrounding its views.
	EdgeInsets() foundation.NSEdgeInsets
	SetEdgeInsets(value foundation.NSEdgeInsets)
	// The spacing and sizing distribution of stacked views along the primary axis.
	Distribution() NSStackViewDistribution
	SetDistribution(value NSStackViewDistribution)

	// Topic: Configuring Views in a Stack View

	// Returns the custom spacing, in points, between a specified view in the stack view and the view that follows it.
	CustomSpacingAfterView(view INSView) float64
	// Specifies the custom spacing, in points, between a specified view and the view that follows it in the stack view.
	SetCustomSpacingAfterView(spacing float64, view INSView)
	// Returns the visibility priority for a specified view in the stack view.
	VisibilityPriorityForView(view INSView) NSStackViewVisibilityPriority
	// Sets the Auto Layout priority for a view to remain attached to the stack view when Auto Layout reduces the stack view’s size.
	SetVisibilityPriorityForView(priority NSStackViewVisibilityPriority, view INSView)

	// Topic: Configuring Dynamic Behavior for a Stack View

	// A Boolean value that indicates whether the stack view removes hidden views from its view hierarchy.
	DetachesHiddenViews() bool
	SetDetachesHiddenViews(value bool)
	// Sets the Auto Layout priority for resisting clipping of views in the stack view when Auto Layout attempts to reduce the stack view’s size.
	SetClippingResistancePriorityForOrientation(clippingResistancePriority NSLayoutPriority, orientation NSLayoutConstraintOrientation)
	// Sets the Auto Layout priority for the stack view to minimize its size, for a specified user interface axis.
	SetHuggingPriorityForOrientation(huggingPriority NSLayoutPriority, orientation NSLayoutConstraintOrientation)

	// A Boolean value indicating whether the view is hidden.
	IsHidden() bool
	SetIsHidden(value bool)
	// A required constraint.
	Required() NSLayoutPriority
}

// Init initializes the instance.
func (s NSStackView) Init() NSStackView {
	rv := objc.Send[NSStackView](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSStackView) Autorelease() NSStackView {
	rv := objc.Send[NSStackView](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSStackView creates a new NSStackView instance.
func NewNSStackView() NSStackView {
	class := getNSStackViewClass()
	rv := objc.Send[NSStackView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a view using from data in the specified coder object.
//
// coder: The coder object that contains the view’s configuration details.
//
// # Return Value
//
// An initialized view or `nil` if AppKit couldn’t create the object.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/init(coder:)
func NewStackViewWithCoder(coder foundation.INSCoder) NSStackView {
	instance := getNSStackViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSStackViewFromID(rv)
}

// Initializes and returns a newly allocated [NSView] object with a specified
// frame rectangle.
//
// frameRect: The frame rectangle for the created view object.
//
// # Return Value
//
// An initialized view or `nil` if AppKit couldn’t create the object.
//
// # Discussion
//
// Insert the view into your window’s view hieararchy before you can do
// anything with it. This method is the designated initializer for the
// [NSView] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/init(frame:)
func NewStackViewWithFrame(frameRect corefoundation.CGRect) NSStackView {
	instance := getNSStackViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSStackViewFromID(rv)
}

// Creates and returns a stack view with a specified array of views.
//
// views: The array of views for the new stack view.
//
// # Return Value
//
// A stack view initialized with the specified array of views.
//
// # Discussion
//
// The returned stack view has horizontal layout direction and its
// [TranslatesAutoresizingMaskIntoConstraints] property is set to the Boolean
// value false. The views you provide in the `views` parameter are placed into
// the stack view’s leading gravity area.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/init(views:)
func NewStackViewWithViews(views []NSView) NSStackView {
	rv := objc.Send[objc.ID](objc.ID(getNSStackViewClass().class), objc.Sel("stackViewWithViews:"), objectivec.IObjectSliceToNSArray(views))
	return NSStackViewFromID(rv)
}

// Adds a view to the end of the stack view gravity area.
//
// view: The view to add to the specified gravity area.
//
// gravity: The gravity area that you are adding the specified view to. Valid values
// are those in the [NSStackView.Gravity] enumeration.
//
// # Discussion
//
// The location of a newly added view depends on the stack view layout
// direction and, for a horizontal stack view, on user interface language:
//
// - : A newly added view appears at the trailing edge of the specified
// gravity area, as determined by the value of the inherited
// [UserInterfaceLayoutDirection] property of the stack view. For a left to
// right language, a new view appears at the right side of the gravity area. -
// : A newly added view appears at the bottom of the specified gravity area.
//
// Calling this method updates the stack view’s layout, which can change the
// stack view size. As a result, views could detach or clip according to the
// clipping resistance of the stack view and the visibility priorities of its
// views.
//
// A view in a detached state is not present in the stack view’s view
// hierarchy, but it still consumes memory. To respond to detachment and
// reattachment of views, implement an [NSStackViewDelegate] object and assign
// it to the [Delegate] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/addView(_:in:)
//
// [NSStackView.Gravity]: https://developer.apple.com/documentation/AppKit/NSStackView/Gravity
func (s NSStackView) AddViewInGravity(view INSView, gravity NSStackViewGravity) {
	objc.Send[objc.ID](s.ID, objc.Sel("addView:inGravity:"), view, gravity)
}

// Adds a view to a stack view gravity area at a specified index position.
//
// view: The view to add to the specified gravity area.
//
// index: The index position, within the gravity area, for the new view. The position
// of index `0` depends on the stack view layout direction and, for a
// horizontal stack view, on the user interface layout direction:
//
// - : The `0` index for a gravity area is at the leading side, as determined
// by the value of the inherited [UserInterfaceLayoutDirection] property of
// the stack view. For a left to right language, index `0` is at the left of
// the gravity area. - : The `0` index for a gravity area is at the top.
//
// See the [UserInterfaceLayoutDirection] property and the
// [UserInterfaceLayoutDirection] method.
//
// gravity: The gravity area that you are adding the specified view to. Valid values
// are those in the [NSStackView.Gravity] enumeration.
//
// # Discussion
//
// Calling this method updates the stack view’s layout, which can change the
// stack view size. As a result, views could detach or clip according to the
// clipping resistance of the stack view and the visibility priorities of its
// views.
//
// A view in a detached state is not present in the stack view’s view
// hierarchy, but it still consumes memory. To respond to detachment and
// reattachment of views, implement an [NSStackViewDelegate] object and assign
// it to the [Delegate] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/insertView(_:at:in:)
//
// [NSStackView.Gravity]: https://developer.apple.com/documentation/AppKit/NSStackView/Gravity
func (s NSStackView) InsertViewAtIndexInGravity(view INSView, index uint, gravity NSStackViewGravity) {
	objc.Send[objc.ID](s.ID, objc.Sel("insertView:atIndex:inGravity:"), view, index, gravity)
}

// Specifies an array of views for a specified gravity area in the stack view,
// replacing any previous views in that area.
//
// views: The array of views you are specifying for the gravity area.
//
// gravity: The gravity area that you’re specifying the array of views for. Valid
// values are those in the [NSStackView.Gravity] enumeration, according to the
// stack view’s layout direction.
//
// # Discussion
//
// Calling this method updates the stack view’s layout, which can change the
// stack view size. As a result, views could detach or clip according to the
// clipping resistance of the stack view and the visibility priorities of its
// views.
//
// A view in a detached state is not present in the stack view’s view
// hierarchy, but it still consumes memory. To respond to detachment and
// reattachment of views, implement an [NSStackViewDelegate] object and assign
// it to the [Delegate] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/setViews(_:in:)
//
// [NSStackView.Gravity]: https://developer.apple.com/documentation/AppKit/NSStackView/Gravity
func (s NSStackView) SetViewsInGravity(views []NSView, gravity NSStackViewGravity) {
	objc.Send[objc.ID](s.ID, objc.Sel("setViews:inGravity:"), objectivec.IObjectSliceToNSArray(views), gravity)
}

// Removes a specified view from the stack view.
//
// view: The view you want to remove from the stack view.
//
// # Discussion
//
// This method removes a view from a stack view whether the view is attached
// or detached. For an attached view only, you can alternatively call the
// [removeFromSuperview()] method on the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/removeView(_:)
//
// [removeFromSuperview()]: https://developer.apple.com/documentation/UIKit/UIView/removeFromSuperview()
func (s NSStackView) RemoveView(view INSView) {
	objc.Send[objc.ID](s.ID, objc.Sel("removeView:"), view)
}

// Adds the specified view to the end of the arranged subviews list.
//
// view: The view to add to the end of the [ArrangedSubviews] array.
//
// # Discussion
//
// The stack view ensures that the [ArrangedSubviews] array is always a subset
// of its [Subviews] array. This method automatically adds the provided view
// as a subview of the stack view, if it is not already. If the view is
// already a subview, this operation does not alter the subview ordering.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/addArrangedSubview(_:)
func (s NSStackView) AddArrangedSubview(view INSView) {
	objc.Send[objc.ID](s.ID, objc.Sel("addArrangedSubview:"), view)
}

// Adds the provided view to the array of arranged subviews at the specified
// index.
//
// view: The view to be added to the array of arranged views managed by the stack.
//
// index: The index where the stack inserts the new view in its [ArrangedSubviews]
// array. This value must not be greater than the number of views currently in
// this array. If the index is out of bounds, this method throws an
// [internalInconsistencyException] exception.
//
// # Discussion
//
// If index is already occupied, the stack view increases the size of the
// [ArrangedSubviews] array and shifts all of its contents at the index and
// above to the next higher space in the array. Then the stack view stores the
// provided view at the index.
//
// The stack view also ensures that the [ArrangedSubviews] array is always a
// subset of its [Subviews] array. This method automatically adds the provided
// view as a subview of the stack view, if it is not already. When adding
// subviews, the stack view appends the view to the end of its [Subviews]
// array. The index only affects the order of views in the [ArrangedSubviews]
// array. It does not affect the ordering of views in the [Subviews] array.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/insertArrangedSubview(_:at:)
//
// [internalInconsistencyException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/internalInconsistencyException
func (s NSStackView) InsertArrangedSubviewAtIndex(view INSView, index int) {
	objc.Send[objc.ID](s.ID, objc.Sel("insertArrangedSubview:atIndex:"), view, index)
}

// Removes the provided view from the stack’s array of arranged subviews.
//
// view: The view to be removed from the array of views arranged by the stack.
//
// # Discussion
//
// This method removes the provided view from the stack’s [ArrangedSubviews]
// array. After calling this method, the stack view no longer manages the
// view’s position and size. However, this method does not remove the
// provided view from the stack’s [Subviews] array; therefore, the view
// still appears in the view hierarchy.
//
// To prevent the view from appearing on screen after calling this method,
// explicitly call the view’s [RemoveFromSuperview] method, or set its
// [Hidden] property to true.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/removeArrangedSubview(_:)
func (s NSStackView) RemoveArrangedSubview(view INSView) {
	objc.Send[objc.ID](s.ID, objc.Sel("removeArrangedSubview:"), view)
}

// Returns the array of views in the specified gravity area in the stack view.
//
// gravity: The gravity area whose view array you want to get. Valid values are those
// in the [NSStackView.Gravity] enumeration, according to the stack view’s
// layout direction..
//
// # Return Value
//
// The array of views in a specified gravity area.
//
// # Discussion
//
// The returned array contains all of the views in the specified gravity area
// of the stack view, regardless of whether they are attached. The index
// position of each view in the array matches the view ordering within the
// gravity area. A detached view’s index position is its gravity area
// position when attached.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/views(in:)
//
// [NSStackView.Gravity]: https://developer.apple.com/documentation/AppKit/NSStackView/Gravity
func (s NSStackView) ViewsInGravity(gravity NSStackViewGravity) []NSView {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("viewsInGravity:"), gravity)
	return objc.ConvertSlice(rv, func(id objc.ID) NSView {
		return NSViewFromID(id)
	})
}

// Returns the Auto Layout priority for resisting clipping of views in the
// stack view when Auto Layout attempts to reduce the stack view’s size.
//
// orientation: The stack view layout direction to which the clipping resistance priority
// applies.
//
// # Return Value
//
// A layout constraint priority that identifies the clipping resistance for
// the stack view.
//
// # Discussion
//
// For an explanation of clipping resistance and how to use it for a stack
// view, see the [SetClippingResistancePriorityForOrientation] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/clippingResistancePriority(for:)
func (s NSStackView) ClippingResistancePriorityForOrientation(orientation NSLayoutConstraintOrientation) NSLayoutPriority {
	rv := objc.Send[NSLayoutPriority](s.ID, objc.Sel("clippingResistancePriorityForOrientation:"), orientation)
	return NSLayoutPriority(rv)
}

// Returns the Auto Layout priority for the stack view to minimize its size to
// fit its contained views as closely as possible, for a specified user
// interface axis.
//
// orientation: The user interface axis (horizontal or vertical) whose hugging priority you
// want to get from the stack view. Valid values are those in the
// [NSUserInterfaceLayoutOrientation] enumeration.
//
// # Return Value
//
// The Auto Layout priority for the stack view to minimize its size.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/huggingPriority(for:)
//
// [NSUserInterfaceLayoutOrientation]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutOrientation
func (s NSStackView) HuggingPriorityForOrientation(orientation NSLayoutConstraintOrientation) NSLayoutPriority {
	rv := objc.Send[NSLayoutPriority](s.ID, objc.Sel("huggingPriorityForOrientation:"), orientation)
	return NSLayoutPriority(rv)
}

// Returns the custom spacing, in points, between a specified view in the
// stack view and the view that follows it.
//
// view: The view whose trailing spacing you are getting.
//
// # Return Value
//
// The number of points between the trailing edge of the specified view and
// the one that follows it (that is, the one with the next highest index
// order).
//
// # Discussion
//
// If you set custom spacing for a view using the [SetCustomSpacingAfterView]
// method, it overrides the stack view’s default spacing as set in the
// [Spacing] property.
//
// A stack view retains custom spacing across layout updates. Custom spacing
// for a view is lost if you remove the view from the stack view or specify a
// new value.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/customSpacing(after:)
func (s NSStackView) CustomSpacingAfterView(view INSView) float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("customSpacingAfterView:"), view)
	return rv
}

// Specifies the custom spacing, in points, between a specified view and the
// view that follows it in the stack view.
//
// spacing: The custom trailing space to use between the `aView` view and the one that
// follows it, in points.
//
// Default value is [useDefaultSpacing], which indicates that the view does
// not use custom spacing.
//
// view: The view whose trailing spacing you are setting.
//
// # Discussion
//
// For a horizontal stack view, this method sets custom spacing between a
// specified view and the view to its right when the user interface direction
// is left to right. (See the inherited [UserInterfaceLayoutDirection]
// property for information on layout direction.) For a vertical stack view,
// this method sets custom spacing below a specified view.
//
// If you set custom spacing for a view, it overrides the stack view’s
// default spacing for that view, as set in the [Spacing] property.
//
// A stack view retains custom spacing across layout updates. Custom spacing
// for a view is lost if you remove the view from the stack view or specify a
// new value.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/setCustomSpacing(_:after:)
//
// [useDefaultSpacing]: https://developer.apple.com/documentation/AppKit/NSStackView/useDefaultSpacing
func (s NSStackView) SetCustomSpacingAfterView(spacing float64, view INSView) {
	objc.Send[objc.ID](s.ID, objc.Sel("setCustomSpacing:afterView:"), spacing, view)
}

// Returns the visibility priority for a specified view in the stack view.
//
// view: The view that you are getting the visibility priority for.
//
// # Return Value
//
// The visibility priority for the specified view.
//
// # Discussion
//
// Visibility priority is the Auto Layout priority for a view to remain
// attached to a stack view when Auto Layout reduces the stack view’s size
// (such as when a user reduces the enclosing window’s size).
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/visibilityPriority(for:)
func (s NSStackView) VisibilityPriorityForView(view INSView) NSStackViewVisibilityPriority {
	rv := objc.Send[NSStackViewVisibilityPriority](s.ID, objc.Sel("visibilityPriorityForView:"), view)
	return NSStackViewVisibilityPriority(rv)
}

// Sets the Auto Layout priority for a view to remain attached to the stack
// view when Auto Layout reduces the stack view’s size.
//
// priority: The visibility priority for a specified view. Valid values are those in the
// [NSStackViewVisibilityPriority] enumeration.
//
// view: The view whose visibility priority you are setting.
//
// # Discussion
//
// When Auto Layout reduces the stack view’s size (such as when a user
// reduces the size of the window containing the stack view), causing one or
// more views to no longer fit, the stack view detaches views in order of
// increasing . A view with lower visibility priority detaches before a view
// with higher visibility priority. A set of views with identical, detachable
// visibility priority are all detached or reattached together. A view with
// the highest possible visibility priority never detaches.
//
// A view in a detached state is not present in the stack view’s view
// hierarchy, but it still consumes memory.
//
// The default visibility priority for a view is [mustHold], resulting in the
// view never detaching.
//
// To allow a view to detach as needed by the stack view, set a visibility
// priority of [detachOnlyIfNecessary]. To force a view to detach regardless
// of the enclosing view’s size, set a visibility priority of [notVisible].
// To explicitly reattach a view to a stack view, set a visibility priority of
// [mustHold].
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/setVisibilityPriority(_:for:)
//
// [detachOnlyIfNecessary]: https://developer.apple.com/documentation/AppKit/NSStackView/VisibilityPriority/detachOnlyIfNecessary
// [mustHold]: https://developer.apple.com/documentation/AppKit/NSStackView/VisibilityPriority/mustHold
// [notVisible]: https://developer.apple.com/documentation/AppKit/NSStackView/VisibilityPriority/notVisible
func (s NSStackView) SetVisibilityPriorityForView(priority NSStackViewVisibilityPriority, view INSView) {
	objc.Send[objc.ID](s.ID, objc.Sel("setVisibilityPriority:forView:"), priority, view)
}

// Sets the Auto Layout priority for resisting clipping of views in the stack
// view when Auto Layout attempts to reduce the stack view’s size.
//
// clippingResistancePriority: The clipping resistance Auto Layout priority you want to apply to the stack
// view for a given user interface axis. The default value is [Required],
// which disallows clipping. Other valid values are those in the
// [NSLayoutPriority] enumeration.
//
// orientation: The horizontal or vertical user interface axis that the clipping resistance
// priority applies to; one of the constants from the
// [NSLayoutConstraint.Orientation] enumeration.
//
// # Discussion
//
// A clipped view is one that is partially hidden beyond the border of its
// enclosing stack view. When Auto Layout attempts to reduce the stack
// view’s size (such as when a user attempts to reduce the size of the
// enclosing window), causing a view to no longer fit, the stack view clips or
// detaches the view, or else prevents further reduction of the stack view’s
// size.
//
// To allow view clipping, set a clipping resistance lower than the default
// value of NSLayoutPriorityRequired and set the visibility priority of all
// the stack view’s views to [mustHold].
//
// To ensure that views detach rather than clip, lower the clipping resistance
// for the stack view to a value less than the default of [Required] and set
// the visibility priority for at least one view to a value less than
// [mustHold].
//
// If you disallow view clipping and disallow view detachment, which is the
// default behavior for a stack view, Auto Layout prevents the stack view from
// being reduced in size beyond the minimum needed to show all of its views.
//
// Clipping begins from the right and bottom sides of a stack view.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/setClippingResistancePriority(_:for:)
//
// [NSLayoutConstraint.Orientation]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Orientation
// [mustHold]: https://developer.apple.com/documentation/AppKit/NSStackView/VisibilityPriority/mustHold
func (s NSStackView) SetClippingResistancePriorityForOrientation(clippingResistancePriority NSLayoutPriority, orientation NSLayoutConstraintOrientation) {
	objc.Send[objc.ID](s.ID, objc.Sel("setClippingResistancePriority:forOrientation:"), clippingResistancePriority, orientation)
}

// Sets the Auto Layout priority for the stack view to minimize its size, for
// a specified user interface axis.
//
// huggingPriority: The Auto Layout priority for the stack view to minimize its size. The
// default value is [defaultLow]. Other valid values are those in the
// [NSLayoutPriority] enumeration.
//
// orientation: The horizontal or vertical user interface axis for which you’re setting
// the stack view’s hugging priority; one of the constants from the
// [NSUserInterfaceLayoutOrientation] enumeration.
//
// - To specify horizontal-axis hugging for any stack view (whether it uses
// vertical or horizontal layout), use the
// [NSUserInterfaceLayoutOrientationHorizontal] constant. - To specify
// vertical-axis hugging for any stack view (whether it uses vertical or
// horizontal layout), use the [NSUserInterfaceLayoutOrientationVertical]
// constant.
//
// # Discussion
//
// This method lets you specify a different hugging priority for each user
// interface axis. The default value for hugging priority, on both user
// interface axes, is [defaultLow]. If you have not added constraints between
// the stack view and its enclosing view, the stack view stays as small as
// possible to fully contain its views—independent of the size of the view
// that contains it.
//
// To configure a stack view to grow and shrink according to the size of its
// enclosing view, add constraints between the stack view and its enclosing
// view by using Auto Layout priorities higher than the hugging priority.
//
// To configure a stack view to prevent its enclosing view from growing, use
// priorities for the constraints between the stack view and its enclosing
// view that are lower than the hugging priority.
//
// The value of the hugging priority also affects spacing between views and
// between gravity areas, as described in the discussion for the [Spacing]
// property.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/setHuggingPriority(_:for:)
//
// [defaultLow]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Priority-swift.struct/defaultLow
// [NSUserInterfaceLayoutOrientation]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutOrientation
//
// [defaultLow]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Priority-swift.struct/defaultLow
func (s NSStackView) SetHuggingPriorityForOrientation(huggingPriority NSLayoutPriority, orientation NSLayoutConstraintOrientation) {
	objc.Send[objc.ID](s.ID, objc.Sel("setHuggingPriority:forOrientation:"), huggingPriority, orientation)
}

// The delegate object for the stack view.
//
// # Discussion
//
// The system calls a delegate method when resizing has caused a view to be
// detached from or reattached to the stack view. For more information, see
// [NSStackViewDelegate].
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/delegate
func (s NSStackView) Delegate() NSStackViewDelegate {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("delegate"))
	return NSStackViewDelegateObjectFromID(rv)
}
func (s NSStackView) SetDelegate(value NSStackViewDelegate) {
	objc.Send[struct{}](s.ID, objc.Sel("setDelegate:"), value)
}

// The array of views arranged by the stack view.
//
// # Discussion
//
// The stack view ensures that the contents of this array are always a subset
// of its [Subviews] array.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/arrangedSubviews
func (s NSStackView) ArrangedSubviews() []NSView {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("arrangedSubviews"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSView {
		return NSViewFromID(id)
	})
}

// The array of views owned by the stack view.
//
// # Discussion
//
// The `views` array always contains all of the views owned and managed by the
// stack view, regardless of their gravity area placement and regardless of
// whether or not they are attached. The index position of each view in the
// array matches the view ordering within the stack view. A detached view’s
// index position is its stack view position when attached.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/views
func (s NSStackView) Views() []NSView {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("views"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSView {
		return NSViewFromID(id)
	})
}

// An array that contains the detached views from all the stack view’s
// gravity areas.
//
// # Discussion
//
// A view in a detached state is not present in the stack view’s view
// hierarchy, but it still consumes memory.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/detachedViews
func (s NSStackView) DetachedViews() []NSView {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("detachedViews"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSView {
		return NSViewFromID(id)
	})
}

// The horizontal or vertical layout direction of the stack view.
//
// # Discussion
//
// Default value is [NSUserInterfaceLayoutOrientationHorizontal]. For values
// that apply to this property, see [NSUserInterfaceLayoutOrientation].
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/orientation
//
// [NSUserInterfaceLayoutOrientation]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutOrientation
func (s NSStackView) Orientation() NSUserInterfaceLayoutOrientation {
	rv := objc.Send[NSUserInterfaceLayoutOrientation](s.ID, objc.Sel("orientation"))
	return NSUserInterfaceLayoutOrientation(rv)
}
func (s NSStackView) SetOrientation(value NSUserInterfaceLayoutOrientation) {
	objc.Send[struct{}](s.ID, objc.Sel("setOrientation:"), value)
}

// The view alignment within the stack view.
//
// # Discussion
//
// The default value for this property depends on whether the stack view is
// horizontal or vertical:
//
// - : The default value is [NSLayoutConstraint.Attribute.centerY]. - : The
// default value is [NSLayoutConstraint.Attribute.centerX].
//
// These constants are described as part of the [NSLayoutConstraint.Attribute]
// enumeration in [NSLayoutConstraint]; see that enumeration for the other
// possible alignment values.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/alignment
//
// [NSLayoutConstraint.Attribute.centerX]: https://developer.apple.com/documentation/UIKit/NSLayoutConstraint/Attribute/centerX
// [NSLayoutConstraint.Attribute.centerY]: https://developer.apple.com/documentation/UIKit/NSLayoutConstraint/Attribute/centerY
// [NSLayoutConstraint.Attribute]: https://developer.apple.com/documentation/UIKit/NSLayoutConstraint/Attribute
// [NSLayoutConstraint]: https://developer.apple.com/documentation/UIKit/NSLayoutConstraint
func (s NSStackView) Alignment() NSLayoutAttribute {
	rv := objc.Send[NSLayoutAttribute](s.ID, objc.Sel("alignment"))
	return NSLayoutAttribute(rv)
}
func (s NSStackView) SetAlignment(value NSLayoutAttribute) {
	objc.Send[struct{}](s.ID, objc.Sel("setAlignment:"), value)
}

// The minimum spacing, in points, between adjacent views in the stack view.
//
// # Discussion
//
// A stack view uses this property to define the minimum distance between
// views within a gravity area and between neighboring views in adjacent
// gravity areas. The default value for the [Spacing] property is `8.0`
// points.
//
// The automatically applied Auto Layout constraints for [Spacing] are shown
// in the table below.
//
// [Table data omitted]
//
// The first row indicates that inter-view spacing is constrained to equal the
// value of the [Spacing] property with a priority of at least [defaultHigh];
// you can increase this by setting a higher stack view hugging priority with
// the [SetHuggingPriorityForOrientation] method.
//
// The second row indicates that the spacing between adjacent views in
// neighboring gravity areas is constrained to equal the value of the
// [Spacing] property with the priority of the stack view’s hugging
// priority.
//
// The third row indicates that inter-view spacing is allowed to grow larger
// than the value of the [Spacing] property with a priority of
// NSLayoutPriorityRequired.
//
// In combination, these constraints result in the following typical stack
// view behavior: In a stack view whose [hasEqualSpacing] property is set to
// false (the default) and whose hugging priority is left at [defaultLow] (the
// default), views within a gravity area remain a fixed distance from each
// other (equal to the value of the [Spacing] property), and the distance
// between gravity areas grows and shrinks as the stack view grows and shrinks
// along its layout direction axis. If you set the [hasEqualSpacing] property
// to true and use the default hugging priority, then the distance between all
// adjacent views, including adjacent views in neighboring gravity areas,
// grows and shrinks as the stack view grows and shrinks.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/spacing
//
// [defaultHigh]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Priority-swift.struct/defaultHigh
// [defaultLow]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Priority-swift.struct/defaultLow
// [hasEqualSpacing]: https://developer.apple.com/documentation/AppKit/NSStackView/hasEqualSpacing
func (s NSStackView) Spacing() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("spacing"))
	return rv
}
func (s NSStackView) SetSpacing(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setSpacing:"), value)
}

// The geometric padding, in points, inside the stack view, surrounding its
// views.
//
// # Discussion
//
// The default value is `(0, 0, 0, 0)`. Edge insets remain as they are if you
// change the value of a stack view’s [Orientation] property or the value of
// its inherited [UserInterfaceLayoutDirection] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/edgeInsets
func (s NSStackView) EdgeInsets() foundation.NSEdgeInsets {
	rv := objc.Send[foundation.NSEdgeInsets](s.ID, objc.Sel("edgeInsets"))
	return foundation.NSEdgeInsets(rv)
}
func (s NSStackView) SetEdgeInsets(value foundation.NSEdgeInsets) {
	objc.Send[struct{}](s.ID, objc.Sel("setEdgeInsets:"), value)
}

// The spacing and sizing distribution of stacked views along the primary
// axis.
//
// # Discussion
//
// The default value is `gravityAreas`.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/distribution-swift.property
func (s NSStackView) Distribution() NSStackViewDistribution {
	rv := objc.Send[NSStackViewDistribution](s.ID, objc.Sel("distribution"))
	return NSStackViewDistribution(rv)
}
func (s NSStackView) SetDistribution(value NSStackViewDistribution) {
	objc.Send[struct{}](s.ID, objc.Sel("setDistribution:"), value)
}

// A Boolean value that indicates whether the stack view removes hidden views
// from its view hierarchy.
//
// # Discussion
//
// When the value of this property is true, setting the [Hidden] property of a
// view to true causes the stack view to remove hidden views from its view
// hierarchy and put them in the [DetachedViews] property. Changing the
// view’s [Hidden] property to false causes the stack view to add the view
// back to the view hierarchy. When the value of this property is false, views
// remain in the view hierarchy, even when they are hidden. The default value
// of this property is true.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/detachesHiddenViews
func (s NSStackView) DetachesHiddenViews() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("detachesHiddenViews"))
	return rv
}
func (s NSStackView) SetDetachesHiddenViews(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setDetachesHiddenViews:"), value)
}

// A Boolean value indicating whether the view is hidden.
//
// See: https://developer.apple.com/documentation/appkit/nsview/ishidden
func (s NSStackView) IsHidden() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hidden"))
	return rv
}
func (s NSStackView) SetIsHidden(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setHidden:"), value)
}

// A required constraint.
//
// See: https://developer.apple.com/documentation/appkit/nslayoutconstraint/priority-swift.struct/required
func (s NSStackView) Required() NSLayoutPriority {
	rv := objc.Send[NSLayoutPriority](s.ID, objc.Sel("NSLayoutPriorityRequired"))
	return NSLayoutPriority(rv)
}
