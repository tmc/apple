// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"fmt"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// The methods for presenting native user interface elements on behalf of a webpage.
//
// See: https://developer.apple.com/documentation/WebKit/WKUIDelegate
type WKUIDelegate interface {
	objectivec.IObject
}

// WKUIDelegateObject wraps an existing Objective-C object that conforms to the WKUIDelegate protocol.
type WKUIDelegateObject struct {
	objectivec.Object
}

func (o WKUIDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// WKUIDelegateObjectFromID constructs a [WKUIDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func WKUIDelegateObjectFromID(id objc.ID) WKUIDelegateObject {
	return WKUIDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Creates a new web view.
//
// webView: The web view invoking the delegate method.
//
// configuration: The configuration to use when creating the new web view.
//
// navigationAction: The navigation action causing the new web view to be created.
//
// windowFeatures: Window features requested by the webpage.
//
// # Return Value
//
// A new web view or `nil`.
//
// # Discussion
//
// The web view returned must be created with the specified configuration.
// WebKit loads the request in the returned web view.
//
// See: https://developer.apple.com/documentation/WebKit/WKUIDelegate/webView(_:createWebViewWith:for:windowFeatures:)
func (o WKUIDelegateObject) WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures(webView IWKWebView, configuration IWKWebViewConfiguration, navigationAction IWKNavigationAction, windowFeatures IWKWindowFeatures) IWKWebView {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("webView:createWebViewWithConfiguration:forNavigationAction:windowFeatures:"), webView, configuration, navigationAction, windowFeatures)
	return WKWebViewFromID(rv)
}

// Notifies your app that the DOM window closed successfully.
//
// webView: The web view invoking the delegate method.
//
// # Discussion
//
// Your app should remove the web view from the view hierarchy and update the
// UI as needed, for instance by closing the containing browser tab or window.
//
// See: https://developer.apple.com/documentation/WebKit/WKUIDelegate/webViewDidClose(_:)
func (o WKUIDelegateObject) WebViewDidClose(webView IWKWebView) {
	objc.Send[struct{}](o.ID, objc.Sel("webViewDidClose:"), webView)
}

// Displays a JavaScript alert panel.
//
// webView: The web view invoking the delegate method.
//
// message: The message to be displayed.
//
// frame: Information about the frame whose JavaScript process initiated this call.
//
// completionHandler: The completion handler to call after the alert panel has been dismissed.
//
// # Discussion
//
// For user security, implementations of this method should call attention to
// the fact that a specific website controls the content in this panel. A
// simple formula for identifying the controlling website is
// `frame.Request().URL.Host()`. The panel should have a single OK button.
//
// See: https://developer.apple.com/documentation/WebKit/WKUIDelegate/webView(_:runJavaScriptAlertPanelWithMessage:initiatedByFrame:completionHandler:)
func (o WKUIDelegateObject) WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler(webView IWKWebView, message string, frame IWKFrameInfo, completionHandler VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:runJavaScriptAlertPanelWithMessage:initiatedByFrame:completionHandler:"), webView, objc.String(message), frame, completionHandler)
}

// Displays a JavaScript confirm panel.
//
// webView: The web view invoking the delegate method.
//
// message: The message to be displayed.
//
// frame: Information about the frame whose JavaScript process initiated this call.
//
// completionHandler: The completion handler to call after the confirm panel has been dismissed.
// Pass true if the user chose OK, and pass false if the user chose Cancel.
//
// # Discussion
//
// For user security, implementations of this method should call attention to
// the fact that a specific website controls the content in this panel. A
// simple formula for identifying the controlling website is
// `frame.Request().URL.Host()`. The panel should have two buttons, typically
// OK and Cancel.
//
// See: https://developer.apple.com/documentation/WebKit/WKUIDelegate/webView(_:runJavaScriptConfirmPanelWithMessage:initiatedByFrame:completionHandler:)
func (o WKUIDelegateObject) WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler(webView IWKWebView, message string, frame IWKFrameInfo, completionHandler BoolHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:runJavaScriptConfirmPanelWithMessage:initiatedByFrame:completionHandler:"), webView, objc.String(message), frame, completionHandler)
}

