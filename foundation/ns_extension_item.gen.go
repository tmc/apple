// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSExtensionItem] class.
var (
	_NSExtensionItemClass     NSExtensionItemClass
	_NSExtensionItemClassOnce sync.Once
)

func getNSExtensionItemClass() NSExtensionItemClass {
	_NSExtensionItemClassOnce.Do(func() {
		_NSExtensionItemClass = NSExtensionItemClass{class: objc.GetClass("NSExtensionItem")}
	})
	return _NSExtensionItemClass
}

// GetNSExtensionItemClass returns the class object for NSExtensionItem.
func GetNSExtensionItemClass() NSExtensionItemClass {
	return getNSExtensionItemClass()
}

type NSExtensionItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSExtensionItemClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSExtensionItemClass) Alloc() NSExtensionItem {
	rv := objc.Send[NSExtensionItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An immutable collection of values representing different aspects of an item
// for an extension to act upon.
//
// # Identifying the Item
//
//   - [NSExtensionItem.AttributedTitle]: An optional title for the item.
//   - [NSExtensionItem.SetAttributedTitle]
//   - [NSExtensionItem.UserInfo]: An optional dictionary of keys and values corresponding to the extension item’s properties.
//   - [NSExtensionItem.SetUserInfo]
//
// # Item Contents
//
//   - [NSExtensionItem.Attachments]: An optional array of media data associated with the extension item.
//   - [NSExtensionItem.SetAttachments]
//   - [NSExtensionItem.AttributedContentText]: An optional string describing the extension item content.
//   - [NSExtensionItem.SetAttributedContentText]
//
// See: https://developer.apple.com/documentation/Foundation/NSExtensionItem
type NSExtensionItem struct {
	objectivec.Object
}

// NSExtensionItemFromID constructs a [NSExtensionItem] from an objc.ID.
//
// An immutable collection of values representing different aspects of an item
// for an extension to act upon.
func NSExtensionItemFromID(id objc.ID) NSExtensionItem {
	return NSExtensionItem{objectivec.Object{ID: id}}
}
// NOTE: NSExtensionItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSExtensionItem] class.
//
// # Identifying the Item
//
//   - [INSExtensionItem.AttributedTitle]: An optional title for the item.
//   - [INSExtensionItem.SetAttributedTitle]
//   - [INSExtensionItem.UserInfo]: An optional dictionary of keys and values corresponding to the extension item’s properties.
//   - [INSExtensionItem.SetUserInfo]
//
// # Item Contents
//
//   - [INSExtensionItem.Attachments]: An optional array of media data associated with the extension item.
//   - [INSExtensionItem.SetAttachments]
//   - [INSExtensionItem.AttributedContentText]: An optional string describing the extension item content.
//   - [INSExtensionItem.SetAttributedContentText]
//
// See: https://developer.apple.com/documentation/Foundation/NSExtensionItem
type INSExtensionItem interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Identifying the Item

	// An optional title for the item.
	AttributedTitle() INSAttributedString
	SetAttributedTitle(value INSAttributedString)
	// An optional dictionary of keys and values corresponding to the extension item’s properties.
	UserInfo() INSDictionary
	SetUserInfo(value INSDictionary)

	// Topic: Item Contents

	// An optional array of media data associated with the extension item.
	Attachments() []NSItemProvider
	SetAttachments(value []NSItemProvider)
	// An optional string describing the extension item content.
	AttributedContentText() INSAttributedString
	SetAttributedContentText(value INSAttributedString)
}

// Init initializes the instance.
func (e NSExtensionItem) Init() NSExtensionItem {
	rv := objc.Send[NSExtensionItem](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e NSExtensionItem) Autorelease() NSExtensionItem {
	rv := objc.Send[NSExtensionItem](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSExtensionItem creates a new NSExtensionItem instance.
func NewNSExtensionItem() NSExtensionItem {
	class := getNSExtensionItemClass()
	rv := objc.Send[NSExtensionItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewExtensionItemWithCoder(coder INSCoder) NSExtensionItem {
	instance := getNSExtensionItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSExtensionItemFromID(rv)
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (e NSExtensionItem) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](e.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (e NSExtensionItem) InitWithCoder(coder INSCoder) NSExtensionItem {
	rv := objc.Send[NSExtensionItem](e.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// An optional title for the item.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/NSExtensionItem/attributedTitle
func (e NSExtensionItem) AttributedTitle() INSAttributedString {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("attributedTitle"))
	return NSAttributedStringFromID(objc.ID(rv))
}
func (e NSExtensionItem) SetAttributedTitle(value INSAttributedString) {
	objc.Send[struct{}](e.ID, objc.Sel("setAttributedTitle:"), value)
}
// An optional dictionary of keys and values corresponding to the extension
// item’s properties.
//
// # Discussion
// 
// If applicable to a particular extension type, additional information may be
// available in the `userInfo` dictionary. For example, in the context of an
// Action extension, the `userInfo` dictionary may contain values for the keys
// [NSExtensionItemAttachmentsKey], [NSExtensionItemAttributedContentTextKey],
// and [NSExtensionItemAttributedTitleKey].
//
// [NSExtensionItemAttachmentsKey]: https://developer.apple.com/documentation/Foundation/NSExtensionItemAttachmentsKey
// [NSExtensionItemAttributedContentTextKey]: https://developer.apple.com/documentation/Foundation/NSExtensionItemAttributedContentTextKey
// [NSExtensionItemAttributedTitleKey]: https://developer.apple.com/documentation/Foundation/NSExtensionItemAttributedTitleKey
//
// See: https://developer.apple.com/documentation/Foundation/NSExtensionItem/userInfo
func (e NSExtensionItem) UserInfo() INSDictionary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("userInfo"))
	return NSDictionaryFromID(objc.ID(rv))
}
func (e NSExtensionItem) SetUserInfo(value INSDictionary) {
	objc.Send[struct{}](e.ID, objc.Sel("setUserInfo:"), value)
}
// An optional array of media data associated with the extension item.
//
// # Discussion
// 
// Populate this array with images, videos, URLs, and so on. It’s not meant
// to be an array of alternative data formats or types, but is instead a
// collection to include in a social media post, for example. These items are
// always typed [NSItemProvider].
//
// See: https://developer.apple.com/documentation/Foundation/NSExtensionItem/attachments
func (e NSExtensionItem) Attachments() []NSItemProvider {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("attachments"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSItemProvider {
		return NSItemProviderFromID(id)
	})
}
func (e NSExtensionItem) SetAttachments(value []NSItemProvider) {
	objc.Send[struct{}](e.ID, objc.Sel("setAttachments:"), objectivec.IObjectSliceToNSArray(value))
}
// An optional string describing the extension item content.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/NSExtensionItem/attributedContentText
func (e NSExtensionItem) AttributedContentText() INSAttributedString {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("attributedContentText"))
	return NSAttributedStringFromID(objc.ID(rv))
}
func (e NSExtensionItem) SetAttributedContentText(value INSAttributedString) {
	objc.Send[struct{}](e.ID, objc.Sel("setAttributedContentText:"), value)
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

