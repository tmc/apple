// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A protocol you implement to track download progress and handle redirects, authentication challenges, and failures.
//
// See: https://developer.apple.com/documentation/WebKit/WKDownloadDelegate
type WKDownloadDelegate interface {
	objectivec.IObject

	// Asks the delegate to provide a file destination where the system should write the download data.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKDownloadDelegate/download(_:decideDestinationUsing:suggestedFilename:completionHandler:)
	DownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler(download IWKDownload, response foundation.NSURLResponse, suggestedFilename string, completionHandler URLHandler)
}

// WKDownloadDelegateObject wraps an existing Objective-C object that conforms to the WKDownloadDelegate protocol.
type WKDownloadDelegateObject struct {
	objectivec.Object
}

func (o WKDownloadDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// WKDownloadDelegateObjectFromID constructs a [WKDownloadDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func WKDownloadDelegateObjectFromID(id objc.ID) WKDownloadDelegateObject {
	return WKDownloadDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Asks the delegate to provide a file destination where the system should
// write the download data.
//
// download: The download that needs a file destination where the systems should write
// the download data.
//
// response: A response from the server for an HTTP request, or a synthesized response
// for a blob download.
//
// suggestedFilename: A string with a filename suggestion to use in creating the file
// destination.
//
// completionHandler: A closure you invoke with a destination file URL to begin the download, or
// `nil` to cancel the download.
//
// # Discussion
//
// The suggested filename can come from the response or from the web content.
//
// The destination file URL must meet the following requirements:
//
// - It’s a file that doesn’t exist. - It’s in a directory that exists.
// - It’s in a directory that WebKit can write to.
//
// See: https://developer.apple.com/documentation/WebKit/WKDownloadDelegate/download(_:decideDestinationUsing:suggestedFilename:completionHandler:)
func (o WKDownloadDelegateObject) DownloadDecideDestinationUsingResponseSuggestedFilenameCompletionHandler(download IWKDownload, response foundation.NSURLResponse, suggestedFilename string, completionHandler URLHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("download:decideDestinationUsingResponse:suggestedFilename:completionHandler:"), download, response, objc.String(suggestedFilename), completionHandler)
}

// Tells the delegate that the download finished.
//
// download: The download that finished.
//
// See: https://developer.apple.com/documentation/WebKit/WKDownloadDelegate/downloadDidFinish(_:)
func (o WKDownloadDelegateObject) DownloadDidFinish(download IWKDownload) {
	objc.Send[struct{}](o.ID, objc.Sel("downloadDidFinish:"), download)
}

// Tells the delegate that the download failed, with error information and
// data you can use to restart the download.
//
// download: The download that failed.
//
// error: An error describing what caused the download to fail.
//
// resumeData: A data object you use to restart the download.
//
// # Discussion
//
// To restart a failed download, call
// [ResumeDownloadFromResumeDataCompletionHandler] with `resumeData`.
//
// See: https://developer.apple.com/documentation/WebKit/WKDownloadDelegate/download(_:didFailWithError:resumeData:)
func (o WKDownloadDelegateObject) DownloadDidFailWithErrorResumeData(download IWKDownload, error_ foundation.INSError, resumeData foundation.INSData) {
	objc.Send[struct{}](o.ID, objc.Sel("download:didFailWithError:resumeData:"), download, error_, resumeData)
}

// Asks the delegate to respond to an authentication challenge.
//
// download: The download that received the authentication challenge.
//
// challenge: The authentication challenge.
//
// completionHandler: A closure you must invoke to respond to the authentication challenge.
// Provide the closure with a disposition that describes how to respond to the
// authorization challenge, and optional credentials.
//
// # Discussion
//
// Determine how to respond to the authentication challenge in this method.
// Then invoke `completionHandler` with a disposition that describes how to
// respond to the authorization challenge, and optional credentials.
//
// If you don’t implement this method, the web view responds to the
// challenge with [URLSession.AuthChallengeDisposition.rejectProtectionSpace].
//
// See: https://developer.apple.com/documentation/WebKit/WKDownloadDelegate/download(_:didReceive:completionHandler:)
//
// [URLSession.AuthChallengeDisposition.rejectProtectionSpace]: https://developer.apple.com/documentation/Foundation/URLSession/AuthChallengeDisposition/rejectProtectionSpace
func (o WKDownloadDelegateObject) DownloadDidReceiveAuthenticationChallengeCompletionHandler(download IWKDownload, challenge foundation.NSURLAuthenticationChallenge, completionHandler URLCredentialHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("download:didReceiveAuthenticationChallenge:completionHandler:"), download, challenge, completionHandler)
}

