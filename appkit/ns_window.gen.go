// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/quartzcore"
)

// The class instance for the [NSWindow] class.
var (
	_NSWindowClass     NSWindowClass
	_NSWindowClassOnce sync.Once
)

func getNSWindowClass() NSWindowClass {
	_NSWindowClassOnce.Do(func() {
		_NSWindowClass = NSWindowClass{class: objc.GetClass("NSWindow")}
	})
	return _NSWindowClass
}

// GetNSWindowClass returns the class object for NSWindow.
func GetNSWindowClass() NSWindowClass {
	return getNSWindowClass()
}

type NSWindowClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSWindowClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSWindowClass) Alloc() NSWindow {
	rv := objc.Send[NSWindow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A window that an app displays on the screen.
//
// # Overview
//
// A single [NSWindow] object corresponds to, at most, one on-screen window.
// Windows perform two principal functions:
//
// - To place views in a provided area - To accept and distribute mouse and
// keyboard events the user generates to the appropriate views
//
// # Creating a Window
//
//   - [NSWindow.InitWithContentRectStyleMaskBackingDefer]: Initializes the window with the specified values.
//   - [NSWindow.InitWithContentRectStyleMaskBackingDeferScreen]: Initializes an allocated window with the specified values.
//
// # Managing the Window’s Behavior
//
//   - [NSWindow.Delegate]: The window’s delegate.
//   - [NSWindow.SetDelegate]
//
// # Configuring the Window’s Content
//
//   - [NSWindow.ContentViewController]: The main content view controller for the window.
//   - [NSWindow.SetContentViewController]
//   - [NSWindow.ContentView]: The window’s content view, the highest accessible view object in the window’s view hierarchy.
//   - [NSWindow.SetContentView]
//
// # Configuring the Window’s Appearance
//
//   - [NSWindow.StyleMask]: Flags that describe the window’s current style, such as if it’s resizable or in full-screen mode.
//   - [NSWindow.SetStyleMask]
//   - [NSWindow.ToggleFullScreen]: Takes the window into or out of fullscreen mode,
//   - [NSWindow.WorksWhenModal]: A Boolean value that indicates whether the window is able to receive keyboard and mouse events even when some other window is being run modally.
//   - [NSWindow.AlphaValue]: The window’s alpha value.
//   - [NSWindow.SetAlphaValue]
//   - [NSWindow.BackgroundColor]: The color of the window’s background.
//   - [NSWindow.SetBackgroundColor]
//   - [NSWindow.ColorSpace]: The window’s color space.
//   - [NSWindow.SetColorSpace]
//   - [NSWindow.SetDynamicDepthLimit]: Sets a Boolean value that indicates whether the window’s depth limit can change to match the depth of the screen it’s on.
//   - [NSWindow.CanHide]: A Boolean value that indicates whether the window can hide when its application becomes hidden.
//   - [NSWindow.SetCanHide]
//   - [NSWindow.OnActiveSpace]: A Boolean value that indicates whether the window is on the currently active space.
//   - [NSWindow.HidesOnDeactivate]: A Boolean value that indicates whether the window is removed from the screen when its application becomes inactive.
//   - [NSWindow.SetHidesOnDeactivate]
//   - [NSWindow.CollectionBehavior]: A value that identifies the window’s behavior in window collections.
//   - [NSWindow.SetCollectionBehavior]
//   - [NSWindow.Opaque]: A Boolean value that indicates whether the window is opaque.
//   - [NSWindow.SetOpaque]
//   - [NSWindow.HasShadow]: A Boolean value that indicates whether the window has a shadow.
//   - [NSWindow.SetHasShadow]
//   - [NSWindow.InvalidateShadow]: Invalidates the window shadow so that it is recomputed based on the current window shape.
//   - [NSWindow.AutorecalculatesContentBorderThicknessForEdge]: Indicates whether the window calculates the thickness of a given border automatically.
//   - [NSWindow.SetAutorecalculatesContentBorderThicknessForEdge]: Specifies whether the window calculates the thickness of a given border automatically.
//   - [NSWindow.ContentBorderThicknessForEdge]: Indicates the thickness of a given border of the window.
//   - [NSWindow.SetContentBorderThicknessForEdge]: Specifies the thickness of a given border of the window.
//   - [NSWindow.PreventsApplicationTerminationWhenModal]: A Boolean value that indicates whether the window prevents application termination when modal.
//   - [NSWindow.SetPreventsApplicationTerminationWhenModal]
//   - [NSWindow.AppearanceSource]: An object that the window inherits its appearance from.
//   - [NSWindow.SetAppearanceSource]
//
// # Accessing Window Information
//
//   - [NSWindow.DepthLimit]: The depth limit of the window.
//   - [NSWindow.SetDepthLimit]
//   - [NSWindow.HasDynamicDepthLimit]: A Boolean value that indicates whether the window’s depth limit can change to match the depth of the screen it’s on.
//   - [NSWindow.WindowNumber]: The window number of the window’s window device.
//   - [NSWindow.DeviceDescription]: A dictionary containing information about the window’s resolution, such as color, depth, and so on.
//   - [NSWindow.CanBecomeVisibleWithoutLogin]: A Boolean value that indicates whether the window can be displayed at the login window.
//   - [NSWindow.SetCanBecomeVisibleWithoutLogin]
//   - [NSWindow.SharingType]: A Boolean value that indicates the level of access other processes have to the window’s content.
//   - [NSWindow.SetSharingType]
//   - [NSWindow.BackingType]: The window’s backing store type.
//   - [NSWindow.SetBackingType]
//   - [NSWindow.DisplayLinkWithTargetSelector]: Returns a new display link whose callback will be invoked in-sync with the display the window is on.
//
// # Getting Layout Information
//
//   - [NSWindow.ContentRectForFrameRect]: Returns the window’s content rectangle with a given frame rectangle.
//   - [NSWindow.FrameRectForContentRect]: Returns the window’s frame rectangle with a given content rectangle.
//
// # Managing Windows
//
//   - [NSWindow.WindowController]: The window’s window controller.
//   - [NSWindow.SetWindowController]
//
// # Managing Sheets
//
//   - [NSWindow.AttachedSheet]: The sheet attached to the window.
//   - [NSWindow.Sheet]: A Boolean value that indicates whether the window has ever run as a modal sheet.
//   - [NSWindow.BeginSheetCompletionHandler]: Starts a document-modal session and presents—or queues for presentation—a sheet.
//   - [NSWindow.BeginCriticalSheetCompletionHandler]: Starts a document-modal session and presents the specified critical sheet.
//   - [NSWindow.EndSheet]: Ends a document-modal session and dismisses the specified sheet.
//   - [NSWindow.EndSheetReturnCode]: Ends a document-modal session and dismisses the specified sheet.
//   - [NSWindow.SheetParent]: The window to which the sheet is attached.
//   - [NSWindow.Sheets]: An array of the sheets currently attached to the window.
//
// # Sizing Windows
//
//   - [NSWindow.Frame]: The window’s frame rectangle in screen coordinates, including the title bar.
//   - [NSWindow.SetFrameOrigin]: Positions the bottom-left corner of the window’s frame rectangle at a given point in screen coordinates.
//   - [NSWindow.SetFrameTopLeftPoint]: Positions the top-left corner of the window’s frame rectangle at a given point in screen coordinates.
//   - [NSWindow.ConstrainFrameRectToScreen]: Modifies and returns a frame rectangle so that its top edge lies on a specific screen.
//   - [NSWindow.CascadeTopLeftFromPoint]: Positions the window’s top-left to a given point.
//   - [NSWindow.SetFrameDisplay]: Sets the origin and size of the window’s frame rectangle according to a given frame rectangle, thereby setting its position and size onscreen.
//   - [NSWindow.SetFrameDisplayAnimate]: Sets the origin and size of the window’s frame rectangle, with optional animation, according to a given frame rectangle, thereby setting its position and size onscreen.
//   - [NSWindow.AnimationResizeTime]: Specifies the duration of a smooth frame-size change.
//   - [NSWindow.AspectRatio]: The window’s aspect ratio, which constrains the size of its frame rectangle to integral multiples of this ratio when the user resizes it.
//   - [NSWindow.SetAspectRatio]
//   - [NSWindow.MinSize]: The minimum size to which the window’s frame (including its title bar) can be sized.
//   - [NSWindow.SetMinSize]
//   - [NSWindow.MaxSize]: The maximum size to which the window’s frame (including its title bar) can be sized.
//   - [NSWindow.SetMaxSize]
//   - [NSWindow.Zoomed]: A Boolean value that indicates whether the window is in a zoomed state.
//   - [NSWindow.PerformZoom]: This action method simulates the user clicking the zoom box by momentarily highlighting the button and then zooming the window.
//   - [NSWindow.Zoom]: Toggles the size and location of the window between its standard state (which the application provides as the best size to display the window’s data) and its user state (a new size and location the user may have set by moving or resizing the window).
//   - [NSWindow.ResizeFlags]: The flags field of the event record for the mouse-down event that initiated the resizing session.
//   - [NSWindow.ResizeIncrements]: The window’s resizing increments.
//   - [NSWindow.SetResizeIncrements]
//   - [NSWindow.PreservesContentDuringLiveResize]: A Boolean value that indicates whether the window tries to optimize user-initiated resize operations by preserving the content of views that have not changed.
//   - [NSWindow.SetPreservesContentDuringLiveResize]
//   - [NSWindow.InLiveResize]: A Boolean value that indicates whether the window is being resized by the user.
//
// # Sizing Content
//
//   - [NSWindow.ContentAspectRatio]: The window’s content aspect ratio.
//   - [NSWindow.SetContentAspectRatio]
//   - [NSWindow.ContentMinSize]: The minimum size of the window’s content view in the window’s base coordinate system.
//   - [NSWindow.SetContentMinSize]
//   - [NSWindow.SetContentSize]: Sets the size of the window’s content view to a given size, which is expressed in the window’s base coordinate system.
//   - [NSWindow.ContentMaxSize]: The maximum size of the window’s content view in the window’s base coordinate system.
//   - [NSWindow.SetContentMaxSize]
//   - [NSWindow.ContentResizeIncrements]: The window’s content-view resizing increments.
//   - [NSWindow.SetContentResizeIncrements]
//   - [NSWindow.ContentLayoutGuide]: A value used by Auto Layout constraints to automatically bind to the value of [contentLayoutRect](<doc://com.apple.appkit/documentation/AppKit/NSWindow/contentLayoutRect>).
//   - [NSWindow.ContentLayoutRect]: The area inside the window that is for non-obscured content, in window coordinates.
//   - [NSWindow.MaxFullScreenContentSize]: A maximum size that is used to determine if a window can fit when it is in full screen in a tile.
//   - [NSWindow.SetMaxFullScreenContentSize]
//   - [NSWindow.MinFullScreenContentSize]: A minimum size that is used to determine if a window can fit when it is in full screen in a tile.
//   - [NSWindow.SetMinFullScreenContentSize]
//
// # Managing Window Layers
//
//   - [NSWindow.OrderOut]: Removes the window from the screen list, which hides the window.
//   - [NSWindow.OrderBack]: Moves the window to the back of its level in the screen list, without changing either the key window or the main window.
//   - [NSWindow.OrderFront]: Moves the window to the front of its level in the screen list, without changing either the key window or the main window.
//   - [NSWindow.OrderFrontRegardless]: Moves the window to the front of its level, even if its application isn’t active, without changing either the key window or the main window.
//   - [NSWindow.OrderWindowRelativeTo]: Repositions the window’s window device in the window server’s screen list.
//   - [NSWindow.Level]: The window level of the window.
//   - [NSWindow.SetLevel]
//
// # Managing Window Visibility and Occlusion State
//
//   - [NSWindow.Visible]: A Boolean value that indicates whether the window is visible onscreen (even when it’s obscured by other windows).
//   - [NSWindow.OcclusionState]: The occlusion state of the window.
//
// # Managing Window Frames in User Defaults
//
//   - [NSWindow.SetFrameUsingName]: Sets the window’s frame rectangle by reading the rectangle data stored under a given name from the defaults system.
//   - [NSWindow.SetFrameUsingNameForce]: Sets the window’s frame rectangle by reading the rectangle data stored under a given name from the defaults system. Can operate on non-resizable windows.
//   - [NSWindow.SaveFrameUsingName]: Saves the window’s frame rectangle in the user defaults system under a given name.
//   - [NSWindow.FrameAutosaveName]: The name used to automatically save the window’s frame rectangle data in the defaults system.
//   - [NSWindow.StringWithSavedFrame]: A string representation of the window’s frame rectangle.
//   - [NSWindow.SetFrameFromString]: Sets the window’s frame rectangle from a given string representation.
//
// # Managing Key Status
//
//   - [NSWindow.KeyWindow]: A Boolean value that indicates whether the window is the key window for the application.
//   - [NSWindow.CanBecomeKeyWindow]: A Boolean value that indicates whether the window can become the key window.
//   - [NSWindow.MakeKeyWindow]: Makes the window the key window.
//   - [NSWindow.MakeKeyAndOrderFront]: Moves the window to the front of the screen list, within its level, and makes it the key window; that is, it shows the window.
//   - [NSWindow.BecomeKeyWindow]: Informs the window that it has become the key window.
//   - [NSWindow.ResignKeyWindow]: Resigns the window’s key window status.
//
// # Managing Main Status
//
//   - [NSWindow.MainWindow]: A Boolean value that indicates whether the window is the application’s main window.
//   - [NSWindow.CanBecomeMainWindow]: A Boolean value that indicates whether the window can become the application’s main window.
//   - [NSWindow.MakeMainWindow]: Makes the window the main window.
//   - [NSWindow.BecomeMainWindow]: Informs the window that it has become the main window.
//   - [NSWindow.ResignMainWindow]: Resigns the window’s main window status.
//
// # Managing Toolbars
//
//   - [NSWindow.Toolbar]: The window’s toolbar.
//   - [NSWindow.SetToolbar]
//   - [NSWindow.ToggleToolbarShown]: Toggles the visibility of the window’s toolbar.
//   - [NSWindow.RunToolbarCustomizationPalette]: Presents the toolbar customization user interface.
//
// # Managing Attached Windows
//
//   - [NSWindow.ChildWindows]: An array of the window’s attached child windows.
//   - [NSWindow.AddChildWindowOrdered]: Adds a given window as a child window of the window.
//   - [NSWindow.RemoveChildWindow]: Detaches a given child window from the window.
//   - [NSWindow.ParentWindow]: The parent window to which the window is attached as a child.
//   - [NSWindow.SetParentWindow]
//
// # Managing Default Buttons
//
//   - [NSWindow.DefaultButtonCell]: The button cell that performs as if clicked when the window receives a Return (or Enter) key event.
//   - [NSWindow.SetDefaultButtonCell]
//   - [NSWindow.EnableKeyEquivalentForDefaultButtonCell]: Reenables the default button cell’s key equivalent, so it performs a click when the user presses Return (or Enter).
//   - [NSWindow.DisableKeyEquivalentForDefaultButtonCell]: Disables the default button cell’s key equivalent, so it doesn’t perform a click when the user presses Return (or Enter).
//
// # Managing Field Editors
//
//   - [NSWindow.FieldEditorForObject]: Returns the window’s field editor, creating it if requested.
//   - [NSWindow.EndEditingFor]: Forces the field editor to give up its first responder status and prepares it for its next assignment.
//
// # Managing the Window Menu
//
//   - [NSWindow.ExcludedFromWindowsMenu]: A Boolean value that indicates whether the window is excluded from the application’s Windows menu.
//   - [NSWindow.SetExcludedFromWindowsMenu]
//
// # Managing Cursor Rectangles
//
//   - [NSWindow.AreCursorRectsEnabled]: A Boolean value that indicates whether the window’s cursor rectangles are enabled.
//   - [NSWindow.EnableCursorRects]: Reenables cursor rectangle management within the window after a [disableCursorRects()](<doc://com.apple.appkit/documentation/AppKit/NSWindow/disableCursorRects()>) message.
//   - [NSWindow.DisableCursorRects]: Disables all cursor rectangle management within the window.
//   - [NSWindow.DiscardCursorRects]: Invalidates all cursor rectangles in the window.
//   - [NSWindow.InvalidateCursorRectsForView]: Marks as invalid the cursor rectangles of a given view object in the window, so they’ll be set up again when the window becomes key.
//   - [NSWindow.ResetCursorRects]: Clears the window’s cursor rectangles and the cursor rectangles of the [NSView](<doc://com.apple.appkit/documentation/AppKit/NSView>) objects in its view hierarchy.
//
// # Managing Title Bars
//
//   - [NSWindow.StandardWindowButton]: Returns the window button of a given window button kind in the window’s view hierarchy.
//   - [NSWindow.ShowsToolbarButton]: A Boolean value that indicates whether the toolbar control button is currently displayed.
//   - [NSWindow.SetShowsToolbarButton]
//   - [NSWindow.TitlebarAppearsTransparent]: A Boolean value that indicates whether the title bar draws its background.
//   - [NSWindow.SetTitlebarAppearsTransparent]
//   - [NSWindow.ToolbarStyle]: The style that determines the appearance and location of the toolbar in relation to the title bar.
//   - [NSWindow.SetToolbarStyle]
//   - [NSWindow.TitlebarSeparatorStyle]: The type of separator that the app displays between the title bar and content of a window.
//   - [NSWindow.SetTitlebarSeparatorStyle]
//   - [NSWindow.WindowTitlebarLayoutDirection]: The direction the window’s title bar lays text out, either left to right or right to left.
//
// # Managing Title Bar Accessories
//
//   - [NSWindow.AddTitlebarAccessoryViewController]: Adds the specified title bar accessory view controller to the window.
//   - [NSWindow.InsertTitlebarAccessoryViewControllerAtIndex]: Inserts the view controller into the window’s array of title bar accessory view controllers at the specified index.
//   - [NSWindow.RemoveTitlebarAccessoryViewControllerAtIndex]: Removes the view controller at the specified index from the window’s array of title bar accessory view controllers.
//   - [NSWindow.TitlebarAccessoryViewControllers]: An array of title bar accessory view controllers that are currently added to the window.
//   - [NSWindow.SetTitlebarAccessoryViewControllers]
//
// # Managing Window Tabs
//
//   - [NSWindow.Tab]: An object that represents information about a window when it displays as a tab.
//   - [NSWindow.TabbingIdentifier]: A value that allows a group of related windows.
//   - [NSWindow.SetTabbingIdentifier]
//   - [NSWindow.AddTabbedWindowOrdered]: Adds the provided window as a new tab in a tabbed window using the specified ordering instruction.
//   - [NSWindow.TabbingMode]: A value that indicates when a window displays tabs.
//   - [NSWindow.SetTabbingMode]
//   - [NSWindow.TabbedWindows]: An array of windows that display as tabs.
//   - [NSWindow.MergeAllWindows]: Merges all open windows into a single tabbed window.
//   - [NSWindow.SelectNextTab]: Selects the next tab in the tab group in the trailing direction.
//   - [NSWindow.SelectPreviousTab]: Selects the previous tab in the tab group in the leading direction.
//   - [NSWindow.MoveTabToNewWindow]: Moves the tab to a new containing window.
//   - [NSWindow.ToggleTabBar]: Shows or hides the tab bar.
//   - [NSWindow.ToggleTabOverview]: Shows or hides the tab overview.
//   - [NSWindow.TabGroup]: A group of windows that display together as a tab group.
//
// # Managing Tooltips
//
//   - [NSWindow.AllowsToolTipsWhenApplicationIsInactive]: A Boolean value that indicates whether the window can display tooltips even when the application is in the background.
//   - [NSWindow.SetAllowsToolTipsWhenApplicationIsInactive]
//
// # Handling Events
//
//   - [NSWindow.CurrentEvent]: The event currently being processed by the application.
//   - [NSWindow.NextEventMatchingMask]: Returns the next event matching a given mask.
//   - [NSWindow.NextEventMatchingMaskUntilDateInModeDequeue]: Forwards the message to the global application object.
//   - [NSWindow.DiscardEventsMatchingMaskBeforeEvent]: Forwards the message to the global application object.
//   - [NSWindow.PostEventAtStart]: Forwards the message to the global application object.
//   - [NSWindow.SendEvent]: This action method dispatches mouse and keyboard events the global application object sends to the window.
//
// # Managing Responders
//
//   - [NSWindow.InitialFirstResponder]: The view that’s made first responder (also called the key view) the first time the window is placed onscreen.
//   - [NSWindow.SetInitialFirstResponder]
//   - [NSWindow.FirstResponder]: The window’s first responder.
//   - [NSWindow.MakeFirstResponder]: Attempts to make a given responder the first responder for the window.
//
// # Managing the Key View Loop
//
//   - [NSWindow.SelectKeyViewPrecedingView]: Gives key view status to the view that precedes the given view.
//   - [NSWindow.SelectKeyViewFollowingView]: Gives key view status to the view that follows the given view.
//   - [NSWindow.SelectPreviousKeyView]: Searches for a candidate previous key view and, if it finds one, tries to make it the first responder.
//   - [NSWindow.SelectNextKeyView]: Searches for a candidate next key view and, if it finds one, tries to make it the first responder.
//   - [NSWindow.KeyViewSelectionDirection]: The direction the window is currently using to change the key view.
//   - [NSWindow.AutorecalculatesKeyViewLoop]: A Boolean value that indicates whether the window automatically recalculates the key view loop when views are added.
//   - [NSWindow.SetAutorecalculatesKeyViewLoop]
//   - [NSWindow.RecalculateKeyViewLoop]: Marks the key view loop as “dirty” and in need of recalculation.
//
// # Managing Window Sharing
//
//   - [NSWindow.TransferWindowSharingToWindowCompletionHandler]
//   - [NSWindow.HasActiveWindowSharingSession]: Indicates whether the receiver is the subject of an active SharePlay sharing session.
//
// # Handling Mouse Events
//
//   - [NSWindow.AcceptsMouseMovedEvents]: A Boolean value that indicates whether the window accepts mouse-moved events.
//   - [NSWindow.SetAcceptsMouseMovedEvents]
//   - [NSWindow.IgnoresMouseEvents]: A Boolean value that indicates whether the window is transparent to mouse events.
//   - [NSWindow.SetIgnoresMouseEvents]
//   - [NSWindow.MouseLocationOutsideOfEventStream]: The current location of the pointer reckoned in the window’s base coordinate system, regardless of the current event being handled or of any events pending.
//   - [NSWindow.TrackEventsMatchingMaskTimeoutModeHandler]: Tracks events that match the specified mask using the specified tracking handler until the tracking handler explicitly terminates tracking.
//   - [NSWindow.PerformWindowDragWithEvent]: Starts a window drag based on the specified mouse-down event.
//
// # Handling Window Restoration
//
//   - [NSWindow.Restorable]: A Boolean value indicating whether the window configuration is preserved between application launches.
//   - [NSWindow.SetRestorable]
//   - [NSWindow.RestorationClass]: The restoration class associated with the window.
//   - [NSWindow.SetRestorationClass]
//   - [NSWindow.DisableSnapshotRestoration]: Disables snapshot restoration.
//   - [NSWindow.EnableSnapshotRestoration]: Enables snapshot restoration.
//
// # Drawing Windows
//
//   - [NSWindow.Display]: Passes a display message down the window’s view hierarchy, thus redrawing all views within the window.
//   - [NSWindow.DisplayIfNeeded]: Passes a display message down the window’s view hierarchy, thus redrawing all views that need displaying.
//   - [NSWindow.ViewsNeedDisplay]: A Boolean value that indicates whether any of the window’s views need to be displayed.
//   - [NSWindow.SetViewsNeedDisplay]
//   - [NSWindow.AllowsConcurrentViewDrawing]: A Boolean value that indicates whether the window allows multithreaded view drawing.
//   - [NSWindow.SetAllowsConcurrentViewDrawing]
//
// # Window Animation
//
//   - [NSWindow.AnimationBehavior]: The window’s automatic animation behavior.
//   - [NSWindow.SetAnimationBehavior]
//
// # Updating Windows
//
//   - [NSWindow.Update]: Updates the window.
//
// # Dragging Items
//
//   - [NSWindow.RegisterForDraggedTypes]: Registers a set of pasteboard types that the window accepts as the destination of an image-dragging session.
//   - [NSWindow.UnregisterDraggedTypes]: Unregisters the window as a possible destination for dragging operations.
//
// # Accessing Edited Status
//
//   - [NSWindow.DocumentEdited]: A Boolean value that indicates whether the window’s document has been edited.
//   - [NSWindow.SetDocumentEdited]
//
// # Converting Coordinates
//
//   - [NSWindow.BackingScaleFactor]: The backing scale factor.
//   - [NSWindow.BackingAlignedRectOptions]: Returns a backing store pixel-aligned rectangle in window coordinates.
//   - [NSWindow.ConvertRectFromBacking]: Converts a rectangle from its pixel-aligned backing store coordinate system to the window’s coordinate system.
//   - [NSWindow.ConvertRectFromScreen]: Converts a rectangle from the screen coordinate system to the window’s coordinate system.
//   - [NSWindow.ConvertPointFromBacking]: Converts a point from its pixel-aligned backing store coordinate system to the window’s coordinate system.
//   - [NSWindow.ConvertPointFromScreen]: Converts a point from the screen coordinate system to the window’s coordinate system.
//   - [NSWindow.ConvertRectToBacking]: Converts a rectangle from the window’s coordinate system to its pixel-aligned backing store coordinate system.
//   - [NSWindow.ConvertRectToScreen]: Converts a rectangle to the screen coordinate system from the window’s coordinate system.
//   - [NSWindow.ConvertPointToBacking]: Converts a point from the window’s coordinate system to its pixel-aligned backing store coordinate system.
//   - [NSWindow.ConvertPointToScreen]: Converts a point to the screen coordinate system from the window’s coordinate system.
//
// # Managing Titles
//
//   - [NSWindow.Title]: The string that appears in the title bar of the window or the path to the represented file.
//   - [NSWindow.SetTitle]
//   - [NSWindow.Subtitle]: A secondary line of text that appears in the title bar of the window.
//   - [NSWindow.SetSubtitle]
//   - [NSWindow.TitleVisibility]: A value that indicates the visibility of the window’s title and title bar buttons.
//   - [NSWindow.SetTitleVisibility]
//   - [NSWindow.SetTitleWithRepresentedFilename]: Sets a given path as the window’s title, formatting it as a file-system path, and records this path as the window’s associated file.
//   - [NSWindow.RepresentedFilename]: The path to the file of the window’s represented file.
//   - [NSWindow.SetRepresentedFilename]
//   - [NSWindow.RepresentedURL]: The URL of the file the window represents.
//   - [NSWindow.SetRepresentedURL]
//
// # Accessing Screen Information
//
//   - [NSWindow.Screen]: The screen the window is on.
//   - [NSWindow.DeepestScreen]: The deepest screen the window is on (it may be split over several screens).
//   - [NSWindow.DisplaysWhenScreenProfileChanges]: A Boolean value that indicates whether the window context should be updated when the screen profile changes or when the window moves to a different screen.
//   - [NSWindow.SetDisplaysWhenScreenProfileChanges]
//
// # Moving Windows
//
//   - [NSWindow.MovableByWindowBackground]: A Boolean value that indicates whether the window is movable by clicking and dragging anywhere in its background.
//   - [NSWindow.SetMovableByWindowBackground]
//   - [NSWindow.Movable]: A Boolean value that indicates whether the window can be dragged by clicking in its title bar or background.
//   - [NSWindow.SetMovable]
//   - [NSWindow.Center]: Sets the window’s location to the center of the screen.
//
// # Closing Windows
//
//   - [NSWindow.PerformClose]: Simulates the user clicking the close button by momentarily highlighting the button and then closing the window.
//   - [NSWindow.Close]: Removes the window from the screen.
//   - [NSWindow.ReleasedWhenClosed]: A Boolean value that indicates whether the window is released when it receives the `close` message.
//   - [NSWindow.SetReleasedWhenClosed]
//
// # Minimizing Windows
//
//   - [NSWindow.Miniaturized]: A Boolean value that indicates whether the window is minimized.
//   - [NSWindow.PerformMiniaturize]: Simulates the user clicking the minimize button by momentarily highlighting the button, then minimizing the window.
//   - [NSWindow.Miniaturize]: Removes the window from the screen list and displays the minimized window in the Dock.
//   - [NSWindow.Deminiaturize]: De-minimizes the window.
//   - [NSWindow.MiniwindowImage]: The custom miniaturized window image of the window.
//   - [NSWindow.SetMiniwindowImage]
//   - [NSWindow.MiniwindowTitle]: The title displayed in the window’s minimized window.
//   - [NSWindow.SetMiniwindowTitle]
//
// # Getting the Dock Tile
//
//   - [NSWindow.DockTile]: The application’s Dock tile.
//
// # Printing Windows
//
//   - [NSWindow.Print]: Runs the Print panel, and if the user chooses an option other than canceling, prints the window (its frame view and all subviews).
//   - [NSWindow.DataWithEPSInsideRect]: Returns EPS data that draws the region of the window within a given rectangle.
//   - [NSWindow.DataWithPDFInsideRect]: Returns PDF data that draws the region of the window within a given rectangle.
//
// # Triggering Constraint-Based Layout
//
//   - [NSWindow.UpdateConstraintsIfNeeded]: Updates the constraints based on changes to views in the window since the last layout.
//   - [NSWindow.LayoutIfNeeded]: Updates the layout of views in the window based on the current views and constraints.
//
// # Debugging Constraint-Based Layout
//
//   - [NSWindow.VisualizeConstraints]: Displays a visual representation of the supplied constraints in the window.
//
// # Constraint-Based Layouts
//
//   - [NSWindow.AnchorAttributeForOrientation]: Returns the part of the window that stays stationary during constraint-based layout.
//   - [NSWindow.SetAnchorAttributeForOrientation]: Sets the part of the window that stays stationary during constraint-based layout.
//
// # Working with Window Depths
//
//   - [NSWindow.BitsPerPixel]: Returns the bits per pixel for the specified window depth.
//   - [NSWindow.SetNSBitsPerPixelFromDepth]
//   - [NSWindow.BitsPerSample]: Returns the bits per sample for the specified window depth.
//   - [NSWindow.SetNSBitsPerSampleFromDepth]
//   - [NSWindow.ColorSpaceName]: Returns the name of the color space corresponding to the passed window depth.
//   - [NSWindow.SetNSColorSpaceFromDepth]
//   - [NSWindow.NumberOfColorComponents]: Returns the number of color components in the specified color space.
//   - [NSWindow.SetNSNumberOfColorComponents]
//   - [NSWindow.IsPlanar]: Returns whether the specified window depth is planar.
//   - [NSWindow.SetNSPlanarFromDepth]
//   - [NSWindow.CanRepresentDisplayGamut]: A Boolean value that indicates if the window and its screen use a color space that can represent the specified display gamut.
//
// # Getting Information About Scripting Attributes
//
//   - [NSWindow.HasCloseBox]: A Boolean value that indicates if the window has a close box.
//   - [NSWindow.HasTitleBar]: A Boolean value that indicates if the window has a title bar.
//   - [NSWindow.ModalPanel]: A Boolean value that indicates whether the window is a modal panel.
//   - [NSWindow.FloatingPanel]: A Boolean value that indicates whether the window is a floating panel.
//   - [NSWindow.Zoomable]: A Boolean value that indicates whether the window allows zooming.
//   - [NSWindow.Resizable]: A Boolean value that indicates if the user can resize the window.
//   - [NSWindow.Miniaturizable]: A Boolean value that indicates whether the window can minimize.
//   - [NSWindow.OrderedIndex]: The zero-based position of the window, based on its order from front to back among all visible application windows.
//   - [NSWindow.SetOrderedIndex]
//
// # Setting Scripting Attributes
//
//   - [NSWindow.SetIsMiniaturized]: Sets the window’s miniaturized state to the value you specify.
//   - [NSWindow.SetIsVisible]: Sets the window’s visible state to the value you specify.
//   - [NSWindow.SetIsZoomed]: Sets the window’s zoomed state to the value you specify.
//
// # Handling Script Commands
//
//   - [NSWindow.HandleCloseScriptCommand]: Handles the AppleScript command to close the window (and its associated document, if any).
//   - [NSWindow.HandlePrintScriptCommand]: Handles the AppleScript command to print the contents of the window (or its associated document, if any).
//   - [NSWindow.HandleSaveScriptCommand]: Handles the AppleScript command to save the window (and its associated document, if any).
//
// # Instance Properties
//
//   - [NSWindow.CascadingReferenceFrame]
//
// # Instance Methods
//
//   - [NSWindow.BeginDraggingSessionWithItemsEventSource]
//   - [NSWindow.RequestSharingOfWindowCompletionHandler]
//   - [NSWindow.RequestSharingOfWindowUsingPreviewTitleCompletionHandler]
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow
type NSWindow struct {
	NSResponder
}

// NSWindowFromID constructs a [NSWindow] from an objc.ID.
//
// A window that an app displays on the screen.
func NSWindowFromID(id objc.ID) NSWindow {
	return NSWindow{NSResponder: NSResponderFromID(id)}
}

// NOTE: NSWindow adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSWindow] class.
//
// # Creating a Window
//
//   - [INSWindow.InitWithContentRectStyleMaskBackingDefer]: Initializes the window with the specified values.
//   - [INSWindow.InitWithContentRectStyleMaskBackingDeferScreen]: Initializes an allocated window with the specified values.
//
// # Managing the Window’s Behavior
//
//   - [INSWindow.Delegate]: The window’s delegate.
//   - [INSWindow.SetDelegate]
//
// # Configuring the Window’s Content
//
//   - [INSWindow.ContentViewController]: The main content view controller for the window.
//   - [INSWindow.SetContentViewController]
//   - [INSWindow.ContentView]: The window’s content view, the highest accessible view object in the window’s view hierarchy.
//   - [INSWindow.SetContentView]
//
// # Configuring the Window’s Appearance
//
//   - [INSWindow.StyleMask]: Flags that describe the window’s current style, such as if it’s resizable or in full-screen mode.
//   - [INSWindow.SetStyleMask]
//   - [INSWindow.ToggleFullScreen]: Takes the window into or out of fullscreen mode,
//   - [INSWindow.WorksWhenModal]: A Boolean value that indicates whether the window is able to receive keyboard and mouse events even when some other window is being run modally.
//   - [INSWindow.AlphaValue]: The window’s alpha value.
//   - [INSWindow.SetAlphaValue]
//   - [INSWindow.BackgroundColor]: The color of the window’s background.
//   - [INSWindow.SetBackgroundColor]
//   - [INSWindow.ColorSpace]: The window’s color space.
//   - [INSWindow.SetColorSpace]
//   - [INSWindow.SetDynamicDepthLimit]: Sets a Boolean value that indicates whether the window’s depth limit can change to match the depth of the screen it’s on.
//   - [INSWindow.CanHide]: A Boolean value that indicates whether the window can hide when its application becomes hidden.
//   - [INSWindow.SetCanHide]
//   - [INSWindow.OnActiveSpace]: A Boolean value that indicates whether the window is on the currently active space.
//   - [INSWindow.HidesOnDeactivate]: A Boolean value that indicates whether the window is removed from the screen when its application becomes inactive.
//   - [INSWindow.SetHidesOnDeactivate]
//   - [INSWindow.CollectionBehavior]: A value that identifies the window’s behavior in window collections.
//   - [INSWindow.SetCollectionBehavior]
//   - [INSWindow.Opaque]: A Boolean value that indicates whether the window is opaque.
//   - [INSWindow.SetOpaque]
//   - [INSWindow.HasShadow]: A Boolean value that indicates whether the window has a shadow.
//   - [INSWindow.SetHasShadow]
//   - [INSWindow.InvalidateShadow]: Invalidates the window shadow so that it is recomputed based on the current window shape.
//   - [INSWindow.AutorecalculatesContentBorderThicknessForEdge]: Indicates whether the window calculates the thickness of a given border automatically.
//   - [INSWindow.SetAutorecalculatesContentBorderThicknessForEdge]: Specifies whether the window calculates the thickness of a given border automatically.
//   - [INSWindow.ContentBorderThicknessForEdge]: Indicates the thickness of a given border of the window.
//   - [INSWindow.SetContentBorderThicknessForEdge]: Specifies the thickness of a given border of the window.
//   - [INSWindow.PreventsApplicationTerminationWhenModal]: A Boolean value that indicates whether the window prevents application termination when modal.
//   - [INSWindow.SetPreventsApplicationTerminationWhenModal]
//   - [INSWindow.AppearanceSource]: An object that the window inherits its appearance from.
//   - [INSWindow.SetAppearanceSource]
//
// # Accessing Window Information
//
//   - [INSWindow.DepthLimit]: The depth limit of the window.
//   - [INSWindow.SetDepthLimit]
//   - [INSWindow.HasDynamicDepthLimit]: A Boolean value that indicates whether the window’s depth limit can change to match the depth of the screen it’s on.
//   - [INSWindow.WindowNumber]: The window number of the window’s window device.
//   - [INSWindow.DeviceDescription]: A dictionary containing information about the window’s resolution, such as color, depth, and so on.
//   - [INSWindow.CanBecomeVisibleWithoutLogin]: A Boolean value that indicates whether the window can be displayed at the login window.
//   - [INSWindow.SetCanBecomeVisibleWithoutLogin]
//   - [INSWindow.SharingType]: A Boolean value that indicates the level of access other processes have to the window’s content.
//   - [INSWindow.SetSharingType]
//   - [INSWindow.BackingType]: The window’s backing store type.
//   - [INSWindow.SetBackingType]
//   - [INSWindow.DisplayLinkWithTargetSelector]: Returns a new display link whose callback will be invoked in-sync with the display the window is on.
//
// # Getting Layout Information
//
//   - [INSWindow.ContentRectForFrameRect]: Returns the window’s content rectangle with a given frame rectangle.
//   - [INSWindow.FrameRectForContentRect]: Returns the window’s frame rectangle with a given content rectangle.
//
// # Managing Windows
//
//   - [INSWindow.WindowController]: The window’s window controller.
//   - [INSWindow.SetWindowController]
//
// # Managing Sheets
//
//   - [INSWindow.AttachedSheet]: The sheet attached to the window.
//   - [INSWindow.Sheet]: A Boolean value that indicates whether the window has ever run as a modal sheet.
//   - [INSWindow.BeginSheetCompletionHandler]: Starts a document-modal session and presents—or queues for presentation—a sheet.
//   - [INSWindow.BeginCriticalSheetCompletionHandler]: Starts a document-modal session and presents the specified critical sheet.
//   - [INSWindow.EndSheet]: Ends a document-modal session and dismisses the specified sheet.
//   - [INSWindow.EndSheetReturnCode]: Ends a document-modal session and dismisses the specified sheet.
//   - [INSWindow.SheetParent]: The window to which the sheet is attached.
//   - [INSWindow.Sheets]: An array of the sheets currently attached to the window.
//
// # Sizing Windows
//
//   - [INSWindow.Frame]: The window’s frame rectangle in screen coordinates, including the title bar.
//   - [INSWindow.SetFrameOrigin]: Positions the bottom-left corner of the window’s frame rectangle at a given point in screen coordinates.
//   - [INSWindow.SetFrameTopLeftPoint]: Positions the top-left corner of the window’s frame rectangle at a given point in screen coordinates.
//   - [INSWindow.ConstrainFrameRectToScreen]: Modifies and returns a frame rectangle so that its top edge lies on a specific screen.
//   - [INSWindow.CascadeTopLeftFromPoint]: Positions the window’s top-left to a given point.
//   - [INSWindow.SetFrameDisplay]: Sets the origin and size of the window’s frame rectangle according to a given frame rectangle, thereby setting its position and size onscreen.
//   - [INSWindow.SetFrameDisplayAnimate]: Sets the origin and size of the window’s frame rectangle, with optional animation, according to a given frame rectangle, thereby setting its position and size onscreen.
//   - [INSWindow.AnimationResizeTime]: Specifies the duration of a smooth frame-size change.
//   - [INSWindow.AspectRatio]: The window’s aspect ratio, which constrains the size of its frame rectangle to integral multiples of this ratio when the user resizes it.
//   - [INSWindow.SetAspectRatio]
//   - [INSWindow.MinSize]: The minimum size to which the window’s frame (including its title bar) can be sized.
//   - [INSWindow.SetMinSize]
//   - [INSWindow.MaxSize]: The maximum size to which the window’s frame (including its title bar) can be sized.
//   - [INSWindow.SetMaxSize]
//   - [INSWindow.Zoomed]: A Boolean value that indicates whether the window is in a zoomed state.
//   - [INSWindow.PerformZoom]: This action method simulates the user clicking the zoom box by momentarily highlighting the button and then zooming the window.
//   - [INSWindow.Zoom]: Toggles the size and location of the window between its standard state (which the application provides as the best size to display the window’s data) and its user state (a new size and location the user may have set by moving or resizing the window).
//   - [INSWindow.ResizeFlags]: The flags field of the event record for the mouse-down event that initiated the resizing session.
//   - [INSWindow.ResizeIncrements]: The window’s resizing increments.
//   - [INSWindow.SetResizeIncrements]
//   - [INSWindow.PreservesContentDuringLiveResize]: A Boolean value that indicates whether the window tries to optimize user-initiated resize operations by preserving the content of views that have not changed.
//   - [INSWindow.SetPreservesContentDuringLiveResize]
//   - [INSWindow.InLiveResize]: A Boolean value that indicates whether the window is being resized by the user.
//
// # Sizing Content
//
//   - [INSWindow.ContentAspectRatio]: The window’s content aspect ratio.
//   - [INSWindow.SetContentAspectRatio]
//   - [INSWindow.ContentMinSize]: The minimum size of the window’s content view in the window’s base coordinate system.
//   - [INSWindow.SetContentMinSize]
//   - [INSWindow.SetContentSize]: Sets the size of the window’s content view to a given size, which is expressed in the window’s base coordinate system.
//   - [INSWindow.ContentMaxSize]: The maximum size of the window’s content view in the window’s base coordinate system.
//   - [INSWindow.SetContentMaxSize]
//   - [INSWindow.ContentResizeIncrements]: The window’s content-view resizing increments.
//   - [INSWindow.SetContentResizeIncrements]
//   - [INSWindow.ContentLayoutGuide]: A value used by Auto Layout constraints to automatically bind to the value of [contentLayoutRect](<doc://com.apple.appkit/documentation/AppKit/NSWindow/contentLayoutRect>).
//   - [INSWindow.ContentLayoutRect]: The area inside the window that is for non-obscured content, in window coordinates.
//   - [INSWindow.MaxFullScreenContentSize]: A maximum size that is used to determine if a window can fit when it is in full screen in a tile.
//   - [INSWindow.SetMaxFullScreenContentSize]
//   - [INSWindow.MinFullScreenContentSize]: A minimum size that is used to determine if a window can fit when it is in full screen in a tile.
//   - [INSWindow.SetMinFullScreenContentSize]
//
// # Managing Window Layers
//
//   - [INSWindow.OrderOut]: Removes the window from the screen list, which hides the window.
//   - [INSWindow.OrderBack]: Moves the window to the back of its level in the screen list, without changing either the key window or the main window.
//   - [INSWindow.OrderFront]: Moves the window to the front of its level in the screen list, without changing either the key window or the main window.
//   - [INSWindow.OrderFrontRegardless]: Moves the window to the front of its level, even if its application isn’t active, without changing either the key window or the main window.
//   - [INSWindow.OrderWindowRelativeTo]: Repositions the window’s window device in the window server’s screen list.
//   - [INSWindow.Level]: The window level of the window.
//   - [INSWindow.SetLevel]
//
// # Managing Window Visibility and Occlusion State
//
//   - [INSWindow.Visible]: A Boolean value that indicates whether the window is visible onscreen (even when it’s obscured by other windows).
//   - [INSWindow.OcclusionState]: The occlusion state of the window.
//
// # Managing Window Frames in User Defaults
//
//   - [INSWindow.SetFrameUsingName]: Sets the window’s frame rectangle by reading the rectangle data stored under a given name from the defaults system.
//   - [INSWindow.SetFrameUsingNameForce]: Sets the window’s frame rectangle by reading the rectangle data stored under a given name from the defaults system. Can operate on non-resizable windows.
//   - [INSWindow.SaveFrameUsingName]: Saves the window’s frame rectangle in the user defaults system under a given name.
//   - [INSWindow.FrameAutosaveName]: The name used to automatically save the window’s frame rectangle data in the defaults system.
//   - [INSWindow.StringWithSavedFrame]: A string representation of the window’s frame rectangle.
//   - [INSWindow.SetFrameFromString]: Sets the window’s frame rectangle from a given string representation.
//
// # Managing Key Status
//
//   - [INSWindow.KeyWindow]: A Boolean value that indicates whether the window is the key window for the application.
//   - [INSWindow.CanBecomeKeyWindow]: A Boolean value that indicates whether the window can become the key window.
//   - [INSWindow.MakeKeyWindow]: Makes the window the key window.
//   - [INSWindow.MakeKeyAndOrderFront]: Moves the window to the front of the screen list, within its level, and makes it the key window; that is, it shows the window.
//   - [INSWindow.BecomeKeyWindow]: Informs the window that it has become the key window.
//   - [INSWindow.ResignKeyWindow]: Resigns the window’s key window status.
//
// # Managing Main Status
//
//   - [INSWindow.MainWindow]: A Boolean value that indicates whether the window is the application’s main window.
//   - [INSWindow.CanBecomeMainWindow]: A Boolean value that indicates whether the window can become the application’s main window.
//   - [INSWindow.MakeMainWindow]: Makes the window the main window.
//   - [INSWindow.BecomeMainWindow]: Informs the window that it has become the main window.
//   - [INSWindow.ResignMainWindow]: Resigns the window’s main window status.
//
// # Managing Toolbars
//
//   - [INSWindow.Toolbar]: The window’s toolbar.
//   - [INSWindow.SetToolbar]
//   - [INSWindow.ToggleToolbarShown]: Toggles the visibility of the window’s toolbar.
//   - [INSWindow.RunToolbarCustomizationPalette]: Presents the toolbar customization user interface.
//
// # Managing Attached Windows
//
//   - [INSWindow.ChildWindows]: An array of the window’s attached child windows.
//   - [INSWindow.AddChildWindowOrdered]: Adds a given window as a child window of the window.
//   - [INSWindow.RemoveChildWindow]: Detaches a given child window from the window.
//   - [INSWindow.ParentWindow]: The parent window to which the window is attached as a child.
//   - [INSWindow.SetParentWindow]
//
// # Managing Default Buttons
//
//   - [INSWindow.DefaultButtonCell]: The button cell that performs as if clicked when the window receives a Return (or Enter) key event.
//   - [INSWindow.SetDefaultButtonCell]
//   - [INSWindow.EnableKeyEquivalentForDefaultButtonCell]: Reenables the default button cell’s key equivalent, so it performs a click when the user presses Return (or Enter).
//   - [INSWindow.DisableKeyEquivalentForDefaultButtonCell]: Disables the default button cell’s key equivalent, so it doesn’t perform a click when the user presses Return (or Enter).
//
// # Managing Field Editors
//
//   - [INSWindow.FieldEditorForObject]: Returns the window’s field editor, creating it if requested.
//   - [INSWindow.EndEditingFor]: Forces the field editor to give up its first responder status and prepares it for its next assignment.
//
// # Managing the Window Menu
//
//   - [INSWindow.ExcludedFromWindowsMenu]: A Boolean value that indicates whether the window is excluded from the application’s Windows menu.
//   - [INSWindow.SetExcludedFromWindowsMenu]
//
// # Managing Cursor Rectangles
//
//   - [INSWindow.AreCursorRectsEnabled]: A Boolean value that indicates whether the window’s cursor rectangles are enabled.
//   - [INSWindow.EnableCursorRects]: Reenables cursor rectangle management within the window after a [disableCursorRects()](<doc://com.apple.appkit/documentation/AppKit/NSWindow/disableCursorRects()>) message.
//   - [INSWindow.DisableCursorRects]: Disables all cursor rectangle management within the window.
//   - [INSWindow.DiscardCursorRects]: Invalidates all cursor rectangles in the window.
//   - [INSWindow.InvalidateCursorRectsForView]: Marks as invalid the cursor rectangles of a given view object in the window, so they’ll be set up again when the window becomes key.
//   - [INSWindow.ResetCursorRects]: Clears the window’s cursor rectangles and the cursor rectangles of the [NSView](<doc://com.apple.appkit/documentation/AppKit/NSView>) objects in its view hierarchy.
//
// # Managing Title Bars
//
//   - [INSWindow.StandardWindowButton]: Returns the window button of a given window button kind in the window’s view hierarchy.
//   - [INSWindow.ShowsToolbarButton]: A Boolean value that indicates whether the toolbar control button is currently displayed.
//   - [INSWindow.SetShowsToolbarButton]
//   - [INSWindow.TitlebarAppearsTransparent]: A Boolean value that indicates whether the title bar draws its background.
//   - [INSWindow.SetTitlebarAppearsTransparent]
//   - [INSWindow.ToolbarStyle]: The style that determines the appearance and location of the toolbar in relation to the title bar.
//   - [INSWindow.SetToolbarStyle]
//   - [INSWindow.TitlebarSeparatorStyle]: The type of separator that the app displays between the title bar and content of a window.
//   - [INSWindow.SetTitlebarSeparatorStyle]
//   - [INSWindow.WindowTitlebarLayoutDirection]: The direction the window’s title bar lays text out, either left to right or right to left.
//
// # Managing Title Bar Accessories
//
//   - [INSWindow.AddTitlebarAccessoryViewController]: Adds the specified title bar accessory view controller to the window.
//   - [INSWindow.InsertTitlebarAccessoryViewControllerAtIndex]: Inserts the view controller into the window’s array of title bar accessory view controllers at the specified index.
//   - [INSWindow.RemoveTitlebarAccessoryViewControllerAtIndex]: Removes the view controller at the specified index from the window’s array of title bar accessory view controllers.
//   - [INSWindow.TitlebarAccessoryViewControllers]: An array of title bar accessory view controllers that are currently added to the window.
//   - [INSWindow.SetTitlebarAccessoryViewControllers]
//
// # Managing Window Tabs
//
//   - [INSWindow.Tab]: An object that represents information about a window when it displays as a tab.
//   - [INSWindow.TabbingIdentifier]: A value that allows a group of related windows.
//   - [INSWindow.SetTabbingIdentifier]
//   - [INSWindow.AddTabbedWindowOrdered]: Adds the provided window as a new tab in a tabbed window using the specified ordering instruction.
//   - [INSWindow.TabbingMode]: A value that indicates when a window displays tabs.
//   - [INSWindow.SetTabbingMode]
//   - [INSWindow.TabbedWindows]: An array of windows that display as tabs.
//   - [INSWindow.MergeAllWindows]: Merges all open windows into a single tabbed window.
//   - [INSWindow.SelectNextTab]: Selects the next tab in the tab group in the trailing direction.
//   - [INSWindow.SelectPreviousTab]: Selects the previous tab in the tab group in the leading direction.
//   - [INSWindow.MoveTabToNewWindow]: Moves the tab to a new containing window.
//   - [INSWindow.ToggleTabBar]: Shows or hides the tab bar.
//   - [INSWindow.ToggleTabOverview]: Shows or hides the tab overview.
//   - [INSWindow.TabGroup]: A group of windows that display together as a tab group.
//
// # Managing Tooltips
//
//   - [INSWindow.AllowsToolTipsWhenApplicationIsInactive]: A Boolean value that indicates whether the window can display tooltips even when the application is in the background.
//   - [INSWindow.SetAllowsToolTipsWhenApplicationIsInactive]
//
// # Handling Events
//
//   - [INSWindow.CurrentEvent]: The event currently being processed by the application.
//   - [INSWindow.NextEventMatchingMask]: Returns the next event matching a given mask.
//   - [INSWindow.NextEventMatchingMaskUntilDateInModeDequeue]: Forwards the message to the global application object.
//   - [INSWindow.DiscardEventsMatchingMaskBeforeEvent]: Forwards the message to the global application object.
//   - [INSWindow.PostEventAtStart]: Forwards the message to the global application object.
//   - [INSWindow.SendEvent]: This action method dispatches mouse and keyboard events the global application object sends to the window.
//
// # Managing Responders
//
//   - [INSWindow.InitialFirstResponder]: The view that’s made first responder (also called the key view) the first time the window is placed onscreen.
//   - [INSWindow.SetInitialFirstResponder]
//   - [INSWindow.FirstResponder]: The window’s first responder.
//   - [INSWindow.MakeFirstResponder]: Attempts to make a given responder the first responder for the window.
//
// # Managing the Key View Loop
//
//   - [INSWindow.SelectKeyViewPrecedingView]: Gives key view status to the view that precedes the given view.
//   - [INSWindow.SelectKeyViewFollowingView]: Gives key view status to the view that follows the given view.
//   - [INSWindow.SelectPreviousKeyView]: Searches for a candidate previous key view and, if it finds one, tries to make it the first responder.
//   - [INSWindow.SelectNextKeyView]: Searches for a candidate next key view and, if it finds one, tries to make it the first responder.
//   - [INSWindow.KeyViewSelectionDirection]: The direction the window is currently using to change the key view.
//   - [INSWindow.AutorecalculatesKeyViewLoop]: A Boolean value that indicates whether the window automatically recalculates the key view loop when views are added.
//   - [INSWindow.SetAutorecalculatesKeyViewLoop]
//   - [INSWindow.RecalculateKeyViewLoop]: Marks the key view loop as “dirty” and in need of recalculation.
//
// # Managing Window Sharing
//
//   - [INSWindow.TransferWindowSharingToWindowCompletionHandler]
//   - [INSWindow.HasActiveWindowSharingSession]: Indicates whether the receiver is the subject of an active SharePlay sharing session.
//
// # Handling Mouse Events
//
//   - [INSWindow.AcceptsMouseMovedEvents]: A Boolean value that indicates whether the window accepts mouse-moved events.
//   - [INSWindow.SetAcceptsMouseMovedEvents]
//   - [INSWindow.IgnoresMouseEvents]: A Boolean value that indicates whether the window is transparent to mouse events.
//   - [INSWindow.SetIgnoresMouseEvents]
//   - [INSWindow.MouseLocationOutsideOfEventStream]: The current location of the pointer reckoned in the window’s base coordinate system, regardless of the current event being handled or of any events pending.
//   - [INSWindow.TrackEventsMatchingMaskTimeoutModeHandler]: Tracks events that match the specified mask using the specified tracking handler until the tracking handler explicitly terminates tracking.
//   - [INSWindow.PerformWindowDragWithEvent]: Starts a window drag based on the specified mouse-down event.
//
// # Handling Window Restoration
//
//   - [INSWindow.Restorable]: A Boolean value indicating whether the window configuration is preserved between application launches.
//   - [INSWindow.SetRestorable]
//   - [INSWindow.RestorationClass]: The restoration class associated with the window.
//   - [INSWindow.SetRestorationClass]
//   - [INSWindow.DisableSnapshotRestoration]: Disables snapshot restoration.
//   - [INSWindow.EnableSnapshotRestoration]: Enables snapshot restoration.
//
// # Drawing Windows
//
//   - [INSWindow.Display]: Passes a display message down the window’s view hierarchy, thus redrawing all views within the window.
//   - [INSWindow.DisplayIfNeeded]: Passes a display message down the window’s view hierarchy, thus redrawing all views that need displaying.
//   - [INSWindow.ViewsNeedDisplay]: A Boolean value that indicates whether any of the window’s views need to be displayed.
//   - [INSWindow.SetViewsNeedDisplay]
//   - [INSWindow.AllowsConcurrentViewDrawing]: A Boolean value that indicates whether the window allows multithreaded view drawing.
//   - [INSWindow.SetAllowsConcurrentViewDrawing]
//
// # Window Animation
//
//   - [INSWindow.AnimationBehavior]: The window’s automatic animation behavior.
//   - [INSWindow.SetAnimationBehavior]
//
// # Updating Windows
//
//   - [INSWindow.Update]: Updates the window.
//
// # Dragging Items
//
//   - [INSWindow.RegisterForDraggedTypes]: Registers a set of pasteboard types that the window accepts as the destination of an image-dragging session.
//   - [INSWindow.UnregisterDraggedTypes]: Unregisters the window as a possible destination for dragging operations.
//
// # Accessing Edited Status
//
//   - [INSWindow.DocumentEdited]: A Boolean value that indicates whether the window’s document has been edited.
//   - [INSWindow.SetDocumentEdited]
//
// # Converting Coordinates
//
//   - [INSWindow.BackingScaleFactor]: The backing scale factor.
//   - [INSWindow.BackingAlignedRectOptions]: Returns a backing store pixel-aligned rectangle in window coordinates.
//   - [INSWindow.ConvertRectFromBacking]: Converts a rectangle from its pixel-aligned backing store coordinate system to the window’s coordinate system.
//   - [INSWindow.ConvertRectFromScreen]: Converts a rectangle from the screen coordinate system to the window’s coordinate system.
//   - [INSWindow.ConvertPointFromBacking]: Converts a point from its pixel-aligned backing store coordinate system to the window’s coordinate system.
//   - [INSWindow.ConvertPointFromScreen]: Converts a point from the screen coordinate system to the window’s coordinate system.
//   - [INSWindow.ConvertRectToBacking]: Converts a rectangle from the window’s coordinate system to its pixel-aligned backing store coordinate system.
//   - [INSWindow.ConvertRectToScreen]: Converts a rectangle to the screen coordinate system from the window’s coordinate system.
//   - [INSWindow.ConvertPointToBacking]: Converts a point from the window’s coordinate system to its pixel-aligned backing store coordinate system.
//   - [INSWindow.ConvertPointToScreen]: Converts a point to the screen coordinate system from the window’s coordinate system.
//
// # Managing Titles
//
//   - [INSWindow.Title]: The string that appears in the title bar of the window or the path to the represented file.
//   - [INSWindow.SetTitle]
//   - [INSWindow.Subtitle]: A secondary line of text that appears in the title bar of the window.
//   - [INSWindow.SetSubtitle]
//   - [INSWindow.TitleVisibility]: A value that indicates the visibility of the window’s title and title bar buttons.
//   - [INSWindow.SetTitleVisibility]
//   - [INSWindow.SetTitleWithRepresentedFilename]: Sets a given path as the window’s title, formatting it as a file-system path, and records this path as the window’s associated file.
//   - [INSWindow.RepresentedFilename]: The path to the file of the window’s represented file.
//   - [INSWindow.SetRepresentedFilename]
//   - [INSWindow.RepresentedURL]: The URL of the file the window represents.
//   - [INSWindow.SetRepresentedURL]
//
// # Accessing Screen Information
//
//   - [INSWindow.Screen]: The screen the window is on.
//   - [INSWindow.DeepestScreen]: The deepest screen the window is on (it may be split over several screens).
//   - [INSWindow.DisplaysWhenScreenProfileChanges]: A Boolean value that indicates whether the window context should be updated when the screen profile changes or when the window moves to a different screen.
//   - [INSWindow.SetDisplaysWhenScreenProfileChanges]
//
// # Moving Windows
//
//   - [INSWindow.MovableByWindowBackground]: A Boolean value that indicates whether the window is movable by clicking and dragging anywhere in its background.
//   - [INSWindow.SetMovableByWindowBackground]
//   - [INSWindow.Movable]: A Boolean value that indicates whether the window can be dragged by clicking in its title bar or background.
//   - [INSWindow.SetMovable]
//   - [INSWindow.Center]: Sets the window’s location to the center of the screen.
//
// # Closing Windows
//
//   - [INSWindow.PerformClose]: Simulates the user clicking the close button by momentarily highlighting the button and then closing the window.
//   - [INSWindow.Close]: Removes the window from the screen.
//   - [INSWindow.ReleasedWhenClosed]: A Boolean value that indicates whether the window is released when it receives the `close` message.
//   - [INSWindow.SetReleasedWhenClosed]
//
// # Minimizing Windows
//
//   - [INSWindow.Miniaturized]: A Boolean value that indicates whether the window is minimized.
//   - [INSWindow.PerformMiniaturize]: Simulates the user clicking the minimize button by momentarily highlighting the button, then minimizing the window.
//   - [INSWindow.Miniaturize]: Removes the window from the screen list and displays the minimized window in the Dock.
//   - [INSWindow.Deminiaturize]: De-minimizes the window.
//   - [INSWindow.MiniwindowImage]: The custom miniaturized window image of the window.
//   - [INSWindow.SetMiniwindowImage]
//   - [INSWindow.MiniwindowTitle]: The title displayed in the window’s minimized window.
//   - [INSWindow.SetMiniwindowTitle]
//
// # Getting the Dock Tile
//
//   - [INSWindow.DockTile]: The application’s Dock tile.
//
// # Printing Windows
//
//   - [INSWindow.Print]: Runs the Print panel, and if the user chooses an option other than canceling, prints the window (its frame view and all subviews).
//   - [INSWindow.DataWithEPSInsideRect]: Returns EPS data that draws the region of the window within a given rectangle.
//   - [INSWindow.DataWithPDFInsideRect]: Returns PDF data that draws the region of the window within a given rectangle.
//
// # Triggering Constraint-Based Layout
//
//   - [INSWindow.UpdateConstraintsIfNeeded]: Updates the constraints based on changes to views in the window since the last layout.
//   - [INSWindow.LayoutIfNeeded]: Updates the layout of views in the window based on the current views and constraints.
//
// # Debugging Constraint-Based Layout
//
//   - [INSWindow.VisualizeConstraints]: Displays a visual representation of the supplied constraints in the window.
//
// # Constraint-Based Layouts
//
//   - [INSWindow.AnchorAttributeForOrientation]: Returns the part of the window that stays stationary during constraint-based layout.
//   - [INSWindow.SetAnchorAttributeForOrientation]: Sets the part of the window that stays stationary during constraint-based layout.
//
// # Working with Window Depths
//
//   - [INSWindow.BitsPerPixel]: Returns the bits per pixel for the specified window depth.
//   - [INSWindow.SetNSBitsPerPixelFromDepth]
//   - [INSWindow.BitsPerSample]: Returns the bits per sample for the specified window depth.
//   - [INSWindow.SetNSBitsPerSampleFromDepth]
//   - [INSWindow.ColorSpaceName]: Returns the name of the color space corresponding to the passed window depth.
//   - [INSWindow.SetNSColorSpaceFromDepth]
//   - [INSWindow.NumberOfColorComponents]: Returns the number of color components in the specified color space.
//   - [INSWindow.SetNSNumberOfColorComponents]
//   - [INSWindow.IsPlanar]: Returns whether the specified window depth is planar.
//   - [INSWindow.SetNSPlanarFromDepth]
//   - [INSWindow.CanRepresentDisplayGamut]: A Boolean value that indicates if the window and its screen use a color space that can represent the specified display gamut.
//
// # Getting Information About Scripting Attributes
//
//   - [INSWindow.HasCloseBox]: A Boolean value that indicates if the window has a close box.
//   - [INSWindow.HasTitleBar]: A Boolean value that indicates if the window has a title bar.
//   - [INSWindow.ModalPanel]: A Boolean value that indicates whether the window is a modal panel.
//   - [INSWindow.FloatingPanel]: A Boolean value that indicates whether the window is a floating panel.
//   - [INSWindow.Zoomable]: A Boolean value that indicates whether the window allows zooming.
//   - [INSWindow.Resizable]: A Boolean value that indicates if the user can resize the window.
//   - [INSWindow.Miniaturizable]: A Boolean value that indicates whether the window can minimize.
//   - [INSWindow.OrderedIndex]: The zero-based position of the window, based on its order from front to back among all visible application windows.
//   - [INSWindow.SetOrderedIndex]
//
// # Setting Scripting Attributes
//
//   - [INSWindow.SetIsMiniaturized]: Sets the window’s miniaturized state to the value you specify.
//   - [INSWindow.SetIsVisible]: Sets the window’s visible state to the value you specify.
//   - [INSWindow.SetIsZoomed]: Sets the window’s zoomed state to the value you specify.
//
// # Handling Script Commands
//
//   - [INSWindow.HandleCloseScriptCommand]: Handles the AppleScript command to close the window (and its associated document, if any).
//   - [INSWindow.HandlePrintScriptCommand]: Handles the AppleScript command to print the contents of the window (or its associated document, if any).
//   - [INSWindow.HandleSaveScriptCommand]: Handles the AppleScript command to save the window (and its associated document, if any).
//
// # Instance Properties
//
//   - [INSWindow.CascadingReferenceFrame]
//
// # Instance Methods
//
//   - [INSWindow.BeginDraggingSessionWithItemsEventSource]
//   - [INSWindow.RequestSharingOfWindowCompletionHandler]
//   - [INSWindow.RequestSharingOfWindowUsingPreviewTitleCompletionHandler]
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow
type INSWindow interface {
	INSResponder

	// Topic: Creating a Window

	// Initializes the window with the specified values.
	InitWithContentRectStyleMaskBackingDefer(contentRect corefoundation.CGRect, style NSWindowStyleMask, backingStoreType NSBackingStoreType, flag bool) NSWindow
	// Initializes an allocated window with the specified values.
	InitWithContentRectStyleMaskBackingDeferScreen(contentRect corefoundation.CGRect, style NSWindowStyleMask, backingStoreType NSBackingStoreType, flag bool, screen INSScreen) NSWindow

	// Topic: Managing the Window’s Behavior

	// The window’s delegate.
	Delegate() NSWindowDelegate
	SetDelegate(value NSWindowDelegate)

	// Topic: Configuring the Window’s Content

	// The main content view controller for the window.
	ContentViewController() INSViewController
	SetContentViewController(value INSViewController)
	// The window’s content view, the highest accessible view object in the window’s view hierarchy.
	ContentView() INSView
	SetContentView(value INSView)

	// Topic: Configuring the Window’s Appearance

	// Flags that describe the window’s current style, such as if it’s resizable or in full-screen mode.
	StyleMask() NSWindowStyleMask
	SetStyleMask(value NSWindowStyleMask)
	// Takes the window into or out of fullscreen mode,
	ToggleFullScreen(sender objectivec.IObject)
	// A Boolean value that indicates whether the window is able to receive keyboard and mouse events even when some other window is being run modally.
	WorksWhenModal() bool
	// The window’s alpha value.
	AlphaValue() float64
	SetAlphaValue(value float64)
	// The color of the window’s background.
	BackgroundColor() INSColor
	SetBackgroundColor(value INSColor)
	// The window’s color space.
	ColorSpace() INSColorSpace
	SetColorSpace(value INSColorSpace)
	// Sets a Boolean value that indicates whether the window’s depth limit can change to match the depth of the screen it’s on.
	SetDynamicDepthLimit(flag bool)
	// A Boolean value that indicates whether the window can hide when its application becomes hidden.
	CanHide() bool
	SetCanHide(value bool)
	// A Boolean value that indicates whether the window is on the currently active space.
	OnActiveSpace() bool
	// A Boolean value that indicates whether the window is removed from the screen when its application becomes inactive.
	HidesOnDeactivate() bool
	SetHidesOnDeactivate(value bool)
	// A value that identifies the window’s behavior in window collections.
	CollectionBehavior() NSWindowCollectionBehavior
	SetCollectionBehavior(value NSWindowCollectionBehavior)
	// A Boolean value that indicates whether the window is opaque.
	Opaque() bool
	SetOpaque(value bool)
	// A Boolean value that indicates whether the window has a shadow.
	HasShadow() bool
	SetHasShadow(value bool)
	// Invalidates the window shadow so that it is recomputed based on the current window shape.
	InvalidateShadow()
	// Indicates whether the window calculates the thickness of a given border automatically.
	AutorecalculatesContentBorderThicknessForEdge(edge foundation.NSRectEdge) bool
	// Specifies whether the window calculates the thickness of a given border automatically.
	SetAutorecalculatesContentBorderThicknessForEdge(flag bool, edge foundation.NSRectEdge)
	// Indicates the thickness of a given border of the window.
	ContentBorderThicknessForEdge(edge foundation.NSRectEdge) float64
	// Specifies the thickness of a given border of the window.
	SetContentBorderThicknessForEdge(thickness float64, edge foundation.NSRectEdge)
	// A Boolean value that indicates whether the window prevents application termination when modal.
	PreventsApplicationTerminationWhenModal() bool
	SetPreventsApplicationTerminationWhenModal(value bool)
	// An object that the window inherits its appearance from.
	AppearanceSource() objectivec.Object
	SetAppearanceSource(value objectivec.Object)

	// Topic: Accessing Window Information

	// The depth limit of the window.
	DepthLimit() NSWindowDepth
	SetDepthLimit(value NSWindowDepth)
	// A Boolean value that indicates whether the window’s depth limit can change to match the depth of the screen it’s on.
	HasDynamicDepthLimit() bool
	// The window number of the window’s window device.
	WindowNumber() int
	// A dictionary containing information about the window’s resolution, such as color, depth, and so on.
	DeviceDescription() foundation.INSDictionary
	// A Boolean value that indicates whether the window can be displayed at the login window.
	CanBecomeVisibleWithoutLogin() bool
	SetCanBecomeVisibleWithoutLogin(value bool)
	// A Boolean value that indicates the level of access other processes have to the window’s content.
	SharingType() NSWindowSharingType
	SetSharingType(value NSWindowSharingType)
	// The window’s backing store type.
	BackingType() NSBackingStoreType
	SetBackingType(value NSBackingStoreType)
	// Returns a new display link whose callback will be invoked in-sync with the display the window is on.
	DisplayLinkWithTargetSelector(target objectivec.IObject, selector objc.SEL) quartzcore.CADisplayLink

	// Topic: Getting Layout Information

	// Returns the window’s content rectangle with a given frame rectangle.
	ContentRectForFrameRect(frameRect corefoundation.CGRect) corefoundation.CGRect
	// Returns the window’s frame rectangle with a given content rectangle.
	FrameRectForContentRect(contentRect corefoundation.CGRect) corefoundation.CGRect

	// Topic: Managing Windows

	// The window’s window controller.
	WindowController() INSWindowController
	SetWindowController(value INSWindowController)

	// Topic: Managing Sheets

	// The sheet attached to the window.
	AttachedSheet() INSWindow
	// A Boolean value that indicates whether the window has ever run as a modal sheet.
	Sheet() bool
	// Starts a document-modal session and presents—or queues for presentation—a sheet.
	BeginSheetCompletionHandler(sheetWindow INSWindow, handler ModalResponseHandler)
	// Starts a document-modal session and presents the specified critical sheet.
	BeginCriticalSheetCompletionHandler(sheetWindow INSWindow, handler ModalResponseHandler)
	// Ends a document-modal session and dismisses the specified sheet.
	EndSheet(sheetWindow INSWindow)
	// Ends a document-modal session and dismisses the specified sheet.
	EndSheetReturnCode(sheetWindow INSWindow, returnCode NSModalResponse)
	// The window to which the sheet is attached.
	SheetParent() INSWindow
	// An array of the sheets currently attached to the window.
	Sheets() []NSWindow

	// Topic: Sizing Windows

	// The window’s frame rectangle in screen coordinates, including the title bar.
	Frame() corefoundation.CGRect
	// Positions the bottom-left corner of the window’s frame rectangle at a given point in screen coordinates.
	SetFrameOrigin(point corefoundation.CGPoint)
	// Positions the top-left corner of the window’s frame rectangle at a given point in screen coordinates.
	SetFrameTopLeftPoint(point corefoundation.CGPoint)
	// Modifies and returns a frame rectangle so that its top edge lies on a specific screen.
	ConstrainFrameRectToScreen(frameRect corefoundation.CGRect, screen INSScreen) corefoundation.CGRect
	// Positions the window’s top-left to a given point.
	CascadeTopLeftFromPoint(topLeftPoint corefoundation.CGPoint) corefoundation.CGPoint
	// Sets the origin and size of the window’s frame rectangle according to a given frame rectangle, thereby setting its position and size onscreen.
	SetFrameDisplay(frameRect corefoundation.CGRect, flag bool)
	// Sets the origin and size of the window’s frame rectangle, with optional animation, according to a given frame rectangle, thereby setting its position and size onscreen.
	SetFrameDisplayAnimate(frameRect corefoundation.CGRect, displayFlag bool, animateFlag bool)
	// Specifies the duration of a smooth frame-size change.
	AnimationResizeTime(newFrame corefoundation.CGRect) float64
	// The window’s aspect ratio, which constrains the size of its frame rectangle to integral multiples of this ratio when the user resizes it.
	AspectRatio() corefoundation.CGSize
	SetAspectRatio(value corefoundation.CGSize)
	// The minimum size to which the window’s frame (including its title bar) can be sized.
	MinSize() corefoundation.CGSize
	SetMinSize(value corefoundation.CGSize)
	// The maximum size to which the window’s frame (including its title bar) can be sized.
	MaxSize() corefoundation.CGSize
	SetMaxSize(value corefoundation.CGSize)
	// A Boolean value that indicates whether the window is in a zoomed state.
	Zoomed() bool
	// This action method simulates the user clicking the zoom box by momentarily highlighting the button and then zooming the window.
	PerformZoom(sender objectivec.IObject)
	// Toggles the size and location of the window between its standard state (which the application provides as the best size to display the window’s data) and its user state (a new size and location the user may have set by moving or resizing the window).
	Zoom(sender objectivec.IObject)
	// The flags field of the event record for the mouse-down event that initiated the resizing session.
	ResizeFlags() NSEventModifierFlags
	// The window’s resizing increments.
	ResizeIncrements() corefoundation.CGSize
	SetResizeIncrements(value corefoundation.CGSize)
	// A Boolean value that indicates whether the window tries to optimize user-initiated resize operations by preserving the content of views that have not changed.
	PreservesContentDuringLiveResize() bool
	SetPreservesContentDuringLiveResize(value bool)
	// A Boolean value that indicates whether the window is being resized by the user.
	InLiveResize() bool

	// Topic: Sizing Content

	// The window’s content aspect ratio.
	ContentAspectRatio() corefoundation.CGSize
	SetContentAspectRatio(value corefoundation.CGSize)
	// The minimum size of the window’s content view in the window’s base coordinate system.
	ContentMinSize() corefoundation.CGSize
	SetContentMinSize(value corefoundation.CGSize)
	// Sets the size of the window’s content view to a given size, which is expressed in the window’s base coordinate system.
	SetContentSize(size corefoundation.CGSize)
	// The maximum size of the window’s content view in the window’s base coordinate system.
	ContentMaxSize() corefoundation.CGSize
	SetContentMaxSize(value corefoundation.CGSize)
	// The window’s content-view resizing increments.
	ContentResizeIncrements() corefoundation.CGSize
	SetContentResizeIncrements(value corefoundation.CGSize)
	// A value used by Auto Layout constraints to automatically bind to the value of [contentLayoutRect](<doc://com.apple.appkit/documentation/AppKit/NSWindow/contentLayoutRect>).
	ContentLayoutGuide() objectivec.IObject
	// The area inside the window that is for non-obscured content, in window coordinates.
	ContentLayoutRect() corefoundation.CGRect
	// A maximum size that is used to determine if a window can fit when it is in full screen in a tile.
	MaxFullScreenContentSize() corefoundation.CGSize
	SetMaxFullScreenContentSize(value corefoundation.CGSize)
	// A minimum size that is used to determine if a window can fit when it is in full screen in a tile.
	MinFullScreenContentSize() corefoundation.CGSize
	SetMinFullScreenContentSize(value corefoundation.CGSize)

	// Topic: Managing Window Layers

	// Removes the window from the screen list, which hides the window.
	OrderOut(sender objectivec.IObject)
	// Moves the window to the back of its level in the screen list, without changing either the key window or the main window.
	OrderBack(sender objectivec.IObject)
	// Moves the window to the front of its level in the screen list, without changing either the key window or the main window.
	OrderFront(sender objectivec.IObject)
	// Moves the window to the front of its level, even if its application isn’t active, without changing either the key window or the main window.
	OrderFrontRegardless()
	// Repositions the window’s window device in the window server’s screen list.
	OrderWindowRelativeTo(place NSWindowOrderingMode, otherWin int)
	// The window level of the window.
	Level() NSWindowLevel
	SetLevel(value NSWindowLevel)

	// Topic: Managing Window Visibility and Occlusion State

	// A Boolean value that indicates whether the window is visible onscreen (even when it’s obscured by other windows).
	Visible() bool
	// The occlusion state of the window.
	OcclusionState() NSWindowOcclusionState

	// Topic: Managing Window Frames in User Defaults

	// Sets the window’s frame rectangle by reading the rectangle data stored under a given name from the defaults system.
	SetFrameUsingName(name NSWindowFrameAutosaveName) bool
	// Sets the window’s frame rectangle by reading the rectangle data stored under a given name from the defaults system. Can operate on non-resizable windows.
	SetFrameUsingNameForce(name NSWindowFrameAutosaveName, force bool) bool
	// Saves the window’s frame rectangle in the user defaults system under a given name.
	SaveFrameUsingName(name NSWindowFrameAutosaveName)
	// The name used to automatically save the window’s frame rectangle data in the defaults system.
	FrameAutosaveName() NSWindowFrameAutosaveName
	// A string representation of the window’s frame rectangle.
	StringWithSavedFrame() NSWindowPersistableFrameDescriptor
	// Sets the window’s frame rectangle from a given string representation.
	SetFrameFromString(string_ NSWindowPersistableFrameDescriptor)

	// Topic: Managing Key Status

	// A Boolean value that indicates whether the window is the key window for the application.
	KeyWindow() bool
	// A Boolean value that indicates whether the window can become the key window.
	CanBecomeKeyWindow() bool
	// Makes the window the key window.
	MakeKeyWindow()
	// Moves the window to the front of the screen list, within its level, and makes it the key window; that is, it shows the window.
	MakeKeyAndOrderFront(sender objectivec.IObject)
	// Informs the window that it has become the key window.
	BecomeKeyWindow()
	// Resigns the window’s key window status.
	ResignKeyWindow()

	// Topic: Managing Main Status

	// A Boolean value that indicates whether the window is the application’s main window.
	MainWindow() bool
	// A Boolean value that indicates whether the window can become the application’s main window.
	CanBecomeMainWindow() bool
	// Makes the window the main window.
	MakeMainWindow()
	// Informs the window that it has become the main window.
	BecomeMainWindow()
	// Resigns the window’s main window status.
	ResignMainWindow()

	// Topic: Managing Toolbars

	// The window’s toolbar.
	Toolbar() INSToolbar
	SetToolbar(value INSToolbar)
	// Toggles the visibility of the window’s toolbar.
	ToggleToolbarShown(sender objectivec.IObject)
	// Presents the toolbar customization user interface.
	RunToolbarCustomizationPalette(sender objectivec.IObject)

	// Topic: Managing Attached Windows

	// An array of the window’s attached child windows.
	ChildWindows() []NSWindow
	// Adds a given window as a child window of the window.
	AddChildWindowOrdered(childWin INSWindow, place NSWindowOrderingMode)
	// Detaches a given child window from the window.
	RemoveChildWindow(childWin INSWindow)
	// The parent window to which the window is attached as a child.
	ParentWindow() INSWindow
	SetParentWindow(value INSWindow)

	// Topic: Managing Default Buttons

	// The button cell that performs as if clicked when the window receives a Return (or Enter) key event.
	DefaultButtonCell() INSButtonCell
	SetDefaultButtonCell(value INSButtonCell)
	// Reenables the default button cell’s key equivalent, so it performs a click when the user presses Return (or Enter).
	EnableKeyEquivalentForDefaultButtonCell()
	// Disables the default button cell’s key equivalent, so it doesn’t perform a click when the user presses Return (or Enter).
	DisableKeyEquivalentForDefaultButtonCell()

	// Topic: Managing Field Editors

	// Returns the window’s field editor, creating it if requested.
	FieldEditorForObject(createFlag bool, object objectivec.IObject) INSText
	// Forces the field editor to give up its first responder status and prepares it for its next assignment.
	EndEditingFor(object objectivec.IObject)

	// Topic: Managing the Window Menu

	// A Boolean value that indicates whether the window is excluded from the application’s Windows menu.
	ExcludedFromWindowsMenu() bool
	SetExcludedFromWindowsMenu(value bool)

	// Topic: Managing Cursor Rectangles

	// A Boolean value that indicates whether the window’s cursor rectangles are enabled.
	AreCursorRectsEnabled() bool
	// Reenables cursor rectangle management within the window after a [disableCursorRects()](<doc://com.apple.appkit/documentation/AppKit/NSWindow/disableCursorRects()>) message.
	EnableCursorRects()
	// Disables all cursor rectangle management within the window.
	DisableCursorRects()
	// Invalidates all cursor rectangles in the window.
	DiscardCursorRects()
	// Marks as invalid the cursor rectangles of a given view object in the window, so they’ll be set up again when the window becomes key.
	InvalidateCursorRectsForView(view INSView)
	// Clears the window’s cursor rectangles and the cursor rectangles of the [NSView](<doc://com.apple.appkit/documentation/AppKit/NSView>) objects in its view hierarchy.
	ResetCursorRects()

	// Topic: Managing Title Bars

	// Returns the window button of a given window button kind in the window’s view hierarchy.
	StandardWindowButton(b NSWindowButton) INSButton
	// A Boolean value that indicates whether the toolbar control button is currently displayed.
	ShowsToolbarButton() bool
	SetShowsToolbarButton(value bool)
	// A Boolean value that indicates whether the title bar draws its background.
	TitlebarAppearsTransparent() bool
	SetTitlebarAppearsTransparent(value bool)
	// The style that determines the appearance and location of the toolbar in relation to the title bar.
	ToolbarStyle() NSWindowToolbarStyle
	SetToolbarStyle(value NSWindowToolbarStyle)
	// The type of separator that the app displays between the title bar and content of a window.
	TitlebarSeparatorStyle() NSTitlebarSeparatorStyle
	SetTitlebarSeparatorStyle(value NSTitlebarSeparatorStyle)
	// The direction the window’s title bar lays text out, either left to right or right to left.
	WindowTitlebarLayoutDirection() NSUserInterfaceLayoutDirection

	// Topic: Managing Title Bar Accessories

	// Adds the specified title bar accessory view controller to the window.
	AddTitlebarAccessoryViewController(childViewController INSTitlebarAccessoryViewController)
	// Inserts the view controller into the window’s array of title bar accessory view controllers at the specified index.
	InsertTitlebarAccessoryViewControllerAtIndex(childViewController INSTitlebarAccessoryViewController, index int)
	// Removes the view controller at the specified index from the window’s array of title bar accessory view controllers.
	RemoveTitlebarAccessoryViewControllerAtIndex(index int)
	// An array of title bar accessory view controllers that are currently added to the window.
	TitlebarAccessoryViewControllers() []NSTitlebarAccessoryViewController
	SetTitlebarAccessoryViewControllers(value []NSTitlebarAccessoryViewController)

	// Topic: Managing Window Tabs

	// An object that represents information about a window when it displays as a tab.
	Tab() INSWindowTab
	// A value that allows a group of related windows.
	TabbingIdentifier() NSWindowTabbingIdentifier
	SetTabbingIdentifier(value NSWindowTabbingIdentifier)
	// Adds the provided window as a new tab in a tabbed window using the specified ordering instruction.
	AddTabbedWindowOrdered(window INSWindow, ordered NSWindowOrderingMode)
	// A value that indicates when a window displays tabs.
	TabbingMode() NSWindowTabbingMode
	SetTabbingMode(value NSWindowTabbingMode)
	// An array of windows that display as tabs.
	TabbedWindows() []NSWindow
	// Merges all open windows into a single tabbed window.
	MergeAllWindows(sender objectivec.IObject)
	// Selects the next tab in the tab group in the trailing direction.
	SelectNextTab(sender objectivec.IObject)
	// Selects the previous tab in the tab group in the leading direction.
	SelectPreviousTab(sender objectivec.IObject)
	// Moves the tab to a new containing window.
	MoveTabToNewWindow(sender objectivec.IObject)
	// Shows or hides the tab bar.
	ToggleTabBar(sender objectivec.IObject)
	// Shows or hides the tab overview.
	ToggleTabOverview(sender objectivec.IObject)
	// A group of windows that display together as a tab group.
	TabGroup() INSWindowTabGroup

	// Topic: Managing Tooltips

	// A Boolean value that indicates whether the window can display tooltips even when the application is in the background.
	AllowsToolTipsWhenApplicationIsInactive() bool
	SetAllowsToolTipsWhenApplicationIsInactive(value bool)

	// Topic: Handling Events

	// The event currently being processed by the application.
	CurrentEvent() INSEvent
	// Returns the next event matching a given mask.
	NextEventMatchingMask(mask NSEventMask) INSEvent
	// Forwards the message to the global application object.
	NextEventMatchingMaskUntilDateInModeDequeue(mask NSEventMask, expiration foundation.NSDate, mode foundation.NSString, deqFlag bool) INSEvent
	// Forwards the message to the global application object.
	DiscardEventsMatchingMaskBeforeEvent(mask NSEventMask, lastEvent INSEvent)
	// Forwards the message to the global application object.
	PostEventAtStart(event INSEvent, flag bool)
	// This action method dispatches mouse and keyboard events the global application object sends to the window.
	SendEvent(event INSEvent)

	// Topic: Managing Responders

	// The view that’s made first responder (also called the key view) the first time the window is placed onscreen.
	InitialFirstResponder() INSView
	SetInitialFirstResponder(value INSView)
	// The window’s first responder.
	FirstResponder() INSResponder
	// Attempts to make a given responder the first responder for the window.
	MakeFirstResponder(responder INSResponder) bool

	// Topic: Managing the Key View Loop

	// Gives key view status to the view that precedes the given view.
	SelectKeyViewPrecedingView(view INSView)
	// Gives key view status to the view that follows the given view.
	SelectKeyViewFollowingView(view INSView)
	// Searches for a candidate previous key view and, if it finds one, tries to make it the first responder.
	SelectPreviousKeyView(sender objectivec.IObject)
	// Searches for a candidate next key view and, if it finds one, tries to make it the first responder.
	SelectNextKeyView(sender objectivec.IObject)
	// The direction the window is currently using to change the key view.
	KeyViewSelectionDirection() NSSelectionDirection
	// A Boolean value that indicates whether the window automatically recalculates the key view loop when views are added.
	AutorecalculatesKeyViewLoop() bool
	SetAutorecalculatesKeyViewLoop(value bool)
	// Marks the key view loop as “dirty” and in need of recalculation.
	RecalculateKeyViewLoop()

	// Topic: Managing Window Sharing

	TransferWindowSharingToWindowCompletionHandler(window INSWindow, completionHandler ErrorHandler)
	// Indicates whether the receiver is the subject of an active SharePlay sharing session.
	HasActiveWindowSharingSession() bool

	// Topic: Handling Mouse Events

	// A Boolean value that indicates whether the window accepts mouse-moved events.
	AcceptsMouseMovedEvents() bool
	SetAcceptsMouseMovedEvents(value bool)
	// A Boolean value that indicates whether the window is transparent to mouse events.
	IgnoresMouseEvents() bool
	SetIgnoresMouseEvents(value bool)
	// The current location of the pointer reckoned in the window’s base coordinate system, regardless of the current event being handled or of any events pending.
	MouseLocationOutsideOfEventStream() corefoundation.CGPoint
	// Tracks events that match the specified mask using the specified tracking handler until the tracking handler explicitly terminates tracking.
	TrackEventsMatchingMaskTimeoutModeHandler(mask NSEventMask, timeout float64, mode foundation.NSString, trackingHandler EventHandler)
	// Starts a window drag based on the specified mouse-down event.
	PerformWindowDragWithEvent(event INSEvent)

	// Topic: Handling Window Restoration

	// A Boolean value indicating whether the window configuration is preserved between application launches.
	Restorable() bool
	SetRestorable(value bool)
	// The restoration class associated with the window.
	RestorationClass() objc.Class
	SetRestorationClass(value objc.Class)
	// Disables snapshot restoration.
	DisableSnapshotRestoration()
	// Enables snapshot restoration.
	EnableSnapshotRestoration()

	// Topic: Drawing Windows

	// Passes a display message down the window’s view hierarchy, thus redrawing all views within the window.
	Display()
	// Passes a display message down the window’s view hierarchy, thus redrawing all views that need displaying.
	DisplayIfNeeded()
	// A Boolean value that indicates whether any of the window’s views need to be displayed.
	ViewsNeedDisplay() bool
	SetViewsNeedDisplay(value bool)
	// A Boolean value that indicates whether the window allows multithreaded view drawing.
	AllowsConcurrentViewDrawing() bool
	SetAllowsConcurrentViewDrawing(value bool)

	// Topic: Window Animation

	// The window’s automatic animation behavior.
	AnimationBehavior() NSWindowAnimationBehavior
	SetAnimationBehavior(value NSWindowAnimationBehavior)

	// Topic: Updating Windows

	// Updates the window.
	Update()

	// Topic: Dragging Items

	// Registers a set of pasteboard types that the window accepts as the destination of an image-dragging session.
	RegisterForDraggedTypes(newTypes []string)
	// Unregisters the window as a possible destination for dragging operations.
	UnregisterDraggedTypes()

	// Topic: Accessing Edited Status

	// A Boolean value that indicates whether the window’s document has been edited.
	DocumentEdited() bool
	SetDocumentEdited(value bool)

	// Topic: Converting Coordinates

	// The backing scale factor.
	BackingScaleFactor() float64
	// Returns a backing store pixel-aligned rectangle in window coordinates.
	BackingAlignedRectOptions(rect corefoundation.CGRect, options foundation.AlignmentOptions) corefoundation.CGRect
	// Converts a rectangle from its pixel-aligned backing store coordinate system to the window’s coordinate system.
	ConvertRectFromBacking(rect corefoundation.CGRect) corefoundation.CGRect
	// Converts a rectangle from the screen coordinate system to the window’s coordinate system.
	ConvertRectFromScreen(rect corefoundation.CGRect) corefoundation.CGRect
	// Converts a point from its pixel-aligned backing store coordinate system to the window’s coordinate system.
	ConvertPointFromBacking(point corefoundation.CGPoint) corefoundation.CGPoint
	// Converts a point from the screen coordinate system to the window’s coordinate system.
	ConvertPointFromScreen(point corefoundation.CGPoint) corefoundation.CGPoint
	// Converts a rectangle from the window’s coordinate system to its pixel-aligned backing store coordinate system.
	ConvertRectToBacking(rect corefoundation.CGRect) corefoundation.CGRect
	// Converts a rectangle to the screen coordinate system from the window’s coordinate system.
	ConvertRectToScreen(rect corefoundation.CGRect) corefoundation.CGRect
	// Converts a point from the window’s coordinate system to its pixel-aligned backing store coordinate system.
	ConvertPointToBacking(point corefoundation.CGPoint) corefoundation.CGPoint
	// Converts a point to the screen coordinate system from the window’s coordinate system.
	ConvertPointToScreen(point corefoundation.CGPoint) corefoundation.CGPoint

	// Topic: Managing Titles

	// The string that appears in the title bar of the window or the path to the represented file.
	Title() string
	SetTitle(value string)
	// A secondary line of text that appears in the title bar of the window.
	Subtitle() string
	SetSubtitle(value string)
	// A value that indicates the visibility of the window’s title and title bar buttons.
	TitleVisibility() NSWindowTitleVisibility
	SetTitleVisibility(value NSWindowTitleVisibility)
	// Sets a given path as the window’s title, formatting it as a file-system path, and records this path as the window’s associated file.
	SetTitleWithRepresentedFilename(filename string)
	// The path to the file of the window’s represented file.
	RepresentedFilename() string
	SetRepresentedFilename(value string)
	// The URL of the file the window represents.
	RepresentedURL() foundation.NSURL
	SetRepresentedURL(value foundation.NSURL)

	// Topic: Accessing Screen Information

	// The screen the window is on.
	Screen() INSScreen
	// The deepest screen the window is on (it may be split over several screens).
	DeepestScreen() INSScreen
	// A Boolean value that indicates whether the window context should be updated when the screen profile changes or when the window moves to a different screen.
	DisplaysWhenScreenProfileChanges() bool
	SetDisplaysWhenScreenProfileChanges(value bool)

	// Topic: Moving Windows

	// A Boolean value that indicates whether the window is movable by clicking and dragging anywhere in its background.
	MovableByWindowBackground() bool
	SetMovableByWindowBackground(value bool)
	// A Boolean value that indicates whether the window can be dragged by clicking in its title bar or background.
	Movable() bool
	SetMovable(value bool)
	// Sets the window’s location to the center of the screen.
	Center()

	// Topic: Closing Windows

	// Simulates the user clicking the close button by momentarily highlighting the button and then closing the window.
	PerformClose(sender objectivec.IObject)
	// Removes the window from the screen.
	Close()
	// A Boolean value that indicates whether the window is released when it receives the `close` message.
	ReleasedWhenClosed() bool
	SetReleasedWhenClosed(value bool)

	// Topic: Minimizing Windows

	// A Boolean value that indicates whether the window is minimized.
	Miniaturized() bool
	// Simulates the user clicking the minimize button by momentarily highlighting the button, then minimizing the window.
	PerformMiniaturize(sender objectivec.IObject)
	// Removes the window from the screen list and displays the minimized window in the Dock.
	Miniaturize(sender objectivec.IObject)
	// De-minimizes the window.
	Deminiaturize(sender objectivec.IObject)
	// The custom miniaturized window image of the window.
	MiniwindowImage() INSImage
	SetMiniwindowImage(value INSImage)
	// The title displayed in the window’s minimized window.
	MiniwindowTitle() string
	SetMiniwindowTitle(value string)

	// Topic: Getting the Dock Tile

	// The application’s Dock tile.
	DockTile() INSDockTile

	// Topic: Printing Windows

	// Runs the Print panel, and if the user chooses an option other than canceling, prints the window (its frame view and all subviews).
	Print(sender objectivec.IObject)
	// Returns EPS data that draws the region of the window within a given rectangle.
	DataWithEPSInsideRect(rect corefoundation.CGRect) foundation.NSData
	// Returns PDF data that draws the region of the window within a given rectangle.
	DataWithPDFInsideRect(rect corefoundation.CGRect) foundation.NSData

	// Topic: Triggering Constraint-Based Layout

	// Updates the constraints based on changes to views in the window since the last layout.
	UpdateConstraintsIfNeeded()
	// Updates the layout of views in the window based on the current views and constraints.
	LayoutIfNeeded()

	// Topic: Debugging Constraint-Based Layout

	// Displays a visual representation of the supplied constraints in the window.
	VisualizeConstraints(constraints []NSLayoutConstraint)

	// Topic: Constraint-Based Layouts

	// Returns the part of the window that stays stationary during constraint-based layout.
	AnchorAttributeForOrientation(orientation NSLayoutConstraintOrientation) NSLayoutAttribute
	// Sets the part of the window that stays stationary during constraint-based layout.
	SetAnchorAttributeForOrientation(attr NSLayoutAttribute, orientation NSLayoutConstraintOrientation)

	// Topic: Working with Window Depths

	// Returns the bits per pixel for the specified window depth.
	BitsPerPixel() int
	SetNSBitsPerPixelFromDepth(value int)
	// Returns the bits per sample for the specified window depth.
	BitsPerSample() int
	SetNSBitsPerSampleFromDepth(value int)
	// Returns the name of the color space corresponding to the passed window depth.
	ColorSpaceName() NSColorSpaceName
	SetNSColorSpaceFromDepth(value NSColorSpaceName)
	// Returns the number of color components in the specified color space.
	NumberOfColorComponents() int
	SetNSNumberOfColorComponents(value int)
	// Returns whether the specified window depth is planar.
	IsPlanar() bool
	SetNSPlanarFromDepth(value bool)
	// A Boolean value that indicates if the window and its screen use a color space that can represent the specified display gamut.
	CanRepresentDisplayGamut(displayGamut NSDisplayGamut) bool

	// Topic: Getting Information About Scripting Attributes

	// A Boolean value that indicates if the window has a close box.
	HasCloseBox() bool
	// A Boolean value that indicates if the window has a title bar.
	HasTitleBar() bool
	// A Boolean value that indicates whether the window is a modal panel.
	ModalPanel() bool
	// A Boolean value that indicates whether the window is a floating panel.
	FloatingPanel() bool
	// A Boolean value that indicates whether the window allows zooming.
	Zoomable() bool
	// A Boolean value that indicates if the user can resize the window.
	Resizable() bool
	// A Boolean value that indicates whether the window can minimize.
	Miniaturizable() bool
	// The zero-based position of the window, based on its order from front to back among all visible application windows.
	OrderedIndex() int
	SetOrderedIndex(value int)

	// Topic: Setting Scripting Attributes

	// Sets the window’s miniaturized state to the value you specify.
	SetIsMiniaturized(flag bool)
	// Sets the window’s visible state to the value you specify.
	SetIsVisible(flag bool)
	// Sets the window’s zoomed state to the value you specify.
	SetIsZoomed(flag bool)

	// Topic: Handling Script Commands

	// Handles the AppleScript command to close the window (and its associated document, if any).
	HandleCloseScriptCommand(command foundation.NSCloseCommand) objectivec.IObject
	// Handles the AppleScript command to print the contents of the window (or its associated document, if any).
	HandlePrintScriptCommand(command foundation.NSScriptCommand) objectivec.IObject
	// Handles the AppleScript command to save the window (and its associated document, if any).
	HandleSaveScriptCommand(command foundation.NSScriptCommand) objectivec.IObject

	// Topic: Instance Properties

	CascadingReferenceFrame() corefoundation.CGRect

	// Topic: Instance Methods

	BeginDraggingSessionWithItemsEventSource(items []NSDraggingItem, event INSEvent, source NSDraggingSource) INSDraggingSession
	RequestSharingOfWindowCompletionHandler(window INSWindow, completionHandler ErrorHandler)
	RequestSharingOfWindowUsingPreviewTitleCompletionHandler(image INSImage, title string, completionHandler ErrorHandler)

	// Name of an exception that occurs when you pass an invalid argument to a method, such as a `nil` pointer where a non-`nil` object is required.
	InvalidArgumentException() foundation.NSString
	// A Boolean value that indicates whether the window’s resize indicator is visible.
	ShowsResizeIndicator() bool
	SetShowsResizeIndicator(value bool)
	// The Carbon window reference associated with the window, creating one if necessary.
	WindowRef() WindowRef
	SetWindowRef(value WindowRef)
	// Returns the animation that should be performed for the specified key.
	AnimationForKey(key NSAnimatablePropertyKey) objectivec.IObject
	// Sets the option dictionary that maps event trigger keys to animation objects.
	Animations() foundation.INSDictionary
	// Returns a proxy object for the receiver that can be used to initiate implied animation for property changes.
	Animator() INSWindow
}

// Init initializes the instance.
func (w NSWindow) Init() NSWindow {
	rv := objc.Send[NSWindow](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w NSWindow) Autorelease() NSWindow {
	rv := objc.Send[NSWindow](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSWindow creates a new NSWindow instance.
func NewNSWindow() NSWindow {
	class := getNSWindowClass()
	rv := objc.Send[NSWindow](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new responder object with data in an unarchiver.
//
// coder: An unarchiver object.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/init(coder:)
func NewWindowWithCoder(coder foundation.INSCoder) NSWindow {
	instance := getNSWindowClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSWindowFromID(rv)
}

// Initializes the window with the specified values.
//
// contentRect: Origin and size of the window’s content area in screen coordinates. Note
// that the window server limits window position coordinates to ±16,000 and
// sizes to 10,000.
//
// style: The window’s style. It can be [NSBorderlessWindowMask], or it can contain
// any of the options described in [NSWindow.StyleMask], combined using the C
// bitwise OR operator. Borderless windows display none of the usual
// peripheral elements and are generally useful only for display or caching
// purposes; you should normally not need to create them. Also, note that a
// window’s style mask should include [NSTitledWindowMask] if it includes
// any of the others.
//
// backingStoreType: Specifies how the drawing done in the window is buffered by the window
// device, and possible values are described in [NSWindow.BackingStoreType].
//
// flag: Specifies whether the window server creates a window device for the window
// immediately. When true, the window server defers creating the window device
// until the window is moved onscreen. All display messages sent to the window
// or its views are postponed until the window is created, just before it’s
// moved onscreen.
//
// # Return Value
//
// The initialized window.
//
// # Discussion
//
// This method is the designated initializer for the [NSWindow] class.
//
// Deferring the creation of the window improves launch time and minimizes the
// virtual memory load on the window server.
//
// The new window creates a view to be its default content view. You can
// replace it with your own object by setting the [ContentView] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/init(contentRect:styleMask:backing:defer:)
//
// [NSWindow.StyleMask]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
// [NSWindow.BackingStoreType]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType
func NewWindowWithContentRectStyleMaskBackingDefer(contentRect corefoundation.CGRect, style NSWindowStyleMask, backingStoreType NSBackingStoreType, flag bool) NSWindow {
	instance := getNSWindowClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentRect:styleMask:backing:defer:"), contentRect, style, backingStoreType, flag)
	return NSWindowFromID(rv)
}

// Initializes an allocated window with the specified values.
//
// contentRect: Origin and size of the window’s content area in screen coordinates. The
// origin is relative to the origin of the provided screen. Note that the
// window server limits window position coordinates to ±16,000 and sizes to
// 10,000.
//
// style: The window’s style. It can be [NSBorderlessWindowMask], or it can contain
// any of the options described in [NSWindow.StyleMask], combined using the C
// bitwise OR operator. Borderless windows display none of the usual
// peripheral elements and are generally useful only for display or caching
// purposes; you should not usually need to create them. Also, note that a
// window’s style mask should include [NSTitledWindowMask] if it includes
// any of the others.
//
// backingStoreType: Specifies how the drawing done in the window is buffered by the window
// device; possible values are described in [NSWindow.BackingStoreType].
//
// flag: Specifies whether the window server creates a window device for the window
// immediately. When true, the window server defers creating the window device
// until the window is moved onscreen. All display messages sent to the window
// or its views are postponed until the window is created, just before it’s
// moved onscreen.
//
// screen: Specifies the screen on which the window is positioned. The content
// rectangle is positioned relative to the bottom-left corner of `screen`.
// When `nil`, the content rectangle is positioned relative to (0, 0), which
// is the origin of the primary screen.
//
// # Return Value
//
// The initialized window.
//
// # Discussion
//
// The primary screen is the one that contains the current key window or, if
// there is no key window, the one that contains the main menu. If there’s
// neither a key window nor a main menu (if there’s no active application),
// the primary screen is the one where the origin of the screen coordinate
// system is located.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/init(contentRect:styleMask:backing:defer:screen:)
//
// [NSWindow.StyleMask]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
// [NSWindow.BackingStoreType]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType
func NewWindowWithContentRectStyleMaskBackingDeferScreen(contentRect corefoundation.CGRect, style NSWindowStyleMask, backingStoreType NSBackingStoreType, flag bool, screen INSScreen) NSWindow {
	instance := getNSWindowClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, screen)
	return NSWindowFromID(rv)
}

// Creates a titled window that contains the specified content view
// controller.
//
// contentViewController: The view controller that provides the main content view for the window. The
// window’s [ContentView] property is set to
// `contentViewController“XCUIElementTypeView`.
//
// # Return Value
//
// A window with the content view controller set to the passed-in view
// controller object.
//
// # Discussion
//
// This method creates a basic window object that is titled, closable,
// resizable, and miniaturizable. By default, the window’s title is
// automatically bound to the title of `contentViewController`. You can
// control the size of the window by using Auto Layout and applying size
// constraints to the view or its subviews. The initial size of the window is
// set to the initial size of [ContentView] (that is, the size of
// `contentViewController“XCUIElementTypeView`). The newly created window has
// [ReleasedWhenClosed] set to false, and it must be explicitly retained to
// keep the window instance alive.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/init(contentViewController:)
func NewWindowWithContentViewController(contentViewController INSViewController) NSWindow {
	rv := objc.Send[objc.ID](objc.ID(getNSWindowClass().class), objc.Sel("windowWithContentViewController:"), contentViewController)
	return NSWindowFromID(rv)
}

// Returns a Cocoa window created from a Carbon window.
//
// windowRef: The Carbon [WindowRef] object to use to create the Cocoa window.
//
// # Return Value
//
// A Cocoa window created from `windowRef`.
//
// # Discussion
//
// For more information on Carbon-Cocoa integration, see Using a Carbon User
// Interface in a Cocoa Application in Carbon-Cocoa Integration Guide.
//
// # Special Considerations
//
// For historical reasons, contrary to normal memory management policy “ does
// retain `windowRef`. It is therefore recommended that you make sure you
// retain `windowRef` before calling this method. If `windowRef` is still
// valid when the Cocoa window is deallocated, the Cocoa window will release
// it.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/init(windowRef:)
func NewWindowWithWindowRef(windowRef WindowRef) NSWindow {
	instance := getNSWindowClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWindowRef:"), windowRef)
	return NSWindowFromID(rv)
}

// Initializes the window with the specified values.
//
// contentRect: Origin and size of the window’s content area in screen coordinates. Note
// that the window server limits window position coordinates to ±16,000 and
// sizes to 10,000.
//
// style: The window’s style. It can be [NSBorderlessWindowMask], or it can contain
// any of the options described in [NSWindow.StyleMask], combined using the C
// bitwise OR operator. Borderless windows display none of the usual
// peripheral elements and are generally useful only for display or caching
// purposes; you should normally not need to create them. Also, note that a
// window’s style mask should include [NSTitledWindowMask] if it includes
// any of the others.
//
// backingStoreType: Specifies how the drawing done in the window is buffered by the window
// device, and possible values are described in [NSWindow.BackingStoreType].
//
// flag: Specifies whether the window server creates a window device for the window
// immediately. When true, the window server defers creating the window device
// until the window is moved onscreen. All display messages sent to the window
// or its views are postponed until the window is created, just before it’s
// moved onscreen.
//
// # Return Value
//
// The initialized window.
//
// # Discussion
//
// This method is the designated initializer for the [NSWindow] class.
//
// Deferring the creation of the window improves launch time and minimizes the
// virtual memory load on the window server.
//
// The new window creates a view to be its default content view. You can
// replace it with your own object by setting the [ContentView] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/init(contentRect:styleMask:backing:defer:)
//
// [NSWindow.StyleMask]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
// [NSWindow.BackingStoreType]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType
func (w NSWindow) InitWithContentRectStyleMaskBackingDefer(contentRect corefoundation.CGRect, style NSWindowStyleMask, backingStoreType NSBackingStoreType, flag bool) NSWindow {
	rv := objc.Send[NSWindow](w.ID, objc.Sel("initWithContentRect:styleMask:backing:defer:"), contentRect, style, backingStoreType, flag)
	return rv
}

// Initializes an allocated window with the specified values.
//
// contentRect: Origin and size of the window’s content area in screen coordinates. The
// origin is relative to the origin of the provided screen. Note that the
// window server limits window position coordinates to ±16,000 and sizes to
// 10,000.
//
// style: The window’s style. It can be [NSBorderlessWindowMask], or it can contain
// any of the options described in [NSWindow.StyleMask], combined using the C
// bitwise OR operator. Borderless windows display none of the usual
// peripheral elements and are generally useful only for display or caching
// purposes; you should not usually need to create them. Also, note that a
// window’s style mask should include [NSTitledWindowMask] if it includes
// any of the others.
//
// backingStoreType: Specifies how the drawing done in the window is buffered by the window
// device; possible values are described in [NSWindow.BackingStoreType].
//
// flag: Specifies whether the window server creates a window device for the window
// immediately. When true, the window server defers creating the window device
// until the window is moved onscreen. All display messages sent to the window
// or its views are postponed until the window is created, just before it’s
// moved onscreen.
//
// screen: Specifies the screen on which the window is positioned. The content
// rectangle is positioned relative to the bottom-left corner of `screen`.
// When `nil`, the content rectangle is positioned relative to (0, 0), which
// is the origin of the primary screen.
//
// # Return Value
//
// The initialized window.
//
// # Discussion
//
// The primary screen is the one that contains the current key window or, if
// there is no key window, the one that contains the main menu. If there’s
// neither a key window nor a main menu (if there’s no active application),
// the primary screen is the one where the origin of the screen coordinate
// system is located.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/init(contentRect:styleMask:backing:defer:screen:)
//
// [NSWindow.StyleMask]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
// [NSWindow.BackingStoreType]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType
func (w NSWindow) InitWithContentRectStyleMaskBackingDeferScreen(contentRect corefoundation.CGRect, style NSWindowStyleMask, backingStoreType NSBackingStoreType, flag bool, screen INSScreen) NSWindow {
	rv := objc.Send[NSWindow](w.ID, objc.Sel("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, screen)
	return rv
}

// Takes the window into or out of fullscreen mode,
//
// sender: The object that sent the message.
//
// # Discussion
//
// If an application supports fullscreen, it should add a menu item to the
// View menu with “ as the action, and `nil` as the target.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/toggleFullScreen(_:)
func (w NSWindow) ToggleFullScreen(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("toggleFullScreen:"), sender)
}

// Sets a Boolean value that indicates whether the window’s depth limit can
// change to match the depth of the screen it’s on.
//
// flag: true if the window’s depth can change; otherwise, false.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/setDynamicDepthLimit(_:)
func (w NSWindow) SetDynamicDepthLimit(flag bool) {
	objc.Send[objc.ID](w.ID, objc.Sel("setDynamicDepthLimit:"), flag)
}

// Invalidates the window shadow so that it is recomputed based on the current
// window shape.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/invalidateShadow()
func (w NSWindow) InvalidateShadow() {
	objc.Send[objc.ID](w.ID, objc.Sel("invalidateShadow"))
}

// Indicates whether the window calculates the thickness of a given border
// automatically.
//
// edge: The border to check:
//
// - [NSMaxYEdge]: Top border.
// - [NSMinYEdge]: Bottom border.
//
// # Return Value
//
// true when the window auto-recalculates the given border’s thickness;
// otherwise, false.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/autorecalculatesContentBorderThickness(for:)
func (w NSWindow) AutorecalculatesContentBorderThicknessForEdge(edge foundation.NSRectEdge) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("autorecalculatesContentBorderThicknessForEdge:"), edge)
	return rv
}

// Specifies whether the window calculates the thickness of a given border
// automatically.
//
// flag: If true, the window calculates the thickness of the edge automatically; if
// false, it does not.
//
// edge: The border to set auto-recalculation on or off:
//
// - [NSMaxYEdge]: Top border.
// - [NSMinYEdge]: Bottom border.
//
// # Discussion
//
// Turning off a border’s auto-recalculation status sets its border
// thickness to `0.0`.
//
// In a nontextured window calling “ passing [NSMaxYEdge] will raise an
// exception (in a nontextured window, it’s only valid to set the content
// border thickness of the bottom edge). It is only valid to set the content
// border thickness of the top edge in a textured window.
//
// Typically, if you call “, you should also call `NO `.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/setAutorecalculatesContentBorderThickness(_:for:)
func (w NSWindow) SetAutorecalculatesContentBorderThicknessForEdge(flag bool, edge foundation.NSRectEdge) {
	objc.Send[objc.ID](w.ID, objc.Sel("setAutorecalculatesContentBorderThickness:forEdge:"), flag, edge)
}

// Indicates the thickness of a given border of the window.
//
// edge: The border whose thickness to get:
//
// - [NSMaxYEdge]: Top border.
// - [NSMinYEdge]: Bottom border.
//
// # Return Value
//
// Thickness of the given border, in points.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/contentBorderThickness(for:)
func (w NSWindow) ContentBorderThicknessForEdge(edge foundation.NSRectEdge) float64 {
	rv := objc.Send[float64](w.ID, objc.Sel("contentBorderThicknessForEdge:"), edge)
	return rv
}

// Specifies the thickness of a given border of the window.
//
// thickness: The thickness for `edge`, in points.
//
// edge: The border whose thickness to set:
//
// - [NSMaxYEdge]: Top border.
// - [NSMinYEdge]: Bottom border.
//
// # Discussion
//
// In a nontextured window calling “ passing [NSMaxYEdge] will raise an
// exception (in a nontextured window, it’s only valid to set the content
// border thickness of the bottom edge). It is only valid to set the content
// border thickness of the top edge in a textured window.
//
// Typically, if you call “, you should also call `NO `.
//
// The `contentBorder` does not include the title bar or toolbar, so a
// textured window that just wants the gradient in the title bar and toolbar
// should have a `thickness` of `0` for [NSMaxYEdge].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/setContentBorderThickness(_:for:)
func (w NSWindow) SetContentBorderThicknessForEdge(thickness float64, edge foundation.NSRectEdge) {
	objc.Send[objc.ID](w.ID, objc.Sel("setContentBorderThickness:forEdge:"), thickness, edge)
}

// Returns a new display link whose callback will be invoked in-sync with the
// display the window is on.
//
// # Discussion
//
// If the window is not on any display the callback will not be invoked.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/displayLink(target:selector:)
func (w NSWindow) DisplayLinkWithTargetSelector(target objectivec.IObject, selector objc.SEL) quartzcore.CADisplayLink {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("displayLinkWithTarget:selector:"), target, selector)
	return quartzcore.CADisplayLinkFromID(rv)
}

// Returns the window’s content rectangle with a given frame rectangle.
//
// frameRect: The frame rectangle for the window expressed in screen coordinates.
//
// # Return Value
//
// The window’s content rectangle, expressed in screen coordinates, with
// f`rameRect`.
//
// # Discussion
//
// The window uses its current style mask in computing the content rectangle.
// See [NSWindow.StyleMask] for a list of style mask values. The main
// advantage of this instance-method counterpart to
// [ContentRectForFrameRectStyleMask] is that it allows you to take toolbars
// into account when converting between content and frame rectangles. (The
// toolbar is not included in the content rectangle.)
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/contentRect(forFrameRect:)
//
// [NSWindow.StyleMask]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
func (w NSWindow) ContentRectForFrameRect(frameRect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](w.ID, objc.Sel("contentRectForFrameRect:"), frameRect)
	return corefoundation.CGRect(rv)
}

// Returns the window’s frame rectangle with a given content rectangle.
//
// contentRect: The content rectangle for the window expressed in screen coordinates.
//
// # Return Value
//
// The window’s frame rectangle, expressed in screen coordinates, with
// `contentRect`.
//
// # Discussion
//
// The window uses its current style mask in computing the frame rectangle.
// See [NSWindow.StyleMask] for a list of style mask values. The major
// advantage of this instance-method counterpart to
// [FrameRectForContentRectStyleMask] is that it allows you to take toolbars
// into account when converting between content and frame rectangles. (The
// toolbar is included in the frame rectangle but not the content rectangle.)
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/frameRect(forContentRect:)
//
// [NSWindow.StyleMask]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
func (w NSWindow) FrameRectForContentRect(contentRect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](w.ID, objc.Sel("frameRectForContentRect:"), contentRect)
	return corefoundation.CGRect(rv)
}

// Starts a document-modal session and presents—or queues for
// presentation—a sheet.
//
// sheetWindow: The window object that represents the sheet to present.
//
// handler: The completion handler that gets called when the sheet’s modal session
// ends.
//
// # Discussion
//
// If the window already has a presented sheet, this method queues the
// specified sheet for presentation after the current sheet is dismissed and
// then returns control to the caller.
//
// If the window has no presented sheets, this method displays the specified
// sheet, makes it key, and returns control to the caller. While the sheet
// remains visible, most events targeted at the receiver are prohibited. The
// runloop does not enter any special mode to accomplish this.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/beginSheet(_:completionHandler:)
func (w NSWindow) BeginSheetCompletionHandler(sheetWindow INSWindow, handler ModalResponseHandler) {
	_block1, _ := NewModalResponseBlock(handler)
	objc.Send[objc.ID](w.ID, objc.Sel("beginSheet:completionHandler:"), sheetWindow, _block1)
}

// Starts a document-modal session and presents the specified critical sheet.
//
// sheetWindow: The window object that represents the critical sheet to present. A critical
// sheet contains content that is time-critical or very important to the user.
//
// handler: The completion handler that gets called when the sheet’s modal session
// ends.
//
// # Discussion
//
// This method displays the sheet—on top of the window’s current sheet, if
// one exists—makes it key and returns control to the caller. While the
// sheet remains visible, most events targeted at the receiver are prohibited.
// The runloop does not enter any special mode to accomplish this.
//
// If the window already has a sheet when this method runs, the existing sheet
// is temporarily disabled while the critical sheet is presented. When the
// critical sheet is dismissed, the previously presented sheet continues its
// standard operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/beginCriticalSheet(_:completionHandler:)
func (w NSWindow) BeginCriticalSheetCompletionHandler(sheetWindow INSWindow, handler ModalResponseHandler) {
	_block1, _ := NewModalResponseBlock(handler)
	objc.Send[objc.ID](w.ID, objc.Sel("beginCriticalSheet:completionHandler:"), sheetWindow, _block1)
}

// Ends a document-modal session and dismisses the specified sheet.
//
// sheetWindow: The window object that represents the sheet to be dismissed.
//
// # Discussion
//
// This method ends the modal session with the return code
// [NSModalResponseStop].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/endSheet(_:)-4dmmq
func (w NSWindow) EndSheet(sheetWindow INSWindow) {
	objc.Send[objc.ID](w.ID, objc.Sel("endSheet:"), sheetWindow)
}

// Ends a document-modal session and dismisses the specified sheet.
//
// sheetWindow: The window object that represents the sheet to dismiss.
//
// returnCode: The return code to send to the completion handler. You can use a custom
// value that you define or one of the return codes defined in the
// [NSModalResponse] enumeration or `Additional NSModalResponse Values`.
//
// # Discussion
//
// This method ends the modal session with the specified return code.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/endSheet(_:returnCode:)
func (w NSWindow) EndSheetReturnCode(sheetWindow INSWindow, returnCode NSModalResponse) {
	objc.Send[objc.ID](w.ID, objc.Sel("endSheet:returnCode:"), sheetWindow, returnCode)
}

// Positions the bottom-left corner of the window’s frame rectangle at a
// given point in screen coordinates.
//
// point: The new position of the window’s bottom-left corner in screen
// coordinates.
//
// # Discussion
//
// Note that the window server limits window position coordinates to ±16,000.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/setFrameOrigin(_:)
func (w NSWindow) SetFrameOrigin(point corefoundation.CGPoint) {
	objc.Send[objc.ID](w.ID, objc.Sel("setFrameOrigin:"), point)
}

// Positions the top-left corner of the window’s frame rectangle at a given
// point in screen coordinates.
//
// point: The new position of the window’s top-left corner in screen coordinates.
//
// # Discussion
//
// Note that the window server limits window position coordinates to ±16,000;
// if necessary, adjust `aPoint` relative to the window’s lower-left corner
// to account for this limit.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/setFrameTopLeftPoint(_:)
func (w NSWindow) SetFrameTopLeftPoint(point corefoundation.CGPoint) {
	objc.Send[objc.ID](w.ID, objc.Sel("setFrameTopLeftPoint:"), point)
}

// Modifies and returns a frame rectangle so that its top edge lies on a
// specific screen.
//
// frameRect: The proposed frame rectangle to adjust.
//
// screen: The screen on which the top edge of the window’s frame is to lie.
//
// # Return Value
//
// The adjusted frame rectangle.
//
// # Discussion
//
// If the window is resizable and the window’s height is greater than the
// screen height, the rectangle’s height is adjusted to fit within the
// screen as well. The rectangle’s width and horizontal location are
// unaffected. You shouldn’t need to invoke this method yourself; it’s
// invoked automatically (and the modified frame is used to locate and set the
// size of the window) whenever a titled [NSWindow] object is placed onscreen
// and whenever its size is changed.
//
// Subclasses can override this method to prevent their instances from being
// constrained or to constrain them differently.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/constrainFrameRect(_:to:)
func (w NSWindow) ConstrainFrameRectToScreen(frameRect corefoundation.CGRect, screen INSScreen) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](w.ID, objc.Sel("constrainFrameRect:toScreen:"), frameRect, screen)
	return corefoundation.CGRect(rv)
}

