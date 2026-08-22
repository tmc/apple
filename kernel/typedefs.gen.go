// Code generated from Apple documentation. DO NOT EDIT.

package kernel

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// See: https://developer.apple.com/documentation/kernel/ataoperationtype
type ATAOperationType = uint32

// See: https://developer.apple.com/documentation/kernel/atapiclientdata
// ATAPIClientData is an unresolved C aggregate typedef.
type ATAPIClientData unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/atapicmdpacket
// ATAPICmdPacket is opaque storage with the size and alignment C gives ATAPICmdPacket:
// 18 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 18 into.
type ATAPICmdPacket [9]uint16

// See: https://developer.apple.com/documentation/kernel/atarequestidentifier
type ATARequestIdentifier = uintptr

// See: https://developer.apple.com/documentation/kernel/avcconnecttargetplugsinparams
// AVCConnectTargetPlugsInParams is opaque storage with the size and alignment C gives AVCConnectTargetPlugsInParams:
// 28 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 28 into.
type AVCConnectTargetPlugsInParams [7]uint32

// See: https://developer.apple.com/documentation/kernel/avcconnecttargetplugsoutparams
// AVCConnectTargetPlugsOutParams is opaque storage with the size and alignment C gives AVCConnectTargetPlugsOutParams:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type AVCConnectTargetPlugsOutParams [2]uint32

// See: https://developer.apple.com/documentation/kernel/avcgettargetplugconnectioninparams
// AVCGetTargetPlugConnectionInParams is opaque storage with the size and alignment C gives AVCGetTargetPlugConnectionInParams:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type AVCGetTargetPlugConnectionInParams [3]uint32

// See: https://developer.apple.com/documentation/kernel/avcgettargetplugconnectionoutparams
// AVCGetTargetPlugConnectionOutParams is opaque storage with the size and alignment C gives AVCGetTargetPlugConnectionOutParams:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type AVCGetTargetPlugConnectionOutParams [4]uint32

// See: https://developer.apple.com/documentation/kernel/avcsubunitplugrecord
// AVCSubunitPlugRecord is opaque storage with the size and alignment C gives AVCSubunitPlugRecord:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type AVCSubunitPlugRecord [2]uint32

// See: https://developer.apple.com/documentation/kernel/avcunitplugrecord
// AVCUnitPlugRecord is opaque storage with the size and alignment C gives AVCUnitPlugRecord:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type AVCUnitPlugRecord [1]uint32

// See: https://developer.apple.com/documentation/kernel/avcunitplugs
// AVCUnitPlugs is opaque storage with the size and alignment C gives AVCUnitPlugs:
// 512 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 512 into.
type AVCUnitPlugs [128]uint32

// See: https://developer.apple.com/documentation/kernel/avidtype
type AVIDType = uint32

// See: https://developer.apple.com/documentation/kernel/absolutetime
type AbsoluteTime = uint64

// See: https://developer.apple.com/documentation/kernel/bddiscinfo
// BDDiscInfo is opaque storage with the size and alignment C gives BDDiscInfo:
// 34 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 34 into.
type BDDiscInfo [34]byte

// See: https://developer.apple.com/documentation/kernel/bdfeatures
type BDFeatures = uint32

// See: https://developer.apple.com/documentation/kernel/bdmediatype
type BDMediaType = uint32

// See: https://developer.apple.com/documentation/kernel/bdtrackinfo
// BDTrackInfo is opaque storage with the size and alignment C gives BDTrackInfo:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type BDTrackInfo [36]byte

// See: https://developer.apple.com/documentation/kernel/btheaderrec
// BTHeaderRec is opaque storage with the size and alignment C gives BTHeaderRec:
// 106 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 106 into.
type BTHeaderRec [53]uint16

// See: https://developer.apple.com/documentation/kernel/btnodedescriptor
// BTNodeDescriptor is opaque storage with the size and alignment C gives BTNodeDescriptor:
// 14 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 14 into.
type BTNodeDescriptor [7]uint16

// See: https://developer.apple.com/documentation/kernel/block0
// Block0 is opaque storage with the size and alignment C gives Block0:
// 512 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 512 into.
type Block0 [512]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothafhhostchannelclassification
// BluetoothAFHHostChannelClassification is opaque storage with the size and alignment C gives BluetoothAFHHostChannelClassification:
// 10 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 10 into.
type BluetoothAFHHostChannelClassification [10]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothafhmode
type BluetoothAFHMode = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothafhresults
// BluetoothAFHResults is opaque storage with the size and alignment C gives BluetoothAFHResults:
// 14 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 14 into.
type BluetoothAFHResults [7]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothampcommandrejectreason
type BluetoothAMPCommandRejectReason = int

// See: https://developer.apple.com/documentation/kernel/bluetoothampcreatephysicallinkresponsestatus
type BluetoothAMPCreatePhysicalLinkResponseStatus = int

// See: https://developer.apple.com/documentation/kernel/bluetoothampdisconnectphysicallinkresponsestatus
type BluetoothAMPDisconnectPhysicalLinkResponseStatus = int

// See: https://developer.apple.com/documentation/kernel/bluetoothampdiscoverresponsecontrollerstatus
type BluetoothAMPDiscoverResponseControllerStatus = int

// See: https://developer.apple.com/documentation/kernel/bluetoothampgetassocresponsestatus
type BluetoothAMPGetAssocResponseStatus = int

// See: https://developer.apple.com/documentation/kernel/bluetoothampgetinforesponsestatus
type BluetoothAMPGetInfoResponseStatus = int

// See: https://developer.apple.com/documentation/kernel/bluetoothampmanagercode
type BluetoothAMPManagerCode = int

// See: https://developer.apple.com/documentation/kernel/bluetoothairmode
type BluetoothAirMode = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothallowroleswitch
type BluetoothAllowRoleSwitch = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothauthenticationrequirements
type BluetoothAuthenticationRequirements = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothclassofdevice
type BluetoothClassOfDevice = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothclockoffset
type BluetoothClockOffset = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothconnectionhandle
type BluetoothConnectionHandle = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothdeviceaddress
// BluetoothDeviceAddress is opaque storage with the size and alignment C gives BluetoothDeviceAddress:
// 6 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 6 into.
type BluetoothDeviceAddress [6]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothdeviceclassmajor
type BluetoothDeviceClassMajor = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothdeviceclassminor
type BluetoothDeviceClassMinor = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothdevicename
type BluetoothDeviceName = byte

// See: https://developer.apple.com/documentation/kernel/bluetoothencryptionenable
type BluetoothEncryptionEnable = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothenhancedsynchronousconnectioninfo
// BluetoothEnhancedSynchronousConnectionInfo is opaque storage with the size and alignment C gives BluetoothEnhancedSynchronousConnectionInfo:
// 80 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 80 into.
type BluetoothEnhancedSynchronousConnectionInfo [10]uint64

// See: https://developer.apple.com/documentation/kernel/bluetootheventfiltercondition
// BluetoothEventFilterCondition is opaque storage with the size and alignment C gives BluetoothEventFilterCondition:
// 7 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 7 into.
type BluetoothEventFilterCondition [7]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothhciacldatabytecount
type BluetoothHCIACLDataByteCount = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhciafhchannelassessmentmode
type BluetoothHCIAFHChannelAssessmentMode = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciacceptsynchronousconnectionrequestparams
// BluetoothHCIAcceptSynchronousConnectionRequestParams is opaque storage with the size and alignment C gives BluetoothHCIAcceptSynchronousConnectionRequestParams:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type BluetoothHCIAcceptSynchronousConnectionRequestParams [4]uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhciauthenticationenable
type BluetoothHCIAuthenticationEnable = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciautomaticflushtimeout
type BluetoothHCIAutomaticFlushTimeout = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhciautomaticflushtimeoutinfo
// BluetoothHCIAutomaticFlushTimeoutInfo is opaque storage with the size and alignment C gives BluetoothHCIAutomaticFlushTimeoutInfo:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type BluetoothHCIAutomaticFlushTimeoutInfo [2]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcibuffersize
// BluetoothHCIBufferSize is opaque storage with the size and alignment C gives BluetoothHCIBufferSize:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type BluetoothHCIBufferSize [4]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcicommandopcode
type BluetoothHCICommandOpCode = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcicommandopcodecommand
type BluetoothHCICommandOpCodeCommand = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcicommandopcodegroup
type BluetoothHCICommandOpCodeGroup = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciconnectionaccepttimeout
type BluetoothHCIConnectionAcceptTimeout = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhciconnectionmode
type BluetoothHCIConnectionMode = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcicontentformat
type BluetoothHCIContentFormat = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcicountrycode
type BluetoothHCICountryCode = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcicurrentinquiryaccesscodes
// BluetoothHCICurrentInquiryAccessCodes is opaque storage with the size and alignment C gives BluetoothHCICurrentInquiryAccessCodes:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type BluetoothHCICurrentInquiryAccessCodes [2]uint64

// See: https://developer.apple.com/documentation/kernel/bluetoothhcicurrentinquiryaccesscodesforwrite
// BluetoothHCICurrentInquiryAccessCodesForWrite is opaque storage with the size and alignment C gives BluetoothHCICurrentInquiryAccessCodesForWrite:
// 193 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 193 into.
type BluetoothHCICurrentInquiryAccessCodesForWrite [193]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothhcidataid
type BluetoothHCIDataID = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhcideletestoredlinkkeyflag
type BluetoothHCIDeleteStoredLinkKeyFlag = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciencryptionkeysize
type BluetoothHCIEncryptionKeySize = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciencryptionkeysizeinfo
// BluetoothHCIEncryptionKeySizeInfo is opaque storage with the size and alignment C gives BluetoothHCIEncryptionKeySizeInfo:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type BluetoothHCIEncryptionKeySizeInfo [2]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhciencryptionmode
type BluetoothHCIEncryptionMode = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcienhancedacceptsynchronousconnectionrequestparams
// BluetoothHCIEnhancedAcceptSynchronousConnectionRequestParams is opaque storage with the size and alignment C gives BluetoothHCIEnhancedAcceptSynchronousConnectionRequestParams:
// 80 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 80 into.
type BluetoothHCIEnhancedAcceptSynchronousConnectionRequestParams [10]uint64

// See: https://developer.apple.com/documentation/kernel/bluetoothhcienhancedsetupsynchronousconnectionparams
// BluetoothHCIEnhancedSetupSynchronousConnectionParams is opaque storage with the size and alignment C gives BluetoothHCIEnhancedSetupSynchronousConnectionParams:
// 80 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 80 into.
type BluetoothHCIEnhancedSetupSynchronousConnectionParams [10]uint64

// See: https://developer.apple.com/documentation/kernel/bluetoothhcierroneousdatareporting
type BluetoothHCIErroneousDataReporting = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventauthenticationcompleteresults
// BluetoothHCIEventAuthenticationCompleteResults is opaque storage with the size and alignment C gives BluetoothHCIEventAuthenticationCompleteResults:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type BluetoothHCIEventAuthenticationCompleteResults [1]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventchangeconnectionlinkkeycompleteresults
// BluetoothHCIEventChangeConnectionLinkKeyCompleteResults is opaque storage with the size and alignment C gives BluetoothHCIEventChangeConnectionLinkKeyCompleteResults:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type BluetoothHCIEventChangeConnectionLinkKeyCompleteResults [1]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventcode
type BluetoothHCIEventCode = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventconnectioncompleteresults
// BluetoothHCIEventConnectionCompleteResults is opaque storage with the size and alignment C gives BluetoothHCIEventConnectionCompleteResults:
// 10 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 10 into.
type BluetoothHCIEventConnectionCompleteResults [5]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventconnectionpackettyperesults
// BluetoothHCIEventConnectionPacketTypeResults is opaque storage with the size and alignment C gives BluetoothHCIEventConnectionPacketTypeResults:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type BluetoothHCIEventConnectionPacketTypeResults [2]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventconnectionrequestresults
// BluetoothHCIEventConnectionRequestResults is opaque storage with the size and alignment C gives BluetoothHCIEventConnectionRequestResults:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type BluetoothHCIEventConnectionRequestResults [4]uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventdatabufferoverflowresults
// BluetoothHCIEventDataBufferOverflowResults is opaque storage with the size and alignment C gives BluetoothHCIEventDataBufferOverflowResults:
// 1 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 1 into.
type BluetoothHCIEventDataBufferOverflowResults [1]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventdisconnectioncompleteresults
// BluetoothHCIEventDisconnectionCompleteResults is opaque storage with the size and alignment C gives BluetoothHCIEventDisconnectionCompleteResults:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type BluetoothHCIEventDisconnectionCompleteResults [2]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventencryptionchangeresults
// BluetoothHCIEventEncryptionChangeResults is opaque storage with the size and alignment C gives BluetoothHCIEventEncryptionChangeResults:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type BluetoothHCIEventEncryptionChangeResults [2]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventencryptionkeyrefreshcompleteresults
// BluetoothHCIEventEncryptionKeyRefreshCompleteResults is opaque storage with the size and alignment C gives BluetoothHCIEventEncryptionKeyRefreshCompleteResults:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type BluetoothHCIEventEncryptionKeyRefreshCompleteResults [1]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventflowspecificationdata
// BluetoothHCIEventFlowSpecificationData is opaque storage with the size and alignment C gives BluetoothHCIEventFlowSpecificationData:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type BluetoothHCIEventFlowSpecificationData [6]uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventflushoccurredresults
// BluetoothHCIEventFlushOccurredResults is opaque storage with the size and alignment C gives BluetoothHCIEventFlushOccurredResults:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type BluetoothHCIEventFlushOccurredResults [1]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventhardwareerrorresults
// BluetoothHCIEventHardwareErrorResults is opaque storage with the size and alignment C gives BluetoothHCIEventHardwareErrorResults:
// 1 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 1 into.
type BluetoothHCIEventHardwareErrorResults [1]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventid
type BluetoothHCIEventID = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventleconnectioncompleteresults
// BluetoothHCIEventLEConnectionCompleteResults is opaque storage with the size and alignment C gives BluetoothHCIEventLEConnectionCompleteResults:
// 17 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 17 into.
type BluetoothHCIEventLEConnectionCompleteResults [17]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventleconnectionupdatecompleteresults
// BluetoothHCIEventLEConnectionUpdateCompleteResults is opaque storage with the size and alignment C gives BluetoothHCIEventLEConnectionUpdateCompleteResults:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type BluetoothHCIEventLEConnectionUpdateCompleteResults [8]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventleenhancedconnectioncompleteresults
// BluetoothHCIEventLEEnhancedConnectionCompleteResults is opaque storage with the size and alignment C gives BluetoothHCIEventLEEnhancedConnectionCompleteResults:
// 29 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 29 into.
type BluetoothHCIEventLEEnhancedConnectionCompleteResults [29]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventlelongtermkeyrequestresults
// BluetoothHCIEventLELongTermKeyRequestResults is opaque storage with the size and alignment C gives BluetoothHCIEventLELongTermKeyRequestResults:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type BluetoothHCIEventLELongTermKeyRequestResults [6]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventlemetaresults
// BluetoothHCIEventLEMetaResults is opaque storage with the size and alignment C gives BluetoothHCIEventLEMetaResults:
// 256 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 256 into.
type BluetoothHCIEventLEMetaResults [256]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventlereadremoteusedfeaturescompleteresults
// BluetoothHCIEventLEReadRemoteUsedFeaturesCompleteResults is opaque storage with the size and alignment C gives BluetoothHCIEventLEReadRemoteUsedFeaturesCompleteResults:
// 10 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 10 into.
type BluetoothHCIEventLEReadRemoteUsedFeaturesCompleteResults [10]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventlinkkeynotificationresults
// BluetoothHCIEventLinkKeyNotificationResults is opaque storage with the size and alignment C gives BluetoothHCIEventLinkKeyNotificationResults:
// 23 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 23 into.
type BluetoothHCIEventLinkKeyNotificationResults [23]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventmask
type BluetoothHCIEventMask = uint64

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventmasterlinkkeycompleteresults
// BluetoothHCIEventMasterLinkKeyCompleteResults is opaque storage with the size and alignment C gives BluetoothHCIEventMasterLinkKeyCompleteResults:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type BluetoothHCIEventMasterLinkKeyCompleteResults [2]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventmaxslotschangeresults
// BluetoothHCIEventMaxSlotsChangeResults is opaque storage with the size and alignment C gives BluetoothHCIEventMaxSlotsChangeResults:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type BluetoothHCIEventMaxSlotsChangeResults [2]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventmodechangeresults
// BluetoothHCIEventModeChangeResults is opaque storage with the size and alignment C gives BluetoothHCIEventModeChangeResults:
// 6 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 6 into.
type BluetoothHCIEventModeChangeResults [3]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventpagescanmodechangeresults
// BluetoothHCIEventPageScanModeChangeResults is opaque storage with the size and alignment C gives BluetoothHCIEventPageScanModeChangeResults:
// 7 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 7 into.
type BluetoothHCIEventPageScanModeChangeResults [7]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventpagescanrepetitionmodechangeresults
// BluetoothHCIEventPageScanRepetitionModeChangeResults is opaque storage with the size and alignment C gives BluetoothHCIEventPageScanRepetitionModeChangeResults:
// 7 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 7 into.
type BluetoothHCIEventPageScanRepetitionModeChangeResults [7]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventqossetupcompleteresults
// BluetoothHCIEventQoSSetupCompleteResults is opaque storage with the size and alignment C gives BluetoothHCIEventQoSSetupCompleteResults:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type BluetoothHCIEventQoSSetupCompleteResults [6]uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventqosviolationresults
// BluetoothHCIEventQoSViolationResults is opaque storage with the size and alignment C gives BluetoothHCIEventQoSViolationResults:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type BluetoothHCIEventQoSViolationResults [1]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventreadclockoffsetresults
// BluetoothHCIEventReadClockOffsetResults is opaque storage with the size and alignment C gives BluetoothHCIEventReadClockOffsetResults:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type BluetoothHCIEventReadClockOffsetResults [2]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventreadextendedfeaturesresults
// BluetoothHCIEventReadExtendedFeaturesResults is opaque storage with the size and alignment C gives BluetoothHCIEventReadExtendedFeaturesResults:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type BluetoothHCIEventReadExtendedFeaturesResults [6]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventreadremoteextendedfeaturesresults
// BluetoothHCIEventReadRemoteExtendedFeaturesResults is opaque storage with the size and alignment C gives BluetoothHCIEventReadRemoteExtendedFeaturesResults:
// 14 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 14 into.
type BluetoothHCIEventReadRemoteExtendedFeaturesResults [7]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventreadremotesupportedfeaturesresults
// BluetoothHCIEventReadRemoteSupportedFeaturesResults is opaque storage with the size and alignment C gives BluetoothHCIEventReadRemoteSupportedFeaturesResults:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type BluetoothHCIEventReadRemoteSupportedFeaturesResults [6]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventreadremoteversioninforesults
// BluetoothHCIEventReadRemoteVersionInfoResults is opaque storage with the size and alignment C gives BluetoothHCIEventReadRemoteVersionInfoResults:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type BluetoothHCIEventReadRemoteVersionInfoResults [4]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventreadsupportedfeaturesresults
// BluetoothHCIEventReadSupportedFeaturesResults is opaque storage with the size and alignment C gives BluetoothHCIEventReadSupportedFeaturesResults:
// 10 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 10 into.
type BluetoothHCIEventReadSupportedFeaturesResults [5]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventremotenamerequestresults
// BluetoothHCIEventRemoteNameRequestResults is opaque storage with the size and alignment C gives BluetoothHCIEventRemoteNameRequestResults:
// 254 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 254 into.
type BluetoothHCIEventRemoteNameRequestResults [254]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventreturnlinkkeysresults
// BluetoothHCIEventReturnLinkKeysResults is opaque storage with the size and alignment C gives BluetoothHCIEventReturnLinkKeysResults:
// 23 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 23 into.
type BluetoothHCIEventReturnLinkKeysResults [23]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventrolechangeresults
// BluetoothHCIEventRoleChangeResults is opaque storage with the size and alignment C gives BluetoothHCIEventRoleChangeResults:
// 10 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 10 into.
type BluetoothHCIEventRoleChangeResults [5]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventsimplepairingcompleteresults
// BluetoothHCIEventSimplePairingCompleteResults is opaque storage with the size and alignment C gives BluetoothHCIEventSimplePairingCompleteResults:
// 6 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 6 into.
type BluetoothHCIEventSimplePairingCompleteResults [6]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventsniffsubratingresults
// BluetoothHCIEventSniffSubratingResults is opaque storage with the size and alignment C gives BluetoothHCIEventSniffSubratingResults:
// 10 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 10 into.
type BluetoothHCIEventSniffSubratingResults [5]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventstatus
type BluetoothHCIEventStatus = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventsynchronousconnectionchangedresults
// BluetoothHCIEventSynchronousConnectionChangedResults is opaque storage with the size and alignment C gives BluetoothHCIEventSynchronousConnectionChangedResults:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type BluetoothHCIEventSynchronousConnectionChangedResults [4]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventsynchronousconnectioncompleteresults
// BluetoothHCIEventSynchronousConnectionCompleteResults is opaque storage with the size and alignment C gives BluetoothHCIEventSynchronousConnectionCompleteResults:
// 18 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 18 into.
type BluetoothHCIEventSynchronousConnectionCompleteResults [9]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventvendorspecificresults
// BluetoothHCIEventVendorSpecificResults is opaque storage with the size and alignment C gives BluetoothHCIEventVendorSpecificResults:
// 256 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 256 into.
type BluetoothHCIEventVendorSpecificResults [256]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothhciextendedfeaturesinfo
// BluetoothHCIExtendedFeaturesInfo is opaque storage with the size and alignment C gives BluetoothHCIExtendedFeaturesInfo:
// 10 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 10 into.
type BluetoothHCIExtendedFeaturesInfo [10]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothhciextendedinquiryresponse
// BluetoothHCIExtendedInquiryResponse is opaque storage with the size and alignment C gives BluetoothHCIExtendedInquiryResponse:
// 240 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 240 into.
type BluetoothHCIExtendedInquiryResponse [240]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothhciextendedinquiryresponsedatatype
type BluetoothHCIExtendedInquiryResponseDataType = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciextendedinquiryresult
// BluetoothHCIExtendedInquiryResult is opaque storage with the size and alignment C gives BluetoothHCIExtendedInquiryResult:
// 260 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 260 into.
type BluetoothHCIExtendedInquiryResult [65]uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhcifecrequired
type BluetoothHCIFECRequired = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcifailedcontactcount
type BluetoothHCIFailedContactCount = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcifailedcontactinfo
// BluetoothHCIFailedContactInfo is opaque storage with the size and alignment C gives BluetoothHCIFailedContactInfo:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type BluetoothHCIFailedContactInfo [2]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhciflowcontrolstate
type BluetoothHCIFlowControlState = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciholdmodeactivity
type BluetoothHCIHoldModeActivity = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciinputbandwidth
type BluetoothHCIInputBandwidth = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhciinputcodeddatasize
type BluetoothHCIInputCodedDataSize = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhciinputcodingformat
type BluetoothHCIInputCodingFormat = uint64

// See: https://developer.apple.com/documentation/kernel/bluetoothhciinputdatapath
type BluetoothHCIInputDataPath = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciinputpcmdataformat
type BluetoothHCIInputPCMDataFormat = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciinputpcmsamplepayloadmsbposition
type BluetoothHCIInputPCMSamplePayloadMSBPosition = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciinputtransportunitsize
type BluetoothHCIInputTransportUnitSize = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciinquiryaccesscode
// BluetoothHCIInquiryAccessCode is opaque storage with the size and alignment C gives BluetoothHCIInquiryAccessCode:
// 3 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 3 into.
type BluetoothHCIInquiryAccessCode [3]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothhciinquiryaccesscodecount
type BluetoothHCIInquiryAccessCodeCount = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciinquirylength
type BluetoothHCIInquiryLength = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciinquirymode
type BluetoothHCIInquiryMode = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciinquiryresult
// BluetoothHCIInquiryResult is opaque storage with the size and alignment C gives BluetoothHCIInquiryResult:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type BluetoothHCIInquiryResult [5]uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhciinquiryresults
// BluetoothHCIInquiryResults is opaque storage with the size and alignment C gives BluetoothHCIInquiryResults:
// 1004 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 1004 into.
type BluetoothHCIInquiryResults [251]uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhciinquiryscantype
type BluetoothHCIInquiryScanType = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciinquirywithrssiresult
// BluetoothHCIInquiryWithRSSIResult is opaque storage with the size and alignment C gives BluetoothHCIInquiryWithRSSIResult:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type BluetoothHCIInquiryWithRSSIResult [4]uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhciinquirywithrssiresults
// BluetoothHCIInquiryWithRSSIResults is opaque storage with the size and alignment C gives BluetoothHCIInquiryWithRSSIResults:
// 804 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 804 into.
type BluetoothHCIInquiryWithRSSIResults [201]uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhcilebuffersize
// BluetoothHCILEBufferSize is opaque storage with the size and alignment C gives BluetoothHCILEBufferSize:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type BluetoothHCILEBufferSize [2]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcilesupportedfeatures
type BluetoothHCILESupportedFeatures = BluetoothHCISupportedFeatures

// See: https://developer.apple.com/documentation/kernel/bluetoothhcileusedfeatures
type BluetoothHCILEUsedFeatures = BluetoothHCISupportedFeatures

// See: https://developer.apple.com/documentation/kernel/bluetoothhcilinkpolicysettings
type BluetoothHCILinkPolicySettings = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcilinkpolicysettingsinfo
// BluetoothHCILinkPolicySettingsInfo is opaque storage with the size and alignment C gives BluetoothHCILinkPolicySettingsInfo:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type BluetoothHCILinkPolicySettingsInfo [2]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcilinkquality
type BluetoothHCILinkQuality = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcilinkqualityinfo
// BluetoothHCILinkQualityInfo is opaque storage with the size and alignment C gives BluetoothHCILinkQualityInfo:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type BluetoothHCILinkQualityInfo [2]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcilinksupervisiontimeout
// BluetoothHCILinkSupervisionTimeout is opaque storage with the size and alignment C gives BluetoothHCILinkSupervisionTimeout:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type BluetoothHCILinkSupervisionTimeout [2]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhciloopbackmode
type BluetoothHCILoopbackMode = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcimaxlatency
type BluetoothHCIMaxLatency = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcimodeinterval
type BluetoothHCIModeInterval = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcinumbroadcastretransmissions
type BluetoothHCINumBroadcastRetransmissions = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcinumlinkkeysdeleted
type BluetoothHCINumLinkKeysDeleted = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcinumlinkkeystowrite
type BluetoothHCINumLinkKeysToWrite = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcioperationid
type BluetoothHCIOperationID = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhcioutputbandwidth
type BluetoothHCIOutputBandwidth = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhcioutputcodeddatasize
type BluetoothHCIOutputCodedDataSize = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcioutputcodingformat
type BluetoothHCIOutputCodingFormat = uint64

// See: https://developer.apple.com/documentation/kernel/bluetoothhcioutputdatapath
type BluetoothHCIOutputDataPath = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcioutputpcmdataformat
type BluetoothHCIOutputPCMDataFormat = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcioutputpcmsamplepayloadmsbposition
type BluetoothHCIOutputPCMSamplePayloadMSBPosition = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcioutputtransportunitsize
type BluetoothHCIOutputTransportUnitSize = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcipagenumber
type BluetoothHCIPageNumber = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcipagescanenablestate
type BluetoothHCIPageScanEnableState = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcipagescanmode
type BluetoothHCIPageScanMode = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcipagescanperiodmode
type BluetoothHCIPageScanPeriodMode = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcipagescantype
type BluetoothHCIPageScanType = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcipagetimeout
type BluetoothHCIPageTimeout = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhciparambytecount
type BluetoothHCIParamByteCount = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciparkmodebeaconinterval
type BluetoothHCIParkModeBeaconInterval = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcipowerstate
type BluetoothHCIPowerState = int

// See: https://developer.apple.com/documentation/kernel/bluetoothhciqosflags
type BluetoothHCIQoSFlags = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciqualityofservicesetupparams
// BluetoothHCIQualityOfServiceSetupParams is opaque storage with the size and alignment C gives BluetoothHCIQualityOfServiceSetupParams:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type BluetoothHCIQualityOfServiceSetupParams [5]uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhcirssiinfo
// BluetoothHCIRSSIInfo is opaque storage with the size and alignment C gives BluetoothHCIRSSIInfo:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type BluetoothHCIRSSIInfo [2]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcirssivalue
type BluetoothHCIRSSIValue = int8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcireadextendedinquiryresponseresults
// BluetoothHCIReadExtendedInquiryResponseResults is opaque storage with the size and alignment C gives BluetoothHCIReadExtendedInquiryResponseResults:
// 241 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 241 into.
type BluetoothHCIReadExtendedInquiryResponseResults [241]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothhcireadlmphandleresults
// BluetoothHCIReadLMPHandleResults is opaque storage with the size and alignment C gives BluetoothHCIReadLMPHandleResults:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type BluetoothHCIReadLMPHandleResults [2]uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhcireadlocaloobdataresults
// BluetoothHCIReadLocalOOBDataResults is opaque storage with the size and alignment C gives BluetoothHCIReadLocalOOBDataResults:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type BluetoothHCIReadLocalOOBDataResults [32]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothhcireadstoredlinkkeysflag
type BluetoothHCIReadStoredLinkKeysFlag = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcireceivebandwidth
type BluetoothHCIReceiveBandwidth = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhcireceivecodecframesize
type BluetoothHCIReceiveCodecFrameSize = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcireceivecodingformat
type BluetoothHCIReceiveCodingFormat = uint64

// See: https://developer.apple.com/documentation/kernel/bluetoothhcirequestcallbackinfo
// BluetoothHCIRequestCallbackInfo is opaque storage with the size and alignment C gives BluetoothHCIRequestCallbackInfo:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type BluetoothHCIRequestCallbackInfo [5]uint64

// See: https://developer.apple.com/documentation/kernel/bluetoothhcirequestid
type BluetoothHCIRequestID = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhciresponsecount
type BluetoothHCIResponseCount = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciretransmissioneffort
type BluetoothHCIRetransmissionEffort = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcirole
type BluetoothHCIRole = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciroleinfo
// BluetoothHCIRoleInfo is opaque storage with the size and alignment C gives BluetoothHCIRoleInfo:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type BluetoothHCIRoleInfo [2]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhciscodatabytecount
type BluetoothHCISCODataByteCount = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciscanactivity
// BluetoothHCIScanActivity is opaque storage with the size and alignment C gives BluetoothHCIScanActivity:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type BluetoothHCIScanActivity [2]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcisetupsynchronousconnectionparams
// BluetoothHCISetupSynchronousConnectionParams is opaque storage with the size and alignment C gives BluetoothHCISetupSynchronousConnectionParams:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type BluetoothHCISetupSynchronousConnectionParams [4]uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhcisignalid
type BluetoothHCISignalID = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhcisimplepairingmode
type BluetoothHCISimplePairingMode = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcisimplepairingoobdata
// BluetoothHCISimplePairingOOBData is opaque storage with the size and alignment C gives BluetoothHCISimplePairingOOBData:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type BluetoothHCISimplePairingOOBData [16]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothhcisniffattemptcount
type BluetoothHCISniffAttemptCount = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcisnifftimeout
type BluetoothHCISniffTimeout = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcistatus
type BluetoothHCIStatus = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcistoredlinkkeysinfo
// BluetoothHCIStoredLinkKeysInfo is opaque storage with the size and alignment C gives BluetoothHCIStoredLinkKeysInfo:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type BluetoothHCIStoredLinkKeysInfo [2]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcisupportedcommands
// BluetoothHCISupportedCommands is opaque storage with the size and alignment C gives BluetoothHCISupportedCommands:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type BluetoothHCISupportedCommands [64]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothhcisupportedfeatures
// BluetoothHCISupportedFeatures is opaque storage with the size and alignment C gives BluetoothHCISupportedFeatures:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type BluetoothHCISupportedFeatures [8]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothhcisupportediac
type BluetoothHCISupportedIAC = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcitransmitbandwidth
type BluetoothHCITransmitBandwidth = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhcitransmitcodecframesize
type BluetoothHCITransmitCodecFrameSize = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcitransmitcodingformat
type BluetoothHCITransmitCodingFormat = uint64

// See: https://developer.apple.com/documentation/kernel/bluetoothhcitransmitpowerlevel
type BluetoothHCITransmitPowerLevel = int8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcitransmitpowerlevelinfo
// BluetoothHCITransmitPowerLevelInfo is opaque storage with the size and alignment C gives BluetoothHCITransmitPowerLevelInfo:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type BluetoothHCITransmitPowerLevelInfo [2]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcitransmitpowerleveltype
type BluetoothHCITransmitPowerLevelType = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcitransportcommandid
type BluetoothHCITransportCommandID = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhcitransportid
type BluetoothHCITransportID = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhcivendorcommandselector
type BluetoothHCIVendorCommandSelector = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhciversioninfo
// BluetoothHCIVersionInfo is opaque storage with the size and alignment C gives BluetoothHCIVersionInfo:
// 10 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 10 into.
type BluetoothHCIVersionInfo [5]uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhciversions
type BluetoothHCIVersions = int

// See: https://developer.apple.com/documentation/kernel/bluetoothhcivoicesetting
type BluetoothHCIVoiceSetting = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothiocapability
type BluetoothIOCapability = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothiocapabilityresponse
// BluetoothIOCapabilityResponse is opaque storage with the size and alignment C gives BluetoothIOCapabilityResponse:
// 9 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 9 into.
type BluetoothIOCapabilityResponse [9]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothirk
// BluetoothIRK is opaque storage with the size and alignment C gives BluetoothIRK:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type BluetoothIRK [16]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothkey
type BluetoothKey = string

// See: https://developer.apple.com/documentation/kernel/bluetoothkeyflag
type BluetoothKeyFlag = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothkeytype
type BluetoothKeyType = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothkeypressnotification
// BluetoothKeypressNotification is opaque storage with the size and alignment C gives BluetoothKeypressNotification:
// 7 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 7 into.
type BluetoothKeypressNotification [7]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothkeypressnotificationtype
type BluetoothKeypressNotificationType = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothl2capbytecount
type BluetoothL2CAPByteCount = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothl2capchannelid
type BluetoothL2CAPChannelID = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothl2capcommandbytecount
type BluetoothL2CAPCommandByteCount = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothl2capcommandcode
type BluetoothL2CAPCommandCode = int

// See: https://developer.apple.com/documentation/kernel/bluetoothl2capcommandid
type BluetoothL2CAPCommandID = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothl2capcommandrejectreason
type BluetoothL2CAPCommandRejectReason = int

// See: https://developer.apple.com/documentation/kernel/bluetoothl2capconfigurationoption
type BluetoothL2CAPConfigurationOption = int

// See: https://developer.apple.com/documentation/kernel/bluetoothl2capconfigurationresult
type BluetoothL2CAPConfigurationResult = int

// See: https://developer.apple.com/documentation/kernel/bluetoothl2capconfigurationretransmissionandflowcontrolflags
type BluetoothL2CAPConfigurationRetransmissionAndFlowControlFlags = uint

// See: https://developer.apple.com/documentation/kernel/bluetoothl2capconnectionresult
type BluetoothL2CAPConnectionResult = int

// See: https://developer.apple.com/documentation/kernel/bluetoothl2capconnectionstatus
type BluetoothL2CAPConnectionStatus = int

// See: https://developer.apple.com/documentation/kernel/bluetoothl2capflushtimeout
type BluetoothL2CAPFlushTimeout = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothl2capgroupid
type BluetoothL2CAPGroupID = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothl2capinformationextendedfeaturesmask
type BluetoothL2CAPInformationExtendedFeaturesMask = uint

// See: https://developer.apple.com/documentation/kernel/bluetoothl2capinformationresult
type BluetoothL2CAPInformationResult = int

// See: https://developer.apple.com/documentation/kernel/bluetoothl2capinformationtype
type BluetoothL2CAPInformationType = int

// See: https://developer.apple.com/documentation/kernel/bluetoothl2caplinktimeout
type BluetoothL2CAPLinkTimeout = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothl2capmtu
type BluetoothL2CAPMTU = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothl2cappsm
type BluetoothL2CAPPSM = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothl2capqostype
type BluetoothL2CAPQoSType = int

// See: https://developer.apple.com/documentation/kernel/bluetoothl2capqualityofserviceoptions
type BluetoothL2CAPQualityOfServiceOptions = uint

// See: https://developer.apple.com/documentation/kernel/bluetoothl2capretransmissionandflowcontroloptions
type BluetoothL2CAPRetransmissionAndFlowControlOptions = uint

// See: https://developer.apple.com/documentation/kernel/bluetoothl2capsegmentationandreassembly
type BluetoothL2CAPSegmentationAndReassembly = int

// See: https://developer.apple.com/documentation/kernel/bluetoothl2capsupervisoryfuctiontype
type BluetoothL2CAPSupervisoryFuctionType = int

// See: https://developer.apple.com/documentation/kernel/bluetoothlap
type BluetoothLAP = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothleaddresstype
type BluetoothLEAddressType = int

// See: https://developer.apple.com/documentation/kernel/bluetoothleadvertisingtype
type BluetoothLEAdvertisingType = int

// See: https://developer.apple.com/documentation/kernel/bluetoothleconnectioninterval
type BluetoothLEConnectionInterval = int

// See: https://developer.apple.com/documentation/kernel/bluetoothlescan
type BluetoothLEScan = int

// See: https://developer.apple.com/documentation/kernel/bluetoothlescanduplicatefilter
type BluetoothLEScanDuplicateFilter = int

// See: https://developer.apple.com/documentation/kernel/bluetoothlescanfilter
type BluetoothLEScanFilter = int

// See: https://developer.apple.com/documentation/kernel/bluetoothlescantype
type BluetoothLEScanType = int

// See: https://developer.apple.com/documentation/kernel/bluetoothlesecuritymanagercommandcode
type BluetoothLESecurityManagerCommandCode = int

// See: https://developer.apple.com/documentation/kernel/bluetoothlesecuritymanageriocapability
type BluetoothLESecurityManagerIOCapability = int

// See: https://developer.apple.com/documentation/kernel/bluetoothlesecuritymanagerkeypressnotificationtype
type BluetoothLESecurityManagerKeypressNotificationType = int

// See: https://developer.apple.com/documentation/kernel/bluetoothlesecuritymanageroobdata
type BluetoothLESecurityManagerOOBData = int

// See: https://developer.apple.com/documentation/kernel/bluetoothlesecuritymanagerpairingfailedreasoncode
type BluetoothLESecurityManagerPairingFailedReasonCode = int

// See: https://developer.apple.com/documentation/kernel/bluetoothlesecuritymanageruserinputcapability
type BluetoothLESecurityManagerUserInputCapability = int

// See: https://developer.apple.com/documentation/kernel/bluetoothlesecuritymanageruseroutputcapability
type BluetoothLESecurityManagerUserOutputCapability = int

// See: https://developer.apple.com/documentation/kernel/bluetoothlmphandle
type BluetoothLMPHandle = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothlmpsubversion
type BluetoothLMPSubversion = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothlmpversion
type BluetoothLMPVersion = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothlmpversions
type BluetoothLMPVersions = int

// See: https://developer.apple.com/documentation/kernel/bluetoothlinktype
type BluetoothLinkType = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothmanufacturername
type BluetoothManufacturerName = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothmaxslots
type BluetoothMaxSlots = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothnumericvalue
type BluetoothNumericValue = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothoobdatapresence
type BluetoothOOBDataPresence = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothpincode
// BluetoothPINCode is opaque storage with the size and alignment C gives BluetoothPINCode:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type BluetoothPINCode [16]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothpintype
type BluetoothPINType = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothpackettype
type BluetoothPacketType = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothpagescanmode
type BluetoothPageScanMode = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothpagescanperiodmode
type BluetoothPageScanPeriodMode = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothpagescanrepetitionmode
type BluetoothPageScanRepetitionMode = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothpasskey
type BluetoothPasskey = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothrfcommchannelid
type BluetoothRFCOMMChannelID = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothrfcommlinestatus
type BluetoothRFCOMMLineStatus = int

// See: https://developer.apple.com/documentation/kernel/bluetoothrfcommmtu
type BluetoothRFCOMMMTU = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothrfcommparitytype
type BluetoothRFCOMMParityType = int

// See: https://developer.apple.com/documentation/kernel/bluetoothreadclockinfo
// BluetoothReadClockInfo is opaque storage with the size and alignment C gives BluetoothReadClockInfo:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type BluetoothReadClockInfo [3]uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothreasoncode
type BluetoothReasonCode = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothremotehostsupportedfeaturesnotification
// BluetoothRemoteHostSupportedFeaturesNotification is opaque storage with the size and alignment C gives BluetoothRemoteHostSupportedFeaturesNotification:
// 14 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 14 into.
type BluetoothRemoteHostSupportedFeaturesNotification [14]byte

// See: https://developer.apple.com/documentation/kernel/bluetoothrole
type BluetoothRole = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothsdpdataelementsizedescriptor
type BluetoothSDPDataElementSizeDescriptor = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothsdpdataelementtypedescriptor
type BluetoothSDPDataElementTypeDescriptor = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothsdperrorcode
type BluetoothSDPErrorCode = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothsdppduid
type BluetoothSDPPDUID = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothsdpserviceattributeid
type BluetoothSDPServiceAttributeID = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothsdpservicerecordhandle
type BluetoothSDPServiceRecordHandle = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothsdptransactionid
type BluetoothSDPTransactionID = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothsdpuuid16
type BluetoothSDPUUID16 = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothsdpuuid32
type BluetoothSDPUUID32 = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothserviceclassmajor
type BluetoothServiceClassMajor = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothseteventmask
type BluetoothSetEventMask = uint

// See: https://developer.apple.com/documentation/kernel/bluetoothsimplepairingdebugmode
type BluetoothSimplePairingDebugMode = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothsynchronousconnectioninfo
// BluetoothSynchronousConnectionInfo is opaque storage with the size and alignment C gives BluetoothSynchronousConnectionInfo:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type BluetoothSynchronousConnectionInfo [4]uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothtransportinfo
// BluetoothTransportInfo is opaque storage with the size and alignment C gives BluetoothTransportInfo:
// 120 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 120 into.
type BluetoothTransportInfo [15]uint64

// See: https://developer.apple.com/documentation/kernel/bluetoothtransportinfoptr
type BluetoothTransportInfoPtr = *BluetoothTransportInfo

// See: https://developer.apple.com/documentation/kernel/bluetoothuserconfirmationrequest
// BluetoothUserConfirmationRequest is opaque storage with the size and alignment C gives BluetoothUserConfirmationRequest:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type BluetoothUserConfirmationRequest [3]uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothuserpasskeynotification
// BluetoothUserPasskeyNotification is opaque storage with the size and alignment C gives BluetoothUserPasskeyNotification:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type BluetoothUserPasskeyNotification [3]uint32

// See: https://developer.apple.com/documentation/kernel/boolean
type Boolean = bool

// See: https://developer.apple.com/documentation/kernel/boot_video
// Boot_Video is opaque storage with the size and alignment C gives Boot_Video:
// 56 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 56 into.
type Boot_Video [7]uint64

// See: https://developer.apple.com/documentation/kernel/boot_videov1
// Boot_VideoV1 is opaque storage with the size and alignment C gives Boot_VideoV1:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Boot_VideoV1 [6]uint32

// See: https://developer.apple.com/documentation/kernel/bounds
type Bounds = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/byte
type Byte = byte

// ByteCount is abst_ByteCount.
//
// See: https://developer.apple.com/documentation/kernel/bytecount
type ByteCount = uint

// See: https://developer.apple.com/documentation/kernel/byteptr
type BytePtr = *uint8

// See: https://developer.apple.com/documentation/kernel/bytef
type Bytef = byte

// See: https://developer.apple.com/documentation/kernel/cdatip
// CDATIP is opaque storage with the size and alignment C gives CDATIP:
// 28 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 28 into.
type CDATIP [28]byte

// See: https://developer.apple.com/documentation/kernel/cdaudiostatus
// CDAudioStatus is opaque storage with the size and alignment C gives CDAudioStatus:
// 9 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 9 into.
type CDAudioStatus [9]byte

// See: https://developer.apple.com/documentation/kernel/cddiscinfo
// CDDiscInfo is opaque storage with the size and alignment C gives CDDiscInfo:
// 34 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 34 into.
type CDDiscInfo [34]byte

// See: https://developer.apple.com/documentation/kernel/cdfeatures
type CDFeatures = uint32

// See: https://developer.apple.com/documentation/kernel/cdisrc
type CDISRC = int8

// See: https://developer.apple.com/documentation/kernel/cdmcn
type CDMCN = int8

// See: https://developer.apple.com/documentation/kernel/cdmsf
// CDMSF is opaque storage with the size and alignment C gives CDMSF:
// 3 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 3 into.
type CDMSF [3]byte

// See: https://developer.apple.com/documentation/kernel/cdmediatype
type CDMediaType = uint32

// See: https://developer.apple.com/documentation/kernel/cdpma
// CDPMA is opaque storage with the size and alignment C gives CDPMA:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type CDPMA [4]byte

// See: https://developer.apple.com/documentation/kernel/cdpmadescriptor
// CDPMADescriptor is opaque storage with the size and alignment C gives CDPMADescriptor:
// 11 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 11 into.
type CDPMADescriptor [11]byte

// See: https://developer.apple.com/documentation/kernel/cdsectorarea
type CDSectorArea = int

// See: https://developer.apple.com/documentation/kernel/cdsectorsize
type CDSectorSize = int

// See: https://developer.apple.com/documentation/kernel/cdsectortype
type CDSectorType = int

// See: https://developer.apple.com/documentation/kernel/cdtext
// CDTEXT is opaque storage with the size and alignment C gives CDTEXT:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type CDTEXT [4]byte

// See: https://developer.apple.com/documentation/kernel/cdtextdescriptor
// CDTEXTDescriptor is opaque storage with the size and alignment C gives CDTEXTDescriptor:
// 18 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 18 into.
type CDTEXTDescriptor [18]byte

// See: https://developer.apple.com/documentation/kernel/cdtoc
// CDTOC is opaque storage with the size and alignment C gives CDTOC:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type CDTOC [4]byte

// See: https://developer.apple.com/documentation/kernel/cdtocdescriptor
// CDTOCDescriptor is opaque storage with the size and alignment C gives CDTOCDescriptor:
// 11 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 11 into.
type CDTOCDescriptor [11]byte

// See: https://developer.apple.com/documentation/kernel/cdtocformat
type CDTOCFormat = byte

// See: https://developer.apple.com/documentation/kernel/cdtrackinfo
// CDTrackInfo is opaque storage with the size and alignment C gives CDTrackInfo:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type CDTrackInfo [36]byte

// See: https://developer.apple.com/documentation/kernel/cdtrackinfoaddresstype
type CDTrackInfoAddressType = byte

// See: https://developer.apple.com/documentation/kernel/complex
// COMPLEX is opaque storage with the size and alignment C gives COMPLEX:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type COMPLEX [2]uint32

// See: https://developer.apple.com/documentation/kernel/complex_split
// COMPLEX_SPLIT is opaque storage with the size and alignment C gives COMPLEX_SPLIT:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type COMPLEX_SPLIT [2]uint64

// See: https://developer.apple.com/documentation/kernel/csrnodeuniqueid
type CSRNodeUniqueID = uint64

// See: https://developer.apple.com/documentation/kernel/cs_blobindex
// CS_BlobIndex is opaque storage with the size and alignment C gives CS_BlobIndex:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type CS_BlobIndex [2]uint32

// See: https://developer.apple.com/documentation/kernel/cs_codedirectory
// CS_CodeDirectory is opaque storage with the size and alignment C gives CS_CodeDirectory:
// 112 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 112 into.
type CS_CodeDirectory [14]uint64

// See: https://developer.apple.com/documentation/kernel/cs_genericblob
// CS_GenericBlob is opaque storage with the size and alignment C gives CS_GenericBlob:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type CS_GenericBlob [2]uint32

// See: https://developer.apple.com/documentation/kernel/cs_superblob
// CS_SuperBlob is opaque storage with the size and alignment C gives CS_SuperBlob:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type CS_SuperBlob [3]uint32

// See: https://developer.apple.com/documentation/kernel/colorspec
// ColorSpec is opaque storage with the size and alignment C gives ColorSpec:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type ColorSpec [4]uint16

// See: https://developer.apple.com/documentation/kernel/colorspecptr
type ColorSpecPtr = *ColorSpec

// See: https://developer.apple.com/documentation/kernel/consthfsunistr255param
type ConstHFSUniStr255Param = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dasdmodeparameterblockdescriptor
// DASDModeParameterBlockDescriptor is opaque storage with the size and alignment C gives DASDModeParameterBlockDescriptor:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type DASDModeParameterBlockDescriptor [8]byte

// See: https://developer.apple.com/documentation/kernel/dclcallcommandproc
type DCLCallCommandProc = *objc.ID

// See: https://developer.apple.com/documentation/kernel/dclcallcommandprocptr
type DCLCallCommandProcPtr = *DCLCallCommandProc

// See: https://developer.apple.com/documentation/kernel/dclcallproc
// DCLCallProc is opaque storage with the size and alignment C gives DCLCallProc:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type DCLCallProc [4]uint64

// See: https://developer.apple.com/documentation/kernel/dclcallprocdatatype
type DCLCallProcDataType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dclcallprocptr
type DCLCallProcPtr = *DCLCallProc

// See: https://developer.apple.com/documentation/kernel/dclcommand
// DCLCommand is opaque storage with the size and alignment C gives DCLCommand:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type DCLCommand [3]uint64

// See: https://developer.apple.com/documentation/kernel/dclcommandptr
type DCLCommandPtr = *DCLCommand

// See: https://developer.apple.com/documentation/kernel/dclcompilerdatatype
type DCLCompilerDataType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dcljump
// DCLJump is opaque storage with the size and alignment C gives DCLJump:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type DCLJump [3]uint64

// See: https://developer.apple.com/documentation/kernel/dcljumpptr
type DCLJumpPtr = *DCLJump

// See: https://developer.apple.com/documentation/kernel/dcllabel
// DCLLabel is opaque storage with the size and alignment C gives DCLLabel:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type DCLLabel [2]uint64

// See: https://developer.apple.com/documentation/kernel/dcllabelptr
type DCLLabelPtr = *DCLLabel

// See: https://developer.apple.com/documentation/kernel/dclnudclleader
// DCLNuDCLLeader is opaque storage with the size and alignment C gives DCLNuDCLLeader:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type DCLNuDCLLeader [3]uint64

// See: https://developer.apple.com/documentation/kernel/dclptrtimestamp
// DCLPtrTimeStamp is opaque storage with the size and alignment C gives DCLPtrTimeStamp:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type DCLPtrTimeStamp [3]uint64

// See: https://developer.apple.com/documentation/kernel/dclptrtimestampptr
type DCLPtrTimeStampPtr = *DCLPtrTimeStamp

// See: https://developer.apple.com/documentation/kernel/dclsettagsyncbits
// DCLSetTagSyncBits is opaque storage with the size and alignment C gives DCLSetTagSyncBits:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type DCLSetTagSyncBits [3]uint64

// See: https://developer.apple.com/documentation/kernel/dclsettagsyncbitsptr
type DCLSetTagSyncBitsPtr = *DCLSetTagSyncBits

// See: https://developer.apple.com/documentation/kernel/dcltimestamp
// DCLTimeStamp is opaque storage with the size and alignment C gives DCLTimeStamp:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type DCLTimeStamp [3]uint64

// See: https://developer.apple.com/documentation/kernel/dcltimestampptr
type DCLTimeStampPtr = *DCLTimeStamp

// See: https://developer.apple.com/documentation/kernel/dcltransferbuffer
// DCLTransferBuffer is opaque storage with the size and alignment C gives DCLTransferBuffer:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type DCLTransferBuffer [5]uint64

// See: https://developer.apple.com/documentation/kernel/dcltransferbufferptr
type DCLTransferBufferPtr = *DCLTransferBuffer

// See: https://developer.apple.com/documentation/kernel/dcltransferpacket
// DCLTransferPacket is opaque storage with the size and alignment C gives DCLTransferPacket:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type DCLTransferPacket [4]uint64

// See: https://developer.apple.com/documentation/kernel/dcltransferpacketptr
type DCLTransferPacketPtr = *DCLTransferPacket

// See: https://developer.apple.com/documentation/kernel/dclupdatedcllist
// DCLUpdateDCLList is opaque storage with the size and alignment C gives DCLUpdateDCLList:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type DCLUpdateDCLList [4]uint64

// See: https://developer.apple.com/documentation/kernel/dclupdatedcllistptr
type DCLUpdateDCLListPtr = *DCLUpdateDCLList

// See: https://developer.apple.com/documentation/kernel/ddmap
// DDMap is opaque storage with the size and alignment C gives DDMap:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type DDMap [8]byte

// See: https://developer.apple.com/documentation/kernel/double_complex
// DOUBLE_COMPLEX is opaque storage with the size and alignment C gives DOUBLE_COMPLEX:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type DOUBLE_COMPLEX [2]uint64

// See: https://developer.apple.com/documentation/kernel/double_complex_split
// DOUBLE_COMPLEX_SPLIT is opaque storage with the size and alignment C gives DOUBLE_COMPLEX_SPLIT:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type DOUBLE_COMPLEX_SPLIT [2]uint64

// See: https://developer.apple.com/documentation/kernel/dpme
// DPME is opaque storage with the size and alignment C gives DPME:
// 512 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 512 into.
type DPME [512]byte

// DSPComplex is used to hold a complex value.
//
// See: https://developer.apple.com/documentation/kernel/dspcomplex
// DSPComplex is opaque storage with the size and alignment C gives DSPComplex:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type DSPComplex [2]uint32

// DSPDoubleComplex is used to hold a double-precision complex value.
//
// See: https://developer.apple.com/documentation/kernel/dspdoublecomplex
// DSPDoubleComplex is opaque storage with the size and alignment C gives DSPDoubleComplex:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type DSPDoubleComplex [2]uint64

// DSPDoubleSplitComplex is used to represent a double-precision complex number when the real and imaginary parts are stored in separate arrays.
//
// See: https://developer.apple.com/documentation/kernel/dspdoublesplitcomplex
// DSPDoubleSplitComplex is opaque storage with the size and alignment C gives DSPDoubleSplitComplex:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type DSPDoubleSplitComplex [2]uint64

// See: https://developer.apple.com/documentation/kernel/dspsplitcomplex
// DSPSplitComplex is opaque storage with the size and alignment C gives DSPSplitComplex:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type DSPSplitComplex [2]uint64

// See: https://developer.apple.com/documentation/kernel/dtentry
type DTEntry = uintptr

// See: https://developer.apple.com/documentation/kernel/dtentryiterator
type DTEntryIterator = uintptr

// See: https://developer.apple.com/documentation/kernel/dtentrynamebuf
type DTEntryNameBuf = int8

// See: https://developer.apple.com/documentation/kernel/dtmemorymaprange
// DTMemoryMapRange is opaque storage with the size and alignment C gives DTMemoryMapRange:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type DTMemoryMapRange [2]uint64

// See: https://developer.apple.com/documentation/kernel/dtpropertyiterator
type DTPropertyIterator = uintptr

// See: https://developer.apple.com/documentation/kernel/dtpropertynamebuf
type DTPropertyNameBuf = int8

// See: https://developer.apple.com/documentation/kernel/dtsavedscopeptr
// DTSavedScopePtr is an unresolved C aggregate typedef.
type DTSavedScopePtr unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dvdauthenticationgrantidinfo
// DVDAuthenticationGrantIDInfo is opaque storage with the size and alignment C gives DVDAuthenticationGrantIDInfo:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type DVDAuthenticationGrantIDInfo [8]byte

// See: https://developer.apple.com/documentation/kernel/dvdauthenticationsuccessflaginfo
// DVDAuthenticationSuccessFlagInfo is opaque storage with the size and alignment C gives DVDAuthenticationSuccessFlagInfo:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type DVDAuthenticationSuccessFlagInfo [8]byte

// See: https://developer.apple.com/documentation/kernel/dvdbooktype
type DVDBookType = byte

// See: https://developer.apple.com/documentation/kernel/dvdcprmregioncode
type DVDCPRMRegionCode = byte

// See: https://developer.apple.com/documentation/kernel/dvdchallengekeyinfo
// DVDChallengeKeyInfo is opaque storage with the size and alignment C gives DVDChallengeKeyInfo:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type DVDChallengeKeyInfo [16]byte

// See: https://developer.apple.com/documentation/kernel/dvdcopyrightinfo
// DVDCopyrightInfo is opaque storage with the size and alignment C gives DVDCopyrightInfo:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type DVDCopyrightInfo [8]byte

// See: https://developer.apple.com/documentation/kernel/dvddiscinfo
// DVDDiscInfo is opaque storage with the size and alignment C gives DVDDiscInfo:
// 34 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 34 into.
type DVDDiscInfo [34]byte

// See: https://developer.apple.com/documentation/kernel/dvddisckeyinfo
// DVDDiscKeyInfo is opaque storage with the size and alignment C gives DVDDiscKeyInfo:
// 2052 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2052 into.
type DVDDiscKeyInfo [2052]byte

// See: https://developer.apple.com/documentation/kernel/dvdfeatures
type DVDFeatures = uint32

// See: https://developer.apple.com/documentation/kernel/dvdkey1info
// DVDKey1Info is opaque storage with the size and alignment C gives DVDKey1Info:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type DVDKey1Info [12]byte

// See: https://developer.apple.com/documentation/kernel/dvdkey2info
// DVDKey2Info is opaque storage with the size and alignment C gives DVDKey2Info:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type DVDKey2Info [12]byte

// See: https://developer.apple.com/documentation/kernel/dvdkeyclass
type DVDKeyClass = byte

// See: https://developer.apple.com/documentation/kernel/dvdkeyformat
type DVDKeyFormat = byte

// See: https://developer.apple.com/documentation/kernel/dvdmanufacturinginfo
// DVDManufacturingInfo is opaque storage with the size and alignment C gives DVDManufacturingInfo:
// 2052 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2052 into.
type DVDManufacturingInfo [2052]byte

// See: https://developer.apple.com/documentation/kernel/dvdmediatype
type DVDMediaType = uint32

// See: https://developer.apple.com/documentation/kernel/dvdphysicalformatinfo
// DVDPhysicalFormatInfo is opaque storage with the size and alignment C gives DVDPhysicalFormatInfo:
// 2052 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2052 into.
type DVDPhysicalFormatInfo [2052]byte

// See: https://developer.apple.com/documentation/kernel/dvdrzoneinfo
// DVDRZoneInfo is opaque storage with the size and alignment C gives DVDRZoneInfo:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type DVDRZoneInfo [36]byte

// See: https://developer.apple.com/documentation/kernel/dvdrzoneinfoaddresstype
type DVDRZoneInfoAddressType = byte

// See: https://developer.apple.com/documentation/kernel/dvdregionplaybackcontrolinfo
// DVDRegionPlaybackControlInfo is opaque storage with the size and alignment C gives DVDRegionPlaybackControlInfo:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type DVDRegionPlaybackControlInfo [8]byte

// See: https://developer.apple.com/documentation/kernel/dvdregionalplaybackcontrolscheme
type DVDRegionalPlaybackControlScheme = byte

// See: https://developer.apple.com/documentation/kernel/dvdstructureformat
type DVDStructureFormat = byte

// See: https://developer.apple.com/documentation/kernel/dvdtitlekeyinfo
// DVDTitleKeyInfo is opaque storage with the size and alignment C gives DVDTitleKeyInfo:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type DVDTitleKeyInfo [12]byte

// See: https://developer.apple.com/documentation/kernel/depthmode
type DepthMode = uint16

// See: https://developer.apple.com/documentation/kernel/devicetreenode
// DeviceTreeNode is opaque storage with the size and alignment C gives DeviceTreeNode:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type DeviceTreeNode [2]uint32

// See: https://developer.apple.com/documentation/kernel/devicetreenodeproperty
// DeviceTreeNodeProperty is opaque storage with the size and alignment C gives DeviceTreeNodeProperty:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type DeviceTreeNodeProperty [9]uint32

// See: https://developer.apple.com/documentation/kernel/displayidtype
type DisplayIDType = uint32

// See: https://developer.apple.com/documentation/kernel/displaymodeid
type DisplayModeID = int32

// See: https://developer.apple.com/documentation/kernel/driverdescversion
type DriverDescVersion = uint32

// See: https://developer.apple.com/documentation/kernel/driverdescription
// DriverDescription is opaque storage with the size and alignment C gives DriverDescription:
// 128 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 128 into.
type DriverDescription [32]uint32

// See: https://developer.apple.com/documentation/kernel/driverdescriptionptr
type DriverDescriptionPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/driverosruntime
// DriverOSRuntime is opaque storage with the size and alignment C gives DriverOSRuntime:
// 68 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 68 into.
type DriverOSRuntime [17]uint32

// See: https://developer.apple.com/documentation/kernel/driverosruntimeptr
type DriverOSRuntimePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/driverosservice
// DriverOSService is opaque storage with the size and alignment C gives DriverOSService:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type DriverOSService [4]uint32

// See: https://developer.apple.com/documentation/kernel/driverosserviceptr
type DriverOSServicePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/driverserviceinfo
// DriverServiceInfo is opaque storage with the size and alignment C gives DriverServiceInfo:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type DriverServiceInfo [3]uint32

// See: https://developer.apple.com/documentation/kernel/driverserviceinfoptr
type DriverServiceInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/drivertype
// DriverType is opaque storage with the size and alignment C gives DriverType:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type DriverType [36]byte

// See: https://developer.apple.com/documentation/kernel/drivertypeptr
type DriverTypePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/efi_boolean
type EFI_BOOLEAN = byte

// See: https://developer.apple.com/documentation/kernel/efi_char16
type EFI_CHAR16 = int16

// See: https://developer.apple.com/documentation/kernel/efi_char32
type EFI_CHAR32 = int32

// See: https://developer.apple.com/documentation/kernel/efi_char64
type EFI_CHAR64 = int64

// See: https://developer.apple.com/documentation/kernel/efi_char8
type EFI_CHAR8 = int8

// See: https://developer.apple.com/documentation/kernel/efi_configuration_table_32
// EFI_CONFIGURATION_TABLE_32 is opaque storage with the size and alignment C gives EFI_CONFIGURATION_TABLE_32:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type EFI_CONFIGURATION_TABLE_32 [5]uint32

// See: https://developer.apple.com/documentation/kernel/efi_configuration_table_64
// EFI_CONFIGURATION_TABLE_64 is opaque storage with the size and alignment C gives EFI_CONFIGURATION_TABLE_64:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type EFI_CONFIGURATION_TABLE_64 [3]uint64

// See: https://developer.apple.com/documentation/kernel/efi_guid
// EFI_GUID is opaque storage with the size and alignment C gives EFI_GUID:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type EFI_GUID [4]uint32

// See: https://developer.apple.com/documentation/kernel/efi_handle32
type EFI_HANDLE32 = uint32

// See: https://developer.apple.com/documentation/kernel/efi_handle64
type EFI_HANDLE64 = uint64

// See: https://developer.apple.com/documentation/kernel/efi_int16
type EFI_INT16 = int16

// See: https://developer.apple.com/documentation/kernel/efi_int32
type EFI_INT32 = int32

// See: https://developer.apple.com/documentation/kernel/efi_int64
type EFI_INT64 = int64

// See: https://developer.apple.com/documentation/kernel/efi_int8
type EFI_INT8 = int8

// See: https://developer.apple.com/documentation/kernel/efi_memory_descriptor
// EFI_MEMORY_DESCRIPTOR is opaque storage with the size and alignment C gives EFI_MEMORY_DESCRIPTOR:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type EFI_MEMORY_DESCRIPTOR [5]uint64

// See: https://developer.apple.com/documentation/kernel/efi_memory_type
type EFI_MEMORY_TYPE = uint32

// See: https://developer.apple.com/documentation/kernel/efi_physical_address
type EFI_PHYSICAL_ADDRESS = uint64

// See: https://developer.apple.com/documentation/kernel/efi_ptr32
type EFI_PTR32 = uint32

// See: https://developer.apple.com/documentation/kernel/efi_ptr64
type EFI_PTR64 = uint64

// See: https://developer.apple.com/documentation/kernel/efi_reset_type
type EFI_RESET_TYPE = uint32

// See: https://developer.apple.com/documentation/kernel/efi_runtime_services_32
// EFI_RUNTIME_SERVICES_32 is opaque storage with the size and alignment C gives EFI_RUNTIME_SERVICES_32:
// 72 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 72 into.
type EFI_RUNTIME_SERVICES_32 [9]uint64

// See: https://developer.apple.com/documentation/kernel/efi_runtime_services_64
// EFI_RUNTIME_SERVICES_64 is opaque storage with the size and alignment C gives EFI_RUNTIME_SERVICES_64:
// 112 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 112 into.
type EFI_RUNTIME_SERVICES_64 [14]uint64

// See: https://developer.apple.com/documentation/kernel/efi_status
type EFI_STATUS = uint32

// See: https://developer.apple.com/documentation/kernel/efi_system_table_32
// EFI_SYSTEM_TABLE_32 is opaque storage with the size and alignment C gives EFI_SYSTEM_TABLE_32:
// 72 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 72 into.
type EFI_SYSTEM_TABLE_32 [9]uint64

// See: https://developer.apple.com/documentation/kernel/efi_system_table_64
// EFI_SYSTEM_TABLE_64 is opaque storage with the size and alignment C gives EFI_SYSTEM_TABLE_64:
// 120 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 120 into.
type EFI_SYSTEM_TABLE_64 [15]uint64

// See: https://developer.apple.com/documentation/kernel/efi_table_header
// EFI_TABLE_HEADER is opaque storage with the size and alignment C gives EFI_TABLE_HEADER:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type EFI_TABLE_HEADER [3]uint64

// See: https://developer.apple.com/documentation/kernel/efi_time
// EFI_TIME is opaque storage with the size and alignment C gives EFI_TIME:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type EFI_TIME [4]uint32

// See: https://developer.apple.com/documentation/kernel/efi_time_capabilities
// EFI_TIME_CAPABILITIES is opaque storage with the size and alignment C gives EFI_TIME_CAPABILITIES:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type EFI_TIME_CAPABILITIES [3]uint32

// See: https://developer.apple.com/documentation/kernel/efi_uint16
type EFI_UINT16 = uint16

// See: https://developer.apple.com/documentation/kernel/efi_uint32
type EFI_UINT32 = uint32

// See: https://developer.apple.com/documentation/kernel/efi_uint64
type EFI_UINT64 = uint64

// See: https://developer.apple.com/documentation/kernel/efi_uint8
type EFI_UINT8 = byte

// See: https://developer.apple.com/documentation/kernel/efi_uintn
type EFI_UINTN = uint32

// See: https://developer.apple.com/documentation/kernel/efi_virtual_address
type EFI_VIRTUAL_ADDRESS = uint64

// See: https://developer.apple.com/documentation/kernel/evscreen
type EVScreen = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/exbrightmessage
// EXBrightMessage is opaque storage with the size and alignment C gives EXBrightMessage:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type EXBrightMessage [3]uint32

// See: https://developer.apple.com/documentation/kernel/exbrightmessagetype
type EXBrightMessageType = uint32

// See: https://developer.apple.com/documentation/kernel/exdisplaypipehealthrecord
// EXDisplayPipeHealthRecord is opaque storage with the size and alignment C gives EXDisplayPipeHealthRecord:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type EXDisplayPipeHealthRecord [2]uint64

// See: https://developer.apple.com/documentation/kernel/exdisplaypipehealthreport
// EXDisplayPipeHealthReport is opaque storage with the size and alignment C gives EXDisplayPipeHealthReport:
// 88 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 88 into.
type EXDisplayPipeHealthReport [11]uint64

// See: https://developer.apple.com/documentation/kernel/exdisplaypipeindicator
type EXDisplayPipeIndicator = uint32

// See: https://developer.apple.com/documentation/kernel/exdisplaypipeindicatorparams
// EXDisplayPipeIndicatorParams is opaque storage with the size and alignment C gives EXDisplayPipeIndicatorParams:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type EXDisplayPipeIndicatorParams [2]uint32

// See: https://developer.apple.com/documentation/kernel/exdisplaypipesecuretestatus
// EXDisplayPipeSecureTEStatus is opaque storage with the size and alignment C gives EXDisplayPipeSecureTEStatus:
// 544 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 544 into.
type EXDisplayPipeSecureTEStatus [68]uint64

// See: https://developer.apple.com/documentation/kernel/exdisplaypipestatus
// EXDisplayPipeStatus is opaque storage with the size and alignment C gives EXDisplayPipeStatus:
// 352 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 352 into.
type EXDisplayPipeStatus [44]uint64

// See: https://developer.apple.com/documentation/kernel/efimemoryrange
// EfiMemoryRange is opaque storage with the size and alignment C gives EfiMemoryRange:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type EfiMemoryRange [5]uint64

// See: https://developer.apple.com/documentation/kernel/evcmd
type EvCmd = int

// See: https://developer.apple.com/documentation/kernel/evglobals
// EvGlobals is opaque storage with the size and alignment C gives EvGlobals:
// 23264 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 23264 into.
type EvGlobals [5816]uint32

// See: https://developer.apple.com/documentation/kernel/evoffsets
// EvOffsets is opaque storage with the size and alignment C gives EvOffsets:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type EvOffsets [2]uint32

// See: https://developer.apple.com/documentation/kernel/extendedsensecode
type ExtendedSenseCode = byte

// FFTDirection is specifies whether to perform a forward or inverse FFT.
//
// See: https://developer.apple.com/documentation/kernel/fftdirection
type FFTDirection = int32

// FFTRadix is the size of the FFT decomposition.
//
// See: https://developer.apple.com/documentation/kernel/fftradix
type FFTRadix = int32

// FFTSetup is an opaque type that contains setup information for a given FFT transform.
//
// See: https://developer.apple.com/documentation/kernel/fftsetup
type FFTSetup = uintptr

// FFTSetupD is an opaque type that contains setup information for a given double-precision FFT transform.
//
// See: https://developer.apple.com/documentation/kernel/fftsetupd
type FFTSetupD = uintptr

// See: https://developer.apple.com/documentation/kernel/fwaddress
// FWAddress is opaque storage with the size and alignment C gives FWAddress:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type FWAddress [2]uint32

// See: https://developer.apple.com/documentation/kernel/fwaddressptr
type FWAddressPtr = uintptr

// See: https://developer.apple.com/documentation/kernel/fwclientcommandid
type FWClientCommandID = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fwisochchannelforcestopnotificationproc
type FWIsochChannelForceStopNotificationProc = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fwisochchannelforcestopnotificationprocptr
type FWIsochChannelForceStopNotificationProcPtr = *uintptr

// See: https://developer.apple.com/documentation/kernel/fwmultiisochreceivelistenerparams
// FWMultiIsochReceiveListenerParams is opaque storage with the size and alignment C gives FWMultiIsochReceiveListenerParams:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type FWMultiIsochReceiveListenerParams [3]uint32

// See: https://developer.apple.com/documentation/kernel/fwsbp2logincompleteparams
// FWSBP2LoginCompleteParams is opaque storage with the size and alignment C gives FWSBP2LoginCompleteParams:
// 1 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 1 into.
type FWSBP2LoginCompleteParams [1]byte

// See: https://developer.apple.com/documentation/kernel/fwsbp2logincompleteparamsptr
type FWSBP2LoginCompleteParamsPtr = uintptr

// See: https://developer.apple.com/documentation/kernel/fwsbp2loginresponse
// FWSBP2LoginResponse is opaque storage with the size and alignment C gives FWSBP2LoginResponse:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type FWSBP2LoginResponse [4]uint32

// See: https://developer.apple.com/documentation/kernel/fwsbp2loginresponseptr
type FWSBP2LoginResponsePtr = uintptr

// See: https://developer.apple.com/documentation/kernel/fwsbp2logoutcompleteparams
// FWSBP2LogoutCompleteParams is opaque storage with the size and alignment C gives FWSBP2LogoutCompleteParams:
// 1 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 1 into.
type FWSBP2LogoutCompleteParams [1]byte

// See: https://developer.apple.com/documentation/kernel/fwsbp2logoutcompleteparamsptr
type FWSBP2LogoutCompleteParamsPtr = uintptr

// See: https://developer.apple.com/documentation/kernel/fwsbp2notifyparams
// FWSBP2NotifyParams is opaque storage with the size and alignment C gives FWSBP2NotifyParams:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type FWSBP2NotifyParams [4]uint64

// See: https://developer.apple.com/documentation/kernel/fwsbp2notifyparamsptr
type FWSBP2NotifyParamsPtr = uintptr

// See: https://developer.apple.com/documentation/kernel/fwsbp2reconnectparams
// FWSBP2ReconnectParams is opaque storage with the size and alignment C gives FWSBP2ReconnectParams:
// 1 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 1 into.
type FWSBP2ReconnectParams [1]byte

// See: https://developer.apple.com/documentation/kernel/fwsbp2reconnectparamsptr
type FWSBP2ReconnectParamsPtr = uintptr

// See: https://developer.apple.com/documentation/kernel/fwsbp2statusblock
// FWSBP2StatusBlock is opaque storage with the size and alignment C gives FWSBP2StatusBlock:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type FWSBP2StatusBlock [8]uint32

// See: https://developer.apple.com/documentation/kernel/fixed
type Fixed = int32

// See: https://developer.apple.com/documentation/kernel/fixedptr
type FixedPtr = *int32

// Float32 is convenience type that represent a 32-bit floating point number.
//
// See: https://developer.apple.com/documentation/kernel/float32
type Float32 = float32

// Float64 is convenience type that represent a 64-bit floating point number.
//
// See: https://developer.apple.com/documentation/kernel/float64
type Float64 = float64

// See: https://developer.apple.com/documentation/kernel/fndrdirinfo
// FndrDirInfo is opaque storage with the size and alignment C gives FndrDirInfo:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type FndrDirInfo [8]uint16

// See: https://developer.apple.com/documentation/kernel/fndrfileinfo
// FndrFileInfo is opaque storage with the size and alignment C gives FndrFileInfo:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type FndrFileInfo [8]uint16

// See: https://developer.apple.com/documentation/kernel/fndropaqueinfo
// FndrOpaqueInfo is opaque storage with the size and alignment C gives FndrOpaqueInfo:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type FndrOpaqueInfo [8]uint16

// See: https://developer.apple.com/documentation/kernel/fourcharcode
type FourCharCode = uint32

// Fract is represents a type used by the Compression and Decompression API.
//
// See: https://developer.apple.com/documentation/kernel/fract
type Fract = int32

// See: https://developer.apple.com/documentation/kernel/fractptr
type FractPtr = *int32

// See: https://developer.apple.com/documentation/kernel/gammatableid
type GammaTableID = uint32

// See: https://developer.apple.com/documentation/kernel/gammatbl
// GammaTbl is opaque storage with the size and alignment C gives GammaTbl:
// 14 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 14 into.
type GammaTbl [7]uint16

// See: https://developer.apple.com/documentation/kernel/gammatblptr
type GammaTblPtr = *GammaTbl

// See: https://developer.apple.com/documentation/kernel/hfscatalogfile
// HFSCatalogFile is opaque storage with the size and alignment C gives HFSCatalogFile:
// 102 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 102 into.
type HFSCatalogFile [51]uint16

// See: https://developer.apple.com/documentation/kernel/hfscatalogfolder
// HFSCatalogFolder is opaque storage with the size and alignment C gives HFSCatalogFolder:
// 70 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 70 into.
type HFSCatalogFolder [35]uint16

// See: https://developer.apple.com/documentation/kernel/hfscatalogkey
// HFSCatalogKey is opaque storage with the size and alignment C gives HFSCatalogKey:
// 38 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 38 into.
type HFSCatalogKey [19]uint16

// See: https://developer.apple.com/documentation/kernel/hfscatalogthread
// HFSCatalogThread is opaque storage with the size and alignment C gives HFSCatalogThread:
// 46 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 46 into.
type HFSCatalogThread [23]uint16

// See: https://developer.apple.com/documentation/kernel/hfsextentdescriptor
// HFSExtentDescriptor is opaque storage with the size and alignment C gives HFSExtentDescriptor:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type HFSExtentDescriptor [2]uint16

// See: https://developer.apple.com/documentation/kernel/hfsextentkey
// HFSExtentKey is opaque storage with the size and alignment C gives HFSExtentKey:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type HFSExtentKey [4]uint16

// See: https://developer.apple.com/documentation/kernel/hfsextentrecord
type HFSExtentRecord = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hfsmasterdirectoryblock
// HFSMasterDirectoryBlock is opaque storage with the size and alignment C gives HFSMasterDirectoryBlock:
// 162 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 162 into.
type HFSMasterDirectoryBlock [81]uint16

// See: https://developer.apple.com/documentation/kernel/hfsplusattrdata
// HFSPlusAttrData is opaque storage with the size and alignment C gives HFSPlusAttrData:
// 18 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 18 into.
type HFSPlusAttrData [9]uint16

// See: https://developer.apple.com/documentation/kernel/hfsplusattrextents
// HFSPlusAttrExtents is opaque storage with the size and alignment C gives HFSPlusAttrExtents:
// 72 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 72 into.
type HFSPlusAttrExtents [36]uint16

// See: https://developer.apple.com/documentation/kernel/hfsplusattrforkdata
// HFSPlusAttrForkData is opaque storage with the size and alignment C gives HFSPlusAttrForkData:
// 88 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 88 into.
type HFSPlusAttrForkData [44]uint16

// See: https://developer.apple.com/documentation/kernel/hfsplusattrinlinedata
// HFSPlusAttrInlineData is opaque storage with the size and alignment C gives HFSPlusAttrInlineData:
// 14 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 14 into.
type HFSPlusAttrInlineData [7]uint16

// See: https://developer.apple.com/documentation/kernel/hfsplusattrkey
// HFSPlusAttrKey is opaque storage with the size and alignment C gives HFSPlusAttrKey:
// 268 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 268 into.
type HFSPlusAttrKey [134]uint16

// See: https://developer.apple.com/documentation/kernel/hfsplusbsdinfo
// HFSPlusBSDInfo is opaque storage with the size and alignment C gives HFSPlusBSDInfo:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type HFSPlusBSDInfo [8]uint16

// See: https://developer.apple.com/documentation/kernel/hfspluscatalogfile
// HFSPlusCatalogFile is opaque storage with the size and alignment C gives HFSPlusCatalogFile:
// 248 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 248 into.
type HFSPlusCatalogFile [124]uint16

// See: https://developer.apple.com/documentation/kernel/hfspluscatalogfolder
// HFSPlusCatalogFolder is opaque storage with the size and alignment C gives HFSPlusCatalogFolder:
// 88 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 88 into.
type HFSPlusCatalogFolder [44]uint16

// See: https://developer.apple.com/documentation/kernel/hfspluscatalogkey
// HFSPlusCatalogKey is opaque storage with the size and alignment C gives HFSPlusCatalogKey:
// 518 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 518 into.
type HFSPlusCatalogKey [259]uint16

// See: https://developer.apple.com/documentation/kernel/hfspluscatalogthread
// HFSPlusCatalogThread is opaque storage with the size and alignment C gives HFSPlusCatalogThread:
// 520 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 520 into.
type HFSPlusCatalogThread [260]uint16

// See: https://developer.apple.com/documentation/kernel/hfsplusextentdescriptor
// HFSPlusExtentDescriptor is opaque storage with the size and alignment C gives HFSPlusExtentDescriptor:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type HFSPlusExtentDescriptor [4]uint16

// See: https://developer.apple.com/documentation/kernel/hfsplusextentkey
// HFSPlusExtentKey is opaque storage with the size and alignment C gives HFSPlusExtentKey:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type HFSPlusExtentKey [6]uint16

// See: https://developer.apple.com/documentation/kernel/hfsplusextentrecord
type HFSPlusExtentRecord = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hfsplusforkdata
// HFSPlusForkData is opaque storage with the size and alignment C gives HFSPlusForkData:
// 80 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 80 into.
type HFSPlusForkData [40]uint16

// See: https://developer.apple.com/documentation/kernel/hfsplusvolumeheader
// HFSPlusVolumeHeader is opaque storage with the size and alignment C gives HFSPlusVolumeHeader:
// 512 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 512 into.
type HFSPlusVolumeHeader [256]uint16

// See: https://developer.apple.com/documentation/kernel/hfsunistr255
// HFSUniStr255 is opaque storage with the size and alignment C gives HFSUniStr255:
// 512 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 512 into.
type HFSUniStr255 [256]uint16

// See: https://developer.apple.com/documentation/kernel/hidreportcommandtype
type HIDReportCommandType = int

// See: https://developer.apple.com/documentation/kernel/handle
type Handle = *Ptr

// See: https://developer.apple.com/documentation/kernel/hardwarecursordescriptorptr
type HardwareCursorDescriptorPtr = *HardwareCursorDescriptorRec

// See: https://developer.apple.com/documentation/kernel/hardwarecursordescriptorrec
type HardwareCursorDescriptorRec = IOHardwareCursorDescriptor

// See: https://developer.apple.com/documentation/kernel/hardwarecursorinfoptr
type HardwareCursorInfoPtr = *HardwareCursorInfoRec

// See: https://developer.apple.com/documentation/kernel/hardwarecursorinforec
type HardwareCursorInfoRec = IOHardwareCursorInfo

// IIRChannel is constants that specify which channels a stereo biquadratic filter operates.
//
// See: https://developer.apple.com/documentation/kernel/iirchannel
type IIRChannel = int32

// See: https://developer.apple.com/documentation/kernel/ioacpiaddressspaceid
type IOACPIAddressSpaceID = uint32

// See: https://developer.apple.com/documentation/kernel/ioatacompletionfunction
type IOATACompletionFunction = *objc.ID

// See: https://developer.apple.com/documentation/kernel/ioataregptr16
type IOATARegPtr16 = *IOATAReg16

// See: https://developer.apple.com/documentation/kernel/ioataregptr32
type IOATARegPtr32 = *IOATAReg32

// See: https://developer.apple.com/documentation/kernel/ioataregptr8
type IOATARegPtr8 = *IOATAReg8

// See: https://developer.apple.com/documentation/kernel/ioaccelbounds
// IOAccelBounds is opaque storage with the size and alignment C gives IOAccelBounds:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOAccelBounds [4]uint16

// See: https://developer.apple.com/documentation/kernel/ioacceldeviceregion
// IOAccelDeviceRegion is opaque storage with the size and alignment C gives IOAccelDeviceRegion:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type IOAccelDeviceRegion [3]uint32

// See: https://developer.apple.com/documentation/kernel/ioaccelid
type IOAccelID = int32

// See: https://developer.apple.com/documentation/kernel/ioaccelsize
// IOAccelSize is opaque storage with the size and alignment C gives IOAccelSize:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type IOAccelSize [2]uint16

// See: https://developer.apple.com/documentation/kernel/ioaccelsurfaceinformation
// IOAccelSurfaceInformation is opaque storage with the size and alignment C gives IOAccelSurfaceInformation:
// 88 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 88 into.
type IOAccelSurfaceInformation [11]uint64

// See: https://developer.apple.com/documentation/kernel/ioaccelsurfacereaddata
// IOAccelSurfaceReadData is opaque storage with the size and alignment C gives IOAccelSurfaceReadData:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type IOAccelSurfaceReadData [4]uint64

// See: https://developer.apple.com/documentation/kernel/ioaccelsurfacescaling
// IOAccelSurfaceScaling is opaque storage with the size and alignment C gives IOAccelSurfaceScaling:
// 44 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 44 into.
type IOAccelSurfaceScaling [11]uint32

// See: https://developer.apple.com/documentation/kernel/ioaddressrange
// IOAddressRange is opaque storage with the size and alignment C gives IOAddressRange:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type IOAddressRange [2]uint64

// See: https://developer.apple.com/documentation/kernel/ioalignment
type IOAlignment = uint32

// See: https://developer.apple.com/documentation/kernel/ioappletimingid
type IOAppleTimingID = uint32

// See: https://developer.apple.com/documentation/kernel/ioasyncmethod
type IOAsyncMethod = *objc.ID

// See: https://developer.apple.com/documentation/kernel/ioaudiobufferdatadescriptor
// IOAudioBufferDataDescriptor is opaque storage with the size and alignment C gives IOAudioBufferDataDescriptor:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type IOAudioBufferDataDescriptor [5]uint32

// See: https://developer.apple.com/documentation/kernel/ioaudioclientbuffer
// IOAudioClientBuffer is opaque storage with the size and alignment C gives IOAudioClientBuffer:
// 1 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 1 into.
type IOAudioClientBuffer [1]byte

// See: https://developer.apple.com/documentation/kernel/ioaudioclientbuffer64
// IOAudioClientBuffer64 is opaque storage with the size and alignment C gives IOAudioClientBuffer64:
// 1 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 1 into.
type IOAudioClientBuffer64 [1]byte

// See: https://developer.apple.com/documentation/kernel/ioaudioclientbufferextendedinfo
// IOAudioClientBufferExtendedInfo is opaque storage with the size and alignment C gives IOAudioClientBufferExtendedInfo:
// 1 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 1 into.
type IOAudioClientBufferExtendedInfo [1]byte

// See: https://developer.apple.com/documentation/kernel/ioaudioclientbufferextendedinfo64
// IOAudioClientBufferExtendedInfo64 is opaque storage with the size and alignment C gives IOAudioClientBufferExtendedInfo64:
// 1 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 1 into.
type IOAudioClientBufferExtendedInfo64 [1]byte

// IOAudioDevicePowerState is identifies the power state of the audio device.
//
// See: https://developer.apple.com/documentation/kernel/ioaudiodevicepowerstate
type IOAudioDevicePowerState = uint32

// See: https://developer.apple.com/documentation/kernel/ioaudioenginenotifications
type IOAudioEngineNotifications = int

// IOAudioEnginePosition is represents a position in an audio audio engine.
//
// See: https://developer.apple.com/documentation/kernel/ioaudioengineposition
// IOAudioEnginePosition is opaque storage with the size and alignment C gives IOAudioEnginePosition:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOAudioEnginePosition [2]uint32

// See: https://developer.apple.com/documentation/kernel/ioaudioenginetraps
type IOAudioEngineTraps = int

// See: https://developer.apple.com/documentation/kernel/ioaudiosampleintervaldescriptor
// IOAudioSampleIntervalDescriptor is opaque storage with the size and alignment C gives IOAudioSampleIntervalDescriptor:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOAudioSampleIntervalDescriptor [2]uint32

// See: https://developer.apple.com/documentation/kernel/ioaudiosamplerate
// IOAudioSampleRate is opaque storage with the size and alignment C gives IOAudioSampleRate:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOAudioSampleRate [2]uint32

// See: https://developer.apple.com/documentation/kernel/ioaudiostreamdatadescriptor
// IOAudioStreamDataDescriptor is opaque storage with the size and alignment C gives IOAudioStreamDataDescriptor:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type IOAudioStreamDataDescriptor [3]uint32

// See: https://developer.apple.com/documentation/kernel/ioaudiostreamformat
// IOAudioStreamFormat is opaque storage with the size and alignment C gives IOAudioStreamFormat:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IOAudioStreamFormat [6]uint32

// See: https://developer.apple.com/documentation/kernel/ioaudiostreamformatextension
// IOAudioStreamFormatExtension is opaque storage with the size and alignment C gives IOAudioStreamFormatExtension:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type IOAudioStreamFormatExtension [4]uint32

// See: https://developer.apple.com/documentation/kernel/ioaudiotimestamp
// IOAudioTimeStamp is opaque storage with the size and alignment C gives IOAudioTimeStamp:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type IOAudioTimeStamp [8]uint64

// See: https://developer.apple.com/documentation/kernel/ioblitcopyrectangle
// IOBlitCopyRectangle is opaque storage with the size and alignment C gives IOBlitCopyRectangle:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IOBlitCopyRectangle [6]uint32

// See: https://developer.apple.com/documentation/kernel/ioblitcopyrectangles
// IOBlitCopyRectangles is opaque storage with the size and alignment C gives IOBlitCopyRectangles:
// 116 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 116 into.
type IOBlitCopyRectangles [29]uint32

// See: https://developer.apple.com/documentation/kernel/ioblitcopyregion
// IOBlitCopyRegion is opaque storage with the size and alignment C gives IOBlitCopyRegion:
// 104 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 104 into.
type IOBlitCopyRegion [13]uint64

// See: https://developer.apple.com/documentation/kernel/ioblitcursor
// IOBlitCursor is opaque storage with the size and alignment C gives IOBlitCursor:
// 104 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 104 into.
type IOBlitCursor [26]uint32

// See: https://developer.apple.com/documentation/kernel/ioblitmemory
// IOBlitMemory is opaque storage with the size and alignment C gives IOBlitMemory:
// 120 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 120 into.
type IOBlitMemory [15]uint64

// See: https://developer.apple.com/documentation/kernel/ioblitmemoryref
type IOBlitMemoryRef uintptr

// See: https://developer.apple.com/documentation/kernel/ioblitoperation
// IOBlitOperation is opaque storage with the size and alignment C gives IOBlitOperation:
// 88 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 88 into.
type IOBlitOperation [22]uint32

// See: https://developer.apple.com/documentation/kernel/ioblitrectangle
// IOBlitRectangle is opaque storage with the size and alignment C gives IOBlitRectangle:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type IOBlitRectangle [4]uint32

// See: https://developer.apple.com/documentation/kernel/ioblitrectangles
// IOBlitRectangles is opaque storage with the size and alignment C gives IOBlitRectangles:
// 108 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 108 into.
type IOBlitRectangles [27]uint32

// See: https://developer.apple.com/documentation/kernel/ioblitscanlines
// IOBlitScanlines is opaque storage with the size and alignment C gives IOBlitScanlines:
// 108 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 108 into.
type IOBlitScanlines [27]uint32

// See: https://developer.apple.com/documentation/kernel/ioblitsourcetype
type IOBlitSourceType = uint32

// See: https://developer.apple.com/documentation/kernel/ioblitsurface
// IOBlitSurface is opaque storage with the size and alignment C gives IOBlitSurface:
// 120 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 120 into.
type IOBlitSurface [15]uint64

// See: https://developer.apple.com/documentation/kernel/ioblittype
type IOBlitType = uint32

// See: https://developer.apple.com/documentation/kernel/ioblitvertex
// IOBlitVertex is opaque storage with the size and alignment C gives IOBlitVertex:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOBlitVertex [2]uint32

// See: https://developer.apple.com/documentation/kernel/ioblitvertices
// IOBlitVertices is opaque storage with the size and alignment C gives IOBlitVertices:
// 108 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 108 into.
type IOBlitVertices [27]uint32

// See: https://developer.apple.com/documentation/kernel/iobytecount
type IOByteCount = uint

// See: https://developer.apple.com/documentation/kernel/iobytecount32
type IOByteCount32 = uint32

// See: https://developer.apple.com/documentation/kernel/iobytecount64
type IOByteCount64 = uint64

// See: https://developer.apple.com/documentation/kernel/iocachemode
type IOCacheMode = uint32

// See: https://developer.apple.com/documentation/kernel/iocolorcomponent
type IOColorComponent = uint16

// See: https://developer.apple.com/documentation/kernel/iocolorentry
// IOColorEntry is opaque storage with the size and alignment C gives IOColorEntry:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOColorEntry [4]uint16

// See: https://developer.apple.com/documentation/kernel/iocommandcode
type IOCommandCode = uint32

// See: https://developer.apple.com/documentation/kernel/iocommandid
type IOCommandID = uintptr

// See: https://developer.apple.com/documentation/kernel/iocommandkind
type IOCommandKind = uint32

// See: https://developer.apple.com/documentation/kernel/iodataqueueclientdequeueentryblock
type IODataQueueClientDequeueEntryBlock = func(data unsafe.Pointer, dataSize uintptr)

// See: https://developer.apple.com/documentation/kernel/iodataqueueclientenqueueentryblock
type IODataQueueClientEnqueueEntryBlock = func(data unsafe.Pointer, dataSize uintptr)

// See: https://developer.apple.com/documentation/kernel/iodebuggerlockstate
type IODebuggerLockState = uint32

// See: https://developer.apple.com/documentation/kernel/iodetailedtiminginformation
type IODetailedTimingInformation = IODetailedTimingInformationV2

// See: https://developer.apple.com/documentation/kernel/iodetailedtiminginformationv1
// IODetailedTimingInformationV1 is opaque storage with the size and alignment C gives IODetailedTimingInformationV1:
// 44 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 44 into.
type IODetailedTimingInformationV1 [11]uint32

// See: https://developer.apple.com/documentation/kernel/iodetailedtiminginformationv2
// IODetailedTimingInformationV2 is opaque storage with the size and alignment C gives IODetailedTimingInformationV2:
// 160 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 160 into.
type IODetailedTimingInformationV2 [20]uint64

// See: https://developer.apple.com/documentation/kernel/iodirection
type IODirection = uint32

// See: https://developer.apple.com/documentation/kernel/iodispatchaction
type IODispatchAction = func() Kern_return_t

// See: https://developer.apple.com/documentation/kernel/iodispatchblock
type IODispatchBlock = func()

// See: https://developer.apple.com/documentation/kernel/iodispatchqueuecancelhandler
type IODispatchQueueCancelHandler = func()

// See: https://developer.apple.com/documentation/kernel/iodispatchqueuename
type IODispatchQueueName = int8

// See: https://developer.apple.com/documentation/kernel/iodispatchsourcecancelhandler
type IODispatchSourceCancelHandler = func()

// See: https://developer.apple.com/documentation/kernel/iodisplaymodeid
type IODisplayModeID = int32

// See: https://developer.apple.com/documentation/kernel/iodisplaymodeinformation
// IODisplayModeInformation is opaque storage with the size and alignment C gives IODisplayModeInformation:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type IODisplayModeInformation [9]uint32

// See: https://developer.apple.com/documentation/kernel/iodisplayproductid
type IODisplayProductID = uint32

// See: https://developer.apple.com/documentation/kernel/iodisplayscalerinformation
// IODisplayScalerInformation is opaque storage with the size and alignment C gives IODisplayScalerInformation:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type IODisplayScalerInformation [12]uint32

// See: https://developer.apple.com/documentation/kernel/iodisplaytimingrange
type IODisplayTimingRange = IODisplayTimingRangeV2

// See: https://developer.apple.com/documentation/kernel/iodisplaytimingrangev1
// IODisplayTimingRangeV1 is opaque storage with the size and alignment C gives IODisplayTimingRangeV1:
// 240 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 240 into.
type IODisplayTimingRangeV1 [30]uint64

// See: https://developer.apple.com/documentation/kernel/iodisplaytimingrangev2
// IODisplayTimingRangeV2 is opaque storage with the size and alignment C gives IODisplayTimingRangeV2:
// 312 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 312 into.
type IODisplayTimingRangeV2 [39]uint64

// See: https://developer.apple.com/documentation/kernel/iodisplayvendorid
type IODisplayVendorID = uint32

// See: https://developer.apple.com/documentation/kernel/ioenetmulticastmode
type IOEnetMulticastMode = bool

// See: https://developer.apple.com/documentation/kernel/ioenetpromiscuousmode
type IOEnetPromiscuousMode = bool

// See: https://developer.apple.com/documentation/kernel/ioethernetcontrolleravbstate
type IOEthernetControllerAVBState = uint32

// See: https://developer.apple.com/documentation/kernel/ioethernetcontrolleravbstateevent
type IOEthernetControllerAVBStateEvent = uint32

// See: https://developer.apple.com/documentation/kernel/ioethernetcontrolleravbtimesyncsupport
type IOEthernetControllerAVBTimeSyncSupport = uint32

// See: https://developer.apple.com/documentation/kernel/iofbcursorcontrolattribute
// IOFBCursorControlAttribute is opaque storage with the size and alignment C gives IOFBCursorControlAttribute:
// 144 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 144 into.
type IOFBCursorControlAttribute [18]uint64

// See: https://developer.apple.com/documentation/kernel/iofbcursorcontrolcallouts
// IOFBCursorControlCallouts is opaque storage with the size and alignment C gives IOFBCursorControlCallouts:
// 136 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 136 into.
type IOFBCursorControlCallouts [17]uint64

// See: https://developer.apple.com/documentation/kernel/iofbcursorref
type IOFBCursorRef = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iofbdplinkconfig
// IOFBDPLinkConfig is opaque storage with the size and alignment C gives IOFBDPLinkConfig:
// 28 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 28 into.
type IOFBDPLinkConfig [14]uint16

// See: https://developer.apple.com/documentation/kernel/iofbdisplaymodedescription
// IOFBDisplayModeDescription is opaque storage with the size and alignment C gives IOFBDisplayModeDescription:
// 204 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 204 into.
type IOFBDisplayModeDescription [51]uint32

// See: https://developer.apple.com/documentation/kernel/iofbhdrmetadata
// IOFBHDRMetaData is opaque storage with the size and alignment C gives IOFBHDRMetaData:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type IOFBHDRMetaData [8]uint64

// See: https://developer.apple.com/documentation/kernel/iofbhdrmetadatav1
// IOFBHDRMetaDataV1 is opaque storage with the size and alignment C gives IOFBHDRMetaDataV1:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type IOFBHDRMetaDataV1 [8]uint64

// See: https://developer.apple.com/documentation/kernel/iofwavcasynccommandstate
type IOFWAVCAsyncCommandState = int

// See: https://developer.apple.com/documentation/kernel/iofwavcplugtypes
type IOFWAVCPlugTypes = int

// See: https://developer.apple.com/documentation/kernel/iofwavcsubunitplugmessages
type IOFWAVCSubunitPlugMessages = int

// See: https://developer.apple.com/documentation/kernel/iofwduplicateguidrec
// IOFWDuplicateGUIDRec is opaque storage with the size and alignment C gives IOFWDuplicateGUIDRec:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IOFWDuplicateGUIDRec [3]uint64

// See: https://developer.apple.com/documentation/kernel/iofwrequestrefcon
type IOFWRequestRefCon = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iofirewiresessionref
type IOFireWireSessionRef uintptr

// See: https://developer.apple.com/documentation/kernel/iofixed
type IOFixed = int32

// See: https://developer.apple.com/documentation/kernel/iofixed1616
type IOFixed1616 = uint32

// See: https://developer.apple.com/documentation/kernel/iofixedpoint32
// IOFixedPoint32 is opaque storage with the size and alignment C gives IOFixedPoint32:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOFixedPoint32 [2]uint32

// See: https://developer.apple.com/documentation/kernel/iofourcharcode
type IOFourCharCode = uint32

// See: https://developer.apple.com/documentation/kernel/iogbounds
// IOGBounds is opaque storage with the size and alignment C gives IOGBounds:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOGBounds [4]uint16

// See: https://developer.apple.com/documentation/kernel/iogpoint
// IOGPoint is opaque storage with the size and alignment C gives IOGPoint:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type IOGPoint [2]uint16

// See: https://developer.apple.com/documentation/kernel/iogsize
// IOGSize is opaque storage with the size and alignment C gives IOGSize:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type IOGSize [2]uint16

// See: https://developer.apple.com/documentation/kernel/iohidaccelerationalgorithmtype
type IOHIDAccelerationAlgorithmType = byte

// See: https://developer.apple.com/documentation/kernel/iohidbiometriceventtype
type IOHIDBiometricEventType = uint32

// See: https://developer.apple.com/documentation/kernel/iohidbuttonmodes
type IOHIDButtonModes = int

// IOHIDCompletion is struct specifying action to perform when set/get report completes.
//
// See: https://developer.apple.com/documentation/kernel/iohidcompletion
// IOHIDCompletion is opaque storage with the size and alignment C gives IOHIDCompletion:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IOHIDCompletion [3]uint64

// See: https://developer.apple.com/documentation/kernel/iohiddigitizerstylusdata
// IOHIDDigitizerStylusData is an unresolved C aggregate typedef.
type IOHIDDigitizerStylusData unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iohiddigitizertouchdata
// IOHIDDigitizerTouchData is an unresolved C aggregate typedef.
type IOHIDDigitizerTouchData unsafe.Pointer

// IOHIDElementCollectionType is describes different types of HID collections.
//
// See: https://developer.apple.com/documentation/kernel/iohidelementcollectiontype
type IOHIDElementCollectionType = int

// See: https://developer.apple.com/documentation/kernel/iohidelementcommitdirection
type IOHIDElementCommitDirection = int

// IOHIDElementCookie is abstract data type used as a unique identifier for an element.
//
// See: https://developer.apple.com/documentation/kernel/iohidelementcookie
type IOHIDElementCookie = uint32

// See: https://developer.apple.com/documentation/kernel/iohidelementflags
type IOHIDElementFlags = uint32

// IOHIDElementType is describes different types of HID elements.
//
// See: https://developer.apple.com/documentation/kernel/iohidelementtype
type IOHIDElementType = int

// See: https://developer.apple.com/documentation/kernel/iohideventtype
type IOHIDEventType = uint32

// See: https://developer.apple.com/documentation/kernel/iohidkeyboardphysicallayouttype
type IOHIDKeyboardPhysicalLayoutType = uint32

// See: https://developer.apple.com/documentation/kernel/iohidkind
type IOHIDKind = uint32

// IOHIDOptionsType is options for opening a device via IOHIDLib.
//
// See: https://developer.apple.com/documentation/kernel/iohidoptionstype
type IOHIDOptionsType = uint32

// IOHIDQueueOptionsType is options for creating a queue via IOHIDLib.
//
// See: https://developer.apple.com/documentation/kernel/iohidqueueoptionstype
type IOHIDQueueOptionsType = uint32

// IOHIDReportType is describes different type of HID reports.
//
// See: https://developer.apple.com/documentation/kernel/iohidreporttype
type IOHIDReportType = int

// IOHIDStandardType is type to define what industrial standard the device is referencing.
//
// See: https://developer.apple.com/documentation/kernel/iohidstandardtype
type IOHIDStandardType = uint32

// IOHIDValueOptions is describes options for gathering element values.
//
// See: https://developer.apple.com/documentation/kernel/iohidvalueoptions
type IOHIDValueOptions = uint32

// IOHIDValueScaleType is describes different types of scaling that can be performed on element values.
//
// See: https://developer.apple.com/documentation/kernel/iohidvaluescaletype
type IOHIDValueScaleType = uint32

// See: https://developer.apple.com/documentation/kernel/iohardwarecursordescriptor
// IOHardwareCursorDescriptor is opaque storage with the size and alignment C gives IOHardwareCursorDescriptor:
// 104 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 104 into.
type IOHardwareCursorDescriptor [13]uint64

// See: https://developer.apple.com/documentation/kernel/iohardwarecursorinfo
// IOHardwareCursorInfo is opaque storage with the size and alignment C gives IOHardwareCursorInfo:
// 56 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 56 into.
type IOHardwareCursorInfo [7]uint64

// See: https://developer.apple.com/documentation/kernel/iohistreportinfo
// IOHistReportInfo is opaque storage with the size and alignment C gives IOHistReportInfo:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type IOHistReportInfo [1]uint32

// See: https://developer.apple.com/documentation/kernel/iohistogramreportvalues
// IOHistogramReportValues is opaque storage with the size and alignment C gives IOHistogramReportValues:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type IOHistogramReportValues [32]byte

// See: https://developer.apple.com/documentation/kernel/iohistogramsegmentconfig
// IOHistogramSegmentConfig is opaque storage with the size and alignment C gives IOHistogramSegmentConfig:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type IOHistogramSegmentConfig [16]byte

// See: https://developer.apple.com/documentation/kernel/ioindex
type IOIndex = int32

// See: https://developer.apple.com/documentation/kernel/iointerruptactionblock
type IOInterruptActionBlock = func(nub *IOService, source int32)

// See: https://developer.apple.com/documentation/kernel/iointerruptsource
// IOInterruptSource is an unresolved C aggregate typedef.
type IOInterruptSource unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iointerruptstate
type IOInterruptState = int32

// See: https://developer.apple.com/documentation/kernel/iointerruptvector
// IOInterruptVector is opaque storage with the size and alignment C gives IOInterruptVector:
// 1 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 1 into.
type IOInterruptVector [1]byte

// See: https://developer.apple.com/documentation/kernel/iointerruptvectornumber
type IOInterruptVectorNumber = int32

// See: https://developer.apple.com/documentation/kernel/ioitemcount
type IOItemCount = uint32

// See: https://developer.apple.com/documentation/kernel/iokitdiagnosticsparameters
// IOKitDiagnosticsParameters is opaque storage with the size and alignment C gives IOKitDiagnosticsParameters:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type IOKitDiagnosticsParameters [8]uint64

// See: https://developer.apple.com/documentation/kernel/iolock
// IOLock is an unresolved C aggregate typedef.
type IOLock unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iolockstate
type IOLockState = uint32

// See: https://developer.apple.com/documentation/kernel/iologicaladdress
type IOLogicalAddress = uint64

// See: https://developer.apple.com/documentation/kernel/iomediaattributemask
type IOMediaAttributeMask = uint32

// See: https://developer.apple.com/documentation/kernel/iomediastate
type IOMediaState = uint32

// See: https://developer.apple.com/documentation/kernel/iomessage
type IOMessage = uint32

// See: https://developer.apple.com/documentation/kernel/iomethod
type IOMethod = *objc.ID

// See: https://developer.apple.com/documentation/kernel/ionamedvalue
// IONamedValue is opaque storage with the size and alignment C gives IONamedValue:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type IONamedValue [2]uint64

// See: https://developer.apple.com/documentation/kernel/ionormdistreportvalues
// IONormDistReportValues is opaque storage with the size and alignment C gives IONormDistReportValues:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type IONormDistReportValues [32]byte

// See: https://developer.apple.com/documentation/kernel/ionotificationref
type IONotificationRef = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iooptionbits
type IOOptionBits = uint32

// See: https://developer.apple.com/documentation/kernel/iooutputaction
type IOOutputAction = *uintptr

// See: https://developer.apple.com/documentation/kernel/iopcidevicecrashnotification_t
type IOPCIDeviceCrashNotification_t = uint32

// See: https://developer.apple.com/documentation/kernel/iopmcalendarstruct
// IOPMCalendarStruct is opaque storage with the size and alignment C gives IOPMCalendarStruct:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type IOPMCalendarStruct [3]uint32

// See: https://developer.apple.com/documentation/kernel/iopmdriverassertionid
type IOPMDriverAssertionID = uint64

// See: https://developer.apple.com/documentation/kernel/iopmdriverassertionlevel
type IOPMDriverAssertionLevel = uint32

// See: https://developer.apple.com/documentation/kernel/iopmdriverassertiontype
type IOPMDriverAssertionType = uint64

// IOPMPowerFlags is bits are used in defining capabilityFlags, inputPowerRequirements, and outputPowerCharacter in the IOPMPowerState structure.
//
// See: https://developer.apple.com/documentation/kernel/iopmpowerflags
type IOPMPowerFlags = uint

// See: https://developer.apple.com/documentation/kernel/iopmpowerstate
// IOPMPowerState is opaque storage with the size and alignment C gives IOPMPowerState:
// 96 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 96 into.
type IOPMPowerState [12]uint64

// See: https://developer.apple.com/documentation/kernel/iopacketbufferconstraints
// IOPacketBufferConstraints is an unresolved C aggregate typedef.
type IOPacketBufferConstraints unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iophysicaladdress
type IOPhysicalAddress = uint64

// See: https://developer.apple.com/documentation/kernel/iophysicaladdress32
type IOPhysicalAddress32 = uint32

// See: https://developer.apple.com/documentation/kernel/iophysicaladdress64
type IOPhysicalAddress64 = uint64

// See: https://developer.apple.com/documentation/kernel/iophysicallength
type IOPhysicalLength = uint64

// See: https://developer.apple.com/documentation/kernel/iophysicallength32
type IOPhysicalLength32 = uint32

// See: https://developer.apple.com/documentation/kernel/iophysicallength64
type IOPhysicalLength64 = uint64

// See: https://developer.apple.com/documentation/kernel/iophysicalrange
// IOPhysicalRange is opaque storage with the size and alignment C gives IOPhysicalRange:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type IOPhysicalRange [2]uint64

// See: https://developer.apple.com/documentation/kernel/iopixelaperture
type IOPixelAperture = int32

// See: https://developer.apple.com/documentation/kernel/iopixelencoding
type IOPixelEncoding = int8

// See: https://developer.apple.com/documentation/kernel/iopixelinformation
// IOPixelInformation is opaque storage with the size and alignment C gives IOPixelInformation:
// 172 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 172 into.
type IOPixelInformation [43]uint32

// See: https://developer.apple.com/documentation/kernel/iopowerstatechangenotification
// IOPowerStateChangeNotification is opaque storage with the size and alignment C gives IOPowerStateChangeNotification:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type IOPowerStateChangeNotification [4]uint64

// See: https://developer.apple.com/documentation/kernel/iopropertyname
type IOPropertyName = int8

// See: https://developer.apple.com/documentation/kernel/iorpc
// IORPC is opaque storage with the size and alignment C gives IORPC:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IORPC [3]uint64

// See: https://developer.apple.com/documentation/kernel/iorpcmessage
// IORPCMessage is opaque storage with the size and alignment C gives IORPCMessage:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IORPCMessage [6]uint32

// See: https://developer.apple.com/documentation/kernel/iorpcmessageerrorreturncontent
// IORPCMessageErrorReturnContent is opaque storage with the size and alignment C gives IORPCMessageErrorReturnContent:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type IORPCMessageErrorReturnContent [8]uint32

// See: https://developer.apple.com/documentation/kernel/iorpcmessagemach
// IORPCMessageMach is opaque storage with the size and alignment C gives IORPCMessageMach:
// 28 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 28 into.
type IORPCMessageMach [7]uint32

// See: https://developer.apple.com/documentation/kernel/iorwlock
// IORWLock is an unresolved C aggregate typedef.
type IORWLock unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iorangescalar
type IORangeScalar = uint64

// See: https://developer.apple.com/documentation/kernel/iorecursivelock
// IORecursiveLock is an unresolved C aggregate typedef.
type IORecursiveLock unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioregistryplanename
type IORegistryPlaneName = int8

// See: https://developer.apple.com/documentation/kernel/ioreportcategories
type IOReportCategories = uint16

// See: https://developer.apple.com/documentation/kernel/ioreportchannel
// IOReportChannel is opaque storage with the size and alignment C gives IOReportChannel:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type IOReportChannel [2]uint64

// See: https://developer.apple.com/documentation/kernel/ioreportchannellist
// IOReportChannelList is opaque storage with the size and alignment C gives IOReportChannelList:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOReportChannelList [1]uint64

// See: https://developer.apple.com/documentation/kernel/ioreportchanneltype
// IOReportChannelType is opaque storage with the size and alignment C gives IOReportChannelType:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOReportChannelType [8]byte

// See: https://developer.apple.com/documentation/kernel/ioreportconfigureaction
type IOReportConfigureAction = uint32

// See: https://developer.apple.com/documentation/kernel/ioreportelement
// IOReportElement is opaque storage with the size and alignment C gives IOReportElement:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type IOReportElement [64]byte

// See: https://developer.apple.com/documentation/kernel/ioreportelementvalues
// IOReportElementValues is opaque storage with the size and alignment C gives IOReportElementValues:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type IOReportElementValues [32]byte

// See: https://developer.apple.com/documentation/kernel/ioreportformat
type IOReportFormat = byte

// See: https://developer.apple.com/documentation/kernel/ioreportinterest
// IOReportInterest is opaque storage with the size and alignment C gives IOReportInterest:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IOReportInterest [3]uint64

// See: https://developer.apple.com/documentation/kernel/ioreportinterestlist
// IOReportInterestList is opaque storage with the size and alignment C gives IOReportInterestList:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOReportInterestList [1]uint64

// See: https://developer.apple.com/documentation/kernel/ioreportquantity
type IOReportQuantity = byte

// See: https://developer.apple.com/documentation/kernel/ioreportscalefactor
type IOReportScaleFactor = uint64

// See: https://developer.apple.com/documentation/kernel/ioreportunit
type IOReportUnit = uint64

// See: https://developer.apple.com/documentation/kernel/ioreportunits
type IOReportUnits = uint64

// See: https://developer.apple.com/documentation/kernel/ioreportupdateaction
type IOReportUpdateAction = uint32

// See: https://developer.apple.com/documentation/kernel/ioreturn
type IOReturn = int32

// See: https://developer.apple.com/documentation/kernel/ioselect
type IOSelect = uint32

// See: https://developer.apple.com/documentation/kernel/ioserviceapplierblock
type IOServiceApplierBlock = func(service *IOService)

// See: https://developer.apple.com/documentation/kernel/ioserviceinteresthandlerblock
type IOServiceInterestHandlerBlock = func(messageType uint32, provider *IOService, messageArgument unsafe.Pointer, argSize uintptr) int32

// See: https://developer.apple.com/documentation/kernel/ioservicematchingnotificationhandlerblock
type IOServiceMatchingNotificationHandlerBlock = func(newService *IOService, notifier *IONotifier) bool

// See: https://developer.apple.com/documentation/kernel/ioservicename
type IOServiceName = int8

// See: https://developer.apple.com/documentation/kernel/ioservicenotificationblock
type IOServiceNotificationBlock = func(type_ uint64, service *IOService, options uint64)

// See: https://developer.apple.com/documentation/kernel/iosimplearrayreportvalues
// IOSimpleArrayReportValues is opaque storage with the size and alignment C gives IOSimpleArrayReportValues:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type IOSimpleArrayReportValues [32]byte

// See: https://developer.apple.com/documentation/kernel/iosimplelock
// IOSimpleLock is an unresolved C aggregate typedef.
type IOSimpleLock unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iosimplereportvalues
// IOSimpleReportValues is opaque storage with the size and alignment C gives IOSimpleReportValues:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type IOSimpleReportValues [32]byte

// See: https://developer.apple.com/documentation/kernel/iostatenotificationhandler
type IOStateNotificationHandler = func() Kern_return_t

// See: https://developer.apple.com/documentation/kernel/iostatenotificationlistenerref
type IOStateNotificationListenerRef = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iostatereportinfo
// IOStateReportInfo is opaque storage with the size and alignment C gives IOStateReportInfo:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type IOStateReportInfo [2]uint64

// See: https://developer.apple.com/documentation/kernel/iostatereportvalues
// IOStateReportValues is opaque storage with the size and alignment C gives IOStateReportValues:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type IOStateReportValues [32]byte

// See: https://developer.apple.com/documentation/kernel/iostorageaccess
type IOStorageAccess = uint32

// See: https://developer.apple.com/documentation/kernel/iostoragegetprovisionstatusoptions
type IOStorageGetProvisionStatusOptions = uint64

// See: https://developer.apple.com/documentation/kernel/iostorageoptions
type IOStorageOptions = uint16

// See: https://developer.apple.com/documentation/kernel/iostoragepriority
type IOStoragePriority = byte

// See: https://developer.apple.com/documentation/kernel/iostoragesynchronizeoptions
type IOStorageSynchronizeOptions = uint32

// See: https://developer.apple.com/documentation/kernel/iostorageunmapoptions
type IOStorageUnmapOptions = uint32

// See: https://developer.apple.com/documentation/kernel/iostreammode
type IOStreamMode = uint

// See: https://developer.apple.com/documentation/kernel/iotvector
// IOTVector is opaque storage with the size and alignment C gives IOTVector:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type IOTVector [2]uint64

// See: https://developer.apple.com/documentation/kernel/iothread
type IOThread = Thread_t

// See: https://developer.apple.com/documentation/kernel/iotiminginformation
// IOTimingInformation is opaque storage with the size and alignment C gives IOTimingInformation:
// 168 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 168 into.
type IOTimingInformation [21]uint64

// See: https://developer.apple.com/documentation/kernel/iotrap
type IOTrap = *objc.ID

// IOUSB20HubDescriptor is a structure that defines the descriptor for a USB 2.0 hub.
//
// See: https://developer.apple.com/documentation/kernel/iousb20hubdescriptor
// IOUSB20HubDescriptor is opaque storage with the size and alignment C gives IOUSB20HubDescriptor:
// 11 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 11 into.
type IOUSB20HubDescriptor [11]byte

// IOUSB3HubDescriptor is a structure that defines the descriptor for a USB 3.0 hub.
//
// See: https://developer.apple.com/documentation/kernel/iousb3hubdescriptor
// IOUSB3HubDescriptor is an unresolved C aggregate typedef.
type IOUSB3HubDescriptor unsafe.Pointer

// IOUSBBOSDescriptor is the structure for storing a binary object store (BOS) descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbbosdescriptor
// IOUSBBOSDescriptor is opaque storage with the size and alignment C gives IOUSBBOSDescriptor:
// 5 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 5 into.
type IOUSBBOSDescriptor [5]byte

// IOUSBBOSDescriptorPtr is a pointer to a structure for storing a binary object store (BOS) descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbbosdescriptorptr
type IOUSBBOSDescriptorPtr = *IOUSBBOSDescriptor

// IOUSBBulkPipeReq is the structure that represents a bulk pipe request.
//
// See: https://developer.apple.com/documentation/kernel/iousbbulkpipereq
// IOUSBBulkPipeReq is opaque storage with the size and alignment C gives IOUSBBulkPipeReq:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type IOUSBBulkPipeReq [4]uint64

// IOUSBCompletion is the structure that specifies the action to perform when the USB input/output request completes.
//
// See: https://developer.apple.com/documentation/kernel/iousbcompletion
// IOUSBCompletion is opaque storage with the size and alignment C gives IOUSBCompletion:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IOUSBCompletion [3]uint64

// IOUSBCompletionWithTimeStamp is a structure specifying action to perform when the USB input/output request completes.
//
// See: https://developer.apple.com/documentation/kernel/iousbcompletionwithtimestamp
// IOUSBCompletionWithTimeStamp is opaque storage with the size and alignment C gives IOUSBCompletionWithTimeStamp:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IOUSBCompletionWithTimeStamp [3]uint64

// IOUSBConfigurationDescHeader is the header of a configuration descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbconfigurationdescheader
// IOUSBConfigurationDescHeader is opaque storage with the size and alignment C gives IOUSBConfigurationDescHeader:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type IOUSBConfigurationDescHeader [4]byte

// IOUSBConfigurationDescHeaderPtr is a pointer to a configuration descriptor header.
//
// See: https://developer.apple.com/documentation/kernel/iousbconfigurationdescheaderptr
type IOUSBConfigurationDescHeaderPtr = *IOUSBConfigurationDescHeader

// IOUSBConfigurationDescriptor is the structure for storing a USB configuration descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbconfigurationdescriptor
// IOUSBConfigurationDescriptor is opaque storage with the size and alignment C gives IOUSBConfigurationDescriptor:
// 9 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 9 into.
type IOUSBConfigurationDescriptor [9]byte

// IOUSBConfigurationDescriptorPtr is a pointer to a configuration descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbconfigurationdescriptorptr
type IOUSBConfigurationDescriptorPtr = *IOUSBConfigurationDescriptor

// IOUSBDFUDescriptor is a structure that defines the USB device firmware update descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbdfudescriptor
// IOUSBDFUDescriptor is opaque storage with the size and alignment C gives IOUSBDFUDescriptor:
// 7 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 7 into.
type IOUSBDFUDescriptor [7]byte

// IOUSBDFUDescriptorPtr is a pointer to a structure that defines the USB device firmware update descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbdfudescriptorptr
type IOUSBDFUDescriptorPtr = *IOUSBDFUDescriptor

// IOUSBDescriptor is the base descriptor type.
//
// See: https://developer.apple.com/documentation/kernel/iousbdescriptor
// IOUSBDescriptor is opaque storage with the size and alignment C gives IOUSBDescriptor:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type IOUSBDescriptor [2]byte

// IOUSBDescriptorHeader is the base descriptor header.
//
// See: https://developer.apple.com/documentation/kernel/iousbdescriptorheader
// IOUSBDescriptorHeader is opaque storage with the size and alignment C gives IOUSBDescriptorHeader:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type IOUSBDescriptorHeader [2]byte

// IOUSBDescriptorHeaderPtr is a pointer to a USB descriptor header.
//
// See: https://developer.apple.com/documentation/kernel/iousbdescriptorheaderptr
type IOUSBDescriptorHeaderPtr = *IOUSBDescriptorHeader

// IOUSBDevReqOOL is an internal structure to pass parameters between IOUSBLib and UserClient.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevreqool
// IOUSBDevReqOOL is opaque storage with the size and alignment C gives IOUSBDevReqOOL:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IOUSBDevReqOOL [3]uint64

// IOUSBDevReqOOLTO is an internal structure to pass parameters between IOUSBLib and UserClient.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevreqoolto
// IOUSBDevReqOOLTO is opaque storage with the size and alignment C gives IOUSBDevReqOOLTO:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type IOUSBDevReqOOLTO [4]uint64

// IOUSBDevRequest is a structure that defines a standard device request.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevrequest
// IOUSBDevRequest is opaque storage with the size and alignment C gives IOUSBDevRequest:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IOUSBDevRequest [3]uint64

// IOUSBDevRequestTO is a structure that defines a standard device request with timeout.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevrequestto
// IOUSBDevRequestTO is opaque storage with the size and alignment C gives IOUSBDevRequestTO:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type IOUSBDevRequestTO [4]uint64

// IOUSBDeviceCapabilityBillboard is the structure for the billboard device capability.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitybillboard
// IOUSBDeviceCapabilityBillboard is opaque storage with the size and alignment C gives IOUSBDeviceCapabilityBillboard:
// 44 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 44 into.
type IOUSBDeviceCapabilityBillboard [44]byte

// IOUSBDeviceCapabilityBillboardAltConfig is the structure for the billboard alternative configuration device capability.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitybillboardaltconfig
// IOUSBDeviceCapabilityBillboardAltConfig is opaque storage with the size and alignment C gives IOUSBDeviceCapabilityBillboardAltConfig:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type IOUSBDeviceCapabilityBillboardAltConfig [4]byte

// IOUSBDeviceCapabilityBillboardAltConfigCompatibility is the structure for the billboard alternative configuration compatibility device capability.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitybillboardaltconfigcompatibility
// IOUSBDeviceCapabilityBillboardAltConfigCompatibility is opaque storage with the size and alignment C gives IOUSBDeviceCapabilityBillboardAltConfigCompatibility:
// 7 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 7 into.
type IOUSBDeviceCapabilityBillboardAltConfigCompatibility [7]byte

// IOUSBDeviceCapabilityBillboardAltConfigPtr is a pointer to a USB device capability billboard alternative configuration structure.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitybillboardaltconfigptr
type IOUSBDeviceCapabilityBillboardAltConfigPtr = *IOUSBDeviceCapabilityBillboardAltConfig

// IOUSBDeviceCapabilityBillboardAltMode is the structure for the billboard alternative mode device capability.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitybillboardaltmode
type IOUSBDeviceCapabilityBillboardAltMode = uint

// IOUSBDeviceCapabilityBillboardAltModePtr is a pointer to a USB device capability billboard alternative mode structure.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitybillboardaltmodeptr
type IOUSBDeviceCapabilityBillboardAltModePtr = *IOUSBDeviceCapabilityBillboardAltMode

// IOUSBDeviceCapabilityBillboardPtr is a pointer to a USB device capability billboard object.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitybillboardptr
type IOUSBDeviceCapabilityBillboardPtr = *IOUSBDeviceCapabilityBillboard

// IOUSBDeviceCapabilityContainerID is the structure for the container ID device capability.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitycontainerid
// IOUSBDeviceCapabilityContainerID is opaque storage with the size and alignment C gives IOUSBDeviceCapabilityContainerID:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type IOUSBDeviceCapabilityContainerID [20]byte

// IOUSBDeviceCapabilityContainerIDPtr is a pointer to a USB device capability container ID.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitycontaineridptr
type IOUSBDeviceCapabilityContainerIDPtr = unsafe.Pointer

// IOUSBDeviceCapabilityDescriptorHeader is the device capability descriptor header.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitydescriptorheader
// IOUSBDeviceCapabilityDescriptorHeader is opaque storage with the size and alignment C gives IOUSBDeviceCapabilityDescriptorHeader:
// 3 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 3 into.
type IOUSBDeviceCapabilityDescriptorHeader [3]byte

// IOUSBDeviceCapabilityDescriptorHeaderPtr is a pointer to a device capability descriptor header.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitydescriptorheaderptr
type IOUSBDeviceCapabilityDescriptorHeaderPtr = *IOUSBDeviceCapabilityDescriptorHeader

// IOUSBDeviceCapabilitySuperSpeedPlusUSB is the structure for the SuperSpeedPlus USB device capability.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitysuperspeedplususb
// IOUSBDeviceCapabilitySuperSpeedPlusUSB is opaque storage with the size and alignment C gives IOUSBDeviceCapabilitySuperSpeedPlusUSB:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type IOUSBDeviceCapabilitySuperSpeedPlusUSB [12]byte

// IOUSBDeviceCapabilitySuperSpeedPlusUSBPtr is a pointer to a SuperSpeedPlus USB device capability structure.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitysuperspeedplususbptr
type IOUSBDeviceCapabilitySuperSpeedPlusUSBPtr = *IOUSBDeviceCapabilitySuperSpeedPlusUSB

// IOUSBDeviceCapabilitySuperSpeedUSB is the structure for the SuperSpeed USB device capability.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitysuperspeedusb
// IOUSBDeviceCapabilitySuperSpeedUSB is opaque storage with the size and alignment C gives IOUSBDeviceCapabilitySuperSpeedUSB:
// 10 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 10 into.
type IOUSBDeviceCapabilitySuperSpeedUSB [10]byte

// IOUSBDeviceCapabilitySuperSpeedUSBPtr is a pointer to a SuperSpeed USB device capability structure.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitysuperspeedusbptr
type IOUSBDeviceCapabilitySuperSpeedUSBPtr = *IOUSBDeviceCapabilitySuperSpeedUSB

// IOUSBDeviceCapabilityUSB2Extension is the structure for the USB 2.0 extension device capability.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilityusb2extension
// IOUSBDeviceCapabilityUSB2Extension is opaque storage with the size and alignment C gives IOUSBDeviceCapabilityUSB2Extension:
// 7 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 7 into.
type IOUSBDeviceCapabilityUSB2Extension [7]byte

// IOUSBDeviceCapabilityUSB2ExtensionPtr is a pointer to a USB 2.0 extension device capability structure.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilityusb2extensionptr
type IOUSBDeviceCapabilityUSB2ExtensionPtr = *IOUSBDeviceCapabilityUSB2Extension

// IOUSBDeviceDescriptor is the structure for storing a USB device descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicedescriptor
// IOUSBDeviceDescriptor is opaque storage with the size and alignment C gives IOUSBDeviceDescriptor:
// 18 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 18 into.
type IOUSBDeviceDescriptor [18]byte

// IOUSBDeviceDescriptorPtr is a pointer to a USB device descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicedescriptorptr
type IOUSBDeviceDescriptorPtr = *IOUSBDeviceDescriptor

// IOUSBDeviceQualifierDescriptor is the structure for describing a high-speed capable USB device.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicequalifierdescriptor
// IOUSBDeviceQualifierDescriptor is opaque storage with the size and alignment C gives IOUSBDeviceQualifierDescriptor:
// 10 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 10 into.
type IOUSBDeviceQualifierDescriptor [10]byte

// IOUSBDeviceQualifierDescriptorPtr is a pointer to a qualifier descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicequalifierdescriptorptr
type IOUSBDeviceQualifierDescriptorPtr = *IOUSBDeviceQualifierDescriptor

// IOUSBDeviceRequest is a structure that defines a standard device request.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicerequest
// IOUSBDeviceRequest is opaque storage with the size and alignment C gives IOUSBDeviceRequest:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOUSBDeviceRequest [8]byte

// IOUSBDeviceRequestPtr is a pointer to a structure that defines a standard device request.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicerequestptr
type IOUSBDeviceRequestPtr = *IOUSBDevRequest

// IOUSBDeviceRequestSetSELData is the structure for receiving system exit latency values.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicerequestsetseldata
// IOUSBDeviceRequestSetSELData is opaque storage with the size and alignment C gives IOUSBDeviceRequestSetSELData:
// 6 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 6 into.
type IOUSBDeviceRequestSetSELData [6]byte

// IOUSBEndpointDescriptor is the structure for storing an endpoint descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbendpointdescriptor
// IOUSBEndpointDescriptor is opaque storage with the size and alignment C gives IOUSBEndpointDescriptor:
// 7 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 7 into.
type IOUSBEndpointDescriptor [7]byte

// IOUSBEndpointDescriptorPtr is a pointer to the endpoint descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbendpointdescriptorptr
type IOUSBEndpointDescriptorPtr = *IOUSBEndpointDescriptor

// IOUSBEndpointProperties is a structure that holds USB endpoint properties.
//
// See: https://developer.apple.com/documentation/kernel/iousbendpointproperties
// IOUSBEndpointProperties is opaque storage with the size and alignment C gives IOUSBEndpointProperties:
// 15 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 15 into.
type IOUSBEndpointProperties [15]byte

// IOUSBEndpointPropertiesPtr is a pointer to an endpoint properties object.
//
// See: https://developer.apple.com/documentation/kernel/iousbendpointpropertiesptr
type IOUSBEndpointPropertiesPtr = *IOUSBEndpointProperties

// IOUSBFindEndpointRequest is the structure that represents an endoint request to locate.
//
// See: https://developer.apple.com/documentation/kernel/iousbfindendpointrequest
// IOUSBFindEndpointRequest is opaque storage with the size and alignment C gives IOUSBFindEndpointRequest:
// 6 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 6 into.
type IOUSBFindEndpointRequest [3]uint16

// IOUSBFindInterfaceRequest is the structure for finding an interface request.
//
// See: https://developer.apple.com/documentation/kernel/iousbfindinterfacerequest
// IOUSBFindInterfaceRequest is opaque storage with the size and alignment C gives IOUSBFindInterfaceRequest:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOUSBFindInterfaceRequest [4]uint16

// IOUSBGetFrameStruct is a structure that contains frame information.
//
// See: https://developer.apple.com/documentation/kernel/iousbgetframestruct
// IOUSBGetFrameStruct is opaque storage with the size and alignment C gives IOUSBGetFrameStruct:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type IOUSBGetFrameStruct [2]uint64

// IOUSBHIDData is data related to the mouse and keyboard.
//
// See: https://developer.apple.com/documentation/kernel/iousbhiddata
// IOUSBHIDData is opaque storage with the size and alignment C gives IOUSBHIDData:
// 66 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 66 into.
type IOUSBHIDData [33]uint16

// IOUSBHIDDataPtr is a pointer to a structure related to mouse and keyboard data.
//
// See: https://developer.apple.com/documentation/kernel/iousbhiddataptr
type IOUSBHIDDataPtr = *IOUSBHIDData

// IOUSBHIDDescriptor is a structure that defines the USB HID descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbhiddescriptor
// IOUSBHIDDescriptor is opaque storage with the size and alignment C gives IOUSBHIDDescriptor:
// 9 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 9 into.
type IOUSBHIDDescriptor [9]byte

// IOUSBHIDDescriptorPtr is a pointer to a structure that defines the USB HID descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbhiddescriptorptr
type IOUSBHIDDescriptorPtr = *IOUSBHIDDescriptor

// IOUSBHIDReportDesc is a structure that defines the USB HID report descriptor header.
//
// See: https://developer.apple.com/documentation/kernel/iousbhidreportdesc
// IOUSBHIDReportDesc is opaque storage with the size and alignment C gives IOUSBHIDReportDesc:
// 3 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 3 into.
type IOUSBHIDReportDesc [3]byte

// IOUSBHIDReportDescPtr is a pointer to a structure that defines the USB HID report descriptor header.
//
// See: https://developer.apple.com/documentation/kernel/iousbhidreportdescptr
type IOUSBHIDReportDescPtr = *IOUSBHIDReportDesc

// See: https://developer.apple.com/documentation/kernel/iousbhostcimessage
// IOUSBHostCIMessage is opaque storage with the size and alignment C gives IOUSBHostCIMessage:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type IOUSBHostCIMessage [16]byte

// See: https://developer.apple.com/documentation/kernel/iousbhostciuserclientversion
type IOUSBHostCIUserClientVersion = int

// IOUSBHostIOSourceClientRecordLink is a structure that represents a USB host input/output source client record entry.
//
// See: https://developer.apple.com/documentation/kernel/iousbhostiosourceclientrecordlink
// IOUSBHostIOSourceClientRecordLink is opaque storage with the size and alignment C gives IOUSBHostIOSourceClientRecordLink:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type IOUSBHostIOSourceClientRecordLink [2]uint64

// IOUSBHostIOSourceClientRecordList is a structure that represents a list of USB host input/output source client records.
//
// See: https://developer.apple.com/documentation/kernel/iousbhostiosourceclientrecordlist
// IOUSBHostIOSourceClientRecordList is opaque storage with the size and alignment C gives IOUSBHostIOSourceClientRecordList:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOUSBHostIOSourceClientRecordList [1]uint64

// IOUSBHubDescriptor is a structure that defines the descriptor for a USB hub.
//
// See: https://developer.apple.com/documentation/kernel/iousbhubdescriptor
// IOUSBHubDescriptor is an unresolved C aggregate typedef.
type IOUSBHubDescriptor unsafe.Pointer

// IOUSBHubPortReEnumerateParam is a structure for USB hub port reenumeration.
//
// See: https://developer.apple.com/documentation/kernel/iousbhubportreenumerateparam
// IOUSBHubPortReEnumerateParam is an unresolved C aggregate typedef.
type IOUSBHubPortReEnumerateParam unsafe.Pointer

// IOUSBHubPortStatus is a structure that contains the USB hub port status.
//
// See: https://developer.apple.com/documentation/kernel/iousbhubportstatus
type IOUSBHubPortStatus = IOUSBHubStatus

// IOUSBHubStatus is a structure that represents the USB hub status.
//
// See: https://developer.apple.com/documentation/kernel/iousbhubstatus
// IOUSBHubStatus is an unresolved C aggregate typedef.
type IOUSBHubStatus unsafe.Pointer

// IOUSBHubStatusPtr is a pointer to a USB hub status structure.
//
// See: https://developer.apple.com/documentation/kernel/iousbhubstatusptr
type IOUSBHubStatusPtr = unsafe.Pointer

// IOUSBInterfaceAssociationDescriptor is the descriptor that associates multiple interfaces to the same function.
//
// See: https://developer.apple.com/documentation/kernel/iousbinterfaceassociationdescriptor
// IOUSBInterfaceAssociationDescriptor is opaque storage with the size and alignment C gives IOUSBInterfaceAssociationDescriptor:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOUSBInterfaceAssociationDescriptor [8]byte

// IOUSBInterfaceAssociationDescriptorPtr is a pointer to a USB interface association descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbinterfaceassociationdescriptorptr
type IOUSBInterfaceAssociationDescriptorPtr = *IOUSBInterfaceAssociationDescriptor

// IOUSBInterfaceDescriptor is a descriptor for a specific interface of a USB device.
//
// See: https://developer.apple.com/documentation/kernel/iousbinterfacedescriptor
// IOUSBInterfaceDescriptor is opaque storage with the size and alignment C gives IOUSBInterfaceDescriptor:
// 9 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 9 into.
type IOUSBInterfaceDescriptor [9]byte

// IOUSBInterfaceDescriptorPtr is a pointer to a USB interface descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbinterfacedescriptorptr
type IOUSBInterfaceDescriptorPtr = *IOUSBInterfaceDescriptor

// IOUSBIsocCompletion is a structure specifying the action to perform when an isochronous USB input/output operation completes.
//
// See: https://developer.apple.com/documentation/kernel/iousbisoccompletion
// IOUSBIsocCompletion is opaque storage with the size and alignment C gives IOUSBIsocCompletion:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IOUSBIsocCompletion [3]uint64

// IOUSBIsocFrame is a structure for encoding information about a single frame in an isochronous transfer.
//
// See: https://developer.apple.com/documentation/kernel/iousbisocframe
// IOUSBIsocFrame is opaque storage with the size and alignment C gives IOUSBIsocFrame:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOUSBIsocFrame [2]uint32

// IOUSBIsocStruct is an internal structure to pass parameters between IOUSBLib and UserClient.
//
// See: https://developer.apple.com/documentation/kernel/iousbisocstruct
// IOUSBIsocStruct is opaque storage with the size and alignment C gives IOUSBIsocStruct:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type IOUSBIsocStruct [6]uint64

// IOUSBIsochronousFrame is a structure representing a single frame in an isochronous transfer.
//
// See: https://developer.apple.com/documentation/kernel/iousbisochronousframe
// IOUSBIsochronousFrame is opaque storage with the size and alignment C gives IOUSBIsochronousFrame:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IOUSBIsochronousFrame [24]byte

// IOUSBKeyboardData is a structure containing USB keyboard data.
//
// See: https://developer.apple.com/documentation/kernel/iousbkeyboarddata
// IOUSBKeyboardData is opaque storage with the size and alignment C gives IOUSBKeyboardData:
// 66 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 66 into.
type IOUSBKeyboardData [33]uint16

// IOUSBKeyboardDataPtr is a pointer to a structure containing USB keyboard data.
//
// See: https://developer.apple.com/documentation/kernel/iousbkeyboarddataptr
type IOUSBKeyboardDataPtr = *IOUSBKeyboardData

// IOUSBLowLatencyIsocCompletion is the function that executes when the low-latency isochronous USB input/output request completes.
//
// See: https://developer.apple.com/documentation/kernel/iousblowlatencyisoccompletion
// IOUSBLowLatencyIsocCompletion is opaque storage with the size and alignment C gives IOUSBLowLatencyIsocCompletion:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IOUSBLowLatencyIsocCompletion [3]uint64

// IOUSBLowLatencyIsocFrame is a structure for encoding information about each low-latency isochronous frame.
//
// See: https://developer.apple.com/documentation/kernel/iousblowlatencyisocframe
// IOUSBLowLatencyIsocFrame is opaque storage with the size and alignment C gives IOUSBLowLatencyIsocFrame:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type IOUSBLowLatencyIsocFrame [4]uint32

// IOUSBLowLatencyIsocStruct is an internal structure to pass parameters between IOUSBLib and UserClient.
//
// See: https://developer.apple.com/documentation/kernel/iousblowlatencyisocstruct
// IOUSBLowLatencyIsocStruct is opaque storage with the size and alignment C gives IOUSBLowLatencyIsocStruct:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type IOUSBLowLatencyIsocStruct [5]uint64

// IOUSBMatch is a structure for matching USB devices.
//
// See: https://developer.apple.com/documentation/kernel/iousbmatch
// IOUSBMatch is opaque storage with the size and alignment C gives IOUSBMatch:
// 10 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 10 into.
type IOUSBMatch [5]uint16

// IOUSBMouseData is a structure containing USB mouse data.
//
// See: https://developer.apple.com/documentation/kernel/iousbmousedata
// IOUSBMouseData is opaque storage with the size and alignment C gives IOUSBMouseData:
// 6 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 6 into.
type IOUSBMouseData [3]uint16

// IOUSBMouseDataPtr is a pointer to a structure containing USB mouse data.
//
// See: https://developer.apple.com/documentation/kernel/iousbmousedataptr
type IOUSBMouseDataPtr = *IOUSBMouseData

// IOUSBPlatformCapabilityDescriptor is the structure for the platform capability descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbplatformcapabilitydescriptor
// IOUSBPlatformCapabilityDescriptor is opaque storage with the size and alignment C gives IOUSBPlatformCapabilityDescriptor:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type IOUSBPlatformCapabilityDescriptor [20]byte

// IOUSBPlatformCapabilityDescriptorPtr is a pointer to a USB platform capability descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbplatformcapabilitydescriptorptr
type IOUSBPlatformCapabilityDescriptorPtr = *IOUSBPlatformCapabilityDescriptor

// IOUSBStringDescriptor is the structure for storing a string descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbstringdescriptor
// IOUSBStringDescriptor is opaque storage with the size and alignment C gives IOUSBStringDescriptor:
// 3 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 3 into.
type IOUSBStringDescriptor [3]byte

// IOUSBStringDescriptorPtr is a pointer to a string descriptor structure.
//
// See: https://developer.apple.com/documentation/kernel/iousbstringdescriptorptr
type IOUSBStringDescriptorPtr = *IOUSBStringDescriptor

// IOUSBSuperSpeedEndpointCompanionDescriptor is the descriptor for a SuperSpeed USB endpoint companion.
//
// See: https://developer.apple.com/documentation/kernel/iousbsuperspeedendpointcompaniondescriptor
// IOUSBSuperSpeedEndpointCompanionDescriptor is opaque storage with the size and alignment C gives IOUSBSuperSpeedEndpointCompanionDescriptor:
// 6 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 6 into.
type IOUSBSuperSpeedEndpointCompanionDescriptor [6]byte

// IOUSBSuperSpeedEndpointCompanionDescriptorPtr is a pointer to a SuperSpeed USB endpoint companion descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbsuperspeedendpointcompaniondescriptorptr
type IOUSBSuperSpeedEndpointCompanionDescriptorPtr = *IOUSBSuperSpeedEndpointCompanionDescriptor

// IOUSBSuperSpeedHubDescriptor is a structure that defines the descriptor for a SuperSpeed USB hub.
//
// See: https://developer.apple.com/documentation/kernel/iousbsuperspeedhubdescriptor
// IOUSBSuperSpeedHubDescriptor is opaque storage with the size and alignment C gives IOUSBSuperSpeedHubDescriptor:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type IOUSBSuperSpeedHubDescriptor [12]byte

// IOUSBSuperSpeedPlusIsochronousEndpointCompanionDescriptor is the descriptor for a SuperSpeedPlus isochronous USB endpoint companion.
//
// See: https://developer.apple.com/documentation/kernel/iousbsuperspeedplusisochronousendpointcompaniondescriptor
// IOUSBSuperSpeedPlusIsochronousEndpointCompanionDescriptor is opaque storage with the size and alignment C gives IOUSBSuperSpeedPlusIsochronousEndpointCompanionDescriptor:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOUSBSuperSpeedPlusIsochronousEndpointCompanionDescriptor [8]byte

// IOUSBSuperSpeedPlusIsochronousEndpointCompanionDescriptorPtr is a pointer to a SuperSpeedPlus isochronous USB endpoint companion descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbsuperspeedplusisochronousendpointcompaniondescriptorptr
type IOUSBSuperSpeedPlusIsochronousEndpointCompanionDescriptorPtr = *IOUSBSuperSpeedPlusIsochronousEndpointCompanionDescriptor

// See: https://developer.apple.com/documentation/kernel/iouserclientasyncargumentsarray
type IOUserClientAsyncArgumentsArray = uint64

// See: https://developer.apple.com/documentation/kernel/iouserclientasyncreferencearray
type IOUserClientAsyncReferenceArray = uint64

// See: https://developer.apple.com/documentation/kernel/iouserclientscalararray
type IOUserClientScalarArray = uint64

// See: https://developer.apple.com/documentation/kernel/ioversion
type IOVersion = uint32

// See: https://developer.apple.com/documentation/kernel/iovideodevicenotification
// IOVideoDeviceNotification is opaque storage with the size and alignment C gives IOVideoDeviceNotification:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type IOVideoDeviceNotification [4]uint64

// See: https://developer.apple.com/documentation/kernel/iovideodevicenotificationmessage
// IOVideoDeviceNotificationMessage is opaque storage with the size and alignment C gives IOVideoDeviceNotificationMessage:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type IOVideoDeviceNotificationMessage [8]uint64

// See: https://developer.apple.com/documentation/kernel/iovideostreamdescription
// IOVideoStreamDescription is opaque storage with the size and alignment C gives IOVideoStreamDescription:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IOVideoStreamDescription [6]uint32

// See: https://developer.apple.com/documentation/kernel/iovirtualaddress
type IOVirtualAddress = uint64

// See: https://developer.apple.com/documentation/kernel/iovirtualrange
// IOVirtualRange is opaque storage with the size and alignment C gives IOVirtualRange:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type IOVirtualRange [2]uint64

// See: https://developer.apple.com/documentation/kernel/interruptserviceidptr
type InterruptServiceIDPtr = *InterruptServiceIDType

// See: https://developer.apple.com/documentation/kernel/interruptserviceidtype
type InterruptServiceIDType = uintptr

// See: https://developer.apple.com/documentation/kernel/interruptservicetype
type InterruptServiceType = uint32

// ItemCount is abst_ItemCount.
//
// See: https://developer.apple.com/documentation/kernel/itemcount
type ItemCount = uint

// See: https://developer.apple.com/documentation/kernel/journalinfoblock
// JournalInfoBlock is opaque storage with the size and alignment C gives JournalInfoBlock:
// 180 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 180 into.
type JournalInfoBlock [90]uint16

// See: https://developer.apple.com/documentation/kernel/kuncusernotificationid
type KUNCUserNotificationID = uint

// See: https://developer.apple.com/documentation/kernel/kernelid
type KernelID = uintptr

// See: https://developer.apple.com/documentation/kernel/logicaladdress
type LogicalAddress = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/longlbamodeparameterblockdescriptor
// LongLBAModeParameterBlockDescriptor is opaque storage with the size and alignment C gives LongLBAModeParameterBlockDescriptor:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type LongLBAModeParameterBlockDescriptor [16]byte

// See: https://developer.apple.com/documentation/kernel/lowlatencyuserbufferinfo
// LowLatencyUserBufferInfo is opaque storage with the size and alignment C gives LowLatencyUserBufferInfo:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type LowLatencyUserBufferInfo [5]uint64

// See: https://developer.apple.com/documentation/kernel/lowlatencyuserbufferinfov2
// LowLatencyUserBufferInfoV2 is opaque storage with the size and alignment C gives LowLatencyUserBufferInfoV2:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type LowLatencyUserBufferInfoV2 [6]uint64

// See: https://developer.apple.com/documentation/kernel/lowlatencyuserbufferinfov3
// LowLatencyUserBufferInfoV3 is opaque storage with the size and alignment C gives LowLatencyUserBufferInfoV3:
// 56 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 56 into.
type LowLatencyUserBufferInfoV3 [7]uint64

// See: https://developer.apple.com/documentation/kernel/md5_ctx
// MD5_CTX is opaque storage with the size and alignment C gives MD5_CTX:
// 88 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 88 into.
type MD5_CTX [22]uint32

// See: https://developer.apple.com/documentation/kernel/masteraudiofunctions
// MasterAudioFunctions is an unresolved C aggregate typedef.
type MasterAudioFunctions unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mastermuteupdate
type MasterMuteUpdate = bool

// See: https://developer.apple.com/documentation/kernel/mastervolumeupdate
type MasterVolumeUpdate = uint16

// See: https://developer.apple.com/documentation/kernel/modepageformatheader
// ModePageFormatHeader is opaque storage with the size and alignment C gives ModePageFormatHeader:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type ModePageFormatHeader [2]byte

// See: https://developer.apple.com/documentation/kernel/modeparameterblockdescriptor
// ModeParameterBlockDescriptor is opaque storage with the size and alignment C gives ModeParameterBlockDescriptor:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type ModeParameterBlockDescriptor [8]byte

// See: https://developer.apple.com/documentation/kernel/ndr_record_t
// NDR_record_t is opaque storage with the size and alignment C gives NDR_record_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type NDR_record_t [8]byte

// See: https://developer.apple.com/documentation/kernel/nxeqelement
// NXEQElement is opaque storage with the size and alignment C gives NXEQElement:
// 96 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 96 into.
type NXEQElement [24]uint32

// See: https://developer.apple.com/documentation/kernel/nxevent
// NXEvent is opaque storage with the size and alignment C gives NXEvent:
// 88 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 88 into.
type NXEvent [22]uint32

// See: https://developer.apple.com/documentation/kernel/nxeventdata
// NXEventData is opaque storage with the size and alignment C gives NXEventData:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type NXEventData [12]uint32

// See: https://developer.apple.com/documentation/kernel/nxeventext
// NXEventExt is opaque storage with the size and alignment C gives NXEventExt:
// 124 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 124 into.
type NXEventExt [31]uint32

// See: https://developer.apple.com/documentation/kernel/nxeventextension
// NXEventExtension is opaque storage with the size and alignment C gives NXEventExtension:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type NXEventExtension [9]uint32

// See: https://developer.apple.com/documentation/kernel/nxeventptr
type NXEventPtr = uintptr

// See: https://developer.apple.com/documentation/kernel/nxeventsystemdevice
// NXEventSystemDevice is opaque storage with the size and alignment C gives NXEventSystemDevice:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type NXEventSystemDevice [4]uint32

// See: https://developer.apple.com/documentation/kernel/nxeventsystemdevicelist
// NXEventSystemDeviceList is opaque storage with the size and alignment C gives NXEventSystemDeviceList:
// 256 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 256 into.
type NXEventSystemDeviceList [64]uint32

// See: https://developer.apple.com/documentation/kernel/nxeventsysteminfodata
type NXEventSystemInfoData = int32

// See: https://developer.apple.com/documentation/kernel/nxeventsysteminfotype
type NXEventSystemInfoType = *int32

// See: https://developer.apple.com/documentation/kernel/nxkeymapping
// NXKeyMapping is opaque storage with the size and alignment C gives NXKeyMapping:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type NXKeyMapping [2]uint64

// See: https://developer.apple.com/documentation/kernel/nxmousescaling
// NXMouseScaling is opaque storage with the size and alignment C gives NXMouseScaling:
// 84 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 84 into.
type NXMouseScaling [21]uint32

// See: https://developer.apple.com/documentation/kernel/nxparsedkeymapping
// NXParsedKeyMapping is opaque storage with the size and alignment C gives NXParsedKeyMapping:
// 3552 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 3552 into.
type NXParsedKeyMapping [444]uint64

// See: https://developer.apple.com/documentation/kernel/nxswappeddouble
type NXSwappedDouble = uint64

// See: https://developer.apple.com/documentation/kernel/nxswappedfloat
type NXSwappedFloat = uint

// See: https://developer.apple.com/documentation/kernel/nxtabletpointdata
// NXTabletPointData is opaque storage with the size and alignment C gives NXTabletPointData:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type NXTabletPointData [8]uint32

// See: https://developer.apple.com/documentation/kernel/nxtabletpointdataptr
type NXTabletPointDataPtr = uintptr

// See: https://developer.apple.com/documentation/kernel/nxtabletproximitydata
// NXTabletProximityData is opaque storage with the size and alignment C gives NXTabletProximityData:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type NXTabletProximityData [8]uint32

// See: https://developer.apple.com/documentation/kernel/nxtabletproximitydataptr
type NXTabletProximityDataPtr = uintptr

// See: https://developer.apple.com/documentation/kernel/nudclflags
type NuDCLFlags = uint

// See: https://developer.apple.com/documentation/kernel/nudclreceivepacketref
type NuDCLReceivePacketRef uintptr

// See: https://developer.apple.com/documentation/kernel/nudclref
type NuDCLRef uintptr

// See: https://developer.apple.com/documentation/kernel/nudclsendpacketref
type NuDCLSendPacketRef uintptr

// See: https://developer.apple.com/documentation/kernel/nudclskipcycleref
type NuDCLSkipCycleRef uintptr

// See: https://developer.apple.com/documentation/kernel/numversion
// NumVersion is opaque storage with the size and alignment C gives NumVersion:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type NumVersion [4]byte

// See: https://developer.apple.com/documentation/kernel/osactionabortedhandler
type OSActionAbortedHandler = func()

// See: https://developer.apple.com/documentation/kernel/osactioncancelhandler
type OSActionCancelHandler = func()

// See: https://developer.apple.com/documentation/kernel/osarrayptr
type OSArrayPtr = *OSArray

// See: https://developer.apple.com/documentation/kernel/osasyncreference64
type OSAsyncReference64 = uint64

// See: https://developer.apple.com/documentation/kernel/osasyncreference
type OSAsyncReference = Natural_t

// See: https://developer.apple.com/documentation/kernel/osbooleanptr
type OSBooleanPtr = *OSBoolean

// See: https://developer.apple.com/documentation/kernel/oscollectioniteratorptr
type OSCollectionIteratorPtr = *OSCollectionIterator

// See: https://developer.apple.com/documentation/kernel/oscollectionptr
type OSCollectionPtr = *OSCollection

// See: https://developer.apple.com/documentation/kernel/oscontainer
type OSContainer = OSObject

// See: https://developer.apple.com/documentation/kernel/osdataconstptr
type OSDataConstPtr = *OSData

// See: https://developer.apple.com/documentation/kernel/osdataptr
type OSDataPtr = *OSData

// See: https://developer.apple.com/documentation/kernel/osdictionaryptr
type OSDictionaryPtr = *OSDictionary

// OSKextRequestTag is identifies a kext request made to user space.
//
// See: https://developer.apple.com/documentation/kernel/oskextrequesttag
type OSKextRequestTag = uint32

// OSMallocTag is an opaque type used to track memory allocations.
//
// See: https://developer.apple.com/documentation/kernel/osmalloctag
type OSMallocTag = uintptr

// OSMallocTag_t is see [OSMallocTag].
//
// See: https://developer.apple.com/documentation/kernel/osmalloctag_t
type OSMallocTag_t = uintptr

// See: https://developer.apple.com/documentation/kernel/osnumberptr
type OSNumberPtr = *OSNumber

// See: https://developer.apple.com/documentation/kernel/osobjectapplierblock
type OSObjectApplierBlock = func(object *OSObject)

// See: https://developer.apple.com/documentation/kernel/osobjectptr
type OSObjectPtr = *OSObject

// See: https://developer.apple.com/documentation/kernel/osobjectref
type OSObjectRef = uint64

// See: https://developer.apple.com/documentation/kernel/osorderedsetptr
type OSOrderedSetPtr = *OSOrderedSet

// See: https://developer.apple.com/documentation/kernel/osserializeptr
type OSSerializePtr = *OSSerialize

// See: https://developer.apple.com/documentation/kernel/osserializerblock
type OSSerializerBlock = func(serializer *OSSerialize) bool

// See: https://developer.apple.com/documentation/kernel/osserializerptr
type OSSerializerPtr = *OSSerializer

// See: https://developer.apple.com/documentation/kernel/ossetptr
type OSSetPtr = *OSSet

// See: https://developer.apple.com/documentation/kernel/osstringconstptr
type OSStringConstPtr = *OSString

// See: https://developer.apple.com/documentation/kernel/osstringptr
type OSStringPtr = *OSString

// See: https://developer.apple.com/documentation/kernel/ossymbolconstptr
type OSSymbolConstPtr = *OSSymbol

// See: https://developer.apple.com/documentation/kernel/ossymbolptr
type OSSymbolPtr = *OSSymbol

// See: https://developer.apple.com/documentation/kernel/ostype
type OSType = uint32

// See: https://developer.apple.com/documentation/kernel/ostypeptr
type OSTypePtr = *uint32

// See: https://developer.apple.com/documentation/kernel/opaquedtentryiterator
// OpaqueDTEntryIterator is opaque storage with the size and alignment C gives OpaqueDTEntryIterator:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type OpaqueDTEntryIterator [5]uint64

// See: https://developer.apple.com/documentation/kernel/opaquedtpropertyiterator
// OpaqueDTPropertyIterator is opaque storage with the size and alignment C gives OpaqueDTPropertyIterator:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type OpaqueDTPropertyIterator [3]uint64

// See: https://developer.apple.com/documentation/kernel/pbversion
type PBVersion = uint32

// See: https://developer.apple.com/documentation/kernel/pe_video
// PE_Video is opaque storage with the size and alignment C gives PE_Video:
// 144 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 144 into.
type PE_Video [18]uint64

// See: https://developer.apple.com/documentation/kernel/pe_state_t
// PE_state_t is opaque storage with the size and alignment C gives PE_state_t:
// 176 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 176 into.
type PE_state_t [22]uint64

// See: https://developer.apple.com/documentation/kernel/ptr
type Ptr = *byte

// See: https://developer.apple.com/documentation/kernel/raw_header
// RAW_header is opaque storage with the size and alignment C gives RAW_header:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type RAW_header [3]uint64

// See: https://developer.apple.com/documentation/kernel/report_luns_logical_unit_addressing
// REPORT_LUNS_LOGICAL_UNIT_ADDRESSING is opaque storage with the size and alignment C gives REPORT_LUNS_LOGICAL_UNIT_ADDRESSING:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type REPORT_LUNS_LOGICAL_UNIT_ADDRESSING [1]uint16

// See: https://developer.apple.com/documentation/kernel/report_luns_peripheral_device_addressing
// REPORT_LUNS_PERIPHERAL_DEVICE_ADDRESSING is opaque storage with the size and alignment C gives REPORT_LUNS_PERIPHERAL_DEVICE_ADDRESSING:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type REPORT_LUNS_PERIPHERAL_DEVICE_ADDRESSING [1]uint16

// See: https://developer.apple.com/documentation/kernel/rgbcolor
// RGBColor is opaque storage with the size and alignment C gives RGBColor:
// 6 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 6 into.
type RGBColor [3]uint16

// See: https://developer.apple.com/documentation/kernel/rgbcolorhdl
type RGBColorHdl = *RGBColorPtr

// See: https://developer.apple.com/documentation/kernel/rgbcolorptr
type RGBColorPtr = *RGBColor

// See: https://developer.apple.com/documentation/kernel/rawsensecode
type RawSenseCode = byte

// See: https://developer.apple.com/documentation/kernel/realdtentry
type RealDTEntry = unsafe.Pointer

// RectPtr is represents a type used by the Video Components API.
//
// See: https://developer.apple.com/documentation/kernel/rectptr
type RectPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/regcstrentryname
type RegCStrEntryName = int8

// See: https://developer.apple.com/documentation/kernel/regcstrentrynamebuf
type RegCStrEntryNameBuf = int8

// See: https://developer.apple.com/documentation/kernel/regcstrentrynameptr
type RegCStrEntryNamePtr = *byte

// See: https://developer.apple.com/documentation/kernel/regcstrpathname
type RegCStrPathName = int8

// See: https://developer.apple.com/documentation/kernel/regentryid
// RegEntryID is opaque storage with the size and alignment C gives RegEntryID:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type RegEntryID [4]uint64

// See: https://developer.apple.com/documentation/kernel/regentryidptr
type RegEntryIDPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/regentryiter
type RegEntryIter = *IORegistryIterator

// See: https://developer.apple.com/documentation/kernel/regentryiterationop
type RegEntryIterationOp = uint32

// See: https://developer.apple.com/documentation/kernel/regentrymodifiers
type RegEntryModifiers = uint32

// See: https://developer.apple.com/documentation/kernel/regiterationop
type RegIterationOp = uint32

// See: https://developer.apple.com/documentation/kernel/regmodifiers
type RegModifiers = uint32

// See: https://developer.apple.com/documentation/kernel/regpathnamesize
type RegPathNameSize = uint32

// See: https://developer.apple.com/documentation/kernel/regpropertyiter
type RegPropertyIter = *OSIterator

// See: https://developer.apple.com/documentation/kernel/regpropertymodifiers
type RegPropertyModifiers = uint32

// See: https://developer.apple.com/documentation/kernel/regpropertyname
type RegPropertyName = int8

// See: https://developer.apple.com/documentation/kernel/regpropertynamebuf
type RegPropertyNameBuf = int8

// See: https://developer.apple.com/documentation/kernel/regpropertynameptr
type RegPropertyNamePtr = *byte

// See: https://developer.apple.com/documentation/kernel/regpropertyvalue
type RegPropertyValue = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/regpropertyvaluesize
type RegPropertyValueSize = uint32

// See: https://developer.apple.com/documentation/kernel/restype
type ResType = uint32

// See: https://developer.apple.com/documentation/kernel/restypeptr
type ResTypePtr = *uint32

// See: https://developer.apple.com/documentation/kernel/runtimeoptions
type RuntimeOptions = uint32

// See: https://developer.apple.com/documentation/kernel/sbcmodepagecaching
// SBCModePageCaching is opaque storage with the size and alignment C gives SBCModePageCaching:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type SBCModePageCaching [20]byte

// See: https://developer.apple.com/documentation/kernel/sbcmodepageflexibledisk
// SBCModePageFlexibleDisk is opaque storage with the size and alignment C gives SBCModePageFlexibleDisk:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type SBCModePageFlexibleDisk [32]byte

// See: https://developer.apple.com/documentation/kernel/sbcmodepageformatdevice
// SBCModePageFormatDevice is opaque storage with the size and alignment C gives SBCModePageFormatDevice:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type SBCModePageFormatDevice [24]byte

// See: https://developer.apple.com/documentation/kernel/sbcmodepagerigiddiskgeometry
// SBCModePageRigidDiskGeometry is opaque storage with the size and alignment C gives SBCModePageRigidDiskGeometry:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type SBCModePageRigidDiskGeometry [24]byte

// See: https://developer.apple.com/documentation/kernel/scsicmdfield10bit
type SCSICmdField10Bit = uint16

// See: https://developer.apple.com/documentation/kernel/scsicmdfield11bit
type SCSICmdField11Bit = uint16

// See: https://developer.apple.com/documentation/kernel/scsicmdfield12bit
type SCSICmdField12Bit = uint16

// See: https://developer.apple.com/documentation/kernel/scsicmdfield13bit
type SCSICmdField13Bit = uint16

// See: https://developer.apple.com/documentation/kernel/scsicmdfield14bit
type SCSICmdField14Bit = uint16

// See: https://developer.apple.com/documentation/kernel/scsicmdfield15bit
type SCSICmdField15Bit = uint16

// See: https://developer.apple.com/documentation/kernel/scsicmdfield17bit
type SCSICmdField17Bit = uint32

// See: https://developer.apple.com/documentation/kernel/scsicmdfield18bit
type SCSICmdField18Bit = uint32

// See: https://developer.apple.com/documentation/kernel/scsicmdfield19bit
type SCSICmdField19Bit = uint32

// See: https://developer.apple.com/documentation/kernel/scsicmdfield1bit
type SCSICmdField1Bit = byte

// See: https://developer.apple.com/documentation/kernel/scsicmdfield1byte
type SCSICmdField1Byte = byte

// See: https://developer.apple.com/documentation/kernel/scsicmdfield20bit
type SCSICmdField20Bit = uint32

// See: https://developer.apple.com/documentation/kernel/scsicmdfield21bit
type SCSICmdField21Bit = uint32

// See: https://developer.apple.com/documentation/kernel/scsicmdfield22bit
type SCSICmdField22Bit = uint32

// See: https://developer.apple.com/documentation/kernel/scsicmdfield23bit
type SCSICmdField23Bit = uint32

// See: https://developer.apple.com/documentation/kernel/scsicmdfield25bit
type SCSICmdField25Bit = uint32

// See: https://developer.apple.com/documentation/kernel/scsicmdfield26bit
type SCSICmdField26Bit = uint32

// See: https://developer.apple.com/documentation/kernel/scsicmdfield27bit
type SCSICmdField27Bit = uint32

// See: https://developer.apple.com/documentation/kernel/scsicmdfield28bit
type SCSICmdField28Bit = uint32

// See: https://developer.apple.com/documentation/kernel/scsicmdfield29bit
type SCSICmdField29Bit = uint32

// See: https://developer.apple.com/documentation/kernel/scsicmdfield2bit
type SCSICmdField2Bit = byte

// See: https://developer.apple.com/documentation/kernel/scsicmdfield2byte
type SCSICmdField2Byte = uint16

// See: https://developer.apple.com/documentation/kernel/scsicmdfield30bit
type SCSICmdField30Bit = uint32

// See: https://developer.apple.com/documentation/kernel/scsicmdfield31bit
type SCSICmdField31Bit = uint32

// See: https://developer.apple.com/documentation/kernel/scsicmdfield33bit
type SCSICmdField33Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield34bit
type SCSICmdField34Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield35bit
type SCSICmdField35Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield36bit
type SCSICmdField36Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield37bit
type SCSICmdField37Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield38bit
type SCSICmdField38Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield39bit
type SCSICmdField39Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield3bit
type SCSICmdField3Bit = byte

// See: https://developer.apple.com/documentation/kernel/scsicmdfield3byte
type SCSICmdField3Byte = uint32

// See: https://developer.apple.com/documentation/kernel/scsicmdfield41bit
type SCSICmdField41Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield42bit
type SCSICmdField42Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield43bit
type SCSICmdField43Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield44bit
type SCSICmdField44Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield45bit
type SCSICmdField45Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield46bit
type SCSICmdField46Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield47bit
type SCSICmdField47Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield49bit
type SCSICmdField49Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield4bit
type SCSICmdField4Bit = byte

// See: https://developer.apple.com/documentation/kernel/scsicmdfield4byte
type SCSICmdField4Byte = uint32

// See: https://developer.apple.com/documentation/kernel/scsicmdfield50bit
type SCSICmdField50Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield51bit
type SCSICmdField51Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield52bit
type SCSICmdField52Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield53bit
type SCSICmdField53Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield54bit
type SCSICmdField54Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield55bit
type SCSICmdField55Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield57bit
type SCSICmdField57Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield58bit
type SCSICmdField58Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield59bit
type SCSICmdField59Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield5bit
type SCSICmdField5Bit = byte

// See: https://developer.apple.com/documentation/kernel/scsicmdfield5byte
type SCSICmdField5Byte = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield60bit
type SCSICmdField60Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield61bit
type SCSICmdField61Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield62bit
type SCSICmdField62Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield63bit
type SCSICmdField63Bit = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield6bit
type SCSICmdField6Bit = byte

// See: https://developer.apple.com/documentation/kernel/scsicmdfield6byte
type SCSICmdField6Byte = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield7bit
type SCSICmdField7Bit = byte

// See: https://developer.apple.com/documentation/kernel/scsicmdfield7byte
type SCSICmdField7Byte = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield8byte
type SCSICmdField8Byte = uint64

// See: https://developer.apple.com/documentation/kernel/scsicmdfield9bit
type SCSICmdField9Bit = uint16

// See: https://developer.apple.com/documentation/kernel/scsicmd_inquiry_pagecx_header
// SCSICmd_INQUIRY_PAGECx_Header is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_PAGECx_Header:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type SCSICmd_INQUIRY_PAGECx_Header [4]byte

// See: https://developer.apple.com/documentation/kernel/scsicmd_inquiry_page00_header_spc_16
// SCSICmd_INQUIRY_Page00_Header_SPC_16 is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_Page00_Header_SPC_16:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type SCSICmd_INQUIRY_Page00_Header_SPC_16 [2]uint16

// See: https://developer.apple.com/documentation/kernel/scsicmd_inquiry_page80_header_spc_16
// SCSICmd_INQUIRY_Page80_Header_SPC_16 is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_Page80_Header_SPC_16:
// 6 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 6 into.
type SCSICmd_INQUIRY_Page80_Header_SPC_16 [3]uint16

// See: https://developer.apple.com/documentation/kernel/scsicmd_inquiry_pageb0_data
// SCSICmd_INQUIRY_PageB0_Data is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_PageB0_Data:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type SCSICmd_INQUIRY_PageB0_Data [64]byte

// See: https://developer.apple.com/documentation/kernel/scsicmd_inquiry_pageb2_data
// SCSICmd_INQUIRY_PageB2_Data is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_PageB2_Data:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type SCSICmd_INQUIRY_PageB2_Data [8]byte

// See: https://developer.apple.com/documentation/kernel/scsicmd_inquiry_pageb2_provisioning_group_descriptor
// SCSICmd_INQUIRY_PageB2_Provisioning_Group_Descriptor is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_PageB2_Provisioning_Group_Descriptor:
// 38 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 38 into.
type SCSICmd_INQUIRY_PageB2_Provisioning_Group_Descriptor [38]byte

// See: https://developer.apple.com/documentation/kernel/scsicmd_inquiry_pagec0_data
// SCSICmd_INQUIRY_PageC0_Data is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_PageC0_Data:
// 116 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 116 into.
type SCSICmd_INQUIRY_PageC0_Data [116]byte

// See: https://developer.apple.com/documentation/kernel/scsicmd_inquiry_pagec1_data
// SCSICmd_INQUIRY_PageC1_Data is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_PageC1_Data:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type SCSICmd_INQUIRY_PageC1_Data [12]byte

// See: https://developer.apple.com/documentation/kernel/scsicmd_inquiry_standarddataptr
type SCSICmd_INQUIRY_StandardDataPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmd_report_luns_header
// SCSICmd_REPORT_LUNS_Header is opaque storage with the size and alignment C gives SCSICmd_REPORT_LUNS_Header:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type SCSICmd_REPORT_LUNS_Header [4]uint32

// See: https://developer.apple.com/documentation/kernel/scsicmd_report_luns_lun_entry
// SCSICmd_REPORT_LUNS_LUN_ENTRY is opaque storage with the size and alignment C gives SCSICmd_REPORT_LUNS_LUN_ENTRY:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type SCSICmd_REPORT_LUNS_LUN_ENTRY [4]uint16

// See: https://developer.apple.com/documentation/kernel/scsicommanddescriptorblock
type SCSICommandDescriptorBlock = byte

// SCSIDeviceIdentifier is 64-bit number to represent a SCSI Device.
//
// See: https://developer.apple.com/documentation/kernel/scsideviceidentifier
type SCSIDeviceIdentifier = uint64

// SCSIInitiatorIdentifier is 64-bit number to represent a SCSI Initiator Device.
//
// See: https://developer.apple.com/documentation/kernel/scsiinitiatoridentifier
type SCSIInitiatorIdentifier = uint64

// See: https://developer.apple.com/documentation/kernel/scsilogicalunitbytes
type SCSILogicalUnitBytes = byte

// See: https://developer.apple.com/documentation/kernel/scsilogicalunitnumber
type SCSILogicalUnitNumber = uint64

// See: https://developer.apple.com/documentation/kernel/scsiparalleltaskidentifier
type SCSIParallelTaskIdentifier = *OSObject

// SCSIPortStatus is 32-bit number to represent a SCSIPortStatus.
//
// See: https://developer.apple.com/documentation/kernel/scsiportstatus
type SCSIPortStatus = uint32

// See: https://developer.apple.com/documentation/kernel/scsiprotocolfeature
type SCSIProtocolFeature = uint32

// See: https://developer.apple.com/documentation/kernel/scsiprotocolpowerstate
type SCSIProtocolPowerState = uint32

// SCSIServiceResponse is attributes for task service response.
//
// See: https://developer.apple.com/documentation/kernel/scsiserviceresponse
type SCSIServiceResponse = uint32

// SCSITaggedTaskIdentifier is 64-bit number to represent a unique task identifier.
//
// See: https://developer.apple.com/documentation/kernel/scsitaggedtaskidentifier
type SCSITaggedTaskIdentifier = uint64

// SCSITargetIdentifier is 64-bit number to represent a SCSI Target Device.
//
// See: https://developer.apple.com/documentation/kernel/scsitargetidentifier
type SCSITargetIdentifier = uint64

// SCSITaskAttribute is attributes for task delivery.
//
// See: https://developer.apple.com/documentation/kernel/scsitaskattribute
type SCSITaskAttribute = uint32

// See: https://developer.apple.com/documentation/kernel/scsitaskidentifier
type SCSITaskIdentifier = *OSObject

// See: https://developer.apple.com/documentation/kernel/scsitaskmode
type SCSITaskMode = uint32

// SCSITaskState is attributes for task state.
//
// See: https://developer.apple.com/documentation/kernel/scsitaskstate
type SCSITaskState = uint32

// SCSITaskStatus is attributes for task status.
//
// See: https://developer.apple.com/documentation/kernel/scsitaskstatus
type SCSITaskStatus = uint32

// See: https://developer.apple.com/documentation/kernel/scsi_sense_data
// SCSI_Sense_Data is opaque storage with the size and alignment C gives SCSI_Sense_Data:
// 18 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 18 into.
type SCSI_Sense_Data [18]byte

// See: https://developer.apple.com/documentation/kernel/sc_scatter
// SC_Scatter is opaque storage with the size and alignment C gives SC_Scatter:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type SC_Scatter [3]uint64

// See: https://developer.apple.com/documentation/kernel/sha1_ctx
// SHA1_CTX is opaque storage with the size and alignment C gives SHA1_CTX:
// 104 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 104 into.
type SHA1_CTX [13]uint64

// See: https://developer.apple.com/documentation/kernel/sint
type SInt = int32

// See: https://developer.apple.com/documentation/kernel/sint16
type SInt16 = int16

// See: https://developer.apple.com/documentation/kernel/sint32
type SInt32 = int32

// See: https://developer.apple.com/documentation/kernel/sint64
type SInt64 = int64

// See: https://developer.apple.com/documentation/kernel/sint8
type SInt8 = int8

// See: https://developer.apple.com/documentation/kernel/spcmodepagepowercondition
// SPCModePagePowerCondition is opaque storage with the size and alignment C gives SPCModePagePowerCondition:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type SPCModePagePowerCondition [12]byte

// See: https://developer.apple.com/documentation/kernel/spcmodeparameterheader10
// SPCModeParameterHeader10 is opaque storage with the size and alignment C gives SPCModeParameterHeader10:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type SPCModeParameterHeader10 [8]byte

// See: https://developer.apple.com/documentation/kernel/spcmodeparameterheader6
// SPCModeParameterHeader6 is opaque storage with the size and alignment C gives SPCModeParameterHeader6:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type SPCModeParameterHeader6 [4]byte

// See: https://developer.apple.com/documentation/kernel/servicecount
type ServiceCount = uint32

// See: https://developer.apple.com/documentation/kernel/signedbyte
type SignedByte = int8

// See: https://developer.apple.com/documentation/kernel/stickykeys_modifierinfo
// StickyKeys_ModifierInfo is opaque storage with the size and alignment C gives StickyKeys_ModifierInfo:
// 3 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 3 into.
type StickyKeys_ModifierInfo [3]byte

// See: https://developer.apple.com/documentation/kernel/stickykeys_toggleinfo
// StickyKeys_ToggleInfo is opaque storage with the size and alignment C gives StickyKeys_ToggleInfo:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type StickyKeys_ToggleInfo [5]uint64

// See: https://developer.apple.com/documentation/kernel/str31
type Str31 = byte

// See: https://developer.apple.com/documentation/kernel/transmissionpower
type TransmissionPower = int8

// See: https://developer.apple.com/documentation/kernel/uaspipedescriptor
// UASPipeDescriptor is opaque storage with the size and alignment C gives UASPipeDescriptor:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type UASPipeDescriptor [4]byte

// See: https://developer.apple.com/documentation/kernel/uaspipedescriptorptr
type UASPipeDescriptorPtr = *UASPipeDescriptor

// See: https://developer.apple.com/documentation/kernel/uint16
type UInt16 = uint16

// See: https://developer.apple.com/documentation/kernel/uint32
type UInt32 = uint32

// See: https://developer.apple.com/documentation/kernel/uint32ptr
type UInt32Ptr = *uint32

// See: https://developer.apple.com/documentation/kernel/uint64
type UInt64 = uint64

// See: https://developer.apple.com/documentation/kernel/uint8
type UInt8 = byte

// See: https://developer.apple.com/documentation/kernel/undkey
type UNDKey = *byte

// See: https://developer.apple.com/documentation/kernel/undlabel
type UNDLabel = *byte

// See: https://developer.apple.com/documentation/kernel/undmessage
type UNDMessage = *byte

// See: https://developer.apple.com/documentation/kernel/undpath
type UNDPath = *byte

// See: https://developer.apple.com/documentation/kernel/undreplyref
type UNDReplyRef = uint32

// See: https://developer.apple.com/documentation/kernel/undserverref
type UNDServerRef = uint32

// USBDeviceAddress is a USB device address.
//
// See: https://developer.apple.com/documentation/kernel/usbdeviceaddress
type USBDeviceAddress = uint16

// USBDeviceInformationBits is the state of a USB device.
//
// See: https://developer.apple.com/documentation/kernel/usbdeviceinformationbits
type USBDeviceInformationBits = int

// USBLowLatencyBufferType is specifies which kind of low-latency buffer to create.
//
// See: https://developer.apple.com/documentation/kernel/usblowlatencybuffertype
type USBLowLatencyBufferType = int

// USBPhysicalAddress32 is a 32-bit USB physical address.
//
// See: https://developer.apple.com/documentation/kernel/usbphysicaladdress32
type USBPhysicalAddress32 = uint32

// USBPowerRequestTypes is specifies the kind of power to reserve.
//
// See: https://developer.apple.com/documentation/kernel/usbpowerrequesttypes
type USBPowerRequestTypes = int

// USBStatus is the value of the USB device status.
//
// See: https://developer.apple.com/documentation/kernel/usbstatus
type USBStatus = uint16

// USBStatusPtr is a pointer to a USB status.
//
// See: https://developer.apple.com/documentation/kernel/usbstatusptr
type USBStatusPtr = *USBStatus

// See: https://developer.apple.com/documentation/kernel/unichar
type UniChar = uint16

// See: https://developer.apple.com/documentation/kernel/userexportdclcallcommandproc
type UserExportDCLCallCommandProc = *objc.ID

// See: https://developer.apple.com/documentation/kernel/userexportdclcallproc
// UserExportDCLCallProc is opaque storage with the size and alignment C gives UserExportDCLCallProc:
// 44 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 44 into.
type UserExportDCLCallProc [44]byte

// See: https://developer.apple.com/documentation/kernel/userexportdclcommand
// UserExportDCLCommand is opaque storage with the size and alignment C gives UserExportDCLCommand:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type UserExportDCLCommand [32]byte

// See: https://developer.apple.com/documentation/kernel/userexportdcljump
// UserExportDCLJump is opaque storage with the size and alignment C gives UserExportDCLJump:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type UserExportDCLJump [36]byte

// See: https://developer.apple.com/documentation/kernel/userexportdcllabel
// UserExportDCLLabel is opaque storage with the size and alignment C gives UserExportDCLLabel:
// 28 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 28 into.
type UserExportDCLLabel [28]byte

// See: https://developer.apple.com/documentation/kernel/userexportdclnudclleader
// UserExportDCLNuDCLLeader is opaque storage with the size and alignment C gives UserExportDCLNuDCLLeader:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type UserExportDCLNuDCLLeader [36]byte

// See: https://developer.apple.com/documentation/kernel/userexportdclptrtimestamp
// UserExportDCLPtrTimeStamp is opaque storage with the size and alignment C gives UserExportDCLPtrTimeStamp:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type UserExportDCLPtrTimeStamp [36]byte

// See: https://developer.apple.com/documentation/kernel/userexportdclsettagsyncbits
// UserExportDCLSetTagSyncBits is opaque storage with the size and alignment C gives UserExportDCLSetTagSyncBits:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type UserExportDCLSetTagSyncBits [32]byte

// See: https://developer.apple.com/documentation/kernel/userexportdcltimestamp
// UserExportDCLTimeStamp is opaque storage with the size and alignment C gives UserExportDCLTimeStamp:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type UserExportDCLTimeStamp [32]byte

// See: https://developer.apple.com/documentation/kernel/userexportdcltransferbuffer
// UserExportDCLTransferBuffer is opaque storage with the size and alignment C gives UserExportDCLTransferBuffer:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type UserExportDCLTransferBuffer [48]byte

// See: https://developer.apple.com/documentation/kernel/userexportdcltransferpacket
// UserExportDCLTransferPacket is opaque storage with the size and alignment C gives UserExportDCLTransferPacket:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type UserExportDCLTransferPacket [40]byte

// See: https://developer.apple.com/documentation/kernel/userexportdclupdatedcllist
// UserExportDCLUpdateDCLList is opaque storage with the size and alignment C gives UserExportDCLUpdateDCLList:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type UserExportDCLUpdateDCLList [40]byte

// See: https://developer.apple.com/documentation/kernel/vdclutbehavior
type VDClutBehavior = uint32

// See: https://developer.apple.com/documentation/kernel/vdclutbehaviorptr
type VDClutBehaviorPtr = *VDClutBehavior

// See: https://developer.apple.com/documentation/kernel/vdcommunicationinfoptr
type VDCommunicationInfoPtr = *VDCommunicationInfoRec

// See: https://developer.apple.com/documentation/kernel/vdcommunicationinforec
// VDCommunicationInfoRec is opaque storage with the size and alignment C gives VDCommunicationInfoRec:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type VDCommunicationInfoRec [12]uint32

// See: https://developer.apple.com/documentation/kernel/vdcommunicationptr
type VDCommunicationPtr = *VDCommunicationRec

// See: https://developer.apple.com/documentation/kernel/vdcommunicationrec
// VDCommunicationRec is opaque storage with the size and alignment C gives VDCommunicationRec:
// 80 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 80 into.
type VDCommunicationRec [10]uint64

// See: https://developer.apple.com/documentation/kernel/vdconfigurationfeaturelistrec
// VDConfigurationFeatureListRec is opaque storage with the size and alignment C gives VDConfigurationFeatureListRec:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type VDConfigurationFeatureListRec [4]uint64

// See: https://developer.apple.com/documentation/kernel/vdconfigurationfeaturelistrecptr
type VDConfigurationFeatureListRecPtr = *VDConfigurationFeatureListRec

// See: https://developer.apple.com/documentation/kernel/vdconfigurationptr
type VDConfigurationPtr = *VDConfigurationRec

// See: https://developer.apple.com/documentation/kernel/vdconfigurationrec
// VDConfigurationRec is opaque storage with the size and alignment C gives VDConfigurationRec:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type VDConfigurationRec [4]uint64

// See: https://developer.apple.com/documentation/kernel/vdconvolutioninfoptr
type VDConvolutionInfoPtr = *VDConvolutionInfoRec

// See: https://developer.apple.com/documentation/kernel/vdconvolutioninforec
// VDConvolutionInfoRec is opaque storage with the size and alignment C gives VDConvolutionInfoRec:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type VDConvolutionInfoRec [5]uint32

// See: https://developer.apple.com/documentation/kernel/vdddcblockptr
type VDDDCBlockPtr = *VDDDCBlockRec

// See: https://developer.apple.com/documentation/kernel/vdddcblockrec
// VDDDCBlockRec is opaque storage with the size and alignment C gives VDDDCBlockRec:
// 144 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 144 into.
type VDDDCBlockRec [36]uint32

// See: https://developer.apple.com/documentation/kernel/vddefmode
type VDDefMode = uint

// See: https://developer.apple.com/documentation/kernel/vddefmodeptr
type VDDefModePtr = *VDDefMode

// See: https://developer.apple.com/documentation/kernel/vddetailedtimingptr
type VDDetailedTimingPtr = *VDDetailedTimingRec

// See: https://developer.apple.com/documentation/kernel/vddetailedtimingrec
// VDDetailedTimingRec is opaque storage with the size and alignment C gives VDDetailedTimingRec:
// 160 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 160 into.
type VDDetailedTimingRec [20]uint64

// See: https://developer.apple.com/documentation/kernel/vddisplayconnectinfoptr
type VDDisplayConnectInfoPtr = *VDDisplayConnectInfoRec

// See: https://developer.apple.com/documentation/kernel/vddisplayconnectinforec
// VDDisplayConnectInfoRec is opaque storage with the size and alignment C gives VDDisplayConnectInfoRec:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type VDDisplayConnectInfoRec [3]uint64

// See: https://developer.apple.com/documentation/kernel/vddisplaytimingrangeptr
type VDDisplayTimingRangePtr = *VDDisplayTimingRangeRec

// See: https://developer.apple.com/documentation/kernel/vddisplaytimingrangerec
// VDDisplayTimingRangeRec is opaque storage with the size and alignment C gives VDDisplayTimingRangeRec:
// 240 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 240 into.
type VDDisplayTimingRangeRec [30]uint64

// See: https://developer.apple.com/documentation/kernel/vddrawhardwarecursorptr
type VDDrawHardwareCursorPtr = *VDDrawHardwareCursorRec

// See: https://developer.apple.com/documentation/kernel/vddrawhardwarecursorrec
// VDDrawHardwareCursorRec is opaque storage with the size and alignment C gives VDDrawHardwareCursorRec:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type VDDrawHardwareCursorRec [5]uint32

// See: https://developer.apple.com/documentation/kernel/vdentrecptr
type VDEntRecPtr = *VDEntryRecord

// See: https://developer.apple.com/documentation/kernel/vdentryrecord
// VDEntryRecord is opaque storage with the size and alignment C gives VDEntryRecord:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type VDEntryRecord [1]uint64

// See: https://developer.apple.com/documentation/kernel/vdflagrecptr
type VDFlagRecPtr = *VDFlagRecord

// See: https://developer.apple.com/documentation/kernel/vdflagrecord
// VDFlagRecord is opaque storage with the size and alignment C gives VDFlagRecord:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type VDFlagRecord [2]byte

// VDGamRecPtr is represents a type used by the Video Components API.
//
// See: https://developer.apple.com/documentation/kernel/vdgamrecptr
type VDGamRecPtr = *VDGammaRecord

// See: https://developer.apple.com/documentation/kernel/vdgammainfoptr
type VDGammaInfoPtr = *VDGammaInfoRec

// See: https://developer.apple.com/documentation/kernel/vdgammainforec
// VDGammaInfoRec is opaque storage with the size and alignment C gives VDGammaInfoRec:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type VDGammaInfoRec [3]uint64

// See: https://developer.apple.com/documentation/kernel/vdgammarecord
// VDGammaRecord is opaque storage with the size and alignment C gives VDGammaRecord:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type VDGammaRecord [1]uint64

// See: https://developer.apple.com/documentation/kernel/vdgetgammalistptr
type VDGetGammaListPtr = *VDGetGammaListRec

// See: https://developer.apple.com/documentation/kernel/vdgetgammalistrec
// VDGetGammaListRec is opaque storage with the size and alignment C gives VDGetGammaListRec:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type VDGetGammaListRec [3]uint64

// See: https://developer.apple.com/documentation/kernel/vdgrayptr
type VDGrayPtr = *VDGrayRecord

// See: https://developer.apple.com/documentation/kernel/vdgrayrecord
// VDGrayRecord is opaque storage with the size and alignment C gives VDGrayRecord:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type VDGrayRecord [2]byte

// See: https://developer.apple.com/documentation/kernel/vdhardwarecursordrawstateptr
type VDHardwareCursorDrawStatePtr = *VDHardwareCursorDrawStateRec

// See: https://developer.apple.com/documentation/kernel/vdhardwarecursordrawstaterec
// VDHardwareCursorDrawStateRec is opaque storage with the size and alignment C gives VDHardwareCursorDrawStateRec:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type VDHardwareCursorDrawStateRec [6]uint32

// See: https://developer.apple.com/documentation/kernel/vdmirrorptr
type VDMirrorPtr = *VDMirrorRec

// See: https://developer.apple.com/documentation/kernel/vdmirrorrec
// VDMirrorRec is opaque storage with the size and alignment C gives VDMirrorRec:
// 104 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 104 into.
type VDMirrorRec [13]uint64

// See: https://developer.apple.com/documentation/kernel/vdmulticonnectinfoptr
type VDMultiConnectInfoPtr = *VDMultiConnectInfoRec

// See: https://developer.apple.com/documentation/kernel/vdmulticonnectinforec
// VDMultiConnectInfoRec is opaque storage with the size and alignment C gives VDMultiConnectInfoRec:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type VDMultiConnectInfoRec [4]uint64

// See: https://developer.apple.com/documentation/kernel/vdpageinfo
// VDPageInfo is opaque storage with the size and alignment C gives VDPageInfo:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type VDPageInfo [3]uint64

// See: https://developer.apple.com/documentation/kernel/vdpginfoptr
type VDPgInfoPtr = *VDPageInfo

// See: https://developer.apple.com/documentation/kernel/vdpowerstateptr
type VDPowerStatePtr = *VDPowerStateRec

// See: https://developer.apple.com/documentation/kernel/vdpowerstaterec
// VDPowerStateRec is opaque storage with the size and alignment C gives VDPowerStateRec:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type VDPowerStateRec [3]uint64

// See: https://developer.apple.com/documentation/kernel/vdprivateselectordatarec
// VDPrivateSelectorDataRec is opaque storage with the size and alignment C gives VDPrivateSelectorDataRec:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type VDPrivateSelectorDataRec [4]uint64

// See: https://developer.apple.com/documentation/kernel/vdprivateselectorrec
// VDPrivateSelectorRec is opaque storage with the size and alignment C gives VDPrivateSelectorRec:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type VDPrivateSelectorRec [5]uint64

// See: https://developer.apple.com/documentation/kernel/vdresolutioninfoptr
type VDResolutionInfoPtr = *VDResolutionInfoRec

// See: https://developer.apple.com/documentation/kernel/vdresolutioninforec
// VDResolutionInfoRec is opaque storage with the size and alignment C gives VDResolutionInfoRec:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type VDResolutionInfoRec [5]uint64

// See: https://developer.apple.com/documentation/kernel/vdretrievegammaptr
type VDRetrieveGammaPtr = *VDRetrieveGammaRec

// See: https://developer.apple.com/documentation/kernel/vdretrievegammarec
// VDRetrieveGammaRec is opaque storage with the size and alignment C gives VDRetrieveGammaRec:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type VDRetrieveGammaRec [2]uint64

// See: https://developer.apple.com/documentation/kernel/vdscalerinfoptr
type VDScalerInfoPtr = *VDScalerInfoRec

// See: https://developer.apple.com/documentation/kernel/vdscalerinforec
// VDScalerInfoRec is opaque storage with the size and alignment C gives VDScalerInfoRec:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type VDScalerInfoRec [12]uint32

// See: https://developer.apple.com/documentation/kernel/vdscalerptr
type VDScalerPtr = *VDScalerRec

// See: https://developer.apple.com/documentation/kernel/vdscalerrec
// VDScalerRec is opaque storage with the size and alignment C gives VDScalerRec:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type VDScalerRec [16]uint32

// See: https://developer.apple.com/documentation/kernel/vdsetentryptr
type VDSetEntryPtr = *VDSetEntryRecord

// See: https://developer.apple.com/documentation/kernel/vdsetentryrecord
// VDSetEntryRecord is opaque storage with the size and alignment C gives VDSetEntryRecord:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type VDSetEntryRecord [2]uint64

// See: https://developer.apple.com/documentation/kernel/vdsethardwarecursorptr
type VDSetHardwareCursorPtr = *VDSetHardwareCursorRec

// See: https://developer.apple.com/documentation/kernel/vdsethardwarecursorrec
// VDSetHardwareCursorRec is opaque storage with the size and alignment C gives VDSetHardwareCursorRec:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type VDSetHardwareCursorRec [2]uint64

// See: https://developer.apple.com/documentation/kernel/vdsettings
// VDSettings is opaque storage with the size and alignment C gives VDSettings:
// 38 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 38 into.
type VDSettings [19]uint16

// See: https://developer.apple.com/documentation/kernel/vdsettingsptr
type VDSettingsPtr = *VDSettings

// See: https://developer.apple.com/documentation/kernel/vdsizeinfo
// VDSizeInfo is opaque storage with the size and alignment C gives VDSizeInfo:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type VDSizeInfo [4]uint16

// See: https://developer.apple.com/documentation/kernel/vdsupportshardwarecursorptr
type VDSupportsHardwareCursorPtr = *VDSupportsHardwareCursorRec

// See: https://developer.apple.com/documentation/kernel/vdsupportshardwarecursorrec
// VDSupportsHardwareCursorRec is opaque storage with the size and alignment C gives VDSupportsHardwareCursorRec:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type VDSupportsHardwareCursorRec [3]uint32

// See: https://developer.apple.com/documentation/kernel/vdswitchinfoptr
type VDSwitchInfoPtr = *VDSwitchInfoRec

// See: https://developer.apple.com/documentation/kernel/vdswitchinforec
// VDSwitchInfoRec is opaque storage with the size and alignment C gives VDSwitchInfoRec:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type VDSwitchInfoRec [4]uint64

// See: https://developer.apple.com/documentation/kernel/vdsyncinfoptr
type VDSyncInfoPtr = *VDSyncInfoRec

// See: https://developer.apple.com/documentation/kernel/vdsyncinforec
// VDSyncInfoRec is opaque storage with the size and alignment C gives VDSyncInfoRec:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type VDSyncInfoRec [2]byte

// See: https://developer.apple.com/documentation/kernel/vdszinfoptr
type VDSzInfoPtr = *VDSizeInfo

// See: https://developer.apple.com/documentation/kernel/vdtiminginfoptr
type VDTimingInfoPtr = *VDTimingInfoRec

// See: https://developer.apple.com/documentation/kernel/vdtiminginforec
// VDTimingInfoRec is opaque storage with the size and alignment C gives VDTimingInfoRec:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type VDTimingInfoRec [4]uint64

// See: https://developer.apple.com/documentation/kernel/vdvideoparametersinfoptr
type VDVideoParametersInfoPtr = *VDVideoParametersInfoRec

// See: https://developer.apple.com/documentation/kernel/vdvideoparametersinforec
// VDVideoParametersInfoRec is opaque storage with the size and alignment C gives VDVideoParametersInfoRec:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type VDVideoParametersInfoRec [4]uint64

// See: https://developer.apple.com/documentation/kernel/void
type VOID = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vpblock
// VPBlock is opaque storage with the size and alignment C gives VPBlock:
// 44 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 44 into.
type VPBlock [11]uint32

// See: https://developer.apple.com/documentation/kernel/vpblockptr
type VPBlockPtr = *VPBlock

// See: https://developer.apple.com/documentation/kernel/videodevicetype
type VideoDeviceType = uint32

// See: https://developer.apple.com/documentation/kernel/wk_word
type WK_word = uint32

// See: https://developer.apple.com/documentation/kernel/addr64_t
type Addr64_t = uint64

// See: https://developer.apple.com/documentation/kernel/aid_t
type Aid_t = uint64

// See: https://developer.apple.com/documentation/kernel/alarm_port_t
type Alarm_port_t = Alarm_t

// See: https://developer.apple.com/documentation/kernel/alarm_t
type Alarm_t = uint32

// See: https://developer.apple.com/documentation/kernel/alarm_type_t
type Alarm_type_t = int32

// See: https://developer.apple.com/documentation/kernel/arcade_register_t
type Arcade_register_t = uint32

// See: https://developer.apple.com/documentation/kernel/arm_debug_info_t
// Arm_debug_info_t is opaque storage with the size and alignment C gives arm_debug_info_t:
// 1 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 1 into.
type Arm_debug_info_t [1]byte

// See: https://developer.apple.com/documentation/kernel/arm_exception_state32_t
// Arm_exception_state32_t is opaque storage with the size and alignment C gives arm_exception_state32_t:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type Arm_exception_state32_t [3]uint32

// See: https://developer.apple.com/documentation/kernel/arm_feature_bits_t
// Arm_feature_bits_t is opaque storage with the size and alignment C gives arm_feature_bits_t:
// 1 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 1 into.
type Arm_feature_bits_t [1]byte

// See: https://developer.apple.com/documentation/kernel/arm_neon_state32_t
// Arm_neon_state32_t is an unresolved C aggregate typedef.
type Arm_neon_state32_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/arm_state_hdr_t
// Arm_state_hdr_t is opaque storage with the size and alignment C gives arm_state_hdr_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Arm_state_hdr_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/arm_thread_state32_t
// Arm_thread_state32_t is opaque storage with the size and alignment C gives arm_thread_state32_t:
// 68 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 68 into.
type Arm_thread_state32_t [17]uint32

// See: https://developer.apple.com/documentation/kernel/arm_unified_thread_state_t
// Arm_unified_thread_state_t is opaque storage with the size and alignment C gives arm_unified_thread_state_t:
// 280 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 280 into.
type Arm_unified_thread_state_t [35]uint64

// See: https://developer.apple.com/documentation/kernel/ataregisterimage
// AtaRegisterImage is opaque storage with the size and alignment C gives ataRegisterImage:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type AtaRegisterImage [6]uint16

// See: https://developer.apple.com/documentation/kernel/atataskfile
// AtaTaskFile is opaque storage with the size and alignment C gives ataTaskFile:
// 7 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 7 into.
type AtaTaskFile [7]byte

// See: https://developer.apple.com/documentation/kernel/atm_action_t
type Atm_action_t = uint32

// See: https://developer.apple.com/documentation/kernel/atm_aid_t
type Atm_aid_t = uint64

// See: https://developer.apple.com/documentation/kernel/atm_guard_t
type Atm_guard_t = uint64

// See: https://developer.apple.com/documentation/kernel/atm_mailbox_offset_t
type Atm_mailbox_offset_t = uint64

// See: https://developer.apple.com/documentation/kernel/atm_memory_descriptor_array_t
type Atm_memory_descriptor_array_t = *Atm_memory_descriptor_t

// See: https://developer.apple.com/documentation/kernel/atm_memory_descriptor_t
type Atm_memory_descriptor_t = uint32

// See: https://developer.apple.com/documentation/kernel/atm_memory_size_array_t
type Atm_memory_size_array_t = *uint64

// See: https://developer.apple.com/documentation/kernel/atm_subaid32_t
type Atm_subaid32_t = uint32

// See: https://developer.apple.com/documentation/kernel/attrgroup_t
type Attrgroup_t = uint32

// See: https://developer.apple.com/documentation/kernel/attribute_set_t
// Attribute_set_t is opaque storage with the size and alignment C gives attribute_set_t:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type Attribute_set_t [5]uint32

// See: https://developer.apple.com/documentation/kernel/attrreference_t
// Attrreference_t is opaque storage with the size and alignment C gives attrreference_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Attrreference_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/au_asflgs_t
type Au_asflgs_t = uint64

// See: https://developer.apple.com/documentation/kernel/au_asid_t
type Au_asid_t = int32

// See: https://developer.apple.com/documentation/kernel/au_class_t
type Au_class_t = uint32

// See: https://developer.apple.com/documentation/kernel/au_ctlmode_t
type Au_ctlmode_t = byte

// See: https://developer.apple.com/documentation/kernel/au_emod_t
type Au_emod_t = uint16

// See: https://developer.apple.com/documentation/kernel/au_evclass_map_t
// Au_evclass_map_t is opaque storage with the size and alignment C gives au_evclass_map_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Au_evclass_map_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/au_event_t
type Au_event_t = uint16

// See: https://developer.apple.com/documentation/kernel/au_expire_after_t
// Au_expire_after_t is opaque storage with the size and alignment C gives au_expire_after_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Au_expire_after_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/au_fstat_t
// Au_fstat_t is opaque storage with the size and alignment C gives au_fstat_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Au_fstat_t [2]uint64

// See: https://developer.apple.com/documentation/kernel/au_id_t
type Au_id_t = uint32

// See: https://developer.apple.com/documentation/kernel/au_mask_t
// Au_mask_t is opaque storage with the size and alignment C gives au_mask_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Au_mask_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/au_qctrl_t
// Au_qctrl_t is opaque storage with the size and alignment C gives au_qctrl_t:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type Au_qctrl_t [5]uint32

// See: https://developer.apple.com/documentation/kernel/au_session_t
// Au_session_t is opaque storage with the size and alignment C gives au_session_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Au_session_t [2]uint64

// See: https://developer.apple.com/documentation/kernel/au_stat_t
// Au_stat_t is opaque storage with the size and alignment C gives au_stat_t:
// 56 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 56 into.
type Au_stat_t [14]uint32

// See: https://developer.apple.com/documentation/kernel/au_tid_addr_t
// Au_tid_addr_t is opaque storage with the size and alignment C gives au_tid_addr_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Au_tid_addr_t [6]uint32

// See: https://developer.apple.com/documentation/kernel/au_tid_t
// Au_tid_t is opaque storage with the size and alignment C gives au_tid_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Au_tid_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/audit_token_t
type Audit_token_t = [32]byte

// See: https://developer.apple.com/documentation/kernel/auditinfo_addr_t
// Auditinfo_addr_t is opaque storage with the size and alignment C gives auditinfo_addr_t:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type Auditinfo_addr_t [6]uint64

// See: https://developer.apple.com/documentation/kernel/auditinfo_t
// Auditinfo_t is opaque storage with the size and alignment C gives auditinfo_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Auditinfo_t [6]uint32

// See: https://developer.apple.com/documentation/kernel/auditpinfo_addr_t
// Auditpinfo_addr_t is opaque storage with the size and alignment C gives auditpinfo_addr_t:
// 56 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 56 into.
type Auditpinfo_addr_t [7]uint64

// See: https://developer.apple.com/documentation/kernel/auditpinfo_t
// Auditpinfo_t is opaque storage with the size and alignment C gives auditpinfo_t:
// 28 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 28 into.
type Auditpinfo_t [7]uint32

// See: https://developer.apple.com/documentation/kernel/backtrace_flags_t
type Backtrace_flags_t = uint32

// See: https://developer.apple.com/documentation/kernel/backtrace_info_t
type Backtrace_info_t = uint32

// See: https://developer.apple.com/documentation/kernel/backtrace_pack_t
type Backtrace_pack_t = uint32

// See: https://developer.apple.com/documentation/kernel/bank_action_t
type Bank_action_t = uint32

// See: https://developer.apple.com/documentation/kernel/blkcnt_t
type Blkcnt_t = int64

// See: https://developer.apple.com/documentation/kernel/blksize_t
type Blksize_t = int32

// See: https://developer.apple.com/documentation/kernel/block_hint_t
type Block_hint_t = byte

// See: https://developer.apple.com/documentation/kernel/boolean_t
type Boolean_t = int32

// See: https://developer.apple.com/documentation/kernel/boot_args
// Boot_args is opaque storage with the size and alignment C gives boot_args:
// 4096 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4096 into.
type Boot_args [512]uint64

// See: https://developer.apple.com/documentation/kernel/boot_icon_element
// Boot_icon_element is opaque storage with the size and alignment C gives boot_icon_element:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type Boot_icon_element [8]uint32

// See: https://developer.apple.com/documentation/kernel/bootstrap_t
type Bootstrap_t = uint32

// See: https://developer.apple.com/documentation/kernel/bpf_int32
type Bpf_int32 = int32

// Bpf_tap_mode is mode for tapping. BPF_MODE_DISABLED/BPF_MODE_INPUT_OUTPUT etc.
//
// See: https://developer.apple.com/documentation/kernel/bpf_tap_mode
type Bpf_tap_mode = uint32

// See: https://developer.apple.com/documentation/kernel/bpf_u_int32
type Bpf_u_int32 = uint32

// See: https://developer.apple.com/documentation/kernel/buf_bptr_t
type Buf_bptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/buf_ptr_ref_t
type Buf_ptr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/buf_ptr_t
type Buf_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/buf_ref_ptr_t
type Buf_ref_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/buf_ref_ref_t
type Buf_ref_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/buf_ref_t
type Buf_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/buf_t
type Buf_t = uintptr

// See: https://developer.apple.com/documentation/kernel/bufattr_bptr_t
type Bufattr_bptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/bufattr_ptr_ref_t
type Bufattr_ptr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/bufattr_ptr_t
type Bufattr_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/bufattr_ref_ptr_t
type Bufattr_ref_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/bufattr_ref_ref_t
type Bufattr_ref_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/bufattr_ref_t
type Bufattr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/bufattr_t
type Bufattr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/cache_type_t
type Cache_type_t = uint32

// See: https://developer.apple.com/documentation/kernel/caddr_t
type Caddr_t = *byte

// See: https://developer.apple.com/documentation/kernel/caddr_ut
type Caddr_ut = Caddr_t

// See: https://developer.apple.com/documentation/kernel/call_gate_t
// Call_gate_t is opaque storage with the size and alignment C gives call_gate_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Call_gate_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/cc_t
type Cc_t = byte

// See: https://developer.apple.com/documentation/kernel/charf
type Charf = int8

// See: https://developer.apple.com/documentation/kernel/circle_queue_head_t
// Circle_queue_head_t is opaque storage with the size and alignment C gives circle_queue_head_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Circle_queue_head_t [1]uint64

// See: https://developer.apple.com/documentation/kernel/circle_queue_t
type Circle_queue_t = uintptr

// See: https://developer.apple.com/documentation/kernel/cl_direct_read_lock_t
// Cl_direct_read_lock_t is an unresolved C aggregate typedef.
type Cl_direct_read_lock_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/clock_attr_t
type Clock_attr_t = *int32

// See: https://developer.apple.com/documentation/kernel/clock_ctrl_port_t
type Clock_ctrl_port_t = Clock_ctrl_t

// See: https://developer.apple.com/documentation/kernel/clock_ctrl_t
type Clock_ctrl_t = uint32

// See: https://developer.apple.com/documentation/kernel/clock_flavor_t
type Clock_flavor_t = int32

// See: https://developer.apple.com/documentation/kernel/clock_frequency_info_t
// Clock_frequency_info_t is opaque storage with the size and alignment C gives clock_frequency_info_t:
// 200 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 200 into.
type Clock_frequency_info_t [25]uint64

// See: https://developer.apple.com/documentation/kernel/clock_id_t
type Clock_id_t = int32

// See: https://developer.apple.com/documentation/kernel/clock_nsec_t
type Clock_nsec_t = uint32

// See: https://developer.apple.com/documentation/kernel/clock_reply_t
type Clock_reply_t = uint32

// See: https://developer.apple.com/documentation/kernel/clock_res_t
type Clock_res_t = int32

// See: https://developer.apple.com/documentation/kernel/clock_sec_t
type Clock_sec_t = uint

// See: https://developer.apple.com/documentation/kernel/clock_serv_port_t
type Clock_serv_port_t = Clock_serv_t

// See: https://developer.apple.com/documentation/kernel/clock_serv_t
type Clock_serv_t = uint32

// See: https://developer.apple.com/documentation/kernel/clock_t
type Clock_t = uint

// See: https://developer.apple.com/documentation/kernel/clock_usec_t
type Clock_usec_t = uint32

// See: https://developer.apple.com/documentation/kernel/cluster_type_t
type Cluster_type_t = uint32

// See: https://developer.apple.com/documentation/kernel/coalition_t
type Coalition_t = uint32

// See: https://developer.apple.com/documentation/kernel/code_desc_t
// Code_desc_t is opaque storage with the size and alignment C gives code_desc_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Code_desc_t [4]uint16

// See: https://developer.apple.com/documentation/kernel/conninfo_multipathtcp_t
// Conninfo_multipathtcp_t is opaque storage with the size and alignment C gives conninfo_multipathtcp_t:
// 336 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 336 into.
type Conninfo_multipathtcp_t [42]uint64

// See: https://developer.apple.com/documentation/kernel/conninfo_tcp_t
// Conninfo_tcp_t is opaque storage with the size and alignment C gives conninfo_tcp_t:
// 424 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 424 into.
type Conninfo_tcp_t [106]uint32

// See: https://developer.apple.com/documentation/kernel/coprocessor_type_t
type Coprocessor_type_t = uint32

// See: https://developer.apple.com/documentation/kernel/cpu_id_t
type Cpu_id_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cpu_subtype_t
type Cpu_subtype_t = int32

// See: https://developer.apple.com/documentation/kernel/cpu_threadtype_t
type Cpu_threadtype_t = int32

// See: https://developer.apple.com/documentation/kernel/cpu_type_t
type Cpu_type_t = int32

// See: https://developer.apple.com/documentation/kernel/cpuid_arch_perf_leaf_t
// Cpuid_arch_perf_leaf_t is opaque storage with the size and alignment C gives cpuid_arch_perf_leaf_t:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type Cpuid_arch_perf_leaf_t [3]uint32

// See: https://developer.apple.com/documentation/kernel/cpuid_cache_desc_t
// Cpuid_cache_desc_t is opaque storage with the size and alignment C gives cpuid_cache_desc_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Cpuid_cache_desc_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/cpuid_mwait_leaf_t
// Cpuid_mwait_leaf_t is opaque storage with the size and alignment C gives cpuid_mwait_leaf_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Cpuid_mwait_leaf_t [4]uint32

// See: https://developer.apple.com/documentation/kernel/cpuid_register_t
type Cpuid_register_t = uint32

// See: https://developer.apple.com/documentation/kernel/cpuid_thermal_leaf_t
// Cpuid_thermal_leaf_t is opaque storage with the size and alignment C gives cpuid_thermal_leaf_t:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type Cpuid_thermal_leaf_t [10]uint32

// See: https://developer.apple.com/documentation/kernel/cpuid_tsc_leaf_t
// Cpuid_tsc_leaf_t is opaque storage with the size and alignment C gives cpuid_tsc_leaf_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Cpuid_tsc_leaf_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/cpuid_xsave_leaf_t
// Cpuid_xsave_leaf_t is opaque storage with the size and alignment C gives cpuid_xsave_leaf_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Cpuid_xsave_leaf_t [4]uint32

// See: https://developer.apple.com/documentation/kernel/cr0_t
// Cr0_t is opaque storage with the size and alignment C gives cr0_t:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type Cr0_t [1]uint32

// See: https://developer.apple.com/documentation/kernel/cryptex_auth_type_t
type Cryptex_auth_type_t = uint32

// See: https://developer.apple.com/documentation/kernel/crypto_random_ctx_t
type Crypto_random_ctx_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/crypto_random_kmem_ctx_size_fn_t
type Crypto_random_kmem_ctx_size_fn_t = uintptr

// See: https://developer.apple.com/documentation/kernel/cs_launch_type_t
type Cs_launch_type_t = CsLaunchType

// See: https://developer.apple.com/documentation/kernel/ct_rune_t
type Ct_rune_t = int32

// See: https://developer.apple.com/documentation/kernel/d_devtotty_t
type D_devtotty_t = *objc.ID

// See: https://developer.apple.com/documentation/kernel/daddr64_t
type Daddr64_t = int64

// See: https://developer.apple.com/documentation/kernel/daddr_t
type Daddr_t = int32

// See: https://developer.apple.com/documentation/kernel/data_desc_t
// Data_desc_t is opaque storage with the size and alignment C gives data_desc_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Data_desc_t [4]uint16

// See: https://developer.apple.com/documentation/kernel/debug_header_entry
type Debug_header_entry = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/debug_header_t
// Debug_header_t is opaque storage with the size and alignment C gives debug_header_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Debug_header_t [16]byte

// See: https://developer.apple.com/documentation/kernel/debug_trailer_t
// Debug_trailer_t is opaque storage with the size and alignment C gives debug_trailer_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Debug_trailer_t [2]uint64

// See: https://developer.apple.com/documentation/kernel/descriptor_options
type Descriptor_options = uint32

// See: https://developer.apple.com/documentation/kernel/dev_t
type Dev_t = int32

// See: https://developer.apple.com/documentation/kernel/dir_clone_authorizer_op_t
type Dir_clone_authorizer_op_t = uint32

// See: https://developer.apple.com/documentation/kernel/dk_bd_read_disc_info_t
// Dk_bd_read_disc_info_t is opaque storage with the size and alignment C gives dk_bd_read_disc_info_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_bd_read_disc_info_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/dk_bd_read_structure_t
// Dk_bd_read_structure_t is opaque storage with the size and alignment C gives dk_bd_read_structure_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_bd_read_structure_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/dk_bd_read_track_info_t
// Dk_bd_read_track_info_t is opaque storage with the size and alignment C gives dk_bd_read_track_info_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_bd_read_track_info_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/dk_bd_report_key_t
// Dk_bd_report_key_t is opaque storage with the size and alignment C gives dk_bd_report_key_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_bd_report_key_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/dk_bd_send_key_t
// Dk_bd_send_key_t is opaque storage with the size and alignment C gives dk_bd_send_key_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_bd_send_key_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/dk_cd_read_disc_info_t
// Dk_cd_read_disc_info_t is opaque storage with the size and alignment C gives dk_cd_read_disc_info_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_cd_read_disc_info_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/dk_cd_read_isrc_t
// Dk_cd_read_isrc_t is opaque storage with the size and alignment C gives dk_cd_read_isrc_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Dk_cd_read_isrc_t [16]byte

// See: https://developer.apple.com/documentation/kernel/dk_cd_read_mcn_t
// Dk_cd_read_mcn_t is opaque storage with the size and alignment C gives dk_cd_read_mcn_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Dk_cd_read_mcn_t [16]byte

// See: https://developer.apple.com/documentation/kernel/dk_cd_read_t
// Dk_cd_read_t is opaque storage with the size and alignment C gives dk_cd_read_t:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type Dk_cd_read_t [4]uint64

// See: https://developer.apple.com/documentation/kernel/dk_cd_read_toc_t
// Dk_cd_read_toc_t is opaque storage with the size and alignment C gives dk_cd_read_toc_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_cd_read_toc_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/dk_cd_read_track_info_t
// Dk_cd_read_track_info_t is opaque storage with the size and alignment C gives dk_cd_read_track_info_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_cd_read_track_info_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/dk_corestorage_info_t
// Dk_corestorage_info_t is opaque storage with the size and alignment C gives dk_corestorage_info_t:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type Dk_corestorage_info_t [8]uint64

// See: https://developer.apple.com/documentation/kernel/dk_dvd_read_disc_info_t
// Dk_dvd_read_disc_info_t is opaque storage with the size and alignment C gives dk_dvd_read_disc_info_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_dvd_read_disc_info_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/dk_dvd_read_rzone_info_t
// Dk_dvd_read_rzone_info_t is opaque storage with the size and alignment C gives dk_dvd_read_rzone_info_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_dvd_read_rzone_info_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/dk_dvd_read_structure_t
// Dk_dvd_read_structure_t is opaque storage with the size and alignment C gives dk_dvd_read_structure_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_dvd_read_structure_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/dk_dvd_report_key_t
// Dk_dvd_report_key_t is opaque storage with the size and alignment C gives dk_dvd_report_key_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_dvd_report_key_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/dk_dvd_send_key_t
// Dk_dvd_send_key_t is opaque storage with the size and alignment C gives dk_dvd_send_key_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_dvd_send_key_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/dk_error_description_t
// Dk_error_description_t is opaque storage with the size and alignment C gives dk_error_description_t:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type Dk_error_description_t [4]uint64

// See: https://developer.apple.com/documentation/kernel/dk_extent_t
// Dk_extent_t is opaque storage with the size and alignment C gives dk_extent_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Dk_extent_t [2]uint64

// See: https://developer.apple.com/documentation/kernel/dk_firmware_path_t
// Dk_firmware_path_t is opaque storage with the size and alignment C gives dk_firmware_path_t:
// 128 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 128 into.
type Dk_firmware_path_t [128]byte

// See: https://developer.apple.com/documentation/kernel/dk_format_capacities_t
// Dk_format_capacities_t is opaque storage with the size and alignment C gives dk_format_capacities_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Dk_format_capacities_t [2]uint64

// See: https://developer.apple.com/documentation/kernel/dk_format_capacity_t
// Dk_format_capacity_t is opaque storage with the size and alignment C gives dk_format_capacity_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Dk_format_capacity_t [2]uint64

// See: https://developer.apple.com/documentation/kernel/dk_physical_extent_t
// Dk_physical_extent_t is opaque storage with the size and alignment C gives dk_physical_extent_t:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type Dk_physical_extent_t [4]uint64

// See: https://developer.apple.com/documentation/kernel/dk_provision_extent_t
// Dk_provision_extent_t is opaque storage with the size and alignment C gives dk_provision_extent_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_provision_extent_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/dk_provision_status_t
// Dk_provision_status_t is opaque storage with the size and alignment C gives dk_provision_status_t:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type Dk_provision_status_t [5]uint64

// See: https://developer.apple.com/documentation/kernel/dk_set_tier_t
// Dk_set_tier_t is opaque storage with the size and alignment C gives dk_set_tier_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Dk_set_tier_t [2]uint64

// See: https://developer.apple.com/documentation/kernel/dk_synchronize_t
// Dk_synchronize_t is opaque storage with the size and alignment C gives dk_synchronize_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_synchronize_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/dk_unmap_t
// Dk_unmap_t is opaque storage with the size and alignment C gives dk_unmap_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Dk_unmap_t [2]uint64

// See: https://developer.apple.com/documentation/kernel/double_t
type Double_t = float64

// See: https://developer.apple.com/documentation/kernel/dump_fcn_t
type Dump_fcn_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dyld_kernel_image_info_array_t
type Dyld_kernel_image_info_array_t = *Dyld_kernel_image_info_t

// See: https://developer.apple.com/documentation/kernel/dyld_kernel_image_info_t
// Dyld_kernel_image_info_t is opaque storage with the size and alignment C gives dyld_kernel_image_info_t:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type Dyld_kernel_image_info_t [5]uint64

// See: https://developer.apple.com/documentation/kernel/dyld_kernel_process_info_t
// Dyld_kernel_process_info_t is opaque storage with the size and alignment C gives dyld_kernel_process_info_t:
// 72 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 72 into.
type Dyld_kernel_process_info_t [9]uint64

// See: https://developer.apple.com/documentation/kernel/eioaccelsurfacelockbits
type EIOAccelSurfaceLockBits = int

// See: https://developer.apple.com/documentation/kernel/eioaccelsurfacemodebits
type EIOAccelSurfaceModeBits = int

// See: https://developer.apple.com/documentation/kernel/eioaccelsurfacescalebits
type EIOAccelSurfaceScaleBits = int

// See: https://developer.apple.com/documentation/kernel/eioaccelsurfaceshapebits
type EIOAccelSurfaceShapeBits = int

// See: https://developer.apple.com/documentation/kernel/eioaccelsurfacestatebits
type EIOAccelSurfaceStateBits = int

// See: https://developer.apple.com/documentation/kernel/ecc_event_t
// Ecc_event_t is opaque storage with the size and alignment C gives ecc_event_t:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type Ecc_event_t [5]uint64

// See: https://developer.apple.com/documentation/kernel/ecc_flags_t
type Ecc_flags_t = uint32

// See: https://developer.apple.com/documentation/kernel/ecc_version_t
type Ecc_version_t = uint32

// See: https://developer.apple.com/documentation/kernel/empty_fcn_t
type Empty_fcn_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/emulation_vector_t
type Emulation_vector_t = *Mach_vm_offset_t

// See: https://developer.apple.com/documentation/kernel/eph_panic_flags_t
type Eph_panic_flags_t = uint64

// See: https://developer.apple.com/documentation/kernel/er_t
// Er_t is an unresolved C aggregate typedef.
type Er_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/errno_t
type Errno_t = int32

// See: https://developer.apple.com/documentation/kernel/ether_addr_t
type Ether_addr_t = [6]byte

// See: https://developer.apple.com/documentation/kernel/ether_header_t
// Ether_header_t is opaque storage with the size and alignment C gives ether_header_t:
// 14 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 14 into.
type Ether_header_t [7]uint16

// See: https://developer.apple.com/documentation/kernel/event64_t
type Event64_t = uint64

// See: https://developer.apple.com/documentation/kernel/event_t
type Event_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/eventlink_port_pair_t
type Eventlink_port_pair_t = uintptr

// See: https://developer.apple.com/documentation/kernel/eviospecialkeymsg_t
type EvioSpecialKeyMsg_t = uintptr

// See: https://developer.apple.com/documentation/kernel/evsioevsioccsindices
type EvsioEVSIOCCSIndices = int

// See: https://developer.apple.com/documentation/kernel/evsioevsioscsindices
type EvsioEVSIOSCSIndices = int

// See: https://developer.apple.com/documentation/kernel/ex_cb_action_t
type Ex_cb_action_t = uint32

// See: https://developer.apple.com/documentation/kernel/ex_cb_class_t
type Ex_cb_class_t = uint32

// See: https://developer.apple.com/documentation/kernel/ex_cb_state_t
// Ex_cb_state_t is opaque storage with the size and alignment C gives ex_cb_state_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Ex_cb_state_t [1]uint64

// See: https://developer.apple.com/documentation/kernel/exception_behavior_array_t
type Exception_behavior_array_t = *Exception_behavior_t

// See: https://developer.apple.com/documentation/kernel/exception_behavior_t
type Exception_behavior_t = int32

// See: https://developer.apple.com/documentation/kernel/exception_data_t
type Exception_data_t = *Exception_data_type_t

// See: https://developer.apple.com/documentation/kernel/exception_data_type_t
type Exception_data_type_t = int32

// See: https://developer.apple.com/documentation/kernel/exception_flavor_array_t
type Exception_flavor_array_t = *Thread_state_flavor_t

// See: https://developer.apple.com/documentation/kernel/exception_handler_array_t
type Exception_handler_array_t = *Exception_handler_t

// See: https://developer.apple.com/documentation/kernel/exception_handler_info_array_t
type Exception_handler_info_array_t = *Ipc_info_port_t

// See: https://developer.apple.com/documentation/kernel/exception_handler_info_t
type Exception_handler_info_t = Ipc_info_port_t

// See: https://developer.apple.com/documentation/kernel/exception_handler_t
type Exception_handler_t = uint32

// See: https://developer.apple.com/documentation/kernel/exception_mask_array_t
type Exception_mask_array_t = *Exception_mask_t

// See: https://developer.apple.com/documentation/kernel/exception_mask_t
type Exception_mask_t = uint32

// See: https://developer.apple.com/documentation/kernel/exception_port_arrary_t
type Exception_port_arrary_t = Exception_handler_array_t

// See: https://developer.apple.com/documentation/kernel/exception_port_array_t
type Exception_port_array_t = *uint32

// See: https://developer.apple.com/documentation/kernel/exception_port_info_array_t
type Exception_port_info_array_t = *Ipc_info_port_t

// See: https://developer.apple.com/documentation/kernel/exception_port_t
type Exception_port_t = uint32

// See: https://developer.apple.com/documentation/kernel/exception_type_t
type Exception_type_t = int32

// See: https://developer.apple.com/documentation/kernel/exclave_ecstackentry_addr_t
type Exclave_ecstackentry_addr_t = uint64

// See: https://developer.apple.com/documentation/kernel/ext_paniclog_create_options_t
type Ext_paniclog_create_options_t = uint32

// See: https://developer.apple.com/documentation/kernel/ext_paniclog_flags_t
type Ext_paniclog_flags_t = ExtPaniclogFlags

// See: https://developer.apple.com/documentation/kernel/extentrecord
// Extentrecord is opaque storage with the size and alignment C gives extentrecord:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type Extentrecord [16]uint32

// See: https://developer.apple.com/documentation/kernel/fattributiontag_t
// Fattributiontag_t is opaque storage with the size and alignment C gives fattributiontag_t:
// 272 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 272 into.
type Fattributiontag_t [34]uint64

// See: https://developer.apple.com/documentation/kernel/fchecklv_t
// Fchecklv_t is opaque storage with the size and alignment C gives fchecklv_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Fchecklv_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/fd_mask
type Fd_mask = int32

// See: https://developer.apple.com/documentation/kernel/fd_set
// Fd_set is opaque storage with the size and alignment C gives fd_set:
// 128 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 128 into.
type Fd_set [32]uint32

// See: https://developer.apple.com/documentation/kernel/fgetsigsinfo_t
// Fgetsigsinfo_t is opaque storage with the size and alignment C gives fgetsigsinfo_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Fgetsigsinfo_t [2]uint64

// See: https://developer.apple.com/documentation/kernel/fhandle_t
// Fhandle_t is opaque storage with the size and alignment C gives fhandle_t:
// 132 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 132 into.
type Fhandle_t [33]uint32

// See: https://developer.apple.com/documentation/kernel/file_bptr_t
type File_bptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/file_ptr_ref_t
type File_ptr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/file_ptr_t
type File_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/file_ref_ptr_t
type File_ref_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/file_ref_ref_t
type File_ref_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/file_ref_t
type File_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/file_t
type File_t = uintptr

// See: https://developer.apple.com/documentation/kernel/filesec_t
type Filesec_t = uintptr

// See: https://developer.apple.com/documentation/kernel/fixpt_t
type Fixpt_t = uint32

// See: https://developer.apple.com/documentation/kernel/float_t
type Float_t = float32

// See: https://developer.apple.com/documentation/kernel/fp_control_t
// Fp_control_t is an unresolved C aggregate typedef.
type Fp_control_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fp_status_t
// Fp_status_t is an unresolved C aggregate typedef.
type Fp_status_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fpunchhole_t
// Fpunchhole_t is opaque storage with the size and alignment C gives fpunchhole_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Fpunchhole_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/frame_type_bitmask_t
type Frame_type_bitmask_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fs_role_mount_args_t
// Fs_role_mount_args_t is opaque storage with the size and alignment C gives fs_role_mount_args_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Fs_role_mount_args_t [2]uint64

// See: https://developer.apple.com/documentation/kernel/fsblkcnt_t
type Fsblkcnt_t = uint32

// See: https://developer.apple.com/documentation/kernel/fsfilcnt_t
type Fsfilcnt_t = uint32

// See: https://developer.apple.com/documentation/kernel/fsfile_type_t
type Fsfile_type_t = uint32

// See: https://developer.apple.com/documentation/kernel/fsid_t
// Fsid_t is opaque storage with the size and alignment C gives fsid_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Fsid_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/fsignatures_t
// Fsignatures_t is opaque storage with the size and alignment C gives fsignatures_t:
// 56 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 56 into.
type Fsignatures_t [7]uint64

// See: https://developer.apple.com/documentation/kernel/fsobj_id_t
// Fsobj_id_t is opaque storage with the size and alignment C gives fsobj_id_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Fsobj_id_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/fsobj_tag_t
type Fsobj_tag_t = uint32

// See: https://developer.apple.com/documentation/kernel/fsobj_type_t
type Fsobj_type_t = uint32

// See: https://developer.apple.com/documentation/kernel/fspecread_t
// Fspecread_t is opaque storage with the size and alignment C gives fspecread_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Fspecread_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/fstore_t
// Fstore_t is opaque storage with the size and alignment C gives fstore_t:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type Fstore_t [4]uint64

// See: https://developer.apple.com/documentation/kernel/fsupplement_t
// Fsupplement_t is opaque storage with the size and alignment C gives fsupplement_t:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type Fsupplement_t [4]uint64

// See: https://developer.apple.com/documentation/kernel/fsvolid_t
type Fsvolid_t = uint32

// See: https://developer.apple.com/documentation/kernel/ftrimactivefile_t
// Ftrimactivefile_t is opaque storage with the size and alignment C gives ftrimactivefile_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Ftrimactivefile_t [2]uint64

// See: https://developer.apple.com/documentation/kernel/gdt_t
// Gdt_t is opaque storage with the size and alignment C gives gdt_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Gdt_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/gid_t
type Gid_t = uint32

// See: https://developer.apple.com/documentation/kernel/gpu_descriptor
// Gpu_descriptor is opaque storage with the size and alignment C gives gpu_descriptor:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Gpu_descriptor [2]uint32

// See: https://developer.apple.com/documentation/kernel/gpu_descriptor_t
type Gpu_descriptor_t = *Gpu_descriptor

// See: https://developer.apple.com/documentation/kernel/gpu_energy_data
// Gpu_energy_data is opaque storage with the size and alignment C gives gpu_energy_data:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type Gpu_energy_data [8]uint32

// See: https://developer.apple.com/documentation/kernel/gpu_energy_data_t
type Gpu_energy_data_t = *Gpu_energy_data

// See: https://developer.apple.com/documentation/kernel/graftdmg_type_t
type Graftdmg_type_t = uint32

// See: https://developer.apple.com/documentation/kernel/gssd_byte_buffer
type Gssd_byte_buffer = *uint8

// See: https://developer.apple.com/documentation/kernel/gssd_cred
type Gssd_cred = uint64

// See: https://developer.apple.com/documentation/kernel/gssd_ctx
type Gssd_ctx = uint64

// See: https://developer.apple.com/documentation/kernel/gssd_dstring
type Gssd_dstring = *byte

// See: https://developer.apple.com/documentation/kernel/gssd_etype_list
type Gssd_etype_list = *int32

// See: https://developer.apple.com/documentation/kernel/gssd_gid_list
type Gssd_gid_list = *uint32

// See: https://developer.apple.com/documentation/kernel/gssd_mechtype
type Gssd_mechtype = GssdMechtype

// See: https://developer.apple.com/documentation/kernel/gssd_nametype
type Gssd_nametype = GssdNametype

// See: https://developer.apple.com/documentation/kernel/gssd_string
type Gssd_string = *byte

// See: https://developer.apple.com/documentation/kernel/gz_header
// Gz_header is opaque storage with the size and alignment C gives gz_header:
// 80 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 80 into.
type Gz_header [10]uint64

// See: https://developer.apple.com/documentation/kernel/gz_headerp
type Gz_headerp = *Gz_header

// See: https://developer.apple.com/documentation/kernel/hash_info_bucket_array_t
type Hash_info_bucket_array_t = *Hash_info_bucket_t

// See: https://developer.apple.com/documentation/kernel/hash_info_bucket_t
// Hash_info_bucket_t is opaque storage with the size and alignment C gives hash_info_bucket_t:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type Hash_info_bucket_t [1]uint32

// See: https://developer.apple.com/documentation/kernel/host_basic_info_data_t
// Host_basic_info_data_t is opaque storage with the size and alignment C gives host_basic_info_data_t:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type Host_basic_info_data_t [12]uint32

// See: https://developer.apple.com/documentation/kernel/host_basic_info_t
type Host_basic_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/host_can_has_debugger_info_data_t
// Host_can_has_debugger_info_data_t is opaque storage with the size and alignment C gives host_can_has_debugger_info_data_t:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type Host_can_has_debugger_info_data_t [1]uint32

// See: https://developer.apple.com/documentation/kernel/host_can_has_debugger_info_t
type Host_can_has_debugger_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/host_cpu_load_info_data_t
// Host_cpu_load_info_data_t is opaque storage with the size and alignment C gives host_cpu_load_info_data_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Host_cpu_load_info_data_t [4]uint32

// See: https://developer.apple.com/documentation/kernel/host_cpu_load_info_t
type Host_cpu_load_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/host_flavor_t
type Host_flavor_t = int32

// See: https://developer.apple.com/documentation/kernel/host_info64_t
type Host_info64_t = *Integer_t

// See: https://developer.apple.com/documentation/kernel/host_info_data_t
type Host_info_data_t = int32

// See: https://developer.apple.com/documentation/kernel/host_info_t
type Host_info_t = *Integer_t

// See: https://developer.apple.com/documentation/kernel/host_load_info_data_t
// Host_load_info_data_t is opaque storage with the size and alignment C gives host_load_info_data_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Host_load_info_data_t [6]uint32

// See: https://developer.apple.com/documentation/kernel/host_load_info_t
type Host_load_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/host_name_port_t
type Host_name_port_t = Host_t

// See: https://developer.apple.com/documentation/kernel/host_name_t
type Host_name_t = Host_t

// See: https://developer.apple.com/documentation/kernel/host_preferred_user_arch_data_t
// Host_preferred_user_arch_data_t is opaque storage with the size and alignment C gives host_preferred_user_arch_data_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Host_preferred_user_arch_data_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/host_preferred_user_arch_t
type Host_preferred_user_arch_t = uintptr

// See: https://developer.apple.com/documentation/kernel/host_priority_info_data_t
// Host_priority_info_data_t is opaque storage with the size and alignment C gives host_priority_info_data_t:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type Host_priority_info_data_t [8]uint32

// See: https://developer.apple.com/documentation/kernel/host_priority_info_t
type Host_priority_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/host_priv_t
type Host_priv_t = uint32

// See: https://developer.apple.com/documentation/kernel/host_purgable_info_data_t
// Host_purgable_info_data_t is opaque storage with the size and alignment C gives host_purgable_info_data_t:
// 272 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 272 into.
type Host_purgable_info_data_t [34]uint64

// See: https://developer.apple.com/documentation/kernel/host_purgable_info_t
type Host_purgable_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/host_sched_info_data_t
// Host_sched_info_data_t is opaque storage with the size and alignment C gives host_sched_info_data_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Host_sched_info_data_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/host_sched_info_t
type Host_sched_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/host_security_t
type Host_security_t = uint32

// See: https://developer.apple.com/documentation/kernel/host_t
type Host_t = uint32

// See: https://developer.apple.com/documentation/kernel/hv_callbacks_t
// Hv_callbacks_t is opaque storage with the size and alignment C gives hv_callbacks_t:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type Hv_callbacks_t [8]uint64

// See: https://developer.apple.com/documentation/kernel/hv_trap_table_t
// Hv_trap_table_t is opaque storage with the size and alignment C gives hv_trap_table_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Hv_trap_table_t [2]uint64

// See: https://developer.apple.com/documentation/kernel/hv_trap_type_t
type Hv_trap_type_t = uint32

// See: https://developer.apple.com/documentation/kernel/hv_volatile_state_t
type Hv_volatile_state_t = uint32

// See: https://developer.apple.com/documentation/kernel/hvg_hcall_args_t
// Hvg_hcall_args_t is opaque storage with the size and alignment C gives hvg_hcall_args_t:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type Hvg_hcall_args_t [6]uint64

// See: https://developer.apple.com/documentation/kernel/hvg_hcall_code_t
type Hvg_hcall_code_t = HvgHcallCode

// See: https://developer.apple.com/documentation/kernel/hvg_hcall_dump_option_t
type Hvg_hcall_dump_option_t = HvgHcallDumpOption

// See: https://developer.apple.com/documentation/kernel/hvg_hcall_output_t
// Hvg_hcall_output_t is opaque storage with the size and alignment C gives hvg_hcall_output_t:
// 56 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 56 into.
type Hvg_hcall_output_t [7]uint64

// See: https://developer.apple.com/documentation/kernel/hvg_hcall_return_t
type Hvg_hcall_return_t = HvgHcallReturn

// See: https://developer.apple.com/documentation/kernel/hvg_hcall_vmcore_file_t
// Hvg_hcall_vmcore_file_t is opaque storage with the size and alignment C gives hvg_hcall_vmcore_file_t:
// 57 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 57 into.
type Hvg_hcall_vmcore_file_t [57]byte

// See: https://developer.apple.com/documentation/kernel/hw_spin_policy_t
type Hw_spin_policy_t = uintptr

// See: https://developer.apple.com/documentation/kernel/i386_cpu_info_t
// I386_cpu_info_t is opaque storage with the size and alignment C gives i386_cpu_info_t:
// 608 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 608 into.
type I386_cpu_info_t [76]uint64

// See: https://developer.apple.com/documentation/kernel/i386_ioport_t
type I386_ioport_t = uint16

// See: https://developer.apple.com/documentation/kernel/id_t
type Id_t = uint32

// See: https://developer.apple.com/documentation/kernel/idle_tickle_t
type Idle_tickle_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/idt_t
// Idt_t is opaque storage with the size and alignment C gives idt_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Idt_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/idtype_t
type Idtype_t = uint32

// See: https://developer.apple.com/documentation/kernel/if_clone_bptr_t
type If_clone_bptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/if_clone_ptr_ref_t
type If_clone_ptr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/if_clone_ptr_t
type If_clone_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/if_clone_ref_ptr_t
type If_clone_ref_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/if_clone_ref_ref_t
type If_clone_ref_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/if_clone_ref_t
type If_clone_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/if_clone_t
type If_clone_t = uintptr

// See: https://developer.apple.com/documentation/kernel/if_netem_model_t
type If_netem_model_t = uint32

// See: https://developer.apple.com/documentation/kernel/ifaddr_bptr_t
type Ifaddr_bptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ifaddr_ptr_ref_t
type Ifaddr_ptr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ifaddr_ptr_t
type Ifaddr_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ifaddr_ref_ptr_t
type Ifaddr_ref_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ifaddr_ref_ref_t
type Ifaddr_ref_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ifaddr_ref_t
type Ifaddr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ifaddr_t
type Ifaddr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ifmultiaddr_bptr_t
type Ifmultiaddr_bptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ifmultiaddr_ptr_ref_t
type Ifmultiaddr_ptr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ifmultiaddr_ptr_t
type Ifmultiaddr_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ifmultiaddr_ref_ptr_t
type Ifmultiaddr_ref_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ifmultiaddr_ref_ref_t
type Ifmultiaddr_ref_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ifmultiaddr_ref_t
type Ifmultiaddr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ifmultiaddr_t
type Ifmultiaddr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ifnet_bptr_t
type Ifnet_bptr_t = uintptr

// Ifnet_family_t is storage type for the interface family.
//
// See: https://developer.apple.com/documentation/kernel/ifnet_family_t
type Ifnet_family_t = uint32

// See: https://developer.apple.com/documentation/kernel/ifnet_filter_bptr_t
type Ifnet_filter_bptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ifnet_filter_ptr_ref_t
type Ifnet_filter_ptr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ifnet_filter_ptr_t
type Ifnet_filter_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ifnet_filter_ref_ptr_t
type Ifnet_filter_ref_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ifnet_filter_ref_ref_t
type Ifnet_filter_ref_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ifnet_filter_ref_t
type Ifnet_filter_ref_t = uintptr

// Ifnet_offload_t is flags indicating the offload support of the interface.
//
// See: https://developer.apple.com/documentation/kernel/ifnet_offload_t
type Ifnet_offload_t = uint32

// See: https://developer.apple.com/documentation/kernel/ifnet_ptr_ref_t
type Ifnet_ptr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ifnet_ptr_t
type Ifnet_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ifnet_ref_ptr_t
type Ifnet_ref_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ifnet_ref_ref_t
type Ifnet_ref_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ifnet_ref_t
type Ifnet_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ifnet_t
type Ifnet_t = uintptr

// See: https://developer.apple.com/documentation/kernel/in6_addr_t
// In6_addr_t is opaque storage with the size and alignment C gives in6_addr_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type In6_addr_t [4]uint32

// See: https://developer.apple.com/documentation/kernel/in6_clat46_evhdlr_code_t
type In6_clat46_evhdlr_code_t = uint32

// See: https://developer.apple.com/documentation/kernel/in_addr_t
type In_addr_t = uint32

// See: https://developer.apple.com/documentation/kernel/in_port_t
type In_port_t = uint16

// See: https://developer.apple.com/documentation/kernel/ino64_t
type Ino64_t = uint64

// See: https://developer.apple.com/documentation/kernel/ino_t
type Ino_t = uint64

// See: https://developer.apple.com/documentation/kernel/inp_gen_t
type Inp_gen_t = uint64

// See: https://developer.apple.com/documentation/kernel/int16_t
type Int16_t = int16

// See: https://developer.apple.com/documentation/kernel/int32_t
type Int32_t = int32

// See: https://developer.apple.com/documentation/kernel/int64_t
type Int64_t = int64

// See: https://developer.apple.com/documentation/kernel/int8_t
type Int8_t = int8

// See: https://developer.apple.com/documentation/kernel/int_fast16_t
type Int_fast16_t = int16

// See: https://developer.apple.com/documentation/kernel/int_fast32_t
type Int_fast32_t = int32

// See: https://developer.apple.com/documentation/kernel/int_fast64_t
type Int_fast64_t = int64

// See: https://developer.apple.com/documentation/kernel/int_fast8_t
type Int_fast8_t = int8

// See: https://developer.apple.com/documentation/kernel/int_least16_t
type Int_least16_t = int16

// See: https://developer.apple.com/documentation/kernel/int_least32_t
type Int_least32_t = int32

// See: https://developer.apple.com/documentation/kernel/int_least64_t
type Int_least64_t = int64

// See: https://developer.apple.com/documentation/kernel/int_least8_t
type Int_least8_t = int8

// See: https://developer.apple.com/documentation/kernel/integer_t
type Integer_t = int32

// See: https://developer.apple.com/documentation/kernel/interface_filter_t
type Interface_filter_t = uintptr

// See: https://developer.apple.com/documentation/kernel/intf
type Intf = int32

// See: https://developer.apple.com/documentation/kernel/intmax_t
type Intmax_t = int

// See: https://developer.apple.com/documentation/kernel/intptr_t
type Intptr_t = int

// See: https://developer.apple.com/documentation/kernel/intr_gate_t
// Intr_gate_t is opaque storage with the size and alignment C gives intr_gate_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Intr_gate_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/io_addr_t
type Io_addr_t = uint16

// See: https://developer.apple.com/documentation/kernel/io_async_ref64_t
type Io_async_ref64_t = uint64

// See: https://developer.apple.com/documentation/kernel/io_async_ref_t
type Io_async_ref_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/io_buf_ptr_t
type Io_buf_ptr_t = *byte

// See: https://developer.apple.com/documentation/kernel/io_compression_stats_t
// Io_compression_stats_t is an unresolved C aggregate typedef.
type Io_compression_stats_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/io_connect_t
type Io_connect_t = uint32

// See: https://developer.apple.com/documentation/kernel/io_enumerator_t
type Io_enumerator_t = uint32

// See: https://developer.apple.com/documentation/kernel/io_ident_t
type Io_ident_t = uint32

// See: https://developer.apple.com/documentation/kernel/io_iterator_t
type Io_iterator_t = uint32

// See: https://developer.apple.com/documentation/kernel/io_len_t
type Io_len_t = uint16

// See: https://developer.apple.com/documentation/kernel/io_main_t
type Io_main_t = uint32

// See: https://developer.apple.com/documentation/kernel/io_name_t
type Io_name_t = int8

// See: https://developer.apple.com/documentation/kernel/io_object_t
type Io_object_t = uint32

// See: https://developer.apple.com/documentation/kernel/io_registry_entry_t
type Io_registry_entry_t = uint32

// See: https://developer.apple.com/documentation/kernel/io_scalar_inband64_t
type Io_scalar_inband64_t = uint64

// See: https://developer.apple.com/documentation/kernel/io_scalar_inband_t
type Io_scalar_inband_t = int32

// See: https://developer.apple.com/documentation/kernel/io_service_t
type Io_service_t = uint32

// See: https://developer.apple.com/documentation/kernel/io_stat_info_t
// Io_stat_info_t is an unresolved C aggregate typedef.
type Io_stat_info_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/io_string_inband_t
type Io_string_inband_t = int8

// See: https://developer.apple.com/documentation/kernel/io_string_t
type Io_string_t = int8

// See: https://developer.apple.com/documentation/kernel/io_struct_inband_t
type Io_struct_inband_t = int8

// See: https://developer.apple.com/documentation/kernel/io_user_reference_t
type Io_user_reference_t = uint64

// See: https://developer.apple.com/documentation/kernel/io_user_scalar_t
type Io_user_scalar_t = uint64

// See: https://developer.apple.com/documentation/kernel/ioctl_fcn_t
type Ioctl_fcn_t = *objc.ID

// See: https://developer.apple.com/documentation/kernel/ipc_eventlink_t
type Ipc_eventlink_t = uint32

// See: https://developer.apple.com/documentation/kernel/ipc_info_name_array_t
type Ipc_info_name_array_t = *Ipc_info_name_t

// See: https://developer.apple.com/documentation/kernel/ipc_info_name_t
// Ipc_info_name_t is opaque storage with the size and alignment C gives ipc_info_name_t:
// 28 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 28 into.
type Ipc_info_name_t [7]uint32

// See: https://developer.apple.com/documentation/kernel/ipc_info_port_t
// Ipc_info_port_t is opaque storage with the size and alignment C gives ipc_info_port_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Ipc_info_port_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/ipc_info_space_basic_t
// Ipc_info_space_basic_t is opaque storage with the size and alignment C gives ipc_info_space_basic_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Ipc_info_space_basic_t [6]uint32

// See: https://developer.apple.com/documentation/kernel/ipc_info_space_t
// Ipc_info_space_t is opaque storage with the size and alignment C gives ipc_info_space_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Ipc_info_space_t [6]uint32

// See: https://developer.apple.com/documentation/kernel/ipc_info_tree_name_array_t
type Ipc_info_tree_name_array_t = *Ipc_info_tree_name_t

// See: https://developer.apple.com/documentation/kernel/ipc_info_tree_name_t
// Ipc_info_tree_name_t is opaque storage with the size and alignment C gives ipc_info_tree_name_t:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type Ipc_info_tree_name_t [9]uint32

// See: https://developer.apple.com/documentation/kernel/ipc_object_t
type Ipc_object_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ipc_port_t
type Ipc_port_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ipc_pthread_priority_value_t
type Ipc_pthread_priority_value_t = uint32

// See: https://developer.apple.com/documentation/kernel/ipc_space_inspect_t
type Ipc_space_inspect_t = uint32

// See: https://developer.apple.com/documentation/kernel/ipc_space_port_t
type Ipc_space_port_t = Ipc_space_t

// See: https://developer.apple.com/documentation/kernel/ipc_space_read_t
type Ipc_space_read_t = uint32

// See: https://developer.apple.com/documentation/kernel/ipc_space_t
type Ipc_space_t = uint32

// See: https://developer.apple.com/documentation/kernel/ipc_voucher_attr_control_t
type Ipc_voucher_attr_control_t = uint32

// See: https://developer.apple.com/documentation/kernel/ipc_voucher_attr_manager_t
type Ipc_voucher_attr_manager_t = uint32

// See: https://developer.apple.com/documentation/kernel/ipc_voucher_t
type Ipc_voucher_t = uint32

// See: https://developer.apple.com/documentation/kernel/ipi_handler_t
type Ipi_handler_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ipsec_dscp_mapping_t
type Ipsec_dscp_mapping_t = IpsecDscpMapping

// See: https://developer.apple.com/documentation/kernel/kusbconnectable
type KUSBConnectable = int

// See: https://developer.apple.com/documentation/kernel/kauth_ace_rights_t
type Kauth_ace_rights_t = uint32

// See: https://developer.apple.com/documentation/kernel/kauth_acl_eval_t
// Kauth_acl_eval_t is an unresolved C aggregate typedef.
type Kauth_acl_eval_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kauth_action_t
type Kauth_action_t = int32

// See: https://developer.apple.com/documentation/kernel/kauth_cred_t
// Kauth_cred_t is an unresolved C aggregate typedef.
type Kauth_cred_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kauth_listener_t
type Kauth_listener_t = uintptr

// See: https://developer.apple.com/documentation/kernel/kauth_scope_t
type Kauth_scope_t = uintptr

// See: https://developer.apple.com/documentation/kernel/kbdbitvector
type KbdBitVector = *uint32

// See: https://developer.apple.com/documentation/kernel/kbufinfo_t
// Kbufinfo_t is opaque storage with the size and alignment C gives kbufinfo_t:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type Kbufinfo_t [5]uint32

// See: https://developer.apple.com/documentation/kernel/kc_format_t
type Kc_format_t = uint32

// See: https://developer.apple.com/documentation/kernel/kc_kind_t
type Kc_kind_t = int32

// See: https://developer.apple.com/documentation/kernel/kcd_compression_type_t
type Kcd_compression_type_t = uint64

// See: https://developer.apple.com/documentation/kernel/kcdata_descriptor_t
type Kcdata_descriptor_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kcdata_item_t
type Kcdata_item_t = uintptr

// See: https://developer.apple.com/documentation/kernel/kcdata_iter_t
// Kcdata_iter_t is opaque storage with the size and alignment C gives kcdata_iter_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Kcdata_iter_t [2]uint64

// See: https://developer.apple.com/documentation/kernel/kcdata_object_t
type Kcdata_object_t = uint32

// See: https://developer.apple.com/documentation/kernel/kcdata_subtype_descriptor_t
// Kcdata_subtype_descriptor_t is an unresolved C aggregate typedef.
type Kcdata_subtype_descriptor_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kctype_subtype_t
type Kctype_subtype_t = uint32

// See: https://developer.apple.com/documentation/kernel/kd_buf
// Kd_buf is opaque storage with the size and alignment C gives kd_buf:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type Kd_buf [8]uint64

// See: https://developer.apple.com/documentation/kernel/kd_buf_argtype
type Kd_buf_argtype = uint64

// See: https://developer.apple.com/documentation/kernel/kd_callback_t
// Kd_callback_t is opaque storage with the size and alignment C gives kd_callback_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Kd_callback_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/kd_callback_type
type Kd_callback_type = uint32

// See: https://developer.apple.com/documentation/kernel/kd_cpumap
// Kd_cpumap is opaque storage with the size and alignment C gives kd_cpumap:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Kd_cpumap [4]uint32

// See: https://developer.apple.com/documentation/kernel/kd_cpumap_ext
// Kd_cpumap_ext is opaque storage with the size and alignment C gives kd_cpumap_ext:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type Kd_cpumap_ext [10]uint32

// See: https://developer.apple.com/documentation/kernel/kd_cpumap_header
// Kd_cpumap_header is opaque storage with the size and alignment C gives kd_cpumap_header:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Kd_cpumap_header [2]uint32

// See: https://developer.apple.com/documentation/kernel/kd_event_matcher
// Kd_event_matcher is opaque storage with the size and alignment C gives kd_event_matcher:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type Kd_event_matcher [5]uint64

// See: https://developer.apple.com/documentation/kernel/kd_regtype
// Kd_regtype is opaque storage with the size and alignment C gives kd_regtype:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type Kd_regtype [5]uint32

// See: https://developer.apple.com/documentation/kernel/kd_threadmap
// Kd_threadmap is opaque storage with the size and alignment C gives kd_threadmap:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type Kd_threadmap [4]uint64

// See: https://developer.apple.com/documentation/kernel/kdebug_coproc_flags_t
type Kdebug_coproc_flags_t = uint32

// See: https://developer.apple.com/documentation/kernel/kdebug_flags_t
type Kdebug_flags_t = KdebugFlags

// See: https://developer.apple.com/documentation/kernel/kdebug_live_flags_t
type Kdebug_live_flags_t = KdebugLiveFlags

// See: https://developer.apple.com/documentation/kernel/kdebug_test_t
type Kdebug_test_t = uint32

// See: https://developer.apple.com/documentation/kernel/kdp_event_t
type Kdp_event_t = KdpEvent

// See: https://developer.apple.com/documentation/kernel/kern_ctl_ref
type Kern_ctl_ref = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kern_return_t
type Kern_return_t = int32

// See: https://developer.apple.com/documentation/kernel/kernel_boot_info_t
type Kernel_boot_info_t = int8

// See: https://developer.apple.com/documentation/kernel/kernel_resource_sizes_data_t
// Kernel_resource_sizes_data_t is opaque storage with the size and alignment C gives kernel_resource_sizes_data_t:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type Kernel_resource_sizes_data_t [5]uint32

// See: https://developer.apple.com/documentation/kernel/kernel_resource_sizes_t
type Kernel_resource_sizes_t = uintptr

// See: https://developer.apple.com/documentation/kernel/kernel_version_t
type Kernel_version_t = int8

// See: https://developer.apple.com/documentation/kernel/key_t
type Key_t = int32

// See: https://developer.apple.com/documentation/kernel/kf_override_flag_t
type Kf_override_flag_t = uint32

// See: https://developer.apple.com/documentation/kernel/kmod_args_t
type Kmod_args_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kmod_control_flavor_t
type Kmod_control_flavor_t = int32

// See: https://developer.apple.com/documentation/kernel/kmod_info_32_v1_t
// Kmod_info_32_v1_t is opaque storage with the size and alignment C gives kmod_info_32_v1_t:
// 168 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 168 into.
type Kmod_info_32_v1_t [42]uint32

// See: https://developer.apple.com/documentation/kernel/kmod_info_64_v1_t
// Kmod_info_64_v1_t is opaque storage with the size and alignment C gives kmod_info_64_v1_t:
// 196 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 196 into.
type Kmod_info_64_v1_t [49]uint32

// See: https://developer.apple.com/documentation/kernel/kmod_info_array_t
type Kmod_info_array_t = *Kmod_info_t

// See: https://developer.apple.com/documentation/kernel/kmod_info_t
// Kmod_info_t is opaque storage with the size and alignment C gives kmod_info_t:
// 196 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 196 into.
type Kmod_info_t [49]uint32

// See: https://developer.apple.com/documentation/kernel/kmod_reference_t
// Kmod_reference_t is opaque storage with the size and alignment C gives kmod_reference_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Kmod_reference_t [4]uint32

// See: https://developer.apple.com/documentation/kernel/kmod_start_func_t
type Kmod_start_func_t = *objc.ID

// See: https://developer.apple.com/documentation/kernel/kmod_stop_func_t
type Kmod_stop_func_t = *objc.ID

// See: https://developer.apple.com/documentation/kernel/kmod_t
type Kmod_t = int32

// See: https://developer.apple.com/documentation/kernel/kobject_description_t
type Kobject_description_t = int8

// See: https://developer.apple.com/documentation/kernel/kpc_config_t
type Kpc_config_t = uint64

// See: https://developer.apple.com/documentation/kernel/kpc_pm_handler_t
type Kpc_pm_handler_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kperf_kpc_flags_t
type Kperf_kpc_flags_t = uint16

// See: https://developer.apple.com/documentation/kernel/labelstr_t
type Labelstr_t = *byte

// See: https://developer.apple.com/documentation/kernel/launch_constraint_data_t
// Launch_constraint_data_t is an unresolved C aggregate typedef.
type Launch_constraint_data_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/lck_attr_t
// Lck_attr_t is an unresolved C aggregate typedef.
type Lck_attr_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/lck_grp_attr_t
// Lck_grp_attr_t is an unresolved C aggregate typedef.
type Lck_grp_attr_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/lck_grp_t
// Lck_grp_t is an unresolved C aggregate typedef.
type Lck_grp_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/lck_mtx_ext_t
// Lck_mtx_ext_t is an unresolved C aggregate typedef.
type Lck_mtx_ext_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/lck_mtx_t
// Lck_mtx_t is an unresolved C aggregate typedef.
type Lck_mtx_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/lck_rw_t
// Lck_rw_t is an unresolved C aggregate typedef.
type Lck_rw_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/lck_rw_type_t
type Lck_rw_type_t = uint32

// See: https://developer.apple.com/documentation/kernel/lck_sleep_action_t
type Lck_sleep_action_t = uint32

// See: https://developer.apple.com/documentation/kernel/lck_spin_t
// Lck_spin_t is an unresolved C aggregate typedef.
type Lck_spin_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/lck_wake_action_t
type Lck_wake_action_t = uint32

// See: https://developer.apple.com/documentation/kernel/ldt_desc_t
// Ldt_desc_t is opaque storage with the size and alignment C gives ldt_desc_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Ldt_desc_t [4]uint16

// See: https://developer.apple.com/documentation/kernel/ldt_t
// Ldt_t is opaque storage with the size and alignment C gives ldt_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Ldt_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/ledger_amount_t
type Ledger_amount_t = int64

// See: https://developer.apple.com/documentation/kernel/ledger_array_t
type Ledger_array_t = *Ledger_t

// See: https://developer.apple.com/documentation/kernel/ledger_item_t
type Ledger_item_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/ledger_port_array_t
type Ledger_port_array_t = Ledger_array_t

// See: https://developer.apple.com/documentation/kernel/ledger_port_t
type Ledger_port_t = Ledger_t

// See: https://developer.apple.com/documentation/kernel/ledger_t
type Ledger_t = uint32

// See: https://developer.apple.com/documentation/kernel/libsptm_cpu_state_t
type Libsptm_cpu_state_t = uint8

// See: https://developer.apple.com/documentation/kernel/libsptm_error_t
type Libsptm_error_t = uint8

// See: https://developer.apple.com/documentation/kernel/libsptm_refcnt_type_t
type Libsptm_refcnt_type_t = uint8

// See: https://developer.apple.com/documentation/kernel/libsptm_state_t
// Libsptm_state_t is opaque storage with the size and alignment C gives libsptm_state_t:
// 312 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 312 into.
type Libsptm_state_t [39]uint64

// See: https://developer.apple.com/documentation/kernel/listxattrs_result_t
// Listxattrs_result_t is opaque storage with the size and alignment C gives listxattrs_result_t:
// 34920 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 34920 into.
type Listxattrs_result_t [4365]uint64

// See: https://developer.apple.com/documentation/kernel/lock_set_port_t
type Lock_set_port_t = Lock_set_t

// See: https://developer.apple.com/documentation/kernel/lock_set_t
type Lock_set_t = uint32

// See: https://developer.apple.com/documentation/kernel/lockgroup_info_array_t
type Lockgroup_info_array_t = *Lockgroup_info_t

// See: https://developer.apple.com/documentation/kernel/lockgroup_info_t
// Lockgroup_info_t is opaque storage with the size and alignment C gives lockgroup_info_t:
// 264 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 264 into.
type Lockgroup_info_t [33]uint64

// See: https://developer.apple.com/documentation/kernel/lz4_hash_entry_t
// Lz4_hash_entry_t is opaque storage with the size and alignment C gives lz4_hash_entry_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Lz4_hash_entry_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/mach_assert_type_t
type Mach_assert_type_t = byte

// See: https://developer.apple.com/documentation/kernel/mach_atm_subaid_t
type Mach_atm_subaid_t = uint64

// See: https://developer.apple.com/documentation/kernel/mach_bridge_regwrite_timestamp_func_t
type Mach_bridge_regwrite_timestamp_func_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_dead_name_notification_t
// Mach_dead_name_notification_t is opaque storage with the size and alignment C gives mach_dead_name_notification_t:
// 56 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 56 into.
type Mach_dead_name_notification_t [14]uint32

// See: https://developer.apple.com/documentation/kernel/mach_error_fn_t
type Mach_error_fn_t = int32

// See: https://developer.apple.com/documentation/kernel/mach_error_t
type Mach_error_t = int32

// See: https://developer.apple.com/documentation/kernel/mach_eventlink_t
type Mach_eventlink_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_exception_code_t
type Mach_exception_code_t = int64

// See: https://developer.apple.com/documentation/kernel/mach_exception_data_t
type Mach_exception_data_t = *Mach_exception_data_type_t

// See: https://developer.apple.com/documentation/kernel/mach_exception_data_type_t
type Mach_exception_data_type_t = int64

// See: https://developer.apple.com/documentation/kernel/mach_exception_subcode_t
type Mach_exception_subcode_t = int64

// See: https://developer.apple.com/documentation/kernel/mach_memory_info_array_t
type Mach_memory_info_array_t = *Mach_memory_info_t

// See: https://developer.apple.com/documentation/kernel/mach_memory_info_t
// Mach_memory_info_t is opaque storage with the size and alignment C gives mach_memory_info_t:
// 176 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 176 into.
type Mach_memory_info_t [22]uint64

// See: https://developer.apple.com/documentation/kernel/mach_msg_base_t
// Mach_msg_base_t is opaque storage with the size and alignment C gives mach_msg_base_t:
// 28 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 28 into.
type Mach_msg_base_t [7]uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_bits_t
type Mach_msg_bits_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_context_trailer_t
// Mach_msg_context_trailer_t is opaque storage with the size and alignment C gives mach_msg_context_trailer_t:
// 60 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 60 into.
type Mach_msg_context_trailer_t [15]uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_copy_options_t
type Mach_msg_copy_options_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_descriptor_type_t
type Mach_msg_descriptor_type_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_empty_rcv_t
// Mach_msg_empty_rcv_t is opaque storage with the size and alignment C gives mach_msg_empty_rcv_t:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type Mach_msg_empty_rcv_t [8]uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_empty_send_t
// Mach_msg_empty_send_t is opaque storage with the size and alignment C gives mach_msg_empty_send_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Mach_msg_empty_send_t [6]uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_filter_id
type Mach_msg_filter_id = int32

// See: https://developer.apple.com/documentation/kernel/mach_msg_format_0_trailer_t
type Mach_msg_format_0_trailer_t = Mach_msg_security_trailer_t

// See: https://developer.apple.com/documentation/kernel/mach_msg_guard_flags_t
type Mach_msg_guard_flags_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_guarded_port_descriptor32_t
// Mach_msg_guarded_port_descriptor32_t is opaque storage with the size and alignment C gives mach_msg_guarded_port_descriptor32_t:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type Mach_msg_guarded_port_descriptor32_t [3]uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_guarded_port_descriptor64_t
// Mach_msg_guarded_port_descriptor64_t is opaque storage with the size and alignment C gives mach_msg_guarded_port_descriptor64_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Mach_msg_guarded_port_descriptor64_t [4]uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_guarded_port_descriptor_t
// Mach_msg_guarded_port_descriptor_t is opaque storage with the size and alignment C gives mach_msg_guarded_port_descriptor_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Mach_msg_guarded_port_descriptor_t [4]uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_header_t
// Mach_msg_header_t is opaque storage with the size and alignment C gives mach_msg_header_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Mach_msg_header_t [6]uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_id_t
type Mach_msg_id_t = int32

// See: https://developer.apple.com/documentation/kernel/mach_msg_mac_trailer_t
// Mach_msg_mac_trailer_t is opaque storage with the size and alignment C gives mach_msg_mac_trailer_t:
// 68 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 68 into.
type Mach_msg_mac_trailer_t [17]uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_max_trailer_t
type Mach_msg_max_trailer_t = Mach_msg_mac_trailer_t

// See: https://developer.apple.com/documentation/kernel/mach_msg_ool_descriptor32_t
// Mach_msg_ool_descriptor32_t is opaque storage with the size and alignment C gives mach_msg_ool_descriptor32_t:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type Mach_msg_ool_descriptor32_t [3]uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_ool_descriptor64_t
// Mach_msg_ool_descriptor64_t is opaque storage with the size and alignment C gives mach_msg_ool_descriptor64_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Mach_msg_ool_descriptor64_t [4]uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_ool_ports_descriptor32_t
// Mach_msg_ool_ports_descriptor32_t is opaque storage with the size and alignment C gives mach_msg_ool_ports_descriptor32_t:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type Mach_msg_ool_ports_descriptor32_t [3]uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_ool_ports_descriptor64_t
// Mach_msg_ool_ports_descriptor64_t is opaque storage with the size and alignment C gives mach_msg_ool_ports_descriptor64_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Mach_msg_ool_ports_descriptor64_t [4]uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_ool_ports_descriptor_t
// Mach_msg_ool_ports_descriptor_t is opaque storage with the size and alignment C gives mach_msg_ool_ports_descriptor_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Mach_msg_ool_ports_descriptor_t [4]uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_option_t
type Mach_msg_option_t = int32

// See: https://developer.apple.com/documentation/kernel/mach_msg_options_t
type Mach_msg_options_t = int32

// See: https://developer.apple.com/documentation/kernel/mach_msg_priority_t
type Mach_msg_priority_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_return_t
type Mach_msg_return_t = int32

// See: https://developer.apple.com/documentation/kernel/mach_msg_security_trailer_t
// Mach_msg_security_trailer_t is opaque storage with the size and alignment C gives mach_msg_security_trailer_t:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type Mach_msg_security_trailer_t [5]uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_seqno_trailer_t
// Mach_msg_seqno_trailer_t is opaque storage with the size and alignment C gives mach_msg_seqno_trailer_t:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type Mach_msg_seqno_trailer_t [3]uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_size_t
type Mach_msg_size_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_timeout_t
type Mach_msg_timeout_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/mach_msg_trailer_info_t
type Mach_msg_trailer_info_t = *byte

// See: https://developer.apple.com/documentation/kernel/mach_msg_trailer_size_t
type Mach_msg_trailer_size_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_trailer_type_t
type Mach_msg_trailer_type_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_type_name_t
type Mach_msg_type_name_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_type_number_t
type Mach_msg_type_number_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/mach_msg_type_size_t
type Mach_msg_type_size_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/mach_no_senders_notification_t
// Mach_no_senders_notification_t is opaque storage with the size and alignment C gives mach_no_senders_notification_t:
// 56 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 56 into.
type Mach_no_senders_notification_t [14]uint32

// See: https://developer.apple.com/documentation/kernel/mach_port_array_t
type Mach_port_array_t = *uint32

// See: https://developer.apple.com/documentation/kernel/mach_port_context_t
type Mach_port_context_t = uint64

// See: https://developer.apple.com/documentation/kernel/mach_port_deleted_notification_t
// Mach_port_deleted_notification_t is opaque storage with the size and alignment C gives mach_port_deleted_notification_t:
// 56 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 56 into.
type Mach_port_deleted_notification_t [14]uint32

// See: https://developer.apple.com/documentation/kernel/mach_port_delta_t
type Mach_port_delta_t = int32

// See: https://developer.apple.com/documentation/kernel/mach_port_destroyed_notification_t
// Mach_port_destroyed_notification_t is opaque storage with the size and alignment C gives mach_port_destroyed_notification_t:
// 60 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 60 into.
type Mach_port_destroyed_notification_t [15]uint32

// See: https://developer.apple.com/documentation/kernel/mach_port_flavor_t
type Mach_port_flavor_t = int32

// See: https://developer.apple.com/documentation/kernel/mach_port_guard_info_t
// Mach_port_guard_info_t is opaque storage with the size and alignment C gives mach_port_guard_info_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Mach_port_guard_info_t [1]uint64

// See: https://developer.apple.com/documentation/kernel/mach_port_info_ext_t
// Mach_port_info_ext_t is opaque storage with the size and alignment C gives mach_port_info_ext_t:
// 68 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 68 into.
type Mach_port_info_ext_t [17]uint32

// See: https://developer.apple.com/documentation/kernel/mach_port_info_t
type Mach_port_info_t = *Integer_t

// See: https://developer.apple.com/documentation/kernel/mach_port_limits_t
// Mach_port_limits_t is opaque storage with the size and alignment C gives mach_port_limits_t:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type Mach_port_limits_t [1]uint32

// See: https://developer.apple.com/documentation/kernel/mach_port_mscount_t
type Mach_port_mscount_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/mach_port_msgcount_t
type Mach_port_msgcount_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/mach_port_name_array_t
type Mach_port_name_array_t = *Mach_port_name_t

// See: https://developer.apple.com/documentation/kernel/mach_port_name_t
type Mach_port_name_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/mach_port_options_ptr_t
type Mach_port_options_ptr_t = *Mach_port_options_t

// See: https://developer.apple.com/documentation/kernel/mach_port_options_t
// Mach_port_options_t is opaque storage with the size and alignment C gives mach_port_options_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Mach_port_options_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/mach_port_qos_t
// Mach_port_qos_t is opaque storage with the size and alignment C gives mach_port_qos_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Mach_port_qos_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/mach_port_right_t
type Mach_port_right_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/mach_port_rights_t
type Mach_port_rights_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/mach_port_seqno_t
type Mach_port_seqno_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/mach_port_srights_t
type Mach_port_srights_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_port_status_t
// Mach_port_status_t is opaque storage with the size and alignment C gives mach_port_status_t:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type Mach_port_status_t [10]uint32

// See: https://developer.apple.com/documentation/kernel/mach_port_t
type Mach_port_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_port_type_array_t
type Mach_port_type_array_t = *Mach_port_type_t

// See: https://developer.apple.com/documentation/kernel/mach_port_type_t
type Mach_port_type_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/mach_port_urefs_t
type Mach_port_urefs_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/mach_send_once_notification_t
// Mach_send_once_notification_t is opaque storage with the size and alignment C gives mach_send_once_notification_t:
// 44 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 44 into.
type Mach_send_once_notification_t [11]uint32

// See: https://developer.apple.com/documentation/kernel/mach_send_possible_notification_t
// Mach_send_possible_notification_t is opaque storage with the size and alignment C gives mach_send_possible_notification_t:
// 56 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 56 into.
type Mach_send_possible_notification_t [14]uint32

// See: https://developer.apple.com/documentation/kernel/mach_service_port_info_data_t
// Mach_service_port_info_data_t is opaque storage with the size and alignment C gives mach_service_port_info_data_t:
// 256 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 256 into.
type Mach_service_port_info_data_t [256]byte

// See: https://developer.apple.com/documentation/kernel/mach_service_port_info_t
type Mach_service_port_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/mach_task_basic_info_data_t
// Mach_task_basic_info_data_t is opaque storage with the size and alignment C gives mach_task_basic_info_data_t:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type Mach_task_basic_info_data_t [12]uint32

// See: https://developer.apple.com/documentation/kernel/mach_task_basic_info_t
type Mach_task_basic_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/mach_task_flavor_t
type Mach_task_flavor_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_thread_flavor_t
type Mach_thread_flavor_t = uint32

// Mach_timebase_info_data_t is raw Mach Time API In general prefer to use the API clock_gettime_nsec_np(3), which deals in the same clocks (and more) in ns units. Conversion of ns to (resp. from) tick units as returned by the mach time APIs is performed by division (resp. multiplication) with the fraction returned by mach_timebase_info().
//
// See: https://developer.apple.com/documentation/kernel/mach_timebase_info_data_t
// Mach_timebase_info_data_t is opaque storage with the size and alignment C gives mach_timebase_info_data_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Mach_timebase_info_data_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/mach_timespec_t
// Mach_timespec_t is opaque storage with the size and alignment C gives mach_timespec_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Mach_timespec_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/mach_vm_address_t
type Mach_vm_address_t = uint64

// See: https://developer.apple.com/documentation/kernel/mach_vm_address_ut
type Mach_vm_address_ut = uint64

// See: https://developer.apple.com/documentation/kernel/mach_vm_info_region_t
// Mach_vm_info_region_t is opaque storage with the size and alignment C gives mach_vm_info_region_t:
// 56 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 56 into.
type Mach_vm_info_region_t [14]uint32

// See: https://developer.apple.com/documentation/kernel/mach_vm_offset_t
type Mach_vm_offset_t = uint64

// See: https://developer.apple.com/documentation/kernel/mach_vm_offset_ut
type Mach_vm_offset_ut = uint64

// See: https://developer.apple.com/documentation/kernel/mach_vm_range_flags_t
type Mach_vm_range_flags_t = MachVmRangeFlags

// See: https://developer.apple.com/documentation/kernel/mach_vm_range_flavor_t
type Mach_vm_range_flavor_t = MachVmRangeFlavor

// See: https://developer.apple.com/documentation/kernel/mach_vm_range_recipe_t
type Mach_vm_range_recipe_t = Mach_vm_range_recipe_v1_t

// See: https://developer.apple.com/documentation/kernel/mach_vm_range_recipe_v1_t
// Mach_vm_range_recipe_v1_t is opaque storage with the size and alignment C gives mach_vm_range_recipe_v1_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Mach_vm_range_recipe_v1_t [24]byte

// See: https://developer.apple.com/documentation/kernel/mach_vm_range_recipe_v1_ut
type Mach_vm_range_recipe_v1_ut = Mach_vm_range_recipe_v1_t

// See: https://developer.apple.com/documentation/kernel/mach_vm_range_recipes_raw_t
type Mach_vm_range_recipes_raw_t = *uint8

// See: https://developer.apple.com/documentation/kernel/mach_vm_range_t
type Mach_vm_range_t = uintptr

// See: https://developer.apple.com/documentation/kernel/mach_vm_range_tag_t
type Mach_vm_range_tag_t = MachVmRangeTag

// See: https://developer.apple.com/documentation/kernel/mach_vm_read_entry_t
// Mach_vm_read_entry_t is opaque storage with the size and alignment C gives mach_vm_read_entry_t:
// 4096 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4096 into.
type Mach_vm_read_entry_t [1024]uint32

// See: https://developer.apple.com/documentation/kernel/mach_vm_size_t
type Mach_vm_size_t = uint64

// See: https://developer.apple.com/documentation/kernel/mach_vm_size_ut
type Mach_vm_size_ut = uint64

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_command_t
type Mach_voucher_attr_command_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_content_size_t
type Mach_voucher_attr_content_size_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_content_t
type Mach_voucher_attr_content_t = *uint8

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_control_flags_t
type Mach_voucher_attr_control_flags_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_control_t
type Mach_voucher_attr_control_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_importance_refs
type Mach_voucher_attr_importance_refs = uint32

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_key_array_t
type Mach_voucher_attr_key_array_t = *Mach_voucher_attr_key_t

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_key_t
type Mach_voucher_attr_key_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_manager_t
type Mach_voucher_attr_manager_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_raw_recipe_array_size_t
type Mach_voucher_attr_raw_recipe_array_size_t = Mach_msg_type_number_t

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_raw_recipe_array_t
type Mach_voucher_attr_raw_recipe_array_t = Mach_voucher_attr_raw_recipe_t

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_raw_recipe_size_t
type Mach_voucher_attr_raw_recipe_size_t = Mach_msg_type_number_t

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_raw_recipe_t
type Mach_voucher_attr_raw_recipe_t = *uint8

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_recipe_command_array_t
type Mach_voucher_attr_recipe_command_array_t = *Mach_voucher_attr_recipe_command_t

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_recipe_command_t
type Mach_voucher_attr_recipe_command_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_recipe_data_t
// Mach_voucher_attr_recipe_data_t is opaque storage with the size and alignment C gives mach_voucher_attr_recipe_data_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Mach_voucher_attr_recipe_data_t [16]byte

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_recipe_size_t
type Mach_voucher_attr_recipe_size_t = Mach_msg_type_number_t

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_recipe_t
type Mach_voucher_attr_recipe_t = *Mach_voucher_attr_recipe_data_t

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_value_flags_t
type Mach_voucher_attr_value_flags_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_value_handle_array_size_t
type Mach_voucher_attr_value_handle_array_size_t = Mach_msg_type_number_t

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_value_handle_array_t
type Mach_voucher_attr_value_handle_array_t = *Mach_voucher_attr_value_handle_t

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_value_handle_t
type Mach_voucher_attr_value_handle_t = uint64

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_value_reference_t
type Mach_voucher_attr_value_reference_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_voucher_name_array_t
type Mach_voucher_name_array_t = *Mach_voucher_name_t

// See: https://developer.apple.com/documentation/kernel/mach_voucher_name_t
type Mach_voucher_name_t = Mach_port_name_t

// See: https://developer.apple.com/documentation/kernel/mach_voucher_selector_t
type Mach_voucher_selector_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_voucher_t
type Mach_voucher_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_zone_info_array_t
type Mach_zone_info_array_t = *Mach_zone_info_t

// See: https://developer.apple.com/documentation/kernel/mach_zone_info_t
// Mach_zone_info_t is opaque storage with the size and alignment C gives mach_zone_info_t:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type Mach_zone_info_t [8]uint64

// See: https://developer.apple.com/documentation/kernel/mach_zone_name_array_t
type Mach_zone_name_array_t = *Mach_zone_name_t

// See: https://developer.apple.com/documentation/kernel/mach_zone_name_t
// Mach_zone_name_t is opaque storage with the size and alignment C gives mach_zone_name_t:
// 80 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 80 into.
type Mach_zone_name_t [80]byte

// See: https://developer.apple.com/documentation/kernel/mailbox_offset_t
type Mailbox_offset_t = uint64

// See: https://developer.apple.com/documentation/kernel/mb_class_stat_t
// Mb_class_stat_t is opaque storage with the size and alignment C gives mb_class_stat_t:
// 136 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 136 into.
type Mb_class_stat_t [17]uint64

// See: https://developer.apple.com/documentation/kernel/mb_stat_t
// Mb_stat_t is opaque storage with the size and alignment C gives mb_stat_t:
// 144 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 144 into.
type Mb_stat_t [18]uint64

// See: https://developer.apple.com/documentation/kernel/mbstate_t
// Mbstate_t is opaque storage with the size and alignment C gives mbstate_t:
// 128 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 128 into.
type Mbstate_t [16]uint64

// See: https://developer.apple.com/documentation/kernel/mbuf_bptr_t
type Mbuf_bptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/mbuf_csum_performed_flags_t
type Mbuf_csum_performed_flags_t = uint32

// See: https://developer.apple.com/documentation/kernel/mbuf_csum_request_flags_t
type Mbuf_csum_request_flags_t = uint32

// See: https://developer.apple.com/documentation/kernel/mbuf_flags_t
type Mbuf_flags_t = uint32

// See: https://developer.apple.com/documentation/kernel/mbuf_how_t
type Mbuf_how_t = uint32

// See: https://developer.apple.com/documentation/kernel/mbuf_ptr_ref_t
type Mbuf_ptr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/mbuf_ptr_t
type Mbuf_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/mbuf_ref_ptr_t
type Mbuf_ref_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/mbuf_ref_ref_t
type Mbuf_ref_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/mbuf_ref_t
type Mbuf_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/mbuf_t
type Mbuf_t = uintptr

// See: https://developer.apple.com/documentation/kernel/mbuf_tag_id_t
type Mbuf_tag_id_t = uint32

// See: https://developer.apple.com/documentation/kernel/mbuf_tag_type_t
type Mbuf_tag_type_t = uint16

// Mbuf_traffic_class_t is traffic class of a packet.
//
// See: https://developer.apple.com/documentation/kernel/mbuf_traffic_class_t
type Mbuf_traffic_class_t = uint32

// See: https://developer.apple.com/documentation/kernel/mbuf_tso_request_flags_t
type Mbuf_tso_request_flags_t = uint32

// See: https://developer.apple.com/documentation/kernel/mbuf_type_t
type Mbuf_type_t = uint32

// See: https://developer.apple.com/documentation/kernel/mcc_ecc_event_t
// Mcc_ecc_event_t is opaque storage with the size and alignment C gives mcc_ecc_event_t:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type Mcc_ecc_event_t [10]uint32

// See: https://developer.apple.com/documentation/kernel/mcc_ecc_version_t
type Mcc_ecc_version_t = uint32

// See: https://developer.apple.com/documentation/kernel/mcc_flags_t
type Mcc_flags_t = uint32

// See: https://developer.apple.com/documentation/kernel/mcontext_t
// Mcontext_t is an unresolved C aggregate typedef.
type Mcontext_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mem_entry_name_port_t
type Mem_entry_name_port_t = uint32

// See: https://developer.apple.com/documentation/kernel/memory_object_array_t
type Memory_object_array_t = *Memory_object_t

// See: https://developer.apple.com/documentation/kernel/memory_object_attr_info_data_t
// Memory_object_attr_info_data_t is opaque storage with the size and alignment C gives memory_object_attr_info_data_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Memory_object_attr_info_data_t [4]uint32

// See: https://developer.apple.com/documentation/kernel/memory_object_attr_info_t
type Memory_object_attr_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/memory_object_behave_info_data_t
// Memory_object_behave_info_data_t is opaque storage with the size and alignment C gives memory_object_behave_info_data_t:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type Memory_object_behave_info_data_t [5]uint32

// See: https://developer.apple.com/documentation/kernel/memory_object_behave_info_t
type Memory_object_behave_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/memory_object_cluster_size_t
type Memory_object_cluster_size_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/memory_object_control_t
type Memory_object_control_t = uint32

// See: https://developer.apple.com/documentation/kernel/memory_object_copy_strategy_t
type Memory_object_copy_strategy_t = int32

// See: https://developer.apple.com/documentation/kernel/memory_object_default_t
type Memory_object_default_t = uint32

// See: https://developer.apple.com/documentation/kernel/memory_object_fault_info_t
type Memory_object_fault_info_t = *Natural_t

// See: https://developer.apple.com/documentation/kernel/memory_object_flavor_t
type Memory_object_flavor_t = int32

// See: https://developer.apple.com/documentation/kernel/memory_object_info_data_t
type Memory_object_info_data_t = int32

// See: https://developer.apple.com/documentation/kernel/memory_object_info_t
type Memory_object_info_t = *int32

// See: https://developer.apple.com/documentation/kernel/memory_object_name_t
type Memory_object_name_t = uint32

// See: https://developer.apple.com/documentation/kernel/memory_object_offset_t
type Memory_object_offset_t = uint64

// See: https://developer.apple.com/documentation/kernel/memory_object_offset_ut
type Memory_object_offset_ut = uint64

// See: https://developer.apple.com/documentation/kernel/memory_object_perf_info_data_t
// Memory_object_perf_info_data_t is opaque storage with the size and alignment C gives memory_object_perf_info_data_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Memory_object_perf_info_data_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/memory_object_perf_info_t
type Memory_object_perf_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/memory_object_return_t
type Memory_object_return_t = int32

// See: https://developer.apple.com/documentation/kernel/memory_object_size_t
type Memory_object_size_t = uint64

// See: https://developer.apple.com/documentation/kernel/memory_object_size_ut
type Memory_object_size_ut = uint64

// See: https://developer.apple.com/documentation/kernel/memory_object_t
type Memory_object_t = uint32

// See: https://developer.apple.com/documentation/kernel/microstackshot_flags_t
type Microstackshot_flags_t = MicrostackshotFlags

// See: https://developer.apple.com/documentation/kernel/mig_impl_routine_t
type Mig_impl_routine_t = int32

// See: https://developer.apple.com/documentation/kernel/mig_reply_error_t
// Mig_reply_error_t is opaque storage with the size and alignment C gives mig_reply_error_t:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type Mig_reply_error_t [9]uint32

// See: https://developer.apple.com/documentation/kernel/mig_routine_arg_descriptor_t
type Mig_routine_arg_descriptor_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mig_routine_descriptor
// Mig_routine_descriptor is opaque storage with the size and alignment C gives mig_routine_descriptor:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type Mig_routine_descriptor [5]uint64

// See: https://developer.apple.com/documentation/kernel/mig_routine_descriptor_t
type Mig_routine_descriptor_t = *Mig_routine_descriptor

// See: https://developer.apple.com/documentation/kernel/mig_routine_t
type Mig_routine_t = uintptr

// See: https://developer.apple.com/documentation/kernel/mig_subsystem_t
// Mig_subsystem_t is an unresolved C aggregate typedef.
type Mig_subsystem_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mig_symtab_t
// Mig_symtab_t is opaque storage with the size and alignment C gives mig_symtab_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Mig_symtab_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/ml_cpu_info_t
// Ml_cpu_info_t is opaque storage with the size and alignment C gives ml_cpu_info_t:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type Ml_cpu_info_t [8]uint64

// See: https://developer.apple.com/documentation/kernel/ml_page_protection_t
type Ml_page_protection_t = int32

// See: https://developer.apple.com/documentation/kernel/ml_processor_info_t
// Ml_processor_info_t is opaque storage with the size and alignment C gives ml_processor_info_t:
// 144 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 144 into.
type Ml_processor_info_t [18]uint64

// See: https://developer.apple.com/documentation/kernel/mmap_fcn_t
type Mmap_fcn_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mode_t
type Mode_t = uint16

// See: https://developer.apple.com/documentation/kernel/mount_bptr_t
type Mount_bptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/mount_ptr_ref_t
type Mount_ptr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/mount_ptr_t
type Mount_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/mount_ref_ptr_t
type Mount_ref_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/mount_ref_ref_t
type Mount_ref_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/mount_ref_t
type Mount_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/mount_t
type Mount_t = uintptr

// See: https://developer.apple.com/documentation/kernel/mph_panic_flags_t
type Mph_panic_flags_t = uint64

// See: https://developer.apple.com/documentation/kernel/mpsc_queue_chain_t
// Mpsc_queue_chain_t is an unresolved C aggregate typedef.
type Mpsc_queue_chain_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mpsc_queue_head_t
// Mpsc_queue_head_t is an unresolved C aggregate typedef.
type Mpsc_queue_head_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/msg_labels_t
// Msg_labels_t is opaque storage with the size and alignment C gives msg_labels_t:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type Msg_labels_t [1]uint32

// See: https://developer.apple.com/documentation/kernel/msglen_t
type Msglen_t = uint

// See: https://developer.apple.com/documentation/kernel/msgqnum_t
type Msgqnum_t = uint

// See: https://developer.apple.com/documentation/kernel/n_long
type N_long = uint32

// See: https://developer.apple.com/documentation/kernel/n_short
type N_short = uint16

// See: https://developer.apple.com/documentation/kernel/n_time
type N_time = uint32

// See: https://developer.apple.com/documentation/kernel/natural_t
type Natural_t = uint32

// See: https://developer.apple.com/documentation/kernel/net_init_func_ptr
type Net_init_func_ptr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/netaddr_t
type Netaddr_t = uint32

// See: https://developer.apple.com/documentation/kernel/network_port_t
// Network_port_t is opaque storage with the size and alignment C gives network_port_t:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type Network_port_t [5]uint64

// See: https://developer.apple.com/documentation/kernel/nfs_fsid
// Nfs_fsid is opaque storage with the size and alignment C gives nfs_fsid:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Nfs_fsid [2]uint64

// See: https://developer.apple.com/documentation/kernel/nfs_handle
type Nfs_handle = byte

// See: https://developer.apple.com/documentation/kernel/nfs_specdata
// Nfs_specdata is opaque storage with the size and alignment C gives nfs_specdata:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Nfs_specdata [2]uint32

// See: https://developer.apple.com/documentation/kernel/nfs_stateid
// Nfs_stateid is opaque storage with the size and alignment C gives nfs_stateid:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Nfs_stateid [4]uint32

// See: https://developer.apple.com/documentation/kernel/nfs_supported_kerberos_etypes
type Nfs_supported_kerberos_etypes = uint32

// See: https://developer.apple.com/documentation/kernel/nfserr_info_t
// Nfserr_info_t is opaque storage with the size and alignment C gives nfserr_info_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Nfserr_info_t [2]uint64

// See: https://developer.apple.com/documentation/kernel/nfstype
type Nfstype = uint32

// See: https://developer.apple.com/documentation/kernel/nfsuint64
// Nfsuint64 is opaque storage with the size and alignment C gives nfsuint64:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Nfsuint64 [2]uint32

// See: https://developer.apple.com/documentation/kernel/nlink_t
type Nlink_t = uint16

// See: https://developer.apple.com/documentation/kernel/notify_port_t
type Notify_port_t = uint32

// See: https://developer.apple.com/documentation/kernel/np_uid_t
// Np_uid_t is opaque storage with the size and alignment C gives np_uid_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Np_uid_t [2]uint64

// See: https://developer.apple.com/documentation/kernel/nspace_name_t
type Nspace_name_t = int8

// See: https://developer.apple.com/documentation/kernel/nspace_path_t
type Nspace_path_t = int8

// See: https://developer.apple.com/documentation/kernel/ntsid_t
// Ntsid_t is opaque storage with the size and alignment C gives ntsid_t:
// 72 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 72 into.
type Ntsid_t [72]byte

// See: https://developer.apple.com/documentation/kernel/off_t
type Off_t = int64

// See: https://developer.apple.com/documentation/kernel/open_close_fcn_t
type Open_close_fcn_t = *objc.ID

// See: https://developer.apple.com/documentation/kernel/os_block_t
type Os_block_t = func()

// See: https://developer.apple.com/documentation/kernel/os_log_coproc_reg_t
type Os_log_coproc_reg_t = uint32

// See: https://developer.apple.com/documentation/kernel/os_log_t
type Os_log_t = objectivec.Object

// See: https://developer.apple.com/documentation/kernel/os_log_type_t
type Os_log_type_t = OsLogType

// See: https://developer.apple.com/documentation/kernel/packed_uchar16
type Packed_uchar16 = uint8

// See: https://developer.apple.com/documentation/kernel/packed_uchar32
type Packed_uchar32 = uint8

// See: https://developer.apple.com/documentation/kernel/packed_uchar64
type Packed_uchar64 = uint8

// See: https://developer.apple.com/documentation/kernel/packed_ushort4
type Packed_ushort4 = uint16

// See: https://developer.apple.com/documentation/kernel/page_address_array_t
type Page_address_array_t = *Vm_offset_t

// See: https://developer.apple.com/documentation/kernel/pid_t
type Pid_t = int32

// See: https://developer.apple.com/documentation/kernel/pkthdr_bptr_t
type Pkthdr_bptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/pkthdr_ptr_ref_t
type Pkthdr_ptr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/pkthdr_ptr_t
type Pkthdr_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/pkthdr_ref_ptr_t
type Pkthdr_ref_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/pkthdr_ref_ref_t
type Pkthdr_ref_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/pkthdr_ref_t
type Pkthdr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/pkthdr_t
type Pkthdr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/pointer_t
type Pointer_t = uintptr

// See: https://developer.apple.com/documentation/kernel/pointer_ut
type Pointer_ut = uintptr

// See: https://developer.apple.com/documentation/kernel/policy_base_data_t
// Policy_base_data_t is opaque storage with the size and alignment C gives policy_base_data_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Policy_base_data_t [4]uint32

// See: https://developer.apple.com/documentation/kernel/policy_base_t
type Policy_base_t = *Integer_t

// See: https://developer.apple.com/documentation/kernel/policy_fifo_base_data_t
// Policy_fifo_base_data_t is opaque storage with the size and alignment C gives policy_fifo_base_data_t:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type Policy_fifo_base_data_t [1]uint32

// See: https://developer.apple.com/documentation/kernel/policy_fifo_base_t
type Policy_fifo_base_t = uintptr

// See: https://developer.apple.com/documentation/kernel/policy_fifo_info_data_t
// Policy_fifo_info_data_t is opaque storage with the size and alignment C gives policy_fifo_info_data_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Policy_fifo_info_data_t [4]uint32

// See: https://developer.apple.com/documentation/kernel/policy_fifo_info_t
type Policy_fifo_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/policy_fifo_limit_data_t
// Policy_fifo_limit_data_t is opaque storage with the size and alignment C gives policy_fifo_limit_data_t:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type Policy_fifo_limit_data_t [1]uint32

// See: https://developer.apple.com/documentation/kernel/policy_fifo_limit_t
type Policy_fifo_limit_t = uintptr

// See: https://developer.apple.com/documentation/kernel/policy_info_data_t
// Policy_info_data_t is opaque storage with the size and alignment C gives policy_info_data_t:
// 56 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 56 into.
type Policy_info_data_t [14]uint32

// See: https://developer.apple.com/documentation/kernel/policy_info_t
type Policy_info_t = *Integer_t

// See: https://developer.apple.com/documentation/kernel/policy_limit_data_t
// Policy_limit_data_t is opaque storage with the size and alignment C gives policy_limit_data_t:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type Policy_limit_data_t [3]uint32

// See: https://developer.apple.com/documentation/kernel/policy_limit_t
type Policy_limit_t = *Integer_t

// See: https://developer.apple.com/documentation/kernel/policy_rr_base_data_t
// Policy_rr_base_data_t is opaque storage with the size and alignment C gives policy_rr_base_data_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Policy_rr_base_data_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/policy_rr_base_t
type Policy_rr_base_t = uintptr

// See: https://developer.apple.com/documentation/kernel/policy_rr_info_data_t
// Policy_rr_info_data_t is opaque storage with the size and alignment C gives policy_rr_info_data_t:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type Policy_rr_info_data_t [5]uint32

// See: https://developer.apple.com/documentation/kernel/policy_rr_info_t
type Policy_rr_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/policy_rr_limit_data_t
// Policy_rr_limit_data_t is opaque storage with the size and alignment C gives policy_rr_limit_data_t:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type Policy_rr_limit_data_t [1]uint32

// See: https://developer.apple.com/documentation/kernel/policy_rr_limit_t
type Policy_rr_limit_t = uintptr

// See: https://developer.apple.com/documentation/kernel/policy_t
type Policy_t = int32

// See: https://developer.apple.com/documentation/kernel/policy_timeshare_base_data_t
// Policy_timeshare_base_data_t is opaque storage with the size and alignment C gives policy_timeshare_base_data_t:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type Policy_timeshare_base_data_t [1]uint32

// See: https://developer.apple.com/documentation/kernel/policy_timeshare_base_t
type Policy_timeshare_base_t = uintptr

// See: https://developer.apple.com/documentation/kernel/policy_timeshare_info_data_t
// Policy_timeshare_info_data_t is opaque storage with the size and alignment C gives policy_timeshare_info_data_t:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type Policy_timeshare_info_data_t [5]uint32

// See: https://developer.apple.com/documentation/kernel/policy_timeshare_info_t
type Policy_timeshare_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/policy_timeshare_limit_data_t
// Policy_timeshare_limit_data_t is opaque storage with the size and alignment C gives policy_timeshare_limit_data_t:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type Policy_timeshare_limit_data_t [1]uint32

// See: https://developer.apple.com/documentation/kernel/policy_timeshare_limit_t
type Policy_timeshare_limit_t = uintptr

// See: https://developer.apple.com/documentation/kernel/port_name_array_t
type Port_name_array_t = *Mach_port_name_t

// See: https://developer.apple.com/documentation/kernel/port_name_t
type Port_name_t = Mach_port_name_t

// See: https://developer.apple.com/documentation/kernel/port_t
type Port_t = uint32

// See: https://developer.apple.com/documentation/kernel/posix_cred_t
type Posix_cred_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ppnum_t
type Ppnum_t = uint32

// See: https://developer.apple.com/documentation/kernel/priority_queue_compare_fn_t
type Priority_queue_compare_fn_t = func(e1 unsafe.Pointer, e2 unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/kernel/priority_queue_entry_deadline_t
// Priority_queue_entry_deadline_t is an unresolved C aggregate typedef.
type Priority_queue_entry_deadline_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/priority_queue_entry_sched_modifier_t
type Priority_queue_entry_sched_modifier_t = uint8

// See: https://developer.apple.com/documentation/kernel/priority_queue_entry_sched_t
// Priority_queue_entry_sched_t is an unresolved C aggregate typedef.
type Priority_queue_entry_sched_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/priority_queue_entry_stable_t
// Priority_queue_entry_stable_t is an unresolved C aggregate typedef.
type Priority_queue_entry_stable_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/priority_queue_entry_t
// Priority_queue_entry_t is an unresolved C aggregate typedef.
type Priority_queue_entry_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/priority_queue_key_t
type Priority_queue_key_t = uint16

// See: https://developer.apple.com/documentation/kernel/proc_bptr_t
type Proc_bptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/proc_ident_bptr_t
type Proc_ident_bptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/proc_ident_ptr_ref_t
type Proc_ident_ptr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/proc_ident_ptr_t
type Proc_ident_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/proc_ident_ref_ptr_t
type Proc_ident_ref_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/proc_ident_ref_ref_t
type Proc_ident_ref_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/proc_ident_ref_t
type Proc_ident_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/proc_ident_t
type Proc_ident_t = uintptr

// See: https://developer.apple.com/documentation/kernel/proc_ptr_ref_t
type Proc_ptr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/proc_ptr_t
type Proc_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/proc_ref_ptr_t
type Proc_ref_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/proc_ref_ref_t
type Proc_ref_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/proc_ref_t
type Proc_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/proc_t
type Proc_t = uintptr

// See: https://developer.apple.com/documentation/kernel/processor_array_t
type Processor_array_t = *Processor_t

// See: https://developer.apple.com/documentation/kernel/processor_basic_info_data_t
// Processor_basic_info_data_t is opaque storage with the size and alignment C gives processor_basic_info_data_t:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type Processor_basic_info_data_t [5]uint32

// See: https://developer.apple.com/documentation/kernel/processor_basic_info_t
type Processor_basic_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/processor_cpu_load_info_data_t
// Processor_cpu_load_info_data_t is opaque storage with the size and alignment C gives processor_cpu_load_info_data_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Processor_cpu_load_info_data_t [4]uint32

// See: https://developer.apple.com/documentation/kernel/processor_cpu_load_info_t
type Processor_cpu_load_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/processor_cpu_stat64_data_t
// Processor_cpu_stat64_data_t is opaque storage with the size and alignment C gives processor_cpu_stat64_data_t:
// 80 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 80 into.
type Processor_cpu_stat64_data_t [20]uint32

// See: https://developer.apple.com/documentation/kernel/processor_cpu_stat64_t
type Processor_cpu_stat64_t = uintptr

// See: https://developer.apple.com/documentation/kernel/processor_cpu_stat_data_t
// Processor_cpu_stat_data_t is opaque storage with the size and alignment C gives processor_cpu_stat_data_t:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type Processor_cpu_stat_data_t [9]uint32

// See: https://developer.apple.com/documentation/kernel/processor_cpu_stat_t
type Processor_cpu_stat_t = uintptr

// See: https://developer.apple.com/documentation/kernel/processor_flavor_t
type Processor_flavor_t = int32

// See: https://developer.apple.com/documentation/kernel/processor_info_array_t
type Processor_info_array_t = *Integer_t

// See: https://developer.apple.com/documentation/kernel/processor_info_data_t
type Processor_info_data_t = int32

// See: https://developer.apple.com/documentation/kernel/processor_info_t
type Processor_info_t = *Integer_t

// See: https://developer.apple.com/documentation/kernel/processor_port_array_t
type Processor_port_array_t = Processor_array_t

// See: https://developer.apple.com/documentation/kernel/processor_port_t
type Processor_port_t = Processor_t

// See: https://developer.apple.com/documentation/kernel/processor_set_array_t
type Processor_set_array_t = *Processor_set_t

// See: https://developer.apple.com/documentation/kernel/processor_set_basic_info_data_t
// Processor_set_basic_info_data_t is opaque storage with the size and alignment C gives processor_set_basic_info_data_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Processor_set_basic_info_data_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/processor_set_basic_info_t
type Processor_set_basic_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/processor_set_control_port_t
type Processor_set_control_port_t = Processor_set_t

// See: https://developer.apple.com/documentation/kernel/processor_set_control_t
type Processor_set_control_t = uint32

// See: https://developer.apple.com/documentation/kernel/processor_set_flavor_t
type Processor_set_flavor_t = int32

// See: https://developer.apple.com/documentation/kernel/processor_set_info_data_t
type Processor_set_info_data_t = int32

// See: https://developer.apple.com/documentation/kernel/processor_set_info_t
type Processor_set_info_t = *Integer_t

// See: https://developer.apple.com/documentation/kernel/processor_set_load_info_data_t
// Processor_set_load_info_data_t is opaque storage with the size and alignment C gives processor_set_load_info_data_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Processor_set_load_info_data_t [4]uint32

// See: https://developer.apple.com/documentation/kernel/processor_set_load_info_t
type Processor_set_load_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/processor_set_name_array_t
type Processor_set_name_array_t = *Processor_set_t

// See: https://developer.apple.com/documentation/kernel/processor_set_name_port_array_t
type Processor_set_name_port_array_t = Processor_set_array_t

// See: https://developer.apple.com/documentation/kernel/processor_set_name_port_t
type Processor_set_name_port_t = Processor_set_t

// See: https://developer.apple.com/documentation/kernel/processor_set_name_t
type Processor_set_name_t = Processor_set_t

// See: https://developer.apple.com/documentation/kernel/processor_set_port_t
type Processor_set_port_t = Processor_set_t

// See: https://developer.apple.com/documentation/kernel/processor_set_t
type Processor_set_t = uint32

// See: https://developer.apple.com/documentation/kernel/processor_t
type Processor_t = uint32

// Protocol_family_t is storage type for the protocol family.
//
// See: https://developer.apple.com/documentation/kernel/protocol_family_t
type Protocol_family_t = uint32

// See: https://developer.apple.com/documentation/kernel/psize_fcn_t
type Psize_fcn_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ptrdiff_t
type Ptrdiff_t = int

// See: https://developer.apple.com/documentation/kernel/qaddr_t
type Qaddr_t = *Quad_t

// See: https://developer.apple.com/documentation/kernel/quad_t
type Quad_t = int64

// See: https://developer.apple.com/documentation/kernel/queue_chain_t
// Queue_chain_t is opaque storage with the size and alignment C gives queue_chain_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Queue_chain_t [2]uint64

// See: https://developer.apple.com/documentation/kernel/queue_entry_t
type Queue_entry_t = uintptr

// See: https://developer.apple.com/documentation/kernel/queue_head_t
// Queue_head_t is opaque storage with the size and alignment C gives queue_head_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Queue_head_t [2]uint64

// See: https://developer.apple.com/documentation/kernel/queue_t
type Queue_t = uintptr

// See: https://developer.apple.com/documentation/kernel/read_write_fcn_t
type Read_write_fcn_t = *objc.ID

// See: https://developer.apple.com/documentation/kernel/reg64_t
type Reg64_t = uint32

// See: https://developer.apple.com/documentation/kernel/register_t
type Register_t = int64

// See: https://developer.apple.com/documentation/kernel/reset_fcn_t
type Reset_fcn_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/rlim_t
type Rlim_t = uint64

// See: https://developer.apple.com/documentation/kernel/route_t
type Route_t = uintptr

// See: https://developer.apple.com/documentation/kernel/routine_arg_descriptor
// Routine_arg_descriptor is opaque storage with the size and alignment C gives routine_arg_descriptor:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type Routine_arg_descriptor [3]uint32

// See: https://developer.apple.com/documentation/kernel/routine_arg_descriptor_t
type Routine_arg_descriptor_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/routine_arg_offset
type Routine_arg_offset = uint32

// See: https://developer.apple.com/documentation/kernel/routine_arg_size
type Routine_arg_size = uint32

// See: https://developer.apple.com/documentation/kernel/routine_arg_type
type Routine_arg_type = uint32

// See: https://developer.apple.com/documentation/kernel/routine_descriptor_t
type Routine_descriptor_t = uintptr

// See: https://developer.apple.com/documentation/kernel/rpc_routine_arg_descriptor_t
// Rpc_routine_arg_descriptor_t is an unresolved C aggregate typedef.
type Rpc_routine_arg_descriptor_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/rpc_routine_descriptor_t
// Rpc_routine_descriptor_t is an unresolved C aggregate typedef.
type Rpc_routine_descriptor_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/rpc_subsystem_t
// Rpc_subsystem_t is an unresolved C aggregate typedef.
type Rpc_subsystem_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/rsize_t
type Rsize_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/rsvd_fcn_t
type Rsvd_fcn_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/rtentry_bptr_t
type Rtentry_bptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/rtentry_ptr_ref_t
type Rtentry_ptr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/rtentry_ptr_t
type Rtentry_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/rtentry_ref_ptr_t
type Rtentry_ref_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/rtentry_ref_ref_t
type Rtentry_ref_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/rtentry_ref_t
type Rtentry_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/rune_t
type Rune_t = int32

// See: https://developer.apple.com/documentation/kernel/rusage_info_current
// Rusage_info_current is opaque storage with the size and alignment C gives rusage_info_current:
// 464 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 464 into.
type Rusage_info_current [58]uint64

// See: https://developer.apple.com/documentation/kernel/rusage_info_t
type Rusage_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/sa_endpoints_t
// Sa_endpoints_t is opaque storage with the size and alignment C gives sa_endpoints_t:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type Sa_endpoints_t [5]uint64

// See: https://developer.apple.com/documentation/kernel/sa_family_t
type Sa_family_t = uint8

// See: https://developer.apple.com/documentation/kernel/sae_associd_t
type Sae_associd_t = uint32

// See: https://developer.apple.com/documentation/kernel/sae_connid_t
type Sae_connid_t = uint32

// See: https://developer.apple.com/documentation/kernel/secure_boot_cryptex_args_t
// Secure_boot_cryptex_args_t is opaque storage with the size and alignment C gives secure_boot_cryptex_args_t:
// 28 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 28 into.
type Secure_boot_cryptex_args_t [7]uint32

// See: https://developer.apple.com/documentation/kernel/security_token_t
// Security_token_t is opaque storage with the size and alignment C gives security_token_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Security_token_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/segsz_t
type Segsz_t = int32

// See: https://developer.apple.com/documentation/kernel/sel_t
// Sel_t is opaque storage with the size and alignment C gives sel_t:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type Sel_t [1]uint16

// See: https://developer.apple.com/documentation/kernel/select_fcn_t
type Select_fcn_t = *objc.ID

// See: https://developer.apple.com/documentation/kernel/semaphore_port_t
type Semaphore_port_t = Semaphore_t

// See: https://developer.apple.com/documentation/kernel/semaphore_t
type Semaphore_t = uint32

// See: https://developer.apple.com/documentation/kernel/sflt_data_flag_t
type Sflt_data_flag_t = uint32

// See: https://developer.apple.com/documentation/kernel/sflt_event_t
type Sflt_event_t = uint32

// See: https://developer.apple.com/documentation/kernel/sflt_flags
type Sflt_flags = uint32

// Sflt_handle is a 4 byte identifier used with the SO_NKE socket option to identify the socket filter to be attached.
//
// See: https://developer.apple.com/documentation/kernel/sflt_handle
type Sflt_handle = uint32

// See: https://developer.apple.com/documentation/kernel/shared_file_mapping_slide_np_t
// Shared_file_mapping_slide_np_t is opaque storage with the size and alignment C gives shared_file_mapping_slide_np_t:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type Shared_file_mapping_slide_np_t [6]uint64

// See: https://developer.apple.com/documentation/kernel/shared_file_mapping_slide_np_ut
type Shared_file_mapping_slide_np_ut = Shared_file_mapping_slide_np_t

// See: https://developer.apple.com/documentation/kernel/shmatt_t
type Shmatt_t = uint16

// See: https://developer.apple.com/documentation/kernel/sig_atomic_t
type Sig_atomic_t = int32

// See: https://developer.apple.com/documentation/kernel/sig_t
type Sig_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/siginfo_t
// Siginfo_t is opaque storage with the size and alignment C gives siginfo_t:
// 104 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 104 into.
type Siginfo_t [13]uint64

// See: https://developer.apple.com/documentation/kernel/sigset_t
type Sigset_t = uint32

// See: https://developer.apple.com/documentation/kernel/size_t
type Size_t = uintptr

// See: https://developer.apple.com/documentation/kernel/size_ut
type Size_ut = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/sleepwakenote
// SleepWakeNote is opaque storage with the size and alignment C gives sleepWakeNote:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type SleepWakeNote [4]uint64

// See: https://developer.apple.com/documentation/kernel/sleep_type_t
type Sleep_type_t = int32

// See: https://developer.apple.com/documentation/kernel/smr_cb_t
type Smr_cb_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/smr_node_t
// Smr_node_t is an unresolved C aggregate typedef.
type Smr_node_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/smr_seq_t
type Smr_seq_t = uint

// See: https://developer.apple.com/documentation/kernel/smr_t
type Smr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/so_gen_t
type So_gen_t = uint64

// See: https://developer.apple.com/documentation/kernel/sock_storage
type Sock_storage = uint32

// See: https://developer.apple.com/documentation/kernel/sockaddr_ptr_ref_t
type Sockaddr_ptr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/sockaddr_ref_ptr_t
type Sockaddr_ref_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/sockaddr_ref_ref_t
type Sockaddr_ref_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/sockaddr_ref_t
// Sockaddr_ref_t is an unresolved C aggregate typedef.
type Sockaddr_ref_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/sockaddr_storage_ptr_ref_t
type Sockaddr_storage_ptr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/sockaddr_storage_ref_ptr_t
type Sockaddr_storage_ref_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/sockaddr_storage_ref_ref_t
type Sockaddr_storage_ref_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/sockaddr_storage_ref_t
// Sockaddr_storage_ref_t is an unresolved C aggregate typedef.
type Sockaddr_storage_ref_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/socket_bptr_t
type Socket_bptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/socket_ptr_ref_t
type Socket_ptr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/socket_ptr_t
type Socket_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/socket_ref_ptr_t
type Socket_ref_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/socket_ref_ref_t
type Socket_ref_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/socket_ref_t
type Socket_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/socket_t
type Socket_t = uintptr

// See: https://developer.apple.com/documentation/kernel/socklen_t
type Socklen_t = uint32

// See: https://developer.apple.com/documentation/kernel/sockopt_bptr_t
type Sockopt_bptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/sockopt_dir
type Sockopt_dir = byte

// See: https://developer.apple.com/documentation/kernel/sockopt_ptr_ref_t
type Sockopt_ptr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/sockopt_ptr_t
type Sockopt_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/sockopt_ref_ptr_t
type Sockopt_ref_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/sockopt_ref_ref_t
type Sockopt_ref_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/sockopt_ref_t
type Sockopt_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/sockopt_t
type Sockopt_t = uintptr

// See: https://developer.apple.com/documentation/kernel/speed_t
type Speed_t = uint

// See: https://developer.apple.com/documentation/kernel/sptm_asid_t
type Sptm_asid_t = uint16

// See: https://developer.apple.com/documentation/kernel/sptm_call_regs_t
// Sptm_call_regs_t is opaque storage with the size and alignment C gives sptm_call_regs_t:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type Sptm_call_regs_t [8]uint64

// See: https://developer.apple.com/documentation/kernel/sptm_consistent_debug_t
// Sptm_consistent_debug_t is opaque storage with the size and alignment C gives sptm_consistent_debug_t:
// 2672 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2672 into.
type Sptm_consistent_debug_t [2672]byte

// See: https://developer.apple.com/documentation/kernel/sptm_dispatch_endpoint_id_t
type Sptm_dispatch_endpoint_id_t = byte

// See: https://developer.apple.com/documentation/kernel/sptm_dispatch_table_id_t
type Sptm_dispatch_table_id_t = byte

// See: https://developer.apple.com/documentation/kernel/sptm_dispatch_target_t
type Sptm_dispatch_target_t = uint64

// See: https://developer.apple.com/documentation/kernel/sptm_domain_t
type Sptm_domain_t = byte

// See: https://developer.apple.com/documentation/kernel/sptm_frame_type_t
type Sptm_frame_type_t = byte

// See: https://developer.apple.com/documentation/kernel/sptm_instance_id_t
type Sptm_instance_id_t = uint16

// See: https://developer.apple.com/documentation/kernel/sptm_iommu_id_t
type Sptm_iommu_id_t = byte

// See: https://developer.apple.com/documentation/kernel/sptm_iommu_retype_params_t
type Sptm_iommu_retype_params_t = uint32

// See: https://developer.apple.com/documentation/kernel/sptm_paddr_t
type Sptm_paddr_t = uint64

// See: https://developer.apple.com/documentation/kernel/sptm_papt_t
type Sptm_papt_t = uint64

// See: https://developer.apple.com/documentation/kernel/sptm_poff_t
type Sptm_poff_t = uint64

// See: https://developer.apple.com/documentation/kernel/sptm_ppnum_t
type Sptm_ppnum_t = uint32

// See: https://developer.apple.com/documentation/kernel/sptm_pt_level_t
type Sptm_pt_level_t = byte

// See: https://developer.apple.com/documentation/kernel/sptm_pte_t
type Sptm_pte_t = uint64

// See: https://developer.apple.com/documentation/kernel/sptm_return_t
type Sptm_return_t = uint32

// See: https://developer.apple.com/documentation/kernel/sptm_retype_params_t
// Sptm_retype_params_t is opaque storage with the size and alignment C gives sptm_retype_params_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Sptm_retype_params_t [1]uint64

// See: https://developer.apple.com/documentation/kernel/sptm_trace_buffer_t
// Sptm_trace_buffer_t is an unresolved C aggregate typedef.
type Sptm_trace_buffer_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/sptm_trace_t
// Sptm_trace_t is opaque storage with the size and alignment C gives sptm_trace_t:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type Sptm_trace_t [8]uint64

// See: https://developer.apple.com/documentation/kernel/sptm_tte_t
type Sptm_tte_t = uint64

// See: https://developer.apple.com/documentation/kernel/sptm_vaddr_t
type Sptm_vaddr_t = uint64

// See: https://developer.apple.com/documentation/kernel/sptm_vector_type_t
type Sptm_vector_type_t = uint8

// See: https://developer.apple.com/documentation/kernel/sptm_vmid_t
type Sptm_vmid_t = uint16

// See: https://developer.apple.com/documentation/kernel/sptm_voff_t
type Sptm_voff_t = uint64

// See: https://developer.apple.com/documentation/kernel/ssize_t
type Ssize_t = int

// See: https://developer.apple.com/documentation/kernel/stack_t
// Stack_t is opaque storage with the size and alignment C gives stack_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Stack_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/stackshot_flags_t
type Stackshot_flags_t = StackshotFlags

// See: https://developer.apple.com/documentation/kernel/stop_fcn_t
type Stop_fcn_t = *objc.ID

// See: https://developer.apple.com/documentation/kernel/strategy_fcn_t
type Strategy_fcn_t = *objc.ID

// See: https://developer.apple.com/documentation/kernel/string_t
type String_t = *byte

// See: https://developer.apple.com/documentation/kernel/subaid_t
type Subaid_t = uint64

// See: https://developer.apple.com/documentation/kernel/suseconds_t
type Suseconds_t = int32

// See: https://developer.apple.com/documentation/kernel/swblk_t
type Swblk_t = int32

// See: https://developer.apple.com/documentation/kernel/symtab_name_t
type Symtab_name_t = int8

// See: https://developer.apple.com/documentation/kernel/sync_policy_t
type Sync_policy_t = int32

// See: https://developer.apple.com/documentation/kernel/syscall_arg_t
type Syscall_arg_t = uint64

// See: https://developer.apple.com/documentation/kernel/syscp_id_instructions_feat_1_reg
// Syscp_ID_instructions_feat_1_reg is opaque storage with the size and alignment C gives syscp_ID_instructions_feat_1_reg:
// 1 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 1 into.
type Syscp_ID_instructions_feat_1_reg [1]byte

// See: https://developer.apple.com/documentation/kernel/tdevicerequestdirection
type TDeviceRequestDirection = uint32

// See: https://developer.apple.com/documentation/kernel/tdevicerequestrecipient
type TDeviceRequestRecipient = uint32

// See: https://developer.apple.com/documentation/kernel/tdevicerequesttype
type TDeviceRequestType = uint32

// See: https://developer.apple.com/documentation/kernel/tendpointdirection
type TEndpointDirection = uint32

// See: https://developer.apple.com/documentation/kernel/tendpointsynchronizationtype
type TEndpointSynchronizationType = uint32

// See: https://developer.apple.com/documentation/kernel/tendpointtype
type TEndpointType = uint32

// See: https://developer.apple.com/documentation/kernel/tendpointusagetype
type TEndpointUsageType = uint32

// See: https://developer.apple.com/documentation/kernel/tiopcideviceresetoptions
type TIOPCIDeviceResetOptions = uint32

// See: https://developer.apple.com/documentation/kernel/tiopcideviceresettypes
type TIOPCIDeviceResetTypes = uint32

// See: https://developer.apple.com/documentation/kernel/tiopcilinkspeed
type TIOPCILinkSpeed = uint32

// TIOUSBDescriptorSize is constants for the number of bytes in descriptor structures.
//
// See: https://developer.apple.com/documentation/kernel/tiousbdescriptorsize
type TIOUSBDescriptorSize = int

// TIOUSBDescriptorType is constants describing the types of descriptors available for a USB device.
//
// See: https://developer.apple.com/documentation/kernel/tiousbdescriptortype
type TIOUSBDescriptorType = int

// TIOUSBDeviceCapabilityType is constants for the device capability types.
//
// See: https://developer.apple.com/documentation/kernel/tiousbdevicecapabilitytype
type TIOUSBDeviceCapabilityType = int

// TIOUSBDeviceRequestDirectionValue is enumerated device request direction values.
//
// See: https://developer.apple.com/documentation/kernel/tiousbdevicerequestdirectionvalue
type TIOUSBDeviceRequestDirectionValue = int

// TIOUSBDeviceRequestRecipientValue is constants indicating the type of object that receives the results of a request.
//
// See: https://developer.apple.com/documentation/kernel/tiousbdevicerequestrecipientvalue
type TIOUSBDeviceRequestRecipientValue = int

// TIOUSBDeviceRequestTypeValue is constants indicating the type of request to make from a device.
//
// See: https://developer.apple.com/documentation/kernel/tiousbdevicerequesttypevalue
type TIOUSBDeviceRequestTypeValue = int

// TIOUSBEndpointDirection is the direction of data transfers on an endpoint.
//
// See: https://developer.apple.com/documentation/kernel/tiousbendpointdirection
type TIOUSBEndpointDirection = int

// TIOUSBEndpointSynchronizationType is constants for the endpoint synchronization types.
//
// See: https://developer.apple.com/documentation/kernel/tiousbendpointsynchronizationtype
type TIOUSBEndpointSynchronizationType = int

// TIOUSBEndpointType is constants describing the types of endpoints.
//
// See: https://developer.apple.com/documentation/kernel/tiousbendpointtype
type TIOUSBEndpointType = int

// TIOUSBEndpointUsageType is constants for the endpoint usage types.
//
// See: https://developer.apple.com/documentation/kernel/tiousbendpointusagetype
type TIOUSBEndpointUsageType = int

// TIOUSBLanguageID is constants for the USB language identifiers.
//
// See: https://developer.apple.com/documentation/kernel/tiousblanguageid
type TIOUSBLanguageID = string

// See: https://developer.apple.com/documentation/kernel/tusbctypecabletype
type TUSBCTypeCableType = uint32

// See: https://developer.apple.com/documentation/kernel/tusbdevicelpmstatus
type TUSBDeviceLPMStatus = uint32

// See: https://developer.apple.com/documentation/kernel/tusbhostconnectortype
type TUSBHostConnectorType = uint32

// TUSBHostDeviceAddress is the USB host device address.
//
// See: https://developer.apple.com/documentation/kernel/tusbhostdeviceaddress
type TUSBHostDeviceAddress = uint16

// See: https://developer.apple.com/documentation/kernel/tusbhostportconnectable
type TUSBHostPortConnectable = uint32

// See: https://developer.apple.com/documentation/kernel/tusbhostpowersourcetype
type TUSBHostPowerSourceType = uint32

// See: https://developer.apple.com/documentation/kernel/tusblinkstate
type TUSBLinkState = uint32

// See: https://developer.apple.com/documentation/kernel/task_absolutetime_info_data_t
// Task_absolutetime_info_data_t is opaque storage with the size and alignment C gives task_absolutetime_info_data_t:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type Task_absolutetime_info_data_t [8]uint32

// See: https://developer.apple.com/documentation/kernel/task_absolutetime_info_t
type Task_absolutetime_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_affinity_tag_info_data_t
// Task_affinity_tag_info_data_t is opaque storage with the size and alignment C gives task_affinity_tag_info_data_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Task_affinity_tag_info_data_t [4]uint32

// See: https://developer.apple.com/documentation/kernel/task_affinity_tag_info_t
type Task_affinity_tag_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_array_t
type Task_array_t = *Task_t

// See: https://developer.apple.com/documentation/kernel/task_basic_info_32_data_t
// Task_basic_info_32_data_t is opaque storage with the size and alignment C gives task_basic_info_32_data_t:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type Task_basic_info_32_data_t [8]uint32

// See: https://developer.apple.com/documentation/kernel/task_basic_info_32_t
type Task_basic_info_32_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_basic_info_64_2_data_t
// Task_basic_info_64_2_data_t is opaque storage with the size and alignment C gives task_basic_info_64_2_data_t:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type Task_basic_info_64_2_data_t [10]uint32

// See: https://developer.apple.com/documentation/kernel/task_basic_info_64_2_t
type Task_basic_info_64_2_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_basic_info_64_data_t
// Task_basic_info_64_data_t is opaque storage with the size and alignment C gives task_basic_info_64_data_t:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type Task_basic_info_64_data_t [10]uint32

// See: https://developer.apple.com/documentation/kernel/task_basic_info_64_t
type Task_basic_info_64_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_basic_info_data_t
// Task_basic_info_data_t is opaque storage with the size and alignment C gives task_basic_info_data_t:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type Task_basic_info_data_t [10]uint32

// See: https://developer.apple.com/documentation/kernel/task_basic_info_t
type Task_basic_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_category_policy_data_t
// Task_category_policy_data_t is opaque storage with the size and alignment C gives task_category_policy_data_t:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type Task_category_policy_data_t [1]uint32

// See: https://developer.apple.com/documentation/kernel/task_category_policy_t
type Task_category_policy_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_corpse_forking_behavior_t
type Task_corpse_forking_behavior_t = uint32

// See: https://developer.apple.com/documentation/kernel/task_crashinfo_item_t
// Task_crashinfo_item_t is an unresolved C aggregate typedef.
type Task_crashinfo_item_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_dyld_info_data_t
// Task_dyld_info_data_t is opaque storage with the size and alignment C gives task_dyld_info_data_t:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type Task_dyld_info_data_t [5]uint32

// See: https://developer.apple.com/documentation/kernel/task_dyld_info_t
type Task_dyld_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_events_info_data_t
// Task_events_info_data_t is opaque storage with the size and alignment C gives task_events_info_data_t:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type Task_events_info_data_t [8]uint32

// See: https://developer.apple.com/documentation/kernel/task_events_info_t
type Task_events_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_exc_guard_behavior_t
type Task_exc_guard_behavior_t = uint32

// See: https://developer.apple.com/documentation/kernel/task_extmod_info_data_t
// Task_extmod_info_data_t is opaque storage with the size and alignment C gives task_extmod_info_data_t:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type Task_extmod_info_data_t [16]uint32

// See: https://developer.apple.com/documentation/kernel/task_extmod_info_t
type Task_extmod_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_flags_info_data_t
// Task_flags_info_data_t is opaque storage with the size and alignment C gives task_flags_info_data_t:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type Task_flags_info_data_t [1]uint32

// See: https://developer.apple.com/documentation/kernel/task_flags_info_t
type Task_flags_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_flavor_t
type Task_flavor_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/task_gate_t
// Task_gate_t is opaque storage with the size and alignment C gives task_gate_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Task_gate_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/task_id_token_t
type Task_id_token_t = uint32

// See: https://developer.apple.com/documentation/kernel/task_info_data_t
type Task_info_data_t = int32

// See: https://developer.apple.com/documentation/kernel/task_info_t
type Task_info_t = *Integer_t

// See: https://developer.apple.com/documentation/kernel/task_inspect_basic_counts_data_t
// Task_inspect_basic_counts_data_t is opaque storage with the size and alignment C gives task_inspect_basic_counts_data_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Task_inspect_basic_counts_data_t [2]uint64

// See: https://developer.apple.com/documentation/kernel/task_inspect_basic_counts_t
type Task_inspect_basic_counts_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_inspect_flavor_t
type Task_inspect_flavor_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/task_inspect_info_t
type Task_inspect_info_t = *Integer_t

// See: https://developer.apple.com/documentation/kernel/task_inspect_t
type Task_inspect_t = uint32

// See: https://developer.apple.com/documentation/kernel/task_kernelmemory_info_data_t
// Task_kernelmemory_info_data_t is opaque storage with the size and alignment C gives task_kernelmemory_info_data_t:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type Task_kernelmemory_info_data_t [8]uint32

// See: https://developer.apple.com/documentation/kernel/task_kernelmemory_info_t
type Task_kernelmemory_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_latency_qos_t
type Task_latency_qos_t = int32

// See: https://developer.apple.com/documentation/kernel/task_name_t
type Task_name_t = uint32

// See: https://developer.apple.com/documentation/kernel/task_policy_flavor_t
type Task_policy_flavor_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/task_policy_get_t
type Task_policy_get_t = uint32

// See: https://developer.apple.com/documentation/kernel/task_policy_set_t
type Task_policy_set_t = uint32

// See: https://developer.apple.com/documentation/kernel/task_policy_t
type Task_policy_t = *Integer_t

// See: https://developer.apple.com/documentation/kernel/task_port_array_t
type Task_port_array_t = Task_array_t

// See: https://developer.apple.com/documentation/kernel/task_port_t
type Task_port_t = Task_t

// See: https://developer.apple.com/documentation/kernel/task_power_info_data_t
// Task_power_info_data_t is opaque storage with the size and alignment C gives task_power_info_data_t:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type Task_power_info_data_t [12]uint32

// See: https://developer.apple.com/documentation/kernel/task_power_info_t
type Task_power_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_power_info_v2_data_t
// Task_power_info_v2_data_t is opaque storage with the size and alignment C gives task_power_info_v2_data_t:
// 104 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 104 into.
type Task_power_info_v2_data_t [26]uint32

// See: https://developer.apple.com/documentation/kernel/task_power_info_v2_t
type Task_power_info_v2_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_purgable_info_t
// Task_purgable_info_t is opaque storage with the size and alignment C gives task_purgable_info_t:
// 272 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 272 into.
type Task_purgable_info_t [34]uint64

// See: https://developer.apple.com/documentation/kernel/task_qos_policy_t
// Task_qos_policy_t is an unresolved C aggregate typedef.
type Task_qos_policy_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_read_t
type Task_read_t = uint32

// See: https://developer.apple.com/documentation/kernel/task_restartable_range_array_t
type Task_restartable_range_array_t = *Task_restartable_range_t

// See: https://developer.apple.com/documentation/kernel/task_restartable_range_t
// Task_restartable_range_t is opaque storage with the size and alignment C gives task_restartable_range_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Task_restartable_range_t [2]uint64

// See: https://developer.apple.com/documentation/kernel/task_role_t
type Task_role_t = int32

// See: https://developer.apple.com/documentation/kernel/task_special_port_t
type Task_special_port_t = int32

// See: https://developer.apple.com/documentation/kernel/task_suspension_token_t
type Task_suspension_token_t = uint32

// See: https://developer.apple.com/documentation/kernel/task_t
type Task_t = uint32

// See: https://developer.apple.com/documentation/kernel/task_thread_times_info_data_t
// Task_thread_times_info_data_t is opaque storage with the size and alignment C gives task_thread_times_info_data_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Task_thread_times_info_data_t [4]uint32

// See: https://developer.apple.com/documentation/kernel/task_thread_times_info_t
type Task_thread_times_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_throughput_qos_t
type Task_throughput_qos_t = int32

// See: https://developer.apple.com/documentation/kernel/task_trace_memory_info_data_t
// Task_trace_memory_info_data_t is opaque storage with the size and alignment C gives task_trace_memory_info_data_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Task_trace_memory_info_data_t [6]uint32

// See: https://developer.apple.com/documentation/kernel/task_trace_memory_info_t
type Task_trace_memory_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_vm_info_data_t
// Task_vm_info_data_t is opaque storage with the size and alignment C gives task_vm_info_data_t:
// 372 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 372 into.
type Task_vm_info_data_t [93]uint32

// See: https://developer.apple.com/documentation/kernel/task_vm_info_t
type Task_vm_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_wait_state_info_data_t
// Task_wait_state_info_data_t is opaque storage with the size and alignment C gives task_wait_state_info_data_t:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type Task_wait_state_info_data_t [8]uint32

// See: https://developer.apple.com/documentation/kernel/task_wait_state_info_t
type Task_wait_state_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_zone_info_array_t
type Task_zone_info_array_t = *Task_zone_info_t

// See: https://developer.apple.com/documentation/kernel/task_zone_info_t
// Task_zone_info_t is opaque storage with the size and alignment C gives task_zone_info_t:
// 88 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 88 into.
type Task_zone_info_t [11]uint64

// See: https://developer.apple.com/documentation/kernel/tcflag_t
type Tcflag_t = uint

// See: https://developer.apple.com/documentation/kernel/tcp_cc
type Tcp_cc = uint32

// See: https://developer.apple.com/documentation/kernel/tcp_connection_client_accurate_ecn_state_t
type Tcp_connection_client_accurate_ecn_state_t = uint32

// See: https://developer.apple.com/documentation/kernel/tcp_connection_server_accurate_ecn_state_t
type Tcp_connection_server_accurate_ecn_state_t = uint32

// See: https://developer.apple.com/documentation/kernel/tcp_notify_ack_id_t
type Tcp_notify_ack_id_t = uint32

// See: https://developer.apple.com/documentation/kernel/tcp_seq
type Tcp_seq = uint32

// See: https://developer.apple.com/documentation/kernel/telemetry_notice_t
type Telemetry_notice_t = TelemetryNotice

// See: https://developer.apple.com/documentation/kernel/text_encoding_t
type Text_encoding_t = uint32

// See: https://developer.apple.com/documentation/kernel/thread_act_array_t
type Thread_act_array_t = *Thread_act_t

// See: https://developer.apple.com/documentation/kernel/thread_act_port_array_t
type Thread_act_port_array_t = Thread_act_array_t

// See: https://developer.apple.com/documentation/kernel/thread_act_port_t
type Thread_act_port_t = Thread_act_t

// See: https://developer.apple.com/documentation/kernel/thread_act_t
type Thread_act_t = uint32

// See: https://developer.apple.com/documentation/kernel/thread_affinity_policy_data_t
// Thread_affinity_policy_data_t is opaque storage with the size and alignment C gives thread_affinity_policy_data_t:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type Thread_affinity_policy_data_t [1]uint32

// See: https://developer.apple.com/documentation/kernel/thread_affinity_policy_t
type Thread_affinity_policy_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_array_t
type Thread_array_t = *Thread_t

// See: https://developer.apple.com/documentation/kernel/thread_background_policy_data_t
// Thread_background_policy_data_t is opaque storage with the size and alignment C gives thread_background_policy_data_t:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type Thread_background_policy_data_t [1]uint32

// See: https://developer.apple.com/documentation/kernel/thread_background_policy_t
type Thread_background_policy_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_basic_info_data_t
// Thread_basic_info_data_t is opaque storage with the size and alignment C gives thread_basic_info_data_t:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type Thread_basic_info_data_t [10]uint32

// See: https://developer.apple.com/documentation/kernel/thread_basic_info_t
type Thread_basic_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_call_options_t
type Thread_call_options_t = uint32

// See: https://developer.apple.com/documentation/kernel/thread_call_param_t
type Thread_call_param_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/thread_call_priority_t
type Thread_call_priority_t = ThreadCallPriority

// See: https://developer.apple.com/documentation/kernel/thread_call_t
type Thread_call_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_extended_info_data_t
// Thread_extended_info_data_t is opaque storage with the size and alignment C gives thread_extended_info_data_t:
// 112 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 112 into.
type Thread_extended_info_data_t [14]uint64

// See: https://developer.apple.com/documentation/kernel/thread_extended_info_t
type Thread_extended_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_extended_policy_data_t
// Thread_extended_policy_data_t is opaque storage with the size and alignment C gives thread_extended_policy_data_t:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type Thread_extended_policy_data_t [1]uint32

// See: https://developer.apple.com/documentation/kernel/thread_extended_policy_t
type Thread_extended_policy_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_flavor_t
type Thread_flavor_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/thread_identifier_info_data_t
// Thread_identifier_info_data_t is opaque storage with the size and alignment C gives thread_identifier_info_data_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Thread_identifier_info_data_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/thread_identifier_info_t
type Thread_identifier_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_info_data_t
type Thread_info_data_t = int32

// See: https://developer.apple.com/documentation/kernel/thread_info_t
type Thread_info_t = *Integer_t

// See: https://developer.apple.com/documentation/kernel/thread_inspect_t
type Thread_inspect_t = uint32

// See: https://developer.apple.com/documentation/kernel/thread_latency_qos_policy_data_t
// Thread_latency_qos_policy_data_t is opaque storage with the size and alignment C gives thread_latency_qos_policy_data_t:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type Thread_latency_qos_policy_data_t [1]uint32

// See: https://developer.apple.com/documentation/kernel/thread_latency_qos_policy_t
type Thread_latency_qos_policy_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_latency_qos_t
type Thread_latency_qos_t = int32

// See: https://developer.apple.com/documentation/kernel/thread_policy_flavor_t
type Thread_policy_flavor_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/thread_policy_t
type Thread_policy_t = *Integer_t

// See: https://developer.apple.com/documentation/kernel/thread_port_array_t
type Thread_port_array_t = Thread_array_t

// See: https://developer.apple.com/documentation/kernel/thread_port_t
type Thread_port_t = Thread_t

// See: https://developer.apple.com/documentation/kernel/thread_precedence_policy_data_t
// Thread_precedence_policy_data_t is opaque storage with the size and alignment C gives thread_precedence_policy_data_t:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type Thread_precedence_policy_data_t [1]uint32

// See: https://developer.apple.com/documentation/kernel/thread_precedence_policy_t
type Thread_precedence_policy_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_read_t
type Thread_read_t = uint32

// See: https://developer.apple.com/documentation/kernel/thread_selfcounts_kind_t
type Thread_selfcounts_kind_t = uint32

// See: https://developer.apple.com/documentation/kernel/thread_standard_policy_data_t
// Thread_standard_policy_data_t is opaque storage with the size and alignment C gives thread_standard_policy_data_t:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type Thread_standard_policy_data_t [1]uint32

// See: https://developer.apple.com/documentation/kernel/thread_standard_policy_t
type Thread_standard_policy_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_state_data_t
type Thread_state_data_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/thread_state_flavor_array_t
type Thread_state_flavor_array_t = *Thread_state_flavor_t

// See: https://developer.apple.com/documentation/kernel/thread_state_flavor_t
type Thread_state_flavor_t = int32

// See: https://developer.apple.com/documentation/kernel/thread_state_t
type Thread_state_t = *Natural_t

// See: https://developer.apple.com/documentation/kernel/thread_t
type Thread_t = uint32

// See: https://developer.apple.com/documentation/kernel/thread_throughput_qos_policy_data_t
// Thread_throughput_qos_policy_data_t is opaque storage with the size and alignment C gives thread_throughput_qos_policy_data_t:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type Thread_throughput_qos_policy_data_t [1]uint32

// See: https://developer.apple.com/documentation/kernel/thread_throughput_qos_policy_t
type Thread_throughput_qos_policy_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_throughput_qos_t
type Thread_throughput_qos_t = int32

// See: https://developer.apple.com/documentation/kernel/thread_time_constraint_policy_data_t
// Thread_time_constraint_policy_data_t is opaque storage with the size and alignment C gives thread_time_constraint_policy_data_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Thread_time_constraint_policy_data_t [4]uint32

// See: https://developer.apple.com/documentation/kernel/thread_time_constraint_policy_t
type Thread_time_constraint_policy_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_turnstileinfo_t
// Thread_turnstileinfo_t is opaque storage with the size and alignment C gives thread_turnstileinfo_t:
// 26 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 26 into.
type Thread_turnstileinfo_t [26]byte

// See: https://developer.apple.com/documentation/kernel/thread_turnstileinfo_v2_t
// Thread_turnstileinfo_v2_t is opaque storage with the size and alignment C gives thread_turnstileinfo_v2_t:
// 28 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 28 into.
type Thread_turnstileinfo_v2_t [28]byte

// See: https://developer.apple.com/documentation/kernel/thread_waitinfo_t
// Thread_waitinfo_t is opaque storage with the size and alignment C gives thread_waitinfo_t:
// 25 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 25 into.
type Thread_waitinfo_t [25]byte

// See: https://developer.apple.com/documentation/kernel/thread_waitinfo_v2_t
// Thread_waitinfo_v2_t is opaque storage with the size and alignment C gives thread_waitinfo_v2_t:
// 31 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 31 into.
type Thread_waitinfo_v2_t [31]byte

// See: https://developer.apple.com/documentation/kernel/throttle_info_handle_t
type Throttle_info_handle_t = uintptr

// See: https://developer.apple.com/documentation/kernel/time_t
type Time_t = int64

// See: https://developer.apple.com/documentation/kernel/time_value_t
// Time_value_t is opaque storage with the size and alignment C gives time_value_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Time_value_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/token_t
// Token_t is an unresolved C aggregate typedef.
type Token_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/trap_gate_t
// Trap_gate_t is opaque storage with the size and alignment C gives trap_gate_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Trap_gate_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/tss_desc_t
// Tss_desc_t is opaque storage with the size and alignment C gives tss_desc_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Tss_desc_t [4]uint16

// See: https://developer.apple.com/documentation/kernel/tss_t
// Tss_t is opaque storage with the size and alignment C gives tss_t:
// 104 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 104 into.
type Tss_t [26]uint32

// See: https://developer.apple.com/documentation/kernel/uint-4rm
type UInt = uint32

// See: https://developer.apple.com/documentation/kernel/uintf
type UIntf = uint32

// See: https://developer.apple.com/documentation/kernel/ulong
type ULong = uint

// See: https://developer.apple.com/documentation/kernel/ulongf
type ULongf = uint

// See: https://developer.apple.com/documentation/kernel/u_char
type U_char = byte

// See: https://developer.apple.com/documentation/kernel/u_int
type U_int = uint32

// See: https://developer.apple.com/documentation/kernel/u_int16_t
type U_int16_t = uint16

// See: https://developer.apple.com/documentation/kernel/u_int32_t
type U_int32_t = uint32

// See: https://developer.apple.com/documentation/kernel/u_int64_t
type U_int64_t = uint64

// See: https://developer.apple.com/documentation/kernel/u_int8_t
type U_int8_t = byte

// See: https://developer.apple.com/documentation/kernel/u_long
type U_long = uint

// See: https://developer.apple.com/documentation/kernel/u_quad_t
type U_quad_t = uint64

// See: https://developer.apple.com/documentation/kernel/u_short
type U_short = uint16

// See: https://developer.apple.com/documentation/kernel/ucontext64_t
// Ucontext64_t is opaque storage with the size and alignment C gives ucontext64_t:
// 56 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 56 into.
type Ucontext64_t [7]uint64

// See: https://developer.apple.com/documentation/kernel/ucontext_t
// Ucontext_t is opaque storage with the size and alignment C gives ucontext_t:
// 56 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 56 into.
type Ucontext_t [7]uint64

// See: https://developer.apple.com/documentation/kernel/uext_object_t
type Uext_object_t = uint32

// See: https://developer.apple.com/documentation/kernel/uid_t
type Uid_t = uint32

// See: https://developer.apple.com/documentation/kernel/uint
type Uint = uint32

// See: https://developer.apple.com/documentation/kernel/uint16_t
type Uint16_t = uint16

// See: https://developer.apple.com/documentation/kernel/uint32_t
type Uint32_t = uint32

// See: https://developer.apple.com/documentation/kernel/uint64_t
type Uint64_t = uint64

// See: https://developer.apple.com/documentation/kernel/uint8_t
type Uint8_t = uint8

// See: https://developer.apple.com/documentation/kernel/uint_fast16_t
type Uint_fast16_t = uint16

// See: https://developer.apple.com/documentation/kernel/uint_fast32_t
type Uint_fast32_t = uint32

// See: https://developer.apple.com/documentation/kernel/uint_fast64_t
type Uint_fast64_t = uint64

// See: https://developer.apple.com/documentation/kernel/uint_fast8_t
type Uint_fast8_t = byte

// See: https://developer.apple.com/documentation/kernel/uint_least16_t
type Uint_least16_t = uint16

// See: https://developer.apple.com/documentation/kernel/uint_least32_t
type Uint_least32_t = uint32

// See: https://developer.apple.com/documentation/kernel/uint_least64_t
type Uint_least64_t = uint64

// See: https://developer.apple.com/documentation/kernel/uint_least8_t
type Uint_least8_t = byte

// See: https://developer.apple.com/documentation/kernel/uintmax_t
type Uintmax_t = uint

// See: https://developer.apple.com/documentation/kernel/uintptr_t
type Uintptr_t = uint

// See: https://developer.apple.com/documentation/kernel/uio_bptr_t
type Uio_bptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/uio_ptr_ref_t
type Uio_ptr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/uio_ptr_t
type Uio_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/uio_ref_ptr_t
type Uio_ref_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/uio_ref_ref_t
type Uio_ref_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/uio_ref_t
type Uio_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/uio_t
type Uio_t = uintptr

// See: https://developer.apple.com/documentation/kernel/unp_gen_t
type Unp_gen_t = uint64

// See: https://developer.apple.com/documentation/kernel/upl_control_flags_t
type Upl_control_flags_t = uint64

// See: https://developer.apple.com/documentation/kernel/upl_offset_t
type Upl_offset_t = uint32

// See: https://developer.apple.com/documentation/kernel/upl_page_info_array_t
type Upl_page_info_array_t = *Upl_page_info_t

// See: https://developer.apple.com/documentation/kernel/upl_page_info_t
// Upl_page_info_t is opaque storage with the size and alignment C gives upl_page_info_t:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type Upl_page_info_t [2]uint32

// See: https://developer.apple.com/documentation/kernel/upl_page_list_ptr_t
type Upl_page_list_ptr_t = Upl_page_info_array_t

// See: https://developer.apple.com/documentation/kernel/upl_size_t
type Upl_size_t = uint32

// See: https://developer.apple.com/documentation/kernel/upl_t
type Upl_t = uint32

// See: https://developer.apple.com/documentation/kernel/useconds_t
type Useconds_t = uint32

// See: https://developer.apple.com/documentation/kernel/user32_addr_t
type User32_addr_t = uint32

// See: https://developer.apple.com/documentation/kernel/user32_fchecklv_t
// User32_fchecklv_t is opaque storage with the size and alignment C gives user32_fchecklv_t:
// 1 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 1 into.
type User32_fchecklv_t [1]byte

// See: https://developer.apple.com/documentation/kernel/user32_fsignatures_t
// User32_fsignatures_t is opaque storage with the size and alignment C gives user32_fsignatures_t:
// 1 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 1 into.
type User32_fsignatures_t [1]byte

// See: https://developer.apple.com/documentation/kernel/user32_long_t
type User32_long_t = int32

// See: https://developer.apple.com/documentation/kernel/user32_msglen_t
type User32_msglen_t = uint32

// See: https://developer.apple.com/documentation/kernel/user32_msgqnum_t
type User32_msgqnum_t = uint32

// See: https://developer.apple.com/documentation/kernel/user32_off_t
type User32_off_t = int64

// See: https://developer.apple.com/documentation/kernel/user32_size_t
type User32_size_t = uint32

// See: https://developer.apple.com/documentation/kernel/user32_ssize_t
type User32_ssize_t = int32

// See: https://developer.apple.com/documentation/kernel/user32_time_t
type User32_time_t = int32

// See: https://developer.apple.com/documentation/kernel/user32_ulong_t
type User32_ulong_t = uint32

// See: https://developer.apple.com/documentation/kernel/user64_addr_t
type User64_addr_t = uint64

// See: https://developer.apple.com/documentation/kernel/user64_long_t
type User64_long_t = int64

// See: https://developer.apple.com/documentation/kernel/user64_msglen_t
type User64_msglen_t = uint64

// See: https://developer.apple.com/documentation/kernel/user64_msgqnum_t
type User64_msgqnum_t = uint64

// See: https://developer.apple.com/documentation/kernel/user64_off_t
type User64_off_t = int64

// See: https://developer.apple.com/documentation/kernel/user64_size_t
type User64_size_t = uint64

// See: https://developer.apple.com/documentation/kernel/user64_ssize_t
type User64_ssize_t = int64

// See: https://developer.apple.com/documentation/kernel/user64_time_t
type User64_time_t = int64

// See: https://developer.apple.com/documentation/kernel/user64_ulong_t
type User64_ulong_t = uint64

// See: https://developer.apple.com/documentation/kernel/user_addr_t
type User_addr_t = uint64

// See: https://developer.apple.com/documentation/kernel/user_addr_ut
type User_addr_ut = uint64

// See: https://developer.apple.com/documentation/kernel/user_fchecklv_t
// User_fchecklv_t is opaque storage with the size and alignment C gives user_fchecklv_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type User_fchecklv_t [3]uint64

// See: https://developer.apple.com/documentation/kernel/user_fsignatures_t
// User_fsignatures_t is opaque storage with the size and alignment C gives user_fsignatures_t:
// 56 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 56 into.
type User_fsignatures_t [7]uint64

// See: https://developer.apple.com/documentation/kernel/user_fsupplement_t
// User_fsupplement_t is opaque storage with the size and alignment C gives user_fsupplement_t:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type User_fsupplement_t [4]uint64

// See: https://developer.apple.com/documentation/kernel/user_long_t
type User_long_t = int64

// See: https://developer.apple.com/documentation/kernel/user_msglen_t
type User_msglen_t = uint64

// See: https://developer.apple.com/documentation/kernel/user_msgqnum_t
type User_msgqnum_t = uint64

// See: https://developer.apple.com/documentation/kernel/user_off_t
type User_off_t = int64

// See: https://developer.apple.com/documentation/kernel/user_size_t
type User_size_t = uint64

// See: https://developer.apple.com/documentation/kernel/user_size_ut
type User_size_ut = uint64

// See: https://developer.apple.com/documentation/kernel/user_speed_t
type User_speed_t = uint64

// See: https://developer.apple.com/documentation/kernel/user_ssize_t
type User_ssize_t = int64

// See: https://developer.apple.com/documentation/kernel/user_subsystem_t
type User_subsystem_t = *byte

// See: https://developer.apple.com/documentation/kernel/user_tcflag_t
type User_tcflag_t = uint64

// See: https://developer.apple.com/documentation/kernel/user_time_t
type User_time_t = int64

// See: https://developer.apple.com/documentation/kernel/user_ulong_t
type User_ulong_t = uint64

// See: https://developer.apple.com/documentation/kernel/ushort
type Ushort = uint16

// See: https://developer.apple.com/documentation/kernel/uuid_string_t
// Uuid_string_t is opaque storage with the size and alignment C gives uuid_string_t:
// 37 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 37 into.
type Uuid_string_t [37]byte

// See: https://developer.apple.com/documentation/kernel/uuid_t
type Uuid_t = [16]byte

// VDSP_Length is used for numbers of elements in arrays and indices of elements in arrays. It is also used for the base-two logarithm of numbers of elements.
//
// See: https://developer.apple.com/documentation/kernel/vdsp_length
type VDSP_Length = uint

// VDSP_Stride is used to hold differences between indices of elements, including the lengths of strides.
//
// See: https://developer.apple.com/documentation/kernel/vdsp_stride
type VDSP_Stride = int

// See: https://developer.apple.com/documentation/kernel/vdsp_biquad_setup
type VDSP_biquad_Setup = uintptr

// See: https://developer.apple.com/documentation/kernel/vdsp_biquad_setupd
type VDSP_biquad_SetupD = uintptr

// See: https://developer.apple.com/documentation/kernel/vdsp_biquadm_setup
type VDSP_biquadm_Setup = uintptr

// See: https://developer.apple.com/documentation/kernel/vdsp_biquadm_setupd
type VDSP_biquadm_SetupD = uintptr

// VDSP_int24 is a data structure that holds a 24-bit signed integer value.
//
// See: https://developer.apple.com/documentation/kernel/vdsp_int24
// VDSP_int24 is opaque storage with the size and alignment C gives vDSP_int24:
// 3 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 3 into.
type VDSP_int24 [3]byte

// See: https://developer.apple.com/documentation/kernel/vdsp_uint24
// VDSP_uint24 is opaque storage with the size and alignment C gives vDSP_uint24:
// 3 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 3 into.
type VDSP_uint24 [3]byte

// VDouble is a 128-bit vector packed with `double` values.
//
// See: https://developer.apple.com/documentation/Accelerate/vDouble
type VDouble = float64

// See: https://developer.apple.com/documentation/kernel/va_list
type Va_list = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vc_progress_user_options
// Vc_progress_user_options is opaque storage with the size and alignment C gives vc_progress_user_options:
// 44 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 44 into.
type Vc_progress_user_options [11]uint32

// See: https://developer.apple.com/documentation/kernel/vector_int2
type Vector_int2 = [2]int32

// See: https://developer.apple.com/documentation/kernel/vector_int4
type Vector_int4 = [4]int32

// See: https://developer.apple.com/documentation/kernel/vector_int8
type Vector_int8 = int32

// See: https://developer.apple.com/documentation/kernel/vector_uchar16
type Vector_uchar16 = uint8

// See: https://developer.apple.com/documentation/kernel/vector_uchar32
type Vector_uchar32 = uint8

// See: https://developer.apple.com/documentation/kernel/vector_uchar64
type Vector_uchar64 = uint8

// See: https://developer.apple.com/documentation/kernel/vector_uchar8
type Vector_uchar8 = uint8

// See: https://developer.apple.com/documentation/kernel/vector_uint4
type Vector_uint4 = [4]uint32

// See: https://developer.apple.com/documentation/kernel/vector_ushort4
type Vector_ushort4 = uint16

// See: https://developer.apple.com/documentation/kernel/vfs_context_bptr_t
type Vfs_context_bptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vfs_context_ptr_ref_t
type Vfs_context_ptr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vfs_context_ptr_t
type Vfs_context_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vfs_context_ref_ptr_t
type Vfs_context_ref_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vfs_context_ref_ref_t
type Vfs_context_ref_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vfs_context_ref_t
type Vfs_context_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vfs_context_t
type Vfs_context_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vfs_path_t
type Vfs_path_t = int8

// See: https://developer.apple.com/documentation/kernel/vfs_rename_flags_t
type Vfs_rename_flags_t = uint32

// See: https://developer.apple.com/documentation/kernel/vfs_roles_t
type Vfs_roles_t = uint32

// See: https://developer.apple.com/documentation/kernel/vfstable_bptr_t
type Vfstable_bptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vfstable_ptr_ref_t
type Vfstable_ptr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vfstable_ptr_t
type Vfstable_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vfstable_ref_ptr_t
type Vfstable_ref_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vfstable_ref_ref_t
type Vfstable_ref_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vfstable_ref_t
type Vfstable_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vfstable_t
type Vfstable_t = uintptr

// See: https://developer.apple.com/documentation/kernel/virtual_memory_guard_exception_code_t
type Virtual_memory_guard_exception_code_t = VirtualMemoryGuardExceptionCode

// See: https://developer.apple.com/documentation/kernel/vm32_addr_struct_t
type Vm32_addr_struct_t = uint32

// See: https://developer.apple.com/documentation/kernel/vm32_address_t
type Vm32_address_t = uint32

// See: https://developer.apple.com/documentation/kernel/vm32_object_id_t
type Vm32_object_id_t = uint32

// See: https://developer.apple.com/documentation/kernel/vm32_offset_t
type Vm32_offset_t = uint32

// See: https://developer.apple.com/documentation/kernel/vm32_size_struct_t
type Vm32_size_struct_t = uint32

// See: https://developer.apple.com/documentation/kernel/vm32_size_t
type Vm32_size_t = uint32

// See: https://developer.apple.com/documentation/kernel/vm_addr_struct_t
type Vm_addr_struct_t = uint64

// See: https://developer.apple.com/documentation/kernel/vm_address_t
type Vm_address_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_address_ut
type Vm_address_ut = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_behavior_t
type Vm_behavior_t = int32

// See: https://developer.apple.com/documentation/kernel/vm_behavior_ut
type Vm_behavior_ut = int32

// See: https://developer.apple.com/documentation/kernel/vm_extmod_statistics_data_t
// Vm_extmod_statistics_data_t is opaque storage with the size and alignment C gives vm_extmod_statistics_data_t:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type Vm_extmod_statistics_data_t [6]uint64

// See: https://developer.apple.com/documentation/kernel/vm_extmod_statistics_t
type Vm_extmod_statistics_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_info_object_array_t
type Vm_info_object_array_t = *Vm_info_object_t

// See: https://developer.apple.com/documentation/kernel/vm_info_object_t
// Vm_info_object_t is opaque storage with the size and alignment C gives vm_info_object_t:
// 88 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 88 into.
type Vm_info_object_t [22]uint32

// See: https://developer.apple.com/documentation/kernel/vm_info_region_64_t
// Vm_info_region_64_t is opaque storage with the size and alignment C gives vm_info_region_64_t:
// 44 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 44 into.
type Vm_info_region_64_t [11]uint32

// See: https://developer.apple.com/documentation/kernel/vm_info_region_t
// Vm_info_region_t is opaque storage with the size and alignment C gives vm_info_region_t:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type Vm_info_region_t [10]uint32

// See: https://developer.apple.com/documentation/kernel/vm_inherit_t
type Vm_inherit_t = uint32

// See: https://developer.apple.com/documentation/kernel/vm_inherit_ut
type Vm_inherit_ut = uint32

// See: https://developer.apple.com/documentation/kernel/vm_machine_attribute_t
type Vm_machine_attribute_t = uint32

// See: https://developer.apple.com/documentation/kernel/vm_machine_attribute_val_t
type Vm_machine_attribute_val_t = int32

// See: https://developer.apple.com/documentation/kernel/vm_map_address_t
type Vm_map_address_t = uint64

// See: https://developer.apple.com/documentation/kernel/vm_map_address_ut
type Vm_map_address_ut = uint64

// See: https://developer.apple.com/documentation/kernel/vm_map_inspect_t
type Vm_map_inspect_t = uint32

// See: https://developer.apple.com/documentation/kernel/vm_map_offset_t
type Vm_map_offset_t = uint64

// See: https://developer.apple.com/documentation/kernel/vm_map_offset_ut
type Vm_map_offset_ut = uint64

// See: https://developer.apple.com/documentation/kernel/vm_map_read_t
type Vm_map_read_t = uint32

// See: https://developer.apple.com/documentation/kernel/vm_map_size_t
type Vm_map_size_t = uint64

// See: https://developer.apple.com/documentation/kernel/vm_map_size_ut
type Vm_map_size_ut = uint64

// See: https://developer.apple.com/documentation/kernel/vm_map_t
type Vm_map_t = uint32

// See: https://developer.apple.com/documentation/kernel/vm_named_entry_t
type Vm_named_entry_t = uint32

// See: https://developer.apple.com/documentation/kernel/vm_object_id_t
type Vm_object_id_t = uint64

// See: https://developer.apple.com/documentation/kernel/vm_object_offset_t
type Vm_object_offset_t = uint64

// See: https://developer.apple.com/documentation/kernel/vm_object_offset_ut
type Vm_object_offset_ut = uint64

// See: https://developer.apple.com/documentation/kernel/vm_object_size_t
type Vm_object_size_t = uint64

// See: https://developer.apple.com/documentation/kernel/vm_object_size_ut
type Vm_object_size_ut = uint64

// See: https://developer.apple.com/documentation/kernel/vm_offset_t
type Vm_offset_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_offset_ut
type Vm_offset_ut = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_page_info_basic_data_t
// Vm_page_info_basic_data_t is opaque storage with the size and alignment C gives vm_page_info_basic_data_t:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type Vm_page_info_basic_data_t [4]uint64

// See: https://developer.apple.com/documentation/kernel/vm_page_info_basic_t
type Vm_page_info_basic_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_page_info_data_t
type Vm_page_info_data_t = int32

// See: https://developer.apple.com/documentation/kernel/vm_page_info_flavor_t
type Vm_page_info_flavor_t = int32

// See: https://developer.apple.com/documentation/kernel/vm_page_info_t
type Vm_page_info_t = *int32

// See: https://developer.apple.com/documentation/kernel/vm_prot_t
type Vm_prot_t = int32

// See: https://developer.apple.com/documentation/kernel/vm_prot_ut
type Vm_prot_ut = int32

// See: https://developer.apple.com/documentation/kernel/vm_purgable_t
type Vm_purgable_t = int32

// See: https://developer.apple.com/documentation/kernel/vm_purgeable_info_t
type Vm_purgeable_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_purgeable_stat_t
// Vm_purgeable_stat_t is opaque storage with the size and alignment C gives vm_purgeable_stat_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Vm_purgeable_stat_t [2]uint64

// See: https://developer.apple.com/documentation/kernel/vm_read_entry_t
// Vm_read_entry_t is opaque storage with the size and alignment C gives vm_read_entry_t:
// 4096 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4096 into.
type Vm_read_entry_t [1024]uint32

// See: https://developer.apple.com/documentation/kernel/vm_region_basic_info_64_t
type Vm_region_basic_info_64_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_region_basic_info_data_64_t
// Vm_region_basic_info_data_64_t is opaque storage with the size and alignment C gives vm_region_basic_info_data_64_t:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type Vm_region_basic_info_data_64_t [9]uint32

// See: https://developer.apple.com/documentation/kernel/vm_region_basic_info_data_t
// Vm_region_basic_info_data_t is opaque storage with the size and alignment C gives vm_region_basic_info_data_t:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type Vm_region_basic_info_data_t [8]uint32

// See: https://developer.apple.com/documentation/kernel/vm_region_basic_info_t
type Vm_region_basic_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_region_extended_info_data_t
// Vm_region_extended_info_data_t is opaque storage with the size and alignment C gives vm_region_extended_info_data_t:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type Vm_region_extended_info_data_t [9]uint32

// See: https://developer.apple.com/documentation/kernel/vm_region_extended_info_t
type Vm_region_extended_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_region_flavor_t
type Vm_region_flavor_t = int32

// See: https://developer.apple.com/documentation/kernel/vm_region_info_64_t
type Vm_region_info_64_t = *int32

// See: https://developer.apple.com/documentation/kernel/vm_region_info_data_t
type Vm_region_info_data_t = int32

// See: https://developer.apple.com/documentation/kernel/vm_region_info_t
type Vm_region_info_t = *int32

// See: https://developer.apple.com/documentation/kernel/vm_region_recurse_info_64_t
type Vm_region_recurse_info_64_t = *int32

// See: https://developer.apple.com/documentation/kernel/vm_region_recurse_info_t
type Vm_region_recurse_info_t = *int32

// See: https://developer.apple.com/documentation/kernel/vm_region_submap_info_64_t
type Vm_region_submap_info_64_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_region_submap_info_data_64_t
// Vm_region_submap_info_data_64_t is opaque storage with the size and alignment C gives vm_region_submap_info_data_64_t:
// 76 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 76 into.
type Vm_region_submap_info_data_64_t [19]uint32

// See: https://developer.apple.com/documentation/kernel/vm_region_submap_info_data_t
// Vm_region_submap_info_data_t is opaque storage with the size and alignment C gives vm_region_submap_info_data_t:
// 60 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 60 into.
type Vm_region_submap_info_data_t [15]uint32

// See: https://developer.apple.com/documentation/kernel/vm_region_submap_info_t
type Vm_region_submap_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_region_submap_short_info_64_t
type Vm_region_submap_short_info_64_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_region_submap_short_info_data_64_t
// Vm_region_submap_short_info_data_64_t is opaque storage with the size and alignment C gives vm_region_submap_short_info_data_64_t:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type Vm_region_submap_short_info_data_64_t [12]uint32

// See: https://developer.apple.com/documentation/kernel/vm_region_top_info_data_t
// Vm_region_top_info_data_t is opaque storage with the size and alignment C gives vm_region_top_info_data_t:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type Vm_region_top_info_data_t [5]uint32

// See: https://developer.apple.com/documentation/kernel/vm_region_top_info_t
type Vm_region_top_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_size_struct_t
type Vm_size_struct_t = uint64

// See: https://developer.apple.com/documentation/kernel/vm_size_t
type Vm_size_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_size_ut
type Vm_size_ut = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_statistics64_data_t
// Vm_statistics64_data_t is opaque storage with the size and alignment C gives vm_statistics64_data_t:
// 248 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 248 into.
type Vm_statistics64_data_t [31]uint64

// See: https://developer.apple.com/documentation/kernel/vm_statistics64_t
type Vm_statistics64_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_statistics_data_t
// Vm_statistics_data_t is opaque storage with the size and alignment C gives vm_statistics_data_t:
// 60 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 60 into.
type Vm_statistics_data_t [15]uint32

// See: https://developer.apple.com/documentation/kernel/vm_statistics_t
type Vm_statistics_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_sync_t
type Vm_sync_t = uint32

// See: https://developer.apple.com/documentation/kernel/vm_task_entry_t
type Vm_task_entry_t = uint32

// See: https://developer.apple.com/documentation/kernel/vnode_bptr_t
type Vnode_bptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vnode_ptr_ref_t
type Vnode_ptr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vnode_ptr_t
type Vnode_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vnode_ref_ptr_t
type Vnode_ref_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vnode_ref_ref_t
type Vnode_ref_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vnode_ref_t
type Vnode_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vnode_t
type Vnode_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vnode_verify_flags_t
type Vnode_verify_flags_t = uint32

// See: https://developer.apple.com/documentation/kernel/voidp
type Voidp = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/voidpc
type Voidpc = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/voidpf
type Voidpf = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vol_attributes_attr_t
// Vol_attributes_attr_t is opaque storage with the size and alignment C gives vol_attributes_attr_t:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type Vol_attributes_attr_t [10]uint32

// See: https://developer.apple.com/documentation/kernel/vol_capabilities_attr_t
// Vol_capabilities_attr_t is opaque storage with the size and alignment C gives vol_capabilities_attr_t:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type Vol_capabilities_attr_t [8]uint32

// See: https://developer.apple.com/documentation/kernel/vol_capabilities_set_t
type Vol_capabilities_set_t = uint32

// See: https://developer.apple.com/documentation/kernel/vsock_gen_t
type Vsock_gen_t = uint64

// See: https://developer.apple.com/documentation/kernel/wait_interrupt_t
type Wait_interrupt_t = int32

// See: https://developer.apple.com/documentation/kernel/wait_result_t
type Wait_result_t = int32

// See: https://developer.apple.com/documentation/kernel/wait_timeout_urgency_t
type Wait_timeout_urgency_t = int32

// See: https://developer.apple.com/documentation/kernel/wint_t
type Wint_t = int32

// See: https://developer.apple.com/documentation/kernel/x86_avx512_state_t
// X86_avx512_state_t is an unresolved C aggregate typedef.
type X86_avx512_state_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/x86_avx_state_t
// X86_avx_state_t is an unresolved C aggregate typedef.
type X86_avx_state_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/x86_debug_state_t
// X86_debug_state_t is an unresolved C aggregate typedef.
type X86_debug_state_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/x86_exception_state32_t
// X86_exception_state32_t is an unresolved C aggregate typedef.
type X86_exception_state32_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/x86_exception_state_t
// X86_exception_state_t is an unresolved C aggregate typedef.
type X86_exception_state_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/x86_float_state32_t
// X86_float_state32_t is an unresolved C aggregate typedef.
type X86_float_state32_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/x86_float_state_t
// X86_float_state_t is an unresolved C aggregate typedef.
type X86_float_state_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/x86_state_hdr_t
// X86_state_hdr_t is an unresolved C aggregate typedef.
type X86_state_hdr_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/x86_thread_state32_t
// X86_thread_state32_t is an unresolved C aggregate typedef.
type X86_thread_state32_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/x86_thread_state_t
// X86_thread_state_t is an unresolved C aggregate typedef.
type X86_thread_state_t unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/xattrname
type Xattrname = int8

// See: https://developer.apple.com/documentation/kernel/xcred
type Xcred = uint32

// See: https://developer.apple.com/documentation/kernel/xdrbuf_type
type Xdrbuf_type = uint32

// See: https://developer.apple.com/documentation/kernel/xmldata_t
type XmlData_t = *byte

// See: https://developer.apple.com/documentation/kernel/z_stream
// Z_stream is opaque storage with the size and alignment C gives z_stream:
// 112 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 112 into.
type Z_stream [14]uint64

// See: https://developer.apple.com/documentation/kernel/z_streamp
type Z_streamp = *Z_stream

// See: https://developer.apple.com/documentation/kernel/zone_btrecord_array_t
type Zone_btrecord_array_t = *Zone_btrecord_t

// See: https://developer.apple.com/documentation/kernel/zone_btrecord_t
// Zone_btrecord_t is opaque storage with the size and alignment C gives zone_btrecord_t:
// 128 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 128 into.
type Zone_btrecord_t [16]uint64

// See: https://developer.apple.com/documentation/kernel/zone_info_array_t
type Zone_info_array_t = *Zone_info_t

// See: https://developer.apple.com/documentation/kernel/zone_info_t
// Zone_info_t is opaque storage with the size and alignment C gives zone_info_t:
// 56 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 56 into.
type Zone_info_t [7]uint64

// See: https://developer.apple.com/documentation/kernel/zone_name_array_t
type Zone_name_array_t = *Zone_name_t

// See: https://developer.apple.com/documentation/kernel/zone_name_t
// Zone_name_t is opaque storage with the size and alignment C gives zone_name_t:
// 80 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 80 into.
type Zone_name_t [80]byte

// BootVideo is a Go-name alias for Boot_Video.
type BootVideo = Boot_Video

// BootVideoV1 is a Go-name alias for Boot_VideoV1.
type BootVideoV1 = Boot_VideoV1

// ComplexSplit is a Go-name alias for COMPLEX_SPLIT.
type ComplexSplit = COMPLEX_SPLIT

// CsBlobIndex is a Go-name alias for CS_BlobIndex.
type CsBlobIndex = CS_BlobIndex

// CsCodeDirectory is a Go-name alias for CS_CodeDirectory.
type CsCodeDirectory = CS_CodeDirectory

// CsGenericBlob is a Go-name alias for CS_GenericBlob.
type CsGenericBlob = CS_GenericBlob

// CsSuperBlob is a Go-name alias for CS_SuperBlob.
type CsSuperBlob = CS_SuperBlob

// DoubleComplex is a Go-name alias for DOUBLE_COMPLEX.
type DoubleComplex = DOUBLE_COMPLEX

// DoubleComplexSplit is a Go-name alias for DOUBLE_COMPLEX_SPLIT.
type DoubleComplexSplit = DOUBLE_COMPLEX_SPLIT

// EfiBoolean is a Go-name alias for EFI_BOOLEAN.
type EfiBoolean = EFI_BOOLEAN

// EfiChar16 is a Go-name alias for EFI_CHAR16.
type EfiChar16 = EFI_CHAR16

// EfiChar32 is a Go-name alias for EFI_CHAR32.
type EfiChar32 = EFI_CHAR32

// EfiChar64 is a Go-name alias for EFI_CHAR64.
type EfiChar64 = EFI_CHAR64

// EfiChar8 is a Go-name alias for EFI_CHAR8.
type EfiChar8 = EFI_CHAR8

// EfiConfigurationTable32 is a Go-name alias for EFI_CONFIGURATION_TABLE_32.
type EfiConfigurationTable32 = EFI_CONFIGURATION_TABLE_32

// EfiConfigurationTable64 is a Go-name alias for EFI_CONFIGURATION_TABLE_64.
type EfiConfigurationTable64 = EFI_CONFIGURATION_TABLE_64

// EfiGuid is a Go-name alias for EFI_GUID.
type EfiGuid = EFI_GUID

// EfiHandle32 is a Go-name alias for EFI_HANDLE32.
type EfiHandle32 = EFI_HANDLE32

// EfiHandle64 is a Go-name alias for EFI_HANDLE64.
type EfiHandle64 = EFI_HANDLE64

// EfiInt16 is a Go-name alias for EFI_INT16.
type EfiInt16 = EFI_INT16

// EfiInt32 is a Go-name alias for EFI_INT32.
type EfiInt32 = EFI_INT32

// EfiInt64 is a Go-name alias for EFI_INT64.
type EfiInt64 = EFI_INT64

// EfiInt8 is a Go-name alias for EFI_INT8.
type EfiInt8 = EFI_INT8

// EfiMemoryDescriptor is a Go-name alias for EFI_MEMORY_DESCRIPTOR.
type EfiMemoryDescriptor = EFI_MEMORY_DESCRIPTOR

// EfiMemoryType is a Go-name alias for EFI_MEMORY_TYPE.
type EfiMemoryType = EFI_MEMORY_TYPE

// EfiPhysicalAddress is a Go-name alias for EFI_PHYSICAL_ADDRESS.
type EfiPhysicalAddress = EFI_PHYSICAL_ADDRESS

// EfiPtr32 is a Go-name alias for EFI_PTR32.
type EfiPtr32 = EFI_PTR32

// EfiPtr64 is a Go-name alias for EFI_PTR64.
type EfiPtr64 = EFI_PTR64

// EfiResetType is a Go-name alias for EFI_RESET_TYPE.
type EfiResetType = EFI_RESET_TYPE

// EfiRuntimeServices32 is a Go-name alias for EFI_RUNTIME_SERVICES_32.
type EfiRuntimeServices32 = EFI_RUNTIME_SERVICES_32

// EfiRuntimeServices64 is a Go-name alias for EFI_RUNTIME_SERVICES_64.
type EfiRuntimeServices64 = EFI_RUNTIME_SERVICES_64

// EfiStatus is a Go-name alias for EFI_STATUS.
type EfiStatus = EFI_STATUS

// EfiSystemTable32 is a Go-name alias for EFI_SYSTEM_TABLE_32.
type EfiSystemTable32 = EFI_SYSTEM_TABLE_32

// EfiSystemTable64 is a Go-name alias for EFI_SYSTEM_TABLE_64.
type EfiSystemTable64 = EFI_SYSTEM_TABLE_64

// EfiTableHeader is a Go-name alias for EFI_TABLE_HEADER.
type EfiTableHeader = EFI_TABLE_HEADER

// EfiTime is a Go-name alias for EFI_TIME.
type EfiTime = EFI_TIME

// EfiTimeCapabilities is a Go-name alias for EFI_TIME_CAPABILITIES.
type EfiTimeCapabilities = EFI_TIME_CAPABILITIES

// EfiUint16 is a Go-name alias for EFI_UINT16.
type EfiUint16 = EFI_UINT16

// EfiUint32 is a Go-name alias for EFI_UINT32.
type EfiUint32 = EFI_UINT32

// EfiUint64 is a Go-name alias for EFI_UINT64.
type EfiUint64 = EFI_UINT64

// EfiUint8 is a Go-name alias for EFI_UINT8.
type EfiUint8 = EFI_UINT8

// EfiUintn is a Go-name alias for EFI_UINTN.
type EfiUintn = EFI_UINTN

// EfiVirtualAddress is a Go-name alias for EFI_VIRTUAL_ADDRESS.
type EfiVirtualAddress = EFI_VIRTUAL_ADDRESS

// IOPCIDeviceCrashNotification is a Go-name alias for IOPCIDeviceCrashNotification_t.
type IOPCIDeviceCrashNotification = IOPCIDeviceCrashNotification_t

// Md5Ctx is a Go-name alias for MD5_CTX.
type Md5Ctx = MD5_CTX

// NdrRecord is a Go-name alias for NDR_record_t.
type NdrRecord = NDR_record_t

// PeVideo is a Go-name alias for PE_Video.
type PeVideo = PE_Video

// PeState is a Go-name alias for PE_state_t.
type PeState = PE_state_t

// RawHeader is a Go-name alias for RAW_header.
type RawHeader = RAW_header

// ReportLunsLogicalUnitAddressing is a Go-name alias for REPORT_LUNS_LOGICAL_UNIT_ADDRESSING.
type ReportLunsLogicalUnitAddressing = REPORT_LUNS_LOGICAL_UNIT_ADDRESSING

// ReportLunsPeripheralDeviceAddressing is a Go-name alias for REPORT_LUNS_PERIPHERAL_DEVICE_ADDRESSING.
type ReportLunsPeripheralDeviceAddressing = REPORT_LUNS_PERIPHERAL_DEVICE_ADDRESSING

// SCSICmdInquiryPAGECxHeader is a Go-name alias for SCSICmd_INQUIRY_PAGECx_Header.
type SCSICmdInquiryPAGECxHeader = SCSICmd_INQUIRY_PAGECx_Header

// SCSICmdInquiryPage00HeaderSpc16 is a Go-name alias for SCSICmd_INQUIRY_Page00_Header_SPC_16.
type SCSICmdInquiryPage00HeaderSpc16 = SCSICmd_INQUIRY_Page00_Header_SPC_16

// SCSICmdInquiryPage80HeaderSpc16 is a Go-name alias for SCSICmd_INQUIRY_Page80_Header_SPC_16.
type SCSICmdInquiryPage80HeaderSpc16 = SCSICmd_INQUIRY_Page80_Header_SPC_16

// SCSICmdInquiryPageB0Data is a Go-name alias for SCSICmd_INQUIRY_PageB0_Data.
type SCSICmdInquiryPageB0Data = SCSICmd_INQUIRY_PageB0_Data

// SCSICmdInquiryPageB2Data is a Go-name alias for SCSICmd_INQUIRY_PageB2_Data.
type SCSICmdInquiryPageB2Data = SCSICmd_INQUIRY_PageB2_Data

// SCSICmdInquiryPageB2ProvisioningGroupDescriptor is a Go-name alias for SCSICmd_INQUIRY_PageB2_Provisioning_Group_Descriptor.
type SCSICmdInquiryPageB2ProvisioningGroupDescriptor = SCSICmd_INQUIRY_PageB2_Provisioning_Group_Descriptor

// SCSICmdInquiryPageC0Data is a Go-name alias for SCSICmd_INQUIRY_PageC0_Data.
type SCSICmdInquiryPageC0Data = SCSICmd_INQUIRY_PageC0_Data

// SCSICmdInquiryPageC1Data is a Go-name alias for SCSICmd_INQUIRY_PageC1_Data.
type SCSICmdInquiryPageC1Data = SCSICmd_INQUIRY_PageC1_Data

// SCSICmdInquiryStandardDataPtr is a Go-name alias for SCSICmd_INQUIRY_StandardDataPtr.
type SCSICmdInquiryStandardDataPtr = SCSICmd_INQUIRY_StandardDataPtr

// SCSICmdReportLunsHeader is a Go-name alias for SCSICmd_REPORT_LUNS_Header.
type SCSICmdReportLunsHeader = SCSICmd_REPORT_LUNS_Header

// SCSICmdReportLunsLunEntry is a Go-name alias for SCSICmd_REPORT_LUNS_LUN_ENTRY.
type SCSICmdReportLunsLunEntry = SCSICmd_REPORT_LUNS_LUN_ENTRY

// ScsiSenseData is a Go-name alias for SCSI_Sense_Data.
type ScsiSenseData = SCSI_Sense_Data

// ScScatter is a Go-name alias for SC_Scatter.
type ScScatter = SC_Scatter

// Sha1Ctx is a Go-name alias for SHA1_CTX.
type Sha1Ctx = SHA1_CTX

// StickyKeysModifierInfo is a Go-name alias for StickyKeys_ModifierInfo.
type StickyKeysModifierInfo = StickyKeys_ModifierInfo

// StickyKeysToggleInfo is a Go-name alias for StickyKeys_ToggleInfo.
type StickyKeysToggleInfo = StickyKeys_ToggleInfo

// WkWord is a Go-name alias for WK_word.
type WkWord = WK_word

// Addr64 is a Go-name alias for Addr64_t.
type Addr64 = Addr64_t

// Aid is a Go-name alias for Aid_t.
type Aid = Aid_t

// AlarmPort is a Go-name alias for Alarm_port_t.
type AlarmPort = Alarm_port_t

// Alarm is a Go-name alias for Alarm_t.
type Alarm = Alarm_t

// AlarmType is a Go-name alias for Alarm_type_t.
type AlarmType = Alarm_type_t

// ArcadeRegister is a Go-name alias for Arcade_register_t.
type ArcadeRegister = Arcade_register_t

// ArmDebugInfo is a Go-name alias for Arm_debug_info_t.
type ArmDebugInfo = Arm_debug_info_t

// ArmExceptionState32 is a Go-name alias for Arm_exception_state32_t.
type ArmExceptionState32 = Arm_exception_state32_t

// ArmFeatureBits is a Go-name alias for Arm_feature_bits_t.
type ArmFeatureBits = Arm_feature_bits_t

// ArmNeonState32 is a Go-name alias for Arm_neon_state32_t.
type ArmNeonState32 = Arm_neon_state32_t

// ArmStateHdr is a Go-name alias for Arm_state_hdr_t.
type ArmStateHdr = Arm_state_hdr_t

// ArmThreadState32 is a Go-name alias for Arm_thread_state32_t.
type ArmThreadState32 = Arm_thread_state32_t

// ArmUnifiedThreadState is a Go-name alias for Arm_unified_thread_state_t.
type ArmUnifiedThreadState = Arm_unified_thread_state_t

// AtmAction is a Go-name alias for Atm_action_t.
type AtmAction = Atm_action_t

// AtmAid is a Go-name alias for Atm_aid_t.
type AtmAid = Atm_aid_t

// AtmGuard is a Go-name alias for Atm_guard_t.
type AtmGuard = Atm_guard_t

// AtmMailboxOffset is a Go-name alias for Atm_mailbox_offset_t.
type AtmMailboxOffset = Atm_mailbox_offset_t

// AtmMemoryDescriptorArray is a Go-name alias for Atm_memory_descriptor_array_t.
type AtmMemoryDescriptorArray = Atm_memory_descriptor_array_t

// AtmMemoryDescriptor is a Go-name alias for Atm_memory_descriptor_t.
type AtmMemoryDescriptor = Atm_memory_descriptor_t

// AtmMemorySizeArray is a Go-name alias for Atm_memory_size_array_t.
type AtmMemorySizeArray = Atm_memory_size_array_t

// AtmSubaid32 is a Go-name alias for Atm_subaid32_t.
type AtmSubaid32 = Atm_subaid32_t

// Attrgroup is a Go-name alias for Attrgroup_t.
type Attrgroup = Attrgroup_t

// AttributeSet is a Go-name alias for Attribute_set_t.
type AttributeSet = Attribute_set_t

// Attrreference is a Go-name alias for Attrreference_t.
type Attrreference = Attrreference_t

// AuAsflgs is a Go-name alias for Au_asflgs_t.
type AuAsflgs = Au_asflgs_t

// AuAsid is a Go-name alias for Au_asid_t.
type AuAsid = Au_asid_t

// AuClass is a Go-name alias for Au_class_t.
type AuClass = Au_class_t

// AuCtlmode is a Go-name alias for Au_ctlmode_t.
type AuCtlmode = Au_ctlmode_t

// AuEmod is a Go-name alias for Au_emod_t.
type AuEmod = Au_emod_t

// AuEvclassMap is a Go-name alias for Au_evclass_map_t.
type AuEvclassMap = Au_evclass_map_t

// AuEvent is a Go-name alias for Au_event_t.
type AuEvent = Au_event_t

// AuExpireAfter is a Go-name alias for Au_expire_after_t.
type AuExpireAfter = Au_expire_after_t

// AuFstat is a Go-name alias for Au_fstat_t.
type AuFstat = Au_fstat_t

// AuID is a Go-name alias for Au_id_t.
type AuID = Au_id_t

// AuMask is a Go-name alias for Au_mask_t.
type AuMask = Au_mask_t

// AuQctrl is a Go-name alias for Au_qctrl_t.
type AuQctrl = Au_qctrl_t

// AuSession is a Go-name alias for Au_session_t.
type AuSession = Au_session_t

// AuStat is a Go-name alias for Au_stat_t.
type AuStat = Au_stat_t

// AuTidAddr is a Go-name alias for Au_tid_addr_t.
type AuTidAddr = Au_tid_addr_t

// AuTid is a Go-name alias for Au_tid_t.
type AuTid = Au_tid_t

// AuditToken is a Go-name alias for Audit_token_t.
type AuditToken = Audit_token_t

// AuditinfoAddr is a Go-name alias for Auditinfo_addr_t.
type AuditinfoAddr = Auditinfo_addr_t

// Auditinfo is a Go-name alias for Auditinfo_t.
type Auditinfo = Auditinfo_t

// AuditpinfoAddr is a Go-name alias for Auditpinfo_addr_t.
type AuditpinfoAddr = Auditpinfo_addr_t

// Auditpinfo is a Go-name alias for Auditpinfo_t.
type Auditpinfo = Auditpinfo_t

// BacktraceFlags is a Go-name alias for Backtrace_flags_t.
type BacktraceFlags = Backtrace_flags_t

// BacktraceInfo is a Go-name alias for Backtrace_info_t.
type BacktraceInfo = Backtrace_info_t

// BacktracePack is a Go-name alias for Backtrace_pack_t.
type BacktracePack = Backtrace_pack_t

// BankAction is a Go-name alias for Bank_action_t.
type BankAction = Bank_action_t

// Blkcnt is a Go-name alias for Blkcnt_t.
type Blkcnt = Blkcnt_t

// Blksize is a Go-name alias for Blksize_t.
type Blksize = Blksize_t

// BlockHint is a Go-name alias for Block_hint_t.
type BlockHint = Block_hint_t

// BootArgs is a Go-name alias for Boot_args.
type BootArgs = Boot_args

// BootIconElement is a Go-name alias for Boot_icon_element.
type BootIconElement = Boot_icon_element

// Bootstrap is a Go-name alias for Bootstrap_t.
type Bootstrap = Bootstrap_t

// BpfInt32 is a Go-name alias for Bpf_int32.
type BpfInt32 = Bpf_int32

// BpfTapMode is a Go-name alias for Bpf_tap_mode.
type BpfTapMode = Bpf_tap_mode

// BpfUInt32 is a Go-name alias for Bpf_u_int32.
type BpfUInt32 = Bpf_u_int32

// BufBptr is a Go-name alias for Buf_bptr_t.
type BufBptr = Buf_bptr_t

// BufPtrRef is a Go-name alias for Buf_ptr_ref_t.
type BufPtrRef = Buf_ptr_ref_t

// BufPtr is a Go-name alias for Buf_ptr_t.
type BufPtr = Buf_ptr_t

// BufRefPtr is a Go-name alias for Buf_ref_ptr_t.
type BufRefPtr = Buf_ref_ptr_t

// BufRefRef is a Go-name alias for Buf_ref_ref_t.
type BufRefRef = Buf_ref_ref_t

// BufRef is a Go-name alias for Buf_ref_t.
type BufRef = Buf_ref_t

// Buf is a Go-name alias for Buf_t.
type Buf = Buf_t

// BufattrBptr is a Go-name alias for Bufattr_bptr_t.
type BufattrBptr = Bufattr_bptr_t

// BufattrPtrRef is a Go-name alias for Bufattr_ptr_ref_t.
type BufattrPtrRef = Bufattr_ptr_ref_t

// BufattrPtr is a Go-name alias for Bufattr_ptr_t.
type BufattrPtr = Bufattr_ptr_t

// BufattrRefPtr is a Go-name alias for Bufattr_ref_ptr_t.
type BufattrRefPtr = Bufattr_ref_ptr_t

// BufattrRefRef is a Go-name alias for Bufattr_ref_ref_t.
type BufattrRefRef = Bufattr_ref_ref_t

// BufattrRef is a Go-name alias for Bufattr_ref_t.
type BufattrRef = Bufattr_ref_t

// Bufattr is a Go-name alias for Bufattr_t.
type Bufattr = Bufattr_t

// CacheType is a Go-name alias for Cache_type_t.
type CacheType = Cache_type_t

// Caddr is a Go-name alias for Caddr_t.
type Caddr = Caddr_t

// CaddrUt is a Go-name alias for Caddr_ut.
type CaddrUt = Caddr_ut

// CallGate is a Go-name alias for Call_gate_t.
type CallGate = Call_gate_t

// Cc is a Go-name alias for Cc_t.
type Cc = Cc_t

// CircleQueueHead is a Go-name alias for Circle_queue_head_t.
type CircleQueueHead = Circle_queue_head_t

// CircleQueue is a Go-name alias for Circle_queue_t.
type CircleQueue = Circle_queue_t

// ClDirectReadLock is a Go-name alias for Cl_direct_read_lock_t.
type ClDirectReadLock = Cl_direct_read_lock_t

// ClockAttr is a Go-name alias for Clock_attr_t.
type ClockAttr = Clock_attr_t

// ClockCtrlPort is a Go-name alias for Clock_ctrl_port_t.
type ClockCtrlPort = Clock_ctrl_port_t

// ClockCtrl is a Go-name alias for Clock_ctrl_t.
type ClockCtrl = Clock_ctrl_t

// ClockFlavor is a Go-name alias for Clock_flavor_t.
type ClockFlavor = Clock_flavor_t

// ClockFrequencyInfo is a Go-name alias for Clock_frequency_info_t.
type ClockFrequencyInfo = Clock_frequency_info_t

// ClockID is a Go-name alias for Clock_id_t.
type ClockID = Clock_id_t

// ClockNsec is a Go-name alias for Clock_nsec_t.
type ClockNsec = Clock_nsec_t

// ClockReply is a Go-name alias for Clock_reply_t.
type ClockReply = Clock_reply_t

// ClockRes is a Go-name alias for Clock_res_t.
type ClockRes = Clock_res_t

// ClockSec is a Go-name alias for Clock_sec_t.
type ClockSec = Clock_sec_t

// ClockServPort is a Go-name alias for Clock_serv_port_t.
type ClockServPort = Clock_serv_port_t

// ClockServ is a Go-name alias for Clock_serv_t.
type ClockServ = Clock_serv_t

// Clock is a Go-name alias for Clock_t.
type Clock = Clock_t

// ClockUsec is a Go-name alias for Clock_usec_t.
type ClockUsec = Clock_usec_t

// ClusterType is a Go-name alias for Cluster_type_t.
type ClusterType = Cluster_type_t

// Coalition is a Go-name alias for Coalition_t.
type Coalition = Coalition_t

// CodeDesc is a Go-name alias for Code_desc_t.
type CodeDesc = Code_desc_t

// ConninfoMultipathtcp is a Go-name alias for Conninfo_multipathtcp_t.
type ConninfoMultipathtcp = Conninfo_multipathtcp_t

// ConninfoTCP is a Go-name alias for Conninfo_tcp_t.
type ConninfoTCP = Conninfo_tcp_t

// CoprocessorType is a Go-name alias for Coprocessor_type_t.
type CoprocessorType = Coprocessor_type_t

// CPUID is a Go-name alias for Cpu_id_t.
type CPUID = Cpu_id_t

// CPUSubtype is a Go-name alias for Cpu_subtype_t.
type CPUSubtype = Cpu_subtype_t

// CPUThreadtype is a Go-name alias for Cpu_threadtype_t.
type CPUThreadtype = Cpu_threadtype_t

// CPUType is a Go-name alias for Cpu_type_t.
type CPUType = Cpu_type_t

// CpuidArchPerfLeaf is a Go-name alias for Cpuid_arch_perf_leaf_t.
type CpuidArchPerfLeaf = Cpuid_arch_perf_leaf_t

// CpuidCacheDesc is a Go-name alias for Cpuid_cache_desc_t.
type CpuidCacheDesc = Cpuid_cache_desc_t

// CpuidMwaitLeaf is a Go-name alias for Cpuid_mwait_leaf_t.
type CpuidMwaitLeaf = Cpuid_mwait_leaf_t

// CpuidRegister is a Go-name alias for Cpuid_register_t.
type CpuidRegister = Cpuid_register_t

// CpuidThermalLeaf is a Go-name alias for Cpuid_thermal_leaf_t.
type CpuidThermalLeaf = Cpuid_thermal_leaf_t

// CpuidTscLeaf is a Go-name alias for Cpuid_tsc_leaf_t.
type CpuidTscLeaf = Cpuid_tsc_leaf_t

// CpuidXsaveLeaf is a Go-name alias for Cpuid_xsave_leaf_t.
type CpuidXsaveLeaf = Cpuid_xsave_leaf_t

// Cr0 is a Go-name alias for Cr0_t.
type Cr0 = Cr0_t

// CryptexAuthType is a Go-name alias for Cryptex_auth_type_t.
type CryptexAuthType = Cryptex_auth_type_t

// CryptoRandomCtx is a Go-name alias for Crypto_random_ctx_t.
type CryptoRandomCtx = Crypto_random_ctx_t

// CryptoRandomKmemCtxSizeFn is a Go-name alias for Crypto_random_kmem_ctx_size_fn_t.
type CryptoRandomKmemCtxSizeFn = Crypto_random_kmem_ctx_size_fn_t

// CtRune is a Go-name alias for Ct_rune_t.
type CtRune = Ct_rune_t

// DDevtotty is a Go-name alias for D_devtotty_t.
type DDevtotty = D_devtotty_t

// Daddr64 is a Go-name alias for Daddr64_t.
type Daddr64 = Daddr64_t

// Daddr is a Go-name alias for Daddr_t.
type Daddr = Daddr_t

// DataDesc is a Go-name alias for Data_desc_t.
type DataDesc = Data_desc_t

// DebugHeaderEntry is a Go-name alias for Debug_header_entry.
type DebugHeaderEntry = Debug_header_entry

// DebugHeader is a Go-name alias for Debug_header_t.
type DebugHeader = Debug_header_t

// DebugTrailer is a Go-name alias for Debug_trailer_t.
type DebugTrailer = Debug_trailer_t

// DescriptorOptions is a Go-name alias for Descriptor_options.
type DescriptorOptions = Descriptor_options

// Dev is a Go-name alias for Dev_t.
type Dev = Dev_t

// DirCloneAuthorizerOp is a Go-name alias for Dir_clone_authorizer_op_t.
type DirCloneAuthorizerOp = Dir_clone_authorizer_op_t

// DkBdReadDiscInfo is a Go-name alias for Dk_bd_read_disc_info_t.
type DkBdReadDiscInfo = Dk_bd_read_disc_info_t

// DkBdReadStructure is a Go-name alias for Dk_bd_read_structure_t.
type DkBdReadStructure = Dk_bd_read_structure_t

// DkBdReadTrackInfo is a Go-name alias for Dk_bd_read_track_info_t.
type DkBdReadTrackInfo = Dk_bd_read_track_info_t

// DkBdReportKey is a Go-name alias for Dk_bd_report_key_t.
type DkBdReportKey = Dk_bd_report_key_t

// DkBdSendKey is a Go-name alias for Dk_bd_send_key_t.
type DkBdSendKey = Dk_bd_send_key_t

// DkCdReadDiscInfo is a Go-name alias for Dk_cd_read_disc_info_t.
type DkCdReadDiscInfo = Dk_cd_read_disc_info_t

// DkCdReadIsrc is a Go-name alias for Dk_cd_read_isrc_t.
type DkCdReadIsrc = Dk_cd_read_isrc_t

// DkCdReadMcn is a Go-name alias for Dk_cd_read_mcn_t.
type DkCdReadMcn = Dk_cd_read_mcn_t

// DkCdRead is a Go-name alias for Dk_cd_read_t.
type DkCdRead = Dk_cd_read_t

// DkCdReadToc is a Go-name alias for Dk_cd_read_toc_t.
type DkCdReadToc = Dk_cd_read_toc_t

// DkCdReadTrackInfo is a Go-name alias for Dk_cd_read_track_info_t.
type DkCdReadTrackInfo = Dk_cd_read_track_info_t

// DkCorestorageInfo is a Go-name alias for Dk_corestorage_info_t.
type DkCorestorageInfo = Dk_corestorage_info_t

// DkDvdReadDiscInfo is a Go-name alias for Dk_dvd_read_disc_info_t.
type DkDvdReadDiscInfo = Dk_dvd_read_disc_info_t

// DkDvdReadRzoneInfo is a Go-name alias for Dk_dvd_read_rzone_info_t.
type DkDvdReadRzoneInfo = Dk_dvd_read_rzone_info_t

// DkDvdReadStructure is a Go-name alias for Dk_dvd_read_structure_t.
type DkDvdReadStructure = Dk_dvd_read_structure_t

// DkDvdReportKey is a Go-name alias for Dk_dvd_report_key_t.
type DkDvdReportKey = Dk_dvd_report_key_t

// DkDvdSendKey is a Go-name alias for Dk_dvd_send_key_t.
type DkDvdSendKey = Dk_dvd_send_key_t

// DkErrorDescription is a Go-name alias for Dk_error_description_t.
type DkErrorDescription = Dk_error_description_t

// DkExtent is a Go-name alias for Dk_extent_t.
type DkExtent = Dk_extent_t

// DkFirmwarePath is a Go-name alias for Dk_firmware_path_t.
type DkFirmwarePath = Dk_firmware_path_t

// DkFormatCapacities is a Go-name alias for Dk_format_capacities_t.
type DkFormatCapacities = Dk_format_capacities_t

// DkFormatCapacity is a Go-name alias for Dk_format_capacity_t.
type DkFormatCapacity = Dk_format_capacity_t

// DkPhysicalExtent is a Go-name alias for Dk_physical_extent_t.
type DkPhysicalExtent = Dk_physical_extent_t

// DkProvisionExtent is a Go-name alias for Dk_provision_extent_t.
type DkProvisionExtent = Dk_provision_extent_t

// DkProvisionStatus is a Go-name alias for Dk_provision_status_t.
type DkProvisionStatus = Dk_provision_status_t

// DkSetTier is a Go-name alias for Dk_set_tier_t.
type DkSetTier = Dk_set_tier_t

// DkSynchronize is a Go-name alias for Dk_synchronize_t.
type DkSynchronize = Dk_synchronize_t

// DkUnmap is a Go-name alias for Dk_unmap_t.
type DkUnmap = Dk_unmap_t

// Double is a Go-name alias for Double_t.
type Double = Double_t

// DumpFcn is a Go-name alias for Dump_fcn_t.
type DumpFcn = Dump_fcn_t

// DyldKernelImageInfoArray is a Go-name alias for Dyld_kernel_image_info_array_t.
type DyldKernelImageInfoArray = Dyld_kernel_image_info_array_t

// DyldKernelImageInfo is a Go-name alias for Dyld_kernel_image_info_t.
type DyldKernelImageInfo = Dyld_kernel_image_info_t

// DyldKernelProcessInfo is a Go-name alias for Dyld_kernel_process_info_t.
type DyldKernelProcessInfo = Dyld_kernel_process_info_t

// EccEvent is a Go-name alias for Ecc_event_t.
type EccEvent = Ecc_event_t

// EccFlags is a Go-name alias for Ecc_flags_t.
type EccFlags = Ecc_flags_t

// EccVersion is a Go-name alias for Ecc_version_t.
type EccVersion = Ecc_version_t

// EmptyFcn is a Go-name alias for Empty_fcn_t.
type EmptyFcn = Empty_fcn_t

// EmulationVector is a Go-name alias for Emulation_vector_t.
type EmulationVector = Emulation_vector_t

// EphPanicFlags is a Go-name alias for Eph_panic_flags_t.
type EphPanicFlags = Eph_panic_flags_t

// Er is a Go-name alias for Er_t.
type Er = Er_t

// Errno is a Go-name alias for Errno_t.
type Errno = Errno_t

// EtherAddr is a Go-name alias for Ether_addr_t.
type EtherAddr = Ether_addr_t

// EtherHeader is a Go-name alias for Ether_header_t.
type EtherHeader = Ether_header_t

// Event64 is a Go-name alias for Event64_t.
type Event64 = Event64_t

// Event is a Go-name alias for Event_t.
type Event = Event_t

// EventlinkPortPair is a Go-name alias for Eventlink_port_pair_t.
type EventlinkPortPair = Eventlink_port_pair_t

// EvioSpecialKeyMsg is a Go-name alias for EvioSpecialKeyMsg_t.
type EvioSpecialKeyMsg = EvioSpecialKeyMsg_t

// ExCbAction is a Go-name alias for Ex_cb_action_t.
type ExCbAction = Ex_cb_action_t

// ExCbClass is a Go-name alias for Ex_cb_class_t.
type ExCbClass = Ex_cb_class_t

// ExCbState is a Go-name alias for Ex_cb_state_t.
type ExCbState = Ex_cb_state_t

// ExceptionBehaviorArray is a Go-name alias for Exception_behavior_array_t.
type ExceptionBehaviorArray = Exception_behavior_array_t

// ExceptionBehavior is a Go-name alias for Exception_behavior_t.
type ExceptionBehavior = Exception_behavior_t

// ExceptionData is a Go-name alias for Exception_data_t.
type ExceptionData = Exception_data_t

// ExceptionDataType is a Go-name alias for Exception_data_type_t.
type ExceptionDataType = Exception_data_type_t

// ExceptionFlavorArray is a Go-name alias for Exception_flavor_array_t.
type ExceptionFlavorArray = Exception_flavor_array_t

// ExceptionHandlerArray is a Go-name alias for Exception_handler_array_t.
type ExceptionHandlerArray = Exception_handler_array_t

// ExceptionHandlerInfoArray is a Go-name alias for Exception_handler_info_array_t.
type ExceptionHandlerInfoArray = Exception_handler_info_array_t

// ExceptionHandlerInfo is a Go-name alias for Exception_handler_info_t.
type ExceptionHandlerInfo = Exception_handler_info_t

// ExceptionHandler is a Go-name alias for Exception_handler_t.
type ExceptionHandler = Exception_handler_t

// ExceptionMaskArray is a Go-name alias for Exception_mask_array_t.
type ExceptionMaskArray = Exception_mask_array_t

// ExceptionMask is a Go-name alias for Exception_mask_t.
type ExceptionMask = Exception_mask_t

// ExceptionPortArrary is a Go-name alias for Exception_port_arrary_t.
type ExceptionPortArrary = Exception_port_arrary_t

// ExceptionPortArray is a Go-name alias for Exception_port_array_t.
type ExceptionPortArray = Exception_port_array_t

// ExceptionPortInfoArray is a Go-name alias for Exception_port_info_array_t.
type ExceptionPortInfoArray = Exception_port_info_array_t

// ExceptionPort is a Go-name alias for Exception_port_t.
type ExceptionPort = Exception_port_t

// ExceptionType is a Go-name alias for Exception_type_t.
type ExceptionType = Exception_type_t

// ExclaveEcstackentryAddr is a Go-name alias for Exclave_ecstackentry_addr_t.
type ExclaveEcstackentryAddr = Exclave_ecstackentry_addr_t

// ExtPaniclogCreateOptions is a Go-name alias for Ext_paniclog_create_options_t.
type ExtPaniclogCreateOptions = Ext_paniclog_create_options_t

// Fattributiontag is a Go-name alias for Fattributiontag_t.
type Fattributiontag = Fattributiontag_t

// Fchecklv is a Go-name alias for Fchecklv_t.
type Fchecklv = Fchecklv_t

// FdMask is a Go-name alias for Fd_mask.
type FdMask = Fd_mask

// FdSet is a Go-name alias for Fd_set.
type FdSet = Fd_set

// Fgetsigsinfo is a Go-name alias for Fgetsigsinfo_t.
type Fgetsigsinfo = Fgetsigsinfo_t

// Fhandle is a Go-name alias for Fhandle_t.
type Fhandle = Fhandle_t

// FileBptr is a Go-name alias for File_bptr_t.
type FileBptr = File_bptr_t

// FilePtrRef is a Go-name alias for File_ptr_ref_t.
type FilePtrRef = File_ptr_ref_t

// FilePtr is a Go-name alias for File_ptr_t.
type FilePtr = File_ptr_t

// FileRefPtr is a Go-name alias for File_ref_ptr_t.
type FileRefPtr = File_ref_ptr_t

// FileRefRef is a Go-name alias for File_ref_ref_t.
type FileRefRef = File_ref_ref_t

// FileRef is a Go-name alias for File_ref_t.
type FileRef = File_ref_t

// File is a Go-name alias for File_t.
type File = File_t

// Filesec is a Go-name alias for Filesec_t.
type Filesec = Filesec_t

// Fixpt is a Go-name alias for Fixpt_t.
type Fixpt = Fixpt_t

// Float is a Go-name alias for Float_t.
type Float = Float_t

// FPControl is a Go-name alias for Fp_control_t.
type FPControl = Fp_control_t

// FPStatus is a Go-name alias for Fp_status_t.
type FPStatus = Fp_status_t

// Fpunchhole is a Go-name alias for Fpunchhole_t.
type Fpunchhole = Fpunchhole_t

// FrameTypeBitmask is a Go-name alias for Frame_type_bitmask_t.
type FrameTypeBitmask = Frame_type_bitmask_t

// FsRoleMountArgs is a Go-name alias for Fs_role_mount_args_t.
type FsRoleMountArgs = Fs_role_mount_args_t

// Fsblkcnt is a Go-name alias for Fsblkcnt_t.
type Fsblkcnt = Fsblkcnt_t

// Fsfilcnt is a Go-name alias for Fsfilcnt_t.
type Fsfilcnt = Fsfilcnt_t

// FsfileType is a Go-name alias for Fsfile_type_t.
type FsfileType = Fsfile_type_t

// Fsid is a Go-name alias for Fsid_t.
type Fsid = Fsid_t

// Fsignatures is a Go-name alias for Fsignatures_t.
type Fsignatures = Fsignatures_t

// FsobjID is a Go-name alias for Fsobj_id_t.
type FsobjID = Fsobj_id_t

// FsobjTag is a Go-name alias for Fsobj_tag_t.
type FsobjTag = Fsobj_tag_t

// FsobjType is a Go-name alias for Fsobj_type_t.
type FsobjType = Fsobj_type_t

// Fspecread is a Go-name alias for Fspecread_t.
type Fspecread = Fspecread_t

// Fstore is a Go-name alias for Fstore_t.
type Fstore = Fstore_t

// Fsupplement is a Go-name alias for Fsupplement_t.
type Fsupplement = Fsupplement_t

// Fsvolid is a Go-name alias for Fsvolid_t.
type Fsvolid = Fsvolid_t

// Ftrimactivefile is a Go-name alias for Ftrimactivefile_t.
type Ftrimactivefile = Ftrimactivefile_t

// Gdt is a Go-name alias for Gdt_t.
type Gdt = Gdt_t

// Gid is a Go-name alias for Gid_t.
type Gid = Gid_t

// GPUDescriptor is a Go-name alias for Gpu_descriptor.
type GPUDescriptor = Gpu_descriptor

// GPUEnergyData is a Go-name alias for Gpu_energy_data.
type GPUEnergyData = Gpu_energy_data

// GraftdmgType is a Go-name alias for Graftdmg_type_t.
type GraftdmgType = Graftdmg_type_t

// GssdByteBuffer is a Go-name alias for Gssd_byte_buffer.
type GssdByteBuffer = Gssd_byte_buffer

// GssdCred is a Go-name alias for Gssd_cred.
type GssdCred = Gssd_cred

// GssdCtx is a Go-name alias for Gssd_ctx.
type GssdCtx = Gssd_ctx

// GssdDstring is a Go-name alias for Gssd_dstring.
type GssdDstring = Gssd_dstring

// GssdEtypeList is a Go-name alias for Gssd_etype_list.
type GssdEtypeList = Gssd_etype_list

// GssdGidList is a Go-name alias for Gssd_gid_list.
type GssdGidList = Gssd_gid_list

// GssdString is a Go-name alias for Gssd_string.
type GssdString = Gssd_string

// GzHeader is a Go-name alias for Gz_header.
type GzHeader = Gz_header

// GzHeaderp is a Go-name alias for Gz_headerp.
type GzHeaderp = Gz_headerp

// HashInfoBucketArray is a Go-name alias for Hash_info_bucket_array_t.
type HashInfoBucketArray = Hash_info_bucket_array_t

// HashInfoBucket is a Go-name alias for Hash_info_bucket_t.
type HashInfoBucket = Hash_info_bucket_t

// HostBasicInfoData is a Go-name alias for Host_basic_info_data_t.
type HostBasicInfoData = Host_basic_info_data_t

// HostBasicInfo is a Go-name alias for Host_basic_info_t.
type HostBasicInfo = Host_basic_info_t

// HostCanHasDebuggerInfoData is a Go-name alias for Host_can_has_debugger_info_data_t.
type HostCanHasDebuggerInfoData = Host_can_has_debugger_info_data_t

// HostCanHasDebuggerInfo is a Go-name alias for Host_can_has_debugger_info_t.
type HostCanHasDebuggerInfo = Host_can_has_debugger_info_t

// HostCPULoadInfoData is a Go-name alias for Host_cpu_load_info_data_t.
type HostCPULoadInfoData = Host_cpu_load_info_data_t

// HostCPULoadInfo is a Go-name alias for Host_cpu_load_info_t.
type HostCPULoadInfo = Host_cpu_load_info_t

// HostFlavor is a Go-name alias for Host_flavor_t.
type HostFlavor = Host_flavor_t

// HostInfo64 is a Go-name alias for Host_info64_t.
type HostInfo64 = Host_info64_t

// HostInfoData is a Go-name alias for Host_info_data_t.
type HostInfoData = Host_info_data_t

// HostInfo is a Go-name alias for Host_info_t.
type HostInfo = Host_info_t

// HostLoadInfoData is a Go-name alias for Host_load_info_data_t.
type HostLoadInfoData = Host_load_info_data_t

// HostLoadInfo is a Go-name alias for Host_load_info_t.
type HostLoadInfo = Host_load_info_t

// HostNamePort is a Go-name alias for Host_name_port_t.
type HostNamePort = Host_name_port_t

// HostName is a Go-name alias for Host_name_t.
type HostName = Host_name_t

// HostPreferredUserArchData is a Go-name alias for Host_preferred_user_arch_data_t.
type HostPreferredUserArchData = Host_preferred_user_arch_data_t

// HostPreferredUserArch is a Go-name alias for Host_preferred_user_arch_t.
type HostPreferredUserArch = Host_preferred_user_arch_t

// HostPriorityInfoData is a Go-name alias for Host_priority_info_data_t.
type HostPriorityInfoData = Host_priority_info_data_t

// HostPriorityInfo is a Go-name alias for Host_priority_info_t.
type HostPriorityInfo = Host_priority_info_t

// HostPriv is a Go-name alias for Host_priv_t.
type HostPriv = Host_priv_t

// HostPurgableInfoData is a Go-name alias for Host_purgable_info_data_t.
type HostPurgableInfoData = Host_purgable_info_data_t

// HostPurgableInfo is a Go-name alias for Host_purgable_info_t.
type HostPurgableInfo = Host_purgable_info_t

// HostSchedInfoData is a Go-name alias for Host_sched_info_data_t.
type HostSchedInfoData = Host_sched_info_data_t

// HostSchedInfo is a Go-name alias for Host_sched_info_t.
type HostSchedInfo = Host_sched_info_t

// HostSecurity is a Go-name alias for Host_security_t.
type HostSecurity = Host_security_t

// Host is a Go-name alias for Host_t.
type Host = Host_t

// HVCallbacks is a Go-name alias for Hv_callbacks_t.
type HVCallbacks = Hv_callbacks_t

// HVTrapTable is a Go-name alias for Hv_trap_table_t.
type HVTrapTable = Hv_trap_table_t

// HVTrapType is a Go-name alias for Hv_trap_type_t.
type HVTrapType = Hv_trap_type_t

// HVVolatileState is a Go-name alias for Hv_volatile_state_t.
type HVVolatileState = Hv_volatile_state_t

// HvgHcallArgs is a Go-name alias for Hvg_hcall_args_t.
type HvgHcallArgs = Hvg_hcall_args_t

// HvgHcallOutput is a Go-name alias for Hvg_hcall_output_t.
type HvgHcallOutput = Hvg_hcall_output_t

// HvgHcallVmcoreFile is a Go-name alias for Hvg_hcall_vmcore_file_t.
type HvgHcallVmcoreFile = Hvg_hcall_vmcore_file_t

// HwSpinPolicy is a Go-name alias for Hw_spin_policy_t.
type HwSpinPolicy = Hw_spin_policy_t

// I386CPUInfo is a Go-name alias for I386_cpu_info_t.
type I386CPUInfo = I386_cpu_info_t

// I386Ioport is a Go-name alias for I386_ioport_t.
type I386Ioport = I386_ioport_t

// ID is a Go-name alias for Id_t.
type ID = Id_t

// IdleTickle is a Go-name alias for Idle_tickle_t.
type IdleTickle = Idle_tickle_t

// Idt is a Go-name alias for Idt_t.
type Idt = Idt_t

// Idtype is a Go-name alias for Idtype_t.
type Idtype = Idtype_t

// IfCloneBptr is a Go-name alias for If_clone_bptr_t.
type IfCloneBptr = If_clone_bptr_t

// IfClonePtrRef is a Go-name alias for If_clone_ptr_ref_t.
type IfClonePtrRef = If_clone_ptr_ref_t

// IfClonePtr is a Go-name alias for If_clone_ptr_t.
type IfClonePtr = If_clone_ptr_t

// IfCloneRefPtr is a Go-name alias for If_clone_ref_ptr_t.
type IfCloneRefPtr = If_clone_ref_ptr_t

// IfCloneRefRef is a Go-name alias for If_clone_ref_ref_t.
type IfCloneRefRef = If_clone_ref_ref_t

// IfCloneRef is a Go-name alias for If_clone_ref_t.
type IfCloneRef = If_clone_ref_t

// IfClone is a Go-name alias for If_clone_t.
type IfClone = If_clone_t

// IfNetemModel is a Go-name alias for If_netem_model_t.
type IfNetemModel = If_netem_model_t

// IfaddrBptr is a Go-name alias for Ifaddr_bptr_t.
type IfaddrBptr = Ifaddr_bptr_t

// IfaddrPtrRef is a Go-name alias for Ifaddr_ptr_ref_t.
type IfaddrPtrRef = Ifaddr_ptr_ref_t

// IfaddrPtr is a Go-name alias for Ifaddr_ptr_t.
type IfaddrPtr = Ifaddr_ptr_t

// IfaddrRefPtr is a Go-name alias for Ifaddr_ref_ptr_t.
type IfaddrRefPtr = Ifaddr_ref_ptr_t

// IfaddrRefRef is a Go-name alias for Ifaddr_ref_ref_t.
type IfaddrRefRef = Ifaddr_ref_ref_t

// IfaddrRef is a Go-name alias for Ifaddr_ref_t.
type IfaddrRef = Ifaddr_ref_t

// Ifaddr is a Go-name alias for Ifaddr_t.
type Ifaddr = Ifaddr_t

// IfmultiaddrBptr is a Go-name alias for Ifmultiaddr_bptr_t.
type IfmultiaddrBptr = Ifmultiaddr_bptr_t

// IfmultiaddrPtrRef is a Go-name alias for Ifmultiaddr_ptr_ref_t.
type IfmultiaddrPtrRef = Ifmultiaddr_ptr_ref_t

// IfmultiaddrPtr is a Go-name alias for Ifmultiaddr_ptr_t.
type IfmultiaddrPtr = Ifmultiaddr_ptr_t

// IfmultiaddrRefPtr is a Go-name alias for Ifmultiaddr_ref_ptr_t.
type IfmultiaddrRefPtr = Ifmultiaddr_ref_ptr_t

// IfmultiaddrRefRef is a Go-name alias for Ifmultiaddr_ref_ref_t.
type IfmultiaddrRefRef = Ifmultiaddr_ref_ref_t

// IfmultiaddrRef is a Go-name alias for Ifmultiaddr_ref_t.
type IfmultiaddrRef = Ifmultiaddr_ref_t

// Ifmultiaddr is a Go-name alias for Ifmultiaddr_t.
type Ifmultiaddr = Ifmultiaddr_t

// IfnetBptr is a Go-name alias for Ifnet_bptr_t.
type IfnetBptr = Ifnet_bptr_t

// IfnetFilterBptr is a Go-name alias for Ifnet_filter_bptr_t.
type IfnetFilterBptr = Ifnet_filter_bptr_t

// IfnetFilterPtrRef is a Go-name alias for Ifnet_filter_ptr_ref_t.
type IfnetFilterPtrRef = Ifnet_filter_ptr_ref_t

// IfnetFilterPtr is a Go-name alias for Ifnet_filter_ptr_t.
type IfnetFilterPtr = Ifnet_filter_ptr_t

// IfnetFilterRefPtr is a Go-name alias for Ifnet_filter_ref_ptr_t.
type IfnetFilterRefPtr = Ifnet_filter_ref_ptr_t

// IfnetFilterRefRef is a Go-name alias for Ifnet_filter_ref_ref_t.
type IfnetFilterRefRef = Ifnet_filter_ref_ref_t

// IfnetFilterRef is a Go-name alias for Ifnet_filter_ref_t.
type IfnetFilterRef = Ifnet_filter_ref_t

// IfnetOffload is a Go-name alias for Ifnet_offload_t.
type IfnetOffload = Ifnet_offload_t

// IfnetPtrRef is a Go-name alias for Ifnet_ptr_ref_t.
type IfnetPtrRef = Ifnet_ptr_ref_t

// IfnetPtr is a Go-name alias for Ifnet_ptr_t.
type IfnetPtr = Ifnet_ptr_t

// IfnetRefPtr is a Go-name alias for Ifnet_ref_ptr_t.
type IfnetRefPtr = Ifnet_ref_ptr_t

// IfnetRefRef is a Go-name alias for Ifnet_ref_ref_t.
type IfnetRefRef = Ifnet_ref_ref_t

// IfnetRef is a Go-name alias for Ifnet_ref_t.
type IfnetRef = Ifnet_ref_t

// In6Addr is a Go-name alias for In6_addr_t.
type In6Addr = In6_addr_t

// In6Clat46EvhdlrCode is a Go-name alias for In6_clat46_evhdlr_code_t.
type In6Clat46EvhdlrCode = In6_clat46_evhdlr_code_t

// InAddr is a Go-name alias for In_addr_t.
type InAddr = In_addr_t

// InPort is a Go-name alias for In_port_t.
type InPort = In_port_t

// Ino64 is a Go-name alias for Ino64_t.
type Ino64 = Ino64_t

// Ino is a Go-name alias for Ino_t.
type Ino = Ino_t

// InpGen is a Go-name alias for Inp_gen_t.
type InpGen = Inp_gen_t

// Int16 is a Go-name alias for Int16_t.
type Int16 = Int16_t

// Int32 is a Go-name alias for Int32_t.
type Int32 = Int32_t

// Int64 is a Go-name alias for Int64_t.
type Int64 = Int64_t

// Int8 is a Go-name alias for Int8_t.
type Int8 = Int8_t

// IntFast16 is a Go-name alias for Int_fast16_t.
type IntFast16 = Int_fast16_t

// IntFast32 is a Go-name alias for Int_fast32_t.
type IntFast32 = Int_fast32_t

// IntFast64 is a Go-name alias for Int_fast64_t.
type IntFast64 = Int_fast64_t

// IntFast8 is a Go-name alias for Int_fast8_t.
type IntFast8 = Int_fast8_t

// IntLeast16 is a Go-name alias for Int_least16_t.
type IntLeast16 = Int_least16_t

// IntLeast32 is a Go-name alias for Int_least32_t.
type IntLeast32 = Int_least32_t

// IntLeast64 is a Go-name alias for Int_least64_t.
type IntLeast64 = Int_least64_t

// IntLeast8 is a Go-name alias for Int_least8_t.
type IntLeast8 = Int_least8_t

// Integer is a Go-name alias for Integer_t.
type Integer = Integer_t

// InterfaceFilter is a Go-name alias for Interface_filter_t.
type InterfaceFilter = Interface_filter_t

// Intmax is a Go-name alias for Intmax_t.
type Intmax = Intmax_t

// Intptr is a Go-name alias for Intptr_t.
type Intptr = Intptr_t

// IntrGate is a Go-name alias for Intr_gate_t.
type IntrGate = Intr_gate_t

// IOAddr is a Go-name alias for Io_addr_t.
type IOAddr = Io_addr_t

// IOAsyncRef64 is a Go-name alias for Io_async_ref64_t.
type IOAsyncRef64 = Io_async_ref64_t

// IOAsyncRef is a Go-name alias for Io_async_ref_t.
type IOAsyncRef = Io_async_ref_t

// IOBufPtr is a Go-name alias for Io_buf_ptr_t.
type IOBufPtr = Io_buf_ptr_t

// IOCompressionStats is a Go-name alias for Io_compression_stats_t.
type IOCompressionStats = Io_compression_stats_t

// IOConnect is a Go-name alias for Io_connect_t.
type IOConnect = Io_connect_t

// IOEnumerator is a Go-name alias for Io_enumerator_t.
type IOEnumerator = Io_enumerator_t

// IOIdent is a Go-name alias for Io_ident_t.
type IOIdent = Io_ident_t

// IOIterator is a Go-name alias for Io_iterator_t.
type IOIterator = Io_iterator_t

// IOLen is a Go-name alias for Io_len_t.
type IOLen = Io_len_t

// IOMain is a Go-name alias for Io_main_t.
type IOMain = Io_main_t

// IOName is a Go-name alias for Io_name_t.
type IOName = Io_name_t

// IOObject is a Go-name alias for Io_object_t.
type IOObject = Io_object_t

// IOScalarInband64 is a Go-name alias for Io_scalar_inband64_t.
type IOScalarInband64 = Io_scalar_inband64_t

// IOScalarInband is a Go-name alias for Io_scalar_inband_t.
type IOScalarInband = Io_scalar_inband_t

// IOStatInfo is a Go-name alias for Io_stat_info_t.
type IOStatInfo = Io_stat_info_t

// IOStringInband is a Go-name alias for Io_string_inband_t.
type IOStringInband = Io_string_inband_t

// IOString is a Go-name alias for Io_string_t.
type IOString = Io_string_t

// IOStructInband is a Go-name alias for Io_struct_inband_t.
type IOStructInband = Io_struct_inband_t

// IOUserReference is a Go-name alias for Io_user_reference_t.
type IOUserReference = Io_user_reference_t

// IOUserScalar is a Go-name alias for Io_user_scalar_t.
type IOUserScalar = Io_user_scalar_t

// IoctlFcn is a Go-name alias for Ioctl_fcn_t.
type IoctlFcn = Ioctl_fcn_t

// IpcEventlink is a Go-name alias for Ipc_eventlink_t.
type IpcEventlink = Ipc_eventlink_t

// IpcInfoNameArray is a Go-name alias for Ipc_info_name_array_t.
type IpcInfoNameArray = Ipc_info_name_array_t

// IpcInfoName is a Go-name alias for Ipc_info_name_t.
type IpcInfoName = Ipc_info_name_t

// IpcInfoPort is a Go-name alias for Ipc_info_port_t.
type IpcInfoPort = Ipc_info_port_t

// IpcInfoSpaceBasic is a Go-name alias for Ipc_info_space_basic_t.
type IpcInfoSpaceBasic = Ipc_info_space_basic_t

// IpcInfoSpace is a Go-name alias for Ipc_info_space_t.
type IpcInfoSpace = Ipc_info_space_t

// IpcInfoTreeNameArray is a Go-name alias for Ipc_info_tree_name_array_t.
type IpcInfoTreeNameArray = Ipc_info_tree_name_array_t

// IpcInfoTreeName is a Go-name alias for Ipc_info_tree_name_t.
type IpcInfoTreeName = Ipc_info_tree_name_t

// IpcObject is a Go-name alias for Ipc_object_t.
type IpcObject = Ipc_object_t

// IpcPort is a Go-name alias for Ipc_port_t.
type IpcPort = Ipc_port_t

// IpcPthreadPriorityValue is a Go-name alias for Ipc_pthread_priority_value_t.
type IpcPthreadPriorityValue = Ipc_pthread_priority_value_t

// IpcSpaceInspect is a Go-name alias for Ipc_space_inspect_t.
type IpcSpaceInspect = Ipc_space_inspect_t

// IpcSpacePort is a Go-name alias for Ipc_space_port_t.
type IpcSpacePort = Ipc_space_port_t

// IpcSpaceRead is a Go-name alias for Ipc_space_read_t.
type IpcSpaceRead = Ipc_space_read_t

// IpcSpace is a Go-name alias for Ipc_space_t.
type IpcSpace = Ipc_space_t

// IpcVoucherAttrControl is a Go-name alias for Ipc_voucher_attr_control_t.
type IpcVoucherAttrControl = Ipc_voucher_attr_control_t

// IpcVoucherAttrManager is a Go-name alias for Ipc_voucher_attr_manager_t.
type IpcVoucherAttrManager = Ipc_voucher_attr_manager_t

// IpcVoucher is a Go-name alias for Ipc_voucher_t.
type IpcVoucher = Ipc_voucher_t

// IpiHandler is a Go-name alias for Ipi_handler_t.
type IpiHandler = Ipi_handler_t

// KauthAceRights is a Go-name alias for Kauth_ace_rights_t.
type KauthAceRights = Kauth_ace_rights_t

// KauthAclEval is a Go-name alias for Kauth_acl_eval_t.
type KauthAclEval = Kauth_acl_eval_t

// KauthAction is a Go-name alias for Kauth_action_t.
type KauthAction = Kauth_action_t

// KauthCred is a Go-name alias for Kauth_cred_t.
type KauthCred = Kauth_cred_t

// KauthListener is a Go-name alias for Kauth_listener_t.
type KauthListener = Kauth_listener_t

// KauthScope is a Go-name alias for Kauth_scope_t.
type KauthScope = Kauth_scope_t

// Kbufinfo is a Go-name alias for Kbufinfo_t.
type Kbufinfo = Kbufinfo_t

// KcFormat is a Go-name alias for Kc_format_t.
type KcFormat = Kc_format_t

// KcKind is a Go-name alias for Kc_kind_t.
type KcKind = Kc_kind_t

// KcdCompressionType is a Go-name alias for Kcd_compression_type_t.
type KcdCompressionType = Kcd_compression_type_t

// KcdataDescriptor is a Go-name alias for Kcdata_descriptor_t.
type KcdataDescriptor = Kcdata_descriptor_t

// KcdataItem is a Go-name alias for Kcdata_item_t.
type KcdataItem = Kcdata_item_t

// KcdataIter is a Go-name alias for Kcdata_iter_t.
type KcdataIter = Kcdata_iter_t

// KcdataObject is a Go-name alias for Kcdata_object_t.
type KcdataObject = Kcdata_object_t

// KcdataSubtypeDescriptor is a Go-name alias for Kcdata_subtype_descriptor_t.
type KcdataSubtypeDescriptor = Kcdata_subtype_descriptor_t

// KctypeSubtype is a Go-name alias for Kctype_subtype_t.
type KctypeSubtype = Kctype_subtype_t

// KdBuf is a Go-name alias for Kd_buf.
type KdBuf = Kd_buf

// KdBufArgtype is a Go-name alias for Kd_buf_argtype.
type KdBufArgtype = Kd_buf_argtype

// KdCallbackType is a Go-name alias for Kd_callback_type.
type KdCallbackType = Kd_callback_type

// KdCpumap is a Go-name alias for Kd_cpumap.
type KdCpumap = Kd_cpumap

// KdCpumapExt is a Go-name alias for Kd_cpumap_ext.
type KdCpumapExt = Kd_cpumap_ext

// KdCpumapHeader is a Go-name alias for Kd_cpumap_header.
type KdCpumapHeader = Kd_cpumap_header

// KdEventMatcher is a Go-name alias for Kd_event_matcher.
type KdEventMatcher = Kd_event_matcher

// KdRegtype is a Go-name alias for Kd_regtype.
type KdRegtype = Kd_regtype

// KdThreadmap is a Go-name alias for Kd_threadmap.
type KdThreadmap = Kd_threadmap

// KdebugCoprocFlags is a Go-name alias for Kdebug_coproc_flags_t.
type KdebugCoprocFlags = Kdebug_coproc_flags_t

// KdebugTest is a Go-name alias for Kdebug_test_t.
type KdebugTest = Kdebug_test_t

// KernCtlRef is a Go-name alias for Kern_ctl_ref.
type KernCtlRef = Kern_ctl_ref

// KernReturn is a Go-name alias for Kern_return_t.
type KernReturn = Kern_return_t

// KernelBootInfo is a Go-name alias for Kernel_boot_info_t.
type KernelBootInfo = Kernel_boot_info_t

// KernelResourceSizesData is a Go-name alias for Kernel_resource_sizes_data_t.
type KernelResourceSizesData = Kernel_resource_sizes_data_t

// KernelResourceSizes is a Go-name alias for Kernel_resource_sizes_t.
type KernelResourceSizes = Kernel_resource_sizes_t

// KernelVersion is a Go-name alias for Kernel_version_t.
type KernelVersion = Kernel_version_t

// KfOverrideFlag is a Go-name alias for Kf_override_flag_t.
type KfOverrideFlag = Kf_override_flag_t

// KmodArgs is a Go-name alias for Kmod_args_t.
type KmodArgs = Kmod_args_t

// KmodControlFlavor is a Go-name alias for Kmod_control_flavor_t.
type KmodControlFlavor = Kmod_control_flavor_t

// KmodInfo32V1 is a Go-name alias for Kmod_info_32_v1_t.
type KmodInfo32V1 = Kmod_info_32_v1_t

// KmodInfo64V1 is a Go-name alias for Kmod_info_64_v1_t.
type KmodInfo64V1 = Kmod_info_64_v1_t

// KmodInfoArray is a Go-name alias for Kmod_info_array_t.
type KmodInfoArray = Kmod_info_array_t

// KmodInfo is a Go-name alias for Kmod_info_t.
type KmodInfo = Kmod_info_t

// KmodReference is a Go-name alias for Kmod_reference_t.
type KmodReference = Kmod_reference_t

// KmodStartFunc is a Go-name alias for Kmod_start_func_t.
type KmodStartFunc = Kmod_start_func_t

// KmodStopFunc is a Go-name alias for Kmod_stop_func_t.
type KmodStopFunc = Kmod_stop_func_t

// Kmod is a Go-name alias for Kmod_t.
type Kmod = Kmod_t

// KobjectDescription is a Go-name alias for Kobject_description_t.
type KobjectDescription = Kobject_description_t

// KpcConfig is a Go-name alias for Kpc_config_t.
type KpcConfig = Kpc_config_t

// KpcPmHandler is a Go-name alias for Kpc_pm_handler_t.
type KpcPmHandler = Kpc_pm_handler_t

// KperfKpcFlags is a Go-name alias for Kperf_kpc_flags_t.
type KperfKpcFlags = Kperf_kpc_flags_t

// Labelstr is a Go-name alias for Labelstr_t.
type Labelstr = Labelstr_t

// LaunchConstraintData is a Go-name alias for Launch_constraint_data_t.
type LaunchConstraintData = Launch_constraint_data_t

// LckAttr is a Go-name alias for Lck_attr_t.
type LckAttr = Lck_attr_t

// LckGrpAttr is a Go-name alias for Lck_grp_attr_t.
type LckGrpAttr = Lck_grp_attr_t

// LckGrp is a Go-name alias for Lck_grp_t.
type LckGrp = Lck_grp_t

// LckMtxExt is a Go-name alias for Lck_mtx_ext_t.
type LckMtxExt = Lck_mtx_ext_t

// LckMtx is a Go-name alias for Lck_mtx_t.
type LckMtx = Lck_mtx_t

// LckRw is a Go-name alias for Lck_rw_t.
type LckRw = Lck_rw_t

// LckRwType is a Go-name alias for Lck_rw_type_t.
type LckRwType = Lck_rw_type_t

// LckSleepAction is a Go-name alias for Lck_sleep_action_t.
type LckSleepAction = Lck_sleep_action_t

// LckSpin is a Go-name alias for Lck_spin_t.
type LckSpin = Lck_spin_t

// LckWakeAction is a Go-name alias for Lck_wake_action_t.
type LckWakeAction = Lck_wake_action_t

// LdtDesc is a Go-name alias for Ldt_desc_t.
type LdtDesc = Ldt_desc_t

// Ldt is a Go-name alias for Ldt_t.
type Ldt = Ldt_t

// LedgerAmount is a Go-name alias for Ledger_amount_t.
type LedgerAmount = Ledger_amount_t

// LedgerArray is a Go-name alias for Ledger_array_t.
type LedgerArray = Ledger_array_t

// LedgerItem is a Go-name alias for Ledger_item_t.
type LedgerItem = Ledger_item_t

// LedgerPortArray is a Go-name alias for Ledger_port_array_t.
type LedgerPortArray = Ledger_port_array_t

// LedgerPort is a Go-name alias for Ledger_port_t.
type LedgerPort = Ledger_port_t

// Ledger is a Go-name alias for Ledger_t.
type Ledger = Ledger_t

// LibsptmCPUState is a Go-name alias for Libsptm_cpu_state_t.
type LibsptmCPUState = Libsptm_cpu_state_t

// LibsptmError is a Go-name alias for Libsptm_error_t.
type LibsptmError = Libsptm_error_t

// LibsptmRefcntType is a Go-name alias for Libsptm_refcnt_type_t.
type LibsptmRefcntType = Libsptm_refcnt_type_t

// LibsptmState is a Go-name alias for Libsptm_state_t.
type LibsptmState = Libsptm_state_t

// ListxattrsResult is a Go-name alias for Listxattrs_result_t.
type ListxattrsResult = Listxattrs_result_t

// LockSetPort is a Go-name alias for Lock_set_port_t.
type LockSetPort = Lock_set_port_t

// LockSet is a Go-name alias for Lock_set_t.
type LockSet = Lock_set_t

// LockgroupInfoArray is a Go-name alias for Lockgroup_info_array_t.
type LockgroupInfoArray = Lockgroup_info_array_t

// LockgroupInfo is a Go-name alias for Lockgroup_info_t.
type LockgroupInfo = Lockgroup_info_t

// Lz4HashEntry is a Go-name alias for Lz4_hash_entry_t.
type Lz4HashEntry = Lz4_hash_entry_t

// MachAssertType is a Go-name alias for Mach_assert_type_t.
type MachAssertType = Mach_assert_type_t

// MachAtmSubaid is a Go-name alias for Mach_atm_subaid_t.
type MachAtmSubaid = Mach_atm_subaid_t

// MachBridgeRegwriteTimestampFunc is a Go-name alias for Mach_bridge_regwrite_timestamp_func_t.
type MachBridgeRegwriteTimestampFunc = Mach_bridge_regwrite_timestamp_func_t

// MachDeadNameNotification is a Go-name alias for Mach_dead_name_notification_t.
type MachDeadNameNotification = Mach_dead_name_notification_t

// MachErrorFn is a Go-name alias for Mach_error_fn_t.
type MachErrorFn = Mach_error_fn_t

// MachError is a Go-name alias for Mach_error_t.
type MachError = Mach_error_t

// MachEventlink is a Go-name alias for Mach_eventlink_t.
type MachEventlink = Mach_eventlink_t

// MachExceptionCode is a Go-name alias for Mach_exception_code_t.
type MachExceptionCode = Mach_exception_code_t

// MachExceptionData is a Go-name alias for Mach_exception_data_t.
type MachExceptionData = Mach_exception_data_t

// MachExceptionDataType is a Go-name alias for Mach_exception_data_type_t.
type MachExceptionDataType = Mach_exception_data_type_t

// MachExceptionSubcode is a Go-name alias for Mach_exception_subcode_t.
type MachExceptionSubcode = Mach_exception_subcode_t

// MachMemoryInfoArray is a Go-name alias for Mach_memory_info_array_t.
type MachMemoryInfoArray = Mach_memory_info_array_t

// MachMemoryInfo is a Go-name alias for Mach_memory_info_t.
type MachMemoryInfo = Mach_memory_info_t

// MachMsgBase is a Go-name alias for Mach_msg_base_t.
type MachMsgBase = Mach_msg_base_t

// MachMsgBits is a Go-name alias for Mach_msg_bits_t.
type MachMsgBits = Mach_msg_bits_t

// MachMsgContextTrailer is a Go-name alias for Mach_msg_context_trailer_t.
type MachMsgContextTrailer = Mach_msg_context_trailer_t

// MachMsgCopyOptions is a Go-name alias for Mach_msg_copy_options_t.
type MachMsgCopyOptions = Mach_msg_copy_options_t

// MachMsgDescriptorType is a Go-name alias for Mach_msg_descriptor_type_t.
type MachMsgDescriptorType = Mach_msg_descriptor_type_t

// MachMsgEmptyRcv is a Go-name alias for Mach_msg_empty_rcv_t.
type MachMsgEmptyRcv = Mach_msg_empty_rcv_t

// MachMsgEmptySend is a Go-name alias for Mach_msg_empty_send_t.
type MachMsgEmptySend = Mach_msg_empty_send_t

// MachMsgFilterID is a Go-name alias for Mach_msg_filter_id.
type MachMsgFilterID = Mach_msg_filter_id

// MachMsgFormat0Trailer is a Go-name alias for Mach_msg_format_0_trailer_t.
type MachMsgFormat0Trailer = Mach_msg_format_0_trailer_t

// MachMsgGuardFlags is a Go-name alias for Mach_msg_guard_flags_t.
type MachMsgGuardFlags = Mach_msg_guard_flags_t

// MachMsgGuardedPortDescriptor32 is a Go-name alias for Mach_msg_guarded_port_descriptor32_t.
type MachMsgGuardedPortDescriptor32 = Mach_msg_guarded_port_descriptor32_t

// MachMsgGuardedPortDescriptor64 is a Go-name alias for Mach_msg_guarded_port_descriptor64_t.
type MachMsgGuardedPortDescriptor64 = Mach_msg_guarded_port_descriptor64_t

// MachMsgGuardedPortDescriptor is a Go-name alias for Mach_msg_guarded_port_descriptor_t.
type MachMsgGuardedPortDescriptor = Mach_msg_guarded_port_descriptor_t

// MachMsgHeader is a Go-name alias for Mach_msg_header_t.
type MachMsgHeader = Mach_msg_header_t

// MachMsgID is a Go-name alias for Mach_msg_id_t.
type MachMsgID = Mach_msg_id_t

// MachMsgMacTrailer is a Go-name alias for Mach_msg_mac_trailer_t.
type MachMsgMacTrailer = Mach_msg_mac_trailer_t

// MachMsgMaxTrailer is a Go-name alias for Mach_msg_max_trailer_t.
type MachMsgMaxTrailer = Mach_msg_max_trailer_t

// MachMsgOolDescriptor32 is a Go-name alias for Mach_msg_ool_descriptor32_t.
type MachMsgOolDescriptor32 = Mach_msg_ool_descriptor32_t

// MachMsgOolDescriptor64 is a Go-name alias for Mach_msg_ool_descriptor64_t.
type MachMsgOolDescriptor64 = Mach_msg_ool_descriptor64_t

// MachMsgOolPortsDescriptor32 is a Go-name alias for Mach_msg_ool_ports_descriptor32_t.
type MachMsgOolPortsDescriptor32 = Mach_msg_ool_ports_descriptor32_t

// MachMsgOolPortsDescriptor64 is a Go-name alias for Mach_msg_ool_ports_descriptor64_t.
type MachMsgOolPortsDescriptor64 = Mach_msg_ool_ports_descriptor64_t

// MachMsgOolPortsDescriptor is a Go-name alias for Mach_msg_ool_ports_descriptor_t.
type MachMsgOolPortsDescriptor = Mach_msg_ool_ports_descriptor_t

// MachMsgOption is a Go-name alias for Mach_msg_option_t.
type MachMsgOption = Mach_msg_option_t

// MachMsgOptions is a Go-name alias for Mach_msg_options_t.
type MachMsgOptions = Mach_msg_options_t

// MachMsgPriority is a Go-name alias for Mach_msg_priority_t.
type MachMsgPriority = Mach_msg_priority_t

// MachMsgReturn is a Go-name alias for Mach_msg_return_t.
type MachMsgReturn = Mach_msg_return_t

// MachMsgSecurityTrailer is a Go-name alias for Mach_msg_security_trailer_t.
type MachMsgSecurityTrailer = Mach_msg_security_trailer_t

// MachMsgSeqnoTrailer is a Go-name alias for Mach_msg_seqno_trailer_t.
type MachMsgSeqnoTrailer = Mach_msg_seqno_trailer_t

// MachMsgSize is a Go-name alias for Mach_msg_size_t.
type MachMsgSize = Mach_msg_size_t

// MachMsgTimeout is a Go-name alias for Mach_msg_timeout_t.
type MachMsgTimeout = Mach_msg_timeout_t

// MachMsgTrailerInfo is a Go-name alias for Mach_msg_trailer_info_t.
type MachMsgTrailerInfo = Mach_msg_trailer_info_t

// MachMsgTrailerSize is a Go-name alias for Mach_msg_trailer_size_t.
type MachMsgTrailerSize = Mach_msg_trailer_size_t

// MachMsgTrailerType is a Go-name alias for Mach_msg_trailer_type_t.
type MachMsgTrailerType = Mach_msg_trailer_type_t

// MachMsgTypeName is a Go-name alias for Mach_msg_type_name_t.
type MachMsgTypeName = Mach_msg_type_name_t

// MachMsgTypeNumber is a Go-name alias for Mach_msg_type_number_t.
type MachMsgTypeNumber = Mach_msg_type_number_t

// MachMsgTypeSize is a Go-name alias for Mach_msg_type_size_t.
type MachMsgTypeSize = Mach_msg_type_size_t

// MachNoSendersNotification is a Go-name alias for Mach_no_senders_notification_t.
type MachNoSendersNotification = Mach_no_senders_notification_t

// MachPortArray is a Go-name alias for Mach_port_array_t.
type MachPortArray = Mach_port_array_t

// MachPortContext is a Go-name alias for Mach_port_context_t.
type MachPortContext = Mach_port_context_t

// MachPortDeletedNotification is a Go-name alias for Mach_port_deleted_notification_t.
type MachPortDeletedNotification = Mach_port_deleted_notification_t

// MachPortDelta is a Go-name alias for Mach_port_delta_t.
type MachPortDelta = Mach_port_delta_t

// MachPortDestroyedNotification is a Go-name alias for Mach_port_destroyed_notification_t.
type MachPortDestroyedNotification = Mach_port_destroyed_notification_t

// MachPortFlavor is a Go-name alias for Mach_port_flavor_t.
type MachPortFlavor = Mach_port_flavor_t

// MachPortGuardInfo is a Go-name alias for Mach_port_guard_info_t.
type MachPortGuardInfo = Mach_port_guard_info_t

// MachPortInfoExt is a Go-name alias for Mach_port_info_ext_t.
type MachPortInfoExt = Mach_port_info_ext_t

// MachPortInfo is a Go-name alias for Mach_port_info_t.
type MachPortInfo = Mach_port_info_t

// MachPortLimits is a Go-name alias for Mach_port_limits_t.
type MachPortLimits = Mach_port_limits_t

// MachPortMscount is a Go-name alias for Mach_port_mscount_t.
type MachPortMscount = Mach_port_mscount_t

// MachPortMsgcount is a Go-name alias for Mach_port_msgcount_t.
type MachPortMsgcount = Mach_port_msgcount_t

// MachPortNameArray is a Go-name alias for Mach_port_name_array_t.
type MachPortNameArray = Mach_port_name_array_t

// MachPortName is a Go-name alias for Mach_port_name_t.
type MachPortName = Mach_port_name_t

// MachPortOptionsPtr is a Go-name alias for Mach_port_options_ptr_t.
type MachPortOptionsPtr = Mach_port_options_ptr_t

// MachPortOptions is a Go-name alias for Mach_port_options_t.
type MachPortOptions = Mach_port_options_t

// MachPortQos is a Go-name alias for Mach_port_qos_t.
type MachPortQos = Mach_port_qos_t

// MachPortRight is a Go-name alias for Mach_port_right_t.
type MachPortRight = Mach_port_right_t

// MachPortRights is a Go-name alias for Mach_port_rights_t.
type MachPortRights = Mach_port_rights_t

// MachPortSeqno is a Go-name alias for Mach_port_seqno_t.
type MachPortSeqno = Mach_port_seqno_t

// MachPortSrights is a Go-name alias for Mach_port_srights_t.
type MachPortSrights = Mach_port_srights_t

// MachPortStatus is a Go-name alias for Mach_port_status_t.
type MachPortStatus = Mach_port_status_t

// MachPort is a Go-name alias for Mach_port_t.
type MachPort = Mach_port_t

// MachPortTypeArray is a Go-name alias for Mach_port_type_array_t.
type MachPortTypeArray = Mach_port_type_array_t

// MachPortType is a Go-name alias for Mach_port_type_t.
type MachPortType = Mach_port_type_t

// MachPortUrefs is a Go-name alias for Mach_port_urefs_t.
type MachPortUrefs = Mach_port_urefs_t

// MachSendOnceNotification is a Go-name alias for Mach_send_once_notification_t.
type MachSendOnceNotification = Mach_send_once_notification_t

// MachSendPossibleNotification is a Go-name alias for Mach_send_possible_notification_t.
type MachSendPossibleNotification = Mach_send_possible_notification_t

// MachServicePortInfoData is a Go-name alias for Mach_service_port_info_data_t.
type MachServicePortInfoData = Mach_service_port_info_data_t

// MachServicePortInfo is a Go-name alias for Mach_service_port_info_t.
type MachServicePortInfo = Mach_service_port_info_t

// MachTaskBasicInfoData is a Go-name alias for Mach_task_basic_info_data_t.
type MachTaskBasicInfoData = Mach_task_basic_info_data_t

// MachTaskBasicInfo is a Go-name alias for Mach_task_basic_info_t.
type MachTaskBasicInfo = Mach_task_basic_info_t

// MachTaskFlavor is a Go-name alias for Mach_task_flavor_t.
type MachTaskFlavor = Mach_task_flavor_t

// MachThreadFlavor is a Go-name alias for Mach_thread_flavor_t.
type MachThreadFlavor = Mach_thread_flavor_t

// MachTimebaseInfoData is a Go-name alias for Mach_timebase_info_data_t.
type MachTimebaseInfoData = Mach_timebase_info_data_t

// MachTimespec is a Go-name alias for Mach_timespec_t.
type MachTimespec = Mach_timespec_t

// MachVmAddress is a Go-name alias for Mach_vm_address_t.
type MachVmAddress = Mach_vm_address_t

// MachVmAddressUt is a Go-name alias for Mach_vm_address_ut.
type MachVmAddressUt = Mach_vm_address_ut

// MachVmInfoRegion is a Go-name alias for Mach_vm_info_region_t.
type MachVmInfoRegion = Mach_vm_info_region_t

// MachVmOffset is a Go-name alias for Mach_vm_offset_t.
type MachVmOffset = Mach_vm_offset_t

// MachVmOffsetUt is a Go-name alias for Mach_vm_offset_ut.
type MachVmOffsetUt = Mach_vm_offset_ut

// MachVmRangeRecipe is a Go-name alias for Mach_vm_range_recipe_t.
type MachVmRangeRecipe = Mach_vm_range_recipe_t

// MachVmRangeRecipeV1 is a Go-name alias for Mach_vm_range_recipe_v1_t.
type MachVmRangeRecipeV1 = Mach_vm_range_recipe_v1_t

// MachVmRangeRecipeV1Ut is a Go-name alias for Mach_vm_range_recipe_v1_ut.
type MachVmRangeRecipeV1Ut = Mach_vm_range_recipe_v1_ut

// MachVmRangeRecipesRaw is a Go-name alias for Mach_vm_range_recipes_raw_t.
type MachVmRangeRecipesRaw = Mach_vm_range_recipes_raw_t

// MachVmRange is a Go-name alias for Mach_vm_range_t.
type MachVmRange = Mach_vm_range_t

// MachVmReadEntry is a Go-name alias for Mach_vm_read_entry_t.
type MachVmReadEntry = Mach_vm_read_entry_t

// MachVmSize is a Go-name alias for Mach_vm_size_t.
type MachVmSize = Mach_vm_size_t

// MachVmSizeUt is a Go-name alias for Mach_vm_size_ut.
type MachVmSizeUt = Mach_vm_size_ut

// MachVoucherAttrCommand is a Go-name alias for Mach_voucher_attr_command_t.
type MachVoucherAttrCommand = Mach_voucher_attr_command_t

// MachVoucherAttrContentSize is a Go-name alias for Mach_voucher_attr_content_size_t.
type MachVoucherAttrContentSize = Mach_voucher_attr_content_size_t

// MachVoucherAttrContent is a Go-name alias for Mach_voucher_attr_content_t.
type MachVoucherAttrContent = Mach_voucher_attr_content_t

// MachVoucherAttrControlFlags is a Go-name alias for Mach_voucher_attr_control_flags_t.
type MachVoucherAttrControlFlags = Mach_voucher_attr_control_flags_t

// MachVoucherAttrControl is a Go-name alias for Mach_voucher_attr_control_t.
type MachVoucherAttrControl = Mach_voucher_attr_control_t

// MachVoucherAttrImportanceRefs is a Go-name alias for Mach_voucher_attr_importance_refs.
type MachVoucherAttrImportanceRefs = Mach_voucher_attr_importance_refs

// MachVoucherAttrKeyArray is a Go-name alias for Mach_voucher_attr_key_array_t.
type MachVoucherAttrKeyArray = Mach_voucher_attr_key_array_t

// MachVoucherAttrKey is a Go-name alias for Mach_voucher_attr_key_t.
type MachVoucherAttrKey = Mach_voucher_attr_key_t

// MachVoucherAttrManager is a Go-name alias for Mach_voucher_attr_manager_t.
type MachVoucherAttrManager = Mach_voucher_attr_manager_t

// MachVoucherAttrRawRecipeArraySize is a Go-name alias for Mach_voucher_attr_raw_recipe_array_size_t.
type MachVoucherAttrRawRecipeArraySize = Mach_voucher_attr_raw_recipe_array_size_t

// MachVoucherAttrRawRecipeArray is a Go-name alias for Mach_voucher_attr_raw_recipe_array_t.
type MachVoucherAttrRawRecipeArray = Mach_voucher_attr_raw_recipe_array_t

// MachVoucherAttrRawRecipeSize is a Go-name alias for Mach_voucher_attr_raw_recipe_size_t.
type MachVoucherAttrRawRecipeSize = Mach_voucher_attr_raw_recipe_size_t

// MachVoucherAttrRawRecipe is a Go-name alias for Mach_voucher_attr_raw_recipe_t.
type MachVoucherAttrRawRecipe = Mach_voucher_attr_raw_recipe_t

// MachVoucherAttrRecipeCommandArray is a Go-name alias for Mach_voucher_attr_recipe_command_array_t.
type MachVoucherAttrRecipeCommandArray = Mach_voucher_attr_recipe_command_array_t

// MachVoucherAttrRecipeCommand is a Go-name alias for Mach_voucher_attr_recipe_command_t.
type MachVoucherAttrRecipeCommand = Mach_voucher_attr_recipe_command_t

// MachVoucherAttrRecipeData is a Go-name alias for Mach_voucher_attr_recipe_data_t.
type MachVoucherAttrRecipeData = Mach_voucher_attr_recipe_data_t

// MachVoucherAttrRecipeSize is a Go-name alias for Mach_voucher_attr_recipe_size_t.
type MachVoucherAttrRecipeSize = Mach_voucher_attr_recipe_size_t

// MachVoucherAttrRecipe is a Go-name alias for Mach_voucher_attr_recipe_t.
type MachVoucherAttrRecipe = Mach_voucher_attr_recipe_t

// MachVoucherAttrValueFlags is a Go-name alias for Mach_voucher_attr_value_flags_t.
type MachVoucherAttrValueFlags = Mach_voucher_attr_value_flags_t

// MachVoucherAttrValueHandleArraySize is a Go-name alias for Mach_voucher_attr_value_handle_array_size_t.
type MachVoucherAttrValueHandleArraySize = Mach_voucher_attr_value_handle_array_size_t

// MachVoucherAttrValueHandleArray is a Go-name alias for Mach_voucher_attr_value_handle_array_t.
type MachVoucherAttrValueHandleArray = Mach_voucher_attr_value_handle_array_t

// MachVoucherAttrValueHandle is a Go-name alias for Mach_voucher_attr_value_handle_t.
type MachVoucherAttrValueHandle = Mach_voucher_attr_value_handle_t

// MachVoucherAttrValueReference is a Go-name alias for Mach_voucher_attr_value_reference_t.
type MachVoucherAttrValueReference = Mach_voucher_attr_value_reference_t

// MachVoucherNameArray is a Go-name alias for Mach_voucher_name_array_t.
type MachVoucherNameArray = Mach_voucher_name_array_t

// MachVoucherName is a Go-name alias for Mach_voucher_name_t.
type MachVoucherName = Mach_voucher_name_t

// MachVoucherSelector is a Go-name alias for Mach_voucher_selector_t.
type MachVoucherSelector = Mach_voucher_selector_t

// MachVoucher is a Go-name alias for Mach_voucher_t.
type MachVoucher = Mach_voucher_t

// MachZoneInfoArray is a Go-name alias for Mach_zone_info_array_t.
type MachZoneInfoArray = Mach_zone_info_array_t

// MachZoneInfo is a Go-name alias for Mach_zone_info_t.
type MachZoneInfo = Mach_zone_info_t

// MachZoneNameArray is a Go-name alias for Mach_zone_name_array_t.
type MachZoneNameArray = Mach_zone_name_array_t

// MachZoneName is a Go-name alias for Mach_zone_name_t.
type MachZoneName = Mach_zone_name_t

// MailboxOffset is a Go-name alias for Mailbox_offset_t.
type MailboxOffset = Mailbox_offset_t

// MbClassStat is a Go-name alias for Mb_class_stat_t.
type MbClassStat = Mb_class_stat_t

// MbStat is a Go-name alias for Mb_stat_t.
type MbStat = Mb_stat_t

// Mbstate is a Go-name alias for Mbstate_t.
type Mbstate = Mbstate_t

// MbufBptr is a Go-name alias for Mbuf_bptr_t.
type MbufBptr = Mbuf_bptr_t

// MbufCsumPerformedFlags is a Go-name alias for Mbuf_csum_performed_flags_t.
type MbufCsumPerformedFlags = Mbuf_csum_performed_flags_t

// MbufCsumRequestFlags is a Go-name alias for Mbuf_csum_request_flags_t.
type MbufCsumRequestFlags = Mbuf_csum_request_flags_t

// MbufFlags is a Go-name alias for Mbuf_flags_t.
type MbufFlags = Mbuf_flags_t

// MbufHow is a Go-name alias for Mbuf_how_t.
type MbufHow = Mbuf_how_t

// MbufPtrRef is a Go-name alias for Mbuf_ptr_ref_t.
type MbufPtrRef = Mbuf_ptr_ref_t

// MbufPtr is a Go-name alias for Mbuf_ptr_t.
type MbufPtr = Mbuf_ptr_t

// MbufRefPtr is a Go-name alias for Mbuf_ref_ptr_t.
type MbufRefPtr = Mbuf_ref_ptr_t

// MbufRefRef is a Go-name alias for Mbuf_ref_ref_t.
type MbufRefRef = Mbuf_ref_ref_t

// MbufRef is a Go-name alias for Mbuf_ref_t.
type MbufRef = Mbuf_ref_t

// Mbuf is a Go-name alias for Mbuf_t.
type Mbuf = Mbuf_t

// MbufTagID is a Go-name alias for Mbuf_tag_id_t.
type MbufTagID = Mbuf_tag_id_t

// MbufTagType is a Go-name alias for Mbuf_tag_type_t.
type MbufTagType = Mbuf_tag_type_t

// MbufTrafficClass is a Go-name alias for Mbuf_traffic_class_t.
type MbufTrafficClass = Mbuf_traffic_class_t

// MbufTsoRequestFlags is a Go-name alias for Mbuf_tso_request_flags_t.
type MbufTsoRequestFlags = Mbuf_tso_request_flags_t

// MccEccEvent is a Go-name alias for Mcc_ecc_event_t.
type MccEccEvent = Mcc_ecc_event_t

// MccEccVersion is a Go-name alias for Mcc_ecc_version_t.
type MccEccVersion = Mcc_ecc_version_t

// MccFlags is a Go-name alias for Mcc_flags_t.
type MccFlags = Mcc_flags_t

// Mcontext is a Go-name alias for Mcontext_t.
type Mcontext = Mcontext_t

// MemEntryNamePort is a Go-name alias for Mem_entry_name_port_t.
type MemEntryNamePort = Mem_entry_name_port_t

// MemoryObjectArray is a Go-name alias for Memory_object_array_t.
type MemoryObjectArray = Memory_object_array_t

// MemoryObjectAttrInfoData is a Go-name alias for Memory_object_attr_info_data_t.
type MemoryObjectAttrInfoData = Memory_object_attr_info_data_t

// MemoryObjectAttrInfo is a Go-name alias for Memory_object_attr_info_t.
type MemoryObjectAttrInfo = Memory_object_attr_info_t

// MemoryObjectBehaveInfoData is a Go-name alias for Memory_object_behave_info_data_t.
type MemoryObjectBehaveInfoData = Memory_object_behave_info_data_t

// MemoryObjectBehaveInfo is a Go-name alias for Memory_object_behave_info_t.
type MemoryObjectBehaveInfo = Memory_object_behave_info_t

// MemoryObjectClusterSize is a Go-name alias for Memory_object_cluster_size_t.
type MemoryObjectClusterSize = Memory_object_cluster_size_t

// MemoryObjectControl is a Go-name alias for Memory_object_control_t.
type MemoryObjectControl = Memory_object_control_t

// MemoryObjectCopyStrategy is a Go-name alias for Memory_object_copy_strategy_t.
type MemoryObjectCopyStrategy = Memory_object_copy_strategy_t

// MemoryObjectDefault is a Go-name alias for Memory_object_default_t.
type MemoryObjectDefault = Memory_object_default_t

// MemoryObjectFaultInfo is a Go-name alias for Memory_object_fault_info_t.
type MemoryObjectFaultInfo = Memory_object_fault_info_t

// MemoryObjectFlavor is a Go-name alias for Memory_object_flavor_t.
type MemoryObjectFlavor = Memory_object_flavor_t

// MemoryObjectInfoData is a Go-name alias for Memory_object_info_data_t.
type MemoryObjectInfoData = Memory_object_info_data_t

// MemoryObjectInfo is a Go-name alias for Memory_object_info_t.
type MemoryObjectInfo = Memory_object_info_t

// MemoryObjectName is a Go-name alias for Memory_object_name_t.
type MemoryObjectName = Memory_object_name_t

// MemoryObjectOffset is a Go-name alias for Memory_object_offset_t.
type MemoryObjectOffset = Memory_object_offset_t

// MemoryObjectOffsetUt is a Go-name alias for Memory_object_offset_ut.
type MemoryObjectOffsetUt = Memory_object_offset_ut

// MemoryObjectPerfInfoData is a Go-name alias for Memory_object_perf_info_data_t.
type MemoryObjectPerfInfoData = Memory_object_perf_info_data_t

// MemoryObjectPerfInfo is a Go-name alias for Memory_object_perf_info_t.
type MemoryObjectPerfInfo = Memory_object_perf_info_t

// MemoryObjectReturn is a Go-name alias for Memory_object_return_t.
type MemoryObjectReturn = Memory_object_return_t

// MemoryObjectSize is a Go-name alias for Memory_object_size_t.
type MemoryObjectSize = Memory_object_size_t

// MemoryObjectSizeUt is a Go-name alias for Memory_object_size_ut.
type MemoryObjectSizeUt = Memory_object_size_ut

// MemoryObject is a Go-name alias for Memory_object_t.
type MemoryObject = Memory_object_t

// MigImplRoutine is a Go-name alias for Mig_impl_routine_t.
type MigImplRoutine = Mig_impl_routine_t

// MigReplyError is a Go-name alias for Mig_reply_error_t.
type MigReplyError = Mig_reply_error_t

// MigRoutineArgDescriptor is a Go-name alias for Mig_routine_arg_descriptor_t.
type MigRoutineArgDescriptor = Mig_routine_arg_descriptor_t

// MigRoutineDescriptor is a Go-name alias for Mig_routine_descriptor.
type MigRoutineDescriptor = Mig_routine_descriptor

// MigRoutine is a Go-name alias for Mig_routine_t.
type MigRoutine = Mig_routine_t

// MigSubsystem is a Go-name alias for Mig_subsystem_t.
type MigSubsystem = Mig_subsystem_t

// MigSymtab is a Go-name alias for Mig_symtab_t.
type MigSymtab = Mig_symtab_t

// MlCPUInfo is a Go-name alias for Ml_cpu_info_t.
type MlCPUInfo = Ml_cpu_info_t

// MlPageProtection is a Go-name alias for Ml_page_protection_t.
type MlPageProtection = Ml_page_protection_t

// MlProcessorInfo is a Go-name alias for Ml_processor_info_t.
type MlProcessorInfo = Ml_processor_info_t

// MmapFcn is a Go-name alias for Mmap_fcn_t.
type MmapFcn = Mmap_fcn_t

// Mode is a Go-name alias for Mode_t.
type Mode = Mode_t

// MountBptr is a Go-name alias for Mount_bptr_t.
type MountBptr = Mount_bptr_t

// MountPtrRef is a Go-name alias for Mount_ptr_ref_t.
type MountPtrRef = Mount_ptr_ref_t

// MountPtr is a Go-name alias for Mount_ptr_t.
type MountPtr = Mount_ptr_t

// MountRefPtr is a Go-name alias for Mount_ref_ptr_t.
type MountRefPtr = Mount_ref_ptr_t

// MountRefRef is a Go-name alias for Mount_ref_ref_t.
type MountRefRef = Mount_ref_ref_t

// MountRef is a Go-name alias for Mount_ref_t.
type MountRef = Mount_ref_t

// Mount is a Go-name alias for Mount_t.
type Mount = Mount_t

// MphPanicFlags is a Go-name alias for Mph_panic_flags_t.
type MphPanicFlags = Mph_panic_flags_t

// MpscQueueChain is a Go-name alias for Mpsc_queue_chain_t.
type MpscQueueChain = Mpsc_queue_chain_t

// MpscQueueHead is a Go-name alias for Mpsc_queue_head_t.
type MpscQueueHead = Mpsc_queue_head_t

// MsgLabels is a Go-name alias for Msg_labels_t.
type MsgLabels = Msg_labels_t

// Msglen is a Go-name alias for Msglen_t.
type Msglen = Msglen_t

// Msgqnum is a Go-name alias for Msgqnum_t.
type Msgqnum = Msgqnum_t

// NLong is a Go-name alias for N_long.
type NLong = N_long

// NShort is a Go-name alias for N_short.
type NShort = N_short

// NTime is a Go-name alias for N_time.
type NTime = N_time

// Natural is a Go-name alias for Natural_t.
type Natural = Natural_t

// NetInitFuncPtr is a Go-name alias for Net_init_func_ptr.
type NetInitFuncPtr = Net_init_func_ptr

// Netaddr is a Go-name alias for Netaddr_t.
type Netaddr = Netaddr_t

// NetworkPort is a Go-name alias for Network_port_t.
type NetworkPort = Network_port_t

// NfsFsid is a Go-name alias for Nfs_fsid.
type NfsFsid = Nfs_fsid

// NfsHandle is a Go-name alias for Nfs_handle.
type NfsHandle = Nfs_handle

// NfsSpecdata is a Go-name alias for Nfs_specdata.
type NfsSpecdata = Nfs_specdata

// NfsStateid is a Go-name alias for Nfs_stateid.
type NfsStateid = Nfs_stateid

// NfsSupportedKerberosEtypes is a Go-name alias for Nfs_supported_kerberos_etypes.
type NfsSupportedKerberosEtypes = Nfs_supported_kerberos_etypes

// NfserrInfo is a Go-name alias for Nfserr_info_t.
type NfserrInfo = Nfserr_info_t

// Nlink is a Go-name alias for Nlink_t.
type Nlink = Nlink_t

// NotifyPort is a Go-name alias for Notify_port_t.
type NotifyPort = Notify_port_t

// NpUid is a Go-name alias for Np_uid_t.
type NpUid = Np_uid_t

// NspaceName is a Go-name alias for Nspace_name_t.
type NspaceName = Nspace_name_t

// NspacePath is a Go-name alias for Nspace_path_t.
type NspacePath = Nspace_path_t

// Ntsid is a Go-name alias for Ntsid_t.
type Ntsid = Ntsid_t

// Off is a Go-name alias for Off_t.
type Off = Off_t

// OpenCloseFcn is a Go-name alias for Open_close_fcn_t.
type OpenCloseFcn = Open_close_fcn_t

// OSBlock is a Go-name alias for Os_block_t.
type OSBlock = Os_block_t

// OSLogCoprocReg is a Go-name alias for Os_log_coproc_reg_t.
type OSLogCoprocReg = Os_log_coproc_reg_t

// OSLog is a Go-name alias for Os_log_t.
type OSLog = Os_log_t

// OSLogType is a Go-name alias for Os_log_type_t.
type OSLogType = Os_log_type_t

// PackedUchar16 is a Go-name alias for Packed_uchar16.
type PackedUchar16 = Packed_uchar16

// PackedUchar32 is a Go-name alias for Packed_uchar32.
type PackedUchar32 = Packed_uchar32

// PackedUchar64 is a Go-name alias for Packed_uchar64.
type PackedUchar64 = Packed_uchar64

// PackedUshort4 is a Go-name alias for Packed_ushort4.
type PackedUshort4 = Packed_ushort4

// PageAddressArray is a Go-name alias for Page_address_array_t.
type PageAddressArray = Page_address_array_t

// Pid is a Go-name alias for Pid_t.
type Pid = Pid_t

// PkthdrBptr is a Go-name alias for Pkthdr_bptr_t.
type PkthdrBptr = Pkthdr_bptr_t

// PkthdrPtrRef is a Go-name alias for Pkthdr_ptr_ref_t.
type PkthdrPtrRef = Pkthdr_ptr_ref_t

// PkthdrPtr is a Go-name alias for Pkthdr_ptr_t.
type PkthdrPtr = Pkthdr_ptr_t

// PkthdrRefPtr is a Go-name alias for Pkthdr_ref_ptr_t.
type PkthdrRefPtr = Pkthdr_ref_ptr_t

// PkthdrRefRef is a Go-name alias for Pkthdr_ref_ref_t.
type PkthdrRefRef = Pkthdr_ref_ref_t

// PkthdrRef is a Go-name alias for Pkthdr_ref_t.
type PkthdrRef = Pkthdr_ref_t

// Pkthdr is a Go-name alias for Pkthdr_t.
type Pkthdr = Pkthdr_t

// Pointer is a Go-name alias for Pointer_t.
type Pointer = Pointer_t

// PointerUt is a Go-name alias for Pointer_ut.
type PointerUt = Pointer_ut

// PolicyBaseData is a Go-name alias for Policy_base_data_t.
type PolicyBaseData = Policy_base_data_t

// PolicyBase is a Go-name alias for Policy_base_t.
type PolicyBase = Policy_base_t

// PolicyFifoBaseData is a Go-name alias for Policy_fifo_base_data_t.
type PolicyFifoBaseData = Policy_fifo_base_data_t

// PolicyFifoBase is a Go-name alias for Policy_fifo_base_t.
type PolicyFifoBase = Policy_fifo_base_t

// PolicyFifoInfoData is a Go-name alias for Policy_fifo_info_data_t.
type PolicyFifoInfoData = Policy_fifo_info_data_t

// PolicyFifoInfo is a Go-name alias for Policy_fifo_info_t.
type PolicyFifoInfo = Policy_fifo_info_t

// PolicyFifoLimitData is a Go-name alias for Policy_fifo_limit_data_t.
type PolicyFifoLimitData = Policy_fifo_limit_data_t

// PolicyFifoLimit is a Go-name alias for Policy_fifo_limit_t.
type PolicyFifoLimit = Policy_fifo_limit_t

// PolicyInfoData is a Go-name alias for Policy_info_data_t.
type PolicyInfoData = Policy_info_data_t

// PolicyInfo is a Go-name alias for Policy_info_t.
type PolicyInfo = Policy_info_t

// PolicyLimitData is a Go-name alias for Policy_limit_data_t.
type PolicyLimitData = Policy_limit_data_t

// PolicyLimit is a Go-name alias for Policy_limit_t.
type PolicyLimit = Policy_limit_t

// PolicyRrBaseData is a Go-name alias for Policy_rr_base_data_t.
type PolicyRrBaseData = Policy_rr_base_data_t

// PolicyRrBase is a Go-name alias for Policy_rr_base_t.
type PolicyRrBase = Policy_rr_base_t

// PolicyRrInfoData is a Go-name alias for Policy_rr_info_data_t.
type PolicyRrInfoData = Policy_rr_info_data_t

// PolicyRrInfo is a Go-name alias for Policy_rr_info_t.
type PolicyRrInfo = Policy_rr_info_t

// PolicyRrLimitData is a Go-name alias for Policy_rr_limit_data_t.
type PolicyRrLimitData = Policy_rr_limit_data_t

// PolicyRrLimit is a Go-name alias for Policy_rr_limit_t.
type PolicyRrLimit = Policy_rr_limit_t

// Policy is a Go-name alias for Policy_t.
type Policy = Policy_t

// PolicyTimeshareBaseData is a Go-name alias for Policy_timeshare_base_data_t.
type PolicyTimeshareBaseData = Policy_timeshare_base_data_t

// PolicyTimeshareBase is a Go-name alias for Policy_timeshare_base_t.
type PolicyTimeshareBase = Policy_timeshare_base_t

// PolicyTimeshareInfoData is a Go-name alias for Policy_timeshare_info_data_t.
type PolicyTimeshareInfoData = Policy_timeshare_info_data_t

// PolicyTimeshareInfo is a Go-name alias for Policy_timeshare_info_t.
type PolicyTimeshareInfo = Policy_timeshare_info_t

// PolicyTimeshareLimitData is a Go-name alias for Policy_timeshare_limit_data_t.
type PolicyTimeshareLimitData = Policy_timeshare_limit_data_t

// PolicyTimeshareLimit is a Go-name alias for Policy_timeshare_limit_t.
type PolicyTimeshareLimit = Policy_timeshare_limit_t

// PortNameArray is a Go-name alias for Port_name_array_t.
type PortNameArray = Port_name_array_t

// PortName is a Go-name alias for Port_name_t.
type PortName = Port_name_t

// Port is a Go-name alias for Port_t.
type Port = Port_t

// PosixCred is a Go-name alias for Posix_cred_t.
type PosixCred = Posix_cred_t

// Ppnum is a Go-name alias for Ppnum_t.
type Ppnum = Ppnum_t

// PriorityQueueCompareFn is a Go-name alias for Priority_queue_compare_fn_t.
type PriorityQueueCompareFn = Priority_queue_compare_fn_t

// PriorityQueueEntryDeadline is a Go-name alias for Priority_queue_entry_deadline_t.
type PriorityQueueEntryDeadline = Priority_queue_entry_deadline_t

// PriorityQueueEntrySchedModifier is a Go-name alias for Priority_queue_entry_sched_modifier_t.
type PriorityQueueEntrySchedModifier = Priority_queue_entry_sched_modifier_t

// PriorityQueueEntrySched is a Go-name alias for Priority_queue_entry_sched_t.
type PriorityQueueEntrySched = Priority_queue_entry_sched_t

// PriorityQueueEntryStable is a Go-name alias for Priority_queue_entry_stable_t.
type PriorityQueueEntryStable = Priority_queue_entry_stable_t

// PriorityQueueKey is a Go-name alias for Priority_queue_key_t.
type PriorityQueueKey = Priority_queue_key_t

// ProcBptr is a Go-name alias for Proc_bptr_t.
type ProcBptr = Proc_bptr_t

// ProcIdentBptr is a Go-name alias for Proc_ident_bptr_t.
type ProcIdentBptr = Proc_ident_bptr_t

// ProcIdentPtrRef is a Go-name alias for Proc_ident_ptr_ref_t.
type ProcIdentPtrRef = Proc_ident_ptr_ref_t

// ProcIdentPtr is a Go-name alias for Proc_ident_ptr_t.
type ProcIdentPtr = Proc_ident_ptr_t

// ProcIdentRefPtr is a Go-name alias for Proc_ident_ref_ptr_t.
type ProcIdentRefPtr = Proc_ident_ref_ptr_t

// ProcIdentRefRef is a Go-name alias for Proc_ident_ref_ref_t.
type ProcIdentRefRef = Proc_ident_ref_ref_t

// ProcIdentRef is a Go-name alias for Proc_ident_ref_t.
type ProcIdentRef = Proc_ident_ref_t

// ProcIdent is a Go-name alias for Proc_ident_t.
type ProcIdent = Proc_ident_t

// ProcPtrRef is a Go-name alias for Proc_ptr_ref_t.
type ProcPtrRef = Proc_ptr_ref_t

// ProcPtr is a Go-name alias for Proc_ptr_t.
type ProcPtr = Proc_ptr_t

// ProcRefPtr is a Go-name alias for Proc_ref_ptr_t.
type ProcRefPtr = Proc_ref_ptr_t

// ProcRefRef is a Go-name alias for Proc_ref_ref_t.
type ProcRefRef = Proc_ref_ref_t

// ProcRef is a Go-name alias for Proc_ref_t.
type ProcRef = Proc_ref_t

// Proc is a Go-name alias for Proc_t.
type Proc = Proc_t

// ProcessorArray is a Go-name alias for Processor_array_t.
type ProcessorArray = Processor_array_t

// ProcessorBasicInfoData is a Go-name alias for Processor_basic_info_data_t.
type ProcessorBasicInfoData = Processor_basic_info_data_t

// ProcessorBasicInfo is a Go-name alias for Processor_basic_info_t.
type ProcessorBasicInfo = Processor_basic_info_t

// ProcessorCPULoadInfoData is a Go-name alias for Processor_cpu_load_info_data_t.
type ProcessorCPULoadInfoData = Processor_cpu_load_info_data_t

// ProcessorCPULoadInfo is a Go-name alias for Processor_cpu_load_info_t.
type ProcessorCPULoadInfo = Processor_cpu_load_info_t

// ProcessorCPUStat64Data is a Go-name alias for Processor_cpu_stat64_data_t.
type ProcessorCPUStat64Data = Processor_cpu_stat64_data_t

// ProcessorCPUStat64 is a Go-name alias for Processor_cpu_stat64_t.
type ProcessorCPUStat64 = Processor_cpu_stat64_t

// ProcessorCPUStatData is a Go-name alias for Processor_cpu_stat_data_t.
type ProcessorCPUStatData = Processor_cpu_stat_data_t

// ProcessorCPUStat is a Go-name alias for Processor_cpu_stat_t.
type ProcessorCPUStat = Processor_cpu_stat_t

// ProcessorFlavor is a Go-name alias for Processor_flavor_t.
type ProcessorFlavor = Processor_flavor_t

// ProcessorInfoArray is a Go-name alias for Processor_info_array_t.
type ProcessorInfoArray = Processor_info_array_t

// ProcessorInfoData is a Go-name alias for Processor_info_data_t.
type ProcessorInfoData = Processor_info_data_t

// ProcessorInfo is a Go-name alias for Processor_info_t.
type ProcessorInfo = Processor_info_t

// ProcessorPortArray is a Go-name alias for Processor_port_array_t.
type ProcessorPortArray = Processor_port_array_t

// ProcessorPort is a Go-name alias for Processor_port_t.
type ProcessorPort = Processor_port_t

// ProcessorSetArray is a Go-name alias for Processor_set_array_t.
type ProcessorSetArray = Processor_set_array_t

// ProcessorSetBasicInfoData is a Go-name alias for Processor_set_basic_info_data_t.
type ProcessorSetBasicInfoData = Processor_set_basic_info_data_t

// ProcessorSetBasicInfo is a Go-name alias for Processor_set_basic_info_t.
type ProcessorSetBasicInfo = Processor_set_basic_info_t

// ProcessorSetControlPort is a Go-name alias for Processor_set_control_port_t.
type ProcessorSetControlPort = Processor_set_control_port_t

// ProcessorSetControl is a Go-name alias for Processor_set_control_t.
type ProcessorSetControl = Processor_set_control_t

// ProcessorSetFlavor is a Go-name alias for Processor_set_flavor_t.
type ProcessorSetFlavor = Processor_set_flavor_t

// ProcessorSetInfoData is a Go-name alias for Processor_set_info_data_t.
type ProcessorSetInfoData = Processor_set_info_data_t

// ProcessorSetInfo is a Go-name alias for Processor_set_info_t.
type ProcessorSetInfo = Processor_set_info_t

// ProcessorSetLoadInfoData is a Go-name alias for Processor_set_load_info_data_t.
type ProcessorSetLoadInfoData = Processor_set_load_info_data_t

// ProcessorSetLoadInfo is a Go-name alias for Processor_set_load_info_t.
type ProcessorSetLoadInfo = Processor_set_load_info_t

// ProcessorSetNameArray is a Go-name alias for Processor_set_name_array_t.
type ProcessorSetNameArray = Processor_set_name_array_t

// ProcessorSetNamePortArray is a Go-name alias for Processor_set_name_port_array_t.
type ProcessorSetNamePortArray = Processor_set_name_port_array_t

// ProcessorSetNamePort is a Go-name alias for Processor_set_name_port_t.
type ProcessorSetNamePort = Processor_set_name_port_t

// ProcessorSetName is a Go-name alias for Processor_set_name_t.
type ProcessorSetName = Processor_set_name_t

// ProcessorSetPort is a Go-name alias for Processor_set_port_t.
type ProcessorSetPort = Processor_set_port_t

// ProcessorSet is a Go-name alias for Processor_set_t.
type ProcessorSet = Processor_set_t

// ProtocolFamily is a Go-name alias for Protocol_family_t.
type ProtocolFamily = Protocol_family_t

// PsizeFcn is a Go-name alias for Psize_fcn_t.
type PsizeFcn = Psize_fcn_t

// Ptrdiff is a Go-name alias for Ptrdiff_t.
type Ptrdiff = Ptrdiff_t

// Qaddr is a Go-name alias for Qaddr_t.
type Qaddr = Qaddr_t

// Quad is a Go-name alias for Quad_t.
type Quad = Quad_t

// QueueChain is a Go-name alias for Queue_chain_t.
type QueueChain = Queue_chain_t

// QueueEntry is a Go-name alias for Queue_entry_t.
type QueueEntry = Queue_entry_t

// QueueHead is a Go-name alias for Queue_head_t.
type QueueHead = Queue_head_t

// Queue is a Go-name alias for Queue_t.
type Queue = Queue_t

// ReadWriteFcn is a Go-name alias for Read_write_fcn_t.
type ReadWriteFcn = Read_write_fcn_t

// Reg64 is a Go-name alias for Reg64_t.
type Reg64 = Reg64_t

// Register is a Go-name alias for Register_t.
type Register = Register_t

// ResetFcn is a Go-name alias for Reset_fcn_t.
type ResetFcn = Reset_fcn_t

// Rlim is a Go-name alias for Rlim_t.
type Rlim = Rlim_t

// Route is a Go-name alias for Route_t.
type Route = Route_t

// RoutineArgDescriptor is a Go-name alias for Routine_arg_descriptor.
type RoutineArgDescriptor = Routine_arg_descriptor

// RoutineArgOffset is a Go-name alias for Routine_arg_offset.
type RoutineArgOffset = Routine_arg_offset

// RoutineArgSize is a Go-name alias for Routine_arg_size.
type RoutineArgSize = Routine_arg_size

// RoutineArgType is a Go-name alias for Routine_arg_type.
type RoutineArgType = Routine_arg_type

// RoutineDescriptor is a Go-name alias for Routine_descriptor_t.
type RoutineDescriptor = Routine_descriptor_t

// RpcRoutineArgDescriptor is a Go-name alias for Rpc_routine_arg_descriptor_t.
type RpcRoutineArgDescriptor = Rpc_routine_arg_descriptor_t

// RpcRoutineDescriptor is a Go-name alias for Rpc_routine_descriptor_t.
type RpcRoutineDescriptor = Rpc_routine_descriptor_t

// RpcSubsystem is a Go-name alias for Rpc_subsystem_t.
type RpcSubsystem = Rpc_subsystem_t

// Rsize is a Go-name alias for Rsize_t.
type Rsize = Rsize_t

// RsvdFcn is a Go-name alias for Rsvd_fcn_t.
type RsvdFcn = Rsvd_fcn_t

// RtentryBptr is a Go-name alias for Rtentry_bptr_t.
type RtentryBptr = Rtentry_bptr_t

// RtentryPtrRef is a Go-name alias for Rtentry_ptr_ref_t.
type RtentryPtrRef = Rtentry_ptr_ref_t

// RtentryPtr is a Go-name alias for Rtentry_ptr_t.
type RtentryPtr = Rtentry_ptr_t

// RtentryRefPtr is a Go-name alias for Rtentry_ref_ptr_t.
type RtentryRefPtr = Rtentry_ref_ptr_t

// RtentryRefRef is a Go-name alias for Rtentry_ref_ref_t.
type RtentryRefRef = Rtentry_ref_ref_t

// RtentryRef is a Go-name alias for Rtentry_ref_t.
type RtentryRef = Rtentry_ref_t

// Rune is a Go-name alias for Rune_t.
type Rune = Rune_t

// RusageInfoCurrent is a Go-name alias for Rusage_info_current.
type RusageInfoCurrent = Rusage_info_current

// RusageInfo is a Go-name alias for Rusage_info_t.
type RusageInfo = Rusage_info_t

// SaEndpoints is a Go-name alias for Sa_endpoints_t.
type SaEndpoints = Sa_endpoints_t

// SaFamily is a Go-name alias for Sa_family_t.
type SaFamily = Sa_family_t

// SaeAssocid is a Go-name alias for Sae_associd_t.
type SaeAssocid = Sae_associd_t

// SaeConnid is a Go-name alias for Sae_connid_t.
type SaeConnid = Sae_connid_t

// SecureBootCryptexArgs is a Go-name alias for Secure_boot_cryptex_args_t.
type SecureBootCryptexArgs = Secure_boot_cryptex_args_t

// SecurityToken is a Go-name alias for Security_token_t.
type SecurityToken = Security_token_t

// Segsz is a Go-name alias for Segsz_t.
type Segsz = Segsz_t

// Sel is a Go-name alias for Sel_t.
type Sel = Sel_t

// SelectFcn is a Go-name alias for Select_fcn_t.
type SelectFcn = Select_fcn_t

// SemaphorePort is a Go-name alias for Semaphore_port_t.
type SemaphorePort = Semaphore_port_t

// Semaphore is a Go-name alias for Semaphore_t.
type Semaphore = Semaphore_t

// SfltDataFlag is a Go-name alias for Sflt_data_flag_t.
type SfltDataFlag = Sflt_data_flag_t

// SfltEvent is a Go-name alias for Sflt_event_t.
type SfltEvent = Sflt_event_t

// SfltFlags is a Go-name alias for Sflt_flags.
type SfltFlags = Sflt_flags

// SfltHandle is a Go-name alias for Sflt_handle.
type SfltHandle = Sflt_handle

// SharedFileMappingSlideNp is a Go-name alias for Shared_file_mapping_slide_np_t.
type SharedFileMappingSlideNp = Shared_file_mapping_slide_np_t

// SharedFileMappingSlideNpUt is a Go-name alias for Shared_file_mapping_slide_np_ut.
type SharedFileMappingSlideNpUt = Shared_file_mapping_slide_np_ut

// Shmatt is a Go-name alias for Shmatt_t.
type Shmatt = Shmatt_t

// SigAtomic is a Go-name alias for Sig_atomic_t.
type SigAtomic = Sig_atomic_t

// Sig is a Go-name alias for Sig_t.
type Sig = Sig_t

// Siginfo is a Go-name alias for Siginfo_t.
type Siginfo = Siginfo_t

// Sigset is a Go-name alias for Sigset_t.
type Sigset = Sigset_t

// Size is a Go-name alias for Size_t.
type Size = Size_t

// SizeUt is a Go-name alias for Size_ut.
type SizeUt = Size_ut

// SleepType is a Go-name alias for Sleep_type_t.
type SleepType = Sleep_type_t

// SmrCb is a Go-name alias for Smr_cb_t.
type SmrCb = Smr_cb_t

// SmrNode is a Go-name alias for Smr_node_t.
type SmrNode = Smr_node_t

// SmrSeq is a Go-name alias for Smr_seq_t.
type SmrSeq = Smr_seq_t

// Smr is a Go-name alias for Smr_t.
type Smr = Smr_t

// SoGen is a Go-name alias for So_gen_t.
type SoGen = So_gen_t

// SockStorage is a Go-name alias for Sock_storage.
type SockStorage = Sock_storage

// SockaddrPtrRef is a Go-name alias for Sockaddr_ptr_ref_t.
type SockaddrPtrRef = Sockaddr_ptr_ref_t

// SockaddrRefPtr is a Go-name alias for Sockaddr_ref_ptr_t.
type SockaddrRefPtr = Sockaddr_ref_ptr_t

// SockaddrRefRef is a Go-name alias for Sockaddr_ref_ref_t.
type SockaddrRefRef = Sockaddr_ref_ref_t

// SockaddrRef is a Go-name alias for Sockaddr_ref_t.
type SockaddrRef = Sockaddr_ref_t

// SockaddrStoragePtrRef is a Go-name alias for Sockaddr_storage_ptr_ref_t.
type SockaddrStoragePtrRef = Sockaddr_storage_ptr_ref_t

// SockaddrStorageRefPtr is a Go-name alias for Sockaddr_storage_ref_ptr_t.
type SockaddrStorageRefPtr = Sockaddr_storage_ref_ptr_t

// SockaddrStorageRefRef is a Go-name alias for Sockaddr_storage_ref_ref_t.
type SockaddrStorageRefRef = Sockaddr_storage_ref_ref_t

// SockaddrStorageRef is a Go-name alias for Sockaddr_storage_ref_t.
type SockaddrStorageRef = Sockaddr_storage_ref_t

// SocketBptr is a Go-name alias for Socket_bptr_t.
type SocketBptr = Socket_bptr_t

// SocketPtrRef is a Go-name alias for Socket_ptr_ref_t.
type SocketPtrRef = Socket_ptr_ref_t

// SocketPtr is a Go-name alias for Socket_ptr_t.
type SocketPtr = Socket_ptr_t

// SocketRefPtr is a Go-name alias for Socket_ref_ptr_t.
type SocketRefPtr = Socket_ref_ptr_t

// SocketRefRef is a Go-name alias for Socket_ref_ref_t.
type SocketRefRef = Socket_ref_ref_t

// SocketRef is a Go-name alias for Socket_ref_t.
type SocketRef = Socket_ref_t

// Socket is a Go-name alias for Socket_t.
type Socket = Socket_t

// Socklen is a Go-name alias for Socklen_t.
type Socklen = Socklen_t

// SockoptBptr is a Go-name alias for Sockopt_bptr_t.
type SockoptBptr = Sockopt_bptr_t

// SockoptDir is a Go-name alias for Sockopt_dir.
type SockoptDir = Sockopt_dir

// SockoptPtrRef is a Go-name alias for Sockopt_ptr_ref_t.
type SockoptPtrRef = Sockopt_ptr_ref_t

// SockoptPtr is a Go-name alias for Sockopt_ptr_t.
type SockoptPtr = Sockopt_ptr_t

// SockoptRefPtr is a Go-name alias for Sockopt_ref_ptr_t.
type SockoptRefPtr = Sockopt_ref_ptr_t

// SockoptRefRef is a Go-name alias for Sockopt_ref_ref_t.
type SockoptRefRef = Sockopt_ref_ref_t

// SockoptRef is a Go-name alias for Sockopt_ref_t.
type SockoptRef = Sockopt_ref_t

// Speed is a Go-name alias for Speed_t.
type Speed = Speed_t

// SptmAsid is a Go-name alias for Sptm_asid_t.
type SptmAsid = Sptm_asid_t

// SptmCallRegs is a Go-name alias for Sptm_call_regs_t.
type SptmCallRegs = Sptm_call_regs_t

// SptmConsistentDebug is a Go-name alias for Sptm_consistent_debug_t.
type SptmConsistentDebug = Sptm_consistent_debug_t

// SptmDispatchEndpointID is a Go-name alias for Sptm_dispatch_endpoint_id_t.
type SptmDispatchEndpointID = Sptm_dispatch_endpoint_id_t

// SptmDispatchTableID is a Go-name alias for Sptm_dispatch_table_id_t.
type SptmDispatchTableID = Sptm_dispatch_table_id_t

// SptmDispatchTarget is a Go-name alias for Sptm_dispatch_target_t.
type SptmDispatchTarget = Sptm_dispatch_target_t

// SptmDomain is a Go-name alias for Sptm_domain_t.
type SptmDomain = Sptm_domain_t

// SptmFrameType is a Go-name alias for Sptm_frame_type_t.
type SptmFrameType = Sptm_frame_type_t

// SptmInstanceID is a Go-name alias for Sptm_instance_id_t.
type SptmInstanceID = Sptm_instance_id_t

// SptmIommuID is a Go-name alias for Sptm_iommu_id_t.
type SptmIommuID = Sptm_iommu_id_t

// SptmIommuRetypeParams is a Go-name alias for Sptm_iommu_retype_params_t.
type SptmIommuRetypeParams = Sptm_iommu_retype_params_t

// SptmPaddr is a Go-name alias for Sptm_paddr_t.
type SptmPaddr = Sptm_paddr_t

// SptmPapt is a Go-name alias for Sptm_papt_t.
type SptmPapt = Sptm_papt_t

// SptmPoff is a Go-name alias for Sptm_poff_t.
type SptmPoff = Sptm_poff_t

// SptmPpnum is a Go-name alias for Sptm_ppnum_t.
type SptmPpnum = Sptm_ppnum_t

// SptmPtLevel is a Go-name alias for Sptm_pt_level_t.
type SptmPtLevel = Sptm_pt_level_t

// SptmPte is a Go-name alias for Sptm_pte_t.
type SptmPte = Sptm_pte_t

// SptmReturn is a Go-name alias for Sptm_return_t.
type SptmReturn = Sptm_return_t

// SptmRetypeParams is a Go-name alias for Sptm_retype_params_t.
type SptmRetypeParams = Sptm_retype_params_t

// SptmTraceBuffer is a Go-name alias for Sptm_trace_buffer_t.
type SptmTraceBuffer = Sptm_trace_buffer_t

// SptmTrace is a Go-name alias for Sptm_trace_t.
type SptmTrace = Sptm_trace_t

// SptmTte is a Go-name alias for Sptm_tte_t.
type SptmTte = Sptm_tte_t

// SptmVaddr is a Go-name alias for Sptm_vaddr_t.
type SptmVaddr = Sptm_vaddr_t

// SptmVectorType is a Go-name alias for Sptm_vector_type_t.
type SptmVectorType = Sptm_vector_type_t

// SptmVmid is a Go-name alias for Sptm_vmid_t.
type SptmVmid = Sptm_vmid_t

// SptmVoff is a Go-name alias for Sptm_voff_t.
type SptmVoff = Sptm_voff_t

// Ssize is a Go-name alias for Ssize_t.
type Ssize = Ssize_t

// Stack is a Go-name alias for Stack_t.
type Stack = Stack_t

// StopFcn is a Go-name alias for Stop_fcn_t.
type StopFcn = Stop_fcn_t

// StrategyFcn is a Go-name alias for Strategy_fcn_t.
type StrategyFcn = Strategy_fcn_t

// String is a Go-name alias for String_t.
type String = String_t

// Subaid is a Go-name alias for Subaid_t.
type Subaid = Subaid_t

// Suseconds is a Go-name alias for Suseconds_t.
type Suseconds = Suseconds_t

// Swblk is a Go-name alias for Swblk_t.
type Swblk = Swblk_t

// SymtabName is a Go-name alias for Symtab_name_t.
type SymtabName = Symtab_name_t

// SyncPolicy is a Go-name alias for Sync_policy_t.
type SyncPolicy = Sync_policy_t

// SyscallArg is a Go-name alias for Syscall_arg_t.
type SyscallArg = Syscall_arg_t

// SyscpIDInstructionsFeat1Reg is a Go-name alias for Syscp_ID_instructions_feat_1_reg.
type SyscpIDInstructionsFeat1Reg = Syscp_ID_instructions_feat_1_reg

// TaskAbsolutetimeInfoData is a Go-name alias for Task_absolutetime_info_data_t.
type TaskAbsolutetimeInfoData = Task_absolutetime_info_data_t

// TaskAbsolutetimeInfo is a Go-name alias for Task_absolutetime_info_t.
type TaskAbsolutetimeInfo = Task_absolutetime_info_t

// TaskAffinityTagInfoData is a Go-name alias for Task_affinity_tag_info_data_t.
type TaskAffinityTagInfoData = Task_affinity_tag_info_data_t

// TaskAffinityTagInfo is a Go-name alias for Task_affinity_tag_info_t.
type TaskAffinityTagInfo = Task_affinity_tag_info_t

// TaskArray is a Go-name alias for Task_array_t.
type TaskArray = Task_array_t

// TaskBasicInfo32Data is a Go-name alias for Task_basic_info_32_data_t.
type TaskBasicInfo32Data = Task_basic_info_32_data_t

// TaskBasicInfo32 is a Go-name alias for Task_basic_info_32_t.
type TaskBasicInfo32 = Task_basic_info_32_t

// TaskBasicInfo642Data is a Go-name alias for Task_basic_info_64_2_data_t.
type TaskBasicInfo642Data = Task_basic_info_64_2_data_t

// TaskBasicInfo642 is a Go-name alias for Task_basic_info_64_2_t.
type TaskBasicInfo642 = Task_basic_info_64_2_t

// TaskBasicInfo64Data is a Go-name alias for Task_basic_info_64_data_t.
type TaskBasicInfo64Data = Task_basic_info_64_data_t

// TaskBasicInfo64 is a Go-name alias for Task_basic_info_64_t.
type TaskBasicInfo64 = Task_basic_info_64_t

// TaskBasicInfoData is a Go-name alias for Task_basic_info_data_t.
type TaskBasicInfoData = Task_basic_info_data_t

// TaskBasicInfo is a Go-name alias for Task_basic_info_t.
type TaskBasicInfo = Task_basic_info_t

// TaskCategoryPolicyData is a Go-name alias for Task_category_policy_data_t.
type TaskCategoryPolicyData = Task_category_policy_data_t

// TaskCategoryPolicy is a Go-name alias for Task_category_policy_t.
type TaskCategoryPolicy = Task_category_policy_t

// TaskCorpseForkingBehavior is a Go-name alias for Task_corpse_forking_behavior_t.
type TaskCorpseForkingBehavior = Task_corpse_forking_behavior_t

// TaskCrashinfoItem is a Go-name alias for Task_crashinfo_item_t.
type TaskCrashinfoItem = Task_crashinfo_item_t

// TaskDyldInfoData is a Go-name alias for Task_dyld_info_data_t.
type TaskDyldInfoData = Task_dyld_info_data_t

// TaskDyldInfo is a Go-name alias for Task_dyld_info_t.
type TaskDyldInfo = Task_dyld_info_t

// TaskEventsInfoData is a Go-name alias for Task_events_info_data_t.
type TaskEventsInfoData = Task_events_info_data_t

// TaskEventsInfo is a Go-name alias for Task_events_info_t.
type TaskEventsInfo = Task_events_info_t

// TaskExcGuardBehavior is a Go-name alias for Task_exc_guard_behavior_t.
type TaskExcGuardBehavior = Task_exc_guard_behavior_t

// TaskExtmodInfoData is a Go-name alias for Task_extmod_info_data_t.
type TaskExtmodInfoData = Task_extmod_info_data_t

// TaskExtmodInfo is a Go-name alias for Task_extmod_info_t.
type TaskExtmodInfo = Task_extmod_info_t

// TaskFlagsInfoData is a Go-name alias for Task_flags_info_data_t.
type TaskFlagsInfoData = Task_flags_info_data_t

// TaskFlagsInfo is a Go-name alias for Task_flags_info_t.
type TaskFlagsInfo = Task_flags_info_t

// TaskFlavor is a Go-name alias for Task_flavor_t.
type TaskFlavor = Task_flavor_t

// TaskGate is a Go-name alias for Task_gate_t.
type TaskGate = Task_gate_t

// TaskIDToken is a Go-name alias for Task_id_token_t.
type TaskIDToken = Task_id_token_t

// TaskInfoData is a Go-name alias for Task_info_data_t.
type TaskInfoData = Task_info_data_t

// TaskInfo is a Go-name alias for Task_info_t.
type TaskInfo = Task_info_t

// TaskInspectBasicCountsData is a Go-name alias for Task_inspect_basic_counts_data_t.
type TaskInspectBasicCountsData = Task_inspect_basic_counts_data_t

// TaskInspectBasicCounts is a Go-name alias for Task_inspect_basic_counts_t.
type TaskInspectBasicCounts = Task_inspect_basic_counts_t

// TaskInspectFlavor is a Go-name alias for Task_inspect_flavor_t.
type TaskInspectFlavor = Task_inspect_flavor_t

// TaskInspectInfo is a Go-name alias for Task_inspect_info_t.
type TaskInspectInfo = Task_inspect_info_t

// TaskInspect is a Go-name alias for Task_inspect_t.
type TaskInspect = Task_inspect_t

// TaskKernelmemoryInfoData is a Go-name alias for Task_kernelmemory_info_data_t.
type TaskKernelmemoryInfoData = Task_kernelmemory_info_data_t

// TaskKernelmemoryInfo is a Go-name alias for Task_kernelmemory_info_t.
type TaskKernelmemoryInfo = Task_kernelmemory_info_t

// TaskLatencyQos is a Go-name alias for Task_latency_qos_t.
type TaskLatencyQos = Task_latency_qos_t

// TaskName is a Go-name alias for Task_name_t.
type TaskName = Task_name_t

// TaskPolicyFlavor is a Go-name alias for Task_policy_flavor_t.
type TaskPolicyFlavor = Task_policy_flavor_t

// TaskPolicyGet is a Go-name alias for Task_policy_get_t.
type TaskPolicyGet = Task_policy_get_t

// TaskPolicySet is a Go-name alias for Task_policy_set_t.
type TaskPolicySet = Task_policy_set_t

// TaskPolicy is a Go-name alias for Task_policy_t.
type TaskPolicy = Task_policy_t

// TaskPortArray is a Go-name alias for Task_port_array_t.
type TaskPortArray = Task_port_array_t

// TaskPort is a Go-name alias for Task_port_t.
type TaskPort = Task_port_t

// TaskPowerInfoData is a Go-name alias for Task_power_info_data_t.
type TaskPowerInfoData = Task_power_info_data_t

// TaskPowerInfo is a Go-name alias for Task_power_info_t.
type TaskPowerInfo = Task_power_info_t

// TaskPowerInfoV2Data is a Go-name alias for Task_power_info_v2_data_t.
type TaskPowerInfoV2Data = Task_power_info_v2_data_t

// TaskPowerInfoV2 is a Go-name alias for Task_power_info_v2_t.
type TaskPowerInfoV2 = Task_power_info_v2_t

// TaskPurgableInfo is a Go-name alias for Task_purgable_info_t.
type TaskPurgableInfo = Task_purgable_info_t

// TaskQosPolicy is a Go-name alias for Task_qos_policy_t.
type TaskQosPolicy = Task_qos_policy_t

// TaskRead is a Go-name alias for Task_read_t.
type TaskRead = Task_read_t

// TaskRestartableRangeArray is a Go-name alias for Task_restartable_range_array_t.
type TaskRestartableRangeArray = Task_restartable_range_array_t

// TaskRestartableRange is a Go-name alias for Task_restartable_range_t.
type TaskRestartableRange = Task_restartable_range_t

// TaskRole is a Go-name alias for Task_role_t.
type TaskRole = Task_role_t

// TaskSpecialPort is a Go-name alias for Task_special_port_t.
type TaskSpecialPort = Task_special_port_t

// TaskSuspensionToken is a Go-name alias for Task_suspension_token_t.
type TaskSuspensionToken = Task_suspension_token_t

// TaskThreadTimesInfoData is a Go-name alias for Task_thread_times_info_data_t.
type TaskThreadTimesInfoData = Task_thread_times_info_data_t

// TaskThreadTimesInfo is a Go-name alias for Task_thread_times_info_t.
type TaskThreadTimesInfo = Task_thread_times_info_t

// TaskThroughputQos is a Go-name alias for Task_throughput_qos_t.
type TaskThroughputQos = Task_throughput_qos_t

// TaskTraceMemoryInfoData is a Go-name alias for Task_trace_memory_info_data_t.
type TaskTraceMemoryInfoData = Task_trace_memory_info_data_t

// TaskTraceMemoryInfo is a Go-name alias for Task_trace_memory_info_t.
type TaskTraceMemoryInfo = Task_trace_memory_info_t

// TaskVmInfoData is a Go-name alias for Task_vm_info_data_t.
type TaskVmInfoData = Task_vm_info_data_t

// TaskVmInfo is a Go-name alias for Task_vm_info_t.
type TaskVmInfo = Task_vm_info_t

// TaskWaitStateInfoData is a Go-name alias for Task_wait_state_info_data_t.
type TaskWaitStateInfoData = Task_wait_state_info_data_t

// TaskWaitStateInfo is a Go-name alias for Task_wait_state_info_t.
type TaskWaitStateInfo = Task_wait_state_info_t

// TaskZoneInfoArray is a Go-name alias for Task_zone_info_array_t.
type TaskZoneInfoArray = Task_zone_info_array_t

// TaskZoneInfo is a Go-name alias for Task_zone_info_t.
type TaskZoneInfo = Task_zone_info_t

// Tcflag is a Go-name alias for Tcflag_t.
type Tcflag = Tcflag_t

// TCPCc is a Go-name alias for Tcp_cc.
type TCPCc = Tcp_cc

// TCPConnectionClientAccurateEcnState is a Go-name alias for Tcp_connection_client_accurate_ecn_state_t.
type TCPConnectionClientAccurateEcnState = Tcp_connection_client_accurate_ecn_state_t

// TCPConnectionServerAccurateEcnState is a Go-name alias for Tcp_connection_server_accurate_ecn_state_t.
type TCPConnectionServerAccurateEcnState = Tcp_connection_server_accurate_ecn_state_t

// TCPNotifyAckID is a Go-name alias for Tcp_notify_ack_id_t.
type TCPNotifyAckID = Tcp_notify_ack_id_t

// TCPSeq is a Go-name alias for Tcp_seq.
type TCPSeq = Tcp_seq

// TextEncoding is a Go-name alias for Text_encoding_t.
type TextEncoding = Text_encoding_t

// ThreadActArray is a Go-name alias for Thread_act_array_t.
type ThreadActArray = Thread_act_array_t

// ThreadActPortArray is a Go-name alias for Thread_act_port_array_t.
type ThreadActPortArray = Thread_act_port_array_t

// ThreadActPort is a Go-name alias for Thread_act_port_t.
type ThreadActPort = Thread_act_port_t

// ThreadAct is a Go-name alias for Thread_act_t.
type ThreadAct = Thread_act_t

// ThreadAffinityPolicyData is a Go-name alias for Thread_affinity_policy_data_t.
type ThreadAffinityPolicyData = Thread_affinity_policy_data_t

// ThreadAffinityPolicy is a Go-name alias for Thread_affinity_policy_t.
type ThreadAffinityPolicy = Thread_affinity_policy_t

// ThreadArray is a Go-name alias for Thread_array_t.
type ThreadArray = Thread_array_t

// ThreadBackgroundPolicyData is a Go-name alias for Thread_background_policy_data_t.
type ThreadBackgroundPolicyData = Thread_background_policy_data_t

// ThreadBackgroundPolicy is a Go-name alias for Thread_background_policy_t.
type ThreadBackgroundPolicy = Thread_background_policy_t

// ThreadBasicInfoData is a Go-name alias for Thread_basic_info_data_t.
type ThreadBasicInfoData = Thread_basic_info_data_t

// ThreadBasicInfo is a Go-name alias for Thread_basic_info_t.
type ThreadBasicInfo = Thread_basic_info_t

// ThreadCallParam is a Go-name alias for Thread_call_param_t.
type ThreadCallParam = Thread_call_param_t

// ThreadCall is a Go-name alias for Thread_call_t.
type ThreadCall = Thread_call_t

// ThreadExtendedInfoData is a Go-name alias for Thread_extended_info_data_t.
type ThreadExtendedInfoData = Thread_extended_info_data_t

// ThreadExtendedInfo is a Go-name alias for Thread_extended_info_t.
type ThreadExtendedInfo = Thread_extended_info_t

// ThreadExtendedPolicyData is a Go-name alias for Thread_extended_policy_data_t.
type ThreadExtendedPolicyData = Thread_extended_policy_data_t

// ThreadExtendedPolicy is a Go-name alias for Thread_extended_policy_t.
type ThreadExtendedPolicy = Thread_extended_policy_t

// ThreadFlavor is a Go-name alias for Thread_flavor_t.
type ThreadFlavor = Thread_flavor_t

// ThreadIdentifierInfoData is a Go-name alias for Thread_identifier_info_data_t.
type ThreadIdentifierInfoData = Thread_identifier_info_data_t

// ThreadIdentifierInfo is a Go-name alias for Thread_identifier_info_t.
type ThreadIdentifierInfo = Thread_identifier_info_t

// ThreadInfoData is a Go-name alias for Thread_info_data_t.
type ThreadInfoData = Thread_info_data_t

// ThreadInfo is a Go-name alias for Thread_info_t.
type ThreadInfo = Thread_info_t

// ThreadInspect is a Go-name alias for Thread_inspect_t.
type ThreadInspect = Thread_inspect_t

// ThreadLatencyQosPolicyData is a Go-name alias for Thread_latency_qos_policy_data_t.
type ThreadLatencyQosPolicyData = Thread_latency_qos_policy_data_t

// ThreadLatencyQosPolicy is a Go-name alias for Thread_latency_qos_policy_t.
type ThreadLatencyQosPolicy = Thread_latency_qos_policy_t

// ThreadLatencyQos is a Go-name alias for Thread_latency_qos_t.
type ThreadLatencyQos = Thread_latency_qos_t

// ThreadPolicyFlavor is a Go-name alias for Thread_policy_flavor_t.
type ThreadPolicyFlavor = Thread_policy_flavor_t

// ThreadPolicy is a Go-name alias for Thread_policy_t.
type ThreadPolicy = Thread_policy_t

// ThreadPortArray is a Go-name alias for Thread_port_array_t.
type ThreadPortArray = Thread_port_array_t

// ThreadPort is a Go-name alias for Thread_port_t.
type ThreadPort = Thread_port_t

// ThreadPrecedencePolicyData is a Go-name alias for Thread_precedence_policy_data_t.
type ThreadPrecedencePolicyData = Thread_precedence_policy_data_t

// ThreadPrecedencePolicy is a Go-name alias for Thread_precedence_policy_t.
type ThreadPrecedencePolicy = Thread_precedence_policy_t

// ThreadRead is a Go-name alias for Thread_read_t.
type ThreadRead = Thread_read_t

// ThreadSelfcountsKind is a Go-name alias for Thread_selfcounts_kind_t.
type ThreadSelfcountsKind = Thread_selfcounts_kind_t

// ThreadStandardPolicyData is a Go-name alias for Thread_standard_policy_data_t.
type ThreadStandardPolicyData = Thread_standard_policy_data_t

// ThreadStandardPolicy is a Go-name alias for Thread_standard_policy_t.
type ThreadStandardPolicy = Thread_standard_policy_t

// ThreadStateData is a Go-name alias for Thread_state_data_t.
type ThreadStateData = Thread_state_data_t

// ThreadStateFlavorArray is a Go-name alias for Thread_state_flavor_array_t.
type ThreadStateFlavorArray = Thread_state_flavor_array_t

// ThreadStateFlavor is a Go-name alias for Thread_state_flavor_t.
type ThreadStateFlavor = Thread_state_flavor_t

// ThreadState is a Go-name alias for Thread_state_t.
type ThreadState = Thread_state_t

// Thread is a Go-name alias for Thread_t.
type Thread = Thread_t

// ThreadThroughputQosPolicyData is a Go-name alias for Thread_throughput_qos_policy_data_t.
type ThreadThroughputQosPolicyData = Thread_throughput_qos_policy_data_t

// ThreadThroughputQosPolicy is a Go-name alias for Thread_throughput_qos_policy_t.
type ThreadThroughputQosPolicy = Thread_throughput_qos_policy_t

// ThreadThroughputQos is a Go-name alias for Thread_throughput_qos_t.
type ThreadThroughputQos = Thread_throughput_qos_t

// ThreadTimeConstraintPolicyData is a Go-name alias for Thread_time_constraint_policy_data_t.
type ThreadTimeConstraintPolicyData = Thread_time_constraint_policy_data_t

// ThreadTimeConstraintPolicy is a Go-name alias for Thread_time_constraint_policy_t.
type ThreadTimeConstraintPolicy = Thread_time_constraint_policy_t

// ThreadTurnstileinfo is a Go-name alias for Thread_turnstileinfo_t.
type ThreadTurnstileinfo = Thread_turnstileinfo_t

// ThreadTurnstileinfoV2 is a Go-name alias for Thread_turnstileinfo_v2_t.
type ThreadTurnstileinfoV2 = Thread_turnstileinfo_v2_t

// ThreadWaitinfo is a Go-name alias for Thread_waitinfo_t.
type ThreadWaitinfo = Thread_waitinfo_t

// ThreadWaitinfoV2 is a Go-name alias for Thread_waitinfo_v2_t.
type ThreadWaitinfoV2 = Thread_waitinfo_v2_t

// ThrottleInfoHandle is a Go-name alias for Throttle_info_handle_t.
type ThrottleInfoHandle = Throttle_info_handle_t

// Time is a Go-name alias for Time_t.
type Time = Time_t

// TimeValue is a Go-name alias for Time_value_t.
type TimeValue = Time_value_t

// Token is a Go-name alias for Token_t.
type Token = Token_t

// TrapGate is a Go-name alias for Trap_gate_t.
type TrapGate = Trap_gate_t

// TssDesc is a Go-name alias for Tss_desc_t.
type TssDesc = Tss_desc_t

// Tss is a Go-name alias for Tss_t.
type Tss = Tss_t

// UChar is a Go-name alias for U_char.
type UChar = U_char

// UQuad is a Go-name alias for U_quad_t.
type UQuad = U_quad_t

// UShort is a Go-name alias for U_short.
type UShort = U_short

// Ucontext64 is a Go-name alias for Ucontext64_t.
type Ucontext64 = Ucontext64_t

// Ucontext is a Go-name alias for Ucontext_t.
type Ucontext = Ucontext_t

// UextObject is a Go-name alias for Uext_object_t.
type UextObject = Uext_object_t

// Uid is a Go-name alias for Uid_t.
type Uid = Uid_t

// Uint16 is a Go-name alias for Uint16_t.
type Uint16 = Uint16_t

// Uint32 is a Go-name alias for Uint32_t.
type Uint32 = Uint32_t

// Uint64 is a Go-name alias for Uint64_t.
type Uint64 = Uint64_t

// Uint8 is a Go-name alias for Uint8_t.
type Uint8 = Uint8_t

// UintFast16 is a Go-name alias for Uint_fast16_t.
type UintFast16 = Uint_fast16_t

// UintFast32 is a Go-name alias for Uint_fast32_t.
type UintFast32 = Uint_fast32_t

// UintFast64 is a Go-name alias for Uint_fast64_t.
type UintFast64 = Uint_fast64_t

// UintFast8 is a Go-name alias for Uint_fast8_t.
type UintFast8 = Uint_fast8_t

// UintLeast16 is a Go-name alias for Uint_least16_t.
type UintLeast16 = Uint_least16_t

// UintLeast32 is a Go-name alias for Uint_least32_t.
type UintLeast32 = Uint_least32_t

// UintLeast64 is a Go-name alias for Uint_least64_t.
type UintLeast64 = Uint_least64_t

// UintLeast8 is a Go-name alias for Uint_least8_t.
type UintLeast8 = Uint_least8_t

// Uintmax is a Go-name alias for Uintmax_t.
type Uintmax = Uintmax_t

// Uintptr is a Go-name alias for Uintptr_t.
type Uintptr = Uintptr_t

// UioBptr is a Go-name alias for Uio_bptr_t.
type UioBptr = Uio_bptr_t

// UioPtrRef is a Go-name alias for Uio_ptr_ref_t.
type UioPtrRef = Uio_ptr_ref_t

// UioPtr is a Go-name alias for Uio_ptr_t.
type UioPtr = Uio_ptr_t

// UioRefPtr is a Go-name alias for Uio_ref_ptr_t.
type UioRefPtr = Uio_ref_ptr_t

// UioRefRef is a Go-name alias for Uio_ref_ref_t.
type UioRefRef = Uio_ref_ref_t

// UioRef is a Go-name alias for Uio_ref_t.
type UioRef = Uio_ref_t

// Uio is a Go-name alias for Uio_t.
type Uio = Uio_t

// UnpGen is a Go-name alias for Unp_gen_t.
type UnpGen = Unp_gen_t

// UplControlFlags is a Go-name alias for Upl_control_flags_t.
type UplControlFlags = Upl_control_flags_t

// UplOffset is a Go-name alias for Upl_offset_t.
type UplOffset = Upl_offset_t

// UplPageInfoArray is a Go-name alias for Upl_page_info_array_t.
type UplPageInfoArray = Upl_page_info_array_t

// UplPageInfo is a Go-name alias for Upl_page_info_t.
type UplPageInfo = Upl_page_info_t

// UplPageListPtr is a Go-name alias for Upl_page_list_ptr_t.
type UplPageListPtr = Upl_page_list_ptr_t

// UplSize is a Go-name alias for Upl_size_t.
type UplSize = Upl_size_t

// Upl is a Go-name alias for Upl_t.
type Upl = Upl_t

// Useconds is a Go-name alias for Useconds_t.
type Useconds = Useconds_t

// User32Addr is a Go-name alias for User32_addr_t.
type User32Addr = User32_addr_t

// User32Fchecklv is a Go-name alias for User32_fchecklv_t.
type User32Fchecklv = User32_fchecklv_t

// User32Fsignatures is a Go-name alias for User32_fsignatures_t.
type User32Fsignatures = User32_fsignatures_t

// User32Long is a Go-name alias for User32_long_t.
type User32Long = User32_long_t

// User32Msglen is a Go-name alias for User32_msglen_t.
type User32Msglen = User32_msglen_t

// User32Msgqnum is a Go-name alias for User32_msgqnum_t.
type User32Msgqnum = User32_msgqnum_t

// User32Off is a Go-name alias for User32_off_t.
type User32Off = User32_off_t

// User32Size is a Go-name alias for User32_size_t.
type User32Size = User32_size_t

// User32Ssize is a Go-name alias for User32_ssize_t.
type User32Ssize = User32_ssize_t

// User32Time is a Go-name alias for User32_time_t.
type User32Time = User32_time_t

// User32Ulong is a Go-name alias for User32_ulong_t.
type User32Ulong = User32_ulong_t

// User64Addr is a Go-name alias for User64_addr_t.
type User64Addr = User64_addr_t

// User64Long is a Go-name alias for User64_long_t.
type User64Long = User64_long_t

// User64Msglen is a Go-name alias for User64_msglen_t.
type User64Msglen = User64_msglen_t

// User64Msgqnum is a Go-name alias for User64_msgqnum_t.
type User64Msgqnum = User64_msgqnum_t

// User64Off is a Go-name alias for User64_off_t.
type User64Off = User64_off_t

// User64Size is a Go-name alias for User64_size_t.
type User64Size = User64_size_t

// User64Ssize is a Go-name alias for User64_ssize_t.
type User64Ssize = User64_ssize_t

// User64Time is a Go-name alias for User64_time_t.
type User64Time = User64_time_t

// User64Ulong is a Go-name alias for User64_ulong_t.
type User64Ulong = User64_ulong_t

// UserAddr is a Go-name alias for User_addr_t.
type UserAddr = User_addr_t

// UserAddrUt is a Go-name alias for User_addr_ut.
type UserAddrUt = User_addr_ut

// UserFchecklv is a Go-name alias for User_fchecklv_t.
type UserFchecklv = User_fchecklv_t

// UserFsignatures is a Go-name alias for User_fsignatures_t.
type UserFsignatures = User_fsignatures_t

// UserFsupplement is a Go-name alias for User_fsupplement_t.
type UserFsupplement = User_fsupplement_t

// UserLong is a Go-name alias for User_long_t.
type UserLong = User_long_t

// UserMsglen is a Go-name alias for User_msglen_t.
type UserMsglen = User_msglen_t

// UserMsgqnum is a Go-name alias for User_msgqnum_t.
type UserMsgqnum = User_msgqnum_t

// UserOff is a Go-name alias for User_off_t.
type UserOff = User_off_t

// UserSize is a Go-name alias for User_size_t.
type UserSize = User_size_t

// UserSizeUt is a Go-name alias for User_size_ut.
type UserSizeUt = User_size_ut

// UserSpeed is a Go-name alias for User_speed_t.
type UserSpeed = User_speed_t

// UserSsize is a Go-name alias for User_ssize_t.
type UserSsize = User_ssize_t

// UserSubsystem is a Go-name alias for User_subsystem_t.
type UserSubsystem = User_subsystem_t

// UserTcflag is a Go-name alias for User_tcflag_t.
type UserTcflag = User_tcflag_t

// UserTime is a Go-name alias for User_time_t.
type UserTime = User_time_t

// UserUlong is a Go-name alias for User_ulong_t.
type UserUlong = User_ulong_t

// UuidString is a Go-name alias for Uuid_string_t.
type UuidString = Uuid_string_t

// Uuid is a Go-name alias for Uuid_t.
type Uuid = Uuid_t

// VdspLength is a Go-name alias for VDSP_Length.
type VdspLength = VDSP_Length

// VdspStride is a Go-name alias for VDSP_Stride.
type VdspStride = VDSP_Stride

// VdspBiquadSetup is a Go-name alias for VDSP_biquad_Setup.
type VdspBiquadSetup = VDSP_biquad_Setup

// VdspBiquadSetupD is a Go-name alias for VDSP_biquad_SetupD.
type VdspBiquadSetupD = VDSP_biquad_SetupD

// VdspBiquadmSetup is a Go-name alias for VDSP_biquadm_Setup.
type VdspBiquadmSetup = VDSP_biquadm_Setup

// VdspBiquadmSetupD is a Go-name alias for VDSP_biquadm_SetupD.
type VdspBiquadmSetupD = VDSP_biquadm_SetupD

// VdspInt24 is a Go-name alias for VDSP_int24.
type VdspInt24 = VDSP_int24

// VdspUint24 is a Go-name alias for VDSP_uint24.
type VdspUint24 = VDSP_uint24

// VaList is a Go-name alias for Va_list.
type VaList = Va_list

// VcProgressUserOptions is a Go-name alias for Vc_progress_user_options.
type VcProgressUserOptions = Vc_progress_user_options

// VectorInt2 is a Go-name alias for Vector_int2.
type VectorInt2 = Vector_int2

// VectorInt4 is a Go-name alias for Vector_int4.
type VectorInt4 = Vector_int4

// VectorInt8 is a Go-name alias for Vector_int8.
type VectorInt8 = Vector_int8

// VectorUchar16 is a Go-name alias for Vector_uchar16.
type VectorUchar16 = Vector_uchar16

// VectorUchar32 is a Go-name alias for Vector_uchar32.
type VectorUchar32 = Vector_uchar32

// VectorUchar64 is a Go-name alias for Vector_uchar64.
type VectorUchar64 = Vector_uchar64

// VectorUchar8 is a Go-name alias for Vector_uchar8.
type VectorUchar8 = Vector_uchar8

// VectorUint4 is a Go-name alias for Vector_uint4.
type VectorUint4 = Vector_uint4

// VectorUshort4 is a Go-name alias for Vector_ushort4.
type VectorUshort4 = Vector_ushort4

// VfsContextBptr is a Go-name alias for Vfs_context_bptr_t.
type VfsContextBptr = Vfs_context_bptr_t

// VfsContextPtrRef is a Go-name alias for Vfs_context_ptr_ref_t.
type VfsContextPtrRef = Vfs_context_ptr_ref_t

// VfsContextPtr is a Go-name alias for Vfs_context_ptr_t.
type VfsContextPtr = Vfs_context_ptr_t

// VfsContextRefPtr is a Go-name alias for Vfs_context_ref_ptr_t.
type VfsContextRefPtr = Vfs_context_ref_ptr_t

// VfsContextRefRef is a Go-name alias for Vfs_context_ref_ref_t.
type VfsContextRefRef = Vfs_context_ref_ref_t

// VfsContextRef is a Go-name alias for Vfs_context_ref_t.
type VfsContextRef = Vfs_context_ref_t

// VfsContext is a Go-name alias for Vfs_context_t.
type VfsContext = Vfs_context_t

// VfsPath is a Go-name alias for Vfs_path_t.
type VfsPath = Vfs_path_t

// VfsRenameFlags is a Go-name alias for Vfs_rename_flags_t.
type VfsRenameFlags = Vfs_rename_flags_t

// VfsRoles is a Go-name alias for Vfs_roles_t.
type VfsRoles = Vfs_roles_t

// VfstableBptr is a Go-name alias for Vfstable_bptr_t.
type VfstableBptr = Vfstable_bptr_t

// VfstablePtrRef is a Go-name alias for Vfstable_ptr_ref_t.
type VfstablePtrRef = Vfstable_ptr_ref_t

// VfstablePtr is a Go-name alias for Vfstable_ptr_t.
type VfstablePtr = Vfstable_ptr_t

// VfstableRefPtr is a Go-name alias for Vfstable_ref_ptr_t.
type VfstableRefPtr = Vfstable_ref_ptr_t

// VfstableRefRef is a Go-name alias for Vfstable_ref_ref_t.
type VfstableRefRef = Vfstable_ref_ref_t

// VfstableRef is a Go-name alias for Vfstable_ref_t.
type VfstableRef = Vfstable_ref_t

// Vfstable is a Go-name alias for Vfstable_t.
type Vfstable = Vfstable_t

// Vm32AddrStruct is a Go-name alias for Vm32_addr_struct_t.
type Vm32AddrStruct = Vm32_addr_struct_t

// Vm32Address is a Go-name alias for Vm32_address_t.
type Vm32Address = Vm32_address_t

// Vm32ObjectID is a Go-name alias for Vm32_object_id_t.
type Vm32ObjectID = Vm32_object_id_t

// Vm32Offset is a Go-name alias for Vm32_offset_t.
type Vm32Offset = Vm32_offset_t

// Vm32SizeStruct is a Go-name alias for Vm32_size_struct_t.
type Vm32SizeStruct = Vm32_size_struct_t

// Vm32Size is a Go-name alias for Vm32_size_t.
type Vm32Size = Vm32_size_t

// VmAddrStruct is a Go-name alias for Vm_addr_struct_t.
type VmAddrStruct = Vm_addr_struct_t

// VmAddress is a Go-name alias for Vm_address_t.
type VmAddress = Vm_address_t

// VmAddressUt is a Go-name alias for Vm_address_ut.
type VmAddressUt = Vm_address_ut

// VmBehavior is a Go-name alias for Vm_behavior_t.
type VmBehavior = Vm_behavior_t

// VmBehaviorUt is a Go-name alias for Vm_behavior_ut.
type VmBehaviorUt = Vm_behavior_ut

// VmExtmodStatisticsData is a Go-name alias for Vm_extmod_statistics_data_t.
type VmExtmodStatisticsData = Vm_extmod_statistics_data_t

// VmExtmodStatistics is a Go-name alias for Vm_extmod_statistics_t.
type VmExtmodStatistics = Vm_extmod_statistics_t

// VmInfoObjectArray is a Go-name alias for Vm_info_object_array_t.
type VmInfoObjectArray = Vm_info_object_array_t

// VmInfoObject is a Go-name alias for Vm_info_object_t.
type VmInfoObject = Vm_info_object_t

// VmInfoRegion64 is a Go-name alias for Vm_info_region_64_t.
type VmInfoRegion64 = Vm_info_region_64_t

// VmInfoRegion is a Go-name alias for Vm_info_region_t.
type VmInfoRegion = Vm_info_region_t

// VmInherit is a Go-name alias for Vm_inherit_t.
type VmInherit = Vm_inherit_t

// VmInheritUt is a Go-name alias for Vm_inherit_ut.
type VmInheritUt = Vm_inherit_ut

// VmMachineAttribute is a Go-name alias for Vm_machine_attribute_t.
type VmMachineAttribute = Vm_machine_attribute_t

// VmMachineAttributeVal is a Go-name alias for Vm_machine_attribute_val_t.
type VmMachineAttributeVal = Vm_machine_attribute_val_t

// VmMapAddress is a Go-name alias for Vm_map_address_t.
type VmMapAddress = Vm_map_address_t

// VmMapAddressUt is a Go-name alias for Vm_map_address_ut.
type VmMapAddressUt = Vm_map_address_ut

// VmMapInspect is a Go-name alias for Vm_map_inspect_t.
type VmMapInspect = Vm_map_inspect_t

// VmMapOffset is a Go-name alias for Vm_map_offset_t.
type VmMapOffset = Vm_map_offset_t

// VmMapOffsetUt is a Go-name alias for Vm_map_offset_ut.
type VmMapOffsetUt = Vm_map_offset_ut

// VmMapRead is a Go-name alias for Vm_map_read_t.
type VmMapRead = Vm_map_read_t

// VmMapSize is a Go-name alias for Vm_map_size_t.
type VmMapSize = Vm_map_size_t

// VmMapSizeUt is a Go-name alias for Vm_map_size_ut.
type VmMapSizeUt = Vm_map_size_ut

// VmMap is a Go-name alias for Vm_map_t.
type VmMap = Vm_map_t

// VmNamedEntry is a Go-name alias for Vm_named_entry_t.
type VmNamedEntry = Vm_named_entry_t

// VmObjectID is a Go-name alias for Vm_object_id_t.
type VmObjectID = Vm_object_id_t

// VmObjectOffset is a Go-name alias for Vm_object_offset_t.
type VmObjectOffset = Vm_object_offset_t

// VmObjectOffsetUt is a Go-name alias for Vm_object_offset_ut.
type VmObjectOffsetUt = Vm_object_offset_ut

// VmObjectSize is a Go-name alias for Vm_object_size_t.
type VmObjectSize = Vm_object_size_t

// VmObjectSizeUt is a Go-name alias for Vm_object_size_ut.
type VmObjectSizeUt = Vm_object_size_ut

// VmOffset is a Go-name alias for Vm_offset_t.
type VmOffset = Vm_offset_t

// VmOffsetUt is a Go-name alias for Vm_offset_ut.
type VmOffsetUt = Vm_offset_ut

// VmPageInfoBasicData is a Go-name alias for Vm_page_info_basic_data_t.
type VmPageInfoBasicData = Vm_page_info_basic_data_t

// VmPageInfoBasic is a Go-name alias for Vm_page_info_basic_t.
type VmPageInfoBasic = Vm_page_info_basic_t

// VmPageInfoData is a Go-name alias for Vm_page_info_data_t.
type VmPageInfoData = Vm_page_info_data_t

// VmPageInfoFlavor is a Go-name alias for Vm_page_info_flavor_t.
type VmPageInfoFlavor = Vm_page_info_flavor_t

// VmPageInfo is a Go-name alias for Vm_page_info_t.
type VmPageInfo = Vm_page_info_t

// VmProt is a Go-name alias for Vm_prot_t.
type VmProt = Vm_prot_t

// VmProtUt is a Go-name alias for Vm_prot_ut.
type VmProtUt = Vm_prot_ut

// VmPurgable is a Go-name alias for Vm_purgable_t.
type VmPurgable = Vm_purgable_t

// VmPurgeableInfo is a Go-name alias for Vm_purgeable_info_t.
type VmPurgeableInfo = Vm_purgeable_info_t

// VmPurgeableStat is a Go-name alias for Vm_purgeable_stat_t.
type VmPurgeableStat = Vm_purgeable_stat_t

// VmReadEntry is a Go-name alias for Vm_read_entry_t.
type VmReadEntry = Vm_read_entry_t

// VmRegionBasicInfo64 is a Go-name alias for Vm_region_basic_info_64_t.
type VmRegionBasicInfo64 = Vm_region_basic_info_64_t

// VmRegionBasicInfoData64 is a Go-name alias for Vm_region_basic_info_data_64_t.
type VmRegionBasicInfoData64 = Vm_region_basic_info_data_64_t

// VmRegionBasicInfoData is a Go-name alias for Vm_region_basic_info_data_t.
type VmRegionBasicInfoData = Vm_region_basic_info_data_t

// VmRegionBasicInfo is a Go-name alias for Vm_region_basic_info_t.
type VmRegionBasicInfo = Vm_region_basic_info_t

// VmRegionExtendedInfoData is a Go-name alias for Vm_region_extended_info_data_t.
type VmRegionExtendedInfoData = Vm_region_extended_info_data_t

// VmRegionExtendedInfo is a Go-name alias for Vm_region_extended_info_t.
type VmRegionExtendedInfo = Vm_region_extended_info_t

// VmRegionFlavor is a Go-name alias for Vm_region_flavor_t.
type VmRegionFlavor = Vm_region_flavor_t

// VmRegionInfo64 is a Go-name alias for Vm_region_info_64_t.
type VmRegionInfo64 = Vm_region_info_64_t

// VmRegionInfoData is a Go-name alias for Vm_region_info_data_t.
type VmRegionInfoData = Vm_region_info_data_t

// VmRegionInfo is a Go-name alias for Vm_region_info_t.
type VmRegionInfo = Vm_region_info_t

// VmRegionRecurseInfo64 is a Go-name alias for Vm_region_recurse_info_64_t.
type VmRegionRecurseInfo64 = Vm_region_recurse_info_64_t

// VmRegionRecurseInfo is a Go-name alias for Vm_region_recurse_info_t.
type VmRegionRecurseInfo = Vm_region_recurse_info_t

// VmRegionSubmapInfo64 is a Go-name alias for Vm_region_submap_info_64_t.
type VmRegionSubmapInfo64 = Vm_region_submap_info_64_t

// VmRegionSubmapInfoData64 is a Go-name alias for Vm_region_submap_info_data_64_t.
type VmRegionSubmapInfoData64 = Vm_region_submap_info_data_64_t

// VmRegionSubmapInfoData is a Go-name alias for Vm_region_submap_info_data_t.
type VmRegionSubmapInfoData = Vm_region_submap_info_data_t

// VmRegionSubmapInfo is a Go-name alias for Vm_region_submap_info_t.
type VmRegionSubmapInfo = Vm_region_submap_info_t

// VmRegionSubmapShortInfo64 is a Go-name alias for Vm_region_submap_short_info_64_t.
type VmRegionSubmapShortInfo64 = Vm_region_submap_short_info_64_t

// VmRegionSubmapShortInfoData64 is a Go-name alias for Vm_region_submap_short_info_data_64_t.
type VmRegionSubmapShortInfoData64 = Vm_region_submap_short_info_data_64_t

// VmRegionTopInfoData is a Go-name alias for Vm_region_top_info_data_t.
type VmRegionTopInfoData = Vm_region_top_info_data_t

// VmRegionTopInfo is a Go-name alias for Vm_region_top_info_t.
type VmRegionTopInfo = Vm_region_top_info_t

// VmSizeStruct is a Go-name alias for Vm_size_struct_t.
type VmSizeStruct = Vm_size_struct_t

// VmSize is a Go-name alias for Vm_size_t.
type VmSize = Vm_size_t

// VmSizeUt is a Go-name alias for Vm_size_ut.
type VmSizeUt = Vm_size_ut

// VmStatistics64Data is a Go-name alias for Vm_statistics64_data_t.
type VmStatistics64Data = Vm_statistics64_data_t

// VmStatistics64 is a Go-name alias for Vm_statistics64_t.
type VmStatistics64 = Vm_statistics64_t

// VmStatisticsData is a Go-name alias for Vm_statistics_data_t.
type VmStatisticsData = Vm_statistics_data_t

// VmStatistics is a Go-name alias for Vm_statistics_t.
type VmStatistics = Vm_statistics_t

// VmSync is a Go-name alias for Vm_sync_t.
type VmSync = Vm_sync_t

// VmTaskEntry is a Go-name alias for Vm_task_entry_t.
type VmTaskEntry = Vm_task_entry_t

// VnodeBptr is a Go-name alias for Vnode_bptr_t.
type VnodeBptr = Vnode_bptr_t

// VnodePtrRef is a Go-name alias for Vnode_ptr_ref_t.
type VnodePtrRef = Vnode_ptr_ref_t

// VnodePtr is a Go-name alias for Vnode_ptr_t.
type VnodePtr = Vnode_ptr_t

// VnodeRefPtr is a Go-name alias for Vnode_ref_ptr_t.
type VnodeRefPtr = Vnode_ref_ptr_t

// VnodeRefRef is a Go-name alias for Vnode_ref_ref_t.
type VnodeRefRef = Vnode_ref_ref_t

// VnodeRef is a Go-name alias for Vnode_ref_t.
type VnodeRef = Vnode_ref_t

// Vnode is a Go-name alias for Vnode_t.
type Vnode = Vnode_t

// VnodeVerifyFlags is a Go-name alias for Vnode_verify_flags_t.
type VnodeVerifyFlags = Vnode_verify_flags_t

// VolAttributesAttr is a Go-name alias for Vol_attributes_attr_t.
type VolAttributesAttr = Vol_attributes_attr_t

// VolCapabilitiesAttr is a Go-name alias for Vol_capabilities_attr_t.
type VolCapabilitiesAttr = Vol_capabilities_attr_t

// VolCapabilitiesSet is a Go-name alias for Vol_capabilities_set_t.
type VolCapabilitiesSet = Vol_capabilities_set_t

// VsockGen is a Go-name alias for Vsock_gen_t.
type VsockGen = Vsock_gen_t

// WaitInterrupt is a Go-name alias for Wait_interrupt_t.
type WaitInterrupt = Wait_interrupt_t

// WaitResult is a Go-name alias for Wait_result_t.
type WaitResult = Wait_result_t

// WaitTimeoutUrgency is a Go-name alias for Wait_timeout_urgency_t.
type WaitTimeoutUrgency = Wait_timeout_urgency_t

// Wint is a Go-name alias for Wint_t.
type Wint = Wint_t

// X86Avx512State is a Go-name alias for X86_avx512_state_t.
type X86Avx512State = X86_avx512_state_t

// X86AvxState is a Go-name alias for X86_avx_state_t.
type X86AvxState = X86_avx_state_t

// X86DebugState is a Go-name alias for X86_debug_state_t.
type X86DebugState = X86_debug_state_t

// X86ExceptionState32 is a Go-name alias for X86_exception_state32_t.
type X86ExceptionState32 = X86_exception_state32_t

// X86ExceptionState is a Go-name alias for X86_exception_state_t.
type X86ExceptionState = X86_exception_state_t

// X86FloatState32 is a Go-name alias for X86_float_state32_t.
type X86FloatState32 = X86_float_state32_t

// X86FloatState is a Go-name alias for X86_float_state_t.
type X86FloatState = X86_float_state_t

// X86StateHdr is a Go-name alias for X86_state_hdr_t.
type X86StateHdr = X86_state_hdr_t

// X86ThreadState32 is a Go-name alias for X86_thread_state32_t.
type X86ThreadState32 = X86_thread_state32_t

// X86ThreadState is a Go-name alias for X86_thread_state_t.
type X86ThreadState = X86_thread_state_t

// XdrbufType is a Go-name alias for Xdrbuf_type.
type XdrbufType = Xdrbuf_type

// XmlData is a Go-name alias for XmlData_t.
type XmlData = XmlData_t

// ZStream is a Go-name alias for Z_stream.
type ZStream = Z_stream

// ZStreamp is a Go-name alias for Z_streamp.
type ZStreamp = Z_streamp

// ZoneBtrecordArray is a Go-name alias for Zone_btrecord_array_t.
type ZoneBtrecordArray = Zone_btrecord_array_t

// ZoneBtrecord is a Go-name alias for Zone_btrecord_t.
type ZoneBtrecord = Zone_btrecord_t

// ZoneInfoArray is a Go-name alias for Zone_info_array_t.
type ZoneInfoArray = Zone_info_array_t

// ZoneInfo is a Go-name alias for Zone_info_t.
type ZoneInfo = Zone_info_t

// ZoneNameArray is a Go-name alias for Zone_name_array_t.
type ZoneNameArray = Zone_name_array_t

// ZoneName is a Go-name alias for Zone_name_t.
type ZoneName = Zone_name_t

// Uio_rw aliases the generated uio enum family for vnode read/write APIs.
type Uio_rw = Uio

// Uio_seg aliases the generated uio enum family for vnode segment-space APIs.
type Uio_seg = Uio

// Ifnet_interface_advisory_direction aliases the generated advisory direction enum.
type Ifnet_interface_advisory_direction = IfInterfaceAdvisoryDirection

// IfnetInterfaceAdvisoryDirection aliases the generated advisory direction enum.
type IfnetInterfaceAdvisoryDirection = IfInterfaceAdvisoryDirection

// Ifnet_interface_advisory_wifi_freq_band aliases the generated advisory Wi-Fi frequency-band enum.
type Ifnet_interface_advisory_wifi_freq_band = IfInterfaceAdvisoryFreqBand

// Ifnet_interface_advisory_interface_type aliases the generated advisory interface type enum.
type Ifnet_interface_advisory_interface_type = IfInterfaceAdvisoryInterfaceType

// IfnetInterfaceAdvisoryInterfaceType aliases the generated advisory interface type enum.
type IfnetInterfaceAdvisoryInterfaceType = IfInterfaceAdvisoryInterfaceType

// Ifnet_interface_advisory_version is a uint8-backed advisory version enum.
type Ifnet_interface_advisory_version = uint8

// Ifnet_interface_advisory_rate_trend is an int32-backed advisory rate-trend enum.
type Ifnet_interface_advisory_rate_trend = int32
