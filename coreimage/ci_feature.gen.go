// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CIFeature] class.
var (
	_CIFeatureClass     CIFeatureClass
	_CIFeatureClassOnce sync.Once
)

func getCIFeatureClass() CIFeatureClass {
	_CIFeatureClassOnce.Do(func() {
		_CIFeatureClass = CIFeatureClass{class: objc.GetClass("CIFeature")}
	})
	return _CIFeatureClass
}

// GetCIFeatureClass returns the class object for CIFeature.
func GetCIFeatureClass() CIFeatureClass {
	return getCIFeatureClass()
}

type CIFeatureClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CIFeatureClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIFeatureClass) Alloc() CIFeature {
	rv := objc.Send[CIFeature](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The abstract superclass for objects representing notable features detected
// in an image.
//
// # Overview
// 
// A [CIFeature] object represents a portion of an image that a detector
// believes matches its criteria. Subclasses of CIFeature holds additional
// information specific to the detector that discovered the feature.
//
// # Feature Properties
//
//   - [CIFeature.Bounds]: The rectangle that holds discovered feature.
//   - [CIFeature.Type]: The type of feature that was discovered.
//
// # Feature Types
//
//   - [CIFeature.CIFeatureTypeFace]: A Core Image feature type for person’s face.
//   - [CIFeature.CIFeatureTypeRectangle]: A Core Image feature type for rectangular object.
//   - [CIFeature.CIFeatureTypeQRCode]: A Core Image feature type for QR code object.
//   - [CIFeature.CIFeatureTypeText]: A Core Image feature type for text.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFeature
type CIFeature struct {
	objectivec.Object
}

// CIFeatureFromID constructs a [CIFeature] from an objc.ID.
//
// The abstract superclass for objects representing notable features detected
// in an image.
func CIFeatureFromID(id objc.ID) CIFeature {
	return CIFeature{objectivec.Object{ID: id}}
}
// NOTE: CIFeature adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIFeature] class.
//
// # Feature Properties
//
//   - [ICIFeature.Bounds]: The rectangle that holds discovered feature.
//   - [ICIFeature.Type]: The type of feature that was discovered.
//
// # Feature Types
//
//   - [ICIFeature.CIFeatureTypeFace]: A Core Image feature type for person’s face.
//   - [ICIFeature.CIFeatureTypeRectangle]: A Core Image feature type for rectangular object.
//   - [ICIFeature.CIFeatureTypeQRCode]: A Core Image feature type for QR code object.
//   - [ICIFeature.CIFeatureTypeText]: A Core Image feature type for text.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFeature
type ICIFeature interface {
	objectivec.IObject

	// Topic: Feature Properties

	// The rectangle that holds discovered feature.
	Bounds() corefoundation.CGRect
	// The type of feature that was discovered.
	Type() string

	// Topic: Feature Types

	// A Core Image feature type for person’s face.
	CIFeatureTypeFace() string
	// A Core Image feature type for rectangular object.
	CIFeatureTypeRectangle() string
	// A Core Image feature type for QR code object.
	CIFeatureTypeQRCode() string
	// A Core Image feature type for text.
	CIFeatureTypeText() string
}

// Init initializes the instance.
func (f CIFeature) Init() CIFeature {
	rv := objc.Send[CIFeature](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f CIFeature) Autorelease() CIFeature {
	rv := objc.Send[CIFeature](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIFeature creates a new CIFeature instance.
func NewCIFeature() CIFeature {
	class := getCIFeatureClass()
	rv := objc.Send[CIFeature](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The rectangle that holds discovered feature.
//
// # Discussion
// 
// The rectangle is in the coordinate system of the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFeature/bounds
func (f CIFeature) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](f.ID, objc.Sel("bounds"))
	return corefoundation.CGRect(rv)
}
// The type of feature that was discovered.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFeature/type
func (f CIFeature) Type() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("type"))
	return foundation.NSStringFromID(rv).String()
}
// A Core Image feature type for person’s face.
//
// See: https://developer.apple.com/documentation/coreimage/cifeaturetypeface
func (f CIFeature) CIFeatureTypeFace() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("CIFeatureTypeFace"))
	return foundation.NSStringFromID(rv).String()
}
// A Core Image feature type for rectangular object.
//
// See: https://developer.apple.com/documentation/coreimage/cifeaturetyperectangle
func (f CIFeature) CIFeatureTypeRectangle() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("CIFeatureTypeRectangle"))
	return foundation.NSStringFromID(rv).String()
}
// A Core Image feature type for QR code object.
//
// See: https://developer.apple.com/documentation/coreimage/cifeaturetypeqrcode
func (f CIFeature) CIFeatureTypeQRCode() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("CIFeatureTypeQRCode"))
	return foundation.NSStringFromID(rv).String()
}
// A Core Image feature type for text.
//
// See: https://developer.apple.com/documentation/coreimage/cifeaturetypetext
func (f CIFeature) CIFeatureTypeText() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("CIFeatureTypeText"))
	return foundation.NSStringFromID(rv).String()
}

