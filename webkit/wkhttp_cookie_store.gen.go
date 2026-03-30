// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKHTTPCookieStore] class.
var (
	_WKHTTPCookieStoreClass     WKHTTPCookieStoreClass
	_WKHTTPCookieStoreClassOnce sync.Once
)

func getWKHTTPCookieStoreClass() WKHTTPCookieStoreClass {
	_WKHTTPCookieStoreClassOnce.Do(func() {
		_WKHTTPCookieStoreClass = WKHTTPCookieStoreClass{class: objc.GetClass("WKHTTPCookieStore")}
	})
	return _WKHTTPCookieStoreClass
}

// GetWKHTTPCookieStoreClass returns the class object for WKHTTPCookieStore.
func GetWKHTTPCookieStoreClass() WKHTTPCookieStoreClass {
	return getWKHTTPCookieStoreClass()
}

type WKHTTPCookieStoreClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKHTTPCookieStoreClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKHTTPCookieStoreClass) Alloc() WKHTTPCookieStore {
	rv := objc.Send[WKHTTPCookieStore](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that manages the HTTP cookies associated with a particular web
// view.
//
// # Overview
//
// Use a [WKHTTPCookieStore] to specify the initial cookies for your webpages,
// and to manage cookies for your web content. For example, you might use this
// object to delete the cookie for the current session when the user logs out.
// To detect when the webpage changes a cookie, install a cookie observer
// using the [WKHTTPCookieStore.AddObserver] method.
//
// You don’t create a [WKHTTPCookieStore] object directly. Instead, retrieve
// this object from the [WKWebsiteDataStore] object in your web view’s
// configuration object.
//
// # Managing cookies
//
//   - [WKHTTPCookieStore.SetCookieCompletionHandler]: Adds a cookie to the cookie store.
//   - [WKHTTPCookieStore.DeleteCookieCompletionHandler]: Deletes the specified cookie.
//
// # Permitting cookie storage
//
//   - [WKHTTPCookieStore.GetCookiePolicy]: Returns a cookie policy that indicates whether the cookie store allows cookie storage.
//   - [WKHTTPCookieStore.SetCookiePolicyCompletionHandler]: Sets a cookie policy that indicates whether the cookie store allows cookie storage.
//
// # Observing cookie store changes
//
//   - [WKHTTPCookieStore.AddObserver]: Adds an observer to the cookie store.
//   - [WKHTTPCookieStore.RemoveObserver]: Removes an observer from the cookie store.
//
// # Instance Methods
//
//   - [WKHTTPCookieStore.SetCookiesCompletionHandler]
//
// See: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore
type WKHTTPCookieStore struct {
	objectivec.Object
}

// WKHTTPCookieStoreFromID constructs a [WKHTTPCookieStore] from an objc.ID.
//
// An object that manages the HTTP cookies associated with a particular web
// view.
func WKHTTPCookieStoreFromID(id objc.ID) WKHTTPCookieStore {
	return WKHTTPCookieStore{objectivec.Object{ID: id}}
}

// NOTE: WKHTTPCookieStore adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKHTTPCookieStore] class.
//
// # Managing cookies
//
//   - [IWKHTTPCookieStore.SetCookieCompletionHandler]: Adds a cookie to the cookie store.
//   - [IWKHTTPCookieStore.DeleteCookieCompletionHandler]: Deletes the specified cookie.
//
// # Permitting cookie storage
//
//   - [IWKHTTPCookieStore.GetCookiePolicy]: Returns a cookie policy that indicates whether the cookie store allows cookie storage.
//   - [IWKHTTPCookieStore.SetCookiePolicyCompletionHandler]: Sets a cookie policy that indicates whether the cookie store allows cookie storage.
//
// # Observing cookie store changes
//
//   - [IWKHTTPCookieStore.AddObserver]: Adds an observer to the cookie store.
//   - [IWKHTTPCookieStore.RemoveObserver]: Removes an observer from the cookie store.
//
// # Instance Methods
//
//   - [IWKHTTPCookieStore.SetCookiesCompletionHandler]
//
// See: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore
type IWKHTTPCookieStore interface {
	objectivec.IObject

	// Topic: Managing cookies

	// Adds a cookie to the cookie store.
	SetCookieCompletionHandler(cookie foundation.NSHTTPCookie, completionHandler VoidHandler)
	// Deletes the specified cookie.
	DeleteCookieCompletionHandler(cookie foundation.NSHTTPCookie, completionHandler VoidHandler)

	// Topic: Permitting cookie storage

	// Returns a cookie policy that indicates whether the cookie store allows cookie storage.
	GetCookiePolicy(completionHandler WKCookiePolicyHandler)
	// Sets a cookie policy that indicates whether the cookie store allows cookie storage.
	SetCookiePolicyCompletionHandler(policy WKCookiePolicy, completionHandler VoidHandler)

	// Topic: Observing cookie store changes

	// Adds an observer to the cookie store.
	AddObserver(observer WKHTTPCookieStoreObserver)
	// Removes an observer from the cookie store.
	RemoveObserver(observer WKHTTPCookieStoreObserver)

	// Topic: Instance Methods

	SetCookiesCompletionHandler(cookies []foundation.NSHTTPCookie, completionHandler VoidHandler)

	// The local files WebKit can access when loading content.
	ReadAccessURL() foundation.NSString
}

// Init initializes the instance.
func (h WKHTTPCookieStore) Init() WKHTTPCookieStore {
	rv := objc.Send[WKHTTPCookieStore](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h WKHTTPCookieStore) Autorelease() WKHTTPCookieStore {
	rv := objc.Send[WKHTTPCookieStore](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKHTTPCookieStore creates a new WKHTTPCookieStore instance.
func NewWKHTTPCookieStore() WKHTTPCookieStore {
	class := getWKHTTPCookieStoreClass()
	rv := objc.Send[WKHTTPCookieStore](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Adds a cookie to the cookie store.
//
// cookie: The cookie to add.
//
// completionHandler: A completion handler block to execute asynchronously after the method
// successfully stores the cookie. This block has no return value and no
// parameters.
//
// See: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore/setCookie(_:completionHandler:)
func (h WKHTTPCookieStore) SetCookieCompletionHandler(cookie foundation.NSHTTPCookie, completionHandler VoidHandler) {
	_block1, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](h.ID, objc.Sel("setCookie:completionHandler:"), cookie, _block1)
}

// Deletes the specified cookie.
//
// cookie: The cookie to delete.
//
// completionHandler: A completion handler block to execute asynchronously after the method
// successfully deletes the cookie. This block has no return value and no
// parameters.
//
// See: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore/delete(_:completionHandler:)
func (h WKHTTPCookieStore) DeleteCookieCompletionHandler(cookie foundation.NSHTTPCookie, completionHandler VoidHandler) {
	_block1, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](h.ID, objc.Sel("deleteCookie:completionHandler:"), cookie, _block1)
}

// Returns a cookie policy that indicates whether the cookie store allows
// cookie storage.
//
// completionHandler: The completion handler block to execute asynchronously with the cookie
// policy. This block has no return value, and takes the following parameter:
//
// cookiePolicy: A [WKHTTPCookieStore.CookiePolicy] case that indicates
// whether the cookie store allows cookie storage.
//
// See: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore/getCookiePolicy(_:)
//
// [WKHTTPCookieStore.CookiePolicy]: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore/CookiePolicy
func (h WKHTTPCookieStore) GetCookiePolicy(completionHandler WKCookiePolicyHandler) {
	_block0, _ := NewWKCookiePolicyBlock(completionHandler)
	objc.Send[objc.ID](h.ID, objc.Sel("getCookiePolicy:"), _block0)
}

// Sets a cookie policy that indicates whether the cookie store allows cookie
// storage.
//
// policy: A cookie policy that indicates whether the cookie store allows cookie
// storage.
//
// completionHandler: A block the system invokes after it sets the cookie policy.
//
// See: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore/setCookiePolicy(_:completionHandler:)
func (h WKHTTPCookieStore) SetCookiePolicyCompletionHandler(policy WKCookiePolicy, completionHandler VoidHandler) {
	_block1, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](h.ID, objc.Sel("setCookiePolicy:completionHandler:"), policy, _block1)
}

// Adds an observer to the cookie store.
//
// observer: The observer object to add. The cookie store doesn’t maintain a strong
// reference to the object you specify. You are responsible for removing your
// observer object before it becomes invalid.
//
// See: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore/add(_:)
func (h WKHTTPCookieStore) AddObserver(observer WKHTTPCookieStoreObserver) {
	objc.Send[objc.ID](h.ID, objc.Sel("addObserver:"), observer)
}

// Removes an observer from the cookie store.
//
// observer: The observer object to remove.
//
// See: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore/remove(_:)
func (h WKHTTPCookieStore) RemoveObserver(observer WKHTTPCookieStoreObserver) {
	objc.Send[objc.ID](h.ID, objc.Sel("removeObserver:"), observer)
}

// cookies: An array of cookies to set.
//
// completionHandler: A block to invoke once the cookies have been stored.
//
// # Discussion
//
// Set multiple cookies.
//
// See: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStore/setCookies(_:completionHandler:)
func (h WKHTTPCookieStore) SetCookiesCompletionHandler(cookies []foundation.NSHTTPCookie, completionHandler VoidHandler) {
	_block1, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](h.ID, objc.Sel("setCookies:completionHandler:"), cookies, _block1)
}

// The local files WebKit can access when loading content.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentReadingOptionKey/readAccessURL
func (h WKHTTPCookieStore) ReadAccessURL() foundation.NSString {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("readAccessURL"))
	return foundation.NSStringFromID(objc.ID(rv))
}

// SetCookie is a synchronous wrapper around [WKHTTPCookieStore.SetCookieCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (h WKHTTPCookieStore) SetCookie(ctx context.Context, cookie foundation.NSHTTPCookie) error {
	done := make(chan struct{}, 1)
	h.SetCookieCompletionHandler(cookie, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// DeleteCookie is a synchronous wrapper around [WKHTTPCookieStore.DeleteCookieCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (h WKHTTPCookieStore) DeleteCookie(ctx context.Context, cookie foundation.NSHTTPCookie) error {
	done := make(chan struct{}, 1)
	h.DeleteCookieCompletionHandler(cookie, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// GetCookiePolicySync is a synchronous wrapper around [WKHTTPCookieStore.GetCookiePolicy].
// It blocks until the completion handler fires or the context is cancelled.
func (h WKHTTPCookieStore) GetCookiePolicySync(ctx context.Context) (int, error) {
	done := make(chan int, 1)
	h.GetCookiePolicy(func(val int) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

// SetCookiePolicy is a synchronous wrapper around [WKHTTPCookieStore.SetCookiePolicyCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (h WKHTTPCookieStore) SetCookiePolicy(ctx context.Context, policy WKCookiePolicy) error {
	done := make(chan struct{}, 1)
	h.SetCookiePolicyCompletionHandler(policy, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
