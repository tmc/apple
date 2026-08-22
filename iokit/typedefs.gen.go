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
// ATASMARTData is opaque storage with the size and alignment C gives ATASMARTData:
// 512 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 512 into.
type ATASMARTData [512]byte

// See: https://developer.apple.com/documentation/iokit/atasmartdatathresholds
// ATASMARTDataThresholds is opaque storage with the size and alignment C gives ATASMARTDataThresholds:
// 512 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 512 into.
type ATASMARTDataThresholds [512]byte

// See: https://developer.apple.com/documentation/iokit/atasmartlogdirectory
// ATASMARTLogDirectory is opaque storage with the size and alignment C gives ATASMARTLogDirectory:
// 512 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 512 into.
type ATASMARTLogDirectory [512]byte

// See: https://developer.apple.com/documentation/iokit/atasmartlogentry
// ATASMARTLogEntry is opaque storage with the size and alignment C gives ATASMARTLogEntry:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type ATASMARTLogEntry [2]byte

// See: https://developer.apple.com/documentation/iokit/avidtype
type AVIDType = uint32

// See: https://developer.apple.com/documentation/iokit/bddiscinfo
// BDDiscInfo is opaque storage with the size and alignment C gives BDDiscInfo:
// 34 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 34 into.
type BDDiscInfo [34]byte

// See: https://developer.apple.com/documentation/iokit/bdfeatures
type BDFeatures = uint32

// See: https://developer.apple.com/documentation/iokit/bdmediatype
type BDMediaType = uint32

// See: https://developer.apple.com/documentation/iokit/bdtrackinfo
// BDTrackInfo is opaque storage with the size and alignment C gives BDTrackInfo:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type BDTrackInfo [36]byte

// See: https://developer.apple.com/documentation/iokit/block0
// Block0 is opaque storage with the size and alignment C gives Block0:
// 512 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 512 into.
type Block0 [512]byte

// See: https://developer.apple.com/documentation/iokit/cdatip
// CDATIP is opaque storage with the size and alignment C gives CDATIP:
// 28 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 28 into.
type CDATIP [28]byte

// See: https://developer.apple.com/documentation/iokit/cdaudiostatus
// CDAudioStatus is opaque storage with the size and alignment C gives CDAudioStatus:
// 9 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 9 into.
type CDAudioStatus [9]byte

// See: https://developer.apple.com/documentation/iokit/cddiscinfo
// CDDiscInfo is opaque storage with the size and alignment C gives CDDiscInfo:
// 34 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 34 into.
type CDDiscInfo [34]byte

// See: https://developer.apple.com/documentation/iokit/cdfeatures
type CDFeatures = uint32

// See: https://developer.apple.com/documentation/iokit/cdisrc
type CDISRC = int8

// See: https://developer.apple.com/documentation/iokit/cdmcn
type CDMCN = int8

// See: https://developer.apple.com/documentation/iokit/cdmsf
// CDMSF is opaque storage with the size and alignment C gives CDMSF:
// 3 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 3 into.
type CDMSF [3]byte

// See: https://developer.apple.com/documentation/iokit/cdmediatype
type CDMediaType = uint32

// See: https://developer.apple.com/documentation/iokit/cdpma
// CDPMA is opaque storage with the size and alignment C gives CDPMA:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type CDPMA [4]byte

// See: https://developer.apple.com/documentation/iokit/cdpmadescriptor
// CDPMADescriptor is opaque storage with the size and alignment C gives CDPMADescriptor:
// 11 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 11 into.
type CDPMADescriptor [11]byte

// See: https://developer.apple.com/documentation/iokit/cdsectorarea
type CDSectorArea = uint32

// See: https://developer.apple.com/documentation/iokit/cdsectorsize
type CDSectorSize = uint32

// See: https://developer.apple.com/documentation/iokit/cdsectortype
type CDSectorType = uint32

// See: https://developer.apple.com/documentation/iokit/cdtext
// CDTEXT is opaque storage with the size and alignment C gives CDTEXT:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type CDTEXT [4]byte

// See: https://developer.apple.com/documentation/iokit/cdtextdescriptor
// CDTEXTDescriptor is opaque storage with the size and alignment C gives CDTEXTDescriptor:
// 18 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 18 into.
type CDTEXTDescriptor [18]byte

// See: https://developer.apple.com/documentation/iokit/cdtoc
// CDTOC is opaque storage with the size and alignment C gives CDTOC:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type CDTOC [4]byte

// See: https://developer.apple.com/documentation/iokit/cdtocdescriptor
// CDTOCDescriptor is opaque storage with the size and alignment C gives CDTOCDescriptor:
// 11 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 11 into.
type CDTOCDescriptor [11]byte

// See: https://developer.apple.com/documentation/iokit/cdtocformat
type CDTOCFormat = uint8

// See: https://developer.apple.com/documentation/iokit/cdtrackinfo
// CDTrackInfo is opaque storage with the size and alignment C gives CDTrackInfo:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type CDTrackInfo [36]byte

// See: https://developer.apple.com/documentation/iokit/cdtrackinfoaddresstype
type CDTrackInfoAddressType = uint8

// See: https://developer.apple.com/documentation/iokit/csrnodeuniqueid
type CSRNodeUniqueID = uint64

// See: https://developer.apple.com/documentation/iokit/colorspec
// ColorSpec is opaque storage with the size and alignment C gives ColorSpec:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type ColorSpec [4]uint16

// See: https://developer.apple.com/documentation/iokit/colorspecptr
type ColorSpecPtr = *applicationservices.ColorSpec

// See: https://developer.apple.com/documentation/iokit/dasdmodeparameterblockdescriptor
// DASDModeParameterBlockDescriptor is opaque storage with the size and alignment C gives DASDModeParameterBlockDescriptor:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type DASDModeParameterBlockDescriptor [8]byte

// See: https://developer.apple.com/documentation/iokit/dclcallcommandproc
type DCLCallCommandProc = *kernel.ID

// See: https://developer.apple.com/documentation/iokit/dclcallcommandprocptr
type DCLCallCommandProcPtr = *DCLCallCommandProc

// See: https://developer.apple.com/documentation/iokit/dclcallproc
// DCLCallProc is opaque storage with the size and alignment C gives DCLCallProc:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type DCLCallProc [4]uint64

// See: https://developer.apple.com/documentation/iokit/dclcallprocdatatype
type DCLCallProcDataType = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/dclcallprocptr
type DCLCallProcPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/dclcommand
// DCLCommand is opaque storage with the size and alignment C gives DCLCommand:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type DCLCommand [3]uint64

// See: https://developer.apple.com/documentation/iokit/dclcommandptr
type DCLCommandPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/dclcompilerdatatype
type DCLCompilerDataType = uint32

// See: https://developer.apple.com/documentation/iokit/dcljump
// DCLJump is opaque storage with the size and alignment C gives DCLJump:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type DCLJump [3]uint64

// See: https://developer.apple.com/documentation/iokit/dcljumpptr
type DCLJumpPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/dcllabel
// DCLLabel is opaque storage with the size and alignment C gives DCLLabel:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type DCLLabel [2]uint64

// See: https://developer.apple.com/documentation/iokit/dcllabelptr
type DCLLabelPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/dclnudclleader
// DCLNuDCLLeader is opaque storage with the size and alignment C gives DCLNuDCLLeader:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type DCLNuDCLLeader [3]uint64

// See: https://developer.apple.com/documentation/iokit/dclptrtimestamp
// DCLPtrTimeStamp is opaque storage with the size and alignment C gives DCLPtrTimeStamp:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type DCLPtrTimeStamp [3]uint64

// See: https://developer.apple.com/documentation/iokit/dclptrtimestampptr
type DCLPtrTimeStampPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/dclsettagsyncbits
// DCLSetTagSyncBits is opaque storage with the size and alignment C gives DCLSetTagSyncBits:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type DCLSetTagSyncBits [3]uint64

// See: https://developer.apple.com/documentation/iokit/dclsettagsyncbitsptr
type DCLSetTagSyncBitsPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/dcltimestamp
// DCLTimeStamp is opaque storage with the size and alignment C gives DCLTimeStamp:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type DCLTimeStamp [3]uint64

// See: https://developer.apple.com/documentation/iokit/dcltimestampptr
type DCLTimeStampPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/dcltransferbuffer
// DCLTransferBuffer is opaque storage with the size and alignment C gives DCLTransferBuffer:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type DCLTransferBuffer [5]uint64

// See: https://developer.apple.com/documentation/iokit/dcltransferbufferptr
type DCLTransferBufferPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/dcltransferpacket
// DCLTransferPacket is opaque storage with the size and alignment C gives DCLTransferPacket:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type DCLTransferPacket [4]uint64

// See: https://developer.apple.com/documentation/iokit/dcltransferpacketptr
type DCLTransferPacketPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/dclupdatedcllist
// DCLUpdateDCLList is opaque storage with the size and alignment C gives DCLUpdateDCLList:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type DCLUpdateDCLList [4]uint64

// See: https://developer.apple.com/documentation/iokit/dclupdatedcllistptr
type DCLUpdateDCLListPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/ddmap
// DDMap is opaque storage with the size and alignment C gives DDMap:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type DDMap [8]byte

// See: https://developer.apple.com/documentation/iokit/dpme
// DPME is opaque storage with the size and alignment C gives DPME:
// 512 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 512 into.
type DPME [512]byte

// See: https://developer.apple.com/documentation/iokit/dvdauthenticationgrantidinfo
// DVDAuthenticationGrantIDInfo is opaque storage with the size and alignment C gives DVDAuthenticationGrantIDInfo:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type DVDAuthenticationGrantIDInfo [8]byte

// See: https://developer.apple.com/documentation/iokit/dvdauthenticationsuccessflaginfo
// DVDAuthenticationSuccessFlagInfo is opaque storage with the size and alignment C gives DVDAuthenticationSuccessFlagInfo:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type DVDAuthenticationSuccessFlagInfo [8]byte

// See: https://developer.apple.com/documentation/iokit/dvdbooktype
type DVDBookType = uint8

// See: https://developer.apple.com/documentation/iokit/dvdcprmregioncode
type DVDCPRMRegionCode = uint8

// See: https://developer.apple.com/documentation/iokit/dvdchallengekeyinfo
// DVDChallengeKeyInfo is opaque storage with the size and alignment C gives DVDChallengeKeyInfo:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type DVDChallengeKeyInfo [16]byte

// See: https://developer.apple.com/documentation/iokit/dvdcopyrightinfo
// DVDCopyrightInfo is opaque storage with the size and alignment C gives DVDCopyrightInfo:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type DVDCopyrightInfo [8]byte

// See: https://developer.apple.com/documentation/iokit/dvddiscinfo
// DVDDiscInfo is opaque storage with the size and alignment C gives DVDDiscInfo:
// 34 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 34 into.
type DVDDiscInfo [34]byte

// See: https://developer.apple.com/documentation/iokit/dvddisckeyinfo
// DVDDiscKeyInfo is opaque storage with the size and alignment C gives DVDDiscKeyInfo:
// 2052 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2052 into.
type DVDDiscKeyInfo [2052]byte

// See: https://developer.apple.com/documentation/iokit/dvdfeatures
type DVDFeatures = uint32

// See: https://developer.apple.com/documentation/iokit/dvdkey1info
// DVDKey1Info is opaque storage with the size and alignment C gives DVDKey1Info:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type DVDKey1Info [12]byte

// See: https://developer.apple.com/documentation/iokit/dvdkey2info
// DVDKey2Info is opaque storage with the size and alignment C gives DVDKey2Info:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type DVDKey2Info [12]byte

// See: https://developer.apple.com/documentation/iokit/dvdkeyclass
type DVDKeyClass = uint8

// See: https://developer.apple.com/documentation/iokit/dvdkeyformat
type DVDKeyFormat = uint8

// See: https://developer.apple.com/documentation/iokit/dvdmanufacturinginfo
// DVDManufacturingInfo is opaque storage with the size and alignment C gives DVDManufacturingInfo:
// 2052 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2052 into.
type DVDManufacturingInfo [2052]byte

// See: https://developer.apple.com/documentation/iokit/dvdmediatype
type DVDMediaType = uint32

// See: https://developer.apple.com/documentation/iokit/dvdphysicalformatinfo
// DVDPhysicalFormatInfo is opaque storage with the size and alignment C gives DVDPhysicalFormatInfo:
// 2052 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2052 into.
type DVDPhysicalFormatInfo [2052]byte

// See: https://developer.apple.com/documentation/iokit/dvdrzoneinfo
// DVDRZoneInfo is opaque storage with the size and alignment C gives DVDRZoneInfo:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type DVDRZoneInfo [36]byte

// See: https://developer.apple.com/documentation/iokit/dvdrzoneinfoaddresstype
type DVDRZoneInfoAddressType = uint8

// See: https://developer.apple.com/documentation/iokit/dvdregionplaybackcontrolinfo
// DVDRegionPlaybackControlInfo is opaque storage with the size and alignment C gives DVDRegionPlaybackControlInfo:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type DVDRegionPlaybackControlInfo [8]byte

// See: https://developer.apple.com/documentation/iokit/dvdregionalplaybackcontrolscheme
type DVDRegionalPlaybackControlScheme = uint8

// See: https://developer.apple.com/documentation/iokit/dvdstructureformat
type DVDStructureFormat = uint8

// See: https://developer.apple.com/documentation/iokit/dvdtitlekeyinfo
// DVDTitleKeyInfo is opaque storage with the size and alignment C gives DVDTitleKeyInfo:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type DVDTitleKeyInfo [12]byte

// See: https://developer.apple.com/documentation/iokit/depthmode
type DepthMode = uint16

// See: https://developer.apple.com/documentation/iokit/displayidtype
type DisplayIDType = uint32

// See: https://developer.apple.com/documentation/iokit/displaymodeid
type DisplayModeID = int32

// See: https://developer.apple.com/documentation/iokit/evcmd
type EvCmd = uint32

// See: https://developer.apple.com/documentation/iokit/evglobals
// EvGlobals is opaque storage with the size and alignment C gives EvGlobals:
// 23264 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 23264 into.
type EvGlobals [5816]uint32

// See: https://developer.apple.com/documentation/iokit/evoffsets
// EvOffsets is opaque storage with the size and alignment C gives EvOffsets:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type EvOffsets [2]uint32

// See: https://developer.apple.com/documentation/iokit/extendedsensecode
type ExtendedSenseCode = uint8

// See: https://developer.apple.com/documentation/iokit/fwaddress
// FWAddress is opaque storage with the size and alignment C gives FWAddress:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type FWAddress [2]uint32

// See: https://developer.apple.com/documentation/iokit/fwaddressptr
type FWAddressPtr = uintptr

// See: https://developer.apple.com/documentation/iokit/fwaddressspaceflags
type FWAddressSpaceFlags = uint32

// See: https://developer.apple.com/documentation/iokit/fwclientcommandid
type FWClientCommandID = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/fwsbp2logincompleteparams
// FWSBP2LoginCompleteParams is opaque storage with the size and alignment C gives FWSBP2LoginCompleteParams:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type FWSBP2LoginCompleteParams [5]uint64

// See: https://developer.apple.com/documentation/iokit/fwsbp2loginresponse
// FWSBP2LoginResponse is opaque storage with the size and alignment C gives FWSBP2LoginResponse:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type FWSBP2LoginResponse [4]uint32

// See: https://developer.apple.com/documentation/iokit/fwsbp2logoutcompleteparams
// FWSBP2LogoutCompleteParams is opaque storage with the size and alignment C gives FWSBP2LogoutCompleteParams:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type FWSBP2LogoutCompleteParams [4]uint64

// See: https://developer.apple.com/documentation/iokit/fwsbp2notifyparams
// FWSBP2NotifyParams is opaque storage with the size and alignment C gives FWSBP2NotifyParams:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type FWSBP2NotifyParams [4]uint64

// See: https://developer.apple.com/documentation/iokit/fwsbp2reconnectparams
// FWSBP2ReconnectParams is opaque storage with the size and alignment C gives FWSBP2ReconnectParams:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type FWSBP2ReconnectParams [4]uint64

// See: https://developer.apple.com/documentation/iokit/fwsbp2statusblock
// FWSBP2StatusBlock is opaque storage with the size and alignment C gives FWSBP2StatusBlock:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type FWSBP2StatusBlock [8]uint32

// See: https://developer.apple.com/documentation/iokit/fwsbp2virtualrange
// FWSBP2VirtualRange is opaque storage with the size and alignment C gives FWSBP2VirtualRange:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type FWSBP2VirtualRange [2]uint64

// See: https://developer.apple.com/documentation/iokit/gammatableid
type GammaTableID = uint32

// See: https://developer.apple.com/documentation/iokit/gammatbl
// GammaTbl is opaque storage with the size and alignment C gives GammaTbl:
// 14 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 14 into.
type GammaTbl [7]uint16

// See: https://developer.apple.com/documentation/iokit/gammatblptr
type GammaTblPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/hidreportcommandtype
type HIDReportCommandType = uint32

