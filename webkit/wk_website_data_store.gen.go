// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKWebsiteDataStore] class.
var (
	_WKWebsiteDataStoreClass     WKWebsiteDataStoreClass
	_WKWebsiteDataStoreClassOnce sync.Once
)

func getWKWebsiteDataStoreClass() WKWebsiteDataStoreClass {
	_WKWebsiteDataStoreClassOnce.Do(func() {
		_WKWebsiteDataStoreClass = WKWebsiteDataStoreClass{class: objc.GetClass("WKWebsiteDataStore")}
	})
	return _WKWebsiteDataStoreClass
}

// GetWKWebsiteDataStoreClass returns the class object for WKWebsiteDataStore.
func GetWKWebsiteDataStoreClass() WKWebsiteDataStoreClass {
	return getWKWebsiteDataStoreClass()
}

type WKWebsiteDataStoreClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKWebsiteDataStoreClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKWebsiteDataStoreClass) Alloc() WKWebsiteDataStore {
	rv := objc.Send[WKWebsiteDataStore](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that manages cookies, disk and memory caches, and other types of
// data for a web view.
//
// # Overview
//
// Use a [WKWebsiteDataStore] object to configure and manage web site data.
// Specifically, use this object to:
//
// - Manage cookies that your web site uses - Learn about the types of data
// that websites store - Remove unwanted web site data
//
// Create a data store object and assign it to the [WebsiteDataStore] property
// of a [WKWebViewConfiguration] object before you create your web view.
//
// By default, [WKWebViewConfiguration] uses the default data store returned
// by the [WKWebsiteDataStore.DefaultDataStore] method, which saves website data persistently to
// disk.
//
// To implement private browsing, create a nonpersistent data store using the
// [WKWebsiteDataStore.NonPersistentDataStore] method instead.
//
// To implement profile browsing, create a persistent data store using the
// [WKWebsiteDataStore.DataStoreForIdentifier] method, passing an identifier that you use to
// identify the data store.
//
// # Inspecting data store properties
//
//   - [WKWebsiteDataStore.Identifier]: An identifier that uniquely identifies a data store.
//   - [WKWebsiteDataStore.Persistent]: A Boolean value that indicates whether this object stores data to disk.
//
// # Retrieving a cookie store
//
//   - [WKWebsiteDataStore.HttpCookieStore]: The object that manages the HTTP cookies for your website.
//
// # Removing specific types of data
//
//   - [WKWebsiteDataStore.RemoveDataOfTypesForDataRecordsCompletionHandler]: Removes the specified types of website data from one or more data records.
//   - [WKWebsiteDataStore.RemoveDataOfTypesModifiedSinceCompletionHandler]: Removes website data that changed after the specified date.
//
// # Instance Methods
//
//   - [WKWebsiteDataStore.FetchDataOfTypesCompletionHandler]
//   - [WKWebsiteDataStore.RestoreDataCompletionHandler]
//
// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore
type WKWebsiteDataStore struct {
	objectivec.Object
}

// WKWebsiteDataStoreFromID constructs a [WKWebsiteDataStore] from an objc.ID.
//
// An object that manages cookies, disk and memory caches, and other types of
// data for a web view.
func WKWebsiteDataStoreFromID(id objc.ID) WKWebsiteDataStore {
	return WKWebsiteDataStore{objectivec.Object{ID: id}}
}

// NOTE: WKWebsiteDataStore adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKWebsiteDataStore] class.
//
// # Inspecting data store properties
//
//   - [IWKWebsiteDataStore.Identifier]: An identifier that uniquely identifies a data store.
//   - [IWKWebsiteDataStore.Persistent]: A Boolean value that indicates whether this object stores data to disk.
//
// # Retrieving a cookie store
//
//   - [IWKWebsiteDataStore.HttpCookieStore]: The object that manages the HTTP cookies for your website.
//
// # Removing specific types of data
//
//   - [IWKWebsiteDataStore.RemoveDataOfTypesForDataRecordsCompletionHandler]: Removes the specified types of website data from one or more data records.
//   - [IWKWebsiteDataStore.RemoveDataOfTypesModifiedSinceCompletionHandler]: Removes website data that changed after the specified date.
//
// # Instance Methods
//
//   - [IWKWebsiteDataStore.FetchDataOfTypesCompletionHandler]
//   - [IWKWebsiteDataStore.RestoreDataCompletionHandler]
//
// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore
type IWKWebsiteDataStore interface {
	objectivec.IObject

	// Topic: Inspecting data store properties

	// An identifier that uniquely identifies a data store.
	Identifier() foundation.NSUUID
	// A Boolean value that indicates whether this object stores data to disk.
	Persistent() bool

	// Topic: Retrieving a cookie store

	// The object that manages the HTTP cookies for your website.
	HttpCookieStore() IWKHTTPCookieStore

	// Topic: Removing specific types of data

	// Removes the specified types of website data from one or more data records.
	RemoveDataOfTypesForDataRecordsCompletionHandler(dataTypes foundation.INSSet, dataRecords []WKWebsiteDataRecord, completionHandler VoidHandler)
	// Removes website data that changed after the specified date.
	RemoveDataOfTypesModifiedSinceCompletionHandler(dataTypes foundation.INSSet, date foundation.INSDate, completionHandler VoidHandler)

	// Topic: Instance Methods

	FetchDataOfTypesCompletionHandler(dataTypes foundation.INSSet, completionHandler DataErrorHandler)
	RestoreDataCompletionHandler(data foundation.INSData, completionHandler ErrorHandler)

	ProxyConfigurations() []objectivec.Object
	SetProxyConfigurations(value []objectivec.Object)
	// The local files WebKit can access when loading content.
	ReadAccessURL() foundation.NSString
	// The object you use to get and set the site’s cookies and to track the cached data objects.
	WebsiteDataStore() IWKWebsiteDataStore
	SetWebsiteDataStore(value IWKWebsiteDataStore)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (w WKWebsiteDataStore) Init() WKWebsiteDataStore {
	rv := objc.Send[WKWebsiteDataStore](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WKWebsiteDataStore) Autorelease() WKWebsiteDataStore {
	rv := objc.Send[WKWebsiteDataStore](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKWebsiteDataStore creates a new WKWebsiteDataStore instance.
func NewWKWebsiteDataStore() WKWebsiteDataStore {
	class := getWKWebsiteDataStoreClass()
	rv := objc.Send[WKWebsiteDataStore](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the persistent data store with the unique identifier you provide.
//
// identifier: An identifier that uniquely identifies a data store.
//
// # Discussion
//
// If the data store for the unique identifier you provide does not exist, the
// system creates and returns it. Use this method to get a data store for a
// profile.
//
// This method throws an exception if the identifier is not valid.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/init(forIdentifier:)
func NewWebsiteDataStoreForIdentifier(identifier foundation.NSUUID) WKWebsiteDataStore {
	rv := objc.Send[objc.ID](objc.ID(getWKWebsiteDataStoreClass().class), objc.Sel("dataStoreForIdentifier:"), identifier)
	return WKWebsiteDataStoreFromID(rv)
}

// Removes the specified types of website data from one or more data records.
//
// dataTypes: The website data types to remove from the records.
//
// dataRecords: The records that contain the data.
//
// completionHandler: The completion handler block to execute asynchronously after the web view
// removes the specified data. This block has no return value and takes no
// parameters.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/removeData(ofTypes:for:completionHandler:)
func (w WKWebsiteDataStore) RemoveDataOfTypesForDataRecordsCompletionHandler(dataTypes foundation.INSSet, dataRecords []WKWebsiteDataRecord, completionHandler VoidHandler) {
	_block2, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("removeDataOfTypes:forDataRecords:completionHandler:"), dataTypes, dataRecords, _block2)
}

// Removes website data that changed after the specified date.
//
// dataTypes: The website data types to remove.
//
// date: The target date for the data removal. The data store removes data that a
// website changed after this date.
//
// completionHandler: The completion handler block to execute asynchronously after the web view
// removes the specified data. This block has no return value and takes no
// parameters.
//
// # Discussion
//
// This method removes the specified data type from all records, but only if a
// website modified the record’s data after the specified `date`.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/removeData(ofTypes:modifiedSince:completionHandler:)
func (w WKWebsiteDataStore) RemoveDataOfTypesModifiedSinceCompletionHandler(dataTypes foundation.INSSet, date foundation.INSDate, completionHandler VoidHandler) {
	_block2, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("removeDataOfTypes:modifiedSince:completionHandler:"), dataTypes, date, _block2)
}

// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/fetchData(of:completionHandler:)
func (w WKWebsiteDataStore) FetchDataOfTypesCompletionHandler(dataTypes foundation.INSSet, completionHandler DataErrorHandler) {
	_block1, _ := NewDataErrorBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("fetchDataOfTypes:completionHandler:"), dataTypes, _block1)
}

// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/restoreData(_:completionHandler:)
func (w WKWebsiteDataStore) RestoreDataCompletionHandler(data foundation.INSData, completionHandler ErrorHandler) {
	_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("restoreData:completionHandler:"), data, _block1)
}
func (w WKWebsiteDataStore) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](w.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns the default data store, which stores data persistently to disk.
//
// # Return Value
//
// The default data store for websites.
//
// # Discussion
//
// A web view configured with the default data store saves website data
// persistenly to disk. Use this data store to retain the state of web content
// between browsing sessions.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/default()
func (_WKWebsiteDataStoreClass WKWebsiteDataStoreClass) DefaultDataStore() WKWebsiteDataStore {
	rv := objc.Send[objc.ID](objc.ID(_WKWebsiteDataStoreClass.class), objc.Sel("defaultDataStore"))
	return WKWebsiteDataStoreFromID(rv)
}

// Creates a new data store object that stores website data in memory, and
// doesn’t write that data to disk.
//
// # Return Value
//
// A new data store object that doesn’t save data to disk.
//
// # Discussion
//
// Use a nonpersistent data store to implement private browsing in your web
// view. This method creates a new data store that stores data only in memory,
// and doesn’t write that data to disk.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/nonPersistent()
func (_WKWebsiteDataStoreClass WKWebsiteDataStoreClass) NonPersistentDataStore() WKWebsiteDataStore {
	rv := objc.Send[objc.ID](objc.ID(_WKWebsiteDataStoreClass.class), objc.Sel("nonPersistentDataStore"))
	return WKWebsiteDataStoreFromID(rv)
}

// Returns the set of all the available data types.
//
// # Discussion
//
// Potential values in the set are [WKWebsiteDataRecord] constants.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/allWebsiteDataTypes()
func (_WKWebsiteDataStoreClass WKWebsiteDataStoreClass) AllWebsiteDataTypes() foundation.INSSet {
	rv := objc.Send[objc.ID](objc.ID(_WKWebsiteDataStoreClass.class), objc.Sel("allWebsiteDataTypes"))
	return foundation.NSSetFromID(rv)
}

// Removes the data store that matches the identifier you provide.
//
// identifier: An identifier that uniquely identifies a data store.
//
// completionHandler: A block the system invokes after it removes the data store. This block has
// no return value, and takes the following parameter:
//
// error: An error, if the system could not remove the data store.
//
// # Discussion
//
// Call this method to remove the data store for the unique identifier.
// Release any [WKWebView] instances using the data store before you call this
// method. If the system cannot complete removal of the data store, this
// throws an error.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/remove(forIdentifier:completionHandler:)
func (_WKWebsiteDataStoreClass WKWebsiteDataStoreClass) RemoveDataStoreForIdentifierCompletionHandler(identifier foundation.NSUUID, completionHandler ErrorHandler) {
	_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](objc.ID(_WKWebsiteDataStoreClass.class), objc.Sel("removeDataStoreForIdentifier:completionHandler:"), identifier, _block1)
}

// An identifier that uniquely identifies a data store.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/identifier
func (w WKWebsiteDataStore) Identifier() foundation.NSUUID {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("identifier"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}

// A Boolean value that indicates whether this object stores data to disk.
//
// # Discussion
//
// The value of this property is true if the data store writes data to disk,
// or false if it doesn’t.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/isPersistent
func (w WKWebsiteDataStore) Persistent() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isPersistent"))
	return rv
}

// The object that manages the HTTP cookies for your website.
//
// # Discussion
//
// Use the [WKHTTPCookieStore] object in this property to manage your
// website’s cookies. A cookie store object contains methods to add new
// cookies, and to remove cookies that don’t apply to your content. For
// example, you might initially use this object to add a cookie for the
// user’s login credentials and then delete that cookie when the user logs
// out.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/httpCookieStore
func (w WKWebsiteDataStore) HttpCookieStore() IWKHTTPCookieStore {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("httpCookieStore"))
	return WKHTTPCookieStoreFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataStore/proxyConfigurations-6g21z
func (w WKWebsiteDataStore) ProxyConfigurations() []objectivec.Object {
	rv := objc.Send[[]objc.ID](w.ID, objc.Sel("proxyConfigurations"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.Object {
		return objectivec.ObjectFromID(id)
	})
}
func (w WKWebsiteDataStore) SetProxyConfigurations(value []objectivec.Object) {
	objc.Send[struct{}](w.ID, objc.Sel("setProxyConfigurations:"), objectivec.IObjectSliceToNSArray(value))
}

// The local files WebKit can access when loading content.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentReadingOptionKey/readAccessURL
func (w WKWebsiteDataStore) ReadAccessURL() foundation.NSString {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("readAccessURL"))
	return foundation.NSStringFromID(objc.ID(rv))
}

// The object you use to get and set the site’s cookies and to track the
// cached data objects.
//
// See: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/websitedatastore
func (w WKWebsiteDataStore) WebsiteDataStore() IWKWebsiteDataStore {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("websiteDataStore"))
	return WKWebsiteDataStoreFromID(objc.ID(rv))
}
func (w WKWebsiteDataStore) SetWebsiteDataStore(value IWKWebsiteDataStore) {
	objc.Send[struct{}](w.ID, objc.Sel("setWebsiteDataStore:"), value)
}

// RemoveDataOfTypesModifiedSince is a synchronous wrapper around [WKWebsiteDataStore.RemoveDataOfTypesModifiedSinceCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w WKWebsiteDataStore) RemoveDataOfTypesModifiedSince(ctx context.Context, dataTypes foundation.INSSet, date foundation.INSDate) error {
	done := make(chan struct{}, 1)
	w.RemoveDataOfTypesModifiedSinceCompletionHandler(dataTypes, date, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RemoveDataStoreForIdentifier is a synchronous wrapper around [WKWebsiteDataStore.RemoveDataStoreForIdentifierCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (wc WKWebsiteDataStoreClass) RemoveDataStoreForIdentifier(ctx context.Context, identifier foundation.NSUUID) error {
	done := make(chan error, 1)
	wc.RemoveDataStoreForIdentifierCompletionHandler(identifier, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// FetchDataOfTypes is a synchronous wrapper around [WKWebsiteDataStore.FetchDataOfTypesCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w WKWebsiteDataStore) FetchDataOfTypes(ctx context.Context, dataTypes foundation.INSSet) (*foundation.NSData, error) {
	type result struct {
		val *foundation.NSData
		err error
	}
	done := make(chan result, 1)
	w.FetchDataOfTypesCompletionHandler(dataTypes, func(val *foundation.NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// RestoreData is a synchronous wrapper around [WKWebsiteDataStore.RestoreDataCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w WKWebsiteDataStore) RestoreData(ctx context.Context, data foundation.INSData) error {
	done := make(chan error, 1)
	w.RestoreDataCompletionHandler(data, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
