// Code generated from Apple documentation. DO NOT EDIT.

package coreservices

import (
	"unsafe"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/security"
)

// AEAddressDesc is a descriptor that contains the address of an application, used to describe the target application for an Apple event.
//
// See: https://developer.apple.com/documentation/coreservices/aeaddressdesc
type AEAddressDesc = AEDesc

// AEArrayData is stores array information to be put into a descriptor listwith the [AEPutArray] functionor extracted from a descriptor list with the [AEGetArray] function.
//
// See: https://developer.apple.com/documentation/coreservices/1443170-aearraydata
type AEArrayData = unsafe.Pointer

// AEArrayDataPointer is a pointer to a union of type [AEArrayData].
//
// See: https://developer.apple.com/documentation/coreservices/aearraydatapointer
type AEArrayDataPointer = *AEArrayData

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
type AECoerceDescProcPtr = func(objectivec.IObject, uint32, uintptr, objectivec.IObject) int16

// AECoerceDescUPP is defines a data type for the universal procedure pointer for the [AECoerceDescProcPtr] callback function pointer.
//
// See: https://developer.apple.com/documentation/coreservices/aecoercedescupp
type AECoerceDescUPP = unsafe.Pointer

// AECoercePtrProcPtr is defines a pointer to a function that coerces data stored in a buffer. Your pointer coercion callback routine coerces the data from the passed buffer to the specified type, returning the coerced data in a descriptor.
//
// See: https://developer.apple.com/documentation/coreservices/aecoerceptrprocptr
type AECoercePtrProcPtr = func(uint32, unsafe.Pointer, corefoundation.CGSize, uint32, uintptr, objectivec.IObject) int16

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
type AEDisposeExternalProcPtr = func(unsafe.Pointer, corefoundation.CGSize, uintptr)

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
type AEEventHandlerProcPtr = func(objectivec.IObject, objectivec.IObject, uintptr) int16

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
type AERecord = unsafe.Pointer

// AERemoteProcessResolverCallback is defines a pointer to a function the Apple Event Manager calls when the asynchronous execution of a remote process resolver completes, either due to success or failure, after a call to the [AERemoteProcessResolverScheduleWithRunLoop] function. Your callback function can use the reference passed to it to get the remote process information.
//
// See: https://developer.apple.com/documentation/coreservices/aeremoteprocessresolvercallback
type AERemoteProcessResolverCallback = func(AERemoteProcessResolverRef, unsafe.Pointer)

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
type AFPAlternateAddress = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/afpserversignature
type AFPServerSignature = uint8

// See: https://developer.apple.com/documentation/coreservices/afptagdata
type AFPTagData = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/afpvolmountinfo
type AFPVolMountInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/afpvolmountinfoptr
type AFPVolMountInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/afpxvolmountinfo
type AFPXVolMountInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/afpxvolmountinfoptr
type AFPXVolMountInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/aiffloop
type AIFFLoop = unsafe.Pointer

// AliasHandle is abst_AliasHandle.
//
// See: https://developer.apple.com/documentation/coreservices/aliashandle
type AliasHandle = unsafe.Pointer

// AliasInfoType is defines the alias record information type used in the index parameter of [GetAliasInfo].
//
// See: https://developer.apple.com/documentation/coreservices/aliasinfotype
type AliasInfoType = int16

// AliasPtr is abst_AliasPtr.
//
// See: https://developer.apple.com/documentation/coreservices/aliasptr
type AliasPtr = objectivec.IObject

// AliasRecord is defines an alias record.
//
// See: https://developer.apple.com/documentation/coreservices/aliasrecord
type AliasRecord = unsafe.Pointer

// AppleEvent is a descriptor whose data is a list of descriptors containing both attributes and parameters that make up an Apple event.
//
// See: https://developer.apple.com/documentation/coreservices/appleevent
type AppleEvent = AEDesc

// See: https://developer.apple.com/documentation/coreservices/appleeventptr
type AppleEventPtr = *AEDesc

// See: https://developer.apple.com/documentation/coreservices/applicationspecificchunk
type ApplicationSpecificChunk = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/applicationspecificchunkptr
type ApplicationSpecificChunkPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/areaid
type AreaID = uintptr

// See: https://developer.apple.com/documentation/coreservices/audiorecordingchunk
type AudioRecordingChunk = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/audiorecordingchunkptr
type AudioRecordingChunkPtr = unsafe.Pointer

// BigEndianFixed is protects a big-endian Fixed value from being changed bylittle-endian code.
//
// See: https://developer.apple.com/documentation/coreservices/bigendianfixed
type BigEndianFixed = unsafe.Pointer

// BigEndianLong is protects a big-endian long value from being changed bylittle-endian code.
//
// See: https://developer.apple.com/documentation/coreservices/bigendianlong
type BigEndianLong = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/bigendianostype
type BigEndianOSType = unsafe.Pointer

// BigEndianShort is protects a big-endian short value from being changed bylittle-endian code.
//
// See: https://developer.apple.com/documentation/coreservices/bigendianshort
type BigEndianShort = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/bigendianuint32
type BigEndianUInt32 = unsafe.Pointer

// BigEndianUnsignedFixed is protects a big-endian unsigned Fixed value from beingchanged by little-endian code.
//
// See: https://developer.apple.com/documentation/coreservices/bigendianunsignedfixed
type BigEndianUnsignedFixed = unsafe.Pointer

