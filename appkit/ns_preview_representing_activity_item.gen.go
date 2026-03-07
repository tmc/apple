// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPreviewRepresentingActivityItem] class.
var (
	_NSPreviewRepresentingActivityItemClass     NSPreviewRepresentingActivityItemClass
	_NSPreviewRepresentingActivityItemClassOnce sync.Once
)

func getNSPreviewRepresentingActivityItemClass() NSPreviewRepresentingActivityItemClass {
	_NSPreviewRepresentingActivityItemClassOnce.Do(func() {
		_NSPreviewRepresentingActivityItemClass = NSPreviewRepresentingActivityItemClass{class: objc.GetClass("NSPreviewRepresentingActivityItem")}
	})
	return _NSPreviewRepresentingActivityItemClass
}

// GetNSPreviewRepresentingActivityItemClass returns the class object for NSPreviewRepresentingActivityItem.
func GetNSPreviewRepresentingActivityItemClass() NSPreviewRepresentingActivityItemClass {
	return getNSPreviewRepresentingActivityItemClass()
}

type NSPreviewRepresentingActivityItemClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPreviewRepresentingActivityItemClass) Alloc() NSPreviewRepresentingActivityItem {
	rv := objc.Send[NSPreviewRepresentingActivityItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A type that adds metadata to an item you share using the macOS share sheet.
//
// # Overview
// 
// An [NSPreviewRepresentingActivityItem] object provides a concrete
// implementation of the [NSPreviewRepresentableActivityItem] protocol. Use it
// to create shareable items for common types such as strings or images, or
// when you don’t want to adopt the [NSPreviewRepresentableActivityItem]
// protocol directly in your app’s objects. To share the item from your app,
// initialize the [NSSharingServicePicker] object with this object.
//
// # Creating a Preview Activity Item
//
//   - [NSPreviewRepresentingActivityItem.InitWithItemTitleImageIcon]: Creates a metadata object with the title, image, and icon for a shareable item.
//   - [NSPreviewRepresentingActivityItem.InitWithItemTitleImageProviderIconProvider]: Creates a metadata object that provides a title and images for a shareable item.
//
// See: https://developer.apple.com/documentation/AppKit/NSPreviewRepresentingActivityItem
type NSPreviewRepresentingActivityItem struct {
	objectivec.Object
}

// NSPreviewRepresentingActivityItemFromID constructs a [NSPreviewRepresentingActivityItem] from an objc.ID.
//
// A type that adds metadata to an item you share using the macOS share sheet.
func NSPreviewRepresentingActivityItemFromID(id objc.ID) NSPreviewRepresentingActivityItem {
	return NSPreviewRepresentingActivityItem{objectivec.Object{id}}
}
// NOTE: NSPreviewRepresentingActivityItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSPreviewRepresentingActivityItem] class.
//
// # Creating a Preview Activity Item
//
//   - [INSPreviewRepresentingActivityItem.InitWithItemTitleImageIcon]: Creates a metadata object with the title, image, and icon for a shareable item.
//   - [INSPreviewRepresentingActivityItem.InitWithItemTitleImageProviderIconProvider]: Creates a metadata object that provides a title and images for a shareable item.
//
// See: https://developer.apple.com/documentation/AppKit/NSPreviewRepresentingActivityItem
type INSPreviewRepresentingActivityItem interface {
	objectivec.IObject
	NSPreviewRepresentableActivityItem

	// Topic: Creating a Preview Activity Item

	// Creates a metadata object with the title, image, and icon for a shareable item.
	InitWithItemTitleImageIcon(item objectivec.IObject, title string, image INSImage, icon INSImage) NSPreviewRepresentingActivityItem
	// Creates a metadata object that provides a title and images for a shareable item.
	InitWithItemTitleImageProviderIconProvider(item objectivec.IObject, title string, imageProvider foundation.NSItemProvider, iconProvider foundation.NSItemProvider) NSPreviewRepresentingActivityItem
}





// Init initializes the instance.
func (p NSPreviewRepresentingActivityItem) Init() NSPreviewRepresentingActivityItem {
	rv := objc.Send[NSPreviewRepresentingActivityItem](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPreviewRepresentingActivityItem) Autorelease() NSPreviewRepresentingActivityItem {
	rv := objc.Send[NSPreviewRepresentingActivityItem](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPreviewRepresentingActivityItem creates a new NSPreviewRepresentingActivityItem instance.
func NewNSPreviewRepresentingActivityItem() NSPreviewRepresentingActivityItem {
	class := getNSPreviewRepresentingActivityItemClass()
	rv := objc.Send[NSPreviewRepresentingActivityItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a metadata object with the title, image, and icon for a shareable
// item.
//
// item: The item to share from the Mac share sheet. The item must conform to the
// [NSPasteboardWriting] protocol, or be an [NSItemProvider] or [NSDocument]
// object.
// //
// [NSItemProvider]: https://developer.apple.com/documentation/Foundation/NSItemProvider
//
// title: The localized name of the item.
//
// image: An image to display as a preview for the item. For example, when you share
// a URL for a webpage, you might specify the hero image for that page or a
// rendering of the webpage itself.
//
// icon: A thumbnail-size image of the source of the item. Typically, you specify
// your app’s icon but you can provide a different icon if the content has a
// different source.
//
// # Return Value
// 
// An initialized item to share.
//
// See: https://developer.apple.com/documentation/AppKit/NSPreviewRepresentingActivityItem/init(item:title:image:icon:)
func NewPreviewRepresentingActivityItemWithItemTitleImageIcon(item objectivec.IObject, title string, image INSImage, icon INSImage) NSPreviewRepresentingActivityItem {
	instance := getNSPreviewRepresentingActivityItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithItem:title:image:icon:"), item, objc.String(title), image, icon)
	return NSPreviewRepresentingActivityItemFromID(rv)
}


// Creates a metadata object that provides a title and images for a shareable
// item.
//
// item: The item to share from the Mac share sheet. The item must conform to the
// [NSPasteboardWriting] protocol, or be an [NSItemProvider] or [NSDocument]
// object.
// //
// [NSItemProvider]: https://developer.apple.com/documentation/Foundation/NSItemProvider
//
// title: The localized name of the item.
//
// imageProvider: An object that provides the image to display as a preview for the item. For
// example, when you share a URL for a webpage, you might specify the hero
// image for that page or a rendering of the webpage itself.
//
// iconProvider: An object that a thumbnail-size image of the source of the item. Typically,
// you specify your app’s icon but you can provide a different icon if the
// content has a different source.
//
// # Return Value
// 
// An initialized item to share.
//
// See: https://developer.apple.com/documentation/AppKit/NSPreviewRepresentingActivityItem/init(item:title:imageProvider:iconProvider:)
func NewPreviewRepresentingActivityItemWithItemTitleImageProviderIconProvider(item objectivec.IObject, title string, imageProvider foundation.NSItemProvider, iconProvider foundation.NSItemProvider) NSPreviewRepresentingActivityItem {
	instance := getNSPreviewRepresentingActivityItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithItem:title:imageProvider:iconProvider:"), item, objc.String(title), imageProvider, iconProvider)
	return NSPreviewRepresentingActivityItemFromID(rv)
}







// Creates a metadata object with the title, image, and icon for a shareable
// item.
//
// item: The item to share from the Mac share sheet. The item must conform to the
// [NSPasteboardWriting] protocol, or be an [NSItemProvider] or [NSDocument]
// object.
// //
// [NSItemProvider]: https://developer.apple.com/documentation/Foundation/NSItemProvider
//
// title: The localized name of the item.
//
// image: An image to display as a preview for the item. For example, when you share
// a URL for a webpage, you might specify the hero image for that page or a
// rendering of the webpage itself.
//
// icon: A thumbnail-size image of the source of the item. Typically, you specify
// your app’s icon but you can provide a different icon if the content has a
// different source.
//
// # Return Value
// 
// An initialized item to share.
//
// See: https://developer.apple.com/documentation/AppKit/NSPreviewRepresentingActivityItem/init(item:title:image:icon:)
func (p NSPreviewRepresentingActivityItem) InitWithItemTitleImageIcon(item objectivec.IObject, title string, image INSImage, icon INSImage) NSPreviewRepresentingActivityItem {
	rv := objc.Send[NSPreviewRepresentingActivityItem](p.ID, objc.Sel("initWithItem:title:image:icon:"), item, objc.String(title), image, icon)
	return rv
}

// Creates a metadata object that provides a title and images for a shareable
// item.
//
// item: The item to share from the Mac share sheet. The item must conform to the
// [NSPasteboardWriting] protocol, or be an [NSItemProvider] or [NSDocument]
// object.
// //
// [NSItemProvider]: https://developer.apple.com/documentation/Foundation/NSItemProvider
//
// title: The localized name of the item.
//
// imageProvider: An object that provides the image to display as a preview for the item. For
// example, when you share a URL for a webpage, you might specify the hero
// image for that page or a rendering of the webpage itself.
//
// iconProvider: An object that a thumbnail-size image of the source of the item. Typically,
// you specify your app’s icon but you can provide a different icon if the
// content has a different source.
//
// # Return Value
// 
// An initialized item to share.
//
// See: https://developer.apple.com/documentation/AppKit/NSPreviewRepresentingActivityItem/init(item:title:imageProvider:iconProvider:)
func (p NSPreviewRepresentingActivityItem) InitWithItemTitleImageProviderIconProvider(item objectivec.IObject, title string, imageProvider foundation.NSItemProvider, iconProvider foundation.NSItemProvider) NSPreviewRepresentingActivityItem {
	rv := objc.Send[NSPreviewRepresentingActivityItem](p.ID, objc.Sel("initWithItem:title:imageProvider:iconProvider:"), item, objc.String(title), imageProvider, iconProvider)
	return rv
}












// An object that provides an icon that represents the item’s source.
//
// # Discussion
// 
// Typically, the icon is a thumbnail-sized representation of the source app
// for the content. For example, provide your app’s icon for content you
// manage.
//
// See: https://developer.apple.com/documentation/AppKit/NSPreviewRepresentableActivityItem/iconProvider
func (p NSPreviewRepresentingActivityItem) IconProvider() foundation.NSItemProvider {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("iconProvider"))
	return foundation.NSItemProviderFromID(objc.ID(rv))
}



// An object that provides a visual representation of the item.
//
// # Discussion
// 
// Provide a full-size representation of the content you’re sharing. For
// example, if the shared item is a link to a webpage, provide the hero image
// for that webpage or a rendering of the page.
//
// See: https://developer.apple.com/documentation/AppKit/NSPreviewRepresentableActivityItem/imageProvider
func (p NSPreviewRepresentingActivityItem) ImageProvider() foundation.NSItemProvider {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("imageProvider"))
	return foundation.NSItemProviderFromID(objc.ID(rv))
}



// A localized string that contains the name of the item.
//
// See: https://developer.apple.com/documentation/AppKit/NSPreviewRepresentableActivityItem/title
func (p NSPreviewRepresentingActivityItem) Title() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}



// The app-specific item you want to share.
//
// # Discussion
// 
// Use this property to provide the data you want to pass to the sharing
// service. The item must conform to the [NSPasteboardWriting] protocol, or be
// an [NSItemProvider] or [NSDocument] object.
//
// [NSItemProvider]: https://developer.apple.com/documentation/Foundation/NSItemProvider
//
// See: https://developer.apple.com/documentation/AppKit/NSPreviewRepresentableActivityItem/item
func (p NSPreviewRepresentingActivityItem) Item() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("item"))
	return objectivec.Object{ID: rv}
}














			// Protocol methods for NSPreviewRepresentableActivityItem
			