// Displays a custom Lockdown Mode first use message.
//
// webView: The web view that is requesting to display the Lockdown Mode first use
// dialog.
//
// message: The message for the web view to display if the delegate does not display
// the first use dialog.
//
// completionHandler: A block you must invoke to resume after the web view displays the first use
// dialog. The block does not return a value, and accepts the following
// parameter:
//
// dialogResult: A display result case that indicates how the method handled
// the display request.
//
// # Discussion
//
// Implement this method to display a custom Lockdown Mode message, or to
// suppress the message. Return, or call the completion handler, with a
// [WKDialogResult] case that indicates how your method handled the display
// request. For more information about Lockdown Mode, see [About Lockdown
// Mode].
//
// If you don’t implement this method, the web view displays the default
// message.
//
// See: https://developer.apple.com/documentation/WebKit/WKUIDelegate/webView(_:showLockdownModeFirstUseMessage:completionHandler:)
//
// [About Lockdown Mode]: https://support.apple.com/en-us/HT212650
// [WKDialogResult]: https://developer.apple.com/documentation/WebKit/WKDialogResult
func (o WKUIDelegateObject) WebViewShowLockdownModeFirstUseMessageCompletionHandler(webView IWKWebView, message string, completionHandler WKDialogResultHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:showLockdownModeFirstUseMessage:completionHandler:"), webView, objc.String(message), completionHandler)
}

// Tells the delegate that a contextual menu interaction began.
//
// webView: The web view in which the interaction occurred.
//
// elementInfo: An object that contains information about the element involved in the
// interaction.
//
// completionHandler: The completion handler for you to call with information about how you want
// to handle the interaction. This handler block has no return value and takes
// the following parameter:
//
// configuration: The [UIContextMenuConfiguration] object that contains the
// details of how you want to handle the interaction. Specify `nil` for this
// parameter if you don’t want to show a contextual menu.
//
// See: https://developer.apple.com/documentation/WebKit/WKUIDelegate/webView(_:contextMenuConfigurationForElement:completionHandler:)
//
// [UIContextMenuConfiguration]: https://developer.apple.com/documentation/UIKit/UIContextMenuConfiguration
func (o WKUIDelegateObject) WebViewContextMenuConfigurationForElementCompletionHandler(webView IWKWebView, elementInfo unsafe.Pointer, completionHandler UIContextMenuConfigurationHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:contextMenuConfigurationForElement:completionHandler:"), webView, elementInfo, completionHandler)
}

// Provides the delegate with the animator object that the web view uses to
// display the contextual menu.
//
// webView: The web view in which the interaction occurred.
//
// elementInfo: An object that contains information about the element involved in the
// interaction.
//
// animator: The animator object to use to commit additional animations related to the
// appearance of the contextual menu.
//
// See: https://developer.apple.com/documentation/WebKit/WKUIDelegate/webView(_:contextMenuForElement:willCommitWithAnimator:)
func (o WKUIDelegateObject) WebViewContextMenuForElementWillCommitWithAnimator(webView IWKWebView, elementInfo unsafe.Pointer, animator objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:contextMenuForElement:willCommitWithAnimator:"), webView, elementInfo, animator)
}

// Tells the delegate that the web view is about to present the contextual
// menu for the specified element.
//
// webView: The web view in which the interaction occurred.
//
// elementInfo: An object that contains information about the element involved in the
// interaction.
//
// See: https://developer.apple.com/documentation/WebKit/WKUIDelegate/webView(_:contextMenuWillPresentForElement:)
func (o WKUIDelegateObject) WebViewContextMenuWillPresentForElement(webView IWKWebView, elementInfo unsafe.Pointer) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:contextMenuWillPresentForElement:"), webView, elementInfo)
}

// Tells the delegate that the web view dismissed the contextual menu for the
// specified element.
//
// webView: The web view in which the interaction occurred.
//
// elementInfo: An object that contains information about the element involved in the
// interaction.
//
// See: https://developer.apple.com/documentation/WebKit/WKUIDelegate/webView(_:contextMenuDidEndForElement:)
func (o WKUIDelegateObject) WebViewContextMenuDidEndForElement(webView IWKWebView, elementInfo unsafe.Pointer) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:contextMenuDidEndForElement:"), webView, elementInfo)
}

// Tells the delegate that the web view is about to dismiss an edit menu.
//
// See: https://developer.apple.com/documentation/WebKit/WKUIDelegate/webView(_:willDismissEditMenuWithAnimator:)
func (o WKUIDelegateObject) WebViewWillDismissEditMenuWithAnimator(webView IWKWebView, animator objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:willDismissEditMenuWithAnimator:"), webView, animator)
}

