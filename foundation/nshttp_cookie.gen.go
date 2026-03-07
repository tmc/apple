// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [HTTPCookie] class.
var (
	_HTTPCookieClass     HTTPCookieClass
	_HTTPCookieClassOnce sync.Once
)

func getHTTPCookieClass() HTTPCookieClass {
	_HTTPCookieClassOnce.Do(func() {
		_HTTPCookieClass = HTTPCookieClass{class: objc.GetClass("NSHTTPCookie")}
	})
	return _HTTPCookieClass
}

// GetHTTPCookieClass returns the class object for NSHTTPCookie.
func GetHTTPCookieClass() HTTPCookieClass {
	return getHTTPCookieClass()
}

type HTTPCookieClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (hc HTTPCookieClass) Alloc() HTTPCookie {
	rv := objc.Send[HTTPCookie](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}







// A representation of an HTTP cookie.
//
// # Overview
// 
// An [NSHTTPCookie] object is immutable, initialized from a dictionary that
// contains the attributes of the cookie. This class supports two different
// cookie versions:
// 
// - Version 0: The original cookie format defined by Netscape. Most cookies
// are in this format. - Version 1: The cookie format defined in [RFC 6265],
// HTTP State Management Mechanism.
//
// [RFC 6265]: https://tools.ietf.org/html/rfc6265
//
// # Creating cookies
//
//   - [HTTPCookie.InitWithProperties]: Initializes an HTTP cookie object with the given cookie properties.
//
// # Getting cookie host properties
//
//   - [HTTPCookie.Domain]: The domain of the cookie.
//   - [HTTPCookie.Path]: The cookie’s path.
//   - [HTTPCookie.PortList]: The cookie’s port list.
//
// # Getting cookie metadata
//
//   - [HTTPCookie.Name]: The cookie’s name.
//   - [HTTPCookie.Value]: The cookie’s string value.
//   - [HTTPCookie.Version]: The cookie’s version.
//
// # Determining cookie lifespan
//
//   - [HTTPCookie.ExpiresDate]: The cookie’s expiration date.
//   - [HTTPCookie.SessionOnly]: A Boolean value that indicates whether the cookie should be discarded at the end of the session (regardless of expiration date).
//
// # Securing cookies
//
//   - [HTTPCookie.HTTPOnly]: A Boolean value that indicates whether the cookie should only be sent to HTTP servers.
//   - [HTTPCookie.Secure]: A Boolean value that indicates whether the cookie may only be sent over secure channels.
//   - [HTTPCookie.SameSitePolicy]: A Boolean value that indicates whether to restrict the cookie to requests sent back to the same site that created it.
//
// # Accessing cookie properties as key-value pairs
//
//   - [HTTPCookie.Properties]: The cookie’s properties.
//
// # Getting user-readable cookie metadata
//
//   - [HTTPCookie.Comment]: The cookie’s comment string.
//   - [HTTPCookie.CommentURL]: The cookie’s comment URL.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookie
type HTTPCookie struct {
	objectivec.Object
}

// HTTPCookieFromID constructs a [HTTPCookie] from an objc.ID.
//
// A representation of an HTTP cookie.
func HTTPCookieFromID(id objc.ID) HTTPCookie {
	return NSHTTPCookie{objectivec.Object{id}}
}

// NSHTTPCookieFromID is an alias for [HTTPCookieFromID] for cross-framework compatibility.
func NSHTTPCookieFromID(id objc.ID) HTTPCookie { return HTTPCookieFromID(id) }
// NOTE: HTTPCookie adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [HTTPCookie] class.
//
// # Creating cookies
//
//   - [IHTTPCookie.InitWithProperties]: Initializes an HTTP cookie object with the given cookie properties.
//
// # Getting cookie host properties
//
//   - [IHTTPCookie.Domain]: The domain of the cookie.
//   - [IHTTPCookie.Path]: The cookie’s path.
//   - [IHTTPCookie.PortList]: The cookie’s port list.
//
// # Getting cookie metadata
//
//   - [IHTTPCookie.Name]: The cookie’s name.
//   - [IHTTPCookie.Value]: The cookie’s string value.
//   - [IHTTPCookie.Version]: The cookie’s version.
//
// # Determining cookie lifespan
//
//   - [IHTTPCookie.ExpiresDate]: The cookie’s expiration date.
//   - [IHTTPCookie.SessionOnly]: A Boolean value that indicates whether the cookie should be discarded at the end of the session (regardless of expiration date).
//
// # Securing cookies
//
//   - [IHTTPCookie.HTTPOnly]: A Boolean value that indicates whether the cookie should only be sent to HTTP servers.
//   - [IHTTPCookie.Secure]: A Boolean value that indicates whether the cookie may only be sent over secure channels.
//   - [IHTTPCookie.SameSitePolicy]: A Boolean value that indicates whether to restrict the cookie to requests sent back to the same site that created it.
//
// # Accessing cookie properties as key-value pairs
//
//   - [IHTTPCookie.Properties]: The cookie’s properties.
//
// # Getting user-readable cookie metadata
//
//   - [IHTTPCookie.Comment]: The cookie’s comment string.
//   - [IHTTPCookie.CommentURL]: The cookie’s comment URL.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookie
type IHTTPCookie interface {
	objectivec.IObject

	// Topic: Creating cookies

	// Initializes an HTTP cookie object with the given cookie properties.
	InitWithProperties(properties INSDictionary) HTTPCookie

	// Topic: Getting cookie host properties

	// The domain of the cookie.
	Domain() string
	// The cookie’s path.
	Path() string
	// The cookie’s port list.
	PortList() []NSNumber

	// Topic: Getting cookie metadata

	// The cookie’s name.
	Name() string
	// The cookie’s string value.
	Value() string
	// The cookie’s version.
	Version() uint

	// Topic: Determining cookie lifespan

	// The cookie’s expiration date.
	ExpiresDate() INSDate
	// A Boolean value that indicates whether the cookie should be discarded at the end of the session (regardless of expiration date).
	SessionOnly() bool

	// Topic: Securing cookies

	// A Boolean value that indicates whether the cookie should only be sent to HTTP servers.
	HTTPOnly() bool
	// A Boolean value that indicates whether the cookie may only be sent over secure channels.
	Secure() bool
	// A Boolean value that indicates whether to restrict the cookie to requests sent back to the same site that created it.
	SameSitePolicy() NSHTTPCookieStringPolicy

	// Topic: Accessing cookie properties as key-value pairs

	// The cookie’s properties.
	Properties() INSDictionary

	// Topic: Getting user-readable cookie metadata

	// The cookie’s comment string.
	Comment() string
	// The cookie’s comment URL.
	CommentURL() INSURL
}





// Init initializes the instance.
func (h HTTPCookie) Init() HTTPCookie {
	rv := objc.Send[HTTPCookie](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h HTTPCookie) Autorelease() HTTPCookie {
	rv := objc.Send[HTTPCookie](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewHTTPCookie creates a new HTTPCookie instance.
func NewHTTPCookie() HTTPCookie {
	class := getHTTPCookieClass()
	rv := objc.Send[HTTPCookie](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Initializes an HTTP cookie object with the given cookie properties.
//
// properties: The properties for the new cookie object, expressed as key-value pairs.
//
// # Return Value
// 
// A new cookie object, with the given properies.
//
// # Discussion
// 
// This initializer returns `nil` if the provided properties are invalid. To
// successfully create a cookie, you must provide values for (at least) the
// [path], [name], and [value] keys, and either the [originURL] key or the
// [domain] key.
// 
// See Accepting cookies for more information on the available cookie
// attribute constants and the constraints imposed on the values in the
// dictionary.
//
// [domain]: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/domain
// [name]: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/name
// [originURL]: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/originURL
// [path]: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/path
// [value]: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/value
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookie/init(properties:)
func NewHTTPCookieWithProperties(properties INSDictionary) HTTPCookie {
	instance := getHTTPCookieClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProperties:"), properties)
	return HTTPCookieFromID(rv)
}







// Initializes an HTTP cookie object with the given cookie properties.
//
// properties: The properties for the new cookie object, expressed as key-value pairs.
//
// # Return Value
// 
// A new cookie object, with the given properies.
//
// # Discussion
// 
// This initializer returns `nil` if the provided properties are invalid. To
// successfully create a cookie, you must provide values for (at least) the
// [path], [name], and [value] keys, and either the [originURL] key or the
// [domain] key.
// 
// See Accepting cookies for more information on the available cookie
// attribute constants and the constraints imposed on the values in the
// dictionary.
//
// [domain]: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/domain
// [name]: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/name
// [originURL]: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/originURL
// [path]: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/path
// [value]: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/value
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookie/init(properties:)
func (h HTTPCookie) InitWithProperties(properties INSDictionary) HTTPCookie {
	rv := objc.Send[HTTPCookie](h.ID, objc.Sel("initWithProperties:"), properties)
	return rv
}





// Creates an array of HTTP cookies that corresponds to the provided response
// header fields for the provided URL.
//
// headerFields: The header fields used to create the [NSHTTPCookie] objects.
//
// URL: The URL associated with the created cookies.
//
// # Return Value
// 
// The array of created cookies.
//
// # Discussion
// 
// This method ignores irrelevant header fields in `headerFields`, allowing
// dictionaries to contain additional data.
// 
// If `headerFields` doesn’t specify a domain for a given cookie, the cookie
// is created with a default domain value of [URL].
// 
// If `headerFields` doesn’t specify a path for a given cookie, the cookie
// is created with a default path value of `"/"`.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookie/cookies(withResponseHeaderFields:for:)
func (_HTTPCookieClass HTTPCookieClass) CookiesWithResponseHeaderFieldsForURL(headerFields INSDictionary, URL INSURL) []HTTPCookie {
	rv := objc.Send[[]objc.ID](objc.ID(_HTTPCookieClass.class), objc.Sel("cookiesWithResponseHeaderFields:forURL:"), headerFields, URL)
	return objc.ConvertSlice(rv, func(id objc.ID) NSHTTPCookie {
		return NSHTTPCookieFromID(id)
	})
}

// Converts an array of cookies to a dictionary of header fields.
//
// cookies: The cookies from which the header fields are created.
//
// # Return Value
// 
// The dictionary of header fields created from the provided cookies.
//
// # Discussion
// 
// To send these headers as part of a URL request to a remote server, create
// an [NSMutableURLRequest] object, then call the [AllHTTPHeaderFields] or
// [SetValueForHTTPHeaderField] method to set the provided headers for the
// request. Finally, initialize and start an [NSURLSessionTask] instance based
// on that request object.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookie/requestHeaderFields(with:)
func (_HTTPCookieClass HTTPCookieClass) RequestHeaderFieldsWithCookies(cookies []NSHTTPCookie) INSDictionary {
	rv := objc.Send[objc.ID](objc.ID(_HTTPCookieClass.class), objc.Sel("requestHeaderFieldsWithCookies:"), objectivec.IObjectSliceToNSArray(cookies))
	return NSDictionaryFromID(rv)
}

// Creates and initializes an HTTP cookie object using the provided
// properties.
//
// properties: The properties for the new cookie object, expressed as key value pairs.
//
// # Return Value
// 
// The newly created cookie object. Returns `nil` if the provided properties
// are invalid.
//
// # Discussion
// 
// To successfully create a cookie, you must provide values for (at least) the
// [path], [name], and [value] keys, and either the [originURL] key or the
// [domain] key.
// 
// See Accepting cookies for more information on the available cookie
// attribute constants and the constraints imposed on the values in the
// dictionary.
//
// [domain]: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/domain
// [name]: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/name
// [originURL]: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/originURL
// [path]: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/path
// [value]: https://developer.apple.com/documentation/Foundation/HTTPCookiePropertyKey/value
//
// See: https://developer.apple.com/documentation/Foundation/NSHTTPCookie/cookieWithProperties:
func (_HTTPCookieClass HTTPCookieClass) CookieWithProperties(properties INSDictionary) HTTPCookie {
	rv := objc.Send[objc.ID](objc.ID(_HTTPCookieClass.class), objc.Sel("cookieWithProperties:"), properties)
	return NSHTTPCookieFromID(rv)
}








// The domain of the cookie.
//
// # Discussion
// 
// If the domain does not start with a dot, then the cookie is only sent to
// the exact host specified by the domain. If the domain does start with a
// dot, then the cookie is sent to other hosts in that domain as well, subject
// to certain restrictions. See [RFC 6265] for more detail.
//
// [RFC 6265]: https://tools.ietf.org/html/rfc6265.html
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookie/domain
func (h HTTPCookie) Domain() string {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("domain"))
	return NSStringFromID(rv).String()
}



// The cookie’s path.
//
// # Discussion
// 
// The cookie will be sent with requests for this path in the cookie’s
// domain, and all paths that have this prefix. A path of `"/"` means the
// cookie will be sent for all URLs in the domain.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookie/path
func (h HTTPCookie) Path() string {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("path"))
	return NSStringFromID(rv).String()
}



// The cookie’s port list.
//
// # Discussion
// 
// The list of ports for the cookie, returned as an array of [NSNumber]
// objects containing integers. If the cookie has no port list, the value of
// this property is `nil` and the cookie will be sent to any port. Otherwise,
// the cookie is only sent to ports specified in the port list.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookie/portList
func (h HTTPCookie) PortList() []NSNumber {
	rv := objc.Send[[]objc.ID](h.ID, objc.Sel("portList"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSNumber {
		return NSNumberFromID(id)
	})
}



// The cookie’s name.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookie/name
func (h HTTPCookie) Name() string {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("name"))
	return NSStringFromID(rv).String()
}



// The cookie’s string value.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookie/value
func (h HTTPCookie) Value() string {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("value"))
	return NSStringFromID(rv).String()
}



// The cookie’s version.
//
// # Discussion
// 
// Version 0 maps to “old-style” Netscape cookies. Version 1 maps to [RFC
// 6265] cookies.
//
// [RFC 6265]: https://tools.ietf.org/html/rfc6265
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookie/version
func (h HTTPCookie) Version() uint {
	rv := objc.Send[uint](h.ID, objc.Sel("version"))
	return rv
}



// The cookie’s expiration date.
//
// # Discussion
// 
// This value is `nil` if there is no specific expiration date, as with
// session-only cookies. The expiration date is the date when the cookie
// should be deleted.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookie/expiresDate
func (h HTTPCookie) ExpiresDate() INSDate {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("expiresDate"))
	return NSDateFromID(objc.ID(rv))
}



// A Boolean value that indicates whether the cookie should be discarded at
// the end of the session (regardless of expiration date).
//
// # Discussion
// 
// This value is [true] if the cookie should be discarded at the end of the
// session (regardless of expiration date), otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookie/isSessionOnly
func (h HTTPCookie) SessionOnly() bool {
	rv := objc.Send[bool](h.ID, objc.Sel("isSessionOnly"))
	return rv
}



// A Boolean value that indicates whether the cookie should only be sent to
// HTTP servers.
//
// # Discussion
// 
// The value of this property is [true] if the cookie should only be sent
// using HTTP headers, [false] otherwise.
// 
// Cookies can be marked as HTTP-only by a server (or by JavaScript code).
// Cookies marked as such must only be sent via HTTP Headers in HTTP requests
// for URLs that match both the path and domain of the respective cookies.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookie/isHTTPOnly
func (h HTTPCookie) HTTPOnly() bool {
	rv := objc.Send[bool](h.ID, objc.Sel("isHTTPOnly"))
	return rv
}



// A Boolean value that indicates whether the cookie may only be sent over
// secure channels.
//
// # Discussion
// 
// This value is [true] if this cookie should only be sent over secure
// channels, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookie/isSecure
func (h HTTPCookie) Secure() bool {
	rv := objc.Send[bool](h.ID, objc.Sel("isSecure"))
	return rv
}



// A Boolean value that indicates whether to restrict the cookie to requests
// sent back to the same site that created it.
//
// # Discussion
// 
// Along with the policy values defined by [NSHTTPCookieStringPolicy], this
// property may also be `nil`. In this case, cross-site requests include the
// cookie.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookie/sameSitePolicy
func (h HTTPCookie) SameSitePolicy() NSHTTPCookieStringPolicy {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("sameSitePolicy"))
	return HTTPCookieStringPolicy(NSStringFromID(rv).String())
}



// The cookie’s properties.
//
// # Discussion
// 
// This dictionary can be used with [InitWithProperties] (or
// [CookieWithProperties] in Objective-C) to create an equivalent
// [NSHTTPCookie] object.
// 
// See [InitWithProperties] for more information on the constraints imposed on
// the `properties` dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookie/properties
func (h HTTPCookie) Properties() INSDictionary {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("properties"))
	return NSDictionaryFromID(objc.ID(rv))
}



// The cookie’s comment string.
//
// # Discussion
// 
// This value is `nil` if the cookie has no comment. You can present this
// string to the user, explaining the contents and purpose of this cookie.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookie/comment
func (h HTTPCookie) Comment() string {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("comment"))
	return NSStringFromID(rv).String()
}



// The cookie’s comment URL.
//
// # Discussion
// 
// This value is `nil` if the cookie has no comment URL. This value specifies
// a URL that can be presented to the user as a link for further information
// about this cookie.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookie/commentURL
func (h HTTPCookie) CommentURL() INSURL {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("commentURL"))
	return NSURLFromID(objc.ID(rv))
}

























