// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSAppleEventDescriptor] class.
var (
	_NSAppleEventDescriptorClass     NSAppleEventDescriptorClass
	_NSAppleEventDescriptorClassOnce sync.Once
)

func getNSAppleEventDescriptorClass() NSAppleEventDescriptorClass {
	_NSAppleEventDescriptorClassOnce.Do(func() {
		_NSAppleEventDescriptorClass = NSAppleEventDescriptorClass{class: objc.GetClass("NSAppleEventDescriptor")}
	})
	return _NSAppleEventDescriptorClass
}

// GetNSAppleEventDescriptorClass returns the class object for NSAppleEventDescriptor.
func GetNSAppleEventDescriptorClass() NSAppleEventDescriptorClass {
	return getNSAppleEventDescriptorClass()
}

type NSAppleEventDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAppleEventDescriptorClass) Alloc() NSAppleEventDescriptor {
	rv := objc.Send[NSAppleEventDescriptor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A wrapper for the Apple event descriptor data type.
//
// # Overview
// 
// An instance of [NSAppleEventDescriptor] represents a descriptor—the basic
// building block for Apple events. This class is a wrapper for the underlying
// Apple event descriptor data type, [AEDesc]. Scriptable Cocoa applications
// frequently work with instances of [NSAppleEventDescriptor], but should
// rarely need to work directly with the [AEDesc] data structure.
// 
// A is a data structure that stores data and an accompanying four-character
// code. A descriptor can store a value, or it can store a list of other
// descriptors (which may also be lists). All the information in an Apple
// event is stored in descriptors and lists of descriptors, and every Apple
// event is itself a descriptor list that matches certain criteria.
// 
// Descriptors can be used to build arbitrarily complex containers, so that
// one Apple event can represent a script statement such as `tell application
// "TextEdit" to get word 3 of paragraph 6 of document 3`.
// 
// In working with Apple event descriptors, it can be useful to understand
// some of the underlying data types. You’ll find terms such as descriptor,
// descriptor list, Apple event record, and Apple event defined in Building an
// Apple Event in Apple Events Programming Guide. You’ll also find
// information on the four-character codes used to identify information within
// a descriptor. Apple event data types are defined in [Apple Event Manager].
// The values of many four-character codes used by Apple (and in some cases
// reused by developers) can be found in [AppleScript Terminology and Apple
// Event Codes].
// 
// The most common reason to construct an Apple event with an instance of
// [NSAppleEventDescriptor] is to supply information in a return Apple event.
// The most common situation where you might need to extract information from
// an Apple event (as an instance of [NSAppleEventDescriptor]) is when an
// Apple event handler installed by your application is invoked, as described
// in “Installing an Apple Event Handler” in [How Cocoa Applications
// Handle Apple Events]. In addition, if you execute an AppleScript script
// using the [NSAppleScript] class, you get an instance of
// [NSAppleEventDescriptor] as the return value, from which you can extract
// any required information.
// 
// When you work with an instance of [NSAppleEventDescriptor], you can access
// the underlying descriptor directly, if necessary, with the [NSAppleEventDescriptor.AeDesc] method.
// Other methods, including [NSAppleEventDescriptor.DescriptorWithDescriptorTypeBytesLength] make it
// possible to create and initialize instances of [NSAppleEventDescriptor]
// without creating temporary instances of [NSData].
// 
// The designated initializer for [NSAppleEventDescriptor] is
// [NSAppleEventDescriptor.InitWithAEDescNoCopy]. However, it is unlikely that you will need to
// create a subclass of [NSAppleEventDescriptor].
// 
// Cocoa doesn’t currently provide a mechanism for applications to directly
// send raw Apple events (though compiling and executing an AppleScript script
// with [NSAppleScript] may result in Apple events being sent). However, Cocoa
// applications have full access to the Apple Event Manager C APIs for working
// with Apple events. So, for example, you might use an instance of
// [NSAppleEventDescriptor] to assemble an Apple event and call the Apple
// Event Manager function `AESend(_:_:_:_:_:_:_:)` to send it.
// 
// If you need to send Apple events, or if you need more information on some
// of the Apple event concepts described here, see Apple Events Programming
// Guide and [Apple Event Manager].
//
// [AEDesc]: https://developer.apple.com/documentation/coreservices/aedesc
// [Apple Event Manager]: https://developer.apple.com/documentation/applicationservices/apple_event_manager
// [AppleScript Terminology and Apple Event Codes]: http://developer.apple.com/releasenotes/AppleScript/ASTerminology_AppleEventCodes/TermsAndCodes.html
// [How Cocoa Applications Handle Apple Events]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ScriptableCocoaApplications/SApps_handle_AEs/SAppsHandleAEs.html#//apple_ref/doc/uid/20001239
//
// # Creating and Initializing Descriptors
//
//   - [NSAppleEventDescriptor.InitListDescriptor]: Initializes a newly allocated instance as an empty list descriptor.
//   - [NSAppleEventDescriptor.InitRecordDescriptor]: Initializes a newly allocated instance as a descriptor that is an Apple event record.
//   - [NSAppleEventDescriptor.InitWithAEDescNoCopy]: Initializes a newly allocated instance as a descriptor for the specified Carbon [AEDesc] structure.
//   - [NSAppleEventDescriptor.InitWithDescriptorTypeBytesLength]: Initializes a newly allocated instance as a descriptor with the specified descriptor type and data (from an arbitrary sequence of bytes and a length count).
//   - [NSAppleEventDescriptor.InitWithDescriptorTypeData]: Initializes a newly allocated instance as a descriptor with the specified descriptor type and data (from an instance of [NSData]).
//   - [NSAppleEventDescriptor.InitWithEventClassEventIDTargetDescriptorReturnIDTransactionID]: Initializes a newly allocated instance as a descriptor for an Apple event, initialized with the specified values.
//
// # Getting Information About a Descriptor
//
//   - [NSAppleEventDescriptor.AeDesc]: The [AEDesc] structure encapsulated by the receiver, if it has one.
//   - [NSAppleEventDescriptor.BooleanValue]: The contents of the receiver as a Boolean value, coercing (to `typeBoolean`) if necessary.
//   - [NSAppleEventDescriptor.CoerceToDescriptorType]: Returns a descriptor obtained by coercing the receiver to the specified type.
//   - [NSAppleEventDescriptor.Data]: The receiver’s data.
//   - [NSAppleEventDescriptor.DescriptorType]: The descriptor type of the receiver.
//   - [NSAppleEventDescriptor.EnumCodeValue]: The contents of the receiver as an enumeration type, coercing to `typeEnumerated` if necessary.
//   - [NSAppleEventDescriptor.Int32Value]: The contents of the receiver as an integer, coercing (to `typeSInt32`) if necessary.
//   - [NSAppleEventDescriptor.NumberOfItems]: The number of descriptors in the receiver’s descriptor list.
//   - [NSAppleEventDescriptor.StringValue]: The contents of the receiver as a Unicode text string, coercing to `typeUnicodeText` if necessary.
//   - [NSAppleEventDescriptor.TypeCodeValue]: The contents of the receiver as a type, coercing to `typeType` if necessary.
//
// # Working With List Descriptors
//
//   - [NSAppleEventDescriptor.DescriptorAtIndex]: Returns the descriptor at the specified (one-based) position in the receiving descriptor list.
//   - [NSAppleEventDescriptor.InsertDescriptorAtIndex]: Inserts a descriptor at the specified (one-based) position in the receiving descriptor list, replacing the existing descriptor, if any, at that position.
//   - [NSAppleEventDescriptor.RemoveDescriptorAtIndex]: Removes the descriptor at the specified (one-based) position in the receiving descriptor list.
//
// # Working With Record Descriptors
//
//   - [NSAppleEventDescriptor.DescriptorForKeyword]: Returns the receiver’s descriptor for the specified keyword.
//   - [NSAppleEventDescriptor.KeywordForDescriptorAtIndex]: Returns the keyword for the descriptor at the specified (one-based) position in the receiver.
//   - [NSAppleEventDescriptor.RemoveDescriptorWithKeyword]: Removes the receiver’s descriptor identified by the specified keyword.
//   - [NSAppleEventDescriptor.SetDescriptorForKeyword]: Adds a descriptor, identified by a keyword, to the receiver.
//
// # Working With Apple Event Descriptors
//
//   - [NSAppleEventDescriptor.AttributeDescriptorForKeyword]: Returns a descriptor for the receiver’s Apple event attribute identified by the specified keyword.
//   - [NSAppleEventDescriptor.EventClass]: The event class for the receiver.
//   - [NSAppleEventDescriptor.EventID]: The event ID for the receiver.
//   - [NSAppleEventDescriptor.ParamDescriptorForKeyword]: Returns a descriptor for the receiver’s Apple event parameter identified by the specified keyword.
//   - [NSAppleEventDescriptor.RemoveParamDescriptorWithKeyword]: Removes the receiver’s parameter descriptor identified by the specified keyword.
//   - [NSAppleEventDescriptor.ReturnID]: The receiver’s return ID (the ID for a reply Apple event).
//   - [NSAppleEventDescriptor.SetAttributeDescriptorForKeyword]: Adds a descriptor to the receiver as an attribute identified by the specified keyword.
//   - [NSAppleEventDescriptor.SetParamDescriptorForKeyword]: Adds a descriptor to the receiver as an Apple event parameter identified by the specified keyword.
//   - [NSAppleEventDescriptor.TransactionID]: The receiver’s transaction ID, if any.
//
// # Instance Properties
//
//   - [NSAppleEventDescriptor.DateValue]
//   - [NSAppleEventDescriptor.DoubleValue]
//   - [NSAppleEventDescriptor.FileURLValue]
//   - [NSAppleEventDescriptor.IsRecordDescriptor]
//
// # Instance Methods
//
//   - [NSAppleEventDescriptor.SendEventWithOptionsTimeoutError]
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor
type NSAppleEventDescriptor struct {
	objectivec.Object
}

// NSAppleEventDescriptorFromID constructs a [NSAppleEventDescriptor] from an objc.ID.
//
// A wrapper for the Apple event descriptor data type.
func NSAppleEventDescriptorFromID(id objc.ID) NSAppleEventDescriptor {
	return NSAppleEventDescriptor{objectivec.Object{id}}
}
// NOTE: NSAppleEventDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSAppleEventDescriptor] class.
//
// # Creating and Initializing Descriptors
//
//   - [INSAppleEventDescriptor.InitListDescriptor]: Initializes a newly allocated instance as an empty list descriptor.
//   - [INSAppleEventDescriptor.InitRecordDescriptor]: Initializes a newly allocated instance as a descriptor that is an Apple event record.
//   - [INSAppleEventDescriptor.InitWithAEDescNoCopy]: Initializes a newly allocated instance as a descriptor for the specified Carbon [AEDesc] structure.
//   - [INSAppleEventDescriptor.InitWithDescriptorTypeBytesLength]: Initializes a newly allocated instance as a descriptor with the specified descriptor type and data (from an arbitrary sequence of bytes and a length count).
//   - [INSAppleEventDescriptor.InitWithDescriptorTypeData]: Initializes a newly allocated instance as a descriptor with the specified descriptor type and data (from an instance of [NSData]).
//   - [INSAppleEventDescriptor.InitWithEventClassEventIDTargetDescriptorReturnIDTransactionID]: Initializes a newly allocated instance as a descriptor for an Apple event, initialized with the specified values.
//
// # Getting Information About a Descriptor
//
//   - [INSAppleEventDescriptor.AeDesc]: The [AEDesc] structure encapsulated by the receiver, if it has one.
//   - [INSAppleEventDescriptor.BooleanValue]: The contents of the receiver as a Boolean value, coercing (to `typeBoolean`) if necessary.
//   - [INSAppleEventDescriptor.CoerceToDescriptorType]: Returns a descriptor obtained by coercing the receiver to the specified type.
//   - [INSAppleEventDescriptor.Data]: The receiver’s data.
//   - [INSAppleEventDescriptor.DescriptorType]: The descriptor type of the receiver.
//   - [INSAppleEventDescriptor.EnumCodeValue]: The contents of the receiver as an enumeration type, coercing to `typeEnumerated` if necessary.
//   - [INSAppleEventDescriptor.Int32Value]: The contents of the receiver as an integer, coercing (to `typeSInt32`) if necessary.
//   - [INSAppleEventDescriptor.NumberOfItems]: The number of descriptors in the receiver’s descriptor list.
//   - [INSAppleEventDescriptor.StringValue]: The contents of the receiver as a Unicode text string, coercing to `typeUnicodeText` if necessary.
//   - [INSAppleEventDescriptor.TypeCodeValue]: The contents of the receiver as a type, coercing to `typeType` if necessary.
//
// # Working With List Descriptors
//
//   - [INSAppleEventDescriptor.DescriptorAtIndex]: Returns the descriptor at the specified (one-based) position in the receiving descriptor list.
//   - [INSAppleEventDescriptor.InsertDescriptorAtIndex]: Inserts a descriptor at the specified (one-based) position in the receiving descriptor list, replacing the existing descriptor, if any, at that position.
//   - [INSAppleEventDescriptor.RemoveDescriptorAtIndex]: Removes the descriptor at the specified (one-based) position in the receiving descriptor list.
//
// # Working With Record Descriptors
//
//   - [INSAppleEventDescriptor.DescriptorForKeyword]: Returns the receiver’s descriptor for the specified keyword.
//   - [INSAppleEventDescriptor.KeywordForDescriptorAtIndex]: Returns the keyword for the descriptor at the specified (one-based) position in the receiver.
//   - [INSAppleEventDescriptor.RemoveDescriptorWithKeyword]: Removes the receiver’s descriptor identified by the specified keyword.
//   - [INSAppleEventDescriptor.SetDescriptorForKeyword]: Adds a descriptor, identified by a keyword, to the receiver.
//
// # Working With Apple Event Descriptors
//
//   - [INSAppleEventDescriptor.AttributeDescriptorForKeyword]: Returns a descriptor for the receiver’s Apple event attribute identified by the specified keyword.
//   - [INSAppleEventDescriptor.EventClass]: The event class for the receiver.
//   - [INSAppleEventDescriptor.EventID]: The event ID for the receiver.
//   - [INSAppleEventDescriptor.ParamDescriptorForKeyword]: Returns a descriptor for the receiver’s Apple event parameter identified by the specified keyword.
//   - [INSAppleEventDescriptor.RemoveParamDescriptorWithKeyword]: Removes the receiver’s parameter descriptor identified by the specified keyword.
//   - [INSAppleEventDescriptor.ReturnID]: The receiver’s return ID (the ID for a reply Apple event).
//   - [INSAppleEventDescriptor.SetAttributeDescriptorForKeyword]: Adds a descriptor to the receiver as an attribute identified by the specified keyword.
//   - [INSAppleEventDescriptor.SetParamDescriptorForKeyword]: Adds a descriptor to the receiver as an Apple event parameter identified by the specified keyword.
//   - [INSAppleEventDescriptor.TransactionID]: The receiver’s transaction ID, if any.
//
// # Instance Properties
//
//   - [INSAppleEventDescriptor.DateValue]
//   - [INSAppleEventDescriptor.DoubleValue]
//   - [INSAppleEventDescriptor.FileURLValue]
//   - [INSAppleEventDescriptor.IsRecordDescriptor]
//
// # Instance Methods
//
//   - [INSAppleEventDescriptor.SendEventWithOptionsTimeoutError]
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor
type INSAppleEventDescriptor interface {
	objectivec.IObject
	NSCopying

	// Topic: Creating and Initializing Descriptors

	// Initializes a newly allocated instance as an empty list descriptor.
	InitListDescriptor() NSAppleEventDescriptor
	// Initializes a newly allocated instance as a descriptor that is an Apple event record.
	InitRecordDescriptor() NSAppleEventDescriptor
	// Initializes a newly allocated instance as a descriptor for the specified Carbon [AEDesc] structure.
	InitWithAEDescNoCopy(aeDesc objectivec.IObject) NSAppleEventDescriptor
	// Initializes a newly allocated instance as a descriptor with the specified descriptor type and data (from an arbitrary sequence of bytes and a length count).
	InitWithDescriptorTypeBytesLength(descriptorType uint32, bytes unsafe.Pointer, byteCount uint) NSAppleEventDescriptor
	// Initializes a newly allocated instance as a descriptor with the specified descriptor type and data (from an instance of [NSData]).
	InitWithDescriptorTypeData(descriptorType uint32, data INSData) NSAppleEventDescriptor
	// Initializes a newly allocated instance as a descriptor for an Apple event, initialized with the specified values.
	InitWithEventClassEventIDTargetDescriptorReturnIDTransactionID(eventClass uint32, eventID uint32, targetDescriptor INSAppleEventDescriptor, returnID int16, transactionID int32) NSAppleEventDescriptor

	// Topic: Getting Information About a Descriptor

	// The [AEDesc] structure encapsulated by the receiver, if it has one.
	AeDesc() objectivec.IObject
	// The contents of the receiver as a Boolean value, coercing (to `typeBoolean`) if necessary.
	BooleanValue() bool
	// Returns a descriptor obtained by coercing the receiver to the specified type.
	CoerceToDescriptorType(descriptorType uint32) INSAppleEventDescriptor
	// The receiver’s data.
	Data() INSData
	// The descriptor type of the receiver.
	DescriptorType() uint32
	// The contents of the receiver as an enumeration type, coercing to `typeEnumerated` if necessary.
	EnumCodeValue() uint32
	// The contents of the receiver as an integer, coercing (to `typeSInt32`) if necessary.
	Int32Value() int32
	// The number of descriptors in the receiver’s descriptor list.
	NumberOfItems() int
	// The contents of the receiver as a Unicode text string, coercing to `typeUnicodeText` if necessary.
	StringValue() string
	// The contents of the receiver as a type, coercing to `typeType` if necessary.
	TypeCodeValue() uint32

	// Topic: Working With List Descriptors

	// Returns the descriptor at the specified (one-based) position in the receiving descriptor list.
	DescriptorAtIndex(index int) INSAppleEventDescriptor
	// Inserts a descriptor at the specified (one-based) position in the receiving descriptor list, replacing the existing descriptor, if any, at that position.
	InsertDescriptorAtIndex(descriptor INSAppleEventDescriptor, index int)
	// Removes the descriptor at the specified (one-based) position in the receiving descriptor list.
	RemoveDescriptorAtIndex(index int)

	// Topic: Working With Record Descriptors

	// Returns the receiver’s descriptor for the specified keyword.
	DescriptorForKeyword(keyword uint32) INSAppleEventDescriptor
	// Returns the keyword for the descriptor at the specified (one-based) position in the receiver.
	KeywordForDescriptorAtIndex(index int) uint32
	// Removes the receiver’s descriptor identified by the specified keyword.
	RemoveDescriptorWithKeyword(keyword uint32)
	// Adds a descriptor, identified by a keyword, to the receiver.
	SetDescriptorForKeyword(descriptor INSAppleEventDescriptor, keyword uint32)

	// Topic: Working With Apple Event Descriptors

	// Returns a descriptor for the receiver’s Apple event attribute identified by the specified keyword.
	AttributeDescriptorForKeyword(keyword uint32) INSAppleEventDescriptor
	// The event class for the receiver.
	EventClass() uint32
	// The event ID for the receiver.
	EventID() uint32
	// Returns a descriptor for the receiver’s Apple event parameter identified by the specified keyword.
	ParamDescriptorForKeyword(keyword uint32) INSAppleEventDescriptor
	// Removes the receiver’s parameter descriptor identified by the specified keyword.
	RemoveParamDescriptorWithKeyword(keyword uint32)
	// The receiver’s return ID (the ID for a reply Apple event).
	ReturnID() int16
	// Adds a descriptor to the receiver as an attribute identified by the specified keyword.
	SetAttributeDescriptorForKeyword(descriptor INSAppleEventDescriptor, keyword uint32)
	// Adds a descriptor to the receiver as an Apple event parameter identified by the specified keyword.
	SetParamDescriptorForKeyword(descriptor INSAppleEventDescriptor, keyword uint32)
	// The receiver’s transaction ID, if any.
	TransactionID() int32

	// Topic: Instance Properties

	DateValue() INSDate
	DoubleValue() float64
	FileURLValue() INSURL
	IsRecordDescriptor() bool

	// Topic: Instance Methods

	SendEventWithOptionsTimeoutError(sendOptions NSAppleEventSendOptions, timeoutInSeconds float64) (INSAppleEventDescriptor, error)

	// Encodes the receiver using a given archiver.
	EncodeWithCoder(coder INSCoder)
	InitWithCoder(coder INSCoder) NSAppleEventDescriptor
}





// Init initializes the instance.
func (a NSAppleEventDescriptor) Init() NSAppleEventDescriptor {
	rv := objc.Send[NSAppleEventDescriptor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAppleEventDescriptor) Autorelease() NSAppleEventDescriptor {
	rv := objc.Send[NSAppleEventDescriptor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAppleEventDescriptor creates a new NSAppleEventDescriptor instance.
func NewNSAppleEventDescriptor() NSAppleEventDescriptor {
	class := getNSAppleEventDescriptorClass()
	rv := objc.Send[NSAppleEventDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Initializes a newly allocated instance as an empty list descriptor.
//
// # Return Value
// 
// An empty list descriptor, or `nil` if an error occurs.
//
// # Discussion
// 
// You can add items to the empty list descriptor with
// [InsertDescriptorAtIndex]. The list indices are one-based.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(listDescriptor:)
func NewAppleEventDescriptorListDescriptor() NSAppleEventDescriptor {
	instance := getNSAppleEventDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initListDescriptor"))
	return NSAppleEventDescriptorFromID(rv)
}


// Initializes a newly allocated instance as a descriptor that is an Apple
// event record.
//
// # Return Value
// 
// The initialized Apple event record, or `nil` if an error occurs.
//
// # Discussion
// 
// An Apple event record is a descriptor whose data is a set of descriptors
// keyed by four-character codes. You can add information to the descriptor
// with methods such as [SetAttributeDescriptorForKeyword],
// [SetDescriptorForKeyword], and [SetParamDescriptorForKeyword].
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(recordDescriptor:)
func NewAppleEventDescriptorRecordDescriptor() NSAppleEventDescriptor {
	instance := getNSAppleEventDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initRecordDescriptor"))
	return NSAppleEventDescriptorFromID(rv)
}


// Initializes a newly allocated instance as a descriptor for the specified
// Carbon [AEDesc] structure.
//
// aeDesc: A pointer to the [AEDesc] structure to associate with the descriptor.
//
// # Return Value
// 
// An instance of [NSAppleEventDescriptor] that is associated with the
// structure pointed to by `aeDesc`, or `nil` if an error occurs.
//
// # Discussion
// 
// The initialized object takes responsibility for calling the [AEDisposeDesc]
// function on the [AEDesc] at object deallocation time. This is the
// designated initializer for this class.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(aeDescNoCopy:)
func NewAppleEventDescriptorWithAEDescNoCopy(aeDesc objectivec.IObject) NSAppleEventDescriptor {
	instance := getNSAppleEventDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAEDescNoCopy:"), aeDesc)
	return NSAppleEventDescriptorFromID(rv)
}


//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(applicationURL:)
func NewAppleEventDescriptorWithApplicationURL(applicationURL INSURL) NSAppleEventDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getNSAppleEventDescriptorClass().class), objc.Sel("descriptorWithApplicationURL:"), applicationURL)
	return NSAppleEventDescriptorFromID(rv)
}


// Creates a descriptor initialized with type `typeBoolean` that stores the
// specified Boolean value.
//
// boolean: The Boolean value to be set in the returned descriptor.
//
// # Return Value
// 
// A descriptor with the specified Boolean value, or `nil` if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(boolean:)
func NewAppleEventDescriptorWithBoolean(boolean bool) NSAppleEventDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getNSAppleEventDescriptorClass().class), objc.Sel("descriptorWithBoolean:"), boolean)
	return NSAppleEventDescriptorFromID(rv)
}


//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(bundleIdentifier:)
func NewAppleEventDescriptorWithBundleIdentifier(bundleIdentifier string) NSAppleEventDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getNSAppleEventDescriptorClass().class), objc.Sel("descriptorWithBundleIdentifier:"), objc.String(bundleIdentifier))
	return NSAppleEventDescriptorFromID(rv)
}


