// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// An interface you implement that represents an abstract location inside your document’s content.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLocation
type NSTextLocation interface {
	objectivec.IObject

	// Compares and returns the logical ordering to location.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextLocation/compare(_:)
	Compare(location NSTextLocation) foundation.NSComparisonResult
}



// NSTextLocationObject wraps an existing Objective-C object that conforms to the NSTextLocation protocol.
type NSTextLocationObject struct {
	objectivec.Object
}
func (o NSTextLocationObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSTextLocationObjectFromID constructs a [NSTextLocationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTextLocationObjectFromID(id objc.ID) NSTextLocationObject {
	return NSTextLocationObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Compares and returns the logical ordering to location.
//
// location: The location to compare the current location to.
//
// # Return Value
// 
// A [ComparisonResult].
//
// [ComparisonResult]: https://developer.apple.com/documentation/Foundation/ComparisonResult
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLocation/compare(_:)

func (o NSTextLocationObject) Compare(location NSTextLocation) foundation.NSComparisonResult {
	
	rv := objc.Send[foundation.NSComparisonResult](o.ID, objc.Sel("compare:"), location)
	return rv
	}







