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
type IOBlitMemoryRef uintptr

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
type IODispatchAction = func() int32

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
type IOFireWireSessionRef uintptr

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
type IOStateNotificationHandler = func() int32

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

// See: https://developer.apple.com/documentation/kernel/iousbhostcidoorbell
type IOUSBHostCIDoorbell = uint32

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
type NuDCLReceivePacketRef uintptr

// See: https://developer.apple.com/documentation/kernel/nudclref
type NuDCLRef uintptr

// See: https://developer.apple.com/documentation/kernel/nudclsendpacketref
type NuDCLSendPacketRef uintptr

// See: https://developer.apple.com/documentation/kernel/nudclskipcycleref
type NuDCLSkipCycleRef uintptr

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
type OSAsyncReference = unsafe.Pointer

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
type Io_async_ref_t = unsafe.Pointer

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
type Ledger_item_t = unsafe.Pointer

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
type Mach_msg_timeout_t = unsafe.Pointer

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
type Mach_msg_type_number_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_msg_type_size_t
type Mach_msg_type_size_t = unsafe.Pointer

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
type Mach_port_mscount_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_port_msgcount_t
type Mach_port_msgcount_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_port_name_array_t
type Mach_port_name_array_t = Mach_port_name_t

// See: https://developer.apple.com/documentation/kernel/mach_port_name_t
type Mach_port_name_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_port_options_ptr_t
type Mach_port_options_ptr_t = Mach_port_options_t

// See: https://developer.apple.com/documentation/kernel/mach_port_options_t
type Mach_port_options_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_port_qos_t
type Mach_port_qos_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_port_right_t
type Mach_port_right_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_port_rights_t
type Mach_port_rights_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_port_seqno_t
type Mach_port_seqno_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_port_srights_t
type Mach_port_srights_t = uint

// See: https://developer.apple.com/documentation/kernel/mach_port_status_t
type Mach_port_status_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_port_t
type Mach_port_t = uint32

// See: https://developer.apple.com/documentation/kernel/mach_port_type_array_t
type Mach_port_type_array_t = Mach_port_type_t

// See: https://developer.apple.com/documentation/kernel/mach_port_type_t
type Mach_port_type_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/mach_port_urefs_t
type Mach_port_urefs_t = unsafe.Pointer

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
type Memory_object_cluster_size_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/memory_object_control_t
type Memory_object_control_t = uint32

// See: https://developer.apple.com/documentation/kernel/memory_object_copy_strategy_t
type Memory_object_copy_strategy_t = int

// See: https://developer.apple.com/documentation/kernel/memory_object_default_t
type Memory_object_default_t = uint32

// See: https://developer.apple.com/documentation/kernel/memory_object_fault_info_t
type Memory_object_fault_info_t = unsafe.Pointer

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

// See: https://developer.apple.com/documentation/kernel/tusblpmexitlatency
type TUSBLPMExitLatency = unsafe.Pointer

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
type Task_flavor_t = unsafe.Pointer

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
type Task_inspect_flavor_t = unsafe.Pointer

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
type Task_policy_flavor_t = unsafe.Pointer

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
type Thread_flavor_t = unsafe.Pointer

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
type Thread_policy_flavor_t = unsafe.Pointer

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
type Thread_state_data_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/thread_state_flavor_array_t
type Thread_state_flavor_array_t = Thread_state_flavor_t

// See: https://developer.apple.com/documentation/kernel/thread_state_flavor_t
type Thread_state_flavor_t = int

// See: https://developer.apple.com/documentation/kernel/thread_state_t
type Thread_state_t = unsafe.Pointer

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
type VDouble = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/va_list
type Va_list = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vc_progress_user_options
type Vc_progress_user_options = unsafe.Pointer

// See: https://developer.apple.com/documentation/kernel/vector_int2
type Vector_int2 = int

// See: https://developer.apple.com/documentation/kernel/vector_int4
type Vector_int4 = int

// See: https://developer.apple.com/documentation/kernel/vector_int8
type Vector_int8 = int

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

// NXByteOrder is the canonical enum type used by byte-order helpers.

type NXByteOrder = Nx

// Uio_rw aliases the generated uio enum family for vnode read/write APIs.

type Uio_rw = Uio

// Uio_seg aliases the generated uio enum family for vnode segment-space APIs.

type Uio_seg = Uio

// Ifnet_interface_advisory_interface_type aliases the generated advisory interface type enum.

type Ifnet_interface_advisory_interface_type = IfInterfaceAdvisoryInterfaceType

// IfnetInterfaceAdvisoryInterfaceType aliases the generated advisory interface type enum.

type IfnetInterfaceAdvisoryInterfaceType = IfInterfaceAdvisoryInterfaceType

// Ifnet_interface_advisory_version is a uint8-backed advisory version enum.

type Ifnet_interface_advisory_version = uint8
