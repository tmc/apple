// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
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
//   - [NSWindow.DisplayLinkWithTargetSelector]
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
//   - [NSWindow.HasActiveWindowSharingSession]
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
//   - [NSWindow.SetBitsPerPixel]
//   - [NSWindow.BitsPerSample]: Returns the bits per sample for the specified window depth.
//   - [NSWindow.SetBitsPerSample]
//   - [NSWindow.ColorSpaceName]: Returns the name of the color space corresponding to the passed window depth.
//   - [NSWindow.SetColorSpaceName]
//   - [NSWindow.NumberOfColorComponents]: Returns the number of color components in the specified color space.
//   - [NSWindow.SetNumberOfColorComponents]
//   - [NSWindow.IsPlanar]: Returns whether the specified window depth is planar.
//   - [NSWindow.SetIsPlanar]
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
//   - [INSWindow.DisplayLinkWithTargetSelector]
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
//   - [INSWindow.HasActiveWindowSharingSession]
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
//   - [INSWindow.SetBitsPerPixel]
//   - [INSWindow.BitsPerSample]: Returns the bits per sample for the specified window depth.
//   - [INSWindow.SetBitsPerSample]
//   - [INSWindow.ColorSpaceName]: Returns the name of the color space corresponding to the passed window depth.
//   - [INSWindow.SetColorSpaceName]
//   - [INSWindow.NumberOfColorComponents]: Returns the number of color components in the specified color space.
//   - [INSWindow.SetNumberOfColorComponents]
//   - [INSWindow.IsPlanar]: Returns whether the specified window depth is planar.
//   - [INSWindow.SetIsPlanar]
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
	NSAccessibilityElementProtocol
	NSAccessibilityProtocol
	NSAppearanceCustomization
	NSMenuItemValidation
	NSUserInterfaceItemIdentification
	NSUserInterfaceValidations

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
	DisplayLinkWithTargetSelector(target objectivec.IObject, selector objc.SEL) objectivec.IObject

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
	NextEventMatchingMaskUntilDateInModeDequeue(mask NSEventMask, expiration foundation.INSDate, mode foundation.NSString, deqFlag bool) INSEvent
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
	RepresentedURL() foundation.INSURL
	SetRepresentedURL(value foundation.INSURL)

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
	DataWithEPSInsideRect(rect corefoundation.CGRect) foundation.INSData
	// Returns PDF data that draws the region of the window within a given rectangle.
	DataWithPDFInsideRect(rect corefoundation.CGRect) foundation.INSData

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
	SetBitsPerPixel(value int)
	// Returns the bits per sample for the specified window depth.
	BitsPerSample() int
	SetBitsPerSample(value int)
	// Returns the name of the color space corresponding to the passed window depth.
	ColorSpaceName() NSColorSpaceName
	SetColorSpaceName(value NSColorSpaceName)
	// Returns the number of color components in the specified color space.
	NumberOfColorComponents() int
	SetNumberOfColorComponents(value int)
	// Returns whether the specified window depth is planar.
	IsPlanar() bool
	SetIsPlanar(value bool)
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

	// The location of the window’s backing store.
	BackingLocation() objectivec.IObject
	SetBackingLocation(value objectivec.IObject)
	// The collection of drawers associated with the window.
	Drawers() objc.ID
	SetDrawers(value objc.ID)
	// The graphics context associated with the window for the current thread.
	GraphicsContext() INSGraphicsContext
	SetGraphicsContext(value INSGraphicsContext)
	// Name of an exception that occurs when you pass an invalid argument to a method, such as a `nil` pointer where a non-`nil` object is required.
	InvalidArgumentException() foundation.NSString
	// A Boolean value that indicates whether the window automatically displays views that need to be displayed.
	IsAutodisplay() bool
	SetIsAutodisplay(value bool)
	// A Boolean value that indicates whether the window’s flushing ability is disabled.
	IsFlushWindowDisabled() bool
	SetIsFlushWindowDisabled(value bool)
	// A Boolean value that indicates whether the window device the window manages is freed when it’s removed from the screen list.
	IsOneShot() bool
	SetIsOneShot(value bool)
	// A Boolean value that indicates the preferred location for the window’s backing store.
	PreferredBackingLocation() objectivec.IObject
	SetPreferredBackingLocation(value objectivec.IObject)
	// A Boolean value that indicates whether the window’s resize indicator is visible.
	ShowsResizeIndicator() bool
	SetShowsResizeIndicator(value bool)
	// The Carbon window reference associated with the window, creating one if necessary.
	WindowRef() WindowRef
	SetWindowRef(value WindowRef)
	// Returns the animation that should be performed for the specified key.
	AnimationForKey(key NSAnimatablePropertyKey) objectivec.IObject
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
// //
// [NSWindow.StyleMask]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
//
// backingStoreType: Specifies how the drawing done in the window is buffered by the window
// device, and possible values are described in [NSWindow.BackingStoreType].
// //
// [NSWindow.BackingStoreType]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType
//
// flag: Specifies whether the window server creates a window device for the window
// immediately. When [true], the window server defers creating the window
// device until the window is moved onscreen. All display messages sent to the
// window or its views are postponed until the window is created, just before
// it’s moved onscreen.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
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
// //
// [NSWindow.StyleMask]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
//
// backingStoreType: Specifies how the drawing done in the window is buffered by the window
// device; possible values are described in [NSWindow.BackingStoreType].
// //
// [NSWindow.BackingStoreType]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType
//
// flag: Specifies whether the window server creates a window device for the window
// immediately. When [true], the window server defers creating the window
// device until the window is moved onscreen. All display messages sent to the
// window or its views are postponed until the window is created, just before
// it’s moved onscreen.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
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
// `contentViewController``XCUIElementTypeView`.
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
// `contentViewController``XCUIElementTypeView`). The newly created window has
// [ReleasedWhenClosed] set to [false], and it must be explicitly retained to
// keep the window instance alive.
//
// [false]: https://developer.apple.com/documentation/Swift/false
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
// For historical reasons, contrary to normal memory management policy `` does
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
// //
// [NSWindow.StyleMask]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
//
// backingStoreType: Specifies how the drawing done in the window is buffered by the window
// device, and possible values are described in [NSWindow.BackingStoreType].
// //
// [NSWindow.BackingStoreType]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType
//
// flag: Specifies whether the window server creates a window device for the window
// immediately. When [true], the window server defers creating the window
// device until the window is moved onscreen. All display messages sent to the
// window or its views are postponed until the window is created, just before
// it’s moved onscreen.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
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
// //
// [NSWindow.StyleMask]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
//
// backingStoreType: Specifies how the drawing done in the window is buffered by the window
// device; possible values are described in [NSWindow.BackingStoreType].
// //
// [NSWindow.BackingStoreType]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType
//
// flag: Specifies whether the window server creates a window device for the window
// immediately. When [true], the window server defers creating the window
// device until the window is moved onscreen. All display messages sent to the
// window or its views are postponed until the window is created, just before
// it’s moved onscreen.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
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
// View menu with `` as the action, and `nil` as the target.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/toggleFullScreen(_:)
func (w NSWindow) ToggleFullScreen(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("toggleFullScreen:"), sender)
}
// Sets a Boolean value that indicates whether the window’s depth limit can
// change to match the depth of the screen it’s on.
//
// flag: [true] if the window’s depth can change; otherwise, [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [true] when the window auto-recalculates the given border’s thickness;
// otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/autorecalculatesContentBorderThickness(for:)
func (w NSWindow) AutorecalculatesContentBorderThicknessForEdge(edge foundation.NSRectEdge) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("autorecalculatesContentBorderThicknessForEdge:"), edge)
	return rv
}
// Specifies whether the window calculates the thickness of a given border
// automatically.
//
// flag: If [true], the window calculates the thickness of the edge automatically;
// if [false], it does not.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// In a nontextured window calling `` passing [NSMaxYEdge] will raise an
// exception (in a nontextured window, it’s only valid to set the content
// border thickness of the bottom edge). It is only valid to set the content
// border thickness of the top edge in a textured window.
// 
// Typically, if you call ``, you should also call `NO `.
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
// In a nontextured window calling `` passing [NSMaxYEdge] will raise an
// exception (in a nontextured window, it’s only valid to set the content
// border thickness of the bottom edge). It is only valid to set the content
// border thickness of the top edge in a textured window.
// 
// Typically, if you call ``, you should also call `NO `.
// 
// The `contentBorder` does not include the title bar or toolbar, so a
// textured window that just wants the gradient in the title bar and toolbar
// should have a `thickness` of `0` for [NSMaxYEdge].
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/setContentBorderThickness(_:for:)
func (w NSWindow) SetContentBorderThicknessForEdge(thickness float64, edge foundation.NSRectEdge) {
	objc.Send[objc.ID](w.ID, objc.Sel("setContentBorderThickness:forEdge:"), thickness, edge)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/displayLink(target:selector:)
func (w NSWindow) DisplayLinkWithTargetSelector(target objectivec.IObject, selector objc.SEL) objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("displayLinkWithTarget:selector:"), target, selector)
	return objectivec.Object{ID: rv}
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
// [NSWindow.StyleMask]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/contentRect(forFrameRect:)
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
// [NSWindow.StyleMask]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/frameRect(forContentRect:)
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
_block1, _cleanup1 := NewModalResponseBlock(handler)
	defer _cleanup1()
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
_block1, _cleanup1 := NewModalResponseBlock(handler)
	defer _cleanup1()
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
// //
// [NSZeroPoint]: https://developer.apple.com/documentation/Foundation/NSZeroPoint
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
// When [true] the window sends a [DisplayIfNeeded] message down its view
// hierarchy, thus redrawing all views.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
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
// When [true] the window sends a [DisplayIfNeeded] message down its view
// hierarchy, thus redrawing all views.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// animateFlag: Specifies whether the window performs a smooth resize. [true] to perform
// the animation, whose duration is specified by [AnimationResizeTime].
// //
// [true]: https://developer.apple.com/documentation/Swift/true
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
// the window’s version. [WindowShouldZoomToFrame] returns [false] if
// zooming isn’t currently allowed. - If the window currently allows
// zooming, sets the new frame.
//
// [false]: https://developer.apple.com/documentation/Swift/false
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
// place: - [NSWindow.OrderingMode.out]: The window is removed from the screen list
// and `otherWin` is ignored.
// 
// - [NSWindow.OrderingMode.above]: The window is ordered immediately in front
// of the window whose window number is `otherWin` -
// [NSWindow.OrderingMode.below]: The window is placed immediately behind the
// window represented by `otherWin`.
// //
// [NSWindow.OrderingMode.above]: https://developer.apple.com/documentation/AppKit/NSWindow/OrderingMode/above
// [NSWindow.OrderingMode.below]: https://developer.apple.com/documentation/AppKit/NSWindow/OrderingMode/below
// [NSWindow.OrderingMode.out]: https://developer.apple.com/documentation/AppKit/NSWindow/OrderingMode/out
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
// [true] when `name` is read and the frame is set successfully; otherwise,
// [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// force: [true] to use [SetFrameUsingName] on a non-resizable window; [false] to
// fail on a non-resizable window.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// [true] when `name` is read and the frame is set successfully; otherwise,
// [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [UserDefaults]: https://developer.apple.com/documentation/Foundation/UserDefaults
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/saveFrame(usingName:)
func (w NSWindow) SaveFrameUsingName(name NSWindowFrameAutosaveName) {
	objc.Send[objc.ID](w.ID, objc.Sel("saveFrameUsingName:"), objc.String(string(name)))
}
// Sets the name AppKit uses to automatically save the window’s frame
// rectangle data in the defaults system.
//
// name: The name to use for the window’s frame rectangle.
//
// # Return Value
// 
// [true] if the specified name is acceptable and another window isn’t
// already using it; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/setFrameAutosaveName(_:)
func (w NSWindow) SetFrameAutosaveName(name NSWindowFrameAutosaveName) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("setFrameAutosaveName:"), objc.String(string(name)))
	return rv
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
// [didBecomeKeyNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didBecomeKeyNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/becomeKey()
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
// [didResignKeyNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didResignKeyNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/resignKey()
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
// [didBecomeMainNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didBecomeMainNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/becomeMain()
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
// [didResignMainNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didResignMainNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/resignMain()
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
// place: - [NSWindow.OrderingMode.above]: `childWin` is ordered immediately in front
// of the window.
// 
// - [NSWindow.OrderingMode.below]: `childWin` is ordered immediately behind
// the window.
// //
// [NSWindow.OrderingMode.above]: https://developer.apple.com/documentation/AppKit/NSWindow/OrderingMode/above
// [NSWindow.OrderingMode.below]: https://developer.apple.com/documentation/AppKit/NSWindow/OrderingMode/below
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
// createFlag: If [true], creates a field editor if one doesn’t exist; if [false], does
// not create a field editor.
// 
// A freshly created [NSWindow] object doesn’t have a field editor. After a
// field editor has been created for a window, the `createFlag` argument is
// ignored. By passing [false] for `createFlag` and testing the return value,
// however, you can predicate an action on the existence of the field editor.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// is [false] and if the field editor doesn’t exist.
//
// [false]: https://developer.apple.com/documentation/Swift/false
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
// status even if its [ResignFirstResponder] method returns [false]. This
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
// [false]: https://developer.apple.com/documentation/Swift/false
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
// parameter is [true].
//
// [NSEventTrackingRunLoopMode]: https://developer.apple.com/documentation/AppKit/NSEventTrackingRunLoopMode
// [distantFuture]: https://developer.apple.com/documentation/Foundation/NSDate/distantFuture
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/nextEvent(matching:)
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
// deqFlag: [true] to remove the returned event from the event queue; [false] to leave
// the returned event in the queue.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// The next event whose mask matches the specified `mask`; otherwise, `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/nextEvent(matching:until:inMode:dequeue:)
func (w NSWindow) NextEventMatchingMaskUntilDateInModeDequeue(mask NSEventMask, expiration foundation.INSDate, mode foundation.NSString, deqFlag bool) INSEvent {
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
// flag: [true] to place the event in the front of the queue; [false] to place it in
// the back.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [true] when the operation is successful; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If `responder` isn’t already the first responder, this method first sends
// a [ResignFirstResponder] message to the object that is the first responder.
// If that object refuses to resign, it remains the first responder, and this
// method immediately returns [false]. If the current first responder resigns,
// this method sends a [BecomeFirstResponder] message to `responder`. If
// `responder` does not accept first responder status, the [NSWindow] object
// becomes first responder; in this case, the method returns [true] even if
// `responder` refuses first responder status.
// 
// If `responder` is `nil`, this method still sends [ResignFirstResponder] to
// the current first responder. If the current first responder refuses to
// resign, it remains the first responder and this method immediately returns
// [false]. If the current first responder returns [true] from
// [ResignFirstResponder], the window is made its own first responder and this
// method returns [true].
// 
// The Application Kit framework uses this method to alter the first responder
// in response to mouse-down events; you can also use it to explicitly set the
// first responder from within your program. The `responder` object is
// typically an [NSView] object in the window’s view hierarchy. If this
// method is called explicitly, first send [AcceptsFirstResponder] to
// `responder`, and do not call [FirstResponder] if [AcceptsFirstResponder]
// returns [false].
// 
// Use [InitialFirstResponder] to the set the first responder to be used when
// the window is brought onscreen for the first time.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// Sends the [previousValidKeyView] message to `view` and, if that message
// returns an [NSView] object, invokes [FirstResponder] with the returned
// object.
//
// [previousValidKeyView]: https://developer.apple.com/documentation/AppKit/NSView/previousValidKeyView
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
// Sends the [nextValidKeyView] message to `view` and, if that message returns
// an [NSView] object, invokes [FirstResponder] with the returned object.
//
// [nextValidKeyView]: https://developer.apple.com/documentation/AppKit/NSView/nextValidKeyView
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
// [previousValidKeyView] method of [NSView] returns - The
// [InitialFirstResponder] designates as the window’s initial first
// responder if it returns [true] to an [AcceptsFirstResponder] message -
// Otherwise, the initial first responder’s previous valid key view, which
// may be `nil`
//
// [previousValidKeyView]: https://developer.apple.com/documentation/AppKit/NSView/previousValidKeyView
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [nextValidKeyView] method of [NSView] returns - The object
// [InitialFirstResponder] designates as the window’s initial first
// responder if it returns [true] to an [AcceptsFirstResponder] message -
// Otherwise, the initial first responder’s next valid key view, which may
// be `nil`
//
// [nextValidKeyView]: https://developer.apple.com/documentation/AppKit/NSView/nextValidKeyView
// [true]: https://developer.apple.com/documentation/Swift/true
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
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/transferWindowSharing(to:completionHandler:)
func (w NSWindow) TransferWindowSharingToWindowCompletionHandler(window INSWindow, completionHandler ErrorHandler) {
_block1, _cleanup1 := NewErrorBlock(completionHandler)
	defer _cleanup1()
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
// add [EventMaskPressure] to the event mask. This method returns when
// tracking terminates.
// 
// Each event is removed from the event queue and then passed to the tracking
// handler. If a matching event does not exist in the event queue, the main
// thread blocks in the specified runloop mode until an event of the requested
// type is received or the specified timeout expires. If the timeout expires,
// the tracking handler is called with a `nil` event (a negative timeout is
// interpreted as `0`). Use [NSEventDurationForever] to prevent timing out.
// Tracking continues until you set `stop` to [true]. Note that calls to
// [NextEventMatchingMask] are allowed inside the `trackingHandler` block.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/trackEvents(matching:timeout:mode:handler:)
func (w NSWindow) TrackEventsMatchingMaskTimeoutModeHandler(mask NSEventMask, timeout float64, mode foundation.NSString, trackingHandler EventHandler) {
_block3, _cleanup3 := NewEventBlock(trackingHandler)
	defer _cleanup3()
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
// [didUpdateNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didUpdateNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/update()
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
// //
// [AlignmentOptions]: https://developer.apple.com/documentation/Foundation/AlignmentOptions
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
// [NSIntegralRectWithOptions(_:_:)]: https://developer.apple.com/documentation/Foundation/NSIntegralRectWithOptions(_:_:)
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/backingAlignedRect(_:options:)
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
// [WindowShouldClose] method returns [false], the window doesn’t close. If
// neither the window nor the delegate implement [WindowShouldClose], or it
// returns [true], this method invokes [Close] to close the window.
// 
// If the window doesn’t have a close button or can’t close (for example,
// if the delegate replies [false] to a [WindowShouldClose] message), the
// system emits the alert sound.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [willCloseNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/willCloseNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/close()
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
func (w NSWindow) DataWithEPSInsideRect(rect corefoundation.CGRect) foundation.INSData {
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
func (w NSWindow) DataWithPDFInsideRect(rect corefoundation.CGRect) foundation.INSData {
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
// //
// [NSLayoutConstraint.Orientation]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Orientation
//
// # Return Value
// 
// Returns the layout attribute.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/anchorAttribute(for:)
func (w NSWindow) AnchorAttributeForOrientation(orientation NSLayoutConstraintOrientation) NSLayoutAttribute {
	rv := objc.Send[NSLayoutAttribute](w.ID, objc.Sel("anchorAttributeForOrientation:"), orientation)
	return NSLayoutAttribute(rv)
}
// Sets the part of the window that stays stationary during constraint-based
// layout.
//
// attr: The layout attribute. [NSLayoutConstraint.Attribute] specifies the possible
// values.
// //
// [NSLayoutConstraint.Attribute]: https://developer.apple.com/documentation/UIKit/NSLayoutConstraint/Attribute
//
// orientation: The window drag orientation. [NSLayoutConstraint.Orientation] specifies the
// possible values.
// //
// [NSLayoutConstraint.Orientation]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Orientation
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/setAnchorAttribute(_:for:)
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
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/beginDraggingSession(items:event:source:)
func (w NSWindow) BeginDraggingSessionWithItemsEventSource(items []NSDraggingItem, event INSEvent, source NSDraggingSource) INSDraggingSession {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("beginDraggingSessionWithItems:event:source:"), objectivec.IObjectSliceToNSArray(items), event, source)
	return NSDraggingSessionFromID(rv)
}
//
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
_block1, _cleanup1 := NewErrorBlock(completionHandler)
	defer _cleanup1()
	objc.Send[objc.ID](w.ID, objc.Sel("requestSharingOfWindow:completionHandler:"), window, _block1)
}
//
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
_block2, _cleanup2 := NewErrorBlock(completionHandler)
	defer _cleanup2()
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
// [CAAnimation]: https://developer.apple.com/documentation/QuartzCore/CAAnimation
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
// [CALayer]: https://developer.apple.com/documentation/QuartzCore/CALayer
// [action(forKey:)]: https://developer.apple.com/documentation/QuartzCore/CALayer/action(forKey:)
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/animation(forKey:)
func (w NSWindow) AnimationForKey(key NSAnimatablePropertyKey) objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("animationForKey:"), objc.String(string(key)))
	return objectivec.Object{ID: rv}
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
// Implemented to override the default action of enabling or disabling a
// specific menu item.
//
// menuItem: An [NSMenuItem] object that represents the menu item.
//
// # Return Value
// 
// [true] to enable `menuItem`, [false] to disable it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [true] if the user interface item should be enabled, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// //
// [NSWindow.NumberListOptions]: https://developer.apple.com/documentation/AppKit/NSWindow/NumberListOptions
//
// # Return Value
// 
// An array of window numbers for all visible windows satisfying the specified
// options. (Windows on the active space are returned in z-order; that is,
// front to back.)
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/windowNumbers(options:)
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
// //
// [NSWindow.StyleMask]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
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
// //
// [NSWindow.StyleMask]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
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
// //
// [NSWindow.StyleMask]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
//
// # Return Value
// 
// The minimum width of the window’s frame, using `style`, in order to
// display `title`.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/minFrameWidth(withTitle:styleMask:)
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
// //
// [NSWindow.StyleMask]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
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
// [CAAnimation]: https://developer.apple.com/documentation/QuartzCore/CAAnimation
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
// [CAAnimation]: https://developer.apple.com/documentation/QuartzCore/CAAnimation
// [NSAnimationTriggerOrderIn]: https://developer.apple.com/documentation/AppKit/NSAnimationTriggerOrderIn
// [NSAnimationTriggerOrderOut]: https://developer.apple.com/documentation/AppKit/NSAnimationTriggerOrderOut
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/defaultAnimation(forKey:)
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
// The value of this property is [true] if the window is able to receive
// keyboard and mouse events even when some other window is being run modally;
// otherwise, [false]. By default, the [NSWindow] value of this property is
// [false]. Only subclasses of [NSPanel] should override this default.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// The value of this property is [true] if the window can hide when its
// application becomes hidden (during execution of the
// `NSApplication```NSApplication/hide(_:)`` method); otherwise, [false]. By
// default, the value of the property is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// The value of this property is [true] if the window is on the currently
// active space; otherwise, [false]. For visible windows, this property
// indicates whether the window is currently visible on the active space. For
// nonvisible windows, it indicates whether ordering the window onscreen would
// cause it to be on the active space.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// The value of this property is [true] if the window is removed from the
// screen when its application is deactivated; [false] if it remains onscreen.
// The default value for [NSWindow] is [false]; the default value for
// [NSPanel] is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [NSWindow.CollectionBehavior]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/collectionBehavior-swift.property
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
// The value of this property is [true] when the window is opaque; otherwise,
// [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// The value of this property is [true] when the window has a shadow;
// otherwise, [false]. If you change the value of this property, the window
// shadow is invalidated, forcing the window shadow to be recomputed.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// The value of this property is [true] if the window prevents application
// termination when modal; otherwise, [false]. The default value is [true].
// 
// Usually, application termination is prevented when a modal window or sheet
// is open, without consulting the application delegate. Some windows may wish
// not to prevent termination, however. Setting this property to [false]
// overrides the default behavior and allows termination to proceed even if
// the window is open, either through the sudden termination path if enabled,
// or after consulting the application delegate.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [NSApp]: https://developer.apple.com/documentation/AppKit/NSApp
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
// ([NSWindow.Depth.twentyfourBitRGB] is the default).
//
// [NSWindow.Depth.twentyfourBitRGB]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/twentyfourBitRGB
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
// The value of this property is [true] when the window has a dynamic depth
// limit; otherwise, [false]. When the value of [HasDynamicDepthLimit] is
// [false], the window uses either its preset depth limit or the default depth
// limit. A different, and non-dynamic, depth limit can be set using the
// [DepthLimit] property.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// The value of this property is [true] when the window can be displayed at
// the login window; otherwise, [false]. By default, the value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [NSWindow.SharingType]: https://developer.apple.com/documentation/AppKit/NSWindow/SharingType-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/sharingType-swift.property
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
// [NSWindow.BackingStoreType]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/backingType
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
// The value of this property is [true] if the window has ever run as a modal
// sheet; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// The value of this property is [true] if the window is in a zoomed state;
// otherwise, [false].
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
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// `increments``XCUIElementTypeWidth` and `increments``XCUIElementTypeHeight`,
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
// The value of this property is [true] if the window tries to optimize live
// resize operations by preserving the content of views that have not moved;
// otherwise, [false]. By default, live-resize optimization is turned on.
// 
// When live-resize optimization is active, the window redraws only those
// views that moved (or do not support this optimization) during a live resize
// operation. You might consider disabling this optimization for the window if
// none of the window’s contained views can take advantage of it. Disabling
// the optimization for the window prevents it from checking each view to see
// if the optimization is supported.
// 
// See [preservesContentDuringLiveResize] in [NSView] for additional
// information on how to support this optimization.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [preservesContentDuringLiveResize]: https://developer.apple.com/documentation/AppKit/NSView/preservesContentDuringLiveResize
// [true]: https://developer.apple.com/documentation/Swift/true
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
// The value of this property is [true] if the window is being live resized;
// otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [NSFullSizeContentViewWindowMask]: https://developer.apple.com/documentation/AppKit/NSFullSizeContentViewWindowMask
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/contentLayoutRect
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
// The value of this property is [true] when the window is onscreen (even if
// it’s obscured by other windows); otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// When the value of this property is [WindowOcclusionStateVisible], at least
// part of the window is visible; otherwise, the window is fully occluded.
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
// The value of this property is [true] if the window is the key window for
// the application; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// The value of this property is [true] if the window can become the key
// window, otherwise, [false].
// 
// Attempts to make the window the key window are abandoned if the value of
// this property is [false]. The value of this property is [true] if the
// window has a title bar or a resize bar, or [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// The value of this property is [true] when the window is the main window for
// the application; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// The value of this property is [true] when the window can become the main
// window; otherwise, [false].
// 
// Attempts to make the window the main window are abandoned if the value of
// this property is [false]. The value of the property is [true] if the window
// is visible, is not an [NSPanel] object, and has a title bar or a resize
// mechanism. Otherwise, the value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// The value of this property is [true] when the window is excluded from the
// Windows menu; otherwise, [false]. The default initial setting is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// The value of this property is [true] when cursor rectangles are enabled;
// otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// The value of this property is [true] if the standard toolbar button is
// currently displayed; otherwise, [false]. When clicked, the toolbar control
// button shows or hides a window’s toolbar. The toolbar control button
// appears in a window’s title bar. If the window does not have a toolbar,
// this property has no effect.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// When the value of this property is [true], the title bar does not draw its
// background, which allows all content underneath it to show through. It only
// makes sense to set this property to [true] when
// [NSFullSizeContentViewWindowMask] is also set.
//
// [NSFullSizeContentViewWindowMask]: https://developer.apple.com/documentation/AppKit/NSFullSizeContentViewWindowMask
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/titlebarAppearsTransparent
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
// The default value is [NSWindow.ToolbarStyle.automatic].
//
// [NSWindow.ToolbarStyle.automatic]: https://developer.apple.com/documentation/AppKit/NSWindow/ToolbarStyle-swift.enum/automatic
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
// The default value is [NSTitlebarSeparatorStyle.automatic]. Changing this
// value overrides any preference by [NSSplitViewItem].
//
// [NSTitlebarSeparatorStyle.automatic]: https://developer.apple.com/documentation/AppKit/NSTitlebarSeparatorStyle/automatic
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
// general, this returns [NSUserInterfaceLayoutDirection.rightToLeft] if the
// primary system language is right to left. The layout direction may be right
// to left even in applications that don’t have a right to left language
// localization. Refer to this value if the application uses
// [TitlebarAppearsTransparent] and places controls under the title bar.
//
// [NSUserInterfaceLayoutDirection.rightToLeft]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutDirection/rightToLeft
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
// default value is [NSWindow.TabbingMode.automatic]. When the value is
// [NSWindow.TabbingMode.automatic], the system uses [UserTabbingPreference]
// to determine tabbing behavior.
// 
// For a list of possible values, see [NSWindow.TabbingMode].
//
// [NSWindow.TabbingMode.automatic]: https://developer.apple.com/documentation/AppKit/NSWindow/TabbingMode-swift.enum/automatic
// [NSWindow.TabbingMode]: https://developer.apple.com/documentation/AppKit/NSWindow/TabbingMode-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/tabbingMode-swift.property
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
// The value of this property is [true] if the window can display tooltips
// even when the application is in the background; otherwise, [false]. The
// default is [false]. Changing the value of this property does not take
// effect until the window changes to an active state.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [Event Handling Basics]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/EventOverview/EventHandlingBasics/EventHandlingBasics.html#//apple_ref/doc/uid/10000060i-CH5
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/firstResponder
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
// [NSWindow.SelectionDirection]: https://developer.apple.com/documentation/AppKit/NSWindow/SelectionDirection
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/keyViewSelectionDirection
func (w NSWindow) KeyViewSelectionDirection() NSSelectionDirection {
	rv := objc.Send[NSSelectionDirection](w.ID, objc.Sel("keyViewSelectionDirection"))
	return NSSelectionDirection(rv)
}
// A Boolean value that indicates whether the window automatically
// recalculates the key view loop when views are added.
//
// # Discussion
// 
// The value of this property is [true] if the window automatically
// recalculates the key view loop when views are added; otherwise, [false]. If
// [AutorecalculatesKeyViewLoop] is [false], the client code must update the
// key view loop manually or call [RecalculateKeyViewLoop] to have the window
// recalculate it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/autorecalculatesKeyViewLoop
func (w NSWindow) AutorecalculatesKeyViewLoop() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("autorecalculatesKeyViewLoop"))
	return rv
}
func (w NSWindow) SetAutorecalculatesKeyViewLoop(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setAutorecalculatesKeyViewLoop:"), value)
}
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
// The value of this property is [true] when the window accepts (and
// distributes) mouse-moved events; otherwise, [false]. By default the value
// is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// The value of this property is [true] when the window is transparent to
// mouse events; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// Set this property to [true] if you want the window to be preserved or
// [false] if you do not want it preserved. By default, the value of this
// property is [true] if the window’s [StyleMask] property includes the
// [NSTitledWindowMask] flag. For other windows, the value is [false]. Setting
// a value explicitly overrides the default values.
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
// [NSTitledWindowMask]: https://developer.apple.com/documentation/AppKit/NSTitledWindowMask
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/isRestorable
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
// The value of this property is [true] when any of the window’s views need
// to be displayed; otherwise, [false]. You should rarely need to set this
// property; the [NSView] method [needsDisplay] and similar methods set it
// automatically.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [needsDisplay]: https://developer.apple.com/documentation/AppKit/NSView/needsDisplay
// [true]: https://developer.apple.com/documentation/Swift/true
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
// The value of this property is [true] if the window allows multithreaded
// view drawing; otherwise, [false]. The default value is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [NSWindow.AnimationBehavior.default], which causes AppKit to determine the
// style of animation to use automatically based on its inference of a
// window’s “type” from various window properties. A window’s
// animation behavior can be set to [NSWindow.AnimationBehavior.none] to
// disable AppKit’s automatic animations for the window, which may be useful
// if that animation interferes with an animation that your application
// implements.
// 
// The animation behavior can also be set to one of the other non-default
// [NSWindow.AnimationBehavior] values to override AppKit’s automatic
// inference of appropriate animation behavior based on the window’s
// apparent type, although this is not recommended.
//
// [NSWindow.AnimationBehavior.default]: https://developer.apple.com/documentation/AppKit/NSWindow/AnimationBehavior-swift.enum/default
// [NSWindow.AnimationBehavior.none]: https://developer.apple.com/documentation/AppKit/NSWindow/AnimationBehavior-swift.enum/none
// [NSWindow.AnimationBehavior]: https://developer.apple.com/documentation/AppKit/NSWindow/AnimationBehavior-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/animationBehavior-swift.property
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
// The value of this property is [true] when the window’s document has been
// edited; otherwise, [false]. Initially, by default, [NSWindow] objects are
// in the “not edited” state.
// 
// You should set [DocumentEdited] to [true] every time the window’s
// document changes in such a way that it needs to be saved. Conversely, when
// the document is saved, you should set the property to [true] when the
// window’s document has been edited; otherwise, [false]. Then, before
// closing the window you can examine the value of the property to determine
// whether to allow the user a chance to save the document.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// By default, the value of this property is
// [NSWindow.TitleVisibility.visible].
//
// [NSWindow.TitleVisibility.visible]: https://developer.apple.com/documentation/AppKit/NSWindow/TitleVisibility-swift.enum/visible
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
func (w NSWindow) RepresentedURL() foundation.INSURL {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("representedURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (w NSWindow) SetRepresentedURL(value foundation.INSURL) {
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
// The value of this property is [true] when the window context should be
// updated when the ColorSync profile of the current screen changes or when a
// majority of the window is moved to a different screen whose profile is
// different than the previous screen; otherwise, [false]. The default value
// is [false].
// 
// After the window context is updated, the window is told to display itself.
// If you need to update offscreen caches for the window, you should register
// to receive the [didChangeScreenProfileNotification] notification.
//
// [didChangeScreenProfileNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didChangeScreenProfileNotification
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/displaysWhenScreenProfileChanges
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
// The value of this property is [true] when the window is movable by clicking
// and dragging anywhere in its background; otherwise, [false].
// 
// A window with a style mask of [NSTexturedBackgroundWindowMask] is movable
// by background by default. Sheets and drawers cannot be movable by window
// background.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// The value of this property is [true] if the window can be moved by the
// user; otherwise, [false].
// 
// When a window’s [Movable] property is [false], the value of the
// [MovableByWindowBackground] property is ignored. When the value of
// [Movable] is [false], the window can only be dragged between spaces in F8
// mode, and its relative screen position is always preserved. Note that a
// resizable window may still be resized, and the window frame may be changed
// programmatically. A nonmovable window will not be moved or resized by the
// system in response to a display reconfiguration. Applications may choose to
// enable application-controlled window dragging after disabling
// user-initiating dragging by handling the
// [MouseDown]/[MouseDragged]/[MouseUp] sequence in [SendEvent] in an
// [NSWindow] subclass.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// The value of this property is [true] if the window is automatically
// released after being closed; [false] if it’s simply removed from the
// screen.
// 
// The default for [NSWindow] is [true]; the default for [NSPanel] is [false].
// Release when closed, however, is ignored for windows owned by window
// controllers. Another strategy for releasing an [NSWindow] object is to have
// its delegate autorelease it on receiving a [WindowShouldClose] message.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// The value of this property is [true] if the window is minimized; otherwise,
// [false]. A minimized window is removed from the screen and replaced by a
// image, icon, or button that represents it, called the counterpart.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// the [AppleDockIconEnabled] key to [true] when first registering your
// application’s user defaults. You must set this key prior to calling the
// `init` method of [NSApplication], which reads the current value of the key.
//
// [true]: https://developer.apple.com/documentation/Swift/true
// [willMiniaturizeNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/willMiniaturizeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/miniwindowImage
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
func (w NSWindow) SetBitsPerPixel(value int) {
	objc.Send[struct{}](w.ID, objc.Sel("setNSBitsPerPixelFromDepth:"), value)
}
// Returns the bits per sample for the specified window depth.
//
// See: https://developer.apple.com/documentation/appkit/nswindow/depth/bitspersample
func (w NSWindow) BitsPerSample() int {
	rv := objc.Send[int](w.ID, objc.Sel("NSBitsPerSampleFromDepth"))
	return rv
}
func (w NSWindow) SetBitsPerSample(value int) {
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
func (w NSWindow) SetColorSpaceName(value NSColorSpaceName) {
	objc.Send[struct{}](w.ID, objc.Sel("setNSColorSpaceFromDepth:"), objc.String(string(value)))
}
// Returns the number of color components in the specified color space.
//
// See: https://developer.apple.com/documentation/appkit/nscolorspacename/numberofcolorcomponents
func (w NSWindow) NumberOfColorComponents() int {
	rv := objc.Send[int](w.ID, objc.Sel("NSNumberOfColorComponents"))
	return rv
}
func (w NSWindow) SetNumberOfColorComponents(value int) {
	objc.Send[struct{}](w.ID, objc.Sel("setNSNumberOfColorComponents:"), value)
}
// Returns whether the specified window depth is planar.
//
// See: https://developer.apple.com/documentation/appkit/nswindow/depth/isplanar
func (w NSWindow) IsPlanar() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("NSPlanarFromDepth"))
	return rv
}
func (w NSWindow) SetIsPlanar(value bool) {
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
//
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
// Sets the option dictionary that maps event trigger keys to animation
// objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/animations
func (w NSWindow) Animations() foundation.INSDictionary {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("animations"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (w NSWindow) SetAnimations(value foundation.INSDictionary) {
	objc.Send[struct{}](w.ID, objc.Sel("setAnimations:"), value)
}
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
func (w NSWindow) Appearance() INSAppearance {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("appearance"))
	return NSAppearanceFromID(objc.ID(rv))
}
func (w NSWindow) SetAppearance(value INSAppearance) {
	objc.Send[struct{}](w.ID, objc.Sel("setAppearance:"), value)
}
// The location of the window’s backing store.
//
// See: https://developer.apple.com/documentation/appkit/nswindow/backinglocation-swift.property
func (w NSWindow) BackingLocation() objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("backingLocation"))
	return objectivec.Object{ID: rv}
}
func (w NSWindow) SetBackingLocation(value objectivec.IObject) {
	objc.Send[struct{}](w.ID, objc.Sel("setBackingLocation:"), value)
}
// The collection of drawers associated with the window.
//
// See: https://developer.apple.com/documentation/appkit/nswindow/drawers
func (w NSWindow) Drawers() objc.ID {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("drawers"))
	return rv
}
func (w NSWindow) SetDrawers(value objc.ID) {
	objc.Send[struct{}](w.ID, objc.Sel("setDrawers:"), value)
}
// The appearance that will be used when the receiver is drawn onscreen, in an
// [NSAppearance] object. (read-only)
//
// # Discussion
// 
// The default value for this property is provided by the nearest ancestor of
// the receiver that has set an appearance.
// 
// You can use this property to ensure that an offscreen view sets the
// appropriate current appearance when it draws onscreen.
//
// See: https://developer.apple.com/documentation/AppKit/NSAppearanceCustomization/effectiveAppearance
func (w NSWindow) EffectiveAppearance() INSAppearance {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("effectiveAppearance"))
	return NSAppearanceFromID(objc.ID(rv))
}
// The graphics context associated with the window for the current thread.
//
// See: https://developer.apple.com/documentation/appkit/nswindow/graphicscontext
func (w NSWindow) GraphicsContext() INSGraphicsContext {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("graphicsContext"))
	return NSGraphicsContextFromID(objc.ID(rv))
}
func (w NSWindow) SetGraphicsContext(value INSGraphicsContext) {
	objc.Send[struct{}](w.ID, objc.Sel("setGraphicsContext:"), value)
}
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
// [Mac Technology Overview]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/OSX_Technology_Overview/About/About.html#//apple_ref/doc/uid/TP40001067
// [OS X Frameworks]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/OSX_Technology_Overview/SystemFrameworks/SystemFrameworks.html#//apple_ref/doc/uid/TP40001067-CH210
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceItemIdentification/identifier
func (w NSWindow) Identifier() NSUserInterfaceItemIdentifier {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("identifier"))
	return NSUserInterfaceItemIdentifier(foundation.NSStringFromID(rv).String())
}
func (w NSWindow) SetIdentifier(value NSUserInterfaceItemIdentifier) {
	objc.Send[struct{}](w.ID, objc.Sel("setIdentifier:"), objc.String(string(value)))
}
// Name of an exception that occurs when you pass an invalid argument to a
// method, such as a `nil` pointer where a non-`nil` object is required.
//
// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
func (w NSWindow) InvalidArgumentException() foundation.NSString {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("invalidArgumentException"))
	return foundation.NSStringFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the window automatically displays
