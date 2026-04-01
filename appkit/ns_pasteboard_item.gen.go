// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPasteboardItem] class.
var (
	_NSPasteboardItemClass     NSPasteboardItemClass
	_NSPasteboardItemClassOnce sync.Once
)

func getNSPasteboardItemClass() NSPasteboardItemClass {
	_NSPasteboardItemClassOnce.Do(func() {
		_NSPasteboardItemClass = NSPasteboardItemClass{class: objc.GetClass("NSPasteboardItem")}
	})
	return _NSPasteboardItemClass
}

// GetNSPasteboardItemClass returns the class object for NSPasteboardItem.
func GetNSPasteboardItemClass() NSPasteboardItemClass {
	return getNSPasteboardItemClass()
}

type NSPasteboardItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPasteboardItemClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPasteboardItemClass) Alloc() NSPasteboardItem {
	rv := objc.Send[NSPasteboardItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An item on a pasteboard.
//
// # Overview
//
// There are three main uses for an [NSPasteboardItem] object:
//
// - Providing data on the pasteboard.
//
// You can create one or more pasteboard items, set data or data providers for
// types, and write them to the pasteboard.
//
// - Customizing data already on the pasteboard.
//
// As a delegate or subclass, you can retrieve the pasteboard items currently
// on the pasteboard, read the existing types and data, and set new data and
// data providers for types as necessary.
//
// - Retrieving data from the pasteboard.
//
// You can retrieve pasteboard items from the pasteboard and then read the
// data for types you’re interested in.
//
// A pasteboard item can be associated with a single pasteboard. When you
// create an item, you can write it to any pasteboard. When you pass an item
// to a pasteboard in [WriteObjects], that item becomes bound to the
// pasteboard it writes to. When you retrieve items from a pasteboard using
// [NSPasteboardItem.PasteboardItems] or [ReadObjectsForClassesOptions], the returned items are
// associated with the messaged pasteboard. Passing an item that is already
// associated with a pasteboard into [WriteObjects] causes an exception.
//
// Use pasteboard items during a single pasteboard interaction, rather than
// retaining and reusing them. A pasteboard item is only valid until the owner
// of the pasteboard changes.
//
// # Getting types
//
//   - [NSPasteboardItem.Types]: An array of uniform type identifier strings of the data types that the receiver supports.
//   - [NSPasteboardItem.AvailableTypeFromArray]: Returns from a given array of types the first type within the pasteboard item, according to the ordering of types.
//
// # Setting the data provider
//
//   - [NSPasteboardItem.SetDataProviderForTypes]: Sets the data provider for the specified types.
//
// # Setting values
//
//   - [NSPasteboardItem.SetDataForType]: Sets the value for a specified type as a data object.
//   - [NSPasteboardItem.SetStringForType]: Sets the value for a specified type as a string.
//   - [NSPasteboardItem.SetPropertyListForType]: Sets the value for a specified type as a property list.
//
// # Getting values
//
//   - [NSPasteboardItem.DataForType]: Returns the value for the specified type as a data object.
//   - [NSPasteboardItem.StringForType]: Returns the value for the specified type as a string.
//   - [NSPasteboardItem.PropertyListForType]: Returns the value for the specified type as a property list.
//
// # Instance Properties
//
//   - [NSPasteboardItem.CollaborationMetadata]: A model object you use for conveying data during a collaboration.
//   - [NSPasteboardItem.SetCollaborationMetadata]
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardItem
type NSPasteboardItem struct {
	objectivec.Object
}

// NSPasteboardItemFromID constructs a [NSPasteboardItem] from an objc.ID.
//
// An item on a pasteboard.
func NSPasteboardItemFromID(id objc.ID) NSPasteboardItem {
	return NSPasteboardItem{objectivec.Object{ID: id}}
}

// NOTE: NSPasteboardItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPasteboardItem] class.
//
// # Getting types
//
//   - [INSPasteboardItem.Types]: An array of uniform type identifier strings of the data types that the receiver supports.
//   - [INSPasteboardItem.AvailableTypeFromArray]: Returns from a given array of types the first type within the pasteboard item, according to the ordering of types.
//
// # Setting the data provider
//
//   - [INSPasteboardItem.SetDataProviderForTypes]: Sets the data provider for the specified types.
//
// # Setting values
//
//   - [INSPasteboardItem.SetDataForType]: Sets the value for a specified type as a data object.
//   - [INSPasteboardItem.SetStringForType]: Sets the value for a specified type as a string.
//   - [INSPasteboardItem.SetPropertyListForType]: Sets the value for a specified type as a property list.
//
// # Getting values
//
//   - [INSPasteboardItem.DataForType]: Returns the value for the specified type as a data object.
//   - [INSPasteboardItem.StringForType]: Returns the value for the specified type as a string.
//   - [INSPasteboardItem.PropertyListForType]: Returns the value for the specified type as a property list.
//
// # Instance Properties
//
//   - [INSPasteboardItem.CollaborationMetadata]: A model object you use for conveying data during a collaboration.
//   - [INSPasteboardItem.SetCollaborationMetadata]
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardItem
type INSPasteboardItem interface {
	objectivec.IObject
	NSPasteboardWriting

	// Topic: Getting types

	// An array of uniform type identifier strings of the data types that the receiver supports.
	Types() []string
	// Returns from a given array of types the first type within the pasteboard item, according to the ordering of types.
	AvailableTypeFromArray(types []string) NSPasteboardType

	// Topic: Setting the data provider

	// Sets the data provider for the specified types.
	SetDataProviderForTypes(dataProvider NSPasteboardItemDataProvider, types []string) bool

	// Topic: Setting values

	// Sets the value for a specified type as a data object.
	SetDataForType(data foundation.INSData, type_ NSPasteboardType) bool
	// Sets the value for a specified type as a string.
	SetStringForType(string_ string, type_ NSPasteboardType) bool
	// Sets the value for a specified type as a property list.
	SetPropertyListForType(propertyList objectivec.IObject, type_ NSPasteboardType) bool

	// Topic: Getting values

	// Returns the value for the specified type as a data object.
	DataForType(type_ NSPasteboardType) foundation.INSData
	// Returns the value for the specified type as a string.
	StringForType(type_ NSPasteboardType) string
	// Returns the value for the specified type as a property list.
	PropertyListForType(type_ NSPasteboardType) objectivec.IObject

	// Topic: Instance Properties

	// A model object you use for conveying data during a collaboration.
	CollaborationMetadata() unsafe.Pointer
	SetCollaborationMetadata(value unsafe.Pointer)

	// An array that contains all the items held by the pasteboard.
	PasteboardItems() INSPasteboardItem
	SetPasteboardItems(value INSPasteboardItem)
	// Initializes an instance with a property list object and a type string.
	InitWithPasteboardPropertyListOfType(propertyList objectivec.IObject, type_ NSPasteboardType) NSPasteboardItem
}

// Init initializes the instance.
func (p NSPasteboardItem) Init() NSPasteboardItem {
	rv := objc.Send[NSPasteboardItem](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPasteboardItem) Autorelease() NSPasteboardItem {
	rv := objc.Send[NSPasteboardItem](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPasteboardItem creates a new NSPasteboardItem instance.
func NewNSPasteboardItem() NSPasteboardItem {
	class := getNSPasteboardItemClass()
	rv := objc.Send[NSPasteboardItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes an instance with a property list object and a type string.
//
// propertyList: A property list containing data to initialize the receiver.
//
// By default, the property list object is an instance of [NSData]. If you
// implement [ReadingOptionsForTypePasteboard] and specify an option other
// than [NSPasteboardReadingAsData], the `propertyList` may be any other
// property list object.
//
// type: A UTI supported by the receiver for reading (one of the types returned by
// [ReadableTypesForPasteboard]).
//
// # Return Value
//
// An object initialized using the data in `propertyList`.
//
// # Discussion
//
// This method is considered optional because, if [ReadableTypesForPasteboard]
// returns just a single type, and that type uses the
// [NSPasteboardReadingAsKeyedArchive] reading option, then instances are
// initialized using [init(coder:)] instead of this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/init(pasteboardPropertyList:ofType:)
//
// [init(coder:)]: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewPasteboardItemWithPasteboardPropertyListOfType(propertyList objectivec.IObject, type_ NSPasteboardType) NSPasteboardItem {
	instance := getNSPasteboardItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPasteboardPropertyList:ofType:"), propertyList, objc.String(string(type_)))
	return NSPasteboardItemFromID(rv)
}

// Returns from a given array of types the first type within the pasteboard
// item, according to the ordering of types.
//
// types: An array of strings representing UTIs, arranged in order of preference
// (most preferred as the 0th element in the array).
//
// # Return Value
//
// The first (according to the sender’s ordering of `types`) type in `types`
// contained in the pasteboard item, or `nil` if the receiver does not contain
// any types given in `types`.
//
// # Discussion
//
// The method checks for UTI conformance of the requested types, preferring an
// exact match to conformance.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/availableType(from:)
func (p NSPasteboardItem) AvailableTypeFromArray(types []string) NSPasteboardType {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("availableTypeFromArray:"), objectivec.StringSliceToNSArray(types))
	return NSPasteboardType(foundation.NSStringFromID(rv).String())
}

// Sets the data provider for the specified types.
//
// dataProvider: A pasteboard data provider.
//
// types: An array of strings indicating the UTIs for the data representations
// `dataProvider` may provide.
//
// # Return Value
//
// true if the data provider was set successfully, otherwise false.
//
// # Discussion
//
// This method registers the data provider to be messaged to provide the data
// for any of the specified types when requested.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/setDataProvider(_:forTypes:)
func (p NSPasteboardItem) SetDataProviderForTypes(dataProvider NSPasteboardItemDataProvider, types []string) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("setDataProvider:forTypes:"), dataProvider, objectivec.StringSliceToNSArray(types))
	return rv
}

// Sets the value for a specified type as a data object.
//
// data: An [NSData] object containing the value for the representation specified by
// `type`.
//
// type: A uniform type identifier string.
//
// # Return Value
//
// true if the value was set successfully, otherwise false.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/setData(_:forType:)
func (p NSPasteboardItem) SetDataForType(data foundation.INSData, type_ NSPasteboardType) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("setData:forType:"), data, objc.String(string(type_)))
	return rv
}

// Sets the value for a specified type as a string.
//
// string: A string for the representation specified by `type`.
//
// type: A uniform type identifier string.
//
// # Return Value
//
// true if the value was set successfully, otherwise false.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/setString(_:forType:)
func (p NSPasteboardItem) SetStringForType(string_ string, type_ NSPasteboardType) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("setString:forType:"), objc.String(string_), objc.String(string(type_)))
	return rv
}

