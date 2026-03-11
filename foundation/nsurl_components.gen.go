// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSURLComponents] class.
var (
	_NSURLComponentsClass     NSURLComponentsClass
	_NSURLComponentsClassOnce sync.Once
)

func getNSURLComponentsClass() NSURLComponentsClass {
	_NSURLComponentsClassOnce.Do(func() {
		_NSURLComponentsClass = NSURLComponentsClass{class: objc.GetClass("NSURLComponents")}
	})
	return _NSURLComponentsClass
}

// GetNSURLComponentsClass returns the class object for NSURLComponents.
func GetNSURLComponentsClass() NSURLComponentsClass {
	return getNSURLComponentsClass()
}

type NSURLComponentsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSURLComponentsClass) Alloc() NSURLComponents {
	rv := objc.Send[NSURLComponents](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An object that parses URLs into and constructs URLs from their constituent
// parts.
//
// # Overview
// 
// In Swift, this object bridges to [URLComponents]; use [NSURLComponents]
// when you need reference semantics or other Foundation-specific behavior.
// 
// The [NSURLComponents] class is a class that is designed to parse URLs based
// on [RFC 3986] and to construct URLs from their constituent parts. Its
// behavior differs subtly from the [NSURL] class, which conforms to older
// RFCs. However, you can easily obtain an [NSURL] object based on the
// contents of a URL components object or vice versa.
// 
// You create a URL components object in one of three ways: from an [NSString]
// object that contains a URL, from an [NSURL] object, or from scratch by
// using the default initializer. From there, you can modify the URL’s
// individual components and subcomponents by modifying various properties,
// either in unencoded form or in URL-encoded form. If you set the unencoded
// property, you can then obtain the encoded equivalent by reading the encoded
// property value and vice versa.
//
// [RFC 3986]: http://www.ietf.org/rfc/rfc3986.txt
// [URLComponents]: https://developer.apple.com/documentation/Foundation/URLComponents
//
// # Creating URL components
//
//   - [NSURLComponents.InitWithString]: Creates a URL components object by parsing a URL in string form.
//   - [NSURLComponents.InitWithStringEncodingInvalidCharacters]: Creates a URL components instance from the provided string, optionally IDNA- and percent-encoding any invalid characters.
//   - [NSURLComponents.InitWithURLResolvingAgainstBaseURL]: Creates a URL components object by parsing the URL from an [NSURL] object.
//
// # Getting the URL
//
//   - [NSURLComponents.String]: A URL derived from the components object, in string form.
//   - [NSURLComponents.URL]: A URL object derived from the components object.
//   - [NSURLComponents.URLRelativeToURL]: Returns a URL object derived from the components object.
//
// # Accessing components in native format
//
//   - [NSURLComponents.Fragment]: The fragment URL component (the part after a `#` symbol), or nil if not present.
//   - [NSURLComponents.SetFragment]
//   - [NSURLComponents.Host]: The host URL subcomponent, or nil if not present.
//   - [NSURLComponents.SetHost]
//   - [NSURLComponents.EncodedHost]: The host subcomponent, percent-encoded.
//   - [NSURLComponents.SetEncodedHost]
//   - [NSURLComponents.Password]: The password URL subcomponent, or nil if not present.
//   - [NSURLComponents.SetPassword]
//   - [NSURLComponents.Path]: The path URL component, or nil if not present.
//   - [NSURLComponents.SetPath]
//   - [NSURLComponents.Port]: The port number URL component, or nil if not present.
//   - [NSURLComponents.SetPort]
//   - [NSURLComponents.Query]: The query URL component as a string, or nil if not present.
//   - [NSURLComponents.SetQuery]
//   - [NSURLComponents.QueryItems]: The query URL component as an array of name/value pairs.
//   - [NSURLComponents.SetQueryItems]
//   - [NSURLComponents.Scheme]: The scheme URL component, or nil if not present.
//   - [NSURLComponents.SetScheme]
//   - [NSURLComponents.User]: The username URL subcomponent, or nil if not present.
//   - [NSURLComponents.SetUser]
//
// # Accessing components in URL-encoded format
//
//   - [NSURLComponents.PercentEncodedFragment]: The fragment URL component (the part after a `#` symbol) expressed as a URL-encoded string, or `nil` if not present.
//   - [NSURLComponents.SetPercentEncodedFragment]
//   - [NSURLComponents.PercentEncodedHost]: The host URL subcomponent expressed as a URL-encoded string, or `nil` if not present.
//   - [NSURLComponents.SetPercentEncodedHost]
//   - [NSURLComponents.PercentEncodedPassword]: The password URL subcomponent expressed as a URL-encoded string, or `nil` if not present.
//   - [NSURLComponents.SetPercentEncodedPassword]
//   - [NSURLComponents.PercentEncodedPath]: The path URL component expressed as a URL-encoded string, or `nil` if not present.
//   - [NSURLComponents.SetPercentEncodedPath]
//   - [NSURLComponents.PercentEncodedQuery]: The query URL component expressed as a URL-encoded string, or `nil` if not present.
//   - [NSURLComponents.SetPercentEncodedQuery]
//   - [NSURLComponents.PercentEncodedUser]: The username URL subcomponent expressed as a URL-encoded string, or `nil` if not present.
//   - [NSURLComponents.SetPercentEncodedUser]
//
// # Locating components in the URL string representation
//
//   - [NSURLComponents.PercentEncodedQueryItems]
//   - [NSURLComponents.SetPercentEncodedQueryItems]
//   - [NSURLComponents.RangeOfFragment]: Returns the character range of the fragment in the string returned by the string property.
//   - [NSURLComponents.RangeOfHost]: Returns the character range of the host in the string returned by the string property.
//   - [NSURLComponents.RangeOfPassword]: Returns the character range of the password in the string returned by the string property.
//   - [NSURLComponents.RangeOfPath]: Returns the character range of the path in the string returned by the string property.
//   - [NSURLComponents.RangeOfPort]: Returns the character range of the port in the string returned by the string property.
//   - [NSURLComponents.RangeOfQuery]: Returns the character range of the query in the string returned by the string property.
//   - [NSURLComponents.RangeOfScheme]: Returns the character range of the scheme in the string returned by the string property.
//   - [NSURLComponents.RangeOfUser]: Returns the character range of the user in the string returned by the string property.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents
type NSURLComponents struct {
	objectivec.Object
}

// NSURLComponentsFromID constructs a [NSURLComponents] from an objc.ID.
//
// An object that parses URLs into and constructs URLs from their constituent
// parts.
func NSURLComponentsFromID(id objc.ID) NSURLComponents {
	return NSURLComponents{objectivec.Object{id}}
}
// NOTE: NSURLComponents adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSURLComponents] class.
//
// # Creating URL components
//
//   - [INSURLComponents.InitWithString]: Creates a URL components object by parsing a URL in string form.
//   - [INSURLComponents.InitWithStringEncodingInvalidCharacters]: Creates a URL components instance from the provided string, optionally IDNA- and percent-encoding any invalid characters.
//   - [INSURLComponents.InitWithURLResolvingAgainstBaseURL]: Creates a URL components object by parsing the URL from an [NSURL] object.
//
// # Getting the URL
//
//   - [INSURLComponents.String]: A URL derived from the components object, in string form.
//   - [INSURLComponents.URL]: A URL object derived from the components object.
//   - [INSURLComponents.URLRelativeToURL]: Returns a URL object derived from the components object.
//
// # Accessing components in native format
//
//   - [INSURLComponents.Fragment]: The fragment URL component (the part after a `#` symbol), or nil if not present.
//   - [INSURLComponents.SetFragment]
//   - [INSURLComponents.Host]: The host URL subcomponent, or nil if not present.
//   - [INSURLComponents.SetHost]
//   - [INSURLComponents.EncodedHost]: The host subcomponent, percent-encoded.
//   - [INSURLComponents.SetEncodedHost]
//   - [INSURLComponents.Password]: The password URL subcomponent, or nil if not present.
//   - [INSURLComponents.SetPassword]
//   - [INSURLComponents.Path]: The path URL component, or nil if not present.
//   - [INSURLComponents.SetPath]
//   - [INSURLComponents.Port]: The port number URL component, or nil if not present.
//   - [INSURLComponents.SetPort]
//   - [INSURLComponents.Query]: The query URL component as a string, or nil if not present.
//   - [INSURLComponents.SetQuery]
//   - [INSURLComponents.QueryItems]: The query URL component as an array of name/value pairs.
//   - [INSURLComponents.SetQueryItems]
//   - [INSURLComponents.Scheme]: The scheme URL component, or nil if not present.
//   - [INSURLComponents.SetScheme]
//   - [INSURLComponents.User]: The username URL subcomponent, or nil if not present.
//   - [INSURLComponents.SetUser]
//
// # Accessing components in URL-encoded format
//
//   - [INSURLComponents.PercentEncodedFragment]: The fragment URL component (the part after a `#` symbol) expressed as a URL-encoded string, or `nil` if not present.
//   - [INSURLComponents.SetPercentEncodedFragment]
//   - [INSURLComponents.PercentEncodedHost]: The host URL subcomponent expressed as a URL-encoded string, or `nil` if not present.
//   - [INSURLComponents.SetPercentEncodedHost]
//   - [INSURLComponents.PercentEncodedPassword]: The password URL subcomponent expressed as a URL-encoded string, or `nil` if not present.
//   - [INSURLComponents.SetPercentEncodedPassword]
//   - [INSURLComponents.PercentEncodedPath]: The path URL component expressed as a URL-encoded string, or `nil` if not present.
//   - [INSURLComponents.SetPercentEncodedPath]
//   - [INSURLComponents.PercentEncodedQuery]: The query URL component expressed as a URL-encoded string, or `nil` if not present.
//   - [INSURLComponents.SetPercentEncodedQuery]
//   - [INSURLComponents.PercentEncodedUser]: The username URL subcomponent expressed as a URL-encoded string, or `nil` if not present.
//   - [INSURLComponents.SetPercentEncodedUser]
//
// # Locating components in the URL string representation
//
//   - [INSURLComponents.PercentEncodedQueryItems]
//   - [INSURLComponents.SetPercentEncodedQueryItems]
//   - [INSURLComponents.RangeOfFragment]: Returns the character range of the fragment in the string returned by the string property.
//   - [INSURLComponents.RangeOfHost]: Returns the character range of the host in the string returned by the string property.
//   - [INSURLComponents.RangeOfPassword]: Returns the character range of the password in the string returned by the string property.
//   - [INSURLComponents.RangeOfPath]: Returns the character range of the path in the string returned by the string property.
//   - [INSURLComponents.RangeOfPort]: Returns the character range of the port in the string returned by the string property.
//   - [INSURLComponents.RangeOfQuery]: Returns the character range of the query in the string returned by the string property.
//   - [INSURLComponents.RangeOfScheme]: Returns the character range of the scheme in the string returned by the string property.
//   - [INSURLComponents.RangeOfUser]: Returns the character range of the user in the string returned by the string property.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents
type INSURLComponents interface {
	objectivec.IObject
	NSCopying

	// Topic: Creating URL components

	// Creates a URL components object by parsing a URL in string form.
	InitWithString(URLString string) NSURLComponents
	// Creates a URL components instance from the provided string, optionally IDNA- and percent-encoding any invalid characters.
	InitWithStringEncodingInvalidCharacters(URLString string, encodingInvalidCharacters bool) NSURLComponents
	// Creates a URL components object by parsing the URL from an [NSURL] object.
	InitWithURLResolvingAgainstBaseURL(url INSURL, resolve bool) NSURLComponents

	// Topic: Getting the URL

	// A URL derived from the components object, in string form.
	String() string
	// A URL object derived from the components object.
	URL() INSURL
	// Returns a URL object derived from the components object.
	URLRelativeToURL(baseURL INSURL) INSURL

	// Topic: Accessing components in native format

	// The fragment URL component (the part after a `#` symbol), or nil if not present.
	Fragment() string
	SetFragment(value string)
	// The host URL subcomponent, or nil if not present.
	Host() string
	SetHost(value string)
	// The host subcomponent, percent-encoded.
	EncodedHost() string
	SetEncodedHost(value string)
	// The password URL subcomponent, or nil if not present.
	Password() string
	SetPassword(value string)
	// The path URL component, or nil if not present.
	Path() string
	SetPath(value string)
	// The port number URL component, or nil if not present.
	Port() INSNumber
	SetPort(value INSNumber)
	// The query URL component as a string, or nil if not present.
	Query() string
	SetQuery(value string)
	// The query URL component as an array of name/value pairs.
	QueryItems() []NSURLQueryItem
	SetQueryItems(value []NSURLQueryItem)
	// The scheme URL component, or nil if not present.
	Scheme() string
	SetScheme(value string)
	// The username URL subcomponent, or nil if not present.
	User() string
	SetUser(value string)

	// Topic: Accessing components in URL-encoded format

	// The fragment URL component (the part after a `#` symbol) expressed as a URL-encoded string, or `nil` if not present.
	PercentEncodedFragment() string
	SetPercentEncodedFragment(value string)
	// The host URL subcomponent expressed as a URL-encoded string, or `nil` if not present.
	PercentEncodedHost() string
	SetPercentEncodedHost(value string)
	// The password URL subcomponent expressed as a URL-encoded string, or `nil` if not present.
	PercentEncodedPassword() string
	SetPercentEncodedPassword(value string)
	// The path URL component expressed as a URL-encoded string, or `nil` if not present.
	PercentEncodedPath() string
	SetPercentEncodedPath(value string)
	// The query URL component expressed as a URL-encoded string, or `nil` if not present.
	PercentEncodedQuery() string
	SetPercentEncodedQuery(value string)
	// The username URL subcomponent expressed as a URL-encoded string, or `nil` if not present.
	PercentEncodedUser() string
	SetPercentEncodedUser(value string)

	// Topic: Locating components in the URL string representation

	PercentEncodedQueryItems() []NSURLQueryItem
	SetPercentEncodedQueryItems(value []NSURLQueryItem)
	// Returns the character range of the fragment in the string returned by the string property.
	RangeOfFragment() NSRange
	// Returns the character range of the host in the string returned by the string property.
	RangeOfHost() NSRange
	// Returns the character range of the password in the string returned by the string property.
	RangeOfPassword() NSRange
	// Returns the character range of the path in the string returned by the string property.
	RangeOfPath() NSRange
	// Returns the character range of the port in the string returned by the string property.
	RangeOfPort() NSRange
	// Returns the character range of the query in the string returned by the string property.
	RangeOfQuery() NSRange
	// Returns the character range of the scheme in the string returned by the string property.
	RangeOfScheme() NSRange
	// Returns the character range of the user in the string returned by the string property.
	RangeOfUser() NSRange
}





// Init initializes the instance.
func (u NSURLComponents) Init() NSURLComponents {
	rv := objc.Send[NSURLComponents](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u NSURLComponents) Autorelease() NSURLComponents {
	rv := objc.Send[NSURLComponents](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSURLComponents creates a new NSURLComponents instance.
func NewNSURLComponents() NSURLComponents {
	class := getNSURLComponentsClass()
	rv := objc.Send[NSURLComponents](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a URL components object by parsing a URL in string form.
//
// URLString: The URL string to parse.
//
// # Return Value
// 
// Returns the initialized URL components object, or `nil` if the URL string
// could not be parsed.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/init(string:)
func NewURLComponentsWithString(URLString string) NSURLComponents {
	instance := getNSURLComponentsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:"), objc.String(URLString))
	return NSURLComponentsFromID(rv)
}


// Creates a URL components instance from the provided string, optionally
// IDNA- and percent-encoding any invalid characters.
//
// URLString: The URL string to parse.
//
// encodingInvalidCharacters: A Boolean value that indicates whether the initializer attempts to encode
// any invalid characters in [URLString].
//
// # Discussion
// 
// If `encodingInvalidCharacters` is `true`, this initializer tries to encode
// the string to create a valid URL. If the URL string is still invalid after
// encoding, the initializer returns `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/init(string:encodingInvalidCharacters:)
func NewURLComponentsWithStringEncodingInvalidCharacters(URLString string, encodingInvalidCharacters bool) NSURLComponents {
	instance := getNSURLComponentsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:encodingInvalidCharacters:"), objc.String(URLString), encodingInvalidCharacters)
	return NSURLComponentsFromID(rv)
}


// Creates a URL components object by parsing the URL from an [NSURL] object.
//
// url: The URL to parse.
//
// resolve: Controls whether the URL should be resolved against its base URL before
// parsing. If [true], and if the `url` parameter contains a relative URL, the
// original URL is resolved against its base URL before parsing by calling the
// [AbsoluteURL] method. Otherwise, the string portion is used by itself.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// Returns the initialized URL components object, or `nil` if the URL could
// not be parsed.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/init(url:resolvingAgainstBaseURL:)
func NewURLComponentsWithURLResolvingAgainstBaseURL(url INSURL, resolve bool) NSURLComponents {
	instance := getNSURLComponentsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:resolvingAgainstBaseURL:"), url, resolve)
	return NSURLComponentsFromID(rv)
}







// Creates a URL components object by parsing a URL in string form.
//
// URLString: The URL string to parse.
//
// # Return Value
// 
// Returns the initialized URL components object, or `nil` if the URL string
// could not be parsed.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/init(string:)
func (u NSURLComponents) InitWithString(URLString string) NSURLComponents {
	rv := objc.Send[NSURLComponents](u.ID, objc.Sel("initWithString:"), objc.String(URLString))
	return rv
}

// Creates a URL components instance from the provided string, optionally
// IDNA- and percent-encoding any invalid characters.
//
// URLString: The URL string to parse.
//
// encodingInvalidCharacters: A Boolean value that indicates whether the initializer attempts to encode
// any invalid characters in [URLString].
//
// # Discussion
// 
// If `encodingInvalidCharacters` is `true`, this initializer tries to encode
// the string to create a valid URL. If the URL string is still invalid after
// encoding, the initializer returns `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/init(string:encodingInvalidCharacters:)
func (u NSURLComponents) InitWithStringEncodingInvalidCharacters(URLString string, encodingInvalidCharacters bool) NSURLComponents {
	rv := objc.Send[NSURLComponents](u.ID, objc.Sel("initWithString:encodingInvalidCharacters:"), objc.String(URLString), encodingInvalidCharacters)
	return rv
}

// Creates a URL components object by parsing the URL from an [NSURL] object.
//
// url: The URL to parse.
//
// resolve: Controls whether the URL should be resolved against its base URL before
// parsing. If [true], and if the `url` parameter contains a relative URL, the
// original URL is resolved against its base URL before parsing by calling the
// [AbsoluteURL] method. Otherwise, the string portion is used by itself.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// Returns the initialized URL components object, or `nil` if the URL could
// not be parsed.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/init(url:resolvingAgainstBaseURL:)
func (u NSURLComponents) InitWithURLResolvingAgainstBaseURL(url INSURL, resolve bool) NSURLComponents {
	rv := objc.Send[NSURLComponents](u.ID, objc.Sel("initWithURL:resolvingAgainstBaseURL:"), url, resolve)
	return rv
}

// Returns a URL object derived from the components object.
//
// baseURL: If non-`nil`, this URL is used as the base URL portion of the resulting URL
// object.
//
// # Discussion
// 
// If the components object has an authority component (user, password, host,
// or port) and a path component, then the path must either begin with `"/"`
// or be an empty string. Otherwise, this property contains `nil`.
// 
// If the [NSURLComponents] have an authority component (user, password, host,
// or port) and has a path component, the path component must not start with
// `"//"`. If it does, this property contains `nil`.
// 
// To configure a components object based on an existing URL, call either the
// [ComponentsWithURLResolvingAgainstBaseURL] or
// [InitWithURLResolvingAgainstBaseURL] method.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/url(relativeTo:)
func (u NSURLComponents) URLRelativeToURL(baseURL INSURL) INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("URLRelativeToURL:"), baseURL)
	return NSURLFromID(rv)
}





// Returns a URL components object by parsing a URL in string form.
//
// URLString: The URL string to parse.
//
// # Return Value
// 
// Returns the new URL components object, or `nil` if the URL string could not
// be parsed.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/componentsWithString:
func (_NSURLComponentsClass NSURLComponentsClass) ComponentsWithString(URLString string) NSURLComponents {
	rv := objc.Send[objc.ID](objc.ID(_NSURLComponentsClass.class), objc.Sel("componentsWithString:"), objc.String(URLString))
	return NSURLComponentsFromID(rv)
}

// Returns a URL components instance from the provided string, optionally
// IDNA- and percent-encoding any invalid characters.
//
// URLString: The URL string to parse.
//
// encodingInvalidCharacters: A Boolean value that indicates whether the initializer attempts to encode
// any invalid characters in [URLString].
//
// # Return Value
// 
// A URL components instance from the provided string, optionally with invalid
// characters percent-encoded.
//
// # Discussion
// 
// If `encodingInvalidCharacters` is `true`, this initializer tries to encode
// the string to create a valid URL. If the URL string is still invalid after
// encoding, the method returns `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/componentsWithString:encodingInvalidCharacters:
func (_NSURLComponentsClass NSURLComponentsClass) ComponentsWithStringEncodingInvalidCharacters(URLString string, encodingInvalidCharacters bool) NSURLComponents {
	rv := objc.Send[objc.ID](objc.ID(_NSURLComponentsClass.class), objc.Sel("componentsWithString:encodingInvalidCharacters:"), objc.String(URLString), encodingInvalidCharacters)
	return NSURLComponentsFromID(rv)
}

// Returns a URL components object by parsing the URL from an [NSURL] object.
//
// url: The URL to parse.
//
// resolve: Controls whether the URL should be resolved against its base URL before
// parsing. If [true], and if the `url` parameter contains a relative URL, the
// original URL is resolved against its base URL before parsing by calling the
// [AbsoluteURL] method. Otherwise, the string portion is used by itself.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// Returns the new URL components object, or `nil` if the URL could not be
// parsed.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/componentsWithURL:resolvingAgainstBaseURL:
func (_NSURLComponentsClass NSURLComponentsClass) ComponentsWithURLResolvingAgainstBaseURL(url INSURL, resolve bool) NSURLComponents {
	rv := objc.Send[objc.ID](objc.ID(_NSURLComponentsClass.class), objc.Sel("componentsWithURL:resolvingAgainstBaseURL:"), url, resolve)
	return NSURLComponentsFromID(rv)
}








// A URL derived from the components object, in string form.
//
// # Discussion
// 
// If the receiver has an authority component (user, password, host, or port)
// and a path component, then the path must either begin with `"/"` or be an
// empty string. Otherwise, this property contains `nil`.
// 
// If the receiver have an authority component (user, password, host, or port)
// and has a path component, the path component must not start with `"//"`. If
// it does, this property contains `nil`.
// 
// This property can be used only to obtain a URL string based on the values
// of the other properties. To configure a components object based on an
// existing URL string, call either the [ComponentsWithString] or
// [InitWithString] method.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/string
func (u NSURLComponents) String() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("string"))
	return NSStringFromID(rv).String()
}



// A URL object derived from the components object.
//
// # Discussion
// 
// If the receiver has an authority component (user, password, host, or port)
// and a path component, then the path must either begin with `"/"` or be an
// empty string. Otherwise, this property contains `nil`.
// 
// If the receiver have an authority component (user, password, host, or port)
// and has a path component, the path component must not start with `"//"`. If
// it does, this property contains `nil`.
// 
// If the receiver has `nil` values for all component properties, such as when
// initializing with [Init], this property returns an [NSURL] object with an
// empty string, because a URL always has a path—even if it’s an empty
// string.
// 
// This property can be used only to obtain a URL based on the values of the
// other properties. To configure a components object based on an existing
// URL, call either the [ComponentsWithURLResolvingAgainstBaseURL] or
// [InitWithURLResolvingAgainstBaseURL] method.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/url
func (u NSURLComponents) URL() INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("URL"))
	return NSURLFromID(objc.ID(rv))
}



// The fragment URL component (the part after a `#` symbol), or nil if not
// present.
//
// # Discussion
// 
// For example, in the URL
// `//www.ExampleXCUIElementTypeCom()/index.Html()#jumpLocation`, the fragment
// is `jumpLocation`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/fragment
func (u NSURLComponents) Fragment() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("fragment"))
	return NSStringFromID(rv).String()
}
func (u NSURLComponents) SetFragment(value string) {
	objc.Send[struct{}](u.ID, objc.Sel("setFragment:"), objc.String(value))
}



