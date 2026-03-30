// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The methods to adopt in an object that monitors changes to a webpage’s cookies.
//
// See: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStoreObserver
type WKHTTPCookieStoreObserver interface {
	objectivec.IObject
}

// WKHTTPCookieStoreObserverObject wraps an existing Objective-C object that conforms to the WKHTTPCookieStoreObserver protocol.
type WKHTTPCookieStoreObserverObject struct {
	objectivec.Object
}

func (o WKHTTPCookieStoreObserverObject) BaseObject() objectivec.Object {
	return o.Object
}

// WKHTTPCookieStoreObserverObjectFromID constructs a [WKHTTPCookieStoreObserverObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func WKHTTPCookieStoreObserverObjectFromID(id objc.ID) WKHTTPCookieStoreObserverObject {
	return WKHTTPCookieStoreObserverObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate that the cookies in the specified cookie store changed.
//
// cookieStore: The cookie store that contains the modified cookies.
//
// # Discussion
//
// When the value of a cookie changes, the cookie store calls this method on
// all registered observer objects. Use this method to fetch the new cookie
// values and update any app-specific data structures or JavaScript
// environment variables that use those values.
//
// See: https://developer.apple.com/documentation/WebKit/WKHTTPCookieStoreObserver/cookiesDidChange(in:)
func (o WKHTTPCookieStoreObserverObject) CookiesDidChangeInCookieStore(cookieStore IWKHTTPCookieStore) {
	objc.Send[struct{}](o.ID, objc.Sel("cookiesDidChangeInCookieStore:"), cookieStore)
}
