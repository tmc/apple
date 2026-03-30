// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol with methods that represent a tab to web extensions.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab
type WKWebExtensionTab interface {
	objectivec.IObject
}

// WKWebExtensionTabObject wraps an existing Objective-C object that conforms to the WKWebExtensionTab protocol.
type WKWebExtensionTabObject struct {
	objectivec.Object
}

func (o WKWebExtensionTabObject) BaseObject() objectivec.Object {
	return o.Object
}

// WKWebExtensionTabObjectFromID constructs a [WKWebExtensionTabObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func WKWebExtensionTabObjectFromID(id objc.ID) WKWebExtensionTabObject {
	return WKWebExtensionTabObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Called to activate the tab, making it frontmost.
//
// context: The context in which the web extension is running.
//
// completionHandler: A block that must be called upon completion. It takes a single error
// argument, which should be provided if any errors occurred.
//
// # Discussion
//
// Upon activation, the tab should become the frontmost and either be the sole
// selected tab or be included among the selected tabs. No action is performed
// if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/activate(for:completionHandler:)
func (o WKWebExtensionTabObject) ActivateForWebExtensionContextCompletionHandler(context IWKWebExtensionContext, completionHandler ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("activateForWebExtensionContext:completionHandler:"), context, completionHandler)
}

// Called to close the tab.
//
// context: The context in which the web extension is running.
//
// completionHandler: A block that must be called upon completion. It takes a single error
// argument, which should be provided if any errors occurred.
//
// # Discussion
//
// No action is performed if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/close(for:completionHandler:)
func (o WKWebExtensionTabObject) CloseForWebExtensionContextCompletionHandler(context IWKWebExtensionContext, completionHandler ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("closeForWebExtensionContext:completionHandler:"), context, completionHandler)
}

// Called to detect the locale of the webpage currently loaded in the tab.
//
// context: The context in which the web extension is running.
//
// completionHandler: A block that must be called upon completion. The block takes two arguments:
// the detected locale (or `nil` if the locale is unknown) and an error, which
// should be provided if any errors occurred.
//
// # Discussion
//
// No action is performed if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/detectWebpageLocale(for:completionHandler:)
func (o WKWebExtensionTabObject) DetectWebpageLocaleForWebExtensionContextCompletionHandler(context IWKWebExtensionContext, completionHandler LocaleErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("detectWebpageLocaleForWebExtensionContext:completionHandler:"), context, completionHandler)
}

// Called to duplicate the tab.
//
// configuration: The tab configuration influencing the duplicated tab’s properties.
//
// context: The context in which the web extension is running.
//
// completionHandler: A block that must be called upon completion. It takes two arguments: the
// duplicated tab (or \c nil if no tab was created) and an error, which should
// be provided if any errors occurred.
//
// # Discussion
//
// This is equivalent to the user selecting to duplicate the tab through a
// menu item, with the specified configuration.
//
// No action is performed if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/duplicate(using:for:completionHandler:)
func (o WKWebExtensionTabObject) DuplicateUsingConfigurationForWebExtensionContextCompletionHandler(configuration IWKWebExtensionTabConfiguration, context IWKWebExtensionContext, completionHandler WKWebExtensionTabErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("duplicateUsingConfiguration:forWebExtensionContext:completionHandler:"), configuration, context, completionHandler)
}

// Called to navigate the tab to the previous page in its history.
//
// context: The context in which the web extension is running.
//
// completionHandler: A block that must be called upon completion. It takes a single error
// argument, which should be provided if any errors occurred.
//
// # Discussion
//
// Navigates to the previous page in the tab’s web view via [GoBack] if not
// implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/goBack(for:completionHandler:)
func (o WKWebExtensionTabObject) GoBackForWebExtensionContextCompletionHandler(context IWKWebExtensionContext, completionHandler ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("goBackForWebExtensionContext:completionHandler:"), context, completionHandler)
}

// Called to navigate the tab to the next page in its history.
//
// context: The context in which the web extension is running.
//
// completionHandler: A block that must be called upon completion. It takes a single error
// argument, which should be provided if any errors occurred.
//
// # Discussion
//
// Navigates to the next page in the tab’s web view via [GoForward] if not
// implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/goForward(for:completionHandler:)
func (o WKWebExtensionTabObject) GoForwardForWebExtensionContextCompletionHandler(context IWKWebExtensionContext, completionHandler ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("goForwardForWebExtensionContext:completionHandler:"), context, completionHandler)
}

