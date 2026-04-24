// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
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
	rv := objc.Send[SLDisplayPresetDevice](objc.ID(sc.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice
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
//
// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice
type ISLDisplayPresetDevice interface {
	objectivec.IObject

	// Topic: Methods

	ActivePresetIndex() uint32
	ContainerId() objectivec.IObject
	CopyCFContainerId() string
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
	GetUserAdjustmentRangeWithInput(range_ uint32, input objectivec.IObject) objectivec.IObject
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
	PresetUUIDAtIndexToBytes(index uint32, bytes objectivec.IObject) bool
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
	rv := objc.Send[SLDisplayPresetDevice](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLDisplayPresetDevice) Autorelease() SLDisplayPresetDevice {
	rv := objc.Send[SLDisplayPresetDevice](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLDisplayPresetDevice creates a new SLDisplayPresetDevice instance.
func NewSLDisplayPresetDevice() SLDisplayPresetDevice {
	class := getSLDisplayPresetDeviceClass()
	rv := objc.Send[SLDisplayPresetDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/initWithService:
func NewSLDisplayPresetDeviceWithService(service uint32) SLDisplayPresetDevice {
	instance := getSLDisplayPresetDeviceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithService:"), service)
	return SLDisplayPresetDeviceFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/activePresetIndex
func (s SLDisplayPresetDevice) ActivePresetIndex() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("activePresetIndex"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/containerId
func (s SLDisplayPresetDevice) ContainerId() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("containerId"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/copyCFContainerId
func (s SLDisplayPresetDevice) CopyCFContainerId() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("copyCFContainerId"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/copyCalibrationInfo
func (s SLDisplayPresetDevice) CopyCalibrationInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("copyCalibrationInfo"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/copyCustomPresetInfo
func (s SLDisplayPresetDevice) CopyCustomPresetInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("copyCustomPresetInfo"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/copyPresetAtIndex:
func (s SLDisplayPresetDevice) CopyPresetAtIndex(index uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("copyPresetAtIndex:"), index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/copyPresetDataAtIndex:
func (s SLDisplayPresetDevice) CopyPresetDataAtIndex(index uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("copyPresetDataAtIndex:"), index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/copyUserAdjustment
func (s SLDisplayPresetDevice) CopyUserAdjustment() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("copyUserAdjustment"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/copyUserAdjustmentData
func (s SLDisplayPresetDevice) CopyUserAdjustmentData() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("copyUserAdjustmentData"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/copyUserAdjustmentForPreset:
func (s SLDisplayPresetDevice) CopyUserAdjustmentForPreset(preset uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("copyUserAdjustmentForPreset:"), preset)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/factoryDefaultPresetIndex
func (s SLDisplayPresetDevice) FactoryDefaultPresetIndex() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("factoryDefaultPresetIndex"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/factoryResetWithType:
func (s SLDisplayPresetDevice) FactoryResetWithType(type_ byte) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("factoryResetWithType:"), type_)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/getUserAdjustmentPowerLimit
func (s SLDisplayPresetDevice) GetUserAdjustmentPowerLimit() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("getUserAdjustmentPowerLimit"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/getUserAdjustmentRange:withInput:
func (s SLDisplayPresetDevice) GetUserAdjustmentRangeWithInput(range_ uint32, input objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("getUserAdjustmentRange:withInput:"), range_, input)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/invalidateLiveUserAdjustmentForPreset:
func (s SLDisplayPresetDevice) InvalidateLiveUserAdjustmentForPreset(preset uint32) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("invalidateLiveUserAdjustmentForPreset:"), preset)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/invalidateUserAdjustment
func (s SLDisplayPresetDevice) InvalidateUserAdjustment() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("invalidateUserAdjustment"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/invalidateUserAdjustmentForPreset:
func (s SLDisplayPresetDevice) InvalidateUserAdjustmentForPreset(preset uint32) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("invalidateUserAdjustmentForPreset:"), preset)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/isLiveUserAdjustmentSupported
func (s SLDisplayPresetDevice) IsLiveUserAdjustmentSupported() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isLiveUserAdjustmentSupported"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/isPerPresetUserAdjustmentSupported
func (s SLDisplayPresetDevice) IsPerPresetUserAdjustmentSupported() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isPerPresetUserAdjustmentSupported"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/isPresetValidAtIndex:
func (s SLDisplayPresetDevice) IsPresetValidAtIndex(index uint32) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isPresetValidAtIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/isPresetWritableAtIndex:
func (s SLDisplayPresetDevice) IsPresetWritableAtIndex(index uint32) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isPresetWritableAtIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/isUserAdjustmentSupported
func (s SLDisplayPresetDevice) IsUserAdjustmentSupported() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isUserAdjustmentSupported"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/isUserAdjustmentValid
func (s SLDisplayPresetDevice) IsUserAdjustmentValid() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isUserAdjustmentValid"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/isUserAdjustmentValidForAnyPreset
func (s SLDisplayPresetDevice) IsUserAdjustmentValidForAnyPreset() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isUserAdjustmentValidForAnyPreset"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/isUserAdjustmentValidForPreset:
func (s SLDisplayPresetDevice) IsUserAdjustmentValidForPreset(preset uint32) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isUserAdjustmentValidForPreset:"), preset)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/presetCapabilities
func (s SLDisplayPresetDevice) PresetCapabilities() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("presetCapabilities"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/presetCount
func (s SLDisplayPresetDevice) PresetCount() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("presetCount"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/presetUUIDAtIndex:toBytes:
func (s SLDisplayPresetDevice) PresetUUIDAtIndexToBytes(index uint32, bytes objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("presetUUIDAtIndex:toBytes:"), index, bytes)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/resetPresetAtIndex:
func (s SLDisplayPresetDevice) ResetPresetAtIndex(index uint32) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("resetPresetAtIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/resetProController
func (s SLDisplayPresetDevice) ResetProController() {
	objc.Send[objc.ID](s.ID, objc.Sel("resetProController"))
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/setActivePresetIndex:
func (s SLDisplayPresetDevice) SetActivePresetIndex(index uint32) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("setActivePresetIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/setCustomPresetDataAtIndex:withData:
func (s SLDisplayPresetDevice) SetCustomPresetDataAtIndexWithData(index uint32, data objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("setCustomPresetDataAtIndex:withData:"), index, data)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/setPresetAtIndex:withData:
func (s SLDisplayPresetDevice) SetPresetAtIndexWithData(index uint32, data objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("setPresetAtIndex:withData:"), index, data)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/setPresetDataAtIndex:withData:
func (s SLDisplayPresetDevice) SetPresetDataAtIndexWithData(index uint32, data objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("setPresetDataAtIndex:withData:"), index, data)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/setUserAdjustment:
func (s SLDisplayPresetDevice) SetUserAdjustment(adjustment objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("setUserAdjustment:"), adjustment)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/setUserAdjustmentData:
func (s SLDisplayPresetDevice) SetUserAdjustmentData(data objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("setUserAdjustmentData:"), data)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/setUserAdjustmentForPreset:withData:
func (s SLDisplayPresetDevice) SetUserAdjustmentForPresetWithData(preset uint32, data objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("setUserAdjustmentForPreset:withData:"), preset, data)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/initWithService:
func (s SLDisplayPresetDevice) InitWithService(service uint32) SLDisplayPresetDevice {
	rv := objc.Send[SLDisplayPresetDevice](s.ID, objc.Sel("initWithService:"), service)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/luminanceCorrectionFactorForWhitepoint:
func (_SLDisplayPresetDeviceClass SLDisplayPresetDeviceClass) LuminanceCorrectionFactorForWhitepoint(whitepoint objectivec.IObject) float32 {
	rv := objc.Send[float32](objc.ID(_SLDisplayPresetDeviceClass.class), objc.Sel("luminanceCorrectionFactorForWhitepoint:"), whitepoint)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/presetDeviceWithService:
func (_SLDisplayPresetDeviceClass SLDisplayPresetDeviceClass) PresetDeviceWithService(service uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLDisplayPresetDeviceClass.class), objc.Sel("presetDeviceWithService:"), service)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDisplayPresetDevice/userAdjustmentLuminanceCorrectionFactorForWhitepoint:
func (_SLDisplayPresetDeviceClass SLDisplayPresetDeviceClass) UserAdjustmentLuminanceCorrectionFactorForWhitepoint(whitepoint objectivec.IObject) float32 {
	rv := objc.Send[float32](objc.ID(_SLDisplayPresetDeviceClass.class), objc.Sel("userAdjustmentLuminanceCorrectionFactorForWhitepoint:"), whitepoint)
	return rv
}
