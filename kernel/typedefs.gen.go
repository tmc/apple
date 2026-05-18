// Code generated from Apple documentation. DO NOT EDIT.

package kernel

import (
	"unsafe"
)

// See: https://developer.apple.com/documentation/kernel/ataoperationtype
type ATAOperationType = uint32

// See: https://developer.apple.com/documentation/kernel/atapiclientdata
type ATAPIClientData = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/atapicmdpacket
type ATAPICmdPacket = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/atarequestidentifier
type ATARequestIdentifier = uintptr

// See: https://developer.apple.com/documentation/kernel/avcconnecttargetplugsinparams
type AVCConnectTargetPlugsInParams = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/avcconnecttargetplugsoutparams
type AVCConnectTargetPlugsOutParams = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/avcgettargetplugconnectioninparams
type AVCGetTargetPlugConnectionInParams = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/avcgettargetplugconnectionoutparams
type AVCGetTargetPlugConnectionOutParams = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/avcsubunitplugrecord
type AVCSubunitPlugRecord = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/avcunitplugrecord
type AVCUnitPlugRecord = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/avcunitplugs
type AVCUnitPlugs = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/avidtype
type AVIDType = uint32

// See: https://developer.apple.com/documentation/kernel/absolutetime
type AbsoluteTime = uint64

// See: https://developer.apple.com/documentation/kernel/bddiscinfo
type BDDiscInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bdfeatures
type BDFeatures = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bdmediatype
type BDMediaType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bdtrackinfo
type BDTrackInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/btheaderrec
type BTHeaderRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/btnodedescriptor
type BTNodeDescriptor = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/block0
type Block0 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothafhhostchannelclassification
type BluetoothAFHHostChannelClassification = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothafhmode
type BluetoothAFHMode = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothafhresults
type BluetoothAFHResults = unsafe.Pointer

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
type BluetoothDeviceAddress = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothdeviceclassmajor
type BluetoothDeviceClassMajor = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothdeviceclassminor
type BluetoothDeviceClassMinor = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothdevicename
type BluetoothDeviceName = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothencryptionenable
type BluetoothEncryptionEnable = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothenhancedsynchronousconnectioninfo
type BluetoothEnhancedSynchronousConnectionInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetootheventfiltercondition
type BluetoothEventFilterCondition = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhciacldatabytecount
type BluetoothHCIACLDataByteCount = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhciafhchannelassessmentmode
type BluetoothHCIAFHChannelAssessmentMode = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciacceptsynchronousconnectionrequestparams
type BluetoothHCIAcceptSynchronousConnectionRequestParams = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhciauthenticationenable
type BluetoothHCIAuthenticationEnable = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciautomaticflushtimeout
type BluetoothHCIAutomaticFlushTimeout = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhciautomaticflushtimeoutinfo
type BluetoothHCIAutomaticFlushTimeoutInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcibuffersize
type BluetoothHCIBufferSize = unsafe.Pointer

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
type BluetoothHCICurrentInquiryAccessCodes = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcicurrentinquiryaccesscodesforwrite
type BluetoothHCICurrentInquiryAccessCodesForWrite = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcidataid
type BluetoothHCIDataID = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhcideletestoredlinkkeyflag
type BluetoothHCIDeleteStoredLinkKeyFlag = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciencryptionkeysize
type BluetoothHCIEncryptionKeySize = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciencryptionkeysizeinfo
type BluetoothHCIEncryptionKeySizeInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhciencryptionmode
type BluetoothHCIEncryptionMode = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcienhancedacceptsynchronousconnectionrequestparams
type BluetoothHCIEnhancedAcceptSynchronousConnectionRequestParams = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcienhancedsetupsynchronousconnectionparams
type BluetoothHCIEnhancedSetupSynchronousConnectionParams = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcierroneousdatareporting
type BluetoothHCIErroneousDataReporting = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventauthenticationcompleteresults
type BluetoothHCIEventAuthenticationCompleteResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventchangeconnectionlinkkeycompleteresults
type BluetoothHCIEventChangeConnectionLinkKeyCompleteResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventcode
type BluetoothHCIEventCode = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventconnectioncompleteresults
type BluetoothHCIEventConnectionCompleteResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventconnectionpackettyperesults
type BluetoothHCIEventConnectionPacketTypeResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventconnectionrequestresults
type BluetoothHCIEventConnectionRequestResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventdatabufferoverflowresults
type BluetoothHCIEventDataBufferOverflowResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventdisconnectioncompleteresults
type BluetoothHCIEventDisconnectionCompleteResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventencryptionchangeresults
type BluetoothHCIEventEncryptionChangeResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventencryptionkeyrefreshcompleteresults
type BluetoothHCIEventEncryptionKeyRefreshCompleteResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventflowspecificationdata
type BluetoothHCIEventFlowSpecificationData = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventflushoccurredresults
type BluetoothHCIEventFlushOccurredResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventhardwareerrorresults
type BluetoothHCIEventHardwareErrorResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventid
type BluetoothHCIEventID = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventleconnectioncompleteresults
type BluetoothHCIEventLEConnectionCompleteResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventleconnectionupdatecompleteresults
type BluetoothHCIEventLEConnectionUpdateCompleteResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventleenhancedconnectioncompleteresults
type BluetoothHCIEventLEEnhancedConnectionCompleteResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventlelongtermkeyrequestresults
type BluetoothHCIEventLELongTermKeyRequestResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventlemetaresults
type BluetoothHCIEventLEMetaResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventlereadremoteusedfeaturescompleteresults
type BluetoothHCIEventLEReadRemoteUsedFeaturesCompleteResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventlinkkeynotificationresults
type BluetoothHCIEventLinkKeyNotificationResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventmask
type BluetoothHCIEventMask = uint64

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventmasterlinkkeycompleteresults
type BluetoothHCIEventMasterLinkKeyCompleteResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventmaxslotschangeresults
type BluetoothHCIEventMaxSlotsChangeResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventmodechangeresults
type BluetoothHCIEventModeChangeResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventpagescanmodechangeresults
type BluetoothHCIEventPageScanModeChangeResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventpagescanrepetitionmodechangeresults
type BluetoothHCIEventPageScanRepetitionModeChangeResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventqossetupcompleteresults
type BluetoothHCIEventQoSSetupCompleteResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventqosviolationresults
type BluetoothHCIEventQoSViolationResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventreadclockoffsetresults
type BluetoothHCIEventReadClockOffsetResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventreadextendedfeaturesresults
type BluetoothHCIEventReadExtendedFeaturesResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventreadremoteextendedfeaturesresults
type BluetoothHCIEventReadRemoteExtendedFeaturesResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventreadremotesupportedfeaturesresults
type BluetoothHCIEventReadRemoteSupportedFeaturesResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventreadremoteversioninforesults
type BluetoothHCIEventReadRemoteVersionInfoResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventreadsupportedfeaturesresults
type BluetoothHCIEventReadSupportedFeaturesResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventremotenamerequestresults
type BluetoothHCIEventRemoteNameRequestResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventreturnlinkkeysresults
type BluetoothHCIEventReturnLinkKeysResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventrolechangeresults
type BluetoothHCIEventRoleChangeResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventsimplepairingcompleteresults
type BluetoothHCIEventSimplePairingCompleteResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventsniffsubratingresults
type BluetoothHCIEventSniffSubratingResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventstatus
type BluetoothHCIEventStatus = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventsynchronousconnectionchangedresults
type BluetoothHCIEventSynchronousConnectionChangedResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventsynchronousconnectioncompleteresults
type BluetoothHCIEventSynchronousConnectionCompleteResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcieventvendorspecificresults
type BluetoothHCIEventVendorSpecificResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhciextendedfeaturesinfo
type BluetoothHCIExtendedFeaturesInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhciextendedinquiryresponse
type BluetoothHCIExtendedInquiryResponse = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhciextendedinquiryresponsedatatype
type BluetoothHCIExtendedInquiryResponseDataType = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciextendedinquiryresult
type BluetoothHCIExtendedInquiryResult = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcifecrequired
type BluetoothHCIFECRequired = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcifailedcontactcount
type BluetoothHCIFailedContactCount = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcifailedcontactinfo
type BluetoothHCIFailedContactInfo = unsafe.Pointer

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
type BluetoothHCIInquiryAccessCode = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhciinquiryaccesscodecount
type BluetoothHCIInquiryAccessCodeCount = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciinquirylength
type BluetoothHCIInquiryLength = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciinquirymode
type BluetoothHCIInquiryMode = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciinquiryresult
type BluetoothHCIInquiryResult = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhciinquiryresults
type BluetoothHCIInquiryResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhciinquiryscantype
type BluetoothHCIInquiryScanType = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciinquirywithrssiresult
type BluetoothHCIInquiryWithRSSIResult = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhciinquirywithrssiresults
type BluetoothHCIInquiryWithRSSIResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcilebuffersize
type BluetoothHCILEBufferSize = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcilesupportedfeatures
type BluetoothHCILESupportedFeatures = BluetoothHCISupportedFeatures

// See: https://developer.apple.com/documentation/kernel/bluetoothhcileusedfeatures
type BluetoothHCILEUsedFeatures = BluetoothHCISupportedFeatures

// See: https://developer.apple.com/documentation/kernel/bluetoothhcilinkpolicysettings
type BluetoothHCILinkPolicySettings = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcilinkpolicysettingsinfo
type BluetoothHCILinkPolicySettingsInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcilinkquality
type BluetoothHCILinkQuality = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcilinkqualityinfo
type BluetoothHCILinkQualityInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcilinksupervisiontimeout
type BluetoothHCILinkSupervisionTimeout = unsafe.Pointer

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
type BluetoothHCIQualityOfServiceSetupParams = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcirssiinfo
type BluetoothHCIRSSIInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcirssivalue
type BluetoothHCIRSSIValue = int8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcireadextendedinquiryresponseresults
type BluetoothHCIReadExtendedInquiryResponseResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcireadlmphandleresults
type BluetoothHCIReadLMPHandleResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcireadlocaloobdataresults
type BluetoothHCIReadLocalOOBDataResults = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcireadstoredlinkkeysflag
type BluetoothHCIReadStoredLinkKeysFlag = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcireceivebandwidth
type BluetoothHCIReceiveBandwidth = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhcireceivecodecframesize
type BluetoothHCIReceiveCodecFrameSize = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcireceivecodingformat
type BluetoothHCIReceiveCodingFormat = uint64

// See: https://developer.apple.com/documentation/kernel/bluetoothhcirequestcallbackinfo
type BluetoothHCIRequestCallbackInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcirequestid
type BluetoothHCIRequestID = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhciresponsecount
type BluetoothHCIResponseCount = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciretransmissioneffort
type BluetoothHCIRetransmissionEffort = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcirole
type BluetoothHCIRole = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciroleinfo
type BluetoothHCIRoleInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhciscodatabytecount
type BluetoothHCISCODataByteCount = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhciscanactivity
type BluetoothHCIScanActivity = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcisetupsynchronousconnectionparams
type BluetoothHCISetupSynchronousConnectionParams = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcisignalid
type BluetoothHCISignalID = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhcisimplepairingmode
type BluetoothHCISimplePairingMode = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcisimplepairingoobdata
type BluetoothHCISimplePairingOOBData = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcisniffattemptcount
type BluetoothHCISniffAttemptCount = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcisnifftimeout
type BluetoothHCISniffTimeout = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothhcistatus
type BluetoothHCIStatus = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcistoredlinkkeysinfo
type BluetoothHCIStoredLinkKeysInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcisupportedcommands
type BluetoothHCISupportedCommands = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcisupportedfeatures
type BluetoothHCISupportedFeatures = unsafe.Pointer

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
type BluetoothHCITransmitPowerLevelInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhcitransmitpowerleveltype
type BluetoothHCITransmitPowerLevelType = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothhcitransportcommandid
type BluetoothHCITransportCommandID = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhcitransportid
type BluetoothHCITransportID = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhcivendorcommandselector
type BluetoothHCIVendorCommandSelector = uint32

// See: https://developer.apple.com/documentation/kernel/bluetoothhciversioninfo
type BluetoothHCIVersionInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothhciversions
type BluetoothHCIVersions = int

// See: https://developer.apple.com/documentation/kernel/bluetoothhcivoicesetting
type BluetoothHCIVoiceSetting = uint16

// See: https://developer.apple.com/documentation/kernel/bluetoothiocapability
type BluetoothIOCapability = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothiocapabilityresponse
type BluetoothIOCapabilityResponse = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothirk
type BluetoothIRK = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothkey
type BluetoothKey = string

// See: https://developer.apple.com/documentation/kernel/bluetoothkeyflag
type BluetoothKeyFlag = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothkeytype
type BluetoothKeyType = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothkeypressnotification
type BluetoothKeypressNotification = unsafe.Pointer

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
type BluetoothL2CAPGroupID = string

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
type BluetoothPINCode = unsafe.Pointer

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
type BluetoothReadClockInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothreasoncode
type BluetoothReasonCode = uint8

// See: https://developer.apple.com/documentation/kernel/bluetoothremotehostsupportedfeaturesnotification
type BluetoothRemoteHostSupportedFeaturesNotification = unsafe.Pointer

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
type BluetoothSynchronousConnectionInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothtransportinfo
type BluetoothTransportInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothtransportinfoptr
type BluetoothTransportInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothuserconfirmationrequest
type BluetoothUserConfirmationRequest = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bluetoothuserpasskeynotification
type BluetoothUserPasskeyNotification = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/boolean
type Boolean = bool

// See: https://developer.apple.com/documentation/kernel/boot_video
type Boot_Video = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/boot_videov1
type Boot_VideoV1 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bounds
type Bounds = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/byte
type Byte = uint8

// ByteCount is abst_ByteCount.
//
// See: https://developer.apple.com/documentation/kernel/bytecount
type ByteCount = uint

// See: https://developer.apple.com/documentation/kernel/byteptr
type BytePtr = uint8

// See: https://developer.apple.com/documentation/kernel/bytef
type Bytef = uint8

// See: https://developer.apple.com/documentation/kernel/cdatip
type CDATIP = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cdaudiostatus
type CDAudioStatus = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cddiscinfo
type CDDiscInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cdfeatures
type CDFeatures = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cdisrc
type CDISRC = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cdmcn
type CDMCN = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cdmsf
type CDMSF = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cdmediatype
type CDMediaType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cdpma
type CDPMA = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cdpmadescriptor
type CDPMADescriptor = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cdsectorarea
type CDSectorArea = int

// See: https://developer.apple.com/documentation/kernel/cdsectorsize
type CDSectorSize = int

// See: https://developer.apple.com/documentation/kernel/cdsectortype
type CDSectorType = int

// See: https://developer.apple.com/documentation/kernel/cdtext
type CDTEXT = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cdtextdescriptor
type CDTEXTDescriptor = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cdtoc
type CDTOC = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cdtocdescriptor
type CDTOCDescriptor = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cdtocformat
type CDTOCFormat = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cdtrackinfo
type CDTrackInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cdtrackinfoaddresstype
type CDTrackInfoAddressType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/complex
type COMPLEX = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/complex_split
type COMPLEX_SPLIT = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/csrnodeuniqueid
type CSRNodeUniqueID = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cs_blobindex
type CS_BlobIndex = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cs_codedirectory
type CS_CodeDirectory = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cs_genericblob
type CS_GenericBlob = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cs_superblob
type CS_SuperBlob = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/colorspec
type ColorSpec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/colorspecptr
type ColorSpecPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/consthfsunistr255param
type ConstHFSUniStr255Param = HFSUniStr255

// See: https://developer.apple.com/documentation/kernel/dasdmodeparameterblockdescriptor
type DASDModeParameterBlockDescriptor = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dclcallcommandproc
type DCLCallCommandProc = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dclcallcommandprocptr
type DCLCallCommandProcPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dclcallproc
type DCLCallProc = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dclcallprocdatatype
type DCLCallProcDataType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dclcallprocptr
type DCLCallProcPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dclcommand
type DCLCommand = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dclcommandptr
type DCLCommandPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dclcompilerdatatype
type DCLCompilerDataType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dcljump
type DCLJump = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dcljumpptr
type DCLJumpPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dcllabel
type DCLLabel = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dcllabelptr
type DCLLabelPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dclnudclleader
type DCLNuDCLLeader = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dclptrtimestamp
type DCLPtrTimeStamp = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dclptrtimestampptr
type DCLPtrTimeStampPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dclsettagsyncbits
type DCLSetTagSyncBits = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dclsettagsyncbitsptr
type DCLSetTagSyncBitsPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dcltimestamp
type DCLTimeStamp = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dcltimestampptr
type DCLTimeStampPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dcltransferbuffer
type DCLTransferBuffer = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dcltransferbufferptr
type DCLTransferBufferPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dcltransferpacket
type DCLTransferPacket = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dcltransferpacketptr
type DCLTransferPacketPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dclupdatedcllist
type DCLUpdateDCLList = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dclupdatedcllistptr
type DCLUpdateDCLListPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ddmap
type DDMap = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/double_complex
type DOUBLE_COMPLEX = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/double_complex_split
type DOUBLE_COMPLEX_SPLIT = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dpme
type DPME = unsafe.Pointer

// DSPComplex is used to hold a complex value.
//
// See: https://developer.apple.com/documentation/kernel/dspcomplex
type DSPComplex = unsafe.Pointer

// DSPDoubleComplex is used to hold a double-precision complex value.
//
// See: https://developer.apple.com/documentation/kernel/dspdoublecomplex
type DSPDoubleComplex = unsafe.Pointer

// DSPDoubleSplitComplex is used to represent a double-precision complex number when the real and imaginary parts are stored in separate arrays.
//
// See: https://developer.apple.com/documentation/kernel/dspdoublesplitcomplex
type DSPDoubleSplitComplex = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dspsplitcomplex
type DSPSplitComplex = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dtentry
type DTEntry = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dtentryiterator
type DTEntryIterator = uintptr

// See: https://developer.apple.com/documentation/kernel/dtentrynamebuf
type DTEntryNameBuf = int8

// See: https://developer.apple.com/documentation/kernel/dtmemorymaprange
type DTMemoryMapRange = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dtpropertyiterator
type DTPropertyIterator = uintptr

// See: https://developer.apple.com/documentation/kernel/dtpropertynamebuf
type DTPropertyNameBuf = int8

// See: https://developer.apple.com/documentation/kernel/dtsavedscopeptr
type DTSavedScopePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dvdauthenticationgrantidinfo
type DVDAuthenticationGrantIDInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dvdauthenticationsuccessflaginfo
type DVDAuthenticationSuccessFlagInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dvdbooktype
type DVDBookType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dvdcprmregioncode
type DVDCPRMRegionCode = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dvdchallengekeyinfo
type DVDChallengeKeyInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dvdcopyrightinfo
type DVDCopyrightInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dvddiscinfo
type DVDDiscInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dvddisckeyinfo
type DVDDiscKeyInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dvdfeatures
type DVDFeatures = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dvdkey1info
type DVDKey1Info = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dvdkey2info
type DVDKey2Info = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dvdkeyclass
type DVDKeyClass = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dvdkeyformat
type DVDKeyFormat = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dvdmanufacturinginfo
type DVDManufacturingInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dvdmediatype
type DVDMediaType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dvdphysicalformatinfo
type DVDPhysicalFormatInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dvdrzoneinfo
type DVDRZoneInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dvdrzoneinfoaddresstype
type DVDRZoneInfoAddressType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dvdregionplaybackcontrolinfo
type DVDRegionPlaybackControlInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dvdregionalplaybackcontrolscheme
type DVDRegionalPlaybackControlScheme = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dvdstructureformat
type DVDStructureFormat = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dvdtitlekeyinfo
type DVDTitleKeyInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/depthmode
type DepthMode = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/devicetreenode
type DeviceTreeNode = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/devicetreenodeproperty
type DeviceTreeNodeProperty = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/displayidtype
type DisplayIDType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/displaymodeid
type DisplayModeID = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/driverdescversion
type DriverDescVersion = uint32

// See: https://developer.apple.com/documentation/kernel/driverdescription
type DriverDescription = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/driverdescriptionptr
type DriverDescriptionPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/driverosruntime
type DriverOSRuntime = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/driverosruntimeptr
type DriverOSRuntimePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/driverosservice
type DriverOSService = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/driverosserviceptr
type DriverOSServicePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/driverserviceinfo
type DriverServiceInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/driverserviceinfoptr
type DriverServiceInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/drivertype
type DriverType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/drivertypeptr
type DriverTypePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/efi_boolean
type EFI_BOOLEAN = uint8

// See: https://developer.apple.com/documentation/kernel/efi_char16
type EFI_CHAR16 = int16

// See: https://developer.apple.com/documentation/kernel/efi_char32
type EFI_CHAR32 = int32

// See: https://developer.apple.com/documentation/kernel/efi_char64
type EFI_CHAR64 = int64

// See: https://developer.apple.com/documentation/kernel/efi_char8
type EFI_CHAR8 = int8

// See: https://developer.apple.com/documentation/kernel/efi_configuration_table_32
type EFI_CONFIGURATION_TABLE_32 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/efi_configuration_table_64
type EFI_CONFIGURATION_TABLE_64 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/efi_guid
type EFI_GUID = string

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
type EFI_MEMORY_DESCRIPTOR = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/efi_memory_type
type EFI_MEMORY_TYPE = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/efi_physical_address
type EFI_PHYSICAL_ADDRESS = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/efi_ptr32
type EFI_PTR32 = uint32

// See: https://developer.apple.com/documentation/kernel/efi_ptr64
type EFI_PTR64 = uint64

// See: https://developer.apple.com/documentation/kernel/efi_reset_type
type EFI_RESET_TYPE = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/efi_runtime_services_32
type EFI_RUNTIME_SERVICES_32 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/efi_runtime_services_64
type EFI_RUNTIME_SERVICES_64 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/efi_status
type EFI_STATUS = uint32

// See: https://developer.apple.com/documentation/kernel/efi_system_table_32
type EFI_SYSTEM_TABLE_32 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/efi_system_table_64
type EFI_SYSTEM_TABLE_64 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/efi_table_header
type EFI_TABLE_HEADER = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/efi_time
type EFI_TIME = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/efi_time_capabilities
type EFI_TIME_CAPABILITIES = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/efi_uint16
type EFI_UINT16 = uint16

// See: https://developer.apple.com/documentation/kernel/efi_uint32
type EFI_UINT32 = uint32

// See: https://developer.apple.com/documentation/kernel/efi_uint64
type EFI_UINT64 = uint64

// See: https://developer.apple.com/documentation/kernel/efi_uint8
type EFI_UINT8 = uint8

// See: https://developer.apple.com/documentation/kernel/efi_uintn
type EFI_UINTN = uint32

