// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSMutableURLRequest] class.
var (
	_NSMutableURLRequestClass     NSMutableURLRequestClass
	_NSMutableURLRequestClassOnce sync.Once
)

func getNSMutableURLRequestClass() NSMutableURLRequestClass {
	_NSMutableURLRequestClassOnce.Do(func() {
		_NSMutableURLRequestClass = NSMutableURLRequestClass{class: objc.GetClass("NSMutableURLRequest")}
	})
	return _NSMutableURLRequestClass
}

// GetNSMutableURLRequestClass returns the class object for NSMutableURLRequest.
func GetNSMutableURLRequestClass() NSMutableURLRequestClass {
	return getNSMutableURLRequestClass()
}

type NSMutableURLRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMutableURLRequestClass) Alloc() NSMutableURLRequest {
	rv := objc.Send[NSMutableURLRequest](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A mutable URL load request that is independent of protocol or URL scheme.
//
// # Overview
// 
// In Swift, this object bridges to [NSURLRequest] and you use when you need
// reference semantics or other Foundation-specific behavior.
// 
// [NSMutableURLRequest] is a subclass of [NSURLRequest] that allows you to
// change the request’s properties.
// 
// [NSMutableURLRequest] only represents information about the request. Use
// other classes, such as [NSURLSession], to send the request to a server. See
// [Fetching website data into memory] and [Uploading data to a website] for
// an introduction to these techniques.
// 
// Classes that create a network operation based on a request make a deep copy
// of that request. Thus, changing the request after creating a network
// operation has no effect on the ongoing operation. For example, if you use
// [DataTaskWithRequestCompletionHandler] to create a data task from a
// request, and then later change the request, the data task continues using
// the original request.
//
// [Fetching website data into memory]: https://developer.apple.com/documentation/Foundation/fetching-website-data-into-memory
// [Uploading data to a website]: https://developer.apple.com/documentation/Foundation/uploading-data-to-a-website
//
// # Accessing header fields
//
//   - [NSMutableURLRequest.AddValueForHTTPHeaderField]: Adds a value to the header field.
//   - [NSMutableURLRequest.SetValueForHTTPHeaderField]: Sets a value for the header field.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest
type NSMutableURLRequest struct {
	NSURLRequest
}

// NSMutableURLRequestFromID constructs a [NSMutableURLRequest] from an objc.ID.
//
// A mutable URL load request that is independent of protocol or URL scheme.
func NSMutableURLRequestFromID(id objc.ID) NSMutableURLRequest {
	return NSMutableURLRequest{NSURLRequest: NSURLRequestFromID(id)}
}
// NOTE: NSMutableURLRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSMutableURLRequest] class.
//
// # Accessing header fields
//
//   - [INSMutableURLRequest.AddValueForHTTPHeaderField]: Adds a value to the header field.
//   - [INSMutableURLRequest.SetValueForHTTPHeaderField]: Sets a value for the header field.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest
type INSMutableURLRequest interface {
	INSURLRequest

	// Topic: Accessing header fields

	// Adds a value to the header field.
	AddValueForHTTPHeaderField(value string, field string)
	// Sets a value for the header field.
	SetValueForHTTPHeaderField(value string, field string)
}





// Init initializes the instance.
func (m NSMutableURLRequest) Init() NSMutableURLRequest {
	rv := objc.Send[NSMutableURLRequest](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMutableURLRequest) Autorelease() NSMutableURLRequest {
	rv := objc.Send[NSMutableURLRequest](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMutableURLRequest creates a new NSMutableURLRequest instance.
func NewNSMutableURLRequest() NSMutableURLRequest {
	class := getNSMutableURLRequestClass()
	rv := objc.Send[NSMutableURLRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewMutableURLRequestWithCoder(coder INSCoder) NSMutableURLRequest {
	instance := getNSMutableURLRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSMutableURLRequestFromID(rv)
}


// Creates a URL request for a specified URL.
//
// URL: The URL for the request.
//
// # Return Value
// 
// The initialized URL request.
//
// # Discussion
// 
// The request is created with the with the default cache policy
// ([URLRequestUseProtocolCachePolicy]), and the default timeout interval (60
// seconds).
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/init(url:)
func NewMutableURLRequestWithURL(URL INSURL) NSMutableURLRequest {
	instance := getNSMutableURLRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), URL)
	return NSMutableURLRequestFromID(rv)
}


// Creates a URL request with the specified URL, cache policy, and timeout
// values.
//
// URL: The URL for the request.
//
// cachePolicy: The cache policy for the request.
//
// timeoutInterval: The timeout interval for the request, in seconds.
//
// # Return Value
// 
// The initialized URL request.
//
// # Discussion
// 
// This is the designated initializer for [NSURLRequest].
//
// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/init(url:cachePolicy:timeoutInterval:)
func NewMutableURLRequestWithURLCachePolicyTimeoutInterval(URL INSURL, cachePolicy NSURLRequestCachePolicy, timeoutInterval float64) NSMutableURLRequest {
	instance := getNSMutableURLRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:cachePolicy:timeoutInterval:"), URL, cachePolicy, timeoutInterval)
	return NSMutableURLRequestFromID(rv)
}







// Adds a value to the header field.
//
// value: The value for the header field.
//
// field: The name of the header field. In keeping with the HTTP RFC, HTTP header
// field names are case insensitive.
//
// # Discussion
// 
// This method provides the ability to add values to header fields
// incrementally. If a value was previously set for the specified field, the
// supplied value is appended to the existing value using the appropriate
// field delimiter (a comma).
// 
// Certain header fields are reserved (see [NSURLRequest]). Do not use this
// method to change such headers.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/addValue(_:forHTTPHeaderField:)
func (m NSMutableURLRequest) AddValueForHTTPHeaderField(value string, field string) {
	objc.Send[objc.ID](m.ID, objc.Sel("addValue:forHTTPHeaderField:"), objc.String(value), objc.String(field))
}

// Sets a value for the header field.
//
// value: The new value for the header field. Any existing value for the field is
// replaced by the new value.
//
// field: The name of the header field to set. In keeping with the HTTP RFC, HTTP
// header field names are case insensitive.
//
// # Discussion
// 
// Certain header fields are reserved. Do not use this method to set such
// headers. Specifically, there is no need for you to set the `Content-Length`
// header. See [NSURLRequest].
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/setValue(_:forHTTPHeaderField:)
func (m NSMutableURLRequest) SetValueForHTTPHeaderField(value string, field string) {
	objc.Send[objc.ID](m.ID, objc.Sel("setValue:forHTTPHeaderField:"), objc.String(value), objc.String(field))
}





































