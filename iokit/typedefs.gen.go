// Code generated from Apple documentation. DO NOT EDIT.

package iokit

import (
	"unsafe"

	"github.com/tmc/apple/applicationservices"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/kernel"
)

// See: https://developer.apple.com/documentation/iokit/ataoperationtype
type ATAOperationType = uint32

// See: https://developer.apple.com/documentation/iokit/atasmartdata
type ATASMARTData = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/atasmartdatathresholds
type ATASMARTDataThresholds = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/atasmartlogdirectory
type ATASMARTLogDirectory = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/atasmartlogentry
type ATASMARTLogEntry = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/avidtype
type AVIDType = uint32

// See: https://developer.apple.com/documentation/iokit/bddiscinfo
type BDDiscInfo = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/bdfeatures
type BDFeatures = uint32

// See: https://developer.apple.com/documentation/iokit/bdmediatype
type BDMediaType = uint32

// See: https://developer.apple.com/documentation/iokit/bdtrackinfo
type BDTrackInfo = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/block0
type Block0 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/cdatip
type CDATIP = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/cdaudiostatus
type CDAudioStatus = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/cddiscinfo
type CDDiscInfo = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/cdfeatures
type CDFeatures = uint32

// See: https://developer.apple.com/documentation/iokit/cdisrc
type CDISRC = int8

// See: https://developer.apple.com/documentation/iokit/cdmcn
type CDMCN = int8

// See: https://developer.apple.com/documentation/iokit/cdmsf
type CDMSF = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/cdmediatype
type CDMediaType = uint32

// See: https://developer.apple.com/documentation/iokit/cdpma
type CDPMA = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/cdpmadescriptor
type CDPMADescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/cdsectorarea
type CDSectorArea = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/cdsectorsize
type CDSectorSize = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/cdsectortype
type CDSectorType = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/cdtext
type CDTEXT = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/cdtextdescriptor
type CDTEXTDescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/cdtoc
type CDTOC = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/cdtocdescriptor
type CDTOCDescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/cdtocformat
type CDTOCFormat = uint8

// See: https://developer.apple.com/documentation/iokit/cdtrackinfo
type CDTrackInfo = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/cdtrackinfoaddresstype
type CDTrackInfoAddressType = uint8

// See: https://developer.apple.com/documentation/iokit/csrnodeuniqueid
type CSRNodeUniqueID = uint64

// See: https://developer.apple.com/documentation/iokit/colorspec
type ColorSpec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/colorspecptr
type ColorSpecPtr = applicationservices.ColorSpec

// See: https://developer.apple.com/documentation/iokit/dasdmodeparameterblockdescriptor
type DASDModeParameterBlockDescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dclcallcommandproc
type DCLCallCommandProc = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dclcallcommandprocptr
type DCLCallCommandProcPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/dclcallproc
type DCLCallProc = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dclcallprocdatatype
type DCLCallProcDataType = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/dclcallprocptr
type DCLCallProcPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/dclcommand
type DCLCommand = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dclcommandptr
type DCLCommandPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/dclcompilerdatatype
type DCLCompilerDataType = uint32

// See: https://developer.apple.com/documentation/iokit/dcljump
type DCLJump = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dcljumpptr
type DCLJumpPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/dcllabel
type DCLLabel = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dcllabelptr
type DCLLabelPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/dclnudclleader
type DCLNuDCLLeader = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dclptrtimestamp
type DCLPtrTimeStamp = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dclptrtimestampptr
type DCLPtrTimeStampPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/dclsettagsyncbits
type DCLSetTagSyncBits = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dclsettagsyncbitsptr
type DCLSetTagSyncBitsPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/dcltimestamp
type DCLTimeStamp = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dcltimestampptr
type DCLTimeStampPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/dcltransferbuffer
type DCLTransferBuffer = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dcltransferbufferptr
type DCLTransferBufferPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/dcltransferpacket
type DCLTransferPacket = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dcltransferpacketptr
type DCLTransferPacketPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/dclupdatedcllist
type DCLUpdateDCLList = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dclupdatedcllistptr
type DCLUpdateDCLListPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/ddmap
type DDMap = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dpme
type DPME = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dvdauthenticationgrantidinfo
type DVDAuthenticationGrantIDInfo = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dvdauthenticationsuccessflaginfo
type DVDAuthenticationSuccessFlagInfo = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dvdbooktype
type DVDBookType = uint8

// See: https://developer.apple.com/documentation/iokit/dvdcprmregioncode
type DVDCPRMRegionCode = uint8

// See: https://developer.apple.com/documentation/iokit/dvdchallengekeyinfo
type DVDChallengeKeyInfo = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dvdcopyrightinfo
type DVDCopyrightInfo = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dvddiscinfo
type DVDDiscInfo = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dvddisckeyinfo
type DVDDiscKeyInfo = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dvdfeatures
type DVDFeatures = uint32

// See: https://developer.apple.com/documentation/iokit/dvdkey1info
type DVDKey1Info = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dvdkey2info
type DVDKey2Info = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dvdkeyclass
type DVDKeyClass = uint8

// See: https://developer.apple.com/documentation/iokit/dvdkeyformat
type DVDKeyFormat = uint8

// See: https://developer.apple.com/documentation/iokit/dvdmanufacturinginfo
type DVDManufacturingInfo = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dvdmediatype
type DVDMediaType = uint32

// See: https://developer.apple.com/documentation/iokit/dvdphysicalformatinfo
type DVDPhysicalFormatInfo = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dvdrzoneinfo
type DVDRZoneInfo = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dvdrzoneinfoaddresstype
type DVDRZoneInfoAddressType = uint8

// See: https://developer.apple.com/documentation/iokit/dvdregionplaybackcontrolinfo
type DVDRegionPlaybackControlInfo = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dvdregionalplaybackcontrolscheme
type DVDRegionalPlaybackControlScheme = uint8

// See: https://developer.apple.com/documentation/iokit/dvdstructureformat
type DVDStructureFormat = uint8

// See: https://developer.apple.com/documentation/iokit/dvdtitlekeyinfo
type DVDTitleKeyInfo = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/depthmode
type DepthMode = uint16

// See: https://developer.apple.com/documentation/iokit/displayidtype
type DisplayIDType = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/displaymodeid
type DisplayModeID = string

// See: https://developer.apple.com/documentation/iokit/evcmd
type EvCmd = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/evglobals
type EvGlobals = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/evoffsets
type EvOffsets = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/extendedsensecode
type ExtendedSenseCode = uint8

// See: https://developer.apple.com/documentation/iokit/fwaddress
type FWAddress = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/fwaddressptr
type FWAddressPtr = uintptr

// See: https://developer.apple.com/documentation/iokit/fwaddressspaceflags
type FWAddressSpaceFlags = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/fwclientcommandid
type FWClientCommandID = string

// See: https://developer.apple.com/documentation/iokit/fwsbp2logincompleteparams
type FWSBP2LoginCompleteParams = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/fwsbp2loginresponse
type FWSBP2LoginResponse = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/fwsbp2logoutcompleteparams
type FWSBP2LogoutCompleteParams = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/fwsbp2notifyparams
type FWSBP2NotifyParams = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/fwsbp2reconnectparams
type FWSBP2ReconnectParams = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/fwsbp2statusblock
type FWSBP2StatusBlock = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/fwsbp2virtualrange
type FWSBP2VirtualRange = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/gammatableid
type GammaTableID = uint32

// See: https://developer.apple.com/documentation/iokit/gammatbl
type GammaTbl = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/gammatblptr
type GammaTblPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/hidreportcommandtype
type HIDReportCommandType = kernel.Pointer

// IOATASMARTInterface is self-Monitoring, Analysis, and Reporting Technology Interface.
//
// See: https://developer.apple.com/documentation/iokit/ioatasmartinterface
type IOATASMARTInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioavccommandresponse
type IOAVCCommandResponse = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioavcframefields
type IOAVCFrameFields = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioavcopcodes
type IOAVCOpcodes = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioavcunittypes
type IOAVCUnitTypes = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioaccelbounds
type IOAccelBounds = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioacceldeviceregion
type IOAccelDeviceRegion = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioaccelid
type IOAccelID = int32

// See: https://developer.apple.com/documentation/iokit/ioaccelsize
type IOAccelSize = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioaccelsurfaceinformation
type IOAccelSurfaceInformation = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioaccelsurfacereaddata
type IOAccelSurfaceReadData = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioaccelsurfacescaling
type IOAccelSurfaceScaling = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioaddressrange
type IOAddressRange = IOVirtualRange

// See: https://developer.apple.com/documentation/iokit/ioalignment
type IOAlignment = uint32

// See: https://developer.apple.com/documentation/iokit/ioappletimingid
type IOAppleTimingID = uint32

// IOAsyncCallback is standard callback function for asynchronous I/O requests with lots of extra arguments beyond a refcon and result code.
//
// See: https://developer.apple.com/documentation/iokit/ioasynccallback
type IOAsyncCallback = func(unsafe.Pointer, int, unsafe.Pointer, uint32)

// IOAsyncCallback0 is standard callback function for asynchronous I/O requests with no extra arguments beyond a refcon and result code.
//
// See: https://developer.apple.com/documentation/iokit/ioasynccallback0
type IOAsyncCallback0 = func(unsafe.Pointer, int)

// IOAsyncCallback1 is standard callback function for asynchronous I/O requests with one extra argument beyond a refcon and result code. This is often a count of the number of bytes transferred.
//
// See: https://developer.apple.com/documentation/iokit/ioasynccallback1
type IOAsyncCallback1 = func(unsafe.Pointer, int, unsafe.Pointer)

// IOAsyncCallback2 is standard callback function for asynchronous I/O requests with two extra arguments beyond a refcon and result code.
//
// See: https://developer.apple.com/documentation/iokit/ioasynccallback2
type IOAsyncCallback2 = func(unsafe.Pointer, int, unsafe.Pointer, unsafe.Pointer)

// See: https://developer.apple.com/documentation/iokit/ioaudiobufferdatadescriptor
type IOAudioBufferDataDescriptor = kernel.Pointer

// IOAudioControlCalls is the set of constants passed to IOAudioControlUserClient::getExternalMethodForIndex() when making calls from the IOAudioFamily user client code.
//
// See: https://developer.apple.com/documentation/iokit/ioaudiocontrolcalls
type IOAudioControlCalls = kernel.Pointer