// BigEndianUnsignedLong is protects a big-endian unsigned long value from being changedby little-endian code.
//
// See: https://developer.apple.com/documentation/coreservices/bigendianunsignedlong
type BigEndianUnsignedLong = unsafe.Pointer

// BigEndianUnsignedShort is protects a big-endian unsigned short value from beingchanged by little-endian code.
//
// See: https://developer.apple.com/documentation/coreservices/bigendianunsignedshort
type BigEndianUnsignedShort = unsafe.Pointer

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
type CSIdentityStatusUpdatedCallback = func(CSIdentityRef, int, corefoundation.CFErrorRef, unsafe.Pointer)

// See: https://developer.apple.com/documentation/coreservices/callingconventiontype
type CallingConventionType = uint16

// See: https://developer.apple.com/documentation/coreservices/catpositionrec
type CatPositionRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/chunkheader
type ChunkHeader = unsafe.Pointer

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
type Comment = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/commentschunk
type CommentsChunk = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/commentschunkptr
type CommentsChunkPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/commonchunk
type CommonChunk = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/commonchunkptr
type CommonChunkPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/component
type Component = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/componentaliasresource
type ComponentAliasResource = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/componentdescription
type ComponentDescription = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/componentfunctionupp
type ComponentFunctionUPP = unsafe.Pointer

// ComponentInstance is abst_ComponentInstance.
//
// See: https://developer.apple.com/documentation/coreservices/componentinstance
type ComponentInstance = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/componentinstancerecord
type ComponentInstanceRecord = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/componentmpworkfunctionheaderrecord
type ComponentMPWorkFunctionHeaderRecord = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/componentmpworkfunctionheaderrecordptr
type ComponentMPWorkFunctionHeaderRecordPtr = unsafe.Pointer

// ComponentMPWorkFunctionUPP is represents a type used by the Image Codec API.
//
// See: https://developer.apple.com/documentation/coreservices/componentmpworkfunctionupp
type ComponentMPWorkFunctionUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/componentparameters
type ComponentParameters = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/componentplatforminfo
type ComponentPlatformInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/componentplatforminfoarray
type ComponentPlatformInfoArray = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/componentrecord
type ComponentRecord = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/componentresource
type ComponentResource = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/componentresourceextension
type ComponentResourceExtension = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/componentresourcehandle
type ComponentResourceHandle = unsafe.Pointer

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
type ConstFSSpecPtr = objectivec.IObject

// ConstScriptCodeRunPtr is defines a constant script code run pointer.
//
// See: https://developer.apple.com/documentation/coreservices/constscriptcoderunptr
type ConstScriptCodeRunPtr = ScriptCodeRun

// ConstTextEncodingRunPtr is defines a constant text encoding run pointer.
//
// See: https://developer.apple.com/documentation/coreservices/consttextencodingrunptr
type ConstTextEncodingRunPtr = TextEncodingRun

// ConstTextPtr is defines a constant text pointer.
//
// See: https://developer.apple.com/documentation/coreservices/consttextptr
type ConstTextPtr = uint8

// ConstTextToUnicodeInfo is defines a constant text to Unicode converter object.
//
// See: https://developer.apple.com/documentation/coreservices/consttexttounicodeinfo
type ConstTextToUnicodeInfo = TextToUnicodeInfo

// ConstUniCharArrayPtr is defines a constant Unicode character array pointer.
//
// See: https://developer.apple.com/documentation/coreservices/constunichararrayptr
type ConstUniCharArrayPtr = uint16

// ConstUnicodeMappingPtr is defines a constant Unicode mapping pointer.
//
// See: https://developer.apple.com/documentation/coreservices/constunicodemappingptr
type ConstUnicodeMappingPtr = UnicodeMapping

// ConstUnicodeToTextInfo is defines a constant Unicode to text converter object.
//
// See: https://developer.apple.com/documentation/coreservices/constunicodetotextinfo
type ConstUnicodeToTextInfo = UnicodeToTextInfo

// See: https://developer.apple.com/documentation/coreservices/containerchunk
type ContainerChunk = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/custombadgeresource
type CustomBadgeResource = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/custombadgeresourcehandle
type CustomBadgeResourceHandle = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/custombadgeresourceptr
type CustomBadgeResourcePtr = unsafe.Pointer

// DCSDictionaryRef is an opaque object that represents a dictionary file.
//
// See: https://developer.apple.com/documentation/coreservices/dcsdictionaryref
type DCSDictionaryRef uintptr

// See: https://developer.apple.com/documentation/coreservices/dinfo
type DInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/dxinfo
type DXInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/datecacheptr
type DateCachePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/datecacherecord
type DateCacheRecord = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/datedelta
type DateDelta = int8

// See: https://developer.apple.com/documentation/coreservices/dateform
type DateForm = int8

// See: https://developer.apple.com/documentation/coreservices/dateorders
type DateOrders = int8

// See: https://developer.apple.com/documentation/coreservices/datetimerec
type DateTimeRec = unsafe.Pointer

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
type DeferredTask = unsafe.Pointer

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
type ExceptionInformation = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/exceptioninformationpowerpc
type ExceptionInformationPowerPC = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/exceptionkind
type ExceptionKind = uint

