// Code generated from Apple documentation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

// NSPointerFunctionsOptions values.
const (

	// HashTableCopyIn is equal to [copyIn].
	//
	// See: https://developer.apple.com/documentation/Foundation/NSHashTableCopyIn
	HashTableCopyIn NSPointerFunctionsOptions = 65536

	// HashTableObjectPointerPersonality is equal to [objectPointerPersonality].
	//
	// See: https://developer.apple.com/documentation/Foundation/NSHashTableObjectPointerPersonality
	HashTableObjectPointerPersonality NSPointerFunctionsOptions = 512

	// HashTableStrongMemory is equal to [strongMemory].
	//
	// See: https://developer.apple.com/documentation/Foundation/NSHashTableStrongMemory
	HashTableStrongMemory NSPointerFunctionsOptions = 0

	// HashTableWeakMemory is equal to [weakMemory]. Uses weak read and write barriers appropriate for ARC or GC. Using [weakMemory] object references will turn to [NULL] on last release.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSHashTableWeakMemory
	HashTableWeakMemory NSPointerFunctionsOptions = 5

	// MapTableCopyIn is equivalent to [copyIn].
	//
	// See: https://developer.apple.com/documentation/Foundation/NSMapTableCopyIn
	MapTableCopyIn NSPointerFunctionsOptions = 65536

	// MapTableObjectPointerPersonality is equivalent to [objectPointerPersonality].
	//
	// See: https://developer.apple.com/documentation/Foundation/NSMapTableObjectPointerPersonality
	MapTableObjectPointerPersonality NSPointerFunctionsOptions = 512

	// MapTableStrongMemory is equivalent to [strongMemory].
	//
	// See: https://developer.apple.com/documentation/Foundation/NSMapTableStrongMemory
	MapTableStrongMemory NSPointerFunctionsOptions = 0

	// MapTableWeakMemory is equivalent to [weakMemory].
	//
	// See: https://developer.apple.com/documentation/Foundation/NSMapTableWeakMemory
	MapTableWeakMemory NSPointerFunctionsOptions = 5
)

// int values.
const (

	// NotFound is a value indicating that a requested item couldn’t be found or doesn’t exist.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotFound-4qp9h
	NotFound int = 9223372036854775807

	// OperationQueueDefaultMaxConcurrentOperationCount is the default maximum number of operations to invoke concurrently in a queue.
	//
	// See: https://developer.apple.com/documentation/Foundation/OperationQueue/defaultMaxConcurrentOperationCount
	OperationQueueDefaultMaxConcurrentOperationCount int = -1
)

// uint values.
const (

	// UndoCloseGroupingRunLoopOrdering is a priority to use when using a run loop to close an undo group.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSUndoCloseGroupingRunLoopOrdering
	UndoCloseGroupingRunLoopOrdering uint = 350000
)

var (
	// AlternateDescriptionAttributeName is an alternate description for a URL or image.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/alternateDescription
	AlternateDescriptionAttributeName NSAttributedStringKey
	// ImageURLAttributeName is the URL for an image in Markdown text.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/imageURL
	ImageURLAttributeName NSAttributedStringKey
	// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/agreeWithArgument
	InflectionAgreementArgumentAttributeName NSAttributedStringKey
	// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/agreeWithConcept
	InflectionAgreementConceptAttributeName NSAttributedStringKey
	// InflectionAlternativeAttributeName is the alternative translation for a string when no suitable inflection exists.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/inflectionAlternative
	InflectionAlternativeAttributeName NSAttributedStringKey
	// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/referentConcept
	InflectionReferentConceptAttributeName NSAttributedStringKey
	// InflectionRuleAttributeName is an attribute that tells the system how to apply grammar rules and other modifiers to the range of text.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/inflectionRule
	InflectionRuleAttributeName NSAttributedStringKey
	// InlinePresentationIntentAttributeName is an attribute that provides details for an inline Markdown element.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/inlinePresentationIntent
	InlinePresentationIntentAttributeName NSAttributedStringKey
	// LanguageIdentifierAttributeName is the language identifier associated with the range of text.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/languageIdentifier
	LanguageIdentifierAttributeName NSAttributedStringKey
	// ListItemDelimiterAttributeName is the delimiter used when declaring the current list item.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/listItemDelimiter
	ListItemDelimiterAttributeName NSAttributedStringKey
	// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/localizedNumberFormat
	LocalizedNumberFormatAttributeName NSAttributedStringKey
	// MarkdownSourcePositionAttributeName is the position in a Markdown source string corresponding to some attributed text.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/markdownSourcePosition
	MarkdownSourcePositionAttributeName NSAttributedStringKey
	// MorphologyAttributeName is an attribute that contains grammatical properties to apply to the text.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/morphology
	MorphologyAttributeName NSAttributedStringKey
	// PresentationIntentAttributeName is an attribute that provides details for a block-level Markdown element.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/presentationIntentAttributeName
	PresentationIntentAttributeName NSAttributedStringKey
	// ReplacementIndexAttributeName is the replacement position associated with a format string specifier.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/replacementIndex
	ReplacementIndexAttributeName NSAttributedStringKey
)

var (
	// AppleEventManagerWillProcessFirstEventNotification is posted by [NSAppleEventManager] before it first dispatches an Apple event. Your application can use this notification to avoid registering any Apple event handlers until the first time at which they may be needed.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSAppleEventManagerWillProcessFirstEvent
	AppleEventManagerWillProcessFirstEventNotification NSNotificationName
	// BundleDidLoadNotification is a notification that lets observers know when classes are dynamically loaded.
	//
	// See: https://developer.apple.com/documentation/Foundation/Bundle/didLoadNotification
	BundleDidLoadNotification NSNotificationName
	// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsbundleresourcerequestlowdiskspace
	NSBundleResourceRequestLowDiskSpaceNotification NSNotificationName
	// CalendarDayChangedNotification is a notification that is posted whenever the calendar day of the system changes, as determined by the system calendar, locale, and time zone.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSCalendarDayChanged
	CalendarDayChangedNotification NSNotificationName
	// ClassDescriptionNeededForClassNotification is posted by [init(for:)] when a class description cannot be found for a class.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSClassDescriptionNeededForClass
	ClassDescriptionNeededForClassNotification NSNotificationName
	// CurrentLocaleDidChangeNotification is a notification that indicates that the user’s locale changed.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocale/currentLocaleDidChangeNotification
	CurrentLocaleDidChangeNotification NSNotificationName
	// DidBecomeSingleThreadedNotification is not implemented.
	//
	// Deprecated: Deprecated since macOS 26.0. Programs no longer transition to single-threaded mode from threaded environments
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSDidBecomeSingleThreaded
	DidBecomeSingleThreadedNotification NSNotificationName
	// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsextensionhostdidbecomeactive
	NSExtensionHostDidBecomeActiveNotification NSNotificationName
	// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsextensionhostdidenterbackground
	NSExtensionHostDidEnterBackgroundNotification NSNotificationName
	// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsextensionhostwillenterforeground
	NSExtensionHostWillEnterForegroundNotification NSNotificationName
	// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsextensionhostwillresignactive
	NSExtensionHostWillResignActiveNotification NSNotificationName
	// FileHandleConnectionAcceptedNotification is posted when a file handle object establishes a socket connection between two processes, creates a file handle object for one end of the connection, and makes this object available to observers.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSFileHandleConnectionAccepted
	FileHandleConnectionAcceptedNotification NSNotificationName
	// FileHandleDataAvailableNotification is posted when the file handle determines that data is currently available for reading in a file or at a communications channel.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSFileHandleDataAvailable
	FileHandleDataAvailableNotification NSNotificationName
	// FileHandleReadCompletionNotification is posted when the file handle reads the data currently available in a file or at a communications channel.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileHandle/readCompletionNotification
	FileHandleReadCompletionNotification NSNotificationName
	// FileHandleReadToEndOfFileCompletionNotification is posted when the file handle reads all data in the file or, in a communications channel, until the other process signals the end of data.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSFileHandleReadToEndOfFileCompletion
	FileHandleReadToEndOfFileCompletionNotification NSNotificationName
	// HTTPCookieManagerCookiesChangedNotification is a notification posted when the cookies stored in the cookie storage have changed.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSHTTPCookieManagerCookiesChanged
	HTTPCookieManagerCookiesChangedNotification NSNotificationName
	// MetadataQueryDidFinishGatheringNotification is posted when the receiver has finished with the initial result-gathering phase of the query.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSMetadataQueryDidFinishGathering
	MetadataQueryDidFinishGatheringNotification NSNotificationName
	// MetadataQueryDidStartGatheringNotification is posted when the receiver begins with the initial result-gathering phase of the query.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSMetadataQueryDidStartGathering
	MetadataQueryDidStartGatheringNotification NSNotificationName
	// MetadataQueryDidUpdateNotification is posted when the receiver’s results have changed during the live-update phase of the query.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSMetadataQueryDidUpdate
	MetadataQueryDidUpdateNotification NSNotificationName
	// MetadataQueryGatheringProgressNotification is posted as the receiver is collecting results during the initial result-gathering phase of the query.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSMetadataQueryGatheringProgress
	MetadataQueryGatheringProgressNotification NSNotificationName
	// PortDidBecomeInvalidNotification is posted from the [invalidate()] method, which is invoked when the [NSPort] is deallocated or when it notices that its communication channel has been damaged. The notification object is the [NSPort] object that has become invalid. This notification does not contain a `userInfo` dictionary.
	//
	// See: https://developer.apple.com/documentation/Foundation/Port/didBecomeInvalidNotification
	PortDidBecomeInvalidNotification NSNotificationName
	// ProcessInfoPowerStateDidChangeNotification is posts when the power state of a device changes.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSProcessInfoPowerStateDidChange
	ProcessInfoPowerStateDidChangeNotification NSNotificationName
	// ProcessInfoThermalStateDidChangeNotification is posts when the thermal state of the system changes.
	//
	// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/thermalStateDidChangeNotification
	ProcessInfoThermalStateDidChangeNotification NSNotificationName
	// SystemClockDidChangeNotification is a notification posted whenever the system clock is changed.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSSystemClockDidChange
	SystemClockDidChangeNotification NSNotificationName
	// SystemTimeZoneDidChangeNotification is a notification posted when the time zone changes.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSSystemTimeZoneDidChange
	SystemTimeZoneDidChangeNotification NSNotificationName
	// TaskDidTerminateNotification is posted when the task has stopped execution.
	//
	// See: https://developer.apple.com/documentation/Foundation/Process/didTerminateNotification
	TaskDidTerminateNotification NSNotificationName
	// ThreadWillExitNotification is an [NSThread] object posts this notification when it receives the [exit()] message, before the thread exits. Observer methods invoked to receive this notification execute in the exiting thread, before it exits.
	//
	// Deprecated: Deprecated since macOS 26.0. This notification does not protect against data races
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSThreadWillExit
	ThreadWillExitNotification NSNotificationName
	// UbiquitousKeyValueStoreDidChangeExternallyNotification is posted when the value of one or more keys changes due to incoming data from iCloud.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/didChangeExternallyNotification
	UbiquitousKeyValueStoreDidChangeExternallyNotification NSNotificationName
	// UbiquityIdentityDidChangeNotification is sent after the iCloud (“ubiquity”) identity has changed.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSUbiquityIdentityDidChange
	UbiquityIdentityDidChangeNotification NSNotificationName
	// UndoManagerCheckpointNotification is posted whenever an undo manager opens or closes an undo group (except when it opens a top-level group) and when checking the redo stack.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSUndoManagerCheckpoint
	UndoManagerCheckpointNotification NSNotificationName
	// UndoManagerDidCloseUndoGroupNotification is posted after an undo manager closes an undo group.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSUndoManagerDidCloseUndoGroup
	UndoManagerDidCloseUndoGroupNotification NSNotificationName
	// UndoManagerDidOpenUndoGroupNotification is posted whenever an undo manager opens an undo group.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSUndoManagerDidOpenUndoGroup
	UndoManagerDidOpenUndoGroupNotification NSNotificationName
	// UndoManagerDidRedoChangeNotification is posted just after an undo manager performs a redo operation.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSUndoManagerDidRedoChange
	UndoManagerDidRedoChangeNotification NSNotificationName
	// UndoManagerDidUndoChangeNotification is posted just after an undo manager performs an undo operation.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSUndoManagerDidUndoChange
	UndoManagerDidUndoChangeNotification NSNotificationName
	// UndoManagerWillCloseUndoGroupNotification is posted before an undo manager closes an undo group.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSUndoManagerWillCloseUndoGroup
	UndoManagerWillCloseUndoGroupNotification NSNotificationName
	// UndoManagerWillRedoChangeNotification is posted just before an undo manager performs a redo operation.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSUndoManagerWillRedoChange
	UndoManagerWillRedoChangeNotification NSNotificationName
	// UndoManagerWillUndoChangeNotification is posted just before an undo manager performs an undo operation.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSUndoManagerWillUndoChange
	UndoManagerWillUndoChangeNotification NSNotificationName
	// UserDefaultsDidChangeNotification is posted when the current process changes the value of a setting.
	//
	// See: https://developer.apple.com/documentation/Foundation/UserDefaults/didChangeNotification
	UserDefaultsDidChangeNotification NSNotificationName
	// WillBecomeMultiThreadedNotification is posted when the first thread is detached from the current thread. The [NSThread] class posts this notification at most once—the first time a thread is detached using [detachNewThreadSelector(_:toTarget:with:)] or the [start()] method. Subsequent invocations of those methods do not post this notification. Observers of this notification have their notification method invoked in the main thread, not the new thread. The observer notification methods always execute before the new thread begins executing.
	//
	// Deprecated: Deprecated since macOS 26.0. This notification does not protect against data races
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSWillBecomeMultiThreaded
	WillBecomeMultiThreadedNotification NSNotificationName
)

var (
	// AppleEventTimeOutDefault is specifies that an event-processing operation should continue until a timeout occurs based on a value determined by the Apple Event Manager (about 1 minute). Not currently used by applications.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSAppleEventTimeOutDefault
	AppleEventTimeOutDefault float64
	// AppleEventTimeOutNone is specifies that the application is willing to wait indefinitely for the current operation to complete. Not currently used by applications.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSAppleEventTimeOutNone
	AppleEventTimeOutNone float64
	// FoundationVersionNumber is the version of the Foundation framework in the current environment.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSFoundationVersionNumber
	FoundationVersionNumber float64
)

