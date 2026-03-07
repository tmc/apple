// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf


// A set of optional methods that a window’s delegate can implement to respond to events, such as window resizing, moving, exposing, and minimizing.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate
type NSWindowDelegate interface {
	objectivec.IObject
}



// NSWindowDelegateObject wraps an existing Objective-C object that conforms to the NSWindowDelegate protocol.
type NSWindowDelegateObject struct {
	objectivec.Object
}
func (o NSWindowDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSWindowDelegateObjectFromID constructs a [NSWindowDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSWindowDelegateObjectFromID(id objc.ID) NSWindowDelegateObject {
	return NSWindowDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Tells the delegate that the window is about to show a sheet at the
// specified location, giving it the opportunity to return a custom location
// for the attachment of the sheet to the window.
//
// window: The window containing the sheet to be animated.
//
// sheet: The sheet to be shown.
//
// rect: The default sheet location, just under the title bar of the window, aligned
// with the left and right edges of the window.
//
// # Return Value
// 
// The custom location specified.
//
// # Discussion
// 
// This method is also called whenever the user resizes `window` while `sheet`
// is attached.
// 
// This method is useful in many situations. If your window has a toolbar, for
// example, you can specify a location for the sheet that is just below it. If
// you want the sheet associated with a certain control or view, you could
// position the sheet so that it appears to originate from the object (through
// animation) or is positioned next to it.
// 
// Neither the `rect` parameter nor the returned [NSRect] value define the
// boundary of the sheet. They indicate where the top-left edge of the sheet
// is attached to the window. The origin is expressed in window coordinates;
// the default `origin.Y()` value is the height of the content view and the
// default `origin.X()` value is 0. The `size.Width()` value indicates the
// width and behavior of the initial animation; if `size.Width()` is narrower
// than the sheet, the sheet genies out from the specified location, and if
// `size.Width()` is wider than the sheet, the sheet slides out. You cannot
// affect the size of the sheet through the `size.Width()` and `size.Height()`
// fields. It is recommended that you specify zero for the `size.Height()`
// value as this field may have additional meaning in a future release.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/window(_:willPositionSheet:using:)

func (o NSWindowDelegateObject) WindowWillPositionSheetUsingRect(window INSWindow, sheet INSWindow, rect corefoundation.CGRect) corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("window:willPositionSheet:usingRect:"), window, sheet, rect)
	return rv
	}

// Notifies the delegate that the window is about to open a sheet.
//
// notification: A notification named [willBeginSheetNotification].
// //
// [willBeginSheetNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/willBeginSheetNotification
//
// # Discussion
// 
// You can retrieve the window object in question by sending [object] to
// `notification`.
//
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowWillBeginSheet(_:)

func (o NSWindowDelegateObject) WindowWillBeginSheet(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowWillBeginSheet:"), notification)
	}

// Tells the delegate that the window has closed a sheet.
//
// notification: A notification named [didEndSheetNotification].
// //
// [didEndSheetNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didEndSheetNotification
//
// # Discussion
// 
// You can retrieve the window object in question by sending [object] to
// `notification`.
//
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowDidEndSheet(_:)

func (o NSWindowDelegateObject) WindowDidEndSheet(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowDidEndSheet:"), notification)
	}

// Tells the delegate that the window is being resized (whether by the user or
// through one of the `setFrame...` methods other than [SetFrameDisplay]).
//
// sender: The window being resized.
//
// frameSize: The size to which the specified window is being resized.
//
// # Return Value
// 
// A custom size to which the specified window will be resized.
//
// # Discussion
// 
// The `frameSize` contains the size (in screen coordinates) `sender` will be
// resized to. To resize to a different size, simply return the desired size
// from this method; to avoid resizing, return the current size. `sender`’s
// minimum and maximum size constraints have already been applied when this
// method is called.
// 
// While the user is resizing a window, the delegate is sent a series of
// [WindowWillResizeToSize] messages as the window’s frame continues to
// change size.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowWillResize(_:to:)

func (o NSWindowDelegateObject) WindowWillResizeToSize(sender INSWindow, frameSize corefoundation.CGSize) corefoundation.CGSize {
	
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("windowWillResize:toSize:"), sender, frameSize)
	return rv
	}

// Tells the delegate that the window has been resized.
//
// notification: A notification named [didResizeNotification].
// //
// [didResizeNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didResizeNotification
//
// # Discussion
// 
// You can retrieve the window object in question by sending [object] to
// `notification`.
//
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowDidResize(_:)

func (o NSWindowDelegateObject) WindowDidResize(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowDidResize:"), notification)
	}

// Tells the delegate that the window is about to be live resized.
//
// notification: A notification named [willStartLiveResizeNotification].
// //
// [willStartLiveResizeNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/willStartLiveResizeNotification
//
// # Discussion
// 
// You can retrieve the window object in question by sending [object] to
// `notification`.
//
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowWillStartLiveResize(_:)

func (o NSWindowDelegateObject) WindowWillStartLiveResize(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowWillStartLiveResize:"), notification)
	}

// Tells the delegate that a live resize operation on the window has ended.
//
// notification: A notification named [didEndLiveResizeNotification].
// //
// [didEndLiveResizeNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didEndLiveResizeNotification
//
// # Discussion
// 
// You can retrieve the window object in question by sending [object] to
// `notification`.
//
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowDidEndLiveResize(_:)

func (o NSWindowDelegateObject) WindowDidEndLiveResize(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowDidEndLiveResize:"), notification)
	}

// Tells the delegate that the window is about to be minimized.
//
// notification: A notification named [willMiniaturizeNotification].
// //
// [willMiniaturizeNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/willMiniaturizeNotification
//
// # Discussion
// 
// You can retrieve the [NSWindow] object in question by sending [object] to
// `notification`.
//
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowWillMiniaturize(_:)

func (o NSWindowDelegateObject) WindowWillMiniaturize(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowWillMiniaturize:"), notification)
	}

