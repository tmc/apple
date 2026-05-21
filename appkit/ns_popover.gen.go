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

// Protocol methods for NSAccessibilityElementProtocol

// Returns the accessibility element’s frame in screen coordinates.
//
// # Return Value
//
// The element’s frame in screen coordinates.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityFrame] property. This method is called whenever accessibility
// clients request the [size] or [position] attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityFrame()
//
// [accessibilityFrame]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrame
// [position]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/position
// [size]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/size
func (o NSPopover) AccessibilityFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("accessibilityFrame"))
	return rv
}

// Returns the accessibility element’s parent in the accessibility
// hierarchy.
//
// # Return Value
//
// The element’s parent in the accessibility hierarchy.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityParent] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityParent()
//
// [accessibilityParent]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityParent
func (o NSPopover) AccessibilityParent() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityParent"))
	return objectivec.Object{ID: rv}
}

// Returns the accessibility element’s identity.
//
// # Return Value
//
// Returns the unique ID for the accessibility element. It is often used in
// automated testing.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityIdentifier] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityIdentifier()
//
// [accessibilityIdentifier]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIdentifier
func (o NSPopover) AccessibilityIdentifier() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// Returns a Boolean value that indicates whether the accessibility element
// has the keyboard focus.
//
// # Return Value
//
// true if this element has the keyboard focus; otherwise, false.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityFocused] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/isAccessibilityFocused()
//
// [accessibilityFocused]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFocused
func (o NSPopover) IsAccessibilityFocused() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
}

// Protocol methods for NSAccessibilityProtocol

// Returns a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityElement()
func (o NSPopover) IsAccessibilityElement() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityElement"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityEnabled()
func (o NSPopover) IsAccessibilityEnabled() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEnabled"))
	return rv
}

// Returns a Boolean value that indicates whether assistive apps can invoke
// the specified selector on the accessibility element.
//
// selector: The selector to check.
//
// # Return Value
//
// true, if accessibility clients can call the selector; otherwise, false.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilitySelectorAllowed(_:)
func (o NSPopover) IsAccessibilitySelectorAllowed(selector objc.SEL) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelectorAllowed:"), selector)
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityProtectedContent()
func (o NSPopover) IsAccessibilityProtectedContent() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityProtectedContent"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// is currently in a selected state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilitySelected()
func (o NSPopover) IsAccessibilitySelected() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelected"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// must have content for successful submission of a form.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityRequired()
func (o NSPopover) IsAccessibilityRequired() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityRequired"))
	return rv
}

// Returns the substring for the specified range.
//
// range: A range of characters contained by the element.
//
// # Return Value
//
// The substring specified by the given range.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityString(for:)
func (o NSPopover) AccessibilityStringForRange(range_ foundation.NSRange) string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityStringForRange:"), range_)
	return foundation.NSStringFromID(rv).String()
}

// Returns the attributed substring for the specified range of characters.
//
// range: The range of characters.
//
// # Return Value
//
// An attributed string representing the specified characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityAttributedString(for:)
func (o NSPopover) AccessibilityAttributedStringForRange(range_ foundation.NSRange) foundation.NSAttributedString {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAttributedStringForRange:"), range_)
	return foundation.NSAttributedStringFromID(rv)
}

// Returns the rich text format (RTF) data that describes the specified range
// of characters.
//
// range: The range of characters.
//
// # Return Value
//
// A data object containing an RTF representation of the specified characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRTF(for:)
func (o NSPopover) AccessibilityRTFForRange(range_ foundation.NSRange) foundation.NSData {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRTFForRange:"), range_)
	return foundation.NSDataFromID(rv)
}

// Returns the rectangle that encloses the specified range of characters.
//
// range: The range of characters.
//
// # Return Value
//
// The rectangle that encloses the specified characters.
//
// # Discussion
//
// If the range crosses a line boundary, the returned rectangle fully encloses
// all the lines of characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFrame(for:)
func (o NSPopover) AccessibilityFrameForRange(range_ foundation.NSRange) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("accessibilityFrameForRange:"), range_)
	return rv
}

// Returns the line number for the line that contains the specified character
// index.
//
// index: The index for a character.
//
// # Return Value
//
// The line number for the line holding the specified character index.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLine(for:)
func (o NSPopover) AccessibilityLineForIndex(index int) int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityLineForIndex:"), index)
	return rv
}

// Returns the range of characters for the glyph that includes the specified
// character.
//
// index: The specified character.
//
// # Return Value
//
// The range of characters for the glyph.
//
// # Discussion
//
// This value always includes the specified character but may include
// additional characters if that character is part of a multicharacter glyph.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRange(for:)-6kv3
func (o NSPopover) AccessibilityRangeForIndex(index int) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForIndex:"), index)
	return rv
}

// Returns a range of characters that all have the same style as the specified
// character.
//
// index: The index of the specified character.
//
// # Return Value
//
// A range of characters with the same style as the specified character.
//
// # Discussion
//
// This method returns a range of characters that meet two conditions: The
// range must include the specified character, and all the other characters in
// the range must match the specified character’s style. If none of the
// adjacent characters match the specified character’s style, the method
// returns only the specified character.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityStyleRange(for:)
func (o NSPopover) AccessibilityStyleRangeForIndex(index int) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityStyleRangeForIndex:"), index)
	return rv
}

// Returns the range of characters in the specified line.
//
// line: The line number to be examined.
//
// # Return Value
//
// The range of characters for the specified line number. If the line ends
// with a newline character, including the newline is preferred.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRange(forLine:)
func (o NSPopover) AccessibilityRangeForLine(line int) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForLine:"), line)
	return rv
}

// Returns the range of characters for the glyph at the specified point.
//
// point: A point in screen coordinates.
//
// # Return Value
//
// The range of characters that make up the glyph at the given point.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRange(for:)-1iudm
func (o NSPopover) AccessibilityRangeForPosition(point corefoundation.CGPoint) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForPosition:"), point)
	return rv
}

// Returns the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityAlternateUIVisible()
func (o NSPopover) IsAccessibilityAlternateUIVisible() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityAlternateUIVisible"))
	return rv
}

// Returns a Boolean value that determines whether the window is the app’s
// main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityMain()
func (o NSPopover) IsAccessibilityMain() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMain"))
	return rv
}

// Returns the Boolean value that determines whether the window is in a
// minimized state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityMinimized()
func (o NSPopover) IsAccessibilityMinimized() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMinimized"))
	return rv
}

