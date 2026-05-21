// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"
)

// C struct types

// CGColor
type CGColor struct {
}

// CGDisplayStream
type CGDisplayStream struct {
}

// CGDisplayStreamUpdate
type CGDisplayStreamUpdate struct {
}

// CGSRegionObject
type CGSRegionObject struct {
}

// CGXConnection
type CGXConnection struct {
}

// CGXSession
type CGXSession struct {
}

// CGXSessionProcessData
type CGXSessionProcessData struct {
}

// CPSProcessRec
type CPSProcessRec struct {
}

// CPSProcessSerNum
type CPSProcessSerNum struct {
}

// CPXEventProcessorContext
type CPXEventProcessorContext struct {
}

// IONotificationPort
type IONotificationPort struct {
}

// MessageInitData
type MessageInitData struct {
}

// OpaqueCUIRendererRef
type OpaqueCUIRendererRef struct {
}

// PresetDeviceFlags
type PresetDeviceFlags struct {
}

// ProDisplayController
type ProDisplayController struct {
	_load_legacy_preset_data           unsafe.Pointer
	_erase_legacy_preset_data          unsafe.Pointer
	_load_legacy_user_adjustment_data  unsafe.Pointer
	_erase_legacy_user_adjustment_data unsafe.Pointer
	_preset_update_callback            unsafe.Pointer
	_ua_update_callback                unsafe.Pointer
	_shim                              unsafe.Pointer
}

// ProcessSerialNumber
type ProcessSerialNumber struct {
	HighLongOfPSN uint32
	LowLongOfPSN  uint32
}

// SLSBrightnessPolicyTxState
type SLSBrightnessPolicyTxState struct {
	Shielding_policy uint8
	Dim_policy       uint8
	Sleep_policy     uint8
	Mask             uint
}

// SLSBrightnessTimeoutTxState
type SLSBrightnessTimeoutTxState struct {
	Shielding_timeout float64
	Dim_timeout       float64
	Sleep_timeout     float64
	Mask              uint
}

// SLSBrightnessTxState
type SLSBrightnessTxState struct {
	Ambient                    float32
	Filtered_ambient           float32
	Sdr_brightness             float32
	Brightness_limit           float32
	Headroom                   float32
	Potential_headroom         float32
	Reference_headroom         float32
	Contrast_preservation      float32
	Low_ambient_adaptation     float32
	High_ambient_adaptation    float32
	Indicator_brightness       float32
	Indicator_brightness_limit float32
	Contrast_enhancer          float32
	Mask                       uint
}

// SLSEventRecord
type SLSEventRecord struct {
}

// SLSTransaction
type SLSTransaction struct {
}

// WSMainThreadBlockHoist
type WSMainThreadBlockHoist struct {
}

// WSStructuralRegion
type WSStructuralRegion struct {
}

// Cfuuid
type Cfuuid struct {
}

// CGEvent
type CGEvent struct {
}

// IOHIDEvent
type IOHIDEvent struct {
}

// Lsasn
type Lsasn struct {
}
