// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSplitViewItem] class.
var (
	_NSSplitViewItemClass     NSSplitViewItemClass
	_NSSplitViewItemClassOnce sync.Once
)

func getNSSplitViewItemClass() NSSplitViewItemClass {
	_NSSplitViewItemClassOnce.Do(func() {
		_NSSplitViewItemClass = NSSplitViewItemClass{class: objc.GetClass("NSSplitViewItem")}
	})
	return _NSSplitViewItemClass
}

// GetNSSplitViewItemClass returns the class object for NSSplitViewItem.
func GetNSSplitViewItemClass() NSSplitViewItemClass {
	return getNSSplitViewItemClass()
}

type NSSplitViewItemClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSplitViewItemClass) Alloc() NSSplitViewItem {
	rv := objc.Send[NSSplitViewItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An item in a split view controller.
//
// # Overview
// 
// A split view item represents a single pane in a split view controller
// ([NSSplitViewController]). Each split view item contains information about
// a child view controller in the split view controller, like its preferred
// thickness, holding priority, and collapsed state.
// 
// To add one or more accessory views to the top or bottom of a split view
// item, such as a search field above a list, use the
// [NSSplitViewItem.TopAlignedAccessoryViewControllers] and
// [NSSplitViewItem.BottomAlignedAccessoryViewControllers] properties to specify
// [NSSplitViewItemAccessoryViewController] types.
//
// # Managing the item thickness
//
//   - [NSSplitViewItem.AutomaticMaximumThickness]: The maximum thickness of the split view item when it resizes due to automatic sizing.
//   - [NSSplitViewItem.SetAutomaticMaximumThickness]
//   - [NSSplitViewItem.PreferredThicknessFraction]: The preferred thickness of the split view item relative to the split view.
//   - [NSSplitViewItem.SetPreferredThicknessFraction]
//   - [NSSplitViewItem.MinimumThickness]: The minimum thickness of the split view item.
//   - [NSSplitViewItem.SetMinimumThickness]
//   - [NSSplitViewItem.MaximumThickness]: The maximum thickness of the split view item.
//   - [NSSplitViewItem.SetMaximumThickness]
//
// # Getting Auto Layout behaviors
//
//   - [NSSplitViewItem.HoldingPriority]: The priority for a split view item to hold its size.
//   - [NSSplitViewItem.SetHoldingPriority]
//   - [NSSplitViewItem.AutomaticallyAdjustsSafeAreaInsets]: When YES, other items such as sidebars or inspectors may appear overlaid on top of this item’s `viewController` and this item’s `safeAreaInsets` will be adjusted with respect to overlaid content. Defaults to [NO].
//   - [NSSplitViewItem.SetAutomaticallyAdjustsSafeAreaInsets]
//
// # Getting the item behavior
//
//   - [NSSplitViewItem.Behavior]: The standard behavior type of the split view item.
//
// # Collapsing and expanding the item
//
//   - [NSSplitViewItem.Collapsed]: A Boolean value that determines whether the child view controller that corresponds to the split view item is in a collapsed state in the split view controller.
//   - [NSSplitViewItem.SetCollapsed]
//   - [NSSplitViewItem.CanCollapse]: A Boolean value that determines whether a user interaction can collapse the child view controller that corresponds to the split view item.
//   - [NSSplitViewItem.SetCanCollapse]
//   - [NSSplitViewItem.CollapseBehavior]: The resizing behavior when the split view item toggles its collapsed state.
//   - [NSSplitViewItem.SetCollapseBehavior]
//   - [NSSplitViewItem.SpringLoaded]: A Boolean value that determines whether the split view item can temporarily expand during a drag.
//   - [NSSplitViewItem.SetSpringLoaded]
//   - [NSSplitViewItem.CanCollapseFromWindowResize]
//   - [NSSplitViewItem.SetCanCollapseFromWindowResize]
//
// # Customizing appearance
//
//   - [NSSplitViewItem.AllowsFullHeightLayout]: A Boolean value that indicates whether full-height sidebars appear in the window after you set a style mask.
//   - [NSSplitViewItem.SetAllowsFullHeightLayout]
//   - [NSSplitViewItem.TitlebarSeparatorStyle]: The type of separator that the app displays between the title bar and content of a window.
//   - [NSSplitViewItem.SetTitlebarSeparatorStyle]
//
// # Configuring accessory views
//
//   - [NSSplitViewItem.TopAlignedAccessoryViewControllers]: The following methods allow you to add accessory views to the top/bottom of this splitViewItem. See [NSSplitViewItemAccessoryViewController] for more details.
//   - [NSSplitViewItem.SetTopAlignedAccessoryViewControllers]
//   - [NSSplitViewItem.BottomAlignedAccessoryViewControllers]
//   - [NSSplitViewItem.SetBottomAlignedAccessoryViewControllers]
//   - [NSSplitViewItem.AddTopAlignedAccessoryViewController]
//   - [NSSplitViewItem.InsertTopAlignedAccessoryViewControllerAtIndex]
//   - [NSSplitViewItem.RemoveTopAlignedAccessoryViewControllerAtIndex]: NOTE: you can use this method, or `-`, whichever is easier.
//   - [NSSplitViewItem.AddBottomAlignedAccessoryViewController]
//   - [NSSplitViewItem.InsertBottomAlignedAccessoryViewControllerAtIndex]
//   - [NSSplitViewItem.RemoveBottomAlignedAccessoryViewControllerAtIndex]: NOTE: you can use this method, or `-`, whichever is easier.
//
// # Getting the View Controller
//
//   - [NSSplitViewItem.ViewController]: The view controller that the split view item represents.
//   - [NSSplitViewItem.SetViewController]
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem
type NSSplitViewItem struct {
	objectivec.Object
}

// NSSplitViewItemFromID constructs a [NSSplitViewItem] from an objc.ID.
//
// An item in a split view controller.
func NSSplitViewItemFromID(id objc.ID) NSSplitViewItem {
	return NSSplitViewItem{objectivec.Object{ID: id}}
}
// NOTE: NSSplitViewItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSplitViewItem] class.
//
// # Managing the item thickness
//
//   - [INSSplitViewItem.AutomaticMaximumThickness]: The maximum thickness of the split view item when it resizes due to automatic sizing.
//   - [INSSplitViewItem.SetAutomaticMaximumThickness]
//   - [INSSplitViewItem.PreferredThicknessFraction]: The preferred thickness of the split view item relative to the split view.
//   - [INSSplitViewItem.SetPreferredThicknessFraction]
//   - [INSSplitViewItem.MinimumThickness]: The minimum thickness of the split view item.
//   - [INSSplitViewItem.SetMinimumThickness]
//   - [INSSplitViewItem.MaximumThickness]: The maximum thickness of the split view item.
//   - [INSSplitViewItem.SetMaximumThickness]
//
// # Getting Auto Layout behaviors
//
//   - [INSSplitViewItem.HoldingPriority]: The priority for a split view item to hold its size.
//   - [INSSplitViewItem.SetHoldingPriority]
//   - [INSSplitViewItem.AutomaticallyAdjustsSafeAreaInsets]: When YES, other items such as sidebars or inspectors may appear overlaid on top of this item’s `viewController` and this item’s `safeAreaInsets` will be adjusted with respect to overlaid content. Defaults to [NO].
//   - [INSSplitViewItem.SetAutomaticallyAdjustsSafeAreaInsets]
//
// # Getting the item behavior
//
//   - [INSSplitViewItem.Behavior]: The standard behavior type of the split view item.
//
// # Collapsing and expanding the item
//
//   - [INSSplitViewItem.Collapsed]: A Boolean value that determines whether the child view controller that corresponds to the split view item is in a collapsed state in the split view controller.
//   - [INSSplitViewItem.SetCollapsed]
//   - [INSSplitViewItem.CanCollapse]: A Boolean value that determines whether a user interaction can collapse the child view controller that corresponds to the split view item.
//   - [INSSplitViewItem.SetCanCollapse]
//   - [INSSplitViewItem.CollapseBehavior]: The resizing behavior when the split view item toggles its collapsed state.
//   - [INSSplitViewItem.SetCollapseBehavior]
//   - [INSSplitViewItem.SpringLoaded]: A Boolean value that determines whether the split view item can temporarily expand during a drag.
//   - [INSSplitViewItem.SetSpringLoaded]
//   - [INSSplitViewItem.CanCollapseFromWindowResize]
//   - [INSSplitViewItem.SetCanCollapseFromWindowResize]
//
// # Customizing appearance
//
//   - [INSSplitViewItem.AllowsFullHeightLayout]: A Boolean value that indicates whether full-height sidebars appear in the window after you set a style mask.
//   - [INSSplitViewItem.SetAllowsFullHeightLayout]
//   - [INSSplitViewItem.TitlebarSeparatorStyle]: The type of separator that the app displays between the title bar and content of a window.
//   - [INSSplitViewItem.SetTitlebarSeparatorStyle]
//
// # Configuring accessory views
//
//   - [INSSplitViewItem.TopAlignedAccessoryViewControllers]: The following methods allow you to add accessory views to the top/bottom of this splitViewItem. See [NSSplitViewItemAccessoryViewController] for more details.
//   - [INSSplitViewItem.SetTopAlignedAccessoryViewControllers]
//   - [INSSplitViewItem.BottomAlignedAccessoryViewControllers]
//   - [INSSplitViewItem.SetBottomAlignedAccessoryViewControllers]
//   - [INSSplitViewItem.AddTopAlignedAccessoryViewController]
//   - [INSSplitViewItem.InsertTopAlignedAccessoryViewControllerAtIndex]
//   - [INSSplitViewItem.RemoveTopAlignedAccessoryViewControllerAtIndex]: NOTE: you can use this method, or `-`, whichever is easier.
//   - [INSSplitViewItem.AddBottomAlignedAccessoryViewController]
//   - [INSSplitViewItem.InsertBottomAlignedAccessoryViewControllerAtIndex]
//   - [INSSplitViewItem.RemoveBottomAlignedAccessoryViewControllerAtIndex]: NOTE: you can use this method, or `-`, whichever is easier.
//
// # Getting the View Controller
//
//   - [INSSplitViewItem.ViewController]: The view controller that the split view item represents.
//   - [INSSplitViewItem.SetViewController]
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem
type INSSplitViewItem interface {
	objectivec.IObject

	// Topic: Managing the item thickness

	// The maximum thickness of the split view item when it resizes due to automatic sizing.
	AutomaticMaximumThickness() float64
	SetAutomaticMaximumThickness(value float64)
	// The preferred thickness of the split view item relative to the split view.
	PreferredThicknessFraction() float64
	SetPreferredThicknessFraction(value float64)
	// The minimum thickness of the split view item.
	MinimumThickness() float64
	SetMinimumThickness(value float64)
	// The maximum thickness of the split view item.
	MaximumThickness() float64
	SetMaximumThickness(value float64)

	// Topic: Getting Auto Layout behaviors

	// The priority for a split view item to hold its size.
	HoldingPriority() NSLayoutPriority
	SetHoldingPriority(value NSLayoutPriority)
	// When YES, other items such as sidebars or inspectors may appear overlaid on top of this item’s `viewController` and this item’s `safeAreaInsets` will be adjusted with respect to overlaid content. Defaults to [NO].
	AutomaticallyAdjustsSafeAreaInsets() bool
	SetAutomaticallyAdjustsSafeAreaInsets(value bool)

	// Topic: Getting the item behavior

	// The standard behavior type of the split view item.
	Behavior() NSSplitViewItemBehavior

	// Topic: Collapsing and expanding the item

	// A Boolean value that determines whether the child view controller that corresponds to the split view item is in a collapsed state in the split view controller.
	Collapsed() bool
	SetCollapsed(value bool)
	// A Boolean value that determines whether a user interaction can collapse the child view controller that corresponds to the split view item.
	CanCollapse() bool
	SetCanCollapse(value bool)
	// The resizing behavior when the split view item toggles its collapsed state.
	CollapseBehavior() NSSplitViewItemCollapseBehavior
	SetCollapseBehavior(value NSSplitViewItemCollapseBehavior)
	// A Boolean value that determines whether the split view item can temporarily expand during a drag.
	SpringLoaded() bool
	SetSpringLoaded(value bool)
	CanCollapseFromWindowResize() bool
	SetCanCollapseFromWindowResize(value bool)

	// Topic: Customizing appearance

	// A Boolean value that indicates whether full-height sidebars appear in the window after you set a style mask.
	AllowsFullHeightLayout() bool
	SetAllowsFullHeightLayout(value bool)
	// The type of separator that the app displays between the title bar and content of a window.
	TitlebarSeparatorStyle() NSTitlebarSeparatorStyle
	SetTitlebarSeparatorStyle(value NSTitlebarSeparatorStyle)

	// Topic: Configuring accessory views

	// The following methods allow you to add accessory views to the top/bottom of this splitViewItem. See [NSSplitViewItemAccessoryViewController] for more details.
	TopAlignedAccessoryViewControllers() []NSSplitViewItemAccessoryViewController
	SetTopAlignedAccessoryViewControllers(value []NSSplitViewItemAccessoryViewController)
	BottomAlignedAccessoryViewControllers() []NSSplitViewItemAccessoryViewController
	SetBottomAlignedAccessoryViewControllers(value []NSSplitViewItemAccessoryViewController)
	AddTopAlignedAccessoryViewController(childViewController INSSplitViewItemAccessoryViewController)
	InsertTopAlignedAccessoryViewControllerAtIndex(childViewController INSSplitViewItemAccessoryViewController, index int)
	// NOTE: you can use this method, or `-`, whichever is easier.
	RemoveTopAlignedAccessoryViewControllerAtIndex(index int)
	AddBottomAlignedAccessoryViewController(childViewController INSSplitViewItemAccessoryViewController)
	InsertBottomAlignedAccessoryViewControllerAtIndex(childViewController INSSplitViewItemAccessoryViewController, index int)
	// NOTE: you can use this method, or `-`, whichever is easier.
	RemoveBottomAlignedAccessoryViewControllerAtIndex(index int)

	// Topic: Getting the View Controller

	// The view controller that the split view item represents.
	ViewController() INSViewController
	SetViewController(value INSViewController)

	// Returns the animation that should be performed for the specified key.
	AnimationForKey(key NSAnimatablePropertyKey) objectivec.IObject
	// Returns a proxy object for the receiver that can be used to initiate implied animation for property changes.
	Animator() INSSplitViewItem
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s NSSplitViewItem) Init() NSSplitViewItem {
	rv := objc.Send[NSSplitViewItem](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSplitViewItem) Autorelease() NSSplitViewItem {
	rv := objc.Send[NSSplitViewItem](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSplitViewItem creates a new NSSplitViewItem instance.
func NewNSSplitViewItem() NSSplitViewItem {
	class := getNSSplitViewItemClass()
	rv := objc.Send[NSSplitViewItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a split view item that represents a content list for the specified
// view controller.
//
// # Discussion
// 
// You use a content list to display information like the Mail app’s list of
// messages or the Notes app’s list of notes.
// 
// Content lists use standard system default values for these properties:
// 
// - [MinimumThickness], [MaximumThickness], and [AutomaticMaximumThickness]
// use the standard system defaults for content lists -
// [PreferredThicknessFraction] uses the standard fraction for content lists
// (`0.28` if an adjacent sidebar is visible, `0.33` if not)
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(contentListWithViewController:)
func NewSplitViewItemContentListWithViewController(viewController INSViewController) NSSplitViewItem {
	rv := objc.Send[objc.ID](objc.ID(getNSSplitViewItemClass().class), objc.Sel("contentListWithViewController:"), viewController)
	return NSSplitViewItemFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(inspectorWithViewController:)
func NewSplitViewItemInspectorWithViewController(viewController INSViewController) NSSplitViewItem {
	rv := objc.Send[objc.ID](objc.ID(getNSSplitViewItemClass().class), objc.Sel("inspectorWithViewController:"), viewController)
	return NSSplitViewItemFromID(rv)
}

// Creates a split view item that represents a sidebar for the specified view
// controller.
//
// # Discussion
// 
// Sidebar items take advantage of the standard system appearance and behavior
// for sidebars, including:
// 
// - The translucent material background - The ability to collapse and expand
// on split view size changes - The ability to overlay at small split view
// sizes in full-screen mode
// 
// Additionally, sidebars use standard system default values for these
// properties:
// 
// - [CanCollapse] and [SpringLoaded] are [true] - [MinimumThickness] and
// [MaximumThickness] use the standard minimum and maximum sidebar size -
// [PreferredThicknessFraction] uses the standard fraction for sidebars
// (`0.15`)
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(sidebarWithViewController:)
func NewSplitViewItemSidebarWithViewController(viewController INSViewController) NSSplitViewItem {
	rv := objc.Send[objc.ID](objc.ID(getNSSplitViewItemClass().class), objc.Sel("sidebarWithViewController:"), viewController)
	return NSSplitViewItemFromID(rv)
}

// Creates a split view item that represents the specified view controller.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/init(viewController:)
func NewSplitViewItemWithViewController(viewController INSViewController) NSSplitViewItem {
	rv := objc.Send[objc.ID](objc.ID(getNSSplitViewItemClass().class), objc.Sel("splitViewItemWithViewController:"), viewController)
	return NSSplitViewItemFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/addTopAlignedAccessoryViewController(_:)
func (s NSSplitViewItem) AddTopAlignedAccessoryViewController(childViewController INSSplitViewItemAccessoryViewController) {
	objc.Send[objc.ID](s.ID, objc.Sel("addTopAlignedAccessoryViewController:"), childViewController)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/insertTopAlignedAccessoryViewController(_:at:)
func (s NSSplitViewItem) InsertTopAlignedAccessoryViewControllerAtIndex(childViewController INSSplitViewItemAccessoryViewController, index int) {
	objc.Send[objc.ID](s.ID, objc.Sel("insertTopAlignedAccessoryViewController:atIndex:"), childViewController, index)
}

// NOTE: you can use this method, or `-`, whichever is easier.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/removeTopAlignedAccessoryViewController(at:)
func (s NSSplitViewItem) RemoveTopAlignedAccessoryViewControllerAtIndex(index int) {
	objc.Send[objc.ID](s.ID, objc.Sel("removeTopAlignedAccessoryViewControllerAtIndex:"), index)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/addBottomAlignedAccessoryViewController(_:)
func (s NSSplitViewItem) AddBottomAlignedAccessoryViewController(childViewController INSSplitViewItemAccessoryViewController) {
	objc.Send[objc.ID](s.ID, objc.Sel("addBottomAlignedAccessoryViewController:"), childViewController)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/insertBottomAlignedAccessoryViewController(_:at:)
func (s NSSplitViewItem) InsertBottomAlignedAccessoryViewControllerAtIndex(childViewController INSSplitViewItemAccessoryViewController, index int) {
	objc.Send[objc.ID](s.ID, objc.Sel("insertBottomAlignedAccessoryViewController:atIndex:"), childViewController, index)
}

// NOTE: you can use this method, or `-`, whichever is easier.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/removeBottomAlignedAccessoryViewController(at:)
func (s NSSplitViewItem) RemoveBottomAlignedAccessoryViewControllerAtIndex(index int) {
	objc.Send[objc.ID](s.ID, objc.Sel("removeBottomAlignedAccessoryViewControllerAtIndex:"), index)
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
func (s NSSplitViewItem) AnimationForKey(key NSAnimatablePropertyKey) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("animationForKey:"), objc.String(string(key)))
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
func (s NSSplitViewItem) Animator() INSSplitViewItem {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("animator"))
	return NSSplitViewItemFromID(rv)
}
func (s NSSplitViewItem) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
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
func (_NSSplitViewItemClass NSSplitViewItemClass) DefaultAnimationForKey(key NSAnimatablePropertyKey) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NSSplitViewItemClass.class), objc.Sel("defaultAnimationForKey:"), objc.String(string(key)))
	return objectivec.Object{ID: rv}
}

// The maximum thickness of the split view item when it resizes due to
// automatic sizing.
//
// # Discussion
// 
// Automatic sizing may happen when the split view item has a set
// [PreferredThicknessFraction] and the app enters full-screen mode, or when
// other split view items cause the item to change size. The user can still
// resize the item up to its absolute maximum size in [MaximumThickness] by
// dragging the divider.
// 
// The default value of this property is [unspecifiedDimension], which means
// the split view item doesn’t enforce an automatic maximum size.
//
// [unspecifiedDimension]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/unspecifiedDimension
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/automaticMaximumThickness
func (s NSSplitViewItem) AutomaticMaximumThickness() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("automaticMaximumThickness"))
	return rv
}
func (s NSSplitViewItem) SetAutomaticMaximumThickness(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setAutomaticMaximumThickness:"), value)
}

// The preferred thickness of the split view item relative to the split view.
//
// # Discussion
// 
// This value represents the proportion of the split view that you want the
// split view item to occupy on a scale of `0.0` to `1.0`. The system uses
// this value to adjust the thickness of the item in relation to other items
// when a user double-clicks a neighboring divider, and when the app enters
// full-screen mode.
// 
// The default value of this property is [unspecifiedDimension], which means
// the item doesn’t resize when a user double-clicks the divider, and the
// system preserves the absolute size when the app enters full-screen mode.
//
// [unspecifiedDimension]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/unspecifiedDimension
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/preferredThicknessFraction
func (s NSSplitViewItem) PreferredThicknessFraction() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("preferredThicknessFraction"))
	return rv
}
func (s NSSplitViewItem) SetPreferredThicknessFraction(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setPreferredThicknessFraction:"), value)
}

// The minimum thickness of the split view item.
//
// # Discussion
// 
// This value affects the split view item’s width (for a vertical split
// view) or height (for a horizontal split view).
// 
// The default value of this property is [unspecifiedDimension], which means
// the split view item doesn’t enforce a minimum size. However, layout
// constraints in the contained view hierarchy might specify a minimum size
// regardless.
//
// [unspecifiedDimension]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/unspecifiedDimension
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/minimumThickness
func (s NSSplitViewItem) MinimumThickness() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("minimumThickness"))
	return rv
}
func (s NSSplitViewItem) SetMinimumThickness(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setMinimumThickness:"), value)
}

// The maximum thickness of the split view item.
//
// # Discussion
// 
// This value affects the split view item’s width (for a vertical split
// view) or height (for a horizontal split view).
// 
// The default value of this property is [unspecifiedDimension], which means
// the split view item doesn’t enforce a maximum size. However, layout
// constraints in the contained view hierarchy might specify a maximum size
// regardless.
//
// [unspecifiedDimension]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/unspecifiedDimension
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/maximumThickness
func (s NSSplitViewItem) MaximumThickness() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("maximumThickness"))
	return rv
}
func (s NSSplitViewItem) SetMaximumThickness(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setMaximumThickness:"), value)
}