// Returns a Boolean value that determines whether the window is modal.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityModal()
func (o NSPopover) IsAccessibilityModal() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityModal"))
	return rv
}

// Returns a Boolean value that determines whether the app is the frontmost
// app.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityFrontmost()
func (o NSPopover) IsAccessibilityFrontmost() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFrontmost"))
	return rv
}

// Returns a Boolean value that determines whether the app is in a hidden
// state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityHidden()
func (o NSPopover) IsAccessibilityHidden() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityHidden"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility
// element’s grid is in row major order or in column major order.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityOrderedByRow()
func (o NSPopover) IsAccessibilityOrderedByRow() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityOrderedByRow"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// is in an expanded state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityExpanded()
func (o NSPopover) IsAccessibilityExpanded() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityExpanded"))
	return rv
}

// Returns a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityDisclosed()
func (o NSPopover) IsAccessibilityDisclosed() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityDisclosed"))
	return rv
}

// Returns the cell at the specified column and row.
//
// column: The column index.
//
// row: The row index.
//
// # Return Value
//
// The cell specified by the column and row indexes.
//
// # Discussion
//
// This property is required for all elements that function as cell-based
// tables.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCell(forColumn:row:)
func (o NSPopover) AccessibilityCellForColumnRow(column int, row int) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCellForColumn:row:"), column, row)
	return objectivec.Object{ID: rv}
}

// Converts the provided point in screen coordinates to a point in the layout
// area’s coordinate system.
//
// point: A point in the screen’s coordinate system.
//
// # Return Value
//
// A point in the layout area’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLayoutPoint(forScreenPoint:)
func (o NSPopover) AccessibilityLayoutPointForScreenPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityLayoutPointForScreenPoint:"), point)
	return rv
}

// Converts the provided size in screen coordinates to a size in the layout
// area’s coordinate system.
//
// size: A size in the screen’s coordinate system.
//
// # Return Value
//
// A size in the layout area’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLayoutSize(forScreenSize:)
func (o NSPopover) AccessibilityLayoutSizeForScreenSize(size corefoundation.CGSize) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("accessibilityLayoutSizeForScreenSize:"), size)
	return rv
}

// Converts the provided point in the layout area’s coordinates to a point
// in the screen’s coordinate system.
//
// point: A point in the layout area’s coordinate system.
//
// # Return Value
//
// A point in the screen’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityScreenPoint(forLayoutPoint:)
func (o NSPopover) AccessibilityScreenPointForLayoutPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityScreenPointForLayoutPoint:"), point)
	return rv
}

// Converts the provided size in the layout area’s coordinates to a size in
// the screen’s coordinate system.
//
// size: A size in the layout area’s coordinate system.
//
// # Return Value
//
// A size in the screen’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityScreenSize(forLayoutSize:)
func (o NSPopover) AccessibilityScreenSizeForLayoutSize(size corefoundation.CGSize) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("accessibilityScreenSizeForLayoutSize:"), size)
	return rv
}

// Returns a Boolean value that indicates whether the accessibility element is
// in an edited state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityEdited()
func (o NSPopover) IsAccessibilityEdited() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEdited"))
	return rv
}

// Cancels the current operation.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformCancel()
func (o NSPopover) AccessibilityPerformCancel() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformCancel"))
	return rv
}

// Simulates pressing Return in the accessibility element.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method on elements that take keyboard input, such as a text field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformConfirm()
func (o NSPopover) AccessibilityPerformConfirm() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformConfirm"))
	return rv
}

// Selects the accessibility element.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method on selectable elements, such as a menu item.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformPick()
func (o NSPopover) AccessibilityPerformPick() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformPick"))
	return rv
}

// Simulates clicking the accessibility element.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method on elements that behave like buttons.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformPress()
func (o NSPopover) AccessibilityPerformPress() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformPress"))
	return rv
}

// Displays the accessibility element’s alternative UI.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method to trigger changes to the UI due to a mouse-hover or
// similar event.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowAlternateUI()
func (o NSPopover) AccessibilityPerformShowAlternateUI() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowAlternateUI"))
	return rv
}

// Returns to the accessibility element’s original UI.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Call this method after successfully calling
// [AccessibilityPerformShowAlternateUI] to return to the original UI.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowDefaultUI()
func (o NSPopover) AccessibilityPerformShowDefaultUI() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowDefaultUI"))
	return rv
}

// Displays the menu accessibility element.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method to display the contextual menu for the element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowMenu()
func (o NSPopover) AccessibilityPerformShowMenu() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowMenu"))
	return rv
}

// Brings the window to the front.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// The window behaves as if you had clicked on the window’s title bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformRaise()
func (o NSPopover) AccessibilityPerformRaise() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformRaise"))
	return rv
}

// Increments the accessibility element’s value.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method on elements that have an adjustable [accessibilityValue]
// property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformIncrement()
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
func (o NSPopover) AccessibilityPerformIncrement() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformIncrement"))
	return rv
}

// Decrements the accessibility element’s value.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method on elements that have an adjustable [accessibilityValue]
// property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformDecrement()
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
func (o NSPopover) AccessibilityPerformDecrement() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformDecrement"))
	return rv
}

// Deletes the accessibility element’s value.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method on elements with values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformDelete()
func (o NSPopover) AccessibilityPerformDelete() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformDelete"))
	return rv
}

// The activation point for the user interface element.
//
// # Discussion
//
// The activation point in screen coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityActivationPoint
func (o NSPopover) AccessibilityActivationPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityActivationPoint"))
	return corefoundation.CGPoint(rv)
}

func (o NSPopover) SetAccessibilityActivationPoint(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityActivationPoint:"), value)
}

// The allowed values for the slider accessibility element.
//
// # Discussion
//
// Use this property if the slider can be set only to predefined values (for
// example, if the slider’s level indicator automatically snaps to the
// closest integer values between 0 and 100).
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityAllowedValues
func (o NSPopover) AccessibilityAllowedValues() []foundation.NSNumber {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityAllowedValues"))
	result := make([]foundation.NSNumber, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = foundation.NSNumberFromID(id)
	}
	return result
}

func (o NSPopover) SetAccessibilityAllowedValues(value []foundation.NSNumber) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAllowedValues:"), objectivec.IObjectSliceToNSArray(value))
}

