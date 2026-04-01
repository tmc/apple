// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"context"
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKWebExtensionContext] class.
var (
	_WKWebExtensionContextClass     WKWebExtensionContextClass
	_WKWebExtensionContextClassOnce sync.Once
)

func getWKWebExtensionContextClass() WKWebExtensionContextClass {
	_WKWebExtensionContextClassOnce.Do(func() {
		_WKWebExtensionContextClass = WKWebExtensionContextClass{class: objc.GetClass("WKWebExtensionContext")}
	})
	return _WKWebExtensionContextClass
}

// GetWKWebExtensionContextClass returns the class object for WKWebExtensionContext.
func GetWKWebExtensionContextClass() WKWebExtensionContextClass {
	return getWKWebExtensionContextClass()
}

type WKWebExtensionContextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKWebExtensionContextClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKWebExtensionContextClass) Alloc() WKWebExtensionContext {
	rv := objc.Send[WKWebExtensionContext](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents the runtime environment for a web extension.
//
// # Overview
//
// This class provides methods for managing the extension’s permissions,
// allowing it to inject content, run background logic, show popovers, and
// display other web-based UI to the user.
//
// # Initializers
//
//   - [WKWebExtensionContext.InitForExtension]: Returns a web extension context initialized with a specified extension.
//
// # Instance Properties
//
//   - [WKWebExtensionContext.BaseURL]: The base URL the context uses for loading extension resources or injecting content into webpages.
//   - [WKWebExtensionContext.SetBaseURL]
//   - [WKWebExtensionContext.Commands]: The commands associated with the extension.
//   - [WKWebExtensionContext.CurrentPermissionMatchPatterns]: The currently granted permission match patterns that have not expired.
//   - [WKWebExtensionContext.CurrentPermissions]: The currently granted permissions that have not expired.
//   - [WKWebExtensionContext.DeniedPermissionMatchPatterns]: The currently denied permission match patterns and their expiration dates.
//   - [WKWebExtensionContext.SetDeniedPermissionMatchPatterns]
//   - [WKWebExtensionContext.DeniedPermissions]: The currently denied permissions and their expiration dates.
//   - [WKWebExtensionContext.SetDeniedPermissions]
//   - [WKWebExtensionContext.Errors]: All errors that occurred in the extension context.
//   - [WKWebExtensionContext.FocusedWindow]: The window that currently has focus for this extension.
//   - [WKWebExtensionContext.GrantedPermissionMatchPatterns]: The currently granted permission match patterns and their expiration dates.
//   - [WKWebExtensionContext.SetGrantedPermissionMatchPatterns]
//   - [WKWebExtensionContext.GrantedPermissions]: The currently granted permissions and their expiration dates.
//   - [WKWebExtensionContext.SetGrantedPermissions]
//   - [WKWebExtensionContext.HasAccessToAllHosts]: A Boolean value indicating if the currently granted permission match patterns set contains the `<all_urls>` pattern or any `*` host patterns.
//   - [WKWebExtensionContext.HasAccessToAllURLs]: A Boolean value indicating if the currently granted permission match patterns set contains the `<all_urls>` pattern.
//   - [WKWebExtensionContext.HasAccessToPrivateData]: A Boolean value indicating if the extension has access to private data.
//   - [WKWebExtensionContext.SetHasAccessToPrivateData]
//   - [WKWebExtensionContext.HasContentModificationRules]: A boolean value indicating whether the extension includes rules used for content modification or blocking.
//   - [WKWebExtensionContext.HasInjectedContent]: A Boolean value indicating whether the extension has script or stylesheet content that can be injected into webpages.
//   - [WKWebExtensionContext.HasRequestedOptionalAccessToAllHosts]: A Boolean value indicating if the extension has requested optional access to all hosts.
//   - [WKWebExtensionContext.SetHasRequestedOptionalAccessToAllHosts]
//   - [WKWebExtensionContext.InspectionName]: The name shown when inspecting the background web view.
//   - [WKWebExtensionContext.SetInspectionName]
//   - [WKWebExtensionContext.Inspectable]: Determines whether Web Inspector can inspect the [WKWebView](<doc://com.apple.webkit/documentation/WebKit/WKWebView>) instances for this context.
//   - [WKWebExtensionContext.SetInspectable]
//   - [WKWebExtensionContext.Loaded]: A Boolean value indicating if this context is loaded in an extension controller.
//   - [WKWebExtensionContext.OpenTabs]: A set of open tabs in all open windows that are exposed to this extension.
//   - [WKWebExtensionContext.OpenWindows]: The open windows that are exposed to this extension.
//   - [WKWebExtensionContext.OptionsPageURL]: The URL of the extension’s options page, if the extension has one.
//   - [WKWebExtensionContext.OverrideNewTabPageURL]: The URL to use as an alternative to the default new tab page, if the extension has one.
//   - [WKWebExtensionContext.UniqueIdentifier]: A unique identifier used to distinguish the extension from other extensions and target it for messages.
//   - [WKWebExtensionContext.SetUniqueIdentifier]
//   - [WKWebExtensionContext.UnsupportedAPIs]: Specifies unsupported APIs for this extension, making them `undefined` in JavaScript.
//   - [WKWebExtensionContext.SetUnsupportedAPIs]
//   - [WKWebExtensionContext.WebExtension]: The extension this context represents.
//   - [WKWebExtensionContext.WebExtensionController]: The extension controller this context is loaded in, otherwise `nil` if it isn’t loaded.
//   - [WKWebExtensionContext.WebViewConfiguration]: The web view configuration to use for web views that load pages from this extension.
//
// # Instance Methods
//
//   - [WKWebExtensionContext.ActionForTab]: Retrieves the extension action for a given tab, or the default action if `nil` is passed.
//   - [WKWebExtensionContext.ClearUserGestureInTab]: Called by the app to clear a user gesture in a specific tab.
//   - [WKWebExtensionContext.CommandForEvent]: Retrieves the command associated with the given event without performing it.
//   - [WKWebExtensionContext.DidChangeTabPropertiesForTab]: Called by the app when the properties of a tab are changed to fire appropriate events with only this extension.
//   - [WKWebExtensionContext.DidCloseWindow]: Called by the app when a window is closed to fire appropriate events with only this extension.
//   - [WKWebExtensionContext.DidDeselectTabs]: Called by the app when tabs are deselected to fire appropriate events with only this extension.
//   - [WKWebExtensionContext.DidFocusWindow]: Called by the app when a window gains focus to fire appropriate events with only this extension.
//   - [WKWebExtensionContext.DidOpenTab]: Called by the app when a new tab is opened to fire appropriate events with only this extension.
//   - [WKWebExtensionContext.DidOpenWindow]: Called by the app when a new window is opened to fire appropriate events with only this extension.
//   - [WKWebExtensionContext.DidReplaceTabWithTab]: Called by the app when a tab is replaced by another tab to fire appropriate events with only this extension.
//   - [WKWebExtensionContext.DidSelectTabs]: Called by the app when tabs are selected to fire appropriate events with only this extension.
//   - [WKWebExtensionContext.HasAccessToURL]: Checks the specified URL against the currently granted permission match patterns.
//   - [WKWebExtensionContext.HasAccessToURLInTab]: Checks the specified URL against the currently granted permission match patterns in a specific tab.
//   - [WKWebExtensionContext.HasActiveUserGestureInTab]: Indicates if a user gesture is currently active in the specified tab.
//   - [WKWebExtensionContext.HasInjectedContentForURL]: Checks if the extension has script or stylesheet content that can be injected into the specified URL.
//   - [WKWebExtensionContext.HasPermission]: Checks the specified permission against the currently granted permissions.
//   - [WKWebExtensionContext.HasPermissionInTab]: Checks the specified permission against the currently granted permissions in a specific tab.
//   - [WKWebExtensionContext.LoadBackgroundContentWithCompletionHandler]: Loads the background content if needed for the extension.
//   - [WKWebExtensionContext.MenuItemsForTab]: Retrieves the menu items for a given tab.
//   - [WKWebExtensionContext.PerformActionForTab]: Performs the extension action associated with the specified tab or performs the default action if `nil` is passed.
//   - [WKWebExtensionContext.PerformCommand]: Performs the specified command, triggering events specific to this extension.
//   - [WKWebExtensionContext.PerformCommandForEvent]: Performs the command associated with the given event.
//   - [WKWebExtensionContext.PermissionStatusForPermission]: Checks the specified permission against the currently denied, granted, and requested permissions.
//   - [WKWebExtensionContext.PermissionStatusForMatchPattern]: Checks the specified match pattern against the currently denied, granted, and requested permission match patterns.
//   - [WKWebExtensionContext.PermissionStatusForURL]: Checks the specified URL against the currently denied, granted, and requested permission match patterns.
//   - [WKWebExtensionContext.PermissionStatusForPermissionInTab]: Checks the specified permission against the currently denied, granted, and requested permissions.
//   - [WKWebExtensionContext.PermissionStatusForURLInTab]: Checks the specified URL against the currently denied, granted, and requested permission match patterns.
//   - [WKWebExtensionContext.PermissionStatusForMatchPatternInTab]: Checks the specified match pattern against the currently denied, granted, and requested permission match patterns.
//   - [WKWebExtensionContext.SetPermissionStatusForPermission]: Sets the status of a permission with a distant future expiration date.
//   - [WKWebExtensionContext.SetPermissionStatusForURL]: Sets the permission status of a URL with a distant future expiration date.
//   - [WKWebExtensionContext.SetPermissionStatusForMatchPattern]: Sets the status of a match pattern with a distant future expiration date.
//   - [WKWebExtensionContext.SetPermissionStatusForURLExpirationDate]: Sets the permission status of a URL with a distant future expiration date.
//   - [WKWebExtensionContext.SetPermissionStatusForPermissionExpirationDate]: Sets the status of a permission with a specific expiration date.
//   - [WKWebExtensionContext.SetPermissionStatusForMatchPatternExpirationDate]: Sets the status of a match pattern with a specific expiration date.
//   - [WKWebExtensionContext.UserGesturePerformedInTab]: Should be called by the app when a user gesture is performed in a specific tab.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext
type WKWebExtensionContext struct {
	objectivec.Object
}

// WKWebExtensionContextFromID constructs a [WKWebExtensionContext] from an objc.ID.
//
// An object that represents the runtime environment for a web extension.
func WKWebExtensionContextFromID(id objc.ID) WKWebExtensionContext {
	return WKWebExtensionContext{objectivec.Object{ID: id}}
}

// NOTE: WKWebExtensionContext adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKWebExtensionContext] class.
//
// # Initializers
//
//   - [IWKWebExtensionContext.InitForExtension]: Returns a web extension context initialized with a specified extension.
//
// # Instance Properties
//
//   - [IWKWebExtensionContext.BaseURL]: The base URL the context uses for loading extension resources or injecting content into webpages.
//   - [IWKWebExtensionContext.SetBaseURL]
//   - [IWKWebExtensionContext.Commands]: The commands associated with the extension.
//   - [IWKWebExtensionContext.CurrentPermissionMatchPatterns]: The currently granted permission match patterns that have not expired.
//   - [IWKWebExtensionContext.CurrentPermissions]: The currently granted permissions that have not expired.
//   - [IWKWebExtensionContext.DeniedPermissionMatchPatterns]: The currently denied permission match patterns and their expiration dates.
//   - [IWKWebExtensionContext.SetDeniedPermissionMatchPatterns]
//   - [IWKWebExtensionContext.DeniedPermissions]: The currently denied permissions and their expiration dates.
//   - [IWKWebExtensionContext.SetDeniedPermissions]
//   - [IWKWebExtensionContext.Errors]: All errors that occurred in the extension context.
//   - [IWKWebExtensionContext.FocusedWindow]: The window that currently has focus for this extension.
//   - [IWKWebExtensionContext.GrantedPermissionMatchPatterns]: The currently granted permission match patterns and their expiration dates.
//   - [IWKWebExtensionContext.SetGrantedPermissionMatchPatterns]
//   - [IWKWebExtensionContext.GrantedPermissions]: The currently granted permissions and their expiration dates.
//   - [IWKWebExtensionContext.SetGrantedPermissions]
//   - [IWKWebExtensionContext.HasAccessToAllHosts]: A Boolean value indicating if the currently granted permission match patterns set contains the `<all_urls>` pattern or any `*` host patterns.
//   - [IWKWebExtensionContext.HasAccessToAllURLs]: A Boolean value indicating if the currently granted permission match patterns set contains the `<all_urls>` pattern.
//   - [IWKWebExtensionContext.HasAccessToPrivateData]: A Boolean value indicating if the extension has access to private data.
//   - [IWKWebExtensionContext.SetHasAccessToPrivateData]
//   - [IWKWebExtensionContext.HasContentModificationRules]: A boolean value indicating whether the extension includes rules used for content modification or blocking.
//   - [IWKWebExtensionContext.HasInjectedContent]: A Boolean value indicating whether the extension has script or stylesheet content that can be injected into webpages.
//   - [IWKWebExtensionContext.HasRequestedOptionalAccessToAllHosts]: A Boolean value indicating if the extension has requested optional access to all hosts.
//   - [IWKWebExtensionContext.SetHasRequestedOptionalAccessToAllHosts]
//   - [IWKWebExtensionContext.InspectionName]: The name shown when inspecting the background web view.
//   - [IWKWebExtensionContext.SetInspectionName]
//   - [IWKWebExtensionContext.Inspectable]: Determines whether Web Inspector can inspect the [WKWebView](<doc://com.apple.webkit/documentation/WebKit/WKWebView>) instances for this context.
//   - [IWKWebExtensionContext.SetInspectable]
//   - [IWKWebExtensionContext.Loaded]: A Boolean value indicating if this context is loaded in an extension controller.
//   - [IWKWebExtensionContext.OpenTabs]: A set of open tabs in all open windows that are exposed to this extension.
//   - [IWKWebExtensionContext.OpenWindows]: The open windows that are exposed to this extension.
//   - [IWKWebExtensionContext.OptionsPageURL]: The URL of the extension’s options page, if the extension has one.
//   - [IWKWebExtensionContext.OverrideNewTabPageURL]: The URL to use as an alternative to the default new tab page, if the extension has one.
//   - [IWKWebExtensionContext.UniqueIdentifier]: A unique identifier used to distinguish the extension from other extensions and target it for messages.
//   - [IWKWebExtensionContext.SetUniqueIdentifier]
//   - [IWKWebExtensionContext.UnsupportedAPIs]: Specifies unsupported APIs for this extension, making them `undefined` in JavaScript.
//   - [IWKWebExtensionContext.SetUnsupportedAPIs]
//   - [IWKWebExtensionContext.WebExtension]: The extension this context represents.
//   - [IWKWebExtensionContext.WebExtensionController]: The extension controller this context is loaded in, otherwise `nil` if it isn’t loaded.
//   - [IWKWebExtensionContext.WebViewConfiguration]: The web view configuration to use for web views that load pages from this extension.
//
// # Instance Methods
//
//   - [IWKWebExtensionContext.ActionForTab]: Retrieves the extension action for a given tab, or the default action if `nil` is passed.
//   - [IWKWebExtensionContext.ClearUserGestureInTab]: Called by the app to clear a user gesture in a specific tab.
//   - [IWKWebExtensionContext.CommandForEvent]: Retrieves the command associated with the given event without performing it.
//   - [IWKWebExtensionContext.DidChangeTabPropertiesForTab]: Called by the app when the properties of a tab are changed to fire appropriate events with only this extension.
//   - [IWKWebExtensionContext.DidCloseWindow]: Called by the app when a window is closed to fire appropriate events with only this extension.
//   - [IWKWebExtensionContext.DidDeselectTabs]: Called by the app when tabs are deselected to fire appropriate events with only this extension.
//   - [IWKWebExtensionContext.DidFocusWindow]: Called by the app when a window gains focus to fire appropriate events with only this extension.
//   - [IWKWebExtensionContext.DidOpenTab]: Called by the app when a new tab is opened to fire appropriate events with only this extension.
//   - [IWKWebExtensionContext.DidOpenWindow]: Called by the app when a new window is opened to fire appropriate events with only this extension.
//   - [IWKWebExtensionContext.DidReplaceTabWithTab]: Called by the app when a tab is replaced by another tab to fire appropriate events with only this extension.
//   - [IWKWebExtensionContext.DidSelectTabs]: Called by the app when tabs are selected to fire appropriate events with only this extension.
//   - [IWKWebExtensionContext.HasAccessToURL]: Checks the specified URL against the currently granted permission match patterns.
//   - [IWKWebExtensionContext.HasAccessToURLInTab]: Checks the specified URL against the currently granted permission match patterns in a specific tab.
//   - [IWKWebExtensionContext.HasActiveUserGestureInTab]: Indicates if a user gesture is currently active in the specified tab.
//   - [IWKWebExtensionContext.HasInjectedContentForURL]: Checks if the extension has script or stylesheet content that can be injected into the specified URL.
//   - [IWKWebExtensionContext.HasPermission]: Checks the specified permission against the currently granted permissions.
//   - [IWKWebExtensionContext.HasPermissionInTab]: Checks the specified permission against the currently granted permissions in a specific tab.
//   - [IWKWebExtensionContext.LoadBackgroundContentWithCompletionHandler]: Loads the background content if needed for the extension.
//   - [IWKWebExtensionContext.MenuItemsForTab]: Retrieves the menu items for a given tab.
//   - [IWKWebExtensionContext.PerformActionForTab]: Performs the extension action associated with the specified tab or performs the default action if `nil` is passed.
//   - [IWKWebExtensionContext.PerformCommand]: Performs the specified command, triggering events specific to this extension.
//   - [IWKWebExtensionContext.PerformCommandForEvent]: Performs the command associated with the given event.
//   - [IWKWebExtensionContext.PermissionStatusForPermission]: Checks the specified permission against the currently denied, granted, and requested permissions.
//   - [IWKWebExtensionContext.PermissionStatusForMatchPattern]: Checks the specified match pattern against the currently denied, granted, and requested permission match patterns.
//   - [IWKWebExtensionContext.PermissionStatusForURL]: Checks the specified URL against the currently denied, granted, and requested permission match patterns.
//   - [IWKWebExtensionContext.PermissionStatusForPermissionInTab]: Checks the specified permission against the currently denied, granted, and requested permissions.
//   - [IWKWebExtensionContext.PermissionStatusForURLInTab]: Checks the specified URL against the currently denied, granted, and requested permission match patterns.
//   - [IWKWebExtensionContext.PermissionStatusForMatchPatternInTab]: Checks the specified match pattern against the currently denied, granted, and requested permission match patterns.
//   - [IWKWebExtensionContext.SetPermissionStatusForPermission]: Sets the status of a permission with a distant future expiration date.
//   - [IWKWebExtensionContext.SetPermissionStatusForURL]: Sets the permission status of a URL with a distant future expiration date.
//   - [IWKWebExtensionContext.SetPermissionStatusForMatchPattern]: Sets the status of a match pattern with a distant future expiration date.
//   - [IWKWebExtensionContext.SetPermissionStatusForURLExpirationDate]: Sets the permission status of a URL with a distant future expiration date.
//   - [IWKWebExtensionContext.SetPermissionStatusForPermissionExpirationDate]: Sets the status of a permission with a specific expiration date.
//   - [IWKWebExtensionContext.SetPermissionStatusForMatchPatternExpirationDate]: Sets the status of a match pattern with a specific expiration date.
//   - [IWKWebExtensionContext.UserGesturePerformedInTab]: Should be called by the app when a user gesture is performed in a specific tab.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext
type IWKWebExtensionContext interface {
	objectivec.IObject

	// Topic: Initializers

	// Returns a web extension context initialized with a specified extension.
	InitForExtension(extension IWKWebExtension) WKWebExtensionContext

	// Topic: Instance Properties

	// The base URL the context uses for loading extension resources or injecting content into webpages.
	BaseURL() foundation.INSURL
	SetBaseURL(value foundation.INSURL)
	// The commands associated with the extension.
	Commands() []WKWebExtensionCommand
	// The currently granted permission match patterns that have not expired.
	CurrentPermissionMatchPatterns() foundation.INSSet
	// The currently granted permissions that have not expired.
	CurrentPermissions() foundation.INSSet
	// The currently denied permission match patterns and their expiration dates.
	DeniedPermissionMatchPatterns() foundation.INSDictionary
	SetDeniedPermissionMatchPatterns(value foundation.INSDictionary)
	// The currently denied permissions and their expiration dates.
	DeniedPermissions() foundation.INSDictionary
	SetDeniedPermissions(value foundation.INSDictionary)
	// All errors that occurred in the extension context.
	Errors() []foundation.NSError
	// The window that currently has focus for this extension.
	FocusedWindow() WKWebExtensionWindow
	// The currently granted permission match patterns and their expiration dates.
	GrantedPermissionMatchPatterns() foundation.INSDictionary
	SetGrantedPermissionMatchPatterns(value foundation.INSDictionary)
	// The currently granted permissions and their expiration dates.
	GrantedPermissions() foundation.INSDictionary
	SetGrantedPermissions(value foundation.INSDictionary)
	// A Boolean value indicating if the currently granted permission match patterns set contains the `<all_urls>` pattern or any `*` host patterns.
	HasAccessToAllHosts() bool
	// A Boolean value indicating if the currently granted permission match patterns set contains the `<all_urls>` pattern.
	HasAccessToAllURLs() bool
	// A Boolean value indicating if the extension has access to private data.
	HasAccessToPrivateData() bool
	SetHasAccessToPrivateData(value bool)
	// A boolean value indicating whether the extension includes rules used for content modification or blocking.
	HasContentModificationRules() bool
	// A Boolean value indicating whether the extension has script or stylesheet content that can be injected into webpages.
	HasInjectedContent() bool
	// A Boolean value indicating if the extension has requested optional access to all hosts.
	HasRequestedOptionalAccessToAllHosts() bool
	SetHasRequestedOptionalAccessToAllHosts(value bool)
	// The name shown when inspecting the background web view.
	InspectionName() string
	SetInspectionName(value string)
	// Determines whether Web Inspector can inspect the [WKWebView](<doc://com.apple.webkit/documentation/WebKit/WKWebView>) instances for this context.
	Inspectable() bool
	SetInspectable(value bool)
	// A Boolean value indicating if this context is loaded in an extension controller.
	Loaded() bool
	// A set of open tabs in all open windows that are exposed to this extension.
	OpenTabs() foundation.INSSet
	// The open windows that are exposed to this extension.
	OpenWindows() []objectivec.IObject
	// The URL of the extension’s options page, if the extension has one.
	OptionsPageURL() foundation.INSURL
	// The URL to use as an alternative to the default new tab page, if the extension has one.
	OverrideNewTabPageURL() foundation.INSURL
	// A unique identifier used to distinguish the extension from other extensions and target it for messages.
	UniqueIdentifier() string
	SetUniqueIdentifier(value string)
	// Specifies unsupported APIs for this extension, making them `undefined` in JavaScript.
	UnsupportedAPIs() foundation.INSSet
	SetUnsupportedAPIs(value foundation.INSSet)
	// The extension this context represents.
	WebExtension() IWKWebExtension
	// The extension controller this context is loaded in, otherwise `nil` if it isn’t loaded.
	WebExtensionController() IWKWebExtensionController
	// The web view configuration to use for web views that load pages from this extension.
	WebViewConfiguration() IWKWebViewConfiguration

	// Topic: Instance Methods

	// Retrieves the extension action for a given tab, or the default action if `nil` is passed.
	ActionForTab(tab WKWebExtensionTab) IWKWebExtensionAction
	// Called by the app to clear a user gesture in a specific tab.
	ClearUserGestureInTab(tab WKWebExtensionTab)
	// Retrieves the command associated with the given event without performing it.
	CommandForEvent(event appkit.NSEvent) IWKWebExtensionCommand
	// Called by the app when the properties of a tab are changed to fire appropriate events with only this extension.
	DidChangeTabPropertiesForTab(properties WKWebExtensionTabChangedProperties, changedTab WKWebExtensionTab)
	// Called by the app when a window is closed to fire appropriate events with only this extension.
	DidCloseWindow(closedWindow WKWebExtensionWindow)
	// Called by the app when tabs are deselected to fire appropriate events with only this extension.
	DidDeselectTabs(deselectedTabs []objectivec.IObject)
	// Called by the app when a window gains focus to fire appropriate events with only this extension.
	DidFocusWindow(focusedWindow WKWebExtensionWindow)
	// Called by the app when a new tab is opened to fire appropriate events with only this extension.
	DidOpenTab(newTab WKWebExtensionTab)
	// Called by the app when a new window is opened to fire appropriate events with only this extension.
	DidOpenWindow(newWindow WKWebExtensionWindow)
	// Called by the app when a tab is replaced by another tab to fire appropriate events with only this extension.
	DidReplaceTabWithTab(oldTab WKWebExtensionTab, newTab WKWebExtensionTab)
	// Called by the app when tabs are selected to fire appropriate events with only this extension.
	DidSelectTabs(selectedTabs []objectivec.IObject)
	// Checks the specified URL against the currently granted permission match patterns.
	HasAccessToURL(url foundation.INSURL) bool
	// Checks the specified URL against the currently granted permission match patterns in a specific tab.
	HasAccessToURLInTab(url foundation.INSURL, tab WKWebExtensionTab) bool
	// Indicates if a user gesture is currently active in the specified tab.
	HasActiveUserGestureInTab(tab WKWebExtensionTab) bool
	// Checks if the extension has script or stylesheet content that can be injected into the specified URL.
	HasInjectedContentForURL(url foundation.INSURL) bool
	// Checks the specified permission against the currently granted permissions.
	HasPermission(permission WKWebExtensionPermission) bool
	// Checks the specified permission against the currently granted permissions in a specific tab.
	HasPermissionInTab(permission WKWebExtensionPermission, tab WKWebExtensionTab) bool
	// Loads the background content if needed for the extension.
	LoadBackgroundContentWithCompletionHandler(completionHandler ErrorHandler)
	// Retrieves the menu items for a given tab.
	MenuItemsForTab(tab WKWebExtensionTab) []appkit.NSMenuItem
	// Performs the extension action associated with the specified tab or performs the default action if `nil` is passed.
	PerformActionForTab(tab WKWebExtensionTab)
	// Performs the specified command, triggering events specific to this extension.
	PerformCommand(command IWKWebExtensionCommand)
	// Performs the command associated with the given event.
	PerformCommandForEvent(event appkit.NSEvent) bool
	// Checks the specified permission against the currently denied, granted, and requested permissions.
	PermissionStatusForPermission(permission WKWebExtensionPermission) WKWebExtensionContextPermissionStatus
	// Checks the specified match pattern against the currently denied, granted, and requested permission match patterns.
	PermissionStatusForMatchPattern(pattern IWKWebExtensionMatchPattern) WKWebExtensionContextPermissionStatus
	// Checks the specified URL against the currently denied, granted, and requested permission match patterns.
	PermissionStatusForURL(url foundation.INSURL) WKWebExtensionContextPermissionStatus
	// Checks the specified permission against the currently denied, granted, and requested permissions.
	PermissionStatusForPermissionInTab(permission WKWebExtensionPermission, tab WKWebExtensionTab) WKWebExtensionContextPermissionStatus
	// Checks the specified URL against the currently denied, granted, and requested permission match patterns.
	PermissionStatusForURLInTab(url foundation.INSURL, tab WKWebExtensionTab) WKWebExtensionContextPermissionStatus
	// Checks the specified match pattern against the currently denied, granted, and requested permission match patterns.
	PermissionStatusForMatchPatternInTab(pattern IWKWebExtensionMatchPattern, tab WKWebExtensionTab) WKWebExtensionContextPermissionStatus
	// Sets the status of a permission with a distant future expiration date.
	SetPermissionStatusForPermission(status WKWebExtensionContextPermissionStatus, permission WKWebExtensionPermission)
	// Sets the permission status of a URL with a distant future expiration date.
	SetPermissionStatusForURL(status WKWebExtensionContextPermissionStatus, url foundation.INSURL)
	// Sets the status of a match pattern with a distant future expiration date.
	SetPermissionStatusForMatchPattern(status WKWebExtensionContextPermissionStatus, pattern IWKWebExtensionMatchPattern)
	// Sets the permission status of a URL with a distant future expiration date.
	SetPermissionStatusForURLExpirationDate(status WKWebExtensionContextPermissionStatus, url foundation.INSURL, expirationDate foundation.INSDate)
	// Sets the status of a permission with a specific expiration date.
	SetPermissionStatusForPermissionExpirationDate(status WKWebExtensionContextPermissionStatus, permission WKWebExtensionPermission, expirationDate foundation.INSDate)
	// Sets the status of a match pattern with a specific expiration date.
	SetPermissionStatusForMatchPatternExpirationDate(status WKWebExtensionContextPermissionStatus, pattern IWKWebExtensionMatchPattern, expirationDate foundation.INSDate)
	// Should be called by the app when a user gesture is performed in a specific tab.
	UserGesturePerformedInTab(tab WKWebExtensionTab)

	// Called by the app when a tab is activated to notify only this specific extension.
	DidActivateTabPreviousActiveTab(activatedTab WKWebExtensionTab, previousTab WKWebExtensionTab)
	// Called by the app when a tab is closed to fire appropriate events with only this extension.
	DidCloseTabWindowIsClosing(closedTab WKWebExtensionTab, windowIsClosing bool)
	// Called by the app when a tab is moved to fire appropriate events with only this extension.
	DidMoveTabFromIndexInWindow(movedTab WKWebExtensionTab, index uint, oldWindow WKWebExtensionWindow)
}

// Init initializes the instance.
func (w WKWebExtensionContext) Init() WKWebExtensionContext {
	rv := objc.Send[WKWebExtensionContext](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WKWebExtensionContext) Autorelease() WKWebExtensionContext {
	rv := objc.Send[WKWebExtensionContext](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKWebExtensionContext creates a new WKWebExtensionContext instance.
func NewWKWebExtensionContext() WKWebExtensionContext {
	class := getWKWebExtensionContextClass()
	rv := objc.Send[WKWebExtensionContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a web extension context initialized with a specified extension.
//
// extension: The extension to use for the new web extension context.
//
// # Return Value
//
// An initialized web extension context.
//
// # Discussion
//
// This is a designated initializer.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/init(for:)
func NewWebExtensionContextForExtension(extension IWKWebExtension) WKWebExtensionContext {
	instance := getWKWebExtensionContextClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initForExtension:"), extension)
	return WKWebExtensionContextFromID(rv)
}

// Returns a web extension context initialized with a specified extension.
//
// extension: The extension to use for the new web extension context.
//
// # Return Value
//
// An initialized web extension context.
//
// # Discussion
//
// This is a designated initializer.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/init(for:)
func (w WKWebExtensionContext) InitForExtension(extension IWKWebExtension) WKWebExtensionContext {
	rv := objc.Send[WKWebExtensionContext](w.ID, objc.Sel("initForExtension:"), extension)
	return rv
}

// Retrieves the extension action for a given tab, or the default action if
// `nil` is passed.
//
// tab: The tab for which to retrieve the extension action, or `nil` to get the
// default action.
//
// # Discussion
//
// The returned object represents the action specific to the tab when
// provided; otherwise, it returns the default action. The default action is
// useful when the context is unrelated to a specific tab. When possible,
// specify the tab to get the most context-relevant action.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/action(for:)
func (w WKWebExtensionContext) ActionForTab(tab WKWebExtensionTab) IWKWebExtensionAction {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("actionForTab:"), tab)
	return WKWebExtensionActionFromID(rv)
}

// Called by the app to clear a user gesture in a specific tab.
//
// tab: The tab from which the user gesture should be cleared.
//
// # Discussion
//
// When a user gesture is no longer relevant in a tab, this method should be
// called to update the extension context.
//
// This will revoke the extension’s access to features that require active
// user interaction, such as `activeTab`. User gestures are automatically
// cleared during navigation in certain scenarios; this method is needed if
// the app intends to clear the gesture more aggressively.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/clearUserGesture(in:)
func (w WKWebExtensionContext) ClearUserGestureInTab(tab WKWebExtensionTab) {
	objc.Send[objc.ID](w.ID, objc.Sel("clearUserGestureInTab:"), tab)
}

// Retrieves the command associated with the given event without performing
// it.
//
// event: The event for which to retrieve the corresponding command.
//
// # Return Value
//
// The command associated with the event, or `nil` if there is no such
// command.
//
// # Discussion
//
// Returns the command that corresponds to the provided event, if such a
// command exists. This provides a way to programmatically determine what
// action would occur for a given event, without triggering the command.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/command(for:)
func (w WKWebExtensionContext) CommandForEvent(event appkit.NSEvent) IWKWebExtensionCommand {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("commandForEvent:"), event)
	return WKWebExtensionCommandFromID(rv)
}

// Called by the app when the properties of a tab are changed to fire
// appropriate events with only this extension.
//
// properties: The properties of the tab that were changed.
//
// changedTab: The tab whose properties were changed.
//
// # Discussion
//
// This method informs only the specific extension of the changes to a tab’s
// properties. If the intention is to inform all loaded extensions
// consistently, you should use the respective method on the extension
// controller instead.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/didChangeTabProperties(_:for:)
func (w WKWebExtensionContext) DidChangeTabPropertiesForTab(properties WKWebExtensionTabChangedProperties, changedTab WKWebExtensionTab) {
	objc.Send[objc.ID](w.ID, objc.Sel("didChangeTabProperties:forTab:"), properties, changedTab)
}

// Called by the app when a window is closed to fire appropriate events with
// only this extension.
//
// closedWindow: The window that was closed.
//
// # Discussion
//
// This method informs only the specific extension of the closure of a window.
// If the intention is to inform all loaded extensions consistently, you
// should use the respective method on the extension controller instead.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/didCloseWindow(_:)
func (w WKWebExtensionContext) DidCloseWindow(closedWindow WKWebExtensionWindow) {
	objc.Send[objc.ID](w.ID, objc.Sel("didCloseWindow:"), closedWindow)
}

// Called by the app when tabs are deselected to fire appropriate events with
// only this extension.
//
// deselectedTabs: The set of tabs that were deselected.
//
// # Discussion
//
// This method informs only the specific extension that tabs have been
// deselected. If the intention is to inform all loaded extensions
// consistently, you should use the respective method on the extension
// controller instead.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/didDeselectTabs(_:)
func (w WKWebExtensionContext) DidDeselectTabs(deselectedTabs []objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("didDeselectTabs:"), objectivec.IObjectSliceToNSArray(deselectedTabs))
}

// Called by the app when a window gains focus to fire appropriate events with
// only this extension.
//
// focusedWindow: The window that gained focus, or `nil` if no window has focus or a window
// has focus that is not visible to this extension.
//
// # Discussion
//
// This method informs only the specific extension that a window has gained
// focus. If the intention is to inform all loaded extensions consistently,
// you should use the respective method on the extension controller instead.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/didFocusWindow(_:)
func (w WKWebExtensionContext) DidFocusWindow(focusedWindow WKWebExtensionWindow) {
	objc.Send[objc.ID](w.ID, objc.Sel("didFocusWindow:"), focusedWindow)
}

// Called by the app when a new tab is opened to fire appropriate events with
// only this extension.
//
// newTab: The newly opened tab.
//
// # Discussion
//
// This method informs only the specific extension of the opening of a new
// tab. If the intention is to inform all loaded extensions consistently, you
// should use the respective method on the extension controller instead.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/didOpenTab(_:)
func (w WKWebExtensionContext) DidOpenTab(newTab WKWebExtensionTab) {
	objc.Send[objc.ID](w.ID, objc.Sel("didOpenTab:"), newTab)
}

// Called by the app when a new window is opened to fire appropriate events
// with only this extension.
//
// newWindow: The newly opened window.
//
// # Discussion
//
// This method informs only the specific extension of the opening of a new
// window. If the intention is to inform all loaded extensions consistently,
// you should use the respective method on the extension controller instead.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/didOpenWindow(_:)
func (w WKWebExtensionContext) DidOpenWindow(newWindow WKWebExtensionWindow) {
	objc.Send[objc.ID](w.ID, objc.Sel("didOpenWindow:"), newWindow)
}

// Called by the app when a tab is replaced by another tab to fire appropriate
// events with only this extension.
//
// oldTab: The tab that was replaced.
//
// newTab: The tab that replaced the old tab.
//
// # Discussion
//
// This method informs only the specific extension that a tab has been
// replaced. If the intention is to inform all loaded extensions consistently,
// you should use the respective method on the extension controller instead.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/didReplaceTab(_:with:)
func (w WKWebExtensionContext) DidReplaceTabWithTab(oldTab WKWebExtensionTab, newTab WKWebExtensionTab) {
	objc.Send[objc.ID](w.ID, objc.Sel("didReplaceTab:withTab:"), oldTab, newTab)
}

// Called by the app when tabs are selected to fire appropriate events with
// only this extension.
//
// selectedTabs: The set of tabs that were selected.
//
// # Discussion
//
// This method informs only the specific extension that tabs have been
// selected. If the intention is to inform all loaded extensions consistently,
// you should use the respective method on the extension controller instead.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/didSelectTabs(_:)
func (w WKWebExtensionContext) DidSelectTabs(selectedTabs []objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("didSelectTabs:"), objectivec.IObjectSliceToNSArray(selectedTabs))
}

// Checks the specified URL against the currently granted permission match
// patterns.
//
// url: The URL for which to return the status.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasAccess(to:)
func (w WKWebExtensionContext) HasAccessToURL(url foundation.INSURL) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasAccessToURL:"), url)
	return rv
}

// Checks the specified URL against the currently granted permission match
// patterns in a specific tab.
//
// url: The URL for which to return the status.
//
// tab: The tab in which to return the permission status, or `nil` if the tab is
// not known or the global status is desired.
//
// # Discussion
//
// Some match patterns can be granted on a per-tab basis. When the tab is
// known, access checks should always use this method.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasAccess(to:in:)
func (w WKWebExtensionContext) HasAccessToURLInTab(url foundation.INSURL, tab WKWebExtensionTab) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasAccessToURL:inTab:"), url, tab)
	return rv
}

// Indicates if a user gesture is currently active in the specified tab.
//
// tab: The tab for which to check for an active user gesture.
//
// # Discussion
//
// An active user gesture may influence the availability of certain
// permissions, such as `activeTab`. User gestures can be triggered by various
// user interactions with the web extension, including clicking on extension
// menu items, executing extension commands, or interacting with extension
// actions. A tab as having an active user gesture enables the extension to
// access features that require user interaction.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasActiveUserGesture(in:)
func (w WKWebExtensionContext) HasActiveUserGestureInTab(tab WKWebExtensionTab) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasActiveUserGestureInTab:"), tab)
	return rv
}