// Tells the delegate that the window has been minimized.
//
// notification: A notification named [didMiniaturizeNotification].
// //
// [didMiniaturizeNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didMiniaturizeNotification
//
// # Discussion
// 
// You can retrieve the [NSWindow] object in question by sending [object] to
// `notification`.
//
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowDidMiniaturize(_:)

func (o NSWindowDelegateObject) WindowDidMiniaturize(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowDidMiniaturize:"), notification)
	}

// Tells the delegate that the window has been deminimized.
//
// notification: A notification named [didDeminiaturizeNotification]
// //
// [didDeminiaturizeNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didDeminiaturizeNotification
//
// # Discussion
// 
// You can retrieve the [NSWindow] object in question by sending [object] to
// `notification`.
//
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowDidDeminiaturize(_:)

func (o NSWindowDelegateObject) WindowDidDeminiaturize(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowDidDeminiaturize:"), notification)
	}

// Called by [NSWindow]’s [Zoom] method while determining the frame a window
// may be zoomed to.
//
// window: The window whose frame size is being determined.
//
// newFrame: The size of the current screen, which is the screen containing the largest
// part of the window’s current frame, possibly reduced on the top, bottom,
// left, or right, depending on the current interface style. The frame is
// reduced on the top to leave room for the menu bar.
//
// # Return Value
// 
// The specified window’s standard frame.
//
// # Discussion
// 
// The standard frame for a window should supply the size and location that
// are “best” for the type of information shown in the window, taking into
// account the available display or displays. For example, the best width for
// a window that displays a word-processing document is the width of a page or
// the width of the display, whichever is smaller. The best height can be
// determined similarly. On return from this method, the [Zoom] method
// modifies the returned standard frame, if necessary, to fit on the current
// screen.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowWillUseStandardFrame(_:defaultFrame:)

func (o NSWindowDelegateObject) WindowWillUseStandardFrameDefaultFrame(window INSWindow, newFrame corefoundation.CGRect) corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("windowWillUseStandardFrame:defaultFrame:"), window, newFrame)
	return rv
	}

// Asks the delegate whether the specified window should zoom to the specified
// frame.
//
// window: The window being zoomed.
//
// newFrame: The rectangle to which the specified window is being zoomed.
//
// # Return Value
// 
// [true] to allow `window`’s frame to become `newFrame`; otherwise,
// [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowShouldZoom(_:toFrame:)

func (o NSWindowDelegateObject) WindowShouldZoomToFrame(window INSWindow, newFrame corefoundation.CGRect) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("windowShouldZoom:toFrame:"), window, newFrame)
	return rv
	}

// Called to allow the delegate to modify the full-screen content size.
//
// window: The window to enter to full-screen mode.
//
// proposedSize: The proposed window size.
//
// # Return Value
// 
// The window size to use when displaying content size.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/window(_:willUseFullScreenContentSize:)

func (o NSWindowDelegateObject) WindowWillUseFullScreenContentSize(window INSWindow, proposedSize corefoundation.CGSize) corefoundation.CGSize {
	
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("window:willUseFullScreenContentSize:"), window, proposedSize)
	return rv
	}

// Returns the presentation options the window uses when transitioning to
// full-screen mode.
//
// window: The window to enter to full-screen mode.
//
// proposedOptions: The proposed options. See [NSApplication.PresentationOptions] for the
// possible values.
// //
// [NSApplication.PresentationOptions]: https://developer.apple.com/documentation/AppKit/NSApplication/PresentationOptions-swift.struct
//
// # Return Value
// 
// The options the window should use when transitioning to full-screen mode.
// These may be the same as the `proposedOptions` or may be modified.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/window(_:willUseFullScreenPresentationOptions:)

func (o NSWindowDelegateObject) WindowWillUseFullScreenPresentationOptions(window INSWindow, proposedOptions NSApplicationPresentationOptions) NSApplicationPresentationOptions {
	
	rv := objc.Send[NSApplicationPresentationOptions](o.ID, objc.Sel("window:willUseFullScreenPresentationOptions:"), window, proposedOptions)
	return rv
	}

// The window is about to enter full-screen mode.
//
// notification: A notification named [willEnterFullScreenNotification].
// //
// [willEnterFullScreenNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/willEnterFullScreenNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowWillEnterFullScreen(_:)

func (o NSWindowDelegateObject) WindowWillEnterFullScreen(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowWillEnterFullScreen:"), notification)
	}

// The window has entered full-screen mode.
//
// notification: A notification named [didEnterFullScreenNotification].
// //
// [didEnterFullScreenNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didEnterFullScreenNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowDidEnterFullScreen(_:)

func (o NSWindowDelegateObject) WindowDidEnterFullScreen(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowDidEnterFullScreen:"), notification)
	}

// The window is about to exit full-screen mode.
//
// notification: A notification named [willExitFullScreenNotification].
// //
// [willExitFullScreenNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/willExitFullScreenNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowWillExitFullScreen(_:)

func (o NSWindowDelegateObject) WindowWillExitFullScreen(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowWillExitFullScreen:"), notification)
	}

// The window has left full-screen mode.
//
// notification: A notification named [didExitFullScreenNotification].
// //
// [didExitFullScreenNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didExitFullScreenNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowDidExitFullScreen(_:)

func (o NSWindowDelegateObject) WindowDidExitFullScreen(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowDidExitFullScreen:"), notification)
	}

// Called when the window is about to enter full-screen mode.
//
// window: The window to enter full-screen mode.
//
// # Return Value
// 
// An array of windows to use for the animation to full-screen mode for
// `window`; otherwise `nil`.
//
// # Discussion
// 
// This method lets a window delegate customize the animation when the window
// is about to enter full-screen mode by providing a custom window or windows
// containing layers or other effects. If you don’t want to perform custom
// animation, you can omit the implementation of this method, or it can return
// `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/customWindowsToEnterFullScreen(for:)

func (o NSWindowDelegateObject) CustomWindowsToEnterFullScreenForWindow(window INSWindow) []NSWindow {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("customWindowsToEnterFullScreenForWindow:"), window)
	return objc.ConvertSlice(rv, func(id objc.ID) NSWindow {
		return NSWindowFromID(id)
	})
	}

