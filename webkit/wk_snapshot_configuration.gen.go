// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKSnapshotConfiguration] class.
var (
	_WKSnapshotConfigurationClass     WKSnapshotConfigurationClass
	_WKSnapshotConfigurationClassOnce sync.Once
)

func getWKSnapshotConfigurationClass() WKSnapshotConfigurationClass {
	_WKSnapshotConfigurationClassOnce.Do(func() {
		_WKSnapshotConfigurationClass = WKSnapshotConfigurationClass{class: objc.GetClass("WKSnapshotConfiguration")}
	})
	return _WKSnapshotConfigurationClass
}

// GetWKSnapshotConfigurationClass returns the class object for WKSnapshotConfiguration.
func GetWKSnapshotConfigurationClass() WKSnapshotConfigurationClass {
	return getWKSnapshotConfigurationClass()
}

type WKSnapshotConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKSnapshotConfigurationClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKSnapshotConfigurationClass) Alloc() WKSnapshotConfiguration {
	rv := objc.Send[WKSnapshotConfiguration](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// The configuration data to use when generating an image from a web view’s
// contents.
//
// # Overview
//
// Create a [WKSnapshotConfiguration] object when you want to generate an
// image based on your web view’s content. Use this object to specify the
// portion of the web view to capture and the capture behavior. To generate
// the snapshot, pass the configuration object to the
// [TakeSnapshotWithConfigurationCompletionHandler] method of [WKWebView],
// which returns a platform-native image for you to use.
//
// # Specifying the snapshot dimensions
//
//   - [WKSnapshotConfiguration.Rect]: The portion of your web view to capture, specified as a rectangle in the view’s coordinate system.
//   - [WKSnapshotConfiguration.SetRect]
//   - [WKSnapshotConfiguration.SnapshotWidth]: The width of the captured image, in points.
//   - [WKSnapshotConfiguration.SetSnapshotWidth]
//
// # Configuring the capture behavior
//
//   - [WKSnapshotConfiguration.AfterScreenUpdates]: A Boolean value that indicates whether to take the snapshot after incorporating any pending screen updates.
//   - [WKSnapshotConfiguration.SetAfterScreenUpdates]
//
// See: https://developer.apple.com/documentation/WebKit/WKSnapshotConfiguration
type WKSnapshotConfiguration struct {
	objectivec.Object
}

// WKSnapshotConfigurationFromID constructs a [WKSnapshotConfiguration] from an objc.ID.
//
// The configuration data to use when generating an image from a web view’s
// contents.
func WKSnapshotConfigurationFromID(id objc.ID) WKSnapshotConfiguration {
	return WKSnapshotConfiguration{objectivec.Object{ID: id}}
}

// NOTE: WKSnapshotConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKSnapshotConfiguration] class.
//
// # Specifying the snapshot dimensions
//
//   - [IWKSnapshotConfiguration.Rect]: The portion of your web view to capture, specified as a rectangle in the view’s coordinate system.
//   - [IWKSnapshotConfiguration.SetRect]
//   - [IWKSnapshotConfiguration.SnapshotWidth]: The width of the captured image, in points.
//   - [IWKSnapshotConfiguration.SetSnapshotWidth]
//
// # Configuring the capture behavior
//
//   - [IWKSnapshotConfiguration.AfterScreenUpdates]: A Boolean value that indicates whether to take the snapshot after incorporating any pending screen updates.
//   - [IWKSnapshotConfiguration.SetAfterScreenUpdates]
//
// See: https://developer.apple.com/documentation/WebKit/WKSnapshotConfiguration
type IWKSnapshotConfiguration interface {
	objectivec.IObject

	// Topic: Specifying the snapshot dimensions

	// The portion of your web view to capture, specified as a rectangle in the view’s coordinate system.
	Rect() corefoundation.CGRect
	SetRect(value corefoundation.CGRect)
	// The width of the captured image, in points.
	SnapshotWidth() foundation.NSNumber
	SetSnapshotWidth(value foundation.NSNumber)

	// Topic: Configuring the capture behavior

	// A Boolean value that indicates whether to take the snapshot after incorporating any pending screen updates.
	AfterScreenUpdates() bool
	SetAfterScreenUpdates(value bool)
}

// Init initializes the instance.
func (s WKSnapshotConfiguration) Init() WKSnapshotConfiguration {
	rv := objc.Send[WKSnapshotConfiguration](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s WKSnapshotConfiguration) Autorelease() WKSnapshotConfiguration {
	rv := objc.Send[WKSnapshotConfiguration](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKSnapshotConfiguration creates a new WKSnapshotConfiguration instance.
func NewWKSnapshotConfiguration() WKSnapshotConfiguration {
	class := getWKSnapshotConfigurationClass()
	rv := objc.Send[WKSnapshotConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
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
// See: https://developer.apple.com/documentation/WebKit/WKSnapshotConfiguration/rect
//
// [CGRectNull]: https://developer.apple.com/documentation/CoreGraphics/CGRectNull
func (s WKSnapshotConfiguration) Rect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("rect"))
	return corefoundation.CGRect(rv)
}
func (s WKSnapshotConfiguration) SetRect(value corefoundation.CGRect) {
	objc.Send[struct{}](s.ID, objc.Sel("setRect:"), value)
}

// The width of the captured image, in points.
//
// # Discussion
//
// Use this property to scale the generated image to the specified width. The
// web view maintains the aspect ratio of the captured content, but scales it
// to match the width you specify.
//
// The default value of this property is `nil`, which returns an image whose
// size matches the original size of the captured rectangle.
//
// See: https://developer.apple.com/documentation/WebKit/WKSnapshotConfiguration/snapshotWidth
func (s WKSnapshotConfiguration) SnapshotWidth() foundation.NSNumber {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("snapshotWidth"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (s WKSnapshotConfiguration) SetSnapshotWidth(value foundation.NSNumber) {
	objc.Send[struct{}](s.ID, objc.Sel("setSnapshotWidth:"), value)
}

// A Boolean value that indicates whether to take the snapshot after
// incorporating any pending screen updates.
//
// # Discussion
//
// The default value of this property is true, which causes the web view to
// incorporate any recent changes to the view’s content and then generate
// the snapshot. If you change the value to false, the web view takes the
// snapshot immediately, and before incorporating any new changes.
//
// See: https://developer.apple.com/documentation/WebKit/WKSnapshotConfiguration/afterScreenUpdates
func (s WKSnapshotConfiguration) AfterScreenUpdates() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("afterScreenUpdates"))
	return rv
}
func (s WKSnapshotConfiguration) SetAfterScreenUpdates(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setAfterScreenUpdates:"), value)
}
