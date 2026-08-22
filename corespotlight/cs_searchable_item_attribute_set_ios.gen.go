// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.
//go:build ios
// +build ios

package corespotlight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/uniformtypeidentifiers"
)

// The identifiers that specify custom actions the app supports for the item.
//
// # Discussion
//
// The identifiers correspond to the [CoreSpotlightActionIdentifier] values
// you specify in the [CoreSpotlightActions] key of the app’s `Info.Plist()`
// file.
//
// When the user selects a custom action on an indexed item, the system
// launches your app and invokes
// [application(_:continue:restorationHandler:)]. The `userInfo` dictionary of
// the specified [NSUserActivity] includes the corresponding `Info.Plist()`
// entry using the key [CSActionIdentifier].
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/actionIdentifiers
//
// [CSActionIdentifier]: https://developer.apple.com/documentation/CoreSpotlight/CSActionIdentifier
// [CoreSpotlightActionIdentifier]: https://developer.apple.com/documentation/BundleResources/Information-Property-List/CoreSpotlightActions/CoreSpotlightActionIdentifier
// [CoreSpotlightActions]: https://developer.apple.com/documentation/BundleResources/Information-Property-List/CoreSpotlightActions
// [NSUserActivity]: https://developer.apple.com/documentation/Foundation/NSUserActivity
// [application(_:continue:restorationHandler:)]: https://developer.apple.com/documentation/UIKit/UIApplicationDelegate/application(_:continue:restorationHandler:)
func (c CSSearchableItemAttributeSet) ActionIdentifiers() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("actionIdentifiers"))
	return objc.ConvertSliceToStrings(rv)
}
func (c CSSearchableItemAttributeSet) SetActionIdentifiers(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setActionIdentifiers:"), objectivec.StringSliceToNSArray(value))
}

// The file type of the item to enable the user to share items from Spotlight.
//
// # Discussion
//
// Core Spotlight uses this property to determine the correct type identifier
// to pass to the
// [FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError] method.
//
// Spotlight enables sharing the item even if your app doesn’t set
// sharedItemContentType, but does support drag and drop for URL-backed types.
// Similarly, Spotlight enables copying items if your app supports drag and
// drop of Core Spotlight items.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/sharedItemContentType
func (c CSSearchableItemAttributeSet) SharedItemContentType() uniformtypeidentifiers.UTType {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("sharedItemContentType"))
	return uniformtypeidentifiers.UTTypeFromID(rv)
}
func (c CSSearchableItemAttributeSet) SetSharedItemContentType(value uniformtypeidentifiers.UTType) {
	objc.Send[struct{}](c.ID, objc.Sel("setSharedItemContentType:"), value)
}
