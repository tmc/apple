// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
// An instance of the [NSTouchBarItem] class is called an item. It appears to
// the user on the Touch Bar, typically along with other items, within the
// (invisible) bounds of the view for an [NSTouchBar] object, called a bar.
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
// - An [NSCandidateListTouchBarItem] object (a candidate-list item), along
// with its delegate, provides a list of textual suggestions for the current
// text view - An [NSColorPickerTouchBarItem] object (a color picker item)
// provides a system-defined color picker - An [NSCustomTouchBarItem] object
// (a custom item) contains a responder of your choice, such as a view, a
// button, or a scrubber (an instance of the [NSScrubber] class) - An
// [NSGroupTouchBarItem] object (a group item) provides a bar to contain other
// items - An [NSPopoverTouchBarItem] object (a popover item) provides a
// two-state control that, when touched or pressed, expands into its second
// state, showing the contents of a bar it owns - An
// [NSSharingServicePickerTouchBarItem] object (a sharing service picker
// item), along with its delegate, provides a list of objects eligible for
// sharing - An [NSSliderTouchBarItem] object (a slider item) provides a
// slider control for choosing a value in a range
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
// You typically use a custom item (an instance of the [NSCustomTouchBarItem]
// class) to hold a view. For example, to place a button in the Touch Bar,
// proceed as follows:
//
// - Use an [NSButton] convenience initializer such as
// [NSStatusBarButtonClass.ButtonWithTitleImageTargetAction] to create and
// configure the button. - Set the [NSTouchBarItem.View] property for a custom
// item to point to the new button.
//
// # Popover items
//
// A popover item (an instance of the [NSPopoverTouchBarItem] class) — the
// second commonly-used type — lets you provide a new bar (an [NSTouchBar]
// object) when a user taps, or presses-and-holds, on the collapsed
// representation of the popover item.
//
// In its expanded state, a popover appears as an overlay above other items in
// the Touch Bar.
//
// To show a bar when a user taps a popover item, specify a bar in the
// item’s [NSPopoverTouchBarItem.PopoverTouchBar] property. Enable
// press-and-hold by specifying a bar in the
// [NSPopoverTouchBarItem.PressAndHoldTouchBar] property. The press-and-hold
// feature is suitable only for a simple popover, such as one that contains a
// single segmented control (an instance of the [NSSegmentedControl] class) or
// slider (an instance of the [NSSliderTouchBarItem] class).
//
// The system automatically shows a chevron in the popover item under the
// following conditions: You specify the same [NSTouchBar] object for both
// [NSPopoverTouchBarItem.PressAndHoldTouchBar] and
// [NSPopoverTouchBarItem.PopoverTouchBar] properties, and you use the default
// view for the popover item’s
// [NSPopoverTouchBarItem.CollapsedRepresentation] property.
//
// If you provide a popover item that contains a scrubber (an [NSScrubber]
// instance), you’ll likely want to dismiss both the scrubber and the
// popover after the user makes their selection in the scrubber. A good
// approach to achieve this user interaction is to subclass
// [NSPopoverTouchBarItem], employing your instance of the subclass as the
// scrubber’s delegate. You can then configure the delegate object, within
// its [DidFinishInteractingWithScrubber] method, to call the popover’s
// [NSPopoverTouchBarItem.DismissPopover] method.
//
// If you place a segmented control in a bar for a popover item, take care not
// to use [NSSegmentSwitchTrackingMomentary] option of the
// [NSSegmentedControl.SwitchTracking] enumeration because doing so interferes
// with the user’s operation of the control.
//
// # Other common item types
//
// To provide a slider item, always use the [NSSliderTouchBarItem] class,
// which employs a standard slider but is optimized for user interaction with
// the Touch Bar. (That is, don’t instead add an [NSSlider] object directly
// to a custom item.)
//
// A group item (an instance of the [NSGroupTouchBarItem] class) is a
// container that provides a bar, in its [NSGroupTouchBarItem.GroupTouchBar]
// property, with its own array of items. You can enable customization for the
// items in a group’s contained bar, in the same way you would for items
// directly within a top-level bar. Using a group item lets you provide
// different user customization rules for different parts of the Touch Bar.
// Using a group item also lets you enable centering of the group within the
// Touch Bar.
//
// A spacing item lets you add custom spacing between items in a bar. Specify
// a spacing item for a bar by assigning the [fixedSpaceSmall],
// [fixedSpaceLarge], or [flexibleSpace] identifier to an item, and adding
// that item to the bar’s items array. The system automatically instantiates
// and configures spacing items based on the identifiers you specify.
//
// # Configuration
//
// You must configure each item with a unique identifier, and can optionally
// assign a visibility priority or tag it as a principal item.
//
// NSTouchBarItem identification. You must provide a unique identifier for
// each item in the bar, apart from spacing items. Specify an identifier, of
// type [NSTouchBarItemIdentifier] (called an item identifier), for each item
// when you initialize it. The item identifier serves as a persistable weak
// reference to the item. The system uses item identifiers to populate bars
// and to track and record changes for user customization.
//
// NSTouchBarItem priority for visibility. If the system is showing a bar in
// the Touch Bar, but horizontal space is constrained and the bar defines more
// items than will fit, the system hides some of the items. You influence this
// hide/show behavior by setting a value for the
// [NSTouchBarItem.VisibilityPriority] property of each item.
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
// Principal Items. Within a bar, you can optionally specify an item as having
// special significance by employing the [NSTouchBar.PrincipalItemIdentifier]
// property. The system attempts to center a principal item within the Touch
// Bar. If you want a group of items to appear centered in the Touch Bar,
// designate the group item (of type [NSTouchBarItem]) as the principal item.
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
// [NSFontClass.SystemFontOfSize] class method (or related methods) of the
// [NSFont] class. Use a font size of `0` to automatically obtain appropriate
// sizing for the Touch Bar.
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
// - [NSColorClass.LabelColor] - [NSColorClass.SecondaryLabelColor] -
// [NSColorClass.TertiaryLabelColor] - [NSColorClass.QuaternaryLabelColor]
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
// [NSGestureRecognizer.TouchesCancelledWithEvent] responder method, because
// users can perform touch interactions that result in canceled touches.
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
//   - [NSTouchBarItem.IsVisible]: A Boolean value that reflects whether or not the item is visible.
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
//
// [Creating and Customizing the Touch Bar]: https://developer.apple.com/documentation/AppKit/creating-and-customizing-the-touch-bar
// [Integrating a Toolbar and Touch Bar into Your App]: https://developer.apple.com/documentation/AppKit/integrating-a-toolbar-and-touch-bar-into-your-app
// [NSSegmentedControl.SwitchTracking]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/SwitchTracking
// [fixedSpaceLarge]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/Identifier-swift.struct/fixedSpaceLarge
// [fixedSpaceSmall]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/Identifier-swift.struct/fixedSpaceSmall
// [flexibleSpace]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/Identifier-swift.struct/flexibleSpace
// [high]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/Priority/high
// [low]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/Priority/low
// [normal]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/Priority/normal
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
//   - [INSTouchBarItem.IsVisible]: A Boolean value that reflects whether or not the item is visible.
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
	IsVisible() bool

	// Topic: Configuring bar customization

	// The user-visible string identifying this item during bar customization.
	CustomizationLabel() string

	// Topic: Subclassing bar items

	// The view controller associated with this item.
	ViewController() INSViewController
	// The view associated with this item.
	View() INSView

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
// When true, this item is shown in a currently visible bar. This property is
// always false for spaces, proxy items, and groups.
//
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/isVisible
func (t NSTouchBarItem) IsVisible() bool {
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
