// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSharingServicePickerTouchBarItem] class.
var (
	_NSSharingServicePickerTouchBarItemClass     NSSharingServicePickerTouchBarItemClass
	_NSSharingServicePickerTouchBarItemClassOnce sync.Once
)

func getNSSharingServicePickerTouchBarItemClass() NSSharingServicePickerTouchBarItemClass {
	_NSSharingServicePickerTouchBarItemClassOnce.Do(func() {
		_NSSharingServicePickerTouchBarItemClass = NSSharingServicePickerTouchBarItemClass{class: objc.GetClass("NSSharingServicePickerTouchBarItem")}
	})
	return _NSSharingServicePickerTouchBarItemClass
}

// GetNSSharingServicePickerTouchBarItemClass returns the class object for NSSharingServicePickerTouchBarItem.
func GetNSSharingServicePickerTouchBarItemClass() NSSharingServicePickerTouchBarItemClass {
	return getNSSharingServicePickerTouchBarItemClass()
}

type NSSharingServicePickerTouchBarItemClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSharingServicePickerTouchBarItemClass) Alloc() NSSharingServicePickerTouchBarItem {
	rv := objc.Send[NSSharingServicePickerTouchBarItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A bar item that, along with its delegate, provides a list of objects
// eligible for sharing.
//
// # Setting the delegate
//
//   - [NSSharingServicePickerTouchBarItem.Delegate]: The object that acts as the delegate of the sharing service picker bar item.
//   - [NSSharingServicePickerTouchBarItem.SetDelegate]
//
// # Configuring the appearance
//
//   - [NSSharingServicePickerTouchBarItem.ButtonImage]: The image displayed in the sharing service picker item button.
//   - [NSSharingServicePickerTouchBarItem.SetButtonImage]
//   - [NSSharingServicePickerTouchBarItem.ButtonTitle]: The text displayed in the sharing service picker item button.
//   - [NSSharingServicePickerTouchBarItem.SetButtonTitle]
//
// # Enabling the item
//
//   - [NSSharingServicePickerTouchBarItem.Enabled]: A Boolean value that specifies whether the sharing service picker item is enabled.
//   - [NSSharingServicePickerTouchBarItem.SetEnabled]
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem
type NSSharingServicePickerTouchBarItem struct {
	NSTouchBarItem
}

// NSSharingServicePickerTouchBarItemFromID constructs a [NSSharingServicePickerTouchBarItem] from an objc.ID.
//
// A bar item that, along with its delegate, provides a list of objects
// eligible for sharing.
func NSSharingServicePickerTouchBarItemFromID(id objc.ID) NSSharingServicePickerTouchBarItem {
	return NSSharingServicePickerTouchBarItem{NSTouchBarItem: NSTouchBarItemFromID(id)}
}
// NOTE: NSSharingServicePickerTouchBarItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSSharingServicePickerTouchBarItem] class.
//
// # Setting the delegate
//
//   - [INSSharingServicePickerTouchBarItem.Delegate]: The object that acts as the delegate of the sharing service picker bar item.
//   - [INSSharingServicePickerTouchBarItem.SetDelegate]
//
// # Configuring the appearance
//
//   - [INSSharingServicePickerTouchBarItem.ButtonImage]: The image displayed in the sharing service picker item button.
//   - [INSSharingServicePickerTouchBarItem.SetButtonImage]
//   - [INSSharingServicePickerTouchBarItem.ButtonTitle]: The text displayed in the sharing service picker item button.
//   - [INSSharingServicePickerTouchBarItem.SetButtonTitle]
//
// # Enabling the item
//
//   - [INSSharingServicePickerTouchBarItem.Enabled]: A Boolean value that specifies whether the sharing service picker item is enabled.
//   - [INSSharingServicePickerTouchBarItem.SetEnabled]
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem
type INSSharingServicePickerTouchBarItem interface {
	INSTouchBarItem

	// Topic: Setting the delegate

	// The object that acts as the delegate of the sharing service picker bar item.
	Delegate() NSSharingServicePickerTouchBarItemDelegate
	SetDelegate(value NSSharingServicePickerTouchBarItemDelegate)

	// Topic: Configuring the appearance

	// The image displayed in the sharing service picker item button.
	ButtonImage() objectivec.Object
	SetButtonImage(value objectivec.Object)
	// The text displayed in the sharing service picker item button.
	ButtonTitle() string
	SetButtonTitle(value string)

	// Topic: Enabling the item

	// A Boolean value that specifies whether the sharing service picker item is enabled.
	Enabled() bool
	SetEnabled(value bool)

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (s NSSharingServicePickerTouchBarItem) Init() NSSharingServicePickerTouchBarItem {
	rv := objc.Send[NSSharingServicePickerTouchBarItem](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSharingServicePickerTouchBarItem) Autorelease() NSSharingServicePickerTouchBarItem {
	rv := objc.Send[NSSharingServicePickerTouchBarItem](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSharingServicePickerTouchBarItem creates a new NSSharingServicePickerTouchBarItem instance.
func NewNSSharingServicePickerTouchBarItem() NSSharingServicePickerTouchBarItem {
	class := getNSSharingServicePickerTouchBarItemClass()
	rv := objc.Send[NSSharingServicePickerTouchBarItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Initializes and returns a new item from a storyboard or nib file.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(coder:)
func NewSharingServicePickerTouchBarItemWithCoder(coder foundation.INSCoder) NSSharingServicePickerTouchBarItem {
	instance := getNSSharingServicePickerTouchBarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSSharingServicePickerTouchBarItemFromID(rv)
}


// Creates a new item with the specified identifier.
//
// # Discussion
// 
// The designated initializer. The identifier must be globally unique for
// every item, except for space items.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(identifier:)
func NewSharingServicePickerTouchBarItemWithIdentifier(identifier NSTouchBarItemIdentifier) NSSharingServicePickerTouchBarItem {
	instance := getNSSharingServicePickerTouchBarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifier:"), objc.String(string(identifier)))
	return NSSharingServicePickerTouchBarItemFromID(rv)
}






func (s NSSharingServicePickerTouchBarItem) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The object that acts as the delegate of the sharing service picker bar
// item.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem/delegate
func (s NSSharingServicePickerTouchBarItem) Delegate() NSSharingServicePickerTouchBarItemDelegate {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("delegate"))
	return NSSharingServicePickerTouchBarItemDelegateObjectFromID(rv)
}
func (s NSSharingServicePickerTouchBarItem) SetDelegate(value NSSharingServicePickerTouchBarItemDelegate) {
	objc.Send[struct{}](s.ID, objc.Sel("setDelegate:"), value)
}



// The image displayed in the sharing service picker item button.
//
// # Discussion
// 
// The default value of this property is a sharing service picker image.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem/buttonImage
func (s NSSharingServicePickerTouchBarItem) ButtonImage() objectivec.Object {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("buttonImage"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (s NSSharingServicePickerTouchBarItem) SetButtonImage(value objectivec.Object) {
	objc.Send[struct{}](s.ID, objc.Sel("setButtonImage:"), value)
}



// The text displayed in the sharing service picker item button.
//
// # Discussion
// 
// The default value of this property is the empty string.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem/buttonTitle
func (s NSSharingServicePickerTouchBarItem) ButtonTitle() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("buttonTitle"))
	return foundation.NSStringFromID(rv).String()
}
func (s NSSharingServicePickerTouchBarItem) SetButtonTitle(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setButtonTitle:"), objc.String(value))
}



// A Boolean value that specifies whether the sharing service picker item is
// enabled.
//
// # Discussion
// 
// If [true], the sharing button is enabled.
// 
// If the sharing popover is currently visible when this property is changed
// to [false], the popover is dismissed.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem/isEnabled
func (s NSSharingServicePickerTouchBarItem) Enabled() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isEnabled"))
	return rv
}
func (s NSSharingServicePickerTouchBarItem) SetEnabled(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setEnabled:"), value)
}

