// A Boolean value that determines whether the accessibility element’s
// alternative UI is currently visible.
//
// # Discussion
//
// Use this property for elements that present an alternative UI—for
// example, when the pointer hovers over an interface element for a few
// seconds.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityAlternateUIVisible
func (o NSPopover) AccessibilityAlternateUIVisible() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityAlternateUIVisible"))
	return bool(rv)
}

func (o NSPopover) SetAccessibilityAlternateUIVisible(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAlternateUIVisible:"), value)
}

// The child accessibility element with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityApplicationFocusedUIElement
func (o NSPopover) AccessibilityApplicationFocusedUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityApplicationFocusedUIElement"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityApplicationFocusedUIElement(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityApplicationFocusedUIElement:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityAttributedUserInputLabels
func (o NSPopover) AccessibilityAttributedUserInputLabels() []foundation.NSAttributedString {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityAttributedUserInputLabels"))
	result := make([]foundation.NSAttributedString, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = foundation.NSAttributedStringFromID(id)
	}
	return result
}

func (o NSPopover) SetAccessibilityAttributedUserInputLabels(value []foundation.NSAttributedString) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAttributedUserInputLabels:"), objectivec.IObjectSliceToNSArray(value))
}

// The child accessibility element that represents the window’s cancel
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityCancelButton
func (o NSPopover) AccessibilityCancelButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCancelButton"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityCancelButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCancelButton:"), value)
}

// The child accessibility elements in the accessibility hierarchy.
//
// # Discussion
//
// This property contains references to child elements in the accessibility
// hierarchy. If you create an [NSView] subclass, you don’t typically need
// to set this value. The system automatically populates the
// `accessibilityChildren` property with descendants in the view hierarchy
// that are also in the accessibility hierarchy. If you use an
// [NSAccessibilityElement] subclass to represent an interface element that is
// not backed by a view, you can either set the `accessibilityChildren`
// property or you can call the
// [NSAccessibilityElement.AccessibilityAddChildElement] convenience method.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityChildren
func (o NSPopover) AccessibilityChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityChildren"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilityChildren(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChildren:"), value)
}

// An array of child accessibility elements in order for linear navigation.
//
// # Discussion
//
// The array should match all elements found in [accessibilityChildren],
// rearranged in an easily navigable order.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityChildrenInNavigationOrder
//
// [accessibilityChildren]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityChildren
func (o NSPopover) AccessibilityChildrenInNavigationOrder() []objectivec.IObject {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityChildrenInNavigationOrder"))
	result := make([]objectivec.IObject, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = objectivec.Object{ID: id}
	}
	return result
}

func (o NSPopover) SetAccessibilityChildrenInNavigationOrder(value []objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChildrenInNavigationOrder:"), objectivec.IObjectSliceToNSArray(value))
}

// The clear button for the search field.
//
// # Discussion
//
// Use this property on a search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityClearButton
func (o NSPopover) AccessibilityClearButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityClearButton"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityClearButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityClearButton:"), value)
}

// The child accessibility element that represents the window’s close
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityCloseButton
func (o NSPopover) AccessibilityCloseButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCloseButton"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityCloseButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCloseButton:"), value)
}

// The number of columns in the accessibility element’s grid.
//
// # Discussion
//
// Use this property for UI elements that present a grid of child elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumnCount
func (o NSPopover) AccessibilityColumnCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityColumnCount"))
	return int(rv)
}

func (o NSPopover) SetAccessibilityColumnCount(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnCount:"), value)
}

// The column header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumnHeaderUIElements
func (o NSPopover) AccessibilityColumnHeaderUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnHeaderUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilityColumnHeaderUIElements(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnHeaderUIElements:"), value)
}

// The column index range of the cell.
//
// # Discussion
//
// This property contains the column’s starting index and index span in the
// table. Use this property in the elements representing a table’s cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumnIndexRange
func (o NSPopover) AccessibilityColumnIndexRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityColumnIndexRange"))
	return foundation.NSRange(rv)
}

func (o NSPopover) SetAccessibilityColumnIndexRange(value foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnIndexRange:"), value)
}

// The column titles for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumnTitles
func (o NSPopover) AccessibilityColumnTitles() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnTitles"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilityColumnTitles(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnTitles:"), value)
}

// The column accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumns
func (o NSPopover) AccessibilityColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumns"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilityColumns(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumns:"), value)
}

// The contents of the current accessibility element.
//
// # Discussion
//
// This property is used by container elements. It holds an array of the
// container’s contents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityContents
func (o NSPopover) AccessibilityContents() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityContents"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilityContents(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityContents:"), value)
}

// The critical value for the level indicator.
//
// # Discussion
//
// Use this property for elements such as the battery level indicator. This
// property sets a boundary value. If the element’s value exceeds the
// boundary value, the element has reached a critical stage.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityCriticalValue
func (o NSPopover) AccessibilityCriticalValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCriticalValue"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityCriticalValue(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCriticalValue:"), value)
}

// The custom actions of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityCustomActions
func (o NSPopover) AccessibilityCustomActions() []NSAccessibilityCustomAction {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityCustomActions"))
	result := make([]NSAccessibilityCustomAction, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = NSAccessibilityCustomActionFromID(id)
	}
	return result
}

func (o NSPopover) SetAccessibilityCustomActions(value []NSAccessibilityCustomAction) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomActions:"), objectivec.IObjectSliceToNSArray(value))
}

// The custom rotors of the current accessibility element.
//
// # Discussion
//
// Custom rotors are lists of items of a specific category. For example, a
// “headings” rotor returns a list of headings a given document.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityCustomRotors
func (o NSPopover) AccessibilityCustomRotors() []NSAccessibilityCustomRotor {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityCustomRotors"))
	result := make([]NSAccessibilityCustomRotor, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = NSAccessibilityCustomRotorFromID(id)
	}
	return result
}

func (o NSPopover) SetAccessibilityCustomRotors(value []NSAccessibilityCustomRotor) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomRotors:"), objectivec.IObjectSliceToNSArray(value))
}

// The decrement button for the stepper accessibility element.
//
// # Discussion
//
// Use this property on a stepper.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDecrementButton
func (o NSPopover) AccessibilityDecrementButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDecrementButton"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityDecrementButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDecrementButton:"), value)
}

// The child accessibility element that represents the window’s default
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDefaultButton
func (o NSPopover) AccessibilityDefaultButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDefaultButton"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityDefaultButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDefaultButton:"), value)
}

