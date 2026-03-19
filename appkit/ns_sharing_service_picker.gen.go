// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSharingServicePicker] class.
var (
	_NSSharingServicePickerClass     NSSharingServicePickerClass
	_NSSharingServicePickerClassOnce sync.Once
)

func getNSSharingServicePickerClass() NSSharingServicePickerClass {
	_NSSharingServicePickerClassOnce.Do(func() {
		_NSSharingServicePickerClass = NSSharingServicePickerClass{class: objc.GetClass("NSSharingServicePicker")}
	})
	return _NSSharingServicePickerClass
}

// GetNSSharingServicePickerClass returns the class object for NSSharingServicePicker.
func GetNSSharingServicePickerClass() NSSharingServicePickerClass {
	return getNSSharingServicePickerClass()
}

type NSSharingServicePickerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSharingServicePickerClass) Alloc() NSSharingServicePicker {
	rv := objc.Send[NSSharingServicePicker](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A list of sharing services that the user can choose from.
//
// # Overview
// 
// An [NSSharingServicePicker] object presents an interface for sharing one or
// more items using a specific service. In macOS 12 and earlier, this picker
// displays a menu with a list of services that someone can use to share the
// item. In macOS 13 and later, the picker displays a popover with a preview
// of the item and the list of services. When someone chooses a service, the
// picker automatically shares the proposed item with that service.
// 
// Create a sharing service picker and configure it with a delegate object to
// monitor interactions. Your delegate must conform to the
// [NSSharingServicePickerDelegate] protocol. Present the picker from your
// interface using the [NSSharingServicePicker.ShowRelativeToRectOfViewPreferredEdge] method.
//
// # Creating a sharing service picker
//
//   - [NSSharingServicePicker.InitWithItems]: Creates a new sharing service picker for the selected items.
//
// # Managing the sharing service picker
//
//   - [NSSharingServicePicker.Delegate]: The object for managing the sharing service picker.
//   - [NSSharingServicePicker.SetDelegate]
//
// # Displaying the sharing service picker
//
//   - [NSSharingServicePicker.ShowRelativeToRectOfViewPreferredEdge]: Shows the picker interface and populates it with the relevant sharing services.
//   - [NSSharingServicePicker.Close]: Closes the picker interface.
//
// # Retrieving the sharing menu item
//
//   - [NSSharingServicePicker.StandardShareMenuItem]: A menu item suitable to display the picker for the specified items.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker
type NSSharingServicePicker struct {
	objectivec.Object
}

// NSSharingServicePickerFromID constructs a [NSSharingServicePicker] from an objc.ID.
//
// A list of sharing services that the user can choose from.
func NSSharingServicePickerFromID(id objc.ID) NSSharingServicePicker {
	return NSSharingServicePicker{objectivec.Object{ID: id}}
}
// NOTE: NSSharingServicePicker adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSharingServicePicker] class.
//
// # Creating a sharing service picker
//
//   - [INSSharingServicePicker.InitWithItems]: Creates a new sharing service picker for the selected items.
//
// # Managing the sharing service picker
//
//   - [INSSharingServicePicker.Delegate]: The object for managing the sharing service picker.
//   - [INSSharingServicePicker.SetDelegate]
//
// # Displaying the sharing service picker
//
//   - [INSSharingServicePicker.ShowRelativeToRectOfViewPreferredEdge]: Shows the picker interface and populates it with the relevant sharing services.
//   - [INSSharingServicePicker.Close]: Closes the picker interface.
//
// # Retrieving the sharing menu item
//
//   - [INSSharingServicePicker.StandardShareMenuItem]: A menu item suitable to display the picker for the specified items.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker
type INSSharingServicePicker interface {
	objectivec.IObject

	// Topic: Creating a sharing service picker

	// Creates a new sharing service picker for the selected items.
	InitWithItems(items foundation.INSArray) NSSharingServicePicker

	// Topic: Managing the sharing service picker

	// The object for managing the sharing service picker.
	Delegate() NSSharingServicePickerDelegate
	SetDelegate(value NSSharingServicePickerDelegate)

	// Topic: Displaying the sharing service picker

	// Shows the picker interface and populates it with the relevant sharing services.
	ShowRelativeToRectOfViewPreferredEdge(rect corefoundation.CGRect, view INSView, preferredEdge foundation.NSRectEdge)
	// Closes the picker interface.
	Close()

	// Topic: Retrieving the sharing menu item

	// A menu item suitable to display the picker for the specified items.
	StandardShareMenuItem() INSMenuItem
}

// Init initializes the instance.
func (s NSSharingServicePicker) Init() NSSharingServicePicker {
	rv := objc.Send[NSSharingServicePicker](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSharingServicePicker) Autorelease() NSSharingServicePicker {
	rv := objc.Send[NSSharingServicePicker](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSharingServicePicker creates a new NSSharingServicePicker instance.
func NewNSSharingServicePicker() NSSharingServicePicker {
	class := getNSSharingServicePickerClass()
	rv := objc.Send[NSSharingServicePicker](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new sharing service picker for the selected items.
//
// items: The items to be shared. The items in the array must conform to the
// [NSPasteboardWriting] or [NSPreviewRepresentableActivityItem] protocol. For
// example, you can specify an [NSString], [NSImage], [NSURL], or similar type
// directly. You can also specify [NSItemProvider] or [NSDocument] objects in
// the array to share those types.
// //
// [NSItemProvider]: https://developer.apple.com/documentation/Foundation/NSItemProvider
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [NSURL]: https://developer.apple.com/documentation/Foundation/NSURL
//
// # Return Value
// 
// A configured sharing service picker.
//
// # Discussion
// 
// If an item is an [NSURL] object and contains a file URL (pointing to a
// video on the local disk for example), the picker shares the content of the
// file. If the URL is remote, then the picker shares the URL instead of the
// contents.
//
// [NSURL]: https://developer.apple.com/documentation/Foundation/NSURL
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/init(items:)
func NewSharingServicePickerWithItems(items foundation.INSArray) NSSharingServicePicker {
	instance := getNSSharingServicePickerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithItems:"), items)
	return NSSharingServicePickerFromID(rv)
}

// Creates a new sharing service picker for the selected items.
//
// items: The items to be shared. The items in the array must conform to the
// [NSPasteboardWriting] or [NSPreviewRepresentableActivityItem] protocol. For
// example, you can specify an [NSString], [NSImage], [NSURL], or similar type
// directly. You can also specify [NSItemProvider] or [NSDocument] objects in
// the array to share those types.
// //
// [NSItemProvider]: https://developer.apple.com/documentation/Foundation/NSItemProvider
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [NSURL]: https://developer.apple.com/documentation/Foundation/NSURL
//
// # Return Value
// 
// A configured sharing service picker.
//
// # Discussion
// 
// If an item is an [NSURL] object and contains a file URL (pointing to a
// video on the local disk for example), the picker shares the content of the
// file. If the URL is remote, then the picker shares the URL instead of the
// contents.
//
// [NSURL]: https://developer.apple.com/documentation/Foundation/NSURL
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/init(items:)
func (s NSSharingServicePicker) InitWithItems(items foundation.INSArray) NSSharingServicePicker {
	rv := objc.Send[NSSharingServicePicker](s.ID, objc.Sel("initWithItems:"), items)
	return rv
}
// Shows the picker interface and populates it with the relevant sharing
// services.
//
// rect: The rectangle the picker should be showed relative to. The coordinates are
// in the `view` coordinate system. Passing [NSZeroRect] causes the view
// bounds to be used.
// //
// [NSZeroRect]: https://developer.apple.com/documentation/Foundation/NSZeroRect
//
// view: The view.
//
// preferredEdge: The preferred edge of the view to display the picker. See [NSRectEdge] for
// the possible values.
// //
// [NSRectEdge]: https://developer.apple.com/documentation/Foundation/NSRectEdge
//
// # Discussion
// 
// In macOS 13 and later, the picker displays the macOS share sheet in a
// popover. In earlier versions of macOS, the picker displays a menu with the
// available services. When someone chooses a service for sharing the items,
// the picker notifies its delegate and then shares the content.
// 
// The following example shows the basic code to display the picker in
// response to an action. Update the code to include the items you want to
// share and to position the picker interface in an appropriate part of your
// window.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/show(relativeTo:of:preferredEdge:)
func (s NSSharingServicePicker) ShowRelativeToRectOfViewPreferredEdge(rect corefoundation.CGRect, view INSView, preferredEdge foundation.NSRectEdge) {
	objc.Send[objc.ID](s.ID, objc.Sel("showRelativeToRect:ofView:preferredEdge:"), rect, view, preferredEdge)
}
// Closes the picker interface.
//
// # Discussion
// 
// The [SharingServicePickerDidChooseSharingService] method will be invoked if
// the [Delegate] is set to `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/close()
func (s NSSharingServicePicker) Close() {
	objc.Send[objc.ID](s.ID, objc.Sel("close"))
}

// The object for managing the sharing service picker.
//
// # Discussion
// 
// The delegate object must conform to the [NSSharingServicePickerDelegate]
// delegate.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/delegate
func (s NSSharingServicePicker) Delegate() NSSharingServicePickerDelegate {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("delegate"))
	return NSSharingServicePickerDelegateObjectFromID(rv)
}
func (s NSSharingServicePicker) SetDelegate(value NSSharingServicePickerDelegate) {
	objc.Send[struct{}](s.ID, objc.Sel("setDelegate:"), value)
}
// A menu item suitable to display the picker for the specified items.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/standardShareMenuItem
func (s NSSharingServicePicker) StandardShareMenuItem() INSMenuItem {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("standardShareMenuItem"))
	return NSMenuItemFromID(objc.ID(rv))
}

