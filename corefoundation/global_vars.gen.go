// Code generated from Apple documentation. DO NOT EDIT.

package corefoundation

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// KCFAllocatorDefault is this is a synonym for [NULL].
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFAllocatorDefault
	KCFAllocatorDefault CFAllocatorRef
	// KCFAllocatorMalloc is this allocator uses `malloc()`, `realloc()`, and `free()`.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFAllocatorMalloc
	KCFAllocatorMalloc CFAllocatorRef
	// KCFAllocatorMallocZone is this allocator explicitly uses the default malloc zone, returned by `malloc_default_zone()`.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFAllocatorMallocZone
	KCFAllocatorMallocZone CFAllocatorRef
	// KCFAllocatorNull is this allocator does nothing—it allocates no memory.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFAllocatorNull
	KCFAllocatorNull CFAllocatorRef
	// KCFAllocatorSystemDefault is default system allocator.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFAllocatorSystemDefault
	KCFAllocatorSystemDefault CFAllocatorRef
	// KCFAllocatorUseContext is special allocator argument to [CFAllocatorCreate(_:_:)]—it uses the functions given in the context to allocate the allocator.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFAllocatorUseContext
	KCFAllocatorUseContext CFAllocatorRef
)

var (
	// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier/banglaCalendar
	KCFBanglaCalendar CFCalendarIdentifier
	// KCFBuddhistCalendar is specifies the Buddhist calendar.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier/buddhistCalendar
	KCFBuddhistCalendar CFCalendarIdentifier
	// KCFChineseCalendar is specifies the Chinese calendar.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier/chineseCalendar
	KCFChineseCalendar CFCalendarIdentifier
	// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier/dangiCalendar
	KCFDangiCalendar CFCalendarIdentifier
	// KCFGregorianCalendar is the name of the calendar currently supported by the [calendarName] property.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier/gregorianCalendar
	KCFGregorianCalendar CFCalendarIdentifier
	// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier/gujaratiCalendar
	KCFGujaratiCalendar CFCalendarIdentifier
	// KCFHebrewCalendar is specifies the Hebrew calendar.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier/hebrewCalendar
	KCFHebrewCalendar CFCalendarIdentifier
	// KCFISO8601Calendar is specifies the ISO 8601 calendar.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier/cfiso8601Calendar
	KCFISO8601Calendar CFCalendarIdentifier
	// KCFIndianCalendar is specifies the Indian calendar.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier/indianCalendar
	KCFIndianCalendar CFCalendarIdentifier
	// KCFIslamicCalendar is specifies the Islamic calendar.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier/islamicCalendar
	KCFIslamicCalendar CFCalendarIdentifier
	// KCFIslamicCivilCalendar is specifies the Islamic tabular calendar with Friday (civil) origin.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier/islamicCivilCalendar
	KCFIslamicCivilCalendar CFCalendarIdentifier
	// KCFIslamicTabularCalendar is specifies the Islamic tabular calendar with Thursday (astronomical) origin.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier/islamicTabularCalendar
	KCFIslamicTabularCalendar CFCalendarIdentifier
	// KCFIslamicUmmAlQuraCalendar is specifies the Islamic Umm Al Qura calendar.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier/islamicUmmAlQuraCalendar
	KCFIslamicUmmAlQuraCalendar CFCalendarIdentifier
	// KCFJapaneseCalendar is specifies the Japanese calendar.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier/japaneseCalendar
	KCFJapaneseCalendar CFCalendarIdentifier
	// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier/kannadaCalendar
	KCFKannadaCalendar CFCalendarIdentifier
	// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier/malayalamCalendar
	KCFMalayalamCalendar CFCalendarIdentifier
	// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier/marathiCalendar
	KCFMarathiCalendar CFCalendarIdentifier
	// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier/odiaCalendar
	KCFOdiaCalendar CFCalendarIdentifier
	// KCFPersianCalendar is specifies the Persian calendar.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier/persianCalendar
	KCFPersianCalendar CFCalendarIdentifier
	// KCFRepublicOfChinaCalendar is specifies the calendar for the Republic of China.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier/republicOfChinaCalendar
	KCFRepublicOfChinaCalendar CFCalendarIdentifier
	// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier/tamilCalendar
	KCFTamilCalendar CFCalendarIdentifier
	// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier/teluguCalendar
	KCFTeluguCalendar CFCalendarIdentifier
	// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier/vietnameseCalendar
	KCFVietnameseCalendar CFCalendarIdentifier
	// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier/vikramCalendar
	KCFVikramCalendar CFCalendarIdentifier
)

var (
	// KCFBooleanFalse is boolean false value.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFBooleanFalse
	KCFBooleanFalse CFBooleanRef
	// KCFBooleanTrue is boolean true value.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFBooleanTrue
	KCFBooleanTrue CFBooleanRef
)

