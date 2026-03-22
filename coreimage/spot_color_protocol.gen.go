// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a spot color filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CISpotColor
type CISpotColor interface {
	objectivec.IObject
	CIFilterProtocol

	// The center value of the first color range to replace.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/centerColor1
	CenterColor1() ICIColor

	// The center value of the second color range to replace.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/centerColor2
	CenterColor2() ICIColor

	// The center value of the third color range to replace.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/centerColor3
	CenterColor3() ICIColor

	// A value that indicates how closely the first color must match before it’s replaced.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/closeness1
	Closeness1() float32

	// A value that indicates how closely the second color must match before it’s replaced.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/closeness2
	Closeness2() float32

	// A value that indicates how closely the third color must match before it’s replaced.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/closeness3
	Closeness3() float32

	// The contrast of the first replacement color.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/contrast1
	Contrast1() float32

	// The contrast of the second replacement color.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/contrast2
	Contrast2() float32

	// The contrast of the third replacement color.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/contrast3
	Contrast3() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/inputImage
	InputImage() ICIImage

	// A replacement color for the first color range.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/replacementColor1
	ReplacementColor1() ICIColor

	// A replacement color for the second color range.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/replacementColor2
	ReplacementColor2() ICIColor

	// A replacement color for the third color range.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/replacementColor3
	ReplacementColor3() ICIColor

	// The center value of the first color range to replace.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/centerColor1
	SetCenterColor1(value ICIColor)

	// The center value of the second color range to replace.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/centerColor2
	SetCenterColor2(value ICIColor)

	// The center value of the third color range to replace.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/centerColor3
	SetCenterColor3(value ICIColor)

	// A value that indicates how closely the first color must match before it’s replaced.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/closeness1
	SetCloseness1(value float32)

	// A value that indicates how closely the second color must match before it’s replaced.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/closeness2
	SetCloseness2(value float32)

	// A value that indicates how closely the third color must match before it’s replaced.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/closeness3
	SetCloseness3(value float32)

	// The contrast of the first replacement color.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/contrast1
	SetContrast1(value float32)

	// The contrast of the second replacement color.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/contrast2
	SetContrast2(value float32)

	// The contrast of the third replacement color.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/contrast3
	SetContrast3(value float32)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/inputImage
	SetInputImage(value ICIImage)

	// A replacement color for the first color range.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/replacementColor1
	SetReplacementColor1(value ICIColor)

	// A replacement color for the second color range.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/replacementColor2
	SetReplacementColor2(value ICIColor)

	// A replacement color for the third color range.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/replacementColor3
	SetReplacementColor3(value ICIColor)
}

// CISpotColorObject wraps an existing Objective-C object that conforms to the CISpotColor protocol.
type CISpotColorObject struct {
	objectivec.Object
}
func (o CISpotColorObject) BaseObject() objectivec.Object {
	return o.Object
}

// CISpotColorObjectFromID constructs a [CISpotColorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CISpotColorObjectFromID(id objc.ID) CISpotColorObject {
	return CISpotColorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The center value of the first color range to replace.
//
// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/centerColor1
func (o CISpotColorObject) CenterColor1() ICIColor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("centerColor1"))
	return CIColorFromID(rv)
	}
// The center value of the second color range to replace.
//
// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/centerColor2
func (o CISpotColorObject) CenterColor2() ICIColor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("centerColor2"))
	return CIColorFromID(rv)
	}
// The center value of the third color range to replace.
//
// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/centerColor3
func (o CISpotColorObject) CenterColor3() ICIColor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("centerColor3"))
	return CIColorFromID(rv)
	}
// A value that indicates how closely the first color must match before it’s
// replaced.
//
// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/closeness1
func (o CISpotColorObject) Closeness1() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("closeness1"))
	return rv
	}
// A value that indicates how closely the second color must match before
// it’s replaced.
//
// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/closeness2
func (o CISpotColorObject) Closeness2() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("closeness2"))
	return rv
	}
// A value that indicates how closely the third color must match before it’s
// replaced.
//
// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/closeness3
func (o CISpotColorObject) Closeness3() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("closeness3"))
	return rv
	}
// The contrast of the first replacement color.
//
// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/contrast1
func (o CISpotColorObject) Contrast1() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("contrast1"))
	return rv
	}
// The contrast of the second replacement color.
//
// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/contrast2
func (o CISpotColorObject) Contrast2() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("contrast2"))
	return rv
	}
// The contrast of the third replacement color.
//
// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/contrast3
func (o CISpotColorObject) Contrast3() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("contrast3"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/inputImage
func (o CISpotColorObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A replacement color for the first color range.
//
// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/replacementColor1
func (o CISpotColorObject) ReplacementColor1() ICIColor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("replacementColor1"))
	return CIColorFromID(rv)
	}
// A replacement color for the second color range.
//
// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/replacementColor2
func (o CISpotColorObject) ReplacementColor2() ICIColor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("replacementColor2"))
	return CIColorFromID(rv)
	}
// A replacement color for the third color range.
//
// See: https://developer.apple.com/documentation/CoreImage/CISpotColor/replacementColor3
func (o CISpotColorObject) ReplacementColor3() ICIColor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("replacementColor3"))
	return CIColorFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CISpotColorObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CISpotColorObject) SetCenterColor1(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenterColor1:"), value)
}

func (o CISpotColorObject) SetCenterColor2(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenterColor2:"), value)
}

func (o CISpotColorObject) SetCenterColor3(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenterColor3:"), value)
}

func (o CISpotColorObject) SetCloseness1(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setCloseness1:"), value)
}

func (o CISpotColorObject) SetCloseness2(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setCloseness2:"), value)
}

func (o CISpotColorObject) SetCloseness3(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setCloseness3:"), value)
}

func (o CISpotColorObject) SetContrast1(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setContrast1:"), value)
}

func (o CISpotColorObject) SetContrast2(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setContrast2:"), value)
}

func (o CISpotColorObject) SetContrast3(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setContrast3:"), value)
}

func (o CISpotColorObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CISpotColorObject) SetReplacementColor1(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setReplacementColor1:"), value)
}

func (o CISpotColorObject) SetReplacementColor2(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setReplacementColor2:"), value)
}

func (o CISpotColorObject) SetReplacementColor3(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setReplacementColor3:"), value)
}