// IOAudioControlNotifications is the set of constants passed in the type field of IOAudioControlUserClient::registerNotificaitonPort().
//
// See: https://developer.apple.com/documentation/iokit/ioaudiocontrolnotifications
type IOAudioControlNotifications = kernel.Pointer

// IOAudioEngineCalls is the set of constants passed to IOAudioEngineUserClient::getExternalMethodForIndex() when making calls from the IOAudioFamily user client code.
//
// See: https://developer.apple.com/documentation/iokit/ioaudioenginecalls
type IOAudioEngineCalls = kernel.Pointer

// IOAudioEngineMemory is used to identify the type of memory requested by a client process to be mapped into its process space.
//
// See: https://developer.apple.com/documentation/iokit/ioaudioenginememory
type IOAudioEngineMemory = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioaudioenginenotifications
type IOAudioEngineNotifications = kernel.Pointer

// IOAudioEngineState is represents the state of an IOAudioEngine.
//
// See: https://developer.apple.com/documentation/iokit/ioaudioenginestate
type IOAudioEngineState = kernel.Pointer

// IOAudioEngineStatus is shared-memory structure giving audio engine status.
//
// See: https://developer.apple.com/documentation/iokit/ioaudioenginestatus
type IOAudioEngineStatus = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioaudioenginetraps
type IOAudioEngineTraps = kernel.Pointer

// IOAudioNotificationMessage is used in the mach message for IOAudio notifications.
//
// See: https://developer.apple.com/documentation/iokit/ioaudionotificationmessage
type IOAudioNotificationMessage = kernel.Pointer

// IOAudioSMPTETime is a structure for holding a SMPTE time.
//
// See: https://developer.apple.com/documentation/iokit/ioaudiosmptetime
type IOAudioSMPTETime = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioaudiosampleintervaldescriptor
type IOAudioSampleIntervalDescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioaudiosamplerate
type IOAudioSampleRate = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioaudiostreamdatadescriptor
type IOAudioStreamDataDescriptor = kernel.Pointer

// IOAudioStreamDirection is represents the direction of an IOAudioStream.
//
// See: https://developer.apple.com/documentation/iokit/ioaudiostreamdirection
type IOAudioStreamDirection = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioaudiostreamformat
type IOAudioStreamFormat = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioaudiostreamformatextension
type IOAudioStreamFormatExtension = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioaudiotimestamp
type IOAudioTimeStamp = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioblitcompletiontoken
type IOBlitCompletionToken = int32

// See: https://developer.apple.com/documentation/iokit/ioblitcopyrectangle
type IOBlitCopyRectangle = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioblitcopyrectangles
type IOBlitCopyRectangles = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioblitcopyregion
type IOBlitCopyRegion = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioblitcursor
type IOBlitCursor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioblitmemory
type IOBlitMemory = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/ioblitmemoryref
type IOBlitMemoryRef uintptr

// See: https://developer.apple.com/documentation/iokit/ioblitoperation
type IOBlitOperation = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioblitrectangle
type IOBlitRectangle = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioblitrectangles
type IOBlitRectangles = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioblitscanlines
type IOBlitScanlines = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioblitsourcedesttype
type IOBlitSourceDestType = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/ioblitsourcetype
type IOBlitSourceType = uint32

// See: https://developer.apple.com/documentation/iokit/ioblitsurface
type IOBlitSurface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioblittype
type IOBlitType = uint32

// See: https://developer.apple.com/documentation/iokit/ioblitvertex
type IOBlitVertex = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioblitvertices
type IOBlitVertices = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iobytecount
type IOByteCount = uint

// See: https://developer.apple.com/documentation/iokit/iobytecount32
type IOByteCount32 = uint32

// See: https://developer.apple.com/documentation/iokit/iobytecount64
type IOByteCount64 = uint64

// See: https://developer.apple.com/documentation/iokit/iocfplugininterfacestruct
type IOCFPlugInInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iocsrkeytype
type IOCSRKeyType = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iocachemode
type IOCacheMode = uint32

// See: https://developer.apple.com/documentation/iokit/iocolorcomponent
type IOColorComponent = uint16

// See: https://developer.apple.com/documentation/iokit/iocolorentry
type IOColorEntry = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioconfigkeytype
type IOConfigKeyType = kernel.Pointer

// IODataQueueAppendix is a struct mapping to the appendix region of a data queue.
//
// See: https://developer.apple.com/documentation/iokit/iodataqueueappendix
type IODataQueueAppendix = unsafe.Pointer

// IODataQueueEntry is represents an entry within the data queue.
//
// See: https://developer.apple.com/documentation/iokit/iodataqueueentry
type IODataQueueEntry = unsafe.Pointer

// IODataQueueMemory is a struct mapping to the header region of a data queue.
//
// See: https://developer.apple.com/documentation/iokit/iodataqueuememory
type IODataQueueMemory = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iodetailedtiminginformation
type IODetailedTimingInformation = IODetailedTimingInformationV2

// See: https://developer.apple.com/documentation/iokit/iodetailedtiminginformationv1
type IODetailedTimingInformationV1 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iodetailedtiminginformationv2
type IODetailedTimingInformationV2 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iodevicenumber
type IODeviceNumber = uint32

// See: https://developer.apple.com/documentation/iokit/iodisplaymodeinformation
type IODisplayModeInformation = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iodisplayproductid
type IODisplayProductID = uint32

// See: https://developer.apple.com/documentation/iokit/iodisplayscalerinformation
type IODisplayScalerInformation = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iodisplaytimingrange
type IODisplayTimingRange = IODisplayTimingRangeV2

// See: https://developer.apple.com/documentation/iokit/iodisplaytimingrangev1
type IODisplayTimingRangeV1 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iodisplaytimingrangev2
type IODisplayTimingRangeV2 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iodisplayvendorid
type IODisplayVendorID = uint32

// See: https://developer.apple.com/documentation/iokit/iodot3collentry
type IODot3CollEntry = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iodot3rxextraentry
type IODot3RxExtraEntry = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iodot3statsentry
type IODot3StatsEntry = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iodot3txextraentry
type IODot3TxExtraEntry = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioethernetstats
type IOEthernetStats = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofbdplinkconfig
type IOFBDPLinkConfig = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofbdisplaymodedescription
type IOFBDisplayModeDescription = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofbhdrmetadatav1
type IOFBHDRMetaDataV1 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofwavcasynccommandstate
type IOFWAVCAsyncCommandState = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofwavcplugtypes
type IOFWAVCPlugTypes = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofwavcsubunitplugmessages
type IOFWAVCSubunitPlugMessages = kernel.Pointer

// IOFWAsyncStreamListenerInterface is represents and provides management functions for a asyn stream listener object.
//
// See: https://developer.apple.com/documentation/iokit/iofwasyncstreamlistenerinterface
type IOFWAsyncStreamListenerInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofwasyncstreamlistenerinterfaceref
type IOFWAsyncStreamListenerInterfaceRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofwdclnotificationtype
type IOFWDCLNotificationType = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofwisochportoptions
type IOFWIsochPortOptions = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofwisochresourceflags
type IOFWIsochResourceFlags = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofwspeed
type IOFWSpeed = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofirewireavclibasynchronouscommand
type IOFireWireAVCLibAsynchronousCommand = kernel.Pointer

// IOFireWireAVCLibConsumerInterface is interface for an asynchronous connection consumer.
//
// See: https://developer.apple.com/documentation/iokit/iofirewireavclibconsumerinterface
type IOFireWireAVCLibConsumerInterface = kernel.Pointer

// IOFireWireAVCLibProtocolInterface is initial interface discovered for all AVC protocol drivers.
//
// See: https://developer.apple.com/documentation/iokit/iofirewireavclibprotocolinterface
type IOFireWireAVCLibProtocolInterface = kernel.Pointer

// IOFireWireAVCLibUnitInterface is initial interface discovered for all AVC Unit drivers.
//
// See: https://developer.apple.com/documentation/iokit/iofirewireavclibunitinterface
type IOFireWireAVCLibUnitInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofirewireasyncstreamcommandinterface
type IOFireWireAsyncStreamCommandInterface = kernel.Pointer

// IOFireWireCommandInterface is iOFireWireLib command object.
//
// See: https://developer.apple.com/documentation/iokit/iofirewirecommandinterface
type IOFireWireCommandInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofirewirecompareswapcommandinterface
type IOFireWireCompareSwapCommandInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofirewirecompareswapcommandinterface_v3
type IOFireWireCompareSwapCommandInterface_v3 = kernel.Pointer

// IOFireWireConfigDirectoryInterface is iOFireWireLib device config ROM browsing interface.
//
// See: https://developer.apple.com/documentation/iokit/iofirewireconfigdirectoryinterface
type IOFireWireConfigDirectoryInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofirewiredclcommandpoolinterface
type IOFireWireDCLCommandPoolInterface = kernel.Pointer

// IOFireWireDeviceInterface is iOFireWireDeviceInterface is your primary gateway to the functionality contained in IOFireWireLib.
//
// See: https://developer.apple.com/documentation/iokit/iofirewiredeviceinterface
type IOFireWireDeviceInterface = kernel.Pointer

// IOFireWireIsochChannelInterface is fireWire user client isochronous channel object.
//
// See: https://developer.apple.com/documentation/iokit/iofirewireisochchannelinterface
type IOFireWireIsochChannelInterface = kernel.Pointer

// IOFireWireIsochPortInterface is fireWire user client isochronous port interface.
//
// See: https://developer.apple.com/documentation/iokit/iofirewireisochportinterface
type IOFireWireIsochPortInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofirewirelibasyncstreamcommandref
type IOFireWireLibAsyncStreamCommandRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewirelibcommandref
type IOFireWireLibCommandRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewirelibcompareswapcommandref
type IOFireWireLibCompareSwapCommandRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewirelibcompareswapcommandv3ref
type IOFireWireLibCompareSwapCommandV3Ref uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewirelibconfigdirectoryref
type IOFireWireLibConfigDirectoryRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewirelibdclcommandpoolref
type IOFireWireLibDCLCommandPoolRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewirelibdeviceref
type IOFireWireLibDeviceRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewirelibirmallocationinterface
type IOFireWireLibIRMAllocationInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofirewirelibirmallocationref
type IOFireWireLibIRMAllocationRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewirelibisochchannelref
type IOFireWireLibIsochChannelRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewirelibisochportref
type IOFireWireLibIsochPortRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewireliblocalisochportref
type IOFireWireLibLocalIsochPortRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewireliblocalunitdirectoryref
type IOFireWireLibLocalUnitDirectoryRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewirelibnudclpoolref
type IOFireWireLibNuDCLPoolRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewirelibnubref
type IOFireWireLibNubRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewirelibphycommandref
type IOFireWireLibPHYCommandRef uintptr

