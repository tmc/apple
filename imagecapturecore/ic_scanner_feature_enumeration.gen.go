// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ICScannerFeatureEnumeration] class.
var (
	_ICScannerFeatureEnumerationClass     ICScannerFeatureEnumerationClass
	_ICScannerFeatureEnumerationClassOnce sync.Once
)

func getICScannerFeatureEnumerationClass() ICScannerFeatureEnumerationClass {
	_ICScannerFeatureEnumerationClassOnce.Do(func() {
		_ICScannerFeatureEnumerationClass = ICScannerFeatureEnumerationClass{class: objc.GetClass("ICScannerFeatureEnumeration")}
	})
	return _ICScannerFeatureEnumerationClass
}

// GetICScannerFeatureEnumerationClass returns the class object for ICScannerFeatureEnumeration.
func GetICScannerFeatureEnumerationClass() ICScannerFeatureEnumerationClass {
	return getICScannerFeatureEnumerationClass()
}

type ICScannerFeatureEnumerationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic ICScannerFeatureEnumerationClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic ICScannerFeatureEnumerationClass) Alloc() ICScannerFeatureEnumeration {
	rv := objc.Send[ICScannerFeatureEnumeration](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// A feature that can have one of several discrete values, strings or numbers.
//
// # Instance Properties
//
//   - [ICScannerFeatureEnumeration.CurrentValue]
//   - [ICScannerFeatureEnumeration.SetCurrentValue]
//   - [ICScannerFeatureEnumeration.DefaultValue]
//   - [ICScannerFeatureEnumeration.MenuItemLabels]
//   - [ICScannerFeatureEnumeration.MenuItemLabelsTooltips]
//   - [ICScannerFeatureEnumeration.Values]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureEnumeration
type ICScannerFeatureEnumeration struct {
	ICScannerFeature
}

// ICScannerFeatureEnumerationFromID constructs a [ICScannerFeatureEnumeration] from an objc.ID.
//
// A feature that can have one of several discrete values, strings or numbers.
func ICScannerFeatureEnumerationFromID(id objc.ID) ICScannerFeatureEnumeration {
	return ICScannerFeatureEnumeration{ICScannerFeature: ICScannerFeatureFromID(id)}
}

// NOTE: ICScannerFeatureEnumeration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ICScannerFeatureEnumeration] class.
//
// # Instance Properties
//
//   - [IICScannerFeatureEnumeration.CurrentValue]
//   - [IICScannerFeatureEnumeration.SetCurrentValue]
//   - [IICScannerFeatureEnumeration.DefaultValue]
//   - [IICScannerFeatureEnumeration.MenuItemLabels]
//   - [IICScannerFeatureEnumeration.MenuItemLabelsTooltips]
//   - [IICScannerFeatureEnumeration.Values]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureEnumeration
type IICScannerFeatureEnumeration interface {
	IICScannerFeature

	// Topic: Instance Properties

	CurrentValue() objectivec.IObject
	SetCurrentValue(value objectivec.IObject)
	DefaultValue() objectivec.IObject
	MenuItemLabels() []string
	MenuItemLabelsTooltips() []string
	Values() []foundation.NSNumber
}

// Init initializes the instance.
func (s ICScannerFeatureEnumeration) Init() ICScannerFeatureEnumeration {
	rv := objc.Send[ICScannerFeatureEnumeration](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s ICScannerFeatureEnumeration) Autorelease() ICScannerFeatureEnumeration {
	rv := objc.Send[ICScannerFeatureEnumeration](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerFeatureEnumeration creates a new ICScannerFeatureEnumeration instance.
func NewICScannerFeatureEnumeration() ICScannerFeatureEnumeration {
	class := getICScannerFeatureEnumerationClass()
	rv := objc.Send[ICScannerFeatureEnumeration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureEnumeration/currentValue
func (s ICScannerFeatureEnumeration) CurrentValue() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("currentValue"))
	return objectivec.Object{ID: rv}
}
func (s ICScannerFeatureEnumeration) SetCurrentValue(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setCurrentValue:"), value)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureEnumeration/defaultValue
func (s ICScannerFeatureEnumeration) DefaultValue() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("defaultValue"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureEnumeration/menuItemLabels
func (s ICScannerFeatureEnumeration) MenuItemLabels() []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("menuItemLabels"))
	return objc.ConvertSliceToStrings(rv)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureEnumeration/menuItemLabelsTooltips
func (s ICScannerFeatureEnumeration) MenuItemLabelsTooltips() []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("menuItemLabelsTooltips"))
	return objc.ConvertSliceToStrings(rv)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureEnumeration/values
func (s ICScannerFeatureEnumeration) Values() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("values"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