// The host URL subcomponent, or nil if not present.
//
// # Discussion
// 
// For example, in the URL `//www.ExampleXCUIElementTypeCom()/index.Html()`,
// the host is `www.ExampleXCUIElementTypeCom()`.
// 
// The getter for this property removes any percent encoding this component
// may have (if the component allows percent encoding). Setting this property
// assumes the subcomponent or component string isn’t percent encoded and
// adds percent encoding (if the component allows percent encoding).
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/host
func (u NSURLComponents) Host() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("host"))
	return NSStringFromID(rv).String()
}
func (u NSURLComponents) SetHost(value string) {
	objc.Send[struct{}](u.ID, objc.Sel("setHost:"), objc.String(value))
}



// The host subcomponent, percent-encoded.
//
// # Discussion
// 
// The getter for this property retains any percent-encoding this component
// may have. Setting this property assumes the component string already has
// the correct percent-encoding. Attempting to set an incorrectly
// percent-encoded string raises [fatalError(_:file:line:)].
//
// [fatalError(_:file:line:)]: https://developer.apple.com/documentation/Swift/fatalError(_:file:line:)
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/encodedHost
func (u NSURLComponents) EncodedHost() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("encodedHost"))
	return NSStringFromID(rv).String()
}
func (u NSURLComponents) SetEncodedHost(value string) {
	objc.Send[struct{}](u.ID, objc.Sel("setEncodedHost:"), objc.String(value))
}



