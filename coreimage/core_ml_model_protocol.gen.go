// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/coreml"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a Core ML model filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CICoreMLModel
type CICoreMLModel interface {
	objectivec.IObject
	CIFilterProtocol

	// A number that specifies which output of a multihead Core ML model applies the effect on the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICoreMLModel/headIndex
	HeadIndex() float32
	SetHeadIndex(value float32)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICoreMLModel/inputImage
	InputImage() ICIImage
	SetInputImage(value ICIImage)

	// The Core ML model used to apply the effect on the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICoreMLModel/model
	Model() *coreml.MLModel
	SetModel(value *coreml.MLModel)

	// A Boolean value that specifies whether to apply Softmax normalization to the output of the model.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICoreMLModel/softmaxNormalization
	SoftmaxNormalization() bool
	SetSoftmaxNormalization(value bool)
}

// CICoreMLModelObject wraps an existing Objective-C object that conforms to the CICoreMLModel protocol.
type CICoreMLModelObject struct {
	objectivec.Object
}

func (o CICoreMLModelObject) BaseObject() objectivec.Object {
	return o.Object
}

// CICoreMLModelObjectFromID constructs a [CICoreMLModelObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CICoreMLModelObjectFromID(id objc.ID) CICoreMLModelObject {
	return CICoreMLModelObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CICoreMLModelObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// A number that specifies which output of a multihead Core ML model applies
// the effect on the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CICoreMLModel/headIndex
func (o CICoreMLModelObject) HeadIndex() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("headIndex"))
	return float32(rv)
}

func (o CICoreMLModelObject) SetHeadIndex(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setHeadIndex:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CICoreMLModel/inputImage
func (o CICoreMLModelObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

func (o CICoreMLModelObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// The Core ML model used to apply the effect on the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CICoreMLModel/model
func (o CICoreMLModelObject) Model() *coreml.MLModel {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("model"))
	val := coreml.MLModelFromID(rv)
	return &val
}

func (o CICoreMLModelObject) SetModel(value *coreml.MLModel) {
	objc.Send[struct{}](o.ID, objc.Sel("setModel:"), value)
}

// A Boolean value that specifies whether to apply Softmax normalization to
// the output of the model.
//
// See: https://developer.apple.com/documentation/CoreImage/CICoreMLModel/softmaxNormalization
func (o CICoreMLModelObject) SoftmaxNormalization() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("softmaxNormalization"))
	return bool(rv)
}

func (o CICoreMLModelObject) SetSoftmaxNormalization(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setSoftmaxNormalization:"), value)
}
