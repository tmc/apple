// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKWebExtensionAction] class.
var (
	_WKWebExtensionActionClass     WKWebExtensionActionClass
	_WKWebExtensionActionClassOnce sync.Once
)

func getWKWebExtensionActionClass() WKWebExtensionActionClass {
	_WKWebExtensionActionClassOnce.Do(func() {
		_WKWebExtensionActionClass = WKWebExtensionActionClass{class: objc.GetClass("WKWebExtensionAction")}
	})
	return _WKWebExtensionActionClass
}

// GetWKWebExtensionActionClass returns the class object for WKWebExtensionAction.
func GetWKWebExtensionActionClass() WKWebExtensionActionClass {
	return getWKWebExtensionActionClass()
}

type WKWebExtensionActionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKWebExtensionActionClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKWebExtensionActionClass) Alloc() WKWebExtensionAction {
	rv := objc.Send[WKWebExtensionAction](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that encapsulates the properties for an individual web extension
// action.
//
// # Overview
//
// This class provides access to action properties, such as pop-up, icon, or
// title, with tab-specific values.
//
// # Instance Properties
//
//   - [WKWebExtensionAction.AssociatedTab]: The tab that this action is associated with, or `nil` if it’s the default action.
//   - [WKWebExtensionAction.BadgeText]: The badge text for the action.
//   - [WKWebExtensionAction.HasUnreadBadgeText]: A Boolean value indicating whether the badge text is unread.
//   - [WKWebExtensionAction.SetHasUnreadBadgeText]
//   - [WKWebExtensionAction.InspectionName]: The name shown when inspecting the pop-up web view.
//   - [WKWebExtensionAction.SetInspectionName]
//   - [WKWebExtensionAction.Enabled]: A Boolean value indicating whether the action is enabled.
//   - [WKWebExtensionAction.Label]: The localized display label for the action.
//   - [WKWebExtensionAction.MenuItems]: The menu items provided by the extension for this action.
//   - [WKWebExtensionAction.PopupPopover]: A popover that presents a web view loaded with the pop-up page for this action, or `nil` if no popup is specified.
//   - [WKWebExtensionAction.PopupWebView]: A web view loaded with the pop-up page for this action, or `nil` if no pop-up is specified.
//   - [WKWebExtensionAction.PresentsPopup]: A Boolean value indicating whether the action has a pop-up.
//   - [WKWebExtensionAction.WebExtensionContext]: The extension context to which this action is related.
//
// # Instance Methods
//
//   - [WKWebExtensionAction.ClosePopup]: Triggers the dismissal process of the pop-up.
//   - [WKWebExtensionAction.IconForSize]: Returns the action icon for the specified size.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action
type WKWebExtensionAction struct {
	objectivec.Object
}

// WKWebExtensionActionFromID constructs a [WKWebExtensionAction] from an objc.ID.
//
// An object that encapsulates the properties for an individual web extension
// action.
func WKWebExtensionActionFromID(id objc.ID) WKWebExtensionAction {
	return WKWebExtensionAction{objectivec.Object{ID: id}}
}

// NOTE: WKWebExtensionAction adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKWebExtensionAction] class.
//
// # Instance Properties
//
//   - [IWKWebExtensionAction.AssociatedTab]: The tab that this action is associated with, or `nil` if it’s the default action.
//   - [IWKWebExtensionAction.BadgeText]: The badge text for the action.
//   - [IWKWebExtensionAction.HasUnreadBadgeText]: A Boolean value indicating whether the badge text is unread.
//   - [IWKWebExtensionAction.SetHasUnreadBadgeText]
//   - [IWKWebExtensionAction.InspectionName]: The name shown when inspecting the pop-up web view.
//   - [IWKWebExtensionAction.SetInspectionName]
//   - [IWKWebExtensionAction.Enabled]: A Boolean value indicating whether the action is enabled.
//   - [IWKWebExtensionAction.Label]: The localized display label for the action.
//   - [IWKWebExtensionAction.MenuItems]: The menu items provided by the extension for this action.
//   - [IWKWebExtensionAction.PopupPopover]: A popover that presents a web view loaded with the pop-up page for this action, or `nil` if no popup is specified.
//   - [IWKWebExtensionAction.PopupWebView]: A web view loaded with the pop-up page for this action, or `nil` if no pop-up is specified.
//   - [IWKWebExtensionAction.PresentsPopup]: A Boolean value indicating whether the action has a pop-up.
//   - [IWKWebExtensionAction.WebExtensionContext]: The extension context to which this action is related.
//
// # Instance Methods
//
//   - [IWKWebExtensionAction.ClosePopup]: Triggers the dismissal process of the pop-up.
//   - [IWKWebExtensionAction.IconForSize]: Returns the action icon for the specified size.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action
type IWKWebExtensionAction interface {
	objectivec.IObject

	// Topic: Instance Properties

	// The tab that this action is associated with, or `nil` if it’s the default action.
	AssociatedTab() WKWebExtensionTab
	// The badge text for the action.
	BadgeText() string
	// A Boolean value indicating whether the badge text is unread.
	HasUnreadBadgeText() bool
	SetHasUnreadBadgeText(value bool)
	// The name shown when inspecting the pop-up web view.
	InspectionName() string
	SetInspectionName(value string)
	// A Boolean value indicating whether the action is enabled.
	Enabled() bool
	// The localized display label for the action.
	Label() string
	// The menu items provided by the extension for this action.
	MenuItems() []appkit.NSMenuItem
	// A popover that presents a web view loaded with the pop-up page for this action, or `nil` if no popup is specified.
	PopupPopover() appkit.NSPopover
	// A web view loaded with the pop-up page for this action, or `nil` if no pop-up is specified.
	PopupWebView() IWKWebView
	// A Boolean value indicating whether the action has a pop-up.
	PresentsPopup() bool
	// The extension context to which this action is related.
	WebExtensionContext() IWKWebExtensionContext

	// Topic: Instance Methods

	// Triggers the dismissal process of the pop-up.
	ClosePopup()
	// Returns the action icon for the specified size.
	IconForSize(size corefoundation.CGSize) appkit.NSImage
}

// Init initializes the instance.
func (w WKWebExtensionAction) Init() WKWebExtensionAction {
	rv := objc.Send[WKWebExtensionAction](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WKWebExtensionAction) Autorelease() WKWebExtensionAction {
	rv := objc.Send[WKWebExtensionAction](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKWebExtensionAction creates a new WKWebExtensionAction instance.
func NewWKWebExtensionAction() WKWebExtensionAction {
	class := getWKWebExtensionActionClass()
	rv := objc.Send[WKWebExtensionAction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Triggers the dismissal process of the pop-up.
//
// # Discussion
//
// Invoke this method to manage the pop-up’s lifecycle, ensuring the web
// view is unloaded and resources are released once the pop-up closes. This
// method is automatically called upon the dismissal of the action’s
// [UIViewController] or [NSPopover]. For custom scenarios where the
// pop-up’s lifecycle is manually managed, it must be explicitly invoked to
// ensure proper closure.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/closePopup()
//
// [NSPopover]: https://developer.apple.com/documentation/AppKit/NSPopover
// [UIViewController]: https://developer.apple.com/documentation/UIKit/UIViewController
func (w WKWebExtensionAction) ClosePopup() {
	objc.Send[objc.ID](w.ID, objc.Sel("closePopup"))
}

// Returns the action icon for the specified size.
//
// size: The size to use when looking up the action icon.
//
// # Return Value
//
// The action icon, or `nil` if the icon was unable to be loaded.
//
// # Discussion
//
// This icon should represent the extension in action sheets or toolbars. The
// returned image will be the best match for the specified size that is
// available in the extension’s action icon set. If no matching icon is
// available, the method will fall back to the extension’s icon.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/icon(for:)
func (w WKWebExtensionAction) IconForSize(size corefoundation.CGSize) appkit.NSImage {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("iconForSize:"), size)
	return appkit.NSImageFromID(rv)
}

// The tab that this action is associated with, or `nil` if it’s the default
// action.
//
// # Discussion
//
// When this property is `nil`, it indicates that the action is the default
// action and not associated with a specific tab.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/associatedTab
func (w WKWebExtensionAction) AssociatedTab() WKWebExtensionTab {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("associatedTab"))
	return WKWebExtensionTabObjectFromID(rv)
}

// The badge text for the action.
//
// # Discussion
//
// Provides the text that appears on the badge for the action. An empty string
// signifies that no badge should be shown.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/badgeText
func (w WKWebExtensionAction) BadgeText() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("badgeText"))
	return foundation.NSStringFromID(rv).String()
}

// A Boolean value indicating whether the badge text is unread.
//
// # Discussion
//
// This property is automatically set to [YES] when [BadgeText] changes and is
// not empty. If [BadgeText] becomes empty or the popup associated with the
// action is presented, this property is automatically set to [NO].
// Additionally, it should be set to [NO] by the app when the badge has been
// presented to the user. This property is useful for higher-level
// notification badges when extensions might be hidden behind an action sheet.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/hasUnreadBadgeText
func (w WKWebExtensionAction) HasUnreadBadgeText() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasUnreadBadgeText"))
	return rv
}
func (w WKWebExtensionAction) SetHasUnreadBadgeText(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setHasUnreadBadgeText:"), value)
}

// The name shown when inspecting the pop-up web view.
//
// # Discussion
//
// This is the text that will appear when inspecting the pop-up web view.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/inspectionName
func (w WKWebExtensionAction) InspectionName() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("inspectionName"))
	return foundation.NSStringFromID(rv).String()
}
func (w WKWebExtensionAction) SetInspectionName(value string) {
	objc.Send[struct{}](w.ID, objc.Sel("setInspectionName:"), objc.String(value))
}