var (
	// AppleScriptErrorAppName is an [NSString] that specifies the name of the application that generated the error.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSAppleScript/errorAppName
	AppleScriptErrorAppName string
	// AppleScriptErrorBriefMessage is an [NSString] that provides a brief description of the error.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSAppleScript/errorBriefMessage
	AppleScriptErrorBriefMessage string
	// AppleScriptErrorMessage is an [NSString] that supplies a detailed description of the error condition.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSAppleScript/errorMessage
	AppleScriptErrorMessage string
	// AppleScriptErrorNumber is an [NSNumber] that specifies the error number.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSAppleScript/errorNumber
	AppleScriptErrorNumber string
	// AppleScriptErrorRange is an [NSValue] that specifies a range.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSAppleScript/errorRange
	AppleScriptErrorRange string
	// ArgumentDomain is the identifier for the domain that contains command-line settings.
	//
	// See: https://developer.apple.com/documentation/Foundation/UserDefaults/argumentDomain
	ArgumentDomain string
	// AssertionHandlerKey is a key with a corresponding value in the thread dictionary.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSAssertionHandlerKey
	AssertionHandlerKey string
	// ExtensionItemAttachmentsKey is an optional array of media data associated with the extension item.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExtensionItemAttachmentsKey
	ExtensionItemAttachmentsKey string
	// ExtensionItemAttributedContentTextKey is an optional string describing the extension item content.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExtensionItemAttributedContentTextKey
	ExtensionItemAttributedContentTextKey string
	// ExtensionItemAttributedTitleKey is an optional title of the extension item.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExtensionItemAttributedTitleKey
	ExtensionItemAttributedTitleKey string
	// ExtensionItemsAndErrorsKey is the extension items and errors key.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExtensionItemsAndErrorsKey
	ExtensionItemsAndErrorsKey string
	// ExtensionJavaScriptPreprocessingResultsKey is a key whose value is an item of type `kUTTypePropertyList`. The item contains an [NSDictionary] that contains the object returned by the JavaScript code to its completion function.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExtensionJavaScriptPreprocessingResultsKey
	ExtensionJavaScriptPreprocessingResultsKey string
	// FileHandleNotificationDataItem is a key in the userinfo dictionary in a [readCompletionNotification] and [NSFileHandleReadToEndOfFileCompletion].
	//
	// See: https://developer.apple.com/documentation/Foundation/NSFileHandleNotificationDataItem
	FileHandleNotificationDataItem string
	// FileHandleNotificationFileHandleItem is a key in the userinfo dictionary in a [NSFileHandleConnectionAccepted] notification.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSFileHandleNotificationFileHandleItem
	FileHandleNotificationFileHandleItem string
	// FileManagerUnmountDissentingProcessIdentifierErrorKey is the process identifier of the process that prevented a volume from unmounting.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSFileManagerUnmountDissentingProcessIdentifierErrorKey
	FileManagerUnmountDissentingProcessIdentifierErrorKey string
	// GlobalDomain is the identifier for the domain that contains system-specified settings for all apps.
	//
	// See: https://developer.apple.com/documentation/Foundation/UserDefaults/globalDomain
	GlobalDomain string
	// GrammarCorrections is the value for the [NSGrammarCorrections] key should be an [NSArray] of [NSStrings] representing potential substitutions to correct the problem, but it is expected that this may not be available in all cases. [NSGrammarUserDescription] or [NSGrammarCorrections] must be supplied in order for correction guidance to be presented to the user.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSGrammarCorrections
	GrammarCorrections string
	// GrammarRange is the value for the [NSGrammarRange] dictionary key should be an [NSValue] containing an [NSRange], a subrange of the sentence range used as the return value, whose location should be an offset from the beginning of the sentence–so, for example, an [NSGrammarRange] for the first four characters of the overall sentence range should be `{0, 4}`. If the [NSGrammarRange] key is not present in the dictionary it is assumed to be equal to the overall sentence range.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSGrammarRange
	GrammarRange string
	// GrammarUserDescription is the value for the [NSGrammarUserDescription] dictionary key should be an [NSString] containing descriptive text about that range, to be presented directly to the user; it is intended that the user description should provide enough information to allow the user to correct the problem. It is recommended that [NSGrammarUserDescription] be supplied in all cases, however, [NSGrammarUserDescription] or [NSGrammarCorrections] must be supplied in order for correction guidance to be presented to the user.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSGrammarUserDescription
	GrammarUserDescription string
	// ItemProviderErrorDomain is the error domain associated with the item provider.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/errorDomain
	ItemProviderErrorDomain string
	// ItemProviderPreferredImageSizeKey is a key provided to the options dictionary to indicate a preferred image size.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSItemProviderPreferredImageSizeKey
	ItemProviderPreferredImageSizeKey string
	// KeyedArchiveRootObjectKey is archives created using the class method [archivedData(withRootObject:)] use this key for the root object in the hierarchy of encoded objects. The [NSKeyedUnarchiver] class method [unarchiveObject(with:)] looks for this root key as well.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSKeyedArchiveRootObjectKey
	KeyedArchiveRootObjectKey string
	// LoadedClasses is a constant used as a key for the `userInfo` dictionary of a [didLoadNotification] notification that corresponds to an array of names of each class that was loaded.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLoadedClasses
	LoadedClasses string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemAcquisitionMakeKey
	MetadataItemAcquisitionMakeKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemAcquisitionModelKey
	MetadataItemAcquisitionModelKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemAlbumKey
	MetadataItemAlbumKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemAltitudeKey
	MetadataItemAltitudeKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemApertureKey
	MetadataItemApertureKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemAppleLoopDescriptorsKey
	MetadataItemAppleLoopDescriptorsKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemAppleLoopsKeyFilterTypeKey
	MetadataItemAppleLoopsKeyFilterTypeKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemAppleLoopsLoopModeKey
	MetadataItemAppleLoopsLoopModeKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemAppleLoopsRootKeyKey
	MetadataItemAppleLoopsRootKeyKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemApplicationCategoriesKey
	MetadataItemApplicationCategoriesKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemAttributeChangeDateKey
	MetadataItemAttributeChangeDateKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemAudiencesKey
	MetadataItemAudiencesKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemAudioBitRateKey
	MetadataItemAudioBitRateKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemAudioChannelCountKey
	MetadataItemAudioChannelCountKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemAudioEncodingApplicationKey
	MetadataItemAudioEncodingApplicationKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemAudioSampleRateKey
	MetadataItemAudioSampleRateKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemAudioTrackNumberKey
	MetadataItemAudioTrackNumberKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemAuthorAddressesKey
	MetadataItemAuthorAddressesKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemAuthorEmailAddressesKey
	MetadataItemAuthorEmailAddressesKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemAuthorsKey
	MetadataItemAuthorsKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemBitsPerSampleKey
	MetadataItemBitsPerSampleKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemCFBundleIdentifierKey
	MetadataItemCFBundleIdentifierKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemCameraOwnerKey
	MetadataItemCameraOwnerKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemCityKey
	MetadataItemCityKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemCodecsKey
	MetadataItemCodecsKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemColorSpaceKey
	MetadataItemColorSpaceKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemCommentKey
	MetadataItemCommentKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemComposerKey
	MetadataItemComposerKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemContactKeywordsKey
	MetadataItemContactKeywordsKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemContentCreationDateKey
	MetadataItemContentCreationDateKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemContentModificationDateKey
	MetadataItemContentModificationDateKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemContentTypeKey
	MetadataItemContentTypeKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemContentTypeTreeKey
	MetadataItemContentTypeTreeKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemContributorsKey
	MetadataItemContributorsKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemCopyrightKey
	MetadataItemCopyrightKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemCountryKey
	MetadataItemCountryKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemCoverageKey
	MetadataItemCoverageKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemCreatorKey
	MetadataItemCreatorKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemDateAddedKey
	MetadataItemDateAddedKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemDeliveryTypeKey
	MetadataItemDeliveryTypeKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemDescriptionKey
	MetadataItemDescriptionKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemDirectorKey
	MetadataItemDirectorKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemDisplayNameKey
	MetadataItemDisplayNameKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemDownloadedDateKey
	MetadataItemDownloadedDateKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemDueDateKey
	MetadataItemDueDateKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemDurationSecondsKey
	MetadataItemDurationSecondsKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemEXIFGPSVersionKey
	MetadataItemEXIFGPSVersionKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemEXIFVersionKey
	MetadataItemEXIFVersionKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemEditorsKey
	MetadataItemEditorsKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemEmailAddressesKey
	MetadataItemEmailAddressesKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemEncodingApplicationsKey
	MetadataItemEncodingApplicationsKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemExecutableArchitecturesKey
	MetadataItemExecutableArchitecturesKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemExecutablePlatformKey
	MetadataItemExecutablePlatformKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemExposureModeKey
	MetadataItemExposureModeKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemExposureProgramKey
	MetadataItemExposureProgramKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemExposureTimeSecondsKey
	MetadataItemExposureTimeSecondsKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemExposureTimeStringKey
	MetadataItemExposureTimeStringKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemFNumberKey
	MetadataItemFNumberKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemFSContentChangeDateKey
	MetadataItemFSContentChangeDateKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemFSCreationDateKey
	MetadataItemFSCreationDateKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemFSNameKey
	MetadataItemFSNameKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemFSSizeKey
	MetadataItemFSSizeKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemFinderCommentKey
	MetadataItemFinderCommentKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemFlashOnOffKey
	MetadataItemFlashOnOffKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemFocalLength35mmKey
	MetadataItemFocalLength35mmKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemFocalLengthKey
	MetadataItemFocalLengthKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemFontsKey
	MetadataItemFontsKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemGPSAreaInformationKey
	MetadataItemGPSAreaInformationKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemGPSDOPKey
	MetadataItemGPSDOPKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemGPSDateStampKey
	MetadataItemGPSDateStampKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemGPSDestBearingKey
	MetadataItemGPSDestBearingKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemGPSDestDistanceKey
	MetadataItemGPSDestDistanceKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemGPSDestLatitudeKey
	MetadataItemGPSDestLatitudeKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemGPSDestLongitudeKey
	MetadataItemGPSDestLongitudeKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemGPSDifferentalKey
	MetadataItemGPSDifferentalKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemGPSMapDatumKey
	MetadataItemGPSMapDatumKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemGPSMeasureModeKey
	MetadataItemGPSMeasureModeKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemGPSProcessingMethodKey
	MetadataItemGPSProcessingMethodKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemGPSStatusKey
	MetadataItemGPSStatusKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemGPSTrackKey
	MetadataItemGPSTrackKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemGenreKey
	MetadataItemGenreKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemHasAlphaChannelKey
	MetadataItemHasAlphaChannelKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemHeadlineKey
	MetadataItemHeadlineKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemISOSpeedKey
	MetadataItemISOSpeedKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemIdentifierKey
	MetadataItemIdentifierKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemImageDirectionKey
	MetadataItemImageDirectionKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemInformationKey
	MetadataItemInformationKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemInstantMessageAddressesKey
	MetadataItemInstantMessageAddressesKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemInstructionsKey
	MetadataItemInstructionsKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemIsApplicationManagedKey
	MetadataItemIsApplicationManagedKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemIsGeneralMIDISequenceKey
	MetadataItemIsGeneralMIDISequenceKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemIsLikelyJunkKey
	MetadataItemIsLikelyJunkKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemIsUbiquitousKey
	MetadataItemIsUbiquitousKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemKeySignatureKey
	MetadataItemKeySignatureKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemKeywordsKey
	MetadataItemKeywordsKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemKindKey
	MetadataItemKindKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemLanguagesKey
	MetadataItemLanguagesKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemLastUsedDateKey
	MetadataItemLastUsedDateKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemLatitudeKey
	MetadataItemLatitudeKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemLayerNamesKey
	MetadataItemLayerNamesKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemLensModelKey
	MetadataItemLensModelKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemLongitudeKey
	MetadataItemLongitudeKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemLyricistKey
	MetadataItemLyricistKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemMaxApertureKey
	MetadataItemMaxApertureKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemMediaTypesKey
	MetadataItemMediaTypesKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemMeteringModeKey
	MetadataItemMeteringModeKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemMusicalGenreKey
	MetadataItemMusicalGenreKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemMusicalInstrumentCategoryKey
	MetadataItemMusicalInstrumentCategoryKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemMusicalInstrumentNameKey
	MetadataItemMusicalInstrumentNameKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemNamedLocationKey
	MetadataItemNamedLocationKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemNumberOfPagesKey
	MetadataItemNumberOfPagesKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemOrganizationsKey
	MetadataItemOrganizationsKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemOrientationKey
	MetadataItemOrientationKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemOriginalFormatKey
	MetadataItemOriginalFormatKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemOriginalSourceKey
	MetadataItemOriginalSourceKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemPageHeightKey
	MetadataItemPageHeightKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemPageWidthKey
	MetadataItemPageWidthKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemParticipantsKey
	MetadataItemParticipantsKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemPathKey
	MetadataItemPathKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemPerformersKey
	MetadataItemPerformersKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemPhoneNumbersKey
	MetadataItemPhoneNumbersKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemPixelCountKey
	MetadataItemPixelCountKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemPixelHeightKey
	MetadataItemPixelHeightKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemPixelWidthKey
	MetadataItemPixelWidthKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemProducerKey
	MetadataItemProducerKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemProfileNameKey
	MetadataItemProfileNameKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemProjectsKey
	MetadataItemProjectsKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemPublishersKey
	MetadataItemPublishersKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemRecipientAddressesKey
	MetadataItemRecipientAddressesKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemRecipientEmailAddressesKey
	MetadataItemRecipientEmailAddressesKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemRecipientsKey
	MetadataItemRecipientsKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemRecordingDateKey
	MetadataItemRecordingDateKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemRecordingYearKey
	MetadataItemRecordingYearKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemRedEyeOnOffKey
	MetadataItemRedEyeOnOffKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemResolutionHeightDPIKey
	MetadataItemResolutionHeightDPIKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemResolutionWidthDPIKey
	MetadataItemResolutionWidthDPIKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemRightsKey
	MetadataItemRightsKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemSecurityMethodKey
	MetadataItemSecurityMethodKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemSpeedKey
	MetadataItemSpeedKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemStarRatingKey
	MetadataItemStarRatingKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemStateOrProvinceKey
	MetadataItemStateOrProvinceKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemStreamableKey
	MetadataItemStreamableKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemSubjectKey
	MetadataItemSubjectKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemTempoKey
	MetadataItemTempoKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemTextContentKey
	MetadataItemTextContentKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemThemeKey
	MetadataItemThemeKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemTimeSignatureKey
	MetadataItemTimeSignatureKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemTimestampKey
	MetadataItemTimestampKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemTitleKey
	MetadataItemTitleKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemTotalBitRateKey
	MetadataItemTotalBitRateKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemURLKey
	MetadataItemURLKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemVersionKey
	MetadataItemVersionKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemVideoBitRateKey
	MetadataItemVideoBitRateKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemWhereFromsKey
	MetadataItemWhereFromsKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataItemWhiteBalanceKey
	MetadataItemWhiteBalanceKey string
	// MetadataQueryAccessibleUbiquitousExternalDocumentsScope is search for documents outside the app’s container. This search can locate iCloud documents that the user previously opened using a document picker view controller. This lets your app access the documents again without requiring direct user interaction. The result’s [NSMetadataItemURLKey] attributes return security-scoped NSURLs. For more information on working with security-scoped URLs, see [Security-Scoped URLs] in [NSURL].
	//
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataQueryAccessibleUbiquitousExternalDocumentsScope
	MetadataQueryAccessibleUbiquitousExternalDocumentsScope string
	// MetadataQueryIndexedLocalComputerScope is search all indexed local mounted volumes including the current user’s home directory (even if the home directory is remote).
	//
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataQueryIndexedLocalComputerScope
	MetadataQueryIndexedLocalComputerScope string
	// MetadataQueryIndexedNetworkScope is search all indexed user-mounted remote volumes.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataQueryIndexedNetworkScope
	MetadataQueryIndexedNetworkScope string
	// MetadataQueryLocalComputerScope is search all local mounted volumes, including the user home directory. The user’s home directory is searched even if it is a remote volume.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataQueryLocalComputerScope
	MetadataQueryLocalComputerScope string
	// MetadataQueryNetworkScope is search all user-mounted remote volumes.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataQueryNetworkScope
	MetadataQueryNetworkScope string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataQueryResultContentRelevanceAttribute
	MetadataQueryResultContentRelevanceAttribute string
	// MetadataQueryUbiquitousDataScope is search all files not in the [Documents] directories of the app’s iCloud container directories.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataQueryUbiquitousDataScope
	MetadataQueryUbiquitousDataScope string
	// MetadataQueryUbiquitousDocumentsScope is search all files in the [Documents] directories of the app’s iCloud container directories.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataQueryUbiquitousDocumentsScope
	MetadataQueryUbiquitousDocumentsScope string
	// MetadataQueryUpdateAddedItemsKey is the key for retrieving an array of items added to the query result. By default, this array contains [NSMetadataItem] objects, representing the query’s results; however, the query’s delegate can substitute these objects with instances of a different class.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataQueryUpdateAddedItemsKey
	MetadataQueryUpdateAddedItemsKey string
	// MetadataQueryUpdateChangedItemsKey is the key for retrieving an array of items that have changed in the query result. By default, this array contains [NSMetadataItem] objects, representing the query’s results; however, the query’s delegate can substitute these objects with instances of a different class.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataQueryUpdateChangedItemsKey
	MetadataQueryUpdateChangedItemsKey string
	// MetadataQueryUpdateRemovedItemsKey is the key for retrieving an array of items removed from the query result. By default, this array contains [NSMetadataItem] objects, representing the query’s results; however, the query’s delegate can substitute these objects with instances of a different class.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataQueryUpdateRemovedItemsKey
	MetadataQueryUpdateRemovedItemsKey string
	// MetadataQueryUserHomeScope is search the user’s home directory.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataQueryUserHomeScope
	MetadataQueryUserHomeScope string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousItemContainerDisplayNameKey
	MetadataUbiquitousItemContainerDisplayNameKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousItemDownloadRequestedKey
	MetadataUbiquitousItemDownloadRequestedKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousItemDownloadingErrorKey
	MetadataUbiquitousItemDownloadingErrorKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousItemDownloadingStatusCurrent
	MetadataUbiquitousItemDownloadingStatusCurrent string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousItemDownloadingStatusDownloaded
	MetadataUbiquitousItemDownloadingStatusDownloaded string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousItemDownloadingStatusKey
	MetadataUbiquitousItemDownloadingStatusKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousItemDownloadingStatusNotDownloaded
	MetadataUbiquitousItemDownloadingStatusNotDownloaded string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousItemHasUnresolvedConflictsKey
	MetadataUbiquitousItemHasUnresolvedConflictsKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousItemIsDownloadingKey
	MetadataUbiquitousItemIsDownloadingKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousItemIsExternalDocumentKey
	MetadataUbiquitousItemIsExternalDocumentKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousItemIsSharedKey
	MetadataUbiquitousItemIsSharedKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousItemIsUploadedKey
	MetadataUbiquitousItemIsUploadedKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousItemIsUploadingKey
	MetadataUbiquitousItemIsUploadingKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousItemPercentDownloadedKey
	MetadataUbiquitousItemPercentDownloadedKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousItemPercentUploadedKey
	MetadataUbiquitousItemPercentUploadedKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousItemURLInLocalContainerKey
	MetadataUbiquitousItemURLInLocalContainerKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousItemUploadingErrorKey
	MetadataUbiquitousItemUploadingErrorKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousSharedItemCurrentUserPermissionsKey
	MetadataUbiquitousSharedItemCurrentUserPermissionsKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousSharedItemCurrentUserRoleKey
	MetadataUbiquitousSharedItemCurrentUserRoleKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousSharedItemMostRecentEditorNameComponentsKey
	MetadataUbiquitousSharedItemMostRecentEditorNameComponentsKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousSharedItemOwnerNameComponentsKey
	MetadataUbiquitousSharedItemOwnerNameComponentsKey string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousSharedItemPermissionsReadOnly
	MetadataUbiquitousSharedItemPermissionsReadOnly string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousSharedItemPermissionsReadWrite
	MetadataUbiquitousSharedItemPermissionsReadWrite string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousSharedItemRoleOwner
	MetadataUbiquitousSharedItemRoleOwner string
	// See: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousSharedItemRoleParticipant
	MetadataUbiquitousSharedItemRoleParticipant string
	// NetServicesErrorCode is this key identifies the error that occurred during the most recent operation.
	//
	// See: https://developer.apple.com/documentation/Foundation/NetService/errorCode-swift.type.property
	NetServicesErrorCode string
	// See: https://developer.apple.com/documentation/Foundation/NSOperationNotSupportedForKeyException
	OperationNotSupportedForKeyException string
	// See: https://developer.apple.com/documentation/Foundation/NSPersonNameComponentDelimiter
	PersonNameComponentDelimiter string
	// See: https://developer.apple.com/documentation/Foundation/NSPersonNameComponentFamilyName
	PersonNameComponentFamilyName string
	// See: https://developer.apple.com/documentation/Foundation/NSPersonNameComponentGivenName
	PersonNameComponentGivenName string
	// See: https://developer.apple.com/documentation/Foundation/NSPersonNameComponentKey
	PersonNameComponentKey string
	// See: https://developer.apple.com/documentation/Foundation/NSPersonNameComponentMiddleName
	PersonNameComponentMiddleName string
	// See: https://developer.apple.com/documentation/Foundation/NSPersonNameComponentNickname
	PersonNameComponentNickname string
	// See: https://developer.apple.com/documentation/Foundation/NSPersonNameComponentPrefix
	PersonNameComponentPrefix string
	// See: https://developer.apple.com/documentation/Foundation/NSPersonNameComponentSuffix
	PersonNameComponentSuffix string
	// RegistrationDomain is the identifier for the domain that contains your app’s registered default values.
	//
	// See: https://developer.apple.com/documentation/Foundation/UserDefaults/registrationDomain
	RegistrationDomain string
	// URLAuthenticationMethodClientCertificate is use client certificate authentication for this protection space.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLAuthenticationMethodClientCertificate
	URLAuthenticationMethodClientCertificate string
	// URLAuthenticationMethodDefault is use the default authentication method for a protocol.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLAuthenticationMethodDefault
	URLAuthenticationMethodDefault string
	// URLAuthenticationMethodHTMLForm is use HTML form authentication for this protection space.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLAuthenticationMethodHTMLForm
	URLAuthenticationMethodHTMLForm string
	// URLAuthenticationMethodHTTPBasic is use HTTP basic authentication for this protection space.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLAuthenticationMethodHTTPBasic
	URLAuthenticationMethodHTTPBasic string
	// URLAuthenticationMethodHTTPDigest is use HTTP digest authentication for this protection space.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLAuthenticationMethodHTTPDigest
	URLAuthenticationMethodHTTPDigest string
	// URLAuthenticationMethodNTLM is use NTLM authentication for this protection space.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLAuthenticationMethodNTLM
	URLAuthenticationMethodNTLM string
	// URLAuthenticationMethodNegotiate is negotiate whether to use Kerberos or NTLM authentication for this protection space.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLAuthenticationMethodNegotiate
	URLAuthenticationMethodNegotiate string
	// URLAuthenticationMethodServerTrust is perform server trust authentication (certificate validation) for this protection space.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLAuthenticationMethodServerTrust
	URLAuthenticationMethodServerTrust string
	// URLCredentialStorageRemoveSynchronizableCredentials is the corresponding value is an [NSNumber] object representing a Boolean value that indicates whether credentials which contain the [URLCredential.Persistence.synchronizable] attribute should be removed.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLCredentialStorageRemoveSynchronizableCredentials
	URLCredentialStorageRemoveSynchronizableCredentials string
	// URLErrorBackgroundTaskCancelledReasonKey is a key in the error dictionary that provides the reason for canceling a background task.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLErrorBackgroundTaskCancelledReasonKey
	URLErrorBackgroundTaskCancelledReasonKey string
	// URLErrorFailingURLErrorKey is the URL which caused a load to fail.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLErrorFailingURLErrorKey
	URLErrorFailingURLErrorKey string
	// URLErrorFailingURLPeerTrustErrorKey is the state of a failed SSL handshake.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLErrorFailingURLPeerTrustErrorKey
	URLErrorFailingURLPeerTrustErrorKey string
	// URLErrorFailingURLStringErrorKey is the URL which caused a load to fail.
	//
	// Deprecated: Deprecated since macOS 15.4. Use NSURLErrorFailingURLErrorKey instead
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLErrorFailingURLStringErrorKey
	URLErrorFailingURLStringErrorKey string
	// URLFileScheme is identifies a URL that points to a file on a mounted volume.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLFileScheme
	URLFileScheme string
	// URLProtectionSpaceFTP is the protocol type for FTP.
	//
	// Deprecated: Deprecated since macOS 15.4. FTP is deprecated and only supported in the classic loading mode
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLProtectionSpaceFTP
	URLProtectionSpaceFTP string
	// URLProtectionSpaceFTPProxy is the proxy type for FTP proxies.
	//
	// Deprecated: Deprecated since macOS 15.4. FTP is deprecated and only supported in the classic loading mode
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLProtectionSpaceFTPProxy
	URLProtectionSpaceFTPProxy string
	// URLProtectionSpaceHTTP is the protocol type for HTTP.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLProtectionSpaceHTTP
	URLProtectionSpaceHTTP string
	// URLProtectionSpaceHTTPProxy is the proxy type for HTTP proxies.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLProtectionSpaceHTTPProxy
	URLProtectionSpaceHTTPProxy string
	// URLProtectionSpaceHTTPS is the protocol type for HTTPS.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLProtectionSpaceHTTPS
	URLProtectionSpaceHTTPS string
	// URLProtectionSpaceHTTPSProxy is the proxy type for HTTPS proxies.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLProtectionSpaceHTTPSProxy
	URLProtectionSpaceHTTPSProxy string
	// URLProtectionSpaceSOCKSProxy is the proxy type for SOCKS proxies.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLProtectionSpaceSOCKSProxy
	URLProtectionSpaceSOCKSProxy string
	// URLSessionDownloadTaskResumeData is a key in the error dictionary that provides resume data.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLSessionDownloadTaskResumeData
	URLSessionDownloadTaskResumeData string
	// URLSessionUploadTaskResumeData is key in the userInfo dictionary of an NSError received during a failed upload.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLSessionUploadTaskResumeData
	URLSessionUploadTaskResumeData string
	// UbiquitousKeyValueStoreChangeReasonKey is a key that indicates the reason why the key-value store changed.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStoreChangeReasonKey
	UbiquitousKeyValueStoreChangeReasonKey string
	// UbiquitousKeyValueStoreChangedKeysKey is a key that indicates which keys changed in the iCloud key-value store.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStoreChangedKeysKey
	UbiquitousKeyValueStoreChangedKeysKey string
	// UndoManagerGroupIsDiscardableKey is a key, used in a notification’s user info, that indicates the undo group contains only discardable actions.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSUndoManagerGroupIsDiscardableKey
	UndoManagerGroupIsDiscardableKey string
	// UserActivityTypeBrowsingWeb is an activity that continues from Handoff or a universal link.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSUserActivityTypeBrowsingWeb
	UserActivityTypeBrowsingWeb string
)

var (
	// AverageKeyValueOperator is the @avg array operator.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSKeyValueOperator/averageKeyValueOperator
	AverageKeyValueOperator NSKeyValueOperator
	// CountKeyValueOperator is the @count array operator.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSKeyValueOperator/countKeyValueOperator
	CountKeyValueOperator NSKeyValueOperator
	// DistinctUnionOfArraysKeyValueOperator is the @distinctUnionOfArrays array operator.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSKeyValueOperator/distinctUnionOfArraysKeyValueOperator
	DistinctUnionOfArraysKeyValueOperator NSKeyValueOperator
	// DistinctUnionOfObjectsKeyValueOperator is the @distinctUnionOfObjects array operator.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSKeyValueOperator/distinctUnionOfObjectsKeyValueOperator
	DistinctUnionOfObjectsKeyValueOperator NSKeyValueOperator
	// DistinctUnionOfSetsKeyValueOperator is the @distinctUnionOfSets array operator.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSKeyValueOperator/distinctUnionOfSetsKeyValueOperator
	DistinctUnionOfSetsKeyValueOperator NSKeyValueOperator
	// MaximumKeyValueOperator is the @max array operator.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSKeyValueOperator/maximumKeyValueOperator
	MaximumKeyValueOperator NSKeyValueOperator
	// MinimumKeyValueOperator is the @min array operator.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSKeyValueOperator/minimumKeyValueOperator
	MinimumKeyValueOperator NSKeyValueOperator
	// SumKeyValueOperator is the @sum array operator.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSKeyValueOperator/sumKeyValueOperator
	SumKeyValueOperator NSKeyValueOperator
	// UnionOfArraysKeyValueOperator is the @unionOfArrays array operator.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSKeyValueOperator/unionOfArraysKeyValueOperator
	UnionOfArraysKeyValueOperator NSKeyValueOperator
	// UnionOfObjectsKeyValueOperator is the @unionOfObjects array operator.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSKeyValueOperator/unionOfObjectsKeyValueOperator
	UnionOfObjectsKeyValueOperator NSKeyValueOperator
	// UnionOfSetsKeyValueOperator is the @unionOfSets array operator.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSKeyValueOperator/unionOfSetsKeyValueOperator
	UnionOfSetsKeyValueOperator NSKeyValueOperator
)

var (
)

