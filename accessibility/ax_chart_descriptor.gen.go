// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AXChartDescriptor] class.
var (
	_AXChartDescriptorClass     AXChartDescriptorClass
	_AXChartDescriptorClassOnce sync.Once
)

func getAXChartDescriptorClass() AXChartDescriptorClass {
	_AXChartDescriptorClassOnce.Do(func() {
		_AXChartDescriptorClass = AXChartDescriptorClass{class: objc.GetClass("AXChartDescriptor")}
	})
	return _AXChartDescriptorClass
}

// GetAXChartDescriptorClass returns the class object for AXChartDescriptor.
func GetAXChartDescriptorClass() AXChartDescriptorClass {
	return getAXChartDescriptorClass()
}

type AXChartDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXChartDescriptorClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXChartDescriptorClass) Alloc() AXChartDescriptor {
	rv := objc.Send[AXChartDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that contains all the semantic information about an accessible
// chart.
//
// # Specifying the chart title
//
//   - [AXChartDescriptor.Title]: The title of the chart.
//   - [AXChartDescriptor.SetTitle]
//   - [AXChartDescriptor.AttributedTitle]: An attributed version of the chart title.
//   - [AXChartDescriptor.SetAttributedTitle]
//
// # Specifying the chart summary
//
//   - [AXChartDescriptor.Summary]: A description of the key takeaways or features of the chart.
//   - [AXChartDescriptor.SetSummary]
//
// # Specifying the axes
//
//   - [AXChartDescriptor.XAxis]: The axis descriptor for the chart’s x-axis.
//   - [AXChartDescriptor.SetXAxis]
//   - [AXChartDescriptor.YAxis]: The axis descriptor for the chart’s y-axis.
//   - [AXChartDescriptor.SetYAxis]
//   - [AXChartDescriptor.AdditionalAxes]: The descriptors for additional categorical or numerical axes beyond the x-axis and y-axis.
//   - [AXChartDescriptor.SetAdditionalAxes]
//
// # Specifying a series of data points
//
//   - [AXChartDescriptor.Series]: The descriptors for each data series in the chart.
//   - [AXChartDescriptor.SetSeries]
//
// # Specifying the content layout
//
//   - [AXChartDescriptor.ContentFrame]: The bounds of the view, in screen coordinates, for visually rendering data values.
//   - [AXChartDescriptor.SetContentFrame]
//   - [AXChartDescriptor.ContentDirection]: The direction of the content in the chart.
//   - [AXChartDescriptor.SetContentDirection]
//
// See: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor
type AXChartDescriptor struct {
	objectivec.Object
}

// AXChartDescriptorFromID constructs a [AXChartDescriptor] from an objc.ID.
//
// An object that contains all the semantic information about an accessible
// chart.
func AXChartDescriptorFromID(id objc.ID) AXChartDescriptor {
	return AXChartDescriptor{objectivec.Object{ID: id}}
}

// NOTE: AXChartDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXChartDescriptor] class.
//
// # Specifying the chart title
//
//   - [IAXChartDescriptor.Title]: The title of the chart.
//   - [IAXChartDescriptor.SetTitle]
//   - [IAXChartDescriptor.AttributedTitle]: An attributed version of the chart title.
//   - [IAXChartDescriptor.SetAttributedTitle]
//
// # Specifying the chart summary
//
//   - [IAXChartDescriptor.Summary]: A description of the key takeaways or features of the chart.
//   - [IAXChartDescriptor.SetSummary]
//
// # Specifying the axes
//
//   - [IAXChartDescriptor.XAxis]: The axis descriptor for the chart’s x-axis.
//   - [IAXChartDescriptor.SetXAxis]
//   - [IAXChartDescriptor.YAxis]: The axis descriptor for the chart’s y-axis.
//   - [IAXChartDescriptor.SetYAxis]
//   - [IAXChartDescriptor.AdditionalAxes]: The descriptors for additional categorical or numerical axes beyond the x-axis and y-axis.
//   - [IAXChartDescriptor.SetAdditionalAxes]
//
// # Specifying a series of data points
//
//   - [IAXChartDescriptor.Series]: The descriptors for each data series in the chart.
//   - [IAXChartDescriptor.SetSeries]
//
// # Specifying the content layout
//
//   - [IAXChartDescriptor.ContentFrame]: The bounds of the view, in screen coordinates, for visually rendering data values.
//   - [IAXChartDescriptor.SetContentFrame]
//   - [IAXChartDescriptor.ContentDirection]: The direction of the content in the chart.
//   - [IAXChartDescriptor.SetContentDirection]
//
// See: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor
type IAXChartDescriptor interface {
	objectivec.IObject

	// Topic: Specifying the chart title

	// The title of the chart.
	Title() string
	SetTitle(value string)
	// An attributed version of the chart title.
	AttributedTitle() foundation.NSAttributedString
	SetAttributedTitle(value foundation.NSAttributedString)

	// Topic: Specifying the chart summary

	// A description of the key takeaways or features of the chart.
	Summary() string
	SetSummary(value string)

	// Topic: Specifying the axes

	// The axis descriptor for the chart’s x-axis.
	XAxis() AXDataAxisDescriptor
	SetXAxis(value AXDataAxisDescriptor)
	// The axis descriptor for the chart’s y-axis.
	YAxis() IAXNumericDataAxisDescriptor
	SetYAxis(value IAXNumericDataAxisDescriptor)
	// The descriptors for additional categorical or numerical axes beyond the x-axis and y-axis.
	AdditionalAxes() []objectivec.IObject
	SetAdditionalAxes(value []objectivec.IObject)

	// Topic: Specifying a series of data points

	// The descriptors for each data series in the chart.
	Series() []AXDataSeriesDescriptor
	SetSeries(value []AXDataSeriesDescriptor)

	// Topic: Specifying the content layout

	// The bounds of the view, in screen coordinates, for visually rendering data values.
	ContentFrame() corefoundation.CGRect
	SetContentFrame(value corefoundation.CGRect)
	// The direction of the content in the chart.
	ContentDirection() AXChartDescriptorContentDirection
	SetContentDirection(value AXChartDescriptorContentDirection)

	// Creates a chart descriptor with the specified attributed title, summary, x-axis descriptor, y-axis descriptor, descriptors for additional axes, and array of data series.
	InitWithAttributedTitleSummaryXAxisDescriptorYAxisDescriptorAdditionalAxesSeries(attributedTitle foundation.NSAttributedString, summary string, xAxis AXDataAxisDescriptor, yAxis IAXNumericDataAxisDescriptor, additionalAxes []objectivec.IObject, series []AXDataSeriesDescriptor) AXChartDescriptor
	// Creates a chart descriptor with the specified attributed title, summary, x-axis descriptor, y-axis descriptor, and array of data series.
	InitWithAttributedTitleSummaryXAxisDescriptorYAxisDescriptorSeries(attributedTitle foundation.NSAttributedString, summary string, xAxis AXDataAxisDescriptor, yAxis IAXNumericDataAxisDescriptor, series []AXDataSeriesDescriptor) AXChartDescriptor
	// Creates a chart descriptor with the specified title, summary, x-axis descriptor, y-axis descriptor, descriptors for additional axes, and array of data series.
	InitWithTitleSummaryXAxisDescriptorYAxisDescriptorAdditionalAxesSeries(title string, summary string, xAxis AXDataAxisDescriptor, yAxis IAXNumericDataAxisDescriptor, additionalAxes []objectivec.IObject, series []AXDataSeriesDescriptor) AXChartDescriptor
	// Creates a chart descriptor with the specified title, summary, x-axis descriptor, y-axis descriptor, descriptors for additional axes, and array of data series.
	InitWithTitleSummaryXAxisDescriptorYAxisDescriptorSeries(title string, summary string, xAxis AXDataAxisDescriptor, yAxis IAXNumericDataAxisDescriptor, series []AXDataSeriesDescriptor) AXChartDescriptor
}

// Init initializes the instance.
func (a AXChartDescriptor) Init() AXChartDescriptor {
	rv := objc.Send[AXChartDescriptor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXChartDescriptor) Autorelease() AXChartDescriptor {
	rv := objc.Send[AXChartDescriptor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXChartDescriptor creates a new AXChartDescriptor instance.
func NewAXChartDescriptor() AXChartDescriptor {
	class := getAXChartDescriptorClass()
	rv := objc.Send[AXChartDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a chart descriptor with the specified attributed title, summary,
// x-axis descriptor, y-axis descriptor, descriptors for additional axes, and
// array of data series.
//
// See: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/initWithAttributedTitle:summary:xAxisDescriptor:yAxisDescriptor:additionalAxes:series:
func NewAXChartDescriptorWithAttributedTitleSummaryXAxisDescriptorYAxisDescriptorAdditionalAxesSeries(attributedTitle foundation.NSAttributedString, summary string, xAxis AXDataAxisDescriptor, yAxis IAXNumericDataAxisDescriptor, additionalAxes []objectivec.IObject, series []AXDataSeriesDescriptor) AXChartDescriptor {
	instance := getAXChartDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAttributedTitle:summary:xAxisDescriptor:yAxisDescriptor:additionalAxes:series:"), attributedTitle, objc.String(summary), xAxis, yAxis, objectivec.IObjectSliceToNSArray(additionalAxes), objectivec.IObjectSliceToNSArray(series))
	return AXChartDescriptorFromID(rv)
}

// Creates a chart descriptor with the specified attributed title, summary,
// x-axis descriptor, y-axis descriptor, and array of data series.
//
// See: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/initWithAttributedTitle:summary:xAxisDescriptor:yAxisDescriptor:series:
func NewAXChartDescriptorWithAttributedTitleSummaryXAxisDescriptorYAxisDescriptorSeries(attributedTitle foundation.NSAttributedString, summary string, xAxis AXDataAxisDescriptor, yAxis IAXNumericDataAxisDescriptor, series []AXDataSeriesDescriptor) AXChartDescriptor {
	instance := getAXChartDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAttributedTitle:summary:xAxisDescriptor:yAxisDescriptor:series:"), attributedTitle, objc.String(summary), xAxis, yAxis, objectivec.IObjectSliceToNSArray(series))
	return AXChartDescriptorFromID(rv)
}

// Creates a chart descriptor with the specified title, summary, x-axis
// descriptor, y-axis descriptor, descriptors for additional axes, and array
// of data series.
//
// See: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/initWithTitle:summary:xAxisDescriptor:yAxisDescriptor:additionalAxes:series:
func NewAXChartDescriptorWithTitleSummaryXAxisDescriptorYAxisDescriptorAdditionalAxesSeries(title string, summary string, xAxis AXDataAxisDescriptor, yAxis IAXNumericDataAxisDescriptor, additionalAxes []objectivec.IObject, series []AXDataSeriesDescriptor) AXChartDescriptor {
	instance := getAXChartDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTitle:summary:xAxisDescriptor:yAxisDescriptor:additionalAxes:series:"), objc.String(title), objc.String(summary), xAxis, yAxis, objectivec.IObjectSliceToNSArray(additionalAxes), objectivec.IObjectSliceToNSArray(series))
	return AXChartDescriptorFromID(rv)
}

// Creates a chart descriptor with the specified title, summary, x-axis
// descriptor, y-axis descriptor, descriptors for additional axes, and array
// of data series.
//
// See: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/initWithTitle:summary:xAxisDescriptor:yAxisDescriptor:series:
func NewAXChartDescriptorWithTitleSummaryXAxisDescriptorYAxisDescriptorSeries(title string, summary string, xAxis AXDataAxisDescriptor, yAxis IAXNumericDataAxisDescriptor, series []AXDataSeriesDescriptor) AXChartDescriptor {
	instance := getAXChartDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTitle:summary:xAxisDescriptor:yAxisDescriptor:series:"), objc.String(title), objc.String(summary), xAxis, yAxis, objectivec.IObjectSliceToNSArray(series))
	return AXChartDescriptorFromID(rv)
}

// Creates a chart descriptor with the specified attributed title, summary,
// x-axis descriptor, y-axis descriptor, descriptors for additional axes, and
// array of data series.
//
// See: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/initWithAttributedTitle:summary:xAxisDescriptor:yAxisDescriptor:additionalAxes:series:
func (a AXChartDescriptor) InitWithAttributedTitleSummaryXAxisDescriptorYAxisDescriptorAdditionalAxesSeries(attributedTitle foundation.NSAttributedString, summary string, xAxis AXDataAxisDescriptor, yAxis IAXNumericDataAxisDescriptor, additionalAxes []objectivec.IObject, series []AXDataSeriesDescriptor) AXChartDescriptor {
	rv := objc.Send[AXChartDescriptor](a.ID, objc.Sel("initWithAttributedTitle:summary:xAxisDescriptor:yAxisDescriptor:additionalAxes:series:"), attributedTitle, objc.String(summary), xAxis, yAxis, objectivec.IObjectSliceToNSArray(additionalAxes), objectivec.IObjectSliceToNSArray(series))
	return rv
}

// Creates a chart descriptor with the specified attributed title, summary,
// x-axis descriptor, y-axis descriptor, and array of data series.
//
// See: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/initWithAttributedTitle:summary:xAxisDescriptor:yAxisDescriptor:series:
func (a AXChartDescriptor) InitWithAttributedTitleSummaryXAxisDescriptorYAxisDescriptorSeries(attributedTitle foundation.NSAttributedString, summary string, xAxis AXDataAxisDescriptor, yAxis IAXNumericDataAxisDescriptor, series []AXDataSeriesDescriptor) AXChartDescriptor {
	rv := objc.Send[AXChartDescriptor](a.ID, objc.Sel("initWithAttributedTitle:summary:xAxisDescriptor:yAxisDescriptor:series:"), attributedTitle, objc.String(summary), xAxis, yAxis, objectivec.IObjectSliceToNSArray(series))
	return rv
}

// Creates a chart descriptor with the specified title, summary, x-axis
// descriptor, y-axis descriptor, descriptors for additional axes, and array
// of data series.
//
// See: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/initWithTitle:summary:xAxisDescriptor:yAxisDescriptor:additionalAxes:series:
func (a AXChartDescriptor) InitWithTitleSummaryXAxisDescriptorYAxisDescriptorAdditionalAxesSeries(title string, summary string, xAxis AXDataAxisDescriptor, yAxis IAXNumericDataAxisDescriptor, additionalAxes []objectivec.IObject, series []AXDataSeriesDescriptor) AXChartDescriptor {
	rv := objc.Send[AXChartDescriptor](a.ID, objc.Sel("initWithTitle:summary:xAxisDescriptor:yAxisDescriptor:additionalAxes:series:"), objc.String(title), objc.String(summary), xAxis, yAxis, objectivec.IObjectSliceToNSArray(additionalAxes), objectivec.IObjectSliceToNSArray(series))
	return rv
}

// Creates a chart descriptor with the specified title, summary, x-axis
// descriptor, y-axis descriptor, descriptors for additional axes, and array
// of data series.
//
// See: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/initWithTitle:summary:xAxisDescriptor:yAxisDescriptor:series:
func (a AXChartDescriptor) InitWithTitleSummaryXAxisDescriptorYAxisDescriptorSeries(title string, summary string, xAxis AXDataAxisDescriptor, yAxis IAXNumericDataAxisDescriptor, series []AXDataSeriesDescriptor) AXChartDescriptor {
	rv := objc.Send[AXChartDescriptor](a.ID, objc.Sel("initWithTitle:summary:xAxisDescriptor:yAxisDescriptor:series:"), objc.String(title), objc.String(summary), xAxis, yAxis, objectivec.IObjectSliceToNSArray(series))
	return rv
}

// The title of the chart.
//
// See: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/title
func (a AXChartDescriptor) Title() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}
func (a AXChartDescriptor) SetTitle(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setTitle:"), objc.String(value))
}

// An attributed version of the chart title.
//
// # Discussion
//
// If you set the value of this property, the system uses this value instead
// of [AXChartDescriptor.Title].
//
// See: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/attributedTitle
func (a AXChartDescriptor) AttributedTitle() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("attributedTitle"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (a AXChartDescriptor) SetAttributedTitle(value foundation.NSAttributedString) {
	objc.Send[struct{}](a.ID, objc.Sel("setAttributedTitle:"), value)
}

// A description of the key takeaways or features of the chart.
//
// # Discussion
//
// Provide a string that includes a brief description of the main takeaways or
// insights that the chart offers, for example, “The chart shows that fuel
// efficiency decreases as vehicle weight increases.”
//
// See: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/summary
func (a AXChartDescriptor) Summary() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("summary"))
	return foundation.NSStringFromID(rv).String()
}
func (a AXChartDescriptor) SetSummary(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setSummary:"), objc.String(value))
}

// The axis descriptor for the chart’s x-axis.
//
// See: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/xAxis-6dnxd
func (a AXChartDescriptor) XAxis() AXDataAxisDescriptor {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("xAxis"))
	return AXDataAxisDescriptorObjectFromID(rv)
}
func (a AXChartDescriptor) SetXAxis(value AXDataAxisDescriptor) {
	objc.Send[struct{}](a.ID, objc.Sel("setXAxis:"), value)
}

// The axis descriptor for the chart’s y-axis.
//
// See: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/yAxis
func (a AXChartDescriptor) YAxis() IAXNumericDataAxisDescriptor {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("yAxis"))
	return AXNumericDataAxisDescriptorFromID(objc.ID(rv))
}
func (a AXChartDescriptor) SetYAxis(value IAXNumericDataAxisDescriptor) {
	objc.Send[struct{}](a.ID, objc.Sel("setYAxis:"), value)
}

// The descriptors for additional categorical or numerical axes beyond the
// x-axis and y-axis.
//
// See: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/additionalAxes-9ldc0
func (a AXChartDescriptor) AdditionalAxes() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("additionalAxes"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (a AXChartDescriptor) SetAdditionalAxes(value []objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setAdditionalAxes:"), objectivec.IObjectSliceToNSArray(value))
}

// The descriptors for each data series in the chart.
//
// See: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/series
func (a AXChartDescriptor) Series() []AXDataSeriesDescriptor {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("series"))
	return objc.ConvertSlice(rv, func(id objc.ID) AXDataSeriesDescriptor {
		return AXDataSeriesDescriptorFromID(id)
	})
}
func (a AXChartDescriptor) SetSeries(value []AXDataSeriesDescriptor) {
	objc.Send[struct{}](a.ID, objc.Sel("setSeries:"), objectivec.IObjectSliceToNSArray(value))
}

// The bounds of the view, in screen coordinates, for visually rendering data
// values.
//
// See: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/contentFrame
func (a AXChartDescriptor) ContentFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](a.ID, objc.Sel("contentFrame"))
	return corefoundation.CGRect(rv)
}
func (a AXChartDescriptor) SetContentFrame(value corefoundation.CGRect) {
	objc.Send[struct{}](a.ID, objc.Sel("setContentFrame:"), value)
}

// The direction of the content in the chart.
//
// See: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/contentDirection-swift.property
func (a AXChartDescriptor) ContentDirection() AXChartDescriptorContentDirection {
	rv := objc.Send[AXChartDescriptorContentDirection](a.ID, objc.Sel("contentDirection"))
	return AXChartDescriptorContentDirection(rv)
}
func (a AXChartDescriptor) SetContentDirection(value AXChartDescriptorContentDirection) {
	objc.Send[struct{}](a.ID, objc.Sel("setContentDirection:"), value)
}
