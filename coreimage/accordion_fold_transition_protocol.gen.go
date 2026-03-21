// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure an accordion fold transition filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAccordionFoldTransition
type CIAccordionFoldTransition interface {
	objectivec.IObject
	CIFilterProtocol
	CITransitionFilter

	// The height of the accordion-fold part of the transition.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAccordionFoldTransition/bottomHeight
	BottomHeight() float32

	// A value that specifies the intensity of the shadow in the transtion.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAccordionFoldTransition/foldShadowAmount
	FoldShadowAmount() float32

	// The number of folds used in the transition.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAccordionFoldTransition/numberOfFolds
	NumberOfFolds() float32

	// The height of the accordion-fold part of the transition.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAccordionFoldTransition/bottomHeight
	SetBottomHeight(value float32)

	// A value that specifies the intensity of the shadow in the transtion.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAccordionFoldTransition/foldShadowAmount
	SetFoldShadowAmount(value float32)

	// The number of folds used in the transition.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAccordionFoldTransition/numberOfFolds
	SetNumberOfFolds(value float32)
}

// CIAccordionFoldTransitionObject wraps an existing Objective-C object that conforms to the CIAccordionFoldTransition protocol.
type CIAccordionFoldTransitionObject struct {
	objectivec.Object
}
func (o CIAccordionFoldTransitionObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIAccordionFoldTransitionObjectFromID constructs a [CIAccordionFoldTransitionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIAccordionFoldTransitionObjectFromID(id objc.ID) CIAccordionFoldTransitionObject {
	return CIAccordionFoldTransitionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The height of the accordion-fold part of the transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAccordionFoldTransition/bottomHeight
func (o CIAccordionFoldTransitionObject) BottomHeight() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("bottomHeight"))
	return rv
	}
// A value that specifies the intensity of the shadow in the transtion.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAccordionFoldTransition/foldShadowAmount
func (o CIAccordionFoldTransitionObject) FoldShadowAmount() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("foldShadowAmount"))
	return rv
	}
// The number of folds used in the transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAccordionFoldTransition/numberOfFolds
func (o CIAccordionFoldTransitionObject) NumberOfFolds() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("numberOfFolds"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIAccordionFoldTransitionObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/inputImage
func (o CIAccordionFoldTransitionObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The target image for a transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/targetImage
func (o CIAccordionFoldTransitionObject) TargetImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("targetImage"))
	return CIImageFromID(rv)
	}
// The parametric time of the transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/time
func (o CIAccordionFoldTransitionObject) Time() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("time"))
	return rv
	}

func (o CIAccordionFoldTransitionObject) SetBottomHeight(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setBottomHeight:"), value)
}

func (o CIAccordionFoldTransitionObject) SetFoldShadowAmount(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setFoldShadowAmount:"), value)
}

func (o CIAccordionFoldTransitionObject) SetNumberOfFolds(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setNumberOfFolds:"), value)
}

func (o CIAccordionFoldTransitionObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIAccordionFoldTransitionObject) SetTargetImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setTargetImage:"), value)
}

func (o CIAccordionFoldTransitionObject) SetTime(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setTime:"), value)
}