// Called when the window is about to enter full-screen mode.
//
// window: The window to enter full-screen mode.
//
// screen: The display screen on which the window will enter full-screen mode.
//
// # Return Value
// 
// An array of windows to use for the animation to full-screen mode for
// `window`; otherwise `nil`.
//
// # Discussion
// 
// This method lets a window delegate customize the animation when the window
// is about to enter full-screen mode by providing a custom window or windows
// containing layers or other effects. If you don’t want to perform custom
// animation, you can omit the implementation of this method, or it can return
// `nil`.
// 
// # Special Considerations
// 
// If this method and [CustomWindowsToEnterFullScreenForWindow] are both
// implemented, this method is called.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/customWindowsToEnterFullScreen(for:on:)

func (o NSWindowDelegateObject) CustomWindowsToEnterFullScreenForWindowOnScreen(window INSWindow, screen INSScreen) []NSWindow {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("customWindowsToEnterFullScreenForWindow:onScreen:"), window, screen)
	return objc.ConvertSlice(rv, func(id objc.ID) NSWindow {
		return NSWindowFromID(id)
	})
	}

// This method is called to start the window animation into full-screen mode,
// including transitioning to a new space.
//
// window: The window to enter full-screen mode.
//
// duration: The duration of the presentation change.
//
// # Discussion
// 
// You can implement this method to perform custom animation with the given
// duration to be in sync with the system animation.
// 
// # Special Considerations
// 
// This method is called only if [CustomWindowsToEnterFullScreenForWindow]
// returns non-`nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/window(_:startCustomAnimationToEnterFullScreenWithDuration:)

func (o NSWindowDelegateObject) WindowStartCustomAnimationToEnterFullScreenWithDuration(window INSWindow, duration float64) {
	
	objc.Send[struct{}](o.ID, objc.Sel("window:startCustomAnimationToEnterFullScreenWithDuration:"), window, duration)
	}

// This method is called to start the window animation into full-screen mode,
// including transitioning to a new space.
//
// window: The window to enter full-screen mode.
//
// screen: The display screen on which the window will enter full-screen mode.
//
// duration: The duration of the presentation change.
//
// # Discussion
// 
// You can implement this method to perform custom animation with the given
// duration to be in sync with the system animation.
// 
// # Special Considerations
// 
// This method is called only if [CustomWindowsToEnterFullScreenForWindow]
// returns non-`nil`. If
// [WindowStartCustomAnimationToEnterFullScreenWithDuration] and this method
// are both implemented, this method is called.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/window(_:startCustomAnimationToEnterFullScreenOn:withDuration:)

func (o NSWindowDelegateObject) WindowStartCustomAnimationToEnterFullScreenOnScreenWithDuration(window INSWindow, screen INSScreen, duration float64) {
	
	objc.Send[struct{}](o.ID, objc.Sel("window:startCustomAnimationToEnterFullScreenOnScreen:withDuration:"), window, screen, duration)
	}

// Called if the window failed to enter full-screen mode.
//
// window: The window that failed to enter to full-screen mode.
//
// # Discussion
// 
// In some cases, the transition to enter full-screen mode can fail, due to
// being in the midst of handling some other animation or user gesture. This
// method indicates that there was an error, and you should clean up any work
// you may have done to prepare to enter full-screen mode.
// 
// This message is sent whether or not the delegate indicated a custom
// animation by returning non-`nil` from
// [CustomWindowsToEnterFullScreenForWindow].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowDidFailToEnterFullScreen(_:)

func (o NSWindowDelegateObject) WindowDidFailToEnterFullScreen(window INSWindow) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowDidFailToEnterFullScreen:"), window)
	}

// Called when the window is about to exit full-screen mode.
//
// window: The window to exit full-screen mode.
//
// # Return Value
// 
// An array of windows involved in the animation out of full-screen mode for
// `window`; otherwise `nil`.
//
// # Discussion
// 
// This method lets the window delegate customize the animation when the
// window is about to exit full-screen mode by providing a custom window or
// windows containing layers or other effects. If an you do not want to
// perform custom animation, you can omit the implementation of this method,
// or it can return `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/customWindowsToExitFullScreen(for:)

func (o NSWindowDelegateObject) CustomWindowsToExitFullScreenForWindow(window INSWindow) []NSWindow {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("customWindowsToExitFullScreenForWindow:"), window)
	return objc.ConvertSlice(rv, func(id objc.ID) NSWindow {
		return NSWindowFromID(id)
	})
	}

// This method is called to start the window animation out of full-screen
// mode, including transitioning back to the desktop space.
//
// window: The window to exit to full-screen mode.
//
// duration: The duration of the presentation change.
//
// # Discussion
// 
// You can implement this method to perform custom animation with the given
// duration to be in sync with the system animation.
// 
// # Special Considerations
// 
// This method is called only if [CustomWindowsToExitFullScreenForWindow]
// returns non-`nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/window(_:startCustomAnimationToExitFullScreenWithDuration:)

func (o NSWindowDelegateObject) WindowStartCustomAnimationToExitFullScreenWithDuration(window INSWindow, duration float64) {
	
	objc.Send[struct{}](o.ID, objc.Sel("window:startCustomAnimationToExitFullScreenWithDuration:"), window, duration)
	}

// Called if the window failed to exit full-screen mode.
//
// window: The window that failed to exit to full-screen mode.
//
// # Discussion
// 
// In some cases, the transition to exit full-screen mode can fail, due to
// being in the midst of handling some other animation or user gesture. This
// method indicates that there was an error, and you should clean up any work
// you may have done to prepare to exit full-screen mode.
// 
// This message is sent whether or not the delegate indicated a custom
// animation by returning non-`nil` from
// [CustomWindowsToExitFullScreenForWindow].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowDidFailToExitFullScreen(_:)

func (o NSWindowDelegateObject) WindowDidFailToExitFullScreen(window INSWindow) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowDidFailToExitFullScreen:"), window)
	}

