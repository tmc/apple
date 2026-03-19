// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSWritingToolsCoordinator] class.
var (
	_NSWritingToolsCoordinatorClass     NSWritingToolsCoordinatorClass
	_NSWritingToolsCoordinatorClassOnce sync.Once
)

func getNSWritingToolsCoordinatorClass() NSWritingToolsCoordinatorClass {
	_NSWritingToolsCoordinatorClassOnce.Do(func() {
		_NSWritingToolsCoordinatorClass = NSWritingToolsCoordinatorClass{class: objc.GetClass("NSWritingToolsCoordinator")}
	})
	return _NSWritingToolsCoordinatorClass
}

// GetNSWritingToolsCoordinatorClass returns the class object for NSWritingToolsCoordinator.
func GetNSWritingToolsCoordinatorClass() NSWritingToolsCoordinatorClass {
	return getNSWritingToolsCoordinatorClass()
}

type NSWritingToolsCoordinatorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSWritingToolsCoordinatorClass) Alloc() NSWritingToolsCoordinator {
	rv := objc.Send[NSWritingToolsCoordinator](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that manages interactions between Writing Tools and your custom
// text view.
//
// # Overview
// 
// Add a [NSWritingToolsCoordinator] object to a custom view when you want to
// add Writing Tools support to that view. The coordinator manages
// interactions between your view and the Writing Tools UI and back-end
// capabilities. When creating a coordinator, you supply a delegate object to
// respond to requests from the system and provide needed information. Your
// delegate delivers your view’s text to Writing Tools, incorporates
// suggested changes back into your text storage, and supports the animations
// that Writing Tools creates to show the state of an operation.
// 
// Create the [NSWritingToolsCoordinator] object when setting up your UI, and
// initialize it with a custom object that adopts the
// [NSWritingToolsCoordinatorDelegate] protocol. Add the coordinator to the
// [WritingToolsCoordinator] property of your view. When a coordinator is
// present on a view, the system adds UI elements to initiate Writing Tools
// operations.
// 
// When defining the delegate, choose an object from your app that has access
// to your view and its text storage. You can adopt the
// [NSWritingToolsCoordinatorDelegate] protocol in the view itself, or in
// another type that your view uses to manage content. During the interactions
// with Writing Tools, the delegate gets and sets the contents of the view’s
// text storage and supports Writing Tools behaviors.
//
// # Creating a coordinator object
//
//   - [NSWritingToolsCoordinator.InitWithDelegate]: Creates a writing tools coordinator and assigns the specified delegate object to it.
//
// # Managing Writing Tools interactions
//
//   - [NSWritingToolsCoordinator.Delegate]: The object that handles Writing Tools interactions for your view.
//   - [NSWritingToolsCoordinator.View]: The view that currently uses the writing tools coordinator.
//
// # Getting the host views for effects
//
//   - [NSWritingToolsCoordinator.EffectContainerView]: The view that Writing Tools uses to display visual effects during the text-rewriting process.
//   - [NSWritingToolsCoordinator.SetEffectContainerView]
//   - [NSWritingToolsCoordinator.DecorationContainerView]: The view that Writing Tools uses to display background decorations such as proofreading marks.
//   - [NSWritingToolsCoordinator.SetDecorationContainerView]
//
// # Configuring the experience
//
//   - [NSWritingToolsCoordinator.PreferredBehavior]: The level of Writing Tools support you want the system to provide for your view.
//   - [NSWritingToolsCoordinator.SetPreferredBehavior]
//   - [NSWritingToolsCoordinator.Behavior]: The actual level of Writing Tools support the system provides for your view.
//   - [NSWritingToolsCoordinator.PreferredResultOptions]: The type of content you allow Writing Tools to generate for your custom text view.
//   - [NSWritingToolsCoordinator.SetPreferredResultOptions]
//   - [NSWritingToolsCoordinator.ResultOptions]: The type of content the system generates for your custom text view.
//
// # Reporting changes to Writing Tools
//
//   - [NSWritingToolsCoordinator.UpdateRangeWithTextReasonForContextWithIdentifier]: Informs the coordinator about changes your app made to the text in the specified context object.
//   - [NSWritingToolsCoordinator.UpdateForReflowedTextInContextWithIdentifier]: Informs the coordinator that a change occurred to the view or its text that requires a layout update.
//
// # Managing the current state
//
//   - [NSWritingToolsCoordinator.StopWritingTools]: Stops the current Writing Tools operation and dismisses the system UI.
//   - [NSWritingToolsCoordinator.State]: The current level of Writing Tools activity in your view.
//
// # Instance Properties
//
//   - [NSWritingToolsCoordinator.IncludesTextListMarkers]
//   - [NSWritingToolsCoordinator.SetIncludesTextListMarkers]
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator
type NSWritingToolsCoordinator struct {
	objectivec.Object
}

// NSWritingToolsCoordinatorFromID constructs a [NSWritingToolsCoordinator] from an objc.ID.
//
// An object that manages interactions between Writing Tools and your custom
// text view.
func NSWritingToolsCoordinatorFromID(id objc.ID) NSWritingToolsCoordinator {
	return NSWritingToolsCoordinator{objectivec.Object{ID: id}}
}
// NOTE: NSWritingToolsCoordinator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSWritingToolsCoordinator] class.
//
// # Creating a coordinator object
//
//   - [INSWritingToolsCoordinator.InitWithDelegate]: Creates a writing tools coordinator and assigns the specified delegate object to it.
//
// # Managing Writing Tools interactions
//
//   - [INSWritingToolsCoordinator.Delegate]: The object that handles Writing Tools interactions for your view.
//   - [INSWritingToolsCoordinator.View]: The view that currently uses the writing tools coordinator.
//
// # Getting the host views for effects
//
//   - [INSWritingToolsCoordinator.EffectContainerView]: The view that Writing Tools uses to display visual effects during the text-rewriting process.
//   - [INSWritingToolsCoordinator.SetEffectContainerView]
//   - [INSWritingToolsCoordinator.DecorationContainerView]: The view that Writing Tools uses to display background decorations such as proofreading marks.
//   - [INSWritingToolsCoordinator.SetDecorationContainerView]
//
// # Configuring the experience
//
//   - [INSWritingToolsCoordinator.PreferredBehavior]: The level of Writing Tools support you want the system to provide for your view.
//   - [INSWritingToolsCoordinator.SetPreferredBehavior]
//   - [INSWritingToolsCoordinator.Behavior]: The actual level of Writing Tools support the system provides for your view.
//   - [INSWritingToolsCoordinator.PreferredResultOptions]: The type of content you allow Writing Tools to generate for your custom text view.
//   - [INSWritingToolsCoordinator.SetPreferredResultOptions]
//   - [INSWritingToolsCoordinator.ResultOptions]: The type of content the system generates for your custom text view.
//
// # Reporting changes to Writing Tools
//
//   - [INSWritingToolsCoordinator.UpdateRangeWithTextReasonForContextWithIdentifier]: Informs the coordinator about changes your app made to the text in the specified context object.
//   - [INSWritingToolsCoordinator.UpdateForReflowedTextInContextWithIdentifier]: Informs the coordinator that a change occurred to the view or its text that requires a layout update.
//
// # Managing the current state
//
//   - [INSWritingToolsCoordinator.StopWritingTools]: Stops the current Writing Tools operation and dismisses the system UI.
//   - [INSWritingToolsCoordinator.State]: The current level of Writing Tools activity in your view.
//
// # Instance Properties
//
//   - [INSWritingToolsCoordinator.IncludesTextListMarkers]
//   - [INSWritingToolsCoordinator.SetIncludesTextListMarkers]
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator
type INSWritingToolsCoordinator interface {
	objectivec.IObject

	// Topic: Creating a coordinator object

	// Creates a writing tools coordinator and assigns the specified delegate object to it.
	InitWithDelegate(delegate NSWritingToolsCoordinatorDelegate) NSWritingToolsCoordinator

	// Topic: Managing Writing Tools interactions

	// The object that handles Writing Tools interactions for your view.
	Delegate() NSWritingToolsCoordinatorDelegate
	// The view that currently uses the writing tools coordinator.
	View() INSView

	// Topic: Getting the host views for effects

	// The view that Writing Tools uses to display visual effects during the text-rewriting process.
	EffectContainerView() INSView
	SetEffectContainerView(value INSView)
	// The view that Writing Tools uses to display background decorations such as proofreading marks.
	DecorationContainerView() INSView
	SetDecorationContainerView(value INSView)

	// Topic: Configuring the experience

	// The level of Writing Tools support you want the system to provide for your view.
	PreferredBehavior() NSWritingToolsBehavior
	SetPreferredBehavior(value NSWritingToolsBehavior)
	// The actual level of Writing Tools support the system provides for your view.
	Behavior() NSWritingToolsBehavior
	// The type of content you allow Writing Tools to generate for your custom text view.
	PreferredResultOptions() NSWritingToolsResultOptions
	SetPreferredResultOptions(value NSWritingToolsResultOptions)
	// The type of content the system generates for your custom text view.
	ResultOptions() NSWritingToolsResultOptions

	// Topic: Reporting changes to Writing Tools

	// Informs the coordinator about changes your app made to the text in the specified context object.
	UpdateRangeWithTextReasonForContextWithIdentifier(range_ foundation.NSRange, replacementText foundation.NSAttributedString, reason NSWritingToolsCoordinatorTextUpdateReason, contextID foundation.NSUUID)
	// Informs the coordinator that a change occurred to the view or its text that requires a layout update.
	UpdateForReflowedTextInContextWithIdentifier(contextID foundation.NSUUID)

	// Topic: Managing the current state

	// Stops the current Writing Tools operation and dismisses the system UI.
	StopWritingTools()
	// The current level of Writing Tools activity in your view.
	State() NSWritingToolsCoordinatorState

	// Topic: Instance Properties

	IncludesTextListMarkers() bool
	SetIncludesTextListMarkers(value bool)

	WritingToolsCoordinator() INSWritingToolsCoordinator
	SetWritingToolsCoordinator(value INSWritingToolsCoordinator)
}

// Init initializes the instance.
func (w NSWritingToolsCoordinator) Init() NSWritingToolsCoordinator {
	rv := objc.Send[NSWritingToolsCoordinator](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w NSWritingToolsCoordinator) Autorelease() NSWritingToolsCoordinator {
	rv := objc.Send[NSWritingToolsCoordinator](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSWritingToolsCoordinator creates a new NSWritingToolsCoordinator instance.
func NewNSWritingToolsCoordinator() NSWritingToolsCoordinator {
	class := getNSWritingToolsCoordinatorClass()
	rv := objc.Send[NSWritingToolsCoordinator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a writing tools coordinator and assigns the specified delegate
// object to it.
//
// delegate: An object capable of handling Writing Tools interactions for your view. The
// delegate must be able to modify your view’s text storage and refresh the
// view’s layout and appearance.
//
// # Discussion
// 
// Create the coordinator object during your view’s initialization, and
// assign the object to your view. Assign the coordinator to the
// [WritingToolsCoordinator] property of your view.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/init(delegate:)
func NewWritingToolsCoordinatorWithDelegate(delegate NSWritingToolsCoordinatorDelegate) NSWritingToolsCoordinator {
	instance := getNSWritingToolsCoordinatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDelegate:"), delegate)
	return NSWritingToolsCoordinatorFromID(rv)
}

// Creates a writing tools coordinator and assigns the specified delegate
// object to it.
//
// delegate: An object capable of handling Writing Tools interactions for your view. The
// delegate must be able to modify your view’s text storage and refresh the
// view’s layout and appearance.
//
// # Discussion
// 
// Create the coordinator object during your view’s initialization, and
// assign the object to your view. Assign the coordinator to the
// [WritingToolsCoordinator] property of your view.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/init(delegate:)
func (w NSWritingToolsCoordinator) InitWithDelegate(delegate NSWritingToolsCoordinatorDelegate) NSWritingToolsCoordinator {
	rv := objc.Send[NSWritingToolsCoordinator](w.ID, objc.Sel("initWithDelegate:"), delegate)
	return rv
}
// Informs the coordinator about changes your app made to the text in the
// specified context object.
//
// range: The range of text to replace. This range is relative to the starting
// location of the specified context object’s text in your view’s text
// storage. If you initialized the context object with the entire contents of
// your view’s text storage, specify the range of text you’re replacing in
// your text storage. However, if you initialized the context object with only
// a portion of your view’s text, specify a range that is relative to the
// starting location of the context object’s text.
//
// replacementText: The text that replaces the previous content in `range`. Specify an empty
// string to delete the text in the specified range.
//
// reason: The reason you updated the text.
//
// contextID: The unique identifier of the context object that contains the text you
// modified.
//
// # Discussion
// 
// If you make any changes to the text Writing Tools is evaluating, call this
// method to report those changes to your view’s coordinator object. You
// might make changes in response to an undo command or when someone types
// into the same part of your view’s text. Calling this method keeps the
// coordinator object informed of any changes, and ensures it delivers
// accurate information to its delegate. In response, the coordinator
// refreshes previews and other information related to your view. If the scope
// of the update is significantly large, the coordinator can optionally cancel
// the Writing Tools session altogether.
// 
// Use this method to report changes that precisely intersect your context
// object’s text. The first time you call this method for a context object,
// report changes only to the original attributed string in that object. If
// you call this method more than once, report changes to the newly modified
// version of that string. Don’t use this method to report changes to text
// that comes before or after the context object. If you make changes before
// your context object, report those changes separately using the
// [UpdateForReflowedTextInContextWithIdentifier] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/updateRange(_:with:reason:forContextWithIdentifier:)
func (w NSWritingToolsCoordinator) UpdateRangeWithTextReasonForContextWithIdentifier(range_ foundation.NSRange, replacementText foundation.NSAttributedString, reason NSWritingToolsCoordinatorTextUpdateReason, contextID foundation.NSUUID) {
	objc.Send[objc.ID](w.ID, objc.Sel("updateRange:withText:reason:forContextWithIdentifier:"), range_, replacementText, reason, contextID)
}
// Informs the coordinator that a change occurred to the view or its text that
// requires a layout update.
//
// contextID: The unique identifier of the context object affected by the change. Pass
// the identifier for the context object that comes after the changes.
//
// # Discussion
// 
// Use this method to inform Writing Tools when the geometry of your view
// changes, or when the text that precedes one of your context objects
// changes. Changes to the view’s geometry or text can affect the flow of
// any remaining text, and require a layout update. Writing Tools uses this
// method to refresh any layout-dependent information it’s currently
// tracking. For example, it uses it to refresh the location of proofreading
// marks it’s displaying in your view.
// 
// If a text change affects the text inside a context object, call the
// [UpdateRangeWithTextReasonForContextWithIdentifier] method to report that
// change instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/updateForReflowedTextInContextWithIdentifier(_:)
func (w NSWritingToolsCoordinator) UpdateForReflowedTextInContextWithIdentifier(contextID foundation.NSUUID) {
	objc.Send[objc.ID](w.ID, objc.Sel("updateForReflowedTextInContextWithIdentifier:"), contextID)
}
// Stops the current Writing Tools operation and dismisses the system UI.
//
// # Discussion
// 
// Call this method to abort the current Writing Tools operation. This method
// dismisses the system’s Writing Tools UI and stops any in-flight
// interactions with your view. This method does not undo any changes that
// Writing Tools already made to your view’s content.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/stopWritingTools()
func (w NSWritingToolsCoordinator) StopWritingTools() {
	objc.Send[objc.ID](w.ID, objc.Sel("stopWritingTools"))
}

// The object that handles Writing Tools interactions for your view.
//
// # Discussion
// 
// Specify this object at initialization time when creating your
// [NSWritingToolsCoordinator] object. The object must adopt the
// [NSWritingToolsCoordinatorDelegate] protocol, and be capable of modifying
// your view’s text storage and refreshing the view’s layout and
// appearance.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/delegate-swift.property
func (w NSWritingToolsCoordinator) Delegate() NSWritingToolsCoordinatorDelegate {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("delegate"))
	return NSWritingToolsCoordinatorDelegateObjectFromID(rv)
}
// The view that currently uses the writing tools coordinator.
//
// # Discussion
// 
// Use this property to refer to the view that currently owns the coordinator
// object. The system updates this property automatically when you assign the
// coordinator to the [WritingToolsCoordinator] property of your view. The
// value of this property is `nil` if there is no associated view.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/view
func (w NSWritingToolsCoordinator) View() INSView {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("view"))
	return NSViewFromID(objc.ID(rv))
}
// The view that Writing Tools uses to display visual effects during the
// text-rewriting process.
//
// # Discussion
// 
// Writing Tools uses the view in this property to host the visual effects it
// creates when making interactive changes to your view’s content. These
// visual effects let people know the state of the text and provide feedback
// about what’s happening to it. Set this property to a subview that sits
// visually above, and covers, all of the text in your custom text view. If
// you don’t assign a value to this property, the coordinator places its own
// effect view in front of the subviews in your custom view. The default value
// of this property is `nil`.
// 
// If you display your view’s text using multiple text containers, implement
// the `NSWritingToolsCoordinator/Delegate/writingToolsCoordinator(_:)` method
// to request multiple previews.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/effectContainerView
func (w NSWritingToolsCoordinator) EffectContainerView() INSView {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("effectContainerView"))
	return NSViewFromID(objc.ID(rv))
}
func (w NSWritingToolsCoordinator) SetEffectContainerView(value INSView) {
	objc.Send[struct{}](w.ID, objc.Sel("setEffectContainerView:"), value)
}
// The view that Writing Tools uses to display background decorations such as
// proofreading marks.
//
// # Discussion
// 
// Writing Tools uses the view in this property to host proofreading marks and
// other visual elements that show any suggested changes. Set this property to
// a subview situated visibly below the text in your custom text view. It’s
// also satisfactory to place this view visually in front of the text. Make
// sure the size of the view is big enough to cover all of the affected text.
// If you don’t assign a value to this property, the coordinator places its
// own decoration view behind the subviews in your custom view. The default
// value of this property is `nil`.
// 
// If you display your view’s text using multiple text containers, implement
// the `NSWritingToolsCoordinator/Delegate/writingToolsCoordinator(_:)` and
// `NSWritingToolsCoordinator/Delegate/writingToolsCoordinator(_:)` methods to
// provide separate decoration views for each container.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/decorationContainerView
func (w NSWritingToolsCoordinator) DecorationContainerView() INSView {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("decorationContainerView"))
	return NSViewFromID(objc.ID(rv))
}
func (w NSWritingToolsCoordinator) SetDecorationContainerView(value INSView) {
	objc.Send[struct{}](w.ID, objc.Sel("setDecorationContainerView:"), value)
}
// The level of Writing Tools support you want the system to provide for your
// view.
//
// # Discussion
// 
// Use this property to request an inline or panel-based experience, or to
// disable Writing Tools for your view altogether. The default value of this
// property is [NSWritingToolsBehavior.default].
//
// [NSWritingToolsBehavior.default]: https://developer.apple.com/documentation/AppKit/NSWritingToolsBehavior/default
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/preferredBehavior
func (w NSWritingToolsCoordinator) PreferredBehavior() NSWritingToolsBehavior {
	rv := objc.Send[NSWritingToolsBehavior](w.ID, objc.Sel("preferredBehavior"))
	return NSWritingToolsBehavior(rv)
}
func (w NSWritingToolsCoordinator) SetPreferredBehavior(value NSWritingToolsBehavior) {
	objc.Send[struct{}](w.ID, objc.Sel("setPreferredBehavior:"), value)
}
// The actual level of Writing Tools support the system provides for your
// view.
//
// # Discussion
// 
// The system chooses this value based on the device capabilities, and takes
// the value in the [PreferredBehavior] property into consideration when
// making the choice. The value in this property is never the default option,
// and is instead one of the specific options such as
// [NSWritingToolsBehavior.none], [NSWritingToolsBehavior.limited], or
// [NSWritingToolsBehavior.complete].
//
// [NSWritingToolsBehavior.complete]: https://developer.apple.com/documentation/AppKit/NSWritingToolsBehavior/complete
// [NSWritingToolsBehavior.limited]: https://developer.apple.com/documentation/AppKit/NSWritingToolsBehavior/limited
// [NSWritingToolsBehavior.none]: https://developer.apple.com/documentation/AppKit/NSWritingToolsBehavior/none
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/behavior
func (w NSWritingToolsCoordinator) Behavior() NSWritingToolsBehavior {
	rv := objc.Send[NSWritingToolsBehavior](w.ID, objc.Sel("behavior"))
	return NSWritingToolsBehavior(rv)
}
// The type of content you allow Writing Tools to generate for your custom
// text view.
//
// # Discussion
// 
// Writing Tools can create plain text or rich text, and it can format text
// using lists or tables as needed. If your view doesn’t support specific
// types of content, specify the types you do support in this property. The
// default value of this property is `NSWritingToolsResult/default`, which
// lets the system determine the type of content to generate.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/preferredResultOptions
func (w NSWritingToolsCoordinator) PreferredResultOptions() NSWritingToolsResultOptions {
	rv := objc.Send[NSWritingToolsResultOptions](w.ID, objc.Sel("preferredResultOptions"))
	return NSWritingToolsResultOptions(rv)
}
func (w NSWritingToolsCoordinator) SetPreferredResultOptions(value NSWritingToolsResultOptions) {
	objc.Send[struct{}](w.ID, objc.Sel("setPreferredResultOptions:"), value)
}
// The type of content the system generates for your custom text view.
//
// # Discussion
// 
// This property contains the set of options that Writing Tools outputs for
// your view. Writing Tools takes the value in the [PreferredResultOptions]
// property into consideration when determining this value.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/resultOptions
func (w NSWritingToolsCoordinator) ResultOptions() NSWritingToolsResultOptions {
	rv := objc.Send[NSWritingToolsResultOptions](w.ID, objc.Sel("resultOptions"))
	return NSWritingToolsResultOptions(rv)
}
// The current level of Writing Tools activity in your view.
//
// # Discussion
// 
// Use this property to determine when Writing Tools is actively making
// changes to your view. During the course of Writing Tools interactions, the
// system reports state changes to the delegate’s
// [WritingToolsCoordinatorWillChangeToStateCompletion] method and updates
// this property accordingly.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/state-swift.property
func (w NSWritingToolsCoordinator) State() NSWritingToolsCoordinatorState {
	rv := objc.Send[NSWritingToolsCoordinatorState](w.ID, objc.Sel("state"))
	return NSWritingToolsCoordinatorState(rv)
}
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/includesTextListMarkers
func (w NSWritingToolsCoordinator) IncludesTextListMarkers() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("includesTextListMarkers"))
	return rv
}
func (w NSWritingToolsCoordinator) SetIncludesTextListMarkers(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setIncludesTextListMarkers:"), value)
}
// See: https://developer.apple.com/documentation/appkit/nsview/writingtoolscoordinator
func (w NSWritingToolsCoordinator) WritingToolsCoordinator() INSWritingToolsCoordinator {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("writingToolsCoordinator"))
	return NSWritingToolsCoordinatorFromID(objc.ID(rv))
}
func (w NSWritingToolsCoordinator) SetWritingToolsCoordinator(value INSWritingToolsCoordinator) {
	objc.Send[struct{}](w.ID, objc.Sel("setWritingToolsCoordinator:"), value)
}

// A Boolean value that indicates whether Writing Tools features are currently
// available.
//
// # Discussion
// 
// The value of this property is `true` when Writing Tools features are
// available, and `false` when they aren’t. Writing Tools support might be
// unavailable because of device constraints or because the system isn’t
// ready to process Writing Tools requests.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/isWritingToolsAvailable
func (_NSWritingToolsCoordinatorClass NSWritingToolsCoordinatorClass) IsWritingToolsAvailable() bool {
	rv := objc.Send[bool](objc.ID(_NSWritingToolsCoordinatorClass.class), objc.Sel("isWritingToolsAvailable"))
	return rv
}

