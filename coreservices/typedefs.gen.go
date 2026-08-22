// Code generated from Apple documentation. DO NOT EDIT.

package coreservices

import (
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/security"
)

// AEAddressDesc is a descriptor that contains the address of an application, used to describe the target application for an Apple event.
//
// See: https://developer.apple.com/documentation/coreservices/aeaddressdesc
type AEAddressDesc = AEDesc

// AEArrayData is stores array information to be put into a descriptor listwith the [AEPutArray] functionor extracted from a descriptor list with the [AEGetArray] function.
//
// See: https://developer.apple.com/documentation/coreservices/1443170-aearraydata
// AEArrayData is opaque storage with the size and alignment C gives AEArrayData:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type AEArrayData [8]uint16

// AEArrayDataPointer is a pointer to a union of type [AEArrayData].
//
// See: https://developer.apple.com/documentation/coreservices/aearraydatapointer
type AEArrayDataPointer = unsafe.Pointer

// AEArrayType is stores a value that specifies an array type.
//
// See: https://developer.apple.com/documentation/coreservices/aearraytype
type AEArrayType = int8

// AEBuildErrorCode is represents syntax errors found by an Apple Event build routine.
//
// See: https://developer.apple.com/documentation/coreservices/aebuilderrorcode
type AEBuildErrorCode = uint32

// AECoerceDescProcPtr is defines a pointer to a function that coerces data stored in a descriptor. Your descriptor coercion callback function coerces the data from the passed descriptor to the specified type, returning the coerced data in a second descriptor.
//
// See: https://developer.apple.com/documentation/coreservices/aecoercedescprocptr
type AECoerceDescProcPtr = func(fromDesc unsafe.Pointer, toType uint32, handlerRefcon uintptr, toDesc unsafe.Pointer) int16

// AECoerceDescUPP is defines a data type for the universal procedure pointer for the [AECoerceDescProcPtr] callback function pointer.
//
// See: https://developer.apple.com/documentation/coreservices/aecoercedescupp
type AECoerceDescUPP = unsafe.Pointer

// AECoercePtrProcPtr is defines a pointer to a function that coerces data stored in a buffer. Your pointer coercion callback routine coerces the data from the passed buffer to the specified type, returning the coerced data in a descriptor.
//
// See: https://developer.apple.com/documentation/coreservices/aecoerceptrprocptr
type AECoercePtrProcPtr = func(typeCode uint32, dataPtr unsafe.Pointer, dataSize int, toType uint32, handlerRefcon uintptr, result unsafe.Pointer) int16

// AECoercePtrUPP is defines a data type for the universal procedure pointer for the [AECoercePtrProcPtr] callback function pointer.
//
// See: https://developer.apple.com/documentation/coreservices/aecoerceptrupp
type AECoercePtrUPP = unsafe.Pointer

// AECoercionHandlerUPP is defines a data type for the universal procedure pointer for the [AECoercionHandlerUPP] callback function pointer.
//
// See: https://developer.apple.com/documentation/coreservices/aecoercionhandlerupp
type AECoercionHandlerUPP = unsafe.Pointer

// AEDataStorage is a pointer to an opaque data type that provides storage for an [AEDesc] descriptor.
//
// See: https://developer.apple.com/documentation/coreservices/aedatastorage
type AEDataStorage = unsafe.Pointer

// AEDataStorageType is an opaque data type used to store data in Apple event descriptors.
//
// See: https://developer.apple.com/documentation/coreservices/aedatastoragetype
type AEDataStorageType = unsafe.Pointer

// AEDescList is a descriptor whose data consists of a list of one or more descriptors.
//
// See: https://developer.apple.com/documentation/coreservices/aedesclist
type AEDescList = AEDesc

// See: https://developer.apple.com/documentation/coreservices/aedescptr
type AEDescPtr = *AEDesc

// AEDisposeExternalProcPtr is defines a pointer to a function the Apple Event Manager calls to dispose of a descriptor created by the [AECreateDescFromExternalPtr] function. Your callback function disposes of the buffer you originally passed to that function.
//
// See: https://developer.apple.com/documentation/coreservices/aedisposeexternalprocptr
type AEDisposeExternalProcPtr = func(dataPtr unsafe.Pointer, dataLength int, refcon uintptr)

// AEDisposeExternalUPP is defines a universal procedure pointer to a function the Apple Event Manager calls to dispose of a descriptor created by the [AECreateDescFromExternalPtr] function.
//
// See: https://developer.apple.com/documentation/coreservices/aedisposeexternalupp
type AEDisposeExternalUPP = unsafe.Pointer

// AEEventClass is specifies the event class of an Apple event.
//
// See: https://developer.apple.com/documentation/coreservices/aeeventclass
type AEEventClass = uint32

// AEEventHandlerProcPtr is defines a pointer to a function that handles one or more Apple events. Your Apple event handler function performs any action requested by the Apple event, adds parameters to the reply Apple event if appropriate (possibly including error information), and returns a result code.
//
// See: https://developer.apple.com/documentation/coreservices/aeeventhandlerprocptr
type AEEventHandlerProcPtr = func(theAppleEvent unsafe.Pointer, reply unsafe.Pointer, handlerRefcon uintptr) int16

// AEEventHandlerUPP is defines a data type for the universal procedure pointer for the [AEEventHandlerUPP] callback function pointer.
//
// See: https://developer.apple.com/documentation/coreservices/aeeventhandlerupp
type AEEventHandlerUPP = unsafe.Pointer

// AEEventID is specifies the event ID of an Apple event.
//
// See: https://developer.apple.com/documentation/coreservices/aeeventid
type AEEventID = uint32

// AEEventSource is a data type for values that specify how an Apple event was delivered.
//
// See: https://developer.apple.com/documentation/coreservices/aeeventsource
type AEEventSource = int8

// AEKeyword is a four-character code that uniquely identifies a descriptor in an Apple event record or an Apple event.
//
// See: https://developer.apple.com/documentation/coreservices/aekeyword
type AEKeyword = uint32

// AERecord is a descriptor whose data is a list of keyword-specified descriptors.
//
// See: https://developer.apple.com/documentation/coreservices/aerecord
// AERecord is opaque storage with the size and alignment C gives AERecord:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type AERecord [6]uint16

// AERemoteProcessResolverCallback is defines a pointer to a function the Apple Event Manager calls when the asynchronous execution of a remote process resolver completes, either due to success or failure, after a call to the [AERemoteProcessResolverScheduleWithRunLoop] function. Your callback function can use the reference passed to it to get the remote process information.
//
// See: https://developer.apple.com/documentation/coreservices/aeremoteprocessresolvercallback
type AERemoteProcessResolverCallback = func(ref AERemoteProcessResolverRef, info unsafe.Pointer)

// AERemoteProcessResolverRef is an opaque reference to an object that encapsulates the mechanism for obtaining a list of processes running on a remote machine.
//
// See: https://developer.apple.com/documentation/coreservices/aeremoteprocessresolverref
type AERemoteProcessResolverRef uintptr

// AEReturnID is specifies a return ID for a created Apple event.
//
// See: https://developer.apple.com/documentation/coreservices/aereturnid
type AEReturnID = int16

// AESendMode is specify send preferences to the [AESend] function.
//
// See: https://developer.apple.com/documentation/coreservices/aesendmode
type AESendMode = int32

// AESendPriority is specifies the processing priority for a sent Apple event.
//
// See: https://developer.apple.com/documentation/coreservices/aesendpriority
type AESendPriority = int16

// AEStreamRef is an opaque data structure for storing stream-based descriptor data.
//
// See: https://developer.apple.com/documentation/coreservices/aestreamref
type AEStreamRef uintptr

// AETransactionID is specifies a transaction ID.
//
// See: https://developer.apple.com/documentation/coreservices/aetransactionid
type AETransactionID = int32

// See: https://developer.apple.com/documentation/coreservices/afpalternateaddress
// AFPAlternateAddress is opaque storage with the size and alignment C gives AFPAlternateAddress:
// 3 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 3 into.
type AFPAlternateAddress [3]byte

// See: https://developer.apple.com/documentation/coreservices/afpserversignature
type AFPServerSignature = uint8

// See: https://developer.apple.com/documentation/coreservices/afptagdata
// AFPTagData is opaque storage with the size and alignment C gives AFPTagData:
// 3 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 3 into.
type AFPTagData [3]byte

// See: https://developer.apple.com/documentation/coreservices/afpvolmountinfo
// AFPVolMountInfo is opaque storage with the size and alignment C gives AFPVolMountInfo:
// 168 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 168 into.
type AFPVolMountInfo [84]uint16

// See: https://developer.apple.com/documentation/coreservices/afpvolmountinfoptr
type AFPVolMountInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/afpxvolmountinfo
// AFPXVolMountInfo is opaque storage with the size and alignment C gives AFPXVolMountInfo:
// 206 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 206 into.
type AFPXVolMountInfo [103]uint16

// See: https://developer.apple.com/documentation/coreservices/afpxvolmountinfoptr
type AFPXVolMountInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/aiffloop
// AIFFLoop is opaque storage with the size and alignment C gives AIFFLoop:
// 6 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 6 into.
type AIFFLoop [3]uint16

// AliasHandle is abst_AliasHandle.
//
// See: https://developer.apple.com/documentation/coreservices/aliashandle
type AliasHandle = *AliasPtr

// AliasInfoType is defines the alias record information type used in the index parameter of [GetAliasInfo].
//
// See: https://developer.apple.com/documentation/coreservices/aliasinfotype
type AliasInfoType = int16

// AliasPtr is abst_AliasPtr.
//
// See: https://developer.apple.com/documentation/coreservices/aliasptr
type AliasPtr = unsafe.Pointer

// AliasRecord is defines an alias record.
//
// See: https://developer.apple.com/documentation/coreservices/aliasrecord
// AliasRecord is opaque storage with the size and alignment C gives AliasRecord:
// 6 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 6 into.
type AliasRecord [6]byte

// AppleEvent is a descriptor whose data is a list of descriptors containing both attributes and parameters that make up an Apple event.
//
// See: https://developer.apple.com/documentation/coreservices/appleevent
type AppleEvent = AEDesc

// See: https://developer.apple.com/documentation/coreservices/appleeventptr
type AppleEventPtr = *AEDesc

// See: https://developer.apple.com/documentation/coreservices/applicationspecificchunk
// ApplicationSpecificChunk is opaque storage with the size and alignment C gives ApplicationSpecificChunk:
// 14 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 14 into.
type ApplicationSpecificChunk [7]uint16

// See: https://developer.apple.com/documentation/coreservices/applicationspecificchunkptr
type ApplicationSpecificChunkPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/areaid
type AreaID = uintptr

// See: https://developer.apple.com/documentation/coreservices/audiorecordingchunk
// AudioRecordingChunk is opaque storage with the size and alignment C gives AudioRecordingChunk:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type AudioRecordingChunk [16]uint16

// See: https://developer.apple.com/documentation/coreservices/audiorecordingchunkptr
type AudioRecordingChunkPtr = unsafe.Pointer

// BigEndianFixed is protects a big-endian Fixed value from being changed bylittle-endian code.
//
// See: https://developer.apple.com/documentation/coreservices/bigendianfixed
// BigEndianFixed is opaque storage with the size and alignment C gives BigEndianFixed:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type BigEndianFixed [2]uint16

// BigEndianLong is protects a big-endian long value from being changed bylittle-endian code.
//
// See: https://developer.apple.com/documentation/coreservices/bigendianlong
// BigEndianLong is opaque storage with the size and alignment C gives BigEndianLong:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type BigEndianLong [4]uint16

// See: https://developer.apple.com/documentation/coreservices/bigendianostype
// BigEndianOSType is opaque storage with the size and alignment C gives BigEndianOSType:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type BigEndianOSType [2]uint16

