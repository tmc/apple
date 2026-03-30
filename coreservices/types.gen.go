// Code generated from Apple documentation for coreservices. DO NOT EDIT.

package coreservices

import (
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// C struct types

// AEBuildError - Defines a structure for storing additional error codeinformation for “AEBuild” routines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/aebuilderror
type AEBuildError struct {
	FError    uint32 // The error code. See [AEBuildErrorCode](<doc://com.apple.documentation/documentation/coreservices/aebuilderrorcode>) for alist of errors.
	FErrorPos uint32 // The character position where the parser detectedthe error.

}

// AEDesc - Stores data and an accompanying descriptor type to formthe basic building block of all Apple Events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/aedesc
type AEDesc struct {
	DescriptorType uint32         // A four-character code of type [DescType](<doc://com.apple.documentation/documentation/coreservices/desctype>) that indicates the type ofdata in the structure. See [DescType](<doc://com.apple.documentation/documentation/coreservices/desctype>).
	DataHandle     unsafe.Pointer // An opaque storage type that points to the storagefor the descriptor data. Your application doesn’t access thisdata directly—rather, it calls one of the functions [AEGetDescDataSize(_:)](<doc://com.apple.documentation/documentation/coreservices/1450119-aegetdescdatasize>), [AEGetDescData(_:_:_:)](<doc://com.apple.documentation/documentation/coreservices/1444427-aegetdescdata>), or [AEReplaceDescData(_:_:_:_:)](<doc://com.apple.documentation/documentation/coreservices/1446695-aereplacedescdata>).See [AEDataStorage](<doc://com.apple.documentation/documentation/coreservices/aedatastorage>).

}

// AEKeyDesc - Associates a keyword with a descriptor to form a keyword-specifieddescriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/aekeydesc
type AEKeyDesc struct {
	DescKey     uint32 // A four-character code of type [AEKeyword](<doc://com.apple.documentation/documentation/coreservices/aekeyword>) that uniquely identifies thekey that is associated with the data in the structure. Some keywordconstants are described in [Keyword Attribute Constants](<doc://com.apple.documentation/documentation/coreservices/apple_events/1542920-keyword_attribute_constants>) and [Keyword Parameter Constants](<doc://com.apple.documentation/documentation/coreservices/apple_events/1527206-keyword_parameter_constants>).See [AEKeyword](<doc://com.apple.documentation/documentation/coreservices/aekeyword>).
	DescContent AEDesc // A descriptor of type [AEDesc](<doc://com.apple.documentation/documentation/coreservices/aedesc>) that stores the keyword descriptordata. See [AEDesc](<doc://com.apple.documentation/documentation/coreservices/aedesc>).

}

// AERemoteProcessResolverContext - Supplied as a parameter when performing asynchronous resolutionof remote processes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/aeremoteprocessresolvercontext
type AERemoteProcessResolverContext struct {
	Version         int                                               // This should be set to zero (0).
	Info            unsafe.Pointer                                    // A pointer to arbitrary information. The pointeris retained and passed to the callback, allowing you to provideinformation to that routine.
	Retain          corefoundation.CFAllocatorRetainCallBack          // A prototype for a function callback that retainsthe specified data. Called on the info pointer. This field may be [NULL].
	Release         corefoundation.CFAllocatorReleaseCallBack         // A prototype for a function callback that releasesthe specified data. Called on the info pointer. This field may be [NULL].
	CopyDescription corefoundation.CFAllocatorCopyDescriptionCallBack // A prototype for a function callback that providesa description of the specified data. Called on the info pointer.This field may be [NULL].

}

// CSIdentityClientContext
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/csidentityclientcontext
type CSIdentityClientContext struct {
	Version         int
	Info            unsafe.Pointer
	Retain          corefoundation.CFAllocatorRetainCallBack
	Release         corefoundation.CFAllocatorReleaseCallBack
	CopyDescription corefoundation.CFAllocatorCopyDescriptionCallBack
	StatusUpdated   unsafe.Pointer
}

// CSIdentityQueryClientContext
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/csidentityqueryclientcontext
type CSIdentityQueryClientContext struct {
	Version             int
	Info                unsafe.Pointer
	RetainInfo          corefoundation.CFAllocatorRetainCallBack
	ReleaseInfo         corefoundation.CFAllocatorReleaseCallBack
	CopyInfoDescription corefoundation.CFAllocatorCopyDescriptionCallBack
	ReceiveEvent        unsafe.Pointer
}

// FSEventStreamContext
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/fseventstreamcontext
type FSEventStreamContext struct {
	Version         int                                               // Currently the only valid value is zero.
	Info            unsafe.Pointer                                    // An arbitrary client-defined value (for instance, a pointer) to be associated with the stream and passed to the callback when it is invoked. If a non-NULL value is supplied for the retain callback the framework will use it to retain this value. If a non-NULL value is supplied for the release callback then when the stream is deallocated it will be used to release this value. This can be NULL.
	Retain          corefoundation.CFAllocatorRetainCallBack          // The callback used retain the info pointer. This can be NULL.
	Release         corefoundation.CFAllocatorReleaseCallBack         // The callback used release a retain on the info pointer. This can be NULL.
	CopyDescription corefoundation.CFAllocatorCopyDescriptionCallBack // The callback used to create a descriptive string representation of the info pointer (or the data pointed to by the info pointer) for debugging purposes. This can be NULL.

}

// IntlText - International text consists of an ordered series of bytes, beginning with a 4-byte language code and a 4-byte script code that together determine the format of the bytes that follow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/intltext
type IntlText struct {
	TheScriptCode unsafe.Pointer
	TheLangCode   unsafe.Pointer
	TheText       unsafe.Pointer
}

// LSApplicationParameters - The specification that defines the app, launch flags, and additional parameters that control how an app launches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/lsapplicationparameters
type LSApplicationParameters struct {
	Version           int                            // The version of the structure. The value of thisfield must be `0`.
	Flags             LSLaunchFlags                  // Launch flags. For possible values, see [LSLaunchFlags](<doc://com.apple.documentation/documentation/coreservices/lslaunchflags>).
	Application       objectivec.IObject             // The [FSRef] ofthe application to open.
	AsyncLaunchRefCon unsafe.Pointer                 // The client `refCon` thatis to appear in subsequent launch notifications.
	Environment       corefoundation.CFDictionaryRef // A dictionary of [CFStringRef] keysand values for environment variables to set in the launched process.The value of this field can be [NULL].
	Argv              corefoundation.CFArrayRef      // An array of values of type [CFString](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFString>) that specify the arguments that are to be passed to `main()` in the launched process. The value of this field can be [NULL]. This field is ignored in macOS 10.4.
	InitialEvent      AEDesc                         // The first Apple Event to send to the launchedprocess. The value of this field can be [NULL].

}

// LSItemInfoRecord - The specification that contains requested information about an item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/lsiteminforecord
type LSItemInfoRecord struct {
	Flags     LSItemInfoFlags            // Item-information flags specifying informationabout the item; see [LSItemInfoFlags](<doc://com.apple.documentation/documentation/coreservices/lsiteminfoflags>) for a description of these flags.
	Filetype  uint32                     // The item’s file type.
	Creator   uint32                     // The item’s creator signature.
	Extension corefoundation.CFStringRef // A Core Foundation string object specifying theitem’s filename extension; see the  inthe Core Foundation Reference Documentation for a description ofthe [CFStringRef] datatype.

}

// LSLaunchFSRefSpec - The specification that defines, by file-system reference, an app to launch, items to open, or both, along with related information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/lslaunchfsrefspec
type LSLaunchFSRefSpec struct {
	AppRef         objectivec.IObject // A pointer to a file-system reference designatingthe application to launch; see the  inthe Carbon File Management Documentation for a description of the [FSRef] datatype. Set this field to [NULL] torequest that each item in the `itemRefs` arraybe opened in its own preferred application.
	NumDocs        int                // The number of elements in the array specifiedby the `itemRefs` field.The value of this field can be `0`,in which case the application designated by `appRef` islaunched without opening any items.
	ItemRefs       objectivec.IObject // An array of file-system references designatingthe item or items to open. If the value of `numDocs` is `0`,this field is ignored and can be set to [NULL].
	PassThruParams AEDesc             // A pointer to an Apple event descriptor that ispassed untouched as an optional parameter, with keyword `keyAEPropData` (`'prdt'`),in the Apple event sent to each application launched or activated(whether individual preferred applications or the application designatedby `appRef`). See the  in the Carbon Interapplication CommunicationDocumentation for a description of the [AEDesc] datatype. The value of this field can be [NULL].
	LaunchFlags    LSLaunchFlags      // Launch flags specifying how to launch each application(including whether to print or merely open documents); see [LSLaunchFlags](<doc://com.apple.documentation/documentation/coreservices/lslaunchflags>) fora description of these flags.
	AsyncRefCon    unsafe.Pointer     // A pointer to an arbitrary application-definedvalue, passed in the Carbon event notifying you of an application’slaunch or termination (if you have registered for such notification).The value of this field can be [NULL].

}

// LSLaunchURLSpec - The specification for launching an app, opening items, or both, along with related information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/lslaunchurlspec
type LSLaunchURLSpec struct {
	AppURL         corefoundation.CFURLRef   // A Core Foundation URL reference designating the application to launch; see the  in the Core Foundation Reference Documentation for a description of the [CFURLRef] data type.The URL must have scheme `file` and contain a valid path to an application file or application bundle. Set this field to [NULL] to request that each item in the `itemURLs` array be opened in its own preferred application.
	ItemURLs       corefoundation.CFArrayRef // A reference to an array of Core Foundation URL references designating the item or items to open; see the  in the Core Foundation Reference Documentation for a description of the [CFArrayRef] data type. The value of this field can be [NULL], in which case the application designated by `appURL` will be launched without opening any items.
	PassThruParams AEDesc                    // A pointer to an Apple event descriptor that is passed untouched as an optional parameter, with keyword `keyAEPropData` (`'prdt'`), in the Apple event sent to each application launched or activated (whether individual preferred applications or the application designated by `appURL`). See the  in the Carbon Interapplication Communication Documentation for a description of the [AEDesc] data type. The value of this field can be [NULL].
	LaunchFlags    LSLaunchFlags             // Launch flags specifying how to launch each application (including whether to print or merely open documents); see [LSLaunchFlags](<doc://com.apple.documentation/documentation/coreservices/lslaunchflags>) for a description of these flags.
	AsyncRefCon    unsafe.Pointer            // A pointer to an arbitrary application-defined value, passed in the Carbon event notifying you of an application’s launch or termination (if you have registered for such notification). The value of this field can be [NULL].

}

// MDExporterInterfaceStruct
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/mdexporterinterfacestruct
type MDExporterInterfaceStruct struct {
	ImporterExportData unsafe.Pointer
	Release            corefoundation.ULONG
	AddRef             corefoundation.ULONG
	QueryInterface     corefoundation.HRESULT
}

// MDImporterBundleWrapperURLInterfaceStruct
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/mdimporterbundlewrapperurlinterfacestruct
type MDImporterBundleWrapperURLInterfaceStruct struct {
	AddRef                             corefoundation.ULONG
	ImporterImportBundleWrapperURLData unsafe.Pointer
	QueryInterface                     corefoundation.HRESULT
	Release                            corefoundation.ULONG
}

// MDImporterInterfaceStruct
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/mdimporterinterfacestruct
type MDImporterInterfaceStruct struct {
	Release            corefoundation.ULONG
	AddRef             corefoundation.ULONG
	ImporterImportData unsafe.Pointer
	QueryInterface     corefoundation.HRESULT
}

// MDImporterURLInterfaceStruct
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/mdimporterurlinterfacestruct
type MDImporterURLInterfaceStruct struct {
	QueryInterface        corefoundation.HRESULT
	ImporterImportURLData unsafe.Pointer
	AddRef                corefoundation.ULONG
	Release               corefoundation.ULONG
}

// MDQueryBatchingParams - Structure containing the progress notification batchingparameters of a MDQuery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/mdquerybatchingparams
type MDQueryBatchingParams struct {
	First_max_num    int // The maximum number of results that can accumulatebefore the first progress notification is sent. This value is usedonly during the initial result-gathering phase of a query.
	First_max_ms     int // The maximum number of milliseconds that can passbefore the first progress notification is sent. This value is advisory,in that the notification will be triggered at some point after `first_max_ms` millisecondshave passed since the query began accumulating results. This valueis used only during the initial result-gathering phase of a query.
	Progress_max_num int // The maximum number of results that can accumulatebefore additional progress notifications are sent. This value isused only during the initial result-gathering phase of a query.
	Progress_max_ms  int // The maximum number of milliseconds that can passbefore additional progress notifications are sent. This value isadvisory, in that the notification will be triggered at some pointafter `progress_max_ms` millisecondshave passed since the query began accumulating results. This valueis used only during the initial result-gathering phase of a query.
	Update_max_num   int // The maximum number of results that can accumulatebefore an update notification is sent. This value is used only duringthe live-update phase of a query.
	Update_max_ms    int // The maximum number of milliseconds that can passbefore an update notification is sent. This value is advisory, inthat the notification will be triggered at some point after `update_max_ms` millisecondshave passed since the query began accumulating results. This valueis used only during the live-update phase of a query.

}

// OffsetArray - Specifies offsets of ranges of text. Not typically used by developers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/offsetarray
type OffsetArray struct {
	FNumOfOffsets unsafe.Pointer
	FOffset       unsafe.Pointer
}

// TScriptingSizeResource - Defines a data type to store stack and heap information. Not typically used by developers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/tscriptingsizeresource
type TScriptingSizeResource struct {
	ScriptingSizeFlags unsafe.Pointer
	MinStackSize       uint32
	PreferredStackSize uint32
	MaxStackSize       uint32
	MinHeapSize        uint32
	PreferredHeapSize  uint32
	MaxHeapSize        uint32
}

// TextRange - Specifies a range of text. Not typically used by developers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/textrange
type TextRange struct {
	FStart       unsafe.Pointer
	FEnd         unsafe.Pointer
	FHiliteStyle unsafe.Pointer
}

// TextRangeArray - Specifies an array of text ranges. Not typically used by developers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/textrangearray
type TextRangeArray struct {
	FNumOfRanges unsafe.Pointer
	FRange       TextRange
}

// UCKeyLayoutFeatureInfo - Specifies the longest possible output string to be produced by the current `'uchr'` resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/uckeylayoutfeatureinfo
type UCKeyLayoutFeatureInfo struct {
	KeyLayoutFeatureInfoFormat uint16 // An unsigned 16-bit integer identifying the format of the [UCKeyLayoutFeatureInfo] structure. Set to `kUCKeyLayoutFeatureInfoFormat`.
	Reserved                   uint16 // Reserved. Set to 0.
	MaxOutputStringLength      uint32 // An unsigned 32-bit integer specifying the longest possible output string of Unicode characters to be produced by this `'uchr'` resource.

}

// UCKeyModifiersToTableNum - Maps a modifier key combination to a particular key-code-to-character table number in a `'uchr'` resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/uckeymodifierstotablenum
type UCKeyModifiersToTableNum struct {
	KeyModifiersToTableNumFormat uint16 // An unsigned 16-bit integer identifying the format of the [UCKeyModifiersToTableNum] structure. Set to `kUCKeyModifiersToTableNumFormat`.
	DefaultTableNum              uint16 // An unsigned 16-bit integer identifying the table number to use for modifier combinations that are outside of the range included in the `tableNum` field.
	ModifiersCount               uint32 // An unsigned 32-bit integer specifying the range of modifier bit combinations for which there are entries in the `tableNum[]` field.
	TableNum                     uint8  // An array of unsigned 8-bit integers that map modifier bit combinations to table numbers. These values are indexes into the `keyToCharTableOffsets` array in a [UCKeyToCharTableIndex](<doc://com.apple.documentation/documentation/coreservices/uckeytochartableindex>)structure; these, in turn, are offsets to the actual key-code-to character tables, which follow the [UCKeyToCharTableIndex] structure in the `'uchr'` resource.

}

// UCKeySequenceDataIndex - Contains offsets to a list of character sequences for a `'uchr'` resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/uckeysequencedataindex
type UCKeySequenceDataIndex struct {
	KeySequenceDataIndexFormat uint16 // An unsigned 16-bit integer identifying the format of the [UCKeySequenceDataIndex] structure. Set to `kUCKeySequenceDataIndexFormat`.
	CharSequenceCount          uint16 // An unsigned 16-bit integer specifying the number of Unicode character sequences that follow the end of the [UCKeySequenceDataIndex] structure.
	CharSequenceOffsets        uint16 // An array of offsets from the beginning of the [UCKeySequenceDataIndex] structure to the Unicode character sequences that follow it. Because a given offset indicates both the beginning of a new character sequence and the end of the sequence that precedes it, the length of each sequence is determined by the difference between the offset to that sequence and the value of the next offset in the array. The array contains one more entry than the number of character sequences; the final entry is the offset to the end of the final character sequence.

}

// UCKeyStateEntryRange - Maps from a dead-key state to either the resultant Unicode character(s) or the new dead key state produced when the current state is terminated by a given character key for a `'uchr'` resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/uckeystateentryrange
type UCKeyStateEntryRange struct {
	CurStateStart   uint16 // An unsigned 16-bit integer specifying the beginning of a given dead-key state range.
	CurStateRange   uint8  // An unsigned 8-bit integer specifying the number of entries in a given dead-key state range.
	DeltaMultiplier uint8  // An unsigned 8-bit integer.
	CharData        uint16 // A value of type [UCKeyCharSeq]. This base character value is used to determine the actual Unicode character(s) produced when a given dead-key state terminates.
	NextState       uint16 // An unsigned 16-bit integer. This base dead-key state value is used to determine the following dead-key state, if any.

}

// UCKeyStateEntryTerminal - Maps from a dead-key state to the Unicode character(s) produced when that state is terminated by a given character key for a `'uchr'` resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/uckeystateentryterminal
type UCKeyStateEntryTerminal struct {
	CurState uint16 // An unsigned 16-bit integer specifying the current dead-key state.
	CharData uint16 // A value of type [UCKeyCharSeq] specifying the Unicode character(s) produced when a given character key is pressed.

}

// UCKeyStateRecord - Determines dead-key state transitions in a `'uchr'` resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/uckeystaterecord
type UCKeyStateRecord struct {
	StateZeroCharData  uint16 // A value of type [UCKeyCharSeq] specifying the Unicode character(s) produced from a given key code while no dead-key state is in effect.
	StateZeroNextState uint16 // An unsigned 16-bit integer specifying the dead-key state produced from a given key code when no previous dead-key state is in effect. If the [UCKeyStateRecord] structure does not initiate a dead-key state (but only provides terminators for other dead-key states), this will be 0. A non-zero value specifies the resulting new dead-key state and refers to the current state entry within the `stateEntryData[]` field for the following dead-key state record that is applied.
	StateEntryCount    uint16 // An unsigned 16-bit integer specifying the number of elements in the `stateEntryData` field’s array for a given dead-key state record.
	StateEntryFormat   uint16 // An unsigned 16-bit integer specifying the format of the data in the `stateEntryData` field’s array. This should be 0 if the `stateEntryCount` field is set to 0. Currently available values are `kUCKeyStateEntryTerminalFormat` and `kUCKeyStateEntryRangeFormat`; see [Key State Entry Formats](<doc://com.apple.documentation/documentation/coreservices/carbon_core/unicode_utilities/1390376-key_state_entry_formats>) for descriptions of these values.
	StateEntryData     uint32 // An array of dead-key state entries, whose size depends on their format, but which will always be a multiple of 4 bytes. Each entry maps from the current dead-key state to the Unicode character(s) that result when a given character key is pressed or to the next dead-key state, if any. The format of the entry is specified by the `stateEntryFormat` field to be either that of type [UCKeyStateEntryTerminal](<doc://com.apple.documentation/documentation/coreservices/uckeystateentryterminal>) or [UCKeyStateEntryRange](<doc://com.apple.documentation/documentation/coreservices/uckeystateentryrange>).

}

// UCKeyStateRecordsIndex - Provides a count of, and offsets to, dead-key state records in a `'uchr'` resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/uckeystaterecordsindex
type UCKeyStateRecordsIndex struct {
	KeyStateRecordsIndexFormat uint16 // An unsigned 16-bit integer identifying the format of the [UCKeyStateRecordsIndex] structure. Set to `kUCKeyStateRecordsIndexFormat`.
	KeyStateRecordCount        uint16 // An unsigned 16-bit integer specifying the number of dead-key state records that are included in the resource.
	KeyStateRecordOffsets      uint32 // An array of offsets from the beginning of the resource to each of the [UCKeyStateRecord] values that follow this structure in the `'uchr'` resource.

}

// UCKeyStateTerminators - Lists the default terminators for each dead-key state handled by a `'uchr'` resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/uckeystateterminators
type UCKeyStateTerminators struct {
	KeyStateTerminatorsFormat uint16 // An unsigned 16-bit integer identifying the format of the [UCKeyStateTerminators] structure. Set to `kUCKeyStateTerminatorsFormat`.
	KeyStateTerminatorCount   uint16 // An unsigned 16-bit integer specifying the number of default dead-key state terminators contained in the `keyStateTerminators[]` array.
	KeyStateTerminators       uint16 // An array of default dead-key state terminators, described as values of type [UCKeyCharSeq](<doc://com.apple.documentation/documentation/coreservices/uckeycharseq>); the value `keyStateTerminators[0]` is the terminator for state 1, and so on.

}

// UCKeyToCharTableIndex - Provides a count of, and offsets to, key-code-to-character tables in a `'uchr'` resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/uckeytochartableindex
type UCKeyToCharTableIndex struct {
	KeyToCharTableIndexFormat uint16 // An unsigned 16-bit integer identifying the format of the [UCKeyToCharTableIndex] structure. Set to `kUCKeyToCharTableIndexFormat`.
	KeyToCharTableSize        uint16 // An unsigned 16-bit integer specifying the number of virtual key codes supported by this resource; for ADB keyboards this is 128 (with virtual key codes ranging from 0 to 127).
	KeyToCharTableCount       uint32 // An unsigned 32-bit integer specifying the number of key-code-to-character tables, typically 6 to 12.
	KeyToCharTableOffsets     uint32 // An array of offsets from the beginning of the `'uchr'` resource to each of the [UCKeyOutput] key-code-to-character tables in the `keyToCharData[]` array that follows this structure in the resource.

}

// UCKeyboardLayout - Provides header data for a `'uchr'` resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/uckeyboardlayout
type UCKeyboardLayout struct {
	KeyLayoutHeaderFormat      uint16               // An unsigned 16-bit integer identifying the format of the structure. Set to `kUCLayoutHeaderFormat`.
	KeyLayoutDataVersion       uint16               // An unsigned 16-bit integer identifying the version of the data in the resource, in binary code decimal format. For example, 0x0100 would equal version 1.0.
	KeyLayoutFeatureInfoOffset uint32               // An unsigned 32-bit integer providing an offset to a structure of type [UCKeyLayoutFeatureInfo](<doc://com.apple.documentation/documentation/coreservices/uckeylayoutfeatureinfo>), if such is used in the resource. May be 0 if no [UCKeyLayoutFeatureInfo] table is included in the resource.
	KeyboardTypeCount          uint32               // An unsigned 32-bit integer specifying the number of [UCKeyboardTypeHeader] structures in the `keyboardTypeList[]` field’s array.
	KeyboardTypeList           UCKeyboardTypeHeader // A variable-length array containing structures of type [UCKeyboardTypeHeader]. Each [UCKeyboardTypeHeader] entry specifies a range of physical keyboard types and contains offsets to each of the key mapping sections to be used for that range of keyboard types.

}

// UCKeyboardTypeHeader - Specifies a range of physical keyboard types in a `'uchr'` resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/uckeyboardtypeheader
type UCKeyboardTypeHeader struct {
	KeyboardTypeFirst            uint32 // An unsigned 32-bit integer specifying the first keyboard type in this entry. For the initial entry (that is, the default entry) in an array of [UCKeyboardTypeHeader] structures, you should set this value to 0. The initial [UCKeyboardTypeHeader] entry is used if the keyboard type passed to the function [UCKeyTranslate(_:_:_:_:_:_:_:_:_:_:)](<doc://com.apple.documentation/documentation/coreservices/1390584-uckeytranslate>) does not match any other entry, that is, if it is not within the range of values specified by `keyboardTypeFirst` and `keyboardTypeLast` for any entry.
	KeyboardTypeLast             uint32 // An unsigned 32-bit integer specifying the last keyboard type in this entry. For the initial entry (that is, the default entry) in an array of [UCKeyboardTypeHeader] structures, you should set this value to 0.
	KeyModifiersToTableNumOffset uint32 // An unsigned 32-bit integer providing an offset to a structure of type [UCKeyModifiersToTableNum](<doc://com.apple.documentation/documentation/coreservices/uckeymodifierstotablenum>). The `'uchr'` resource requires a [UCKeyModifiersToTableNum] structure, therefore this field must contain a non-zero value.
	KeyToCharTableIndexOffset    uint32 // An unsigned 32-bit integer providing an offset to a structure of type [UCKeyToCharTableIndex](<doc://com.apple.documentation/documentation/coreservices/uckeytochartableindex>). The `'uchr'` resource requires a [UCKeyToCharTableIndex] structure, therefore this field must contain a non-zero value.
	KeyStateRecordsIndexOffset   uint32 // An unsigned 32-bit integer providing an offset to a structure of type [UCKeyStateRecordsIndex](<doc://com.apple.documentation/documentation/coreservices/uckeystaterecordsindex>), if such is used in the resource. This value may be 0 if no dead-key state records are included in the resource.
	KeyStateTerminatorsOffset    uint32 // An unsigned 32-bit integer providing an offset to a structure of type [UCKeyStateTerminators](<doc://com.apple.documentation/documentation/coreservices/uckeystateterminators>), if such is used in the resource. This value may be 0 if no dead-key state terminators are included in the resource.
	KeySequenceDataIndexOffset   uint32 // An unsigned 32-bit integer providing an offset to a structure of type [UCKeySequenceDataIndex](<doc://com.apple.documentation/documentation/coreservices/uckeysequencedataindex>), if such is used in the resource. This value may be 0 if no character key sequences are included in the resource.

}

// WritingCode
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/writingcode
type WritingCode struct {
	TheScriptCode unsafe.Pointer
	TheLangCode   unsafe.Pointer
}

// CcntTokenRecord - Stores token information used by the AEResolve functionwhile locating a range of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/ccnttokenrecord
type CcntTokenRecord struct {
	TokenClass uint32 // The class ID of the container represented by the `token` parameter.See [DescType](<doc://com.apple.documentation/documentation/coreservices/desctype>).
	Token      AEDesc // A token for the current container. (Token is definedin [AEDisposeToken(_:)](<doc://com.apple.documentation/documentation/coreservices/1446783-aedisposetoken>). See [AEDesc](<doc://com.apple.documentation/documentation/coreservices/aedesc>).

}
