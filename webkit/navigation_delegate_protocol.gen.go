// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// Methods for accepting or rejecting navigation changes, and for tracking the progress of navigation requests.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationDelegate
type WKNavigationDelegate interface {
	objectivec.IObject
}

// WKNavigationDelegateObject wraps an existing Objective-C object that conforms to the WKNavigationDelegate protocol.
type WKNavigationDelegateObject struct {
	objectivec.Object
}

func (o WKNavigationDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// WKNavigationDelegateObjectFromID constructs a [WKNavigationDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func WKNavigationDelegateObjectFromID(id objc.ID) WKNavigationDelegateObject {
	return WKNavigationDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Asks the delegate for permission to navigate to new content based on the
// specified preferences and action information.
//
// webView: The web view from which the navigation request began.
//
// navigationAction: Details about the action that triggered the navigation request.
//
// preferences: The default preferences to use when displaying the new webpage. Specify the
// default preferences for pages using the [DefaultWebpagePreferences]
// property of [WKWebViewConfiguration] when you create your web view.
//
// decisionHandler: A completion handler block to call with the results about whether to allow
// or cancel the navigation. This handler has no return value and takes the
// following parameters:
//
// policy: A constant that indicates whether to cancel or allow the
// navigation. For a list of possible values, see [WKNavigationActionPolicy].
// preferences: The set of preferences to apply to the page if the navigation
// is allowed. You may pass the object from the `preferences` parameter or
// configure a new preferences object and pass it instead.
//
// # Discussion
//
// Use this method to allow or deny a navigation request that originated with
// the specified action. The web view calls this method after the interaction
// occurs but before it attempts to load any content. If you implement this
// method, always execute the `decisionHandler` block at some point. You may
// execute it synchronously from your delegate method’s implementation, or
// execute it asynchronously after your method returns.
//
// If your delegate object implements this method, the web view doesn’t call
// the [WebViewDecidePolicyForNavigationActionDecisionHandler] method.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationDelegate/webView(_:decidePolicyFor:preferences:decisionHandler:)
//
// [WKNavigationActionPolicy]: https://developer.apple.com/documentation/WebKit/WKNavigationActionPolicy
func (o WKNavigationDelegateObject) WebViewDecidePolicyForNavigationActionPreferencesDecisionHandler(webView IWKWebView, navigationAction IWKNavigationAction, preferences IWKWebpagePreferences, decisionHandler WKWebpagePreferencesHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:decidePolicyForNavigationAction:preferences:decisionHandler:"), webView, navigationAction, preferences, decisionHandler)
}

// Asks the delegate for permission to navigate to new content based on the
// specified action information.
//
// webView: The web view from which the navigation request began.
//
// navigationAction: Details about the action that triggered the navigation request.
//
// decisionHandler: A completion handler block to call with the results about whether to allow
// or cancel the navigation. This handler has no return value and takes the
// following parameter:
//
// policy: A constant that indicates whether to cancel or allow the
// navigation. For a list of possible values, see [WKNavigationActionPolicy].
//
// # Discussion
//
// Use this method to allow or deny a navigation request that originated with
// the specified action. The web view calls this method after the interaction
// occurs but before it attempts to load any content. If you implement this
// method, always execute the `decisionHandler` block at some point. You may
// execute it synchronously from your delegate method’s implementation, or
// execute it asynchronously after your method returns.
//
// If your delegate object implements the
// [WebViewDecidePolicyForNavigationActionPreferencesDecisionHandler] method,
// the web view doesn’t call this method.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationDelegate/webView(_:decidePolicyFor:decisionHandler:)-2ni62
//
// [WKNavigationActionPolicy]: https://developer.apple.com/documentation/WebKit/WKNavigationActionPolicy
func (o WKNavigationDelegateObject) WebViewDecidePolicyForNavigationActionDecisionHandler(webView IWKWebView, navigationAction IWKNavigationAction, decisionHandler WKNavigationActionPolicyHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:decidePolicyForNavigationAction:decisionHandler:"), webView, navigationAction, decisionHandler)
}

// Asks the delegate for permission to navigate to new content after the
// response to the navigation request is known.
//
// webView: The web view from which the navigation request began.
//
// navigationResponse: Descriptive information about the navigation response.
//
// decisionHandler: A completion handler block to call with the results about whether to allow
// or cancel the navigation. This handler has no return value and takes the
// following parameter:
//
// policy: A constant that indicates whether to cancel or allow the
// navigation. For a list of possible values, see
// [WKNavigationResponsePolicy].
//
// # Discussion
//
// Use this method to allow or deny a navigation request after the web view
// receives the response to its original URL request. The `navigationResponse`
// parameter contains the details of the response, including the type of data
// that the response contains. If you implement this method, always execute
// the `decisionHandler` block at some point. You may execute it synchronously
// from your delegate method’s implementation, or execute it asynchronously
// after your method returns.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationDelegate/webView(_:decidePolicyFor:decisionHandler:)-19mn2
//
// [WKNavigationResponsePolicy]: https://developer.apple.com/documentation/WebKit/WKNavigationResponsePolicy
func (o WKNavigationDelegateObject) WebViewDecidePolicyForNavigationResponseDecisionHandler(webView IWKWebView, navigationResponse IWKNavigationResponse, decisionHandler WKNavigationResponsePolicyHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:decidePolicyForNavigationResponse:decisionHandler:"), webView, navigationResponse, decisionHandler)
}

// Tells the delegate that navigation from the main frame has started.
//
// webView: The web view that is loading the content.
//
// navigation: The navigation object associated with the load request.
//
// # Discussion
//
// The web view calls this method after it receives provisional approval to
// process a navigation request, but before it receives a response to that
// request.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationDelegate/webView(_:didStartProvisionalNavigation:)
func (o WKNavigationDelegateObject) WebViewDidStartProvisionalNavigation(webView IWKWebView, navigation IWKNavigation) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:didStartProvisionalNavigation:"), webView, navigation)
}

