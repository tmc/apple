// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AXDataPoint] class.
var (
	_AXDataPointClass     AXDataPointClass
	_AXDataPointClassOnce sync.Once
)

func getAXDataPointClass() AXDataPointClass {
	_AXDataPointClassOnce.Do(func() {
		_AXDataPointClass = AXDataPointClass{class: objc.GetClass("AXDataPoint")}
	})
	return _AXDataPointClass
}

// GetAXDataPointClass returns the class object for AXDataPoint.
func GetAXDataPointClass() AXDataPointClass {
	return getAXDataPointClass()
}

type AXDataPointClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXDataPointClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXDataPointClass) Alloc() AXDataPoint {
	rv := objc.Send[AXDataPoint](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a single data point in a chart.
//
// # Specifying the data value
//
//   - [AXDataPoint.XValue]: The value of the x-axis for the data point.
//   - [AXDataPoint.SetXValue]
//   - [AXDataPoint.YValue]: The value of the y-axis for the data point.
//   - [AXDataPoint.SetYValue]
//
// # Specifying the label
//
//   - [AXDataPoint.Label]: The label for the data point.
//   - [AXDataPoint.SetLabel]
//   - [AXDataPoint.AttributedLabel]: An attributed version of the label for the data point.
//   - [AXDataPoint.SetAttributedLabel]
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataPoint
type AXDataPoint struct {
	objectivec.Object
}

// AXDataPointFromID constructs a [AXDataPoint] from an objc.ID.
//
// An object that represents a single data point in a chart.
func AXDataPointFromID(id objc.ID) AXDataPoint {
	return AXDataPoint{objectivec.Object{ID: id}}
}

// NOTE: AXDataPoint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXDataPoint] class.
//
// # Specifying the data value
//
//   - [IAXDataPoint.XValue]: The value of the x-axis for the data point.
//   - [IAXDataPoint.SetXValue]
//   - [IAXDataPoint.YValue]: The value of the y-axis for the data point.
//   - [IAXDataPoint.SetYValue]
//
// # Specifying the label
//
//   - [IAXDataPoint.Label]: The label for the data point.
//   - [IAXDataPoint.SetLabel]
//   - [IAXDataPoint.AttributedLabel]: An attributed version of the label for the data point.
//   - [IAXDataPoint.SetAttributedLabel]
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataPoint
type IAXDataPoint interface {
	objectivec.IObject

	// Topic: Specifying the data value

	// The value of the x-axis for the data point.
	XValue() IAXDataPointValue
	SetXValue(value IAXDataPointValue)
	// The value of the y-axis for the data point.
	YValue() IAXDataPointValue
	SetYValue(value IAXDataPointValue)

	// Topic: Specifying the label

	// The label for the data point.
	Label() string
	SetLabel(value string)
	// An attributed version of the label for the data point.
	AttributedLabel() foundation.NSAttributedString
	SetAttributedLabel(value foundation.NSAttributedString)

	// An array of values for additional axes for the data point.
	AdditionalValues() []AXDataPointValue
	SetAdditionalValues(value []AXDataPointValue)
	// Creates a data point with the specified x- and y-values.
	InitWithXY(xValue IAXDataPointValue, yValue IAXDataPointValue) AXDataPoint
	// Creates a data point with the specified x-value, y-value, and additional values.
	InitWithXYAdditionalValues(xValue IAXDataPointValue, yValue IAXDataPointValue, additionalValues []AXDataPointValue) AXDataPoint
	// Creates a data point with the specified x-value, y-value, additional values, and label.
	InitWithXYAdditionalValuesLabel(xValue IAXDataPointValue, yValue IAXDataPointValue, additionalValues []AXDataPointValue, label string) AXDataPoint
}

// Init initializes the instance.
func (a AXDataPoint) Init() AXDataPoint {
	rv := objc.Send[AXDataPoint](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXDataPoint) Autorelease() AXDataPoint {
	rv := objc.Send[AXDataPoint](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXDataPoint creates a new AXDataPoint instance.
func NewAXDataPoint() AXDataPoint {
	class := getAXDataPointClass()
	rv := objc.Send[AXDataPoint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a data point with the specified x- and y-values.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataPoint/initWithX:y:
func NewAXDataPointWithXY(xValue IAXDataPointValue, yValue IAXDataPointValue) AXDataPoint {
	instance := getAXDataPointClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithX:y:"), xValue, yValue)
	return AXDataPointFromID(rv)
}

// Creates a data point with the specified x-value, y-value, and additional
// values.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataPoint/initWithX:y:additionalValues:
func NewAXDataPointWithXYAdditionalValues(xValue IAXDataPointValue, yValue IAXDataPointValue, additionalValues []AXDataPointValue) AXDataPoint {
	instance := getAXDataPointClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithX:y:additionalValues:"), xValue, yValue, objectivec.IObjectSliceToNSArray(additionalValues))
	return AXDataPointFromID(rv)
}

// Creates a data point with the specified x-value, y-value, additional
// values, and label.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataPoint/initWithX:y:additionalValues:label:
func NewAXDataPointWithXYAdditionalValuesLabel(xValue IAXDataPointValue, yValue IAXDataPointValue, additionalValues []AXDataPointValue, label string) AXDataPoint {
	instance := getAXDataPointClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithX:y:additionalValues:label:"), xValue, yValue, objectivec.IObjectSliceToNSArray(additionalValues), objc.String(label))
	return AXDataPointFromID(rv)
}

// Creates a data point with the specified x- and y-values.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataPoint/initWithX:y:
func (a AXDataPoint) InitWithXY(xValue IAXDataPointValue, yValue IAXDataPointValue) AXDataPoint {
	rv := objc.Send[AXDataPoint](a.ID, objc.Sel("initWithX:y:"), xValue, yValue)
	return rv
}

// Creates a data point with the specified x-value, y-value, and additional
// values.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataPoint/initWithX:y:additionalValues:
func (a AXDataPoint) InitWithXYAdditionalValues(xValue IAXDataPointValue, yValue IAXDataPointValue, additionalValues []AXDataPointValue) AXDataPoint {
	rv := objc.Send[AXDataPoint](a.ID, objc.Sel("initWithX:y:additionalValues:"), xValue, yValue, objectivec.IObjectSliceToNSArray(additionalValues))
	return rv
}

// Creates a data point with the specified x-value, y-value, additional
// values, and label.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataPoint/initWithX:y:additionalValues:label:
func (a AXDataPoint) InitWithXYAdditionalValuesLabel(xValue IAXDataPointValue, yValue IAXDataPointValue, additionalValues []AXDataPointValue, label string) AXDataPoint {
	rv := objc.Send[AXDataPoint](a.ID, objc.Sel("initWithX:y:additionalValues:label:"), xValue, yValue, objectivec.IObjectSliceToNSArray(additionalValues), objc.String(label))
	return rv
}

// The value of the x-axis for the data point.
//
// # Discussion
//
// Use a `double` value for a numeric x-axis, or an [NSString] value for a
// categorical x-axis.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataPoint/xValue
//
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
func (a AXDataPoint) XValue() IAXDataPointValue {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("xValue"))
	return AXDataPointValueFromID(objc.ID(rv))
}
func (a AXDataPoint) SetXValue(value IAXDataPointValue) {
	objc.Send[struct{}](a.ID, objc.Sel("setXValue:"), value)
}

// The value of the y-axis for the data point.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataPoint/yValue
func (a AXDataPoint) YValue() IAXDataPointValue {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("yValue"))
	return AXDataPointValueFromID(objc.ID(rv))
}
func (a AXDataPoint) SetYValue(value IAXDataPointValue) {
	objc.Send[struct{}](a.ID, objc.Sel("setYValue:"), value)
}

// The label for the data point.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataPoint/label
func (a AXDataPoint) Label() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (a AXDataPoint) SetLabel(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setLabel:"), objc.String(value))
}

// An attributed version of the label for the data point.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataPoint/attributedLabel
func (a AXDataPoint) AttributedLabel() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("attributedLabel"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (a AXDataPoint) SetAttributedLabel(value foundation.NSAttributedString) {
	objc.Send[struct{}](a.ID, objc.Sel("setAttributedLabel:"), value)
}

// An array of values for additional axes for the data point.
//
// # Discussion
//
// Provide these values in the same order as their corresponding
// [AXDataAxisDescriptor] objects in [additionalAxes].
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataPoint/additionalValues
//
// [additionalAxes]: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/additionalAxes-2adwh
func (a AXDataPoint) AdditionalValues() []AXDataPointValue {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("additionalValues"))
	return objc.ConvertSlice(rv, func(id objc.ID) AXDataPointValue {
		return AXDataPointValueFromID(id)
	})
}
func (a AXDataPoint) SetAdditionalValues(value []AXDataPointValue) {
	objc.Send[struct{}](a.ID, objc.Sel("setAdditionalValues:"), objectivec.IObjectSliceToNSArray(value))
}
