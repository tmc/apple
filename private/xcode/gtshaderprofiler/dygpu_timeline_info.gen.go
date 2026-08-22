// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[DYGPUTimelineInfo](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

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
type IDYGPUTimelineInfo interface {
	objectivec.IObject

	// Topic: Methods

	ActiveCoreInfoMasksPerPeriodicSample() foundation.NSData
	SetActiveCoreInfoMasksPerPeriodicSample(value foundation.NSData)
	ActiveShadersPerPeriodicSample() foundation.NSData
	SetActiveShadersPerPeriodicSample(value foundation.NSData)
	DerivedCounterNames() foundation.INSArray
	SetDerivedCounterNames(value foundation.INSArray)
	DerivedCounters() foundation.NSData
	SetDerivedCounters(value foundation.NSData)
	EncodeWithCoder(coder foundation.INSCoder)
	EncoderTimelineInfos() foundation.NSData
	SetEncoderTimelineInfos(value foundation.NSData)
	EnumerateActiveShadersForAllSamples(samples VoidHandler)
	EnumerateActiveShadersForSampleAtIndexWithBlock(index uint32, block VoidHandler)
	MetalFXTimelineInfo() foundation.NSData
	SetMetalFXTimelineInfo(value foundation.NSData)
	NumActiveShadersPerPeriodicSample() foundation.NSData
	SetNumActiveShadersPerPeriodicSample(value foundation.NSData)
	NumPeriodicSamples() uint32
	SetNumPeriodicSamples(value uint32)
	Timestamps() foundation.NSData
	SetTimestamps(value foundation.NSData)
	InitWithCoder(coder foundation.INSCoder) DYGPUTimelineInfo
}

// Init initializes the instance.
func (d DYGPUTimelineInfo) Init() DYGPUTimelineInfo {
	rv := objc.SendIfResponds[DYGPUTimelineInfo](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DYGPUTimelineInfo) Autorelease() DYGPUTimelineInfo {
	rv := objc.SendIfResponds[DYGPUTimelineInfo](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDYGPUTimelineInfo creates a new DYGPUTimelineInfo instance.
func NewDYGPUTimelineInfo() DYGPUTimelineInfo {
	class := getDYGPUTimelineInfoClass()
	rv := objc.SendIfResponds[DYGPUTimelineInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewDYGPUTimelineInfoWithCoder(coder objectivec.IObject) DYGPUTimelineInfo {
	instance := getDYGPUTimelineInfoClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DYGPUTimelineInfoFromID(rv)
}

func (d DYGPUTimelineInfo) EncodeWithCoder(coder foundation.INSCoder) {
	objc.SendIfResponds[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (d DYGPUTimelineInfo) EnumerateActiveShadersForAllSamples(samples VoidHandler) {
	_block0, _ := NewVoidBlock(samples)
	objc.SendIfResponds[objc.ID](d.ID, objc.Sel("enumerateActiveShadersForAllSamples:"), _block0)
}
func (d DYGPUTimelineInfo) EnumerateActiveShadersForSampleAtIndexWithBlock(index uint32, block VoidHandler) {
	_block1, _ := NewVoidBlock(block)
	objc.SendIfResponds[objc.ID](d.ID, objc.Sel("enumerateActiveShadersForSampleAtIndex:withBlock:"), index, _block1)
}
func (d DYGPUTimelineInfo) InitWithCoder(coder foundation.INSCoder) DYGPUTimelineInfo {
	rv := objc.SendIfResponds[DYGPUTimelineInfo](d.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

func (_DYGPUTimelineInfoClass DYGPUTimelineInfoClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_DYGPUTimelineInfoClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (d DYGPUTimelineInfo) ActiveCoreInfoMasksPerPeriodicSample() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("activeCoreInfoMasksPerPeriodicSample"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (d DYGPUTimelineInfo) SetActiveCoreInfoMasksPerPeriodicSample(value foundation.NSData) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setActiveCoreInfoMasksPerPeriodicSample:"), value)
}
func (d DYGPUTimelineInfo) ActiveShadersPerPeriodicSample() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("activeShadersPerPeriodicSample"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (d DYGPUTimelineInfo) SetActiveShadersPerPeriodicSample(value foundation.NSData) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setActiveShadersPerPeriodicSample:"), value)
}
func (d DYGPUTimelineInfo) DerivedCounterNames() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("derivedCounterNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (d DYGPUTimelineInfo) SetDerivedCounterNames(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setDerivedCounterNames:"), value)
}
func (d DYGPUTimelineInfo) DerivedCounters() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("derivedCounters"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (d DYGPUTimelineInfo) SetDerivedCounters(value foundation.NSData) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setDerivedCounters:"), value)
}
func (d DYGPUTimelineInfo) EncoderTimelineInfos() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("encoderTimelineInfos"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (d DYGPUTimelineInfo) SetEncoderTimelineInfos(value foundation.NSData) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setEncoderTimelineInfos:"), value)
}
func (d DYGPUTimelineInfo) MetalFXTimelineInfo() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("metalFXTimelineInfo"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (d DYGPUTimelineInfo) SetMetalFXTimelineInfo(value foundation.NSData) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setMetalFXTimelineInfo:"), value)
}
func (d DYGPUTimelineInfo) NumActiveShadersPerPeriodicSample() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("numActiveShadersPerPeriodicSample"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (d DYGPUTimelineInfo) SetNumActiveShadersPerPeriodicSample(value foundation.NSData) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setNumActiveShadersPerPeriodicSample:"), value)
}
func (d DYGPUTimelineInfo) NumPeriodicSamples() uint32 {
	rv := objc.SendIfResponds[uint32](d.ID, objc.Sel("numPeriodicSamples"))
	return rv
}
func (d DYGPUTimelineInfo) SetNumPeriodicSamples(value uint32) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setNumPeriodicSamples:"), value)
}
func (d DYGPUTimelineInfo) Timestamps() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("timestamps"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (d DYGPUTimelineInfo) SetTimestamps(value foundation.NSData) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setTimestamps:"), value)
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
