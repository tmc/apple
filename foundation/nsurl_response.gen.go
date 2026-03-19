// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [URLResponse] class.
var (
	_URLResponseClass     URLResponseClass
	_URLResponseClassOnce sync.Once
)

func getURLResponseClass() URLResponseClass {
	_URLResponseClassOnce.Do(func() {
		_URLResponseClass = URLResponseClass{class: objc.GetClass("NSURLResponse")}
	})
	return _URLResponseClass
}

// GetURLResponseClass returns the class object for NSURLResponse.
func GetURLResponseClass() URLResponseClass {
	return getURLResponseClass()
}

type URLResponseClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc URLResponseClass) Alloc() URLResponse {
	rv := objc.Send[URLResponse](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// The metadata associated with the response to a URL load request,
// independent of protocol and URL scheme.
//
// # Overview
// 
// The related [NSHTTPURLResponse] class is a commonly used subclass of
// [NSURLResponse] whose objects represent a response to an HTTP URL load
// request and store additional protocol-specific information such as the
// response headers. Whenever you make an HTTP request, the [NSURLResponse]
// object you get back is actually an instance of the [NSHTTPURLResponse]
// class.
//
// # Creating a response
//
//   - [URLResponse.InitWithURLMIMETypeExpectedContentLengthTextEncodingName]: Creates an initialized [URLResponse](<doc://com.apple.foundation/documentation/Foundation/URLResponse>) object with the URL, MIME type, length, and text encoding set to given values.
//
// # Getting the response properties
//
//   - [URLResponse.ExpectedContentLength]: The expected length of the response’s content.
//   - [URLResponse.SuggestedFilename]: A suggested filename for the response data.
//   - [URLResponse.MIMEType]: The MIME type of the response.
//   - [URLResponse.TextEncodingName]: The name of the text encoding provided by the response’s originating source.
//   - [URLResponse.URL]: The URL for the response.
//
// See: https://developer.apple.com/documentation/Foundation/URLResponse
type URLResponse struct {
	objectivec.Object
}

// URLResponseFromID constructs a [URLResponse] from an objc.ID.
//
// The metadata associated with the response to a URL load request,
// independent of protocol and URL scheme.
func URLResponseFromID(id objc.ID) URLResponse {
	return NSURLResponse{objectivec.Object{ID: id}}
}

// NSURLResponseFromID is an alias for [URLResponseFromID] for cross-framework compatibility.
func NSURLResponseFromID(id objc.ID) URLResponse { return URLResponseFromID(id) }
// NOTE: URLResponse adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [URLResponse] class.
//
// # Creating a response
//
//   - [IURLResponse.InitWithURLMIMETypeExpectedContentLengthTextEncodingName]: Creates an initialized [URLResponse](<doc://com.apple.foundation/documentation/Foundation/URLResponse>) object with the URL, MIME type, length, and text encoding set to given values.
//
// # Getting the response properties
//
//   - [IURLResponse.ExpectedContentLength]: The expected length of the response’s content.
//   - [IURLResponse.SuggestedFilename]: A suggested filename for the response data.
//   - [IURLResponse.MIMEType]: The MIME type of the response.
//   - [IURLResponse.TextEncodingName]: The name of the text encoding provided by the response’s originating source.
//   - [IURLResponse.URL]: The URL for the response.
//
// See: https://developer.apple.com/documentation/Foundation/URLResponse
type IURLResponse interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Creating a response

	// Creates an initialized [URLResponse](<doc://com.apple.foundation/documentation/Foundation/URLResponse>) object with the URL, MIME type, length, and text encoding set to given values.
	InitWithURLMIMETypeExpectedContentLengthTextEncodingName(URL INSURL, MIMEType string, length int, name string) URLResponse

	// Topic: Getting the response properties

	// The expected length of the response’s content.
	ExpectedContentLength() int64
	// A suggested filename for the response data.
	SuggestedFilename() string
	// The MIME type of the response.
	MIMEType() string
	// The name of the text encoding provided by the response’s originating source.
	TextEncodingName() string
	// The URL for the response.
	URL() INSURL
}

// Init initializes the instance.
func (u URLResponse) Init() URLResponse {
	rv := objc.Send[URLResponse](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u URLResponse) Autorelease() URLResponse {
	rv := objc.Send[URLResponse](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLResponse creates a new URLResponse instance.
func NewURLResponse() URLResponse {
	class := getURLResponseClass()
	rv := objc.Send[URLResponse](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewURLResponseWithCoder(coder INSCoder) URLResponse {
	instance := getURLResponseClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return URLResponseFromID(rv)
}

// Creates an initialized [NSURLResponse] object with the URL, MIME type,
// length, and text encoding set to given values.
//
// URL: The URL for the new object.
//
// MIMEType: The MIME type.
//
// length: The expected content length.This value should be `–1` if the expected
// length is undetermined
//
// name: The text encoding name. This value may be `nil`.
//
// # Return Value
// 
// The initialized URL response.
//
// # Discussion
// 
// This is the designated initializer for [NSURLResponse].
//
// See: https://developer.apple.com/documentation/Foundation/URLResponse/init(url:mimeType:expectedContentLength:textEncodingName:)
func NewURLResponseWithURLMIMETypeExpectedContentLengthTextEncodingName(URL INSURL, MIMEType string, length int, name string) URLResponse {
	instance := getURLResponseClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:MIMEType:expectedContentLength:textEncodingName:"), URL, objc.String(MIMEType), length, objc.String(name))
	return URLResponseFromID(rv)
}

// Creates an initialized [NSURLResponse] object with the URL, MIME type,
// length, and text encoding set to given values.
//
// URL: The URL for the new object.
//
// MIMEType: The MIME type.
//
// length: The expected content length.This value should be `–1` if the expected
// length is undetermined
//
// name: The text encoding name. This value may be `nil`.
//
// # Return Value
// 
// The initialized URL response.
//
// # Discussion
// 
// This is the designated initializer for [NSURLResponse].
//
// See: https://developer.apple.com/documentation/Foundation/URLResponse/init(url:mimeType:expectedContentLength:textEncodingName:)
func (u URLResponse) InitWithURLMIMETypeExpectedContentLengthTextEncodingName(URL INSURL, MIMEType string, length int, name string) URLResponse {
	rv := objc.Send[URLResponse](u.ID, objc.Sel("initWithURL:MIMEType:expectedContentLength:textEncodingName:"), URL, objc.String(MIMEType), length, objc.String(name))
	return rv
}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (u URLResponse) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (u URLResponse) InitWithCoder(coder INSCoder) URLResponse {
	rv := objc.Send[URLResponse](u.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// The expected length of the response’s content.
//
// # Discussion
// 
// This property’s value is [NSURLResponseUnknownLength] if the length
// can’t be determined.
// 
// Some protocol implementations report the content length as part of the
// response, but not all protocols guarantee to deliver that amount of data.
// Your app should be prepared to deal with more or less data.
//
// [NSURLResponseUnknownLength]: https://developer.apple.com/documentation/Foundation/NSURLResponseUnknownLength
//
// See: https://developer.apple.com/documentation/Foundation/URLResponse/expectedContentLength
func (u URLResponse) ExpectedContentLength() int64 {
	rv := objc.Send[int64](u.ID, objc.Sel("expectedContentLength"))
	return rv
}
// A suggested filename for the response data.
//
// # Discussion
// 
// Accessing this property attempts to generate a filename using the following
// information, in order:
// 
// - A filename specified using the content disposition header. - The last
// path component of the URL. - The host of the URL.
// 
// If the host of URL can’t be converted to a valid filename, the filename
// “unknown” is used.
// 
// In most cases, this property appends the proper file extension based on the
// MIME type. Accessing this property always returns a valid filename
// regardless of whether the resource is saved to disk.
//
// See: https://developer.apple.com/documentation/Foundation/URLResponse/suggestedFilename
func (u URLResponse) SuggestedFilename() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("suggestedFilename"))
	return NSStringFromID(rv).String()
}
// The MIME type of the response.
//
// # Discussion
// 
// The MIME type is often provided by the response’s originating source.
// However, that value may be changed or corrected by a protocol
// implementation if it can be determined that the response’s source
// reported the information incorrectly.
// 
// If the response’s originating source does not provide a MIME type, an
// attempt to guess the MIME type may be made.
//
// See: https://developer.apple.com/documentation/Foundation/URLResponse/mimeType
func (u URLResponse) MIMEType() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("MIMEType"))
	return NSStringFromID(rv).String()
}
// The name of the text encoding provided by the response’s originating
// source.
//
// # Discussion
// 
// If no text encoding was provided by the protocol, this property’s value
// is `nil`.
// 
// You can convert this string to a [CFStringEncoding] value by calling
// [CFStringConvertIANACharSetNameToEncoding(_:)]. You can subsequently
// convert that value to an [NSStringEncoding] value by calling
// [CFStringConvertEncodingToNSStringEncoding(_:)].
//
// [CFStringConvertEncodingToNSStringEncoding(_:)]: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertEncodingToNSStringEncoding(_:)
// [CFStringConvertIANACharSetNameToEncoding(_:)]: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertIANACharSetNameToEncoding(_:)
//
// See: https://developer.apple.com/documentation/Foundation/URLResponse/textEncodingName
func (u URLResponse) TextEncodingName() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("textEncodingName"))
	return NSStringFromID(rv).String()
}
// The URL for the response.
//
// See: https://developer.apple.com/documentation/Foundation/URLResponse/url
func (u URLResponse) URL() INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("URL"))
	return NSURLFromID(objc.ID(rv))
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

