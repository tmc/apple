// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [HTTPURLResponse] class.
var (
	_HTTPURLResponseClass     HTTPURLResponseClass
	_HTTPURLResponseClassOnce sync.Once
)

func getHTTPURLResponseClass() HTTPURLResponseClass {
	_HTTPURLResponseClassOnce.Do(func() {
		_HTTPURLResponseClass = HTTPURLResponseClass{class: objc.GetClass("NSHTTPURLResponse")}
	})
	return _HTTPURLResponseClass
}

// GetHTTPURLResponseClass returns the class object for NSHTTPURLResponse.
func GetHTTPURLResponseClass() HTTPURLResponseClass {
	return getHTTPURLResponseClass()
}

type HTTPURLResponseClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (hc HTTPURLResponseClass) Alloc() HTTPURLResponse {
	rv := objc.Send[HTTPURLResponse](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}







// The metadata associated with the response to an HTTP protocol URL load
// request.
//
// # Overview
// 
// The [NSHTTPURLResponse] class is a subclass of [NSURLResponse] that
// provides methods for accessing information specific to HTTP protocol
// responses. Whenever you make HTTP URL load requests, any response objects
// you get back from the [NSURLSession], [NSURLConnection], or [NSURLDownload]
// class are instances of the [NSHTTPURLResponse] class.
//
// # Initializing a response object
//
//   - [HTTPURLResponse.InitWithURLStatusCodeHTTPVersionHeaderFields]: Initializes an HTTP URL response object with a status code, protocol version, and response headers.
//
// # Getting HTTP response headers
//
//   - [HTTPURLResponse.AllHeaderFields]: All HTTP header fields of the response.
//   - [HTTPURLResponse.ValueForHTTPHeaderField]: Returns the value that corresponds to the given header field.
//
// # Getting response status codes
//
//   - [HTTPURLResponse.StatusCode]: The response’s HTTP status code.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPURLResponse
type HTTPURLResponse struct {
	NSURLResponse
}

// HTTPURLResponseFromID constructs a [HTTPURLResponse] from an objc.ID.
//
// The metadata associated with the response to an HTTP protocol URL load
// request.
func HTTPURLResponseFromID(id objc.ID) HTTPURLResponse {
	return NSHTTPURLResponse{NSURLResponse: NSURLResponseFromID(id)}
}

// NSHTTPURLResponseFromID is an alias for [HTTPURLResponseFromID] for cross-framework compatibility.
func NSHTTPURLResponseFromID(id objc.ID) HTTPURLResponse { return HTTPURLResponseFromID(id) }
// NOTE: HTTPURLResponse adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [HTTPURLResponse] class.
//
// # Initializing a response object
//
//   - [IHTTPURLResponse.InitWithURLStatusCodeHTTPVersionHeaderFields]: Initializes an HTTP URL response object with a status code, protocol version, and response headers.
//
// # Getting HTTP response headers
//
//   - [IHTTPURLResponse.AllHeaderFields]: All HTTP header fields of the response.
//   - [IHTTPURLResponse.ValueForHTTPHeaderField]: Returns the value that corresponds to the given header field.
//
// # Getting response status codes
//
//   - [IHTTPURLResponse.StatusCode]: The response’s HTTP status code.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPURLResponse
type IHTTPURLResponse interface {
	INSURLResponse

	// Topic: Initializing a response object

	// Initializes an HTTP URL response object with a status code, protocol version, and response headers.
	InitWithURLStatusCodeHTTPVersionHeaderFields(url INSURL, statusCode int, HTTPVersion string, headerFields INSDictionary) HTTPURLResponse

	// Topic: Getting HTTP response headers

	// All HTTP header fields of the response.
	AllHeaderFields() INSDictionary
	// Returns the value that corresponds to the given header field.
	ValueForHTTPHeaderField(field string) string

	// Topic: Getting response status codes

	// The response’s HTTP status code.
	StatusCode() int
}





// Init initializes the instance.
func (h HTTPURLResponse) Init() HTTPURLResponse {
	rv := objc.Send[HTTPURLResponse](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h HTTPURLResponse) Autorelease() HTTPURLResponse {
	rv := objc.Send[HTTPURLResponse](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewHTTPURLResponse creates a new HTTPURLResponse instance.
func NewHTTPURLResponse() HTTPURLResponse {
	class := getHTTPURLResponseClass()
	rv := objc.Send[HTTPURLResponse](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewHTTPURLResponseWithCoder(coder INSCoder) HTTPURLResponse {
	instance := getHTTPURLResponseClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return HTTPURLResponseFromID(rv)
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
func NewHTTPURLResponseWithURLMIMETypeExpectedContentLengthTextEncodingName(URL INSURL, MIMEType string, length int, name string) HTTPURLResponse {
	instance := getHTTPURLResponseClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:MIMEType:expectedContentLength:textEncodingName:"), URL, objc.String(MIMEType), length, objc.String(name))
	return HTTPURLResponseFromID(rv)
}


// Initializes an HTTP URL response object with a status code, protocol
// version, and response headers.
//
// url: The URL from which the response was generated.
//
// statusCode: The HTTP status code to return (`404`, for example). See [RFC 2616] for
// details.
// //
// [RFC 2616]: http://www.ietf.org/rfc/rfc2616.txt
//
// HTTPVersion: The version of the HTTP response as returned by the server. This is
// typically represented as “HTTP/1.1”.
//
// headerFields: A dictionary representing the keys and values from the server’s response
// header.
//
// # Return Value
// 
// An initialized [NSHTTPURLResponse] object or `nil` if an error occurred
// during initialization.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPURLResponse/init(url:statusCode:httpVersion:headerFields:)
func NewHTTPURLResponseWithURLStatusCodeHTTPVersionHeaderFields(url INSURL, statusCode int, HTTPVersion string, headerFields INSDictionary) HTTPURLResponse {
	instance := getHTTPURLResponseClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:statusCode:HTTPVersion:headerFields:"), url, statusCode, objc.String(HTTPVersion), headerFields)
	return HTTPURLResponseFromID(rv)
}







// Initializes an HTTP URL response object with a status code, protocol
// version, and response headers.
//
// url: The URL from which the response was generated.
//
// statusCode: The HTTP status code to return (`404`, for example). See [RFC 2616] for
// details.
// //
// [RFC 2616]: http://www.ietf.org/rfc/rfc2616.txt
//
// HTTPVersion: The version of the HTTP response as returned by the server. This is
// typically represented as “HTTP/1.1”.
//
// headerFields: A dictionary representing the keys and values from the server’s response
// header.
//
// # Return Value
// 
// An initialized [NSHTTPURLResponse] object or `nil` if an error occurred
// during initialization.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPURLResponse/init(url:statusCode:httpVersion:headerFields:)
func (h HTTPURLResponse) InitWithURLStatusCodeHTTPVersionHeaderFields(url INSURL, statusCode int, HTTPVersion string, headerFields INSDictionary) HTTPURLResponse {
	rv := objc.Send[HTTPURLResponse](h.ID, objc.Sel("initWithURL:statusCode:HTTPVersion:headerFields:"), url, statusCode, objc.String(HTTPVersion), headerFields)
	return rv
}

// Returns the value that corresponds to the given header field.
//
// field: The name of the header field you want to retrieve. The name is
// case-insensitive.
//
// # Return Value
// 
// The value associated with the given header field, or `nil` if no value is
// associated with the field.
//
// # Discussion
// 
// In keeping with the HTTP RFC, HTTP header field names are case-insensitive.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPURLResponse/value(forHTTPHeaderField:)
func (h HTTPURLResponse) ValueForHTTPHeaderField(field string) string {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("valueForHTTPHeaderField:"), objc.String(field))
	return NSStringFromID(rv).String()
}





// Returns a localized string corresponding to a specified HTTP status code.
//
// statusCode: The HTTP status code. See [RFC 2616] for details.
// //
// [RFC 2616]: http://www.ietf.org/rfc/rfc2616.txt
//
// # Return Value
// 
// A localized string suitable for displaying to users that describes the
// specified status code.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPURLResponse/localizedString(forStatusCode:)
func (_HTTPURLResponseClass HTTPURLResponseClass) LocalizedStringForStatusCode(statusCode int) string {
	rv := objc.Send[objc.ID](objc.ID(_HTTPURLResponseClass.class), objc.Sel("localizedStringForStatusCode:"), statusCode)
	return NSStringFromID(rv).String()
}








// All HTTP header fields of the response.
//
// # Discussion
// 
// The value of this property is a dictionary that contains all the HTTP
// header fields received as part of the server’s response. By examining
// this dictionary, clients can see the “raw” header information returned
// by the HTTP server.
// 
// The keys in this dictionary are the header field names, as received from
// the server. See [RFC 2616] for a list of commonly used HTTP header fields.
// 
// HTTP headers are case insensitive. To simplify your code, URL Loading
// System canonicalizes certain header field names into their standard form.
// For example, if the server sends a `content-length` header, it’s
// automatically adjusted to be `Content-Length`.
// 
// When using Swift, this property is a standard dictionary, so its keys are
// case-sensitive. To perform a case-insensitive header lookup, use the
// [ValueForHTTPHeaderField] method instead.
// 
// In Objective-C, the returned dictionary of headers is case-preserving
// during the set operation (unless the key already exists with a different
// case), and case-insensitive when looking up keys.
// 
// For example, if you set the header `X-foo`, and then later set the header
// `X-Foo`, the dictionary’s key is be `X-foo`, but the value comes from the
// `X-Foo` header.
// 
// # Special considerations
// 
// Prior to OS X v10.7 and iOS 5, canonicalization occurred for all header
// fields. The case-preserving dictionary was first introduced in OS X v10.7.2
// and iOS 5.
//
// [RFC 2616]: http://www.ietf.org/rfc/rfc2616.txt
//
// See: https://developer.apple.com/documentation/Foundation/HTTPURLResponse/allHeaderFields
func (h HTTPURLResponse) AllHeaderFields() INSDictionary {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("allHeaderFields"))
	return NSDictionaryFromID(objc.ID(rv))
}



// The response’s HTTP status code.
//
// # Discussion
// 
// See [RFC 2616] for details.
//
// [RFC 2616]: http://www.ietf.org/rfc/rfc2616.txt
//
// See: https://developer.apple.com/documentation/Foundation/HTTPURLResponse/statusCode
func (h HTTPURLResponse) StatusCode() int {
	rv := objc.Send[int](h.ID, objc.Sel("statusCode"))
	return rv
}




