// Positions the window’s top-left to a given point.
//
// topLeftPoint: The new top-left point, in screen coordinates, for the window. When
// [NSZeroPoint], the window is not moved, except as needed to constrain to
// the visible screen
//
// # Return Value
//
// The point shifted from top left of the window in screen coordinates.
//
// # Discussion
//
// The returned point can be passed to a subsequent invocation of
// [CascadeTopLeftFromPoint] to position the next window so the title bars of
// both windows are fully visible.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/cascadeTopLeft(from:)
//
// [NSZeroPoint]: https://developer.apple.com/documentation/Foundation/NSZeroPoint
func (w NSWindow) CascadeTopLeftFromPoint(topLeftPoint corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](w.ID, objc.Sel("cascadeTopLeftFromPoint:"), topLeftPoint)
	return corefoundation.CGPoint(rv)
}

// Sets the origin and size of the window’s frame rectangle according to a
// given frame rectangle, thereby setting its position and size onscreen.
//
// frameRect: The frame rectangle for the window, including the title bar.
//
// flag: Specifies whether the window redraws the views that need to be displayed.
// When true the window sends a [DisplayIfNeeded] message down its view
// hierarchy, thus redrawing all views.
//
// # Discussion
//
// Note that the window server limits window position coordinates to ±16,000
// and sizes to 10,000.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/setFrame(_:display:)
func (w NSWindow) SetFrameDisplay(frameRect corefoundation.CGRect, flag bool) {
	objc.Send[objc.ID](w.ID, objc.Sel("setFrame:display:"), frameRect, flag)
}