var (
	// KCFBundleDevelopmentRegionKey is the name of the development language of the bundle.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFBundleDevelopmentRegionKey
	KCFBundleDevelopmentRegionKey string
	// KCFBundleExecutableKey is the name of the executable in this bundle (if any).
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFBundleExecutableKey
	KCFBundleExecutableKey string
	// KCFBundleIdentifierKey is the bundle identifier.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFBundleIdentifierKey
	KCFBundleIdentifierKey string
	// KCFBundleInfoDictionaryVersionKey is the version of the information property list format.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFBundleInfoDictionaryVersionKey
	KCFBundleInfoDictionaryVersionKey string
	// KCFBundleLocalizationsKey is allows an unbundled application that handles localization itself to specify which localizations it has available.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFBundleLocalizationsKey
	KCFBundleLocalizationsKey string
	// KCFBundleNameKey is the human-readable name of the bundle.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFBundleNameKey
	KCFBundleNameKey string
	// KCFBundleVersionKey is the version number of the bundle.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFBundleVersionKey
	KCFBundleVersionKey string
	// KCFErrorDescriptionKey is key to identify the description in the `userInfo` dictionary.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFErrorDescriptionKey
	KCFErrorDescriptionKey string
	// KCFErrorFilePathKey is key to identify associated file path in the `userInfo` dictionary.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFErrorFilePathKey
	KCFErrorFilePathKey string
	// KCFErrorLocalizedDescriptionKey is key to identify the user-presentable description in the `userInfo` dictionary.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFErrorLocalizedDescriptionKey
	KCFErrorLocalizedDescriptionKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFErrorLocalizedFailureKey
	KCFErrorLocalizedFailureKey string
	// KCFErrorLocalizedFailureReasonKey is key to identify the user-presentable failure reason in the `userInfo` dictionary.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFErrorLocalizedFailureReasonKey
	KCFErrorLocalizedFailureReasonKey string
	// KCFErrorLocalizedRecoverySuggestionKey is key to identify the user-presentable recovery suggestion in the `userInfo` dictionary.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFErrorLocalizedRecoverySuggestionKey
	KCFErrorLocalizedRecoverySuggestionKey string
	// KCFErrorURLKey is key to identify associated URL in the `userInfo` dictionary.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFErrorURLKey
	KCFErrorURLKey string
	// KCFErrorUnderlyingErrorKey is key to identify the underlying error in the `userInfo` dictionary.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFErrorUnderlyingErrorKey
	KCFErrorUnderlyingErrorKey string
	// KCFPlugInDynamicRegisterFunctionKey is used to specify a plug-in’s registration function.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFPlugInDynamicRegisterFunctionKey
	KCFPlugInDynamicRegisterFunctionKey string
	// KCFPlugInDynamicRegistrationKey is indicates whether a plug-in requires dynamic registration.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFPlugInDynamicRegistrationKey
	KCFPlugInDynamicRegistrationKey string
	// KCFPlugInFactoriesKey is used to statically register factory functions.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFPlugInFactoriesKey
	KCFPlugInFactoriesKey string
	// KCFPlugInTypesKey is used to statically register the factories that can create each supported type.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFPlugInTypesKey
	KCFPlugInTypesKey string
	// KCFPlugInUnloadFunctionKey is used to specify a plug-in’s unload function.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFPlugInUnloadFunctionKey
	KCFPlugInUnloadFunctionKey string
	// KCFPreferencesAnyApplication is indicates a preference that applies to any application.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFPreferencesAnyApplication
	KCFPreferencesAnyApplication string
	// KCFPreferencesAnyHost is indicates a preference that applies to any host.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFPreferencesAnyHost
	KCFPreferencesAnyHost string
	// KCFPreferencesAnyUser is indicates a preference that applies to any user.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFPreferencesAnyUser
	KCFPreferencesAnyUser string
	// KCFPreferencesCurrentApplication is indicates a preference that applies only to the current application.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFPreferencesCurrentApplication
	KCFPreferencesCurrentApplication string
	// KCFPreferencesCurrentHost is indicates a preference that applies only to the current host.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFPreferencesCurrentHost
	KCFPreferencesCurrentHost string
	// KCFPreferencesCurrentUser is indicates a preference that applies only to the current user.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFPreferencesCurrentUser
	KCFPreferencesCurrentUser string
	// KCFSocketCommandKey is not used.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFSocketCommandKey
	KCFSocketCommandKey string
	// KCFSocketErrorKey is not used.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFSocketErrorKey
	KCFSocketErrorKey string
	// KCFSocketNameKey is not used.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFSocketNameKey
	KCFSocketNameKey string
	// KCFSocketRegisterCommand is not used.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFSocketRegisterCommand
	KCFSocketRegisterCommand string
	// KCFSocketResultKey is not used.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFSocketResultKey
	KCFSocketResultKey string
	// KCFSocketRetrieveCommand is not used.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFSocketRetrieveCommand
	KCFSocketRetrieveCommand string
	// KCFSocketValueKey is not used.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFSocketValueKey
	KCFSocketValueKey string
	// KCFStreamPropertySOCKSPassword is constant for the key required to set a user’s password.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStreamPropertySOCKSPassword
	KCFStreamPropertySOCKSPassword string
	// KCFStreamPropertySOCKSProxy is sOCKS proxy property key.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStreamPropertySOCKSProxy
	KCFStreamPropertySOCKSProxy string
	// KCFStreamPropertySOCKSProxyHost is constant for the SOCKS proxy host key.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStreamPropertySOCKSProxyHost
	KCFStreamPropertySOCKSProxyHost string
	// KCFStreamPropertySOCKSProxyPort is constant for the SOCKS proxy host port key.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStreamPropertySOCKSProxyPort
	KCFStreamPropertySOCKSProxyPort string
	// KCFStreamPropertySOCKSUser is constant for the key required to set a user name.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStreamPropertySOCKSUser
	KCFStreamPropertySOCKSUser string
	// KCFStreamPropertySOCKSVersion is constant for the SOCKS version key.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStreamPropertySOCKSVersion
	KCFStreamPropertySOCKSVersion string
	// KCFStreamPropertyShouldCloseNativeSocket is should Close Native Socket property key.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStreamPropertyShouldCloseNativeSocket
	KCFStreamPropertyShouldCloseNativeSocket string
	// KCFStreamPropertySocketSecurityLevel is socket Security Level property key.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStreamPropertySocketSecurityLevel
	KCFStreamPropertySocketSecurityLevel string
	// KCFStreamSocketSOCKSVersion4 is constant used in the `kCFStreamSockerSOCKSVersion` key to specify SOCKS4 as the SOCKS version for the stream.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStreamSocketSOCKSVersion4
	KCFStreamSocketSOCKSVersion4 string
	// KCFStreamSocketSOCKSVersion5 is constant used in the `kCFStreamSOCKSVersion` key to specify SOCKS5 as the SOCKS version for the stream.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStreamSocketSOCKSVersion5
	KCFStreamSocketSOCKSVersion5 string
	// KCFStreamSocketSecurityLevelNegotiatedSSL is specifies that the highest level security protocol that can be negotiated be set as the security protocol for a socket stream.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStreamSocketSecurityLevelNegotiatedSSL
	KCFStreamSocketSecurityLevelNegotiatedSSL string
	// KCFStreamSocketSecurityLevelNone is specifies that no security level be set.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStreamSocketSecurityLevelNone
	KCFStreamSocketSecurityLevelNone string
	// KCFStreamSocketSecurityLevelTLSv1 is specifies that TLS version 1 be set as the security protocol for a socket stream.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStreamSocketSecurityLevelTLSv1
	KCFStreamSocketSecurityLevelTLSv1 string
	// KCFStringTransformFullwidthHalfwidth is the identifier of a reversible transform to convert full-width characters to their half-width equivalents.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStringTransformFullwidthHalfwidth
	KCFStringTransformFullwidthHalfwidth string
	// KCFStringTransformHiraganaKatakana is the identifier of a reversible transform to transliterate text to Katakana from Hiragana.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStringTransformHiraganaKatakana
	KCFStringTransformHiraganaKatakana string
	// KCFStringTransformLatinArabic is the identifier of a reversible transform to transliterate text to Arabic from Latin.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStringTransformLatinArabic
	KCFStringTransformLatinArabic string
	// KCFStringTransformLatinCyrillic is the identifier of a reversible transform to transliterate text to Cyrillic from Latin.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStringTransformLatinCyrillic
	KCFStringTransformLatinCyrillic string
	// KCFStringTransformLatinGreek is the identifier of a reversible transform to transliterate text to Greek from Latin.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStringTransformLatinGreek
	KCFStringTransformLatinGreek string
	// KCFStringTransformLatinHangul is the identifier of a reversible transform to transliterate text to Hangul from Latin.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStringTransformLatinHangul
	KCFStringTransformLatinHangul string
	// KCFStringTransformLatinHebrew is the identifier of a reversible transform to transliterate text to Hebrew from Latin.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStringTransformLatinHebrew
	KCFStringTransformLatinHebrew string
	// KCFStringTransformLatinHiragana is the identifier of a reversible transform to transliterate text to Hiragana from Latin.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStringTransformLatinHiragana
	KCFStringTransformLatinHiragana string
	// KCFStringTransformLatinKatakana is the identifier of a reversible transform to transliterate text to Katakana from Latin.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStringTransformLatinKatakana
	KCFStringTransformLatinKatakana string
	// KCFStringTransformLatinThai is the identifier of a reversible transform to transliterate text to Thai from Latin.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStringTransformLatinThai
	KCFStringTransformLatinThai string
	// KCFStringTransformMandarinLatin is the identifier of a transform to transliterate text to Latin from ideographs interpreted as Mandarin Chinese. This transform is not reversible.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStringTransformMandarinLatin
	KCFStringTransformMandarinLatin string
	// KCFStringTransformStripCombiningMarks is the identifier of a transform to strip combining marks (accents or diacritics).
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStringTransformStripCombiningMarks
	KCFStringTransformStripCombiningMarks string
	// KCFStringTransformStripDiacritics is the identifier of a transform to remove diacritic markings.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStringTransformStripDiacritics
	KCFStringTransformStripDiacritics string
	// KCFStringTransformToLatin is the identifier of a transform to transliterate all text possible to Latin script. Ideographs are transliterated as Mandarin Chinese.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStringTransformToLatin
	KCFStringTransformToLatin string
	// KCFStringTransformToUnicodeName is the identifier of a reversible transform to transliterate characters other than printable ASCII to their Unicode character name in braces.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStringTransformToUnicodeName
	KCFStringTransformToUnicodeName string
	// KCFStringTransformToXMLHex is the identifier of a reversible transform to transliterate characters other than printable ASCII to XML/HTML numeric entities.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStringTransformToXMLHex
	KCFStringTransformToXMLHex string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLAddedToDirectoryDateKey
	KCFURLAddedToDirectoryDateKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLApplicationIsScriptableKey
	KCFURLApplicationIsScriptableKey string
	// KCFURLAttributeModificationDateKey is key for the last time the resource’s attributes were modified, returned as a [CFDate] object if the volume supports attribute modification dates, or `nil` if attribute modification dates are unsupported.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLAttributeModificationDateKey
	KCFURLAttributeModificationDateKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLCanonicalPathKey
	KCFURLCanonicalPathKey string
	// KCFURLContentAccessDateKey is key for the last time the resource was accessed, returned as a [CFDate] object if the volume supports access dates, or `nil` if access dates are unsupported.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLContentAccessDateKey
	KCFURLContentAccessDateKey string
	// KCFURLContentModificationDateKey is key for the last time the resource was modified, returned as a [CFDate] object if the volume supports modification dates, or `nil` if modification dates are unsupported.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLContentModificationDateKey
	KCFURLContentModificationDateKey string
	// KCFURLCreationDateKey is key for the resource’s creation date, returned as a [CFDate] object if the volume supports creation dates, or `nil` if creation dates are unsupported.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLCreationDateKey
	KCFURLCreationDateKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLDirectoryEntryCountKey
	KCFURLDirectoryEntryCountKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLDocumentIdentifierKey
	KCFURLDocumentIdentifierKey string
	// KCFURLFileAllocatedSizeKey is key for the total size allocated on disk for the file, returned as an [CFNumber] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLFileAllocatedSizeKey
	KCFURLFileAllocatedSizeKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLFileContentIdentifierKey
	KCFURLFileContentIdentifierKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLFileIdentifierKey
	KCFURLFileIdentifierKey string
	// KCFURLFileResourceIdentifierKey is key for the resource’s unique identifier, returned as a [CFType] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLFileResourceIdentifierKey
	KCFURLFileResourceIdentifierKey string
	// KCFURLFileResourceTypeBlockSpecial is the resource is a block special file.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLFileResourceTypeBlockSpecial
	KCFURLFileResourceTypeBlockSpecial string
	// KCFURLFileResourceTypeCharacterSpecial is the resource is a character special file.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLFileResourceTypeCharacterSpecial
	KCFURLFileResourceTypeCharacterSpecial string
	// KCFURLFileResourceTypeDirectory is the resource is a directory.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLFileResourceTypeDirectory
	KCFURLFileResourceTypeDirectory string
	// KCFURLFileResourceTypeKey is key for the resource’s object type, returned as a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLFileResourceTypeKey
	KCFURLFileResourceTypeKey string
	// KCFURLFileResourceTypeNamedPipe is the resource is a named pipe.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLFileResourceTypeNamedPipe
	KCFURLFileResourceTypeNamedPipe string
	// KCFURLFileResourceTypeRegular is the resource is a regular file.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLFileResourceTypeRegular
	KCFURLFileResourceTypeRegular string
	// KCFURLFileResourceTypeSocket is the resource is a socket.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLFileResourceTypeSocket
	KCFURLFileResourceTypeSocket string
	// KCFURLFileResourceTypeSymbolicLink is the resource is a symbolic link.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLFileResourceTypeSymbolicLink
	KCFURLFileResourceTypeSymbolicLink string
	// KCFURLFileResourceTypeUnknown is the resource’s type is unknown.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLFileResourceTypeUnknown
	KCFURLFileResourceTypeUnknown string
	// KCFURLFileSecurityKey is key for the resource’s security information, returned as a [CFFileSecurity] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLFileSecurityKey
	KCFURLFileSecurityKey string
	// KCFURLFileSizeKey is key for the file’s size in bytes, returned as a [CFNumber] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLFileSizeKey
	KCFURLFileSizeKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLGenerationIdentifierKey
	KCFURLGenerationIdentifierKey string
	// KCFURLHasHiddenExtensionKey is key for determining whether the resource’s extension is normally removed from its localized name, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLHasHiddenExtensionKey
	KCFURLHasHiddenExtensionKey string
	// KCFURLIsAliasFileKey is key for determining whether the file is an alias, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLIsAliasFileKey
	KCFURLIsAliasFileKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLIsApplicationKey
	KCFURLIsApplicationKey string
	// KCFURLIsDirectoryKey is key for determining whether the resource is a directory, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLIsDirectoryKey
	KCFURLIsDirectoryKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLIsExcludedFromBackupKey
	KCFURLIsExcludedFromBackupKey string
	// KCFURLIsExecutableKey is key for determining whether the current process (as determined by the EUID) can execute the resource (if it is a file) or search the resource (if it is a directory), returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLIsExecutableKey
	KCFURLIsExecutableKey string
	// KCFURLIsHiddenKey is key for determining whether the resource is normally not displayed to users, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLIsHiddenKey
	KCFURLIsHiddenKey string
	// KCFURLIsMountTriggerKey is key for determining whether the URL is a file system trigger directory, returned as a [CFBoolean] object. Traversing or opening a file system trigger directory causes an attempt to mount a file system on the directory.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLIsMountTriggerKey
	KCFURLIsMountTriggerKey string
	// KCFURLIsPackageKey is key for determining whether the resource is a packaged directory, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLIsPackageKey
	KCFURLIsPackageKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLIsPurgeableKey
	KCFURLIsPurgeableKey string
	// KCFURLIsReadableKey is key for determining whether the current process (as determined by the EUID) can read the resource, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLIsReadableKey
	KCFURLIsReadableKey string
	// KCFURLIsRegularFileKey is key for determining whether the resource is a regular file, as opposed to a directory or a symbolic link. Returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLIsRegularFileKey
	KCFURLIsRegularFileKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLIsSparseKey
	KCFURLIsSparseKey string
	// KCFURLIsSymbolicLinkKey is key for determining whether the resource is a symbolic link, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLIsSymbolicLinkKey
	KCFURLIsSymbolicLinkKey string
	// KCFURLIsSystemImmutableKey is key for determining whether the resource’s system immutable bit is set, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLIsSystemImmutableKey
	KCFURLIsSystemImmutableKey string
	// KCFURLIsUbiquitousItemKey is a [CFBoolean] value that tells whether the item is synced to the cloud. (read-only).
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLIsUbiquitousItemKey
	KCFURLIsUbiquitousItemKey string
	// KCFURLIsUserImmutableKey is key for determining whether the resource’s user immutable bit is set, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLIsUserImmutableKey
	KCFURLIsUserImmutableKey string
	// KCFURLIsVolumeKey is key for determining whether the resource is the root directory of a volume, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLIsVolumeKey
	KCFURLIsVolumeKey string
	// KCFURLIsWritableKey is key for determining whether the current process (as determined by the EUID) can write to the resource, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLIsWritableKey
	KCFURLIsWritableKey string
	// KCFURLKeysOfUnsetValuesKey is key for the resource properties that have not been set after the [CFURLSetResourcePropertiesForKeys(_:_:_:)] function returns an error, returned as an array of [CFString] objects.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLKeysOfUnsetValuesKey
	KCFURLKeysOfUnsetValuesKey string
	// KCFURLLabelNumberKey is key for the resource’s label number, returned as a [CFNumber] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLLabelNumberKey
	KCFURLLabelNumberKey string
	// KCFURLLinkCountKey is key for the number of hard links to the resource, returned as a [CFNumber] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLLinkCountKey
	KCFURLLinkCountKey string
	// KCFURLLocalizedLabelKey is key for the resource’s localized label text, returned as a [CFString] object, or [NULL] if the resource has no localized label text.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLLocalizedLabelKey
	KCFURLLocalizedLabelKey string
	// KCFURLLocalizedNameKey is key for the resource’s localized or extension-hidden name, retuned as a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLLocalizedNameKey
	KCFURLLocalizedNameKey string
	// KCFURLLocalizedTypeDescriptionKey is key for the resource’s localized type description, returned as a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLLocalizedTypeDescriptionKey
	KCFURLLocalizedTypeDescriptionKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLMayHaveExtendedAttributesKey
	KCFURLMayHaveExtendedAttributesKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLMayShareFileContentKey
	KCFURLMayShareFileContentKey string
	// KCFURLNameKey is key for the resource’s name in the file system, returned as a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLNameKey
	KCFURLNameKey string
	// KCFURLParentDirectoryURLKey is key for the parent directory of the resource, returned as a [CFURL] object, or `nil` if the resource is the root directory of its volume.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLParentDirectoryURLKey
	KCFURLParentDirectoryURLKey string
	// KCFURLPathKey is a [CFString] value containing the URL’s path as a file system path. (read-only).
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLPathKey
	KCFURLPathKey string
	// KCFURLPreferredIOBlockSizeKey is key for the optimal block size to use when reading or writing this file’s data, returned as a [CFNumber] object, or [NULL] if the preferred size is not available.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLPreferredIOBlockSizeKey
	KCFURLPreferredIOBlockSizeKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLQuarantinePropertiesKey
	KCFURLQuarantinePropertiesKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLTagNamesKey
	KCFURLTagNamesKey string
	// KCFURLTotalFileAllocatedSizeKey is key for the total allocated size of the file in bytes, returned as a [CFNumber] object. This includes the size of any file metadata.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLTotalFileAllocatedSizeKey
	KCFURLTotalFileAllocatedSizeKey string
	// KCFURLTotalFileSizeKey is key for the total displayable size of the file in bytes, returned as a [CFNumber] object. This includes the size of any file metadata.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLTotalFileSizeKey
	KCFURLTotalFileSizeKey string
	// KCFURLTypeIdentifierKey is key for the resource’s uniform type identifier (UTI), returned as a [CFString] object.
	//
	// Deprecated: Deprecated since macOS 26.2. Use NSURLContentTypeKey instead
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLTypeIdentifierKey
	KCFURLTypeIdentifierKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLUbiquitousItemDownloadingErrorKey
	KCFURLUbiquitousItemDownloadingErrorKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLUbiquitousItemDownloadingStatusCurrent
	KCFURLUbiquitousItemDownloadingStatusCurrent string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLUbiquitousItemDownloadingStatusDownloaded
	KCFURLUbiquitousItemDownloadingStatusDownloaded string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLUbiquitousItemDownloadingStatusKey
	KCFURLUbiquitousItemDownloadingStatusKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLUbiquitousItemDownloadingStatusNotDownloaded
	KCFURLUbiquitousItemDownloadingStatusNotDownloaded string
	// KCFURLUbiquitousItemHasUnresolvedConflictsKey is a [CFBoolean] value that tells whether the item has conflicts outstanding. (read-only).
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLUbiquitousItemHasUnresolvedConflictsKey
	KCFURLUbiquitousItemHasUnresolvedConflictsKey string
	// KCFURLUbiquitousItemIsDownloadingKey is a [CFBoolean] value that tells whether data for the item is being downloaded. (read-only).
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLUbiquitousItemIsDownloadingKey
	KCFURLUbiquitousItemIsDownloadingKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLUbiquitousItemIsExcludedFromSyncKey
	KCFURLUbiquitousItemIsExcludedFromSyncKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLUbiquitousItemIsSyncPausedKey
	KCFURLUbiquitousItemIsSyncPausedKey string
	// KCFURLUbiquitousItemIsUploadedKey is a [CFBoolean] value that tells whether there is data present in the cloud for this item. (read-only).
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLUbiquitousItemIsUploadedKey
	KCFURLUbiquitousItemIsUploadedKey string
	// KCFURLUbiquitousItemIsUploadingKey is a [CFBoolean] value that tells whether data for the item is being uploaded. (read-only).
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLUbiquitousItemIsUploadingKey
	KCFURLUbiquitousItemIsUploadingKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLUbiquitousItemSupportedSyncControlsKey
	KCFURLUbiquitousItemSupportedSyncControlsKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLUbiquitousItemUploadingErrorKey
	KCFURLUbiquitousItemUploadingErrorKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeAvailableCapacityForImportantUsageKey
	KCFURLVolumeAvailableCapacityForImportantUsageKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeAvailableCapacityForOpportunisticUsageKey
	KCFURLVolumeAvailableCapacityForOpportunisticUsageKey string
	// KCFURLVolumeAvailableCapacityKey is key for the volume’s available capacity in bytes, returned as a [CFNumber] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeAvailableCapacityKey
	KCFURLVolumeAvailableCapacityKey string
	// KCFURLVolumeCreationDateKey is key for the volume’s creation date, returned as a [CFDate] object, or [NULL] if it cannot be determined.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeCreationDateKey
	KCFURLVolumeCreationDateKey string
	// KCFURLVolumeIdentifierKey is key for the unique identifier of the resource’s volume, returned as a [CFType] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeIdentifierKey
	KCFURLVolumeIdentifierKey string
	// KCFURLVolumeIsAutomountedKey is key for determining whether the volume is automounted, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeIsAutomountedKey
	KCFURLVolumeIsAutomountedKey string
	// KCFURLVolumeIsBrowsableKey is key for determining whether the volume is visible in GUI-based file-browsing environments, such as the Desktop or the Finder application, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeIsBrowsableKey
	KCFURLVolumeIsBrowsableKey string
	// KCFURLVolumeIsEjectableKey is key for determining whether the volume is ejectable from the drive mechanism under software control, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeIsEjectableKey
	KCFURLVolumeIsEjectableKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeIsEncryptedKey
	KCFURLVolumeIsEncryptedKey string
	// KCFURLVolumeIsInternalKey is key for determining whether the volume is connected to an internal bus, returned as a [CFBoolean] object, or [NULL] if it cannot be determined.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeIsInternalKey
	KCFURLVolumeIsInternalKey string
	// KCFURLVolumeIsJournalingKey is key for determining whether the volume is currently journaling, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeIsJournalingKey
	KCFURLVolumeIsJournalingKey string
	// KCFURLVolumeIsLocalKey is key for determining whether the volume is stored on a local device, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeIsLocalKey
	KCFURLVolumeIsLocalKey string
	// KCFURLVolumeIsReadOnlyKey is key for determining whether the volume is read-only, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeIsReadOnlyKey
	KCFURLVolumeIsReadOnlyKey string
	// KCFURLVolumeIsRemovableKey is key for determining whether the volume is removable from the drive mechanism, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeIsRemovableKey
	KCFURLVolumeIsRemovableKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeIsRootFileSystemKey
	KCFURLVolumeIsRootFileSystemKey string
	// KCFURLVolumeLocalizedFormatDescriptionKey is key for the volume’s descriptive format name, returned as a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeLocalizedFormatDescriptionKey
	KCFURLVolumeLocalizedFormatDescriptionKey string
	// KCFURLVolumeLocalizedNameKey is the user-presentable name of the volume, returned as a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeLocalizedNameKey
	KCFURLVolumeLocalizedNameKey string
	// KCFURLVolumeMaximumFileSizeKey is key for the largest file size supported by the volume in bytes, returned as a [CFNumber] object, or [NULL] if it cannot be determined.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeMaximumFileSizeKey
	KCFURLVolumeMaximumFileSizeKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeMountFromLocationKey
	KCFURLVolumeMountFromLocationKey string
	// KCFURLVolumeNameKey is the name of the volume, returned as a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeNameKey
	KCFURLVolumeNameKey string
	// KCFURLVolumeResourceCountKey is key for the total number of resources on the volume, returned as a [CFNumber] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeResourceCountKey
	KCFURLVolumeResourceCountKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeSubtypeKey
	KCFURLVolumeSubtypeKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeSupportsAccessPermissionsKey
	KCFURLVolumeSupportsAccessPermissionsKey string
	// KCFURLVolumeSupportsAdvisoryFileLockingKey is key for determining whether the volume implements whole-file advisory locks in the style of flock, along with the `O_EXLOCK` and `O_SHLOCK` flags of the open function, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeSupportsAdvisoryFileLockingKey
	KCFURLVolumeSupportsAdvisoryFileLockingKey string
	// KCFURLVolumeSupportsCasePreservedNamesKey is key for determining whether the volume supports case-preserved names, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeSupportsCasePreservedNamesKey
	KCFURLVolumeSupportsCasePreservedNamesKey string
	// KCFURLVolumeSupportsCaseSensitiveNamesKey is key for determining whether the volume supports case-sensitive names, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeSupportsCaseSensitiveNamesKey
	KCFURLVolumeSupportsCaseSensitiveNamesKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeSupportsCompressionKey
	KCFURLVolumeSupportsCompressionKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeSupportsExclusiveRenamingKey
	KCFURLVolumeSupportsExclusiveRenamingKey string
	// KCFURLVolumeSupportsExtendedSecurityKey is key for determining whether the volume supports extended security (access control lists), returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeSupportsExtendedSecurityKey
	KCFURLVolumeSupportsExtendedSecurityKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeSupportsFileCloningKey
	KCFURLVolumeSupportsFileCloningKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeSupportsFileProtectionKey
	KCFURLVolumeSupportsFileProtectionKey string
	// KCFURLVolumeSupportsHardLinksKey is key for determining whether the volume supports hard links, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeSupportsHardLinksKey
	KCFURLVolumeSupportsHardLinksKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeSupportsImmutableFilesKey
	KCFURLVolumeSupportsImmutableFilesKey string
	// KCFURLVolumeSupportsJournalingKey is key for determining whether the volume supports journaling, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeSupportsJournalingKey
	KCFURLVolumeSupportsJournalingKey string
	// KCFURLVolumeSupportsPersistentIDsKey is key for determining whether the volume supports persistent IDs, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeSupportsPersistentIDsKey
	KCFURLVolumeSupportsPersistentIDsKey string
	// KCFURLVolumeSupportsRenamingKey is key for determining whether the volume can be renamed, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeSupportsRenamingKey
	KCFURLVolumeSupportsRenamingKey string
	// KCFURLVolumeSupportsRootDirectoryDatesKey is key for determining whether the volume supports reliable storage of times for the root directory, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeSupportsRootDirectoryDatesKey
	KCFURLVolumeSupportsRootDirectoryDatesKey string
	// KCFURLVolumeSupportsSparseFilesKey is key for determining whether the volume supports sparse files, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeSupportsSparseFilesKey
	KCFURLVolumeSupportsSparseFilesKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeSupportsSwapRenamingKey
	KCFURLVolumeSupportsSwapRenamingKey string
	// KCFURLVolumeSupportsSymbolicLinksKey is key for determining whether the volume supports symbolic links, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeSupportsSymbolicLinksKey
	KCFURLVolumeSupportsSymbolicLinksKey string
	// KCFURLVolumeSupportsVolumeSizesKey is key for determining whether the volume supports returning volume size information, returned as a [CFBoolean] object. If `true`, volume size information is available as values of the [kCFURLVolumeTotalCapacityKey] and [kCFURLVolumeAvailableCapacityKey] keys.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeSupportsVolumeSizesKey
	KCFURLVolumeSupportsVolumeSizesKey string
	// KCFURLVolumeSupportsZeroRunsKey is key for determining whether the volume supports zero runs, returned as a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeSupportsZeroRunsKey
	KCFURLVolumeSupportsZeroRunsKey string
	// KCFURLVolumeTotalCapacityKey is key for the volume’s capacity in bytes, returned as a [CFNumber] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeTotalCapacityKey
	KCFURLVolumeTotalCapacityKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeTypeNameKey
	KCFURLVolumeTypeNameKey string
	// KCFURLVolumeURLForRemountingKey is key for the URL needed to remount the network volume, returned as a [CFURL] object, or [NULL] if a URL is not available.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeURLForRemountingKey
	KCFURLVolumeURLForRemountingKey string
	// KCFURLVolumeURLKey is key for the root directory of the resource’s volume, returned as a [CFURL] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeURLKey
	KCFURLVolumeURLKey string
	// KCFURLVolumeUUIDStringKey is key for the volume’s persistent UUID, returned as a [CFString] object, or [NULL] if a persistent UUID is not available.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFURLVolumeUUIDStringKey
	KCFURLVolumeUUIDStringKey string
	// KCFUserNotificationAlertHeaderKey is the title of the notification dialog.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFUserNotificationAlertHeaderKey
	KCFUserNotificationAlertHeaderKey string
	// KCFUserNotificationAlertMessageKey is the message string to display in the dialog.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFUserNotificationAlertMessageKey
	KCFUserNotificationAlertMessageKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFUserNotificationAlertTopMostKey
	KCFUserNotificationAlertTopMostKey string
	// KCFUserNotificationAlternateButtonTitleKey is the title of an optional alternate button.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFUserNotificationAlternateButtonTitleKey
	KCFUserNotificationAlternateButtonTitleKey string
	// KCFUserNotificationCheckBoxTitlesKey is the list of titles for all the checkboxes or radio buttons to display.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFUserNotificationCheckBoxTitlesKey
	KCFUserNotificationCheckBoxTitlesKey string
	// KCFUserNotificationDefaultButtonTitleKey is the title of the default button.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFUserNotificationDefaultButtonTitleKey
	KCFUserNotificationDefaultButtonTitleKey string
	// KCFUserNotificationIconURLKey is a file URL pointing to the icon to display in the dialog.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFUserNotificationIconURLKey
	KCFUserNotificationIconURLKey string
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFUserNotificationKeyboardTypesKey
	KCFUserNotificationKeyboardTypesKey string
	// KCFUserNotificationLocalizationURLKey is a file URL pointing to a bundle that contains localized versions of the strings displayed in the dialog.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFUserNotificationLocalizationURLKey
	KCFUserNotificationLocalizationURLKey string
	// KCFUserNotificationOtherButtonTitleKey is the title of an optional third button.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFUserNotificationOtherButtonTitleKey
	KCFUserNotificationOtherButtonTitleKey string
	// KCFUserNotificationPopUpSelectionKey is the item that was selected from a pop-up menu.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFUserNotificationPopUpSelectionKey
	KCFUserNotificationPopUpSelectionKey string
	// KCFUserNotificationPopUpTitlesKey is the list of strings to display in a pop-up menu.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFUserNotificationPopUpTitlesKey
	KCFUserNotificationPopUpTitlesKey string
	// KCFUserNotificationProgressIndicatorValueKey is a value to indicate the progress of an operation.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFUserNotificationProgressIndicatorValueKey
	KCFUserNotificationProgressIndicatorValueKey string
	// KCFUserNotificationSoundURLKey is a file URL pointing to a sound that will be played when the alert appears.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFUserNotificationSoundURLKey
	KCFUserNotificationSoundURLKey string
	// KCFUserNotificationTextFieldTitlesKey is the list of titles for all the text fields to display.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFUserNotificationTextFieldTitlesKey
	KCFUserNotificationTextFieldTitlesKey string
	// KCFUserNotificationTextFieldValuesKey is the list of values to put into the text fields. If only one text field is to be displayed, you can pass its value string directly without putting it into an array first.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFUserNotificationTextFieldValuesKey
	KCFUserNotificationTextFieldValuesKey string
	// KCFXMLTreeErrorDescription is dictionary key whose value is a CFString containing a readable description of the error.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFXMLTreeErrorDescription
	KCFXMLTreeErrorDescription string
	// KCFXMLTreeErrorLineNumber is dictionary key whose value is a CFNumber containing the line number where the error was detected. This may not be the line number where the actual XML error is located.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFXMLTreeErrorLineNumber
	KCFXMLTreeErrorLineNumber string
	// KCFXMLTreeErrorLocation is dictionary key whose value is a CFNumber containing the byte location where the error was detected.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFXMLTreeErrorLocation
	KCFXMLTreeErrorLocation string
	// KCFXMLTreeErrorStatusCode is dictionary key whose value is a CFNumber containing the error status code. See [CFXMLParser] for possible status code values.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFXMLTreeErrorStatusCode
	KCFXMLTreeErrorStatusCode string
)