// views that need to be displayed.
//
// See: https://developer.apple.com/documentation/appkit/nswindow/isautodisplay
func (w NSWindow) IsAutodisplay() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("autodisplay"))
	return rv
}
func (w NSWindow) SetIsAutodisplay(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setAutodisplay:"), value)
}
// A Boolean value that indicates whether the window’s flushing ability is
// disabled.
//
// See: https://developer.apple.com/documentation/appkit/nswindow/isflushwindowdisabled
func (w NSWindow) IsFlushWindowDisabled() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("flushWindowDisabled"))
	return rv
}
func (w NSWindow) SetIsFlushWindowDisabled(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setFlushWindowDisabled:"), value)
}
// A Boolean value that indicates whether the window device the window manages
// is freed when it’s removed from the screen list.
//
// See: https://developer.apple.com/documentation/appkit/nswindow/isoneshot
func (w NSWindow) IsOneShot() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("oneShot"))
	return rv
}
func (w NSWindow) SetIsOneShot(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setOneShot:"), value)
}
// A Boolean value that indicates the preferred location for the window’s
// backing store.
//
// See: https://developer.apple.com/documentation/appkit/nswindow/preferredbackinglocation
func (w NSWindow) PreferredBackingLocation() objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("preferredBackingLocation"))
	return objectivec.Object{ID: rv}
}
func (w NSWindow) SetPreferredBackingLocation(value objectivec.IObject) {
	objc.Send[struct{}](w.ID, objc.Sel("setPreferredBackingLocation:"), value)
}
// A Boolean value that indicates whether the window’s resize indicator is
// visible.
//
// See: https://developer.apple.com/documentation/appkit/nswindow/showsresizeindicator
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
// See: https://developer.apple.com/documentation/appkit/nswindow/windowref
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
// [NSWindow.UserTabbingPreference]: https://developer.apple.com/documentation/AppKit/NSWindow/UserTabbingPreference-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/userTabbingPreference-swift.type.property
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

			// Protocol methods for NSAccessibilityElementProtocol
			