// Sets the origin and size of the window’s frame rectangle, with optional
// animation, according to a given frame rectangle, thereby setting its
// position and size onscreen.
//
// frameRect: The frame rectangle for the window, including the title bar.
//
// displayFlag: Specifies whether the window redraws the views that need to be displayed.
// When true the window sends a [DisplayIfNeeded] message down its view
// hierarchy, thus redrawing all views.
//
// animateFlag: Specifies whether the window performs a smooth resize. true to perform the
// animation, whose duration is specified by [AnimationResizeTime].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/setFrame(_:display:animate:)
func (w NSWindow) SetFrameDisplayAnimate(frameRect corefoundation.CGRect, displayFlag bool, animateFlag bool) {
	objc.Send[objc.ID](w.ID, objc.Sel("setFrame:display:animate:"), frameRect, displayFlag, animateFlag)
}

// Specifies the duration of a smooth frame-size change.
//
// newFrame: The frame rectangle specified in [SetFrameDisplayAnimate].
//
// # Return Value
//
// The duration of the frame size change.
//
// # Discussion
//
// Subclasses can override this method to control the total time for the frame
// change.
//
// The [NSWindow] implementation uses the value from the [NSWindowResizeTime]
// user default as the time in seconds to resize by 150 pixels. If this value
// is unspecified, [NSWindowResizeTime] is 0.20 seconds (this default value
// may be different in different releases of macOS).
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/animationResizeTime(_:)
func (w NSWindow) AnimationResizeTime(newFrame corefoundation.CGRect) float64 {
	rv := objc.Send[float64](w.ID, objc.Sel("animationResizeTime:"), newFrame)
	return rv
}