// Called when the index of the tab in the window is needed.
//
// context: The context in which the web extension is running.
//
// # Discussion
//
// This method should be implemented for better performance. Defaults to the
// window’s [TabsForWebExtensionContext] method to find the index if not
// implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/indexInWindow(for:)
func (o WKWebExtensionTabObject) IndexInWindowForWebExtensionContext(context IWKWebExtensionContext) uint {
	rv := objc.Send[uint](o.ID, objc.Sel("indexInWindowForWebExtensionContext:"), context)
	return rv
}

// Called to check if the tab has finished loading.
//
// context: The context in which the web extension is running.
//
// # Discussion
//
// Defaults to [Loading] of the tab’s web view if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/isLoadingComplete(for:)
func (o WKWebExtensionTabObject) IsLoadingCompleteForWebExtensionContext(context IWKWebExtensionContext) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isLoadingCompleteForWebExtensionContext:"), context)
	return rv
}

// Called to check if the tab is currently muted.
//
// context: The context in which the web extension is running.
//
// # Discussion
//
// Defaults to [NO] if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/isMuted(for:)
func (o WKWebExtensionTabObject) IsMutedForWebExtensionContext(context IWKWebExtensionContext) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isMutedForWebExtensionContext:"), context)
	return rv
}

// Called when the pinned state of the tab is needed.
//
// context: The context in which the web extension is running.
//
// # Discussion
//
// Defaults to [NO] if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/isPinned(for:)
func (o WKWebExtensionTabObject) IsPinnedForWebExtensionContext(context IWKWebExtensionContext) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isPinnedForWebExtensionContext:"), context)
	return rv
}

// Called to check if the tab is currently playing audio.
//
// context: The context in which the web extension is running.
//
// # Discussion
//
// Defaults to [NO] if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/isPlayingAudio(for:)
func (o WKWebExtensionTabObject) IsPlayingAudioForWebExtensionContext(context IWKWebExtensionContext) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isPlayingAudioForWebExtensionContext:"), context)
	return rv
}

// Called to check if the tab is currently showing reader mode.
//
// context: The context in which the web extension is running.
//
// # Discussion
//
// Defaults to [NO] if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/isReaderModeActive(for:)
func (o WKWebExtensionTabObject) IsReaderModeActiveForWebExtensionContext(context IWKWebExtensionContext) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isReaderModeActiveForWebExtensionContext:"), context)
	return rv
}

// Called to check if reader mode is available for the tab.
//
// context: The context in which the web extension is running.
//
// # Discussion
//
// Defaults to [NO] if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/isReaderModeAvailable(for:)
func (o WKWebExtensionTabObject) IsReaderModeAvailableForWebExtensionContext(context IWKWebExtensionContext) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isReaderModeAvailableForWebExtensionContext:"), context)
	return rv
}

// Called when the selected state of the tab is needed.
//
// context: The context in which the web extension is running.
//
// # Discussion
//
// Defaults to [YES] for the active tab and [NO] for other tabs if not
// implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/isSelected(for:)
func (o WKWebExtensionTabObject) IsSelectedForWebExtensionContext(context IWKWebExtensionContext) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isSelectedForWebExtensionContext:"), context)
	return rv
}

// Called to load a URL in the tab.
//
// url: The URL to be loaded in the tab.
//
// context: The context in which the web extension is running.
//
// completionHandler: A block that must be called upon completion. It takes a single error
// argument, which should be provided if any errors occurred.
//
// # Discussion
//
// If the tab is already loading a page, calling this method should stop the
// current page from loading and start loading the new URL. Loads the URL in
// the tab’s web view via [LoadRequest] if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/loadURL(_:for:completionHandler:)
func (o WKWebExtensionTabObject) LoadURLForWebExtensionContextCompletionHandler(url foundation.INSURL, context IWKWebExtensionContext, completionHandler ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("loadURL:forWebExtensionContext:completionHandler:"), url, context, completionHandler)
}

// Called when the parent tab for the tab is needed.
//
// context: The context in which the web extension is running.
//
// # Discussion
//
// Defaults to `nil` if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/parentTab(for:)
func (o WKWebExtensionTabObject) ParentTabForWebExtensionContext(context IWKWebExtensionContext) WKWebExtensionTab {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("parentTabForWebExtensionContext:"), context)
	return WKWebExtensionTabObjectFromID(rv)
}