// Sets the value for a specified type as a property list.
//
// propertyList: A property list object containing the value for the representation
// specified by `type`.
//
// type: A uniform type identifier string.
//
// # Return Value
//
// true if the value was set successfully, otherwise false.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/setPropertyList(_:forType:)
func (p NSPasteboardItem) SetPropertyListForType(propertyList objectivec.IObject, type_ NSPasteboardType) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("setPropertyList:forType:"), propertyList, objc.String(string(type_)))
	return rv
}

// Returns the value for the specified type as a data object.
//
// type: A uniform type identifier string.
//
// # Return Value
//
// The value for the specified type as an [NSData] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/data(forType:)
func (p NSPasteboardItem) DataForType(type_ NSPasteboardType) foundation.INSData {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("dataForType:"), objc.String(string(type_)))
	return foundation.NSDataFromID(rv)
}

// Returns the value for the specified type as a string.
//
// type: A uniform type identifier string.
//
// # Return Value
//
// The value for the specified type as a string.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/string(forType:)
func (p NSPasteboardItem) StringForType(type_ NSPasteboardType) string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("stringForType:"), objc.String(string(type_)))
	return foundation.NSStringFromID(rv).String()
}

// Returns the value for the specified type as a property list.
//
// type: A uniform type identifier string.
//
// # Return Value
//
// The value for the specified type as a property list.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/propertyList(forType:)
func (p NSPasteboardItem) PropertyListForType(type_ NSPasteboardType) objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("propertyListForType:"), objc.String(string(type_)))
	return objectivec.Object{ID: rv}
}

