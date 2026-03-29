// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTouchBarItem] class.
var (
	_NSTouchBarItemClass     NSTouchBarItemClass
	_NSTouchBarItemClassOnce sync.Once
)

func getNSTouchBarItemClass() NSTouchBarItemClass {
	_NSTouchBarItemClassOnce.Do(func() {
		_NSTouchBarItemClass = NSTouchBarItemClass{class: objc.GetClass("NSTouchBarItem")}
	})
	return _NSTouchBarItemClass
}

// GetNSTouchBarItemClass returns the class object for NSTouchBarItem.
func GetNSTouchBarItemClass() NSTouchBarItemClass {
	return getNSTouchBarItemClass()
}

type NSTouchBarItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTouchBarItemClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTouchBarItemClass) Alloc() NSTouchBarItem {
	rv := objc.Send[NSTouchBarItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A UI control shown in the Touch Bar on supported models of MacBook Pro.
//
// # Overview
// 
// An instance of the [NSTouchBarItem] class is called an . It appears to the
// user on the Touch Bar, typically along with other items, within the
// (invisible) bounds of the view for an [NSTouchBar] object, called a .
// 
// You use an item by adding it or its identifier to one or another of a
// bar’s arrays, depending on your app’s architecture and on the user
// customization you want to support. Because of the close interaction between
// bars and items, be sure you have read the overview for the [NSTouchBar]
// class before continuing here to learn about items.
// 
// AppKit provides a rich set of subclasses of [NSTouchBarItem], each of which
// is described in the corresponding class reference document:
// 
// - An [NSCandidateListTouchBarItem] object (a ), along with its delegate,
// provides a list of textual suggestions for the current text view - An
// [NSColorPickerTouchBarItem] object (a ) provides a system-defined color
// picker - An [NSCustomTouchBarItem] object (a ) contains a responder of your
// choice, such as a view, a button, or a scrubber (an instance of the
// [NSScrubber] class) - An [NSGroupTouchBarItem] object (a ) provides a bar
// to contain other items - An [NSPopoverTouchBarItem] object (a ) provides a
// two-state control that, when touched or pressed, expands into its second
// state, showing the contents of a bar it owns - An
// [NSSharingServicePickerTouchBarItem] object (a ), along with its delegate,
// provides a list of objects eligible for sharing - An [NSSliderTouchBarItem]
// object (a ) provides a slider control for choosing a value in a range
// 
// The two most commonly-used item classes are [NSCustomTouchBarItem] and
// [NSPopoverTouchBarItem].
// 
// Refer to the following sample code projects which demonstrate how to use
// [NSTouchBarItem] and related classes:
// 
// - [Creating and Customizing the Touch Bar] - [Integrating a Toolbar and
// Touch Bar into Your App]
// 
// # Custom items
// 
// You typically use a (an instance of the [NSCustomTouchBarItem] class) to
// hold a view. For example, to place a button in the Touch Bar, proceed as
// follows:
// 
// - Use an [NSButton] convenience initializer such as
// [ButtonWithTitleImageTargetAction] to create and configure the button. -
// Set the [View] property for a custom item to point to the new button.
// 
// # Popover items
// 
// A (an instance of the [NSPopoverTouchBarItem] class) — the second
// commonly-used type — lets you provide a new bar (an [NSTouchBar] object)
// when a user taps, or presses-and-holds, on the collapsed representation of
// the popover item.
// 
// In its expanded state, a popover appears as an overlay above other items in
// the Touch Bar.
// 
// To show a bar when a user taps a popover item, specify a bar in the
// item’s [NSTouchBarItem.PopoverTouchBar] property. Enable press-and-hold by specifying a
// bar in the [NSTouchBarItem.PressAndHoldTouchBar] property. The press-and-hold feature is
// suitable only for a simple popover, such as one that contains a single
// segmented control (an instance of the [NSSegmentedControl] class) or slider
// (an instance of the [NSSliderTouchBarItem] class).
// 
// The system automatically shows a chevron in the popover item under the
// following conditions: You specify the same [NSTouchBar] object for both
// [NSTouchBarItem.PressAndHoldTouchBar] and [NSTouchBarItem.PopoverTouchBar] properties, you use the
// default view for the popover item’s [NSTouchBarItem.CollapsedRepresentation] property.
// 
// If you provide a popover item that contains a scrubber (an [NSScrubber]
// instance), you’ll likely want to dismiss both the scrubber and the
// popover after the user makes their selection in the scrubber. A good
// approach to achieve this user interaction is to subclass
// [NSPopoverTouchBarItem], employing your instance of the subclass as the
// scrubber’s delegate. You can then configure the delegate object, within
// its [DidFinishInteractingWithScrubber] method, to call the popover’s
// [DismissPopover] method.
// 
// If you place a segmented control in a bar for a popover item, take care to
// use [NSSegmentedControl.SwitchTracking.momentary] option of the
// [NSSegmentedControl.SwitchTracking] enumeration because doing so interferes
// with the user’s operation of the control.
// 
// # Other common item types
// 
// To provide a , always use the [NSSliderTouchBarItem] class, which employs a
// standard slider but is optimized for user interaction with the Touch Bar.
// (That is, don’t instead add an [NSSlider] object directly to a custom
// item.)
// 
// A (an instance of the [NSGroupTouchBarItem] class) is a container that
// provides a bar, in its [NSTouchBarItem.GroupTouchBar] property, with its own array of
// items. You can enable customization for the items in a group’s contained
// bar, in the same way you would for items directly within a top-level bar.
// Using a group item lets you provide different user customization rules for
// different parts of the Touch Bar. Using a group item also lets you enable
// centering of the group within the Touch Bar.
// 
// A lets you add custom spacing between items in a bar. Specify a spacing
// item for a bar by assigning the [fixedSpaceSmall], [fixedSpaceLarge], or
// [flexibleSpace] identifier to an item, and adding that item to the bar’s
// items array. The system automatically instantiates and configures spacing
// items based on the identifiers you specify.
// 
// # Configuration
// 
// You must configure each item with a unique identifier, and can optionally
// assign a visibility priority or tag it as a principal item.
// 
// You must provide a unique identifier for each item in the bar, apart from
// spacing items. Specify an identifier, of type [NSTouchBarItemIdentifier]
// (called an ), for each item when you initialize it. The item identifier
// serves as a persistable weak reference to the item. The system uses item
// identifiers to populate bars and to track and record changes for user
// customization.
// 
// If the system is showing a bar in the Touch Bar, but horizontal space is
// constrained and the bar defines more items than will fit, the system hides
// some of the items. You influence this hide/show behavior by setting a value
// for the [NSTouchBarItem.VisibilityPriority] property of each item.
// 
// Lower-visibility-priority items get hidden by the system, as needed, before
// higher-visibility-priority items do.
// 
// To set visibility priority, use the constants in the
// [NSTouchBarItemPriority] enumeration, or assign an integer value. The value
// `0` indicates [normal] visibility priority. Visibility priority increases
// with increasing numerical value. The [low] constant provides a value of
// `-1000`; the [high] constant, a value `+1000`. You can use integers outside
// of this range if you need to.
// 
// The system hides or shows groups of identical-priority items (defined
// within a single bar) together. The one exception to this rule is for items
// whose visibility priority is [normal]; these items get hidden one-by-one,
// with the normal-priority item farthest to the right getting hidden first.
// If horizontal space later increases in the Touch Bar, and hidden,
// normal-priority items become eligible for display, the system first shows
// the most recently-hidden of those items.
// 
// Within a bar, you can optionally specify an item as having special
// significance by employing the [NSTouchBarItem.PrincipalItemIdentifier] property. The
// system attempts to center a principal item within the Touch Bar. If you
// want a group of items to appear centered in the Touch Bar, designate the
// group item (of type [NSTouchBarItem]) as the principal item.
// 
// If more than one bar in the responder chain is eligible to be visible in
// the Touch Bar, and more than one of those has a principal item, the system
// determines which one to center in the Touch Bar.
// 
// # Fonts, images, and colors
// 
// When using a button in a custom item, don’t attempt to set the button
// title’s font. In the Touch Bar, the system specifies fonts for standard
// controls.
// 
// If you need to specify a font, such as for custom drawing, use the
// [SystemFontOfSize] class method (or related methods) of the [NSFont] class.
// Use a font size of `0` to automatically obtain appropriate sizing for the
// Touch Bar.
// 
// If you use an image in a button or other control in the Touch Bar, take
// care to employ a template image. Template images in the Touch Bar respond
// automatically to system white-point changes, and automatically react to
// user interactions. The overview in this document lists the built-in Touch
// Bar template images.
// 
// To use your own image assets, use Retina-resolution images, designated as
// `@2x` in your asset catalog and with a maximum height of 30 points
// (corresponding to 60 pixels).
// 
// To set colors on objects within an [NSTouchBarItem] object, use AppKit
// named colors and use a bezel color property (available starting in macOS
// 10.12.1). Named colors appear correctly in the Touch Bar, support
// appearance vibrancy, and respond to system white-point changes. In a button
// or a segmented control, employ the bezel color property to ensure
// appropriate appearance in the Touch Bar.
// 
// To set the background color on a button within a custom item, use code like
// this:
// 
// To set color on text and glyphs in the Touch Bar, use the following colors
// from the [NSColor] class:
// 
// - [labelColor] - [secondaryLabelColor] - [tertiaryLabelColor] -
// [quaternaryLabelColor]
// 
// The system automatically changes the relative brightness and the
// white-point of these colors, depending on the ambient light, and depending
// on other factors such as keyboard backlight level. Always use these colors,
// or colors that dynamically derive from these colors, for control
// backgrounds, text, icons, and glyphs in the Touch Bar.
// 
// # Handling touch events
// 
// The easiest way to handle touch events in an item is to use AppKit
// controls, such as by adding a button, a segmented control, or a scrubber to
// the item. Standard AppKit controls convey touch events to your specified
// targets automatically, so use standard controls whenever possible in your
// app.
// 
// If standard controls are insufficient, you can create composite views with
// a combination of standard controls, custom views, and gesture recognizers
// that you manually add to those custom views.
// 
// If you require the lowest-level of control for touch event processing, you
// can use the [NSTouch] class directly. You might go this route, for example,
// to provide good user feedback in the case of a control placed within a
// scroll view.Direct use of touch methods allows fine-grained control over
// interaction. You can, for example, highlight a control immediately upon a
// user touching it, and then remove the highlight if the user then, without
// lifting the finger, performs a scroll gesture.
// 
// If using the [NSTouch] class directly, be sure to implement the
// [TouchesCancelledWithEvent] responder method, because users can perform
// touch interactions that result in canceled touches.
//
// [Creating and Customizing the Touch Bar]: https://developer.apple.com/documentation/AppKit/creating-and-customizing-the-touch-bar
// [Integrating a Toolbar and Touch Bar into Your App]: https://developer.apple.com/documentation/AppKit/integrating-a-toolbar-and-touch-bar-into-your-app
// [NSSegmentedControl.SwitchTracking.momentary]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/SwitchTracking/momentary
// [NSSegmentedControl.SwitchTracking]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/SwitchTracking
// [fixedSpaceLarge]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/Identifier-swift.struct/fixedSpaceLarge
// [fixedSpaceSmall]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/Identifier-swift.struct/fixedSpaceSmall
// [flexibleSpace]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/Identifier-swift.struct/flexibleSpace
// [high]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/Priority/high
// [labelColor]: https://developer.apple.com/documentation/AppKit/NSColor/labelColor
// [low]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/Priority/low
// [normal]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/Priority/normal
// [quaternaryLabelColor]: https://developer.apple.com/documentation/AppKit/NSColor/quaternaryLabelColor
// [secondaryLabelColor]: https://developer.apple.com/documentation/AppKit/NSColor/secondaryLabelColor
// [tertiaryLabelColor]: https://developer.apple.com/documentation/AppKit/NSColor/tertiaryLabelColor
//
// # Creating a bar item
//
//   - [NSTouchBarItem.InitWithIdentifier]: Creates a new item with the specified identifier.
//   - [NSTouchBarItem.InitWithCoder]: Initializes and returns a new item from a storyboard or nib file.
//
// # Identifying a bar item
//
//   - [NSTouchBarItem.Identifier]: The identifier for this item.
//
// # Managing item visibility
//
//   - [NSTouchBarItem.VisibilityPriority]: Determines which items are shown in a bar when space is limited.
//   - [NSTouchBarItem.SetVisibilityPriority]
//   - [NSTouchBarItem.Visible]: A Boolean value that reflects whether or not the item is visible.
//
// # Configuring bar customization
//
//   - [NSTouchBarItem.CustomizationLabel]: The user-visible string identifying this item during bar customization.
//
// # Subclassing bar items
//
//   - [NSTouchBarItem.ViewController]: The view controller associated with this item.
//   - [NSTouchBarItem.View]: The view associated with this item.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem
type NSTouchBarItem struct {
	objectivec.Object
}

// NSTouchBarItemFromID constructs a [NSTouchBarItem] from an objc.ID.
//
// A UI control shown in the Touch Bar on supported models of MacBook Pro.
func NSTouchBarItemFromID(id objc.ID) NSTouchBarItem {
	return NSTouchBarItem{objectivec.Object{ID: id}}
}
// NOTE: NSTouchBarItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTouchBarItem] class.
//
// # Creating a bar item
//
//   - [INSTouchBarItem.InitWithIdentifier]: Creates a new item with the specified identifier.
//   - [INSTouchBarItem.InitWithCoder]: Initializes and returns a new item from a storyboard or nib file.
//
// # Identifying a bar item
//
//   - [INSTouchBarItem.Identifier]: The identifier for this item.
//
// # Managing item visibility
//
//   - [INSTouchBarItem.VisibilityPriority]: Determines which items are shown in a bar when space is limited.
//   - [INSTouchBarItem.SetVisibilityPriority]
//   - [INSTouchBarItem.Visible]: A Boolean value that reflects whether or not the item is visible.
//
// # Configuring bar customization
//
//   - [INSTouchBarItem.CustomizationLabel]: The user-visible string identifying this item during bar customization.
//
// # Subclassing bar items
//
//   - [INSTouchBarItem.ViewController]: The view controller associated with this item.
//   - [INSTouchBarItem.View]: The view associated with this item.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem
type INSTouchBarItem interface {
	objectivec.IObject

	// Topic: Creating a bar item

	// Creates a new item with the specified identifier.
	InitWithIdentifier(identifier NSTouchBarItemIdentifier) NSTouchBarItem
	// Initializes and returns a new item from a storyboard or nib file.
	InitWithCoder(coder foundation.INSCoder) NSTouchBarItem

	// Topic: Identifying a bar item

	// The identifier for this item.
	Identifier() NSTouchBarItemIdentifier

	// Topic: Managing item visibility

	// Determines which items are shown in a bar when space is limited.
	VisibilityPriority() NSTouchBarItemPriority
	SetVisibilityPriority(value NSTouchBarItemPriority)
	// A Boolean value that reflects whether or not the item is visible.
	Visible() bool

	// Topic: Configuring bar customization

	// The user-visible string identifying this item during bar customization.
	CustomizationLabel() string

	// Topic: Subclassing bar items

	// The view controller associated with this item.
	ViewController() INSViewController
	// The view associated with this item.
	View() INSView

	// The view displayed when this item is displayed in its parent bar.
	CollapsedRepresentation() INSView
	SetCollapsedRepresentation(value INSView)
	// A bar that holds this group’s items.
	GroupTouchBar() INSTouchBar
	SetGroupTouchBar(value INSTouchBar)
	// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
	IsContinuous() bool
	SetIsContinuous(value bool)
	// The bar displayed when this item is “popped.”
	PopoverTouchBar() INSTouchBar
	SetPopoverTouchBar(value INSTouchBar)
	// The bar that is displayed when a user press-and-holds on the popover item.
	PressAndHoldTouchBar() INSTouchBar
	SetPressAndHoldTouchBar(value INSTouchBar)
	// The identifier of an item you want the system to center in the Touch Bar.
	PrincipalItemIdentifier() NSTouchBarItemIdentifier
	SetPrincipalItemIdentifier(value NSTouchBarItemIdentifier)
	// The type of tracking behavior the control exhibits.
	TrackingMode() NSSegmentSwitchTracking
	SetTrackingMode(value NSSegmentSwitchTracking)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (t NSTouchBarItem) Init() NSTouchBarItem {
	rv := objc.Send[NSTouchBarItem](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTouchBarItem) Autorelease() NSTouchBarItem {
	rv := objc.Send[NSTouchBarItem](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTouchBarItem creates a new NSTouchBarItem instance.
func NewNSTouchBarItem() NSTouchBarItem {
	class := getNSTouchBarItemClass()
	rv := objc.Send[NSTouchBarItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes and returns a new item from a storyboard or nib file.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(coder:)
func NewTouchBarItemWithCoder(coder foundation.INSCoder) NSTouchBarItem {
	instance := getNSTouchBarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTouchBarItemFromID(rv)
}

// Creates a new item with the specified identifier.
//
// # Discussion
// 
// The designated initializer. The identifier must be globally unique for
// every item, except for space items.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(identifier:)
func NewTouchBarItemWithIdentifier(identifier NSTouchBarItemIdentifier) NSTouchBarItem {
	instance := getNSTouchBarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifier:"), objc.String(string(identifier)))
	return NSTouchBarItemFromID(rv)
}

// Creates a new item with the specified identifier.
//
// # Discussion
// 
// The designated initializer. The identifier must be globally unique for
// every item, except for space items.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(identifier:)
func (t NSTouchBarItem) InitWithIdentifier(identifier NSTouchBarItemIdentifier) NSTouchBarItem {
	rv := objc.Send[NSTouchBarItem](t.ID, objc.Sel("initWithIdentifier:"), objc.String(string(identifier)))
	return rv
}
// Initializes and returns a new item from a storyboard or nib file.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(coder:)
func (t NSTouchBarItem) InitWithCoder(coder foundation.INSCoder) NSTouchBarItem {
	rv := objc.Send[NSTouchBarItem](t.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (t NSTouchBarItem) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The identifier for this item.
//
// # Discussion
// 
// This read-only property returns the value the item was initialized with.
// 
// For all items other than spaces, this value must be globally unique.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/identifier-swift.property
func (t NSTouchBarItem) Identifier() NSTouchBarItemIdentifier {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("identifier"))
	return NSTouchBarItemIdentifier(foundation.NSStringFromID(rv).String())
}
// Determines which items are shown in a bar when space is limited.
//
// # Discussion
// 
// The bar hides items of lower priority when there is not enough space to
// show all items. Use this property to specify the relative priority of the
// items in your bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/visibilityPriority
func (t NSTouchBarItem) VisibilityPriority() NSTouchBarItemPriority {
	rv := objc.Send[NSTouchBarItemPriority](t.ID, objc.Sel("visibilityPriority"))
	return NSTouchBarItemPriority(rv)
}
func (t NSTouchBarItem) SetVisibilityPriority(value NSTouchBarItemPriority) {
	objc.Send[struct{}](t.ID, objc.Sel("setVisibilityPriority:"), value)
}
// A Boolean value that reflects whether or not the item is visible.
//
// # Discussion
// 
// When [true], this item is shown in a currently visible bar. This property
// is always [false] for spaces, proxy items, and groups.
// 
// This property is key-value observable.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/isVisible
func (t NSTouchBarItem) Visible() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isVisible"))
	return rv
}
// The user-visible string identifying this item during bar customization.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/customizationLabel
func (t NSTouchBarItem) CustomizationLabel() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("customizationLabel"))
	return foundation.NSStringFromID(rv).String()
}
// The view controller associated with this item.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/viewController
func (t NSTouchBarItem) ViewController() INSViewController {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("viewController"))
	return NSViewControllerFromID(objc.ID(rv))
}
// The view associated with this item.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/view
func (t NSTouchBarItem) View() INSView {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("view"))
	return NSViewFromID(objc.ID(rv))
}
// The view displayed when this item is displayed in its parent bar.
//
// See: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/collapsedrepresentation
func (t NSTouchBarItem) CollapsedRepresentation() INSView {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("collapsedRepresentation"))
	return NSViewFromID(objc.ID(rv))
}
func (t NSTouchBarItem) SetCollapsedRepresentation(value INSView) {
	objc.Send[struct{}](t.ID, objc.Sel("setCollapsedRepresentation:"), value)
}
// A bar that holds this group’s items.
//
// See: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/grouptouchbar
func (t NSTouchBarItem) GroupTouchBar() INSTouchBar {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("groupTouchBar"))
	return NSTouchBarFromID(objc.ID(rv))
}
func (t NSTouchBarItem) SetGroupTouchBar(value INSTouchBar) {
	objc.Send[struct{}](t.ID, objc.Sel("setGroupTouchBar:"), value)
}
// A Boolean value indicating whether the receiver’s cell sends its action
// message continuously to its target during mouse tracking.
//
// See: https://developer.apple.com/documentation/appkit/nscontrol/iscontinuous
func (t NSTouchBarItem) IsContinuous() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("continuous"))
	return rv
}
func (t NSTouchBarItem) SetIsContinuous(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setContinuous:"), value)
}
// The bar displayed when this item is “popped.”
//
// See: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/popovertouchbar
func (t NSTouchBarItem) PopoverTouchBar() INSTouchBar {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("popoverTouchBar"))
	return NSTouchBarFromID(objc.ID(rv))
}
func (t NSTouchBarItem) SetPopoverTouchBar(value INSTouchBar) {
	objc.Send[struct{}](t.ID, objc.Sel("setPopoverTouchBar:"), value)
}
// The bar that is displayed when a user press-and-holds on the popover item.
//
// See: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/pressandholdtouchbar
func (t NSTouchBarItem) PressAndHoldTouchBar() INSTouchBar {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("pressAndHoldTouchBar"))
	return NSTouchBarFromID(objc.ID(rv))
}
func (t NSTouchBarItem) SetPressAndHoldTouchBar(value INSTouchBar) {
	objc.Send[struct{}](t.ID, objc.Sel("setPressAndHoldTouchBar:"), value)
}
// The identifier of an item you want the system to center in the Touch Bar.
//
// See: https://developer.apple.com/documentation/appkit/nstouchbar/principalitemidentifier
func (t NSTouchBarItem) PrincipalItemIdentifier() NSTouchBarItemIdentifier {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("principalItemIdentifier"))
	return NSTouchBarItemIdentifier(foundation.NSStringFromID(rv).String())
}
func (t NSTouchBarItem) SetPrincipalItemIdentifier(value NSTouchBarItemIdentifier) {
	objc.Send[struct{}](t.ID, objc.Sel("setPrincipalItemIdentifier:"), objc.String(string(value)))
}
// The type of tracking behavior the control exhibits.
//
// See: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/trackingmode
func (t NSTouchBarItem) TrackingMode() NSSegmentSwitchTracking {
	rv := objc.Send[NSSegmentSwitchTracking](t.ID, objc.Sel("trackingMode"))
	return NSSegmentSwitchTracking(rv)
}
func (t NSTouchBarItem) SetTrackingMode(value NSSegmentSwitchTracking) {
	objc.Send[struct{}](t.ID, objc.Sel("setTrackingMode:"), value)
}