// Called when the pending URL of the tab is needed.
//
// context: The context in which the web extension is running.
//
// # Discussion
//
// The pending URL is the URL of a page that is in the process of loading. If
// there is no pending URL, return `nil`.
//
// Defaults to `nil` if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/pendingURL(for:)
func (o WKWebExtensionTabObject) PendingURLForWebExtensionContext(context IWKWebExtensionContext) foundation.INSURL {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("pendingURLForWebExtensionContext:"), context)
	return foundation.NSURLFromID(rv)
}

// Called to reload the current page in the tab.
//
// fromOrigin: A boolean value indicating whether to reload the tab from the origin,
// bypassing the cache.
//
// context: The context in which the web extension is running.
//
// completionHandler: A block that must be called upon completion. It takes a single error
// argument, which should be provided if any errors occurred.
//
// # Discussion
//
// Reloads the tab’s web view via [Reload] or [ReloadFromOrigin] if not
// implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/reload(fromOrigin:for:completionHandler:)
func (o WKWebExtensionTabObject) ReloadFromOriginForWebExtensionContextCompletionHandler(fromOrigin bool, context IWKWebExtensionContext, completionHandler ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("reloadFromOrigin:forWebExtensionContext:completionHandler:"), fromOrigin, context, completionHandler)
}

// Called to set the mute state of the tab.
//
// muted: A boolean indicating whether the tab should be muted.
//
// context: The context in which the web extension is running.
//
// completionHandler: A block that must be called upon completion. It takes a single error
// argument, which should be provided if any errors occurred.
//
// # Discussion
//
// No action is performed if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/setMuted(_:for:completionHandler:)
func (o WKWebExtensionTabObject) SetMutedForWebExtensionContextCompletionHandler(muted bool, context IWKWebExtensionContext, completionHandler ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("setMuted:forWebExtensionContext:completionHandler:"), muted, context, completionHandler)
}

// Called to set or clear the parent tab for the tab.
//
// parentTab: The tab that should be set as the parent of the tab. If \c nil is provided,
// the current parent tab should be cleared.
//
// context: The context in which the web extension is running.
//
// completionHandler: A block that must be called upon completion. It takes a single error
// argument, which should be provided if any errors occurred.
//
// # Discussion
//
// No action is performed if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/setParentTab(_:for:completionHandler:)
func (o WKWebExtensionTabObject) SetParentTabForWebExtensionContextCompletionHandler(parentTab WKWebExtensionTab, context IWKWebExtensionContext, completionHandler ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("setParentTab:forWebExtensionContext:completionHandler:"), parentTab, context, completionHandler)
}

// Called to set the pinned state of the tab.
//
// pinned: A boolean value indicating whether to pin the tab.
//
// context: The context in which the web extension is running.
//
// completionHandler: A block that must be called upon completion. It takes a single error
// argument, which should be provided if any errors occurred.
//
// # Discussion
//
// This is equivalent to the user selecting to pin or unpin the tab through a
// menu item. When a tab is pinned, it should be moved to the front of the tab
// bar and usually reduced in size. When a tab is unpinned, it should be
// restored to a normal size and position in the tab bar. No action is
// performed if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/setPinned(_:for:completionHandler:)
func (o WKWebExtensionTabObject) SetPinnedForWebExtensionContextCompletionHandler(pinned bool, context IWKWebExtensionContext, completionHandler ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("setPinned:forWebExtensionContext:completionHandler:"), pinned, context, completionHandler)
}

// Called to set the reader mode for the tab.
//
// active: A boolean value indicating whether to activate reader mode.
//
// context: The context in which the web extension is running.
//
// completionHandler: A block that must be called upon completion. It takes a single error
// argument, which should be provided if any errors occurred.
//
// # Discussion
//
// No action is performed if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/setReaderModeActive(_:for:completionHandler:)
func (o WKWebExtensionTabObject) SetReaderModeActiveForWebExtensionContextCompletionHandler(active bool, context IWKWebExtensionContext, completionHandler ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("setReaderModeActive:forWebExtensionContext:completionHandler:"), active, context, completionHandler)
}

// Called to set the selected state of the tab.
//
// selected: A boolean value indicating whether to select the tab.
//
// context: The context in which the web extension is running.
//
// completionHandler: A block that must be called upon completion. It takes a single error
// argument, which should be provided if any errors occurred.
//
// # Discussion
//
// This is equivalent to the user command-clicking on the tab to add it to or
// remove it from a selection.
//
// The method should update the tab’s selection state without changing the
// active tab. No action is performed if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/setSelected(_:for:completionHandler:)
func (o WKWebExtensionTabObject) SetSelectedForWebExtensionContextCompletionHandler(selected bool, context IWKWebExtensionContext, completionHandler ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("setSelected:forWebExtensionContext:completionHandler:"), selected, context, completionHandler)
}

