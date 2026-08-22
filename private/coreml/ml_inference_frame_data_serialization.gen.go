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
	rv := objc.SendIfResponds[MLInferenceFrameDataSerialization](objc.ID(mc.class), objc.Sel("alloc"))
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
type IMLInferenceFrameDataSerialization interface {
	objectivec.IObject

	// Topic: Methods

	ModelIOFrameData() bool
	SetModelIOFrameData(value bool)
	OutputDirectoryURL() foundation.NSURL
	SetOutputDirectoryURL(value foundation.NSURL)
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
func (m MLInferenceFrameDataSerialization) Init() MLInferenceFrameDataSerialization {
	rv := objc.SendIfResponds[MLInferenceFrameDataSerialization](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLInferenceFrameDataSerialization) Autorelease() MLInferenceFrameDataSerialization {
	rv := objc.SendIfResponds[MLInferenceFrameDataSerialization](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLInferenceFrameDataSerialization creates a new MLInferenceFrameDataSerialization instance.
func NewMLInferenceFrameDataSerialization() MLInferenceFrameDataSerialization {
	class := getMLInferenceFrameDataSerializationClass()
	rv := objc.SendIfResponds[MLInferenceFrameDataSerialization](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewInferenceFrameDataSerializationWithOutputDirectoryPrefix(directory objectivec.IObject, prefix objectivec.IObject) MLInferenceFrameDataSerialization {
	instance := getMLInferenceFrameDataSerializationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithOutputDirectory:prefix:"), directory, prefix)
	return MLInferenceFrameDataSerializationFromID(rv)
}

func NewInferenceFrameDataSerializationWithOutputDirectoryPrefixShouldOverwriteModelIOFrameDataSegmentIOFrameData(directory objectivec.IObject, prefix objectivec.IObject, overwrite bool, data bool, data2 bool) MLInferenceFrameDataSerialization {
	instance := getMLInferenceFrameDataSerializationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithOutputDirectory:prefix:shouldOverwrite:modelIOFrameData:segmentIOFrameData:"), directory, prefix, overwrite, data, data2)
	return MLInferenceFrameDataSerializationFromID(rv)
}

func (m MLInferenceFrameDataSerialization) InitWithOutputDirectoryPrefix(directory objectivec.IObject, prefix objectivec.IObject) MLInferenceFrameDataSerialization {
	rv := objc.SendIfResponds[MLInferenceFrameDataSerialization](m.ID, objc.Sel("initWithOutputDirectory:prefix:"), directory, prefix)
	return rv
}
func (m MLInferenceFrameDataSerialization) InitWithOutputDirectoryPrefixShouldOverwriteModelIOFrameDataSegmentIOFrameData(directory objectivec.IObject, prefix objectivec.IObject, overwrite bool, data bool, data2 bool) MLInferenceFrameDataSerialization {
	rv := objc.SendIfResponds[MLInferenceFrameDataSerialization](m.ID, objc.Sel("initWithOutputDirectory:prefix:shouldOverwrite:modelIOFrameData:segmentIOFrameData:"), directory, prefix, overwrite, data, data2)
	return rv
}

func (m MLInferenceFrameDataSerialization) ModelIOFrameData() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("modelIOFrameData"))
	return rv
}
func (m MLInferenceFrameDataSerialization) SetModelIOFrameData(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setModelIOFrameData:"), value)
}
func (m MLInferenceFrameDataSerialization) OutputDirectoryURL() foundation.NSURL {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputDirectoryURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (m MLInferenceFrameDataSerialization) SetOutputDirectoryURL(value foundation.NSURL) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setOutputDirectoryURL:"), value)
}
func (m MLInferenceFrameDataSerialization) Prefix() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("prefix"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLInferenceFrameDataSerialization) SetPrefix(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setPrefix:"), objc.String(value))
}
func (m MLInferenceFrameDataSerialization) SegmentIOFrameData() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("segmentIOFrameData"))
	return rv
}
func (m MLInferenceFrameDataSerialization) SetSegmentIOFrameData(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setSegmentIOFrameData:"), value)
}
func (m MLInferenceFrameDataSerialization) ShouldOverwrite() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("shouldOverwrite"))
	return rv
}
func (m MLInferenceFrameDataSerialization) SetShouldOverwrite(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setShouldOverwrite:"), value)
}