// The priority for a split view item to hold its size.
//
// # Discussion
// 
// This priority affects the split view item’s width (for a vertical split
// view) or height (for a horizontal split view).
// 
// The view with the lowest priority is the first to gain additional width if
// the split view grows or shrinks. The default of this property is
// [defaultLow].
//
// [defaultLow]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Priority-swift.struct/defaultLow
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/holdingPriority
func (s NSSplitViewItem) HoldingPriority() NSLayoutPriority {
	rv := objc.Send[NSLayoutPriority](s.ID, objc.Sel("holdingPriority"))
	return NSLayoutPriority(rv)
}
func (s NSSplitViewItem) SetHoldingPriority(value NSLayoutPriority) {
	objc.Send[struct{}](s.ID, objc.Sel("setHoldingPriority:"), value)
}

// When YES, other items such as sidebars or inspectors may appear overlaid on
// top of this item’s `viewController` and this item’s `safeAreaInsets`
// will be adjusted with respect to overlaid content. Defaults to [NO].
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/automaticallyAdjustsSafeAreaInsets
func (s NSSplitViewItem) AutomaticallyAdjustsSafeAreaInsets() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("automaticallyAdjustsSafeAreaInsets"))
	return rv
}
func (s NSSplitViewItem) SetAutomaticallyAdjustsSafeAreaInsets(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setAutomaticallyAdjustsSafeAreaInsets:"), value)
}

