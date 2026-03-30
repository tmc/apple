// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSValue] class.
var (
	_NSValueClass     NSValueClass
	_NSValueClassOnce sync.Once
)

func getNSValueClass() NSValueClass {
	_NSValueClassOnce.Do(func() {
		_NSValueClass = NSValueClass{class: objc.GetClass("NSValue")}
	})
	return _NSValueClass
}

// GetNSValueClass returns the class object for NSValue.
func GetNSValueClass() NSValueClass {
	return getNSValueClass()
}

type NSValueClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSValueClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSValueClass) Alloc() NSValue {
	rv := objc.Send[NSValue](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A simple container for a single C or Objective-C data item.
//
// # Overview
//
// An [NSValue] object can hold any of the scalar types such as `int`,
// `float`, and `char`, as well as pointers, structures, and object `id`
// references. Use this class to work with such data types in collections
// (such as [NSArray] and [NSSet]), [Key-value coding], and other APIs that
// require Objective-C objects. [NSValue] objects are always immutable.
//
// # Subclassing Notes
//
// The abstract [NSValue] class is the public interface of a class cluster
// consisting mostly of private, concrete classes that create and return a
// value object appropriate for a given situation. It is possible to subclass
// [NSValue], but doing so requires providing storage facilities for the value
// (which is not inherited by subclasses) and implementing two primitive
// methods.
//
// # Methods to Override
//
// Any subclass of [NSValue] override the primitive instance methods
// [NSValue.GetValue] and [NSValue.ObjCType]. These methods must operate on the storage that
// you provide for the value.
//
// You might want to implement an initializer for your subclass that is suited
// to the storage you provide. The [NSValue] class does not have a designated
// initializer, so your initializer need only invoke the [init()] method of
// `super`. The [NSValue] class adopts the [NSCopying] and [NSSecureCoding]
// protocols; if you want instances of your own custom subclass created from
// copying or coding, override the methods in these protocols.
//
// You may also wish to implement the [NSValue.Hash] method to make your subclass work
// well in collections.
//
// # Alternatives to Subclassing
//
// If you need only to use [NSValue] objects for wrap a custom data types or
// structures defined by your app, you need not create an [NSValue] subclass.
// Instead, create a category that uses existing [NSValue] methods to store
// and retrieve data of your custom type. For example, the code below defines
// a custom Polyhedron structure and creates [NSValue] convenience methods to
// store and retrieve it:
//
// # Working with Raw Values
//
//   - [NSValue.InitWithBytesObjCType]: Initializes a value object to contain the specified value, interpreted with the specified Objective-C type.
//   - [NSValue.ObjCType]: A C string containing the Objective-C type of the data contained in the value object.
//
// # Working with Pointer and Object Values
//
//   - [NSValue.PointerValue]: Returns the value as an untyped pointer.
//   - [NSValue.NonretainedObjectValue]: The value as a non-retained pointer to an object.
//
// # Working with Range Values
//
//   - [NSValue.RangeValue]: The Foundation range structure representation of the value.
//
// # Working with Foundation Geometry Values
//
//   - [NSValue.PointValue]: The Foundation point structure representation of the value.
//   - [NSValue.SizeValue]: The Foundation size structure representation of the value.
//   - [NSValue.RectValue]: The Foundation rectangle structure representation of the value.
//
// # Working with CoreAnimation Transform Values
//
//   - [NSValue.CATransform3DValue]: The CoreAnimation transform structure representation of the value.
//
// # Working with Media Time Values
//
//   - [NSValue.CMTimeValue]: The CoreMedia time structure representation of the value.
//   - [NSValue.CMTimeRangeValue]: The CoreMedia time range structure representation of the value.
//   - [NSValue.CMTimeMappingValue]: The CoreMedia time mapping structure representation of the value.
//
// # Working with Geographic Coordinate Values
//
//   - [NSValue.MKCoordinateValue]: The CoreLocation geographic coordinate structure representation of the value.
//   - [NSValue.MKCoordinateSpanValue]: The MapKit coordinate span structure representation of the value.
//
// # Working with SceneKit Vector and Matrix Values
//
//   - [NSValue.SCNVector3Value]: The three-element Scene Kit vector representation of the value.
//   - [NSValue.SCNVector4Value]: The four-element Scene Kit vector representation of the value.
//   - [NSValue.SCNMatrix4Value]: The Scene Kit 4 x 4 matrix representation of the value.
//
// # Comparing Value Objects
//
//   - [NSValue.IsEqualToValue]: Returns a Boolean value that indicates whether the value object and another value object are equal.
//
// # Instance Properties
//
//   - [NSValue.EdgeInsetsValue]
//   - [NSValue.GCPoint2Value]
//   - [NSValue.CMVideoDimensionsValue]
//
// # Instance Methods
//
//   - [NSValue.GetValueSize]
//
// See: https://developer.apple.com/documentation/Foundation/NSValue
//
// [Key-value coding]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/KeyValueCoding.html#//apple_ref/doc/uid/TP40008195-CH25
// [init()]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/init()
type NSValue struct {
	objectivec.Object
}

// NSValueFromID constructs a [NSValue] from an objc.ID.
//
// A simple container for a single C or Objective-C data item.
func NSValueFromID(id objc.ID) NSValue {
	return NSValue{objectivec.Object{ID: id}}
}

// NOTE: NSValue adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSValue] class.
//
// # Working with Raw Values
//
//   - [INSValue.InitWithBytesObjCType]: Initializes a value object to contain the specified value, interpreted with the specified Objective-C type.
//   - [INSValue.ObjCType]: A C string containing the Objective-C type of the data contained in the value object.
//
// # Working with Pointer and Object Values
//
//   - [INSValue.PointerValue]: Returns the value as an untyped pointer.
//   - [INSValue.NonretainedObjectValue]: The value as a non-retained pointer to an object.
//
// # Working with Range Values
//
//   - [INSValue.RangeValue]: The Foundation range structure representation of the value.
//
// # Working with Foundation Geometry Values
//
//   - [INSValue.PointValue]: The Foundation point structure representation of the value.
//   - [INSValue.SizeValue]: The Foundation size structure representation of the value.
//   - [INSValue.RectValue]: The Foundation rectangle structure representation of the value.
//
// # Working with CoreAnimation Transform Values
//
//   - [INSValue.CATransform3DValue]: The CoreAnimation transform structure representation of the value.
//
// # Working with Media Time Values
//
//   - [INSValue.CMTimeValue]: The CoreMedia time structure representation of the value.
//   - [INSValue.CMTimeRangeValue]: The CoreMedia time range structure representation of the value.
//   - [INSValue.CMTimeMappingValue]: The CoreMedia time mapping structure representation of the value.
//
// # Working with Geographic Coordinate Values
//
//   - [INSValue.MKCoordinateValue]: The CoreLocation geographic coordinate structure representation of the value.
//   - [INSValue.MKCoordinateSpanValue]: The MapKit coordinate span structure representation of the value.
//
// # Working with SceneKit Vector and Matrix Values
//
//   - [INSValue.SCNVector3Value]: The three-element Scene Kit vector representation of the value.
//   - [INSValue.SCNVector4Value]: The four-element Scene Kit vector representation of the value.
//   - [INSValue.SCNMatrix4Value]: The Scene Kit 4 x 4 matrix representation of the value.
//
// # Comparing Value Objects
//
//   - [INSValue.IsEqualToValue]: Returns a Boolean value that indicates whether the value object and another value object are equal.
//
// # Instance Properties
//
//   - [INSValue.EdgeInsetsValue]
//   - [INSValue.GCPoint2Value]
//   - [INSValue.CMVideoDimensionsValue]
//
// # Instance Methods
//
//   - [INSValue.GetValueSize]
//
// See: https://developer.apple.com/documentation/Foundation/NSValue
type INSValue interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Working with Raw Values

	// Initializes a value object to contain the specified value, interpreted with the specified Objective-C type.
	InitWithBytesObjCType(value unsafe.Pointer, type_ string) NSValue
	// A C string containing the Objective-C type of the data contained in the value object.
	ObjCType() string

	// Topic: Working with Pointer and Object Values

	// Returns the value as an untyped pointer.
	PointerValue() unsafe.Pointer
	// The value as a non-retained pointer to an object.
	NonretainedObjectValue() objectivec.IObject

	// Topic: Working with Range Values

	// The Foundation range structure representation of the value.
	RangeValue() NSRange

	// Topic: Working with Foundation Geometry Values

	// The Foundation point structure representation of the value.
	PointValue() NSPoint
	// The Foundation size structure representation of the value.
	SizeValue() NSSize
	// The Foundation rectangle structure representation of the value.
	RectValue() NSRect

	// Topic: Working with CoreAnimation Transform Values

	// The CoreAnimation transform structure representation of the value.
	CATransform3DValue() objectivec.IObject

	// Topic: Working with Media Time Values

	// The CoreMedia time structure representation of the value.
	CMTimeValue() objectivec.IObject
	// The CoreMedia time range structure representation of the value.
	CMTimeRangeValue() objectivec.IObject
	// The CoreMedia time mapping structure representation of the value.
	CMTimeMappingValue() objectivec.IObject

	// Topic: Working with Geographic Coordinate Values

	// The CoreLocation geographic coordinate structure representation of the value.
	MKCoordinateValue() objectivec.IObject
	// The MapKit coordinate span structure representation of the value.
	MKCoordinateSpanValue() objectivec.IObject

	// Topic: Working with SceneKit Vector and Matrix Values

	// The three-element Scene Kit vector representation of the value.
	SCNVector3Value() objectivec.IObject
	// The four-element Scene Kit vector representation of the value.
	SCNVector4Value() objectivec.IObject
	// The Scene Kit 4 x 4 matrix representation of the value.
	SCNMatrix4Value() objectivec.IObject

	// Topic: Comparing Value Objects

	// Returns a Boolean value that indicates whether the value object and another value object are equal.
	IsEqualToValue(value INSValue) bool

	// Topic: Instance Properties

	EdgeInsetsValue() NSEdgeInsets
	GCPoint2Value() objectivec.IObject
	CMVideoDimensionsValue() objectivec.IObject

	// Topic: Instance Methods

	GetValueSize(value unsafe.Pointer, size uint)

	// Returns an integer that can be used as a table address in a hash table structure.
	Hash() int
	SetHash(value int)
}

// Init initializes the instance.
func (v NSValue) Init() NSValue {
	rv := objc.Send[NSValue](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v NSValue) Autorelease() NSValue {
	rv := objc.Send[NSValue](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSValue creates a new NSValue instance.
func NewNSValue() NSValue {
	class := getNSValueClass()
	rv := objc.Send[NSValue](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a value object to contain the specified value, interpreted with
// the specified Objective-C type.
//
// value: A pointer to data to be stored in the new value object.
//
// type: The Objective-C type of `value`, as provided by the `@encode()` compiler
// directive. Do not hard-code this parameter as a C string.
//
// # Return Value
//
// An initialized value object that contains `value`, which is interpreted as
// being of the Objective-C type `type`. The returned object might be
// different than the original receiver.
//
// # Discussion
//
// See [Number and Value Programming Topics] for other considerations in
// creating a value object.
//
// This is the designated initializer for the [NSValue] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(bytes:objCType:)
//
// [Number and Value Programming Topics]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/NumbersandValues/NumbersandValues.html#//apple_ref/doc/uid/10000038i
func NewValueWithBytesObjCType(value unsafe.Pointer, type_ string) NSValue {
	instance := getNSValueClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytes:objCType:"), value, unsafe.Pointer(unsafe.StringData(type_+"\x00")))
	return NSValueFromID(rv)
}

// Creates a new value object containing the specified CoreAnimation transform
// structure.
//
// t: The value for the new object.
//
// # Return Value
//
// A new value object that contains the transform information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(CATransform3D:)
// t is a [quartzcore.CATransform3D].
func NewValueWithCATransform3D(t objectivec.IObject) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithCATransform3D:"), t)
	return NSValueFromID(rv)
}

// Creates a new value object containing the specified CoreGraphics affine
// transform structure.
//
// transform: The value for the new object.
//
// # Return Value
//
// A new value object that contains the affine transform information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(CGAffineTransform:)
func NewValueWithCGAffineTransform(transform corefoundation.CGAffineTransform) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithCGAffineTransform:"), transform)
	return NSValueFromID(rv)
}

// Creates a new value object containing the specified CoreGraphics point
// structure.
//
// point: The value for the new object.
//
// # Return Value
//
// A new value object that contains the point information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(CGPoint:)
func NewValueWithCGPoint(point corefoundation.CGPoint) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithCGPoint:"), point)
	return NSValueFromID(rv)
}

// Creates a new value object containing the specified CoreGraphics rectangle
// structure.
//
// rect: The value for the new object.
//
// # Return Value
//
// A new value object that contains the rectangle information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(CGRect:)
func NewValueWithCGRect(rect corefoundation.CGRect) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithCGRect:"), rect)
	return NSValueFromID(rv)
}

// Creates a new value object containing the specified CoreGraphics size
// structure.
//
// size: The value for the new object.
//
// # Return Value
//
// A new value object that contains the size information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(CGSize:)
func NewValueWithCGSize(size corefoundation.CGSize) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithCGSize:"), size)
	return NSValueFromID(rv)
}

// Creates a new value object containing the specified CoreGraphics vector
// structure.
//
// vector: The value for the new object.
//
// # Return Value
//
// A new value object that contains the vector information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(CGVector:)
func NewValueWithCGVector(vector corefoundation.CGVector) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithCGVector:"), vector)
	return NSValueFromID(rv)
}

