// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/uniformtypeidentifiers"
)

// The class instance for the [NSPasteboard] class.
var (
	_NSPasteboardClass     NSPasteboardClass
	_NSPasteboardClassOnce sync.Once
)

func getNSPasteboardClass() NSPasteboardClass {
	_NSPasteboardClassOnce.Do(func() {
		_NSPasteboardClass = NSPasteboardClass{class: objc.GetClass("NSPasteboard")}
	})
	return _NSPasteboardClass
}

// GetNSPasteboardClass returns the class object for NSPasteboard.
func GetNSPasteboardClass() NSPasteboardClass {
	return getNSPasteboardClass()
}

type NSPasteboardClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPasteboardClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPasteboardClass) Alloc() NSPasteboard {
	rv := objc.Send[NSPasteboard](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that transfers data to and from the pasteboard server.
//
// # Overview
// 
// The pasteboard server is shared by all running apps. It contains data that
// the user has cut or copied, as well as other data that one application
// wants to transfer to another. [NSPasteboard] objects are an application’s
// sole interface to the server and to all pasteboard operations.
// 
// An [NSPasteboard] object is also used to transfer data between apps and
// service providers listed in each application’s Services menu. The drag
// pasteboard is used to transfer data that is being dragged by the user.
// 
// A pasteboard can contain multiple items. You can directly write or read any
// object that implements the [NSPasteboardWriting] or [NSPasteboardReading]
// [Protocol] respectively. This allows you to write and read common items
// such as URLs, colors, images, strings, attributed strings, and sounds
// without an intermediary object. Your custom classes can also implement
// these protocols for use with the pasteboard.
// 
// Writing methods such as [NSPasteboard.SetDataForType] provide a convenient means of
// writing to the first pasteboard item, without having to create the first
// pasteboard item. You can use code like this, for example:
// 
// The general pasteboard, available by way of the [NSPasteboard.GeneralPasteboard] class
// method, automatically participates with the Universal Clipboard feature in
// macOS 10.12 and later and in iOS 10.0 and later. There is no macOS API for
// interacting with this feature.
//
// [Protocol]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/Protocol.html#//apple_ref/doc/uid/TP40008195-CH45
//
// # Creating and releasing a pasteboard
//
//   - [NSPasteboard.ReleaseGlobally]: Releases the receiver’s resources in the pasteboard server.
//
// # Determining pasteboard access
//
//   - [NSPasteboard.AccessBehavior]: The current pasteboard access behavior. The user can customize this behavior per-app in System Settings for any app that has triggered a pasteboard access alert in the past.
//
// # Writing data
//
//   - [NSPasteboard.ClearContents]: Clears the existing contents of the pasteboard.
//   - [NSPasteboard.WriteObjects]: Writes an array of objects to the receiver.
//   - [NSPasteboard.SetDataForType]: Sets the data as the representation for the specified type for the first item on the receiver.
//   - [NSPasteboard.SetPropertyListForType]: Sets the given property list as the representation for the specified type for the first item on the receiver.
//   - [NSPasteboard.SetStringForType]: Sets the given string as the representation for the specified type for the first item on the receiver.
//
// # Reading data
//
//   - [NSPasteboard.ReadObjectsForClassesOptions]: Reads from the receiver objects that best match the specified array of classes.
//   - [NSPasteboard.PasteboardItems]: An array that contains all the items held by the pasteboard.
//   - [NSPasteboard.IndexOfPasteboardItem]: Returns the index of the specified pasteboard item.
//   - [NSPasteboard.DataForType]: Returns the data for the specified type from the first item in the receiver that contains the type.
//   - [NSPasteboard.PropertyListForType]: Returns the property list for the specified type from the first item in the receiver that contains the type.
//   - [NSPasteboard.StringForType]: Returns a concatenation of the strings for the specified type from all the items in the receiver that contain the type.
//
// # Validating contents
//
//   - [NSPasteboard.AvailableTypeFromArray]: Scans the specified types for a type that the receiver supports.
//   - [NSPasteboard.CanReadItemWithDataConformingToTypes]: Returns a Boolean value that indicates whether the receiver contains any items that conform to the specified UTIs.
//   - [NSPasteboard.CanReadObjectForClassesOptions]: Returns a Boolean value that indicates whether the receiver contains any items that can be represented as an instance of any class in a given array.
//   - [NSPasteboard.Types]: An array of the receiver’s supported data types.
//
// # Preparing the pasteboard for content
//
//   - [NSPasteboard.PrepareForNewContentsWithOptions]: Prepares the pasteboard to receive new contents, removing the existing pasteboard contents.
//
// # Getting information about a pasteboard
//
//   - [NSPasteboard.Name]: The receiver’s name.
//   - [NSPasteboard.ChangeCount]: The receiver’s change count.
//
// # Writing data (macOS 10.5 and earlier)
//
//   - [NSPasteboard.DeclareTypesOwner]: Prepares the receiver for a change in its contents by declaring the new types of data it will contain and a new owner.
//   - [NSPasteboard.AddTypesOwner]: Adds promises for the specified types to the first pasteboard item.
//   - [NSPasteboard.WriteFileContents]: Writes the contents of the specified file to the pasteboard.
//   - [NSPasteboard.WriteFileWrapper]: Writes the serialized contents of the specified file wrapper to the pasteboard.
//
// # Reading data (macOS 10.5 and earlier)
//
//   - [NSPasteboard.ReadFileContentsTypeToFile]: Reads data representing a file’s contents from the receiver and writes it to the specified file.
//   - [NSPasteboard.ReadFileWrapper]: Reads data representing a file’s contents from the receiver and returns it as a file wrapper.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard
type NSPasteboard struct {
	objectivec.Object
}

// NSPasteboardFromID constructs a [NSPasteboard] from an objc.ID.
//
// An object that transfers data to and from the pasteboard server.
func NSPasteboardFromID(id objc.ID) NSPasteboard {
	return NSPasteboard{objectivec.Object{ID: id}}
}
// NOTE: NSPasteboard adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPasteboard] class.
//
// # Creating and releasing a pasteboard
//
//   - [INSPasteboard.ReleaseGlobally]: Releases the receiver’s resources in the pasteboard server.
//
// # Determining pasteboard access
//
//   - [INSPasteboard.AccessBehavior]: The current pasteboard access behavior. The user can customize this behavior per-app in System Settings for any app that has triggered a pasteboard access alert in the past.
//
// # Writing data
//
//   - [INSPasteboard.ClearContents]: Clears the existing contents of the pasteboard.
//   - [INSPasteboard.WriteObjects]: Writes an array of objects to the receiver.
//   - [INSPasteboard.SetDataForType]: Sets the data as the representation for the specified type for the first item on the receiver.
//   - [INSPasteboard.SetPropertyListForType]: Sets the given property list as the representation for the specified type for the first item on the receiver.
//   - [INSPasteboard.SetStringForType]: Sets the given string as the representation for the specified type for the first item on the receiver.
//
// # Reading data
//
//   - [INSPasteboard.ReadObjectsForClassesOptions]: Reads from the receiver objects that best match the specified array of classes.
//   - [INSPasteboard.PasteboardItems]: An array that contains all the items held by the pasteboard.
//   - [INSPasteboard.IndexOfPasteboardItem]: Returns the index of the specified pasteboard item.
//   - [INSPasteboard.DataForType]: Returns the data for the specified type from the first item in the receiver that contains the type.
//   - [INSPasteboard.PropertyListForType]: Returns the property list for the specified type from the first item in the receiver that contains the type.
//   - [INSPasteboard.StringForType]: Returns a concatenation of the strings for the specified type from all the items in the receiver that contain the type.
//
// # Validating contents
//
//   - [INSPasteboard.AvailableTypeFromArray]: Scans the specified types for a type that the receiver supports.
//   - [INSPasteboard.CanReadItemWithDataConformingToTypes]: Returns a Boolean value that indicates whether the receiver contains any items that conform to the specified UTIs.
//   - [INSPasteboard.CanReadObjectForClassesOptions]: Returns a Boolean value that indicates whether the receiver contains any items that can be represented as an instance of any class in a given array.
//   - [INSPasteboard.Types]: An array of the receiver’s supported data types.
//
// # Preparing the pasteboard for content
//
//   - [INSPasteboard.PrepareForNewContentsWithOptions]: Prepares the pasteboard to receive new contents, removing the existing pasteboard contents.
//
// # Getting information about a pasteboard
//
//   - [INSPasteboard.Name]: The receiver’s name.
//   - [INSPasteboard.ChangeCount]: The receiver’s change count.
//
// # Writing data (macOS 10.5 and earlier)
//
//   - [INSPasteboard.DeclareTypesOwner]: Prepares the receiver for a change in its contents by declaring the new types of data it will contain and a new owner.
//   - [INSPasteboard.AddTypesOwner]: Adds promises for the specified types to the first pasteboard item.
//   - [INSPasteboard.WriteFileContents]: Writes the contents of the specified file to the pasteboard.
//   - [INSPasteboard.WriteFileWrapper]: Writes the serialized contents of the specified file wrapper to the pasteboard.
//
// # Reading data (macOS 10.5 and earlier)
//
//   - [INSPasteboard.ReadFileContentsTypeToFile]: Reads data representing a file’s contents from the receiver and writes it to the specified file.
//   - [INSPasteboard.ReadFileWrapper]: Reads data representing a file’s contents from the receiver and returns it as a file wrapper.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard
type INSPasteboard interface {
	objectivec.IObject

	// Topic: Creating and releasing a pasteboard

	// Releases the receiver’s resources in the pasteboard server.
	ReleaseGlobally()

	// Topic: Determining pasteboard access

	// The current pasteboard access behavior. The user can customize this behavior per-app in System Settings for any app that has triggered a pasteboard access alert in the past.
	AccessBehavior() NSPasteboardAccessBehavior

	// Topic: Writing data

	// Clears the existing contents of the pasteboard.
	ClearContents() int
	// Writes an array of objects to the receiver.
	WriteObjects(objects []objectivec.IObject) bool
	// Sets the data as the representation for the specified type for the first item on the receiver.
	SetDataForType(data foundation.INSData, dataType NSPasteboardType) bool
	// Sets the given property list as the representation for the specified type for the first item on the receiver.
	SetPropertyListForType(plist objectivec.IObject, dataType NSPasteboardType) bool
	// Sets the given string as the representation for the specified type for the first item on the receiver.
	SetStringForType(string_ string, dataType NSPasteboardType) bool

	// Topic: Reading data

	// Reads from the receiver objects that best match the specified array of classes.
	ReadObjectsForClassesOptions(classArray []objc.Class, options foundation.INSDictionary) foundation.INSArray
	// An array that contains all the items held by the pasteboard.
	PasteboardItems() []NSPasteboardItem
	// Returns the index of the specified pasteboard item.
	IndexOfPasteboardItem(pasteboardItem INSPasteboardItem) uint
	// Returns the data for the specified type from the first item in the receiver that contains the type.
	DataForType(dataType NSPasteboardType) foundation.INSData
	// Returns the property list for the specified type from the first item in the receiver that contains the type.
	PropertyListForType(dataType NSPasteboardType) objectivec.IObject
	// Returns a concatenation of the strings for the specified type from all the items in the receiver that contain the type.
	StringForType(dataType NSPasteboardType) string

	// Topic: Validating contents

	// Scans the specified types for a type that the receiver supports.
	AvailableTypeFromArray(types []string) NSPasteboardType
	// Returns a Boolean value that indicates whether the receiver contains any items that conform to the specified UTIs.
	CanReadItemWithDataConformingToTypes(types []string) bool
	// Returns a Boolean value that indicates whether the receiver contains any items that can be represented as an instance of any class in a given array.
	CanReadObjectForClassesOptions(classArray []objc.Class, options foundation.INSDictionary) bool
	// An array of the receiver’s supported data types.
	Types() []string

	// Topic: Preparing the pasteboard for content

	// Prepares the pasteboard to receive new contents, removing the existing pasteboard contents.
	PrepareForNewContentsWithOptions(options NSPasteboardContentsOptions) int

	// Topic: Getting information about a pasteboard

	// The receiver’s name.
	Name() NSPasteboardName
	// The receiver’s change count.
	ChangeCount() int

	// Topic: Writing data (macOS 10.5 and earlier)

	// Prepares the receiver for a change in its contents by declaring the new types of data it will contain and a new owner.
	DeclareTypesOwner(newTypes []string, newOwner objectivec.IObject) int
	// Adds promises for the specified types to the first pasteboard item.
	AddTypesOwner(newTypes []string, newOwner objectivec.IObject) int
	// Writes the contents of the specified file to the pasteboard.
	WriteFileContents(filename string) bool
	// Writes the serialized contents of the specified file wrapper to the pasteboard.
	WriteFileWrapper(wrapper foundation.NSFileWrapper) bool

	// Topic: Reading data (macOS 10.5 and earlier)

	// Reads data representing a file’s contents from the receiver and writes it to the specified file.
	ReadFileContentsTypeToFile(type_ NSPasteboardType, filename string) string
	// Reads data representing a file’s contents from the receiver and returns it as a file wrapper.
	ReadFileWrapper() foundation.NSFileWrapper

	// An array of calendar events that the data detection system identifies.
	CalendarEvents() objectivec.IObject
	SetCalendarEvents(value objectivec.IObject)
	// The content type of a file that the data detection system identifies when the pasteboard contains a file URL.
	ContentType() uniformtypeidentifiers.UTType
	SetContentType(value uniformtypeidentifiers.UTType)
	// An array of email addresses that the data detection system identifies.
	EmailAddresses() objectivec.IObject
	SetEmailAddresses(value objectivec.IObject)
	// An array of flight numbers that the data detection system identifies.
	FlightNumbers() objectivec.IObject
	SetFlightNumbers(value objectivec.IObject)
	// An array of web links that the data detection system identifies.
	Links() objectivec.IObject
	SetLinks(value objectivec.IObject)
	// A set of key paths that represent metadata types that the data detection system identifies.
	MetadataTypes() objectivec.IObject
	SetMetadataTypes(value objectivec.IObject)
	// An array of money amounts and currencies that the data detection system identifies.
	MoneyAmounts() objectivec.IObject
	SetMoneyAmounts(value objectivec.IObject)
	// A number that the data detection system identifies.
	Number() float64
	SetNumber(value float64)
	// A set of key paths that represent patterns that the data detection system identifies.
	Patterns() objectivec.IObject
	SetPatterns(value objectivec.IObject)
	// An array of phone numbers that the data detection system identifies.
	PhoneNumbers() objectivec.IObject
	SetPhoneNumbers(value objectivec.IObject)
	// An array of postal addresses that the data detection system identifies.
	PostalAddresses() objectivec.IObject
	SetPostalAddresses(value objectivec.IObject)
	// A string that the data detection system identifies as a probable web search item, suitable for implementing “Paste and Search”.
	ProbableWebSearch() string
	SetProbableWebSearch(value string)
	// A string that the data detection system identifies as a probable web URL, suitable for implementing “Paste and Go”.
	ProbableWebURL() string
	SetProbableWebURL(value string)
	// An array of parcel tracking numbers and carriers that the data detection system identifies.
	ShipmentTrackingNumbers() objectivec.IObject
	SetShipmentTrackingNumbers(value objectivec.IObject)
}

// Init initializes the instance.
func (p NSPasteboard) Init() NSPasteboard {
	rv := objc.Send[NSPasteboard](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPasteboard) Autorelease() NSPasteboard {
	rv := objc.Send[NSPasteboard](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPasteboard creates a new NSPasteboard instance.
func NewNSPasteboard() NSPasteboard {
	class := getNSPasteboardClass()
	rv := objc.Send[NSPasteboard](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new pasteboard object that supplies the specified data in as many
// types as possible based on the available filter services.
//
// data: The data to be placed on the pasteboard.
//
// type: The type of data in the `data` parameter.
//
// # Return Value
// 
// The new pasteboard object.
//
// # Discussion
// 
// The returned pasteboard also declares data of the supplied `type`.
// 
// No filter service is invoked until the data is actually requested, so
// invoking this method is reasonably inexpensive.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/init(byFilteringData:ofType:)
func NewPasteboardByFilteringDataOfType(data foundation.INSData, type_ NSPasteboardType) NSPasteboard {
	rv := objc.Send[objc.ID](objc.ID(getNSPasteboardClass().class), objc.Sel("pasteboardByFilteringData:ofType:"), data, objc.String(string(type_)))
	return NSPasteboardFromID(rv)
}

// Creates a new pasteboard object that supplies the specified file in as many
// types as possible based on the available filter services.
//
// filename: The filename to put on the pasteboard.
//
// # Return Value
// 
// The new pasteboard object.
//
// # Discussion
// 
// No filter service is invoked until the data is actually requested, so
// invoking this method is reasonably inexpensive.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/init(byFilteringFile:)
func NewPasteboardByFilteringFile(filename string) NSPasteboard {
	rv := objc.Send[objc.ID](objc.ID(getNSPasteboardClass().class), objc.Sel("pasteboardByFilteringFile:"), objc.String(filename))
	return NSPasteboardFromID(rv)
}

// Creates a new pasteboard object that supplies the specified pasteboard data
// in as many types as possible based on the available filter services.
//
// pboard: The original pasteboard object.
//
// # Return Value
// 
// The new pasteboard object. This method returns the object in the
// `pasteboard` parameter if the pasteboard was returned by one of the
// `pasteboardByFiltering...` methods. This prevents a pasteboard from being
// expanded multiple times.
//
// # Discussion
// 
// This process can be thought of as expanding the pasteboard, because the new
// pasteboard generally contains more representations of the data than
// `pasteboard`.
// 
// This method only returns the original types and the types that can be
// created as a result of a single filter; the pasteboard does not have
// defined types that are the result of translation by multiple filters.
// 
// No filter service is invoked until the data is actually requested, so
// invoking this method is reasonably inexpensive.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/init(byFilteringTypesInPasteboard:)
func NewPasteboardByFilteringTypesInPasteboard(pboard INSPasteboard) NSPasteboard {
	rv := objc.Send[objc.ID](objc.ID(getNSPasteboardClass().class), objc.Sel("pasteboardByFilteringTypesInPasteboard:"), pboard)
	return NSPasteboardFromID(rv)
}

// Returns the pasteboard with the specified name.
//
// name: The name of the pasteboard. The names of standard pasteboards are given in
// `Pasteboard Names`.
//
// # Return Value
// 
// The pasteboard associated with the given name, or a new [NSPasteboard]
// object if the application does not yet have a pasteboard object for the
// specified name.
//
// # Discussion
// 
// Other names can be assigned to create private pasteboards for other
// purposes.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/init(name:)
func NewPasteboardWithName(name NSPasteboardName) NSPasteboard {
	rv := objc.Send[objc.ID](objc.ID(getNSPasteboardClass().class), objc.Sel("pasteboardWithName:"), objc.String(string(name)))
	return NSPasteboardFromID(rv)
}

// Releases the receiver’s resources in the pasteboard server.
//
// # Discussion
// 
// After this method is invoked, no other application can use the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/releaseGlobally()
func (p NSPasteboard) ReleaseGlobally() {
	objc.Send[objc.ID](p.ID, objc.Sel("releaseGlobally"))
}
// Clears the existing contents of the pasteboard.
//
// # Return Value
// 
// The change count of the receiver.
//
// # Discussion
// 
// Clears the existing contents of the pasteboard, preparing it for new
// contents. This is the first step in providing data on the pasteboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/clearContents()
func (p NSPasteboard) ClearContents() int {
	rv := objc.Send[int](p.ID, objc.Sel("clearContents"))
	return rv
}
// Writes an array of objects to the receiver.
//
// objects: An array of objects that implement the [NSPasteboardWriting] protocol
// (including instances of [NSPasteboardItem]).
//
// # Return Value
// 
// [true] if the array was successfully added, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/writeObjects(_:)
func (p NSPasteboard) WriteObjects(objects []objectivec.IObject) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("writeObjects:"), objectivec.IObjectSliceToNSArray(objects))
	return rv
}
// Sets the data as the representation for the specified type for the first
// item on the receiver.
//
// data: The data to write to the pasteboard.
//
// dataType: The type of data in the `data` parameter. The type must have been declared
// by a previous [DeclareTypesOwner] message.
//
// # Return Value
// 
// [true] if the data was written successfully, otherwise [false] if ownership
// of the pasteboard has changed. Any other error raises an
// [NSPasteboardCommunicationException].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/setData(_:forType:)
func (p NSPasteboard) SetDataForType(data foundation.INSData, dataType NSPasteboardType) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("setData:forType:"), data, objc.String(string(dataType)))
	return rv
}
// Sets the given property list as the representation for the specified type
// for the first item on the receiver.
//
// plist: The property list data to write to the pasteboard.
//
// dataType: The type of property-list data in the `propertyList` parameter. The type
// must have been declared by a previous [DeclareTypesOwner] message.
//
// # Return Value
// 
// [true] if the data was written successfully, otherwise [false] if ownership
// of the pasteboard has changed. Any other error raises an
// [NSPasteboardCommunicationException].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method invokes [SetDataForType] with a serialized property list
// parameter.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/setPropertyList(_:forType:)
func (p NSPasteboard) SetPropertyListForType(plist objectivec.IObject, dataType NSPasteboardType) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("setPropertyList:forType:"), plist, objc.String(string(dataType)))
	return rv
}
// Sets the given string as the representation for the specified type for the
// first item on the receiver.
//
// string: The string to write to the pasteboard.
//
// dataType: The type of string data. The type must have been declared by a previous
// [DeclareTypesOwner] message.
//
// # Return Value
// 
// [true] if the data was written successfully, otherwise [false] if ownership
// of the pasteboard has changed. Any other error raises an
// [NSPasteboardCommunicationException].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method invokes [SetDataForType] to perform the write.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/setString(_:forType:)
func (p NSPasteboard) SetStringForType(string_ string, dataType NSPasteboardType) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("setString:forType:"), objc.String(string_), objc.String(string(dataType)))
	return rv
}
// Reads from the receiver objects that best match the specified array of
// classes.
//
// classArray: An array of class objects.
// 
// Because this method creates an instance of the first class that can read a
// given pasteboard item, you can order the classes in `classArray` to match
// your preferred order of representation. Classes in the array must conform
// to the [NSPasteboardReading] protocol.
//
// options: A dictionary that specifies options to refine the search for pasteboard
// items, for example to restrict the search to file URLs with particular
// content types. For valid dictionary keys, see `Pasteboard Reading Options`.
//
// # Return Value
// 
// An array containing the best match (if any) for each of the items on the
// receiver that can be represented by a class specified in `classArray`.
// Returns `nil` if there is an error in retrieving the requested items from
// the pasteboard, or an empty array if no objects of the specified classes
// can be created.
//
// # Discussion
// 
// Classes in `classArray` must implement the [NSPasteboardReading] protocol.
// Cocoa classes that implement this protocol include [NSImage], [NSString],
// [NSURL], [NSColor], [NSAttributedString], and [NSPasteboardItem]. For every
// item on the pasteboard, each class in the provided array will be queried
// for the types it can read using [ReadableTypesForPasteboard]. An instance
// is created of the first class found in the provided array whose readable
// types match a conforming type contained in that pasteboard item. Any
// instances that could be created from pasteboard item data is returned to
// the caller. Additional options, such as restricting the search to file URLs
// with particular content types, can be specified with an options dictionary.
// 
// Only objects of the requested classes are returned. You can always ensure
// to receive one object per item on the pasteboard by including the
// [NSPasteboardItem] class in the array of classes.
// 
// Consider the following example: there are five items on the pasteboard, two
// contain TIFF data, two contain RTF data, one contains a private data type.
// The following table shows what objects you get back in the returned array
// for different classes in `classArray`.
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/readObjects(forClasses:options:)
func (p NSPasteboard) ReadObjectsForClassesOptions(classArray []objc.Class, options foundation.INSDictionary) foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("readObjectsForClasses:options:"), objectivec.ClassSliceToNSArray(classArray), options)
	return foundation.NSArrayFromID(rv)
}
// Returns the index of the specified pasteboard item.
//
// pasteboardItem: A pasteboard item.
//
// # Return Value
// 
// The index of the specified pasteboard item. If `pasteboardItem` has not
// been added to any pasteboard, or is owned by another pasteboard, returns
// [NSNotFound].
//
// # Discussion
// 
// An item’s index in the pasteboard is useful for a pasteboard item data
// provider that has promised data for multiple items, to be able to easily
// match the pasteboard item to an array of source data from which to derive
// the promised data.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/index(of:)
func (p NSPasteboard) IndexOfPasteboardItem(pasteboardItem INSPasteboardItem) uint {
	rv := objc.Send[uint](p.ID, objc.Sel("indexOfPasteboardItem:"), pasteboardItem)
	return rv
}
// Returns the data for the specified type from the first item in the receiver
// that contains the type.
//
// dataType: The type of data you want to read from the pasteboard. This value should be
// one of the types returned by [Types] or [AvailableTypeFromArray].
//
// # Return Value
// 
// A data object containing the data for the specified type from the first
// item in the receiver that contains the type, or `nil` if the contents of
// the pasteboard changed since they were last checked.
//
// # Discussion
// 
// This method may also return `nil` if the pasteboard server cannot supply
// the data in time—for example, if the pasteboard’s owner is slow in
// responding to a [pasteboard:provideDataForType:] message and the
// interprocess communication times out.
// 
// # Discussion
// 
// Errors other than a timeout raise an [NSPasteboardCommunicationException].
// 
// If `nil` is returned, the application should put up a panel informing the
// user that it was unable to carry out the paste operation. Note that sending
// [Types] or [AvailableTypeFromArray] before invoking [DataForType] can help
// you determine whether a `nil` result from a reading method is due to
// something like a pasteboard timeout.
// 
// # Special Considerations
// 
// For standard text data types such as string, RTF, and RTFD, the text data
// from each item is returned as one combined result separated by newlines.
//
// [pasteboard:provideDataForType:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/pasteboard:provideDataForType:
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/data(forType:)
func (p NSPasteboard) DataForType(dataType NSPasteboardType) foundation.INSData {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("dataForType:"), objc.String(string(dataType)))
	return foundation.NSDataFromID(rv)
}
// Returns the property list for the specified type from the first item in the
// receiver that contains the type.
//
// dataType: The pasteboard data type containing the property-list data.
//
// # Return Value
// 
// A property list of objects of the specified type, obtained from the first
// item in the receiver that contains the type. The returned property list can
// contain any combination of objects, as long as each object is a valid
// property-list type (for a list of types, see [Property list]).
//
// [Property list]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/PropertyList.html#//apple_ref/doc/uid/TP40008195-CH44
//
// # Discussion
// 
// This method invokes the [DataForType] method.
// 
// # Special Considerations
// 
// It’s a good idea to check [Types] or call [AvailableTypeFromArray] before
// invoking [PropertyListForType]. Although performing this check isn’t
// required, doing so can help you determine if a `nil` result from a reading
// method is due to something like a pasteboard timeout.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/propertyList(forType:)
func (p NSPasteboard) PropertyListForType(dataType NSPasteboardType) objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("propertyListForType:"), objc.String(string(dataType)))
	return objectivec.Object{ID: rv}
}
// Returns a concatenation of the strings for the specified type from all the
// items in the receiver that contain the type.
//
// dataType: The pasteboard data type to read.
//
// # Return Value
// 
// A concatenation of the strings for the specified type from all the items in
// the receiver that contain the type, or `nil` if none of the items contain
// strings of the specified type.
//
// # Discussion
// 
// This method invokes [DataForType] to obtain the string. If the string
// cannot be obtained, [StringForType] returns `nil`. See [DataForType] for a
// description of what will cause `nil` to be returned.
// 
// In macOS 10.6 and later, if the receiver contains multiple items that can
// provide string, RTF, or RTFD data, the text data from each item is returned
// as a combined result separated by newlines.
// 
// # Special Considerations
// 
// It’s a good idea to check [Types] or call [AvailableTypeFromArray] before
// invoking [StringForType]. Although performing this check isn’t required,
// doing so can help you determine if a `nil` result from a reading method is
// due to something like a pasteboard timeout.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/string(forType:)
func (p NSPasteboard) StringForType(dataType NSPasteboardType) string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("stringForType:"), objc.String(string(dataType)))
	return foundation.NSStringFromID(rv).String()
}
// Scans the specified types for a type that the receiver supports.
//
// types: An array of [NSString] objects specifying the pasteboard types your
// application supports, in preferred order.
//
// # Return Value
// 
// The first pasteboard type in `types` that is available on the pasteboard,
// or `nil` if the receiver does not contain any of the types in `types`.
//
// # Discussion
// 
// You use this method to determine the best representation available on the
// pasteboard. For example, if your application supports RTFD, RTF, and string
// data, then you might invoke the method as follows:
// 
// If the pasteboard contains RTF and string data, then `bestType` would
// contain [NSRTFPboardType]. If the pasteboard contains none of the types in
// `supportedTypes`, then `bestType` would be `nil`.
// 
// You must send a [Types] or [AvailableTypeFromArray] message before reading
// any data from an [NSPasteboard] object. If you need to see if a type in the
// returned array matches a type string you have stored locally, use the
// [isEqual(to:)] method to perform the comparison.
//
// [isEqual(to:)]: https://developer.apple.com/documentation/Foundation/NSString/isEqual(to:)
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/availableType(from:)
func (p NSPasteboard) AvailableTypeFromArray(types []string) NSPasteboardType {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("availableTypeFromArray:"), objectivec.StringSliceToNSArray(types))
	return NSPasteboardType(foundation.NSStringFromID(rv).String())
}
// Returns a Boolean value that indicates whether the receiver contains any
// items that conform to the specified UTIs.
//
// types: An array of [NSString] objects containing UTIs.
//
// # Return Value
// 
// [true] if the receiver contains any items that conform to the UTIs
// specified in `types`, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/canReadItem(withDataConformingToTypes:)
func (p NSPasteboard) CanReadItemWithDataConformingToTypes(types []string) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canReadItemWithDataConformingToTypes:"), objectivec.StringSliceToNSArray(types))
	return rv
}
// Returns a Boolean value that indicates whether the receiver contains any
// items that can be represented as an instance of any class in a given array.
//
// classArray: An array of class objects.
// 
// Classes in the array must conform to the [NSPasteboardReading] protocol.
//
// options: A dictionary that specifies options to refine the search for pasteboard
// items, for example to restrict the search to file URLs with particular
// content types. For valid dictionary keys, see `Pasteboard Reading Options`.
//
// # Return Value
// 
// [true] if the receiver contains any items that can be represented as an
// instance of a class specified in `classArray`, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/canReadObject(forClasses:options:)
func (p NSPasteboard) CanReadObjectForClassesOptions(classArray []objc.Class, options foundation.INSDictionary) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canReadObjectForClasses:options:"), objectivec.ClassSliceToNSArray(classArray), options)
	return rv
}
// Prepares the pasteboard to receive new contents, removing the existing
// pasteboard contents.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/prepareForNewContents(with:)
func (p NSPasteboard) PrepareForNewContentsWithOptions(options NSPasteboardContentsOptions) int {
	rv := objc.Send[int](p.ID, objc.Sel("prepareForNewContentsWithOptions:"), options)
	return rv
}
// Prepares the receiver for a change in its contents by declaring the new
// types of data it will contain and a new owner.
//
// newTypes: An array of [NSString] objects that specify the types of data that may be
// added to the new pasteboard. The types should be ordered according to the
// preference of the source application, with the most preferred type coming
// first (typically, the richest representation).
//
// newOwner: The object responsible for writing data to the pasteboard, or `nil` if you
// provide data for all types immediately. If you specify a `newOwner` object,
// it must support all of the types declared in the `newTypes` parameter and
// must remain alive for as long as the data is on the pasteboard.
//
// # Return Value
// 
// The receiver’s new change count.
//
// # Discussion
// 
// This method is the equivalent of invoking [ClearContents], implicitly
// writing the first pasteboard item, and then calling [AddTypesOwner] to
// promise types for the first pasteboard item.
// 
// # Special Considerations
// 
// In general, you should not use this method with [WriteObjects], since
// [WriteObjects] will always write additional items to the pasteboard, and
// will not affect items already on the pasteboard, including the item
// implicitly created by this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/declareTypes(_:owner:)
func (p NSPasteboard) DeclareTypesOwner(newTypes []string, newOwner objectivec.IObject) int {
	rv := objc.Send[int](p.ID, objc.Sel("declareTypes:owner:"), objectivec.StringSliceToNSArray(newTypes), newOwner)
	return rv
}
// Adds promises for the specified types to the first pasteboard item.
//
// newTypes: An array of [NSString] objects, each of which specifies a type of data that
// can be provided to the pasteboard.
//
// newOwner: The object that provides the data for the specified types.
// 
// If the data for those types is provided immediately, the owner can be
// `nil`. If the data for the added types will be provided lazily when
// requested from the pasteboard, an owner object must be provided that
// implements the -[pasteboard:provideDataForType:] method of the
// [NSPasteboardOwner] informal protocol.
// //
// [pasteboard:provideDataForType:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/pasteboard:provideDataForType:
//
// # Return Value
// 
// The new change count, or `0` if there was an error adding the types.
//
// # Discussion
// 
// This method adds promises for the specified types to the first pasteboard
// item.
// 
// You use this methods to declare additional types of data for the first
// pasteboard item in the receiver. You can also use it to replace existing
// types added by a previous [DeclareTypesOwner] or [AddTypesOwner] message.
// 
// The `newTypes` parameter specifies the types of data you are promising to
// the pasteboard. The types should be ordered according to the preference of
// the source application, with the most preferred type coming first
// (typically, the richest representation). New types are added to the end of
// the list containing any existing types, if any.
// 
// If you specify a type that has already been declared, this method replaces
// the owner of that type with the value in `newOwner`. In addition, any data
// already written to the pasteboard for that type is removed.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/addTypes(_:owner:)
func (p NSPasteboard) AddTypesOwner(newTypes []string, newOwner objectivec.IObject) int {
	rv := objc.Send[int](p.ID, objc.Sel("addTypes:owner:"), objectivec.StringSliceToNSArray(newTypes), newOwner)
	return rv
}
// Writes the contents of the specified file to the pasteboard.
//
// filename: The name of the file to write to the pasteboard.
//
// # Return Value
// 
// [true] if the data was successfully written, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Writes the contents of the file `filename` to the receiver and declares the
// data to be of type [NSFileContentsPboardType] and also of a type
// appropriate for the file’s extension (as returned by the
// [NSCreateFileContentsPboardType] function when passed the files extension),
// if it has one.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/writeFileContents(_:)
func (p NSPasteboard) WriteFileContents(filename string) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("writeFileContents:"), objc.String(filename))
	return rv
}
// Writes the serialized contents of the specified file wrapper to the
// pasteboard.
//
// wrapper: The file wrapper to write to the pasteboard.
//
// # Return Value
// 
// [true] if the data was successfully written, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Writes the serialized contents of the file wrapper `wrapper` to the
// receiver and declares the data to be of type [NSFileContentsPboardType] and
// also of a type appropriate for the file’s extension (as returned by the
// [NSCreateFileContentsPboardType] function when passed the files extension),
// if it has one. If `wrapper` does not have a preferred filename, this method
// raises an exception.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/write(_:)
func (p NSPasteboard) WriteFileWrapper(wrapper foundation.NSFileWrapper) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("writeFileWrapper:"), wrapper)
	return rv
}
// Reads data representing a file’s contents from the receiver and writes it
// to the specified file.
//
// type: The pasteboard data type to read. You should generally specify a value for
// this parameter. If you specify `nil`, the filename extension (in
// combination with the [NSCreateFileContentsPboardType] function) is used to
// determine the type.
//
// filename: The file to receive the pasteboard data.
//
// # Return Value
// 
// The name of the file into which the data was actually written.
//
// # Discussion
// 
// Data of any file contents type should only be read using this method. If
// data matching the specified type is not found on the pasteboard, data of
// type [NSFileContentsPboardType] is requested.
// 
// # Special Considerations
// 
// You must send an [AvailableTypeFromArray] or [Types] message before
// invoking [ReadFileContentsTypeToFile].
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/readFileContentsType(_:toFile:)
func (p NSPasteboard) ReadFileContentsTypeToFile(type_ NSPasteboardType, filename string) string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("readFileContentsType:toFile:"), objc.String(string(type_)), objc.String(filename))
	return foundation.NSStringFromID(rv).String()
}
// Reads data representing a file’s contents from the receiver and returns
// it as a file wrapper.
//
// # Return Value
// 
// A file wrapper containing the pasteboard data, or `nil` if the receiver
// contained no data of type [NSFileContentsPboardType].
//
// # Discussion
// 
// In macOS 10.5 and earlier, the file contents pboard type allowed you to
// synthesize a pboard type for a file’s contents based on the file’s
// extension. In macOS 10.5 and later, using the UTI of a file to represent
// its contents now replaces this functionality.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/readFileWrapper()
func (p NSPasteboard) ReadFileWrapper() foundation.NSFileWrapper {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("readFileWrapper"))
	return foundation.NSFileWrapperFromID(rv)
}

