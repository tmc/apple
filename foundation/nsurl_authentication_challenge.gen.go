// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [URLAuthenticationChallenge] class.
var (
	_URLAuthenticationChallengeClass     URLAuthenticationChallengeClass
	_URLAuthenticationChallengeClassOnce sync.Once
)

func getURLAuthenticationChallengeClass() URLAuthenticationChallengeClass {
	_URLAuthenticationChallengeClassOnce.Do(func() {
		_URLAuthenticationChallengeClass = URLAuthenticationChallengeClass{class: objc.GetClass("NSURLAuthenticationChallenge")}
	})
	return _URLAuthenticationChallengeClass
}

// GetURLAuthenticationChallengeClass returns the class object for NSURLAuthenticationChallenge.
func GetURLAuthenticationChallengeClass() URLAuthenticationChallengeClass {
	return getURLAuthenticationChallengeClass()
}

type URLAuthenticationChallengeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc URLAuthenticationChallengeClass) Alloc() URLAuthenticationChallenge {
	rv := objc.Send[URLAuthenticationChallenge](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A challenge from a server requiring authentication from the client.
//
// # Overview
// 
// Your app receives authentication challenges in various [NSURLSession],
// [NSURLConnection], and [NSURLDownload] delegate methods, such as
// [URLSessionTaskDidReceiveChallengeCompletionHandler]. These objects provide
// the information you’ll need when deciding how to handle a server’s
// request for authentication.
// 
// At the core of that authentication challenge is a that defines the type of
// authentication being requested, the host and port number, the networking
// protocol, and (where applicable) the authentication realm (a group of
// related URLs on the same server that share a single set of credentials).
//
// # Creating an authentication challenge instance
//
//   - [URLAuthenticationChallenge.InitWithAuthenticationChallengeSender]: Creates an authentication challenge from an existing challenge instance.
//   - [URLAuthenticationChallenge.InitWithProtectionSpaceProposedCredentialPreviousFailureCountFailureResponseErrorSender]: Initializes an authentication challenge from parameters you provide.
//
// # Inspecting the authentication challenge
//
//   - [URLAuthenticationChallenge.ProtectionSpace]: The receiver’s protection space.
//
// # Getting properties of previous authentication attempts
//
//   - [URLAuthenticationChallenge.FailureResponse]: The URL response object representing the last authentication failure.
//   - [URLAuthenticationChallenge.PreviousFailureCount]: The receiver’s count of failed authentication attempts.
//   - [URLAuthenticationChallenge.ProposedCredential]: The proposed credential for this challenge.
//
// # Getting authentication errors
//
//   - [URLAuthenticationChallenge.Error]: The error object representing the last authentication failure.
//
// # Legacy
//
//   - [URLAuthenticationChallenge.Sender]: The sender of the challenge.
//
// See: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge
type URLAuthenticationChallenge struct {
	objectivec.Object
}

// URLAuthenticationChallengeFromID constructs a [URLAuthenticationChallenge] from an objc.ID.
//
// A challenge from a server requiring authentication from the client.
func URLAuthenticationChallengeFromID(id objc.ID) URLAuthenticationChallenge {
	return NSURLAuthenticationChallenge{objectivec.Object{ID: id}}
}

// NSURLAuthenticationChallengeFromID is an alias for [URLAuthenticationChallengeFromID] for cross-framework compatibility.
func NSURLAuthenticationChallengeFromID(id objc.ID) URLAuthenticationChallenge { return URLAuthenticationChallengeFromID(id) }
// NOTE: URLAuthenticationChallenge adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [URLAuthenticationChallenge] class.
//
// # Creating an authentication challenge instance
//
//   - [IURLAuthenticationChallenge.InitWithAuthenticationChallengeSender]: Creates an authentication challenge from an existing challenge instance.
//   - [IURLAuthenticationChallenge.InitWithProtectionSpaceProposedCredentialPreviousFailureCountFailureResponseErrorSender]: Initializes an authentication challenge from parameters you provide.
//
// # Inspecting the authentication challenge
//
//   - [IURLAuthenticationChallenge.ProtectionSpace]: The receiver’s protection space.
//
// # Getting properties of previous authentication attempts
//
//   - [IURLAuthenticationChallenge.FailureResponse]: The URL response object representing the last authentication failure.
//   - [IURLAuthenticationChallenge.PreviousFailureCount]: The receiver’s count of failed authentication attempts.
//   - [IURLAuthenticationChallenge.ProposedCredential]: The proposed credential for this challenge.
//
// # Getting authentication errors
//
//   - [IURLAuthenticationChallenge.Error]: The error object representing the last authentication failure.
//
// # Legacy
//
//   - [IURLAuthenticationChallenge.Sender]: The sender of the challenge.
//
// See: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge
type IURLAuthenticationChallenge interface {
	objectivec.IObject
	NSCoding
	NSSecureCoding

	// Topic: Creating an authentication challenge instance

	// Creates an authentication challenge from an existing challenge instance.
	InitWithAuthenticationChallengeSender(challenge INSURLAuthenticationChallenge, sender NSURLAuthenticationChallengeSender) URLAuthenticationChallenge
	// Initializes an authentication challenge from parameters you provide.
	InitWithProtectionSpaceProposedCredentialPreviousFailureCountFailureResponseErrorSender(space INSURLProtectionSpace, credential INSURLCredential, previousFailureCount int, response INSURLResponse, error_ INSError, sender NSURLAuthenticationChallengeSender) URLAuthenticationChallenge

	// Topic: Inspecting the authentication challenge

	// The receiver’s protection space.
	ProtectionSpace() INSURLProtectionSpace

	// Topic: Getting properties of previous authentication attempts

	// The URL response object representing the last authentication failure.
	FailureResponse() INSURLResponse
	// The receiver’s count of failed authentication attempts.
	PreviousFailureCount() int
	// The proposed credential for this challenge.
	ProposedCredential() INSURLCredential

	// Topic: Getting authentication errors

	// The error object representing the last authentication failure.
	Error() INSError

	// Topic: Legacy

	// The sender of the challenge.
	Sender() NSURLAuthenticationChallengeSender
}

// Init initializes the instance.
func (u URLAuthenticationChallenge) Init() URLAuthenticationChallenge {
	rv := objc.Send[URLAuthenticationChallenge](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u URLAuthenticationChallenge) Autorelease() URLAuthenticationChallenge {
	rv := objc.Send[URLAuthenticationChallenge](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLAuthenticationChallenge creates a new URLAuthenticationChallenge instance.
func NewURLAuthenticationChallenge() URLAuthenticationChallenge {
	class := getURLAuthenticationChallengeClass()
	rv := objc.Send[URLAuthenticationChallenge](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an authentication challenge from an existing challenge instance.
//
// challenge: The challenge that you want to copy. Usually, this is a challenge received
// by an existing [NSURLProtocol] subclass that you are subclassing.
//
// sender: The sender that you want to use for the new object. Typically, the sender
// is the instance of your custom [NSURLProtocol] subclass that called this
// method.
//
// # Return Value
// 
// A new authentication challenge object, based on an existing challenge.
//
// # Discussion
// 
// Most apps don’t create [NSURLAuthenticationChallenge] instances
// themselves. Instead, they handle received challenges in the
// [URLSessionTaskDidReceiveChallengeCompletionHandler] method of
// [NSURLSessionTaskDelegate].
// 
// However, you might need to create authentication challenge objects when
// adding support for custom networking protocols, as part of a custom
// [NSURLProtocol] subclass. When you subclass an existing [NSURLProtocol]
// subclass, this initializer lets you modify challenges issued by the
// existing class so that your subclass receives any responses to those
// challenges.
//
// See: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge/init(authenticationChallenge:sender:)
func NewURLAuthenticationChallengeWithAuthenticationChallengeSender(challenge INSURLAuthenticationChallenge, sender NSURLAuthenticationChallengeSender) URLAuthenticationChallenge {
	instance := getURLAuthenticationChallengeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAuthenticationChallenge:sender:"), challenge, sender)
	return URLAuthenticationChallengeFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewURLAuthenticationChallengeWithCoder(coder INSCoder) URLAuthenticationChallenge {
	instance := getURLAuthenticationChallengeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return URLAuthenticationChallengeFromID(rv)
}

// Initializes an authentication challenge from parameters you provide.
//
// space: The protection space for the authentication challenge. This provides
// additional information about the authentication request, such as the host,
// port, authentication realm, and so on.
//
// credential: The proposed credential, or `nil`.
//
// previousFailureCount: The total number of previous failures for this request, including failures
// for other protection spaces.
//
// response: An instance of [NSURLResponse] containing the server response that caused
// you to generate an authentication challenge, or `nil` if no response object
// is applicable to the challenge.
//
// error: An `NS``Error` instance describing the authentication failure, or `nil` if
// it is not applicable to the challenge.
//
// sender: The object that initiated the authentication challenge (typically, the
// object that called this method).
//
// # Return Value
// 
// A new authentication challenge object, with the given properties.
//
// # Discussion
// 
// Most apps don’t create [NSURLAuthenticationChallenge] instances
// themselves. Instead, they handle received challenges in the
// [URLSessionTaskDidReceiveChallengeCompletionHandler] method of
// [NSURLSessionTaskDelegate].
// 
// However, you might need to create authentication challenge objects when
// adding support for custom networking protocols, as part of your custom
// [NSURLProtocol] subclasses.
//
// See: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge/init(protectionSpace:proposedCredential:previousFailureCount:failureResponse:error:sender:)
func NewURLAuthenticationChallengeWithProtectionSpaceProposedCredentialPreviousFailureCountFailureResponseErrorSender(space INSURLProtectionSpace, credential INSURLCredential, previousFailureCount int, response INSURLResponse, error_ INSError, sender NSURLAuthenticationChallengeSender) URLAuthenticationChallenge {
	instance := getURLAuthenticationChallengeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProtectionSpace:proposedCredential:previousFailureCount:failureResponse:error:sender:"), space, credential, previousFailureCount, response, error_, sender)
	return URLAuthenticationChallengeFromID(rv)
}

// Creates an authentication challenge from an existing challenge instance.
//
// challenge: The challenge that you want to copy. Usually, this is a challenge received
// by an existing [NSURLProtocol] subclass that you are subclassing.
//
// sender: The sender that you want to use for the new object. Typically, the sender
// is the instance of your custom [NSURLProtocol] subclass that called this
// method.
//
// # Return Value
// 
// A new authentication challenge object, based on an existing challenge.
//
// # Discussion
// 
// Most apps don’t create [NSURLAuthenticationChallenge] instances
// themselves. Instead, they handle received challenges in the
// [URLSessionTaskDidReceiveChallengeCompletionHandler] method of
// [NSURLSessionTaskDelegate].
// 
// However, you might need to create authentication challenge objects when
// adding support for custom networking protocols, as part of a custom
// [NSURLProtocol] subclass. When you subclass an existing [NSURLProtocol]
// subclass, this initializer lets you modify challenges issued by the
// existing class so that your subclass receives any responses to those
// challenges.
//
// See: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge/init(authenticationChallenge:sender:)
func (u URLAuthenticationChallenge) InitWithAuthenticationChallengeSender(challenge INSURLAuthenticationChallenge, sender NSURLAuthenticationChallengeSender) URLAuthenticationChallenge {
	rv := objc.Send[URLAuthenticationChallenge](u.ID, objc.Sel("initWithAuthenticationChallenge:sender:"), challenge, sender)
	return rv
}
// Initializes an authentication challenge from parameters you provide.
//
// space: The protection space for the authentication challenge. This provides
// additional information about the authentication request, such as the host,
// port, authentication realm, and so on.
//
// credential: The proposed credential, or `nil`.
//
// previousFailureCount: The total number of previous failures for this request, including failures
// for other protection spaces.
//
// response: An instance of [NSURLResponse] containing the server response that caused
// you to generate an authentication challenge, or `nil` if no response object
// is applicable to the challenge.
//
// error: An `NS``Error` instance describing the authentication failure, or `nil` if
// it is not applicable to the challenge.
//
// sender: The object that initiated the authentication challenge (typically, the
// object that called this method).
//
// # Return Value
// 
// A new authentication challenge object, with the given properties.
//
// # Discussion
// 
// Most apps don’t create [NSURLAuthenticationChallenge] instances
// themselves. Instead, they handle received challenges in the
// [URLSessionTaskDidReceiveChallengeCompletionHandler] method of
// [NSURLSessionTaskDelegate].
// 
// However, you might need to create authentication challenge objects when
// adding support for custom networking protocols, as part of your custom
// [NSURLProtocol] subclasses.
//
// See: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge/init(protectionSpace:proposedCredential:previousFailureCount:failureResponse:error:sender:)
func (u URLAuthenticationChallenge) InitWithProtectionSpaceProposedCredentialPreviousFailureCountFailureResponseErrorSender(space INSURLProtectionSpace, credential INSURLCredential, previousFailureCount int, response INSURLResponse, error_ INSError, sender NSURLAuthenticationChallengeSender) URLAuthenticationChallenge {
	rv := objc.Send[URLAuthenticationChallenge](u.ID, objc.Sel("initWithProtectionSpace:proposedCredential:previousFailureCount:failureResponse:error:sender:"), space, credential, previousFailureCount, response, error_, sender)
	return rv
}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (u URLAuthenticationChallenge) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (u URLAuthenticationChallenge) InitWithCoder(coder INSCoder) URLAuthenticationChallenge {
	rv := objc.Send[URLAuthenticationChallenge](u.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// The receiver’s protection space.
//
// # Discussion
// 
// A protection space object provides additional information about the
// authentication request, such as the host, port, authentication realm, and
// so on. The protection space also tells you whether the authentication
// challenge is asking you to provide the user’s credentials or to verify
// the TLS credentials provided by the server.
//
// See: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge/protectionSpace
func (u URLAuthenticationChallenge) ProtectionSpace() INSURLProtectionSpace {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("protectionSpace"))
	return NSURLProtectionSpaceFromID(objc.ID(rv))
}
// The URL response object representing the last authentication failure.
//
// # Discussion
// 
// This value is `nil` if the protocol doesn’t use responses to indicate an
// authentication failure.
//
// See: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge/failureResponse
func (u URLAuthenticationChallenge) FailureResponse() INSURLResponse {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("failureResponse"))
	return NSURLResponseFromID(objc.ID(rv))
}
// The receiver’s count of failed authentication attempts.
//
// # Discussion
// 
// The previous failure count includes failures from protection spaces, not
// just the current one.
//
// See: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge/previousFailureCount
func (u URLAuthenticationChallenge) PreviousFailureCount() int {
	rv := objc.Send[int](u.ID, objc.Sel("previousFailureCount"))
	return rv
}
// The proposed credential for this challenge.
//
// # Discussion
// 
// This method returns `nil` if there is no default credential for this
// challenge.
// 
// If you have previously attempted to authenticate and failed, this method
// returns the most recent failed credential.
// 
// If the proposed credential is not `nil` and returns [true] when you call
// its [HasPassword] method, then the credential is ready to use as-is. If the
// proposed credential’s [HasPassword] method returns [false], then the
// credential provides a default user name, and the client must prompt the
// user for a corresponding password.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge/proposedCredential
func (u URLAuthenticationChallenge) ProposedCredential() INSURLCredential {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("proposedCredential"))
	return NSURLCredentialFromID(objc.ID(rv))
}
// The error object representing the last authentication failure.
//
// # Discussion
// 
// This value is `nil` if the protocol doesn’t use errors to indicate an
// authentication failure.
//
// See: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge/error
func (u URLAuthenticationChallenge) Error() INSError {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("error"))
	return NSErrorFromID(objc.ID(rv))
}
// The sender of the challenge.
//
// # Discussion
// 
// If you are using the [NSURLSession] API, this value is purely
// informational, because you respond to authentication challenges in your
// [NSURLSessionDelegate] or [NSURLSessionTaskDelegate] implementations, by
// passing [URLSession.AuthChallengeDisposition] constants to the provided
// completion handler blocks.
// 
// However, if you are using the legacy [NSURLConnection] or [NSURLDownload]
// API, you use this object directly in your authentication handler delegate
// method. With these APIs, after you finish processing the authentication
// challenge, you respond by calling methods defined in the
// [NSURLAuthenticationChallengeSender] protocol on this sender.
//
// [URLSession.AuthChallengeDisposition]: https://developer.apple.com/documentation/Foundation/URLSession/AuthChallengeDisposition
//
// See: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallenge/sender
func (u URLAuthenticationChallenge) Sender() NSURLAuthenticationChallengeSender {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("sender"))
	return NSURLAuthenticationChallengeSenderObjectFromID(rv)
}

			// Protocol methods for NSSecureCoding
			

