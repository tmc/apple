// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSAppleEventManager] class.
var (
	_NSAppleEventManagerClass     NSAppleEventManagerClass
	_NSAppleEventManagerClassOnce sync.Once
)

func getNSAppleEventManagerClass() NSAppleEventManagerClass {
	_NSAppleEventManagerClassOnce.Do(func() {
		_NSAppleEventManagerClass = NSAppleEventManagerClass{class: objc.GetClass("NSAppleEventManager")}
	})
	return _NSAppleEventManagerClass
}

// GetNSAppleEventManagerClass returns the class object for NSAppleEventManager.
func GetNSAppleEventManagerClass() NSAppleEventManagerClass {
	return getNSAppleEventManagerClass()
}

type NSAppleEventManagerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAppleEventManagerClass) Alloc() NSAppleEventManager {
	rv := objc.Send[NSAppleEventManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A mechanism for registering handler routines for specific types of Apple
// events and dispatching events to those handlers.
//
// # Overview
// 
// Cocoa provides built-in scriptability support that uses scriptability
// information supplied by an application to automatically convert Apple
// events into script command objects that perform the desired operation.
// However, some applications may want to perform more basic Apple event
// handling, in which an application registers handlers for the Apple events
// it can process, then calls on the Apple Event Manager to dispatch received
// Apple events to the appropriate handler. [NSAppleEventManager] supports
// these mechanisms by providing methods to register and remove handlers and
// to dispatch Apple events to the appropriate handler, if one exists. For
// related information, see [How Cocoa Applications Handle Apple Events]
// 
// Each application has at most one instance of [NSAppleEventManager]. To
// obtain a reference to it, you call the class method
// [NSAppleEventManager.SharedAppleEventManager], which creates the instance if it doesn’t
// already exist.
// 
// For information about the Apple Event Manager, see [Apple Event Manager]
// and Apple Events Programming Guide.
//
// [Apple Event Manager]: https://developer.apple.com/documentation/applicationservices/apple_event_manager
// [How Cocoa Applications Handle Apple Events]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ScriptableCocoaApplications/SApps_handle_AEs/SAppsHandleAEs.html#//apple_ref/doc/uid/20001239
//
// # Working with event handlers
//
//   - [NSAppleEventManager.RemoveEventHandlerForEventClassAndEventID]: If an Apple event handler has been registered for the event specified by `eventClass` and `eventID`, removes it.
//   - [NSAppleEventManager.SetEventHandlerAndSelectorForEventClassAndEventID]: Registers the Apple event handler specified by `handler` for the event specified by `eventClass` and `eventID`.
//
// # Working with events
//
//   - [NSAppleEventManager.DispatchRawAppleEventWithRawReplyHandlerRefCon]: Causes the Apple event specified by `theAppleEvent` to be dispatched to the appropriate Apple event handler, if one has been registered by calling [setEventHandler(_:andSelector:forEventClass:andEventID:)](<doc://com.apple.foundation/documentation/Foundation/NSAppleEventManager/setEventHandler(_:andSelector:forEventClass:andEventID:)>).
//
// # Suspending and resuming Apple events
//
//   - [NSAppleEventManager.AppleEventForSuspensionID]: Given a nonzero `suspensionID` returned by an invocation of [suspendCurrentAppleEvent()](<doc://com.apple.foundation/documentation/Foundation/NSAppleEventManager/suspendCurrentAppleEvent()>), returns the descriptor for the event whose handling was suspended.
//   - [NSAppleEventManager.CurrentAppleEvent]: Returns the descriptor for `currentAppleEvent` if an Apple event is being handled on the current thread.
//   - [NSAppleEventManager.CurrentReplyAppleEvent]: Returns the corresponding reply event descriptor if an Apple event is being handled on the current thread.
//   - [NSAppleEventManager.ReplyAppleEventForSuspensionID]: Given a nonzero `suspensionID` returned by an invocation of [suspendCurrentAppleEvent()](<doc://com.apple.foundation/documentation/Foundation/NSAppleEventManager/suspendCurrentAppleEvent()>), returns the corresponding reply event descriptor.
//   - [NSAppleEventManager.ResumeWithSuspensionID]: Given a nonzero `suspensionID` returned by an invocation of [suspendCurrentAppleEvent()](<doc://com.apple.foundation/documentation/Foundation/NSAppleEventManager/suspendCurrentAppleEvent()>), signal that handling of the suspended event may now continue.
//   - [NSAppleEventManager.SetCurrentAppleEventAndReplyEventWithSuspensionID]: Given a nonzero `suspensionID` returned by an invocation of [suspendCurrentAppleEvent()](<doc://com.apple.foundation/documentation/Foundation/NSAppleEventManager/suspendCurrentAppleEvent()>), sets the values that will be returned by subsequent invocations of [currentAppleEvent](<doc://com.apple.foundation/documentation/Foundation/NSAppleEventManager/currentAppleEvent>) and [currentReplyAppleEvent](<doc://com.apple.foundation/documentation/Foundation/NSAppleEventManager/currentReplyAppleEvent>) to be the event whose handling was suspended and its corresponding reply event, respectively.
//   - [NSAppleEventManager.SuspendCurrentAppleEvent]: Suspends the handling of the current event and returns an ID that must be used to resume the handling of the event if an Apple event is being handled on the current thread.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventManager
type NSAppleEventManager struct {
	objectivec.Object
}

// NSAppleEventManagerFromID constructs a [NSAppleEventManager] from an objc.ID.
//
// A mechanism for registering handler routines for specific types of Apple
// events and dispatching events to those handlers.
func NSAppleEventManagerFromID(id objc.ID) NSAppleEventManager {
	return NSAppleEventManager{objectivec.Object{ID: id}}
}
// NOTE: NSAppleEventManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSAppleEventManager] class.
//
// # Working with event handlers
//
//   - [INSAppleEventManager.RemoveEventHandlerForEventClassAndEventID]: If an Apple event handler has been registered for the event specified by `eventClass` and `eventID`, removes it.
//   - [INSAppleEventManager.SetEventHandlerAndSelectorForEventClassAndEventID]: Registers the Apple event handler specified by `handler` for the event specified by `eventClass` and `eventID`.
//
// # Working with events
//
//   - [INSAppleEventManager.DispatchRawAppleEventWithRawReplyHandlerRefCon]: Causes the Apple event specified by `theAppleEvent` to be dispatched to the appropriate Apple event handler, if one has been registered by calling [setEventHandler(_:andSelector:forEventClass:andEventID:)](<doc://com.apple.foundation/documentation/Foundation/NSAppleEventManager/setEventHandler(_:andSelector:forEventClass:andEventID:)>).
//
// # Suspending and resuming Apple events
//
//   - [INSAppleEventManager.AppleEventForSuspensionID]: Given a nonzero `suspensionID` returned by an invocation of [suspendCurrentAppleEvent()](<doc://com.apple.foundation/documentation/Foundation/NSAppleEventManager/suspendCurrentAppleEvent()>), returns the descriptor for the event whose handling was suspended.
//   - [INSAppleEventManager.CurrentAppleEvent]: Returns the descriptor for `currentAppleEvent` if an Apple event is being handled on the current thread.
//   - [INSAppleEventManager.CurrentReplyAppleEvent]: Returns the corresponding reply event descriptor if an Apple event is being handled on the current thread.
//   - [INSAppleEventManager.ReplyAppleEventForSuspensionID]: Given a nonzero `suspensionID` returned by an invocation of [suspendCurrentAppleEvent()](<doc://com.apple.foundation/documentation/Foundation/NSAppleEventManager/suspendCurrentAppleEvent()>), returns the corresponding reply event descriptor.
//   - [INSAppleEventManager.ResumeWithSuspensionID]: Given a nonzero `suspensionID` returned by an invocation of [suspendCurrentAppleEvent()](<doc://com.apple.foundation/documentation/Foundation/NSAppleEventManager/suspendCurrentAppleEvent()>), signal that handling of the suspended event may now continue.
//   - [INSAppleEventManager.SetCurrentAppleEventAndReplyEventWithSuspensionID]: Given a nonzero `suspensionID` returned by an invocation of [suspendCurrentAppleEvent()](<doc://com.apple.foundation/documentation/Foundation/NSAppleEventManager/suspendCurrentAppleEvent()>), sets the values that will be returned by subsequent invocations of [currentAppleEvent](<doc://com.apple.foundation/documentation/Foundation/NSAppleEventManager/currentAppleEvent>) and [currentReplyAppleEvent](<doc://com.apple.foundation/documentation/Foundation/NSAppleEventManager/currentReplyAppleEvent>) to be the event whose handling was suspended and its corresponding reply event, respectively.
//   - [INSAppleEventManager.SuspendCurrentAppleEvent]: Suspends the handling of the current event and returns an ID that must be used to resume the handling of the event if an Apple event is being handled on the current thread.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventManager
type INSAppleEventManager interface {
	objectivec.IObject

	// Topic: Working with event handlers

	// If an Apple event handler has been registered for the event specified by `eventClass` and `eventID`, removes it.
	RemoveEventHandlerForEventClassAndEventID(eventClass uint32, eventID uint32)
	// Registers the Apple event handler specified by `handler` for the event specified by `eventClass` and `eventID`.
	SetEventHandlerAndSelectorForEventClassAndEventID(handler ErrorHandler, handleEventSelector objc.SEL, eventClass uint32, eventID uint32)

	// Topic: Working with events

	// Causes the Apple event specified by `theAppleEvent` to be dispatched to the appropriate Apple event handler, if one has been registered by calling [setEventHandler(_:andSelector:forEventClass:andEventID:)](<doc://com.apple.foundation/documentation/Foundation/NSAppleEventManager/setEventHandler(_:andSelector:forEventClass:andEventID:)>).
	DispatchRawAppleEventWithRawReplyHandlerRefCon(theAppleEvent objectivec.IObject, theReply objectivec.IObject, handlerRefCon ErrorHandler) objectivec.IObject

	// Topic: Suspending and resuming Apple events

	// Given a nonzero `suspensionID` returned by an invocation of [suspendCurrentAppleEvent()](<doc://com.apple.foundation/documentation/Foundation/NSAppleEventManager/suspendCurrentAppleEvent()>), returns the descriptor for the event whose handling was suspended.
	AppleEventForSuspensionID(suspensionID NSAppleEventManagerSuspensionID) INSAppleEventDescriptor
	// Returns the descriptor for `currentAppleEvent` if an Apple event is being handled on the current thread.
	CurrentAppleEvent() INSAppleEventDescriptor
	// Returns the corresponding reply event descriptor if an Apple event is being handled on the current thread.
	CurrentReplyAppleEvent() INSAppleEventDescriptor
	// Given a nonzero `suspensionID` returned by an invocation of [suspendCurrentAppleEvent()](<doc://com.apple.foundation/documentation/Foundation/NSAppleEventManager/suspendCurrentAppleEvent()>), returns the corresponding reply event descriptor.
	ReplyAppleEventForSuspensionID(suspensionID NSAppleEventManagerSuspensionID) INSAppleEventDescriptor
	// Given a nonzero `suspensionID` returned by an invocation of [suspendCurrentAppleEvent()](<doc://com.apple.foundation/documentation/Foundation/NSAppleEventManager/suspendCurrentAppleEvent()>), signal that handling of the suspended event may now continue.
	ResumeWithSuspensionID(suspensionID NSAppleEventManagerSuspensionID)
	// Given a nonzero `suspensionID` returned by an invocation of [suspendCurrentAppleEvent()](<doc://com.apple.foundation/documentation/Foundation/NSAppleEventManager/suspendCurrentAppleEvent()>), sets the values that will be returned by subsequent invocations of [currentAppleEvent](<doc://com.apple.foundation/documentation/Foundation/NSAppleEventManager/currentAppleEvent>) and [currentReplyAppleEvent](<doc://com.apple.foundation/documentation/Foundation/NSAppleEventManager/currentReplyAppleEvent>) to be the event whose handling was suspended and its corresponding reply event, respectively.
	SetCurrentAppleEventAndReplyEventWithSuspensionID(suspensionID NSAppleEventManagerSuspensionID)
	// Suspends the handling of the current event and returns an ID that must be used to resume the handling of the event if an Apple event is being handled on the current thread.
	SuspendCurrentAppleEvent() NSAppleEventManagerSuspensionID
}

// Init initializes the instance.
func (a NSAppleEventManager) Init() NSAppleEventManager {
	rv := objc.Send[NSAppleEventManager](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAppleEventManager) Autorelease() NSAppleEventManager {
	rv := objc.Send[NSAppleEventManager](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAppleEventManager creates a new NSAppleEventManager instance.
func NewNSAppleEventManager() NSAppleEventManager {
	class := getNSAppleEventManagerClass()
	rv := objc.Send[NSAppleEventManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// If an Apple event handler has been registered for the event specified by
// `eventClass` and `eventID`, removes it.
//
// # Discussion
// 
// Otherwise does nothing.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/removeEventHandler(forEventClass:andEventID:)
func (a NSAppleEventManager) RemoveEventHandlerForEventClassAndEventID(eventClass uint32, eventID uint32) {
	objc.Send[objc.ID](a.ID, objc.Sel("removeEventHandlerForEventClass:andEventID:"), eventClass, eventID)
}
// Registers the Apple event handler specified by `handler` for the event
// specified by `eventClass` and `eventID`.
//
// # Discussion
// 
// If an event handler is already registered for the specified event class and
// event ID, removes it. The signature for `handler` should match the
// following:
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/setEventHandler(_:andSelector:forEventClass:andEventID:)
func (a NSAppleEventManager) SetEventHandlerAndSelectorForEventClassAndEventID(handler ErrorHandler, handleEventSelector objc.SEL, eventClass uint32, eventID uint32) {
_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](a.ID, objc.Sel("setEventHandler:andSelector:forEventClass:andEventID:"), _block0, handleEventSelector, eventClass, eventID)
}
// Causes the Apple event specified by `theAppleEvent` to be dispatched to the
// appropriate Apple event handler, if one has been registered by calling
// [SetEventHandlerAndSelectorForEventClassAndEventID].
//
// # Discussion
// 
// The `theReply` parameter always specifies a reply Apple event, never `nil`.
// However, the handler should not fill out the reply if the descriptor type
// for the reply event is `typeNull`, indicating the sender does not want a
// reply.
// 
// The `handlerRefcon` parameter provides 4 bytes of data to the handler; a
// common use for this parameter is to pass a pointer to additional data.
// 
// This method is primarily intended for Cocoa’s internal use. Note that an
// event means routing an event to an appropriate handler in the current
// application. You cannot use this method to an event to other applications.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/dispatchRawAppleEvent(_:withRawReply:handlerRefCon:)
func (a NSAppleEventManager) DispatchRawAppleEventWithRawReplyHandlerRefCon(theAppleEvent objectivec.IObject, theReply objectivec.IObject, handlerRefCon ErrorHandler) objectivec.IObject {
_block2, _cleanup2 := NewErrorBlock(handlerRefCon)
	defer _cleanup2()
	rv := objc.Send[objc.ID](a.ID, objc.Sel("dispatchRawAppleEvent:withRawReply:handlerRefCon:"), theAppleEvent, theReply, _block2)
	return objectivec.Object{ID: rv}
}
// Given a nonzero `suspensionID` returned by an invocation of
// [SuspendCurrentAppleEvent], returns the descriptor for the event whose
// handling was suspended.
//
// # Discussion
// 
// The effects of mutating or retaining the returned descriptor are undefined,
// although it may be copied. [AppleEventForSuspensionID] may be invoked in
// any thread, not just the one in which the corresponding invocation of
// [SuspendCurrentAppleEvent] occurred.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/appleEvent(forSuspensionID:)
func (a NSAppleEventManager) AppleEventForSuspensionID(suspensionID NSAppleEventManagerSuspensionID) INSAppleEventDescriptor {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("appleEventForSuspensionID:"), suspensionID)
	return NSAppleEventDescriptorFromID(rv)
}
// Given a nonzero `suspensionID` returned by an invocation of
// [SuspendCurrentAppleEvent], returns the corresponding reply event
// descriptor.
//
// # Discussion
// 
// This descriptor, including any mutations, will be returned to the sender of
// the suspended event when handling of the event is resumed, if the sender
// has requested a reply. The effects of retaining the descriptor are
// undefined; it may be copied, but mutations of the copy are returned to the
// sender of the suspended event. [ReplyAppleEventForSuspensionID] may be
// invoked in any thread, not just the one in which the corresponding
// invocation of [SuspendCurrentAppleEvent] occurred.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/replyAppleEvent(forSuspensionID:)
func (a NSAppleEventManager) ReplyAppleEventForSuspensionID(suspensionID NSAppleEventManagerSuspensionID) INSAppleEventDescriptor {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("replyAppleEventForSuspensionID:"), suspensionID)
	return NSAppleEventDescriptorFromID(rv)
}
// Given a nonzero `suspensionID` returned by an invocation of
// [SuspendCurrentAppleEvent], signal that handling of the suspended event may
// now continue.
//
// # Discussion
// 
// This may result in the immediate sending of the reply event to the sender
// of the suspended event, if the sender has requested a reply. If
// `suspensionID` has been used in a previous invocation of
// [SetCurrentAppleEventAndReplyEventWithSuspensionID] the effects of that
// invocation are completely undone. Redundant invocations of
// [ResumeWithSuspensionID] are ignored. Subsequent invocations of other
// [NSAppleEventManager] methods using the same suspension ID are invalid.
// [ResumeWithSuspensionID] may be invoked in any thread, not just the one in
// which the corresponding invocation of [SuspendCurrentAppleEvent] occurred.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/resume(withSuspensionID:)
func (a NSAppleEventManager) ResumeWithSuspensionID(suspensionID NSAppleEventManagerSuspensionID) {
	objc.Send[objc.ID](a.ID, objc.Sel("resumeWithSuspensionID:"), suspensionID)
}
// Given a nonzero `suspensionID` returned by an invocation of
// [SuspendCurrentAppleEvent], sets the values that will be returned by
// subsequent invocations of [CurrentAppleEvent] and [CurrentReplyAppleEvent]
// to be the event whose handling was suspended and its corresponding reply
// event, respectively.
//
// # Discussion
// 
// Redundant invocations of
// [SetCurrentAppleEventAndReplyEventWithSuspensionID] are ignored.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/setCurrentAppleEventAndReplyEventWithSuspensionID(_:)
func (a NSAppleEventManager) SetCurrentAppleEventAndReplyEventWithSuspensionID(suspensionID NSAppleEventManagerSuspensionID) {
	objc.Send[objc.ID](a.ID, objc.Sel("setCurrentAppleEventAndReplyEventWithSuspensionID:"), suspensionID)
}
// Suspends the handling of the current event and returns an ID that must be
// used to resume the handling of the event if an Apple event is being handled
// on the current thread.
//
// # Discussion
// 
// An Apple event is being handled on the current thread if
// [CurrentAppleEvent] does not return `nil`. Returns zero otherwise. The
// suspended event is no longer the current event after this method returns.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/suspendCurrentAppleEvent()
func (a NSAppleEventManager) SuspendCurrentAppleEvent() NSAppleEventManagerSuspensionID {
	rv := objc.Send[NSAppleEventManagerSuspensionID](a.ID, objc.Sel("suspendCurrentAppleEvent"))
	return NSAppleEventManagerSuspensionID(rv)
}

// Returns the single instance of [NSAppleEventManager], creating it first if
// it doesn’t exist.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/shared()
func (_NSAppleEventManagerClass NSAppleEventManagerClass) SharedAppleEventManager() NSAppleEventManager {
	rv := objc.Send[objc.ID](objc.ID(_NSAppleEventManagerClass.class), objc.Sel("sharedAppleEventManager"))
	return NSAppleEventManagerFromID(rv)
}

// Returns the descriptor for `currentAppleEvent` if an Apple event is being
// handled on the current thread.
//
// # Discussion
// 
// An Apple event is being handled on the current thread if a handler that was
// registered with [SetEventHandlerAndSelectorForEventClassAndEventID] is
// being messaged at this instant or
// [SetCurrentAppleEventAndReplyEventWithSuspensionID] has just been invoked.
// Returns `nil` otherwise. The effects of mutating or retaining the returned
// descriptor are undefined, although it may be copied.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/currentAppleEvent
func (a NSAppleEventManager) CurrentAppleEvent() INSAppleEventDescriptor {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("currentAppleEvent"))
	return NSAppleEventDescriptorFromID(objc.ID(rv))
}
// Returns the corresponding reply event descriptor if an Apple event is being
// handled on the current thread.
//
// # Discussion
// 
// An Apple event is being handled on the current thread if
// [CurrentAppleEvent] does not return `nil`. Returns `nil` otherwise. This
// descriptor, including any mutations, will be returned to the sender of the
// current event when all handling of the event has been completed, if the
// sender has requested a reply. The effects of retaining the descriptor are
// undefined; it may be copied, but mutations of the copy are not returned to
// the sender of the current event.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/currentReplyAppleEvent
func (a NSAppleEventManager) CurrentReplyAppleEvent() INSAppleEventDescriptor {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("currentReplyAppleEvent"))
	return NSAppleEventDescriptorFromID(objc.ID(rv))
}

