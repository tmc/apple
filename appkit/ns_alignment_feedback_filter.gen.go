// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSAlignmentFeedbackFilter] class.
var (
	_NSAlignmentFeedbackFilterClass     NSAlignmentFeedbackFilterClass
	_NSAlignmentFeedbackFilterClassOnce sync.Once
)

func getNSAlignmentFeedbackFilterClass() NSAlignmentFeedbackFilterClass {
	_NSAlignmentFeedbackFilterClassOnce.Do(func() {
		_NSAlignmentFeedbackFilterClass = NSAlignmentFeedbackFilterClass{class: objc.GetClass("NSAlignmentFeedbackFilter")}
	})
	return _NSAlignmentFeedbackFilterClass
}

// GetNSAlignmentFeedbackFilterClass returns the class object for NSAlignmentFeedbackFilter.
func GetNSAlignmentFeedbackFilterClass() NSAlignmentFeedbackFilterClass {
	return getNSAlignmentFeedbackFilterClass()
}

type NSAlignmentFeedbackFilterClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAlignmentFeedbackFilterClass) Alloc() NSAlignmentFeedbackFilter {
	rv := objc.Send[NSAlignmentFeedbackFilter](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An object that can filter the movement of an object and provides haptic
// feedback when alignment occurs.
//
// # Overview
// 
// With a Force Touch trackpad, apps can produce tactile feedback to
// complement user actions. If your app implements alignment features, you can
// use the [NSAlignmentFeedbackFilter] class to filter object movements and
// provide haptic feedback to the user at appropriate times. As the user drags
// objects into alignment with a guide or another object, the user actually
// feels a physical bump as the object snaps into place.
// 
// # Implementing Alignment Feedback
// 
// To implement alignment feedback in your custom alignment controller class,
// set up the class to receive events for tracking the movement of an object.
// These can be events matching the [NSAlignmentFeedbackFilter.InputEventMask] value of an
// [NSAlignmentFeedbackFilter] object, or events from a gesture recognizer
// ([NSGestureRecognizer]). For each event received:
// 
// - Create an instance of an [NSAlignmentFeedbackFilter] object. For example:
// 
// - Inform the alignment feedback filter object about the event. To do this,
// call one of the following methods:
// 
// - [NSAlignmentFeedbackFilter.UpdateWithEvent]
// - [NSAlignmentFeedbackFilter.UpdateWithPanRecognizer]
// 
// - Store the location of the object before it moves in response to the
// event. This is considered the location of the object. - Move the object to
// its new location in response to the event. This is the location where the
// object will reside if no alignment occurs. - Store the new location of the
// object. This is considered the location of the object. - Determine where
// the object will move to be aligned. This is considered the location of the
// object. - Request a feedback token based on the previous location, default
// location, and aligned location. To do this, call one of the following
// methods:
// 
// -
// [NSAlignmentFeedbackFilter.AlignmentFeedbackTokenForMovementInViewPreviousPointAlignedPointDefaultPoint]
// - If the object will be moved both horizontally and vertically to become
// aligned. -
// [NSAlignmentFeedbackFilter.AlignmentFeedbackTokenForHorizontalMovementInViewPreviousXAlignedXDefaultX]
// - If the object will be moved horizontally only to become aligned. -
// [NSAlignmentFeedbackFilter.AlignmentFeedbackTokenForVerticalMovementInViewPreviousYAlignedYDefaultY]
// - If the object will be moved vertically only to become aligned.
// 
// - If a feedback token is successfully prepared, call
// [NSAlignmentFeedbackFilter.PerformFeedbackPerformanceTime] to perform the haptic feedback. Then, move
// the object to the aligned location.
// 
// If a value of `null` is returned, rather than a feedback token, then the
// system has determined that alignment and feedback are not appropriate.
// Perhaps the cursor is moving too fast or the distance to the aligned
// location is not significant enough to produce a visual snap. Move the
// object to its default location.
//
// # Informing the Filter About Events
//
//   - [NSAlignmentFeedbackFilter.UpdateWithEvent]: Informs the feedback filter about a new event.
//   - [NSAlignmentFeedbackFilter.UpdateWithPanRecognizer]: Informs the feedback filter about a new pan (drag) gesture recognizer event.
//
// # Preparing Haptic Feedback for Alignment
//
//   - [NSAlignmentFeedbackFilter.AlignmentFeedbackTokenForMovementInViewPreviousPointAlignedPointDefaultPoint]: Requests a feedback token for the alignment of an object requiring horizontal and vertical movement.
//   - [NSAlignmentFeedbackFilter.AlignmentFeedbackTokenForHorizontalMovementInViewPreviousXAlignedXDefaultX]: Requests a feedback token for the alignment of an object requiring horizontal movement only.
//   - [NSAlignmentFeedbackFilter.AlignmentFeedbackTokenForVerticalMovementInViewPreviousYAlignedYDefaultY]: Requests a feedback token for the alignment of an object requiring vertical movement only.
//
// # Providing Feedback to the User
//
//   - [NSAlignmentFeedbackFilter.PerformFeedbackPerformanceTime]: Performs the haptic feedback described by one or more alignment feedback tokens.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlignmentFeedbackFilter
type NSAlignmentFeedbackFilter struct {
	objectivec.Object
}

// NSAlignmentFeedbackFilterFromID constructs a [NSAlignmentFeedbackFilter] from an objc.ID.
//
// An object that can filter the movement of an object and provides haptic
// feedback when alignment occurs.
func NSAlignmentFeedbackFilterFromID(id objc.ID) NSAlignmentFeedbackFilter {
	return NSAlignmentFeedbackFilter{objectivec.Object{id}}
}
// NOTE: NSAlignmentFeedbackFilter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSAlignmentFeedbackFilter] class.
//
// # Informing the Filter About Events
//
//   - [INSAlignmentFeedbackFilter.UpdateWithEvent]: Informs the feedback filter about a new event.
//   - [INSAlignmentFeedbackFilter.UpdateWithPanRecognizer]: Informs the feedback filter about a new pan (drag) gesture recognizer event.
//
// # Preparing Haptic Feedback for Alignment
//
//   - [INSAlignmentFeedbackFilter.AlignmentFeedbackTokenForMovementInViewPreviousPointAlignedPointDefaultPoint]: Requests a feedback token for the alignment of an object requiring horizontal and vertical movement.
//   - [INSAlignmentFeedbackFilter.AlignmentFeedbackTokenForHorizontalMovementInViewPreviousXAlignedXDefaultX]: Requests a feedback token for the alignment of an object requiring horizontal movement only.
//   - [INSAlignmentFeedbackFilter.AlignmentFeedbackTokenForVerticalMovementInViewPreviousYAlignedYDefaultY]: Requests a feedback token for the alignment of an object requiring vertical movement only.
//
// # Providing Feedback to the User
//
//   - [INSAlignmentFeedbackFilter.PerformFeedbackPerformanceTime]: Performs the haptic feedback described by one or more alignment feedback tokens.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlignmentFeedbackFilter
type INSAlignmentFeedbackFilter interface {
	objectivec.IObject

	// Topic: Informing the Filter About Events

	// Informs the feedback filter about a new event.
	UpdateWithEvent(event INSEvent)
	// Informs the feedback filter about a new pan (drag) gesture recognizer event.
	UpdateWithPanRecognizer(panRecognizer INSPanGestureRecognizer)

	// Topic: Preparing Haptic Feedback for Alignment

	// Requests a feedback token for the alignment of an object requiring horizontal and vertical movement.
	AlignmentFeedbackTokenForMovementInViewPreviousPointAlignedPointDefaultPoint(view INSView, previousPoint corefoundation.CGPoint, alignedPoint corefoundation.CGPoint, defaultPoint corefoundation.CGPoint) NSAlignmentFeedbackToken
	// Requests a feedback token for the alignment of an object requiring horizontal movement only.
	AlignmentFeedbackTokenForHorizontalMovementInViewPreviousXAlignedXDefaultX(view INSView, previousX float64, alignedX float64, defaultX float64) NSAlignmentFeedbackToken
	// Requests a feedback token for the alignment of an object requiring vertical movement only.
	AlignmentFeedbackTokenForVerticalMovementInViewPreviousYAlignedYDefaultY(view INSView, previousY float64, alignedY float64, defaultY float64) NSAlignmentFeedbackToken

	// Topic: Providing Feedback to the User

	// Performs the haptic feedback described by one or more alignment feedback tokens.
	PerformFeedbackPerformanceTime(alignmentFeedbackTokens []objectivec.IObject, performanceTime NSHapticFeedbackPerformanceTime)
}





// Init initializes the instance.
func (a NSAlignmentFeedbackFilter) Init() NSAlignmentFeedbackFilter {
	rv := objc.Send[NSAlignmentFeedbackFilter](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAlignmentFeedbackFilter) Autorelease() NSAlignmentFeedbackFilter {
	rv := objc.Send[NSAlignmentFeedbackFilter](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAlignmentFeedbackFilter creates a new NSAlignmentFeedbackFilter instance.
func NewNSAlignmentFeedbackFilter() NSAlignmentFeedbackFilter {
	class := getNSAlignmentFeedbackFilterClass()
	rv := objc.Send[NSAlignmentFeedbackFilter](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Informs the feedback filter about a new event.
//
// event: An event ([NSEvent]) to be filtered, which matches an event type returned
// by the [InputEventMask] method.
//
// # Discussion
// 
// This method informs the feedback filter about a new event to be filtered,
// which matches an event type returned by the [InputEventMask] method. Call
// the `` method instead of [UpdateWithPanRecognizer] if you are using a
// tracking loop controller for event tracking.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlignmentFeedbackFilter/update(with:)
func (a NSAlignmentFeedbackFilter) UpdateWithEvent(event INSEvent) {
	objc.Send[objc.ID](a.ID, objc.Sel("updateWithEvent:"), event)
}

// Informs the feedback filter about a new pan (drag) gesture recognizer
// event.
//
// panRecognizer: The gesture recognizer ([NSPanGestureRecognizer]) that produced the event.
//
// # Discussion
// 
// This method informs the feedback filter about a new pan (drag) gesture
// recognizer event. Call this method instead of [UpdateWithEvent] if your
// event tracking uses gesture recognizers.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlignmentFeedbackFilter/update(withPanRecognizer:)
func (a NSAlignmentFeedbackFilter) UpdateWithPanRecognizer(panRecognizer INSPanGestureRecognizer) {
	objc.Send[objc.ID](a.ID, objc.Sel("updateWithPanRecognizer:"), panRecognizer)
}

// Requests a feedback token for the alignment of an object requiring
// horizontal and vertical movement.
//
// view: The view ([NSView]) in which the object was moved.
//
// previousPoint: The location ([NSPoint]) of the object prior to its move.
// //
// [NSPoint]: https://developer.apple.com/documentation/Foundation/NSPoint
//
// alignedPoint: The new location ([NSPoint]) of the object if alignment occurs.
// //
// [NSPoint]: https://developer.apple.com/documentation/Foundation/NSPoint
//
// defaultPoint: The current location ([NSPoint]) of the object. This is where the user
// actually moved the object. This location may be offset from the location of
// the cursor.
// //
// [NSPoint]: https://developer.apple.com/documentation/Foundation/NSPoint
//
// # Return Value
// 
// A null value if the system determines that the alignment should not occur.
// Otherwise, a feedback token of type [NSAlignmentFeedbackToken] is returned.
//
// # Discussion
// 
// This method requests a feedback token for the alignment of an object
// requiring horizontal and vertical movement.
// 
// If a feedback token is returned, call [PerformFeedbackPerformanceTime] to
// initiate haptic feedback. Then, move the object to its aligned location.
// 
// If no feedback token is returned, don’t perform the alignment or request
// haptic feedback. Even if this joint horizontal and vertical alignment
// fails, be sure to check other alignments. For example, an individual
// horizontal or vertical alignment may still be allowed. If no alignments
// will occur, move the object to its default location.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlignmentFeedbackFilter/alignmentFeedbackTokenForMovement(in:previousPoint:alignedPoint:defaultPoint:)
func (a NSAlignmentFeedbackFilter) AlignmentFeedbackTokenForMovementInViewPreviousPointAlignedPointDefaultPoint(view INSView, previousPoint corefoundation.CGPoint, alignedPoint corefoundation.CGPoint, defaultPoint corefoundation.CGPoint) NSAlignmentFeedbackToken {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("alignmentFeedbackTokenForMovementInView:previousPoint:alignedPoint:defaultPoint:"), view, previousPoint, alignedPoint, defaultPoint)
	return NSAlignmentFeedbackTokenObjectFromID(rv)
}

// Requests a feedback token for the alignment of an object requiring
// horizontal movement only.
//
// view: The view ([NSView]) in which the object was moved.
//
// previousX: The horizontal location ([CGFloat]) of the object prior to its move.
// //
// [CGFloat]: https://developer.apple.com/documentation/CoreFoundation/CGFloat-swift.struct
//
// alignedX: The new horizontal location ([CGFloat]) of the object if alignment occurs.
// //
// [CGFloat]: https://developer.apple.com/documentation/CoreFoundation/CGFloat-swift.struct
//
// defaultX: The current horizontal location ([CGFloat]) of the object. This is where
// the user actually moved the object. This location may be offset from the
// location of the cursor.
// //
// [CGFloat]: https://developer.apple.com/documentation/CoreFoundation/CGFloat-swift.struct
//
// # Return Value
// 
// If the system determines that the alignment should not occur, a null value
// is returned. Otherwise, a feedback token of type [NSAlignmentFeedbackToken]
// is returned.
//
// # Discussion
// 
// This method requests a feedback token for the alignment of an object
// requiring horizontal movement only.
// 
// If a feedback token is returned, call [PerformFeedbackPerformanceTime] to
// initiate haptic feedback. Then, move the object to its aligned location.
// 
// If no feedback token is returned, don’t perform the horizontal alignment
// or request haptic feedback. Even if this horizontal alignment fails, be
// sure to check other alignments. For example, a vertical alignment may still
// be allowed. If no alignments will occur, move the object to its default
// location.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlignmentFeedbackFilter/alignmentFeedbackTokenForHorizontalMovement(in:previousX:alignedX:defaultX:)
func (a NSAlignmentFeedbackFilter) AlignmentFeedbackTokenForHorizontalMovementInViewPreviousXAlignedXDefaultX(view INSView, previousX float64, alignedX float64, defaultX float64) NSAlignmentFeedbackToken {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("alignmentFeedbackTokenForHorizontalMovementInView:previousX:alignedX:defaultX:"), view, previousX, alignedX, defaultX)
	return NSAlignmentFeedbackTokenObjectFromID(rv)
}

// Requests a feedback token for the alignment of an object requiring vertical
// movement only.
//
// view: The view ([NSView]) in which the object was moved.
//
// previousY: The vertical location ([CGFloat]) of the object prior to its move.
// //
// [CGFloat]: https://developer.apple.com/documentation/CoreFoundation/CGFloat-swift.struct
//
// alignedY: The new vertical location ([CGFloat]) of the object if alignment occurs.
// //
// [CGFloat]: https://developer.apple.com/documentation/CoreFoundation/CGFloat-swift.struct
//
// defaultY: The current vertical location ([CGFloat]) of the object. This is where the
// user actually moved the object. This location may be offset from the
// location of the cursor.
// //
// [CGFloat]: https://developer.apple.com/documentation/CoreFoundation/CGFloat-swift.struct
//
// # Return Value
// 
// If the system determines that the alignment should not occur, a null value
// is returned. Otherwise, a feedback token of type [NSAlignmentFeedbackToken]
// is returned.
//
// # Discussion
// 
// This method requests a feedback token for the alignment of an object
// requiring vertical movement only.
// 
// If a feedback token is returned, call [PerformFeedbackPerformanceTime] to
// initiate haptic feedback. Then, move the object to its aligned location.
// 
// If no feedback token is returned, don’t perform the vertical alignment or
// request haptic feedback. Even if this vertical alignment fails, be sure to
// check other alignments. For example, a horizontal alignment may still be
// allowed. If no alignments will occur, move the object to its default
// location.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlignmentFeedbackFilter/alignmentFeedbackTokenForVerticalMovement(in:previousY:alignedY:defaultY:)
func (a NSAlignmentFeedbackFilter) AlignmentFeedbackTokenForVerticalMovementInViewPreviousYAlignedYDefaultY(view INSView, previousY float64, alignedY float64, defaultY float64) NSAlignmentFeedbackToken {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("alignmentFeedbackTokenForVerticalMovementInView:previousY:alignedY:defaultY:"), view, previousY, alignedY, defaultY)
	return NSAlignmentFeedbackTokenObjectFromID(rv)
}

// Performs the haptic feedback described by one or more alignment feedback
// tokens.
//
// alignmentFeedbackTokens: One or more feedback tokens prepared for specific alignment locations by
// calling
// [AlignmentFeedbackTokenForMovementInViewPreviousPointAlignedPointDefaultPoint],
// [AlignmentFeedbackTokenForHorizontalMovementInViewPreviousXAlignedXDefaultX],
// or
// [AlignmentFeedbackTokenForVerticalMovementInViewPreviousYAlignedYDefaultY].
// Typically, no more than one feedback token per dimension
// (horizontal/vertical) should be provided.
//
// performanceTime: The time, of type [NSHapticFeedbackPerformanceTime], when the feedback
// should be provided to the user.
//
// # Discussion
// 
// Call this method to initiate haptic feedback to the user. Pass it one or
// more alignment feedback tokens of type [NSAlignmentFeedbackToken] and a
// time to execute of type [NSHapticFeedbackPerformanceTime]. Call this method
// immediately before moving the object to its new aligned position.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlignmentFeedbackFilter/performFeedback(_:performanceTime:)
func (a NSAlignmentFeedbackFilter) PerformFeedbackPerformanceTime(alignmentFeedbackTokens []objectivec.IObject, performanceTime NSHapticFeedbackPerformanceTime) {
	objc.Send[objc.ID](a.ID, objc.Sel("performFeedback:performanceTime:"), objectivec.IObjectSliceToNSArray(alignmentFeedbackTokens), performanceTime)
}
















// Retrieves the event types the filter accepts.
//
// # Discussion
// 
// This method retrieves the event types that the filter accepts. This
// information may be used by an event tracking loop to watch for events that
// can be passed to the filter.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlignmentFeedbackFilter/inputEventMask
func (_NSAlignmentFeedbackFilterClass NSAlignmentFeedbackFilterClass) InputEventMask() NSEventMask {
	rv := objc.Send[NSEventMask](objc.ID(_NSAlignmentFeedbackFilterClass.class), objc.Sel("inputEventMask"))
	return NSEventMask(rv)
}



















