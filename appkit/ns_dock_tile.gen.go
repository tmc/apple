// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSDockTile] class.
var (
	_NSDockTileClass     NSDockTileClass
	_NSDockTileClassOnce sync.Once
)

func getNSDockTileClass() NSDockTileClass {
	_NSDockTileClassOnce.Do(func() {
		_NSDockTileClass = NSDockTileClass{class: objc.GetClass("NSDockTile")}
	})
	return _NSDockTileClass
}

// GetNSDockTileClass returns the class object for NSDockTile.
func GetNSDockTileClass() NSDockTileClass {
	return getNSDockTileClass()
}

type NSDockTileClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSDockTileClass) Alloc() NSDockTile {
	rv := objc.Send[NSDockTile](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The visual representation of your app’s miniaturized windows and app icon
// as they appear in the Dock.
//
// # Overview
// 
// You do not create Dock tile objects explicitly in your app. Instead, you
// retrieve the Dock tile for an existing window or for the app by calling
// that object’s [DockTile] method. Also, you do not subclass the
// [NSDockTile] class; instead, you use the methods of the class to make the
// following customizations:
// 
// - Badge the tile with a custom string. - Remove or show the application
// icon badge. - Draw the tile content yourself.
// 
// If you decide to draw the tile content yourself, you must provide a custom
// content view to handle the drawing.
// 
// # Application Dock Tiles
// 
// An application Dock tile defaults to display the application’s
// [NSDockTile.ApplicationIconImage].
// 
// The application Dock tile never shows a smaller application icon badge.
// 
// Whether using the default or custom view, the application Dock tile may be
// badged with a short custom string.
// 
// # Window Dock Tiles
// 
// A window Dock tile defaults to display a miniaturized version of the
// windows contents with a badge derived from the application Dock icon,
// including any customized application Dock icon. The default window Dock
// tile image may not be badged with a custom string.
// 
// A window Dock tile can use a custom view to draw the Dock icon. If a custom
// view is used, no application badge will be added, but the text label will
// be overlaid on top of the icon.
//
// # Drawing the Tile’s Content
//
//   - [NSDockTile.ContentView]: The view to use for drawing the dock tile contents.
//   - [NSDockTile.SetContentView]
//
// # Getting the Tile Information
//
//   - [NSDockTile.Size]: The size of the tile.
//   - [NSDockTile.Owner]: The object represented by the dock tile.
//
// # Applying Badge Icons to the Tile
//
//   - [NSDockTile.ShowsApplicationBadge]: A Boolean showing whether the tile is badged with the application’s icon
//   - [NSDockTile.SetShowsApplicationBadge]
//   - [NSDockTile.BadgeLabel]: The string to be displayed in the tile’s badging area.
//   - [NSDockTile.SetBadgeLabel]
//
// # Updating the Dock Tile
//
//   - [NSDockTile.Display]: Redraws the dock tile’s content.
//
// See: https://developer.apple.com/documentation/AppKit/NSDockTile
type NSDockTile struct {
	objectivec.Object
}

// NSDockTileFromID constructs a [NSDockTile] from an objc.ID.
//
// The visual representation of your app’s miniaturized windows and app icon
// as they appear in the Dock.
func NSDockTileFromID(id objc.ID) NSDockTile {
	return NSDockTile{objectivec.Object{ID: id}}
}
// NOTE: NSDockTile adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSDockTile] class.
//
// # Drawing the Tile’s Content
//
//   - [INSDockTile.ContentView]: The view to use for drawing the dock tile contents.
//   - [INSDockTile.SetContentView]
//
// # Getting the Tile Information
//
//   - [INSDockTile.Size]: The size of the tile.
//   - [INSDockTile.Owner]: The object represented by the dock tile.
//
// # Applying Badge Icons to the Tile
//
//   - [INSDockTile.ShowsApplicationBadge]: A Boolean showing whether the tile is badged with the application’s icon
//   - [INSDockTile.SetShowsApplicationBadge]
//   - [INSDockTile.BadgeLabel]: The string to be displayed in the tile’s badging area.
//   - [INSDockTile.SetBadgeLabel]
//
// # Updating the Dock Tile
//
//   - [INSDockTile.Display]: Redraws the dock tile’s content.
//
// See: https://developer.apple.com/documentation/AppKit/NSDockTile
type INSDockTile interface {
	objectivec.IObject

	// Topic: Drawing the Tile’s Content

	// The view to use for drawing the dock tile contents.
	ContentView() INSView
	SetContentView(value INSView)

	// Topic: Getting the Tile Information

	// The size of the tile.
	Size() corefoundation.CGSize
	// The object represented by the dock tile.
	Owner() objectivec.IObject

	// Topic: Applying Badge Icons to the Tile

	// A Boolean showing whether the tile is badged with the application’s icon
	ShowsApplicationBadge() bool
	SetShowsApplicationBadge(value bool)
	// The string to be displayed in the tile’s badging area.
	BadgeLabel() string
	SetBadgeLabel(value string)

	// Topic: Updating the Dock Tile

	// Redraws the dock tile’s content.
	Display()

	// The image used for the app’s icon.
	ApplicationIconImage() INSImage
	SetApplicationIconImage(value INSImage)
	// The application’s Dock tile.
	DockTile() INSDockTile
	SetDockTile(value INSDockTile)
}

// Init initializes the instance.
func (d NSDockTile) Init() NSDockTile {
	rv := objc.Send[NSDockTile](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NSDockTile) Autorelease() NSDockTile {
	rv := objc.Send[NSDockTile](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSDockTile creates a new NSDockTile instance.
func NewNSDockTile() NSDockTile {
	class := getNSDockTileClass()
	rv := objc.Send[NSDockTile](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Redraws the dock tile’s content.
//
// # Discussion
// 
// If a custom content view is provided, Cocoa calls the `` method of that
// view (and its subviews) to draw the tile’s content.
// 
// You can call this method to force the redrawing of the dock tile contents.
// You might do this if the contents of the underlying application or window
// change in a way that would require a refreshing of the tile. Some types of
// system activity, such as resizing the dock, may trigger automatic redraws
// of the tile. In most cases, however, your application is responsible for
// triggering redraws.
// 
// Cocoa does not automatically redraw the contents of your dock tile.
// Instead, your application must explicitly send `display` messages to the
// dock tile object whenever the contents of your view change and need to be
// redrawn.
//
// See: https://developer.apple.com/documentation/AppKit/NSDockTile/display()
func (d NSDockTile) Display() {
	objc.Send[objc.ID](d.ID, objc.Sel("display"))
}

// The view to use for drawing the dock tile contents.
//
// # Discussion
// 
// The view you specify should be height and width resizable.
// 
// Cocoa does not automatically redraw the contents of your dock tile.
// Instead, your application must explicitly send display messages to the dock
// tile object whenever the contents of your view change and need to be
// redrawn. Your dock tile view is responsible for drawing the entire contents
// of the dock tile. Your view does not need to draw the application or custom
// string badges.
//
// See: https://developer.apple.com/documentation/AppKit/NSDockTile/contentView
func (d NSDockTile) ContentView() INSView {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("contentView"))
	return NSViewFromID(objc.ID(rv))
}
func (d NSDockTile) SetContentView(value INSView) {
	objc.Send[struct{}](d.ID, objc.Sel("setContentView:"), value)
}
// The size of the tile.
//
// # Discussion
// 
// This corresponds to the size of the backing store in the dock, which may be
// bigger than the actual tile displayed on the screen.
//
// See: https://developer.apple.com/documentation/AppKit/NSDockTile/size
func (d NSDockTile) Size() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](d.ID, objc.Sel("size"))
	return corefoundation.CGSize(rv)
}
// The object represented by the dock tile.
//
// See: https://developer.apple.com/documentation/AppKit/NSDockTile/owner
func (d NSDockTile) Owner() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("owner"))
	return objectivec.Object{ID: rv}
}
// A Boolean showing whether the tile is badged with the application’s icon
//
// # Discussion
// 
// Miniaturized windows include the application badge by default to convey the
// associated application to the user. In macOS 10.5 and later, application
// tiles do not support the application badge. A miniaturized window with a
// custom view does not draw the application badge.
// 
// The application icon is positioned automatically in the tile by the
// [NSDockTile] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSDockTile/showsApplicationBadge
func (d NSDockTile) ShowsApplicationBadge() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("showsApplicationBadge"))
	return rv
}
func (d NSDockTile) SetShowsApplicationBadge(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setShowsApplicationBadge:"), value)
}
// The string to be displayed in the tile’s badging area.
//
// # Discussion
// 
// The appearance of the badge area is system defined.
//
// See: https://developer.apple.com/documentation/AppKit/NSDockTile/badgeLabel
func (d NSDockTile) BadgeLabel() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("badgeLabel"))
	return foundation.NSStringFromID(rv).String()
}
func (d NSDockTile) SetBadgeLabel(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setBadgeLabel:"), objc.String(value))
}
// The image used for the app’s icon.
//
// See: https://developer.apple.com/documentation/appkit/nsapplication/applicationiconimage
func (d NSDockTile) ApplicationIconImage() INSImage {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("applicationIconImage"))
	return NSImageFromID(objc.ID(rv))
}
func (d NSDockTile) SetApplicationIconImage(value INSImage) {
	objc.Send[struct{}](d.ID, objc.Sel("setApplicationIconImage:"), value)
}
// The application’s Dock tile.
//
// See: https://developer.apple.com/documentation/appkit/nswindow/docktile
func (d NSDockTile) DockTile() INSDockTile {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("dockTile"))
	return NSDockTileFromID(objc.ID(rv))
}
func (d NSDockTile) SetDockTile(value INSDockTile) {
	objc.Send[struct{}](d.ID, objc.Sel("setDockTile:"), value)
}