// IOFireWireLibPHYPacketListenerInterface is represents and provides management functions for a phy packet listener object.
//
// See: https://developer.apple.com/documentation/iokit/iofirewirelibphypacketlistenerinterface
type IOFireWireLibPHYPacketListenerInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofirewirelibphypacketlistenerref
type IOFireWireLibPHYPacketListenerRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewirelibphysicaladdressspaceref
type IOFireWireLibPhysicalAddressSpaceRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewirelibpseudoaddressspaceref
type IOFireWireLibPseudoAddressSpaceRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewirelibreadcommandref
type IOFireWireLibReadCommandRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewirelibreadquadletcommandref
type IOFireWireLibReadQuadletCommandRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewirelibremoteisochportref
type IOFireWireLibRemoteIsochPortRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewirelibunitref
type IOFireWireLibUnitRef uintptr

// IOFireWireLibVectorCommandInterface is iOFireWireLib command object for grouping commands execution.
//
// See: https://developer.apple.com/documentation/iokit/iofirewirelibvectorcommandinterface
type IOFireWireLibVectorCommandInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofirewirelibvectorcommandref
type IOFireWireLibVectorCommandRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewirelibwritecommandref
type IOFireWireLibWriteCommandRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewirelibwritequadletcommandref
type IOFireWireLibWriteQuadletCommandRef uintptr

// IOFireWireLocalIsochPortInterface is fireWire user client local isochronous port object.
//
// See: https://developer.apple.com/documentation/iokit/iofirewirelocalisochportinterface
type IOFireWireLocalIsochPortInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofirewirelocalunitdirectoryinterface
type IOFireWireLocalUnitDirectoryInterface = kernel.Pointer

// IOFireWireNuDCLPoolInterface is use this interface to build NuDCL-based DCL programs.
//
// See: https://developer.apple.com/documentation/iokit/iofirewirenudclpoolinterface
type IOFireWireNuDCLPoolInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofirewirenubinterface
type IOFireWireNubInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofirewirephycommandinterface
type IOFireWirePHYCommandInterface = kernel.Pointer

// IOFireWirePhysicalAddressSpaceInterface is iOFireWireLib physical address space object. ( interface name: IOFireWirePhysicalAddressSpaceInterface ).
//
// See: https://developer.apple.com/documentation/iokit/iofirewirephysicaladdressspaceinterface
type IOFireWirePhysicalAddressSpaceInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofirewirepseudoaddressspaceinterface
type IOFireWirePseudoAddressSpaceInterface = kernel.Pointer

// IOFireWireReadCommandInterface is iOFireWireLib block read command object.
//
// See: https://developer.apple.com/documentation/iokit/iofirewirereadcommandinterface
type IOFireWireReadCommandInterface = kernel.Pointer

// IOFireWireReadQuadletCommandInterface is iOFireWireReadQuadletCommandInterface -- IOFireWireLib quadlet read command object.
//
// See: https://developer.apple.com/documentation/iokit/iofirewirereadquadletcommandinterface
type IOFireWireReadQuadletCommandInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofirewireremoteisochportinterface
type IOFireWireRemoteIsochPortInterface = kernel.Pointer

// IOFireWireSBP2LibLUNInterface is initial interface disovered for all drivers.
//
// See: https://developer.apple.com/documentation/iokit/iofirewiresbp2libluninterface
type IOFireWireSBP2LibLUNInterface = kernel.Pointer

// IOFireWireSBP2LibLoginInterface is supplies the login maintenance and Normal Command ORB execution portions of the API.
//
// See: https://developer.apple.com/documentation/iokit/iofirewiresbp2liblogininterface
type IOFireWireSBP2LibLoginInterface = kernel.Pointer

// IOFireWireSBP2LibMgmtORBInterface is supplies non login related management ORBs. Management ORBs can be executed independent of a login, if necessary. Management ORBs are created using the IOFireWireSBP2LibLUNInterface.
//
// See: https://developer.apple.com/documentation/iokit/iofirewiresbp2libmgmtorbinterface
type IOFireWireSBP2LibMgmtORBInterface = kernel.Pointer

// IOFireWireSBP2LibORBInterface is represents an SBP2 normal command ORB. Supplies the APIs for configuring normal command ORBs. This includes setting the command block and writing the page tables for I/O. The ORBs are executed using the submitORB method in IOFireWireSBP2LibLoginInterface.
//
// See: https://developer.apple.com/documentation/iokit/iofirewiresbp2liborbinterface
type IOFireWireSBP2LibORBInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofirewiresessionref
type IOFireWireSessionRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewireunitinterface
type IOFireWireUnitInterface = kernel.Pointer

// IOFireWireWriteCommandInterface is iOFireWireLib block read command object.
//
// See: https://developer.apple.com/documentation/iokit/iofirewirewritecommandinterface
type IOFireWireWriteCommandInterface = kernel.Pointer

// IOFireWireWriteQuadletCommandInterface is iOFireWireLib quadlet read command object.
//
// See: https://developer.apple.com/documentation/iokit/iofirewirewritequadletcommandinterface
type IOFireWireWriteQuadletCommandInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofixed
type IOFixed = int32

// See: https://developer.apple.com/documentation/iokit/iofixed1616
type IOFixed1616 = uint32

// See: https://developer.apple.com/documentation/iokit/iofixedpoint32
type IOFixedPoint32 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iofourcharcode
type IOFourCharCode = uint32

// See: https://developer.apple.com/documentation/iokit/ioframebufferinformation
type IOFramebufferInformation = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iogbounds
type IOGBounds = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iogpoint
type IOGPoint = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iogsize
type IOGSize = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iographicsacceleratorinterface
type IOGraphicsAcceleratorInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iohidaccelerationalgorithmtype
type IOHIDAccelerationAlgorithmType = uint8

// See: https://developer.apple.com/documentation/iokit/iohidaccesstype
type IOHIDAccessType = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iohidbuttonmodes
type IOHIDButtonModes = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iohidcompletion
type IOHIDCompletion = kernel.Pointer

// IOHIDDeviceDeviceInterface is the object you use to access HID devices from user space, returned by version 1.5 of the IOHIDFamily.
//
// See: https://developer.apple.com/documentation/iokit/iohiddevicedeviceinterface
type IOHIDDeviceDeviceInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iohiddevicegetvalueoptions
type IOHIDDeviceGetValueOptions = kernel.Pointer

// IOHIDDeviceInterface is cFPlugin object subclass which provides the primary interface to HID devices.
//
// See: https://developer.apple.com/documentation/iokit/iohiddeviceinterface
type IOHIDDeviceInterface = kernel.Pointer

// IOHIDDeviceInterface121 is cFPlugin object subclass which provides the primary interface to HID devices. This class is a subclass of IOHIDDeviceInterface.
//
// See: https://developer.apple.com/documentation/iokit/iohiddeviceinterface121
type IOHIDDeviceInterface121 = kernel.Pointer

// IOHIDDeviceInterface122 is cFPlugin object subclass which provides the primary interface to HID devices. This class is a subclass of IOHIDDeviceInterface121.
//
// See: https://developer.apple.com/documentation/iokit/iohiddeviceinterface122
type IOHIDDeviceInterface122 = kernel.Pointer

// IOHIDDeviceQueueInterface is the object you use to access a HID queue from user space, returned by version 1.5 of the IOHIDFamily.
//
// See: https://developer.apple.com/documentation/iokit/iohiddevicequeueinterface
type IOHIDDeviceQueueInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iohiddeviceref
type IOHIDDeviceRef = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iohiddevicetimestampeddeviceinterface
type IOHIDDeviceTimeStampedDeviceInterface = kernel.Pointer

// IOHIDDeviceTransactionInterface is the object you use to access a HID transaction from user space, returned by version 1.5 of the IOHIDFamily.
//
// See: https://developer.apple.com/documentation/iokit/iohiddevicetransactioninterface
type IOHIDDeviceTransactionInterface = kernel.Pointer

// IOHIDElementCollectionType is describes different types of HID collections.
//
// See: https://developer.apple.com/documentation/iokit/iohidelementcollectiontype
type IOHIDElementCollectionType = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iohidelementcommitdirection
type IOHIDElementCommitDirection = kernel.Pointer

// IOHIDElementCookie is abstract data type used as a unique identifier for an element.
//
// See: https://developer.apple.com/documentation/iokit/iohidelementcookie
type IOHIDElementCookie = uint32

// See: https://developer.apple.com/documentation/iokit/iohidelementflags
type IOHIDElementFlags = uint32

// See: https://developer.apple.com/documentation/iokit/iohidelementref
type IOHIDElementRef = kernel.Pointer

// IOHIDElementType is describes different types of HID elements.
//
// See: https://developer.apple.com/documentation/iokit/iohidelementtype
type IOHIDElementType = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iohideventstruct
type IOHIDEventStruct = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iohideventsystemclientref
type IOHIDEventSystemClientRef = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iohidkeyboardeventoptions
type IOHIDKeyboardEventOptions = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iohidkeyboardphysicallayouttype
type IOHIDKeyboardPhysicalLayoutType = uint32

// IOHIDManagerOptions is various options that can be supplied to IOHIDManager functions.
//
// See: https://developer.apple.com/documentation/iokit/iohidmanageroptions
type IOHIDManagerOptions = uint32

// IOHIDManagerRef is this is the type of a reference to the IOHIDManager.
//
// See: https://developer.apple.com/documentation/iokit/iohidmanagerref
type IOHIDManagerRef = kernel.Pointer

// IOHIDOptionsType is options for opening a device via IOHIDLib.
//
// See: https://developer.apple.com/documentation/iokit/iohidoptionstype
type IOHIDOptionsType = uint32

// IOHIDOutputTransactionInterface is cFPlugin object subclass which privides interface for output transactions to HID devices. Created by a IOHIDDeviceInterface object.
//
// See: https://developer.apple.com/documentation/iokit/iohidoutputtransactioninterface
type IOHIDOutputTransactionInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iohidpointereventoptions
type IOHIDPointerEventOptions = kernel.Pointer