// Checks if the extension has script or stylesheet content that can be
// injected into the specified URL.
//
// url: The webpage URL to check.
//
// # Return Value
//
// Returns [YES] if the extension has content that can be injected by matching
// the URL against the extension’s requested match patterns.
//
// # Discussion
//
// The extension context will still need to be loaded and have granted website
// permissions for its content to actually be injected.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasInjectedContent(for:)
func (w WKWebExtensionContext) HasInjectedContentForURL(url foundation.INSURL) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasInjectedContentForURL:"), url)
	return rv
}

// Checks the specified permission against the currently granted permissions.
//
// permission: The permission for which to return the status.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasPermission(_:)
func (w WKWebExtensionContext) HasPermission(permission WKWebExtensionPermission) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasPermission:"), objc.String(string(permission)))
	return rv
}

// Checks the specified permission against the currently granted permissions
// in a specific tab.
//
// permission: The permission for which to return the status.
//
// tab: The tab in which to return the permission status, or `nil` if the tab is
// not known or the global status is desired.
//
// # Discussion
//
// Permissions can be granted on a per-tab basis. When the tab is known,
// permission checks should always use this method.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasPermission(_:in:)
func (w WKWebExtensionContext) HasPermissionInTab(permission WKWebExtensionPermission, tab WKWebExtensionTab) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasPermission:inTab:"), objc.String(string(permission)), tab)
	return rv
}