// The password URL subcomponent, or nil if not present.
//
// # Discussion
// 
// For example, in the URL
// `//password@www.ExampleXCUIElementTypeCom()/index.Html()`, the password is
// `password`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/password
func (u NSURLComponents) Password() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("password"))
	return NSStringFromID(rv).String()
}
func (u NSURLComponents) SetPassword(value string) {
	objc.Send[struct{}](u.ID, objc.Sel("setPassword:"), objc.String(value))
}



// The path URL component, or nil if not present.
//
// # Discussion
// 
// For example, in the URL `//www.ExampleXCUIElementTypeCom()/index.Html()`,
// the path is `/index.Html()`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/path
func (u NSURLComponents) Path() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("path"))
	return NSStringFromID(rv).String()
}
func (u NSURLComponents) SetPath(value string) {
	objc.Send[struct{}](u.ID, objc.Sel("setPath:"), objc.String(value))
}



// The port number URL component, or nil if not present.
//
// # Discussion
// 
// For example, in the URL `//www.Example().8080/index.Php()`, the port number
// is `8080`.
// 
// If you attempt to set the port to a negative port number, this class throws
// an exception.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/port
func (u NSURLComponents) Port() INSNumber {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("port"))
	return NSNumberFromID(objc.ID(rv))
}
func (u NSURLComponents) SetPort(value INSNumber) {
	objc.Send[struct{}](u.ID, objc.Sel("setPort:"), value)
}