// Tells the delegate that the web view received a server redirect for a
// request.
//
// webView: The web view that is loading the content.
//
// navigation: The navigation object that received a server redirect.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationDelegate/webView(_:didReceiveServerRedirectForProvisionalNavigation:)
func (o WKNavigationDelegateObject) WebViewDidReceiveServerRedirectForProvisionalNavigation(webView IWKWebView, navigation IWKNavigation) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:didReceiveServerRedirectForProvisionalNavigation:"), webView, navigation)
}

// Tells the delegate that the web view has started to receive content for the
// main frame.
//
// webView: The web view that is loading the content.
//
// navigation: The navigation object that uniquely identifies the load request.
//
// # Discussion
//
// After the navigation delegate’s
// [WebViewDecidePolicyForNavigationResponseDecisionHandler] method approves
// the navigation response, the web view begins processing it. As changes
// become ready, the web view calls this method immediately before it starts
// to update the main frame.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationDelegate/webView(_:didCommit:)
func (o WKNavigationDelegateObject) WebViewDidCommitNavigation(webView IWKWebView, navigation IWKNavigation) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:didCommitNavigation:"), webView, navigation)
}

// Tells the delegate that navigation is complete.
//
// webView: The web view that loaded the content.
//
// navigation: The navigation object that finished.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationDelegate/webView(_:didFinish:)
func (o WKNavigationDelegateObject) WebViewDidFinishNavigation(webView IWKWebView, navigation IWKNavigation) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:didFinishNavigation:"), webView, navigation)
}