// Returns the accessibility element’s frame in screen coordinates.
//
// # Return Value
// 
// The element’s frame in screen coordinates.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityFrame] property. This method is called whenever accessibility
// clients request the [size] or [position] attributes.
//
// [accessibilityFrame]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrame
// [position]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/position
// [size]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/size
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityFrame()
func (o NSWindow) AccessibilityFrame() corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("accessibilityFrame"))
	return rv
	}
// Returns the accessibility element’s parent in the accessibility
// hierarchy.
//
// # Return Value
// 
// The element’s parent in the accessibility hierarchy.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityParent] property.
//
// [accessibilityParent]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityParent
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityParent()
func (o NSWindow) AccessibilityParent() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityParent"))
	return objectivec.Object{ID: rv}
	}
// Returns the accessibility element’s identity.
//
// # Return Value
// 
// Returns the unique ID for the accessibility element. It is often used in
// automated testing.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityIdentifier] property.
//
// [accessibilityIdentifier]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIdentifier
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityIdentifier()
func (o NSWindow) AccessibilityIdentifier() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIdentifier"))
	return foundation.NSStringFromID(rv).String()
	}
// Returns a Boolean value that indicates whether the accessibility element
// has the keyboard focus.
//
// # Return Value
// 
// [true] if this element has the keyboard focus; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityFocused] property.
//
// [accessibilityFocused]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFocused
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/isAccessibilityFocused()
func (o NSWindow) IsAccessibilityFocused() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
	}

			// Protocol methods for NSAccessibilityProtocol
			