// See: https://developer.apple.com/documentation/kernel/efi_virtual_address
type EFI_VIRTUAL_ADDRESS = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/evscreen
type EVScreen = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/exbrightmessage
type EXBrightMessage = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/exbrightmessagetype
type EXBrightMessageType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/exdisplaypipehealthrecord
type EXDisplayPipeHealthRecord = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/exdisplaypipehealthreport
type EXDisplayPipeHealthReport = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/exdisplaypipeindicator
type EXDisplayPipeIndicator = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/exdisplaypipeindicatorparams
type EXDisplayPipeIndicatorParams = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/exdisplaypipesecuretestatus
type EXDisplayPipeSecureTEStatus = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/exdisplaypipestatus
type EXDisplayPipeStatus = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/efimemoryrange
type EfiMemoryRange = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/evcmd
type EvCmd = int

// See: https://developer.apple.com/documentation/kernel/evglobals
type EvGlobals = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/evoffsets
type EvOffsets = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/extendedsensecode
type ExtendedSenseCode = unsafe.Pointer

// FFTDirection is specifies whether to perform a forward or inverse FFT.
//
// See: https://developer.apple.com/documentation/kernel/fftdirection
type FFTDirection = int

// FFTRadix is the size of the FFT decomposition.
//
// See: https://developer.apple.com/documentation/kernel/fftradix
type FFTRadix = int

// FFTSetup is an opaque type that contains setup information for a given FFT transform.
//
// See: https://developer.apple.com/documentation/kernel/fftsetup
type FFTSetup = uintptr

// FFTSetupD is an opaque type that contains setup information for a given double-precision FFT transform.
//
// See: https://developer.apple.com/documentation/kernel/fftsetupd
type FFTSetupD = uintptr

// See: https://developer.apple.com/documentation/kernel/fwaddress
type FWAddress = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fwaddressptr
type FWAddressPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fwclientcommandid
type FWClientCommandID = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fwisochchannelforcestopnotificationproc
type FWIsochChannelForceStopNotificationProc = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fwisochchannelforcestopnotificationprocptr
type FWIsochChannelForceStopNotificationProcPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fwmultiisochreceivelistenerparams
type FWMultiIsochReceiveListenerParams = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fwsbp2logincompleteparams
type FWSBP2LoginCompleteParams = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fwsbp2logincompleteparamsptr
type FWSBP2LoginCompleteParamsPtr = uintptr

// See: https://developer.apple.com/documentation/kernel/fwsbp2loginresponse
type FWSBP2LoginResponse = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fwsbp2loginresponseptr
type FWSBP2LoginResponsePtr = uintptr

// See: https://developer.apple.com/documentation/kernel/fwsbp2logoutcompleteparams
type FWSBP2LogoutCompleteParams = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fwsbp2logoutcompleteparamsptr
type FWSBP2LogoutCompleteParamsPtr = uintptr

// See: https://developer.apple.com/documentation/kernel/fwsbp2notifyparams
type FWSBP2NotifyParams = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fwsbp2notifyparamsptr
type FWSBP2NotifyParamsPtr = uintptr

// See: https://developer.apple.com/documentation/kernel/fwsbp2reconnectparams
type FWSBP2ReconnectParams = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fwsbp2reconnectparamsptr
type FWSBP2ReconnectParamsPtr = uintptr

// See: https://developer.apple.com/documentation/kernel/fwsbp2statusblock
type FWSBP2StatusBlock = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fixed
type Fixed = uint32

// See: https://developer.apple.com/documentation/kernel/fixedptr
type FixedPtr = unsafe.Pointer

// Float32 is convenience type that represent a 32-bit floating point number.
//
// See: https://developer.apple.com/documentation/kernel/float32
type Float32 = float32

// Float64 is convenience type that represent a 64-bit floating point number.
//
// See: https://developer.apple.com/documentation/kernel/float64
type Float64 = float64

// See: https://developer.apple.com/documentation/kernel/fndrdirinfo
type FndrDirInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fndrfileinfo
type FndrFileInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fndropaqueinfo
type FndrOpaqueInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fourcharcode
type FourCharCode = uint32

// Fract is represents a type used by the Compression and Decompression API.
//
// See: https://developer.apple.com/documentation/kernel/fract
type Fract = uint32

// See: https://developer.apple.com/documentation/kernel/fractptr
type FractPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/gammatableid
type GammaTableID = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/gammatbl
type GammaTbl = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/gammatblptr
type GammaTblPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hfscatalogfile
type HFSCatalogFile = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hfscatalogfolder
type HFSCatalogFolder = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hfscatalogkey
type HFSCatalogKey = string

// See: https://developer.apple.com/documentation/kernel/hfscatalogthread
type HFSCatalogThread = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hfsextentdescriptor
type HFSExtentDescriptor = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hfsextentkey
type HFSExtentKey = string

// See: https://developer.apple.com/documentation/kernel/hfsextentrecord
type HFSExtentRecord = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hfsmasterdirectoryblock
type HFSMasterDirectoryBlock = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hfsplusattrdata
type HFSPlusAttrData = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hfsplusattrextents
type HFSPlusAttrExtents = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hfsplusattrforkdata
type HFSPlusAttrForkData = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hfsplusattrinlinedata
type HFSPlusAttrInlineData = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hfsplusattrkey
type HFSPlusAttrKey = string

// See: https://developer.apple.com/documentation/kernel/hfsplusbsdinfo
type HFSPlusBSDInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hfspluscatalogfile
type HFSPlusCatalogFile = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hfspluscatalogfolder
type HFSPlusCatalogFolder = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hfspluscatalogkey
type HFSPlusCatalogKey = string

// See: https://developer.apple.com/documentation/kernel/hfspluscatalogthread
type HFSPlusCatalogThread = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hfsplusextentdescriptor
type HFSPlusExtentDescriptor = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hfsplusextentkey
type HFSPlusExtentKey = string

// See: https://developer.apple.com/documentation/kernel/hfsplusextentrecord
type HFSPlusExtentRecord = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hfsplusforkdata
type HFSPlusForkData = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hfsplusvolumeheader
type HFSPlusVolumeHeader = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hfsunistr255
type HFSUniStr255 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hidreportcommandtype
type HIDReportCommandType = int

// See: https://developer.apple.com/documentation/kernel/handle
type Handle = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hardwarecursordescriptorptr
type HardwareCursorDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hardwarecursordescriptorrec
type HardwareCursorDescriptorRec = IOHardwareCursorDescriptor

// See: https://developer.apple.com/documentation/kernel/hardwarecursorinfoptr
type HardwareCursorInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hardwarecursorinforec
type HardwareCursorInfoRec = IOHardwareCursorInfo

// IIRChannel is constants that specify which channels a stereo biquadratic filter operates.
//
// See: https://developer.apple.com/documentation/kernel/iirchannel
type IIRChannel = int

// See: https://developer.apple.com/documentation/kernel/ioacpiaddressspaceid
type IOACPIAddressSpaceID = uint32

// See: https://developer.apple.com/documentation/kernel/ioatacompletionfunction
type IOATACompletionFunction = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioataregptr16
type IOATARegPtr16 = IOATAReg16

// See: https://developer.apple.com/documentation/kernel/ioataregptr32
type IOATARegPtr32 = IOATAReg32

// See: https://developer.apple.com/documentation/kernel/ioataregptr8
type IOATARegPtr8 = IOATAReg8

// See: https://developer.apple.com/documentation/kernel/ioavccommandresponse
type IOAVCCommandResponse = int

// See: https://developer.apple.com/documentation/kernel/ioavcframefields
type IOAVCFrameFields = int

// See: https://developer.apple.com/documentation/kernel/ioavcopcodes
type IOAVCOpcodes = int

// See: https://developer.apple.com/documentation/kernel/ioavcunittypes
type IOAVCUnitTypes = int

// See: https://developer.apple.com/documentation/kernel/ioaccelbounds
type IOAccelBounds = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioacceldeviceregion
type IOAccelDeviceRegion = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioaccelid
type IOAccelID = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioaccelsize
type IOAccelSize = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioaccelsurfaceinformation
type IOAccelSurfaceInformation = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioaccelsurfacereaddata
type IOAccelSurfaceReadData = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioaccelsurfacescaling
type IOAccelSurfaceScaling = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioaddressrange
type IOAddressRange = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioalignment
type IOAlignment = uint

// See: https://developer.apple.com/documentation/kernel/ioappletimingid
type IOAppleTimingID = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioasyncmethod
type IOAsyncMethod = *uintptr

// See: https://developer.apple.com/documentation/kernel/ioaudiobufferdatadescriptor
type IOAudioBufferDataDescriptor = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioaudioclientbuffer
type IOAudioClientBuffer = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioaudioclientbuffer64
type IOAudioClientBuffer64 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioaudioclientbufferextendedinfo
type IOAudioClientBufferExtendedInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioaudioclientbufferextendedinfo64
type IOAudioClientBufferExtendedInfo64 = unsafe.Pointer

// IOAudioDevicePowerState is identifies the power state of the audio device.
//
// See: https://developer.apple.com/documentation/kernel/ioaudiodevicepowerstate
type IOAudioDevicePowerState = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioaudioenginenotifications
type IOAudioEngineNotifications = unsafe.Pointer

// IOAudioEnginePosition is represents a position in an audio audio engine.
//
// See: https://developer.apple.com/documentation/kernel/ioaudioengineposition
type IOAudioEnginePosition = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioaudioenginetraps
type IOAudioEngineTraps = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioaudiosampleintervaldescriptor
type IOAudioSampleIntervalDescriptor = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioaudiosamplerate
type IOAudioSampleRate = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioaudiostreamdatadescriptor
type IOAudioStreamDataDescriptor = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioaudiostreamformat
type IOAudioStreamFormat = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioaudiostreamformatextension
type IOAudioStreamFormatExtension = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioaudiotimestamp
type IOAudioTimeStamp = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioblitcopyrectangle
type IOBlitCopyRectangle = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioblitcopyrectangles
type IOBlitCopyRectangles = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioblitcopyregion
type IOBlitCopyRegion = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioblitcursor
type IOBlitCursor = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioblitmemory
type IOBlitMemory = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioblitmemoryref
type IOBlitMemoryRef = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioblitoperation
type IOBlitOperation = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioblitrectangle
type IOBlitRectangle = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioblitrectangles
type IOBlitRectangles = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioblitscanlines
type IOBlitScanlines = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioblitsourcetype
type IOBlitSourceType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioblitsurface
type IOBlitSurface = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioblittype
type IOBlitType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioblitvertex
type IOBlitVertex = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioblitvertices
type IOBlitVertices = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iobytecount32
type IOByteCount32 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iobytecount64
type IOByteCount64 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iocachemode
type IOCacheMode = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iocolorcomponent
type IOColorComponent = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iocolorentry
type IOColorEntry = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iocommandcode
type IOCommandCode = uint32

// See: https://developer.apple.com/documentation/kernel/iocommandid
type IOCommandID = uintptr

// See: https://developer.apple.com/documentation/kernel/iocommandkind
type IOCommandKind = uint32

// See: https://developer.apple.com/documentation/kernel/iodataqueueclientdequeueentryblock
type IODataQueueClientDequeueEntryBlock = func(unsafe.Pointer, uintptr)

// See: https://developer.apple.com/documentation/kernel/iodataqueueclientenqueueentryblock
type IODataQueueClientEnqueueEntryBlock = func(unsafe.Pointer, uintptr)

// See: https://developer.apple.com/documentation/kernel/iodebuggerlockstate
type IODebuggerLockState = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iodetailedtiminginformation
type IODetailedTimingInformation = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iodetailedtiminginformationv1
type IODetailedTimingInformationV1 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iodetailedtiminginformationv2
type IODetailedTimingInformationV2 = unsafe.Pointer

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

// See: https://developer.apple.com/documentation/kernel/iodisplaymodeinformation
type IODisplayModeInformation = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iodisplayproductid
type IODisplayProductID = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iodisplayscalerinformation
type IODisplayScalerInformation = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iodisplaytimingrange
type IODisplayTimingRange = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iodisplaytimingrangev1
type IODisplayTimingRangeV1 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iodisplaytimingrangev2
type IODisplayTimingRangeV2 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iodisplayvendorid
type IODisplayVendorID = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioenetmulticastmode
type IOEnetMulticastMode = bool

// See: https://developer.apple.com/documentation/kernel/ioenetpromiscuousmode
type IOEnetPromiscuousMode = bool

// See: https://developer.apple.com/documentation/kernel/ioethernetcontrolleravbstate
type IOEthernetControllerAVBState = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioethernetcontrolleravbstateevent
type IOEthernetControllerAVBStateEvent = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioethernetcontrolleravbtimesyncsupport
type IOEthernetControllerAVBTimeSyncSupport = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iofbcursorcontrolattribute
type IOFBCursorControlAttribute = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iofbcursorcontrolcallouts
type IOFBCursorControlCallouts = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iofbcursorref
type IOFBCursorRef uintptr

// See: https://developer.apple.com/documentation/kernel/iofbdplinkconfig
type IOFBDPLinkConfig = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iofbdisplaymodedescription
type IOFBDisplayModeDescription = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iofbhdrmetadata
type IOFBHDRMetaData = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iofbhdrmetadatav1
type IOFBHDRMetaDataV1 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iofwavcasynccommandstate
type IOFWAVCAsyncCommandState = int

// See: https://developer.apple.com/documentation/kernel/iofwavcplugtypes
type IOFWAVCPlugTypes = int

// See: https://developer.apple.com/documentation/kernel/iofwavcsubunitplugmessages
type IOFWAVCSubunitPlugMessages = int

// See: https://developer.apple.com/documentation/kernel/iofwdclnotificationtype
type IOFWDCLNotificationType = int

// See: https://developer.apple.com/documentation/kernel/iofwduplicateguidrec
type IOFWDuplicateGUIDRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iofwisochportoptions
type IOFWIsochPortOptions = uint

// See: https://developer.apple.com/documentation/kernel/iofwisochresourceflags
type IOFWIsochResourceFlags = uint

// See: https://developer.apple.com/documentation/kernel/iofwrequestrefcon
type IOFWRequestRefCon = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iofwspeed
type IOFWSpeed = int

// See: https://developer.apple.com/documentation/kernel/iofirewiresessionref
type IOFireWireSessionRef = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iofixed1616
type IOFixed1616 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iofixedpoint32
type IOFixedPoint32 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iogbounds
type IOGBounds = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iogpoint
type IOGPoint = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iogsize
type IOGSize = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iohidaccelerationalgorithmtype
type IOHIDAccelerationAlgorithmType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iohidbiometriceventtype
type IOHIDBiometricEventType = uint32

// See: https://developer.apple.com/documentation/kernel/iohidbuttonmodes
type IOHIDButtonModes = int

// IOHIDCompletion is struct specifying action to perform when set/get report completes.
//
// See: https://developer.apple.com/documentation/kernel/iohidcompletion
type IOHIDCompletion = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iohiddigitizerstylusdata
type IOHIDDigitizerStylusData = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iohiddigitizertouchdata
type IOHIDDigitizerTouchData = unsafe.Pointer

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
type IOHIDKind = uint

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
type IOHardwareCursorDescriptor = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iohardwarecursorinfo
type IOHardwareCursorInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iohistreportinfo
type IOHistReportInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iohistogramreportvalues
type IOHistogramReportValues = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iohistogramsegmentconfig
type IOHistogramSegmentConfig = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioindex
type IOIndex = int32

// See: https://developer.apple.com/documentation/kernel/iointerruptactionblock
type IOInterruptActionBlock = func(*IOService, int)

// See: https://developer.apple.com/documentation/kernel/iointerruptsource
type IOInterruptSource = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iointerruptstate
type IOInterruptState = bool

// See: https://developer.apple.com/documentation/kernel/iointerruptvector
type IOInterruptVector = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iointerruptvectornumber
type IOInterruptVectorNumber = int32

// See: https://developer.apple.com/documentation/kernel/ioitemcount
type IOItemCount = uint

// See: https://developer.apple.com/documentation/kernel/iokitdiagnosticsparameters
type IOKitDiagnosticsParameters = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iolock
type IOLock = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iolockstate
type IOLockState = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iologicaladdress
type IOLogicalAddress = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iomediaattributemask
type IOMediaAttributeMask = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iomediastate
type IOMediaState = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iomessage
type IOMessage = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iomethod
type IOMethod = *uintptr

// See: https://developer.apple.com/documentation/kernel/ionamedvalue
type IONamedValue = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ionormdistreportvalues
type IONormDistReportValues = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ionotificationref
type IONotificationRef uintptr

// See: https://developer.apple.com/documentation/kernel/iooptionbits
type IOOptionBits = uint32

// See: https://developer.apple.com/documentation/kernel/iooutputaction
type IOOutputAction = *uintptr

// See: https://developer.apple.com/documentation/kernel/iopcidevicecrashnotification_t
type IOPCIDeviceCrashNotification_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iopmcalendarstruct
type IOPMCalendarStruct = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iopmdriverassertionid
type IOPMDriverAssertionID = uint64

// See: https://developer.apple.com/documentation/kernel/iopmdriverassertionlevel
type IOPMDriverAssertionLevel = uint32

// See: https://developer.apple.com/documentation/kernel/iopmdriverassertiontype
type IOPMDriverAssertionType = uint64

// IOPMPowerFlags is bits are used in defining capabilityFlags, inputPowerRequirements, and outputPowerCharacter in the IOPMPowerState structure.
//
// See: https://developer.apple.com/documentation/kernel/iopmpowerflags
type IOPMPowerFlags = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iopmpowerstate
type IOPMPowerState = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iopacketbufferconstraints
type IOPacketBufferConstraints = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iophysicaladdress
type IOPhysicalAddress = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iophysicaladdress32
type IOPhysicalAddress32 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iophysicaladdress64
type IOPhysicalAddress64 = uint64

// See: https://developer.apple.com/documentation/kernel/iophysicallength
type IOPhysicalLength = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iophysicallength32
type IOPhysicalLength32 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iophysicallength64
type IOPhysicalLength64 = uint64

// See: https://developer.apple.com/documentation/kernel/iophysicalrange
type IOPhysicalRange = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iopixelaperture
type IOPixelAperture = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iopixelencoding
type IOPixelEncoding = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iopixelinformation
type IOPixelInformation = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iopowerstatechangenotification
type IOPowerStateChangeNotification = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iopropertyname
type IOPropertyName = int8

// See: https://developer.apple.com/documentation/kernel/iorpc
type IORPC = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iorpcmessage
type IORPCMessage = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iorpcmessageerrorreturncontent
type IORPCMessageErrorReturnContent = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iorpcmessagemach
type IORPCMessageMach = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iorwlock
type IORWLock = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iorangescalar
type IORangeScalar = uint

// See: https://developer.apple.com/documentation/kernel/iorecursivelock
type IORecursiveLock = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioregistryplanename
type IORegistryPlaneName = int8

// See: https://developer.apple.com/documentation/kernel/ioreportcategories
type IOReportCategories = uint16

// See: https://developer.apple.com/documentation/kernel/ioreportchannel
type IOReportChannel = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioreportchannellist
type IOReportChannelList = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioreportchanneltype
type IOReportChannelType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioreportconfigureaction
type IOReportConfigureAction = uint32

// See: https://developer.apple.com/documentation/kernel/ioreportelement
type IOReportElement = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioreportelementvalues
type IOReportElementValues = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioreportformat
type IOReportFormat = uint8

// See: https://developer.apple.com/documentation/kernel/ioreportinterest
type IOReportInterest = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioreportinterestlist
type IOReportInterestList = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioreportquantity
type IOReportQuantity = uint8

// See: https://developer.apple.com/documentation/kernel/ioreportscalefactor
type IOReportScaleFactor = uint64

// See: https://developer.apple.com/documentation/kernel/ioreportunit
type IOReportUnit = uint64

// See: https://developer.apple.com/documentation/kernel/ioreportunits
type IOReportUnits = uint64

// See: https://developer.apple.com/documentation/kernel/ioreportupdateaction
type IOReportUpdateAction = uint32

// See: https://developer.apple.com/documentation/kernel/ioreturn
type IOReturn = int

// See: https://developer.apple.com/documentation/kernel/ioselect
type IOSelect = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ioserviceapplierblock
type IOServiceApplierBlock = func(*IOService)

// See: https://developer.apple.com/documentation/kernel/ioserviceinteresthandlerblock
type IOServiceInterestHandlerBlock = func(uint32, *IOService, unsafe.Pointer, uintptr) int

// See: https://developer.apple.com/documentation/kernel/ioservicematchingnotificationhandlerblock
type IOServiceMatchingNotificationHandlerBlock = func(*IOService, *IONotifier) bool

// See: https://developer.apple.com/documentation/kernel/ioservicename
type IOServiceName = int8

// See: https://developer.apple.com/documentation/kernel/ioservicenotificationblock
type IOServiceNotificationBlock = func(uint64, *IOService, uint64)

// See: https://developer.apple.com/documentation/kernel/iosimplearrayreportvalues
type IOSimpleArrayReportValues = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iosimplelock
type IOSimpleLock = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iosimplereportvalues
type IOSimpleReportValues = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iostatenotificationhandler
type IOStateNotificationHandler = func() Kern_return_t

// See: https://developer.apple.com/documentation/kernel/iostatenotificationlistenerref
type IOStateNotificationListenerRef uintptr

// See: https://developer.apple.com/documentation/kernel/iostatereportinfo
type IOStateReportInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iostatereportvalues
type IOStateReportValues = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iostorageaccess
type IOStorageAccess = uint32

// See: https://developer.apple.com/documentation/kernel/iostoragegetprovisionstatusoptions
type IOStorageGetProvisionStatusOptions = uint64

// See: https://developer.apple.com/documentation/kernel/iostorageoptions
type IOStorageOptions = uint16

// See: https://developer.apple.com/documentation/kernel/iostoragepriority
type IOStoragePriority = uint8

// See: https://developer.apple.com/documentation/kernel/iostoragesynchronizeoptions
type IOStorageSynchronizeOptions = uint32

// See: https://developer.apple.com/documentation/kernel/iostorageunmapoptions
type IOStorageUnmapOptions = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iostreammode
type IOStreamMode = uint

// See: https://developer.apple.com/documentation/kernel/iotvector
type IOTVector = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iothread
type IOThread = Thread_t

// See: https://developer.apple.com/documentation/kernel/iotiminginformation
type IOTimingInformation = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iotrap
type IOTrap = *uintptr

// IOUSB20HubDescriptor is a structure that defines the descriptor for a USB 2.0 hub.
//
// See: https://developer.apple.com/documentation/kernel/iousb20hubdescriptor
type IOUSB20HubDescriptor = unsafe.Pointer

// IOUSB3HubDescriptor is a structure that defines the descriptor for a USB 3.0 hub.
//
// See: https://developer.apple.com/documentation/kernel/iousb3hubdescriptor
type IOUSB3HubDescriptor = unsafe.Pointer

