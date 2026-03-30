// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a checkerboard generator filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CICheckerboardGenerator
type CICheckerboardGenerator interface {
	objectivec.IObject
	CIFilterProtocol

	// The center of the effect as x and y coordinates.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICheckerboardGenerator/center
	Center() corefoundation.CGPoint

	// A color to use for the first set of squares.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICheckerboardGenerator/color0
	Color0() ICIColor

	// A color to use for the second set of squares.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICheckerboardGenerator/color1
	Color1() ICIColor

	// The sharpness of the edges in the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICheckerboardGenerator/sharpness
	Sharpness() float32

	// The width of the squares in the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICheckerboardGenerator/width
	Width() float32

	// The center of the effect as x and y coordinates.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICheckerboardGenerator/center
	SetCenter(value corefoundation.CGPoint)

	// A color to use for the first set of squares.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICheckerboardGenerator/color0
	SetColor0(value ICIColor)

	// A color to use for the second set of squares.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICheckerboardGenerator/color1
	SetColor1(value ICIColor)

	// The sharpness of the edges in the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICheckerboardGenerator/sharpness
	SetSharpness(value float32)

	// The width of the squares in the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICheckerboardGenerator/width
	SetWidth(value float32)
}

// CICheckerboardGeneratorObject wraps an existing Objective-C object that conforms to the CICheckerboardGenerator protocol.
type CICheckerboardGeneratorObject struct {
	objectivec.Object
}

func (o CICheckerboardGeneratorObject) BaseObject() objectivec.Object {
	return o.Object
}

// CICheckerboardGeneratorObjectFromID constructs a [CICheckerboardGeneratorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CICheckerboardGeneratorObjectFromID(id objc.ID) CICheckerboardGeneratorObject {
	return CICheckerboardGeneratorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The center of the effect as x and y coordinates.
//
// See: https://developer.apple.com/documentation/CoreImage/CICheckerboardGenerator/center
func (o CICheckerboardGeneratorObject) Center() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
}

// A color to use for the first set of squares.
//
// See: https://developer.apple.com/documentation/CoreImage/CICheckerboardGenerator/color0
func (o CICheckerboardGeneratorObject) Color0() ICIColor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color0"))
	return CIColorFromID(rv)
}

// A color to use for the second set of squares.
//
// See: https://developer.apple.com/documentation/CoreImage/CICheckerboardGenerator/color1
func (o CICheckerboardGeneratorObject) Color1() ICIColor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color1"))
	return CIColorFromID(rv)
}

// The sharpness of the edges in the pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CICheckerboardGenerator/sharpness
func (o CICheckerboardGeneratorObject) Sharpness() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("sharpness"))
	return rv
}

// The width of the squares in the pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CICheckerboardGenerator/width
func (o CICheckerboardGeneratorObject) Width() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("width"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CICheckerboardGeneratorObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The center of the effect as x and y coordinates.
//
// See: https://developer.apple.com/documentation/CoreImage/CICheckerboardGenerator/center
func (o CICheckerboardGeneratorObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

// A color to use for the first set of squares.
//
// See: https://developer.apple.com/documentation/CoreImage/CICheckerboardGenerator/color0
func (o CICheckerboardGeneratorObject) SetColor0(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor0:"), value)
}

// A color to use for the second set of squares.
//
// See: https://developer.apple.com/documentation/CoreImage/CICheckerboardGenerator/color1
func (o CICheckerboardGeneratorObject) SetColor1(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor1:"), value)
}

// The sharpness of the edges in the pattern.
//
// # Discussion
//
// The smaller the value, the more blurry the pattern. Values range from 0.0
// to 1.0.
//
// See: https://developer.apple.com/documentation/CoreImage/CICheckerboardGenerator/sharpness
func (o CICheckerboardGeneratorObject) SetSharpness(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setSharpness:"), value)
}

// The width of the squares in the pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CICheckerboardGenerator/width
func (o CICheckerboardGeneratorObject) SetWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setWidth:"), value)
}
