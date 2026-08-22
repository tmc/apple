// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines a type that receives enumerations of search results from your extension.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchEnumerationObserver
type NSFileProviderSearchEnumerationObserver interface {
	objectivec.IObject

	// Delivers an array of search results to the observer.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchEnumerationObserver/didEnumerate(_:)
	DidEnumerateSearchResults(searchResults []objectivec.IObject)

	// Finish enumerating a page of results, and optionally provide a location within the results to continue the enumeration.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchEnumerationObserver/finishEnumerating(upTo:)
	FinishEnumeratingUpToPage(nextPage NSFileProviderPage)

	// Finishes a search enumeration by sending an error to the framework.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchEnumerationObserver/finishEnumeratingWithError(_:)
	FinishEnumeratingWithError(error_ foundation.NSError)

	// The maximum number of results to return in a single page enumeration.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchEnumerationObserver/maximumNumberOfResultsPerPage
	MaximumNumberOfResultsPerPage() int
}

// NSFileProviderSearchEnumerationObserverObject wraps an existing Objective-C object that conforms to the NSFileProviderSearchEnumerationObserver protocol.
type NSFileProviderSearchEnumerationObserverObject struct {
	objectivec.Object
}

func (o NSFileProviderSearchEnumerationObserverObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderSearchEnumerationObserverObjectFromID constructs a [NSFileProviderSearchEnumerationObserverObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderSearchEnumerationObserverObjectFromID(id objc.ID) NSFileProviderSearchEnumerationObserverObject {
	return NSFileProviderSearchEnumerationObserverObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Delivers an array of search results to the observer.
//
// # Discussion
//
// For files stored on your server, consult search indexes on the server and
// use them to create an array of [NSFileProviderSearchResult] instances that
// you provide to this method.
//
// You can call this method multiple times prior to calling
// [FinishEnumeratingUpToPage] or [FinishEnumeratingWithError], as long as the
// total number of results doesn’t exceed
// `NSFileProviderSearchEnumerationObserver/maxNumberOfResults`.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchEnumerationObserver/didEnumerate(_:)
func (o NSFileProviderSearchEnumerationObserverObject) DidEnumerateSearchResults(searchResults []objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("didEnumerateSearchResults:"), objectivec.IObjectSliceToNSArray(searchResults))
}

// Finish enumerating a page of results, and optionally provide a location
// within the results to continue the enumeration.
//
// # Discussion
//
// Call this method after you make one or more calls to
// [DidEnumerateSearchResults] to provide results to the observer. The
// collective results you provide in these calls constitues a “page” of
// results.
//
// Finish your page before sending
// `NSFileProviderSearchEnumerationObserver/maxNumberOfResults`. If you have
// more results to provide, use the `nextPage` parameter to indicate where to
// continue in your result set. The system sends the `nextPage` parameter the
// next time it calls your [EnumerateSearchResultsForObserverStartingAtPage]
// method.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchEnumerationObserver/finishEnumerating(upTo:)
func (o NSFileProviderSearchEnumerationObserverObject) FinishEnumeratingUpToPage(nextPage NSFileProviderPage) {
	objc.Send[struct{}](o.ID, objc.Sel("finishEnumeratingUpToPage:"), nextPage)
}

// Finishes a search enumeration by sending an error to the framework.
//
// # Discussion
//
// Finishing with an error causes the system to stop requesting additional
// pages of results. The system doesn’t retry after you call this method. If
// an error is potentially recoverable, you can perform your own retry in your
// implementation of [EnumerateSearchResultsForObserverStartingAtPage] and
// continue if successful, or end the query by calling this method.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchEnumerationObserver/finishEnumeratingWithError(_:)
func (o NSFileProviderSearchEnumerationObserverObject) FinishEnumeratingWithError(error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("finishEnumeratingWithError:"), error_)
}

// The maximum number of results to return in a single page enumeration.
//
// # Discussion
//
// If the extension returns more than this number of results in a single page
// enumeration, the system will crash the extension process.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchEnumerationObserver/maximumNumberOfResultsPerPage
func (o NSFileProviderSearchEnumerationObserverObject) MaximumNumberOfResultsPerPage() int {
	rv := objc.Send[int](o.ID, objc.Sel("maximumNumberOfResultsPerPage"))
	return int(rv)
}