// Creates a new value object containing the specified CoreMedia time
// structure.
//
// time: The value for the new object.
//
// # Return Value
//
// A new value object that contains the media time information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(CMTime:)
// time is a [coremedia.CMTime].
func NewValueWithCMTime(time objectivec.IObject) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithCMTime:"), time)
	return NSValueFromID(rv)
}

// Creates a new value object containing the specified CoreMedia time mapping
// structure.
//
// timeMapping: The value for the new object.
//
// # Return Value
//
// A new value object that contains the time mapping information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(CMTimeMapping:)
// timeMapping is a [coremedia.CMTimeMapping].
func NewValueWithCMTimeMapping(timeMapping objectivec.IObject) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithCMTimeMapping:"), timeMapping)
	return NSValueFromID(rv)
}

// Creates a new value object containing the specified CoreMedia time range
// structure.
//
// timeRange: The value for the new object.
//
// # Return Value
//
// A new value object that contains the time range information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(CMTimeRange:)
// timeRange is a [coremedia.CMTimeRange].
func NewValueWithCMTimeRange(timeRange objectivec.IObject) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithCMTimeRange:"), timeRange)
	return NSValueFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSValue/init(CMVideoDimensions:)
// dimensions is a [coremedia.CMVideoDimensions].
func NewValueWithCMVideoDimensions(dimensions objectivec.IObject) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithCMVideoDimensions:"), dimensions)
	return NSValueFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSValue/init(coder:)
