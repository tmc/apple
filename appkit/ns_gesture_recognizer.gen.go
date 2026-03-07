// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSGestureRecognizer] class.
var (
	_NSGestureRecognizerClass     NSGestureRecognizerClass
	_NSGestureRecognizerClassOnce sync.Once
)

func getNSGestureRecognizerClass() NSGestureRecognizerClass {
	_NSGestureRecognizerClassOnce.Do(func() {
		_NSGestureRecognizerClass = NSGestureRecognizerClass{class: objc.GetClass("NSGestureRecognizer")}
	})
	return _NSGestureRecognizerClass
}

// GetNSGestureRecognizerClass returns the class object for NSGestureRecognizer.
func GetNSGestureRecognizerClass() NSGestureRecognizerClass {
	return getNSGestureRecognizerClass()
}

type NSGestureRecognizerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSGestureRecognizerClass) Alloc() NSGestureRecognizer {
	rv := objc.Send[NSGestureRecognizer](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An object that monitors events and calls its action method when a
// predefined sequence of events occur.
//
// # Overview
// 
// A gesture recognizer might recognize a single click, a click and drag, or a
// sequence of events that imply rotation. You do not create instances of this
// class directly. This class is an abstract base class that defines the
// common behavior for all gesture recognizers. When using a gesture
// recognizer in your app, create an instance of one of the concrete
// subclasses.
// 
// The concrete subclasses of [NSGestureRecognizer] are the following:
// 
// - [NSClickGestureRecognizer] - [NSMagnificationGestureRecognizer] -
// [NSPanGestureRecognizer] - [NSPressGestureRecognizer] -
// [NSRotationGestureRecognizer]
// 
// A gesture recognizer operates on events in a specific view (or in any of
// that view’s subviews). After creating a gesture recognizer, attach it to
// one of your views using the [AddGestureRecognizer] method. Events received
// by your app are forwarded automatically to any relevant gesture recognizers
// before they are sent to the corresponding view. The gesture recognizer can
// delay the further progression of the events until recognition is complete
// or allow the events to be delivered normally.
// 
// A gesture recognizer can detect gestures that are either discrete or
// continuous in nature. A click gesture is discrete because it involves a
// mouse-down and mouse-up event without any mouse movements in between. By
// contrast, a pan or rotation gesture is continuous because it involves
// tracking mouse movements over a period of time.
// 
// During the gesture recognition process, a gesture recognizer calls the
// action method of its associated target object to report the state of the
// recognition process. For discrete gestures, the action method is typically
// called only once when the gesture is recognized. For continuous gestures,
// it may be called multiple times depending on the current state of the
// gesture recognizer. In that situation, you can use your action method to
// perform appropriate tasks, such as creating animations for any
// mouse-related movements, in addition to handling the final results of the
// gesture recognition process.
// 
// A gesture recognizer has only one action method and one target object, and
// the method must conform to one of the following signatures:
// 
// When your code needs additional information about the particulars of a
// gesture, define your action method to include the gesture recognizer
// parameter. You almost always want the gesture recognizer object when
// handling continuous gestures. For example, for a rotation gesture, you
// would use the gesture recognizer object to get the updated rotation value.
// You can also use the gesture recognizer object to get the location of where
// the gesture occurred.
// 
// # State Transitions
// 
// Gesture recognizers operate within a predefined state machine,
// transitioning from state to state as they handle events. All gesture
// recognizers begin in the Possible ([NSGestureRecognizer.State.possible])
// state, but the possible transitions differ for continuous and discrete
// gestures.
// 
// Discrete gestures transition from the Possible state directly to the
// Recognized ([recognized]) or Failed ([NSGestureRecognizer.State.failed])
// state, depending on whether they successfully interpret the gesture. When a
// discrete gesture recognizer transitions to the Recognized state, it calls
// the action method of its target object.
// 
// For continuous gestures, the state transitions are as follows:
// 
// - Possible —> Began —> [Changed] —> Cancelled - Possible —> Began
// —> [Changed] —> Ended
// 
// The Changed state is optional and may occur multiple times before the
// Cancelled or Ended state is reached. Many state transitions cause the
// gesture recognizer to call its action method. Setting the [NSGestureRecognizer.State] property
// to [NSGestureRecognizer.State.changed] while monitoring events also calls
// the action method. You can use these calls to update the state of your app
// or update any custom animations.
// 
// For a list of possible states, see the constants in
// [NSGestureRecognizer.State].
// 
// # Subclassing Notes
// 
// You may create a subclass of [NSGestureRecognizer] that recognizes a
// distinctive gesture—for example, a “check mark” gesture. A custom
// gesture recognizer implements any appropriate event-related methods to
// detect its gesture along with a few other methods for managing state
// information.
// 
// All gesture recognizers must update the value in the state property at
// appropriate times. Specifically, you must update it for all state
// transitions. For more information about the possible state transitions of a
// gesture recognizer, see [NSGestureRecognizer].
// 
// # Methods to Override
// 
// When creating your own gesture recognizer subclass:
// 
// - Implement the [NSGestureRecognizer.Reset] method and any other relevant methods in Methods
// for Subclasses. - Override the [NSGestureRecognizer.LocationInView] method as needed to specify
// an appropriate point for your gesture.
// 
// AppKit waits for a mouse-down event, magnify event, or rotation event to
// occur before starting the gesture recognition process. A gesture recognizer
// that used only key-down events to recognize its gesture would not have its
// [NSGestureRecognizer.KeyDown] method called until a mouse-down, magnify, or rotation event
// started the recognition process.
// 
// # Alternatives to Subclassing
// 
// The [NSGestureRecognizer] class defines the common behaviors that can be
// configured for all concrete gesture recognizers. It also supports a
// delegate—an object that adopts the [NSGestureRecognizerDelegate]
// protocol—for handling finer-grained customization of some behaviors
// without the need for subclassing. For example, you can use the delegate to
// create dependencies between specific gesture recognizer objects.
// 
// For more information about using the delegate to control the behavior of
// your gesture recognizers, see [NSGestureRecognizerDelegate].
//
// [NSGestureRecognizer.State.changed]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/State-swift.enum/changed
// [NSGestureRecognizer.State.failed]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/State-swift.enum/failed
// [NSGestureRecognizer.State.possible]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/State-swift.enum/possible
// [NSGestureRecognizer.State]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/State-swift.enum
// [recognized]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/State-swift.enum/recognized
//
// # Initializing a Gesture Recognizer
//
//   - [NSGestureRecognizer.InitWithTargetAction]: Initializes the gesture recognizer with the specified target and action information.
//
// # Accessing the Target and Action
//
//   - [NSGestureRecognizer.Action]: The action method to call when the gesture is recognized.
//   - [NSGestureRecognizer.SetAction]
//   - [NSGestureRecognizer.Target]: The object that implements the action method.
//   - [NSGestureRecognizer.SetTarget]
//
// # Getting the Location of Events
//
//   - [NSGestureRecognizer.LocationInView]: Returns the point computed as the location of the gesture.
//
// # Accessing the Recognizer’s State
//
//   - [NSGestureRecognizer.State]: The current state of the gesture recognizer.
//   - [NSGestureRecognizer.View]: The view to which the gesture recognizer is attached.
//   - [NSGestureRecognizer.Enabled]: A Boolean value indicating whether the gesture recognizer is able to handle events.
//   - [NSGestureRecognizer.SetEnabled]
//
// # Delaying Events
//
//   - [NSGestureRecognizer.DelaysPrimaryMouseButtonEvents]: A Boolean value that indicates whether primary mouse button events are delivered only after gesture recognition fails.
//   - [NSGestureRecognizer.SetDelaysPrimaryMouseButtonEvents]
//   - [NSGestureRecognizer.DelaysSecondaryMouseButtonEvents]: A Boolean value that indicates whether secondary mouse button events are delivered only after gesture recognition fails.
//   - [NSGestureRecognizer.SetDelaysSecondaryMouseButtonEvents]
//   - [NSGestureRecognizer.DelaysOtherMouseButtonEvents]: A Boolean value that indicates whether other mouse button events are delivered only after gesture recognition fails.
//   - [NSGestureRecognizer.SetDelaysOtherMouseButtonEvents]
//   - [NSGestureRecognizer.DelaysKeyEvents]: A Boolean value that indicates whether key events are delivered only after gesture recognition fails.
//   - [NSGestureRecognizer.SetDelaysKeyEvents]
//   - [NSGestureRecognizer.DelaysMagnificationEvents]: A Boolean value that indicates whether magnification events are delivered only after gesture recognition fails.
//   - [NSGestureRecognizer.SetDelaysMagnificationEvents]
//   - [NSGestureRecognizer.DelaysRotationEvents]: A Boolean value that indicates whether rotation events are delivered only after gesture recognition fails.
//   - [NSGestureRecognizer.SetDelaysRotationEvents]
//
// # Accessing the Delegate
//
//   - [NSGestureRecognizer.Delegate]: The delegate of the gesture recognizer.
//   - [NSGestureRecognizer.SetDelegate]
//
// # Methods for Subclasses
//
//   - [NSGestureRecognizer.Reset]: Overridden to reset the internal state of the gesture recognizer when an attempt completes.
//   - [NSGestureRecognizer.MouseDown]: Informs the gesture recognizer that the user pressed the left mouse button.
//   - [NSGestureRecognizer.MouseDragged]: Informs the gesture recognizer that the user moved the mouse with the left button pressed.
//   - [NSGestureRecognizer.MouseUp]: Informs the gesture recognizer that the user released the left mouse button.
//   - [NSGestureRecognizer.OtherMouseDown]: Informs the gesture recognizer that the user pressed a mouse button other than the left or right one.
//   - [NSGestureRecognizer.OtherMouseDragged]: Informs the gesture recognizer that the user moved the mouse with a button other than the left or right one pressed.
//   - [NSGestureRecognizer.OtherMouseUp]: Informs the gesture recognizer that the user released a mouse button other than the left or right one.
//   - [NSGestureRecognizer.RightMouseDown]: Informs the gesture recognizer that the user pressed the right mouse button.
//   - [NSGestureRecognizer.RightMouseDragged]: Informs the gesture recognizer that the user moved the mouse with the right button pressed.
//   - [NSGestureRecognizer.RightMouseUp]: Informs the gesture recognizer that the user released the right mouse button.
//   - [NSGestureRecognizer.MagnifyWithEvent]: Informs the gesture recognizer that the user is performing a pinch gesture.
//   - [NSGestureRecognizer.RotateWithEvent]: Informs the gesture recognizer that the user is performing a rotation gesture.
//   - [NSGestureRecognizer.CanBePreventedByGestureRecognizer]: Overridden to indicate that the specified gesture recognizer can prevent the current object from recognizing a gesture.
//   - [NSGestureRecognizer.CanPreventGestureRecognizer]: Overridden to indicate that the current object can prevent the specified gesture recognizer from recognizing its gesture.
//   - [NSGestureRecognizer.ShouldBeRequiredToFailByGestureRecognizer]: Overridden to indicate that the current object must fail before the specified gesture recognizer begins recognizing its gesture.
//   - [NSGestureRecognizer.ShouldRequireFailureOfGestureRecognizer]: Overridden to indicate that the specified gesture recognizer must fail before the current object begins recognizing its gesture.
//   - [NSGestureRecognizer.KeyDown]: Informs the gesture recognizer that the user has pressed a key.
//   - [NSGestureRecognizer.KeyUp]: Informs the gesture recognizer that the user released a key.
//   - [NSGestureRecognizer.TabletPoint]: Informs the user that a tablet-point event occurred.
//   - [NSGestureRecognizer.FlagsChanged]: Informs the current object that the user pressed or released a modifier key (Shift, Control, and so on).
//   - [NSGestureRecognizer.PressureChangeWithEvent]: Informs the current object that a pressure change occurred on a system that supports pressure sensitivity.
//
// # Configuring Pressure
//
//   - [NSGestureRecognizer.PressureConfiguration]: Configures the behavior and progression of the Force Touch trackpad when responding to recognized pressure gestures.
//   - [NSGestureRecognizer.SetPressureConfiguration]
//
// # Initializers
//
//   - [NSGestureRecognizer.InitWithCoder]
//
// # Instance Properties
//
//   - [NSGestureRecognizer.AllowedTouchTypes]
//   - [NSGestureRecognizer.SetAllowedTouchTypes]
//   - [NSGestureRecognizer.ModifierFlags]
//   - [NSGestureRecognizer.Name]
//   - [NSGestureRecognizer.SetName]
//
// # Instance Methods
//
//   - [NSGestureRecognizer.TouchesBeganWithEvent]: Called when one or more fingers first make contact with an [NSTouchBar](<doc://com.apple.appkit/documentation/AppKit/NSTouchBar>) instance on the Touch Bar.
//   - [NSGestureRecognizer.TouchesCancelledWithEvent]: Called when a system event, such as a low-memory warning, cancels an in-progress touch event in an [NSTouchBar](<doc://com.apple.appkit/documentation/AppKit/NSTouchBar>) object.
//   - [NSGestureRecognizer.TouchesEndedWithEvent]: Called when one or more fingers are removed from contact with an [NSTouchBar](<doc://com.apple.appkit/documentation/AppKit/NSTouchBar>) instance on the Touch Bar.
//   - [NSGestureRecognizer.TouchesMovedWithEvent]: Called when one or more fingers, associated with an in-progress event, move within an [NSTouchBar](<doc://com.apple.appkit/documentation/AppKit/NSTouchBar>) instance on the Touch Bar.
//   - [NSGestureRecognizer.MouseCancelled]
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer
type NSGestureRecognizer struct {
	objectivec.Object
}

// NSGestureRecognizerFromID constructs a [NSGestureRecognizer] from an objc.ID.
//
// An object that monitors events and calls its action method when a
// predefined sequence of events occur.
func NSGestureRecognizerFromID(id objc.ID) NSGestureRecognizer {
	return NSGestureRecognizer{objectivec.Object{id}}
}
// NOTE: NSGestureRecognizer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSGestureRecognizer] class.
//
// # Initializing a Gesture Recognizer
//
//   - [INSGestureRecognizer.InitWithTargetAction]: Initializes the gesture recognizer with the specified target and action information.
//
// # Accessing the Target and Action
//
//   - [INSGestureRecognizer.Action]: The action method to call when the gesture is recognized.
//   - [INSGestureRecognizer.SetAction]
//   - [INSGestureRecognizer.Target]: The object that implements the action method.
//   - [INSGestureRecognizer.SetTarget]
//
// # Getting the Location of Events
//
//   - [INSGestureRecognizer.LocationInView]: Returns the point computed as the location of the gesture.
//
// # Accessing the Recognizer’s State
//
//   - [INSGestureRecognizer.State]: The current state of the gesture recognizer.
//   - [INSGestureRecognizer.View]: The view to which the gesture recognizer is attached.
//   - [INSGestureRecognizer.Enabled]: A Boolean value indicating whether the gesture recognizer is able to handle events.
//   - [INSGestureRecognizer.SetEnabled]
//
// # Delaying Events
//
//   - [INSGestureRecognizer.DelaysPrimaryMouseButtonEvents]: A Boolean value that indicates whether primary mouse button events are delivered only after gesture recognition fails.
//   - [INSGestureRecognizer.SetDelaysPrimaryMouseButtonEvents]
//   - [INSGestureRecognizer.DelaysSecondaryMouseButtonEvents]: A Boolean value that indicates whether secondary mouse button events are delivered only after gesture recognition fails.
//   - [INSGestureRecognizer.SetDelaysSecondaryMouseButtonEvents]
//   - [INSGestureRecognizer.DelaysOtherMouseButtonEvents]: A Boolean value that indicates whether other mouse button events are delivered only after gesture recognition fails.
//   - [INSGestureRecognizer.SetDelaysOtherMouseButtonEvents]
//   - [INSGestureRecognizer.DelaysKeyEvents]: A Boolean value that indicates whether key events are delivered only after gesture recognition fails.
//   - [INSGestureRecognizer.SetDelaysKeyEvents]
//   - [INSGestureRecognizer.DelaysMagnificationEvents]: A Boolean value that indicates whether magnification events are delivered only after gesture recognition fails.
//   - [INSGestureRecognizer.SetDelaysMagnificationEvents]
//   - [INSGestureRecognizer.DelaysRotationEvents]: A Boolean value that indicates whether rotation events are delivered only after gesture recognition fails.
//   - [INSGestureRecognizer.SetDelaysRotationEvents]
//
// # Accessing the Delegate
//
//   - [INSGestureRecognizer.Delegate]: The delegate of the gesture recognizer.
//   - [INSGestureRecognizer.SetDelegate]
//
// # Methods for Subclasses
//
//   - [INSGestureRecognizer.Reset]: Overridden to reset the internal state of the gesture recognizer when an attempt completes.
//   - [INSGestureRecognizer.MouseDown]: Informs the gesture recognizer that the user pressed the left mouse button.
//   - [INSGestureRecognizer.MouseDragged]: Informs the gesture recognizer that the user moved the mouse with the left button pressed.
//   - [INSGestureRecognizer.MouseUp]: Informs the gesture recognizer that the user released the left mouse button.
//   - [INSGestureRecognizer.OtherMouseDown]: Informs the gesture recognizer that the user pressed a mouse button other than the left or right one.
//   - [INSGestureRecognizer.OtherMouseDragged]: Informs the gesture recognizer that the user moved the mouse with a button other than the left or right one pressed.
//   - [INSGestureRecognizer.OtherMouseUp]: Informs the gesture recognizer that the user released a mouse button other than the left or right one.
//   - [INSGestureRecognizer.RightMouseDown]: Informs the gesture recognizer that the user pressed the right mouse button.
//   - [INSGestureRecognizer.RightMouseDragged]: Informs the gesture recognizer that the user moved the mouse with the right button pressed.
//   - [INSGestureRecognizer.RightMouseUp]: Informs the gesture recognizer that the user released the right mouse button.
//   - [INSGestureRecognizer.MagnifyWithEvent]: Informs the gesture recognizer that the user is performing a pinch gesture.
//   - [INSGestureRecognizer.RotateWithEvent]: Informs the gesture recognizer that the user is performing a rotation gesture.
//   - [INSGestureRecognizer.CanBePreventedByGestureRecognizer]: Overridden to indicate that the specified gesture recognizer can prevent the current object from recognizing a gesture.
//   - [INSGestureRecognizer.CanPreventGestureRecognizer]: Overridden to indicate that the current object can prevent the specified gesture recognizer from recognizing its gesture.
//   - [INSGestureRecognizer.ShouldBeRequiredToFailByGestureRecognizer]: Overridden to indicate that the current object must fail before the specified gesture recognizer begins recognizing its gesture.
//   - [INSGestureRecognizer.ShouldRequireFailureOfGestureRecognizer]: Overridden to indicate that the specified gesture recognizer must fail before the current object begins recognizing its gesture.
//   - [INSGestureRecognizer.KeyDown]: Informs the gesture recognizer that the user has pressed a key.
//   - [INSGestureRecognizer.KeyUp]: Informs the gesture recognizer that the user released a key.
//   - [INSGestureRecognizer.TabletPoint]: Informs the user that a tablet-point event occurred.
//   - [INSGestureRecognizer.FlagsChanged]: Informs the current object that the user pressed or released a modifier key (Shift, Control, and so on).
//   - [INSGestureRecognizer.PressureChangeWithEvent]: Informs the current object that a pressure change occurred on a system that supports pressure sensitivity.
//
// # Configuring Pressure
//
//   - [INSGestureRecognizer.PressureConfiguration]: Configures the behavior and progression of the Force Touch trackpad when responding to recognized pressure gestures.
//   - [INSGestureRecognizer.SetPressureConfiguration]
//
// # Initializers
//
//   - [INSGestureRecognizer.InitWithCoder]
//
// # Instance Properties
//
//   - [INSGestureRecognizer.AllowedTouchTypes]
//   - [INSGestureRecognizer.SetAllowedTouchTypes]
//   - [INSGestureRecognizer.ModifierFlags]
//   - [INSGestureRecognizer.Name]
//   - [INSGestureRecognizer.SetName]
//
// # Instance Methods
//
//   - [INSGestureRecognizer.TouchesBeganWithEvent]: Called when one or more fingers first make contact with an [NSTouchBar](<doc://com.apple.appkit/documentation/AppKit/NSTouchBar>) instance on the Touch Bar.
//   - [INSGestureRecognizer.TouchesCancelledWithEvent]: Called when a system event, such as a low-memory warning, cancels an in-progress touch event in an [NSTouchBar](<doc://com.apple.appkit/documentation/AppKit/NSTouchBar>) object.
//   - [INSGestureRecognizer.TouchesEndedWithEvent]: Called when one or more fingers are removed from contact with an [NSTouchBar](<doc://com.apple.appkit/documentation/AppKit/NSTouchBar>) instance on the Touch Bar.
//   - [INSGestureRecognizer.TouchesMovedWithEvent]: Called when one or more fingers, associated with an in-progress event, move within an [NSTouchBar](<doc://com.apple.appkit/documentation/AppKit/NSTouchBar>) instance on the Touch Bar.
//   - [INSGestureRecognizer.MouseCancelled]
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer
type INSGestureRecognizer interface {
	objectivec.IObject

	// Topic: Initializing a Gesture Recognizer

	// Initializes the gesture recognizer with the specified target and action information.
	InitWithTargetAction(target objectivec.IObject, action objc.SEL) NSGestureRecognizer

	// Topic: Accessing the Target and Action

	// The action method to call when the gesture is recognized.
	Action() objc.SEL
	SetAction(value objc.SEL)
	// The object that implements the action method.
	Target() objectivec.IObject
	SetTarget(value objectivec.IObject)

	// Topic: Getting the Location of Events

	// Returns the point computed as the location of the gesture.
	LocationInView(view INSView) corefoundation.CGPoint

	// Topic: Accessing the Recognizer’s State

	// The current state of the gesture recognizer.
	State() NSGestureRecognizerState
	// The view to which the gesture recognizer is attached.
	View() INSView
	// A Boolean value indicating whether the gesture recognizer is able to handle events.
	Enabled() bool
	SetEnabled(value bool)

	// Topic: Delaying Events

	// A Boolean value that indicates whether primary mouse button events are delivered only after gesture recognition fails.
	DelaysPrimaryMouseButtonEvents() bool
	SetDelaysPrimaryMouseButtonEvents(value bool)
	// A Boolean value that indicates whether secondary mouse button events are delivered only after gesture recognition fails.
	DelaysSecondaryMouseButtonEvents() bool
	SetDelaysSecondaryMouseButtonEvents(value bool)
	// A Boolean value that indicates whether other mouse button events are delivered only after gesture recognition fails.
	DelaysOtherMouseButtonEvents() bool
	SetDelaysOtherMouseButtonEvents(value bool)
	// A Boolean value that indicates whether key events are delivered only after gesture recognition fails.
	DelaysKeyEvents() bool
	SetDelaysKeyEvents(value bool)
	// A Boolean value that indicates whether magnification events are delivered only after gesture recognition fails.
	DelaysMagnificationEvents() bool
	SetDelaysMagnificationEvents(value bool)
	// A Boolean value that indicates whether rotation events are delivered only after gesture recognition fails.
	DelaysRotationEvents() bool
	SetDelaysRotationEvents(value bool)

	// Topic: Accessing the Delegate

	// The delegate of the gesture recognizer.
	Delegate() NSGestureRecognizerDelegate
	SetDelegate(value NSGestureRecognizerDelegate)

	// Topic: Methods for Subclasses

	// Overridden to reset the internal state of the gesture recognizer when an attempt completes.
	Reset()
	// Informs the gesture recognizer that the user pressed the left mouse button.
	MouseDown(event INSEvent)
	// Informs the gesture recognizer that the user moved the mouse with the left button pressed.
	MouseDragged(event INSEvent)
	// Informs the gesture recognizer that the user released the left mouse button.
	MouseUp(event INSEvent)
	// Informs the gesture recognizer that the user pressed a mouse button other than the left or right one.
	OtherMouseDown(event INSEvent)
	// Informs the gesture recognizer that the user moved the mouse with a button other than the left or right one pressed.
	OtherMouseDragged(event INSEvent)
	// Informs the gesture recognizer that the user released a mouse button other than the left or right one.
	OtherMouseUp(event INSEvent)
	// Informs the gesture recognizer that the user pressed the right mouse button.
	RightMouseDown(event INSEvent)
	// Informs the gesture recognizer that the user moved the mouse with the right button pressed.
	RightMouseDragged(event INSEvent)
	// Informs the gesture recognizer that the user released the right mouse button.
	RightMouseUp(event INSEvent)
	// Informs the gesture recognizer that the user is performing a pinch gesture.
	MagnifyWithEvent(event INSEvent)
	// Informs the gesture recognizer that the user is performing a rotation gesture.
	RotateWithEvent(event INSEvent)
	// Overridden to indicate that the specified gesture recognizer can prevent the current object from recognizing a gesture.
	CanBePreventedByGestureRecognizer(preventingGestureRecognizer INSGestureRecognizer) bool
	// Overridden to indicate that the current object can prevent the specified gesture recognizer from recognizing its gesture.
	CanPreventGestureRecognizer(preventedGestureRecognizer INSGestureRecognizer) bool
	// Overridden to indicate that the current object must fail before the specified gesture recognizer begins recognizing its gesture.
	ShouldBeRequiredToFailByGestureRecognizer(otherGestureRecognizer INSGestureRecognizer) bool
	// Overridden to indicate that the specified gesture recognizer must fail before the current object begins recognizing its gesture.
	ShouldRequireFailureOfGestureRecognizer(otherGestureRecognizer INSGestureRecognizer) bool
	// Informs the gesture recognizer that the user has pressed a key.
	KeyDown(event INSEvent)
	// Informs the gesture recognizer that the user released a key.
	KeyUp(event INSEvent)
	// Informs the user that a tablet-point event occurred.
	TabletPoint(event INSEvent)
	// Informs the current object that the user pressed or released a modifier key (Shift, Control, and so on).
	FlagsChanged(event INSEvent)
	// Informs the current object that a pressure change occurred on a system that supports pressure sensitivity.
	PressureChangeWithEvent(event INSEvent)

	// Topic: Configuring Pressure

	// Configures the behavior and progression of the Force Touch trackpad when responding to recognized pressure gestures.
	PressureConfiguration() INSPressureConfiguration
	SetPressureConfiguration(value INSPressureConfiguration)

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) NSGestureRecognizer

	// Topic: Instance Properties

	AllowedTouchTypes() NSTouchTypeMask
	SetAllowedTouchTypes(value NSTouchTypeMask)
	ModifierFlags() NSEventModifierFlags
	Name() string
	SetName(value string)

	// Topic: Instance Methods

	// Called when one or more fingers first make contact with an [NSTouchBar](<doc://com.apple.appkit/documentation/AppKit/NSTouchBar>) instance on the Touch Bar.
	TouchesBeganWithEvent(event INSEvent)
	// Called when a system event, such as a low-memory warning, cancels an in-progress touch event in an [NSTouchBar](<doc://com.apple.appkit/documentation/AppKit/NSTouchBar>) object.
	TouchesCancelledWithEvent(event INSEvent)
	// Called when one or more fingers are removed from contact with an [NSTouchBar](<doc://com.apple.appkit/documentation/AppKit/NSTouchBar>) instance on the Touch Bar.
	TouchesEndedWithEvent(event INSEvent)
	// Called when one or more fingers, associated with an in-progress event, move within an [NSTouchBar](<doc://com.apple.appkit/documentation/AppKit/NSTouchBar>) instance on the Touch Bar.
	TouchesMovedWithEvent(event INSEvent)
	MouseCancelled(event INSEvent)

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (g NSGestureRecognizer) Init() NSGestureRecognizer {
	rv := objc.Send[NSGestureRecognizer](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g NSGestureRecognizer) Autorelease() NSGestureRecognizer {
	rv := objc.Send[NSGestureRecognizer](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSGestureRecognizer creates a new NSGestureRecognizer instance.
func NewNSGestureRecognizer() NSGestureRecognizer {
	class := getNSGestureRecognizerClass()
	rv := objc.Send[NSGestureRecognizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/init(coder:)
func NewGestureRecognizerWithCoder(coder foundation.INSCoder) NSGestureRecognizer {
	instance := getNSGestureRecognizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSGestureRecognizerFromID(rv)
}


// Initializes the gesture recognizer with the specified target and action
// information.
//
// target: The object whose action method is called when the gesture is recognized.
// You must not specify `nil` for this parameter.
//
// action: A selector that identifies the method to call when the gesture is
// recognized. This method must be implemented by the object in `target`. You
// must not specify `nil` for this parameter.
//
// # Return Value
// 
// The initialized gesture recognizer object or `nil` if an error occurred.
//
// # Discussion
// 
// This method is the designated initializer. Subclasses must call this method
// from their own custom initialization methods. Call the method before
// performing other tasks.
// 
// This method records the specified `target` and `action` values and prepares
// the gesture recognizer for use.
// 
// The `action` method must have one of the following signatures:
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/init(target:action:)
func NewGestureRecognizerWithTargetAction(target objectivec.IObject, action objc.SEL) NSGestureRecognizer {
	instance := getNSGestureRecognizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTarget:action:"), target, action)
	return NSGestureRecognizerFromID(rv)
}







// Initializes the gesture recognizer with the specified target and action
// information.
//
// target: The object whose action method is called when the gesture is recognized.
// You must not specify `nil` for this parameter.
//
// action: A selector that identifies the method to call when the gesture is
// recognized. This method must be implemented by the object in `target`. You
// must not specify `nil` for this parameter.
//
// # Return Value
// 
// The initialized gesture recognizer object or `nil` if an error occurred.
//
// # Discussion
// 
// This method is the designated initializer. Subclasses must call this method
// from their own custom initialization methods. Call the method before
// performing other tasks.
// 
// This method records the specified `target` and `action` values and prepares
// the gesture recognizer for use.
// 
// The `action` method must have one of the following signatures:
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/init(target:action:)
func (g NSGestureRecognizer) InitWithTargetAction(target objectivec.IObject, action objc.SEL) NSGestureRecognizer {
	rv := objc.Send[NSGestureRecognizer](g.ID, objc.Sel("initWithTarget:action:"), target, action)
	return rv
}

// Returns the point computed as the location of the gesture.
//
// view: The view whose coordinate system you want to use for determining the
// location of the gesture. Specify `nil` to return the point in the
// coordinate system of the window.
//
// # Return Value
// 
// The point at which the gesture occurred. The returned point is in the
// coordinate system of the specified view, or in the coordinate system of the
// window if you specified `nil` for the view parameter.
//
// # Discussion
// 
// Use this method to determine the location at which the gesture occurred.
// Subclasses are responsible for overriding this method and returning an
// appropriate value based on the type of gesture.
// 
// For specific information about what the returned point represents, see the
// specific gesture recognizer subclass.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/location(in:)
func (g NSGestureRecognizer) LocationInView(view INSView) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](g.ID, objc.Sel("locationInView:"), view)
	return corefoundation.CGPoint(rv)
}

// Overridden to reset the internal state of the gesture recognizer when an
// attempt completes.
//
// # Discussion
// 
// AppKit calls this method after the gesture recognizer state has been set to
// any of the terminal states: [NSGestureRecognizer.State.ended],
// [NSGestureRecognizer.State.cancelled], [NSGestureRecognizer.State.failed],
// or [recognized]. Subclasses should override this method and use it to reset
// any internal state of the gesture recognizer in preparation for a new
// recognition attempt. After this method is called, the gesture recognizer
// receives no further updates for events that began but have not yet ended.
//
// [NSGestureRecognizer.State.cancelled]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/State-swift.enum/cancelled
// [NSGestureRecognizer.State.ended]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/State-swift.enum/ended
// [NSGestureRecognizer.State.failed]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/State-swift.enum/failed
// [recognized]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/State-swift.enum/recognized
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/reset()
func (g NSGestureRecognizer) Reset() {
	objc.Send[objc.ID](g.ID, objc.Sel("reset"))
}

// Informs the gesture recognizer that the user pressed the left mouse button.
//
// event: An object encapsulating information about the mouse-down event.
//
// # Discussion
// 
// The default implementation of this method does nothing. Use this method to
// start tracking the gesture in whatever way is appropriate.
// 
// A gesture recognizer monitors events that occur in its view (and any
// subviews) but does not take part in the responder chain itself. The gesture
// recognizer receives events before any views do. Use the
// [DelaysPrimaryMouseButtonEvents] property to control whether `event` is
// propagated to the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/mouseDown(with:)
func (g NSGestureRecognizer) MouseDown(event INSEvent) {
	objc.Send[objc.ID](g.ID, objc.Sel("mouseDown:"), event)
}

// Informs the gesture recognizer that the user moved the mouse with the left
// button pressed.
//
// event: An object encapsulating information about the mouse-dragged event.
//
// # Discussion
// 
// The default implementation of this method does nothing. Use this method to
// update the state of your gesture recognizer in whatever way is appropriate.
// 
// A gesture recognizer monitors events that occur in its view (and any
// subviews) but does not take part in the responder chain itself. The gesture
// recognizer receives events before any views do. Use the
// [DelaysPrimaryMouseButtonEvents] property to control whether `event` is
// propagated to the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/mouseDragged(with:)
func (g NSGestureRecognizer) MouseDragged(event INSEvent) {
	objc.Send[objc.ID](g.ID, objc.Sel("mouseDragged:"), event)
}

// Informs the gesture recognizer that the user released the left mouse
// button.
//
// event: An object encapsulating information about the mouse-up event.
//
// # Discussion
// 
// The default implementation of this method does nothing. Use this method to
// update the state of your gesture recognizer in whatever way is appropriate.
// 
// A gesture recognizer monitors events that occur in its view (and any
// subviews) but does not take part in the responder chain itself. The gesture
// recognizer receives events before any views do. Use the
// [DelaysPrimaryMouseButtonEvents] property to control whether `event` is
// propagated to the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/mouseUp(with:)
func (g NSGestureRecognizer) MouseUp(event INSEvent) {
	objc.Send[objc.ID](g.ID, objc.Sel("mouseUp:"), event)
}

// Informs the gesture recognizer that the user pressed a mouse button other
// than the left or right one.
//
// event: An object encapsulating information about the mouse-down event.
//
// # Discussion
// 
// The default implementation of this method does nothing. Use this method to
// start tracking the gesture in whatever way is appropriate.
// 
// A gesture recognizer monitors events that occur in its view (and any
// subviews) but does not take part in the responder chain itself. The gesture
// recognizer receives events before any views do. Use the
// [DelaysOtherMouseButtonEvents] property to control whether `event` is
// propagated to the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/otherMouseDown(with:)
func (g NSGestureRecognizer) OtherMouseDown(event INSEvent) {
	objc.Send[objc.ID](g.ID, objc.Sel("otherMouseDown:"), event)
}

// Informs the gesture recognizer that the user moved the mouse with a button
// other than the left or right one pressed.
//
// event: An object encapsulating information about the mouse-dragged event.
//
// # Discussion
// 
// The default implementation of this method does nothing. Use this method to
// update the state of your gesture recognizer in whatever way is appropriate.
// 
// A gesture recognizer monitors events that occur in its view (and any
// subviews) but does not take part in the responder chain itself. The gesture
// recognizer receives events before any views do. Use the
// [DelaysOtherMouseButtonEvents] property to control whether `event` is
// propagated to the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/otherMouseDragged(with:)
func (g NSGestureRecognizer) OtherMouseDragged(event INSEvent) {
	objc.Send[objc.ID](g.ID, objc.Sel("otherMouseDragged:"), event)
}

// Informs the gesture recognizer that the user released a mouse button other
// than the left or right one.
//
// event: An object encapsulating information about the mouse-up event.
//
// # Discussion
// 
// The default implementation of this method does nothing. Use this method to
// update the state of your gesture recognizer in whatever way is appropriate.
// 
// A gesture recognizer monitors events that occur in its view (and any
// subviews) but does not take part in the responder chain itself. The gesture
// recognizer receives events before any views do. Use the
// [DelaysOtherMouseButtonEvents] property to control whether `event` is
// propagated to the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/otherMouseUp(with:)
func (g NSGestureRecognizer) OtherMouseUp(event INSEvent) {
	objc.Send[objc.ID](g.ID, objc.Sel("otherMouseUp:"), event)
}

// Informs the gesture recognizer that the user pressed the right mouse
// button.
//
// event: An object encapsulating information about the mouse-down event.
//
// # Discussion
// 
// The default implementation of this method does nothing. Use this method to
// start tracking the gesture in whatever way is appropriate.
// 
// A gesture recognizer monitors events that occur in its view (and any
// subviews) but does not take part in the responder chain itself. The gesture
// recognizer receives events before any views do. Use the
// [DelaysSecondaryMouseButtonEvents] property to control whether `event` is
// propagated to the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/rightMouseDown(with:)
func (g NSGestureRecognizer) RightMouseDown(event INSEvent) {
	objc.Send[objc.ID](g.ID, objc.Sel("rightMouseDown:"), event)
}

// Informs the gesture recognizer that the user moved the mouse with the right
// button pressed.
//
// event: An object encapsulating information about the mouse-dragged event.
//
// # Discussion
// 
// The default implementation of this method does nothing. Use this method to
// update the state of your gesture recognizer in whatever way is appropriate.
// 
// A gesture recognizer monitors events that occur in its view (and any
// subviews) but does not take part in the responder chain itself. The gesture
// recognizer receives events before any views do. Use the
// [DelaysSecondaryMouseButtonEvents] property to control whether `event` is
// propagated to the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/rightMouseDragged(with:)
func (g NSGestureRecognizer) RightMouseDragged(event INSEvent) {
	objc.Send[objc.ID](g.ID, objc.Sel("rightMouseDragged:"), event)
}

// Informs the gesture recognizer that the user released the right mouse
// button.
//
// event: An object encapsulating information about the mouse-up event.
//
// # Discussion
// 
// The default implementation of this method does nothing. Use this method to
// update the state of your gesture recognizer in whatever way is appropriate.
// 
// A gesture recognizer monitors events that occur in its view (and any
// subviews) but does not take part in the responder chain itself. The gesture
// recognizer receives events before any views do. Use the
// [DelaysSecondaryMouseButtonEvents] property to control whether `event` is
// propagated to the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/rightMouseUp(with:)
func (g NSGestureRecognizer) RightMouseUp(event INSEvent) {
	objc.Send[objc.ID](g.ID, objc.Sel("rightMouseUp:"), event)
}

// Informs the gesture recognizer that the user is performing a pinch gesture.
//
// event: An object encapsulating information about the magnify event.
//
// # Discussion
// 
// The default implementation of this method does nothing. Use this method to
// start tracking the gesture in whatever way is appropriate.
// 
// A gesture recognizer monitors events that occur in its view (and any
// subviews) but does not take part in the responder chain itself. The gesture
// recognizer receives events before any views do. Use the
// [DelaysMagnificationEvents] property to control whether `event` is
// propagated to the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/magnify(with:)
func (g NSGestureRecognizer) MagnifyWithEvent(event INSEvent) {
	objc.Send[objc.ID](g.ID, objc.Sel("magnifyWithEvent:"), event)
}

// Informs the gesture recognizer that the user is performing a rotation
// gesture.
//
// event: An object encapsulating information about the rotate event.
//
// # Discussion
// 
// The default implementation of this method does nothing. Use this method to
// start tracking the gesture in whatever way is appropriate.
// 
// A gesture recognizer monitors events that occur in its view (and any
// subviews) but does not take part in the responder chain itself. The gesture
// recognizer receives events before any views do. Use the
// [DelaysRotationEvents] property to control whether `event` is propagated to
// the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/rotate(with:)
func (g NSGestureRecognizer) RotateWithEvent(event INSEvent) {
	objc.Send[objc.ID](g.ID, objc.Sel("rotateWithEvent:"), event)
}

// Overridden to indicate that the specified gesture recognizer can prevent
// the current object from recognizing a gesture.
//
// preventingGestureRecognizer: The gesture recognizer that can prevent the current object from recognizing
// its gesture.
//
// # Return Value
// 
// [true] to indicate that `preventingGestureRecognizer` can block the current
// gesture recognizer from recognizing its gesture, or [false] if both gesture
// recognizers can operate simultaneously.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method enables similar behavior as the [GestureRecognizerShouldBegin]
// and [GestureRecognizerShouldRequireFailureOfGestureRecognizer] methods of
// the gesture recognizer’s delegate. Using this method lets you define
// rules that apply to all instances of your custom gesture recognizer class.
// For example, a gesture recognizer of a given class might want to prevent
// instances of the same class from recognizing at the same time.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/canBePrevented(by:)
func (g NSGestureRecognizer) CanBePreventedByGestureRecognizer(preventingGestureRecognizer INSGestureRecognizer) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("canBePreventedByGestureRecognizer:"), preventingGestureRecognizer)
	return rv
}

// Overridden to indicate that the current object can prevent the specified
// gesture recognizer from recognizing its gesture.
//
// preventedGestureRecognizer: The gesture recognizer to be prevented from recognizing its gesture.
//
// # Return Value
// 
// [true] to indicate that `preventedGestureRecognizer` should be blocked from
// recognizing its gesture, or [false] if both gesture recognizers can operate
// simultaneously.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method enables similar behavior as the [GestureRecognizerShouldBegin]
// and [GestureRecognizerShouldRequireFailureOfGestureRecognizer] methods of
// the gesture recognizer’s delegate. Using this method lets you define
// rules that apply to all instances of your custom gesture recognizer class.
// For example, an [NSClickGestureRecognizer] object does not prevent another
// [NSClickGestureRecognizer] object with a higher click count from
// recognizing its gesture.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/canPrevent(_:)
func (g NSGestureRecognizer) CanPreventGestureRecognizer(preventedGestureRecognizer INSGestureRecognizer) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("canPreventGestureRecognizer:"), preventedGestureRecognizer)
	return rv
}

// Overridden to indicate that the current object must fail before the
// specified gesture recognizer begins recognizing its gesture.
//
// otherGestureRecognizer: The gesture recognizer that is not allowed to recognize its gesture until
// the current object fails.
//
// # Return Value
// 
// [true] to cause `otherGestureRecognizer` to wait until the current object
// fails before beginning its own gesture recognition.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Using this method lets you define rules that apply to all instances of your
// custom gesture recognizer class.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/shouldBeRequiredToFail(by:)
func (g NSGestureRecognizer) ShouldBeRequiredToFailByGestureRecognizer(otherGestureRecognizer INSGestureRecognizer) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("shouldBeRequiredToFailByGestureRecognizer:"), otherGestureRecognizer)
	return rv
}