// A Boolean value that determines whether the row is disclosing other rows.
//
// # Discussion
//
// Use this property in the elements representing an outline’s row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDisclosed
func (o NSPopover) AccessibilityDisclosed() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityDisclosed"))
	return bool(rv)
}

func (o NSPopover) SetAccessibilityDisclosed(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosed:"), value)
}

// The row disclosing the current row.
//
// # Discussion
//
// Use this property in the elements representing an outline’s row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDisclosedByRow
func (o NSPopover) AccessibilityDisclosedByRow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedByRow"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityDisclosedByRow(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedByRow:"), value)
}

// The rows that the current row discloses.
//
// # Discussion
//
// Use this property in the elements representing an outline’s row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDisclosedRows
func (o NSPopover) AccessibilityDisclosedRows() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedRows"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityDisclosedRows(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedRows:"), value)
}

// The indention level for the row.
//
// # Discussion
//
// Use this property in the elements representing an outline’s row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDisclosureLevel
func (o NSPopover) AccessibilityDisclosureLevel() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityDisclosureLevel"))
	return int(rv)
}

func (o NSPopover) SetAccessibilityDisclosureLevel(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosureLevel:"), value)
}

// The URL for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDocument
func (o NSPopover) AccessibilityDocument() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDocument"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSPopover) SetAccessibilityDocument(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDocument:"), objc.String(value))
}

// A Boolean value that indicates whether the accessibility element is in an
// edited state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityEdited
func (o NSPopover) AccessibilityEdited() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEdited"))
	return bool(rv)
}

func (o NSPopover) SetAccessibilityEdited(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityEdited:"), value)
}

// A Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// # Discussion
//
// Use this property to expose this object to accessibility clients as a
// functional interface element. For example, when you place a button in a
// window, the system typically creates a button cell inside a button control
// inside a container view inside a window. Users, however, don’t care about
// the view hierarchy details. They should only be told that there’s a
// button in a window.
//
// If this property is set to false, accessibility clients ignore this
// element. By default, [NSView] and its subclasses set this value to false;
// however, if your [NSView] subclass adopts one of the accessibility
// protocols, the system changes the default value to true.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityElement
func (o NSPopover) AccessibilityElement() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityElement"))
	return bool(rv)
}

func (o NSPopover) SetAccessibilityElement(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityElement:"), value)
}

// A Boolean value that determines whether the accessibility element responds
// to user events.
//
// # Discussion
//
// Returns YES if the element is enabled; otherwise, NO. Enabled elements
// respond to user events.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityEnabled
func (o NSPopover) AccessibilityEnabled() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEnabled"))
	return bool(rv)
}

func (o NSPopover) SetAccessibilityEnabled(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityEnabled:"), value)
}

// A Boolean value that determines whether the accessibility element is in an
// expanded state.
//
// # Discussion
//
// Use this property on elements that can expand to reveal additional
// information, such as outline rows and combo boxes.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityExpanded
func (o NSPopover) AccessibilityExpanded() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityExpanded"))
	return bool(rv)
}

func (o NSPopover) SetAccessibilityExpanded(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExpanded:"), value)
}

// The icon for the app’s menu bar extra.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityExtrasMenuBar
func (o NSPopover) AccessibilityExtrasMenuBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityExtrasMenuBar"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityExtrasMenuBar(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExtrasMenuBar:"), value)
}

// The filename for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFilename
func (o NSPopover) AccessibilityFilename() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFilename"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSPopover) SetAccessibilityFilename(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFilename:"), objc.String(value))
}

// A Boolean value that determines whether the accessibility element has the
// keyboard focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFocused
func (o NSPopover) AccessibilityFocused() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return bool(rv)
}

func (o NSPopover) SetAccessibilityFocused(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFocused:"), value)
}

// The child window with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFocusedWindow
func (o NSPopover) AccessibilityFocusedWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFocusedWindow"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityFocusedWindow(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFocusedWindow:"), value)
}

// The accessibility element’s frame in screen coordinates.
//
// # Discussion
//
// This property is accessed by the system whenever an accessibility client
// requests the element’s size or position.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrame
func (o NSPopover) SetAccessibilityFrame(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrame:"), value)
}

// A Boolean value that determines whether the app is the frontmost app.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrontmost
func (o NSPopover) AccessibilityFrontmost() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFrontmost"))
	return bool(rv)
}

func (o NSPopover) SetAccessibilityFrontmost(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrontmost:"), value)
}

// The child accessibility element that represents the window’s full-screen
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFullScreenButton
func (o NSPopover) AccessibilityFullScreenButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFullScreenButton"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityFullScreenButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFullScreenButton:"), value)
}

// The child accessibility element that represents the window’s grow area.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityGrowArea
func (o NSPopover) AccessibilityGrowArea() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityGrowArea"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityGrowArea(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityGrowArea:"), value)
}

// The drag handle accessibility elements for the layout item element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHandles
func (o NSPopover) AccessibilityHandles() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHandles"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilityHandles(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHandles:"), value)
}

// The header for the table view.
//
// # Discussion
//
// Use this property on a table view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHeader
func (o NSPopover) AccessibilityHeader() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHeader"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityHeader(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHeader:"), value)
}

// The help text for the accessibility element.
//
// # Discussion
//
// Use this property only when the results of activating this element are not
// obvious from the element’s label. This string functions as a tooltip. For
// example, VoiceOver reads this string when you pause over a control. To help
// ensure that accessibility clients like VoiceOver read the help text with
// the proper inflection, begin this string with a verb, capitalize the first
// letter, and end the string with a period. Always localize this string. The
// default value is `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHelp
func (o NSPopover) AccessibilityHelp() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHelp"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSPopover) SetAccessibilityHelp(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHelp:"), objc.String(value))
}

// A Boolean value that determines whether the app is in a hidden state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHidden
func (o NSPopover) AccessibilityHidden() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityHidden"))
	return bool(rv)
}

func (o NSPopover) SetAccessibilityHidden(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHidden:"), value)
}

// The horizontal scroll bar for the scroll view.
//
// # Discussion
//
// Use this property on a scrollable view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHorizontalScrollBar
func (o NSPopover) AccessibilityHorizontalScrollBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalScrollBar"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityHorizontalScrollBar(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalScrollBar:"), value)
}