// IOATASMARTInterface is self-Monitoring, Analysis, and Reporting Technology Interface.
//
// See: https://developer.apple.com/documentation/iokit/ioatasmartinterface
// IOATASMARTInterface is opaque storage with the size and alignment C gives IOATASMARTInterface:
// 128 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 128 into.
type IOATASMARTInterface [16]uint64

// See: https://developer.apple.com/documentation/iokit/ioavccommandresponse
type IOAVCCommandResponse = uint32

// See: https://developer.apple.com/documentation/iokit/ioavcframefields
type IOAVCFrameFields = uint32

// See: https://developer.apple.com/documentation/iokit/ioavcopcodes
type IOAVCOpcodes = uint32

// See: https://developer.apple.com/documentation/iokit/ioavcunittypes
type IOAVCUnitTypes = uint32

// See: https://developer.apple.com/documentation/iokit/ioaccelbounds
// IOAccelBounds is opaque storage with the size and alignment C gives IOAccelBounds:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOAccelBounds [4]uint16

// See: https://developer.apple.com/documentation/iokit/ioacceldeviceregion
// IOAccelDeviceRegion is opaque storage with the size and alignment C gives IOAccelDeviceRegion:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type IOAccelDeviceRegion [3]uint32

// See: https://developer.apple.com/documentation/iokit/ioaccelid
type IOAccelID = int32

// See: https://developer.apple.com/documentation/iokit/ioaccelsize
// IOAccelSize is opaque storage with the size and alignment C gives IOAccelSize:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type IOAccelSize [2]uint16

// See: https://developer.apple.com/documentation/iokit/ioaccelsurfaceinformation
// IOAccelSurfaceInformation is opaque storage with the size and alignment C gives IOAccelSurfaceInformation:
// 88 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 88 into.
type IOAccelSurfaceInformation [11]uint64

// See: https://developer.apple.com/documentation/iokit/ioaccelsurfacereaddata
// IOAccelSurfaceReadData is opaque storage with the size and alignment C gives IOAccelSurfaceReadData:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type IOAccelSurfaceReadData [4]uint64

// See: https://developer.apple.com/documentation/iokit/ioaccelsurfacescaling
// IOAccelSurfaceScaling is opaque storage with the size and alignment C gives IOAccelSurfaceScaling:
// 44 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 44 into.
type IOAccelSurfaceScaling [11]uint32

// See: https://developer.apple.com/documentation/iokit/ioaddressrange
type IOAddressRange = IOVirtualRange

// See: https://developer.apple.com/documentation/iokit/ioalignment
type IOAlignment = uint32

// See: https://developer.apple.com/documentation/iokit/ioappletimingid
type IOAppleTimingID = uint32

// IOAsyncCallback is standard callback function for asynchronous I/O requests with lots of extra arguments beyond a refcon and result code.
//
// See: https://developer.apple.com/documentation/iokit/ioasynccallback
type IOAsyncCallback = func(refcon unsafe.Pointer, result int32, args unsafe.Pointer, numArgs uint32)

// IOAsyncCallback0 is standard callback function for asynchronous I/O requests with no extra arguments beyond a refcon and result code.
//
// See: https://developer.apple.com/documentation/iokit/ioasynccallback0
type IOAsyncCallback0 = func(refcon unsafe.Pointer, result int32)

// IOAsyncCallback1 is standard callback function for asynchronous I/O requests with one extra argument beyond a refcon and result code. This is often a count of the number of bytes transferred.
//
// See: https://developer.apple.com/documentation/iokit/ioasynccallback1
type IOAsyncCallback1 = func(refcon unsafe.Pointer, result int32, arg0 unsafe.Pointer)

// IOAsyncCallback2 is standard callback function for asynchronous I/O requests with two extra arguments beyond a refcon and result code.
//
// See: https://developer.apple.com/documentation/iokit/ioasynccallback2
type IOAsyncCallback2 = func(refcon unsafe.Pointer, result int32, arg0 unsafe.Pointer, arg1 unsafe.Pointer)

// See: https://developer.apple.com/documentation/iokit/ioaudiobufferdatadescriptor
// IOAudioBufferDataDescriptor is opaque storage with the size and alignment C gives IOAudioBufferDataDescriptor:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type IOAudioBufferDataDescriptor [5]uint32

// IOAudioControlCalls is the set of constants passed to IOAudioControlUserClient::getExternalMethodForIndex() when making calls from the IOAudioFamily user client code.
//
// See: https://developer.apple.com/documentation/iokit/ioaudiocontrolcalls
type IOAudioControlCalls = uint32

// IOAudioControlNotifications is the set of constants passed in the type field of IOAudioControlUserClient::registerNotificaitonPort().
//
// See: https://developer.apple.com/documentation/iokit/ioaudiocontrolnotifications
type IOAudioControlNotifications = uint32

// IOAudioEngineCalls is the set of constants passed to IOAudioEngineUserClient::getExternalMethodForIndex() when making calls from the IOAudioFamily user client code.
//
// See: https://developer.apple.com/documentation/iokit/ioaudioenginecalls
type IOAudioEngineCalls = uint32

// IOAudioEngineMemory is used to identify the type of memory requested by a client process to be mapped into its process space.
//
// See: https://developer.apple.com/documentation/iokit/ioaudioenginememory
type IOAudioEngineMemory = uint32

// See: https://developer.apple.com/documentation/iokit/ioaudioenginenotifications
type IOAudioEngineNotifications = uint32

// IOAudioEngineState is represents the state of an IOAudioEngine.
//
// See: https://developer.apple.com/documentation/iokit/ioaudioenginestate
type IOAudioEngineState = uint32

// IOAudioEngineStatus is shared-memory structure giving audio engine status.
//
// See: https://developer.apple.com/documentation/iokit/ioaudioenginestatus
// IOAudioEngineStatus is opaque storage with the size and alignment C gives IOAudioEngineStatus:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type IOAudioEngineStatus [5]uint32

// See: https://developer.apple.com/documentation/iokit/ioaudioenginetraps
type IOAudioEngineTraps = uint32

// IOAudioNotificationMessage is used in the mach message for IOAudio notifications.
//
// See: https://developer.apple.com/documentation/iokit/ioaudionotificationmessage
// IOAudioNotificationMessage is opaque storage with the size and alignment C gives IOAudioNotificationMessage:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type IOAudioNotificationMessage [5]uint64

// IOAudioSMPTETime is a structure for holding a SMPTE time.
//
// See: https://developer.apple.com/documentation/iokit/ioaudiosmptetime
// IOAudioSMPTETime is opaque storage with the size and alignment C gives IOAudioSMPTETime:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IOAudioSMPTETime [6]uint32

// See: https://developer.apple.com/documentation/iokit/ioaudiosampleintervaldescriptor
// IOAudioSampleIntervalDescriptor is opaque storage with the size and alignment C gives IOAudioSampleIntervalDescriptor:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOAudioSampleIntervalDescriptor [2]uint32

// See: https://developer.apple.com/documentation/iokit/ioaudiosamplerate
// IOAudioSampleRate is opaque storage with the size and alignment C gives IOAudioSampleRate:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOAudioSampleRate [2]uint32

// See: https://developer.apple.com/documentation/iokit/ioaudiostreamdatadescriptor
// IOAudioStreamDataDescriptor is opaque storage with the size and alignment C gives IOAudioStreamDataDescriptor:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type IOAudioStreamDataDescriptor [3]uint32

// IOAudioStreamDirection is represents the direction of an IOAudioStream.
//
// See: https://developer.apple.com/documentation/iokit/ioaudiostreamdirection
type IOAudioStreamDirection = uint32

// See: https://developer.apple.com/documentation/iokit/ioaudiostreamformat
// IOAudioStreamFormat is opaque storage with the size and alignment C gives IOAudioStreamFormat:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IOAudioStreamFormat [6]uint32

// See: https://developer.apple.com/documentation/iokit/ioaudiostreamformatextension
// IOAudioStreamFormatExtension is opaque storage with the size and alignment C gives IOAudioStreamFormatExtension:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type IOAudioStreamFormatExtension [4]uint32

// See: https://developer.apple.com/documentation/iokit/ioaudiotimestamp
// IOAudioTimeStamp is opaque storage with the size and alignment C gives IOAudioTimeStamp:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type IOAudioTimeStamp [8]uint64

// See: https://developer.apple.com/documentation/iokit/ioblitcompletiontoken
type IOBlitCompletionToken = int32

// See: https://developer.apple.com/documentation/iokit/ioblitcopyrectangle
// IOBlitCopyRectangle is opaque storage with the size and alignment C gives IOBlitCopyRectangle:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IOBlitCopyRectangle [6]uint32

// See: https://developer.apple.com/documentation/iokit/ioblitcopyrectangles
// IOBlitCopyRectangles is opaque storage with the size and alignment C gives IOBlitCopyRectangles:
// 116 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 116 into.
type IOBlitCopyRectangles [29]uint32

// See: https://developer.apple.com/documentation/iokit/ioblitcopyregion
// IOBlitCopyRegion is opaque storage with the size and alignment C gives IOBlitCopyRegion:
// 104 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 104 into.
type IOBlitCopyRegion [13]uint64

// See: https://developer.apple.com/documentation/iokit/ioblitcursor
// IOBlitCursor is opaque storage with the size and alignment C gives IOBlitCursor:
// 104 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 104 into.
type IOBlitCursor [26]uint32

// See: https://developer.apple.com/documentation/iokit/ioblitmemory
// IOBlitMemory is opaque storage with the size and alignment C gives IOBlitMemory:
// 120 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 120 into.
type IOBlitMemory [15]uint64

// See: https://developer.apple.com/documentation/iokit/ioblitmemoryref
type IOBlitMemoryRef uintptr

// See: https://developer.apple.com/documentation/iokit/ioblitoperation
// IOBlitOperation is opaque storage with the size and alignment C gives IOBlitOperation:
// 88 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 88 into.
type IOBlitOperation [22]uint32

// See: https://developer.apple.com/documentation/iokit/ioblitrectangle
// IOBlitRectangle is opaque storage with the size and alignment C gives IOBlitRectangle:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type IOBlitRectangle [4]uint32

// See: https://developer.apple.com/documentation/iokit/ioblitrectangles
// IOBlitRectangles is opaque storage with the size and alignment C gives IOBlitRectangles:
// 108 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 108 into.
type IOBlitRectangles [27]uint32

// See: https://developer.apple.com/documentation/iokit/ioblitscanlines
// IOBlitScanlines is opaque storage with the size and alignment C gives IOBlitScanlines:
// 108 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 108 into.
type IOBlitScanlines [27]uint32

// See: https://developer.apple.com/documentation/iokit/ioblitsourcedesttype
type IOBlitSourceDestType = uint32

// See: https://developer.apple.com/documentation/iokit/ioblitsourcetype
type IOBlitSourceType = uint32

// See: https://developer.apple.com/documentation/iokit/ioblitsurface
// IOBlitSurface is opaque storage with the size and alignment C gives IOBlitSurface:
// 120 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 120 into.
type IOBlitSurface [15]uint64

// See: https://developer.apple.com/documentation/iokit/ioblittype
type IOBlitType = uint32

// See: https://developer.apple.com/documentation/iokit/ioblitvertex
// IOBlitVertex is opaque storage with the size and alignment C gives IOBlitVertex:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOBlitVertex [2]uint32

// See: https://developer.apple.com/documentation/iokit/ioblitvertices
// IOBlitVertices is opaque storage with the size and alignment C gives IOBlitVertices:
// 108 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 108 into.
type IOBlitVertices [27]uint32

// See: https://developer.apple.com/documentation/iokit/iobytecount
type IOByteCount = uint

// See: https://developer.apple.com/documentation/iokit/iobytecount32
type IOByteCount32 = uint32

// See: https://developer.apple.com/documentation/iokit/iobytecount64
type IOByteCount64 = uint64

// See: https://developer.apple.com/documentation/iokit/iocfplugininterfacestruct
// IOCFPlugInInterface is opaque storage with the size and alignment C gives IOCFPlugInInterface:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type IOCFPlugInInterface [8]uint64

// See: https://developer.apple.com/documentation/iokit/iocsrkeytype
type IOCSRKeyType = uint32

// See: https://developer.apple.com/documentation/iokit/iocachemode
type IOCacheMode = uint32

// See: https://developer.apple.com/documentation/iokit/iocolorcomponent
type IOColorComponent = uint16

// See: https://developer.apple.com/documentation/iokit/iocolorentry
// IOColorEntry is opaque storage with the size and alignment C gives IOColorEntry:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOColorEntry [4]uint16

// See: https://developer.apple.com/documentation/iokit/ioconfigkeytype
type IOConfigKeyType = uint32

// IODataQueueAppendix is a struct mapping to the appendix region of a data queue.
//
// See: https://developer.apple.com/documentation/iokit/iodataqueueappendix
// IODataQueueAppendix is opaque storage with the size and alignment C gives IODataQueueAppendix:
// 28 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 28 into.
type IODataQueueAppendix [7]uint32

// IODataQueueEntry is represents an entry within the data queue.
//
// See: https://developer.apple.com/documentation/iokit/iodataqueueentry
// IODataQueueEntry is opaque storage with the size and alignment C gives IODataQueueEntry:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IODataQueueEntry [2]uint32

// IODataQueueMemory is a struct mapping to the header region of a data queue.
//
// See: https://developer.apple.com/documentation/iokit/iodataqueuememory
// IODataQueueMemory is opaque storage with the size and alignment C gives IODataQueueMemory:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type IODataQueueMemory [5]uint32

// See: https://developer.apple.com/documentation/iokit/iodetailedtiminginformation
type IODetailedTimingInformation = IODetailedTimingInformationV2

// See: https://developer.apple.com/documentation/iokit/iodetailedtiminginformationv1
// IODetailedTimingInformationV1 is opaque storage with the size and alignment C gives IODetailedTimingInformationV1:
// 44 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 44 into.
type IODetailedTimingInformationV1 [11]uint32

// See: https://developer.apple.com/documentation/iokit/iodetailedtiminginformationv2
// IODetailedTimingInformationV2 is opaque storage with the size and alignment C gives IODetailedTimingInformationV2:
// 160 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 160 into.
type IODetailedTimingInformationV2 [20]uint64

// See: https://developer.apple.com/documentation/iokit/iodevicenumber
type IODeviceNumber = uint32

// See: https://developer.apple.com/documentation/iokit/iodisplaymodeid
type IODisplayModeID = int32

// See: https://developer.apple.com/documentation/iokit/iodisplaymodeinformation
// IODisplayModeInformation is opaque storage with the size and alignment C gives IODisplayModeInformation:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type IODisplayModeInformation [9]uint32

// See: https://developer.apple.com/documentation/iokit/iodisplayproductid
type IODisplayProductID = uint32

// See: https://developer.apple.com/documentation/iokit/iodisplayscalerinformation
// IODisplayScalerInformation is opaque storage with the size and alignment C gives IODisplayScalerInformation:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type IODisplayScalerInformation [12]uint32

// See: https://developer.apple.com/documentation/iokit/iodisplaytimingrange
type IODisplayTimingRange = IODisplayTimingRangeV2

// See: https://developer.apple.com/documentation/iokit/iodisplaytimingrangev1
// IODisplayTimingRangeV1 is opaque storage with the size and alignment C gives IODisplayTimingRangeV1:
// 240 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 240 into.
type IODisplayTimingRangeV1 [30]uint64

// See: https://developer.apple.com/documentation/iokit/iodisplaytimingrangev2
// IODisplayTimingRangeV2 is opaque storage with the size and alignment C gives IODisplayTimingRangeV2:
// 312 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 312 into.
type IODisplayTimingRangeV2 [39]uint64

// See: https://developer.apple.com/documentation/iokit/iodisplayvendorid
type IODisplayVendorID = uint32

// See: https://developer.apple.com/documentation/iokit/iodot3collentry
// IODot3CollEntry is opaque storage with the size and alignment C gives IODot3CollEntry:
// 1 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 1 into.
type IODot3CollEntry [1]byte

// See: https://developer.apple.com/documentation/iokit/iodot3rxextraentry
// IODot3RxExtraEntry is opaque storage with the size and alignment C gives IODot3RxExtraEntry:
// 1 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 1 into.
type IODot3RxExtraEntry [1]byte

// See: https://developer.apple.com/documentation/iokit/iodot3statsentry
// IODot3StatsEntry is opaque storage with the size and alignment C gives IODot3StatsEntry:
// 1 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 1 into.
type IODot3StatsEntry [1]byte

// See: https://developer.apple.com/documentation/iokit/iodot3txextraentry
// IODot3TxExtraEntry is opaque storage with the size and alignment C gives IODot3TxExtraEntry:
// 1 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 1 into.
type IODot3TxExtraEntry [1]byte

// See: https://developer.apple.com/documentation/iokit/ioethernetstats
// IOEthernetStats is opaque storage with the size and alignment C gives IOEthernetStats:
// 1 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 1 into.
type IOEthernetStats [1]byte