// BigEndianShort is protects a big-endian short value from being changed bylittle-endian code.
//
// See: https://developer.apple.com/documentation/coreservices/bigendianshort
// BigEndianShort is opaque storage with the size and alignment C gives BigEndianShort:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type BigEndianShort [1]uint16

// See: https://developer.apple.com/documentation/coreservices/bigendianuint32
// BigEndianUInt32 is opaque storage with the size and alignment C gives BigEndianUInt32:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type BigEndianUInt32 [2]uint16

// BigEndianUnsignedFixed is protects a big-endian unsigned Fixed value from beingchanged by little-endian code.
//
// See: https://developer.apple.com/documentation/coreservices/bigendianunsignedfixed
// BigEndianUnsignedFixed is opaque storage with the size and alignment C gives BigEndianUnsignedFixed:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type BigEndianUnsignedFixed [2]uint16

// BigEndianUnsignedLong is protects a big-endian unsigned long value from being changedby little-endian code.
//
// See: https://developer.apple.com/documentation/coreservices/bigendianunsignedlong
// BigEndianUnsignedLong is opaque storage with the size and alignment C gives BigEndianUnsignedLong:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type BigEndianUnsignedLong [4]uint16

// BigEndianUnsignedShort is protects a big-endian unsigned short value from beingchanged by little-endian code.
//
// See: https://developer.apple.com/documentation/coreservices/bigendianunsignedshort
// BigEndianUnsignedShort is opaque storage with the size and alignment C gives BigEndianUnsignedShort:
// 2 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 2 into.
type BigEndianUnsignedShort [1]uint16

// See: https://developer.apple.com/documentation/coreservices/cscomponentsthreadmode
type CSComponentsThreadMode = uint32

// See: https://developer.apple.com/documentation/coreservices/csdiskspacerecoverycallback
type CSDiskSpaceRecoveryCallback = func(bool, uint64, corefoundation.CFErrorRef)

// See: https://developer.apple.com/documentation/coreservices/csdiskspacerecoveryoptions
type CSDiskSpaceRecoveryOptions = int32

// See: https://developer.apple.com/documentation/coreservices/csidentityauthorityref
type CSIdentityAuthorityRef uintptr

// See: https://developer.apple.com/documentation/coreservices/csidentityclass
type CSIdentityClass = int32

// See: https://developer.apple.com/documentation/coreservices/csidentityflags
type CSIdentityFlags = uint32

// See: https://developer.apple.com/documentation/coreservices/csidentityqueryevent
type CSIdentityQueryEvent = int

// See: https://developer.apple.com/documentation/coreservices/csidentityqueryflags
type CSIdentityQueryFlags = uint32

// See: https://developer.apple.com/documentation/coreservices/csidentityqueryreceiveeventcallback
type CSIdentityQueryReceiveEventCallback = func(CSIdentityQueryRef, CSIdentityQueryEvent, corefoundation.CFArrayRef, corefoundation.CFErrorRef, unsafe.Pointer)

// See: https://developer.apple.com/documentation/coreservices/csidentityqueryref
type CSIdentityQueryRef uintptr

// See: https://developer.apple.com/documentation/coreservices/csidentityquerystringcomparisonmethod
type CSIdentityQueryStringComparisonMethod = int

// See: https://developer.apple.com/documentation/coreservices/csidentityref
type CSIdentityRef uintptr

// See: https://developer.apple.com/documentation/coreservices/csidentitystatusupdatedcallback
type CSIdentityStatusUpdatedCallback = func(CSIdentityRef, corefoundation.CFIndex, corefoundation.CFErrorRef, unsafe.Pointer)

// See: https://developer.apple.com/documentation/coreservices/callingconventiontype
type CallingConventionType = uint16

// See: https://developer.apple.com/documentation/coreservices/catpositionrec
// CatPositionRec is opaque storage with the size and alignment C gives CatPositionRec:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type CatPositionRec [8]uint16

// See: https://developer.apple.com/documentation/coreservices/chunkheader
// ChunkHeader is opaque storage with the size and alignment C gives ChunkHeader:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type ChunkHeader [4]uint16

// CollatorRef is refers to an opaque object that encapsulates locale and collation information for the purpose of performing Unicode string comparison.
//
// See: https://developer.apple.com/documentation/coreservices/collatorref
type CollatorRef uintptr

// See: https://developer.apple.com/documentation/coreservices/collectionexceptionupp
type CollectionExceptionUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/collectionflattenupp
type CollectionFlattenUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/collectiontag
type CollectionTag = uint32

// See: https://developer.apple.com/documentation/coreservices/comment
// Comment is opaque storage with the size and alignment C gives Comment:
// 10 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 10 into.
type Comment [5]uint16

// See: https://developer.apple.com/documentation/coreservices/commentschunk
// CommentsChunk is opaque storage with the size and alignment C gives CommentsChunk:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type CommentsChunk [10]uint16

// See: https://developer.apple.com/documentation/coreservices/commentschunkptr
type CommentsChunkPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/commonchunk
// CommonChunk is opaque storage with the size and alignment C gives CommonChunk:
// 26 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 26 into.
type CommonChunk [13]uint16

// See: https://developer.apple.com/documentation/coreservices/commonchunkptr
type CommonChunkPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/componentaliasresource
// ComponentAliasResource is opaque storage with the size and alignment C gives ComponentAliasResource:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type ComponentAliasResource [32]uint16

// See: https://developer.apple.com/documentation/coreservices/componentdescription
// ComponentDescription is opaque storage with the size and alignment C gives ComponentDescription:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type ComponentDescription [10]uint16

// See: https://developer.apple.com/documentation/coreservices/componentfunctionupp
type ComponentFunctionUPP = unsafe.Pointer

// ComponentInstance is abst_ComponentInstance.
//
// See: https://developer.apple.com/documentation/coreservices/componentinstance
type ComponentInstance = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/componentinstancerecord
// ComponentInstanceRecord is opaque storage with the size and alignment C gives ComponentInstanceRecord:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type ComponentInstanceRecord [4]uint16

// See: https://developer.apple.com/documentation/coreservices/componentmpworkfunctionheaderrecord
// ComponentMPWorkFunctionHeaderRecord is opaque storage with the size and alignment C gives ComponentMPWorkFunctionHeaderRecord:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type ComponentMPWorkFunctionHeaderRecord [8]uint16

// See: https://developer.apple.com/documentation/coreservices/componentmpworkfunctionheaderrecordptr
type ComponentMPWorkFunctionHeaderRecordPtr = unsafe.Pointer

// ComponentMPWorkFunctionUPP is represents a type used by the Image Codec API.
//
// See: https://developer.apple.com/documentation/coreservices/componentmpworkfunctionupp
type ComponentMPWorkFunctionUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/componentparameters
// ComponentParameters is opaque storage with the size and alignment C gives ComponentParameters:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type ComponentParameters [8]uint16

// See: https://developer.apple.com/documentation/coreservices/componentplatforminfo
// ComponentPlatformInfo is opaque storage with the size and alignment C gives ComponentPlatformInfo:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type ComponentPlatformInfo [6]uint16

// See: https://developer.apple.com/documentation/coreservices/componentplatforminfoarray
// ComponentPlatformInfoArray is opaque storage with the size and alignment C gives ComponentPlatformInfoArray:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type ComponentPlatformInfoArray [8]uint16

// See: https://developer.apple.com/documentation/coreservices/componentrecord
// ComponentRecord is opaque storage with the size and alignment C gives ComponentRecord:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type ComponentRecord [4]uint16

// See: https://developer.apple.com/documentation/coreservices/componentresource
// ComponentResource is opaque storage with the size and alignment C gives ComponentResource:
// 44 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 44 into.
type ComponentResource [22]uint16

// See: https://developer.apple.com/documentation/coreservices/componentresourceextension
// ComponentResourceExtension is opaque storage with the size and alignment C gives ComponentResourceExtension:
// 10 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 10 into.
type ComponentResourceExtension [5]uint16

// See: https://developer.apple.com/documentation/coreservices/componentresourcehandle
type ComponentResourceHandle = *ComponentResourcePtr

// See: https://developer.apple.com/documentation/coreservices/componentresourceptr
type ComponentResourcePtr = unsafe.Pointer

// ComponentResult is abst_ComponentResult.
//
// See: https://developer.apple.com/documentation/coreservices/componentresult
type ComponentResult = int32

// See: https://developer.apple.com/documentation/coreservices/componentroutineupp
type ComponentRoutineUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/constfseventstreamref
type ConstFSEventStreamRef uintptr

// See: https://developer.apple.com/documentation/coreservices/constfsspecptr
type ConstFSSpecPtr = unsafe.Pointer

// ConstScriptCodeRunPtr is defines a constant script code run pointer.
//
// See: https://developer.apple.com/documentation/coreservices/constscriptcoderunptr
type ConstScriptCodeRunPtr = unsafe.Pointer

// ConstTextEncodingRunPtr is defines a constant text encoding run pointer.
//
// See: https://developer.apple.com/documentation/coreservices/consttextencodingrunptr
type ConstTextEncodingRunPtr = unsafe.Pointer

// ConstTextPtr is defines a constant text pointer.
//
// See: https://developer.apple.com/documentation/coreservices/consttextptr
type ConstTextPtr = *uint8

// ConstTextToUnicodeInfo is defines a constant text to Unicode converter object.
//
// See: https://developer.apple.com/documentation/coreservices/consttexttounicodeinfo
type ConstTextToUnicodeInfo = TextToUnicodeInfo

// ConstUniCharArrayPtr is defines a constant Unicode character array pointer.
//
// See: https://developer.apple.com/documentation/coreservices/constunichararrayptr
type ConstUniCharArrayPtr = *uint16

// ConstUnicodeMappingPtr is defines a constant Unicode mapping pointer.
//
// See: https://developer.apple.com/documentation/coreservices/constunicodemappingptr
type ConstUnicodeMappingPtr = unsafe.Pointer

// ConstUnicodeToTextInfo is defines a constant Unicode to text converter object.
//
// See: https://developer.apple.com/documentation/coreservices/constunicodetotextinfo
type ConstUnicodeToTextInfo = UnicodeToTextInfo

// See: https://developer.apple.com/documentation/coreservices/containerchunk
// ContainerChunk is opaque storage with the size and alignment C gives ContainerChunk:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type ContainerChunk [6]uint16

// See: https://developer.apple.com/documentation/coreservices/custombadgeresource
// CustomBadgeResource is opaque storage with the size and alignment C gives CustomBadgeResource:
// 28 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 28 into.
type CustomBadgeResource [14]uint16

// See: https://developer.apple.com/documentation/coreservices/custombadgeresourcehandle
type CustomBadgeResourceHandle = *CustomBadgeResourcePtr

// See: https://developer.apple.com/documentation/coreservices/custombadgeresourceptr
type CustomBadgeResourcePtr = unsafe.Pointer

// DCSDictionaryRef is an opaque object that represents a dictionary file.
//
// See: https://developer.apple.com/documentation/coreservices/dcsdictionaryref
type DCSDictionaryRef uintptr

// See: https://developer.apple.com/documentation/coreservices/dinfo
// DInfo is opaque storage with the size and alignment C gives DInfo:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type DInfo [8]uint16

// See: https://developer.apple.com/documentation/coreservices/dxinfo
// DXInfo is opaque storage with the size and alignment C gives DXInfo:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type DXInfo [8]uint16

// See: https://developer.apple.com/documentation/coreservices/datecacheptr
type DateCachePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/datecacherecord
// DateCacheRecord is opaque storage with the size and alignment C gives DateCacheRecord:
// 512 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 512 into.
type DateCacheRecord [256]uint16

// See: https://developer.apple.com/documentation/coreservices/datedelta
type DateDelta = int8

// See: https://developer.apple.com/documentation/coreservices/dateform
type DateForm = int8

// See: https://developer.apple.com/documentation/coreservices/dateorders
type DateOrders = int8

// See: https://developer.apple.com/documentation/coreservices/datetimerec
// DateTimeRec is opaque storage with the size and alignment C gives DateTimeRec:
// 14 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 14 into.
type DateTimeRec [7]uint16

