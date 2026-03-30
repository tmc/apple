// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A group of methods you use to customize web extension interactions.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionControllerDelegate
type WKWebExtensionControllerDelegate interface {
	objectivec.IObject
}

// WKWebExtensionControllerDelegateObject wraps an existing Objective-C object that conforms to the WKWebExtensionControllerDelegate protocol.
type WKWebExtensionControllerDelegateObject struct {
	objectivec.Object
}

func (o WKWebExtensionControllerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// WKWebExtensionControllerDelegateObjectFromID constructs a [WKWebExtensionControllerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func WKWebExtensionControllerDelegateObjectFromID(id objc.ID) WKWebExtensionControllerDelegateObject {
	return WKWebExtensionControllerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Called when an extension context wants to establish a persistent connection
// to an application.
//
// controller: The web extension controller that is managing the extension.
//
// port: A port object for handling the message exchange.
//
// extensionContext: The context in which the web extension is running.
//
// completionHandler: A block to be called when the connection is ready to use, taking an
// optional error. If the connection is successfully established, the error
// should be \c nil.
//
// # Discussion
//
// This method should be implemented by the app to handle establishing
// connections to applications.
//
// The provided [WKWebExtensionMessagePort] object can be used to handle
// message sending, receiving, and disconnection.
//
// You should retain the port object for as long as the connection remains
// active. Releasing the port will disconnect it.
//
// If not implemented, the default behavior is to pass the messages to the app
// extension handler within the extension’s bundle, if the extension was
// loaded from an app extension bundle; otherwise, no action is performed if
// not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionControllerDelegate/webExtensionController(_:connectUsing:for:completionHandler:)
func (o WKWebExtensionControllerDelegateObject) WebExtensionControllerConnectUsingMessagePortForExtensionContextCompletionHandler(controller IWKWebExtensionController, port IWKWebExtensionMessagePort, extensionContext IWKWebExtensionContext, completionHandler ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("webExtensionController:connectUsingMessagePort:forExtensionContext:completionHandler:"), controller, port, extensionContext, completionHandler)
}

// Called when an action’s properties are updated.
//
// controller: The web extension controller initiating the request.
//
// action: The web extension action whose properties are updated.
//
// context: The context within which the web extension is running.
//
// # Discussion
//
// This method is called when an action’s properties are updated and should
// be reflected in the app’s user interface.
//
// The app should ensure that any visible changes, such as icons and labels,
// are updated accordingly.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionControllerDelegate/webExtensionController(_:didUpdate:forExtensionContext:)
func (o WKWebExtensionControllerDelegateObject) WebExtensionControllerDidUpdateActionForExtensionContext(controller IWKWebExtensionController, action IWKWebExtensionAction, context IWKWebExtensionContext) {
	objc.Send[struct{}](o.ID, objc.Sel("webExtensionController:didUpdateAction:forExtensionContext:"), controller, action, context)
}

// Called when an extension context requests the currently focused window.
//
// controller: The web extension controller that is managing the extension.
//
// extensionContext: The context in which the web extension is running.
//
// # Discussion
//
// This method can be optionally implemented by the app to designate the
// window currently in focus to the extension.
//
// If not implemented, the first window in the result of
// [WebExtensionControllerOpenWindowsForExtensionContext] is used.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionControllerDelegate/webExtensionController(_:focusedWindowFor:)
func (o WKWebExtensionControllerDelegateObject) WebExtensionControllerFocusedWindowForExtensionContext(controller IWKWebExtensionController, extensionContext IWKWebExtensionContext) WKWebExtensionWindow {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("webExtensionController:focusedWindowForExtensionContext:"), controller, extensionContext)
	return WKWebExtensionWindowObjectFromID(rv)
}

