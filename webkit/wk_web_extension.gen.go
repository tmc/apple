// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"context"
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKWebExtension] class.
var (
	_WKWebExtensionClass     WKWebExtensionClass
	_WKWebExtensionClassOnce sync.Once
)

func getWKWebExtensionClass() WKWebExtensionClass {
	_WKWebExtensionClassOnce.Do(func() {
		_WKWebExtensionClass = WKWebExtensionClass{class: objc.GetClass("WKWebExtension")}
	})
	return _WKWebExtensionClass
}

// GetWKWebExtensionClass returns the class object for WKWebExtension.
func GetWKWebExtensionClass() WKWebExtensionClass {
	return getWKWebExtensionClass()
}

type WKWebExtensionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKWebExtensionClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKWebExtensionClass) Alloc() WKWebExtension {
	rv := objc.Send[WKWebExtension](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that encapsulates a web extension’s resources that the manifest
// file defines.
//
// # Overview
//
// This class reads and parses the `manifest.Json()` file along with the
// supporting resources like icons and localizations.
//
// # Instance Properties
//
//   - [WKWebExtension.AllRequestedMatchPatterns]: The set of websites that the extension requires access to for injected content and for receiving messages from websites.
//   - [WKWebExtension.DefaultLocale]: The default locale for the extension.
//   - [WKWebExtension.DisplayActionLabel]: The default localized extension action label.
//   - [WKWebExtension.DisplayDescription]: The localized extension description.
//   - [WKWebExtension.DisplayName]: The localized extension name.
//   - [WKWebExtension.DisplayShortName]: The localized extension short name.
//   - [WKWebExtension.DisplayVersion]: The localized extension display version.
//   - [WKWebExtension.Errors]: An array of all errors that occurred during the processing of the extension.
//   - [WKWebExtension.HasBackgroundContent]: A Boolean value indicating whether the extension has background content that can run when needed.
//   - [WKWebExtension.HasCommands]: A Boolean value indicating whether the extension includes commands that users can invoke.
//   - [WKWebExtension.HasContentModificationRules]: A Boolean value indicating whether the extension includes rules used for content modification or blocking.
//   - [WKWebExtension.HasInjectedContent]: A Boolean value indicating whether the extension has script or stylesheet content that can be injected into webpages.
//   - [WKWebExtension.HasOptionsPage]: A Boolean value indicating whether the extension has an options page.
//   - [WKWebExtension.HasOverrideNewTabPage]: A Boolean value indicating whether the extension provides an alternative to the default new tab page.
//   - [WKWebExtension.HasPersistentBackgroundContent]: A Boolean value indicating whether the extension has background content that stays in memory as long as the extension is loaded.
//   - [WKWebExtension.Manifest]: The parsed manifest as a dictionary.
//   - [WKWebExtension.ManifestVersion]: The parsed manifest version, or `0` if there is no version specified in the manifest.
//   - [WKWebExtension.OptionalPermissionMatchPatterns]: The set of websites that the extension may need access to for optional functionality.
//   - [WKWebExtension.OptionalPermissions]: The set of permissions that the extension may need for optional functionality.
//   - [WKWebExtension.RequestedPermissionMatchPatterns]: The set of websites that the extension requires access to for its base functionality.
//   - [WKWebExtension.RequestedPermissions]: The set of permissions that the extension requires for its base functionality.
//   - [WKWebExtension.Version]: The extension version.
//
// # Instance Methods
//
//   - [WKWebExtension.ActionIconForSize]: Returns the default action icon for the specified size.
//   - [WKWebExtension.IconForSize]: Returns the extension’s icon image for the specified size.
//   - [WKWebExtension.SupportsManifestVersion]: Checks if a manifest version is supported by the extension.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension
type WKWebExtension struct {
	objectivec.Object
}

// WKWebExtensionFromID constructs a [WKWebExtension] from an objc.ID.
//
// An object that encapsulates a web extension’s resources that the manifest
// file defines.
func WKWebExtensionFromID(id objc.ID) WKWebExtension {
	return WKWebExtension{objectivec.Object{ID: id}}
}

// NOTE: WKWebExtension adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKWebExtension] class.
//
// # Instance Properties
//
//   - [IWKWebExtension.AllRequestedMatchPatterns]: The set of websites that the extension requires access to for injected content and for receiving messages from websites.
//   - [IWKWebExtension.DefaultLocale]: The default locale for the extension.
//   - [IWKWebExtension.DisplayActionLabel]: The default localized extension action label.
//   - [IWKWebExtension.DisplayDescription]: The localized extension description.
//   - [IWKWebExtension.DisplayName]: The localized extension name.
//   - [IWKWebExtension.DisplayShortName]: The localized extension short name.
//   - [IWKWebExtension.DisplayVersion]: The localized extension display version.
//   - [IWKWebExtension.Errors]: An array of all errors that occurred during the processing of the extension.
//   - [IWKWebExtension.HasBackgroundContent]: A Boolean value indicating whether the extension has background content that can run when needed.
//   - [IWKWebExtension.HasCommands]: A Boolean value indicating whether the extension includes commands that users can invoke.
//   - [IWKWebExtension.HasContentModificationRules]: A Boolean value indicating whether the extension includes rules used for content modification or blocking.
//   - [IWKWebExtension.HasInjectedContent]: A Boolean value indicating whether the extension has script or stylesheet content that can be injected into webpages.
//   - [IWKWebExtension.HasOptionsPage]: A Boolean value indicating whether the extension has an options page.
//   - [IWKWebExtension.HasOverrideNewTabPage]: A Boolean value indicating whether the extension provides an alternative to the default new tab page.
//   - [IWKWebExtension.HasPersistentBackgroundContent]: A Boolean value indicating whether the extension has background content that stays in memory as long as the extension is loaded.
//   - [IWKWebExtension.Manifest]: The parsed manifest as a dictionary.
//   - [IWKWebExtension.ManifestVersion]: The parsed manifest version, or `0` if there is no version specified in the manifest.
//   - [IWKWebExtension.OptionalPermissionMatchPatterns]: The set of websites that the extension may need access to for optional functionality.
//   - [IWKWebExtension.OptionalPermissions]: The set of permissions that the extension may need for optional functionality.
//   - [IWKWebExtension.RequestedPermissionMatchPatterns]: The set of websites that the extension requires access to for its base functionality.
//   - [IWKWebExtension.RequestedPermissions]: The set of permissions that the extension requires for its base functionality.
//   - [IWKWebExtension.Version]: The extension version.
//
// # Instance Methods
//
//   - [IWKWebExtension.ActionIconForSize]: Returns the default action icon for the specified size.
//   - [IWKWebExtension.IconForSize]: Returns the extension’s icon image for the specified size.
//   - [IWKWebExtension.SupportsManifestVersion]: Checks if a manifest version is supported by the extension.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension
type IWKWebExtension interface {
	objectivec.IObject

	// Topic: Instance Properties

	// The set of websites that the extension requires access to for injected content and for receiving messages from websites.
	AllRequestedMatchPatterns() foundation.INSSet
	// The default locale for the extension.
	DefaultLocale() foundation.NSLocale
	// The default localized extension action label.
	DisplayActionLabel() string
	// The localized extension description.
	DisplayDescription() string
	// The localized extension name.
	DisplayName() string
	// The localized extension short name.
	DisplayShortName() string
	// The localized extension display version.
	DisplayVersion() string
	// An array of all errors that occurred during the processing of the extension.
	Errors() []foundation.NSError
	// A Boolean value indicating whether the extension has background content that can run when needed.
	HasBackgroundContent() bool
	// A Boolean value indicating whether the extension includes commands that users can invoke.
	HasCommands() bool
	// A Boolean value indicating whether the extension includes rules used for content modification or blocking.
	HasContentModificationRules() bool
	// A Boolean value indicating whether the extension has script or stylesheet content that can be injected into webpages.
	HasInjectedContent() bool
	// A Boolean value indicating whether the extension has an options page.
	HasOptionsPage() bool
	// A Boolean value indicating whether the extension provides an alternative to the default new tab page.
	HasOverrideNewTabPage() bool
	// A Boolean value indicating whether the extension has background content that stays in memory as long as the extension is loaded.
	HasPersistentBackgroundContent() bool
	// The parsed manifest as a dictionary.
	Manifest() foundation.INSDictionary
	// The parsed manifest version, or `0` if there is no version specified in the manifest.
	ManifestVersion() float64
	// The set of websites that the extension may need access to for optional functionality.
	OptionalPermissionMatchPatterns() foundation.INSSet
	// The set of permissions that the extension may need for optional functionality.
	OptionalPermissions() foundation.INSSet
	// The set of websites that the extension requires access to for its base functionality.
	RequestedPermissionMatchPatterns() foundation.INSSet
	// The set of permissions that the extension requires for its base functionality.
	RequestedPermissions() foundation.INSSet
	// The extension version.
	Version() string

	// Topic: Instance Methods

	// Returns the default action icon for the specified size.
	ActionIconForSize(size corefoundation.CGSize) objc.ID
	// Returns the extension’s icon image for the specified size.
	IconForSize(size corefoundation.CGSize) objc.ID
	// Checks if a manifest version is supported by the extension.
	SupportsManifestVersion(manifestVersion float64) bool
}

// Init initializes the instance.
func (w WKWebExtension) Init() WKWebExtension {
	rv := objc.Send[WKWebExtension](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WKWebExtension) Autorelease() WKWebExtension {
	rv := objc.Send[WKWebExtension](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKWebExtension creates a new WKWebExtension instance.
func NewWKWebExtension() WKWebExtension {
	class := getWKWebExtensionClass()
	rv := objc.Send[WKWebExtension](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the default action icon for the specified size.
//
// size: The size to use when looking up the action icon.
//
// # Return Value
//
// The action icon, or `nil` if the icon was unable to be loaded.
//
// # Discussion
//
// This icon serves as a default and should be used to represent the extension
// in contexts like action sheets or toolbars prior to the extension being
// loaded into an extension context. Once the extension is loaded, use the
// [ActionForTab] API to get the tab-specific icon.
//
// The returned image will be the best match for the specified size that is
// available in the extension’s action icon set. If no matching icon is
// available, the method will fall back to the extension’s icon.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/actionIcon(for:)
func (w WKWebExtension) ActionIconForSize(size corefoundation.CGSize) objc.ID {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("actionIconForSize:"), size)
	return rv
}

// Returns the extension’s icon image for the specified size.
//
// size: The size to use when looking up the icon.
//
// # Return Value
//
// The extension’s icon image, or `nil` if the icon was unable to be loaded.
//
// # Discussion
//
// This icon should represent the extension in settings or other areas that
// show the extension. The returned image will be the best match for the
// specified size that is available in the extension’s icon set. If no
// matching icon can be found, the method will return `nil`.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/icon(for:)
func (w WKWebExtension) IconForSize(size corefoundation.CGSize) objc.ID {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("iconForSize:"), size)
	return rv
}

// Checks if a manifest version is supported by the extension.
//
// manifestVersion: The version number to check.
//
// # Return Value
//
// Returns [YES] if the extension specified a manifest version that is greater
// than or equal to `manifestVersion`.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/supportsManifestVersion(_:)
func (w WKWebExtension) SupportsManifestVersion(manifestVersion float64) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("supportsManifestVersion:"), manifestVersion)
	return rv
}

// Returns a web extension initialized with a specified app extension bundle.
//
// appExtensionBundle: The bundle to use for the new web extension.
//
// completionHandler: A block to be called with an initialized web extension, or `nil` if the
// object could not be initialized due to an error.
//
// # Discussion
//
// The app extension bundle must contain a `manifest.Json()` file in its
// resources directory. If the manifest is invalid or missing, or the bundle
// is otherwise improperly configured, an error will be returned.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/extensionWithAppExtensionBundle:completionHandler:
func (_WKWebExtensionClass WKWebExtensionClass) ExtensionWithAppExtensionBundleCompletionHandler(appExtensionBundle foundation.NSBundle, completionHandler WKWebExtensionErrorHandler) {
	_block1, _ := NewWKWebExtensionErrorBlock(completionHandler)
	objc.Send[objc.ID](objc.ID(_WKWebExtensionClass.class), objc.Sel("extensionWithAppExtensionBundle:completionHandler:"), appExtensionBundle, _block1)
}

// Returns a web extension initialized with a specified resource base URL,
// which can point to either a directory or a ZIP archive.
//
// resourceBaseURL: The file URL to use for the new web extension.
//
// completionHandler: A block to be called with an initialized web extension, or `nil` if the
// object could not be initialized due to an error.
//
// # Discussion
//
// The URL must be a file URL that points to either a directory with a
// `manifest.Json()` file or a ZIP archive containing a `manifest.Json()`
// file. If the manifest is invalid or missing, or the URL points to an
// unsupported format or invalid archive, an error will be returned.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/extensionWithResourceBaseURL:completionHandler:
func (_WKWebExtensionClass WKWebExtensionClass) ExtensionWithResourceBaseURLCompletionHandler(resourceBaseURL foundation.INSURL, completionHandler WKWebExtensionErrorHandler) {
	_block1, _ := NewWKWebExtensionErrorBlock(completionHandler)
	objc.Send[objc.ID](objc.ID(_WKWebExtensionClass.class), objc.Sel("extensionWithResourceBaseURL:completionHandler:"), resourceBaseURL, _block1)
}

// The set of websites that the extension requires access to for injected
// content and for receiving messages from websites.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/allRequestedMatchPatterns
func (w WKWebExtension) AllRequestedMatchPatterns() foundation.INSSet {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("allRequestedMatchPatterns"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The default locale for the extension.
//
// # Discussion
//
// Returns `nil` if there was no default locale specified.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/defaultLocale
func (w WKWebExtension) DefaultLocale() foundation.NSLocale {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("defaultLocale"))
	return foundation.NSLocaleFromID(objc.ID(rv))
}

// The default localized extension action label.
//
// # Discussion
//
// Returns `nil` if there was no default action label specified.
//
// This label serves as a default and should be used to represent the
// extension in contexts like action sheets or toolbars prior to the extension
// being loaded into an extension context. Once the extension is loaded, use
// the [ActionForTab] API to get the tab-specific label.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/displayActionLabel
func (w WKWebExtension) DisplayActionLabel() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("displayActionLabel"))
	return foundation.NSStringFromID(rv).String()
}

// The localized extension description.
//
// # Discussion
//
// Returns `nil` if there was no description specified.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/displayDescription
func (w WKWebExtension) DisplayDescription() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("displayDescription"))
	return foundation.NSStringFromID(rv).String()
}

