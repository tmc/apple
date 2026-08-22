// Code generated from Apple documentation for ScriptingBridge. DO NOT EDIT.

package scriptingbridge

import (
	"sync"

	"github.com/tmc/apple/coreservices"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SBObject] class.
var (
	_SBObjectClass     SBObjectClass
	_SBObjectClassOnce sync.Once
)

func getSBObjectClass() SBObjectClass {
	_SBObjectClassOnce.Do(func() {
		_SBObjectClass = SBObjectClass{class: objc.GetClass("SBObject")}
	})
	return _SBObjectClass
}

// GetSBObjectClass returns the class object for SBObject.
func GetSBObjectClass() SBObjectClass {
	return getSBObjectClass()
}

type SBObjectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SBObjectClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SBObjectClass) Alloc() SBObject {
	rv := objc.Send[SBObject](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// The [SBObject] class declares methods that can be invoked on any object in
// a scriptable application. It defines methods for getting elements and
// properties of an object, as well as setting a given object to a new value.
//
// # Overview
//
// Each [SBObject] is built around an object specifier, which tells Scripting
// Bridge how to locate the object. Therefore, you can think of an [SBObject]
// as a reference to an object in an target application rather than an object
// itself. To bypass this reference-based approach and force evaluation, use
// the [SBObject.Get] method.
//
// Typically, rather than create [SBObject] instances explictly, you receive
// [SBObject] objects by calling methods of an [SBApplication] subclass. For
// example, if you wanted to get an [SBObject] representing the current iTunes
// track, you would use code like this (where `iTunesTrack` is a subclass of
// [SBObject]):
//
// You can discover the names of dynamically generated classes such as
// `iTunesApplication` and `iTunesTrack` by examining the header file created
// by the `sdp` tool. Alternatively, you give these variables the dynamic
// Objective-C type `id`.
//
// # Initializing a Scripting Bridge Object
//
//   - [SBObject.InitWithData]: Returns an instance of an [SBObject] subclass initialized with the given data.
//   - [SBObject.InitWithProperties]: Returns an instance of an [SBObject] subclass initialized with the specified properties.
//   - [SBObject.InitWithElementCodePropertiesData]: Returns an instance of an [SBObject] subclass initialized with the specified properties and data and added to the designated element array.
//
// # Getting Referenced Data
//
//   - [SBObject.Get]: Forces evaluation of the receiver, causing the real object to be returned immediately.
//
// # Sending Apple Events
//
//   - [SBObject.SetTo]: Sets the receiver to a specified value.
//
// # Getting Properties and Elements
//
//   - [SBObject.PropertyWithClassCode]: Returns an object of the designated scripting class representing the specified property of the receiver
//   - [SBObject.PropertyWithCode]: Returns an object representing the specified property of the receiver.
//   - [SBObject.ElementArrayWithCode]: Returns an array containing every child of the receiver with the given class-type code.
//
// # Instance Methods
//
//   - [SBObject.LastError]: The error from the last event this object sent, or nil if it succeeded.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBObject
type SBObject struct {
	objectivec.Object
}

// SBObjectFromID constructs a [SBObject] from an objc.ID.
//
// The [SBObject] class declares methods that can be invoked on any object in
// a scriptable application. It defines methods for getting elements and
// properties of an object, as well as setting a given object to a new value.
func SBObjectFromID(id objc.ID) SBObject {
	return SBObject{objectivec.Object{ID: id}}
}

// NOTE: SBObject adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SBObject] class.
//
// # Initializing a Scripting Bridge Object
//
//   - [ISBObject.InitWithData]: Returns an instance of an [SBObject] subclass initialized with the given data.
//   - [ISBObject.InitWithProperties]: Returns an instance of an [SBObject] subclass initialized with the specified properties.
//   - [ISBObject.InitWithElementCodePropertiesData]: Returns an instance of an [SBObject] subclass initialized with the specified properties and data and added to the designated element array.
//
// # Getting Referenced Data
//
//   - [ISBObject.Get]: Forces evaluation of the receiver, causing the real object to be returned immediately.
//
// # Sending Apple Events
//
//   - [ISBObject.SetTo]: Sets the receiver to a specified value.
//
// # Getting Properties and Elements
//
//   - [ISBObject.PropertyWithClassCode]: Returns an object of the designated scripting class representing the specified property of the receiver
//   - [ISBObject.PropertyWithCode]: Returns an object representing the specified property of the receiver.
//   - [ISBObject.ElementArrayWithCode]: Returns an array containing every child of the receiver with the given class-type code.
//
// # Instance Methods
//
//   - [ISBObject.LastError]: The error from the last event this object sent, or nil if it succeeded.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBObject
type ISBObject interface {
	objectivec.IObject

	// Topic: Initializing a Scripting Bridge Object

	// Returns an instance of an [SBObject] subclass initialized with the given data.
	InitWithData(data objectivec.IObject) SBObject
	// Returns an instance of an [SBObject] subclass initialized with the specified properties.
	InitWithProperties(properties foundation.INSDictionary) SBObject
	// Returns an instance of an [SBObject] subclass initialized with the specified properties and data and added to the designated element array.
	InitWithElementCodePropertiesData(code coreservices.DescType, properties foundation.INSDictionary, data objectivec.IObject) SBObject

	// Topic: Getting Referenced Data

	// Forces evaluation of the receiver, causing the real object to be returned immediately.
	Get() objectivec.IObject

	// Topic: Sending Apple Events

	// Sets the receiver to a specified value.
	SetTo(value objectivec.IObject)

	// Topic: Getting Properties and Elements

	// Returns an object of the designated scripting class representing the specified property of the receiver
	PropertyWithClassCode(cls objectivec.Class, code coreservices.AEKeyword) ISBObject
	// Returns an object representing the specified property of the receiver.
	PropertyWithCode(code coreservices.AEKeyword) ISBObject
	// Returns an array containing every child of the receiver with the given class-type code.
	ElementArrayWithCode(code coreservices.DescType) ISBElementArray

	// Topic: Instance Methods

	// The error from the last event this object sent, or nil if it succeeded.
	LastError() foundation.NSError

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s SBObject) Init() SBObject {
	rv := objc.Send[SBObject](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SBObject) Autorelease() SBObject {
	rv := objc.Send[SBObject](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSBObject creates a new SBObject instance.
func NewSBObject() SBObject {
	class := getSBObjectClass()
	rv := objc.Send[SBObject](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an instance of an [SBObject] subclass initialized with the given
// data.
//
// data: An object containing data for the new [SBObject] object. The data varies
// according to the type of scripting object to be created.
//
// # Return Value
//
// An [SBObject] object or `nil` if the object could not be initialized.
//
// # Discussion
//
// Scripting Bridge does not actually create an object in the target
// application until you add the object returned from this method to an
// element array ([SBElementArray]).
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBObject/init(data:)
func NewSBObjectWithData(data objectivec.IObject) SBObject {
	instance := getSBObjectClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:"), data)
	return SBObjectFromID(rv)
}

// Returns an instance of an [SBObject] subclass initialized with the
// specified properties and data and added to the designated element array.
//
// code: A four-character code used to identify an element in the target
// application’s scripting interface. See [Apple Event Manager] for details.
//
// properties: A dictionary with [NSNumber] keys specifying the four-character codes of
// properties (that is, attributes or to-one relationships) and the values for
// those properties. Pass `nil` if you are initializing the object by `data`
// only.
//
// data: An object containing data for the new [SBObject] object. The data varies
// according to the type of scripting object to be created. Pass `nil` if you
// initializing the object by `properties` only.
//
// # Return Value
//
// An [SBObject] object or `nil` if the object could not be initialized.
//
// # Discussion
//
// Unlike the other initializers of this class, this method not only
// initializes the [SBObject] object but adds it to a specified element array.
// This method is the designated initializer.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBObject/init(elementCode:properties:data:)
//
// [Apple Event Manager]: https://developer.apple.com/documentation/applicationservices/apple_event_manager
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func NewSBObjectWithElementCodePropertiesData(code coreservices.DescType, properties foundation.INSDictionary, data objectivec.IObject) SBObject {
	instance := getSBObjectClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithElementCode:properties:data:"), code, properties, data)
	return SBObjectFromID(rv)
}

// Returns an instance of an [SBObject] subclass initialized with the
// specified properties.
//
// properties: A dictionary with keys specifying the names of properties (that is,
// attributes or to-one relationships) and the values for those properties.
//
// # Return Value
//
// An [SBObject] object or `nil` if the object could not be initialized.
//
// # Discussion
//
// Scripting Bridge does not actually create an object in the target
// application until you add the object returned from this method to an
// element array ([SBElementArray]).
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBObject/init(properties:)
func NewSBObjectWithProperties(properties foundation.INSDictionary) SBObject {
	instance := getSBObjectClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProperties:"), properties)
	return SBObjectFromID(rv)
}

// Returns an instance of an [SBObject] subclass initialized with the given
// data.
//
// data: An object containing data for the new [SBObject] object. The data varies
// according to the type of scripting object to be created.
//
// # Return Value
//
// An [SBObject] object or `nil` if the object could not be initialized.
//
// # Discussion
//
// Scripting Bridge does not actually create an object in the target
// application until you add the object returned from this method to an
// element array ([SBElementArray]).
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBObject/init(data:)
func (s SBObject) InitWithData(data objectivec.IObject) SBObject {
	rv := objc.Send[SBObject](s.ID, objc.Sel("initWithData:"), data)
	return rv
}

// Returns an instance of an [SBObject] subclass initialized with the
// specified properties.
//
// properties: A dictionary with keys specifying the names of properties (that is,
// attributes or to-one relationships) and the values for those properties.
//
// # Return Value
//
// An [SBObject] object or `nil` if the object could not be initialized.
//
// # Discussion
//
// Scripting Bridge does not actually create an object in the target
// application until you add the object returned from this method to an
// element array ([SBElementArray]).
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBObject/init(properties:)
func (s SBObject) InitWithProperties(properties foundation.INSDictionary) SBObject {
	rv := objc.Send[SBObject](s.ID, objc.Sel("initWithProperties:"), properties)
	return rv
}

// Returns an instance of an [SBObject] subclass initialized with the
// specified properties and data and added to the designated element array.
//
// code: A four-character code used to identify an element in the target
// application’s scripting interface. See [Apple Event Manager] for details.
//
// properties: A dictionary with [NSNumber] keys specifying the four-character codes of
// properties (that is, attributes or to-one relationships) and the values for
// those properties. Pass `nil` if you are initializing the object by `data`
// only.
//
// data: An object containing data for the new [SBObject] object. The data varies
// according to the type of scripting object to be created. Pass `nil` if you
// initializing the object by `properties` only.
//
// # Return Value
//
// An [SBObject] object or `nil` if the object could not be initialized.
//
// # Discussion
//
// Unlike the other initializers of this class, this method not only
// initializes the [SBObject] object but adds it to a specified element array.
// This method is the designated initializer.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBObject/init(elementCode:properties:data:)
//
// [Apple Event Manager]: https://developer.apple.com/documentation/applicationservices/apple_event_manager
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (s SBObject) InitWithElementCodePropertiesData(code coreservices.DescType, properties foundation.INSDictionary, data objectivec.IObject) SBObject {
	rv := objc.Send[SBObject](s.ID, objc.Sel("initWithElementCode:properties:data:"), code, properties, data)
	return rv
}

// Forces evaluation of the receiver, causing the real object to be returned
// immediately.
//
// # Return Value
//
// For most properties, the result is a Foundation object such as an
// [NSString]. For properties with no Foundation equivalent, the result is an
// [NSAppleEventDescriptor] or another [SBObject] for most elements.
//
// # Discussion
//
// This method forces the current object reference (the receiver) to be
// evaluated, resulting in the return of the referenced object. By default,
// Scripting Bridge deals with references to objects until you actually
// request some concrete data from them or until you call the `get` method.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBObject/get()
func (s SBObject) Get() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("get"))
	return objectivec.Object{ID: rv}
}

// Sets the receiver to a specified value.
//
// value: The data the receiver should be set to. It can be an [NSString],
// [NSNumber], [NSArray], [SBObject], or any other type of object supported by
// the Scripting Bridge framework.
//
// # Discussion
//
// You should not call this method directly.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBObject/setTo(_:)
//
// [NSArray]: https://developer.apple.com/documentation/Foundation/NSArray
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
func (s SBObject) SetTo(value objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("setTo:"), value)
}

// Returns an object of the designated scripting class representing the
// specified property of the receiver
//
// code: A four-character code that uniquely identifies a property of the receiver.
//
// # Return Value
//
// An instance of the designated `class` that represents the receiver’s
// property identified by `code`.
//
// # Discussion
//
// [SBObject] subclasses use this method to implement application-specific
// property accessor methods. You should not need to call this method
// directly.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBObject/property(with:code:)
func (s SBObject) PropertyWithClassCode(cls objectivec.Class, code coreservices.AEKeyword) ISBObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("propertyWithClass:code:"), cls, code)
	return SBObjectFromID(rv)
}

// Returns an object representing the specified property of the receiver.
//
// code: A four-character code that uniquely identifies a property of the receiver.
//
// # Return Value
//
// An object representing the receiver’s property as identified by `code`.
//
// # Discussion
//
// [SBObject] subclasses use this method to implement application-specific
// property accessor methods. You should not need to call this method
// directly.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBObject/property(withCode:)
func (s SBObject) PropertyWithCode(code coreservices.AEKeyword) ISBObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("propertyWithCode:"), code)
	return SBObjectFromID(rv)
}

// Returns an array containing every child of the receiver with the given
// class-type code.
//
// code: A four-character code that identifies a scripting class.
//
// # Return Value
//
// An [SBElementArray] object containing every child of the receiver whose
// class matches `code`.
//
// # Discussion
//
// [SBObject] subclasses use this method to implement application-specific
// property accessor methods. You should not need to call this method
// directly.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBObject/elementArray(withCode:)
func (s SBObject) ElementArrayWithCode(code coreservices.DescType) ISBElementArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("elementArrayWithCode:"), code)
	return SBElementArrayFromID(rv)
}

// The error from the last event this object sent, or nil if it succeeded.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBObject/lastError()
func (s SBObject) LastError() foundation.NSError {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("lastError"))
	return foundation.NSErrorFromID(rv)
}
func (s SBObject) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}
