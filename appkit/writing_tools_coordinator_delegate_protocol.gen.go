// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An interface that you use to manage interactions between Writing Tools and your custom text view.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol
type NSWritingToolsCoordinatorDelegate interface {
	objectivec.IObject

	// Asks your delegate to provide the text to evaluate during the Writing Tools operation.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol/writingToolsCoordinator(_:requestsContextsFor:completion:)
	WritingToolsCoordinatorRequestsContextsForScopeCompletion(writingToolsCoordinator INSWritingToolsCoordinator, scope NSWritingToolsCoordinatorContextScope, completion VoidHandler)

	// Tells the delegate that there are text changes to incorporate into the view.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol/writingToolsCoordinator(_:replace:in:proposedText:reason:animationParameters:completion:)
	WritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion(writingToolsCoordinator INSWritingToolsCoordinator, range_ foundation.NSRange, context objectivec.IObject, replacementText foundation.NSAttributedString, reason NSWritingToolsCoordinatorTextReplacementReason, animationParameters objectivec.IObject, completion AttributedStringHandler)

	// Asks the delegate to update your view’s current text selection.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol/writingToolsCoordinator(_:select:in:completion:)
	WritingToolsCoordinatorSelectRangesInContextCompletion(writingToolsCoordinator INSWritingToolsCoordinator, ranges []foundation.NSValue, context objectivec.IObject, completion VoidHandler)

	// Asks the delegate for a preview image and layout information for the specified text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol/writingToolsCoordinator(_:requestsPreviewFor:in:completion:)
	WritingToolsCoordinatorRequestsPreviewForRectInContextCompletion(writingToolsCoordinator INSWritingToolsCoordinator, rect corefoundation.CGRect, context objectivec.IObject, completion TextPreviewHandler)

	// Prepare for animations for the content that Writing Tools is evaluating.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol/writingToolsCoordinator(_:prepareFor:for:in:completion:)
	WritingToolsCoordinatorPrepareForTextAnimationForRangeInContextCompletion(writingToolsCoordinator INSWritingToolsCoordinator, textAnimation NSWritingToolsCoordinatorTextAnimation, range_ foundation.NSRange, context objectivec.IObject, completion VoidHandler)

	// Asks the delegate to clean up any state related to the specified Writing Tools animation.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol/writingToolsCoordinator(_:finish:for:in:completion:)
	WritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion(writingToolsCoordinator INSWritingToolsCoordinator, textAnimation NSWritingToolsCoordinatorTextAnimation, range_ foundation.NSRange, context objectivec.IObject, completion VoidHandler)
}

// NSWritingToolsCoordinatorDelegateObject wraps an existing Objective-C object that conforms to the NSWritingToolsCoordinatorDelegate protocol.
type NSWritingToolsCoordinatorDelegateObject struct {
	objectivec.Object
}

func (o NSWritingToolsCoordinatorDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSWritingToolsCoordinatorDelegateObjectFromID constructs a [NSWritingToolsCoordinatorDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSWritingToolsCoordinatorDelegateObjectFromID(id objc.ID) NSWritingToolsCoordinatorDelegateObject {
	return NSWritingToolsCoordinatorDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Asks your delegate to provide the text to evaluate during the Writing Tools
// operation.
//
// writingToolsCoordinator: The coordinator object requesting information from your custom view.
//
// scope: The amount of text the coordinator requested. Use this property to
// determine if Writing Tools is evaluating all of your text or only a subset.
//
// completion: A completion block to execute with the required information. You must
// execute this block once at end of your method’s implementation. The block
// has no return value and takes an array of
// [NSWritingToolsCoordinatorContext] objects that contain the requested
// information.
//
// # Discussion
//
// At the start of every operation, the [NSWritingToolsCoordinator] object
// calls this method to request the text to evaluate. Use this method to
// create one or more [NSWritingToolsCoordinatorContext] objects with your
// view’s text. Create only one [NSWritingToolsCoordinatorContext] object if
// your view has only one text storage object. If your view contains multiple
// text storage objects, create separate [NSWritingToolsCoordinatorContext]
// objects for each text storage object. Writing Tools returns updates for
// each context object separately, making it easier for you to incorporate
// changes into the correct text storage object. Don’t create multiple
// context objects if your view has only one text storage object.
//
// The `scope` parameter tells you what content Writing Tools expects you to
// provide in your context object. For example, Writing Tools expects you to
// provide the selected text when the parameter contains the
// [NSWritingToolsCoordinatorContextScopeUserSelection] option. When Writing
// Tools requests a subset of your overall text, include some of the
// surrounding text in your context object too. Writing Tools can use the
// extra text you provide to improve the results of its evaluation. For
// example, it might use an entire paragraph, instead of only the selected
// sentence, to evaluate ways to rewrite that sentence. It’s best to include
// the text up to the nearest paragraph boundary before and after the
// selection. If you include extra text in your context object, set the
// [Range] property to the range of the selected text.
//
// Pass the context objects you create to the provided completion handler
// before your method returns. Writing Tools waits for you to call the
// completion handler before proceeding with the operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol/writingToolsCoordinator(_:requestsContextsFor:completion:)
func (o NSWritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorRequestsContextsForScopeCompletion(writingToolsCoordinator INSWritingToolsCoordinator, scope NSWritingToolsCoordinatorContextScope, completion VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("writingToolsCoordinator:requestsContextsForScope:completion:"), writingToolsCoordinator, scope, completion)
}

// Tells the delegate that there are text changes to incorporate into the
// view.
//
// writingToolsCoordinator: The coordinator object providing the changes to your custom view.
//
// range: A range of text to update. This range is relative to the text in your
// context object, and it’s your responsibility to match that location to
// the correct location in your text storage. If you initialized the context
// object with the entire contents of your view’s text storage, you can use
// `range` as-is to access that text storage. However, if you initialized the
// context object with only a portion of your view’s text, add the starting
// location of your context object’s text to this value to get the correct
// range for that text storage.
//
// context: The context object that contains the original text to modify. Use this
// object to locate the correct text storage object for your view.
//
// replacementText: The text to insert in place of the current text at `range`. You can insert
// this text as-is, insert a modified version of this string, or reject the
// replacement text altogether.
//
// reason: The type of replacement Writing Tools performs. This parameter indicates
// whether Writing Tools is replacing the text with or without animations.
//
// animationParameters: The animation parameters for any interactive changes, or `nil` if the
// changes aren’t interactive. Use this object to create any additional
// animations for the system to run alongside the changes Writing Tools makes.
// For example, use it to update other views that contain related information.
//
// completion: A completion handler to execute with the results of the operation. The
// handler has no return value and takes an optional attributed string as a
// parameter. If you incorporate the replacement text, either as-is or with
// modifications, pass the actual string you incorporated to the completion
// block. If you reject the suggested change and leave the original text
// unchanged, specify `nil` for this parameter.
//
// # Discussion
//
// Use this method to update your view’s text storage with the proposed
// changes. Writing Tools can call this method multiple times during the
// course of a session to notify you of changes to different ranges of text.
// Incorporate the changes into your view’s text storage and notify your
// layout manager so it can refresh the view.
//
// Remove the text in the appropriate range of your text storage, and replace
// it with the contents of `replacementText`. When you finish, call the
// completion handler and pass in the replacement text you inserted. If you
// change the string in `replacementText` before incorporating it into your
// text storage, return your modified string instead. Returning the string
// lets Writing Tools track any alterations you made to it. You can also pass
// `nil` to the completion handler if you don’t incorporate the replacement
// text.
//
// For interactive changes, Writing Tools works with your delegate to animate
// the removal of the old text and the insertion of any replacement text. If
// you need to modify other parts of your interface to reflect the changes,
// use the provided [NSWritingToolsCoordinatorAnimationParameters] object to
// create additional animations to run at the same time as the system-provided
// animations.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol/writingToolsCoordinator(_:replace:in:proposedText:reason:animationParameters:completion:)
func (o NSWritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion(writingToolsCoordinator INSWritingToolsCoordinator, range_ foundation.NSRange, context objectivec.IObject, replacementText foundation.NSAttributedString, reason NSWritingToolsCoordinatorTextReplacementReason, animationParameters objectivec.IObject, completion AttributedStringHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("writingToolsCoordinator:replaceRange:inContext:proposedText:reason:animationParameters:completion:"), writingToolsCoordinator, range_, context, replacementText, reason, animationParameters, completion)
}

// Asks the delegate to update your view’s current text selection.
//
// writingToolsCoordinator: The coordinator object making the change to your view.
//
// ranges: One or more ranges of text to select. Each range is relative to the text in
// your context object, and it’s your responsibility to match each location
// to the correct location in your text storage. If you initialized the
// context object with the entire contents of your view’s text storage, you
// can use the ranges as-is to access that text storage. However, if you
// initialized the context object with only a portion of your view’s text,
// add the starting location of your context object’s text to each value to
// get the correct range for that text storage.
//
// context: The context object you use to identify the associated text storage.
//
// completion: The completion handler to execute when your delegate finishes updating the
// selection. The handler has no parameters or return value. You must call
// this handler at some point during the implementation of your method.
//
// # Discussion
//
// As Writing Tools suggests changes to your view’s text, it calls this
// method to update the text selection accordingly. Use this method to update
// the current selection in your view’s text storage. When you finish making
// the changes, call the provided completion block to let Writing Tools know
// you’re finished.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol/writingToolsCoordinator(_:select:in:completion:)
func (o NSWritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorSelectRangesInContextCompletion(writingToolsCoordinator INSWritingToolsCoordinator, ranges []foundation.NSValue, context objectivec.IObject, completion VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("writingToolsCoordinator:selectRanges:inContext:completion:"), writingToolsCoordinator, objectivec.IObjectSliceToNSArray(ranges), context, completion)
}

// Asks the delegate for a preview image and layout information for the
// specified text.
//
// writingToolsCoordinator: The coordinator object notifying you that animations are about to begin.
//
// rect: The portion of your view that contains the requested text. This rectangle
// is in the view’s coordinate system.
//
// context: The context object that contains the original text. Use this object to
// fetch the current text, and to match that text to your underlying text
// storage.
//
// completion: A completion handler to execute when you are done. The handler has no
// return value and takes an [NSTextPreview] object as a parameter. You must
// call this handler at some point during your implementation.
//
// # Discussion
//
// During an interactive evaluation of your view’s text, Writing Tools
// creates different animations to provide feedback on what’s happening. As
// part of the preparation for those animations, Writing Tools asks you to
// provide a preview of the affected content in your view. Writing Tools uses
// this preview to build and execute the animations in the view stored in the
// [EffectContainerView] property of the coordinator object.
//
// To build a preview of your content in macOS, render the requested portion
// of your view into an image with a transparent background and use that image
// to create your [NSTextPreview] object directly. Set the [PresentationFrame]
// property to the specified rectangle. Set the [CandidateRects] property to
// the selection rectangles for the associated text, which you get from your
// view’s layout manager. Writing Tools uses this information to place your
// image directly above the text in your view.
//
// For a single animation type, the system calls the
// [WritingToolsCoordinatorPrepareForTextAnimationForRangeInContextCompletion]
// method, followed sequentially by this method and then the
// [WritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion]
// method. Each method executes asynchronously, but the system calls the next
// method in the sequence only after you call the completion handler of the
// previous method. However, multiple animations can run simultaneously, so
// check the `textAnimation` parameter to differentiate sequences.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol/writingToolsCoordinator(_:requestsPreviewFor:in:completion:)
func (o NSWritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorRequestsPreviewForRectInContextCompletion(writingToolsCoordinator INSWritingToolsCoordinator, rect corefoundation.CGRect, context objectivec.IObject, completion TextPreviewHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("writingToolsCoordinator:requestsPreviewForRect:inContext:completion:"), writingToolsCoordinator, rect, context, completion)
}

// Prepare for animations for the content that Writing Tools is evaluating.
//
// writingToolsCoordinator: The coordinator object notifying you that animations are about to begin.
//
// textAnimation: The type of animation Writing Tools is preparing.
//
// range: The range of text affected by the animation. This range is relative to the
// text in your context object, and it’s your responsibility to match that
// location to the correct location in your text storage. If you initialized
// the context object with the entire contents of your view’s text storage,
// you can use `range` as-is to access that text storage. However, if you
// initialized the context object with only a portion of your view’s text,
// add the starting location of your context object’s text to this value to
// get the correct range for that text storage.
//
// context: The context object that contains the original text. Use this object to
// fetch the current text, and to match that text to your underlying text
// storage.
//
// completion: A completion handler to execute when you are done. The handler has no
// return value and takes no parameters. You must call this handler at some
// point during your implementation.
//
// # Discussion
//
// During an interactive evaluation of your view’s text, Writing Tools
// creates different animations to provide feedback on what’s happening. For
// example, it creates an [NSWritingToolsCoordinatorTextAnimationAnticipate]
// animation to let people know the system is evaluating the text. The
// `textAnimation` parameter tells you what type of animation to prepare for.
//
// Use this method to prepare for the system-provided animations of your
// view’s content. For interactive animations, hide the text in the
// specified range temporarily while the system animations run. For
// non-interactive animations, dim the text for the duration of the animation
// to indicate it’s not editable. For animations to remove or insert text,
// you can also use this method to set up animations to reflow your view’s
// content to match the changes. At the end of a given animation, use your
// delegate’s
// [WritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion]
// method to undo any changes you make to your content.
//
// For a single animation type, the system calls this method, followed
// sequentially by the
// [WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion]
// and [WritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion]
// methods. Each method executes asynchronously, but the system calls the next
// method in the sequence only after you call the completion handler of the
// previous method. However, multiple animations can run simultaneously, so
// check the `textAnimation` and `range` parameters to differentiate
// sequences.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol/writingToolsCoordinator(_:prepareFor:for:in:completion:)
func (o NSWritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorPrepareForTextAnimationForRangeInContextCompletion(writingToolsCoordinator INSWritingToolsCoordinator, textAnimation NSWritingToolsCoordinatorTextAnimation, range_ foundation.NSRange, context objectivec.IObject, completion VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("writingToolsCoordinator:prepareForTextAnimation:forRange:inContext:completion:"), writingToolsCoordinator, textAnimation, range_, context, completion)
}

// Asks the delegate to clean up any state related to the specified Writing
// Tools animation.
//
// writingToolsCoordinator: The coordinator object notifying you that animations are about to begin.
//
// textAnimation: The type of animation Writing Tools finished.
//
// range: The range of text that finished animating. This range is relative to the
// text in your context object, and it’s your responsibility to match that
// location to the correct location in your text storage. If you initialized
// the context object with the entire contents of your view’s text storage,
// you can use `range` as-is to access that text storage. However, if you
// initialized the context object with only a portion of your view’s text,
// add the starting location of your context object’s text to this value to
// get the correct range for that text storage.
//
// context: The context object that contains the original text.
//
// completion: A completion handler to execute when you are done. The handler has no
// return value and takes no parameters. You must call this handler at some
// point during your implementation.
//
// # Discussion
//
// Use this method to clean up any data structures you created to support the
// specified type of Writing Tools animation. You can also use this method to
// restore the visibility of any text you hid previously. When you finish your
// cleanup work, call the completion handler to notify Writing Tools.
//
// Writing Tools calls this method only after previous calls to the
// [WritingToolsCoordinatorPrepareForTextAnimationForRangeInContextCompletion]
// and
// [WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion]
// methods for the same animation type. However, Writing Tools can interleave
// calls to this method with calls to prepare an animation of a different
// type. In your implementation of this method, make sure the actions you take
// don’t interfere with other in-flight animations.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol/writingToolsCoordinator(_:finish:for:in:completion:)
func (o NSWritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion(writingToolsCoordinator INSWritingToolsCoordinator, textAnimation NSWritingToolsCoordinatorTextAnimation, range_ foundation.NSRange, context objectivec.IObject, completion VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("writingToolsCoordinator:finishTextAnimation:forRange:inContext:completion:"), writingToolsCoordinator, textAnimation, range_, context, completion)
}

// Notifies your delegate of relevant state changes when Writing Tools is
// running in your view.
//
// writingToolsCoordinator: The coordinator object providing information to your custom view.
//
// completion: A handler to execute when your delegate finishes processing the change of
// state. The handler has no parameters or return value. You must call this
// handler at some point during the implementation of your method.
//
// # Discussion
//
// Use state transitions to perform actions related to your view or text
// storage. When Writing Tools is active, it updates its state to indicate
// what task it’s currently performing. Writing Tools starts in the
// [NSWritingToolsCoordinatorStateInactive] state and moves to other states as
// it presents UI and starts interacting with your view’s content. For
// example, it moves to the [NSWritingToolsCoordinatorStateInteractiveResting]
// state when it’s making changes to your view’s text storage.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol/writingToolsCoordinator(_:willChangeTo:completion:)
func (o NSWritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorWillChangeToStateCompletion(writingToolsCoordinator INSWritingToolsCoordinator, newState NSWritingToolsCoordinatorState, completion VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("writingToolsCoordinator:willChangeToState:completion:"), writingToolsCoordinator, newState, completion)
}

// Asks the delegate to provide a decoration view for the specified range of
// text.
//
// writingToolsCoordinator: The coordinator object requesting information from your custom view.
//
// range: The range of text to consider in the specified `context` object. The
// location value of this range is relative to the beginning of the text in
// your context object, and it’s your responsibility to match that location
// to the correct location in your text storage. If you initialized the
// context object with the entire contents of your view’s text storage, you
// can use `range` as-is to access that text storage. However, if you
// initialized the context object with only a portion of your view’s text,
// add the starting location of your context object’s text to this value to
// get the correct range for that text storage.
//
// context: The context object that contains the text to consider. Use this object to
// locate the appropriate text storage object for your view.
//
// completion: A completion handler to execute when you are done. The handler has no
// return value and takes an [NSView] object as a parameter. You must call
// this handler at some point during your implementation.
//
// # Discussion
//
// If your view uses multiple [NSTextContainer] objects to draw text in
// different regions, use this method to provide Writing Tools with the view
// to use for the specified range of text. After calling your delegate’s
// [WritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion]
// method, Writing Tools calls this method for each subrange of text you
// provided. Find or provide a view situated visibly below the specified text
// in your text view. It’s also satisfactory to provide a view that’s
// visually in front of the text. Writing Tools uses the provided view to host
// any proofreading marks for the specified range of text.
//
// If your view has only one text container, use the coordinator’s
// [DecorationContainerView] property to specify the view to use for
// proofreading marks.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol/writingToolsCoordinator(_:requestsDecorationContainerViewFor:in:completion:)
func (o NSWritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion(writingToolsCoordinator INSWritingToolsCoordinator, range_ foundation.NSRange, context objectivec.IObject, completion ViewHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("writingToolsCoordinator:requestsDecorationContainerViewForRange:inContext:completion:"), writingToolsCoordinator, range_, context, completion)
}
