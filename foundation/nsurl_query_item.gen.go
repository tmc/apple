// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSURLQueryItem] class.
var (
	_NSURLQueryItemClass     NSURLQueryItemClass
	_NSURLQueryItemClassOnce sync.Once
)

func getNSURLQueryItemClass() NSURLQueryItemClass {
	_NSURLQueryItemClassOnce.Do(func() {
		_NSURLQueryItemClass = NSURLQueryItemClass{class: objc.GetClass("NSURLQueryItem")}
	})
	return _NSURLQueryItemClass
}

// GetNSURLQueryItemClass returns the class object for NSURLQueryItem.
func GetNSURLQueryItemClass() NSURLQueryItemClass {
	return getNSURLQueryItemClass()
}

type NSURLQueryItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSURLQueryItemClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSURLQueryItemClass) Alloc() NSURLQueryItem {
	rv := objc.Send[NSURLQueryItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object representing a single name/value pair for an item in the query
// portion of a URL.
//
// # Overview
// 
// In Swift, this object bridges to [URLQueryItem]; use [NSURLQueryItem] when
// you need reference semantics or other Foundation-specific behavior.
// 
// You use query items with the [NSURLQueryItem.QueryItems] property of an [NSURLComponents]
// object.
//
// [URLQueryItem]: https://developer.apple.com/documentation/Foundation/URLQueryItem
//
// # Creating a Query Item
//
//   - [NSURLQueryItem.InitWithNameValue]: Initializes a newly allocated query item with the specified name and value.
//
// # Reading a Query Item’s Name and Value
//
//   - [NSURLQueryItem.Name]: The name of the query item.
//   - [NSURLQueryItem.Value]: The value for the query item.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLQueryItem
type NSURLQueryItem struct {
	objectivec.Object
}

// NSURLQueryItemFromID constructs a [NSURLQueryItem] from an objc.ID.
//
// An object representing a single name/value pair for an item in the query
// portion of a URL.
func NSURLQueryItemFromID(id objc.ID) NSURLQueryItem {
	return NSURLQueryItem{objectivec.Object{ID: id}}
}
// NOTE: NSURLQueryItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSURLQueryItem] class.
//
// # Creating a Query Item
//
//   - [INSURLQueryItem.InitWithNameValue]: Initializes a newly allocated query item with the specified name and value.
//
// # Reading a Query Item’s Name and Value
//
//   - [INSURLQueryItem.Name]: The name of the query item.
//   - [INSURLQueryItem.Value]: The value for the query item.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLQueryItem
type INSURLQueryItem interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Creating a Query Item

	// Initializes a newly allocated query item with the specified name and value.
	InitWithNameValue(name string, value string) NSURLQueryItem

	// Topic: Reading a Query Item’s Name and Value

	// The name of the query item.
	Name() string
	// The value for the query item.
	Value() string

	// The query URL component as an array of name/value pairs.
	QueryItems() INSURLQueryItem
	SetQueryItems(value INSURLQueryItem)
}

// Init initializes the instance.
func (u NSURLQueryItem) Init() NSURLQueryItem {
	rv := objc.Send[NSURLQueryItem](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u NSURLQueryItem) Autorelease() NSURLQueryItem {
	rv := objc.Send[NSURLQueryItem](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSURLQueryItem creates a new NSURLQueryItem instance.
func NewNSURLQueryItem() NSURLQueryItem {
	class := getNSURLQueryItemClass()
	rv := objc.Send[NSURLQueryItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewURLQueryItemWithCoder(coder INSCoder) NSURLQueryItem {
	instance := getNSURLQueryItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSURLQueryItemFromID(rv)
}

// Initializes a newly allocated query item with the specified name and value.
//
// name: The name of the query item. For example, in the URL
// `//www.AppleXCUIElementTypeCom()/search/?q=iPad`, the `name` parameter is
// `q`.
//
// value: The value for the query item. For example, in the URL
// `//www.AppleXCUIElementTypeCom()/search/?q=iPad`, the `value` parameter is
// `iPad`.
//
// # Return Value
// 
// An initialized query item object.
//
// # Discussion
// 
// To use the newly initialized query item in composing a URL, add it to the
// [QueryItems] array of an [NSURLComponents] instance. Because assigning an
// array of query items to an [NSURLComponents] instance automatically encodes
// the name and value properties, you should not percent-encode these strings.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLQueryItem/init(name:value:)
func NewURLQueryItemWithNameValue(name string, value string) NSURLQueryItem {
	instance := getNSURLQueryItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:value:"), objc.String(name), objc.String(value))
	return NSURLQueryItemFromID(rv)
}

// Initializes a newly allocated query item with the specified name and value.
//
// name: The name of the query item. For example, in the URL
// `//www.AppleXCUIElementTypeCom()/search/?q=iPad`, the `name` parameter is
// `q`.
//
// value: The value for the query item. For example, in the URL
// `//www.AppleXCUIElementTypeCom()/search/?q=iPad`, the `value` parameter is
// `iPad`.
//
// # Return Value
// 
// An initialized query item object.
//
// # Discussion
// 
// To use the newly initialized query item in composing a URL, add it to the
// [QueryItems] array of an [NSURLComponents] instance. Because assigning an
// array of query items to an [NSURLComponents] instance automatically encodes
// the name and value properties, you should not percent-encode these strings.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLQueryItem/init(name:value:)
func (u NSURLQueryItem) InitWithNameValue(name string, value string) NSURLQueryItem {
	rv := objc.Send[NSURLQueryItem](u.ID, objc.Sel("initWithName:value:"), objc.String(name), objc.String(value))
	return rv
}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (u NSURLQueryItem) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (u NSURLQueryItem) InitWithCoder(coder INSCoder) NSURLQueryItem {
	rv := objc.Send[NSURLQueryItem](u.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Creates a new query item with the specified name and value.
//
// name: The name of the query item. For example, in the URL
// `//www.AppleXCUIElementTypeCom()/search/?q=iPad`, the `name` parameter is
// `q`.
//
// value: The value for the query item. For example, in the URL
// `//www.AppleXCUIElementTypeCom()/search/?q=iPad`, the `value` parameter is
// `iPad`.
//
// # Return Value
// 
// A new query item object.
//
// # Discussion
// 
// To use the newly initialized query item in composing a URL, add it to the
// [QueryItems] array of an [NSURLComponents] instance. Because assigning an
// array of query items to an [NSURLComponents] instance automatically encodes
// the name and value properties, you should not percent-encode these strings.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLQueryItem/queryItemWithName:value:
func (_NSURLQueryItemClass NSURLQueryItemClass) QueryItemWithNameValue(name string, value string) NSURLQueryItem {
	rv := objc.Send[objc.ID](objc.ID(_NSURLQueryItemClass.class), objc.Sel("queryItemWithName:value:"), objc.String(name), objc.String(value))
	return NSURLQueryItemFromID(rv)
}

// The name of the query item.
//
// # Discussion
// 
// For example, in the URL `//www.AppleXCUIElementTypeCom()/search/?q=iPad`,
// the `name` parameter is `q`.
// 
// This string is not percent-encoded.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLQueryItem/name
func (u NSURLQueryItem) Name() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("name"))
	return NSStringFromID(rv).String()
}
// The value for the query item.
//
// # Discussion
// 
// For example, in the URL `//www.AppleXCUIElementTypeCom()/search/?q=iPad`,
// the `value` parameter is `iPad`.
// 
// This string is not percent-encoded.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLQueryItem/value
func (u NSURLQueryItem) Value() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("value"))
	return NSStringFromID(rv).String()
}
// The query URL component as an array of name/value pairs.
//
// See: https://developer.apple.com/documentation/foundation/nsurlcomponents/queryitems
func (u NSURLQueryItem) QueryItems() INSURLQueryItem {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("queryItems"))
	return NSURLQueryItemFromID(objc.ID(rv))
}
func (u NSURLQueryItem) SetQueryItems(value INSURLQueryItem) {
	objc.Send[struct{}](u.ID, objc.Sel("setQueryItems:"), value)
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