func NewValueWithCoder(coder INSCoder) NSValue {
	instance := getNSValueClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSValueFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSValue/init(directionalEdgeInsets:)
// insets is a [appkit.NSDirectionalEdgeInsets].
func NewValueWithDirectionalEdgeInsets(insets objectivec.IObject) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithDirectionalEdgeInsets:"), insets)
	return NSValueFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSValue/init(edgeInsets:)
func NewValueWithEdgeInsets(insets NSEdgeInsets) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithEdgeInsets:"), insets)
	return NSValueFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSValue/init(GCPoint2:)
// point is a [gamecontroller.GCPoint2].
func NewValueWithGCPoint2(point objectivec.IObject) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithGCPoint2:"), point)
	return NSValueFromID(rv)
}

// Creates a new value object containing the specified CoreLocation geographic
// coordinate structure.
//
// coordinate: The value for the new object.
//
// # Return Value
//
// A new value object that contains the geographic coordinate information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(MKCoordinate:)
// coordinate is a [corelocation.CLLocationCoordinate2D].
func NewValueWithMKCoordinate(coordinate objectivec.IObject) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithMKCoordinate:"), coordinate)
	return NSValueFromID(rv)
}

// Creates a new value object containing the specified MapKit coordinate span
// structure.
//
// span: The value for the new object.
//
// # Return Value
//
// A new value object that contains the coordinate span information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(MKCoordinateSpan:)
// span is a [mapkit.MKCoordinateSpan].
func NewValueWithMKCoordinateSpan(span objectivec.IObject) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithMKCoordinateSpan:"), span)
	return NSValueFromID(rv)
}