// See: https://developer.apple.com/documentation/iokit/iofbdplinkconfig
// IOFBDPLinkConfig is opaque storage with the size and alignment C gives IOFBDPLinkConfig:
// 28 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 28 into.
type IOFBDPLinkConfig [14]uint16

// See: https://developer.apple.com/documentation/iokit/iofbdisplaymodedescription
// IOFBDisplayModeDescription is opaque storage with the size and alignment C gives IOFBDisplayModeDescription:
// 204 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 204 into.
type IOFBDisplayModeDescription [51]uint32

// See: https://developer.apple.com/documentation/iokit/iofbhdrmetadatav1
// IOFBHDRMetaDataV1 is opaque storage with the size and alignment C gives IOFBHDRMetaDataV1:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type IOFBHDRMetaDataV1 [8]uint64

// See: https://developer.apple.com/documentation/iokit/iofwavcasynccommandstate
type IOFWAVCAsyncCommandState = uint32

// See: https://developer.apple.com/documentation/iokit/iofwavcplugtypes
type IOFWAVCPlugTypes = uint32

// See: https://developer.apple.com/documentation/iokit/iofwavcsubunitplugmessages
type IOFWAVCSubunitPlugMessages = uint32

// IOFWAsyncStreamListenerInterface is represents and provides management functions for a asyn stream listener object.
//
// See: https://developer.apple.com/documentation/iokit/iofwasyncstreamlistenerinterface
// IOFWAsyncStreamListenerInterface is opaque storage with the size and alignment C gives IOFWAsyncStreamListenerInterface:
// 120 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 120 into.
type IOFWAsyncStreamListenerInterface [15]uint64

// See: https://developer.apple.com/documentation/iokit/iofwasyncstreamlistenerinterfaceref
type IOFWAsyncStreamListenerInterfaceRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofwdclnotificationtype
type IOFWDCLNotificationType = uint32

// See: https://developer.apple.com/documentation/iokit/iofwisochportoptions
type IOFWIsochPortOptions = uint32

// See: https://developer.apple.com/documentation/iokit/iofwisochresourceflags
type IOFWIsochResourceFlags = uint32

// See: https://developer.apple.com/documentation/iokit/iofwspeed
type IOFWSpeed = uint32

// See: https://developer.apple.com/documentation/iokit/iofirewireavclibasynchronouscommand
// IOFireWireAVCLibAsynchronousCommand is opaque storage with the size and alignment C gives IOFireWireAVCLibAsynchronousCommand:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type IOFireWireAVCLibAsynchronousCommand [8]uint64

// IOFireWireAVCLibConsumerInterface is interface for an asynchronous connection consumer.
//
// See: https://developer.apple.com/documentation/iokit/iofirewireavclibconsumerinterface
// IOFireWireAVCLibConsumerInterface is opaque storage with the size and alignment C gives IOFireWireAVCLibConsumerInterface:
// 152 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 152 into.
type IOFireWireAVCLibConsumerInterface [19]uint64

// IOFireWireAVCLibProtocolInterface is initial interface discovered for all AVC protocol drivers.
//
// See: https://developer.apple.com/documentation/iokit/iofirewireavclibprotocolinterface
// IOFireWireAVCLibProtocolInterface is opaque storage with the size and alignment C gives IOFireWireAVCLibProtocolInterface:
// 240 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 240 into.
type IOFireWireAVCLibProtocolInterface [30]uint64

// IOFireWireAVCLibUnitInterface is initial interface discovered for all AVC Unit drivers.
//
// See: https://developer.apple.com/documentation/iokit/iofirewireavclibunitinterface
// IOFireWireAVCLibUnitInterface is opaque storage with the size and alignment C gives IOFireWireAVCLibUnitInterface:
// 232 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 232 into.
type IOFireWireAVCLibUnitInterface [29]uint64

// See: https://developer.apple.com/documentation/iokit/iofirewireasyncstreamcommandinterface
// IOFireWireAsyncStreamCommandInterface is opaque storage with the size and alignment C gives IOFireWireAsyncStreamCommandInterface:
// 232 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 232 into.
type IOFireWireAsyncStreamCommandInterface [29]uint64

// IOFireWireCommandInterface is iOFireWireLib command object.
//
// See: https://developer.apple.com/documentation/iokit/iofirewirecommandinterface
// IOFireWireCommandInterface is opaque storage with the size and alignment C gives IOFireWireCommandInterface:
// 208 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 208 into.
type IOFireWireCommandInterface [26]uint64

// See: https://developer.apple.com/documentation/iokit/iofirewirecompareswapcommandinterface
// IOFireWireCompareSwapCommandInterface is opaque storage with the size and alignment C gives IOFireWireCompareSwapCommandInterface:
// 176 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 176 into.
type IOFireWireCompareSwapCommandInterface [22]uint64

// See: https://developer.apple.com/documentation/iokit/iofirewirecompareswapcommandinterface_v3
// IOFireWireCompareSwapCommandInterface_v3 is opaque storage with the size and alignment C gives IOFireWireCompareSwapCommandInterface_v3:
// 248 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 248 into.
type IOFireWireCompareSwapCommandInterface_v3 [31]uint64

// IOFireWireConfigDirectoryInterface is iOFireWireLib device config ROM browsing interface.
//
// See: https://developer.apple.com/documentation/iokit/iofirewireconfigdirectoryinterface
// IOFireWireConfigDirectoryInterface is opaque storage with the size and alignment C gives IOFireWireConfigDirectoryInterface:
// 192 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 192 into.
type IOFireWireConfigDirectoryInterface [24]uint64

// See: https://developer.apple.com/documentation/iokit/iofirewiredclcommandpoolinterface
// IOFireWireDCLCommandPoolInterface is opaque storage with the size and alignment C gives IOFireWireDCLCommandPoolInterface:
// 208 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 208 into.
type IOFireWireDCLCommandPoolInterface [26]uint64

// IOFireWireDeviceInterface is iOFireWireDeviceInterface is your primary gateway to the functionality contained in IOFireWireLib.
//
// See: https://developer.apple.com/documentation/iokit/iofirewiredeviceinterface
// IOFireWireDeviceInterface is opaque storage with the size and alignment C gives IOFireWireDeviceInterface:
// 632 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 632 into.
type IOFireWireDeviceInterface [79]uint64

// IOFireWireIsochChannelInterface is fireWire user client isochronous channel object.
//
// See: https://developer.apple.com/documentation/iokit/iofirewireisochchannelinterface
// IOFireWireIsochChannelInterface is opaque storage with the size and alignment C gives IOFireWireIsochChannelInterface:
// 144 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 144 into.
type IOFireWireIsochChannelInterface [18]uint64

// IOFireWireIsochPortInterface is fireWire user client isochronous port interface.
//
// See: https://developer.apple.com/documentation/iokit/iofirewireisochportinterface
// IOFireWireIsochPortInterface is opaque storage with the size and alignment C gives IOFireWireIsochPortInterface:
// 96 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 96 into.
type IOFireWireIsochPortInterface [12]uint64

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
// IOFireWireLibIRMAllocationInterface is opaque storage with the size and alignment C gives IOFireWireLibIRMAllocationInterface:
// 112 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 112 into.
type IOFireWireLibIRMAllocationInterface [14]uint64

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
// IOFireWireLibPHYPacketListenerInterface is opaque storage with the size and alignment C gives IOFireWireLibPHYPacketListenerInterface:
// 120 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 120 into.
type IOFireWireLibPHYPacketListenerInterface [15]uint64

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
// IOFireWireLibVectorCommandInterface is opaque storage with the size and alignment C gives IOFireWireLibVectorCommandInterface:
// 176 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 176 into.
type IOFireWireLibVectorCommandInterface [22]uint64

// See: https://developer.apple.com/documentation/iokit/iofirewirelibvectorcommandref
type IOFireWireLibVectorCommandRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewirelibwritecommandref
type IOFireWireLibWriteCommandRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewirelibwritequadletcommandref
type IOFireWireLibWriteQuadletCommandRef uintptr

// IOFireWireLocalIsochPortInterface is fireWire user client local isochronous port object.
//
// See: https://developer.apple.com/documentation/iokit/iofirewirelocalisochportinterface
// IOFireWireLocalIsochPortInterface is opaque storage with the size and alignment C gives IOFireWireLocalIsochPortInterface:
// 160 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 160 into.
type IOFireWireLocalIsochPortInterface [20]uint64

// See: https://developer.apple.com/documentation/iokit/iofirewirelocalunitdirectoryinterface
// IOFireWireLocalUnitDirectoryInterface is opaque storage with the size and alignment C gives IOFireWireLocalUnitDirectoryInterface:
// 80 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 80 into.
type IOFireWireLocalUnitDirectoryInterface [10]uint64

// IOFireWireNuDCLPoolInterface is use this interface to build NuDCL-based DCL programs.
//
// See: https://developer.apple.com/documentation/iokit/iofirewirenudclpoolinterface
// IOFireWireNuDCLPoolInterface is opaque storage with the size and alignment C gives IOFireWireNuDCLPoolInterface:
// 424 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 424 into.
type IOFireWireNuDCLPoolInterface [53]uint64

// See: https://developer.apple.com/documentation/iokit/iofirewirenubinterface
// IOFireWireNubInterface is opaque storage with the size and alignment C gives IOFireWireNubInterface:
// 632 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 632 into.
type IOFireWireNubInterface [79]uint64

// See: https://developer.apple.com/documentation/iokit/iofirewirephycommandinterface
// IOFireWirePHYCommandInterface is opaque storage with the size and alignment C gives IOFireWirePHYCommandInterface:
// 216 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 216 into.
type IOFireWirePHYCommandInterface [27]uint64

// IOFireWirePhysicalAddressSpaceInterface is iOFireWireLib physical address space object. ( interface name: IOFireWirePhysicalAddressSpaceInterface ).
//
// See: https://developer.apple.com/documentation/iokit/iofirewirephysicaladdressspaceinterface
// IOFireWirePhysicalAddressSpaceInterface is opaque storage with the size and alignment C gives IOFireWirePhysicalAddressSpaceInterface:
// 88 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 88 into.
type IOFireWirePhysicalAddressSpaceInterface [11]uint64

// See: https://developer.apple.com/documentation/iokit/iofirewirepseudoaddressspaceinterface
// IOFireWirePseudoAddressSpaceInterface is opaque storage with the size and alignment C gives IOFireWirePseudoAddressSpaceInterface:
// 128 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 128 into.
type IOFireWirePseudoAddressSpaceInterface [16]uint64

// IOFireWireReadCommandInterface is iOFireWireLib block read command object.
//
// See: https://developer.apple.com/documentation/iokit/iofirewirereadcommandinterface
// IOFireWireReadCommandInterface is opaque storage with the size and alignment C gives IOFireWireReadCommandInterface:
// 208 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 208 into.
type IOFireWireReadCommandInterface [26]uint64

// IOFireWireReadQuadletCommandInterface is iOFireWireReadQuadletCommandInterface -- IOFireWireLib quadlet read command object.
//
// See: https://developer.apple.com/documentation/iokit/iofirewirereadquadletcommandinterface
// IOFireWireReadQuadletCommandInterface is opaque storage with the size and alignment C gives IOFireWireReadQuadletCommandInterface:
// 136 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 136 into.
type IOFireWireReadQuadletCommandInterface [17]uint64

// See: https://developer.apple.com/documentation/iokit/iofirewireremoteisochportinterface
// IOFireWireRemoteIsochPortInterface is opaque storage with the size and alignment C gives IOFireWireRemoteIsochPortInterface:
// 136 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 136 into.
type IOFireWireRemoteIsochPortInterface [17]uint64

// IOFireWireSBP2LibLUNInterface is initial interface disovered for all drivers.
//
// See: https://developer.apple.com/documentation/iokit/iofirewiresbp2libluninterface
// IOFireWireSBP2LibLUNInterface is opaque storage with the size and alignment C gives IOFireWireSBP2LibLUNInterface:
// 128 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 128 into.
type IOFireWireSBP2LibLUNInterface [16]uint64

// IOFireWireSBP2LibLoginInterface is supplies the login maintenance and Normal Command ORB execution portions of the API.
//
// See: https://developer.apple.com/documentation/iokit/iofirewiresbp2liblogininterface
// IOFireWireSBP2LibLoginInterface is opaque storage with the size and alignment C gives IOFireWireSBP2LibLoginInterface:
// 216 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 216 into.
type IOFireWireSBP2LibLoginInterface [27]uint64

// IOFireWireSBP2LibMgmtORBInterface is supplies non login related management ORBs. Management ORBs can be executed independent of a login, if necessary. Management ORBs are created using the IOFireWireSBP2LibLUNInterface.
//
// See: https://developer.apple.com/documentation/iokit/iofirewiresbp2libmgmtorbinterface
// IOFireWireSBP2LibMgmtORBInterface is opaque storage with the size and alignment C gives IOFireWireSBP2LibMgmtORBInterface:
// 104 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 104 into.
type IOFireWireSBP2LibMgmtORBInterface [13]uint64

// IOFireWireSBP2LibORBInterface is represents an SBP2 normal command ORB. Supplies the APIs for configuring normal command ORBs. This includes setting the command block and writing the page tables for I/O. The ORBs are executed using the submitORB method in IOFireWireSBP2LibLoginInterface.
//
// See: https://developer.apple.com/documentation/iokit/iofirewiresbp2liborbinterface
// IOFireWireSBP2LibORBInterface is opaque storage with the size and alignment C gives IOFireWireSBP2LibORBInterface:
// 136 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 136 into.
type IOFireWireSBP2LibORBInterface [17]uint64

// See: https://developer.apple.com/documentation/iokit/iofirewiresessionref
type IOFireWireSessionRef uintptr

// See: https://developer.apple.com/documentation/iokit/iofirewireunitinterface
// IOFireWireUnitInterface is opaque storage with the size and alignment C gives IOFireWireUnitInterface:
// 632 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 632 into.
type IOFireWireUnitInterface [79]uint64

// IOFireWireWriteCommandInterface is iOFireWireLib block read command object.
//
// See: https://developer.apple.com/documentation/iokit/iofirewirewritecommandinterface
// IOFireWireWriteCommandInterface is opaque storage with the size and alignment C gives IOFireWireWriteCommandInterface:
// 208 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 208 into.
type IOFireWireWriteCommandInterface [26]uint64

// IOFireWireWriteQuadletCommandInterface is iOFireWireLib quadlet read command object.
//
// See: https://developer.apple.com/documentation/iokit/iofirewirewritequadletcommandinterface
// IOFireWireWriteQuadletCommandInterface is opaque storage with the size and alignment C gives IOFireWireWriteQuadletCommandInterface:
// 136 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 136 into.
type IOFireWireWriteQuadletCommandInterface [17]uint64

// See: https://developer.apple.com/documentation/iokit/iofixed
type IOFixed = int32

// See: https://developer.apple.com/documentation/iokit/iofixed1616
type IOFixed1616 = uint32

// See: https://developer.apple.com/documentation/iokit/iofixedpoint32
// IOFixedPoint32 is opaque storage with the size and alignment C gives IOFixedPoint32:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOFixedPoint32 [2]uint32

// See: https://developer.apple.com/documentation/iokit/iofourcharcode
type IOFourCharCode = uint32

// See: https://developer.apple.com/documentation/iokit/ioframebufferinformation
// IOFramebufferInformation is opaque storage with the size and alignment C gives IOFramebufferInformation:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type IOFramebufferInformation [8]uint64

// See: https://developer.apple.com/documentation/iokit/iogbounds
// IOGBounds is opaque storage with the size and alignment C gives IOGBounds:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOGBounds [4]uint16

// See: https://developer.apple.com/documentation/iokit/iogpoint
// IOGPoint is opaque storage with the size and alignment C gives IOGPoint:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type IOGPoint [2]uint16

// See: https://developer.apple.com/documentation/iokit/iogsize
// IOGSize is opaque storage with the size and alignment C gives IOGSize:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type IOGSize [2]uint16

// See: https://developer.apple.com/documentation/iokit/iographicsacceleratorinterface
// IOGraphicsAcceleratorInterface is opaque storage with the size and alignment C gives IOGraphicsAcceleratorInterface:
// 376 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 376 into.
type IOGraphicsAcceleratorInterface [47]uint64

// See: https://developer.apple.com/documentation/iokit/iohidaccelerationalgorithmtype
type IOHIDAccelerationAlgorithmType = uint8

// See: https://developer.apple.com/documentation/iokit/iohidaccesstype
type IOHIDAccessType = uint32

// See: https://developer.apple.com/documentation/iokit/iohidbuttonmodes
type IOHIDButtonModes = uint32

// See: https://developer.apple.com/documentation/iokit/iohidcompletion
// IOHIDCompletion is opaque storage with the size and alignment C gives IOHIDCompletion:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IOHIDCompletion [3]uint64

