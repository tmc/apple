// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PropertyListSerialization] class.
var (
	_PropertyListSerializationClass     PropertyListSerializationClass
	_PropertyListSerializationClassOnce sync.Once
)

func getPropertyListSerializationClass() PropertyListSerializationClass {
	_PropertyListSerializationClassOnce.Do(func() {
		_PropertyListSerializationClass = PropertyListSerializationClass{class: objc.GetClass("NSPropertyListSerialization")}
	})
	return _PropertyListSerializationClass
}

// GetPropertyListSerializationClass returns the class object for NSPropertyListSerialization.
func GetPropertyListSerializationClass() PropertyListSerializationClass {
	return getPropertyListSerializationClass()
}

type PropertyListSerializationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (pc PropertyListSerializationClass) Alloc() PropertyListSerialization {
	rv := objc.Send[PropertyListSerialization](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}







// An object that converts between a property list and one of several
// serialized representations.
//
// # Overview
// 
// The [NSPropertyListSerialization] class provides methods that convert a
// property list to and from several serialized formats. A property list is
// itself an array or dictionary that contains only [NSData], [NSString],
// [NSArray], [NSDictionary], [NSDate], and [NSNumber] objects.
// 
// Property list objects are toll-free bridged with their respective Core
// Foundation types ([CFData], [CFString], and so on). See [Toll-Free
// Bridging] for more information on toll-free bridging.
//
// [CFData]: https://developer.apple.com/documentation/CoreFoundation/CFData
// [CFString]: https://developer.apple.com/documentation/CoreFoundation/CFString
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
//
// # Error Codes
//
//   - [PropertyListSerialization.NSPropertyListReadCorruptError]: Parsing of the property list failed.
//   - [PropertyListSerialization.SetNSPropertyListReadCorruptError]
//   - [PropertyListSerialization.NSPropertyListReadUnknownVersionError]: The version number of the property list cannot be determined.
//   - [PropertyListSerialization.SetNSPropertyListReadUnknownVersionError]
//   - [PropertyListSerialization.NSPropertyListReadStreamError]: Reading of the property list failed.
//   - [PropertyListSerialization.SetNSPropertyListReadStreamError]
//   - [PropertyListSerialization.NSPropertyListWriteStreamError]: Writing to the property list failed.
//   - [PropertyListSerialization.SetNSPropertyListWriteStreamError]
//   - [PropertyListSerialization.NSPropertyListWriteInvalidError]: Writing failed because of an invalid property list object, or an invalid property list type was specified.
//   - [PropertyListSerialization.SetNSPropertyListWriteInvalidError]
//   - [PropertyListSerialization.NSPropertyListErrorMinimum]: The start of the range of error codes reserved for property list errors.
//   - [PropertyListSerialization.SetNSPropertyListErrorMinimum]
//   - [PropertyListSerialization.NSPropertyListErrorMaximum]: The end of the range of error codes reserved for property list errors.
//   - [PropertyListSerialization.SetNSPropertyListErrorMaximum]
//
// See: https://developer.apple.com/documentation/Foundation/PropertyListSerialization
type PropertyListSerialization struct {
	objectivec.Object
}

// PropertyListSerializationFromID constructs a [PropertyListSerialization] from an objc.ID.
//
// An object that converts between a property list and one of several
// serialized representations.
func PropertyListSerializationFromID(id objc.ID) PropertyListSerialization {
	return NSPropertyListSerialization{objectivec.Object{id}}
}

// NSPropertyListSerializationFromID is an alias for [PropertyListSerializationFromID] for cross-framework compatibility.
func NSPropertyListSerializationFromID(id objc.ID) PropertyListSerialization { return PropertyListSerializationFromID(id) }
// NOTE: PropertyListSerialization adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [PropertyListSerialization] class.
//
// # Error Codes
//
//   - [IPropertyListSerialization.NSPropertyListReadCorruptError]: Parsing of the property list failed.
//   - [IPropertyListSerialization.SetNSPropertyListReadCorruptError]
//   - [IPropertyListSerialization.NSPropertyListReadUnknownVersionError]: The version number of the property list cannot be determined.
//   - [IPropertyListSerialization.SetNSPropertyListReadUnknownVersionError]
//   - [IPropertyListSerialization.NSPropertyListReadStreamError]: Reading of the property list failed.
//   - [IPropertyListSerialization.SetNSPropertyListReadStreamError]
//   - [IPropertyListSerialization.NSPropertyListWriteStreamError]: Writing to the property list failed.
//   - [IPropertyListSerialization.SetNSPropertyListWriteStreamError]
//   - [IPropertyListSerialization.NSPropertyListWriteInvalidError]: Writing failed because of an invalid property list object, or an invalid property list type was specified.
//   - [IPropertyListSerialization.SetNSPropertyListWriteInvalidError]
//   - [IPropertyListSerialization.NSPropertyListErrorMinimum]: The start of the range of error codes reserved for property list errors.
//   - [IPropertyListSerialization.SetNSPropertyListErrorMinimum]
//   - [IPropertyListSerialization.NSPropertyListErrorMaximum]: The end of the range of error codes reserved for property list errors.
//   - [IPropertyListSerialization.SetNSPropertyListErrorMaximum]
//
// See: https://developer.apple.com/documentation/Foundation/PropertyListSerialization
type IPropertyListSerialization interface {
	objectivec.IObject

	// Topic: Error Codes

	// Parsing of the property list failed.
	NSPropertyListReadCorruptError() int
	SetNSPropertyListReadCorruptError(value int)
	// The version number of the property list cannot be determined.
	NSPropertyListReadUnknownVersionError() int
	SetNSPropertyListReadUnknownVersionError(value int)
	// Reading of the property list failed.
	NSPropertyListReadStreamError() int
	SetNSPropertyListReadStreamError(value int)
	// Writing to the property list failed.
	NSPropertyListWriteStreamError() int
	SetNSPropertyListWriteStreamError(value int)
	// Writing failed because of an invalid property list object, or an invalid property list type was specified.
	NSPropertyListWriteInvalidError() int
	SetNSPropertyListWriteInvalidError(value int)
	// The start of the range of error codes reserved for property list errors.
	NSPropertyListErrorMinimum() int
	SetNSPropertyListErrorMinimum(value int)
	// The end of the range of error codes reserved for property list errors.
	NSPropertyListErrorMaximum() int
	SetNSPropertyListErrorMaximum(value int)
}





// Init initializes the instance.
func (p PropertyListSerialization) Init() PropertyListSerialization {
	rv := objc.Send[PropertyListSerialization](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PropertyListSerialization) Autorelease() PropertyListSerialization {
	rv := objc.Send[PropertyListSerialization](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPropertyListSerialization creates a new PropertyListSerialization instance.
func NewPropertyListSerialization() PropertyListSerialization {
	class := getPropertyListSerializationClass()
	rv := objc.Send[PropertyListSerialization](objc.ID(class.class), objc.Sel("new"))
	return rv
}














// Returns an [NSData] object containing a given property list in a specified
// format.
//
// plist: A property list object.
//
// format: A property list format. For possible values, see
// [PropertyListSerialization.PropertyListFormat].
// //
// [PropertyListSerialization.PropertyListFormat]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/PropertyListFormat
//
// opt: The `opt` parameter is currently unused. No options should be specified.
//
// # Return Value
// 
// An [NSData] object containing `plist` in the format specified by `format`.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/data(fromPropertyList:format:options:)
func (_PropertyListSerializationClass PropertyListSerializationClass) DataWithPropertyListFormatOptionsError(plist objectivec.IObject, format NSPropertyListFormat, opt NSPropertyListWriteOptions) (NSData, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_PropertyListSerializationClass.class), objc.Sel("dataWithPropertyList:format:options:error:"), plist, format, opt, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSData{}, NSErrorFrom(errorPtr)
	}
	return NSDataFromID(rv), nil

}

// Writes a property list to the specified stream.
//
// plist: The property list that you want to write out.
//
// stream: An [NSOutputStream] instance that is open and ready to receive the property
// list data.
//
// format: One of the property list formats defined in
// [PropertyListSerialization.PropertyListFormat].
// //
// [PropertyListSerialization.PropertyListFormat]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/PropertyListFormat
//
// opt: Currently unused. Set to `0`.
//
// error: A pointer that the function may set to an [NSError] object when an error
// occurs to provide additional information about the error.
//
// # Return Value
// 
// The number of bytes written to the stream. A return value of `0` indicates
// that an error occurred.
//
// See: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/writePropertyList(_:to:format:options:error:)
func (_PropertyListSerializationClass PropertyListSerializationClass) WritePropertyListToStreamFormatOptionsError(plist objectivec.IObject, stream INSOutputStream, format NSPropertyListFormat, opt NSPropertyListWriteOptions) (int, error) {
			var errorPtr objc.ID
	rv := objc.Send[int](objc.ID(_PropertyListSerializationClass.class), objc.Sel("writePropertyList:toStream:format:options:error:"), plist, stream, format, opt, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// Creates and returns a property list from the specified data.
//
// data: A data object containing a serialized property list.
//
// opt: The options used to create the property list. For possible values, see
// [PropertyListSerialization.MutabilityOptions].
// //
// [PropertyListSerialization.MutabilityOptions]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/MutabilityOptions
//
// format: Upon return, contains the format that the property list was stored in. Pass
// `nil` if you do not need to know the format.
//
// # Return Value
// 
// A property list object corresponding to the representation in `data`. If
// data is not in a supported format, returns `nil`.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/propertyList(from:options:format:)
func (_PropertyListSerializationClass PropertyListSerializationClass) PropertyListWithDataOptionsFormatError(data INSData, opt NSPropertyListMutabilityOptions, format NSPropertyListFormat) (objectivec.IObject, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_PropertyListSerializationClass.class), objc.Sel("propertyListWithData:options:format:error:"), data, opt, format, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// Creates and returns a property list by reading from the specified stream.
//
// stream: An [NSStream] object. The stream should be open and configured for reading.
//
// opt: The options used to create the property list. For possible values, see
// [PropertyListSerialization.MutabilityOptions].
// //
// [PropertyListSerialization.MutabilityOptions]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/MutabilityOptions
//
// format: Upon return, contains the format that the property list was stored in. Pass
// `nil` if you do not need to know the format.
//
// # Return Value
// 
// A property list object corresponding to the representation in `data`. If
// data is not in a supported format, returns `nil`.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/propertyList(with:options:format:)
func (_PropertyListSerializationClass PropertyListSerializationClass) PropertyListWithStreamOptionsFormatError(stream INSInputStream, opt NSPropertyListMutabilityOptions, format NSPropertyListFormat) (objectivec.IObject, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_PropertyListSerializationClass.class), objc.Sel("propertyListWithStream:options:format:error:"), stream, opt, format, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// Returns a Boolean value that indicates whether a given property list is
// valid for a given format.
//
// plist: A property list object.
//
// format: A property list format. For possible values, see
// [PropertyListSerialization.PropertyListFormat].
// //
// [PropertyListSerialization.PropertyListFormat]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/PropertyListFormat
//
// # Return Value
// 
// [true] if `plist` is a valid property list in format `format`, otherwise
// [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/propertyList(_:isValidFor:)
func (_PropertyListSerializationClass PropertyListSerializationClass) PropertyListIsValidForFormat(plist objectivec.IObject, format NSPropertyListFormat) bool {
	rv := objc.Send[bool](objc.ID(_PropertyListSerializationClass.class), objc.Sel("propertyList:isValidForFormat:"), plist, format)
	return rv
}








// Parsing of the property list failed.
//
// See: https://developer.apple.com/documentation/foundation/nspropertylistreadcorrupterror-swift.var
func (p PropertyListSerialization) NSPropertyListReadCorruptError() int {
	rv := objc.Send[int](p.ID, objc.Sel("NSPropertyListReadCorruptError"))
	return rv
}
func (p PropertyListSerialization) SetNSPropertyListReadCorruptError(value int) {
	objc.Send[struct{}](p.ID, objc.Sel("setNSPropertyListReadCorruptError:"), value)
}



// The version number of the property list cannot be determined.
//
// See: https://developer.apple.com/documentation/foundation/nspropertylistreadunknownversionerror-swift.var
func (p PropertyListSerialization) NSPropertyListReadUnknownVersionError() int {
	rv := objc.Send[int](p.ID, objc.Sel("NSPropertyListReadUnknownVersionError"))
	return rv
}
func (p PropertyListSerialization) SetNSPropertyListReadUnknownVersionError(value int) {
	objc.Send[struct{}](p.ID, objc.Sel("setNSPropertyListReadUnknownVersionError:"), value)
}



// Reading of the property list failed.
//
// See: https://developer.apple.com/documentation/foundation/nspropertylistreadstreamerror-swift.var
func (p PropertyListSerialization) NSPropertyListReadStreamError() int {
	rv := objc.Send[int](p.ID, objc.Sel("NSPropertyListReadStreamError"))
	return rv
}
func (p PropertyListSerialization) SetNSPropertyListReadStreamError(value int) {
	objc.Send[struct{}](p.ID, objc.Sel("setNSPropertyListReadStreamError:"), value)
}



// Writing to the property list failed.
//
// See: https://developer.apple.com/documentation/foundation/nspropertylistwritestreamerror-swift.var
func (p PropertyListSerialization) NSPropertyListWriteStreamError() int {
	rv := objc.Send[int](p.ID, objc.Sel("NSPropertyListWriteStreamError"))
	return rv
}
func (p PropertyListSerialization) SetNSPropertyListWriteStreamError(value int) {
	objc.Send[struct{}](p.ID, objc.Sel("setNSPropertyListWriteStreamError:"), value)
}



// Writing failed because of an invalid property list object, or an invalid
// property list type was specified.
//
// See: https://developer.apple.com/documentation/foundation/nspropertylistwriteinvaliderror-swift.var
func (p PropertyListSerialization) NSPropertyListWriteInvalidError() int {
	rv := objc.Send[int](p.ID, objc.Sel("NSPropertyListWriteInvalidError"))
	return rv
}
func (p PropertyListSerialization) SetNSPropertyListWriteInvalidError(value int) {
	objc.Send[struct{}](p.ID, objc.Sel("setNSPropertyListWriteInvalidError:"), value)
}



// The start of the range of error codes reserved for property list errors.
//
// See: https://developer.apple.com/documentation/foundation/nspropertylisterrorminimum-swift.var
func (p PropertyListSerialization) NSPropertyListErrorMinimum() int {
	rv := objc.Send[int](p.ID, objc.Sel("NSPropertyListErrorMinimum"))
	return rv
}
func (p PropertyListSerialization) SetNSPropertyListErrorMinimum(value int) {
	objc.Send[struct{}](p.ID, objc.Sel("setNSPropertyListErrorMinimum:"), value)
}



// The end of the range of error codes reserved for property list errors.
//
// See: https://developer.apple.com/documentation/foundation/nspropertylisterrormaximum-swift.var
func (p PropertyListSerialization) NSPropertyListErrorMaximum() int {
	rv := objc.Send[int](p.ID, objc.Sel("NSPropertyListErrorMaximum"))
	return rv
}
func (p PropertyListSerialization) SetNSPropertyListErrorMaximum(value int) {
	objc.Send[struct{}](p.ID, objc.Sel("setNSPropertyListErrorMaximum:"), value)
}























