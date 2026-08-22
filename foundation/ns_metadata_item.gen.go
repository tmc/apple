// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMetadataItem] class.
var (
	_NSMetadataItemClass     NSMetadataItemClass
	_NSMetadataItemClassOnce sync.Once
)

func getNSMetadataItemClass() NSMetadataItemClass {
	_NSMetadataItemClassOnce.Do(func() {
		_NSMetadataItemClass = NSMetadataItemClass{class: objc.GetClass("NSMetadataItem")}
	})
	return _NSMetadataItemClass
}

// GetNSMetadataItemClass returns the class object for NSMetadataItem.
func GetNSMetadataItemClass() NSMetadataItemClass {
	return getNSMetadataItemClass()
}

type NSMetadataItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSMetadataItemClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMetadataItemClass) Alloc() NSMetadataItem {
	rv := objc.Send[NSMetadataItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The metadata associated with a file.
//
// # Overview
//
// Metadata items provide a simple interface to retrieve the available
// attribute names and values.
//
// # Creating a Metadata Item
//
//   - [NSMetadataItem.InitWithURL]: Initializes a metadata item with a given URL.
//
// # Getting Item Attributes
//
//   - [NSMetadataItem.Attributes]: An array containing the attribute keys for the metadata item’s values.
//   - [NSMetadataItem.ValueForAttribute]: Returns the receiver’s metadata attribute name specified by a given key.
//   - [NSMetadataItem.ValuesForAttributes]: Returns a dictionary containing the key-value pairs for the attribute names specified by a given array of keys.
//
// See: https://developer.apple.com/documentation/Foundation/NSMetadataItem
type NSMetadataItem struct {
	objectivec.Object
}

// NSMetadataItemFromID constructs a [NSMetadataItem] from an objc.ID.
//
// The metadata associated with a file.
func NSMetadataItemFromID(id objc.ID) NSMetadataItem {
	return NSMetadataItem{objectivec.Object{ID: id}}
}

// NOTE: NSMetadataItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMetadataItem] class.
//
// # Creating a Metadata Item
//
//   - [INSMetadataItem.InitWithURL]: Initializes a metadata item with a given URL.
//
// # Getting Item Attributes
//
//   - [INSMetadataItem.Attributes]: An array containing the attribute keys for the metadata item’s values.
//   - [INSMetadataItem.ValueForAttribute]: Returns the receiver’s metadata attribute name specified by a given key.
//   - [INSMetadataItem.ValuesForAttributes]: Returns a dictionary containing the key-value pairs for the attribute names specified by a given array of keys.
//
// See: https://developer.apple.com/documentation/Foundation/NSMetadataItem
type INSMetadataItem interface {
	objectivec.IObject

	// Topic: Creating a Metadata Item

	// Initializes a metadata item with a given URL.
	InitWithURL(url INSURL) NSMetadataItem

	// Topic: Getting Item Attributes

	// An array containing the attribute keys for the metadata item’s values.
	Attributes() []string
	// Returns the receiver’s metadata attribute name specified by a given key.
	ValueForAttribute(key string) objectivec.IObject
	// Returns a dictionary containing the key-value pairs for the attribute names specified by a given array of keys.
	ValuesForAttributes(keys []string) INSDictionary
}

// Init initializes the instance.
func (m NSMetadataItem) Init() NSMetadataItem {
	rv := objc.Send[NSMetadataItem](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMetadataItem) Autorelease() NSMetadataItem {
	rv := objc.Send[NSMetadataItem](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMetadataItem creates a new NSMetadataItem instance.
func NewNSMetadataItem() NSMetadataItem {
	class := getNSMetadataItemClass()
	rv := objc.Send[NSMetadataItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a metadata item with a given URL.
//
// url: The URL for the metadata item.
//
// # Return Value
//
// A metadata item for the file identified by `url`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMetadataItem/init(url:)
func NewMetadataItemWithURL(url INSURL) NSMetadataItem {
	instance := getNSMetadataItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), url)
	return NSMetadataItemFromID(rv)
}

// Initializes a metadata item with a given URL.
//
// url: The URL for the metadata item.
//
// # Return Value
//
// A metadata item for the file identified by `url`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMetadataItem/init(url:)
func (m NSMetadataItem) InitWithURL(url INSURL) NSMetadataItem {
	rv := objc.Send[NSMetadataItem](m.ID, objc.Sel("initWithURL:"), url)
	return rv
}

// Returns the receiver’s metadata attribute name specified by a given key.
//
// key: The name of a metadata attribute. See the “Constants” section for a
// list of possible keys.
//
// # Return Value
//
// The receiver’s metadata attribute name specified by `key`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMetadataItem/value(forAttribute:)
func (m NSMetadataItem) ValueForAttribute(key string) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("valueForAttribute:"), objc.String(key))
	return objectivec.Object{ID: rv}
}

// Returns a dictionary containing the key-value pairs for the attribute names
// specified by a given array of keys.
//
// keys: An array containing [NSString] objects that specify the names of a metadata
// attributes. See the “Constants” section for a list of possible keys.
//
// # Return Value
//
// A dictionary containing the key-value pairs for the attribute names
// specified by `keys`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMetadataItem/values(forAttributes:)
func (m NSMetadataItem) ValuesForAttributes(keys []string) INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("valuesForAttributes:"), objectivec.StringSliceToNSArray(keys))
	return NSDictionaryFromID(rv)
}

// An array containing the attribute keys for the metadata item’s values.
//
// # Discussion
//
// This property contains an array of attribute keys, representing the values
// available from this metadata item. For a list of possible keys, see
// `Attribute Keys`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMetadataItem/attributes
func (m NSMetadataItem) Attributes() []string {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("attributes"))
	return objc.ConvertSliceToStrings(rv)
}