// Overridden to indicate that the specified gesture recognizer must fail
// before the current object begins recognizing its gesture.
//
// otherGestureRecognizer: The gesture recognizer that must fail before the current object is allowed
// to recognize its gesture.
//
// # Return Value
// 
// [true] to cause the current object to wait to recognize its own gesture
// until the object in otherGestureRecognizer fails.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Using this method lets you define rules that apply to all instances of your
// custom gesture recognizer class.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/shouldRequireFailure(of:)
func (g NSGestureRecognizer) ShouldRequireFailureOfGestureRecognizer(otherGestureRecognizer INSGestureRecognizer) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("shouldRequireFailureOfGestureRecognizer:"), otherGestureRecognizer)
	return rv
}

// Informs the gesture recognizer that the user has pressed a key.
//
// event: An object encapsulating information about the key-down event.
//
// # Discussion
// 
// The default implementation of this method does nothing. Use this method to
// update the state of your gesture recognizer in whatever way is appropriate.
// 
// A gesture recognizer monitors events that occur in its view (and any
// subviews) but does not take part in the responder chain itself. The gesture
// recognizer receives events before any views do. Use the [DelaysKeyEvents]
// property to control whether `event` is propagated to the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/keyDown(with:)
func (g NSGestureRecognizer) KeyDown(event INSEvent) {
	objc.Send[objc.ID](g.ID, objc.Sel("keyDown:"), event)
}