// Tells the delegate that the window is about to move.
//
// notification: A notification named [willMoveNotification].
// //
// [willMoveNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/willMoveNotification
//
// # Discussion
// 
// You can retrieve the [NSWindow] object in question by sending [object] to
// `notification`.
//
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowWillMove(_:)

func (o NSWindowDelegateObject) WindowWillMove(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowWillMove:"), notification)
	}

// Tells the delegate that the window has moved.
//
// notification: A notification named [didMoveNotification].
// //
// [didMoveNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didMoveNotification
//
// # Discussion
// 
// You can retrieve the [NSWindow] object in question by sending [object] to
// `notification`.
//
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowDidMove(_:)

func (o NSWindowDelegateObject) WindowDidMove(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowDidMove:"), notification)
	}

// Tells the delegate that the window has changed screens.
//
// notification: A notification named [didChangeScreenNotification].
// //
// [didChangeScreenNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didChangeScreenNotification
//
// # Discussion
// 
// You can retrieve the [NSWindow] object in question by sending [object] to
// `notification`.
//
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowDidChangeScreen(_:)

func (o NSWindowDelegateObject) WindowDidChangeScreen(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowDidChangeScreen:"), notification)
	}

// Tells the delegate that the window has changed screen display profiles.
//
// notification: A notification named [didChangeScreenProfileNotification].
// //
// [didChangeScreenProfileNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didChangeScreenProfileNotification
//
// # Discussion
// 
// You can retrieve the [NSWindow] object in question by sending [object] to
// `notification`.
// 
// If your app runs in macOS 10.7.3 or later, you should instead watch for the
// notification [NSWindowDidChangeBackingPropertiesNotification].
//
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowDidChangeScreenProfile(_:)

func (o NSWindowDelegateObject) WindowDidChangeScreenProfile(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowDidChangeScreenProfile:"), notification)
	}

// Tells the delegate that the window backing properties changed.
//
// notification: A notification named [NSWindowDidChangeBackingPropertiesNotification].
//
// # Discussion
// 
// You can retrieve the [NSWindow] object in question by sending [object] to
// `notification`.
// 
// The notification [NSWindowDidChangeBackingPropertiesNotification] is posted
// in macOS 10.7.3 or later when a window’s backing scale factor or its
// color space changes. You should watch for this notification instead of
// [NSWindowDidChangeScreenProfileNotification] if your app runs on a system
// version on which the backing properties notification is available.
// 
// Many apps won’t have the need to watch for this notification, but those
// that perform sophisticated color handling or manually manage their own
// cache of window-resolution or color-space-appropriate bitmapped images will
// find this notification useful as a prompt to invalidate caches or schedule
// other reassessment for the new resolution or color space. The
// notification’s `userInfo` dictionary specifies the window’s previous
// backing scale factor (retrieved with the
// key[NSBackingPropertyOldScaleFactorKey]) and color space (retrieved with
// the key [NSBackingPropertyOldColorSpaceKey]). You can compare these with
// the window’s new previous backing scale factor and color space at the
// time of the notification to determine which properties changed.
//
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowDidChangeBackingProperties(_:)

func (o NSWindowDelegateObject) WindowDidChangeBackingProperties(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowDidChangeBackingProperties:"), notification)
	}

// Tells the delegate that the user has attempted to close a window or the
// window has received a [PerformClose] message.
//
// sender: The window being closed.
//
// # Return Value
// 
// [true] to allow `sender` to be closed; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method may not always be called during window closing. Specifically,
// this method is not called when a user quits an application.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowShouldClose(_:)

func (o NSWindowDelegateObject) WindowShouldClose(sender INSWindow) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("windowShouldClose:"), sender)
	return rv
	}

// Tells the delegate that the window is about to close.
//
// notification: A notification named [willCloseNotification].
// //
// [willCloseNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/willCloseNotification
//
// # Discussion
// 
// You can retrieve the [NSWindow] object in question by sending [object] to
// `notification`.
//
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowWillClose(_:)

func (o NSWindowDelegateObject) WindowWillClose(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowWillClose:"), notification)
	}

// Tells the delegate that the window has become the key window.
//
// notification: A notification named [didBecomeKeyNotification].
// //
// [didBecomeKeyNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didBecomeKeyNotification
//
// # Discussion
// 
// You can retrieve the window object in question by sending [object] to
// `notification`.
//
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowDidBecomeKey(_:)

func (o NSWindowDelegateObject) WindowDidBecomeKey(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowDidBecomeKey:"), notification)
	}

// Tells the delegate that the window has resigned key window status.
//
// notification: A notification named [didResignKeyNotification].
// //
// [didResignKeyNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didResignKeyNotification
//
// # Discussion
// 
// You can retrieve the window object in question by sending [object] to
// `notification`.
//
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowDidResignKey(_:)

func (o NSWindowDelegateObject) WindowDidResignKey(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowDidResignKey:"), notification)
	}

// Tells the delegate that the window has become main.
//
// notification: A notification named [didBecomeMainNotification].
// //
// [didBecomeMainNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didBecomeMainNotification
//
// # Discussion
// 
// You can retrieve the window object in question by sending [object] to
// `notification`.
//
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowDidBecomeMain(_:)

func (o NSWindowDelegateObject) WindowDidBecomeMain(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowDidBecomeMain:"), notification)
	}

// Tells the delegate that the window has resigned main window status.
//
// notification: A notification named [didResignMainNotification].
// //
// [didResignMainNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didResignMainNotification
//
// # Discussion
// 
// You can retrieve the window object in question by sending [object] to
// `notification`.
//
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowDidResignMain(_:)

func (o NSWindowDelegateObject) WindowDidResignMain(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowDidResignMain:"), notification)
	}

// Tells the delegate that the field editor for a text-displaying object has
// been requested.
//
// sender: The window requesting the field editor from the delegate.
//
// client: A text-displaying object to be associated with the field editor. If `nil`,
// the requested field editor is the default.
//
// # Return Value
// 
// The field editor for `client`; returns `nil` when the delegate has no field
// editor to assign.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowWillReturnFieldEditor(_:to:)