// A template image for showing additional detail for an item.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbaradddetailtemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarAddDetailTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarAddDetailTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for creating a new item.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbaraddtemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarAddTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarAddTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for setting or showing an alarm.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbaralarmtemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarAlarmTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarAlarmTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for muting audio input or denoting that audio input is
// muted.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbaraudioinputmutetemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarAudioInputMuteTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarAudioInputMuteTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for denoting audio input.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbaraudioinputtemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarAudioInputTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarAudioInputTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for muting audio output or for denoting that audio output
// is muted.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbaraudiooutputmutetemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarAudioOutputMuteTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarAudioOutputMuteTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for setting the audio output volume to a high level, or
// for denoting that the audio is at the peak volume.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbaraudiooutputvolumehightemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarAudioOutputVolumeHighTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarAudioOutputVolumeHighTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for setting the audio output volume to a low level, or for
// denoting that it is set to a low level.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbaraudiooutputvolumelowtemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarAudioOutputVolumeLowTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarAudioOutputVolumeLowTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for setting the audio output volume to a medium level, or
// for denoting that it is set to a medium level.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbaraudiooutputvolumemediumtemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarAudioOutputVolumeMediumTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarAudioOutputVolumeMediumTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for setting the audio output volume to silent, or for
// denoting that it is set to silent.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbaraudiooutputvolumeofftemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarAudioOutputVolumeOffTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarAudioOutputVolumeOffTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for showing app-specific bookmarks.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarbookmarkstemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarBookmarksTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarBookmarksTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for showing a color picker so the user can select a fill
// color.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarcolorpickerfillname
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarColorPickerFillName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarColorPickerFill"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for showing a color picker so the user can select a text
// color.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarcolorpickerfontname
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarColorPickerFontName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarColorPickerFont"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for showing a color picker so the user can select a stroke
// color.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarcolorpickerstrokename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarColorPickerStrokeName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarColorPickerStroke"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for initiating or denoting audio communication.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarcommunicationaudiotemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarCommunicationAudioTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarCommunicationAudioTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for initiating or denoting video communication.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarcommunicationvideotemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarCommunicationVideoTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarCommunicationVideoTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for opening a new document or view in edit mode.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarcomposetemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarComposeTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarComposeTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for deleting the current or selected item.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbardeletetemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarDeleteTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarDeleteTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for downloading an item.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbardownloadtemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarDownloadTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarDownloadTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for entering full screen mode.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarenterfullscreentemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarEnterFullScreenTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarEnterFullScreenTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for exiting full screen mode.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarexitfullscreentemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarExitFullScreenTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarExitFullScreenTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for moving forward through media playback or slides.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarfastforwardtemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarFastForwardTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarFastForwardTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for opening or representing a folder.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarfoldertemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarFolderTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarFolderTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for copying an item to a destination.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarfoldercopytotemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarFolderCopyToTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarFolderCopyToTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for moving an item to a destination.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarfoldermovetotemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarFolderMoveToTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarFolderMoveToTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for showing information about an item.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbargetinfotemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarGetInfoTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarGetInfoTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for returning to the previous screen or location.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbargobacktemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarGoBackTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarGoBackTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for moving to the next item in a list.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbargodowntemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarGoDownTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarGoDownTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for moving to the next screen or location.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbargoforwardtemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarGoForwardTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarGoForwardTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for moving to the previous item in a list.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbargouptemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarGoUpTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarGoUpTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for showing history information, such as recent downloads.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarhistorytemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarHistoryTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarHistoryTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for showing items in an icon view.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbariconviewtemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarIconViewTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarIconViewTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for showing items in a list view.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarlistviewtemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarListViewTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarListViewTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for creating an email message.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarmailtemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarMailTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarMailTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for creating a new folder.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarnewfoldertemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarNewFolderTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarNewFolderTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for creating a new message, or for denoting the use of
// messaging.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarnewmessagetemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarNewMessageTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarNewMessageTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for opening an item in the user’s browser.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbaropeninbrowsertemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarOpenInBrowserTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarOpenInBrowserTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for pausing media playback or slides.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarpausetemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarPauseTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarPauseTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for starting or resuming playback of media or slides.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarplaytemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarPlayTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarPlayTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for toggling between playing and pausing media or slides.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarplaypausetemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarPlayPauseTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarPlayPauseTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for denoting the current playback position within a
// timeline track.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarplayheadtemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarPlayheadTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarPlayheadTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for opening an item in Quick Look.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarquicklooktemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarQuickLookTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarQuickLookTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for starting recording.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarrecordstarttemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarRecordStartTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarRecordStartTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for stopping recording or stopping playback of media or
// slides.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarrecordstoptemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarRecordStopTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarRecordStopTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for refreshing displayed data.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarrefreshtemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarRefreshTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarRefreshTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for moving backwards through media or slides.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarrewindtemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarRewindTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarRewindTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for rotating an item counterclockwise.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarrotatelefttemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarRotateLeftTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarRotateLeftTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for rotating an item clockwise.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarrotaterighttemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarRotateRightTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarRotateRightTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for showing a search field or for initiating a search.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarsearchtemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarSearchTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarSearchTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for sharing content with others directly or via social
// media.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarsharetemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarShareTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarShareTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for showing a sidebar in the current view.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarsidebartemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarSidebarTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarSidebarTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for skipping to the previous chapter or location during
// media playback.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarskipbacktemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarSkipBackTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarSkipBackTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for skipping to the start of media playback.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarskiptostarttemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarSkipToStartTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarSkipToStartTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for skipping back 30 seconds during media playback.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarskipback30secondstemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarSkipBack30SecondsTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarSkipBack30SecondsTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for skipping back 15 seconds during media playback.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarskipback15secondstemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarSkipBack15SecondsTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarSkipBack15SecondsTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for skipping ahead 15 seconds during media playback.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarskipahead15secondstemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarSkipAhead15SecondsTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarSkipAhead15SecondsTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for skipping ahead 30 seconds during media playback.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarskipahead30secondstemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarSkipAhead30SecondsTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarSkipAhead30SecondsTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for skipping to the end of media playback.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarskiptoendtemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarSkipToEndTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarSkipToEndTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for skipping to the next chapter or location during media
// playback.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarskipaheadtemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarSkipAheadTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarSkipAheadTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for starting a slideshow.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarslideshowtemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarSlideshowTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarSlideshowTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for applying a tag to an item.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbartagicontemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarTagIconTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarTagIconTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for inserting a text box.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbartextboxtemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarTextBoxTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarTextBoxTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for inserting a list or converting text to list form.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbartextlisttemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarTextListTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarTextListTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for making selected text bold.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbartextboldtemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarTextBoldTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarTextBoldTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for italicizing the selected text.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbartextitalictemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarTextItalicTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarTextItalicTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for underlining text.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbartextunderlinetemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarTextUnderlineTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarTextUnderlineTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for striking through text.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbartextstrikethroughtemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarTextStrikethroughTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarTextStrikethroughTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for fully justifying text.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbartextjustifiedaligntemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarTextJustifiedAlignTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarTextJustifiedAlignTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for aligning text to the left.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbartextleftaligntemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarTextLeftAlignTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarTextLeftAlignTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for centering text.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbartextcenteraligntemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarTextCenterAlignTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarTextCenterAlignTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for aligning text to the right.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbartextrightaligntemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarTextRightAlignTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarTextRightAlignTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for showing or representing user information.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarusertemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarUserTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarUserTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for creating a new user account.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbaruseraddtemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarUserAddTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarUserAddTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for showing or representing a group of users.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarusergrouptemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarUserGroupTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarUserGroupTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for increasing the audio output volume.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarvolumeuptemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarVolumeUpTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarVolumeUpTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// The primary color to use for text labels.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/labelcolor
func (_NSTouchBarItemClass NSTouchBarItemClass) LabelColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("labelColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSTouchBarItemClass NSTouchBarItemClass) SetLabelColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSTouchBarItemClass.class), objc.Sel("setLabelColor:"), value)
}
// The quaternary color to use for text labels and separators.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/quaternarylabelcolor
func (_NSTouchBarItemClass NSTouchBarItemClass) QuaternaryLabelColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("quaternaryLabelColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSTouchBarItemClass NSTouchBarItemClass) SetQuaternaryLabelColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSTouchBarItemClass.class), objc.Sel("setQuaternaryLabelColor:"), value)
}
// The secondary color to use for text labels.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/secondarylabelcolor
func (_NSTouchBarItemClass NSTouchBarItemClass) SecondaryLabelColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("secondaryLabelColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSTouchBarItemClass NSTouchBarItemClass) SetSecondaryLabelColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSTouchBarItemClass.class), objc.Sel("setSecondaryLabelColor:"), value)
}
// The tertiary color to use for text labels.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/tertiarylabelcolor
func (_NSTouchBarItemClass NSTouchBarItemClass) TertiaryLabelColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("tertiaryLabelColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSTouchBarItemClass NSTouchBarItemClass) SetTertiaryLabelColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSTouchBarItemClass.class), objc.Sel("setTertiaryLabelColor:"), value)
}
// A template image for reducing the audio output volume.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarvolumedowntemplatename
func (_NSTouchBarItemClass NSTouchBarItemClass) TouchBarVolumeDownTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarItemClass.class), objc.Sel("NSImageNameTouchBarVolumeDownTemplate"))
	return foundation.NSStringFromID(rv).String()
}

