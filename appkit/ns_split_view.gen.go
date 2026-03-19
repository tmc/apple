// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSSplitView] class.
var (
	_NSSplitViewClass     NSSplitViewClass
	_NSSplitViewClassOnce sync.Once
)

func getNSSplitViewClass() NSSplitViewClass {
	_NSSplitViewClassOnce.Do(func() {
		_NSSplitViewClass = NSSplitViewClass{class: objc.GetClass("NSSplitView")}
	})
	return _NSSplitViewClass
}

// GetNSSplitViewClass returns the class object for NSSplitView.
func GetNSSplitViewClass() NSSplitViewClass {
	return getNSSplitViewClass()
}

type NSSplitViewClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSplitViewClass) Alloc() NSSplitView {
	rv := objc.Send[NSSplitView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A view that arranges two or more views in a linear stack running
// horizontally or vertically.
//
// # Overview
// 
// A split view manages the dividers and orientation for a split view
// controller ([NSSplitViewController]). By default, dividers have a
// horizontal orientation so that the split view arranges its panes vertically
// from top to bottom.
// 
// Divider indices are zero-based. If the [NSSplitView.Vertical] property is [false],
// which is the default value, the top divider has an index of `0`. If
// [NSSplitView.Vertical] is [true], the leading divider has an index of `0`.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Customizing the Split View Behavior
//
//   - [NSSplitView.Delegate]: The split view’s delegate.
//   - [NSSplitView.SetDelegate]
//
// # Arranging Subviews
//
//   - [NSSplitView.ArrangesAllSubviews]: A Boolean value that determines whether the split view arranges all of its subviews as split panes.
//   - [NSSplitView.SetArrangesAllSubviews]
//   - [NSSplitView.ArrangedSubviews]: The array of views that the split view arranges as its split panes.
//   - [NSSplitView.AddArrangedSubview]: Adds a view as an arranged split pane.
//   - [NSSplitView.InsertArrangedSubviewAtIndex]: Adds a view as an arranged split pane at the specified index.
//   - [NSSplitView.RemoveArrangedSubview]: Removes a view as an arranged split pane.
//
// # Managing Subviews
//
//   - [NSSplitView.AdjustSubviews]: Adjusts the sizes of the split view’s subviews so they (plus the dividers) fill the split view.
//   - [NSSplitView.IsSubviewCollapsed]: Returns whether the specified view is in a collapsed state.
//   - [NSSplitView.HoldingPriorityForSubviewAtIndex]: Returns the priority of the subview’s width or height when resizing.
//   - [NSSplitView.SetHoldingPriorityForSubviewAtIndex]: Sets the priority for split view subviews to maintain their width or height.
//
// # Managing Divider Orientation
//
//   - [NSSplitView.Vertical]: A Boolean value that determines the geometric orientation of the split view’s dividers.
//   - [NSSplitView.SetVertical]
//
// # Configuring and Drawing Dividers
//
//   - [NSSplitView.DividerStyle]: The style of divider between views.
//   - [NSSplitView.SetDividerStyle]
//   - [NSSplitView.DividerColor]: The color of the dividers that the split view draws between subviews.
//   - [NSSplitView.DividerThickness]: The thickness of the dividers for the split view.
//   - [NSSplitView.DrawDividerInRect]: Draws a divider between two of the split view’s subviews.
//
// # Saving Subview Positions
//
//   - [NSSplitView.AutosaveName]: The name to use when the system automatically saves the split view’s divider configuration.
//   - [NSSplitView.SetAutosaveName]
//
// # Constraining Split Position
//
//   - [NSSplitView.MinPossiblePositionOfDividerAtIndex]: Returns the minimum possible position of the divider at the specified index.
//   - [NSSplitView.MaxPossiblePositionOfDividerAtIndex]: Returns the maximum possible position of the divider at the specified index.
//   - [NSSplitView.SetPositionOfDividerAtIndex]: Updates the location of a divider you specify by index.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitView
type NSSplitView struct {
	NSView
}

// NSSplitViewFromID constructs a [NSSplitView] from an objc.ID.
//
// A view that arranges two or more views in a linear stack running
// horizontally or vertically.
func NSSplitViewFromID(id objc.ID) NSSplitView {
	return NSSplitView{NSView: NSViewFromID(id)}
}
// NOTE: NSSplitView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSplitView] class.
//
// # Customizing the Split View Behavior
//
//   - [INSSplitView.Delegate]: The split view’s delegate.
//   - [INSSplitView.SetDelegate]
//
// # Arranging Subviews
//
//   - [INSSplitView.ArrangesAllSubviews]: A Boolean value that determines whether the split view arranges all of its subviews as split panes.
//   - [INSSplitView.SetArrangesAllSubviews]
//   - [INSSplitView.ArrangedSubviews]: The array of views that the split view arranges as its split panes.
//   - [INSSplitView.AddArrangedSubview]: Adds a view as an arranged split pane.
//   - [INSSplitView.InsertArrangedSubviewAtIndex]: Adds a view as an arranged split pane at the specified index.
//   - [INSSplitView.RemoveArrangedSubview]: Removes a view as an arranged split pane.
//
// # Managing Subviews
//
//   - [INSSplitView.AdjustSubviews]: Adjusts the sizes of the split view’s subviews so they (plus the dividers) fill the split view.
//   - [INSSplitView.IsSubviewCollapsed]: Returns whether the specified view is in a collapsed state.
//   - [INSSplitView.HoldingPriorityForSubviewAtIndex]: Returns the priority of the subview’s width or height when resizing.
//   - [INSSplitView.SetHoldingPriorityForSubviewAtIndex]: Sets the priority for split view subviews to maintain their width or height.
//
// # Managing Divider Orientation
//
//   - [INSSplitView.Vertical]: A Boolean value that determines the geometric orientation of the split view’s dividers.
//   - [INSSplitView.SetVertical]
//
// # Configuring and Drawing Dividers
//
//   - [INSSplitView.DividerStyle]: The style of divider between views.
//   - [INSSplitView.SetDividerStyle]
//   - [INSSplitView.DividerColor]: The color of the dividers that the split view draws between subviews.
//   - [INSSplitView.DividerThickness]: The thickness of the dividers for the split view.
//   - [INSSplitView.DrawDividerInRect]: Draws a divider between two of the split view’s subviews.
//
// # Saving Subview Positions
//
//   - [INSSplitView.AutosaveName]: The name to use when the system automatically saves the split view’s divider configuration.
//   - [INSSplitView.SetAutosaveName]
//
// # Constraining Split Position
//
//   - [INSSplitView.MinPossiblePositionOfDividerAtIndex]: Returns the minimum possible position of the divider at the specified index.
//   - [INSSplitView.MaxPossiblePositionOfDividerAtIndex]: Returns the maximum possible position of the divider at the specified index.
//   - [INSSplitView.SetPositionOfDividerAtIndex]: Updates the location of a divider you specify by index.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitView
type INSSplitView interface {
	INSView

	// Topic: Customizing the Split View Behavior

	// The split view’s delegate.
	Delegate() NSSplitViewDelegate
	SetDelegate(value NSSplitViewDelegate)

	// Topic: Arranging Subviews

	// A Boolean value that determines whether the split view arranges all of its subviews as split panes.
	ArrangesAllSubviews() bool
	SetArrangesAllSubviews(value bool)
	// The array of views that the split view arranges as its split panes.
	ArrangedSubviews() []NSView
	// Adds a view as an arranged split pane.
	AddArrangedSubview(view INSView)
	// Adds a view as an arranged split pane at the specified index.
	InsertArrangedSubviewAtIndex(view INSView, index int)
	// Removes a view as an arranged split pane.
	RemoveArrangedSubview(view INSView)

	// Topic: Managing Subviews

	// Adjusts the sizes of the split view’s subviews so they (plus the dividers) fill the split view.
	AdjustSubviews()
	// Returns whether the specified view is in a collapsed state.
	IsSubviewCollapsed(subview INSView) bool
	// Returns the priority of the subview’s width or height when resizing.
	HoldingPriorityForSubviewAtIndex(subviewIndex int) NSLayoutPriority
	// Sets the priority for split view subviews to maintain their width or height.
	SetHoldingPriorityForSubviewAtIndex(priority NSLayoutPriority, subviewIndex int)

	// Topic: Managing Divider Orientation

	// A Boolean value that determines the geometric orientation of the split view’s dividers.
	Vertical() bool
	SetVertical(value bool)

	// Topic: Configuring and Drawing Dividers

	// The style of divider between views.
	DividerStyle() NSSplitViewDividerStyle
	SetDividerStyle(value NSSplitViewDividerStyle)
	// The color of the dividers that the split view draws between subviews.
	DividerColor() INSColor
	// The thickness of the dividers for the split view.
	DividerThickness() float64
	// Draws a divider between two of the split view’s subviews.
	DrawDividerInRect(rect corefoundation.CGRect)

	// Topic: Saving Subview Positions

	// The name to use when the system automatically saves the split view’s divider configuration.
	AutosaveName() NSSplitViewAutosaveName
	SetAutosaveName(value NSSplitViewAutosaveName)

	// Topic: Constraining Split Position

	// Returns the minimum possible position of the divider at the specified index.
	MinPossiblePositionOfDividerAtIndex(dividerIndex int) float64
	// Returns the maximum possible position of the divider at the specified index.
	MaxPossiblePositionOfDividerAtIndex(dividerIndex int) float64
	// Updates the location of a divider you specify by index.
	SetPositionOfDividerAtIndex(position float64, dividerIndex int)
}

// Init initializes the instance.
func (s NSSplitView) Init() NSSplitView {
	rv := objc.Send[NSSplitView](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSplitView) Autorelease() NSSplitView {
	rv := objc.Send[NSSplitView](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSplitView creates a new NSSplitView instance.
func NewNSSplitView() NSSplitView {
	class := getNSSplitViewClass()
	rv := objc.Send[NSSplitView](objc.ID(class.class), objc.Sel("new"))
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
func NewSplitViewWithCoder(coder foundation.INSCoder) NSSplitView {
	instance := getNSSplitViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSSplitViewFromID(rv)
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
func NewSplitViewWithFrame(frameRect corefoundation.CGRect) NSSplitView {
	instance := getNSSplitViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSSplitViewFromID(rv)
}

// Adds a view as an arranged split pane.
//
// # Discussion
// 
// If the view isn’t a subview of the split view, calling this method adds
// it to the split view’s [subviews] array.
//
// [subviews]: https://developer.apple.com/documentation/AppKit/NSView/subviews
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitView/addArrangedSubview(_:)
func (s NSSplitView) AddArrangedSubview(view INSView) {
	objc.Send[objc.ID](s.ID, objc.Sel("addArrangedSubview:"), view)
}
// Adds a view as an arranged split pane at the specified index.
//
// # Discussion
// 
// If the view is already an arranged view, calling this method moves the view
// to the specified index in the [ArrangedSubviews] array. This change
// doesn’t affect the view’s index in the split view’s [subviews] array.
// 
// If the view isn’t a subview of the split view, calling this method adds
// it to the split view’s [subviews] array.
//
// [subviews]: https://developer.apple.com/documentation/AppKit/NSView/subviews
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitView/insertArrangedSubview(_:at:)
func (s NSSplitView) InsertArrangedSubviewAtIndex(view INSView, index int) {
	objc.Send[objc.ID](s.ID, objc.Sel("insertArrangedSubview:atIndex:"), view, index)
}
// Removes a view as an arranged split pane.
//
// # Discussion
// 
// If the value of [ArrangesAllSubviews] is [false], calling this method
// doesn’t remove the view as a subview; it remains in the split view’s
// [subviews] array.
// 
// If you remove a view as a subview (either by calling [RemoveFromSuperview]
// or removing it from the split view’s [subviews] array), the system
// automatically removes the view as an arranged subview.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [subviews]: https://developer.apple.com/documentation/AppKit/NSView/subviews
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitView/removeArrangedSubview(_:)
func (s NSSplitView) RemoveArrangedSubview(view INSView) {
	objc.Send[objc.ID](s.ID, objc.Sel("removeArrangedSubview:"), view)
}
// Adjusts the sizes of the split view’s subviews so they (plus the
// dividers) fill the split view.
//
// # Discussion
// 
// When you call this method, the split view’s subviews resize
// proportionally; the relative sizes of the subviews don’t change.
// 
// The default implementation of this method resizes subviews proportionally
// so that the ratio of heights (when using horizontal dividers) or widths
// (when using vertical dividers) doesn’t change, even though the absolute
// sizes change.
// 
// Call this method on split views where you’ve added or removed subviews to
// reestablish the consistency of subview placement.
// 
// This method invalidates the cursor when it is over a divider, ensuring the
// cursor is always of the correct type during and after resizing animations.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitView/adjustSubviews()
func (s NSSplitView) AdjustSubviews() {
	objc.Send[objc.ID](s.ID, objc.Sel("adjustSubviews"))
}
// Returns whether the specified view is in a collapsed state.
//
// subview: The subview in the split view.
//
// # Return Value
// 
// [true] if `subview` is in a collapsed state; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitView/isSubviewCollapsed(_:)
func (s NSSplitView) IsSubviewCollapsed(subview INSView) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isSubviewCollapsed:"), subview)
	return rv
}
// Returns the priority of the subview’s width or height when resizing.
//
// subviewIndex: The index of the subview.
//
// # Return Value
// 
// The layout priority of the subview at the index.
//
// # Discussion
// 
// The priority is the manner that the split view subviews use to maintain
// their width (for a vertical split view) or height (for a horizontal split
// view). During a split view resize, subviews with higher priorities maintain
// their sizes before subviews with lower priorities. The subview with the
// lowest priority is the first to gain additional thickness if the split view
// grows or shrinks.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitView/holdingPriorityForSubview(at:)
func (s NSSplitView) HoldingPriorityForSubviewAtIndex(subviewIndex int) NSLayoutPriority {
	rv := objc.Send[NSLayoutPriority](s.ID, objc.Sel("holdingPriorityForSubviewAtIndex:"), subviewIndex)
	return NSLayoutPriority(rv)
}
// Sets the priority for split view subviews to maintain their width or
// height.
//
// priority: The priority.
//
// subviewIndex: The index of the subview
//
// # Discussion
// 
// Calling this method sets the priority that split view subviews use to
// maintain their width (for a vertical split view) or height (for a
// horizontal split view). During a split view resize, subviews with higher
// priorities maintain their sizes before subviews with lower priorities. The
// subview with the lowest priority is the first to gain additional thickness
// if the split view grows or shrinks.
// 
// The default priority is [defaultLow]. Use priorities less than
// [dragThatCannotResizeWindow].
//
// [defaultLow]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Priority-swift.struct/defaultLow
// [dragThatCannotResizeWindow]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Priority-swift.struct/dragThatCannotResizeWindow
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitView/setHoldingPriority(_:forSubviewAt:)
func (s NSSplitView) SetHoldingPriorityForSubviewAtIndex(priority NSLayoutPriority, subviewIndex int) {
	objc.Send[objc.ID](s.ID, objc.Sel("setHoldingPriority:forSubviewAtIndex:"), priority, subviewIndex)
}
// Draws a divider between two of the split view’s subviews.
//
// rect: The entire divider rectangle in the split view’s flipped coordinate
// system.
//
// # Discussion
// 
// If you override this method to use a custom style for the divider, you may
// need to change the size of the divider.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitView/drawDivider(in:)
func (s NSSplitView) DrawDividerInRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](s.ID, objc.Sel("drawDividerInRect:"), rect)
}
// Returns the minimum possible position of the divider at the specified
// index.
//
// dividerIndex: The index of the divider.
//
// # Return Value
// 
// A [CGFloat] that specifies the minimum possible position of the divider.
//
// [CGFloat]: https://developer.apple.com/documentation/CoreFoundation/CGFloat-swift.struct
//
// # Discussion
// 
// The position is because the bounds of the split view and the current
// position of other dividers dictate it. positions result from letting the
// delegate apply constraints to the possible positions.
// 
// You can invoke this method to determine the range of values that you can
// pass to [SetPositionOfDividerAtIndex]. You can also invoke it from delegate
// methods like [SplitViewConstrainSplitPositionOfSubviewAt] to implement
// relatively complex behaviors that depend on the current state of the split
// view.
// 
// The result of invoking this method when you haven’t invoked
// [AdjustSubviews], and the subview frames are invalid, is undefined.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitView/minPossiblePositionOfDivider(at:)
func (s NSSplitView) MinPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("minPossiblePositionOfDividerAtIndex:"), dividerIndex)
	return rv
}
// Returns the maximum possible position of the divider at the specified
// index.
//
// dividerIndex: The index of the divider.
//
// # Return Value
// 
// A [CGFloat] that specifies the maximum possible position of the divider.
//
// [CGFloat]: https://developer.apple.com/documentation/CoreFoundation/CGFloat-swift.struct
//
// # Discussion
// 
// The position is because the bounds of the split view and the current
// position of other dividers dictate it. positions result from letting the
// delegate apply constraints to the possible positions.
// 
// You can invoke this method to determine the range of values that you can
// pass to [SetPositionOfDividerAtIndex]. You can also invoke it from delegate
// methods like [SplitViewConstrainSplitPositionOfSubviewAt] to implement
// relatively complex behaviors that depend on the current state of the split
// view.
// 
// The result of invoking this method when you haven’t invoked
// [AdjustSubviews], and the subview frames are invalid, is undefined.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitView/maxPossiblePositionOfDivider(at:)
func (s NSSplitView) MaxPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("maxPossiblePositionOfDividerAtIndex:"), dividerIndex)
	return rv
}
// Updates the location of a divider you specify by index.
//
// position: The position of the divider.
//
// dividerIndex: The index of the divider.
//
// # Discussion
// 
// One of the views adjacent to the divider may collapse because the
// method’s default implementation assumes a person is dragging the divider
// to the new location. The Auto Layout system collapses the view if it
// can’t satisfy the view’s constraints — typically imposed by its
// delegate — with the divider’s new location.
// 
// [NSSplitView] doesn’t invoke this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitView/setPosition(_:ofDividerAt:)
func (s NSSplitView) SetPositionOfDividerAtIndex(position float64, dividerIndex int) {
	objc.Send[objc.ID](s.ID, objc.Sel("setPosition:ofDividerAtIndex:"), position, dividerIndex)
}