// A description of the layout area’s horizontal units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHorizontalUnitDescription
func (o NSPopover) AccessibilityHorizontalUnitDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSPopover) SetAccessibilityHorizontalUnitDescription(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalUnitDescription:"), objc.String(value))
}

// The units that the layout area uses for horizontal values.
//
// # Discussion
//
// For a list of possible values, see [NSAccessibilityUnits].
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHorizontalUnits
//
// [NSAccessibilityUnits]: https://developer.apple.com/documentation/AppKit/NSAccessibilityUnits
func (o NSPopover) AccessibilityHorizontalUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityHorizontalUnits"))
	return NSAccessibilityUnits(rv)
}

func (o NSPopover) SetAccessibilityHorizontalUnits(value NSAccessibilityUnits) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalUnits:"), value)
}

// The accessibility element’s identity.
//
// # Discussion
//
// This property holds the unique ID for the accessibility element. It is
// often used in automated testing.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIdentifier
func (o NSPopover) SetAccessibilityIdentifier(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIdentifier:"), objc.String(value))
}

// The increment button for the stepper accessibility element.
//
// # Discussion
//
// Use this property on a stepper.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIncrementButton
func (o NSPopover) AccessibilityIncrementButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIncrementButton"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityIncrementButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIncrementButton:"), value)
}

// The index of the row or column that the accessibility element represents.
//
// # Discussion
//
// Use this property for any element that can be accessed through an index:
// cells, rows, columns, and so forth.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIndex
func (o NSPopover) AccessibilityIndex() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityIndex"))
	return int(rv)
}

func (o NSPopover) SetAccessibilityIndex(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIndex:"), value)
}

// The line number that contains the insertion point.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityInsertionPointLineNumber
func (o NSPopover) AccessibilityInsertionPointLineNumber() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityInsertionPointLineNumber"))
	return int(rv)
}

func (o NSPopover) SetAccessibilityInsertionPointLineNumber(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityInsertionPointLineNumber:"), value)
}

// A short description of the accessibility element.
//
// # Discussion
//
// Do not include the accessibility element’s type in the label (for
// example, write [Play], not `Play button`.). If possible, use a single word.
// To help ensure that accessibility clients such as VoiceOver read the label
// with the correct intonation, start this label with a capital letter. Do not
// put a period at the end. Always localize the label.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityLabel
func (o NSPopover) AccessibilityLabel() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabel"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSPopover) SetAccessibilityLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabel:"), objc.String(value))
}

// The child label elements for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityLabelUIElements
func (o NSPopover) AccessibilityLabelUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabelUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilityLabelUIElements(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabelUIElements:"), value)
}

// The value of the label accessibility element.
//
// # Discussion
//
// Use this property on a slider element’s labels.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityLabelValue
func (o NSPopover) AccessibilityLabelValue() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("accessibilityLabelValue"))
	return float32(rv)
}

func (o NSPopover) SetAccessibilityLabelValue(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabelValue:"), value)
}

// The elements that have links with the accessibility element.
//
// # Discussion
//
// Use this property to define a relationship between different user interface
// elements. For example, use this property to link a list item with contents
// displayed in another pane or window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityLinkedUIElements
func (o NSPopover) AccessibilityLinkedUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLinkedUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilityLinkedUIElements(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLinkedUIElements:"), value)
}

// A Boolean value that determines whether the window is the app’s main
// window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMain
func (o NSPopover) AccessibilityMain() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMain"))
	return bool(rv)
}

func (o NSPopover) SetAccessibilityMain(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMain:"), value)
}

// The app’s main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMainWindow
func (o NSPopover) AccessibilityMainWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMainWindow"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityMainWindow(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMainWindow:"), value)
}

// The user interface element that functions as a marker group for the ruler
// accessibility element.
//
// # Discussion
//
// Use this property on a ruler element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMarkerGroupUIElement
func (o NSPopover) AccessibilityMarkerGroupUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerGroupUIElement"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityMarkerGroupUIElement(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerGroupUIElement:"), value)
}

// A human-readable description of the marker type.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMarkerTypeDescription
func (o NSPopover) AccessibilityMarkerTypeDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerTypeDescription"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSPopover) SetAccessibilityMarkerTypeDescription(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerTypeDescription:"), objc.String(value))
}

// An array of marker accessibility elements for the ruler.
//
// # Discussion
//
// Use this property on a ruler element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMarkerUIElements
func (o NSPopover) AccessibilityMarkerUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilityMarkerUIElements(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerUIElements:"), value)
}

// The marker values for the ruler.
//
// # Discussion
//
// Use this property on a ruler element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMarkerValues
func (o NSPopover) AccessibilityMarkerValues() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerValues"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityMarkerValues(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerValues:"), value)
}

// The maximum value for the accessibility element.
//
// # Discussion
//
// This property is set to `nil` by default. Only a few AppKit controls (for
// example, [NSSliderCell]) support this value. Set this property only when
// the element has an [accessibilityValue] property and you want to define the
// maximum possible value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMaxValue
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
func (o NSPopover) AccessibilityMaxValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMaxValue"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityMaxValue(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMaxValue:"), value)
}

// The app’s menu bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMenuBar
func (o NSPopover) AccessibilityMenuBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMenuBar"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityMenuBar(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMenuBar:"), value)
}

// The minimum value for the accessibility element.
//
// # Discussion
//
// This property is set to `nil` by default. Only a few AppKit controls (for
// example, [NSSliderCell]) support this value. Set this property only when
// the element has an [accessibilityValue] property and you want to define the
// minimum possible value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMinValue
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
func (o NSPopover) AccessibilityMinValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinValue"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityMinValue(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinValue:"), value)
}

// The child accessibility element that represents the window’s minimize
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMinimizeButton
func (o NSPopover) AccessibilityMinimizeButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinimizeButton"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityMinimizeButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimizeButton:"), value)
}

// A Boolean value that determines whether this window is in a minimized
// state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMinimized
func (o NSPopover) AccessibilityMinimized() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMinimized"))
	return bool(rv)
}

func (o NSPopover) SetAccessibilityMinimized(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimized:"), value)
}

// A Boolean value that determines whether the window is modal.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityModal
func (o NSPopover) AccessibilityModal() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityModal"))
	return bool(rv)
}

func (o NSPopover) SetAccessibilityModal(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityModal:"), value)
}

