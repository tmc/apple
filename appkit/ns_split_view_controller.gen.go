// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSplitViewController] class.
var (
	_NSSplitViewControllerClass     NSSplitViewControllerClass
	_NSSplitViewControllerClassOnce sync.Once
)

func getNSSplitViewControllerClass() NSSplitViewControllerClass {
	_NSSplitViewControllerClassOnce.Do(func() {
		_NSSplitViewControllerClass = NSSplitViewControllerClass{class: objc.GetClass("NSSplitViewController")}
	})
	return _NSSplitViewControllerClass
}

// GetNSSplitViewControllerClass returns the class object for NSSplitViewController.
func GetNSSplitViewControllerClass() NSSplitViewControllerClass {
	return getNSSplitViewControllerClass()
}

type NSSplitViewControllerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSplitViewControllerClass) Alloc() NSSplitViewController {
	rv := objc.Send[NSSplitViewController](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that manages an array of adjacent child views, and has a split
// view object for managing dividers between those views.
//
// # Overview
// 
// A split view controller manages a set of child views that it displays next
// to each other in a side-by-side or top-to-bottom arrangement.
// 
// A split view controller owns an array of split view items
// ([NSSplitViewItem]), each of which has a view controller
// ([NSViewController]) and corresponding view. The split view controller’s
// [SplitView] object manages those child views and the dividers between them.
// 
// By default, a split view arranges its child views vertically from top to
// bottom. To specify a horizontal (side-by-side) arrangement, implement the
// [Vertical] property of the [SplitView] object to return [true].
// 
// The split view controller serves as the delegate of its [SplitView] object.
// If you override a split view delegate method, your override must call
// `super`.
// 
// To use a split view controller, you must use Auto Layout for the child
// views and to support animations that collapse and reveal child views. For
// example, if you design a layout that contains two views, a content area and
// an optional sidebar, you employ Auto Layout constraints to specify whether
// the content area shrinks or remains the same size when the sidebar becomes
// visible.
// 
// A split view controller employs lazy loading of its views. For example,
// adding a collapsed split view item as a new child doesn’t load the
// associated view until it shows.
// 
// For more information about using [NSSplitViewController] in your app, see
// [Navigating Hierarchical Data Using Outline and Split Views].
//
// [Navigating Hierarchical Data Using Outline and Split Views]: https://developer.apple.com/documentation/AppKit/navigating-hierarchical-data-using-outline-and-split-views
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Configuring and Managing a Split View Controller
//
//   - [NSSplitViewController.SplitView]: The split view that the split view controller manages.
//   - [NSSplitViewController.SetSplitView]
//   - [NSSplitViewController.SplitViewItemForViewController]: Returns the corresponding split view item for the specified child view controller of the split view controller.
//   - [NSSplitViewController.SplitViewItems]: The array of split view items that correspond to the split view controller’s child view controllers.
//   - [NSSplitViewController.SetSplitViewItems]
//
// # Modifying a Split View Controller
//
//   - [NSSplitViewController.AddSplitViewItem]: Adds a split view item to the end of the array of split view items.
//   - [NSSplitViewController.InsertSplitViewItemAtIndex]: Adds a split view item to the array of split view items at the specified index position.
//   - [NSSplitViewController.RemoveSplitViewItem]: Removes a specified split view item from the split view controller.
//
// # Managing Sidebars
//
//   - [NSSplitViewController.ToggleSidebar]: Collapses or expands the first sidebar in the split view controller using an animation.
//   - [NSSplitViewController.MinimumThicknessForInlineSidebars]: The minimum thickness for a sidebar before it automatically collapses.
//   - [NSSplitViewController.SetMinimumThicknessForInlineSidebars]
//
// # Managing Inspectors
//
//   - [NSSplitViewController.ToggleInspector]
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewController
type NSSplitViewController struct {
	NSViewController
}

// NSSplitViewControllerFromID constructs a [NSSplitViewController] from an objc.ID.
//
// An object that manages an array of adjacent child views, and has a split
// view object for managing dividers between those views.
func NSSplitViewControllerFromID(id objc.ID) NSSplitViewController {
	return NSSplitViewController{NSViewController: NSViewControllerFromID(id)}
}
// NOTE: NSSplitViewController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSplitViewController] class.
//
// # Configuring and Managing a Split View Controller
//
//   - [INSSplitViewController.SplitView]: The split view that the split view controller manages.
//   - [INSSplitViewController.SetSplitView]
//   - [INSSplitViewController.SplitViewItemForViewController]: Returns the corresponding split view item for the specified child view controller of the split view controller.
//   - [INSSplitViewController.SplitViewItems]: The array of split view items that correspond to the split view controller’s child view controllers.
//   - [INSSplitViewController.SetSplitViewItems]
//
// # Modifying a Split View Controller
//
//   - [INSSplitViewController.AddSplitViewItem]: Adds a split view item to the end of the array of split view items.
//   - [INSSplitViewController.InsertSplitViewItemAtIndex]: Adds a split view item to the array of split view items at the specified index position.
//   - [INSSplitViewController.RemoveSplitViewItem]: Removes a specified split view item from the split view controller.
//
// # Managing Sidebars
//
//   - [INSSplitViewController.ToggleSidebar]: Collapses or expands the first sidebar in the split view controller using an animation.
//   - [INSSplitViewController.MinimumThicknessForInlineSidebars]: The minimum thickness for a sidebar before it automatically collapses.
//   - [INSSplitViewController.SetMinimumThicknessForInlineSidebars]
//
// # Managing Inspectors
//
//   - [INSSplitViewController.ToggleInspector]
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewController
type INSSplitViewController interface {
	INSViewController
	NSSplitViewDelegate
	NSUserInterfaceValidations

	// Topic: Configuring and Managing a Split View Controller

	// The split view that the split view controller manages.
	SplitView() INSSplitView
	SetSplitView(value INSSplitView)
	// Returns the corresponding split view item for the specified child view controller of the split view controller.
	SplitViewItemForViewController(viewController INSViewController) INSSplitViewItem
	// The array of split view items that correspond to the split view controller’s child view controllers.
	SplitViewItems() []NSSplitViewItem
	SetSplitViewItems(value []NSSplitViewItem)

	// Topic: Modifying a Split View Controller

	// Adds a split view item to the end of the array of split view items.
	AddSplitViewItem(splitViewItem INSSplitViewItem)
	// Adds a split view item to the array of split view items at the specified index position.
	InsertSplitViewItemAtIndex(splitViewItem INSSplitViewItem, index int)
	// Removes a specified split view item from the split view controller.
	RemoveSplitViewItem(splitViewItem INSSplitViewItem)

	// Topic: Managing Sidebars

	// Collapses or expands the first sidebar in the split view controller using an animation.
	ToggleSidebar(sender objectivec.IObject)
	// The minimum thickness for a sidebar before it automatically collapses.
	MinimumThicknessForInlineSidebars() float64
	SetMinimumThicknessForInlineSidebars(value float64)

	// Topic: Managing Inspectors

	ToggleInspector(sender objectivec.IObject)

	// A Boolean value that determines the geometric orientation of the split view’s dividers.
	IsVertical() bool
	SetIsVertical(value bool)
}

// Init initializes the instance.
func (s NSSplitViewController) Init() NSSplitViewController {
	rv := objc.Send[NSSplitViewController](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSplitViewController) Autorelease() NSSplitViewController {
	rv := objc.Send[NSSplitViewController](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSplitViewController creates a new NSSplitViewController instance.
func NewNSSplitViewController() NSSplitViewController {
	class := getNSSplitViewControllerClass()
	rv := objc.Send[NSSplitViewController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/init(coder:)
func NewSplitViewControllerWithCoder(coder foundation.INSCoder) NSSplitViewController {
	instance := getNSSplitViewControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSSplitViewControllerFromID(rv)
}

// Returns a view controller object initialized to the nib file in the
// specified bundle.
//
// nibNameOrNil: The name of the nib file, without any leading path information.
//
// nibBundleOrNil: The bundle in which to search for the nib file. If you specify `nil`, this
// method looks for the nib file in the main bundle.
//
// # Return Value
// 
// The initialized [NSViewController] object.
//
// # Discussion
// 
// The [NSViewController] object looks for the nib file in the bundle’s
// language-specific project directories first, followed by the Resources
// directory.
// 
// The specified nib file should typically have the class of the file’s
// owner set to [NSViewController], or a custom subclass, with the `view`
// outlet connected to a view.
// 
// If you pass in `nil` for `nibNameOrNil`, [NibName] returns `nil` and
// [LoadView] throws an exception; in this case you must set [View] before
// [View] is invoked, or override [LoadView].
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/init(nibName:bundle:)
func NewSplitViewControllerWithNibNameBundle(nibNameOrNil NSNibName, nibBundleOrNil foundation.NSBundle) NSSplitViewController {
	instance := getNSSplitViewControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNibName:bundle:"), objc.String(string(nibNameOrNil)), nibBundleOrNil)
	return NSSplitViewControllerFromID(rv)
}

// Returns the corresponding split view item for the specified child view
// controller of the split view controller.
//
// viewController: The child view controller with the corresponding split view item you want.
//
// # Return Value
// 
// The corresponding split view item, or `nil` if `viewController` isn’t a
// child of the split view controller.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewController/splitViewItem(for:)
func (s NSSplitViewController) SplitViewItemForViewController(viewController INSViewController) INSSplitViewItem {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("splitViewItemForViewController:"), viewController)
	return NSSplitViewItemFromID(rv)
}

// Adds a split view item to the end of the array of split view items.
//
// splitViewItem: The split view item to add.
//
// # Discussion
// 
// This is a convenience method you can use in place of the
// [InsertSplitViewItemAtIndex] method when you want to add a split view item
// to the end of the [SplitViewItems] array. Calling this method implicitly
// calls the [InsertSplitViewItemAtIndex] method.
// 
// If you subclass the [NSSplitViewController] class, don’t call this method
// in your custom object to add a split view item. Instead, call the
// [InsertSplitViewItemAtIndex] method directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewController/addSplitViewItem(_:)
func (s NSSplitViewController) AddSplitViewItem(splitViewItem INSSplitViewItem) {
	objc.Send[objc.ID](s.ID, objc.Sel("addSplitViewItem:"), splitViewItem)
}

// Adds a split view item to the array of split view items at the specified
// index position.
//
// splitViewItem: The split view item to add.
//
// index: The index position for adding the split view item in the [SplitViewItems]
// array.
//
// # Discussion
// 
// If the split view controller’s view finishes loading, and the split view
// item that you add is visible, the system loads the split view item’s view
// controller’s view and adds it to the split view.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewController/insertSplitViewItem(_:at:)
func (s NSSplitViewController) InsertSplitViewItemAtIndex(splitViewItem INSSplitViewItem, index int) {
	objc.Send[objc.ID](s.ID, objc.Sel("insertSplitViewItem:atIndex:"), splitViewItem, index)
}

// Removes a specified split view item from the split view controller.
//
// splitViewItem: The split view item to remove.
//
// # Discussion
// 
// After you remove a split view item, the system adjusts the layout of the
// split view accordingly.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewController/removeSplitViewItem(_:)
func (s NSSplitViewController) RemoveSplitViewItem(splitViewItem INSSplitViewItem) {
	objc.Send[objc.ID](s.ID, objc.Sel("removeSplitViewItem:"), splitViewItem)
}

// Collapses or expands the first sidebar in the split view controller using
// an animation.
//
// # Discussion
// 
// If the split view controller doesn’t contain a sidebar, calling this
// method does nothing.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewController/toggleSidebar(_:)
func (s NSSplitViewController) ToggleSidebar(sender objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("toggleSidebar:"), sender)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewController/toggleInspector(_:)
func (s NSSplitViewController) ToggleInspector(sender objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("toggleInspector:"), sender)
}

// Allows the split view controller to return an additional rectangle where
// mouse clicks can initiate divider dragging.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewController/splitView(_:additionalEffectiveRectOfDividerAt:)
func (s NSSplitViewController) SplitViewAdditionalEffectiveRectOfDividerAtIndex(splitView INSSplitView, dividerIndex int) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("splitView:additionalEffectiveRectOfDividerAtIndex:"), splitView, dividerIndex)
	return corefoundation.CGRect(rv)
}

// Allows the split view controller to determine whether the user can collapse
// and expand the specified subview.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewController/splitView(_:canCollapseSubview:)
func (s NSSplitViewController) SplitViewCanCollapseSubview(splitView INSSplitView, subview INSView) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("splitView:canCollapseSubview:"), splitView, subview)
	return rv
}

// Allows the delegate to constrain the maximum coordinate limit of a divider
// when the user drags it.
//
// splitView: The split view that sends the message.
//
// proposedMaximumPosition: The proposed maximum coordinate limit of the subview in the split view’s
// flipped coordinate system.
//
// dividerIndex: Specifies the divider the user is moving, with the first divider being 0
// and increasing from top to bottom (or left to right).
//
// # Return Value
// 
// The maximum coordinate limit of the divider.
//
// # Discussion
// 
// The delegate receives this message before the split view begins tracking
// the mouse to position a divider. You can further constrain the limits, but
// you can’t extend the divider limits.
// 
// If the split bars are horizontal (views are one on top of the other),
// `proposedMax` is the bottom limit. If the split bars are vertical (views
// are side by side), `proposedMax` is the right limit. The initial value of
// `proposedMax` is the bottom (or right side) of the subview after the
// divider.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewDelegate/splitView(_:constrainMaxCoordinate:ofSubviewAt:)
func (s NSSplitViewController) SplitViewConstrainMaxCoordinateOfSubviewAt(splitView INSSplitView, proposedMaximumPosition float64, dividerIndex int) float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("splitView:constrainMaxCoordinate:ofSubviewAt:"), splitView, proposedMaximumPosition, dividerIndex)
	return rv
}