var (
	// CharacterConversionException is [NSString] raises an [NSCharacterConversionException] if a string cannot be represented in a file-system or string encoding.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/characterConversionException
	CharacterConversionException NSExceptionName
	// DecimalNumberDivideByZeroException is the exception raised on divide by zero.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/decimalNumberDivideByZeroException
	DecimalNumberDivideByZeroException NSExceptionName
	// DecimalNumberExactnessException is the exception raised if there is an exactness error.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/decimalNumberExactnessException
	DecimalNumberExactnessException NSExceptionName
	// DecimalNumberOverflowException is the exception raised on overflow.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/decimalNumberOverflowException
	DecimalNumberOverflowException NSExceptionName
	// DecimalNumberUnderflowException is the exception raised on underflow.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/decimalNumberUnderflowException
	DecimalNumberUnderflowException NSExceptionName
	// DestinationInvalidException is name of an exception that occurs when an internal assertion fails and implies an unexpected condition within the distributed objects.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/destinationInvalidException
	DestinationInvalidException NSExceptionName
	// FileHandleOperationException is raised by [NSFileHandle] if attempts to determine file-handle type fail or if attempts to read from a file or channel fail.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/fileHandleOperationException
	FileHandleOperationException NSExceptionName
	// GenericException is a generic name for an exception.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/genericException
	GenericException NSExceptionName
	// InconsistentArchiveException is the name of an exception raised by [NSArchiver] if there are problems initializing or encoding.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/inconsistentArchiveException
	InconsistentArchiveException NSExceptionName
	// InternalInconsistencyException is name of an exception that occurs when an internal assertion fails and implies an unexpected condition within the called code.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/internalInconsistencyException
	InternalInconsistencyException NSExceptionName
	// InvalidArchiveOperationException is the name of the exception raised by [NSKeyedArchiver] if there is a problem creating an archive.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArchiveOperationException
	InvalidArchiveOperationException NSExceptionName
	// InvalidArgumentException is name of an exception that occurs when you pass an invalid argument to a method, such as a `nil` pointer where a non-`nil` object is required.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
	InvalidArgumentException NSExceptionName
	// InvalidReceivePortException is name of an exception that occurs when the receive port of an [NSConnection] has become invalid.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidReceivePortException
	InvalidReceivePortException NSExceptionName
	// InvalidSendPortException is name of an exception that occurs when the send port of an [NSConnection] has become invalid.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidSendPortException
	InvalidSendPortException NSExceptionName
	// InvalidUnarchiveOperationException is the name of the exception raised by [NSKeyedArchiver] if there is a problem extracting an archive.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidUnarchiveOperationException
	InvalidUnarchiveOperationException NSExceptionName
	// InvocationOperationCancelledException is the name of the exception raised if the [result] method is called after the operation was cancelled.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/invocationOperationCancelledException
	InvocationOperationCancelledException NSExceptionName
	// InvocationOperationVoidResultException is the name of the exception raised if the [result] method is called for an invocation method with a `void` return type.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/invocationOperationVoidResultException
	InvocationOperationVoidResultException NSExceptionName
	// MallocException is obsolete; not currently used.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/mallocException
	MallocException NSExceptionName
	// ObjectInaccessibleException is name of an exception that occurs when a remote object is accessed from a thread that should not access it.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/objectInaccessibleException
	ObjectInaccessibleException NSExceptionName
	// ObjectNotAvailableException is name of an exception that occurs when the remote side of the [NSConnection] refused to send the message to the object because the object has never been vended.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/objectNotAvailableException
	ObjectNotAvailableException NSExceptionName
	// OldStyleException is no longer used.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/oldStyleException
	OldStyleException NSExceptionName
	// ParseErrorException is [NSString] raises an [NSParseErrorException] if a string cannot be parsed as a property list.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/parseErrorException
	ParseErrorException NSExceptionName
	// PortReceiveException is generic error occurred on receive.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/portReceiveException
	PortReceiveException NSExceptionName
	// PortSendException is generic error occurred on send.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/portSendException
	PortSendException NSExceptionName
	// PortTimeoutException is name of an exception that occurs when a timeout set on a port expires during a send or receive operation.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/portTimeoutException
	PortTimeoutException NSExceptionName
	// RangeException is name of an exception that occurs when attempting to access outside the bounds of some data, such as beyond the end of a string.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/rangeException
	RangeException NSExceptionName
	// UndefinedKeyException is raised when a key value coding operation fails.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/undefinedKeyException
	UndefinedKeyException NSExceptionName
)

var (
	// CocoaErrorDomain is cocoa errors.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSCocoaErrorDomain
	CocoaErrorDomain NSErrorDomain
	// MachErrorDomain is mach errors.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSMachErrorDomain
	MachErrorDomain NSErrorDomain
	// NetServicesErrorDomain is this key identifies the originator of the error, which is either the [NSNetService] object or the mach network layer. For most errors, you should not need the value provided by this key.
	//
	// See: https://developer.apple.com/documentation/Foundation/NetService/errorDomain
	NetServicesErrorDomain NSErrorDomain
	// OSStatusErrorDomain is mac OS 9/Carbon errors.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSOSStatusErrorDomain
	OSStatusErrorDomain NSErrorDomain
	// POSIXErrorDomain is pOSIX/BSD errors.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSPOSIXErrorDomain
	POSIXErrorDomain NSErrorDomain
	// StreamSOCKSErrorDomain is the error domain used by [NSError] when reporting SOCKS errors.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSStreamSOCKSErrorDomain
	StreamSOCKSErrorDomain NSErrorDomain
	// StreamSocketSSLErrorDomain is the error domain used by [NSError] when reporting SSL errors.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSStreamSocketSSLErrorDomain
	StreamSocketSSLErrorDomain NSErrorDomain
	// URLErrorDomain is uRL loading system errors.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLErrorDomain
	URLErrorDomain NSErrorDomain
	// XMLParserErrorDomain is indicates an error in XML parsing.
	//
	// See: https://developer.apple.com/documentation/Foundation/XMLParser/errorDomain
	XMLParserErrorDomain NSErrorDomain
)

var (
	// DeallocateZombies is a global variable that determines whether or not the memory of zombie objects is deallocated.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSDeallocateZombies
	DeallocateZombies bool
	// DebugEnabled is a global variable that can be used to enable debug behavior in your app, such as extra logging.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSDebugEnabled
	DebugEnabled bool
	// KeepAllocationStatistics is a no-longer-used global variable related to keeping statistics.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSKeepAllocationStatistics
	KeepAllocationStatistics bool
	// ZombieEnabled is a global variable related to zombie objects that in practice has no effect.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSZombieEnabled
	ZombieEnabled bool
)

var (
	// See: https://developer.apple.com/documentation/Foundation/NSDebugDescriptionErrorKey
	DebugDescriptionErrorKey NSErrorUserInfoKey
	// FilePathErrorKey is contains the file path of the error.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSFilePathErrorKey
	FilePathErrorKey NSErrorUserInfoKey
	// HelpAnchorErrorKey is the corresponding value is an [NSString] containing the localized help corresponding to the help button. See [helpAnchor] for more information.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSHelpAnchorErrorKey
	HelpAnchorErrorKey NSErrorUserInfoKey
	// LocalizedDescriptionKey is the corresponding value is a localized string representation of the error that, if present, will be returned by [localizedDescription].
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocalizedDescriptionKey
	LocalizedDescriptionKey NSErrorUserInfoKey
	// See: https://developer.apple.com/documentation/Foundation/NSLocalizedFailureErrorKey
	LocalizedFailureErrorKey NSErrorUserInfoKey
	// LocalizedFailureReasonErrorKey is the corresponding value is a localized string representation containing the reason for the failure that, if present, will be returned by [localizedFailureReason].
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocalizedFailureReasonErrorKey
	LocalizedFailureReasonErrorKey NSErrorUserInfoKey
	// LocalizedRecoveryOptionsErrorKey is the corresponding value is an array containing the localized titles of buttons appropriate for displaying in an alert panel.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocalizedRecoveryOptionsErrorKey
	LocalizedRecoveryOptionsErrorKey NSErrorUserInfoKey
	// LocalizedRecoverySuggestionErrorKey is the corresponding value is a string containing the localized recovery suggestion for the error.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocalizedRecoverySuggestionErrorKey
	LocalizedRecoverySuggestionErrorKey NSErrorUserInfoKey
	// See: https://developer.apple.com/documentation/Foundation/NSMultipleUnderlyingErrorsKey
	MultipleUnderlyingErrorsKey NSErrorUserInfoKey
	// RecoveryAttempterErrorKey is the corresponding value is an object that conforms to the NSErrorRecoveryAttempting informal protocol.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSRecoveryAttempterErrorKey
	RecoveryAttempterErrorKey NSErrorUserInfoKey
	// StringEncodingErrorKey is the corresponding value is an [NSNumber] object containing the [NSStringEncoding] value.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSStringEncodingErrorKey
	StringEncodingErrorKey NSErrorUserInfoKey
	// URLErrorKey is the corresponding value is an [NSURL] object.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLErrorKey
	URLErrorKey NSErrorUserInfoKey
	// URLErrorNetworkUnavailableReasonKey is the reason the network was unavailable for a task.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLErrorNetworkUnavailableReasonKey
	URLErrorNetworkUnavailableReasonKey NSErrorUserInfoKey
	// UnderlyingErrorKey is the corresponding value is an error that was encountered in an underlying implementation and caused the error that the receiver represents to occur.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSUnderlyingErrorKey
	UnderlyingErrorKey NSErrorUserInfoKey
)

var (
	// RunLoopDefaultMode is the mode set to handle input sources other than connection objects.
	//
	// See: https://developer.apple.com/documentation/Foundation/RunLoop/Mode/default
	RunLoopDefaultMode NSRunLoopMode
	// RunLoopCommonModes is a pseudo-mode that includes one or more other run loop modes.
	//
	// See: https://developer.apple.com/documentation/Foundation/RunLoop/Mode/common
	RunLoopCommonModes NSRunLoopMode
)

var (
	// EdgeInsetsZero is an edge insets structure with a zero inset on each edge.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSEdgeInsetsZero
	EdgeInsetsZero NSEdgeInsets
)

var (
	// FileAppendOnly is the key in a file attribute dictionary whose value indicates whether the file is read-only.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey/appendOnly
	FileAppendOnly NSFileAttributeKey
	// FileBusy is the key in a file attribute dictionary whose value indicates whether the file is busy.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey/busy
	FileBusy NSFileAttributeKey
	// FileCreationDate is the key in a file attribute dictionary whose value indicates the file’s creation date.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey/creationDate
	FileCreationDate NSFileAttributeKey
	// FileDeviceIdentifier is the key in a file attribute dictionary whose value indicates the identifier for the device on which the file resides.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey/deviceIdentifier
	FileDeviceIdentifier NSFileAttributeKey
	// FileExtensionHidden is the key in a file attribute dictionary whose value indicates whether the file’s extension is hidden.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey/extensionHidden
	FileExtensionHidden NSFileAttributeKey
	// FileGroupOwnerAccountID is the key in a file attribute dictionary whose value indicates the file’s group ID.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey/groupOwnerAccountID
	FileGroupOwnerAccountID NSFileAttributeKey
	// FileGroupOwnerAccountName is the key in a file attribute dictionary whose value indicates the group name of the file’s owner.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey/groupOwnerAccountName
	FileGroupOwnerAccountName NSFileAttributeKey
	// FileHFSCreatorCode is the key in a file attribute dictionary whose value indicates the file’s HFS creator code.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey/hfsCreatorCode
	FileHFSCreatorCode NSFileAttributeKey
	// FileHFSTypeCode is the key in a file attribute dictionary whose value indicates the file’s HFS type code.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey/hfsTypeCode
	FileHFSTypeCode NSFileAttributeKey
	// FileImmutable is the key in a file attribute dictionary whose value indicates whether the file is mutable.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey/immutable
	FileImmutable NSFileAttributeKey
	// FileModificationDate is the key in a file attribute dictionary whose value indicates the file’s last modified date.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey/modificationDate
	FileModificationDate NSFileAttributeKey
	// FileOwnerAccountID is the key in a file attribute dictionary whose value indicates the file’s owner’s account ID.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey/ownerAccountID
	FileOwnerAccountID NSFileAttributeKey
	// FileOwnerAccountName is the key in a file attribute dictionary whose value indicates the name of the file’s owner.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey/ownerAccountName
	FileOwnerAccountName NSFileAttributeKey
	// FilePosixPermissions is the key in a file attribute dictionary whose value indicates the file’s Posix permissions.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey/posixPermissions
	FilePosixPermissions NSFileAttributeKey
	// FileProtectionKey is the key in a file attribute dictionary whose value identifies the protection level for this file.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey/protectionKey
	FileProtectionKey NSFileAttributeKey
	// FileReferenceCount is the key in a file attribute dictionary whose value indicates the file’s reference count.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey/referenceCount
	FileReferenceCount NSFileAttributeKey
	// FileSize is the key in a file attribute dictionary whose value indicates the file’s size in bytes.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey/size
	FileSize NSFileAttributeKey
	// FileSystemFileNumber is the key in a file attribute dictionary whose value indicates the file’s filesystem file number.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey/systemFileNumber
	FileSystemFileNumber NSFileAttributeKey
	// FileSystemFreeNodes is the key in a file system attribute dictionary whose value indicates the number of free nodes in the file system.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey/systemFreeNodes
	FileSystemFreeNodes NSFileAttributeKey
	// FileSystemFreeSize is the key in a file system attribute dictionary whose value indicates the amount of free space on the file system.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey/systemFreeSize
	FileSystemFreeSize NSFileAttributeKey
	// FileSystemNodes is the key in a file system attribute dictionary whose value indicates the number of nodes in the file system.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey/systemNodes
	FileSystemNodes NSFileAttributeKey
	// FileSystemNumber is the key in a file system attribute dictionary whose value indicates the filesystem number of the file system.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey/systemNumber
	FileSystemNumber NSFileAttributeKey
	// FileSystemSize is the key in a file system attribute dictionary whose value indicates the size of the file system.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey/systemSize
	FileSystemSize NSFileAttributeKey
	// FileType is the key in a file attribute dictionary whose value indicates the file’s type.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeKey/type
	FileType NSFileAttributeKey
)

var (
	// FileProtectionComplete is the file is stored in an encrypted format on disk and cannot be read from or written to while the device is locked or booting.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileProtectionType/complete
	FileProtectionComplete NSFileProtectionType
	// FileProtectionCompleteUnlessOpen is the file is stored in an encrypted format on disk after it is closed.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileProtectionType/completeUnlessOpen
	FileProtectionCompleteUnlessOpen NSFileProtectionType
	// FileProtectionCompleteUntilFirstUserAuthentication is the file is stored in an encrypted format on disk and cannot be accessed until after the device has booted.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileProtectionType/completeUntilFirstUserAuthentication
	FileProtectionCompleteUntilFirstUserAuthentication NSFileProtectionType
	// FileProtectionNone is the file has no special protections associated with it.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileProtectionType/none
	FileProtectionNone NSFileProtectionType
)

var (
	// FileTypeBlockSpecial is a block special file.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeType/typeBlockSpecial
	FileTypeBlockSpecial NSFileAttributeType
	// FileTypeCharacterSpecial is a character special file.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeType/typeCharacterSpecial
	FileTypeCharacterSpecial NSFileAttributeType
	// FileTypeDirectory is a directory.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeType/typeDirectory
	FileTypeDirectory NSFileAttributeType
	// FileTypeRegular is a regular file.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeType/typeRegular
	FileTypeRegular NSFileAttributeType
	// FileTypeSocket is a socket.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeType/typeSocket
	FileTypeSocket NSFileAttributeType
	// FileTypeSymbolicLink is a symbolic link.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeType/typeSymbolicLink
	FileTypeSymbolicLink NSFileAttributeType
	// FileTypeUnknown is a file whose type is unknown.
	//
	// See: https://developer.apple.com/documentation/Foundation/FileAttributeType/typeUnknown
	FileTypeUnknown NSFileAttributeType
)

var (
	// HTTPCookieComment is an [NSString] object containing the comment for the cookie.
	//
	// See: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/comment
	HTTPCookieComment NSHTTPCookiePropertyKey
	// HTTPCookieCommentURL is an [NSURL] object or [NSString] object containing the comment URL for the cookie.
	//
	// See: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/commentURL
	HTTPCookieCommentURL NSHTTPCookiePropertyKey
	// HTTPCookieDiscard is an [NSString] object stating whether the cookie should be discarded at the end of the session.
	//
	// See: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/discard
	HTTPCookieDiscard NSHTTPCookiePropertyKey
	// HTTPCookieDomain is an [NSString] object containing the domain for the cookie.
	//
	// See: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/domain
	HTTPCookieDomain NSHTTPCookiePropertyKey
	// HTTPCookieExpires is an [NSDate] object or [NSString] object specifying the expiration date for the cookie.
	//
	// See: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/expires
	HTTPCookieExpires NSHTTPCookiePropertyKey
	// HTTPCookieMaximumAge is an [NSString] object containing an integer value stating how long in seconds the cookie should be kept, at most.
	//
	// See: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/maximumAge
	HTTPCookieMaximumAge NSHTTPCookiePropertyKey
	// HTTPCookieName is an [NSString] object containing the name of the cookie (required).
	//
	// See: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/name
	HTTPCookieName NSHTTPCookiePropertyKey
	// HTTPCookieOriginURL is an NSURL or [NSString] object containing the URL that set this cookie.
	//
	// See: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/originURL
	HTTPCookieOriginURL NSHTTPCookiePropertyKey
	// HTTPCookiePath is an [NSString] object containing the path for the cookie.
	//
	// See: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/path
	HTTPCookiePath NSHTTPCookiePropertyKey
	// HTTPCookiePort is an [NSString] object containing comma-separated integer values specifying the ports for the cookie.
	//
	// See: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/port
	HTTPCookiePort NSHTTPCookiePropertyKey
	// HTTPCookieSameSitePolicy is a string indicating the same-site policy for the cookie.
	//
	// See: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/sameSitePolicy
	HTTPCookieSameSitePolicy NSHTTPCookiePropertyKey
	// HTTPCookieSecure is an [NSString] object indicating that the cookie should be transmitted only over secure channels.
	//
	// See: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/secure
	HTTPCookieSecure NSHTTPCookiePropertyKey
	// See: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/setByJavaScript
	HTTPCookieSetByJavaScript NSHTTPCookiePropertyKey
	// HTTPCookieValue is an [NSString] object containing the value of the cookie.
	//
	// See: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/value
	HTTPCookieValue NSHTTPCookiePropertyKey
	// HTTPCookieVersion is an [NSString] object that specifies the version of the cookie.
	//
	// See: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/version
	HTTPCookieVersion NSHTTPCookiePropertyKey
)

var (
	// HTTPCookieSameSiteLax is a policy that allows certain cross-site requests to include the cookie.
	//
	// See: https://developer.apple.com/documentation/Foundation/HTTPCookieStringPolicy/sameSiteLax
	HTTPCookieSameSiteLax NSHTTPCookieStringPolicy
	// HTTPCookieSameSiteStrict is a policy that prohibits a cross-site request from including the cookie.
	//
	// See: https://developer.apple.com/documentation/Foundation/HTTPCookieStringPolicy/sameSiteStrict
	HTTPCookieSameSiteStrict NSHTTPCookieStringPolicy
)

var (
	// See: https://developer.apple.com/documentation/Foundation/NSAttributedStringFormattingContextKey/inflectionConceptsKey
	InflectionConceptsKey NSAttributedStringFormattingContextKey
)

var (
	// IntegerHashCallBacks is for sets of [NSInteger]-sized quantities or smaller (for example, `int`, `long`, or `unichar`).
	//
	// See: https://developer.apple.com/documentation/Foundation/NSIntegerHashCallBacks
	IntegerHashCallBacks NSHashTableCallBacks
	// NonOwnedPointerHashCallBacks is for sets of pointers, hashed by address.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNonOwnedPointerHashCallBacks
	NonOwnedPointerHashCallBacks NSHashTableCallBacks
	// NonRetainedObjectHashCallBacks is for sets of objects, but without retaining/releasing.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNonRetainedObjectHashCallBacks
	NonRetainedObjectHashCallBacks NSHashTableCallBacks
	// ObjectHashCallBacks is for sets of objects (similar to [NSSet]).
	//
	// See: https://developer.apple.com/documentation/Foundation/NSObjectHashCallBacks
	ObjectHashCallBacks NSHashTableCallBacks
	// OwnedObjectIdentityHashCallBacks is for sets of objects, with transfer of ownership upon insertion, using pointer equality.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSOwnedObjectIdentityHashCallBacks
	OwnedObjectIdentityHashCallBacks NSHashTableCallBacks
	// OwnedPointerHashCallBacks is for sets of pointers, with transfer of ownership upon insertion.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSOwnedPointerHashCallBacks
	OwnedPointerHashCallBacks NSHashTableCallBacks
	// PointerToStructHashCallBacks is for sets of pointers to structs, when the first field of the struct is `int`-sized.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSPointerToStructHashCallBacks
	PointerToStructHashCallBacks NSHashTableCallBacks
)

var (
	// IntegerMapKeyCallBacks is for keys that are pointer-sized quantities or smaller (for example, `int`, `long`, or `unichar`).
	//
	// See: https://developer.apple.com/documentation/Foundation/NSIntegerMapKeyCallBacks
	IntegerMapKeyCallBacks NSMapTableKeyCallBacks
	// NonOwnedPointerMapKeyCallBacks is for keys that are pointers not freed.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNonOwnedPointerMapKeyCallBacks
	NonOwnedPointerMapKeyCallBacks NSMapTableKeyCallBacks
	// NonOwnedPointerOrNullMapKeyCallBacks is for keys that are pointers not freed, or [NULL].
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNonOwnedPointerOrNullMapKeyCallBacks
	NonOwnedPointerOrNullMapKeyCallBacks NSMapTableKeyCallBacks
	// NonRetainedObjectMapKeyCallBacks is for sets of objects, but without retaining/releasing.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNonRetainedObjectMapKeyCallBacks
	NonRetainedObjectMapKeyCallBacks NSMapTableKeyCallBacks
	// ObjectMapKeyCallBacks is for keys that are objects.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSObjectMapKeyCallBacks
	ObjectMapKeyCallBacks NSMapTableKeyCallBacks
	// OwnedPointerMapKeyCallBacks is for keys that are pointers, with transfer of ownership upon insertion.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSOwnedPointerMapKeyCallBacks
	OwnedPointerMapKeyCallBacks NSMapTableKeyCallBacks
)