// See: https://developer.apple.com/documentation/coreservices/debugassertoutputhandlerupp
type DebugAssertOutputHandlerUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/debugcomponentcallbackupp
type DebugComponentCallbackUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/debuggerdisposethreadtpp
type DebuggerDisposeThreadTPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/debuggerdisposethreadupp
type DebuggerDisposeThreadUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/debuggernewthreadtpp
type DebuggerNewThreadTPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/debuggernewthreadupp
type DebuggerNewThreadUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/debuggerthreadschedulertpp
type DebuggerThreadSchedulerTPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/debuggerthreadschedulerupp
type DebuggerThreadSchedulerUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/deferredtask
// DeferredTask is opaque storage with the size and alignment C gives DeferredTask:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type DeferredTask [18]uint16

// See: https://developer.apple.com/documentation/coreservices/deferredtaskptr
type DeferredTaskPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/deferredtaskupp
type DeferredTaskUPP = unsafe.Pointer

// DescType is specifies the type of the data stored in an [AEDesc] descriptor.
//
// See: https://developer.apple.com/documentation/coreservices/desctype
type DescType = uint32

// See: https://developer.apple.com/documentation/coreservices/exceptionhandler
type ExceptionHandler = uintptr

// See: https://developer.apple.com/documentation/coreservices/exceptionhandlertpp
type ExceptionHandlerTPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/exceptionhandlerupp
type ExceptionHandlerUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/exceptioninformation
// ExceptionInformation is opaque storage with the size and alignment C gives ExceptionInformation:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type ExceptionInformation [6]uint64

// See: https://developer.apple.com/documentation/coreservices/exceptioninformationpowerpc
// ExceptionInformationPowerPC is opaque storage with the size and alignment C gives ExceptionInformationPowerPC:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type ExceptionInformationPowerPC [6]uint64

// See: https://developer.apple.com/documentation/coreservices/exceptionkind
type ExceptionKind = uint

// See: https://developer.apple.com/documentation/coreservices/extcommonchunk
// ExtCommonChunk is opaque storage with the size and alignment C gives ExtCommonChunk:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type ExtCommonChunk [16]uint16

// See: https://developer.apple.com/documentation/coreservices/extcommonchunkptr
type ExtCommonChunkPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/extcomponentresource
// ExtComponentResource is opaque storage with the size and alignment C gives ExtComponentResource:
// 70 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 70 into.
type ExtComponentResource [35]uint16

// See: https://developer.apple.com/documentation/coreservices/extcomponentresourcehandle
type ExtComponentResourceHandle = *ExtComponentResourcePtr

// See: https://developer.apple.com/documentation/coreservices/extcomponentresourceptr
type ExtComponentResourcePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/extendedfileinfo
// ExtendedFileInfo is opaque storage with the size and alignment C gives ExtendedFileInfo:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type ExtendedFileInfo [8]uint16

// See: https://developer.apple.com/documentation/coreservices/extendedfolderinfo
// ExtendedFolderInfo is opaque storage with the size and alignment C gives ExtendedFolderInfo:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type ExtendedFolderInfo [8]uint16

// See: https://developer.apple.com/documentation/coreservices/finfo
// FInfo is opaque storage with the size and alignment C gives FInfo:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type FInfo [8]uint16

// See: https://developer.apple.com/documentation/coreservices/fnmessage
type FNMessage = uint32

// See: https://developer.apple.com/documentation/coreservices/fnsubscriptionref
type FNSubscriptionRef uintptr

// See: https://developer.apple.com/documentation/coreservices/fnsubscriptionupp
type FNSubscriptionUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fpregintel
type FPRegIntel = byte

// See: https://developer.apple.com/documentation/coreservices/fpuinformation
// FPUInformation is opaque storage with the size and alignment C gives FPUInformation:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type FPUInformation [1]uint64

// See: https://developer.apple.com/documentation/coreservices/fpuinformationintel64
// FPUInformationIntel64 is an unresolved C aggregate typedef.
type FPUInformationIntel64 unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fpuinformationpowerpc
// FPUInformationPowerPC is opaque storage with the size and alignment C gives FPUInformationPowerPC:
// 272 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 272 into.
type FPUInformationPowerPC [34]uint64

// FSAliasInfo is defines an information block passed to the [FSCopyAliasInfo] function.
//
// See: https://developer.apple.com/documentation/coreservices/fsaliasinfo
// FSAliasInfo is opaque storage with the size and alignment C gives FSAliasInfo:
// 42 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 42 into.
type FSAliasInfo [21]uint16

// FSAliasInfoBitmap is returned by the [FSCopyAliasInfo] function to indicate which fields of the alias information structure contain valid data.
//
// See: https://developer.apple.com/documentation/coreservices/fsaliasinfobitmap
type FSAliasInfoBitmap = uint32

// See: https://developer.apple.com/documentation/coreservices/fsaliasinfoptr
type FSAliasInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsallocationflags
type FSAllocationFlags = uint16

// See: https://developer.apple.com/documentation/coreservices/fscatalogbulkparam
// FSCatalogBulkParam is opaque storage with the size and alignment C gives FSCatalogBulkParam:
// 112 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 112 into.
type FSCatalogBulkParam [56]uint16

// See: https://developer.apple.com/documentation/coreservices/fscatalogbulkparamptr
type FSCatalogBulkParamPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fscataloginfo
// FSCatalogInfo is opaque storage with the size and alignment C gives FSCatalogInfo:
// 148 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 148 into.
type FSCatalogInfo [74]uint16

// See: https://developer.apple.com/documentation/coreservices/fscataloginfobitmap
type FSCatalogInfoBitmap = uint32

// See: https://developer.apple.com/documentation/coreservices/fscataloginfoptr
type FSCatalogInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsejectstatus
type FSEjectStatus = uint32

// See: https://developer.apple.com/documentation/coreservices/fseventstreamcallback
type FSEventStreamCallback = func(streamRef ConstFSEventStreamRef, clientCallBackInfo unsafe.Pointer, numEvents int32, eventPaths unsafe.Pointer, eventFlags unsafe.Pointer, eventIds unsafe.Pointer)

// See: https://developer.apple.com/documentation/coreservices/fseventstreamcreateflags
type FSEventStreamCreateFlags = uint32

// See: https://developer.apple.com/documentation/coreservices/fseventstreameventflags
type FSEventStreamEventFlags = uint32

// See: https://developer.apple.com/documentation/coreservices/fseventstreameventid
type FSEventStreamEventId = uint64

// See: https://developer.apple.com/documentation/coreservices/fseventstreamref
type FSEventStreamRef uintptr

// See: https://developer.apple.com/documentation/coreservices/fsfileoperationclientcontext
// FSFileOperationClientContext is opaque storage with the size and alignment C gives FSFileOperationClientContext:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type FSFileOperationClientContext [20]uint16

// See: https://developer.apple.com/documentation/coreservices/fsfileoperationref
type FSFileOperationRef uintptr

// See: https://developer.apple.com/documentation/coreservices/fsfileoperationstage
type FSFileOperationStage = uint32

// See: https://developer.apple.com/documentation/coreservices/fsfilesecurityref
type FSFileSecurityRef uintptr

// See: https://developer.apple.com/documentation/coreservices/fsforkcbinfoparam
// FSForkCBInfoParam is opaque storage with the size and alignment C gives FSForkCBInfoParam:
// 66 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 66 into.
type FSForkCBInfoParam [33]uint16

// See: https://developer.apple.com/documentation/coreservices/fsforkcbinfoparamptr
type FSForkCBInfoParamPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsforkioparam
// FSForkIOParam is opaque storage with the size and alignment C gives FSForkIOParam:
// 130 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 130 into.
type FSForkIOParam [65]uint16

// See: https://developer.apple.com/documentation/coreservices/fsforkioparamptr
type FSForkIOParamPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsforkinfo
// FSForkInfo is opaque storage with the size and alignment C gives FSForkInfo:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type FSForkInfo [24]uint16

// See: https://developer.apple.com/documentation/coreservices/fsforkinfoflags
type FSForkInfoFlags = uint8

// See: https://developer.apple.com/documentation/coreservices/fsforkinfoptr
type FSForkInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsiorefnum
type FSIORefNum = int32

// See: https://developer.apple.com/documentation/coreservices/fsiterator
type FSIterator = uintptr

// See: https://developer.apple.com/documentation/coreservices/fsiteratorflags
type FSIteratorFlags = uint32

// See: https://developer.apple.com/documentation/coreservices/fsmountstatus
type FSMountStatus = uint32

// See: https://developer.apple.com/documentation/coreservices/fspermissioninfo
// FSPermissionInfo is opaque storage with the size and alignment C gives FSPermissionInfo:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type FSPermissionInfo [10]uint16

// See: https://developer.apple.com/documentation/coreservices/fsrangelockparam
// FSRangeLockParam is opaque storage with the size and alignment C gives FSRangeLockParam:
// 60 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 60 into.
type FSRangeLockParam [30]uint16

// See: https://developer.apple.com/documentation/coreservices/fsrangelockparamptr
type FSRangeLockParamPtr = unsafe.Pointer

// FSRef is identifies a directory or file, including a volume’s root directory.
//
// See: https://developer.apple.com/documentation/coreservices/fsref
// FSRef is opaque storage with the size and alignment C gives FSRef:
// 80 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 80 into.
type FSRef [80]byte

// See: https://developer.apple.com/documentation/coreservices/fsrefforkioparam
// FSRefForkIOParam is opaque storage with the size and alignment C gives FSRefForkIOParam:
// 96 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 96 into.
type FSRefForkIOParam [48]uint16

// See: https://developer.apple.com/documentation/coreservices/fsrefforkioparamptr
type FSRefForkIOParamPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsrefparam
// FSRefParam is opaque storage with the size and alignment C gives FSRefParam:
// 120 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 120 into.
type FSRefParam [60]uint16

// See: https://developer.apple.com/documentation/coreservices/fsrefparamptr
type FSRefParamPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsrefptr
type FSRefPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fssearchparams
// FSSearchParams is opaque storage with the size and alignment C gives FSSearchParams:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type FSSearchParams [20]uint16

// See: https://developer.apple.com/documentation/coreservices/fssearchparamsptr
type FSSearchParamsPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsspec
// FSSpec is opaque storage with the size and alignment C gives FSSpec:
// 70 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 70 into.
type FSSpec [70]byte

// See: https://developer.apple.com/documentation/coreservices/fsspecarrayptr
type FSSpecArrayPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsspechandle
type FSSpecHandle = *FSSpecPtr

// See: https://developer.apple.com/documentation/coreservices/fsspecptr
type FSSpecPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsunmountstatus
type FSUnmountStatus = uint32

// See: https://developer.apple.com/documentation/coreservices/fsvolumeejectupp
type FSVolumeEjectUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsvolumeinfo
// FSVolumeInfo is opaque storage with the size and alignment C gives FSVolumeInfo:
// 128 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 128 into.
type FSVolumeInfo [64]uint16

// See: https://developer.apple.com/documentation/coreservices/fsvolumeinfobitmap
type FSVolumeInfoBitmap = uint32

// See: https://developer.apple.com/documentation/coreservices/fsvolumeinfoparam
// FSVolumeInfoParam is opaque storage with the size and alignment C gives FSVolumeInfoParam:
// 72 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 72 into.
type FSVolumeInfoParam [36]uint16

// See: https://developer.apple.com/documentation/coreservices/fsvolumeinfoparamptr
type FSVolumeInfoParamPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsvolumeinfoptr
type FSVolumeInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsvolumemountupp
type FSVolumeMountUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsvolumeoperation
type FSVolumeOperation = uintptr

// See: https://developer.apple.com/documentation/coreservices/fsvolumerefnum
type FSVolumeRefNum = int16

// See: https://developer.apple.com/documentation/coreservices/fsvolumeunmountupp
type FSVolumeUnmountUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fvector
// FVector is opaque storage with the size and alignment C gives FVector:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type FVector [2]uint16