// Asks the delegate to respond to an authentication challenge.
//
// webView: The web view that receives the authentication challenge.
//
// challenge: The authentication challenge.
//
// completionHandler: A completion handler block to execute with the response. This handler has
// no return value and takes the following parameters:
//
// disposition: The option to use to handle the challenge. For a list of
// options, see [URLSession.AuthChallengeDisposition]. credential: The
// credential to use for authentication when the `disposition` parameter
// contains the value [URLSession.AuthChallengeDisposition.useCredential].
// Specify `nil` to continue without a credential.
//
// # Discussion
//
// If you don’t implement this method, the web view responds to the
// authentication challenge with the
// [URLSession.AuthChallengeDisposition.rejectProtectionSpace] disposition.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationDelegate/webView(_:didReceive:completionHandler:)
//
// [URLSession.AuthChallengeDisposition.useCredential]: https://developer.apple.com/documentation/Foundation/URLSession/AuthChallengeDisposition/useCredential
// [URLSession.AuthChallengeDisposition]: https://developer.apple.com/documentation/Foundation/URLSession/AuthChallengeDisposition
// [URLSession.AuthChallengeDisposition.rejectProtectionSpace]: https://developer.apple.com/documentation/Foundation/URLSession/AuthChallengeDisposition/rejectProtectionSpace
func (o WKNavigationDelegateObject) WebViewDidReceiveAuthenticationChallengeCompletionHandler(webView IWKWebView, challenge foundation.NSURLAuthenticationChallenge, completionHandler URLCredentialHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:didReceiveAuthenticationChallenge:completionHandler:"), webView, challenge, completionHandler)
}

// Asks the delegate whether to continue with a connection that uses a
// deprecated version of TLS.
//
// webView: The web view that receives the authentication challenge.
//
// challenge: The authentication challenge.
//
// decisionHandler: The completion handler block to execute with the decision. This handler has
// no return value and takes the following parameter:
//
// decision: A Boolean value that indicates whether to continue to use a
// deprecated version of TLS. Specify true to continue, or false to reject the
// connection.
//
// # Discussion
//
// If you don’t implement this method, the web view uses system settings to
// determine whether to allow the use of deprecated versions of TLS.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationDelegate/webView(_:authenticationChallenge:shouldAllowDeprecatedTLS:)
func (o WKNavigationDelegateObject) WebViewAuthenticationChallengeShouldAllowDeprecatedTLS(webView IWKWebView, challenge foundation.NSURLAuthenticationChallenge, decisionHandler BoolHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:authenticationChallenge:shouldAllowDeprecatedTLS:"), webView, challenge, decisionHandler)
}

// Tells the delegate that an error occurred during navigation.
//
// webView: The web view that reported the error.
//
// navigation: The navigation object for the operation. This object corresponds to a
// [WKNavigation] object that WebKit returned when the load operation began.
// You use it to track the progress of that operation.
//
// error: The error that occurred.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationDelegate/webView(_:didFail:withError:)
func (o WKNavigationDelegateObject) WebViewDidFailNavigationWithError(webView IWKWebView, navigation IWKNavigation, error_ foundation.INSError) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:didFailNavigation:withError:"), webView, navigation, error_)
}

// Tells the delegate that an error occurred during the early navigation
// process.
//
// webView: The web view that called the delegate method.
//
// navigation: The navigation object for the operation. This object corresponds to a
// [WKNavigation] object that WebKit returned when the load operation began.
// You use it to track the progress of that operation.
//
// error: The error that occurred.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationDelegate/webView(_:didFailProvisionalNavigation:withError:)
func (o WKNavigationDelegateObject) WebViewDidFailProvisionalNavigationWithError(webView IWKWebView, navigation IWKNavigation, error_ foundation.INSError) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:didFailProvisionalNavigation:withError:"), webView, navigation, error_)
}

// Tells the delegate that the web view’s content process was terminated.
//
// webView: The web view whose underlying web content process was terminated.
//
// # Discussion
//
// Web views use a separate process to render and manage web content. WebKit
// calls this method when the process for the specified web view terminates
// for any reason.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationDelegate/webViewWebContentProcessDidTerminate(_:)
func (o WKNavigationDelegateObject) WebViewWebContentProcessDidTerminate(webView IWKWebView) {
	objc.Send[struct{}](o.ID, objc.Sel("webViewWebContentProcessDidTerminate:"), webView)
}

// Tells the delegate that a navigation response became a download.
//
// webView: The web view in which the navigation response took place.
//
// navigationResponse: Descriptive information about the navigation response that turned into a
// download.
//
// download: An object that represents the download of a web resource.
//
// # Discussion
//
// Implement this method to begin tracking download progress.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationDelegate/webView(_:navigationResponse:didBecome:)
func (o WKNavigationDelegateObject) WebViewNavigationResponseDidBecomeDownload(webView IWKWebView, navigationResponse IWKNavigationResponse, download IWKDownload) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:navigationResponse:didBecomeDownload:"), webView, navigationResponse, download)
}