// This action method simulates the user clicking the zoom box by momentarily
// highlighting the button and then zooming the window.
//
// sender: The object sending the message.
//
// # Discussion
//
// If the window doesn’t have a zoom box or can’t be zoomed for some
// reason, the computer beeps.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/performZoom(_:)
func (w NSWindow) PerformZoom(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("performZoom:"), sender)
}

// Toggles the size and location of the window between its standard state
// (which the application provides as the best size to display the window’s
// data) and its user state (a new size and location the user may have set by
// moving or resizing the window).
//
// sender: The object sending the message.
//
// # Discussion
//
// For more information on the standard and user states, see
// [WindowWillUseStandardFrameDefaultFrame].
//
// Typically, the system invokes the [Zoom] method after a user clicks the
// window’s zoom box, and [PerformZoom] may also invoke [Zoom]
// programmatically. It performs the following steps:
//
// - Invokes the [WindowWillUseStandardFrameDefaultFrame] method, if the
// delegate or the window class implements it, to obtain a “best fit”
// frame for the window. If neither the delegate nor the window class
// implements the method, [Zoom] uses a default frame. The default frame
// nearly fills the current screen that contains the largest part of the
// window’s current frame. - Adjusts the resulting frame, if necessary, to
// fit on the current screen. - Compares the resulting frame to the current
// frame to determine whether the window’s standard frame is currently
// displayed. If the current frame is within a few pixels of the standard
// frame in size and location, the system considers it a match. - Determines a
// new frame. If the window is currently in the standard state, the new frame
// represents the user state, saved during a previous zoom. If the window is
// currently in the user state, the new frame represents the standard state,
// computed in step 1 above. If there’s no saved user state because there
// has been no previous zoom, the size and location of the window don’t
// change. - Determines whether the window currently allows zooming. By
// default, zooming is allowed. If the window’s delegate implements the
// [WindowShouldZoomToFrame] method, [Zoom] invokes that method. If the
// delegate doesn’t implement the method but the window does, [Zoom] invokes
// the window’s version. [WindowShouldZoomToFrame] returns false if zooming
// isn’t currently allowed. - If the window currently allows zooming, sets
// the new frame.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/zoom(_:)
func (w NSWindow) Zoom(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("zoom:"), sender)
}

// Sets the size of the window’s content view to a given size, which is
// expressed in the window’s base coordinate system.
//
// size: The new size of the window’s content view in the window’s base
// coordinate system.
//
// # Discussion
//
// This size in turn alters the size of the [NSWindow] object itself. Note
// that the window server limits window sizes to 10,000; if necessary, be sure
// to limit `aSize` relative to the frame rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/setContentSize(_:)
func (w NSWindow) SetContentSize(size corefoundation.CGSize) {
	objc.Send[objc.ID](w.ID, objc.Sel("setContentSize:"), size)
}

// Removes the window from the screen list, which hides the window.
//
// sender: The window to remove.
//
// # Discussion
//
// If the window is the key or main window, the window object immediately
// behind it is made key or main in its place. Calling [OrderOut] causes the
// window to be removed from the screen, but does not cause it to be released.
// See the [Close] method for information on when a window is released.
// Calling [OrderOut] on a child window causes the window to be removed from
// its parent window before being removed.
//
// The default animation based on the window type will be used when the window
// is ordered out unless it has been modified by the [AnimationBehavior]
// property.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/orderOut(_:)
func (w NSWindow) OrderOut(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("orderOut:"), sender)
}

// Moves the window to the back of its level in the screen list, without
// changing either the key window or the main window.
//
// sender: Message originator.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/orderBack(_:)
func (w NSWindow) OrderBack(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("orderBack:"), sender)
}

// Moves the window to the front of its level in the screen list, without
// changing either the key window or the main window.
//
// sender: The message’s sender.
//
// # Discussion
//
// The default animation based on the window type will be used when the window
// is ordered front unless it has been modified by the [AnimationBehavior]
// property.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/orderFront(_:)
func (w NSWindow) OrderFront(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("orderFront:"), sender)
}

// Moves the window to the front of its level, even if its application isn’t
// active, without changing either the key window or the main window.
//
// # Discussion
//
// Normally an [NSWindow] object can’t be moved in front of the key window
// unless it and the key window are in the same application. You should rarely
// need to invoke this method; it’s designed to be used when applications
// are cooperating in such a way that an active application (with the key
// window) is using another application to display data.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/orderFrontRegardless()
func (w NSWindow) OrderFrontRegardless() {
	objc.Send[objc.ID](w.ID, objc.Sel("orderFrontRegardless"))
}

// Repositions the window’s window device in the window server’s screen
// list.
//
// place: - [NSWindowOut]: The window is removed from the screen list and `otherWin`
// is ignored.
//
// - [NSWindowAbove]: The window is ordered immediately in front of the window
// whose window number is `otherWin` - [NSWindowBelow]: The window is placed
// immediately behind the window represented by `otherWin`.
//
// otherWin: The number of the window the window is to be placed in front of or behind.
// Pass `0` to place the window in front of (when `place` is [NSWindowAbove])
// or behind (when `place` is [NSWindowBelow]) all other windows in its level.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/order(_:relativeTo:)
func (w NSWindow) OrderWindowRelativeTo(place NSWindowOrderingMode, otherWin int) {
	objc.Send[objc.ID](w.ID, objc.Sel("orderWindow:relativeTo:"), place, otherWin)
}

// Sets the window’s frame rectangle by reading the rectangle data stored
// under a given name from the defaults system.
//
// name: The name of the frame to read.
//
// # Return Value
//
// true when `name` is read and the frame is set successfully; otherwise,
// false.
//
// # Discussion
//
// The frame is constrained according to the window’s minimum and maximum
// size settings. This method causes a [WindowWillResizeToSize] message to be
// sent to the delegate.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/setFrameUsingName(_:)
func (w NSWindow) SetFrameUsingName(name NSWindowFrameAutosaveName) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("setFrameUsingName:"), objc.String(string(name)))
	return rv
}

// Sets the window’s frame rectangle by reading the rectangle data stored
// under a given name from the defaults system. Can operate on non-resizable
// windows.
//
// name: The name of the frame to read.
//
// force: true to use [SetFrameUsingName] on a non-resizable window; false to fail on
// a non-resizable window.
//
// # Return Value
//
// true when `name` is read and the frame is set successfully; otherwise,
// false.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/setFrameUsingName(_:force:)
func (w NSWindow) SetFrameUsingNameForce(name NSWindowFrameAutosaveName, force bool) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("setFrameUsingName:force:"), objc.String(string(name)), force)
	return rv
}

// Saves the window’s frame rectangle in the user defaults system under a
// given name.
//
// name: The name under which the frame is to be saved.
//
// # Discussion
//
// With the companion method [SetFrameUsingName], you can save and reset an
// [NSWindow] object’s frame over various launches of an application. The
// default is owned by the application and stored under the name “`NSWindow
// Frame name`”. See [UserDefaults] for more information.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/saveFrame(usingName:)
//
// [UserDefaults]: https://developer.apple.com/documentation/Foundation/UserDefaults
func (w NSWindow) SaveFrameUsingName(name NSWindowFrameAutosaveName) {
	objc.Send[objc.ID](w.ID, objc.Sel("saveFrameUsingName:"), objc.String(string(name)))
}