// See: https://developer.apple.com/documentation/coreservices/fxinfo
// FXInfo is opaque storage with the size and alignment C gives FXInfo:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type FXInfo [8]uint16

// See: https://developer.apple.com/documentation/coreservices/fileinfo
// FileInfo is opaque storage with the size and alignment C gives FileInfo:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type FileInfo [8]uint16

// See: https://developer.apple.com/documentation/coreservices/folderclass
type FolderClass = uint32

// See: https://developer.apple.com/documentation/coreservices/folderdesc
// FolderDesc is opaque storage with the size and alignment C gives FolderDesc:
// 100 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 100 into.
type FolderDesc [50]uint16

// See: https://developer.apple.com/documentation/coreservices/folderdescflags
type FolderDescFlags = uint32

// See: https://developer.apple.com/documentation/coreservices/folderdescptr
type FolderDescPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/folderinfo
// FolderInfo is opaque storage with the size and alignment C gives FolderInfo:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type FolderInfo [8]uint16

// See: https://developer.apple.com/documentation/coreservices/folderlocation
type FolderLocation = uint32

// See: https://developer.apple.com/documentation/coreservices/foldermanagernotificationupp
type FolderManagerNotificationUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/folderrouting
// FolderRouting is opaque storage with the size and alignment C gives FolderRouting:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type FolderRouting [12]uint16

// See: https://developer.apple.com/documentation/coreservices/folderroutingptr
type FolderRoutingPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/foldertype
type FolderType = uint32

// See: https://developer.apple.com/documentation/coreservices/formatclass
type FormatClass = int8

// See: https://developer.apple.com/documentation/coreservices/formatresulttype
type FormatResultType = int8

// See: https://developer.apple.com/documentation/coreservices/formatstatus
type FormatStatus = int16

// See: https://developer.apple.com/documentation/coreservices/formatversionchunk
// FormatVersionChunk is opaque storage with the size and alignment C gives FormatVersionChunk:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type FormatVersionChunk [6]uint16

// See: https://developer.apple.com/documentation/coreservices/formatversionchunkptr
type FormatVersionChunkPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/getmissingcomponentresourceupp
type GetMissingComponentResourceUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/getvolparmsinfobuffer
// GetVolParmsInfoBuffer is opaque storage with the size and alignment C gives GetVolParmsInfoBuffer:
// 44 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 44 into.
type GetVolParmsInfoBuffer [22]uint16

// See: https://developer.apple.com/documentation/coreservices/hfscatalognodeid
type HFSCatalogNodeID = uint32

// See: https://developer.apple.com/documentation/coreservices/iocompletionupp
type IOCompletionUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/isatype
type ISAType = int8

// See: https://developer.apple.com/documentation/coreservices/iconfamilyelement
// IconFamilyElement is opaque storage with the size and alignment C gives IconFamilyElement:
// 10 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 10 into.
type IconFamilyElement [5]uint16

// See: https://developer.apple.com/documentation/coreservices/iconfamilyhandle
type IconFamilyHandle = *IconFamilyPtr

// See: https://developer.apple.com/documentation/coreservices/iconfamilyptr
type IconFamilyPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/iconfamilyresource
// IconFamilyResource is opaque storage with the size and alignment C gives IconFamilyResource:
// 18 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 18 into.
type IconFamilyResource [9]uint16

// See: https://developer.apple.com/documentation/coreservices/iconref
type IconRef uintptr

// See: https://developer.apple.com/documentation/coreservices/iconservicesusageflags
type IconServicesUsageFlags = uint32

// See: https://developer.apple.com/documentation/coreservices/indextoucstringprocptr
type IndexToUCStringProcPtr = func(uint32, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool

// See: https://developer.apple.com/documentation/coreservices/indextoucstringupp
type IndexToUCStringUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/instrumentchunk
// InstrumentChunk is opaque storage with the size and alignment C gives InstrumentChunk:
// 28 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 28 into.
type InstrumentChunk [14]uint16

// See: https://developer.apple.com/documentation/coreservices/instrumentchunkptr
type InstrumentChunkPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/intl0hndl
type Intl0Hndl = *Intl0Ptr

// See: https://developer.apple.com/documentation/coreservices/intl0ptr
type Intl0Ptr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/intl0rec
// Intl0Rec is opaque storage with the size and alignment C gives Intl0Rec:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type Intl0Rec [16]uint16

// See: https://developer.apple.com/documentation/coreservices/intl1hndl
type Intl1Hndl = *Intl1Ptr

// See: https://developer.apple.com/documentation/coreservices/intl1ptr
type Intl1Ptr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/intl1rec
// Intl1Rec is opaque storage with the size and alignment C gives Intl1Rec:
// 332 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 332 into.
type Intl1Rec [166]uint16

// See: https://developer.apple.com/documentation/coreservices/itl1extrec
// Itl1ExtRec is opaque storage with the size and alignment C gives Itl1ExtRec:
// 380 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 380 into.
type Itl1ExtRec [190]uint16

// See: https://developer.apple.com/documentation/coreservices/itl4handle
type Itl4Handle = *Itl4Ptr

// See: https://developer.apple.com/documentation/coreservices/itl4ptr
type Itl4Ptr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/itl4rec
// Itl4Rec is opaque storage with the size and alignment C gives Itl4Rec:
// 52 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 52 into.
type Itl4Rec [26]uint16

// See: https://developer.apple.com/documentation/coreservices/itl5record
// Itl5Record is opaque storage with the size and alignment C gives Itl5Record:
// 28 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 28 into.
type Itl5Record [14]uint16

// See: https://developer.apple.com/documentation/coreservices/itlbextrecord
// ItlbExtRecord is opaque storage with the size and alignment C gives ItlbExtRecord:
// 50 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 50 into.
type ItlbExtRecord [25]uint16

// See: https://developer.apple.com/documentation/coreservices/itlbrecord
// ItlbRecord is opaque storage with the size and alignment C gives ItlbRecord:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type ItlbRecord [10]uint16

// See: https://developer.apple.com/documentation/coreservices/itlcrecord
// ItlcRecord is opaque storage with the size and alignment C gives ItlcRecord:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type ItlcRecord [24]uint16

// See: https://developer.apple.com/documentation/coreservices/kcattrtype
type KCAttrType = uint32

// See: https://developer.apple.com/documentation/coreservices/kcattribute
type KCAttribute = security.SecKeychainAttribute

// See: https://developer.apple.com/documentation/coreservices/kcattributelist
type KCAttributeList = security.SecKeychainAttributeList

// See: https://developer.apple.com/documentation/coreservices/kcauthtype
type KCAuthType = uint32

// See: https://developer.apple.com/documentation/coreservices/kccallbackinfo
// KCCallbackInfo is opaque storage with the size and alignment C gives KCCallbackInfo:
// 44 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 44 into.
type KCCallbackInfo [22]uint16

// See: https://developer.apple.com/documentation/coreservices/kccallbackupp
type KCCallbackUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/kccertaddoptions
type KCCertAddOptions = uint32

// See: https://developer.apple.com/documentation/coreservices/kccertsearchoptions
type KCCertSearchOptions = uint32

// See: https://developer.apple.com/documentation/coreservices/kcevent
type KCEvent = uint16

// See: https://developer.apple.com/documentation/coreservices/kceventmask
type KCEventMask = uint16

// See: https://developer.apple.com/documentation/coreservices/kcitemattr
type KCItemAttr = uint32

// See: https://developer.apple.com/documentation/coreservices/kcitemclass
type KCItemClass = uint32

// See: https://developer.apple.com/documentation/coreservices/kcitemref
type KCItemRef = security.SecKeychainItemRef

// See: https://developer.apple.com/documentation/coreservices/kcprotocoltype
type KCProtocolType = uint32

// See: https://developer.apple.com/documentation/coreservices/kcpublickeyhash
type KCPublicKeyHash = uint8

// See: https://developer.apple.com/documentation/coreservices/kcref
type KCRef = security.SecKeychainRef

// See: https://developer.apple.com/documentation/coreservices/kcsearchref
type KCSearchRef = security.SecKeychainSearchRef

// See: https://developer.apple.com/documentation/coreservices/kcstatus
type KCStatus = uint32

// See: https://developer.apple.com/documentation/coreservices/kcverifystopon
type KCVerifyStopOn = uint16

// See: https://developer.apple.com/documentation/coreservices/lssharedfilelistchangedprocptr
type LSSharedFileListChangedProcPtr = func(LSSharedFileListRef, unsafe.Pointer)

// LSSharedFileListItemRef is a file-system object in the shared file list.
//
// See: https://developer.apple.com/documentation/coreservices/lssharedfilelistitemref
type LSSharedFileListItemRef uintptr

// LSSharedFileListRef is a persistent list of file-system objects.
//
// See: https://developer.apple.com/documentation/coreservices/lssharedfilelistref
type LSSharedFileListRef uintptr

// See: https://developer.apple.com/documentation/coreservices/lssharedfilelistresolutionflags
type LSSharedFileListResolutionFlags = uint32

// See: https://developer.apple.com/documentation/coreservices/localdatetime
// LocalDateTime is opaque storage with the size and alignment C gives LocalDateTime:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type LocalDateTime [4]uint16

// See: https://developer.apple.com/documentation/coreservices/localdatetimehandle
type LocalDateTimeHandle = *LocalDateTimePtr

// See: https://developer.apple.com/documentation/coreservices/localdatetimeptr
type LocalDateTimePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/localeandvariant
// LocaleAndVariant is opaque storage with the size and alignment C gives LocaleAndVariant:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type LocaleAndVariant [6]uint16

// See: https://developer.apple.com/documentation/coreservices/localenamemask
type LocaleNameMask = uint32

// See: https://developer.apple.com/documentation/coreservices/localeoperationclass
type LocaleOperationClass = uint32

// See: https://developer.apple.com/documentation/coreservices/localeoperationvariant
type LocaleOperationVariant = uint32

// See: https://developer.apple.com/documentation/coreservices/localepartmask
type LocalePartMask = uint32

// See: https://developer.apple.com/documentation/coreservices/localeref
type LocaleRef uintptr

// See: https://developer.apple.com/documentation/coreservices/longdatefield
type LongDateField = int8

// See: https://developer.apple.com/documentation/coreservices/longdatetime
type LongDateTime = int64

// MDItemRef is a reference to a MDItem object.
//
// See: https://developer.apple.com/documentation/coreservices/mditemref
type MDItemRef uintptr

// See: https://developer.apple.com/documentation/coreservices/mdlabelref
type MDLabelRef uintptr

// MDQueryCreateResultFunction is callback function used to create the result objects stored and returned by a query.
//
// See: https://developer.apple.com/documentation/coreservices/mdquerycreateresultfunction
type MDQueryCreateResultFunction = func(query MDQueryRef, item MDItemRef, context unsafe.Pointer) unsafe.Pointer

// MDQueryCreateValueFunction is callback function usedto create the value objects stored and returned by a query.
//
// See: https://developer.apple.com/documentation/coreservices/mdquerycreatevaluefunction
type MDQueryCreateValueFunction = func(query MDQueryRef, attrName corefoundation.CFStringRef, attrValue corefoundation.CFTypeRef, context unsafe.Pointer) unsafe.Pointer

// MDQueryRef is a reference to a MDQuery object.
//
// See: https://developer.apple.com/documentation/coreservices/mdqueryref
type MDQueryRef uintptr

// MDQuerySortComparatorFunction is callback function used to sort the results of a query.
//
// See: https://developer.apple.com/documentation/coreservices/mdquerysortcomparatorfunction
type MDQuerySortComparatorFunction = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) corefoundation.CFComparisonResult

// See: https://developer.apple.com/documentation/coreservices/mididatachunk
// MIDIDataChunk is opaque storage with the size and alignment C gives MIDIDataChunk:
// 10 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 10 into.
type MIDIDataChunk [5]uint16