// Allows the delegate to constrain the minimum coordinate limit of a divider
// when the user drags it.
//
// splitView: The split view that sends the message.
//
// proposedMinimumPosition: The proposed minimum coordinate limit of the subview in the split view’s
// flipped coordinate system.
//
// dividerIndex: Specifies the divider the user is moving, with the first divider being 0
// and increasing from top to bottom (or left to right).
//
// # Return Value
// 
// The minimum coordinate limit of the divider.
//
// # Discussion
// 
// The delegate receives this message before the split view begins tracking
// the cursor to position a divider. You can further constrain the limits, but
// you can’t extend the divider limits.
// 
// If the split bars are horizontal (views are one on top of the other),
// `proposedMin` is the top limit. If the split bars are vertical (views are
// side by side), `proposedMin` is the left limit. The initial value of
// `proposedMin` is the top (or left side) of the subview before the divider.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewDelegate/splitView(_:constrainMinCoordinate:ofSubviewAt:)
func (s NSSplitViewController) SplitViewConstrainMinCoordinateOfSubviewAt(splitView INSSplitView, proposedMinimumPosition float64, dividerIndex int) float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("splitView:constrainMinCoordinate:ofSubviewAt:"), splitView, proposedMinimumPosition, dividerIndex)
	return rv
}