// Creates and returns a new pasteboard with a name that is guaranteed to be
// unique with respect to other pasteboards in the system.
//
// # Return Value
// 
// The new pasteboard object.
//
// # Discussion
// 
// This method is useful for apps that implement their own interprocess
// communication using pasteboards. Because the lifetime of a unique
// pasteboard is not related to the lifetime of the creating app, you must
// release a unique pasteboard by calling [ReleaseGlobally] to avoid possible
// leaks.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/withUniqueName()
func (_NSPasteboardClass NSPasteboardClass) PasteboardWithUniqueName() NSPasteboard {
	rv := objc.Send[objc.ID](objc.ID(_NSPasteboardClass.class), objc.Sel("pasteboardWithUniqueName"))
	return NSPasteboardFromID(rv)
}
// Returns the data types that can be converted to the specified type using
// the available filter services.
//
// type: The target data type.
//
// # Return Value
// 
// An array of [NSString] objects containing the types that can be converted
// to the target data type.
//
// # Discussion
// 
// The array also contains the original type.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/types(filterableTo:)
func (_NSPasteboardClass NSPasteboardClass) TypesFilterableTo(type_ NSPasteboardType) []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSPasteboardClass.class), objc.Sel("typesFilterableTo:"), objc.String(string(type_)))
	return objc.ConvertSliceToStrings(rv)
}