// See: https://developer.apple.com/documentation/coreservices/mididatachunkptr
type MIDIDataChunkPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/mpaddressspaceid
type MPAddressSpaceID = uintptr

// See: https://developer.apple.com/documentation/coreservices/mpaddressspaceinfo
// MPAddressSpaceInfo is opaque storage with the size and alignment C gives MPAddressSpaceInfo:
// 96 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 96 into.
type MPAddressSpaceInfo [12]uint64

// See: https://developer.apple.com/documentation/coreservices/mpareaid
type MPAreaID = uintptr

// MPCoherenceID is represents a memory coherence group.
//
// See: https://developer.apple.com/documentation/coreservices/mpcoherenceid
type MPCoherenceID = uintptr

// See: https://developer.apple.com/documentation/coreservices/mpconsoleid
type MPConsoleID = uintptr

// MPCpuID is represents a CPU ID.
//
// See: https://developer.apple.com/documentation/coreservices/mpcpuid
type MPCpuID = uintptr

// MPCriticalRegionID is represents a critical region ID, which Multiprocessing Services uses to manipulate critical regions.
//
// See: https://developer.apple.com/documentation/coreservices/mpcriticalregionid
type MPCriticalRegionID = uintptr

// See: https://developer.apple.com/documentation/coreservices/mpcriticalregioninfo
// MPCriticalRegionInfo is opaque storage with the size and alignment C gives MPCriticalRegionInfo:
// 56 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 56 into.
type MPCriticalRegionInfo [7]uint64

// MPDebuggerLevel is indicates the debugger level.
//
// See: https://developer.apple.com/documentation/coreservices/mpdebuggerlevel
type MPDebuggerLevel = uint32

// MPEventFlags is represents event information for an event group.
//
// See: https://developer.apple.com/documentation/coreservices/mpeventflags
type MPEventFlags = uint32

// MPEventID is represents an event group ID, which Multiprocessing Services uses to manipulate event groups.
//
// See: https://developer.apple.com/documentation/coreservices/mpeventid
type MPEventID = uintptr

// See: https://developer.apple.com/documentation/coreservices/mpeventinfo
// MPEventInfo is opaque storage with the size and alignment C gives MPEventInfo:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type MPEventInfo [6]uint64

// MPExceptionKind is represents the kind of exception thrown.
//
// See: https://developer.apple.com/documentation/coreservices/mpexceptionkind
type MPExceptionKind = uint32

// See: https://developer.apple.com/documentation/coreservices/mpisfullyinitializedproc
type MPIsFullyInitializedProc = bool

// MPNotificationID is represents a notification ID, which Multiprocessing Services uses to manipulate kernel notifications.
//
// See: https://developer.apple.com/documentation/coreservices/mpnotificationid
type MPNotificationID = uintptr

// See: https://developer.apple.com/documentation/coreservices/mpnotificationinfo
// MPNotificationInfo is opaque storage with the size and alignment C gives MPNotificationInfo:
// 80 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 80 into.
type MPNotificationInfo [10]uint64

// MPOpaqueID is represents a generic notification ID (that is, an ID that could be a queue ID, event ID, kernel notification ID, or semaphore ID).
//
// See: https://developer.apple.com/documentation/coreservices/mpopaqueid
type MPOpaqueID = uintptr

// See: https://developer.apple.com/documentation/coreservices/mpopaqueidclass
type MPOpaqueIDClass = uint32

// See: https://developer.apple.com/documentation/coreservices/mppagesizeclass
type MPPageSizeClass = uint32

// MPProcessID is represents a process ID.
//
// See: https://developer.apple.com/documentation/coreservices/mpprocessid
type MPProcessID = uintptr

// MPQueueID is represents a queue ID, which Multiprocessing Services uses to manipulate message queues.
//
// See: https://developer.apple.com/documentation/coreservices/mpqueueid
type MPQueueID = uintptr

// See: https://developer.apple.com/documentation/coreservices/mpqueueinfo
// MPQueueInfo is opaque storage with the size and alignment C gives MPQueueInfo:
// 80 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 80 into.
type MPQueueInfo [10]uint64

// MPRemoteContext is specify which contexts are allowed to execute the callback function when using [MPRemoteCall].
//
// See: https://developer.apple.com/documentation/coreservices/mpremotecontext
type MPRemoteContext = uint8

// MPSemaphoreCount is represents a semaphore count.
//
// See: https://developer.apple.com/documentation/coreservices/mpsemaphorecount
type MPSemaphoreCount = uint

// MPSemaphoreID is represents a semaphore ID, which Multiprocessing Services uses to manipulate semaphores.
//
// See: https://developer.apple.com/documentation/coreservices/mpsemaphoreid
type MPSemaphoreID = uintptr

// See: https://developer.apple.com/documentation/coreservices/mpsemaphoreinfo
// MPSemaphoreInfo is opaque storage with the size and alignment C gives MPSemaphoreInfo:
// 56 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 56 into.
type MPSemaphoreInfo [7]uint64

// MPTaskID is represents a task ID.
//
// See: https://developer.apple.com/documentation/coreservices/mptaskid
type MPTaskID = uintptr

// MPTaskInfo is contains information about a task.
//
// See: https://developer.apple.com/documentation/coreservices/mptaskinfo
// MPTaskInfo is opaque storage with the size and alignment C gives MPTaskInfo:
// 128 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 128 into.
type MPTaskInfo [16]uint64

// See: https://developer.apple.com/documentation/coreservices/mptaskinfoversion2
// MPTaskInfoVersion2 is opaque storage with the size and alignment C gives MPTaskInfoVersion2:
// 88 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 88 into.
type MPTaskInfoVersion2 [11]uint64

// MPTaskOptions is specify optional actions when calling the [MPCreateTask] function.
//
// See: https://developer.apple.com/documentation/coreservices/mptaskoptions
type MPTaskOptions = uint32

// See: https://developer.apple.com/documentation/coreservices/mptaskstatekind
type MPTaskStateKind = uint32

// MPTaskWeight is represents the relative processor weighting of a task.
//
// See: https://developer.apple.com/documentation/coreservices/mptaskweight
type MPTaskWeight = uint32

// MPTimerID is represents a timer ID.
//
// See: https://developer.apple.com/documentation/coreservices/mptimerid
type MPTimerID = uintptr

// See: https://developer.apple.com/documentation/coreservices/machineinformation
// MachineInformation is opaque storage with the size and alignment C gives MachineInformation:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type MachineInformation [1]uint64

// See: https://developer.apple.com/documentation/coreservices/machineinformationintel64
// MachineInformationIntel64 is an unresolved C aggregate typedef.
type MachineInformationIntel64 unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/machineinformationpowerpc
// MachineInformationPowerPC is opaque storage with the size and alignment C gives MachineInformationPowerPC:
// 88 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 88 into.
type MachineInformationPowerPC [11]uint64

// See: https://developer.apple.com/documentation/coreservices/machinelocation
// MachineLocation is opaque storage with the size and alignment C gives MachineLocation:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type MachineLocation [8]uint16

// See: https://developer.apple.com/documentation/coreservices/marker
// Marker is opaque storage with the size and alignment C gives Marker:
// 262 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 262 into.
type Marker [131]uint16

// See: https://developer.apple.com/documentation/coreservices/markerchunk
// MarkerChunk is opaque storage with the size and alignment C gives MarkerChunk:
// 272 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 272 into.
type MarkerChunk [136]uint16

// See: https://developer.apple.com/documentation/coreservices/markerchunkptr
type MarkerChunkPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/markeridtype
type MarkerIdType = int16

// See: https://developer.apple.com/documentation/coreservices/memoryexceptioninformation
// MemoryExceptionInformation is opaque storage with the size and alignment C gives MemoryExceptionInformation:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type MemoryExceptionInformation [4]uint64

// See: https://developer.apple.com/documentation/coreservices/memoryreferencekind
type MemoryReferenceKind = uint

// See: https://developer.apple.com/documentation/coreservices/mixedmodestaterecord
// MixedModeStateRecord is opaque storage with the size and alignment C gives MixedModeStateRecord:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type MixedModeStateRecord [8]uint16

// See: https://developer.apple.com/documentation/coreservices/nitl4handle
type NItl4Handle = *NItl4Ptr

// See: https://developer.apple.com/documentation/coreservices/nitl4ptr
type NItl4Ptr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/nitl4rec
// NItl4Rec is opaque storage with the size and alignment C gives NItl4Rec:
// 68 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 68 into.
type NItl4Rec [34]uint16

// See: https://developer.apple.com/documentation/coreservices/nanoseconds
type Nanoseconds = uint64

// See: https://developer.apple.com/documentation/coreservices/numformatstring
// NumFormatString is opaque storage with the size and alignment C gives NumFormatString:
// 256 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 256 into.
type NumFormatString [256]byte

// See: https://developer.apple.com/documentation/coreservices/numformatstringrec
// NumFormatStringRec is opaque storage with the size and alignment C gives NumFormatStringRec:
// 256 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 256 into.
type NumFormatStringRec [256]byte

// See: https://developer.apple.com/documentation/coreservices/numberparts
// NumberParts is opaque storage with the size and alignment C gives NumberParts:
// 172 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 172 into.
type NumberParts [86]uint16

// See: https://developer.apple.com/documentation/coreservices/numberpartsptr
type NumberPartsPtr = unsafe.Pointer

// OSLAccessorProcPtr is your object accessor function either finds elements or properties of an Apple event object.
//
// See: https://developer.apple.com/documentation/coreservices/oslaccessorprocptr
type OSLAccessorProcPtr = func(desiredClass uint32, container unsafe.Pointer, containerClass uint32, form uint32, selectionData unsafe.Pointer, value unsafe.Pointer, accessorRefcon uintptr) int16

// OSLAccessorUPP is defines a data type for the universal procedure pointer for the [OSLAccessorProcPtr] callback function pointer.
//
// See: https://developer.apple.com/documentation/coreservices/oslaccessorupp
type OSLAccessorUPP = unsafe.Pointer

// OSLAdjustMarksProcPtr is defines a pointer to an adjust marks callback function. Your adjust marks function unmarks objects previously marked by a call to your marking function.
//
// See: https://developer.apple.com/documentation/coreservices/osladjustmarksprocptr
type OSLAdjustMarksProcPtr = func(newStart int32, newStop int32, markToken unsafe.Pointer) int16

// OSLAdjustMarksUPP is defines a data type for the universal procedure pointer for the [OSLAdjustMarksProcPtr] callback function pointer.
//
// See: https://developer.apple.com/documentation/coreservices/osladjustmarksupp
type OSLAdjustMarksUPP = unsafe.Pointer

// OSLCompareProcPtr is defines a pointer to an object comparison callback function. Your object comparison function compares one Apple event object to another or to the data for a descriptor.
//
// See: https://developer.apple.com/documentation/coreservices/oslcompareprocptr
type OSLCompareProcPtr = func(oper uint32, obj1 unsafe.Pointer, obj2 unsafe.Pointer, result unsafe.Pointer) int16

// OSLCompareUPP is defines a data type for the universal procedure pointer for the [OSLCompareProcPtr] callback function pointer.
//
// See: https://developer.apple.com/documentation/coreservices/oslcompareupp
type OSLCompareUPP = unsafe.Pointer

// OSLCountProcPtr is defines a pointer to an object counting callback function. Your object counting function counts the number of Apple event objects of a specified class in a specified container object.
//
// See: https://developer.apple.com/documentation/coreservices/oslcountprocptr
type OSLCountProcPtr = func(desiredType uint32, containerClass uint32, container unsafe.Pointer, result unsafe.Pointer) int16

// OSLCountUPP is defines a data type for the universal procedure pointer for the [OSLCountProcPtr] callback function pointer.
//
// See: https://developer.apple.com/documentation/coreservices/oslcountupp
type OSLCountUPP = unsafe.Pointer

