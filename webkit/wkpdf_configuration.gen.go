// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKPDFConfiguration] class.
var (
	_WKPDFConfigurationClass     WKPDFConfigurationClass
	_WKPDFConfigurationClassOnce sync.Once
)

func getWKPDFConfigurationClass() WKPDFConfigurationClass {
	_WKPDFConfigurationClassOnce.Do(func() {
		_WKPDFConfigurationClass = WKPDFConfigurationClass{class: objc.GetClass("WKPDFConfiguration")}
	})
	return _WKPDFConfigurationClass
}

// GetWKPDFConfigurationClass returns the class object for WKPDFConfiguration.
func GetWKPDFConfigurationClass() WKPDFConfigurationClass {
	return getWKPDFConfigurationClass()
}

type WKPDFConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKPDFConfigurationClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKPDFConfigurationClass) Alloc() WKPDFConfiguration {
	rv := objc.Send[WKPDFConfiguration](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// The configuration data to use when generating a PDF representation of a web
// view’s contents.
//
// # Overview
//
// Create a [WKPDFConfiguration] object when you want to generate a PDF
// version of your web view’s content. Use this object to specify the
// portion of the web view to capture. To generate the PDF content, pass the
// configuration object to the [createPDF(configuration:completionHandler:)]
// method of [WKWebView], which returns the PDF data for you to use.
//
// # Specifying snapshot properties
//
//   - [WKPDFConfiguration.AllowTransparentBackground]: A Boolean value that indicates whether the PDF may have a transparent background.
//   - [WKPDFConfiguration.SetAllowTransparentBackground]
//
// See: https://developer.apple.com/documentation/WebKit/WKPDFConfiguration
//
// [createPDF(configuration:completionHandler:)]: https://developer.apple.com/documentation/WebKit/WKWebView/createPDF(configuration:completionHandler:)
type WKPDFConfiguration struct {
	objectivec.Object
}

// WKPDFConfigurationFromID constructs a [WKPDFConfiguration] from an objc.ID.
//
// The configuration data to use when generating a PDF representation of a web
// view’s contents.
func WKPDFConfigurationFromID(id objc.ID) WKPDFConfiguration {
	return WKPDFConfiguration{objectivec.Object{ID: id}}
}

// NOTE: WKPDFConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKPDFConfiguration] class.
//
// # Specifying snapshot properties
//
//   - [IWKPDFConfiguration.AllowTransparentBackground]: A Boolean value that indicates whether the PDF may have a transparent background.
//   - [IWKPDFConfiguration.SetAllowTransparentBackground]
//
// See: https://developer.apple.com/documentation/WebKit/WKPDFConfiguration
type IWKPDFConfiguration interface {
	objectivec.IObject

	// Topic: Specifying snapshot properties

	// A Boolean value that indicates whether the PDF may have a transparent background.
	AllowTransparentBackground() bool
	SetAllowTransparentBackground(value bool)

	// The portion of your web view to capture, specified as a rectangle in the view’s coordinate system.
	Rect() corefoundation.CGRect
	SetRect(value corefoundation.CGRect)
}

// Init initializes the instance.
func (p WKPDFConfiguration) Init() WKPDFConfiguration {
	rv := objc.Send[WKPDFConfiguration](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p WKPDFConfiguration) Autorelease() WKPDFConfiguration {
	rv := objc.Send[WKPDFConfiguration](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKPDFConfiguration creates a new WKPDFConfiguration instance.
func NewWKPDFConfiguration() WKPDFConfiguration {
	class := getWKPDFConfigurationClass()
	rv := objc.Send[WKPDFConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether the PDF may have a transparent
// background.
//
// See: https://developer.apple.com/documentation/WebKit/WKPDFConfiguration/allowTransparentBackground
func (p WKPDFConfiguration) AllowTransparentBackground() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("allowTransparentBackground"))
	return rv
}
func (p WKPDFConfiguration) SetAllowTransparentBackground(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAllowTransparentBackground:"), value)
}

// The portion of your web view to capture, specified as a rectangle in the
// view’s coordinate system.
//
// # Discussion
//
// The default value of this property is [CGRectNull], which captures
// everything in the view’s bounds rectangle. If you specify a custom
// rectangle, it must lie within the bounds rectangle of the [WKWebView]
// object.
//
// See: https://developer.apple.com/documentation/WebKit/WKPDFConfiguration/rect-3xww9
//
// [CGRectNull]: https://developer.apple.com/documentation/CoreGraphics/CGRectNull
func (p WKPDFConfiguration) Rect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p.ID, objc.Sel("rect"))
	return corefoundation.CGRect(rv)
}
func (p WKPDFConfiguration) SetRect(value corefoundation.CGRect) {
	objc.Send[struct{}](p.ID, objc.Sel("setRect:"), value)
}