// Allows the delegate to constrain the divider to certain positions.
//
// splitView: The split view that sends the message.
//
// proposedPosition: The cursor’s current position, and the proposed position of the divider.
//
// dividerIndex: The index of the divider the user is moving, with the first divider being
// `0` and increasing from top to bottom (or left to right).
//
// # Return Value
// 
// The position for constraining the divider.
//
// # Discussion
// 
// If the delegate implements this method, the split view calls it repeatedly
// as the user moves the divider.
// 
// If a subview’s height must be a multiple of a certain number, use this
// method to return the multiple nearest to `proposedPosition`.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewDelegate/splitView(_:constrainSplitPosition:ofSubviewAt:)
func (s NSSplitViewController) SplitViewConstrainSplitPositionOfSubviewAt(splitView INSSplitView, proposedPosition float64, dividerIndex int) float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("splitView:constrainSplitPosition:ofSubviewAt:"), splitView, proposedPosition, dividerIndex)
	return rv
}

// Notifies the delegate when the split view resizes its subviews.
//
// notification: A notification named [didResizeSubviewsNotification], which posts after a
// change to the size of some or all subviews of a split view.
// //
// [didResizeSubviewsNotification]: https://developer.apple.com/documentation/AppKit/NSSplitView/didResizeSubviewsNotification
//
// # Discussion
// 
// If the delegate implements this method, the system automatically registers
// it to receive this notification.
// 
// The default notification center invokes this method after the split view
// resizes two of its subviews in response to the repositioning of a divider.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewDelegate/splitViewDidResizeSubviews(_:)
func (s NSSplitViewController) SplitViewDidResizeSubviews(notification foundation.NSNotification) {
	objc.Send[objc.ID](s.ID, objc.Sel("splitViewDidResizeSubviews:"), notification)
}