// OSLDisposeTokenProcPtr is defines a pointer to a dispose token callback function. Your dispose token function, required only if you use a complex token format, disposes of the specified token.
//
// See: https://developer.apple.com/documentation/coreservices/osldisposetokenprocptr
type OSLDisposeTokenProcPtr = func(unneededToken unsafe.Pointer) int16

// OSLDisposeTokenUPP is defines a data type for the universal procedure pointer for the [OSLDisposeTokenProcPtr] callback function pointer.
//
// See: https://developer.apple.com/documentation/coreservices/osldisposetokenupp
type OSLDisposeTokenUPP = unsafe.Pointer

// OSLGetErrDescProcPtr is defines a pointer to an error descriptor callback function. Your error descriptor callback function supplies a pointer to an address where the Apple Event Manager can store the current descriptor if an error occurs during a call to the [AEResolve] function.
//
// See: https://developer.apple.com/documentation/coreservices/oslgeterrdescprocptr
type OSLGetErrDescProcPtr = func(appDescPtr unsafe.Pointer) int16

// OSLGetErrDescUPP is defines a data type for the universal procedure pointer for the [OSLGetErrDescProcPtr] callback function pointer.
//
// See: https://developer.apple.com/documentation/coreservices/oslgeterrdescupp
type OSLGetErrDescUPP = unsafe.Pointer

// OSLGetMarkTokenProcPtr is defines a pointer to a mark token callback function. Your mark token function returns a mark token.
//
// See: https://developer.apple.com/documentation/coreservices/oslgetmarktokenprocptr
type OSLGetMarkTokenProcPtr = func(dContainerToken unsafe.Pointer, containerClass uint32, result unsafe.Pointer) int16

// OSLGetMarkTokenUPP is defines a data type for the universal procedure pointer for the [OSLGetMarkTokenProcPtr] callback function pointer.
//
// See: https://developer.apple.com/documentation/coreservices/oslgetmarktokenupp
type OSLGetMarkTokenUPP = unsafe.Pointer

// OSLMarkProcPtr is defines a pointer to an object marking callback function. Your object-marking function marks a specific Apple event object.
//
// See: https://developer.apple.com/documentation/coreservices/oslmarkprocptr
type OSLMarkProcPtr = func(dToken unsafe.Pointer, markToken unsafe.Pointer, index int32) int16

// OSLMarkUPP is defines a data type for the universal procedure pointer for the [OSLMarkProcPtr] callback function pointer.
//
// See: https://developer.apple.com/documentation/coreservices/oslmarkupp
type OSLMarkUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/offpair
// OffPair is opaque storage with the size and alignment C gives OffPair:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type OffPair [2]uint16

// OffsetArrayHandle is defines a data type that points to an [OffsetArray]. Not typically used by developers.
//
// See: https://developer.apple.com/documentation/coreservices/offsetarrayhandle
type OffsetArrayHandle = *OffsetArrayPtr

// See: https://developer.apple.com/documentation/coreservices/offsetarrayptr
type OffsetArrayPtr = *OffsetArray

// See: https://developer.apple.com/documentation/coreservices/offsettable
type OffsetTable = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/pefcontainerheader
// PEFContainerHeader is opaque storage with the size and alignment C gives PEFContainerHeader:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type PEFContainerHeader [20]uint16

// See: https://developer.apple.com/documentation/coreservices/pefexportedsymbol
// PEFExportedSymbol is opaque storage with the size and alignment C gives PEFExportedSymbol:
// 10 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 10 into.
type PEFExportedSymbol [5]uint16

// See: https://developer.apple.com/documentation/coreservices/pefexportedsymbolhashslot
// PEFExportedSymbolHashSlot is opaque storage with the size and alignment C gives PEFExportedSymbolHashSlot:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type PEFExportedSymbolHashSlot [2]uint16

// See: https://developer.apple.com/documentation/coreservices/pefexportedsymbolkey
// PEFExportedSymbolKey is opaque storage with the size and alignment C gives PEFExportedSymbolKey:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type PEFExportedSymbolKey [2]uint16

// See: https://developer.apple.com/documentation/coreservices/pefimportedlibrary
// PEFImportedLibrary is opaque storage with the size and alignment C gives PEFImportedLibrary:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type PEFImportedLibrary [12]uint16

// See: https://developer.apple.com/documentation/coreservices/pefimportedsymbol
// PEFImportedSymbol is opaque storage with the size and alignment C gives PEFImportedSymbol:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type PEFImportedSymbol [2]uint16

// See: https://developer.apple.com/documentation/coreservices/pefloaderinfoheader
// PEFLoaderInfoHeader is opaque storage with the size and alignment C gives PEFLoaderInfoHeader:
// 56 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 56 into.
type PEFLoaderInfoHeader [28]uint16

// See: https://developer.apple.com/documentation/coreservices/pefloaderrelocationheader
// PEFLoaderRelocationHeader is opaque storage with the size and alignment C gives PEFLoaderRelocationHeader:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type PEFLoaderRelocationHeader [6]uint16

// See: https://developer.apple.com/documentation/coreservices/pefrelocchunk
type PEFRelocChunk = uint16

// See: https://developer.apple.com/documentation/coreservices/pefsectionheader
// PEFSectionHeader is opaque storage with the size and alignment C gives PEFSectionHeader:
// 28 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 28 into.
type PEFSectionHeader [14]uint16

// See: https://developer.apple.com/documentation/coreservices/pefsplithashword
// PEFSplitHashWord is opaque storage with the size and alignment C gives PEFSplitHashWord:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type PEFSplitHashWord [2]uint16

// See: https://developer.apple.com/documentation/coreservices/paramblockrec
// ParamBlockRec is an unresolved C aggregate typedef.
type ParamBlockRec unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/parmblkptr
type ParmBlkPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/procinfotype
type ProcInfoType = uint

// See: https://developer.apple.com/documentation/coreservices/qelem
// QElem is opaque storage with the size and alignment C gives QElem:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type QElem [6]uint16

// See: https://developer.apple.com/documentation/coreservices/qelemptr
type QElemPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/qhdr
// QHdr is opaque storage with the size and alignment C gives QHdr:
// 18 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 18 into.
type QHdr [9]uint16

// QHdrPtr is represents a type used by the Compression and Decompression API.
//
// See: https://developer.apple.com/documentation/coreservices/qhdrptr
type QHdrPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/qtypes
type QTypes = int8

// See: https://developer.apple.com/documentation/coreservices/rdflagstype
type RDFlagsType = uint8

// See: https://developer.apple.com/documentation/coreservices/rtatype
type RTAType = int8

// See: https://developer.apple.com/documentation/coreservices/registerinformation
// RegisterInformation is opaque storage with the size and alignment C gives RegisterInformation:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type RegisterInformation [1]uint64

// See: https://developer.apple.com/documentation/coreservices/registerinformationintel64
// RegisterInformationIntel64 is an unresolved C aggregate typedef.
type RegisterInformationIntel64 unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/registerinformationpowerpc
// RegisterInformationPowerPC is opaque storage with the size and alignment C gives RegisterInformationPowerPC:
// 256 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 256 into.
type RegisterInformationPowerPC [128]uint16

// See: https://developer.apple.com/documentation/coreservices/registeredcomponentinstancerecord
// RegisteredComponentInstanceRecord is opaque storage with the size and alignment C gives RegisteredComponentInstanceRecord:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type RegisteredComponentInstanceRecord [4]uint16

// See: https://developer.apple.com/documentation/coreservices/registeredcomponentinstancerecordptr
type RegisteredComponentInstanceRecordPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/registeredcomponentrecord
// RegisteredComponentRecord is opaque storage with the size and alignment C gives RegisteredComponentRecord:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type RegisteredComponentRecord [4]uint16

// See: https://developer.apple.com/documentation/coreservices/registeredcomponentrecordptr
type RegisteredComponentRecordPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/resattributes
type ResAttributes = int16

// See: https://developer.apple.com/documentation/coreservices/reserrupp
type ResErrUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/resfileattributes
type ResFileAttributes = int16

// See: https://developer.apple.com/documentation/coreservices/resfilerefnum
type ResFileRefNum = int32

// See: https://developer.apple.com/documentation/coreservices/resid
type ResID = int16

// See: https://developer.apple.com/documentation/coreservices/resourcecount
type ResourceCount = int16

// See: https://developer.apple.com/documentation/coreservices/resourceindex
type ResourceIndex = int16

// See: https://developer.apple.com/documentation/coreservices/resourcespec
// ResourceSpec is opaque storage with the size and alignment C gives ResourceSpec:
// 6 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 6 into.
type ResourceSpec [3]uint16

// See: https://developer.apple.com/documentation/coreservices/routinedescriptor
// RoutineDescriptor is opaque storage with the size and alignment C gives RoutineDescriptor:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type RoutineDescriptor [20]uint16

// See: https://developer.apple.com/documentation/coreservices/routinedescriptorhandle
type RoutineDescriptorHandle = *RoutineDescriptorPtr

// See: https://developer.apple.com/documentation/coreservices/routinedescriptorptr
type RoutineDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/routineflagstype
type RoutineFlagsType = uint16

// See: https://developer.apple.com/documentation/coreservices/routinerecord
// RoutineRecord is opaque storage with the size and alignment C gives RoutineRecord:
// 28 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 28 into.
type RoutineRecord [14]uint16

// See: https://developer.apple.com/documentation/coreservices/routinerecordhandle
type RoutineRecordHandle = *RoutineRecordPtr

// See: https://developer.apple.com/documentation/coreservices/routinerecordptr
type RoutineRecordPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/routingflags
type RoutingFlags = uint32

// See: https://developer.apple.com/documentation/coreservices/routingresourceentry
// RoutingResourceEntry is opaque storage with the size and alignment C gives RoutingResourceEntry:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type RoutingResourceEntry [10]uint16

// See: https://developer.apple.com/documentation/coreservices/routingresourcehandle
type RoutingResourceHandle = *RoutingResourcePtr

// See: https://developer.apple.com/documentation/coreservices/routingresourceptr
type RoutingResourcePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/rsrcchainlocation
type RsrcChainLocation = int16

// See: https://developer.apple.com/documentation/coreservices/rulebasedtrslrecord
// RuleBasedTrslRecord is opaque storage with the size and alignment C gives RuleBasedTrslRecord:
// 10 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 10 into.
type RuleBasedTrslRecord [5]uint16

// SKDocument is defines an opaque data type representing a document’s URL.
//
// See: https://developer.apple.com/documentation/coreservices/skdocument
type SKDocument = corefoundation.CFTypeRef

// SKDocumentID is defines an opaque data type representing a lightweight document identifier.
//
// See: https://developer.apple.com/documentation/coreservices/skdocumentid
type SKDocumentID = int

// SKDocumentRef is defines an opaque data type representing a document’s URL.
//
// See: https://developer.apple.com/documentation/coreservices/skdocumentref
type SKDocumentRef = corefoundation.CFTypeRef

// SKIndexDocumentIteratorRef is defines an opaque data type representing an index-based document iterator.
//
// See: https://developer.apple.com/documentation/coreservices/skindexdocumentiteratorref
type SKIndexDocumentIteratorRef uintptr

// SKIndexRef is defines an opaque data type representing an index.
//
// See: https://developer.apple.com/documentation/coreservices/skindexref
type SKIndexRef uintptr

// SKSearchGroupRef is deprecated. Use asynchronous searching with SKSearchCreate instead, which does not employ search groups.
//
// See: https://developer.apple.com/documentation/coreservices/sksearchgroupref
type SKSearchGroupRef uintptr

// SKSearchOptions is specifies the search options available for the [SKSearchCreate(_:_:_:)] function.
//
// See: https://developer.apple.com/documentation/coreservices/sksearchoptions
type SKSearchOptions = uint32

// SKSearchRef is defines an opaque data type representing an asynchronous search.
//
// See: https://developer.apple.com/documentation/coreservices/sksearchref
type SKSearchRef uintptr