// The standard behavior type of the split view item.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/behavior-swift.property
func (s NSSplitViewItem) Behavior() NSSplitViewItemBehavior {
	rv := objc.Send[NSSplitViewItemBehavior](s.ID, objc.Sel("behavior"))
	return NSSplitViewItemBehavior(rv)
}

// A Boolean value that determines whether the child view controller that
// corresponds to the split view item is in a collapsed state in the split
// view controller.
//
// # Discussion
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/isCollapsed
func (s NSSplitViewItem) Collapsed() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isCollapsed"))
	return rv
}
func (s NSSplitViewItem) SetCollapsed(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setCollapsed:"), value)
}

// A Boolean value that determines whether a user interaction can collapse the
// child view controller that corresponds to the split view item.
//
// # Discussion
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/canCollapse
func (s NSSplitViewItem) CanCollapse() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("canCollapse"))
	return rv
}
func (s NSSplitViewItem) SetCanCollapse(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setCanCollapse:"), value)
}

// The resizing behavior when the split view item toggles its collapsed state.
//
// # Discussion
// 
// The default value of this property is
// [NSSplitViewItem.CollapseBehavior.default].
//
// [NSSplitViewItem.CollapseBehavior.default]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/CollapseBehavior-swift.enum/default
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/collapseBehavior-swift.property
func (s NSSplitViewItem) CollapseBehavior() NSSplitViewItemCollapseBehavior {
	rv := objc.Send[NSSplitViewItemCollapseBehavior](s.ID, objc.Sel("collapseBehavior"))
	return NSSplitViewItemCollapseBehavior(rv)
}
func (s NSSplitViewItem) SetCollapseBehavior(value NSSplitViewItemCollapseBehavior) {
	objc.Send[struct{}](s.ID, objc.Sel("setCollapseBehavior:"), value)
}