// IOHIDDeviceDeviceInterface is the object you use to access HID devices from user space, returned by version 1.5 of the IOHIDFamily.
//
// See: https://developer.apple.com/documentation/iokit/iohiddevicedeviceinterface
// IOHIDDeviceDeviceInterface is opaque storage with the size and alignment C gives IOHIDDeviceDeviceInterface:
// 120 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 120 into.
type IOHIDDeviceDeviceInterface [15]uint64

// See: https://developer.apple.com/documentation/iokit/iohiddevicegetvalueoptions
type IOHIDDeviceGetValueOptions = uint32

// IOHIDDeviceInterface is cFPlugin object subclass which provides the primary interface to HID devices.
//
// See: https://developer.apple.com/documentation/iokit/iohiddeviceinterface
// IOHIDDeviceInterface is opaque storage with the size and alignment C gives IOHIDDeviceInterface:
// 160 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 160 into.
type IOHIDDeviceInterface [20]uint64

// IOHIDDeviceInterface121 is cFPlugin object subclass which provides the primary interface to HID devices. This class is a subclass of IOHIDDeviceInterface.
//
// See: https://developer.apple.com/documentation/iokit/iohiddeviceinterface121
// IOHIDDeviceInterface121 is opaque storage with the size and alignment C gives IOHIDDeviceInterface121:
// 160 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 160 into.
type IOHIDDeviceInterface121 [20]uint64

// IOHIDDeviceInterface122 is cFPlugin object subclass which provides the primary interface to HID devices. This class is a subclass of IOHIDDeviceInterface121.
//
// See: https://developer.apple.com/documentation/iokit/iohiddeviceinterface122
// IOHIDDeviceInterface122 is opaque storage with the size and alignment C gives IOHIDDeviceInterface122:
// 176 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 176 into.
type IOHIDDeviceInterface122 [22]uint64

// IOHIDDeviceQueueInterface is the object you use to access a HID queue from user space, returned by version 1.5 of the IOHIDFamily.
//
// See: https://developer.apple.com/documentation/iokit/iohiddevicequeueinterface
// IOHIDDeviceQueueInterface is opaque storage with the size and alignment C gives IOHIDDeviceQueueInterface:
// 112 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 112 into.
type IOHIDDeviceQueueInterface [14]uint64

// See: https://developer.apple.com/documentation/iokit/iohiddeviceref
type IOHIDDeviceRef uintptr

// See: https://developer.apple.com/documentation/iokit/iohiddevicetimestampeddeviceinterface
// IOHIDDeviceTimeStampedDeviceInterface is opaque storage with the size and alignment C gives IOHIDDeviceTimeStampedDeviceInterface:
// 128 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 128 into.
type IOHIDDeviceTimeStampedDeviceInterface [16]uint64

// IOHIDDeviceTransactionInterface is the object you use to access a HID transaction from user space, returned by version 1.5 of the IOHIDFamily.
//
// See: https://developer.apple.com/documentation/iokit/iohiddevicetransactioninterface
// IOHIDDeviceTransactionInterface is opaque storage with the size and alignment C gives IOHIDDeviceTransactionInterface:
// 112 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 112 into.
type IOHIDDeviceTransactionInterface [14]uint64

// IOHIDElementCollectionType is describes different types of HID collections.
//
// See: https://developer.apple.com/documentation/iokit/iohidelementcollectiontype
type IOHIDElementCollectionType = uint32

// See: https://developer.apple.com/documentation/iokit/iohidelementcommitdirection
type IOHIDElementCommitDirection = uint32

// IOHIDElementCookie is abstract data type used as a unique identifier for an element.
//
// See: https://developer.apple.com/documentation/iokit/iohidelementcookie
type IOHIDElementCookie = uint32

// See: https://developer.apple.com/documentation/iokit/iohidelementflags
type IOHIDElementFlags = uint32

// See: https://developer.apple.com/documentation/iokit/iohidelementref
type IOHIDElementRef uintptr

// IOHIDElementType is describes different types of HID elements.
//
// See: https://developer.apple.com/documentation/iokit/iohidelementtype
type IOHIDElementType = uint32

// See: https://developer.apple.com/documentation/iokit/iohideventstruct
// IOHIDEventStruct is opaque storage with the size and alignment C gives IOHIDEventStruct:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type IOHIDEventStruct [4]uint64

// See: https://developer.apple.com/documentation/iokit/iohideventsystemclientref
type IOHIDEventSystemClientRef uintptr

// See: https://developer.apple.com/documentation/iokit/iohidkeyboardeventoptions
type IOHIDKeyboardEventOptions = uint32

// See: https://developer.apple.com/documentation/iokit/iohidkeyboardphysicallayouttype
type IOHIDKeyboardPhysicalLayoutType = uint32

// IOHIDManagerOptions is various options that can be supplied to IOHIDManager functions.
//
// See: https://developer.apple.com/documentation/iokit/iohidmanageroptions
type IOHIDManagerOptions = uint32

// IOHIDManagerRef is this is the type of a reference to the IOHIDManager.
//
// See: https://developer.apple.com/documentation/iokit/iohidmanagerref
type IOHIDManagerRef uintptr

// IOHIDOptionsType is options for opening a device via IOHIDLib.
//
// See: https://developer.apple.com/documentation/iokit/iohidoptionstype
type IOHIDOptionsType = uint32

// IOHIDOutputTransactionInterface is cFPlugin object subclass which privides interface for output transactions to HID devices. Created by a IOHIDDeviceInterface object.
//
// See: https://developer.apple.com/documentation/iokit/iohidoutputtransactioninterface
// IOHIDOutputTransactionInterface is opaque storage with the size and alignment C gives IOHIDOutputTransactionInterface:
// 152 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 152 into.
type IOHIDOutputTransactionInterface [19]uint64

// See: https://developer.apple.com/documentation/iokit/iohidpointereventoptions
type IOHIDPointerEventOptions = uint32

// IOHIDQueueInterface is cFPlugin object subclass which provides an interface for input queues from HID devices. Created by an IOHIDDeviceInterface object.
//
// See: https://developer.apple.com/documentation/iokit/iohidqueueinterface
// IOHIDQueueInterface is opaque storage with the size and alignment C gives IOHIDQueueInterface:
// 144 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 144 into.
type IOHIDQueueInterface [18]uint64

// IOHIDQueueOptionsType is options for creating a queue via IOHIDLib.
//
// See: https://developer.apple.com/documentation/iokit/iohidqueueoptionstype
type IOHIDQueueOptionsType = uint32

// See: https://developer.apple.com/documentation/iokit/iohidqueueref
type IOHIDQueueRef uintptr

// IOHIDReportType is describes different type of HID reports.
//
// See: https://developer.apple.com/documentation/iokit/iohidreporttype
type IOHIDReportType = uint32

// See: https://developer.apple.com/documentation/iokit/iohidrequesttype
type IOHIDRequestType = uint32

// See: https://developer.apple.com/documentation/iokit/iohidscrolleventoptions
type IOHIDScrollEventOptions = uint32

// See: https://developer.apple.com/documentation/iokit/iohidserviceclientref
type IOHIDServiceClientRef uintptr

// IOHIDStandardType is type to define what industrial standard the device is referencing.
//
// See: https://developer.apple.com/documentation/iokit/iohidstandardtype
type IOHIDStandardType = uint32

// IOHIDTransactionDirectionType is direction for an IOHIDDeviceTransactionInterface.
//
// See: https://developer.apple.com/documentation/iokit/iohidtransactiondirectiontype
type IOHIDTransactionDirectionType = uint32

// See: https://developer.apple.com/documentation/iokit/iohidtransactionoptions
type IOHIDTransactionOptions = uint32

// See: https://developer.apple.com/documentation/iokit/iohidtransactionref
type IOHIDTransactionRef uintptr

// See: https://developer.apple.com/documentation/iokit/iohiduserdevicegetreportblock
type IOHIDUserDeviceGetReportBlock = func(type_ IOHIDReportType, reportID uint32, report *uint8, reportLength *corefoundation.CFIndex) int32

// See: https://developer.apple.com/documentation/iokit/iohiduserdeviceoptions
type IOHIDUserDeviceOptions = uint

// See: https://developer.apple.com/documentation/iokit/iohiduserdeviceref
type IOHIDUserDeviceRef uintptr

// See: https://developer.apple.com/documentation/iokit/iohiduserdevicesetreportblock
type IOHIDUserDeviceSetReportBlock = func(type_ IOHIDReportType, reportID uint32, report *uint8, reportLength corefoundation.CFIndex) int32

// IOHIDValueOptions is describes options for gathering element values.
//
// See: https://developer.apple.com/documentation/iokit/iohidvalueoptions
type IOHIDValueOptions = uint32

// See: https://developer.apple.com/documentation/iokit/iohidvalueref
type IOHIDValueRef uintptr

// IOHIDValueScaleType is describes different types of scaling that can be performed on element values.
//
// See: https://developer.apple.com/documentation/iokit/iohidvaluescaletype
type IOHIDValueScaleType = uint32

// See: https://developer.apple.com/documentation/iokit/iohardwarecursordescriptor
// IOHardwareCursorDescriptor is opaque storage with the size and alignment C gives IOHardwareCursorDescriptor:
// 104 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 104 into.
type IOHardwareCursorDescriptor [13]uint64

// See: https://developer.apple.com/documentation/iokit/iohardwarecursorinfo
// IOHardwareCursorInfo is opaque storage with the size and alignment C gives IOHardwareCursorInfo:
// 56 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 56 into.
type IOHardwareCursorInfo [7]uint64

// See: https://developer.apple.com/documentation/iokit/ioi2cbuffer
// IOI2CBuffer is an unresolved C aggregate typedef.
type IOI2CBuffer unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/ioi2cbustiming
// IOI2CBusTiming is opaque storage with the size and alignment C gives IOI2CBusTiming:
// 80 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 80 into.
type IOI2CBusTiming [20]uint32

// See: https://developer.apple.com/documentation/iokit/ioi2cconnectref
type IOI2CConnectRef uintptr

// See: https://developer.apple.com/documentation/iokit/ioi2crequest
// IOI2CRequest is opaque storage with the size and alignment C gives IOI2CRequest:
// 124 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 124 into.
type IOI2CRequest [31]uint32

// See: https://developer.apple.com/documentation/iokit/ioindex
type IOIndex = int32

// See: https://developer.apple.com/documentation/iokit/ioitemcount
type IOItemCount = uint32

// See: https://developer.apple.com/documentation/iokit/iologicaladdress
type IOLogicalAddress = kernel.IOVirtualAddress

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
// IONVMeSMARTInterface is opaque storage with the size and alignment C gives IONVMeSMARTInterface:
// 256 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 256 into.
type IONVMeSMARTInterface [32]uint64

// See: https://developer.apple.com/documentation/iokit/ionetworkstats
// IONetworkStats is opaque storage with the size and alignment C gives IONetworkStats:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type IONetworkStats [5]uint32

// See: https://developer.apple.com/documentation/iokit/ionotificationportref
type IONotificationPortRef uintptr

// See: https://developer.apple.com/documentation/iokit/iooptionbits
type IOOptionBits = uint32

// See: https://developer.apple.com/documentation/iokit/iooutputqueuestats
// IOOutputQueueStats is opaque storage with the size and alignment C gives IOOutputQueueStats:
// 44 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 44 into.
type IOOutputQueueStats [11]uint32

// IOPMAssertionID is type for AssertionID arguments to [IOPMAssertionCreateWithProperties] and [IOPMAssertionRelease].
//
// See: https://developer.apple.com/documentation/iokit/iopmassertionid
type IOPMAssertionID = uint32

// IOPMAssertionLevel is type for AssertionLevel argument to IOPMAssertionCreate.
//
// See: https://developer.apple.com/documentation/iokit/iopmassertionlevel
type IOPMAssertionLevel = uint32

// See: https://developer.apple.com/documentation/iokit/iopmcalendarstruct
// IOPMCalendarStruct is opaque storage with the size and alignment C gives IOPMCalendarStruct:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type IOPMCalendarStruct [3]uint32

// IOPMPowerFlags is bits are used in defining capabilityFlags, inputPowerRequirements, and outputPowerCharacter in the IOPMPowerState structure.
//
// See: https://developer.apple.com/documentation/iokit/iopmpowerflags
type IOPMPowerFlags = uint

// See: https://developer.apple.com/documentation/iokit/iopmuseractivetype
type IOPMUserActiveType = uint32

// IOPSLowBatteryWarningLevel is the battery can provide no more than 10 minutes of runtime.
//
// See: https://developer.apple.com/documentation/iokit/iopslowbatterywarninglevel
type IOPSLowBatteryWarningLevel = uint32

// See: https://developer.apple.com/documentation/iokit/iophysicaladdress
type IOPhysicalAddress = uint64

// See: https://developer.apple.com/documentation/iokit/iophysicaladdress32
type IOPhysicalAddress32 = uint32

// See: https://developer.apple.com/documentation/iokit/iophysicaladdress64
type IOPhysicalAddress64 = uint64

// See: https://developer.apple.com/documentation/iokit/iophysicallength
type IOPhysicalLength = uint64

// See: https://developer.apple.com/documentation/iokit/iophysicallength32
type IOPhysicalLength32 = uint32

// See: https://developer.apple.com/documentation/iokit/iophysicallength64
type IOPhysicalLength64 = uint64

// See: https://developer.apple.com/documentation/iokit/iopixelaperture
type IOPixelAperture = int32

// See: https://developer.apple.com/documentation/iokit/iopixelencoding
type IOPixelEncoding = int8

// See: https://developer.apple.com/documentation/iokit/iopixelinformation
// IOPixelInformation is opaque storage with the size and alignment C gives IOPixelInformation:
// 172 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 172 into.
type IOPixelInformation [43]uint32

// See: https://developer.apple.com/documentation/iokit/iopowerstatechangenotification
// IOPowerStateChangeNotification is opaque storage with the size and alignment C gives IOPowerStateChangeNotification:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type IOPowerStateChangeNotification [4]uint64

// See: https://developer.apple.com/documentation/iokit/ioreturn
type IOReturn = int32

// See: https://developer.apple.com/documentation/iokit/ioselect
type IOSelect = uint32

// IOServiceInterestCallback is callback function to be notified of changes in state of an IOService.
//
// See: https://developer.apple.com/documentation/iokit/ioserviceinterestcallback
type IOServiceInterestCallback = func(refcon unsafe.Pointer, service uint32, messageType uint32, messageArgument unsafe.Pointer)

// IOServiceMatchingCallback is callback function to be notified of IOService publication.
//
// See: https://developer.apple.com/documentation/iokit/ioservicematchingcallback
type IOServiceMatchingCallback = func(refcon unsafe.Pointer, iterator uint32)

// See: https://developer.apple.com/documentation/iokit/iostorageunmapoptions
type IOStorageUnmapOptions = uint32

// See: https://developer.apple.com/documentation/iokit/iostreambufferid
type IOStreamBufferID = uint32

// See: https://developer.apple.com/documentation/iokit/iostreaminterface
// IOStreamInterface is opaque storage with the size and alignment C gives IOStreamInterface:
// 248 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 248 into.
type IOStreamInterface [31]uint64

// See: https://developer.apple.com/documentation/iokit/iostreammode
type IOStreamMode = uint32

// See: https://developer.apple.com/documentation/iokit/iostreamref
type IOStreamRef uintptr

// IOSystemLoadAdvisoryLevel is return type for IOGetSystemLoadAdvisory.
//
// See: https://developer.apple.com/documentation/iokit/iosystemloadadvisorylevel
type IOSystemLoadAdvisoryLevel = int32

// See: https://developer.apple.com/documentation/iokit/iotiminginformation
// IOTimingInformation is opaque storage with the size and alignment C gives IOTimingInformation:
// 168 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 168 into.
type IOTimingInformation [21]uint64

// IOUPSPlugInInterface is represents and provides management functions for a UPS device.
//
// See: https://developer.apple.com/documentation/iokit/ioupsplugininterface
// IOUPSPlugInInterface is opaque storage with the size and alignment C gives IOUPSPlugInInterface:
// 72 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 72 into.
type IOUPSPlugInInterface [9]uint64

// See: https://developer.apple.com/documentation/iokit/ioupsplugininterface_v140
// IOUPSPlugInInterface_v140 is opaque storage with the size and alignment C gives IOUPSPlugInInterface_v140:
// 80 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 80 into.
type IOUPSPlugInInterface_v140 [10]uint64

// See: https://developer.apple.com/documentation/iokit/iousb20hubdescriptor
// IOUSB20HubDescriptor is opaque storage with the size and alignment C gives IOUSB20HubDescriptor:
// 11 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 11 into.
type IOUSB20HubDescriptor [11]byte

// See: https://developer.apple.com/documentation/iokit/iousbbosdescriptor
// IOUSBBOSDescriptor is opaque storage with the size and alignment C gives IOUSBBOSDescriptor:
// 5 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 5 into.
type IOUSBBOSDescriptor [5]byte

