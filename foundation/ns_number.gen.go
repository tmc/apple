// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSNumber] class.
var (
	_NSNumberClass     NSNumberClass
	_NSNumberClassOnce sync.Once
)

func getNSNumberClass() NSNumberClass {
	_NSNumberClassOnce.Do(func() {
		_NSNumberClass = NSNumberClass{class: objc.GetClass("NSNumber")}
	})
	return _NSNumberClass
}

// GetNSNumberClass returns the class object for NSNumber.
func GetNSNumberClass() NSNumberClass {
	return getNSNumberClass()
}

type NSNumberClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSNumberClass) Alloc() NSNumber {
	rv := objc.Send[NSNumber](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An object wrapper for primitive scalar numeric values.
//
// # Overview
// 
// [NSNumber] is a subclass of [NSValue] that offers a value as any C scalar
// (numeric) type. It defines a set of methods specifically for setting and
// accessing the value as a signed or unsigned `char`, `short int`, `int`,
// `long int`, `long long int`, `float`, or `double` or as a [BOOL]. (Note
// that number objects do not necessarily preserve the type they are created
// with.) It also defines a [NSNumber.Compare] method to determine the ordering of two
// [NSNumber] objects.
// 
// [NSNumber] is “toll-free bridged” with its Core Foundation
// counterparts: [CFNumber] for integer and floating point values, and
// [CFBoolean] for Boolean values. See [Toll-Free Bridging] for more
// information on toll-free bridging.
// 
// # Value Conversions
// 
// [NSNumber] provides readonly properties that return the object’s stored
// value converted to a particular Boolean, integer, unsigned integer, or
// floating point C scalar type. Because numeric types have different storage
// capabilities, attempting to initialize with a value of one type and access
// the value of another type may produce an erroneous result—for example,
// initializing with a `double` value exceeding `FLT_MAX` and accessing its
// [NSNumber.FloatValue], or initializing with an negative integer value and accessing
// its [NSNumber.UnsignedIntegerValue]. In some cases, attempting to initialize with a
// value of a type and access the value of another type may result in loss of
// precision—for example, initializing with a `double` value with many
// significant digits and accessing its [NSNumber.FloatValue], or initializing with a
// large integer value and accessing its [NSNumber.CharValue].
// 
// An [NSNumber] object initialized with a value of a particular type
// accessing the converted value of a different of type, such as `unsigned
// int` and `float`, will convert its stored value to that converted type in
// the following ways:
// 
// [Table data omitted]
// 
// [Table data omitted]
// 
// [Table data omitted]
// 
// [Table data omitted]
// 
// # Subclassing Notes
// 
// As with any class cluster, subclasses of [NSNumber] must override the
// primitive methods of its superclass, [NSValue]. In addition, there are two
// requirements around the data type your subclass represents:
// 
// - Your implementation of [NSNumber.ObjCType] must return one of “`c`”,
// “[C]”, “`s`”, “[S]”, “`i`”, “[I]”, “`l`”,
// “[L]”, “`q`”, “[Q]”, “`f`”, and “`d`”. This is required
// for the other methods of [NSNumber] to behave correctly. - Your subclass
// must override the accessor method that corresponds to the declared
// type—for example, if your implementation of [NSNumber.ObjCType] returns “`i`”,
// you must override [NSNumber.IntValue].
//
// [CFBoolean]: https://developer.apple.com/documentation/CoreFoundation/CFBoolean
// [CFNumber]: https://developer.apple.com/documentation/CoreFoundation/CFNumber
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
//
// # Initializing an NSNumber Object
//
//   - [NSNumber.InitWithBool]: Returns an [NSNumber] object initialized to contain a given value, treated as a [BOOL].
//   - [NSNumber.InitWithChar]: Returns an [NSNumber] object initialized to contain a given value, treated as a signed `char`.
//   - [NSNumber.InitWithDouble]: Returns an [NSNumber] object initialized to contain `value`, treated as a `double`.
//   - [NSNumber.InitWithFloat]: Returns an [NSNumber] object initialized to contain a given value, treated as a `float`.
//   - [NSNumber.InitWithInt]: Returns an [NSNumber] object initialized to contain a given value, treated as a signed `int`.
//   - [NSNumber.InitWithInteger]: Returns an [NSNumber] object initialized to contain a given value, treated as an [NSInteger].
//   - [NSNumber.InitWithLongLong]: Returns an [NSNumber] object initialized to contain `value`, treated as a signed `long long`.
//   - [NSNumber.InitWithShort]: Returns an [NSNumber] object initialized to contain a given value, treated as a signed `short`.
//   - [NSNumber.InitWithUnsignedChar]: Returns an [NSNumber] object initialized to contain a given value, treated as an `unsigned char`.
//   - [NSNumber.InitWithUnsignedInt]: Returns an [NSNumber] object initialized to contain a given value, treated as an `unsigned int`.
//   - [NSNumber.InitWithUnsignedInteger]: Returns an [NSNumber] object initialized to contain a given value, treated as an [NSUInteger].
//   - [NSNumber.InitWithUnsignedLongLong]: Returns an [NSNumber] object initialized to contain a given value, treated as an `unsigned long long`.
//   - [NSNumber.InitWithUnsignedShort]: Returns an [NSNumber] object initialized to contain a given value, treated as an `unsigned short`.
//
// # Accessing Numeric Values
//
//   - [NSNumber.BoolValue]: The number object’s value expressed as a Boolean value.
//   - [NSNumber.CharValue]: The number object’s value expressed as a `char`.
//   - [NSNumber.DecimalValue]: The number object’s value expressed as an [Decimal](<doc://com.apple.foundation/documentation/Foundation/Decimal>) structure.
//   - [NSNumber.DoubleValue]: The number object’s value expressed as a `double`, converted as necessary.
//   - [NSNumber.FloatValue]: The number object’s value expressed as a `float`, converted as necessary.
//   - [NSNumber.IntValue]: The number object’s value expressed as an `int`, converted as necessary.
//   - [NSNumber.IntegerValue]: The number object’s value expressed as an [NSInteger] object, converted as necessary.
//   - [NSNumber.LongLongValue]: The number object’s value expressed as a `long long`, converted as necessary.
//   - [NSNumber.ShortValue]: The number object’s value expressed as a `short`, converted as necessary.
//   - [NSNumber.UnsignedCharValue]: The number object’s value expressed as an unsigned `char`, converted as necessary.
//   - [NSNumber.UnsignedIntegerValue]: The number object’s value expressed as an [NSUInteger] object, converted as necessary.
//   - [NSNumber.UnsignedIntValue]: The number object’s value expressed as an unsigned `int`, converted as necessary.
//   - [NSNumber.UnsignedLongLongValue]: The number object’s value expressed as an unsigned `long long`, converted as necessary.
//   - [NSNumber.UnsignedShortValue]: The number object’s value expressed as an unsigned `short`, converted as necessary.
//
// # Retrieving String Representations
//
//   - [NSNumber.DescriptionWithLocale]: Returns a string that represents the contents of the number object for a given locale.
//   - [NSNumber.StringValue]: The number object’s value expressed as a human-readable string.
//
// # Comparing NSNumber Objects
//
//   - [NSNumber.Compare]: Returns an [NSComparisonResult] value that indicates whether the number object’s value is greater than, equal to, or less than a given number.
//   - [NSNumber.IsEqualToNumber]: Returns a Boolean value that indicates whether the number object’s value and a given number are equal.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber
type NSNumber struct {
	NSValue
}

// NSNumberFromID constructs a [NSNumber] from an objc.ID.
//
// An object wrapper for primitive scalar numeric values.
func NSNumberFromID(id objc.ID) NSNumber {
	return NSNumber{NSValue: NSValueFromID(id)}
}
// NOTE: NSNumber adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSNumber] class.
//
// # Initializing an NSNumber Object
//
//   - [INSNumber.InitWithBool]: Returns an [NSNumber] object initialized to contain a given value, treated as a [BOOL].
//   - [INSNumber.InitWithChar]: Returns an [NSNumber] object initialized to contain a given value, treated as a signed `char`.
//   - [INSNumber.InitWithDouble]: Returns an [NSNumber] object initialized to contain `value`, treated as a `double`.
//   - [INSNumber.InitWithFloat]: Returns an [NSNumber] object initialized to contain a given value, treated as a `float`.
//   - [INSNumber.InitWithInt]: Returns an [NSNumber] object initialized to contain a given value, treated as a signed `int`.
//   - [INSNumber.InitWithInteger]: Returns an [NSNumber] object initialized to contain a given value, treated as an [NSInteger].
//   - [INSNumber.InitWithLongLong]: Returns an [NSNumber] object initialized to contain `value`, treated as a signed `long long`.
//   - [INSNumber.InitWithShort]: Returns an [NSNumber] object initialized to contain a given value, treated as a signed `short`.
//   - [INSNumber.InitWithUnsignedChar]: Returns an [NSNumber] object initialized to contain a given value, treated as an `unsigned char`.
//   - [INSNumber.InitWithUnsignedInt]: Returns an [NSNumber] object initialized to contain a given value, treated as an `unsigned int`.
//   - [INSNumber.InitWithUnsignedInteger]: Returns an [NSNumber] object initialized to contain a given value, treated as an [NSUInteger].
//   - [INSNumber.InitWithUnsignedLongLong]: Returns an [NSNumber] object initialized to contain a given value, treated as an `unsigned long long`.
//   - [INSNumber.InitWithUnsignedShort]: Returns an [NSNumber] object initialized to contain a given value, treated as an `unsigned short`.
//
// # Accessing Numeric Values
//
//   - [INSNumber.BoolValue]: The number object’s value expressed as a Boolean value.
//   - [INSNumber.CharValue]: The number object’s value expressed as a `char`.
//   - [INSNumber.DecimalValue]: The number object’s value expressed as an [Decimal](<doc://com.apple.foundation/documentation/Foundation/Decimal>) structure.
//   - [INSNumber.DoubleValue]: The number object’s value expressed as a `double`, converted as necessary.
//   - [INSNumber.FloatValue]: The number object’s value expressed as a `float`, converted as necessary.
//   - [INSNumber.IntValue]: The number object’s value expressed as an `int`, converted as necessary.
//   - [INSNumber.IntegerValue]: The number object’s value expressed as an [NSInteger] object, converted as necessary.
//   - [INSNumber.LongLongValue]: The number object’s value expressed as a `long long`, converted as necessary.
//   - [INSNumber.ShortValue]: The number object’s value expressed as a `short`, converted as necessary.
//   - [INSNumber.UnsignedCharValue]: The number object’s value expressed as an unsigned `char`, converted as necessary.
//   - [INSNumber.UnsignedIntegerValue]: The number object’s value expressed as an [NSUInteger] object, converted as necessary.
//   - [INSNumber.UnsignedIntValue]: The number object’s value expressed as an unsigned `int`, converted as necessary.
//   - [INSNumber.UnsignedLongLongValue]: The number object’s value expressed as an unsigned `long long`, converted as necessary.
//   - [INSNumber.UnsignedShortValue]: The number object’s value expressed as an unsigned `short`, converted as necessary.
//
// # Retrieving String Representations
//
//   - [INSNumber.DescriptionWithLocale]: Returns a string that represents the contents of the number object for a given locale.
//   - [INSNumber.StringValue]: The number object’s value expressed as a human-readable string.
//
// # Comparing NSNumber Objects
//
//   - [INSNumber.Compare]: Returns an [NSComparisonResult] value that indicates whether the number object’s value is greater than, equal to, or less than a given number.
//   - [INSNumber.IsEqualToNumber]: Returns a Boolean value that indicates whether the number object’s value and a given number are equal.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber
type INSNumber interface {
	INSValue

	// Topic: Initializing an NSNumber Object

	// Returns an [NSNumber] object initialized to contain a given value, treated as a [BOOL].
	InitWithBool(value bool) NSNumber
	// Returns an [NSNumber] object initialized to contain a given value, treated as a signed `char`.
	InitWithChar(value int8) NSNumber
	// Returns an [NSNumber] object initialized to contain `value`, treated as a `double`.
	InitWithDouble(value float64) NSNumber
	// Returns an [NSNumber] object initialized to contain a given value, treated as a `float`.
	InitWithFloat(value float32) NSNumber
	// Returns an [NSNumber] object initialized to contain a given value, treated as a signed `int`.
	InitWithInt(value int) NSNumber
	// Returns an [NSNumber] object initialized to contain a given value, treated as an [NSInteger].
	InitWithInteger(value int) NSNumber
	// Returns an [NSNumber] object initialized to contain `value`, treated as a signed `long long`.
	InitWithLongLong(value int64) NSNumber
	// Returns an [NSNumber] object initialized to contain a given value, treated as a signed `short`.
	InitWithShort(value int16) NSNumber
	// Returns an [NSNumber] object initialized to contain a given value, treated as an `unsigned char`.
	InitWithUnsignedChar(value byte) NSNumber
	// Returns an [NSNumber] object initialized to contain a given value, treated as an `unsigned int`.
	InitWithUnsignedInt(value uint32) NSNumber
	// Returns an [NSNumber] object initialized to contain a given value, treated as an [NSUInteger].
	InitWithUnsignedInteger(value uint) NSNumber
	// Returns an [NSNumber] object initialized to contain a given value, treated as an `unsigned long long`.
	InitWithUnsignedLongLong(value uint64) NSNumber
	// Returns an [NSNumber] object initialized to contain a given value, treated as an `unsigned short`.
	InitWithUnsignedShort(value uint16) NSNumber

	// Topic: Accessing Numeric Values

	// The number object’s value expressed as a Boolean value.
	BoolValue() bool
	// The number object’s value expressed as a `char`.
	CharValue() int8
	// The number object’s value expressed as an [Decimal](<doc://com.apple.foundation/documentation/Foundation/Decimal>) structure.
	DecimalValue() Decimal
	// The number object’s value expressed as a `double`, converted as necessary.
	DoubleValue() float64
	// The number object’s value expressed as a `float`, converted as necessary.
	FloatValue() float32
	// The number object’s value expressed as an `int`, converted as necessary.
	IntValue() int
	// The number object’s value expressed as an [NSInteger] object, converted as necessary.
	IntegerValue() int
	// The number object’s value expressed as a `long long`, converted as necessary.
	LongLongValue() int64
	// The number object’s value expressed as a `short`, converted as necessary.
	ShortValue() int16
	// The number object’s value expressed as an unsigned `char`, converted as necessary.
	UnsignedCharValue() byte
	// The number object’s value expressed as an [NSUInteger] object, converted as necessary.
	UnsignedIntegerValue() uint
	// The number object’s value expressed as an unsigned `int`, converted as necessary.
	UnsignedIntValue() uint32
	// The number object’s value expressed as an unsigned `long long`, converted as necessary.
	UnsignedLongLongValue() uint64
	// The number object’s value expressed as an unsigned `short`, converted as necessary.
	UnsignedShortValue() uint16

	// Topic: Retrieving String Representations

	// Returns a string that represents the contents of the number object for a given locale.
	DescriptionWithLocale(locale objectivec.IObject) string
	// The number object’s value expressed as a human-readable string.
	StringValue() string

	// Topic: Comparing NSNumber Objects

	// Returns an [NSComparisonResult] value that indicates whether the number object’s value is greater than, equal to, or less than a given number.
	Compare(otherNumber INSNumber) ComparisonResult
	// Returns a Boolean value that indicates whether the number object’s value and a given number are equal.
	IsEqualToNumber(number INSNumber) bool

	// The number object’s value expressed as a `long`, converted as necessary.
	LongValue() int
	// The number object’s value expressed as an unsigned `long`, converted as necessary.
	UnsignedLongValue() uint64
	// Returns an [NSNumber] object initialized to contain a given value, treated as a signed `long`.
	InitWithLong(value int) NSNumber
	// Returns an [NSNumber] object initialized to contain a given value, treated as an `unsigned long`.
	InitWithUnsignedLong(value uint64) NSNumber
}





// Init initializes the instance.
func (n NSNumber) Init() NSNumber {
	rv := objc.Send[NSNumber](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NSNumber) Autorelease() NSNumber {
	rv := objc.Send[NSNumber](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSNumber creates a new NSNumber instance.
func NewNSNumber() NSNumber {
	class := getNSNumberClass()
	rv := objc.Send[NSNumber](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(GCPoint2:)
// point is a [gamecontroller.GCPoint2].
func NewNumberValueWithGCPoint2(point objectivec.IObject) NSNumber {
	rv := objc.Send[objc.ID](objc.ID(getNSNumberClass().class), objc.Sel("valueWithGCPoint2:"), point)
	return NSNumberFromID(rv)
}


// Returns an [NSNumber] object initialized to contain a given value, treated
// as a [BOOL].
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as a [BOOL].
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-1ojz2
func NewNumberWithBool(value bool) NSNumber {
	instance := getNSNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBool:"), value)
	return NSNumberFromID(rv)
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
// [Number and Value Programming Topics]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/NumbersandValues/NumbersandValues.html#//apple_ref/doc/uid/10000038i
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(bytes:objCType:)
func NewNumberWithBytesObjCType(value unsafe.Pointer, type_ []byte) NSNumber {
	instance := getNSNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytes:objCType:"), value, unsafe.Pointer(unsafe.SliceData(type_)))
	return NSNumberFromID(rv)
}


// Returns an [NSNumber] object initialized to contain a given value, treated
// as a signed `char`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as a signed `char`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-8krjs
func NewNumberWithChar(value int8) NSNumber {
	instance := getNSNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithChar:"), value)
	return NSNumberFromID(rv)
}


//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(coder:)
func NewNumberWithCoder(coder INSCoder) NSNumber {
	instance := getNSNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSNumberFromID(rv)
}


// Returns an [NSNumber] object initialized to contain `value`, treated as a
// `double`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as a `double`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-15chk
func NewNumberWithDouble(value float64) NSNumber {
	instance := getNSNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDouble:"), value)
	return NSNumberFromID(rv)
}


// Returns an [NSNumber] object initialized to contain a given value, treated
// as a `float`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as a `float`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-2vlwk
func NewNumberWithFloat(value float32) NSNumber {
	instance := getNSNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFloat:"), value)
	return NSNumberFromID(rv)
}


// Returns an [NSNumber] object initialized to contain a given value, treated
// as a signed `int`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as a signed `int`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-7jvmg
func NewNumberWithInt(value int) NSNumber {
	instance := getNSNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInt:"), value)
	return NSNumberFromID(rv)
}


// Returns an [NSNumber] object initialized to contain a given value, treated
// as an [NSInteger].
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as an [NSInteger].
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-5jcjl
func NewNumberWithInteger(value int) NSNumber {
	instance := getNSNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInteger:"), value)
	return NSNumberFromID(rv)
}


// Returns an [NSNumber] object initialized to contain a given value, treated
// as a signed `long`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as a signed `long`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/initWithLong:
func NewNumberWithLong(value int) NSNumber {
	instance := getNSNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLong:"), value)
	return NSNumberFromID(rv)
}


// Returns an [NSNumber] object initialized to contain `value`, treated as a
// signed `long long`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as a signed `long
// long`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-40ad0
func NewNumberWithLongLong(value int64) NSNumber {
	instance := getNSNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLongLong:"), value)
	return NSNumberFromID(rv)
}


// Returns an [NSNumber] object initialized to contain a given value, treated
// as a signed `short`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as a signed `short`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-16drx
func NewNumberWithShort(value int16) NSNumber {
	instance := getNSNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithShort:"), value)
	return NSNumberFromID(rv)
}


// Returns an [NSNumber] object initialized to contain a given value, treated
// as an `unsigned char`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as an `unsigned char`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-8se67
func NewNumberWithUnsignedChar(value byte) NSNumber {
	instance := getNSNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUnsignedChar:"), value)
	return NSNumberFromID(rv)
}


// Returns an [NSNumber] object initialized to contain a given value, treated
// as an `unsigned int`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as an `unsigned int`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-47coa
func NewNumberWithUnsignedInt(value uint32) NSNumber {
	instance := getNSNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUnsignedInt:"), value)
	return NSNumberFromID(rv)
}


// Returns an [NSNumber] object initialized to contain a given value, treated
// as an [NSUInteger].
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as an [NSUInteger].
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-3l4ek
func NewNumberWithUnsignedInteger(value uint) NSNumber {
	instance := getNSNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUnsignedInteger:"), value)
	return NSNumberFromID(rv)
}


// Returns an [NSNumber] object initialized to contain a given value, treated
// as an `unsigned long`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as an `unsigned long`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/initWithUnsignedLong:
func NewNumberWithUnsignedLong(value uint64) NSNumber {
	instance := getNSNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUnsignedLong:"), value)
	return NSNumberFromID(rv)
}


// Returns an [NSNumber] object initialized to contain a given value, treated
// as an `unsigned long long`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as an `unsigned long
// long`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-43lc7
func NewNumberWithUnsignedLongLong(value uint64) NSNumber {
	instance := getNSNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUnsignedLongLong:"), value)
	return NSNumberFromID(rv)
}


// Returns an [NSNumber] object initialized to contain a given value, treated
// as an `unsigned short`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as an `unsigned
// short`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-87y9m
func NewNumberWithUnsignedShort(value uint16) NSNumber {
	instance := getNSNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUnsignedShort:"), value)
	return NSNumberFromID(rv)
}







// Returns an [NSNumber] object initialized to contain a given value, treated
// as a [BOOL].
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as a [BOOL].
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-1ojz2
func (n NSNumber) InitWithBool(value bool) NSNumber {
	rv := objc.Send[NSNumber](n.ID, objc.Sel("initWithBool:"), value)
	return rv
}

// Returns an [NSNumber] object initialized to contain a given value, treated
// as a signed `char`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as a signed `char`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-8krjs
func (n NSNumber) InitWithChar(value int8) NSNumber {
	rv := objc.Send[NSNumber](n.ID, objc.Sel("initWithChar:"), value)
	return rv
}

// Returns an [NSNumber] object initialized to contain `value`, treated as a
// `double`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as a `double`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-15chk
func (n NSNumber) InitWithDouble(value float64) NSNumber {
	rv := objc.Send[NSNumber](n.ID, objc.Sel("initWithDouble:"), value)
	return rv
}

// Returns an [NSNumber] object initialized to contain a given value, treated
// as a `float`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as a `float`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-2vlwk
func (n NSNumber) InitWithFloat(value float32) NSNumber {
	rv := objc.Send[NSNumber](n.ID, objc.Sel("initWithFloat:"), value)
	return rv
}

// Returns an [NSNumber] object initialized to contain a given value, treated
// as a signed `int`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as a signed `int`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-7jvmg
func (n NSNumber) InitWithInt(value int) NSNumber {
	rv := objc.Send[NSNumber](n.ID, objc.Sel("initWithInt:"), value)
	return rv
}

// Returns an [NSNumber] object initialized to contain a given value, treated
// as an [NSInteger].
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as an [NSInteger].
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-5jcjl
func (n NSNumber) InitWithInteger(value int) NSNumber {
	rv := objc.Send[NSNumber](n.ID, objc.Sel("initWithInteger:"), value)
	return rv
}

// Returns an [NSNumber] object initialized to contain `value`, treated as a
// signed `long long`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as a signed `long
// long`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-40ad0
func (n NSNumber) InitWithLongLong(value int64) NSNumber {
	rv := objc.Send[NSNumber](n.ID, objc.Sel("initWithLongLong:"), value)
	return rv
}

// Returns an [NSNumber] object initialized to contain a given value, treated
// as a signed `short`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as a signed `short`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-16drx
func (n NSNumber) InitWithShort(value int16) NSNumber {
	rv := objc.Send[NSNumber](n.ID, objc.Sel("initWithShort:"), value)
	return rv
}

// Returns an [NSNumber] object initialized to contain a given value, treated
// as an `unsigned char`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as an `unsigned char`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-8se67
func (n NSNumber) InitWithUnsignedChar(value byte) NSNumber {
	rv := objc.Send[NSNumber](n.ID, objc.Sel("initWithUnsignedChar:"), value)
	return rv
}

// Returns an [NSNumber] object initialized to contain a given value, treated
// as an `unsigned int`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as an `unsigned int`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-47coa
func (n NSNumber) InitWithUnsignedInt(value uint32) NSNumber {
	rv := objc.Send[NSNumber](n.ID, objc.Sel("initWithUnsignedInt:"), value)
	return rv
}

// Returns an [NSNumber] object initialized to contain a given value, treated
// as an [NSUInteger].
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as an [NSUInteger].
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-3l4ek
func (n NSNumber) InitWithUnsignedInteger(value uint) NSNumber {
	rv := objc.Send[NSNumber](n.ID, objc.Sel("initWithUnsignedInteger:"), value)
	return rv
}

// Returns an [NSNumber] object initialized to contain a given value, treated
// as an `unsigned long long`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as an `unsigned long
// long`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-43lc7
func (n NSNumber) InitWithUnsignedLongLong(value uint64) NSNumber {
	rv := objc.Send[NSNumber](n.ID, objc.Sel("initWithUnsignedLongLong:"), value)
	return rv
}

// Returns an [NSNumber] object initialized to contain a given value, treated
// as an `unsigned short`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as an `unsigned
// short`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/init(value:)-87y9m
func (n NSNumber) InitWithUnsignedShort(value uint16) NSNumber {
	rv := objc.Send[NSNumber](n.ID, objc.Sel("initWithUnsignedShort:"), value)
	return rv
}

// Returns a string that represents the contents of the number object for a
// given locale.
//
// locale: An object containing locale information with which to format the
// description. Use `nil` if you don’t want the description formatted.
//
// # Return Value
// 
// A string that represents the contents of the number object formatted using
// the locale information in `locale`.
//
// # Discussion
// 
// For example, if you have an [NSNumber] object that has the integer value
// 522, sending it the [DescriptionWithLocale] message returns the string
// “522”.
// 
// To obtain the string representation, this method invokes [NSString]’s
// [InitWithFormatLocale] method, supplying the format based on the type the
// [NSNumber] object was created with:
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/description(withLocale:)
func (n NSNumber) DescriptionWithLocale(locale objectivec.IObject) string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("descriptionWithLocale:"), locale)
	return NSStringFromID(rv).String()
}

// Returns an [NSComparisonResult] value that indicates whether the number
// object’s value is greater than, equal to, or less than a given number.
//
// otherNumber: The number to compare to the number object’s value.
// 
// This value must not be `nil`. If the value is `nil`, the behavior is
// undefined and may change in future versions of macOS.
//
// # Return Value
// 
// [NSOrderedAscending] if the value of `otherNumber` is greater than the
// number object’s, [NSOrderedSame] if they’re equal, and
// [NSOrderedDescending] if the value of `otherNumber` is less than the number
// object’s.
//
// # Discussion
// 
// The [Compare] method follows the standard C rules for type conversion. For
// example, if you compare an [NSNumber] object that has an integer value with
// an [NSNumber] object that has a floating point value, the integer value is
// converted to a floating-point value for comparison.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/compare(_:)
func (n NSNumber) Compare(otherNumber INSNumber) ComparisonResult {
	rv := objc.Send[ComparisonResult](n.ID, objc.Sel("compare:"), otherNumber)
	return ComparisonResult(rv)
}

// Returns a Boolean value that indicates whether the number object’s value
// and a given number are equal.
//
// number: The number to compare to the number object’s value.
//
// # Return Value
// 
// [true] if the number object’s value and `number` are equal, otherwise
// [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Two [NSNumber] objects are considered equal if they have the same id values
// or if they have equivalent values (as determined by the [Compare] method).
// 
// This method is more efficient than [Compare] if you know the two objects
// are numbers.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/isEqual(to:)
func (n NSNumber) IsEqualToNumber(number INSNumber) bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isEqualToNumber:"), number)
	return rv
}

// Returns an [NSNumber] object initialized to contain a given value, treated
// as a signed `long`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as a signed `long`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/initWithLong:
func (n NSNumber) InitWithLong(value int) NSNumber {
	rv := objc.Send[NSNumber](n.ID, objc.Sel("initWithLong:"), value)
	return rv
}

// Returns an [NSNumber] object initialized to contain a given value, treated
// as an `unsigned long`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as an `unsigned long`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/initWithUnsignedLong:
func (n NSNumber) InitWithUnsignedLong(value uint64) NSNumber {
	rv := objc.Send[NSNumber](n.ID, objc.Sel("initWithUnsignedLong:"), value)
	return rv
}





// Creates and returns an [NSNumber] object containing a given value, treating
// it as a [BOOL].
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as a [BOOL].
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithBool:
func (_NSNumberClass NSNumberClass) NumberWithBool(value bool) NSNumber {
	rv := objc.Send[objc.ID](objc.ID(_NSNumberClass.class), objc.Sel("numberWithBool:"), value)
	return NSNumberFromID(rv)
}

// Creates and returns an [NSNumber] object containing a given value, treating
// it as a signed `char`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as a signed `char`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithChar:
func (_NSNumberClass NSNumberClass) NumberWithChar(value int8) NSNumber {
	rv := objc.Send[objc.ID](objc.ID(_NSNumberClass.class), objc.Sel("numberWithChar:"), value)
	return NSNumberFromID(rv)
}

// Creates and returns an [NSNumber] object containing a given value, treating
// it as a `double`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as a `double`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithDouble:
func (_NSNumberClass NSNumberClass) NumberWithDouble(value float64) NSNumber {
	rv := objc.Send[objc.ID](objc.ID(_NSNumberClass.class), objc.Sel("numberWithDouble:"), value)
	return NSNumberFromID(rv)
}

// Creates and returns an [NSNumber] object containing a given value, treating
// it as a `float`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as a `float`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithFloat:
func (_NSNumberClass NSNumberClass) NumberWithFloat(value float32) NSNumber {
	rv := objc.Send[objc.ID](objc.ID(_NSNumberClass.class), objc.Sel("numberWithFloat:"), value)
	return NSNumberFromID(rv)
}

// Creates and returns an [NSNumber] object containing a given value, treating
// it as a signed `int`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as a signed `int`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithInt:
func (_NSNumberClass NSNumberClass) NumberWithInt(value int) NSNumber {
	rv := objc.Send[objc.ID](objc.ID(_NSNumberClass.class), objc.Sel("numberWithInt:"), value)
	return NSNumberFromID(rv)
}

// Creates and returns an [NSNumber] object containing a given value, treating
// it as an [NSInteger].
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as an [NSInteger].
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithInteger:
func (_NSNumberClass NSNumberClass) NumberWithInteger(value int) NSNumber {
	rv := objc.Send[objc.ID](objc.ID(_NSNumberClass.class), objc.Sel("numberWithInteger:"), value)
	return NSNumberFromID(rv)
}

// Creates and returns an [NSNumber] object containing a given value, treating
// it as a signed `long`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as a signed `long`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithLong:
func (_NSNumberClass NSNumberClass) NumberWithLong(value int) NSNumber {
	rv := objc.Send[objc.ID](objc.ID(_NSNumberClass.class), objc.Sel("numberWithLong:"), value)
	return NSNumberFromID(rv)
}

// Creates and returns an [NSNumber] object containing a given value, treating
// it as a signed `long long`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as a signed `long
// long`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithLongLong:
func (_NSNumberClass NSNumberClass) NumberWithLongLong(value int64) NSNumber {
	rv := objc.Send[objc.ID](objc.ID(_NSNumberClass.class), objc.Sel("numberWithLongLong:"), value)
	return NSNumberFromID(rv)
}

// Creates and returns an [NSNumber] object containing `value`, treating it as
// a signed `short`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as a signed `short`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithShort:
func (_NSNumberClass NSNumberClass) NumberWithShort(value int16) NSNumber {
	rv := objc.Send[objc.ID](objc.ID(_NSNumberClass.class), objc.Sel("numberWithShort:"), value)
	return NSNumberFromID(rv)
}

// Creates and returns an [NSNumber] object containing a given value, treating
// it as an `unsigned char`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as an `unsigned char`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithUnsignedChar:
func (_NSNumberClass NSNumberClass) NumberWithUnsignedChar(value byte) NSNumber {
	rv := objc.Send[objc.ID](objc.ID(_NSNumberClass.class), objc.Sel("numberWithUnsignedChar:"), value)
	return NSNumberFromID(rv)
}

// Creates and returns an [NSNumber] object containing a given value, treating
// it as an `unsigned int`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as an `unsigned int`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithUnsignedInt:
func (_NSNumberClass NSNumberClass) NumberWithUnsignedInt(value uint32) NSNumber {
	rv := objc.Send[objc.ID](objc.ID(_NSNumberClass.class), objc.Sel("numberWithUnsignedInt:"), value)
	return NSNumberFromID(rv)
}

// Creates and returns an [NSNumber] object containing a given value, treating
// it as an [NSUInteger].
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as an [NSUInteger].
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithUnsignedInteger:
func (_NSNumberClass NSNumberClass) NumberWithUnsignedInteger(value uint) NSNumber {
	rv := objc.Send[objc.ID](objc.ID(_NSNumberClass.class), objc.Sel("numberWithUnsignedInteger:"), value)
	return NSNumberFromID(rv)
}

// Creates and returns an [NSNumber] object containing a given value, treating
// it as an `unsigned long`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as an `unsigned long`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithUnsignedLong:
func (_NSNumberClass NSNumberClass) NumberWithUnsignedLong(value uint64) NSNumber {
	rv := objc.Send[objc.ID](objc.ID(_NSNumberClass.class), objc.Sel("numberWithUnsignedLong:"), value)
	return NSNumberFromID(rv)
}

// Creates and returns an [NSNumber] object containing a given value, treating
// it as an `unsigned long long`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as an `unsigned long
// long`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithUnsignedLongLong:
func (_NSNumberClass NSNumberClass) NumberWithUnsignedLongLong(value uint64) NSNumber {
	rv := objc.Send[objc.ID](objc.ID(_NSNumberClass.class), objc.Sel("numberWithUnsignedLongLong:"), value)
	return NSNumberFromID(rv)
}

// Creates and returns an [NSNumber] object containing a given value, treating
// it as an `unsigned short`.
//
// value: The value for the new number.
//
// # Return Value
// 
// An [NSNumber] object containing `value`, treating it as an `unsigned
// short`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/numberWithUnsignedShort:
func (_NSNumberClass NSNumberClass) NumberWithUnsignedShort(value uint16) NSNumber {
	rv := objc.Send[objc.ID](objc.ID(_NSNumberClass.class), objc.Sel("numberWithUnsignedShort:"), value)
	return NSNumberFromID(rv)
}








// The number object’s value expressed as a Boolean value.
//
// # Discussion
// 
// A `0` value always means [false], and any nonzero value is interpreted as
// [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/boolValue
func (n NSNumber) BoolValue() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("boolValue"))
	return rv
}



// The number object’s value expressed as a `char`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/int8Value
func (n NSNumber) CharValue() int8 {
	rv := objc.Send[int8](n.ID, objc.Sel("charValue"))
	return rv
}



// The number object’s value expressed as an [Decimal] structure.
//
// [Decimal]: https://developer.apple.com/documentation/Foundation/Decimal
//
// # Discussion
// 
// The [Decimal] value isn’t guaranteed to be exact for `float` and `double`
// values.
//
// [Decimal]: https://developer.apple.com/documentation/Foundation/Decimal
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/decimalValue
func (n NSNumber) DecimalValue() Decimal {
	rv := objc.Send[Decimal](n.ID, objc.Sel("decimalValue"))
	return Decimal(rv)
}



// The number object’s value expressed as a `double`, converted as
// necessary.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/doubleValue
func (n NSNumber) DoubleValue() float64 {
	rv := objc.Send[float64](n.ID, objc.Sel("doubleValue"))
	return rv
}



// The number object’s value expressed as a `float`, converted as necessary.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/floatValue
func (n NSNumber) FloatValue() float32 {
	rv := objc.Send[float32](n.ID, objc.Sel("floatValue"))
	return rv
}



// The number object’s value expressed as an `int`, converted as necessary.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/int32Value
func (n NSNumber) IntValue() int {
	rv := objc.Send[int](n.ID, objc.Sel("intValue"))
	return rv
}



// The number object’s value expressed as an [NSInteger] object, converted
// as necessary.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/intValue-95zzp
func (n NSNumber) IntegerValue() int {
	rv := objc.Send[int](n.ID, objc.Sel("integerValue"))
	return rv
}



// The number object’s value expressed as a `long long`, converted as
// necessary.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/int64Value
func (n NSNumber) LongLongValue() int64 {
	rv := objc.Send[int64](n.ID, objc.Sel("longLongValue"))
	return rv
}



// The number object’s value expressed as a `short`, converted as necessary.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/int16Value
func (n NSNumber) ShortValue() int16 {
	rv := objc.Send[int16](n.ID, objc.Sel("shortValue"))
	return rv
}



// The number object’s value expressed as an unsigned `char`, converted as
// necessary.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/uint8Value
func (n NSNumber) UnsignedCharValue() byte {
	rv := objc.Send[byte](n.ID, objc.Sel("unsignedCharValue"))
	return rv
}



// The number object’s value expressed as an [NSUInteger] object, converted
// as necessary.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/uintValue
func (n NSNumber) UnsignedIntegerValue() uint {
	rv := objc.Send[uint](n.ID, objc.Sel("unsignedIntegerValue"))
	return rv
}



// The number object’s value expressed as an unsigned `int`, converted as
// necessary.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/uint32Value
func (n NSNumber) UnsignedIntValue() uint32 {
	rv := objc.Send[uint32](n.ID, objc.Sel("unsignedIntValue"))
	return rv
}



// The number object’s value expressed as an unsigned `long long`, converted
// as necessary.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/uint64Value
func (n NSNumber) UnsignedLongLongValue() uint64 {
	rv := objc.Send[uint64](n.ID, objc.Sel("unsignedLongLongValue"))
	return rv
}



// The number object’s value expressed as an unsigned `short`, converted as
// necessary.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/uint16Value
func (n NSNumber) UnsignedShortValue() uint16 {
	rv := objc.Send[uint16](n.ID, objc.Sel("unsignedShortValue"))
	return rv
}



// The number object’s value expressed as a human-readable string.
//
// # Discussion
// 
// The string is created by invoking [DescriptionWithLocale] where locale is
// `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/stringValue
func (n NSNumber) StringValue() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("stringValue"))
	return NSStringFromID(rv).String()
}



// The number object’s value expressed as a `long`, converted as necessary.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/longValue
func (n NSNumber) LongValue() int {
	rv := objc.Send[int](n.ID, objc.Sel("longValue"))
	return rv
}



// The number object’s value expressed as an unsigned `long`, converted as
// necessary.
//
// See: https://developer.apple.com/documentation/Foundation/NSNumber/unsignedLongValue
func (n NSNumber) UnsignedLongValue() uint64 {
	rv := objc.Send[uint64](n.ID, objc.Sel("unsignedLongValue"))
	return rv
}



