// A Boolean value that determines whether the split view item can temporarily
// expand during a drag.
//
// # Discussion
// 
// If the value of this property is [true], the split view item can
// temporarily expand during a drag if the user hovers or force clicks its
// neighboring divider.
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/isSpringLoaded
func (s NSSplitViewItem) SpringLoaded() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isSpringLoaded"))
	return rv
}
func (s NSSplitViewItem) SetSpringLoaded(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setSpringLoaded:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/canCollapseFromWindowResize
func (s NSSplitViewItem) CanCollapseFromWindowResize() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("canCollapseFromWindowResize"))
	return rv
}
func (s NSSplitViewItem) SetCanCollapseFromWindowResize(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setCanCollapseFromWindowResize:"), value)
}

// A Boolean value that indicates whether full-height sidebars appear in the
// window after you set a style mask.
//
// # Discussion
// 
// This property only applies to [NSSplitViewItem.Behavior.sidebar] and
// [NSSplitViewItem.Behavior.inspector]. The default value is [true].
//
// [NSSplitViewItem.Behavior.inspector]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/Behavior-swift.enum/inspector
// [NSSplitViewItem.Behavior.sidebar]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/Behavior-swift.enum/sidebar
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/allowsFullHeightLayout
func (s NSSplitViewItem) AllowsFullHeightLayout() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("allowsFullHeightLayout"))
	return rv
}
func (s NSSplitViewItem) SetAllowsFullHeightLayout(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setAllowsFullHeightLayout:"), value)
}

