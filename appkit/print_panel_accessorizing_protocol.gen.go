// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that a Print panel object can use to get information from a printing accessory controller.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintPanelAccessorizing
type NSPrintPanelAccessorizing interface {
	objectivec.IObject

	// Returns an array of dictionaries containing the localized user setting summary strings.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSPrintPanelAccessorizing/localizedSummaryItems()
	LocalizedSummaryItems() foundation.INSDictionary
}

// NSPrintPanelAccessorizingObject wraps an existing Objective-C object that conforms to the NSPrintPanelAccessorizing protocol.
type NSPrintPanelAccessorizingObject struct {
	objectivec.Object
}
func (o NSPrintPanelAccessorizingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSPrintPanelAccessorizingObjectFromID constructs a [NSPrintPanelAccessorizingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSPrintPanelAccessorizingObjectFromID(id objc.ID) NSPrintPanelAccessorizingObject {
	return NSPrintPanelAccessorizingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns an array of dictionaries containing the localized user setting
// summary strings.
//
// # Return Value
// 
// An array of [NSDictionary] objects, each of which contains a
// [NSPrintPanelAccessorySummaryItemNameKey] and
// [NSPrintPanelAccessorySummaryItemDescriptionKey] key. The values for the
// keys are both strings. This method must not return `nil`.
//
// # Discussion
// 
// Accessory panels must implement this method to return information about the
// panel’s current settings. The returned array should contain a dictionary
// for each setting that is managed by the accessory panel and each dictionary
// should contain two key-value pairs identifying the name of the setting and
// its current value.
// 
// Your accessory view must be KVO-compliant for the `localizedSummaryItems`
// key path because [NSPrintPanel] object observes that key path and uses it
// to keep the contents of the summary view up to date. This means your view
// should manually send KVO notifications to observers for the
// `localizedSummaryItems` key path whenever the contents of the set of
// summary items changes. For more information on supporting key-value
// observing and manual notifications, see [Key-Value Observing Programming
// Guide].
//
// [Key-Value Observing Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueObserving/KeyValueObserving.html#//apple_ref/doc/uid/10000177i
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintPanelAccessorizing/localizedSummaryItems()

func (o NSPrintPanelAccessorizingObject) LocalizedSummaryItems() foundation.INSDictionary {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("localizedSummaryItems"))
	return foundation.NSDictionaryFromID(rv)
	}

// Returns a set of strings identifying the key paths for any properties that
// might affect the built-in print preview.
//
// # Return Value
// 
// A set of [NSString] objects identifying one or more key paths. Only key
// paths for properties that might affect the contents of the print preview
// should be returned.
//
// # Discussion
// 
// If an accessory view modifies printing-related properties that are used by
// the print preview, you should implement this method to return the key paths
// for those properties. For example, if you write an accessory view that lets
// the user change the left and right document margins in the current
// [NSPrintInfo] object, you would return the following key paths:
// `representedObject.LeftMargin()`, `representedObject.RightMargin()`. (The
// [NSPrintInfo] object is the represented object of the accessory
// controller.)
// 
// Implementation of this method is optional. You do not need to implement
// this method if you are not using the [NSPrintPanel] object’s built-in
// preview facilities. If you do use these facilities, however, you should
// implement this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintPanelAccessorizing/keyPathsForValuesAffectingPreview()

func (o NSPrintPanelAccessorizingObject) KeyPathsForValuesAffectingPreview() foundation.INSSet {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("keyPathsForValuesAffectingPreview"))
	return foundation.NSSetFromID(rv)
	}