// See: https://developer.apple.com/documentation/coreservices/extcommonchunk
type ExtCommonChunk = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/extcommonchunkptr
type ExtCommonChunkPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/extcomponentresource
type ExtComponentResource = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/extcomponentresourcehandle
type ExtComponentResourceHandle = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/extcomponentresourceptr
type ExtComponentResourcePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/extendedfileinfo
type ExtendedFileInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/extendedfolderinfo
type ExtendedFolderInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/finfo
type FInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fnmessage
type FNMessage = uint32

// See: https://developer.apple.com/documentation/coreservices/fnsubscriptionref
type FNSubscriptionRef uintptr

// See: https://developer.apple.com/documentation/coreservices/fnsubscriptionupp
type FNSubscriptionUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fpregintel
type FPRegIntel = uint8

// See: https://developer.apple.com/documentation/coreservices/fpuinformation
type FPUInformation = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fpuinformationintel64
type FPUInformationIntel64 = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fpuinformationpowerpc
type FPUInformationPowerPC = unsafe.Pointer

// FSAliasInfo is defines an information block passed to the [FSCopyAliasInfo] function.
//
// See: https://developer.apple.com/documentation/coreservices/fsaliasinfo
type FSAliasInfo = unsafe.Pointer

// FSAliasInfoBitmap is returned by the [FSCopyAliasInfo] function to indicate which fields of the alias information structure contain valid data.
//
// See: https://developer.apple.com/documentation/coreservices/fsaliasinfobitmap
type FSAliasInfoBitmap = uint32

// See: https://developer.apple.com/documentation/coreservices/fsaliasinfoptr
type FSAliasInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsallocationflags
type FSAllocationFlags = uint16

// See: https://developer.apple.com/documentation/coreservices/fscatalogbulkparam
type FSCatalogBulkParam = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fscatalogbulkparamptr
type FSCatalogBulkParamPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fscataloginfo
type FSCatalogInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fscataloginfobitmap
type FSCatalogInfoBitmap = uint32

// See: https://developer.apple.com/documentation/coreservices/fscataloginfoptr
type FSCatalogInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsejectstatus
type FSEjectStatus = uint32

// See: https://developer.apple.com/documentation/coreservices/fseventstreamcallback
type FSEventStreamCallback = func(ConstFSEventStreamRef, unsafe.Pointer, int, unsafe.Pointer, objectivec.IObject, objectivec.IObject)

// See: https://developer.apple.com/documentation/coreservices/fseventstreamcreateflags
type FSEventStreamCreateFlags = uint32

// See: https://developer.apple.com/documentation/coreservices/fseventstreameventflags
type FSEventStreamEventFlags = uint32

// See: https://developer.apple.com/documentation/coreservices/fseventstreameventid
type FSEventStreamEventId = uint64

// See: https://developer.apple.com/documentation/coreservices/fseventstreamref
type FSEventStreamRef uintptr

// See: https://developer.apple.com/documentation/coreservices/fsfileoperationclientcontext
type FSFileOperationClientContext = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsfileoperationref
type FSFileOperationRef uintptr

// See: https://developer.apple.com/documentation/coreservices/fsfileoperationstage
type FSFileOperationStage = uint32

// See: https://developer.apple.com/documentation/coreservices/fsfilesecurityref
type FSFileSecurityRef uintptr

// See: https://developer.apple.com/documentation/coreservices/fsforkcbinfoparam
type FSForkCBInfoParam = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsforkcbinfoparamptr
type FSForkCBInfoParamPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsforkioparam
type FSForkIOParam = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsforkioparamptr
type FSForkIOParamPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsforkinfo
type FSForkInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsforkinfoflags
type FSForkInfoFlags = uint8

// See: https://developer.apple.com/documentation/coreservices/fsforkinfoptr
type FSForkInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsiorefnum
type FSIORefNum = int

// See: https://developer.apple.com/documentation/coreservices/fsiterator
type FSIterator = uintptr

// See: https://developer.apple.com/documentation/coreservices/fsiteratorflags
type FSIteratorFlags = objectivec.IObject

// See: https://developer.apple.com/documentation/coreservices/fsmountstatus
type FSMountStatus = uint32

// See: https://developer.apple.com/documentation/coreservices/fspermissioninfo
type FSPermissionInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsrangelockparam
type FSRangeLockParam = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsrangelockparamptr
type FSRangeLockParamPtr = unsafe.Pointer

// FSRef is identifies a directory or file, including a volume’s root directory.
//
// See: https://developer.apple.com/documentation/coreservices/fsref
type FSRef uintptr

// See: https://developer.apple.com/documentation/coreservices/fsrefforkioparam
type FSRefForkIOParam = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsrefforkioparamptr
type FSRefForkIOParamPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsrefparam
type FSRefParam = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsrefparamptr
type FSRefParamPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsrefptr
type FSRefPtr = objectivec.IObject

// See: https://developer.apple.com/documentation/coreservices/fssearchparams
type FSSearchParams = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fssearchparamsptr
type FSSearchParamsPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsspec
type FSSpec = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsspecarrayptr
type FSSpecArrayPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsspechandle
type FSSpecHandle = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsspecptr
type FSSpecPtr = objectivec.IObject

// See: https://developer.apple.com/documentation/coreservices/fsunmountstatus
type FSUnmountStatus = uint32

// See: https://developer.apple.com/documentation/coreservices/fsvolumeejectupp
type FSVolumeEjectUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsvolumeinfo
type FSVolumeInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fsvolumeinfobitmap
type FSVolumeInfoBitmap = uint32