// Tells the delegate that a navigation action became a download.
//
// webView: The web view in which the navigation action took place.
//
// navigationAction: Descriptive information about the navigation response that turned into a
// download.
//
// download: An object that represents the download of a web resource.
//
// # Discussion
//
// Implement this method to begin tracking download progress.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationDelegate/webView(_:navigationAction:didBecome:)
func (o WKNavigationDelegateObject) WebViewNavigationActionDidBecomeDownload(webView IWKWebView, navigationAction IWKNavigationAction, download IWKDownload) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:navigationAction:didBecomeDownload:"), webView, navigationAction, download)
}

// See: https://developer.apple.com/documentation/WebKit/WKNavigationDelegate/webView(_:shouldGoTo:willUseInstantBack:completionHandler:)
func (o WKNavigationDelegateObject) WebViewShouldGoToBackForwardListItemWillUseInstantBackCompletionHandler(webView IWKWebView, backForwardListItem IWKBackForwardListItem, willUseInstantBack bool, completionHandler BoolHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:shouldGoToBackForwardListItem:willUseInstantBack:completionHandler:"), webView, backForwardListItem, willUseInstantBack, completionHandler)
}

// WKNavigationDelegateConfig holds optional typed callbacks for [WKNavigationDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/webkit/wknavigationdelegate
type WKNavigationDelegateConfig struct {

	// Tracking the load progress of a request
	// WebViewDidStartProvisionalNavigation — Tells the delegate that navigation from the main frame has started.
	WebViewDidStartProvisionalNavigation func(webView WKWebView, navigation WKNavigation)
	// WebViewDidReceiveServerRedirectForProvisionalNavigation — Tells the delegate that the web view received a server redirect for a request.
	WebViewDidReceiveServerRedirectForProvisionalNavigation func(webView WKWebView, navigation WKNavigation)

	// Responding to navigation errors
	// WebViewDidFailProvisionalNavigationWithError — Tells the delegate that an error occurred during the early navigation process.
	WebViewDidFailProvisionalNavigationWithError func(webView WKWebView, navigation WKNavigation, error_ foundation.NSError)
	// WebViewWebContentProcessDidTerminate — Tells the delegate that the web view’s content process was terminated.
	WebViewWebContentProcessDidTerminate func(webView WKWebView)

	// Other Methods
	// WebViewDidCommitNavigation — Tells the delegate that the web view has started to receive content for the main frame.
	WebViewDidCommitNavigation func(webView WKWebView, navigation WKNavigation)
	// WebViewDidFinishNavigation — Tells the delegate that navigation is complete.
	WebViewDidFinishNavigation func(webView WKWebView, navigation WKNavigation)
	// WebViewDidFailNavigationWithError — Tells the delegate that an error occurred during navigation.
	WebViewDidFailNavigationWithError func(webView WKWebView, navigation WKNavigation, error_ foundation.NSError)
	// WebViewNavigationResponseDidBecomeDownload — Tells the delegate that a navigation response became a download.
	WebViewNavigationResponseDidBecomeDownload func(webView WKWebView, navigationResponse WKNavigationResponse, download WKDownload)
	// WebViewNavigationActionDidBecomeDownload — Tells the delegate that a navigation action became a download.
	WebViewNavigationActionDidBecomeDownload func(webView WKWebView, navigationAction WKNavigationAction, download WKDownload)
}