// Sets the window’s frame rectangle from a given string representation.
//
// string: A string representation of a frame rectangle, previously accessed using
// [StringWithSavedFrame].
//
// # Discussion
//
// If the window is not resizable, this method will not resize the window. The
// frame is constrained according to the window’s minimum and maximum size
// settings. This method causes a [WindowWillResizeToSize] message to be sent
// to the delegate.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/setFrame(from:)
func (w NSWindow) SetFrameFromString(string_ NSWindowPersistableFrameDescriptor) {
	objc.Send[objc.ID](w.ID, objc.Sel("setFrameFromString:"), objc.String(string(string_)))
}

// Makes the window the key window.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/makeKey()
func (w NSWindow) MakeKeyWindow() {
	objc.Send[objc.ID](w.ID, objc.Sel("makeKeyWindow"))
}

// Moves the window to the front of the screen list, within its level, and
// makes it the key window; that is, it shows the window.
//
// sender: The message’s sender.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/makeKeyAndOrderFront(_:)
func (w NSWindow) MakeKeyAndOrderFront(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("makeKeyAndOrderFront:"), sender)
}

// Informs the window that it has become the key window.
//
// # Discussion
//
// This method reestablishes the window’s first responder, sends the
// `becomeKeyWindow` message to that object if it responds, and posts
// [didBecomeKeyNotification] to the default notification center.
//
// Never invoke this method directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/becomeKey()
//
// [didBecomeKeyNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didBecomeKeyNotification
func (w NSWindow) BecomeKeyWindow() {
	objc.Send[objc.ID](w.ID, objc.Sel("becomeKeyWindow"))
}

// Resigns the window’s key window status.
//
// # Discussion
//
// This method sends [ResignKeyWindow] to the window’s first responder,
// sends [WindowDidResignKey] to the window’s delegate, and posts
// [didResignKeyNotification] to the default notification center.
//
// Never invoke this method directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/resignKey()
//
// [didResignKeyNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didResignKeyNotification
func (w NSWindow) ResignKeyWindow() {
	objc.Send[objc.ID](w.ID, objc.Sel("resignKeyWindow"))
}

// Makes the window the main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/makeMain()
func (w NSWindow) MakeMainWindow() {
	objc.Send[objc.ID](w.ID, objc.Sel("makeMainWindow"))
}

// Informs the window that it has become the main window.
//
// # Discussion
//
// This method posts an [didBecomeMainNotification] to the default
// notification center.
//
// Never invoke this method directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/becomeMain()
//
// [didBecomeMainNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didBecomeMainNotification
func (w NSWindow) BecomeMainWindow() {
	objc.Send[objc.ID](w.ID, objc.Sel("becomeMainWindow"))
}

// Resigns the window’s main window status.
//
// # Discussion
//
// This method sends [WindowDidResignMain] to the window’s delegate and
// posts [didResignMainNotification] to the default notification center.
//
// Never invoke this method directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/resignMain()
//
// [didResignMainNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didResignMainNotification
func (w NSWindow) ResignMainWindow() {
	objc.Send[objc.ID](w.ID, objc.Sel("resignMainWindow"))
}

// Toggles the visibility of the window’s toolbar.
//
// sender: The message’s sender.
//
// # Discussion
//
// See the [NSToolbar] class description for additional information.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/toggleToolbarShown(_:)
func (w NSWindow) ToggleToolbarShown(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("toggleToolbarShown:"), sender)
}

// Presents the toolbar customization user interface.
//
// sender: The message’s sender.
//
// # Discussion
//
// See the [NSToolbar] class description for additional information.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/runToolbarCustomizationPalette(_:)
func (w NSWindow) RunToolbarCustomizationPalette(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("runToolbarCustomizationPalette:"), sender)
}

// Adds a given window as a child window of the window.
//
// childWin: The child window to order.
//
// place: - [NSWindowAbove]: `childWin` is ordered immediately in front of the
// window.
//
// - [NSWindowBelow]: `childWin` is ordered immediately behind the window.
//
// # Discussion
//
// After the `childWin` is added as a child of the window, it is maintained in
// relative position indicated by `place` for subsequent ordering operations
// involving either window. While this attachment is active, moving `childWin`
// will not cause the window to move (as in sliding a drawer in or out), but
// moving the window will cause `childWin` to move.
//
// Note that you should not create cycles between parent and child windows.
// For example, you should not add window B as child of window A, then add
// window A as a child of window B.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/addChildWindow(_:ordered:)
func (w NSWindow) AddChildWindowOrdered(childWin INSWindow, place NSWindowOrderingMode) {
	objc.Send[objc.ID](w.ID, objc.Sel("addChildWindow:ordered:"), childWin, place)
}

// Detaches a given child window from the window.
//
// childWin: The child window to detach.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/removeChildWindow(_:)
func (w NSWindow) RemoveChildWindow(childWin INSWindow) {
	objc.Send[objc.ID](w.ID, objc.Sel("removeChildWindow:"), childWin)
}

// Reenables the default button cell’s key equivalent, so it performs a
// click when the user presses Return (or Enter).
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/enableKeyEquivalentForDefaultButtonCell()
func (w NSWindow) EnableKeyEquivalentForDefaultButtonCell() {
	objc.Send[objc.ID](w.ID, objc.Sel("enableKeyEquivalentForDefaultButtonCell"))
}

// Disables the default button cell’s key equivalent, so it doesn’t
// perform a click when the user presses Return (or Enter).
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/disableKeyEquivalentForDefaultButtonCell()
func (w NSWindow) DisableKeyEquivalentForDefaultButtonCell() {
	objc.Send[objc.ID](w.ID, objc.Sel("disableKeyEquivalentForDefaultButtonCell"))
}

// Returns the window’s field editor, creating it if requested.
//
// createFlag: If true, creates a field editor if one doesn’t exist; if false, does not
// create a field editor.
//
// A freshly created [NSWindow] object doesn’t have a field editor. After a
// field editor has been created for a window, the `createFlag` argument is
// ignored. By passing false for `createFlag` and testing the return value,
// however, you can predicate an action on the existence of the field editor.
//
// object: A text-displaying object for which the delegate (in
// [WindowWillReturnFieldEditorToObject]) assigns a custom field editor. Pass
// `nil` to get the default field editor, which can be the [NSWindow] field
// editor or a custom field editor returned by the delegate.
//
// # Return Value
//
// Returns the field editor for the designated object (`object`) or, if
// `object` is `nil`, the default field editor. Returns `nil` if `createFlag`
// is false and if the field editor doesn’t exist.
//
// # Discussion
//
// The field editor is a single [NSTextView] object that is shared among all
// the controls in a window for light text-editing needs. It is automatically
// instantiated when needed, and it can be used however your application sees
// fit. Typically, the field editor is used by simple text-bearing
// objects—for example, an [NSTextField] object uses its window’s field
// editor to display and manipulate text. The field editor can be shared by
// any number of objects, and so its state may be constantly changing.
// Therefore, it shouldn’t be used to display text that demands
// sophisticated layout (for this you should create a dedicated [NSTextView]
// object).
//
// The field editor may be in use by some view object, so be sure to properly
// dissociate it from that object before actually using it yourself (the
// appropriate way to do this is illustrated in the description of
// [EndEditingFor]). Once you retrieve the field editor, you can insert it in
// the view hierarchy, set a delegate to interpret text events, and have it
// perform whatever editing is needed. Then, when it sends a
// [TextDidEndEditing] message to the delegate, you can get its text to
// display or store and remove the field editor using [EndEditingFor].
//
// The window’s delegate can substitute a custom field editor in place of
// the window’s field editor by implementing
// [WindowWillReturnFieldEditorToObject]. The custom field editor can become
// the default editor (common to all text-displaying objects) or specific to a
// particular text-displaying object (`object`). The window sends this message
// to its delegate with itself and `object` as the arguments; if the delegate
// returns a non-`nil` value, the window returns that object instead of its
// field editor in [FieldEditorForObject]. However, note the following:
//
// - If the window’s delegate is identical to `object`,
// [WindowWillReturnFieldEditorToObject] isn’t sent to the delegate. - The
// object returned by the delegate method, though it may become first
// responder, does not become the window’s default field editor. Other
// objects continue to use the window’s default field editor.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/fieldEditor(_:for:)
func (w NSWindow) FieldEditorForObject(createFlag bool, object objectivec.IObject) INSText {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("fieldEditor:forObject:"), createFlag, object)
	return NSTextFromID(rv)
}

// Forces the field editor to give up its first responder status and prepares
// it for its next assignment.
//
// object: The object that is using the window’s field editor.
//
// # Discussion
//
// If the field editor is the first responder, it’s made to resign that
// status even if its [ResignFirstResponder] method returns false. This
// registration forces the field editor to send a [TextDidEndEditing] message
// to its delegate. The field editor is then removed from the view hierarchy,
// its delegate is set to `nil`, and it’s emptied of any text it may
// contain.
//
// This method is typically invoked by the object using the field editor when
// it’s finished. Other objects normally change the first responder by
// simply using [FirstResponder], which allows a field editor or other object
// to retain its first responder status if, for example, the user has entered
// an invalid value. The [EndEditingFor] method should be used only as a last
// resort if the field editor refuses to resign first responder status. Even
// in this case, you should always allow the field editor a chance to validate
// its text and take whatever other action it needs first. You can do this by
// first trying to make the [NSWindow] object the first responder:
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/endEditing(for:)
func (w NSWindow) EndEditingFor(object objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("endEditingFor:"), object)
}

// Reenables cursor rectangle management within the window after a
// [DisableCursorRects] message.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/enableCursorRects()
func (w NSWindow) EnableCursorRects() {
	objc.Send[objc.ID](w.ID, objc.Sel("enableCursorRects"))
}

// Disables all cursor rectangle management within the window.
//
// # Discussion
//
// Use this method when you need to do some special cursor manipulation and
// you don’t want the Application Kit interfering.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/disableCursorRects()
func (w NSWindow) DisableCursorRects() {
	objc.Send[objc.ID](w.ID, objc.Sel("disableCursorRects"))
}

// Invalidates all cursor rectangles in the window.
//
// # Discussion
//
// This method is invoked by [ResetCursorRects] to clear out existing cursor
// rectangles before resetting them. You shouldn’t invoke it in the code you
// write, but you might want to override it to change its behavior.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/discardCursorRects()
func (w NSWindow) DiscardCursorRects() {
	objc.Send[objc.ID](w.ID, objc.Sel("discardCursorRects"))
}

// Marks as invalid the cursor rectangles of a given view object in the
// window, so they’ll be set up again when the window becomes key.
//
// view: The view in the window’s view hierarchy.
//
// # Discussion
//
// If the window is current the key window, window resets the cursor
// rectangles immediately.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/invalidateCursorRects(for:)
func (w NSWindow) InvalidateCursorRectsForView(view INSView) {
	objc.Send[objc.ID](w.ID, objc.Sel("invalidateCursorRectsForView:"), view)
}

// Clears the window’s cursor rectangles and the cursor rectangles of the
// [NSView] objects in its view hierarchy.
//
// # Discussion
//
// Invokes [DiscardCursorRects] to clear the window’s cursor rectangles,
// then sends [ResetCursorRects] to every [NSView] object in the window’s
// view hierarchy.
//
// This method is typically invoked by the NSApplication object when it
// detects that the key window’s cursor rectangles are invalid. In program
// code, it’s more efficient to invoke [InvalidateCursorRectsForView].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/resetCursorRects()
func (w NSWindow) ResetCursorRects() {
	objc.Send[objc.ID](w.ID, objc.Sel("resetCursorRects"))
}

// Returns the window button of a given window button kind in the window’s
// view hierarchy.
//
// b: The type of standard window button to return.
//
// # Return Value
//
// Window button in the window’s view hierarchy of the type identified by
// `b`; `nil` when such button is not in the window’s view hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/standardWindowButton(_:)
func (w NSWindow) StandardWindowButton(b NSWindowButton) INSButton {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("standardWindowButton:"), b)
	return NSButtonFromID(rv)
}

// Adds the specified title bar accessory view controller to the window.
//
// childViewController: An instance of [NSTitlebarAccessoryViewController] containing the view to
// add, along with where to place it and how it should behave in full screen
// mode.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/addTitlebarAccessoryViewController(_:)
func (w NSWindow) AddTitlebarAccessoryViewController(childViewController INSTitlebarAccessoryViewController) {
	objc.Send[objc.ID](w.ID, objc.Sel("addTitlebarAccessoryViewController:"), childViewController)
}

// Inserts the view controller into the window’s array of title bar
// accessory view controllers at the specified index.
//
// childViewController: The title bar accessory view controller to insert.
//
// index: The index at which to insert `childViewController`.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/insertTitlebarAccessoryViewController(_:at:)
func (w NSWindow) InsertTitlebarAccessoryViewControllerAtIndex(childViewController INSTitlebarAccessoryViewController, index int) {
	objc.Send[objc.ID](w.ID, objc.Sel("insertTitlebarAccessoryViewController:atIndex:"), childViewController, index)
}

// Removes the view controller at the specified index from the window’s
// array of title bar accessory view controllers.
//
// index: The index in the array of title bar view controllers from which to remove
// the view controller.
//
// # Discussion
//
// You can also use [RemoveFromParentViewController] to remove a specific
// title bar accessory view controller.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/removeTitlebarAccessoryViewController(at:)
func (w NSWindow) RemoveTitlebarAccessoryViewControllerAtIndex(index int) {
	objc.Send[objc.ID](w.ID, objc.Sel("removeTitlebarAccessoryViewControllerAtIndex:"), index)
}

// Adds the provided window as a new tab in a tabbed window using the
// specified ordering instruction.
//
// window: The window to add as a tabbed window.
//
// ordered: A value that indicates the order of the added window relative to other
// windows.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/addTabbedWindow(_:ordered:)
func (w NSWindow) AddTabbedWindowOrdered(window INSWindow, ordered NSWindowOrderingMode) {
	objc.Send[objc.ID](w.ID, objc.Sel("addTabbedWindow:ordered:"), window, ordered)
}

// Merges all open windows into a single tabbed window.
//
// sender: The object that initiated the action to merge all windows.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/mergeAllWindows(_:)
func (w NSWindow) MergeAllWindows(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("mergeAllWindows:"), sender)
}

// Selects the next tab in the tab group in the trailing direction.
//
// sender: The object that initiated the action to select the next tab.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/selectNextTab(_:)
func (w NSWindow) SelectNextTab(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("selectNextTab:"), sender)
}

// Selects the previous tab in the tab group in the leading direction.
//
// sender: The object that initiated the action to select the previous tab.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/selectPreviousTab(_:)
func (w NSWindow) SelectPreviousTab(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("selectPreviousTab:"), sender)
}

// Moves the tab to a new containing window.
//
// sender: The object that initiated the action to move the tab to a new window.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/moveTabToNewWindow(_:)
func (w NSWindow) MoveTabToNewWindow(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("moveTabToNewWindow:"), sender)
}

// Shows or hides the tab bar.
//
// sender: The object that initiated the action to toggle the tab bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/toggleTabBar(_:)
func (w NSWindow) ToggleTabBar(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("toggleTabBar:"), sender)
}

// Shows or hides the tab overview.
//
// sender: The object that initiated the action to toggle the tab overview.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/toggleTabOverview(_:)
func (w NSWindow) ToggleTabOverview(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("toggleTabOverview:"), sender)
}

// Returns the next event matching a given mask.
//
// mask: The mask that the event to return must match. Events with non-matching
// masks are left in the queue. See [DiscardEventsMatchingMaskBeforeEvent] in
// [NSApplication] for the list of mask values.
//
// # Return Value
//
// The next event whose mask matches `mask`; `nil` when no matching event was
// found.
//
// # Discussion
//
// This method calls the [NextEventMatchingMaskUntilDateInModeDequeue] method,
// where the matching mask parameter is the specified `mask`, the `until`
// (Swift) or `untilDate` (Objective-C) parameter is [distantFuture], the
// `inMode` parameter is [NSEventTrackingRunLoopMode], and the `dequeue`
// parameter is true.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/nextEvent(matching:)
//
// [NSEventTrackingRunLoopMode]: https://developer.apple.com/documentation/AppKit/NSEventTrackingRunLoopMode
// [distantFuture]: https://developer.apple.com/documentation/Foundation/NSDate/distantFuture
func (w NSWindow) NextEventMatchingMask(mask NSEventMask) INSEvent {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("nextEventMatchingMask:"), mask)
	return NSEventFromID(rv)
}

// Forwards the message to the global application object.
//
// mask: The mask that the event to return must match.
//
// expiration: The date until which to wait for events.
//
// mode: The run loop mode to use while waiting for events
//
// deqFlag: true to remove the returned event from the event queue; false to leave the
// returned event in the queue.
//
// # Return Value
//
// The next event whose mask matches the specified `mask`; otherwise, `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/nextEvent(matching:until:inMode:dequeue:)
func (w NSWindow) NextEventMatchingMaskUntilDateInModeDequeue(mask NSEventMask, expiration foundation.NSDate, mode foundation.NSString, deqFlag bool) INSEvent {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("nextEventMatchingMask:untilDate:inMode:dequeue:"), mask, expiration, mode, deqFlag)
	return NSEventFromID(rv)
}

// Forwards the message to the global application object.
//
// mask: The mask of the events to discard.
//
// lastEvent: The event up to which queued events are discarded from the queue.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/discardEvents(matching:before:)
func (w NSWindow) DiscardEventsMatchingMaskBeforeEvent(mask NSEventMask, lastEvent INSEvent) {
	objc.Send[objc.ID](w.ID, objc.Sel("discardEventsMatchingMask:beforeEvent:"), mask, lastEvent)
}

// Forwards the message to the global application object.
//
// event: The event to add to the window’s event queue.
//
// flag: true to place the event in the front of the queue; false to place it in the
// back.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/postEvent(_:atStart:)
func (w NSWindow) PostEventAtStart(event INSEvent, flag bool) {
	objc.Send[objc.ID](w.ID, objc.Sel("postEvent:atStart:"), event, flag)
}

// This action method dispatches mouse and keyboard events the global
// application object sends to the window.
//
// event: The mouse or keyboard event to process.
//
// # Discussion
//
// Never invoke this method directly. A right mouse-down event in a window of
// an inactive application isn’t delivered to the corresponding [NSWindow]
// object. Instead, a [SendEvent] message with a window number of `0` delivers
// it to the NSApplication object.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/sendEvent(_:)
func (w NSWindow) SendEvent(event INSEvent) {
	objc.Send[objc.ID](w.ID, objc.Sel("sendEvent:"), event)
}

// Attempts to make a given responder the first responder for the window.
//
// responder: The responder to set as the window’s first responder. `nil` makes the
// window its first responder.
//
// # Return Value
//
// true when the operation is successful; otherwise, false.
//
// # Discussion
//
// If `responder` isn’t already the first responder, this method first sends
// a [ResignFirstResponder] message to the object that is the first responder.
// If that object refuses to resign, it remains the first responder, and this
// method immediately returns false. If the current first responder resigns,
// this method sends a [BecomeFirstResponder] message to `responder`. If
// `responder` does not accept first responder status, the [NSWindow] object
// becomes first responder; in this case, the method returns true even if
// `responder` refuses first responder status.
//
// If `responder` is `nil`, this method still sends [ResignFirstResponder] to
// the current first responder. If the current first responder refuses to
// resign, it remains the first responder and this method immediately returns
// false. If the current first responder returns true from
// [ResignFirstResponder], the window is made its own first responder and this
// method returns true.
//
// The Application Kit framework uses this method to alter the first responder
// in response to mouse-down events; you can also use it to explicitly set the
// first responder from within your program. The `responder` object is
// typically an [NSView] object in the window’s view hierarchy. If this
// method is called explicitly, first send [AcceptsFirstResponder] to
// `responder`, and do not call [FirstResponder] if [AcceptsFirstResponder]
// returns false.
//
// Use [InitialFirstResponder] to the set the first responder to be used when
// the window is brought onscreen for the first time.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/makeFirstResponder(_:)
func (w NSWindow) MakeFirstResponder(responder INSResponder) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("makeFirstResponder:"), responder)
	return rv
}

// Gives key view status to the view that precedes the given view.
//
// view: The view whose preceding view in the key view loop to seek.
//
// # Discussion
//
// Sends the [PreviousValidKeyView] message to `view` and, if that message
// returns an [NSView] object, invokes [FirstResponder] with the returned
// object.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/selectKeyView(preceding:)
func (w NSWindow) SelectKeyViewPrecedingView(view INSView) {
	objc.Send[objc.ID](w.ID, objc.Sel("selectKeyViewPrecedingView:"), view)
}

// Gives key view status to the view that follows the given view.
//
// view: The view whose following view in the key view loop to seek.
//
// # Discussion
//
// Sends the [NextValidKeyView] message to `view` and, if that message returns
// an [NSView] object, invokes [FirstResponder] with the returned object.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/selectKeyView(following:)
func (w NSWindow) SelectKeyViewFollowingView(view INSView) {
	objc.Send[objc.ID](w.ID, objc.Sel("selectKeyViewFollowingView:"), view)
}

// Searches for a candidate previous key view and, if it finds one, tries to
// make it the first responder.
//
// sender: The message’s sender.
//
// # Discussion
//
// The candidate is one of the following (which this function searches for in
// this order):
//
// - The current first responder’s previous valid key view, which the
// [PreviousValidKeyView] method of [NSView] returns - The
// [InitialFirstResponder] designates as the window’s initial first
// responder if it returns true to an [AcceptsFirstResponder] message -
// Otherwise, the initial first responder’s previous valid key view, which
// may be `nil`
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/selectPreviousKeyView(_:)
func (w NSWindow) SelectPreviousKeyView(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("selectPreviousKeyView:"), sender)
}

// Searches for a candidate next key view and, if it finds one, tries to make
// it the first responder.
//
// sender: The message’s sender.
//
// # Discussion
//
// The candidate is one of the following (which this function searches for in
// this order):
//
// - The current first responder’s next valid key view, which the
// [NextValidKeyView] method of [NSView] returns - The object
// [InitialFirstResponder] designates as the window’s initial first
// responder if it returns true to an [AcceptsFirstResponder] message -
// Otherwise, the initial first responder’s next valid key view, which may
// be `nil`
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/selectNextKeyView(_:)
func (w NSWindow) SelectNextKeyView(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("selectNextKeyView:"), sender)
}

// Marks the key view loop as “dirty” and in need of recalculation.
//
// # Discussion
//
// The key view loop is recalculated the next time someone requests the next
// or previous key view of the window. The recalculated loop is based on the
// geometric order of the views in the window.
//
// If you don’t want to maintain the key view loop of your window manually,
// you can use this method to do it for you. When it’s first loaded,
// [NSWindow] calls this method automatically if your window doesn’t have a
// key view loop already established. If you add or remove views later, you
// can call this method manually to update the window’s key view loop. You
// can also set the [AutorecalculatesKeyViewLoop] property to have the window
// recalculate the loop automatically.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/recalculateKeyViewLoop()
func (w NSWindow) RecalculateKeyViewLoop() {
	objc.Send[objc.ID](w.ID, objc.Sel("recalculateKeyViewLoop"))
}

// # Discussion
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/transferWindowSharing(to:completionHandler:)
func (w NSWindow) TransferWindowSharingToWindowCompletionHandler(window INSWindow, completionHandler ErrorHandler) {
	_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("transferWindowSharingToWindow:completionHandler:"), window, _block1)
}

// Tracks events that match the specified mask using the specified tracking
// handler until the tracking handler explicitly terminates tracking.
//
// mask: The event mask (see [NSEventMask] in [NSEvent] for possible values).
//
// timeout: The maximum time interval the system waits for an event before passing
// `nil` to the handler.
//
// mode: The run loop mode.
//
// trackingHandler: A block that is called to track the events. The block takes the following
// parameters:
//
// event: The event to examine. stop: A Boolean value that indicates when
// tracking should stop.
//
// # Discussion
//
// You can use this method in a tracking loop to get pressure events when you
// add [NSEventMaskPressure] to the event mask. This method returns when
// tracking terminates.
//
// Each event is removed from the event queue and then passed to the tracking
// handler. If a matching event does not exist in the event queue, the main
// thread blocks in the specified runloop mode until an event of the requested
// type is received or the specified timeout expires. If the timeout expires,
// the tracking handler is called with a `nil` event (a negative timeout is
// interpreted as `0`). Use [NSEventDurationForever] to prevent timing out.
// Tracking continues until you set `stop` to true. Note that calls to
// [NextEventMatchingMask] are allowed inside the `trackingHandler` block.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/trackEvents(matching:timeout:mode:handler:)
func (w NSWindow) TrackEventsMatchingMaskTimeoutModeHandler(mask NSEventMask, timeout float64, mode foundation.NSString, trackingHandler EventHandler) {
	_block3, _ := NewEventBlock(trackingHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("trackEventsMatchingMask:timeout:mode:handler:"), mask, timeout, mode, _block3)
}

// Starts a window drag based on the specified mouse-down event.
//
// event: The original mouse-down event received by the application or a view.
//
// # Discussion
//
// Your application (or a view) can call this method after receiving and
// examining a mouse-down event. Upon examination of the event, a view may
// allow that portion of the window to start a window drag and can hand off
// the work to the Window Server process by calling this method. Doing so
// allows the window to participate in space switching and other system
// features.
//
// This method returns right away, and a mouse-up event may not get sent.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/performDrag(with:)
func (w NSWindow) PerformWindowDragWithEvent(event INSEvent) {
	objc.Send[objc.ID](w.ID, objc.Sel("performWindowDragWithEvent:"), event)
}

// Disables snapshot restoration.
//
// # Discussion
//
// After disabling snapshot restoration, the system doesn’t snapshot the
// window’s restorable state.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/disableSnapshotRestoration()
func (w NSWindow) DisableSnapshotRestoration() {
	objc.Send[objc.ID](w.ID, objc.Sel("disableSnapshotRestoration"))
}

// Enables snapshot restoration.
//
// # Discussion
//
// While snapshot restoration is enabled, the system snapshots the window’s
// restorable state.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/enableSnapshotRestoration()
func (w NSWindow) EnableSnapshotRestoration() {
	objc.Send[objc.ID](w.ID, objc.Sel("enableSnapshotRestoration"))
}

// Passes a display message down the window’s view hierarchy, thus redrawing
// all views within the window.
//
// # Discussion
//
// You rarely need to invoke this method. [NSWindow] objects normally record
// which of their views need displaying and display them automatically on each
// pass through the event loop.
//
// This method includes the frame view that draws the border, title bar, and
// other peripheral elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/display()
func (w NSWindow) Display() {
	objc.Send[objc.ID](w.ID, objc.Sel("display"))
}

// Passes a display message down the window’s view hierarchy, thus redrawing
// all views that need displaying.
//
// # Discussion
//
// This method includes the frame view that draws the border, title bar, and
// other peripheral elements. It’s useful when you want to modify some
// number of views and then display only the ones that you modified.
//
// You rarely need to invoke this method. [NSWindow] objects normally record
// which of their views need displaying and display them automatically on each
// pass through the event loop.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/displayIfNeeded()
func (w NSWindow) DisplayIfNeeded() {
	objc.Send[objc.ID](w.ID, objc.Sel("displayIfNeeded"))
}

// Updates the window.
//
// # Discussion
//
// The [NSWindow] implementation of this method does nothing more than post an
// [didUpdateNotification] notification to the default notification center. A
// subclass can override this method to perform specialized operations, but it
// should send an update message to `super` just before returning. For
// example, the [NSMenu] class implements this method to disable and enable
// menu commands.
//
// An [NSWindow] object is automatically sent an `update` message on every
// pass through the event loop and before it’s displayed onscreen. You can
// manually cause an `update` message to be sent to all visible [NSWindow]
// objects through the [NSApplication] [UpdateWindows] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/update()
//
// [didUpdateNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didUpdateNotification
func (w NSWindow) Update() {
	objc.Send[objc.ID](w.ID, objc.Sel("update"))
}

// Registers a set of pasteboard types that the window accepts as the
// destination of an image-dragging session.
//
// newTypes: An array of the pasteboard types the window accepts as the destination of
// an image-dragging session.
//
// # Discussion
//
// Registering an [NSWindow] object for dragged types automatically makes it a
// candidate destination object for a dragging session. [NSWindow] has a
// default implementation for many of the methods in the
// [NSDraggingDestination] protocol. The default implementation forwards each
// message to the delegate if the delegate responds to the selector of the
// message. The messages forwarded this way are [DraggingEntered],
// [DraggingUpdated], [DraggingExited], [PrepareForDragOperation],
// [PerformDragOperation], and [ConcludeDragOperation].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/registerForDraggedTypes(_:)
func (w NSWindow) RegisterForDraggedTypes(newTypes []string) {
	objc.Send[objc.ID](w.ID, objc.Sel("registerForDraggedTypes:"), objectivec.StringSliceToNSArray(newTypes))
}

// Unregisters the window as a possible destination for dragging operations.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/unregisterDraggedTypes()
func (w NSWindow) UnregisterDraggedTypes() {
	objc.Send[objc.ID](w.ID, objc.Sel("unregisterDraggedTypes"))
}

// Returns a backing store pixel-aligned rectangle in window coordinates.
//
// rect: The rectangle in view coordinates.
//
// options: The alignment options. [AlignmentOptions] specifies the possible values.
//
// # Return Value
//
// A rectangle, in window coordinates, aligned to the backing store pixels
// according to the specified options.
//
// # Discussion
//
// This method uses [NSIntegralRectWithOptions(_:_:)] to align the input
// rectangle, and produces a backing store pixel-aligned rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/backingAlignedRect(_:options:)
//
// [AlignmentOptions]: https://developer.apple.com/documentation/Foundation/AlignmentOptions
// [NSIntegralRectWithOptions(_:_:)]: https://developer.apple.com/documentation/Foundation/NSIntegralRectWithOptions(_:_:)
func (w NSWindow) BackingAlignedRectOptions(rect corefoundation.CGRect, options foundation.AlignmentOptions) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](w.ID, objc.Sel("backingAlignedRect:options:"), rect, options)
	return corefoundation.CGRect(rv)
}

// Converts a rectangle from its pixel-aligned backing store coordinate system
// to the window’s coordinate system.
//
// rect: The rectangle aligned to the pixel backing store coordinate system.
//
// # Return Value
//
// A rectangle in the window’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/convertFromBacking(_:)
func (w NSWindow) ConvertRectFromBacking(rect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](w.ID, objc.Sel("convertRectFromBacking:"), rect)
	return corefoundation.CGRect(rv)
}

// Converts a rectangle from the screen coordinate system to the window’s
// coordinate system.
//
// rect: A rectangle in the screen’s coordinate system.
//
// # Return Value
//
// A rectangle in the window’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/convertFromScreen(_:)
func (w NSWindow) ConvertRectFromScreen(rect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](w.ID, objc.Sel("convertRectFromScreen:"), rect)
	return corefoundation.CGRect(rv)
}

// Converts a point from its pixel-aligned backing store coordinate system to
// the window’s coordinate system.
//
// point: The point in the pixel-aligned backing store coordinate system.
//
// # Return Value
//
// A point in the window’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/convertPointFromBacking(_:)
func (w NSWindow) ConvertPointFromBacking(point corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](w.ID, objc.Sel("convertPointFromBacking:"), point)
	return corefoundation.CGPoint(rv)
}

// Converts a point from the screen coordinate system to the window’s
// coordinate system.
//
// point: A point in the screen’s coordinate system.
//
// # Return Value
//
// A point in the window’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/convertPoint(fromScreen:)
func (w NSWindow) ConvertPointFromScreen(point corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](w.ID, objc.Sel("convertPointFromScreen:"), point)
	return corefoundation.CGPoint(rv)
}

// Converts a rectangle from the window’s coordinate system to its
// pixel-aligned backing store coordinate system.
//
// rect: A rectangle in the window’s coordinate system.
//
// # Return Value
//
// A rectangle in its pixel-aligned backing store coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/convertToBacking(_:)
func (w NSWindow) ConvertRectToBacking(rect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](w.ID, objc.Sel("convertRectToBacking:"), rect)
	return corefoundation.CGRect(rv)
}

// Converts a rectangle to the screen coordinate system from the window’s
// coordinate system.
//
// rect: A rectangle in the window’s coordinate system.
//
// # Return Value
//
// A rectangle in the screen’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/convertToScreen(_:)
func (w NSWindow) ConvertRectToScreen(rect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](w.ID, objc.Sel("convertRectToScreen:"), rect)
	return corefoundation.CGRect(rv)
}

// Converts a point from the window’s coordinate system to its pixel-aligned
// backing store coordinate system.
//
// point: A point in the window’s coordinate system.
//
// # Return Value
//
// A point in its pixel-aligned backing store coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/convertPointToBacking(_:)
func (w NSWindow) ConvertPointToBacking(point corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](w.ID, objc.Sel("convertPointToBacking:"), point)
	return corefoundation.CGPoint(rv)
}