// IOHIDQueueInterface is cFPlugin object subclass which provides an interface for input queues from HID devices. Created by an IOHIDDeviceInterface object.
//
// See: https://developer.apple.com/documentation/iokit/iohidqueueinterface
type IOHIDQueueInterface = kernel.Pointer

// IOHIDQueueOptionsType is options for creating a queue via IOHIDLib.
//
// See: https://developer.apple.com/documentation/iokit/iohidqueueoptionstype
type IOHIDQueueOptionsType = uint32

// See: https://developer.apple.com/documentation/iokit/iohidqueueref
type IOHIDQueueRef = kernel.Pointer

// IOHIDReportType is describes different type of HID reports.
//
// See: https://developer.apple.com/documentation/iokit/iohidreporttype
type IOHIDReportType = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iohidrequesttype
type IOHIDRequestType = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iohidscrolleventoptions
type IOHIDScrollEventOptions = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iohidserviceclientref
type IOHIDServiceClientRef = kernel.Pointer

// IOHIDStandardType is type to define what industrial standard the device is referencing.
//
// See: https://developer.apple.com/documentation/iokit/iohidstandardtype
type IOHIDStandardType = uint32

// IOHIDTransactionDirectionType is direction for an IOHIDDeviceTransactionInterface.
//
// See: https://developer.apple.com/documentation/iokit/iohidtransactiondirectiontype
type IOHIDTransactionDirectionType = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iohidtransactionoptions
type IOHIDTransactionOptions = uint32

// See: https://developer.apple.com/documentation/iokit/iohidtransactionref
type IOHIDTransactionRef = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iohiduserdevicegetreportblock
type IOHIDUserDeviceGetReportBlock = func(IOHIDReportType, uint32, *uint8, *corefoundation.CFIndex) int

// See: https://developer.apple.com/documentation/iokit/iohiduserdeviceoptions
type IOHIDUserDeviceOptions = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iohiduserdeviceref
type IOHIDUserDeviceRef = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iohiduserdevicesetreportblock
type IOHIDUserDeviceSetReportBlock = func(IOHIDReportType, uint32, *uint8, corefoundation.CFIndex) int

// IOHIDValueOptions is describes options for gathering element values.
//
// See: https://developer.apple.com/documentation/iokit/iohidvalueoptions
type IOHIDValueOptions = uint32

// See: https://developer.apple.com/documentation/iokit/iohidvalueref
type IOHIDValueRef = kernel.Pointer

// IOHIDValueScaleType is describes different types of scaling that can be performed on element values.
//
// See: https://developer.apple.com/documentation/iokit/iohidvaluescaletype
type IOHIDValueScaleType = uint32

// See: https://developer.apple.com/documentation/iokit/iohardwarecursordescriptor
type IOHardwareCursorDescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iohardwarecursorinfo
type IOHardwareCursorInfo = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioi2cbuffer
type IOI2CBuffer = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/ioi2cbustiming
type IOI2CBusTiming = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioi2cconnectref
type IOI2CConnectRef uintptr

// See: https://developer.apple.com/documentation/iokit/ioi2crequest
type IOI2CRequest = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioindex
type IOIndex = int32

// See: https://developer.apple.com/documentation/iokit/ioitemcount
type IOItemCount = uint32

// See: https://developer.apple.com/documentation/iokit/iologicaladdress
type IOLogicalAddress = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iomediaattributemask
type IOMediaAttributeMask = uint32

// See: https://developer.apple.com/documentation/iokit/iomediastate
type IOMediaState = uint32

// See: https://developer.apple.com/documentation/iokit/iomediumtype
type IOMediumType = uint32

// See: https://developer.apple.com/documentation/iokit/iomessage
type IOMessage = uint32

// See: https://developer.apple.com/documentation/iokit/iondhandle
type IONDHandle = uint32

// See: https://developer.apple.com/documentation/iokit/ionvmesmartinterface
type IONVMeSMARTInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ionetworkstats
type IONetworkStats = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ionotificationportref
type IONotificationPortRef = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iooptionbits
type IOOptionBits = uint32

// See: https://developer.apple.com/documentation/iokit/iooutputqueuestats
type IOOutputQueueStats = kernel.Pointer

// IOPMAssertionID is type for AssertionID arguments to [IOPMAssertionCreateWithProperties] and [IOPMAssertionRelease].
//
// See: https://developer.apple.com/documentation/iokit/iopmassertionid
type IOPMAssertionID = uint32

// IOPMAssertionLevel is type for AssertionLevel argument to IOPMAssertionCreate.
//
// See: https://developer.apple.com/documentation/iokit/iopmassertionlevel
type IOPMAssertionLevel = uint32

// See: https://developer.apple.com/documentation/iokit/iopmcalendarstruct
type IOPMCalendarStruct = kernel.Pointer

// IOPMPowerFlags is bits are used in defining capabilityFlags, inputPowerRequirements, and outputPowerCharacter in the IOPMPowerState structure.
//
// See: https://developer.apple.com/documentation/iokit/iopmpowerflags
type IOPMPowerFlags = uint

// See: https://developer.apple.com/documentation/iokit/iopmuseractivetype
type IOPMUserActiveType = kernel.Pointer

// IOPSLowBatteryWarningLevel is the battery can provide no more than 10 minutes of runtime.
//
// See: https://developer.apple.com/documentation/iokit/iopslowbatterywarninglevel
type IOPSLowBatteryWarningLevel = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iophysicaladdress
type IOPhysicalAddress = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iophysicaladdress32
type IOPhysicalAddress32 = uint32

// See: https://developer.apple.com/documentation/iokit/iophysicaladdress64
type IOPhysicalAddress64 = uint64

// See: https://developer.apple.com/documentation/iokit/iophysicallength
type IOPhysicalLength = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iophysicallength32
type IOPhysicalLength32 = uint32

// See: https://developer.apple.com/documentation/iokit/iophysicallength64
type IOPhysicalLength64 = uint64

// See: https://developer.apple.com/documentation/iokit/iopixelaperture
type IOPixelAperture = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iopixelencoding
type IOPixelEncoding = int8

// See: https://developer.apple.com/documentation/iokit/iopixelinformation
type IOPixelInformation = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iopowerstatechangenotification
type IOPowerStateChangeNotification = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioreturn
type IOReturn = int

// See: https://developer.apple.com/documentation/iokit/ioselect
type IOSelect = uint32

// IOServiceInterestCallback is callback function to be notified of changes in state of an IOService.
//
// See: https://developer.apple.com/documentation/iokit/ioserviceinterestcallback
type IOServiceInterestCallback = func(unsafe.Pointer, uintptr, uint32, unsafe.Pointer)

// IOServiceMatchingCallback is callback function to be notified of IOService publication.
//
// See: https://developer.apple.com/documentation/iokit/ioservicematchingcallback
type IOServiceMatchingCallback = func(unsafe.Pointer, uintptr)

// See: https://developer.apple.com/documentation/iokit/iostorageunmapoptions
type IOStorageUnmapOptions = uint32

// See: https://developer.apple.com/documentation/iokit/iostreambufferid
type IOStreamBufferID = uint32

// See: https://developer.apple.com/documentation/iokit/iostreaminterface
type IOStreamInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iostreammode
type IOStreamMode = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iostreamref
type IOStreamRef uintptr

// IOSystemLoadAdvisoryLevel is return type for IOGetSystemLoadAdvisory.
//
// See: https://developer.apple.com/documentation/iokit/iosystemloadadvisorylevel
type IOSystemLoadAdvisoryLevel = int

// See: https://developer.apple.com/documentation/iokit/iotiminginformation
type IOTimingInformation = kernel.Pointer

// IOUPSPlugInInterface is represents and provides management functions for a UPS device.
//
// See: https://developer.apple.com/documentation/iokit/ioupsplugininterface
type IOUPSPlugInInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/ioupsplugininterface_v140
type IOUPSPlugInInterface_v140 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousb20hubdescriptor
type IOUSB20HubDescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbbosdescriptor
type IOUSBBOSDescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbbosdescriptorptr
type IOUSBBOSDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbbulkpipereq
type IOUSBBulkPipeReq = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbcompletion
type IOUSBCompletion = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbcompletionwithtimestamp
type IOUSBCompletionWithTimeStamp = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbconfigurationdescheader
type IOUSBConfigurationDescHeader = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbconfigurationdescheaderptr
type IOUSBConfigurationDescHeaderPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbconfigurationdescriptor
type IOUSBConfigurationDescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbconfigurationdescriptorptr
type IOUSBConfigurationDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdfudescriptor
type IOUSBDFUDescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdfudescriptorptr
type IOUSBDFUDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdescriptor
type IOUSBDescriptor = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdescriptorheader
type IOUSBDescriptorHeader = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdescriptorheaderptr
type IOUSBDescriptorHeaderPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevreqool
type IOUSBDevReqOOL = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevreqoolto
type IOUSBDevReqOOLTO = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevrequest
type IOUSBDevRequest = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevrequestto
type IOUSBDevRequestTO = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitybillboard
type IOUSBDeviceCapabilityBillboard = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitybillboardaltconfig
type IOUSBDeviceCapabilityBillboardAltConfig = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitybillboardaltconfigcompatibility
type IOUSBDeviceCapabilityBillboardAltConfigCompatibility = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitybillboardaltconfigptr
type IOUSBDeviceCapabilityBillboardAltConfigPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitybillboardaltmode
type IOUSBDeviceCapabilityBillboardAltMode = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitybillboardaltmodeptr
type IOUSBDeviceCapabilityBillboardAltModePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitybillboardptr
type IOUSBDeviceCapabilityBillboardPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitycontainerid
type IOUSBDeviceCapabilityContainerID = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitycontaineridptr
type IOUSBDeviceCapabilityContainerIDPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitydescriptorheader
type IOUSBDeviceCapabilityDescriptorHeader = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitydescriptorheaderptr
type IOUSBDeviceCapabilityDescriptorHeaderPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitysuperspeedplususb
type IOUSBDeviceCapabilitySuperSpeedPlusUSB = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitysuperspeedplususbptr
type IOUSBDeviceCapabilitySuperSpeedPlusUSBPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitysuperspeedusb
type IOUSBDeviceCapabilitySuperSpeedUSB = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitysuperspeedusbptr
type IOUSBDeviceCapabilitySuperSpeedUSBPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilityusb2extension
type IOUSBDeviceCapabilityUSB2Extension = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilityusb2extensionptr
type IOUSBDeviceCapabilityUSB2ExtensionPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicedescriptor
type IOUSBDeviceDescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicedescriptorptr
type IOUSBDeviceDescriptorPtr = unsafe.Pointer