// See: https://developer.apple.com/documentation/iokit/iousbbosdescriptorptr
type IOUSBBOSDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbbulkpipereq
// IOUSBBulkPipeReq is opaque storage with the size and alignment C gives IOUSBBulkPipeReq:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type IOUSBBulkPipeReq [4]uint64

// See: https://developer.apple.com/documentation/iokit/iousbcompletion
// IOUSBCompletion is opaque storage with the size and alignment C gives IOUSBCompletion:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IOUSBCompletion [3]uint64

// See: https://developer.apple.com/documentation/iokit/iousbcompletionwithtimestamp
// IOUSBCompletionWithTimeStamp is opaque storage with the size and alignment C gives IOUSBCompletionWithTimeStamp:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IOUSBCompletionWithTimeStamp [3]uint64

// See: https://developer.apple.com/documentation/iokit/iousbconfigurationdescheader
// IOUSBConfigurationDescHeader is opaque storage with the size and alignment C gives IOUSBConfigurationDescHeader:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type IOUSBConfigurationDescHeader [4]byte

// See: https://developer.apple.com/documentation/iokit/iousbconfigurationdescheaderptr
type IOUSBConfigurationDescHeaderPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbconfigurationdescriptor
// IOUSBConfigurationDescriptor is opaque storage with the size and alignment C gives IOUSBConfigurationDescriptor:
// 9 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 9 into.
type IOUSBConfigurationDescriptor [9]byte

// See: https://developer.apple.com/documentation/iokit/iousbconfigurationdescriptorptr
type IOUSBConfigurationDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdfudescriptor
// IOUSBDFUDescriptor is opaque storage with the size and alignment C gives IOUSBDFUDescriptor:
// 7 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 7 into.
type IOUSBDFUDescriptor [7]byte

// See: https://developer.apple.com/documentation/iokit/iousbdfudescriptorptr
type IOUSBDFUDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdescriptor
// IOUSBDescriptor is opaque storage with the size and alignment C gives IOUSBDescriptor:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type IOUSBDescriptor [2]byte

// See: https://developer.apple.com/documentation/iokit/iousbdescriptorheader
// IOUSBDescriptorHeader is opaque storage with the size and alignment C gives IOUSBDescriptorHeader:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type IOUSBDescriptorHeader [2]byte

// See: https://developer.apple.com/documentation/iokit/iousbdescriptorheaderptr
type IOUSBDescriptorHeaderPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevreqool
// IOUSBDevReqOOL is opaque storage with the size and alignment C gives IOUSBDevReqOOL:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IOUSBDevReqOOL [3]uint64

// See: https://developer.apple.com/documentation/iokit/iousbdevreqoolto
// IOUSBDevReqOOLTO is opaque storage with the size and alignment C gives IOUSBDevReqOOLTO:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type IOUSBDevReqOOLTO [4]uint64

// See: https://developer.apple.com/documentation/iokit/iousbdevrequest
// IOUSBDevRequest is opaque storage with the size and alignment C gives IOUSBDevRequest:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IOUSBDevRequest [3]uint64

// See: https://developer.apple.com/documentation/iokit/iousbdevrequestto
// IOUSBDevRequestTO is opaque storage with the size and alignment C gives IOUSBDevRequestTO:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type IOUSBDevRequestTO [4]uint64

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitybillboard
// IOUSBDeviceCapabilityBillboard is opaque storage with the size and alignment C gives IOUSBDeviceCapabilityBillboard:
// 44 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 44 into.
type IOUSBDeviceCapabilityBillboard [44]byte

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitybillboardaltconfig
// IOUSBDeviceCapabilityBillboardAltConfig is opaque storage with the size and alignment C gives IOUSBDeviceCapabilityBillboardAltConfig:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type IOUSBDeviceCapabilityBillboardAltConfig [4]byte

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitybillboardaltconfigcompatibility
// IOUSBDeviceCapabilityBillboardAltConfigCompatibility is opaque storage with the size and alignment C gives IOUSBDeviceCapabilityBillboardAltConfigCompatibility:
// 7 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 7 into.
type IOUSBDeviceCapabilityBillboardAltConfigCompatibility [7]byte

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitybillboardaltconfigptr
type IOUSBDeviceCapabilityBillboardAltConfigPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitybillboardaltmode
type IOUSBDeviceCapabilityBillboardAltMode = uint

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitybillboardaltmodeptr
type IOUSBDeviceCapabilityBillboardAltModePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitybillboardptr
type IOUSBDeviceCapabilityBillboardPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitycontainerid
// IOUSBDeviceCapabilityContainerID is opaque storage with the size and alignment C gives IOUSBDeviceCapabilityContainerID:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type IOUSBDeviceCapabilityContainerID [20]byte

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitycontaineridptr
type IOUSBDeviceCapabilityContainerIDPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitydescriptorheader
// IOUSBDeviceCapabilityDescriptorHeader is opaque storage with the size and alignment C gives IOUSBDeviceCapabilityDescriptorHeader:
// 3 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 3 into.
type IOUSBDeviceCapabilityDescriptorHeader [3]byte

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitydescriptorheaderptr
type IOUSBDeviceCapabilityDescriptorHeaderPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitysuperspeedplususb
// IOUSBDeviceCapabilitySuperSpeedPlusUSB is opaque storage with the size and alignment C gives IOUSBDeviceCapabilitySuperSpeedPlusUSB:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type IOUSBDeviceCapabilitySuperSpeedPlusUSB [12]byte

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitysuperspeedplususbptr
type IOUSBDeviceCapabilitySuperSpeedPlusUSBPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitysuperspeedusb
// IOUSBDeviceCapabilitySuperSpeedUSB is opaque storage with the size and alignment C gives IOUSBDeviceCapabilitySuperSpeedUSB:
// 10 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 10 into.
type IOUSBDeviceCapabilitySuperSpeedUSB [10]byte

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilitysuperspeedusbptr
type IOUSBDeviceCapabilitySuperSpeedUSBPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilityusb2extension
// IOUSBDeviceCapabilityUSB2Extension is opaque storage with the size and alignment C gives IOUSBDeviceCapabilityUSB2Extension:
// 7 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 7 into.
type IOUSBDeviceCapabilityUSB2Extension [7]byte

// See: https://developer.apple.com/documentation/iokit/iousbdevicecapabilityusb2extensionptr
type IOUSBDeviceCapabilityUSB2ExtensionPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicedescriptor
// IOUSBDeviceDescriptor is opaque storage with the size and alignment C gives IOUSBDeviceDescriptor:
// 18 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 18 into.
type IOUSBDeviceDescriptor [18]byte

// See: https://developer.apple.com/documentation/iokit/iousbdevicedescriptorptr
type IOUSBDeviceDescriptorPtr = unsafe.Pointer

// IOUSBDeviceInterface is the object you use to access USB devices from user space, returned by all versions of the IOUSBFamily currently shipping.
//
// See: https://developer.apple.com/documentation/iokit/iousbdeviceinterface
// IOUSBDeviceInterface is opaque storage with the size and alignment C gives IOUSBDeviceInterface:
// 408 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 408 into.
type IOUSBDeviceInterface [51]uint64

// See: https://developer.apple.com/documentation/iokit/iousbdeviceinterface100
// IOUSBDeviceInterface100 is opaque storage with the size and alignment C gives IOUSBDeviceInterface100:
// 232 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 232 into.
type IOUSBDeviceInterface100 [29]uint64

// IOUSBDeviceInterface182 is the object you use to access USB devices from user space, returned by the IOUSBFamily version 1.8.2 and above.
//
// See: https://developer.apple.com/documentation/iokit/iousbdeviceinterface182
// IOUSBDeviceInterface182 is opaque storage with the size and alignment C gives IOUSBDeviceInterface182:
// 296 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 296 into.
type IOUSBDeviceInterface182 [37]uint64

// IOUSBDeviceInterface187 is the object you use to access USB devices from user space, returned by the IOUSBFamily version 10.8.7 and above.
//
// See: https://developer.apple.com/documentation/iokit/iousbdeviceinterface187
// IOUSBDeviceInterface187 is opaque storage with the size and alignment C gives IOUSBDeviceInterface187:
// 304 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 304 into.
type IOUSBDeviceInterface187 [38]uint64

// IOUSBDeviceInterface197 is the object you use to access USB devices from user space, returned by the IOUSBFamily version 1.9.7 and above.
//
// See: https://developer.apple.com/documentation/iokit/iousbdeviceinterface197
// IOUSBDeviceInterface197 is opaque storage with the size and alignment C gives IOUSBDeviceInterface197:
// 320 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 320 into.
type IOUSBDeviceInterface197 [40]uint64

// IOUSBDeviceInterface245 is the object you use to access USB devices from user space, returned by the IOUSBFamily version 2.4.5 and above.
//
// See: https://developer.apple.com/documentation/iokit/iousbdeviceinterface245
// IOUSBDeviceInterface245 is opaque storage with the size and alignment C gives IOUSBDeviceInterface245:
// 320 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 320 into.
type IOUSBDeviceInterface245 [40]uint64

// IOUSBDeviceInterface300 is the object you use to access USB devices from user space, returned by the IOUSBFamily version 3.0.0 and above.
//
// See: https://developer.apple.com/documentation/iokit/iousbdeviceinterface300
// IOUSBDeviceInterface300 is opaque storage with the size and alignment C gives IOUSBDeviceInterface300:
// 328 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 328 into.
type IOUSBDeviceInterface300 [41]uint64

// IOUSBDeviceInterface320 is the object you use to access USB devices from user space, returned by the IOUSBFamily version 3.2.0 and above.
//
// See: https://developer.apple.com/documentation/iokit/iousbdeviceinterface320
// IOUSBDeviceInterface320 is opaque storage with the size and alignment C gives IOUSBDeviceInterface320:
// 360 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 360 into.
type IOUSBDeviceInterface320 [45]uint64

// See: https://developer.apple.com/documentation/iokit/iousbdeviceinterface400
// IOUSBDeviceInterface400 is opaque storage with the size and alignment C gives IOUSBDeviceInterface400:
// 360 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 360 into.
type IOUSBDeviceInterface400 [45]uint64

// IOUSBDeviceInterface500 is the object you use to access USB devices from user space, returned by the IOUSBFamily version 3.2.0 and above.
//
// See: https://developer.apple.com/documentation/iokit/iousbdeviceinterface500
// IOUSBDeviceInterface500 is opaque storage with the size and alignment C gives IOUSBDeviceInterface500:
// 368 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 368 into.
type IOUSBDeviceInterface500 [46]uint64

// See: https://developer.apple.com/documentation/iokit/iousbdeviceinterface650
// IOUSBDeviceInterface650 is opaque storage with the size and alignment C gives IOUSBDeviceInterface650:
// 400 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 400 into.
type IOUSBDeviceInterface650 [50]uint64

// See: https://developer.apple.com/documentation/iokit/iousbdeviceinterface942
// IOUSBDeviceInterface942 is opaque storage with the size and alignment C gives IOUSBDeviceInterface942:
// 408 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 408 into.
type IOUSBDeviceInterface942 [51]uint64

// See: https://developer.apple.com/documentation/iokit/iousbdevicequalifierdescriptor
// IOUSBDeviceQualifierDescriptor is opaque storage with the size and alignment C gives IOUSBDeviceQualifierDescriptor:
// 10 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 10 into.
type IOUSBDeviceQualifierDescriptor [10]byte

// See: https://developer.apple.com/documentation/iokit/iousbdevicequalifierdescriptorptr
type IOUSBDeviceQualifierDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicerequest
// IOUSBDeviceRequest is opaque storage with the size and alignment C gives IOUSBDeviceRequest:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOUSBDeviceRequest [8]byte

// See: https://developer.apple.com/documentation/iokit/iousbdevicerequestptr
type IOUSBDeviceRequestPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbdevicerequestsetseldata
// IOUSBDeviceRequestSetSELData is opaque storage with the size and alignment C gives IOUSBDeviceRequestSetSELData:
// 6 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 6 into.
type IOUSBDeviceRequestSetSELData [6]byte

// See: https://developer.apple.com/documentation/iokit/iousbendpointdescriptor
// IOUSBEndpointDescriptor is opaque storage with the size and alignment C gives IOUSBEndpointDescriptor:
// 7 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 7 into.
type IOUSBEndpointDescriptor [7]byte

// See: https://developer.apple.com/documentation/iokit/iousbendpointdescriptorptr
type IOUSBEndpointDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbendpointproperties
// IOUSBEndpointProperties is opaque storage with the size and alignment C gives IOUSBEndpointProperties:
// 15 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 15 into.
type IOUSBEndpointProperties [15]byte

// See: https://developer.apple.com/documentation/iokit/iousbendpointpropertiesptr
type IOUSBEndpointPropertiesPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbfindendpointrequest
// IOUSBFindEndpointRequest is opaque storage with the size and alignment C gives IOUSBFindEndpointRequest:
// 6 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 6 into.
type IOUSBFindEndpointRequest [3]uint16

// See: https://developer.apple.com/documentation/iokit/iousbfindinterfacerequest
// IOUSBFindInterfaceRequest is opaque storage with the size and alignment C gives IOUSBFindInterfaceRequest:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOUSBFindInterfaceRequest [4]uint16

// See: https://developer.apple.com/documentation/iokit/iousbgetframestruct
// IOUSBGetFrameStruct is opaque storage with the size and alignment C gives IOUSBGetFrameStruct:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type IOUSBGetFrameStruct [2]uint64

// See: https://developer.apple.com/documentation/iokit/iousbhiddataptr
type IOUSBHIDDataPtr = *kernel.IOUSBHIDData

// See: https://developer.apple.com/documentation/iokit/iousbhiddescriptor
// IOUSBHIDDescriptor is opaque storage with the size and alignment C gives IOUSBHIDDescriptor:
// 9 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 9 into.
type IOUSBHIDDescriptor [9]byte

// See: https://developer.apple.com/documentation/iokit/iousbhiddescriptorptr
type IOUSBHIDDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbhidreportdesc
// IOUSBHIDReportDesc is opaque storage with the size and alignment C gives IOUSBHIDReportDesc:
// 3 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 3 into.
type IOUSBHIDReportDesc [3]byte

// See: https://developer.apple.com/documentation/iokit/iousbhidreportdescptr
type IOUSBHIDReportDescPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceassociationdescriptor
// IOUSBInterfaceAssociationDescriptor is opaque storage with the size and alignment C gives IOUSBInterfaceAssociationDescriptor:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOUSBInterfaceAssociationDescriptor [8]byte

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceassociationdescriptorptr
type IOUSBInterfaceAssociationDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbinterfacedescriptor
// IOUSBInterfaceDescriptor is opaque storage with the size and alignment C gives IOUSBInterfaceDescriptor:
// 9 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 9 into.
type IOUSBInterfaceDescriptor [9]byte

// See: https://developer.apple.com/documentation/iokit/iousbinterfacedescriptorptr
type IOUSBInterfaceDescriptorPtr = unsafe.Pointer

// IOUSBInterfaceInterface is the object you use to access a USB device interface from user space, returned by all versions of the IOUSBFamily currently shipping.
//
// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface
// IOUSBInterfaceInterface is opaque storage with the size and alignment C gives IOUSBInterfaceInterface:
// 616 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 616 into.
type IOUSBInterfaceInterface [77]uint64

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface100
// IOUSBInterfaceInterface100 is opaque storage with the size and alignment C gives IOUSBInterfaceInterface100:
// 296 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 296 into.
type IOUSBInterfaceInterface100 [37]uint64

// IOUSBInterfaceInterface182 is the object you use to access a USB device interface from user space, returned by the IOUSBFamily version 1.8.2 and above.
//
// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface182
// IOUSBInterfaceInterface182 is opaque storage with the size and alignment C gives IOUSBInterfaceInterface182:
// 352 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 352 into.
type IOUSBInterfaceInterface182 [44]uint64

// IOUSBInterfaceInterface183 is the object you use to access a USB device interface from user space, returned by the IOUSBFamily version 1.8.3 and above.
//
// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface183
// IOUSBInterfaceInterface183 is opaque storage with the size and alignment C gives IOUSBInterfaceInterface183:
// 360 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 360 into.
type IOUSBInterfaceInterface183 [45]uint64

// IOUSBInterfaceInterface190 is the object you use to access a USB device interface from user space, returned by the IOUSBFamily version 1.9 and above.
//
// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface190
// IOUSBInterfaceInterface190 is opaque storage with the size and alignment C gives IOUSBInterfaceInterface190:
// 392 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 392 into.
type IOUSBInterfaceInterface190 [49]uint64

// IOUSBInterfaceInterface192 is the object you use to access a USB device interface from user space, returned by the IOUSBFamily version 1.9.2 and above.
//
// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface192
// IOUSBInterfaceInterface192 is opaque storage with the size and alignment C gives IOUSBInterfaceInterface192:
// 424 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 424 into.
type IOUSBInterfaceInterface192 [53]uint64