// Informs the gesture recognizer that the user released a key.
//
// event: An object encapsulating information about the key-up event.
//
// # Discussion
// 
// The default implementation of this method does nothing. Use this method to
// update the state of your gesture recognizer in whatever way is appropriate.
// 
// A gesture recognizer monitors events that occur in its view (and any
// subviews) but does not take part in the responder chain itself. The gesture
// recognizer receives events before any views do. Use the [DelaysKeyEvents]
// property to control whether `event` is propagated to the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/keyUp(with:)
func (g NSGestureRecognizer) KeyUp(event INSEvent) {
	objc.Send[objc.ID](g.ID, objc.Sel("keyUp:"), event)
}

// Informs the user that a tablet-point event occurred.
//
// event: An object encapsulating information about the tablet-point event.
// Tablet-point events describe the current state of a pointing device that is
// in proximity to its tablet, reflecting changes such as location, pressure,
// tilt, and rotation. For more information, see [NSEvent].
//
// # Discussion
// 
// The default implementation of this method does nothing. Use this method to
// update the state of your gesture recognizer in whatever way is appropriate.
// 
// A gesture recognizer monitors events that occur in its view (and any
// subviews) but does not take part in the responder chain itself. The gesture
// recognizer receives events before any views do.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/tabletPoint(with:)
func (g NSGestureRecognizer) TabletPoint(event INSEvent) {
	objc.Send[objc.ID](g.ID, objc.Sel("tabletPoint:"), event)
}