// IOUSBDeviceInterface is the object you use to access USB devices from user space, returned by all versions of the IOUSBFamily currently shipping.
//
// See: https://developer.apple.com/documentation/iokit/iousbdeviceinterface
type IOUSBDeviceInterface = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdeviceinterface100
type IOUSBDeviceInterface100 = kernel.Pointer

// IOUSBDeviceInterface182 is the object you use to access USB devices from user space, returned by the IOUSBFamily version 1.8.2 and above.
//
// See: https://developer.apple.com/documentation/iokit/iousbdeviceinterface182
type IOUSBDeviceInterface182 = kernel.Pointer

// IOUSBDeviceInterface187 is the object you use to access USB devices from user space, returned by the IOUSBFamily version 10.8.7 and above.
//
// See: https://developer.apple.com/documentation/iokit/iousbdeviceinterface187
type IOUSBDeviceInterface187 = kernel.Pointer

// IOUSBDeviceInterface197 is the object you use to access USB devices from user space, returned by the IOUSBFamily version 1.9.7 and above.
//
// See: https://developer.apple.com/documentation/iokit/iousbdeviceinterface197
type IOUSBDeviceInterface197 = kernel.Pointer

// IOUSBDeviceInterface245 is the object you use to access USB devices from user space, returned by the IOUSBFamily version 2.4.5 and above.
//
// See: https://developer.apple.com/documentation/iokit/iousbdeviceinterface245
type IOUSBDeviceInterface245 = kernel.Pointer

// IOUSBDeviceInterface300 is the object you use to access USB devices from user space, returned by the IOUSBFamily version 3.0.0 and above.
//
// See: https://developer.apple.com/documentation/iokit/iousbdeviceinterface300
type IOUSBDeviceInterface300 = kernel.Pointer

// IOUSBDeviceInterface320 is the object you use to access USB devices from user space, returned by the IOUSBFamily version 3.2.0 and above.
//
// See: https://developer.apple.com/documentation/iokit/iousbdeviceinterface320
type IOUSBDeviceInterface320 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdeviceinterface400
type IOUSBDeviceInterface400 = kernel.Pointer

// IOUSBDeviceInterface500 is the object you use to access USB devices from user space, returned by the IOUSBFamily version 3.2.0 and above.
//
// See: https://developer.apple.com/documentation/iokit/iousbdeviceinterface500
type IOUSBDeviceInterface500 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdeviceinterface650
type IOUSBDeviceInterface650 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdeviceinterface942
type IOUSBDeviceInterface942 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicequalifierdescriptor
type IOUSBDeviceQualifierDescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicequalifierdescriptorptr
type IOUSBDeviceQualifierDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicerequest
type IOUSBDeviceRequest = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicerequestptr
type IOUSBDeviceRequestPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicerequestsetseldata
type IOUSBDeviceRequestSetSELData = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbendpointdescriptor
type IOUSBEndpointDescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbendpointdescriptorptr
type IOUSBEndpointDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbendpointproperties
type IOUSBEndpointProperties = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbendpointpropertiesptr
type IOUSBEndpointPropertiesPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbfindendpointrequest
type IOUSBFindEndpointRequest = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbfindinterfacerequest
type IOUSBFindInterfaceRequest = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbgetframestruct
type IOUSBGetFrameStruct = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbhiddataptr
type IOUSBHIDDataPtr = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbhiddescriptor
type IOUSBHIDDescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbhiddescriptorptr
type IOUSBHIDDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbhidreportdesc
type IOUSBHIDReportDesc = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbhidreportdescptr
type IOUSBHIDReportDescPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceassociationdescriptor
type IOUSBInterfaceAssociationDescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceassociationdescriptorptr
type IOUSBInterfaceAssociationDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbinterfacedescriptor
type IOUSBInterfaceDescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbinterfacedescriptorptr
type IOUSBInterfaceDescriptorPtr = unsafe.Pointer

// IOUSBInterfaceInterface is the object you use to access a USB device interface from user space, returned by all versions of the IOUSBFamily currently shipping.
//
// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface
type IOUSBInterfaceInterface = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface100
type IOUSBInterfaceInterface100 = kernel.Pointer

// IOUSBInterfaceInterface182 is the object you use to access a USB device interface from user space, returned by the IOUSBFamily version 1.8.2 and above.
//
// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface182
type IOUSBInterfaceInterface182 = kernel.Pointer

// IOUSBInterfaceInterface183 is the object you use to access a USB device interface from user space, returned by the IOUSBFamily version 1.8.3 and above.
//
// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface183
type IOUSBInterfaceInterface183 = kernel.Pointer

// IOUSBInterfaceInterface190 is the object you use to access a USB device interface from user space, returned by the IOUSBFamily version 1.9 and above.
//
// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface190
type IOUSBInterfaceInterface190 = kernel.Pointer

// IOUSBInterfaceInterface192 is the object you use to access a USB device interface from user space, returned by the IOUSBFamily version 1.9.2 and above.
//
// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface192
type IOUSBInterfaceInterface192 = kernel.Pointer

// IOUSBInterfaceInterface197 is the object you use to access a USB device interface from user space, returned by the IOUSBFamily version 1.9.7 and above.
//
// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface197
type IOUSBInterfaceInterface197 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface220
type IOUSBInterfaceInterface220 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface245
type IOUSBInterfaceInterface245 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface300
type IOUSBInterfaceInterface300 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface398
type IOUSBInterfaceInterface398 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface400
type IOUSBInterfaceInterface400 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface500
type IOUSBInterfaceInterface500 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface550
type IOUSBInterfaceInterface550 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface650
type IOUSBInterfaceInterface650 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface700
type IOUSBInterfaceInterface700 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface800
type IOUSBInterfaceInterface800 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface942
type IOUSBInterfaceInterface942 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbisoccompletion
type IOUSBIsocCompletion = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbisocframe
type IOUSBIsocFrame = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbisocstruct
type IOUSBIsocStruct = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbkeyboarddata
type IOUSBKeyboardData = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbkeyboarddataptr
type IOUSBKeyboardDataPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousblowlatencyisoccompletion
type IOUSBLowLatencyIsocCompletion = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousblowlatencyisocframe
type IOUSBLowLatencyIsocFrame = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousblowlatencyisocstruct
type IOUSBLowLatencyIsocStruct = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbmatch
type IOUSBMatch = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbmousedata
type IOUSBMouseData = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbmousedataptr
type IOUSBMouseDataPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbplatformcapabilitydescriptor
type IOUSBPlatformCapabilityDescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbplatformcapabilitydescriptorptr
type IOUSBPlatformCapabilityDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbstringdescriptor
type IOUSBStringDescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbstringdescriptorptr
type IOUSBStringDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbsuperspeedendpointcompaniondescriptor
type IOUSBSuperSpeedEndpointCompanionDescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbsuperspeedendpointcompaniondescriptorptr
type IOUSBSuperSpeedEndpointCompanionDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbsuperspeedhubdescriptor
type IOUSBSuperSpeedHubDescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbsuperspeedplusisochronousendpointcompaniondescriptor
type IOUSBSuperSpeedPlusIsochronousEndpointCompanionDescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbsuperspeedplusisochronousendpointcompaniondescriptorptr
type IOUSBSuperSpeedPlusIsochronousEndpointCompanionDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/ioversion
type IOVersion = uint32

// See: https://developer.apple.com/documentation/iokit/iovideodeviceinterface
type IOVideoDeviceInterface = IOVideoDeviceInterface_v1_t

// IOVideoDeviceInterface_v1_t is forward declaration of IOVideoDeviceInterface_v1_t.
//
// See: https://developer.apple.com/documentation/iokit/iovideodeviceinterface_v1_t
type IOVideoDeviceInterface_v1_t = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iovideodevicenotification
type IOVideoDeviceNotification = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iovideodevicenotificationmessage
type IOVideoDeviceNotificationMessage = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iovideodeviceref
type IOVideoDeviceRef uintptr

// See: https://developer.apple.com/documentation/iokit/iovideostreamdescription
type IOVideoStreamDescription = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/iovirtualaddress
type IOVirtualAddress = kernel.MachVmAddress

// See: https://developer.apple.com/documentation/iokit/longlbamodeparameterblockdescriptor
type LongLBAModeParameterBlockDescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/lowlatencyuserbufferinfo
type LowLatencyUserBufferInfo = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/lowlatencyuserbufferinfov2
type LowLatencyUserBufferInfoV2 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/lowlatencyuserbufferinfov3
type LowLatencyUserBufferInfoV3 = kernel.Pointer

// MMCDeviceInterface is basic interface for an MMC-2 Compliant Device.
//
// See: https://developer.apple.com/documentation/iokit/mmcdeviceinterface
type MMCDeviceInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/modepageformatheader
type ModePageFormatHeader = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/modeparameterblockdescriptor
type ModeParameterBlockDescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/nvmeidentifycontrollerstruct
type NVMeIdentifyControllerStruct = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/nvmeidentifynamespacestruct
type NVMeIdentifyNamespaceStruct = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/nvmelbaformatdatastruct
type NVMeLBAFormatDataStruct = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/nvmepowerstatedescriptor
type NVMePowerStateDescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/nvmesmartdata
type NVMeSMARTData = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/nxcoord
type NXCoord = float32

// See: https://developer.apple.com/documentation/iokit/nxeqelement
type NXEQElement = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/nxevent
type NXEvent = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/nxeventext
type NXEventExt = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/nxeventextension
type NXEventExtension = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/nxeventhandle
type NXEventHandle = uint32

// See: https://developer.apple.com/documentation/iokit/nxeventptr
type NXEventPtr = uintptr

// See: https://developer.apple.com/documentation/iokit/nxeventsystemdevice
type NXEventSystemDevice = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/nxeventsystemdevicelist
type NXEventSystemDeviceList = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/nxeventsysteminfodata
type NXEventSystemInfoData = int

// See: https://developer.apple.com/documentation/iokit/nxeventsysteminfotype
type NXEventSystemInfoType = int