var (
	// KCFCopyStringBagCallBacks is predefined [CFBagCallBacks] structure containing a set of callbacks appropriate for use when the values in a CFBag are all CFString objects. The bag makes immutable copies of the strings placed into it.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFCopyStringBagCallBacks
	KCFCopyStringBagCallBacks CFBagCallBacks
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFTypeBagCallBacks
	KCFTypeBagCallBacks CFBagCallBacks
)

var (
	// KCFCopyStringDictionaryKeyCallBacks is predefined [CFDictionaryKeyCallBacks] structure containing a set of callbacks appropriate for use when the keys of a CFDictionary are all CFString objects, which may be mutable and need to be copied in order to serve as constant keys for the values in the dictionary.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFCopyStringDictionaryKeyCallBacks
	KCFCopyStringDictionaryKeyCallBacks CFDictionaryKeyCallBacks
	// KCFTypeDictionaryKeyCallBacks is predefined [CFDictionaryKeyCallBacks] structure containing a set of callbacks appropriate for use when the keys of a CFDictionary are all CFType-derived objects.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFTypeDictionaryKeyCallBacks
	KCFTypeDictionaryKeyCallBacks CFDictionaryKeyCallBacks
)

var (
	// KCFCopyStringSetCallBacks is predefined [CFSetCallBacks] structure containing a set of callbacks appropriate for use when the values in a set are all CFString objects. The retain callback makes an immutable copy of strings added to the set.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFCopyStringSetCallBacks
	KCFCopyStringSetCallBacks CFSetCallBacks
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFTypeSetCallBacks
	KCFTypeSetCallBacks CFSetCallBacks
)