func (o NSWindowDelegateObject) WindowWillReturnFieldEditorToObject(sender INSWindow, client objectivec.IObject) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("windowWillReturnFieldEditor:toObject:"), sender, client)
	return objectivec.Object{ID: rv}
	}

// Tells the delegate that the window received an [Update] message.
//
// notification: A notification named [didUpdateNotification]
// //
// [didUpdateNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didUpdateNotification
//
// # Discussion
// 
// You can retrieve the window object in question by sending [object] to
// `notification`.
//
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowDidUpdate(_:)

func (o NSWindowDelegateObject) WindowDidUpdate(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowDidUpdate:"), notification)
	}

// Tells the delegate that the window has been exposed.
//
// notification: A notification named [didExposeNotification].
// //
// [didExposeNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didExposeNotification
//
// # Discussion
// 
// You can retrieve the window object in question by sending [object] to
// `notification`.
//
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowDidExpose(_:)

func (o NSWindowDelegateObject) WindowDidExpose(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowDidExpose:"), notification)
	}

// Tells the delegate that the window changed its occlusion state.
//
// notification: An [didChangeOcclusionStateNotification] notification.
// //
// [didChangeOcclusionStateNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didChangeOcclusionStateNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowDidChangeOcclusionState(_:)

func (o NSWindowDelegateObject) WindowDidChangeOcclusionState(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowDidChangeOcclusionState:"), notification)
	}

// Asks the delegate whether a user can drag the document icon from the
// window’s title bar.
//
// window: The window containing the document icon the user wants to drag.
//
// event: The left-mouse down event that triggered the dragging operation.
//
// dragImageLocation: The location of the origin of the document icon, in window coordinates,
// when the user started the dragging operation.
//
// pasteboard: The pasteboard containing the contents of the document, which the delegate
// can modify.
//
// # Return Value
// 
// [true] to allow the drag to proceed; [false] to prevent it. Before turning
// no the delegate can implement its own dragging behavior as described below.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Implementing this method enables an application to customize the process of
// dragging the window’s document icon. The delegate can prohibit the drag
// by returning [false]. Before returning [false], the delegate can implement
// its own dragging behavior using
// [DragImageAtOffsetEventPasteboardSourceSlideBack].
// 
// Alternatively, the delegate can enable a drag by returning [true], for
// example, to override the default [NSWindow] behavior of prohibiting the
// drag of an edited document. In addition, the delegate can customize the
// pasteboard contents before returning [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/window(_:shouldDragDocumentWith:from:with:)

func (o NSWindowDelegateObject) WindowShouldDragDocumentWithEventFromWithPasteboard(window INSWindow, event INSEvent, dragImageLocation corefoundation.CGPoint, pasteboard INSPasteboard) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("window:shouldDragDocumentWithEvent:from:withPasteboard:"), window, event, dragImageLocation, pasteboard)
	return rv
	}

// Tells the delegate that the window’s undo manager has been requested.
// Returns the appropriate undo manager for the window.
//
// window: The window whose undo manager is being requested.
//
// # Return Value
// 
// The appropriate undo manager for the specified window.
//
// # Discussion
// 
// If this method is not implemented by the delegate, the window creates
// an[UndoManager] for `window`. Further, after a window creates its own undo
// manager, this method is never again called on the delegate.
//
// [UndoManager]: https://developer.apple.com/documentation/Foundation/UndoManager
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowWillReturnUndoManager(_:)

func (o NSWindowDelegateObject) WindowWillReturnUndoManager(window INSWindow) foundation.NSUndoManager {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("windowWillReturnUndoManager:"), window)
	return foundation.NSUndoManagerFromID(rv)
	}

// Asks the delegate whether the window displays the title pop-up menu in
// response to a Command-click or Control-click on its title.
//
// window: The window whose title the user Command-clicked or Control-clicked.
//
// menu: The menu the window will display, if allowed. By default, its items are the
// path components of the file represented by `window`.
//
// # Return Value
// 
// [true] to allow the display of the title pop-up menu; [false] to prevent
// it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/window(_:shouldPopUpDocumentPathMenu:)

func (o NSWindowDelegateObject) WindowShouldPopUpDocumentPathMenu(window INSWindow, menu INSMenu) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("window:shouldPopUpDocumentPathMenu:"), window, menu)
	return rv
	}

// Tells the delegate the window is about to add its restorable state to a
// given archiver.
//
// window: The window adding its restorable state to an archive.
//
// state: The coder creating the archive.
//
// # Discussion
// 
// This method is called during the window’s
// [EncodeRestorableStateWithCoder] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/window(_:willEncodeRestorableState:)

func (o NSWindowDelegateObject) WindowWillEncodeRestorableState(window INSWindow, state foundation.INSCoder) {
	
	objc.Send[struct{}](o.ID, objc.Sel("window:willEncodeRestorableState:"), window, state)
	}

// Tells the delegate the window is has extracted its restorable state from a
// given archiver.
//
// window: The window extracting its restorable state from an archive.
//
// state: The coder extracting the archive.
//
// # Discussion
// 
// This method is called during the window’s [RestoreStateWithCoder] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/window(_:didDecodeRestorableState:)

func (o NSWindowDelegateObject) WindowDidDecodeRestorableState(window INSWindow, state foundation.INSCoder) {
	
	objc.Send[struct{}](o.ID, objc.Sel("window:didDecodeRestorableState:"), window, state)
	}

// Tells the delegate the window will resize for presentation during version
// browsing.
//
// window: The window being presented in a version browser.
//
// maxPreferredFrameSize: The maximum size the version browser would prefer the window to be.
//
// maxAllowedFrameSize: The maximum allowed size for the window (the full-screen frame minus the
// margins required to ensure the Versions controls are still visible).
//
// # Return Value
// 
// The size that the window should be.
//
// # Discussion
// 
// Windows entering the version browser will be resized to the size returned
// by this method. If either dimension of the returned size is larger than the
// `maxPreferredFrameSize`, the window will also be scaled down to ensure it
// fits properly in the version browser.
// 
// If this method is not implemented, the version browser will use
// [WindowWillUseStandardFrameDefaultFrame] to determine the resulting window
// frame size.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/window(_:willResizeForVersionBrowserWithMaxPreferredSize:maxAllowedSize:)