//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewAppleEventDescriptorWithCoder(coder INSCoder) NSAppleEventDescriptor {
	instance := getNSAppleEventDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSAppleEventDescriptorFromID(rv)
}


//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(date:)
func NewAppleEventDescriptorWithDate(date INSDate) NSAppleEventDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getNSAppleEventDescriptorClass().class), objc.Sel("descriptorWithDate:"), date)
	return NSAppleEventDescriptorFromID(rv)
}


// Initializes a newly allocated instance as a descriptor with the specified
// descriptor type and data (from an arbitrary sequence of bytes and a length
// count).
//
// descriptorType: The descriptor type to be set in the returned descriptor.
//
// bytes: The data, as a sequence of bytes, to be set in the returned descriptor.
//
// byteCount: The length, in bytes, of the data to be set in the returned descriptor.
//
// # Return Value
// 
// An instance of [NSAppleEventDescriptor] with the specified type and data.
// Returns `nil` if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(descriptorType:bytes:length:)
func NewAppleEventDescriptorWithDescriptorTypeBytesLength(descriptorType uint32, bytes unsafe.Pointer, byteCount uint) NSAppleEventDescriptor {
	instance := getNSAppleEventDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescriptorType:bytes:length:"), descriptorType, bytes, byteCount)
	return NSAppleEventDescriptorFromID(rv)
}


