// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AXDataSeriesDescriptor] class.
var (
	_AXDataSeriesDescriptorClass     AXDataSeriesDescriptorClass
	_AXDataSeriesDescriptorClassOnce sync.Once
)

func getAXDataSeriesDescriptorClass() AXDataSeriesDescriptorClass {
	_AXDataSeriesDescriptorClassOnce.Do(func() {
		_AXDataSeriesDescriptorClass = AXDataSeriesDescriptorClass{class: objc.GetClass("AXDataSeriesDescriptor")}
	})
	return _AXDataSeriesDescriptorClass
}

// GetAXDataSeriesDescriptorClass returns the class object for AXDataSeriesDescriptor.
func GetAXDataSeriesDescriptorClass() AXDataSeriesDescriptorClass {
	return getAXDataSeriesDescriptorClass()
}

type AXDataSeriesDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXDataSeriesDescriptorClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXDataSeriesDescriptorClass) Alloc() AXDataSeriesDescriptor {
	rv := objc.Send[AXDataSeriesDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a series of data points.
//
// # Creating a data series
//
//   - [AXDataSeriesDescriptor.InitWithNameIsContinuousDataPoints]: Creates a data series with the specified name, a Boolean value that indicates whether the series is continuous, and data points.
//   - [AXDataSeriesDescriptor.InitWithAttributedNameIsContinuousDataPoints]: Creates a data series with the specified attributed name, a Boolean value that indicates whether the series is continuous, and data points.
//
// # Naming the series
//
//   - [AXDataSeriesDescriptor.Name]: The name of the data series.
//   - [AXDataSeriesDescriptor.SetName]
//   - [AXDataSeriesDescriptor.AttributedName]: An attributed version of the data series name.
//   - [AXDataSeriesDescriptor.SetAttributedName]
//
// # Configuring the data points
//
//   - [AXDataSeriesDescriptor.IsContinuous]: A Boolean value that determines whether the data series is continuous.
//   - [AXDataSeriesDescriptor.SetIsContinuous]
//   - [AXDataSeriesDescriptor.DataPoints]: The data points that the series contains.
//   - [AXDataSeriesDescriptor.SetDataPoints]
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor
type AXDataSeriesDescriptor struct {
	objectivec.Object
}

// AXDataSeriesDescriptorFromID constructs a [AXDataSeriesDescriptor] from an objc.ID.
//
// An object that represents a series of data points.
func AXDataSeriesDescriptorFromID(id objc.ID) AXDataSeriesDescriptor {
	return AXDataSeriesDescriptor{objectivec.Object{ID: id}}
}

// NOTE: AXDataSeriesDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXDataSeriesDescriptor] class.
//
// # Creating a data series
//
//   - [IAXDataSeriesDescriptor.InitWithNameIsContinuousDataPoints]: Creates a data series with the specified name, a Boolean value that indicates whether the series is continuous, and data points.
//   - [IAXDataSeriesDescriptor.InitWithAttributedNameIsContinuousDataPoints]: Creates a data series with the specified attributed name, a Boolean value that indicates whether the series is continuous, and data points.
//
// # Naming the series
//
//   - [IAXDataSeriesDescriptor.Name]: The name of the data series.
//   - [IAXDataSeriesDescriptor.SetName]
//   - [IAXDataSeriesDescriptor.AttributedName]: An attributed version of the data series name.
//   - [IAXDataSeriesDescriptor.SetAttributedName]
//
// # Configuring the data points
//
//   - [IAXDataSeriesDescriptor.IsContinuous]: A Boolean value that determines whether the data series is continuous.
//   - [IAXDataSeriesDescriptor.SetIsContinuous]
//   - [IAXDataSeriesDescriptor.DataPoints]: The data points that the series contains.
//   - [IAXDataSeriesDescriptor.SetDataPoints]
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor
type IAXDataSeriesDescriptor interface {
	objectivec.IObject

	// Topic: Creating a data series

	// Creates a data series with the specified name, a Boolean value that indicates whether the series is continuous, and data points.
	InitWithNameIsContinuousDataPoints(name string, isContinuous bool, dataPoints []AXDataPoint) AXDataSeriesDescriptor
	// Creates a data series with the specified attributed name, a Boolean value that indicates whether the series is continuous, and data points.
	InitWithAttributedNameIsContinuousDataPoints(attributedName foundation.NSAttributedString, isContinuous bool, dataPoints []AXDataPoint) AXDataSeriesDescriptor

	// Topic: Naming the series

	// The name of the data series.
	Name() string
	SetName(value string)
	// An attributed version of the data series name.
	AttributedName() foundation.NSAttributedString
	SetAttributedName(value foundation.NSAttributedString)

	// Topic: Configuring the data points

	// A Boolean value that determines whether the data series is continuous.
	IsContinuous() bool
	SetIsContinuous(value bool)
	// The data points that the series contains.
	DataPoints() []AXDataPoint
	SetDataPoints(value []AXDataPoint)
}

// Init initializes the instance.
func (a AXDataSeriesDescriptor) Init() AXDataSeriesDescriptor {
	rv := objc.Send[AXDataSeriesDescriptor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXDataSeriesDescriptor) Autorelease() AXDataSeriesDescriptor {
	rv := objc.Send[AXDataSeriesDescriptor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXDataSeriesDescriptor creates a new AXDataSeriesDescriptor instance.
func NewAXDataSeriesDescriptor() AXDataSeriesDescriptor {
	class := getAXDataSeriesDescriptorClass()
	rv := objc.Send[AXDataSeriesDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a data series with the specified attributed name, a Boolean value
// that indicates whether the series is continuous, and data points.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/init(attributedName:isContinuous:dataPoints:)
func NewAXDataSeriesDescriptorWithAttributedNameIsContinuousDataPoints(attributedName foundation.NSAttributedString, isContinuous bool, dataPoints []AXDataPoint) AXDataSeriesDescriptor {
	instance := getAXDataSeriesDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAttributedName:isContinuous:dataPoints:"), attributedName, isContinuous, objectivec.IObjectSliceToNSArray(dataPoints))
	return AXDataSeriesDescriptorFromID(rv)
}

// Creates a data series with the specified name, a Boolean value that
// indicates whether the series is continuous, and data points.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/init(name:isContinuous:dataPoints:)
func NewAXDataSeriesDescriptorWithNameIsContinuousDataPoints(name string, isContinuous bool, dataPoints []AXDataPoint) AXDataSeriesDescriptor {
	instance := getAXDataSeriesDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:isContinuous:dataPoints:"), objc.String(name), isContinuous, objectivec.IObjectSliceToNSArray(dataPoints))
	return AXDataSeriesDescriptorFromID(rv)
}

// Creates a data series with the specified name, a Boolean value that
// indicates whether the series is continuous, and data points.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/init(name:isContinuous:dataPoints:)
func (a AXDataSeriesDescriptor) InitWithNameIsContinuousDataPoints(name string, isContinuous bool, dataPoints []AXDataPoint) AXDataSeriesDescriptor {
	rv := objc.Send[AXDataSeriesDescriptor](a.ID, objc.Sel("initWithName:isContinuous:dataPoints:"), objc.String(name), isContinuous, objectivec.IObjectSliceToNSArray(dataPoints))
	return rv
}

// Creates a data series with the specified attributed name, a Boolean value
// that indicates whether the series is continuous, and data points.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/init(attributedName:isContinuous:dataPoints:)
func (a AXDataSeriesDescriptor) InitWithAttributedNameIsContinuousDataPoints(attributedName foundation.NSAttributedString, isContinuous bool, dataPoints []AXDataPoint) AXDataSeriesDescriptor {
	rv := objc.Send[AXDataSeriesDescriptor](a.ID, objc.Sel("initWithAttributedName:isContinuous:dataPoints:"), attributedName, isContinuous, objectivec.IObjectSliceToNSArray(dataPoints))
	return rv
}

// The name of the data series.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/name
func (a AXDataSeriesDescriptor) Name() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (a AXDataSeriesDescriptor) SetName(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setName:"), objc.String(value))
}

// An attributed version of the data series name.
//
// # Discussion
//
// If you set the value of this property, the system uses this value instead
// of [AXDataSeriesDescriptor.Name].
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/attributedName
func (a AXDataSeriesDescriptor) AttributedName() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("attributedName"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (a AXDataSeriesDescriptor) SetAttributedName(value foundation.NSAttributedString) {
	objc.Send[struct{}](a.ID, objc.Sel("setAttributedName:"), value)
}

// A Boolean value that determines whether the data series is continuous.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/isContinuous
func (a AXDataSeriesDescriptor) IsContinuous() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isContinuous"))
	return rv
}
func (a AXDataSeriesDescriptor) SetIsContinuous(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setIsContinuous:"), value)
}

// The data points that the series contains.
//
// See: https://developer.apple.com/documentation/Accessibility/AXDataSeriesDescriptor/dataPoints
func (a AXDataSeriesDescriptor) DataPoints() []AXDataPoint {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("dataPoints"))
	return objc.ConvertSlice(rv, func(id objc.ID) AXDataPoint {
		return AXDataPointFromID(id)
	})
}
func (a AXDataSeriesDescriptor) SetDataPoints(value []AXDataPoint) {
	objc.Send[struct{}](a.ID, objc.Sel("setDataPoints:"), objectivec.IObjectSliceToNSArray(value))
}
