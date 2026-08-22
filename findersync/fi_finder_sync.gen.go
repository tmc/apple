// Code generated from Apple documentation for FinderSync. DO NOT EDIT.

package findersync

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FIFinderSync] class.
var (
	_FIFinderSyncClass     FIFinderSyncClass
	_FIFinderSyncClassOnce sync.Once
)

func getFIFinderSyncClass() FIFinderSyncClass {
	_FIFinderSyncClassOnce.Do(func() {
		_FIFinderSyncClass = FIFinderSyncClass{class: objc.GetClass("FIFinderSync")}
	})
	return _FIFinderSyncClass
}

// GetFIFinderSyncClass returns the class object for FIFinderSync.
func GetFIFinderSyncClass() FIFinderSyncClass {
	return getFIFinderSyncClass()
}

type FIFinderSyncClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FIFinderSyncClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FIFinderSyncClass) Alloc() FIFinderSync {
	rv := objc.Send[FIFinderSync](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// A type to subclass to add badges, custom shortcut menus, and toolbar
// buttons to the Finder.
//
// # Overview
//
// Subclass the FIFinderSync class when you want to customize the appearance
// of the Finder. Although the FIFinderSync class doesn’t provide any
// developer accessible API, it does adopt the [FIFinderSyncProtocol]
// protocol. This protocol declares methods you can implement to modify the
// appearance of the Finder. For more information on these methods, see
// [FIFinderSyncProtocol]. To learn more about creating a Finder Sync
// extension, see [Finder Sync] in [App Extension Programming Guide].
//
// See: https://developer.apple.com/documentation/FinderSync/FIFinderSync-swift.class
//
// [App Extension Programming Guide]: https://developer.apple.com/library/archive/documentation/General/Conceptual/ExtensibilityPG/index.html#//apple_ref/doc/uid/TP40014214
// [Finder Sync]: https://developer.apple.com/library/archive/documentation/General/Conceptual/ExtensibilityPG/Finder.html#//apple_ref/doc/uid/TP40014214-CH15
type FIFinderSync struct {
	objectivec.Object
}

// FIFinderSyncFromID constructs a [FIFinderSync] from an objc.ID.
//
// A type to subclass to add badges, custom shortcut menus, and toolbar
// buttons to the Finder.
func FIFinderSyncFromID(id objc.ID) FIFinderSync {
	return FIFinderSync{objectivec.Object{ID: id}}
}

// NOTE: FIFinderSync adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FIFinderSync] class.
//
// See: https://developer.apple.com/documentation/FinderSync/FIFinderSync-swift.class
type IFIFinderSync interface {
	objectivec.IObject

	// Tells the extension that the user is looking at a monitored directory or at one of its subdirectories.
	BeginObservingDirectoryAtURL(url foundation.NSURL)
	// Tells the extension that the user has stopped looking at a monitored directory or at one of its subdirectories.
	EndObservingDirectoryAtURL(url foundation.NSURL)
	MakeListenerEndpointForServiceNameItemURLAndReturnError(serviceName foundation.NSFileProviderServiceName, itemURL foundation.NSURL) (foundation.NSXPCListenerEndpoint, error)
	// Requests a custom menu from the extension.
	MenuForMenuKind(menu FIMenuKind) appkit.NSMenu
	// Requests a badge for the given file or directory.
	RequestBadgeIdentifierForURL(url foundation.NSURL)
	SupportedServiceNamesForItemWithURL(itemURL foundation.NSURL) []string
	// The image for the extension’s toolbar button.
	ToolbarItemImage() appkit.NSImage
	// The name of the extension’s toolbar button.
	ToolbarItemName() string
	// The tooltip text for the extension’s toolbar button.
	ToolbarItemToolTip() string
}

// Init initializes the instance.
func (f FIFinderSync) Init() FIFinderSync {
	rv := objc.Send[FIFinderSync](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f FIFinderSync) Autorelease() FIFinderSync {
	rv := objc.Send[FIFinderSync](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewFIFinderSync creates a new FIFinderSync instance.
func NewFIFinderSync() FIFinderSync {
	class := getFIFinderSyncClass()
	rv := objc.Send[FIFinderSync](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Tells the extension that the user is looking at a monitored directory or at
// one of its subdirectories.
//
// url: The URL of the directory.
//
// # Discussion
//
// Override this method to receive notifications when the user opens the
// contents of a monitored directory or one of its subdirectories in the
// Finder. The system calls “ only once for each unique URL. As long as the
// content remains visible in at least one Finder window, any additional
// Finder windows that open to the same URL are ignored.
//
// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncProtocol/beginObservingDirectory(at:)
func (f FIFinderSync) BeginObservingDirectoryAtURL(url foundation.NSURL) {
	objc.Send[objc.ID](f.ID, objc.Sel("beginObservingDirectoryAtURL:"), url)
}

// Tells the extension that the user has stopped looking at a monitored
// directory or at one of its subdirectories.
//
// url: The URL of the directory.
//
// # Discussion
//
// Override this method to receive notifications when the user is no longer
// looking at the contents of the given URL. As with
// [BeginObservingDirectoryAtURL], the Open and Save dialogs are tracked
// separately from the Finder.
//
// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncProtocol/endObservingDirectory(at:)
func (f FIFinderSync) EndObservingDirectoryAtURL(url foundation.NSURL) {
	objc.Send[objc.ID](f.ID, objc.Sel("endObservingDirectoryAtURL:"), url)
}

// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncProtocol/makeListenerEndpoint(forServiceName:itemURL:)
func (f FIFinderSync) MakeListenerEndpointForServiceNameItemURLAndReturnError(serviceName foundation.NSFileProviderServiceName, itemURL foundation.NSURL) (foundation.NSXPCListenerEndpoint, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](f.ID, objc.Sel("makeListenerEndpointForServiceName:itemURL:andReturnError:"), serviceName, itemURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSXPCListenerEndpoint{}, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSXPCListenerEndpointFromID(rv), nil

}

// Requests a custom menu from the extension.
//
// menu: The type of menu being displayed. For a list of possible values, see
// [FIMenuKind].
//
// # Return Value
//
// A custom menu.
//
// # Discussion
//
// Override this method to provide custom menus in the Finder. You can
// customize this menu based both on the menu’s kind and on the selected and
// targeted items (if any). You can get the selected and targeted items from
// the extension’s [FIFinderSyncController].
//
// If `kind` is [FIMenuKindToolbarItemMenu], the system always calls this
// method even if the target and selection are not related to the extension.
//
// The extension’s principal object provides a method for each menu item’s
// assigned action.
//
// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncProtocol/menu(for:)
//
// [FIMenuKind]: https://developer.apple.com/documentation/FinderSync/FIMenuKind
func (f FIFinderSync) MenuForMenuKind(menu FIMenuKind) appkit.NSMenu {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("menuForMenuKind:"), menu)
	return appkit.NSMenuFromID(rv)
}

// Requests a badge for the given file or directory.
//
// url: The URL of a file or directory inside the extension’s monitored
// directories.
//
// # Discussion
//
// Override this method to receive notifications whenever a new item becomes
// visible in the Finder. Check the item’s state, and call
// [FIFinderSyncController.SetBadgeIdentifierForURL] to set an appropriate
// badge.
//
// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncProtocol/requestBadgeIdentifier(for:)
func (f FIFinderSync) RequestBadgeIdentifierForURL(url foundation.NSURL) {
	objc.Send[objc.ID](f.ID, objc.Sel("requestBadgeIdentifierForURL:"), url)
}

// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncProtocol/supportedServiceNamesForItem(with:)
func (f FIFinderSync) SupportedServiceNamesForItemWithURL(itemURL foundation.NSURL) []string {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("supportedServiceNamesForItemWithURL:"), itemURL)
	return objc.ConvertSliceToStrings(rv)
}

// The image for the extension’s toolbar button.
//
// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncProtocol/toolbarItemImage
func (f FIFinderSync) ToolbarItemImage() appkit.NSImage {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("toolbarItemImage"))
	return appkit.NSImageFromID(rv)
}

// The name of the extension’s toolbar button.
//
// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncProtocol/toolbarItemName
func (f FIFinderSync) ToolbarItemName() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("toolbarItemName"))
	return foundation.NSStringFromID(rv).String()
}

// The tooltip text for the extension’s toolbar button.
//
// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncProtocol/toolbarItemToolTip
func (f FIFinderSync) ToolbarItemToolTip() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("toolbarItemToolTip"))
	return foundation.NSStringFromID(rv).String()
}

// Protocol methods for FIFinderSyncProtocol
