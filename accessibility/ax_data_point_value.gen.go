// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AXDataPointValue] class.
var (
	_AXDataPointValueClass     AXDataPointValueClass
	_AXDataPointValueClassOnce sync.Once
)

func getAXDataPointValueClass() AXDataPointValueClass {
	_AXDataPointValueClassOnce.Do(func() {
		_AXDataPointValueClass = AXDataPointValueClass{class: objc.GetClass("AXDataPointValue")}
	})
	return _AXDataPointValueClass
}

// GetAXDataPointValueClass returns the class object for AXDataPointValue.
func GetAXDataPointValueClass() AXDataPointValueClass {
	return getAXDataPointValueClass()
}

type AXDataPointValueClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXDataPointValueClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXDataPointValueClass) Alloc() AXDataPointValue {
	rv := objc.Send[AXDataPointValue](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A single data value.
//
// # Overview
//
// An [AXDataPointValue] can be either numeric or categorical. Data points in
// a numeric axis use the [AXDataPointValue.Number] property, and data points
// in a categorical axis use the [AXDataPointValue.Category] property.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataPointValue
type AXDataPointValue struct {
	objectivec.Object
}

// AXDataPointValueFromID constructs a [AXDataPointValue] from an objc.ID.
//
// A single data value.
func AXDataPointValueFromID(id objc.ID) AXDataPointValue {
	return AXDataPointValue{objectivec.Object{ID: id}}
}

// NOTE: AXDataPointValue adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXDataPointValue] class.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataPointValue
type IAXDataPointValue interface {
	objectivec.IObject

	// A string that represents the categorical data value.
	Category() string
	SetCategory(value string)
	// A number that represents the numeric data value.
	Number() float64
	SetNumber(value float64)
}

// Init initializes the instance.
func (a AXDataPointValue) Init() AXDataPointValue {
	rv := objc.Send[AXDataPointValue](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXDataPointValue) Autorelease() AXDataPointValue {
	rv := objc.Send[AXDataPointValue](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXDataPointValue creates a new AXDataPointValue instance.
func NewAXDataPointValue() AXDataPointValue {
	class := getAXDataPointValueClass()
	rv := objc.Send[AXDataPointValue](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a categorical data value with the specified category string.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataPointValue/valueWithCategory:
func (_AXDataPointValueClass AXDataPointValueClass) ValueWithCategory(category string) AXDataPointValue {
	rv := objc.Send[objc.ID](objc.ID(_AXDataPointValueClass.class), objc.Sel("valueWithCategory:"), objc.String(category))
	return AXDataPointValueFromID(rv)
}

// Creates a numeric data value with the specified number.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataPointValue/valueWithNumber:
func (_AXDataPointValueClass AXDataPointValueClass) ValueWithNumber(number float64) AXDataPointValue {
	rv := objc.Send[objc.ID](objc.ID(_AXDataPointValueClass.class), objc.Sel("valueWithNumber:"), number)
	return AXDataPointValueFromID(rv)
}

// A string that represents the categorical data value.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataPointValue/category
func (a AXDataPointValue) Category() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("category"))
	return foundation.NSStringFromID(rv).String()
}
func (a AXDataPointValue) SetCategory(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setCategory:"), objc.String(value))
}

// A number that represents the numeric data value.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataPointValue/number
func (a AXDataPointValue) Number() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("number"))
	return rv
}
func (a AXDataPointValue) SetNumber(value float64) {
	objc.Send[struct{}](a.ID, objc.Sel("setNumber:"), value)
}
