// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKUserContentController] class.
var (
	_WKUserContentControllerClass     WKUserContentControllerClass
	_WKUserContentControllerClassOnce sync.Once
)

func getWKUserContentControllerClass() WKUserContentControllerClass {
	_WKUserContentControllerClassOnce.Do(func() {
		_WKUserContentControllerClass = WKUserContentControllerClass{class: objc.GetClass("WKUserContentController")}
	})
	return _WKUserContentControllerClass
}

// GetWKUserContentControllerClass returns the class object for WKUserContentController.
func GetWKUserContentControllerClass() WKUserContentControllerClass {
	return getWKUserContentControllerClass()
}

type WKUserContentControllerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKUserContentControllerClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKUserContentControllerClass) Alloc() WKUserContentController {
	rv := objc.Send[WKUserContentController](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object for managing interactions between JavaScript code and your web
// view, and for filtering content in your web view.
//
// # Overview
//
// A [WKUserContentController] object provides a bridge between your app and
// the JavaScript code running in the web view. Use this object to do the
// following:
//
// - Inject JavaScript code into webpages running in your web view. - Install
// custom JavaScript functions that call through to your app’s native code.
// - Specify custom filters to prevent the webpage from loading restricted
// content.
//
// Create and configure a [WKUserContentController] object as part of your
// overall web view setup. Assign the object to the [UserContentController]
// property of your [WKWebViewConfiguration] object before creating your web
// view.
//
// # Adding and Removing Custom Scripts
//
//   - [WKUserContentController.AddUserScript]: Injects the specified script into the webpage’s content.
//   - [WKUserContentController.RemoveAllUserScripts]: Removes all user scripts from the web view.
//   - [WKUserContentController.UserScripts]: The user scripts associated with the user content controller.
//
// # Adding and Removing Message Handlers
//
//   - [WKUserContentController.AddScriptMessageHandlerName]: Installs a message handler that you can call from your JavaScript code.
//   - [WKUserContentController.AddScriptMessageHandlerContentWorldName]: Installs a message handler that you can call from the specified content world in your JavaScript code.
//   - [WKUserContentController.AddScriptMessageHandlerWithReplyContentWorldName]: Installs a message handler that returns a reply to your JavaScript code.
//   - [WKUserContentController.RemoveScriptMessageHandlerForName]: Uninstalls the custom message handler with the specified name from your JavaScript code.
//   - [WKUserContentController.RemoveScriptMessageHandlerForNameContentWorld]: Uninstalls a custom message handler from the specified content world in your JavaScript code.
//   - [WKUserContentController.RemoveAllScriptMessageHandlersFromContentWorld]: Uninstalls all custom message handlers from the specified content world in your JavaScript code.
//   - [WKUserContentController.RemoveAllScriptMessageHandlers]: Uninstalls all custom message handlers associated with the user content controller.
//
// # Adding and Removing Content Rules
//
//   - [WKUserContentController.AddContentRuleList]: Adds the specified content rule list to the content controller object.
//   - [WKUserContentController.RemoveContentRuleList]: Removes the specified rule list from the content controller object.
//   - [WKUserContentController.RemoveAllContentRuleLists]: Removes all rules lists from the content controller.
//
// See: https://developer.apple.com/documentation/WebKit/WKUserContentController
type WKUserContentController struct {
	objectivec.Object
}

// WKUserContentControllerFromID constructs a [WKUserContentController] from an objc.ID.
//
// An object for managing interactions between JavaScript code and your web
// view, and for filtering content in your web view.
func WKUserContentControllerFromID(id objc.ID) WKUserContentController {
	return WKUserContentController{objectivec.Object{ID: id}}
}

// NOTE: WKUserContentController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKUserContentController] class.
//
// # Adding and Removing Custom Scripts
//
//   - [IWKUserContentController.AddUserScript]: Injects the specified script into the webpage’s content.
//   - [IWKUserContentController.RemoveAllUserScripts]: Removes all user scripts from the web view.
//   - [IWKUserContentController.UserScripts]: The user scripts associated with the user content controller.
//
// # Adding and Removing Message Handlers
//
//   - [IWKUserContentController.AddScriptMessageHandlerName]: Installs a message handler that you can call from your JavaScript code.
//   - [IWKUserContentController.AddScriptMessageHandlerContentWorldName]: Installs a message handler that you can call from the specified content world in your JavaScript code.
//   - [IWKUserContentController.AddScriptMessageHandlerWithReplyContentWorldName]: Installs a message handler that returns a reply to your JavaScript code.
//   - [IWKUserContentController.RemoveScriptMessageHandlerForName]: Uninstalls the custom message handler with the specified name from your JavaScript code.
//   - [IWKUserContentController.RemoveScriptMessageHandlerForNameContentWorld]: Uninstalls a custom message handler from the specified content world in your JavaScript code.
//   - [IWKUserContentController.RemoveAllScriptMessageHandlersFromContentWorld]: Uninstalls all custom message handlers from the specified content world in your JavaScript code.
//   - [IWKUserContentController.RemoveAllScriptMessageHandlers]: Uninstalls all custom message handlers associated with the user content controller.
//
// # Adding and Removing Content Rules
//
//   - [IWKUserContentController.AddContentRuleList]: Adds the specified content rule list to the content controller object.
//   - [IWKUserContentController.RemoveContentRuleList]: Removes the specified rule list from the content controller object.
//   - [IWKUserContentController.RemoveAllContentRuleLists]: Removes all rules lists from the content controller.
//
// See: https://developer.apple.com/documentation/WebKit/WKUserContentController
type IWKUserContentController interface {
	objectivec.IObject

	// Topic: Adding and Removing Custom Scripts

	// Injects the specified script into the webpage’s content.
	AddUserScript(userScript IWKUserScript)
	// Removes all user scripts from the web view.
	RemoveAllUserScripts()
	// The user scripts associated with the user content controller.
	UserScripts() []WKUserScript

	// Topic: Adding and Removing Message Handlers

	// Installs a message handler that you can call from your JavaScript code.
	AddScriptMessageHandlerName(scriptMessageHandler WKScriptMessageHandler, name string)
	// Installs a message handler that you can call from the specified content world in your JavaScript code.
	AddScriptMessageHandlerContentWorldName(scriptMessageHandler WKScriptMessageHandler, world IWKContentWorld, name string)
	// Installs a message handler that returns a reply to your JavaScript code.
	AddScriptMessageHandlerWithReplyContentWorldName(scriptMessageHandlerWithReply WKScriptMessageHandlerWithReply, contentWorld IWKContentWorld, name string)
	// Uninstalls the custom message handler with the specified name from your JavaScript code.
	RemoveScriptMessageHandlerForName(name string)
	// Uninstalls a custom message handler from the specified content world in your JavaScript code.
	RemoveScriptMessageHandlerForNameContentWorld(name string, contentWorld IWKContentWorld)
	// Uninstalls all custom message handlers from the specified content world in your JavaScript code.
	RemoveAllScriptMessageHandlersFromContentWorld(contentWorld IWKContentWorld)
	// Uninstalls all custom message handlers associated with the user content controller.
	RemoveAllScriptMessageHandlers()

	// Topic: Adding and Removing Content Rules

	// Adds the specified content rule list to the content controller object.
	AddContentRuleList(contentRuleList IWKContentRuleList)
	// Removes the specified rule list from the content controller object.
	RemoveContentRuleList(contentRuleList IWKContentRuleList)
	// Removes all rules lists from the content controller.
	RemoveAllContentRuleLists()

	// The object that coordinates interactions between your app’s native code and the webpage’s scripts and other content.
	UserContentController() IWKUserContentController
	SetUserContentController(value IWKUserContentController)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (u WKUserContentController) Init() WKUserContentController {
	rv := objc.Send[WKUserContentController](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u WKUserContentController) Autorelease() WKUserContentController {
	rv := objc.Send[WKUserContentController](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKUserContentController creates a new WKUserContentController instance.
func NewWKUserContentController() WKUserContentController {
	class := getWKUserContentControllerClass()
	rv := objc.Send[WKUserContentController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Injects the specified script into the webpage’s content.
//
// userScript: The user script to add to the web view’s current page.
//
// See: https://developer.apple.com/documentation/WebKit/WKUserContentController/addUserScript(_:)
func (u WKUserContentController) AddUserScript(userScript IWKUserScript) {
	objc.Send[objc.ID](u.ID, objc.Sel("addUserScript:"), userScript)
}

// Removes all user scripts from the web view.
//
// See: https://developer.apple.com/documentation/WebKit/WKUserContentController/removeAllUserScripts()
func (u WKUserContentController) RemoveAllUserScripts() {
	objc.Send[objc.ID](u.ID, objc.Sel("removeAllUserScripts"))
}

// Installs a message handler that you can call from your JavaScript code.
//
// scriptMessageHandler: The message handler object that implements your custom code. This object
// must adopt the [WKScriptMessageHandler] protocol.
//
// name: The name of the message handler. This parameter must be unique within the
// user content controller and must not be an empty string.
//
// The user content controller uses this parameter to define a JavaScript
// function for your message handler in the page’s main content world. The
// name of this function is
// `window.WebkitXCUIElementTypeMessageHandlers().“XCUIElementTypePostMessage()`,
// where corresponds to the value of this parameter. For example, if you
// specify the string [MyFunction], the user content controller defines the
// `window.WebkitXCUIElementTypeMessageHandlers().MyFunction.PostMessage()`
// function in JavaScript.
//
// # Discussion
//
// To execute your handler’s code, call the JavaScript function that this
// method defines. You may pass a parameter value to the method. The user
// content controller packages that value into an appropriate type and
// delivers it to your content handler’s delegate method.
//
// This method uses the content world from the [PageWorld] property of
// [WKContentWorld].
//
// See: https://developer.apple.com/documentation/WebKit/WKUserContentController/add(_:name:)
func (u WKUserContentController) AddScriptMessageHandlerName(scriptMessageHandler WKScriptMessageHandler, name string) {
	objc.Send[objc.ID](u.ID, objc.Sel("addScriptMessageHandler:name:"), scriptMessageHandler, objc.String(name))
}

// Installs a message handler that you can call from the specified content
// world in your JavaScript code.
//
// scriptMessageHandler: The message handler object that implements your custom code. This object
// must adopt the [WKScriptMessageHandler] protocol.
//
// world: The content world in which to install the message handler. Use the content
// world to scope your message handler to specific parts of your JavaScript
// code.
//
// name: The name of the message handler. This parameter must be unique within the
// user content controller and must not be an empty string.
//
// The user content controller uses this parameter to define a JavaScript
// function for your message handler in all frames in the specified content
// world. The name of this function is
// `window.WebkitXCUIElementTypeMessageHandlers().“XCUIElementTypePostMessage()`,
// where corresponds to the value of this parameter. For example, if you
// specify the string [MyFunction], the user content controller defines the
// `window.WebkitXCUIElementTypeMessageHandlers().MyFunction.PostMessage()`
// function in JavaScript.
//
// # Discussion
//
// To execute your handler’s code, call the JavaScript function that this
// method defines. You may pass a parameter value to the method. The user
// content controller packages that value into an appropriate type and
// delivers it to your content handler’s delegate method.
//
// See: https://developer.apple.com/documentation/WebKit/WKUserContentController/add(_:contentWorld:name:)
func (u WKUserContentController) AddScriptMessageHandlerContentWorldName(scriptMessageHandler WKScriptMessageHandler, world IWKContentWorld, name string) {
	objc.Send[objc.ID](u.ID, objc.Sel("addScriptMessageHandler:contentWorld:name:"), scriptMessageHandler, world, objc.String(name))
}

// Installs a message handler that returns a reply to your JavaScript code.
//
// scriptMessageHandlerWithReply: The message handler object that implements your custom code. This object
// must adopt the [WKScriptMessageHandlerWithReply] protocol.
//
// contentWorld: The content world in which to install the message handler. Use the content
// world to scope your message handler to specific parts of your JavaScript
// code.
//
// name: The name of the message handler. This parameter must be unique within the
// user content controller and must not be an empty string.
//
// The user content controller uses this parameter to define a JavaScript
// function for your message handler in all frames in the specified content
// world. The name of this function is
// `window.WebkitXCUIElementTypeMessageHandlers().“XCUIElementTypePostMessage()`,
// where corresponds to the value of this parameter. For example, if you
// specify the string [MyFunction], the user content controller defines the
// `window.WebkitXCUIElementTypeMessageHandlers().MyFunction.PostMessage()`
// function in JavaScript.
//
// # Discussion
//
// To execute your handler’s code, call the JavaScript function that this
// method defines. You may pass a parameter value to the method. The user
// content controller packages that value into an appropriate type and
// delivers it to your content handler’s delegate method.
//
// See: https://developer.apple.com/documentation/WebKit/WKUserContentController/addScriptMessageHandler(_:contentWorld:name:)
func (u WKUserContentController) AddScriptMessageHandlerWithReplyContentWorldName(scriptMessageHandlerWithReply WKScriptMessageHandlerWithReply, contentWorld IWKContentWorld, name string) {
	objc.Send[objc.ID](u.ID, objc.Sel("addScriptMessageHandlerWithReply:contentWorld:name:"), scriptMessageHandlerWithReply, contentWorld, objc.String(name))
}

// Uninstalls the custom message handler with the specified name from your
// JavaScript code.
//
// name: The name of the message handler to remove. If no message handler with this
// name exists in the user content controller, this method does nothing.
//
// # Discussion
//
// Use this method to remove a message handler that you previously installed
// using the [AddScriptMessageHandlerName] method. This method removes the
// message handler from the page content world — that is, the content world
// available from the [PageWorld] property of [WKContentWorld]. If you
// installed the message handler in a different content world, this method
// doesn’t remove it.
//
// See: https://developer.apple.com/documentation/WebKit/WKUserContentController/removeScriptMessageHandler(forName:)
func (u WKUserContentController) RemoveScriptMessageHandlerForName(name string) {
	objc.Send[objc.ID](u.ID, objc.Sel("removeScriptMessageHandlerForName:"), objc.String(name))
}

// Uninstalls a custom message handler from the specified content world in
// your JavaScript code.
//
// name: The name of the message handler to remove. If no message handler with this
// name exists in the user content controller, this method does nothing.
//
// contentWorld: The content world from which to remove the message handler. For more
// information about content worlds, see [WKContentWorld].
//
// See: https://developer.apple.com/documentation/WebKit/WKUserContentController/removeScriptMessageHandler(forName:contentWorld:)
func (u WKUserContentController) RemoveScriptMessageHandlerForNameContentWorld(name string, contentWorld IWKContentWorld) {
	objc.Send[objc.ID](u.ID, objc.Sel("removeScriptMessageHandlerForName:contentWorld:"), objc.String(name), contentWorld)
}

// Uninstalls all custom message handlers from the specified content world in
// your JavaScript code.
//
// contentWorld: The content world from which to remove the message handler. For more
// information about content worlds, see [WKContentWorld].
//
// # Discussion
//
// Use this method to remove all message handlers in the specified content
// world. Message handlers in other content worlds are unaffected.
//
// See: https://developer.apple.com/documentation/WebKit/WKUserContentController/removeAllScriptMessageHandlers(from:)
func (u WKUserContentController) RemoveAllScriptMessageHandlersFromContentWorld(contentWorld IWKContentWorld) {
	objc.Send[objc.ID](u.ID, objc.Sel("removeAllScriptMessageHandlersFromContentWorld:"), contentWorld)
}

// Uninstalls all custom message handlers associated with the user content
// controller.
//
// # Discussion
//
// Use this method to remove all message handlers in all content worlds in
// your JavaScript code.
//
// See: https://developer.apple.com/documentation/WebKit/WKUserContentController/removeAllScriptMessageHandlers()
func (u WKUserContentController) RemoveAllScriptMessageHandlers() {
	objc.Send[objc.ID](u.ID, objc.Sel("removeAllScriptMessageHandlers"))
}

// Adds the specified content rule list to the content controller object.
//
// contentRuleList: The rule list to add. Create and retrieve rules lists using a
// [WKContentExtensionStore] object.
//
// # Discussion
//
// Call this method to apply a set of content filtering rules to your web
// view’s configuration.
//
// See: https://developer.apple.com/documentation/WebKit/WKUserContentController/add(_:)
func (u WKUserContentController) AddContentRuleList(contentRuleList IWKContentRuleList) {
	objc.Send[objc.ID](u.ID, objc.Sel("addContentRuleList:"), contentRuleList)
}

// Removes the specified rule list from the content controller object.
//
// contentRuleList: The rule list to remove.
//
// # Discussion
//
// This method removes the rule list only from the user content controller
// object. You can still access the rule list from the
// [WKContentRuleListStore] object you used to create it.
//
// See: https://developer.apple.com/documentation/WebKit/WKUserContentController/remove(_:)
func (u WKUserContentController) RemoveContentRuleList(contentRuleList IWKContentRuleList) {
	objc.Send[objc.ID](u.ID, objc.Sel("removeContentRuleList:"), contentRuleList)
}

// Removes all rules lists from the content controller.
//
// # Discussion
//
// This method removes the rule lists only from the user content controller
// object. You can still access rule lists from the [WKContentRuleListStore]
// objects you used to create them.
//
// See: https://developer.apple.com/documentation/WebKit/WKUserContentController/removeAllContentRuleLists()
func (u WKUserContentController) RemoveAllContentRuleLists() {
	objc.Send[objc.ID](u.ID, objc.Sel("removeAllContentRuleLists"))
}
func (u WKUserContentController) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The user scripts associated with the user content controller.
//
// See: https://developer.apple.com/documentation/WebKit/WKUserContentController/userScripts
func (u WKUserContentController) UserScripts() []WKUserScript {
	rv := objc.Send[[]objc.ID](u.ID, objc.Sel("userScripts"))
	return objc.ConvertSlice(rv, func(id objc.ID) WKUserScript {
		return WKUserScriptFromID(id)
	})
}

// The object that coordinates interactions between your app’s native code
// and the webpage’s scripts and other content.
//
// See: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/usercontentcontroller
func (u WKUserContentController) UserContentController() IWKUserContentController {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("userContentController"))
	return WKUserContentControllerFromID(objc.ID(rv))
}
func (u WKUserContentController) SetUserContentController(value IWKUserContentController) {
	objc.Send[struct{}](u.ID, objc.Sel("setUserContentController:"), value)
}