func (o NSWindowDelegateObject) WindowWillResizeForVersionBrowserWithMaxPreferredSizeMaxAllowedSize(window INSWindow, maxPreferredFrameSize corefoundation.CGSize, maxAllowedFrameSize corefoundation.CGSize) corefoundation.CGSize {
	
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("window:willResizeForVersionBrowserWithMaxPreferredSize:maxAllowedSize:"), window, maxPreferredFrameSize, maxAllowedFrameSize)
	return rv
	}

// Tells the delegate the window is about to enter version browsing.
//
// notification: An [willEnterVersionBrowserNotification] notification.
// //
// [willEnterVersionBrowserNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/willEnterVersionBrowserNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowWillEnterVersionBrowser(_:)

func (o NSWindowDelegateObject) WindowWillEnterVersionBrowser(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowWillEnterVersionBrowser:"), notification)
	}

// Tells the delegate that the window has entered version browsing.
//
// notification: An [didEnterVersionBrowserNotification] notification.
// //
// [didEnterVersionBrowserNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didEnterVersionBrowserNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowDidEnterVersionBrowser(_:)

func (o NSWindowDelegateObject) WindowDidEnterVersionBrowser(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowDidEnterVersionBrowser:"), notification)
	}

// Tells the delegate that the window is about to leave version browsing.
//
// notification: An [willExitVersionBrowserNotification] notification.
// //
// [willExitVersionBrowserNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/willExitVersionBrowserNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowWillExitVersionBrowser(_:)

func (o NSWindowDelegateObject) WindowWillExitVersionBrowser(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowWillExitVersionBrowser:"), notification)
	}

// Tells the delegate that the window has left version browsing.
//
// notification: An [didExitVersionBrowserNotification] notification.
// //
// [didExitVersionBrowserNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didExitVersionBrowserNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowDidExitVersionBrowser(_:)

func (o NSWindowDelegateObject) WindowDidExitVersionBrowser(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("windowDidExitVersionBrowser:"), notification)
	}

//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/previewRepresentableActivityItems(for:)

func (o NSWindowDelegateObject) PreviewRepresentableActivityItemsForWindow(window INSWindow) []objectivec.IObject {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("previewRepresentableActivityItemsForWindow:"), window)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
	}

// Method called to get the window to share once sharing is confirmed, after a
// request is initiated by
// requestSharingOfWindowUsingPreview:title:completionHandler:. Implement this
// on the delegate of the requesting window
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowDelegate/windowForSharingRequest(from:)

func (o NSWindowDelegateObject) WindowForSharingRequestFromWindow(window INSWindow) INSWindow {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("windowForSharingRequestFromWindow:"), window)
	return NSWindowFromID(rv)
	}





