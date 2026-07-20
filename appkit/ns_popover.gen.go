// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPopover] class.
var (
	_NSPopoverClass     NSPopoverClass
	_NSPopoverClassOnce sync.Once
)

func getNSPopoverClass() NSPopoverClass {
	_NSPopoverClassOnce.Do(func() {
		_NSPopoverClass = NSPopoverClass{class: objc.GetClass("NSPopover")}
	})
	return _NSPopoverClass
}

// GetNSPopoverClass returns the class object for NSPopover.
func GetNSPopoverClass() NSPopoverClass {
	return getNSPopoverClass()
}

type NSPopoverClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPopoverClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPopoverClass) Alloc() NSPopover {
	rv := objc.Send[NSPopover](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A means to display additional content related to existing content on the
// screen.
//
// # Overview
//
// The popover is positioned relative to the existing content and an anchor is
// used to express the relation between these two units of content. A popover
// has an appearance that specifies its visual characteristics, as well as a
// behavior that determines which user interactions will cause the popover to
// close. A transient popover is closed in response to most user interactions,
// whereas a semi-transient popover is closed when the user interacts with the
// window containing the popover’s positioning view. Popovers with
// application-defined behavior are not usually closed on the developer’s
// behalf.
//
// The system automatically positions each popover relative to its positioning
// view and moves the popover whenever its positioning view moves. A
// positioning rectangle within the positioning view can be specified for
// additional granularity.
//
// Popovers can be detached to become a separate window when they are dragged
// by implementing the appropriate delegate method.
//
// # Accessing a Popover’s Content View Controller
//
//   - [NSPopover.ContentViewController]: The view controller that manages the content of the popover.
//   - [NSPopover.SetContentViewController]
//
// # Managing a Popover’s Position and Size
//
//   - [NSPopover.Behavior]: Specifies the behavior of the popover.
//   - [NSPopover.SetBehavior]
//   - [NSPopover.ShowRelativeToRectOfViewPreferredEdge]: Shows the popover anchored to the specified view.
//   - [NSPopover.PositioningRect]: The rectangle within the positioning view relative to which the popover should be positioned.
//   - [NSPopover.SetPositioningRect]
//
// # Managing a Popover’s Appearance
//
//   - [NSPopover.Animates]: Specifies if the popover is to be animated.
//   - [NSPopover.SetAnimates]
//   - [NSPopover.ContentSize]: The content size of the popover.
//   - [NSPopover.SetContentSize]
//   - [NSPopover.IsShown]: The display state of the popover.
//   - [NSPopover.IsDetached]: A Boolean value that indicates whether the window created by a popover’s detachment is automatically created.
//
// # Closing a Popover
//
//   - [NSPopover.PerformClose]: Attempts to close the popover.
//   - [NSPopover.Close]: Forces the popover to close without consulting its delegate.
//
// # Getting and Setting the Delegate
//
//   - [NSPopover.Delegate]: The delegate of the popover.
//   - [NSPopover.SetDelegate]
//
// # Instance Properties
//
//   - [NSPopover.HasFullSizeContent]: A Boolean value that indicates whether the content view of the popover extends into the arrow region.
//   - [NSPopover.SetHasFullSizeContent]
//
// # Instance Methods
//
//   - [NSPopover.ShowRelativeToToolbarItem]: Shows the popover anchored to the specified toolbar item.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover
type NSPopover struct {
	NSResponder
}

// NSPopoverFromID constructs a [NSPopover] from an objc.ID.
//
// A means to display additional content related to existing content on the
// screen.
func NSPopoverFromID(id objc.ID) NSPopover {
	return NSPopover{NSResponder: NSResponderFromID(id)}
}

// NOTE: NSPopover adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPopover] class.
//
// # Accessing a Popover’s Content View Controller
//
//   - [INSPopover.ContentViewController]: The view controller that manages the content of the popover.
//   - [INSPopover.SetContentViewController]
//
// # Managing a Popover’s Position and Size
//
//   - [INSPopover.Behavior]: Specifies the behavior of the popover.
//   - [INSPopover.SetBehavior]
//   - [INSPopover.ShowRelativeToRectOfViewPreferredEdge]: Shows the popover anchored to the specified view.
//   - [INSPopover.PositioningRect]: The rectangle within the positioning view relative to which the popover should be positioned.
//   - [INSPopover.SetPositioningRect]
//
// # Managing a Popover’s Appearance
//
//   - [INSPopover.Animates]: Specifies if the popover is to be animated.
//   - [INSPopover.SetAnimates]
//   - [INSPopover.ContentSize]: The content size of the popover.
//   - [INSPopover.SetContentSize]
//   - [INSPopover.IsShown]: The display state of the popover.
//   - [INSPopover.IsDetached]: A Boolean value that indicates whether the window created by a popover’s detachment is automatically created.
//
// # Closing a Popover
//
//   - [INSPopover.PerformClose]: Attempts to close the popover.
//   - [INSPopover.Close]: Forces the popover to close without consulting its delegate.
//
// # Getting and Setting the Delegate
//
//   - [INSPopover.Delegate]: The delegate of the popover.
//   - [INSPopover.SetDelegate]
//
// # Instance Properties
//
//   - [INSPopover.HasFullSizeContent]: A Boolean value that indicates whether the content view of the popover extends into the arrow region.
//   - [INSPopover.SetHasFullSizeContent]
//
// # Instance Methods
//
//   - [INSPopover.ShowRelativeToToolbarItem]: Shows the popover anchored to the specified toolbar item.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover
type INSPopover interface {
	INSResponder
	NSAppearanceCustomization

	// Topic: Accessing a Popover’s Content View Controller

	// The view controller that manages the content of the popover.
	ContentViewController() INSViewController
	SetContentViewController(value INSViewController)

	// Topic: Managing a Popover’s Position and Size

	// Specifies the behavior of the popover.
	Behavior() NSPopoverBehavior
	SetBehavior(value NSPopoverBehavior)
	// Shows the popover anchored to the specified view.
	ShowRelativeToRectOfViewPreferredEdge(positioningRect corefoundation.CGRect, positioningView INSView, preferredEdge foundation.NSRectEdge)
	// The rectangle within the positioning view relative to which the popover should be positioned.
	PositioningRect() corefoundation.CGRect
	SetPositioningRect(value corefoundation.CGRect)

	// Topic: Managing a Popover’s Appearance

	// Specifies if the popover is to be animated.
	Animates() bool
	SetAnimates(value bool)
	// The content size of the popover.
	ContentSize() corefoundation.CGSize
	SetContentSize(value corefoundation.CGSize)
	// The display state of the popover.
	IsShown() bool
	// A Boolean value that indicates whether the window created by a popover’s detachment is automatically created.
	IsDetached() bool

	// Topic: Closing a Popover

	// Attempts to close the popover.
	PerformClose(sender objectivec.IObject)
	// Forces the popover to close without consulting its delegate.
	Close()

	// Topic: Getting and Setting the Delegate

	// The delegate of the popover.
	Delegate() NSPopoverDelegate
	SetDelegate(value NSPopoverDelegate)

	// Topic: Instance Properties

	// A Boolean value that indicates whether the content view of the popover extends into the arrow region.
	HasFullSizeContent() bool
	SetHasFullSizeContent(value bool)

	// Topic: Instance Methods

	// Shows the popover anchored to the specified toolbar item.
	ShowRelativeToToolbarItem(toolbarItem INSToolbarItem)
}

// Init initializes the instance.
func (p NSPopover) Init() NSPopover {
	rv := objc.Send[NSPopover](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPopover) Autorelease() NSPopover {
	rv := objc.Send[NSPopover](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPopover creates a new NSPopover instance.
func NewNSPopover() NSPopover {
	class := getNSPopoverClass()
	rv := objc.Send[NSPopover](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSPopover/init(coder:)
func NewPopoverWithCoder(coder foundation.INSCoder) NSPopover {
	instance := getNSPopoverClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSPopoverFromID(rv)
}

// Shows the popover anchored to the specified view.
//
// positioningRect: The rectangle within `positioningView` relative to which the popover should
// be positioned. Normally set to the bounds of `positioningView`. May be an
// empty rectangle, which will default to the bounds of `positioningView`.
//
// positioningView: The view relative to which the popover should be positioned. Causes the
// method to raise [invalidArgumentException] if `nil`.
//
// preferredEdge: The edge of `positioningView` the popover should prefer to be anchored to.
//
// # Discussion
//
// This method raises [internalInconsistencyException] if
// [NSPopover.ContentViewController] or the view controller’s view is `nil`.
// If the popover is already being shown, this method updates the anchored
// view, rectangle, and preferred edge. If the positioning view is not
// visible, this method does nothing.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/show(relativeTo:of:preferredEdge:)
//
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
// [internalInconsistencyException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/internalInconsistencyException
func (p NSPopover) ShowRelativeToRectOfViewPreferredEdge(positioningRect corefoundation.CGRect, positioningView INSView, preferredEdge foundation.NSRectEdge) {
	objc.Send[objc.ID](p.ID, objc.Sel("showRelativeToRect:ofView:preferredEdge:"), positioningRect, positioningView, preferredEdge)
}

// Attempts to close the popover.
//
// sender: The sender of the action message.
//
// # Discussion
//
// The popover will not be closed if it has a delegate and the delegate
// implements the returns [PopoverShouldClose] method returning false, or if a
// subclass of the NSPopover class implements “ and returns false).
//
// The operation will fail if the popover is displaying a nested popover or if
// it has a child window. A window will attempt to close its popovers when it
// receives a [NSWindow.PerformClose] message.
//
// The popover animates out when closed unless the [NSPopover.Animates]
// property is set to false.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/performClose(_:)
func (p NSPopover) PerformClose(sender objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("performClose:"), sender)
}

// Forces the popover to close without consulting its delegate.
//
// # Discussion
//
// Any popovers nested within the popovers will also receive a
// [NSPopover.Close] message. When a window is closed in response to the
// [NSWindow.Close] message being sent, all of its popovers are closed. The
// popover animates out when closed unless the [NSPopover.Animates] property
// is set to false.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/close()
func (p NSPopover) Close() {
	objc.Send[objc.ID](p.ID, objc.Sel("close"))
}

// Shows the popover anchored to the specified toolbar item.
//
// toolbarItem: The toolbar item anchoring the popover.
//
// # Discussion
//
// Use this method to display a popover relative to a toolbar item. When the
// item is in the overflow menu, the popover presents itself from another
// appropriate affordance in the window. See
// [NSPopover.ShowRelativeToRectOfViewPreferredEdge] for popover behavior.
//
// This method raises an [invalidArgumentException] if it can’t locate the
// toolbar item. This could occur if the item isn’t in a toolbar, or because
// the toolbar isn’t in the window.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/show(relativeTo:)
//
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
func (p NSPopover) ShowRelativeToToolbarItem(toolbarItem INSToolbarItem) {
	objc.Send[objc.ID](p.ID, objc.Sel("showRelativeToToolbarItem:"), toolbarItem)
}

// The view controller that manages the content of the popover.
//
// # Discussion
//
// You must set the content view controller of the popover before the popover
// is shown. Changes to the popover’s content view controller while the
// popover is shown will cause the popover to animate if the
// [NSPopover.Animates] property is true.
//
// The default value is `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/contentViewController
func (p NSPopover) ContentViewController() INSViewController {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("contentViewController"))
	return NSViewControllerFromID(objc.ID(rv))
}
func (p NSPopover) SetContentViewController(value INSViewController) {
	objc.Send[struct{}](p.ID, objc.Sel("setContentViewController:"), value)
}

// Specifies the behavior of the popover.
//
// # Discussion
//
// The default value is [NSPopoverBehaviorApplicationDefined]. See
// [NSPopover.Behavior] for possible value.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/behavior-swift.property
//
// [NSPopover.Behavior]: https://developer.apple.com/documentation/AppKit/NSPopover/Behavior-swift.enum
func (p NSPopover) Behavior() NSPopoverBehavior {
	rv := objc.Send[NSPopoverBehavior](p.ID, objc.Sel("behavior"))
	return NSPopoverBehavior(rv)
}
func (p NSPopover) SetBehavior(value NSPopoverBehavior) {
	objc.Send[struct{}](p.ID, objc.Sel("setBehavior:"), value)
}

// The rectangle within the positioning view relative to which the popover
// should be positioned.
//
// # Discussion
//
// Popovers are positioned relative to a positioning view and are
// automatically moved when the location or size of the positioning view
// changes.
//
// Sometimes it is desirable to position popovers relative to a rectangle
// within the positioning view. In this case, you must update the
// `positioningRect` property whenever this rectangle changes.
//
// This property is exposed as a read-only binding.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/positioningRect
func (p NSPopover) PositioningRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p.ID, objc.Sel("positioningRect"))
	return corefoundation.CGRect(rv)
}
func (p NSPopover) SetPositioningRect(value corefoundation.CGRect) {
	objc.Send[struct{}](p.ID, objc.Sel("setPositioningRect:"), value)
}

// The appearance of the popover.
//
// # Discussion
//
// If no appearance is specified, the popover’s effective appearance
// defaults to [vibrantLight].
//
// In apps that run in macOS 10.10 and later, the previous property type of
// [NSPopover.Appearance] is deprecated. In apps that run in OS X v10.9 and
// earlier, the [aqua] appearance is automatically set on popover content.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/appearance-swift.property
//
// [NSPopover.Appearance]: https://developer.apple.com/documentation/AppKit/NSPopover/Appearance-swift.enum
// [aqua]: https://developer.apple.com/documentation/AppKit/NSAppearance/Name-swift.struct/aqua
// [vibrantLight]: https://developer.apple.com/documentation/AppKit/NSAppearance/Name-swift.struct/vibrantLight
func (p NSPopover) Appearance() INSAppearance {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("appearance"))
	return NSAppearanceFromID(objc.ID(rv))
}
func (p NSPopover) SetAppearance(value INSAppearance) {
	objc.Send[struct{}](p.ID, objc.Sel("setAppearance:"), value)
}

// The appearance that will be used when the popover is displayed onscreen.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/effectiveAppearance
func (p NSPopover) EffectiveAppearance() INSAppearance {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("effectiveAppearance"))
	return NSAppearanceFromID(objc.ID(rv))
}

// Specifies if the popover is to be animated.
//
// # Discussion
//
// A popover may be animated when it shows, closes, moves, or appears to
// transition to a detachable window. This property also controls whether the
// popover animates when the content view or content size changes.
//
// The system does not guarantee which behaviors will be animated or that this
// property will be respected; it is regarded as a hint.
//
// The default value is true.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/animates
func (p NSPopover) Animates() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("animates"))
	return rv
}
func (p NSPopover) SetAnimates(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAnimates:"), value)
}

