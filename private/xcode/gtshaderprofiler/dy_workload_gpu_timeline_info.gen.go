// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DYWorkloadGPUTimelineInfo] class.
var (
	_DYWorkloadGPUTimelineInfoClass     DYWorkloadGPUTimelineInfoClass
	_DYWorkloadGPUTimelineInfoClassOnce sync.Once
)

func getDYWorkloadGPUTimelineInfoClass() DYWorkloadGPUTimelineInfoClass {
	_DYWorkloadGPUTimelineInfoClassOnce.Do(func() {
		_DYWorkloadGPUTimelineInfoClass = DYWorkloadGPUTimelineInfoClass{class: objc.GetClass("DYWorkloadGPUTimelineInfo")}
	})
	return _DYWorkloadGPUTimelineInfoClass
}

// GetDYWorkloadGPUTimelineInfoClass returns the class object for DYWorkloadGPUTimelineInfo.
func GetDYWorkloadGPUTimelineInfoClass() DYWorkloadGPUTimelineInfoClass {
	return getDYWorkloadGPUTimelineInfoClass()
}

type DYWorkloadGPUTimelineInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DYWorkloadGPUTimelineInfoClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DYWorkloadGPUTimelineInfoClass) Alloc() DYWorkloadGPUTimelineInfo {
	rv := objc.Send[DYWorkloadGPUTimelineInfo](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DYWorkloadGPUTimelineInfo.AggregatedGPUTimelineInfo]
//   - [DYWorkloadGPUTimelineInfo.SetAggregatedGPUTimelineInfo]
//   - [DYWorkloadGPUTimelineInfo.CoalescedEncoderInfo]
//   - [DYWorkloadGPUTimelineInfo.SetCoalescedEncoderInfo]
//   - [DYWorkloadGPUTimelineInfo.ConsistentStateAchieved]
//   - [DYWorkloadGPUTimelineInfo.SetConsistentStateAchieved]
//   - [DYWorkloadGPUTimelineInfo.CoreCounts]
//   - [DYWorkloadGPUTimelineInfo.SetCoreCounts]
//   - [DYWorkloadGPUTimelineInfo.CounterGroups]
//   - [DYWorkloadGPUTimelineInfo.SetCounterGroups]
//   - [DYWorkloadGPUTimelineInfo.CreateCounterGroup]
//   - [DYWorkloadGPUTimelineInfo.DerivedEncoderCounterInfo]
//   - [DYWorkloadGPUTimelineInfo.SetDerivedEncoderCounterInfo]
//   - [DYWorkloadGPUTimelineInfo.EncodeWithCoder]
//   - [DYWorkloadGPUTimelineInfo.EnumerateEncoderDerivedData]
//   - [DYWorkloadGPUTimelineInfo.EnumerateEncoderDerivedDataAtIndexWithBlock]
//   - [DYWorkloadGPUTimelineInfo.IsMio]
//   - [DYWorkloadGPUTimelineInfo.MGPUTimelineInfoAtIndex]
//   - [DYWorkloadGPUTimelineInfo.MGPUTimelineInfos]
//   - [DYWorkloadGPUTimelineInfo.SetMGPUTimelineInfos]
//   - [DYWorkloadGPUTimelineInfo.MetalFXCallDuration]
//   - [DYWorkloadGPUTimelineInfo.PerRingSampledDerivedCounters]
//   - [DYWorkloadGPUTimelineInfo.SetPerRingSampledDerivedCounters]
//   - [DYWorkloadGPUTimelineInfo.ProfiledState]
//   - [DYWorkloadGPUTimelineInfo.SetProfiledState]
//   - [DYWorkloadGPUTimelineInfo.RestoreTimestamps]
//   - [DYWorkloadGPUTimelineInfo.SetRestoreTimestamps]
//   - [DYWorkloadGPUTimelineInfo.TimeBaseDenominator]
//   - [DYWorkloadGPUTimelineInfo.SetTimeBaseDenominator]
//   - [DYWorkloadGPUTimelineInfo.TimeBaseNumerator]
//   - [DYWorkloadGPUTimelineInfo.SetTimeBaseNumerator]
//   - [DYWorkloadGPUTimelineInfo.InitWithCoder]
//   - [DYWorkloadGPUTimelineInfo.Version]
//   - [DYWorkloadGPUTimelineInfo.SetVersion]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo
type DYWorkloadGPUTimelineInfo struct {
	objectivec.Object
}

// DYWorkloadGPUTimelineInfoFromID constructs a [DYWorkloadGPUTimelineInfo] from an objc.ID.
func DYWorkloadGPUTimelineInfoFromID(id objc.ID) DYWorkloadGPUTimelineInfo {
	return DYWorkloadGPUTimelineInfo{objectivec.Object{ID: id}}
}

// Ensure DYWorkloadGPUTimelineInfo implements IDYWorkloadGPUTimelineInfo.
var _ IDYWorkloadGPUTimelineInfo = DYWorkloadGPUTimelineInfo{}

// An interface definition for the [DYWorkloadGPUTimelineInfo] class.
//
// # Methods
//
//   - [IDYWorkloadGPUTimelineInfo.AggregatedGPUTimelineInfo]
//   - [IDYWorkloadGPUTimelineInfo.SetAggregatedGPUTimelineInfo]
//   - [IDYWorkloadGPUTimelineInfo.CoalescedEncoderInfo]
//   - [IDYWorkloadGPUTimelineInfo.SetCoalescedEncoderInfo]
//   - [IDYWorkloadGPUTimelineInfo.ConsistentStateAchieved]
//   - [IDYWorkloadGPUTimelineInfo.SetConsistentStateAchieved]
//   - [IDYWorkloadGPUTimelineInfo.CoreCounts]
//   - [IDYWorkloadGPUTimelineInfo.SetCoreCounts]
//   - [IDYWorkloadGPUTimelineInfo.CounterGroups]
//   - [IDYWorkloadGPUTimelineInfo.SetCounterGroups]
//   - [IDYWorkloadGPUTimelineInfo.CreateCounterGroup]
//   - [IDYWorkloadGPUTimelineInfo.DerivedEncoderCounterInfo]
//   - [IDYWorkloadGPUTimelineInfo.SetDerivedEncoderCounterInfo]
//   - [IDYWorkloadGPUTimelineInfo.EncodeWithCoder]
//   - [IDYWorkloadGPUTimelineInfo.EnumerateEncoderDerivedData]
//   - [IDYWorkloadGPUTimelineInfo.EnumerateEncoderDerivedDataAtIndexWithBlock]
//   - [IDYWorkloadGPUTimelineInfo.IsMio]
//   - [IDYWorkloadGPUTimelineInfo.MGPUTimelineInfoAtIndex]
//   - [IDYWorkloadGPUTimelineInfo.MGPUTimelineInfos]
//   - [IDYWorkloadGPUTimelineInfo.SetMGPUTimelineInfos]
//   - [IDYWorkloadGPUTimelineInfo.MetalFXCallDuration]
//   - [IDYWorkloadGPUTimelineInfo.PerRingSampledDerivedCounters]
//   - [IDYWorkloadGPUTimelineInfo.SetPerRingSampledDerivedCounters]
//   - [IDYWorkloadGPUTimelineInfo.ProfiledState]
//   - [IDYWorkloadGPUTimelineInfo.SetProfiledState]
//   - [IDYWorkloadGPUTimelineInfo.RestoreTimestamps]
//   - [IDYWorkloadGPUTimelineInfo.SetRestoreTimestamps]
//   - [IDYWorkloadGPUTimelineInfo.TimeBaseDenominator]
//   - [IDYWorkloadGPUTimelineInfo.SetTimeBaseDenominator]
//   - [IDYWorkloadGPUTimelineInfo.TimeBaseNumerator]
//   - [IDYWorkloadGPUTimelineInfo.SetTimeBaseNumerator]
//   - [IDYWorkloadGPUTimelineInfo.InitWithCoder]
//   - [IDYWorkloadGPUTimelineInfo.Version]
//   - [IDYWorkloadGPUTimelineInfo.SetVersion]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo
type IDYWorkloadGPUTimelineInfo interface {
	objectivec.IObject

	// Topic: Methods

	AggregatedGPUTimelineInfo() IDYGPUTimelineInfo
	SetAggregatedGPUTimelineInfo(value IDYGPUTimelineInfo)
	CoalescedEncoderInfo() foundation.INSDictionary
	SetCoalescedEncoderInfo(value foundation.INSDictionary)
	ConsistentStateAchieved() bool
	SetConsistentStateAchieved(value bool)
	CoreCounts() foundation.INSArray
	SetCoreCounts(value foundation.INSArray)
	CounterGroups() foundation.INSArray
	SetCounterGroups(value foundation.INSArray)
	CreateCounterGroup() objectivec.IObject
	DerivedEncoderCounterInfo() IDYGPUDerivedEncoderCounterInfo
	SetDerivedEncoderCounterInfo(value IDYGPUDerivedEncoderCounterInfo)
	EncodeWithCoder(coder foundation.INSCoder)
	EnumerateEncoderDerivedData(data VoidHandler)
	EnumerateEncoderDerivedDataAtIndexWithBlock(index uint32, block VoidHandler)
	IsMio() bool
	MGPUTimelineInfoAtIndex(index uint64) objectivec.IObject
	MGPUTimelineInfos() foundation.INSArray
	SetMGPUTimelineInfos(value foundation.INSArray)
	MetalFXCallDuration(duration uint64) uint64
	PerRingSampledDerivedCounters() foundation.INSArray
	SetPerRingSampledDerivedCounters(value foundation.INSArray)
	ProfiledState() uint32
	SetProfiledState(value uint32)
	RestoreTimestamps() foundation.INSArray
	SetRestoreTimestamps(value foundation.INSArray)
	TimeBaseDenominator() uint32
	SetTimeBaseDenominator(value uint32)
	TimeBaseNumerator() uint32
	SetTimeBaseNumerator(value uint32)
	InitWithCoder(coder foundation.INSCoder) DYWorkloadGPUTimelineInfo
	Version() uint32
	SetVersion(value uint32)
}

// Init initializes the instance.
func (d DYWorkloadGPUTimelineInfo) Init() DYWorkloadGPUTimelineInfo {
	rv := objc.Send[DYWorkloadGPUTimelineInfo](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DYWorkloadGPUTimelineInfo) Autorelease() DYWorkloadGPUTimelineInfo {
	rv := objc.Send[DYWorkloadGPUTimelineInfo](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDYWorkloadGPUTimelineInfo creates a new DYWorkloadGPUTimelineInfo instance.
func NewDYWorkloadGPUTimelineInfo() DYWorkloadGPUTimelineInfo {
	class := getDYWorkloadGPUTimelineInfoClass()
	rv := objc.Send[DYWorkloadGPUTimelineInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo/initWithCoder:
func NewDYWorkloadGPUTimelineInfoWithCoder(coder objectivec.IObject) DYWorkloadGPUTimelineInfo {
	instance := getDYWorkloadGPUTimelineInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DYWorkloadGPUTimelineInfoFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo/createCounterGroup
func (d DYWorkloadGPUTimelineInfo) CreateCounterGroup() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("createCounterGroup"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo/encodeWithCoder:
func (d DYWorkloadGPUTimelineInfo) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo/enumerateEncoderDerivedData:
func (d DYWorkloadGPUTimelineInfo) EnumerateEncoderDerivedData(data VoidHandler) {
	_block0, _ := NewVoidBlock(data)
	objc.Send[objc.ID](d.ID, objc.Sel("enumerateEncoderDerivedData:"), _block0)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo/enumerateEncoderDerivedDataAtIndex:withBlock:
func (d DYWorkloadGPUTimelineInfo) EnumerateEncoderDerivedDataAtIndexWithBlock(index uint32, block VoidHandler) {
	_block1, _ := NewVoidBlock(block)
	objc.Send[objc.ID](d.ID, objc.Sel("enumerateEncoderDerivedDataAtIndex:withBlock:"), index, _block1)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo/isMio
func (d DYWorkloadGPUTimelineInfo) IsMio() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isMio"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo/mGPUTimelineInfoAtIndex:
func (d DYWorkloadGPUTimelineInfo) MGPUTimelineInfoAtIndex(index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("mGPUTimelineInfoAtIndex:"), index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo/metalFXCallDuration:
func (d DYWorkloadGPUTimelineInfo) MetalFXCallDuration(duration uint64) uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("metalFXCallDuration:"), duration)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo/initWithCoder:
func (d DYWorkloadGPUTimelineInfo) InitWithCoder(coder foundation.INSCoder) DYWorkloadGPUTimelineInfo {
	rv := objc.Send[DYWorkloadGPUTimelineInfo](d.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo/supportsSecureCoding
func (_DYWorkloadGPUTimelineInfoClass DYWorkloadGPUTimelineInfoClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_DYWorkloadGPUTimelineInfoClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo/aggregatedGPUTimelineInfo
func (d DYWorkloadGPUTimelineInfo) AggregatedGPUTimelineInfo() IDYGPUTimelineInfo {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("aggregatedGPUTimelineInfo"))
	return DYGPUTimelineInfoFromID(objc.ID(rv))
}
func (d DYWorkloadGPUTimelineInfo) SetAggregatedGPUTimelineInfo(value IDYGPUTimelineInfo) {
	objc.Send[struct{}](d.ID, objc.Sel("setAggregatedGPUTimelineInfo:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo/coalescedEncoderInfo
func (d DYWorkloadGPUTimelineInfo) CoalescedEncoderInfo() foundation.INSDictionary {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("coalescedEncoderInfo"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (d DYWorkloadGPUTimelineInfo) SetCoalescedEncoderInfo(value foundation.INSDictionary) {
	objc.Send[struct{}](d.ID, objc.Sel("setCoalescedEncoderInfo:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo/consistentStateAchieved
func (d DYWorkloadGPUTimelineInfo) ConsistentStateAchieved() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("consistentStateAchieved"))
	return rv
}
func (d DYWorkloadGPUTimelineInfo) SetConsistentStateAchieved(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setConsistentStateAchieved:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo/coreCounts
func (d DYWorkloadGPUTimelineInfo) CoreCounts() foundation.INSArray {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("coreCounts"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (d DYWorkloadGPUTimelineInfo) SetCoreCounts(value foundation.INSArray) {
	objc.Send[struct{}](d.ID, objc.Sel("setCoreCounts:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo/counterGroups
func (d DYWorkloadGPUTimelineInfo) CounterGroups() foundation.INSArray {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("counterGroups"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (d DYWorkloadGPUTimelineInfo) SetCounterGroups(value foundation.INSArray) {
	objc.Send[struct{}](d.ID, objc.Sel("setCounterGroups:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo/derivedEncoderCounterInfo
func (d DYWorkloadGPUTimelineInfo) DerivedEncoderCounterInfo() IDYGPUDerivedEncoderCounterInfo {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("derivedEncoderCounterInfo"))
	return DYGPUDerivedEncoderCounterInfoFromID(objc.ID(rv))
}
func (d DYWorkloadGPUTimelineInfo) SetDerivedEncoderCounterInfo(value IDYGPUDerivedEncoderCounterInfo) {
	objc.Send[struct{}](d.ID, objc.Sel("setDerivedEncoderCounterInfo:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo/mGPUTimelineInfos
func (d DYWorkloadGPUTimelineInfo) MGPUTimelineInfos() foundation.INSArray {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("mGPUTimelineInfos"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (d DYWorkloadGPUTimelineInfo) SetMGPUTimelineInfos(value foundation.INSArray) {
	objc.Send[struct{}](d.ID, objc.Sel("setMGPUTimelineInfos:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo/perRingSampledDerivedCounters
func (d DYWorkloadGPUTimelineInfo) PerRingSampledDerivedCounters() foundation.INSArray {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("perRingSampledDerivedCounters"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (d DYWorkloadGPUTimelineInfo) SetPerRingSampledDerivedCounters(value foundation.INSArray) {
	objc.Send[struct{}](d.ID, objc.Sel("setPerRingSampledDerivedCounters:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo/profiledState
func (d DYWorkloadGPUTimelineInfo) ProfiledState() uint32 {
	rv := objc.Send[uint32](d.ID, objc.Sel("profiledState"))
	return rv
}
func (d DYWorkloadGPUTimelineInfo) SetProfiledState(value uint32) {
	objc.Send[struct{}](d.ID, objc.Sel("setProfiledState:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo/restoreTimestamps
func (d DYWorkloadGPUTimelineInfo) RestoreTimestamps() foundation.INSArray {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("restoreTimestamps"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (d DYWorkloadGPUTimelineInfo) SetRestoreTimestamps(value foundation.INSArray) {
	objc.Send[struct{}](d.ID, objc.Sel("setRestoreTimestamps:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo/timeBaseDenominator
func (d DYWorkloadGPUTimelineInfo) TimeBaseDenominator() uint32 {
	rv := objc.Send[uint32](d.ID, objc.Sel("timeBaseDenominator"))
	return rv
}
func (d DYWorkloadGPUTimelineInfo) SetTimeBaseDenominator(value uint32) {
	objc.Send[struct{}](d.ID, objc.Sel("setTimeBaseDenominator:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo/timeBaseNumerator
func (d DYWorkloadGPUTimelineInfo) TimeBaseNumerator() uint32 {
	rv := objc.Send[uint32](d.ID, objc.Sel("timeBaseNumerator"))
	return rv
}
func (d DYWorkloadGPUTimelineInfo) SetTimeBaseNumerator(value uint32) {
	objc.Send[struct{}](d.ID, objc.Sel("setTimeBaseNumerator:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/DYWorkloadGPUTimelineInfo/version
func (d DYWorkloadGPUTimelineInfo) Version() uint32 {
	rv := objc.Send[uint32](d.ID, objc.Sel("version"))
	return rv
}
func (d DYWorkloadGPUTimelineInfo) SetVersion(value uint32) {
	objc.Send[struct{}](d.ID, objc.Sel("setVersion:"), value)
}

// EnumerateEncoderDerivedDataSync is a synchronous wrapper around [DYWorkloadGPUTimelineInfo.EnumerateEncoderDerivedData].
// It blocks until the completion handler fires or the context is cancelled.
func (d DYWorkloadGPUTimelineInfo) EnumerateEncoderDerivedDataSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	d.EnumerateEncoderDerivedData(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateEncoderDerivedDataAtIndexWithBlockSync is a synchronous wrapper around [DYWorkloadGPUTimelineInfo.EnumerateEncoderDerivedDataAtIndexWithBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (d DYWorkloadGPUTimelineInfo) EnumerateEncoderDerivedDataAtIndexWithBlockSync(ctx context.Context, index uint32) error {
	done := make(chan struct{}, 1)
	d.EnumerateEncoderDerivedDataAtIndexWithBlock(index, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