// Converts a point to the screen coordinate system from the window’s
// coordinate system.
//
// point: A point in the window’s coordinate system.
//
// # Return Value
//
// A point in the screen’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/convertPoint(toScreen:)
func (w NSWindow) ConvertPointToScreen(point corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](w.ID, objc.Sel("convertPointToScreen:"), point)
	return corefoundation.CGPoint(rv)
}

// Sets a given path as the window’s title, formatting it as a file-system
// path, and records this path as the window’s associated file.
//
// filename: The file path to set as the window’s title.
//
// # Discussion
//
// The windows’ title bar displays the filename, not the file’s path.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/setTitleWithRepresentedFilename(_:)
func (w NSWindow) SetTitleWithRepresentedFilename(filename string) {
	objc.Send[objc.ID](w.ID, objc.Sel("setTitleWithRepresentedFilename:"), objc.String(filename))
}

// Sets the window’s location to the center of the screen.
//
// # Discussion
//
// The window is placed exactly in the center horizontally and somewhat above
// center vertically. Such a placement carries a certain visual immediacy and
// importance. This method doesn’t put the window onscreen, however; use
// [KeyAndOrderFront] to do that.
//
// You typically use this method to place a window—most likely an alert
// dialog—where the user can’t miss it. This method is invoked
// automatically when a panel is placed on the screen by the
// [RunModalForWindow] method of the [NSApplication] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/center()
func (w NSWindow) Center() {
	objc.Send[objc.ID](w.ID, objc.Sel("center"))
}

// Simulates the user clicking the close button by momentarily highlighting
// the button and then closing the window.
//
// sender: The message’s sender.
//
// # Discussion
//
// If the window’s delegate or the window itself implements
// [WindowShouldClose], the window sends that message with the window as the
// argument. The window sends only one such message; if both the delegate and
// the window implement the method, the delegate receives the message. If the
// [WindowShouldClose] method returns false, the window doesn’t close. If
// neither the window nor the delegate implement [WindowShouldClose], or it
// returns true, this method invokes [Close] to close the window.
//
// If the window doesn’t have a close button or can’t close (for example,
// if the delegate replies false to a [WindowShouldClose] message), the system
// emits the alert sound.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/performClose(_:)
func (w NSWindow) PerformClose(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("performClose:"), sender)
}

// Removes the window from the screen.
//
// # Discussion
//
// If the window is set to be released when closed, a `release` message is
// sent to the object after the current event is completed. For an [NSWindow]
// object, the default is to be released on closing, while for an [NSPanel]
// object, the default is not to be released. You can use the
// [ReleasedWhenClosed] property to change the default behavior.
//
// A window doesn’t have to be visible to receive the close message. For
// example, when the application terminates, it sends the close message to all
// windows in its window list, even those that are not currently visible.
//
// The close method posts a [willCloseNotification] notification to the
// default notification center.
//
// The close method differs in two important ways from the [PerformClose]
// method:
//
// - It does not attempt to send a [WindowShouldClose] message to the window
// or its delegate. - It does not simulate the user clicking the close button
// by momentarily highlighting the button.
//
// Use [PerformClose] if you need these features.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/close()
//
// [willCloseNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/willCloseNotification
func (w NSWindow) Close() {
	objc.Send[objc.ID](w.ID, objc.Sel("close"))
}

// Simulates the user clicking the minimize button by momentarily highlighting
// the button, then minimizing the window.
//
// sender: The message’s sender.
//
// # Discussion
//
// If the window doesn’t have a minimize button or can’t be minimized for
// some reason, the system emits the alert sound.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/performMiniaturize(_:)
func (w NSWindow) PerformMiniaturize(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("performMiniaturize:"), sender)
}

// Removes the window from the screen list and displays the minimized window
// in the Dock.
//
// sender: The message’s sender.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/miniaturize(_:)
func (w NSWindow) Miniaturize(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("miniaturize:"), sender)
}

// De-minimizes the window.
//
// sender: The message’s sender.
//
// # Discussion
//
// Invoke this method to programmatically deminimize a minimized window in the
// Dock.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/deminiaturize(_:)
func (w NSWindow) Deminiaturize(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("deminiaturize:"), sender)
}

// Runs the Print panel, and if the user chooses an option other than
// canceling, prints the window (its frame view and all subviews).
//
// sender: The message’s sender.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/printWindow(_:)
func (w NSWindow) Print(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("print:"), sender)
}

// Returns EPS data that draws the region of the window within a given
// rectangle.
//
// rect: A rectangle (expressed in the window’s coordinate system) that identifies
// the region to be expressed as EPS data.
//
// # Return Value
//
// The region in the window (identified by `rect`) as EPS data.
//
// # Discussion
//
// This data can be placed on a pasteboard, written to a file, or used to
// create an [NSImage] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/dataWithEPS(inside:)
func (w NSWindow) DataWithEPSInsideRect(rect corefoundation.CGRect) foundation.NSData {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("dataWithEPSInsideRect:"), rect)
	return foundation.NSDataFromID(rv)
}

// Returns PDF data that draws the region of the window within a given
// rectangle.
//
// rect: A rectangle (expressed in the window’s coordinate system) that identifies
// the region to be expressed as PDF data.
//
// # Return Value
//
// The region in the window (identified by `rect`) as PDF data.
//
// # Discussion
//
// This data can be placed on a pasteboard, written to a file, or used to
// create an [NSImage] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/dataWithPDF(inside:)
func (w NSWindow) DataWithPDFInsideRect(rect corefoundation.CGRect) foundation.NSData {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("dataWithPDFInsideRect:"), rect)
	return foundation.NSDataFromID(rv)
}

// Updates the constraints based on changes to views in the window since the
// last layout.
//
// # Discussion
//
// When a new layout pass is triggered for a window, the system invokes this
// method to ensure that any constraints for views in the window are updated
// with information from the current view hierarchy and its constraints. This
// method is called automatically by the system, but may be invoked manually
// if you need to examine the most up to date constraints.
//
// Subclasses should not override this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/updateConstraintsIfNeeded()
func (w NSWindow) UpdateConstraintsIfNeeded() {
	objc.Send[objc.ID](w.ID, objc.Sel("updateConstraintsIfNeeded"))
}

// Updates the layout of views in the window based on the current views and
// constraints.
//
// # Discussion
//
// Before displaying a window that uses constraints-based layout the system
// invokes this method to ensure that the layout of all views is up to date.
// This method updates the layout if needed, first invoking
// [UpdateConstraintsIfNeeded] to ensure that all constraints are up to date.
// This method is called automatically by the system, but may be invoked
// manually if you need to examine the most up to date layout.
//
// Subclasses should not override this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/layoutIfNeeded()
func (w NSWindow) LayoutIfNeeded() {
	objc.Send[objc.ID](w.ID, objc.Sel("layoutIfNeeded"))
}

// Displays a visual representation of the supplied constraints in the window.
//
// constraints: The constraints to visualize. All constraints must be held by views in the
// window.
//
// # Discussion
//
// The constraints to visualize are typically discovered by identifying a view
// whose layout is unexpected and then calling
// [ConstraintsAffectingLayoutForOrientation] on that view.
//
// This method should only be used for debugging constraint-based layout. No
// application should ship with calls to this method as part of its operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/visualizeConstraints(_:)
func (w NSWindow) VisualizeConstraints(constraints []NSLayoutConstraint) {
	objc.Send[objc.ID](w.ID, objc.Sel("visualizeConstraints:"), objectivec.IObjectSliceToNSArray(constraints))
}

// Returns the part of the window that stays stationary during
// constraint-based layout.
//
// orientation: The attribute for orientation. [NSLayoutConstraint.Orientation]specifies
// the possible values.
//
// # Return Value
//
// Returns the layout attribute.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/anchorAttribute(for:)
//
// [NSLayoutConstraint.Orientation]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Orientation
func (w NSWindow) AnchorAttributeForOrientation(orientation NSLayoutConstraintOrientation) NSLayoutAttribute {
	rv := objc.Send[NSLayoutAttribute](w.ID, objc.Sel("anchorAttributeForOrientation:"), orientation)
	return NSLayoutAttribute(rv)
}

// Sets the part of the window that stays stationary during constraint-based
// layout.
//
// attr: The layout attribute. [NSLayoutConstraint.Attribute] specifies the possible
// values.
//
// orientation: The window drag orientation. [NSLayoutConstraint.Orientation] specifies the
// possible values.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/setAnchorAttribute(_:for:)
//
// [NSLayoutConstraint.Attribute]: https://developer.apple.com/documentation/UIKit/NSLayoutConstraint/Attribute
// [NSLayoutConstraint.Orientation]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Orientation
func (w NSWindow) SetAnchorAttributeForOrientation(attr NSLayoutAttribute, orientation NSLayoutConstraintOrientation) {
	objc.Send[objc.ID](w.ID, objc.Sel("setAnchorAttribute:forOrientation:"), attr, orientation)
}

// A Boolean value that indicates if the window and its screen use a color
// space that can represent the specified display gamut.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/canRepresent(_:)
func (w NSWindow) CanRepresentDisplayGamut(displayGamut NSDisplayGamut) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("canRepresentDisplayGamut:"), displayGamut)
	return rv
}

// Sets the window’s miniaturized state to the value you specify.
//
// # Discussion
//
// Depending on the current miniaturized state and the value of `flag`, the
// window may minimize to the Dock or expand from the Dock.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/setIsMiniaturized(_:)
func (w NSWindow) SetIsMiniaturized(flag bool) {
	objc.Send[objc.ID](w.ID, objc.Sel("setIsMiniaturized:"), flag)
}

// Sets the window’s visible state to the value you specify.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/setIsVisible(_:)
func (w NSWindow) SetIsVisible(flag bool) {
	objc.Send[objc.ID](w.ID, objc.Sel("setIsVisible:"), flag)
}

// Sets the window’s zoomed state to the value you specify.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/setIsZoomed(_:)
func (w NSWindow) SetIsZoomed(flag bool) {
	objc.Send[objc.ID](w.ID, objc.Sel("setIsZoomed:"), flag)
}

// Handles the AppleScript command to close the window (and its associated
// document, if any).
//
// # Discussion
//
// Extracts `close` command arguments from the `command` object and uses them
// to determine how to close the associated document—specifically, whether
// to ignore unsaved changes, save changes automatically, or ask the
// user—and identifies the file to save the document to. By default, the
// window saves the document to the file that was opened or previously saved
// to. Otherwise, the window saves it with an “untitled” name.
//
// If there’s a corresponding document and the window is the main window of
// the document, this function forwards the `close` command to the
// corresponding document; otherwise, the window sends itself a `performClose`
// message, if it has a close box.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/handleClose(_:)
func (w NSWindow) HandleCloseScriptCommand(command foundation.NSCloseCommand) objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("handleCloseScriptCommand:"), command)
	return objectivec.Object{ID: rv}
}

// Handles the AppleScript command to print the contents of the window (or its
// associated document, if any).
//
// # Discussion
//
// If there’s a corresponding document and the window is the main window of
// the document, it forwards the `print` command to the corresponding
// document. Otherwise, the window sends itself a `print` message.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/handlePrint(_:)
func (w NSWindow) HandlePrintScriptCommand(command foundation.NSScriptCommand) objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("handlePrintScriptCommand:"), command)
	return objectivec.Object{ID: rv}
}

// Handles the AppleScript command to save the window (and its associated
// document, if any).
//
// # Discussion
//
// If there’s a corresponding document and the window is the main window of
// the document, it forwards the `save` command to the corresponding document.
// Otherwise, this method does nothing.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/handleSave(_:)
func (w NSWindow) HandleSaveScriptCommand(command foundation.NSScriptCommand) objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("handleSaveScriptCommand:"), command)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppKit/NSWindow/beginDraggingSession(items:event:source:)
func (w NSWindow) BeginDraggingSessionWithItemsEventSource(items []NSDraggingItem, event INSEvent, source NSDraggingSource) INSDraggingSession {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("beginDraggingSessionWithItems:event:source:"), objectivec.IObjectSliceToNSArray(items), event, source)
	return NSDraggingSessionFromID(rv)
}

// window: The window to share
//
// completionHandler: A completion block that is called after the request finishes.
//
// # Discussion
//
// Request sharing of window. If there is an available ScreenCaptureKit
// sharing session, an alert will be presented asking the user to confirm the
// share
//
// The error will be non-nil if the request does not result in a window being
// shared. The error will be NSUserCancelledError if there is no
// ScreenCaptureKit session, or if the user rejects the offer to share. If
// sharing fails for some other reason, the error will provide the details.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/requestSharingOfWindow(_:completionHandler:)
func (w NSWindow) RequestSharingOfWindowCompletionHandler(window INSWindow, completionHandler ErrorHandler) {
	_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("requestSharingOfWindow:completionHandler:"), window, _block1)
}

// image: An image showing a preview of the window to share
//
// title: The title to show in a confirmation dialog
//
// completionHandler: A completion block that is called after the request finishes.
//
// # Discussion
//
// Request sharing of window to be provided later. If there is an available
// ScreenCaptureKit sharing session, an alert will be presented asking the
// user to confirm the share. The delegate will be asked to provide the window
// to share via windowForSharingRequestFromWindow:
//
// The error will be non-nil if the request does not result in a window being
// shared. The error will be NSUserCancelledError if there is no
// ScreenCaptureKit session, or if the user rejects the offer to share. If
// sharing fails for some other reason, the error will provide the details.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/requestSharingOfWindow(usingPreview:title:completionHandler:)
func (w NSWindow) RequestSharingOfWindowUsingPreviewTitleCompletionHandler(image INSImage, title string, completionHandler ErrorHandler) {
	_block2, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("requestSharingOfWindowUsingPreview:title:completionHandler:"), image, objc.String(title), _block2)
}

// Returns the animation that should be performed for the specified key.
//
// key: The action name or property specified as a string.
//
// # Return Value
//
// The animation to perform. A subclass of [CAAnimation].
//
// # Discussion
//
// When the action specified by `key` is triggered for an object, this method
// is consulted to find the animation, if any, that should be performed in
// response.
//
// Like its Core Animation [CALayer] counterpart, [action(forKey:)], this
// method is a funnel point that defines the order in which the search for an
// animation proceeds.It first checks the receiver’s Getting the Animator
// Proxy dictionary for a value matching `key`, then falls back to [Animator]
// for the receiver’s class.
//
// Subclasses should not typically need to override this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/animation(forKey:)
//
// [CAAnimation]: https://developer.apple.com/documentation/QuartzCore/CAAnimation
// [CALayer]: https://developer.apple.com/documentation/QuartzCore/CALayer
// [action(forKey:)]: https://developer.apple.com/documentation/QuartzCore/CALayer/action(forKey:)
func (w NSWindow) AnimationForKey(key NSAnimatablePropertyKey) objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("animationForKey:"), objc.String(string(key)))
	return objectivec.Object{ID: rv}
}

// Sets the option dictionary that maps event trigger keys to animation
// objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/animations
func (w NSWindow) Animations() foundation.INSDictionary {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("animations"))
	return foundation.NSDictionaryFromID(rv)
}

// Returns a proxy object for the receiver that can be used to initiate
// implied animation for property changes.
//
// # Return Value
//
// Returns a proxy object for the receiver that can initiate implied
// animations in response to property changes.
//
// # Discussion
//
// The animator proxy object should be treated as if it was the receiver
// itself, and may be passed to any code that accepts the receiver as a
// parameter.
//
// Sending key-value coding compliant “set” messages to the proxy will
// trigger animation for automatically animated properties of its target
// object, if the active [NSAnimationContext] in the current thread has a
// duration value greater than zero, and an animation for the property key is
// found by the [NSAnimatablePropertyContainer] search mechanism.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/animator()
func (w NSWindow) Animator() INSWindow {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("animator"))
	return NSWindowFromID(rv)
}

// The appearance of the receiver, in an [NSAppearance] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSAppearanceCustomization/appearance
func (w NSWindow) Appearance() INSAppearance {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("appearance"))
	return NSAppearanceFromID(rv)
}

// The appearance that will be used when the receiver is drawn onscreen, in an
// [NSAppearance] object. (read-only)
//
// See: https://developer.apple.com/documentation/AppKit/NSAppearanceCustomization/effectiveAppearance
func (w NSWindow) EffectiveAppearance() INSAppearance {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("effectiveAppearance"))
	return NSAppearanceFromID(rv)
}

// A string that identifies the user interface item.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceItemIdentification/identifier
func (w NSWindow) Identifier() NSUserInterfaceItemIdentifier {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("identifier"))
	return NSUserInterfaceItemIdentifier(foundation.NSStringFromID(rv).String())
}

// Implemented to override the default action of enabling or disabling a
// specific menu item.
//
// menuItem: An [NSMenuItem] object that represents the menu item.
//
// # Return Value
//
// true to enable `menuItem`, false to disable it.
//
// # Discussion
//
// The object implementing this method must be the target of `menuItem`. You
// can determine which menu item `menuItem` is by querying it for its tag or
// action.
//
// The following example disables the menu item associated with the
// `nextRecord` action method when the selected line in a table view is the
// last one; conversely, it disables the menu item with `priorRecord` as its
// action method when the selected row is the first one in the table view.
// (The `countryOrRegionKeys` array contains names that appear in the table
// view.)
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemValidation/validateMenuItem(_:)
func (w NSWindow) ValidateMenuItem(menuItem INSMenuItem) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("validateMenuItem:"), menuItem)
	return rv
}

// Returns a Boolean value that indicates whether the sender should be
// enabled.
//
// item: The user interface item to validate. You can send `anItem` the [Action] and
// [Tag] messages.
//
// # Return Value
//
// true if the user interface item should be enabled, otherwise false.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceValidations/validateUserInterfaceItem(_:)
func (w NSWindow) ValidateUserInterfaceItem(item NSValidatedUserInterfaceItem) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("validateUserInterfaceItem:"), item)
	return rv
}

// Returns the window numbers for all visible windows satisfying the specified
// options.
//
// options: The possible options are specified in [NSWindow.NumberListOptions].
//
// # Return Value
//
// An array of window numbers for all visible windows satisfying the specified
// options. (Windows on the active space are returned in z-order; that is,
// front to back.)
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/windowNumbers(options:)
//
// [NSWindow.NumberListOptions]: https://developer.apple.com/documentation/AppKit/NSWindow/NumberListOptions
func (_NSWindowClass NSWindowClass) WindowNumbersWithOptions(options NSWindowNumberListOptions) []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](objc.ID(_NSWindowClass.class), objc.Sel("windowNumbersWithOptions:"), options)
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// Returns the content rectangle used by a window with a given frame rectangle
// and window style.
//
// fRect: The frame rectangle for the window expressed in screen coordinates.
//
// style: The window style for the window. See [NSWindow.StyleMask] for a list of
// style mask values.
//
// # Return Value
//
// The content rectangle, expressed in screen coordinates, used by the window
// with `fRect` and `style`.
//
// # Discussion
//
// When a [NSWindow] instance is available, you should use
// [ContentRectForFrameRect] instead of this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/contentRect(forFrameRect:styleMask:)
//
// [NSWindow.StyleMask]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
func (_NSWindowClass NSWindowClass) ContentRectForFrameRectStyleMask(fRect corefoundation.CGRect, style NSWindowStyleMask) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](objc.ID(_NSWindowClass.class), objc.Sel("contentRectForFrameRect:styleMask:"), fRect, style)
	return corefoundation.CGRect(rv)
}

// Returns the frame rectangle used by a window with a given content rectangle
// and window style.
//
// cRect: The content rectangle for a window expressed in screen coordinates.
//
// style: The window style for the window. See [NSWindow.StyleMask] for a list of
// style mask values.
//
// # Return Value
//
// The frame rectangle, expressed in screen coordinates, used by the window
// with `cRect` and `style`.
//
// # Discussion
//
// When a [NSWindow] instance is available, you should use
// [FrameRectForContentRect] instead of this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/frameRect(forContentRect:styleMask:)
//
// [NSWindow.StyleMask]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
func (_NSWindowClass NSWindowClass) FrameRectForContentRectStyleMask(cRect corefoundation.CGRect, style NSWindowStyleMask) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](objc.ID(_NSWindowClass.class), objc.Sel("frameRectForContentRect:styleMask:"), cRect, style)
	return corefoundation.CGRect(rv)
}

// Returns the minimum width a window’s frame rectangle must have for it to
// display a title, with a given window style.
//
// title: The title for the window.
//
// style: The window style for the window. See [NSWindow.StyleMask] for a list of
// style mask values.
//
// # Return Value
//
// The minimum width of the window’s frame, using `style`, in order to
// display `title`.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/minFrameWidth(withTitle:styleMask:)
//
// [NSWindow.StyleMask]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
func (_NSWindowClass NSWindowClass) MinFrameWidthWithTitleStyleMask(title string, style NSWindowStyleMask) float64 {
	rv := objc.Send[float64](objc.ID(_NSWindowClass.class), objc.Sel("minFrameWidthWithTitle:styleMask:"), objc.String(title), style)
	return rv
}

// Removes the frame data stored under a given name from the application’s
// user defaults.
//
// name: The name of the frame to remove.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/removeFrame(usingName:)
func (_NSWindowClass NSWindowClass) RemoveFrameUsingName(name NSWindowFrameAutosaveName) {
	objc.Send[objc.ID](objc.ID(_NSWindowClass.class), objc.Sel("removeFrameUsingName:"), objc.String(string(name)))
}

// Returns a new instance of a given standard window button, sized
// appropriately for a given window style.
//
// b: The type of standard window button to return.
//
// styleMask: The window style for which `b` is to be sized. See [NSWindow.StyleMask] for
// the list of allowable values.
//
// # Return Value
//
// The new window button of the type identified by `b`; `nil` when no such
// button type exists.
//
// # Discussion
//
// The caller is responsible for adding the button to the view hierarchy and
// for setting the target to be the window.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/standardWindowButton(_:for:)
//
// [NSWindow.StyleMask]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
func (_NSWindowClass NSWindowClass) StandardWindowButtonForStyleMask(b NSWindowButton, styleMask NSWindowStyleMask) NSButton {
	rv := objc.Send[objc.ID](objc.ID(_NSWindowClass.class), objc.Sel("standardWindowButton:forStyleMask:"), b, styleMask)
	return NSButtonFromID(rv)
}

// Returns the number of the frontmost window that would be hit by a
// mouse-down at the specified screen location.
//
// point: The location of the mouse-down in screen coordinates.
//
// windowNumber: If non-0, the search will start below `windowNumber` window in z-order.
//
// # Return Value
//
// The window number of the window under the point. The window number returned
// may correspond to a window in another application.
//
// # Discussion
//
// Because this method uses the same rules as mouse-down hit-testing, windows
// with transparency at the given point, and windows that ignore mouse events,
// will not be returned.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/windowNumber(at:belowWindowWithWindowNumber:)
func (_NSWindowClass NSWindowClass) WindowNumberAtPointBelowWindowWithWindowNumber(point corefoundation.CGPoint, windowNumber int) int {
	rv := objc.Send[int](objc.ID(_NSWindowClass.class), objc.Sel("windowNumberAtPoint:belowWindowWithWindowNumber:"), point, windowNumber)
	return rv
}

// Returns the default animation that should be performed for the specified
// key.
//
// key: The action name or property specified as a string.
//
// # Return Value
//
// The animation to perform. A subclass of [CAAnimation].
//
// # Discussion
//
// The [NSAnimatablePropertyContainer] method consults this class method when
// its search of the receivers Getting the Animator Proxy dictionary fails to
// return an animation for `key`.
//
// An animatable property container should implement this method to return a
// default animation to be performed for each key that it wants to make
// auto-animatable, where `key` usually references a property of the receiver,
// but can also specify a special animation trigger
// ([NSAnimationTriggerOrderIn] or [NSAnimationTriggerOrderOut]).
//
// A developer implementing a custom view subclass, can enable automatic
// animation for properties by overriding this method, and having it return
// the desired default [CAAnimation] subclass to use for each of the property
// keys of interest. The override should defer to super for any keys it
// doesn’t specifically handle, facilitating inheritance of default
// animation specifications. The following is an example of such an
// implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/defaultAnimation(forKey:)
//
// [CAAnimation]: https://developer.apple.com/documentation/QuartzCore/CAAnimation
// [NSAnimationTriggerOrderIn]: https://developer.apple.com/documentation/AppKit/NSAnimationTriggerOrderIn
// [NSAnimationTriggerOrderOut]: https://developer.apple.com/documentation/AppKit/NSAnimationTriggerOrderOut
//
// [CAAnimation]: https://developer.apple.com/documentation/QuartzCore/CAAnimation
func (_NSWindowClass NSWindowClass) DefaultAnimationForKey(key NSAnimatablePropertyKey) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NSWindowClass.class), objc.Sel("defaultAnimationForKey:"), objc.String(string(key)))
	return objectivec.Object{ID: rv}
}

// The window’s delegate.
//
// # Discussion
//
// The value of this property is `nil` if the window doesn’t have a
// delegate.
//
// A window object’s delegate is inserted in the responder chain after the
// window itself and is informed of various actions by the window through
// delegation messages.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/delegate
func (w NSWindow) Delegate() NSWindowDelegate {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("delegate"))
	return NSWindowDelegateObjectFromID(rv)
}
func (w NSWindow) SetDelegate(value NSWindowDelegate) {
	objc.Send[struct{}](w.ID, objc.Sel("setDelegate:"), value)
}

// The main content view controller for the window.
//
// # Discussion
//
// The value of this property provides the content view of the window. Setting
// this value removes the existing value of [ContentView] and makes the
// `contentViewController.View()` the main content view for the window. By
// default, the value of this property is `nil`.
//
// The content view controller controls only the [ContentView] object, and not
// the title of the window. The window title can easily be bound to the
// [ContentViewController] object using code such as: `[window NSTitleBinding
// contentViewController @"title" nil]`. Setting [ContentViewController]
// causes the window to resize based on the current size of the
// [ContentViewController]; to restrict the size of the window, use Auto
// Layout (note that the value of this property is encoded in the NIB).
// Directly assigning a [ContentView] value clears out the root view
// controller.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/contentViewController
func (w NSWindow) ContentViewController() INSViewController {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("contentViewController"))
	return NSViewControllerFromID(objc.ID(rv))
}
func (w NSWindow) SetContentViewController(value INSViewController) {
	objc.Send[struct{}](w.ID, objc.Sel("setContentViewController:"), value)
}

// The window’s content view, the highest accessible view object in the
// window’s view hierarchy.
//
// # Discussion
//
// The window retains the new content view and owns it thereafter. The `view`
// object is resized to fit precisely within the content area of the window.
// You can modify the content view’s coordinate system through its bounds
// rectangle, but you can’t alter its frame rectangle (its size or location)
// directly.
//
// Setting this property releases the old content view. If you plan to reuse
// it, be sure to retain it before changing the property value and to release
// it as appropriate when adding it to another [NSWindow] or [NSView] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/contentView
func (w NSWindow) ContentView() INSView {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("contentView"))
	return NSViewFromID(objc.ID(rv))
}
func (w NSWindow) SetContentView(value INSView) {
	objc.Send[struct{}](w.ID, objc.Sel("setContentView:"), value)
}

// Flags that describe the window’s current style, such as if it’s
// resizable or in full-screen mode.
//
// # Discussion
//
// The [StyleMask] is settable on macOS 10.6 and later. Setting this property
// has the same restrictions as the `styleMask` parameter of
// [InitWithContentRectStyleMaskBackingDefer]. Changing the style mask may
// cause the view hierarchy to be rebuilt.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/styleMask-swift.property
func (w NSWindow) StyleMask() NSWindowStyleMask {
	rv := objc.Send[NSWindowStyleMask](w.ID, objc.Sel("styleMask"))
	return NSWindowStyleMask(rv)
}
func (w NSWindow) SetStyleMask(value NSWindowStyleMask) {
	objc.Send[struct{}](w.ID, objc.Sel("setStyleMask:"), value)
}

// A Boolean value that indicates whether the window is able to receive
// keyboard and mouse events even when some other window is being run modally.
//
// # Discussion
//
// The value of this property is true if the window is able to receive
// keyboard and mouse events even when some other window is being run modally;
// otherwise, false. By default, the [NSWindow] value of this property is
// false. Only subclasses of [NSPanel] should override this default.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/worksWhenModal
func (w NSWindow) WorksWhenModal() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("worksWhenModal"))
	return rv
}

// The window’s alpha value.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/alphaValue
func (w NSWindow) AlphaValue() float64 {
	rv := objc.Send[float64](w.ID, objc.Sel("alphaValue"))
	return rv
}
func (w NSWindow) SetAlphaValue(value float64) {
	objc.Send[struct{}](w.ID, objc.Sel("setAlphaValue:"), value)
}

// The color of the window’s background.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/backgroundColor
func (w NSWindow) BackgroundColor() INSColor {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("backgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (w NSWindow) SetBackgroundColor(value INSColor) {
	objc.Send[struct{}](w.ID, objc.Sel("setBackgroundColor:"), value)
}

// The window’s color space.
//
// # Discussion
//
// The value of this property is `nil` if the window does not have a backing
// store, and is off-screen.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/colorSpace
func (w NSWindow) ColorSpace() INSColorSpace {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("colorSpace"))
	return NSColorSpaceFromID(objc.ID(rv))
}
func (w NSWindow) SetColorSpace(value INSColorSpace) {
	objc.Send[struct{}](w.ID, objc.Sel("setColorSpace:"), value)
}

// A Boolean value that indicates whether the window can hide when its
// application becomes hidden.
//
// # Discussion
//
// The value of this property is true if the window can hide when its
// application becomes hidden (during execution of the
// `NSApplication```NSApplication/hide(_:)“ method); otherwise, false. By
// default, the value of the property is true.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/canHide
func (w NSWindow) CanHide() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("canHide"))
	return rv
}
func (w NSWindow) SetCanHide(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setCanHide:"), value)
}

// A Boolean value that indicates whether the window is on the currently
// active space.
//
// # Discussion
//
// The value of this property is true if the window is on the currently active
// space; otherwise, false. For visible windows, this property indicates
// whether the window is currently visible on the active space. For nonvisible
// windows, it indicates whether ordering the window onscreen would cause it
// to be on the active space.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/isOnActiveSpace
func (w NSWindow) OnActiveSpace() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isOnActiveSpace"))
	return rv
}

// A Boolean value that indicates whether the window is removed from the
// screen when its application becomes inactive.
//
// # Discussion
//
// The value of this property is true if the window is removed from the screen
// when its application is deactivated; false if it remains onscreen. The
// default value for [NSWindow] is false; the default value for [NSPanel] is
// true.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/hidesOnDeactivate
func (w NSWindow) HidesOnDeactivate() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hidesOnDeactivate"))
	return rv
}
func (w NSWindow) SetHidesOnDeactivate(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setHidesOnDeactivate:"), value)
}

// A value that identifies the window’s behavior in window collections.
//
// # Discussion
//
// The possible values for this property are listed in
// [NSWindow.CollectionBehavior].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/collectionBehavior-swift.property
//
// [NSWindow.CollectionBehavior]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct
func (w NSWindow) CollectionBehavior() NSWindowCollectionBehavior {
	rv := objc.Send[NSWindowCollectionBehavior](w.ID, objc.Sel("collectionBehavior"))
	return NSWindowCollectionBehavior(rv)
}
func (w NSWindow) SetCollectionBehavior(value NSWindowCollectionBehavior) {
	objc.Send[struct{}](w.ID, objc.Sel("setCollectionBehavior:"), value)
}

// A Boolean value that indicates whether the window is opaque.
//
// # Discussion
//
// The value of this property is true when the window is opaque; otherwise,
// false.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/isOpaque
func (w NSWindow) Opaque() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isOpaque"))
	return rv
}
func (w NSWindow) SetOpaque(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setOpaque:"), value)
}

// A Boolean value that indicates whether the window has a shadow.
//
// # Discussion
//
// The value of this property is true when the window has a shadow; otherwise,
// false. If you change the value of this property, the window shadow is
// invalidated, forcing the window shadow to be recomputed.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/hasShadow
func (w NSWindow) HasShadow() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasShadow"))
	return rv
}
func (w NSWindow) SetHasShadow(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setHasShadow:"), value)
}