// Tells the delegate that the web view is about to present an edit menu.
//
// See: https://developer.apple.com/documentation/WebKit/WKUIDelegate/webView(_:willPresentEditMenuWithAnimator:)
func (o WKUIDelegateObject) WebViewWillPresentEditMenuWithAnimator(webView IWKWebView, animator objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:willPresentEditMenuWithAnimator:"), webView, animator)
}

// Determines whether a web resource, which the security origin object
// describes, can access the device’s orientation and motion.
//
// webView: The web view requesting permission for orientation and motion information.
//
// origin: An object that identifies the host name, protocol, and port number for a
// web resource.
//
// frame: The frame that initiates the request in the web view.
//
// decisionHandler: A closure that you call from your delegate method. Pass the permission
// decision you determine to the closure.
//
// # Discussion
//
// If you don’t implement this method in your delegate, the system returns
// [WKPermissionDecisionPrompt].
//
// See: https://developer.apple.com/documentation/WebKit/WKUIDelegate/webView(_:requestDeviceOrientationAndMotionPermissionFor:initiatedByFrame:decisionHandler:)
func (o WKUIDelegateObject) WebViewRequestDeviceOrientationAndMotionPermissionForOriginInitiatedByFrameDecisionHandler(webView IWKWebView, origin IWKSecurityOrigin, frame IWKFrameInfo, decisionHandler WKPermissionDecisionHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:requestDeviceOrientationAndMotionPermissionForOrigin:initiatedByFrame:decisionHandler:"), webView, origin, frame, decisionHandler)
}

// Determines whether a web resource, which the security origin object
// describes, can access to the device’s microphone audio and camera video.
//
// webView: The web view requesting permission for microphone audio and camera video.
//
// origin: An object that identifies the host name, protocol, and port number for a
// web resource.
//
// frame: The frame that initiates the request in the web view.
//
// type: An enumeration case representing a type of media capture device, like a
// microphone or camera.
//
// decisionHandler: A closure that you call from your delegate method. Pass the permission
// decision you determine to the closure.
//
// # Discussion
//
// If you don’t implement this method in your delegate, the system returns
// [WKPermissionDecisionPrompt].
//
// See: https://developer.apple.com/documentation/WebKit/WKUIDelegate/webView(_:requestMediaCapturePermissionFor:initiatedByFrame:type:decisionHandler:)
func (o WKUIDelegateObject) WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler(webView IWKWebView, origin IWKSecurityOrigin, frame IWKFrameInfo, type_ WKMediaCaptureType, decisionHandler WKPermissionDecisionHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:requestMediaCapturePermissionForOrigin:initiatedByFrame:type:decisionHandler:"), webView, origin, frame, type_, decisionHandler)
}

// WKUIDelegateConfig holds optional typed callbacks for [WKUIDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/webkit/wkuidelegate
type WKUIDelegateConfig struct {

	// Creating and closing the web view
	// WebViewDidClose — Notifies your app that the DOM window closed successfully.
	WebViewDidClose func(webView WKWebView)

	// Other Methods
	// WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures — Creates a new web view.
	WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures func(webView WKWebView, configuration WKWebViewConfiguration, navigationAction WKNavigationAction, windowFeatures WKWindowFeatures) WKWebView
}

// NewWKUIDelegate creates an Objective-C object implementing the [WKUIDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [WKUIDelegateObject] satisfies the [WKUIDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/webkit/wkuidelegate
func NewWKUIDelegate(config WKUIDelegateConfig) WKUIDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoWKUIDelegate_%d", n)

	var methods []objc.MethodDef

	if config.WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures != nil {
		fn := config.WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("webView:createWebViewWithConfiguration:forNavigationAction:windowFeatures:"),
			Fn: func(self objc.ID, _cmd objc.SEL, webViewID objc.ID, configurationID objc.ID, navigationActionID objc.ID, windowFeaturesID objc.ID) objc.ID {
				webView := WKWebViewFromID(webViewID)
				configuration := WKWebViewConfigurationFromID(configurationID)
				navigationAction := WKNavigationActionFromID(navigationActionID)
				windowFeatures := WKWindowFeaturesFromID(windowFeaturesID)
				return fn(webView, configuration, navigationAction, windowFeatures).GetID()
			},
		})
	}

	if config.WebViewDidClose != nil {
		fn := config.WebViewDidClose
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("webViewDidClose:"),
			Fn: func(self objc.ID, _cmd objc.SEL, webViewID objc.ID) {
				webView := WKWebViewFromID(webViewID)
				fn(webView)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("WKUIDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewWKUIDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return WKUIDelegateObjectFromID(instance)
}
