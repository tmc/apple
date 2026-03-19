// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSError] class.
var (
	_NSErrorClass     NSErrorClass
	_NSErrorClassOnce sync.Once
)

func getNSErrorClass() NSErrorClass {
	_NSErrorClassOnce.Do(func() {
		_NSErrorClass = NSErrorClass{class: objc.GetClass("NSError")}
	})
	return _NSErrorClass
}

// GetNSErrorClass returns the class object for NSError.
func GetNSErrorClass() NSErrorClass {
	return getNSErrorClass()
}

type NSErrorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSErrorClass) Alloc() NSError {
	rv := objc.Send[NSError](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// Information about an error condition including a domain, a domain-specific
// error code, and application-specific information.
//
// # Overview
// 
// Objective-C methods can signal an error condition by returning an [NSError]
// object by reference, which provides additional information about the kind
// of error and any underlying cause, if one can be determined. An [NSError]
// object may also provide localized error descriptions suitable for display
// to the user in its user info dictionary. See [Error Handling Programming
// Guide] for more information.
// 
// Methods in Foundation and other Cocoa frameworks most often produce errors
// in the Cocoa error domain ([NSError.NSCocoaErrorDomain]); error codes for the Cocoa
// Error Domain are documented in the [Foundation Constants]. There are also
// predefined domains corresponding to Mach ([NSError.NSMachErrorDomain]), POSIX
// ([NSError.NSPOSIXErrorDomain]), and Carbon ([NSError.NSOSStatusErrorDomain]) errors.
// 
// [NSError] is “toll-free bridged” with its Core Foundation counterpart,
// [CFError]. See [Toll-Free Bridging] for more information.
// 
// # Subclassing Notes
// 
// Applications may choose to create subclasses of [NSError], for example, to
// provide better localized error strings by overriding
// [NSError.LocalizedDescription].
//
// [CFError]: https://developer.apple.com/documentation/CoreFoundation/CFError
// [Error Handling Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ErrorHandlingCocoa/ErrorHandling/ErrorHandling.html#//apple_ref/doc/uid/TP40001806
// [Foundation Constants]: https://developer.apple.com/documentation/Foundation/foundation-constants
// [NSError.NSCocoaErrorDomain]: https://developer.apple.com/documentation/Foundation/NSCocoaErrorDomain
// [NSError.NSMachErrorDomain]: https://developer.apple.com/documentation/Foundation/NSMachErrorDomain
// [NSError.NSOSStatusErrorDomain]: https://developer.apple.com/documentation/Foundation/NSOSStatusErrorDomain
// [NSError.NSPOSIXErrorDomain]: https://developer.apple.com/documentation/Foundation/NSPOSIXErrorDomain
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
//
// # Creating Error Objects
//
//   - [NSError.InitWithDomainCodeUserInfo]: Returns an [NSError] object initialized for a given domain and code with a given `userInfo` dictionary.
//
// # Getting Error Properties
//
//   - [NSError.Code]: The error code.
//   - [NSError.Domain]: A string containing the error domain.
//   - [NSError.UserInfo]: The user info dictionary.
//
// # Getting a Localized Error Description
//
//   - [NSError.LocalizedDescription]: A string containing the localized description of the error.
//   - [NSError.LocalizedRecoveryOptions]: An array containing the localized titles of buttons appropriate for displaying in an alert panel.
//   - [NSError.LocalizedRecoverySuggestion]: A string containing the localized recovery suggestion for the error.
//   - [NSError.LocalizedFailureReason]: A string containing the localized explanation of the reason for the error.
//
// # Getting the Error Recovery Attempter
//
//   - [NSError.RecoveryAttempter]: The object in the user info dictionary corresponding to the [NSRecoveryAttempterErrorKey](<doc://com.apple.foundation/documentation/Foundation/NSRecoveryAttempterErrorKey>) key.
//
// # Displaying a Help Anchor
//
//   - [NSError.HelpAnchor]: A string to display in response to an alert panel help anchor button being pressed.
//
// # Error Domains
//
//   - [NSError.NSCocoaErrorDomain]: Cocoa errors
//   - [NSError.NSPOSIXErrorDomain]: POSIX/BSD errors
//   - [NSError.NSOSStatusErrorDomain]: Mac OS 9/Carbon errors
//   - [NSError.NSMachErrorDomain]: Mach errors
//   - [NSError.NSURLErrorDomain]: URL loading system errors
//   - [NSError.NSStreamSOCKSErrorDomain]: The error domain used by 
//   - [NSError.NSStreamSocketSSLErrorDomain]: The error domain used by 
//
// # Instance Properties
//
//   - [NSError.UnderlyingErrors]
//
// See: https://developer.apple.com/documentation/Foundation/NSError
type NSError struct {
	objectivec.Object
}

// NSErrorFromID constructs a [NSError] from an objc.ID.
//
// Information about an error condition including a domain, a domain-specific
// error code, and application-specific information.
func NSErrorFromID(id objc.ID) NSError {
	return NSError{objectivec.Object{ID: id}}
}
// NOTE: NSError adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSError] class.
//
// # Creating Error Objects
//
//   - [INSError.InitWithDomainCodeUserInfo]: Returns an [NSError] object initialized for a given domain and code with a given `userInfo` dictionary.
//
// # Getting Error Properties
//
//   - [INSError.Code]: The error code.
//   - [INSError.Domain]: A string containing the error domain.
//   - [INSError.UserInfo]: The user info dictionary.
//
// # Getting a Localized Error Description
//
//   - [INSError.LocalizedDescription]: A string containing the localized description of the error.
//   - [INSError.LocalizedRecoveryOptions]: An array containing the localized titles of buttons appropriate for displaying in an alert panel.
//   - [INSError.LocalizedRecoverySuggestion]: A string containing the localized recovery suggestion for the error.
//   - [INSError.LocalizedFailureReason]: A string containing the localized explanation of the reason for the error.
//
// # Getting the Error Recovery Attempter
//
//   - [INSError.RecoveryAttempter]: The object in the user info dictionary corresponding to the [NSRecoveryAttempterErrorKey](<doc://com.apple.foundation/documentation/Foundation/NSRecoveryAttempterErrorKey>) key.
//
// # Displaying a Help Anchor
//
//   - [INSError.HelpAnchor]: A string to display in response to an alert panel help anchor button being pressed.
//
// # Error Domains
//
//   - [INSError.NSCocoaErrorDomain]: Cocoa errors
//   - [INSError.NSPOSIXErrorDomain]: POSIX/BSD errors
//   - [INSError.NSOSStatusErrorDomain]: Mac OS 9/Carbon errors
//   - [INSError.NSMachErrorDomain]: Mach errors
//   - [INSError.NSURLErrorDomain]: URL loading system errors
//   - [INSError.NSStreamSOCKSErrorDomain]: The error domain used by 
//   - [INSError.NSStreamSocketSSLErrorDomain]: The error domain used by 
//
// # Instance Properties
//
//   - [INSError.UnderlyingErrors]
//
// See: https://developer.apple.com/documentation/Foundation/NSError
type INSError interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Creating Error Objects

	// Returns an [NSError] object initialized for a given domain and code with a given `userInfo` dictionary.
	InitWithDomainCodeUserInfo(domain NSErrorDomain, code int, dict INSDictionary) NSError

	// Topic: Getting Error Properties

	// The error code.
	Code() int
	// A string containing the error domain.
	Domain() NSErrorDomain
	// The user info dictionary.
	UserInfo() INSDictionary

	// Topic: Getting a Localized Error Description

	// A string containing the localized description of the error.
	LocalizedDescription() string
	// An array containing the localized titles of buttons appropriate for displaying in an alert panel.
	LocalizedRecoveryOptions() []string
	// A string containing the localized recovery suggestion for the error.
	LocalizedRecoverySuggestion() string
	// A string containing the localized explanation of the reason for the error.
	LocalizedFailureReason() string

	// Topic: Getting the Error Recovery Attempter

	// The object in the user info dictionary corresponding to the [NSRecoveryAttempterErrorKey](<doc://com.apple.foundation/documentation/Foundation/NSRecoveryAttempterErrorKey>) key.
	RecoveryAttempter() objectivec.IObject

	// Topic: Displaying a Help Anchor

	// A string to display in response to an alert panel help anchor button being pressed.
	HelpAnchor() string

	// Topic: Error Domains

	// Cocoa errors
	NSCocoaErrorDomain() string
	// POSIX/BSD errors
	NSPOSIXErrorDomain() string
	// Mac OS 9/Carbon errors
	NSOSStatusErrorDomain() string
	// Mach errors
	NSMachErrorDomain() string
	// URL loading system errors
	NSURLErrorDomain() string
	// The error domain used by 
	NSStreamSOCKSErrorDomain() string
	// The error domain used by 
	NSStreamSocketSSLErrorDomain() string

	// Topic: Instance Properties

	UnderlyingErrors() []NSError

	// The corresponding value is an object that conforms to the NSErrorRecoveryAttempting informal protocol.
	NSRecoveryAttempterErrorKey() string
	Error() string
}

// Init initializes the instance.
func (e NSError) Init() NSError {
	rv := objc.Send[NSError](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e NSError) Autorelease() NSError {
	rv := objc.Send[NSError](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSError creates a new NSError instance.
func NewNSError() NSError {
	class := getNSErrorClass()
	rv := objc.Send[NSError](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewErrorWithCoder(coder INSCoder) NSError {
	instance := getNSErrorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSErrorFromID(rv)
}

// Returns an [NSError] object initialized for a given domain and code with a
// given `userInfo` dictionary.
//
// domain: The error domain—this can be one of the predefined [NSError] domains, or
// an arbitrary string describing a custom domain. `domain` must not be `nil`.
// See `Error Domains` for a list of predefined domains.
//
// code: The error code for the error.
//
// dict: The `userInfo` dictionary for the error. `userInfo` may be `nil`.
//
// # Return Value
// 
// An [NSError] object initialized for `domain` with the specified error
// `code` and the dictionary of arbitrary data `userInfo`.
//
// # Discussion
// 
// This is the designated initializer for [NSError].
//
// See: https://developer.apple.com/documentation/Foundation/NSError/init(domain:code:userInfo:)
func NewErrorWithDomainCodeUserInfo(domain NSErrorDomain, code int, dict INSDictionary) NSError {
	instance := getNSErrorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDomain:code:userInfo:"), objc.String(string(domain)), code, dict)
	return NSErrorFromID(rv)
}

// Returns an [NSError] object initialized for a given domain and code with a
// given `userInfo` dictionary.
//
// domain: The error domain—this can be one of the predefined [NSError] domains, or
// an arbitrary string describing a custom domain. `domain` must not be `nil`.
// See `Error Domains` for a list of predefined domains.
//
// code: The error code for the error.
//
// dict: The `userInfo` dictionary for the error. `userInfo` may be `nil`.
//
// # Return Value
// 
// An [NSError] object initialized for `domain` with the specified error
// `code` and the dictionary of arbitrary data `userInfo`.
//
// # Discussion
// 
// This is the designated initializer for [NSError].
//
// See: https://developer.apple.com/documentation/Foundation/NSError/init(domain:code:userInfo:)
func (e NSError) InitWithDomainCodeUserInfo(domain NSErrorDomain, code int, dict INSDictionary) NSError {
	rv := objc.Send[NSError](e.ID, objc.Sel("initWithDomain:code:userInfo:"), objc.String(string(domain)), code, dict)
	return rv
}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (e NSError) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](e.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (e NSError) InitWithCoder(coder INSCoder) NSError {
	rv := objc.Send[NSError](e.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Specifies a block to call when the corresponding property is not present in
// the user info dictionary.
//
// errorDomain: The error domain of the provider.
//
// provider: A block to be executed synchronously at the time a corresponding property
// is accessed.
// 
// err: The error object that is being accessed. userInfoKey: The user info
// key corresponding to the accessed property.
//
// # Discussion
// 
// This method specifies a block that is called from the implementations of
// [LocalizedDescription], [LocalizedFailureReason],
// [LocalizedRecoverySuggestion], [LocalizedRecoveryOptions],
// [RecoveryAttempter], and [HelpAnchor] when the underlying value for any of
// those properties is not present in the [UserInfo] dictionary of NSError
// instances with the specified domain.
// 
// A user info provider is optional, and allows localization and formatting of
// error messages to be done lazily, rather than populating the [UserInfo] at
// the time of creation. It is expected that only the “owner” of an
// [NSError] domain specifies the provider for the domain, and that this is
// done at most once. This method is not meant for consumers of errors to
// customize the [UserInfo] entries, and should not be used to customize the
// behaviors of error domains provided by the system.
// 
// The keys of a provider’s [UserInfo] dictionary correspond to the queried
// property, such as [NSLocalizedDescriptionKey] for the
// [LocalizedDescription] property. The provider should return `nil` for any
// keys that it is unable to provide, as well as any keys it does not
// recognize (since the list of error keys may be extended in future
// releases). If an appropriate result for the requested key cannot be
// provided, return `nil` rather than choosing to manufacture a generic
// fallback response.
// 
// The provider block is executed synchronously at the time when a
// corresponding property is accessed. The results are not cached.
//
// [NSLocalizedDescriptionKey]: https://developer.apple.com/documentation/Foundation/NSLocalizedDescriptionKey
//
// See: https://developer.apple.com/documentation/Foundation/NSError/setUserInfoValueProvider(forDomain:provider:)
func (_NSErrorClass NSErrorClass) SetUserInfoValueProviderForDomainProvider(errorDomain NSErrorDomain, provider ErrorHandler) {
_block1, _ := NewErrorBlock(provider)
	objc.Send[objc.ID](objc.ID(_NSErrorClass.class), objc.Sel("setUserInfoValueProviderForDomain:provider:"), objc.String(string(errorDomain)), _block1)
}
// Returns any user info provider specified for a given error domain.
//
// errorDomain: The error domain of the user info provider.
//
// # Return Value
// 
// The user info provider of the error domain, or `nil` if none is specified.
//
// See: https://developer.apple.com/documentation/Foundation/NSError/userInfoValueProvider(forDomain:)
func (_NSErrorClass NSErrorClass) UserInfoValueProviderForDomain(errorDomain NSErrorDomain) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NSErrorClass.class), objc.Sel("userInfoValueProviderForDomain:"), objc.String(string(errorDomain)))
	return objectivec.Object{ID: rv}
}
// Returns a properly formatted error object with a
// [NSFileProviderItemCollisionError] error code.
//
// See: https://developer.apple.com/documentation/Foundation/NSError/fileProviderErrorForCollision(with:)
func (_NSErrorClass NSErrorClass) FileProviderErrorForCollisionWithItem(existingItem objectivec.Id) NSError {
	rv := objc.Send[objc.ID](objc.ID(_NSErrorClass.class), objc.Sel("fileProviderErrorForCollisionWithItem:"), existingItem)
	return NSErrorFromID(rv)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSError/fileProviderErrorForNonExistentItem(withIdentifier:)
func (_NSErrorClass NSErrorClass) FileProviderErrorForNonExistentItemWithIdentifier(itemIdentifier INSString) NSError {
	rv := objc.Send[objc.ID](objc.ID(_NSErrorClass.class), objc.Sel("fileProviderErrorForNonExistentItemWithIdentifier:"), itemIdentifier)
	return NSErrorFromID(rv)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSError/fileProviderErrorForRejectedDeletion(of:)
func (_NSErrorClass NSErrorClass) FileProviderErrorForRejectedDeletionOfItem(updatedVersion objectivec.Id) NSError {
	rv := objc.Send[objc.ID](objc.ID(_NSErrorClass.class), objc.Sel("fileProviderErrorForRejectedDeletionOfItem:"), updatedVersion)
	return NSErrorFromID(rv)
}
// Creates and initializes an [NSError] object for a given domain and code
// with a given `userInfo` dictionary.
//
// domain: The error domain—this can be one of the predefined [NSError] domains, or
// an arbitrary string describing a custom domain. `domain` must not be `nil`.
// See `Error Domains` for a list of predefined domains.
//
// code: The error code for the error.
//
// dict: The `userInfo` dictionary for the error. `userInfo` may be `nil`.
//
// # Return Value
// 
// An [NSError] object for `domain` with the specified error `code` and the
// dictionary of arbitrary data `userInfo`.
//
// See: https://developer.apple.com/documentation/Foundation/NSError/errorWithDomain:code:userInfo:
func (_NSErrorClass NSErrorClass) ErrorWithDomainCodeUserInfo(domain NSErrorDomain, code int, dict INSDictionary) NSError {
	rv := objc.Send[objc.ID](objc.ID(_NSErrorClass.class), objc.Sel("errorWithDomain:code:userInfo:"), objc.String(string(domain)), code, dict)
	return NSErrorFromID(rv)
}

// The error code.
//
// # Discussion
// 
// Note that errors are domain-specific.
//
// See: https://developer.apple.com/documentation/Foundation/NSError/code
func (e NSError) Code() int {
	rv := objc.Send[int](e.ID, objc.Sel("code"))
	return rv
}
// A string containing the error domain.
//
// See: https://developer.apple.com/documentation/Foundation/NSError/domain
func (e NSError) Domain() NSErrorDomain {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("domain"))
	return NSErrorDomain(NSStringFromID(rv).String())
}
// The user info dictionary.
//
// # Discussion
// 
// If the user info dictionary has not been set, this property is `nil`.
// 
// On macOS 10.8 or later, if the user info dictionary has not been set, this
// property returns an empty dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/NSError/userInfo
func (e NSError) UserInfo() INSDictionary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("userInfo"))
	return NSDictionaryFromID(objc.ID(rv))
}
// A string containing the localized description of the error.
//
// # Discussion
// 
// The object in the user info dictionary for the key
// [NSLocalizedDescriptionKey]. If the user info dictionary doesn’t contain
// a value for [NSLocalizedDescriptionKey], a default string is constructed
// from the domain and code.
//
// [NSLocalizedDescriptionKey]: https://developer.apple.com/documentation/Foundation/NSLocalizedDescriptionKey
//
// See: https://developer.apple.com/documentation/Foundation/NSError/localizedDescription
func (e NSError) LocalizedDescription() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("localizedDescription"))
	return NSStringFromID(rv).String()
}
// An array containing the localized titles of buttons appropriate for
// displaying in an alert panel.
//
// # Discussion
// 
// The object in the user info dictionary for the key
// [NSLocalizedRecoveryOptionsErrorKey]. If the user info dictionary doesn’t
// contain a value for [NSLocalizedRecoveryOptionsErrorKey], this property is
// `nil`.
// 
// The first string is the title of the right-most and default button, the
// second the one to the left of that, and so on. The recovery options should
// be appropriate for the [LocalizedRecoverySuggestion] property. If the user
// info dictionary doesn’t contain a value for
// [NSLocalizedRecoveryOptionsErrorKey], only an OK button is displayed.
//
// [NSLocalizedRecoveryOptionsErrorKey]: https://developer.apple.com/documentation/Foundation/NSLocalizedRecoveryOptionsErrorKey
//
// See: https://developer.apple.com/documentation/Foundation/NSError/localizedRecoveryOptions
func (e NSError) LocalizedRecoveryOptions() []string {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("localizedRecoveryOptions"))
	return objc.ConvertSliceToStrings(rv)
}
// A string containing the localized recovery suggestion for the error.
//
// # Discussion
// 
// The object in the user info dictionary for the key
// [NSLocalizedRecoverySuggestionErrorKey]. If the user info dictionary
// doesn’t contain a value for [NSLocalizedRecoverySuggestionErrorKey], this
// property is `nil`.
// 
// The returned string is suitable for displaying as the secondary message in
// an alert panel.
//
// [NSLocalizedRecoverySuggestionErrorKey]: https://developer.apple.com/documentation/Foundation/NSLocalizedRecoverySuggestionErrorKey
//
// See: https://developer.apple.com/documentation/Foundation/NSError/localizedRecoverySuggestion
func (e NSError) LocalizedRecoverySuggestion() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("localizedRecoverySuggestion"))
	return NSStringFromID(rv).String()
}
// A string containing the localized explanation of the reason for the error.
//
// # Discussion
// 
// The object in the user info dictionary for the key
// [NSLocalizedFailureReasonErrorKey].
//
// [NSLocalizedFailureReasonErrorKey]: https://developer.apple.com/documentation/Foundation/NSLocalizedFailureReasonErrorKey
//
// See: https://developer.apple.com/documentation/Foundation/NSError/localizedFailureReason
func (e NSError) LocalizedFailureReason() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("localizedFailureReason"))
	return NSStringFromID(rv).String()
}
// The object in the user info dictionary corresponding to the
// [NSRecoveryAttempterErrorKey] key.
//
// [NSRecoveryAttempterErrorKey]: https://developer.apple.com/documentation/Foundation/NSRecoveryAttempterErrorKey
//
// # Discussion
// 
// The recovery attempter must be an instance of a class that conforms to the
// [NSSecureCoding] and NSErrorRecoveryAttempting protocols. It must also be
// able to correctly interpret an index in the [LocalizedRecoveryOptions]
// property.
// 
// If [UserInfo] doesn’t contain a value for [NSRecoveryAttempterErrorKey],
// this property is `nil`.
//
// [NSRecoveryAttempterErrorKey]: https://developer.apple.com/documentation/Foundation/NSRecoveryAttempterErrorKey
//
// See: https://developer.apple.com/documentation/Foundation/NSError/recoveryAttempter
func (e NSError) RecoveryAttempter() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("recoveryAttempter"))
	return objectivec.Object{ID: rv}
}
// A string to display in response to an alert panel help anchor button being
// pressed.
//
// # Discussion
// 
// The object in the user info dictionary for the key [NSHelpAnchorErrorKey].
// If the user info dictionary doesn’t contain a value for
// [NSHelpAnchorErrorKey], this property is `nil`.
// 
// If this property is non-`nil` for an error being presented by
// [init(error:)], the alert panel will include a help anchor button that can
// display this string.
//
// [NSHelpAnchorErrorKey]: https://developer.apple.com/documentation/Foundation/NSHelpAnchorErrorKey
// [init(error:)]: https://developer.apple.com/documentation/AppKit/NSAlert/init(error:)
//
// See: https://developer.apple.com/documentation/Foundation/NSError/helpAnchor
func (e NSError) HelpAnchor() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("helpAnchor"))
	return NSStringFromID(rv).String()
}
// Cocoa errors
//
// See: https://developer.apple.com/documentation/foundation/nscocoaerrordomain
func (e NSError) NSCocoaErrorDomain() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("NSCocoaErrorDomain"))
	return NSStringFromID(rv).String()
}
// POSIX/BSD errors
//
// See: https://developer.apple.com/documentation/foundation/nsposixerrordomain
func (e NSError) NSPOSIXErrorDomain() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("NSPOSIXErrorDomain"))
	return NSStringFromID(rv).String()
}
// Mac OS 9/Carbon errors
//
// See: https://developer.apple.com/documentation/foundation/nsosstatuserrordomain
func (e NSError) NSOSStatusErrorDomain() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("NSOSStatusErrorDomain"))
	return NSStringFromID(rv).String()
}
// Mach errors
//
// See: https://developer.apple.com/documentation/foundation/nsmacherrordomain
func (e NSError) NSMachErrorDomain() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("NSMachErrorDomain"))
	return NSStringFromID(rv).String()
}
// URL loading system errors
//
// See: https://developer.apple.com/documentation/foundation/nsurlerrordomain
func (e NSError) NSURLErrorDomain() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("NSURLErrorDomain"))
	return NSStringFromID(rv).String()
}
// The error domain used by
//
// See: https://developer.apple.com/documentation/foundation/nsstreamsockserrordomain
func (e NSError) NSStreamSOCKSErrorDomain() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("NSStreamSOCKSErrorDomain"))
	return NSStringFromID(rv).String()
}
// The error domain used by
//
// See: https://developer.apple.com/documentation/foundation/nsstreamsocketsslerrordomain
func (e NSError) NSStreamSocketSSLErrorDomain() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("NSStreamSocketSSLErrorDomain"))
	return NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Foundation/NSError/underlyingErrors
func (e NSError) UnderlyingErrors() []NSError {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("underlyingErrors"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSError {
		return NSErrorFromID(id)
	})
}
// The corresponding value is an object that conforms to the
// NSErrorRecoveryAttempting informal protocol.
//
// See: https://developer.apple.com/documentation/foundation/nsrecoveryattemptererrorkey
func (e NSError) NSRecoveryAttempterErrorKey() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("NSRecoveryAttempterErrorKey"))
	return NSStringFromID(rv).String()
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

// SetUserInfoValueProviderForDomainProviderSync is a synchronous wrapper around [NSError.SetUserInfoValueProviderForDomainProvider].
// It blocks until the completion handler fires or the context is cancelled.
func (ec NSErrorClass) SetUserInfoValueProviderForDomainProviderSync(ctx context.Context, errorDomain NSErrorDomain) error {
	done := make(chan error, 1)
	ec.SetUserInfoValueProviderForDomainProvider(errorDomain, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (n_ NSError) Error() string {
	desc := any(n_.LocalizedDescription())
	if s, ok := desc.(string); ok {
		return s
	}
	if s, ok := desc.(interface{ String() string }); ok {
		return s.String()
	}
	return ""
}