// The contents that follow the divider accessibility element.
//
// # Discussion
//
// For example, use this property to set the subview adjacent to a split
// view’s splitter element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityNextContents
func (o NSPopover) AccessibilityNextContents() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityNextContents"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilityNextContents(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityNextContents:"), value)
}

// The number of characters in the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityNumberOfCharacters
func (o NSPopover) AccessibilityNumberOfCharacters() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityNumberOfCharacters"))
	return int(rv)
}

func (o NSPopover) SetAccessibilityNumberOfCharacters(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityNumberOfCharacters:"), value)
}

// A Boolean value that determines whether the accessibility element’s grid
// is in row major order or in column major order.
//
// # Discussion
//
// Use this property for UI elements that present a grid of child elements.
// Set the property to true if the grid is ordered row major; otherwise, set
// to false.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityOrderedByRow
func (o NSPopover) AccessibilityOrderedByRow() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityOrderedByRow"))
	return bool(rv)
}

func (o NSPopover) SetAccessibilityOrderedByRow(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOrderedByRow:"), value)
}

// The orientation of the accessibility element.
//
// # Discussion
//
// This property can hold either the [NSAccessibilityOrientationHorizontal]
// value or the [NSAccessibilityOrientationVertical] value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityOrientation
func (o NSPopover) AccessibilityOrientation() NSAccessibilityOrientation {
	rv := objc.Send[NSAccessibilityOrientation](o.ID, objc.Sel("accessibilityOrientation"))
	return NSAccessibilityOrientation(rv)
}

func (o NSPopover) SetAccessibilityOrientation(value NSAccessibilityOrientation) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOrientation:"), value)
}

// The overflow button for the toolbar.
//
// # Discussion
//
// Use this property on a toolbar element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityOverflowButton
func (o NSPopover) AccessibilityOverflowButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityOverflowButton"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityOverflowButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOverflowButton:"), value)
}

// The accessibility element’s parent in the accessibility hierarchy.
//
// # Discussion
//
// This property must contain a reference to another element in the
// accessibility hierarchy. If you create an [NSView] subclass, you don’t
// typically need to set this value. The system automatically sets the parent
// to the nearest ancestor in the view hierarchy that is also in the
// accessibility hierarchy. If you use an [NSAccessibilityElement] subclass to
// represent an interface element that is not backed by a view, you can either
// set the parent property or you can call the
// [NSAccessibilityElementClass.AccessibilityElementWithRoleFrameLabelParent]
// convenience method, which sets it automatically.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityParent
func (o NSPopover) SetAccessibilityParent(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityParent:"), value)
}

// The placeholder value for the accessibility element.
//
// # Discussion
//
// Use this property for accessibility elements that support placeholder
// values, such as text fields.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityPlaceholderValue
func (o NSPopover) AccessibilityPlaceholderValue() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPlaceholderValue"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSPopover) SetAccessibilityPlaceholderValue(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityPlaceholderValue:"), objc.String(value))
}

// The contents that precede the divider accessibility element.
//
// # Discussion
//
// For example, use this property to set the subview adjacent to a split
// view’s splitter element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityPreviousContents
func (o NSPopover) AccessibilityPreviousContents() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPreviousContents"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilityPreviousContents(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityPreviousContents:"), value)
}

// A Boolean value that determines whether the accessibility element contains
// protected content.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityProtectedContent
func (o NSPopover) AccessibilityProtectedContent() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityProtectedContent"))
	return bool(rv)
}

func (o NSPopover) SetAccessibilityProtectedContent(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityProtectedContent:"), value)
}

// The child accessibility element that represents the window’s proxy icon.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityProxy
func (o NSPopover) AccessibilityProxy() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityProxy"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityProxy(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityProxy:"), value)
}

// A Boolean value that determines whether the accessibility element must have
// content for successful submission of a form.
//
// # Discussion
//
// Returns YES if the element is required to have content; otherwise, NO.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRequired
func (o NSPopover) AccessibilityRequired() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityRequired"))
	return bool(rv)
}

func (o NSPopover) SetAccessibilityRequired(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRequired:"), value)
}

// The type of interface element that the accessibility element represents.
//
// # Discussion
//
// This property contains a nonlocalized string that defines the element’s
// role in the app. For a list of possible roles, see [Roles]. This property
// is set automatically when you adopt one of the accessibility protocols.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRole
func (o NSPopover) AccessibilityRole() NSAccessibilityRole {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRole"))
	return NSAccessibilityRole(foundation.NSStringFromID(rv).String())
}

func (o NSPopover) SetAccessibilityRole(value NSAccessibilityRole) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRole:"), objc.String(string(value)))
}

// A localized, human-intelligible description of the accessibility
// element’s role, such as .
//
// # Discussion
//
// This property is set automatically based on the value of the
// [accessibilityRole] property; however, you can customize the value of this
// property to better describe your element’s role. Keep role descriptions
// short. If possible, use a single word. These descriptions should be noun
// phrases, all lowercase, with no period at the end.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRoleDescription
//
// [accessibilityRole]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRole
func (o NSPopover) AccessibilityRoleDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRoleDescription"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSPopover) SetAccessibilityRoleDescription(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRoleDescription:"), objc.String(value))
}

// The number of rows in the accessibility element’s grid.
//
// # Discussion
//
// Use this property for elements that present a grid of child elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRowCount
func (o NSPopover) AccessibilityRowCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityRowCount"))
	return int(rv)
}

func (o NSPopover) SetAccessibilityRowCount(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowCount:"), value)
}

// The row header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRowHeaderUIElements
func (o NSPopover) AccessibilityRowHeaderUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRowHeaderUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilityRowHeaderUIElements(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowHeaderUIElements:"), value)
}

// The row index range of the cell.
//
// # Discussion
//
// This property contains the row’s starting index and index span in the
// table. Use this property in the elements representing a table’s cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRowIndexRange
func (o NSPopover) AccessibilityRowIndexRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRowIndexRange"))
	return foundation.NSRange(rv)
}

func (o NSPopover) SetAccessibilityRowIndexRange(value foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowIndexRange:"), value)
}

// The row accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRows
func (o NSPopover) AccessibilityRows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRows"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilityRows(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRows:"), value)
}