// See: https://developer.apple.com/documentation/coreservices/fsvolumeinfoparam
type FSVolumeInfoParam = unsafe.Pointer

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
type FVector = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fxinfo
type FXInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/fileinfo
type FileInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/folderclass
type FolderClass = uint32

// See: https://developer.apple.com/documentation/coreservices/folderdesc
type FolderDesc = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/folderdescflags
type FolderDescFlags = uint32

// See: https://developer.apple.com/documentation/coreservices/folderdescptr
type FolderDescPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/folderinfo
type FolderInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/folderlocation
type FolderLocation = uint32

// See: https://developer.apple.com/documentation/coreservices/foldermanagernotificationupp
type FolderManagerNotificationUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/folderrouting
type FolderRouting = unsafe.Pointer

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
type FormatVersionChunk = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/formatversionchunkptr
type FormatVersionChunkPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/getmissingcomponentresourceupp
type GetMissingComponentResourceUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/getvolparmsinfobuffer
type GetVolParmsInfoBuffer = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/hfscatalognodeid
type HFSCatalogNodeID = uint32

// See: https://developer.apple.com/documentation/coreservices/iocompletionupp
type IOCompletionUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/isatype
type ISAType = int8

// See: https://developer.apple.com/documentation/coreservices/iconfamilyelement
type IconFamilyElement = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/iconfamilyhandle
type IconFamilyHandle = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/iconfamilyptr
type IconFamilyPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/iconfamilyresource
type IconFamilyResource = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/iconref
type IconRef uintptr

// See: https://developer.apple.com/documentation/coreservices/iconservicesusageflags
type IconServicesUsageFlags = uint32

// See: https://developer.apple.com/documentation/coreservices/indextoucstringprocptr
type IndexToUCStringProcPtr = func(uint32, unsafe.Pointer, unsafe.Pointer, objectivec.IObject, objectivec.IObject) objectivec.IObject

// See: https://developer.apple.com/documentation/coreservices/indextoucstringupp
type IndexToUCStringUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/instrumentchunk
type InstrumentChunk = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/instrumentchunkptr
type InstrumentChunkPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/intl0hndl
type Intl0Hndl = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/intl0ptr
type Intl0Ptr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/intl0rec
type Intl0Rec = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/intl1hndl
type Intl1Hndl = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/intl1ptr
type Intl1Ptr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/intl1rec
type Intl1Rec = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/itl1extrec
type Itl1ExtRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/itl4handle
type Itl4Handle = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/itl4ptr
type Itl4Ptr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/itl4rec
type Itl4Rec = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/itl5record
type Itl5Record = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/itlbextrecord
type ItlbExtRecord = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/itlbrecord
type ItlbRecord = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/itlcrecord
type ItlcRecord = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/kcattrtype
type KCAttrType = security.SecKeychainAttrType

// See: https://developer.apple.com/documentation/coreservices/kcattribute
type KCAttribute = security.SecKeychainAttribute

// See: https://developer.apple.com/documentation/coreservices/kcattributelist
type KCAttributeList = security.SecKeychainAttributeList

// See: https://developer.apple.com/documentation/coreservices/kcauthtype
type KCAuthType = uint32

// See: https://developer.apple.com/documentation/coreservices/kccallbackinfo
type KCCallbackInfo = unsafe.Pointer

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
type KCStatus = security.SecKeychainStatus

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
type LocalDateTime = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/localdatetimehandle
type LocalDateTimeHandle = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/localdatetimeptr
type LocalDateTimePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/localeandvariant
type LocaleAndVariant = unsafe.Pointer

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
type MDQueryCreateResultFunction = func(MDQueryRef, MDItemRef, unsafe.Pointer) unsafe.Pointer

// MDQueryCreateValueFunction is callback function usedto create the value objects stored and returned by a query.
//
// See: https://developer.apple.com/documentation/coreservices/mdquerycreatevaluefunction
type MDQueryCreateValueFunction = func(MDQueryRef, corefoundation.CFStringRef, corefoundation.CFTypeRef, unsafe.Pointer) unsafe.Pointer

// MDQueryRef is a reference to a MDQuery object.
//
// See: https://developer.apple.com/documentation/coreservices/mdqueryref
type MDQueryRef uintptr

// MDQuerySortComparatorFunction is callback function used to sort the results of a query.
//
// See: https://developer.apple.com/documentation/coreservices/mdquerysortcomparatorfunction
type MDQuerySortComparatorFunction = func(objectivec.IObject, objectivec.IObject, unsafe.Pointer) corefoundation.CFComparisonResult

// See: https://developer.apple.com/documentation/coreservices/mididatachunk
type MIDIDataChunk = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/mididatachunkptr
type MIDIDataChunkPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/mpaddressspaceid
type MPAddressSpaceID = uintptr

// See: https://developer.apple.com/documentation/coreservices/mpaddressspaceinfo
type MPAddressSpaceInfo = unsafe.Pointer

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
type MPCriticalRegionInfo = unsafe.Pointer

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
type MPEventInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/mpisfullyinitializedproc
type MPIsFullyInitializedProc = bool

// MPNotificationID is represents a notification ID, which Multiprocessing Services uses to manipulate kernel notifications.
//
// See: https://developer.apple.com/documentation/coreservices/mpnotificationid
type MPNotificationID = uintptr

