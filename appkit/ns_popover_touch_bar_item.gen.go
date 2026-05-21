// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPopoverTouchBarItem] class.
var (
	_NSPopoverTouchBarItemClass     NSPopoverTouchBarItemClass
	_NSPopoverTouchBarItemClassOnce sync.Once
)

func getNSPopoverTouchBarItemClass() NSPopoverTouchBarItemClass {
	_NSPopoverTouchBarItemClassOnce.Do(func() {
		_NSPopoverTouchBarItemClass = NSPopoverTouchBarItemClass{class: objc.GetClass("NSPopoverTouchBarItem")}
	})
	return _NSPopoverTouchBarItemClass
}

// GetNSPopoverTouchBarItemClass returns the class object for NSPopoverTouchBarItem.
func GetNSPopoverTouchBarItemClass() NSPopoverTouchBarItemClass {
	return getNSPopoverTouchBarItemClass()
}

type NSPopoverTouchBarItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPopoverTouchBarItemClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPopoverTouchBarItemClass) Alloc() NSPopoverTouchBarItem {
	rv := objc.Send[NSPopoverTouchBarItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A bar item that provides a two-state control that can expand into its
// second state, showing the contents of a bar that it owns.
//
// # Configuring the collapsed popover
//
//   - [NSPopoverTouchBarItem.CollapsedRepresentation]: The view displayed when this item is displayed in its parent bar.
//   - [NSPopoverTouchBarItem.SetCollapsedRepresentation]
//   - [NSPopoverTouchBarItem.CollapsedRepresentationImage]: The image displayed by the button for the default collapsed representation.
//   - [NSPopoverTouchBarItem.SetCollapsedRepresentationImage]
//   - [NSPopoverTouchBarItem.CollapsedRepresentationLabel]: The localized string displayed by the button for the default collapsed representation.
//   - [NSPopoverTouchBarItem.SetCollapsedRepresentationLabel]
//
// # Configuring the expanded popover
//
//   - [NSPopoverTouchBarItem.PopoverTouchBar]: The bar displayed when this item is “popped.”
//   - [NSPopoverTouchBarItem.SetPopoverTouchBar]
//   - [NSPopoverTouchBarItem.ShowsCloseButton]: A Boolean value that determines whether a close button should be shown on the popover bar.
//   - [NSPopoverTouchBarItem.SetShowsCloseButton]
//   - [NSPopoverTouchBarItem.PressAndHoldTouchBar]: The bar that is displayed when a user press-and-holds on the popover item.
//   - [NSPopoverTouchBarItem.SetPressAndHoldTouchBar]
//
// # Expanding and collapsing a popover
//
//   - [NSPopoverTouchBarItem.ShowPopover]: Replaces the main bar with this item’s popover bar.
//   - [NSPopoverTouchBarItem.DismissPopover]: Restores the previously visible main bar.
//   - [NSPopoverTouchBarItem.MakeStandardActivatePopoverGestureRecognizer]: Returns a gesture recognizer, configured to invoke the [showPopover(_:)](<doc://com.apple.appkit/documentation/AppKit/NSPopoverTouchBarItem/showPopover(_:)>) method.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem
type NSPopoverTouchBarItem struct {
	NSTouchBarItem
}

// NSPopoverTouchBarItemFromID constructs a [NSPopoverTouchBarItem] from an objc.ID.
//
// A bar item that provides a two-state control that can expand into its
// second state, showing the contents of a bar that it owns.
func NSPopoverTouchBarItemFromID(id objc.ID) NSPopoverTouchBarItem {
	return NSPopoverTouchBarItem{NSTouchBarItem: NSTouchBarItemFromID(id)}
}

// NOTE: NSPopoverTouchBarItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPopoverTouchBarItem] class.
//
// # Configuring the collapsed popover
//
//   - [INSPopoverTouchBarItem.CollapsedRepresentation]: The view displayed when this item is displayed in its parent bar.
//   - [INSPopoverTouchBarItem.SetCollapsedRepresentation]
//   - [INSPopoverTouchBarItem.CollapsedRepresentationImage]: The image displayed by the button for the default collapsed representation.
//   - [INSPopoverTouchBarItem.SetCollapsedRepresentationImage]
//   - [INSPopoverTouchBarItem.CollapsedRepresentationLabel]: The localized string displayed by the button for the default collapsed representation.
//   - [INSPopoverTouchBarItem.SetCollapsedRepresentationLabel]
//
// # Configuring the expanded popover
//
//   - [INSPopoverTouchBarItem.PopoverTouchBar]: The bar displayed when this item is “popped.”
//   - [INSPopoverTouchBarItem.SetPopoverTouchBar]
//   - [INSPopoverTouchBarItem.ShowsCloseButton]: A Boolean value that determines whether a close button should be shown on the popover bar.
//   - [INSPopoverTouchBarItem.SetShowsCloseButton]
//   - [INSPopoverTouchBarItem.PressAndHoldTouchBar]: The bar that is displayed when a user press-and-holds on the popover item.
//   - [INSPopoverTouchBarItem.SetPressAndHoldTouchBar]
//
// # Expanding and collapsing a popover
//
//   - [INSPopoverTouchBarItem.ShowPopover]: Replaces the main bar with this item’s popover bar.
//   - [INSPopoverTouchBarItem.DismissPopover]: Restores the previously visible main bar.
//   - [INSPopoverTouchBarItem.MakeStandardActivatePopoverGestureRecognizer]: Returns a gesture recognizer, configured to invoke the [showPopover(_:)](<doc://com.apple.appkit/documentation/AppKit/NSPopoverTouchBarItem/showPopover(_:)>) method.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem
type INSPopoverTouchBarItem interface {
	INSTouchBarItem

	// Topic: Configuring the collapsed popover

	// The view displayed when this item is displayed in its parent bar.
	CollapsedRepresentation() INSView
	SetCollapsedRepresentation(value INSView)
	// The image displayed by the button for the default collapsed representation.
	CollapsedRepresentationImage() INSImage
	SetCollapsedRepresentationImage(value INSImage)
	// The localized string displayed by the button for the default collapsed representation.
	CollapsedRepresentationLabel() string
	SetCollapsedRepresentationLabel(value string)

	// Topic: Configuring the expanded popover

	// The bar displayed when this item is “popped.”
	PopoverTouchBar() INSTouchBar
	SetPopoverTouchBar(value INSTouchBar)
	// A Boolean value that determines whether a close button should be shown on the popover bar.
	ShowsCloseButton() bool
	SetShowsCloseButton(value bool)
	// The bar that is displayed when a user press-and-holds on the popover item.
	PressAndHoldTouchBar() INSTouchBar
	SetPressAndHoldTouchBar(value INSTouchBar)

	// Topic: Expanding and collapsing a popover

	// Replaces the main bar with this item’s popover bar.
	ShowPopover(sender objectivec.IObject)
	// Restores the previously visible main bar.
	DismissPopover(sender objectivec.IObject)
	// Returns a gesture recognizer, configured to invoke the [showPopover(_:)](<doc://com.apple.appkit/documentation/AppKit/NSPopoverTouchBarItem/showPopover(_:)>) method.
	MakeStandardActivatePopoverGestureRecognizer() INSGestureRecognizer
}

// Init initializes the instance.
func (p NSPopoverTouchBarItem) Init() NSPopoverTouchBarItem {
	rv := objc.Send[NSPopoverTouchBarItem](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPopoverTouchBarItem) Autorelease() NSPopoverTouchBarItem {
	rv := objc.Send[NSPopoverTouchBarItem](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPopoverTouchBarItem creates a new NSPopoverTouchBarItem instance.
func NewNSPopoverTouchBarItem() NSPopoverTouchBarItem {
	class := getNSPopoverTouchBarItemClass()
	rv := objc.Send[NSPopoverTouchBarItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes and returns a new item from a storyboard or nib file.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(coder:)
func NewPopoverTouchBarItemWithCoder(coder foundation.INSCoder) NSPopoverTouchBarItem {
	instance := getNSPopoverTouchBarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSPopoverTouchBarItemFromID(rv)
}

// Creates a new item with the specified identifier.
//
// # Discussion
//
// The designated initializer. The identifier must be globally unique for
// every item, except for space items.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(identifier:)
func NewPopoverTouchBarItemWithIdentifier(identifier NSTouchBarItemIdentifier) NSPopoverTouchBarItem {
	instance := getNSPopoverTouchBarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifier:"), objc.String(string(identifier)))
	return NSPopoverTouchBarItemFromID(rv)
}

// Replaces the main bar with this item’s popover bar.
//
// # Discussion
//
// If this item is not visible, this method will have no effect.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/showPopover(_:)
func (p NSPopoverTouchBarItem) ShowPopover(sender objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("showPopover:"), sender)
}

// Restores the previously visible main bar.
//
// # Discussion
//
// This method has the same effect as the user tapping the optional close
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/dismissPopover(_:)
func (p NSPopoverTouchBarItem) DismissPopover(sender objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("dismissPopover:"), sender)
}

// Returns a gesture recognizer, configured to invoke the
// [NSPopoverTouchBarItem.ShowPopover] method.
//
// # Discussion
//
// Use this method to create a gesture recognizer that you then attach to a
// custom [NSPopoverTouchBarItem.CollapsedRepresentation] view.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/makeStandardActivatePopoverGestureRecognizer()
func (p NSPopoverTouchBarItem) MakeStandardActivatePopoverGestureRecognizer() INSGestureRecognizer {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("makeStandardActivatePopoverGestureRecognizer"))
	return NSGestureRecognizerFromID(rv)
}

// The view displayed when this item is displayed in its parent bar.
//
// # Discussion
//
// By default, this is an [NSButton] whose target is this popover item, whose
// action is [NSPopoverTouchBarItem.ShowPopover], and whose image and title
// are bound to this item’s
// [NSPopoverTouchBarItem.CollapsedRepresentationImage] and
// [NSPopoverTouchBarItem.CollapsedRepresentationImage] respectively.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/collapsedRepresentation
func (p NSPopoverTouchBarItem) CollapsedRepresentation() INSView {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("collapsedRepresentation"))
	return NSViewFromID(objc.ID(rv))
}
func (p NSPopoverTouchBarItem) SetCollapsedRepresentation(value INSView) {
	objc.Send[struct{}](p.ID, objc.Sel("setCollapsedRepresentation:"), value)
}

// The image displayed by the button for the default collapsed representation.
//
// # Discussion
//
// If the [NSPopoverTouchBarItem.CollapsedRepresentation] button has been
// replaced by a different view, this property may not have any effect.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/collapsedRepresentationImage
func (p NSPopoverTouchBarItem) CollapsedRepresentationImage() INSImage {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("collapsedRepresentationImage"))
	return NSImageFromID(objc.ID(rv))
}
func (p NSPopoverTouchBarItem) SetCollapsedRepresentationImage(value INSImage) {
	objc.Send[struct{}](p.ID, objc.Sel("setCollapsedRepresentationImage:"), value)
}

// The localized string displayed by the button for the default collapsed
// representation.
//
// # Discussion
//
// If the [NSPopoverTouchBarItem.CollapsedRepresentation] button has been
// replaced by a different view, this property may not have any effect.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/collapsedRepresentationLabel
func (p NSPopoverTouchBarItem) CollapsedRepresentationLabel() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("collapsedRepresentationLabel"))
	return foundation.NSStringFromID(rv).String()
}
func (p NSPopoverTouchBarItem) SetCollapsedRepresentationLabel(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setCollapsedRepresentationLabel:"), objc.String(value))
}