// Allows the split view controller to modify the rectangle where mouse clicks
// initiate divider dragging.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewController/splitView(_:effectiveRect:forDrawnRect:ofDividerAt:)
func (s NSSplitViewController) SplitViewEffectiveRectForDrawnRectOfDividerAtIndex(splitView INSSplitView, proposedEffectiveRect corefoundation.CGRect, drawnRect corefoundation.CGRect, dividerIndex int) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("splitView:effectiveRect:forDrawnRect:ofDividerAtIndex:"), splitView, proposedEffectiveRect, drawnRect, dividerIndex)
	return corefoundation.CGRect(rv)
}

// Allows the delegate to specify custom sizing behavior for the subviews of
// the split view.
//
// splitView: The split view that sends the message.
//
// oldSize: The size of the split view before the user resizes it.
//
// # Discussion
// 
// If the delegate implements this method, it receives this message after the
// split view resizes.
// 
// Resize the subviews so that the sum of the sizes of the subviews plus the
// sum of the thickness of the dividers equals the size of the new frame of
// the [NSSplitView]. You can get the thickness of a divider through the
// [DividerThickness] method.
// 
// If you implement this delegate method to resize subviews on your own, the
// [NSSplitView] doesn’t perform any error checking for you. However, you
// can invoke [AdjustSubviews] to perform the default sizing behavior.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewDelegate/splitView(_:resizeSubviewsWithOldSize:)
func (s NSSplitViewController) SplitViewResizeSubviewsWithOldSize(splitView INSSplitView, oldSize corefoundation.CGSize) {
	objc.Send[objc.ID](s.ID, objc.Sel("splitView:resizeSubviewsWithOldSize:"), splitView, oldSize)
}