// The query URL component as a string, or nil if not present.
//
// # Discussion
// 
// For example, in the URL
// `//www.ExampleXCUIElementTypeCom()/index.Php()?key1=value1&key2=value2`,
// the query string is `key1=value1&key2=value2`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/query
func (u NSURLComponents) Query() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("query"))
	return NSStringFromID(rv).String()
}
func (u NSURLComponents) SetQuery(value string) {
	objc.Send[struct{}](u.ID, objc.Sel("setQuery:"), objc.String(value))
}



// The query URL component as an array of name/value pairs.
//
// # Discussion
// 
// When you get this property’s value, the [NSURLComponents] class parses
// the [Query] string and returns an array of [NSURLQueryItem] objects, each
// of which represents a single key-value pair, in the order in which they
// appear in the original query string. Because a name may appear more than
// once in a single query string, the [Name] properties of query items are not
// guaranteed to be unique. If the [Query] property is an empty string, the
// [QueryItems] property is an empty array. If the [Query] property is `nil`,
// the [QueryItems] property is also `nil`.
// 
// When you set this property’s value, the [NSURLComponents] class joins
// each name/value pair with a `=` delimiter and joins the array with a `&`
// delimiter, then sets the [Query] property to the resulting string. Setting
// the [QueryItems] property to an empty array sets the [Query] property to an
// empty string, and setting the [QueryItems] property to `nil` sets the
// [Query] property to `nil`.
// 
// To ensure you can compose and decompose URL queries even with empty
// components, the [NSURLComponents] class has the following behavior for
// cases where no name or value is present:
// 
// - If a name-value pair has nothing before its equals sign, the `name`
// property of the corresponding query item is a zero-length string. - If a
// name-value pair has nothing after its equals sign, the `value` property of
// the corresponding query item is a zero-length string. - If a name-value
// pair has no equals sign, the `value` property of the corresponding query
// item is `nil`. - If a name-value pair is empty (that is, the `query` string
// starts with `&`, ends with `&`, or contains `&``&`), the corresponding
// query item has a zero-length `name` and `nil` `value`.
// 
// For example, in the URL
// `//www.ExampleXCUIElementTypeCom()/index.Php()?key1=value1&key2=value2`,
// this property’s value is an array of two [NSURLQueryItem] objects: one
// whose name property is `key1` and whose value property is `value1`, and one
// whose name property is `key2` and whose value property is `value2`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/queryItems
func (u NSURLComponents) QueryItems() []NSURLQueryItem {
	rv := objc.Send[[]objc.ID](u.ID, objc.Sel("queryItems"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSURLQueryItem {
		return NSURLQueryItemFromID(id)
	})
}
func (u NSURLComponents) SetQueryItems(value []NSURLQueryItem) {
	objc.Send[struct{}](u.ID, objc.Sel("setQueryItems:"), objectivec.IObjectSliceToNSArray(value))
}



// The scheme URL component, or nil if not present.
//
// # Discussion
// 
// For example, in the URL `//www.ExampleXCUIElementTypeCom()/index.Html()`,
// the scheme is `http`.
// 
// If you attempt to set the scheme to an invalid scheme string, this class
// throws an exception.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/scheme
func (u NSURLComponents) Scheme() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("scheme"))
	return NSStringFromID(rv).String()
}
func (u NSURLComponents) SetScheme(value string) {
	objc.Send[struct{}](u.ID, objc.Sel("setScheme:"), objc.String(value))
}



