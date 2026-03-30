// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKWebExtensionController] class.
var (
	_WKWebExtensionControllerClass     WKWebExtensionControllerClass
	_WKWebExtensionControllerClassOnce sync.Once
)

func getWKWebExtensionControllerClass() WKWebExtensionControllerClass {
	_WKWebExtensionControllerClassOnce.Do(func() {
		_WKWebExtensionControllerClass = WKWebExtensionControllerClass{class: objc.GetClass("WKWebExtensionController")}
	})
	return _WKWebExtensionControllerClass
}

// GetWKWebExtensionControllerClass returns the class object for WKWebExtensionController.
func GetWKWebExtensionControllerClass() WKWebExtensionControllerClass {
	return getWKWebExtensionControllerClass()
}

type WKWebExtensionControllerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKWebExtensionControllerClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKWebExtensionControllerClass) Alloc() WKWebExtensionController {
	rv := objc.Send[WKWebExtensionController](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that manages a set of loaded extension contexts.
//
// # Overview
//
// You can have one or more extension controller instances, allowing different
// parts of the app to use different sets of extensions.
//
// You can associate a controller with [WKWebView] using the
// [WebExtensionController] property on [WKWebViewConfiguration].
//
// # Initializers
//
//   - [WKWebExtensionController.InitWithConfiguration]: Returns a web extension controller initialized with the specified configuration.
//
// # Instance Properties
//
//   - [WKWebExtensionController.Configuration]: A copy of the configuration with which the web extension controller was initialized.
//   - [WKWebExtensionController.Delegate]: The extension controller delegate.
//   - [WKWebExtensionController.SetDelegate]
//   - [WKWebExtensionController.ExtensionContexts]: A set of all the currently loaded extension contexts.
//   - [WKWebExtensionController.Extensions]: A set of all the currently loaded extensions.
//
// # Instance Methods
//
//   - [WKWebExtensionController.DidChangeTabPropertiesForTab]: Should be called by the app when the properties of a tab are changed to fire appropriate events with all loaded web extensions.
//   - [WKWebExtensionController.DidCloseWindow]: Should be called by the app when a window is closed to fire appropriate events with all loaded web extensions.
//   - [WKWebExtensionController.DidDeselectTabs]: Should be called by the app when tabs are deselected to fire appropriate events with all loaded web extensions.
//   - [WKWebExtensionController.DidFocusWindow]: Should be called by the app when a window gains focus to fire appropriate events with all loaded web extensions.
//   - [WKWebExtensionController.DidOpenTab]: Should be called by the app when a new tab is opened to fire appropriate events with all loaded web extensions.
//   - [WKWebExtensionController.DidOpenWindow]: Should be called by the app when a new window is opened to fire appropriate events with all loaded web extensions.
//   - [WKWebExtensionController.DidReplaceTabWithTab]: Should be called by the app when a tab is replaced by another tab to fire appropriate events with all loaded web extensions.
//   - [WKWebExtensionController.DidSelectTabs]: Should be called by the app when tabs are selected to fire appropriate events with all loaded web extensions.
//   - [WKWebExtensionController.ExtensionContextForURL]: Returns a loaded extension context matching the specified URL.
//   - [WKWebExtensionController.ExtensionContextForExtension]: Returns a loaded extension context for the specified extension.
//   - [WKWebExtensionController.FetchDataRecordOfTypesForExtensionContextCompletionHandler]: Fetches a data record containing the given extension data types for a specific known web extension context.
//   - [WKWebExtensionController.LoadExtensionContextError]: Loads the specified extension context.
//   - [WKWebExtensionController.RemoveDataOfTypesFromDataRecordsCompletionHandler]: Removes extension data of the given types for the given data records.
//   - [WKWebExtensionController.UnloadExtensionContextError]: Unloads the specified extension context.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController
type WKWebExtensionController struct {
	objectivec.Object
}

// WKWebExtensionControllerFromID constructs a [WKWebExtensionController] from an objc.ID.
//
// An object that manages a set of loaded extension contexts.
func WKWebExtensionControllerFromID(id objc.ID) WKWebExtensionController {
	return WKWebExtensionController{objectivec.Object{ID: id}}
}

// NOTE: WKWebExtensionController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKWebExtensionController] class.
//
// # Initializers
//
//   - [IWKWebExtensionController.InitWithConfiguration]: Returns a web extension controller initialized with the specified configuration.
//
// # Instance Properties
//
//   - [IWKWebExtensionController.Configuration]: A copy of the configuration with which the web extension controller was initialized.
//   - [IWKWebExtensionController.Delegate]: The extension controller delegate.
//   - [IWKWebExtensionController.SetDelegate]
//   - [IWKWebExtensionController.ExtensionContexts]: A set of all the currently loaded extension contexts.
//   - [IWKWebExtensionController.Extensions]: A set of all the currently loaded extensions.
//
// # Instance Methods
//
//   - [IWKWebExtensionController.DidChangeTabPropertiesForTab]: Should be called by the app when the properties of a tab are changed to fire appropriate events with all loaded web extensions.
//   - [IWKWebExtensionController.DidCloseWindow]: Should be called by the app when a window is closed to fire appropriate events with all loaded web extensions.
//   - [IWKWebExtensionController.DidDeselectTabs]: Should be called by the app when tabs are deselected to fire appropriate events with all loaded web extensions.
//   - [IWKWebExtensionController.DidFocusWindow]: Should be called by the app when a window gains focus to fire appropriate events with all loaded web extensions.
//   - [IWKWebExtensionController.DidOpenTab]: Should be called by the app when a new tab is opened to fire appropriate events with all loaded web extensions.
//   - [IWKWebExtensionController.DidOpenWindow]: Should be called by the app when a new window is opened to fire appropriate events with all loaded web extensions.
//   - [IWKWebExtensionController.DidReplaceTabWithTab]: Should be called by the app when a tab is replaced by another tab to fire appropriate events with all loaded web extensions.
//   - [IWKWebExtensionController.DidSelectTabs]: Should be called by the app when tabs are selected to fire appropriate events with all loaded web extensions.
//   - [IWKWebExtensionController.ExtensionContextForURL]: Returns a loaded extension context matching the specified URL.
//   - [IWKWebExtensionController.ExtensionContextForExtension]: Returns a loaded extension context for the specified extension.
//   - [IWKWebExtensionController.FetchDataRecordOfTypesForExtensionContextCompletionHandler]: Fetches a data record containing the given extension data types for a specific known web extension context.
//   - [IWKWebExtensionController.LoadExtensionContextError]: Loads the specified extension context.
//   - [IWKWebExtensionController.RemoveDataOfTypesFromDataRecordsCompletionHandler]: Removes extension data of the given types for the given data records.
//   - [IWKWebExtensionController.UnloadExtensionContextError]: Unloads the specified extension context.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController
type IWKWebExtensionController interface {
	objectivec.IObject

	// Topic: Initializers

	// Returns a web extension controller initialized with the specified configuration.
	InitWithConfiguration(configuration IWKWebExtensionControllerConfiguration) WKWebExtensionController

	// Topic: Instance Properties

	// A copy of the configuration with which the web extension controller was initialized.
	Configuration() IWKWebExtensionControllerConfiguration
	// The extension controller delegate.
	Delegate() WKWebExtensionControllerDelegate
	SetDelegate(value WKWebExtensionControllerDelegate)
	// A set of all the currently loaded extension contexts.
	ExtensionContexts() foundation.INSSet
	// A set of all the currently loaded extensions.
	Extensions() foundation.INSSet

	// Topic: Instance Methods

	// Should be called by the app when the properties of a tab are changed to fire appropriate events with all loaded web extensions.
	DidChangeTabPropertiesForTab(properties WKWebExtensionTabChangedProperties, changedTab WKWebExtensionTab)
	// Should be called by the app when a window is closed to fire appropriate events with all loaded web extensions.
	DidCloseWindow(closedWindow WKWebExtensionWindow)
	// Should be called by the app when tabs are deselected to fire appropriate events with all loaded web extensions.
	DidDeselectTabs(deselectedTabs []objectivec.IObject)
	// Should be called by the app when a window gains focus to fire appropriate events with all loaded web extensions.
	DidFocusWindow(focusedWindow WKWebExtensionWindow)
	// Should be called by the app when a new tab is opened to fire appropriate events with all loaded web extensions.
	DidOpenTab(newTab WKWebExtensionTab)
	// Should be called by the app when a new window is opened to fire appropriate events with all loaded web extensions.
	DidOpenWindow(newWindow WKWebExtensionWindow)
	// Should be called by the app when a tab is replaced by another tab to fire appropriate events with all loaded web extensions.
	DidReplaceTabWithTab(oldTab WKWebExtensionTab, newTab WKWebExtensionTab)
	// Should be called by the app when tabs are selected to fire appropriate events with all loaded web extensions.
	DidSelectTabs(selectedTabs []objectivec.IObject)
	// Returns a loaded extension context matching the specified URL.
	ExtensionContextForURL(URL foundation.INSURL) IWKWebExtensionContext
	// Returns a loaded extension context for the specified extension.
	ExtensionContextForExtension(extension IWKWebExtension) IWKWebExtensionContext
	// Fetches a data record containing the given extension data types for a specific known web extension context.
	FetchDataRecordOfTypesForExtensionContextCompletionHandler(dataTypes foundation.INSSet, extensionContext IWKWebExtensionContext, completionHandler WKWebExtensionDataRecordHandler)
	// Loads the specified extension context.
	LoadExtensionContextError(extensionContext IWKWebExtensionContext) (bool, error)
	// Removes extension data of the given types for the given data records.
	RemoveDataOfTypesFromDataRecordsCompletionHandler(dataTypes foundation.INSSet, dataRecords []WKWebExtensionDataRecord, completionHandler VoidHandler)
	// Unloads the specified extension context.
	UnloadExtensionContextError(extensionContext IWKWebExtensionContext) (bool, error)

	WebExtensionController() IWKWebExtensionController
	SetWebExtensionController(value IWKWebExtensionController)
	// Should be called by the app when a tab is activated to notify all loaded web extensions.
	DidActivateTabPreviousActiveTab(activatedTab WKWebExtensionTab, previousTab WKWebExtensionTab)
	// Should be called by the app when a tab is closed to fire appropriate events with all loaded web extensions.
	DidCloseTabWindowIsClosing(closedTab WKWebExtensionTab, windowIsClosing bool)
	// Should be called by the app when a tab is moved to fire appropriate events with all loaded web extensions.
	DidMoveTabFromIndexInWindow(movedTab WKWebExtensionTab, index uint, oldWindow WKWebExtensionWindow)
}

// Init initializes the instance.
func (w WKWebExtensionController) Init() WKWebExtensionController {
	rv := objc.Send[WKWebExtensionController](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WKWebExtensionController) Autorelease() WKWebExtensionController {
	rv := objc.Send[WKWebExtensionController](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKWebExtensionController creates a new WKWebExtensionController instance.
func NewWKWebExtensionController() WKWebExtensionController {
	class := getWKWebExtensionControllerClass()
	rv := objc.Send[WKWebExtensionController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a web extension controller initialized with the specified
// configuration.
//
// configuration: The configuration for the new web extension controller.
//
// # Return Value
//
// An initialized web extension controller, or nil if the object could not be
// initialized.
//
// # Discussion
//
// This is a designated initializer. You can use [Init] to initialize an
// instance with the default configuration. The initializer copies the
// specified configuration, so mutating the configuration after invoking the
// initializer has no effect on the web extension controller.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/init(configuration:)
func NewWebExtensionControllerWithConfiguration(configuration IWKWebExtensionControllerConfiguration) WKWebExtensionController {
	instance := getWKWebExtensionControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return WKWebExtensionControllerFromID(rv)
}

// Returns a web extension controller initialized with the specified
// configuration.
//
// configuration: The configuration for the new web extension controller.
//
// # Return Value
//
// An initialized web extension controller, or nil if the object could not be
// initialized.
//
// # Discussion
//
// This is a designated initializer. You can use [Init] to initialize an
// instance with the default configuration. The initializer copies the
// specified configuration, so mutating the configuration after invoking the
// initializer has no effect on the web extension controller.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/init(configuration:)
func (w WKWebExtensionController) InitWithConfiguration(configuration IWKWebExtensionControllerConfiguration) WKWebExtensionController {
	rv := objc.Send[WKWebExtensionController](w.ID, objc.Sel("initWithConfiguration:"), configuration)
	return rv
}

// Should be called by the app when the properties of a tab are changed to
// fire appropriate events with all loaded web extensions.
//
// properties: The properties of the tab that were changed.
//
// changedTab: The tab whose properties were changed.
//
// # Discussion
//
// This method informs all loaded extensions of changes to tab properties,
// ensuring a unified understanding across extensions.
//
// If the intention is to inform only a specific extension, you should use the
// respective method on that extension’s context instead.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/didChangeTabProperties(_:for:)
func (w WKWebExtensionController) DidChangeTabPropertiesForTab(properties WKWebExtensionTabChangedProperties, changedTab WKWebExtensionTab) {
	objc.Send[objc.ID](w.ID, objc.Sel("didChangeTabProperties:forTab:"), properties, changedTab)
}

// Should be called by the app when a window is closed to fire appropriate
// events with all loaded web extensions.
//
// # Discussion
//
// - closedWindow: The window that was closed.
//
// This method informs all loaded extensions of the closure of a window,
// ensuring consistent understanding across extensions.
//
// If the intention is to inform only a specific extension, you should use the
// respective method on that extension’s context instead.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/didCloseWindow(_:)
func (w WKWebExtensionController) DidCloseWindow(closedWindow WKWebExtensionWindow) {
	objc.Send[objc.ID](w.ID, objc.Sel("didCloseWindow:"), closedWindow)
}

// Should be called by the app when tabs are deselected to fire appropriate
// events with all loaded web extensions.
//
// deselectedTabs: The set of tabs that were deselected.
//
// # Discussion
//
// This method informs all loaded extensions that tabs have been deselected,
// ensuring consistent understanding across extensions.
//
// If the intention is to inform only a specific extension, you should use the
// respective method on that extension’s context instead.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/didDeselectTabs(_:)
func (w WKWebExtensionController) DidDeselectTabs(deselectedTabs []objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("didDeselectTabs:"), objectivec.IObjectSliceToNSArray(deselectedTabs))
}

// Should be called by the app when a window gains focus to fire appropriate
// events with all loaded web extensions.
//
// focusedWindow: The window that gained focus, or \c nil if no window has focus or a window
// has focus that is not visible to extensions.
//
// # Discussion
//
// This method informs all loaded extensions of the focused window, ensuring
// consistent understanding across extensions.
//
// If the intention is to inform only a specific extension, you should use the
// respective method on that extension’s context instead.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/didFocusWindow(_:)
func (w WKWebExtensionController) DidFocusWindow(focusedWindow WKWebExtensionWindow) {
	objc.Send[objc.ID](w.ID, objc.Sel("didFocusWindow:"), focusedWindow)
}

// Should be called by the app when a new tab is opened to fire appropriate
// events with all loaded web extensions.
//
// newTab: The newly opened tab.
//
// # Discussion
//
// This method informs all loaded extensions of the opening of a new tab,
// ensuring consistent understanding across extensions.
//
// If the intention is to inform only a specific extension, you should use the
// respective method on that extension’s context instead.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/didOpenTab(_:)
func (w WKWebExtensionController) DidOpenTab(newTab WKWebExtensionTab) {
	objc.Send[objc.ID](w.ID, objc.Sel("didOpenTab:"), newTab)
}

// Should be called by the app when a new window is opened to fire appropriate
// events with all loaded web extensions.
//
// newWindow: The newly opened window.
//
// # Discussion
//
// This method informs all loaded extensions of the opening of a new window,
// ensuring consistent understanding across extensions.
//
// If the intention is to inform only a specific extension, you should use the
// respective method on that extension’s context instead.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/didOpenWindow(_:)
func (w WKWebExtensionController) DidOpenWindow(newWindow WKWebExtensionWindow) {
	objc.Send[objc.ID](w.ID, objc.Sel("didOpenWindow:"), newWindow)
}

// Should be called by the app when a tab is replaced by another tab to fire
// appropriate events with all loaded web extensions.
//
// oldTab: The tab that was replaced.
//
// newTab: The tab that replaced the old tab.
//
// # Discussion
//
// This method informs all loaded extensions of the replacement of a tab,
// ensuring consistent understanding across extensions.
//
// If the intention is to inform only a specific extension, you should use the
// respective method on that extension’s context instead.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/didReplaceTab(_:with:)
func (w WKWebExtensionController) DidReplaceTabWithTab(oldTab WKWebExtensionTab, newTab WKWebExtensionTab) {
	objc.Send[objc.ID](w.ID, objc.Sel("didReplaceTab:withTab:"), oldTab, newTab)
}

// Should be called by the app when tabs are selected to fire appropriate
// events with all loaded web extensions.
//
// selectedTabs: The set of tabs that were selected.
//
// # Discussion
//
// This method informs all loaded extensions that tabs have been selected,
// ensuring consistent understanding across extensions.
//
// If the intention is to inform only a specific extension, you should use the
// respective method on that extension’s context instead.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/didSelectTabs(_:)
func (w WKWebExtensionController) DidSelectTabs(selectedTabs []objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("didSelectTabs:"), objectivec.IObjectSliceToNSArray(selectedTabs))
}

// Returns a loaded extension context matching the specified URL.
//
// URL: The URL to lookup.
//
// # Return Value
//
// An extension context or `nil` if no match was found.
//
// # Discussion
//
// This method is useful for determining the extension context to use when
// about to navigate to an extension URL. For example, you could use this
// method to retrieve the appropriate extension context and then use its
// [WebViewConfiguration] property to configure a web view for loading that
// URL.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/extensionContext(for:)-2kr4
func (w WKWebExtensionController) ExtensionContextForURL(URL foundation.INSURL) IWKWebExtensionContext {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("extensionContextForURL:"), URL)
	return WKWebExtensionContextFromID(rv)
}

// Returns a loaded extension context for the specified extension.
//
// extension: An extension to lookup.
//
// # Return Value
//
// An extension context or `nil` if no match was found.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/extensionContext(for:)-6ecpm
func (w WKWebExtensionController) ExtensionContextForExtension(extension IWKWebExtension) IWKWebExtensionContext {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("extensionContextForExtension:"), extension)
	return WKWebExtensionContextFromID(rv)
}

// Fetches a data record containing the given extension data types for a
// specific known web extension context.
//
// dataTypes: The extension data types to fetch records for.
//
// extensionContext: The specific web extension context to fetch records for.
//
// completionHandler: A block to invoke when the data record has been fetched.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/fetchDataRecord(ofTypes:for:completionHandler:)
func (w WKWebExtensionController) FetchDataRecordOfTypesForExtensionContextCompletionHandler(dataTypes foundation.INSSet, extensionContext IWKWebExtensionContext, completionHandler WKWebExtensionDataRecordHandler) {
	_block2, _ := NewWKWebExtensionDataRecordBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("fetchDataRecordOfTypes:forExtensionContext:completionHandler:"), dataTypes, extensionContext, _block2)
}

// Loads the specified extension context.
//
// # Discussion
//
// Causes the context to start, loading any background content, and injecting
// any content into relevant tabs.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/load(_:)
func (w WKWebExtensionController) LoadExtensionContextError(extensionContext IWKWebExtensionContext) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](w.ID, objc.Sel("loadExtensionContext:error:"), extensionContext, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("loadExtensionContext:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Removes extension data of the given types for the given data records.
//
// dataTypes: The extension data types that should be removed.
//
// dataRecords: The extension data records to delete data from.
//
// completionHandler: A block to invoke when the data has been removed.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/removeData(ofTypes:from:completionHandler:)
func (w WKWebExtensionController) RemoveDataOfTypesFromDataRecordsCompletionHandler(dataTypes foundation.INSSet, dataRecords []WKWebExtensionDataRecord, completionHandler VoidHandler) {
	_block2, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("removeDataOfTypes:fromDataRecords:completionHandler:"), dataTypes, dataRecords, _block2)
}

// Unloads the specified extension context.
//
// # Discussion
//
// Causes the context to stop running.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/unload(_:)
func (w WKWebExtensionController) UnloadExtensionContextError(extensionContext IWKWebExtensionContext) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](w.ID, objc.Sel("unloadExtensionContext:error:"), extensionContext, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("unloadExtensionContext:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Should be called by the app when a tab is activated to notify all loaded
// web extensions.
//
// activatedTab: The tab that has become active.
//
// previousTab: The tab that was active before. This parameter can be \c nil if there was
// no previously active tab.
//
// # Discussion
//
// This method informs all loaded extensions of the tab activation, ensuring
// consistent state awareness across extensions.
//
// If the intention is to inform only a specific extension, use the respective
// method on that extension’s context instead.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/didActivateTab:previousActiveTab:
func (w WKWebExtensionController) DidActivateTabPreviousActiveTab(activatedTab WKWebExtensionTab, previousTab WKWebExtensionTab) {
	objc.Send[objc.ID](w.ID, objc.Sel("didActivateTab:previousActiveTab:"), activatedTab, previousTab)
}

// Should be called by the app when a tab is closed to fire appropriate events
// with all loaded web extensions.
//
// closedTab: The tab that was closed.
//
// windowIsClosing: A boolean value indicating whether the window containing the tab is also
// closing.
//
// # Discussion
//
// This method informs all loaded extensions of the closing of a tab, ensuring
// consistent understanding across extensions.
//
// If the intention is to inform only a specific extension, you should use the
// respective method on that extension’s context instead.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/didCloseTab:windowIsClosing:
func (w WKWebExtensionController) DidCloseTabWindowIsClosing(closedTab WKWebExtensionTab, windowIsClosing bool) {
	objc.Send[objc.ID](w.ID, objc.Sel("didCloseTab:windowIsClosing:"), closedTab, windowIsClosing)
}

// Should be called by the app when a tab is moved to fire appropriate events
// with all loaded web extensions.
//
// movedTab: The tab that was moved.
//
// index: The old index of the tab within the window.
//
// oldWindow: The window that the tab was moved from, or `nil` if the tab is moving from
// no open window.
//
// # Discussion
//
// This method informs all loaded extensions of the movement of a tab,
// ensuring consistent understanding across extensions.
//
// If the window is staying the same, the current window should be specified.
// If the intention is to inform only a specific extension, use the respective
// method on that extension’s context instead.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/didMoveTab:fromIndex:inWindow:
func (w WKWebExtensionController) DidMoveTabFromIndexInWindow(movedTab WKWebExtensionTab, index uint, oldWindow WKWebExtensionWindow) {
	objc.Send[objc.ID](w.ID, objc.Sel("didMoveTab:fromIndex:inWindow:"), movedTab, index, oldWindow)
}

// A copy of the configuration with which the web extension controller was
// initialized.
//
// # Discussion
//
// Mutating the configuration has no effect on the web extension controller.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/configuration-swift.property
func (w WKWebExtensionController) Configuration() IWKWebExtensionControllerConfiguration {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("configuration"))
	return WKWebExtensionControllerConfigurationFromID(objc.ID(rv))
}

// The extension controller delegate.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/delegate
func (w WKWebExtensionController) Delegate() WKWebExtensionControllerDelegate {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("delegate"))
	return WKWebExtensionControllerDelegateObjectFromID(rv)
}
func (w WKWebExtensionController) SetDelegate(value WKWebExtensionControllerDelegate) {
	objc.Send[struct{}](w.ID, objc.Sel("setDelegate:"), value)
}

// A set of all the currently loaded extension contexts.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/extensionContexts
func (w WKWebExtensionController) ExtensionContexts() foundation.INSSet {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("extensionContexts"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// A set of all the currently loaded extensions.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/extensions
func (w WKWebExtensionController) Extensions() foundation.INSSet {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("extensions"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/webextensioncontroller
func (w WKWebExtensionController) WebExtensionController() IWKWebExtensionController {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("webExtensionController"))
	return WKWebExtensionControllerFromID(objc.ID(rv))
}
func (w WKWebExtensionController) SetWebExtensionController(value IWKWebExtensionController) {
	objc.Send[struct{}](w.ID, objc.Sel("setWebExtensionController:"), value)
}

// Returns a set of all available extension data types.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/allExtensionDataTypes
func (_WKWebExtensionControllerClass WKWebExtensionControllerClass) AllExtensionDataTypes() foundation.INSSet {
	rv := objc.Send[objc.ID](objc.ID(_WKWebExtensionControllerClass.class), objc.Sel("allExtensionDataTypes"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// FetchDataRecordOfTypesForExtensionContext is a synchronous wrapper around [WKWebExtensionController.FetchDataRecordOfTypesForExtensionContextCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w WKWebExtensionController) FetchDataRecordOfTypesForExtensionContext(ctx context.Context, dataTypes foundation.INSSet, extensionContext IWKWebExtensionContext) (*WKWebExtensionDataRecord, error) {
	done := make(chan *WKWebExtensionDataRecord, 1)
	w.FetchDataRecordOfTypesForExtensionContextCompletionHandler(dataTypes, extensionContext, func(val *WKWebExtensionDataRecord) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