// The type of markers for the ruler.
//
// # Discussion
//
// Use this property on a ruler element. For a complete list of marker types,
// see [NSAccessibilityRulerMarkerType].
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRulerMarkerType
//
// [NSAccessibilityRulerMarkerType]: https://developer.apple.com/documentation/AppKit/NSAccessibilityRulerMarkerType
func (o NSPopover) AccessibilityRulerMarkerType() NSAccessibilityRulerMarkerType {
	rv := objc.Send[NSAccessibilityRulerMarkerType](o.ID, objc.Sel("accessibilityRulerMarkerType"))
	return NSAccessibilityRulerMarkerType(rv)
}

func (o NSPopover) SetAccessibilityRulerMarkerType(value NSAccessibilityRulerMarkerType) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRulerMarkerType:"), value)
}

// The search button for the search field.
//
// # Discussion
//
// Use this property on a search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySearchButton
func (o NSPopover) AccessibilitySearchButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchButton"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilitySearchButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchButton:"), value)
}

// The search menu for the search field.
//
// # Discussion
//
// Use this property on a search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySearchMenu
func (o NSPopover) AccessibilitySearchMenu() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchMenu"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilitySearchMenu(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchMenu:"), value)
}

// A Boolean value that determines whether the accessibility element is
// currently in a selected state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelected
func (o NSPopover) AccessibilitySelected() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelected"))
	return bool(rv)
}

func (o NSPopover) SetAccessibilitySelected(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelected:"), value)
}

// The currently selected cells for the table.
//
// # Discussion
//
// This property is required for all elements that act like cell-based tables.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedCells
func (o NSPopover) AccessibilitySelectedCells() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedCells"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilitySelectedCells(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedCells:"), value)
}

// The accessibility element’s currently selected children.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedChildren
func (o NSPopover) AccessibilitySelectedChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedChildren"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilitySelectedChildren(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedChildren:"), value)
}

// The currently selected columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedColumns
func (o NSPopover) AccessibilitySelectedColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedColumns"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilitySelectedColumns(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedColumns:"), value)
}

// The currently selected rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedRows
func (o NSPopover) AccessibilitySelectedRows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedRows"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilitySelectedRows(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedRows:"), value)
}

// The currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedText
func (o NSPopover) AccessibilitySelectedText() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedText"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSPopover) SetAccessibilitySelectedText(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedText:"), objc.String(value))
}

// The range of the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedTextRange
func (o NSPopover) AccessibilitySelectedTextRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySelectedTextRange"))
	return foundation.NSRange(rv)
}

func (o NSPopover) SetAccessibilitySelectedTextRange(value foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedTextRange:"), value)
}

// An array of ranges for the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedTextRanges
func (o NSPopover) AccessibilitySelectedTextRanges() []foundation.NSValue {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilitySelectedTextRanges"))
	result := make([]foundation.NSValue, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = foundation.NSValueFromID(id)
	}
	return result
}

func (o NSPopover) SetAccessibilitySelectedTextRanges(value []foundation.NSValue) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedTextRanges:"), objectivec.IObjectSliceToNSArray(value))
}

// The list of elements that the accessibility element is a title for.
//
// # Discussion
//
// Use on a static text label to associate that label with one or more user
// interface elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityServesAsTitleForUIElements
func (o NSPopover) AccessibilityServesAsTitleForUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityServesAsTitleForUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilityServesAsTitleForUIElements(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityServesAsTitleForUIElements:"), value)
}

// The range of characters that the accessibility element displays.
//
// # Discussion
//
// Use this property to manage text that is split across multiple
// elements—for example, an ebook reader that splits the text into multiple
// pages.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySharedCharacterRange
func (o NSPopover) AccessibilitySharedCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySharedCharacterRange"))
	return foundation.NSRange(rv)
}

func (o NSPopover) SetAccessibilitySharedCharacterRange(value foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedCharacterRange:"), value)
}

// An array of elements that shares the keyboard focus with the accessibility
// element.
//
// # Discussion
//
// Use this property to manage elements that share the keyboard focus—for
// example, a search field with completion menu below it.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySharedFocusElements
func (o NSPopover) AccessibilitySharedFocusElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedFocusElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilitySharedFocusElements(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedFocusElements:"), value)
}

// Other elements that share text with the accessibility element.
//
// # Discussion
//
// Use this property to manage text that is split across multiple
// elements—for example, an ebook reader that splits the text into multiple
// pages.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySharedTextUIElements
func (o NSPopover) AccessibilitySharedTextUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedTextUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilitySharedTextUIElements(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedTextUIElements:"), value)
}

// The menu currently displaying for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityShownMenu
func (o NSPopover) AccessibilityShownMenu() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityShownMenu"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityShownMenu(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityShownMenu:"), value)
}

// The accessibility element’s sort direction.
//
// # Discussion
//
// Used by an element with an [button] role and an
// [NSAccessibilitySortButtonRole] subrole. For a list of possible sort
// directions, see [NSAccessibilitySortDirection].
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySortDirection
//
// [NSAccessibilitySortButtonRole]: https://developer.apple.com/documentation/AppKit/NSAccessibilitySortButtonRole
// [NSAccessibilitySortDirection]: https://developer.apple.com/documentation/AppKit/NSAccessibilitySortDirection
// [button]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/button
func (o NSPopover) AccessibilitySortDirection() NSAccessibilitySortDirection {
	rv := objc.Send[NSAccessibilitySortDirection](o.ID, objc.Sel("accessibilitySortDirection"))
	return NSAccessibilitySortDirection(rv)
}

func (o NSPopover) SetAccessibilitySortDirection(value NSAccessibilitySortDirection) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySortDirection:"), value)
}

// An array that contains the views and splitter bar from the split view.
//
// # Discussion
//
// Use this property on a split view element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySplitters
func (o NSPopover) AccessibilitySplitters() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySplitters"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilitySplitters(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySplitters:"), value)
}

// The specialized interface element type that the accessibility element
// represents.
//
// # Discussion
//
// For a list of possible subroles, see [Subroles].
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySubrole
func (o NSPopover) AccessibilitySubrole() NSAccessibilitySubrole {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySubrole"))
	return NSAccessibilitySubrole(foundation.NSStringFromID(rv).String())
}

func (o NSPopover) SetAccessibilitySubrole(value NSAccessibilitySubrole) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySubrole:"), objc.String(string(value)))
}

// The tab accessibility elements for the tab view.
//
// # Discussion
//
// Use this property on a tab view element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityTabs
func (o NSPopover) AccessibilityTabs() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTabs"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilityTabs(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTabs:"), value)
}