// Initializes a newly allocated instance as a descriptor with the specified
// descriptor type and data (from an instance of [NSData]).
//
// descriptorType: The descriptor type to be set in the initialized descriptor.
//
// data: The data to be set in the initialized descriptor.
//
// # Return Value
// 
// An instance of [NSAppleEventDescriptor] with the specified type and data.
// Returns `nil` if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(descriptorType:data:)
func NewAppleEventDescriptorWithDescriptorTypeData(descriptorType uint32, data INSData) NSAppleEventDescriptor {
	instance := getNSAppleEventDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescriptorType:data:"), descriptorType, data)
	return NSAppleEventDescriptorFromID(rv)
}


//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(double:)
func NewAppleEventDescriptorWithDouble(doubleValue float64) NSAppleEventDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getNSAppleEventDescriptorClass().class), objc.Sel("descriptorWithDouble:"), doubleValue)
	return NSAppleEventDescriptorFromID(rv)
}


// Creates a descriptor initialized with type `typeEnumerated` that stores the
// specified enumerator data type value.
//
// enumerator: A type code that identifies the type of enumerated data to be stored in the
// returned descriptor.
//
// # Return Value
// 
// A descriptor with the specified enumerator data type value, or `nil` if an
// error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(enumCode:)
func NewAppleEventDescriptorWithEnumCode(enumerator uint32) NSAppleEventDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getNSAppleEventDescriptorClass().class), objc.Sel("descriptorWithEnumCode:"), enumerator)
	return NSAppleEventDescriptorFromID(rv)
}