// Creates a value object containing the specified object.
//
// anObject: The value for the new object.
//
// # Return Value
//
// A new value object that contains `anObject`.
//
// # Discussion
//
// This method is equivalent to invoking [ValueWithObjCType] in this manner:
//
// This method is useful if you want to add an object to a [Collection] but
// don’t want the collection to create a strong reference to it.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(nonretainedObject:)
//
// [Collection]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/Collection.html#//apple_ref/doc/uid/TP40008195-CH10
func NewValueWithNonretainedObject(anObject objectivec.IObject) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithNonretainedObject:"), anObject)
	return NSValueFromID(rv)
}

// Creates a value object containing the specified value, interpreted with the
// specified Objective-C type.
//
// value: A pointer to data to be stored in the new value object.
//
// type: The Objective-C type of `value`, as provided by the `@encode()` compiler
// directive. Do not hard-code this parameter as a C string.
//
// # Return Value
//
// A new value object that contains `value`, which is interpreted as being of
// the Objective-C type `type`.
//
// # Discussion
//
// This method has the same effect as [ValueWithBytesObjCType] and may be
// deprecated in a future release. You should use [ValueWithBytesObjCType]
// instead.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(_:withObjCType:)
func NewValueWithObjCType(value unsafe.Pointer, type_ string) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("value:withObjCType:"), value, unsafe.Pointer(unsafe.StringData(type_+"\x00")))
	return NSValueFromID(rv)
}