// The content size of the popover.
//
// # Discussion
//
// The popover’s content size is set to match the size of the content view
// when the content view controller is set.
//
// Changes to the content size of the popover will cause the popover to
// animate while it is shown if the [NSPopover.Animates] property is true.
//
// This property is exposed as a read-only binding.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/contentSize
func (p NSPopover) ContentSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](p.ID, objc.Sel("contentSize"))
	return corefoundation.CGSize(rv)
}
func (p NSPopover) SetContentSize(value corefoundation.CGSize) {
	objc.Send[struct{}](p.ID, objc.Sel("setContentSize:"), value)
}

// The display state of the popover.
//
// # Discussion
//
// The value is true if the popover is being shown, false otherwise.
//
// The popover is considered to be shown from the point when
// [NSPopover.ShowRelativeToRectOfViewPreferredEdge] is invoked. A popover is
// considered closed in response to an invocation of either [NSPopover.Close]
// or [NSPopover.PerformClose].
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/isShown
func (p NSPopover) IsShown() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isShown"))
	return rv
}

// A Boolean value that indicates whether the window created by a popover’s
// detachment is automatically created.
//
// # Discussion
//
// When [NSPopover.Detached] is true, the detached window is automatically
// created. This property does not apply when detaching a popover results in a
// window returned by [DetachableWindowForPopover].
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/isDetached
func (p NSPopover) IsDetached() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isDetached"))
	return rv
}

// The delegate of the popover.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/delegate
func (p NSPopover) Delegate() NSPopoverDelegate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("delegate"))
	return NSPopoverDelegateObjectFromID(rv)
}
func (p NSPopover) SetDelegate(value NSPopoverDelegate) {
	objc.Send[struct{}](p.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates whether the content view of the popover
// extends into the arrow region.
//
// # Discussion
//
// Setting the value of this property to true extends the frame of the content
// view by the height of the arrow region on all four sides of the frame. This
// causes the [NSPopover.ContentViewController] view to extend to the
// window’s bounds.
//
// [media-4304810]
//
// Use the [NSView.SafeAreaLayoutGuide] of the
// [NSPopover.ContentViewController] view to ensure that your content is fully
// visible and doesn’t become clipped when displayed.
//
// [media-4304811]
//
// Setting this value to false doesn’t extend the
// [NSPopover.ContentViewController] view fully into the arrow region. The
// default value for this property is false.
//
// [media-4304812]
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/hasFullSizeContent
func (p NSPopover) HasFullSizeContent() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("hasFullSizeContent"))
	return rv
}
func (p NSPopover) SetHasFullSizeContent(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setHasFullSizeContent:"), value)
}

// Protocol methods for NSAppearanceCustomization