// The localized extension name.
//
// # Discussion
//
// Returns `nil` if there was no name specified.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/displayName
func (w WKWebExtension) DisplayName() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("displayName"))
	return foundation.NSStringFromID(rv).String()
}

// The localized extension short name.
//
// # Discussion
//
// Returns `nil` if there was no short name specified.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/displayShortName
func (w WKWebExtension) DisplayShortName() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("displayShortName"))
	return foundation.NSStringFromID(rv).String()
}

// The localized extension display version.
//
// # Discussion
//
// Returns `nil` if there was no display version specified.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/displayVersion
func (w WKWebExtension) DisplayVersion() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("displayVersion"))
	return foundation.NSStringFromID(rv).String()
}

// An array of all errors that occurred during the processing of the
// extension.
//
// # Discussion
//
// Provides an array of all parse-time errors for the extension, with repeat
// errors consolidated into a single entry for the original occurrence only.
// If no errors occurred, an empty array is returned.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/errors
func (w WKWebExtension) Errors() []foundation.NSError {
	rv := objc.Send[[]objc.ID](w.ID, objc.Sel("errors"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSError {
		return foundation.NSErrorFromID(id)
	})
}

// A Boolean value indicating whether the extension has background content
// that can run when needed.
//
// # Discussion
//
// If this property is [YES], the extension can run in the background even
// when no webpages are open.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/hasBackgroundContent
func (w WKWebExtension) HasBackgroundContent() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasBackgroundContent"))
	return rv
}

