// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSButtonTouchBarItem] class.
var (
	_NSButtonTouchBarItemClass     NSButtonTouchBarItemClass
	_NSButtonTouchBarItemClassOnce sync.Once
)

func getNSButtonTouchBarItemClass() NSButtonTouchBarItemClass {
	_NSButtonTouchBarItemClassOnce.Do(func() {
		_NSButtonTouchBarItemClass = NSButtonTouchBarItemClass{class: objc.GetClass("NSButtonTouchBarItem")}
	})
	return _NSButtonTouchBarItemClass
}

// GetNSButtonTouchBarItemClass returns the class object for NSButtonTouchBarItem.
func GetNSButtonTouchBarItemClass() NSButtonTouchBarItemClass {
	return getNSButtonTouchBarItemClass()
}

type NSButtonTouchBarItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSButtonTouchBarItemClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSButtonTouchBarItemClass) Alloc() NSButtonTouchBarItem {
	rv := objc.Send[NSButtonTouchBarItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A bar item that provides a button.
//
// # Configuring button appearance
//
//   - [NSButtonTouchBarItem.Title]
//   - [NSButtonTouchBarItem.SetTitle]
//   - [NSButtonTouchBarItem.Image]
//   - [NSButtonTouchBarItem.SetImage]
//   - [NSButtonTouchBarItem.BezelColor]
//   - [NSButtonTouchBarItem.SetBezelColor]
//
// # Configuring button state
//
//   - [NSButtonTouchBarItem.Enabled]
//   - [NSButtonTouchBarItem.SetEnabled]
//
// # Handling button interaction
//
//   - [NSButtonTouchBarItem.Target]
//   - [NSButtonTouchBarItem.SetTarget]
//   - [NSButtonTouchBarItem.Action]
//   - [NSButtonTouchBarItem.SetAction]
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem
type NSButtonTouchBarItem struct {
	NSTouchBarItem
}

// NSButtonTouchBarItemFromID constructs a [NSButtonTouchBarItem] from an objc.ID.
//
// A bar item that provides a button.
func NSButtonTouchBarItemFromID(id objc.ID) NSButtonTouchBarItem {
	return NSButtonTouchBarItem{NSTouchBarItem: NSTouchBarItemFromID(id)}
}

// NOTE: NSButtonTouchBarItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSButtonTouchBarItem] class.
//
// # Configuring button appearance
//
//   - [INSButtonTouchBarItem.Title]
//   - [INSButtonTouchBarItem.SetTitle]
//   - [INSButtonTouchBarItem.Image]
//   - [INSButtonTouchBarItem.SetImage]
//   - [INSButtonTouchBarItem.BezelColor]
//   - [INSButtonTouchBarItem.SetBezelColor]
//
// # Configuring button state
//
//   - [INSButtonTouchBarItem.Enabled]
//   - [INSButtonTouchBarItem.SetEnabled]
//
// # Handling button interaction
//
//   - [INSButtonTouchBarItem.Target]
//   - [INSButtonTouchBarItem.SetTarget]
//   - [INSButtonTouchBarItem.Action]
//   - [INSButtonTouchBarItem.SetAction]
//
// See: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem
type INSButtonTouchBarItem interface {
	INSTouchBarItem

	// Topic: Configuring button appearance

	Title() string
	SetTitle(value string)
	Image() INSImage
	SetImage(value INSImage)
	BezelColor() INSColor
	SetBezelColor(value INSColor)

	// Topic: Configuring button state

	Enabled() bool
	SetEnabled(value bool)

	// Topic: Handling button interaction

	Target() objectivec.IObject
	SetTarget(value objectivec.IObject)
	Action() objc.SEL
	SetAction(value objc.SEL)
}

// Init initializes the instance.
func (b NSButtonTouchBarItem) Init() NSButtonTouchBarItem {
	rv := objc.Send[NSButtonTouchBarItem](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b NSButtonTouchBarItem) Autorelease() NSButtonTouchBarItem {
	rv := objc.Send[NSButtonTouchBarItem](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSButtonTouchBarItem creates a new NSButtonTouchBarItem instance.
func NewNSButtonTouchBarItem() NSButtonTouchBarItem {
	class := getNSButtonTouchBarItemClass()
	rv := objc.Send[NSButtonTouchBarItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes and returns a new item from a storyboard or nib file.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(coder:)
func NewButtonTouchBarItemWithCoder(coder foundation.INSCoder) NSButtonTouchBarItem {
	instance := getNSButtonTouchBarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSButtonTouchBarItemFromID(rv)
}

// Creates a new item with the specified identifier.
//
// # Discussion
//
// The designated initializer. The identifier must be globally unique for
// every item, except for space items.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(identifier:)
func NewButtonTouchBarItemWithIdentifier(identifier NSTouchBarItemIdentifier) NSButtonTouchBarItem {
	instance := getNSButtonTouchBarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifier:"), objc.String(string(identifier)))
	return NSButtonTouchBarItemFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/init(identifier:image:target:action:)
func NewButtonTouchBarItemWithIdentifierImageTargetAction(identifier NSTouchBarItemIdentifier, image INSImage, target objectivec.IObject, action objc.SEL) NSButtonTouchBarItem {
	rv := objc.Send[objc.ID](objc.ID(getNSButtonTouchBarItemClass().class), objc.Sel("buttonTouchBarItemWithIdentifier:image:target:action:"), objc.String(string(identifier)), image, target, action)
	return NSButtonTouchBarItemFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/init(identifier:title:image:target:action:)
func NewButtonTouchBarItemWithIdentifierTitleImageTargetAction(identifier NSTouchBarItemIdentifier, title string, image INSImage, target objectivec.IObject, action objc.SEL) NSButtonTouchBarItem {
	rv := objc.Send[objc.ID](objc.ID(getNSButtonTouchBarItemClass().class), objc.Sel("buttonTouchBarItemWithIdentifier:title:image:target:action:"), objc.String(string(identifier)), objc.String(title), image, target, action)
	return NSButtonTouchBarItemFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/init(identifier:title:target:action:)
func NewButtonTouchBarItemWithIdentifierTitleTargetAction(identifier NSTouchBarItemIdentifier, title string, target objectivec.IObject, action objc.SEL) NSButtonTouchBarItem {
	rv := objc.Send[objc.ID](objc.ID(getNSButtonTouchBarItemClass().class), objc.Sel("buttonTouchBarItemWithIdentifier:title:target:action:"), objc.String(string(identifier)), objc.String(title), target, action)
	return NSButtonTouchBarItemFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/title
func (b NSButtonTouchBarItem) Title() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}
func (b NSButtonTouchBarItem) SetTitle(value string) {
	objc.Send[struct{}](b.ID, objc.Sel("setTitle:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/image
func (b NSButtonTouchBarItem) Image() INSImage {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("image"))
	return NSImageFromID(objc.ID(rv))
}
func (b NSButtonTouchBarItem) SetImage(value INSImage) {
	objc.Send[struct{}](b.ID, objc.Sel("setImage:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/bezelColor
func (b NSButtonTouchBarItem) BezelColor() INSColor {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("bezelColor"))
	return NSColorFromID(objc.ID(rv))
}
func (b NSButtonTouchBarItem) SetBezelColor(value INSColor) {
	objc.Send[struct{}](b.ID, objc.Sel("setBezelColor:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/isEnabled
func (b NSButtonTouchBarItem) Enabled() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isEnabled"))
	return rv
}
func (b NSButtonTouchBarItem) SetEnabled(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setEnabled:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/target
func (b NSButtonTouchBarItem) Target() objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("target"))
	return objectivec.Object{ID: rv}
}
func (b NSButtonTouchBarItem) SetTarget(value objectivec.IObject) {
	objc.Send[struct{}](b.ID, objc.Sel("setTarget:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/action
func (b NSButtonTouchBarItem) Action() objc.SEL {
	rv := objc.Send[objc.SEL](b.ID, objc.Sel("action"))
	return rv
}
func (b NSButtonTouchBarItem) SetAction(value objc.SEL) {
	objc.Send[struct{}](b.ID, objc.Sel("setAction:"), value)
}
