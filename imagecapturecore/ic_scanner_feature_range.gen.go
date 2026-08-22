// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [ICScannerFeatureRange] class.
var (
	_ICScannerFeatureRangeClass     ICScannerFeatureRangeClass
	_ICScannerFeatureRangeClassOnce sync.Once
)

func getICScannerFeatureRangeClass() ICScannerFeatureRangeClass {
	_ICScannerFeatureRangeClassOnce.Do(func() {
		_ICScannerFeatureRangeClass = ICScannerFeatureRangeClass{class: objc.GetClass("ICScannerFeatureRange")}
	})
	return _ICScannerFeatureRangeClass
}

// GetICScannerFeatureRangeClass returns the class object for ICScannerFeatureRange.
func GetICScannerFeatureRangeClass() ICScannerFeatureRangeClass {
	return getICScannerFeatureRangeClass()
}

type ICScannerFeatureRangeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic ICScannerFeatureRangeClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic ICScannerFeatureRangeClass) Alloc() ICScannerFeatureRange {
	rv := objc.Send[ICScannerFeatureRange](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// A feature with a value that lies within a range.
//
// # Instance Properties
//
//   - [ICScannerFeatureRange.CurrentValue]
//   - [ICScannerFeatureRange.SetCurrentValue]
//   - [ICScannerFeatureRange.DefaultValue]
//   - [ICScannerFeatureRange.MaxValue]
//   - [ICScannerFeatureRange.MinValue]
//   - [ICScannerFeatureRange.StepSize]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureRange
type ICScannerFeatureRange struct {
	ICScannerFeature
}

// ICScannerFeatureRangeFromID constructs a [ICScannerFeatureRange] from an objc.ID.
//
// A feature with a value that lies within a range.
func ICScannerFeatureRangeFromID(id objc.ID) ICScannerFeatureRange {
	return ICScannerFeatureRange{ICScannerFeature: ICScannerFeatureFromID(id)}
}

// NOTE: ICScannerFeatureRange adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ICScannerFeatureRange] class.
//
// # Instance Properties
//
//   - [IICScannerFeatureRange.CurrentValue]
//   - [IICScannerFeatureRange.SetCurrentValue]
//   - [IICScannerFeatureRange.DefaultValue]
//   - [IICScannerFeatureRange.MaxValue]
//   - [IICScannerFeatureRange.MinValue]
//   - [IICScannerFeatureRange.StepSize]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureRange
type IICScannerFeatureRange interface {
	IICScannerFeature

	// Topic: Instance Properties

	CurrentValue() float64
	SetCurrentValue(value float64)
	DefaultValue() float64
	MaxValue() float64
	MinValue() float64
	StepSize() float64
}

// Init initializes the instance.
func (s ICScannerFeatureRange) Init() ICScannerFeatureRange {
	rv := objc.Send[ICScannerFeatureRange](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s ICScannerFeatureRange) Autorelease() ICScannerFeatureRange {
	rv := objc.Send[ICScannerFeatureRange](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerFeatureRange creates a new ICScannerFeatureRange instance.
func NewICScannerFeatureRange() ICScannerFeatureRange {
	class := getICScannerFeatureRangeClass()
	rv := objc.Send[ICScannerFeatureRange](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureRange/currentValue
func (s ICScannerFeatureRange) CurrentValue() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("currentValue"))
	return rv
}
func (s ICScannerFeatureRange) SetCurrentValue(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setCurrentValue:"), value)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureRange/defaultValue
func (s ICScannerFeatureRange) DefaultValue() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("defaultValue"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureRange/maxValue
func (s ICScannerFeatureRange) MaxValue() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("maxValue"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureRange/minValue
func (s ICScannerFeatureRange) MinValue() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("minValue"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureRange/stepSize
func (s ICScannerFeatureRange) StepSize() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("stepSize"))
	return rv
}
