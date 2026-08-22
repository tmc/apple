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

// CGEventProcess
type CGEventProcess struct {
	Field1 int32
	Field2 uint32
	Field3 uint32
	Field4 uint32
	Field5 uint32
}

// CGGestureData
type CGGestureData struct {
	Field1  uint32
	Field2  uint64
	Field3  bool
	Field4  bool
	Field5  uint8
	Field6  uint8
	Field7  uint32
	Field8  float32
	Field9  uint16
	Field10 uint8
	Field11 unsafe.Pointer
}

// CGSRegionObject
type CGSRegionObject struct {
}

// CGSTabletPointData
type CGSTabletPointData struct {
	Field1  int32
	Field2  int32
	Field3  int32
	Field4  uint16
	Field5  uint16
	Field6  [2]uint16
	Field7  uint16
	Field8  int16
	Field9  uint16
	Field10 int16
	Field11 int16
	Field12 int16
}

// CGSTabletProximityData
type CGSTabletProximityData struct {
	Field1  uint16
	Field2  uint16
	Field3  uint16
	Field4  uint16
	Field5  uint16
	Field6  uint16
	Field7  uint32
	Field8  uint64
	Field9  uint32
	Field10 uint8
	Field11 uint8
	Field12 int16
}

// CGXCaptureState
type CGXCaptureState struct {
}

// CGXConnection
type CGXConnection struct {
	Field1   CGXConnectionRef
	Field2   CGXConnectionRef
	Field3   uint32
	Field4   WSConnectionDatagramInfo
	Field5   uint32
	Field6   CGXSessionRef
	Field7   int32
	Field8   CGXConnectionNotificationContext
	Field9   CGXConnectionNotificationContext
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
	Field41  uint32
	Field42  corefoundation.CFDictionaryRef
	Field43  CGSRegionObjectRef
	Field44  uint32
	Field45  uint32
	Field46  uint32
	Field47  float64
	Field48  float64
	Field49  *uint64
	Field50  uint64
	Field51  uint64
	Field52  uint32
	Field53  uint32
	Field54  CGSRegionObjectRef
	Field55  CGXCursorRef
	Field56  corefoundation.CFDictionaryRef
	Field57  unsafe.Pointer
	Field58  uint32
	Field59  uint32
	Field60  uint32
	Field61  int32
	Field62  uint32
	Field63  CPSProcessSerNum
	Field64  uint32
	Field65  uint32
	Field66  uint32
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
	Field82  uint32
	Field83  unsafe.Pointer
	Field84  unsafe.Pointer
	Field85  float64
	Field86  float64
	Field87  float64
	Field88  float64
	Field89  float64
	Field90  uint32
	Field91  CGXDirtyScreenStateRef
	Field92  uint32
	Field93  bool
	Field94  PKGSpaceRef
	Field95  WSSymbolicHotKeyBitMask
	Field96  int32
	Field97  uint64
	Field98  XListStructRef
	Field99  objectivec.Object
	Field100 objectivec.Object
	Field101 objectivec.Object
	Field102 uint32
	Field103 float64
	Field104 uint32
	Field105 float64
	Field106 SLSStructuralRegionIDRangeRef
	Field107 [13]uint32
	Field108 XListStructRef
	Field109 CGSRegionObjectRef
	Field110 objectivec.Object
	Field111 objectivec.Object
	Field112 uint32
	Field113 bool
	Field114 bool
	Field115 uint64
	Field116 [8]uint32
	Field117 CGSRegionObjectRef
	Field118 bool
	Field119 CGSRegionObjectRef
	Field120 objectivec.Object
	Field121 objectivec.Object
	Field122 objectivec.Object
	Field123 uint64
	Field124 objectivec.Object
	Field125 objectivec.Object
	Field126 int32
	Field127 int32
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
	Field149 uint32
	Field150 objectivec.Object
	Field151 int32
	Field152 objectivec.Object
}

// CGXConnectionCAContextTrackingState
type CGXConnectionCAContextTrackingState struct {
}

// CGXConnectionNotice
type CGXConnectionNotice struct {
}

// CGXConnectionNotificationContext
type CGXConnectionNotificationContext struct {
	Field1 uint32
	Field2 uint32
	Field3 CGXConnectionNoticeRef
}

// CGXCredentials
type CGXCredentials struct {
	Field1 uint32
	Field2 uint32
}

// CGXCursor
type CGXCursor struct {
}

// CGXDirtyScreenState
type CGXDirtyScreenState struct {
}

// CGXEventTap
type CGXEventTap struct {
}