var (
	// IntegerMapValueCallBacks is for values that are pointer-sized quantities, (for example, `int`, `long`, or `unichar`).
	//
	// See: https://developer.apple.com/documentation/Foundation/NSIntegerMapValueCallBacks
	IntegerMapValueCallBacks NSMapTableValueCallBacks
	// NonOwnedPointerMapValueCallBacks is for values that are not owned pointers.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNonOwnedPointerMapValueCallBacks
	NonOwnedPointerMapValueCallBacks NSMapTableValueCallBacks
	// NonRetainedObjectMapValueCallBacks is for sets of objects, but without retaining/releasing.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSNonRetainedObjectMapValueCallBacks
	NonRetainedObjectMapValueCallBacks NSMapTableValueCallBacks
	// ObjectMapValueCallBacks is for values that are objects.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSObjectMapValueCallBacks
	ObjectMapValueCallBacks NSMapTableValueCallBacks
	// OwnedPointerMapValueCallBacks is for values that are owned pointers.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSOwnedPointerMapValueCallBacks
	OwnedPointerMapValueCallBacks NSMapTableValueCallBacks
)

var (
	// IsNilTransformerName is this value transformer returns true if the value is nil.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSValueTransformerName/isNilTransformerName
	IsNilTransformerName NSValueTransformerName
	// IsNotNilTransformerName is this value transformer returns true if the value is non-nil.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSValueTransformerName/isNotNilTransformerName
	IsNotNilTransformerName NSValueTransformerName
	// NegateBooleanTransformerName is this value transformer negates a boolean value, transforming true to false and false to true.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSValueTransformerName/negateBooleanTransformerName
	NegateBooleanTransformerName NSValueTransformerName
	// SecureUnarchiveFromDataTransformerName is the name of the value transformer that creates then returns an object by attempting to unarchive the data to a class that supports secure coding.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSValueTransformerName/secureUnarchiveFromDataTransformerName
	SecureUnarchiveFromDataTransformerName NSValueTransformerName
)

var (
	// KeyValueChangeIndexesKey is if the value of the [kindKey] entry is [NSKeyValueChange.insertion], [NSKeyValueChange.removal], or [NSKeyValueChange.replacement], the value of this key is an [NSIndexSet] object that contains the indexes of the inserted, removed, or replaced objects.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSKeyValueChangeKey/indexesKey
	KeyValueChangeIndexesKey NSKeyValueChangeKey
	// KeyValueChangeKindKey is an [NSNumber] object that contains a value corresponding to one of the [NSKeyValueChange] enums, indicating what sort of change has occurred.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSKeyValueChangeKey/kindKey
	KeyValueChangeKindKey NSKeyValueChangeKey
	// KeyValueChangeNewKey is if the value of the [kindKey] entry is [NSKeyValueChange.setting], and [new] was specified when the observer was registered, the value of this key is the new value for the attribute.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSKeyValueChangeKey/newKey
	KeyValueChangeNewKey NSKeyValueChangeKey
	// KeyValueChangeNotificationIsPriorKey is if the [prior] option was specified when the observer was registered this notification is sent prior to a change.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSKeyValueChangeKey/notificationIsPriorKey
	KeyValueChangeNotificationIsPriorKey NSKeyValueChangeKey
	// KeyValueChangeOldKey is if the value of the [kindKey] entry is [NSKeyValueChange.setting], and [old] was specified when the observer was registered, the value of this key is the value before the attribute was changed.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSKeyValueChangeKey/oldKey
	KeyValueChangeOldKey NSKeyValueChangeKey
)

var (
)

var (
)

var (
	// LocalNotificationCenterType is distributes notifications to all tasks on the sender’s computer.
	//
	// See: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/CenterType/localNotificationCenterType
	LocalNotificationCenterType NSDistributedNotificationCenterType
)

var (
	// LocaleAlternateQuotationBeginDelimiterKey is the alternating begin quotation symbol associated with the locale.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocale/Key/alternateQuotationBeginDelimiterKey
	LocaleAlternateQuotationBeginDelimiterKey NSLocaleKey
	// LocaleAlternateQuotationEndDelimiterKey is the alternate end quotation symbol associated with the locale.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocale/Key/alternateQuotationEndDelimiterKey
	LocaleAlternateQuotationEndDelimiterKey NSLocaleKey
	// LocaleCalendar is the calendar associated with the locale.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocale/Key/calendar
	LocaleCalendar NSLocaleKey
	// LocaleCollationIdentifier is the collation associated with the locale.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocale/Key/collationIdentifier
	LocaleCollationIdentifier NSLocaleKey
	// LocaleCollatorIdentifier is the collation identifier for the locale.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocale/Key/collatorIdentifier
	LocaleCollatorIdentifier NSLocaleKey
	// LocaleCountryCode is the locale country or region code.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocale/Key/countryCode
	LocaleCountryCode NSLocaleKey
	// LocaleCurrencyCode is the currency code associated with the locale.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocale/Key/currencyCode
	LocaleCurrencyCode NSLocaleKey
	// LocaleCurrencySymbol is the currency symbol associated with the locale.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocale/Key/currencySymbol
	LocaleCurrencySymbol NSLocaleKey
	// LocaleDecimalSeparator is the decimal separator associated with the locale.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocale/Key/decimalSeparator
	LocaleDecimalSeparator NSLocaleKey
	// LocaleExemplarCharacterSet is the exemplar character set for the locale.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocale/Key/exemplarCharacterSet
	LocaleExemplarCharacterSet NSLocaleKey
	// LocaleGroupingSeparator is the numeric grouping separator associated with the locale.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocale/Key/groupingSeparator
	LocaleGroupingSeparator NSLocaleKey
	// LocaleIdentifier is the locale identifier.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocale/Key/identifier
	LocaleIdentifier NSLocaleKey
	// LocaleLanguageCode is the locale language code.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocale/Key/languageCode
	LocaleLanguageCode NSLocaleKey
	// LocaleMeasurementSystem is the measurement system associated with the locale.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocale/Key/measurementSystem
	LocaleMeasurementSystem NSLocaleKey
	// LocaleQuotationBeginDelimiterKey is the begin quotation symbol associated with the locale.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocale/Key/quotationBeginDelimiterKey
	LocaleQuotationBeginDelimiterKey NSLocaleKey
	// LocaleQuotationEndDelimiterKey is the end quotation symbol associated with the locale.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocale/Key/quotationEndDelimiterKey
	LocaleQuotationEndDelimiterKey NSLocaleKey
	// LocaleScriptCode is the locale script code.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocale/Key/scriptCode
	LocaleScriptCode NSLocaleKey
	// LocaleUsesMetricSystem is a flag that indicates whether the locale uses the metric system.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocale/Key/usesMetricSystem
	LocaleUsesMetricSystem NSLocaleKey
	// LocaleVariantCode is the locale variant code.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocale/Key/variantCode
	LocaleVariantCode NSLocaleKey
)

var (
	// ProgressEstimatedTimeRemainingKey is a key with a corresponding value that represents the time remaining, in seconds.
	//
	// See: https://developer.apple.com/documentation/Foundation/ProgressUserInfoKey/estimatedTimeRemainingKey
	ProgressEstimatedTimeRemainingKey NSProgressUserInfoKey
	// ProgressFileAnimationImageKey is a key with a corresponding value that is an image, typically an icon to represent the file.
	//
	// See: https://developer.apple.com/documentation/Foundation/ProgressUserInfoKey/fileAnimationImageKey
	ProgressFileAnimationImageKey NSProgressUserInfoKey
	// ProgressFileAnimationImageOriginalRectKey is a key with a corresponding value that indicates the starting location of the image onscreen.
	//
	// See: https://developer.apple.com/documentation/Foundation/ProgressUserInfoKey/fileAnimationImageOriginalRectKey
	ProgressFileAnimationImageOriginalRectKey NSProgressUserInfoKey
	// ProgressFileCompletedCountKey is a key with a corresponding value that represents the number of completed files.
	//
	// See: https://developer.apple.com/documentation/Foundation/ProgressUserInfoKey/fileCompletedCountKey
	ProgressFileCompletedCountKey NSProgressUserInfoKey
	// ProgressFileIconKey is a key with a corresponding value that must be an image, typically an icon to represent the file.
	//
	// See: https://developer.apple.com/documentation/Foundation/ProgressUserInfoKey/fileIconKey
	ProgressFileIconKey NSProgressUserInfoKey
	// ProgressFileOperationKindKey is a key with a corresponding value that indicates the kind of file operation a progress object represents.
	//
	// See: https://developer.apple.com/documentation/Foundation/ProgressUserInfoKey/fileOperationKindKey
	ProgressFileOperationKindKey NSProgressUserInfoKey
	// ProgressFileTotalCountKey is a key with a corresponding value that represents the total number of files within a file operation.
	//
	// See: https://developer.apple.com/documentation/Foundation/ProgressUserInfoKey/fileTotalCountKey
	ProgressFileTotalCountKey NSProgressUserInfoKey
	// ProgressFileURLKey is a key with a corresponding value that represents the file URL of a file operation for the progress object.
	//
	// See: https://developer.apple.com/documentation/Foundation/ProgressUserInfoKey/fileURLKey
	ProgressFileURLKey NSProgressUserInfoKey
	// ProgressThroughputKey is a key with a corresponding value that indicates the speed of data processing, in bytes per second.
	//
	// See: https://developer.apple.com/documentation/Foundation/ProgressUserInfoKey/throughputKey
	ProgressThroughputKey NSProgressUserInfoKey
)

var (
)

var (
	// ProgressKindFile is the value that indicates that the progress is tracking a file operation.
	//
	// See: https://developer.apple.com/documentation/Foundation/ProgressKind/file
	ProgressKindFile NSProgressKind
)

var (
	// StreamDataWrittenToMemoryStreamKey is value is an [NSData] instance containing the data written to a memory stream.
	//
	// See: https://developer.apple.com/documentation/Foundation/Stream/PropertyKey/dataWrittenToMemoryStreamKey
	StreamDataWrittenToMemoryStreamKey NSStreamPropertyKey
	// StreamFileCurrentOffsetKey is value is an [NSNumber] object containing the current absolute offset of the stream.
	//
	// See: https://developer.apple.com/documentation/Foundation/Stream/PropertyKey/fileCurrentOffsetKey
	StreamFileCurrentOffsetKey NSStreamPropertyKey
	// StreamNetworkServiceType is the type of service for the stream. Providing the service type allows the system to properly handle certain attributes of the stream, including routing and suspension behavior. Most streams do not need to set this property. See `Stream Service Types` for a list of possible values.
	//
	// See: https://developer.apple.com/documentation/Foundation/Stream/PropertyKey/networkServiceType
	StreamNetworkServiceType NSStreamPropertyKey
	// StreamSOCKSProxyConfigurationKey is value is an [NSDictionary] object containing SOCKS proxy configuration information.
	//
	// See: https://developer.apple.com/documentation/Foundation/Stream/PropertyKey/socksProxyConfigurationKey
	StreamSOCKSProxyConfigurationKey NSStreamPropertyKey
	// See: https://developer.apple.com/documentation/Foundation/Stream/PropertyKey/socketSecurityLevelKey
	StreamSocketSecurityLevelKey NSStreamPropertyKey
)

var (
	// StreamNetworkServiceTypeBackground is specifies that the stream is providing a background service.
	//
	// See: https://developer.apple.com/documentation/Foundation/StreamNetworkServiceTypeValue/background
	StreamNetworkServiceTypeBackground NSStreamNetworkServiceTypeValue
	// See: https://developer.apple.com/documentation/Foundation/StreamNetworkServiceTypeValue/callSignaling
	StreamNetworkServiceTypeCallSignaling NSStreamNetworkServiceTypeValue
	// StreamNetworkServiceTypeVideo is specifies that the stream is providing video service.
	//
	// See: https://developer.apple.com/documentation/Foundation/StreamNetworkServiceTypeValue/video
	StreamNetworkServiceTypeVideo NSStreamNetworkServiceTypeValue
	// StreamNetworkServiceTypeVoice is specifies that the stream is providing voice service.
	//
	// See: https://developer.apple.com/documentation/Foundation/StreamNetworkServiceTypeValue/voice
	StreamNetworkServiceTypeVoice NSStreamNetworkServiceTypeValue
)

var (
	// StreamSOCKSProxyHostKey is value is an [NSString] object that represents the SOCKS proxy host.
	//
	// See: https://developer.apple.com/documentation/Foundation/StreamSOCKSProxyConfiguration/hostKey
	StreamSOCKSProxyHostKey NSStreamSOCKSProxyConfiguration
	// StreamSOCKSProxyPasswordKey is value is an [NSString] object containing the user’s password.
	//
	// See: https://developer.apple.com/documentation/Foundation/StreamSOCKSProxyConfiguration/passwordKey
	StreamSOCKSProxyPasswordKey NSStreamSOCKSProxyConfiguration
	// StreamSOCKSProxyPortKey is value is an [NSNumber] object containing an integer that represents the port on which the proxy listens.
	//
	// See: https://developer.apple.com/documentation/Foundation/StreamSOCKSProxyConfiguration/portKey
	StreamSOCKSProxyPortKey NSStreamSOCKSProxyConfiguration
	// StreamSOCKSProxyUserKey is value is an [NSString] object containing the user’s name.
	//
	// See: https://developer.apple.com/documentation/Foundation/StreamSOCKSProxyConfiguration/userKey
	StreamSOCKSProxyUserKey NSStreamSOCKSProxyConfiguration
	// StreamSOCKSProxyVersionKey is value is either [NSStreamSOCKSProxyVersion4] or [NSStreamSOCKSProxyVersion5].
	//
	// See: https://developer.apple.com/documentation/Foundation/StreamSOCKSProxyConfiguration/versionKey
	StreamSOCKSProxyVersionKey NSStreamSOCKSProxyConfiguration
)

var (
	// StreamSOCKSProxyVersion4 is possible value for [NSStreamSOCKSProxyVersionKey].
	//
	// See: https://developer.apple.com/documentation/Foundation/StreamSOCKSProxyVersion/version4
	StreamSOCKSProxyVersion4 NSStreamSOCKSProxyVersion
	// StreamSOCKSProxyVersion5 is possible value for [NSStreamSOCKSProxyVersionKey].
	//
	// See: https://developer.apple.com/documentation/Foundation/StreamSOCKSProxyVersion/version5
	StreamSOCKSProxyVersion5 NSStreamSOCKSProxyVersion
)

var (
)

var (
	// See: https://developer.apple.com/documentation/Foundation/StringEncodingDetectionOptionsKey/allowLossyKey
	StringEncodingDetectionAllowLossyKey NSStringEncodingDetectionOptionsKey
	// See: https://developer.apple.com/documentation/Foundation/StringEncodingDetectionOptionsKey/disallowedEncodingsKey
	StringEncodingDetectionDisallowedEncodingsKey NSStringEncodingDetectionOptionsKey
	// See: https://developer.apple.com/documentation/Foundation/StringEncodingDetectionOptionsKey/fromWindowsKey
	StringEncodingDetectionFromWindowsKey NSStringEncodingDetectionOptionsKey
	// See: https://developer.apple.com/documentation/Foundation/StringEncodingDetectionOptionsKey/likelyLanguageKey
	StringEncodingDetectionLikelyLanguageKey NSStringEncodingDetectionOptionsKey
	// See: https://developer.apple.com/documentation/Foundation/StringEncodingDetectionOptionsKey/lossySubstitutionKey
	StringEncodingDetectionLossySubstitutionKey NSStringEncodingDetectionOptionsKey
	// See: https://developer.apple.com/documentation/Foundation/StringEncodingDetectionOptionsKey/suggestedEncodingsKey
	StringEncodingDetectionSuggestedEncodingsKey NSStringEncodingDetectionOptionsKey
	// See: https://developer.apple.com/documentation/Foundation/StringEncodingDetectionOptionsKey/useOnlySuggestedEncodingsKey
	StringEncodingDetectionUseOnlySuggestedEncodingsKey NSStringEncodingDetectionOptionsKey
)

var (
)

var (
	// TextCheckingAirlineKey is a key that corresponds to the airline of a transit result.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingKey/airline
	TextCheckingAirlineKey NSTextCheckingKey
	// TextCheckingCityKey is a key that corresponds to the city component of the address.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingKey/city
	TextCheckingCityKey NSTextCheckingKey
	// TextCheckingCountryKey is a key that corresponds to the country or region component of the address.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingKey/country
	TextCheckingCountryKey NSTextCheckingKey
	// TextCheckingFlightKey is a key that corresponds to the flight component of a transit result.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingKey/flight
	TextCheckingFlightKey NSTextCheckingKey
	// TextCheckingJobTitleKey is a key that corresponds to the job component of the address.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingKey/jobTitle
	TextCheckingJobTitleKey NSTextCheckingKey
	// TextCheckingNameKey is a key that corresponds to the name component of the address.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingKey/name
	TextCheckingNameKey NSTextCheckingKey
	// TextCheckingOrganizationKey is a key that corresponds to the organization component of the address.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingKey/organization
	TextCheckingOrganizationKey NSTextCheckingKey
	// TextCheckingPhoneKey is a key that corresponds to the phone number component of the address.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingKey/phone
	TextCheckingPhoneKey NSTextCheckingKey
	// TextCheckingStateKey is a key that corresponds to the state or province component of the address.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingKey/state
	TextCheckingStateKey NSTextCheckingKey
	// TextCheckingStreetKey is a key that corresponds to the street address component of the address.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingKey/street
	TextCheckingStreetKey NSTextCheckingKey
	// TextCheckingZIPKey is a key that corresponds to the zip code or postal code component of the address.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingKey/zip
	TextCheckingZIPKey NSTextCheckingKey
)