// Informs the current object that the user pressed or released a modifier key
// (Shift, Control, and so on).
//
// event: An object encapsulating information about the modifier-key event.
//
// # Discussion
// 
// The default implementation of this method does nothing. Use this method to
// update the state of your gesture recognizer in whatever way is appropriate.
// 
// A gesture recognizer monitors events that occur in its view (and any
// subviews) but does not take part in the responder chain itself. The gesture
// recognizer receives events before any views do.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/flagsChanged(with:)
func (g NSGestureRecognizer) FlagsChanged(event INSEvent) {
	objc.Send[objc.ID](g.ID, objc.Sel("flagsChanged:"), event)
}

// Informs the current object that a pressure change occurred on a system that
// supports pressure sensitivity.
//
// event: An [NSEvent] object encapsulating information about the event that invoked
// the change in pressure.
//
// # Discussion
// 
// This method is invoked automatically in response to user actions. `event`
// is the event that initiated the change in pressure.
// 
// The default implementation of this method does nothing. Use this method to
// update the state of your gesture recognizer in whatever way is appropriate.
// 
// A gesture recognizer monitors events that occur in its view (and any
// subviews) but does not take part in the responder chain itself. The gesture
// recognizer receives events before any views do.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/pressureChange(with:)
func (g NSGestureRecognizer) PressureChangeWithEvent(event INSEvent) {
	objc.Send[objc.ID](g.ID, objc.Sel("pressureChangeWithEvent:"), event)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/init(coder:)
func (g NSGestureRecognizer) InitWithCoder(coder foundation.INSCoder) NSGestureRecognizer {
	rv := objc.Send[NSGestureRecognizer](g.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Called when one or more fingers first make contact with an [NSTouchBar]
// instance on the Touch Bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/touchesBegan(with:)
func (g NSGestureRecognizer) TouchesBeganWithEvent(event INSEvent) {
	objc.Send[objc.ID](g.ID, objc.Sel("touchesBeganWithEvent:"), event)
}

// Called when a system event, such as a low-memory warning, cancels an
// in-progress touch event in an [NSTouchBar] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/touchesCancelled(with:)
func (g NSGestureRecognizer) TouchesCancelledWithEvent(event INSEvent) {
	objc.Send[objc.ID](g.ID, objc.Sel("touchesCancelledWithEvent:"), event)
}

// Called when one or more fingers are removed from contact with an
// [NSTouchBar] instance on the Touch Bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/touchesEnded(with:)
func (g NSGestureRecognizer) TouchesEndedWithEvent(event INSEvent) {
	objc.Send[objc.ID](g.ID, objc.Sel("touchesEndedWithEvent:"), event)
}

// Called when one or more fingers, associated with an in-progress event, move
// within an [NSTouchBar] instance on the Touch Bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/touchesMoved(with:)
func (g NSGestureRecognizer) TouchesMovedWithEvent(event INSEvent) {
	objc.Send[objc.ID](g.ID, objc.Sel("touchesMovedWithEvent:"), event)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/mouseCancelled(with:)
func (g NSGestureRecognizer) MouseCancelled(event INSEvent) {
	objc.Send[objc.ID](g.ID, objc.Sel("mouseCancelled:"), event)
}
func (g NSGestureRecognizer) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The action method to call when the gesture is recognized.
//
// # Discussion
// 
// You specify this method when initializing your gesture recognizer but can
// also change the method later. The gesture recognizer executes your action
// method during specific states of the gesture recognition process. Use your
// action method to perform any app-specific tasks. For example, you might use
// that method to animate changes onscreen while the gesture is in progress.
// 
// The signature of this method must be one of the following:
// 
// For continuous gestures, it is recommended that you use an action method
// that accepts the gesture recognizer as a parameter. In your method, use the
// provided gesture recognizer object to get the current state of the gesture
// and perform appropriate tasks based on that state.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/action
func (g NSGestureRecognizer) Action() objc.SEL {
	rv := objc.Send[objc.SEL](g.ID, objc.Sel("action"))
	return rv
}
func (g NSGestureRecognizer) SetAction(value objc.SEL) {
	objc.Send[struct{}](g.ID, objc.Sel("setAction:"), value)
}



// The object that implements the action method.
//
// # Discussion
// 
// The object in this property must implement the method specified by the
// [Action] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/target
func (g NSGestureRecognizer) Target() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("target"))
	return objectivec.Object{ID: rv}
}
func (g NSGestureRecognizer) SetTarget(value objectivec.IObject) {
	objc.Send[struct{}](g.ID, objc.Sel("setTarget:"), value)
}



// The current state of the gesture recognizer.
//
// # Discussion
// 
// This property conveys where the gesture recognizer is in the recognition
// process. The default declaration of this property is read-only so that
// external clients (such as other gesture recognizers) can use the value for
// informational purposes. Subclasses can redeclare the property as read-write
// internally. When doing so, you do not need to provide a custom
// implementation to set the value of the property. This class provides an
// implementation that detects state transitions and updates the gesture
// recognizer accordingly.
// 
// For more information about the state transitions that can occur in a
// gesture recognizer, see [NSGestureRecognizer].
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/state-swift.property
func (g NSGestureRecognizer) State() NSGestureRecognizerState {
	rv := objc.Send[NSGestureRecognizerState](g.ID, objc.Sel("state"))
	return NSGestureRecognizerState(rv)
}



// The view to which the gesture recognizer is attached.
//
// # Discussion
// 
// To attach a gesture recognizer to a view, call the [AddGestureRecognizer]
// method of the view. If the gesture recognizer is not attached to a view,
// the value in this property is `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/view
func (g NSGestureRecognizer) View() INSView {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("view"))
	return NSViewFromID(objc.ID(rv))
}



// A Boolean value indicating whether the gesture recognizer is able to handle
// events.
//
// # Discussion
// 
// When the value of this property is [true], the gesture recognizer receives
// events and uses them to determine when its gesture is performed. When the
// value is [false], the gesture recognizer does not receive events. Changing
// the value from [true] to [false] while the gesture recognizer is in the
// process of recognizing a gesture changes the state of the gesture
// recognizer to [NSGestureRecognizer.State.cancelled].
// 
// The default value of this property is [true].
//
// [NSGestureRecognizer.State.cancelled]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/State-swift.enum/cancelled
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/isEnabled
func (g NSGestureRecognizer) Enabled() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isEnabled"))
	return rv
}
func (g NSGestureRecognizer) SetEnabled(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setEnabled:"), value)
}