// A Boolean value indicating whether the action is enabled.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/isEnabled
func (w WKWebExtensionAction) Enabled() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isEnabled"))
	return rv
}

// The localized display label for the action.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/label
func (w WKWebExtensionAction) Label() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}

// The menu items provided by the extension for this action.
//
// # Discussion
//
// Provides menu items supplied by the extension, allowing the user to perform
// extension-defined actions.
//
// The app is responsible for displaying these menu items, typically in a
// context menu or a long-press menu on the action in action sheets or
// toolbars.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/menuItems
func (w WKWebExtensionAction) MenuItems() []appkit.NSMenuItem {
	rv := objc.Send[[]objc.ID](w.ID, objc.Sel("menuItems"))
	return objc.ConvertSlice(rv, func(id objc.ID) appkit.NSMenuItem {
		return appkit.NSMenuItemFromID(id)
	})
}

// A popover that presents a web view loaded with the pop-up page for this
// action, or `nil` if no popup is specified.
//
// # Discussion
//
// This popover contains a view controller with a web view preloaded with the
// pop-up page. It automatically adjusts its size to fit the web view’s
// content size. The [PresentsPopup] property should be checked to determine
// the availability of a pop-up before using this property. Dismissing the
// popover will close the pop-up and unload the web view.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/popupPopover
func (w WKWebExtensionAction) PopupPopover() appkit.NSPopover {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("popupPopover"))
	return appkit.NSPopoverFromID(objc.ID(rv))
}

// A web view loaded with the pop-up page for this action, or `nil` if no
// pop-up is specified.
//
// # Discussion
//
// The web view will be preloaded with the pop-up page upon first access or
// after it has been unloaded. Use the [PresentsPopup] property to determine
// whether a pop-up should be displayed before using this property.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/popupWebView
func (w WKWebExtensionAction) PopupWebView() IWKWebView {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("popupWebView"))
	return WKWebViewFromID(objc.ID(rv))
}

// A Boolean value indicating whether the action has a pop-up.
//
// # Discussion
//
// Use this property to check if the action has a pop-up before attempting to
// show any pop-up views.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/presentsPopup
func (w WKWebExtensionAction) PresentsPopup() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("presentsPopup"))
	return rv
}

// The extension context to which this action is related.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/webExtensionContext
func (w WKWebExtensionAction) WebExtensionContext() IWKWebExtensionContext {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("webExtensionContext"))
	return WKWebExtensionContextFromID(objc.ID(rv))
}