// Called when an extension context requests a new tab to be opened.
//
// controller: The web extension controller that is managing the extension.
//
// configuration: The configuration specifying how the new tab should be created.
//
// extensionContext: The context in which the web extension is running.
//
// completionHandler: A block to be called with the newly created tab or `nil` if the tab
// wasn’t created. An error should be provided if any errors occurred.
//
// # Discussion
//
// This method should be implemented by the app to handle requests to open new
// tabs. The app can decide how to handle the process based on the provided
// configuration and existing tabs. Once handled, the app should call the
// completion handler with the opened tab or `nil` if the request was declined
// or failed. If not implemented, the extension will be unable to open new
// tabs.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionControllerDelegate/webExtensionController(_:openNewTabUsing:for:completionHandler:)
func (o WKWebExtensionControllerDelegateObject) WebExtensionControllerOpenNewTabUsingConfigurationForExtensionContextCompletionHandler(controller IWKWebExtensionController, configuration IWKWebExtensionTabConfiguration, extensionContext IWKWebExtensionContext, completionHandler WKWebExtensionTabErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("webExtensionController:openNewTabUsingConfiguration:forExtensionContext:completionHandler:"), controller, configuration, extensionContext, completionHandler)
}

// Called when an extension context requests a new window to be opened.
//
// controller: The web extension controller that is managing the extension.
//
// configuration: The configuration specifying how the new window should be created.
//
// extensionContext: The context in which the web extension is running.
//
// completionHandler: A block to be called with the newly created window or `nil` if the window
// wasn’t created. An error should be provided if any errors occurred.
//
// # Discussion
//
// This method should be implemented by the app to handle requests to open new
// windows. The app can decide how to handle the process based on the provided
// configuration and existing windows. Once handled, the app should call the
// completion handler with the opened window or `nil` if the request was
// declined or failed. If not implemented, the extension will be unable to
// open new windows.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionControllerDelegate/webExtensionController(_:openNewWindowUsing:for:completionHandler:)
func (o WKWebExtensionControllerDelegateObject) WebExtensionControllerOpenNewWindowUsingConfigurationForExtensionContextCompletionHandler(controller IWKWebExtensionController, configuration IWKWebExtensionWindowConfiguration, extensionContext IWKWebExtensionContext, completionHandler WKWebExtensionWindowErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("webExtensionController:openNewWindowUsingConfiguration:forExtensionContext:completionHandler:"), controller, configuration, extensionContext, completionHandler)
}

// Called when an extension context requests its options page to be opened.
//
// controller: The web extension controller that is managing the extension.
//
// extensionContext: The context in which the web extension is running.
//
// completionHandler: A block to be called once the options page has been displayed or with an
// error if the page could not be shown.
//
// # Discussion
//
// This method should be implemented by the app to handle requests to display
// the extension’s options page. The app can decide how and where to display
// the options page (e.g., in a new tab or a separate window). The app should
// call the completion handler once the options page is visible to the user,
// or with an error if the operation was declined or failed. If not
// implemented, the options page will be opened in a new tab using the
// [WebExtensionControllerOpenNewTabUsingConfigurationForExtensionContextCompletionHandler]
// delegate method.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionControllerDelegate/webExtensionController(_:openOptionsPageFor:completionHandler:)
func (o WKWebExtensionControllerDelegateObject) WebExtensionControllerOpenOptionsPageForExtensionContextCompletionHandler(controller IWKWebExtensionController, extensionContext IWKWebExtensionContext, completionHandler ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("webExtensionController:openOptionsPageForExtensionContext:completionHandler:"), controller, extensionContext, completionHandler)
}