// IOUSBInterfaceInterface197 is the object you use to access a USB device interface from user space, returned by the IOUSBFamily version 1.9.7 and above.
//
// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface197
// IOUSBInterfaceInterface197 is opaque storage with the size and alignment C gives IOUSBInterfaceInterface197:
// 448 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 448 into.
type IOUSBInterfaceInterface197 [56]uint64

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface220
// IOUSBInterfaceInterface220 is opaque storage with the size and alignment C gives IOUSBInterfaceInterface220:
// 464 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 464 into.
type IOUSBInterfaceInterface220 [58]uint64

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface245
// IOUSBInterfaceInterface245 is opaque storage with the size and alignment C gives IOUSBInterfaceInterface245:
// 464 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 464 into.
type IOUSBInterfaceInterface245 [58]uint64

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface300
// IOUSBInterfaceInterface300 is opaque storage with the size and alignment C gives IOUSBInterfaceInterface300:
// 472 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 472 into.
type IOUSBInterfaceInterface300 [59]uint64

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface398
// IOUSBInterfaceInterface398 is opaque storage with the size and alignment C gives IOUSBInterfaceInterface398:
// 472 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 472 into.
type IOUSBInterfaceInterface398 [59]uint64

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface400
// IOUSBInterfaceInterface400 is opaque storage with the size and alignment C gives IOUSBInterfaceInterface400:
// 472 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 472 into.
type IOUSBInterfaceInterface400 [59]uint64

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface500
// IOUSBInterfaceInterface500 is opaque storage with the size and alignment C gives IOUSBInterfaceInterface500:
// 480 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 480 into.
type IOUSBInterfaceInterface500 [60]uint64

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface550
// IOUSBInterfaceInterface550 is opaque storage with the size and alignment C gives IOUSBInterfaceInterface550:
// 560 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 560 into.
type IOUSBInterfaceInterface550 [70]uint64

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface650
// IOUSBInterfaceInterface650 is opaque storage with the size and alignment C gives IOUSBInterfaceInterface650:
// 584 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 584 into.
type IOUSBInterfaceInterface650 [73]uint64

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface700
// IOUSBInterfaceInterface700 is opaque storage with the size and alignment C gives IOUSBInterfaceInterface700:
// 592 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 592 into.
type IOUSBInterfaceInterface700 [74]uint64

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface800
// IOUSBInterfaceInterface800 is opaque storage with the size and alignment C gives IOUSBInterfaceInterface800:
// 608 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 608 into.
type IOUSBInterfaceInterface800 [76]uint64

// See: https://developer.apple.com/documentation/iokit/iousbinterfaceinterface942
// IOUSBInterfaceInterface942 is opaque storage with the size and alignment C gives IOUSBInterfaceInterface942:
// 616 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 616 into.
type IOUSBInterfaceInterface942 [77]uint64

// See: https://developer.apple.com/documentation/iokit/iousbisoccompletion
// IOUSBIsocCompletion is opaque storage with the size and alignment C gives IOUSBIsocCompletion:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IOUSBIsocCompletion [3]uint64

// See: https://developer.apple.com/documentation/iokit/iousbisocframe
// IOUSBIsocFrame is opaque storage with the size and alignment C gives IOUSBIsocFrame:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOUSBIsocFrame [2]uint32

// See: https://developer.apple.com/documentation/iokit/iousbisocstruct
// IOUSBIsocStruct is opaque storage with the size and alignment C gives IOUSBIsocStruct:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type IOUSBIsocStruct [6]uint64

// See: https://developer.apple.com/documentation/iokit/iousbkeyboarddata
// IOUSBKeyboardData is opaque storage with the size and alignment C gives IOUSBKeyboardData:
// 66 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 66 into.
type IOUSBKeyboardData [33]uint16

// See: https://developer.apple.com/documentation/iokit/iousbkeyboarddataptr
type IOUSBKeyboardDataPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousblowlatencyisoccompletion
// IOUSBLowLatencyIsocCompletion is opaque storage with the size and alignment C gives IOUSBLowLatencyIsocCompletion:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IOUSBLowLatencyIsocCompletion [3]uint64

// See: https://developer.apple.com/documentation/iokit/iousblowlatencyisocframe
// IOUSBLowLatencyIsocFrame is opaque storage with the size and alignment C gives IOUSBLowLatencyIsocFrame:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type IOUSBLowLatencyIsocFrame [4]uint32

// See: https://developer.apple.com/documentation/iokit/iousblowlatencyisocstruct
// IOUSBLowLatencyIsocStruct is opaque storage with the size and alignment C gives IOUSBLowLatencyIsocStruct:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type IOUSBLowLatencyIsocStruct [5]uint64

// See: https://developer.apple.com/documentation/iokit/iousbmatch
// IOUSBMatch is opaque storage with the size and alignment C gives IOUSBMatch:
// 10 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 10 into.
type IOUSBMatch [5]uint16

// See: https://developer.apple.com/documentation/iokit/iousbmousedata
// IOUSBMouseData is opaque storage with the size and alignment C gives IOUSBMouseData:
// 6 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 6 into.
type IOUSBMouseData [3]uint16

// See: https://developer.apple.com/documentation/iokit/iousbmousedataptr
type IOUSBMouseDataPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbplatformcapabilitydescriptor
// IOUSBPlatformCapabilityDescriptor is opaque storage with the size and alignment C gives IOUSBPlatformCapabilityDescriptor:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type IOUSBPlatformCapabilityDescriptor [20]byte

// See: https://developer.apple.com/documentation/iokit/iousbplatformcapabilitydescriptorptr
type IOUSBPlatformCapabilityDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbstringdescriptor
// IOUSBStringDescriptor is opaque storage with the size and alignment C gives IOUSBStringDescriptor:
// 3 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 3 into.
type IOUSBStringDescriptor [3]byte

// See: https://developer.apple.com/documentation/iokit/iousbstringdescriptorptr
type IOUSBStringDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbsuperspeedendpointcompaniondescriptor
// IOUSBSuperSpeedEndpointCompanionDescriptor is opaque storage with the size and alignment C gives IOUSBSuperSpeedEndpointCompanionDescriptor:
// 6 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 6 into.
type IOUSBSuperSpeedEndpointCompanionDescriptor [6]byte

// See: https://developer.apple.com/documentation/iokit/iousbsuperspeedendpointcompaniondescriptorptr
type IOUSBSuperSpeedEndpointCompanionDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iousbsuperspeedhubdescriptor
// IOUSBSuperSpeedHubDescriptor is opaque storage with the size and alignment C gives IOUSBSuperSpeedHubDescriptor:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type IOUSBSuperSpeedHubDescriptor [12]byte

// See: https://developer.apple.com/documentation/iokit/iousbsuperspeedplusisochronousendpointcompaniondescriptor
// IOUSBSuperSpeedPlusIsochronousEndpointCompanionDescriptor is opaque storage with the size and alignment C gives IOUSBSuperSpeedPlusIsochronousEndpointCompanionDescriptor:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type IOUSBSuperSpeedPlusIsochronousEndpointCompanionDescriptor [8]byte

// See: https://developer.apple.com/documentation/iokit/iousbsuperspeedplusisochronousendpointcompaniondescriptorptr
type IOUSBSuperSpeedPlusIsochronousEndpointCompanionDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/ioversion
type IOVersion = uint32

// See: https://developer.apple.com/documentation/iokit/iovideodeviceinterface
type IOVideoDeviceInterface = IOVideoDeviceInterface_v1_t

// IOVideoDeviceInterface_v1_t is forward declaration of IOVideoDeviceInterface_v1_t.
//
// See: https://developer.apple.com/documentation/iokit/iovideodeviceinterface_v1_t
// IOVideoDeviceInterface_v1_t is opaque storage with the size and alignment C gives IOVideoDeviceInterface_v1_t:
// 128 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 128 into.
type IOVideoDeviceInterface_v1_t [16]uint64

// See: https://developer.apple.com/documentation/iokit/iovideodevicenotification
// IOVideoDeviceNotification is opaque storage with the size and alignment C gives IOVideoDeviceNotification:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type IOVideoDeviceNotification [4]uint64

// See: https://developer.apple.com/documentation/iokit/iovideodevicenotificationmessage
// IOVideoDeviceNotificationMessage is opaque storage with the size and alignment C gives IOVideoDeviceNotificationMessage:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type IOVideoDeviceNotificationMessage [8]uint64

// See: https://developer.apple.com/documentation/iokit/iovideodeviceref
type IOVideoDeviceRef = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/iovideostreamdescription
// IOVideoStreamDescription is opaque storage with the size and alignment C gives IOVideoStreamDescription:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type IOVideoStreamDescription [6]uint32

// See: https://developer.apple.com/documentation/iokit/iovirtualaddress
type IOVirtualAddress = uint64

// See: https://developer.apple.com/documentation/iokit/longlbamodeparameterblockdescriptor
// LongLBAModeParameterBlockDescriptor is opaque storage with the size and alignment C gives LongLBAModeParameterBlockDescriptor:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type LongLBAModeParameterBlockDescriptor [16]byte

// See: https://developer.apple.com/documentation/iokit/lowlatencyuserbufferinfo
// LowLatencyUserBufferInfo is opaque storage with the size and alignment C gives LowLatencyUserBufferInfo:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type LowLatencyUserBufferInfo [5]uint64

// See: https://developer.apple.com/documentation/iokit/lowlatencyuserbufferinfov2
// LowLatencyUserBufferInfoV2 is opaque storage with the size and alignment C gives LowLatencyUserBufferInfoV2:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type LowLatencyUserBufferInfoV2 [6]uint64

// See: https://developer.apple.com/documentation/iokit/lowlatencyuserbufferinfov3
// LowLatencyUserBufferInfoV3 is opaque storage with the size and alignment C gives LowLatencyUserBufferInfoV3:
// 56 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 56 into.
type LowLatencyUserBufferInfoV3 [7]uint64

// MMCDeviceInterface is basic interface for an MMC-2 Compliant Device.
//
// See: https://developer.apple.com/documentation/iokit/mmcdeviceinterface
// MMCDeviceInterface is opaque storage with the size and alignment C gives MMCDeviceInterface:
// 200 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 200 into.
type MMCDeviceInterface [25]uint64

// See: https://developer.apple.com/documentation/iokit/modepageformatheader
// ModePageFormatHeader is opaque storage with the size and alignment C gives ModePageFormatHeader:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type ModePageFormatHeader [2]byte

// See: https://developer.apple.com/documentation/iokit/modeparameterblockdescriptor
// ModeParameterBlockDescriptor is opaque storage with the size and alignment C gives ModeParameterBlockDescriptor:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type ModeParameterBlockDescriptor [8]byte

// See: https://developer.apple.com/documentation/iokit/nvmeidentifycontrollerstruct
// NVMeIdentifyControllerStruct is opaque storage with the size and alignment C gives NVMeIdentifyControllerStruct:
// 4096 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4096 into.
type NVMeIdentifyControllerStruct [4096]byte

// See: https://developer.apple.com/documentation/iokit/nvmeidentifynamespacestruct
// NVMeIdentifyNamespaceStruct is opaque storage with the size and alignment C gives NVMeIdentifyNamespaceStruct:
// 4096 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4096 into.
type NVMeIdentifyNamespaceStruct [4096]byte

// See: https://developer.apple.com/documentation/iokit/nvmelbaformatdatastruct
// NVMeLBAFormatDataStruct is opaque storage with the size and alignment C gives NVMeLBAFormatDataStruct:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type NVMeLBAFormatDataStruct [4]byte

// See: https://developer.apple.com/documentation/iokit/nvmepowerstatedescriptor
// NVMePowerStateDescriptor is opaque storage with the size and alignment C gives NVMePowerStateDescriptor:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type NVMePowerStateDescriptor [32]byte

// See: https://developer.apple.com/documentation/iokit/nvmesmartdata
// NVMeSMARTData is opaque storage with the size and alignment C gives NVMeSMARTData:
// 512 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 512 into.
type NVMeSMARTData [512]byte

// See: https://developer.apple.com/documentation/iokit/nxcoord
type NXCoord = float32

// See: https://developer.apple.com/documentation/iokit/nxeqelement
// NXEQElement is opaque storage with the size and alignment C gives NXEQElement:
// 96 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 96 into.
type NXEQElement [24]uint32

// See: https://developer.apple.com/documentation/iokit/nxevent
// NXEvent is opaque storage with the size and alignment C gives NXEvent:
// 88 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 88 into.
type NXEvent [22]uint32

// See: https://developer.apple.com/documentation/iokit/nxeventext
// NXEventExt is opaque storage with the size and alignment C gives NXEventExt:
// 124 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 124 into.
type NXEventExt [31]uint32

// See: https://developer.apple.com/documentation/iokit/nxeventextension
// NXEventExtension is opaque storage with the size and alignment C gives NXEventExtension:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type NXEventExtension [9]uint32

// See: https://developer.apple.com/documentation/iokit/nxeventhandle
type NXEventHandle = uint32

// See: https://developer.apple.com/documentation/iokit/nxeventptr
type NXEventPtr = uintptr

// See: https://developer.apple.com/documentation/iokit/nxeventsystemdevice
// NXEventSystemDevice is opaque storage with the size and alignment C gives NXEventSystemDevice:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type NXEventSystemDevice [4]uint32

// See: https://developer.apple.com/documentation/iokit/nxeventsystemdevicelist
// NXEventSystemDeviceList is opaque storage with the size and alignment C gives NXEventSystemDeviceList:
// 256 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 256 into.
type NXEventSystemDeviceList [64]uint32

// See: https://developer.apple.com/documentation/iokit/nxeventsysteminfodata
type NXEventSystemInfoData = int32

// See: https://developer.apple.com/documentation/iokit/nxeventsysteminfotype
type NXEventSystemInfoType = *int32

// See: https://developer.apple.com/documentation/iokit/nxkeymapping
// NXKeyMapping is opaque storage with the size and alignment C gives NXKeyMapping:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type NXKeyMapping [2]uint64

// See: https://developer.apple.com/documentation/iokit/nxmousebutton
type NXMouseButton = uint32

// See: https://developer.apple.com/documentation/iokit/nxmousescaling
// NXMouseScaling is opaque storage with the size and alignment C gives NXMouseScaling:
// 84 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 84 into.
type NXMouseScaling [21]uint32

// See: https://developer.apple.com/documentation/iokit/nxparsedkeymapping
// NXParsedKeyMapping is opaque storage with the size and alignment C gives NXParsedKeyMapping:
// 3552 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 3552 into.
type NXParsedKeyMapping [444]uint64

// See: https://developer.apple.com/documentation/iokit/nxpoint
// NXPoint is opaque storage with the size and alignment C gives NXPoint:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type NXPoint [2]uint32

// See: https://developer.apple.com/documentation/iokit/nxsize
// NXSize is opaque storage with the size and alignment C gives NXSize:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type NXSize [2]uint32

// See: https://developer.apple.com/documentation/iokit/nxtabletpointdata
// NXTabletPointData is opaque storage with the size and alignment C gives NXTabletPointData:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type NXTabletPointData [8]uint32

// See: https://developer.apple.com/documentation/iokit/nxtabletpointdataptr
type NXTabletPointDataPtr = uintptr

// See: https://developer.apple.com/documentation/iokit/nxtabletproximitydata
// NXTabletProximityData is opaque storage with the size and alignment C gives NXTabletProximityData:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type NXTabletProximityData [8]uint32

// See: https://developer.apple.com/documentation/iokit/nxtabletproximitydataptr
type NXTabletProximityDataPtr = uintptr

// See: https://developer.apple.com/documentation/iokit/nudclflags
type NuDCLFlags = uint32

// See: https://developer.apple.com/documentation/iokit/nudclreceivepacketref
type NuDCLReceivePacketRef uintptr

// See: https://developer.apple.com/documentation/iokit/nudclref
type NuDCLRef uintptr

// See: https://developer.apple.com/documentation/iokit/nudclsendpacketref
type NuDCLSendPacketRef uintptr

// See: https://developer.apple.com/documentation/iokit/nudclskipcycleref
type NuDCLSkipCycleRef uintptr

// See: https://developer.apple.com/documentation/iokit/osasyncreference
// OSAsyncReference is opaque storage with the size and alignment C gives OSAsyncReference:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type OSAsyncReference [8]uint32

// See: https://developer.apple.com/documentation/iokit/osasyncreference64
// OSAsyncReference64 is opaque storage with the size and alignment C gives OSAsyncReference64:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type OSAsyncReference64 [8]uint64

// See: https://developer.apple.com/documentation/iokit/osobjectref
type OSObjectRef = uint64

// See: https://developer.apple.com/documentation/iokit/report_luns_logical_unit_addressing
// REPORT_LUNS_LOGICAL_UNIT_ADDRESSING is opaque storage with the size and alignment C gives REPORT_LUNS_LOGICAL_UNIT_ADDRESSING:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type REPORT_LUNS_LOGICAL_UNIT_ADDRESSING [1]uint16