// A Boolean value that indicates whether the window prevents application
// termination when modal.
//
// # Discussion
//
// The value of this property is true if the window prevents application
// termination when modal; otherwise, false. The default value is true.
//
// Usually, application termination is prevented when a modal window or sheet
// is open, without consulting the application delegate. Some windows may wish
// not to prevent termination, however. Setting this property to false
// overrides the default behavior and allows termination to proceed even if
// the window is open, either through the sudden termination path if enabled,
// or after consulting the application delegate.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/preventsApplicationTerminationWhenModal
func (w NSWindow) PreventsApplicationTerminationWhenModal() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("preventsApplicationTerminationWhenModal"))
	return rv
}
func (w NSWindow) SetPreventsApplicationTerminationWhenModal(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setPreventsApplicationTerminationWhenModal:"), value)
}

// An object that the window inherits its appearance from.
//
// # Discussion
//
// The default value of this property is [NSApp]. The window uses key-value
// observing to monitor the source’s [EffectiveAppearance] for changes.
// Typically, you use this property for child windows shown from a parent
// window or specific view.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/appearanceSource
func (w NSWindow) AppearanceSource() objectivec.Object {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("appearanceSource"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (w NSWindow) SetAppearanceSource(value objectivec.Object) {
	objc.Send[struct{}](w.ID, objc.Sel("setAppearanceSource:"), value)
}

// The depth limit of the window.
//
// # Discussion
//
// The value of this property can be examined with the Application Kit
// functions [NSPlanarFromDepth], [NSColorSpaceFromDepth],
// [NSBitsPerSampleFromDepth], and [NSBitsPerPixelFromDepth]. In addition, the
// [NSBestDepth] function provides the best depth limit based on a set of
// parameters.
//
// Setting this property to `0` sets the depth limit to the window’s default
// depth limit. A depth limit of `0` can be useful for reverting a window
// object to its initial depth. You can also use one of the explicit bit
// depths defined in `Explicit Window Depth Limits`
// ([NSWindowDepthTwentyfourBitRGB] is the default).
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/depthLimit
func (w NSWindow) DepthLimit() NSWindowDepth {
	rv := objc.Send[NSWindowDepth](w.ID, objc.Sel("depthLimit"))
	return NSWindowDepth(rv)
}
func (w NSWindow) SetDepthLimit(value NSWindowDepth) {
	objc.Send[struct{}](w.ID, objc.Sel("setDepthLimit:"), value)
}

// A Boolean value that indicates whether the window’s depth limit can
// change to match the depth of the screen it’s on.
//
// # Discussion
//
// The value of this property is true when the window has a dynamic depth
// limit; otherwise, false. When the value of [HasDynamicDepthLimit] is false,
// the window uses either its preset depth limit or the default depth limit. A
// different, and non-dynamic, depth limit can be set using the [DepthLimit]
// property.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/hasDynamicDepthLimit
func (w NSWindow) HasDynamicDepthLimit() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasDynamicDepthLimit"))
	return rv
}

// The window number of the window’s window device.
//
// # Discussion
//
// Each window device in an application is given a unique window number—note
// that this isn’t the same as the global window number assigned by the
// window server. This number can be used to identify the window device with
// the [OrderWindowRelativeTo] method and in the AppKit function
// [NSWindowList].
//
// If the window doesn’t have a window device, the value of this property is
// equal to or less than `0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/windowNumber
func (w NSWindow) WindowNumber() int {
	rv := objc.Send[int](w.ID, objc.Sel("windowNumber"))
	return rv
}

// A dictionary containing information about the window’s resolution, such
// as color, depth, and so on.
//
// # Discussion
//
// This information is useful for tuning images and colors to the window’s
// display capabilities. The contents of the dictionary are described in
// `Display Device—Descriptions`.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/deviceDescription
func (w NSWindow) DeviceDescription() foundation.INSDictionary {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("deviceDescription"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the window can be displayed at the
// login window.
//
// # Discussion
//
// The value of this property is true when the window can be displayed at the
// login window; otherwise, false. By default, the value is false.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/canBecomeVisibleWithoutLogin
func (w NSWindow) CanBecomeVisibleWithoutLogin() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("canBecomeVisibleWithoutLogin"))
	return rv
}
func (w NSWindow) SetCanBecomeVisibleWithoutLogin(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setCanBecomeVisibleWithoutLogin:"), value)
}

// A Boolean value that indicates the level of access other processes have to
// the window’s content.
//
// # Discussion
//
// The value of this property represents the sharing level of the window’s
// content. See [NSWindow.SharingType] for possible values.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/sharingType-swift.property
//
// [NSWindow.SharingType]: https://developer.apple.com/documentation/AppKit/NSWindow/SharingType-swift.enum
func (w NSWindow) SharingType() NSWindowSharingType {
	rv := objc.Send[NSWindowSharingType](w.ID, objc.Sel("sharingType"))
	return NSWindowSharingType(rv)
}
func (w NSWindow) SetSharingType(value NSWindowSharingType) {
	objc.Send[struct{}](w.ID, objc.Sel("setSharingType:"), value)
}

// The window’s backing store type.
//
// # Discussion
//
// The possible values for this property are described in
// [NSWindow.BackingStoreType]. You can set the property only to switch a
// buffered window to retained or vice versa; you can’t change the backing
// type to or from nonretained after initializing a [NSWindow] object (an
// error is generated if you attempt to do so).
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/backingType
//
// [NSWindow.BackingStoreType]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType
func (w NSWindow) BackingType() NSBackingStoreType {
	rv := objc.Send[NSBackingStoreType](w.ID, objc.Sel("backingType"))
	return NSBackingStoreType(rv)
}
func (w NSWindow) SetBackingType(value NSBackingStoreType) {
	objc.Send[struct{}](w.ID, objc.Sel("setBackingType:"), value)
}

// The window’s window controller.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/windowController
func (w NSWindow) WindowController() INSWindowController {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("windowController"))
	return NSWindowControllerFromID(objc.ID(rv))
}
func (w NSWindow) SetWindowController(value INSWindowController) {
	objc.Send[struct{}](w.ID, objc.Sel("setWindowController:"), value)
}

// The sheet attached to the window.
//
// # Discussion
//
// The value of this property is `nil` when the window doesn’t have a sheet
// attached.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/attachedSheet
func (w NSWindow) AttachedSheet() INSWindow {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("attachedSheet"))
	return NSWindowFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the window has ever run as a modal
// sheet.
//
// # Discussion
//
// The value of this property is true if the window has ever run as a modal
// sheet; otherwise, false.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/isSheet
func (w NSWindow) Sheet() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isSheet"))
	return rv
}

// The window to which the sheet is attached.
//
// # Discussion
//
// The value of this property is `nil` if the receiver is not a sheet or has
// no sheet parent.
//
// The window object in this property refers to the window to which the sheet
// is logically attached, regardless of appearance. The parent window–sheet
// relationship begins with the beginning of the sheet (for example, through
// [BeginSheetCompletionHandler]) and ends with the sheet’s dismissal (for
// example, through [EndSheet]).
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/sheetParent
func (w NSWindow) SheetParent() INSWindow {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("sheetParent"))
	return NSWindowFromID(objc.ID(rv))
}

// An array of the sheets currently attached to the window.
//
// # Discussion
//
// The value of this property is an ordered array that contains—in
// top-to-bottom order—the presented sheets that are attached to the window,
// followed by queued sheets, in the order they were queued. The array
// doesn’t include nested sheets or subsheets.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/sheets
func (w NSWindow) Sheets() []NSWindow {
	rv := objc.Send[[]objc.ID](w.ID, objc.Sel("sheets"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSWindow {
		return NSWindowFromID(id)
	})
}

// The window’s frame rectangle in screen coordinates, including the title
// bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/frame
func (w NSWindow) Frame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](w.ID, objc.Sel("frame"))
	return corefoundation.CGRect(rv)
}

// The window’s aspect ratio, which constrains the size of its frame
// rectangle to integral multiples of this ratio when the user resizes it.
//
// # Discussion
//
// The size of the window’s frame rectangle is constrained to integral
// multiples of this ratio when the user resizes it. You can set an [NSWindow]
// object’s size to any ratio programmatically.
//
// An [NSWindow] object’s aspect ratio and its resize increments are
// mutually exclusive attributes. In fact, setting one attribute cancels the
// setting of the other. For example, to cancel an established aspect ratio
// setting for an [NSWindow] object, you can set the [ResizeIncrements]
// property with the width and height set to `1.0`:
//
// The [ContentAspectRatio] property takes precedence over this property.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/aspectRatio
func (w NSWindow) AspectRatio() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](w.ID, objc.Sel("aspectRatio"))
	return corefoundation.CGSize(rv)
}
func (w NSWindow) SetAspectRatio(value corefoundation.CGSize) {
	objc.Send[struct{}](w.ID, objc.Sel("setAspectRatio:"), value)
}

// The minimum size to which the window’s frame (including its title bar)
// can be sized.
//
// # Discussion
//
// The minimum size constraint is enforced for resizing by the user as well as
// for the `setFrame...` methods other than [SetFrameDisplay] and
// [SetFrameDisplayAnimate].
//
// The [NSWindow] method takes precedence over this property.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/minSize
func (w NSWindow) MinSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](w.ID, objc.Sel("minSize"))
	return corefoundation.CGSize(rv)
}
func (w NSWindow) SetMinSize(value corefoundation.CGSize) {
	objc.Send[struct{}](w.ID, objc.Sel("setMinSize:"), value)
}

// The maximum size to which the window’s frame (including its title bar)
// can be sized.
//
// # Discussion
//
// The maximum size constraint is enforced for resizing by the user as well as
// for the `setFrame...` methods other than [SetFrameDisplay] and
// [SetFrameDisplayAnimate]. Note that the window server limits window sizes
// to 10,000.
//
// The default maximum size of a window is `{FLT_MAX, FLT_MAX}` (`FLT_MAX` is
// defined in `/usr/include/float.H()`). When the maximum size of a window has
// been set, there is no way to reset it other than by specifying this default
// maximum size.
//
// The [ContentMaxSize] property takes precedence over this property.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/maxSize
func (w NSWindow) MaxSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](w.ID, objc.Sel("maxSize"))
	return corefoundation.CGSize(rv)
}
func (w NSWindow) SetMaxSize(value corefoundation.CGSize) {
	objc.Send[struct{}](w.ID, objc.Sel("setMaxSize:"), value)
}

// A Boolean value that indicates whether the window is in a zoomed state.
//
// # Discussion
//
// The value of this property is true if the window is in a zoomed state;
// otherwise, false.
//
// The zoomed state of the window is determined using the following steps:
//
// - If the delegate or the window class implements
// [WindowWillUseStandardFrameDefaultFrame], it is invoked to obtain the
// zoomed frame of the window. The value of [Zoomed] is then determined by
// whether or not the current window frame is equal to the zoomed frame. - If
// the neither the delegate nor the window class implements
// [WindowWillUseStandardFrameDefaultFrame], a default frame that nearly fits
// the screen is chosen. If the delegate or window class implements
// [WindowWillUseStandardFrameDefaultFrame], it is invoked to validate the
// proposed zoomed frame. After the zoomed frame is validated, the value of
// [Zoomed] is determined by whether or not the current window frame is equal
// to the zoomed frame.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/isZoomed
func (w NSWindow) Zoomed() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isZoomed"))
	return rv
}

// The flags field of the event record for the mouse-down event that initiated
// the resizing session.
//
// # Discussion
//
// The value of this property is a mask indicating which of the modifier keys
// was held down when the mouse-down event occurred. The flags are listed in
// [NSEvent] class’s [ModifierFlags] method description. The property is
// valid only while the window is being resized.
//
// You can use this property to constrain the direction or amount of resizing.
// Because of its limited validity, this property should only be accessed from
// within an implementation of the delegate method [WindowWillResizeToSize].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/resizeFlags
func (w NSWindow) ResizeFlags() NSEventModifierFlags {
	rv := objc.Send[NSEventModifierFlags](w.ID, objc.Sel("resizeFlags"))
	return NSEventModifierFlags(rv)
}

// The window’s resizing increments.
//
// # Discussion
//
// This property restricts the user’s ability to resize the window so the
// width and height change by multiples of width and height increments. As the
// user resizes the window, its size changes by multiples of
// `increments“XCUIElementTypeWidth` and `increments“XCUIElementTypeHeight`,
// which should be whole numbers, 1.0 or greater. Whatever the current
// resizing increments, you can set an [NSWindow] object’s size to any
// height and width programmatically.
//
// Resize increments and aspect ratio are mutually exclusive attributes. For
// more information, see [AspectRatio].
//
// The [ContentResizeIncrements] property takes precedence over this property.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/resizeIncrements
func (w NSWindow) ResizeIncrements() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](w.ID, objc.Sel("resizeIncrements"))
	return corefoundation.CGSize(rv)
}
func (w NSWindow) SetResizeIncrements(value corefoundation.CGSize) {
	objc.Send[struct{}](w.ID, objc.Sel("setResizeIncrements:"), value)
}

// A Boolean value that indicates whether the window tries to optimize
// user-initiated resize operations by preserving the content of views that
// have not changed.
//
// # Discussion
//
// The value of this property is true if the window tries to optimize live
// resize operations by preserving the content of views that have not moved;
// otherwise, false. By default, live-resize optimization is turned on.
//
// When live-resize optimization is active, the window redraws only those
// views that moved (or do not support this optimization) during a live resize
// operation. You might consider disabling this optimization for the window if
// none of the window’s contained views can take advantage of it. Disabling
// the optimization for the window prevents it from checking each view to see
// if the optimization is supported.
//
// See [PreservesContentDuringLiveResize] in [NSView] for additional
// information on how to support this optimization.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/preservesContentDuringLiveResize
func (w NSWindow) PreservesContentDuringLiveResize() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("preservesContentDuringLiveResize"))
	return rv
}
func (w NSWindow) SetPreservesContentDuringLiveResize(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setPreservesContentDuringLiveResize:"), value)
}

// A Boolean value that indicates whether the window is being resized by the
// user.
//
// # Discussion
//
// The value of this property is true if the window is being live resized;
// otherwise, false.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/inLiveResize
func (w NSWindow) InLiveResize() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("inLiveResize"))
	return rv
}

// The window’s content aspect ratio.
//
// # Discussion
//
// By default, the content aspect ratio (that is, height in relation to width)
// is `(0, 0)`. If you set the aspect ratio of a window’s content view, the
// dimensions of its content rectangle are constrained to integral multiples
// of that ratio when users resize it. You can set a window’s content view
// to any size programmatically, regardless of its aspect ratio. The value of
// this property takes precedence over [AspectRatio].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/contentAspectRatio
func (w NSWindow) ContentAspectRatio() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](w.ID, objc.Sel("contentAspectRatio"))
	return corefoundation.CGSize(rv)
}
func (w NSWindow) SetContentAspectRatio(value corefoundation.CGSize) {
	objc.Send[struct{}](w.ID, objc.Sel("setContentAspectRatio:"), value)
}

// The minimum size of the window’s content view in the window’s base
// coordinate system.
//
// # Discussion
//
// The minimum size constraint is enforced for resizing by the user as well as
// for the [SetContentSize] method and the `setFrame...` methods other than
// [SetFrameDisplay] and [SetFrameDisplayAnimate]. This method takes
// precedence over the [MinSize] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/contentMinSize
func (w NSWindow) ContentMinSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](w.ID, objc.Sel("contentMinSize"))
	return corefoundation.CGSize(rv)
}
func (w NSWindow) SetContentMinSize(value corefoundation.CGSize) {
	objc.Send[struct{}](w.ID, objc.Sel("setContentMinSize:"), value)
}

// The maximum size of the window’s content view in the window’s base
// coordinate system.
//
// # Discussion
//
// The maximum size constraint is enforced for resizing by the user as well as
// for the [SetContentSize] method and the `setFrame...` methods other than
// [SetFrameDisplay] and [SetFrameDisplayAnimate]. This method takes
// precedence over the [MaxSize] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/contentMaxSize
func (w NSWindow) ContentMaxSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](w.ID, objc.Sel("contentMaxSize"))
	return corefoundation.CGSize(rv)
}
func (w NSWindow) SetContentMaxSize(value corefoundation.CGSize) {
	objc.Send[struct{}](w.ID, objc.Sel("setContentMaxSize:"), value)
}

// The window’s content-view resizing increments.
//
// # Discussion
//
// The value of this property restricts the user’s ability to resize the
// window so the width and height of its content view change by multiples of
// width and height increments. As the user resizes the window, the size of
// its content view changes by integral multiples of
// `contentResizeIncrements.Width()` and `contentResizeIncrements.Height()`.
// However, you can set a window’s size to any width and height
// programmatically. This property takes precedence over [ResizeIncrements].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/contentResizeIncrements
func (w NSWindow) ContentResizeIncrements() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](w.ID, objc.Sel("contentResizeIncrements"))
	return corefoundation.CGSize(rv)
}
func (w NSWindow) SetContentResizeIncrements(value corefoundation.CGSize) {
	objc.Send[struct{}](w.ID, objc.Sel("setContentResizeIncrements:"), value)
}

// A value used by Auto Layout constraints to automatically bind to the value
// of [ContentLayoutRect].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/contentLayoutGuide
func (w NSWindow) ContentLayoutGuide() objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("contentLayoutGuide"))
	return objectivec.Object{ID: rv}
}

// The area inside the window that is for non-obscured content, in window
// coordinates.
//
// # Discussion
//
// Typically, the area represented by this property is the same as the frame
// of the [ContentView]. However, for windows with
// [NSFullSizeContentViewWindowMask] set, there needs to be a way to determine
// the portion that is not under the toolbar. The [ContentLayoutRect] property
// contains the portion of the layout that is not obscured under the toolbar.
// This property is KVO compliant.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/contentLayoutRect
//
// [NSFullSizeContentViewWindowMask]: https://developer.apple.com/documentation/AppKit/NSFullSizeContentViewWindowMask
func (w NSWindow) ContentLayoutRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](w.ID, objc.Sel("contentLayoutRect"))
	return corefoundation.CGRect(rv)
}

// A maximum size that is used to determine if a window can fit when it is in
// full screen in a tile.
//
// # Discussion
//
// By default, the system uses Auto Layout to determine the maximum size, so
// applications that don’t change window content upon entering full screen
// should not need to set the value of [MaxFullScreenContentSize]. (If Auto
// Layout is not used, the system queries [ContentMinSize] and
// [ContentMaxSize].) If an application does significant rework of the user
// interface in full screen, then it may be necessary to set the value of
// [MaxFullScreenContentSize]. You can use this property even if the window
// does not support full screen, but can be implicitly opted into supporting a
// full screen tile based on resizing behavior and window properties (for more
// information, see the [CollectionBehavior] property).
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/maxFullScreenContentSize
func (w NSWindow) MaxFullScreenContentSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](w.ID, objc.Sel("maxFullScreenContentSize"))
	return corefoundation.CGSize(rv)
}
func (w NSWindow) SetMaxFullScreenContentSize(value corefoundation.CGSize) {
	objc.Send[struct{}](w.ID, objc.Sel("setMaxFullScreenContentSize:"), value)
}

// A minimum size that is used to determine if a window can fit when it is in
// full screen in a tile.
//
// # Discussion
//
// By default, the system uses Auto Layout to determine the minimum size, so
// applications that don’t change window content upon entering full screen
// should not need to set the value of [MinFullScreenContentSize]. (If Auto
// Layout is not used, the system queries [ContentMinSize] and
// [ContentMaxSize].) If an application does significant rework of the user
// interface in full screen, then it may be necessary to set the value of
// [MinFullScreenContentSize]. You can use this property even if the window
// does not support full screen, but can be implicitly opted into supporting a
// full screen tile based on resizing behavior and window properties (for more
// information, see the [CollectionBehavior] property).
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/minFullScreenContentSize
func (w NSWindow) MinFullScreenContentSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](w.ID, objc.Sel("minFullScreenContentSize"))
	return corefoundation.CGSize(rv)
}
func (w NSWindow) SetMinFullScreenContentSize(value corefoundation.CGSize) {
	objc.Send[struct{}](w.ID, objc.Sel("setMinFullScreenContentSize:"), value)
}

// The window level of the window.
//
// # Discussion
//
// See `Window Levels` for a list of possible values. Each level in the list
// groups windows within it in front of those in all preceding groups.
// Floating windows, for example, appear in front of all normal-level windows.
//
// The constant [NSTornOffMenuWindowLevel] is preferable to its synonym,
// [NSSubmenuWindowLevel].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/level-swift.property
func (w NSWindow) Level() NSWindowLevel {
	rv := objc.Send[NSWindowLevel](w.ID, objc.Sel("level"))
	return NSWindowLevel(rv)
}
func (w NSWindow) SetLevel(value NSWindowLevel) {
	objc.Send[struct{}](w.ID, objc.Sel("setLevel:"), value)
}

// A Boolean value that indicates whether the window is visible onscreen (even
// when it’s obscured by other windows).
//
// # Discussion
//
// The value of this property is true when the window is onscreen (even if
// it’s obscured by other windows); otherwise, false.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/isVisible
func (w NSWindow) Visible() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isVisible"))
	return rv
}

// The occlusion state of the window.
//
// # Discussion
//
// When the value of this property is [NSWindowOcclusionStateVisible], at
// least part of the window is visible; otherwise, the window is fully
// occluded.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/occlusionState-swift.property
func (w NSWindow) OcclusionState() NSWindowOcclusionState {
	rv := objc.Send[NSWindowOcclusionState](w.ID, objc.Sel("occlusionState"))
	return NSWindowOcclusionState(rv)
}

// The name used to automatically save the window’s frame rectangle data in
// the defaults system.
//
// # Discussion
//
// Assigning a value to this property reloads the associated frame, which can
// result in moving the window to that frame’s location.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/frameAutosaveName-swift.property
func (w NSWindow) FrameAutosaveName() NSWindowFrameAutosaveName {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("frameAutosaveName"))
	return NSWindowFrameAutosaveName(foundation.NSStringFromID(rv).String())
}

// A string representation of the window’s frame rectangle.
//
// # Discussion
//
// The value of this property is a string that can be used in a later call to
// [SetFrameFromString].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/frameDescriptor
func (w NSWindow) StringWithSavedFrame() NSWindowPersistableFrameDescriptor {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("stringWithSavedFrame"))
	return NSWindowPersistableFrameDescriptor(foundation.NSStringFromID(rv).String())
}

// A Boolean value that indicates whether the window is the key window for the
// application.
//
// # Discussion
//
// The value of this property is true if the window is the key window for the
// application; otherwise, false.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/isKeyWindow
func (w NSWindow) KeyWindow() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isKeyWindow"))
	return rv
}

// A Boolean value that indicates whether the window can become the key
// window.
//
// # Discussion
//
// The value of this property is true if the window can become the key window,
// otherwise, false.
//
// Attempts to make the window the key window are abandoned if the value of
// this property is false. The value of this property is true if the window
// has a title bar or a resize bar, or false otherwise.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/canBecomeKey
func (w NSWindow) CanBecomeKeyWindow() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("canBecomeKeyWindow"))
	return rv
}

// A Boolean value that indicates whether the window is the application’s
// main window.
//
// # Discussion
//
// The value of this property is true when the window is the main window for
// the application; otherwise, false.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/isMainWindow
func (w NSWindow) MainWindow() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isMainWindow"))
	return rv
}

// A Boolean value that indicates whether the window can become the
// application’s main window.
//
// # Discussion
//
// The value of this property is true when the window can become the main
// window; otherwise, false.
//
// Attempts to make the window the main window are abandoned if the value of
// this property is false. The value of the property is true if the window is
// visible, is not an [NSPanel] object, and has a title bar or a resize
// mechanism. Otherwise, the value is false.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/canBecomeMain
func (w NSWindow) CanBecomeMainWindow() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("canBecomeMainWindow"))
	return rv
}

// The window’s toolbar.
//
// # Discussion
//
// For more information about toolbars, see [NSToolbar].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/toolbar
func (w NSWindow) Toolbar() INSToolbar {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("toolbar"))
	return NSToolbarFromID(objc.ID(rv))
}
func (w NSWindow) SetToolbar(value INSToolbar) {
	objc.Send[struct{}](w.ID, objc.Sel("setToolbar:"), value)
}

// An array of the window’s attached child windows.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/childWindows
func (w NSWindow) ChildWindows() []NSWindow {
	rv := objc.Send[[]objc.ID](w.ID, objc.Sel("childWindows"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSWindow {
		return NSWindowFromID(id)
	})
}

// The parent window to which the window is attached as a child.
//
// # Discussion
//
// This property should be set from a subclass when it is overridden by a
// subclass’s implementation. It should not be set otherwise.
//
// Note that calling [OrderOut] on a child window causes the window to be
// removed from its parent window before it is itself removed.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/parent
func (w NSWindow) ParentWindow() INSWindow {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("parentWindow"))
	return NSWindowFromID(objc.ID(rv))
}
func (w NSWindow) SetParentWindow(value INSWindow) {
	objc.Send[struct{}](w.ID, objc.Sel("setParentWindow:"), value)
}

// The button cell that performs as if clicked when the window receives a
// Return (or Enter) key event.
//
// # Discussion
//
// This cell draws itself as the focal element for keyboard interface control,
// unless another button cell is focused on, in which case the default button
// cell temporarily draws itself as normal and disables its key equivalent.
//
// The window receives a Return key event if no responder in its responder
// chain claims it, or if the user presses the Control key along with the
// Return key.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/defaultButtonCell
func (w NSWindow) DefaultButtonCell() INSButtonCell {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("defaultButtonCell"))
	return NSButtonCellFromID(objc.ID(rv))
}
func (w NSWindow) SetDefaultButtonCell(value INSButtonCell) {
	objc.Send[struct{}](w.ID, objc.Sel("setDefaultButtonCell:"), value)
}

// A Boolean value that indicates whether the window is excluded from the
// application’s Windows menu.
//
// # Discussion
//
// The value of this property is true when the window is excluded from the
// Windows menu; otherwise, false. The default initial setting is false.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/isExcludedFromWindowsMenu
func (w NSWindow) ExcludedFromWindowsMenu() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isExcludedFromWindowsMenu"))
	return rv
}
func (w NSWindow) SetExcludedFromWindowsMenu(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setExcludedFromWindowsMenu:"), value)
}

// A Boolean value that indicates whether the window’s cursor rectangles are
// enabled.
//
// # Discussion
//
// The value of this property is true when cursor rectangles are enabled;
// otherwise, false.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/areCursorRectsEnabled
func (w NSWindow) AreCursorRectsEnabled() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("areCursorRectsEnabled"))
	return rv
}

// A Boolean value that indicates whether the toolbar control button is
// currently displayed.
//
// # Discussion
//
// The value of this property is true if the standard toolbar button is
// currently displayed; otherwise, false. When clicked, the toolbar control
// button shows or hides a window’s toolbar. The toolbar control button
// appears in a window’s title bar. If the window does not have a toolbar,
// this property has no effect.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/showsToolbarButton
func (w NSWindow) ShowsToolbarButton() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("showsToolbarButton"))
	return rv
}
func (w NSWindow) SetShowsToolbarButton(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setShowsToolbarButton:"), value)
}

// A Boolean value that indicates whether the title bar draws its background.
//
// # Discussion
//
// When the value of this property is true, the title bar does not draw its
// background, which allows all content underneath it to show through. It only
// makes sense to set this property to true when
// [NSFullSizeContentViewWindowMask] is also set.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/titlebarAppearsTransparent
//
// [NSFullSizeContentViewWindowMask]: https://developer.apple.com/documentation/AppKit/NSFullSizeContentViewWindowMask
func (w NSWindow) TitlebarAppearsTransparent() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("titlebarAppearsTransparent"))
	return rv
}
func (w NSWindow) SetTitlebarAppearsTransparent(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setTitlebarAppearsTransparent:"), value)
}

// The style that determines the appearance and location of the toolbar in
// relation to the title bar.
//
// # Discussion
//
// The default value is [NSWindowToolbarStyleAutomatic].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/toolbarStyle-swift.property
func (w NSWindow) ToolbarStyle() NSWindowToolbarStyle {
	rv := objc.Send[NSWindowToolbarStyle](w.ID, objc.Sel("toolbarStyle"))
	return NSWindowToolbarStyle(rv)
}
func (w NSWindow) SetToolbarStyle(value NSWindowToolbarStyle) {
	objc.Send[struct{}](w.ID, objc.Sel("setToolbarStyle:"), value)
}

// The type of separator that the app displays between the title bar and
// content of a window.
//
// # Discussion
//
// The default value is [NSTitlebarSeparatorStyleAutomatic]. Changing this
// value overrides any preference by [NSSplitViewItem].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/titlebarSeparatorStyle
func (w NSWindow) TitlebarSeparatorStyle() NSTitlebarSeparatorStyle {
	rv := objc.Send[NSTitlebarSeparatorStyle](w.ID, objc.Sel("titlebarSeparatorStyle"))
	return NSTitlebarSeparatorStyle(rv)
}
func (w NSWindow) SetTitlebarSeparatorStyle(value NSTitlebarSeparatorStyle) {
	objc.Send[struct{}](w.ID, objc.Sel("setTitlebarSeparatorStyle:"), value)
}

// The direction the window’s title bar lays text out, either left to right
// or right to left.
//
// # Discussion
//
// The layout direction of the window title bar includes the standard window
// buttons (close, minimize, maximize) and the title for the window. In
// general, this returns [NSUserInterfaceLayoutDirectionRightToLeft] if the
// primary system language is right to left. The layout direction may be right
// to left even in applications that don’t have a right to left language
// localization. Refer to this value if the application uses
// [TitlebarAppearsTransparent] and places controls under the title bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/windowTitlebarLayoutDirection
func (w NSWindow) WindowTitlebarLayoutDirection() NSUserInterfaceLayoutDirection {
	rv := objc.Send[NSUserInterfaceLayoutDirection](w.ID, objc.Sel("windowTitlebarLayoutDirection"))
	return NSUserInterfaceLayoutDirection(rv)
}

// An array of title bar accessory view controllers that are currently added
// to the window.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/titlebarAccessoryViewControllers
func (w NSWindow) TitlebarAccessoryViewControllers() []NSTitlebarAccessoryViewController {
	rv := objc.Send[[]objc.ID](w.ID, objc.Sel("titlebarAccessoryViewControllers"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSTitlebarAccessoryViewController {
		return NSTitlebarAccessoryViewControllerFromID(id)
	})
}
func (w NSWindow) SetTitlebarAccessoryViewControllers(value []NSTitlebarAccessoryViewController) {
	objc.Send[struct{}](w.ID, objc.Sel("setTitlebarAccessoryViewControllers:"), objectivec.IObjectSliceToNSArray(value))
}

// An object that represents information about a window when it displays as a
// tab.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/tab
func (w NSWindow) Tab() INSWindowTab {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("tab"))
	return NSWindowTabFromID(objc.ID(rv))
}

// A value that allows a group of related windows.
//
// # Discussion
//
// By default, a window generates a tabbing identifier from inherent window
// properties, such as the window class name, the delegate class name, the
// window controller class name, and some additional state. Group windows
// together by using the same tabbing identifier.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/tabbingIdentifier-swift.property
func (w NSWindow) TabbingIdentifier() NSWindowTabbingIdentifier {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("tabbingIdentifier"))
	return NSWindowTabbingIdentifier(foundation.NSStringFromID(rv).String())
}
func (w NSWindow) SetTabbingIdentifier(value NSWindowTabbingIdentifier) {
	objc.Send[struct{}](w.ID, objc.Sel("setTabbingIdentifier:"), objc.String(string(value)))
}

// A value that indicates when a window displays tabs.
//
// # Discussion
//
// Set this to the desired tabbing mode before displaying a window. The
// default value is [NSWindowTabbingModeAutomatic]. When the value is
// [NSWindowTabbingModeAutomatic], the system uses [UserTabbingPreference] to
// determine tabbing behavior.
//
// For a list of possible values, see [NSWindow.TabbingMode].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/tabbingMode-swift.property
//
// [NSWindow.TabbingMode]: https://developer.apple.com/documentation/AppKit/NSWindow/TabbingMode-swift.enum
func (w NSWindow) TabbingMode() NSWindowTabbingMode {
	rv := objc.Send[NSWindowTabbingMode](w.ID, objc.Sel("tabbingMode"))
	return NSWindowTabbingMode(rv)
}
func (w NSWindow) SetTabbingMode(value NSWindowTabbingMode) {
	objc.Send[struct{}](w.ID, objc.Sel("setTabbingMode:"), value)
}