// A Boolean value indicating whether the extension includes commands that
// users can invoke.
//
// # Discussion
//
// If this property is [YES], the extension contains one or more commands that
// can be performed by the user. These commands should be accessible via
// keyboard shortcuts, menu items, or other user interface elements provided
// by the app. The list of commands can be accessed via [Commands] on an
// extension context, and invoked via [PerformCommand].
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/hasCommands
func (w WKWebExtension) HasCommands() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasCommands"))
	return rv
}

// A Boolean value indicating whether the extension includes rules used for
// content modification or blocking.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/hasContentModificationRules
func (w WKWebExtension) HasContentModificationRules() bool {
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
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/hasInjectedContent
func (w WKWebExtension) HasInjectedContent() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasInjectedContent"))
	return rv
}

// A Boolean value indicating whether the extension has an options page.
//
// # Discussion
//
// If this property is [YES], the extension includes a dedicated options page
// where users can customize settings. The app should provide access to this
// page through a user interface element, which can be accessed via
// [OptionsPageURL] on an extension context.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/hasOptionsPage
func (w WKWebExtension) HasOptionsPage() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasOptionsPage"))
	return rv
}

// A Boolean value indicating whether the extension provides an alternative to
// the default new tab page.
//
// # Discussion
//
// If this property is [YES], the extension can specify a custom page that can
// be displayed when a new tab is opened in the app, instead of the default
// new tab page. The app should prompt the user for permission to use the
// extension’s new tab page as the default, which can be accessed via
// [OverrideNewTabPageURL] on an extension context.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/hasOverrideNewTabPage
func (w WKWebExtension) HasOverrideNewTabPage() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasOverrideNewTabPage"))
	return rv
}