// The split view’s delegate.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitView/delegate
func (s NSSplitView) Delegate() NSSplitViewDelegate {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("delegate"))
	return NSSplitViewDelegateObjectFromID(rv)
}
func (s NSSplitView) SetDelegate(value NSSplitViewDelegate) {
	objc.Send[struct{}](s.ID, objc.Sel("setDelegate:"), value)
}
// A Boolean value that determines whether the split view arranges all of its
// subviews as split panes.
//
// # Discussion
// 
// If the value of this property is [true], the split view arranges all of its
// subviews automatically. The [ArrangedSubviews] array is identical to the
// split view’s [subviews] array, so any change to [subviews] reflects in
// the [ArrangedSubviews] array. The default value of this property is [true].
// 
// If the value of this property is [false], you must explicitly add a view as
// an arranged subview to arrange it as a split pane. You add an arranged
// subview using [AddArrangedSubview].
// 
// When you change the value of this property from [true] to [false], all
// existing subviews stay as arranged subviews in [ArrangedSubviews]. When you
// change the value of this property from [false] to [true], all existing
// subviews become arranged subviews, and the value of the [subviews] array
// becomes the [ArrangedSubviews] array.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [subviews]: https://developer.apple.com/documentation/AppKit/NSView/subviews
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitView/arrangesAllSubviews
func (s NSSplitView) ArrangesAllSubviews() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("arrangesAllSubviews"))
	return rv
}
func (s NSSplitView) SetArrangesAllSubviews(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setArrangesAllSubviews:"), value)
}
// The array of views that the split view arranges as its split panes.
//
// # Discussion
// 
// This array contains a subset of the views in the split view’s [subviews]
// property. The views in this array may appear in a different order than in
// the [subviews] array.
// 
// If the value of [ArrangesAllSubviews] is [true], this array is identical to
// the [subviews] array.
//
// [subviews]: https://developer.apple.com/documentation/AppKit/NSView/subviews
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitView/arrangedSubviews
func (s NSSplitView) ArrangedSubviews() []NSView {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("arrangedSubviews"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSView {
		return NSViewFromID(id)
	})
}
// A Boolean value that determines the geometric orientation of the split
// view’s dividers.
//
// # Discussion
// 
// The default value of this property is [false], which indicates horizontal
// dividers and views that stack one above the other (top-to-bottom) in the
// containing split view controller’s view.
// 
// To specify vertical dividers and a horizontal (side-by-side) arrangement of
// views within a split view controller, implement this property to return
// [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitView/isVertical
func (s NSSplitView) Vertical() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isVertical"))
	return rv
}
func (s NSSplitView) SetVertical(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setVertical:"), value)
}
// The style of divider between views.
//
// # Discussion
// 
// See [NSSplitView.DividerStyle] for the possible values.
//
// [NSSplitView.DividerStyle]: https://developer.apple.com/documentation/AppKit/NSSplitView/DividerStyle-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitView/dividerStyle-swift.property
func (s NSSplitView) DividerStyle() NSSplitViewDividerStyle {
	rv := objc.Send[NSSplitViewDividerStyle](s.ID, objc.Sel("dividerStyle"))
	return NSSplitViewDividerStyle(rv)
}
func (s NSSplitView) SetDividerStyle(value NSSplitViewDividerStyle) {
	objc.Send[struct{}](s.ID, objc.Sel("setDividerStyle:"), value)
}
// The color of the dividers that the split view draws between subviews.
//
// # Discussion
// 
// The default implementation of this method returns [clear] when the split
// view’s [DividerStyle] is [NSSplitView.DividerStyle.thick], or when
// [DividerStyle] is [NSSplitView.DividerStyle.paneSplitter] and the split
// view is in a textured window. The system draws all other thin dividers with
// a color that provides appropriate contrast between two white panes.
// 
// You can subclass [NSSplitView] and override this method to change the color
// of dividers.
//
// [NSSplitView.DividerStyle.paneSplitter]: https://developer.apple.com/documentation/AppKit/NSSplitView/DividerStyle-swift.enum/paneSplitter
// [NSSplitView.DividerStyle.thick]: https://developer.apple.com/documentation/AppKit/NSSplitView/DividerStyle-swift.enum/thick
// [clear]: https://developer.apple.com/documentation/AppKit/NSColor/clear
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitView/dividerColor
func (s NSSplitView) DividerColor() INSColor {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("dividerColor"))
	return NSColorFromID(objc.ID(rv))
}
// The thickness of the dividers for the split view.
//
// # Discussion
// 
// You can subclass [NSSplitView] and override this method to change the
// thickness of a split view’s dividers.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitView/dividerThickness
func (s NSSplitView) DividerThickness() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("dividerThickness"))
	return rv
}
// The name to use when the system automatically saves the split view’s
// divider configuration.
//
// # Discussion
// 
// If this property’s value is `nil` or empty, autosaving doesn’t occur.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitView/autosaveName-swift.property
func (s NSSplitView) AutosaveName() NSSplitViewAutosaveName {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("autosaveName"))
	return NSSplitViewAutosaveName(foundation.NSStringFromID(rv).String())
}
func (s NSSplitView) SetAutosaveName(value NSSplitViewAutosaveName) {
	objc.Send[struct{}](s.ID, objc.Sel("setAutosaveName:"), objc.String(string(value)))
}