// Initializes a newly allocated instance as a descriptor for an Apple event,
// initialized with the specified values.
//
// eventClass: The event class to be set in the returned descriptor.
//
// eventID: The event ID to be set in the returned descriptor.
//
// targetDescriptor: A pointer to a descriptor that identifies the target application for the
// Apple event. Passing `nil` results in an Apple event descriptor that has no
// `keyAddressAttr` attribute (it is valid for an Apple event to have no
// target address attribute).
//
// returnID: The return ID to be set in the returned descriptor. If you pass a value of
// `kAutoGenerateReturnID`, the Apple Event Manager assigns the created Apple
// event a return ID that is unique to the current session. If you pass any
// other value, the Apple Event Manager assigns that value for the ID.
//
// transactionID: The transaction ID to be set in the returned descriptor. A transaction is a
// sequence of Apple events that are sent back and forth between client and
// server applications, beginning with the client’s initial request for a
// service. All Apple events that are part of a transaction must have the same
// transaction ID. You can specify `kAnyTransactionID` if the Apple event is
// not one of a series of interdependent Apple events.
//
// # Return Value
// 
// The initialized Apple event (an instance of [NSAppleEventDescriptor]), or
// `nil` if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(eventClass:eventID:targetDescriptor:returnID:transactionID:)
func NewAppleEventDescriptorWithEventClassEventIDTargetDescriptorReturnIDTransactionID(eventClass uint32, eventID uint32, targetDescriptor INSAppleEventDescriptor, returnID int16, transactionID int32) NSAppleEventDescriptor {
	instance := getNSAppleEventDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEventClass:eventID:targetDescriptor:returnID:transactionID:"), eventClass, eventID, targetDescriptor, returnID, transactionID)
	return NSAppleEventDescriptorFromID(rv)
}