var (
	// URLAddedToDirectoryDateKey is the time at which the resource’s was created or renamed into or within its parent directory, returned as an [NSDate]. Inconsistent behavior may be observed when this attribute is requested on hard-linked items. This property is not supported by all volumes. (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/addedToDirectoryDateKey
	URLAddedToDirectoryDateKey NSURLResourceKey
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/applicationIsScriptableKey
	URLApplicationIsScriptableKey NSURLResourceKey
	// URLAttributeModificationDateKey is the time at which the resource’s attributes were most recently modified, returned as an [NSDate] object if the volume supports attribute modification dates, or `nil` if attribute modification dates are unsupported (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/attributeModificationDateKey
	URLAttributeModificationDateKey NSURLResourceKey
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/canonicalPathKey
	URLCanonicalPathKey NSURLResourceKey
	// URLContentAccessDateKey is the time at which the resource was most recently accessed.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/contentAccessDateKey
	URLContentAccessDateKey NSURLResourceKey
	// URLContentModificationDateKey is the time at which the resource was most recently modified.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/contentModificationDateKey
	URLContentModificationDateKey NSURLResourceKey
	// URLContentTypeKey is the resource’s type.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/contentTypeKey
	URLContentTypeKey NSURLResourceKey
	// URLCreationDateKey is the time at which the resource was created.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/creationDateKey
	URLCreationDateKey NSURLResourceKey
	// URLCustomIconKey is the icon stored with the resource, returned as an [NSImage] object, or `nil` if the resource has no custom icon.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/customIconKey
	URLCustomIconKey NSURLResourceKey
	// URLDirectoryEntryCountKey is the key for a count of items in the directory.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/directoryEntryCountKey
	URLDirectoryEntryCountKey NSURLResourceKey
	// URLDocumentIdentifierKey is the document identifier returned as an [NSNumber] (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/documentIdentifierKey
	URLDocumentIdentifierKey NSURLResourceKey
	// URLEffectiveIconKey is the resource’s normal icon, returned as an [NSImage] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/effectiveIconKey
	URLEffectiveIconKey NSURLResourceKey
	// URLFileAllocatedSizeKey is the key for the total allocated size on-disk for the file.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/fileAllocatedSizeKey
	URLFileAllocatedSizeKey NSURLResourceKey
	// URLFileContentIdentifierKey is the key for a value that APFS assigns to identify a file’s content data stream.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/fileContentIdentifierKey
	URLFileContentIdentifierKey NSURLResourceKey
	// URLFileIdentifierKey is the key for the file system’s internal inode identifier for the item.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/fileIdentifierKey
	URLFileIdentifierKey NSURLResourceKey
	// URLFileProtectionKey is the key for the protection level of the file.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/fileProtectionKey
	URLFileProtectionKey NSURLResourceKey
	// URLFileResourceIdentifierKey is the key for the resource’s unique identifier.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/fileResourceIdentifierKey
	URLFileResourceIdentifierKey NSURLResourceKey
	// URLFileResourceTypeKey is the key for the resource’s object type.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/fileResourceTypeKey
	URLFileResourceTypeKey NSURLResourceKey
	// URLFileSecurityKey is the key for the resource’s security information.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/fileSecurityKey
	URLFileSecurityKey NSURLResourceKey
	// URLFileSizeKey is the key for the file’s size, in bytes.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/fileSizeKey
	URLFileSizeKey NSURLResourceKey
	// URLGenerationIdentifierKey is an opaque generation identifier, returned as an `id ` (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/generationIdentifierKey
	URLGenerationIdentifierKey NSURLResourceKey
	// URLHasHiddenExtensionKey is key for determining whether the resource’s extension is normally removed from its localized name, returned as a Boolean [NSNumber] object (read-write).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/hasHiddenExtensionKey
	URLHasHiddenExtensionKey NSURLResourceKey
	// URLIsAliasFileKey is the key for determining whether the file is an alias.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/isAliasFileKey
	URLIsAliasFileKey NSURLResourceKey
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/isApplicationKey
	URLIsApplicationKey NSURLResourceKey
	// URLIsDirectoryKey is a key for determining whether the resource is a directory.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/isDirectoryKey
	URLIsDirectoryKey NSURLResourceKey
	// URLIsExcludedFromBackupKey is a key for indicating whether the system excludes the resource from all backups of app data.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/isExcludedFromBackupKey
	URLIsExcludedFromBackupKey NSURLResourceKey
	// URLIsExecutableKey is key for determining whether the current process (as determined by the EUID) can execute the resource (if it is a file) or search the resource (if it is a directory), returned as a Boolean [NSNumber] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/isExecutableKey
	URLIsExecutableKey NSURLResourceKey
	// URLIsHiddenKey is key for determining whether the resource is normally not displayed to users, returned as a Boolean [NSNumber] object (read-write).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/isHiddenKey
	URLIsHiddenKey NSURLResourceKey
	// URLIsMountTriggerKey is key for determining whether the URL is a file system trigger directory, returned as a Boolean [NSNumber] object (read-only). Traversing or opening a file system trigger directory causes an attempt to mount a file system on the directory.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/isMountTriggerKey
	URLIsMountTriggerKey NSURLResourceKey
	// URLIsPackageKey is the key for determining whether the resource is a file package.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/isPackageKey
	URLIsPackageKey NSURLResourceKey
	// URLIsPurgeableKey is the key for a Boolean value that indicates whether the file system can delete a file when the system needs to free space.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/isPurgeableKey
	URLIsPurgeableKey NSURLResourceKey
	// URLIsReadableKey is key for determining whether the current process (as determined by the EUID) can read the resource, returned as a Boolean [NSNumber] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/isReadableKey
	URLIsReadableKey NSURLResourceKey
	// URLIsRegularFileKey is the key for determining whether the resource is a regular file rather than a directory or a symbolic link.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/isRegularFileKey
	URLIsRegularFileKey NSURLResourceKey
	// URLIsSparseKey is the key for a Boolean value that indicates whether the file has sparse regions.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/isSparseKey
	URLIsSparseKey NSURLResourceKey
	// URLIsSymbolicLinkKey is key for determining whether the resource is a symbolic link, returned as a Boolean [NSNumber] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/isSymbolicLinkKey
	URLIsSymbolicLinkKey NSURLResourceKey
	// URLIsSystemImmutableKey is key for determining whether the resource’s system immutable bit is set, returned as a Boolean [NSNumber] object (read-write).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/isSystemImmutableKey
	URLIsSystemImmutableKey NSURLResourceKey
	// URLIsUbiquitousItemKey is the key for a Boolean value that indicates whether the item is in iCloud storage.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/isUbiquitousItemKey
	URLIsUbiquitousItemKey NSURLResourceKey
	// URLIsUserImmutableKey is key for determining whether the resource’s user immutable bit is set, returned as a Boolean [NSNumber] object (read-write).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/isUserImmutableKey
	URLIsUserImmutableKey NSURLResourceKey
	// URLIsVolumeKey is key for determining whether the resource is the root directory of a volume, returned as a Boolean [NSNumber] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/isVolumeKey
	URLIsVolumeKey NSURLResourceKey
	// URLIsWritableKey is key for determining whether the current process (as determined by the EUID) can write to the resource, returned as a Boolean [NSNumber] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/isWritableKey
	URLIsWritableKey NSURLResourceKey
	// URLKeysOfUnsetValuesKey is key for the resource properties that have not been set after the [setResourceValues(_:)] method returns an error, returned as an array of [NSString] objects.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/keysOfUnsetValuesKey
	URLKeysOfUnsetValuesKey NSURLResourceKey
	// URLLabelColorKey is the resource’s label color, returned as an [NSColor] object, or `nil` if the resource has no label color (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/labelColorKey
	URLLabelColorKey NSURLResourceKey
	// URLLabelNumberKey is the resource’s label number, returned as an [NSNumber] object (read-write).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/labelNumberKey
	URLLabelNumberKey NSURLResourceKey
	// URLLinkCountKey is the number of hard links to the resource, returned as an [NSNumber] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/linkCountKey
	URLLinkCountKey NSURLResourceKey
	// URLLocalizedLabelKey is the resource’s localized label text, returned as an [NSString] object, or `nil` if the resource has no localized label text (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/localizedLabelKey
	URLLocalizedLabelKey NSURLResourceKey
	// URLLocalizedNameKey is the resource’s localized or extension-hidden name, returned as an [NSString] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/localizedNameKey
	URLLocalizedNameKey NSURLResourceKey
	// URLLocalizedTypeDescriptionKey is the resource’s localized type description, returned as an [NSString] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/localizedTypeDescriptionKey
	URLLocalizedTypeDescriptionKey NSURLResourceKey
	// URLMayHaveExtendedAttributesKey is the key for a Boolean value that indicates whether the file has extended attributes.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/mayHaveExtendedAttributesKey
	URLMayHaveExtendedAttributesKey NSURLResourceKey
	// URLMayShareFileContentKey is the key for a Boolean value that indicates whether cloned files and their original files may share data blocks.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/mayShareFileContentKey
	URLMayShareFileContentKey NSURLResourceKey
	// URLNameKey is the resource’s name in the file system, returned as an [NSString] object (read-write).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/nameKey
	URLNameKey NSURLResourceKey
	// URLParentDirectoryURLKey is the container directory of the resource.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/parentDirectoryURLKey
	URLParentDirectoryURLKey NSURLResourceKey
	// URLPathKey is the file system path for the URL, returned as an [NSString] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/pathKey
	URLPathKey NSURLResourceKey
	// URLPreferredIOBlockSizeKey is the key for the optimal block size to use when reading or writing the file’s data.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/preferredIOBlockSizeKey
	URLPreferredIOBlockSizeKey NSURLResourceKey
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/quarantinePropertiesKey
	URLQuarantinePropertiesKey NSURLResourceKey
	// URLTagNamesKey is the names of tags attached to the resource, returned as an array of [NSString] values (read-write).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/tagNamesKey
	URLTagNamesKey NSURLResourceKey
	// URLTotalFileAllocatedSizeKey is the key for the total allocated size of the file, in bytes.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/totalFileAllocatedSizeKey
	URLTotalFileAllocatedSizeKey NSURLResourceKey
	// URLTotalFileSizeKey is the key for the total displayable size of the file, in bytes.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/totalFileSizeKey
	URLTotalFileSizeKey NSURLResourceKey
	// URLTypeIdentifierKey is the resource’s uniform type identifier (UTI), returned as an [NSString] object (read-only).
	//
	// Deprecated: Deprecated since macOS 26.4. Use NSURLContentTypeKey instead
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/typeIdentifierKey
	URLTypeIdentifierKey NSURLResourceKey
	// URLUbiquitousItemContainerDisplayNameKey is the key for a string that contains the name of the item’s container as it appears to the user.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/ubiquitousItemContainerDisplayNameKey
	URLUbiquitousItemContainerDisplayNameKey NSURLResourceKey
	// URLUbiquitousItemDownloadRequestedKey is the key for a Boolean value that indicates whether the system has already made a call [startDownloadingUbiquitousItem(at:)] to download the item.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/ubiquitousItemDownloadRequestedKey
	URLUbiquitousItemDownloadRequestedKey NSURLResourceKey
	// URLUbiquitousItemDownloadingErrorKey is the key for an error object that indicates why downloading the item from iCloud fails.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/ubiquitousItemDownloadingErrorKey
	URLUbiquitousItemDownloadingErrorKey NSURLResourceKey
	// URLUbiquitousItemDownloadingStatusKey is the key for the current download state for the item.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/ubiquitousItemDownloadingStatusKey
	URLUbiquitousItemDownloadingStatusKey NSURLResourceKey
	// URLUbiquitousItemHasUnresolvedConflictsKey is the key for a Boolean value that indicates whether this item has outstanding conflicts.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/ubiquitousItemHasUnresolvedConflictsKey
	URLUbiquitousItemHasUnresolvedConflictsKey NSURLResourceKey
	// See: https://developer.apple.com/documentation/foundation/urlresourcekey/ubiquitousitemisdownloadedkey
	NSURLUbiquitousItemIsDownloadedKey NSURLResourceKey
	// URLUbiquitousItemIsDownloadingKey is the key for a Boolean value that indicates whether the system is downloading the item from iCloud.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/ubiquitousItemIsDownloadingKey
	URLUbiquitousItemIsDownloadingKey NSURLResourceKey
	// URLUbiquitousItemIsExcludedFromSyncKey is the key of a Boolean value that indicates whether the system excludes the item from syncing.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/ubiquitousItemIsExcludedFromSyncKey
	URLUbiquitousItemIsExcludedFromSyncKey NSURLResourceKey
	// URLUbiquitousItemIsSharedKey is the key for a Boolean value that indicates a shared item.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/ubiquitousItemIsSharedKey
	URLUbiquitousItemIsSharedKey NSURLResourceKey
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/ubiquitousItemIsSyncPausedKey
	URLUbiquitousItemIsSyncPausedKey NSURLResourceKey
	// URLUbiquitousItemIsUploadedKey is the key for a Boolean value that indicates whether the system uploads the item’s data to iCloud storage.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/ubiquitousItemIsUploadedKey
	URLUbiquitousItemIsUploadedKey NSURLResourceKey
	// URLUbiquitousItemIsUploadingKey is the key for a Boolean value that indicates whether the system is uploading the item to iCloud.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/ubiquitousItemIsUploadingKey
	URLUbiquitousItemIsUploadingKey NSURLResourceKey
	// See: https://developer.apple.com/documentation/foundation/urlresourcekey/ubiquitousitempercentdownloadedkey
	NSURLUbiquitousItemPercentDownloadedKey NSURLResourceKey
	// See: https://developer.apple.com/documentation/foundation/urlresourcekey/ubiquitousitempercentuploadedkey
	NSURLUbiquitousItemPercentUploadedKey NSURLResourceKey
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/ubiquitousItemSupportedSyncControlsKey
	URLUbiquitousItemSupportedSyncControlsKey NSURLResourceKey
	// URLUbiquitousItemUploadingErrorKey is the key for an error object that indicates why uploading the item to iCloud fails.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/ubiquitousItemUploadingErrorKey
	URLUbiquitousItemUploadingErrorKey NSURLResourceKey
	// URLUbiquitousSharedItemCurrentUserPermissionsKey is the key for the current user’s permissions.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/ubiquitousSharedItemCurrentUserPermissionsKey
	URLUbiquitousSharedItemCurrentUserPermissionsKey NSURLResourceKey
	// URLUbiquitousSharedItemCurrentUserRoleKey is the key for the role of the current user.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/ubiquitousSharedItemCurrentUserRoleKey
	URLUbiquitousSharedItemCurrentUserRoleKey NSURLResourceKey
	// URLUbiquitousSharedItemMostRecentEditorNameComponentsKey is the key for the name components of the most recent editor.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/ubiquitousSharedItemMostRecentEditorNameComponentsKey
	URLUbiquitousSharedItemMostRecentEditorNameComponentsKey NSURLResourceKey
	// URLUbiquitousSharedItemOwnerNameComponentsKey is the key for the name components of the item’s owner.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/ubiquitousSharedItemOwnerNameComponentsKey
	URLUbiquitousSharedItemOwnerNameComponentsKey NSURLResourceKey
	// URLVolumeAvailableCapacityForImportantUsageKey is key for the volume’s available capacity in bytes for storing important resources (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeAvailableCapacityForImportantUsageKey
	URLVolumeAvailableCapacityForImportantUsageKey NSURLResourceKey
	// URLVolumeAvailableCapacityForOpportunisticUsageKey is key for the volume’s available capacity in bytes for storing nonessential resources (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeAvailableCapacityForOpportunisticUsageKey
	URLVolumeAvailableCapacityForOpportunisticUsageKey NSURLResourceKey
	// URLVolumeAvailableCapacityKey is key for the volume’s available capacity in bytes (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeAvailableCapacityKey
	URLVolumeAvailableCapacityKey NSURLResourceKey
	// URLVolumeCreationDateKey is key for the volume’s creation date, returned as an [NSDate] object, or [NULL] if it cannot be determined (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeCreationDateKey
	URLVolumeCreationDateKey NSURLResourceKey
	// URLVolumeIdentifierKey is the unique identifier of the resource’s volume, returned as an `id` (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeIdentifierKey
	URLVolumeIdentifierKey NSURLResourceKey
	// URLVolumeIsAutomountedKey is a key for determining whether the volume is automounted.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeIsAutomountedKey
	URLVolumeIsAutomountedKey NSURLResourceKey
	// URLVolumeIsBrowsableKey is a key for determining whether the volume is visible in GUI-based file-browsing environments, such as the Desktop or the Finder app.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeIsBrowsableKey
	URLVolumeIsBrowsableKey NSURLResourceKey
	// URLVolumeIsEjectableKey is a key for determining whether the volume is ejectable from the drive mechanism under software control.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeIsEjectableKey
	URLVolumeIsEjectableKey NSURLResourceKey
	// URLVolumeIsEncryptedKey is a key for determining whether the volume is encrypted.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeIsEncryptedKey
	URLVolumeIsEncryptedKey NSURLResourceKey
	// URLVolumeIsInternalKey is a key for determining whether the volume is connected to an internal bus.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeIsInternalKey
	URLVolumeIsInternalKey NSURLResourceKey
	// URLVolumeIsJournalingKey is a key for determining whether the volume is currently journaling.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeIsJournalingKey
	URLVolumeIsJournalingKey NSURLResourceKey
	// URLVolumeIsLocalKey is a key for determining whether the volume is on a local device.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeIsLocalKey
	URLVolumeIsLocalKey NSURLResourceKey
	// URLVolumeIsReadOnlyKey is a key for determining whether the volume is read-only.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeIsReadOnlyKey
	URLVolumeIsReadOnlyKey NSURLResourceKey
	// URLVolumeIsRemovableKey is a key for determining whether the volume is removable from the drive mechanism.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeIsRemovableKey
	URLVolumeIsRemovableKey NSURLResourceKey
	// URLVolumeIsRootFileSystemKey is a key for determining whether the volume is the root file system.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeIsRootFileSystemKey
	URLVolumeIsRootFileSystemKey NSURLResourceKey
	// URLVolumeLocalizedFormatDescriptionKey is key for the volume’s descriptive format name, returned as an [NSString] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeLocalizedFormatDescriptionKey
	URLVolumeLocalizedFormatDescriptionKey NSURLResourceKey
	// URLVolumeLocalizedNameKey is the name of the volume as it should be displayed in the user interface, returned as an [NSString] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeLocalizedNameKey
	URLVolumeLocalizedNameKey NSURLResourceKey
	// URLVolumeMaximumFileSizeKey is key for the largest file size supported by the volume in bytes, returned as a Boolean [NSNumber] object, or `nil` if it cannot be determined (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeMaximumFileSizeKey
	URLVolumeMaximumFileSizeKey NSURLResourceKey
	// URLVolumeMountFromLocationKey is the key for the volume mounted-from location.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeMountFromLocationKey
	URLVolumeMountFromLocationKey NSURLResourceKey
	// URLVolumeNameKey is the name of the volume, returned as an string object.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeNameKey
	URLVolumeNameKey NSURLResourceKey
	// URLVolumeResourceCountKey is key for the total number of resources on the volume, returned as an [NSNumber] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeResourceCountKey
	URLVolumeResourceCountKey NSURLResourceKey
	// URLVolumeSubtypeKey is the key for the file system subtype value.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeSubtypeKey
	URLVolumeSubtypeKey NSURLResourceKey
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeSupportsAccessPermissionsKey
	URLVolumeSupportsAccessPermissionsKey NSURLResourceKey
	// URLVolumeSupportsAdvisoryFileLockingKey is key for determining whether the volume implements whole-file advisory locks in the style of flock, along with the `O_EXLOCK` and `O_SHLOCK` flags of the open function, returned as a Boolean [NSNumber] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeSupportsAdvisoryFileLockingKey
	URLVolumeSupportsAdvisoryFileLockingKey NSURLResourceKey
	// URLVolumeSupportsCasePreservedNamesKey is key for determining whether the volume supports case-preserved names, returned as a Boolean [NSNumber] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeSupportsCasePreservedNamesKey
	URLVolumeSupportsCasePreservedNamesKey NSURLResourceKey
	// URLVolumeSupportsCaseSensitiveNamesKey is key for determining whether the volume supports case-sensitive names, returned as a Boolean [NSNumber] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeSupportsCaseSensitiveNamesKey
	URLVolumeSupportsCaseSensitiveNamesKey NSURLResourceKey
	// URLVolumeSupportsCompressionKey is whether the volume supports transparent decompression of compressed files using `decmpfs`, returned as [NSNumber] containing a Boolean value (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeSupportsCompressionKey
	URLVolumeSupportsCompressionKey NSURLResourceKey
	// URLVolumeSupportsExclusiveRenamingKey is whether the volume supports exclusive renaming using `renamex_np(2)` with the `RENAME_EXCL` option, returned as [NSNumber] containing a Boolean value (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeSupportsExclusiveRenamingKey
	URLVolumeSupportsExclusiveRenamingKey NSURLResourceKey
	// URLVolumeSupportsExtendedSecurityKey is key for determining whether the volume supports extended security (access control lists), returned as a Boolean [NSNumber] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeSupportsExtendedSecurityKey
	URLVolumeSupportsExtendedSecurityKey NSURLResourceKey
	// URLVolumeSupportsFileCloningKey is whether the volume supports cloning using `clonefile(2)`, returned as [NSNumber] containing a Boolean value (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeSupportsFileCloningKey
	URLVolumeSupportsFileCloningKey NSURLResourceKey
	// URLVolumeSupportsFileProtectionKey is a Boolean value that indicates the volume supports data protection for files.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeSupportsFileProtectionKey
	URLVolumeSupportsFileProtectionKey NSURLResourceKey
	// URLVolumeSupportsHardLinksKey is key for determining whether the volume supports hard links, returned as a Boolean [NSNumber] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeSupportsHardLinksKey
	URLVolumeSupportsHardLinksKey NSURLResourceKey
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeSupportsImmutableFilesKey
	URLVolumeSupportsImmutableFilesKey NSURLResourceKey
	// URLVolumeSupportsJournalingKey is key for determining whether the volume supports journaling, returned as a Boolean [NSNumber] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeSupportsJournalingKey
	URLVolumeSupportsJournalingKey NSURLResourceKey
	// URLVolumeSupportsPersistentIDsKey is key for determining whether the volume supports persistent IDs, returned as a Boolean [NSNumber] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeSupportsPersistentIDsKey
	URLVolumeSupportsPersistentIDsKey NSURLResourceKey
	// URLVolumeSupportsRenamingKey is key for determining whether the volume can be renamed, returned as a Boolean [NSNumber] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeSupportsRenamingKey
	URLVolumeSupportsRenamingKey NSURLResourceKey
	// URLVolumeSupportsRootDirectoryDatesKey is key for determining whether the volume supports reliable storage of times for the root directory, returned as a Boolean [NSNumber] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeSupportsRootDirectoryDatesKey
	URLVolumeSupportsRootDirectoryDatesKey NSURLResourceKey
	// URLVolumeSupportsSparseFilesKey is key for determining whether the volume supports sparse files, returned as a Boolean [NSNumber] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeSupportsSparseFilesKey
	URLVolumeSupportsSparseFilesKey NSURLResourceKey
	// URLVolumeSupportsSwapRenamingKey is whether the volume supports renaming using `renamex_np(2)` with the `RENAME_SWAP` option, returned as [NSNumber] containing a Boolean value (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeSupportsSwapRenamingKey
	URLVolumeSupportsSwapRenamingKey NSURLResourceKey
	// URLVolumeSupportsSymbolicLinksKey is key for determining whether the volume supports symbolic links, returned as a Boolean [NSNumber] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeSupportsSymbolicLinksKey
	URLVolumeSupportsSymbolicLinksKey NSURLResourceKey
	// URLVolumeSupportsVolumeSizesKey is key for determining whether the volume supports returning volume size information, returned as a Boolean [NSNumber] object (read-only). If `true`, volume size information is available as values of the [volumeTotalCapacityKey] and[volumeAvailableCapacityKey] keys.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeSupportsVolumeSizesKey
	URLVolumeSupportsVolumeSizesKey NSURLResourceKey
	// URLVolumeSupportsZeroRunsKey is key for determining whether the volume supports zero runs, returned as a Boolean [NSNumber] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeSupportsZeroRunsKey
	URLVolumeSupportsZeroRunsKey NSURLResourceKey
	// URLVolumeTotalCapacityKey is key for the volume’s total capacity in bytes (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeTotalCapacityKey
	URLVolumeTotalCapacityKey NSURLResourceKey
	// URLVolumeTypeNameKey is the key for the name of the file system type.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeTypeNameKey
	URLVolumeTypeNameKey NSURLResourceKey
	// URLVolumeURLForRemountingKey is key for the URL needed to remount the network volume, returned as an [NSURL] object, or `nil` if a URL is not available (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeURLForRemountingKey
	URLVolumeURLForRemountingKey NSURLResourceKey
	// URLVolumeURLKey is the root directory of the resource’s volume, returned as an [NSURL] object (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeURLKey
	URLVolumeURLKey NSURLResourceKey
	// URLVolumeUUIDStringKey is key for the volume’s persistent UUID, returned as an [NSString] object, or `nil` if a persistent UUID is not available (read-only).
	//
	// See: https://developer.apple.com/documentation/Foundation/URLResourceKey/volumeUUIDStringKey
	URLVolumeUUIDStringKey NSURLResourceKey
)

var (
	// URLFileProtectionComplete is an option that instructs the system to store the file in an encrypted format on-disk that your app can’t access for reading or writing to while the device is locked or booting.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLFileProtection/complete
	URLFileProtectionComplete NSURLFileProtectionType
	// URLFileProtectionCompleteUnlessOpen is an option that instructs the system to store the file in an encrypted format on-disk after it closes.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLFileProtection/completeUnlessOpen
	URLFileProtectionCompleteUnlessOpen NSURLFileProtectionType
	// URLFileProtectionCompleteUntilFirstUserAuthentication is an option that instructs the system to store the file in an encrypted format on-disk that your app can’t access until after the device boots.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLFileProtection/completeUntilFirstUserAuthentication
	URLFileProtectionCompleteUntilFirstUserAuthentication NSURLFileProtectionType
	// See: https://developer.apple.com/documentation/foundation/urlfileprotection/completewhenuserinactive
	NSURLFileProtectionCompleteWhenUserInactive NSURLFileProtectionType
	// URLFileProtectionNone is an option that indicates the file has no special protections associated with it.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLFileProtection/none
	URLFileProtectionNone NSURLFileProtectionType
)

var (
)