// Initializes an instance with a property list object and a type string.
//
// propertyList: A property list containing data to initialize the receiver.
//
// By default, the property list object is an instance of [NSData]. If you
// implement [ReadingOptionsForTypePasteboard] and specify an option other
// than [NSPasteboardReadingAsData], the `propertyList` may be any other
// property list object.
//
// type: A UTI supported by the receiver for reading (one of the types returned by
// [ReadableTypesForPasteboard]).
//
// # Return Value
//
// An object initialized using the data in `propertyList`.
//
// # Discussion
//
// This method is considered optional because, if [ReadableTypesForPasteboard]
// returns just a single type, and that type uses the
// [NSPasteboardReadingAsKeyedArchive] reading option, then instances are
// initialized using [init(coder:)] instead of this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/init(pasteboardPropertyList:ofType:)
//
// [init(coder:)]: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (p NSPasteboardItem) InitWithPasteboardPropertyListOfType(propertyList objectivec.IObject, type_ NSPasteboardType) NSPasteboardItem {
	rv := objc.Send[NSPasteboardItem](p.ID, objc.Sel("initWithPasteboardPropertyList:ofType:"), propertyList, objc.String(string(type_)))
	return rv
}

// Returns a property list object to represent the receiver on a pasteboard as
// an object of a specified type.
//
// type: One of the types the receiver supports for writing (one of the UTIs
// returned by its implementation of [WritableTypesForPasteboard]).
//
// # Return Value
//
// A property list object to represent the receiver on a pasteboard as an
// object of type `type`.
//
// # Discussion
//
// The returned value will commonly be the [NSData] object for the specified
// data type. However, if this method returns either a string, or any other
// property-list type, the pasteboard will automatically convert these items
// to the correct data format required for the pasteboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardWriting/pasteboardPropertyList(forType:)
func (p NSPasteboardItem) PasteboardPropertyListForType(type_ NSPasteboardType) objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("pasteboardPropertyListForType:"), objc.String(string(type_)))
	return objectivec.Object{ID: rv}
}