// IOUSBBOSDescriptor is the structure for storing a binary object store (BOS) descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbbosdescriptor
type IOUSBBOSDescriptor = unsafe.Pointer

// IOUSBBOSDescriptorPtr is a pointer to a structure for storing a binary object store (BOS) descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbbosdescriptorptr
type IOUSBBOSDescriptorPtr = unsafe.Pointer

// IOUSBBulkPipeReq is the structure that represents a bulk pipe request.
//
// See: https://developer.apple.com/documentation/kernel/iousbbulkpipereq
type IOUSBBulkPipeReq = unsafe.Pointer

// IOUSBCompletion is the structure that specifies the action to perform when the USB input/output request completes.
//
// See: https://developer.apple.com/documentation/kernel/iousbcompletion
type IOUSBCompletion = unsafe.Pointer

// IOUSBCompletionWithTimeStamp is a structure specifying action to perform when the USB input/output request completes.
//
// See: https://developer.apple.com/documentation/kernel/iousbcompletionwithtimestamp
type IOUSBCompletionWithTimeStamp = unsafe.Pointer

// IOUSBConfigurationDescHeader is the header of a configuration descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbconfigurationdescheader
type IOUSBConfigurationDescHeader = unsafe.Pointer

// IOUSBConfigurationDescHeaderPtr is a pointer to a configuration descriptor header.
//
// See: https://developer.apple.com/documentation/kernel/iousbconfigurationdescheaderptr
type IOUSBConfigurationDescHeaderPtr = unsafe.Pointer

// IOUSBConfigurationDescriptor is the structure for storing a USB configuration descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbconfigurationdescriptor
type IOUSBConfigurationDescriptor = unsafe.Pointer

// IOUSBConfigurationDescriptorPtr is a pointer to a configuration descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbconfigurationdescriptorptr
type IOUSBConfigurationDescriptorPtr = unsafe.Pointer

// IOUSBDFUDescriptor is a structure that defines the USB device firmware update descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbdfudescriptor
type IOUSBDFUDescriptor = unsafe.Pointer

// IOUSBDFUDescriptorPtr is a pointer to a structure that defines the USB device firmware update descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbdfudescriptorptr
type IOUSBDFUDescriptorPtr = unsafe.Pointer

// IOUSBDescriptor is the base descriptor type.
//
// See: https://developer.apple.com/documentation/kernel/iousbdescriptor
type IOUSBDescriptor = unsafe.Pointer

// IOUSBDescriptorHeader is the base descriptor header.
//
// See: https://developer.apple.com/documentation/kernel/iousbdescriptorheader
type IOUSBDescriptorHeader = unsafe.Pointer

// IOUSBDescriptorHeaderPtr is a pointer to a USB descriptor header.
//
// See: https://developer.apple.com/documentation/kernel/iousbdescriptorheaderptr
type IOUSBDescriptorHeaderPtr = unsafe.Pointer

// IOUSBDevReqOOL is an internal structure to pass parameters between IOUSBLib and UserClient.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevreqool
type IOUSBDevReqOOL = unsafe.Pointer

// IOUSBDevReqOOLTO is an internal structure to pass parameters between IOUSBLib and UserClient.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevreqoolto
type IOUSBDevReqOOLTO = unsafe.Pointer

// IOUSBDevRequest is a structure that defines a standard device request.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevrequest
type IOUSBDevRequest = unsafe.Pointer

// IOUSBDevRequestTO is a structure that defines a standard device request with timeout.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevrequestto
type IOUSBDevRequestTO = unsafe.Pointer

// IOUSBDeviceCapabilityBillboard is the structure for the billboard device capability.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitybillboard
type IOUSBDeviceCapabilityBillboard = unsafe.Pointer

// IOUSBDeviceCapabilityBillboardAltConfig is the structure for the billboard alternative configuration device capability.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitybillboardaltconfig
type IOUSBDeviceCapabilityBillboardAltConfig = unsafe.Pointer

// IOUSBDeviceCapabilityBillboardAltConfigCompatibility is the structure for the billboard alternative configuration compatibility device capability.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitybillboardaltconfigcompatibility
type IOUSBDeviceCapabilityBillboardAltConfigCompatibility = unsafe.Pointer

// IOUSBDeviceCapabilityBillboardAltConfigPtr is a pointer to a USB device capability billboard alternative configuration structure.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitybillboardaltconfigptr
type IOUSBDeviceCapabilityBillboardAltConfigPtr = unsafe.Pointer

// IOUSBDeviceCapabilityBillboardAltMode is the structure for the billboard alternative mode device capability.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitybillboardaltmode
type IOUSBDeviceCapabilityBillboardAltMode = uint

// IOUSBDeviceCapabilityBillboardAltModePtr is a pointer to a USB device capability billboard alternative mode structure.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitybillboardaltmodeptr
type IOUSBDeviceCapabilityBillboardAltModePtr = unsafe.Pointer

// IOUSBDeviceCapabilityBillboardPtr is a pointer to a USB device capability billboard object.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitybillboardptr
type IOUSBDeviceCapabilityBillboardPtr = unsafe.Pointer

// IOUSBDeviceCapabilityContainerID is the structure for the container ID device capability.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitycontainerid
type IOUSBDeviceCapabilityContainerID = string

// IOUSBDeviceCapabilityContainerIDPtr is a pointer to a USB device capability container ID.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitycontaineridptr
type IOUSBDeviceCapabilityContainerIDPtr = unsafe.Pointer

// IOUSBDeviceCapabilityDescriptorHeader is the device capability descriptor header.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitydescriptorheader
type IOUSBDeviceCapabilityDescriptorHeader = unsafe.Pointer

// IOUSBDeviceCapabilityDescriptorHeaderPtr is a pointer to a device capability descriptor header.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitydescriptorheaderptr
type IOUSBDeviceCapabilityDescriptorHeaderPtr = unsafe.Pointer

// IOUSBDeviceCapabilitySuperSpeedPlusUSB is the structure for the SuperSpeedPlus USB device capability.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitysuperspeedplususb
type IOUSBDeviceCapabilitySuperSpeedPlusUSB = unsafe.Pointer

// IOUSBDeviceCapabilitySuperSpeedPlusUSBPtr is a pointer to a SuperSpeedPlus USB device capability structure.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitysuperspeedplususbptr
type IOUSBDeviceCapabilitySuperSpeedPlusUSBPtr = unsafe.Pointer

// IOUSBDeviceCapabilitySuperSpeedUSB is the structure for the SuperSpeed USB device capability.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitysuperspeedusb
type IOUSBDeviceCapabilitySuperSpeedUSB = unsafe.Pointer

// IOUSBDeviceCapabilitySuperSpeedUSBPtr is a pointer to a SuperSpeed USB device capability structure.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilitysuperspeedusbptr
type IOUSBDeviceCapabilitySuperSpeedUSBPtr = unsafe.Pointer

// IOUSBDeviceCapabilityUSB2Extension is the structure for the USB 2.0 extension device capability.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilityusb2extension
type IOUSBDeviceCapabilityUSB2Extension = unsafe.Pointer

// IOUSBDeviceCapabilityUSB2ExtensionPtr is a pointer to a USB 2.0 extension device capability structure.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicecapabilityusb2extensionptr
type IOUSBDeviceCapabilityUSB2ExtensionPtr = unsafe.Pointer

// IOUSBDeviceDescriptor is the structure for storing a USB device descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicedescriptor
type IOUSBDeviceDescriptor = unsafe.Pointer

// IOUSBDeviceDescriptorPtr is a pointer to a USB device descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicedescriptorptr
type IOUSBDeviceDescriptorPtr = unsafe.Pointer

// IOUSBDeviceQualifierDescriptor is the structure for describing a high-speed capable USB device.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicequalifierdescriptor
type IOUSBDeviceQualifierDescriptor = unsafe.Pointer

// IOUSBDeviceQualifierDescriptorPtr is a pointer to a qualifier descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicequalifierdescriptorptr
type IOUSBDeviceQualifierDescriptorPtr = unsafe.Pointer

// IOUSBDeviceRequest is a structure that defines a standard device request.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicerequest
type IOUSBDeviceRequest = unsafe.Pointer

// IOUSBDeviceRequestPtr is a pointer to a structure that defines a standard device request.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicerequestptr
type IOUSBDeviceRequestPtr = unsafe.Pointer

// IOUSBDeviceRequestSetSELData is the structure for receiving system exit latency values.
//
// See: https://developer.apple.com/documentation/kernel/iousbdevicerequestsetseldata
type IOUSBDeviceRequestSetSELData = unsafe.Pointer

// IOUSBEndpointDescriptor is the structure for storing an endpoint descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbendpointdescriptor
type IOUSBEndpointDescriptor = unsafe.Pointer

// IOUSBEndpointDescriptorPtr is a pointer to the endpoint descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbendpointdescriptorptr
type IOUSBEndpointDescriptorPtr = unsafe.Pointer

// IOUSBEndpointProperties is a structure that holds USB endpoint properties.
//
// See: https://developer.apple.com/documentation/kernel/iousbendpointproperties
type IOUSBEndpointProperties = unsafe.Pointer

// IOUSBEndpointPropertiesPtr is a pointer to an endpoint properties object.
//
// See: https://developer.apple.com/documentation/kernel/iousbendpointpropertiesptr
type IOUSBEndpointPropertiesPtr = unsafe.Pointer

// IOUSBFindEndpointRequest is the structure that represents an endoint request to locate.
//
// See: https://developer.apple.com/documentation/kernel/iousbfindendpointrequest
type IOUSBFindEndpointRequest = unsafe.Pointer

// IOUSBFindInterfaceRequest is the structure for finding an interface request.
//
// See: https://developer.apple.com/documentation/kernel/iousbfindinterfacerequest
type IOUSBFindInterfaceRequest = unsafe.Pointer

// IOUSBGetFrameStruct is a structure that contains frame information.
//
// See: https://developer.apple.com/documentation/kernel/iousbgetframestruct
type IOUSBGetFrameStruct = unsafe.Pointer

// IOUSBHIDData is data related to the mouse and keyboard.
//
// See: https://developer.apple.com/documentation/kernel/iousbhiddata
type IOUSBHIDData = unsafe.Pointer

// IOUSBHIDDataPtr is a pointer to a structure related to mouse and keyboard data.
//
// See: https://developer.apple.com/documentation/kernel/iousbhiddataptr
type IOUSBHIDDataPtr = unsafe.Pointer

// IOUSBHIDDescriptor is a structure that defines the USB HID descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbhiddescriptor
type IOUSBHIDDescriptor = unsafe.Pointer

// IOUSBHIDDescriptorPtr is a pointer to a structure that defines the USB HID descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbhiddescriptorptr
type IOUSBHIDDescriptorPtr = unsafe.Pointer

// IOUSBHIDReportDesc is a structure that defines the USB HID report descriptor header.
//
// See: https://developer.apple.com/documentation/kernel/iousbhidreportdesc
type IOUSBHIDReportDesc = unsafe.Pointer

// IOUSBHIDReportDescPtr is a pointer to a structure that defines the USB HID report descriptor header.
//
// See: https://developer.apple.com/documentation/kernel/iousbhidreportdescptr
type IOUSBHIDReportDescPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iousbhostcimessage
type IOUSBHostCIMessage = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iousbhostciuserclientversion
type IOUSBHostCIUserClientVersion = int

// IOUSBHostIOSourceClientRecordLink is a structure that represents a USB host input/output source client record entry.
//
// See: https://developer.apple.com/documentation/kernel/iousbhostiosourceclientrecordlink
type IOUSBHostIOSourceClientRecordLink = unsafe.Pointer

// IOUSBHostIOSourceClientRecordList is a structure that represents a list of USB host input/output source client records.
//
// See: https://developer.apple.com/documentation/kernel/iousbhostiosourceclientrecordlist
type IOUSBHostIOSourceClientRecordList = unsafe.Pointer

// IOUSBHubDescriptor is a structure that defines the descriptor for a USB hub.
//
// See: https://developer.apple.com/documentation/kernel/iousbhubdescriptor
type IOUSBHubDescriptor = unsafe.Pointer

// IOUSBHubPortReEnumerateParam is a structure for USB hub port reenumeration.
//
// See: https://developer.apple.com/documentation/kernel/iousbhubportreenumerateparam
type IOUSBHubPortReEnumerateParam = unsafe.Pointer

// IOUSBHubPortStatus is a structure that contains the USB hub port status.
//
// See: https://developer.apple.com/documentation/kernel/iousbhubportstatus
type IOUSBHubPortStatus = IOUSBHubStatus

// IOUSBHubStatus is a structure that represents the USB hub status.
//
// See: https://developer.apple.com/documentation/kernel/iousbhubstatus
type IOUSBHubStatus = unsafe.Pointer

// IOUSBHubStatusPtr is a pointer to a USB hub status structure.
//
// See: https://developer.apple.com/documentation/kernel/iousbhubstatusptr
type IOUSBHubStatusPtr = unsafe.Pointer

// IOUSBInterfaceAssociationDescriptor is the descriptor that associates multiple interfaces to the same function.
//
// See: https://developer.apple.com/documentation/kernel/iousbinterfaceassociationdescriptor
type IOUSBInterfaceAssociationDescriptor = unsafe.Pointer

// IOUSBInterfaceAssociationDescriptorPtr is a pointer to a USB interface association descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbinterfaceassociationdescriptorptr
type IOUSBInterfaceAssociationDescriptorPtr = unsafe.Pointer

// IOUSBInterfaceDescriptor is a descriptor for a specific interface of a USB device.
//
// See: https://developer.apple.com/documentation/kernel/iousbinterfacedescriptor
type IOUSBInterfaceDescriptor = unsafe.Pointer

// IOUSBInterfaceDescriptorPtr is a pointer to a USB interface descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbinterfacedescriptorptr
type IOUSBInterfaceDescriptorPtr = unsafe.Pointer

// IOUSBIsocCompletion is a structure specifying the action to perform when an isochronous USB input/output operation completes.
//
// See: https://developer.apple.com/documentation/kernel/iousbisoccompletion
type IOUSBIsocCompletion = unsafe.Pointer

// IOUSBIsocFrame is a structure for encoding information about a single frame in an isochronous transfer.
//
// See: https://developer.apple.com/documentation/kernel/iousbisocframe
type IOUSBIsocFrame = unsafe.Pointer

// IOUSBIsocStruct is an internal structure to pass parameters between IOUSBLib and UserClient.
//
// See: https://developer.apple.com/documentation/kernel/iousbisocstruct
type IOUSBIsocStruct = unsafe.Pointer

// IOUSBIsochronousFrame is a structure representing a single frame in an isochronous transfer.
//
// See: https://developer.apple.com/documentation/kernel/iousbisochronousframe
type IOUSBIsochronousFrame = unsafe.Pointer

// IOUSBKeyboardData is a structure containing USB keyboard data.
//
// See: https://developer.apple.com/documentation/kernel/iousbkeyboarddata
type IOUSBKeyboardData = unsafe.Pointer

// IOUSBKeyboardDataPtr is a pointer to a structure containing USB keyboard data.
//
// See: https://developer.apple.com/documentation/kernel/iousbkeyboarddataptr
type IOUSBKeyboardDataPtr = unsafe.Pointer

// IOUSBLowLatencyIsocCompletion is the function that executes when the low-latency isochronous USB input/output request completes.
//
// See: https://developer.apple.com/documentation/kernel/iousblowlatencyisoccompletion
type IOUSBLowLatencyIsocCompletion = unsafe.Pointer

// IOUSBLowLatencyIsocFrame is a structure for encoding information about each low-latency isochronous frame.
//
// See: https://developer.apple.com/documentation/kernel/iousblowlatencyisocframe
type IOUSBLowLatencyIsocFrame = unsafe.Pointer

// IOUSBLowLatencyIsocStruct is an internal structure to pass parameters between IOUSBLib and UserClient.
//
// See: https://developer.apple.com/documentation/kernel/iousblowlatencyisocstruct
type IOUSBLowLatencyIsocStruct = unsafe.Pointer

// IOUSBMatch is a structure for matching USB devices.
//
// See: https://developer.apple.com/documentation/kernel/iousbmatch
type IOUSBMatch = unsafe.Pointer

// IOUSBMouseData is a structure containing USB mouse data.
//
// See: https://developer.apple.com/documentation/kernel/iousbmousedata
type IOUSBMouseData = unsafe.Pointer

// IOUSBMouseDataPtr is a pointer to a structure containing USB mouse data.
//
// See: https://developer.apple.com/documentation/kernel/iousbmousedataptr
type IOUSBMouseDataPtr = unsafe.Pointer

// IOUSBPlatformCapabilityDescriptor is the structure for the platform capability descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbplatformcapabilitydescriptor
type IOUSBPlatformCapabilityDescriptor = unsafe.Pointer

// IOUSBPlatformCapabilityDescriptorPtr is a pointer to a USB platform capability descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbplatformcapabilitydescriptorptr
type IOUSBPlatformCapabilityDescriptorPtr = unsafe.Pointer

// IOUSBStringDescriptor is the structure for storing a string descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbstringdescriptor
type IOUSBStringDescriptor = unsafe.Pointer

// IOUSBStringDescriptorPtr is a pointer to a string descriptor structure.
//
// See: https://developer.apple.com/documentation/kernel/iousbstringdescriptorptr
type IOUSBStringDescriptorPtr = unsafe.Pointer

// IOUSBSuperSpeedEndpointCompanionDescriptor is the descriptor for a SuperSpeed USB endpoint companion.
//
// See: https://developer.apple.com/documentation/kernel/iousbsuperspeedendpointcompaniondescriptor
type IOUSBSuperSpeedEndpointCompanionDescriptor = unsafe.Pointer

// IOUSBSuperSpeedEndpointCompanionDescriptorPtr is a pointer to a SuperSpeed USB endpoint companion descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbsuperspeedendpointcompaniondescriptorptr
type IOUSBSuperSpeedEndpointCompanionDescriptorPtr = unsafe.Pointer

// IOUSBSuperSpeedHubDescriptor is a structure that defines the descriptor for a SuperSpeed USB hub.
//
// See: https://developer.apple.com/documentation/kernel/iousbsuperspeedhubdescriptor
type IOUSBSuperSpeedHubDescriptor = unsafe.Pointer

// IOUSBSuperSpeedPlusIsochronousEndpointCompanionDescriptor is the descriptor for a SuperSpeedPlus isochronous USB endpoint companion.
//
// See: https://developer.apple.com/documentation/kernel/iousbsuperspeedplusisochronousendpointcompaniondescriptor
type IOUSBSuperSpeedPlusIsochronousEndpointCompanionDescriptor = unsafe.Pointer

// IOUSBSuperSpeedPlusIsochronousEndpointCompanionDescriptorPtr is a pointer to a SuperSpeedPlus isochronous USB endpoint companion descriptor.
//
// See: https://developer.apple.com/documentation/kernel/iousbsuperspeedplusisochronousendpointcompaniondescriptorptr
type IOUSBSuperSpeedPlusIsochronousEndpointCompanionDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iouserclientasyncargumentsarray
type IOUserClientAsyncArgumentsArray = uint64

// See: https://developer.apple.com/documentation/kernel/iouserclientasyncreferencearray
type IOUserClientAsyncReferenceArray = uint64

// See: https://developer.apple.com/documentation/kernel/iouserclientscalararray
type IOUserClientScalarArray = uint64

// See: https://developer.apple.com/documentation/kernel/ioversion
type IOVersion = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iovideodevicenotification
type IOVideoDeviceNotification = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iovideodevicenotificationmessage
type IOVideoDeviceNotificationMessage = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iovideostreamdescription
type IOVideoStreamDescription = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/iovirtualaddress
type IOVirtualAddress = Mach_vm_address_t

// See: https://developer.apple.com/documentation/kernel/iovirtualrange
type IOVirtualRange = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/interruptserviceidptr
type InterruptServiceIDPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/interruptserviceidtype
type InterruptServiceIDType = uintptr

// See: https://developer.apple.com/documentation/kernel/interruptservicetype
type InterruptServiceType = uint32

// See: https://developer.apple.com/documentation/kernel/journalinfoblock
type JournalInfoBlock = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kuncusernotificationid
type KUNCUserNotificationID = uintptr

// See: https://developer.apple.com/documentation/kernel/kernelid
type KernelID = uintptr

// See: https://developer.apple.com/documentation/kernel/longlbamodeparameterblockdescriptor
type LongLBAModeParameterBlockDescriptor = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/lowlatencyuserbufferinfo
type LowLatencyUserBufferInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/lowlatencyuserbufferinfov2
type LowLatencyUserBufferInfoV2 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/lowlatencyuserbufferinfov3
type LowLatencyUserBufferInfoV3 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/md5_ctx
type MD5_CTX = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/masteraudiofunctions
type MasterAudioFunctions = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mastermuteupdate
type MasterMuteUpdate = bool

// See: https://developer.apple.com/documentation/kernel/mastervolumeupdate
type MasterVolumeUpdate = uint16

// See: https://developer.apple.com/documentation/kernel/modepageformatheader
type ModePageFormatHeader = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/modeparameterblockdescriptor
type ModeParameterBlockDescriptor = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ndr_record_t
type NDR_record_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nxeqelement
type NXEQElement = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nxevent
type NXEvent = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nxeventdata
type NXEventData = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nxeventext
type NXEventExt = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nxeventextension
type NXEventExtension = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nxeventptr
type NXEventPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nxeventsystemdevice
type NXEventSystemDevice = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nxeventsystemdevicelist
type NXEventSystemDeviceList = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nxeventsysteminfodata
type NXEventSystemInfoData = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nxeventsysteminfotype
type NXEventSystemInfoType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nxkeymapping
type NXKeyMapping = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nxmousebutton
type NXMouseButton = int

// See: https://developer.apple.com/documentation/kernel/nxmousescaling
type NXMouseScaling = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nxparsedkeymapping
type NXParsedKeyMapping = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nxswappeddouble
type NXSwappedDouble = uint64

// See: https://developer.apple.com/documentation/kernel/nxswappedfloat
type NXSwappedFloat = uint

// See: https://developer.apple.com/documentation/kernel/nxtabletpointdata
type NXTabletPointData = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nxtabletpointdataptr
type NXTabletPointDataPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nxtabletproximitydata
type NXTabletProximityData = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nxtabletproximitydataptr
type NXTabletProximityDataPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nudclflags
type NuDCLFlags = uint

// See: https://developer.apple.com/documentation/kernel/nudclreceivepacketref
type NuDCLReceivePacketRef = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nudclref
type NuDCLRef uintptr

// See: https://developer.apple.com/documentation/kernel/nudclsendpacketref
type NuDCLSendPacketRef = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nudclskipcycleref
type NuDCLSkipCycleRef = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/numversion
type NumVersion = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/osactionabortedhandler
type OSActionAbortedHandler = func()

// See: https://developer.apple.com/documentation/kernel/osactioncancelhandler
type OSActionCancelHandler = func()

