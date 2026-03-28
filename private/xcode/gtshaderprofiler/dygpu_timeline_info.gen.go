// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DYGPUTimelineInfo] class.
var (
	_DYGPUTimelineInfoClass     DYGPUTimelineInfoClass
	_DYGPUTimelineInfoClassOnce sync.Once
)

func getDYGPUTimelineInfoClass() DYGPUTimelineInfoClass {
	_DYGPUTimelineInfoClassOnce.Do(func() {
		_DYGPUTimelineInfoClass = DYGPUTimelineInfoClass{class: objc.GetClass("DYGPUTimelineInfo")}
	})
	return _DYGPUTimelineInfoClass
}

// GetDYGPUTimelineInfoClass returns the class object for DYGPUTimelineInfo.
func GetDYGPUTimelineInfoClass() DYGPUTimelineInfoClass {
	return getDYGPUTimelineInfoClass()
}

type DYGPUTimelineInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DYGPUTimelineInfoClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DYGPUTimelineInfoClass) Alloc() DYGPUTimelineInfo {
	rv := objc.Send[DYGPUTimelineInfo](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [DYGPUTimelineInfo.ActiveCoreInfoMasksPerPeriodicSample]
//   - [DYGPUTimelineInfo.SetActiveCoreInfoMasksPerPeriodicSample]
//   - [DYGPUTimelineInfo.ActiveShadersPerPeriodicSample]
//   - [DYGPUTimelineInfo.SetActiveShadersPerPeriodicSample]
//   - [DYGPUTimelineInfo.DerivedCounterNames]
//   - [DYGPUTimelineInfo.SetDerivedCounterNames]
//   - [DYGPUTimelineInfo.DerivedCounters]
//   - [DYGPUTimelineInfo.SetDerivedCounters]
//   - [DYGPUTimelineInfo.EncodeWithCoder]
//   - [DYGPUTimelineInfo.EncoderTimelineInfos]
//   - [DYGPUTimelineInfo.SetEncoderTimelineInfos]
//   - [DYGPUTimelineInfo.EnumerateActiveShadersForAllSamples]
//   - [DYGPUTimelineInfo.EnumerateActiveShadersForSampleAtIndexWithBlock]
//   - [DYGPUTimelineInfo.MetalFXTimelineInfo]
//   - [DYGPUTimelineInfo.SetMetalFXTimelineInfo]
//   - [DYGPUTimelineInfo.NumActiveShadersPerPeriodicSample]
//   - [DYGPUTimelineInfo.SetNumActiveShadersPerPeriodicSample]
//   - [DYGPUTimelineInfo.NumPeriodicSamples]
//   - [DYGPUTimelineInfo.SetNumPeriodicSamples]
//   - [DYGPUTimelineInfo.Timestamps]
//   - [DYGPUTimelineInfo.SetTimestamps]
//   - [DYGPUTimelineInfo.InitWithCoder]
// See: https://developer.apple.com/documentation/GTShaderProfiler/DYGPUTimelineInfo
type DYGPUTimelineInfo struct {
	objectivec.Object
}

// DYGPUTimelineInfoFromID constructs a [DYGPUTimelineInfo] from an objc.ID.
func DYGPUTimelineInfoFromID(id objc.ID) DYGPUTimelineInfo {
	return DYGPUTimelineInfo{objectivec.Object{ID: id}}
}
// Ensure DYGPUTimelineInfo implements IDYGPUTimelineInfo.
var _ IDYGPUTimelineInfo = DYGPUTimelineInfo{}

// An interface definition for the [DYGPUTimelineInfo] class.
//
// # Methods
//
//   - [IDYGPUTimelineInfo.ActiveCoreInfoMasksPerPeriodicSample]
//   - [IDYGPUTimelineInfo.SetActiveCoreInfoMasksPerPeriodicSample]
//   - [IDYGPUTimelineInfo.ActiveShadersPerPeriodicSample]
//   - [IDYGPUTimelineInfo.SetActiveShadersPerPeriodicSample]
//   - [IDYGPUTimelineInfo.DerivedCounterNames]
//   - [IDYGPUTimelineInfo.SetDerivedCounterNames]
//   - [IDYGPUTimelineInfo.DerivedCounters]
//   - [IDYGPUTimelineInfo.SetDerivedCounters]
//   - [IDYGPUTimelineInfo.EncodeWithCoder]
//   - [IDYGPUTimelineInfo.EncoderTimelineInfos]
//   - [IDYGPUTimelineInfo.SetEncoderTimelineInfos]
//   - [IDYGPUTimelineInfo.EnumerateActiveShadersForAllSamples]
//   - [IDYGPUTimelineInfo.EnumerateActiveShadersForSampleAtIndexWithBlock]
//   - [IDYGPUTimelineInfo.MetalFXTimelineInfo]
//   - [IDYGPUTimelineInfo.SetMetalFXTimelineInfo]
//   - [IDYGPUTimelineInfo.NumActiveShadersPerPeriodicSample]
//   - [IDYGPUTimelineInfo.SetNumActiveShadersPerPeriodicSample]
//   - [IDYGPUTimelineInfo.NumPeriodicSamples]
//   - [IDYGPUTimelineInfo.SetNumPeriodicSamples]
//   - [IDYGPUTimelineInfo.Timestamps]
//   - [IDYGPUTimelineInfo.SetTimestamps]
//   - [IDYGPUTimelineInfo.InitWithCoder]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/DYGPUTimelineInfo
type IDYGPUTimelineInfo interface {
	objectivec.IObject

	// Topic: Methods

	ActiveCoreInfoMasksPerPeriodicSample() foundation.INSData
	SetActiveCoreInfoMasksPerPeriodicSample(value foundation.INSData)
	ActiveShadersPerPeriodicSample() foundation.INSData
	SetActiveShadersPerPeriodicSample(value foundation.INSData)
	DerivedCounterNames() foundation.INSArray
	SetDerivedCounterNames(value foundation.INSArray)
	DerivedCounters() foundation.INSData
	SetDerivedCounters(value foundation.INSData)
	EncodeWithCoder(coder foundation.INSCoder)
	EncoderTimelineInfos() foundation.INSData
	SetEncoderTimelineInfos(value foundation.INSData)
	EnumerateActiveShadersForAllSamples(samples VoidHandler)
	EnumerateActiveShadersForSampleAtIndexWithBlock(index uint32, block VoidHandler)
	MetalFXTimelineInfo() foundation.INSData
	SetMetalFXTimelineInfo(value foundation.INSData)
	NumActiveShadersPerPeriodicSample() foundation.INSData
	SetNumActiveShadersPerPeriodicSample(value foundation.INSData)
	NumPeriodicSamples() uint32
	SetNumPeriodicSamples(value uint32)
	Timestamps() foundation.INSData
	SetTimestamps(value foundation.INSData)
	InitWithCoder(coder foundation.INSCoder) DYGPUTimelineInfo
}

// Init initializes the instance.
func (d DYGPUTimelineInfo) Init() DYGPUTimelineInfo {
	rv := objc.Send[DYGPUTimelineInfo](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DYGPUTimelineInfo) Autorelease() DYGPUTimelineInfo {
	rv := objc.Send[DYGPUTimelineInfo](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDYGPUTimelineInfo creates a new DYGPUTimelineInfo instance.
func NewDYGPUTimelineInfo() DYGPUTimelineInfo {
	class := getDYGPUTimelineInfoClass()
	rv := objc.Send[DYGPUTimelineInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/DYGPUTimelineInfo/initWithCoder:
func NewDYGPUTimelineInfoWithCoder(coder objectivec.IObject) DYGPUTimelineInfo {
	instance := getDYGPUTimelineInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DYGPUTimelineInfoFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/DYGPUTimelineInfo/encodeWithCoder:
func (d DYGPUTimelineInfo) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/DYGPUTimelineInfo/enumerateActiveShadersForAllSamples:
func (d DYGPUTimelineInfo) EnumerateActiveShadersForAllSamples(samples VoidHandler) {
_block0, _ := NewVoidBlock(samples)
	objc.Send[objc.ID](d.ID, objc.Sel("enumerateActiveShadersForAllSamples:"), _block0)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/DYGPUTimelineInfo/enumerateActiveShadersForSampleAtIndex:withBlock:
func (d DYGPUTimelineInfo) EnumerateActiveShadersForSampleAtIndexWithBlock(index uint32, block VoidHandler) {
_block1, _ := NewVoidBlock(block)
	objc.Send[objc.ID](d.ID, objc.Sel("enumerateActiveShadersForSampleAtIndex:withBlock:"), index, _block1)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/DYGPUTimelineInfo/initWithCoder:
func (d DYGPUTimelineInfo) InitWithCoder(coder foundation.INSCoder) DYGPUTimelineInfo {
	rv := objc.Send[DYGPUTimelineInfo](d.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYGPUTimelineInfo/supportsSecureCoding
func (_DYGPUTimelineInfoClass DYGPUTimelineInfoClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_DYGPUTimelineInfoClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYGPUTimelineInfo/activeCoreInfoMasksPerPeriodicSample
func (d DYGPUTimelineInfo) ActiveCoreInfoMasksPerPeriodicSample() foundation.INSData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("activeCoreInfoMasksPerPeriodicSample"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (d DYGPUTimelineInfo) SetActiveCoreInfoMasksPerPeriodicSample(value foundation.INSData) {
	objc.Send[struct{}](d.ID, objc.Sel("setActiveCoreInfoMasksPerPeriodicSample:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/DYGPUTimelineInfo/activeShadersPerPeriodicSample
func (d DYGPUTimelineInfo) ActiveShadersPerPeriodicSample() foundation.INSData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("activeShadersPerPeriodicSample"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (d DYGPUTimelineInfo) SetActiveShadersPerPeriodicSample(value foundation.INSData) {
	objc.Send[struct{}](d.ID, objc.Sel("setActiveShadersPerPeriodicSample:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/DYGPUTimelineInfo/derivedCounterNames
func (d DYGPUTimelineInfo) DerivedCounterNames() foundation.INSArray {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("derivedCounterNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (d DYGPUTimelineInfo) SetDerivedCounterNames(value foundation.INSArray) {
	objc.Send[struct{}](d.ID, objc.Sel("setDerivedCounterNames:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/DYGPUTimelineInfo/derivedCounters
func (d DYGPUTimelineInfo) DerivedCounters() foundation.INSData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("derivedCounters"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (d DYGPUTimelineInfo) SetDerivedCounters(value foundation.INSData) {
	objc.Send[struct{}](d.ID, objc.Sel("setDerivedCounters:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/DYGPUTimelineInfo/encoderTimelineInfos
func (d DYGPUTimelineInfo) EncoderTimelineInfos() foundation.INSData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("encoderTimelineInfos"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (d DYGPUTimelineInfo) SetEncoderTimelineInfos(value foundation.INSData) {
	objc.Send[struct{}](d.ID, objc.Sel("setEncoderTimelineInfos:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/DYGPUTimelineInfo/metalFXTimelineInfo
func (d DYGPUTimelineInfo) MetalFXTimelineInfo() foundation.INSData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("metalFXTimelineInfo"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (d DYGPUTimelineInfo) SetMetalFXTimelineInfo(value foundation.INSData) {
	objc.Send[struct{}](d.ID, objc.Sel("setMetalFXTimelineInfo:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/DYGPUTimelineInfo/numActiveShadersPerPeriodicSample
func (d DYGPUTimelineInfo) NumActiveShadersPerPeriodicSample() foundation.INSData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("numActiveShadersPerPeriodicSample"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (d DYGPUTimelineInfo) SetNumActiveShadersPerPeriodicSample(value foundation.INSData) {
	objc.Send[struct{}](d.ID, objc.Sel("setNumActiveShadersPerPeriodicSample:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/DYGPUTimelineInfo/numPeriodicSamples
func (d DYGPUTimelineInfo) NumPeriodicSamples() uint32 {
	rv := objc.Send[uint32](d.ID, objc.Sel("numPeriodicSamples"))
	return rv
}
func (d DYGPUTimelineInfo) SetNumPeriodicSamples(value uint32) {
	objc.Send[struct{}](d.ID, objc.Sel("setNumPeriodicSamples:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/DYGPUTimelineInfo/timestamps
func (d DYGPUTimelineInfo) Timestamps() foundation.INSData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("timestamps"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (d DYGPUTimelineInfo) SetTimestamps(value foundation.INSData) {
	objc.Send[struct{}](d.ID, objc.Sel("setTimestamps:"), value)
}

// EnumerateActiveShadersForAllSamplesSync is a synchronous wrapper around [DYGPUTimelineInfo.EnumerateActiveShadersForAllSamples].
// It blocks until the completion handler fires or the context is cancelled.
func (d DYGPUTimelineInfo) EnumerateActiveShadersForAllSamplesSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	d.EnumerateActiveShadersForAllSamples(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateActiveShadersForSampleAtIndexWithBlockSync is a synchronous wrapper around [DYGPUTimelineInfo.EnumerateActiveShadersForSampleAtIndexWithBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (d DYGPUTimelineInfo) EnumerateActiveShadersForSampleAtIndexWithBlockSync(ctx context.Context, index uint32) error {
	done := make(chan struct{}, 1)
	d.EnumerateActiveShadersForSampleAtIndexWithBlock(index, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

