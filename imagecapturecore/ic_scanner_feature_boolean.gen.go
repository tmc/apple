// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [ICScannerFeatureBoolean] class.
var (
	_ICScannerFeatureBooleanClass     ICScannerFeatureBooleanClass
	_ICScannerFeatureBooleanClassOnce sync.Once
)

func getICScannerFeatureBooleanClass() ICScannerFeatureBooleanClass {
	_ICScannerFeatureBooleanClassOnce.Do(func() {
		_ICScannerFeatureBooleanClass = ICScannerFeatureBooleanClass{class: objc.GetClass("ICScannerFeatureBoolean")}
	})
	return _ICScannerFeatureBooleanClass
}

// GetICScannerFeatureBooleanClass returns the class object for ICScannerFeatureBoolean.
func GetICScannerFeatureBooleanClass() ICScannerFeatureBooleanClass {
	return getICScannerFeatureBooleanClass()
}

type ICScannerFeatureBooleanClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic ICScannerFeatureBooleanClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic ICScannerFeatureBooleanClass) Alloc() ICScannerFeatureBoolean {
	rv := objc.Send[ICScannerFeatureBoolean](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// A feature with a value of [YES] or [NO].
//
// # Instance Properties
//
//   - [ICScannerFeatureBoolean.Value]
//   - [ICScannerFeatureBoolean.SetValue]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureBoolean
type ICScannerFeatureBoolean struct {
	ICScannerFeature
}

// ICScannerFeatureBooleanFromID constructs a [ICScannerFeatureBoolean] from an objc.ID.
//
// A feature with a value of [YES] or [NO].
func ICScannerFeatureBooleanFromID(id objc.ID) ICScannerFeatureBoolean {
	return ICScannerFeatureBoolean{ICScannerFeature: ICScannerFeatureFromID(id)}
}

// NOTE: ICScannerFeatureBoolean adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ICScannerFeatureBoolean] class.
//
// # Instance Properties
//
//   - [IICScannerFeatureBoolean.Value]
//   - [IICScannerFeatureBoolean.SetValue]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureBoolean
type IICScannerFeatureBoolean interface {
	IICScannerFeature

	// Topic: Instance Properties

	Value() bool
	SetValue(value bool)
}

// Init initializes the instance.
func (s ICScannerFeatureBoolean) Init() ICScannerFeatureBoolean {
	rv := objc.Send[ICScannerFeatureBoolean](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s ICScannerFeatureBoolean) Autorelease() ICScannerFeatureBoolean {
	rv := objc.Send[ICScannerFeatureBoolean](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerFeatureBoolean creates a new ICScannerFeatureBoolean instance.
func NewICScannerFeatureBoolean() ICScannerFeatureBoolean {
	class := getICScannerFeatureBooleanClass()
	rv := objc.Send[ICScannerFeatureBoolean](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureBoolean/value
func (s ICScannerFeatureBoolean) Value() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("value"))
	return rv
}
func (s ICScannerFeatureBoolean) SetValue(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setValue:"), value)
}