// Creates a new value object containing the specified Foundation point
// structure.
//
// point: The value for the new object.
//
// # Return Value
//
// A new value object that contains the point information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(point:)
func NewValueWithPoint(point corefoundation.CGPoint) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithPoint:"), point)
	return NSValueFromID(rv)
}

// Creates a value object containing the specified pointer.
//
// pointer: The value for the new object.
//
// # Return Value
//
// A new value object that contains `aPointer`.
//
// # Discussion
//
// This method is equivalent to invoking [ValueWithObjCType] in this manner:
//
// This method does not copy the contents of `aPointer`, so you must not to
// free the memory at the pointer destination while the [NSValue] object
// exists. [NSData] objects may be more suited for arbitrary pointers than
// [NSValue] objects.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(pointer:)
func NewValueWithPointer(pointer unsafe.Pointer) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithPointer:"), pointer)
	return NSValueFromID(rv)
}

// Creates a new value object containing the specified Foundation range
// structure.
//
// range: The value for the new object.
//
// # Return Value
//
// A new value object that contains the range information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(range:)
func NewValueWithRange(range_ NSRange) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithRange:"), range_)
	return NSValueFromID(rv)
}

// Creates a new value object containing the specified Foundation rectangle
// structure.
//
// rect: The value for the new object.
//
// # Return Value
//
// A new value object that contains the data in the `rect` structure.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(rect:)
func NewValueWithRect(rect corefoundation.CGRect) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithRect:"), rect)
	return NSValueFromID(rv)
}

// Creates a value object that contains the specified SceneKit 4 x 4 matrix.
//
// v: The value for the new object.
//
// # Return Value
//
// A new value object that contains the matrix information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(SCNMatrix4:)
// v is a [scenekit.SCNMatrix4].
func NewValueWithSCNMatrix4(v objectivec.IObject) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithSCNMatrix4:"), v)
	return NSValueFromID(rv)
}

// Creates a value object that contains the specified three-element SceneKit
// vector.
//
// v: The value for the new object.
//
// # Return Value
//
// A new value object that contains the vector information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(SCNVector3:)
// v is a [scenekit.SCNVector3].
func NewValueWithSCNVector3(v objectivec.IObject) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithSCNVector3:"), v)
	return NSValueFromID(rv)
}

// Creates a value object that contains the specified four-element SceneKit
// vector.
//
// v: The value for the new object.
//
// # Return Value
//
// A new value object that contains the vector information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(SCNVector4:)
// v is a [scenekit.SCNVector4].
func NewValueWithSCNVector4(v objectivec.IObject) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithSCNVector4:"), v)
	return NSValueFromID(rv)
}

// Creates a new value object containing the specified Foundation size
// structure.
//
// size: The value for the new object.
//
// # Return Value
//
// A new value object that contains the size information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(size:)
func NewValueWithSize(size corefoundation.CGSize) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithSize:"), size)
	return NSValueFromID(rv)
}

