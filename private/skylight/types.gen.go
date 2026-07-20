// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
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
	Field1   CGXConnectionRef
	Field2   CGXConnectionRef
	Field3   uint
	Field4   unsafe.Pointer
	Field5   uint
	Field6   CGXSessionRef
	Field7   int
	Field8   unsafe.Pointer
	Field9   unsafe.Pointer
	Field10  objectivec.Object
	Field11  objectivec.Object
	Field12  objectivec.Object
	Field13  objectivec.Object
	Field14  objectivec.Object
	Field15  objectivec.Object
	Field16  objectivec.Object
	Field17  objectivec.Object
	Field18  objectivec.Object
	Field19  objectivec.Object
	Field20  objectivec.Object
	Field21  objectivec.Object
	Field22  objectivec.Object
	Field23  objectivec.Object
	Field24  objectivec.Object
	Field25  objectivec.Object
	Field26  objectivec.Object
	Field27  objectivec.Object
	Field28  objectivec.Object
	Field29  objectivec.Object
	Field30  objectivec.Object
	Field31  objectivec.Object
	Field32  objectivec.Object
	Field33  objectivec.Object
	Field34  objectivec.Object
	Field35  objectivec.Object
	Field36  objectivec.Object
	Field37  objectivec.Object
	Field38  objectivec.Object
	Field39  objectivec.Object
	Field40  unsafe.Pointer
	Field41  uint
	Field42  corefoundation.CFDictionaryRef
	Field43  CGSRegionObjectRef
	Field44  uint
	Field45  uint
	Field46  uint
	Field47  float64
	Field48  float64
	Field49  *uint64
	Field50  uint64
	Field51  uint64
	Field52  uint
	Field53  uint
	Field54  CGSRegionObjectRef
	Field55  CGXCursorRef
	Field56  corefoundation.CFDictionaryRef
	Field57  unsafe.Pointer
	Field58  uint
	Field59  uint
	Field60  uint
	Field61  int
	Field62  uint
	Field63  CPSProcessSerNum
	Field64  uint
	Field65  uint
	Field66  uint
	Field67  objectivec.Object
	Field68  objectivec.Object
	Field69  objectivec.Object
	Field70  objectivec.Object
	Field71  objectivec.Object
	Field72  objectivec.Object
	Field73  objectivec.Object
	Field74  objectivec.Object
	Field75  objectivec.Object
	Field76  objectivec.Object
	Field77  objectivec.Object
	Field78  objectivec.Object
	Field79  objectivec.Object
	Field80  CGXEventTapRef
	Field81  CGXCaptureStateRef
	Field82  uint
	Field83  unsafe.Pointer
	Field84  unsafe.Pointer
	Field85  float64
	Field86  float64
	Field87  float64
	Field88  float64
	Field89  float64
	Field90  uint
	Field91  CGXDirtyScreenStateRef
	Field92  uint
	Field93  bool
	Field94  PKGSpaceRef
	Field95  uint
	Field96  int
	Field97  uint64
	Field98  XListStructRef
	Field99  objectivec.Object
	Field100 objectivec.Object
	Field101 objectivec.Object
	Field102 uint
	Field103 float64
	Field104 uint
	Field105 float64
	Field106 SLSStructuralRegionIDRangeRef
	Field107 unsafe.Pointer
	Field108 XListStructRef
	Field109 CGSRegionObjectRef
	Field110 objectivec.Object
	Field111 objectivec.Object
	Field112 uint
	Field113 bool
	Field114 bool
	Field115 uint64
	Field116 unsafe.Pointer
	Field117 CGSRegionObjectRef
	Field118 bool
	Field119 CGSRegionObjectRef
	Field120 objectivec.Object
	Field121 objectivec.Object
	Field122 objectivec.Object
	Field123 uint64
	Field124 objectivec.Object
	Field125 objectivec.Object
	Field126 int
	Field127 int
	Field128 CGXConnectionCAContextTrackingStateRef
	Field129 CGXConnectionBoxRef
	Field130 bool
	Field131 float64
	Field132 corefoundation.CFStringRef
	Field133 objectivec.Object
	Field134 objectivec.Object
	Field135 objectivec.Object
	Field136 objectivec.Object
	Field137 objectivec.Object
	Field138 objectivec.Object
	Field139 objectivec.Object
	Field140 objectivec.Object
	Field141 objectivec.Object
	Field142 objectivec.Object
	Field143 objectivec.Object
	Field144 objectivec.Object
	Field145 objectivec.Object
	Field146 unsafe.Pointer
	Field147 uint64
	Field148 XListStructRef
	Field149 uint
	Field150 objectivec.Object
	Field151 int
	Field152 objectivec.Object
}