// See: https://developer.apple.com/documentation/coreservices/mpnotificationinfo
type MPNotificationInfo = unsafe.Pointer

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
type MPQueueInfo = unsafe.Pointer

// MPRemoteContext is specify which contexts are allowed to execute the callback function when using [MPRemoteCall].
//
// See: https://developer.apple.com/documentation/coreservices/mpremotecontext
type MPRemoteContext = uint8

// MPSemaphoreCount is represents a semaphore count.
//
// See: https://developer.apple.com/documentation/coreservices/mpsemaphorecount
type MPSemaphoreCount = objectivec.IObject

// MPSemaphoreID is represents a semaphore ID, which Multiprocessing Services uses to manipulate semaphores.
//
// See: https://developer.apple.com/documentation/coreservices/mpsemaphoreid
type MPSemaphoreID = uintptr

// See: https://developer.apple.com/documentation/coreservices/mpsemaphoreinfo
type MPSemaphoreInfo = unsafe.Pointer

// MPTaskID is represents a task ID.
//
// See: https://developer.apple.com/documentation/coreservices/mptaskid
type MPTaskID = uintptr

// MPTaskInfo is contains information about a task.
//
// See: https://developer.apple.com/documentation/coreservices/mptaskinfo
type MPTaskInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/mptaskinfoversion2
type MPTaskInfoVersion2 = unsafe.Pointer

// MPTaskOptions is specify optional actions when calling the [MPCreateTask] function.
//
// See: https://developer.apple.com/documentation/coreservices/mptaskoptions
type MPTaskOptions = objectivec.IObject

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
type MachineInformation = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/machineinformationintel64
type MachineInformationIntel64 = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/machineinformationpowerpc
type MachineInformationPowerPC = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/machinelocation
type MachineLocation = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/marker
type Marker = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/markerchunk
type MarkerChunk = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/markerchunkptr
type MarkerChunkPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/markeridtype
type MarkerIdType = int16

// See: https://developer.apple.com/documentation/coreservices/memoryexceptioninformation
type MemoryExceptionInformation = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/memoryreferencekind
type MemoryReferenceKind = uint

// See: https://developer.apple.com/documentation/coreservices/mixedmodestaterecord
type MixedModeStateRecord = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/nitl4handle
type NItl4Handle = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/nitl4ptr
type NItl4Ptr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/nitl4rec
type NItl4Rec = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/nanoseconds
type Nanoseconds = uint64

// See: https://developer.apple.com/documentation/coreservices/numformatstring
type NumFormatString = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/numformatstringrec
type NumFormatStringRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/numberparts
type NumberParts = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/numberpartsptr
type NumberPartsPtr = unsafe.Pointer

// OSLAccessorProcPtr is your object accessor function either finds elements or properties of an Apple event object.
//
// See: https://developer.apple.com/documentation/coreservices/oslaccessorprocptr
type OSLAccessorProcPtr = func(uint32, objectivec.IObject, uint32, uint32, objectivec.IObject, objectivec.IObject, uintptr) int16

// OSLAccessorUPP is defines a data type for the universal procedure pointer for the [OSLAccessorProcPtr] callback function pointer.
//
// See: https://developer.apple.com/documentation/coreservices/oslaccessorupp
type OSLAccessorUPP = unsafe.Pointer

// OSLAdjustMarksProcPtr is defines a pointer to an adjust marks callback function. Your adjust marks function unmarks objects previously marked by a call to your marking function.
//
// See: https://developer.apple.com/documentation/coreservices/osladjustmarksprocptr
type OSLAdjustMarksProcPtr = func(int, int, objectivec.IObject) int16

// OSLAdjustMarksUPP is defines a data type for the universal procedure pointer for the [OSLAdjustMarksProcPtr] callback function pointer.
//
// See: https://developer.apple.com/documentation/coreservices/osladjustmarksupp
type OSLAdjustMarksUPP = unsafe.Pointer

// OSLCompareProcPtr is defines a pointer to an object comparison callback function. Your object comparison function compares one Apple event object to another or to the data for a descriptor.
//
// See: https://developer.apple.com/documentation/coreservices/oslcompareprocptr
type OSLCompareProcPtr = func(uint32, objectivec.IObject, objectivec.IObject, objectivec.IObject) int16

// OSLCompareUPP is defines a data type for the universal procedure pointer for the [OSLCompareProcPtr] callback function pointer.
//
// See: https://developer.apple.com/documentation/coreservices/oslcompareupp
type OSLCompareUPP = unsafe.Pointer

// OSLCountProcPtr is defines a pointer to an object counting callback function. Your object counting function counts the number of Apple event objects of a specified class in a specified container object.
//
// See: https://developer.apple.com/documentation/coreservices/oslcountprocptr
type OSLCountProcPtr = func(uint32, uint32, objectivec.IObject, objectivec.IObject) int16

// OSLCountUPP is defines a data type for the universal procedure pointer for the [OSLCountProcPtr] callback function pointer.
//
// See: https://developer.apple.com/documentation/coreservices/oslcountupp
type OSLCountUPP = unsafe.Pointer

// OSLDisposeTokenProcPtr is defines a pointer to a dispose token callback function. Your dispose token function, required only if you use a complex token format, disposes of the specified token.
//
// See: https://developer.apple.com/documentation/coreservices/osldisposetokenprocptr
type OSLDisposeTokenProcPtr = func(objectivec.IObject) int16