// See: https://developer.apple.com/documentation/kernel/osarrayptr
type OSArrayPtr = OSArray

// See: https://developer.apple.com/documentation/kernel/osasyncreference64
type OSAsyncReference64 = Io_user_reference_t

// See: https://developer.apple.com/documentation/kernel/osasyncreference
type OSAsyncReference = Natural_t

// See: https://developer.apple.com/documentation/kernel/osbooleanptr
type OSBooleanPtr = OSBoolean

// See: https://developer.apple.com/documentation/kernel/oscollectioniteratorptr
type OSCollectionIteratorPtr = OSCollectionIterator

// See: https://developer.apple.com/documentation/kernel/oscollectionptr
type OSCollectionPtr = OSCollection

// See: https://developer.apple.com/documentation/kernel/oscontainer
type OSContainer = OSObject

// See: https://developer.apple.com/documentation/kernel/osdataconstptr
type OSDataConstPtr = OSData

// See: https://developer.apple.com/documentation/kernel/osdataptr
type OSDataPtr = OSData

// See: https://developer.apple.com/documentation/kernel/osdictionaryptr
type OSDictionaryPtr = OSDictionary

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
type OSNumberPtr = OSNumber

// See: https://developer.apple.com/documentation/kernel/osobjectapplierblock
type OSObjectApplierBlock = func(*OSObject)

// See: https://developer.apple.com/documentation/kernel/osobjectptr
type OSObjectPtr = OSObject

// See: https://developer.apple.com/documentation/kernel/osobjectref
type OSObjectRef = uint64

// See: https://developer.apple.com/documentation/kernel/osorderedsetptr
type OSOrderedSetPtr = OSOrderedSet

// See: https://developer.apple.com/documentation/kernel/osserializeptr
type OSSerializePtr = OSSerialize

// See: https://developer.apple.com/documentation/kernel/osserializerblock
type OSSerializerBlock = func(*OSSerialize) bool

// See: https://developer.apple.com/documentation/kernel/osserializerptr
type OSSerializerPtr = OSSerializer

// See: https://developer.apple.com/documentation/kernel/ossetptr
type OSSetPtr = OSSet

// See: https://developer.apple.com/documentation/kernel/osstringconstptr
type OSStringConstPtr = OSString

// See: https://developer.apple.com/documentation/kernel/osstringptr
type OSStringPtr = OSString

// See: https://developer.apple.com/documentation/kernel/ossymbolconstptr
type OSSymbolConstPtr = OSSymbol

// See: https://developer.apple.com/documentation/kernel/ossymbolptr
type OSSymbolPtr = OSSymbol

// See: https://developer.apple.com/documentation/kernel/ostype
type OSType = uint32

// See: https://developer.apple.com/documentation/kernel/ostypeptr
type OSTypePtr = uint32

// See: https://developer.apple.com/documentation/kernel/opaquedtentryiterator
type OpaqueDTEntryIterator = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/opaquedtpropertyiterator
type OpaqueDTPropertyIterator = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/pbversion
type PBVersion = uint32

// See: https://developer.apple.com/documentation/kernel/pe_video
type PE_Video = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/pe_state_t
type PE_state_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ptr
type Ptr = int8

// See: https://developer.apple.com/documentation/kernel/raw_header
type RAW_header = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/report_luns_logical_unit_addressing
type REPORT_LUNS_LOGICAL_UNIT_ADDRESSING = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/report_luns_peripheral_device_addressing
type REPORT_LUNS_PERIPHERAL_DEVICE_ADDRESSING = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/rgbcolor
type RGBColor = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/rgbcolorhdl
type RGBColorHdl = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/rgbcolorptr
type RGBColorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/rawsensecode
type RawSenseCode = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/realdtentry
type RealDTEntry = DeviceTreeNode

// RectPtr is represents a type used by the Video Components API.
//
// See: https://developer.apple.com/documentation/kernel/rectptr
type RectPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/regcstrentryname
type RegCStrEntryName = int8

// See: https://developer.apple.com/documentation/kernel/regcstrentrynamebuf
type RegCStrEntryNameBuf = int8

// See: https://developer.apple.com/documentation/kernel/regcstrentrynameptr
type RegCStrEntryNamePtr = int8

// See: https://developer.apple.com/documentation/kernel/regcstrpathname
type RegCStrPathName = int8

// See: https://developer.apple.com/documentation/kernel/regentryid
type RegEntryID = string

// See: https://developer.apple.com/documentation/kernel/regentryidptr
type RegEntryIDPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/regentryiter
type RegEntryIter = IORegistryIterator

// See: https://developer.apple.com/documentation/kernel/regentryiterationop
type RegEntryIterationOp = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/regentrymodifiers
type RegEntryModifiers = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/regiterationop
type RegIterationOp = uint32

// See: https://developer.apple.com/documentation/kernel/regmodifiers
type RegModifiers = uint32

// See: https://developer.apple.com/documentation/kernel/regpathnamesize
type RegPathNameSize = uint32

// See: https://developer.apple.com/documentation/kernel/regpropertyiter
type RegPropertyIter = OSIterator

// See: https://developer.apple.com/documentation/kernel/regpropertymodifiers
type RegPropertyModifiers = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/regpropertyname
type RegPropertyName = int8

// See: https://developer.apple.com/documentation/kernel/regpropertynamebuf
type RegPropertyNameBuf = int8

// See: https://developer.apple.com/documentation/kernel/regpropertynameptr
type RegPropertyNamePtr = int8

// See: https://developer.apple.com/documentation/kernel/regpropertyvalue
type RegPropertyValue = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/regpropertyvaluesize
type RegPropertyValueSize = uint32

// See: https://developer.apple.com/documentation/kernel/restype
type ResType = uint32

// See: https://developer.apple.com/documentation/kernel/restypeptr
type ResTypePtr = uint32

// See: https://developer.apple.com/documentation/kernel/runtimeoptions
type RuntimeOptions = uint32

// See: https://developer.apple.com/documentation/kernel/sbcmodepagecaching
type SBCModePageCaching = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/sbcmodepageflexibledisk
type SBCModePageFlexibleDisk = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/sbcmodepageformatdevice
type SBCModePageFormatDevice = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/sbcmodepagerigiddiskgeometry
type SBCModePageRigidDiskGeometry = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield10bit
type SCSICmdField10Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield11bit
type SCSICmdField11Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield12bit
type SCSICmdField12Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield13bit
type SCSICmdField13Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield14bit
type SCSICmdField14Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield15bit
type SCSICmdField15Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield17bit
type SCSICmdField17Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield18bit
type SCSICmdField18Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield19bit
type SCSICmdField19Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield1bit
type SCSICmdField1Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield1byte
type SCSICmdField1Byte = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield20bit
type SCSICmdField20Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield21bit
type SCSICmdField21Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield22bit
type SCSICmdField22Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield23bit
type SCSICmdField23Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield25bit
type SCSICmdField25Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield26bit
type SCSICmdField26Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield27bit
type SCSICmdField27Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield28bit
type SCSICmdField28Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield29bit
type SCSICmdField29Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield2bit
type SCSICmdField2Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield2byte
type SCSICmdField2Byte = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield30bit
type SCSICmdField30Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield31bit
type SCSICmdField31Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield33bit
type SCSICmdField33Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield34bit
type SCSICmdField34Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield35bit
type SCSICmdField35Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield36bit
type SCSICmdField36Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield37bit
type SCSICmdField37Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield38bit
type SCSICmdField38Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield39bit
type SCSICmdField39Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield3bit
type SCSICmdField3Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield3byte
type SCSICmdField3Byte = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield41bit
type SCSICmdField41Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield42bit
type SCSICmdField42Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield43bit
type SCSICmdField43Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield44bit
type SCSICmdField44Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield45bit
type SCSICmdField45Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield46bit
type SCSICmdField46Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield47bit
type SCSICmdField47Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield49bit
type SCSICmdField49Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield4bit
type SCSICmdField4Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield4byte
type SCSICmdField4Byte = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield50bit
type SCSICmdField50Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield51bit
type SCSICmdField51Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield52bit
type SCSICmdField52Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield53bit
type SCSICmdField53Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield54bit
type SCSICmdField54Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield55bit
type SCSICmdField55Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield57bit
type SCSICmdField57Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield58bit
type SCSICmdField58Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield59bit
type SCSICmdField59Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield5bit
type SCSICmdField5Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield5byte
type SCSICmdField5Byte = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield60bit
type SCSICmdField60Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield61bit
type SCSICmdField61Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield62bit
type SCSICmdField62Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield63bit
type SCSICmdField63Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield6bit
type SCSICmdField6Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield6byte
type SCSICmdField6Byte = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield7bit
type SCSICmdField7Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield7byte
type SCSICmdField7Byte = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield8byte
type SCSICmdField8Byte = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmdfield9bit
type SCSICmdField9Bit = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmd_inquiry_pagecx_header
type SCSICmd_INQUIRY_PAGECx_Header = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmd_inquiry_page00_header_spc_16
type SCSICmd_INQUIRY_Page00_Header_SPC_16 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmd_inquiry_page80_header_spc_16
type SCSICmd_INQUIRY_Page80_Header_SPC_16 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmd_inquiry_pageb0_data
type SCSICmd_INQUIRY_PageB0_Data = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmd_inquiry_pageb2_data
type SCSICmd_INQUIRY_PageB2_Data = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmd_inquiry_pageb2_provisioning_group_descriptor
type SCSICmd_INQUIRY_PageB2_Provisioning_Group_Descriptor = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmd_inquiry_pagec0_data
type SCSICmd_INQUIRY_PageC0_Data = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmd_inquiry_pagec1_data
type SCSICmd_INQUIRY_PageC1_Data = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmd_inquiry_standarddataptr
type SCSICmd_INQUIRY_StandardDataPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmd_report_luns_header
type SCSICmd_REPORT_LUNS_Header = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicmd_report_luns_lun_entry
type SCSICmd_REPORT_LUNS_LUN_ENTRY = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsicommanddescriptorblock
type SCSICommandDescriptorBlock = unsafe.Pointer

// SCSIDeviceIdentifier is 64-bit number to represent a SCSI Device.
//
// See: https://developer.apple.com/documentation/kernel/scsideviceidentifier
type SCSIDeviceIdentifier = uint64

// SCSIInitiatorIdentifier is 64-bit number to represent a SCSI Initiator Device.
//
// See: https://developer.apple.com/documentation/kernel/scsiinitiatoridentifier
type SCSIInitiatorIdentifier = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsilogicalunitbytes
type SCSILogicalUnitBytes = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsilogicalunitnumber
type SCSILogicalUnitNumber = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsiparalleltaskidentifier
type SCSIParallelTaskIdentifier = OSObject

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
type SCSIServiceResponse = unsafe.Pointer

// SCSITaggedTaskIdentifier is 64-bit number to represent a unique task identifier.
//
// See: https://developer.apple.com/documentation/kernel/scsitaggedtaskidentifier
type SCSITaggedTaskIdentifier = unsafe.Pointer

// SCSITargetIdentifier is 64-bit number to represent a SCSI Target Device.
//
// See: https://developer.apple.com/documentation/kernel/scsitargetidentifier
type SCSITargetIdentifier = unsafe.Pointer

// SCSITaskAttribute is attributes for task delivery.
//
// See: https://developer.apple.com/documentation/kernel/scsitaskattribute
type SCSITaskAttribute = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsitaskidentifier
type SCSITaskIdentifier = OSObject

// See: https://developer.apple.com/documentation/kernel/scsitaskmode
type SCSITaskMode = uint

// SCSITaskState is attributes for task state.
//
// See: https://developer.apple.com/documentation/kernel/scsitaskstate
type SCSITaskState = unsafe.Pointer

// SCSITaskStatus is attributes for task status.
//
// See: https://developer.apple.com/documentation/kernel/scsitaskstatus
type SCSITaskStatus = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/scsi_sense_data
type SCSI_Sense_Data = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/sc_scatter
type SC_Scatter = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/sha1_ctx
type SHA1_CTX = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/sint
type SInt = int

// See: https://developer.apple.com/documentation/kernel/sint16
type SInt16 = int16

// See: https://developer.apple.com/documentation/kernel/sint32
type SInt32 = int32

// See: https://developer.apple.com/documentation/kernel/sint64
type SInt64 = int64

// See: https://developer.apple.com/documentation/kernel/sint8
type SInt8 = int8

// See: https://developer.apple.com/documentation/kernel/spcmodepagepowercondition
type SPCModePagePowerCondition = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/spcmodeparameterheader10
type SPCModeParameterHeader10 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/spcmodeparameterheader6
type SPCModeParameterHeader6 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/servicecount
type ServiceCount = uint32

// See: https://developer.apple.com/documentation/kernel/signedbyte
type SignedByte = int8

// See: https://developer.apple.com/documentation/kernel/stickykeys_modifierinfo
type StickyKeys_ModifierInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/stickykeys_toggleinfo
type StickyKeys_ToggleInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/str31
type Str31 = uint8

// See: https://developer.apple.com/documentation/kernel/transmissionpower
type TransmissionPower = int8

// See: https://developer.apple.com/documentation/kernel/uaspipedescriptor
type UASPipeDescriptor = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/uaspipedescriptorptr
type UASPipeDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/uint16
type UInt16 = uint16

// See: https://developer.apple.com/documentation/kernel/uint32
type UInt32 = uint32

// See: https://developer.apple.com/documentation/kernel/uint32ptr
type UInt32Ptr = uint32

// See: https://developer.apple.com/documentation/kernel/uint64
type UInt64 = uint64

// See: https://developer.apple.com/documentation/kernel/uint8
type UInt8 = uint8

// See: https://developer.apple.com/documentation/kernel/undkey
type UNDKey = int8

// See: https://developer.apple.com/documentation/kernel/undlabel
type UNDLabel = int8

// See: https://developer.apple.com/documentation/kernel/undmessage
type UNDMessage = int8

// See: https://developer.apple.com/documentation/kernel/undpath
type UNDPath = int8

// See: https://developer.apple.com/documentation/kernel/undreplyref
type UNDReplyRef = uint32

// See: https://developer.apple.com/documentation/kernel/undserverref
type UNDServerRef = uint32

// USBDeviceAddress is a USB device address.
//
// See: https://developer.apple.com/documentation/kernel/usbdeviceaddress
type USBDeviceAddress = unsafe.Pointer

// USBDeviceInformationBits is the state of a USB device.
//
// See: https://developer.apple.com/documentation/kernel/usbdeviceinformationbits
type USBDeviceInformationBits = int

// USBLowLatencyBufferType is specifies which kind of low-latency buffer to create.
//
// See: https://developer.apple.com/documentation/kernel/usblowlatencybuffertype
type USBLowLatencyBufferType = int

// USBNotificationTypes is defines types of USB notifications.
//
// See: https://developer.apple.com/documentation/kernel/usbnotificationtypes
type USBNotificationTypes = int

// USBPhysicalAddress32 is a 32-bit USB physical address.
//
// See: https://developer.apple.com/documentation/kernel/usbphysicaladdress32
type USBPhysicalAddress32 = uint32

// USBPowerRequestTypes is specifies the kind of power to reserve.
//
// See: https://developer.apple.com/documentation/kernel/usbpowerrequesttypes
type USBPowerRequestTypes = int

// USBReEnumerateOptions is options for reenumerating a device.
//
// See: https://developer.apple.com/documentation/kernel/usbreenumerateoptions
type USBReEnumerateOptions = uint

// USBStatus is the value of the USB device status.
//
// See: https://developer.apple.com/documentation/kernel/usbstatus
type USBStatus = uint16

// USBStatusPtr is a pointer to a USB status.
//
// See: https://developer.apple.com/documentation/kernel/usbstatusptr
type USBStatusPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/unichar
type UniChar = uint16

// See: https://developer.apple.com/documentation/kernel/userexportdclcallcommandproc
type UserExportDCLCallCommandProc = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/userexportdclcallproc
type UserExportDCLCallProc = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/userexportdclcommand
type UserExportDCLCommand = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/userexportdcljump
type UserExportDCLJump = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/userexportdcllabel
type UserExportDCLLabel = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/userexportdclnudclleader
type UserExportDCLNuDCLLeader = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/userexportdclptrtimestamp
type UserExportDCLPtrTimeStamp = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/userexportdclsettagsyncbits
type UserExportDCLSetTagSyncBits = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/userexportdcltimestamp
type UserExportDCLTimeStamp = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/userexportdcltransferbuffer
type UserExportDCLTransferBuffer = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/userexportdcltransferpacket
type UserExportDCLTransferPacket = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/userexportdclupdatedcllist
type UserExportDCLUpdateDCLList = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdclutbehavior
type VDClutBehavior = uint32

// See: https://developer.apple.com/documentation/kernel/vdclutbehaviorptr
type VDClutBehaviorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdcommunicationinfoptr
type VDCommunicationInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdcommunicationinforec
type VDCommunicationInfoRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdcommunicationptr
type VDCommunicationPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdcommunicationrec
type VDCommunicationRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdconfigurationfeaturelistrec
type VDConfigurationFeatureListRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdconfigurationfeaturelistrecptr
type VDConfigurationFeatureListRecPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdconfigurationptr
type VDConfigurationPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdconfigurationrec
type VDConfigurationRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdconvolutioninfoptr
type VDConvolutionInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdconvolutioninforec
type VDConvolutionInfoRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdddcblockptr
type VDDDCBlockPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdddcblockrec
type VDDDCBlockRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vddefmode
type VDDefMode = uint

// See: https://developer.apple.com/documentation/kernel/vddefmodeptr
type VDDefModePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vddetailedtimingptr
type VDDetailedTimingPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vddetailedtimingrec
type VDDetailedTimingRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vddisplayconnectinfoptr
type VDDisplayConnectInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vddisplayconnectinforec
type VDDisplayConnectInfoRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vddisplaytimingrangeptr
type VDDisplayTimingRangePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vddisplaytimingrangerec
type VDDisplayTimingRangeRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vddrawhardwarecursorptr
type VDDrawHardwareCursorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vddrawhardwarecursorrec
type VDDrawHardwareCursorRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdentrecptr
type VDEntRecPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdentryrecord
type VDEntryRecord = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdflagrecptr
type VDFlagRecPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdflagrecord
type VDFlagRecord = unsafe.Pointer

// VDGamRecPtr is represents a type used by the Video Components API.
//
// See: https://developer.apple.com/documentation/kernel/vdgamrecptr
type VDGamRecPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdgammainfoptr
type VDGammaInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdgammainforec
type VDGammaInfoRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdgammarecord
type VDGammaRecord = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdgetgammalistptr
type VDGetGammaListPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdgetgammalistrec
type VDGetGammaListRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdgrayptr
type VDGrayPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdgrayrecord
type VDGrayRecord = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdhardwarecursordrawstateptr
type VDHardwareCursorDrawStatePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdhardwarecursordrawstaterec
type VDHardwareCursorDrawStateRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdmirrorptr
type VDMirrorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdmirrorrec
type VDMirrorRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdmulticonnectinfoptr
type VDMultiConnectInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdmulticonnectinforec
type VDMultiConnectInfoRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdpageinfo
type VDPageInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdpginfoptr
type VDPgInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdpowerstateptr
type VDPowerStatePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdpowerstaterec
type VDPowerStateRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdprivateselectordatarec
type VDPrivateSelectorDataRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdprivateselectorrec
type VDPrivateSelectorRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdresolutioninfoptr
type VDResolutionInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdresolutioninforec
type VDResolutionInfoRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdretrievegammaptr
type VDRetrieveGammaPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdretrievegammarec
type VDRetrieveGammaRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdscalerinfoptr
type VDScalerInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdscalerinforec
type VDScalerInfoRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdscalerptr
type VDScalerPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdscalerrec
type VDScalerRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdsetentryptr
type VDSetEntryPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdsetentryrecord
type VDSetEntryRecord = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdsethardwarecursorptr
type VDSetHardwareCursorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdsethardwarecursorrec
type VDSetHardwareCursorRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdsettings
type VDSettings = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdsettingsptr
type VDSettingsPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdsizeinfo
type VDSizeInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdsupportshardwarecursorptr
type VDSupportsHardwareCursorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdsupportshardwarecursorrec
type VDSupportsHardwareCursorRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdswitchinfoptr
type VDSwitchInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdswitchinforec
type VDSwitchInfoRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdsyncinfoptr
type VDSyncInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdsyncinforec
type VDSyncInfoRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdszinfoptr
type VDSzInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdtiminginfoptr
type VDTimingInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdtiminginforec
type VDTimingInfoRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdvideoparametersinfoptr
type VDVideoParametersInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdvideoparametersinforec
type VDVideoParametersInfoRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/void
type VOID = string

// See: https://developer.apple.com/documentation/kernel/vpblock
type VPBlock = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vpblockptr
type VPBlockPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/videodevicetype
type VideoDeviceType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/wk_word
type WK_word = uint

// See: https://developer.apple.com/documentation/kernel/addr64_t
type Addr64_t = uint64

// See: https://developer.apple.com/documentation/kernel/aid_t
type Aid_t = uint64

// See: https://developer.apple.com/documentation/kernel/alarm_port_t
type Alarm_port_t = Alarm_t

// See: https://developer.apple.com/documentation/kernel/alarm_t
type Alarm_t = uintptr

// See: https://developer.apple.com/documentation/kernel/alarm_type_t
type Alarm_type_t = int

// See: https://developer.apple.com/documentation/kernel/arcade_register_t
type Arcade_register_t = uintptr

// See: https://developer.apple.com/documentation/kernel/arm_debug_info_t
type Arm_debug_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/arm_exception_state32_t
type Arm_exception_state32_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/arm_feature_bits_t
type Arm_feature_bits_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/arm_neon_state32_t
type Arm_neon_state32_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/arm_state_hdr_t
type Arm_state_hdr_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/arm_thread_state32_t
type Arm_thread_state32_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/arm_unified_thread_state_t
type Arm_unified_thread_state_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ataregisterimage
type AtaRegisterImage = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/atataskfile
type AtaTaskFile = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/atm_action_t
type Atm_action_t = uint32

// See: https://developer.apple.com/documentation/kernel/atm_aid_t
type Atm_aid_t = uint64

// See: https://developer.apple.com/documentation/kernel/atm_guard_t
type Atm_guard_t = uint64

// See: https://developer.apple.com/documentation/kernel/atm_mailbox_offset_t
type Atm_mailbox_offset_t = uint64

// See: https://developer.apple.com/documentation/kernel/atm_memory_descriptor_array_t
type Atm_memory_descriptor_array_t = Atm_memory_descriptor_t

// See: https://developer.apple.com/documentation/kernel/atm_memory_descriptor_t
type Atm_memory_descriptor_t = uint32

// See: https://developer.apple.com/documentation/kernel/atm_memory_size_array_t
type Atm_memory_size_array_t = uint64

// See: https://developer.apple.com/documentation/kernel/atm_subaid32_t
type Atm_subaid32_t = uint32

// See: https://developer.apple.com/documentation/kernel/attrgroup_t
type Attrgroup_t = U_int32_t