var (
	// KCFDateFormatterAMSymbol is specifies the AM symbol property, a CFString object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/amSymbol
	KCFDateFormatterAMSymbol CFDateFormatterKey
	// KCFDateFormatterCalendar is specifies the calendar property, a CFCalendar object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/calendar
	KCFDateFormatterCalendar CFDateFormatterKey
	// KCFDateFormatterCalendarName is specifies the calendar name, a CFString object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/calendarName
	KCFDateFormatterCalendarName CFDateFormatterKey
	// KCFDateFormatterDefaultDate is specifies the default date property, a CFDate object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/defaultDate
	KCFDateFormatterDefaultDate CFDateFormatterKey
	// KCFDateFormatterDefaultFormat is the original format string for the formatter (given the date & time style and locale specified at creation).
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/defaultFormat
	KCFDateFormatterDefaultFormat CFDateFormatterKey
	// KCFDateFormatterDoesRelativeDateFormattingKey is specifies the relative date formatting property, a CFBoolean object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/doesRelativeDateFormattingKey
	KCFDateFormatterDoesRelativeDateFormattingKey CFDateFormatterKey
	// KCFDateFormatterEraSymbols is specifies the era symbols property, a CFArray of CFString objects.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/eraSymbols
	KCFDateFormatterEraSymbols CFDateFormatterKey
	// KCFDateFormatterGregorianStartDate is specifies the Gregorian start date property, a CFDate object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/gregorianStartDate
	KCFDateFormatterGregorianStartDate CFDateFormatterKey
	// KCFDateFormatterIsLenient is specifies the lenient property, a CFBoolean object where a true value indicates that the parsing of strings into date or absolute time values will be fuzzy.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/isLenient
	KCFDateFormatterIsLenient CFDateFormatterKey
	// KCFDateFormatterLongEraSymbols is specifies the long era symbols property, a CFArray of CFString objects.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/longEraSymbols
	KCFDateFormatterLongEraSymbols CFDateFormatterKey
	// KCFDateFormatterMonthSymbols is specifies the month symbols property, a CFArray of CFString objects.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/monthSymbols
	KCFDateFormatterMonthSymbols CFDateFormatterKey
	// KCFDateFormatterPMSymbol is specifies the PM symbol property, a CFString object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/pmSymbol
	KCFDateFormatterPMSymbol CFDateFormatterKey
	// KCFDateFormatterQuarterSymbols is specifies the quarter symbols property, a CFArray of CFString objects.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/quarterSymbols
	KCFDateFormatterQuarterSymbols CFDateFormatterKey
	// KCFDateFormatterShortMonthSymbols is specifies the short month symbols property, a CFArray of CFString objects.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/shortMonthSymbols
	KCFDateFormatterShortMonthSymbols CFDateFormatterKey
	// KCFDateFormatterShortQuarterSymbols is specifies the short quarter symbols property, a CFArray of CFString objects.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/shortQuarterSymbols
	KCFDateFormatterShortQuarterSymbols CFDateFormatterKey
	// KCFDateFormatterShortStandaloneMonthSymbols is specifies the short standalone month symbols property, a CFArray of CFString objects.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/shortStandaloneMonthSymbols
	KCFDateFormatterShortStandaloneMonthSymbols CFDateFormatterKey
	// KCFDateFormatterShortStandaloneQuarterSymbols is specifies the short standalone quarter symbols property, a CFArray of CFString objects.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/shortStandaloneQuarterSymbols
	KCFDateFormatterShortStandaloneQuarterSymbols CFDateFormatterKey
	// KCFDateFormatterShortStandaloneWeekdaySymbols is specifies the short standalone weekday symbols property, a CFArray of CFString objects.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/shortStandaloneWeekdaySymbols
	KCFDateFormatterShortStandaloneWeekdaySymbols CFDateFormatterKey
	// KCFDateFormatterShortWeekdaySymbols is specifies the short weekday symbols property, a CFArray of CFString objects.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/shortWeekdaySymbols
	KCFDateFormatterShortWeekdaySymbols CFDateFormatterKey
	// KCFDateFormatterStandaloneMonthSymbols is specifies the standalone month symbols property, a CFArray of CFString objects.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/standaloneMonthSymbols
	KCFDateFormatterStandaloneMonthSymbols CFDateFormatterKey
	// KCFDateFormatterStandaloneQuarterSymbols is specifies the standalone quarter symbols property, a CFArray of CFString objects.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/standaloneQuarterSymbols
	KCFDateFormatterStandaloneQuarterSymbols CFDateFormatterKey
	// KCFDateFormatterStandaloneWeekdaySymbols is specifies the standalone weekday symbols property, a CFArray of CFString objects.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/standaloneWeekdaySymbols
	KCFDateFormatterStandaloneWeekdaySymbols CFDateFormatterKey
	// KCFDateFormatterTimeZone is specifies the time zone property, a CFTimeZone object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/timeZone
	KCFDateFormatterTimeZone CFDateFormatterKey
	// KCFDateFormatterTwoDigitStartDate is specifies the property representing the date from which two-digit years start, a CFDate object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/twoDigitStartDate
	KCFDateFormatterTwoDigitStartDate CFDateFormatterKey
	// KCFDateFormatterVeryShortMonthSymbols is specifies the very short month symbols property, a CFArray of CFString objects.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/veryShortMonthSymbols
	KCFDateFormatterVeryShortMonthSymbols CFDateFormatterKey
	// KCFDateFormatterVeryShortStandaloneMonthSymbols is specifies the very short standalone month symbols property, a CFArray of CFString objects.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/veryShortStandaloneMonthSymbols
	KCFDateFormatterVeryShortStandaloneMonthSymbols CFDateFormatterKey
	// KCFDateFormatterVeryShortStandaloneWeekdaySymbols is specifies the very short standalone weekday symbols property, a CFArray of CFString objects.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/veryShortStandaloneWeekdaySymbols
	KCFDateFormatterVeryShortStandaloneWeekdaySymbols CFDateFormatterKey
	// KCFDateFormatterVeryShortWeekdaySymbols is specifies the very short weekday symbols property, a CFArray of CFString objects.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/veryShortWeekdaySymbols
	KCFDateFormatterVeryShortWeekdaySymbols CFDateFormatterKey
	// KCFDateFormatterWeekdaySymbols is specifies the weekday symbols property, a CFArray of CFString objects.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey/weekdaySymbols
	KCFDateFormatterWeekdaySymbols CFDateFormatterKey
)

var (
	// KCFErrorDomainCocoa is a constant that specified the Cocoa domain.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFErrorDomainCocoa
	KCFErrorDomainCocoa CFErrorDomain
	// KCFErrorDomainMach is a constant that specified the Mach domain.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFErrorDomainMach
	KCFErrorDomainMach CFErrorDomain
	// KCFErrorDomainOSStatus is a constant that specified the OS domain.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFErrorDomainOSStatus
	KCFErrorDomainOSStatus CFErrorDomain
	// KCFErrorDomainPOSIX is a constant that specified the POSIX domain.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFErrorDomainPOSIX
	KCFErrorDomainPOSIX CFErrorDomain
)

