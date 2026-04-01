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
//   - [NSPopoverTouchBarItem.CollapsedRepresentationImage]: The image displayed by the button for the default collapsed representation.
//   - [NSPopoverTouchBarItem.SetCollapsedRepresentationImage]
//   - [NSPopoverTouchBarItem.CollapsedRepresentationLabel]: The localized string displayed by the button for the default collapsed representation.
//   - [NSPopoverTouchBarItem.SetCollapsedRepresentationLabel]
//
// # Configuring the expanded popover
//
//   - [NSPopoverTouchBarItem.ShowsCloseButton]: A Boolean value that determines whether a close button should be shown on the popover bar.
//   - [NSPopoverTouchBarItem.SetShowsCloseButton]
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
//   - [INSPopoverTouchBarItem.CollapsedRepresentationImage]: The image displayed by the button for the default collapsed representation.
//   - [INSPopoverTouchBarItem.SetCollapsedRepresentationImage]
//   - [INSPopoverTouchBarItem.CollapsedRepresentationLabel]: The localized string displayed by the button for the default collapsed representation.
//   - [INSPopoverTouchBarItem.SetCollapsedRepresentationLabel]
//
// # Configuring the expanded popover
//
//   - [INSPopoverTouchBarItem.ShowsCloseButton]: A Boolean value that determines whether a close button should be shown on the popover bar.
//   - [INSPopoverTouchBarItem.SetShowsCloseButton]
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

	// The image displayed by the button for the default collapsed representation.
	CollapsedRepresentationImage() INSImage
	SetCollapsedRepresentationImage(value INSImage)
	// The localized string displayed by the button for the default collapsed representation.
	CollapsedRepresentationLabel() string
	SetCollapsedRepresentationLabel(value string)

	// Topic: Configuring the expanded popover

	// A Boolean value that determines whether a close button should be shown on the popover bar.
	ShowsCloseButton() bool
	SetShowsCloseButton(value bool)

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

// Returns a gesture recognizer, configured to invoke the [ShowPopover]
// method.
//
// # Discussion
//
// Use this method to create a gesture recognizer that you then attach to a
// custom [CollapsedRepresentation] view.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/makeStandardActivatePopoverGestureRecognizer()
func (p NSPopoverTouchBarItem) MakeStandardActivatePopoverGestureRecognizer() INSGestureRecognizer {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("makeStandardActivatePopoverGestureRecognizer"))
	return NSGestureRecognizerFromID(rv)
}

// The image displayed by the button for the default collapsed representation.
//
// # Discussion
//
// If the [CollapsedRepresentation] button has been replaced by a different
// view, this property may not have any effect.
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
// If the [CollapsedRepresentation] button has been replaced by a different
// view, this property may not have any effect.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem/collapsedRepresentationLabel
func (p NSPopoverTouchBarItem) CollapsedRepresentationLabel() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("collapsedRepresentationLabel"))
	return foundation.NSStringFromID(rv).String()
}
func (p NSPopoverTouchBarItem) SetCollapsedRepresentationLabel(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setCollapsedRepresentationLabel:"), objc.String(value))
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