// Asks the delegate to respond to the download’s redirect response.
//
// download: The download that receives the redirect response.
//
// response: The redirect response.
//
// request: The new request the web view sends as a result of the redirect response.
//
// decisionHandler: A closure you must invoke, providing a download redirect policy that
// indicates whether to proceed with the redirect.
//
// # Discussion
//
// Determine whether to proceed with the redirect. Then invoke the
// decisionHandler closure, providing a download redirect policy that
// indicates whether to proceed with the redirect.
//
// If you don’t implement this method, the web view proceeds with all
// redirects.
//
// See: https://developer.apple.com/documentation/WebKit/WKDownloadDelegate/download(_:willPerformHTTPRedirection:newRequest:decisionHandler:)
func (o WKDownloadDelegateObject) DownloadWillPerformHTTPRedirectionNewRequestDecisionHandler(download IWKDownload, response foundation.NSHTTPURLResponse, request foundation.NSURLRequest, decisionHandler WKDownloadRedirectPolicyHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("download:willPerformHTTPRedirection:newRequest:decisionHandler:"), download, response, request, decisionHandler)
}

// See: https://developer.apple.com/documentation/WebKit/WKDownloadDelegate/download(_:decidePlaceholderPolicy:)
func (o WKDownloadDelegateObject) DownloadDecidePlaceholderPolicy(download IWKDownload, completionHandler URLHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("download:decidePlaceholderPolicy:"), download, completionHandler)
}

// See: https://developer.apple.com/documentation/WebKit/WKDownloadDelegate/download(_:didReceiveFinalURL:)
func (o WKDownloadDelegateObject) DownloadDidReceiveFinalURL(download IWKDownload, url foundation.INSURL) {
	objc.Send[struct{}](o.ID, objc.Sel("download:didReceiveFinalURL:"), download, url)
}

// See: https://developer.apple.com/documentation/WebKit/WKDownloadDelegate/download(_:didReceivePlaceholderURL:completionHandler:)
func (o WKDownloadDelegateObject) DownloadDidReceivePlaceholderURLCompletionHandler(download IWKDownload, url foundation.INSURL, completionHandler VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("download:didReceivePlaceholderURL:completionHandler:"), download, url, completionHandler)
}

// WKDownloadDelegateConfig holds optional typed callbacks for [WKDownloadDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate
type WKDownloadDelegateConfig struct {

	// Tracking Download Progress
	// DownloadDidFinish — Tells the delegate that the download finished.
	DownloadDidFinish func(download WKDownload)
	// DownloadDidFailWithErrorResumeData — Tells the delegate that the download failed, with error information and data you can use to restart the download.
	DownloadDidFailWithErrorResumeData func(download WKDownload, error_ foundation.NSError, resumeData foundation.NSData)

	// Instance Methods
	DownloadDidReceiveFinalURL func(download WKDownload, url foundation.NSURL)
}

// NewWKDownloadDelegate creates an Objective-C object implementing the [WKDownloadDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [WKDownloadDelegateObject] satisfies the [WKDownloadDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/webkit/wkdownloaddelegate
func NewWKDownloadDelegate(config WKDownloadDelegateConfig) WKDownloadDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoWKDownloadDelegate_%d", n)

	var methods []objc.MethodDef

	if config.DownloadDidFinish != nil {
		fn := config.DownloadDidFinish
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("downloadDidFinish:"),
			Fn: func(self objc.ID, _cmd objc.SEL, downloadID objc.ID) {
				download := WKDownloadFromID(downloadID)
				fn(download)
			},
		})
	}

	if config.DownloadDidFailWithErrorResumeData != nil {
		fn := config.DownloadDidFailWithErrorResumeData
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("download:didFailWithError:resumeData:"),
			Fn: func(self objc.ID, _cmd objc.SEL, downloadID objc.ID, error_ID objc.ID, resumeDataID objc.ID) {
				download := WKDownloadFromID(downloadID)
				error_ := foundation.NSErrorFromID(error_ID)
				resumeData := foundation.NSDataFromID(resumeDataID)
				fn(download, error_, resumeData)
			},
		})
	}

	if config.DownloadDidReceiveFinalURL != nil {
		fn := config.DownloadDidReceiveFinalURL
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("download:didReceiveFinalURL:"),
			Fn: func(self objc.ID, _cmd objc.SEL, downloadID objc.ID, urlID objc.ID) {
				download := WKDownloadFromID(downloadID)
				url := foundation.NSURLFromID(urlID)
				fn(download, url)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("WKDownloadDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewWKDownloadDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return WKDownloadDelegateObjectFromID(instance)
}