// SKSearchResultsFilterCallBack is deprecated. Use [SKSearchCreate] and [SKSearchFindMatches] instead, which do not use a callback.
//
// See: https://developer.apple.com/documentation/coreservices/sksearchresultsfiltercallback
type SKSearchResultsFilterCallBack = func(inIndex SKIndexRef, inDocument SKDocumentRef, inContext unsafe.Pointer) bool

// SKSearchResultsRef is deprecated. Use asynchronous searching with SKSearchCreate instead, which does not employ search groups.
//
// See: https://developer.apple.com/documentation/coreservices/sksearchresultsref
type SKSearchResultsRef uintptr

// SKSummaryRef is defines an opaque data type representing summarization information.
//
// See: https://developer.apple.com/documentation/coreservices/sksummaryref
type SKSummaryRef uintptr

// See: https://developer.apple.com/documentation/coreservices/schedulerinforec
// SchedulerInfoRec is opaque storage with the size and alignment C gives SchedulerInfoRec:
// 28 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 28 into.
type SchedulerInfoRec [14]uint16

// See: https://developer.apple.com/documentation/coreservices/schedulerinforecptr
type SchedulerInfoRecPtr = unsafe.Pointer

// ScriptCodeRun is contains script code information for a text run.
//
// See: https://developer.apple.com/documentation/coreservices/scriptcoderun
// ScriptCodeRun is opaque storage with the size and alignment C gives ScriptCodeRun:
// 10 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 10 into.
type ScriptCodeRun [5]uint16

// See: https://developer.apple.com/documentation/coreservices/scriptcoderunptr
type ScriptCodeRunPtr = unsafe.Pointer

// SelectorFunctionUPP is defines a universal procedure pointer to a selector function callback.
//
// See: https://developer.apple.com/documentation/coreservices/selectorfunctionupp
type SelectorFunctionUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/sleepqrec
// SleepQRec is opaque storage with the size and alignment C gives SleepQRec:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type SleepQRec [10]uint16

// See: https://developer.apple.com/documentation/coreservices/sleepqrecptr
type SleepQRecPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/sleepqupp
type SleepQUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/sounddatachunk
// SoundDataChunk is opaque storage with the size and alignment C gives SoundDataChunk:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type SoundDataChunk [8]uint16

// See: https://developer.apple.com/documentation/coreservices/sounddatachunkptr
type SoundDataChunkPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/string2datestatus
type String2DateStatus = int16

// See: https://developer.apple.com/documentation/coreservices/stringtodatestatus
type StringToDateStatus = int16

// See: https://developer.apple.com/documentation/coreservices/syspptr
type SysPPtr = unsafe.Pointer

// TECBufferContextRec is contains buffers for text and text encoding runs.
//
// See: https://developer.apple.com/documentation/coreservices/tecbuffercontextrec
// TECBufferContextRec is opaque storage with the size and alignment C gives TECBufferContextRec:
// 64 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 64 into.
type TECBufferContextRec [8]uint64

// TECConversionInfo is contains text encoding conversion information.
//
// See: https://developer.apple.com/documentation/coreservices/tecconversioninfo
// TECConversionInfo is opaque storage with the size and alignment C gives TECConversionInfo:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type TECConversionInfo [6]uint16

// TECConverterContextRec is contains converter information used by a Text Encoding Converter plug-in.
//
// See: https://developer.apple.com/documentation/coreservices/tecconvertercontextrec
// TECConverterContextRec is opaque storage with the size and alignment C gives TECConverterContextRec:
// 152 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 152 into.
type TECConverterContextRec [19]uint64

// See: https://developer.apple.com/documentation/coreservices/tecencodingpairrec
// TECEncodingPairRec is opaque storage with the size and alignment C gives TECEncodingPairRec:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type TECEncodingPairRec [6]uint32

// See: https://developer.apple.com/documentation/coreservices/tecencodingpairs
// TECEncodingPairs is opaque storage with the size and alignment C gives TECEncodingPairs:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type TECEncodingPairs [8]uint32

// See: https://developer.apple.com/documentation/coreservices/tecencodingpairshandle
type TECEncodingPairsHandle = *TECEncodingPairsPtr

// See: https://developer.apple.com/documentation/coreservices/tecencodingpairsptr
type TECEncodingPairsPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tecencodingpairsrec
// TECEncodingPairsRec is opaque storage with the size and alignment C gives TECEncodingPairsRec:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type TECEncodingPairsRec [9]uint32

// See: https://developer.apple.com/documentation/coreservices/tecencodingslisthandle
type TECEncodingsListHandle = *TECEncodingsListPtr

// See: https://developer.apple.com/documentation/coreservices/tecencodingslistptr
type TECEncodingsListPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tecencodingslistrec
// TECEncodingsListRec is opaque storage with the size and alignment C gives TECEncodingsListRec:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type TECEncodingsListRec [4]uint32

// TECInfo is contains information about the Unicode Converter, the Text Encoding Converter, and Basic Text Types.
//
// See: https://developer.apple.com/documentation/coreservices/tecinfo
// TECInfo is opaque storage with the size and alignment C gives TECInfo:
// 84 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 84 into.
type TECInfo [42]uint16

// See: https://developer.apple.com/documentation/coreservices/tecinfohandle
type TECInfoHandle = *TECInfoPtr

// See: https://developer.apple.com/documentation/coreservices/tecinfoptr
type TECInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tecinternetnamerec
// TECInternetNameRec is opaque storage with the size and alignment C gives TECInternetNameRec:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type TECInternetNameRec [5]uint32

// See: https://developer.apple.com/documentation/coreservices/tecinternetnameusagemask
type TECInternetNameUsageMask = uint32

// See: https://developer.apple.com/documentation/coreservices/tecinternetnameshandle
type TECInternetNamesHandle = *TECInternetNamesPtr

// See: https://developer.apple.com/documentation/coreservices/tecinternetnamesptr
type TECInternetNamesPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tecinternetnamesrec
// TECInternetNamesRec is opaque storage with the size and alignment C gives TECInternetNamesRec:
// 24 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 24 into.
type TECInternetNamesRec [6]uint32

// See: https://developer.apple.com/documentation/coreservices/teclocalelisttoencodinglistptr
type TECLocaleListToEncodingListPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/teclocalelisttoencodinglistrec
// TECLocaleListToEncodingListRec is opaque storage with the size and alignment C gives TECLocaleListToEncodingListRec:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type TECLocaleListToEncodingListRec [3]uint32

// See: https://developer.apple.com/documentation/coreservices/teclocaletoencodingslisthandle
type TECLocaleToEncodingsListHandle = *TECLocaleToEncodingsListPtr

// See: https://developer.apple.com/documentation/coreservices/teclocaletoencodingslistptr
type TECLocaleToEncodingsListPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/teclocaletoencodingslistrec
// TECLocaleToEncodingsListRec is opaque storage with the size and alignment C gives TECLocaleToEncodingsListRec:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type TECLocaleToEncodingsListRec [4]uint32

// TECObjectRef is defines an opaque reference to a converter object.
//
// See: https://developer.apple.com/documentation/coreservices/tecobjectref
type TECObjectRef uintptr

// TECPluginDispatchTable is contains version and signature information and pointers to the callback functions used by a text encoding converter plug-in.
//
// See: https://developer.apple.com/documentation/coreservices/tecplugindispatchtable
// TECPluginDispatchTable is opaque storage with the size and alignment C gives TECPluginDispatchTable:
// 160 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 160 into.
type TECPluginDispatchTable [20]uint64

// TECPluginGetPluginDispatchTablePtr is defines a pointer to a function that returnsa pointer to a plug-in dispatch table.
//
// See: https://developer.apple.com/documentation/coreservices/tecplugingetplugindispatchtableptr
type TECPluginGetPluginDispatchTablePtr = unsafe.Pointer

// TECPluginSig is defines a data type for a Text Encoding Converter plug-in signature.
//
// See: https://developer.apple.com/documentation/coreservices/tecpluginsig
type TECPluginSig = uint32

// TECPluginSignature is defines a data type for a Text Encoding Converter plug-in signature.
//
// See: https://developer.apple.com/documentation/coreservices/tecpluginsignature
type TECPluginSignature = uint32

// TECPluginStateRec is contains state information for a Text Encoding Converter plug-in.
//
// See: https://developer.apple.com/documentation/coreservices/tecpluginstaterec
// TECPluginStateRec is opaque storage with the size and alignment C gives TECPluginStateRec:
// 20 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 20 into.
type TECPluginStateRec [5]uint32

// TECPluginVersion is defines a data type for Text Encoding Converter plug-in version.
//
// See: https://developer.apple.com/documentation/coreservices/tecpluginversion
type TECPluginVersion = uint32

// TECSnifferContextRec is contains infomation used by a sniffer object.
//
// See: https://developer.apple.com/documentation/coreservices/tecsniffercontextrec
// TECSnifferContextRec is opaque storage with the size and alignment C gives TECSnifferContextRec:
// 112 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 112 into.
type TECSnifferContextRec [14]uint64

// TECSnifferObjectRef is defines a reference to an opaque sniffer object.
//
// See: https://developer.apple.com/documentation/coreservices/tecsnifferobjectref
type TECSnifferObjectRef uintptr

// See: https://developer.apple.com/documentation/coreservices/tecsubtextencodingrec
// TECSubTextEncodingRec is opaque storage with the size and alignment C gives TECSubTextEncodingRec:
// 32 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 32 into.
type TECSubTextEncodingRec [8]uint32

// See: https://developer.apple.com/documentation/coreservices/tecsubtextencodingshandle
type TECSubTextEncodingsHandle = *TECSubTextEncodingsPtr

// See: https://developer.apple.com/documentation/coreservices/tecsubtextencodingsptr
type TECSubTextEncodingsPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tecsubtextencodingsrec
// TECSubTextEncodingsRec is opaque storage with the size and alignment C gives TECSubTextEncodingsRec:
// 36 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 36 into.
type TECSubTextEncodingsRec [9]uint32

// See: https://developer.apple.com/documentation/coreservices/tmtask
// TMTask is opaque storage with the size and alignment C gives TMTask:
// 42 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 42 into.
type TMTask [21]uint16

// See: https://developer.apple.com/documentation/coreservices/tmtaskptr
type TMTaskPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tabledirectoryrecord
// TableDirectoryRecord is opaque storage with the size and alignment C gives TableDirectoryRecord:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type TableDirectoryRecord [8]uint16

// TaskStorageIndex is represents a task storage index value used by functions described in “Accessing Per-Task Storage Variables.”.
//
// See: https://developer.apple.com/documentation/coreservices/taskstorageindex
type TaskStorageIndex = uint

// TaskStorageValue is represents a task storage value used by functions described in “Accessing Per-Task Storage Variables.”.
//
// See: https://developer.apple.com/documentation/coreservices/taskstoragevalue
type TaskStorageValue = unsafe.Pointer

// TextBreakLocatorRef is refers to an opaque object that encapsulates locale and text-break information for the purpose of finding boundaries in Unicode text.
//
// See: https://developer.apple.com/documentation/coreservices/textbreaklocatorref
type TextBreakLocatorRef uintptr

// See: https://developer.apple.com/documentation/coreservices/textchunk
// TextChunk is opaque storage with the size and alignment C gives TextChunk:
// 10 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 10 into.
type TextChunk [5]uint16

// See: https://developer.apple.com/documentation/coreservices/textchunkptr
type TextChunkPtr = unsafe.Pointer

// TextEncoding is defines a data type for a text encoding value.
//
// See: https://developer.apple.com/documentation/coreservices/textencoding
type TextEncoding = uint32

// TextEncodingBase is specify base text encodings.
//
// See: https://developer.apple.com/documentation/coreservices/textencodingbase
type TextEncodingBase = uint32

// TextEncodingFormat is specify a text encoding format.
//
// See: https://developer.apple.com/documentation/coreservices/textencodingformat
type TextEncodingFormat = uint32