// See: https://developer.apple.com/documentation/iokit/nxkeymapping
type NXKeyMapping = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/nxmousebutton
type NXMouseButton = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/nxmousescaling
type NXMouseScaling = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/nxparsedkeymapping
type NXParsedKeyMapping = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/nxpoint
type NXPoint = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/nxsize
type NXSize = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/nxtabletpointdata
type NXTabletPointData = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/nxtabletpointdataptr
type NXTabletPointDataPtr = uintptr

// See: https://developer.apple.com/documentation/iokit/nxtabletproximitydata
type NXTabletProximityData = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/nxtabletproximitydataptr
type NXTabletProximityDataPtr = uintptr

// See: https://developer.apple.com/documentation/iokit/nudclflags
type NuDCLFlags = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/nudclreceivepacketref
type NuDCLReceivePacketRef uintptr

// See: https://developer.apple.com/documentation/iokit/nudclref
type NuDCLRef uintptr

// See: https://developer.apple.com/documentation/iokit/nudclsendpacketref
type NuDCLSendPacketRef uintptr

// See: https://developer.apple.com/documentation/iokit/nudclskipcycleref
type NuDCLSkipCycleRef uintptr

// See: https://developer.apple.com/documentation/iokit/osasyncreference
type OSAsyncReference = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/osasyncreference64
type OSAsyncReference64 = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/osobjectref
type OSObjectRef = uint64

// See: https://developer.apple.com/documentation/iokit/report_luns_logical_unit_addressing
type REPORT_LUNS_LOGICAL_UNIT_ADDRESSING = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/report_luns_peripheral_device_addressing
type REPORT_LUNS_PERIPHERAL_DEVICE_ADDRESSING = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/rgbcolor
type RGBColor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/rgbcolorhdl
type RGBColorHdl = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/rgbcolorptr
type RGBColorPtr = applicationservices.RGBColor

// See: https://developer.apple.com/documentation/iokit/rawsensecode
type RawSenseCode = uint8

// See: https://developer.apple.com/documentation/iokit/regentryid
type RegEntryID = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/regentryidptr
type RegEntryIDPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/sbcmodepagecaching
type SBCModePageCaching = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/sbcmodepageflexibledisk
type SBCModePageFlexibleDisk = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/sbcmodepageformatdevice
type SBCModePageFormatDevice = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/sbcmodepagerigiddiskgeometry
type SBCModePageRigidDiskGeometry = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicmdfield10bit
type SCSICmdField10Bit = uint16

// See: https://developer.apple.com/documentation/iokit/scsicmdfield11bit
type SCSICmdField11Bit = uint16

// See: https://developer.apple.com/documentation/iokit/scsicmdfield12bit
type SCSICmdField12Bit = uint16

// See: https://developer.apple.com/documentation/iokit/scsicmdfield13bit
type SCSICmdField13Bit = uint16

// See: https://developer.apple.com/documentation/iokit/scsicmdfield14bit
type SCSICmdField14Bit = uint16

// See: https://developer.apple.com/documentation/iokit/scsicmdfield15bit
type SCSICmdField15Bit = uint16

// See: https://developer.apple.com/documentation/iokit/scsicmdfield17bit
type SCSICmdField17Bit = uint32

// See: https://developer.apple.com/documentation/iokit/scsicmdfield18bit
type SCSICmdField18Bit = uint32

// See: https://developer.apple.com/documentation/iokit/scsicmdfield19bit
type SCSICmdField19Bit = uint32

// See: https://developer.apple.com/documentation/iokit/scsicmdfield1bit
type SCSICmdField1Bit = uint8

// See: https://developer.apple.com/documentation/iokit/scsicmdfield1byte
type SCSICmdField1Byte = uint8

// See: https://developer.apple.com/documentation/iokit/scsicmdfield20bit
type SCSICmdField20Bit = uint32

// See: https://developer.apple.com/documentation/iokit/scsicmdfield21bit
type SCSICmdField21Bit = uint32

// See: https://developer.apple.com/documentation/iokit/scsicmdfield22bit
type SCSICmdField22Bit = uint32

// See: https://developer.apple.com/documentation/iokit/scsicmdfield23bit
type SCSICmdField23Bit = uint32

// See: https://developer.apple.com/documentation/iokit/scsicmdfield25bit
type SCSICmdField25Bit = uint32

// See: https://developer.apple.com/documentation/iokit/scsicmdfield26bit
type SCSICmdField26Bit = uint32

// See: https://developer.apple.com/documentation/iokit/scsicmdfield27bit
type SCSICmdField27Bit = uint32

// See: https://developer.apple.com/documentation/iokit/scsicmdfield28bit
type SCSICmdField28Bit = uint32

// See: https://developer.apple.com/documentation/iokit/scsicmdfield29bit
type SCSICmdField29Bit = uint32

// See: https://developer.apple.com/documentation/iokit/scsicmdfield2bit
type SCSICmdField2Bit = uint8

// See: https://developer.apple.com/documentation/iokit/scsicmdfield2byte
type SCSICmdField2Byte = uint16

// See: https://developer.apple.com/documentation/iokit/scsicmdfield30bit
type SCSICmdField30Bit = uint32

// See: https://developer.apple.com/documentation/iokit/scsicmdfield31bit
type SCSICmdField31Bit = uint32

// See: https://developer.apple.com/documentation/iokit/scsicmdfield33bit
type SCSICmdField33Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield34bit
type SCSICmdField34Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield35bit
type SCSICmdField35Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield36bit
type SCSICmdField36Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield37bit
type SCSICmdField37Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield38bit
type SCSICmdField38Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield39bit
type SCSICmdField39Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield3bit
type SCSICmdField3Bit = uint8

// See: https://developer.apple.com/documentation/iokit/scsicmdfield3byte
type SCSICmdField3Byte = uint32

// See: https://developer.apple.com/documentation/iokit/scsicmdfield41bit
type SCSICmdField41Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield42bit
type SCSICmdField42Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield43bit
type SCSICmdField43Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield44bit
type SCSICmdField44Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield45bit
type SCSICmdField45Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield46bit
type SCSICmdField46Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield47bit
type SCSICmdField47Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield49bit
type SCSICmdField49Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield4bit
type SCSICmdField4Bit = uint8

// See: https://developer.apple.com/documentation/iokit/scsicmdfield4byte
type SCSICmdField4Byte = uint32

// See: https://developer.apple.com/documentation/iokit/scsicmdfield50bit
type SCSICmdField50Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield51bit
type SCSICmdField51Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield52bit
type SCSICmdField52Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield53bit
type SCSICmdField53Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield54bit
type SCSICmdField54Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield55bit
type SCSICmdField55Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield57bit
type SCSICmdField57Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield58bit
type SCSICmdField58Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield59bit
type SCSICmdField59Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield5bit
type SCSICmdField5Bit = uint8

// See: https://developer.apple.com/documentation/iokit/scsicmdfield5byte
type SCSICmdField5Byte = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield60bit
type SCSICmdField60Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield61bit
type SCSICmdField61Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield62bit
type SCSICmdField62Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield63bit
type SCSICmdField63Bit = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield6bit
type SCSICmdField6Bit = uint8

// See: https://developer.apple.com/documentation/iokit/scsicmdfield6byte
type SCSICmdField6Byte = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield7bit
type SCSICmdField7Bit = uint8

// See: https://developer.apple.com/documentation/iokit/scsicmdfield7byte
type SCSICmdField7Byte = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield8byte
type SCSICmdField8Byte = uint64

// See: https://developer.apple.com/documentation/iokit/scsicmdfield9bit
type SCSICmdField9Bit = uint16

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_pagecx_header
type SCSICmd_INQUIRY_PAGECx_Header = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_page00_header
type SCSICmd_INQUIRY_Page00_Header = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_page00_header_spc_16
type SCSICmd_INQUIRY_Page00_Header_SPC_16 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_page80_header
type SCSICmd_INQUIRY_Page80_Header = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_page80_header_spc_16
type SCSICmd_INQUIRY_Page80_Header_SPC_16 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_page83_header
type SCSICmd_INQUIRY_Page83_Header = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_page83_header_spc_16
type SCSICmd_INQUIRY_Page83_Header_SPC_16 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_page83_identification_descriptor
type SCSICmd_INQUIRY_Page83_Identification_Descriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_page83_logicalunitgroup_identifier
type SCSICmd_INQUIRY_Page83_LogicalUnitGroup_Identifier = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_page83_relativetargetport_identifier
type SCSICmd_INQUIRY_Page83_RelativeTargetPort_Identifier = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_page83_targetportgroup_identifier
type SCSICmd_INQUIRY_Page83_TargetPortGroup_Identifier = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_page89_data
type SCSICmd_INQUIRY_Page89_Data = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_pageb0_data
type SCSICmd_INQUIRY_PageB0_Data = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_pageb1_data
type SCSICmd_INQUIRY_PageB1_Data = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_pageb2_data
type SCSICmd_INQUIRY_PageB2_Data = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_pageb2_provisioning_group_descriptor
type SCSICmd_INQUIRY_PageB2_Provisioning_Group_Descriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_pagec0_data
type SCSICmd_INQUIRY_PageC0_Data = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_pagec1_data
type SCSICmd_INQUIRY_PageC1_Data = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_standarddata
type SCSICmd_INQUIRY_StandardData = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_standarddataall
type SCSICmd_INQUIRY_StandardDataAll = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_standarddataptr
type SCSICmd_INQUIRY_StandardDataPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicmd_report_luns_header
type SCSICmd_REPORT_LUNS_Header = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicmd_report_luns_lun_entry
type SCSICmd_REPORT_LUNS_LUN_ENTRY = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicommanddescriptorblock
type SCSICommandDescriptorBlock = uint8

// SCSIDeviceIdentifier is 64-bit number to represent a SCSI Device.
//
// See: https://developer.apple.com/documentation/iokit/scsideviceidentifier
type SCSIDeviceIdentifier = uint64

// SCSIInitiatorIdentifier is 64-bit number to represent a SCSI Initiator Device.
//
// See: https://developer.apple.com/documentation/iokit/scsiinitiatoridentifier
type SCSIInitiatorIdentifier = string

// See: https://developer.apple.com/documentation/iokit/scsilogicalunitbytes
type SCSILogicalUnitBytes = uint8

// See: https://developer.apple.com/documentation/iokit/scsilogicalunitnumber
type SCSILogicalUnitNumber = uint64