// See: https://developer.apple.com/documentation/iokit/report_luns_peripheral_device_addressing
// REPORT_LUNS_PERIPHERAL_DEVICE_ADDRESSING is opaque storage with the size and alignment C gives REPORT_LUNS_PERIPHERAL_DEVICE_ADDRESSING:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type REPORT_LUNS_PERIPHERAL_DEVICE_ADDRESSING [1]uint16

// See: https://developer.apple.com/documentation/iokit/rgbcolor
// RGBColor is opaque storage with the size and alignment C gives RGBColor:
// 6 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 6 into.
type RGBColor [3]uint16

// See: https://developer.apple.com/documentation/iokit/rgbcolorhdl
type RGBColorHdl = *RGBColorPtr

// See: https://developer.apple.com/documentation/iokit/rgbcolorptr
type RGBColorPtr = *applicationservices.RGBColor

// See: https://developer.apple.com/documentation/iokit/rawsensecode
type RawSenseCode = uint8

// See: https://developer.apple.com/documentation/iokit/regentryid
// RegEntryID is opaque storage with the size and alignment C gives RegEntryID:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type RegEntryID [4]uint64

// See: https://developer.apple.com/documentation/iokit/regentryidptr
type RegEntryIDPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/sbcmodepagecaching
// SBCModePageCaching is opaque storage with the size and alignment C gives SBCModePageCaching:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type SBCModePageCaching [20]byte

// See: https://developer.apple.com/documentation/iokit/sbcmodepageflexibledisk
// SBCModePageFlexibleDisk is opaque storage with the size and alignment C gives SBCModePageFlexibleDisk:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type SBCModePageFlexibleDisk [32]byte

// See: https://developer.apple.com/documentation/iokit/sbcmodepageformatdevice
// SBCModePageFormatDevice is opaque storage with the size and alignment C gives SBCModePageFormatDevice:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type SBCModePageFormatDevice [24]byte

// See: https://developer.apple.com/documentation/iokit/sbcmodepagerigiddiskgeometry
// SBCModePageRigidDiskGeometry is opaque storage with the size and alignment C gives SBCModePageRigidDiskGeometry:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type SBCModePageRigidDiskGeometry [24]byte

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
// SCSICmd_INQUIRY_PAGECx_Header is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_PAGECx_Header:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type SCSICmd_INQUIRY_PAGECx_Header [4]byte

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_page00_header
// SCSICmd_INQUIRY_Page00_Header is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_Page00_Header:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type SCSICmd_INQUIRY_Page00_Header [4]byte

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_page00_header_spc_16
// SCSICmd_INQUIRY_Page00_Header_SPC_16 is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_Page00_Header_SPC_16:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type SCSICmd_INQUIRY_Page00_Header_SPC_16 [2]uint16

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_page80_header
// SCSICmd_INQUIRY_Page80_Header is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_Page80_Header:
// 5 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 5 into.
type SCSICmd_INQUIRY_Page80_Header [5]byte

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_page80_header_spc_16
// SCSICmd_INQUIRY_Page80_Header_SPC_16 is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_Page80_Header_SPC_16:
// 6 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 6 into.
type SCSICmd_INQUIRY_Page80_Header_SPC_16 [3]uint16

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_page83_header
// SCSICmd_INQUIRY_Page83_Header is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_Page83_Header:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type SCSICmd_INQUIRY_Page83_Header [4]byte

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_page83_header_spc_16
// SCSICmd_INQUIRY_Page83_Header_SPC_16 is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_Page83_Header_SPC_16:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type SCSICmd_INQUIRY_Page83_Header_SPC_16 [2]uint16

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_page83_identification_descriptor
// SCSICmd_INQUIRY_Page83_Identification_Descriptor is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_Page83_Identification_Descriptor:
// 5 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 5 into.
type SCSICmd_INQUIRY_Page83_Identification_Descriptor [5]byte

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_page83_logicalunitgroup_identifier
// SCSICmd_INQUIRY_Page83_LogicalUnitGroup_Identifier is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_Page83_LogicalUnitGroup_Identifier:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type SCSICmd_INQUIRY_Page83_LogicalUnitGroup_Identifier [2]uint16

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_page83_relativetargetport_identifier
// SCSICmd_INQUIRY_Page83_RelativeTargetPort_Identifier is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_Page83_RelativeTargetPort_Identifier:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type SCSICmd_INQUIRY_Page83_RelativeTargetPort_Identifier [2]uint16

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_page83_targetportgroup_identifier
// SCSICmd_INQUIRY_Page83_TargetPortGroup_Identifier is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_Page83_TargetPortGroup_Identifier:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type SCSICmd_INQUIRY_Page83_TargetPortGroup_Identifier [2]uint16

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_page89_data
// SCSICmd_INQUIRY_Page89_Data is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_Page89_Data:
// 572 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 572 into.
type SCSICmd_INQUIRY_Page89_Data [143]uint32

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_pageb0_data
// SCSICmd_INQUIRY_PageB0_Data is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_PageB0_Data:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type SCSICmd_INQUIRY_PageB0_Data [64]byte

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_pageb1_data
// SCSICmd_INQUIRY_PageB1_Data is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_PageB1_Data:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type SCSICmd_INQUIRY_PageB1_Data [32]uint16

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_pageb2_data
// SCSICmd_INQUIRY_PageB2_Data is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_PageB2_Data:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type SCSICmd_INQUIRY_PageB2_Data [8]byte

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_pageb2_provisioning_group_descriptor
// SCSICmd_INQUIRY_PageB2_Provisioning_Group_Descriptor is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_PageB2_Provisioning_Group_Descriptor:
// 38 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 38 into.
type SCSICmd_INQUIRY_PageB2_Provisioning_Group_Descriptor [38]byte

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_pagec0_data
// SCSICmd_INQUIRY_PageC0_Data is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_PageC0_Data:
// 116 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 116 into.
type SCSICmd_INQUIRY_PageC0_Data [116]byte

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_pagec1_data
// SCSICmd_INQUIRY_PageC1_Data is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_PageC1_Data:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type SCSICmd_INQUIRY_PageC1_Data [12]byte

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_standarddata
// SCSICmd_INQUIRY_StandardData is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_StandardData:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type SCSICmd_INQUIRY_StandardData [36]byte

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_standarddataall
// SCSICmd_INQUIRY_StandardDataAll is opaque storage with the size and alignment C gives SCSICmd_INQUIRY_StandardDataAll:
// 256 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 256 into.
type SCSICmd_INQUIRY_StandardDataAll [128]uint16

// See: https://developer.apple.com/documentation/iokit/scsicmd_inquiry_standarddataptr
type SCSICmd_INQUIRY_StandardDataPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/scsicmd_report_luns_header
// SCSICmd_REPORT_LUNS_Header is opaque storage with the size and alignment C gives SCSICmd_REPORT_LUNS_Header:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type SCSICmd_REPORT_LUNS_Header [4]uint32

// See: https://developer.apple.com/documentation/iokit/scsicmd_report_luns_lun_entry
// SCSICmd_REPORT_LUNS_LUN_ENTRY is opaque storage with the size and alignment C gives SCSICmd_REPORT_LUNS_LUN_ENTRY:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type SCSICmd_REPORT_LUNS_LUN_ENTRY [4]uint16

// See: https://developer.apple.com/documentation/iokit/scsicommanddescriptorblock
type SCSICommandDescriptorBlock = uint8

// SCSIDeviceIdentifier is 64-bit number to represent a SCSI Device.
//
// See: https://developer.apple.com/documentation/iokit/scsideviceidentifier
type SCSIDeviceIdentifier = uint64

// SCSIInitiatorIdentifier is 64-bit number to represent a SCSI Initiator Device.
//
// See: https://developer.apple.com/documentation/iokit/scsiinitiatoridentifier
type SCSIInitiatorIdentifier = uint64

// See: https://developer.apple.com/documentation/iokit/scsilogicalunitbytes
type SCSILogicalUnitBytes = uint8

// See: https://developer.apple.com/documentation/iokit/scsilogicalunitnumber
type SCSILogicalUnitNumber = uint64

// SCSIServiceResponse is attributes for task service response.
//
// See: https://developer.apple.com/documentation/iokit/scsiserviceresponse
type SCSIServiceResponse = uint32

// SCSITaggedTaskIdentifier is 64-bit number to represent a unique task identifier.
//
// See: https://developer.apple.com/documentation/iokit/scsitaggedtaskidentifier
type SCSITaggedTaskIdentifier = uint64

// SCSITargetIdentifier is 64-bit number to represent a SCSI Target Device.
//
// See: https://developer.apple.com/documentation/iokit/scsitargetidentifier
type SCSITargetIdentifier = uint64

// SCSITaskAttribute is attributes for task delivery.
//
// See: https://developer.apple.com/documentation/iokit/scsitaskattribute
type SCSITaskAttribute = uint32

// SCSITaskDeviceInterface is basic interface for a SCSITask Device.
//
// See: https://developer.apple.com/documentation/iokit/scsitaskdeviceinterface
// SCSITaskDeviceInterface is opaque storage with the size and alignment C gives SCSITaskDeviceInterface:
// 88 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 88 into.
type SCSITaskDeviceInterface [11]uint64

// SCSITaskInterface is basic interface for a SCSITask.
//
// See: https://developer.apple.com/documentation/iokit/scsitaskinterface
// SCSITaskInterface is opaque storage with the size and alignment C gives SCSITaskInterface:
// 200 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 200 into.
type SCSITaskInterface [25]uint64

// See: https://developer.apple.com/documentation/iokit/scsitasksgelement
// SCSITaskSGElement is opaque storage with the size and alignment C gives SCSITaskSGElement:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type SCSITaskSGElement [2]uint64

// SCSITaskState is attributes for task state.
//
// See: https://developer.apple.com/documentation/iokit/scsitaskstate
type SCSITaskState = uint32

// SCSITaskStatus is attributes for task status.
//
// See: https://developer.apple.com/documentation/iokit/scsitaskstatus
type SCSITaskStatus = uint32

// See: https://developer.apple.com/documentation/iokit/scsi_capacity_data
// SCSI_Capacity_Data is opaque storage with the size and alignment C gives SCSI_Capacity_Data:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type SCSI_Capacity_Data [2]uint32

// See: https://developer.apple.com/documentation/iokit/scsi_capacity_data_long
// SCSI_Capacity_Data_Long is opaque storage with the size and alignment C gives SCSI_Capacity_Data_Long:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type SCSI_Capacity_Data_Long [4]uint64

// See: https://developer.apple.com/documentation/iokit/scsi_sense_data
// SCSI_Sense_Data is opaque storage with the size and alignment C gives SCSI_Sense_Data:
// 18 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 18 into.
type SCSI_Sense_Data [18]byte

// See: https://developer.apple.com/documentation/iokit/spcmodepagepowercondition
// SPCModePagePowerCondition is opaque storage with the size and alignment C gives SPCModePagePowerCondition:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type SPCModePagePowerCondition [12]byte

// See: https://developer.apple.com/documentation/iokit/spcmodeparameterheader10
// SPCModeParameterHeader10 is opaque storage with the size and alignment C gives SPCModeParameterHeader10:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type SPCModeParameterHeader10 [8]byte

// See: https://developer.apple.com/documentation/iokit/spcmodeparameterheader6
// SPCModeParameterHeader6 is opaque storage with the size and alignment C gives SPCModeParameterHeader6:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type SPCModeParameterHeader6 [4]byte

// See: https://developer.apple.com/documentation/iokit/uaspipedescriptor
// UASPipeDescriptor is opaque storage with the size and alignment C gives UASPipeDescriptor:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type UASPipeDescriptor [4]byte

// See: https://developer.apple.com/documentation/iokit/uaspipedescriptorptr
type UASPipeDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/usbdeviceaddress
type USBDeviceAddress = uint16

// See: https://developer.apple.com/documentation/iokit/usbdeviceinformationbits
type USBDeviceInformationBits = uint32

// See: https://developer.apple.com/documentation/iokit/usblowlatencybuffertype
type USBLowLatencyBufferType = uint32

// See: https://developer.apple.com/documentation/iokit/usbnotificationtypes
type USBNotificationTypes = uint32

// See: https://developer.apple.com/documentation/iokit/usbphysicaladdress32
type USBPhysicalAddress32 = uint32

// See: https://developer.apple.com/documentation/iokit/usbpowerrequesttypes
type USBPowerRequestTypes = uint32

// See: https://developer.apple.com/documentation/iokit/usbreenumerateoptions
type USBReEnumerateOptions = int32

// See: https://developer.apple.com/documentation/iokit/usbstatus
type USBStatus = uint16

// See: https://developer.apple.com/documentation/iokit/usbstatusptr
type USBStatusPtr = *USBStatus

// See: https://developer.apple.com/documentation/iokit/userexportdclcallcommandproc
type UserExportDCLCallCommandProc = *kernel.ID

// See: https://developer.apple.com/documentation/iokit/userexportdclcallproc
// UserExportDCLCallProc is opaque storage with the size and alignment C gives UserExportDCLCallProc:
// 44 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 44 into.
type UserExportDCLCallProc [44]byte

// See: https://developer.apple.com/documentation/iokit/userexportdclcommand
// UserExportDCLCommand is opaque storage with the size and alignment C gives UserExportDCLCommand:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type UserExportDCLCommand [32]byte

// See: https://developer.apple.com/documentation/iokit/userexportdcljump
// UserExportDCLJump is opaque storage with the size and alignment C gives UserExportDCLJump:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type UserExportDCLJump [36]byte

// See: https://developer.apple.com/documentation/iokit/userexportdcllabel
// UserExportDCLLabel is opaque storage with the size and alignment C gives UserExportDCLLabel:
// 28 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 28 into.
type UserExportDCLLabel [28]byte

// See: https://developer.apple.com/documentation/iokit/userexportdclnudclleader
// UserExportDCLNuDCLLeader is opaque storage with the size and alignment C gives UserExportDCLNuDCLLeader:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type UserExportDCLNuDCLLeader [36]byte

// See: https://developer.apple.com/documentation/iokit/userexportdclptrtimestamp
// UserExportDCLPtrTimeStamp is opaque storage with the size and alignment C gives UserExportDCLPtrTimeStamp:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type UserExportDCLPtrTimeStamp [36]byte

// See: https://developer.apple.com/documentation/iokit/userexportdclsettagsyncbits
// UserExportDCLSetTagSyncBits is opaque storage with the size and alignment C gives UserExportDCLSetTagSyncBits:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type UserExportDCLSetTagSyncBits [32]byte

// See: https://developer.apple.com/documentation/iokit/userexportdcltimestamp
// UserExportDCLTimeStamp is opaque storage with the size and alignment C gives UserExportDCLTimeStamp:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type UserExportDCLTimeStamp [32]byte

// See: https://developer.apple.com/documentation/iokit/userexportdcltransferbuffer
// UserExportDCLTransferBuffer is opaque storage with the size and alignment C gives UserExportDCLTransferBuffer:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type UserExportDCLTransferBuffer [48]byte

// See: https://developer.apple.com/documentation/iokit/userexportdcltransferpacket
// UserExportDCLTransferPacket is opaque storage with the size and alignment C gives UserExportDCLTransferPacket:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type UserExportDCLTransferPacket [40]byte

// See: https://developer.apple.com/documentation/iokit/userexportdclupdatedcllist
// UserExportDCLUpdateDCLList is opaque storage with the size and alignment C gives UserExportDCLUpdateDCLList:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type UserExportDCLUpdateDCLList [40]byte

// See: https://developer.apple.com/documentation/iokit/vdclutbehavior
type VDClutBehavior = uint32

// See: https://developer.apple.com/documentation/iokit/vdclutbehaviorptr
type VDClutBehaviorPtr = *VDClutBehavior

// See: https://developer.apple.com/documentation/iokit/vdcommunicationinfoptr
type VDCommunicationInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdcommunicationinforec
// VDCommunicationInfoRec is opaque storage with the size and alignment C gives VDCommunicationInfoRec:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type VDCommunicationInfoRec [12]uint32

// See: https://developer.apple.com/documentation/iokit/vdcommunicationptr
type VDCommunicationPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdcommunicationrec
// VDCommunicationRec is opaque storage with the size and alignment C gives VDCommunicationRec:
// 80 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 80 into.
type VDCommunicationRec [10]uint64

// See: https://developer.apple.com/documentation/iokit/vdconfigurationfeaturelistrec
// VDConfigurationFeatureListRec is opaque storage with the size and alignment C gives VDConfigurationFeatureListRec:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type VDConfigurationFeatureListRec [4]uint64

// See: https://developer.apple.com/documentation/iokit/vdconfigurationfeaturelistrecptr
type VDConfigurationFeatureListRecPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdconfigurationptr
type VDConfigurationPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdconfigurationrec
// VDConfigurationRec is opaque storage with the size and alignment C gives VDConfigurationRec:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type VDConfigurationRec [4]uint64