// See: https://developer.apple.com/documentation/kernel/attribute_set_t
type Attribute_set_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/attrreference_t
type Attrreference_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/au_asflgs_t
type Au_asflgs_t = U_int64_t

// See: https://developer.apple.com/documentation/kernel/au_asid_t
type Au_asid_t = int32

// See: https://developer.apple.com/documentation/kernel/au_class_t
type Au_class_t = U_int32_t

// See: https://developer.apple.com/documentation/kernel/au_ctlmode_t
type Au_ctlmode_t = uint8

// See: https://developer.apple.com/documentation/kernel/au_emod_t
type Au_emod_t = U_int16_t

// See: https://developer.apple.com/documentation/kernel/au_evclass_map_t
type Au_evclass_map_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/au_event_t
type Au_event_t = U_int16_t

// See: https://developer.apple.com/documentation/kernel/au_expire_after_t
type Au_expire_after_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/au_fstat_t
type Au_fstat_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/au_id_t
type Au_id_t = uint32

// See: https://developer.apple.com/documentation/kernel/au_mask_t
type Au_mask_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/au_qctrl_t
type Au_qctrl_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/au_session_t
type Au_session_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/au_stat_t
type Au_stat_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/au_tid_addr_t
type Au_tid_addr_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/au_tid_t
type Au_tid_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/auditinfo_addr_t
type Auditinfo_addr_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/auditinfo_t
type Auditinfo_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/auditpinfo_addr_t
type Auditpinfo_addr_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/auditpinfo_t
type Auditpinfo_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/backtrace_flags_t
type Backtrace_flags_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/backtrace_info_t
type Backtrace_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/backtrace_pack_t
type Backtrace_pack_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bank_action_t
type Bank_action_t = uint32

// See: https://developer.apple.com/documentation/kernel/blkcnt_t
type Blkcnt_t = int64

// See: https://developer.apple.com/documentation/kernel/blksize_t
type Blksize_t = int32

// See: https://developer.apple.com/documentation/kernel/block_hint_t
type Block_hint_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/boolean_t
type Boolean_t = bool

// See: https://developer.apple.com/documentation/kernel/boot_args
type Boot_args = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/boot_icon_element
type Boot_icon_element = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/bootstrap_t
type Bootstrap_t = uint32

// See: https://developer.apple.com/documentation/kernel/bpf_int32
type Bpf_int32 = int32

// Bpf_tap_mode is mode for tapping. BPF_MODE_DISABLED/BPF_MODE_INPUT_OUTPUT etc.
//
// See: https://developer.apple.com/documentation/kernel/bpf_tap_mode
type Bpf_tap_mode = U_int32_t

// See: https://developer.apple.com/documentation/kernel/bpf_u_int32
type Bpf_u_int32 = U_int32_t

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
type Cache_type_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/caddr_t
type Caddr_t = int8

// See: https://developer.apple.com/documentation/kernel/caddr_ut
type Caddr_ut = Caddr_t

// See: https://developer.apple.com/documentation/kernel/call_gate_t
type Call_gate_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cc_t
type Cc_t = uint8

// See: https://developer.apple.com/documentation/kernel/charf
type Charf = int8

// See: https://developer.apple.com/documentation/kernel/circle_queue_head_t
type Circle_queue_head_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/circle_queue_t
type Circle_queue_t = uintptr

// See: https://developer.apple.com/documentation/kernel/cl_direct_read_lock_t
type Cl_direct_read_lock_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/clock_attr_t
type Clock_attr_t = int

// See: https://developer.apple.com/documentation/kernel/clock_ctrl_port_t
type Clock_ctrl_port_t = Clock_ctrl_t

// See: https://developer.apple.com/documentation/kernel/clock_ctrl_t
type Clock_ctrl_t = uintptr

// See: https://developer.apple.com/documentation/kernel/clock_flavor_t
type Clock_flavor_t = int

// See: https://developer.apple.com/documentation/kernel/clock_frequency_info_t
type Clock_frequency_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/clock_id_t
type Clock_id_t = int

// See: https://developer.apple.com/documentation/kernel/clock_nsec_t
type Clock_nsec_t = uint

// See: https://developer.apple.com/documentation/kernel/clock_reply_t
type Clock_reply_t = uint32

// See: https://developer.apple.com/documentation/kernel/clock_res_t
type Clock_res_t = int

// See: https://developer.apple.com/documentation/kernel/clock_sec_t
type Clock_sec_t = uint

// See: https://developer.apple.com/documentation/kernel/clock_serv_port_t
type Clock_serv_port_t = Clock_serv_t

// See: https://developer.apple.com/documentation/kernel/clock_serv_t
type Clock_serv_t = uintptr

// See: https://developer.apple.com/documentation/kernel/clock_t
type Clock_t = uint

// See: https://developer.apple.com/documentation/kernel/clock_usec_t
type Clock_usec_t = uint

// See: https://developer.apple.com/documentation/kernel/cluster_type_t
type Cluster_type_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/coalition_t
type Coalition_t = uintptr

// See: https://developer.apple.com/documentation/kernel/code_desc_t
type Code_desc_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/conninfo_multipathtcp_t
type Conninfo_multipathtcp_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/conninfo_tcp_t
type Conninfo_tcp_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/coprocessor_type_t
type Coprocessor_type_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cpu_id_t
type Cpu_id_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cpu_subtype_t
type Cpu_subtype_t = int32

// See: https://developer.apple.com/documentation/kernel/cpu_threadtype_t
type Cpu_threadtype_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/cpu_type_t
type Cpu_type_t = int32

// See: https://developer.apple.com/documentation/kernel/cpuid_arch_perf_leaf_t
type Cpuid_arch_perf_leaf_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cpuid_cache_desc_t
type Cpuid_cache_desc_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cpuid_mwait_leaf_t
type Cpuid_mwait_leaf_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cpuid_register_t
type Cpuid_register_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cpuid_thermal_leaf_t
type Cpuid_thermal_leaf_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cpuid_tsc_leaf_t
type Cpuid_tsc_leaf_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cpuid_xsave_leaf_t
type Cpuid_xsave_leaf_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cr0_t
type Cr0_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/cryptex_auth_type_t
type Cryptex_auth_type_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/crypto_random_ctx_t
type Crypto_random_ctx_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/crypto_random_kmem_ctx_size_fn_t
type Crypto_random_kmem_ctx_size_fn_t = uintptr

// See: https://developer.apple.com/documentation/kernel/cs_launch_type_t
type Cs_launch_type_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ct_rune_t
type Ct_rune_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/d_devtotty_t
type D_devtotty_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/daddr64_t
type Daddr64_t = int64

// See: https://developer.apple.com/documentation/kernel/daddr_t
type Daddr_t = int32

// See: https://developer.apple.com/documentation/kernel/data_desc_t
type Data_desc_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/debug_header_entry
type Debug_header_entry = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/debug_header_t
type Debug_header_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/debug_trailer_t
type Debug_trailer_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/descriptor_options
type Descriptor_options = uint32

// See: https://developer.apple.com/documentation/kernel/dev_t
type Dev_t = int32

// See: https://developer.apple.com/documentation/kernel/dir_clone_authorizer_op_t
type Dir_clone_authorizer_op_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_bd_read_disc_info_t
type Dk_bd_read_disc_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_bd_read_structure_t
type Dk_bd_read_structure_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_bd_read_track_info_t
type Dk_bd_read_track_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_bd_report_key_t
type Dk_bd_report_key_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_bd_send_key_t
type Dk_bd_send_key_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_cd_read_disc_info_t
type Dk_cd_read_disc_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_cd_read_isrc_t
type Dk_cd_read_isrc_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_cd_read_mcn_t
type Dk_cd_read_mcn_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_cd_read_t
type Dk_cd_read_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_cd_read_toc_t
type Dk_cd_read_toc_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_cd_read_track_info_t
type Dk_cd_read_track_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_corestorage_info_t
type Dk_corestorage_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_dvd_read_disc_info_t
type Dk_dvd_read_disc_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_dvd_read_rzone_info_t
type Dk_dvd_read_rzone_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_dvd_read_structure_t
type Dk_dvd_read_structure_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_dvd_report_key_t
type Dk_dvd_report_key_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_dvd_send_key_t
type Dk_dvd_send_key_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_error_description_t
type Dk_error_description_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_extent_t
type Dk_extent_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_firmware_path_t
type Dk_firmware_path_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_format_capacities_t
type Dk_format_capacities_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_format_capacity_t
type Dk_format_capacity_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_physical_extent_t
type Dk_physical_extent_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_provision_extent_t
type Dk_provision_extent_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_provision_status_t
type Dk_provision_status_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_set_tier_t
type Dk_set_tier_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_synchronize_t
type Dk_synchronize_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dk_unmap_t
type Dk_unmap_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/double_t
type Double_t = float64

// See: https://developer.apple.com/documentation/kernel/dump_fcn_t
type Dump_fcn_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dyld_kernel_image_info_array_t
type Dyld_kernel_image_info_array_t = Dyld_kernel_image_info_t

// See: https://developer.apple.com/documentation/kernel/dyld_kernel_image_info_t
type Dyld_kernel_image_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/dyld_kernel_process_info_t
type Dyld_kernel_process_info_t = unsafe.Pointer

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
type Ecc_event_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ecc_flags_t
type Ecc_flags_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ecc_version_t
type Ecc_version_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/empty_fcn_t
type Empty_fcn_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/emulation_vector_t
type Emulation_vector_t = Mach_vm_offset_t

// See: https://developer.apple.com/documentation/kernel/eph_panic_flags_t
type Eph_panic_flags_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/er_t
type Er_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/errno_t
type Errno_t = int

// See: https://developer.apple.com/documentation/kernel/ether_header_t
type Ether_header_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/event64_t
type Event64_t = uint64

// See: https://developer.apple.com/documentation/kernel/event_t
type Event_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/eventlink_port_pair_t
type Eventlink_port_pair_t = uintptr

// See: https://developer.apple.com/documentation/kernel/eviospecialkeymsg_t
type EvioSpecialKeyMsg_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/evsioevsioccsindices
type EvsioEVSIOCCSIndices = int

// See: https://developer.apple.com/documentation/kernel/evsioevsioscsindices
type EvsioEVSIOSCSIndices = int

// See: https://developer.apple.com/documentation/kernel/ex_cb_action_t
type Ex_cb_action_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ex_cb_class_t
type Ex_cb_class_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ex_cb_state_t
type Ex_cb_state_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/exception_behavior_array_t
type Exception_behavior_array_t = Exception_behavior_t

// See: https://developer.apple.com/documentation/kernel/exception_behavior_t
type Exception_behavior_t = int

// See: https://developer.apple.com/documentation/kernel/exception_data_t
type Exception_data_t = Exception_data_type_t

// See: https://developer.apple.com/documentation/kernel/exception_data_type_t
type Exception_data_type_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/exception_flavor_array_t
type Exception_flavor_array_t = Thread_state_flavor_t

// See: https://developer.apple.com/documentation/kernel/exception_handler_array_t
type Exception_handler_array_t = Exception_handler_t

// See: https://developer.apple.com/documentation/kernel/exception_handler_info_array_t
type Exception_handler_info_array_t = Ipc_info_port_t

// See: https://developer.apple.com/documentation/kernel/exception_handler_info_t
type Exception_handler_info_t = Ipc_info_port_t

// See: https://developer.apple.com/documentation/kernel/exception_handler_t
type Exception_handler_t = uint32

// See: https://developer.apple.com/documentation/kernel/exception_mask_array_t
type Exception_mask_array_t = Exception_mask_t

// See: https://developer.apple.com/documentation/kernel/exception_mask_t
type Exception_mask_t = uint

// See: https://developer.apple.com/documentation/kernel/exception_port_arrary_t
type Exception_port_arrary_t = Exception_handler_array_t

// See: https://developer.apple.com/documentation/kernel/exception_port_array_t
type Exception_port_array_t = uint32

// See: https://developer.apple.com/documentation/kernel/exception_port_info_array_t
type Exception_port_info_array_t = Ipc_info_port_t

// See: https://developer.apple.com/documentation/kernel/exception_port_t
type Exception_port_t = Exception_handler_t

// See: https://developer.apple.com/documentation/kernel/exception_type_t
type Exception_type_t = int

// See: https://developer.apple.com/documentation/kernel/exclave_ecstackentry_addr_t
type Exclave_ecstackentry_addr_t = uint64

// See: https://developer.apple.com/documentation/kernel/ext_paniclog_create_options_t
type Ext_paniclog_create_options_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ext_paniclog_flags_t
type Ext_paniclog_flags_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/extentrecord
type Extentrecord = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fattributiontag_t
type Fattributiontag_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fchecklv_t
type Fchecklv_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fd_mask
type Fd_mask = int32

// See: https://developer.apple.com/documentation/kernel/fd_set
type Fd_set = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fgetsigsinfo_t
type Fgetsigsinfo_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fhandle_t
type Fhandle_t = unsafe.Pointer

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
type Fixpt_t = U_int32_t

// See: https://developer.apple.com/documentation/kernel/float_t
type Float_t = float32

// See: https://developer.apple.com/documentation/kernel/fp_control_t
type Fp_control_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fp_status_t
type Fp_status_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fpunchhole_t
type Fpunchhole_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/frame_type_bitmask_t
type Frame_type_bitmask_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fs_role_mount_args_t
type Fs_role_mount_args_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fsblkcnt_t
type Fsblkcnt_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fsfilcnt_t
type Fsfilcnt_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fsfile_type_t
type Fsfile_type_t = U_int32_t

// See: https://developer.apple.com/documentation/kernel/fsid_t
type Fsid_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fsignatures_t
type Fsignatures_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fsobj_id_t
type Fsobj_id_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fsobj_tag_t
type Fsobj_tag_t = U_int32_t

// See: https://developer.apple.com/documentation/kernel/fsobj_type_t
type Fsobj_type_t = U_int32_t

// See: https://developer.apple.com/documentation/kernel/fspecread_t
type Fspecread_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fstore_t
type Fstore_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fsupplement_t
type Fsupplement_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/fsvolid_t
type Fsvolid_t = U_int32_t

// See: https://developer.apple.com/documentation/kernel/ftrimactivefile_t
type Ftrimactivefile_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/gdt_t
type Gdt_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/gid_t
type Gid_t = uint32

// See: https://developer.apple.com/documentation/kernel/gpu_descriptor
type Gpu_descriptor = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/gpu_descriptor_t
type Gpu_descriptor_t = Gpu_descriptor

// See: https://developer.apple.com/documentation/kernel/gpu_energy_data
type Gpu_energy_data = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/gpu_energy_data_t
type Gpu_energy_data_t = Gpu_energy_data

// See: https://developer.apple.com/documentation/kernel/graftdmg_type_t
type Graftdmg_type_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/gssd_byte_buffer
type Gssd_byte_buffer = uint8

// See: https://developer.apple.com/documentation/kernel/gssd_cred
type Gssd_cred = uint64

// See: https://developer.apple.com/documentation/kernel/gssd_ctx
type Gssd_ctx = uint64

// See: https://developer.apple.com/documentation/kernel/gssd_dstring
type Gssd_dstring = int8

// See: https://developer.apple.com/documentation/kernel/gssd_etype_list
type Gssd_etype_list = int32

// See: https://developer.apple.com/documentation/kernel/gssd_gid_list
type Gssd_gid_list = uint32

// See: https://developer.apple.com/documentation/kernel/gssd_mechtype
type Gssd_mechtype = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/gssd_nametype
type Gssd_nametype = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/gssd_string
type Gssd_string = int8

// See: https://developer.apple.com/documentation/kernel/gz_header
type Gz_header = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/gz_headerp
type Gz_headerp = Gz_header

// See: https://developer.apple.com/documentation/kernel/hash_info_bucket_array_t
type Hash_info_bucket_array_t = Hash_info_bucket_t

// See: https://developer.apple.com/documentation/kernel/hash_info_bucket_t
type Hash_info_bucket_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/host_basic_info_data_t
type Host_basic_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/host_basic_info_t
type Host_basic_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/host_can_has_debugger_info_data_t
type Host_can_has_debugger_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/host_can_has_debugger_info_t
type Host_can_has_debugger_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/host_cpu_load_info_data_t
type Host_cpu_load_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/host_cpu_load_info_t
type Host_cpu_load_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/host_flavor_t
type Host_flavor_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/host_info64_t
type Host_info64_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/host_info_data_t
type Host_info_data_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/host_info_t
type Host_info_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/host_load_info_data_t
type Host_load_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/host_load_info_t
type Host_load_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/host_name_port_t
type Host_name_port_t = Host_t

// See: https://developer.apple.com/documentation/kernel/host_name_t
type Host_name_t = Host_t

// See: https://developer.apple.com/documentation/kernel/host_preferred_user_arch_data_t
type Host_preferred_user_arch_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/host_preferred_user_arch_t
type Host_preferred_user_arch_t = uintptr

// See: https://developer.apple.com/documentation/kernel/host_priority_info_data_t
type Host_priority_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/host_priority_info_t
type Host_priority_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/host_priv_t
type Host_priv_t = uintptr

// See: https://developer.apple.com/documentation/kernel/host_purgable_info_data_t
type Host_purgable_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/host_purgable_info_t
type Host_purgable_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/host_sched_info_data_t
type Host_sched_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/host_sched_info_t
type Host_sched_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/host_security_t
type Host_security_t = uintptr

// See: https://developer.apple.com/documentation/kernel/host_t
type Host_t = uintptr

// See: https://developer.apple.com/documentation/kernel/hv_callbacks_t
type Hv_callbacks_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hv_trap_table_t
type Hv_trap_table_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hv_trap_type_t
type Hv_trap_type_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hv_volatile_state_t
type Hv_volatile_state_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hvg_hcall_args_t
type Hvg_hcall_args_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hvg_hcall_code_t
type Hvg_hcall_code_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hvg_hcall_dump_option_t
type Hvg_hcall_dump_option_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hvg_hcall_output_t
type Hvg_hcall_output_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hvg_hcall_return_t
type Hvg_hcall_return_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hvg_hcall_vmcore_file_t
type Hvg_hcall_vmcore_file_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/hw_spin_policy_t
type Hw_spin_policy_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/i386_cpu_info_t
type I386_cpu_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/i386_ioport_t
type I386_ioport_t = uint16

// See: https://developer.apple.com/documentation/kernel/id_t
type Id_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/idle_tickle_t
type Idle_tickle_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/idt_t
type Idt_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/idtype_t
type Idtype_t = unsafe.Pointer

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
type If_netem_model_t = unsafe.Pointer

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
type Ifnet_family_t = U_int32_t

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
type Ifnet_offload_t = U_int32_t

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
type In6_addr_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/in6_clat46_evhdlr_code_t
type In6_clat46_evhdlr_code_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/in_addr_t
type In_addr_t = uint32

// See: https://developer.apple.com/documentation/kernel/in_port_t
type In_port_t = uint16

// See: https://developer.apple.com/documentation/kernel/ino64_t
type Ino64_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ino_t
type Ino_t = uint64

// See: https://developer.apple.com/documentation/kernel/inp_gen_t
type Inp_gen_t = U_quad_t

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
type Integer_t = int

// See: https://developer.apple.com/documentation/kernel/interface_filter_t
type Interface_filter_t = uintptr

// See: https://developer.apple.com/documentation/kernel/intf
type Intf = int

// See: https://developer.apple.com/documentation/kernel/intmax_t
type Intmax_t = int

// See: https://developer.apple.com/documentation/kernel/intptr_t
type Intptr_t = int

// See: https://developer.apple.com/documentation/kernel/intr_gate_t
type Intr_gate_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/io_addr_t
type Io_addr_t = uint16

// See: https://developer.apple.com/documentation/kernel/io_async_ref64_t
type Io_async_ref64_t = Io_user_reference_t

// See: https://developer.apple.com/documentation/kernel/io_async_ref_t
type Io_async_ref_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/io_buf_ptr_t
type Io_buf_ptr_t = int8

// See: https://developer.apple.com/documentation/kernel/io_compression_stats_t
type Io_compression_stats_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/io_connect_t
type Io_connect_t = uintptr

// See: https://developer.apple.com/documentation/kernel/io_enumerator_t
type Io_enumerator_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/io_ident_t
type Io_ident_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/io_iterator_t
type Io_iterator_t = uintptr

// See: https://developer.apple.com/documentation/kernel/io_len_t
type Io_len_t = uint16

// See: https://developer.apple.com/documentation/kernel/io_main_t
type Io_main_t = uint32

// See: https://developer.apple.com/documentation/kernel/io_name_t
type Io_name_t = int8

// See: https://developer.apple.com/documentation/kernel/io_object_t
type Io_object_t = OSObject

// See: https://developer.apple.com/documentation/kernel/io_registry_entry_t
type Io_registry_entry_t = uintptr

// See: https://developer.apple.com/documentation/kernel/io_scalar_inband64_t
type Io_scalar_inband64_t = Io_user_scalar_t

// See: https://developer.apple.com/documentation/kernel/io_scalar_inband_t
type Io_scalar_inband_t = int

// See: https://developer.apple.com/documentation/kernel/io_service_t
type Io_service_t = uintptr

// See: https://developer.apple.com/documentation/kernel/io_stat_info_t
type Io_stat_info_t = unsafe.Pointer

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
type Ioctl_fcn_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ipc_eventlink_t
type Ipc_eventlink_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ipc_info_name_array_t
type Ipc_info_name_array_t = Ipc_info_name_t

// See: https://developer.apple.com/documentation/kernel/ipc_info_name_t
type Ipc_info_name_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ipc_info_port_t
type Ipc_info_port_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ipc_info_space_basic_t
type Ipc_info_space_basic_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ipc_info_space_t
type Ipc_info_space_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ipc_info_tree_name_array_t
type Ipc_info_tree_name_array_t = Ipc_info_tree_name_t

// See: https://developer.apple.com/documentation/kernel/ipc_info_tree_name_t
type Ipc_info_tree_name_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ipc_object_t
type Ipc_object_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ipc_port_t
type Ipc_port_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ipc_pthread_priority_value_t
type Ipc_pthread_priority_value_t = uint32

// See: https://developer.apple.com/documentation/kernel/ipc_space_inspect_t
type Ipc_space_inspect_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ipc_space_port_t
type Ipc_space_port_t = Ipc_space_t

// See: https://developer.apple.com/documentation/kernel/ipc_space_read_t
type Ipc_space_read_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ipc_space_t
type Ipc_space_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ipc_voucher_attr_control_t
type Ipc_voucher_attr_control_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ipc_voucher_attr_manager_t
type Ipc_voucher_attr_manager_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ipc_voucher_t
type Ipc_voucher_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ipi_handler_t
type Ipi_handler_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ipsec_dscp_mapping_t
type Ipsec_dscp_mapping_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kusbconnectable
type KUSBConnectable = int

// See: https://developer.apple.com/documentation/kernel/kauth_ace_rights_t
type Kauth_ace_rights_t = U_int32_t

// See: https://developer.apple.com/documentation/kernel/kauth_acl_eval_t
type Kauth_acl_eval_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kauth_action_t
type Kauth_action_t = int