// The bar displayed when this item is “popped.”
//
// # Discussion
//
// Set this property to a fully configured instance of [NSTouchBar] that is
// displayed when the user taps on the popover item. By default this is an
// empty bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/popoverTouchBar
func (p NSPopoverTouchBarItem) PopoverTouchBar() INSTouchBar {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("popoverTouchBar"))
	return NSTouchBarFromID(objc.ID(rv))
}
func (p NSPopoverTouchBarItem) SetPopoverTouchBar(value INSTouchBar) {
	objc.Send[struct{}](p.ID, objc.Sel("setPopoverTouchBar:"), value)
}

// A Boolean value that determines whether a close button should be shown on
// the popover bar.
//
// # Discussion
//
// When true, a close button is automatically displayed when the popover bar
// is displayed. When false, it is your responsibility to dismiss the popover
// bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/showsCloseButton
func (p NSPopoverTouchBarItem) ShowsCloseButton() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("showsCloseButton"))
	return rv
}
func (p NSPopoverTouchBarItem) SetShowsCloseButton(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setShowsCloseButton:"), value)
}

// The bar that is displayed when a user press-and-holds on the popover item.
//
// # Discussion
//
// This [NSTouchBar] can be the same as the one used for the
// [NSPopoverTouchBarItem.PopoverTouchBar] property, but does not have to be.
//
// When non-`nil` this touch bar is displayed while the user holds their
// finger down on the collapsed representation of the popover item. When the
// user raises their finger, this bar disappears.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/pressAndHoldTouchBar
func (p NSPopoverTouchBarItem) PressAndHoldTouchBar() INSTouchBar {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("pressAndHoldTouchBar"))
	return NSTouchBarFromID(objc.ID(rv))
}
func (p NSPopoverTouchBarItem) SetPressAndHoldTouchBar(value INSTouchBar) {
	objc.Send[struct{}](p.ID, objc.Sel("setPressAndHoldTouchBar:"), value)
}