// Called when an extension context requests the list of ordered open windows.
//
// controller: The web extension controller that is managing the extension.
//
// extensionContext: The context in which the web extension is running.
//
// # Discussion
//
// This method should be implemented by the app to provide the extension with
// the ordered open windows. Depending on your app’s requirements, you may
// return different windows for each extension or the same windows for all
// extensions. The first window in the returned array must correspond to the
// currently focused window and match the result of
// [WebExtensionControllerFocusedWindowForExtensionContext].
//
// If [WebExtensionControllerFocusedWindowForExtensionContext] returns `nil`,
// indicating that no window has focus or the focused window is not visible to
// the extension, the first window in the list returned by this method will be
// considered the presumed focused window. An empty result indicates no open
// windows are available for the extension. Defaults to an empty array if not
// implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionControllerDelegate/webExtensionController(_:openWindowsFor:)
func (o WKWebExtensionControllerDelegateObject) WebExtensionControllerOpenWindowsForExtensionContext(controller IWKWebExtensionController, extensionContext IWKWebExtensionContext) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("webExtensionController:openWindowsForExtensionContext:"), controller, extensionContext)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Called when a popup is requested to be displayed for a specific action.
//
// controller: The web extension controller initiating the request.
//
// action: The action for which the popup is requested.
//
// context: The context within which the web extension is running.
//
// completionHandler: A block to be called once the popup display operation is completed.
//
// # Discussion
//
// This method is called in response to the extension’s scripts or when
// invoking [PerformActionForTab] if the action has a popup.
//
// The associated tab, if applicable, can be located through the
// [AssociatedTab] property of the `action` parameter. This delegate method is
// called when the web view for the popup is fully loaded and ready to
// display. Implementing this method is needed if the app intends to support
// programmatically showing the popup by the extension, although it is
// recommended for handling both programmatic and user-initiated cases.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionControllerDelegate/webExtensionController(_:presentActionPopup:for:completionHandler:)
func (o WKWebExtensionControllerDelegateObject) WebExtensionControllerPresentPopupForActionForExtensionContextCompletionHandler(controller IWKWebExtensionController, action IWKWebExtensionAction, context IWKWebExtensionContext, completionHandler ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("webExtensionController:presentPopupForAction:forExtensionContext:completionHandler:"), controller, action, context, completionHandler)
}

// Called when an extension context wants to send a one-time message to an
// application.
//
// controller: The web extension controller that is managing the extension.
//
// message: The message to be sent.
//
// applicationIdentifier: The unique identifier for the application, or \c nil if none was specified.
//
// extensionContext: The context in which the web extension is running.
//
// replyHandler: A block to be called with a JSON-serializable reply message or an error.
//
// # Discussion
//
// This method should be implemented by the app to handle one-off messages to
// applications.
//
// If not implemented, the default behavior is to pass the message to the app
// extension handler within the extension’s bundle, if the extension was
// loaded from an app extension bundle; otherwise, no action is performed if
// not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionControllerDelegate/webExtensionController(_:sendMessage:toApplicationWithIdentifier:for:replyHandler:)
func (o WKWebExtensionControllerDelegateObject) WebExtensionControllerSendMessageToApplicationWithIdentifierForExtensionContextReplyHandler(controller IWKWebExtensionController, message objectivec.IObject, applicationIdentifier string, extensionContext IWKWebExtensionContext, replyHandler ObjectErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("webExtensionController:sendMessage:toApplicationWithIdentifier:forExtensionContext:replyHandler:"), controller, message, objc.String(applicationIdentifier), extensionContext, replyHandler)
}

// WKWebExtensionControllerDelegateConfig holds optional typed callbacks for [WKWebExtensionControllerDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/webkit/wkwebextensioncontrollerdelegate
type WKWebExtensionControllerDelegateConfig struct {

	// Other Methods
	// WebExtensionControllerDidUpdateActionForExtensionContext — Called when an action’s properties are updated.
	WebExtensionControllerDidUpdateActionForExtensionContext func(controller WKWebExtensionController, action WKWebExtensionAction, context WKWebExtensionContext)
}

// NewWKWebExtensionControllerDelegate creates an Objective-C object implementing the [WKWebExtensionControllerDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [WKWebExtensionControllerDelegateObject] satisfies the [WKWebExtensionControllerDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/webkit/wkwebextensioncontrollerdelegate
func NewWKWebExtensionControllerDelegate(config WKWebExtensionControllerDelegateConfig) WKWebExtensionControllerDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoWKWebExtensionControllerDelegate_%d", n)

	var methods []objc.MethodDef

	if config.WebExtensionControllerDidUpdateActionForExtensionContext != nil {
		fn := config.WebExtensionControllerDidUpdateActionForExtensionContext
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("webExtensionController:didUpdateAction:forExtensionContext:"),
			Fn: func(self objc.ID, _cmd objc.SEL, controllerID objc.ID, actionID objc.ID, contextID objc.ID) {
				controller := WKWebExtensionControllerFromID(controllerID)
				action := WKWebExtensionActionFromID(actionID)
				context := WKWebExtensionContextFromID(contextID)
				fn(controller, action, context)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("WKWebExtensionControllerDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewWKWebExtensionControllerDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return WKWebExtensionControllerDelegateObjectFromID(instance)
}
