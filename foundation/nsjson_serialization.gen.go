// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [JSONSerialization] class.
var (
	_JSONSerializationClass     JSONSerializationClass
	_JSONSerializationClassOnce sync.Once
)

func getJSONSerializationClass() JSONSerializationClass {
	_JSONSerializationClassOnce.Do(func() {
		_JSONSerializationClass = JSONSerializationClass{class: objc.GetClass("NSJSONSerialization")}
	})
	return _JSONSerializationClass
}

// GetJSONSerializationClass returns the class object for NSJSONSerialization.
func GetJSONSerializationClass() JSONSerializationClass {
	return getJSONSerializationClass()
}

type JSONSerializationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (jc JSONSerializationClass) Class() objc.Class {
	return jc.class
}

// Alloc allocates memory for a new instance of the class.
func (jc JSONSerializationClass) Alloc() JSONSerialization {
	rv := objc.Send[JSONSerialization](objc.ID(jc.class), objc.Sel("alloc"))
	return rv
}

// An object that converts between JSON and the equivalent Foundation objects.
//
// # Overview
// 
// You use the [NSJSONSerialization] class to convert JSON to Foundation
// objects and convert Foundation objects to JSON.
// 
// To convert a Foundation object to JSON, the object must have the following
// properties:
// 
// - The top level object is an [NSArray] or [NSDictionary], unless you set
// the [JSONWritingFragmentsAllowed] option. - All objects are instances of
// [NSString], [NSNumber], [NSArray], [NSDictionary], or [NSNull]. - All
// dictionary keys are instances of [NSString]. - Numbers are neither [NaN]
// nor infinity.
// 
// Other rules may apply. Calling [IsValidJSONObject] or attempting a
// conversion are the definitive ways to tell if the [NSJSONSerialization]
// class can convert given object to JSON data.
//
// See: https://developer.apple.com/documentation/Foundation/JSONSerialization
type JSONSerialization struct {
	objectivec.Object
}

// JSONSerializationFromID constructs a [JSONSerialization] from an objc.ID.
//
// An object that converts between JSON and the equivalent Foundation objects.
func JSONSerializationFromID(id objc.ID) JSONSerialization {
	return NSJSONSerialization{objectivec.Object{ID: id}}
}

// NSJSONSerializationFromID is an alias for [JSONSerializationFromID] for cross-framework compatibility.
func NSJSONSerializationFromID(id objc.ID) JSONSerialization { return JSONSerializationFromID(id) }
// NOTE: JSONSerialization adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [JSONSerialization] class.
//
// See: https://developer.apple.com/documentation/Foundation/JSONSerialization
type IJSONSerialization interface {
	objectivec.IObject

	// Specifies that the parser should allow top-level objects that aren’t arrays or dictionaries.
	FragmentsAllowed() NSJSONWritingOptions
	SetFragmentsAllowed(value NSJSONWritingOptions)
}