// See: https://developer.apple.com/documentation/kernel/kauth_cred_t
type Kauth_cred_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kauth_listener_t
type Kauth_listener_t = uintptr

// See: https://developer.apple.com/documentation/kernel/kauth_scope_t
type Kauth_scope_t = uintptr

// See: https://developer.apple.com/documentation/kernel/kbdbitvector
type KbdBitVector = uint32

// See: https://developer.apple.com/documentation/kernel/kbufinfo_t
type Kbufinfo_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kc_format_t
type Kc_format_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kc_kind_t
type Kc_kind_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kcd_compression_type_t
type Kcd_compression_type_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kcdata_descriptor_t
type Kcdata_descriptor_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kcdata_item_t
type Kcdata_item_t = uintptr

// See: https://developer.apple.com/documentation/kernel/kcdata_iter_t
type Kcdata_iter_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kcdata_object_t
type Kcdata_object_t = uintptr

// See: https://developer.apple.com/documentation/kernel/kcdata_subtype_descriptor_t
type Kcdata_subtype_descriptor_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kctype_subtype_t
type Kctype_subtype_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kd_buf
type Kd_buf = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kd_buf_argtype
type Kd_buf_argtype = uint64

// See: https://developer.apple.com/documentation/kernel/kd_callback_t
type Kd_callback_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kd_callback_type
type Kd_callback_type = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kd_cpumap
type Kd_cpumap = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kd_cpumap_ext
type Kd_cpumap_ext = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kd_cpumap_header
type Kd_cpumap_header = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kd_event_matcher
type Kd_event_matcher = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kd_regtype
type Kd_regtype = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kd_threadmap
type Kd_threadmap = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kdebug_coproc_flags_t
type Kdebug_coproc_flags_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kdebug_flags_t
type Kdebug_flags_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kdebug_live_flags_t
type Kdebug_live_flags_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kdebug_test_t
type Kdebug_test_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kdp_event_t
type Kdp_event_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kern_ctl_ref
type Kern_ctl_ref = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kern_return_t
type Kern_return_t = int32

// See: https://developer.apple.com/documentation/kernel/kernel_boot_info_t
type Kernel_boot_info_t = int8

// See: https://developer.apple.com/documentation/kernel/kernel_resource_sizes_data_t
type Kernel_resource_sizes_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kernel_resource_sizes_t
type Kernel_resource_sizes_t = uintptr

// See: https://developer.apple.com/documentation/kernel/kernel_version_t
type Kernel_version_t = int8

// See: https://developer.apple.com/documentation/kernel/key_t
type Key_t = int32

// See: https://developer.apple.com/documentation/kernel/kf_override_flag_t
type Kf_override_flag_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kmod_args_t
type Kmod_args_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kmod_control_flavor_t
type Kmod_control_flavor_t = int

// See: https://developer.apple.com/documentation/kernel/kmod_info_32_v1_t
type Kmod_info_32_v1_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kmod_info_64_v1_t
type Kmod_info_64_v1_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kmod_info_array_t
type Kmod_info_array_t = Kmod_info_t

// See: https://developer.apple.com/documentation/kernel/kmod_info_t
type Kmod_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kmod_reference_t
type Kmod_reference_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kmod_start_func_t
type Kmod_start_func_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kmod_stop_func_t
type Kmod_stop_func_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kmod_t
type Kmod_t = int

// See: https://developer.apple.com/documentation/kernel/kobject_description_t
type Kobject_description_t = int8

// See: https://developer.apple.com/documentation/kernel/kpc_config_t
type Kpc_config_t = uint64

// See: https://developer.apple.com/documentation/kernel/kpc_pm_handler_t
type Kpc_pm_handler_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/kperf_kpc_flags_t
type Kperf_kpc_flags_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/labelstr_t
type Labelstr_t = int8

// See: https://developer.apple.com/documentation/kernel/launch_constraint_data_t
type Launch_constraint_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/lck_attr_t
type Lck_attr_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/lck_grp_attr_t
type Lck_grp_attr_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/lck_grp_t
type Lck_grp_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/lck_mtx_ext_t
type Lck_mtx_ext_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/lck_mtx_t
type Lck_mtx_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/lck_rw_t
type Lck_rw_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/lck_rw_type_t
type Lck_rw_type_t = uint

// See: https://developer.apple.com/documentation/kernel/lck_sleep_action_t
type Lck_sleep_action_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/lck_spin_t
type Lck_spin_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/lck_wake_action_t
type Lck_wake_action_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ldt_desc_t
type Ldt_desc_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ldt_t
type Ldt_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ledger_amount_t
type Ledger_amount_t = int64

// See: https://developer.apple.com/documentation/kernel/ledger_array_t
type Ledger_array_t = Ledger_t

// See: https://developer.apple.com/documentation/kernel/ledger_item_t
type Ledger_item_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/ledger_port_array_t
type Ledger_port_array_t = Ledger_array_t

// See: https://developer.apple.com/documentation/kernel/ledger_port_t
type Ledger_port_t = Ledger_t

// See: https://developer.apple.com/documentation/kernel/ledger_t
type Ledger_t = uintptr

// See: https://developer.apple.com/documentation/kernel/libsptm_cpu_state_t
type Libsptm_cpu_state_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/libsptm_error_t
type Libsptm_error_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/libsptm_refcnt_type_t
type Libsptm_refcnt_type_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/libsptm_state_t
type Libsptm_state_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/listxattrs_result_t
type Listxattrs_result_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/lock_set_port_t
type Lock_set_port_t = Lock_set_t

// See: https://developer.apple.com/documentation/kernel/lock_set_t
type Lock_set_t = uintptr

// See: https://developer.apple.com/documentation/kernel/lockgroup_info_array_t
type Lockgroup_info_array_t = Lockgroup_info_t

// See: https://developer.apple.com/documentation/kernel/lockgroup_info_t
type Lockgroup_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/lz4_hash_entry_t
type Lz4_hash_entry_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_assert_type_t
type Mach_assert_type_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_atm_subaid_t
type Mach_atm_subaid_t = uint64

// See: https://developer.apple.com/documentation/kernel/mach_bridge_regwrite_timestamp_func_t
type Mach_bridge_regwrite_timestamp_func_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_dead_name_notification_t
type Mach_dead_name_notification_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_error_fn_t
type Mach_error_fn_t = Mach_error_t

// See: https://developer.apple.com/documentation/kernel/mach_error_t
type Mach_error_t = int32

// See: https://developer.apple.com/documentation/kernel/mach_eventlink_t
type Mach_eventlink_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_exception_code_t
type Mach_exception_code_t = Mach_exception_data_type_t

// See: https://developer.apple.com/documentation/kernel/mach_exception_data_t
type Mach_exception_data_t = Mach_exception_data_type_t

// See: https://developer.apple.com/documentation/kernel/mach_exception_data_type_t
type Mach_exception_data_type_t = int64

// See: https://developer.apple.com/documentation/kernel/mach_exception_subcode_t
type Mach_exception_subcode_t = Mach_exception_data_type_t

// See: https://developer.apple.com/documentation/kernel/mach_memory_info_array_t
type Mach_memory_info_array_t = Mach_memory_info_t

// See: https://developer.apple.com/documentation/kernel/mach_memory_info_t
type Mach_memory_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_msg_audit_trailer_t
type Mach_msg_audit_trailer_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_msg_base_t
type Mach_msg_base_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_msg_bits_t
type Mach_msg_bits_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_body_t
type Mach_msg_body_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_msg_context_trailer_t
type Mach_msg_context_trailer_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_msg_copy_options_t
type Mach_msg_copy_options_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_descriptor_type_t
type Mach_msg_descriptor_type_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_empty_rcv_t
type Mach_msg_empty_rcv_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_msg_empty_send_t
type Mach_msg_empty_send_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_msg_filter_id
type Mach_msg_filter_id = int

// See: https://developer.apple.com/documentation/kernel/mach_msg_format_0_trailer_t
type Mach_msg_format_0_trailer_t = Mach_msg_security_trailer_t

// See: https://developer.apple.com/documentation/kernel/mach_msg_guard_flags_t
type Mach_msg_guard_flags_t = uint

// See: https://developer.apple.com/documentation/kernel/mach_msg_guarded_port_descriptor32_t
type Mach_msg_guarded_port_descriptor32_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_msg_guarded_port_descriptor64_t
type Mach_msg_guarded_port_descriptor64_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_msg_guarded_port_descriptor_t
type Mach_msg_guarded_port_descriptor_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_msg_header_t
type Mach_msg_header_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_msg_id_t
type Mach_msg_id_t = int32

// See: https://developer.apple.com/documentation/kernel/mach_msg_mac_trailer_t
type Mach_msg_mac_trailer_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_msg_max_trailer_t
type Mach_msg_max_trailer_t = Mach_msg_mac_trailer_t

// See: https://developer.apple.com/documentation/kernel/mach_msg_ool_descriptor32_t
type Mach_msg_ool_descriptor32_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_msg_ool_descriptor64_t
type Mach_msg_ool_descriptor64_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_msg_ool_descriptor_t
type Mach_msg_ool_descriptor_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_msg_ool_ports_descriptor32_t
type Mach_msg_ool_ports_descriptor32_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_msg_ool_ports_descriptor64_t
type Mach_msg_ool_ports_descriptor64_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_msg_ool_ports_descriptor_t
type Mach_msg_ool_ports_descriptor_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_msg_option_t
type Mach_msg_option_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/mach_msg_options_t
type Mach_msg_options_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/mach_msg_port_descriptor_t
type Mach_msg_port_descriptor_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_msg_priority_t
type Mach_msg_priority_t = uint

// See: https://developer.apple.com/documentation/kernel/mach_msg_return_t
type Mach_msg_return_t = int32

// See: https://developer.apple.com/documentation/kernel/mach_msg_security_trailer_t
type Mach_msg_security_trailer_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_msg_seqno_trailer_t
type Mach_msg_seqno_trailer_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_msg_size_t
type Mach_msg_size_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_timeout_t
type Mach_msg_timeout_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/mach_msg_trailer_info_t
type Mach_msg_trailer_info_t = int8

// See: https://developer.apple.com/documentation/kernel/mach_msg_trailer_size_t
type Mach_msg_trailer_size_t = uint

// See: https://developer.apple.com/documentation/kernel/mach_msg_trailer_t
type Mach_msg_trailer_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_msg_trailer_type_t
type Mach_msg_trailer_type_t = uint

// See: https://developer.apple.com/documentation/kernel/mach_msg_type_descriptor_t
type Mach_msg_type_descriptor_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_msg_type_name_t
type Mach_msg_type_name_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_msg_type_number_t
type Mach_msg_type_number_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/mach_msg_type_size_t
type Mach_msg_type_size_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/mach_no_senders_notification_t
type Mach_no_senders_notification_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_port_array_t
type Mach_port_array_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_port_context_t
type Mach_port_context_t = Mach_vm_address_t

// See: https://developer.apple.com/documentation/kernel/mach_port_deleted_notification_t
type Mach_port_deleted_notification_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_port_delta_t
type Mach_port_delta_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/mach_port_destroyed_notification_t
type Mach_port_destroyed_notification_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_port_flavor_t
type Mach_port_flavor_t = int

// See: https://developer.apple.com/documentation/kernel/mach_port_guard_info_t
type Mach_port_guard_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_port_info_ext_t
type Mach_port_info_ext_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_port_info_t
type Mach_port_info_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/mach_port_limits_t
type Mach_port_limits_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_port_mscount_t
type Mach_port_mscount_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/mach_port_msgcount_t
type Mach_port_msgcount_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/mach_port_name_array_t
type Mach_port_name_array_t = Mach_port_name_t

// See: https://developer.apple.com/documentation/kernel/mach_port_name_t
type Mach_port_name_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/mach_port_options_ptr_t
type Mach_port_options_ptr_t = Mach_port_options_t

// See: https://developer.apple.com/documentation/kernel/mach_port_options_t
type Mach_port_options_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_port_qos_t
type Mach_port_qos_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_port_right_t
type Mach_port_right_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/mach_port_rights_t
type Mach_port_rights_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/mach_port_seqno_t
type Mach_port_seqno_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/mach_port_srights_t
type Mach_port_srights_t = uint

// See: https://developer.apple.com/documentation/kernel/mach_port_status_t
type Mach_port_status_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_port_t
type Mach_port_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_port_type_array_t
type Mach_port_type_array_t = Mach_port_type_t

// See: https://developer.apple.com/documentation/kernel/mach_port_type_t
type Mach_port_type_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/mach_port_urefs_t
type Mach_port_urefs_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/mach_send_once_notification_t
type Mach_send_once_notification_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_send_possible_notification_t
type Mach_send_possible_notification_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_service_port_info_data_t
type Mach_service_port_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_service_port_info_t
type Mach_service_port_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/mach_task_basic_info_data_t
type Mach_task_basic_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_task_basic_info_t
type Mach_task_basic_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/mach_task_flavor_t
type Mach_task_flavor_t = uint

// See: https://developer.apple.com/documentation/kernel/mach_thread_flavor_t
type Mach_thread_flavor_t = uint

// Mach_timebase_info_data_t is raw Mach Time API In general prefer to use the API clock_gettime_nsec_np(3), which deals in the same clocks (and more) in ns units. Conversion of ns to (resp. from) tick units as returned by the mach time APIs is performed by division (resp. multiplication) with the fraction returned by mach_timebase_info().
//
// See: https://developer.apple.com/documentation/kernel/mach_timebase_info_data_t
type Mach_timebase_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_timebase_info_t
type Mach_timebase_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_timespec_t
type Mach_timespec_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_vm_address_t
type Mach_vm_address_t = uint64

// See: https://developer.apple.com/documentation/kernel/mach_vm_address_ut
type Mach_vm_address_ut = Mach_vm_address_t

// See: https://developer.apple.com/documentation/kernel/mach_vm_info_region_t
type Mach_vm_info_region_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_vm_offset_t
type Mach_vm_offset_t = uint64

// See: https://developer.apple.com/documentation/kernel/mach_vm_offset_ut
type Mach_vm_offset_ut = Mach_vm_offset_t

// See: https://developer.apple.com/documentation/kernel/mach_vm_range_flags_t
type Mach_vm_range_flags_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_vm_range_flavor_t
type Mach_vm_range_flavor_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_vm_range_recipe_t
type Mach_vm_range_recipe_t = Mach_vm_range_recipe_v1_t

// See: https://developer.apple.com/documentation/kernel/mach_vm_range_recipe_v1_t
type Mach_vm_range_recipe_v1_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_vm_range_recipe_v1_ut
type Mach_vm_range_recipe_v1_ut = Mach_vm_range_recipe_v1_t

// See: https://developer.apple.com/documentation/kernel/mach_vm_range_recipes_raw_t
type Mach_vm_range_recipes_raw_t = uint8

// See: https://developer.apple.com/documentation/kernel/mach_vm_range_t
type Mach_vm_range_t = uintptr

// See: https://developer.apple.com/documentation/kernel/mach_vm_range_tag_t
type Mach_vm_range_tag_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_vm_read_entry_t
type Mach_vm_read_entry_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_vm_size_t
type Mach_vm_size_t = uint64

// See: https://developer.apple.com/documentation/kernel/mach_vm_size_ut
type Mach_vm_size_ut = Mach_vm_size_t

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_command_t
type Mach_voucher_attr_command_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_content_size_t
type Mach_voucher_attr_content_size_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_content_t
type Mach_voucher_attr_content_t = uint8

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_control_flags_t
type Mach_voucher_attr_control_flags_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_control_t
type Mach_voucher_attr_control_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_importance_refs
type Mach_voucher_attr_importance_refs = uint32

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_key_array_t
type Mach_voucher_attr_key_array_t = Mach_voucher_attr_key_t

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
type Mach_voucher_attr_raw_recipe_t = uint8

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_recipe_command_array_t
type Mach_voucher_attr_recipe_command_array_t = Mach_voucher_attr_recipe_command_t

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_recipe_command_t
type Mach_voucher_attr_recipe_command_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_recipe_data_t
type Mach_voucher_attr_recipe_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_recipe_size_t
type Mach_voucher_attr_recipe_size_t = Mach_msg_type_number_t

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_recipe_t
type Mach_voucher_attr_recipe_t = Mach_voucher_attr_recipe_data_t

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_value_flags_t
type Mach_voucher_attr_value_flags_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_value_handle_array_size_t
type Mach_voucher_attr_value_handle_array_size_t = Mach_msg_type_number_t

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_value_handle_array_t
type Mach_voucher_attr_value_handle_array_t = Mach_voucher_attr_value_handle_t

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_value_handle_t
type Mach_voucher_attr_value_handle_t = uint64

// See: https://developer.apple.com/documentation/kernel/mach_voucher_attr_value_reference_t
type Mach_voucher_attr_value_reference_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_voucher_name_array_t
type Mach_voucher_name_array_t = Mach_voucher_name_t

// See: https://developer.apple.com/documentation/kernel/mach_voucher_name_t
type Mach_voucher_name_t = Mach_port_name_t

// See: https://developer.apple.com/documentation/kernel/mach_voucher_selector_t
type Mach_voucher_selector_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_voucher_t
type Mach_voucher_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_zone_info_array_t
type Mach_zone_info_array_t = Mach_zone_info_t

// See: https://developer.apple.com/documentation/kernel/mach_zone_info_t
type Mach_zone_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_zone_name_array_t
type Mach_zone_name_array_t = Mach_zone_name_t

// See: https://developer.apple.com/documentation/kernel/mach_zone_name_t
type Mach_zone_name_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mailbox_offset_t
type Mailbox_offset_t = uint64

// See: https://developer.apple.com/documentation/kernel/mb_class_stat_t
type Mb_class_stat_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mb_stat_t
type Mb_stat_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mbstate_t
type Mbstate_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mbuf_bptr_t
type Mbuf_bptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/mbuf_csum_performed_flags_t
type Mbuf_csum_performed_flags_t = U_int32_t

// See: https://developer.apple.com/documentation/kernel/mbuf_csum_request_flags_t
type Mbuf_csum_request_flags_t = U_int32_t

// See: https://developer.apple.com/documentation/kernel/mbuf_flags_t
type Mbuf_flags_t = U_int32_t

// See: https://developer.apple.com/documentation/kernel/mbuf_how_t
type Mbuf_how_t = U_int32_t

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
type Mbuf_tag_id_t = U_int32_t

// See: https://developer.apple.com/documentation/kernel/mbuf_tag_type_t
type Mbuf_tag_type_t = U_int16_t

// Mbuf_traffic_class_t is traffic class of a packet.
//
// See: https://developer.apple.com/documentation/kernel/mbuf_traffic_class_t
type Mbuf_traffic_class_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mbuf_tso_request_flags_t
type Mbuf_tso_request_flags_t = U_int32_t

// See: https://developer.apple.com/documentation/kernel/mbuf_type_t
type Mbuf_type_t = U_int32_t

// See: https://developer.apple.com/documentation/kernel/mcc_ecc_event_t
type Mcc_ecc_event_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mcc_ecc_version_t
type Mcc_ecc_version_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mcc_flags_t
type Mcc_flags_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mcontext_t
type Mcontext_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mem_entry_name_port_t
type Mem_entry_name_port_t = uint32

// See: https://developer.apple.com/documentation/kernel/memory_object_array_t
type Memory_object_array_t = Memory_object_t

// See: https://developer.apple.com/documentation/kernel/memory_object_attr_info_data_t
type Memory_object_attr_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/memory_object_attr_info_t
type Memory_object_attr_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/memory_object_behave_info_data_t
type Memory_object_behave_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/memory_object_behave_info_t
type Memory_object_behave_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/memory_object_cluster_size_t
type Memory_object_cluster_size_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/memory_object_control_t
type Memory_object_control_t = uint32

// See: https://developer.apple.com/documentation/kernel/memory_object_copy_strategy_t
type Memory_object_copy_strategy_t = int

// See: https://developer.apple.com/documentation/kernel/memory_object_default_t
type Memory_object_default_t = uint32

// See: https://developer.apple.com/documentation/kernel/memory_object_fault_info_t
type Memory_object_fault_info_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/memory_object_flavor_t
type Memory_object_flavor_t = int

// See: https://developer.apple.com/documentation/kernel/memory_object_info_data_t
type Memory_object_info_data_t = int

// See: https://developer.apple.com/documentation/kernel/memory_object_info_t
type Memory_object_info_t = int

// See: https://developer.apple.com/documentation/kernel/memory_object_name_t
type Memory_object_name_t = uint32

// See: https://developer.apple.com/documentation/kernel/memory_object_offset_t
type Memory_object_offset_t = uint64

// See: https://developer.apple.com/documentation/kernel/memory_object_offset_ut
type Memory_object_offset_ut = Memory_object_offset_t

// See: https://developer.apple.com/documentation/kernel/memory_object_perf_info_data_t
type Memory_object_perf_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/memory_object_perf_info_t
type Memory_object_perf_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/memory_object_return_t
type Memory_object_return_t = int

// See: https://developer.apple.com/documentation/kernel/memory_object_size_t
type Memory_object_size_t = uint64

// See: https://developer.apple.com/documentation/kernel/memory_object_size_ut
type Memory_object_size_ut = Memory_object_size_t

// See: https://developer.apple.com/documentation/kernel/memory_object_t
type Memory_object_t = uint32

// See: https://developer.apple.com/documentation/kernel/microstackshot_flags_t
type Microstackshot_flags_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mig_impl_routine_t
type Mig_impl_routine_t = int32

// See: https://developer.apple.com/documentation/kernel/mig_reply_error_t
type Mig_reply_error_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mig_routine_arg_descriptor_t
type Mig_routine_arg_descriptor_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mig_routine_descriptor
type Mig_routine_descriptor = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mig_routine_descriptor_t
type Mig_routine_descriptor_t = Mig_routine_descriptor

// See: https://developer.apple.com/documentation/kernel/mig_routine_t
type Mig_routine_t = uintptr

// See: https://developer.apple.com/documentation/kernel/mig_subsystem_t
type Mig_subsystem_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mig_symtab_t
type Mig_symtab_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ml_cpu_info_t
type Ml_cpu_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ml_page_protection_t
type Ml_page_protection_t = int

// See: https://developer.apple.com/documentation/kernel/ml_processor_info_t
type Ml_processor_info_t = unsafe.Pointer

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
type Mph_panic_flags_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mpsc_queue_chain_t
type Mpsc_queue_chain_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mpsc_queue_head_t
type Mpsc_queue_head_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/msg_labels_t
type Msg_labels_t = unsafe.Pointer

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
type Natural_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/net_init_func_ptr
type Net_init_func_ptr = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/netaddr_t
type Netaddr_t = uint32

// See: https://developer.apple.com/documentation/kernel/network_port_t
type Network_port_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nfs_fsid
type Nfs_fsid = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nfs_handle
type Nfs_handle = uint8