// CGXSession
type CGXSession struct {
	Field1  uint32
	Field2  int32
	Field3  uint32
	Field4  uint32
	Field5  uint32
	Field6  int32
	Field7  CGXCredentials
	Field8  CGXCredentials
	Field9  bool
	Field10 *byte
	Field11 *byte
	Field12 *byte
	Field13 bool
	Field14 bool
	Field15 bool
	Field16 objectivec.Object
	Field17 uint32
	Field18 bool
	Field19 bool
	Field20 bool
	Field21 bool
	Field22 uint32
	Field23 uint32
	Field24 corefoundation.CFDictionaryRef
	Field25 corefoundation.CFDictionaryRef
	Field26 corefoundation.CFDictionaryRef
	Field27 uint32
	Field28 uint32
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
	Field55 int32
	Field56 int32
	Field57 bool
	Field58 CGXWindowRef
	Field59 CGXWindowRef
	Field60 CGXWindowRef
	Field61 uint32
	Field62 bool
	Field63 uint32
	Field64 float64
	Field65 bool
	Field66 bool
	Field67 CGXSessionWindowOverrideResolutionDataRef
	Field68 CGXSessionWindowPixelDimensionsHintDataRef
	Field69 CGXSessionUtilityDisplayControllerDataRef
	Field70 int32
	Field71 int32
}

// CGXSessionConfigurationData
type CGXSessionConfigurationData struct {
}

// CGXSessionConnectionData
type CGXSessionConnectionData struct {
}

// CGXSessionDFRData
type CGXSessionDFRData struct {
}

// CGXSessionDisplayData
type CGXSessionDisplayData struct {
}

// CGXSessionDisplayStreamData
type CGXSessionDisplayStreamData struct {
}

// CGXSessionDisplayZoomData
type CGXSessionDisplayZoomData struct {
}

// CGXSessionEventData
type CGXSessionEventData struct {
}

// CGXSessionHMDData
type CGXSessionHMDData struct {
}

// CGXSessionPackageData
type CGXSessionPackageData struct {
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
	Field11 [32]CPSProcessRecRef
	Field12 CPSProcessRecRef
	Field13 CPSProcessSerNum
	Field14 CPSKeyFocusInfoRecRef
	Field15 int32
	Field16 uint32
	Field17 bool
	Field18 unsafe.Pointer
	Field19 uint32
	Field20 unsafe.Pointer
}

// CGXSessionUtilityDisplayControllerData
type CGXSessionUtilityDisplayControllerData struct {
}

// CGXSessionWindowData
type CGXSessionWindowData struct {
}

// CGXSessionWindowOverrideResolutionData
type CGXSessionWindowOverrideResolutionData struct {
}

// CGXSessionWindowPixelDimensionsHintData
type CGXSessionWindowPixelDimensionsHintData struct {
}

// CGXWindow
type CGXWindow struct {
}

// CPSKeyFocusInfoRec
type CPSKeyFocusInfoRec struct {
}

// CPSProcessRec
type CPSProcessRec struct {
	Field1  uint32
	Field2  CPSProcessSerNum
	Field3  int32
	Field4  CPSProcessRecRef
	Field5  CPSProcessRecRef
	Field6  CPSProcessRecRef
	Field7  int32
	Field8  bool
	Field9  uint32
	Field10 *byte
	Field11 unsafe.Pointer
	Field12 uint32
	Field13 uint32
	Field14 bool
	Field15 bool
	Field16 int32
	Field17 LSASNRef
	Field18 uint8
	Field19 int32
	Field20 bool
	Field21 uint32
}

// CPSProcessSerNum
type CPSProcessSerNum struct {
	Field1 uint32
	Field2 uint32
}

// CPXEventProcessorContext
type CPXEventProcessorContext struct {
	Field1 CGXSessionRef
	Field2 CGXSessionProcessDataRef
	Field3 CPSProcessRecRef
}

// DesktopEffectsSessionData
type DesktopEffectsSessionData struct {
}

