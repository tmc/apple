// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The protocol for the System Tone Map filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CISystemToneMap
type CISystemToneMap interface {
	objectivec.IObject
	CIFilterProtocol

	// Specifies the current headroom of the intended display.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISystemToneMap/displayHeadroom
	DisplayHeadroom() float32

	// Specifies input image with content headroom and average light level properties.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISystemToneMap/inputImage
	InputImage() ICIImage

	// Specifies the preferred dynamic range behavior of the tone mapping. The value should be kCIDynamicRangeStandard, kCIDynamicRangeConstrainedHigh, kCIDynamicRangeHigh or nil.  If nil then it will behave as kCIDynamicRangeHigh.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISystemToneMap/preferredDynamicRange
	PreferredDynamicRange() CIDynamicRangeOption

	// Specifies the current headroom of the intended display.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISystemToneMap/displayHeadroom
	SetDisplayHeadroom(value float32)

	// Specifies input image with content headroom and average light level properties.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISystemToneMap/inputImage
	SetInputImage(value ICIImage)

	// Specifies the preferred dynamic range behavior of the tone mapping. The value should be kCIDynamicRangeStandard, kCIDynamicRangeConstrainedHigh, kCIDynamicRangeHigh or nil.  If nil then it will behave as kCIDynamicRangeHigh.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISystemToneMap/preferredDynamicRange
	SetPreferredDynamicRange(value CIDynamicRangeOption)
}

// CISystemToneMapObject wraps an existing Objective-C object that conforms to the CISystemToneMap protocol.
type CISystemToneMapObject struct {
	objectivec.Object
}

func (o CISystemToneMapObject) BaseObject() objectivec.Object {
	return o.Object
}

// CISystemToneMapObjectFromID constructs a [CISystemToneMapObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CISystemToneMapObjectFromID(id objc.ID) CISystemToneMapObject {
	return CISystemToneMapObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Specifies the current headroom of the intended display.
//
// See: https://developer.apple.com/documentation/CoreImage/CISystemToneMap/displayHeadroom
func (o CISystemToneMapObject) DisplayHeadroom() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("displayHeadroom"))
	return rv
}

// Specifies input image with content headroom and average light level
// properties.
//
// See: https://developer.apple.com/documentation/CoreImage/CISystemToneMap/inputImage
func (o CISystemToneMapObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// Specifies the preferred dynamic range behavior of the tone mapping. The
// value should be kCIDynamicRangeStandard, kCIDynamicRangeConstrainedHigh,
// kCIDynamicRangeHigh or nil. If nil then it will behave as
// kCIDynamicRangeHigh.
//
// See: https://developer.apple.com/documentation/CoreImage/CISystemToneMap/preferredDynamicRange
func (o CISystemToneMapObject) PreferredDynamicRange() CIDynamicRangeOption {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("preferredDynamicRange"))
	return CIDynamicRangeOption(foundation.NSStringFromID(rv).String())
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CISystemToneMapObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// Specifies the current headroom of the intended display.
//
// See: https://developer.apple.com/documentation/CoreImage/CISystemToneMap/displayHeadroom
func (o CISystemToneMapObject) SetDisplayHeadroom(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setDisplayHeadroom:"), value)
}

// Specifies input image with content headroom and average light level
// properties.
//
// See: https://developer.apple.com/documentation/CoreImage/CISystemToneMap/inputImage
func (o CISystemToneMapObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// Specifies the preferred dynamic range behavior of the tone mapping. The
// value should be kCIDynamicRangeStandard, kCIDynamicRangeConstrainedHigh,
// kCIDynamicRangeHigh or nil. If nil then it will behave as
// kCIDynamicRangeHigh.
//
// See: https://developer.apple.com/documentation/CoreImage/CISystemToneMap/preferredDynamicRange
func (o CISystemToneMapObject) SetPreferredDynamicRange(value CIDynamicRangeOption) {
	objc.Send[struct{}](o.ID, objc.Sel("setPreferredDynamicRange:"), objc.String(string(value)))
}