// See: https://developer.apple.com/documentation/kernel/nfs_specdata
type Nfs_specdata = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nfs_stateid
type Nfs_stateid = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nfs_supported_kerberos_etypes
type Nfs_supported_kerberos_etypes = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nfserr_info_t
type Nfserr_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nfstype
type Nfstype = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nfsuint64
type Nfsuint64 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nlink_t
type Nlink_t = uint16

// See: https://developer.apple.com/documentation/kernel/notify_port_t
type Notify_port_t = uint32

// See: https://developer.apple.com/documentation/kernel/np_uid_t
type Np_uid_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/nspace_name_t
type Nspace_name_t = int8

// See: https://developer.apple.com/documentation/kernel/nspace_path_t
type Nspace_path_t = int8

// See: https://developer.apple.com/documentation/kernel/ntsid_t
type Ntsid_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/off_t
type Off_t = int64

// See: https://developer.apple.com/documentation/kernel/open_close_fcn_t
type Open_close_fcn_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/os_block_t
type Os_block_t = func()

// See: https://developer.apple.com/documentation/kernel/os_log_coproc_reg_t
type Os_log_coproc_reg_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/os_log_t
type Os_log_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/os_log_type_t
type Os_log_type_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/packed_uchar16
type Packed_uchar16 = uint8

// See: https://developer.apple.com/documentation/kernel/packed_uchar32
type Packed_uchar32 = uint8

// See: https://developer.apple.com/documentation/kernel/packed_uchar64
type Packed_uchar64 = uint8

// See: https://developer.apple.com/documentation/kernel/packed_ushort4
type Packed_ushort4 = uint16

// See: https://developer.apple.com/documentation/kernel/page_address_array_t
type Page_address_array_t = Vm_offset_t

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
type Pointer_t = Vm_offset_t

// See: https://developer.apple.com/documentation/kernel/pointer_ut
type Pointer_ut = Pointer_t

// See: https://developer.apple.com/documentation/kernel/policy_base_data_t
type Policy_base_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/policy_base_t
type Policy_base_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/policy_fifo_base_data_t
type Policy_fifo_base_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/policy_fifo_base_t
type Policy_fifo_base_t = uintptr

// See: https://developer.apple.com/documentation/kernel/policy_fifo_info_data_t
type Policy_fifo_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/policy_fifo_info_t
type Policy_fifo_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/policy_fifo_limit_data_t
type Policy_fifo_limit_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/policy_fifo_limit_t
type Policy_fifo_limit_t = uintptr

// See: https://developer.apple.com/documentation/kernel/policy_info_data_t
type Policy_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/policy_info_t
type Policy_info_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/policy_limit_data_t
type Policy_limit_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/policy_limit_t
type Policy_limit_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/policy_rr_base_data_t
type Policy_rr_base_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/policy_rr_base_t
type Policy_rr_base_t = uintptr

// See: https://developer.apple.com/documentation/kernel/policy_rr_info_data_t
type Policy_rr_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/policy_rr_info_t
type Policy_rr_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/policy_rr_limit_data_t
type Policy_rr_limit_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/policy_rr_limit_t
type Policy_rr_limit_t = uintptr

// See: https://developer.apple.com/documentation/kernel/policy_t
type Policy_t = int

// See: https://developer.apple.com/documentation/kernel/policy_timeshare_base_data_t
type Policy_timeshare_base_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/policy_timeshare_base_t
type Policy_timeshare_base_t = uintptr

// See: https://developer.apple.com/documentation/kernel/policy_timeshare_info_data_t
type Policy_timeshare_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/policy_timeshare_info_t
type Policy_timeshare_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/policy_timeshare_limit_data_t
type Policy_timeshare_limit_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/policy_timeshare_limit_t
type Policy_timeshare_limit_t = uintptr

// See: https://developer.apple.com/documentation/kernel/port_name_array_t
type Port_name_array_t = Mach_port_name_t

// See: https://developer.apple.com/documentation/kernel/port_name_t
type Port_name_t = Mach_port_name_t

// See: https://developer.apple.com/documentation/kernel/port_t
type Port_t = uint32

// See: https://developer.apple.com/documentation/kernel/posix_cred_t
type Posix_cred_t = uintptr

// See: https://developer.apple.com/documentation/kernel/ppnum_t
type Ppnum_t = uint32

// See: https://developer.apple.com/documentation/kernel/priority_queue_compare_fn_t
type Priority_queue_compare_fn_t = func(unsafe.Pointer, unsafe.Pointer) int

// See: https://developer.apple.com/documentation/kernel/priority_queue_entry_deadline_t
type Priority_queue_entry_deadline_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/priority_queue_entry_sched_modifier_t
type Priority_queue_entry_sched_modifier_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/priority_queue_entry_sched_t
type Priority_queue_entry_sched_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/priority_queue_entry_stable_t
type Priority_queue_entry_stable_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/priority_queue_entry_t
type Priority_queue_entry_t = unsafe.Pointer

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
type Processor_array_t = Processor_t

// See: https://developer.apple.com/documentation/kernel/processor_basic_info_data_t
type Processor_basic_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/processor_basic_info_t
type Processor_basic_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/processor_cpu_load_info_data_t
type Processor_cpu_load_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/processor_cpu_load_info_t
type Processor_cpu_load_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/processor_cpu_stat64_data_t
type Processor_cpu_stat64_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/processor_cpu_stat64_t
type Processor_cpu_stat64_t = uintptr

// See: https://developer.apple.com/documentation/kernel/processor_cpu_stat_data_t
type Processor_cpu_stat_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/processor_cpu_stat_t
type Processor_cpu_stat_t = uintptr

// See: https://developer.apple.com/documentation/kernel/processor_flavor_t
type Processor_flavor_t = int

// See: https://developer.apple.com/documentation/kernel/processor_info_array_t
type Processor_info_array_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/processor_info_data_t
type Processor_info_data_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/processor_info_t
type Processor_info_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/processor_port_array_t
type Processor_port_array_t = Processor_array_t

// See: https://developer.apple.com/documentation/kernel/processor_port_t
type Processor_port_t = Processor_t

// See: https://developer.apple.com/documentation/kernel/processor_set_array_t
type Processor_set_array_t = Processor_set_t

// See: https://developer.apple.com/documentation/kernel/processor_set_basic_info_data_t
type Processor_set_basic_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/processor_set_basic_info_t
type Processor_set_basic_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/processor_set_control_port_t
type Processor_set_control_port_t = Processor_set_t

// See: https://developer.apple.com/documentation/kernel/processor_set_control_t
type Processor_set_control_t = uintptr

// See: https://developer.apple.com/documentation/kernel/processor_set_flavor_t
type Processor_set_flavor_t = int

// See: https://developer.apple.com/documentation/kernel/processor_set_info_data_t
type Processor_set_info_data_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/processor_set_info_t
type Processor_set_info_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/processor_set_load_info_data_t
type Processor_set_load_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/processor_set_load_info_t
type Processor_set_load_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/processor_set_name_array_t
type Processor_set_name_array_t = Processor_set_t

// See: https://developer.apple.com/documentation/kernel/processor_set_name_port_array_t
type Processor_set_name_port_array_t = Processor_set_array_t

// See: https://developer.apple.com/documentation/kernel/processor_set_name_port_t
type Processor_set_name_port_t = Processor_set_t

// See: https://developer.apple.com/documentation/kernel/processor_set_name_t
type Processor_set_name_t = Processor_set_t

// See: https://developer.apple.com/documentation/kernel/processor_set_port_t
type Processor_set_port_t = Processor_set_t

// See: https://developer.apple.com/documentation/kernel/processor_set_t
type Processor_set_t = uintptr

// See: https://developer.apple.com/documentation/kernel/processor_t
type Processor_t = uintptr

// Protocol_family_t is storage type for the protocol family.
//
// See: https://developer.apple.com/documentation/kernel/protocol_family_t
type Protocol_family_t = U_int32_t

// See: https://developer.apple.com/documentation/kernel/psize_fcn_t
type Psize_fcn_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ptrdiff_t
type Ptrdiff_t = int

// See: https://developer.apple.com/documentation/kernel/qaddr_t
type Qaddr_t = Quad_t

// See: https://developer.apple.com/documentation/kernel/quad_t
type Quad_t = int64

// See: https://developer.apple.com/documentation/kernel/queue_chain_t
type Queue_chain_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/queue_entry_t
type Queue_entry_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/queue_head_t
type Queue_head_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/queue_t
type Queue_t = uintptr

// See: https://developer.apple.com/documentation/kernel/read_write_fcn_t
type Read_write_fcn_t = unsafe.Pointer

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
type Routine_arg_descriptor = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/routine_arg_descriptor_t
type Routine_arg_descriptor_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/routine_arg_offset
type Routine_arg_offset = uint

// See: https://developer.apple.com/documentation/kernel/routine_arg_size
type Routine_arg_size = uint

// See: https://developer.apple.com/documentation/kernel/routine_arg_type
type Routine_arg_type = uint

// See: https://developer.apple.com/documentation/kernel/routine_descriptor_t
type Routine_descriptor_t = uintptr

// See: https://developer.apple.com/documentation/kernel/rpc_routine_arg_descriptor_t
type Rpc_routine_arg_descriptor_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/rpc_routine_descriptor_t
type Rpc_routine_descriptor_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/rpc_subsystem_t
type Rpc_subsystem_t = unsafe.Pointer

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
type Rune_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/rusage_info_current
type Rusage_info_current = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/rusage_info_t
type Rusage_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/sa_endpoints_t
type Sa_endpoints_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/sa_family_t
type Sa_family_t = uint8

// See: https://developer.apple.com/documentation/kernel/sae_associd_t
type Sae_associd_t = uint32

// See: https://developer.apple.com/documentation/kernel/sae_connid_t
type Sae_connid_t = uint32

// See: https://developer.apple.com/documentation/kernel/secure_boot_cryptex_args_t
type Secure_boot_cryptex_args_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/security_token_t
type Security_token_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/segsz_t
type Segsz_t = int32

// See: https://developer.apple.com/documentation/kernel/sel_t
type Sel_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/select_fcn_t
type Select_fcn_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/semaphore_port_t
type Semaphore_port_t = Semaphore_t

// See: https://developer.apple.com/documentation/kernel/semaphore_t
type Semaphore_t = uintptr

// See: https://developer.apple.com/documentation/kernel/sflt_data_flag_t
type Sflt_data_flag_t = U_int32_t

// See: https://developer.apple.com/documentation/kernel/sflt_event_t
type Sflt_event_t = U_int32_t

// See: https://developer.apple.com/documentation/kernel/sflt_flags
type Sflt_flags = U_int32_t

// Sflt_handle is a 4 byte identifier used with the SO_NKE socket option to identify the socket filter to be attached.
//
// See: https://developer.apple.com/documentation/kernel/sflt_handle
type Sflt_handle = U_int32_t

// See: https://developer.apple.com/documentation/kernel/shared_file_mapping_slide_np_t
type Shared_file_mapping_slide_np_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/shared_file_mapping_slide_np_ut
type Shared_file_mapping_slide_np_ut = Shared_file_mapping_slide_np_t

// See: https://developer.apple.com/documentation/kernel/shmatt_t
type Shmatt_t = uint16

// See: https://developer.apple.com/documentation/kernel/sig_atomic_t
type Sig_atomic_t = int

// See: https://developer.apple.com/documentation/kernel/sig_t
type Sig_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/siginfo_t
type Siginfo_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/sigset_t
type Sigset_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/size_t
type Size_t = uintptr

// See: https://developer.apple.com/documentation/kernel/size_ut
type Size_ut = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/sleepwakenote
type SleepWakeNote = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/sleep_type_t
type Sleep_type_t = int

// See: https://developer.apple.com/documentation/kernel/smr_cb_t
type Smr_cb_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/smr_node_t
type Smr_node_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/smr_seq_t
type Smr_seq_t = uint

// See: https://developer.apple.com/documentation/kernel/smr_t
type Smr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/so_gen_t
type So_gen_t = U_quad_t

// See: https://developer.apple.com/documentation/kernel/sock_storage
type Sock_storage = uint32

// See: https://developer.apple.com/documentation/kernel/sockaddr_bptr_t
type Sockaddr_bptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/sockaddr_ptr_ref_t
type Sockaddr_ptr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/sockaddr_ptr_t
type Sockaddr_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/sockaddr_ref_ptr_t
type Sockaddr_ref_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/sockaddr_ref_ref_t
type Sockaddr_ref_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/sockaddr_ref_t
type Sockaddr_ref_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/sockaddr_storage_bptr_t
type Sockaddr_storage_bptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/sockaddr_storage_ptr_ref_t
type Sockaddr_storage_ptr_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/sockaddr_storage_ptr_t
type Sockaddr_storage_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/sockaddr_storage_ref_ptr_t
type Sockaddr_storage_ref_ptr_t = uintptr

// See: https://developer.apple.com/documentation/kernel/sockaddr_storage_ref_ref_t
type Sockaddr_storage_ref_ref_t = uintptr

// See: https://developer.apple.com/documentation/kernel/sockaddr_storage_ref_t
type Sockaddr_storage_ref_t = unsafe.Pointer

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
type Sockopt_dir = U_int8_t

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
type Sptm_call_regs_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/sptm_consistent_debug_t
type Sptm_consistent_debug_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/sptm_dispatch_endpoint_id_t
type Sptm_dispatch_endpoint_id_t = uint8

// See: https://developer.apple.com/documentation/kernel/sptm_dispatch_table_id_t
type Sptm_dispatch_table_id_t = uint8

// See: https://developer.apple.com/documentation/kernel/sptm_dispatch_target_t
type Sptm_dispatch_target_t = uint64

// See: https://developer.apple.com/documentation/kernel/sptm_domain_t
type Sptm_domain_t = uint8

// See: https://developer.apple.com/documentation/kernel/sptm_frame_type_t
type Sptm_frame_type_t = uint8

// See: https://developer.apple.com/documentation/kernel/sptm_instance_id_t
type Sptm_instance_id_t = uint16

// See: https://developer.apple.com/documentation/kernel/sptm_iommu_id_t
type Sptm_iommu_id_t = uint8

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
type Sptm_pt_level_t = uint8

// See: https://developer.apple.com/documentation/kernel/sptm_pte_t
type Sptm_pte_t = uint64

// See: https://developer.apple.com/documentation/kernel/sptm_return_t
type Sptm_return_t = uint32

// See: https://developer.apple.com/documentation/kernel/sptm_retype_params_t
type Sptm_retype_params_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/sptm_trace_buffer_t
type Sptm_trace_buffer_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/sptm_trace_t
type Sptm_trace_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/sptm_tte_t
type Sptm_tte_t = uint64

// See: https://developer.apple.com/documentation/kernel/sptm_vaddr_t
type Sptm_vaddr_t = uint64

// See: https://developer.apple.com/documentation/kernel/sptm_vector_type_t
type Sptm_vector_type_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/sptm_vmid_t
type Sptm_vmid_t = uint16

// See: https://developer.apple.com/documentation/kernel/sptm_voff_t
type Sptm_voff_t = uint64

// See: https://developer.apple.com/documentation/kernel/ssize_t
type Ssize_t = int

// See: https://developer.apple.com/documentation/kernel/stack_t
type Stack_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/stackshot_flags_t
type Stackshot_flags_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/stop_fcn_t
type Stop_fcn_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/strategy_fcn_t
type Strategy_fcn_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/string_t
type String_t = int8

// See: https://developer.apple.com/documentation/kernel/subaid_t
type Subaid_t = uint64

// See: https://developer.apple.com/documentation/kernel/suseconds_t
type Suseconds_t = int32

// See: https://developer.apple.com/documentation/kernel/swblk_t
type Swblk_t = int32

// See: https://developer.apple.com/documentation/kernel/symtab_name_t
type Symtab_name_t = int8

// See: https://developer.apple.com/documentation/kernel/sync_policy_t
type Sync_policy_t = int

// See: https://developer.apple.com/documentation/kernel/syscall_arg_t
type Syscall_arg_t = U_int64_t

// See: https://developer.apple.com/documentation/kernel/syscp_id_instructions_feat_1_reg
type Syscp_ID_instructions_feat_1_reg = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/tdevicerequestdirection
type TDeviceRequestDirection = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/tdevicerequestrecipient
type TDeviceRequestRecipient = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/tdevicerequesttype
type TDeviceRequestType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/tendpointdirection
type TEndpointDirection = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/tendpointsynchronizationtype
type TEndpointSynchronizationType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/tendpointtype
type TEndpointType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/tendpointusagetype
type TEndpointUsageType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/tiopcideviceresetoptions
type TIOPCIDeviceResetOptions = uint

// See: https://developer.apple.com/documentation/kernel/tiopcideviceresettypes
type TIOPCIDeviceResetTypes = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/tiopcilinkspeed
type TIOPCILinkSpeed = unsafe.Pointer

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
type TUSBCTypeCableType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/tusbdevicelpmstatus
type TUSBDeviceLPMStatus = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/tusbhostconnectortype
type TUSBHostConnectorType = unsafe.Pointer

// TUSBHostDeviceAddress is the USB host device address.
//
// See: https://developer.apple.com/documentation/kernel/tusbhostdeviceaddress
type TUSBHostDeviceAddress = uint16

// See: https://developer.apple.com/documentation/kernel/tusbhostportconnectable
type TUSBHostPortConnectable = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/tusbhostpowersourcetype
type TUSBHostPowerSourceType = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/tusblinkstate
type TUSBLinkState = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_absolutetime_info_data_t
type Task_absolutetime_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_absolutetime_info_t
type Task_absolutetime_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_affinity_tag_info_data_t
type Task_affinity_tag_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_affinity_tag_info_t
type Task_affinity_tag_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_array_t
type Task_array_t = Task_t

// See: https://developer.apple.com/documentation/kernel/task_basic_info_32_data_t
type Task_basic_info_32_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_basic_info_32_t
type Task_basic_info_32_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_basic_info_64_2_data_t
type Task_basic_info_64_2_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_basic_info_64_2_t
type Task_basic_info_64_2_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_basic_info_64_data_t
type Task_basic_info_64_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_basic_info_64_t
type Task_basic_info_64_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_basic_info_data_t
type Task_basic_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_basic_info_t
type Task_basic_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_category_policy_data_t
type Task_category_policy_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_category_policy_t
type Task_category_policy_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_corpse_forking_behavior_t
type Task_corpse_forking_behavior_t = uint32

// See: https://developer.apple.com/documentation/kernel/task_crashinfo_item_t
type Task_crashinfo_item_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_dyld_info_data_t
type Task_dyld_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_dyld_info_t
type Task_dyld_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_events_info_data_t
type Task_events_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_events_info_t
type Task_events_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_exc_guard_behavior_t
type Task_exc_guard_behavior_t = uint32

// See: https://developer.apple.com/documentation/kernel/task_extmod_info_data_t
type Task_extmod_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_extmod_info_t
type Task_extmod_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_flags_info_data_t
type Task_flags_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_flags_info_t
type Task_flags_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_flavor_t
type Task_flavor_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/task_gate_t
type Task_gate_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_id_token_t
type Task_id_token_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_info_data_t
type Task_info_data_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/task_info_t
type Task_info_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/task_inspect_basic_counts_data_t
type Task_inspect_basic_counts_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_inspect_basic_counts_t
type Task_inspect_basic_counts_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_inspect_flavor_t
type Task_inspect_flavor_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/task_inspect_info_t
type Task_inspect_info_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/task_inspect_t
type Task_inspect_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_kernelmemory_info_data_t
type Task_kernelmemory_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_kernelmemory_info_t
type Task_kernelmemory_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_latency_qos_t
type Task_latency_qos_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/task_name_t
type Task_name_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_policy_flavor_t
type Task_policy_flavor_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/task_policy_get_t
type Task_policy_get_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_policy_set_t
type Task_policy_set_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_policy_t
type Task_policy_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/task_port_array_t
type Task_port_array_t = Task_array_t

// See: https://developer.apple.com/documentation/kernel/task_port_t
type Task_port_t = Task_t

// See: https://developer.apple.com/documentation/kernel/task_power_info_data_t
type Task_power_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_power_info_t
type Task_power_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_power_info_v2_data_t
type Task_power_info_v2_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_power_info_v2_t
type Task_power_info_v2_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_purgable_info_t
type Task_purgable_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_qos_policy_t
type Task_qos_policy_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_read_t
type Task_read_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_restartable_range_array_t
type Task_restartable_range_array_t = Task_restartable_range_t

// See: https://developer.apple.com/documentation/kernel/task_restartable_range_t
type Task_restartable_range_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_role_t
type Task_role_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_special_port_t
type Task_special_port_t = int

// See: https://developer.apple.com/documentation/kernel/task_suspension_token_t
type Task_suspension_token_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_t
type Task_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_thread_times_info_data_t
type Task_thread_times_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_thread_times_info_t
type Task_thread_times_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_throughput_qos_t
type Task_throughput_qos_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/task_trace_memory_info_data_t
type Task_trace_memory_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_trace_memory_info_t
type Task_trace_memory_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_vm_info_data_t
type Task_vm_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_vm_info_t
type Task_vm_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_wait_state_info_data_t
type Task_wait_state_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/task_wait_state_info_t
type Task_wait_state_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/task_zone_info_array_t
type Task_zone_info_array_t = Task_zone_info_t

// See: https://developer.apple.com/documentation/kernel/task_zone_info_t
type Task_zone_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/tcflag_t
type Tcflag_t = uint

// See: https://developer.apple.com/documentation/kernel/tcp_cc
type Tcp_cc = uint32

// See: https://developer.apple.com/documentation/kernel/tcp_connection_client_accurate_ecn_state_t
type Tcp_connection_client_accurate_ecn_state_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/tcp_connection_server_accurate_ecn_state_t
type Tcp_connection_server_accurate_ecn_state_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/tcp_notify_ack_id_t
type Tcp_notify_ack_id_t = U_int32_t

// See: https://developer.apple.com/documentation/kernel/tcp_seq
type Tcp_seq = uint32

// See: https://developer.apple.com/documentation/kernel/telemetry_notice_t
type Telemetry_notice_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/text_encoding_t
type Text_encoding_t = U_int32_t

// See: https://developer.apple.com/documentation/kernel/thread_act_array_t
type Thread_act_array_t = Thread_act_t

// See: https://developer.apple.com/documentation/kernel/thread_act_port_array_t
type Thread_act_port_array_t = Thread_act_array_t

// See: https://developer.apple.com/documentation/kernel/thread_act_port_t
type Thread_act_port_t = Thread_act_t

// See: https://developer.apple.com/documentation/kernel/thread_act_t
type Thread_act_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_affinity_policy_data_t
type Thread_affinity_policy_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/thread_affinity_policy_t
type Thread_affinity_policy_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_array_t
type Thread_array_t = Thread_t

// See: https://developer.apple.com/documentation/kernel/thread_background_policy_data_t
type Thread_background_policy_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/thread_background_policy_t
type Thread_background_policy_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_basic_info_data_t
type Thread_basic_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/thread_basic_info_t
type Thread_basic_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_call_options_t
type Thread_call_options_t = uint32