// The title of the accessibility element—for example, a button’s visible
// text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityTitle
func (o NSPopover) AccessibilityTitle() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitle"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSPopover) SetAccessibilityTitle(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTitle:"), objc.String(value))
}

// A static text element that represents the accessibility element’s title.
//
// # Discussion
//
// Use this property to associate a static text label with another
// element—for example, to associate a label with its corresponding text
// field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityTitleUIElement
func (o NSPopover) AccessibilityTitleUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitleUIElement"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityTitleUIElement(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTitleUIElement:"), value)
}

// The child accessibility element that represents the window’s toolbar
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityToolbarButton
func (o NSPopover) AccessibilityToolbarButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityToolbarButton"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityToolbarButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityToolbarButton:"), value)
}

// The top-level element that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityTopLevelUIElement
func (o NSPopover) AccessibilityTopLevelUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTopLevelUIElement"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityTopLevelUIElement(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTopLevelUIElement:"), value)
}

// The URL for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityURL
func (o NSPopover) AccessibilityURL() foundation.NSURL {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityURL"))
	return foundation.NSURLFromID(rv)
}

func (o NSPopover) SetAccessibilityURL(value foundation.NSURL) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityURL:"), value)
}

// A human-readable description of the ruler’s units.
//
// # Discussion
//
// Use this property on a ruler element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityUnitDescription
func (o NSPopover) AccessibilityUnitDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSPopover) SetAccessibilityUnitDescription(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUnitDescription:"), objc.String(value))
}

// The units for the ruler.
//
// # Discussion
//
// Use this property on a ruler element. For a complete list of units, see
// [NSAccessibilityUnits].
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityUnits
//
// [NSAccessibilityUnits]: https://developer.apple.com/documentation/AppKit/NSAccessibilityUnits
func (o NSPopover) AccessibilityUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityUnits"))
	return NSAccessibilityUnits(rv)
}

func (o NSPopover) SetAccessibilityUnits(value NSAccessibilityUnits) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUnits:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityUserInputLabels
func (o NSPopover) AccessibilityUserInputLabels() []string {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityUserInputLabels"))
	return objc.ConvertSliceToStrings(rvIDs)
}

func (o NSPopover) SetAccessibilityUserInputLabels(value []string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUserInputLabels:"), objectivec.StringSliceToNSArray(value))
}

// The accessibility element’s value.
//
// # Discussion
//
// The accessibility protocols for roles that support values typically
// redefine this property to take a more specific value type. For example, the
// [staticText] protocol uses [NSString] values, and the [progressIndicator]
// protocol uses [NSNumber] values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [progressIndicator]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/progressIndicator
// [staticText]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/staticText
func (o NSPopover) AccessibilityValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValue"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityValue(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityValue:"), value)
}

// A human-readable description of the accessibility element’s value.
//
// # Discussion
//
// Use this property to provide a more useful description of the accessibility
// element’s raw value. For example, you might set the value to `600`, but
// set the description to `10 minutes`. Always localize this description.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValueDescription
func (o NSPopover) AccessibilityValueDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValueDescription"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSPopover) SetAccessibilityValueDescription(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityValueDescription:"), objc.String(value))
}

// The vertical scroll bar for the scroll view.
//
// # Discussion
//
// Use this property on a scrollable view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVerticalScrollBar
func (o NSPopover) AccessibilityVerticalScrollBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalScrollBar"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityVerticalScrollBar(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalScrollBar:"), value)
}

// A description of the layout area’s vertical units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVerticalUnitDescription
func (o NSPopover) AccessibilityVerticalUnitDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSPopover) SetAccessibilityVerticalUnitDescription(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalUnitDescription:"), objc.String(value))
}

// The units that the layout area uses for vertical values.
//
// # Discussion
//
// For a list of possible values, see [NSAccessibilityUnits].
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVerticalUnits
//
// [NSAccessibilityUnits]: https://developer.apple.com/documentation/AppKit/NSAccessibilityUnits
func (o NSPopover) AccessibilityVerticalUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityVerticalUnits"))
	return NSAccessibilityUnits(rv)
}

func (o NSPopover) SetAccessibilityVerticalUnits(value NSAccessibilityUnits) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalUnits:"), value)
}

// The visible cells for the table.
//
// # Discussion
//
// This property is required for all elements that act like cell-based tables.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleCells
func (o NSPopover) AccessibilityVisibleCells() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleCells"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilityVisibleCells(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleCells:"), value)
}

// The range of visible characters in the document.
//
// # Discussion
//
// Use this property to store the range for entire lines. Characters that are
// horizontally clipped are included in this range.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleCharacterRange
func (o NSPopover) AccessibilityVisibleCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityVisibleCharacterRange"))
	return foundation.NSRange(rv)
}

func (o NSPopover) SetAccessibilityVisibleCharacterRange(value foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleCharacterRange:"), value)
}

// The accessibility element’s visible child accessibility elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleChildren
func (o NSPopover) AccessibilityVisibleChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleChildren"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilityVisibleChildren(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleChildren:"), value)
}

// The visible columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleColumns
func (o NSPopover) AccessibilityVisibleColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleColumns"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilityVisibleColumns(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleColumns:"), value)
}

// The visible rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleRows
func (o NSPopover) AccessibilityVisibleRows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleRows"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilityVisibleRows(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleRows:"), value)
}

// The warning value for the level indicator.
//
// # Discussion
//
// Use this property for elements such as the battery level indicator. This
// property sets a boundary value. If the element’s value exceeds the
// boundary value, the element has reached a warning stage.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityWarningValue
func (o NSPopover) AccessibilityWarningValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWarningValue"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityWarningValue(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWarningValue:"), value)
}

// The window that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityWindow
func (o NSPopover) AccessibilityWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindow"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityWindow(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindow:"), value)
}

// An array that contains all the app’s windows.
//
// # Discussion
//
// Use on the app element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityWindows
func (o NSPopover) AccessibilityWindows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindows"))
	return foundation.NSArrayFromID(rv)
}

func (o NSPopover) SetAccessibilityWindows(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindows:"), value)
}

// The child accessibility element that represents the window’s zoom button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityZoomButton
func (o NSPopover) AccessibilityZoomButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityZoomButton"))
	return objectivec.Object{ID: rv}
}

func (o NSPopover) SetAccessibilityZoomButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityZoomButton:"), value)
}

// Protocol methods for NSAppearanceCustomization
