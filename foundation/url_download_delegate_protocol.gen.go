// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// A protocol that URL download delegates implement to interact with a URL download request.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLDownloadDelegate
type NSURLDownloadDelegate interface {
	objectivec.IObject
}

// NSURLDownloadDelegateObject wraps an existing Objective-C object that conforms to the NSURLDownloadDelegate protocol.
type NSURLDownloadDelegateObject struct {
	objectivec.Object
}
func (o NSURLDownloadDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSURLDownloadDelegateObjectFromID constructs a [NSURLDownloadDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSURLDownloadDelegateObjectFromID(id objc.ID) NSURLDownloadDelegateObject {
	return NSURLDownloadDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Sent to determine whether the delegate is able to respond to a protection
// space’s form of authentication.
//
// connection: The download sending the message.
//
// protectionSpace: The protection space that generates an authentication challenge.
//
// # Discussion
// 
// This method is called before [DownloadDidReceiveAuthenticationChallenge],
// allowing the delegate to inspect a protection space before attempting to
// authenticate against it. By returning [true], the delegate indicates that
// it can handle the form of authentication, which it does in the subsequent
// call to [DownloadDidReceiveAuthenticationChallenge]. Not implementing this
// method is the same as returning [false], in which case default
// authentication handling is used.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSURLDownloadDelegate/download(_:canAuthenticateAgainstProtectionSpace:)
func (o NSURLDownloadDelegateObject) DownloadCanAuthenticateAgainstProtectionSpace(connection INSURLDownload, protectionSpace INSURLProtectionSpace) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("download:canAuthenticateAgainstProtectionSpace:"), connection, protectionSpace)
	return rv
	}
// Sent if an authentication challenge is canceled due to the protocol
// implementation encountering an error.
//
// download: The URL download object sending the message.
//
// challenge: The authentication challenge that caused the download object to cancel the
// download.
//
// # Discussion
// 
// If the delegate receives this message the download will fail and the
// delegate will receive a [DownloadDidFailWithError] message.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLDownloadDelegate/download(_:didCancel:)
func (o NSURLDownloadDelegateObject) DownloadDidCancelAuthenticationChallenge(download INSURLDownload, challenge INSURLAuthenticationChallenge) {
	
	objc.Send[struct{}](o.ID, objc.Sel("download:didCancelAuthenticationChallenge:"), download, challenge)
	}
// Sent when the URL download must authenticate a challenge in order to
// download the request.
//
// download: The URL download object sending the message.
//
// challenge: The URL authentication challenge that must be authenticated in order to
// download the request.
//
// # Discussion
// 
// This method gives the delegate the opportunity to determine the course of
// action taken for the challenge: provide credentials, continue without
// providing credentials or cancel the authentication challenge and the
// download.
// 
// The delegate can determine the number of previous authentication challenges
// by sending the message [PreviousFailureCount] to `challenge`.
// 
// If the previous failure count is 0 and the value returned by
// [ProposedCredential] is `nil`, the delegate can create a new
// NSURLCredential object, providing information specific to the type of
// credential, and send a [UseCredentialForAuthenticationChallenge] message to
// `[challenge sender]`, passing the credential and `challenge` as parameters.
// If [ProposedCredential] is not `nil`, the value is a credential from the
// URL or the shared credential storage that can be provided to the user as
// feedback.
// 
// The delegate may decide to abandon further attempts at authentication at
// any time by sending `[challenge sender]` a
// [ContinueWithoutCredentialForAuthenticationChallenge] or a
// [CancelAuthenticationChallenge] message. The specific action is
// implementation dependent.
// 
// If the delegate implements this method, the download will suspend until
// `[challenge sender]` is sent one of the following messages:
// [UseCredentialForAuthenticationChallenge],
// [ContinueWithoutCredentialForAuthenticationChallenge] or
// [CancelAuthenticationChallenge].
// 
// If the delegate does not implement this method the default implementation
// is used. If a valid credential for the request is provided as part of the
// URL, or is available from the NSURLCredentialStorage the `[challenge
// sender]` is sent a [UseCredentialForAuthenticationChallenge] with the
// credential. If the challenge has no credential or the credentials fail to
// authorize access, then
// [ContinueWithoutCredentialForAuthenticationChallenge] is sent to
// `[challenge sender]` instead.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLDownloadDelegate/download(_:didReceive:)-1pc0v
func (o NSURLDownloadDelegateObject) DownloadDidReceiveAuthenticationChallenge(download INSURLDownload, challenge INSURLAuthenticationChallenge) {
	
	objc.Send[struct{}](o.ID, objc.Sel("download:didReceiveAuthenticationChallenge:"), download, challenge)
	}
// Sent to determine whether the URL loader should consult the credential
// storage to authenticate the download.
//
// download: The connection sending the message.
//
// # Discussion
// 
// This method is called before any attempt to authenticate is made. By
// returning [false], the delegate tells the download not to consult the
// credential storage and makes itself responsible for providing credentials
// for any authentication challenges. Not implementing this method is the same
// as returing [true]. The delegate is free to consult the credential storage
// itself when it receives a [DownloadDidReceiveAuthenticationChallenge]
// message.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSURLDownloadDelegate/downloadShouldUseCredentialStorage(_:)
func (o NSURLDownloadDelegateObject) DownloadShouldUseCredentialStorage(download INSURLDownload) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("downloadShouldUseCredentialStorage:"), download)
	return rv
	}