// Allows the delegate to specify whether to resize the subview.
//
// splitView: The split view that sends the message.
//
// view: The subview to resize.
//
// # Return Value
// 
// If [AdjustSubviews] can change the size of the subview, [true]; otherwise,
// [false]. By returning [false], you lock the size of the split view
// `subview` while the split view resizes.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Regardless of the value that this method returns, [AdjustSubviews] may
// change the origin of the subview. Nonresizable subviews may resize to
// prevent an invalid subview layout.
// 
// If a split view has no delegate, or if its delegate doesn’t respond to
// this message, the split view behaves as if this method returns [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewDelegate/splitView(_:shouldAdjustSizeOfSubview:)
func (s NSSplitViewController) SplitViewShouldAdjustSizeOfSubview(splitView INSSplitView, view INSView) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("splitView:shouldAdjustSizeOfSubview:"), splitView, view)
	return rv
}

// Allows the split view controller to determine whether the user can drag a
// divider or adjust it off the edge of the split view.
//
// # Discussion
// 
// By default, [NSSplitViewController] hides the first and last dividers if
// their outer neighbor is in a collapsed state.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewController/splitView(_:shouldHideDividerAt:)
func (s NSSplitViewController) SplitViewShouldHideDividerAtIndex(splitView INSSplitView, dividerIndex int) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("splitView:shouldHideDividerAtIndex:"), splitView, dividerIndex)
	return rv
}

// Notifies the delegate when the split view is about to resize its subviews.
//
// notification: A notification named [willResizeSubviewsNotification], which posts before a
// change to the size of some or all subviews of a split view.
// //
// [willResizeSubviewsNotification]: https://developer.apple.com/documentation/AppKit/NSSplitView/willResizeSubviewsNotification
//
// # Discussion
// 
// If the delegate implements this method, the system automatically registers
// it to receive this notification.
// 
// The default notification center invokes this method before the split view
// resizes two of its subviews in response to the repositioning of a divider.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewDelegate/splitViewWillResizeSubviews(_:)
func (s NSSplitViewController) SplitViewWillResizeSubviews(notification foundation.NSNotification) {
	objc.Send[objc.ID](s.ID, objc.Sel("splitViewWillResizeSubviews:"), notification)
}

