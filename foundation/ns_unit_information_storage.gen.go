// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [UnitInformationStorage] class.
var (
	_UnitInformationStorageClass     UnitInformationStorageClass
	_UnitInformationStorageClassOnce sync.Once
)

func getUnitInformationStorageClass() UnitInformationStorageClass {
	_UnitInformationStorageClassOnce.Do(func() {
		_UnitInformationStorageClass = UnitInformationStorageClass{class: objc.GetClass("NSUnitInformationStorage")}
	})
	return _UnitInformationStorageClass
}

// GetUnitInformationStorageClass returns the class object for NSUnitInformationStorage.
func GetUnitInformationStorageClass() UnitInformationStorageClass {
	return getUnitInformationStorageClass()
}

type UnitInformationStorageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UnitInformationStorageClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitInformationStorageClass) Alloc() UnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A unit of measure for quantities of information.
//
// # Overview
//
// Use instances of [NSUnitInformationStorage] to represent quantities of
// information using the [NSMeasurement] class. The base unit of measure for
// information is the bit, with a nibble representing four bits and a byte
// representing eight bits.
//
// Larger units of information expand on bits and bytes by orders of magnitude
// in both decimal and binary forms.
//
// # Information Transfer
//
// Units of bits commonly represent the amount of transferred information.
//
// [Table data omitted]
//
// # Information Storage
//
// Units of bytes commonly represent the amount of stored information.
//
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage
type UnitInformationStorage struct {
	NSDimension
}

// UnitInformationStorageFromID constructs a [UnitInformationStorage] from an objc.ID.
//
// A unit of measure for quantities of information.
func UnitInformationStorageFromID(id objc.ID) UnitInformationStorage {
	return NSUnitInformationStorage{NSDimension: NSDimensionFromID(id)}
}

// NSUnitInformationStorageFromID is an alias for [UnitInformationStorageFromID] for cross-framework compatibility.
func NSUnitInformationStorageFromID(id objc.ID) UnitInformationStorage {
	return UnitInformationStorageFromID(id)
}

// NOTE: UnitInformationStorage adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UnitInformationStorage] class.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage
type IUnitInformationStorage interface {
	INSDimension
}

