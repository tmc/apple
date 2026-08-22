// Code generated from Apple documentation for FinderSync. DO NOT EDIT.

package findersync

import (
	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The group of methods to implement for modifying the Finder user interface to express file synchronization status and control.
//
// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncProtocol
type FIFinderSyncProtocol interface {
	objectivec.IObject

	// The image for the extension’s toolbar button.
	//
	// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncProtocol/toolbarItemImage
	ToolbarItemImage() appkit.NSImage

	// The name of the extension’s toolbar button.
	//
	// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncProtocol/toolbarItemName
	ToolbarItemName() string

	// The tooltip text for the extension’s toolbar button.
	//
	// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncProtocol/toolbarItemToolTip
	ToolbarItemToolTip() string
}

// FIFinderSyncProtocolObject wraps an existing Objective-C object that conforms to the FIFinderSyncProtocol protocol.
type FIFinderSyncProtocolObject struct {
	objectivec.Object
}

func (o FIFinderSyncProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// FIFinderSyncProtocolObjectFromID constructs a [FIFinderSyncProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func FIFinderSyncProtocolObjectFromID(id objc.ID) FIFinderSyncProtocolObject {
	return FIFinderSyncProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
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
func (o FIFinderSyncProtocolObject) BeginObservingDirectoryAtURL(url foundation.NSURL) {
	objc.Send[struct{}](o.ID, objc.Sel("beginObservingDirectoryAtURL:"), url)
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
func (o FIFinderSyncProtocolObject) EndObservingDirectoryAtURL(url foundation.NSURL) {
	objc.Send[struct{}](o.ID, objc.Sel("endObservingDirectoryAtURL:"), url)
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
func (o FIFinderSyncProtocolObject) MenuForMenuKind(menu FIMenuKind) appkit.NSMenu {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("menuForMenuKind:"), menu)
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
func (o FIFinderSyncProtocolObject) RequestBadgeIdentifierForURL(url foundation.NSURL) {
	objc.Send[struct{}](o.ID, objc.Sel("requestBadgeIdentifierForURL:"), url)
}

// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncProtocol/makeListenerEndpoint(forServiceName:itemURL:)
func (o FIFinderSyncProtocolObject) MakeListenerEndpointForServiceNameItemURLAndReturnError(serviceName foundation.NSFileProviderServiceName, itemURL foundation.NSURL) (foundation.NSXPCListenerEndpoint, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("makeListenerEndpointForServiceName:itemURL:andReturnError:"), serviceName, itemURL)
	if err != nil {
		return foundation.NSXPCListenerEndpoint{}, err
	}
	return foundation.NSXPCListenerEndpointFromID(rv), nil
}

// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncProtocol/supportedServiceNamesForItem(with:)
func (o FIFinderSyncProtocolObject) SupportedServiceNamesForItemWithURL(itemURL foundation.NSURL) []string {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("supportedServiceNamesForItemWithURL:"), itemURL)
	return objc.ConvertSliceToStrings(rv)
}

// The image for the extension’s toolbar button.
//
// # Discussion
//
// To add a toolbar item to the Finder, override the getter method for the
// toolbar image, name, and tooltip properties.
//
// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncProtocol/toolbarItemImage
func (o FIFinderSyncProtocolObject) ToolbarItemImage() appkit.NSImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("toolbarItemImage"))
	return appkit.NSImageFromID(rv)
}

// The name of the extension’s toolbar button.
//
// # Discussion
//
// To add a toolbar item to the Finder, override the getter method for the
// toolbar image, name, and tooltip properties.
//
// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncProtocol/toolbarItemName
func (o FIFinderSyncProtocolObject) ToolbarItemName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("toolbarItemName"))
	return foundation.NSStringFromID(rv).String()
}

// The tooltip text for the extension’s toolbar button.
//
// # Discussion
//
// To add a toolbar item to the Finder, override the getter method for the
// toolbar image, name, and tooltip properties.
//
// See: https://developer.apple.com/documentation/FinderSync/FIFinderSyncProtocol/toolbarItemToolTip
func (o FIFinderSyncProtocolObject) ToolbarItemToolTip() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("toolbarItemToolTip"))
	return foundation.NSStringFromID(rv).String()
}
