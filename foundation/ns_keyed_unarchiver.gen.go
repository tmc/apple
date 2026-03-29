// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSKeyedUnarchiver] class.
var (
	_NSKeyedUnarchiverClass     NSKeyedUnarchiverClass
	_NSKeyedUnarchiverClassOnce sync.Once
)

func getNSKeyedUnarchiverClass() NSKeyedUnarchiverClass {
	_NSKeyedUnarchiverClassOnce.Do(func() {
		_NSKeyedUnarchiverClass = NSKeyedUnarchiverClass{class: objc.GetClass("NSKeyedUnarchiver")}
	})
	return _NSKeyedUnarchiverClass
}

// GetNSKeyedUnarchiverClass returns the class object for NSKeyedUnarchiver.
func GetNSKeyedUnarchiverClass() NSKeyedUnarchiverClass {
	return getNSKeyedUnarchiverClass()
}

type NSKeyedUnarchiverClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSKeyedUnarchiverClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSKeyedUnarchiverClass) Alloc() NSKeyedUnarchiver {
	rv := objc.Send[NSKeyedUnarchiver](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A decoder that restores data from an archive referenced by keys.
//
// # Overview
// 
// [NSKeyedUnarchiver] is a concrete subclass of [NSCoder] that defines
// methods for decoding a set of named objects (and scalar values) from a
// keyed archive. The [NSKeyedArchiver] class produces archives that this
// class can decode.
// 
// The archiver creates keyed archive as a hierarchy of objects. The archiver
// treats each object as a namespace into which it can encode other objects.
// This means that an unarchiver can only decode objects encoded within the
// immediate scope of their parent object. Objects encoded elsewhere in the
// hierarchy — whether higher than, lower than, or parallel to this
// particular object — aren’t accessible. In this way, the keys used by a
// particular object to encode its instance variables need to be unique only
// within the scope of that object.
// 
// If you invoke one of the `decode`-prefixed methods of this class using a
// key that does not exist in the archive, the return value indicates failure.
// This value varies by decoded type. For example, if a key does not exist in
// an archive, [NSKeyedUnarchiver.DecodeBoolForKey] returns [false], [decodeIntForKey:] returns
// `0`, and [NSKeyedUnarchiver.DecodeObjectForKey] returns `nil`.
// 
// [NSKeyedUnarchiver] supports limited type coercion for numeric types. You
// can use any of the integer decode methods to decode a value encoded as any
// type of integer, whether a standard [Int] or an explicit 32-bit or 64-bit
// integer. Likewise, you can use the [Float]- or [Double]-returning decode
// methods to handle value encoded as a [Float] or [Double]. If an encoded
// value is too large to fit within the coerced type, the decoding method
// throws a [rangeException]. Further, when trying to coerce a value to an
// incompatible type — for example decoding an [Int] as a [Float] — the
// decoding method throws an [invalidUnarchiveOperationException].
//
// [decodeIntForKey:]: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/decodeIntForKey:
// [false]: https://developer.apple.com/documentation/Swift/false
// [invalidUnarchiveOperationException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidUnarchiveOperationException
// [rangeException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
//
// # Creating a Keyed Unarchiver
//
//   - [NSKeyedUnarchiver.InitForReadingFromDataError]: Initializes an archiver to decode data from the specified location.
//
// # Decoding Data
//
//   - [NSKeyedUnarchiver.FinishDecoding]: Tells the receiver that you are finished decoding objects.
//
// # Managing the Delegate
//
//   - [NSKeyedUnarchiver.Delegate]: The receiver’s delegate.
//   - [NSKeyedUnarchiver.SetDelegate]
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver
type NSKeyedUnarchiver struct {
	NSCoder
}

// NSKeyedUnarchiverFromID constructs a [NSKeyedUnarchiver] from an objc.ID.
//
// A decoder that restores data from an archive referenced by keys.
func NSKeyedUnarchiverFromID(id objc.ID) NSKeyedUnarchiver {
	return NSKeyedUnarchiver{NSCoder: NSCoderFromID(id)}
}
// NOTE: NSKeyedUnarchiver adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSKeyedUnarchiver] class.
//
// # Creating a Keyed Unarchiver
//
//   - [INSKeyedUnarchiver.InitForReadingFromDataError]: Initializes an archiver to decode data from the specified location.
//
// # Decoding Data
//
//   - [INSKeyedUnarchiver.FinishDecoding]: Tells the receiver that you are finished decoding objects.
//
// # Managing the Delegate
//
//   - [INSKeyedUnarchiver.Delegate]: The receiver’s delegate.
//   - [INSKeyedUnarchiver.SetDelegate]
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver
type INSKeyedUnarchiver interface {
	INSCoder

	// Topic: Creating a Keyed Unarchiver

	// Initializes an archiver to decode data from the specified location.
	InitForReadingFromDataError(data INSData) (NSKeyedUnarchiver, error)

	// Topic: Decoding Data

	// Tells the receiver that you are finished decoding objects.
	FinishDecoding()

	// Topic: Managing the Delegate

	// The receiver’s delegate.
	Delegate() NSKeyedUnarchiverDelegate
	SetDelegate(value NSKeyedUnarchiverDelegate)

	// The name of the exception raised by 
	InvalidUnarchiveOperationException() NSExceptionName
	// Name of an exception that occurs when attempting to access outside the bounds of some data, such as beyond the end of a string.
	RangeException() NSExceptionName
}

// Init initializes the instance.
func (k NSKeyedUnarchiver) Init() NSKeyedUnarchiver {
	rv := objc.Send[NSKeyedUnarchiver](k.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k NSKeyedUnarchiver) Autorelease() NSKeyedUnarchiver {
	rv := objc.Send[NSKeyedUnarchiver](k.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSKeyedUnarchiver creates a new NSKeyedUnarchiver instance.
func NewNSKeyedUnarchiver() NSKeyedUnarchiver {
	class := getNSKeyedUnarchiverClass()
	rv := objc.Send[NSKeyedUnarchiver](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes an archiver to decode data from the specified location.
//
// data: An archive previously encoded by [NSKeyedArchiver].
//
// # Discussion
// 
// This initializer enables [RequiresSecureCoding] by default, and sets the
// [DecodingFailurePolicy] to [DecodingFailurePolicySetErrorAndReturn].
// 
// Call [FinishDecoding] when you finish decoding data
// 
// This method throws an error if `data` isn’t a valid keyed archive.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/init(forReadingFrom:)
func NewKeyedUnarchiverForReadingFromDataError(data INSData) (NSKeyedUnarchiver, error) {
	var errorPtr objc.ID
	instance := getNSKeyedUnarchiverClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initForReadingFromData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSKeyedUnarchiver{}, NSErrorFrom(errorPtr)
	}
	return NSKeyedUnarchiverFromID(rv), nil
}

// Initializes an archiver to decode data from the specified location.
//
// data: An archive previously encoded by [NSKeyedArchiver].
//
// # Discussion
// 
// This initializer enables [RequiresSecureCoding] by default, and sets the
// [DecodingFailurePolicy] to [DecodingFailurePolicySetErrorAndReturn].
// 
// Call [FinishDecoding] when you finish decoding data
// 
// This method throws an error if `data` isn’t a valid keyed archive.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/init(forReadingFrom:)
func (k NSKeyedUnarchiver) InitForReadingFromDataError(data INSData) (NSKeyedUnarchiver, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](k.ID, objc.Sel("initForReadingFromData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSKeyedUnarchiver{}, NSErrorFrom(errorPtr)
	}
	return NSKeyedUnarchiverFromID(rv), nil

}
// Tells the receiver that you are finished decoding objects.
//
// # Discussion
// 
// Invoking this method allows the receiver to notify its delegate and to
// perform any final operations on the archive. Once this method is invoked,
// the receiver cannot decode any further values.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/finishDecoding()
func (k NSKeyedUnarchiver) FinishDecoding() {
	objc.Send[objc.ID](k.ID, objc.Sel("finishDecoding"))
}

// Decodes a previously-archived object graph, returning the root object as
// one of the specified classes.
//
// classes: A set of classes, at least one of which the root object should conform to.
//
// data: An object graph previously encoded by [NSKeyedArchiver].
//
// # Return Value
// 
// The decoded root of the object graph, as an instance of one of the
// specified classes, or `nil` if an error occurred.
//
// # Discussion
// 
// This method produces an error if data does not contain valid keyed data.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/unarchivedObject(ofClasses:from:)-b9t5
func (_NSKeyedUnarchiverClass NSKeyedUnarchiverClass) UnarchivedObjectOfClassesFromDataError(classes INSSet, data INSData) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSKeyedUnarchiverClass.class), objc.Sel("unarchivedObjectOfClasses:fromData:error:"), classes, data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// Sets a global translation mapping to decode objects encoded with a given
// class name as instances of a given class instead.
//
// cls: The class with which to replace instances of the class named `codedName`.
//
// codedName: The ostensible name of a class in an archive.
//
// # Discussion
// 
// When decoding, the class’s translation mapping is used only if no
// translation is found first in an instance’s separate translation map.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/setClass(_:forClassName:)-swift.type.method
func (_NSKeyedUnarchiverClass NSKeyedUnarchiverClass) SetClassForClassName(cls objc.Class, codedName string) {
	objc.Send[objc.ID](objc.ID(_NSKeyedUnarchiverClass.class), objc.Sel("setClass:forClassName:"), cls, objc.String(codedName))
}
// Returns the class from which this unarchiver instantiates an encoded object
// with a given class name.
//
// codedName: The ostensible name of a class in an archive.
//
// # Return Value
// 
// The class from which [NSKeyedUnarchiver] instantiates an object encoded
// with the class name `codedName`. Returns `nil` if [NSKeyedUnarchiver] does
// not have a translation mapping for `codedName`.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/class(forClassName:)-swift.type.method
func (_NSKeyedUnarchiverClass NSKeyedUnarchiverClass) ClassForClassNameWithCodedName(codedName string) objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_NSKeyedUnarchiverClass.class), objc.Sel("classForClassName:"), objc.String(codedName))
	return rv
}
// Decodes the \c NSArray root object from \c data which should be an \c
// NSArray containing the given non-collection class (no nested arrays or
// arrays of dictionaries, etc) from the given archive, previously encoded by
// \c NSKeyedArchiver.
//
// # Discussion
// 
// Enables \c requiresSecureCoding and sets the \c decodingFailurePolicy to \c
// NSDecodingFailurePolicySetErrorAndReturn.
// 
// Returns \c nil if the given data is not valid or cannot be decoded, and
// sets the \c error out parameter.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/unarchivedArrayOfObjectsOfClass:fromData:error:
func (_NSKeyedUnarchiverClass NSKeyedUnarchiverClass) UnarchivedArrayOfObjectsOfClassFromDataError(cls objc.Class, data INSData) (INSArray, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSKeyedUnarchiverClass.class), objc.Sel("unarchivedArrayOfObjectsOfClass:fromData:error:"), cls, data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return NSArrayFromID(rv), nil

}
// Decodes the \c NSArray root object from \c data which should be an \c
// NSArray, containing the given non-collection classes in \c classes (no
// nested arrays or arrays of dictionaries, etc) from the given archive,
// previously encoded by \c NSKeyedArchiver.
//
// # Discussion
// 
// Enables \c requiresSecureCoding and sets the \c decodingFailurePolicy to \c
// NSDecodingFailurePolicySetErrorAndReturn.
// 
// Returns \c nil if the given data is not valid or cannot be decoded, and
// sets the \c error out parameter.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/unarchivedArrayOfObjectsOfClasses:fromData:error:
func (_NSKeyedUnarchiverClass NSKeyedUnarchiverClass) UnarchivedArrayOfObjectsOfClassesFromDataError(classes INSSet, data INSData) (INSArray, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSKeyedUnarchiverClass.class), objc.Sel("unarchivedArrayOfObjectsOfClasses:fromData:error:"), classes, data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return NSArrayFromID(rv), nil

}
// Decodes the \c NSDictionary root object from \c data which should be an \c
// NSDictionary with keys of type given in \c keyCls and objects of the given
// non-collection class \c objectCls (no nested dictionaries or other
// dictionaries contained in the dictionary, etc) from the given archive,
// previously encoded by \c NSKeyedArchiver.
//
// # Discussion
// 
// Enables \c requiresSecureCoding and sets the \c decodingFailurePolicy to \c
// NSDecodingFailurePolicySetErrorAndReturn.
// 
// Returns \c nil if the given data is not valid or cannot be decoded, and
// sets the \c error out parameter.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/unarchivedDictionaryWithKeysOfClass:objectsOfClass:fromData:error:
func (_NSKeyedUnarchiverClass NSKeyedUnarchiverClass) UnarchivedDictionaryWithKeysOfClassObjectsOfClassFromDataError(keyCls objc.Class, valueCls objc.Class, data INSData) (INSDictionary, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSKeyedUnarchiverClass.class), objc.Sel("unarchivedDictionaryWithKeysOfClass:objectsOfClass:fromData:error:"), keyCls, valueCls, data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return NSDictionaryFromID(rv), nil

}
// Decodes the \c NSDictionary root object from \c data which should be an \c
// NSDictionary, with keys of the types given in \c keyClasses and objects of
// the given non-collection classes in \c objectClasses (no nested
// dictionaries or other dictionaries contained in the dictionary, etc) from
// the given archive, previously encoded by \c NSKeyedArchiver.
//
// # Discussion
// 
// Enables \c requiresSecureCoding and sets the \c decodingFailurePolicy to \c
// NSDecodingFailurePolicySetErrorAndReturn.
// 
// Returns \c nil if the given data is not valid or cannot be decoded, and
// sets the \c error out parameter.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/unarchivedDictionaryWithKeysOfClasses:objectsOfClasses:fromData:error:
func (_NSKeyedUnarchiverClass NSKeyedUnarchiverClass) UnarchivedDictionaryWithKeysOfClassesObjectsOfClassesFromDataError(keyClasses INSSet, valueClasses INSSet, data INSData) (INSDictionary, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSKeyedUnarchiverClass.class), objc.Sel("unarchivedDictionaryWithKeysOfClasses:objectsOfClasses:fromData:error:"), keyClasses, valueClasses, data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return NSDictionaryFromID(rv), nil

}
// Decodes a previously-archived object graph, that returns the root object as
// the specified type.
//
// cls: The expected class of the root object.
//
// data: An object graph previously encoded by [NSKeyedArchiver].
//
// error: If the return value is `nil`, an [NSError] indicating why the unarchive
// operation failed.
//
// # Return Value
// 
// The decoded root of the object graph, or `nil` if an error occurred.
//
// # Discussion
// 
// This method produces an error if `data` does not contain valid keyed data.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/unarchivedObjectOfClass:fromData:error:
func (_NSKeyedUnarchiverClass NSKeyedUnarchiverClass) UnarchivedObjectOfClassFromDataError(cls objc.Class, data INSData) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSKeyedUnarchiverClass.class), objc.Sel("unarchivedObjectOfClass:fromData:error:"), cls, data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// The receiver’s delegate.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiver/delegate
func (k NSKeyedUnarchiver) Delegate() NSKeyedUnarchiverDelegate {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("delegate"))
	return NSKeyedUnarchiverDelegateObjectFromID(rv)
}
func (k NSKeyedUnarchiver) SetDelegate(value NSKeyedUnarchiverDelegate) {
	objc.Send[struct{}](k.ID, objc.Sel("setDelegate:"), value)
}
// The name of the exception raised by
//
// See: https://developer.apple.com/documentation/foundation/nsexceptionname/invalidunarchiveoperationexception
func (k NSKeyedUnarchiver) InvalidUnarchiveOperationException() NSExceptionName {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("NSInvalidUnarchiveOperationException"))
	return NSExceptionName(NSStringFromID(rv).String())
}
// Name of an exception that occurs when attempting to access outside the
// bounds of some data, such as beyond the end of a string.
//
// See: https://developer.apple.com/documentation/foundation/nsexceptionname/rangeexception
func (k NSKeyedUnarchiver) RangeException() NSExceptionName {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("NSRangeException"))
	return NSExceptionName(NSStringFromID(rv).String())
}