// Returns a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityelement()
func (o NSWindow) IsAccessibilityElement() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityElement"))
	return rv
	}
// Sets a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityelement(_:)
func (o NSWindow) SetAccessibilityElement(accessibilityElement bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityElement:"), accessibilityElement)
	}
// Returns a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityenabled()
func (o NSWindow) IsAccessibilityEnabled() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEnabled"))
	return rv
	}
// Sets a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityenabled(_:)
func (o NSWindow) SetAccessibilityEnabled(accessibilityEnabled bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityEnabled:"), accessibilityEnabled)
	}
// Sets the accessibility element’s frame in screen coordinates.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityframe(_:)
func (o NSWindow) SetAccessibilityFrame(accessibilityFrame corefoundation.CGRect) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrame:"), accessibilityFrame)
	}
// Returns the help text for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhelp()
func (o NSWindow) AccessibilityHelp() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHelp"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the help text for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhelp(_:)
func (o NSWindow) SetAccessibilityHelp(accessibilityHelp string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHelp:"), objc.String(accessibilityHelp))
	}
// Returns a short description of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylabel()
func (o NSWindow) AccessibilityLabel() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabel"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets a short description of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylabel(_:)
func (o NSWindow) SetAccessibilityLabel(accessibilityLabel string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabel:"), objc.String(accessibilityLabel))
	}
// Returns the title of the accessibility element—for example, a button’s
// visible text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytitle()
func (o NSWindow) AccessibilityTitle() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitle"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the title of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytitle(_:)
func (o NSWindow) SetAccessibilityTitle(accessibilityTitle string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTitle:"), objc.String(accessibilityTitle))
	}
// Returns the accessibility element’s value.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvalue()
func (o NSWindow) AccessibilityValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValue"))
	return objectivec.Object{ID: rv}
	}
// Sets the accessibility element’s value.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvalue(_:)
func (o NSWindow) SetAccessibilityValue(accessibilityValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityValue:"), accessibilityValue)
	}
// Returns a Boolean value that indicates whether assistive apps can invoke
// the specified selector on the accessibility element.
//
// selector: The selector to check.
//
// # Return Value
// 
// [true], if accessibility clients can call the selector; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilitySelectorAllowed(_:)
func (o NSWindow) IsAccessibilitySelectorAllowed(selector objc.SEL) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelectorAllowed:"), selector)
	return rv
	}
// Returns the contents of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycontents()
func (o NSWindow) AccessibilityContents() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityContents"))
	return objectivec.Object{ID: rv}
	}
// Sets the contents of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycontents(_:)
func (o NSWindow) SetAccessibilityContents(accessibilityContents foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityContents:"), accessibilityContents)
	}
// Returns the critical value for the level indicator.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycriticalvalue()
func (o NSWindow) AccessibilityCriticalValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCriticalValue"))
	return objectivec.Object{ID: rv}
	}
// Sets the critical value for the level indicator.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycriticalvalue(_:)
func (o NSWindow) SetAccessibilityCriticalValue(accessibilityCriticalValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCriticalValue:"), accessibilityCriticalValue)
	}
// Sets the accessibility element’s identity.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityidentifier(_:)
func (o NSWindow) SetAccessibilityIdentifier(accessibilityIdentifier string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIdentifier:"), objc.String(accessibilityIdentifier))
	}
// Returns the maximum value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymaxvalue()
func (o NSWindow) AccessibilityMaxValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMaxValue"))
	return objectivec.Object{ID: rv}
	}
// Sets the maximum value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymaxvalue(_:)
func (o NSWindow) SetAccessibilityMaxValue(accessibilityMaxValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMaxValue:"), accessibilityMaxValue)
	}
// Returns the minimum value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityminvalue()
func (o NSWindow) AccessibilityMinValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinValue"))
	return objectivec.Object{ID: rv}
	}
// Sets the minimum value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityminvalue(_:)
func (o NSWindow) SetAccessibilityMinValue(accessibilityMinValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinValue:"), accessibilityMinValue)
	}
// Returns the orientation of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityorientation()
func (o NSWindow) AccessibilityOrientation() NSAccessibilityOrientation {
	
	rv := objc.Send[NSAccessibilityOrientation](o.ID, objc.Sel("accessibilityOrientation"))
	return rv
	}
// Sets the orientation of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityorientation(_:)
func (o NSWindow) SetAccessibilityOrientation(accessibilityOrientation NSAccessibilityOrientation) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOrientation:"), accessibilityOrientation)
	}
// Returns a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityprotectedcontent()
func (o NSWindow) IsAccessibilityProtectedContent() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityProtectedContent"))
	return rv
	}
// Sets a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityprotectedcontent(_:)
func (o NSWindow) SetAccessibilityProtectedContent(accessibilityProtectedContent bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityProtectedContent:"), accessibilityProtectedContent)
	}
// Returns a Boolean value that determines whether the accessibility element
// is currently in a selected state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityselected()
func (o NSWindow) IsAccessibilitySelected() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelected"))
	return rv
	}
// Sets a Boolean value that determines whether the accessibility element is
// currently in a selected state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselected(_:)
func (o NSWindow) SetAccessibilitySelected(accessibilitySelected bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelected:"), accessibilitySelected)
	}
// Returns the URL for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityurl()
func (o NSWindow) AccessibilityURL() foundation.INSURL {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityURL"))
	return foundation.NSURLFromID(rv)
	}
// Sets the URL for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityurl(_:)
func (o NSWindow) SetAccessibilityURL(accessibilityURL foundation.INSURL) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityURL:"), accessibilityURL)
	}
// Returns the human-readable description of the accessibility element’s
// value.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvaluedescription()
func (o NSWindow) AccessibilityValueDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValueDescription"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the human-readable description of the accessibility element’s value.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvaluedescription(_:)
func (o NSWindow) SetAccessibilityValueDescription(accessibilityValueDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityValueDescription:"), objc.String(accessibilityValueDescription))
	}
// Returns the warning value for the level indicator.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitywarningvalue()
func (o NSWindow) AccessibilityWarningValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWarningValue"))
	return objectivec.Object{ID: rv}
	}
// Sets the warning value for the level indicator.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitywarningvalue(_:)
func (o NSWindow) SetAccessibilityWarningValue(accessibilityWarningValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWarningValue:"), accessibilityWarningValue)
	}
// Returns the child accessibility elements in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitychildren()
func (o NSWindow) AccessibilityChildren() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityChildren"))
	return objectivec.Object{ID: rv}
	}