// The username URL subcomponent, or nil if not present.
//
// # Discussion
// 
// For example, in the URL
// `//password@www.ExampleXCUIElementTypeCom()/index.Html()`, the user is
// `username`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/user
func (u NSURLComponents) User() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("user"))
	return NSStringFromID(rv).String()
}
func (u NSURLComponents) SetUser(value string) {
	objc.Send[struct{}](u.ID, objc.Sel("setUser:"), objc.String(value))
}



// The fragment URL component (the part after a `#` symbol) expressed as a
// URL-encoded string, or `nil` if not present.
//
// # Discussion
// 
// For example, in the URL
// `//www.ExampleXCUIElementTypeCom()/index.Html()#jumpLocation`, the fragment
// is `jumpLocation`.
// 
// If you set this value to something that is not a valid, percent-encoded
// string, this class throws an exception.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/percentEncodedFragment
func (u NSURLComponents) PercentEncodedFragment() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("percentEncodedFragment"))
	return NSStringFromID(rv).String()
}
func (u NSURLComponents) SetPercentEncodedFragment(value string) {
	objc.Send[struct{}](u.ID, objc.Sel("setPercentEncodedFragment:"), objc.String(value))
}



// The host URL subcomponent expressed as a URL-encoded string, or `nil` if
// not present.
//
// # Discussion
// 
// For example, in the URL `//www.ExampleXCUIElementTypeCom()/index.Html()`,
// the host is `www.ExampleXCUIElementTypeCom()`.
// 
// If you set this value to something that is not a valid, percent-encoded
// string, this class throws an exception.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/percentEncodedHost
func (u NSURLComponents) PercentEncodedHost() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("percentEncodedHost"))
	return NSStringFromID(rv).String()
}
func (u NSURLComponents) SetPercentEncodedHost(value string) {
	objc.Send[struct{}](u.ID, objc.Sel("setPercentEncodedHost:"), objc.String(value))
}