// A Boolean value indicating whether the extension has background content
// that stays in memory as long as the extension is loaded.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/hasPersistentBackgroundContent
func (w WKWebExtension) HasPersistentBackgroundContent() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasPersistentBackgroundContent"))
	return rv
}

// The parsed manifest as a dictionary.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/manifest
func (w WKWebExtension) Manifest() foundation.INSDictionary {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("manifest"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The parsed manifest version, or `0` if there is no version specified in the
// manifest.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/manifestVersion
func (w WKWebExtension) ManifestVersion() float64 {
	rv := objc.Send[float64](w.ID, objc.Sel("manifestVersion"))
	return rv
}

// The set of websites that the extension may need access to for optional
// functionality.
//
// # Discussion
//
// These match patterns can be requested by the extension at a later time.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/optionalPermissionMatchPatterns
func (w WKWebExtension) OptionalPermissionMatchPatterns() foundation.INSSet {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("optionalPermissionMatchPatterns"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The set of permissions that the extension may need for optional
// functionality.
//
// # Discussion
//
// These permissions can be requested by the extension at a later time.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/optionalPermissions
func (w WKWebExtension) OptionalPermissions() foundation.INSSet {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("optionalPermissions"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The set of websites that the extension requires access to for its base
// functionality.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/requestedPermissionMatchPatterns
func (w WKWebExtension) RequestedPermissionMatchPatterns() foundation.INSSet {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("requestedPermissionMatchPatterns"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The set of permissions that the extension requires for its base
// functionality.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/requestedPermissions
func (w WKWebExtension) RequestedPermissions() foundation.INSSet {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("requestedPermissions"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The extension version.
//
// # Discussion
//
// Returns `nil` if there was no version specified.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/version
func (w WKWebExtension) Version() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("version"))
	return foundation.NSStringFromID(rv).String()
}

// ExtensionWithAppExtensionBundle is a synchronous wrapper around [WKWebExtension.ExtensionWithAppExtensionBundleCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (wc WKWebExtensionClass) ExtensionWithAppExtensionBundle(ctx context.Context, appExtensionBundle foundation.NSBundle) (*WKWebExtension, error) {
	type result struct {
		val *WKWebExtension
		err error
	}
	done := make(chan result, 1)
	wc.ExtensionWithAppExtensionBundleCompletionHandler(appExtensionBundle, func(val *WKWebExtension, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// ExtensionWithResourceBaseURL is a synchronous wrapper around [WKWebExtension.ExtensionWithResourceBaseURLCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (wc WKWebExtensionClass) ExtensionWithResourceBaseURL(ctx context.Context, resourceBaseURL foundation.INSURL) (*WKWebExtension, error) {
	type result struct {
		val *WKWebExtension
		err error
	}
	done := make(chan result, 1)
	wc.ExtensionWithResourceBaseURLCompletionHandler(resourceBaseURL, func(val *WKWebExtension, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
