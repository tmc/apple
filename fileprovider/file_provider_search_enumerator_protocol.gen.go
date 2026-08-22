// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines methods for providing search results and canceling searches.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchEnumerator
type NSFileProviderSearchEnumerator interface {
	objectivec.IObject

	// Enumerates search results starting from the specified page, in response to a call from the framework.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchEnumerator/enumerateSearchResults(for:startingAt:)
	EnumerateSearchResultsForObserverStartingAtPage(observer NSFileProviderSearchEnumerationObserver, page NSFileProviderPage)

	// Cancels a currently-running enumeration, in respone to a call from the framework.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchEnumerator/invalidate()
	Invalidate()
}

// NSFileProviderSearchEnumeratorObject wraps an existing Objective-C object that conforms to the NSFileProviderSearchEnumerator protocol.
type NSFileProviderSearchEnumeratorObject struct {
	objectivec.Object
}

func (o NSFileProviderSearchEnumeratorObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderSearchEnumeratorObjectFromID constructs a [NSFileProviderSearchEnumeratorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderSearchEnumeratorObjectFromID(id objc.ID) NSFileProviderSearchEnumeratorObject {
	return NSFileProviderSearchEnumeratorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Enumerates search results starting from the specified page, in response to
// a call from the framework.
//
// observer: An [NSFileProviderSearchEnumerationObserver], to which your extension
// provides search results.
//
// page: An indication of a location within the search results to resume
// enumeration. This parameter is non-`nil` if you previously provided a
// `nextPage` parameter to the observer’s [FinishEnumeratingUpToPage]
// method. Make sure the page contains whatever information you need to resume
// the enumeration.
//
// # Discussion
//
// Implement this method to perform your search and deliver pages of results
// to `observer`.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchEnumerator/enumerateSearchResults(for:startingAt:)
func (o NSFileProviderSearchEnumeratorObject) EnumerateSearchResultsForObserverStartingAtPage(observer NSFileProviderSearchEnumerationObserver, page NSFileProviderPage) {
	objc.Send[struct{}](o.ID, objc.Sel("enumerateSearchResultsForObserver:startingAtPage:"), observer, page)
}

// Cancels a currently-running enumeration, in respone to a call from the
// framework.
//
// # Discussion
//
// The framework calls this method to cancel a search if the person using the
// device changes their query, making the results of the current search
// obsolete. The framework also calls this method when it’s finished using
// this enumerator object.
//
// Implement this method by canceling any outstanding requests and cleaning up
// resources.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchEnumerator/invalidate()
func (o NSFileProviderSearchEnumeratorObject) Invalidate() {
	objc.Send[struct{}](o.ID, objc.Sel("invalidate"))
}