//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(fileURL:)
func NewAppleEventDescriptorWithFileURL(fileURL INSURL) NSAppleEventDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getNSAppleEventDescriptorClass().class), objc.Sel("descriptorWithFileURL:"), fileURL)
	return NSAppleEventDescriptorFromID(rv)
}


// Creates a descriptor initialized with Apple event type `typeSInt32` that
// stores the specified integer value.
//
// signedInt: The integer value to be stored in the returned descriptor.
//
// # Return Value
// 
// A descriptor containing the specified integer value, or `nil` if an error
// occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(int32:)
func NewAppleEventDescriptorWithInt32(signedInt int32) NSAppleEventDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getNSAppleEventDescriptorClass().class), objc.Sel("descriptorWithInt32:"), signedInt)
	return NSAppleEventDescriptorFromID(rv)
}


//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(processIdentifier:)
func NewAppleEventDescriptorWithProcessIdentifier(processIdentifier int32) NSAppleEventDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getNSAppleEventDescriptorClass().class), objc.Sel("descriptorWithProcessIdentifier:"), processIdentifier)
	return NSAppleEventDescriptorFromID(rv)
}


// Creates a descriptor initialized with type `typeUnicodeText` that stores
// the text from the specified string.
//
// string: A string that specifies the text to be stored in the returned descriptor.
//
// # Return Value
// 
// A descriptor that contains the text from the specified string, or `nil` if
// an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(string:)
func NewAppleEventDescriptorWithString(string_ string) NSAppleEventDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getNSAppleEventDescriptorClass().class), objc.Sel("descriptorWithString:"), objc.String(string_))
	return NSAppleEventDescriptorFromID(rv)
}


// Creates a descriptor initialized with type `typeType` that stores the
// specified type value.
//
// typeCode: The type value to be set in the returned descriptor.
//
// # Return Value
// 
// A descriptor with the specified type, or `nil` if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(typeCode:)
func NewAppleEventDescriptorWithTypeCode(typeCode uint32) NSAppleEventDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getNSAppleEventDescriptorClass().class), objc.Sel("descriptorWithTypeCode:"), typeCode)
	return NSAppleEventDescriptorFromID(rv)
}







// Initializes a newly allocated instance as an empty list descriptor.
//
// # Return Value
// 
// An empty list descriptor, or `nil` if an error occurs.
//
// # Discussion
// 
// You can add items to the empty list descriptor with
// [InsertDescriptorAtIndex]. The list indices are one-based.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(listDescriptor:)
func (a NSAppleEventDescriptor) InitListDescriptor() NSAppleEventDescriptor {
	rv := objc.Send[NSAppleEventDescriptor](a.ID, objc.Sel("initListDescriptor"))
	return rv
}

// Initializes a newly allocated instance as a descriptor that is an Apple
// event record.
//
// # Return Value
// 
// The initialized Apple event record, or `nil` if an error occurs.
//
// # Discussion
// 
// An Apple event record is a descriptor whose data is a set of descriptors
// keyed by four-character codes. You can add information to the descriptor
// with methods such as [SetAttributeDescriptorForKeyword],
// [SetDescriptorForKeyword], and [SetParamDescriptorForKeyword].
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(recordDescriptor:)
func (a NSAppleEventDescriptor) InitRecordDescriptor() NSAppleEventDescriptor {
	rv := objc.Send[NSAppleEventDescriptor](a.ID, objc.Sel("initRecordDescriptor"))
	return rv
}

// Initializes a newly allocated instance as a descriptor for the specified
// Carbon [AEDesc] structure.
//
// aeDesc: A pointer to the [AEDesc] structure to associate with the descriptor.
//
// # Return Value
// 
// An instance of [NSAppleEventDescriptor] that is associated with the
// structure pointed to by `aeDesc`, or `nil` if an error occurs.
//
// # Discussion
// 
// The initialized object takes responsibility for calling the [AEDisposeDesc]
// function on the [AEDesc] at object deallocation time. This is the
// designated initializer for this class.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(aeDescNoCopy:)
func (a NSAppleEventDescriptor) InitWithAEDescNoCopy(aeDesc objectivec.IObject) NSAppleEventDescriptor {
	rv := objc.Send[NSAppleEventDescriptor](a.ID, objc.Sel("initWithAEDescNoCopy:"), aeDesc)
	return rv
}

// Initializes a newly allocated instance as a descriptor with the specified
// descriptor type and data (from an arbitrary sequence of bytes and a length
// count).
//
// descriptorType: The descriptor type to be set in the returned descriptor.
//
// bytes: The data, as a sequence of bytes, to be set in the returned descriptor.
//
// byteCount: The length, in bytes, of the data to be set in the returned descriptor.
//
// # Return Value
// 
// An instance of [NSAppleEventDescriptor] with the specified type and data.
// Returns `nil` if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(descriptorType:bytes:length:)
func (a NSAppleEventDescriptor) InitWithDescriptorTypeBytesLength(descriptorType uint32, bytes unsafe.Pointer, byteCount uint) NSAppleEventDescriptor {
	rv := objc.Send[NSAppleEventDescriptor](a.ID, objc.Sel("initWithDescriptorType:bytes:length:"), descriptorType, bytes, byteCount)
	return rv
}

// Initializes a newly allocated instance as a descriptor with the specified
// descriptor type and data (from an instance of [NSData]).
//
// descriptorType: The descriptor type to be set in the initialized descriptor.
//
// data: The data to be set in the initialized descriptor.
//
// # Return Value
// 
// An instance of [NSAppleEventDescriptor] with the specified type and data.
// Returns `nil` if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(descriptorType:data:)
func (a NSAppleEventDescriptor) InitWithDescriptorTypeData(descriptorType uint32, data INSData) NSAppleEventDescriptor {
	rv := objc.Send[NSAppleEventDescriptor](a.ID, objc.Sel("initWithDescriptorType:data:"), descriptorType, data)
	return rv
}

