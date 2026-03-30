// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLInferenceFrameDataSerialization] class.
var (
	_MLInferenceFrameDataSerializationClass     MLInferenceFrameDataSerializationClass
	_MLInferenceFrameDataSerializationClassOnce sync.Once
)

func getMLInferenceFrameDataSerializationClass() MLInferenceFrameDataSerializationClass {
	_MLInferenceFrameDataSerializationClassOnce.Do(func() {
		_MLInferenceFrameDataSerializationClass = MLInferenceFrameDataSerializationClass{class: objc.GetClass("MLInferenceFrameDataSerialization")}
	})
	return _MLInferenceFrameDataSerializationClass
}

// GetMLInferenceFrameDataSerializationClass returns the class object for MLInferenceFrameDataSerialization.
func GetMLInferenceFrameDataSerializationClass() MLInferenceFrameDataSerializationClass {
	return getMLInferenceFrameDataSerializationClass()
}

type MLInferenceFrameDataSerializationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLInferenceFrameDataSerializationClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLInferenceFrameDataSerializationClass) Alloc() MLInferenceFrameDataSerialization {
	rv := objc.Send[MLInferenceFrameDataSerialization](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLInferenceFrameDataSerialization.ModelIOFrameData]
//   - [MLInferenceFrameDataSerialization.SetModelIOFrameData]
//   - [MLInferenceFrameDataSerialization.OutputDirectoryURL]
//   - [MLInferenceFrameDataSerialization.SetOutputDirectoryURL]
//   - [MLInferenceFrameDataSerialization.Prefix]
//   - [MLInferenceFrameDataSerialization.SetPrefix]
//   - [MLInferenceFrameDataSerialization.SegmentIOFrameData]
//   - [MLInferenceFrameDataSerialization.SetSegmentIOFrameData]
//   - [MLInferenceFrameDataSerialization.ShouldOverwrite]
//   - [MLInferenceFrameDataSerialization.SetShouldOverwrite]
//   - [MLInferenceFrameDataSerialization.InitWithOutputDirectoryPrefix]
//   - [MLInferenceFrameDataSerialization.InitWithOutputDirectoryPrefixShouldOverwriteModelIOFrameDataSegmentIOFrameData]
//
// See: https://developer.apple.com/documentation/CoreML/MLInferenceFrameDataSerialization
type MLInferenceFrameDataSerialization struct {
	objectivec.Object
}

// MLInferenceFrameDataSerializationFromID constructs a [MLInferenceFrameDataSerialization] from an objc.ID.
func MLInferenceFrameDataSerializationFromID(id objc.ID) MLInferenceFrameDataSerialization {
	return MLInferenceFrameDataSerialization{objectivec.Object{ID: id}}
}

// Ensure MLInferenceFrameDataSerialization implements IMLInferenceFrameDataSerialization.
var _ IMLInferenceFrameDataSerialization = MLInferenceFrameDataSerialization{}

// An interface definition for the [MLInferenceFrameDataSerialization] class.
//
// # Methods
//
//   - [IMLInferenceFrameDataSerialization.ModelIOFrameData]
//   - [IMLInferenceFrameDataSerialization.SetModelIOFrameData]
//   - [IMLInferenceFrameDataSerialization.OutputDirectoryURL]
//   - [IMLInferenceFrameDataSerialization.SetOutputDirectoryURL]
//   - [IMLInferenceFrameDataSerialization.Prefix]
//   - [IMLInferenceFrameDataSerialization.SetPrefix]
//   - [IMLInferenceFrameDataSerialization.SegmentIOFrameData]
//   - [IMLInferenceFrameDataSerialization.SetSegmentIOFrameData]
//   - [IMLInferenceFrameDataSerialization.ShouldOverwrite]
//   - [IMLInferenceFrameDataSerialization.SetShouldOverwrite]
//   - [IMLInferenceFrameDataSerialization.InitWithOutputDirectoryPrefix]
//   - [IMLInferenceFrameDataSerialization.InitWithOutputDirectoryPrefixShouldOverwriteModelIOFrameDataSegmentIOFrameData]
//
// See: https://developer.apple.com/documentation/CoreML/MLInferenceFrameDataSerialization
type IMLInferenceFrameDataSerialization interface {
	objectivec.IObject

	// Topic: Methods

	ModelIOFrameData() bool
	SetModelIOFrameData(value bool)
	OutputDirectoryURL() foundation.INSURL
	SetOutputDirectoryURL(value foundation.INSURL)
	Prefix() string
	SetPrefix(value string)
	SegmentIOFrameData() bool
	SetSegmentIOFrameData(value bool)
	ShouldOverwrite() bool
	SetShouldOverwrite(value bool)
	InitWithOutputDirectoryPrefix(directory objectivec.IObject, prefix objectivec.IObject) MLInferenceFrameDataSerialization
	InitWithOutputDirectoryPrefixShouldOverwriteModelIOFrameDataSegmentIOFrameData(directory objectivec.IObject, prefix objectivec.IObject, overwrite bool, data bool, data2 bool) MLInferenceFrameDataSerialization
}

// Init initializes the instance.
func (i MLInferenceFrameDataSerialization) Init() MLInferenceFrameDataSerialization {
	rv := objc.Send[MLInferenceFrameDataSerialization](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MLInferenceFrameDataSerialization) Autorelease() MLInferenceFrameDataSerialization {
	rv := objc.Send[MLInferenceFrameDataSerialization](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLInferenceFrameDataSerialization creates a new MLInferenceFrameDataSerialization instance.
func NewMLInferenceFrameDataSerialization() MLInferenceFrameDataSerialization {
	class := getMLInferenceFrameDataSerializationClass()
	rv := objc.Send[MLInferenceFrameDataSerialization](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLInferenceFrameDataSerialization/initWithOutputDirectory:prefix:
func NewInferenceFrameDataSerializationWithOutputDirectoryPrefix(directory objectivec.IObject, prefix objectivec.IObject) MLInferenceFrameDataSerialization {
	instance := getMLInferenceFrameDataSerializationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOutputDirectory:prefix:"), directory, prefix)
	return MLInferenceFrameDataSerializationFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLInferenceFrameDataSerialization/initWithOutputDirectory:prefix:shouldOverwrite:modelIOFrameData:segmentIOFrameData:
func NewInferenceFrameDataSerializationWithOutputDirectoryPrefixShouldOverwriteModelIOFrameDataSegmentIOFrameData(directory objectivec.IObject, prefix objectivec.IObject, overwrite bool, data bool, data2 bool) MLInferenceFrameDataSerialization {
	instance := getMLInferenceFrameDataSerializationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOutputDirectory:prefix:shouldOverwrite:modelIOFrameData:segmentIOFrameData:"), directory, prefix, overwrite, data, data2)
	return MLInferenceFrameDataSerializationFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLInferenceFrameDataSerialization/initWithOutputDirectory:prefix:
func (i MLInferenceFrameDataSerialization) InitWithOutputDirectoryPrefix(directory objectivec.IObject, prefix objectivec.IObject) MLInferenceFrameDataSerialization {
	rv := objc.Send[MLInferenceFrameDataSerialization](i.ID, objc.Sel("initWithOutputDirectory:prefix:"), directory, prefix)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLInferenceFrameDataSerialization/initWithOutputDirectory:prefix:shouldOverwrite:modelIOFrameData:segmentIOFrameData:
func (i MLInferenceFrameDataSerialization) InitWithOutputDirectoryPrefixShouldOverwriteModelIOFrameDataSegmentIOFrameData(directory objectivec.IObject, prefix objectivec.IObject, overwrite bool, data bool, data2 bool) MLInferenceFrameDataSerialization {
	rv := objc.Send[MLInferenceFrameDataSerialization](i.ID, objc.Sel("initWithOutputDirectory:prefix:shouldOverwrite:modelIOFrameData:segmentIOFrameData:"), directory, prefix, overwrite, data, data2)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLInferenceFrameDataSerialization/modelIOFrameData
func (i MLInferenceFrameDataSerialization) ModelIOFrameData() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("modelIOFrameData"))
	return rv
}
func (i MLInferenceFrameDataSerialization) SetModelIOFrameData(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setModelIOFrameData:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLInferenceFrameDataSerialization/outputDirectoryURL
func (i MLInferenceFrameDataSerialization) OutputDirectoryURL() foundation.INSURL {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("outputDirectoryURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (i MLInferenceFrameDataSerialization) SetOutputDirectoryURL(value foundation.INSURL) {
	objc.Send[struct{}](i.ID, objc.Sel("setOutputDirectoryURL:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLInferenceFrameDataSerialization/prefix
func (i MLInferenceFrameDataSerialization) Prefix() string {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("prefix"))
	return foundation.NSStringFromID(rv).String()
}
func (i MLInferenceFrameDataSerialization) SetPrefix(value string) {
	objc.Send[struct{}](i.ID, objc.Sel("setPrefix:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLInferenceFrameDataSerialization/segmentIOFrameData
func (i MLInferenceFrameDataSerialization) SegmentIOFrameData() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("segmentIOFrameData"))
	return rv
}
func (i MLInferenceFrameDataSerialization) SetSegmentIOFrameData(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setSegmentIOFrameData:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLInferenceFrameDataSerialization/shouldOverwrite
func (i MLInferenceFrameDataSerialization) ShouldOverwrite() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("shouldOverwrite"))
	return rv
}
func (i MLInferenceFrameDataSerialization) SetShouldOverwrite(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setShouldOverwrite:"), value)
}
