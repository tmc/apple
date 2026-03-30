// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKContentRuleListStore] class.
var (
	_WKContentRuleListStoreClass     WKContentRuleListStoreClass
	_WKContentRuleListStoreClassOnce sync.Once
)

func getWKContentRuleListStoreClass() WKContentRuleListStoreClass {
	_WKContentRuleListStoreClassOnce.Do(func() {
		_WKContentRuleListStoreClass = WKContentRuleListStoreClass{class: objc.GetClass("WKContentRuleListStore")}
	})
	return _WKContentRuleListStoreClass
}

// GetWKContentRuleListStoreClass returns the class object for WKContentRuleListStore.
func GetWKContentRuleListStoreClass() WKContentRuleListStoreClass {
	return getWKContentRuleListStoreClass()
}

type WKContentRuleListStoreClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKContentRuleListStoreClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKContentRuleListStoreClass) Alloc() WKContentRuleListStore {
	rv := objc.Send[WKContentRuleListStore](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that contains the rules for how to load and filter content in the
// web view.
//
// # Overview
//
// Use a [WKContentRuleListStore] to compile and manage rules for filtering
// content in a web view. Rule lists act as content blockers inside your app.
// You use them to prevent the web view from loading specific content, either
// based on the original location of that content or other criteria you
// specify. For example, a corporate app might use rules to prevent the web
// view from loading content that originates from outside the corporate
// network.
//
// Fetch the default [WKContentRuleListStore] object or create a custom store
// object and use it to compile or access the available rules. Each store
// object stores its existing rules persistently in the file system and loads
// those rules at creation time. A store object doesn’t automatically apply
// any of its rules to a particular web view. To apply a rule to a web view,
// add it to the [WKUserContentController] object of the web view’s
// configuration object.
//
// # Creating and Deleting Content Rule Lists
//
//   - [WKContentRuleListStore.CompileContentRuleListForIdentifierEncodedContentRuleListCompletionHandler]: Compiles the specified JSON content into a new rule list and adds it to the current data store.
//   - [WKContentRuleListStore.RemoveContentRuleListForIdentifierCompletionHandler]: Removes a rule list from the current data store asynchronously.
//
// # Accessing the Current Rule Lists
//
//   - [WKContentRuleListStore.GetAvailableContentRuleListIdentifiers]: Fetches the identifiers for all rule lists in the store asynchronously.
//   - [WKContentRuleListStore.LookUpContentRuleListForIdentifierCompletionHandler]: Searches asynchronously for a specific rule list in the data store.
//
// See: https://developer.apple.com/documentation/WebKit/WKContentRuleListStore
type WKContentRuleListStore struct {
	objectivec.Object
}

// WKContentRuleListStoreFromID constructs a [WKContentRuleListStore] from an objc.ID.
//
// An object that contains the rules for how to load and filter content in the
// web view.
func WKContentRuleListStoreFromID(id objc.ID) WKContentRuleListStore {
	return WKContentRuleListStore{objectivec.Object{ID: id}}
}

// NOTE: WKContentRuleListStore adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKContentRuleListStore] class.
//
// # Creating and Deleting Content Rule Lists
//
//   - [IWKContentRuleListStore.CompileContentRuleListForIdentifierEncodedContentRuleListCompletionHandler]: Compiles the specified JSON content into a new rule list and adds it to the current data store.
//   - [IWKContentRuleListStore.RemoveContentRuleListForIdentifierCompletionHandler]: Removes a rule list from the current data store asynchronously.
//
// # Accessing the Current Rule Lists
//
//   - [IWKContentRuleListStore.GetAvailableContentRuleListIdentifiers]: Fetches the identifiers for all rule lists in the store asynchronously.
//   - [IWKContentRuleListStore.LookUpContentRuleListForIdentifierCompletionHandler]: Searches asynchronously for a specific rule list in the data store.
//
// See: https://developer.apple.com/documentation/WebKit/WKContentRuleListStore
type IWKContentRuleListStore interface {
	objectivec.IObject

	// Topic: Creating and Deleting Content Rule Lists

	// Compiles the specified JSON content into a new rule list and adds it to the current data store.
	CompileContentRuleListForIdentifierEncodedContentRuleListCompletionHandler(identifier string, encodedContentRuleList string, completionHandler WKContentRuleListErrorHandler)
	// Removes a rule list from the current data store asynchronously.
	RemoveContentRuleListForIdentifierCompletionHandler(identifier string, completionHandler ErrorHandler)

	// Topic: Accessing the Current Rule Lists

	// Fetches the identifiers for all rule lists in the store asynchronously.
	GetAvailableContentRuleListIdentifiers(completionHandler VoidHandler)
	// Searches asynchronously for a specific rule list in the data store.
	LookUpContentRuleListForIdentifierCompletionHandler(identifier string, completionHandler WKContentRuleListErrorHandler)
}

// Init initializes the instance.
func (c WKContentRuleListStore) Init() WKContentRuleListStore {
	rv := objc.Send[WKContentRuleListStore](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c WKContentRuleListStore) Autorelease() WKContentRuleListStore {
	rv := objc.Send[WKContentRuleListStore](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKContentRuleListStore creates a new WKContentRuleListStore instance.
func NewWKContentRuleListStore() WKContentRuleListStore {
	class := getWKContentRuleListStoreClass()
	rv := objc.Send[WKContentRuleListStore](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new content rule list store in the specified directory.
//
// url: A URL that specifies a directory. The returned object uses this directory
// to store its content rules persistently. For example, you might store your
// app-specific rules in your app’s container directory.
//
// # Return Value
//
// A content rule store associated with the specified directory.
//
// # Discussion
//
// If the specified directory already contains compiled rule lists, this
// method loads those rules and adds them to the returned object. If you
// change any rules after creating this object, the store saves those changes
// to the same directory.
//
// See: https://developer.apple.com/documentation/WebKit/WKContentRuleListStore/init(url:)
func NewContentRuleListStoreWithURL(url foundation.INSURL) WKContentRuleListStore {
	rv := objc.Send[objc.ID](objc.ID(getWKContentRuleListStoreClass().class), objc.Sel("storeWithURL:"), url)
	return WKContentRuleListStoreFromID(rv)
}

// Compiles the specified JSON content into a new rule list and adds it to the
// current data store.
//
// identifier: A unique identifier for the new list. If a list with the specified
// identifier already exists in the store, this method overwrites the old rule
// list with the new content.
//
// encodedContentRuleList: The JSON source for the new rule list. For information about how to format
// the JSON content, see [Creating a content blocker].
//
// completionHandler: A completion handler block to call after compilation finishes. This block
// has no return value and takes the following parameters:
//
// ruleList: The [WKContentRuleList] object that encapsulates the compiled
// rules derived from the `encodedContentRuleList` parameter. This parameter
// is `nil` if an error occurs during compilation. error: `nil` on success, or
// an error object if a problem occurred.
//
// See: https://developer.apple.com/documentation/WebKit/WKContentRuleListStore/compileContentRuleList(forIdentifier:encodedContentRuleList:completionHandler:)
//
// [Creating a content blocker]: https://developer.apple.com/documentation/SafariServices/creating-a-content-blocker
func (c WKContentRuleListStore) CompileContentRuleListForIdentifierEncodedContentRuleListCompletionHandler(identifier string, encodedContentRuleList string, completionHandler WKContentRuleListErrorHandler) {
	_block2, _ := NewWKContentRuleListErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("compileContentRuleListForIdentifier:encodedContentRuleList:completionHandler:"), objc.String(identifier), objc.String(encodedContentRuleList), _block2)
}

// Removes a rule list from the current data store asynchronously.
//
// identifier: The unique identifier for the rule list.
//
// completionHandler: A completion handler block to call after the removal of the content rule
// list. This block has no return value and takes the following parameter:
//
// error: `nil` on success, or an error object if the store encountered an
// error when deleting the rule list.
//
// # Discussion
//
// This method also removes the persistent copy of the rules stored on disk.
//
// See: https://developer.apple.com/documentation/WebKit/WKContentRuleListStore/removeContentRuleList(forIdentifier:completionHandler:)
func (c WKContentRuleListStore) RemoveContentRuleListForIdentifierCompletionHandler(identifier string, completionHandler ErrorHandler) {
	_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("removeContentRuleListForIdentifier:completionHandler:"), objc.String(identifier), _block1)
}

// Fetches the identifiers for all rule lists in the store asynchronously.
//
// completionHandler: A completion handler block to call with the results. This block has no
// return value and takes the following parameter:
//
// identifierArray: An array of strings, each of which corresponds to an
// identifier for a rule list in the data store. Use each string to look up
// the associated [WKContentRuleList] object. If the data store has no rule
// lists, the array is empty.
//
// See: https://developer.apple.com/documentation/WebKit/WKContentRuleListStore/getAvailableContentRuleListIdentifiers(_:)
func (c WKContentRuleListStore) GetAvailableContentRuleListIdentifiers(completionHandler VoidHandler) {
	_block0, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("getAvailableContentRuleListIdentifiers:"), _block0)
}

// Searches asynchronously for a specific rule list in the data store.
//
// identifier: The identifier of the list you want.
//
// completionHandler: A completion handler block to call with the results of the search. This
// block has no return value and takes the following parameters:
//
// ruleList: The [WKContentRuleList] object with the specified identifier.
// This parameter is `nil` if the error occurs during the search. error: `nil`
// on success, or an error object if an error occurs during the search.
//
// See: https://developer.apple.com/documentation/WebKit/WKContentRuleListStore/lookUpContentRuleList(forIdentifier:completionHandler:)
func (c WKContentRuleListStore) LookUpContentRuleListForIdentifierCompletionHandler(identifier string, completionHandler WKContentRuleListErrorHandler) {
	_block1, _ := NewWKContentRuleListErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("lookUpContentRuleListForIdentifier:completionHandler:"), objc.String(identifier), _block1)
}

// Returns the default content rule list store.
//
// # Return Value
//
// The default rule list store.
//
// # Discussion
//
// The default store contains the rules that your app created specifically for
// the current user.
//
// See: https://developer.apple.com/documentation/WebKit/WKContentRuleListStore/default()
func (_WKContentRuleListStoreClass WKContentRuleListStoreClass) DefaultStore() WKContentRuleListStore {
	rv := objc.Send[objc.ID](objc.ID(_WKContentRuleListStoreClass.class), objc.Sel("defaultStore"))
	return WKContentRuleListStoreFromID(rv)
}

// CompileContentRuleListForIdentifierEncodedContentRuleList is a synchronous wrapper around [WKContentRuleListStore.CompileContentRuleListForIdentifierEncodedContentRuleListCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c WKContentRuleListStore) CompileContentRuleListForIdentifierEncodedContentRuleList(ctx context.Context, identifier string, encodedContentRuleList string) (*WKContentRuleList, error) {
	type result struct {
		val *WKContentRuleList
		err error
	}
	done := make(chan result, 1)
	c.CompileContentRuleListForIdentifierEncodedContentRuleListCompletionHandler(identifier, encodedContentRuleList, func(val *WKContentRuleList, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// RemoveContentRuleListForIdentifier is a synchronous wrapper around [WKContentRuleListStore.RemoveContentRuleListForIdentifierCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c WKContentRuleListStore) RemoveContentRuleListForIdentifier(ctx context.Context, identifier string) error {
	done := make(chan error, 1)
	c.RemoveContentRuleListForIdentifierCompletionHandler(identifier, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// GetAvailableContentRuleListIdentifiersSync is a synchronous wrapper around [WKContentRuleListStore.GetAvailableContentRuleListIdentifiers].
// It blocks until the completion handler fires or the context is cancelled.
func (c WKContentRuleListStore) GetAvailableContentRuleListIdentifiersSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	c.GetAvailableContentRuleListIdentifiers(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// LookUpContentRuleListForIdentifier is a synchronous wrapper around [WKContentRuleListStore.LookUpContentRuleListForIdentifierCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c WKContentRuleListStore) LookUpContentRuleListForIdentifier(ctx context.Context, identifier string) (*WKContentRuleList, error) {
	type result struct {
		val *WKContentRuleList
		err error
	}
	done := make(chan result, 1)
	c.LookUpContentRuleListForIdentifierCompletionHandler(identifier, func(val *WKContentRuleList, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