// Returns a Boolean value that indicates whether to enable the specified
// item.
//
// item: The user interface item to validate.
//
// # Return Value
// 
// [true] if the specified item responds to [ToggleSidebar], [false] if it
// doesn’t.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewController/validateUserInterfaceItem(_:)
func (s NSSplitViewController) ValidateUserInterfaceItem(item NSValidatedUserInterfaceItem) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("validateUserInterfaceItem:"), item)
	return rv
}

// The split view that the split view controller manages.
//
// # Discussion
// 
// This property gives you access to the split view controller’s split view
// for querying its attributes or customizing it.
// 
// By default, a split view arranges its child views vertically from top to
// bottom. To specify a horizontal (side-by-side) arrangement, implement the
// [Vertical] property of the split view object to return [true].
// 
// Also by default, a split view has a divider style of
// [NSSplitView.DividerStyle.thin], and doesn’t have an autosave name.
// 
// To provide a custom split view, set this property at any time before you
// call `super` in the inherited [ViewDidLoad] method; that is, before the
// split view controller’s [ViewLoaded] property is [true].
// 
// The split view isn’t always the same object as that in the split view
// controller’s inherited [View] property. To access the split view, always
// use the [SplitView] property.
//
// [NSSplitView.DividerStyle.thin]: https://developer.apple.com/documentation/AppKit/NSSplitView/DividerStyle-swift.enum/thin
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewController/splitView
func (s NSSplitViewController) SplitView() INSSplitView {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("splitView"))
	return NSSplitViewFromID(objc.ID(rv))
}
func (s NSSplitViewController) SetSplitView(value INSSplitView) {
	objc.Send[struct{}](s.ID, objc.Sel("setSplitView:"), value)
}

// The array of split view items that correspond to the split view
// controller’s child view controllers.
//
// # Discussion
// 
// Setting this property implicitly calls the [InsertSplitViewItemAtIndex] or
// [RemoveSplitViewItem] method, as appropriate, to add or remove split view
// items from this array.
// 
// If you add a child view controller to the split view controller, the system
// automatically creates a default split view item for the child view
// controller and adds it to the [SplitViewItems] array.
// 
// If you remove a child view controller, the split view controller removes
// its corresponding split view item from the [SplitViewItems] array.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewController/splitViewItems
func (s NSSplitViewController) SplitViewItems() []NSSplitViewItem {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("splitViewItems"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSSplitViewItem {
		return NSSplitViewItemFromID(id)
	})
}
func (s NSSplitViewController) SetSplitViewItems(value []NSSplitViewItem) {
	objc.Send[struct{}](s.ID, objc.Sel("setSplitViewItems:"), objectivec.IObjectSliceToNSArray(value))
}

// The minimum thickness for a sidebar before it automatically collapses.
//
// # Discussion
// 
// This value describes the minimum thickness in the primary axis of a split
// view—width if the split view’s [Vertical] value is [true], height if
// it’s [false]. This value is the minimum thickness that sidebars can
// shrink to before they automatically collapse. When sidebars autocollapse in
// full-screen mode, they overlay the other split view items.
// 
// Autocollapsed sidebars automatically expand if their thickness increases to
// or above this minimum thickness threshold.
// 
// The default value of this property is [automaticDimension], which
// determines the minimum thickness for sidebars using the effective minimum
// size of the split view item views from the layout constraints in the
// window. If the system can’t apply the constraints that establish the
// minimum size for all noncollapsed split panes, all sidebars automatically
// collapse. In full-screen mode, if a sidebar attempts to expand in this
// state, it overlays instead.
//
// [automaticDimension]: https://developer.apple.com/documentation/AppKit/NSSplitViewController/automaticDimension
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewController/minimumThicknessForInlineSidebars
func (s NSSplitViewController) MinimumThicknessForInlineSidebars() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("minimumThicknessForInlineSidebars"))
	return rv
}
func (s NSSplitViewController) SetMinimumThicknessForInlineSidebars(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setMinimumThicknessForInlineSidebars:"), value)
}

// A Boolean value that determines the geometric orientation of the split
// view’s dividers.
//
// See: https://developer.apple.com/documentation/appkit/nssplitview/isvertical
func (s NSSplitViewController) IsVertical() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("vertical"))
	return rv
}
func (s NSSplitViewController) SetIsVertical(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setVertical:"), value)
}

			// Protocol methods for NSSplitViewDelegate
			

			// Protocol methods for NSUserInterfaceValidations
			