var (
	// KCFLocaleAlternateQuotationBeginDelimiterKey is specifies the alternating begin quotation symbol associated with the locale. In some locales, when quotations are nested, the quotation characters alternate. Thus, [NSLocaleQuotationBeginDelimiterKey], then [NSLocaleAlternateQuotationBeginDelimiterKey], and so on.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleKey/alternateQuotationBeginDelimiterKey
	KCFLocaleAlternateQuotationBeginDelimiterKey CFLocaleKey
	// KCFLocaleAlternateQuotationEndDelimiterKey is specifies the alternating end quotation symbol associated with the locale. In some locales, when quotations are nested, the quotation characters alternate. Thus, [NSLocaleQuotationEndDelimiterKey], then [NSLocaleAlternateQuotationEndDelimiterKey], and so on.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleKey/alternateQuotationEndDelimiterKey
	KCFLocaleAlternateQuotationEndDelimiterKey CFLocaleKey
	// KCFLocaleCalendar is specifies the locale calendar.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleKey/calendar
	KCFLocaleCalendar CFLocaleKey
	// KCFLocaleCalendarIdentifier is specifies the locale calendar identifier.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleKey/calendarIdentifier
	KCFLocaleCalendarIdentifier CFLocaleKey
	// KCFLocaleCollationIdentifier is specifies the locale collation identifier.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleKey/collationIdentifier
	KCFLocaleCollationIdentifier CFLocaleKey
	// KCFLocaleCollatorIdentifier is specifies the collation identifier for the locale.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleKey/collatorIdentifier
	KCFLocaleCollatorIdentifier CFLocaleKey
	// KCFLocaleCountryCode is specifies the locale country code.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleKey/countryCode
	KCFLocaleCountryCode CFLocaleKey
	// KCFLocaleCurrencyCode is specifies the locale currency code.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleKey/currencyCode
	KCFLocaleCurrencyCode CFLocaleKey
	// KCFLocaleCurrencySymbol is specifies the currency symbol.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleKey/currencySymbol
	KCFLocaleCurrencySymbol CFLocaleKey
	// KCFLocaleDecimalSeparator is specifies the decimal point string.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleKey/decimalSeparator
	KCFLocaleDecimalSeparator CFLocaleKey
	// KCFLocaleExemplarCharacterSet is specifies the locale character set.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleKey/exemplarCharacterSet
	KCFLocaleExemplarCharacterSet CFLocaleKey
	// KCFLocaleGroupingSeparator is specifies the separator string between groups of digits.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleKey/groupingSeparator
	KCFLocaleGroupingSeparator CFLocaleKey
	// KCFLocaleIdentifier is specifies locale identifier.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleKey/identifier
	KCFLocaleIdentifier CFLocaleKey
	// KCFLocaleLanguageCode is specifies the locale language code.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleKey/languageCode
	KCFLocaleLanguageCode CFLocaleKey
	// KCFLocaleMeasurementSystem is specifies the measurement system used.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleKey/measurementSystem
	KCFLocaleMeasurementSystem CFLocaleKey
	// KCFLocaleQuotationBeginDelimiterKey is specifies the begin quotation symbol associated with the locale.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleKey/quotationBeginDelimiterKey
	KCFLocaleQuotationBeginDelimiterKey CFLocaleKey
	// KCFLocaleQuotationEndDelimiterKey is specifies the end quotation symbol associated with the locale.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleKey/quotationEndDelimiterKey
	KCFLocaleQuotationEndDelimiterKey CFLocaleKey
	// KCFLocaleScriptCode is specifies the locale script code.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleKey/scriptCode
	KCFLocaleScriptCode CFLocaleKey
	// KCFLocaleUsesMetricSystem is specifies the whether the locale uses the metric system.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleKey/usesMetricSystem
	KCFLocaleUsesMetricSystem CFLocaleKey
	// KCFLocaleVariantCode is specifies the locale variant code.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleKey/variantCode
	KCFLocaleVariantCode CFLocaleKey
)

var (
	// KCFLocaleCurrentLocaleDidChangeNotification is identifier for the notification sent if the current locale changes.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNotificationName/cfLocaleCurrentLocaleDidChange
	KCFLocaleCurrentLocaleDidChangeNotification CFNotificationName
	// KCFTimeZoneSystemTimeZoneDidChangeNotification is name of the notification posted when the system time zone changes.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNotificationName/cfTimeZoneSystemTimeZoneDidChange
	KCFTimeZoneSystemTimeZoneDidChangeNotification CFNotificationName
)

var (
	// KCFNull is the singleton CFNull object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFNull
	KCFNull CFNullRef
)

var (
	// KCFNumberFormatterAlwaysShowDecimalSeparator is specifies if the result of converting a value to a string should always contain the decimal separator, even if the number is an integer.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/alwaysShowDecimalSeparator
	KCFNumberFormatterAlwaysShowDecimalSeparator CFNumberFormatterKey
	// KCFNumberFormatterCurrencyCode is specifies the currency code, a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/currencyCode
	KCFNumberFormatterCurrencyCode CFNumberFormatterKey
	// KCFNumberFormatterCurrencyDecimalSeparator is specifies the currency decimal separator, a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/currencyDecimalSeparator
	KCFNumberFormatterCurrencyDecimalSeparator CFNumberFormatterKey
	// KCFNumberFormatterCurrencyGroupingSeparator is specifies the grouping symbol to use when placing a currency value within a string, a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/currencyGroupingSeparator
	KCFNumberFormatterCurrencyGroupingSeparator CFNumberFormatterKey
	// KCFNumberFormatterCurrencySymbol is specifies the symbol for the currency, a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/currencySymbol
	KCFNumberFormatterCurrencySymbol CFNumberFormatterKey
	// KCFNumberFormatterDecimalSeparator is specifies the decimal separator, a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/decimalSeparator
	KCFNumberFormatterDecimalSeparator CFNumberFormatterKey
	// KCFNumberFormatterDefaultFormat is the original format string for the formatter (given the date and time style and locale specified at creation), a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/defaultFormat
	KCFNumberFormatterDefaultFormat CFNumberFormatterKey
	// KCFNumberFormatterExponentSymbol is specifies the exponent symbol (“E” or “e”) in the scientific notation of numbers (for example, as in `1.0e+56`), a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/exponentSymbol
	KCFNumberFormatterExponentSymbol CFNumberFormatterKey
	// KCFNumberFormatterFormatWidth is specifies the width of a formatted number within a string that is either left justified or right justified based on the value of [paddingPosition], a [CFNumber] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/formatWidth
	KCFNumberFormatterFormatWidth CFNumberFormatterKey
	// KCFNumberFormatterGroupingSeparator is specifies the grouping separator, a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/groupingSeparator
	KCFNumberFormatterGroupingSeparator CFNumberFormatterKey
	// KCFNumberFormatterGroupingSize is specifies how often the “thousands” or grouping separator appears, as in “10,000,000”, a [CFNumber] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/groupingSize
	KCFNumberFormatterGroupingSize CFNumberFormatterKey
	// KCFNumberFormatterInfinitySymbol is specifies the string that is used to represent the symbol for infinity, a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/infinitySymbol
	KCFNumberFormatterInfinitySymbol CFNumberFormatterKey
	// KCFNumberFormatterInternationalCurrencySymbol is specifies the international currency symbol to use when placing a formatted number within a string, a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/internationalCurrencySymbol
	KCFNumberFormatterInternationalCurrencySymbol CFNumberFormatterKey
	// KCFNumberFormatterIsLenient is specifies whether the formatter is lenient, a[CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/isLenient
	KCFNumberFormatterIsLenient CFNumberFormatterKey
	// KCFNumberFormatterMaxFractionDigits is specifies the maximum number of digits after a decimal point, a [CFNumber] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/maxFractionDigits
	KCFNumberFormatterMaxFractionDigits CFNumberFormatterKey
	// KCFNumberFormatterMaxIntegerDigits is specifies the maximum number of integer digits before a decimal point, a [CFNumber] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/maxIntegerDigits
	KCFNumberFormatterMaxIntegerDigits CFNumberFormatterKey
	// KCFNumberFormatterMaxSignificantDigits is specifies the maximum number of significant digits to use, a[CFNumber] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/maxSignificantDigits
	KCFNumberFormatterMaxSignificantDigits CFNumberFormatterKey
	// KCFNumberFormatterMinFractionDigits is specifies the minimum number of digits after a decimal point, a [CFNumber] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/minFractionDigits
	KCFNumberFormatterMinFractionDigits CFNumberFormatterKey
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/minGroupingDigits
	KCFNumberFormatterMinGroupingDigits CFNumberFormatterKey
	// KCFNumberFormatterMinIntegerDigits is specifies the minimum number of integer digits before a decimal point, a [CFNumber] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/minIntegerDigits
	KCFNumberFormatterMinIntegerDigits CFNumberFormatterKey
	// KCFNumberFormatterMinSignificantDigits is specifies the minimum number of significant digits to use, a[CFNumber] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/minSignificantDigits
	KCFNumberFormatterMinSignificantDigits CFNumberFormatterKey
	// KCFNumberFormatterMinusSign is specifies the symbol for the minus sign, a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/minusSign
	KCFNumberFormatterMinusSign CFNumberFormatterKey
	// KCFNumberFormatterMultiplier is specifies the multiplier to use when placing a formatted number within a string, a [CFNumber] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/multiplier
	KCFNumberFormatterMultiplier CFNumberFormatterKey
	// KCFNumberFormatterNaNSymbol is specifies the string that is used to represent NaN (“not a number”) when values are converted to strings, a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/naNSymbol
	KCFNumberFormatterNaNSymbol CFNumberFormatterKey
	// KCFNumberFormatterNegativePrefix is specifies the minus sign prefix symbol to use when placing a formatted number within a string, a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/negativePrefix
	KCFNumberFormatterNegativePrefix CFNumberFormatterKey
	// KCFNumberFormatterNegativeSuffix is specifies the minus sign suffix symbol to use when placing a formatted number within a string, a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/negativeSuffix
	KCFNumberFormatterNegativeSuffix CFNumberFormatterKey
	// KCFNumberFormatterPaddingCharacter is specifies the padding character to use when placing a formatted number within a string, a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/paddingCharacter
	KCFNumberFormatterPaddingCharacter CFNumberFormatterKey
	// KCFNumberFormatterPaddingPosition is specifies the position of a formatted number within a string, a [CFNumber] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/paddingPosition
	KCFNumberFormatterPaddingPosition CFNumberFormatterKey
	// KCFNumberFormatterPerMillSymbol is specifies the per mill (1/1000) symbol to use when placing a formatted number within a string, a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/perMillSymbol
	KCFNumberFormatterPerMillSymbol CFNumberFormatterKey
	// KCFNumberFormatterPercentSymbol is specifies the string that is used to represent the percent symbol, a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/percentSymbol
	KCFNumberFormatterPercentSymbol CFNumberFormatterKey
	// KCFNumberFormatterPlusSign is specifies the symbol for the plus sign, a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/plusSign
	KCFNumberFormatterPlusSign CFNumberFormatterKey
	// KCFNumberFormatterPositivePrefix is specifies the plus sign prefix symbol to use when placing a formatted number within a string, a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/positivePrefix
	KCFNumberFormatterPositivePrefix CFNumberFormatterKey
	// KCFNumberFormatterPositiveSuffix is specifies the plus sign suffix symbol to use when placing a formatted number within a string, a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/positiveSuffix
	KCFNumberFormatterPositiveSuffix CFNumberFormatterKey
	// KCFNumberFormatterRoundingIncrement is specifies a positive rounding increment, or `0.0` to disable rounding, a [CFNumber] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/roundingIncrement
	KCFNumberFormatterRoundingIncrement CFNumberFormatterKey
	// KCFNumberFormatterRoundingMode is specifies how the last digit is rounded, as when `3.1415926535…` is rounded to three decimal places, as in `3.142`. See [CFNumberFormatterRoundingMode] for possible values.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/roundingMode
	KCFNumberFormatterRoundingMode CFNumberFormatterKey
	// KCFNumberFormatterSecondaryGroupingSize is specifies how often the secondary grouping separator appears, a [CFNumber] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/secondaryGroupingSize
	KCFNumberFormatterSecondaryGroupingSize CFNumberFormatterKey
	// KCFNumberFormatterUseGroupingSeparator is specifies if the grouping separator should be used, a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/useGroupingSeparator
	KCFNumberFormatterUseGroupingSeparator CFNumberFormatterKey
	// KCFNumberFormatterUseSignificantDigits is specifies the whether the formatter uses significant digits, a [CFBoolean] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/useSignificantDigits
	KCFNumberFormatterUseSignificantDigits CFNumberFormatterKey
	// KCFNumberFormatterZeroSymbol is specifies the string that is used to represent zero, a [CFString] object.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey/zeroSymbol
	KCFNumberFormatterZeroSymbol CFNumberFormatterKey
)