// Loads the background content if needed for the extension.
//
// completionHandler: A block to be called upon completion of the loading process, with an
// optional error.
//
// # Discussion
//
// This method forces the loading of the background content for the extension
// that will otherwise be loaded on-demand during specific events.
//
// It is useful when the app requires the background content to be loaded for
// other reasons. If the background content is already loaded, the completion
// handler will be called immediately. An error will occur if the extension
// does not have any background content to load or loading fails.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/loadBackgroundContent(completionHandler:)
func (w WKWebExtensionContext) LoadBackgroundContentWithCompletionHandler(completionHandler ErrorHandler) {
	_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("loadBackgroundContentWithCompletionHandler:"), _block0)
}

// Retrieves the menu items for a given tab.
//
// tab: The tab for which to retrieve the menu items.
//
// # Discussion
//
// Returns menu items provided by the extension, allowing the user to perform
// extension-defined actions on the tab.
//
// The app is responsible for displaying these menu items, typically in a
// context menu or a long-press menu on the tab.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/menuItems(for:)
func (w WKWebExtensionContext) MenuItemsForTab(tab WKWebExtensionTab) []appkit.NSMenuItem {
	rv := objc.Send[[]objc.ID](w.ID, objc.Sel("menuItemsForTab:"), tab)
	return objc.ConvertSlice(rv, func(id objc.ID) appkit.NSMenuItem {
		return appkit.NSMenuItemFromID(id)
	})
}