// Creates a new value object containing the specified UIKit edge insets
// structure.
//
// insets: The value for the new object.
//
// # Return Value
//
// A new value object that contains the edge inset information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(UIEdgeInsets:)
// insets is a [uikit.UIEdgeInsets].
func NewValueWithUIEdgeInsets(insets objectivec.IObject) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithUIEdgeInsets:"), insets)
	return NSValueFromID(rv)
}

// Creates a new value object containing the specified UIKit offset structure.
//
// insets: The value for the new object.
//
// # Return Value
//
// A new value object that contains the offset information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(UIOffset:)
// insets is a [uikit.UIOffset].
func NewValueWithUIOffset(insets objectivec.IObject) NSValue {
	rv := objc.Send[objc.ID](objc.ID(getNSValueClass().class), objc.Sel("valueWithUIOffset:"), insets)
	return NSValueFromID(rv)
}

// Initializes a value object to contain the specified value, interpreted with
// the specified Objective-C type.
//
// value: A pointer to data to be stored in the new value object.
//
// type: The Objective-C type of `value`, as provided by the `@encode()` compiler
// directive. Do not hard-code this parameter as a C string.
//
// # Return Value
//
// An initialized value object that contains `value`, which is interpreted as
// being of the Objective-C type `type`. The returned object might be
// different than the original receiver.
//
// # Discussion
//
// See [Number and Value Programming Topics] for other considerations in
// creating a value object.
//
// This is the designated initializer for the [NSValue] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(bytes:objCType:)
//
// [Number and Value Programming Topics]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/NumbersandValues/NumbersandValues.html#//apple_ref/doc/uid/10000038i
func (v NSValue) InitWithBytesObjCType(value unsafe.Pointer, type_ string) NSValue {
	rv := objc.Send[NSValue](v.ID, objc.Sel("initWithBytes:objCType:"), value, unsafe.Pointer(unsafe.StringData(type_+"\x00")))
	return rv
}

// Returns a Boolean value that indicates whether the value object and another
// value object are equal.
//
// value: The other value object with which to compare the value object.
//
// # Return Value
//
// true if both value objects are equal; otherwise, false.
//
// # Discussion
//
// The [NSValue] class compares the type and contents of each value object to
// determine equality.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/isEqual(to:)
func (v NSValue) IsEqualToValue(value INSValue) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isEqualToValue:"), value)
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSValue/init(coder:)
func (v NSValue) InitWithCoder(coder INSCoder) NSValue {
	rv := objc.Send[NSValue](v.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSValue/getValue(_:size:)
func (v NSValue) GetValueSize(value unsafe.Pointer, size uint) {
	objc.Send[objc.ID](v.ID, objc.Sel("getValue:size:"), value, size)
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (v NSValue) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](v.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates a value object containing the specified value, interpreted with the
// specified Objective-C type.
//
// value: A pointer to data to be stored in the new value object.
//
// type: The Objective-C type of `value`, as provided by the `@encode()` compiler
// directive. Do not hard-code this parameter as a C string.
//
// # Return Value
//
// A new value object that contains `value`, which is interpreted as being of
// the Objective-C type `type`.
//
// # Discussion
//
// See [Number and Value Programming Topics] for other considerations in
// creating a value object and code examples.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/valueWithBytes:objCType:
//
// [Number and Value Programming Topics]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/NumbersandValues/NumbersandValues.html#//apple_ref/doc/uid/10000038i
func (_NSValueClass NSValueClass) ValueWithBytesObjCType(value unsafe.Pointer, type_ string) NSValue {
	rv := objc.Send[objc.ID](objc.ID(_NSValueClass.class), objc.Sel("valueWithBytes:objCType:"), value, unsafe.Pointer(unsafe.StringData(type_+"\x00")))
	return NSValueFromID(rv)
}

// A C string containing the Objective-C type of the data contained in the
// value object.
//
// # Discussion
//
// This property provides the same string produced by the `@encode()` compiler
// directive.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/objCType
func (v NSValue) ObjCType() string {
	rv := objc.Send[*byte](v.ID, objc.Sel("objCType"))
	return objc.GoString(rv)
}

// Returns the value as an untyped pointer.
//
// # Return Value
//
// The value as a pointer to void. If the value object was not created to hold
// a pointer-sized data item, the result is undefined.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/pointerValue
func (v NSValue) PointerValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v.ID, objc.Sel("pointerValue"))
	return rv
}

// The value as a non-retained pointer to an object.
//
// # Discussion
//
// If the value was not created to hold a pointer-sized data item, the result
// is undefined.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/nonretainedObjectValue
func (v NSValue) NonretainedObjectValue() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("nonretainedObjectValue"))
	return objectivec.Object{ID: rv}
}

