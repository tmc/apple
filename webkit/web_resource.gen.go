// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WebResource] class.
var (
	_WebResourceClass     WebResourceClass
	_WebResourceClassOnce sync.Once
)

func getWebResourceClass() WebResourceClass {
	_WebResourceClassOnce.Do(func() {
		_WebResourceClass = WebResourceClass{class: objc.GetClass("WebResource")}
	})
	return _WebResourceClass
}

// GetWebResourceClass returns the class object for WebResource.
func GetWebResourceClass() WebResourceClass {
	return getWebResourceClass()
}

type WebResourceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WebResourceClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WebResourceClass) Alloc() WebResource {
	rv := objc.Send[WebResource](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// A [WebResource] object represents a downloaded URL. It encapsulates the
// data of the download as well as other resource properties such as the URL,
// MIME type, and frame name.
//
// # Overview
//
// Use the [WebResource.InitWithDataURLMIMETypeTextEncodingNameFrameName] method to
// initialize a newly created [WebResource] object. Use the other methods in
// this class to get the properties of a [WebResource] object.
//
// # Initializing
//
//   - [WebResource.InitWithDataURLMIMETypeTextEncodingNameFrameName]: Initializes and returns a web resource instance.
//
// # Getting attributes
//
//   - [WebResource.Data]: The receiver’s data.
//   - [WebResource.URL]: The receiver’s URL.
//   - [WebResource.MIMEType]: The receiver’s MIME type.
//   - [WebResource.TextEncodingName]: The receiver’s text encoding name.
//   - [WebResource.FrameName]: The name of the frame. If the receiver does not represent the contents of an entire HTML frame, this is `nil`.
//
// See: https://developer.apple.com/documentation/WebKit/WebResource
type WebResource struct {
	objectivec.Object
}

// WebResourceFromID constructs a [WebResource] from an objc.ID.
//
// A [WebResource] object represents a downloaded URL. It encapsulates the
// data of the download as well as other resource properties such as the URL,
// MIME type, and frame name.
func WebResourceFromID(id objc.ID) WebResource {
	return WebResource{objectivec.Object{ID: id}}
}

// NOTE: WebResource adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WebResource] class.
//
// # Initializing
//
//   - [IWebResource.InitWithDataURLMIMETypeTextEncodingNameFrameName]: Initializes and returns a web resource instance.
//
// # Getting attributes
//
//   - [IWebResource.Data]: The receiver’s data.
//   - [IWebResource.URL]: The receiver’s URL.
//   - [IWebResource.MIMEType]: The receiver’s MIME type.
//   - [IWebResource.TextEncodingName]: The receiver’s text encoding name.
//   - [IWebResource.FrameName]: The name of the frame. If the receiver does not represent the contents of an entire HTML frame, this is `nil`.
//
// See: https://developer.apple.com/documentation/WebKit/WebResource
type IWebResource interface {
	objectivec.IObject

	// Topic: Initializing

	// Initializes and returns a web resource instance.
	InitWithDataURLMIMETypeTextEncodingNameFrameName(data foundation.INSData, URL foundation.INSURL, MIMEType string, textEncodingName string, frameName string) WebResource

	// Topic: Getting attributes

	// The receiver’s data.
	Data() foundation.INSData
	// The receiver’s URL.
	URL() foundation.INSURL
	// The receiver’s MIME type.
	MIMEType() string
	// The receiver’s text encoding name.
	TextEncodingName() string
	// The name of the frame. If the receiver does not represent the contents of an entire HTML frame, this is `nil`.
	FrameName() string

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (w WebResource) Init() WebResource {
	rv := objc.Send[WebResource](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WebResource) Autorelease() WebResource {
	rv := objc.Send[WebResource](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebResource creates a new WebResource instance.
func NewWebResource() WebResource {
	class := getWebResourceClass()
	rv := objc.Send[WebResource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes and returns a web resource instance.
//
// data: The download data.
//
// URL: The download URL.
//
// MIMEType: The MIME type of the data.
//
// textEncodingName: The IANA encoding name (for example, “utf-8” or “utf-16”). This
// parameter may be `nil`.
//
// frameName: The name of the frame. Use this parameter if the resource represents the
// contents of an entire HTML frame; otherwise pass `nil`.
//
// # Return Value
//
// An initialized web resource.
//
// See: https://developer.apple.com/documentation/WebKit/WebResource/init(data:url:mimeType:textEncodingName:frameName:)
func NewWebResourceWithDataURLMIMETypeTextEncodingNameFrameName(data foundation.INSData, URL foundation.INSURL, MIMEType string, textEncodingName string, frameName string) WebResource {
	instance := getWebResourceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:URL:MIMEType:textEncodingName:frameName:"), data, URL, objc.String(MIMEType), objc.String(textEncodingName), objc.String(frameName))
	return WebResourceFromID(rv)
}

// Initializes and returns a web resource instance.
//
// data: The download data.
//
// URL: The download URL.
//
// MIMEType: The MIME type of the data.
//
// textEncodingName: The IANA encoding name (for example, “utf-8” or “utf-16”). This
// parameter may be `nil`.
//
// frameName: The name of the frame. Use this parameter if the resource represents the
// contents of an entire HTML frame; otherwise pass `nil`.
//
// # Return Value
//
// An initialized web resource.
//
// See: https://developer.apple.com/documentation/WebKit/WebResource/init(data:url:mimeType:textEncodingName:frameName:)
func (w WebResource) InitWithDataURLMIMETypeTextEncodingNameFrameName(data foundation.INSData, URL foundation.INSURL, MIMEType string, textEncodingName string, frameName string) WebResource {
	rv := objc.Send[WebResource](w.ID, objc.Sel("initWithData:URL:MIMEType:textEncodingName:frameName:"), data, URL, objc.String(MIMEType), objc.String(textEncodingName), objc.String(frameName))
	return rv
}
func (w WebResource) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](w.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The receiver’s data.
//
// See: https://developer.apple.com/documentation/WebKit/WebResource/data
func (w WebResource) Data() foundation.INSData {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("data"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// The receiver’s URL.
//
// See: https://developer.apple.com/documentation/WebKit/WebResource/url
func (w WebResource) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// The receiver’s MIME type.
//
// See: https://developer.apple.com/documentation/WebKit/WebResource/mimeType
func (w WebResource) MIMEType() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("MIMEType"))
	return foundation.NSStringFromID(rv).String()
}

// The receiver’s text encoding name.
//
// See: https://developer.apple.com/documentation/WebKit/WebResource/textEncodingName
func (w WebResource) TextEncodingName() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("textEncodingName"))
	return foundation.NSStringFromID(rv).String()
}

// The name of the frame. If the receiver does not represent the contents of
// an entire HTML frame, this is `nil`.
//
// See: https://developer.apple.com/documentation/WebKit/WebResource/frameName
func (w WebResource) FrameName() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("frameName"))
	return foundation.NSStringFromID(rv).String()
}
