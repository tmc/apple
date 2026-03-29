// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [HTTPCookieStorage] class.
var (
	_HTTPCookieStorageClass     HTTPCookieStorageClass
	_HTTPCookieStorageClassOnce sync.Once
)

func getHTTPCookieStorageClass() HTTPCookieStorageClass {
	_HTTPCookieStorageClassOnce.Do(func() {
		_HTTPCookieStorageClass = HTTPCookieStorageClass{class: objc.GetClass("NSHTTPCookieStorage")}
	})
	return _HTTPCookieStorageClass
}

// GetHTTPCookieStorageClass returns the class object for NSHTTPCookieStorage.
func GetHTTPCookieStorageClass() HTTPCookieStorageClass {
	return getHTTPCookieStorageClass()
}

type HTTPCookieStorageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (hc HTTPCookieStorageClass) Class() objc.Class {
	return hc.class
}

// Alloc allocates memory for a new instance of the class.
func (hc HTTPCookieStorageClass) Alloc() HTTPCookieStorage {
	rv := objc.Send[HTTPCookieStorage](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// A container that manages the storage of cookies.
//
// # Overview
// 
// Each stored cookie is represented by an instance of the [NSHTTPCookie]
// class.
// 
// # Sharing cookie storage
// 
// The persistent cookie storage returned by [SharedHTTPCookieStorage] may be
// available to app extensions or other apps, subject to the following
// guidelines:
// 
// - iOS — Each app and app extension has a unique data container, meaning
// they have separate cookie stores. You can obtain a common cookie storage by
// using the [SharedCookieStorageForGroupContainerIdentifier] method. - macOS
// (non-sandboxed) — As of macOS 10.11, each app has its own cookie storage.
// Prior to macOS 10.11, a common cookie store is shared among the user’s
// apps. - macOS (sandboxed) — Same as iOS. - [UIWebView] — [UIWebView]
// instances within an app inherit the parent app’s shared cookie storage. -
// [WKWebView] — Each [WKWebView] instance has its own cookie storage. See
// the [WKHTTPCookieStore] class for more information.
// 
// Session cookies (where the cookie object’s [SessionOnly] property is
// [true]) are local to a single process and are not shared.
// 
// # Subclassing notes
// 
// The [NSHTTPCookieStorage] class is usable as-is, but you can subclass it.
// For example, you can override the storage methods like
// [StoreCookiesForTask], [GetCookiesForTaskCompletionHandler] to screen which
// cookies are stored, or reimplement the storage mechanism for security or
// other reasons.
// 
// When overriding methods of this class, be aware that methods that take a
// `task` parameter are preferred by the system to equivalent methods that do
// not. Therefore, you should override the task-based methods when
// subclassing, as follows:
// 
// - Retrieving cookies — Override [GetCookiesForTaskCompletionHandler],
// instead of or in addition to [CookiesForURL]. - Adding cookies — Override
// [StoreCookiesForTask], instead of or in addition to
// [SetCookiesForURLMainDocumentURL].
//
// [UIWebView]: https://developer.apple.com/documentation/UIKit/UIWebView
// [WKHTTPCookieStore]: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore
// [WKWebView]: https://developer.apple.com/documentation/WebKit/WKWebView
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Getting and setting the cookie accept policy
//
//   - [HTTPCookieStorage.CookieAcceptPolicy]: The cookie storage’s cookie accept policy.
//   - [HTTPCookieStorage.SetCookieAcceptPolicy]
//
// # Adding and removing cookies
//
//   - [HTTPCookieStorage.RemoveCookiesSinceDate]: Removes cookies that were stored after a given date.
//   - [HTTPCookieStorage.DeleteCookie]: Deletes the specified cookie from the cookie storage.
//   - [HTTPCookieStorage.SetCookie]: Stores a specified cookie in the cookie storage if the cookie accept policy permits.
//   - [HTTPCookieStorage.SetCookiesForURLMainDocumentURL]: Adds an array of cookies to the cookie storage if the storage’s cookie acceptance policy permits.
//   - [HTTPCookieStorage.StoreCookiesForTask]: Stores an array of cookies in the cookie storage, on behalf of the provided task, if the cookie accept policy permits.
//
// # Retrieving cookies
//
//   - [HTTPCookieStorage.Cookies]: The cookie storage’s cookies.
//   - [HTTPCookieStorage.CookiesForURL]: Returns all the cookie storage’s cookies that are sent to a specified URL.
//   - [HTTPCookieStorage.SortedCookiesUsingDescriptors]: Returns all of the cookie storage’s cookies, sorted according to a given set of sort descriptors.
//
// # Tracking cookie storage changes
//
//   - [HTTPCookieStorage.NSHTTPCookieManagerCookiesChanged]: A notification posted when the cookies stored in the cookie storage have changed.
//   - [HTTPCookieStorage.NSHTTPCookieManagerAcceptPolicyChanged]: A notification posted when the acceptance policy of the cookie storage has changed.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage
type HTTPCookieStorage struct {
	objectivec.Object
}

// HTTPCookieStorageFromID constructs a [HTTPCookieStorage] from an objc.ID.
//
// A container that manages the storage of cookies.
func HTTPCookieStorageFromID(id objc.ID) HTTPCookieStorage {
	return NSHTTPCookieStorage{objectivec.Object{ID: id}}
}

// NSHTTPCookieStorageFromID is an alias for [HTTPCookieStorageFromID] for cross-framework compatibility.
func NSHTTPCookieStorageFromID(id objc.ID) HTTPCookieStorage { return HTTPCookieStorageFromID(id) }
// NOTE: HTTPCookieStorage adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [HTTPCookieStorage] class.
//
// # Getting and setting the cookie accept policy
//
//   - [IHTTPCookieStorage.CookieAcceptPolicy]: The cookie storage’s cookie accept policy.
//   - [IHTTPCookieStorage.SetCookieAcceptPolicy]
//
// # Adding and removing cookies
//
//   - [IHTTPCookieStorage.RemoveCookiesSinceDate]: Removes cookies that were stored after a given date.
//   - [IHTTPCookieStorage.DeleteCookie]: Deletes the specified cookie from the cookie storage.
//   - [IHTTPCookieStorage.SetCookie]: Stores a specified cookie in the cookie storage if the cookie accept policy permits.
//   - [IHTTPCookieStorage.SetCookiesForURLMainDocumentURL]: Adds an array of cookies to the cookie storage if the storage’s cookie acceptance policy permits.
//   - [IHTTPCookieStorage.StoreCookiesForTask]: Stores an array of cookies in the cookie storage, on behalf of the provided task, if the cookie accept policy permits.
//
// # Retrieving cookies
//
//   - [IHTTPCookieStorage.Cookies]: The cookie storage’s cookies.
//   - [IHTTPCookieStorage.CookiesForURL]: Returns all the cookie storage’s cookies that are sent to a specified URL.
//   - [IHTTPCookieStorage.SortedCookiesUsingDescriptors]: Returns all of the cookie storage’s cookies, sorted according to a given set of sort descriptors.
//
// # Tracking cookie storage changes
//
//   - [IHTTPCookieStorage.NSHTTPCookieManagerCookiesChanged]: A notification posted when the cookies stored in the cookie storage have changed.
//   - [IHTTPCookieStorage.NSHTTPCookieManagerAcceptPolicyChanged]: A notification posted when the acceptance policy of the cookie storage has changed.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage
type IHTTPCookieStorage interface {
	objectivec.IObject

	// Topic: Getting and setting the cookie accept policy

	// The cookie storage’s cookie accept policy.
	CookieAcceptPolicy() NSHTTPCookieAcceptPolicy
	SetCookieAcceptPolicy(value NSHTTPCookieAcceptPolicy)

	// Topic: Adding and removing cookies

	// Removes cookies that were stored after a given date.
	RemoveCookiesSinceDate(date INSDate)
	// Deletes the specified cookie from the cookie storage.
	DeleteCookie(cookie INSHTTPCookie)
	// Stores a specified cookie in the cookie storage if the cookie accept policy permits.
	SetCookie(cookie INSHTTPCookie)
	// Adds an array of cookies to the cookie storage if the storage’s cookie acceptance policy permits.
	SetCookiesForURLMainDocumentURL(cookies []NSHTTPCookie, URL INSURL, mainDocumentURL INSURL)
	// Stores an array of cookies in the cookie storage, on behalf of the provided task, if the cookie accept policy permits.
	StoreCookiesForTask(cookies []NSHTTPCookie, task INSURLSessionTask)

	// Topic: Retrieving cookies

	// The cookie storage’s cookies.
	Cookies() []NSHTTPCookie
	// Returns all the cookie storage’s cookies that are sent to a specified URL.
	CookiesForURL(URL INSURL) []NSHTTPCookie
	// Returns all of the cookie storage’s cookies, sorted according to a given set of sort descriptors.
	SortedCookiesUsingDescriptors(sortOrder []NSSortDescriptor) []NSHTTPCookie

	// Topic: Tracking cookie storage changes

	// A notification posted when the cookies stored in the cookie storage have changed.
	NSHTTPCookieManagerCookiesChanged() NSNotificationName
	// A notification posted when the acceptance policy of the cookie storage has changed.
	NSHTTPCookieManagerAcceptPolicyChanged() NSNotificationName

	// A Boolean value that indicates whether the cookie should be discarded at the end of the session (regardless of expiration date).
	IsSessionOnly() bool
	SetIsSessionOnly(value bool)
}

// Init initializes the instance.
func (h HTTPCookieStorage) Init() HTTPCookieStorage {
	rv := objc.Send[HTTPCookieStorage](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h HTTPCookieStorage) Autorelease() HTTPCookieStorage {
	rv := objc.Send[HTTPCookieStorage](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewHTTPCookieStorage creates a new HTTPCookieStorage instance.
func NewHTTPCookieStorage() HTTPCookieStorage {
	class := getHTTPCookieStorageClass()
	rv := objc.Send[HTTPCookieStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Removes cookies that were stored after a given date.
//
// date: The date after which cookies should be removed.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/removeCookies(since:)
func (h HTTPCookieStorage) RemoveCookiesSinceDate(date INSDate) {
	objc.Send[objc.ID](h.ID, objc.Sel("removeCookiesSinceDate:"), date)
}
// Deletes the specified cookie from the cookie storage.
//
// cookie: The cookie to delete.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/deleteCookie(_:)
func (h HTTPCookieStorage) DeleteCookie(cookie INSHTTPCookie) {
	objc.Send[objc.ID](h.ID, objc.Sel("deleteCookie:"), cookie)
}
// Stores a specified cookie in the cookie storage if the cookie accept policy
// permits.
//
// cookie: The cookie to store.
//
// # Discussion
// 
// The cookie replaces an existing cookie with the same name, domain, and
// path, if one exists in the cookie storage. This method accepts the cookie
// only if the storage’s cookie accept policy is
// [HTTPCookie.AcceptPolicy.always] or
// [HTTPCookie.AcceptPolicy.onlyFromMainDocumentDomain]. The cookie is ignored
// if the storage’s cookie accept policy is [HTTPCookie.AcceptPolicy.never].
//
// [HTTPCookie.AcceptPolicy.always]: https://developer.apple.com/documentation/Foundation/HTTPCookie/AcceptPolicy/always
// [HTTPCookie.AcceptPolicy.never]: https://developer.apple.com/documentation/Foundation/HTTPCookie/AcceptPolicy/never
// [HTTPCookie.AcceptPolicy.onlyFromMainDocumentDomain]: https://developer.apple.com/documentation/Foundation/HTTPCookie/AcceptPolicy/onlyFromMainDocumentDomain
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/setCookie(_:)
func (h HTTPCookieStorage) SetCookie(cookie INSHTTPCookie) {
	objc.Send[objc.ID](h.ID, objc.Sel("setCookie:"), cookie)
}
// Adds an array of cookies to the cookie storage if the storage’s cookie
// acceptance policy permits.
//
// cookies: The cookies to add.
//
// URL: The URL associated with the added cookies.
//
// mainDocumentURL: The URL of the main HTML document for the top-level frame, if known. The
// value can be `nil`. This URL is used to determine whether the cookie should
// be accepted if the cookie accept policy is
// [HTTPCookie.AcceptPolicy.onlyFromMainDocumentDomain].
// //
// [HTTPCookie.AcceptPolicy.onlyFromMainDocumentDomain]: https://developer.apple.com/documentation/Foundation/HTTPCookie/AcceptPolicy/onlyFromMainDocumentDomain
//
// # Discussion
// 
// Cookies in the array will replace existing cookies with the same name,
// domain, and path in the cookie storage. If the storage has an accept policy
// of [HTTPCookie.AcceptPolicy.never], the cookies are ignored.
// 
// To store cookies from a set of response headers, an application can use
// [CookiesWithResponseHeaderFieldsForURL] passing a header field dictionary
// and then use this method to store the resulting cookies in accordance with
// the cookie storage’s cookie acceptance policy.
// 
// If you override this method, also override [StoreCookiesForTask].
//
// [HTTPCookie.AcceptPolicy.never]: https://developer.apple.com/documentation/Foundation/HTTPCookie/AcceptPolicy/never
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/setCookies(_:for:mainDocumentURL:)
func (h HTTPCookieStorage) SetCookiesForURLMainDocumentURL(cookies []NSHTTPCookie, URL INSURL, mainDocumentURL INSURL) {
	objc.Send[objc.ID](h.ID, objc.Sel("setCookies:forURL:mainDocumentURL:"), objectivec.IObjectSliceToNSArray(cookies), URL, mainDocumentURL)
}
// Stores an array of cookies in the cookie storage, on behalf of the provided
// task, if the cookie accept policy permits.
//
// cookies: The cookies to add.
//
// task: The task that handles the response. Override this method and inspect this
// parameter if you need to alter your cookie storage strategy based on
// properties of the task.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/storeCookies(_:for:)
func (h HTTPCookieStorage) StoreCookiesForTask(cookies []NSHTTPCookie, task INSURLSessionTask) {
	objc.Send[objc.ID](h.ID, objc.Sel("storeCookies:forTask:"), objectivec.IObjectSliceToNSArray(cookies), task)
}
// Returns all the cookie storage’s cookies that are sent to a specified
// URL.
//
// URL: The URL to filter on.
//
// # Return Value
// 
// An array of cookies whose URL matches the provided URL.
//
// # Discussion
// 
// You can use the [RequestHeaderFieldsWithCookies] method of [NSHTTPCookie]
// to turn the array returned by this method into a set of header fields to
// add to a [URLRequest] object (or [NSMutableURLRequest] in Objective-C).
// 
// If you override this method, also override
// [GetCookiesForTaskCompletionHandler].
//
// [URLRequest]: https://developer.apple.com/documentation/Foundation/URLRequest
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/cookies(for:)
func (h HTTPCookieStorage) CookiesForURL(URL INSURL) []NSHTTPCookie {
	rv := objc.Send[[]objc.ID](h.ID, objc.Sel("cookiesForURL:"), URL)
	return objc.ConvertSlice(rv, func(id objc.ID) NSHTTPCookie {
		return NSHTTPCookieFromID(id)
	})
}
// Returns all of the cookie storage’s cookies, sorted according to a given
// set of sort descriptors.
//
// sortOrder: The sort descriptors to use for sorting, as an array of [NSSortDescriptor]
// objects.
//
// # Return Value
// 
// The cookie storage’s cookies, sorted according to `sortOrder`, as an
// array of [NSHTTPCookie] objects.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/sortedCookies(using:)
func (h HTTPCookieStorage) SortedCookiesUsingDescriptors(sortOrder []NSSortDescriptor) []NSHTTPCookie {
	rv := objc.Send[[]objc.ID](h.ID, objc.Sel("sortedCookiesUsingDescriptors:"), objectivec.IObjectSliceToNSArray(sortOrder))
	return objc.ConvertSlice(rv, func(id objc.ID) NSHTTPCookie {
		return NSHTTPCookieFromID(id)
	})
}

// Returns the cookie storage instance for the container associated with the
// specified app group identifier.
//
// identifier: The app group identifier.
//
// # Discussion
// 
// By default, apps and associated app extensions will have different data
// containers. As a result, the value of the [NSHTTPCookieStorage] class’s
// [SharedHTTPCookieStorage] property will refer to different persistent
// cookie stores when called by the app and by its extensions.You can use this
// method to create a persistent cookie storage available to all apps and
// extensions with access to the same app group.
// 
// Subsequent calls to the this method with the same identifier will return
// the same storage instance.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/sharedCookieStorage(forGroupContainerIdentifier:)
func (_HTTPCookieStorageClass HTTPCookieStorageClass) SharedCookieStorageForGroupContainerIdentifier(identifier string) HTTPCookieStorage {
	rv := objc.Send[objc.ID](objc.ID(_HTTPCookieStorageClass.class), objc.Sel("sharedCookieStorageForGroupContainerIdentifier:"), objc.String(identifier))
	return NSHTTPCookieStorageFromID(rv)
}

// The cookie storage’s cookie accept policy.
//
// # Discussion
// 
// The default cookie accept policy is [HTTPCookie.AcceptPolicy.always].
// Changing the cookie policy affects all currently running applications using
// the cookie storage.
//
// [HTTPCookie.AcceptPolicy.always]: https://developer.apple.com/documentation/Foundation/HTTPCookie/AcceptPolicy/always
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/cookieAcceptPolicy
func (h HTTPCookieStorage) CookieAcceptPolicy() NSHTTPCookieAcceptPolicy {
	rv := objc.Send[NSHTTPCookieAcceptPolicy](h.ID, objc.Sel("cookieAcceptPolicy"))
	return NSHTTPCookieAcceptPolicy(rv)
}
func (h HTTPCookieStorage) SetCookieAcceptPolicy(value NSHTTPCookieAcceptPolicy) {
	objc.Send[struct{}](h.ID, objc.Sel("setCookieAcceptPolicy:"), value)
}
// The cookie storage’s cookies.
//
// # Discussion
// 
// If you want to sort the cookie storage’s cookies, you should use the
// [SortedCookiesUsingDescriptors] method instead of sorting the result of
// this method.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/cookies
func (h HTTPCookieStorage) Cookies() []NSHTTPCookie {
	rv := objc.Send[[]objc.ID](h.ID, objc.Sel("cookies"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSHTTPCookie {
		return NSHTTPCookieFromID(id)
	})
}
// A notification posted when the cookies stored in the cookie storage have
// changed.
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nshttpcookiemanagercookieschanged
func (h HTTPCookieStorage) NSHTTPCookieManagerCookiesChanged() NSNotificationName {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("NSHTTPCookieManagerCookiesChangedNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}
// A notification posted when the acceptance policy of the cookie storage has
// changed.
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nshttpcookiemanageracceptpolicychanged
func (h HTTPCookieStorage) NSHTTPCookieManagerAcceptPolicyChanged() NSNotificationName {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("NSHTTPCookieManagerAcceptPolicyChangedNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}
// A Boolean value that indicates whether the cookie should be discarded at
// the end of the session (regardless of expiration date).
//
// See: https://developer.apple.com/documentation/foundation/httpcookie/issessiononly
func (h HTTPCookieStorage) IsSessionOnly() bool {
	rv := objc.Send[bool](h.ID, objc.Sel("sessionOnly"))
	return rv
}
func (h HTTPCookieStorage) SetIsSessionOnly(value bool) {
	objc.Send[struct{}](h.ID, objc.Sel("setSessionOnly:"), value)
}

// The shared cookie storage instance.
//
// See: https://developer.apple.com/documentation/Foundation/HTTPCookieStorage/shared
func (_HTTPCookieStorageClass HTTPCookieStorageClass) SharedHTTPCookieStorage() HTTPCookieStorage {
	rv := objc.Send[objc.ID](objc.ID(_HTTPCookieStorageClass.class), objc.Sel("sharedHTTPCookieStorage"))
	return NSHTTPCookieStorageFromID(objc.ID(rv))
}