// A Boolean value that indicates whether primary mouse button events are
// delivered only after gesture recognition fails.
//
// # Discussion
// 
// When the value of this property is [true], primary mouse button events are
// delivered to the target view only after gesture recognition fails. Set this
// property to [true] to prevent the view from processing events that might be
// recognized as part of a gesture. Once gesture recognition begins, all types
// of events are delayed until gesture recognition fails.
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/delaysPrimaryMouseButtonEvents
func (g NSGestureRecognizer) DelaysPrimaryMouseButtonEvents() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("delaysPrimaryMouseButtonEvents"))
	return rv
}
func (g NSGestureRecognizer) SetDelaysPrimaryMouseButtonEvents(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setDelaysPrimaryMouseButtonEvents:"), value)
}



// A Boolean value that indicates whether secondary mouse button events are
// delivered only after gesture recognition fails.
//
// # Discussion
// 
// When the value of this property is [true], secondary mouse button events
// are delivered to the target view only after gesture recognition fails. Set
// this property to [true] to prevent the view from processing events that
// might be recognized as part of a gesture. Once gesture recognition begins,
// all types of events are delayed until gesture recognition fails.
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/delaysSecondaryMouseButtonEvents
func (g NSGestureRecognizer) DelaysSecondaryMouseButtonEvents() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("delaysSecondaryMouseButtonEvents"))
	return rv
}
func (g NSGestureRecognizer) SetDelaysSecondaryMouseButtonEvents(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setDelaysSecondaryMouseButtonEvents:"), value)
}