// Sets the child accessibility elements in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitychildren(_:)
func (o NSWindow) SetAccessibilityChildren(accessibilityChildren foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChildren:"), accessibilityChildren)
	}
// Returns the array of child accessibility elements in order for linear
// navigation.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitychildreninnavigationorder()
func (o NSWindow) AccessibilityChildrenInNavigationOrder() unsafe.Pointer {
	
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("accessibilityChildrenInNavigationOrder"))
	return rv
	}
// Sets the array of child accessibility elements in order for linear
// navigation.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitychildreninnavigationorder(_:)
func (o NSWindow) SetAccessibilityChildrenInNavigationOrder(accessibilityChildrenInNavigationOrder foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChildrenInNavigationOrder:"), accessibilityChildrenInNavigationOrder)
	}
// Sets the accessibility element’s parent in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityparent(_:)
func (o NSWindow) SetAccessibilityParent(accessibilityParent objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityParent:"), accessibilityParent)
	}
// Returns the accessibility element’s currently selected children.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedchildren()
func (o NSWindow) AccessibilitySelectedChildren() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedChildren"))
	return objectivec.Object{ID: rv}
	}
// Sets the accessibility element’s currently selected children.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedchildren(_:)
func (o NSWindow) SetAccessibilitySelectedChildren(accessibilitySelectedChildren foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedChildren:"), accessibilitySelectedChildren)
	}
// Returns the top-level element that contains the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytopleveluielement()
func (o NSWindow) AccessibilityTopLevelUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTopLevelUIElement"))
	return objectivec.Object{ID: rv}
	}
// Sets the top-level element that contains the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytopleveluielement(_:)
func (o NSWindow) SetAccessibilityTopLevelUIElement(accessibilityTopLevelUIElement objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTopLevelUIElement:"), accessibilityTopLevelUIElement)
	}
// Returns the accessibility element’s visible child accessibility elements.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblechildren()
func (o NSWindow) AccessibilityVisibleChildren() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleChildren"))
	return objectivec.Object{ID: rv}
	}
// Sets the accessibility element’s visible child accessibility elements.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblechildren(_:)
func (o NSWindow) SetAccessibilityVisibleChildren(accessibilityVisibleChildren foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleChildren:"), accessibilityVisibleChildren)
	}
// Returns the child accessibility element with the current focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityapplicationfocuseduielement()
func (o NSWindow) AccessibilityApplicationFocusedUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityApplicationFocusedUIElement"))
	return objectivec.Object{ID: rv}
	}
// Sets the child accessibility element with the current focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityapplicationfocuseduielement(_:)
func (o NSWindow) SetAccessibilityApplicationFocusedUIElement(accessibilityApplicationFocusedUIElement objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityApplicationFocusedUIElement:"), accessibilityApplicationFocusedUIElement)
	}
// Sets a Boolean value that determines whether the accessibility element has
// the keyboard focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfocused(_:)
func (o NSWindow) SetAccessibilityFocused(accessibilityFocused bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFocused:"), accessibilityFocused)
	}
// Returns the child window with the current focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityfocusedwindow()
func (o NSWindow) AccessibilityFocusedWindow() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFocusedWindow"))
	return objectivec.Object{ID: rv}
	}
// Sets the child window with the current focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfocusedwindow(_:)
func (o NSWindow) SetAccessibilityFocusedWindow(accessibilityFocusedWindow objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFocusedWindow:"), accessibilityFocusedWindow)
	}
// Returns the array of elements that shares the keyboard focus with the
// accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysharedfocuselements()
func (o NSWindow) AccessibilitySharedFocusElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedFocusElements"))
	return objectivec.Object{ID: rv}
	}
// Sets the array of elements that shares the keyboard focus with the
// accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysharedfocuselements(_:)
func (o NSWindow) SetAccessibilitySharedFocusElements(accessibilitySharedFocusElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedFocusElements:"), accessibilitySharedFocusElements)
	}
// Returns a Boolean value that determines whether the accessibility element
// must have content for successful submission of a form.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityrequired()
func (o NSWindow) IsAccessibilityRequired() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityRequired"))
	return rv
	}
// Sets a Boolean value that determines whether the accessibility element must
// have content for successful submission of a form.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrequired(_:)
func (o NSWindow) SetAccessibilityRequired(accessibilityRequired bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRequired:"), accessibilityRequired)
	}
// Returns the type of interface element that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrole()
func (o NSWindow) AccessibilityRole() NSAccessibilityRole {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRole"))
	return NSAccessibilityRole(foundation.NSStringFromID(rv).String())
	}
// Sets the type of interface element that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrole(_:)
func (o NSWindow) SetAccessibilityRole(accessibilityRole NSAccessibilityRole) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRole:"), objc.String(string(accessibilityRole)))
	}
// Returns a localized, human-intelligible description of the accessibility
// element’s role, such as
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityroledescription()
func (o NSWindow) AccessibilityRoleDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRoleDescription"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the localized, human-intelligible description of the accessibility
// element’s role, such as
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityroledescription(_:)
func (o NSWindow) SetAccessibilityRoleDescription(accessibilityRoleDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRoleDescription:"), objc.String(accessibilityRoleDescription))
	}
// Returns the specialized interface element type that the accessibility
// element represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysubrole()
func (o NSWindow) AccessibilitySubrole() NSAccessibilitySubrole {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySubrole"))
	return NSAccessibilitySubrole(foundation.NSStringFromID(rv).String())
	}
// Sets the specialized interface element type that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysubrole(_:)
func (o NSWindow) SetAccessibilitySubrole(accessibilitySubrole NSAccessibilitySubrole) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySubrole:"), objc.String(string(accessibilitySubrole)))
	}
// Returns the custom actions of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycustomactions()
func (o NSWindow) AccessibilityCustomActions() INSAccessibilityCustomAction {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCustomActions"))
	return NSAccessibilityCustomActionFromID(rv)
	}
// Sets the custom actions of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycustomactions(_:)
func (o NSWindow) SetAccessibilityCustomActions(accessibilityCustomActions foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomActions:"), accessibilityCustomActions)
	}
// Returns the custom rotors of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycustomrotors()
func (o NSWindow) AccessibilityCustomRotors() INSAccessibilityCustomRotor {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCustomRotors"))
	return NSAccessibilityCustomRotorFromID(rv)
	}
// Sets the custom rotors of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycustomrotors(_:)
func (o NSWindow) SetAccessibilityCustomRotors(accessibilityCustomRotors foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomRotors:"), accessibilityCustomRotors)
	}
// Returns the line number that contains the insertion point.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityinsertionpointlinenumber()
func (o NSWindow) AccessibilityInsertionPointLineNumber() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityInsertionPointLineNumber"))
	return rv
	}
// Sets the line number that contains the insertion point.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityinsertionpointlinenumber(_:)
func (o NSWindow) SetAccessibilityInsertionPointLineNumber(accessibilityInsertionPointLineNumber int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityInsertionPointLineNumber:"), accessibilityInsertionPointLineNumber)
	}
// Returns the number of characters in the text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitynumberofcharacters()
func (o NSWindow) AccessibilityNumberOfCharacters() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityNumberOfCharacters"))
	return rv
	}
// Sets the number of characters in the text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitynumberofcharacters(_:)
func (o NSWindow) SetAccessibilityNumberOfCharacters(accessibilityNumberOfCharacters int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityNumberOfCharacters:"), accessibilityNumberOfCharacters)
	}
// Returns the placeholder value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityplaceholdervalue()
func (o NSWindow) AccessibilityPlaceholderValue() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPlaceholderValue"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the placeholder value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityplaceholdervalue(_:)
func (o NSWindow) SetAccessibilityPlaceholderValue(accessibilityPlaceholderValue string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityPlaceholderValue:"), objc.String(accessibilityPlaceholderValue))
	}
// Returns the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedtext()
func (o NSWindow) AccessibilitySelectedText() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedText"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedtext(_:)
func (o NSWindow) SetAccessibilitySelectedText(accessibilitySelectedText string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedText:"), objc.String(accessibilitySelectedText))
	}
// Returns the range of the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedtextrange()
func (o NSWindow) AccessibilitySelectedTextRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySelectedTextRange"))
	return rv
	}
// Sets the range of the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedtextrange(_:)
func (o NSWindow) SetAccessibilitySelectedTextRange(accessibilitySelectedTextRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedTextRange:"), accessibilitySelectedTextRange)
	}
// Returns an array of ranges for the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedtextranges()
func (o NSWindow) AccessibilitySelectedTextRanges() foundation.NSValue {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedTextRanges"))
	return foundation.NSValueFromID(rv)
	}
// Sets an array of ranges for the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedtextranges(_:)
func (o NSWindow) SetAccessibilitySelectedTextRanges(accessibilitySelectedTextRanges foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedTextRanges:"), accessibilitySelectedTextRanges)
	}
// Returns the range of characters that the accessibility element displays.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysharedcharacterrange()
func (o NSWindow) AccessibilitySharedCharacterRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySharedCharacterRange"))
	return rv
	}
// Sets the range of characters that the accessibility element displays.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysharedcharacterrange(_:)
func (o NSWindow) SetAccessibilitySharedCharacterRange(accessibilitySharedCharacterRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedCharacterRange:"), accessibilitySharedCharacterRange)
	}
// Returns the other elements that share text with the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysharedtextuielements()
func (o NSWindow) AccessibilitySharedTextUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedTextUIElements"))
	return objectivec.Object{ID: rv}
	}
// Sets the other elements that share text with the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysharedtextuielements(_:)
func (o NSWindow) SetAccessibilitySharedTextUIElements(accessibilitySharedTextUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedTextUIElements:"), accessibilitySharedTextUIElements)
	}
// Returns the range of visible characters in the document.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblecharacterrange()
func (o NSWindow) AccessibilityVisibleCharacterRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityVisibleCharacterRange"))
	return rv
	}
// Sets the range of visible characters in the document.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblecharacterrange(_:)
func (o NSWindow) SetAccessibilityVisibleCharacterRange(accessibilityVisibleCharacterRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleCharacterRange:"), accessibilityVisibleCharacterRange)
	}
// Returns the substring for the specified range.
//
// range: A range of characters contained by the element.
//
// # Return Value
// 
// The substring specified by the given range.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityString(for:)
func (o NSWindow) AccessibilityStringForRange(range_ foundation.NSRange) string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityStringForRange:"), range_)
	return foundation.NSStringFromID(rv).String()
	}
// Returns the attributed substring for the specified range of characters.
//
// range: The range of characters.
//
// # Return Value
// 
// An attributed string representing the specified characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityAttributedString(for:)
func (o NSWindow) AccessibilityAttributedStringForRange(range_ foundation.NSRange) foundation.NSAttributedString {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAttributedStringForRange:"), range_)
	return foundation.NSAttributedStringFromID(rv)
	}
// Returns the rich text format (RTF) data that describes the specified range
// of characters.
//
// range: The range of characters.
//
// # Return Value
// 
// A data object containing an RTF representation of the specified characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRTF(for:)
func (o NSWindow) AccessibilityRTFForRange(range_ foundation.NSRange) foundation.INSData {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRTFForRange:"), range_)
	return foundation.NSDataFromID(rv)
	}
// Returns the rectangle that encloses the specified range of characters.
//
// range: The range of characters.
//
// # Return Value
// 
// The rectangle that encloses the specified characters.
//
// # Discussion
// 
// If the range crosses a line boundary, the returned rectangle fully encloses
// all the lines of characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFrame(for:)
func (o NSWindow) AccessibilityFrameForRange(range_ foundation.NSRange) corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("accessibilityFrameForRange:"), range_)
	return rv
	}
// Returns the line number for the line that contains the specified character
// index.
//
// index: The index for a character.
//
// # Return Value
// 
// The line number for the line holding the specified character index.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLine(for:)
func (o NSWindow) AccessibilityLineForIndex(index int) int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityLineForIndex:"), index)
	return rv
	}
// Returns the range of characters for the glyph that includes the specified
// character.
//
// index: The specified character.
//
// # Return Value
// 
// The range of characters for the glyph.
//
// # Discussion
// 
// This value always includes the specified character but may include
// additional characters if that character is part of a multicharacter glyph.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRange(for:)-6kv3
func (o NSWindow) AccessibilityRangeForIndex(index int) foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForIndex:"), index)
	return rv
	}