// OSLDisposeTokenUPP is defines a data type for the universal procedure pointer for the [OSLDisposeTokenProcPtr] callback function pointer.
//
// See: https://developer.apple.com/documentation/coreservices/osldisposetokenupp
type OSLDisposeTokenUPP = unsafe.Pointer

// OSLGetErrDescProcPtr is defines a pointer to an error descriptor callback function. Your error descriptor callback function supplies a pointer to an address where the Apple Event Manager can store the current descriptor if an error occurs during a call to the [AEResolve] function.
//
// See: https://developer.apple.com/documentation/coreservices/oslgeterrdescprocptr
type OSLGetErrDescProcPtr = func(objectivec.IObject) int16

// OSLGetErrDescUPP is defines a data type for the universal procedure pointer for the [OSLGetErrDescProcPtr] callback function pointer.
//
// See: https://developer.apple.com/documentation/coreservices/oslgeterrdescupp
type OSLGetErrDescUPP = unsafe.Pointer

// OSLGetMarkTokenProcPtr is defines a pointer to a mark token callback function. Your mark token function returns a mark token.
//
// See: https://developer.apple.com/documentation/coreservices/oslgetmarktokenprocptr
type OSLGetMarkTokenProcPtr = func(objectivec.IObject, uint32, objectivec.IObject) int16

// OSLGetMarkTokenUPP is defines a data type for the universal procedure pointer for the [OSLGetMarkTokenProcPtr] callback function pointer.
//
// See: https://developer.apple.com/documentation/coreservices/oslgetmarktokenupp
type OSLGetMarkTokenUPP = unsafe.Pointer

// OSLMarkProcPtr is defines a pointer to an object marking callback function. Your object-marking function marks a specific Apple event object.
//
// See: https://developer.apple.com/documentation/coreservices/oslmarkprocptr
type OSLMarkProcPtr = func(objectivec.IObject, objectivec.IObject, int) int16

// OSLMarkUPP is defines a data type for the universal procedure pointer for the [OSLMarkProcPtr] callback function pointer.
//
// See: https://developer.apple.com/documentation/coreservices/oslmarkupp
type OSLMarkUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/offpair
type OffPair = unsafe.Pointer

// OffsetArrayHandle is defines a data type that points to an [OffsetArray]. Not typically used by developers.
//
// See: https://developer.apple.com/documentation/coreservices/offsetarrayhandle
type OffsetArrayHandle = *OffsetArrayPtr

// See: https://developer.apple.com/documentation/coreservices/offsetarrayptr
type OffsetArrayPtr = *OffsetArray

// See: https://developer.apple.com/documentation/coreservices/offsettable
type OffsetTable = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/pefcontainerheader
type PEFContainerHeader = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/pefexportedsymbol
type PEFExportedSymbol = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/pefexportedsymbolhashslot
type PEFExportedSymbolHashSlot = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/pefexportedsymbolkey
type PEFExportedSymbolKey = string

// See: https://developer.apple.com/documentation/coreservices/pefimportedlibrary
type PEFImportedLibrary = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/pefimportedsymbol
type PEFImportedSymbol = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/pefloaderinfoheader
type PEFLoaderInfoHeader = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/pefloaderrelocationheader
type PEFLoaderRelocationHeader = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/pefrelocchunk
type PEFRelocChunk = uint16

// See: https://developer.apple.com/documentation/coreservices/pefsectionheader
type PEFSectionHeader = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/pefsplithashword
type PEFSplitHashWord = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/paramblockrec
type ParamBlockRec = objectivec.IObject

// See: https://developer.apple.com/documentation/coreservices/parmblkptr
type ParmBlkPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/procinfotype
type ProcInfoType = uint

// See: https://developer.apple.com/documentation/coreservices/qelem
type QElem = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/qelemptr
type QElemPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/qhdr
type QHdr = unsafe.Pointer

// QHdrPtr is represents a type used by the Compression and Decompression API.
//
// See: https://developer.apple.com/documentation/coreservices/qhdrptr
type QHdrPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/qtypes
type QTypes = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/rdflagstype
type RDFlagsType = uint8

// See: https://developer.apple.com/documentation/coreservices/rtatype
type RTAType = int8

// See: https://developer.apple.com/documentation/coreservices/registerinformation
type RegisterInformation = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/registerinformationintel64
type RegisterInformationIntel64 = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/registerinformationpowerpc
type RegisterInformationPowerPC = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/registeredcomponentinstancerecord
type RegisteredComponentInstanceRecord = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/registeredcomponentinstancerecordptr
type RegisteredComponentInstanceRecordPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/registeredcomponentrecord
type RegisteredComponentRecord = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/registeredcomponentrecordptr
type RegisteredComponentRecordPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/resattributes
type ResAttributes = int16

// See: https://developer.apple.com/documentation/coreservices/reserrupp
type ResErrUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/resfileattributes
type ResFileAttributes = int16

// See: https://developer.apple.com/documentation/coreservices/resfilerefnum
type ResFileRefNum = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/resid
type ResID = int16

// See: https://developer.apple.com/documentation/coreservices/resourcecount
type ResourceCount = int16

// See: https://developer.apple.com/documentation/coreservices/resourceindex
type ResourceIndex = int16