// A Boolean value that indicates whether other mouse button events are
// delivered only after gesture recognition fails.
//
// # Discussion
// 
// When the value of this property is [true], other mouse button events are
// delivered to the target view only after gesture recognition fails. Set this
// property to [true] to prevent the view from processing events that might be
// recognized as part of a gesture. Once gesture recognition begins, all types
// of events are delayed until gesture recognition fails.
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/delaysOtherMouseButtonEvents
func (g NSGestureRecognizer) DelaysOtherMouseButtonEvents() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("delaysOtherMouseButtonEvents"))
	return rv
}
func (g NSGestureRecognizer) SetDelaysOtherMouseButtonEvents(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setDelaysOtherMouseButtonEvents:"), value)
}



// A Boolean value that indicates whether key events are delivered only after
// gesture recognition fails.
//
// # Discussion
// 
// When the value of this property is [true], key events are delivered to the
// target view only after gesture recognition fails. Set this property to
// [true] to prevent the view from processing events that might be recognized
// as part of a gesture. Once gesture recognition begins, all types of events
// are delayed until gesture recognition fails.
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/delaysKeyEvents
func (g NSGestureRecognizer) DelaysKeyEvents() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("delaysKeyEvents"))
	return rv
}
func (g NSGestureRecognizer) SetDelaysKeyEvents(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setDelaysKeyEvents:"), value)
}