// An array of windows that display as tabs.
//
// # Discussion
//
// This property is `nil` if the window is not showing a tab bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/tabbedWindows
func (w NSWindow) TabbedWindows() []NSWindow {
	rv := objc.Send[[]objc.ID](w.ID, objc.Sel("tabbedWindows"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSWindow {
		return NSWindowFromID(id)
	})
}

// A group of windows that display together as a tab group.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/tabGroup
func (w NSWindow) TabGroup() INSWindowTabGroup {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("tabGroup"))
	return NSWindowTabGroupFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the window can display tooltips even
// when the application is in the background.
//
// # Discussion
//
// The value of this property is true if the window can display tooltips even
// when the application is in the background; otherwise, false. The default is
// false. Changing the value of this property does not take effect until the
// window changes to an active state.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/allowsToolTipsWhenApplicationIsInactive
func (w NSWindow) AllowsToolTipsWhenApplicationIsInactive() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("allowsToolTipsWhenApplicationIsInactive"))
	return rv
}
func (w NSWindow) SetAllowsToolTipsWhenApplicationIsInactive(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setAllowsToolTipsWhenApplicationIsInactive:"), value)
}

// The event currently being processed by the application.
//
// # Discussion
//
// The value of this property is given by calling by invoking the
// [NSApplication] method [CurrentEvent].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/currentEvent
func (w NSWindow) CurrentEvent() INSEvent {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("currentEvent"))
	return NSEventFromID(objc.ID(rv))
}

// The view that’s made first responder (also called the key view) the first
// time the window is placed onscreen.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/initialFirstResponder
func (w NSWindow) InitialFirstResponder() INSView {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("initialFirstResponder"))
	return NSViewFromID(objc.ID(rv))
}
func (w NSWindow) SetInitialFirstResponder(value INSView) {
	objc.Send[struct{}](w.ID, objc.Sel("setInitialFirstResponder:"), value)
}

// The window’s first responder.
//
// # Discussion
//
// The first responder is usually the first object in a responder chain to
// receive an event or action message. In most cases, the first responder is a
// view object that the user selects or activates with the mouse or keyboard.
//
// You can use the [FirstResponder] property in custom subclasses of responder
// classes ([NSWindow], [NSApplication], [NSView], and subclasses) to
// determine if an instance of the subclass is currently the first responder.
// You can also use it to help locate a text field that currently has
// first-responder status. For more on this subject, see [Event Handling
// Basics]. This property is key-value observing compliant.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/firstResponder
//
// [Event Handling Basics]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/EventOverview/EventHandlingBasics/EventHandlingBasics.html#//apple_ref/doc/uid/10000060i-CH5
func (w NSWindow) FirstResponder() INSResponder {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("firstResponder"))
	return NSResponderFromID(objc.ID(rv))
}

// The direction the window is currently using to change the key view.
//
// # Discussion
//
// The value of this property can be one of the values described in
// [NSWindow.SelectionDirection].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/keyViewSelectionDirection
//
// [NSWindow.SelectionDirection]: https://developer.apple.com/documentation/AppKit/NSWindow/SelectionDirection
func (w NSWindow) KeyViewSelectionDirection() NSSelectionDirection {
	rv := objc.Send[NSSelectionDirection](w.ID, objc.Sel("keyViewSelectionDirection"))
	return NSSelectionDirection(rv)
}

// A Boolean value that indicates whether the window automatically
// recalculates the key view loop when views are added.
//
// # Discussion
//
// The value of this property is true if the window automatically recalculates
// the key view loop when views are added; otherwise, false. If
// [AutorecalculatesKeyViewLoop] is false, the client code must update the key
// view loop manually or call [RecalculateKeyViewLoop] to have the window
// recalculate it.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/autorecalculatesKeyViewLoop
func (w NSWindow) AutorecalculatesKeyViewLoop() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("autorecalculatesKeyViewLoop"))
	return rv
}
func (w NSWindow) SetAutorecalculatesKeyViewLoop(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setAutorecalculatesKeyViewLoop:"), value)
}

// Indicates whether the receiver is the subject of an active SharePlay
// sharing session.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/hasActiveWindowSharingSession
func (w NSWindow) HasActiveWindowSharingSession() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasActiveWindowSharingSession"))
	return rv
}

// A Boolean value that indicates whether the window accepts mouse-moved
// events.
//
// # Discussion
//
// The value of this property is true when the window accepts (and
// distributes) mouse-moved events; otherwise, false. By default the value is
// false.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/acceptsMouseMovedEvents
func (w NSWindow) AcceptsMouseMovedEvents() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("acceptsMouseMovedEvents"))
	return rv
}
func (w NSWindow) SetAcceptsMouseMovedEvents(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setAcceptsMouseMovedEvents:"), value)
}

// A Boolean value that indicates whether the window is transparent to mouse
// events.
//
// # Discussion
//
// The value of this property is true when the window is transparent to mouse
// events; otherwise, false.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/ignoresMouseEvents
func (w NSWindow) IgnoresMouseEvents() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("ignoresMouseEvents"))
	return rv
}
func (w NSWindow) SetIgnoresMouseEvents(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setIgnoresMouseEvents:"), value)
}

// The current location of the pointer reckoned in the window’s base
// coordinate system, regardless of the current event being handled or of any
// events pending.
//
// # Discussion
//
// For the same information in screen coordinates, use [NSEvent]’s
// [MouseLocation].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/mouseLocationOutsideOfEventStream
func (w NSWindow) MouseLocationOutsideOfEventStream() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](w.ID, objc.Sel("mouseLocationOutsideOfEventStream"))
	return corefoundation.CGPoint(rv)
}

// A Boolean value indicating whether the window configuration is preserved
// between application launches.
//
// # Discussion
//
// Set this property to true if you want the window to be preserved or false
// if you do not want it preserved. By default, the value of this property is
// true if the window’s [StyleMask] property includes the
// [NSTitledWindowMask] flag. For other windows, the value is false. Setting a
// value explicitly overrides the default values.
//
// Windows should be preserved between launch cycles to maintain interface
// continuity for the user. During subsequent launch cycles, the system tries
// to recreate the window and restore its configuration to the preserved
// state. Configuration data is updated as needed and saved automatically by
// the system.
//
// If you enable preservation for a given window, you should also specify a
// restoration class for the window using the [RestorationClass] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/isRestorable
//
// [NSTitledWindowMask]: https://developer.apple.com/documentation/AppKit/NSTitledWindowMask
func (w NSWindow) Restorable() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isRestorable"))
	return rv
}
func (w NSWindow) SetRestorable(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setRestorable:"), value)
}

// The restoration class associated with the window.
//
// # Discussion
//
// The value of this property is a class object that conforms to the
// [NSWindowRestoration] protocol corresponding to the class to use to restore
// the window or `nil` if none is set.
//
// The restoration class of a window is responsible for recreating not just
// the window but any other objects needed to manage the window. This almost
// always involves creating a window controller and for multi-window document
// applications also involves creating a document object. Therefore, the
// restoration class must be able to create (or find existing instances of)
// all of these objects at launch time in your application. When prompted by
// AppKit, the restoration class creates or acquires a window that matches the
// same type that was preserved. It then passes that window back to AppKit,
// which proceeds to reconfigure the window with the preserved state
// information.
//
// If you mark your windows as restorable, you must associate a restoration
// class with them. For multi-window document applications, AppKit associates
// the [NSDocumentController] class with any document windows by default. That
// class recreates the preserved document objects, which in turn recreate the
// corresponding window controller and window objects. For other types of
// windows, you must set the restoration class explicitly.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/restorationClass
func (w NSWindow) RestorationClass() objc.Class {
	rv := objc.Send[objc.Class](w.ID, objc.Sel("restorationClass"))
	return rv
}
func (w NSWindow) SetRestorationClass(value objc.Class) {
	objc.Send[struct{}](w.ID, objc.Sel("setRestorationClass:"), value)
}

// A Boolean value that indicates whether any of the window’s views need to
// be displayed.
//
// # Discussion
//
// The value of this property is true when any of the window’s views need to
// be displayed; otherwise, false. You should rarely need to set this
// property; the [NSView] method [NeedsDisplay] and similar methods set it
// automatically.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/viewsNeedDisplay
func (w NSWindow) ViewsNeedDisplay() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("viewsNeedDisplay"))
	return rv
}
func (w NSWindow) SetViewsNeedDisplay(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setViewsNeedDisplay:"), value)
}

// A Boolean value that indicates whether the window allows multithreaded view
// drawing.
//
// # Discussion
//
// The value of this property is true if the window allows multithreaded view
// drawing; otherwise, false. The default value is true.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/allowsConcurrentViewDrawing
func (w NSWindow) AllowsConcurrentViewDrawing() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("allowsConcurrentViewDrawing"))
	return rv
}
func (w NSWindow) SetAllowsConcurrentViewDrawing(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setAllowsConcurrentViewDrawing:"), value)
}

// The window’s automatic animation behavior.
//
// # Discussion
//
// This property controls the automatic window animation behavior used when
// the [OrderFront] or [OrderOut] methods are called. See
// [NSWindow.AnimationBehavior] for the possible values of this property.
//
// By default, a window’s animation behavior is set to
// [NSWindowAnimationBehaviorDefault], which causes AppKit to determine the
// style of animation to use automatically based on its inference of a
// window’s “type” from various window properties. A window’s
// animation behavior can be set to [NSWindowAnimationBehaviorNone] to disable
// AppKit’s automatic animations for the window, which may be useful if that
// animation interferes with an animation that your application implements.
//
// The animation behavior can also be set to one of the other non-default
// [NSWindow.AnimationBehavior] values to override AppKit’s automatic
// inference of appropriate animation behavior based on the window’s
// apparent type, although this is not recommended.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/animationBehavior-swift.property
//
// [NSWindow.AnimationBehavior]: https://developer.apple.com/documentation/AppKit/NSWindow/AnimationBehavior-swift.enum
func (w NSWindow) AnimationBehavior() NSWindowAnimationBehavior {
	rv := objc.Send[NSWindowAnimationBehavior](w.ID, objc.Sel("animationBehavior"))
	return NSWindowAnimationBehavior(rv)
}
func (w NSWindow) SetAnimationBehavior(value NSWindowAnimationBehavior) {
	objc.Send[struct{}](w.ID, objc.Sel("setAnimationBehavior:"), value)
}

// A Boolean value that indicates whether the window’s document has been
// edited.
//
// # Discussion
//
// The value of this property is true when the window’s document has been
// edited; otherwise, false. Initially, by default, [NSWindow] objects are in
// the “not edited” state.
//
// You should set [DocumentEdited] to true every time the window’s document
// changes in such a way that it needs to be saved. Conversely, when the
// document is saved, you should set the property to true when the window’s
// document has been edited; otherwise, false. Then, before closing the window
// you can examine the value of the property to determine whether to allow the
// user a chance to save the document.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/isDocumentEdited
func (w NSWindow) DocumentEdited() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isDocumentEdited"))
	return rv
}
func (w NSWindow) SetDocumentEdited(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setDocumentEdited:"), value)
}

// The backing scale factor.
//
// # Discussion
//
// The value of this property is `2.0` for high-resolution scaled display
// modes, and `1.0` for all other cases.
//
// There are some scenarios where an application that is resolution-aware may
// want to reason on its own about the display environment it is running in.
//
// It is important to note that the value of this property does not represent
// anything concrete, such as pixel density or physical size, because it can
// vary based on the configured display mode. For example, the display may be
// in a mirrored configuration that is still high-resolution scaled, resulting
// in pixel geometry that may not match the native resolution of the display
// device.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/backingScaleFactor
func (w NSWindow) BackingScaleFactor() float64 {
	rv := objc.Send[float64](w.ID, objc.Sel("backingScaleFactor"))
	return rv
}

// The string that appears in the title bar of the window or the path to the
// represented file.
//
// # Discussion
//
// If the title has been set using [SetTitleWithRepresentedFilename], this
// property contains the file’s path. Setting this property also sets the
// title of the window’s miniaturized window.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/title
func (w NSWindow) Title() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}
func (w NSWindow) SetTitle(value string) {
	objc.Send[struct{}](w.ID, objc.Sel("setTitle:"), objc.String(value))
}

// A secondary line of text that appears in the title bar of the window.
//
// # Discussion
//
// When this property is an empty string, the system removes the subtitle from
// the window layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/subtitle
func (w NSWindow) Subtitle() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("subtitle"))
	return foundation.NSStringFromID(rv).String()
}
func (w NSWindow) SetSubtitle(value string) {
	objc.Send[struct{}](w.ID, objc.Sel("setSubtitle:"), objc.String(value))
}

// A value that indicates the visibility of the window’s title and title bar
// buttons.
//
// # Discussion
//
// By default, the value of this property is [NSWindowTitleVisible].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/titleVisibility-swift.property
func (w NSWindow) TitleVisibility() NSWindowTitleVisibility {
	rv := objc.Send[NSWindowTitleVisibility](w.ID, objc.Sel("titleVisibility"))
	return NSWindowTitleVisibility(rv)
}
func (w NSWindow) SetTitleVisibility(value NSWindowTitleVisibility) {
	objc.Send[struct{}](w.ID, objc.Sel("setTitleVisibility:"), value)
}

// The path to the file of the window’s represented file.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/representedFilename
func (w NSWindow) RepresentedFilename() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("representedFilename"))
	return foundation.NSStringFromID(rv).String()
}
func (w NSWindow) SetRepresentedFilename(value string) {
	objc.Send[struct{}](w.ID, objc.Sel("setRepresentedFilename:"), objc.String(value))
}

// The URL of the file the window represents.
//
// # Discussion
//
// When the URL specifies a path, the window shows an icon in its title bar,
// as described in the following table:
//
// [Table data omitted]
//
// You can customize the file icon in the title bar with the following code:
//
// When the URL identifies an existing file, the window’s title offers a
// pop-up menu showing the path components of the URL. (The user displays this
// menu by Command-clicking the title.) The behavior and contents of this menu
// can be controlled with [WindowShouldPopUpDocumentPathMenu].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/representedURL
func (w NSWindow) RepresentedURL() foundation.NSURL {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("representedURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (w NSWindow) SetRepresentedURL(value foundation.NSURL) {
	objc.Send[struct{}](w.ID, objc.Sel("setRepresentedURL:"), value)
}

// The screen the window is on.
//
// # Discussion
//
// The value of this property is the screen where most of the window is on; it
// is `nil` when the window is offscreen.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/screen
func (w NSWindow) Screen() INSScreen {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("screen"))
	return NSScreenFromID(objc.ID(rv))
}

// The deepest screen the window is on (it may be split over several screens).
//
// # Discussion
//
// The value of this property is `nil` when the window is offscreen.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/deepestScreen
func (w NSWindow) DeepestScreen() INSScreen {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("deepestScreen"))
	return NSScreenFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the window context should be updated
// when the screen profile changes or when the window moves to a different
// screen.
//
// # Discussion
//
// The value of this property is true when the window context should be
// updated when the ColorSync profile of the current screen changes or when a
// majority of the window is moved to a different screen whose profile is
// different than the previous screen; otherwise, false. The default value is
// false.
//
// After the window context is updated, the window is told to display itself.
// If you need to update offscreen caches for the window, you should register
// to receive the [didChangeScreenProfileNotification] notification.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/displaysWhenScreenProfileChanges
//
// [didChangeScreenProfileNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didChangeScreenProfileNotification
func (w NSWindow) DisplaysWhenScreenProfileChanges() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("displaysWhenScreenProfileChanges"))
	return rv
}
func (w NSWindow) SetDisplaysWhenScreenProfileChanges(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setDisplaysWhenScreenProfileChanges:"), value)
}

// A Boolean value that indicates whether the window is movable by clicking
// and dragging anywhere in its background.
//
// # Discussion
//
// The value of this property is true when the window is movable by clicking
// and dragging anywhere in its background; otherwise, false.
//
// A window with a style mask of [NSTexturedBackgroundWindowMask] is movable
// by background by default. Sheets and drawers cannot be movable by window
// background.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/isMovableByWindowBackground
func (w NSWindow) MovableByWindowBackground() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isMovableByWindowBackground"))
	return rv
}
func (w NSWindow) SetMovableByWindowBackground(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setMovableByWindowBackground:"), value)
}

// A Boolean value that indicates whether the window can be dragged by
// clicking in its title bar or background.
//
// # Discussion
//
// The value of this property is true if the window can be moved by the user;
// otherwise, false.
//
// When a window’s [Movable] property is false, the value of the
// [MovableByWindowBackground] property is ignored. When the value of
// [Movable] is false, the window can only be dragged between spaces in F8
// mode, and its relative screen position is always preserved. Note that a
// resizable window may still be resized, and the window frame may be changed
// programmatically. A nonmovable window will not be moved or resized by the
// system in response to a display reconfiguration. Applications may choose to
// enable application-controlled window dragging after disabling
// user-initiating dragging by handling the
// [MouseDown]/[MouseDragged]/[MouseUp] sequence in [SendEvent] in an
// [NSWindow] subclass.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/isMovable
func (w NSWindow) Movable() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isMovable"))
	return rv
}
func (w NSWindow) SetMovable(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setMovable:"), value)
}

// A Boolean value that indicates whether the window is released when it
// receives the `close` message.
//
// # Discussion
//
// The value of this property is true if the window is automatically released
// after being closed; false if it’s simply removed from the screen.
//
// The default for [NSWindow] is true; the default for [NSPanel] is false.
// Release when closed, however, is ignored for windows owned by window
// controllers. Another strategy for releasing an [NSWindow] object is to have
// its delegate autorelease it on receiving a [WindowShouldClose] message.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/isReleasedWhenClosed
func (w NSWindow) ReleasedWhenClosed() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isReleasedWhenClosed"))
	return rv
}
func (w NSWindow) SetReleasedWhenClosed(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setReleasedWhenClosed:"), value)
}

// A Boolean value that indicates whether the window is minimized.
//
// # Discussion
//
// The value of this property is true if the window is minimized; otherwise,
// false. A minimized window is removed from the screen and replaced by a
// image, icon, or button that represents it, called the counterpart.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/isMiniaturized
func (w NSWindow) Miniaturized() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isMiniaturized"))
	return rv
}

// The custom miniaturized window image of the window.
//
// # Discussion
//
// The miniaturized window image is the image displayed in the Dock when the
// window is minimized. If you did not assign a custom image to the window,
// the value of this property is `nil`.
//
// When the user minimizes the window, the Dock displays [MiniwindowImage] in
// the corresponding Dock tile, scaling it as needed to fit in the tile. If
// you do not specify a custom image using this property, the Dock creates one
// for you automatically.
//
// You can also set this property as needed to change the minimized window
// image. Typically, you would specify a custom image immediately prior to a
// window being minimized—when the system posts
// [willMiniaturizeNotification]. You can set this property while the window
// is minimized to update the current image in the Dock. However, you should
// not use this property to create complex animations in the Dock.
//
// Support for custom images is disabled by default. To enable support, set
// the [AppleDockIconEnabled] key to true when first registering your
// application’s user defaults. You must set this key prior to calling the
// `init` method of [NSApplication], which reads the current value of the key.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/miniwindowImage
//
// [willMiniaturizeNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/willMiniaturizeNotification
func (w NSWindow) MiniwindowImage() INSImage {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("miniwindowImage"))
	return NSImageFromID(objc.ID(rv))
}
func (w NSWindow) SetMiniwindowImage(value INSImage) {
	objc.Send[struct{}](w.ID, objc.Sel("setMiniwindowImage:"), value)
}

// The title displayed in the window’s minimized window.
//
// # Discussion
//
// A minimized window’s title usually reflects that of its full-size
// counterpart, abbreviated to fit if necessary. Although this property allows
// you to set the minimized window’s title explicitly, changing the
// full-size [NSWindow] object’s title (through [Title] or
// [SetTitleWithRepresentedFilename]) automatically changes the minimized
// window’s title as well.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/miniwindowTitle
func (w NSWindow) MiniwindowTitle() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("miniwindowTitle"))
	return foundation.NSStringFromID(rv).String()
}
func (w NSWindow) SetMiniwindowTitle(value string) {
	objc.Send[struct{}](w.ID, objc.Sel("setMiniwindowTitle:"), objc.String(value))
}

// The application’s Dock tile.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/dockTile
func (w NSWindow) DockTile() INSDockTile {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("dockTile"))
	return NSDockTileFromID(objc.ID(rv))
}

// Returns the bits per pixel for the specified window depth.
//
// See: https://developer.apple.com/documentation/appkit/nswindow/depth/bitsperpixel
func (w NSWindow) BitsPerPixel() int {
	rv := objc.Send[int](w.ID, objc.Sel("NSBitsPerPixelFromDepth"))
	return rv
}
func (w NSWindow) SetNSBitsPerPixelFromDepth(value int) {
	objc.Send[struct{}](w.ID, objc.Sel("setNSBitsPerPixelFromDepth:"), value)
}

// Returns the bits per sample for the specified window depth.
//
// See: https://developer.apple.com/documentation/appkit/nswindow/depth/bitspersample
func (w NSWindow) BitsPerSample() int {
	rv := objc.Send[int](w.ID, objc.Sel("NSBitsPerSampleFromDepth"))
	return rv
}
func (w NSWindow) SetNSBitsPerSampleFromDepth(value int) {
	objc.Send[struct{}](w.ID, objc.Sel("setNSBitsPerSampleFromDepth:"), value)
}

// Returns the name of the color space corresponding to the passed window
// depth.
//
// See: https://developer.apple.com/documentation/appkit/nswindow/depth/colorspacename
func (w NSWindow) ColorSpaceName() NSColorSpaceName {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("NSColorSpaceFromDepth"))
	return NSColorSpaceName(foundation.NSStringFromID(rv).String())
}
func (w NSWindow) SetNSColorSpaceFromDepth(value NSColorSpaceName) {
	objc.Send[struct{}](w.ID, objc.Sel("setNSColorSpaceFromDepth:"), objc.String(string(value)))
}

// Returns the number of color components in the specified color space.
//
// See: https://developer.apple.com/documentation/appkit/nscolorspacename/numberofcolorcomponents
func (w NSWindow) NumberOfColorComponents() int {
	rv := objc.Send[int](w.ID, objc.Sel("NSNumberOfColorComponents"))
	return rv
}
func (w NSWindow) SetNSNumberOfColorComponents(value int) {
	objc.Send[struct{}](w.ID, objc.Sel("setNSNumberOfColorComponents:"), value)
}

// Returns whether the specified window depth is planar.
//
// See: https://developer.apple.com/documentation/appkit/nswindow/depth/isplanar
func (w NSWindow) IsPlanar() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("NSPlanarFromDepth"))
	return rv
}
func (w NSWindow) SetNSPlanarFromDepth(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setNSPlanarFromDepth:"), value)
}

// A Boolean value that indicates if the window has a close box.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/hasCloseBox
func (w NSWindow) HasCloseBox() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasCloseBox"))
	return rv
}

// A Boolean value that indicates if the window has a title bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/hasTitleBar
func (w NSWindow) HasTitleBar() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasTitleBar"))
	return rv
}

// A Boolean value that indicates whether the window is a modal panel.
//
// # Discussion
//
// This property is key-value coding compliant.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/isModalPanel
func (w NSWindow) ModalPanel() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isModalPanel"))
	return rv
}

// A Boolean value that indicates whether the window is a floating panel.
//
// # Discussion
//
// This property is key-value coding compliant.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/isFloatingPanel
func (w NSWindow) FloatingPanel() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isFloatingPanel"))
	return rv
}

// A Boolean value that indicates whether the window allows zooming.
//
// # Discussion
//
// This property is key-value coding compliant.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/isZoomable
func (w NSWindow) Zoomable() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isZoomable"))
	return rv
}

// A Boolean value that indicates if the user can resize the window.
//
// # Discussion
//
// This property is key-value coding compliant.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/isResizable
func (w NSWindow) Resizable() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isResizable"))
	return rv
}

// A Boolean value that indicates whether the window can minimize.
//
// # Discussion
//
// This property is key-value coding compliant.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/isMiniaturizable
func (w NSWindow) Miniaturizable() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isMiniaturizable"))
	return rv
}

// The zero-based position of the window, based on its order from front to
// back among all visible application windows.
//
// # Discussion
//
// If you set this property to an index that’s out of range, the system sets
// the position to the nearest value that’s in range.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/orderedIndex
func (w NSWindow) OrderedIndex() int {
	rv := objc.Send[int](w.ID, objc.Sel("orderedIndex"))
	return rv
}
func (w NSWindow) SetOrderedIndex(value int) {
	objc.Send[struct{}](w.ID, objc.Sel("setOrderedIndex:"), value)
}

// # Discussion
//
// The frame to use when cascading or sizing a new window based on the
// receiver’s position or size. This may be different from `frame` when the
// receiver is positioned by the system.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/cascadingReferenceFrame
func (w NSWindow) CascadingReferenceFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](w.ID, objc.Sel("cascadingReferenceFrame"))
	return corefoundation.CGRect(rv)
}

// Name of an exception that occurs when you pass an invalid argument to a
// method, such as a `nil` pointer where a non-`nil` object is required.
//
// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
func (w NSWindow) InvalidArgumentException() foundation.NSString {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("invalidArgumentException"))
	return foundation.NSStringFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the window’s resize indicator is
// visible.
//
// # Discussion
//
// The value of this property is true when the window’s resize indicator is
// visible; otherwise, false. This property does not affect whether the window
// is resizable.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/showsResizeIndicator
func (w NSWindow) ShowsResizeIndicator() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("showsResizeIndicator"))
	return rv
}
func (w NSWindow) SetShowsResizeIndicator(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setShowsResizeIndicator:"), value)
}

// The Carbon window reference associated with the window, creating one if
// necessary.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/windowRef
func (w NSWindow) WindowRef() WindowRef {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("windowRef"))
	return WindowRef(rv)
}
func (w NSWindow) SetWindowRef(value WindowRef) {
	objc.Send[struct{}](w.ID, objc.Sel("setWindowRef:"), value)
}

// Returns the default depth limit for instances of [NSWindow].
//
// # Return Value
//
// The default depth limit for instances of [NSWindow], determined by the
// depth of the deepest screen level available to the window server.
//
// # Discussion
//
// The value returned can be examined with the Application Kit functions
// [NSPlanarFromDepth], [NSColorSpaceFromDepth], [NSBitsPerSampleFromDepth],
// and [NSBitsPerPixelFromDepth].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/defaultDepthLimit
func (_NSWindowClass NSWindowClass) DefaultDepthLimit() NSWindowDepth {
	rv := objc.Send[NSWindowDepth](objc.ID(_NSWindowClass.class), objc.Sel("defaultDepthLimit"))
	return NSWindowDepth(rv)
}

// A Boolean value that indicates whether the app can automatically organize
// windows into tabs.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/allowsAutomaticWindowTabbing
func (_NSWindowClass NSWindowClass) AllowsAutomaticWindowTabbing() bool {
	rv := objc.Send[bool](objc.ID(_NSWindowClass.class), objc.Sel("allowsAutomaticWindowTabbing"))
	return rv
}
func (_NSWindowClass NSWindowClass) SetAllowsAutomaticWindowTabbing(value bool) {
	objc.Send[struct{}](objc.ID(_NSWindowClass.class), objc.Sel("setAllowsAutomaticWindowTabbing:"), value)
}

// A value that indicates the user’s preference for window tabbing.
//
// # Discussion
//
// This value indicates the user’s preference for window tabbing as set in
// System Preferences. Check this preference any time you create a new window.
// For a list of possible values, see [NSWindow.UserTabbingPreference].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/userTabbingPreference-swift.type.property
//
// [NSWindow.UserTabbingPreference]: https://developer.apple.com/documentation/AppKit/NSWindow/UserTabbingPreference-swift.enum
func (_NSWindowClass NSWindowClass) UserTabbingPreference() NSWindowUserTabbingPreference {
	rv := objc.Send[NSWindowUserTabbingPreference](objc.ID(_NSWindowClass.class), objc.Sel("userTabbingPreference"))
	return NSWindowUserTabbingPreference(rv)
}

// The longest time duration possible.
//
// See: https://developer.apple.com/documentation/appkit/nsevent/foreverduration
func (_NSWindowClass NSWindowClass) ForeverDuration() float64 {
	rv := objc.Send[float64](objc.ID(_NSWindowClass.class), objc.Sel("NSEventDurationForever"))
	return rv
}

// An
//
// See: https://developer.apple.com/documentation/appkit/nswindow/oldcolorspaceuserinfokey
func (_NSWindowClass NSWindowClass) OldColorSpaceUserInfoKey() string {
	rv := objc.Send[objc.ID](objc.ID(_NSWindowClass.class), objc.Sel("NSBackingPropertyOldColorSpaceKey"))
	return foundation.NSStringFromID(rv).String()
}

// An NSNumber containing the old scale factor.
//
// See: https://developer.apple.com/documentation/appkit/nswindow/oldscalefactoruserinfokey
func (_NSWindowClass NSWindowClass) OldScaleFactorUserInfoKey() string {
	rv := objc.Send[objc.ID](objc.ID(_NSWindowClass.class), objc.Sel("NSBackingPropertyOldScaleFactorKey"))
	return foundation.NSStringFromID(rv).String()
}

// Protocol methods for NSAppearanceCustomization

// The appearance of the receiver, in an [NSAppearance] object.
//
// # Discussion
//
// The default value for this property is `nil`, which means that the receiver
// uses the appearance it inherits from the nearest ancestor that has set an
// appearance. When you set `appearance` to a non-`nil` value, the receiver
// and the views it contains use the specified appearance.
//
// See: https://developer.apple.com/documentation/AppKit/NSAppearanceCustomization/appearance
func (o NSWindow) SetAppearance(value INSAppearance) {
	objc.Send[struct{}](o.ID, objc.Sel("setAppearance:"), value)
}

// Protocol methods for NSMenuItemValidation

// Protocol methods for NSUserInterfaceItemIdentification

// A string that identifies the user interface item.
//
// # Discussion
//
// Identifiers are used during window restoration operations to uniquely
// identify the windows of the application. You can set the value of this
// string programmatically or in Interface Builder. If you create an item in
// Interface Builder and do not set a value for this string, a unique value is
// created for the item when the nib file is loaded. For programmatically
// created views, you typically set this value after creating the item but
// before adding it to a window.
//
// You should not change the value of a window’s identifier after adding any
// views to the window. For views and controls in a window, the value you
// specify for this string must be unique on a per-window basis.
//
// The slash (`/`), backslash (`\`), or colon (`:`) characters are reserved
// and must not be used in your custom identifiers. Similarly, Apple reserves
// all identifiers beginning with an underscore (`_`) character. Applications
// and frameworks should use a consistent prefix for their identifiers to
// avoid collisions with other frameworks. For a list of prefixes used by the
// system frameworks, see [OS X Frameworks] in [Mac Technology Overview].
//
// If you are subclassing a class from one of the system frameworks, do not
// override the accessor methods of this protocol.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceItemIdentification/identifier
//
// [Mac Technology Overview]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/OSX_Technology_Overview/About/About.html#//apple_ref/doc/uid/TP40001067
// [OS X Frameworks]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/OSX_Technology_Overview/SystemFrameworks/SystemFrameworks.html#//apple_ref/doc/uid/TP40001067-CH210
func (o NSWindow) SetIdentifier(value NSUserInterfaceItemIdentifier) {
	objc.Send[struct{}](o.ID, objc.Sel("setIdentifier:"), objc.String(string(value)))
}

// Protocol methods for NSUserInterfaceValidations

// BeginSheet is a synchronous wrapper around [NSWindow.BeginSheetCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w NSWindow) BeginSheet(ctx context.Context, sheetWindow INSWindow) (NSModalResponse, error) {
	done := make(chan NSModalResponse, 1)
	w.BeginSheetCompletionHandler(sheetWindow, func(val NSModalResponse) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return *new(NSModalResponse), ctx.Err()
	}
}

// BeginCriticalSheet is a synchronous wrapper around [NSWindow.BeginCriticalSheetCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w NSWindow) BeginCriticalSheet(ctx context.Context, sheetWindow INSWindow) (NSModalResponse, error) {
	done := make(chan NSModalResponse, 1)
	w.BeginCriticalSheetCompletionHandler(sheetWindow, func(val NSModalResponse) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return *new(NSModalResponse), ctx.Err()
	}
}

// TransferWindowSharingToWindow is a synchronous wrapper around [NSWindow.TransferWindowSharingToWindowCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w NSWindow) TransferWindowSharingToWindow(ctx context.Context, window INSWindow) error {
	done := make(chan error, 1)
	w.TransferWindowSharingToWindowCompletionHandler(window, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RequestSharingOfWindow is a synchronous wrapper around [NSWindow.RequestSharingOfWindowCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w NSWindow) RequestSharingOfWindow(ctx context.Context, window INSWindow) error {
	done := make(chan error, 1)
	w.RequestSharingOfWindowCompletionHandler(window, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RequestSharingOfWindowUsingPreviewTitle is a synchronous wrapper around [NSWindow.RequestSharingOfWindowUsingPreviewTitleCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w NSWindow) RequestSharingOfWindowUsingPreviewTitle(ctx context.Context, image INSImage, title string) error {
	done := make(chan error, 1)
	w.RequestSharingOfWindowUsingPreviewTitleCompletionHandler(image, title, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