// See: https://developer.apple.com/documentation/iokit/vdconvolutioninfoptr
type VDConvolutionInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdconvolutioninforec
// VDConvolutionInfoRec is opaque storage with the size and alignment C gives VDConvolutionInfoRec:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type VDConvolutionInfoRec [5]uint32

// See: https://developer.apple.com/documentation/iokit/vdddcblockptr
type VDDDCBlockPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdddcblockrec
// VDDDCBlockRec is opaque storage with the size and alignment C gives VDDDCBlockRec:
// 144 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 144 into.
type VDDDCBlockRec [36]uint32

// See: https://developer.apple.com/documentation/iokit/vddefmode
type VDDefMode = uint

// See: https://developer.apple.com/documentation/iokit/vddefmodeptr
type VDDefModePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vddetailedtimingptr
type VDDetailedTimingPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vddetailedtimingrec
// VDDetailedTimingRec is opaque storage with the size and alignment C gives VDDetailedTimingRec:
// 160 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 160 into.
type VDDetailedTimingRec [20]uint64

// See: https://developer.apple.com/documentation/iokit/vddisplayconnectinfoptr
type VDDisplayConnectInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vddisplayconnectinforec
// VDDisplayConnectInfoRec is opaque storage with the size and alignment C gives VDDisplayConnectInfoRec:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type VDDisplayConnectInfoRec [3]uint64

// See: https://developer.apple.com/documentation/iokit/vddisplaytimingrangeptr
type VDDisplayTimingRangePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vddisplaytimingrangerec
// VDDisplayTimingRangeRec is opaque storage with the size and alignment C gives VDDisplayTimingRangeRec:
// 240 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 240 into.
type VDDisplayTimingRangeRec [30]uint64

// See: https://developer.apple.com/documentation/iokit/vddrawhardwarecursorptr
type VDDrawHardwareCursorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vddrawhardwarecursorrec
// VDDrawHardwareCursorRec is opaque storage with the size and alignment C gives VDDrawHardwareCursorRec:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type VDDrawHardwareCursorRec [5]uint32

// See: https://developer.apple.com/documentation/iokit/vdentrecptr
type VDEntRecPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdentryrecord
// VDEntryRecord is opaque storage with the size and alignment C gives VDEntryRecord:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type VDEntryRecord [1]uint64

// See: https://developer.apple.com/documentation/iokit/vdflagrecptr
type VDFlagRecPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdflagrecord
// VDFlagRecord is opaque storage with the size and alignment C gives VDFlagRecord:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type VDFlagRecord [2]byte

// VDGamRecPtr is represents a type used by the Video Components API.
//
// See: https://developer.apple.com/documentation/iokit/vdgamrecptr
type VDGamRecPtr = *applicationservices.VDGammaRecord

// See: https://developer.apple.com/documentation/iokit/vdgammainfoptr
type VDGammaInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdgammainforec
// VDGammaInfoRec is opaque storage with the size and alignment C gives VDGammaInfoRec:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type VDGammaInfoRec [3]uint64

// See: https://developer.apple.com/documentation/iokit/vdgammarecord
// VDGammaRecord is opaque storage with the size and alignment C gives VDGammaRecord:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type VDGammaRecord [1]uint64

// See: https://developer.apple.com/documentation/iokit/vdgetgammalistptr
type VDGetGammaListPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdgetgammalistrec
// VDGetGammaListRec is opaque storage with the size and alignment C gives VDGetGammaListRec:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type VDGetGammaListRec [3]uint64

// See: https://developer.apple.com/documentation/iokit/vdgrayptr
type VDGrayPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdgrayrecord
// VDGrayRecord is opaque storage with the size and alignment C gives VDGrayRecord:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type VDGrayRecord [2]byte

// See: https://developer.apple.com/documentation/iokit/vdhardwarecursordrawstateptr
type VDHardwareCursorDrawStatePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdhardwarecursordrawstaterec
// VDHardwareCursorDrawStateRec is opaque storage with the size and alignment C gives VDHardwareCursorDrawStateRec:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type VDHardwareCursorDrawStateRec [6]uint32

// See: https://developer.apple.com/documentation/iokit/vdmirrorptr
type VDMirrorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdmirrorrec
// VDMirrorRec is opaque storage with the size and alignment C gives VDMirrorRec:
// 104 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 104 into.
type VDMirrorRec [13]uint64

// See: https://developer.apple.com/documentation/iokit/vdmulticonnectinfoptr
type VDMultiConnectInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdmulticonnectinforec
// VDMultiConnectInfoRec is opaque storage with the size and alignment C gives VDMultiConnectInfoRec:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type VDMultiConnectInfoRec [4]uint64

// See: https://developer.apple.com/documentation/iokit/vdpageinfo
// VDPageInfo is opaque storage with the size and alignment C gives VDPageInfo:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type VDPageInfo [3]uint64

// See: https://developer.apple.com/documentation/iokit/vdpginfoptr
type VDPgInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdpowerstateptr
type VDPowerStatePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdpowerstaterec
// VDPowerStateRec is opaque storage with the size and alignment C gives VDPowerStateRec:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type VDPowerStateRec [3]uint64

// See: https://developer.apple.com/documentation/iokit/vdprivateselectordatarec
// VDPrivateSelectorDataRec is opaque storage with the size and alignment C gives VDPrivateSelectorDataRec:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type VDPrivateSelectorDataRec [4]uint64

// See: https://developer.apple.com/documentation/iokit/vdprivateselectorrec
// VDPrivateSelectorRec is opaque storage with the size and alignment C gives VDPrivateSelectorRec:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type VDPrivateSelectorRec [5]uint64

// See: https://developer.apple.com/documentation/iokit/vdresolutioninfoptr
type VDResolutionInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdresolutioninforec
// VDResolutionInfoRec is opaque storage with the size and alignment C gives VDResolutionInfoRec:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type VDResolutionInfoRec [5]uint64

// See: https://developer.apple.com/documentation/iokit/vdretrievegammaptr
type VDRetrieveGammaPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdretrievegammarec
// VDRetrieveGammaRec is opaque storage with the size and alignment C gives VDRetrieveGammaRec:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type VDRetrieveGammaRec [2]uint64

// See: https://developer.apple.com/documentation/iokit/vdscalerinfoptr
type VDScalerInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdscalerinforec
// VDScalerInfoRec is opaque storage with the size and alignment C gives VDScalerInfoRec:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type VDScalerInfoRec [12]uint32

// See: https://developer.apple.com/documentation/iokit/vdscalerptr
type VDScalerPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdscalerrec
// VDScalerRec is opaque storage with the size and alignment C gives VDScalerRec:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type VDScalerRec [16]uint32

// See: https://developer.apple.com/documentation/iokit/vdsetentryptr
type VDSetEntryPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdsetentryrecord
// VDSetEntryRecord is opaque storage with the size and alignment C gives VDSetEntryRecord:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type VDSetEntryRecord [2]uint64

// See: https://developer.apple.com/documentation/iokit/vdsethardwarecursorptr
type VDSetHardwareCursorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdsethardwarecursorrec
// VDSetHardwareCursorRec is opaque storage with the size and alignment C gives VDSetHardwareCursorRec:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type VDSetHardwareCursorRec [2]uint64

// See: https://developer.apple.com/documentation/iokit/vdsettings
// VDSettings is opaque storage with the size and alignment C gives VDSettings:
// 38 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 38 into.
type VDSettings [19]uint16

// See: https://developer.apple.com/documentation/iokit/vdsettingsptr
type VDSettingsPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdsizeinfo
// VDSizeInfo is opaque storage with the size and alignment C gives VDSizeInfo:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type VDSizeInfo [4]uint16

// See: https://developer.apple.com/documentation/iokit/vdsupportshardwarecursorptr
type VDSupportsHardwareCursorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdsupportshardwarecursorrec
// VDSupportsHardwareCursorRec is opaque storage with the size and alignment C gives VDSupportsHardwareCursorRec:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type VDSupportsHardwareCursorRec [3]uint32

// See: https://developer.apple.com/documentation/iokit/vdswitchinfoptr
type VDSwitchInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdswitchinforec
// VDSwitchInfoRec is opaque storage with the size and alignment C gives VDSwitchInfoRec:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type VDSwitchInfoRec [4]uint64

// See: https://developer.apple.com/documentation/iokit/vdsyncinfoptr
type VDSyncInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdsyncinforec
// VDSyncInfoRec is opaque storage with the size and alignment C gives VDSyncInfoRec:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type VDSyncInfoRec [2]byte

// See: https://developer.apple.com/documentation/iokit/vdszinfoptr
type VDSzInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdtiminginfoptr
type VDTimingInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdtiminginforec
// VDTimingInfoRec is opaque storage with the size and alignment C gives VDTimingInfoRec:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type VDTimingInfoRec [4]uint64

// See: https://developer.apple.com/documentation/iokit/vdvideoparametersinfoptr
type VDVideoParametersInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/vdvideoparametersinforec
// VDVideoParametersInfoRec is opaque storage with the size and alignment C gives VDVideoParametersInfoRec:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type VDVideoParametersInfoRec [4]uint64

// See: https://developer.apple.com/documentation/iokit/vpblock
// VPBlock is opaque storage with the size and alignment C gives VPBlock:
// 44 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 44 into.
type VPBlock [11]uint32

// See: https://developer.apple.com/documentation/iokit/vpblockptr
type VPBlockPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/videodevicetype
type VideoDeviceType = uint32

// See: https://developer.apple.com/documentation/iokit/dk_bd_read_disc_info_t
// Dk_bd_read_disc_info_t is opaque storage with the size and alignment C gives dk_bd_read_disc_info_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_bd_read_disc_info_t [3]uint64

// See: https://developer.apple.com/documentation/iokit/dk_bd_read_structure_t
// Dk_bd_read_structure_t is opaque storage with the size and alignment C gives dk_bd_read_structure_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_bd_read_structure_t [3]uint64

// See: https://developer.apple.com/documentation/iokit/dk_bd_read_track_info_t
// Dk_bd_read_track_info_t is opaque storage with the size and alignment C gives dk_bd_read_track_info_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_bd_read_track_info_t [3]uint64

// See: https://developer.apple.com/documentation/iokit/dk_bd_report_key_t
// Dk_bd_report_key_t is opaque storage with the size and alignment C gives dk_bd_report_key_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_bd_report_key_t [3]uint64

// See: https://developer.apple.com/documentation/iokit/dk_bd_send_key_t
// Dk_bd_send_key_t is opaque storage with the size and alignment C gives dk_bd_send_key_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_bd_send_key_t [3]uint64

// See: https://developer.apple.com/documentation/iokit/dk_cd_read_disc_info_t
// Dk_cd_read_disc_info_t is opaque storage with the size and alignment C gives dk_cd_read_disc_info_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_cd_read_disc_info_t [3]uint64

// See: https://developer.apple.com/documentation/iokit/dk_cd_read_isrc_t
// Dk_cd_read_isrc_t is opaque storage with the size and alignment C gives dk_cd_read_isrc_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Dk_cd_read_isrc_t [16]byte

// See: https://developer.apple.com/documentation/iokit/dk_cd_read_mcn_t
// Dk_cd_read_mcn_t is opaque storage with the size and alignment C gives dk_cd_read_mcn_t:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type Dk_cd_read_mcn_t [16]byte

// See: https://developer.apple.com/documentation/iokit/dk_cd_read_t
// Dk_cd_read_t is opaque storage with the size and alignment C gives dk_cd_read_t:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type Dk_cd_read_t [4]uint64

// See: https://developer.apple.com/documentation/iokit/dk_cd_read_toc_t
// Dk_cd_read_toc_t is opaque storage with the size and alignment C gives dk_cd_read_toc_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_cd_read_toc_t [3]uint64

// See: https://developer.apple.com/documentation/iokit/dk_cd_read_track_info_t
// Dk_cd_read_track_info_t is opaque storage with the size and alignment C gives dk_cd_read_track_info_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_cd_read_track_info_t [3]uint64

// See: https://developer.apple.com/documentation/iokit/dk_dvd_read_disc_info_t
// Dk_dvd_read_disc_info_t is opaque storage with the size and alignment C gives dk_dvd_read_disc_info_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_dvd_read_disc_info_t [3]uint64

// See: https://developer.apple.com/documentation/iokit/dk_dvd_read_rzone_info_t
// Dk_dvd_read_rzone_info_t is opaque storage with the size and alignment C gives dk_dvd_read_rzone_info_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_dvd_read_rzone_info_t [3]uint64

// See: https://developer.apple.com/documentation/iokit/dk_dvd_read_structure_t
// Dk_dvd_read_structure_t is opaque storage with the size and alignment C gives dk_dvd_read_structure_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_dvd_read_structure_t [3]uint64

// See: https://developer.apple.com/documentation/iokit/dk_dvd_report_key_t
// Dk_dvd_report_key_t is opaque storage with the size and alignment C gives dk_dvd_report_key_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_dvd_report_key_t [3]uint64

// See: https://developer.apple.com/documentation/iokit/dk_dvd_send_key_t
// Dk_dvd_send_key_t is opaque storage with the size and alignment C gives dk_dvd_send_key_t:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type Dk_dvd_send_key_t [3]uint64

// See: https://developer.apple.com/documentation/iokit/eioaccelsurfacelockbits
type EIOAccelSurfaceLockBits = uint32

// See: https://developer.apple.com/documentation/iokit/eioaccelsurfacemodebits
type EIOAccelSurfaceModeBits = uint32

// See: https://developer.apple.com/documentation/iokit/eioaccelsurfacescalebits
type EIOAccelSurfaceScaleBits = uint32

// See: https://developer.apple.com/documentation/iokit/eioaccelsurfaceshapebits
type EIOAccelSurfaceShapeBits = uint32

// See: https://developer.apple.com/documentation/iokit/eioaccelsurfacestatebits
type EIOAccelSurfaceStateBits = uint32

// See: https://developer.apple.com/documentation/iokit/eviospecialkeymsg_t
// EvioSpecialKeyMsg_t is an unresolved C aggregate typedef.
type EvioSpecialKeyMsg_t unsafe.Pointer

// See: https://developer.apple.com/documentation/iokit/evsioevsioccsindices
type EvsioEVSIOCCSIndices = uint32

// See: https://developer.apple.com/documentation/iokit/evsioevsioscsindices
type EvsioEVSIOSCSIndices = uint32

// See: https://developer.apple.com/documentation/iokit/io_connect_t
type Io_connect_t = uint32

// See: https://developer.apple.com/documentation/iokit/io_enumerator_t
type Io_enumerator_t = uint32

// See: https://developer.apple.com/documentation/iokit/io_ident_t
type Io_ident_t = uint32

// See: https://developer.apple.com/documentation/iokit/io_iterator_t
type Io_iterator_t = uint32

// See: https://developer.apple.com/documentation/iokit/io_object_t
type Io_object_t = uint32

// See: https://developer.apple.com/documentation/iokit/io_registry_entry_t
type Io_registry_entry_t = uint32

// See: https://developer.apple.com/documentation/iokit/io_service_t
type Io_service_t = uint32

// See: https://developer.apple.com/documentation/iokit/kusbconnectable
type KUSBConnectable = uint32

// See: https://developer.apple.com/documentation/iokit/kusbhostconnectortype
type KUSBHostConnectorType = uint32

// See: https://developer.apple.com/documentation/iokit/sleepwakenote
// SleepWakeNote is opaque storage with the size and alignment C gives sleepWakeNote:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type SleepWakeNote [4]uint64

// See: https://developer.apple.com/documentation/iokit/tiousbdescriptorsize
type TIOUSBDescriptorSize = uint32

// See: https://developer.apple.com/documentation/iokit/tiousbdescriptortype
type TIOUSBDescriptorType = uint32

// See: https://developer.apple.com/documentation/iokit/tiousbdevicecapabilitytype
type TIOUSBDeviceCapabilityType = uint32

// See: https://developer.apple.com/documentation/iokit/tiousbdevicerequestdirectionvalue
type TIOUSBDeviceRequestDirectionValue = uint32

// See: https://developer.apple.com/documentation/iokit/tiousbdevicerequestrecipientvalue
type TIOUSBDeviceRequestRecipientValue = uint32

// See: https://developer.apple.com/documentation/iokit/tiousbdevicerequesttypevalue
type TIOUSBDeviceRequestTypeValue = uint32

// See: https://developer.apple.com/documentation/iokit/tiousbendpointdirection
type TIOUSBEndpointDirection = uint32

// See: https://developer.apple.com/documentation/iokit/tiousbendpointsynchronizationtype
type TIOUSBEndpointSynchronizationType = uint32

// See: https://developer.apple.com/documentation/iokit/tiousbendpointtype
type TIOUSBEndpointType = uint32

// See: https://developer.apple.com/documentation/iokit/tiousbendpointusagetype
type TIOUSBEndpointUsageType = uint32

// See: https://developer.apple.com/documentation/iokit/tiousblanguageid
type TIOUSBLanguageID = uint32

// See: https://developer.apple.com/documentation/iokit/uext_object_t
type Uext_object_t = uint32

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