// The password URL subcomponent expressed as a URL-encoded string, or `nil`
// if not present.
//
// # Discussion
// 
// For example, in the URL
// `//password@www.ExampleXCUIElementTypeCom()/index.Html()`, the password is
// `password`.
// 
// If you set this value to something that is not a valid, percent-encoded
// string, this class throws an exception.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/percentEncodedPassword
func (u NSURLComponents) PercentEncodedPassword() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("percentEncodedPassword"))
	return NSStringFromID(rv).String()
}
func (u NSURLComponents) SetPercentEncodedPassword(value string) {
	objc.Send[struct{}](u.ID, objc.Sel("setPercentEncodedPassword:"), objc.String(value))
}



// The path URL component expressed as a URL-encoded string, or `nil` if not
// present.
//
// # Discussion
// 
// For example, in the URL `//www.ExampleXCUIElementTypeCom()/index.Html()`,
// the path is `/index.Html()`.
// 
// If you set this value to something that is not a valid, percent-encoded
// string, this class throws an exception.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/percentEncodedPath
func (u NSURLComponents) PercentEncodedPath() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("percentEncodedPath"))
	return NSStringFromID(rv).String()
}
func (u NSURLComponents) SetPercentEncodedPath(value string) {
	objc.Send[struct{}](u.ID, objc.Sel("setPercentEncodedPath:"), objc.String(value))
}



