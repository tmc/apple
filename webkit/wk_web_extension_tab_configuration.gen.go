// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKWebExtensionTabConfiguration] class.
var (
	_WKWebExtensionTabConfigurationClass     WKWebExtensionTabConfigurationClass
	_WKWebExtensionTabConfigurationClassOnce sync.Once
)

func getWKWebExtensionTabConfigurationClass() WKWebExtensionTabConfigurationClass {
	_WKWebExtensionTabConfigurationClassOnce.Do(func() {
		_WKWebExtensionTabConfigurationClass = WKWebExtensionTabConfigurationClass{class: objc.GetClass("WKWebExtensionTabConfiguration")}
	})
	return _WKWebExtensionTabConfigurationClass
}

// GetWKWebExtensionTabConfigurationClass returns the class object for WKWebExtensionTabConfiguration.
func GetWKWebExtensionTabConfigurationClass() WKWebExtensionTabConfigurationClass {
	return getWKWebExtensionTabConfigurationClass()
}

type WKWebExtensionTabConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKWebExtensionTabConfigurationClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKWebExtensionTabConfigurationClass) Alloc() WKWebExtensionTabConfiguration {
	rv := objc.Send[WKWebExtensionTabConfiguration](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that encapsulates configuration options for a tab in an
// extension.
//
// # Overview
//
// This class holds various options that influence the behavior and initial
// state of a tab.
//
// The app retains the discretion to disregard any or all of these options, or
// even opt not to create a tab.
//
// # Instance Properties
//
//   - [WKWebExtensionTabConfiguration.Index]: Indicates the position where the tab should be opened within the window.
//   - [WKWebExtensionTabConfiguration.SetIndex]
//   - [WKWebExtensionTabConfiguration.ParentTab]: Indicates the parent tab with which the tab should be related.
//   - [WKWebExtensionTabConfiguration.ShouldAddToSelection]: Indicates whether the tab should be added to the current tab selection.
//   - [WKWebExtensionTabConfiguration.ShouldBeActive]: Indicates whether the tab should be the active tab.
//   - [WKWebExtensionTabConfiguration.ShouldBeMuted]: Indicates whether the tab should be muted.
//   - [WKWebExtensionTabConfiguration.ShouldBePinned]: Indicates whether the tab should be pinned.
//   - [WKWebExtensionTabConfiguration.ShouldReaderModeBeActive]: Indicates whether reader mode in the tab should be active.
//   - [WKWebExtensionTabConfiguration.Url]: Indicates the initial URL for the tab.
//   - [WKWebExtensionTabConfiguration.Window]: Indicates the window where the tab should be opened.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabConfiguration
type WKWebExtensionTabConfiguration struct {
	objectivec.Object
}

// WKWebExtensionTabConfigurationFromID constructs a [WKWebExtensionTabConfiguration] from an objc.ID.
//
// An object that encapsulates configuration options for a tab in an
// extension.
func WKWebExtensionTabConfigurationFromID(id objc.ID) WKWebExtensionTabConfiguration {
	return WKWebExtensionTabConfiguration{objectivec.Object{ID: id}}
}

// NOTE: WKWebExtensionTabConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKWebExtensionTabConfiguration] class.
//
// # Instance Properties
//
//   - [IWKWebExtensionTabConfiguration.Index]: Indicates the position where the tab should be opened within the window.
//   - [IWKWebExtensionTabConfiguration.SetIndex]
//   - [IWKWebExtensionTabConfiguration.ParentTab]: Indicates the parent tab with which the tab should be related.
//   - [IWKWebExtensionTabConfiguration.ShouldAddToSelection]: Indicates whether the tab should be added to the current tab selection.
//   - [IWKWebExtensionTabConfiguration.ShouldBeActive]: Indicates whether the tab should be the active tab.
//   - [IWKWebExtensionTabConfiguration.ShouldBeMuted]: Indicates whether the tab should be muted.
//   - [IWKWebExtensionTabConfiguration.ShouldBePinned]: Indicates whether the tab should be pinned.
//   - [IWKWebExtensionTabConfiguration.ShouldReaderModeBeActive]: Indicates whether reader mode in the tab should be active.
//   - [IWKWebExtensionTabConfiguration.Url]: Indicates the initial URL for the tab.
//   - [IWKWebExtensionTabConfiguration.Window]: Indicates the window where the tab should be opened.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabConfiguration
type IWKWebExtensionTabConfiguration interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Indicates the position where the tab should be opened within the window.
	Index() int
	SetIndex(value int)
	// Indicates the parent tab with which the tab should be related.
	ParentTab() WKWebExtensionTab
	// Indicates whether the tab should be added to the current tab selection.
	ShouldAddToSelection() bool
	// Indicates whether the tab should be the active tab.
	ShouldBeActive() bool
	// Indicates whether the tab should be muted.
	ShouldBeMuted() bool
	// Indicates whether the tab should be pinned.
	ShouldBePinned() bool
	// Indicates whether reader mode in the tab should be active.
	ShouldReaderModeBeActive() bool
	// Indicates the initial URL for the tab.
	Url() foundation.INSURL
	// Indicates the window where the tab should be opened.
	Window() WKWebExtensionWindow
}

// Init initializes the instance.
func (w WKWebExtensionTabConfiguration) Init() WKWebExtensionTabConfiguration {
	rv := objc.Send[WKWebExtensionTabConfiguration](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WKWebExtensionTabConfiguration) Autorelease() WKWebExtensionTabConfiguration {
	rv := objc.Send[WKWebExtensionTabConfiguration](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKWebExtensionTabConfiguration creates a new WKWebExtensionTabConfiguration instance.
func NewWKWebExtensionTabConfiguration() WKWebExtensionTabConfiguration {
	class := getWKWebExtensionTabConfigurationClass()
	rv := objc.Send[WKWebExtensionTabConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Indicates the position where the tab should be opened within the window.
//
// See: https://developer.apple.com/documentation/webkit/wkwebextension/tabconfiguration/index
func (w WKWebExtensionTabConfiguration) Index() int {
	rv := objc.Send[int](w.ID, objc.Sel("index"))
	return rv
}
func (w WKWebExtensionTabConfiguration) SetIndex(value int) {
	objc.Send[struct{}](w.ID, objc.Sel("setIndex:"), value)
}

// Indicates the parent tab with which the tab should be related.
//
// # Discussion
//
// If this property is `nil`, no parent tab was specified.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabConfiguration/parentTab
func (w WKWebExtensionTabConfiguration) ParentTab() WKWebExtensionTab {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("parentTab"))
	return WKWebExtensionTabObjectFromID(rv)
}

// Indicates whether the tab should be added to the current tab selection.
//
// # Discussion
//
// If this property is [YES], the tab should be part of the current selection,
// but not necessarily become the active tab unless [ShouldBeActive] is also
// [YES]. If this property is [NO], the tab shouldn’t be part of the current
// selection.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabConfiguration/shouldAddToSelection
func (w WKWebExtensionTabConfiguration) ShouldAddToSelection() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("shouldAddToSelection"))
	return rv
}

// Indicates whether the tab should be the active tab.
//
// # Discussion
//
// If this property is [YES], the tab should be made active in the window,
// ensuring it is the frontmost tab. Being active implies the tab is also
// selected. If this property is [NO], the tab shouldn’t affect the
// currently active tab.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabConfiguration/shouldBeActive
func (w WKWebExtensionTabConfiguration) ShouldBeActive() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("shouldBeActive"))
	return rv
}

// Indicates whether the tab should be muted.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabConfiguration/shouldBeMuted
func (w WKWebExtensionTabConfiguration) ShouldBeMuted() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("shouldBeMuted"))
	return rv
}

// Indicates whether the tab should be pinned.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabConfiguration/shouldBePinned
func (w WKWebExtensionTabConfiguration) ShouldBePinned() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("shouldBePinned"))
	return rv
}

// Indicates whether reader mode in the tab should be active.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabConfiguration/shouldReaderModeBeActive
func (w WKWebExtensionTabConfiguration) ShouldReaderModeBeActive() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("shouldReaderModeBeActive"))
	return rv
}

// Indicates the initial URL for the tab.
//
// # Discussion
//
// If this property is `nil`, the app’s default start page should appear in
// the tab.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabConfiguration/url
func (w WKWebExtensionTabConfiguration) Url() foundation.INSURL {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// Indicates the window where the tab should be opened.
//
// # Discussion
//
// If this property is `nil`, no window was specified.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabConfiguration/window
func (w WKWebExtensionTabConfiguration) Window() WKWebExtensionWindow {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("window"))
	return WKWebExtensionWindowObjectFromID(rv)
}
