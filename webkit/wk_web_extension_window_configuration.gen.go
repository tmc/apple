// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKWebExtensionWindowConfiguration] class.
var (
	_WKWebExtensionWindowConfigurationClass     WKWebExtensionWindowConfigurationClass
	_WKWebExtensionWindowConfigurationClassOnce sync.Once
)

func getWKWebExtensionWindowConfigurationClass() WKWebExtensionWindowConfigurationClass {
	_WKWebExtensionWindowConfigurationClassOnce.Do(func() {
		_WKWebExtensionWindowConfigurationClass = WKWebExtensionWindowConfigurationClass{class: objc.GetClass("WKWebExtensionWindowConfiguration")}
	})
	return _WKWebExtensionWindowConfigurationClass
}

// GetWKWebExtensionWindowConfigurationClass returns the class object for WKWebExtensionWindowConfiguration.
func GetWKWebExtensionWindowConfigurationClass() WKWebExtensionWindowConfigurationClass {
	return getWKWebExtensionWindowConfigurationClass()
}

type WKWebExtensionWindowConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKWebExtensionWindowConfigurationClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKWebExtensionWindowConfigurationClass) Alloc() WKWebExtensionWindowConfiguration {
	rv := objc.Send[WKWebExtensionWindowConfiguration](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that encapsulates configuration options for a window in an
// extension.
//
// # Overview
//
// This class holds various options that influence the behavior and initial
// state of a window.
//
// The app retains the discretion to disregard any or all of these options, or
// even opt not to create a window.
//
// # Instance Properties
//
//   - [WKWebExtensionWindowConfiguration.Frame]: Indicates the frame where the window should be positioned on the main screen.
//   - [WKWebExtensionWindowConfiguration.ShouldBeFocused]: Indicates whether the window should be focused.
//   - [WKWebExtensionWindowConfiguration.ShouldBePrivate]: Indicates whether the window should be private.
//   - [WKWebExtensionWindowConfiguration.TabURLs]: Indicates the URLs that the window should initially load as tabs.
//   - [WKWebExtensionWindowConfiguration.Tabs]: Indicates the existing tabs that should be moved to the window.
//   - [WKWebExtensionWindowConfiguration.WindowState]: Indicates the window state for the window.
//   - [WKWebExtensionWindowConfiguration.WindowType]: Indicates the window type for the window.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowConfiguration
type WKWebExtensionWindowConfiguration struct {
	objectivec.Object
}

// WKWebExtensionWindowConfigurationFromID constructs a [WKWebExtensionWindowConfiguration] from an objc.ID.
//
// An object that encapsulates configuration options for a window in an
// extension.
func WKWebExtensionWindowConfigurationFromID(id objc.ID) WKWebExtensionWindowConfiguration {
	return WKWebExtensionWindowConfiguration{objectivec.Object{ID: id}}
}

// NOTE: WKWebExtensionWindowConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKWebExtensionWindowConfiguration] class.
//
// # Instance Properties
//
//   - [IWKWebExtensionWindowConfiguration.Frame]: Indicates the frame where the window should be positioned on the main screen.
//   - [IWKWebExtensionWindowConfiguration.ShouldBeFocused]: Indicates whether the window should be focused.
//   - [IWKWebExtensionWindowConfiguration.ShouldBePrivate]: Indicates whether the window should be private.
//   - [IWKWebExtensionWindowConfiguration.TabURLs]: Indicates the URLs that the window should initially load as tabs.
//   - [IWKWebExtensionWindowConfiguration.Tabs]: Indicates the existing tabs that should be moved to the window.
//   - [IWKWebExtensionWindowConfiguration.WindowState]: Indicates the window state for the window.
//   - [IWKWebExtensionWindowConfiguration.WindowType]: Indicates the window type for the window.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowConfiguration
type IWKWebExtensionWindowConfiguration interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Indicates the frame where the window should be positioned on the main screen.
	Frame() corefoundation.CGRect
	// Indicates whether the window should be focused.
	ShouldBeFocused() bool
	// Indicates whether the window should be private.
	ShouldBePrivate() bool
	// Indicates the URLs that the window should initially load as tabs.
	TabURLs() []foundation.NSURL
	// Indicates the existing tabs that should be moved to the window.
	Tabs() []objectivec.IObject
	// Indicates the window state for the window.
	WindowState() WKWebExtensionWindowState
	// Indicates the window type for the window.
	WindowType() WKWebExtensionWindowType
}

// Init initializes the instance.
func (w WKWebExtensionWindowConfiguration) Init() WKWebExtensionWindowConfiguration {
	rv := objc.Send[WKWebExtensionWindowConfiguration](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WKWebExtensionWindowConfiguration) Autorelease() WKWebExtensionWindowConfiguration {
	rv := objc.Send[WKWebExtensionWindowConfiguration](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKWebExtensionWindowConfiguration creates a new WKWebExtensionWindowConfiguration instance.
func NewWKWebExtensionWindowConfiguration() WKWebExtensionWindowConfiguration {
	class := getWKWebExtensionWindowConfigurationClass()
	rv := objc.Send[WKWebExtensionWindowConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Indicates the frame where the window should be positioned on the main
// screen.
//
// # Discussion
//
// This frame should override the app’s default window position and size.
//
// Individual components (e.g., `origin.X()`, `size.Width()`) will be [NaN] if
// not specified.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowConfiguration/frame
func (w WKWebExtensionWindowConfiguration) Frame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](w.ID, objc.Sel("frame"))
	return corefoundation.CGRect(rv)
}

// Indicates whether the window should be focused.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowConfiguration/shouldBeFocused
func (w WKWebExtensionWindowConfiguration) ShouldBeFocused() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("shouldBeFocused"))
	return rv
}

// Indicates whether the window should be private.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowConfiguration/shouldBePrivate
func (w WKWebExtensionWindowConfiguration) ShouldBePrivate() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("shouldBePrivate"))
	return rv
}

// Indicates the URLs that the window should initially load as tabs.
//
// # Discussion
//
// If [TabURLs] and [Tabs] are both empty, the app’s default start page
// should appear in a tab.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowConfiguration/tabURLs
func (w WKWebExtensionWindowConfiguration) TabURLs() []foundation.NSURL {
	rv := objc.Send[[]objc.ID](w.ID, objc.Sel("tabURLs"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSURL {
		return foundation.NSURLFromID(id)
	})
}

// Indicates the existing tabs that should be moved to the window.
//
// # Discussion
//
// If [Tabs] and [TabURLs] are both empty, the app’s default start page
// should appear in a tab.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowConfiguration/tabs
func (w WKWebExtensionWindowConfiguration) Tabs() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](w.ID, objc.Sel("tabs"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Indicates the window state for the window.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowConfiguration/windowState
func (w WKWebExtensionWindowConfiguration) WindowState() WKWebExtensionWindowState {
	rv := objc.Send[WKWebExtensionWindowState](w.ID, objc.Sel("windowState"))
	return WKWebExtensionWindowState(rv)
}

// Indicates the window type for the window.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowConfiguration/windowType
func (w WKWebExtensionWindowConfiguration) WindowType() WKWebExtensionWindowType {
	rv := objc.Send[WKWebExtensionWindowType](w.ID, objc.Sel("windowType"))
	return WKWebExtensionWindowType(rv)
}