// TextEncodingNameSelector is specify the part of an encoding name you want to obtain.
//
// See: https://developer.apple.com/documentation/coreservices/textencodingnameselector
type TextEncodingNameSelector = uint32

// See: https://developer.apple.com/documentation/coreservices/textencodingrec
// TextEncodingRec is opaque storage with the size and alignment C gives TextEncodingRec:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type TextEncodingRec [3]uint32

// TextEncodingRun is contains text encoding information for a text run.
//
// See: https://developer.apple.com/documentation/coreservices/textencodingrun
// TextEncodingRun is opaque storage with the size and alignment C gives TextEncodingRun:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type TextEncodingRun [6]uint16

// See: https://developer.apple.com/documentation/coreservices/textencodingrunptr
type TextEncodingRunPtr = unsafe.Pointer

// TextEncodingVariant is defines a data type for a text encoding variant.
//
// See: https://developer.apple.com/documentation/coreservices/textencodingvariant
type TextEncodingVariant = uint32

// See: https://developer.apple.com/documentation/coreservices/textptr
type TextPtr = *uint8

// See: https://developer.apple.com/documentation/coreservices/textrangearrayhandle
type TextRangeArrayHandle = *TextRangeArrayPtr

// See: https://developer.apple.com/documentation/coreservices/textrangearrayptr
type TextRangeArrayPtr = *TextRangeArray

// See: https://developer.apple.com/documentation/coreservices/textrangehandle
type TextRangeHandle = *TextRangePtr

// See: https://developer.apple.com/documentation/coreservices/textrangeptr
type TextRangePtr = *TextRange

// TextToUnicodeInfo is defines reference to an opaque Unicode converter object.
//
// See: https://developer.apple.com/documentation/coreservices/texttounicodeinfo
type TextToUnicodeInfo = uintptr

// See: https://developer.apple.com/documentation/coreservices/threadentrytpp
type ThreadEntryTPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/threadentryupp
type ThreadEntryUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/threadid
type ThreadID = uint

// See: https://developer.apple.com/documentation/coreservices/threadoptions
type ThreadOptions = uint32

// See: https://developer.apple.com/documentation/coreservices/threadschedulertpp
type ThreadSchedulerTPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/threadschedulerupp
type ThreadSchedulerUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/threadstate
type ThreadState = uint16

// See: https://developer.apple.com/documentation/coreservices/threadstyle
type ThreadStyle = uint32

// See: https://developer.apple.com/documentation/coreservices/threadswitchtpp
type ThreadSwitchTPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/threadswitchupp
type ThreadSwitchUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/threadtaskref
type ThreadTaskRef = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/threadterminationtpp
type ThreadTerminationTPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/threadterminationupp
type ThreadTerminationUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/timerupp
type TimerUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/togglepb
// TogglePB is opaque storage with the size and alignment C gives TogglePB:
// 48 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 48 into.
type TogglePB [24]uint16

// See: https://developer.apple.com/documentation/coreservices/toggleresults
type ToggleResults = int16

// See: https://developer.apple.com/documentation/coreservices/tripleint
type TripleInt = unsafe.Pointer

// UCCharPropertyType is specify property types for a Unicode charater.
//
// See: https://developer.apple.com/documentation/coreservices/uccharpropertytype
type UCCharPropertyType = int32

// UCCharPropertyValue is specify a propery value for a Unicode character.
//
// See: https://developer.apple.com/documentation/coreservices/uccharpropertyvalue
type UCCharPropertyValue = uint32

// UCCollateOptions is specifies options for Unicode string comparison.
//
// See: https://developer.apple.com/documentation/coreservices/uccollateoptions
type UCCollateOptions = uint32

// UCCollationValue is specifies a Unicode collation key.
//
// See: https://developer.apple.com/documentation/coreservices/uccollationvalue
type UCCollationValue = uint32

// UCKeyCharSeq is specifies the output of a dead-key state in a `'uchr'` resource.
//
// See: https://developer.apple.com/documentation/coreservices/uckeycharseq
type UCKeyCharSeq = uint16

// UCKeyOutput is specifies values in key-code-to-character tables in a `'uchr'` resource.
//
// See: https://developer.apple.com/documentation/coreservices/uckeyoutput
type UCKeyOutput = uint16

// See: https://developer.apple.com/documentation/coreservices/uctswalkdirection
type UCTSWalkDirection = uint16

// UCTextBreakOptions is specifies options for locating boundaries in Unicode text.
//
// See: https://developer.apple.com/documentation/coreservices/uctextbreakoptions
type UCTextBreakOptions = uint32

// UCTextBreakType is specifies kinds of text boundaries.
//
// See: https://developer.apple.com/documentation/coreservices/uctextbreaktype
type UCTextBreakType = uint32

// See: https://developer.apple.com/documentation/coreservices/uctypeselectcompareresult
type UCTypeSelectCompareResult = int32

// See: https://developer.apple.com/documentation/coreservices/uctypeselectoptions
type UCTypeSelectOptions = uint16

// See: https://developer.apple.com/documentation/coreservices/uctypeselectref
type UCTypeSelectRef uintptr

// See: https://developer.apple.com/documentation/coreservices/utcdatetime
// UTCDateTime is opaque storage with the size and alignment C gives UTCDateTime:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type UTCDateTime [4]uint16

// See: https://developer.apple.com/documentation/coreservices/utcdatetimehandle
type UTCDateTimeHandle = *UTCDateTimePtr

// See: https://developer.apple.com/documentation/coreservices/utcdatetimeptr
type UTCDateTimePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/unichararrayhandle
type UniCharArrayHandle = *UniCharArrayPtr

// UniCharArrayOffset is represents the boundary between two characters.
//
// See: https://developer.apple.com/documentation/coreservices/unichararrayoffset
type UniCharArrayOffset = uint

// See: https://developer.apple.com/documentation/coreservices/unichararrayptr
type UniCharArrayPtr = *uint16

// UnicodeMapVersion is specify a Unicode mapping version.
//
// See: https://developer.apple.com/documentation/coreservices/unicodemapversion
type UnicodeMapVersion = int32

// UnicodeMapping is contains information for mapping to or from Unicode encoding.
//
// See: https://developer.apple.com/documentation/coreservices/unicodemapping
// UnicodeMapping is opaque storage with the size and alignment C gives UnicodeMapping:
// 12 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 12 into.
type UnicodeMapping [6]uint16

// See: https://developer.apple.com/documentation/coreservices/unicodemappingptr
type UnicodeMappingPtr = unsafe.Pointer

// UnicodeToTextFallbackUPP is defines a universal procedure pointer to a Unicode-to-text-fallback callback function.
//
// See: https://developer.apple.com/documentation/coreservices/unicodetotextfallbackupp
type UnicodeToTextFallbackUPP = unsafe.Pointer

// UnicodeToTextInfo is defines a reference to an opaque Unicode to text converter object.
//
// See: https://developer.apple.com/documentation/coreservices/unicodetotextinfo
type UnicodeToTextInfo = uintptr

// UnicodeToTextRunInfo is defines a reference to an opaque Unicode to text run information converter object.
//
// See: https://developer.apple.com/documentation/coreservices/unicodetotextruninfo
type UnicodeToTextRunInfo = uintptr

// See: https://developer.apple.com/documentation/coreservices/untokentable
// UntokenTable is opaque storage with the size and alignment C gives UntokenTable:
// 516 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 516 into.
type UntokenTable [258]uint16

// See: https://developer.apple.com/documentation/coreservices/untokentablehandle
type UntokenTableHandle = *UntokenTablePtr

// See: https://developer.apple.com/documentation/coreservices/untokentableptr
type UntokenTablePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/vectorinformation
// VectorInformation is opaque storage with the size and alignment C gives VectorInformation:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type VectorInformation [1]uint64

// See: https://developer.apple.com/documentation/coreservices/vectorinformationintel64
// VectorInformationIntel64 is an unresolved C aggregate typedef.
type VectorInformationIntel64 unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/vectorinformationpowerpc
// VectorInformationPowerPC is opaque storage with the size and alignment C gives VectorInformationPowerPC:
// 1064 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 1064 into.
type VectorInformationPowerPC [133]uint64

// See: https://developer.apple.com/documentation/coreservices/volmountinfoheader
// VolMountInfoHeader is opaque storage with the size and alignment C gives VolMountInfoHeader:
// 6 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 6 into.
type VolMountInfoHeader [3]uint16

// See: https://developer.apple.com/documentation/coreservices/volmountinfoptr
type VolMountInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/volumemountinfoheader
// VolumeMountInfoHeader is opaque storage with the size and alignment C gives VolumeMountInfoHeader:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type VolumeMountInfoHeader [4]uint16

// See: https://developer.apple.com/documentation/coreservices/volumemountinfoheaderptr
type VolumeMountInfoHeaderPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/volumetype
type VolumeType = uint32

// WSClientContext is an optional context that can contain data you want passed to your callback.
//
// See: https://developer.apple.com/documentation/coreservices/wsclientcontext
// WSClientContext is opaque storage with the size and alignment C gives WSClientContext:
// 40 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 40 into.
type WSClientContext [20]uint16

// WSMethodInvocationRef is an opaque reference to a web services method invocation.
//
// See: https://developer.apple.com/documentation/coreservices/wsmethodinvocationref
type WSMethodInvocationRef uintptr

// WSProtocolHandlerRef is an opaque reference to a web services protocol handler.
//
// See: https://developer.apple.com/documentation/coreservices/wsprotocolhandlerref
type WSProtocolHandlerRef uintptr

// WSTypeID is web Services Core uses the following enumeration when serializing between Core Foundation and XML types. Because CFTypes are defined at runtime, it isn't always possible to produce a static mapping to a particular CFTypeRef. This enum and associated API allows for static determination of the expected serialization.
//
// See: https://developer.apple.com/documentation/coreservices/wstypeid
type WSTypeID = int32

// See: https://developer.apple.com/documentation/coreservices/widechararr
// WideCharArr is opaque storage with the size and alignment C gives WideCharArr:
// 22 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 22 into.
type WideCharArr [11]uint16

// See: https://developer.apple.com/documentation/coreservices/xlibcontainerheader
// XLibContainerHeader is opaque storage with the size and alignment C gives XLibContainerHeader:
// 80 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 80 into.
type XLibContainerHeader [40]uint16

// See: https://developer.apple.com/documentation/coreservices/xlibexportedsymbol
// XLibExportedSymbol is opaque storage with the size and alignment C gives XLibExportedSymbol:
// 8 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 8 into.
type XLibExportedSymbol [4]uint16

// See: https://developer.apple.com/documentation/coreservices/xlibexportedsymbolhashslot
// XLibExportedSymbolHashSlot is opaque storage with the size and alignment C gives XLibExportedSymbolHashSlot:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type XLibExportedSymbolHashSlot [2]uint16

// See: https://developer.apple.com/documentation/coreservices/xlibexportedsymbolkey
// XLibExportedSymbolKey is opaque storage with the size and alignment C gives XLibExportedSymbolKey:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type XLibExportedSymbolKey [2]uint16

// See: https://developer.apple.com/documentation/coreservices/ccnttokenrechandle
type CcntTokenRecHandle = *CcntTokenRecPtr

// See: https://developer.apple.com/documentation/coreservices/ccnttokenrecptr
type CcntTokenRecPtr = *CcntTokenRecord

// See: https://developer.apple.com/documentation/coreservices/decform
// Decform is opaque storage with the size and alignment C gives decform:
// 4 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 4 into.
type Decform [2]uint16

// See: https://developer.apple.com/documentation/coreservices/decimal
// Decimal is opaque storage with the size and alignment C gives decimal:
// 42 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 42 into.
type Decimal [21]uint16

// See: https://developer.apple.com/documentation/coreservices/registerselectortype
type RegisterSelectorType = uint16

// See: https://developer.apple.com/documentation/coreservices/relop
type Relop = int16

// See: https://developer.apple.com/documentation/coreservices/voidptr
type VoidPtr = unsafe.Pointer
