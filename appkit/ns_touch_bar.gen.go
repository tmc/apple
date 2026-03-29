// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTouchBar] class.
var (
	_NSTouchBarClass     NSTouchBarClass
	_NSTouchBarClassOnce sync.Once
)

func getNSTouchBarClass() NSTouchBarClass {
	_NSTouchBarClassOnce.Do(func() {
		_NSTouchBarClass = NSTouchBarClass{class: objc.GetClass("NSTouchBar")}
	})
	return _NSTouchBarClass
}

// GetNSTouchBarClass returns the class object for NSTouchBar.
func GetNSTouchBarClass() NSTouchBarClass {
	return getNSTouchBarClass()
}

type NSTouchBarClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTouchBarClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTouchBarClass) Alloc() NSTouchBar {
	rv := objc.Send[NSTouchBar](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides dynamic contextual controls in the Touch Bar of
// supported models of MacBook Pro.
//
// # Overview
// 
// On supported MacBook Pro models, the Touch Bar, above the keyboard, shows
// instances of the [NSTouchBar] class from the front-most app. Such an
// instance is called a . You define a bar to provide controls relevant to the
// user’s context. Each such control is an instance of the [NSTouchBarItem]
// class, called an .
// 
// [media-2851859]
// 
// You can provide many bars within your app, one for each responder instance;
// macOS frameworks can provide bars, as well, that can appear alongside your
// app’s bars. The system determines which bars to show at any given time.
// For example, an app that uses standard AppKit objects, such as text fields
// (instances of the [NSTextField] class), obtains appropriate bars along with
// relevant items automatically.
// 
// Refer to the following sample code projects, which demonstrate how to use
// [NSTouchBar] and related classes, including the [NSScrubber] class, with
// its rich API that lets you build a highly customized picker control:
// 
// - [Creating and Customizing the Touch Bar] - [Integrating a Toolbar and
// Touch Bar into Your App]
// 
// To use the Touch Bar, define bars in objects in your app’s responder
// chain. At run time, the system traverses up the responder chain to
// discover, combine, and show bars from your app and from frameworks you link
// against.
// 
// You can configure a bar to support dynamic composition, in which the system
// shows the bar in an expanded form that contains items from bars lower in
// the responder chain (from closer to the first responder). Because of
// dynamic composition and placement of items shown on the Touch Bar, always
// ensure that your bars appear as you expect them to, testing on the versions
// of macOS that you support.
// 
// Instances of the [NSTouchBar] class employ gesture recognizers and take
// advantage of macOS 10.12.1 event enhancements. Because of the physical
// geometry of the Touch Bar, touch events passed to gesture recognizers have
// only a meaningful , or horizontal, component.
// 
// There’s no need, and no API, for your app to know whether or not
// there’s a Touch Bar available. Whether your app is running on a machine
// that supports the Touch Bar or not, your app’s onscreen user interface
// (UI) appears and behaves the same way.
// 
// The Touch Bar is a Retina display, like the screen of a MacBook Pro. To
// perform custom drawing or animation within the Touch Bar, follow the same
// best practices that you would on the screen.
// 
// On the right side of the Touch Bar, the system supplies the
// always-available . The Control Strip gives the user access to standard
// controls for display brightness, sound volume, Siri, and so on. Your
// app’s bars appear to the left of the Control Strip. The user can choose
// to hide the Control Strip, which gives the frontmost app the entire Touch
// Bar width.
// 
// To the right of the Control Strip is a Touch ID sensor. To use Touch ID on
// supported MacBook Pro models, use methods from the [Local Authentication]
// framework.
// 
// The Touch Bar dims automatically and wakes when the user touches it.
// Don’t show alerts in the Touch Bar, and don’t use the Touch Bar for
// widgets.
// 
// For Touch Bar design guidance, read [Human Interface Guidelines].
// 
// # Bar objects
// 
// You can think of an [NSTouchBar] object (or ), with its array of
// [NSTouchBarItem] objects (or ), as analogous to a window toolbar with its
// toolbar items, or a menu with its menu items.
// 
// To provide a bar in your app, define it in an object that meets three
// requirements. The object that defines the bar must:
// 
// - Be a responder (an instance of an [NSResponder] subclass) that’s
// present within a responder chain at runtime. - Implement the [TouchBar]
// delegate method from the [NSTouchBarProvider] protocol.
// 
// The built-in responder classes conform to the [NSTouchBarProvider] protocol
// and support key-value observing (KVO), both of which are used and required
// by the [NSTouchBar] infrastructure. In the context of Touch Bar support, a
// responder instance can also be called a .
// 
// The following code shows an example implementation of the [TouchBar]
// delegate method. In this code snippet, you can see some statements related
// to bar customization.
// 
// You can take advantage of built-in KVO support to keep track of bar state,
// such as which items are visible as the user customizes and interacts with
// the Touch Bar.
// 
// If you explicitly adopt the [NSTouchBarProvider] protocol in the app
// delegate or in a window delegate, you must also explicitly send the
// associated key-value observing notifications from within your
// implementations of [NSTouchBar] methods; this lets the system respond
// appropriately to changes in the bar. To avoid the need to manually support
// KVO, use the as a bar provider, instead of the app delegate, or use a or as
// a bar provider, instead of the associated window delegate.
// 
// To programmatically invalidate a bar associated with a bar provider, such
// as because you’re changing the bar’s state, set its [TouchBar] property
// to a value of `nil`.
// 
// # Item objects
// 
// A bar itself (an [NSTouchBar] object) has no visible representation in a
// MacBook Pro Touch Bar. A user instead sees the bar’s items, each of which
// is an instance of the [NSTouchBarItem] class.
// 
// The items presented in a bar are the elements in a private array owned by
// the bar. To specify the items for a bar, you don’t fill this array
// directly, but rather rely on the bar to manage its items based on various
// groups of items and item identifiers that you do specify directly.
// 
// In specifying items for a bar you have two options, giving you flexibility
// for optimizing resource use and efficiency in your app.
// 
// - The [NSTouchBar.TemplateItems] property is a set that you can directly populate with
// item instances for a bar. Use this option when your items are lightweight
// enough to stay in memory for the duration of your app’s lifetime, and
// when they don’t contain state that might change over time. - The
// [NSTouchBarDelegate] protocol, and its [TouchBarMakeItemForIdentifier]
// delegate method, give your app a way to create items on-demand. Use this
// option when it makes more sense in terms of resource usage and reflecting
// dynamic state.
// 
// Whichever of these two approaches you employ, the system is in charge of
// populating a bar’s private items array based on three things:
// 
// - Your configuration of the bar’s item-identifiers properties - Any
// nesting you have specified - Any customization that the user has specified
// for the bar
// 
// As your app runs, you can obtain the identifiers of the items eligible for
// presentation in a bar — specifically, those in its private items array
// — by accessing the read-only [NSTouchBar.ItemIdentifiers] property. This property
// reflects the current state of the bar instance, including any customization
// that has been performed by the user and any dynamic composition that has
// been performed by the system.
// 
// # Customization
// 
// AppKit provides a rich Touch Bar customization facility for users that
// appears, upon user request, on the main display. Make your bars
// customizable unless you have a specific UI need not to do so.
// 
// A customizable bar automatically obtains onscreen UI which lets the user:
// 
// - Change which items are part of the shown bar - Rearrange items within the
// shown bar
// 
// A user invokes the onscreen customization UI by choosing a dedicated menu
// item.
// 
// To make an [NSTouchBar] object eligible for customization, assign it a
// globally-unique [NSTouchBar.CustomizationIdentifier] identifier. For the identifier
// string, use reverse-DNS style, such as
// “`com.Company()-name.App()-name.Alphanumeric()-ID`”.
// 
// Next, specify the bar’s items and customization possibilities by
// populating its item identifier lists. Each such list is an array, each of
// whose elements is the identifier (of type [NSTouchBarItemIdentifier]) for
// an item (an [NSTouchBarItem] object). A bar’s item identifier lists are:
// 
// , specified in a bar’s [NSTouchBar.DefaultItemIdentifiers] property. Always specify
// this property for an [NSTouchBar] object, even if you elect to make the bar
// noncustomizable. The system:
// 
// - Shows this list’s items by default when the system displays the bar. -
// Includes a preconfigured bar representation, containing these items, in the
// associated customization UI (when you have designated the bar as
// customizable by assigning it a [NSTouchBar.CustomizationIdentifier] property value);
// the user can drag the default bar into the Touch Bar, should they want to
// return to the default configuration.
// 
// , specified in a bar’s [NSTouchBar.CustomizationAllowedItemIdentifiers] property.
// Always configure this property for a customizable bar. The system uses this
// list by showing representations of its items individually in the
// customization UI, arranged in the same order as you specify in the property
// array. When there’s available geometric space, a user can drag in to the
// active bar any of the items in this list. If there isn’t enough space, a
// dragged item replaces the item or items under the spot the new item is
// dropped.
// 
// , specified in a bar’s [NSTouchBar.CustomizationRequiredItemIdentifiers] property.
// Configure this property at your discretion, depending on the design of your
// app. The user can’t remove from the bar any of the items you specify in
// this list.
// 
// To provide textual labels in the customization UI, use the
// [NSTouchBar.CustomizationLabel] property on each [NSTouchBarItem] instance you include
// in a customizable bar. The accessibility system in macOS also makes use of
// these labels.
// 
// If your app design requires a noncustomizable [NSTouchBar] object:
// 
// - List all of the bar’s items in the [NSTouchBar.DefaultItemIdentifiers] property,
// and only in this property. - Don’t use the other properties described in
// this section and, in particular, don’t assign the bar a
// [NSTouchBar.CustomizationIdentifier] property value.
// 
// # Group item, popover item, and composed bar customization
// 
// AppKit lets you specify any bar as customizable or not. The customization
// configuration you provide for a bar remains associated with the bar and its
// items — even when those items are nested by the system into another bar
// higher in the responder chain.
// 
// The system also respects your bar customization configuration when you use
// group and popover items. Each of these item types itself contains one or
// more bars—which can, in turn contain group items and popover items, and
// so on. The rest of this section explains how customization works for the
// bars containing, and bars within, group items and popover items.
// 
// (an instance of the [NSGroupTouchBarItem] class) has one bar, held in the
// object’s [NSTouchBar.GroupTouchBar] property. AppKit supports nesting of group
// items, in that you can configure a [NSTouchBar.GroupTouchBar] bar to itself contain
// one or more group items (or, for that matter, items of any other type,
// guided by what works well in your app).Here are some examples of how
// customization for group items works in practice:
// 
// - If you configure a bar as customizable, and give it a group item whose
// bar you configure as customizable, then the array of items in the
// [NSTouchBar.GroupTouchBar] bar appears in the customization UI as an atomic unit.
// During customization, a user can manipulate the array of items, but
// strictly as a unit: If the (noncustomizable) [NSTouchBar.GroupTouchBar] bar is visible
// in the Touch Bar, the user can remove it as a unit, or can rearrange it
// among the other items in the [NSTouchBar.GroupTouchBar] bar; if the [NSTouchBar.GroupTouchBar]
// bar is instead visible in the customization UI, the user can add it back to
// the Touch Bar, as a unit, placing it within the bar that owns the group
// item. - If you configure a bar as customizable, and give it a group item
// whose bar you configure as customizable, then the [NSTouchBar.GroupTouchBar] bar’s
// items appear in the customization UI as individual items. During
// customization, a user can manipulate each item separately: If an item from
// the [NSTouchBar.GroupTouchBar] bar is visible in the Touch Bar, the user can remove it
// or can rearrange its position individually among the other items in the
// [NSTouchBar.GroupTouchBar] bar; if an item from the [NSTouchBar.GroupTouchBar] bar is instead
// visible in the customization UI, the user can add it back to the Touch Bar,
// individually, placing it anywhere within the [NSTouchBar.GroupTouchBar] bar that owns
// it.
// 
// (an instance of the [NSPopoverTouchBarItem] class) has two bars: one bar
// you specify in its [NSTouchBar.PopoverTouchBar] property and a second, optional bar
// you can specify in its [NSTouchBar.PressAndHoldTouchBar] property. Here are some
// examples of how customization for popover items works in practice:
// 
// - If you configure a bar as customizable, and give it a popover item whose
// [NSTouchBar.PopoverTouchBar] bar you configure as customizable, the [NSTouchBar.PopoverTouchBar]
// bar never appears in the customization UI. If the user invokes the
// customization UI when the (noncustomizable) popover item itself (not the
// button’s associated [NSTouchBar.PopoverTouchBar] bar) is visible in the Touch Bar,
// the customization UI lets the user rearrange the position of the popover
// item relative to the other items in the containing bar. If, on the other
// hand, the user invokes the customization UI when the (noncustomizable)
// [NSTouchBar.PopoverTouchBar] bar is visible in the Touch Bar, the system dismisses the
// popover bar and shows, in the customization UI, customization options for
// the bar that contains the popover item. - If you configure a
// [NSTouchBar.PopoverTouchBar] bar as customizable, the user can invoke the
// [NSTouchBar.PopoverTouchBar] bar (by tapping the popover item that owns it, in the
// Touch Bar) and then use the customization UI to manipulate the items in the
// [NSTouchBar.PopoverTouchBar] bar itself.
// 
// # Customization menu item
// 
// A user invokes the customization UI for a particular [NSTouchBar] object,
// when it’s visible in the Touch Bar, by choosing the bar customization
// menu item. To enable this menu item you must explicitly opt-in, which you
// can do in the following ways:
// 
// - If you want the system to automatically name, place, validate, and
// activate this menu item in your app’s menus, set the
// [isAutomaticCustomizeTouchBarMenuItemEnabled] property of your app object
// (of type [NSApplication]) to [true]. - To explicitly place the
// customization menu item in one of your app’s menus, employ the
// [ToggleTouchBarCustomizationPalette] method of your app object. When you do
// this, the system still names and validates the menu item, and hides it on
// systems that don’t have a Touch Bar.
// 
// If you attempt to employ the customization menu item (using either of these
// two approaches), but do not provide a customization identifier property
// ([NSTouchBar.CustomizationIdentifier]) for a bar, the customization menu item appears
// when that bar is active — but the menu item, in this case, is disabled.
// 
// If your app attempts to use both automatic and explicit placement of the
// customization menu item, the system respects your explicit control and
// doesn’t place the item automatically.
// 
// # Layout
// 
// The user controls the width of the Control Strip and can choose to hide it,
// and the system is in charge of the nesting of [NSTouchBar] instances (for
// the bars you make eligible for composition). As a result, the available
// display width for your bars can vary. There’s no API for you to obtain
// the current available display width.
// 
// In your layout design, don’t depend on a particular Control Strip size.
// Do anticipate dynamic composition and nesting for your bars.
// 
// If you need more horizontal space than might be available, use a popover
// item, a scrubber, or a scroll view — as they fit your design needs, but
// in that, descending, order of preference.
// 
// In geometric-space-constrained scenarios, the system hides [NSTouchBarItem]
// instances according to their visibility priority.
// 
// If you need to center an item in the Touch Bar, designate it as a by
// assigning it to its bar’s [NSTouchBar.PrincipalItemIdentifier] property. Don’t
// hard-code spacing in an attempt to ensure an item is centered. If you want
// a group of items to appear centered in the Touch Bar, designate the group
// item (of type [NSGroupTouchBarItem]) as the principal item.
// 
// # Composition and nesting
// 
// You can configure a bar to support dynamic composition, in which the system
// shows the bar in an expanded form that contains items from bars lower in
// the responder chain (closer to the first responder).
// 
// To allow a bar to serve as a container for nesting, add the
// [otherItemsProxy] item identifier to the bar’s [NSTouchBar.DefaultItemIdentifiers]
// array. A bar that includes this identifier, and that’s relatively higher
// in the responder chain, can then (at runtime) include the items from an
// eligible bar relatively lower in the responder chain.
// 
// The position that you specify for the other-items proxy, within a bar’s
// [NSTouchBar.DefaultItemIdentifiers] array, tells the system where you want nested
// items to be placed.
// 
// The system determines whether or not to compose bars in this way, based on
// system policy and available geometric space in the Touch Bar.
// 
// [NSTouchBar] object nesting can be chained, according to available
// geometric space in the Touch Bar. For example, a view and a text field
// within that view could each contribute their items to the bar defined for a
// parent window controller.
// 
// When the system nests one bar’s items into another bar higher in the
// responder chain, the items appear to the user, in the Touch Bar, as fully
// incorporated into the higher bar. There’s no visual boundary or
// additional spacing to distinguish the items as being nested.
// 
// If a bar doesn’t employ the [otherItemsProxy] identifier, the system
// hides that bar when another bar, lower in the responder chain, is eligible
// for display.
// 
// When determining which items to show in the Touch Bar for the current first
// responder, the system traverses up the entire responder chain. This lets
// the system accommodate any proxy items in bars defined for objects higher
// in the chain, thereby respecting the fact that any bar, defined for an
// object at any position in the responder chain, might include the
// other-items-proxy identifier.
// 
// # Customization for composed bars
// 
// The logical, geometric boundary for a nested [NSTouchBar] object isn’t
// visible to the user in the Touch Bar. However, the boundary remains in
// effect in terms of customization. A user can’t rearrange a nested bar’s
// items outside of its boundary.
// 
// For example, say you have a bar, higher in the responder chain, configured
// like this:
// 
// And say you also have a bar, lower in the responder chain, eligible for
// display in the Touch Bar according to the system and the current app state,
// configured like this:
// 
// The composed bar would correspond to this arrangement in the Touch Bar:
// 
// With the customization UI, the user could then rearrange the items
// represented here by `(A)`, `(B)`, and `(C)`, but only as long as those
// items remained contiguous during the rearrangement, thereby respecting the
// logical boundary of the bar that defines them.
// 
// # Item spacing for composed bars
// 
// When the system nests [NSTouchBar] objects that include spacing items, it
// merges any resulting adjacent spacing. Ensure that your bars appear as you
// expect them to, testing on the versions of macOS that you support. For more
// on spacing items, see [NSTouchBarItem].
// 
// # Bar discovery and the responder chain
// 
// At runtime, the system traverses up the responder chain, starting at the
// object with focus, to discover objects that conform to the
// [NSTouchBarProvider] protocol. Such objects are called . The system then
// populates the Touch Bar, potentially with multiple, nested bars, according
// to system policy and available geometric space.
// 
// Specifically, bar discovery by the system proceeds in the following order:
// 
// - Key window’s first responder - Key window - Key window’s delegate -
// Key window’s controller - Main window’s first responder - Main window -
// Main window’s delegate - Main window’s controller - App object - App
// delegate
// 
// When the system encounters a bar provider that’s an instance of an
// [NSResponder] subclass, the system then additionally searches up the
// responder chain anchored at that object.
// 
// For example, in a complicated but otherwise standard app, bar discovery
// might proceed in this order:
// 
// - Key window’s first responder - View controller of key window’s first
// responder - Intermediate view controllers and views - View that’s closest
// to root of window - View controller that’s closest to root of window -
// Key window - Key window’s controller - App object - App delegate
// 
// The Touch Bar can show one bar nested within another, as described in
// [NSTouchBar].
// 
// # Accessibility and the Touch Bar
// 
// AppKit views and controls adopt the [NSAccessibilityProtocol] protocol and
// automatically send appropriate accessibility notifications. Because the
// Touch Bar is designed to work with AppKit, it’s fully accessible.
// 
// Be sure to use the `customizationLabel` property on every [NSTouchBarItem]
// instance that you designate as customizable. The accessibility system in
// macOS makes use of these labels.
// 
// To learn more about accessibility, read [Accessibility for AppKit].
// 
// # AppKit support for the Touch Bar
// 
// To support the Touch Bar feature, AppKit provides several enhancements,
// first available in macOS 10.12.1:
// 
// - The [NSScrubber] class, along with related APIs, provide a way for you to
// add a flexible, horizontally-oriented picker to a custom item (an instance
// of the [NSCustomTouchBarItem] class). - You can use the
// [NSMagnificationGestureRecognizer] class in bar items. To enable two-finger
// pinch gestures, set the recognizer’s [NSTouchBar.AllowedTouchTypes] mask property,
// on the gesture recognizer, to the [NSTouch.TouchType.direct] constant from
// the [NSTouch.TouchTypeMask] enumeration. - The [NSGestureRecognizer]
// abstract class is enhanced with a set of methods that let you implement
// responses to touch events: [TouchesBeganWithEvent],
// [TouchesCancelledWithEvent], [TouchesEndedWithEvent], and
// [TouchesMovedWithEvent]. - The [NSClickGestureRecognizer],
// [NSPanGestureRecognizer], and [NSPressGestureRecognizer] concrete classes
// are each enhanced with a `numberOfTouchesRequired` property to let you
// specify the number of touches required for a gesture match. - To enable
// touch events in a custom view, you must set the value of a view’s
// [NSTouchBar.AllowedTouchTypes] property to a value of [TouchTypeMaskDirect]. (In macOS
// 10.12.1, the [acceptsTouchEvents] property is deprecated in favor of the
// new [NSTouchBar.AllowedTouchTypes] property.) - The [NSTouch] class has a new property
// and two new methods for supporting the Touch Bar: [NSTouch.TouchType],
// [LocationInView], and [PreviousLocationInView]. - The [NSButton],
// [NSSegmentedControl], and [NSSlider] classes are each enhanced with
// appearance support properties: [NSTouchBar.BezelColor] for buttons,
// [NSTouchBar.SelectedSegmentBezelColor] for segmented controls, and [NSTouchBar.TrackFillColor]
// for sliders. (With the Touch Bar, you employ sliders indirectly, as used by
// slider items.) - Starting in macOS 10.12, you can use a rich set of
// convenience initializers for controls. These initializers simplify the
// definition of bar items and take care of appearance and sizing for the
// Touch Bar. In particular, the [NSButton], [NSSegmentedControl], and
// [NSSlider] classes now offer a variety of convenience initializers such as
// [ButtonWithTitleImageTargetAction]. - Methods and properties in the
// [NSSpellChecker], [NSTextField], and [NSTextView] classes, and in the
// [NSTextFieldDelegate] and [NSTextViewDelegate] protocols, support using the
// Touch Bar for spell checking, predictive text suggestion, text completion,
// and automatic handling of trailing space. For example: - When you use an
// [NSTextView] object, you gain automatic Touch Bar support for text styling
// and predictive text suggestions. - When you use an
// [NSCandidateListTouchBarItem] object, you can use the
// [RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler]
// method. This method provides a completion handler that you can use to
// filter or otherwise manage the candidate text. - AppKit adds many new
// template images for you to use in your [NSTouchBarItem] objects. A few
// examples of these images are: [touchBarAddTemplateName],
// [touchBarComposeTemplateName], [touchBarGoBackTemplateName],
// [touchBarGoForwardTemplateName], and [NSImageNameTouchBarHomeTemplate].
// Always use templates for images in your items: they respond automatically
// to system white-point changes. Note that these images are exclusively for
// use in the Touch Bar and in onscreen windows. For a complete list of these
// template images, see [NSTouchBarItem]. - If your UI for a popover item
// needs more horizontal space, you can use a scroll view (an instance of the
// [NSScrollView] class). , in this case, enable the popover item’s
// press-and-hold option, because doing so interferes with scrolling. - You
// can group bar items by using an instance of the [NSStackView] class.
// However, doing so loses system support for spacing. When you instead place
// items into a group item (an instance of the [NSGroupTouchBarItem] class),
// the system: - Manages inter-item spacing - Supports user customization for
// the individual items
// 
// # Development considerations for the Touch Bar
// 
// The Xcode Touch Bar simulator represents the Touch Bar onscreen and
// supports some user interaction. However, some interactions are unavailable
// in the simulator. For example, you can’t perform two-finger gestures in
// the Touch Bar simulator.
// 
// If you’re adopting Touch Bar support for your app, but running Xcode on a
// Mac without a Touch Bar, you can enable the Xcode Touch Bar simulator by
// choosing Window > Show Touch Bar in Xcode.
// 
// Interface Builder supports development for the Touch Bar with nib objects
// available in the object library. Drag and drop bars and items from the
// object library into your canvas, and then attach them to your app’s
// responders as desired. For more information, see Xcode Help.
// 
// # Performance considerations for the Touch Bar
// 
// The Touch Bar’s display and the MacBook screen share resources, including
// the main CPU and GPU of the MacBook. To ensure that your Touch Bar controls
// perform well, follow the usual best practice of protecting your app’s
// main thread from doing too much work. For example, don’t perform
// rendering work for the main display on the main thread.
// 
// In addition, pay attention to the relative amounts of time your app spends
// on updates to the main display relative to updates to the Touch Bar. The
// optimum ratio can vary according to what the user is doing. For example:
// 
// - When a user is interacting with the Touch Bar watching the Touch Bar,
// such as to control audio, ensure that your app gives priority to updating
// the Touch Bar. - When a user is instead interacting with the Touch Bar but
// watching the main display, such as when using a scrubber to browse pages of
// content, balance your app’s display update work between the Touch Bar and
// the main display.
// 
// Always test Touch Bar performance using the specific MacBook hardware you
// support. Specifically, don’t rely on the Xcode Touch Bar simulator when
// tuning your app for Touch Bar performance.
//
// [Accessibility for AppKit]: https://developer.apple.com/documentation/AppKit/accessibility-for-appkit
// [Creating and Customizing the Touch Bar]: https://developer.apple.com/documentation/AppKit/creating-and-customizing-the-touch-bar
// [Human Interface Guidelines]: https://developer.apple.com/design/human-interface-guidelines/inputs/touch-bar/
// [Integrating a Toolbar and Touch Bar into Your App]: https://developer.apple.com/documentation/AppKit/integrating-a-toolbar-and-touch-bar-into-your-app
// [Local Authentication]: https://developer.apple.com/documentation/LocalAuthentication
// [NSTouch.TouchType.direct]: https://developer.apple.com/documentation/AppKit/NSTouch/TouchType/direct
// [NSTouch.TouchTypeMask]: https://developer.apple.com/documentation/AppKit/NSTouch/TouchTypeMask
// [NSTouch.TouchType]: https://developer.apple.com/documentation/AppKit/NSTouch/TouchType
// [acceptsTouchEvents]: https://developer.apple.com/documentation/AppKit/NSView/acceptsTouchEvents
// [isAutomaticCustomizeTouchBarMenuItemEnabled]: https://developer.apple.com/documentation/AppKit/NSApplication/isAutomaticCustomizeTouchBarMenuItemEnabled
// [otherItemsProxy]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/Identifier-swift.struct/otherItemsProxy
// [touchBarAddTemplateName]: https://developer.apple.com/documentation/AppKit/NSImage/touchBarAddTemplateName
// [touchBarComposeTemplateName]: https://developer.apple.com/documentation/AppKit/NSImage/touchBarComposeTemplateName
// [touchBarGoBackTemplateName]: https://developer.apple.com/documentation/AppKit/NSImage/touchBarGoBackTemplateName
// [touchBarGoForwardTemplateName]: https://developer.apple.com/documentation/AppKit/NSImage/touchBarGoForwardTemplateName
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Creating a bar
//
//   - [NSTouchBar.InitWithCoder]: Creates a Touch Bar object from a coder object provided by a storyboard or NIB file.
//
// # Providing bar items
//
//   - [NSTouchBar.Delegate]: The delegate that provides items to the Touch Bar.
//   - [NSTouchBar.SetDelegate]
//   - [NSTouchBar.TemplateItems]: The primary source of items that the Touch Bar uses to fill its private items array, unless you provide items using a delegate.
//   - [NSTouchBar.SetTemplateItems]
//   - [NSTouchBar.DefaultItemIdentifiers]: A required list of identifiers for items that you want to appear in the Touch Bar after instantiating it.
//   - [NSTouchBar.SetDefaultItemIdentifiers]
//   - [NSTouchBar.PrincipalItemIdentifier]: The identifier of an item you want the system to center in the Touch Bar.
//   - [NSTouchBar.SetPrincipalItemIdentifier]
//   - [NSTouchBar.EscapeKeyReplacementItemIdentifier]: The identifier of an item that replaces the system-provided button in the Touch Bar.
//   - [NSTouchBar.SetEscapeKeyReplacementItemIdentifier]
//
// # Observing bar status
//
//   - [NSTouchBar.Visible]: A Boolean value that Indicates whether the Touch Bar is eligible for display.
//   - [NSTouchBar.ItemIdentifiers]: The list of identifiers for the current items in the Touch Bar.
//   - [NSTouchBar.ItemForIdentifier]: Returns the Touch Bar item that corresponds to a given identifier.
//
// # Configuring user customization
//
//   - [NSTouchBar.CustomizationIdentifier]: A globally unique string that makes the Touch Bar eligible for user customization.
//   - [NSTouchBar.SetCustomizationIdentifier]
//   - [NSTouchBar.CustomizationAllowedItemIdentifiers]: A list of identifiers for items to show in the Touch Bar’s customization UI.
//   - [NSTouchBar.SetCustomizationAllowedItemIdentifiers]
//   - [NSTouchBar.CustomizationRequiredItemIdentifiers]: An optional list of identifiers for items you want to always appear in the Touch Bar and which the user can’t remove during customization.
//   - [NSTouchBar.SetCustomizationRequiredItemIdentifiers]
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBar
type NSTouchBar struct {
	objectivec.Object
}

// NSTouchBarFromID constructs a [NSTouchBar] from an objc.ID.
//
// An object that provides dynamic contextual controls in the Touch Bar of
// supported models of MacBook Pro.
func NSTouchBarFromID(id objc.ID) NSTouchBar {
	return NSTouchBar{objectivec.Object{ID: id}}
}
// NOTE: NSTouchBar adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTouchBar] class.
//
// # Creating a bar
//
//   - [INSTouchBar.InitWithCoder]: Creates a Touch Bar object from a coder object provided by a storyboard or NIB file.
//
// # Providing bar items
//
//   - [INSTouchBar.Delegate]: The delegate that provides items to the Touch Bar.
//   - [INSTouchBar.SetDelegate]
//   - [INSTouchBar.TemplateItems]: The primary source of items that the Touch Bar uses to fill its private items array, unless you provide items using a delegate.
//   - [INSTouchBar.SetTemplateItems]
//   - [INSTouchBar.DefaultItemIdentifiers]: A required list of identifiers for items that you want to appear in the Touch Bar after instantiating it.
//   - [INSTouchBar.SetDefaultItemIdentifiers]
//   - [INSTouchBar.PrincipalItemIdentifier]: The identifier of an item you want the system to center in the Touch Bar.
//   - [INSTouchBar.SetPrincipalItemIdentifier]
//   - [INSTouchBar.EscapeKeyReplacementItemIdentifier]: The identifier of an item that replaces the system-provided button in the Touch Bar.
//   - [INSTouchBar.SetEscapeKeyReplacementItemIdentifier]
//
// # Observing bar status
//
//   - [INSTouchBar.Visible]: A Boolean value that Indicates whether the Touch Bar is eligible for display.
//   - [INSTouchBar.ItemIdentifiers]: The list of identifiers for the current items in the Touch Bar.
//   - [INSTouchBar.ItemForIdentifier]: Returns the Touch Bar item that corresponds to a given identifier.
//
// # Configuring user customization
//
//   - [INSTouchBar.CustomizationIdentifier]: A globally unique string that makes the Touch Bar eligible for user customization.
//   - [INSTouchBar.SetCustomizationIdentifier]
//   - [INSTouchBar.CustomizationAllowedItemIdentifiers]: A list of identifiers for items to show in the Touch Bar’s customization UI.
//   - [INSTouchBar.SetCustomizationAllowedItemIdentifiers]
//   - [INSTouchBar.CustomizationRequiredItemIdentifiers]: An optional list of identifiers for items you want to always appear in the Touch Bar and which the user can’t remove during customization.
//   - [INSTouchBar.SetCustomizationRequiredItemIdentifiers]
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBar
type INSTouchBar interface {
	objectivec.IObject

	// Topic: Creating a bar

	// Creates a Touch Bar object from a coder object provided by a storyboard or NIB file.
	InitWithCoder(coder foundation.INSCoder) NSTouchBar

	// Topic: Providing bar items

	// The delegate that provides items to the Touch Bar.
	Delegate() NSTouchBarDelegate
	SetDelegate(value NSTouchBarDelegate)
	// The primary source of items that the Touch Bar uses to fill its private items array, unless you provide items using a delegate.
	TemplateItems() foundation.INSSet
	SetTemplateItems(value foundation.INSSet)
	// A required list of identifiers for items that you want to appear in the Touch Bar after instantiating it.
	DefaultItemIdentifiers() []string
	SetDefaultItemIdentifiers(value []string)
	// The identifier of an item you want the system to center in the Touch Bar.
	PrincipalItemIdentifier() NSTouchBarItemIdentifier
	SetPrincipalItemIdentifier(value NSTouchBarItemIdentifier)
	// The identifier of an item that replaces the system-provided button in the Touch Bar.
	EscapeKeyReplacementItemIdentifier() NSTouchBarItemIdentifier
	SetEscapeKeyReplacementItemIdentifier(value NSTouchBarItemIdentifier)

	// Topic: Observing bar status

	// A Boolean value that Indicates whether the Touch Bar is eligible for display.
	Visible() bool
	// The list of identifiers for the current items in the Touch Bar.
	ItemIdentifiers() []string
	// Returns the Touch Bar item that corresponds to a given identifier.
	ItemForIdentifier(identifier NSTouchBarItemIdentifier) INSTouchBarItem

	// Topic: Configuring user customization

	// A globally unique string that makes the Touch Bar eligible for user customization.
	CustomizationIdentifier() NSTouchBarCustomizationIdentifier
	SetCustomizationIdentifier(value NSTouchBarCustomizationIdentifier)
	// A list of identifiers for items to show in the Touch Bar’s customization UI.
	CustomizationAllowedItemIdentifiers() []string
	SetCustomizationAllowedItemIdentifiers(value []string)
	// An optional list of identifiers for items you want to always appear in the Touch Bar and which the user can’t remove during customization.
	CustomizationRequiredItemIdentifiers() []string
	SetCustomizationRequiredItemIdentifiers(value []string)

	// A Boolean value indicating whether the view accepts touch events.
	AcceptsTouchEvents() bool
	SetAcceptsTouchEvents(value bool)
	AllowedTouchTypes() NSTouchTypeMask
	SetAllowedTouchTypes(value NSTouchTypeMask)
	// The color of the button’s bezel, in appearances that support it.
	BezelColor() INSColor
	SetBezelColor(value INSColor)
	// The user-visible string identifying this item during bar customization.
	CustomizationLabel() string
	SetCustomizationLabel(value string)
	// A direct touch from a user’s finger on a screen.
	Direct() NSTouchTypeMask
	SetDirect(value NSTouchTypeMask)
	// A bar that holds this group’s items.
	GroupTouchBar() INSTouchBar
	SetGroupTouchBar(value INSTouchBar)
	// The bar displayed when this item is “popped.”
	PopoverTouchBar() INSTouchBar
	SetPopoverTouchBar(value INSTouchBar)
	// The bar that is displayed when a user press-and-holds on the popover item.
	PressAndHoldTouchBar() INSTouchBar
	SetPressAndHoldTouchBar(value INSTouchBar)
	// The color of the selected segment’s bezel, in appearances that support it.
	SelectedSegmentBezelColor() INSColor
	SetSelectedSegmentBezelColor(value INSColor)
	// The property you implement to provide a Touch Bar object.
	TouchBar() INSTouchBar
	SetTouchBar(value INSTouchBar)
	// The color of the filled portion of the slider track, in appearances that support it.
	TrackFillColor() INSColor
	SetTrackFillColor(value INSColor)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (t NSTouchBar) Init() NSTouchBar {
	rv := objc.Send[NSTouchBar](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTouchBar) Autorelease() NSTouchBar {
	rv := objc.Send[NSTouchBar](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTouchBar creates a new NSTouchBar instance.
func NewNSTouchBar() NSTouchBar {
	class := getNSTouchBarClass()
	rv := objc.Send[NSTouchBar](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a Touch Bar object from a coder object provided by a storyboard or
// NIB file.
//
// # Return Value
// 
// A fully initialized Touch Bar object, or `nil` if the coder doesn’t
// define a Touch Bar object.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBar/init(coder:)
func NewTouchBarWithCoder(coder foundation.INSCoder) NSTouchBar {
	instance := getNSTouchBarClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTouchBarFromID(rv)
}

// Creates a Touch Bar object from a coder object provided by a storyboard or
// NIB file.
//
// # Return Value
// 
// A fully initialized Touch Bar object, or `nil` if the coder doesn’t
// define a Touch Bar object.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBar/init(coder:)
func (t NSTouchBar) InitWithCoder(coder foundation.INSCoder) NSTouchBar {
	rv := objc.Send[NSTouchBar](t.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
// Returns the Touch Bar item that corresponds to a given identifier.
//
// # Return Value
// 
// A Touch Bar item if one exists for the given identifier; otherwise, returns
// `nil`.
//
// # Discussion
// 
// The system returns items (instances of the [NSTouchBarItem] class) as it
// finds them, according to the following search order, listed here from
// first-searched to last-searched:
// 
// - Items in the bar’s private array, which are reflected in the value of
// the [ItemIdentifiers] array. - Items in the [TemplateItems] array. - Items
// returned from the bar delegate’s [TouchBarMakeItemForIdentifier] method.
// 
// Your app never needs to instantiate spacing or proxy items because these
// are created by the system directly, according to their identifiers, as
// shown in the table below.
// 
// [Table data omitted]
// 
// For more on the proxy placeholder, see [NSTouchBar].
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBar/item(forIdentifier:)
func (t NSTouchBar) ItemForIdentifier(identifier NSTouchBarItemIdentifier) INSTouchBarItem {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("itemForIdentifier:"), objc.String(string(identifier)))
	return NSTouchBarItemFromID(rv)
}
func (t NSTouchBar) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The delegate that provides items to the Touch Bar.
//
// # Discussion
// 
// Employ a bar delegate, according to the needs of your app, to dynamically
// create items ([NSTouchBarItem] instances). For more information, see
// [NSTouchBar].
// 
// This property is conditionally archived, as described in the
// [encodeConditionalObject(_:forKey:)] method.
//
// [encodeConditionalObject(_:forKey:)]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/encodeConditionalObject(_:forKey:)
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBar/delegate
func (t NSTouchBar) Delegate() NSTouchBarDelegate {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("delegate"))
	return NSTouchBarDelegateObjectFromID(rv)
}
func (t NSTouchBar) SetDelegate(value NSTouchBarDelegate) {
	objc.Send[struct{}](t.ID, objc.Sel("setDelegate:"), value)
}
// The primary source of items that the Touch Bar uses to fill its private
// items array, unless you provide items using a delegate.
//
// # Discussion
// 
// When a bar needs to fill its private items array with items
// ([NSTouchBarItem] instances), it employs a variety of potential sources.
// The first place it looks is this property. For more information, see
// [NSTouchBar].
// 
// The system archives this property.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBar/templateItems
func (t NSTouchBar) TemplateItems() foundation.INSSet {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("templateItems"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (t NSTouchBar) SetTemplateItems(value foundation.INSSet) {
	objc.Send[struct{}](t.ID, objc.Sel("setTemplateItems:"), value)
}
// A required list of identifiers for items that you want to appear in the
// Touch Bar after instantiating it.
//
// # Discussion
// 
// Always specify this property for an [NSTouchBar] object, even if you elect
// to make the bar noncustomizable. The system:
// 
// - Shows this list’s items by default when the system displays the bar. -
// Includes a preconfigured bar, containing these items, in the associated
// customization UI (when you have assigned the bar a
// [CustomizationIdentifier] property value); the user can drag the default
// bar into the Touch Bar, should they want to return to the default
// configuration.
// 
// The system archives this property.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBar/defaultItemIdentifiers
func (t NSTouchBar) DefaultItemIdentifiers() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("defaultItemIdentifiers"))
	return objc.ConvertSliceToStrings(rv)
}
func (t NSTouchBar) SetDefaultItemIdentifiers(value []string) {
	objc.Send[struct{}](t.ID, objc.Sel("setDefaultItemIdentifiers:"), objectivec.StringSliceToNSArray(value))
}
// The identifier of an item you want the system to center in the Touch Bar.
//
// # Discussion
// 
// The system attempts to center a principal item within the Touch Bar. If you
// want a group of items to appear centered in the Touch Bar, designate the
// group item (of type [NSGroupTouchBarItem]) as the principal item.
// 
// If more than one bar in the responder chain is eligible to be visible in
// the Touch Bar, and more than one of those has a principal item, the system
// determines which one to center in the Touch Bar.
// 
// The system archives this property.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBar/principalItemIdentifier
func (t NSTouchBar) PrincipalItemIdentifier() NSTouchBarItemIdentifier {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("principalItemIdentifier"))
	return NSTouchBarItemIdentifier(foundation.NSStringFromID(rv).String())
}
func (t NSTouchBar) SetPrincipalItemIdentifier(value NSTouchBarItemIdentifier) {
	objc.Send[struct{}](t.ID, objc.Sel("setPrincipalItemIdentifier:"), objc.String(string(value)))
}
// The identifier of an item that replaces the system-provided button in the
// Touch Bar.
//
// # Discussion
// 
// To replace the system-provided button, set
// [EscapeKeyReplacementItemIdentifier] to the identifier of an
// [NSTouchBarItem] instance. You must also include the item instance in the
// Touch Bar’s [TemplateItems] array, or return the item when the Touch Bar
// calls the [TouchBarMakeItemForIdentifier] delegate method.
// 
// Devices that include the Touch Bar (second generation) have a dedicated
// Escape key located to the left of the Touch Bar. By default, the Touch Bar
// hides the system-provided button and its replacement on these devices.
// However, if you set the replacement item’s [VisibilityPriority] to
// [high], the Touch Bar shows the item and reduces the visible space of the
// app region.
//
// [high]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/Priority/high
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBar/escapeKeyReplacementItemIdentifier
func (t NSTouchBar) EscapeKeyReplacementItemIdentifier() NSTouchBarItemIdentifier {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("escapeKeyReplacementItemIdentifier"))
	return NSTouchBarItemIdentifier(foundation.NSStringFromID(rv).String())
}
func (t NSTouchBar) SetEscapeKeyReplacementItemIdentifier(value NSTouchBarItemIdentifier) {
	objc.Send[struct{}](t.ID, objc.Sel("setEscapeKeyReplacementItemIdentifier:"), objc.String(string(value)))
}
// A Boolean value that Indicates whether the Touch Bar is eligible for
// display.
//
// # Discussion
// 
// A value of [true] indicates that the bar is attached to an eligible bar
// provider and that its items are displayable, assuming adequate geometric
// space. A is an object that conforms to the [NSTouchBarProvider] protocol.
// For more on bar providers, read [NSTouchBar].
// 
// This property is key–value observable.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBar/isVisible
func (t NSTouchBar) Visible() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isVisible"))
	return rv
}
// The list of identifiers for the current items in the Touch Bar.
//
// # Discussion
// 
// If the user has not customized the bar, this property’s value is the same
// as that of the [DefaultItemIdentifiers] property.
// 
// The system archive this property.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBar/itemIdentifiers
func (t NSTouchBar) ItemIdentifiers() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("itemIdentifiers"))
	return objc.ConvertSliceToStrings(rv)
}
// A globally unique string that makes the Touch Bar eligible for user
// customization.
//
// # Discussion
// 
// To make an [NSTouchBar] object eligible for user customization, assign it a
// globally unique [CustomizationIdentifier] identifier. For the identifier
// string, use reverse-DNS style, such as
// “`com.Company()-name.App()-name.Alphanumeric()-ID`”.
// 
// The system archives this property.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBar/customizationIdentifier-swift.property
func (t NSTouchBar) CustomizationIdentifier() NSTouchBarCustomizationIdentifier {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("customizationIdentifier"))
	return NSTouchBarCustomizationIdentifier(foundation.NSStringFromID(rv).String())
}
func (t NSTouchBar) SetCustomizationIdentifier(value NSTouchBarCustomizationIdentifier) {
	objc.Send[struct{}](t.ID, objc.Sel("setCustomizationIdentifier:"), objc.String(string(value)))
}
// A list of identifiers for items to show in the Touch Bar’s customization
// UI.
//
// # Discussion
// 
// The customization UI shows these items in additional to the items in
// [DefaultItemIdentifiers].
// 
// The items you include in [CustomizationAllowedItemIdentifiers] appear
// individually in the customization UI, arranged in the same order as you
// specify in the array. As long as there’s available geometric space, a
// user can drag in to the associated bar any of the items in this list.
// 
// Always configure this property for a customizable bar.
// 
// The system archives this property.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBar/customizationAllowedItemIdentifiers
func (t NSTouchBar) CustomizationAllowedItemIdentifiers() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("customizationAllowedItemIdentifiers"))
	return objc.ConvertSliceToStrings(rv)
}
func (t NSTouchBar) SetCustomizationAllowedItemIdentifiers(value []string) {
	objc.Send[struct{}](t.ID, objc.Sel("setCustomizationAllowedItemIdentifiers:"), objectivec.StringSliceToNSArray(value))
}
// An optional list of identifiers for items you want to always appear in the
// Touch Bar and which the user can’t remove during customization.
//
// # Discussion
// 
// Configure this property at your discretion, depending on the design of your
// app.
// 
// The system archives this property.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBar/customizationRequiredItemIdentifiers
func (t NSTouchBar) CustomizationRequiredItemIdentifiers() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("customizationRequiredItemIdentifiers"))
	return objc.ConvertSliceToStrings(rv)
}
func (t NSTouchBar) SetCustomizationRequiredItemIdentifiers(value []string) {
	objc.Send[struct{}](t.ID, objc.Sel("setCustomizationRequiredItemIdentifiers:"), objectivec.StringSliceToNSArray(value))
}
// A Boolean value indicating whether the view accepts touch events.
//
// See: https://developer.apple.com/documentation/appkit/nsview/acceptstouchevents
func (t NSTouchBar) AcceptsTouchEvents() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("acceptsTouchEvents"))
	return rv
}
func (t NSTouchBar) SetAcceptsTouchEvents(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAcceptsTouchEvents:"), value)
}
// See: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/allowedtouchtypes
func (t NSTouchBar) AllowedTouchTypes() NSTouchTypeMask {
	rv := objc.Send[NSTouchTypeMask](t.ID, objc.Sel("allowedTouchTypes"))
	return NSTouchTypeMask(rv)
}
func (t NSTouchBar) SetAllowedTouchTypes(value NSTouchTypeMask) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowedTouchTypes:"), value)
}
// The color of the button’s bezel, in appearances that support it.
//
// See: https://developer.apple.com/documentation/appkit/nsbutton/bezelcolor
func (t NSTouchBar) BezelColor() INSColor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("bezelColor"))
	return NSColorFromID(objc.ID(rv))
}
func (t NSTouchBar) SetBezelColor(value INSColor) {
	objc.Send[struct{}](t.ID, objc.Sel("setBezelColor:"), value)
}
// The user-visible string identifying this item during bar customization.
//
// See: https://developer.apple.com/documentation/appkit/nstouchbaritem/customizationlabel
func (t NSTouchBar) CustomizationLabel() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("customizationLabel"))
	return foundation.NSStringFromID(rv).String()
}
func (t NSTouchBar) SetCustomizationLabel(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setCustomizationLabel:"), objc.String(value))
}
// A direct touch from a user’s finger on a screen.
//
// See: https://developer.apple.com/documentation/appkit/nstouch/touchtypemask/direct
func (t NSTouchBar) Direct() NSTouchTypeMask {
	rv := objc.Send[NSTouchTypeMask](t.ID, objc.Sel("NSTouchTypeMaskDirect"))
	return NSTouchTypeMask(rv)
}
func (t NSTouchBar) SetDirect(value NSTouchTypeMask) {
	objc.Send[struct{}](t.ID, objc.Sel("setNSTouchTypeMaskDirect:"), value)
}
// A bar that holds this group’s items.
//
// See: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/grouptouchbar
func (t NSTouchBar) GroupTouchBar() INSTouchBar {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("groupTouchBar"))
	return NSTouchBarFromID(objc.ID(rv))
}
func (t NSTouchBar) SetGroupTouchBar(value INSTouchBar) {
	objc.Send[struct{}](t.ID, objc.Sel("setGroupTouchBar:"), value)
}
// The bar displayed when this item is “popped.”
//
// See: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/popovertouchbar
func (t NSTouchBar) PopoverTouchBar() INSTouchBar {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("popoverTouchBar"))
	return NSTouchBarFromID(objc.ID(rv))
}
func (t NSTouchBar) SetPopoverTouchBar(value INSTouchBar) {
	objc.Send[struct{}](t.ID, objc.Sel("setPopoverTouchBar:"), value)
}
// The bar that is displayed when a user press-and-holds on the popover item.
//
// See: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/pressandholdtouchbar
func (t NSTouchBar) PressAndHoldTouchBar() INSTouchBar {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("pressAndHoldTouchBar"))
	return NSTouchBarFromID(objc.ID(rv))
}
func (t NSTouchBar) SetPressAndHoldTouchBar(value INSTouchBar) {
	objc.Send[struct{}](t.ID, objc.Sel("setPressAndHoldTouchBar:"), value)
}
// The color of the selected segment’s bezel, in appearances that support
// it.
//
// See: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/selectedsegmentbezelcolor
func (t NSTouchBar) SelectedSegmentBezelColor() INSColor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("selectedSegmentBezelColor"))
	return NSColorFromID(objc.ID(rv))
}
func (t NSTouchBar) SetSelectedSegmentBezelColor(value INSColor) {
	objc.Send[struct{}](t.ID, objc.Sel("setSelectedSegmentBezelColor:"), value)
}
// The property you implement to provide a Touch Bar object.
//
// See: https://developer.apple.com/documentation/appkit/nstouchbarprovider/touchbar
func (t NSTouchBar) TouchBar() INSTouchBar {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("touchBar"))
	return NSTouchBarFromID(objc.ID(rv))
}
func (t NSTouchBar) SetTouchBar(value INSTouchBar) {
	objc.Send[struct{}](t.ID, objc.Sel("setTouchBar:"), value)
}
// The color of the filled portion of the slider track, in appearances that
// support it.
//
// See: https://developer.apple.com/documentation/appkit/nsslider/trackfillcolor
func (t NSTouchBar) TrackFillColor() INSColor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("trackFillColor"))
	return NSColorFromID(objc.ID(rv))
}
func (t NSTouchBar) SetTrackFillColor(value INSColor) {
	objc.Send[struct{}](t.ID, objc.Sel("setTrackFillColor:"), value)
}