// Returns an array of UTI strings of data types the receiver can write to a
// given pasteboard.
//
// pasteboard: A pasteboard.
//
// You can use this argument to provide different options based on the
// pasteboard name, if you need to.
//
// # Return Value
//
// An array of UTI strings of data types the receiver can write to
// `pasteboard`.
//
// # Discussion
//
// By default, data for the first returned type is put onto the pasteboard
// immediately, with the remaining types being promised.
//
// To change the default behavior, implement
// -writingOptionsForType:pasteboard: and return [NSPasteboardWritingPromised]
// to lazily provide data for types, return no option to provide the data for
// that type immediately. Use the pasteboard argument to provide different
// types based on the pasteboard name, if desired. Do not perform other
// pasteboard operations in the method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardWriting/writableTypes(for:)
func (p NSPasteboardItem) WritableTypesForPasteboard(pasteboard INSPasteboard) []string {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("writableTypesForPasteboard:"), pasteboard)
	return objc.ConvertSliceToStrings(rv)
}

// Returns options for writing data of a specified type to a given pasteboard.
//
// type: One of the types the receiver supports for writing (one of the UTIs
// returned by its implementation of [WritableTypesForPasteboard]).
//
// pasteboard: A pasteboard.
//
// You can use this argument to provide different options based on the
// pasteboard name, if you need to.
//
// # Return Value
//
// Options for writing data of type type to `pasteboard`. Return `0` for no
// options, or a value given in [Pasteboard Writing Options].
//
// # Discussion
//
// Do not perform other pasteboard operations in the method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardWriting/writingOptions(forType:pasteboard:)
//
// [Pasteboard Writing Options]: https://developer.apple.com/documentation/AppKit/pasteboard-writing-options
func (p NSPasteboardItem) WritingOptionsForTypePasteboard(type_ NSPasteboardType, pasteboard INSPasteboard) NSPasteboardWritingOptions {
	rv := objc.Send[NSPasteboardWritingOptions](p.ID, objc.Sel("writingOptionsForType:pasteboard:"), objc.String(string(type_)), pasteboard)
	return NSPasteboardWritingOptions(rv)
}