var (
	// URLSessionTaskPriorityDefault is the default URL session task priority, used implicitly for any task you have not prioritized.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/defaultPriority
	URLSessionTaskPriorityDefault float32
	// URLSessionTaskPriorityHigh is a high URL session task priority, with a floating point value above the default value and below the maximum of `1.0`.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/highPriority
	URLSessionTaskPriorityHigh float32
	// URLSessionTaskPriorityLow is a low URL session task priority, with a floating point value above the minimum of `0` and below the default value.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/lowPriority
	URLSessionTaskPriorityLow float32
)

var (
	// URLSessionTransferSizeUnknown is the total size of the transfer cannot be determined.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLSessionTransferSizeUnknown
	URLSessionTransferSizeUnknown int64
)

var (
)

var (
)

var (
)

var (
	// See: https://developer.apple.com/documentation/foundation/userdefaults/sizelimitexceedednotification
	NSUserDefaultsSizeLimitExceededNotification NSNotification
)
func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAlternateDescriptionAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AlternateDescriptionAttributeName = NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAppleEventManagerWillProcessFirstEventNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AppleEventManagerWillProcessFirstEventNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAppleEventTimeOutDefault"); err == nil && ptr != 0 {
		AppleEventTimeOutDefault = *(*float64)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAppleEventTimeOutNone"); err == nil && ptr != 0 {
		AppleEventTimeOutNone = *(*float64)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAppleScriptErrorAppName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AppleScriptErrorAppName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAppleScriptErrorBriefMessage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AppleScriptErrorBriefMessage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAppleScriptErrorMessage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AppleScriptErrorMessage = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAppleScriptErrorNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AppleScriptErrorNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAppleScriptErrorRange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AppleScriptErrorRange = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSArgumentDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ArgumentDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAssertionHandlerKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AssertionHandlerKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSAverageKeyValueOperator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AverageKeyValueOperator = NSKeyValueOperator(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSBundleDidLoadNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				BundleDidLoadNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSBundleResourceRequestLowDiskSpaceNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSBundleResourceRequestLowDiskSpaceNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarDayChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CalendarDayChangedNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierBangla"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.Bangla = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierBuddhist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.Buddhist = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierChinese"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.Chinese = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierCoptic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.Coptic = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierDangi"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.Dangi = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierEthiopicAmeteAlem"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.EthiopicAmeteAlem = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierEthiopicAmeteMihret"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.EthiopicAmeteMihret = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierGregorian"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.Gregorian = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierGujarati"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.Gujarati = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierHebrew"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.Hebrew = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierISO8601"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.ISO8601 = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierIndian"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.Indian = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierIslamic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.Islamic = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierIslamicCivil"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.IslamicCivil = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierIslamicTabular"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.IslamicTabular = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierIslamicUmmAlQura"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.IslamicUmmAlQura = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierJapanese"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.Japanese = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierKannada"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.Kannada = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierMalayalam"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.Malayalam = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierMarathi"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.Marathi = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierOdia"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.Odia = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierPersian"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.Persian = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierRepublicOfChina"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.RepublicOfChina = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierTamil"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.Tamil = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierTelugu"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.Telugu = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierVietnamese"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.Vietnamese = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCalendarIdentifierVikram"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSCalendarIdentifiers.Vikram = NSCalendarIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCharacterConversionException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CharacterConversionException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSClassDescriptionNeededForClassNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ClassDescriptionNeededForClassNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCocoaErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CocoaErrorDomain = NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCountKeyValueOperator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CountKeyValueOperator = NSKeyValueOperator(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSCurrentLocaleDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CurrentLocaleDidChangeNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDeallocateZombies"); err == nil && ptr != 0 {
		DeallocateZombies = *(*bool)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDebugDescriptionErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DebugDescriptionErrorKey = NSErrorUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDebugEnabled"); err == nil && ptr != 0 {
		DebugEnabled = *(*bool)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDecimalNumberDivideByZeroException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DecimalNumberDivideByZeroException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDecimalNumberExactnessException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DecimalNumberExactnessException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDecimalNumberOverflowException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DecimalNumberOverflowException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDecimalNumberUnderflowException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DecimalNumberUnderflowException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDefaultRunLoopMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RunLoopDefaultMode = NSRunLoopMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDestinationInvalidException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DestinationInvalidException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDidBecomeSingleThreadedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DidBecomeSingleThreadedNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDistinctUnionOfArraysKeyValueOperator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DistinctUnionOfArraysKeyValueOperator = NSKeyValueOperator(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDistinctUnionOfObjectsKeyValueOperator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DistinctUnionOfObjectsKeyValueOperator = NSKeyValueOperator(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDistinctUnionOfSetsKeyValueOperator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				DistinctUnionOfSetsKeyValueOperator = NSKeyValueOperator(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSEdgeInsetsZero"); err == nil && ptr != 0 {
		EdgeInsetsZero = *(*NSEdgeInsets)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSExtensionHostDidBecomeActiveNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSExtensionHostDidBecomeActiveNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSExtensionHostDidEnterBackgroundNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSExtensionHostDidEnterBackgroundNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSExtensionHostWillEnterForegroundNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSExtensionHostWillEnterForegroundNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSExtensionHostWillResignActiveNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSExtensionHostWillResignActiveNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSExtensionItemAttachmentsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ExtensionItemAttachmentsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSExtensionItemAttributedContentTextKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ExtensionItemAttributedContentTextKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSExtensionItemAttributedTitleKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ExtensionItemAttributedTitleKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSExtensionItemsAndErrorsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ExtensionItemsAndErrorsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSExtensionJavaScriptPreprocessingResultsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ExtensionJavaScriptPreprocessingResultsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileAppendOnly"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileAppendOnly = NSFileAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileBusy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileBusy = NSFileAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileCreationDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileCreationDate = NSFileAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileDeviceIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileDeviceIdentifier = NSFileAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileExtensionHidden"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileExtensionHidden = NSFileAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileGroupOwnerAccountID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileGroupOwnerAccountID = NSFileAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileGroupOwnerAccountName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileGroupOwnerAccountName = NSFileAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileHFSCreatorCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileHFSCreatorCode = NSFileAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileHFSTypeCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileHFSTypeCode = NSFileAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileHandleConnectionAcceptedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileHandleConnectionAcceptedNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileHandleDataAvailableNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileHandleDataAvailableNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileHandleNotificationDataItem"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileHandleNotificationDataItem = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileHandleNotificationFileHandleItem"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileHandleNotificationFileHandleItem = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileHandleOperationException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileHandleOperationException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileHandleReadCompletionNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileHandleReadCompletionNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileHandleReadToEndOfFileCompletionNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileHandleReadToEndOfFileCompletionNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileImmutable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileImmutable = NSFileAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileManagerUnmountDissentingProcessIdentifierErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileManagerUnmountDissentingProcessIdentifierErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileModificationDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileModificationDate = NSFileAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileOwnerAccountID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileOwnerAccountID = NSFileAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileOwnerAccountName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileOwnerAccountName = NSFileAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFilePathErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FilePathErrorKey = NSErrorUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFilePosixPermissions"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FilePosixPermissions = NSFileAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileProtectionComplete"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileProtectionComplete = NSFileProtectionType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileProtectionCompleteUnlessOpen"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileProtectionCompleteUnlessOpen = NSFileProtectionType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileProtectionCompleteUntilFirstUserAuthentication"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileProtectionCompleteUntilFirstUserAuthentication = NSFileProtectionType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileProtectionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileProtectionKey = NSFileAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileProtectionNone"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileProtectionNone = NSFileProtectionType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileReferenceCount"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileReferenceCount = NSFileAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileSize = NSFileAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileSystemFileNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileSystemFileNumber = NSFileAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileSystemFreeNodes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileSystemFreeNodes = NSFileAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileSystemFreeSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileSystemFreeSize = NSFileAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileSystemNodes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileSystemNodes = NSFileAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileSystemNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileSystemNumber = NSFileAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileSystemSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileSystemSize = NSFileAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileType = NSFileAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileTypeBlockSpecial"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileTypeBlockSpecial = NSFileAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileTypeCharacterSpecial"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileTypeCharacterSpecial = NSFileAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileTypeDirectory"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileTypeDirectory = NSFileAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileTypeRegular"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileTypeRegular = NSFileAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileTypeSocket"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileTypeSocket = NSFileAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileTypeSymbolicLink"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileTypeSymbolicLink = NSFileAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFileTypeUnknown"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				FileTypeUnknown = NSFileAttributeType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSFoundationVersionNumber"); err == nil && ptr != 0 {
		FoundationVersionNumber = *(*float64)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSGenericException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GenericException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSGlobalDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GlobalDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSGrammarCorrections"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GrammarCorrections = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSGrammarRange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GrammarRange = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSGrammarUserDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GrammarUserDescription = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSHTTPCookieComment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				HTTPCookieComment = NSHTTPCookiePropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSHTTPCookieCommentURL"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				HTTPCookieCommentURL = NSHTTPCookiePropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSHTTPCookieDiscard"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				HTTPCookieDiscard = NSHTTPCookiePropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSHTTPCookieDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				HTTPCookieDomain = NSHTTPCookiePropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSHTTPCookieExpires"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				HTTPCookieExpires = NSHTTPCookiePropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSHTTPCookieManagerCookiesChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				HTTPCookieManagerCookiesChangedNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSHTTPCookieMaximumAge"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				HTTPCookieMaximumAge = NSHTTPCookiePropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSHTTPCookieName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				HTTPCookieName = NSHTTPCookiePropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSHTTPCookieOriginURL"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				HTTPCookieOriginURL = NSHTTPCookiePropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSHTTPCookiePath"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				HTTPCookiePath = NSHTTPCookiePropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSHTTPCookiePort"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				HTTPCookiePort = NSHTTPCookiePropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSHTTPCookieSameSiteLax"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				HTTPCookieSameSiteLax = NSHTTPCookieStringPolicy(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSHTTPCookieSameSitePolicy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				HTTPCookieSameSitePolicy = NSHTTPCookiePropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSHTTPCookieSameSiteStrict"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				HTTPCookieSameSiteStrict = NSHTTPCookieStringPolicy(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSHTTPCookieSecure"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				HTTPCookieSecure = NSHTTPCookiePropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSHTTPCookieSetByJavaScript"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				HTTPCookieSetByJavaScript = NSHTTPCookiePropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSHTTPCookieValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				HTTPCookieValue = NSHTTPCookiePropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSHTTPCookieVersion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				HTTPCookieVersion = NSHTTPCookiePropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSHelpAnchorErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				HelpAnchorErrorKey = NSErrorUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSImageURLAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ImageURLAttributeName = NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInconsistentArchiveException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InconsistentArchiveException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInflectionAgreementArgumentAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InflectionAgreementArgumentAttributeName = NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInflectionAgreementConceptAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InflectionAgreementConceptAttributeName = NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInflectionAlternativeAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InflectionAlternativeAttributeName = NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInflectionConceptsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InflectionConceptsKey = NSAttributedStringFormattingContextKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInflectionReferentConceptAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InflectionReferentConceptAttributeName = NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInflectionRuleAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InflectionRuleAttributeName = NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInlinePresentationIntentAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InlinePresentationIntentAttributeName = NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSIntegerHashCallBacks"); err == nil && ptr != 0 {
		IntegerHashCallBacks = *(*NSHashTableCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSIntegerMapKeyCallBacks"); err == nil && ptr != 0 {
		IntegerMapKeyCallBacks = *(*NSMapTableKeyCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSIntegerMapValueCallBacks"); err == nil && ptr != 0 {
		IntegerMapValueCallBacks = *(*NSMapTableValueCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInternalInconsistencyException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InternalInconsistencyException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInvalidArchiveOperationException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InvalidArchiveOperationException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInvalidArgumentException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InvalidArgumentException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInvalidReceivePortException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InvalidReceivePortException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInvalidSendPortException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InvalidSendPortException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInvalidUnarchiveOperationException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InvalidUnarchiveOperationException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInvocationOperationCancelledException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InvocationOperationCancelledException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSInvocationOperationVoidResultException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				InvocationOperationVoidResultException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSIsNilTransformerName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IsNilTransformerName = NSValueTransformerName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSIsNotNilTransformerName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IsNotNilTransformerName = NSValueTransformerName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSItemProviderErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ItemProviderErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSItemProviderPreferredImageSizeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ItemProviderPreferredImageSizeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSKeepAllocationStatistics"); err == nil && ptr != 0 {
		KeepAllocationStatistics = *(*bool)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSKeyValueChangeIndexesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KeyValueChangeIndexesKey = NSKeyValueChangeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSKeyValueChangeKindKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KeyValueChangeKindKey = NSKeyValueChangeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSKeyValueChangeNewKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KeyValueChangeNewKey = NSKeyValueChangeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSKeyValueChangeNotificationIsPriorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KeyValueChangeNotificationIsPriorKey = NSKeyValueChangeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSKeyValueChangeOldKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KeyValueChangeOldKey = NSKeyValueChangeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSKeyedArchiveRootObjectKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KeyedArchiveRootObjectKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLanguageIdentifierAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LanguageIdentifierAttributeName = NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagAdjective"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.Adjective = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagAdverb"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.Adverb = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagClassifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.Classifier = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagCloseParenthesis"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.CloseParenthesis = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagCloseQuote"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.CloseQuote = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagConjunction"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.Conjunction = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagDash"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.Dash = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagDeterminer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.Determiner = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagIdiom"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.Idiom = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagInterjection"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.Interjection = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagNoun"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.Noun = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.Number = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagOpenParenthesis"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.OpenParenthesis = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagOpenQuote"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.OpenQuote = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagOrganizationName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.OrganizationName = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagOther"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.Other = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagOtherPunctuation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.OtherPunctuation = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagOtherWhitespace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.OtherWhitespace = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagOtherWord"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.OtherWord = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagParagraphBreak"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.ParagraphBreak = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagParticle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.Particle = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagPersonalName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.PersonalName = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagPlaceName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.PlaceName = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagPreposition"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.Preposition = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagPronoun"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.Pronoun = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagPunctuation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.Punctuation = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagSchemeLanguage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTagSchemes.Language = NSLinguisticTagScheme(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagSchemeLemma"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTagSchemes.Lemma = NSLinguisticTagScheme(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagSchemeLexicalClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTagSchemes.LexicalClass = NSLinguisticTagScheme(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagSchemeNameType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTagSchemes.NameType = NSLinguisticTagScheme(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagSchemeNameTypeOrLexicalClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTagSchemes.NameTypeOrLexicalClass = NSLinguisticTagScheme(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagSchemeScript"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTagSchemes.Script = NSLinguisticTagScheme(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagSchemeTokenType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTagSchemes.TokenType = NSLinguisticTagScheme(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagSentenceTerminator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.SentenceTerminator = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagVerb"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.Verb = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagWhitespace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.Whitespace = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagWord"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.Word = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLinguisticTagWordJoiner"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSLinguisticTags.WordJoiner = NSLinguisticTag(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSListItemDelimiterAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ListItemDelimiterAttributeName = NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLoadedClasses"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LoadedClasses = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocalNotificationCenterType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocalNotificationCenterType = NSDistributedNotificationCenterType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocaleAlternateQuotationBeginDelimiterKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocaleAlternateQuotationBeginDelimiterKey = NSLocaleKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocaleAlternateQuotationEndDelimiterKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocaleAlternateQuotationEndDelimiterKey = NSLocaleKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocaleCalendar"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocaleCalendar = NSLocaleKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocaleCollationIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocaleCollationIdentifier = NSLocaleKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocaleCollatorIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocaleCollatorIdentifier = NSLocaleKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocaleCountryCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocaleCountryCode = NSLocaleKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocaleCurrencyCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocaleCurrencyCode = NSLocaleKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocaleCurrencySymbol"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocaleCurrencySymbol = NSLocaleKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocaleDecimalSeparator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocaleDecimalSeparator = NSLocaleKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocaleExemplarCharacterSet"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocaleExemplarCharacterSet = NSLocaleKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocaleGroupingSeparator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocaleGroupingSeparator = NSLocaleKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocaleIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocaleIdentifier = NSLocaleKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocaleLanguageCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocaleLanguageCode = NSLocaleKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocaleMeasurementSystem"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocaleMeasurementSystem = NSLocaleKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocaleQuotationBeginDelimiterKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocaleQuotationBeginDelimiterKey = NSLocaleKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocaleQuotationEndDelimiterKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocaleQuotationEndDelimiterKey = NSLocaleKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocaleScriptCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocaleScriptCode = NSLocaleKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocaleUsesMetricSystem"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocaleUsesMetricSystem = NSLocaleKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocaleVariantCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocaleVariantCode = NSLocaleKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocalizedDescriptionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocalizedDescriptionKey = NSErrorUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocalizedFailureErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocalizedFailureErrorKey = NSErrorUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocalizedFailureReasonErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocalizedFailureReasonErrorKey = NSErrorUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocalizedNumberFormatAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocalizedNumberFormatAttributeName = NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocalizedRecoveryOptionsErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocalizedRecoveryOptionsErrorKey = NSErrorUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSLocalizedRecoverySuggestionErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				LocalizedRecoverySuggestionErrorKey = NSErrorUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMachErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MachErrorDomain = NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMallocException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MallocException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMarkdownSourcePositionAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MarkdownSourcePositionAttributeName = NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMaximumKeyValueOperator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MaximumKeyValueOperator = NSKeyValueOperator(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemAcquisitionMakeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemAcquisitionMakeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemAcquisitionModelKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemAcquisitionModelKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemAlbumKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemAlbumKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemAltitudeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemAltitudeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemApertureKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemApertureKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemAppleLoopDescriptorsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemAppleLoopDescriptorsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemAppleLoopsKeyFilterTypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemAppleLoopsKeyFilterTypeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemAppleLoopsLoopModeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemAppleLoopsLoopModeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemAppleLoopsRootKeyKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemAppleLoopsRootKeyKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemApplicationCategoriesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemApplicationCategoriesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemAttributeChangeDateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemAttributeChangeDateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemAudiencesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemAudiencesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemAudioBitRateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemAudioBitRateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemAudioChannelCountKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemAudioChannelCountKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemAudioEncodingApplicationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemAudioEncodingApplicationKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemAudioSampleRateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemAudioSampleRateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemAudioTrackNumberKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemAudioTrackNumberKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemAuthorAddressesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemAuthorAddressesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemAuthorEmailAddressesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemAuthorEmailAddressesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemAuthorsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemAuthorsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemBitsPerSampleKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemBitsPerSampleKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemCFBundleIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemCFBundleIdentifierKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemCameraOwnerKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemCameraOwnerKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemCityKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemCityKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemCodecsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemCodecsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemColorSpaceKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemColorSpaceKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemCommentKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemCommentKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemComposerKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemComposerKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemContactKeywordsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemContactKeywordsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemContentCreationDateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemContentCreationDateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemContentModificationDateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemContentModificationDateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemContentTypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemContentTypeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemContentTypeTreeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemContentTypeTreeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemContributorsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemContributorsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemCopyrightKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemCopyrightKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemCountryKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemCountryKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemCoverageKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemCoverageKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemCreatorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemCreatorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemDateAddedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemDateAddedKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemDeliveryTypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemDeliveryTypeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemDescriptionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemDescriptionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemDirectorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemDirectorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemDisplayNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemDisplayNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemDownloadedDateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemDownloadedDateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemDueDateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemDueDateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemDurationSecondsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemDurationSecondsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemEXIFGPSVersionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemEXIFGPSVersionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemEXIFVersionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemEXIFVersionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemEditorsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemEditorsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemEmailAddressesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemEmailAddressesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemEncodingApplicationsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemEncodingApplicationsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemExecutableArchitecturesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemExecutableArchitecturesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemExecutablePlatformKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemExecutablePlatformKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemExposureModeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemExposureModeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemExposureProgramKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemExposureProgramKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemExposureTimeSecondsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemExposureTimeSecondsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemExposureTimeStringKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemExposureTimeStringKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemFNumberKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemFNumberKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemFSContentChangeDateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemFSContentChangeDateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemFSCreationDateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemFSCreationDateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemFSNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemFSNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemFSSizeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemFSSizeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemFinderCommentKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemFinderCommentKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemFlashOnOffKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemFlashOnOffKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemFocalLength35mmKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemFocalLength35mmKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemFocalLengthKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemFocalLengthKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemFontsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemFontsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemGPSAreaInformationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemGPSAreaInformationKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemGPSDOPKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemGPSDOPKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemGPSDateStampKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemGPSDateStampKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemGPSDestBearingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemGPSDestBearingKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemGPSDestDistanceKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemGPSDestDistanceKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemGPSDestLatitudeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemGPSDestLatitudeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemGPSDestLongitudeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemGPSDestLongitudeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemGPSDifferentalKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemGPSDifferentalKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemGPSMapDatumKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemGPSMapDatumKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemGPSMeasureModeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemGPSMeasureModeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemGPSProcessingMethodKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemGPSProcessingMethodKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemGPSStatusKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemGPSStatusKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemGPSTrackKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemGPSTrackKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemGenreKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemGenreKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemHasAlphaChannelKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemHasAlphaChannelKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemHeadlineKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemHeadlineKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemISOSpeedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemISOSpeedKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemIdentifierKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemImageDirectionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemImageDirectionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemInformationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemInformationKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemInstantMessageAddressesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemInstantMessageAddressesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemInstructionsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemInstructionsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemIsApplicationManagedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemIsApplicationManagedKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemIsGeneralMIDISequenceKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemIsGeneralMIDISequenceKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemIsLikelyJunkKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemIsLikelyJunkKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemIsUbiquitousKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemIsUbiquitousKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemKeySignatureKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemKeySignatureKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemKeywordsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemKeywordsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemKindKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemKindKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemLanguagesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemLanguagesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemLastUsedDateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemLastUsedDateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemLatitudeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemLatitudeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemLayerNamesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemLayerNamesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemLensModelKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemLensModelKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemLongitudeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemLongitudeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemLyricistKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemLyricistKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemMaxApertureKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemMaxApertureKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemMediaTypesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemMediaTypesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemMeteringModeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemMeteringModeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemMusicalGenreKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemMusicalGenreKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemMusicalInstrumentCategoryKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemMusicalInstrumentCategoryKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemMusicalInstrumentNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemMusicalInstrumentNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemNamedLocationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemNamedLocationKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemNumberOfPagesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemNumberOfPagesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemOrganizationsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemOrganizationsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemOrientationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemOrientationKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemOriginalFormatKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemOriginalFormatKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemOriginalSourceKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemOriginalSourceKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemPageHeightKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemPageHeightKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemPageWidthKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemPageWidthKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemParticipantsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemParticipantsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemPathKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemPathKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemPerformersKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemPerformersKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemPhoneNumbersKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemPhoneNumbersKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemPixelCountKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemPixelCountKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemPixelHeightKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemPixelHeightKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemPixelWidthKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemPixelWidthKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemProducerKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemProducerKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemProfileNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemProfileNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemProjectsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemProjectsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemPublishersKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemPublishersKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemRecipientAddressesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemRecipientAddressesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemRecipientEmailAddressesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemRecipientEmailAddressesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemRecipientsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemRecipientsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemRecordingDateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemRecordingDateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemRecordingYearKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemRecordingYearKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemRedEyeOnOffKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemRedEyeOnOffKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemResolutionHeightDPIKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemResolutionHeightDPIKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemResolutionWidthDPIKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemResolutionWidthDPIKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemRightsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemRightsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemSecurityMethodKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemSecurityMethodKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemSpeedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemSpeedKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemStarRatingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemStarRatingKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemStateOrProvinceKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemStateOrProvinceKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemStreamableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemStreamableKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemSubjectKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemSubjectKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemTempoKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemTempoKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemTextContentKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemTextContentKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemThemeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemThemeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemTimeSignatureKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemTimeSignatureKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemTimestampKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemTimestampKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemTitleKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemTitleKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemTotalBitRateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemTotalBitRateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemURLKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemURLKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemVersionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemVersionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemVideoBitRateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemVideoBitRateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemWhereFromsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemWhereFromsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataItemWhiteBalanceKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataItemWhiteBalanceKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataQueryAccessibleUbiquitousExternalDocumentsScope"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataQueryAccessibleUbiquitousExternalDocumentsScope = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataQueryDidFinishGatheringNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataQueryDidFinishGatheringNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataQueryDidStartGatheringNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataQueryDidStartGatheringNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataQueryDidUpdateNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataQueryDidUpdateNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataQueryGatheringProgressNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataQueryGatheringProgressNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataQueryIndexedLocalComputerScope"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataQueryIndexedLocalComputerScope = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataQueryIndexedNetworkScope"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataQueryIndexedNetworkScope = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataQueryLocalComputerScope"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataQueryLocalComputerScope = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataQueryNetworkScope"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataQueryNetworkScope = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataQueryResultContentRelevanceAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataQueryResultContentRelevanceAttribute = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataQueryUbiquitousDataScope"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataQueryUbiquitousDataScope = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataQueryUbiquitousDocumentsScope"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataQueryUbiquitousDocumentsScope = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataQueryUpdateAddedItemsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataQueryUpdateAddedItemsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataQueryUpdateChangedItemsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataQueryUpdateChangedItemsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataQueryUpdateRemovedItemsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataQueryUpdateRemovedItemsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataQueryUserHomeScope"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataQueryUserHomeScope = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousItemContainerDisplayNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousItemContainerDisplayNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousItemDownloadRequestedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousItemDownloadRequestedKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousItemDownloadingErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousItemDownloadingErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousItemDownloadingStatusCurrent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousItemDownloadingStatusCurrent = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousItemDownloadingStatusDownloaded"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousItemDownloadingStatusDownloaded = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousItemDownloadingStatusKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousItemDownloadingStatusKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousItemDownloadingStatusNotDownloaded"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousItemDownloadingStatusNotDownloaded = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousItemHasUnresolvedConflictsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousItemHasUnresolvedConflictsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousItemIsDownloadingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousItemIsDownloadingKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousItemIsExternalDocumentKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousItemIsExternalDocumentKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousItemIsSharedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousItemIsSharedKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousItemIsUploadedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousItemIsUploadedKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousItemIsUploadingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousItemIsUploadingKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousItemPercentDownloadedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousItemPercentDownloadedKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousItemPercentUploadedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousItemPercentUploadedKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousItemURLInLocalContainerKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousItemURLInLocalContainerKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousItemUploadingErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousItemUploadingErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousSharedItemCurrentUserPermissionsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousSharedItemCurrentUserPermissionsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousSharedItemCurrentUserRoleKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousSharedItemCurrentUserRoleKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousSharedItemMostRecentEditorNameComponentsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousSharedItemMostRecentEditorNameComponentsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousSharedItemOwnerNameComponentsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousSharedItemOwnerNameComponentsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousSharedItemPermissionsReadOnly"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousSharedItemPermissionsReadOnly = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousSharedItemPermissionsReadWrite"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousSharedItemPermissionsReadWrite = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousSharedItemRoleOwner"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousSharedItemRoleOwner = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMetadataUbiquitousSharedItemRoleParticipant"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MetadataUbiquitousSharedItemRoleParticipant = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMinimumKeyValueOperator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MinimumKeyValueOperator = NSKeyValueOperator(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMorphologyAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MorphologyAttributeName = NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSMultipleUnderlyingErrorsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MultipleUnderlyingErrorsKey = NSErrorUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSNegateBooleanTransformerName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NegateBooleanTransformerName = NSValueTransformerName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSNetServicesErrorCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NetServicesErrorCode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSNetServicesErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NetServicesErrorDomain = NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSNonOwnedPointerHashCallBacks"); err == nil && ptr != 0 {
		NonOwnedPointerHashCallBacks = *(*NSHashTableCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSNonOwnedPointerMapKeyCallBacks"); err == nil && ptr != 0 {
		NonOwnedPointerMapKeyCallBacks = *(*NSMapTableKeyCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSNonOwnedPointerMapValueCallBacks"); err == nil && ptr != 0 {
		NonOwnedPointerMapValueCallBacks = *(*NSMapTableValueCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSNonOwnedPointerOrNullMapKeyCallBacks"); err == nil && ptr != 0 {
		NonOwnedPointerOrNullMapKeyCallBacks = *(*NSMapTableKeyCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSNonRetainedObjectHashCallBacks"); err == nil && ptr != 0 {
		NonRetainedObjectHashCallBacks = *(*NSHashTableCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSNonRetainedObjectMapKeyCallBacks"); err == nil && ptr != 0 {
		NonRetainedObjectMapKeyCallBacks = *(*NSMapTableKeyCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSNonRetainedObjectMapValueCallBacks"); err == nil && ptr != 0 {
		NonRetainedObjectMapValueCallBacks = *(*NSMapTableValueCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSOSStatusErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				OSStatusErrorDomain = NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSObjectHashCallBacks"); err == nil && ptr != 0 {
		ObjectHashCallBacks = *(*NSHashTableCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSObjectInaccessibleException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ObjectInaccessibleException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSObjectMapKeyCallBacks"); err == nil && ptr != 0 {
		ObjectMapKeyCallBacks = *(*NSMapTableKeyCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSObjectMapValueCallBacks"); err == nil && ptr != 0 {
		ObjectMapValueCallBacks = *(*NSMapTableValueCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSObjectNotAvailableException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ObjectNotAvailableException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSOldStyleException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				OldStyleException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSOperationNotSupportedForKeyException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				OperationNotSupportedForKeyException = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSOwnedObjectIdentityHashCallBacks"); err == nil && ptr != 0 {
		OwnedObjectIdentityHashCallBacks = *(*NSHashTableCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSOwnedPointerHashCallBacks"); err == nil && ptr != 0 {
		OwnedPointerHashCallBacks = *(*NSHashTableCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSOwnedPointerMapKeyCallBacks"); err == nil && ptr != 0 {
		OwnedPointerMapKeyCallBacks = *(*NSMapTableKeyCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSOwnedPointerMapValueCallBacks"); err == nil && ptr != 0 {
		OwnedPointerMapValueCallBacks = *(*NSMapTableValueCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPOSIXErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				POSIXErrorDomain = NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSParseErrorException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ParseErrorException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersonNameComponentDelimiter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersonNameComponentDelimiter = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersonNameComponentFamilyName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersonNameComponentFamilyName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersonNameComponentGivenName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersonNameComponentGivenName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersonNameComponentKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersonNameComponentKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersonNameComponentMiddleName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersonNameComponentMiddleName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersonNameComponentNickname"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersonNameComponentNickname = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersonNameComponentPrefix"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersonNameComponentPrefix = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPersonNameComponentSuffix"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PersonNameComponentSuffix = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPointerToStructHashCallBacks"); err == nil && ptr != 0 {
		PointerToStructHashCallBacks = *(*NSHashTableCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPortDidBecomeInvalidNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PortDidBecomeInvalidNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPortReceiveException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PortReceiveException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPortSendException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PortSendException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPortTimeoutException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PortTimeoutException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSPresentationIntentAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PresentationIntentAttributeName = NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSProcessInfoPowerStateDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ProcessInfoPowerStateDidChangeNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSProcessInfoThermalStateDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ProcessInfoThermalStateDidChangeNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSProgressEstimatedTimeRemainingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ProgressEstimatedTimeRemainingKey = NSProgressUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSProgressFileAnimationImageKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ProgressFileAnimationImageKey = NSProgressUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSProgressFileAnimationImageOriginalRectKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ProgressFileAnimationImageOriginalRectKey = NSProgressUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSProgressFileCompletedCountKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ProgressFileCompletedCountKey = NSProgressUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSProgressFileIconKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ProgressFileIconKey = NSProgressUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSProgressFileOperationKindCopying"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSProgressFileOperationKinds.Copying = NSProgressFileOperationKind(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSProgressFileOperationKindDecompressingAfterDownloading"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSProgressFileOperationKinds.DecompressingAfterDownloading = NSProgressFileOperationKind(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSProgressFileOperationKindDownloading"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSProgressFileOperationKinds.Downloading = NSProgressFileOperationKind(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSProgressFileOperationKindDuplicating"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSProgressFileOperationKinds.Duplicating = NSProgressFileOperationKind(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSProgressFileOperationKindKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ProgressFileOperationKindKey = NSProgressUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSProgressFileOperationKindReceiving"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSProgressFileOperationKinds.Receiving = NSProgressFileOperationKind(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSProgressFileOperationKindUploading"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSProgressFileOperationKinds.Uploading = NSProgressFileOperationKind(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSProgressFileTotalCountKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ProgressFileTotalCountKey = NSProgressUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSProgressFileURLKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ProgressFileURLKey = NSProgressUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSProgressKindFile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ProgressKindFile = NSProgressKind(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSProgressThroughputKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ProgressThroughputKey = NSProgressUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRangeException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RangeException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRecoveryAttempterErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RecoveryAttempterErrorKey = NSErrorUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRegistrationDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RegistrationDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSReplacementIndexAttributeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ReplacementIndexAttributeName = NSAttributedStringKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSRunLoopCommonModes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				RunLoopCommonModes = NSRunLoopMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSecureUnarchiveFromDataTransformerName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SecureUnarchiveFromDataTransformerName = NSValueTransformerName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStreamDataWrittenToMemoryStreamKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StreamDataWrittenToMemoryStreamKey = NSStreamPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStreamFileCurrentOffsetKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StreamFileCurrentOffsetKey = NSStreamPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStreamNetworkServiceType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StreamNetworkServiceType = NSStreamPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStreamNetworkServiceTypeBackground"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StreamNetworkServiceTypeBackground = NSStreamNetworkServiceTypeValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStreamNetworkServiceTypeCallSignaling"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StreamNetworkServiceTypeCallSignaling = NSStreamNetworkServiceTypeValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStreamNetworkServiceTypeVideo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StreamNetworkServiceTypeVideo = NSStreamNetworkServiceTypeValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStreamNetworkServiceTypeVoice"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StreamNetworkServiceTypeVoice = NSStreamNetworkServiceTypeValue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStreamSOCKSErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StreamSOCKSErrorDomain = NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStreamSOCKSProxyConfigurationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StreamSOCKSProxyConfigurationKey = NSStreamPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStreamSOCKSProxyHostKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StreamSOCKSProxyHostKey = NSStreamSOCKSProxyConfiguration(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStreamSOCKSProxyPasswordKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StreamSOCKSProxyPasswordKey = NSStreamSOCKSProxyConfiguration(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStreamSOCKSProxyPortKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StreamSOCKSProxyPortKey = NSStreamSOCKSProxyConfiguration(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStreamSOCKSProxyUserKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StreamSOCKSProxyUserKey = NSStreamSOCKSProxyConfiguration(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStreamSOCKSProxyVersion4"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StreamSOCKSProxyVersion4 = NSStreamSOCKSProxyVersion(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStreamSOCKSProxyVersion5"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StreamSOCKSProxyVersion5 = NSStreamSOCKSProxyVersion(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStreamSOCKSProxyVersionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StreamSOCKSProxyVersionKey = NSStreamSOCKSProxyConfiguration(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStreamSocketSSLErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StreamSocketSSLErrorDomain = NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStreamSocketSecurityLevelKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StreamSocketSecurityLevelKey = NSStreamPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStreamSocketSecurityLevelNegotiatedSSL"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StreamSocketSecurityLevels.NegotiatedSSL = NSStreamSocketSecurityLevel(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStreamSocketSecurityLevelNone"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StreamSocketSecurityLevels.None = NSStreamSocketSecurityLevel(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStreamSocketSecurityLevelSSLv2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StreamSocketSecurityLevels.SSLv2 = NSStreamSocketSecurityLevel(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStreamSocketSecurityLevelSSLv3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StreamSocketSecurityLevels.SSLv3 = NSStreamSocketSecurityLevel(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStreamSocketSecurityLevelTLSv1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StreamSocketSecurityLevels.TLSv1 = NSStreamSocketSecurityLevel(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStringEncodingDetectionAllowLossyKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StringEncodingDetectionAllowLossyKey = NSStringEncodingDetectionOptionsKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStringEncodingDetectionDisallowedEncodingsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StringEncodingDetectionDisallowedEncodingsKey = NSStringEncodingDetectionOptionsKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStringEncodingDetectionFromWindowsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StringEncodingDetectionFromWindowsKey = NSStringEncodingDetectionOptionsKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStringEncodingDetectionLikelyLanguageKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StringEncodingDetectionLikelyLanguageKey = NSStringEncodingDetectionOptionsKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStringEncodingDetectionLossySubstitutionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StringEncodingDetectionLossySubstitutionKey = NSStringEncodingDetectionOptionsKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStringEncodingDetectionSuggestedEncodingsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StringEncodingDetectionSuggestedEncodingsKey = NSStringEncodingDetectionOptionsKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStringEncodingDetectionUseOnlySuggestedEncodingsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StringEncodingDetectionUseOnlySuggestedEncodingsKey = NSStringEncodingDetectionOptionsKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStringEncodingErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StringEncodingErrorKey = NSErrorUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStringTransformFullwidthToHalfwidth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StringTransforms.FullwidthToHalfwidth = NSStringTransform(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStringTransformHiraganaToKatakana"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StringTransforms.HiraganaToKatakana = NSStringTransform(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStringTransformLatinToArabic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StringTransforms.LatinToArabic = NSStringTransform(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStringTransformLatinToCyrillic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StringTransforms.LatinToCyrillic = NSStringTransform(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStringTransformLatinToGreek"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StringTransforms.LatinToGreek = NSStringTransform(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStringTransformLatinToHangul"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StringTransforms.LatinToHangul = NSStringTransform(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStringTransformLatinToHebrew"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StringTransforms.LatinToHebrew = NSStringTransform(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStringTransformLatinToHiragana"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StringTransforms.LatinToHiragana = NSStringTransform(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStringTransformLatinToKatakana"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StringTransforms.LatinToKatakana = NSStringTransform(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStringTransformLatinToThai"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StringTransforms.LatinToThai = NSStringTransform(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStringTransformMandarinToLatin"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StringTransforms.MandarinToLatin = NSStringTransform(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStringTransformStripCombiningMarks"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StringTransforms.StripCombiningMarks = NSStringTransform(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStringTransformStripDiacritics"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StringTransforms.StripDiacritics = NSStringTransform(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStringTransformToLatin"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StringTransforms.ToLatin = NSStringTransform(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStringTransformToUnicodeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StringTransforms.ToUnicodeName = NSStringTransform(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSStringTransformToXMLHex"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				StringTransforms.ToXMLHex = NSStringTransform(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSumKeyValueOperator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SumKeyValueOperator = NSKeyValueOperator(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSystemClockDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SystemClockDidChangeNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSSystemTimeZoneDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SystemTimeZoneDidChangeNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTaskDidTerminateNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TaskDidTerminateNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextCheckingAirlineKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextCheckingAirlineKey = NSTextCheckingKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextCheckingCityKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextCheckingCityKey = NSTextCheckingKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextCheckingCountryKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextCheckingCountryKey = NSTextCheckingKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextCheckingFlightKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextCheckingFlightKey = NSTextCheckingKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextCheckingJobTitleKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextCheckingJobTitleKey = NSTextCheckingKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextCheckingNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextCheckingNameKey = NSTextCheckingKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextCheckingOrganizationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextCheckingOrganizationKey = NSTextCheckingKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextCheckingPhoneKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextCheckingPhoneKey = NSTextCheckingKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextCheckingStateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextCheckingStateKey = NSTextCheckingKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextCheckingStreetKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextCheckingStreetKey = NSTextCheckingKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSTextCheckingZIPKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				TextCheckingZIPKey = NSTextCheckingKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSThreadWillExitNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ThreadWillExitNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLAddedToDirectoryDateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLAddedToDirectoryDateKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLApplicationIsScriptableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLApplicationIsScriptableKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLAttributeModificationDateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLAttributeModificationDateKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLAuthenticationMethodClientCertificate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLAuthenticationMethodClientCertificate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLAuthenticationMethodDefault"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLAuthenticationMethodDefault = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLAuthenticationMethodHTMLForm"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLAuthenticationMethodHTMLForm = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLAuthenticationMethodHTTPBasic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLAuthenticationMethodHTTPBasic = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLAuthenticationMethodHTTPDigest"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLAuthenticationMethodHTTPDigest = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLAuthenticationMethodNTLM"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLAuthenticationMethodNTLM = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLAuthenticationMethodNegotiate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLAuthenticationMethodNegotiate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLAuthenticationMethodServerTrust"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLAuthenticationMethodServerTrust = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLCanonicalPathKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLCanonicalPathKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLContentAccessDateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLContentAccessDateKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLContentModificationDateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLContentModificationDateKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLContentTypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLContentTypeKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLCreationDateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLCreationDateKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLCredentialStorageRemoveSynchronizableCredentials"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLCredentialStorageRemoveSynchronizableCredentials = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLCustomIconKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLCustomIconKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLDirectoryEntryCountKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLDirectoryEntryCountKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLDocumentIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLDocumentIdentifierKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLEffectiveIconKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLEffectiveIconKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLErrorBackgroundTaskCancelledReasonKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLErrorBackgroundTaskCancelledReasonKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLErrorDomain = NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLErrorFailingURLErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLErrorFailingURLErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLErrorFailingURLPeerTrustErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLErrorFailingURLPeerTrustErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLErrorFailingURLStringErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLErrorFailingURLStringErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLErrorKey = NSErrorUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLErrorNetworkUnavailableReasonKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLErrorNetworkUnavailableReasonKey = NSErrorUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLFileAllocatedSizeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLFileAllocatedSizeKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLFileContentIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLFileContentIdentifierKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLFileIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLFileIdentifierKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLFileProtectionComplete"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLFileProtectionComplete = NSURLFileProtectionType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLFileProtectionCompleteUnlessOpen"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLFileProtectionCompleteUnlessOpen = NSURLFileProtectionType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLFileProtectionCompleteUntilFirstUserAuthentication"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLFileProtectionCompleteUntilFirstUserAuthentication = NSURLFileProtectionType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLFileProtectionCompleteWhenUserInactive"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSURLFileProtectionCompleteWhenUserInactive = NSURLFileProtectionType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLFileProtectionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLFileProtectionKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLFileProtectionNone"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLFileProtectionNone = NSURLFileProtectionType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLFileResourceIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLFileResourceIdentifierKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLFileResourceTypeBlockSpecial"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLFileResourceTypes.BlockSpecial = NSURLFileResourceType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLFileResourceTypeCharacterSpecial"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLFileResourceTypes.CharacterSpecial = NSURLFileResourceType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLFileResourceTypeDirectory"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLFileResourceTypes.Directory = NSURLFileResourceType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLFileResourceTypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLFileResourceTypeKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLFileResourceTypeNamedPipe"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLFileResourceTypes.NamedPipe = NSURLFileResourceType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLFileResourceTypeRegular"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLFileResourceTypes.Regular = NSURLFileResourceType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLFileResourceTypeSocket"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLFileResourceTypes.Socket = NSURLFileResourceType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLFileResourceTypeSymbolicLink"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLFileResourceTypes.SymbolicLink = NSURLFileResourceType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLFileResourceTypeUnknown"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLFileResourceTypes.Unknown = NSURLFileResourceType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLFileScheme"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLFileScheme = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLFileSecurityKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLFileSecurityKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLFileSizeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLFileSizeKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLGenerationIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLGenerationIdentifierKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLHasHiddenExtensionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLHasHiddenExtensionKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLIsAliasFileKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLIsAliasFileKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLIsApplicationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLIsApplicationKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLIsDirectoryKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLIsDirectoryKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLIsExcludedFromBackupKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLIsExcludedFromBackupKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLIsExecutableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLIsExecutableKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLIsHiddenKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLIsHiddenKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLIsMountTriggerKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLIsMountTriggerKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLIsPackageKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLIsPackageKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLIsPurgeableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLIsPurgeableKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLIsReadableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLIsReadableKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLIsRegularFileKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLIsRegularFileKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLIsSparseKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLIsSparseKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLIsSymbolicLinkKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLIsSymbolicLinkKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLIsSystemImmutableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLIsSystemImmutableKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLIsUbiquitousItemKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLIsUbiquitousItemKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLIsUserImmutableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLIsUserImmutableKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLIsVolumeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLIsVolumeKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLIsWritableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLIsWritableKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLKeysOfUnsetValuesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLKeysOfUnsetValuesKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLLabelColorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLLabelColorKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLLabelNumberKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLLabelNumberKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLLinkCountKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLLinkCountKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLLocalizedLabelKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLLocalizedLabelKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLLocalizedNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLLocalizedNameKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLLocalizedTypeDescriptionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLLocalizedTypeDescriptionKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLMayHaveExtendedAttributesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLMayHaveExtendedAttributesKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLMayShareFileContentKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLMayShareFileContentKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLNameKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLParentDirectoryURLKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLParentDirectoryURLKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLPathKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLPathKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLPreferredIOBlockSizeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLPreferredIOBlockSizeKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLProtectionSpaceFTP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLProtectionSpaceFTP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLProtectionSpaceFTPProxy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLProtectionSpaceFTPProxy = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLProtectionSpaceHTTP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLProtectionSpaceHTTP = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLProtectionSpaceHTTPProxy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLProtectionSpaceHTTPProxy = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLProtectionSpaceHTTPS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLProtectionSpaceHTTPS = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLProtectionSpaceHTTPSProxy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLProtectionSpaceHTTPSProxy = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLProtectionSpaceSOCKSProxy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLProtectionSpaceSOCKSProxy = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLQuarantinePropertiesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLQuarantinePropertiesKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLSessionDownloadTaskResumeData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLSessionDownloadTaskResumeData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLSessionTaskPriorityDefault"); err == nil && ptr != 0 {
		URLSessionTaskPriorityDefault = *(*float32)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLSessionTaskPriorityHigh"); err == nil && ptr != 0 {
		URLSessionTaskPriorityHigh = *(*float32)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLSessionTaskPriorityLow"); err == nil && ptr != 0 {
		URLSessionTaskPriorityLow = *(*float32)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLSessionTransferSizeUnknown"); err == nil && ptr != 0 {
		URLSessionTransferSizeUnknown = *(*int64)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLSessionUploadTaskResumeData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLSessionUploadTaskResumeData = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLTagNamesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLTagNamesKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLTotalFileAllocatedSizeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLTotalFileAllocatedSizeKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLTotalFileSizeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLTotalFileSizeKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLTypeIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLTypeIdentifierKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousItemContainerDisplayNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLUbiquitousItemContainerDisplayNameKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousItemDownloadRequestedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLUbiquitousItemDownloadRequestedKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousItemDownloadingErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLUbiquitousItemDownloadingErrorKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousItemDownloadingStatusCurrent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLUbiquitousItemDownloadingStatuss.Current = NSURLUbiquitousItemDownloadingStatus(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousItemDownloadingStatusDownloaded"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLUbiquitousItemDownloadingStatuss.Downloaded = NSURLUbiquitousItemDownloadingStatus(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousItemDownloadingStatusKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLUbiquitousItemDownloadingStatusKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousItemDownloadingStatusNotDownloaded"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLUbiquitousItemDownloadingStatuss.NotDownloaded = NSURLUbiquitousItemDownloadingStatus(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousItemHasUnresolvedConflictsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLUbiquitousItemHasUnresolvedConflictsKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousItemIsDownloadedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSURLUbiquitousItemIsDownloadedKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousItemIsDownloadingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLUbiquitousItemIsDownloadingKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousItemIsExcludedFromSyncKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLUbiquitousItemIsExcludedFromSyncKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousItemIsSharedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLUbiquitousItemIsSharedKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousItemIsSyncPausedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLUbiquitousItemIsSyncPausedKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousItemIsUploadedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLUbiquitousItemIsUploadedKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousItemIsUploadingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLUbiquitousItemIsUploadingKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousItemPercentDownloadedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSURLUbiquitousItemPercentDownloadedKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousItemPercentUploadedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NSURLUbiquitousItemPercentUploadedKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousItemSupportedSyncControlsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLUbiquitousItemSupportedSyncControlsKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousItemUploadingErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLUbiquitousItemUploadingErrorKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousSharedItemCurrentUserPermissionsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLUbiquitousSharedItemCurrentUserPermissionsKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousSharedItemCurrentUserRoleKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLUbiquitousSharedItemCurrentUserRoleKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousSharedItemMostRecentEditorNameComponentsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLUbiquitousSharedItemMostRecentEditorNameComponentsKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousSharedItemOwnerNameComponentsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLUbiquitousSharedItemOwnerNameComponentsKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousSharedItemPermissionsReadOnly"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLUbiquitousSharedItemPermissionss.ReadOnly = NSURLUbiquitousSharedItemPermissions(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousSharedItemPermissionsReadWrite"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLUbiquitousSharedItemPermissionss.ReadWrite = NSURLUbiquitousSharedItemPermissions(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousSharedItemRoleOwner"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLUbiquitousSharedItemRoles.Owner = NSURLUbiquitousSharedItemRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLUbiquitousSharedItemRoleParticipant"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLUbiquitousSharedItemRoles.Participant = NSURLUbiquitousSharedItemRole(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeAvailableCapacityForImportantUsageKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeAvailableCapacityForImportantUsageKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeAvailableCapacityForOpportunisticUsageKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeAvailableCapacityForOpportunisticUsageKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeAvailableCapacityKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeAvailableCapacityKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeCreationDateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeCreationDateKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeIdentifierKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeIsAutomountedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeIsAutomountedKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeIsBrowsableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeIsBrowsableKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeIsEjectableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeIsEjectableKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeIsEncryptedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeIsEncryptedKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeIsInternalKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeIsInternalKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeIsJournalingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeIsJournalingKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeIsLocalKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeIsLocalKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeIsReadOnlyKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeIsReadOnlyKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeIsRemovableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeIsRemovableKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeIsRootFileSystemKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeIsRootFileSystemKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeLocalizedFormatDescriptionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeLocalizedFormatDescriptionKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeLocalizedNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeLocalizedNameKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeMaximumFileSizeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeMaximumFileSizeKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeMountFromLocationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeMountFromLocationKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeNameKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeResourceCountKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeResourceCountKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeSubtypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeSubtypeKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeSupportsAccessPermissionsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeSupportsAccessPermissionsKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeSupportsAdvisoryFileLockingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeSupportsAdvisoryFileLockingKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeSupportsCasePreservedNamesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeSupportsCasePreservedNamesKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeSupportsCaseSensitiveNamesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeSupportsCaseSensitiveNamesKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeSupportsCompressionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeSupportsCompressionKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeSupportsExclusiveRenamingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeSupportsExclusiveRenamingKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeSupportsExtendedSecurityKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeSupportsExtendedSecurityKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeSupportsFileCloningKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeSupportsFileCloningKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeSupportsFileProtectionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeSupportsFileProtectionKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeSupportsHardLinksKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeSupportsHardLinksKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeSupportsImmutableFilesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeSupportsImmutableFilesKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeSupportsJournalingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeSupportsJournalingKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeSupportsPersistentIDsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeSupportsPersistentIDsKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeSupportsRenamingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeSupportsRenamingKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeSupportsRootDirectoryDatesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeSupportsRootDirectoryDatesKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeSupportsSparseFilesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeSupportsSparseFilesKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeSupportsSwapRenamingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeSupportsSwapRenamingKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeSupportsSymbolicLinksKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeSupportsSymbolicLinksKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeSupportsVolumeSizesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeSupportsVolumeSizesKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeSupportsZeroRunsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeSupportsZeroRunsKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeTotalCapacityKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeTotalCapacityKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeTypeNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeTypeNameKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeURLForRemountingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeURLForRemountingKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeURLKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeURLKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSURLVolumeUUIDStringKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				URLVolumeUUIDStringKey = NSURLResourceKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUbiquitousKeyValueStoreChangeReasonKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UbiquitousKeyValueStoreChangeReasonKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUbiquitousKeyValueStoreChangedKeysKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UbiquitousKeyValueStoreChangedKeysKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUbiquitousKeyValueStoreDidChangeExternallyNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UbiquitousKeyValueStoreDidChangeExternallyNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUbiquityIdentityDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UbiquityIdentityDidChangeNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUndefinedKeyException"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UndefinedKeyException = NSExceptionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUnderlyingErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UnderlyingErrorKey = NSErrorUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUndoManagerCheckpointNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UndoManagerCheckpointNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUndoManagerDidCloseUndoGroupNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UndoManagerDidCloseUndoGroupNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUndoManagerDidOpenUndoGroupNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UndoManagerDidOpenUndoGroupNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUndoManagerDidRedoChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UndoManagerDidRedoChangeNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUndoManagerDidUndoChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UndoManagerDidUndoChangeNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUndoManagerGroupIsDiscardableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UndoManagerGroupIsDiscardableKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUndoManagerWillCloseUndoGroupNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UndoManagerWillCloseUndoGroupNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUndoManagerWillRedoChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UndoManagerWillRedoChangeNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUndoManagerWillUndoChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UndoManagerWillUndoChangeNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUnionOfArraysKeyValueOperator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UnionOfArraysKeyValueOperator = NSKeyValueOperator(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUnionOfObjectsKeyValueOperator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UnionOfObjectsKeyValueOperator = NSKeyValueOperator(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUnionOfSetsKeyValueOperator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UnionOfSetsKeyValueOperator = NSKeyValueOperator(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUserActivityTypeBrowsingWeb"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UserActivityTypeBrowsingWeb = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUserDefaultsDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UserDefaultsDidChangeNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSUserDefaultsSizeLimitExceededNotification"); err == nil && ptr != 0 {
		NSUserDefaultsSizeLimitExceededNotification = *(*NSNotification)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSWillBecomeMultiThreadedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				WillBecomeMultiThreadedNotification = NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSXMLParserErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				XMLParserErrorDomain = NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSZombieEnabled"); err == nil && ptr != 0 {
		ZombieEnabled = *(*bool)(unsafe.Pointer(ptr))
	}

}

// NSCalendarIdentifiers provides typed accessors for [NSCalendarIdentifier] constants.
var NSCalendarIdentifiers struct {
	Bangla NSCalendarIdentifier
	// Buddhist: Identifier for the Buddhist calendar.
	Buddhist NSCalendarIdentifier
	// Chinese: Identifier for the Chinese calendar.
	Chinese NSCalendarIdentifier
	// Coptic: Identifier for the Coptic calendar.
	Coptic NSCalendarIdentifier
	Dangi NSCalendarIdentifier
	// EthiopicAmeteAlem: Identifier for the Ethiopic (Amete Alem) calendar.
	EthiopicAmeteAlem NSCalendarIdentifier
	// EthiopicAmeteMihret: Identifier for the Ethiopic (Amete Mihret) calendar.
	EthiopicAmeteMihret NSCalendarIdentifier
	// Gregorian: Identifier for the Gregorian calendar.
	Gregorian NSCalendarIdentifier
	Gujarati NSCalendarIdentifier
	// Hebrew: Identifier for the Hebrew calendar.
	Hebrew NSCalendarIdentifier
	// ISO8601: Identifier for the ISO8601 calendar.
	ISO8601 NSCalendarIdentifier
	// Indian: Identifier for the Indian calendar.
	Indian NSCalendarIdentifier
	// Islamic: Identifier for the Islamic calendar.
	Islamic NSCalendarIdentifier
	// IslamicCivil: Identifier for the Islamic civil calendar.
	IslamicCivil NSCalendarIdentifier
	// IslamicTabular: Identifier for a tabular Islamic calendar.
	IslamicTabular NSCalendarIdentifier
	// IslamicUmmAlQura: Identifier for the Islamic Umm al-Qura calendar.
	IslamicUmmAlQura NSCalendarIdentifier
	// Japanese: Identifier for the Japanese calendar.
	Japanese NSCalendarIdentifier
	Kannada NSCalendarIdentifier
	Malayalam NSCalendarIdentifier
	Marathi NSCalendarIdentifier
	Odia NSCalendarIdentifier
	// Persian: Identifier for the Persian calendar.
	Persian NSCalendarIdentifier
	// RepublicOfChina: Identifier for the Republic of China calendar.
	RepublicOfChina NSCalendarIdentifier
	Tamil NSCalendarIdentifier
	Telugu NSCalendarIdentifier
	Vietnamese NSCalendarIdentifier
	Vikram NSCalendarIdentifier
}

// NSLinguisticTagSchemes provides typed accessors for [NSLinguisticTagScheme] constants.
var NSLinguisticTagSchemes struct {
	// Language: Supplies the language for a token, if one can be determined.
	Language NSLinguisticTagScheme
	// Lemma: Supplies a stem form of a word token, if known.
	Lemma NSLinguisticTagScheme
	// LexicalClass: Classifies tokens according to class:  part of speech, type of punctuation, or whitespace.
	LexicalClass NSLinguisticTagScheme
	// NameType: Classifies tokens according to whether they are part of a named entity.
	NameType NSLinguisticTagScheme
	// NameTypeOrLexicalClass: Classifies tokens corresponding to names according to [nameType](<doc://com.apple.foundation/documentation/Foundation/NSLinguisticTagScheme/nameType>), and classifies all other tokens according to [lexicalClass](<doc://com.apple.foundation/documentation/Foundation/NSLinguisticTagScheme/lexicalClass>).
	NameTypeOrLexicalClass NSLinguisticTagScheme
	// Script: Supplies the script for a token, if one can be determined.
	Script NSLinguisticTagScheme
	// TokenType: Classifies tokens according to their broad type:  word, punctuation, or whitespace.
	TokenType NSLinguisticTagScheme
}

// NSLinguisticTags provides typed accessors for [NSLinguisticTag] constants.
var NSLinguisticTags struct {
	// Adjective: This token is an adjective
	Adjective NSLinguisticTag
	// Adverb: This token is an adverb.
	Adverb NSLinguisticTag
	// Classifier: This token is a classifier.
	Classifier NSLinguisticTag
	// CloseParenthesis: This token is a close parenthesis.
	CloseParenthesis NSLinguisticTag
	// CloseQuote: This token is a close quote.
	CloseQuote NSLinguisticTag
	// Conjunction: This token is a conjunction.
	Conjunction NSLinguisticTag
	// Dash: This token is a dash.
	Dash NSLinguisticTag
	// Determiner: This token is a determiner.
	Determiner NSLinguisticTag
	// Idiom: This token is an idiom.
	Idiom NSLinguisticTag
	// Interjection: This token is an interjection.
	Interjection NSLinguisticTag
	// Noun: The token is a noun.
	Noun NSLinguisticTag
	// Number: This token is a number.
	Number NSLinguisticTag
	// OpenParenthesis: This token is an open parenthesis.
	OpenParenthesis NSLinguisticTag
	// OpenQuote: This token is an open quote.
	OpenQuote NSLinguisticTag
	// OrganizationName: This token is an organization name.
	OrganizationName NSLinguisticTag
	// Other: The token indicates a non-linguistic item, such as a symbol.
	Other NSLinguisticTag
	// OtherPunctuation: This token is punctuation other than a kind described by other lexical classes (sentence terminator, open or close quote, open or close parenthesis, word joiner, and dash).
	OtherPunctuation NSLinguisticTag
	// OtherWhitespace: This token is whitespace other than a kind described by other lexical classes (paragraph break).
	OtherWhitespace NSLinguisticTag
	// OtherWord: This token is a word other than a kind described by other lexical classes (noun, verb, adjective, adverb, pronoun, determiner, particle, preposition, number, conjunction, interjection, classifier, and idiom).
	OtherWord NSLinguisticTag
	// ParagraphBreak: This token is a paragraph break.
	ParagraphBreak NSLinguisticTag
	// Particle: This token is a particle.
	Particle NSLinguisticTag
	// PersonalName: This token is a personal name.
	PersonalName NSLinguisticTag
	// PlaceName: This token is a place name.
	PlaceName NSLinguisticTag
	// Preposition: This token is a preposition.
	Preposition NSLinguisticTag
	// Pronoun: This token is a pronoun.
	Pronoun NSLinguisticTag
	// Punctuation: The token indicates punctuation.
	Punctuation NSLinguisticTag
	// SentenceTerminator: This token is a sentence terminator.
	SentenceTerminator NSLinguisticTag
	// Verb: This token is a verb.
	Verb NSLinguisticTag
	// Whitespace: The token indicates white space of any sort.
	Whitespace NSLinguisticTag
	// Word: The token indicates a word.
	Word NSLinguisticTag
	// WordJoiner: This token is a word joiner.
	WordJoiner NSLinguisticTag
}

// NSProgressFileOperationKinds provides typed accessors for [NSProgressFileOperationKind] constants.
var NSProgressFileOperationKinds struct {
	// Copying: The progress is tracking the copying of a file from source to destination.
	Copying NSProgressFileOperationKind
	// DecompressingAfterDownloading: The progress is tracking file decompression after a download.
	DecompressingAfterDownloading NSProgressFileOperationKind
	// Downloading: The progress is tracking a file download operation.
	Downloading NSProgressFileOperationKind
	Duplicating NSProgressFileOperationKind
	// Receiving: The progress is tracking the receipt of a file from another source.
	Receiving NSProgressFileOperationKind
	// Uploading: The progress is tracking a file upload operation.
	Uploading NSProgressFileOperationKind
}

// StreamSocketSecurityLevels provides typed accessors for [StreamSocketSecurityLevel] constants.
var StreamSocketSecurityLevels struct {
	// NegotiatedSSL: Specifies that the highest level security protocol that can be negotiated be set as the security protocol for a socket stream.
	NegotiatedSSL StreamSocketSecurityLevel
	// None: Specifies that no security level be set for a socket stream.
	None StreamSocketSecurityLevel
	// SSLv2: Specifies that SSL version 2 be set as the security protocol for a socket stream.
	SSLv2 StreamSocketSecurityLevel
	// SSLv3: Specifies that SSL version 3 be set as the security protocol for a socket stream.
	SSLv3 StreamSocketSecurityLevel
	// TLSv1: Specifies that TLS version 1 be set as the security protocol for a socket stream.
	TLSv1 StreamSocketSecurityLevel
}

// StringTransforms provides typed accessors for [StringTransform] constants.
var StringTransforms struct {
	// FullwidthToHalfwidth: A constant containing the transformation of a string from full-width CJK characters to half-width forms.
	FullwidthToHalfwidth StringTransform
	// HiraganaToKatakana: A constant containing the transliteration of a string from Hiragana script to Katakana script.
	HiraganaToKatakana StringTransform
	// LatinToArabic: A constant containing the transliteration of a string from Latin script to Arabic script.
	LatinToArabic StringTransform
	// LatinToCyrillic: A constant containing the transliteration of a string from Latin script to Cyrillic script.
	LatinToCyrillic StringTransform
	// LatinToGreek: A constant containing the transliteration of a string from Latin script to Greek script.
	LatinToGreek StringTransform
	// LatinToHangul: A constant containing the transliteration of a string from Latin script to Hangul script.
	LatinToHangul StringTransform
	// LatinToHebrew: A constant containing the transliteration of a string from Latin script to Hebrew script.
	LatinToHebrew StringTransform
	// LatinToHiragana: A constant containing the transliteration of a string from Latin script to Hiragana script.
	LatinToHiragana StringTransform
	// LatinToKatakana: A constant containing the transliteration of a string from Latin script to Katakana script.
	LatinToKatakana StringTransform
	// LatinToThai: A constant containing the transliteration of a string from Latin script to Thai script.
	LatinToThai StringTransform
	// MandarinToLatin: A constant containing the transliteration of a string from Han script to Latin.
	MandarinToLatin StringTransform
	// StripCombiningMarks: A constant containing the transformation of a string by removing combining marks.
	StripCombiningMarks StringTransform
	// StripDiacritics: A constant containing the transformation of a string by removing diacritics.
	StripDiacritics StringTransform
	// ToLatin: A constant containing the transliteration of a string from any script to Latin script.
	ToLatin StringTransform
	// ToUnicodeName: An identifier for a transform that converts characters to Unicode names.
	ToUnicodeName StringTransform
	// ToXMLHex: A constant containing the transformation of a string from characters to XML hexadecimal escape codes.
	ToXMLHex StringTransform
}

// URLFileResourceTypes provides typed accessors for [URLFileResourceType] constants.
var URLFileResourceTypes struct {
	// BlockSpecial: The resource is a block special file.
	BlockSpecial URLFileResourceType
	// CharacterSpecial: The resource is a character special file.
	CharacterSpecial URLFileResourceType
	// Directory: The resource is a directory.
	Directory URLFileResourceType
	// NamedPipe: The resource is a named pipe.
	NamedPipe URLFileResourceType
	// Regular: The resource is a regular file.
	Regular URLFileResourceType
	// Socket: The resource is a socket.
	Socket URLFileResourceType
	// SymbolicLink: The resource is a symbolic link.
	SymbolicLink URLFileResourceType
	// Unknown: The resource’s type is unknown.
	Unknown URLFileResourceType
}

// URLUbiquitousItemDownloadingStatuss provides typed accessors for [URLUbiquitousItemDownloadingStatus] constants.
var URLUbiquitousItemDownloadingStatuss struct {
	// Current: A local copy of this item exists and is the most up-to-date version known to the device.
	Current URLUbiquitousItemDownloadingStatus
	// Downloaded: A local copy of this item exists, but it is stale. The most recent version will be downloaded as soon as possible.
	Downloaded URLUbiquitousItemDownloadingStatus
	// NotDownloaded: This item has not been downloaded yet. Use [startDownloadingUbiquitousItem(at:)](<doc://com.apple.foundation/documentation/Foundation/FileManager/startDownloadingUbiquitousItem(at:)>) to download it.
	NotDownloaded URLUbiquitousItemDownloadingStatus
}

// URLUbiquitousSharedItemPermissionss provides typed accessors for [URLUbiquitousSharedItemPermissions] constants.
var URLUbiquitousSharedItemPermissionss struct {
	ReadOnly URLUbiquitousSharedItemPermissions
	ReadWrite URLUbiquitousSharedItemPermissions
}

// URLUbiquitousSharedItemRoles provides typed accessors for [URLUbiquitousSharedItemRole] constants.
var URLUbiquitousSharedItemRoles struct {
	Owner URLUbiquitousSharedItemRole
	Participant URLUbiquitousSharedItemRole
}