// Performs the extension action associated with the specified tab or performs
// the default action if `nil` is passed.
//
// tab: The tab for which to perform the extension action, or `nil` to perform the
// default action.
//
// # Discussion
//
// Performing the action will mark the tab, if specified, as having an active
// user gesture. When the `tab` parameter is `nil`, the default action is
// performed. The action can either trigger an event or display a popup,
// depending on how the extension is configured.
//
// If the action is configured to display a popup, implementing the
// appropriate web extension controller delegate method is required;
// otherwise, no action is performed for popup actions.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/performAction(for:)
func (w WKWebExtensionContext) PerformActionForTab(tab WKWebExtensionTab) {
	objc.Send[objc.ID](w.ID, objc.Sel("performActionForTab:"), tab)
}

// Performs the specified command, triggering events specific to this
// extension.
//
// command: The command to be performed.
//
// # Discussion
//
// This method performs the given command as if it was triggered by a user
// gesture within the context of the focused window and active tab.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/performCommand(_:)
func (w WKWebExtensionContext) PerformCommand(command IWKWebExtensionCommand) {
	objc.Send[objc.ID](w.ID, objc.Sel("performCommand:"), command)
}

// Performs the command associated with the given event.
//
// event: The event representing the user input.
//
// # Return Value
//
// Returns [YES] if a command corresponding to the event was found and
// performed, [NO] otherwise.
//
// # Discussion
//
// This method checks for a command corresponding to the provided event and
// performs it, if available. The app should use this method to perform any
// extension commands at an appropriate time in the app’s event handling,
// like in [sendEvent(_:)] of [NSApplication] or [sendEvent(_:)] of [NSWindow]
// subclasses.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/performCommand(for:)-8btj0
//
// [NSApplication]: https://developer.apple.com/documentation/AppKit/NSApplication
// [NSWindow]: https://developer.apple.com/documentation/AppKit/NSWindow
// [sendEvent(_:)]: https://developer.apple.com/documentation/AppKit/NSWindow/sendEvent(_:)
func (w WKWebExtensionContext) PerformCommandForEvent(event appkit.NSEvent) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("performCommandForEvent:"), event)
	return rv
}