// The delegate receives this message when `download` has determined a
// suggested filename for the downloaded file.
//
// download: The URL download object sending the message.
//
// filename: The suggested filename for the download.
//
// # Discussion
// 
// The suggested filename is either derived from the last path component of
// the URL and the MIME type or, if the download was encoded, from the
// encoding. If the delegate wishes to modify the path, it should send
// [SetDestinationAllowOverwrite] to `download`.
// 
// # Special Considerations
// 
// The delegate will not receive this message if `` has already been called
// for the download.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLDownloadDelegate/download(_:decideDestinationWithSuggestedFilename:)
func (o NSURLDownloadDelegateObject) DownloadDecideDestinationWithSuggestedFilename(download INSURLDownload, filename string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("download:decideDestinationWithSuggestedFilename:"), download, objc.String(filename))
	}
// Sent immediately after a download object begins a download.
//
// download: The URL download object sending the message.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLDownloadDelegate/downloadDidBegin(_:)
func (o NSURLDownloadDelegateObject) DownloadDidBegin(download INSURLDownload) {
	
	objc.Send[struct{}](o.ID, objc.Sel("downloadDidBegin:"), download)
	}
// Sent when the destination file is created.
//
// download: The URL download object sending the message.
//
// path: The path to the destination file.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLDownloadDelegate/download(_:didCreateDestination:)
func (o NSURLDownloadDelegateObject) DownloadDidCreateDestination(download INSURLDownload, path string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("download:didCreateDestination:"), download, objc.String(path))
	}
// Sent when a download object has received sufficient load data to construct
// the NSURLResponse object for the download.
//
// download: The URL download object sending the message.
//
// response: The URL response object received as part of the download. `response` is
// immutable and will not be modified after this method is called.
//
// # Discussion
// 
// In some rare cases, multiple responses may be received for a single
// download. In this case, the client should assume that each new response
// resets the download progress to 0 and should check the new response for the
// expected content length.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLDownloadDelegate/download(_:didReceive:)-817z3
func (o NSURLDownloadDelegateObject) DownloadDidReceiveResponse(download INSURLDownload, response INSURLResponse) {
	
	objc.Send[struct{}](o.ID, objc.Sel("download:didReceiveResponse:"), download, response)
	}
// Sent as a download object receives data incrementally.
//
// download: The URL download object sending the message.
//
// length: The amount of data received in this increment of the download, measured in
// bytes.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLDownloadDelegate/download(_:didReceiveDataOfLength:)
func (o NSURLDownloadDelegateObject) DownloadDidReceiveDataOfLength(download INSURLDownload, length uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("download:didReceiveDataOfLength:"), download, length)
	}