// See: https://developer.apple.com/documentation/kernel/thread_call_param_t
type Thread_call_param_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/thread_call_priority_t
type Thread_call_priority_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/thread_call_t
type Thread_call_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_extended_info_data_t
type Thread_extended_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/thread_extended_info_t
type Thread_extended_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_extended_policy_data_t
type Thread_extended_policy_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/thread_extended_policy_t
type Thread_extended_policy_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_flavor_t
type Thread_flavor_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/thread_identifier_info_data_t
type Thread_identifier_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/thread_identifier_info_t
type Thread_identifier_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_info_data_t
type Thread_info_data_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/thread_info_t
type Thread_info_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/thread_inspect_t
type Thread_inspect_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_latency_qos_policy_data_t
type Thread_latency_qos_policy_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/thread_latency_qos_policy_t
type Thread_latency_qos_policy_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_latency_qos_t
type Thread_latency_qos_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/thread_policy_flavor_t
type Thread_policy_flavor_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/thread_policy_t
type Thread_policy_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/thread_port_array_t
type Thread_port_array_t = Thread_array_t

// See: https://developer.apple.com/documentation/kernel/thread_port_t
type Thread_port_t = Thread_t

// See: https://developer.apple.com/documentation/kernel/thread_precedence_policy_data_t
type Thread_precedence_policy_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/thread_precedence_policy_t
type Thread_precedence_policy_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_read_t
type Thread_read_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_selfcounts_kind_t
type Thread_selfcounts_kind_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/thread_standard_policy_data_t
type Thread_standard_policy_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/thread_standard_policy_t
type Thread_standard_policy_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_state_data_t
type Thread_state_data_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/thread_state_flavor_array_t
type Thread_state_flavor_array_t = Thread_state_flavor_t

// See: https://developer.apple.com/documentation/kernel/thread_state_flavor_t
type Thread_state_flavor_t = int

// See: https://developer.apple.com/documentation/kernel/thread_state_t
type Thread_state_t = Natural_t

// See: https://developer.apple.com/documentation/kernel/thread_t
type Thread_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_throughput_qos_policy_data_t
type Thread_throughput_qos_policy_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/thread_throughput_qos_policy_t
type Thread_throughput_qos_policy_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_throughput_qos_t
type Thread_throughput_qos_t = Integer_t

// See: https://developer.apple.com/documentation/kernel/thread_time_constraint_policy_data_t
type Thread_time_constraint_policy_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/thread_time_constraint_policy_t
type Thread_time_constraint_policy_t = uintptr

// See: https://developer.apple.com/documentation/kernel/thread_turnstileinfo_t
type Thread_turnstileinfo_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/thread_turnstileinfo_v2_t
type Thread_turnstileinfo_v2_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/thread_waitinfo_t
type Thread_waitinfo_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/thread_waitinfo_v2_t
type Thread_waitinfo_v2_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/throttle_info_handle_t
type Throttle_info_handle_t = uintptr

// See: https://developer.apple.com/documentation/kernel/time_t
type Time_t = int64

// See: https://developer.apple.com/documentation/kernel/time_value_t
type Time_value_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/token_t
type Token_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/trap_gate_t
type Trap_gate_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/tss_desc_t
type Tss_desc_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/tss_t
type Tss_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/uint-4rm
type UInt = uint

// See: https://developer.apple.com/documentation/kernel/uintf
type UIntf = uint

// See: https://developer.apple.com/documentation/kernel/ulong
type ULong = uint

// See: https://developer.apple.com/documentation/kernel/ulongf
type ULongf = uint

// See: https://developer.apple.com/documentation/kernel/u_char
type U_char = uint8

// See: https://developer.apple.com/documentation/kernel/u_int
type U_int = uint

// See: https://developer.apple.com/documentation/kernel/u_int16_t
type U_int16_t = uint16

// See: https://developer.apple.com/documentation/kernel/u_int32_t
type U_int32_t = uint

// See: https://developer.apple.com/documentation/kernel/u_int64_t
type U_int64_t = uint64

// See: https://developer.apple.com/documentation/kernel/u_int8_t
type U_int8_t = uint8

// See: https://developer.apple.com/documentation/kernel/u_long
type U_long = uint

// See: https://developer.apple.com/documentation/kernel/u_quad_t
type U_quad_t = U_int64_t

// See: https://developer.apple.com/documentation/kernel/u_short
type U_short = uint16

// See: https://developer.apple.com/documentation/kernel/ucontext64_t
type Ucontext64_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/ucontext_t
type Ucontext_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/uext_object_t
type Uext_object_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/uid_t
type Uid_t = uint32

// See: https://developer.apple.com/documentation/kernel/uint
type Uint = unsafe.Pointer

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
type Uint_fast8_t = uint8

// See: https://developer.apple.com/documentation/kernel/uint_least16_t
type Uint_least16_t = uint16

// See: https://developer.apple.com/documentation/kernel/uint_least32_t
type Uint_least32_t = uint32

// See: https://developer.apple.com/documentation/kernel/uint_least64_t
type Uint_least64_t = uint64

// See: https://developer.apple.com/documentation/kernel/uint_least8_t
type Uint_least8_t = uint8

// See: https://developer.apple.com/documentation/kernel/uintmax_t
type Uintmax_t = uint

// See: https://developer.apple.com/documentation/kernel/uintptr_t
type Uintptr_t = uintptr

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
type Unp_gen_t = U_quad_t

// See: https://developer.apple.com/documentation/kernel/upl_control_flags_t
type Upl_control_flags_t = uint64

// See: https://developer.apple.com/documentation/kernel/upl_offset_t
type Upl_offset_t = uint32

// See: https://developer.apple.com/documentation/kernel/upl_page_info_array_t
type Upl_page_info_array_t = Upl_page_info_t

// See: https://developer.apple.com/documentation/kernel/upl_page_info_t
type Upl_page_info_t = unsafe.Pointer

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
type User32_fchecklv_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/user32_fsignatures_t
type User32_fsignatures_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/user32_long_t
type User32_long_t = int32

// See: https://developer.apple.com/documentation/kernel/user32_msglen_t
type User32_msglen_t = User32_ulong_t

// See: https://developer.apple.com/documentation/kernel/user32_msgqnum_t
type User32_msgqnum_t = User32_ulong_t

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
type User64_msglen_t = User64_ulong_t

// See: https://developer.apple.com/documentation/kernel/user64_msgqnum_t
type User64_msgqnum_t = User64_ulong_t

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
type User_addr_t = U_int64_t

// See: https://developer.apple.com/documentation/kernel/user_addr_ut
type User_addr_ut = User_addr_t

// See: https://developer.apple.com/documentation/kernel/user_fchecklv_t
type User_fchecklv_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/user_fsignatures_t
type User_fsignatures_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/user_fsupplement_t
type User_fsupplement_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/user_long_t
type User_long_t = int64

// See: https://developer.apple.com/documentation/kernel/user_msglen_t
type User_msglen_t = User_ulong_t

// See: https://developer.apple.com/documentation/kernel/user_msgqnum_t
type User_msgqnum_t = User_ulong_t

// See: https://developer.apple.com/documentation/kernel/user_off_t
type User_off_t = int64

// See: https://developer.apple.com/documentation/kernel/user_size_t
type User_size_t = U_int64_t

// See: https://developer.apple.com/documentation/kernel/user_size_ut
type User_size_ut = User_size_t

// See: https://developer.apple.com/documentation/kernel/user_speed_t
type User_speed_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/user_ssize_t
type User_ssize_t = int64

// See: https://developer.apple.com/documentation/kernel/user_subsystem_t
type User_subsystem_t = int8

// See: https://developer.apple.com/documentation/kernel/user_tcflag_t
type User_tcflag_t = uint64

// See: https://developer.apple.com/documentation/kernel/user_time_t
type User_time_t = int64

// See: https://developer.apple.com/documentation/kernel/user_ulong_t
type User_ulong_t = U_int64_t

// See: https://developer.apple.com/documentation/kernel/ushort
type Ushort = uint16

// See: https://developer.apple.com/documentation/kernel/uuid_string_t
type Uuid_string_t = unsafe.Pointer

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
type VDSP_int24 = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vdsp_uint24
type VDSP_uint24 = unsafe.Pointer

// VDouble is a 128-bit vector packed with `double` values.
//
// See: https://developer.apple.com/documentation/Accelerate/vDouble
type VDouble = float64

// See: https://developer.apple.com/documentation/kernel/va_list
type Va_list = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vc_progress_user_options
type Vc_progress_user_options = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vector_int2
type Vector_int2 = int32

// See: https://developer.apple.com/documentation/kernel/vector_int4
type Vector_int4 = int32

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
type Vector_uint4 = uint32

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
type Vfs_rename_flags_t = uint

// See: https://developer.apple.com/documentation/kernel/vfs_roles_t
type Vfs_roles_t = unsafe.Pointer

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
type Virtual_memory_guard_exception_code_t = unsafe.Pointer

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
type Vm_address_t = Vm_offset_t

// See: https://developer.apple.com/documentation/kernel/vm_address_ut
type Vm_address_ut = Vm_address_t

// See: https://developer.apple.com/documentation/kernel/vm_behavior_t
type Vm_behavior_t = int

// See: https://developer.apple.com/documentation/kernel/vm_behavior_ut
type Vm_behavior_ut = Vm_behavior_t

// See: https://developer.apple.com/documentation/kernel/vm_extmod_statistics_data_t
type Vm_extmod_statistics_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vm_extmod_statistics_t
type Vm_extmod_statistics_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_info_object_array_t
type Vm_info_object_array_t = Vm_info_object_t

// See: https://developer.apple.com/documentation/kernel/vm_info_object_t
type Vm_info_object_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vm_info_region_64_t
type Vm_info_region_64_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vm_info_region_t
type Vm_info_region_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vm_inherit_t
type Vm_inherit_t = uint

// See: https://developer.apple.com/documentation/kernel/vm_inherit_ut
type Vm_inherit_ut = Vm_inherit_t

// See: https://developer.apple.com/documentation/kernel/vm_machine_attribute_t
type Vm_machine_attribute_t = uint

// See: https://developer.apple.com/documentation/kernel/vm_machine_attribute_val_t
type Vm_machine_attribute_val_t = int

// See: https://developer.apple.com/documentation/kernel/vm_map_address_t
type Vm_map_address_t = uint64

// See: https://developer.apple.com/documentation/kernel/vm_map_address_ut
type Vm_map_address_ut = Vm_map_address_t

// See: https://developer.apple.com/documentation/kernel/vm_map_inspect_t
type Vm_map_inspect_t = uint32

// See: https://developer.apple.com/documentation/kernel/vm_map_offset_t
type Vm_map_offset_t = uint64

// See: https://developer.apple.com/documentation/kernel/vm_map_offset_ut
type Vm_map_offset_ut = Vm_map_offset_t

// See: https://developer.apple.com/documentation/kernel/vm_map_read_t
type Vm_map_read_t = uint32

// See: https://developer.apple.com/documentation/kernel/vm_map_size_t
type Vm_map_size_t = uint64

// See: https://developer.apple.com/documentation/kernel/vm_map_size_ut
type Vm_map_size_ut = Vm_map_size_t

// See: https://developer.apple.com/documentation/kernel/vm_map_t
type Vm_map_t = uint32

// See: https://developer.apple.com/documentation/kernel/vm_named_entry_t
type Vm_named_entry_t = uint32

// See: https://developer.apple.com/documentation/kernel/vm_object_id_t
type Vm_object_id_t = uint64

// See: https://developer.apple.com/documentation/kernel/vm_object_offset_t
type Vm_object_offset_t = uint64

// See: https://developer.apple.com/documentation/kernel/vm_object_offset_ut
type Vm_object_offset_ut = Vm_object_offset_t

// See: https://developer.apple.com/documentation/kernel/vm_object_size_t
type Vm_object_size_t = uint64

// See: https://developer.apple.com/documentation/kernel/vm_object_size_ut
type Vm_object_size_ut = Vm_object_size_t

// See: https://developer.apple.com/documentation/kernel/vm_offset_t
type Vm_offset_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_offset_ut
type Vm_offset_ut = Vm_offset_t

// See: https://developer.apple.com/documentation/kernel/vm_page_info_basic_data_t
type Vm_page_info_basic_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vm_page_info_basic_t
type Vm_page_info_basic_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_page_info_data_t
type Vm_page_info_data_t = int

// See: https://developer.apple.com/documentation/kernel/vm_page_info_flavor_t
type Vm_page_info_flavor_t = int

// See: https://developer.apple.com/documentation/kernel/vm_page_info_t
type Vm_page_info_t = int

// See: https://developer.apple.com/documentation/kernel/vm_prot_t
type Vm_prot_t = int

// See: https://developer.apple.com/documentation/kernel/vm_prot_ut
type Vm_prot_ut = Vm_prot_t

// See: https://developer.apple.com/documentation/kernel/vm_purgable_t
type Vm_purgable_t = int

// See: https://developer.apple.com/documentation/kernel/vm_purgeable_info_t
type Vm_purgeable_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_purgeable_stat_t
type Vm_purgeable_stat_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vm_read_entry_t
type Vm_read_entry_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vm_region_basic_info_64_t
type Vm_region_basic_info_64_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_region_basic_info_data_64_t
type Vm_region_basic_info_data_64_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vm_region_basic_info_data_t
type Vm_region_basic_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vm_region_basic_info_t
type Vm_region_basic_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_region_extended_info_data_t
type Vm_region_extended_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vm_region_extended_info_t
type Vm_region_extended_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_region_flavor_t
type Vm_region_flavor_t = int

// See: https://developer.apple.com/documentation/kernel/vm_region_info_64_t
type Vm_region_info_64_t = int

// See: https://developer.apple.com/documentation/kernel/vm_region_info_data_t
type Vm_region_info_data_t = int

// See: https://developer.apple.com/documentation/kernel/vm_region_info_t
type Vm_region_info_t = int

// See: https://developer.apple.com/documentation/kernel/vm_region_recurse_info_64_t
type Vm_region_recurse_info_64_t = int

// See: https://developer.apple.com/documentation/kernel/vm_region_recurse_info_t
type Vm_region_recurse_info_t = int

// See: https://developer.apple.com/documentation/kernel/vm_region_submap_info_64_t
type Vm_region_submap_info_64_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_region_submap_info_data_64_t
type Vm_region_submap_info_data_64_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vm_region_submap_info_data_t
type Vm_region_submap_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vm_region_submap_info_t
type Vm_region_submap_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_region_submap_short_info_64_t
type Vm_region_submap_short_info_64_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_region_submap_short_info_data_64_t
type Vm_region_submap_short_info_data_64_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vm_region_top_info_data_t
type Vm_region_top_info_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vm_region_top_info_t
type Vm_region_top_info_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_size_struct_t
type Vm_size_struct_t = uint64

// See: https://developer.apple.com/documentation/kernel/vm_size_t
type Vm_size_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_size_ut
type Vm_size_ut = Vm_size_t

// See: https://developer.apple.com/documentation/kernel/vm_statistics64_data_t
type Vm_statistics64_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vm_statistics64_t
type Vm_statistics64_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_statistics_data_t
type Vm_statistics_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vm_statistics_t
type Vm_statistics_t = uintptr

// See: https://developer.apple.com/documentation/kernel/vm_sync_t
type Vm_sync_t = uint

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
type Vnode_verify_flags_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/voidp
type Voidp = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/voidpc
type Voidpc = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/voidpf
type Voidpf = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vol_attributes_attr_t
type Vol_attributes_attr_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vol_capabilities_attr_t
type Vol_capabilities_attr_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vol_capabilities_set_t
type Vol_capabilities_set_t = U_int32_t

// See: https://developer.apple.com/documentation/kernel/vsock_gen_t
type Vsock_gen_t = U_quad_t

// See: https://developer.apple.com/documentation/kernel/wait_interrupt_t
type Wait_interrupt_t = int

// See: https://developer.apple.com/documentation/kernel/wait_result_t
type Wait_result_t = int

// See: https://developer.apple.com/documentation/kernel/wait_timeout_urgency_t
type Wait_timeout_urgency_t = int

// See: https://developer.apple.com/documentation/kernel/wint_t
type Wint_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/x86_avx512_state_t
type X86_avx512_state_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/x86_avx_state_t
type X86_avx_state_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/x86_debug_state_t
type X86_debug_state_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/x86_exception_state32_t
type X86_exception_state32_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/x86_exception_state_t
type X86_exception_state_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/x86_float_state32_t
type X86_float_state32_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/x86_float_state_t
type X86_float_state_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/x86_state_hdr_t
type X86_state_hdr_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/x86_thread_state32_t
type X86_thread_state32_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/x86_thread_state_t
type X86_thread_state_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/xattrname
type Xattrname = int8

// See: https://developer.apple.com/documentation/kernel/xcred
type Xcred = uint32

// See: https://developer.apple.com/documentation/kernel/xdrbuf_type
type Xdrbuf_type = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/xmldata_t
type XmlData_t = int8

// See: https://developer.apple.com/documentation/kernel/z_stream
type Z_stream = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/z_streamp
type Z_streamp = Z_stream

// See: https://developer.apple.com/documentation/kernel/zone_btrecord_array_t
type Zone_btrecord_array_t = Zone_btrecord_t

// See: https://developer.apple.com/documentation/kernel/zone_btrecord_t
type Zone_btrecord_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/zone_info_array_t
type Zone_info_array_t = Zone_info_t

// See: https://developer.apple.com/documentation/kernel/zone_info_t
type Zone_info_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/zone_name_array_t
type Zone_name_array_t = Zone_name_t

// See: https://developer.apple.com/documentation/kernel/zone_name_t
type Zone_name_t = unsafe.Pointer

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

// GssdMechtype is a Go-name alias for Gssd_mechtype.
type GssdMechtype = Gssd_mechtype

// GssdNametype is a Go-name alias for Gssd_nametype.
type GssdNametype = Gssd_nametype

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

// HvgHcallCode is a Go-name alias for Hvg_hcall_code_t.
type HvgHcallCode = Hvg_hcall_code_t

// HvgHcallOutput is a Go-name alias for Hvg_hcall_output_t.
type HvgHcallOutput = Hvg_hcall_output_t

// HvgHcallReturn is a Go-name alias for Hvg_hcall_return_t.
type HvgHcallReturn = Hvg_hcall_return_t

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

// KdebugFlags is a Go-name alias for Kdebug_flags_t.
type KdebugFlags = Kdebug_flags_t

// KdebugLiveFlags is a Go-name alias for Kdebug_live_flags_t.
type KdebugLiveFlags = Kdebug_live_flags_t

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

// Key is a Go-name alias for Key_t.
type Key = Key_t

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

// MachMsgAuditTrailer is a Go-name alias for Mach_msg_audit_trailer_t.
type MachMsgAuditTrailer = Mach_msg_audit_trailer_t

// MachMsgBase is a Go-name alias for Mach_msg_base_t.
type MachMsgBase = Mach_msg_base_t

// MachMsgBits is a Go-name alias for Mach_msg_bits_t.
type MachMsgBits = Mach_msg_bits_t

// MachMsgBody is a Go-name alias for Mach_msg_body_t.
type MachMsgBody = Mach_msg_body_t

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

// MachMsgOolDescriptor is a Go-name alias for Mach_msg_ool_descriptor_t.
type MachMsgOolDescriptor = Mach_msg_ool_descriptor_t

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

// MachMsgPortDescriptor is a Go-name alias for Mach_msg_port_descriptor_t.
type MachMsgPortDescriptor = Mach_msg_port_descriptor_t

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

// MachMsgTrailer is a Go-name alias for Mach_msg_trailer_t.
type MachMsgTrailer = Mach_msg_trailer_t

// MachMsgTrailerType is a Go-name alias for Mach_msg_trailer_type_t.
type MachMsgTrailerType = Mach_msg_trailer_type_t

// MachMsgTypeDescriptor is a Go-name alias for Mach_msg_type_descriptor_t.
type MachMsgTypeDescriptor = Mach_msg_type_descriptor_t

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

// MachTimebaseInfo is a Go-name alias for Mach_timebase_info_t.
type MachTimebaseInfo = Mach_timebase_info_t

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

// MachVmRangeFlags is a Go-name alias for Mach_vm_range_flags_t.
type MachVmRangeFlags = Mach_vm_range_flags_t

// MachVmRangeRecipe is a Go-name alias for Mach_vm_range_recipe_t.
type MachVmRangeRecipe = Mach_vm_range_recipe_t

// MachVmRangeRecipeV1 is a Go-name alias for Mach_vm_range_recipe_v1_t.
type MachVmRangeRecipeV1 = Mach_vm_range_recipe_v1_t

// MachVmRangeRecipeV1Ut is a Go-name alias for Mach_vm_range_recipe_v1_ut.
type MachVmRangeRecipeV1Ut = Mach_vm_range_recipe_v1_ut

// MachVmRangeRecipesRaw is a Go-name alias for Mach_vm_range_recipes_raw_t.
type MachVmRangeRecipesRaw = Mach_vm_range_recipes_raw_t

// MachVmRangeTag is a Go-name alias for Mach_vm_range_tag_t.
type MachVmRangeTag = Mach_vm_range_tag_t

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

// MicrostackshotFlags is a Go-name alias for Microstackshot_flags_t.
type MicrostackshotFlags = Microstackshot_flags_t

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

// SockaddrBptr is a Go-name alias for Sockaddr_bptr_t.
type SockaddrBptr = Sockaddr_bptr_t

// SockaddrPtrRef is a Go-name alias for Sockaddr_ptr_ref_t.
type SockaddrPtrRef = Sockaddr_ptr_ref_t

// SockaddrPtr is a Go-name alias for Sockaddr_ptr_t.
type SockaddrPtr = Sockaddr_ptr_t

// SockaddrRefPtr is a Go-name alias for Sockaddr_ref_ptr_t.
type SockaddrRefPtr = Sockaddr_ref_ptr_t

// SockaddrRefRef is a Go-name alias for Sockaddr_ref_ref_t.
type SockaddrRefRef = Sockaddr_ref_ref_t

// SockaddrRef is a Go-name alias for Sockaddr_ref_t.
type SockaddrRef = Sockaddr_ref_t

// SockaddrStorageBptr is a Go-name alias for Sockaddr_storage_bptr_t.
type SockaddrStorageBptr = Sockaddr_storage_bptr_t

// SockaddrStoragePtrRef is a Go-name alias for Sockaddr_storage_ptr_ref_t.
type SockaddrStoragePtrRef = Sockaddr_storage_ptr_ref_t

// SockaddrStoragePtr is a Go-name alias for Sockaddr_storage_ptr_t.
type SockaddrStoragePtr = Sockaddr_storage_ptr_t

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

// StackshotFlags is a Go-name alias for Stackshot_flags_t.
type StackshotFlags = Stackshot_flags_t

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

// VirtualMemoryGuardExceptionCode is a Go-name alias for Virtual_memory_guard_exception_code_t.
type VirtualMemoryGuardExceptionCode = Virtual_memory_guard_exception_code_t

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

// NXByteOrder is the canonical enum type used by byte-order helpers.
type NXByteOrder = Nx

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