// Checks the specified permission against the currently denied, granted, and
// requested permissions.
//
// permission: The permission for which to return the status.
//
// # Discussion
//
// Permissions can be granted on a per-tab basis. When the tab is known,
// access checks should always use the method that checks in a tab.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/permissionStatus(for:)-3qq2w
func (w WKWebExtensionContext) PermissionStatusForPermission(permission WKWebExtensionPermission) WKWebExtensionContextPermissionStatus {
	rv := objc.Send[WKWebExtensionContextPermissionStatus](w.ID, objc.Sel("permissionStatusForPermission:"), objc.String(string(permission)))
	return WKWebExtensionContextPermissionStatus(rv)
}

// Checks the specified match pattern against the currently denied, granted,
// and requested permission match patterns.
//
// pattern: The pattern for which to return the status.
//
// # Discussion
//
// Match patterns can be granted on a per-tab basis. When the tab is known,
// access checks should always use the method that checks in a tab.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/permissionStatus(for:)-7mu8
func (w WKWebExtensionContext) PermissionStatusForMatchPattern(pattern IWKWebExtensionMatchPattern) WKWebExtensionContextPermissionStatus {
	rv := objc.Send[WKWebExtensionContextPermissionStatus](w.ID, objc.Sel("permissionStatusForMatchPattern:"), pattern)
	return WKWebExtensionContextPermissionStatus(rv)
}