// SCSIServiceResponse is attributes for task service response.
//
// See: https://developer.apple.com/documentation/iokit/scsiserviceresponse
type SCSIServiceResponse = kernel.Pointer

// SCSITaggedTaskIdentifier is 64-bit number to represent a unique task identifier.
//
// See: https://developer.apple.com/documentation/iokit/scsitaggedtaskidentifier
type SCSITaggedTaskIdentifier = uint64

// SCSITargetIdentifier is 64-bit number to represent a SCSI Target Device.
//
// See: https://developer.apple.com/documentation/iokit/scsitargetidentifier
type SCSITargetIdentifier = string

// SCSITaskAttribute is attributes for task delivery.
//
// See: https://developer.apple.com/documentation/iokit/scsitaskattribute
type SCSITaskAttribute = kernel.Pointer

// SCSITaskDeviceInterface is basic interface for a SCSITask Device.
//
// See: https://developer.apple.com/documentation/iokit/scsitaskdeviceinterface
type SCSITaskDeviceInterface = kernel.Pointer

// SCSITaskInterface is basic interface for a SCSITask.
//
// See: https://developer.apple.com/documentation/iokit/scsitaskinterface
type SCSITaskInterface = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsitasksgelement
type SCSITaskSGElement = unsafe.Pointer

// SCSITaskState is attributes for task state.
//
// See: https://developer.apple.com/documentation/iokit/scsitaskstate
type SCSITaskState = kernel.Pointer

// SCSITaskStatus is attributes for task status.
//
// See: https://developer.apple.com/documentation/iokit/scsitaskstatus
type SCSITaskStatus = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsi_capacity_data
type SCSI_Capacity_Data = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsi_capacity_data_long
type SCSI_Capacity_Data_Long = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/scsi_sense_data
type SCSI_Sense_Data = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/spcmodepagepowercondition
type SPCModePagePowerCondition = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/spcmodeparameterheader10
type SPCModeParameterHeader10 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/spcmodeparameterheader6
type SPCModeParameterHeader6 = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/uaspipedescriptor
type UASPipeDescriptor = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/uaspipedescriptorptr
type UASPipeDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/usbdeviceaddress
type USBDeviceAddress = uint16

// See: https://developer.apple.com/documentation/iokit/usbdeviceinformationbits
type USBDeviceInformationBits = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/usblowlatencybuffertype
type USBLowLatencyBufferType = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/usbnotificationtypes
type USBNotificationTypes = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/usbphysicaladdress32
type USBPhysicalAddress32 = uint32

// See: https://developer.apple.com/documentation/iokit/usbpowerrequesttypes
type USBPowerRequestTypes = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/usbreenumerateoptions
type USBReEnumerateOptions = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/usbstatus
type USBStatus = uint16

// See: https://developer.apple.com/documentation/iokit/usbstatusptr
type USBStatusPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/userexportdclcallcommandproc
type UserExportDCLCallCommandProc = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/userexportdclcallproc
type UserExportDCLCallProc = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/userexportdclcommand
type UserExportDCLCommand = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/userexportdcljump
type UserExportDCLJump = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/userexportdcllabel
type UserExportDCLLabel = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/userexportdclnudclleader
type UserExportDCLNuDCLLeader = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/userexportdclptrtimestamp
type UserExportDCLPtrTimeStamp = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/userexportdclsettagsyncbits
type UserExportDCLSetTagSyncBits = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/userexportdcltimestamp
type UserExportDCLTimeStamp = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/userexportdcltransferbuffer
type UserExportDCLTransferBuffer = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/userexportdcltransferpacket
type UserExportDCLTransferPacket = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/userexportdclupdatedcllist
type UserExportDCLUpdateDCLList = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdclutbehavior
type VDClutBehavior = uint32

// See: https://developer.apple.com/documentation/iokit/vdclutbehaviorptr
type VDClutBehaviorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdcommunicationinfoptr
type VDCommunicationInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdcommunicationinforec
type VDCommunicationInfoRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdcommunicationptr
type VDCommunicationPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdcommunicationrec
type VDCommunicationRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdconfigurationfeaturelistrec
type VDConfigurationFeatureListRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdconfigurationfeaturelistrecptr
type VDConfigurationFeatureListRecPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdconfigurationptr
type VDConfigurationPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdconfigurationrec
type VDConfigurationRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdconvolutioninfoptr
type VDConvolutionInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdconvolutioninforec
type VDConvolutionInfoRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdddcblockptr
type VDDDCBlockPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdddcblockrec
type VDDDCBlockRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vddefmode
type VDDefMode = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vddefmodeptr
type VDDefModePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vddetailedtimingptr
type VDDetailedTimingPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vddetailedtimingrec
type VDDetailedTimingRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vddisplayconnectinfoptr
type VDDisplayConnectInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vddisplayconnectinforec
type VDDisplayConnectInfoRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vddisplaytimingrangeptr
type VDDisplayTimingRangePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vddisplaytimingrangerec
type VDDisplayTimingRangeRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vddrawhardwarecursorptr
type VDDrawHardwareCursorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vddrawhardwarecursorrec
type VDDrawHardwareCursorRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdentrecptr
type VDEntRecPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdentryrecord
type VDEntryRecord = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdflagrecptr
type VDFlagRecPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdflagrecord
type VDFlagRecord = kernel.Pointer

// VDGamRecPtr is represents a type used by the Video Components API.
//
// See: https://developer.apple.com/documentation/iokit/vdgamrecptr
type VDGamRecPtr = applicationservices.VDGammaRecord

// See: https://developer.apple.com/documentation/iokit/vdgammainfoptr
type VDGammaInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdgammainforec
type VDGammaInfoRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdgammarecord
type VDGammaRecord = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdgetgammalistptr
type VDGetGammaListPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdgetgammalistrec
type VDGetGammaListRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdgrayptr
type VDGrayPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdgrayrecord
type VDGrayRecord = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdhardwarecursordrawstateptr
type VDHardwareCursorDrawStatePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdhardwarecursordrawstaterec
type VDHardwareCursorDrawStateRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdmirrorptr
type VDMirrorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdmirrorrec
type VDMirrorRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdmulticonnectinfoptr
type VDMultiConnectInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdmulticonnectinforec
type VDMultiConnectInfoRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdpageinfo
type VDPageInfo = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdpginfoptr
type VDPgInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdpowerstateptr
type VDPowerStatePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdpowerstaterec
type VDPowerStateRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdprivateselectordatarec
type VDPrivateSelectorDataRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdprivateselectorrec
type VDPrivateSelectorRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdresolutioninfoptr
type VDResolutionInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdresolutioninforec
type VDResolutionInfoRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdretrievegammaptr
type VDRetrieveGammaPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdretrievegammarec
type VDRetrieveGammaRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdscalerinfoptr
type VDScalerInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdscalerinforec
type VDScalerInfoRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdscalerptr
type VDScalerPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdscalerrec
type VDScalerRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdsetentryptr
type VDSetEntryPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdsetentryrecord
type VDSetEntryRecord = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdsethardwarecursorptr
type VDSetHardwareCursorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdsethardwarecursorrec
type VDSetHardwareCursorRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdsettings
type VDSettings = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdsettingsptr
type VDSettingsPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdsizeinfo
type VDSizeInfo = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdsupportshardwarecursorptr
type VDSupportsHardwareCursorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdsupportshardwarecursorrec
type VDSupportsHardwareCursorRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdswitchinfoptr
type VDSwitchInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdswitchinforec
type VDSwitchInfoRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdsyncinfoptr
type VDSyncInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdsyncinforec
type VDSyncInfoRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdszinfoptr
type VDSzInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdtiminginfoptr
type VDTimingInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdtiminginforec
type VDTimingInfoRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vdvideoparametersinfoptr
type VDVideoParametersInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdvideoparametersinforec
type VDVideoParametersInfoRec = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vpblock
type VPBlock = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/vpblockptr
type VPBlockPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/videodevicetype
type VideoDeviceType = uint32

// See: https://developer.apple.com/documentation/iokit/dk_bd_read_disc_info_t
type Dk_bd_read_disc_info_t = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dk_bd_read_structure_t
type Dk_bd_read_structure_t = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dk_bd_read_track_info_t
type Dk_bd_read_track_info_t = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dk_bd_report_key_t
type Dk_bd_report_key_t = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dk_bd_send_key_t
type Dk_bd_send_key_t = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dk_cd_read_disc_info_t
type Dk_cd_read_disc_info_t = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dk_cd_read_isrc_t
type Dk_cd_read_isrc_t = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dk_cd_read_mcn_t
type Dk_cd_read_mcn_t = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dk_cd_read_t
type Dk_cd_read_t = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dk_cd_read_toc_t
type Dk_cd_read_toc_t = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dk_cd_read_track_info_t
type Dk_cd_read_track_info_t = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dk_dvd_read_disc_info_t
type Dk_dvd_read_disc_info_t = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dk_dvd_read_rzone_info_t
type Dk_dvd_read_rzone_info_t = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dk_dvd_read_structure_t
type Dk_dvd_read_structure_t = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dk_dvd_report_key_t
type Dk_dvd_report_key_t = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/dk_dvd_send_key_t
type Dk_dvd_send_key_t = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/eioaccelsurfacelockbits
type EIOAccelSurfaceLockBits = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/eioaccelsurfacemodebits
type EIOAccelSurfaceModeBits = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/eioaccelsurfacescalebits
type EIOAccelSurfaceScaleBits = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/eioaccelsurfaceshapebits
type EIOAccelSurfaceShapeBits = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/eioaccelsurfacestatebits
type EIOAccelSurfaceStateBits = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/eviospecialkeymsg_t
type EvioSpecialKeyMsg_t = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/evsioevsioccsindices
type EvsioEVSIOCCSIndices = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/evsioevsioscsindices
type EvsioEVSIOSCSIndices = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/io_connect_t
type Io_connect_t = uintptr

// See: https://developer.apple.com/documentation/iokit/io_enumerator_t
type Io_enumerator_t = uintptr

// See: https://developer.apple.com/documentation/iokit/io_ident_t
type Io_ident_t = uintptr

// See: https://developer.apple.com/documentation/iokit/io_iterator_t
type Io_iterator_t = uintptr

// See: https://developer.apple.com/documentation/iokit/io_object_t
type Io_object_t = uintptr

// See: https://developer.apple.com/documentation/iokit/io_registry_entry_t
type Io_registry_entry_t = uintptr

