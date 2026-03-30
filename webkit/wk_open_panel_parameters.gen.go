// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKOpenPanelParameters] class.
var (
	_WKOpenPanelParametersClass     WKOpenPanelParametersClass
	_WKOpenPanelParametersClassOnce sync.Once
)

func getWKOpenPanelParametersClass() WKOpenPanelParametersClass {
	_WKOpenPanelParametersClassOnce.Do(func() {
		_WKOpenPanelParametersClass = WKOpenPanelParametersClass{class: objc.GetClass("WKOpenPanelParameters")}
	})
	return _WKOpenPanelParametersClass
}

// GetWKOpenPanelParametersClass returns the class object for WKOpenPanelParameters.
func GetWKOpenPanelParametersClass() WKOpenPanelParametersClass {
	return getWKOpenPanelParametersClass()
}

type WKOpenPanelParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKOpenPanelParametersClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKOpenPanelParametersClass) Alloc() WKOpenPanelParameters {
	rv := objc.Send[WKOpenPanelParameters](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// The configuration details of a file upload control in your web content.
//
// # Overview
//
// Use a [WKOpenPanelParameters] to determine the configuration of a file
// upload control. You don’t create this object directly. Instead, a web
// view creates one and passes it to the
// [WebViewRunOpenPanelWithParametersInitiatedByFrameCompletionHandler] method
// of its UI delegate object when it displays a file upload control.
//
// # Configuring the panel parameters
//
//   - [WKOpenPanelParameters.AllowsMultipleSelection]: A Boolean value that indicates whether the file upload control supports multiple files.
//   - [WKOpenPanelParameters.AllowsDirectories]: A Boolean value that indicates whether the file upload control supports the selection of directories.
//
// See: https://developer.apple.com/documentation/WebKit/WKOpenPanelParameters
type WKOpenPanelParameters struct {
	objectivec.Object
}

// WKOpenPanelParametersFromID constructs a [WKOpenPanelParameters] from an objc.ID.
//
// The configuration details of a file upload control in your web content.
func WKOpenPanelParametersFromID(id objc.ID) WKOpenPanelParameters {
	return WKOpenPanelParameters{objectivec.Object{ID: id}}
}

// NOTE: WKOpenPanelParameters adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKOpenPanelParameters] class.
//
// # Configuring the panel parameters
//
//   - [IWKOpenPanelParameters.AllowsMultipleSelection]: A Boolean value that indicates whether the file upload control supports multiple files.
//   - [IWKOpenPanelParameters.AllowsDirectories]: A Boolean value that indicates whether the file upload control supports the selection of directories.
//
// See: https://developer.apple.com/documentation/WebKit/WKOpenPanelParameters
type IWKOpenPanelParameters interface {
	objectivec.IObject

	// Topic: Configuring the panel parameters

	// A Boolean value that indicates whether the file upload control supports multiple files.
	AllowsMultipleSelection() bool
	// A Boolean value that indicates whether the file upload control supports the selection of directories.
	AllowsDirectories() bool
}

// Init initializes the instance.
func (o WKOpenPanelParameters) Init() WKOpenPanelParameters {
	rv := objc.Send[WKOpenPanelParameters](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o WKOpenPanelParameters) Autorelease() WKOpenPanelParameters {
	rv := objc.Send[WKOpenPanelParameters](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKOpenPanelParameters creates a new WKOpenPanelParameters instance.
func NewWKOpenPanelParameters() WKOpenPanelParameters {
	class := getWKOpenPanelParametersClass()
	rv := objc.Send[WKOpenPanelParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether the file upload control supports
// multiple files.
//
// See: https://developer.apple.com/documentation/WebKit/WKOpenPanelParameters/allowsMultipleSelection
func (o WKOpenPanelParameters) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}

// A Boolean value that indicates whether the file upload control supports the
// selection of directories.
//
// See: https://developer.apple.com/documentation/WebKit/WKOpenPanelParameters/allowsDirectories
func (o WKOpenPanelParameters) AllowsDirectories() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("allowsDirectories"))
	return rv
}
