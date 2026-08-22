// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [ICScannerFeatureTemplate] class.
var (
	_ICScannerFeatureTemplateClass     ICScannerFeatureTemplateClass
	_ICScannerFeatureTemplateClassOnce sync.Once
)

func getICScannerFeatureTemplateClass() ICScannerFeatureTemplateClass {
	_ICScannerFeatureTemplateClassOnce.Do(func() {
		_ICScannerFeatureTemplateClass = ICScannerFeatureTemplateClass{class: objc.GetClass("ICScannerFeatureTemplate")}
	})
	return _ICScannerFeatureTemplateClass
}

// GetICScannerFeatureTemplateClass returns the class object for ICScannerFeatureTemplate.
func GetICScannerFeatureTemplateClass() ICScannerFeatureTemplateClass {
	return getICScannerFeatureTemplateClass()
}

type ICScannerFeatureTemplateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic ICScannerFeatureTemplateClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic ICScannerFeatureTemplateClass) Alloc() ICScannerFeatureTemplate {
	rv := objc.Send[ICScannerFeatureTemplate](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// A group of one or more rectangular scan areas that can be used with a
// scanner functional unit.
//
// # Instance Properties
//
//   - [ICScannerFeatureTemplate.Targets]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureTemplate
type ICScannerFeatureTemplate struct {
	ICScannerFeature
}

// ICScannerFeatureTemplateFromID constructs a [ICScannerFeatureTemplate] from an objc.ID.
//
// A group of one or more rectangular scan areas that can be used with a
// scanner functional unit.
func ICScannerFeatureTemplateFromID(id objc.ID) ICScannerFeatureTemplate {
	return ICScannerFeatureTemplate{ICScannerFeature: ICScannerFeatureFromID(id)}
}

// NOTE: ICScannerFeatureTemplate adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ICScannerFeatureTemplate] class.
//
// # Instance Properties
//
//   - [IICScannerFeatureTemplate.Targets]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureTemplate
type IICScannerFeatureTemplate interface {
	IICScannerFeature

	// Topic: Instance Properties

	Targets() []foundation.NSArray
}

// Init initializes the instance.
func (s ICScannerFeatureTemplate) Init() ICScannerFeatureTemplate {
	rv := objc.Send[ICScannerFeatureTemplate](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s ICScannerFeatureTemplate) Autorelease() ICScannerFeatureTemplate {
	rv := objc.Send[ICScannerFeatureTemplate](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerFeatureTemplate creates a new ICScannerFeatureTemplate instance.
func NewICScannerFeatureTemplate() ICScannerFeatureTemplate {
	class := getICScannerFeatureTemplateClass()
	rv := objc.Send[ICScannerFeatureTemplate](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureTemplate/targets
func (s ICScannerFeatureTemplate) Targets() []foundation.NSArray {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("targets"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSArray {
		return foundation.NSArrayFromID(id)
	})
}
