// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMeasurement] class.
var (
	_NSMeasurementClass     NSMeasurementClass
	_NSMeasurementClassOnce sync.Once
)

func getNSMeasurementClass() NSMeasurementClass {
	_NSMeasurementClassOnce.Do(func() {
		_NSMeasurementClass = NSMeasurementClass{class: objc.GetClass("NSMeasurement")}
	})
	return _NSMeasurementClass
}

// GetNSMeasurementClass returns the class object for NSMeasurement.
func GetNSMeasurementClass() NSMeasurementClass {
	return getNSMeasurementClass()
}

type NSMeasurementClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMeasurementClass) Alloc() NSMeasurement {
	rv := objc.Send[NSMeasurement](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A numeric quantity labeled with a unit of measure, with support for unit
// conversion and unit-aware calculations.
//
// # Overview
// 
// Use this object in Swift when you need reference semantics or other
// Foundation-specific behavior.
// 
// An [NSMeasurement] object represents a quantity and unit of measure. The
// [NSMeasurement] class provides a programmatic interface to converting
// measurements into different units, as well as calculating the sum or
// difference between two measurements.
// 
// [NSMeasurement] objects are initialized with an [NSUnit] object and
// `double` value. [NSMeasurement] objects are immutable, and cannot be
// changed after being created.
// 
// You can use the [NSMeasurementFormatter] class to create localized string
// representations of [NSMeasurement] objects.
//
// # Creating Measurements
//
//   - [NSMeasurement.InitWithDoubleValueUnit]: Initializes a new measurement with a specified double-precision floating-point value and unit.
//
// # Accessing Unit and Value
//
//   - [NSMeasurement.Unit]: The unit of measure.
//   - [NSMeasurement.DoubleValue]: The measurement value, represented as a double-precision floating-point number.
//
// # Converting to Other Units
//
//   - [NSMeasurement.CanBeConvertedToUnit]: Indicates whether the measurement can be converted to the given unit.
//   - [NSMeasurement.MeasurementByConvertingToUnit]: Returns a measurement created by converting the receiver to the specified unit.
//
// # Operating on Measurements
//
//   - [NSMeasurement.MeasurementByAddingMeasurement]: Returns a new measurement by adding the receiver to the specified measurement.
//   - [NSMeasurement.MeasurementBySubtractingMeasurement]: Returns a new measurement by subtracting the specified measurement from the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSMeasurement
type NSMeasurement struct {
	objectivec.Object
}

// NSMeasurementFromID constructs a [NSMeasurement] from an objc.ID.
//
// A numeric quantity labeled with a unit of measure, with support for unit
// conversion and unit-aware calculations.
func NSMeasurementFromID(id objc.ID) NSMeasurement {
	return NSMeasurement{objectivec.Object{ID: id}}
}
// NOTE: NSMeasurement adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMeasurement] class.
//
// # Creating Measurements
//
//   - [INSMeasurement.InitWithDoubleValueUnit]: Initializes a new measurement with a specified double-precision floating-point value and unit.
//
// # Accessing Unit and Value
//
//   - [INSMeasurement.Unit]: The unit of measure.
//   - [INSMeasurement.DoubleValue]: The measurement value, represented as a double-precision floating-point number.
//
// # Converting to Other Units
//
//   - [INSMeasurement.CanBeConvertedToUnit]: Indicates whether the measurement can be converted to the given unit.
//   - [INSMeasurement.MeasurementByConvertingToUnit]: Returns a measurement created by converting the receiver to the specified unit.
//
// # Operating on Measurements
//
//   - [INSMeasurement.MeasurementByAddingMeasurement]: Returns a new measurement by adding the receiver to the specified measurement.
//   - [INSMeasurement.MeasurementBySubtractingMeasurement]: Returns a new measurement by subtracting the specified measurement from the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSMeasurement
type INSMeasurement interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Creating Measurements

	// Initializes a new measurement with a specified double-precision floating-point value and unit.
	InitWithDoubleValueUnit(doubleValue float64, unit objectivec.IObject) NSMeasurement

	// Topic: Accessing Unit and Value

	// The unit of measure.
	Unit() objectivec.IObject
	// The measurement value, represented as a double-precision floating-point number.
	DoubleValue() float64

	// Topic: Converting to Other Units

	// Indicates whether the measurement can be converted to the given unit.
	CanBeConvertedToUnit(unit INSUnit) bool
	// Returns a measurement created by converting the receiver to the specified unit.
	MeasurementByConvertingToUnit(unit INSUnit) INSMeasurement

	// Topic: Operating on Measurements

	// Returns a new measurement by adding the receiver to the specified measurement.
	MeasurementByAddingMeasurement(measurement INSMeasurement) INSMeasurement
	// Returns a new measurement by subtracting the specified measurement from the receiver.
	MeasurementBySubtractingMeasurement(measurement INSMeasurement) INSMeasurement
}

// Init initializes the instance.
func (m NSMeasurement) Init() NSMeasurement {
	rv := objc.Send[NSMeasurement](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMeasurement) Autorelease() NSMeasurement {
	rv := objc.Send[NSMeasurement](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMeasurement creates a new NSMeasurement instance.
func NewNSMeasurement() NSMeasurement {
	class := getNSMeasurementClass()
	rv := objc.Send[NSMeasurement](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewMeasurementWithCoder(coder INSCoder) NSMeasurement {
	instance := getNSMeasurementClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSMeasurementFromID(rv)
}

// Initializes a new measurement with a specified double-precision
// floating-point value and unit.
//
// doubleValue: The double-precision floating-point measurement value.
//
// unit: The unit of measure.
//
// # Return Value
// 
// A measurement initialized to have the specified double-precision
// floating-point value and unit.
//
// See: https://developer.apple.com/documentation/Foundation/NSMeasurement/init(doubleValue:unit:)
func NewMeasurementWithDoubleValueUnit(doubleValue float64, unit objectivec.IObject) NSMeasurement {
	instance := getNSMeasurementClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDoubleValue:unit:"), doubleValue, unit)
	return NSMeasurementFromID(rv)
}

// Initializes a new measurement with a specified double-precision
// floating-point value and unit.
//
// doubleValue: The double-precision floating-point measurement value.
//
// unit: The unit of measure.
//
// # Return Value
// 
// A measurement initialized to have the specified double-precision
// floating-point value and unit.
//
// See: https://developer.apple.com/documentation/Foundation/NSMeasurement/init(doubleValue:unit:)
func (m NSMeasurement) InitWithDoubleValueUnit(doubleValue float64, unit objectivec.IObject) NSMeasurement {
	rv := objc.Send[NSMeasurement](m.ID, objc.Sel("initWithDoubleValue:unit:"), doubleValue, unit)
	return rv
}

// Indicates whether the measurement can be converted to the given unit.
//
// See: https://developer.apple.com/documentation/Foundation/NSMeasurement/canBeConverted(to:)
func (m NSMeasurement) CanBeConvertedToUnit(unit INSUnit) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("canBeConvertedToUnit:"), unit)
	return rv
}

// Returns a measurement created by converting the receiver to the specified
// unit.
//
// unit: The unit to convert the measurement into.
//
// # Return Value
// 
// A new measurement with a value calculated by converting into the new unit.
//
// # Discussion
// 
// This method raises an [invalidArgumentException] if the receiver cannot be
// converted to unit.
//
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
//
// See: https://developer.apple.com/documentation/Foundation/NSMeasurement/converting(to:)
func (m NSMeasurement) MeasurementByConvertingToUnit(unit INSUnit) INSMeasurement {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("measurementByConvertingToUnit:"), unit)
	return NSMeasurementFromID(rv)
}

// Returns a new measurement by adding the receiver to the specified
// measurement.
//
// measurement: The measurement to be added.
//
// # Return Value
// 
// A new measurement with a value equal to the receiver’s value plus the
// value of the specified measurement converted into the unit of the receiver.
//
// # Discussion
// 
// This method raises an [invalidArgumentException] if the receiver cannot be
// converted to unit.
// 
// You can use the [CanBeConvertedToUnit] method, passing the unit of the
// specified measurement, to determine whether a measurement can be converted
// to a particular unit before calling this method.
//
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
//
// See: https://developer.apple.com/documentation/Foundation/NSMeasurement/adding(_:)
func (m NSMeasurement) MeasurementByAddingMeasurement(measurement INSMeasurement) INSMeasurement {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("measurementByAddingMeasurement:"), measurement)
	return NSMeasurementFromID(rv)
}

// Returns a new measurement by subtracting the specified measurement from the
// receiver.
//
// measurement: The measurement to be subtracted.
//
// # Return Value
// 
// A new measurement with a value equal to the receiver’s value minus the
// value of the specified measurement converted into the unit of the receiver.
//
// # Discussion
// 
// This method raises an [invalidArgumentException] if the receiver cannot be
// converted to unit.
// 
// You can use the [CanBeConvertedToUnit] method, passing the unit of the
// specified measurement, to determine whether a measurement can be converted
// to a particular unit before calling this method.
//
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
//
// See: https://developer.apple.com/documentation/Foundation/NSMeasurement/subtracting(_:)
func (m NSMeasurement) MeasurementBySubtractingMeasurement(measurement INSMeasurement) INSMeasurement {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("measurementBySubtractingMeasurement:"), measurement)
	return NSMeasurementFromID(rv)
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (m NSMeasurement) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (m NSMeasurement) InitWithCoder(coder INSCoder) NSMeasurement {
	rv := objc.Send[NSMeasurement](m.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// The unit of measure.
//
// See: https://developer.apple.com/documentation/Foundation/NSMeasurement/unit
func (m NSMeasurement) Unit() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("unit"))
	return objectivec.Object{ID: rv}
}

// The measurement value, represented as a double-precision floating-point
// number.
//
// See: https://developer.apple.com/documentation/Foundation/NSMeasurement/doubleValue
func (m NSMeasurement) DoubleValue() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("doubleValue"))
	return rv
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