// Returns a range of characters that all have the same style as the specified
// character.
//
// index: The index of the specified character.
//
// # Return Value
// 
// A range of characters with the same style as the specified character.
//
// # Discussion
// 
// This method returns a range of characters that meet two conditions: The
// range must include the specified character, and all the other characters in
// the range must match the specified character’s style. If none of the
// adjacent characters match the specified character’s style, the method
// returns only the specified character.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityStyleRange(for:)
func (o NSWindow) AccessibilityStyleRangeForIndex(index int) foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityStyleRangeForIndex:"), index)
	return rv
	}
// Returns the range of characters in the specified line.
//
// line: The line number to be examined.
//
// # Return Value
// 
// The range of characters for the specified line number. If the line ends
// with a newline character, including the newline is preferred.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRange(forLine:)
func (o NSWindow) AccessibilityRangeForLine(line int) foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForLine:"), line)
	return rv
	}
// Returns the range of characters for the glyph at the specified point.
//
// point: A point in screen coordinates.
//
// # Return Value
// 
// The range of characters that make up the glyph at the given point.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRange(for:)-1iudm
func (o NSWindow) AccessibilityRangeForPosition(point corefoundation.CGPoint) foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForPosition:"), point)
	return rv
	}
// Returns the activation point for the user interface element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityactivationpoint()
func (o NSWindow) AccessibilityActivationPoint() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityActivationPoint"))
	return rv
	}
// Sets the activation point for the user interface element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityactivationpoint(_:)
func (o NSWindow) SetAccessibilityActivationPoint(accessibilityActivationPoint corefoundation.CGPoint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityActivationPoint:"), accessibilityActivationPoint)
	}
// Returns the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityalternateuivisible()
func (o NSWindow) IsAccessibilityAlternateUIVisible() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityAlternateUIVisible"))
	return rv
	}
// Sets the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityalternateuivisible(_:)
func (o NSWindow) SetAccessibilityAlternateUIVisible(accessibilityAlternateUIVisible bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAlternateUIVisible:"), accessibilityAlternateUIVisible)
	}
// Returns the child accessibility element that represents the window’s
// cancel button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycancelbutton()
func (o NSWindow) AccessibilityCancelButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCancelButton"))
	return objectivec.Object{ID: rv}
	}
// Sets the child accessibility element that represents the window’s cancel
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycancelbutton(_:)
func (o NSWindow) SetAccessibilityCancelButton(accessibilityCancelButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCancelButton:"), accessibilityCancelButton)
	}
// Returns the child accessibility element that represents the window’s
// close button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityclosebutton()
func (o NSWindow) AccessibilityCloseButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCloseButton"))
	return objectivec.Object{ID: rv}
	}
// Sets the child accessibility element that represents the window’s close
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityclosebutton(_:)
func (o NSWindow) SetAccessibilityCloseButton(accessibilityCloseButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCloseButton:"), accessibilityCloseButton)
	}
// Returns the child accessibility element that represents the window’s
// default button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydefaultbutton()
func (o NSWindow) AccessibilityDefaultButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDefaultButton"))
	return objectivec.Object{ID: rv}
	}
// Sets the child accessibility element that represents the window’s default
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydefaultbutton(_:)
func (o NSWindow) SetAccessibilityDefaultButton(accessibilityDefaultButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDefaultButton:"), accessibilityDefaultButton)
	}
// Returns the child accessibility element that represents the window’s
// full-screen button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityfullscreenbutton()
func (o NSWindow) AccessibilityFullScreenButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFullScreenButton"))
	return objectivec.Object{ID: rv}
	}
// Sets the child accessibility element that represents the window’s
// full-screen button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfullscreenbutton(_:)
func (o NSWindow) SetAccessibilityFullScreenButton(accessibilityFullScreenButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFullScreenButton:"), accessibilityFullScreenButton)
	}
// Returns the child accessibility element that represents the window’s grow
// area.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitygrowarea()
func (o NSWindow) AccessibilityGrowArea() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityGrowArea"))
	return objectivec.Object{ID: rv}
	}
// Sets the child accessibility element that represents the window’s grow
// area.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitygrowarea(_:)
func (o NSWindow) SetAccessibilityGrowArea(accessibilityGrowArea objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityGrowArea:"), accessibilityGrowArea)
	}
// Returns a Boolean value that determines whether the window is the app’s
// main window.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilitymain()
func (o NSWindow) IsAccessibilityMain() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMain"))
	return rv
	}
// Sets a Boolean value that determines whether the window is the app’s main
// window.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymain(_:)
func (o NSWindow) SetAccessibilityMain(accessibilityMain bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMain:"), accessibilityMain)
	}
// Returns the child accessibility element that represents the window’s
// minimize button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityminimizebutton()
func (o NSWindow) AccessibilityMinimizeButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinimizeButton"))
	return objectivec.Object{ID: rv}
	}
// Sets the child accessibility element that represents the window’s
// minimize button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityminimizebutton(_:)
func (o NSWindow) SetAccessibilityMinimizeButton(accessibilityMinimizeButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimizeButton:"), accessibilityMinimizeButton)
	}
// Returns the Boolean value that determines whether the window is in a
// minimized state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityminimized()
func (o NSWindow) IsAccessibilityMinimized() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMinimized"))
	return rv
	}
// Sets the Boolean value that determines whether the window is in a minimized
// state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityminimized(_:)
func (o NSWindow) SetAccessibilityMinimized(accessibilityMinimized bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimized:"), accessibilityMinimized)
	}
// Returns a Boolean value that determines whether the window is modal.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilitymodal()
func (o NSWindow) IsAccessibilityModal() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityModal"))
	return rv
	}
// Sets a Boolean value that determines whether the window is modal.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymodal(_:)
func (o NSWindow) SetAccessibilityModal(accessibilityModal bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityModal:"), accessibilityModal)
	}
// Returns the child accessibility element that represents the window’s
// proxy icon.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityproxy()
func (o NSWindow) AccessibilityProxy() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityProxy"))
	return objectivec.Object{ID: rv}
	}
// Sets the child accessibility element that represents the window’s proxy
// icon.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityproxy(_:)
func (o NSWindow) SetAccessibilityProxy(accessibilityProxy objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityProxy:"), accessibilityProxy)
	}
// Returns the menu currently displaying for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityshownmenu()
func (o NSWindow) AccessibilityShownMenu() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityShownMenu"))
	return objectivec.Object{ID: rv}
	}
// Sets the menu currently displaying for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityshownmenu(_:)
func (o NSWindow) SetAccessibilityShownMenu(accessibilityShownMenu objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityShownMenu:"), accessibilityShownMenu)
	}
// Returns the child accessibility element that represents the window’s
// toolbar button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytoolbarbutton()
func (o NSWindow) AccessibilityToolbarButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityToolbarButton"))
	return objectivec.Object{ID: rv}
	}
// Sets the child accessibility element that represents the window’s toolbar
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytoolbarbutton(_:)
func (o NSWindow) SetAccessibilityToolbarButton(accessibilityToolbarButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityToolbarButton:"), accessibilityToolbarButton)
	}
// Returns the window that contains the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitywindow()
func (o NSWindow) AccessibilityWindow() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindow"))
	return objectivec.Object{ID: rv}
	}
// Sets the window that contains the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitywindow(_:)
func (o NSWindow) SetAccessibilityWindow(accessibilityWindow objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindow:"), accessibilityWindow)
	}
// Returns the child accessibility element that represents the window’s zoom
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityzoombutton()
func (o NSWindow) AccessibilityZoomButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityZoomButton"))
	return objectivec.Object{ID: rv}
	}
// Sets the child accessibility element that represents the window’s zoom
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityzoombutton(_:)
func (o NSWindow) SetAccessibilityZoomButton(accessibilityZoomButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityZoomButton:"), accessibilityZoomButton)
	}
// Returns the icon for the app’s menu bar extra.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityextrasmenubar()
func (o NSWindow) AccessibilityExtrasMenuBar() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityExtrasMenuBar"))
	return objectivec.Object{ID: rv}
	}
// Sets the icon for the app’s menu bar extra.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityextrasmenubar(_:)
func (o NSWindow) SetAccessibilityExtrasMenuBar(accessibilityExtrasMenuBar objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExtrasMenuBar:"), accessibilityExtrasMenuBar)
	}
// Returns a Boolean value that determines whether the app is the frontmost
// app.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityfrontmost()
func (o NSWindow) IsAccessibilityFrontmost() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFrontmost"))
	return rv
	}
// Sets a Boolean value that determines whether the app is the frontmost app.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfrontmost(_:)
func (o NSWindow) SetAccessibilityFrontmost(accessibilityFrontmost bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrontmost:"), accessibilityFrontmost)
	}
// Returns a Boolean value that determines whether the app is in a hidden
// state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityhidden()
func (o NSWindow) IsAccessibilityHidden() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityHidden"))
	return rv
	}
// Sets a Boolean value that determines whether the app is in a hidden state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhidden(_:)
func (o NSWindow) SetAccessibilityHidden(accessibilityHidden bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHidden:"), accessibilityHidden)
	}
// Returns the app’s main window.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymainwindow()
func (o NSWindow) AccessibilityMainWindow() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMainWindow"))
	return objectivec.Object{ID: rv}
	}
// Sets the app’s main window.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymainwindow(_:)
func (o NSWindow) SetAccessibilityMainWindow(accessibilityMainWindow objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMainWindow:"), accessibilityMainWindow)
	}
// Returns the app’s menu bar.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymenubar()
func (o NSWindow) AccessibilityMenuBar() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMenuBar"))
	return objectivec.Object{ID: rv}
	}
// Sets the app’s menu bar.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymenubar(_:)
func (o NSWindow) SetAccessibilityMenuBar(accessibilityMenuBar objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMenuBar:"), accessibilityMenuBar)
	}
// Returns an array that contains all the app’s windows.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitywindows()
func (o NSWindow) AccessibilityWindows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindows"))
	return objectivec.Object{ID: rv}
	}
// Sets the array that contains all the app’s windows.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitywindows(_:)
func (o NSWindow) SetAccessibilityWindows(accessibilityWindows foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindows:"), accessibilityWindows)
	}
// Returns the number of columns in the accessibility element’s grid.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumncount()
func (o NSWindow) AccessibilityColumnCount() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityColumnCount"))
	return rv
	}
// Sets the number of columns in the accessibility element’s grid.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumncount(_:)
func (o NSWindow) SetAccessibilityColumnCount(accessibilityColumnCount int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnCount:"), accessibilityColumnCount)
	}
// Returns a Boolean value that determines whether the accessibility
// element’s grid is in row major order or in column major order.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityorderedbyrow()
func (o NSWindow) IsAccessibilityOrderedByRow() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityOrderedByRow"))
	return rv
	}
// Sets a Boolean value that determines whether the element’s grid is in row
// major order or in column major order.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityorderedbyrow(_:)
func (o NSWindow) SetAccessibilityOrderedByRow(accessibilityOrderedByRow bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOrderedByRow:"), accessibilityOrderedByRow)
	}
// Returns the number of rows in the accessibility element’s grid.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrowcount()
func (o NSWindow) AccessibilityRowCount() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityRowCount"))
	return rv
	}
// Sets the number of rows in the accessibility element’s grid.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrowcount(_:)
func (o NSWindow) SetAccessibilityRowCount(accessibilityRowCount int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowCount:"), accessibilityRowCount)
	}
// Returns the horizontal scroll bar for the scroll view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhorizontalscrollbar()
func (o NSWindow) AccessibilityHorizontalScrollBar() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalScrollBar"))
	return objectivec.Object{ID: rv}
	}
// Sets the horizontal scroll bar for the scroll view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhorizontalscrollbar(_:)
func (o NSWindow) SetAccessibilityHorizontalScrollBar(accessibilityHorizontalScrollBar objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalScrollBar:"), accessibilityHorizontalScrollBar)
	}
// Returns the vertical scroll bar for the scroll view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityverticalscrollbar()
func (o NSWindow) AccessibilityVerticalScrollBar() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalScrollBar"))
	return objectivec.Object{ID: rv}
	}
// Sets the vertical scroll bar for the scroll view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityverticalscrollbar(_:)
func (o NSWindow) SetAccessibilityVerticalScrollBar(accessibilityVerticalScrollBar objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalScrollBar:"), accessibilityVerticalScrollBar)
	}
// Returns the column header accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumnheaderuielements()
func (o NSWindow) AccessibilityColumnHeaderUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnHeaderUIElements"))
	return objectivec.Object{ID: rv}
	}
// Sets the column header accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumnheaderuielements(_:)
func (o NSWindow) SetAccessibilityColumnHeaderUIElements(accessibilityColumnHeaderUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnHeaderUIElements:"), accessibilityColumnHeaderUIElements)
	}
