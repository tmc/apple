// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSResponder] class.
var (
	_NSResponderClass     NSResponderClass
	_NSResponderClassOnce sync.Once
)

func getNSResponderClass() NSResponderClass {
	_NSResponderClassOnce.Do(func() {
		_NSResponderClass = NSResponderClass{class: objc.GetClass("NSResponder")}
	})
	return _NSResponderClass
}

// GetNSResponderClass returns the class object for NSResponder.
func GetNSResponderClass() NSResponderClass {
	return getNSResponderClass()
}

type NSResponderClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSResponderClass) Alloc() NSResponder {
	rv := objc.Send[NSResponder](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An abstract class that forms the basis of event and command processing in
// AppKit.
//
// # Overview
// 
// The core classes—[NSApplication], [NSWindow], and [NSView]—inherit from
// [NSResponder], as must any class that handles events. The responder model
// uses three components: event messages, action messages, and the responder
// chain.
// 
// [NSResponder] also plays an important role in the presentation of error
// information. The default implementations of the [NSResponder.PresentError] and
// [NSResponder.PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo] methods
// send [NSResponder.WillPresentError] to `self`, thereby giving subclasses the
// opportunity to customize the localized information presented in error
// alerts. [NSResponder] then forwards the message to the next responder,
// passing it the customized [NSError] object. The exact path up the modified
// responder chain depends on the type of application window:
// 
// - Window that the document owns: view > superviews > window > window
// controller > document object > document controller > the application object
// - Window with window controller but no documents: view > superviews >
// window > window controller > the application object - Window with no window
// controllers: view > superviews > window > the application object
// 
// [NSApplication] displays a document-modal error alert and, if the error
// object has a recovery attempter, gives it a chance to recover from the
// error. A recovery attempter is an object that conforms to the
// [NSErrorRecoveryAttempting] informal protocol.
//
// [NSErrorRecoveryAttempting]: https://developer.apple.com/documentation/Foundation/nserrorrecoveryattempting
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
//
// # Changing the First Responder
//
//   - [NSResponder.AcceptsFirstResponder]: A Boolean value that indicates whether the responder accepts first responder status.
//   - [NSResponder.BecomeFirstResponder]: Notifies the receiver that it’s about to become first responder in its [NSWindow](<doc://com.apple.appkit/documentation/AppKit/NSWindow>).
//   - [NSResponder.ResignFirstResponder]: Notifies the receiver that it’s been asked to relinquish its status as first responder in its window.
//   - [NSResponder.ValidateProposedFirstResponderForEvent]: Allows controls to determine when they should become first responder.
//
// # Managing the Next Responder
//
//   - [NSResponder.NextResponder]: The next responder after this one, or `nil` if it has none.
//   - [NSResponder.SetNextResponder]
//
// # Responding to Mouse Events
//
//   - [NSResponder.MouseDown]: Informs the receiver that the user has pressed the left mouse button.
//   - [NSResponder.MouseDragged]: Informs the receiver that the user has moved the mouse with the left button pressed.
//   - [NSResponder.MouseUp]: Informs the receiver that the user has released the left mouse button.
//   - [NSResponder.MouseMoved]: Informs the receiver that the mouse has moved.
//   - [NSResponder.MouseEntered]: Informs the receiver that the cursor has entered a tracking rectangle.
//   - [NSResponder.MouseExited]: Informs the receiver that the cursor has exited a tracking rectangle.
//   - [NSResponder.RightMouseDown]: Informs the receiver that the user has pressed the right mouse button.
//   - [NSResponder.RightMouseDragged]: Informs the receiver that the user has moved the mouse with the right button pressed.
//   - [NSResponder.RightMouseUp]: Informs the receiver that the user has released the right mouse button.
//   - [NSResponder.OtherMouseDown]: Informs the receiver that the user has pressed a mouse button other than the left or right one.
//   - [NSResponder.OtherMouseDragged]: Informs the receiver that the user has moved the mouse with a button other than the left or right button pressed.
//   - [NSResponder.OtherMouseUp]: Informs the receiver that the user has released a mouse button other than the left or right button.
//
// # Responding to Key Events
//
//   - [NSResponder.KeyDown]: Informs the receiver that the user has pressed a key.
//   - [NSResponder.KeyUp]: Informs the receiver that the user has released a key.
//   - [NSResponder.InterpretKeyEvents]: Handles a series of key events.
//   - [NSResponder.PerformKeyEquivalent]: Handle a key equivalent.
//   - [NSResponder.FlushBufferedKeyEvents]: Clears any unprocessed key events when overridden by subclasses.
//
// # Responding to Pressure Changes
//
//   - [NSResponder.PressureChangeWithEvent]: Indicates a pressure change as the result of a user input event on a system that supports pressure sensitivity.
//
// # Responding to Other Kinds of Events
//
//   - [NSResponder.CursorUpdate]: Informs the receiver that the mouse cursor has moved into a cursor rectangle.
//   - [NSResponder.FlagsChanged]: Informs the receiver that the user has pressed or released a modifier key (Shift, Control, and so on).
//   - [NSResponder.TabletPoint]: Informs the receiver that a tablet-point event has occurred.
//   - [NSResponder.TabletProximity]: Informs the receiver that a tablet-proximity event has occurred.
//   - [NSResponder.HelpRequested]: Displays context-sensitive help for the receiver if help has been registered.
//   - [NSResponder.ScrollWheel]: Informs the receiver that the mouse’s scroll wheel has moved.
//   - [NSResponder.QuickLookWithEvent]: Performs a Quick Look on the content at the location specified by the supplied event.
//   - [NSResponder.ChangeModeWithEvent]: Informs the responder that performed a double-tap on the side of an Apple Pencil.
//
// # Responding to Action Messages
//
//   - [NSResponder.SupplementalTargetForActionSender]: Finds a target for an action method.
//
// # Handling Window Restoration
//
//   - [NSResponder.EncodeRestorableStateWithCoder]: Saves the interface-related state of the responder.
//   - [NSResponder.EncodeRestorableStateWithCoderBackgroundQueue]: Saves the interface-related state of the responder to a keyed archiver either synchronously or asynchronously on the given operation queue.
//   - [NSResponder.RestoreStateWithCoder]: Restores the interface-related state of the responder.
//   - [NSResponder.InvalidateRestorableState]: Marks the responder’s interface-related state as dirty.
//
// # Supporting User Activities
//
//   - [NSResponder.UserActivity]: An object encapsulating a user activity supported by this responder.
//   - [NSResponder.SetUserActivity]
//   - [NSResponder.UpdateUserActivityState]: Updates the state of the given user activity.
//
// # Presenting and Customizing Error Information
//
//   - [NSResponder.PresentError]: Presents an error alert to the user as an application-modal dialog.
//   - [NSResponder.PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo]: Presents an error alert to the user as a document-modal sheet attached to document window.
//   - [NSResponder.WillPresentError]: Returns a custom version of the supplied error object that’s more suitable for presentation in alert sheets and dialogs.
//
// # Dispatching Messages
//
//   - [NSResponder.TryToPerformWith]: Attempts to perform the method indicated by an action with a specified argument.
//
// # Managing a Responder’s Menu
//
//   - [NSResponder.Menu]: Returns the responder’s menu.
//   - [NSResponder.SetMenu]
//
// # Updating the Services Menu
//
//   - [NSResponder.ValidRequestorForSendTypeReturnType]: Overridden by subclasses to determine what services are available.
//
// # Getting the Undo Manager
//
//   - [NSResponder.UndoManager]: The undo manager for this responder.
//
// # Testing Events
//
//   - [NSResponder.ShouldBeTreatedAsInkEvent]: Indicates whether a pen-down event should be treated as an ink event.
//
// # Terminating the Responder Chain
//
//   - [NSResponder.NoResponderFor]: Handles the case where an event or action message falls off the end of the responder chain.
//
// # Touch and Gesture Events
//
//   - [NSResponder.BeginGestureWithEvent]: Informs the receiver that the user has begun a touch gesture.
//   - [NSResponder.EndGestureWithEvent]: Informs the receiver that the user has ended a touch gesture.
//   - [NSResponder.MagnifyWithEvent]: Informs the receiver that the user has begun a pinch gesture.
//   - [NSResponder.RotateWithEvent]: Informs the receiver that the user has begun a rotation gesture.
//   - [NSResponder.SwipeWithEvent]: Informs the receiver that the user has begun a swipe gesture.
//   - [NSResponder.TouchesBeganWithEvent]: Informs the receiver that new set of touches has been recognized.
//   - [NSResponder.TouchesMovedWithEvent]: Informs the receiver that one or more touches has moved.
//   - [NSResponder.TouchesCancelledWithEvent]: Informs the receiver that tracking of touches has been cancelled for any reason.
//   - [NSResponder.TouchesEndedWithEvent]: Returns that a set of touches have been removed.
//   - [NSResponder.WantsForwardedScrollEventsForAxis]: Returns whether to forward elastic scrolling gesture events up the responder.
//   - [NSResponder.SmartMagnifyWithEvent]: Informs the receiver that the user performed a smart zoom gesture.
//   - [NSResponder.WantsScrollEventsForSwipeTrackingOnAxis]: Implement this method to track gesture scroll events such as a swipe.
//
// # Supporting the Touch Bar
//
//   - [NSResponder.MakeTouchBar]: Your custom subclass of the [NSResponder] class should override this method to create and configure your subclass’s default [NSTouchBar](<doc://com.apple.appkit/documentation/AppKit/NSTouchBar>) object.
//
// # Performing Text Find Actions
//
//   - [NSResponder.PerformTextFinderAction]: Performs all find oriented actions.
//
// # Supporting Tabbed Windows
//
//   - [NSResponder.NewWindowForTab]: Creates a new window to show as a tab in a tabbed window.
//
// # Creating Responders
//
//   - [NSResponder.InitWithCoder]: Creates a new responder object with data in an unarchiver.
//
// # Instance Methods
//
//   - [NSResponder.ContextMenuKeyDown]
//   - [NSResponder.MouseCancelled]
//   - [NSResponder.ShowWritingTools]
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder
type NSResponder struct {
	objectivec.Object
}

// NSResponderFromID constructs a [NSResponder] from an objc.ID.
//
// An abstract class that forms the basis of event and command processing in
// AppKit.
func NSResponderFromID(id objc.ID) NSResponder {
	return NSResponder{objectivec.Object{ID: id}}
}
// NOTE: NSResponder adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSResponder] class.
//
// # Changing the First Responder
//
//   - [INSResponder.AcceptsFirstResponder]: A Boolean value that indicates whether the responder accepts first responder status.
//   - [INSResponder.BecomeFirstResponder]: Notifies the receiver that it’s about to become first responder in its [NSWindow](<doc://com.apple.appkit/documentation/AppKit/NSWindow>).
//   - [INSResponder.ResignFirstResponder]: Notifies the receiver that it’s been asked to relinquish its status as first responder in its window.
//   - [INSResponder.ValidateProposedFirstResponderForEvent]: Allows controls to determine when they should become first responder.
//
// # Managing the Next Responder
//
//   - [INSResponder.NextResponder]: The next responder after this one, or `nil` if it has none.
//   - [INSResponder.SetNextResponder]
//
// # Responding to Mouse Events
//
//   - [INSResponder.MouseDown]: Informs the receiver that the user has pressed the left mouse button.
//   - [INSResponder.MouseDragged]: Informs the receiver that the user has moved the mouse with the left button pressed.
//   - [INSResponder.MouseUp]: Informs the receiver that the user has released the left mouse button.
//   - [INSResponder.MouseMoved]: Informs the receiver that the mouse has moved.
//   - [INSResponder.MouseEntered]: Informs the receiver that the cursor has entered a tracking rectangle.
//   - [INSResponder.MouseExited]: Informs the receiver that the cursor has exited a tracking rectangle.
//   - [INSResponder.RightMouseDown]: Informs the receiver that the user has pressed the right mouse button.
//   - [INSResponder.RightMouseDragged]: Informs the receiver that the user has moved the mouse with the right button pressed.
//   - [INSResponder.RightMouseUp]: Informs the receiver that the user has released the right mouse button.
//   - [INSResponder.OtherMouseDown]: Informs the receiver that the user has pressed a mouse button other than the left or right one.
//   - [INSResponder.OtherMouseDragged]: Informs the receiver that the user has moved the mouse with a button other than the left or right button pressed.
//   - [INSResponder.OtherMouseUp]: Informs the receiver that the user has released a mouse button other than the left or right button.
//
// # Responding to Key Events
//
//   - [INSResponder.KeyDown]: Informs the receiver that the user has pressed a key.
//   - [INSResponder.KeyUp]: Informs the receiver that the user has released a key.
//   - [INSResponder.InterpretKeyEvents]: Handles a series of key events.
//   - [INSResponder.PerformKeyEquivalent]: Handle a key equivalent.
//   - [INSResponder.FlushBufferedKeyEvents]: Clears any unprocessed key events when overridden by subclasses.
//
// # Responding to Pressure Changes
//
//   - [INSResponder.PressureChangeWithEvent]: Indicates a pressure change as the result of a user input event on a system that supports pressure sensitivity.
//
// # Responding to Other Kinds of Events
//
//   - [INSResponder.CursorUpdate]: Informs the receiver that the mouse cursor has moved into a cursor rectangle.
//   - [INSResponder.FlagsChanged]: Informs the receiver that the user has pressed or released a modifier key (Shift, Control, and so on).
//   - [INSResponder.TabletPoint]: Informs the receiver that a tablet-point event has occurred.
//   - [INSResponder.TabletProximity]: Informs the receiver that a tablet-proximity event has occurred.
//   - [INSResponder.HelpRequested]: Displays context-sensitive help for the receiver if help has been registered.
//   - [INSResponder.ScrollWheel]: Informs the receiver that the mouse’s scroll wheel has moved.
//   - [INSResponder.QuickLookWithEvent]: Performs a Quick Look on the content at the location specified by the supplied event.
//   - [INSResponder.ChangeModeWithEvent]: Informs the responder that performed a double-tap on the side of an Apple Pencil.
//
// # Responding to Action Messages
//
//   - [INSResponder.SupplementalTargetForActionSender]: Finds a target for an action method.
//
// # Handling Window Restoration
//
//   - [INSResponder.EncodeRestorableStateWithCoder]: Saves the interface-related state of the responder.
//   - [INSResponder.EncodeRestorableStateWithCoderBackgroundQueue]: Saves the interface-related state of the responder to a keyed archiver either synchronously or asynchronously on the given operation queue.
//   - [INSResponder.RestoreStateWithCoder]: Restores the interface-related state of the responder.
//   - [INSResponder.InvalidateRestorableState]: Marks the responder’s interface-related state as dirty.
//
// # Supporting User Activities
//
//   - [INSResponder.UserActivity]: An object encapsulating a user activity supported by this responder.
//   - [INSResponder.SetUserActivity]
//   - [INSResponder.UpdateUserActivityState]: Updates the state of the given user activity.
//
// # Presenting and Customizing Error Information
//
//   - [INSResponder.PresentError]: Presents an error alert to the user as an application-modal dialog.
//   - [INSResponder.PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo]: Presents an error alert to the user as a document-modal sheet attached to document window.
//   - [INSResponder.WillPresentError]: Returns a custom version of the supplied error object that’s more suitable for presentation in alert sheets and dialogs.
//
// # Dispatching Messages
//
//   - [INSResponder.TryToPerformWith]: Attempts to perform the method indicated by an action with a specified argument.
//
// # Managing a Responder’s Menu
//
//   - [INSResponder.Menu]: Returns the responder’s menu.
//   - [INSResponder.SetMenu]
//
// # Updating the Services Menu
//
//   - [INSResponder.ValidRequestorForSendTypeReturnType]: Overridden by subclasses to determine what services are available.
//
// # Getting the Undo Manager
//
//   - [INSResponder.UndoManager]: The undo manager for this responder.
//
// # Testing Events
//
//   - [INSResponder.ShouldBeTreatedAsInkEvent]: Indicates whether a pen-down event should be treated as an ink event.
//
// # Terminating the Responder Chain
//
//   - [INSResponder.NoResponderFor]: Handles the case where an event or action message falls off the end of the responder chain.
//
// # Touch and Gesture Events
//
//   - [INSResponder.BeginGestureWithEvent]: Informs the receiver that the user has begun a touch gesture.
//   - [INSResponder.EndGestureWithEvent]: Informs the receiver that the user has ended a touch gesture.
//   - [INSResponder.MagnifyWithEvent]: Informs the receiver that the user has begun a pinch gesture.
//   - [INSResponder.RotateWithEvent]: Informs the receiver that the user has begun a rotation gesture.
//   - [INSResponder.SwipeWithEvent]: Informs the receiver that the user has begun a swipe gesture.
//   - [INSResponder.TouchesBeganWithEvent]: Informs the receiver that new set of touches has been recognized.
//   - [INSResponder.TouchesMovedWithEvent]: Informs the receiver that one or more touches has moved.
//   - [INSResponder.TouchesCancelledWithEvent]: Informs the receiver that tracking of touches has been cancelled for any reason.
//   - [INSResponder.TouchesEndedWithEvent]: Returns that a set of touches have been removed.
//   - [INSResponder.WantsForwardedScrollEventsForAxis]: Returns whether to forward elastic scrolling gesture events up the responder.
//   - [INSResponder.SmartMagnifyWithEvent]: Informs the receiver that the user performed a smart zoom gesture.
//   - [INSResponder.WantsScrollEventsForSwipeTrackingOnAxis]: Implement this method to track gesture scroll events such as a swipe.
//
// # Supporting the Touch Bar
//
//   - [INSResponder.MakeTouchBar]: Your custom subclass of the [NSResponder] class should override this method to create and configure your subclass’s default [NSTouchBar](<doc://com.apple.appkit/documentation/AppKit/NSTouchBar>) object.
//
// # Performing Text Find Actions
//
//   - [INSResponder.PerformTextFinderAction]: Performs all find oriented actions.
//
// # Supporting Tabbed Windows
//
//   - [INSResponder.NewWindowForTab]: Creates a new window to show as a tab in a tabbed window.
//
// # Creating Responders
//
//   - [INSResponder.InitWithCoder]: Creates a new responder object with data in an unarchiver.
//
// # Instance Methods
//
//   - [INSResponder.ContextMenuKeyDown]
//   - [INSResponder.MouseCancelled]
//   - [INSResponder.ShowWritingTools]
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder
type INSResponder interface {
	objectivec.IObject
	NSStandardKeyBindingResponding
	NSTouchBarProvider
	NSUserActivityRestoring

	// Topic: Changing the First Responder

	// A Boolean value that indicates whether the responder accepts first responder status.
	AcceptsFirstResponder() bool
	// Notifies the receiver that it’s about to become first responder in its [NSWindow](<doc://com.apple.appkit/documentation/AppKit/NSWindow>).
	BecomeFirstResponder() bool
	// Notifies the receiver that it’s been asked to relinquish its status as first responder in its window.
	ResignFirstResponder() bool
	// Allows controls to determine when they should become first responder.
	ValidateProposedFirstResponderForEvent(responder INSResponder, event INSEvent) bool

	// Topic: Managing the Next Responder

	// The next responder after this one, or `nil` if it has none.
	NextResponder() INSResponder
	SetNextResponder(value INSResponder)

	// Topic: Responding to Mouse Events

	// Informs the receiver that the user has pressed the left mouse button.
	MouseDown(event INSEvent)
	// Informs the receiver that the user has moved the mouse with the left button pressed.
	MouseDragged(event INSEvent)
	// Informs the receiver that the user has released the left mouse button.
	MouseUp(event INSEvent)
	// Informs the receiver that the mouse has moved.
	MouseMoved(event INSEvent)
	// Informs the receiver that the cursor has entered a tracking rectangle.
	MouseEntered(event INSEvent)
	// Informs the receiver that the cursor has exited a tracking rectangle.
	MouseExited(event INSEvent)
	// Informs the receiver that the user has pressed the right mouse button.
	RightMouseDown(event INSEvent)
	// Informs the receiver that the user has moved the mouse with the right button pressed.
	RightMouseDragged(event INSEvent)
	// Informs the receiver that the user has released the right mouse button.
	RightMouseUp(event INSEvent)
	// Informs the receiver that the user has pressed a mouse button other than the left or right one.
	OtherMouseDown(event INSEvent)
	// Informs the receiver that the user has moved the mouse with a button other than the left or right button pressed.
	OtherMouseDragged(event INSEvent)
	// Informs the receiver that the user has released a mouse button other than the left or right button.
	OtherMouseUp(event INSEvent)

	// Topic: Responding to Key Events

	// Informs the receiver that the user has pressed a key.
	KeyDown(event INSEvent)
	// Informs the receiver that the user has released a key.
	KeyUp(event INSEvent)
	// Handles a series of key events.
	InterpretKeyEvents(eventArray []NSEvent)
	// Handle a key equivalent.
	PerformKeyEquivalent(event INSEvent) bool
	// Clears any unprocessed key events when overridden by subclasses.
	FlushBufferedKeyEvents()

	// Topic: Responding to Pressure Changes

	// Indicates a pressure change as the result of a user input event on a system that supports pressure sensitivity.
	PressureChangeWithEvent(event INSEvent)

	// Topic: Responding to Other Kinds of Events

	// Informs the receiver that the mouse cursor has moved into a cursor rectangle.
	CursorUpdate(event INSEvent)
	// Informs the receiver that the user has pressed or released a modifier key (Shift, Control, and so on).
	FlagsChanged(event INSEvent)
	// Informs the receiver that a tablet-point event has occurred.
	TabletPoint(event INSEvent)
	// Informs the receiver that a tablet-proximity event has occurred.
	TabletProximity(event INSEvent)
	// Displays context-sensitive help for the receiver if help has been registered.
	HelpRequested(eventPtr INSEvent)
	// Informs the receiver that the mouse’s scroll wheel has moved.
	ScrollWheel(event INSEvent)
	// Performs a Quick Look on the content at the location specified by the supplied event.
	QuickLookWithEvent(event INSEvent)
	// Informs the responder that performed a double-tap on the side of an Apple Pencil.
	ChangeModeWithEvent(event INSEvent)

	// Topic: Responding to Action Messages

	// Finds a target for an action method.
	SupplementalTargetForActionSender(action objc.SEL, sender objectivec.IObject) objectivec.IObject

	// Topic: Handling Window Restoration

	// Saves the interface-related state of the responder.
	EncodeRestorableStateWithCoder(coder foundation.INSCoder)
	// Saves the interface-related state of the responder to a keyed archiver either synchronously or asynchronously on the given operation queue.
	EncodeRestorableStateWithCoderBackgroundQueue(coder foundation.INSCoder, queue foundation.NSOperationQueue)
	// Restores the interface-related state of the responder.
	RestoreStateWithCoder(coder foundation.INSCoder)
	// Marks the responder’s interface-related state as dirty.
	InvalidateRestorableState()

	// Topic: Supporting User Activities

	// An object encapsulating a user activity supported by this responder.
	UserActivity() foundation.NSUserActivity
	SetUserActivity(value foundation.NSUserActivity)
	// Updates the state of the given user activity.
	UpdateUserActivityState(userActivity foundation.NSUserActivity)

	// Topic: Presenting and Customizing Error Information

	// Presents an error alert to the user as an application-modal dialog.
	PresentError(error_ foundation.INSError) bool
	// Presents an error alert to the user as a document-modal sheet attached to document window.
	PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error_ foundation.INSError, window INSWindow, delegate objectivec.IObject, didPresentSelector objc.SEL, contextInfo uintptr)
	// Returns a custom version of the supplied error object that’s more suitable for presentation in alert sheets and dialogs.
	WillPresentError(error_ foundation.INSError) foundation.INSError

	// Topic: Dispatching Messages

	// Attempts to perform the method indicated by an action with a specified argument.
	TryToPerformWith(action objc.SEL, object objectivec.IObject) bool

	// Topic: Managing a Responder’s Menu

	// Returns the responder’s menu.
	Menu() INSMenu
	SetMenu(value INSMenu)

	// Topic: Updating the Services Menu

	// Overridden by subclasses to determine what services are available.
	ValidRequestorForSendTypeReturnType(sendType NSPasteboardType, returnType NSPasteboardType) objectivec.IObject

	// Topic: Getting the Undo Manager

	// The undo manager for this responder.
	UndoManager() foundation.NSUndoManager

	// Topic: Testing Events

	// Indicates whether a pen-down event should be treated as an ink event.
	ShouldBeTreatedAsInkEvent(event INSEvent) bool

	// Topic: Terminating the Responder Chain

	// Handles the case where an event or action message falls off the end of the responder chain.
	NoResponderFor(eventSelector objc.SEL)

	// Topic: Touch and Gesture Events

	// Informs the receiver that the user has begun a touch gesture.
	BeginGestureWithEvent(event INSEvent)
	// Informs the receiver that the user has ended a touch gesture.
	EndGestureWithEvent(event INSEvent)
	// Informs the receiver that the user has begun a pinch gesture.
	MagnifyWithEvent(event INSEvent)
	// Informs the receiver that the user has begun a rotation gesture.
	RotateWithEvent(event INSEvent)
	// Informs the receiver that the user has begun a swipe gesture.
	SwipeWithEvent(event INSEvent)
	// Informs the receiver that new set of touches has been recognized.
	TouchesBeganWithEvent(event INSEvent)
	// Informs the receiver that one or more touches has moved.
	TouchesMovedWithEvent(event INSEvent)
	// Informs the receiver that tracking of touches has been cancelled for any reason.
	TouchesCancelledWithEvent(event INSEvent)
	// Returns that a set of touches have been removed.
	TouchesEndedWithEvent(event INSEvent)
	// Returns whether to forward elastic scrolling gesture events up the responder.
	WantsForwardedScrollEventsForAxis(axis NSEventGestureAxis) bool
	// Informs the receiver that the user performed a smart zoom gesture.
	SmartMagnifyWithEvent(event INSEvent)
	// Implement this method to track gesture scroll events such as a swipe.
	WantsScrollEventsForSwipeTrackingOnAxis(axis NSEventGestureAxis) bool

	// Topic: Supporting the Touch Bar

	// Your custom subclass of the [NSResponder] class should override this method to create and configure your subclass’s default [NSTouchBar](<doc://com.apple.appkit/documentation/AppKit/NSTouchBar>) object.
	MakeTouchBar() INSTouchBar

	// Topic: Performing Text Find Actions

	// Performs all find oriented actions.
	PerformTextFinderAction(sender objectivec.IObject)

	// Topic: Supporting Tabbed Windows

	// Creates a new window to show as a tab in a tabbed window.
	NewWindowForTab(sender objectivec.IObject)

	// Topic: Creating Responders

	// Creates a new responder object with data in an unarchiver.
	InitWithCoder(coder foundation.INSCoder) NSResponder

	// Topic: Instance Methods

	ContextMenuKeyDown(event INSEvent)
	MouseCancelled(event INSEvent)
	ShowWritingTools(sender objectivec.IObject)

	// Implemented by subclasses to invoke the help system, displaying information relevant to the receiver and its current state.
	ShowContextHelp(sender objectivec.IObject)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (r NSResponder) Init() NSResponder {
	rv := objc.Send[NSResponder](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r NSResponder) Autorelease() NSResponder {
	rv := objc.Send[NSResponder](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSResponder creates a new NSResponder instance.
func NewNSResponder() NSResponder {
	class := getNSResponderClass()
	rv := objc.Send[NSResponder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new responder object with data in an unarchiver.
//
// coder: An unarchiver object.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/init(coder:)
func NewResponderWithCoder(coder foundation.INSCoder) NSResponder {
	instance := getNSResponderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSResponderFromID(rv)
}

// Notifies the receiver that it’s about to become first responder in its
// [NSWindow].
//
// # Discussion
// 
// The default implementation returns [true], accepting first responder
// status. Subclasses can override this method to update state or perform some
// action such as highlighting the selection, or to return [false], refusing
// first responder status.
// 
// Use the [NSWindow] [FirstResponder] method, not this method, to make an
// object the first responder. Never invoke this method directly.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/becomeFirstResponder()
func (r NSResponder) BecomeFirstResponder() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("becomeFirstResponder"))
	return rv
}
// Notifies the receiver that it’s been asked to relinquish its status as
// first responder in its window.
//
// # Discussion
// 
// The default implementation returns [true], resigning first responder
// status. Subclasses can override this method to update state or perform some
// action such as unhighlighting the selection, or to return [false], refusing
// to relinquish first responder status.
// 
// Use the [NSWindow] [FirstResponder] method, not this method, to make an
// object the first responder. Never invoke this method directly.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/resignFirstResponder()
func (r NSResponder) ResignFirstResponder() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("resignFirstResponder"))
	return rv
}
// Allows controls to determine when they should become first responder.
//
// responder: The first responder.
//
// event: The event to validate. May be `nil` if there is no applicable event.
//
// # Return Value
// 
// [true] if the control should become first responder, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Some controls, such as [NSTextField], should only become first responder
// when the enclosing NSTableView/NSBrowser indicates that the view can begin
// editing. It is up to the particular control that wants to be validated to
// call this method in its [MouseDown] method (or perhaps at another time) to
// determine if it should attempt to become the first responder or not.
// 
// The [NSTableView], [NSOutlineView], and [NSBrowser] classes implement this
// to allow first responder status only if the responder is a view in a
// selected row. It also delays the first responder assignment if a
// `doubleAction` may be invoked.
// 
// The default implementation returns [true] when there is no [NextResponder]
// set, otherwise, it is forwarded up the responder chain.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/validateProposedFirstResponder(_:for:)
func (r NSResponder) ValidateProposedFirstResponderForEvent(responder INSResponder, event INSEvent) bool {
	rv := objc.Send[bool](r.ID, objc.Sel("validateProposedFirstResponder:forEvent:"), responder, event)
	return rv
}
// Informs the receiver that the user has pressed the left mouse button.
//
// event: An object encapsulating information about the mouse-down event.
//
// # Discussion
// 
// The default implementation simply passes this message to the next
// responder.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/mouseDown(with:)
func (r NSResponder) MouseDown(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("mouseDown:"), event)
}
// Informs the receiver that the user has moved the mouse with the left button
// pressed.
//
// event: An object encapsulating information about the mouse-dragged event.
//
// # Discussion
// 
// The default implementation simply passes this message to the next
// responder.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/mouseDragged(with:)
func (r NSResponder) MouseDragged(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("mouseDragged:"), event)
}
// Informs the receiver that the user has released the left mouse button.
//
// event: An object encapsulating information about the mouse-up event.
//
// # Discussion
// 
// The default implementation simply passes this message to the next
// responder.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/mouseUp(with:)
func (r NSResponder) MouseUp(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("mouseUp:"), event)
}
// Informs the receiver that the mouse has moved.
//
// event: An object encapsulating information about the mouse-moved event.
//
// # Discussion
// 
// The default implementation simply passes this message to the next
// responder.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/mouseMoved(with:)
func (r NSResponder) MouseMoved(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("mouseMoved:"), event)
}
// Informs the receiver that the cursor has entered a tracking rectangle.
//
// event: An object encapsulating information about the mouse-entered event.
//
// # Discussion
// 
// The default implementation simply passes this message to the next
// responder.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/mouseEntered(with:)
func (r NSResponder) MouseEntered(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("mouseEntered:"), event)
}
// Informs the receiver that the cursor has exited a tracking rectangle.
//
// event: An object encapsulating information about the mouse-exited event.
//
// # Discussion
// 
// The default implementation simply passes this message to the next
// responder.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/mouseExited(with:)
func (r NSResponder) MouseExited(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("mouseExited:"), event)
}
// Informs the receiver that the user has pressed the right mouse button.
//
// event: An object encapsulating information about the mouse-down event.
//
// # Discussion
// 
// The default implementation simply passes this message to the next
// responder.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/rightMouseDown(with:)
func (r NSResponder) RightMouseDown(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("rightMouseDown:"), event)
}
// Informs the receiver that the user has moved the mouse with the right
// button pressed.
//
// event: An object encapsulating information about the mouse-dragged event.
//
// # Discussion
// 
// The default implementation simply passes this message to the next
// responder.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/rightMouseDragged(with:)
func (r NSResponder) RightMouseDragged(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("rightMouseDragged:"), event)
}
// Informs the receiver that the user has released the right mouse button.
//
// event: An object encapsulating information about the mouse-up event.
//
// # Discussion
// 
// The default implementation simply passes this message to the next
// responder.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/rightMouseUp(with:)
func (r NSResponder) RightMouseUp(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("rightMouseUp:"), event)
}
// Informs the receiver that the user has pressed a mouse button other than
// the left or right one.
//
// event: An object encapsulating information about the mouse-down event.
//
// # Discussion
// 
// The default implementation simply passes this message to the next
// responder.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/otherMouseDown(with:)
func (r NSResponder) OtherMouseDown(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("otherMouseDown:"), event)
}
// Informs the receiver that the user has moved the mouse with a button other
// than the left or right button pressed.
//
// event: An object encapsulating information about the mouse-dragged event.
//
// # Discussion
// 
// The default implementation simply passes this message to the next
// responder.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/otherMouseDragged(with:)
func (r NSResponder) OtherMouseDragged(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("otherMouseDragged:"), event)
}
// Informs the receiver that the user has released a mouse button other than
// the left or right button.
//
// event: An object encapsulating information about the mouse-up event.
//
// # Discussion
// 
// The default implementation simply passes this message to the next
// responder.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/otherMouseUp(with:)
func (r NSResponder) OtherMouseUp(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("otherMouseUp:"), event)
}
// Informs the receiver that the user has pressed a key.
//
// event: An object encapsulating information about the key-down event.
//
// # Discussion
// 
// The receiver can interpret `event` itself, or pass it to the system input
// manager using [InterpretKeyEvents]. The default implementation simply
// passes this message to the next responder.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/keyDown(with:)
func (r NSResponder) KeyDown(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("keyDown:"), event)
}
// Informs the receiver that the user has released a key.
//
// event: An object encapsulating information about the key-up event.
//
// # Discussion
// 
// The default implementation simply passes this message to the next
// responder.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/keyUp(with:)
func (r NSResponder) KeyUp(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("keyUp:"), event)
}
// Handles a series of key events.
//
// eventArray: An array of key-event characters to give to the system input manager.
//
// # Discussion
// 
// This method, which is invoked by subclasses from the [KeyDown] method,
// sends the character input in `eventArray` to the system input manager for
// interpretation as text to insert or commands to perform. The input manager
// responds to the request by sending [InsertText] and [DoCommandBySelector]
// messages back to the invoker of this method. Subclasses shouldn’t
// override this method.
// 
// See the [NSInputManager] and [NSTextInput] class and protocol
// specifications for more information on input management.
//
// [NSInputManager]: https://developer.apple.com/documentation/AppKit/NSInputManager
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/interpretKeyEvents(_:)
func (r NSResponder) InterpretKeyEvents(eventArray []NSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("interpretKeyEvents:"), objectivec.IObjectSliceToNSArray(eventArray))
}
// Handle a key equivalent.
//
// event: An event object that represents the key equivalent pressed.
//
// # Discussion
// 
// Override to handle key equivalents. If the character code or codes in
// `event` match the receiver’s key equivalent, the receiver should respond
// to the event and return [true]. The default implementation does nothing and
// returns [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/performKeyEquivalent(with:)
func (r NSResponder) PerformKeyEquivalent(event INSEvent) bool {
	rv := objc.Send[bool](r.ID, objc.Sel("performKeyEquivalent:"), event)
	return rv
}
// Clears any unprocessed key events when overridden by subclasses.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/flushBufferedKeyEvents()
func (r NSResponder) FlushBufferedKeyEvents() {
	objc.Send[objc.ID](r.ID, objc.Sel("flushBufferedKeyEvents"))
}
// Indicates a pressure change as the result of a user input event on a system
// that supports pressure sensitivity.
//
// event: An [NSEvent] object encapsulating information about the event that invoked
// the change in pressure.
//
// # Discussion
// 
// This method is invoked automatically in response to user actions. `event`
// is the event that initiated the change in pressure.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/pressureChange(with:)
func (r NSResponder) PressureChangeWithEvent(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("pressureChangeWithEvent:"), event)
}
// Informs the receiver that the mouse cursor has moved into a cursor
// rectangle.
//
// event: An object encapsulating information about the cursor-update event
// ([NSCursorUpdate]).
// //
// [NSCursorUpdate]: https://developer.apple.com/documentation/AppKit/NSCursorUpdate
//
// # Discussion
// 
// Override this method to set the cursor image. The default implementation
// uses cursor rectangles, if cursor rectangles are currently valid. If they
// are not, it calls `super` to send the message up the responder chain.
// 
// If the responder implements this method, but decides not to handle a
// particular event, it should invoke the superclass implementation of this
// method.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/cursorUpdate(with:)
func (r NSResponder) CursorUpdate(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("cursorUpdate:"), event)
}
// Informs the receiver that the user has pressed or released a modifier key
// (Shift, Control, and so on).
//
// event: An object encapsulating information about the modifier-key event.
//
// # Discussion
// 
// The default implementation simply passes this message to the next
// responder.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/flagsChanged(with:)
func (r NSResponder) FlagsChanged(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("flagsChanged:"), event)
}
// Informs the receiver that a tablet-point event has occurred.
//
// event: An object encapsulating information about the tablet-point event.
//
// # Discussion
// 
// Tablet events are represented by [NSEvent] objects of type [NSTabletPoint].
// They describe the current state of a transducer (that is, a pointing
// device) that is in proximity to its tablet, reflecting changes such as
// location, pressure, tilt, and rotation. See the [NSEvent] reference for the
// methods that allow you to extract this and other information from `event`.
// The default implementation of [NSResponder] passes the message to the next
// responder.
//
// [NSTabletPoint]: https://developer.apple.com/documentation/AppKit/NSTabletPoint
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/tabletPoint(with:)
func (r NSResponder) TabletPoint(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("tabletPoint:"), event)
}
// Informs the receiver that a tablet-proximity event has occurred.
//
// event: An object encapsulating information about the tablet-point event.
//
// # Discussion
// 
// Tablet events are represented by [NSEvent] objects of type
// [NSTabletProximity]. Tablet devices generate proximity events when the
// transducer (pointing device) nears a tablet and when it moves away from a
// tablet. From an event object of this type you can extract information about
// the kind of device and its capabilities, as well as the relation of this
// tablet-proximity event to various tablet-point events; see the [NSEvent]
// reference for details. The default implementation passes the message to the
// next responder.
//
// [NSTabletProximity]: https://developer.apple.com/documentation/AppKit/NSTabletProximity
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/tabletProximity(with:)
func (r NSResponder) TabletProximity(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("tabletProximity:"), event)
}
// Displays context-sensitive help for the receiver if help has been
// registered.
//
// eventPtr: An object encapsulating information about the help-request event.
//
// # Discussion
// 
// [NSWindow] invokes this method automatically when the user clicks for help
// and help has been registered using [SetContextHelpForObject]. Otherwise,
// [NSWindow] passes the message to the next responder. Subclasses are not
// required to override this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/helpRequested(_:)
func (r NSResponder) HelpRequested(eventPtr INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("helpRequested:"), eventPtr)
}
// Informs the receiver that the mouse’s scroll wheel has moved.
//
// event: An object encapsulating information about the wheel-scrolling event.
//
// # Discussion
// 
// The default implementation passes this message to the next responder.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/scrollWheel(with:)
func (r NSResponder) ScrollWheel(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("scrollWheel:"), event)
}
// Performs a Quick Look on the content at the location specified by the
// supplied event.
//
// event: An event object containing the location of the Quick Look content.
//
// # Discussion
// 
// The [NSEvent.EventType.quickLook] event type supports this method. The only
// valid properties of an [NSEvent.EventType.quickLook] event are
// [LocationInWindow] and [ModifierFlags]. A Quick Look event does not come in
// through the normal event mechanism; therefore, there is no corresponding
// event mask for it, nor should you attempt to look for it in a [SendEvent]
// message or with the [NextEventMatchingMask] methods.
// 
// If there are no Quick Look items at the location, call super.
// 
// [NSResponder] declares but doesn’t implement this method.
//
// [NSEvent.EventType.quickLook]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/quickLook
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/quickLook(with:)
func (r NSResponder) QuickLookWithEvent(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("quickLookWithEvent:"), event)
}
// Informs the responder that performed a double-tap on the side of an Apple
// Pencil.
//
// event: An object encapsulating information about the change-mode event.
//
// # Discussion
// 
// The default implementation passes the event to the next responder.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/changeMode(with:)
func (r NSResponder) ChangeModeWithEvent(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("changeModeWithEvent:"), event)
}
// Finds a target for an action method.
//
// action: The requested action.
//
// sender: The message sender.
//
// # Return Value
// 
// An object which responds to the action, or `nil`.
//
// # Discussion
// 
// If this [NSResponder] instance does not itself ``, then `` is called.
// 
// This method should return an object which responds to the action; if this
// responder does not have a supplemental object that does that, the
// implementation of this method should call `super`’s ``.
// 
// NSResponder’s implementation returns `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/supplementalTarget(forAction:sender:)
func (r NSResponder) SupplementalTargetForActionSender(action objc.SEL, sender objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("supplementalTargetForAction:sender:"), action, sender)
	return objectivec.Object{ID: rv}
}
// Saves the interface-related state of the responder.
//
// coder: The coder object in which to save the responder’s interface-related
// state.
//
// # Discussion
// 
// This method is part of the window restoration system and is called at
// appropriate times to save the visual state of your responder to the
// specified archive. The default implementation of this method does nothing
// but specific subclasses (such as [NSView] and [NSWindow]) override it to
// save important state information. Therefore, if you override this method,
// you should always call `super` at some point in your implementation.
// 
// Subclasses can override this method and use it to restore any information
// that would be needed to restore the responder to its current state. For
// example, the [NSTabView] class uses this method to save information about
// the currently selected tab. You must store enough data to reconfigure the
// responder and return it to its current state during a subsequent launch of
// the application.
// 
// For information about using a coder object to write data to an archive, see
// [Archives and Serializations Programming Guide].
//
// [Archives and Serializations Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Archiving/Archiving.html#//apple_ref/doc/uid/10000047i
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/encodeRestorableState(with:)
func (r NSResponder) EncodeRestorableStateWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeRestorableStateWithCoder:"), coder)
}
// Saves the interface-related state of the responder to a keyed archiver
// either synchronously or asynchronously on the given operation queue.
//
// coder: A thread-safe keyed archiver to write restorable state into.
//
// queue: A serial background operation queue to perform asynchronous encoding work
// on.
//
// # Discussion
// 
// Don’t call this method directly. The system calls this method on the main
// thread. The receiver may synchronously encode state to `coder` or may
// enqueue asynchronous work to encode additional restorable state using the
// provided serial background [OperationQueue], if it can safely access and
// encode that information. The encoding process finishes when the enqueued
// operations complete.
// 
// If you override this method, call the super implementation.
//
// [OperationQueue]: https://developer.apple.com/documentation/Foundation/OperationQueue
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/encodeRestorableState(with:backgroundQueue:)
func (r NSResponder) EncodeRestorableStateWithCoderBackgroundQueue(coder foundation.INSCoder, queue foundation.NSOperationQueue) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeRestorableStateWithCoder:backgroundQueue:"), coder, queue)
}
// Restores the interface-related state of the responder.
//
// coder: The coder object to use to restore the responder’s interface-related
// state.
//
// # Discussion
// 
// This method is part of the window restoration system and is called at
// launch time to restore the visual state of your responder object. The
// default implementation does nothing but specific subclasses (such as
// [NSView] and [NSWindow]) override it and save important state information.
// Therefore, if you override this method, you should always call `super` at
// some point in your implementation.
// 
// Subclasses can override this method and use it to restore any information
// that was saved in the [EncodeRestorableStateWithCoder] method. You can also
// use this method to reconfigure the responder to its previous appearance.
// 
// For information about using a coder object to read data from an archive,
// see [Encoding and Decoding Custom Types].
//
// [Encoding and Decoding Custom Types]: https://developer.apple.com/documentation/Foundation/encoding-and-decoding-custom-types
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/restoreState(with:)
func (r NSResponder) RestoreStateWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](r.ID, objc.Sel("restoreStateWithCoder:"), coder)
}
// Marks the responder’s interface-related state as dirty.
//
// # Discussion
// 
// Call this method whenever the restorable state of your responder changes.
// This method marks the responder’s state as dirty, which writes the state
// to disk at some point in the future. Don’t override this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/invalidateRestorableState()
func (r NSResponder) InvalidateRestorableState() {
	objc.Send[objc.ID](r.ID, objc.Sel("invalidateRestorableState"))
}
// Updates the state of the given user activity.
//
// userActivity: The user activity to be updated.
//
// # Discussion
// 
// Subclasses override this method to update the state of the supplied
// `userActivity`. Add state representing the user’s activity into
// `userActivity` using its [addUserInfoEntries(from:)] method. When the state
// is dirty, set the [needsSave] property to [true].
// 
// When an [NSUserActivity] object managed by AppKit is updated, an empty
// `userInfo` dictionary is given to the user activity, and all of the objects
// associated with the user activity are sent an [NSResponder] message.
//
// [NSUserActivity]: https://developer.apple.com/documentation/Foundation/NSUserActivity
// [addUserInfoEntries(from:)]: https://developer.apple.com/documentation/Foundation/NSUserActivity/addUserInfoEntries(from:)
// [needsSave]: https://developer.apple.com/documentation/Foundation/NSUserActivity/needsSave
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/updateUserActivityState(_:)
func (r NSResponder) UpdateUserActivityState(userActivity foundation.NSUserActivity) {
	objc.Send[objc.ID](r.ID, objc.Sel("updateUserActivityState:"), userActivity)
}
// Presents an error alert to the user as an application-modal dialog.
//
// error: An object containing information about an error.
//
// # Discussion
// 
// The alert displays information found in the [NSError] object `error`; this
// information can include error description, recovery suggestion, failure
// reason, and button titles (all localized). The method returns [true] if
// error recovery succeeded and [false] otherwise. For error recovery to be
// attempted, an recovery-attempter object (that is, an object conforming to
// the [NSErrorRecoveryAttempting] informal protocol) must be associated with
// `error`.
// 
// The default implementation of this method sends [WillPresentError] to
// `self`. By doing this, [NSResponder] gives subclasses an opportunity to
// customize error presentation. It then forwards the message, passing any
// customized error object, to the next responder; if there is no next
// responder, it passes the error object to [NSApp], which displays a
// document-modal error alert. When the user dismisses the alert, any recovery
// attempter associated with the error object is given a chance to recover
// from the error. See the class description for the precise route up the
// responder chain (plus document and controller objects) this message might
// travel.
// 
// It is not recommended that you attempt to override this method. If you wish
// to customize the error presentation, override [WillPresentError] instead.
//
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/presentError(_:)
func (r NSResponder) PresentError(error_ foundation.INSError) bool {
	rv := objc.Send[bool](r.ID, objc.Sel("presentError:"), error_)
	return rv
}
// Presents an error alert to the user as a document-modal sheet attached to
// document window.
//
// error: The object encapsulating information about the error.
//
// window: The window object identifying the window owning the document-modal sheet.
//
// delegate: The modal delegate for the sheet.
//
// didPresentSelector: A selector identifying the message to be sent to the modal delegate. The
// `didPresentSelector` selector must have the signature:
//
// contextInfo: Supplemental data to be passed to the modal delegate; can be [NULL].
//
// # Discussion
// 
// The information displayed in the alert is extracted from the [NSError]
// object `error`; it may include a description, recovery suggestion, failure
// reason, and button titles (all localized). Once the user dismisses the
// alert and any recovery attempter associated with the error object has had a
// chance to recover from it, the receiver sends a message identified by
// `didPresentSelector` to the modal delegate `delegate`. (A recovery
// attempter is an object that conforms to the NSErrorRecoveryAttempting
// informal protocol.)
// 
// The modal delegate implements the method identified by `didPresentSelector`
// to perform any post-error processing if recovery failed or was not
// attempted (that is, `didRecover` is [false]). Any supplemental data is
// passed to the modal delegate via `contextInfo`.
// 
// The default implementation of this method sends [WillPresentError] to
// `self`. By doing this, [NSResponder] gives subclasses an opportunity to
// customize error presentation. It then forwards the message, passing any
// customized error to the next responder or, if there is no next responder,
// it passes the error object to [NSApp], which displays a document-modal
// error alert. When the user dismisses the alert, any recovery attempter
// associated with the error object is given a chance to recover from the
// error. See the class description for the precise route up the responder
// chain (plus document and controller objects) this message might travel.
// 
// It is not recommended that you attempt to override this method. If you wish
// to customize the error presentation, override [WillPresentError] instead.
//
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/presentError(_:modalFor:delegate:didPresent:contextInfo:)
func (r NSResponder) PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error_ foundation.INSError, window INSWindow, delegate objectivec.IObject, didPresentSelector objc.SEL, contextInfo uintptr) {
	objc.Send[objc.ID](r.ID, objc.Sel("presentError:modalForWindow:delegate:didPresentSelector:contextInfo:"), error_, window, delegate, didPresentSelector, contextInfo)
}
// Returns a custom version of the supplied error object that’s more
// suitable for presentation in alert sheets and dialogs.
//
// error: The error object to customize.
//
// # Return Value
// 
// The customized error object; if you decide not to customize the error
// presentation, return by sending this message to `super` (that is, `return
// [super error]`).
//
// # Discussion
// 
// When overriding this method, you can examine `error` and, if its localized
// description or recovery information is unhelpfully generic, return an error
// object with more specific localized text. If you do this, always use the
// domain and error code of the [NSError] object to distinguish between errors
// whose presentation you want to customize and those you don’t. Don’t
// make decisions based on the localized description, recovery suggestion, or
// recovery options because parsing localized text is problematic.
// 
// The default implementation of this method returns `error` unchanged.
//
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/willPresentError(_:)
func (r NSResponder) WillPresentError(error_ foundation.INSError) foundation.INSError {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("willPresentError:"), error_)
	return foundation.NSErrorFromID(rv)
}
// Attempts to perform the method indicated by an action with a specified
// argument.
//
// action: The selector identifying the action method.
//
// object: The object to use as the sole argument of the action method.
//
// # Return Value
// 
// Returns [false] if no responder is found that responds to `action;`
// otherwise, [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If the receiver responds to `action`, it invokes the method with `object`
// as the argument and returns [true]. If the receiver doesn’t respond, it
// sends this message to its next responder with the same selector and object.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/tryToPerform(_:with:)
func (r NSResponder) TryToPerformWith(action objc.SEL, object objectivec.IObject) bool {
	rv := objc.Send[bool](r.ID, objc.Sel("tryToPerform:with:"), action, object)
	return rv
}
// Overridden by subclasses to determine what services are available.
//
// sendType: A string identifying the send type of pasteboard data. May be an empty
// string (see discussion).
//
// returnType: A string identifying the return type of pasteboard data. May be an empty
// string (see discussion).
//
// # Return Value
// 
// If the receiver can place data of `sendType` on the pasteboard and receive
// data of `returnType`, it should return `self`; otherwise it should return
// either `[super ]` or `[[self nextResponder] ]`, which allows an object
// higher up in the responder chain to have an opportunity to handle the
// message.
//
// # Discussion
// 
// With each event, and for each service in the Services menu, the application
// object sends this message up the responder chain with the send and return
// type for the service being checked. This method is therefore invoked many
// times per event. The default implementation simply forwards this message to
// the next responder, ultimately returning `nil`.
// 
// Either `sendType` or `returnType`—but not both—may be empty. If
// `sendType` is empty, the service doesn’t require input from the
// application requesting the service. If `returnType` is empty, the service
// doesn’t return data.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/validRequestor(forSendType:returnType:)
func (r NSResponder) ValidRequestorForSendTypeReturnType(sendType NSPasteboardType, returnType NSPasteboardType) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("validRequestorForSendType:returnType:"), objc.String(string(sendType)), objc.String(string(returnType)))
	return objectivec.Object{ID: rv}
}
// Indicates whether a pen-down event should be treated as an ink event.
//
// event: An event object representing the event to be tested.
//
// # Return Value
// 
// Returns [true] if the specified event should be treated as an ink event,
// [false] if it should be treated as a mouse event.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method provides the ability to distinguish when a pen-down event
// should start inking versus when a pen-down event should be treated as a
// mouse-down event. This allows for a write-anywhere model for pen-based
// input.
// 
// The default implementation in [NSApplication] sends the method to the
// [NSWindow] object under the pen. If the window is inactive, this method
// returns [true], unless the pen-down is in the window drag region. If the
// window is active, this method is sent to the [NSView] object under the pen.
// 
// The default implementation in [NSView] returns [true], and [NSControl]
// overrides and returns [false]. This allows write-anywhere over most
// [NSView] objects, but allows the pen to be used to track in controls and to
// move windows.
// 
// A custom view should override this method to get the correct behavior for a
// pen-down in the view.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/shouldBeTreatedAsInkEvent(_:)
func (r NSResponder) ShouldBeTreatedAsInkEvent(event INSEvent) bool {
	rv := objc.Send[bool](r.ID, objc.Sel("shouldBeTreatedAsInkEvent:"), event)
	return rv
}
// Handles the case where an event or action message falls off the end of the
// responder chain.
//
// eventSelector: A selector identifying the action or event message.
//
// # Discussion
// 
// The default implementation beeps if `eventSelector` is [KeyDown].
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/noResponder(for:)
func (r NSResponder) NoResponderFor(eventSelector objc.SEL) {
	objc.Send[objc.ID](r.ID, objc.Sel("noResponderFor:"), eventSelector)
}
// Informs the receiver that the user has begun a touch gesture.
//
// event: An event object representing the gesture beginning.
//
// # Discussion
// 
// The event will be sent to the view under the touch in the key window.
// 
// Note that this method is no longer called on apps that link against macOS
// 10.11 and later. If you need to access the phases of a specific gesture,
// you can implement the responder for that gesture and examine its [Phase]
// property instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/beginGesture(with:)
func (r NSResponder) BeginGestureWithEvent(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("beginGestureWithEvent:"), event)
}
// Informs the receiver that the user has ended a touch gesture.
//
// event: An event object representing the gesture end.
//
// # Discussion
// 
// Note that this method is no longer called on apps that link against macOS
// 10.11 and later. If you need to access the phases of a specific gesture,
// you can implement the responder for that gesture and examine its [Phase]
// property instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/endGesture(with:)
func (r NSResponder) EndGestureWithEvent(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("endGestureWithEvent:"), event)
}
// Informs the receiver that the user has begun a pinch gesture.
//
// event: An event object representing the magnify gesture.
//
// # Discussion
// 
// The event will be sent to the view under the touch in the key window.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/magnify(with:)
func (r NSResponder) MagnifyWithEvent(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("magnifyWithEvent:"), event)
}
// Informs the receiver that the user has begun a rotation gesture.
//
// event: An event object representing the rotate gesture.
//
// # Discussion
// 
// The event will be sent to the view under the touch in the key window.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/rotate(with:)
func (r NSResponder) RotateWithEvent(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("rotateWithEvent:"), event)
}
// Informs the receiver that the user has begun a swipe gesture.
//
// event: An event object representing the swipe gesture.
//
// # Discussion
// 
// The event will be sent to the view under the touch in the key window.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/swipe(with:)
func (r NSResponder) SwipeWithEvent(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("swipeWithEvent:"), event)
}
// Informs the receiver that new set of touches has been recognized.
//
// event: An event object representing the beginning of a touch.
//
// # Discussion
// 
// The system sends the event to the view under the touch in the key window.
// To get the set of touches that began for this view (or descendants of this
// view) call [TouchesMatchingPhaseInView] on `event` and pass
// [TouchPhaseBegan] for the phase send.
// 
// This isn’t always the point of contact with the touch device. A touch
// that transitions from resting to active may be part of a
// [TouchesBeganWithEvent] set.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/touchesBegan(with:)
func (r NSResponder) TouchesBeganWithEvent(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("touchesBeganWithEvent:"), event)
}
// Informs the receiver that one or more touches has moved.
//
// event: An event object representing a touch movement.
//
// # Discussion
// 
// The system sends the to the view under the touch in the key window. To get
// the set of touches that moved for this view (or descendants of this view)
// call [TouchesMatchingPhaseInView] on `event` and pass [TouchPhaseMoved] for
// the phase.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/touchesMoved(with:)
func (r NSResponder) TouchesMovedWithEvent(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("touchesMovedWithEvent:"), event)
}
// Informs the receiver that tracking of touches has been cancelled for any
// reason.
//
// event: An event object representing the cancellation of a touch event.
//
// # Discussion
// 
// The event will be sent to the view under the touch in the key window.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/touchesCancelled(with:)
func (r NSResponder) TouchesCancelledWithEvent(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("touchesCancelledWithEvent:"), event)
}
// Returns that a set of touches have been removed.
//
// event: An event object representing the ending of a touch event.
//
// # Discussion
// 
// The system sends the event to the view under the touch in the key window.
// To get the set of touches that ended for this view (or descendants of this
// view) call [TouchesMatchingPhaseInView] on `event` and pass
// [TouchPhaseEnded] for the phase.
// 
// This isn’t always the point of removal with the touch device. A touch
// that transitions from active to resting may be part of an
// [TouchesEndedWithEvent] set.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/touchesEnded(with:)
func (r NSResponder) TouchesEndedWithEvent(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("touchesEndedWithEvent:"), event)
}
// Returns whether to forward elastic scrolling gesture events up the
// responder.
//
// axis: The gesture axis. See [NSEvent.GestureAxis] for the possible values.
// //
// [NSEvent.GestureAxis]: https://developer.apple.com/documentation/AppKit/NSEvent/GestureAxis
//
// # Return Value
// 
// Returns [true] when forward gesture scroll events should be forwarded up
// the responder chain when the scrolling content is already at the edge of
// the scrolled direction at the beginning of the scroll gesture; [false]
// otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Some views process gesture scroll events to perform elastic scrolling. In
// some cases, you may want to track gesture scroll events like a swipe, see
// [TrackSwipeEventWithOptionsDampenAmountThresholdMinMaxUsingHandler].
// 
// Implement this method and return [true] in your swipe controller and views
// that perform elastic scrolling will forward gesture scroll events up the
// responder chain on the following condition: the content to be scrolled is
// already at the edge of the scrolled direction at the beginning of the
// scroll gesture.
// 
// Otherwise, the view will perform elastic scrolling.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/wantsForwardedScrollEvents(for:)
func (r NSResponder) WantsForwardedScrollEventsForAxis(axis NSEventGestureAxis) bool {
	rv := objc.Send[bool](r.ID, objc.Sel("wantsForwardedScrollEventsForAxis:"), axis)
	return rv
}
// Informs the receiver that the user performed a smart zoom gesture.
//
// event: An event object representing the smart zoom gesture.
//
// # Discussion
// 
// The smart zoom gesture is a two-finger double tap on trackpads. In response
// to this event, you should intelligently magnify the content.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/smartMagnify(with:)
func (r NSResponder) SmartMagnifyWithEvent(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("smartMagnifyWithEvent:"), event)
}
// Implement this method to track gesture scroll events such as a swipe.
//
// axis: The event gesture axis of the swipe, which defines the scroll direction.
//
// # Return Value
// 
// [true] if gesture scroll events are to be forwarded up the responder chain;
// otherwise [false]. The default implementation returns [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Implement this method in your swipe controller and return [true] to inform
// views that perform elastic scrolling to forward gesture scroll events up
// the responder chain. The events are forwarded only on the following
// condition: the content to be scrolled is already at the edge of the
// scrolled direction when the scroll gesture begins. Otherwise, the view
// performs elastic scrolling. The default implementation returns [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/wantsScrollEventsForSwipeTracking(on:)
func (r NSResponder) WantsScrollEventsForSwipeTrackingOnAxis(axis NSEventGestureAxis) bool {
	rv := objc.Send[bool](r.ID, objc.Sel("wantsScrollEventsForSwipeTrackingOnAxis:"), axis)
	return rv
}
// Your custom subclass of the [NSResponder] class should override this method
// to create and configure your subclass’s default [NSTouchBar] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/makeTouchBar()
func (r NSResponder) MakeTouchBar() INSTouchBar {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("makeTouchBar"))
	return NSTouchBarFromID(rv)
}
// Performs all find oriented actions.
//
// sender: The sender of the find action.
//
// # Discussion
// 
// When an application performs a find action, it should send this message to
// the responder chain.
// 
// A responder of `` is responsible for creating and owning an instance of
// [NSTextFinder]. Before any other messages are sent to the [NSTextFinder],
// you should set its ‘client’ property to an object which implements the
// [NSTextFinderClient] protocol.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/performTextFinderAction(_:)
func (r NSResponder) PerformTextFinderAction(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("performTextFinderAction:"), sender)
}
// Creates a new window to show as a tab in a tabbed window.
//
// sender: The sender of the action.
//
// # Discussion
// 
// The system automatically calls this method to create a window for a new tab
// when the user clicks the plus button in a tabbed window. The system shows a
// plus button on a tabbed window only if this method exists on one of the
// following objects:
// 
// - In an [NSDocumentController] subclass - In the responder chain starting
// at [NSWindow], such as [NSWindow], the window delegate, the window
// controller, the [NSApplication] delegate, and so on
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/newWindowForTab(_:)
func (r NSResponder) NewWindowForTab(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("newWindowForTab:"), sender)
}
// Creates a new responder object with data in an unarchiver.
//
// coder: An unarchiver object.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/init(coder:)
func (r NSResponder) InitWithCoder(coder foundation.INSCoder) NSResponder {
	rv := objc.Send[NSResponder](r.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
//
// event: The key down event that matches the system-wide context menu hotkey
// combination.
//
// # Discussion
// 
// Handle a key event that should present a context menu at the user focus.
// 
// Most applications should not override this method. Instead, you should
// customize the context menu displayed from a keyboard event by implementing
// `` and `selectionAnchorRect`, or ``, rather than this method.
// 
// You should only override this method when you do not want the
// system-provided default behavior for the context menu hotkey, either for a
// specific key combination, or for the hotkey in general. For example, if
// your application already provides a different behavior for control-Return
// (the default context menu hotkey definition), and you want to preserve that
// behavior, you should override this method to handle that specific key
// combination, and then return without calling `super`. Note that the user
// may customize the hotkey to a different key combination, so in this
// example, if any other key combination is passed to your method, you would
// call `super`.
// 
// An implementation of this method should call `[super event]` to pass the
// request up the responder chain. If the message reaches the application
// object, NSApplication’s implementation of this method will send `` to the
// responder chain. If you do not call `super`, then no further handling of
// the key event will be performed.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/contextMenuKeyDown(_:)
func (r NSResponder) ContextMenuKeyDown(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("contextMenuKeyDown:"), event)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/mouseCancelled(with:)
func (r NSResponder) MouseCancelled(event INSEvent) {
	objc.Send[objc.ID](r.ID, objc.Sel("mouseCancelled:"), event)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/showWritingTools(_:)
func (r NSResponder) ShowWritingTools(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("showWritingTools:"), sender)
}
// Cancels the current operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/cancelOperation(_:)
func (r NSResponder) CancelOperation(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("cancelOperation:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/capitalizeWord(_:)
func (r NSResponder) CapitalizeWord(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("capitalizeWord:"), sender)
}
// Moves the visible content region so the current selection is visually
// centered.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/centerSelectionInVisibleArea(_:)
func (r NSResponder) CenterSelectionInVisibleArea(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("centerSelectionInVisibleArea:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/changeCaseOfLetter(_:)
func (r NSResponder) ChangeCaseOfLetter(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("changeCaseOfLetter:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/complete(_:)
func (r NSResponder) Complete(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("complete:"), sender)
}
// Deletes content moving backward from the current insertion point.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/deleteBackward(_:)
func (r NSResponder) DeleteBackward(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("deleteBackward:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/deleteBackwardByDecomposingPreviousCharacter(_:)
func (r NSResponder) DeleteBackwardByDecomposingPreviousCharacter(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("deleteBackwardByDecomposingPreviousCharacter:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/deleteForward(_:)
func (r NSResponder) DeleteForward(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("deleteForward:"), sender)
}
// Deletes content from the insertion point to the beginning of the current
// line.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/deleteToBeginningOfLine(_:)
func (r NSResponder) DeleteToBeginningOfLine(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("deleteToBeginningOfLine:"), sender)
}
// Deletes content from the insertion point to the beginning of the current
// paragraph.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/deleteToBeginningOfParagraph(_:)
func (r NSResponder) DeleteToBeginningOfParagraph(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("deleteToBeginningOfParagraph:"), sender)
}
// Deletes content from the insertion point to the end of the current line.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/deleteToEndOfLine(_:)
func (r NSResponder) DeleteToEndOfLine(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("deleteToEndOfLine:"), sender)
}
// Deletes content from the insertion point to the end of the current
// paragraph.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/deleteToEndOfParagraph(_:)
func (r NSResponder) DeleteToEndOfParagraph(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("deleteToEndOfParagraph:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/deleteToMark(_:)
func (r NSResponder) DeleteToMark(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("deleteToMark:"), sender)
}
// Deletes the word preceding the current insertion point.
//
// # Discussion
// 
// If the insertion point is in the middle of a word, this method deletes only
// the portion of the word preceding the insertion point.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/deleteWordBackward(_:)
func (r NSResponder) DeleteWordBackward(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("deleteWordBackward:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/deleteWordForward(_:)
func (r NSResponder) DeleteWordForward(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("deleteWordForward:"), sender)
}
// Performs the given selector if possible.
//
// selector: The selector to perform.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/doCommand(by:)
func (r NSResponder) DoCommandBySelector(selector objc.SEL) {
	objc.Send[objc.ID](r.ID, objc.Sel("doCommandBySelector:"), selector)
}
// Indents the content at the current selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/indent(_:)
func (r NSResponder) Indent(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("indent:"), sender)
}
// Inserts a backtab character.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/insertBacktab(_:)
func (r NSResponder) InsertBacktab(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("insertBacktab:"), sender)
}
// Inserts a container break, such as a new page break.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/insertContainerBreak(_:)
func (r NSResponder) InsertContainerBreak(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("insertContainerBreak:"), sender)
}
// Inserts a double quotation mark without substituting a curly quotation
// mark.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/insertDoubleQuoteIgnoringSubstitution(_:)
func (r NSResponder) InsertDoubleQuoteIgnoringSubstitution(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("insertDoubleQuoteIgnoringSubstitution:"), sender)
}
// Inserts a line break character.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/insertLineBreak(_:)
func (r NSResponder) InsertLineBreak(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("insertLineBreak:"), sender)
}
// Inserts a newline character.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/insertNewline(_:)
func (r NSResponder) InsertNewline(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("insertNewline:"), sender)
}
// Inserts a newline character without invoking the field editor’s normal
// handling to end editing.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/insertNewlineIgnoringFieldEditor(_:)
func (r NSResponder) InsertNewlineIgnoringFieldEditor(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("insertNewlineIgnoringFieldEditor:"), sender)
}
// Inserts a paragraph separator.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/insertParagraphSeparator(_:)
func (r NSResponder) InsertParagraphSeparator(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("insertParagraphSeparator:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/insertSingleQuoteIgnoringSubstitution(_:)
func (r NSResponder) InsertSingleQuoteIgnoringSubstitution(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("insertSingleQuoteIgnoringSubstitution:"), sender)
}
// Inserts a tab character.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/insertTab(_:)
func (r NSResponder) InsertTab(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("insertTab:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/insertTabIgnoringFieldEditor(_:)
func (r NSResponder) InsertTabIgnoringFieldEditor(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("insertTabIgnoringFieldEditor:"), sender)
}
// Inserts the text you specify.
//
// insertString: The string to insert in the responder.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/insertText(_:)
func (r NSResponder) InsertText(insertString objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("insertText:"), insertString)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/lowercaseWord(_:)
func (r NSResponder) LowercaseWord(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("lowercaseWord:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/makeBaseWritingDirectionLeftToRight(_:)
func (r NSResponder) MakeBaseWritingDirectionLeftToRight(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("makeBaseWritingDirectionLeftToRight:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/makeBaseWritingDirectionNatural(_:)
func (r NSResponder) MakeBaseWritingDirectionNatural(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("makeBaseWritingDirectionNatural:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/makeBaseWritingDirectionRightToLeft(_:)
func (r NSResponder) MakeBaseWritingDirectionRightToLeft(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("makeBaseWritingDirectionRightToLeft:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/makeTextWritingDirectionLeftToRight(_:)
func (r NSResponder) MakeTextWritingDirectionLeftToRight(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("makeTextWritingDirectionLeftToRight:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/makeTextWritingDirectionNatural(_:)
func (r NSResponder) MakeTextWritingDirectionNatural(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("makeTextWritingDirectionNatural:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/makeTextWritingDirectionRightToLeft(_:)
func (r NSResponder) MakeTextWritingDirectionRightToLeft(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("makeTextWritingDirectionRightToLeft:"), sender)
}
// Moves the insertion pointer backward in the current content.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveBackward(_:)
func (r NSResponder) MoveBackward(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveBackward:"), sender)
}
// Extends the selection to include the content before the current selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveBackwardAndModifySelection(_:)
func (r NSResponder) MoveBackwardAndModifySelection(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveBackwardAndModifySelection:"), sender)
}
// Moves the insertion pointer down in the current content.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveDown(_:)
func (r NSResponder) MoveDown(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveDown:"), sender)
}
// Extends the selection to include the content below the current selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveDownAndModifySelection(_:)
func (r NSResponder) MoveDownAndModifySelection(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveDownAndModifySelection:"), sender)
}
// Moves the insertion pointer forward in the current content.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveForward(_:)
func (r NSResponder) MoveForward(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveForward:"), sender)
}
// Extends the selection to include the content after the current selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveForwardAndModifySelection(_:)
func (r NSResponder) MoveForwardAndModifySelection(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveForwardAndModifySelection:"), sender)
}
// Moves the insertion pointer left in the current content.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveLeft(_:)
func (r NSResponder) MoveLeft(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveLeft:"), sender)
}
// Extends the selection to include the content to the left of the current
// selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveLeftAndModifySelection(_:)
func (r NSResponder) MoveLeftAndModifySelection(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveLeftAndModifySelection:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveParagraphBackwardAndModifySelection(_:)
func (r NSResponder) MoveParagraphBackwardAndModifySelection(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveParagraphBackwardAndModifySelection:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveParagraphForwardAndModifySelection(_:)
func (r NSResponder) MoveParagraphForwardAndModifySelection(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveParagraphForwardAndModifySelection:"), sender)
}
// Moves the insertion pointer right in the current content.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveRight(_:)
func (r NSResponder) MoveRight(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveRight:"), sender)
}
// Extends the selection to include the content to the right of the current
// selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveRightAndModifySelection(_:)
func (r NSResponder) MoveRightAndModifySelection(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveRightAndModifySelection:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToBeginningOfDocument(_:)
func (r NSResponder) MoveToBeginningOfDocument(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveToBeginningOfDocument:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToBeginningOfDocumentAndModifySelection(_:)
func (r NSResponder) MoveToBeginningOfDocumentAndModifySelection(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveToBeginningOfDocumentAndModifySelection:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToBeginningOfLine(_:)
func (r NSResponder) MoveToBeginningOfLine(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveToBeginningOfLine:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToBeginningOfLineAndModifySelection(_:)
func (r NSResponder) MoveToBeginningOfLineAndModifySelection(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveToBeginningOfLineAndModifySelection:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToBeginningOfParagraph(_:)
func (r NSResponder) MoveToBeginningOfParagraph(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveToBeginningOfParagraph:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToBeginningOfParagraphAndModifySelection(_:)
func (r NSResponder) MoveToBeginningOfParagraphAndModifySelection(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveToBeginningOfParagraphAndModifySelection:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToEndOfDocument(_:)
func (r NSResponder) MoveToEndOfDocument(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveToEndOfDocument:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToEndOfDocumentAndModifySelection(_:)
func (r NSResponder) MoveToEndOfDocumentAndModifySelection(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveToEndOfDocumentAndModifySelection:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToEndOfLine(_:)
func (r NSResponder) MoveToEndOfLine(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveToEndOfLine:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToEndOfLineAndModifySelection(_:)
func (r NSResponder) MoveToEndOfLineAndModifySelection(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveToEndOfLineAndModifySelection:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToEndOfParagraph(_:)
func (r NSResponder) MoveToEndOfParagraph(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveToEndOfParagraph:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToEndOfParagraphAndModifySelection(_:)
func (r NSResponder) MoveToEndOfParagraphAndModifySelection(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveToEndOfParagraphAndModifySelection:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToLeftEndOfLine(_:)
func (r NSResponder) MoveToLeftEndOfLine(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveToLeftEndOfLine:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToLeftEndOfLineAndModifySelection(_:)
func (r NSResponder) MoveToLeftEndOfLineAndModifySelection(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveToLeftEndOfLineAndModifySelection:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToRightEndOfLine(_:)
func (r NSResponder) MoveToRightEndOfLine(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveToRightEndOfLine:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToRightEndOfLineAndModifySelection(_:)
func (r NSResponder) MoveToRightEndOfLineAndModifySelection(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveToRightEndOfLineAndModifySelection:"), sender)
}
// Moves the insertion pointer up in the current content.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveUp(_:)
func (r NSResponder) MoveUp(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveUp:"), sender)
}
// Extends the selection to include the content above the current selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveUpAndModifySelection(_:)
func (r NSResponder) MoveUpAndModifySelection(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveUpAndModifySelection:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveWordBackward(_:)
func (r NSResponder) MoveWordBackward(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveWordBackward:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveWordBackwardAndModifySelection(_:)
func (r NSResponder) MoveWordBackwardAndModifySelection(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveWordBackwardAndModifySelection:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveWordForward(_:)
func (r NSResponder) MoveWordForward(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveWordForward:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveWordForwardAndModifySelection(_:)
func (r NSResponder) MoveWordForwardAndModifySelection(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveWordForwardAndModifySelection:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveWordLeft(_:)
func (r NSResponder) MoveWordLeft(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveWordLeft:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveWordLeftAndModifySelection(_:)
func (r NSResponder) MoveWordLeftAndModifySelection(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveWordLeftAndModifySelection:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveWordRight(_:)
func (r NSResponder) MoveWordRight(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveWordRight:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveWordRightAndModifySelection(_:)
func (r NSResponder) MoveWordRightAndModifySelection(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveWordRightAndModifySelection:"), sender)
}
// Moves the visible content region down by a page.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/pageDown(_:)
func (r NSResponder) PageDown(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("pageDown:"), sender)
}
// Moves the visible content region down by a page, and extends the current
// selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/pageDownAndModifySelection(_:)
func (r NSResponder) PageDownAndModifySelection(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("pageDownAndModifySelection:"), sender)
}
// Moves the visible content region up by a page.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/pageUp(_:)
func (r NSResponder) PageUp(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("pageUp:"), sender)
}
// Moves the visible content region up by a page, and extends the current
// selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/pageUpAndModifySelection(_:)
func (r NSResponder) PageUpAndModifySelection(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("pageUpAndModifySelection:"), sender)
}
// Invokes QuickLook to preview the current selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/quickLookPreviewItems(_:)
func (r NSResponder) QuickLookPreviewItems(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("quickLookPreviewItems:"), sender)
}
// Restores the state necessary to continue the specified user activity.
//
// userActivity: The user activity to continue.
//
// # Discussion
// 
// Implement this method to restore an object’s state using the specified
// user activity. The system calls this method on any responders or documents
// passed to the `restorationHandler` in
// [ApplicationContinueUserActivityRestorationHandler]. The system calls this
// method on the main thread. Your implementation should use the state data
// contained in the specified user activity’s [userInfo] dictionary to
// restore the object.
// 
// On macOS, the system can automatically restore activities managed by
// [NSDocument] if you don’t implement
// [ApplicationContinueUserActivityRestorationHandler], or if you return
// [false]. When this occurs, the system opens the document using
// [OpenDocumentWithContentsOfURLDisplayCompletionHandler], and calls
// `restoreUserActivityState` on it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [userInfo]: https://developer.apple.com/documentation/Foundation/NSUserActivity/userInfo
//
// See: https://developer.apple.com/documentation/AppKit/NSUserActivityRestoring/restoreUserActivityState(_:)
func (r NSResponder) RestoreUserActivityState(userActivity foundation.NSUserActivity) {
	objc.Send[objc.ID](r.ID, objc.Sel("restoreUserActivityState:"), userActivity)
}
// Scrolls the content down by a line.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/scrollLineDown(_:)
func (r NSResponder) ScrollLineDown(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("scrollLineDown:"), sender)
}
// Scrolls the content up by a line.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/scrollLineUp(_:)
func (r NSResponder) ScrollLineUp(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("scrollLineUp:"), sender)
}
// Scrolls the content down by a page.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/scrollPageDown(_:)
func (r NSResponder) ScrollPageDown(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("scrollPageDown:"), sender)
}
// Scrolls the content up by a page.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/scrollPageUp(_:)
func (r NSResponder) ScrollPageUp(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("scrollPageUp:"), sender)
}
// Scrolls the content to the beginning of the document.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/scrollToBeginningOfDocument(_:)
func (r NSResponder) ScrollToBeginningOfDocument(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("scrollToBeginningOfDocument:"), sender)
}
// Scrolls the content to the end of the document.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/scrollToEndOfDocument(_:)
func (r NSResponder) ScrollToEndOfDocument(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("scrollToEndOfDocument:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/selectAll(_:)
func (r NSResponder) SelectAll(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("selectAll:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/selectLine(_:)
func (r NSResponder) SelectLine(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("selectLine:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/selectParagraph(_:)
func (r NSResponder) SelectParagraph(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("selectParagraph:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/selectToMark(_:)
func (r NSResponder) SelectToMark(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("selectToMark:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/selectWord(_:)
func (r NSResponder) SelectWord(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("selectWord:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/setMark(_:)
func (r NSResponder) SetMark(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("setMark:"), sender)
}
// Implemented by subclasses to invoke the help system, displaying information
// relevant to the receiver and its current state.
//
// sender: Typically the object that invoked this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/showContextHelp(_:)
func (r NSResponder) ShowContextHelp(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("showContextHelp:"), sender)
}
//
// sender: The object that originated the display of the context menu.
//
// # Discussion
// 
// Present a context menu at the text cursor position, selection, or wherever
// is appropriate for your responder.
// 
// NSView has a default implementation of this method. For any view that
// returns a non-nil value from `-`, the default implementation will display
// that menu automatically. The NSView implementation uses the
// `selectionAnchorRect` method in the [NSViewContentSelectionInfo] protocol
// to determine the location of the selection and of the menu. The NSView
// implementation determines the menu to display by calling `` with a
// right-mouse-down event that is centered on the anchor rect. If
// `selectionAnchorRect` is not implemented, then the NSView implementation
// calls `menuForEvent` with a right-mouse-down event that is centered on the
// view’s bounds, and also displays the menu at that location.
// 
// You should only override this method in a custom view if you need full
// control over the display of a context menu from the keyboard or
// Accessibility, beyond what is provided by default by NSView.
// 
// If the view does not support a context menu, then you should call `[[self
// nextResponder] _cmd sender]` to pass the request up the responder chain.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/showContextMenuForSelection(_:)
func (r NSResponder) ShowContextMenuForSelection(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("showContextMenuForSelection:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/swapWithMark(_:)
func (r NSResponder) SwapWithMark(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("swapWithMark:"), sender)
}
// Transposes the content around the current selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/transpose(_:)
func (r NSResponder) Transpose(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("transpose:"), sender)
}
// Transposes the words around the current selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/transposeWords(_:)
func (r NSResponder) TransposeWords(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("transposeWords:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/uppercaseWord(_:)
func (r NSResponder) UppercaseWord(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("uppercaseWord:"), sender)
}
// Deletes the current selection, placing it in a temporary buffer, such as
// the Clipboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/yank(_:)
func (r NSResponder) Yank(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("yank:"), sender)
}
func (r NSResponder) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns the classes that support secure coding.
//
// keyPath: The key path of the restorable object.
//
// # Return Value
// 
// An array of classes that support secure coding.
//
// # Discussion
// 
// The system calls the function during a secure state restoration and
// restores values only for the allowed classes your app returns in the array.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/allowedClasses(forRestorableStateKeyPath:)
func (_NSResponderClass NSResponderClass) AllowedClassesForRestorableStateKeyPath(keyPath string) []objc.Class {
	rv := objc.Send[[]objc.ID](objc.ID(_NSResponderClass.class), objc.Sel("allowedClassesForRestorableStateKeyPath:"), objc.String(keyPath))
	return objc.ConvertSlice(rv, func(id objc.ID) objc.Class {
		return objc.Class(id)
	})
}

// A Boolean value that indicates whether the responder accepts first
// responder status.
//
// # Discussion
// 
// As first responder, the receiver is the first object in the responder chain
// to be sent key events and action messages. By default, this property is
// [false]. Subclasses set this property to [true] if the receiver accepts
// first responder status.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/acceptsFirstResponder
func (r NSResponder) AcceptsFirstResponder() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("acceptsFirstResponder"))
	return rv
}
// The next responder after this one, or `nil` if it has none.
//
// # Discussion
// 
// The next responder must be an object that inherits, directly or indirectly,
// from [NSResponder].
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/nextResponder
func (r NSResponder) NextResponder() INSResponder {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("nextResponder"))
	return NSResponderFromID(objc.ID(rv))
}
func (r NSResponder) SetNextResponder(value INSResponder) {
	objc.Send[struct{}](r.ID, objc.Sel("setNextResponder:"), value)
}
// An object encapsulating a user activity supported by this responder.
//
// # Discussion
// 
// By setting the [UserActivity] property on a responder, the [NSUserActivity]
// object becomes managed by AppKit. You should override
// [UpdateUserActivityState] to write lazily any state data representing the
// user’s activity to the `userInfo` dictionary. AppKit automatically saves
// user activities it manages at appropriate times. Multiple responders can
// share a single [NSUserActivity] instance, in which case they all get a
// callback, such as [UpdateUserActivityState], when the system updates the
// user activity object.
// 
// In macOS, [NSUserActivity] objects managed by [NSResponder] automatically
// [becomeCurrent()] based on the main window and the responder chain.
// 
// A responder object can set its [UserActivity] property to `nil` if it no
// longer wants to participate. Any [NSUserActivity] objects that AppKit
// manages but have no associated responders (or documents) are automatically
// invalidated.
// 
// You can use this property from any thread, and it’s key-value observable
// (KVO).
//
// [NSUserActivity]: https://developer.apple.com/documentation/Foundation/NSUserActivity
// [becomeCurrent()]: https://developer.apple.com/documentation/Foundation/NSUserActivity/becomeCurrent()
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/userActivity
func (r NSResponder) UserActivity() foundation.NSUserActivity {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("userActivity"))
	return foundation.NSUserActivityFromID(objc.ID(rv))
}
func (r NSResponder) SetUserActivity(value foundation.NSUserActivity) {
	objc.Send[struct{}](r.ID, objc.Sel("setUserActivity:"), value)
}
// Returns the responder’s menu.
//
// # Discussion
// 
// For [NSApplication] this menu is the same as the menu returned by its
// [mainMenu] property.
//
// [mainMenu]: https://developer.apple.com/documentation/AppKit/NSApplication/mainMenu
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/menu
func (r NSResponder) Menu() INSMenu {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("menu"))
	return NSMenuFromID(objc.ID(rv))
}
func (r NSResponder) SetMenu(value INSMenu) {
	objc.Send[struct{}](r.ID, objc.Sel("setMenu:"), value)
}
// The undo manager for this responder.
//
// # Discussion
// 
// The [NSResponder] implementation simply invokes this property on the next
// responder.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/undoManager
func (r NSResponder) UndoManager() foundation.NSUndoManager {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("undoManager"))
	return foundation.NSUndoManagerFromID(objc.ID(rv))
}
// The [NSTouchBar] object associated with the responder.
//
// # Discussion
// 
// If you have not explicitly provided an [NSTouchBar] object for a responder
// by setting this property, the system sends the [TouchBar] message to the
// responder to create the default bar. This property is archived.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/touchBar
func (r NSResponder) TouchBar() INSTouchBar {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("touchBar"))
	return NSTouchBarFromID(objc.ID(rv))
}
func (r NSResponder) SetTouchBar(value INSTouchBar) {
	objc.Send[struct{}](r.ID, objc.Sel("setTouchBar:"), value)
}

// Returns an array of key paths representing the restorable attributes of the
// responder.
//
// # Return Value
// 
// An array of [NSString] objects, each of which contains a key path to one of
// the responder’s attributes.
// 
// # Discussion
// 
// You can use this method instead of, or in addition to, the
// [EncodeRestorableStateWithCoder] and [RestoreStateWithCoder] methods to
// save and restore the state of your responder. The key paths you return must
// refer to attributes that are key-value coding and key-value observing
// compliant. To learn more about these mechanisms, see [Key-Value Coding
// Programming Guide] and [Key-Value Observing Programming Guide].
// 
// When changes are detected, the specified attributes are automatically
// written to disk with the rest of the application’s interface-related
// state. At launch time, the attributes are automatically restored to their
// previous values.
//
// [Key-Value Coding Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueCoding/index.html#//apple_ref/doc/uid/10000107i
// [Key-Value Observing Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueObserving/KeyValueObserving.html#//apple_ref/doc/uid/10000177i
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/restorableStateKeyPaths
func (_NSResponderClass NSResponderClass) RestorableStateKeyPaths() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSResponderClass.class), objc.Sel("restorableStateKeyPaths"))
	return objc.ConvertSliceToStrings(rv)
}

			// Protocol methods for NSStandardKeyBindingResponding
			

			// Protocol methods for NSTouchBarProvider
			

			// Protocol methods for NSUserActivityRestoring
			