// Initializes a newly allocated instance as a descriptor for an Apple event,
// initialized with the specified values.
//
// eventClass: The event class to be set in the returned descriptor.
//
// eventID: The event ID to be set in the returned descriptor.
//
// targetDescriptor: A pointer to a descriptor that identifies the target application for the
// Apple event. Passing `nil` results in an Apple event descriptor that has no
// `keyAddressAttr` attribute (it is valid for an Apple event to have no
// target address attribute).
//
// returnID: The return ID to be set in the returned descriptor. If you pass a value of
// `kAutoGenerateReturnID`, the Apple Event Manager assigns the created Apple
// event a return ID that is unique to the current session. If you pass any
// other value, the Apple Event Manager assigns that value for the ID.
//
// transactionID: The transaction ID to be set in the returned descriptor. A transaction is a
// sequence of Apple events that are sent back and forth between client and
// server applications, beginning with the client’s initial request for a
// service. All Apple events that are part of a transaction must have the same
// transaction ID. You can specify `kAnyTransactionID` if the Apple event is
// not one of a series of interdependent Apple events.
//
// # Return Value
// 
// The initialized Apple event (an instance of [NSAppleEventDescriptor]), or
// `nil` if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(eventClass:eventID:targetDescriptor:returnID:transactionID:)
func (a NSAppleEventDescriptor) InitWithEventClassEventIDTargetDescriptorReturnIDTransactionID(eventClass uint32, eventID uint32, targetDescriptor INSAppleEventDescriptor, returnID int16, transactionID int32) NSAppleEventDescriptor {
	rv := objc.Send[NSAppleEventDescriptor](a.ID, objc.Sel("initWithEventClass:eventID:targetDescriptor:returnID:transactionID:"), eventClass, eventID, targetDescriptor, returnID, transactionID)
	return rv
}

// Returns a descriptor obtained by coercing the receiver to the specified
// type.
//
// descriptorType: The descriptor type to coerce the receiver to.
//
// # Return Value
// 
// A descriptor of the specified type, or `nil` if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/coerce(toDescriptorType:)
func (a NSAppleEventDescriptor) CoerceToDescriptorType(descriptorType uint32) INSAppleEventDescriptor {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("coerceToDescriptorType:"), descriptorType)
	return NSAppleEventDescriptorFromID(rv)
}

// Returns the descriptor at the specified (one-based) position in the
// receiving descriptor list.
//
// index: The one-based descriptor list position of the descriptor to return.
//
// # Return Value
// 
// The descriptor from the specified position (one-based) in the descriptor
// list, or `nil` if the specified descriptor cannot be obtained.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/atIndex(_:)
func (a NSAppleEventDescriptor) DescriptorAtIndex(index int) INSAppleEventDescriptor {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("descriptorAtIndex:"), index)
	return NSAppleEventDescriptorFromID(rv)
}

// Inserts a descriptor at the specified (one-based) position in the receiving
// descriptor list, replacing the existing descriptor, if any, at that
// position.
//
// descriptor: The descriptor to insert in the receiver. Specifying an index of 0 or count
// + 1 causes appending to the end of the list.
//
// index: The one-based descriptor list position at which to insert the descriptor.
//
// # Discussion
// 
// Because it actually replaces the descriptor, if any, at the specified
// position, this method might better be called ``. The receiver must be a
// list descriptor. The indices are one-based. Currently provides no
// indication if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/insert(_:at:)
func (a NSAppleEventDescriptor) InsertDescriptorAtIndex(descriptor INSAppleEventDescriptor, index int) {
	objc.Send[objc.ID](a.ID, objc.Sel("insertDescriptor:atIndex:"), descriptor, index)
}

// Removes the descriptor at the specified (one-based) position in the
// receiving descriptor list.
//
// index: The one-based position of the descriptor to remove.
//
// # Discussion
// 
// The receiver must be a list descriptor. The indices are one-based.
// Currently provides no indication if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/remove(at:)
func (a NSAppleEventDescriptor) RemoveDescriptorAtIndex(index int) {
	objc.Send[objc.ID](a.ID, objc.Sel("removeDescriptorAtIndex:"), index)
}

// Returns the receiver’s descriptor for the specified keyword.
//
// keyword: A keyword (a four-character code) that identifies the descriptor to obtain.
//
// # Return Value
// 
// A descriptor for the specified keyword, or `nil` if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/forKeyword(_:)
func (a NSAppleEventDescriptor) DescriptorForKeyword(keyword uint32) INSAppleEventDescriptor {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("descriptorForKeyword:"), keyword)
	return NSAppleEventDescriptorFromID(rv)
}

// Returns the keyword for the descriptor at the specified (one-based)
// position in the receiver.
//
// index: The one-based descriptor list position of the descriptor to get the keyword
// for.
//
// # Return Value
// 
// The keyword (a four-character code) for the descriptor at the one-based
// location specified by `anIndex`, or 0 if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/keywordForDescriptor(at:)
func (a NSAppleEventDescriptor) KeywordForDescriptorAtIndex(index int) uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("keywordForDescriptorAtIndex:"), index)
	return rv
}

// Removes the receiver’s descriptor identified by the specified keyword.
//
// keyword: A keyword (a four-character code) that identifies the descriptor to remove.
//
// # Discussion
// 
// The receiver must be an Apple event or Apple event record. Currently
// provides no indication if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/remove(withKeyword:)
func (a NSAppleEventDescriptor) RemoveDescriptorWithKeyword(keyword uint32) {
	objc.Send[objc.ID](a.ID, objc.Sel("removeDescriptorWithKeyword:"), keyword)
}

// Adds a descriptor, identified by a keyword, to the receiver.
//
// descriptor: The descriptor to add to the receiver.
//
// keyword: A keyword (a four-character code) that identifies the descriptor to add. If
// a descriptor with that keyword already exists in the receiver, it is
// replaced.
//
// # Discussion
// 
// The receiver must be an Apple event or Apple event record. Currently
// provides no indication if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/setDescriptor(_:forKeyword:)
func (a NSAppleEventDescriptor) SetDescriptorForKeyword(descriptor INSAppleEventDescriptor, keyword uint32) {
	objc.Send[objc.ID](a.ID, objc.Sel("setDescriptor:forKeyword:"), descriptor, keyword)
}

// Returns a descriptor for the receiver’s Apple event attribute identified
// by the specified keyword.
//
// keyword: A keyword (a four-character code) that identifies the descriptor to obtain.
//
// # Return Value
// 
// The attribute descriptor for the specified keyword, or `nil` if an error
// occurs.
//
// # Discussion
// 
// The receiver must be an Apple event.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/attributeDescriptor(forKeyword:)
func (a NSAppleEventDescriptor) AttributeDescriptorForKeyword(keyword uint32) INSAppleEventDescriptor {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("attributeDescriptorForKeyword:"), keyword)
	return NSAppleEventDescriptorFromID(rv)
}

// Returns a descriptor for the receiver’s Apple event parameter identified
// by the specified keyword.
//
// keyword: A keyword (a four-character code) that identifies the parameter descriptor
// to obtain.
//
// # Return Value
// 
// A descriptor for the specified keyword, or `nil` if an error occurs.
//
// # Discussion
// 
// The receiver must be an Apple event.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/paramDescriptor(forKeyword:)
func (a NSAppleEventDescriptor) ParamDescriptorForKeyword(keyword uint32) INSAppleEventDescriptor {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("paramDescriptorForKeyword:"), keyword)
	return NSAppleEventDescriptorFromID(rv)
}

// Removes the receiver’s parameter descriptor identified by the specified
// keyword.
//
// keyword: A keyword (a four-character code) that identifies the parameter descriptor
// to remove. Currently provides no indication if an error occurs.
//
// # Discussion
// 
// The receiver must be an Apple event or Apple event record, both of which
// can contain parameters.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/removeParamDescriptor(withKeyword:)
func (a NSAppleEventDescriptor) RemoveParamDescriptorWithKeyword(keyword uint32) {
	objc.Send[objc.ID](a.ID, objc.Sel("removeParamDescriptorWithKeyword:"), keyword)
}

