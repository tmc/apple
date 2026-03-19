// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMediaLibraryBrowserController] class.
var (
	_NSMediaLibraryBrowserControllerClass     NSMediaLibraryBrowserControllerClass
	_NSMediaLibraryBrowserControllerClassOnce sync.Once
)

func getNSMediaLibraryBrowserControllerClass() NSMediaLibraryBrowserControllerClass {
	_NSMediaLibraryBrowserControllerClassOnce.Do(func() {
		_NSMediaLibraryBrowserControllerClass = NSMediaLibraryBrowserControllerClass{class: objc.GetClass("NSMediaLibraryBrowserController")}
	})
	return _NSMediaLibraryBrowserControllerClass
}

// GetNSMediaLibraryBrowserControllerClass returns the class object for NSMediaLibraryBrowserController.
func GetNSMediaLibraryBrowserControllerClass() NSMediaLibraryBrowserControllerClass {
	return getNSMediaLibraryBrowserControllerClass()
}

type NSMediaLibraryBrowserControllerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMediaLibraryBrowserControllerClass) Alloc() NSMediaLibraryBrowserController {
	rv := objc.Send[NSMediaLibraryBrowserController](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that configures and displays a Media Library Browser panel.
//
// # Overview
// 
// From this panel a user can drag media into views in their app. The class
// provides a standard interface to the MediaLibrary framework content.
// 
// For more information see [MLMediaLibrary], [MLMediaSource], [MLMediaGroup],
// and [MLMediaObject] in [Media Library].
// 
// # Pasteboard Types
// 
// The Media Library Browser defines two pasteboard types for decoding the
// dragged content and retrieving the media content that is appropriate, one
// for mediagroup content and one for individual media items
// 
// - The `com.Apple().MediaLibrary.PboardType.MediaGroupIdentifiersPlist`
// pasteboard type describes media group content and is published when the
// user drags items from the upper media-group-organized pane of the Media
// Library Browser. - The
// `com.Apple().MediaLibrary.PboardType.MediaObjectIdentifiersPlist`
// pasteboard type describes individual media items and is published when the
// user drags media items such as images, movies, or sounds from the media
// item pane of the Media Library Browser.
// 
// There is a third, less specialized, type of media library pasteboard. When
// the user initiates a drag the pasteboard will contain an array with one
// more more [NSFilenamesPboardType] pasteboard items, one for each of the
// files within the group dragged from the media-group-organized pane, and one
// or more items when a media item or items are dragged from the media item
// pane.
// 
// If you do not need access to the associated Media Library metadata, using
// the [NSFilenamesPboardType] pasteboard data is the simplest means of
// retrieving the dragged content, although accessing the media in this manner
// when your app is sandboxed requires that you use the [NSURL]
// [startAccessingSecurityScopedResource()] and
// [stopAccessingSecurityScopedResource()] methods.
// 
// # The Media Group Pasteboard Type
// 
// The `"com.Apple().MediaLibrary.PboardType.MediaGroupIdentifiersPlist"`
// pasteboard type is published when the user drags items from the upper
// media-group-organized pane of the Media Library Browser.
// 
// It consists of an [NSDictionary] plist that contains`keys` with the media
// source identifiers and corresponding `value` that are arrays of media group
// identifiers.
// 
// To decode the pasteboard data and get [MLMediaGroup] instances, assuming
// you have an [MLMediaLibrary] instance, use the techniques illustrated in
// the code listing below.
// 
// Listing 1. Retrieving MLMediaGroup instances from the pasteboard
// 
// # The Media Object Pasteboard Type
// 
// The `"com.Apple().MediaLibrary.PboardType.MediaObjectIdentifiersPlist"`
// pasteboard type is published when the user drags an item from the media
// item pane of the Media Library Browser.
// 
// It consists of an [NSDictionary] plist that contains`keys` with the media
// source identifiers and corresponding `value` that are arrays of media
// object identifiers.
// 
// To decode the pasteboard data and get [MLMediaObject] instances, assuming
// that you have an MLMediaLibrary instance, use the techniques illustrated in
// the code listing below.
// 
// Listing 2. Retrieving MLMediaObject instances from the pasteboard
//
// [MLMediaGroup]: https://developer.apple.com/documentation/MediaLibrary/MLMediaGroup
// [MLMediaLibrary]: https://developer.apple.com/documentation/MediaLibrary/MLMediaLibrary
// [MLMediaObject]: https://developer.apple.com/documentation/MediaLibrary/MLMediaObject
// [MLMediaSource]: https://developer.apple.com/documentation/MediaLibrary/MLMediaSource
// [Media Library]: https://developer.apple.com/documentation/MediaLibrary
// [NSFilenamesPboardType]: https://developer.apple.com/documentation/AppKit/NSFilenamesPboardType
// [startAccessingSecurityScopedResource()]: https://developer.apple.com/documentation/Foundation/NSURL/startAccessingSecurityScopedResource()
// [stopAccessingSecurityScopedResource()]: https://developer.apple.com/documentation/Foundation/NSURL/stopAccessingSecurityScopedResource()
//
// # Displaying the Media Library Browser Panel
//
//   - [NSMediaLibraryBrowserController.Frame]: The frame, in global coordinates, used to display the Media Library Browser panel.
//   - [NSMediaLibraryBrowserController.SetFrame]
//   - [NSMediaLibraryBrowserController.TogglePanel]: Toggles the visibility of the Media Library Browser.
//   - [NSMediaLibraryBrowserController.Visible]: A Boolean value that determines whether the Media Library Browser panel is visible.
//   - [NSMediaLibraryBrowserController.SetVisible]
//
// # Displayed Media Library Types
//
//   - [NSMediaLibraryBrowserController.MediaLibraries]: The media library that is in use.
//   - [NSMediaLibraryBrowserController.SetMediaLibraries]
//
// See: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController
type NSMediaLibraryBrowserController struct {
	objectivec.Object
}

// NSMediaLibraryBrowserControllerFromID constructs a [NSMediaLibraryBrowserController] from an objc.ID.
//
// An object that configures and displays a Media Library Browser panel.
func NSMediaLibraryBrowserControllerFromID(id objc.ID) NSMediaLibraryBrowserController {
	return NSMediaLibraryBrowserController{objectivec.Object{ID: id}}
}
// NOTE: NSMediaLibraryBrowserController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMediaLibraryBrowserController] class.
//
// # Displaying the Media Library Browser Panel
//
//   - [INSMediaLibraryBrowserController.Frame]: The frame, in global coordinates, used to display the Media Library Browser panel.
//   - [INSMediaLibraryBrowserController.SetFrame]
//   - [INSMediaLibraryBrowserController.TogglePanel]: Toggles the visibility of the Media Library Browser.
//   - [INSMediaLibraryBrowserController.Visible]: A Boolean value that determines whether the Media Library Browser panel is visible.
//   - [INSMediaLibraryBrowserController.SetVisible]
//
// # Displayed Media Library Types
//
//   - [INSMediaLibraryBrowserController.MediaLibraries]: The media library that is in use.
//   - [INSMediaLibraryBrowserController.SetMediaLibraries]
//
// See: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController
type INSMediaLibraryBrowserController interface {
	objectivec.IObject

	// Topic: Displaying the Media Library Browser Panel

	// The frame, in global coordinates, used to display the Media Library Browser panel.
	Frame() corefoundation.CGRect
	SetFrame(value corefoundation.CGRect)
	// Toggles the visibility of the Media Library Browser.
	TogglePanel(sender objectivec.IObject)
	// A Boolean value that determines whether the Media Library Browser panel is visible.
	Visible() bool
	SetVisible(value bool)

	// Topic: Displayed Media Library Types

	// The media library that is in use.
	MediaLibraries() NSMediaLibrary
	SetMediaLibraries(value NSMediaLibrary)
}

// Init initializes the instance.
func (m NSMediaLibraryBrowserController) Init() NSMediaLibraryBrowserController {
	rv := objc.Send[NSMediaLibraryBrowserController](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMediaLibraryBrowserController) Autorelease() NSMediaLibraryBrowserController {
	rv := objc.Send[NSMediaLibraryBrowserController](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMediaLibraryBrowserController creates a new NSMediaLibraryBrowserController instance.
func NewNSMediaLibraryBrowserController() NSMediaLibraryBrowserController {
	class := getNSMediaLibraryBrowserControllerClass()
	rv := objc.Send[NSMediaLibraryBrowserController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Toggles the visibility of the Media Library Browser.
//
// sender: The sender of the message.
//
// See: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/togglePanel(_:)
func (m NSMediaLibraryBrowserController) TogglePanel(sender objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("togglePanel:"), sender)
}

// The frame, in global coordinates, used to display the Media Library Browser
// panel.
//
// See: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/frame
func (m NSMediaLibraryBrowserController) Frame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](m.ID, objc.Sel("frame"))
	return corefoundation.CGRect(rv)
}
func (m NSMediaLibraryBrowserController) SetFrame(value corefoundation.CGRect) {
	objc.Send[struct{}](m.ID, objc.Sel("setFrame:"), value)
}
// A Boolean value that determines whether the Media Library Browser panel is
// visible.
//
// # Discussion
// 
// Set this value to [true] to show the Media Library Browser or [false] to
// hide it.
// 
// This value can be read to determine the current visibility status of the
// panel.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/isVisible
func (m NSMediaLibraryBrowserController) Visible() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isVisible"))
	return rv
}
func (m NSMediaLibraryBrowserController) SetVisible(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setVisible:"), value)
}
// The media library that is in use.
//
// # Discussion
// 
// This property will be one of the values in the
// [NSMediaLibraryBrowserController.Library] constants.
// 
// You can set the value to use a specific library (image, audio or movie).
// You can also read the value to determine which is currently displayed.
//
// [NSMediaLibraryBrowserController.Library]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/Library
//
// See: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/mediaLibraries
func (m NSMediaLibraryBrowserController) MediaLibraries() NSMediaLibrary {
	rv := objc.Send[NSMediaLibrary](m.ID, objc.Sel("mediaLibraries"))
	return NSMediaLibrary(rv)
}
func (m NSMediaLibraryBrowserController) SetMediaLibraries(value NSMediaLibrary) {
	objc.Send[struct{}](m.ID, objc.Sel("setMediaLibraries:"), value)
}

// Returns the shared Media Library Browser instance.
//
// # Return Value
// 
// The shared Media Library Browser controller instance.
// 
// # Discussion
// 
// The Media Library Browser panel is a proxy to allow easy display of the
// media library within an app.
//
// See: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/shared
func (_NSMediaLibraryBrowserControllerClass NSMediaLibraryBrowserControllerClass) SharedMediaLibraryBrowserController() NSMediaLibraryBrowserController {
	rv := objc.Send[objc.ID](objc.ID(_NSMediaLibraryBrowserControllerClass.class), objc.Sel("sharedMediaLibraryBrowserController"))
	return NSMediaLibraryBrowserControllerFromID(objc.ID(rv))
}

