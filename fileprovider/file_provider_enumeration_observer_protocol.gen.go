// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An observer that receives batches of items during enumeration.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderEnumerationObserver
type NSFileProviderEnumerationObserver interface {
	objectivec.IObject

	// Provides a batch of enumerated items.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderEnumerationObserver/didEnumerate(_:)
	DidEnumerateItems(updatedItems []objectivec.IObject)

	// Tells the observer that all of the items have been enumerated up to the specified page.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderEnumerationObserver/finishEnumerating(upTo:)
	FinishEnumeratingUpToPage(nextPage NSFileProviderPage)

	// Tells the observer that an error occurred during item enumeration.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderEnumerationObserver/finishEnumeratingWithError(_:)
	FinishEnumeratingWithError(error_ foundation.NSError)

	// The page size that the system recommends.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderEnumerationObserver/suggestedPageSize
	SuggestedPageSize() int
}

// NSFileProviderEnumerationObserverObject wraps an existing Objective-C object that conforms to the NSFileProviderEnumerationObserver protocol.
type NSFileProviderEnumerationObserverObject struct {
	objectivec.Object
}

func (o NSFileProviderEnumerationObserverObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderEnumerationObserverObjectFromID constructs a [NSFileProviderEnumerationObserverObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderEnumerationObserverObjectFromID(id objc.ID) NSFileProviderEnumerationObserverObject {
	return NSFileProviderEnumerationObserverObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Provides a batch of enumerated items.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderEnumerationObserver/didEnumerate(_:)
func (o NSFileProviderEnumerationObserverObject) DidEnumerateItems(updatedItems []objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("didEnumerateItems:"), objectivec.IObjectSliceToNSArray(updatedItems))
}

// Tells the observer that all of the items have been enumerated up to the
// specified page.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderEnumerationObserver/finishEnumerating(upTo:)
func (o NSFileProviderEnumerationObserverObject) FinishEnumeratingUpToPage(nextPage NSFileProviderPage) {
	objc.Send[struct{}](o.ID, objc.Sel("finishEnumeratingUpToPage:"), nextPage)
}

// Tells the observer that an error occurred during item enumeration.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderEnumerationObserver/finishEnumeratingWithError(_:)
func (o NSFileProviderEnumerationObserverObject) FinishEnumeratingWithError(error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("finishEnumeratingWithError:"), error_)
}

// The page size that the system recommends.
//
// # Discussion
//
// The system suggests a page size to optimize performance based on the
// enumeration’s context. The system can request the enumeration of a
// container for various reasons, such as if the user opens the directory in
// Finder, opens a file in an application, or if the system needs to
// materialize the contents of a directory. Each case has its own performance
// profile.
//
// While using the suggested page size helps ensure the best user experience,
// the system enforces a maximum of 100 times the suggested size.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderEnumerationObserver/suggestedPageSize
func (o NSFileProviderEnumerationObserverObject) SuggestedPageSize() int {
	rv := objc.Send[int](o.ID, objc.Sel("suggestedPageSize"))
	return int(rv)
}