// Returns the column accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumns()
func (o NSWindow) AccessibilityColumns() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumns"))
	return objectivec.Object{ID: rv}
	}
// Sets the column accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumns(_:)
func (o NSWindow) SetAccessibilityColumns(accessibilityColumns foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumns:"), accessibilityColumns)
	}
// Returns the column titles for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumntitles()
func (o NSWindow) AccessibilityColumnTitles() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnTitles"))
	return objectivec.Object{ID: rv}
	}
// Sets the column titles for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumntitles(_:)
func (o NSWindow) SetAccessibilityColumnTitles(accessibilityColumnTitles foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnTitles:"), accessibilityColumnTitles)
	}
// Returns a Boolean value that determines whether the accessibility element
// is in an expanded state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityexpanded()
func (o NSWindow) IsAccessibilityExpanded() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityExpanded"))
	return rv
	}
// Sets a Boolean value that determines whether accessibility element is in an
// expanded state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityexpanded(_:)
func (o NSWindow) SetAccessibilityExpanded(accessibilityExpanded bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExpanded:"), accessibilityExpanded)
	}
// Returns the header for the table view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityheader()
func (o NSWindow) AccessibilityHeader() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHeader"))
	return objectivec.Object{ID: rv}
	}
// Sets the header for the table view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityheader(_:)
func (o NSWindow) SetAccessibilityHeader(accessibilityHeader objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHeader:"), accessibilityHeader)
	}
// Returns the index of the row or column that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityindex()
func (o NSWindow) AccessibilityIndex() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityIndex"))
	return rv
	}
// Sets the index of the row or column that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityindex(_:)
func (o NSWindow) SetAccessibilityIndex(accessibilityIndex int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIndex:"), accessibilityIndex)
	}
// Returns the row header accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrowheaderuielements()
func (o NSWindow) AccessibilityRowHeaderUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRowHeaderUIElements"))
	return objectivec.Object{ID: rv}
	}
// Sets the row header accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrowheaderuielements(_:)
func (o NSWindow) SetAccessibilityRowHeaderUIElements(accessibilityRowHeaderUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowHeaderUIElements:"), accessibilityRowHeaderUIElements)
	}
// Returns the row accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrows()
func (o NSWindow) AccessibilityRows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRows"))
	return objectivec.Object{ID: rv}
	}
// Sets the row accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrows(_:)
func (o NSWindow) SetAccessibilityRows(accessibilityRows foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRows:"), accessibilityRows)
	}
// Returns the currently selected columns for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedcolumns()
func (o NSWindow) AccessibilitySelectedColumns() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedColumns"))
	return objectivec.Object{ID: rv}
	}
// Sets the currently selected columns for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedcolumns(_:)
func (o NSWindow) SetAccessibilitySelectedColumns(accessibilitySelectedColumns foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedColumns:"), accessibilitySelectedColumns)
	}
// Returns the currently selected rows for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedrows()
func (o NSWindow) AccessibilitySelectedRows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedRows"))
	return objectivec.Object{ID: rv}
	}
// Sets the currently selected rows for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedrows(_:)
func (o NSWindow) SetAccessibilitySelectedRows(accessibilitySelectedRows foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedRows:"), accessibilitySelectedRows)
	}
// Returns the accessibility element’s sort direction.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysortdirection()
func (o NSWindow) AccessibilitySortDirection() NSAccessibilitySortDirection {
	
	rv := objc.Send[NSAccessibilitySortDirection](o.ID, objc.Sel("accessibilitySortDirection"))
	return rv
	}
// Sets the accessibility element’s sort direction.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysortdirection(_:)
func (o NSWindow) SetAccessibilitySortDirection(accessibilitySortDirection NSAccessibilitySortDirection) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySortDirection:"), accessibilitySortDirection)
	}
// Returns the visible columns for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblecolumns()
func (o NSWindow) AccessibilityVisibleColumns() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleColumns"))
	return objectivec.Object{ID: rv}
	}
// Sets the visible columns for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblecolumns(_:)
func (o NSWindow) SetAccessibilityVisibleColumns(accessibilityVisibleColumns foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleColumns:"), accessibilityVisibleColumns)
	}
// Returns the visible rows for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblerows()
func (o NSWindow) AccessibilityVisibleRows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleRows"))
	return objectivec.Object{ID: rv}
	}
// Sets the visible rows for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblerows(_:)
func (o NSWindow) SetAccessibilityVisibleRows(accessibilityVisibleRows foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleRows:"), accessibilityVisibleRows)
	}
// Returns a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilitydisclosed()
func (o NSWindow) IsAccessibilityDisclosed() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityDisclosed"))
	return rv
	}
// Sets a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosed(_:)
func (o NSWindow) SetAccessibilityDisclosed(accessibilityDisclosed bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosed:"), accessibilityDisclosed)
	}
// Returns the row disclosing the current row.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydisclosedbyrow()
func (o NSWindow) AccessibilityDisclosedByRow() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedByRow"))
	return objectivec.Object{ID: rv}
	}
// Sets the row disclosing the current row.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosedbyrow(_:)
func (o NSWindow) SetAccessibilityDisclosedByRow(accessibilityDisclosedByRow objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedByRow:"), accessibilityDisclosedByRow)
	}
// Returns the rows that the current row discloses.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydisclosedrows()
func (o NSWindow) AccessibilityDisclosedRows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedRows"))
	return objectivec.Object{ID: rv}
	}
// Sets the rows that the current row discloses.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosedrows(_:)
func (o NSWindow) SetAccessibilityDisclosedRows(accessibilityDisclosedRows objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedRows:"), accessibilityDisclosedRows)
	}
// Returns the indention level for the row.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydisclosurelevel()
func (o NSWindow) AccessibilityDisclosureLevel() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityDisclosureLevel"))
	return rv
	}
// Sets the indention level for the row.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosurelevel(_:)
func (o NSWindow) SetAccessibilityDisclosureLevel(accessibilityDisclosureLevel int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosureLevel:"), accessibilityDisclosureLevel)
	}
// Returns the column index range of the cell.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumnindexrange()
func (o NSWindow) AccessibilityColumnIndexRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityColumnIndexRange"))
	return rv
	}
// Sets the column index range of the cell.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumnindexrange(_:)
func (o NSWindow) SetAccessibilityColumnIndexRange(accessibilityColumnIndexRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnIndexRange:"), accessibilityColumnIndexRange)
	}
// Returns the row index range of the cell.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrowindexrange()
func (o NSWindow) AccessibilityRowIndexRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRowIndexRange"))
	return rv
	}
// Sets the row index range of the cell.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrowindexrange(_:)
func (o NSWindow) SetAccessibilityRowIndexRange(accessibilityRowIndexRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowIndexRange:"), accessibilityRowIndexRange)
	}
// Returns the currently selected cells for the table.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedcells()
func (o NSWindow) AccessibilitySelectedCells() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedCells"))
	return objectivec.Object{ID: rv}
	}
// Sets the currently selected cells for the table.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedcells(_:)
func (o NSWindow) SetAccessibilitySelectedCells(accessibilitySelectedCells foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedCells:"), accessibilitySelectedCells)
	}
// Returns the visible cells for the table.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblecells()
func (o NSWindow) AccessibilityVisibleCells() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleCells"))
	return objectivec.Object{ID: rv}
	}
// Sets the visible cells for the table.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblecells(_:)
func (o NSWindow) SetAccessibilityVisibleCells(accessibilityVisibleCells foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleCells:"), accessibilityVisibleCells)
	}
// Returns the cell at the specified column and row.
//
// column: The column index.
//
// row: The row index.
//
// # Return Value
// 
// The cell specified by the column and row indexes.
//
// # Discussion
// 
// This property is required for all elements that function as cell-based
// tables.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCell(forColumn:row:)
func (o NSWindow) AccessibilityCellForColumnRow(column int, row int) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCellForColumn:row:"), column, row)
	return objectivec.Object{ID: rv}
	}
// Returns the drag handle elements for the layout item element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhandles()
func (o NSWindow) AccessibilityHandles() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHandles"))
	return objectivec.Object{ID: rv}
	}
// Sets the drag handle accessibility elements for the layout item element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhandles(_:)
func (o NSWindow) SetAccessibilityHandles(accessibilityHandles foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHandles:"), accessibilityHandles)
	}
// Returns the units that the layout area uses for horizontal values.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhorizontalunits()
func (o NSWindow) AccessibilityHorizontalUnits() NSAccessibilityUnits {
	
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityHorizontalUnits"))
	return rv
	}
// Sets the units that the layout area uses for horizontal values.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhorizontalunits(_:)
func (o NSWindow) SetAccessibilityHorizontalUnits(accessibilityHorizontalUnits NSAccessibilityUnits) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalUnits:"), accessibilityHorizontalUnits)
	}
// Returns the description of the layout area’s horizontal units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhorizontalunitdescription()
func (o NSWindow) AccessibilityHorizontalUnitDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the description of the layout area’s horizontal units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhorizontalunitdescription(_:)
func (o NSWindow) SetAccessibilityHorizontalUnitDescription(accessibilityHorizontalUnitDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalUnitDescription:"), objc.String(accessibilityHorizontalUnitDescription))
	}
// Returns the units that the layout area uses for vertical values.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityverticalunits()
func (o NSWindow) AccessibilityVerticalUnits() NSAccessibilityUnits {
	
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityVerticalUnits"))
	return rv
	}
// Sets the units that the layout area uses for vertical values.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityverticalunits(_:)
func (o NSWindow) SetAccessibilityVerticalUnits(accessibilityVerticalUnits NSAccessibilityUnits) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalUnits:"), accessibilityVerticalUnits)
	}
// Returns the description of the layout area’s vertical units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityverticalunitdescription()
func (o NSWindow) AccessibilityVerticalUnitDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the description of the layout area’s vertical units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityverticalunitdescription(_:)
func (o NSWindow) SetAccessibilityVerticalUnitDescription(accessibilityVerticalUnitDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalUnitDescription:"), objc.String(accessibilityVerticalUnitDescription))
	}
// Converts the provided point in screen coordinates to a point in the layout
// area’s coordinate system.
//
// point: A point in the screen’s coordinate system.
//
// # Return Value
// 
// A point in the layout area’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLayoutPoint(forScreenPoint:)
func (o NSWindow) AccessibilityLayoutPointForScreenPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityLayoutPointForScreenPoint:"), point)
	return rv
	}
// Converts the provided size in screen coordinates to a size in the layout
// area’s coordinate system.
//
// size: A size in the screen’s coordinate system.
//
// # Return Value
// 
// A size in the layout area’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLayoutSize(forScreenSize:)
func (o NSWindow) AccessibilityLayoutSizeForScreenSize(size corefoundation.CGSize) corefoundation.CGSize {
	
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("accessibilityLayoutSizeForScreenSize:"), size)
	return rv
	}
// Converts the provided point in the layout area’s coordinates to a point
// in the screen’s coordinate system.
//
// point: A point in the layout area’s coordinate system.
//
// # Return Value
// 
// A point in the screen’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityScreenPoint(forLayoutPoint:)
func (o NSWindow) AccessibilityScreenPointForLayoutPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityScreenPointForLayoutPoint:"), point)
	return rv
	}
// Converts the provided size in the layout area’s coordinates to a size in
// the screen’s coordinate system.
//
// size: A size in the layout area’s coordinate system.
//
// # Return Value
// 
// A size in the screen’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityScreenSize(forLayoutSize:)
func (o NSWindow) AccessibilityScreenSizeForLayoutSize(size corefoundation.CGSize) corefoundation.CGSize {
	
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("accessibilityScreenSizeForLayoutSize:"), size)
	return rv
	}
// Returns the allowed values for the slider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityallowedvalues()
func (o NSWindow) AccessibilityAllowedValues() foundation.NSNumber {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAllowedValues"))
	return foundation.NSNumberFromID(rv)
	}
// Sets the allowed values for the slider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityallowedvalues(_:)
func (o NSWindow) SetAccessibilityAllowedValues(accessibilityAllowedValues foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAllowedValues:"), accessibilityAllowedValues)
	}
// Returns the child label elements for the slider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylabeluielements()
func (o NSWindow) AccessibilityLabelUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabelUIElements"))
	return objectivec.Object{ID: rv}
	}
// Sets the child label elements for the slider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylabeluielements(_:)
func (o NSWindow) SetAccessibilityLabelUIElements(accessibilityLabelUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabelUIElements:"), accessibilityLabelUIElements)
	}
// Returns the value of the label accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylabelvalue()
func (o NSWindow) AccessibilityLabelValue() float64 {
	
	rv := objc.Send[float64](o.ID, objc.Sel("accessibilityLabelValue"))
	return rv
	}
