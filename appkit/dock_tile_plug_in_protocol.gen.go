// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of methods implemented by plug-ins that allow an app’s Dock tile to be customized while the app is not running.
//
// See: https://developer.apple.com/documentation/AppKit/NSDockTilePlugIn
type NSDockTilePlugIn interface {
	objectivec.IObject

	// Invoked when the plug-in is first loaded and when the application is removed from the Dock.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSDockTilePlugIn/setDockTile(_:)
	SetDockTile(dockTile INSDockTile)
}

// NSDockTilePlugInObject wraps an existing Objective-C object that conforms to the NSDockTilePlugIn protocol.
type NSDockTilePlugInObject struct {
	objectivec.Object
}
func (o NSDockTilePlugInObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSDockTilePlugInObjectFromID constructs a [NSDockTilePlugInObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSDockTilePlugInObjectFromID(id objc.ID) NSDockTilePlugInObject {
	return NSDockTilePlugInObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Invoked when the plug-in is first loaded and when the application is
// removed from the Dock.
//
// dockTile: The dock tile associated with the application, or `nil` if the application
// has been removed from the Dock.
//
// # Discussion
// 
// The plugin is loaded in a system process at login time or when the
// application tile is added to the Dock.
// 
// The principal class of the plug-in must implement the [NSDockTilePlugIn]
// protocol.
//
// See: https://developer.apple.com/documentation/AppKit/NSDockTilePlugIn/setDockTile(_:)
func (o NSDockTilePlugInObject) SetDockTile(dockTile INSDockTile) {
	objc.Send[struct{}](o.ID, objc.Sel("setDockTile:"), dockTile)
	}
// Invoked when the user causes the application’s dock menu to be shown.
//
// # Return Value
// 
// The menu the dock tile displays.
//
// See: https://developer.apple.com/documentation/AppKit/NSDockTilePlugIn/dockMenu()
func (o NSDockTilePlugInObject) DockMenu() INSMenu {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("dockMenu"))
	return NSMenuFromID(rv)
	}

