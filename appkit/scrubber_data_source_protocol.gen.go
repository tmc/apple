// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that a scrubber data source object implements to provide items to the scrubber from an associated data collection in your app.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberDataSource
type NSScrubberDataSource interface {
	objectivec.IObject

	// Asks the data source for the number of items in the scrubber.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSScrubberDataSource/numberOfItems(for:)
	NumberOfItemsForScrubber(scrubber INSScrubber) int

	// Asks the data source object for the view the corresponds to the specified item in the scrubber.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSScrubberDataSource/scrubber(_:viewForItemAt:)
	ScrubberViewForItemAtIndex(scrubber INSScrubber, index int) INSScrubberItemView
}

// NSScrubberDataSourceObject wraps an existing Objective-C object that conforms to the NSScrubberDataSource protocol.
type NSScrubberDataSourceObject struct {
	objectivec.Object
}

func (o NSScrubberDataSourceObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSScrubberDataSourceObjectFromID constructs a [NSScrubberDataSourceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSScrubberDataSourceObjectFromID(id objc.ID) NSScrubberDataSourceObject {
	return NSScrubberDataSourceObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Asks the data source for the number of items in the scrubber.
//
// scrubber: The scrubber whose item count is being requested.
//
// # Return Value
//
// The number of items in the scrubber.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberDataSource/numberOfItems(for:)
func (o NSScrubberDataSourceObject) NumberOfItemsForScrubber(scrubber INSScrubber) int {
	rv := objc.Send[int](o.ID, objc.Sel("numberOfItemsForScrubber:"), scrubber)
	return rv
}

// Asks the data source object for the view the corresponds to the specified
// item in the scrubber.
//
// scrubber: The scrubber requesting the view.
//
// index: The index that specifies the location of the item in the scrubber.
//
// # Return Value
//
// A configured item view object.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberDataSource/scrubber(_:viewForItemAt:)
func (o NSScrubberDataSourceObject) ScrubberViewForItemAtIndex(scrubber INSScrubber, index int) INSScrubberItemView {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("scrubber:viewForItemAtIndex:"), scrubber, index)
	return NSScrubberItemViewFromID(rv)
}