// CGXSession
type CGXSession struct {
	Field1  uint
	Field2  int
	Field3  uint
	Field4  uint
	Field5  uint
	Field6  int
	Field7  unsafe.Pointer
	Field8  unsafe.Pointer
	Field9  bool
	Field10 *byte
	Field11 *byte
	Field12 *byte
	Field13 bool
	Field14 bool
	Field15 bool
	Field16 objectivec.Object
	Field17 uint
	Field18 bool
	Field19 bool
	Field20 bool
	Field21 bool
	Field22 uint
	Field23 uint
	Field24 corefoundation.CFDictionaryRef
	Field25 corefoundation.CFDictionaryRef
	Field26 corefoundation.CFDictionaryRef
	Field27 uint
	Field28 uint
	Field29 bool
	Field30 CGXSessionWindowDataRef
	Field31 ZoomManagerRef
	Field32 CGXSessionDisplayZoomDataRef
	Field33 WSSessionWorkspaceDataRef
	Field34 CGXSessionConnectionDataRef
	Field35 CGXSessionConfigurationDataRef
	Field36 WSSessionDisplayPreferencesDataRef
	Field37 CGXSessionDFRDataRef
	Field38 CGXSessionEventDataRef
	Field39 CGXSessionPackageDataRef
	Field40 CGXSessionProcessDataRef
	Field41 CGXSessionDisplayDataRef
	Field42 SessionDataRef
	Field43 WSSessionDisplayConfigDataRef
	Field44 WSCursorDataRef
	Field45 CGXSessionDisplayStreamDataRef
	Field46 WSMessageTraceSessionDataRef
	Field47 WSSessionStructuralRegionDataRef
	Field48 CGXSessionHMDDataRef
	Field49 WSSessionDisplayUpdateDataRef
	Field50 DesktopEffectsSessionDataRef
	Field51 SchedulerSessionDataRef
	Field52 unsafe.Pointer
	Field53 uint64
	Field54 WSSessionCaptureDataRef
	Field55 int
	Field56 int
	Field57 bool
	Field58 CGXWindowRef
	Field59 CGXWindowRef
	Field60 CGXWindowRef
	Field61 uint
	Field62 bool
	Field63 uint
	Field64 float64
	Field65 bool
	Field66 bool
	Field67 CGXSessionWindowOverrideResolutionDataRef
	Field68 CGXSessionWindowPixelDimensionsHintDataRef
	Field69 CGXSessionUtilityDisplayControllerDataRef
	Field70 int
	Field71 int
}

// CGXSessionProcessData
type CGXSessionProcessData struct {
	Field1  objectivec.Object
	Field2  objectivec.Object
	Field3  objectivec.Object
	Field4  objectivec.Object
	Field5  objectivec.Object
	Field6  objectivec.Object
	Field7  objectivec.Object
	Field8  objectivec.Object
	Field9  objectivec.Object
	Field10 CPSProcessRecRef
	Field11 unsafe.Pointer
	Field12 CPSProcessRecRef
	Field13 CPSProcessSerNum
	Field14 CPSKeyFocusInfoRecRef
	Field15 int
	Field16 uint
	Field17 bool
	Field18 unsafe.Pointer
	Field19 uint
	Field20 unsafe.Pointer
}

// CPSProcessRec
type CPSProcessRec struct {
	Field1  uint
	Field2  CPSProcessSerNum
	Field3  int
	Field4  CPSProcessRecRef
	Field5  CPSProcessRecRef
	Field6  CPSProcessRecRef
	Field7  int
	Field8  bool
	Field9  uint
	Field10 *byte
	Field11 unsafe.Pointer
	Field12 uint
	Field13 uint
	Field14 bool
	Field15 bool
	Field16 int
	Field17 LSASNRef
	Field18 uint8
	Field19 int
	Field20 bool
	Field21 uint
}

// CPSProcessSerNum
type CPSProcessSerNum struct {
	Field1 uint
	Field2 uint
}

// CPXEventProcessorContext
type CPXEventProcessorContext struct {
	Field1 CGXSessionRef
	Field2 CGXSessionProcessDataRef
	Field3 CPSProcessRecRef
}

// IONotificationPort
type IONotificationPort struct {
}

// MessageInitData
type MessageInitData struct {
	Field1 unsafe.Pointer
}

// OpaqueCUIRendererRef
type OpaqueCUIRendererRef struct {
}

// PresetDeviceFlags
type PresetDeviceFlags struct {
	Field1 bool
	Field2 bool
	Field3 bool
	Field4 bool
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
	Field1  uint16
	Field2  uint16
	Field3  uint
	Field4  uint
	Field5  corefoundation.CGPoint
	Field6  corefoundation.CGPoint
	Field7  uint64
	Field8  uint
	Field9  uint
	Field10 uint
	Field11 unsafe.Pointer
	Field12 unsafe.Pointer
	Field13 unsafe.Pointer
	Field14 unsafe.Pointer
	Field15 uint16
	Field16 uint16
	Field17 CGSEventAppendixRef
	Field18 uint
	Field19 bool
	Field20 corefoundation.CFDataRef
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
	Field1 unsafe.Pointer
	Field2 uint
	Field3 SLSEventRecordRef
}

// IOHIDEvent
type IOHIDEvent struct {
}

// Lsasn
type Lsasn struct {
}