// Sets the value of the label accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylabelvalue(_:)
func (o NSWindow) SetAccessibilityLabelValue(accessibilityLabelValue float64) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabelValue:"), accessibilityLabelValue)
	}
// Returns the contents that follow the divider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitynextcontents()
func (o NSWindow) AccessibilityNextContents() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityNextContents"))
	return objectivec.Object{ID: rv}
	}
// Sets the contents that follow the divider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitynextcontents(_:)
func (o NSWindow) SetAccessibilityNextContents(accessibilityNextContents foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityNextContents:"), accessibilityNextContents)
	}
// Returns the contents that precede the divider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitypreviouscontents()
func (o NSWindow) AccessibilityPreviousContents() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPreviousContents"))
	return objectivec.Object{ID: rv}
	}
// Sets the contents that precede the divider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitypreviouscontents(_:)
func (o NSWindow) SetAccessibilityPreviousContents(accessibilityPreviousContents foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityPreviousContents:"), accessibilityPreviousContents)
	}
// Returns an array that contains the views and splitter bar from the split
// view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysplitters()
func (o NSWindow) AccessibilitySplitters() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySplitters"))
	return objectivec.Object{ID: rv}
	}
// Sets the array that contains the views and splitter bar from the split
// view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysplitters(_:)
func (o NSWindow) SetAccessibilitySplitters(accessibilitySplitters foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySplitters:"), accessibilitySplitters)
	}
// Returns the overflow button for the toolbar.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityoverflowbutton()
func (o NSWindow) AccessibilityOverflowButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityOverflowButton"))
	return objectivec.Object{ID: rv}
	}
// Sets the overflow button for the toolbar.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityoverflowbutton(_:)
func (o NSWindow) SetAccessibilityOverflowButton(accessibilityOverflowButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOverflowButton:"), accessibilityOverflowButton)
	}
// Returns the tab accessibility elements for the tab view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytabs()
func (o NSWindow) AccessibilityTabs() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTabs"))
	return objectivec.Object{ID: rv}
	}
// Sets the tab accessibility elements for the tab view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytabs(_:)
func (o NSWindow) SetAccessibilityTabs(accessibilityTabs foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTabs:"), accessibilityTabs)
	}
// Returns the user interface element that functions as a marker group for the
// ruler accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkergroupuielement()
func (o NSWindow) AccessibilityMarkerGroupUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerGroupUIElement"))
	return objectivec.Object{ID: rv}
	}
// Sets the user interface element that functions as a marker group for the
// ruler accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkergroupuielement(_:)
func (o NSWindow) SetAccessibilityMarkerGroupUIElement(accessibilityMarkerGroupUIElement objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerGroupUIElement:"), accessibilityMarkerGroupUIElement)
	}
// Returns the human-readable description of the marker type.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkertypedescription()
func (o NSWindow) AccessibilityMarkerTypeDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerTypeDescription"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the human-readable description of the marker type.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkertypedescription(_:)
func (o NSWindow) SetAccessibilityMarkerTypeDescription(accessibilityMarkerTypeDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerTypeDescription:"), objc.String(accessibilityMarkerTypeDescription))
	}
// Returns the array of marker accessibility elements for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkeruielements()
func (o NSWindow) AccessibilityMarkerUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerUIElements"))
	return objectivec.Object{ID: rv}
	}
// Sets the array of marker accessibility elements for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkeruielements(_:)
func (o NSWindow) SetAccessibilityMarkerUIElements(accessibilityMarkerUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerUIElements:"), accessibilityMarkerUIElements)
	}
// Returns the marker values for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkervalues()
func (o NSWindow) AccessibilityMarkerValues() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerValues"))
	return objectivec.Object{ID: rv}
	}
// Sets the marker values for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkervalues(_:)
func (o NSWindow) SetAccessibilityMarkerValues(accessibilityMarkerValues objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerValues:"), accessibilityMarkerValues)
	}
// Returns the type of markers for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrulermarkertype()
func (o NSWindow) AccessibilityRulerMarkerType() NSAccessibilityRulerMarkerType {
	
	rv := objc.Send[NSAccessibilityRulerMarkerType](o.ID, objc.Sel("accessibilityRulerMarkerType"))
	return rv
	}
// Sets the type of markers for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrulermarkertype(_:)
func (o NSWindow) SetAccessibilityRulerMarkerType(accessibilityRulerMarkerType NSAccessibilityRulerMarkerType) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRulerMarkerType:"), accessibilityRulerMarkerType)
	}
// Returns the units for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityunits()
func (o NSWindow) AccessibilityUnits() NSAccessibilityUnits {
	
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityUnits"))
	return rv
	}
// Sets the units used for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityunits(_:)
func (o NSWindow) SetAccessibilityUnits(accessibilityUnits NSAccessibilityUnits) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUnits:"), accessibilityUnits)
	}
// Returns the human-readable description of the ruler’s units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityunitdescription()
func (o NSWindow) AccessibilityUnitDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityUnitDescription"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the human-readable description of the ruler’s units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityunitdescription(_:)
func (o NSWindow) SetAccessibilityUnitDescription(accessibilityUnitDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUnitDescription:"), objc.String(accessibilityUnitDescription))
	}
// Returns the URL for the file that the accessibility element represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydocument()
func (o NSWindow) AccessibilityDocument() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDocument"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the URL for the file that the accessibility element represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydocument(_:)
func (o NSWindow) SetAccessibilityDocument(accessibilityDocument string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDocument:"), objc.String(accessibilityDocument))
	}
// Returns a Boolean value that indicates whether the accessibility element is
// in an edited state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityedited()
func (o NSWindow) IsAccessibilityEdited() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEdited"))
	return rv
	}
// Sets a Boolean value that indicates whether the accessibility element is in
// an edited state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityedited(_:)
func (o NSWindow) SetAccessibilityEdited(accessibilityEdited bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityEdited:"), accessibilityEdited)
	}
// Returns the filename for the file that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityfilename()
func (o NSWindow) AccessibilityFilename() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFilename"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the filename for the file that the accessibility element represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfilename(_:)
func (o NSWindow) SetAccessibilityFilename(accessibilityFilename string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFilename:"), objc.String(accessibilityFilename))
	}
// Returns the elements that have links with the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylinkeduielements()
func (o NSWindow) AccessibilityLinkedUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLinkedUIElements"))
	return objectivec.Object{ID: rv}
	}
// Sets the elements that have links with the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylinkeduielements(_:)
func (o NSWindow) SetAccessibilityLinkedUIElements(accessibilityLinkedUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLinkedUIElements:"), accessibilityLinkedUIElements)
	}
// Returns the list of elements that the accessibility element is a title for.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityservesastitleforuielements()
func (o NSWindow) AccessibilityServesAsTitleForUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityServesAsTitleForUIElements"))
	return objectivec.Object{ID: rv}
	}
// Sets the list of elements that the accessibility element is a title for.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityservesastitleforuielements(_:)
func (o NSWindow) SetAccessibilityServesAsTitleForUIElements(accessibilityServesAsTitleForUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityServesAsTitleForUIElements:"), accessibilityServesAsTitleForUIElements)
	}
// Returns the static text element that represents the accessibility
// element’s title.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytitleuielement()
func (o NSWindow) AccessibilityTitleUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitleUIElement"))
	return objectivec.Object{ID: rv}
	}
// Sets the static text element that represents the accessibility element’s
// title.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytitleuielement(_:)
func (o NSWindow) SetAccessibilityTitleUIElement(accessibilityTitleUIElement objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTitleUIElement:"), accessibilityTitleUIElement)
	}
// Returns the clear button for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityclearbutton()
func (o NSWindow) AccessibilityClearButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityClearButton"))
	return objectivec.Object{ID: rv}
	}
// Sets the clear button for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityclearbutton(_:)
func (o NSWindow) SetAccessibilityClearButton(accessibilityClearButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityClearButton:"), accessibilityClearButton)
	}
// Returns the search button for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysearchbutton()
func (o NSWindow) AccessibilitySearchButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchButton"))
	return objectivec.Object{ID: rv}
	}
// Sets the search button for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysearchbutton(_:)
func (o NSWindow) SetAccessibilitySearchButton(accessibilitySearchButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchButton:"), accessibilitySearchButton)
	}
// Returns the search menu for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysearchmenu()
func (o NSWindow) AccessibilitySearchMenu() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchMenu"))
	return objectivec.Object{ID: rv}
	}
// Sets the search menu for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysearchmenu(_:)
func (o NSWindow) SetAccessibilitySearchMenu(accessibilitySearchMenu objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchMenu:"), accessibilitySearchMenu)
	}
// Cancels the current operation.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformCancel()
func (o NSWindow) AccessibilityPerformCancel() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformCancel"))
	return rv
	}
// Simulates pressing Return in the accessibility element.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements that take keyboard input, such as a text field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformConfirm()
func (o NSWindow) AccessibilityPerformConfirm() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformConfirm"))
	return rv
	}
// Selects the accessibility element.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on selectable elements, such as a menu item.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformPick()
func (o NSWindow) AccessibilityPerformPick() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformPick"))
	return rv
	}
// Simulates clicking the accessibility element.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements that behave like buttons.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformPress()
func (o NSWindow) AccessibilityPerformPress() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformPress"))
	return rv
	}
// Displays the accessibility element’s alternative UI.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method to trigger changes to the UI due to a mouse-hover or
// similar event.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowAlternateUI()
func (o NSWindow) AccessibilityPerformShowAlternateUI() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowAlternateUI"))
	return rv
	}
// Returns to the accessibility element’s original UI.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Call this method after successfully calling
// [AccessibilityPerformShowAlternateUI] to return to the original UI.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowDefaultUI()
func (o NSWindow) AccessibilityPerformShowDefaultUI() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowDefaultUI"))
	return rv
	}
// Displays the menu accessibility element.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method to display the contextual menu for the element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowMenu()
func (o NSWindow) AccessibilityPerformShowMenu() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowMenu"))
	return rv
	}
// Brings the window to the front.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The window behaves as if you had clicked on the window’s title bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformRaise()
func (o NSWindow) AccessibilityPerformRaise() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformRaise"))
	return rv
	}
// Returns the increment button for the stepper accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityincrementbutton()
func (o NSWindow) AccessibilityIncrementButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIncrementButton"))
	return objectivec.Object{ID: rv}
	}
// Sets the increment button for the stepper accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityincrementbutton(_:)
func (o NSWindow) SetAccessibilityIncrementButton(accessibilityIncrementButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIncrementButton:"), accessibilityIncrementButton)
	}
// Returns the decrement button for the stepper accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydecrementbutton()
func (o NSWindow) AccessibilityDecrementButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDecrementButton"))
	return objectivec.Object{ID: rv}
	}
// Sets the decrement button for the stepper accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydecrementbutton(_:)
func (o NSWindow) SetAccessibilityDecrementButton(accessibilityDecrementButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDecrementButton:"), accessibilityDecrementButton)
	}
// Increments the accessibility element’s value.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements that have an adjustable [accessibilityValue]
// property.
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformIncrement()
func (o NSWindow) AccessibilityPerformIncrement() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformIncrement"))
	return rv
	}
// Decrements the accessibility element’s value.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements that have an adjustable [accessibilityValue]
// property.
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformDecrement()
func (o NSWindow) AccessibilityPerformDecrement() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformDecrement"))
	return rv
	}
// Deletes the accessibility element’s value.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements with values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformDelete()
func (o NSWindow) AccessibilityPerformDelete() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformDelete"))
	return rv
	}
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityattributeduserinputlabels()
func (o NSWindow) AccessibilityAttributedUserInputLabels() foundation.NSAttributedString {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAttributedUserInputLabels"))
	return foundation.NSAttributedStringFromID(rv)
	}
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityuserinputlabels()
func (o NSWindow) AccessibilityUserInputLabels() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityUserInputLabels"))
	return foundation.NSStringFromID(rv).String()
	}
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityattributeduserinputlabels(_:)
func (o NSWindow) SetAccessibilityAttributedUserInputLabels(accessibilityAttributedUserInputLabels foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAttributedUserInputLabels:"), accessibilityAttributedUserInputLabels)
	}
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityuserinputlabels(_:)
func (o NSWindow) SetAccessibilityUserInputLabels(accessibilityUserInputLabels foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUserInputLabels:"), accessibilityUserInputLabels)
	}

			// Protocol methods for NSAppearanceCustomization
			

			// Protocol methods for NSMenuItemValidation
			

			// Protocol methods for NSUserInterfaceItemIdentification
			

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
		return 0, ctx.Err()
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
		return 0, ctx.Err()
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

