// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSKeyedArchiver] class.
var (
	_NSKeyedArchiverClass     NSKeyedArchiverClass
	_NSKeyedArchiverClassOnce sync.Once
)

func getNSKeyedArchiverClass() NSKeyedArchiverClass {
	_NSKeyedArchiverClassOnce.Do(func() {
		_NSKeyedArchiverClass = NSKeyedArchiverClass{class: objc.GetClass("NSKeyedArchiver")}
	})
	return _NSKeyedArchiverClass
}

// GetNSKeyedArchiverClass returns the class object for NSKeyedArchiver.
func GetNSKeyedArchiverClass() NSKeyedArchiverClass {
	return getNSKeyedArchiverClass()
}

type NSKeyedArchiverClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSKeyedArchiverClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSKeyedArchiverClass) Alloc() NSKeyedArchiver {
	rv := objc.Send[NSKeyedArchiver](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An encoder that stores an object’s data to an archive referenced by keys.
//
// # Overview
// 
// [NSKeyedArchiver], a concrete subclass of [NSCoder], provides a way to
// encode objects (and scalar values) into an architecture-independent format
// suitable for storage in a file. When you archive a set of objects, the
// archiver writes the class information and instance variables for each
// object to the archive. The companion class [NSKeyedUnarchiver] decodes the
// data in an archive and creates a set of objects equivalent to the original
// set.
// 
// A keyed archive differs from a non-keyed archive in that all the objects
// and values encoded into the archive have names, or keys. When decoding a
// non-keyed archive, the decoder must decode values in the same order the
// original encoder used. When decoding a keyed archive, the decoder requests
// values by name, meaning it can decode values out of sequence or not at all.
// Keyed archives, therefore, provide better support for forward and backward
// compatibility.
// 
// The keys given to encoded values must be unique only within the scope of
// the currently-encoding object. A keyed archive is hierarchical, so the keys
// used by object A to encode its instance variables don’t conflict with the
// keys used by object B. This is true even if A and B are instances of the
// same class. Within a single object, however, the keys used by a subclass
// can conflict with keys used in its superclasses.
// 
// An [NSArchiver] object can write the archive data to a file or to a
// mutable-data object (an instance of [NSMutableData]) that you provide.
//
// [NSArchiver]: https://developer.apple.com/documentation/Foundation/NSArchiver
//
// # Creating a Keyed Archiver
//
//   - [NSKeyedArchiver.InitRequiringSecureCoding]: Creates an archiver to encode data, and optionally disables secure coding.
//
// # Archiving Data
//
//   - [NSKeyedArchiver.FinishEncoding]: Instructs the receiver to construct the final data stream.
//   - [NSKeyedArchiver.EncodedData]: The encoded data for the archiver.
//   - [NSKeyedArchiver.OutputFormat]: The format in which the receiver encodes its data.
//   - [NSKeyedArchiver.SetOutputFormat]
//
// # Managing the Delegate
//
//   - [NSKeyedArchiver.Delegate]: The archiver’s delegate.
//   - [NSKeyedArchiver.SetDelegate]
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver
type NSKeyedArchiver struct {
	NSCoder
}

// NSKeyedArchiverFromID constructs a [NSKeyedArchiver] from an objc.ID.
//
// An encoder that stores an object’s data to an archive referenced by keys.
func NSKeyedArchiverFromID(id objc.ID) NSKeyedArchiver {
	return NSKeyedArchiver{NSCoder: NSCoderFromID(id)}
}
// NOTE: NSKeyedArchiver adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSKeyedArchiver] class.
//
// # Creating a Keyed Archiver
//
//   - [INSKeyedArchiver.InitRequiringSecureCoding]: Creates an archiver to encode data, and optionally disables secure coding.
//
// # Archiving Data
//
//   - [INSKeyedArchiver.FinishEncoding]: Instructs the receiver to construct the final data stream.
//   - [INSKeyedArchiver.EncodedData]: The encoded data for the archiver.
//   - [INSKeyedArchiver.OutputFormat]: The format in which the receiver encodes its data.
//   - [INSKeyedArchiver.SetOutputFormat]
//
// # Managing the Delegate
//
//   - [INSKeyedArchiver.Delegate]: The archiver’s delegate.
//   - [INSKeyedArchiver.SetDelegate]
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver
type INSKeyedArchiver interface {
	INSCoder

	// Topic: Creating a Keyed Archiver

	// Creates an archiver to encode data, and optionally disables secure coding.
	InitRequiringSecureCoding(requiresSecureCoding bool) NSKeyedArchiver

	// Topic: Archiving Data

	// Instructs the receiver to construct the final data stream.
	FinishEncoding()
	// The encoded data for the archiver.
	EncodedData() INSData
	// The format in which the receiver encodes its data.
	OutputFormat() NSPropertyListFormat
	SetOutputFormat(value NSPropertyListFormat)

	// Topic: Managing the Delegate

	// The archiver’s delegate.
	Delegate() NSKeyedArchiverDelegate
	SetDelegate(value NSKeyedArchiverDelegate)
}

// Init initializes the instance.
func (k NSKeyedArchiver) Init() NSKeyedArchiver {
	rv := objc.Send[NSKeyedArchiver](k.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k NSKeyedArchiver) Autorelease() NSKeyedArchiver {
	rv := objc.Send[NSKeyedArchiver](k.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSKeyedArchiver creates a new NSKeyedArchiver instance.
func NewNSKeyedArchiver() NSKeyedArchiver {
	class := getNSKeyedArchiverClass()
	rv := objc.Send[NSKeyedArchiver](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an archiver to encode data, and optionally disables secure coding.
//
// requiresSecureCoding: A Boolean value indicating whether all encoded objects must conform to
// [NSSecureCoding].
//
// # Discussion
// 
// To prevent the possibility of encoding an object that [NSKeyedUnarchiver]
// can’t decode, set `requiresSecureCoding` to [true] whenever possible.
// This ensures that all encoded objects conform to [NSSecureCoding].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/init(requiringSecureCoding:)
func NewKeyedArchiverRequiringSecureCoding(requiresSecureCoding bool) NSKeyedArchiver {
	instance := getNSKeyedArchiverClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initRequiringSecureCoding:"), requiresSecureCoding)
	return NSKeyedArchiverFromID(rv)
}

// Creates an archiver to encode data, and optionally disables secure coding.
//
// requiresSecureCoding: A Boolean value indicating whether all encoded objects must conform to
// [NSSecureCoding].
//
// # Discussion
// 
// To prevent the possibility of encoding an object that [NSKeyedUnarchiver]
// can’t decode, set `requiresSecureCoding` to [true] whenever possible.
// This ensures that all encoded objects conform to [NSSecureCoding].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/init(requiringSecureCoding:)
func (k NSKeyedArchiver) InitRequiringSecureCoding(requiresSecureCoding bool) NSKeyedArchiver {
	rv := objc.Send[NSKeyedArchiver](k.ID, objc.Sel("initRequiringSecureCoding:"), requiresSecureCoding)
	return rv
}
// Instructs the receiver to construct the final data stream.
//
// # Discussion
// 
// No more values can be encoded after this method is called. You must call
// this method when finished.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/finishEncoding()
func (k NSKeyedArchiver) FinishEncoding() {
	objc.Send[objc.ID](k.ID, objc.Sel("finishEncoding"))
}

// Encodes an object graph with the given root object into a data
// representation, optionally requiring secure coding.
//
// object: The root of the object graph to archive.
//
// requiresSecureCoding: A Boolean value indicating whether all encoded objects must conform to
// [NSSecureCoding].
//
// # Discussion
// 
// To prevent the possibility of encoding an object that [NSKeyedUnarchiver]
// can’t decode, set `requiresSecureCoding` to true whenever possible. This
// ensures that all encoded objects conform to [NSSecureCoding].
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/archivedData(withRootObject:requiringSecureCoding:)
func (_NSKeyedArchiverClass NSKeyedArchiverClass) ArchivedDataWithRootObjectRequiringSecureCodingError(object objectivec.IObject, requiresSecureCoding bool) (NSData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSKeyedArchiverClass.class), objc.Sel("archivedDataWithRootObject:requiringSecureCoding:error:"), object, requiresSecureCoding, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSData{}, NSErrorFrom(errorPtr)
	}
	return NSDataFromID(rv), nil

}
// Sets a global translation mapping to encode instances of a given class with
// the provided name, rather than their real name.
//
// codedName: The name of the class that [NSKeyedArchiver] uses in place of `cls`.
//
// cls: The class for which to set up a translation mapping.
//
// # Discussion
// 
// When encoding, an archiver consults its own translation map before using
// the class’ translation map.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/setClassName(_:for:)-swift.type.method
func (_NSKeyedArchiverClass NSKeyedArchiverClass) SetClassNameForClass(codedName string, cls objc.Class) {
	objc.Send[objc.ID](objc.ID(_NSKeyedArchiverClass.class), objc.Sel("setClassName:forClass:"), objc.String(codedName), cls)
}
// Returns the class name with which the archiver class encodes instances of a
// given class.
//
// cls: The class for which to determine the translation mapping.
//
// # Return Value
// 
// The class name with which [NSKeyedArchiver] encodes instances of `cls`.
// Returns `nil` if [NSKeyedArchiver] does not have a translation mapping for
// `cls`.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/className(for:)-swift.type.method
func (_NSKeyedArchiverClass NSKeyedArchiverClass) ClassNameForClassWithCls(cls objc.Class) string {
	rv := objc.Send[objc.ID](objc.ID(_NSKeyedArchiverClass.class), objc.Sel("classNameForClass:"), cls)
	return NSStringFromID(rv).String()
}

// The encoded data for the archiver.
//
// # Discussion
// 
// If encoding has not yet finished, invoking this property calls
// [FinishEncoding] and populates this property with the encoded data. If you
// initialized the keyed archiver with [init(forWritingWith:)] and a specific
// mutable data instance, this property contains that instance.
//
// [init(forWritingWith:)]: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/init(forWritingWith:)
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/encodedData
func (k NSKeyedArchiver) EncodedData() INSData {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("encodedData"))
	return NSDataFromID(objc.ID(rv))
}
// The format in which the receiver encodes its data.
//
// # Discussion
// 
// The available formats are
// [PropertyListSerialization.PropertyListFormat.xml] and
// [PropertyListSerialization.PropertyListFormat.binary].
//
// [PropertyListSerialization.PropertyListFormat.binary]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/PropertyListFormat/binary
// [PropertyListSerialization.PropertyListFormat.xml]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/PropertyListFormat/xml
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/outputFormat
func (k NSKeyedArchiver) OutputFormat() NSPropertyListFormat {
	rv := objc.Send[NSPropertyListFormat](k.ID, objc.Sel("outputFormat"))
	return NSPropertyListFormat(rv)
}
func (k NSKeyedArchiver) SetOutputFormat(value NSPropertyListFormat) {
	objc.Send[struct{}](k.ID, objc.Sel("setOutputFormat:"), value)
}
// The archiver’s delegate.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedArchiver/delegate
func (k NSKeyedArchiver) Delegate() NSKeyedArchiverDelegate {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("delegate"))
	return NSKeyedArchiverDelegateObjectFromID(rv)
}
func (k NSKeyedArchiver) SetDelegate(value NSKeyedArchiverDelegate) {
	objc.Send[struct{}](k.ID, objc.Sel("setDelegate:"), value)
}