// Sent when a download object determines that the downloaded file is encoded
// to inquire whether the file should be automatically decoded.
//
// download: The URL download object sending the message.
//
// encodingType: The type of encoding used by the downloaded file. The supported encoding
// formats are MacBinary (`"application/macbinary"`), Binhex
// (`"application/mac-binhex40"`) and gzip (`"application/gzip"`).
//
// # Return Value
// 
// [true] to decode the file, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The delegate may receive this message more than once if the file has been
// encoded multiple times. This method is not called if the downloaded file is
// not encoded.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLDownloadDelegate/download(_:shouldDecodeSourceDataOfMIMEType:)
func (o NSURLDownloadDelegateObject) DownloadShouldDecodeSourceDataOfMIMEType(download INSURLDownload, encodingType string) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("download:shouldDecodeSourceDataOfMIMEType:"), download, objc.String(encodingType))
	return rv
	}
// Sent when a download object has received a response from the server after
// attempting to resume a download.
//
// download: The URL download object sending the message.
//
// response: The URL response received from the server in response to an attempt to
// resume a download.
//
// startingByte: The location of the start of the resumed data, in bytes.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLDownloadDelegate/download(_:willResumeWith:fromByte:)
func (o NSURLDownloadDelegateObject) DownloadWillResumeWithResponseFromByte(download INSURLDownload, response INSURLResponse, startingByte int64) {
	
	objc.Send[struct{}](o.ID, objc.Sel("download:willResumeWithResponse:fromByte:"), download, response, startingByte)
	}
// Sent when the download object determines that it must change URLs in order
// to continue loading a request.
//
// download: The URL download object sending the message.
//
// request: The proposed redirected request. The delegate should inspect the redirected
// request to verify that it meets its needs, and create a copy with new
// attributes to return to the connection if necessary.
//
// redirectResponse: The URL response that caused the redirect. May be `nil` in cases where this
// method is not being sent as a result of involving the delegate in redirect
// processing.
//
// # Return Value
// 
// The actual URL request to use in light of the redirection response. The
// delegate may copy and modify `request` as necessary to change its
// attributes, return `request` unmodified, or return `nil`.
//
// # Discussion
// 
// If the delegate wishes to cancel the redirect, it should call the
// `download` object’s [Cancel] method. Alternatively, the delegate method
// can return `nil` to cancel the redirect, and the download will continue to
// process. This has special relevance in the case where `redirectResponse` is
// not `nil`. In this case, any data that is loaded for the download will be
// sent to the delegate, and the delegate will receive a [DownloadDidFinish]
// or [DownloadDidFailWithError] message, as appropriate.
// 
// # Special Considerations
// 
// The delegate can receive this message as a result of transforming a
// request’s URL to its canonical form, or for protocol-specific reasons,
// such as an HTTP redirect. The delegate implementation should be prepared to
// receive this message multiple times.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLDownloadDelegate/download(_:willSend:redirectResponse:)
func (o NSURLDownloadDelegateObject) DownloadWillSendRequestRedirectResponse(download INSURLDownload, request INSURLRequest, redirectResponse INSURLResponse) INSURLRequest {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("download:willSendRequest:redirectResponse:"), download, request, redirectResponse)
	return NSURLRequestFromID(rv)
	}
// Sent if the download fails or if an I/O error occurs when the file is
// written to disk.
//
// download: The URL download object sending the message.
//
// error: The error that caused the failure of the download.
//
// # Discussion
// 
// Any partially downloaded file will be deleted.
// 
// # Special Considerations
// 
// Once the delegate receives this message, it will receive no further
// messages for `download`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLDownloadDelegate/download(_:didFailWithError:)
func (o NSURLDownloadDelegateObject) DownloadDidFailWithError(download INSURLDownload, error_ INSError) {
	
	objc.Send[struct{}](o.ID, objc.Sel("download:didFailWithError:"), download, error_)
	}