// Returns an array of uniform type identifier strings of data types the
// receiver can read from the pasteboard and initialize from.
//
// pasteboard: A pasteboard. You can use the pasteboard argument to provide different
// types based on the pasteboard name, if you need to.
//
// # Return Value
//
// An array of uniform type identifier strings of data types instances that
// the receiver can read from the pasteboard and initialize from.
//
// # Discussion
//
// By default, the system provides the data for a type to
// [InitWithPasteboardPropertyListOfType] as an instance of [NSData]. If you
// implement [ReadingOptionsForTypePasteboard] and specify a different option,
// the system converts the [NSData] object for a type to an [NSString] object
// or any other property list object.
//
// # Special Considerations
//
// Don’t perform other pasteboard operations in the method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/readableTypes(for:)
func (_NSPasteboardItemClass NSPasteboardItemClass) ReadableTypesForPasteboard(pasteboard INSPasteboard) []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSPasteboardItemClass.class), objc.Sel("readableTypesForPasteboard:"), pasteboard)
	return objc.ConvertSliceToStrings(rv)
}

// Returns options for reading data of a specified type from a given
// pasteboard.
//
// type: A UTI supported by instances of the receiver for reading (one of the types
// returned by [ReadableTypesForPasteboard]).
//
// pasteboard: A pasteboard.
//
// You can use the pasteboard argument to provide return different based on
// the pasteboard name, should you need to do so.
//
// # Return Value
//
// Options for reading data of `type` from `pasteboard`. For a list of valid
// values, see [NSPasteboard.ReadingOptions].
//
// # Discussion
//
// Do not perform other pasteboard operations in this method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/readingOptions(forType:pasteboard:)
//
// [NSPasteboard.ReadingOptions]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptions
func (_NSPasteboardItemClass NSPasteboardItemClass) ReadingOptionsForTypePasteboard(type_ NSPasteboardType, pasteboard INSPasteboard) NSPasteboardReadingOptions {
	rv := objc.Send[NSPasteboardReadingOptions](objc.ID(_NSPasteboardItemClass.class), objc.Sel("readingOptionsForType:pasteboard:"), objc.String(string(type_)), pasteboard)
	return NSPasteboardReadingOptions(rv)
}

// An array of uniform type identifier strings of the data types that the
// receiver supports.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/types
func (p NSPasteboardItem) Types() []string {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("types"))
	return objc.ConvertSliceToStrings(rv)
}

// A model object you use for conveying data during a collaboration.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/collaborationMetadata
func (p NSPasteboardItem) CollaborationMetadata() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p.ID, objc.Sel("collaborationMetadata"))
	return rv
}
func (p NSPasteboardItem) SetCollaborationMetadata(value unsafe.Pointer) {
	objc.Send[struct{}](p.ID, objc.Sel("setCollaborationMetadata:"), value)
}

// An array that contains all the items held by the pasteboard.
//
// See: https://developer.apple.com/documentation/appkit/nspasteboard/pasteboarditems
func (p NSPasteboardItem) PasteboardItems() INSPasteboardItem {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("pasteboardItems"))
	return NSPasteboardItemFromID(objc.ID(rv))
}
func (p NSPasteboardItem) SetPasteboardItems(value INSPasteboardItem) {
	objc.Send[struct{}](p.ID, objc.Sel("setPasteboardItems:"), value)
}

// Protocol methods for NSPasteboardWriting
