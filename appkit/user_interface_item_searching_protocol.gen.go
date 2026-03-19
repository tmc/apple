// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of methods an app can implement to provide Spotlight for Help for its own custom help data.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceItemSearching
type NSUserInterfaceItemSearching interface {
	objectivec.IObject
}

// NSUserInterfaceItemSearchingObject wraps an existing Objective-C object that conforms to the NSUserInterfaceItemSearching protocol.
type NSUserInterfaceItemSearchingObject struct {
	objectivec.Object
}
func (o NSUserInterfaceItemSearchingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSUserInterfaceItemSearchingObjectFromID constructs a [NSUserInterfaceItemSearchingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSUserInterfaceItemSearchingObjectFromID(id objc.ID) NSUserInterfaceItemSearchingObject {
	return NSUserInterfaceItemSearchingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns an array of localized strings that will form the help menu item.
//
// item: At item in the help menu.
//
// # Return Value
// 
// An [NSArray] of [NSStrings] (localized for display in the menu) that will
// be combined with separators to form the menu item title.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceItemSearching/localizedTitles(forItem:)
func (o NSUserInterfaceItemSearchingObject) LocalizedTitlesForItem(item objectivec.IObject) []string {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("localizedTitlesForItem:"), item)
	return objc.ConvertSliceToStrings(rv)
	}
// If this method is implemented, a “Show All Help Topics” item will
// appear in the menu and this method is called when the user selects it.
//
// searchString: The search string.
//
// # Discussion
// 
// The application should show all its results for this search, which does not
// include results for menu items. The string for “Show All Help Topics”
// is system defined and localized and cannot be changed by the user.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceItemSearching/showAllHelpTopics(forSearch:)
func (o NSUserInterfaceItemSearchingObject) ShowAllHelpTopicsForSearchString(searchString string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("showAllHelpTopicsForSearchString:"), objc.String(searchString))
	}
// Invoked when the user selects a search result in Help menu.
//
// item: An item in the help menu.
//
// # Discussion
// 
// The default implementation brings up Help Viewer for a Help item.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceItemSearching/performAction(forItem:)
func (o NSUserInterfaceItemSearchingObject) PerformActionForItem(item objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("performActionForItem:"), item)
	}

