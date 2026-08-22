// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ICScannerFeature] class.
var (
	_ICScannerFeatureClass     ICScannerFeatureClass
	_ICScannerFeatureClassOnce sync.Once
)

func getICScannerFeatureClass() ICScannerFeatureClass {
	_ICScannerFeatureClassOnce.Do(func() {
		_ICScannerFeatureClass = ICScannerFeatureClass{class: objc.GetClass("ICScannerFeature")}
	})
	return _ICScannerFeatureClass
}

// GetICScannerFeatureClass returns the class object for ICScannerFeature.
func GetICScannerFeatureClass() ICScannerFeatureClass {
	return getICScannerFeatureClass()
}

type ICScannerFeatureClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic ICScannerFeatureClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic ICScannerFeatureClass) Alloc() ICScannerFeature {
	rv := objc.Send[ICScannerFeature](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// An abstract class that describes a scanner feature.
//
// # Overview
//
// The ImageCaptureCore framework defines three concrete subclasses of scanner
// features: [ICScannerFeatureEnumeration], [ICScannerFeatureRange], and
// [ICScannerFeatureBoolean]. Scanner functional units may have one or more
// instances of these classes to allow users to choose scanner-specific
// settings or operations before performing a scan.
//
// # Instance Properties
//
//   - [ICScannerFeature.HumanReadableName]
//   - [ICScannerFeature.InternalName]
//   - [ICScannerFeature.Tooltip]
//   - [ICScannerFeature.Type]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeature
type ICScannerFeature struct {
	objectivec.Object
}

// ICScannerFeatureFromID constructs a [ICScannerFeature] from an objc.ID.
//
// An abstract class that describes a scanner feature.
func ICScannerFeatureFromID(id objc.ID) ICScannerFeature {
	return ICScannerFeature{objectivec.Object{ID: id}}
}

// NOTE: ICScannerFeature adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ICScannerFeature] class.
//
// # Instance Properties
//
//   - [IICScannerFeature.HumanReadableName]
//   - [IICScannerFeature.InternalName]
//   - [IICScannerFeature.Tooltip]
//   - [IICScannerFeature.Type]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeature
type IICScannerFeature interface {
	objectivec.IObject

	// Topic: Instance Properties

	HumanReadableName() string
	InternalName() string
	Tooltip() string
	Type() ICScannerFeatureType
}

// Init initializes the instance.
func (s ICScannerFeature) Init() ICScannerFeature {
	rv := objc.Send[ICScannerFeature](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s ICScannerFeature) Autorelease() ICScannerFeature {
	rv := objc.Send[ICScannerFeature](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerFeature creates a new ICScannerFeature instance.
func NewICScannerFeature() ICScannerFeature {
	class := getICScannerFeatureClass()
	rv := objc.Send[ICScannerFeature](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeature/humanReadableName
func (s ICScannerFeature) HumanReadableName() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("humanReadableName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeature/internalName
func (s ICScannerFeature) InternalName() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("internalName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeature/tooltip
func (s ICScannerFeature) Tooltip() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("tooltip"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeature/type
func (s ICScannerFeature) Type() ICScannerFeatureType {
	rv := objc.Send[ICScannerFeatureType](s.ID, objc.Sel("type"))
	return ICScannerFeatureType(rv)
}
