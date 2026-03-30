// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKWindowFeatures] class.
var (
	_WKWindowFeaturesClass     WKWindowFeaturesClass
	_WKWindowFeaturesClassOnce sync.Once
)

func getWKWindowFeaturesClass() WKWindowFeaturesClass {
	_WKWindowFeaturesClassOnce.Do(func() {
		_WKWindowFeaturesClass = WKWindowFeaturesClass{class: objc.GetClass("WKWindowFeatures")}
	})
	return _WKWindowFeaturesClass
}

// GetWKWindowFeaturesClass returns the class object for WKWindowFeatures.
func GetWKWindowFeaturesClass() WKWindowFeaturesClass {
	return getWKWindowFeaturesClass()
}

type WKWindowFeaturesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKWindowFeaturesClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKWindowFeaturesClass) Alloc() WKWindowFeatures {
	rv := objc.Send[WKWindowFeatures](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// Display-related attributes that a webpage requests for its window.
//
// # Overview
//
// A [WKWindowFeatures] object contains the attributes that a webpage requests
// from its containing web view. You don’t create a [WKWindowFeatures]
// object directly. When a navigation action results in the display of a new
// web view, [WKWebView] creates this object and passes it to the
// [WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures]
// method of its UI delegate object. The delegate uses the information in this
// object to configure and return the new web view.
//
// # Inspecting Window Position and Dimensions
//
//   - [WKWindowFeatures.AllowsResizing]: A Boolean value that indicates whether to make the containing window window resizable.
//   - [WKWindowFeatures.Height]: The requested height of the containing window.
//   - [WKWindowFeatures.Width]: The requested width of the containing window.
//   - [WKWindowFeatures.X]: The requested x-coordinate of the containing window.
//   - [WKWindowFeatures.Y]: The requested y-coordinate of the containing window.
//
// # Inspecting Visibility Properties
//
//   - [WKWindowFeatures.MenuBarVisibility]: A Boolean value that indicates whether the webpage requests a visible menu bar.
//   - [WKWindowFeatures.StatusBarVisibility]: A Boolean value that indicates whether the webpage requested a visible status bar.
//   - [WKWindowFeatures.ToolbarsVisibility]: A Boolean value that indicates whether the webpage requested a visible toolbar.
//
// See: https://developer.apple.com/documentation/WebKit/WKWindowFeatures
type WKWindowFeatures struct {
	objectivec.Object
}

// WKWindowFeaturesFromID constructs a [WKWindowFeatures] from an objc.ID.
//
// Display-related attributes that a webpage requests for its window.
func WKWindowFeaturesFromID(id objc.ID) WKWindowFeatures {
	return WKWindowFeatures{objectivec.Object{ID: id}}
}

// NOTE: WKWindowFeatures adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKWindowFeatures] class.
//
// # Inspecting Window Position and Dimensions
//
//   - [IWKWindowFeatures.AllowsResizing]: A Boolean value that indicates whether to make the containing window window resizable.
//   - [IWKWindowFeatures.Height]: The requested height of the containing window.
//   - [IWKWindowFeatures.Width]: The requested width of the containing window.
//   - [IWKWindowFeatures.X]: The requested x-coordinate of the containing window.
//   - [IWKWindowFeatures.Y]: The requested y-coordinate of the containing window.
//
// # Inspecting Visibility Properties
//
//   - [IWKWindowFeatures.MenuBarVisibility]: A Boolean value that indicates whether the webpage requests a visible menu bar.
//   - [IWKWindowFeatures.StatusBarVisibility]: A Boolean value that indicates whether the webpage requested a visible status bar.
//   - [IWKWindowFeatures.ToolbarsVisibility]: A Boolean value that indicates whether the webpage requested a visible toolbar.
//
// See: https://developer.apple.com/documentation/WebKit/WKWindowFeatures
type IWKWindowFeatures interface {
	objectivec.IObject

	// Topic: Inspecting Window Position and Dimensions

	// A Boolean value that indicates whether to make the containing window window resizable.
	AllowsResizing() foundation.NSNumber
	// The requested height of the containing window.
	Height() foundation.NSNumber
	// The requested width of the containing window.
	Width() foundation.NSNumber
	// The requested x-coordinate of the containing window.
	X() foundation.NSNumber
	// The requested y-coordinate of the containing window.
	Y() foundation.NSNumber

	// Topic: Inspecting Visibility Properties

	// A Boolean value that indicates whether the webpage requests a visible menu bar.
	MenuBarVisibility() foundation.NSNumber
	// A Boolean value that indicates whether the webpage requested a visible status bar.
	StatusBarVisibility() foundation.NSNumber
	// A Boolean value that indicates whether the webpage requested a visible toolbar.
	ToolbarsVisibility() foundation.NSNumber
}

// Init initializes the instance.
func (w WKWindowFeatures) Init() WKWindowFeatures {
	rv := objc.Send[WKWindowFeatures](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WKWindowFeatures) Autorelease() WKWindowFeatures {
	rv := objc.Send[WKWindowFeatures](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKWindowFeatures creates a new WKWindowFeatures instance.
func NewWKWindowFeatures() WKWindowFeatures {
	class := getWKWindowFeaturesClass()
	rv := objc.Send[WKWindowFeatures](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether to make the containing window window
// resizable.
//
// # Discussion
//
// If the webpage didn’t request a resizable window, this property is `nil`.
//
// See: https://developer.apple.com/documentation/WebKit/WKWindowFeatures/allowsResizing
func (w WKWindowFeatures) AllowsResizing() foundation.NSNumber {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("allowsResizing"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// The requested height of the containing window.
//
// # Discussion
//
// The object in this property contains a [CGFloat] value. If the webpage
// didn’t request a specific window height, this property is `nil`.
//
// See: https://developer.apple.com/documentation/WebKit/WKWindowFeatures/height
//
// [CGFloat]: https://developer.apple.com/documentation/CoreFoundation/CGFloat-swift.struct
func (w WKWindowFeatures) Height() foundation.NSNumber {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("height"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// The requested width of the containing window.
//
// # Discussion
//
// The object in this property contains a [CGFloat] value. If the webpage
// didn’t request a specific window width, this property is `nil`.
//
// See: https://developer.apple.com/documentation/WebKit/WKWindowFeatures/width
//
// [CGFloat]: https://developer.apple.com/documentation/CoreFoundation/CGFloat-swift.struct
func (w WKWindowFeatures) Width() foundation.NSNumber {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("width"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// The requested x-coordinate of the containing window.
//
// # Discussion
//
// The object in this property contains a [CGFloat] value. If the webpage
// didn’t request a specific window position, this property is `nil`.
//
// See: https://developer.apple.com/documentation/WebKit/WKWindowFeatures/x
//
// [CGFloat]: https://developer.apple.com/documentation/CoreFoundation/CGFloat-swift.struct
func (w WKWindowFeatures) X() foundation.NSNumber {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("x"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// The requested y-coordinate of the containing window.
//
// # Discussion
//
// The object in this property contains a [CGFloat] value. If the webpage
// didn’t request a specific window position, this property is `nil`.
//
// See: https://developer.apple.com/documentation/WebKit/WKWindowFeatures/y
//
// [CGFloat]: https://developer.apple.com/documentation/CoreFoundation/CGFloat-swift.struct
func (w WKWindowFeatures) Y() foundation.NSNumber {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("y"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the webpage requests a visible menu
// bar.
//
// # Discussion
//
// If the webpage didn’t request a visible menu bar, this property is `nil`.
//
// See: https://developer.apple.com/documentation/WebKit/WKWindowFeatures/menuBarVisibility
func (w WKWindowFeatures) MenuBarVisibility() foundation.NSNumber {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("menuBarVisibility"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the webpage requested a visible
// status bar.
//
// # Discussion
//
// If the webpage didn’t request a visible status bar, this property is
// `nil`.
//
// See: https://developer.apple.com/documentation/WebKit/WKWindowFeatures/statusBarVisibility
func (w WKWindowFeatures) StatusBarVisibility() foundation.NSNumber {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("statusBarVisibility"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the webpage requested a visible
// toolbar.
//
// # Discussion
//
// If the webpage didn’t request a visible toolbar, this property is `nil`.
//
// See: https://developer.apple.com/documentation/WebKit/WKWindowFeatures/toolbarsVisibility
func (w WKWindowFeatures) ToolbarsVisibility() foundation.NSNumber {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("toolbarsVisibility"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