// A Boolean value that indicates whether magnification events are delivered
// only after gesture recognition fails.
//
// # Discussion
// 
// When the value of this property is [true], magnification events are
// delivered to the target view only after gesture recognition fails. Set this
// property to [true] to prevent the view from processing events that might be
// recognized as part of a gesture. Once gesture recognition begins, all types
// of events are delayed until gesture recognition fails.
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/delaysMagnificationEvents
func (g NSGestureRecognizer) DelaysMagnificationEvents() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("delaysMagnificationEvents"))
	return rv
}
func (g NSGestureRecognizer) SetDelaysMagnificationEvents(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setDelaysMagnificationEvents:"), value)
}



// A Boolean value that indicates whether rotation events are delivered only
// after gesture recognition fails.
//
// # Discussion
// 
// When the value of this property is [true], rotation events are delivered to
// the target view only after gesture recognition fails. Set this property to
// [true] to prevent the view from processing events that might be recognized
// as part of a gesture. Once gesture recognition begins, all types of events
// are delayed until gesture recognition fails.
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/delaysRotationEvents
func (g NSGestureRecognizer) DelaysRotationEvents() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("delaysRotationEvents"))
	return rv
}
func (g NSGestureRecognizer) SetDelaysRotationEvents(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setDelaysRotationEvents:"), value)
}