// Sent when a download object has completed downloading successfully and has
// written its results to disk.
//
// download: The URL download object sending the message.
//
// # Discussion
// 
// The delegate will receive no further messages for `download`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLDownloadDelegate/downloadDidFinish(_:)
func (o NSURLDownloadDelegateObject) DownloadDidFinish(download INSURLDownload) {
	
	objc.Send[struct{}](o.ID, objc.Sel("downloadDidFinish:"), download)
	}

// NSURLDownloadDelegateConfig holds optional typed callbacks for [NSURLDownloadDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsurldownloaddelegate
type NSURLDownloadDelegateConfig struct {

	// Download Authentication
	// DownloadCanAuthenticateAgainstProtectionSpace — Sent to determine whether the delegate is able to respond to a protection space’s form of authentication.
	DownloadCanAuthenticateAgainstProtectionSpace func(connection NSURLDownload, protectionSpace NSURLProtectionSpace) bool
	// DownloadShouldUseCredentialStorage — Sent to determine whether the URL loader should consult the credential storage to authenticate the download.
	DownloadShouldUseCredentialStorage func(download NSURLDownload) bool

	// Download Data and Responses
	// DownloadDidBegin — Sent immediately after a download object begins a download.
	DownloadDidBegin func(download NSURLDownload)
	// DownloadDidReceiveDataOfLength — Sent as a download object receives data incrementally.
	DownloadDidReceiveDataOfLength func(download NSURLDownload, length uint)

	// Download Completion
	// DownloadDidFailWithError — Sent if the download fails or if an I/O error occurs when the file is written to disk.
	DownloadDidFailWithError func(download NSURLDownload, error_ objectivec.Object)
	// DownloadDidFinish — Sent when a download object has completed downloading successfully and has written its results to disk.
	DownloadDidFinish func(download NSURLDownload)

	// Other Methods
	// DownloadDidCancelAuthenticationChallenge — Sent if an authentication challenge is canceled due to the protocol implementation encountering an error.
	DownloadDidCancelAuthenticationChallenge func(download NSURLDownload, challenge NSURLAuthenticationChallenge)
	// DownloadDidReceiveAuthenticationChallenge — Sent when the URL download must authenticate a challenge in order to download the request.
	DownloadDidReceiveAuthenticationChallenge func(download NSURLDownload, challenge NSURLAuthenticationChallenge)
	// DownloadDidReceiveResponse — Sent when a download object has received sufficient load data to construct the NSURLResponse object for the download.
	DownloadDidReceiveResponse func(download NSURLDownload, response NSURLResponse)
	// DownloadWillResumeWithResponseFromByte — Sent when a download object has received a response from the server after attempting to resume a download.
	DownloadWillResumeWithResponseFromByte func(download NSURLDownload, response NSURLResponse, startingByte int64)
	// DownloadWillSendRequestRedirectResponse — Sent when the download object determines that it must change URLs in order to continue loading a request.
	DownloadWillSendRequestRedirectResponse func(download NSURLDownload, request NSURLRequest, redirectResponse NSURLResponse) NSURLRequest
}