// The type of separator that the app displays between the title bar and
// content of a window.
//
// # Discussion
// 
// To apply this value, you must associate the item’s view with its own
// title bar section. The default value is
// [NSTitlebarSeparatorStyle.automatic]. The containing window’s preference
// can override this preference.
//
// [NSTitlebarSeparatorStyle.automatic]: https://developer.apple.com/documentation/AppKit/NSTitlebarSeparatorStyle/automatic
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/titlebarSeparatorStyle
func (s NSSplitViewItem) TitlebarSeparatorStyle() NSTitlebarSeparatorStyle {
	rv := objc.Send[NSTitlebarSeparatorStyle](s.ID, objc.Sel("titlebarSeparatorStyle"))
	return NSTitlebarSeparatorStyle(rv)
}
func (s NSSplitViewItem) SetTitlebarSeparatorStyle(value NSTitlebarSeparatorStyle) {
	objc.Send[struct{}](s.ID, objc.Sel("setTitlebarSeparatorStyle:"), value)
}

// The following methods allow you to add accessory views to the top/bottom of
// this splitViewItem. See [NSSplitViewItemAccessoryViewController] for more
// details.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/topAlignedAccessoryViewControllers
func (s NSSplitViewItem) TopAlignedAccessoryViewControllers() []NSSplitViewItemAccessoryViewController {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("topAlignedAccessoryViewControllers"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSSplitViewItemAccessoryViewController {
		return NSSplitViewItemAccessoryViewControllerFromID(id)
	})
}
func (s NSSplitViewItem) SetTopAlignedAccessoryViewControllers(value []NSSplitViewItemAccessoryViewController) {
	objc.Send[struct{}](s.ID, objc.Sel("setTopAlignedAccessoryViewControllers:"), objectivec.IObjectSliceToNSArray(value))
}

// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/bottomAlignedAccessoryViewControllers
func (s NSSplitViewItem) BottomAlignedAccessoryViewControllers() []NSSplitViewItemAccessoryViewController {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("bottomAlignedAccessoryViewControllers"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSSplitViewItemAccessoryViewController {
		return NSSplitViewItemAccessoryViewControllerFromID(id)
	})
}
func (s NSSplitViewItem) SetBottomAlignedAccessoryViewControllers(value []NSSplitViewItemAccessoryViewController) {
	objc.Send[struct{}](s.ID, objc.Sel("setBottomAlignedAccessoryViewControllers:"), objectivec.IObjectSliceToNSArray(value))
}

// The view controller that the split view item represents.
//
// # Discussion
// 
// Don’t set this property while adding the split view item to a split view
// controller. If you do, the system raises an exception.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/viewController
func (s NSSplitViewItem) ViewController() INSViewController {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("viewController"))
	return NSViewControllerFromID(objc.ID(rv))
}
func (s NSSplitViewItem) SetViewController(value INSViewController) {
	objc.Send[struct{}](s.ID, objc.Sel("setViewController:"), value)
}

// Sets the option dictionary that maps event trigger keys to animation
// objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/animations
func (s NSSplitViewItem) Animations() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("animations"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (s NSSplitViewItem) SetAnimations(value foundation.INSDictionary) {
	objc.Send[struct{}](s.ID, objc.Sel("setAnimations:"), value)
}