var (
	// KCFNumberNaN is “Not a Number.” This value is often the result of an invalid operation, such as the square-root of a negative number.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFNumberNaN
	KCFNumberNaN CFNumberRef
	// KCFNumberNegativeInfinity is designates a negative infinity value.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFNumberNegativeInfinity
	KCFNumberNegativeInfinity CFNumberRef
	// KCFNumberPositiveInfinity is designates a positive infinity value.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFNumberPositiveInfinity
	KCFNumberPositiveInfinity CFNumberRef
)

var (
	// KCFRunLoopCommonModes is objects added to a run loop using this value as the mode are monitored by all run loop modes that have been declared as a member of the set of “common” modes with [CFRunLoopAddCommonMode(_:_:)].
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopMode/commonModes
	KCFRunLoopCommonModes CFRunLoopMode
	// KCFRunLoopDefaultMode is run loop mode that should be used when a thread is in its default, or idle, state, waiting for an event. This mode is used when the run loop is started with [CFRunLoopRun()].
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopMode/defaultMode
	KCFRunLoopDefaultMode CFRunLoopMode
)

var (
	// KCFStreamPropertyAppendToFile is value is a [CFBoolean] value that indicates whether to append the written data to a file, if it already exists, rather than to replace its contents.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFStreamPropertyKey/appendToFile
	KCFStreamPropertyAppendToFile CFStreamPropertyKey
	// KCFStreamPropertyDataWritten is value is a [CFData] object that contains all the bytes written to a writable memory stream. You cannot modify this value.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFStreamPropertyKey/dataWritten
	KCFStreamPropertyDataWritten CFStreamPropertyKey
	// KCFStreamPropertyFileCurrentOffset is value is a [CFNumber] object containing the current file offset.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFStreamPropertyKey/fileCurrentOffset
	KCFStreamPropertyFileCurrentOffset CFStreamPropertyKey
	// KCFStreamPropertySocketNativeHandle is value is a [CFData] object that contains the native handle for a socket stream—of type [CFSocketNativeHandle]—to which the socket stream is connected.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFStreamPropertyKey/socketNativeHandle
	KCFStreamPropertySocketNativeHandle CFStreamPropertyKey
	// KCFStreamPropertySocketRemoteHostName is value is a [CFString] object containing the name of the host to which the socket stream is connected or [NULL] if unknown.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFStreamPropertyKey/socketRemoteHostName
	KCFStreamPropertySocketRemoteHostName CFStreamPropertyKey
	// KCFStreamPropertySocketRemotePortNumber is value is a [CFNumber] object containing the remote port number to which the socket stream is connected or [NULL] if unknown.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/CFStreamPropertyKey/socketRemotePortNumber
	KCFStreamPropertySocketRemotePortNumber CFStreamPropertyKey
)

var (
	// KCFStringBinaryHeapCallBacks is predefined [CFBinaryHeapCallBacks] structure containing a set of callbacks appropriate for use when the values in a binary heap are all [CFString] objects.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFStringBinaryHeapCallBacks
	KCFStringBinaryHeapCallBacks CFBinaryHeapCallBacks
)

var (
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFTypeArrayCallBacks
	KCFTypeArrayCallBacks CFArrayCallBacks
)