// NewNSURLDownloadDelegate creates an Objective-C object implementing the [NSURLDownloadDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSURLDownloadDelegateObject] satisfies the [NSURLDownloadDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsurldownloaddelegate
func NewNSURLDownloadDelegate(config NSURLDownloadDelegateConfig) NSURLDownloadDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSURLDownloadDelegate_%d", n)

	var methods []objc.MethodDef

	if config.DownloadCanAuthenticateAgainstProtectionSpace != nil {
		fn := config.DownloadCanAuthenticateAgainstProtectionSpace
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("download:canAuthenticateAgainstProtectionSpace:"),
			Fn: func(self objc.ID, _cmd objc.SEL, connectionID objc.ID, protectionSpaceID objc.ID) bool {
				connection := NSURLDownloadFromID(connectionID)
				protectionSpace := NSURLProtectionSpaceFromID(protectionSpaceID)
				return fn(connection, protectionSpace)
			},
		})
	}

	if config.DownloadDidCancelAuthenticationChallenge != nil {
		fn := config.DownloadDidCancelAuthenticationChallenge
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("download:didCancelAuthenticationChallenge:"),
			Fn: func(self objc.ID, _cmd objc.SEL, downloadID objc.ID, challengeID objc.ID) {
				download := NSURLDownloadFromID(downloadID)
				challenge := NSURLAuthenticationChallengeFromID(challengeID)
				fn(download, challenge)
			},
		})
	}

	if config.DownloadDidReceiveAuthenticationChallenge != nil {
		fn := config.DownloadDidReceiveAuthenticationChallenge
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("download:didReceiveAuthenticationChallenge:"),
			Fn: func(self objc.ID, _cmd objc.SEL, downloadID objc.ID, challengeID objc.ID) {
				download := NSURLDownloadFromID(downloadID)
				challenge := NSURLAuthenticationChallengeFromID(challengeID)
				fn(download, challenge)
			},
		})
	}

	if config.DownloadShouldUseCredentialStorage != nil {
		fn := config.DownloadShouldUseCredentialStorage
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("downloadShouldUseCredentialStorage:"),
			Fn: func(self objc.ID, _cmd objc.SEL, downloadID objc.ID) bool {
				download := NSURLDownloadFromID(downloadID)
				return fn(download)
			},
		})
	}

	if config.DownloadDidBegin != nil {
		fn := config.DownloadDidBegin
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("downloadDidBegin:"),
			Fn: func(self objc.ID, _cmd objc.SEL, downloadID objc.ID) {
				download := NSURLDownloadFromID(downloadID)
				fn(download)
			},
		})
	}

	if config.DownloadDidReceiveResponse != nil {
		fn := config.DownloadDidReceiveResponse
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("download:didReceiveResponse:"),
			Fn: func(self objc.ID, _cmd objc.SEL, downloadID objc.ID, responseID objc.ID) {
				download := NSURLDownloadFromID(downloadID)
				response := NSURLResponseFromID(responseID)
				fn(download, response)
			},
		})
	}

	if config.DownloadDidReceiveDataOfLength != nil {
		fn := config.DownloadDidReceiveDataOfLength
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("download:didReceiveDataOfLength:"),
			Fn: func(self objc.ID, _cmd objc.SEL, downloadID objc.ID, length uint) {
				download := NSURLDownloadFromID(downloadID)
				fn(download, length)
			},
		})
	}

	if config.DownloadWillResumeWithResponseFromByte != nil {
		fn := config.DownloadWillResumeWithResponseFromByte
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("download:willResumeWithResponse:fromByte:"),
			Fn: func(self objc.ID, _cmd objc.SEL, downloadID objc.ID, responseID objc.ID, startingByte int64) {
				download := NSURLDownloadFromID(downloadID)
				response := NSURLResponseFromID(responseID)
				fn(download, response, startingByte)
			},
		})
	}

	if config.DownloadWillSendRequestRedirectResponse != nil {
		fn := config.DownloadWillSendRequestRedirectResponse
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("download:willSendRequest:redirectResponse:"),
			Fn: func(self objc.ID, _cmd objc.SEL, downloadID objc.ID, requestID objc.ID, redirectResponseID objc.ID) objc.ID {
				download := NSURLDownloadFromID(downloadID)
				request := NSURLRequestFromID(requestID)
				redirectResponse := NSURLResponseFromID(redirectResponseID)
				return fn(download, request, redirectResponse).GetID()
			},
		})
	}

	if config.DownloadDidFailWithError != nil {
		fn := config.DownloadDidFailWithError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("download:didFailWithError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, downloadID objc.ID, error_ID objc.ID) {
				download := NSURLDownloadFromID(downloadID)
				error_ := objectivec.ObjectFromID(error_ID)
				fn(download, error_)
			},
		})
	}

	if config.DownloadDidFinish != nil {
		fn := config.DownloadDidFinish
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("downloadDidFinish:"),
			Fn: func(self objc.ID, _cmd objc.SEL, downloadID objc.ID) {
				download := NSURLDownloadFromID(downloadID)
				fn(download)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSURLDownloadDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSURLDownloadDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSURLDownloadDelegateObjectFromID(instance)
}

