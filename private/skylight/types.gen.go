// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objectivec"
)

// C struct types

// CGColor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/CGColor
type CGColor struct {
}

// CGDisplayStream
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/CGDisplayStream
type CGDisplayStream struct {
}

// CGDisplayStreamUpdate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/CGDisplayStreamUpdate
type CGDisplayStreamUpdate struct {
}

// CGSRegionObject
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/CGSRegionObject
type CGSRegionObject struct {
}

// CGXConnection
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/CGXConnection
type CGXConnection struct {
}

// CGXSession
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/CGXSession
type CGXSession struct {
}

// CGXSessionProcessData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/CGXSessionProcessData
type CGXSessionProcessData struct {
}

// CPSProcessRec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/CPSProcessRec
type CPSProcessRec struct {
}

// CPSProcessSerNum
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/CPSProcessSerNum
type CPSProcessSerNum struct {
}

// CPXEventProcessorContext
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/CPXEventProcessorContext
type CPXEventProcessorContext struct {
}

// IONotificationPort
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/IONotificationPort
type IONotificationPort struct {
}

// MessageInitData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/MessageInitData
type MessageInitData struct {
}

// OpaqueCUIRendererRef
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/OpaqueCUIRendererRef
type OpaqueCUIRendererRef struct {
}

// PresetDeviceFlags
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/PresetDeviceFlags
type PresetDeviceFlags struct {
}

// ProDisplayController
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/ProDisplayController
type ProDisplayController struct {
	_load_legacy_preset_data           objectivec.IObject
	_erase_legacy_preset_data          unsafe.Pointer
	_load_legacy_user_adjustment_data  objectivec.IObject
	_erase_legacy_user_adjustment_data unsafe.Pointer
	_preset_update_callback            unsafe.Pointer
	_ua_update_callback                unsafe.Pointer
	_shim                              unsafe.Pointer
}

// ProcessSerialNumber
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/ProcessSerialNumber
type ProcessSerialNumber struct {
	HighLongOfPSN uint32
	LowLongOfPSN  uint32
}

// SLSBrightnessPolicyTxState
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/SLSBrightnessPolicyTxState
type SLSBrightnessPolicyTxState struct {
	Shielding_policy uint8
	Dim_policy       uint8
	Sleep_policy     uint8
	Mask             uint
}

// SLSBrightnessTimeoutTxState
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/SLSBrightnessTimeoutTxState
type SLSBrightnessTimeoutTxState struct {
	Shielding_timeout float64
	Dim_timeout       float64
	Sleep_timeout     float64
	Mask              uint
}

// SLSBrightnessTxState
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/SLSBrightnessTxState
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
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/_SLSEventRecord
type SLSEventRecord struct {
}

// SLSTransaction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/_SLSTransaction
type SLSTransaction struct {
}

// WSMainThreadBlockHoist
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/WSMainThreadBlockHoist
type WSMainThreadBlockHoist struct {
}

// WSStructuralRegion
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/WSStructuralRegion
type WSStructuralRegion struct {
}

// _CFUUID
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/__CFUUID
type _CFUUID struct {
}

// _CGEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/__CGEvent
type _CGEvent struct {
}

// _IOHIDEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/__IOHIDEvent
type _IOHIDEvent struct {
}

// _LSASN
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SkyLight/__LSASN
type _LSASN struct {
}