// A Boolean value indicating whether the main menu contains an item for
// customizing the contents of the Touch Bar.
//
// # Discussion
// 
// When the value of this property is [true], AppKit adds a standard item to
// the app’s View menu that users can select to customize the Touch Bar
// contents, but only if a Touch Bar is present. If the View menu is
// unavailable, AppKit adds the item to either the Windows or App menu.
// 
// If you prefer to provide a customize menu item, set
// [AutomaticCustomizeTouchBarMenuItemEnabled] to [false], and create the menu
// item with an action of [ToggleTouchBarCustomizationPalette].
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBar/isAutomaticCustomizeTouchBarMenuItemEnabled
func (_NSTouchBarClass NSTouchBarClass) AutomaticCustomizeTouchBarMenuItemEnabled() bool {
	rv := objc.Send[bool](objc.ID(_NSTouchBarClass.class), objc.Sel("isAutomaticCustomizeTouchBarMenuItemEnabled"))
	return rv
}
func (_NSTouchBarClass NSTouchBarClass) SetAutomaticCustomizeTouchBarMenuItemEnabled(value bool) {
	objc.Send[struct{}](objc.ID(_NSTouchBarClass.class), objc.Sel("setAutomaticCustomizeTouchBarMenuItemEnabled:"), value)
}
// A template image for creating a new item.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbaraddtemplatename
func (_NSTouchBarClass NSTouchBarClass) TouchBarAddTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarClass.class), objc.Sel("NSImageNameTouchBarAddTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for opening a new document or view in edit mode.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbarcomposetemplatename
func (_NSTouchBarClass NSTouchBarClass) TouchBarComposeTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarClass.class), objc.Sel("NSImageNameTouchBarComposeTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for returning to the previous screen or location.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbargobacktemplatename
func (_NSTouchBarClass NSTouchBarClass) TouchBarGoBackTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarClass.class), objc.Sel("NSImageNameTouchBarGoBackTemplate"))
	return foundation.NSStringFromID(rv).String()
}
// A template image for moving to the next screen or location.
//
// See: https://developer.apple.com/documentation/appkit/nsimage/touchbargoforwardtemplatename
func (_NSTouchBarClass NSTouchBarClass) TouchBarGoForwardTemplateName() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTouchBarClass.class), objc.Sel("NSImageNameTouchBarGoForwardTemplate"))
	return foundation.NSStringFromID(rv).String()
}