// The delegate of the gesture recognizer.
//
// # Discussion
// 
// Use the delegate for fine-grained control over the recognition of a
// gesture. For example, you can use the delegate to determine whether gesture
// recognition should begin or whether it should start only after other
// gesture recognizers fail.
// 
// The delegate must implement the [NSGestureRecognizerDelegate] protocol.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/delegate
func (g NSGestureRecognizer) Delegate() NSGestureRecognizerDelegate {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("delegate"))
	return NSGestureRecognizerDelegateObjectFromID(rv)
}
func (g NSGestureRecognizer) SetDelegate(value NSGestureRecognizerDelegate) {
	objc.Send[struct{}](g.ID, objc.Sel("setDelegate:"), value)
}



// Configures the behavior and progression of the Force Touch trackpad when
// responding to recognized pressure gestures.
//
// # Discussion
// 
// This property contains a value of type [NSPressureConfiguration], which
// configures the behavior and progression of the Force Touch trackpad when
// responding to recognized pressure gestures.
// 
// Ideally, you should avoid changing the pressure configuration during
// recognition, as the gesture may complete before the configuration has time
// to take effect. If you do need to change the pressure configuration during
// recognition, call the [Set] method of [PressureConfiguration].
// 
// Once the gesture ends or if recognition fails, the pressure configuration
// resets to the current view’s pressure configuration, if any, for the
// remaining duration of the gesture.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/pressureConfiguration
func (g NSGestureRecognizer) PressureConfiguration() INSPressureConfiguration {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("pressureConfiguration"))
	return NSPressureConfigurationFromID(objc.ID(rv))
}
func (g NSGestureRecognizer) SetPressureConfiguration(value INSPressureConfiguration) {
	objc.Send[struct{}](g.ID, objc.Sel("setPressureConfiguration:"), value)
}



// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/allowedTouchTypes
func (g NSGestureRecognizer) AllowedTouchTypes() NSTouchTypeMask {
	rv := objc.Send[NSTouchTypeMask](g.ID, objc.Sel("allowedTouchTypes"))
	return NSTouchTypeMask(rv)
}
func (g NSGestureRecognizer) SetAllowedTouchTypes(value NSTouchTypeMask) {
	objc.Send[struct{}](g.ID, objc.Sel("setAllowedTouchTypes:"), value)
}



// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/modifierFlags
func (g NSGestureRecognizer) ModifierFlags() NSEventModifierFlags {
	rv := objc.Send[NSEventModifierFlags](g.ID, objc.Sel("modifierFlags"))
	return NSEventModifierFlags(rv)
}



// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/name
func (g NSGestureRecognizer) Name() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (g NSGestureRecognizer) SetName(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setName:"), objc.String(value))
}

