// See: https://developer.apple.com/documentation/coreservices/resourcespec
type ResourceSpec = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/routinedescriptor
type RoutineDescriptor = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/routinedescriptorhandle
type RoutineDescriptorHandle = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/routinedescriptorptr
type RoutineDescriptorPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/routineflagstype
type RoutineFlagsType = uint16

// See: https://developer.apple.com/documentation/coreservices/routinerecord
type RoutineRecord = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/routinerecordhandle
type RoutineRecordHandle = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/routinerecordptr
type RoutineRecordPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/routingflags
type RoutingFlags = uint32

// See: https://developer.apple.com/documentation/coreservices/routingresourceentry
type RoutingResourceEntry = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/routingresourcehandle
type RoutingResourceHandle = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/routingresourceptr
type RoutingResourcePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/rsrcchainlocation
type RsrcChainLocation = int16

// See: https://developer.apple.com/documentation/coreservices/rulebasedtrslrecord
type RuleBasedTrslRecord = unsafe.Pointer

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
type SKSearchResultsFilterCallBack = func(SKIndexRef, SKDocumentRef, unsafe.Pointer) objectivec.IObject

// SKSearchResultsRef is deprecated. Use asynchronous searching with SKSearchCreate instead, which does not employ search groups.
//
// See: https://developer.apple.com/documentation/coreservices/sksearchresultsref
type SKSearchResultsRef uintptr

// SKSummaryRef is defines an opaque data type representing summarization information.
//
// See: https://developer.apple.com/documentation/coreservices/sksummaryref
type SKSummaryRef uintptr

// See: https://developer.apple.com/documentation/coreservices/schedulerinforec
type SchedulerInfoRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/schedulerinforecptr
type SchedulerInfoRecPtr = unsafe.Pointer

// ScriptCodeRun is contains script code information for a text run.
//
// See: https://developer.apple.com/documentation/coreservices/scriptcoderun
type ScriptCodeRun = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/scriptcoderunptr
type ScriptCodeRunPtr = unsafe.Pointer

// SelectorFunctionUPP is defines a universal procedure pointer to a selector function callback.
//
// See: https://developer.apple.com/documentation/coreservices/selectorfunctionupp
type SelectorFunctionUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/sleepqrec
type SleepQRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/sleepqrecptr
type SleepQRecPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/sleepqupp
type SleepQUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/sounddatachunk
type SoundDataChunk = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/sounddatachunkptr
type SoundDataChunkPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/string2datestatus
type String2DateStatus = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/stringtodatestatus
type StringToDateStatus = int16

// See: https://developer.apple.com/documentation/coreservices/syspptr
type SysPPtr = unsafe.Pointer

// TECBufferContextRec is contains buffers for text and text encoding runs.
//
// See: https://developer.apple.com/documentation/coreservices/tecbuffercontextrec
type TECBufferContextRec = unsafe.Pointer

// TECConversionInfo is contains text encoding conversion information.
//
// See: https://developer.apple.com/documentation/coreservices/tecconversioninfo
type TECConversionInfo = unsafe.Pointer

// TECConverterContextRec is contains converter information used by a Text Encoding Converter plug-in.
//
// See: https://developer.apple.com/documentation/coreservices/tecconvertercontextrec
type TECConverterContextRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tecencodingpairrec
type TECEncodingPairRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tecencodingpairs
type TECEncodingPairs = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tecencodingpairshandle
type TECEncodingPairsHandle = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tecencodingpairsptr
type TECEncodingPairsPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tecencodingpairsrec
type TECEncodingPairsRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tecencodingslisthandle
type TECEncodingsListHandle = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tecencodingslistptr
type TECEncodingsListPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tecencodingslistrec
type TECEncodingsListRec = unsafe.Pointer

// TECInfo is contains information about the Unicode Converter, the Text Encoding Converter, and Basic Text Types.
//
// See: https://developer.apple.com/documentation/coreservices/tecinfo
type TECInfo = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tecinfohandle
type TECInfoHandle = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tecinfoptr
type TECInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tecinternetnamerec
type TECInternetNameRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tecinternetnameusagemask
type TECInternetNameUsageMask = uint32

// See: https://developer.apple.com/documentation/coreservices/tecinternetnameshandle
type TECInternetNamesHandle = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tecinternetnamesptr
type TECInternetNamesPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tecinternetnamesrec
type TECInternetNamesRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/teclocalelisttoencodinglistptr
type TECLocaleListToEncodingListPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/teclocalelisttoencodinglistrec
type TECLocaleListToEncodingListRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/teclocaletoencodingslisthandle
type TECLocaleToEncodingsListHandle = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/teclocaletoencodingslistptr
type TECLocaleToEncodingsListPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/teclocaletoencodingslistrec
type TECLocaleToEncodingsListRec = unsafe.Pointer

// TECObjectRef is defines an opaque reference to a converter object.
//
// See: https://developer.apple.com/documentation/coreservices/tecobjectref
type TECObjectRef uintptr

// TECPluginDispatchTable is contains version and signature information and pointers to the callback functions used by a text encoding converter plug-in.
//
// See: https://developer.apple.com/documentation/coreservices/tecplugindispatchtable
type TECPluginDispatchTable = unsafe.Pointer

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
type TECPluginStateRec = unsafe.Pointer

// TECPluginVersion is defines a data type for Text Encoding Converter plug-in version.
//
// See: https://developer.apple.com/documentation/coreservices/tecpluginversion
type TECPluginVersion = uint32

