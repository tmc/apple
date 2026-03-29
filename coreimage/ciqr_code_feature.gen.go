// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [CIQRCodeFeature] class.
var (
	_CIQRCodeFeatureClass     CIQRCodeFeatureClass
	_CIQRCodeFeatureClassOnce sync.Once
)

func getCIQRCodeFeatureClass() CIQRCodeFeatureClass {
	_CIQRCodeFeatureClassOnce.Do(func() {
		_CIQRCodeFeatureClass = CIQRCodeFeatureClass{class: objc.GetClass("CIQRCodeFeature")}
	})
	return _CIQRCodeFeatureClass
}

// GetCIQRCodeFeatureClass returns the class object for CIQRCodeFeature.
func GetCIQRCodeFeatureClass() CIQRCodeFeatureClass {
	return getCIQRCodeFeatureClass()
}

type CIQRCodeFeatureClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CIQRCodeFeatureClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIQRCodeFeatureClass) Alloc() CIQRCodeFeature {
	rv := objc.Send[CIQRCodeFeature](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// Information about a Quick Response code detected in a still or video image.
//
// # Overview
// 
// A QR code is a two-dimensional barcode using the ISO/IEC 18004:2006
// standard. The properties of a CIQRCodeFeature object identify the corners
// of the barcode in the image perspective and provide the decoded message.
// 
// To detect QR codes in an image or video, choose [CIQRCodeFeature.CIDetectorTypeQRCode] type
// when initializing a [CIDetector] object.
//
// [CIQRCodeFeature.CIDetectorTypeQRCode]: https://developer.apple.com/documentation/CoreImage/CIDetectorTypeQRCode
//
// # Decoding a Detected Barcode
//
//   - [CIQRCodeFeature.MessageString]: The string decoded from the detected barcode.
//   - [CIQRCodeFeature.SymbolDescriptor]: An abstract representation of a QR Code symbol.
//
// # Identifying the Corners of a Detected Barcode
//
//   - [CIQRCodeFeature.BottomLeft]: The image coordinate of the lower-left corner of the detected QR code.
//   - [CIQRCodeFeature.BottomRight]: The image coordinate of the lower-right corner of the detected QR code.
//   - [CIQRCodeFeature.TopLeft]: The image coordinate of the upper-left corner of the detected QR code.
//   - [CIQRCodeFeature.TopRight]: The image coordinate of the upper-right corner of the detected QR code.
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature
type CIQRCodeFeature struct {
	CIFeature
}

// CIQRCodeFeatureFromID constructs a [CIQRCodeFeature] from an objc.ID.
//
// Information about a Quick Response code detected in a still or video image.
func CIQRCodeFeatureFromID(id objc.ID) CIQRCodeFeature {
	return CIQRCodeFeature{CIFeature: CIFeatureFromID(id)}
}
// NOTE: CIQRCodeFeature adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIQRCodeFeature] class.
//
// # Decoding a Detected Barcode
//
//   - [ICIQRCodeFeature.MessageString]: The string decoded from the detected barcode.
//   - [ICIQRCodeFeature.SymbolDescriptor]: An abstract representation of a QR Code symbol.
//
// # Identifying the Corners of a Detected Barcode
//
//   - [ICIQRCodeFeature.BottomLeft]: The image coordinate of the lower-left corner of the detected QR code.
//   - [ICIQRCodeFeature.BottomRight]: The image coordinate of the lower-right corner of the detected QR code.
//   - [ICIQRCodeFeature.TopLeft]: The image coordinate of the upper-left corner of the detected QR code.
//   - [ICIQRCodeFeature.TopRight]: The image coordinate of the upper-right corner of the detected QR code.
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature
type ICIQRCodeFeature interface {
	ICIFeature

	// Topic: Decoding a Detected Barcode

	// The string decoded from the detected barcode.
	MessageString() string
	// An abstract representation of a QR Code symbol.
	SymbolDescriptor() ICIQRCodeDescriptor

	// Topic: Identifying the Corners of a Detected Barcode

	// The image coordinate of the lower-left corner of the detected QR code.
	BottomLeft() corefoundation.CGPoint
	// The image coordinate of the lower-right corner of the detected QR code.
	BottomRight() corefoundation.CGPoint
	// The image coordinate of the upper-left corner of the detected QR code.
	TopLeft() corefoundation.CGPoint
	// The image coordinate of the upper-right corner of the detected QR code.
	TopRight() corefoundation.CGPoint

	// A detector that searches for Quick Response codes (a type of 2D barcode) in a still image or video, returning 
	CIDetectorTypeQRCode() string
}

// Init initializes the instance.
func (q CIQRCodeFeature) Init() CIQRCodeFeature {
	rv := objc.Send[CIQRCodeFeature](q.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q CIQRCodeFeature) Autorelease() CIQRCodeFeature {
	rv := objc.Send[CIQRCodeFeature](q.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIQRCodeFeature creates a new CIQRCodeFeature instance.
func NewCIQRCodeFeature() CIQRCodeFeature {
	class := getCIQRCodeFeatureClass()
	rv := objc.Send[CIQRCodeFeature](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The string decoded from the detected barcode.
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/messageString
func (q CIQRCodeFeature) MessageString() string {
	rv := objc.Send[objc.ID](q.ID, objc.Sel("messageString"))
	return foundation.NSStringFromID(rv).String()
}
// An abstract representation of a QR Code symbol.
//
// # Discussion
// 
// The property is a [CIQRCodeDescriptor] instance that contains the payload,
// symbol version, mask pattern, and error correction level, so the QR Code
// can be reproduced.
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/symbolDescriptor-swift.property
func (q CIQRCodeFeature) SymbolDescriptor() ICIQRCodeDescriptor {
	rv := objc.Send[objc.ID](q.ID, objc.Sel("symbolDescriptor"))
	return CIQRCodeDescriptorFromID(objc.ID(rv))
}
// The image coordinate of the lower-left corner of the detected QR code.
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/bottomLeft-swift.property
func (q CIQRCodeFeature) BottomLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](q.ID, objc.Sel("bottomLeft"))
	return corefoundation.CGPoint(rv)
}
// The image coordinate of the lower-right corner of the detected QR code.
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/bottomRight-swift.property
func (q CIQRCodeFeature) BottomRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](q.ID, objc.Sel("bottomRight"))
	return corefoundation.CGPoint(rv)
}
// The image coordinate of the upper-left corner of the detected QR code.
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/topLeft-swift.property
func (q CIQRCodeFeature) TopLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](q.ID, objc.Sel("topLeft"))
	return corefoundation.CGPoint(rv)
}
// The image coordinate of the upper-right corner of the detected QR code.
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/topRight-swift.property
func (q CIQRCodeFeature) TopRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](q.ID, objc.Sel("topRight"))
	return corefoundation.CGPoint(rv)
}
// A detector that searches for Quick Response codes (a type of 2D barcode) in
// a still image or video, returning
//
// See: https://developer.apple.com/documentation/coreimage/cidetectortypeqrcode
func (q CIQRCodeFeature) CIDetectorTypeQRCode() string {
	rv := objc.Send[objc.ID](q.ID, objc.Sel("CIDetectorTypeQRCode"))
	return foundation.NSStringFromID(rv).String()
}