// Checks the specified URL against the currently denied, granted, and
// requested permission match patterns.
//
// url: The URL for which to return the status.
//
// # Discussion
//
// URLs and match patterns can be granted on a per-tab basis. When the tab is
// known, access checks should always use the method that checks in a tab.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/permissionStatus(for:)-7ojrb
func (w WKWebExtensionContext) PermissionStatusForURL(url foundation.INSURL) WKWebExtensionContextPermissionStatus {
	rv := objc.Send[WKWebExtensionContextPermissionStatus](w.ID, objc.Sel("permissionStatusForURL:"), url)
	return WKWebExtensionContextPermissionStatus(rv)
}

// Checks the specified permission against the currently denied, granted, and
// requested permissions.
//
// permission: The permission for which to return the status.
//
// tab: The tab in which to return the permission status, or `nil` if the tab is
// not known or the global status is desired.
//
// # Discussion
//
// Permissions can be granted on a per-tab basis. When the tab is known,
// access checks should always specify the tab.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/permissionStatus(for:in:)-4h82n
func (w WKWebExtensionContext) PermissionStatusForPermissionInTab(permission WKWebExtensionPermission, tab WKWebExtensionTab) WKWebExtensionContextPermissionStatus {
	rv := objc.Send[WKWebExtensionContextPermissionStatus](w.ID, objc.Sel("permissionStatusForPermission:inTab:"), objc.String(string(permission)), tab)
	return WKWebExtensionContextPermissionStatus(rv)
}

// Checks the specified URL against the currently denied, granted, and
// requested permission match patterns.
//
// url: The URL for which to return the status.
//
// tab: The tab in which to return the permission status, or `nil` if the tab is
// not known or the global status is desired.
//
// # Discussion
//
// URLs and match patterns can be granted on a per-tab basis. When the tab is
// known, access checks should always use this method.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/permissionStatus(for:in:)-96xaf
func (w WKWebExtensionContext) PermissionStatusForURLInTab(url foundation.INSURL, tab WKWebExtensionTab) WKWebExtensionContextPermissionStatus {
	rv := objc.Send[WKWebExtensionContextPermissionStatus](w.ID, objc.Sel("permissionStatusForURL:inTab:"), url, tab)
	return WKWebExtensionContextPermissionStatus(rv)
}

// Checks the specified match pattern against the currently denied, granted,
// and requested permission match patterns.
//
// pattern: The pattern for which to return the status.
//
// tab: The tab in which to return the permission status, or `nil` if the tab is
// not known or the global status is desired.
//
// # Discussion
//
// Match patterns can be granted on a per-tab basis. When the tab is known,
// access checks should always use this method.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/permissionStatus(for:in:)-nqhm
func (w WKWebExtensionContext) PermissionStatusForMatchPatternInTab(pattern IWKWebExtensionMatchPattern, tab WKWebExtensionTab) WKWebExtensionContextPermissionStatus {
	rv := objc.Send[WKWebExtensionContextPermissionStatus](w.ID, objc.Sel("permissionStatusForMatchPattern:inTab:"), pattern, tab)
	return WKWebExtensionContextPermissionStatus(rv)
}

// Sets the status of a permission with a distant future expiration date.
//
// status: The new permission status to set for the given permission.
//
// permission: The permission for which to set the status.
//
// # Discussion
//
// This method will update [GrantedPermissions] and [DeniedPermissions]. Use
// this method for changing a single permission’s status.
//
// Only [WKWebExtensionContextPermissionStatusDeniedExplicitly],
// [WKWebExtensionContextPermissionStatusUnknown], and
// [WKWebExtensionContextPermissionStatusGrantedExplicitly] states are allowed
// to be set using this method.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/setPermissionStatus(_:for:)-4u95f
func (w WKWebExtensionContext) SetPermissionStatusForPermission(status WKWebExtensionContextPermissionStatus, permission WKWebExtensionPermission) {
	objc.Send[objc.ID](w.ID, objc.Sel("setPermissionStatus:forPermission:"), status, objc.String(string(permission)))
}

// Sets the permission status of a URL with a distant future expiration date.
//
// status: The new permission status to set for the given URL.
//
// url: The URL for which to set the status.
//
// # Discussion
//
// The URL is converted into a match pattern and will update
// [GrantedPermissionMatchPatterns] and [DeniedPermissionMatchPatterns]. Use
// this method for changing a single URL’s status. Only
// [WKWebExtensionContextPermissionStatusDeniedExplicitly],
// [WKWebExtensionContextPermissionStatusUnknown], and
// [WKWebExtensionContextPermissionStatusGrantedExplicitly] states are allowed
// to be set using this method.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/setPermissionStatus(_:for:)-5xahd
func (w WKWebExtensionContext) SetPermissionStatusForURL(status WKWebExtensionContextPermissionStatus, url foundation.INSURL) {
	objc.Send[objc.ID](w.ID, objc.Sel("setPermissionStatus:forURL:"), status, url)
}

// Sets the status of a match pattern with a distant future expiration date.
//
// status: The new permission status to set for the given match pattern.
//
// pattern: The match pattern for which to set the status.
//
// # Discussion
//
// This method will update [GrantedPermissionMatchPatterns] and
// [DeniedPermissionMatchPatterns]. Use this method for changing a single
// match pattern’s status. Only
// [WKWebExtensionContextPermissionStatusDeniedExplicitly],
// [WKWebExtensionContextPermissionStatusUnknown], and
// [WKWebExtensionContextPermissionStatusGrantedExplicitly] states are allowed
// to be set using this method.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/setPermissionStatus(_:for:)-6auqv
func (w WKWebExtensionContext) SetPermissionStatusForMatchPattern(status WKWebExtensionContextPermissionStatus, pattern IWKWebExtensionMatchPattern) {
	objc.Send[objc.ID](w.ID, objc.Sel("setPermissionStatus:forMatchPattern:"), status, pattern)
}

// Sets the permission status of a URL with a distant future expiration date.
//
// status: The new permission status to set for the given URL.
//
// url: The URL for which to set the status.
//
// expirationDate: The expiration date for the new permission status, or `nil` for distant
// future.
//
// # Discussion
//
// The URL is converted into a match pattern and will update
// [GrantedPermissionMatchPatterns] and [DeniedPermissionMatchPatterns]. Use
// this method for changing a single URL’s status. Passing a `nil`
// expiration date will be treated as a distant future date. Only
// [WKWebExtensionContextPermissionStatusDeniedExplicitly],
// [WKWebExtensionContextPermissionStatusUnknown], and
// [WKWebExtensionContextPermissionStatusGrantedExplicitly] states are allowed
// to be set using this method.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/setPermissionStatus(_:for:expirationDate:)-5q9id
func (w WKWebExtensionContext) SetPermissionStatusForURLExpirationDate(status WKWebExtensionContextPermissionStatus, url foundation.INSURL, expirationDate foundation.INSDate) {
	objc.Send[objc.ID](w.ID, objc.Sel("setPermissionStatus:forURL:expirationDate:"), status, url, expirationDate)
}

// Sets the status of a permission with a specific expiration date.
//
// status: The new permission status to set for the given permission.
//
// permission: The permission for which to set the status.
//
// expirationDate: The expiration date for the new permission status, or \c nil for distant
// future.
//
// # Discussion
//
// This method will update [GrantedPermissions] and [DeniedPermissions]. Use
// this method for changing a single permission’s status.
//
// Passing a `nil` expiration date will be treated as a distant future date.
// Only [WKWebExtensionContextPermissionStatusDeniedExplicitly],
// [WKWebExtensionContextPermissionStatusUnknown], and
// [WKWebExtensionContextPermissionStatusGrantedExplicitly] states are allowed
// to be set using this method.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/setPermissionStatus(_:for:expirationDate:)-692ui
func (w WKWebExtensionContext) SetPermissionStatusForPermissionExpirationDate(status WKWebExtensionContextPermissionStatus, permission WKWebExtensionPermission, expirationDate foundation.INSDate) {
	objc.Send[objc.ID](w.ID, objc.Sel("setPermissionStatus:forPermission:expirationDate:"), status, objc.String(string(permission)), expirationDate)
}

// Sets the status of a match pattern with a specific expiration date.
//
// status: The new permission status to set for the given match pattern.
//
// pattern: The match pattern for which to set the status.
//
// expirationDate: The expiration date for the new permission status, or `nil` for distant
// future.
//
// # Discussion
//
// This method will update [GrantedPermissionMatchPatterns] and
// [DeniedPermissionMatchPatterns]. Use this method for changing a single
// match pattern’s status.
//
// Passing a `nil` expiration date will be treated as a distant future date.
// Only [WKWebExtensionContextPermissionStatusDeniedExplicitly],
// [WKWebExtensionContextPermissionStatusUnknown], and
// [WKWebExtensionContextPermissionStatusGrantedExplicitly] states are allowed
// to be set using this method.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/setPermissionStatus(_:for:expirationDate:)-7038f
func (w WKWebExtensionContext) SetPermissionStatusForMatchPatternExpirationDate(status WKWebExtensionContextPermissionStatus, pattern IWKWebExtensionMatchPattern, expirationDate foundation.INSDate) {
	objc.Send[objc.ID](w.ID, objc.Sel("setPermissionStatus:forMatchPattern:expirationDate:"), status, pattern, expirationDate)
}

