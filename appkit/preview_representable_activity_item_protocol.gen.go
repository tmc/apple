// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An interface you adopt in custom objects that you want to share using the macOS share sheet.
//
// See: https://developer.apple.com/documentation/AppKit/NSPreviewRepresentableActivityItem
type NSPreviewRepresentableActivityItem interface {
	objectivec.IObject

	// The app-specific item you want to share.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPreviewRepresentableActivityItem/item
	Item() objectivec.IObject

	// A localized string that contains the name of the item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPreviewRepresentableActivityItem/title
	Title() string

	// An object that provides a visual representation of the item.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPreviewRepresentableActivityItem/imageProvider
	ImageProvider() foundation.NSItemProvider

	// An object that provides an icon that represents the item’s source.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPreviewRepresentableActivityItem/iconProvider
	IconProvider() foundation.NSItemProvider
}

// NSPreviewRepresentableActivityItemObject wraps an existing Objective-C object that conforms to the NSPreviewRepresentableActivityItem protocol.
type NSPreviewRepresentableActivityItemObject struct {
	objectivec.Object
}

func (o NSPreviewRepresentableActivityItemObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSPreviewRepresentableActivityItemObjectFromID constructs a [NSPreviewRepresentableActivityItemObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSPreviewRepresentableActivityItemObjectFromID(id objc.ID) NSPreviewRepresentableActivityItemObject {
	return NSPreviewRepresentableActivityItemObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The app-specific item you want to share.
//
// # Discussion
//
// Use this property to provide the data you want to pass to the sharing
// service. The item must conform to the [NSPasteboardWriting] protocol, or be
// an [NSItemProvider] or [NSDocument] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSPreviewRepresentableActivityItem/item
//
// [NSItemProvider]: https://developer.apple.com/documentation/Foundation/NSItemProvider
func (o NSPreviewRepresentableActivityItemObject) Item() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("item"))
	return objectivec.Object{ID: rv}
}

// A localized string that contains the name of the item.
//
// See: https://developer.apple.com/documentation/AppKit/NSPreviewRepresentableActivityItem/title
func (o NSPreviewRepresentableActivityItemObject) Title() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
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
func (o NSPreviewRepresentableActivityItemObject) ImageProvider() foundation.NSItemProvider {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("imageProvider"))
	return foundation.NSItemProviderFromID(rv)
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
func (o NSPreviewRepresentableActivityItemObject) IconProvider() foundation.NSItemProvider {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("iconProvider"))
	return foundation.NSItemProviderFromID(rv)
}