// The current pasteboard access behavior. The user can customize this
// behavior per-app in System Settings for any app that has triggered a
// pasteboard access alert in the past.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/accessBehavior-86972
func (p NSPasteboard) AccessBehavior() NSPasteboardAccessBehavior {
	rv := objc.Send[NSPasteboardAccessBehavior](p.ID, objc.Sel("accessBehavior"))
	return NSPasteboardAccessBehavior(rv)
}
// An array that contains all the items held by the pasteboard.
//
// # Discussion
// 
// If an error occurs when retrieving the pasteboard items, the value of this
// property is `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/pasteboardItems
func (p NSPasteboard) PasteboardItems() []NSPasteboardItem {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("pasteboardItems"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSPasteboardItem {
		return NSPasteboardItemFromID(id)
	})
}
// An array of the receiver’s supported data types.
//
// # Discussion
// 
// The [Types] array is an array of [NSString] objects containing the union of
// the types of data declared for all the pasteboard items on the receiver.
// The returned types are listed in the order they were declared. It’s a
// good idea to check the value of [Types] (or call [AvailableTypeFromArray])
// before reading any data from an [NSPasteboard] object. If you need to see
// if a type in the [Types] array matches a type string you have stored
// locally, use the [isEqual(to:)] method to perform the comparison.
//
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [isEqual(to:)]: https://developer.apple.com/documentation/Foundation/NSString/isEqual(to:)
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/types
func (p NSPasteboard) Types() []string {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("types"))
	return objc.ConvertSliceToStrings(rv)
}
// The receiver’s name.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/name-swift.property
func (p NSPasteboard) Name() NSPasteboardName {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("name"))
	return NSPasteboardName(foundation.NSStringFromID(rv).String())
}
// The receiver’s change count.
//
// # Discussion
// 
// The change count starts at zero when a client creates the receiver and
// becomes the first owner. The change count subsequently increments each time
// the pasteboard ownership changes.
// 
// The change count is also returned from [ClearContents] and
// [DeclareTypesOwner]. You can therefore record the value of `changeCount` at
// the time that you take ownership of the pasteboard and compare it with a
// later value to determine whether you still have ownership.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/changeCount
func (p NSPasteboard) ChangeCount() int {
	rv := objc.Send[int](p.ID, objc.Sel("changeCount"))
	return rv
}
// An array of calendar events that the data detection system identifies.
//
// See: https://developer.apple.com/documentation/appkit/nspasteboard/detectedvalues/calendarevents
func (p NSPasteboard) CalendarEvents() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("calendarEvents"))
	return objectivec.Object{ID: rv}
}
func (p NSPasteboard) SetCalendarEvents(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setCalendarEvents:"), value)
}
// The content type of a file that the data detection system identifies when
// the pasteboard contains a file URL.
//
// See: https://developer.apple.com/documentation/appkit/nspasteboard/detectedmetadata/contenttype
func (p NSPasteboard) ContentType() uniformtypeidentifiers.UTType {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("contentType"))
	return uniformtypeidentifiers.UTTypeFromID(objc.ID(rv))
}
func (p NSPasteboard) SetContentType(value uniformtypeidentifiers.UTType) {
	objc.Send[struct{}](p.ID, objc.Sel("setContentType:"), value)
}
// An array of email addresses that the data detection system identifies.
//
// See: https://developer.apple.com/documentation/appkit/nspasteboard/detectedvalues/emailaddresses
func (p NSPasteboard) EmailAddresses() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("emailAddresses"))
	return objectivec.Object{ID: rv}
}
func (p NSPasteboard) SetEmailAddresses(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setEmailAddresses:"), value)
}
// An array of flight numbers that the data detection system identifies.
//
// See: https://developer.apple.com/documentation/appkit/nspasteboard/detectedvalues/flightnumbers
func (p NSPasteboard) FlightNumbers() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("flightNumbers"))
	return objectivec.Object{ID: rv}
}
func (p NSPasteboard) SetFlightNumbers(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setFlightNumbers:"), value)
}
// An array of web links that the data detection system identifies.
//
// See: https://developer.apple.com/documentation/appkit/nspasteboard/detectedvalues/links
func (p NSPasteboard) Links() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("links"))
	return objectivec.Object{ID: rv}
}
func (p NSPasteboard) SetLinks(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setLinks:"), value)
}
// A set of key paths that represent metadata types that the data detection
// system identifies.
//
// See: https://developer.apple.com/documentation/appkit/nspasteboard/detectedmetadata/metadatatypes
func (p NSPasteboard) MetadataTypes() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("metadataTypes"))
	return objectivec.Object{ID: rv}
}
func (p NSPasteboard) SetMetadataTypes(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setMetadataTypes:"), value)
}
// An array of money amounts and currencies that the data detection system
// identifies.
//
// See: https://developer.apple.com/documentation/appkit/nspasteboard/detectedvalues/moneyamounts
func (p NSPasteboard) MoneyAmounts() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("moneyAmounts"))
	return objectivec.Object{ID: rv}
}
func (p NSPasteboard) SetMoneyAmounts(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setMoneyAmounts:"), value)
}
// A number that the data detection system identifies.
//
// See: https://developer.apple.com/documentation/appkit/nspasteboard/detectedvalues/number
func (p NSPasteboard) Number() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("number"))
	return rv
}
func (p NSPasteboard) SetNumber(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setNumber:"), value)
}
// A set of key paths that represent patterns that the data detection system
// identifies.
//
// See: https://developer.apple.com/documentation/appkit/nspasteboard/detectedvalues/patterns
func (p NSPasteboard) Patterns() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("patterns"))
	return objectivec.Object{ID: rv}
}
func (p NSPasteboard) SetPatterns(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setPatterns:"), value)
}
// An array of phone numbers that the data detection system identifies.
//
// See: https://developer.apple.com/documentation/appkit/nspasteboard/detectedvalues/phonenumbers
func (p NSPasteboard) PhoneNumbers() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("phoneNumbers"))
	return objectivec.Object{ID: rv}
}
func (p NSPasteboard) SetPhoneNumbers(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setPhoneNumbers:"), value)
}
// An array of postal addresses that the data detection system identifies.
//
// See: https://developer.apple.com/documentation/appkit/nspasteboard/detectedvalues/postaladdresses
func (p NSPasteboard) PostalAddresses() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("postalAddresses"))
	return objectivec.Object{ID: rv}
}
func (p NSPasteboard) SetPostalAddresses(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setPostalAddresses:"), value)
}
// A string that the data detection system identifies as a probable web search
// item, suitable for implementing “Paste and Search”.
//
// See: https://developer.apple.com/documentation/appkit/nspasteboard/detectedvalues/probablewebsearch
func (p NSPasteboard) ProbableWebSearch() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("probableWebSearch"))
	return foundation.NSStringFromID(rv).String()
}
func (p NSPasteboard) SetProbableWebSearch(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setProbableWebSearch:"), objc.String(value))
}
// A string that the data detection system identifies as a probable web URL,
// suitable for implementing “Paste and Go”.
//
// See: https://developer.apple.com/documentation/appkit/nspasteboard/detectedvalues/probableweburl
func (p NSPasteboard) ProbableWebURL() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("probableWebURL"))
	return foundation.NSStringFromID(rv).String()
}
func (p NSPasteboard) SetProbableWebURL(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setProbableWebURL:"), objc.String(value))
}
// An array of parcel tracking numbers and carriers that the data detection
// system identifies.
//
// See: https://developer.apple.com/documentation/appkit/nspasteboard/detectedvalues/shipmenttrackingnumbers
func (p NSPasteboard) ShipmentTrackingNumbers() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("shipmentTrackingNumbers"))
	return objectivec.Object{ID: rv}
}
func (p NSPasteboard) SetShipmentTrackingNumbers(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setShipmentTrackingNumbers:"), value)
}

// The shared pasteboard object to use for general content.
//
// # Return Value
// 
// The general pasteboard.
// 
// # Discussion
// 
// Invokes [PasteboardWithName] to obtain the pasteboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/general
func (_NSPasteboardClass NSPasteboardClass) GeneralPasteboard() NSPasteboard {
	rv := objc.Send[objc.ID](objc.ID(_NSPasteboardClass.class), objc.Sel("generalPasteboard"))
	return NSPasteboardFromID(objc.ID(rv))
}