// Adds a descriptor to the receiver as an attribute identified by the
// specified keyword.
//
// descriptor: The attribute descriptor to add to the receiver.
//
// keyword: A keyword (a four-character code) that identifies the attribute descriptor
// to add. If a descriptor with that keyword already exists in the receiver,
// it is replaced.
//
// # Discussion
// 
// The receiver must be an Apple event. Currently provides no indication if an
// error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/setAttribute(_:forKeyword:)
func (a NSAppleEventDescriptor) SetAttributeDescriptorForKeyword(descriptor INSAppleEventDescriptor, keyword uint32) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAttributeDescriptor:forKeyword:"), descriptor, keyword)
}

// Adds a descriptor to the receiver as an Apple event parameter identified by
// the specified keyword.
//
// descriptor: The parameter descriptor to add to the receiver.
//
// keyword: A keyword (a four-character code) that identifies the parameter descriptor
// to add. If a descriptor with that keyword already exists in the receiver,
// it is replaced.
//
// # Discussion
// 
// The receiver must be an Apple event or Apple event record, both of which
// can contain parameters.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/setParam(_:forKeyword:)
func (a NSAppleEventDescriptor) SetParamDescriptorForKeyword(descriptor INSAppleEventDescriptor, keyword uint32) {
	objc.Send[objc.ID](a.ID, objc.Sel("setParamDescriptor:forKeyword:"), descriptor, keyword)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/sendEvent(options:timeout:)
func (a NSAppleEventDescriptor) SendEventWithOptionsTimeoutError(sendOptions NSAppleEventSendOptions, timeoutInSeconds float64) (INSAppleEventDescriptor, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("sendEventWithOptions:timeout:error:"), sendOptions, timeoutInSeconds, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSAppleEventDescriptor{}, NSErrorFrom(errorPtr)
	}
	return NSAppleEventDescriptorFromID(rv), nil

}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (a NSAppleEventDescriptor) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (a NSAppleEventDescriptor) InitWithCoder(coder INSCoder) NSAppleEventDescriptor {
	rv := objc.Send[NSAppleEventDescriptor](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}





// Creates a descriptor that represents an Apple event, initialized according
// to the specified information.
//
// eventClass: The event class to be set in the returned descriptor.
//
// eventID: The event ID to be set in the returned descriptor.
//
// targetDescriptor: A pointer to a descriptor that identifies the target application for the
// Apple event. Passing `nil` results in an Apple event descriptor that has no
// `keyAddressAttr` attribute (it is valid for an Apple event to have no
// target address attribute).
//
// returnID: The return ID to be set in the returned descriptor. If you pass a value of
// `kAutoGenerateReturnID`, the Apple Event Manager assigns the created Apple
// event a return ID that is unique to the current session. If you pass any
// other value, the Apple Event Manager assigns that value for the ID.
//
// transactionID: The transaction ID to be set in the returned descriptor. A transaction is a
// sequence of Apple events that are sent back and forth between client and
// server applications, beginning with the client’s initial request for a
// service. All Apple events that are part of a transaction must have the same
// transaction ID. You can specify `kAnyTransactionID` if the Apple event is
// not one of a series of interdependent Apple events.
//
// # Return Value
// 
// A descriptor for an Apple event, initialized according to the specified
// parameter values, or `nil` if an error occurs.
//
// # Discussion
// 
// Constants such as `kAutoGenerateReturnID` and `kAnyTransactionID` are
// defined in `AE.Framework()`, a subframework of
// `ApplicationServices.Framework()`.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/appleEvent(withEventClass:eventID:targetDescriptor:returnID:transactionID:)
func (_NSAppleEventDescriptorClass NSAppleEventDescriptorClass) AppleEventWithEventClassEventIDTargetDescriptorReturnIDTransactionID(eventClass uint32, eventID uint32, targetDescriptor INSAppleEventDescriptor, returnID int16, transactionID int32) NSAppleEventDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_NSAppleEventDescriptorClass.class), objc.Sel("appleEventWithEventClass:eventID:targetDescriptor:returnID:transactionID:"), eventClass, eventID, targetDescriptor, returnID, transactionID)
	return NSAppleEventDescriptorFromID(rv)
}

// Creates and initializes an empty list descriptor.
//
// # Return Value
// 
// An empty list descriptor, or `nil` if an error occurs.
//
// # Discussion
// 
// A list descriptor is a descriptor whose data consists of one or more
// descriptors. You can add items to the list by calling
// [InsertDescriptorAtIndex] or remove them with [RemoveDescriptorAtIndex].
// 
// Invoking this method is equivalent to allocating an instance of
// [NSAppleEventDescriptor] and invoking [InitListDescriptor].
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/list()
func (_NSAppleEventDescriptorClass NSAppleEventDescriptorClass) ListDescriptor() NSAppleEventDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_NSAppleEventDescriptorClass.class), objc.Sel("listDescriptor"))
	return NSAppleEventDescriptorFromID(rv)
}

// Creates and initializes a descriptor with no parameter or attribute values
// set.
//
// # Return Value
// 
// A descriptor with no parameter or attribute values set, or `nil` if an
// error occurs.
//
// # Discussion
// 
// You don’t typically call this method, as most [NSAppleEventDescriptor]
// instance methods can’t be safely called on the returned empty descriptor.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/null()
func (_NSAppleEventDescriptorClass NSAppleEventDescriptorClass) NullDescriptor() NSAppleEventDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_NSAppleEventDescriptorClass.class), objc.Sel("nullDescriptor"))
	return NSAppleEventDescriptorFromID(rv)
}

// Creates and initializes a descriptor for an Apple event record whose data
// has yet to be set.
//
// # Return Value
// 
// An Apple event descriptor whose data has yet to be set, or `nil` if an
// error occurs.
//
// # Discussion
// 
// An Apple event record is a descriptor whose data is a set of descriptors
// keyed by four-character codes. You can add information to the descriptor
// with methods such as [SetAttributeDescriptorForKeyword],
// [SetDescriptorForKeyword], and [SetParamDescriptorForKeyword].
// 
// Invoking this method is equivalent to allocating an instance of
// [NSAppleEventDescriptor] and invoking [InitRecordDescriptor].
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/record()
func (_NSAppleEventDescriptorClass NSAppleEventDescriptorClass) RecordDescriptor() NSAppleEventDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_NSAppleEventDescriptorClass.class), objc.Sel("recordDescriptor"))
	return NSAppleEventDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/currentProcess()
func (_NSAppleEventDescriptorClass NSAppleEventDescriptorClass) CurrentProcessDescriptor() NSAppleEventDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_NSAppleEventDescriptorClass.class), objc.Sel("currentProcessDescriptor"))
	return NSAppleEventDescriptorFromID(rv)
}