// NSWindowDelegateConfig holds optional typed callbacks for [NSWindowDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nswindowdelegate
type NSWindowDelegateConfig struct {

	// Managing Sheets
	// WillBeginSheet — Notifies the delegate that the window is about to open a sheet.
	WillBeginSheet func(notification foundation.NSNotification)
	// DidEndSheet — Tells the delegate that the window has closed a sheet.
	DidEndSheet func(notification foundation.NSNotification)

	// Sizing Windows
	// DidResize — Tells the delegate that the window has been resized.
	DidResize func(notification foundation.NSNotification)
	// WillStartLiveResize — Tells the delegate that the window is about to be live resized.
	WillStartLiveResize func(notification foundation.NSNotification)
	// DidEndLiveResize — Tells the delegate that a live resize operation on the window has ended.
	DidEndLiveResize func(notification foundation.NSNotification)

	// Minimizing Windows
	// WillMiniaturize — Tells the delegate that the window is about to be minimized.
	WillMiniaturize func(notification foundation.NSNotification)
	// DidMiniaturize — Tells the delegate that the window has been minimized.
	DidMiniaturize func(notification foundation.NSNotification)
	// DidDeminiaturize — Tells the delegate that the window has been deminimized.
	DidDeminiaturize func(notification foundation.NSNotification)

	// Managing Full-Screen Presentation
	// WillUseFullScreenPresentationOptions — Returns the presentation options the window uses when transitioning to full-screen mode.
	WillUseFullScreenPresentationOptions func(window NSWindow, proposedOptions NSApplicationPresentationOptions) NSApplicationPresentationOptions
	// WillEnterFullScreen — The window is about to enter full-screen mode.
	WillEnterFullScreen func(notification foundation.NSNotification)
	// DidEnterFullScreen — The window has entered full-screen mode.
	DidEnterFullScreen func(notification foundation.NSNotification)
	// WillExitFullScreen — The window is about to exit full-screen mode.
	WillExitFullScreen func(notification foundation.NSNotification)
	// DidExitFullScreen — The window has left full-screen mode.
	DidExitFullScreen func(notification foundation.NSNotification)

	// Custom Full-Screen Presentation Animations
	// StartCustomAnimationToEnterFullScreenWithDuration — This method is called to start the window animation into full-screen mode, including transitioning to a new space.
	StartCustomAnimationToEnterFullScreenWithDuration func(window NSWindow, duration float64)
	// DidFailToEnterFullScreen — Called if the window failed to enter full-screen mode.
	DidFailToEnterFullScreen func(window NSWindow)
	// StartCustomAnimationToExitFullScreenWithDuration — This method is called to start the window animation out of full-screen mode, including transitioning back to the desktop space.
	StartCustomAnimationToExitFullScreenWithDuration func(window NSWindow, duration float64)
	// DidFailToExitFullScreen — Called if the window failed to exit full-screen mode.
	DidFailToExitFullScreen func(window NSWindow)

	// Moving Windows
	// WillMove — Tells the delegate that the window is about to move.
	WillMove func(notification foundation.NSNotification)
	// DidMove — Tells the delegate that the window has moved.
	DidMove func(notification foundation.NSNotification)
	// DidChangeScreen — Tells the delegate that the window has changed screens.
	DidChangeScreen func(notification foundation.NSNotification)
	// DidChangeScreenProfile — Tells the delegate that the window has changed screen display profiles.
	DidChangeScreenProfile func(notification foundation.NSNotification)
	// DidChangeBackingProperties — Tells the delegate that the window backing properties changed.
	DidChangeBackingProperties func(notification foundation.NSNotification)

	// Closing Windows
	// ShouldClose — Tells the delegate that the user has attempted to close a window or the window has received a [performClose(_:)](<doc://com.apple.appkit/documentation/AppKit/NSWindow/performClose(_:)>) message.
	ShouldClose func(sender NSWindow) bool
	// WillClose — Tells the delegate that the window is about to close.
	WillClose func(notification foundation.NSNotification)

	// Managing Key Status
	// DidBecomeKey — Tells the delegate that the window has become the key window.
	DidBecomeKey func(notification foundation.NSNotification)
	// DidResignKey — Tells the delegate that the window has resigned key window status.
	DidResignKey func(notification foundation.NSNotification)

	// Managing Main Status
	// DidBecomeMain — Tells the delegate that the window has become main.
	DidBecomeMain func(notification foundation.NSNotification)
	// DidResignMain — Tells the delegate that the window has resigned main window status.
	DidResignMain func(notification foundation.NSNotification)

	// Updating Windows
	// DidUpdate — Tells the delegate that the window received an [update()](<doc://com.apple.appkit/documentation/AppKit/NSWindow/update()>) message.
	DidUpdate func(notification foundation.NSNotification)

	// Exposing Windows
	// DidExpose — Tells the delegate that the window has been exposed.
	DidExpose func(notification foundation.NSNotification)

	// Managing Occlusion State
	// DidChangeOcclusionState — Tells the delegate that the window changed its occlusion state.
	DidChangeOcclusionState func(notification foundation.NSNotification)

	// Getting the Undo Manager
	// WillReturnUndoManager — Tells the delegate that the window’s undo manager has been requested. Returns the appropriate undo manager for the window.
	WillReturnUndoManager func(window NSWindow) foundation.NSUndoManager

	// Managing Titles
	// ShouldPopUpDocumentPathMenu — Asks the delegate whether the window displays the title pop-up menu in response to a Command-click or Control-click on its title.
	ShouldPopUpDocumentPathMenu func(window NSWindow, menu NSMenu) bool

	// Managing Restorable State
	// WillEncodeRestorableState — Tells the delegate the window is about to add its restorable state to a given archiver.
	WillEncodeRestorableState func(window NSWindow, state foundation.INSCoder)
	// DidDecodeRestorableState — Tells the delegate the window is has extracted its restorable state from a given archiver.
	DidDecodeRestorableState func(window NSWindow, state foundation.INSCoder)

	// Managing Presentation in Version Browsers
	// WillEnterVersionBrowser — Tells the delegate the window is about to enter version browsing.
	WillEnterVersionBrowser func(notification foundation.NSNotification)
	// DidEnterVersionBrowser — Tells the delegate that the window has entered version browsing.
	DidEnterVersionBrowser func(notification foundation.NSNotification)
	// WillExitVersionBrowser — Tells the delegate that the window is about to leave version browsing.
	WillExitVersionBrowser func(notification foundation.NSNotification)
	// DidExitVersionBrowser — Tells the delegate that the window has left version browsing.
	DidExitVersionBrowser func(notification foundation.NSNotification)

	// Other Methods
	// StartCustomAnimationToEnterFullScreenOnScreenWithDuration — This method is called to start the window animation into full-screen mode, including transitioning to a new space.
	StartCustomAnimationToEnterFullScreenOnScreenWithDuration func(window NSWindow, screen NSScreen, duration float64)
	// ForSharingRequestFromWindow — Method called to get the window to share once sharing is confirmed, after a request is initiated by requestSharingOfWindowUsingPreview:title:completionHandler:. Implement this on the delegate of the requesting window
	ForSharingRequestFromWindow func(window NSWindow) NSWindow
}

