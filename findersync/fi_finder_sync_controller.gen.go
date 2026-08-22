// Code generated from Apple documentation for FinderSync. DO NOT EDIT.

package findersync

import (
	"context"
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [FIFinderSyncController] class.
var (
	_FIFinderSyncControllerClass     FIFinderSyncControllerClass
	_FIFinderSyncControllerClassOnce sync.Once
)

func getFIFinderSyncControllerClass() FIFinderSyncControllerClass {
	_FIFinderSyncControllerClassOnce.Do(func() {
		_FIFinderSyncControllerClass = FIFinderSyncControllerClass{class: objc.GetClass("FIFinderSyncController")}
	})
	return _FIFinderSyncControllerClass
}

// GetFIFinderSyncControllerClass returns the class object for FIFinderSyncController.
func GetFIFinderSyncControllerClass() FIFinderSyncControllerClass {
	return getFIFinderSyncControllerClass()
}

type FIFinderSyncControllerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FIFinderSyncControllerClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FIFinderSyncControllerClass) Alloc() FIFinderSyncController {
	rv := objc.Send[FIFinderSyncController](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// A controller that acts as a bridge between your Finder Sync extension and
// the Finder itself.
//
// # Overview
//
// Use the Finder Sync controller to configure your extension, to set badges
// on items in the Finder’s window, and to get a list of selected and
// targeted items.
//
// # Managing the Finder Sync Controller
//
//   - [FIFinderSyncController.DirectoryURLs]: The directories managed by this extension.
//   - [FIFinderSyncController.SetDirectoryURLs]
//   - [FIFinderSyncController.SelectedItemURLs]: Returns an array of selected items.
//   - [FIFinderSyncController.SetBadgeIdentifierForURL]: Sets the badge for a file or directory.
//   - [FIFinderSyncController.SetBadgeImageLabelForBadgeIdentifier]: Sets the badge image and label for the given ID.
//   - [FIFinderSyncController.TargetedURL]: Returns the URL of the Finder’s current target.
//
// # Instance Methods
//
//   - [FIFinderSyncController.LastUsedDateForItemWithURL]
//   - [FIFinderSyncController.SetLastUsedDateForItemWithURLCompletion]
//   - [FIFinderSyncController.SetTagDataForItemWithURLCompletion]
//   - [FIFinderSyncController.TagDataForItemWithURL]
//
// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController
type FIFinderSyncController struct {
	foundation.NSExtensionContext
}

// FIFinderSyncControllerFromID constructs a [FIFinderSyncController] from an objc.ID.
//
// A controller that acts as a bridge between your Finder Sync extension and
// the Finder itself.
func FIFinderSyncControllerFromID(id objc.ID) FIFinderSyncController {
	return FIFinderSyncController{NSExtensionContext: foundation.NSExtensionContextFromID(id)}
}

// NOTE: FIFinderSyncController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FIFinderSyncController] class.
//
// # Managing the Finder Sync Controller
//
//   - [IFIFinderSyncController.DirectoryURLs]: The directories managed by this extension.
//   - [IFIFinderSyncController.SetDirectoryURLs]
//   - [IFIFinderSyncController.SelectedItemURLs]: Returns an array of selected items.
//   - [IFIFinderSyncController.SetBadgeIdentifierForURL]: Sets the badge for a file or directory.
//   - [IFIFinderSyncController.SetBadgeImageLabelForBadgeIdentifier]: Sets the badge image and label for the given ID.
//   - [IFIFinderSyncController.TargetedURL]: Returns the URL of the Finder’s current target.
//
// # Instance Methods
//
//   - [IFIFinderSyncController.LastUsedDateForItemWithURL]
//   - [IFIFinderSyncController.SetLastUsedDateForItemWithURLCompletion]
//   - [IFIFinderSyncController.SetTagDataForItemWithURLCompletion]
//   - [IFIFinderSyncController.TagDataForItemWithURL]
//
// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController
type IFIFinderSyncController interface {
	foundation.INSExtensionContext

	// Topic: Managing the Finder Sync Controller

	// The directories managed by this extension.
	DirectoryURLs() foundation.INSSet
	SetDirectoryURLs(value foundation.INSSet)
	// Returns an array of selected items.
	SelectedItemURLs() []foundation.NSURL
	// Sets the badge for a file or directory.
	SetBadgeIdentifierForURL(badgeID string, url foundation.NSURL)
	// Sets the badge image and label for the given ID.
	SetBadgeImageLabelForBadgeIdentifier(image appkit.NSImage, label string, badgeID string)
	// Returns the URL of the Finder’s current target.
	TargetedURL() foundation.NSURL

	// Topic: Instance Methods

	LastUsedDateForItemWithURL(itemURL foundation.NSURL) foundation.NSDate
	SetLastUsedDateForItemWithURLCompletion(lastUsedDate foundation.NSDate, itemURL foundation.NSURL, completion ErrorHandler)
	SetTagDataForItemWithURLCompletion(tagData foundation.NSData, itemURL foundation.NSURL, completion ErrorHandler)
	TagDataForItemWithURL(itemURL foundation.NSURL) foundation.NSData
}

// Init initializes the instance.
func (f FIFinderSyncController) Init() FIFinderSyncController {
	rv := objc.Send[FIFinderSyncController](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f FIFinderSyncController) Autorelease() FIFinderSyncController {
	rv := objc.Send[FIFinderSyncController](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewFIFinderSyncController creates a new FIFinderSyncController instance.
func NewFIFinderSyncController() FIFinderSyncController {
	class := getFIFinderSyncControllerClass()
	rv := objc.Send[FIFinderSyncController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an array of selected items.
//
// # Return Value
//
// An array of items currently selected in the Finder window.
//
// # Discussion
//
// Use this method when creating a shortcut menu or a menu for the
// extension’s toolbar button. You can then modify the menu’s content
// based on the items currently selected.
//
// This method returns valid values only from the Finder Sync extension’s
// [MenuForMenuKind] method or from one of the menu actions created in this
// method. If the selected items are outside the extension’s managed
// directories (for example, when the user clicks on the toolbar button), this
// method returns `nil`.
//
// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/selectedItemURLs()
func (f FIFinderSyncController) SelectedItemURLs() []foundation.NSURL {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("selectedItemURLs"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSURL {
		return foundation.NSURLFromID(id)
	})
}

// Sets the badge for a file or directory.
//
// badgeID: A unique ID, identifying the badge.
//
// url: The URL of the file or directory.
//
// # Discussion
//
// Adds the specified badge to the given file or directory. Setting the
// identifier to an empty string (`@""`) removes the badge.
//
// Avoid adding badges to items that the Finder hasn’t displayed yet. When
// setting the initial badge, call this method from your Finder Sync
// extension’s [RequestBadgeIdentifierForURL] method. When updating badges,
// call this method only for items that have already received a badge.
//
// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/setBadgeIdentifier(_:for:)
func (f FIFinderSyncController) SetBadgeIdentifierForURL(badgeID string, url foundation.NSURL) {
	objc.Send[objc.ID](f.ID, objc.Sel("setBadgeIdentifier:forURL:"), objc.String(badgeID), url)
}

// Sets the badge image and label for the given ID.
//
// image: An [NSImage] object. The system may or may not draw this image on top of
// the item’s icon; when it does, the system determines the overlay
// position. Don’t add any padding to the image to adjust this positioning.
// The image draws at up to 320 x 320 points.
//
// label: A label describing the sync state represented by this badge. Each label
// should be a short localized string, such as “Waiting.”
//
// badgeID: A unique ID, identifying this badge.
//
// # Discussion
//
// Use this method to configure your badges. Finder may display the image, the
// label or both. Your Finder Sync extension typically sets up a fixed number
// of badges during its `init` method.
//
// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/setBadgeImage(_:label:forBadgeIdentifier:)
//
// [NSImage]: https://developer.apple.com/documentation/AppKit/NSImage
func (f FIFinderSyncController) SetBadgeImageLabelForBadgeIdentifier(image appkit.NSImage, label string, badgeID string) {
	objc.Send[objc.ID](f.ID, objc.Sel("setBadgeImage:label:forBadgeIdentifier:"), image, objc.String(label), objc.String(badgeID))
}

// Returns the URL of the Finder’s current target.
//
// # Return Value
//
// The URL of the Finder’s current target.
//
// # Discussion
//
// Use this method when creating a custom shortcut menu for the Finder. This
// returns the URL of the item that the user Control-clicked, letting you
// customize the menu for that item.
//
// This method returns valid values only from the Finder Sync extension’s
// [MenuForMenuKind] method or from one of the menu actions created in this
// method. If the selected items are outside the extension’s managed
// directories (for example, when the user clicks on the toolbar button), this
// method returns `nil`.
//
// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/targetedURL()
func (f FIFinderSyncController) TargetedURL() foundation.NSURL {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("targetedURL"))
	return foundation.NSURLFromID(rv)
}

// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/lastUsedDateForItem(with:)
func (f FIFinderSyncController) LastUsedDateForItemWithURL(itemURL foundation.NSURL) foundation.NSDate {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("lastUsedDateForItemWithURL:"), itemURL)
	return foundation.NSDateFromID(rv)
}

// # Overview
//
// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/setLastUsedDate(_:forItemWith:completion:)
func (f FIFinderSyncController) SetLastUsedDateForItemWithURLCompletion(lastUsedDate foundation.NSDate, itemURL foundation.NSURL, completion ErrorHandler) {
	_block2, _ := NewErrorBlock(completion)
	objc.Send[objc.ID](f.ID, objc.Sel("setLastUsedDate:forItemWithURL:completion:"), lastUsedDate, itemURL, _block2)
}

// # Overview
//
// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/setTagData(_:forItemWith:completion:)
func (f FIFinderSyncController) SetTagDataForItemWithURLCompletion(tagData foundation.NSData, itemURL foundation.NSURL, completion ErrorHandler) {
	_block2, _ := NewErrorBlock(completion)
	objc.Send[objc.ID](f.ID, objc.Sel("setTagData:forItemWithURL:completion:"), tagData, itemURL, _block2)
}

// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/tagDataForItem(with:)
func (f FIFinderSyncController) TagDataForItemWithURL(itemURL foundation.NSURL) foundation.NSData {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("tagDataForItemWithURL:"), itemURL)
	return foundation.NSDataFromID(rv)
}

// Returns the shared Finder Sync controller object.
//
// # Return Value
//
// The default Finder Sync controller object for this extension.
//
// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/default()
func (_FIFinderSyncControllerClass FIFinderSyncControllerClass) DefaultController() FIFinderSyncController {
	rv := objc.Send[objc.ID](objc.ID(_FIFinderSyncControllerClass.class), objc.Sel("defaultController"))
	return FIFinderSyncControllerFromID(rv)
}

// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/showExtensionManagementInterface()
func (_FIFinderSyncControllerClass FIFinderSyncControllerClass) ShowExtensionManagementInterface() {
	objc.Send[objc.ID](objc.ID(_FIFinderSyncControllerClass.class), objc.Sel("showExtensionManagementInterface"))
}

// The directories managed by this extension.
//
// # Discussion
//
// The extension receives [BeginObservingDirectoryAtURL] and
// [EndObservingDirectoryAtURL] messages for every directory in this set and
// for all of their subdirectories.
//
// Always set `directoryURLs` when the extension starts. If there are no
// directories to watch, set `directoryURLs` to an empty set.
//
// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/directoryURLs
func (f FIFinderSyncController) DirectoryURLs() foundation.INSSet {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("directoryURLs"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (f FIFinderSyncController) SetDirectoryURLs(value foundation.INSSet) {
	objc.Send[struct{}](f.ID, objc.Sel("setDirectoryURLs:"), value)
}

// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncController/isExtensionEnabled
func (_FIFinderSyncControllerClass FIFinderSyncControllerClass) IsExtensionEnabled() bool {
	rv := objc.Send[bool](objc.ID(_FIFinderSyncControllerClass.class), objc.Sel("isExtensionEnabled"))
	return rv
}

// SetLastUsedDateForItemWithURLCompletionSync is a synchronous wrapper around [FIFinderSyncController.SetLastUsedDateForItemWithURLCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (f FIFinderSyncController) SetLastUsedDateForItemWithURLCompletionSync(ctx context.Context, lastUsedDate foundation.NSDate, itemURL foundation.NSURL) error {
	done := make(chan error, 1)
	f.SetLastUsedDateForItemWithURLCompletion(lastUsedDate, itemURL, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetTagDataForItemWithURLCompletionSync is a synchronous wrapper around [FIFinderSyncController.SetTagDataForItemWithURLCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (f FIFinderSyncController) SetTagDataForItemWithURLCompletionSync(ctx context.Context, tagData foundation.NSData, itemURL foundation.NSURL) error {
	done := make(chan error, 1)
	f.SetTagDataForItemWithURLCompletion(tagData, itemURL, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