// See: https://developer.apple.com/documentation/iokit/io_service_t
type Io_service_t = uintptr

// See: https://developer.apple.com/documentation/iokit/kusbconnectable
type KUSBConnectable = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/kusbhostconnectortype
type KUSBHostConnectorType = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/sleepwakenote
type SleepWakeNote = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/tiousbdescriptorsize
type TIOUSBDescriptorSize = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/tiousbdescriptortype
type TIOUSBDescriptorType = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/tiousbdevicecapabilitytype
type TIOUSBDeviceCapabilityType = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/tiousbdevicerequestdirectionvalue
type TIOUSBDeviceRequestDirectionValue = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/tiousbdevicerequestrecipientvalue
type TIOUSBDeviceRequestRecipientValue = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/tiousbdevicerequesttypevalue
type TIOUSBDeviceRequestTypeValue = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/tiousbendpointdirection
type TIOUSBEndpointDirection = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/tiousbendpointsynchronizationtype
type TIOUSBEndpointSynchronizationType = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/tiousbendpointtype
type TIOUSBEndpointType = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/tiousbendpointusagetype
type TIOUSBEndpointUsageType = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/tiousblanguageid
type TIOUSBLanguageID = kernel.Pointer

// See: https://developer.apple.com/documentation/iokit/uext_object_t
type Uext_object_t = uintptr

// See: https://developer.apple.com/documentation/iokit/user_shspeed_t
type User_shspeed_t = uint32

// See: https://developer.apple.com/documentation/iokit/user_speed_t
type User_speed_t = uint64

// See: https://developer.apple.com/documentation/iokit/user_ul_t
type User_ul_t = uint64

// See: https://developer.apple.com/documentation/iokit/user_us_t
type User_us_t = uint32

// IOFireWireCompareSwapCommandInterfaceV3 is a Go-name alias for IOFireWireCompareSwapCommandInterface_v3.
type IOFireWireCompareSwapCommandInterfaceV3 = IOFireWireCompareSwapCommandInterface_v3

// IOUPSPlugInInterfaceV140 is a Go-name alias for IOUPSPlugInInterface_v140.
type IOUPSPlugInInterfaceV140 = IOUPSPlugInInterface_v140

// IOVideoDeviceInterfaceV1 is a Go-name alias for IOVideoDeviceInterface_v1_t.
type IOVideoDeviceInterfaceV1 = IOVideoDeviceInterface_v1_t

// ReportLunsLogicalUnitAddressing is a Go-name alias for REPORT_LUNS_LOGICAL_UNIT_ADDRESSING.
type ReportLunsLogicalUnitAddressing = REPORT_LUNS_LOGICAL_UNIT_ADDRESSING

// ReportLunsPeripheralDeviceAddressing is a Go-name alias for REPORT_LUNS_PERIPHERAL_DEVICE_ADDRESSING.
type ReportLunsPeripheralDeviceAddressing = REPORT_LUNS_PERIPHERAL_DEVICE_ADDRESSING

// SCSICmdInquiryPAGECxHeader is a Go-name alias for SCSICmd_INQUIRY_PAGECx_Header.
type SCSICmdInquiryPAGECxHeader = SCSICmd_INQUIRY_PAGECx_Header

// SCSICmdInquiryPage00Header is a Go-name alias for SCSICmd_INQUIRY_Page00_Header.
type SCSICmdInquiryPage00Header = SCSICmd_INQUIRY_Page00_Header

// SCSICmdInquiryPage00HeaderSpc16 is a Go-name alias for SCSICmd_INQUIRY_Page00_Header_SPC_16.
type SCSICmdInquiryPage00HeaderSpc16 = SCSICmd_INQUIRY_Page00_Header_SPC_16

// SCSICmdInquiryPage80Header is a Go-name alias for SCSICmd_INQUIRY_Page80_Header.
type SCSICmdInquiryPage80Header = SCSICmd_INQUIRY_Page80_Header

// SCSICmdInquiryPage80HeaderSpc16 is a Go-name alias for SCSICmd_INQUIRY_Page80_Header_SPC_16.
type SCSICmdInquiryPage80HeaderSpc16 = SCSICmd_INQUIRY_Page80_Header_SPC_16

// SCSICmdInquiryPage83Header is a Go-name alias for SCSICmd_INQUIRY_Page83_Header.
type SCSICmdInquiryPage83Header = SCSICmd_INQUIRY_Page83_Header

// SCSICmdInquiryPage83HeaderSpc16 is a Go-name alias for SCSICmd_INQUIRY_Page83_Header_SPC_16.
type SCSICmdInquiryPage83HeaderSpc16 = SCSICmd_INQUIRY_Page83_Header_SPC_16

// SCSICmdInquiryPage83IdentificationDescriptor is a Go-name alias for SCSICmd_INQUIRY_Page83_Identification_Descriptor.
type SCSICmdInquiryPage83IdentificationDescriptor = SCSICmd_INQUIRY_Page83_Identification_Descriptor

// SCSICmdInquiryPage83LogicalUnitGroupIdentifier is a Go-name alias for SCSICmd_INQUIRY_Page83_LogicalUnitGroup_Identifier.
type SCSICmdInquiryPage83LogicalUnitGroupIdentifier = SCSICmd_INQUIRY_Page83_LogicalUnitGroup_Identifier

// SCSICmdInquiryPage83RelativeTargetPortIdentifier is a Go-name alias for SCSICmd_INQUIRY_Page83_RelativeTargetPort_Identifier.
type SCSICmdInquiryPage83RelativeTargetPortIdentifier = SCSICmd_INQUIRY_Page83_RelativeTargetPort_Identifier

// SCSICmdInquiryPage83TargetPortGroupIdentifier is a Go-name alias for SCSICmd_INQUIRY_Page83_TargetPortGroup_Identifier.
type SCSICmdInquiryPage83TargetPortGroupIdentifier = SCSICmd_INQUIRY_Page83_TargetPortGroup_Identifier

// SCSICmdInquiryPage89Data is a Go-name alias for SCSICmd_INQUIRY_Page89_Data.
type SCSICmdInquiryPage89Data = SCSICmd_INQUIRY_Page89_Data

// SCSICmdInquiryPageB0Data is a Go-name alias for SCSICmd_INQUIRY_PageB0_Data.
type SCSICmdInquiryPageB0Data = SCSICmd_INQUIRY_PageB0_Data

// SCSICmdInquiryPageB1Data is a Go-name alias for SCSICmd_INQUIRY_PageB1_Data.
type SCSICmdInquiryPageB1Data = SCSICmd_INQUIRY_PageB1_Data

// SCSICmdInquiryPageB2Data is a Go-name alias for SCSICmd_INQUIRY_PageB2_Data.
type SCSICmdInquiryPageB2Data = SCSICmd_INQUIRY_PageB2_Data

// SCSICmdInquiryPageB2ProvisioningGroupDescriptor is a Go-name alias for SCSICmd_INQUIRY_PageB2_Provisioning_Group_Descriptor.
type SCSICmdInquiryPageB2ProvisioningGroupDescriptor = SCSICmd_INQUIRY_PageB2_Provisioning_Group_Descriptor

// SCSICmdInquiryPageC0Data is a Go-name alias for SCSICmd_INQUIRY_PageC0_Data.
type SCSICmdInquiryPageC0Data = SCSICmd_INQUIRY_PageC0_Data

// SCSICmdInquiryPageC1Data is a Go-name alias for SCSICmd_INQUIRY_PageC1_Data.
type SCSICmdInquiryPageC1Data = SCSICmd_INQUIRY_PageC1_Data

// SCSICmdInquiryStandardData is a Go-name alias for SCSICmd_INQUIRY_StandardData.
type SCSICmdInquiryStandardData = SCSICmd_INQUIRY_StandardData

// SCSICmdInquiryStandardDataAll is a Go-name alias for SCSICmd_INQUIRY_StandardDataAll.
type SCSICmdInquiryStandardDataAll = SCSICmd_INQUIRY_StandardDataAll

// SCSICmdInquiryStandardDataPtr is a Go-name alias for SCSICmd_INQUIRY_StandardDataPtr.
type SCSICmdInquiryStandardDataPtr = SCSICmd_INQUIRY_StandardDataPtr

// SCSICmdReportLunsHeader is a Go-name alias for SCSICmd_REPORT_LUNS_Header.
type SCSICmdReportLunsHeader = SCSICmd_REPORT_LUNS_Header

// SCSICmdReportLunsLunEntry is a Go-name alias for SCSICmd_REPORT_LUNS_LUN_ENTRY.
type SCSICmdReportLunsLunEntry = SCSICmd_REPORT_LUNS_LUN_ENTRY

// ScsiCapacityData is a Go-name alias for SCSI_Capacity_Data.
type ScsiCapacityData = SCSI_Capacity_Data

// ScsiCapacityDataLong is a Go-name alias for SCSI_Capacity_Data_Long.
type ScsiCapacityDataLong = SCSI_Capacity_Data_Long

// ScsiSenseData is a Go-name alias for SCSI_Sense_Data.
type ScsiSenseData = SCSI_Sense_Data

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

// EvioSpecialKeyMsg is a Go-name alias for EvioSpecialKeyMsg_t.
type EvioSpecialKeyMsg = EvioSpecialKeyMsg_t

// IOConnect is a Go-name alias for Io_connect_t.
type IOConnect = Io_connect_t

// IOEnumerator is a Go-name alias for Io_enumerator_t.
type IOEnumerator = Io_enumerator_t

// IOIdent is a Go-name alias for Io_ident_t.
type IOIdent = Io_ident_t

// IOIterator is a Go-name alias for Io_iterator_t.
type IOIterator = Io_iterator_t

// IOObject is a Go-name alias for Io_object_t.
type IOObject = Io_object_t

// IORegistryEntry is a Go-name alias for Io_registry_entry_t.
type IORegistryEntry = Io_registry_entry_t

// IOService is a Go-name alias for Io_service_t.
type IOService = Io_service_t

// UextObject is a Go-name alias for Uext_object_t.
type UextObject = Uext_object_t

// UserShspeed is a Go-name alias for User_shspeed_t.
type UserShspeed = User_shspeed_t

// UserSpeed is a Go-name alias for User_speed_t.
type UserSpeed = User_speed_t

// UserUl is a Go-name alias for User_ul_t.
type UserUl = User_ul_t

// UserUs is a Go-name alias for User_us_t.
type UserUs = User_us_t