// The Foundation range structure representation of the value.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/rangeValue
func (v NSValue) RangeValue() NSRange {
	rv := objc.Send[NSRange](v.ID, objc.Sel("rangeValue"))
	return NSRange(rv)
}

// The Foundation point structure representation of the value.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/pointValue
func (v NSValue) PointValue() NSPoint {
	rv := objc.Send[NSPoint](v.ID, objc.Sel("pointValue"))
	return NSPoint(rv)
}

// The Foundation size structure representation of the value.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/sizeValue
func (v NSValue) SizeValue() NSSize {
	rv := objc.Send[NSSize](v.ID, objc.Sel("sizeValue"))
	return NSSize(rv)
}

// The Foundation rectangle structure representation of the value.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/rectValue
func (v NSValue) RectValue() NSRect {
	rv := objc.Send[NSRect](v.ID, objc.Sel("rectValue"))
	return NSRect(rv)
}

// The CoreAnimation transform structure representation of the value.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/caTransform3DValue
func (v NSValue) CATransform3DValue() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("CATransform3DValue"))
	return objectivec.Object{ID: rv}
}

// The CoreMedia time structure representation of the value.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/timeValue
func (v NSValue) CMTimeValue() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("CMTimeValue"))
	return objectivec.Object{ID: rv}
}

// The CoreMedia time range structure representation of the value.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/timeRangeValue
func (v NSValue) CMTimeRangeValue() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("CMTimeRangeValue"))
	return objectivec.Object{ID: rv}
}

// The CoreMedia time mapping structure representation of the value.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/timeMappingValue
func (v NSValue) CMTimeMappingValue() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("CMTimeMappingValue"))
	return objectivec.Object{ID: rv}
}

// The CoreLocation geographic coordinate structure representation of the
// value.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/mkCoordinateValue
func (v NSValue) MKCoordinateValue() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("MKCoordinateValue"))
	return objectivec.Object{ID: rv}
}

// The MapKit coordinate span structure representation of the value.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/mkCoordinateSpanValue
func (v NSValue) MKCoordinateSpanValue() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("MKCoordinateSpanValue"))
	return objectivec.Object{ID: rv}
}

// The three-element Scene Kit vector representation of the value.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/scnVector3Value
func (v NSValue) SCNVector3Value() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("SCNVector3Value"))
	return objectivec.Object{ID: rv}
}

// The four-element Scene Kit vector representation of the value.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/scnVector4Value
func (v NSValue) SCNVector4Value() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("SCNVector4Value"))
	return objectivec.Object{ID: rv}
}

// The Scene Kit 4 x 4 matrix representation of the value.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/scnMatrix4Value
func (v NSValue) SCNMatrix4Value() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("SCNMatrix4Value"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Foundation/NSValue/edgeInsetsValue
func (v NSValue) EdgeInsetsValue() NSEdgeInsets {
	rv := objc.Send[NSEdgeInsets](v.ID, objc.Sel("edgeInsetsValue"))
	return NSEdgeInsets(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSValue/gcPoint2Value
func (v NSValue) GCPoint2Value() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("GCPoint2Value"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Foundation/NSValue/videoDimensionsValue
func (v NSValue) CMVideoDimensionsValue() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("CMVideoDimensionsValue"))
	return objectivec.Object{ID: rv}
}

// Returns an integer that can be used as a table address in a hash table
// structure.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/hash
func (v NSValue) Hash() int {
	rv := objc.Send[int](v.ID, objc.Sel("hash"))
	return rv
}
func (v NSValue) SetHash(value int) {
	objc.Send[struct{}](v.ID, objc.Sel("setHash:"), value)
}

// Protocol methods for NSCopying

// Protocol methods for NSSecureCoding
