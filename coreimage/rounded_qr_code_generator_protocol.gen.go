// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The protocol for the Rounded QR Code Generator filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator
type CIRoundedQRCodeGenerator interface {
	objectivec.IObject
	CIFilterProtocol

	// The fraction of the center space of the QRCode to fill with Color 1. If the size is 0.0 or the Correction Level is L or M, the center of the QRCode will be unaltered. The size will be limited to 0.25 if the Correction Level is Q. The size will be limited to 0.33 if the Correction Level is H.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/centerSpaceSize
	CenterSpaceSize() float32

	// The background color for the QRCode
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/color0
	Color0() ICIColor

	// The foreground color for the QRCode
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/color1
	Color1() ICIColor

	// QR Code correction level L, M, Q, or H.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/correctionLevel
	CorrectionLevel() string

	// The message to encode in the QR Code
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/message
	Message() foundation.INSData

	// If true then the data points in the QRCode should have a rounded appearance.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/roundedData
	RoundedData() bool

	// If 1, then the Finder Patterns in the QRCode should have a rounded appearance. If 2, then the Alignment Patterns will also be rounded
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/roundedMarkers
	RoundedMarkers() int

	// The scale factor to enlarge the QRCode by.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/scale
	Scale() float32

	// The fraction of the center space of the QRCode to fill with Color 1. If the size is 0.0 or the Correction Level is L or M, the center of the QRCode will be unaltered. The size will be limited to 0.25 if the Correction Level is Q. The size will be limited to 0.33 if the Correction Level is H.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/centerSpaceSize
	SetCenterSpaceSize(value float32)

	// The background color for the QRCode
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/color0
	SetColor0(value ICIColor)

	// The foreground color for the QRCode
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/color1
	SetColor1(value ICIColor)

	// QR Code correction level L, M, Q, or H.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/correctionLevel
	SetCorrectionLevel(value string)

	// The message to encode in the QR Code
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/message
	SetMessage(value foundation.INSData)

	// If true then the data points in the QRCode should have a rounded appearance.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/roundedData
	SetRoundedData(value bool)

	// If 1, then the Finder Patterns in the QRCode should have a rounded appearance. If 2, then the Alignment Patterns will also be rounded
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/roundedMarkers
	SetRoundedMarkers(value int)

	// The scale factor to enlarge the QRCode by.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/scale
	SetScale(value float32)
}

// CIRoundedQRCodeGeneratorObject wraps an existing Objective-C object that conforms to the CIRoundedQRCodeGenerator protocol.
type CIRoundedQRCodeGeneratorObject struct {
	objectivec.Object
}

func (o CIRoundedQRCodeGeneratorObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIRoundedQRCodeGeneratorObjectFromID constructs a [CIRoundedQRCodeGeneratorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIRoundedQRCodeGeneratorObjectFromID(id objc.ID) CIRoundedQRCodeGeneratorObject {
	return CIRoundedQRCodeGeneratorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The fraction of the center space of the QRCode to fill with Color 1. If the
// size is 0.0 or the Correction Level is L or M, the center of the QRCode
// will be unaltered. The size will be limited to 0.25 if the Correction Level
// is Q. The size will be limited to 0.33 if the Correction Level is H.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/centerSpaceSize
func (o CIRoundedQRCodeGeneratorObject) CenterSpaceSize() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("centerSpaceSize"))
	return rv
}

// The background color for the QRCode
//
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/color0
func (o CIRoundedQRCodeGeneratorObject) Color0() ICIColor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color0"))
	return CIColorFromID(rv)
}

// The foreground color for the QRCode
//
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/color1
func (o CIRoundedQRCodeGeneratorObject) Color1() ICIColor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color1"))
	return CIColorFromID(rv)
}

// QR Code correction level L, M, Q, or H.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/correctionLevel
func (o CIRoundedQRCodeGeneratorObject) CorrectionLevel() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("correctionLevel"))
	return foundation.NSStringFromID(rv).String()
}

// The message to encode in the QR Code
//
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/message
func (o CIRoundedQRCodeGeneratorObject) Message() foundation.INSData {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("message"))
	return foundation.NSDataFromID(rv)
}

// If true then the data points in the QRCode should have a rounded
// appearance.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/roundedData
func (o CIRoundedQRCodeGeneratorObject) RoundedData() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("roundedData"))
	return rv
}

// If 1, then the Finder Patterns in the QRCode should have a rounded
// appearance. If 2, then the Alignment Patterns will also be rounded
//
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/roundedMarkers
func (o CIRoundedQRCodeGeneratorObject) RoundedMarkers() int {
	rv := objc.Send[int](o.ID, objc.Sel("roundedMarkers"))
	return rv
}

// The scale factor to enlarge the QRCode by.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/scale
func (o CIRoundedQRCodeGeneratorObject) Scale() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("scale"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIRoundedQRCodeGeneratorObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The fraction of the center space of the QRCode to fill with Color 1. If the
// size is 0.0 or the Correction Level is L or M, the center of the QRCode
// will be unaltered. The size will be limited to 0.25 if the Correction Level
// is Q. The size will be limited to 0.33 if the Correction Level is H.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/centerSpaceSize
func (o CIRoundedQRCodeGeneratorObject) SetCenterSpaceSize(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenterSpaceSize:"), value)
}

// The background color for the QRCode
//
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/color0
func (o CIRoundedQRCodeGeneratorObject) SetColor0(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor0:"), value)
}

// The foreground color for the QRCode
//
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/color1
func (o CIRoundedQRCodeGeneratorObject) SetColor1(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor1:"), value)
}

// QR Code correction level L, M, Q, or H.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/correctionLevel
func (o CIRoundedQRCodeGeneratorObject) SetCorrectionLevel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setCorrectionLevel:"), objc.String(value))
}

// The message to encode in the QR Code
//
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/message
func (o CIRoundedQRCodeGeneratorObject) SetMessage(value foundation.INSData) {
	objc.Send[struct{}](o.ID, objc.Sel("setMessage:"), value)
}

// If true then the data points in the QRCode should have a rounded
// appearance.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/roundedData
func (o CIRoundedQRCodeGeneratorObject) SetRoundedData(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setRoundedData:"), value)
}

// If 1, then the Finder Patterns in the QRCode should have a rounded
// appearance. If 2, then the Alignment Patterns will also be rounded
//
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/roundedMarkers
func (o CIRoundedQRCodeGeneratorObject) SetRoundedMarkers(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setRoundedMarkers:"), value)
}

// The scale factor to enlarge the QRCode by.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedQRCodeGenerator/scale
func (o CIRoundedQRCodeGeneratorObject) SetScale(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setScale:"), value)
}