// Init initializes the instance.
func (j JSONSerialization) Init() JSONSerialization {
	rv := objc.Send[JSONSerialization](j.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (j JSONSerialization) Autorelease() JSONSerialization {
	rv := objc.Send[JSONSerialization](j.ID, objc.Sel("autorelease"))
	return rv
}

// NewJSONSerialization creates a new JSONSerialization instance.
func NewJSONSerialization() JSONSerialization {
	class := getJSONSerializationClass()
	rv := objc.Send[JSONSerialization](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a Foundation object from given JSON data.
//
// data: A data object containing JSON data.
//
// opt: Options for reading the JSON data and creating the Foundation objects.
// 
// For possible values, see [JSONSerialization.ReadingOptions].
// //
// [JSONSerialization.ReadingOptions]: https://developer.apple.com/documentation/Foundation/JSONSerialization/ReadingOptions
//
// # Return Value
// 
// A Foundation object from the JSON data in `data`, or `nil` if an error
// occurs.
//
// # Discussion
// 
// The data must be in one of the 5 supported encodings listed in the JSON
// specification: UTF-8, UTF-16LE, UTF-16BE, UTF-32LE, UTF-32BE. The data may
// or may not have a BOM. The most efficient encoding to use for parsing is
// UTF-8, so if you have a choice in encoding the data passed to this method,
// use UTF-8.
//
// See: https://developer.apple.com/documentation/Foundation/JSONSerialization/jsonObject(with:options:)-8demi
func (_JSONSerializationClass JSONSerializationClass) JSONObjectWithDataOptionsError(data INSData, opt NSJSONReadingOptions) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_JSONSerializationClass.class), objc.Sel("JSONObjectWithData:options:error:"), data, opt, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// Returns a Foundation object from JSON data in a given stream.
//
// stream: A stream from which to read JSON data.
// 
// The stream should be open and configured.
//
// opt: Options for reading the JSON data and creating the Foundation objects.
// 
// For possible values, see [JSONSerialization.ReadingOptions].
// //
// [JSONSerialization.ReadingOptions]: https://developer.apple.com/documentation/Foundation/JSONSerialization/ReadingOptions
//
// # Return Value
// 
// A Foundation object from the JSON data in `stream`.
//
// # Discussion
// 
// The data in the stream must be in one of the 5 supported encodings listed
// in the JSON specification: UTF-8, UTF-16LE, UTF-16BE, UTF-32LE, UTF-32BE.
// The data may or may not have a BOM. The most efficient encoding to use for
// parsing is UTF-8, so if you have a choice in encoding the data passed to
// this method, use UTF-8.
//
// See: https://developer.apple.com/documentation/Foundation/JSONSerialization/jsonObject(with:options:)-3afap
func (_JSONSerializationClass JSONSerializationClass) JSONObjectWithStreamOptionsError(stream INSInputStream, opt NSJSONReadingOptions) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_JSONSerializationClass.class), objc.Sel("JSONObjectWithStream:options:error:"), stream, opt, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// Returns JSON data from a Foundation object.
//
// obj: The object from which to generate JSON data. Must not be `nil`.
//
// opt: Options for creating the JSON data.
// 
// See [JSONSerialization.WritingOptions] for possible values.
// //
// [JSONSerialization.WritingOptions]: https://developer.apple.com/documentation/Foundation/JSONSerialization/WritingOptions
//
// # Return Value
// 
// JSON data for `obj`, or `nil` if an internal error occurs. The resulting
// data is encoded in UTF-8.
//
// # Discussion
// 
// If `obj` can’t produce valid JSON, [NSJSONSerialization] throws an
// exception. This exception occurs prior to parsing and represents a
// programming error, not an internal error. Before calling this method, you
// should check whether the input can produce valid JSON by using
// [IsValidJSONObject].
// 
// Setting the [JSONWritingPrettyPrinted] option generates JSON with white
// space designed to make the output more readable. If this option isn’t
// set, [NSJSONSerialization] generates the most compact possible JSON.
//
// See: https://developer.apple.com/documentation/Foundation/JSONSerialization/data(withJSONObject:options:)
func (_JSONSerializationClass JSONSerializationClass) DataWithJSONObjectOptionsError(obj objectivec.IObject, opt NSJSONWritingOptions) (NSData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_JSONSerializationClass.class), objc.Sel("dataWithJSONObject:options:error:"), obj, opt, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSData{}, NSErrorFrom(errorPtr)
	}
	return NSDataFromID(rv), nil

}
// Writes a given JSON object to a stream.
//
// obj: The object to write to `stream`.
//
// stream: The stream to which to write.
// 
// The stream should be open and configured.
//
// opt: Options for writing the JSON data.
// 
// See [JSONSerialization.WritingOptions] for possible values.
// //
// [JSONSerialization.WritingOptions]: https://developer.apple.com/documentation/Foundation/JSONSerialization/WritingOptions
//
// error: If an error occurs, upon return contains an [NSError] object with code
// [NSPropertyListWriteInvalidError] that describes the problem.
// //
// [NSPropertyListWriteInvalidError]: https://developer.apple.com/documentation/Foundation/NSPropertyListWriteInvalidError-swift.var
//
// # Return Value
// 
// The number of bytes written to the stream, or `0` if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/JSONSerialization/writeJSONObject(_:to:options:error:)
func (_JSONSerializationClass JSONSerializationClass) WriteJSONObjectToStreamOptionsError(obj objectivec.IObject, stream INSOutputStream, opt NSJSONWritingOptions) (int, error) {
	var errorPtr objc.ID
	rv := objc.Send[int](objc.ID(_JSONSerializationClass.class), objc.Sel("writeJSONObject:toStream:options:error:"), obj, stream, opt, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, NSErrorFrom(errorPtr)
	}
	return rv, nil

}
// Returns a Boolean value that indicates whether the serializer can convert a
// given object to JSON data.
//
// obj: The object to test.
//
// # Return Value
// 
// `true` if `obj` can be converted to JSON data; otherwise, `false`.
//
// See: https://developer.apple.com/documentation/Foundation/JSONSerialization/isValidJSONObject(_:)
func (_JSONSerializationClass JSONSerializationClass) IsValidJSONObject(obj objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_JSONSerializationClass.class), objc.Sel("isValidJSONObject:"), obj)
	return rv
}

// Specifies that the parser should allow top-level objects that aren’t
// arrays or dictionaries.
//
// See: https://developer.apple.com/documentation/foundation/jsonserialization/writingoptions/fragmentsallowed
func (j JSONSerialization) FragmentsAllowed() NSJSONWritingOptions {
	rv := objc.Send[NSJSONWritingOptions](j.ID, objc.Sel("NSJSONWritingFragmentsAllowed"))
	return NSJSONWritingOptions(rv)
}
func (j JSONSerialization) SetFragmentsAllowed(value NSJSONWritingOptions) {
	objc.Send[struct{}](j.ID, objc.Sel("setNSJSONWritingFragmentsAllowed:"), value)
}