// Should be called by the app when a user gesture is performed in a specific
// tab.
//
// tab: The tab in which the user gesture was performed.
//
// # Discussion
//
// When a user gesture is performed in a tab, this method should be called to
// update the extension context.
//
// This enables the extension to be aware of the user gesture, potentially
// granting it access to features that require user interaction, such as
// `activeTab`. Not required if using [PerformActionForTab].
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/userGesturePerformed(in:)
func (w WKWebExtensionContext) UserGesturePerformedInTab(tab WKWebExtensionTab) {
	objc.Send[objc.ID](w.ID, objc.Sel("userGesturePerformedInTab:"), tab)
}

// Called by the app when a tab is activated to notify only this specific
// extension.
//
// activatedTab: The tab that has become active.
//
// previousTab: The tab that was active before. This parameter can be `nil` if there was no
// previously active tab.
//
// # Discussion
//
// This method informs only the specific extension of the tab activation. If
// the intention is to inform all loaded extensions consistently, you should
// use the respective method on the extension controller instead.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/didActivateTab:previousActiveTab:
func (w WKWebExtensionContext) DidActivateTabPreviousActiveTab(activatedTab WKWebExtensionTab, previousTab WKWebExtensionTab) {
	objc.Send[objc.ID](w.ID, objc.Sel("didActivateTab:previousActiveTab:"), activatedTab, previousTab)
}

// Called by the app when a tab is closed to fire appropriate events with only
// this extension.
//
// closedTab: The tab that was closed.
//
// windowIsClosing: A Boolean value indicating whether the window containing the tab is also
// closing.
//
// # Discussion
//
// This method informs only the specific extension of the closure of a tab. If
// the intention is to inform all loaded extensions consistently, you should
// use the respective method on the extension controller instead.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/didCloseTab:windowIsClosing:
func (w WKWebExtensionContext) DidCloseTabWindowIsClosing(closedTab WKWebExtensionTab, windowIsClosing bool) {
	objc.Send[objc.ID](w.ID, objc.Sel("didCloseTab:windowIsClosing:"), closedTab, windowIsClosing)
}

// Called by the app when a tab is moved to fire appropriate events with only
// this extension.
//
// movedTab: The tab that was moved.
//
// index: The old index of the tab within the window.
//
// oldWindow: The window that the tab was moved from, or `nil` if the tab isn’t moving
// from an open window.
//
// # Discussion
//
// If the window is staying the same, the current window should be specified.
// This method informs only the specific extension that a tab has been moved.
// If the intention is to inform all loaded extensions consistently, you
// should use the respective method on the extension controller instead.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/didMoveTab:fromIndex:inWindow:
func (w WKWebExtensionContext) DidMoveTabFromIndexInWindow(movedTab WKWebExtensionTab, index uint, oldWindow WKWebExtensionWindow) {
	objc.Send[objc.ID](w.ID, objc.Sel("didMoveTab:fromIndex:inWindow:"), movedTab, index, oldWindow)
}

// Returns a web extension context initialized with the specified extension.
//
// extension: The extension to use for the new web extension context.
//
// # Return Value
//
// An initialized web extension context.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/contextForExtension:
func (_WKWebExtensionContextClass WKWebExtensionContextClass) ContextForExtension(extension IWKWebExtension) WKWebExtensionContext {
	rv := objc.Send[objc.ID](objc.ID(_WKWebExtensionContextClass.class), objc.Sel("contextForExtension:"), extension)
	return WKWebExtensionContextFromID(rv)
}