// Init initializes the instance.
func (u UnitInformationStorage) Init() UnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UnitInformationStorage) Autorelease() UnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitInformationStorage creates a new UnitInformationStorage instance.
func NewUnitInformationStorage() UnitInformationStorage {
	class := getUnitInformationStorageClass()
	rv := objc.Send[UnitInformationStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUnitInformationStorageWithCoder(coder INSCoder) UnitInformationStorage {
	instance := getUnitInformationStorageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return UnitInformationStorageFromID(rv)
}

// Initializes a new unit with the specified symbol.
//
// symbol: The symbol used to represent the unit.
//
// # Return Value
//
// A new unit with the specified symbol.
//
// See: https://developer.apple.com/documentation/Foundation/Unit/init(symbol:)
func NewUnitInformationStorageWithSymbol(symbol string) UnitInformationStorage {
	instance := getUnitInformationStorageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return UnitInformationStorageFromID(rv)
}

// Initializes a dimensional unit with the symbol and unit converter you
// specify.
//
// symbol: The symbol used to represent the unit.
//
// converter: The unit converter used to represent the unit in terms of the dimension’s
// base unit.
//
// # Return Value
//
// A new dimensional unit with the specified symbol and unit converter.
//
// # Discussion
//
// This is the designated initializer.
//
// See: https://developer.apple.com/documentation/Foundation/Dimension/init(symbol:converter:)
func NewUnitInformationStorageWithSymbolConverter(symbol string, converter INSUnitConverter) UnitInformationStorage {
	instance := getUnitInformationStorageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	return UnitInformationStorageFromID(rv)
}

// The bits unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/bits
func (_UnitInformationStorageClass UnitInformationStorageClass) Bits() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("bits"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The nibbles unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/nibbles
func (_UnitInformationStorageClass UnitInformationStorageClass) Nibbles() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("nibbles"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The bytes unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/bytes
func (_UnitInformationStorageClass UnitInformationStorageClass) Bytes() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("bytes"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The kibibits unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/kibibits
func (_UnitInformationStorageClass UnitInformationStorageClass) Kibibits() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("kibibits"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The kibibytes unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/kibibytes
func (_UnitInformationStorageClass UnitInformationStorageClass) Kibibytes() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("kibibytes"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The mebibits unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/mebibits
func (_UnitInformationStorageClass UnitInformationStorageClass) Mebibits() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("mebibits"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The mebibytes unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/mebibytes
func (_UnitInformationStorageClass UnitInformationStorageClass) Mebibytes() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("mebibytes"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The gibibits unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/gibibits
func (_UnitInformationStorageClass UnitInformationStorageClass) Gibibits() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("gibibits"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The gibibytes unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/gibibytes
func (_UnitInformationStorageClass UnitInformationStorageClass) Gibibytes() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("gibibytes"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The tebibits unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/tebibits
func (_UnitInformationStorageClass UnitInformationStorageClass) Tebibits() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("tebibits"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The tebibytes unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/tebibytes
func (_UnitInformationStorageClass UnitInformationStorageClass) Tebibytes() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("tebibytes"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The pebibits unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/pebibits
func (_UnitInformationStorageClass UnitInformationStorageClass) Pebibits() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("pebibits"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The pebibytes unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/pebibytes
func (_UnitInformationStorageClass UnitInformationStorageClass) Pebibytes() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("pebibytes"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The exbibits unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/exbibits
func (_UnitInformationStorageClass UnitInformationStorageClass) Exbibits() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("exbibits"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The exbibytes unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/exbibytes
func (_UnitInformationStorageClass UnitInformationStorageClass) Exbibytes() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("exbibytes"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The zebibits unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/zebibits
func (_UnitInformationStorageClass UnitInformationStorageClass) Zebibits() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("zebibits"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The zebibytes unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/zebibytes
func (_UnitInformationStorageClass UnitInformationStorageClass) Zebibytes() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("zebibytes"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The yobibits unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/yobibits
func (_UnitInformationStorageClass UnitInformationStorageClass) Yobibits() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("yobibits"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The yobibytes unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/yobibytes
func (_UnitInformationStorageClass UnitInformationStorageClass) Yobibytes() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("yobibytes"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The kilobits unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/kilobits
func (_UnitInformationStorageClass UnitInformationStorageClass) Kilobits() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("kilobits"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The kilobytes unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/kilobytes
func (_UnitInformationStorageClass UnitInformationStorageClass) Kilobytes() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("kilobytes"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The megabits unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/megabits
func (_UnitInformationStorageClass UnitInformationStorageClass) Megabits() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("megabits"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The megabytes unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/megabytes
func (_UnitInformationStorageClass UnitInformationStorageClass) Megabytes() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("megabytes"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The gigabits unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/gigabits
func (_UnitInformationStorageClass UnitInformationStorageClass) Gigabits() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("gigabits"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The gigabytes unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/gigabytes
func (_UnitInformationStorageClass UnitInformationStorageClass) Gigabytes() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("gigabytes"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The terabits unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/terabits
func (_UnitInformationStorageClass UnitInformationStorageClass) Terabits() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("terabits"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The terrabytes unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/terabytes
func (_UnitInformationStorageClass UnitInformationStorageClass) Terabytes() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("terabytes"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The petabits unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/petabits
func (_UnitInformationStorageClass UnitInformationStorageClass) Petabits() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("petabits"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The petabytes unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/petabytes
func (_UnitInformationStorageClass UnitInformationStorageClass) Petabytes() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("petabytes"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The exabits unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/exabits
func (_UnitInformationStorageClass UnitInformationStorageClass) Exabits() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("exabits"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The exabytes unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/exabytes
func (_UnitInformationStorageClass UnitInformationStorageClass) Exabytes() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("exabytes"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The zettabits unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/zettabits
func (_UnitInformationStorageClass UnitInformationStorageClass) Zettabits() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("zettabits"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The zettabytes unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/zettabytes
func (_UnitInformationStorageClass UnitInformationStorageClass) Zettabytes() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("zettabytes"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The yottabits unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/yottabits
func (_UnitInformationStorageClass UnitInformationStorageClass) Yottabits() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("yottabits"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}

// The yottabytes unit of information.
//
// See: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/yottabytes
func (_UnitInformationStorageClass UnitInformationStorageClass) Yottabytes() UnitInformationStorage {
	rv := objc.Send[objc.ID](objc.ID(_UnitInformationStorageClass.class), objc.Sel("yottabytes"))
	return NSUnitInformationStorageFromID(objc.ID(rv))
}