// The query URL component expressed as a URL-encoded string, or `nil` if not
// present.
//
// # Discussion
// 
// For example, in the URL
// `//www.ExampleXCUIElementTypeCom()/index.Php()?key1=value1&key2=value2`,
// the query string is `key1=value1&key2=value2`.
// 
// If you set this value to something that is not a valid, percent-encoded
// string, this class throws an exception.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/percentEncodedQuery
func (u NSURLComponents) PercentEncodedQuery() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("percentEncodedQuery"))
	return NSStringFromID(rv).String()
}
func (u NSURLComponents) SetPercentEncodedQuery(value string) {
	objc.Send[struct{}](u.ID, objc.Sel("setPercentEncodedQuery:"), objc.String(value))
}



// The username URL subcomponent expressed as a URL-encoded string, or `nil`
// if not present.
//
// # Discussion
// 
// For example, in the URL
// `//password@www.ExampleXCUIElementTypeCom()/index.Html()`, the user is
// `username`.
// 
// If you set this value to something that is not a valid, percent-encoded
// string, this class throws an exception.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/percentEncodedUser
func (u NSURLComponents) PercentEncodedUser() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("percentEncodedUser"))
	return NSStringFromID(rv).String()
}
func (u NSURLComponents) SetPercentEncodedUser(value string) {
	objc.Send[struct{}](u.ID, objc.Sel("setPercentEncodedUser:"), objc.String(value))
}



// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/percentEncodedQueryItems
func (u NSURLComponents) PercentEncodedQueryItems() []NSURLQueryItem {
	rv := objc.Send[[]objc.ID](u.ID, objc.Sel("percentEncodedQueryItems"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSURLQueryItem {
		return NSURLQueryItemFromID(id)
	})
}
func (u NSURLComponents) SetPercentEncodedQueryItems(value []NSURLQueryItem) {
	objc.Send[struct{}](u.ID, objc.Sel("setPercentEncodedQueryItems:"), objectivec.IObjectSliceToNSArray(value))
}



// Returns the character range of the fragment in the string returned by the
// string property.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/rangeOfFragment
func (u NSURLComponents) RangeOfFragment() NSRange {
	rv := objc.Send[NSRange](u.ID, objc.Sel("rangeOfFragment"))
	return NSRange(rv)
}



// Returns the character range of the host in the string returned by the
// string property.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/rangeOfHost
func (u NSURLComponents) RangeOfHost() NSRange {
	rv := objc.Send[NSRange](u.ID, objc.Sel("rangeOfHost"))
	return NSRange(rv)
}



// Returns the character range of the password in the string returned by the
// string property.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/rangeOfPassword
func (u NSURLComponents) RangeOfPassword() NSRange {
	rv := objc.Send[NSRange](u.ID, objc.Sel("rangeOfPassword"))
	return NSRange(rv)
}



// Returns the character range of the path in the string returned by the
// string property.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/rangeOfPath
func (u NSURLComponents) RangeOfPath() NSRange {
	rv := objc.Send[NSRange](u.ID, objc.Sel("rangeOfPath"))
	return NSRange(rv)
}



// Returns the character range of the port in the string returned by the
// string property.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/rangeOfPort
func (u NSURLComponents) RangeOfPort() NSRange {
	rv := objc.Send[NSRange](u.ID, objc.Sel("rangeOfPort"))
	return NSRange(rv)
}



// Returns the character range of the query in the string returned by the
// string property.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/rangeOfQuery
func (u NSURLComponents) RangeOfQuery() NSRange {
	rv := objc.Send[NSRange](u.ID, objc.Sel("rangeOfQuery"))
	return NSRange(rv)
}



// Returns the character range of the scheme in the string returned by the
// string property.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/rangeOfScheme
func (u NSURLComponents) RangeOfScheme() NSRange {
	rv := objc.Send[NSRange](u.ID, objc.Sel("rangeOfScheme"))
	return NSRange(rv)
}



// Returns the character range of the user in the string returned by the
// string property.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLComponents/rangeOfUser
func (u NSURLComponents) RangeOfUser() NSRange {
	rv := objc.Send[NSRange](u.ID, objc.Sel("rangeOfUser"))
	return NSRange(rv)
}















			// Protocol methods for NSCopying
			














