// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLDisplayPresetDevice] class.
var (
	_SLDisplayPresetDeviceClass     SLDisplayPresetDeviceClass
	_SLDisplayPresetDeviceClassOnce sync.Once
)

func getSLDisplayPresetDeviceClass() SLDisplayPresetDeviceClass {
	_SLDisplayPresetDeviceClassOnce.Do(func() {
		_SLDisplayPresetDeviceClass = SLDisplayPresetDeviceClass{class: objc.GetClass("SLDisplayPresetDevice")}
	})
	return _SLDisplayPresetDeviceClass
}

// GetSLDisplayPresetDeviceClass returns the class object for SLDisplayPresetDevice.
func GetSLDisplayPresetDeviceClass() SLDisplayPresetDeviceClass {
	return getSLDisplayPresetDeviceClass()
}

type SLDisplayPresetDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLDisplayPresetDeviceClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLDisplayPresetDeviceClass) Alloc() SLDisplayPresetDevice {
	rv := objc.SendIfResponds[SLDisplayPresetDevice](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLDisplayPresetDevice.ActivePresetIndex]
//   - [SLDisplayPresetDevice.ContainerId]
//   - [SLDisplayPresetDevice.CopyCFContainerId]
//   - [SLDisplayPresetDevice.CopyCalibrationInfo]
//   - [SLDisplayPresetDevice.CopyCustomPresetInfo]
//   - [SLDisplayPresetDevice.CopyPresetAtIndex]
//   - [SLDisplayPresetDevice.CopyPresetDataAtIndex]
//   - [SLDisplayPresetDevice.CopyUserAdjustment]
//   - [SLDisplayPresetDevice.CopyUserAdjustmentData]
//   - [SLDisplayPresetDevice.CopyUserAdjustmentForPreset]
//   - [SLDisplayPresetDevice.FactoryDefaultPresetIndex]
//   - [SLDisplayPresetDevice.FactoryResetWithType]
//   - [SLDisplayPresetDevice.GetUserAdjustmentPowerLimit]
//   - [SLDisplayPresetDevice.GetUserAdjustmentRangeWithInput]
//   - [SLDisplayPresetDevice.InvalidateLiveUserAdjustmentForPreset]
//   - [SLDisplayPresetDevice.InvalidateUserAdjustment]
//   - [SLDisplayPresetDevice.InvalidateUserAdjustmentForPreset]
//   - [SLDisplayPresetDevice.IsLiveUserAdjustmentSupported]
//   - [SLDisplayPresetDevice.IsPerPresetUserAdjustmentSupported]
//   - [SLDisplayPresetDevice.IsPresetValidAtIndex]
//   - [SLDisplayPresetDevice.IsPresetWritableAtIndex]
//   - [SLDisplayPresetDevice.IsUserAdjustmentSupported]
//   - [SLDisplayPresetDevice.IsUserAdjustmentValid]
//   - [SLDisplayPresetDevice.IsUserAdjustmentValidForAnyPreset]
//   - [SLDisplayPresetDevice.IsUserAdjustmentValidForPreset]
//   - [SLDisplayPresetDevice.PresetCapabilities]
//   - [SLDisplayPresetDevice.PresetCount]
//   - [SLDisplayPresetDevice.PresetUUIDAtIndexToBytes]
//   - [SLDisplayPresetDevice.ResetPresetAtIndex]
//   - [SLDisplayPresetDevice.ResetProController]
//   - [SLDisplayPresetDevice.SetActivePresetIndex]
//   - [SLDisplayPresetDevice.SetCustomPresetDataAtIndexWithData]
//   - [SLDisplayPresetDevice.SetPresetAtIndexWithData]
//   - [SLDisplayPresetDevice.SetPresetDataAtIndexWithData]
//   - [SLDisplayPresetDevice.SetUserAdjustment]
//   - [SLDisplayPresetDevice.SetUserAdjustmentData]
//   - [SLDisplayPresetDevice.SetUserAdjustmentForPresetWithData]
//   - [SLDisplayPresetDevice.InitWithService]
type SLDisplayPresetDevice struct {
	objectivec.Object
}

// SLDisplayPresetDeviceFromID constructs a [SLDisplayPresetDevice] from an objc.ID.
func SLDisplayPresetDeviceFromID(id objc.ID) SLDisplayPresetDevice {
	return SLDisplayPresetDevice{objectivec.Object{ID: id}}
}

// Ensure SLDisplayPresetDevice implements ISLDisplayPresetDevice.
var _ ISLDisplayPresetDevice = SLDisplayPresetDevice{}

// An interface definition for the [SLDisplayPresetDevice] class.
//
// # Methods
//
//   - [ISLDisplayPresetDevice.ActivePresetIndex]
//   - [ISLDisplayPresetDevice.ContainerId]
//   - [ISLDisplayPresetDevice.CopyCFContainerId]
//   - [ISLDisplayPresetDevice.CopyCalibrationInfo]
//   - [ISLDisplayPresetDevice.CopyCustomPresetInfo]
//   - [ISLDisplayPresetDevice.CopyPresetAtIndex]
//   - [ISLDisplayPresetDevice.CopyPresetDataAtIndex]
//   - [ISLDisplayPresetDevice.CopyUserAdjustment]
//   - [ISLDisplayPresetDevice.CopyUserAdjustmentData]
//   - [ISLDisplayPresetDevice.CopyUserAdjustmentForPreset]
//   - [ISLDisplayPresetDevice.FactoryDefaultPresetIndex]
//   - [ISLDisplayPresetDevice.FactoryResetWithType]
//   - [ISLDisplayPresetDevice.GetUserAdjustmentPowerLimit]
//   - [ISLDisplayPresetDevice.GetUserAdjustmentRangeWithInput]
//   - [ISLDisplayPresetDevice.InvalidateLiveUserAdjustmentForPreset]
//   - [ISLDisplayPresetDevice.InvalidateUserAdjustment]
//   - [ISLDisplayPresetDevice.InvalidateUserAdjustmentForPreset]
//   - [ISLDisplayPresetDevice.IsLiveUserAdjustmentSupported]
//   - [ISLDisplayPresetDevice.IsPerPresetUserAdjustmentSupported]
//   - [ISLDisplayPresetDevice.IsPresetValidAtIndex]
//   - [ISLDisplayPresetDevice.IsPresetWritableAtIndex]
//   - [ISLDisplayPresetDevice.IsUserAdjustmentSupported]
//   - [ISLDisplayPresetDevice.IsUserAdjustmentValid]
//   - [ISLDisplayPresetDevice.IsUserAdjustmentValidForAnyPreset]
//   - [ISLDisplayPresetDevice.IsUserAdjustmentValidForPreset]
//   - [ISLDisplayPresetDevice.PresetCapabilities]
//   - [ISLDisplayPresetDevice.PresetCount]
//   - [ISLDisplayPresetDevice.PresetUUIDAtIndexToBytes]
//   - [ISLDisplayPresetDevice.ResetPresetAtIndex]
//   - [ISLDisplayPresetDevice.ResetProController]
//   - [ISLDisplayPresetDevice.SetActivePresetIndex]
//   - [ISLDisplayPresetDevice.SetCustomPresetDataAtIndexWithData]
//   - [ISLDisplayPresetDevice.SetPresetAtIndexWithData]
//   - [ISLDisplayPresetDevice.SetPresetDataAtIndexWithData]
//   - [ISLDisplayPresetDevice.SetUserAdjustment]
//   - [ISLDisplayPresetDevice.SetUserAdjustmentData]
//   - [ISLDisplayPresetDevice.SetUserAdjustmentForPresetWithData]
//   - [ISLDisplayPresetDevice.InitWithService]
type ISLDisplayPresetDevice interface {
	objectivec.IObject

	// Topic: Methods

	ActivePresetIndex() uint32
	ContainerId() objectivec.IObject
	CopyCFContainerId() corefoundation.CFUUID
	CopyCalibrationInfo() objectivec.IObject
	CopyCustomPresetInfo() objectivec.IObject
	CopyPresetAtIndex(index uint32) objectivec.IObject
	CopyPresetDataAtIndex(index uint32) objectivec.IObject
	CopyUserAdjustment() objectivec.IObject
	CopyUserAdjustmentData() objectivec.IObject
	CopyUserAdjustmentForPreset(preset uint32) objectivec.IObject
	FactoryDefaultPresetIndex() uint32
	FactoryResetWithType(type_ byte) bool
	GetUserAdjustmentPowerLimit() float32
	GetUserAdjustmentRangeWithInput(range_ uint32, input unsafe.Pointer) unsafe.Pointer
	InvalidateLiveUserAdjustmentForPreset(preset uint32) bool
	InvalidateUserAdjustment() bool
	InvalidateUserAdjustmentForPreset(preset uint32) bool
	IsLiveUserAdjustmentSupported() bool
	IsPerPresetUserAdjustmentSupported() bool
	IsPresetValidAtIndex(index uint32) bool
	IsPresetWritableAtIndex(index uint32) bool
	IsUserAdjustmentSupported() bool
	IsUserAdjustmentValid() bool
	IsUserAdjustmentValidForAnyPreset() bool
	IsUserAdjustmentValidForPreset(preset uint32) bool
	PresetCapabilities() objectivec.IObject
	PresetCount() uint32
	PresetUUIDAtIndexToBytes(index uint32, bytes unsafe.Pointer) bool
	ResetPresetAtIndex(index uint32) bool
	ResetProController()
	SetActivePresetIndex(index uint32) bool
	SetCustomPresetDataAtIndexWithData(index uint32, data objectivec.IObject) bool
	SetPresetAtIndexWithData(index uint32, data objectivec.IObject) bool
	SetPresetDataAtIndexWithData(index uint32, data objectivec.IObject) bool
	SetUserAdjustment(adjustment objectivec.IObject) bool
	SetUserAdjustmentData(data objectivec.IObject) bool
	SetUserAdjustmentForPresetWithData(preset uint32, data objectivec.IObject) bool
	InitWithService(service uint32) SLDisplayPresetDevice
}

// Init initializes the instance.
func (s SLDisplayPresetDevice) Init() SLDisplayPresetDevice {
	rv := objc.SendIfResponds[SLDisplayPresetDevice](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLDisplayPresetDevice) Autorelease() SLDisplayPresetDevice {
	rv := objc.SendIfResponds[SLDisplayPresetDevice](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLDisplayPresetDevice creates a new SLDisplayPresetDevice instance.
func NewSLDisplayPresetDevice() SLDisplayPresetDevice {
	class := getSLDisplayPresetDeviceClass()
	rv := objc.SendIfResponds[SLDisplayPresetDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLDisplayPresetDeviceWithService(service uint32) SLDisplayPresetDevice {
	instance := getSLDisplayPresetDeviceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithService:"), service)
	return SLDisplayPresetDeviceFromID(rv)
}

func (s SLDisplayPresetDevice) ActivePresetIndex() uint32 {
	rv := objc.SendIfResponds[uint32](s.ID, objc.Sel("activePresetIndex"))
	return rv
}
func (s SLDisplayPresetDevice) ContainerId() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("containerId"))
	return objectivec.Object{ID: rv}
}
func (s SLDisplayPresetDevice) CopyCFContainerId() corefoundation.CFUUID {
	rv := objc.SendIfResponds[corefoundation.CFUUIDRef](s.ID, objc.Sel("copyCFContainerId"))
	return corefoundation.CFUUID(rv)
}
func (s SLDisplayPresetDevice) CopyCalibrationInfo() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("copyCalibrationInfo"))
	return objectivec.Object{ID: rv}
}
func (s SLDisplayPresetDevice) CopyCustomPresetInfo() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("copyCustomPresetInfo"))
	return objectivec.Object{ID: rv}
}
func (s SLDisplayPresetDevice) CopyPresetAtIndex(index uint32) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("copyPresetAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (s SLDisplayPresetDevice) CopyPresetDataAtIndex(index uint32) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("copyPresetDataAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (s SLDisplayPresetDevice) CopyUserAdjustment() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("copyUserAdjustment"))
	return objectivec.Object{ID: rv}
}
func (s SLDisplayPresetDevice) CopyUserAdjustmentData() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("copyUserAdjustmentData"))
	return objectivec.Object{ID: rv}
}
func (s SLDisplayPresetDevice) CopyUserAdjustmentForPreset(preset uint32) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("copyUserAdjustmentForPreset:"), preset)
	return objectivec.Object{ID: rv}
}
func (s SLDisplayPresetDevice) FactoryDefaultPresetIndex() uint32 {
	rv := objc.SendIfResponds[uint32](s.ID, objc.Sel("factoryDefaultPresetIndex"))
	return rv
}
func (s SLDisplayPresetDevice) FactoryResetWithType(type_ byte) bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("factoryResetWithType:"), type_)
	return rv
}
func (s SLDisplayPresetDevice) GetUserAdjustmentPowerLimit() float32 {
	rv := objc.SendIfResponds[float32](s.ID, objc.Sel("getUserAdjustmentPowerLimit"))
	return rv
}
func (s SLDisplayPresetDevice) GetUserAdjustmentRangeWithInput(range_ uint32, input unsafe.Pointer) unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](s.ID, objc.Sel("getUserAdjustmentRange:withInput:"), range_, input)
	return rv
}
func (s SLDisplayPresetDevice) InvalidateLiveUserAdjustmentForPreset(preset uint32) bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("invalidateLiveUserAdjustmentForPreset:"), preset)
	return rv
}
func (s SLDisplayPresetDevice) InvalidateUserAdjustment() bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("invalidateUserAdjustment"))
	return rv
}
func (s SLDisplayPresetDevice) InvalidateUserAdjustmentForPreset(preset uint32) bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("invalidateUserAdjustmentForPreset:"), preset)
	return rv
}
func (s SLDisplayPresetDevice) IsLiveUserAdjustmentSupported() bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("isLiveUserAdjustmentSupported"))
	return rv
}
func (s SLDisplayPresetDevice) IsPerPresetUserAdjustmentSupported() bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("isPerPresetUserAdjustmentSupported"))
	return rv
}
func (s SLDisplayPresetDevice) IsPresetValidAtIndex(index uint32) bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("isPresetValidAtIndex:"), index)
	return rv
}
func (s SLDisplayPresetDevice) IsPresetWritableAtIndex(index uint32) bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("isPresetWritableAtIndex:"), index)
	return rv
}
func (s SLDisplayPresetDevice) IsUserAdjustmentSupported() bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("isUserAdjustmentSupported"))
	return rv
}
func (s SLDisplayPresetDevice) IsUserAdjustmentValid() bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("isUserAdjustmentValid"))
	return rv
}
func (s SLDisplayPresetDevice) IsUserAdjustmentValidForAnyPreset() bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("isUserAdjustmentValidForAnyPreset"))
	return rv
}
func (s SLDisplayPresetDevice) IsUserAdjustmentValidForPreset(preset uint32) bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("isUserAdjustmentValidForPreset:"), preset)
	return rv
}
func (s SLDisplayPresetDevice) PresetCapabilities() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("presetCapabilities"))
	return objectivec.Object{ID: rv}
}
func (s SLDisplayPresetDevice) PresetCount() uint32 {
	rv := objc.SendIfResponds[uint32](s.ID, objc.Sel("presetCount"))
	return rv
}
func (s SLDisplayPresetDevice) PresetUUIDAtIndexToBytes(index uint32, bytes unsafe.Pointer) bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("presetUUIDAtIndex:toBytes:"), index, bytes)
	return rv
}
func (s SLDisplayPresetDevice) ResetPresetAtIndex(index uint32) bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("resetPresetAtIndex:"), index)
	return rv
}
func (s SLDisplayPresetDevice) ResetProController() {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("resetProController"))
}
func (s SLDisplayPresetDevice) SetActivePresetIndex(index uint32) bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("setActivePresetIndex:"), index)
	return rv
}
func (s SLDisplayPresetDevice) SetCustomPresetDataAtIndexWithData(index uint32, data objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("setCustomPresetDataAtIndex:withData:"), index, data)
	return rv
}
func (s SLDisplayPresetDevice) SetPresetAtIndexWithData(index uint32, data objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("setPresetAtIndex:withData:"), index, data)
	return rv
}
func (s SLDisplayPresetDevice) SetPresetDataAtIndexWithData(index uint32, data objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("setPresetDataAtIndex:withData:"), index, data)
	return rv
}
func (s SLDisplayPresetDevice) SetUserAdjustment(adjustment objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("setUserAdjustment:"), adjustment)
	return rv
}
func (s SLDisplayPresetDevice) SetUserAdjustmentData(data objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("setUserAdjustmentData:"), data)
	return rv
}
func (s SLDisplayPresetDevice) SetUserAdjustmentForPresetWithData(preset uint32, data objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("setUserAdjustmentForPreset:withData:"), preset, data)
	return rv
}
func (s SLDisplayPresetDevice) InitWithService(service uint32) SLDisplayPresetDevice {
	rv := objc.SendIfResponds[SLDisplayPresetDevice](s.ID, objc.Sel("initWithService:"), service)
	return rv
}

func (_SLDisplayPresetDeviceClass SLDisplayPresetDeviceClass) LuminanceCorrectionFactorForWhitepoint(whitepoint objectivec.IObject) float32 {
	rv := objc.SendIfResponds[float32](objc.ID(_SLDisplayPresetDeviceClass.class), objc.Sel("luminanceCorrectionFactorForWhitepoint:"), whitepoint)
	return rv
}
func (_SLDisplayPresetDeviceClass SLDisplayPresetDeviceClass) PresetDeviceWithService(service uint32) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_SLDisplayPresetDeviceClass.class), objc.Sel("presetDeviceWithService:"), service)
	return objectivec.Object{ID: rv}
}
func (_SLDisplayPresetDeviceClass SLDisplayPresetDeviceClass) UserAdjustmentLuminanceCorrectionFactorForWhitepoint(whitepoint objectivec.IObject) float32 {
	rv := objc.SendIfResponds[float32](objc.ID(_SLDisplayPresetDeviceClass.class), objc.Sel("userAdjustmentLuminanceCorrectionFactorForWhitepoint:"), whitepoint)
	return rv
}
