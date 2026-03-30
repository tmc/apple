// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKFindResult] class.
var (
	_WKFindResultClass     WKFindResultClass
	_WKFindResultClassOnce sync.Once
)

func getWKFindResultClass() WKFindResultClass {
	_WKFindResultClassOnce.Do(func() {
		_WKFindResultClass = WKFindResultClass{class: objc.GetClass("WKFindResult")}
	})
	return _WKFindResultClass
}

// GetWKFindResultClass returns the class object for WKFindResult.
func GetWKFindResultClass() WKFindResultClass {
	return getWKFindResultClass()
}

type WKFindResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKFindResultClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKFindResultClass) Alloc() WKFindResult {
	rv := objc.Send[WKFindResult](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that contains the results of searching the web view’s contents.
//
// # Overview
//
// When you perform a search using the methods of [WKWebView], the web view
// creates a [WKFindResult] object and delivers it to your completion handler.
// You don’t create instances of this class directly. Use the objects that
// the web view provides to determine whether it found a match for the
// content.
//
// # Getting the Search Result
//
//   - [WKFindResult.MatchFound]: A Boolean value that indicates whether the web view found a match during the search.
//
// See: https://developer.apple.com/documentation/WebKit/WKFindResult
type WKFindResult struct {
	objectivec.Object
}

// WKFindResultFromID constructs a [WKFindResult] from an objc.ID.
//
// An object that contains the results of searching the web view’s contents.
func WKFindResultFromID(id objc.ID) WKFindResult {
	return WKFindResult{objectivec.Object{ID: id}}
}

// NOTE: WKFindResult adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKFindResult] class.
//
// # Getting the Search Result
//
//   - [IWKFindResult.MatchFound]: A Boolean value that indicates whether the web view found a match during the search.
//
// See: https://developer.apple.com/documentation/WebKit/WKFindResult
type IWKFindResult interface {
	objectivec.IObject

	// Topic: Getting the Search Result

	// A Boolean value that indicates whether the web view found a match during the search.
	MatchFound() bool
}

// Init initializes the instance.
func (f WKFindResult) Init() WKFindResult {
	rv := objc.Send[WKFindResult](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f WKFindResult) Autorelease() WKFindResult {
	rv := objc.Send[WKFindResult](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKFindResult creates a new WKFindResult instance.
func NewWKFindResult() WKFindResult {
	class := getWKFindResultClass()
	rv := objc.Send[WKFindResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether the web view found a match during
// the search.
//
// # Discussion
//
// The value of this property is true if the web view found a match, or false
// if it didn’t.
//
// See: https://developer.apple.com/documentation/WebKit/WKFindResult/matchFound
func (f WKFindResult) MatchFound() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("matchFound"))
	return rv
}
