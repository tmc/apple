// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelVisionFeaturePrintInfo] class.
var (
	_MLModelVisionFeaturePrintInfoClass     MLModelVisionFeaturePrintInfoClass
	_MLModelVisionFeaturePrintInfoClassOnce sync.Once
)

func getMLModelVisionFeaturePrintInfoClass() MLModelVisionFeaturePrintInfoClass {
	_MLModelVisionFeaturePrintInfoClassOnce.Do(func() {
		_MLModelVisionFeaturePrintInfoClass = MLModelVisionFeaturePrintInfoClass{class: objc.GetClass("MLModelVisionFeaturePrintInfo")}
	})
	return _MLModelVisionFeaturePrintInfoClass
}

// GetMLModelVisionFeaturePrintInfoClass returns the class object for MLModelVisionFeaturePrintInfo.
func GetMLModelVisionFeaturePrintInfoClass() MLModelVisionFeaturePrintInfoClass {
	return getMLModelVisionFeaturePrintInfoClass()
}

type MLModelVisionFeaturePrintInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelVisionFeaturePrintInfoClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelVisionFeaturePrintInfoClass) Alloc() MLModelVisionFeaturePrintInfo {
	rv := objc.Send[MLModelVisionFeaturePrintInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLModelVisionFeaturePrintInfo.FeatureExtractorParameters]
//   - [MLModelVisionFeaturePrintInfo.SetFeatureExtractorParameters]
//   - [MLModelVisionFeaturePrintInfo.PostVisionFeaturePrintModel]
//   - [MLModelVisionFeaturePrintInfo.SetPostVisionFeaturePrintModel]
//   - [MLModelVisionFeaturePrintInfo.Version]
//   - [MLModelVisionFeaturePrintInfo.SetVersion]
// See: https://developer.apple.com/documentation/CoreML/MLModelVisionFeaturePrintInfo
type MLModelVisionFeaturePrintInfo struct {
	objectivec.Object
}

// MLModelVisionFeaturePrintInfoFromID constructs a [MLModelVisionFeaturePrintInfo] from an objc.ID.
func MLModelVisionFeaturePrintInfoFromID(id objc.ID) MLModelVisionFeaturePrintInfo {
	return MLModelVisionFeaturePrintInfo{objectivec.Object{ID: id}}
}
// Ensure MLModelVisionFeaturePrintInfo implements IMLModelVisionFeaturePrintInfo.
var _ IMLModelVisionFeaturePrintInfo = MLModelVisionFeaturePrintInfo{}

// An interface definition for the [MLModelVisionFeaturePrintInfo] class.
//
// # Methods
//
//   - [IMLModelVisionFeaturePrintInfo.FeatureExtractorParameters]
//   - [IMLModelVisionFeaturePrintInfo.SetFeatureExtractorParameters]
//   - [IMLModelVisionFeaturePrintInfo.PostVisionFeaturePrintModel]
//   - [IMLModelVisionFeaturePrintInfo.SetPostVisionFeaturePrintModel]
//   - [IMLModelVisionFeaturePrintInfo.Version]
//   - [IMLModelVisionFeaturePrintInfo.SetVersion]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelVisionFeaturePrintInfo
type IMLModelVisionFeaturePrintInfo interface {
	objectivec.IObject

	// Topic: Methods

	FeatureExtractorParameters() objectivec.IObject
	SetFeatureExtractorParameters(value objectivec.IObject)
	PostVisionFeaturePrintModel() IMLModel
	SetPostVisionFeaturePrintModel(value IMLModel)
	Version() uint64
	SetVersion(value uint64)
}

// Init initializes the instance.
func (m MLModelVisionFeaturePrintInfo) Init() MLModelVisionFeaturePrintInfo {
	rv := objc.Send[MLModelVisionFeaturePrintInfo](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelVisionFeaturePrintInfo) Autorelease() MLModelVisionFeaturePrintInfo {
	rv := objc.Send[MLModelVisionFeaturePrintInfo](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelVisionFeaturePrintInfo creates a new MLModelVisionFeaturePrintInfo instance.
func NewMLModelVisionFeaturePrintInfo() MLModelVisionFeaturePrintInfo {
	class := getMLModelVisionFeaturePrintInfoClass()
	rv := objc.Send[MLModelVisionFeaturePrintInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelVisionFeaturePrintInfo/featureExtractorParameters
func (m MLModelVisionFeaturePrintInfo) FeatureExtractorParameters() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("featureExtractorParameters"))
	return objectivec.Object{ID: rv}
}
func (m MLModelVisionFeaturePrintInfo) SetFeatureExtractorParameters(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setFeatureExtractorParameters:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLModelVisionFeaturePrintInfo/postVisionFeaturePrintModel
func (m MLModelVisionFeaturePrintInfo) PostVisionFeaturePrintModel() IMLModel {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("postVisionFeaturePrintModel"))
	return MLModelFromID(objc.ID(rv))
}
func (m MLModelVisionFeaturePrintInfo) SetPostVisionFeaturePrintModel(value IMLModel) {
	objc.Send[struct{}](m.ID, objc.Sel("setPostVisionFeaturePrintModel:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLModelVisionFeaturePrintInfo/version
func (m MLModelVisionFeaturePrintInfo) Version() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("version"))
	return rv
}
func (m MLModelVisionFeaturePrintInfo) SetVersion(value uint64) {
	objc.Send[struct{}](m.ID, objc.Sel("setVersion:"), value)
}