// NewNSWindowDelegate creates an Objective-C object implementing the [NSWindowDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSWindowDelegateObject] satisfies the [NSWindowDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nswindowdelegate
func NewNSWindowDelegate(config NSWindowDelegateConfig) NSWindowDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSWindowDelegate_%d", n)

	var methods []objc.MethodDef

	if config.WillBeginSheet != nil {
		fn := config.WillBeginSheet
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowWillBeginSheet:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidEndSheet != nil {
		fn := config.DidEndSheet
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowDidEndSheet:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidResize != nil {
		fn := config.DidResize
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowDidResize:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.WillStartLiveResize != nil {
		fn := config.WillStartLiveResize
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowWillStartLiveResize:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidEndLiveResize != nil {
		fn := config.DidEndLiveResize
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowDidEndLiveResize:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.WillMiniaturize != nil {
		fn := config.WillMiniaturize
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowWillMiniaturize:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidMiniaturize != nil {
		fn := config.DidMiniaturize
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowDidMiniaturize:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidDeminiaturize != nil {
		fn := config.DidDeminiaturize
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowDidDeminiaturize:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.WillUseFullScreenPresentationOptions != nil {
		fn := config.WillUseFullScreenPresentationOptions
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("window:willUseFullScreenPresentationOptions:"),
			Fn: func(self objc.ID, _cmd objc.SEL, windowID objc.ID, proposedOptions NSApplicationPresentationOptions) NSApplicationPresentationOptions {
				window := NSWindowFromID(windowID)
				return fn(window, proposedOptions)
			},
		})
	}

	if config.WillEnterFullScreen != nil {
		fn := config.WillEnterFullScreen
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowWillEnterFullScreen:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidEnterFullScreen != nil {
		fn := config.DidEnterFullScreen
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowDidEnterFullScreen:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.WillExitFullScreen != nil {
		fn := config.WillExitFullScreen
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowWillExitFullScreen:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidExitFullScreen != nil {
		fn := config.DidExitFullScreen
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowDidExitFullScreen:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.StartCustomAnimationToEnterFullScreenWithDuration != nil {
		fn := config.StartCustomAnimationToEnterFullScreenWithDuration
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("window:startCustomAnimationToEnterFullScreenWithDuration:"),
			Fn: func(self objc.ID, _cmd objc.SEL, windowID objc.ID, duration float64) {
				window := NSWindowFromID(windowID)
				fn(window, duration)
			},
		})
	}

	if config.StartCustomAnimationToEnterFullScreenOnScreenWithDuration != nil {
		fn := config.StartCustomAnimationToEnterFullScreenOnScreenWithDuration
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("window:startCustomAnimationToEnterFullScreenOnScreen:withDuration:"),
			Fn: func(self objc.ID, _cmd objc.SEL, windowID objc.ID, screenID objc.ID, duration float64) {
				window := NSWindowFromID(windowID)
				screen := NSScreenFromID(screenID)
				fn(window, screen, duration)
			},
		})
	}

	if config.DidFailToEnterFullScreen != nil {
		fn := config.DidFailToEnterFullScreen
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowDidFailToEnterFullScreen:"),
			Fn: func(self objc.ID, _cmd objc.SEL, windowID objc.ID) {
				window := NSWindowFromID(windowID)
				fn(window)
			},
		})
	}

	if config.StartCustomAnimationToExitFullScreenWithDuration != nil {
		fn := config.StartCustomAnimationToExitFullScreenWithDuration
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("window:startCustomAnimationToExitFullScreenWithDuration:"),
			Fn: func(self objc.ID, _cmd objc.SEL, windowID objc.ID, duration float64) {
				window := NSWindowFromID(windowID)
				fn(window, duration)
			},
		})
	}

	if config.DidFailToExitFullScreen != nil {
		fn := config.DidFailToExitFullScreen
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowDidFailToExitFullScreen:"),
			Fn: func(self objc.ID, _cmd objc.SEL, windowID objc.ID) {
				window := NSWindowFromID(windowID)
				fn(window)
			},
		})
	}

	if config.WillMove != nil {
		fn := config.WillMove
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowWillMove:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidMove != nil {
		fn := config.DidMove
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowDidMove:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidChangeScreen != nil {
		fn := config.DidChangeScreen
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowDidChangeScreen:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidChangeScreenProfile != nil {
		fn := config.DidChangeScreenProfile
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowDidChangeScreenProfile:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidChangeBackingProperties != nil {
		fn := config.DidChangeBackingProperties
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowDidChangeBackingProperties:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.ShouldClose != nil {
		fn := config.ShouldClose
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowShouldClose:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID) bool {
				sender := NSWindowFromID(senderID)
				return fn(sender)
			},
		})
	}

	if config.WillClose != nil {
		fn := config.WillClose
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowWillClose:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidBecomeKey != nil {
		fn := config.DidBecomeKey
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowDidBecomeKey:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidResignKey != nil {
		fn := config.DidResignKey
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowDidResignKey:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidBecomeMain != nil {
		fn := config.DidBecomeMain
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowDidBecomeMain:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidResignMain != nil {
		fn := config.DidResignMain
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowDidResignMain:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidUpdate != nil {
		fn := config.DidUpdate
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowDidUpdate:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidExpose != nil {
		fn := config.DidExpose
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowDidExpose:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidChangeOcclusionState != nil {
		fn := config.DidChangeOcclusionState
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowDidChangeOcclusionState:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.WillReturnUndoManager != nil {
		fn := config.WillReturnUndoManager
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowWillReturnUndoManager:"),
			Fn: func(self objc.ID, _cmd objc.SEL, windowID objc.ID) objc.ID {
				window := NSWindowFromID(windowID)
				return fn(window).GetID()
			},
		})
	}

	if config.ShouldPopUpDocumentPathMenu != nil {
		fn := config.ShouldPopUpDocumentPathMenu
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("window:shouldPopUpDocumentPathMenu:"),
			Fn: func(self objc.ID, _cmd objc.SEL, windowID objc.ID, menuID objc.ID) bool {
				window := NSWindowFromID(windowID)
				menu := NSMenuFromID(menuID)
				return fn(window, menu)
			},
		})
	}

	if config.WillEncodeRestorableState != nil {
		fn := config.WillEncodeRestorableState
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("window:willEncodeRestorableState:"),
			Fn: func(self objc.ID, _cmd objc.SEL, windowID objc.ID, stateID objc.ID) {
				window := NSWindowFromID(windowID)
				state := foundation.NSCoderFromID(stateID)
				fn(window, state)
			},
		})
	}

	if config.DidDecodeRestorableState != nil {
		fn := config.DidDecodeRestorableState
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("window:didDecodeRestorableState:"),
			Fn: func(self objc.ID, _cmd objc.SEL, windowID objc.ID, stateID objc.ID) {
				window := NSWindowFromID(windowID)
				state := foundation.NSCoderFromID(stateID)
				fn(window, state)
			},
		})
	}

	if config.WillEnterVersionBrowser != nil {
		fn := config.WillEnterVersionBrowser
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowWillEnterVersionBrowser:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidEnterVersionBrowser != nil {
		fn := config.DidEnterVersionBrowser
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowDidEnterVersionBrowser:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.WillExitVersionBrowser != nil {
		fn := config.WillExitVersionBrowser
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowWillExitVersionBrowser:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidExitVersionBrowser != nil {
		fn := config.DidExitVersionBrowser
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowDidExitVersionBrowser:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.ForSharingRequestFromWindow != nil {
		fn := config.ForSharingRequestFromWindow
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("windowForSharingRequestFromWindow:"),
			Fn: func(self objc.ID, _cmd objc.SEL, windowID objc.ID) objc.ID {
				window := NSWindowFromID(windowID)
				return fn(window).GetID()
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSWindowDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSWindowDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSWindowDelegateObjectFromID(instance)
}