// TECSnifferContextRec is contains infomation used by a sniffer object.
//
// See: https://developer.apple.com/documentation/coreservices/tecsniffercontextrec
type TECSnifferContextRec = unsafe.Pointer

// TECSnifferObjectRef is defines a reference to an opaque sniffer object.
//
// See: https://developer.apple.com/documentation/coreservices/tecsnifferobjectref
type TECSnifferObjectRef uintptr

// See: https://developer.apple.com/documentation/coreservices/tecsubtextencodingrec
type TECSubTextEncodingRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tecsubtextencodingshandle
type TECSubTextEncodingsHandle = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tecsubtextencodingsptr
type TECSubTextEncodingsPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tecsubtextencodingsrec
type TECSubTextEncodingsRec = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tmtask
type TMTask = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tmtaskptr
type TMTaskPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/tabledirectoryrecord
type TableDirectoryRecord = unsafe.Pointer

// TaskStorageIndex is represents a task storage index value used by functions described in “Accessing Per-Task Storage Variables.”.
//
// See: https://developer.apple.com/documentation/coreservices/taskstorageindex
type TaskStorageIndex = objectivec.IObject

// TaskStorageValue is represents a task storage value used by functions described in “Accessing Per-Task Storage Variables.”.
//
// See: https://developer.apple.com/documentation/coreservices/taskstoragevalue
type TaskStorageValue = unsafe.Pointer

// TextBreakLocatorRef is refers to an opaque object that encapsulates locale and text-break information for the purpose of finding boundaries in Unicode text.
//
// See: https://developer.apple.com/documentation/coreservices/textbreaklocatorref
type TextBreakLocatorRef uintptr

// See: https://developer.apple.com/documentation/coreservices/textchunk
type TextChunk = unsafe.Pointer

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
type TextEncodingRec = unsafe.Pointer

// TextEncodingRun is contains text encoding information for a text run.
//
// See: https://developer.apple.com/documentation/coreservices/textencodingrun
type TextEncodingRun = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/textencodingrunptr
type TextEncodingRunPtr = unsafe.Pointer

// TextEncodingVariant is defines a data type for a text encoding variant.
//
// See: https://developer.apple.com/documentation/coreservices/textencodingvariant
type TextEncodingVariant = uint32

// See: https://developer.apple.com/documentation/coreservices/textptr
type TextPtr = uint8

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
type ThreadTaskRef uintptr

// See: https://developer.apple.com/documentation/coreservices/threadterminationtpp
type ThreadTerminationTPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/threadterminationupp
type ThreadTerminationUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/timerupp
type TimerUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/togglepb
type TogglePB = unsafe.Pointer

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
type UTCDateTime = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/utcdatetimehandle
type UTCDateTimeHandle = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/utcdatetimeptr
type UTCDateTimePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/unichararrayhandle
type UniCharArrayHandle = unsafe.Pointer

// UniCharArrayOffset is represents the boundary between two characters.
//
// See: https://developer.apple.com/documentation/coreservices/unichararrayoffset
type UniCharArrayOffset = uint

// See: https://developer.apple.com/documentation/coreservices/unichararrayptr
type UniCharArrayPtr = uint16

// UnicodeMapVersion is specify a Unicode mapping version.
//
// See: https://developer.apple.com/documentation/coreservices/unicodemapversion
type UnicodeMapVersion = int32

// UnicodeMapping is contains information for mapping to or from Unicode encoding.
//
// See: https://developer.apple.com/documentation/coreservices/unicodemapping
type UnicodeMapping = unsafe.Pointer

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
type UntokenTable = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/untokentablehandle
type UntokenTableHandle = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/untokentableptr
type UntokenTablePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/vectorinformation
type VectorInformation = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/vectorinformationintel64
type VectorInformationIntel64 = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/vectorinformationpowerpc
type VectorInformationPowerPC = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/volmountinfoheader
type VolMountInfoHeader = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/volmountinfoptr
type VolMountInfoPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/volumemountinfoheader
type VolumeMountInfoHeader = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/volumemountinfoheaderptr
type VolumeMountInfoHeaderPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/volumetype
type VolumeType = uint32

// WSClientContext is an optional context that can contain data you want passed to your callback.
//
// See: https://developer.apple.com/documentation/coreservices/wsclientcontext
type WSClientContext = unsafe.Pointer

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
type WSTypeID = string

// See: https://developer.apple.com/documentation/coreservices/widechararr
type WideCharArr = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/xlibcontainerheader
type XLibContainerHeader = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/xlibexportedsymbol
type XLibExportedSymbol = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/xlibexportedsymbolhashslot
type XLibExportedSymbolHashSlot = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/xlibexportedsymbolkey
type XLibExportedSymbolKey = string

// See: https://developer.apple.com/documentation/coreservices/ccnttokenrechandle
type CcntTokenRecHandle = *CcntTokenRecPtr

// See: https://developer.apple.com/documentation/coreservices/ccnttokenrecptr
type CcntTokenRecPtr = *CcntTokenRecord

// See: https://developer.apple.com/documentation/coreservices/decform
type Decform = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/decimal
type Decimal = unsafe.Pointer

// See: https://developer.apple.com/documentation/coreservices/registerselectortype
type RegisterSelectorType = uint16

// See: https://developer.apple.com/documentation/coreservices/relop
type Relop = int16

// See: https://developer.apple.com/documentation/coreservices/voidptr
type VoidPtr = unsafe.Pointer

