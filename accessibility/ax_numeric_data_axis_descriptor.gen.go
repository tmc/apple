// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AXNumericDataAxisDescriptor] class.
var (
	_AXNumericDataAxisDescriptorClass     AXNumericDataAxisDescriptorClass
	_AXNumericDataAxisDescriptorClassOnce sync.Once
)

func getAXNumericDataAxisDescriptorClass() AXNumericDataAxisDescriptorClass {
	_AXNumericDataAxisDescriptorClassOnce.Do(func() {
		_AXNumericDataAxisDescriptorClass = AXNumericDataAxisDescriptorClass{class: objc.GetClass("AXNumericDataAxisDescriptor")}
	})
	return _AXNumericDataAxisDescriptorClass
}

// GetAXNumericDataAxisDescriptorClass returns the class object for AXNumericDataAxisDescriptor.
func GetAXNumericDataAxisDescriptorClass() AXNumericDataAxisDescriptorClass {
	return getAXNumericDataAxisDescriptorClass()
}

type AXNumericDataAxisDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXNumericDataAxisDescriptorClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXNumericDataAxisDescriptorClass) Alloc() AXNumericDataAxisDescriptor {
	rv := objc.Send[AXNumericDataAxisDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents an axis of numerical data.
//
// # Specifying the value description
//
//   - [AXNumericDataAxisDescriptor.ValueDescriptionProvider]: A description to speak for a particular data value on the axis.
//   - [AXNumericDataAxisDescriptor.SetValueDescriptionProvider]
//
// # Configuring the axis scale
//
//   - [AXNumericDataAxisDescriptor.ScaleType]: The scale for the axis.
//   - [AXNumericDataAxisDescriptor.SetScaleType]
//
// # Configuring the gridlines
//
//   - [AXNumericDataAxisDescriptor.GridlinePositions]: The positions of the gridlines along the axis.
//   - [AXNumericDataAxisDescriptor.SetGridlinePositions]
//
// See: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor
type AXNumericDataAxisDescriptor struct {
	objectivec.Object
}

// AXNumericDataAxisDescriptorFromID constructs a [AXNumericDataAxisDescriptor] from an objc.ID.
//
// An object that represents an axis of numerical data.
func AXNumericDataAxisDescriptorFromID(id objc.ID) AXNumericDataAxisDescriptor {
	return AXNumericDataAxisDescriptor{objectivec.Object{ID: id}}
}

// NOTE: AXNumericDataAxisDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXNumericDataAxisDescriptor] class.
//
// # Specifying the value description
//
//   - [IAXNumericDataAxisDescriptor.ValueDescriptionProvider]: A description to speak for a particular data value on the axis.
//   - [IAXNumericDataAxisDescriptor.SetValueDescriptionProvider]
//
// # Configuring the axis scale
//
//   - [IAXNumericDataAxisDescriptor.ScaleType]: The scale for the axis.
//   - [IAXNumericDataAxisDescriptor.SetScaleType]
//
// # Configuring the gridlines
//
//   - [IAXNumericDataAxisDescriptor.GridlinePositions]: The positions of the gridlines along the axis.
//   - [IAXNumericDataAxisDescriptor.SetGridlinePositions]
//
// See: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor
type IAXNumericDataAxisDescriptor interface {
	objectivec.IObject

	// Topic: Specifying the value description

	// A description to speak for a particular data value on the axis.
	ValueDescriptionProvider() StringFloat64Handler
	SetValueDescriptionProvider(value StringFloat64Handler)

	// Topic: Configuring the axis scale

	// The scale for the axis.
	ScaleType() AXNumericDataAxisDescriptorScale
	SetScaleType(value AXNumericDataAxisDescriptorScale)

	// Topic: Configuring the gridlines

	// The positions of the gridlines along the axis.
	GridlinePositions() []foundation.NSNumber
	SetGridlinePositions(value []foundation.NSNumber)

	// The minimum displayable value for the axis.
	LowerBound() float64
	SetLowerBound(value float64)
	// The maximum displayable value for the axis.
	UpperBound() float64
	SetUpperBound(value float64)
	// An attributed version of the axis title.
	AttributedTitle() foundation.NSAttributedString
	// Creates a numeric data axis with the specified attributed title, lower bound value, upper bound value, gridline positions, and value description provider block.
	InitWithAttributedTitleLowerBoundUpperBoundGridlinePositionsValueDescriptionProvider(attributedTitle foundation.NSAttributedString, lowerbound float64, upperBound float64, gridlinePositions []foundation.NSNumber, valueDescriptionProvider StringFloat64Handler) AXNumericDataAxisDescriptor
	// Creates a numeric data axis with the specified title, lower bound value, upper bound value, gridline positions, and value description provider block.
	InitWithTitleLowerBoundUpperBoundGridlinePositionsValueDescriptionProvider(title string, lowerbound float64, upperBound float64, gridlinePositions []foundation.NSNumber, valueDescriptionProvider StringFloat64Handler) AXNumericDataAxisDescriptor
	// The title of the axis.
	Title() string
}

// Init initializes the instance.
func (a AXNumericDataAxisDescriptor) Init() AXNumericDataAxisDescriptor {
	rv := objc.Send[AXNumericDataAxisDescriptor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXNumericDataAxisDescriptor) Autorelease() AXNumericDataAxisDescriptor {
	rv := objc.Send[AXNumericDataAxisDescriptor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXNumericDataAxisDescriptor creates a new AXNumericDataAxisDescriptor instance.
func NewAXNumericDataAxisDescriptor() AXNumericDataAxisDescriptor {
	class := getAXNumericDataAxisDescriptorClass()
	rv := objc.Send[AXNumericDataAxisDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An attributed version of the axis title.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataAxisDescriptor/attributedTitle
func (a AXNumericDataAxisDescriptor) AttributedTitle() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("attributedTitle"))
	return foundation.NSAttributedStringFromID(rv)
}

// Creates a numeric data axis with the specified attributed title, lower
// bound value, upper bound value, gridline positions, and value description
// provider block.
//
// See: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/initWithAttributedTitle:lowerBound:upperBound:gridlinePositions:valueDescriptionProvider:
func (a AXNumericDataAxisDescriptor) InitWithAttributedTitleLowerBoundUpperBoundGridlinePositionsValueDescriptionProvider(attributedTitle foundation.NSAttributedString, lowerbound float64, upperBound float64, gridlinePositions []foundation.NSNumber, valueDescriptionProvider StringFloat64Handler) AXNumericDataAxisDescriptor {
	_block4, _ := NewStringFloat64Block(valueDescriptionProvider)
	rv := objc.Send[AXNumericDataAxisDescriptor](a.ID, objc.Sel("initWithAttributedTitle:lowerBound:upperBound:gridlinePositions:valueDescriptionProvider:"), attributedTitle, lowerbound, upperBound, gridlinePositions, _block4)
	return rv
}

// Creates a numeric data axis with the specified title, lower bound value,
// upper bound value, gridline positions, and value description provider
// block.
//
// See: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/initWithTitle:lowerBound:upperBound:gridlinePositions:valueDescriptionProvider:
func (a AXNumericDataAxisDescriptor) InitWithTitleLowerBoundUpperBoundGridlinePositionsValueDescriptionProvider(title string, lowerbound float64, upperBound float64, gridlinePositions []foundation.NSNumber, valueDescriptionProvider StringFloat64Handler) AXNumericDataAxisDescriptor {
	_block4, _ := NewStringFloat64Block(valueDescriptionProvider)
	rv := objc.Send[AXNumericDataAxisDescriptor](a.ID, objc.Sel("initWithTitle:lowerBound:upperBound:gridlinePositions:valueDescriptionProvider:"), objc.String(title), lowerbound, upperBound, gridlinePositions, _block4)
	return rv
}

// The title of the axis.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataAxisDescriptor/title
func (a AXNumericDataAxisDescriptor) Title() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}

// A description to speak for a particular data value on the axis.
//
// # Discussion
//
// Use this property to format data values into string representations that
// include units, dates, times, and more.
//
// See: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/valueDescriptionProvider
func (a AXNumericDataAxisDescriptor) ValueDescriptionProvider() StringFloat64Handler {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("valueDescriptionProvider"))
	_ = rv
	return nil
}
func (a AXNumericDataAxisDescriptor) SetValueDescriptionProvider(value StringFloat64Handler) {
	block, cleanup := NewStringFloat64Block(value)
	defer cleanup()
	objc.Send[struct{}](a.ID, objc.Sel("setValueDescriptionProvider:"), block)
}

// The scale for the axis.
//
// # Discussion
//
// Match the value of this property to the visual representation in the chart.
//
// The default value is [AXScaleTypeLinear].
//
// See: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/scaleType-swift.property
func (a AXNumericDataAxisDescriptor) ScaleType() AXNumericDataAxisDescriptorScale {
	rv := objc.Send[AXNumericDataAxisDescriptorScale](a.ID, objc.Sel("scaleType"))
	return AXNumericDataAxisDescriptorScale(rv)
}
func (a AXNumericDataAxisDescriptor) SetScaleType(value AXNumericDataAxisDescriptorScale) {
	objc.Send[struct{}](a.ID, objc.Sel("setScaleType:"), value)
}

// The positions of the gridlines along the axis.
//
// See: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/gridlinePositions-9z10e
func (a AXNumericDataAxisDescriptor) GridlinePositions() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("gridlinePositions"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
func (a AXNumericDataAxisDescriptor) SetGridlinePositions(value []foundation.NSNumber) {
	objc.Send[struct{}](a.ID, objc.Sel("setGridlinePositions:"), objectivec.IObjectSliceToNSArray(value))
}

// The minimum displayable value for the axis.
//
// See: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/lowerBound
func (a AXNumericDataAxisDescriptor) LowerBound() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("lowerBound"))
	return rv
}
func (a AXNumericDataAxisDescriptor) SetLowerBound(value float64) {
	objc.Send[struct{}](a.ID, objc.Sel("setLowerBound:"), value)
}

// The maximum displayable value for the axis.
//
// See: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/upperBound
func (a AXNumericDataAxisDescriptor) UpperBound() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("upperBound"))
	return rv
}
func (a AXNumericDataAxisDescriptor) SetUpperBound(value float64) {
	objc.Send[struct{}](a.ID, objc.Sel("setUpperBound:"), value)
}

// Protocol methods for AXDataAxisDescriptor

// The title of the axis.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataAxisDescriptor/title
func (o AXNumericDataAxisDescriptor) SetTitle(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setTitle:"), objc.String(value))
}

// An attributed version of the axis title.
//
// # Discussion
//
// If you set the value of this property, the system uses this value instead
// of [Title].
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataAxisDescriptor/attributedTitle
func (o AXNumericDataAxisDescriptor) SetAttributedTitle(value foundation.NSAttributedString) {
	objc.Send[struct{}](o.ID, objc.Sel("setAttributedTitle:"), value)
}