// Called to set the zoom factor of the tab.
//
// zoomFactor: The desired zoom factor for the tab.
//
// context: The context in which the web extension is running.
//
// completionHandler: A block that must be called upon completion. It takes a single error
// argument, which should be provided if any errors occurred.
//
// # Discussion
//
// Sets [PageZoom] of the tab’s web view if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/setZoomFactor(_:for:completionHandler:)
func (o WKWebExtensionTabObject) SetZoomFactorForWebExtensionContextCompletionHandler(zoomFactor float64, context IWKWebExtensionContext, completionHandler ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("setZoomFactor:forWebExtensionContext:completionHandler:"), zoomFactor, context, completionHandler)
}

// Called to determine if the tab should bypass host permission checks.
//
// context: The context in which the web extension is running.
//
// # Discussion
//
// This method allows the app to dynamically control whether a tab can bypass
// standard host permission checks.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/shouldBypassPermissions(for:)
func (o WKWebExtensionTabObject) ShouldBypassPermissionsForWebExtensionContext(context IWKWebExtensionContext) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("shouldBypassPermissionsForWebExtensionContext:"), context)
	return rv
}

// Called to determine if permissions should be granted for the tab on user
// gesture.
//
// context: The context in which the web extension is running.
//
// # Discussion
//
// This method allows the app to control granting of permissions on a per-tab
// basis when triggered by a user gesture. Implementing this method enables
// the app to dynamically manage `activeTab` permissions based on the tab’s
// current state, the content being accessed, or other custom criteria.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/shouldGrantPermissionsOnUserGesture(for:)
func (o WKWebExtensionTabObject) ShouldGrantPermissionsOnUserGestureForWebExtensionContext(context IWKWebExtensionContext) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("shouldGrantPermissionsOnUserGestureForWebExtensionContext:"), context)
	return rv
}

// Called when the size of the tab is needed.
//
// context: The context in which the web extension is running.
//
// # Discussion
//
// Defaults to size of the tab’s web view if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/size(for:)
func (o WKWebExtensionTabObject) SizeForWebExtensionContext(context IWKWebExtensionContext) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("sizeForWebExtensionContext:"), context)
	return rv
}

// Called when the title of the tab is needed.
//
// context: The context in which the web extension is running.
//
// # Discussion
//
// Defaults to [Title] of the tab’s web view if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/title(for:)
func (o WKWebExtensionTabObject) TitleForWebExtensionContext(context IWKWebExtensionContext) string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("titleForWebExtensionContext:"), context)
	return foundation.NSStringFromID(rv).String()
}

// Called when the URL of the tab is needed.
//
// context: The context in which the web extension is running.
//
// # Discussion
//
// Defaults to [URL] of the tab’s web view if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/url(for:)
func (o WKWebExtensionTabObject) UrlForWebExtensionContext(context IWKWebExtensionContext) foundation.INSURL {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("urlForWebExtensionContext:"), context)
	return foundation.NSURLFromID(rv)
}

// Called when the web view for the tab is needed.
//
// context: The context in which the web extension is running.
//
// # Discussion
//
// The web view’s [WKWebViewConfiguration] must have its
// [WebExtensionController] property set to match the controller of the given
// context; otherwise `nil` will be used. Defaults to `nil` if not
// implemented. If `nil`, some critical features will not be available for
// this tab, such as content injection or modification.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/webView(for:)
func (o WKWebExtensionTabObject) WebViewForWebExtensionContext(context IWKWebExtensionContext) IWKWebView {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("webViewForWebExtensionContext:"), context)
	return WKWebViewFromID(rv)
}

// Called when the window containing the tab is needed.
//
// context: The context in which the web extension is running.
//
// # Discussion
//
// Defaults to `nil` if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/window(for:)
func (o WKWebExtensionTabObject) WindowForWebExtensionContext(context IWKWebExtensionContext) WKWebExtensionWindow {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("windowForWebExtensionContext:"), context)
	return WKWebExtensionWindowObjectFromID(rv)
}

// Called when the zoom factor of the tab is needed.
//
// context: The context in which the web extension is running.
//
// # Discussion
//
// Defaults to [PageZoom] of the tab’s web view if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionTab/zoomFactor(for:)
func (o WKWebExtensionTabObject) ZoomFactorForWebExtensionContext(context IWKWebExtensionContext) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("zoomFactorForWebExtensionContext:"), context)
	return rv
}