// NewWKNavigationDelegate creates an Objective-C object implementing the [WKNavigationDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [WKNavigationDelegateObject] satisfies the [WKNavigationDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/webkit/wknavigationdelegate
func NewWKNavigationDelegate(config WKNavigationDelegateConfig) WKNavigationDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoWKNavigationDelegate_%d", n)

	var methods []objc.MethodDef

	if config.WebViewDidStartProvisionalNavigation != nil {
		fn := config.WebViewDidStartProvisionalNavigation
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("webView:didStartProvisionalNavigation:"),
			Fn: func(self objc.ID, _cmd objc.SEL, webViewID objc.ID, navigationID objc.ID) {
				webView := WKWebViewFromID(webViewID)
				navigation := WKNavigationFromID(navigationID)
				fn(webView, navigation)
			},
		})
	}

	if config.WebViewDidReceiveServerRedirectForProvisionalNavigation != nil {
		fn := config.WebViewDidReceiveServerRedirectForProvisionalNavigation
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("webView:didReceiveServerRedirectForProvisionalNavigation:"),
			Fn: func(self objc.ID, _cmd objc.SEL, webViewID objc.ID, navigationID objc.ID) {
				webView := WKWebViewFromID(webViewID)
				navigation := WKNavigationFromID(navigationID)
				fn(webView, navigation)
			},
		})
	}

	if config.WebViewDidCommitNavigation != nil {
		fn := config.WebViewDidCommitNavigation
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("webView:didCommitNavigation:"),
			Fn: func(self objc.ID, _cmd objc.SEL, webViewID objc.ID, navigationID objc.ID) {
				webView := WKWebViewFromID(webViewID)
				navigation := WKNavigationFromID(navigationID)
				fn(webView, navigation)
			},
		})
	}

	if config.WebViewDidFinishNavigation != nil {
		fn := config.WebViewDidFinishNavigation
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("webView:didFinishNavigation:"),
			Fn: func(self objc.ID, _cmd objc.SEL, webViewID objc.ID, navigationID objc.ID) {
				webView := WKWebViewFromID(webViewID)
				navigation := WKNavigationFromID(navigationID)
				fn(webView, navigation)
			},
		})
	}

	if config.WebViewDidFailNavigationWithError != nil {
		fn := config.WebViewDidFailNavigationWithError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("webView:didFailNavigation:withError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, webViewID objc.ID, navigationID objc.ID, error_ID objc.ID) {
				webView := WKWebViewFromID(webViewID)
				navigation := WKNavigationFromID(navigationID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(webView, navigation, error_)
			},
		})
	}

	if config.WebViewDidFailProvisionalNavigationWithError != nil {
		fn := config.WebViewDidFailProvisionalNavigationWithError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("webView:didFailProvisionalNavigation:withError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, webViewID objc.ID, navigationID objc.ID, error_ID objc.ID) {
				webView := WKWebViewFromID(webViewID)
				navigation := WKNavigationFromID(navigationID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(webView, navigation, error_)
			},
		})
	}

	if config.WebViewWebContentProcessDidTerminate != nil {
		fn := config.WebViewWebContentProcessDidTerminate
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("webViewWebContentProcessDidTerminate:"),
			Fn: func(self objc.ID, _cmd objc.SEL, webViewID objc.ID) {
				webView := WKWebViewFromID(webViewID)
				fn(webView)
			},
		})
	}

	if config.WebViewNavigationResponseDidBecomeDownload != nil {
		fn := config.WebViewNavigationResponseDidBecomeDownload
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("webView:navigationResponse:didBecomeDownload:"),
			Fn: func(self objc.ID, _cmd objc.SEL, webViewID objc.ID, navigationResponseID objc.ID, downloadID objc.ID) {
				webView := WKWebViewFromID(webViewID)
				navigationResponse := WKNavigationResponseFromID(navigationResponseID)
				download := WKDownloadFromID(downloadID)
				fn(webView, navigationResponse, download)
			},
		})
	}

	if config.WebViewNavigationActionDidBecomeDownload != nil {
		fn := config.WebViewNavigationActionDidBecomeDownload
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("webView:navigationAction:didBecomeDownload:"),
			Fn: func(self objc.ID, _cmd objc.SEL, webViewID objc.ID, navigationActionID objc.ID, downloadID objc.ID) {
				webView := WKWebViewFromID(webViewID)
				navigationAction := WKNavigationActionFromID(navigationActionID)
				download := WKDownloadFromID(downloadID)
				fn(webView, navigationAction, download)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("WKNavigationDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewWKNavigationDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return WKNavigationDelegateObjectFromID(instance)
}