// Creates a descriptor initialized with the specified event type that stores
// the specified data (from a series of bytes).
//
// descriptorType: The descriptor type to be set in the returned descriptor.
//
// bytes: The data, as a sequence of bytes, to be set in the returned descriptor.
//
// byteCount: The length, in bytes, of the data to be set in the returned descriptor.
//
// # Return Value
// 
// A descriptor with the specified type and data, or `nil` if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/descriptorWithDescriptorType:bytes:length:
func (_NSAppleEventDescriptorClass NSAppleEventDescriptorClass) DescriptorWithDescriptorTypeBytesLength(descriptorType uint32, bytes unsafe.Pointer, byteCount uint) NSAppleEventDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_NSAppleEventDescriptorClass.class), objc.Sel("descriptorWithDescriptorType:bytes:length:"), descriptorType, bytes, byteCount)
	return NSAppleEventDescriptorFromID(rv)
}

// Creates a descriptor initialized with the specified event type that stores
// the specified data (from an instance of [NSData]).
//
// descriptorType: The descriptor type to be set in the returned descriptor.
//
// data: The data, as an instance of [NSData], to be set in the returned descriptor.
//
// # Return Value
// 
// A descriptor with the specified type and data, or `nil` if an error occurs.
//
// # Discussion
// 
// You can use this method to create a descriptor that you can build into a
// complete Apple event by calling methods such as
// [SetAttributeDescriptorForKeyword], [SetDescriptorForKeyword], and
// [SetParamDescriptorForKeyword].
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/descriptorWithDescriptorType:data:
func (_NSAppleEventDescriptorClass NSAppleEventDescriptorClass) DescriptorWithDescriptorTypeData(descriptorType uint32, data INSData) NSAppleEventDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_NSAppleEventDescriptorClass.class), objc.Sel("descriptorWithDescriptorType:data:"), descriptorType, data)
	return NSAppleEventDescriptorFromID(rv)
}








// The [AEDesc] structure encapsulated by the receiver, if it has one.
//
// # Discussion
// 
// If the receiver has a valid [AEDesc] structure, returns a pointer to it;
// otherwise returns `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/aeDesc
func (a NSAppleEventDescriptor) AeDesc() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("aeDesc"))
	return objectivec.Object{ID: rv}
}



// The contents of the receiver as a Boolean value, coercing (to
// `typeBoolean`) if necessary.
//
// # Discussion
// 
// The contents of the descriptor, as a Boolean value, or `false` if an error
// occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/booleanValue
func (a NSAppleEventDescriptor) BooleanValue() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("booleanValue"))
	return rv
}



// The receiver’s data.
//
// # Discussion
// 
// An instance of [NSData] containing the receiver’s data, or `nil` if an
// error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/data
func (a NSAppleEventDescriptor) Data() INSData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("data"))
	return NSDataFromID(objc.ID(rv))
}



// The descriptor type of the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/descriptorType
func (a NSAppleEventDescriptor) DescriptorType() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("descriptorType"))
	return rv
}



// The contents of the receiver as an enumeration type, coercing to
// `typeEnumerated` if necessary.
//
// # Discussion
// 
// The contents of the descriptor, as an enumeration type, or 0 if an error
// occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/enumCodeValue
func (a NSAppleEventDescriptor) EnumCodeValue() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("enumCodeValue"))
	return rv
}



// The contents of the receiver as an integer, coercing (to `typeSInt32`) if
// necessary.
//
// # Discussion
// 
// The contents of the descriptor, as an integer value, or 0 if an error
// occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/int32Value
func (a NSAppleEventDescriptor) Int32Value() int32 {
	rv := objc.Send[int32](a.ID, objc.Sel("int32Value"))
	return rv
}



// The number of descriptors in the receiver’s descriptor list.
//
// # Discussion
// 
// The number of descriptors in the receiver’s descriptor list (possibly 0);
// returns 0 if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/numberOfItems
func (a NSAppleEventDescriptor) NumberOfItems() int {
	rv := objc.Send[int](a.ID, objc.Sel("numberOfItems"))
	return rv
}



// The contents of the receiver as a Unicode text string, coercing to
// `typeUnicodeText` if necessary.
//
// # Discussion
// 
// The contents of the descriptor, as a string, or `nil` if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/stringValue
func (a NSAppleEventDescriptor) StringValue() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("stringValue"))
	return NSStringFromID(rv).String()
}



// The contents of the receiver as a type, coercing to `typeType` if
// necessary.
//
// # Discussion
// 
// The contents of the descriptor, as a type, or 0 if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/typeCodeValue
func (a NSAppleEventDescriptor) TypeCodeValue() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("typeCodeValue"))
	return rv
}



// The event class for the receiver.
//
// # Discussion
// 
// The event class (a four-character code) for the receiver, or 0 if an error
// occurs.
// 
// The receiver must be an Apple event. An Apple event is identified by its
// event class and event ID, a pair of four-character codes stored as 32-bit
// integers. For example, most events in the Standard suite have the
// four-character code `'core'` (defined as the constant `kAECoreSuite` in
// `AE.Framework()`, a subframework of `ApplicationServices.Framework()`). For
// more information on event classes and event IDs, see Building an Apple
// Event in Apple Events Programming Guide.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/eventClass
func (a NSAppleEventDescriptor) EventClass() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("eventClass"))
	return rv
}



// The event ID for the receiver.
//
// # Discussion
// 
// The event ID (a four-character code) for the receiver, or 0 if an error
// occurs.
// 
// The receiver must be an Apple event. An Apple event is identified by its
// event class and event ID, a pair of four-character codes stored as 32-bit
// integers. For example, the `open` Apple event from the Standard suite has
// the four-character code `'odoc'` (defined as the constant `kAEOpen` in
// `AE.Framework()`, a subframework of `ApplicationServices.Framework()`).
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/eventID
func (a NSAppleEventDescriptor) EventID() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("eventID"))
	return rv
}



// The receiver’s return ID (the ID for a reply Apple event).
//
// # Discussion
// 
// The receiver’s return ID (an integer value), or 0 if an error occurs.
// 
// The receiver must be an Apple event.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/returnID
func (a NSAppleEventDescriptor) ReturnID() int16 {
	rv := objc.Send[int16](a.ID, objc.Sel("returnID"))
	return rv
}



// The receiver’s transaction ID, if any.
//
// # Discussion
// 
// The receiver’s transaction ID (an integer value), or 0 if an error
// occurs.
// 
// The receiver must be an Apple event. Currently provides no indication if an
// error occurs. For more information on transactions, see the description for
// [AppleEventWithEventClassEventIDTargetDescriptorReturnIDTransactionID].
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/transactionID
func (a NSAppleEventDescriptor) TransactionID() int32 {
	rv := objc.Send[int32](a.ID, objc.Sel("transactionID"))
	return rv
}



// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/dateValue
func (a NSAppleEventDescriptor) DateValue() INSDate {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("dateValue"))
	return NSDateFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/doubleValue
func (a NSAppleEventDescriptor) DoubleValue() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("doubleValue"))
	return rv
}



// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/fileURLValue
func (a NSAppleEventDescriptor) FileURLValue() INSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("fileURLValue"))
	return NSURLFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/isRecordDescriptor
func (a NSAppleEventDescriptor) IsRecordDescriptor() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isRecordDescriptor"))
	return rv
}














			// Protocol methods for NSCopying
			