// Digester
type Digester struct {
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

// PKGSpace
type PKGSpace struct {
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
	_load_legacy_preset_data           [4]uint64
	_erase_legacy_preset_data          [4]uint64
	_load_legacy_user_adjustment_data  [4]uint64
	_erase_legacy_user_adjustment_data [4]uint64
	_preset_update_callback            [4]uint64
	_ua_update_callback                [4]uint64
	_shim                              [1]uint64
}

// ProDisplayLibraryShim
type ProDisplayLibraryShim struct {
}

// SLSBrightnessPolicyTxState
type SLSBrightnessPolicyTxState struct {
	Shielding_policy uint8
	Dim_policy       uint8
	Sleep_policy     uint8
	Mask             uint32
}

// SLSBrightnessTimeoutTxState
type SLSBrightnessTimeoutTxState struct {
	Shielding_timeout float64
	Dim_timeout       float64
	Sleep_timeout     float64
	Mask              uint32
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
	Mask                       uint32
}

// SLSEventRecord
type SLSEventRecord struct {
	Field1  uint16
	Field2  uint16
	Field3  uint32
	Field4  uint32
	Field5  corefoundation.CGPoint
	Field6  corefoundation.CGPoint
	Field7  uint64
	Field8  uint32
	Field9  uint32
	Field10 uint32
	Field11 CGEventSourceData
	Field12 CGEventProcess
	Field13 unsafe.Pointer
	Field14 unsafe.Pointer
	Field15 uint16
	Field16 uint16
	Field17 CGSEventAppendixRef
	Field18 uint32
	Field19 bool
	Field20 corefoundation.CFDataRef
}

// SLSEventRecord248
type SLSEventRecord248 struct {
	MajorVersion    uint16
	MinorVersion    uint16
	DeclaredLength  uint32
	EventType       uint32
	SubtypeFlags    uint32
	LocationX       float64
	LocationY       float64
	WinLocationX    float64
	WinLocationY    float64
	EventTime       uint64
	EventFlags      uint32
	WindowID        uint32
	ConnectionID    uint32
	Pad0x44         [70]uint8
	ActivationState uint8
	Pad0x8b         [85]uint8
	AppendixPtr     uintptr
	Pad0xe8         [16]uint8
}

// SLSScreenTelemetryResultsSnapshotData
type SLSScreenTelemetryResultsSnapshotData struct {
}

// SLSScreenTelemetryResultsSnapshotPanelData
type SLSScreenTelemetryResultsSnapshotPanelData struct {
}

// SLSScreenTelemetryResultsSnapshotZoneData
type SLSScreenTelemetryResultsSnapshotZoneData struct {
}

// SLSScreenTelemetryResultsSnapshotZoneRowData
type SLSScreenTelemetryResultsSnapshotZoneRowData struct {
}

// SLSStructuralRegionIDRange
type SLSStructuralRegionIDRange struct {
}

// SLSTransaction
type SLSTransaction struct {
}

// SLSeedResolver
type SLSeedResolver struct {
}

// SchedulerSessionData
type SchedulerSessionData struct {
}

// SessionData
type SessionData struct {
}

// SpecialKeyState
type SpecialKeyState struct {
	KeyCode     uint32
	Modifiers   uint32
	Registrant  CPSProcessRecRef
	ConnID      uint32
	KeysPending uint32
}

// WSConnectionDatagramInfo
type WSConnectionDatagramInfo struct {
	Field1  uint32
	Field2  float64
	Field3  float64
	Field4  WSDatagramWriteStreamRef
	Field5  WSNotifyInterestSetRef
	Field6  WSNotifyInterestSetRef
	Field7  int32
	Field8  bool
	Field9  bool
	Field10 bool
	Field11 bool
}

// WSCursorData
type WSCursorData struct {
}

// WSDatagramWriteStream
type WSDatagramWriteStream struct {
}

// WSDisplaySessionData
type WSDisplaySessionData struct {
}

// WSMainThreadBlockHoist
type WSMainThreadBlockHoist struct {
}

// WSMessageTraceSessionData
type WSMessageTraceSessionData struct {
}

// WSNotifyInterestSet
type WSNotifyInterestSet struct {
}

// WSSessionCaptureData
type WSSessionCaptureData struct {
}

// WSSessionDisplayConfigData
type WSSessionDisplayConfigData struct {
}

// WSSessionDisplayPreferencesData
type WSSessionDisplayPreferencesData struct {
}

// WSSessionDisplayUpdateData
type WSSessionDisplayUpdateData struct {
}

// WSSessionStructuralRegionData
type WSSessionStructuralRegionData struct {
}

// WSSessionWorkspaceData
type WSSessionWorkspaceData struct {
}

// WSStructuralRegion
type WSStructuralRegion struct {
}

// WSSymbolicHotKeyBitMask
type WSSymbolicHotKeyBitMask struct {
	Field1 [9]uint32
}

// ZoomManager
type ZoomManager struct {
}

// CFData
type CFData struct {
}

// CFDictionary
type CFDictionary struct {
}

// CFRuntimeBase
type CFRuntimeBase struct {
	Field1 uint64
	Field2 uint64
}

// CFString
type CFString struct {
}

// Cfuuid
type Cfuuid struct {
}

// CGEvent
type CGEvent struct {
	Field1 CFRuntimeBase
	Field2 uint32
	Field3 SLSEventRecordRef
}

// CGEventSourceData
type CGEventSourceData struct {
	Field1  int32
	Field2  uint32
	Field3  uint32
	Field4  uint32
	Field5  uint64
	Field6  uint32
	Field7  uint16
	Field8  uint8
	Field9  uint8
	Field10 uint64
}

// CGSEventAppendix
type CGSEventAppendix struct {
}

// IOHIDEvent
type IOHIDEvent struct {
}

// Lsasn
type Lsasn struct {
}

// SharedWeakCount
type SharedWeakCount struct {
}

// OSUnfairLockS
type OSUnfairLockS struct {
	_os_unfair_lock_opaque uint32
}

// Os_unfair_lock_s is a type alias for OSUnfairLockS for use in objc.Send[T] calls.
type Os_unfair_lock_s = OSUnfairLockS

// Type
type Type struct {
	__data [24]uint8
}

// XListStruct
type XListStruct struct {
}

// X_list_struct is a type alias for XListStruct for use in objc.Send[T] calls.
type X_list_struct = XListStruct