var (
	// KCFTypeDictionaryValueCallBacks is predefined [CFDictionaryValueCallBacks] structure containing a set of callbacks appropriate for use when the values in a CFDictionary are all CFType-derived objects.
	//
	// See: https://developer.apple.com/documentation/CoreFoundation/kCFTypeDictionaryValueCallBacks
	KCFTypeDictionaryValueCallBacks CFDictionaryValueCallBacks
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFAllocatorDefault"); err == nil && ptr != 0 {
		KCFAllocatorDefault = *(*CFAllocatorRef)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFAllocatorMalloc"); err == nil && ptr != 0 {
		KCFAllocatorMalloc = *(*CFAllocatorRef)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFAllocatorMallocZone"); err == nil && ptr != 0 {
		KCFAllocatorMallocZone = *(*CFAllocatorRef)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFAllocatorNull"); err == nil && ptr != 0 {
		KCFAllocatorNull = *(*CFAllocatorRef)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFAllocatorSystemDefault"); err == nil && ptr != 0 {
		KCFAllocatorSystemDefault = *(*CFAllocatorRef)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFAllocatorUseContext"); err == nil && ptr != 0 {
		KCFAllocatorUseContext = *(*CFAllocatorRef)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFBanglaCalendar"); err == nil && ptr != 0 {
		KCFBanglaCalendar = *(*CFCalendarIdentifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFBooleanFalse"); err == nil && ptr != 0 {
		KCFBooleanFalse = *(*CFBooleanRef)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFBooleanTrue"); err == nil && ptr != 0 {
		KCFBooleanTrue = *(*CFBooleanRef)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFBuddhistCalendar"); err == nil && ptr != 0 {
		KCFBuddhistCalendar = *(*CFCalendarIdentifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFBundleDevelopmentRegionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFBundleDevelopmentRegionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFBundleExecutableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFBundleExecutableKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFBundleIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFBundleIdentifierKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFBundleInfoDictionaryVersionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFBundleInfoDictionaryVersionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFBundleLocalizationsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFBundleLocalizationsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFBundleNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFBundleNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFBundleVersionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFBundleVersionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFChineseCalendar"); err == nil && ptr != 0 {
		KCFChineseCalendar = *(*CFCalendarIdentifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFCopyStringBagCallBacks"); err == nil && ptr != 0 {
		KCFCopyStringBagCallBacks = *(*CFBagCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFCopyStringDictionaryKeyCallBacks"); err == nil && ptr != 0 {
		KCFCopyStringDictionaryKeyCallBacks = *(*CFDictionaryKeyCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFCopyStringSetCallBacks"); err == nil && ptr != 0 {
		KCFCopyStringSetCallBacks = *(*CFSetCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDangiCalendar"); err == nil && ptr != 0 {
		KCFDangiCalendar = *(*CFCalendarIdentifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterAMSymbol"); err == nil && ptr != 0 {
		KCFDateFormatterAMSymbol = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterCalendar"); err == nil && ptr != 0 {
		KCFDateFormatterCalendar = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterCalendarName"); err == nil && ptr != 0 {
		KCFDateFormatterCalendarName = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterDefaultDate"); err == nil && ptr != 0 {
		KCFDateFormatterDefaultDate = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterDefaultFormat"); err == nil && ptr != 0 {
		KCFDateFormatterDefaultFormat = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterDoesRelativeDateFormattingKey"); err == nil && ptr != 0 {
		KCFDateFormatterDoesRelativeDateFormattingKey = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterEraSymbols"); err == nil && ptr != 0 {
		KCFDateFormatterEraSymbols = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterGregorianStartDate"); err == nil && ptr != 0 {
		KCFDateFormatterGregorianStartDate = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterIsLenient"); err == nil && ptr != 0 {
		KCFDateFormatterIsLenient = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterLongEraSymbols"); err == nil && ptr != 0 {
		KCFDateFormatterLongEraSymbols = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterMonthSymbols"); err == nil && ptr != 0 {
		KCFDateFormatterMonthSymbols = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterPMSymbol"); err == nil && ptr != 0 {
		KCFDateFormatterPMSymbol = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterQuarterSymbols"); err == nil && ptr != 0 {
		KCFDateFormatterQuarterSymbols = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterShortMonthSymbols"); err == nil && ptr != 0 {
		KCFDateFormatterShortMonthSymbols = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterShortQuarterSymbols"); err == nil && ptr != 0 {
		KCFDateFormatterShortQuarterSymbols = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterShortStandaloneMonthSymbols"); err == nil && ptr != 0 {
		KCFDateFormatterShortStandaloneMonthSymbols = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterShortStandaloneQuarterSymbols"); err == nil && ptr != 0 {
		KCFDateFormatterShortStandaloneQuarterSymbols = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterShortStandaloneWeekdaySymbols"); err == nil && ptr != 0 {
		KCFDateFormatterShortStandaloneWeekdaySymbols = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterShortWeekdaySymbols"); err == nil && ptr != 0 {
		KCFDateFormatterShortWeekdaySymbols = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterStandaloneMonthSymbols"); err == nil && ptr != 0 {
		KCFDateFormatterStandaloneMonthSymbols = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterStandaloneQuarterSymbols"); err == nil && ptr != 0 {
		KCFDateFormatterStandaloneQuarterSymbols = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterStandaloneWeekdaySymbols"); err == nil && ptr != 0 {
		KCFDateFormatterStandaloneWeekdaySymbols = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterTimeZone"); err == nil && ptr != 0 {
		KCFDateFormatterTimeZone = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterTwoDigitStartDate"); err == nil && ptr != 0 {
		KCFDateFormatterTwoDigitStartDate = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterVeryShortMonthSymbols"); err == nil && ptr != 0 {
		KCFDateFormatterVeryShortMonthSymbols = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterVeryShortStandaloneMonthSymbols"); err == nil && ptr != 0 {
		KCFDateFormatterVeryShortStandaloneMonthSymbols = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterVeryShortStandaloneWeekdaySymbols"); err == nil && ptr != 0 {
		KCFDateFormatterVeryShortStandaloneWeekdaySymbols = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterVeryShortWeekdaySymbols"); err == nil && ptr != 0 {
		KCFDateFormatterVeryShortWeekdaySymbols = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFDateFormatterWeekdaySymbols"); err == nil && ptr != 0 {
		KCFDateFormatterWeekdaySymbols = *(*CFDateFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFErrorDescriptionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFErrorDescriptionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFErrorDomainCocoa"); err == nil && ptr != 0 {
		KCFErrorDomainCocoa = *(*CFErrorDomain)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFErrorDomainMach"); err == nil && ptr != 0 {
		KCFErrorDomainMach = *(*CFErrorDomain)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFErrorDomainOSStatus"); err == nil && ptr != 0 {
		KCFErrorDomainOSStatus = *(*CFErrorDomain)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFErrorDomainPOSIX"); err == nil && ptr != 0 {
		KCFErrorDomainPOSIX = *(*CFErrorDomain)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFErrorFilePathKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFErrorFilePathKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFErrorLocalizedDescriptionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFErrorLocalizedDescriptionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFErrorLocalizedFailureKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFErrorLocalizedFailureKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFErrorLocalizedFailureReasonKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFErrorLocalizedFailureReasonKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFErrorLocalizedRecoverySuggestionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFErrorLocalizedRecoverySuggestionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFErrorURLKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFErrorURLKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFErrorUnderlyingErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFErrorUnderlyingErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFGregorianCalendar"); err == nil && ptr != 0 {
		KCFGregorianCalendar = *(*CFCalendarIdentifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFGujaratiCalendar"); err == nil && ptr != 0 {
		KCFGujaratiCalendar = *(*CFCalendarIdentifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFHebrewCalendar"); err == nil && ptr != 0 {
		KCFHebrewCalendar = *(*CFCalendarIdentifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFISO8601Calendar"); err == nil && ptr != 0 {
		KCFISO8601Calendar = *(*CFCalendarIdentifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFIndianCalendar"); err == nil && ptr != 0 {
		KCFIndianCalendar = *(*CFCalendarIdentifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFIslamicCalendar"); err == nil && ptr != 0 {
		KCFIslamicCalendar = *(*CFCalendarIdentifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFIslamicCivilCalendar"); err == nil && ptr != 0 {
		KCFIslamicCivilCalendar = *(*CFCalendarIdentifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFIslamicTabularCalendar"); err == nil && ptr != 0 {
		KCFIslamicTabularCalendar = *(*CFCalendarIdentifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFIslamicUmmAlQuraCalendar"); err == nil && ptr != 0 {
		KCFIslamicUmmAlQuraCalendar = *(*CFCalendarIdentifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFJapaneseCalendar"); err == nil && ptr != 0 {
		KCFJapaneseCalendar = *(*CFCalendarIdentifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFKannadaCalendar"); err == nil && ptr != 0 {
		KCFKannadaCalendar = *(*CFCalendarIdentifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFLocaleAlternateQuotationBeginDelimiterKey"); err == nil && ptr != 0 {
		KCFLocaleAlternateQuotationBeginDelimiterKey = *(*CFLocaleKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFLocaleAlternateQuotationEndDelimiterKey"); err == nil && ptr != 0 {
		KCFLocaleAlternateQuotationEndDelimiterKey = *(*CFLocaleKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFLocaleCalendar"); err == nil && ptr != 0 {
		KCFLocaleCalendar = *(*CFLocaleKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFLocaleCalendarIdentifier"); err == nil && ptr != 0 {
		KCFLocaleCalendarIdentifier = *(*CFLocaleKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFLocaleCollationIdentifier"); err == nil && ptr != 0 {
		KCFLocaleCollationIdentifier = *(*CFLocaleKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFLocaleCollatorIdentifier"); err == nil && ptr != 0 {
		KCFLocaleCollatorIdentifier = *(*CFLocaleKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFLocaleCountryCode"); err == nil && ptr != 0 {
		KCFLocaleCountryCode = *(*CFLocaleKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFLocaleCurrencyCode"); err == nil && ptr != 0 {
		KCFLocaleCurrencyCode = *(*CFLocaleKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFLocaleCurrencySymbol"); err == nil && ptr != 0 {
		KCFLocaleCurrencySymbol = *(*CFLocaleKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFLocaleCurrentLocaleDidChangeNotification"); err == nil && ptr != 0 {
		KCFLocaleCurrentLocaleDidChangeNotification = *(*CFNotificationName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFLocaleDecimalSeparator"); err == nil && ptr != 0 {
		KCFLocaleDecimalSeparator = *(*CFLocaleKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFLocaleExemplarCharacterSet"); err == nil && ptr != 0 {
		KCFLocaleExemplarCharacterSet = *(*CFLocaleKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFLocaleGroupingSeparator"); err == nil && ptr != 0 {
		KCFLocaleGroupingSeparator = *(*CFLocaleKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFLocaleIdentifier"); err == nil && ptr != 0 {
		KCFLocaleIdentifier = *(*CFLocaleKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFLocaleLanguageCode"); err == nil && ptr != 0 {
		KCFLocaleLanguageCode = *(*CFLocaleKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFLocaleMeasurementSystem"); err == nil && ptr != 0 {
		KCFLocaleMeasurementSystem = *(*CFLocaleKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFLocaleQuotationBeginDelimiterKey"); err == nil && ptr != 0 {
		KCFLocaleQuotationBeginDelimiterKey = *(*CFLocaleKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFLocaleQuotationEndDelimiterKey"); err == nil && ptr != 0 {
		KCFLocaleQuotationEndDelimiterKey = *(*CFLocaleKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFLocaleScriptCode"); err == nil && ptr != 0 {
		KCFLocaleScriptCode = *(*CFLocaleKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFLocaleUsesMetricSystem"); err == nil && ptr != 0 {
		KCFLocaleUsesMetricSystem = *(*CFLocaleKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFLocaleVariantCode"); err == nil && ptr != 0 {
		KCFLocaleVariantCode = *(*CFLocaleKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFMalayalamCalendar"); err == nil && ptr != 0 {
		KCFMalayalamCalendar = *(*CFCalendarIdentifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFMarathiCalendar"); err == nil && ptr != 0 {
		KCFMarathiCalendar = *(*CFCalendarIdentifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNull"); err == nil && ptr != 0 {
		KCFNull = *(*CFNullRef)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterAlwaysShowDecimalSeparator"); err == nil && ptr != 0 {
		KCFNumberFormatterAlwaysShowDecimalSeparator = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterCurrencyCode"); err == nil && ptr != 0 {
		KCFNumberFormatterCurrencyCode = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterCurrencyDecimalSeparator"); err == nil && ptr != 0 {
		KCFNumberFormatterCurrencyDecimalSeparator = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterCurrencyGroupingSeparator"); err == nil && ptr != 0 {
		KCFNumberFormatterCurrencyGroupingSeparator = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterCurrencySymbol"); err == nil && ptr != 0 {
		KCFNumberFormatterCurrencySymbol = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterDecimalSeparator"); err == nil && ptr != 0 {
		KCFNumberFormatterDecimalSeparator = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterDefaultFormat"); err == nil && ptr != 0 {
		KCFNumberFormatterDefaultFormat = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterExponentSymbol"); err == nil && ptr != 0 {
		KCFNumberFormatterExponentSymbol = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterFormatWidth"); err == nil && ptr != 0 {
		KCFNumberFormatterFormatWidth = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterGroupingSeparator"); err == nil && ptr != 0 {
		KCFNumberFormatterGroupingSeparator = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterGroupingSize"); err == nil && ptr != 0 {
		KCFNumberFormatterGroupingSize = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterInfinitySymbol"); err == nil && ptr != 0 {
		KCFNumberFormatterInfinitySymbol = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterInternationalCurrencySymbol"); err == nil && ptr != 0 {
		KCFNumberFormatterInternationalCurrencySymbol = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterIsLenient"); err == nil && ptr != 0 {
		KCFNumberFormatterIsLenient = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterMaxFractionDigits"); err == nil && ptr != 0 {
		KCFNumberFormatterMaxFractionDigits = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterMaxIntegerDigits"); err == nil && ptr != 0 {
		KCFNumberFormatterMaxIntegerDigits = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterMaxSignificantDigits"); err == nil && ptr != 0 {
		KCFNumberFormatterMaxSignificantDigits = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterMinFractionDigits"); err == nil && ptr != 0 {
		KCFNumberFormatterMinFractionDigits = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterMinGroupingDigits"); err == nil && ptr != 0 {
		KCFNumberFormatterMinGroupingDigits = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterMinIntegerDigits"); err == nil && ptr != 0 {
		KCFNumberFormatterMinIntegerDigits = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterMinSignificantDigits"); err == nil && ptr != 0 {
		KCFNumberFormatterMinSignificantDigits = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterMinusSign"); err == nil && ptr != 0 {
		KCFNumberFormatterMinusSign = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterMultiplier"); err == nil && ptr != 0 {
		KCFNumberFormatterMultiplier = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterNaNSymbol"); err == nil && ptr != 0 {
		KCFNumberFormatterNaNSymbol = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterNegativePrefix"); err == nil && ptr != 0 {
		KCFNumberFormatterNegativePrefix = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterNegativeSuffix"); err == nil && ptr != 0 {
		KCFNumberFormatterNegativeSuffix = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterPaddingCharacter"); err == nil && ptr != 0 {
		KCFNumberFormatterPaddingCharacter = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterPaddingPosition"); err == nil && ptr != 0 {
		KCFNumberFormatterPaddingPosition = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterPerMillSymbol"); err == nil && ptr != 0 {
		KCFNumberFormatterPerMillSymbol = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterPercentSymbol"); err == nil && ptr != 0 {
		KCFNumberFormatterPercentSymbol = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterPlusSign"); err == nil && ptr != 0 {
		KCFNumberFormatterPlusSign = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterPositivePrefix"); err == nil && ptr != 0 {
		KCFNumberFormatterPositivePrefix = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterPositiveSuffix"); err == nil && ptr != 0 {
		KCFNumberFormatterPositiveSuffix = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterRoundingIncrement"); err == nil && ptr != 0 {
		KCFNumberFormatterRoundingIncrement = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterRoundingMode"); err == nil && ptr != 0 {
		KCFNumberFormatterRoundingMode = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterSecondaryGroupingSize"); err == nil && ptr != 0 {
		KCFNumberFormatterSecondaryGroupingSize = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterUseGroupingSeparator"); err == nil && ptr != 0 {
		KCFNumberFormatterUseGroupingSeparator = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterUseSignificantDigits"); err == nil && ptr != 0 {
		KCFNumberFormatterUseSignificantDigits = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberFormatterZeroSymbol"); err == nil && ptr != 0 {
		KCFNumberFormatterZeroSymbol = *(*CFNumberFormatterKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberNaN"); err == nil && ptr != 0 {
		KCFNumberNaN = *(*CFNumberRef)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberNegativeInfinity"); err == nil && ptr != 0 {
		KCFNumberNegativeInfinity = *(*CFNumberRef)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFNumberPositiveInfinity"); err == nil && ptr != 0 {
		KCFNumberPositiveInfinity = *(*CFNumberRef)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFOdiaCalendar"); err == nil && ptr != 0 {
		KCFOdiaCalendar = *(*CFCalendarIdentifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFPersianCalendar"); err == nil && ptr != 0 {
		KCFPersianCalendar = *(*CFCalendarIdentifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFPlugInDynamicRegisterFunctionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFPlugInDynamicRegisterFunctionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFPlugInDynamicRegistrationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFPlugInDynamicRegistrationKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFPlugInFactoriesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFPlugInFactoriesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFPlugInTypesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFPlugInTypesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFPlugInUnloadFunctionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFPlugInUnloadFunctionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFPreferencesAnyApplication"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFPreferencesAnyApplication = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFPreferencesAnyHost"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFPreferencesAnyHost = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFPreferencesAnyUser"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFPreferencesAnyUser = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFPreferencesCurrentApplication"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFPreferencesCurrentApplication = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFPreferencesCurrentHost"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFPreferencesCurrentHost = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFPreferencesCurrentUser"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFPreferencesCurrentUser = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFRepublicOfChinaCalendar"); err == nil && ptr != 0 {
		KCFRepublicOfChinaCalendar = *(*CFCalendarIdentifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFRunLoopCommonModes"); err == nil && ptr != 0 {
		KCFRunLoopCommonModes = *(*CFRunLoopMode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFRunLoopDefaultMode"); err == nil && ptr != 0 {
		KCFRunLoopDefaultMode = *(*CFRunLoopMode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFSocketCommandKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFSocketCommandKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFSocketErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFSocketErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFSocketNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFSocketNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFSocketRegisterCommand"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFSocketRegisterCommand = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFSocketResultKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFSocketResultKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFSocketRetrieveCommand"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFSocketRetrieveCommand = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFSocketValueKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFSocketValueKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStreamPropertyAppendToFile"); err == nil && ptr != 0 {
		KCFStreamPropertyAppendToFile = *(*CFStreamPropertyKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStreamPropertyDataWritten"); err == nil && ptr != 0 {
		KCFStreamPropertyDataWritten = *(*CFStreamPropertyKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStreamPropertyFileCurrentOffset"); err == nil && ptr != 0 {
		KCFStreamPropertyFileCurrentOffset = *(*CFStreamPropertyKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStreamPropertySOCKSPassword"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStreamPropertySOCKSPassword = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStreamPropertySOCKSProxy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStreamPropertySOCKSProxy = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStreamPropertySOCKSProxyHost"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStreamPropertySOCKSProxyHost = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStreamPropertySOCKSProxyPort"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStreamPropertySOCKSProxyPort = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStreamPropertySOCKSUser"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStreamPropertySOCKSUser = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStreamPropertySOCKSVersion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStreamPropertySOCKSVersion = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStreamPropertyShouldCloseNativeSocket"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStreamPropertyShouldCloseNativeSocket = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStreamPropertySocketNativeHandle"); err == nil && ptr != 0 {
		KCFStreamPropertySocketNativeHandle = *(*CFStreamPropertyKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStreamPropertySocketRemoteHostName"); err == nil && ptr != 0 {
		KCFStreamPropertySocketRemoteHostName = *(*CFStreamPropertyKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStreamPropertySocketRemotePortNumber"); err == nil && ptr != 0 {
		KCFStreamPropertySocketRemotePortNumber = *(*CFStreamPropertyKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStreamPropertySocketSecurityLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStreamPropertySocketSecurityLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStreamSocketSOCKSVersion4"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStreamSocketSOCKSVersion4 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStreamSocketSOCKSVersion5"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStreamSocketSOCKSVersion5 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStreamSocketSecurityLevelNegotiatedSSL"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStreamSocketSecurityLevelNegotiatedSSL = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStreamSocketSecurityLevelNone"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStreamSocketSecurityLevelNone = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStreamSocketSecurityLevelTLSv1"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStreamSocketSecurityLevelTLSv1 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStringBinaryHeapCallBacks"); err == nil && ptr != 0 {
		KCFStringBinaryHeapCallBacks = *(*CFBinaryHeapCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStringTransformFullwidthHalfwidth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStringTransformFullwidthHalfwidth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStringTransformHiraganaKatakana"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStringTransformHiraganaKatakana = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStringTransformLatinArabic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStringTransformLatinArabic = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStringTransformLatinCyrillic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStringTransformLatinCyrillic = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStringTransformLatinGreek"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStringTransformLatinGreek = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStringTransformLatinHangul"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStringTransformLatinHangul = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStringTransformLatinHebrew"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStringTransformLatinHebrew = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStringTransformLatinHiragana"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStringTransformLatinHiragana = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStringTransformLatinKatakana"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStringTransformLatinKatakana = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStringTransformLatinThai"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStringTransformLatinThai = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStringTransformMandarinLatin"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStringTransformMandarinLatin = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStringTransformStripCombiningMarks"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStringTransformStripCombiningMarks = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStringTransformStripDiacritics"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStringTransformStripDiacritics = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStringTransformToLatin"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStringTransformToLatin = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStringTransformToUnicodeName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStringTransformToUnicodeName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFStringTransformToXMLHex"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFStringTransformToXMLHex = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFTamilCalendar"); err == nil && ptr != 0 {
		KCFTamilCalendar = *(*CFCalendarIdentifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFTeluguCalendar"); err == nil && ptr != 0 {
		KCFTeluguCalendar = *(*CFCalendarIdentifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFTimeZoneSystemTimeZoneDidChangeNotification"); err == nil && ptr != 0 {
		KCFTimeZoneSystemTimeZoneDidChangeNotification = *(*CFNotificationName)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFTypeArrayCallBacks"); err == nil && ptr != 0 {
		KCFTypeArrayCallBacks = *(*CFArrayCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFTypeBagCallBacks"); err == nil && ptr != 0 {
		KCFTypeBagCallBacks = *(*CFBagCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFTypeDictionaryKeyCallBacks"); err == nil && ptr != 0 {
		KCFTypeDictionaryKeyCallBacks = *(*CFDictionaryKeyCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFTypeDictionaryValueCallBacks"); err == nil && ptr != 0 {
		KCFTypeDictionaryValueCallBacks = *(*CFDictionaryValueCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFTypeSetCallBacks"); err == nil && ptr != 0 {
		KCFTypeSetCallBacks = *(*CFSetCallBacks)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLAddedToDirectoryDateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLAddedToDirectoryDateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLApplicationIsScriptableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLApplicationIsScriptableKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLAttributeModificationDateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLAttributeModificationDateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLCanonicalPathKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLCanonicalPathKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLContentAccessDateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLContentAccessDateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLContentModificationDateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLContentModificationDateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLCreationDateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLCreationDateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLDirectoryEntryCountKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLDirectoryEntryCountKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLDocumentIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLDocumentIdentifierKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLFileAllocatedSizeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLFileAllocatedSizeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLFileContentIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLFileContentIdentifierKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLFileIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLFileIdentifierKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLFileResourceIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLFileResourceIdentifierKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLFileResourceTypeBlockSpecial"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLFileResourceTypeBlockSpecial = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLFileResourceTypeCharacterSpecial"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLFileResourceTypeCharacterSpecial = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLFileResourceTypeDirectory"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLFileResourceTypeDirectory = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLFileResourceTypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLFileResourceTypeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLFileResourceTypeNamedPipe"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLFileResourceTypeNamedPipe = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLFileResourceTypeRegular"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLFileResourceTypeRegular = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLFileResourceTypeSocket"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLFileResourceTypeSocket = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLFileResourceTypeSymbolicLink"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLFileResourceTypeSymbolicLink = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLFileResourceTypeUnknown"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLFileResourceTypeUnknown = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLFileSecurityKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLFileSecurityKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLFileSizeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLFileSizeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLGenerationIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLGenerationIdentifierKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLHasHiddenExtensionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLHasHiddenExtensionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLIsAliasFileKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLIsAliasFileKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLIsApplicationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLIsApplicationKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLIsDirectoryKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLIsDirectoryKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLIsExcludedFromBackupKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLIsExcludedFromBackupKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLIsExecutableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLIsExecutableKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLIsHiddenKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLIsHiddenKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLIsMountTriggerKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLIsMountTriggerKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLIsPackageKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLIsPackageKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLIsPurgeableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLIsPurgeableKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLIsReadableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLIsReadableKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLIsRegularFileKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLIsRegularFileKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLIsSparseKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLIsSparseKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLIsSymbolicLinkKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLIsSymbolicLinkKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLIsSystemImmutableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLIsSystemImmutableKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLIsUbiquitousItemKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLIsUbiquitousItemKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLIsUserImmutableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLIsUserImmutableKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLIsVolumeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLIsVolumeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLIsWritableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLIsWritableKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLKeysOfUnsetValuesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLKeysOfUnsetValuesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLLabelNumberKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLLabelNumberKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLLinkCountKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLLinkCountKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLLocalizedLabelKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLLocalizedLabelKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLLocalizedNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLLocalizedNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLLocalizedTypeDescriptionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLLocalizedTypeDescriptionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLMayHaveExtendedAttributesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLMayHaveExtendedAttributesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLMayShareFileContentKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLMayShareFileContentKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLParentDirectoryURLKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLParentDirectoryURLKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLPathKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLPathKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLPreferredIOBlockSizeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLPreferredIOBlockSizeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLQuarantinePropertiesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLQuarantinePropertiesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLTagNamesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLTagNamesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLTotalFileAllocatedSizeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLTotalFileAllocatedSizeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLTotalFileSizeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLTotalFileSizeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLTypeIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLTypeIdentifierKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLUbiquitousItemDownloadingErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLUbiquitousItemDownloadingErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLUbiquitousItemDownloadingStatusCurrent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLUbiquitousItemDownloadingStatusCurrent = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLUbiquitousItemDownloadingStatusDownloaded"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLUbiquitousItemDownloadingStatusDownloaded = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLUbiquitousItemDownloadingStatusKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLUbiquitousItemDownloadingStatusKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLUbiquitousItemDownloadingStatusNotDownloaded"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLUbiquitousItemDownloadingStatusNotDownloaded = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLUbiquitousItemHasUnresolvedConflictsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLUbiquitousItemHasUnresolvedConflictsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLUbiquitousItemIsDownloadingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLUbiquitousItemIsDownloadingKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLUbiquitousItemIsExcludedFromSyncKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLUbiquitousItemIsExcludedFromSyncKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLUbiquitousItemIsSyncPausedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLUbiquitousItemIsSyncPausedKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLUbiquitousItemIsUploadedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLUbiquitousItemIsUploadedKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLUbiquitousItemIsUploadingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLUbiquitousItemIsUploadingKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLUbiquitousItemSupportedSyncControlsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLUbiquitousItemSupportedSyncControlsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLUbiquitousItemUploadingErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLUbiquitousItemUploadingErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeAvailableCapacityForImportantUsageKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeAvailableCapacityForImportantUsageKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeAvailableCapacityForOpportunisticUsageKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeAvailableCapacityForOpportunisticUsageKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeAvailableCapacityKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeAvailableCapacityKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeCreationDateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeCreationDateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeIdentifierKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeIsAutomountedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeIsAutomountedKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeIsBrowsableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeIsBrowsableKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeIsEjectableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeIsEjectableKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeIsEncryptedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeIsEncryptedKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeIsInternalKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeIsInternalKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeIsJournalingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeIsJournalingKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeIsLocalKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeIsLocalKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeIsReadOnlyKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeIsReadOnlyKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeIsRemovableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeIsRemovableKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeIsRootFileSystemKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeIsRootFileSystemKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeLocalizedFormatDescriptionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeLocalizedFormatDescriptionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeLocalizedNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeLocalizedNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeMaximumFileSizeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeMaximumFileSizeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeMountFromLocationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeMountFromLocationKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeResourceCountKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeResourceCountKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeSubtypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeSubtypeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeSupportsAccessPermissionsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeSupportsAccessPermissionsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeSupportsAdvisoryFileLockingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeSupportsAdvisoryFileLockingKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeSupportsCasePreservedNamesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeSupportsCasePreservedNamesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeSupportsCaseSensitiveNamesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeSupportsCaseSensitiveNamesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeSupportsCompressionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeSupportsCompressionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeSupportsExclusiveRenamingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeSupportsExclusiveRenamingKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeSupportsExtendedSecurityKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeSupportsExtendedSecurityKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeSupportsFileCloningKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeSupportsFileCloningKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeSupportsFileProtectionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeSupportsFileProtectionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeSupportsHardLinksKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeSupportsHardLinksKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeSupportsImmutableFilesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeSupportsImmutableFilesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeSupportsJournalingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeSupportsJournalingKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeSupportsPersistentIDsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeSupportsPersistentIDsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeSupportsRenamingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeSupportsRenamingKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeSupportsRootDirectoryDatesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeSupportsRootDirectoryDatesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeSupportsSparseFilesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeSupportsSparseFilesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeSupportsSwapRenamingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeSupportsSwapRenamingKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeSupportsSymbolicLinksKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeSupportsSymbolicLinksKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeSupportsVolumeSizesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeSupportsVolumeSizesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeSupportsZeroRunsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeSupportsZeroRunsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeTotalCapacityKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeTotalCapacityKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeTypeNameKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeTypeNameKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeURLForRemountingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeURLForRemountingKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeURLKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeURLKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFURLVolumeUUIDStringKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFURLVolumeUUIDStringKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFUserNotificationAlertHeaderKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFUserNotificationAlertHeaderKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFUserNotificationAlertMessageKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFUserNotificationAlertMessageKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFUserNotificationAlertTopMostKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFUserNotificationAlertTopMostKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFUserNotificationAlternateButtonTitleKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFUserNotificationAlternateButtonTitleKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFUserNotificationCheckBoxTitlesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFUserNotificationCheckBoxTitlesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFUserNotificationDefaultButtonTitleKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFUserNotificationDefaultButtonTitleKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFUserNotificationIconURLKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFUserNotificationIconURLKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFUserNotificationKeyboardTypesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFUserNotificationKeyboardTypesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFUserNotificationLocalizationURLKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFUserNotificationLocalizationURLKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFUserNotificationOtherButtonTitleKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFUserNotificationOtherButtonTitleKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFUserNotificationPopUpSelectionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFUserNotificationPopUpSelectionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFUserNotificationPopUpTitlesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFUserNotificationPopUpTitlesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFUserNotificationProgressIndicatorValueKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFUserNotificationProgressIndicatorValueKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFUserNotificationSoundURLKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFUserNotificationSoundURLKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFUserNotificationTextFieldTitlesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFUserNotificationTextFieldTitlesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFUserNotificationTextFieldValuesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFUserNotificationTextFieldValuesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFVietnameseCalendar"); err == nil && ptr != 0 {
		KCFVietnameseCalendar = *(*CFCalendarIdentifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFVikramCalendar"); err == nil && ptr != 0 {
		KCFVikramCalendar = *(*CFCalendarIdentifier)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFXMLTreeErrorDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFXMLTreeErrorDescription = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFXMLTreeErrorLineNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFXMLTreeErrorLineNumber = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFXMLTreeErrorLocation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFXMLTreeErrorLocation = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCFXMLTreeErrorStatusCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCFXMLTreeErrorStatusCode = objc.GoString(cstr)
			}
		}
	}

}