// The base URL the context uses for loading extension resources or injecting
// content into webpages.
//
// # Discussion
//
// The default value is a unique URL using the `webkit-extension` scheme. The
// base URL can be set to any URL, but only the scheme and host will be used.
// The scheme cannot be a scheme that is already supported by [WKWebView]
// (e.g. http, https, etc.). Setting is only allowed when the context is not
// loaded.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/baseURL
func (w WKWebExtensionContext) BaseURL() foundation.INSURL {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("baseURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (w WKWebExtensionContext) SetBaseURL(value foundation.INSURL) {
	objc.Send[struct{}](w.ID, objc.Sel("setBaseURL:"), value)
}

// The commands associated with the extension.
//
// # Discussion
//
// Provides all commands registered within the extension. Each command
// represents an action or behavior available for the web extension.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/commands
func (w WKWebExtensionContext) Commands() []WKWebExtensionCommand {
	rv := objc.Send[[]objc.ID](w.ID, objc.Sel("commands"))
	return objc.ConvertSlice(rv, func(id objc.ID) WKWebExtensionCommand {
		return WKWebExtensionCommandFromID(id)
	})
}

// The currently granted permission match patterns that have not expired.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/currentPermissionMatchPatterns
func (w WKWebExtensionContext) CurrentPermissionMatchPatterns() foundation.INSSet {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("currentPermissionMatchPatterns"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The currently granted permissions that have not expired.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/currentPermissions
func (w WKWebExtensionContext) CurrentPermissions() foundation.INSSet {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("currentPermissions"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The currently denied permission match patterns and their expiration dates.
//
// # Discussion
//
// Match patterns that don’t expire will have a distant future date. This
// will never include expired entries at time of access.
//
// Setting this property will replace all existing entries. Use this property
// for saving and restoring permission status in bulk.
//
// Match patterns in this dictionary should be explicitly denied by the user
// before being added. Any match pattern in this collection will not be
// presented for approval again until they expire. This value should be saved
// and restored as needed by the app.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/deniedPermissionMatchPatterns
func (w WKWebExtensionContext) DeniedPermissionMatchPatterns() foundation.INSDictionary {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("deniedPermissionMatchPatterns"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (w WKWebExtensionContext) SetDeniedPermissionMatchPatterns(value foundation.INSDictionary) {
	objc.Send[struct{}](w.ID, objc.Sel("setDeniedPermissionMatchPatterns:"), value)
}

// The currently denied permissions and their expiration dates.
//
// # Discussion
//
// Permissions that don’t expire will have a distant future date. This will
// never include expired entries at time of access.
//
// Setting this property will replace all existing entries. Use this property
// for saving and restoring permission status in bulk.
//
// Permissions in this dictionary should be explicitly denied by the user
// before being added. Any match pattern in this collection will not be
// presented for approval again until they expire. This value should be saved
// and restored as needed by the app.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/deniedPermissions
func (w WKWebExtensionContext) DeniedPermissions() foundation.INSDictionary {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("deniedPermissions"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (w WKWebExtensionContext) SetDeniedPermissions(value foundation.INSDictionary) {
	objc.Send[struct{}](w.ID, objc.Sel("setDeniedPermissions:"), value)
}

// All errors that occurred in the extension context.
//
// # Discussion
//
// Provides an array of all parse-time and runtime errors for the extension
// and extension context, with repeat errors consolidated into a single entry
// for the original occurrence. If no errors occurred, an empty array is
// returned.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/errors
func (w WKWebExtensionContext) Errors() []foundation.NSError {
	rv := objc.Send[[]objc.ID](w.ID, objc.Sel("errors"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSError {
		return foundation.NSErrorFromID(id)
	})
}

// The window that currently has focus for this extension.
//
// # Discussion
//
// Provides the window that currently has focus, as set by the
// [DidFocusWindow] method.
//
// It will be `nil` if no window has focus or if a window has focus that is
// not visible to the extension. Initially populated by the window returned by
// the extension controller delegate method
// [WebExtensionControllerFocusedWindowForExtensionContext].
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/focusedWindow
func (w WKWebExtensionContext) FocusedWindow() WKWebExtensionWindow {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("focusedWindow"))
	return WKWebExtensionWindowObjectFromID(rv)
}

// The currently granted permission match patterns and their expiration dates.
//
// # Discussion
//
// Match patterns that don’t expire will have a distant future date. This
// will never include expired entries at time of access.
//
// Setting this property will replace all existing entries. Use this property
// for saving and restoring permission status in bulk.
//
// Match patterns in this dictionary should be explicitly granted by the user
// before being added. Any match pattern in this collection will not be
// presented for approval again until they expire. This value should be saved
// and restored as needed by the app.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/grantedPermissionMatchPatterns
func (w WKWebExtensionContext) GrantedPermissionMatchPatterns() foundation.INSDictionary {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("grantedPermissionMatchPatterns"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (w WKWebExtensionContext) SetGrantedPermissionMatchPatterns(value foundation.INSDictionary) {
	objc.Send[struct{}](w.ID, objc.Sel("setGrantedPermissionMatchPatterns:"), value)
}

// The currently granted permissions and their expiration dates.
//
// # Discussion
//
// Permissions that don’t expire will have a distant future date. This will
// never include expired entries at time of access.
//
// Setting this property will replace all existing entries. Use this property
// for saving and restoring permission status in bulk.
//
// Permissions in this dictionary should be explicitly granted by the user
// before being added. Any permissions in this collection will not be
// presented for approval again until they expire. This value should be saved
// and restored as needed by the app.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/grantedPermissions
func (w WKWebExtensionContext) GrantedPermissions() foundation.INSDictionary {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("grantedPermissions"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (w WKWebExtensionContext) SetGrantedPermissions(value foundation.INSDictionary) {
	objc.Send[struct{}](w.ID, objc.Sel("setGrantedPermissions:"), value)
}

// A Boolean value indicating if the currently granted permission match
// patterns set contains the “ pattern or any `*` host patterns.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasAccessToAllHosts
func (w WKWebExtensionContext) HasAccessToAllHosts() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasAccessToAllHosts"))
	return rv
}

// A Boolean value indicating if the currently granted permission match
// patterns set contains the “ pattern.
//
// # Discussion
//
// This does not check for any `*` host patterns. In most cases you should use
// the broader [HasAccessToAllHosts].
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasAccessToAllURLs
func (w WKWebExtensionContext) HasAccessToAllURLs() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasAccessToAllURLs"))
	return rv
}

// A Boolean value indicating if the extension has access to private data.
//
// # Discussion
//
// If this property is true, the extension is granted permission to interact
// with private windows, tabs, and cookies. Access to private data should be
// explicitly allowed by the user before setting this property. This value
// should be saved and restored as needed by the app.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasAccessToPrivateData
func (w WKWebExtensionContext) HasAccessToPrivateData() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasAccessToPrivateData"))
	return rv
}
func (w WKWebExtensionContext) SetHasAccessToPrivateData(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setHasAccessToPrivateData:"), value)
}

// A boolean value indicating whether the extension includes rules used for
// content modification or blocking.
//
// # Discussion
//
// This includes both static rules available in the extension’s manifest and
// dynamic rules applied during a browsing session.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasContentModificationRules
func (w WKWebExtensionContext) HasContentModificationRules() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasContentModificationRules"))
	return rv
}

// A Boolean value indicating whether the extension has script or stylesheet
// content that can be injected into webpages.
//
// # Discussion
//
// If this property is [YES], the extension has content that can be injected
// by matching against the extension’s requested match patterns.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasInjectedContent
func (w WKWebExtensionContext) HasInjectedContent() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasInjectedContent"))
	return rv
}

// A Boolean value indicating if the extension has requested optional access
// to all hosts.
//
// # Discussion
//
// If this property is [YES], the extension has asked for access to all hosts
// in a call to
// `browser.RuntimeXCUIElementTypePermissionsXCUIElementTypeRequest()`, and
// future permission checks will present discrete hosts for approval as being
// implicitly requested. This value should be saved and restored as needed by
// the app.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/hasRequestedOptionalAccessToAllHosts
func (w WKWebExtensionContext) HasRequestedOptionalAccessToAllHosts() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasRequestedOptionalAccessToAllHosts"))
	return rv
}
func (w WKWebExtensionContext) SetHasRequestedOptionalAccessToAllHosts(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setHasRequestedOptionalAccessToAllHosts:"), value)
}

// The name shown when inspecting the background web view.
//
// # Discussion
//
// This is the text that will appear when inspecting the background web view.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/inspectionName
func (w WKWebExtensionContext) InspectionName() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("inspectionName"))
	return foundation.NSStringFromID(rv).String()
}
func (w WKWebExtensionContext) SetInspectionName(value string) {
	objc.Send[struct{}](w.ID, objc.Sel("setInspectionName:"), objc.String(value))
}

// Determines whether Web Inspector can inspect the [WKWebView] instances for
// this context.
//
// # Discussion
//
// A context can control multiple [WKWebView] instances, from the background
// content, to the popover.
//
// You should set this to [YES] when needed for debugging purposes. The
// default value is [NO].
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/isInspectable
func (w WKWebExtensionContext) Inspectable() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isInspectable"))
	return rv
}
func (w WKWebExtensionContext) SetInspectable(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setInspectable:"), value)
}

// A Boolean value indicating if this context is loaded in an extension
// controller.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/isLoaded
func (w WKWebExtensionContext) Loaded() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isLoaded"))
	return rv
}

// A set of open tabs in all open windows that are exposed to this extension.
//
// # Discussion
//
// Provides a set of tabs in all open windows that are visible to the
// extension, as updated by the [DidOpenTab] and [DidCloseTabWindowIsClosing]
// methods.
//
// Initially populated by the tabs in the windows returned by the extension
// controller delegate method
// [WebExtensionControllerOpenWindowsForExtensionContext].
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/openTabs
func (w WKWebExtensionContext) OpenTabs() foundation.INSSet {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("openTabs"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The open windows that are exposed to this extension.
//
// # Discussion
//
// Provides the windows that are open and visible to the extension, as updated
// by the [DidOpenWindow] and [DidCloseWindow] methods.
//
// Initially populated by the windows returned by the extension controller
// delegate method [WebExtensionControllerOpenWindowsForExtensionContext].
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/openWindows
func (w WKWebExtensionContext) OpenWindows() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](w.ID, objc.Sel("openWindows"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// The URL of the extension’s options page, if the extension has one.
//
// # Discussion
//
// Provides the URL for the dedicated options page, if provided by the
// extension; otherwise `nil` if no page is defined.
//
// The app should provide access to this page through a user interface
// element.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/optionsPageURL
func (w WKWebExtensionContext) OptionsPageURL() foundation.INSURL {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("optionsPageURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// The URL to use as an alternative to the default new tab page, if the
// extension has one.
//
// # Discussion
//
// Provides the URL for a new tab page, if provided by the extension;
// otherwise `nil` if no page is defined.
//
// The app should prompt the user for permission to use the extension’s new
// tab page as the default.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/overrideNewTabPageURL
func (w WKWebExtensionContext) OverrideNewTabPageURL() foundation.INSURL {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("overrideNewTabPageURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// A unique identifier used to distinguish the extension from other extensions
// and target it for messages.
//
// # Discussion
//
// The default value is a unique value that matches the host in the default
// base URL. The identifier can be any value that is unique. Setting is only
// allowed when the context is not loaded. This value is accessible by the
// extension via `browser.RuntimeXCUIElementTypeId()` and is used for
// messaging the extension via `browser.RuntimeXCUIElementTypeSendMessage()`.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/uniqueIdentifier
func (w WKWebExtensionContext) UniqueIdentifier() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("uniqueIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (w WKWebExtensionContext) SetUniqueIdentifier(value string) {
	objc.Send[struct{}](w.ID, objc.Sel("setUniqueIdentifier:"), objc.String(value))
}

// Specifies unsupported APIs for this extension, making them `undefined` in
// JavaScript.
//
// # Discussion
//
// This property allows the app to specify a subset of web extension APIs that
// it chooses not to support, effectively making these APIs `undefined` within
// the extension’s JavaScript contexts. This enables extensions to employ
// feature detection techniques for unsupported APIs, allowing them to adapt
// their behavior based on the APIs actually supported by the app. Setting is
// only allowed when the context is not loaded. Only certain APIs can be
// specified here, particularly those within the `browser` namespace and other
// dynamic functions and properties, anything else will be silently ignored.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/unsupportedAPIs
func (w WKWebExtensionContext) UnsupportedAPIs() foundation.INSSet {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("unsupportedAPIs"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (w WKWebExtensionContext) SetUnsupportedAPIs(value foundation.INSSet) {
	objc.Send[struct{}](w.ID, objc.Sel("setUnsupportedAPIs:"), value)
}

// The extension this context represents.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/webExtension
func (w WKWebExtensionContext) WebExtension() IWKWebExtension {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("webExtension"))
	return WKWebExtensionFromID(objc.ID(rv))
}

// The extension controller this context is loaded in, otherwise `nil` if it
// isn’t loaded.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/webExtensionController
func (w WKWebExtensionContext) WebExtensionController() IWKWebExtensionController {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("webExtensionController"))
	return WKWebExtensionControllerFromID(objc.ID(rv))
}

// The web view configuration to use for web views that load pages from this
// extension.
//
// # Discussion
//
// Returns a customized copy of the configuration, originally set in the web
// extension controller configuration, for this extension.
//
// The app must use this configuration when initializing web views intended to
// navigate to a URL originating from this extension’s base URL.
//
// The app must also swap web views in tabs when navigating to and from web
// extension URLs. This property returns `nil` if the context isn’t
// associated with a web extension controller. The returned configuration copy
// can be customized prior to web view initialization.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/webViewConfiguration
func (w WKWebExtensionContext) WebViewConfiguration() IWKWebViewConfiguration {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("webViewConfiguration"))
	return WKWebViewConfigurationFromID(objc.ID(rv))
}

// LoadBackgroundContent is a synchronous wrapper around [WKWebExtensionContext.LoadBackgroundContentWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w WKWebExtensionContext) LoadBackgroundContent(ctx context.Context) error {
	done := make(chan error, 1)
	w.LoadBackgroundContentWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
