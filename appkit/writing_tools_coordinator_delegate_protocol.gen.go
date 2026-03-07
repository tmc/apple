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
	// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol/writingToolsCoordinator(_:requestsPreviewFor:of:in:completion:)
	WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion(writingToolsCoordinator INSWritingToolsCoordinator, textAnimation NSWritingToolsCoordinatorTextAnimation, range_ foundation.NSRange, context objectivec.IObject, completion ArrayHandler)

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

	// Asks the delegate to provide the bounding paths for the specified text in your view.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol/writingToolsCoordinator(_:requestsBoundingBezierPathsFor:in:completion:)
	WritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion(writingToolsCoordinator INSWritingToolsCoordinator, range_ foundation.NSRange, context objectivec.IObject, completion ArrayHandler)

	// Asks the delegate to provide an underline shape for the specified text during a proofreading session.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol/writingToolsCoordinator(_:requestsUnderlinePathsFor:in:completion:)
	WritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion(writingToolsCoordinator INSWritingToolsCoordinator, range_ foundation.NSRange, context objectivec.IObject, completion ArrayHandler)
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
// [NSWritingToolsCoordinator.Context] objects that contain the requested
// information.
// //
// [NSWritingToolsCoordinator.Context]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Context
//
// # Discussion
// 
// At the start of every operation, the [NSWritingToolsCoordinator] object
// calls this method to request the text to evaluate. Use this method to
// create one or more [NSWritingToolsCoordinator.Context] objects with your
// view’s text. Create only one [NSWritingToolsCoordinator.Context] object
// if your view has only one text storage object. If your view contains
// multiple text storage objects, create separate
// [NSWritingToolsCoordinator.Context] objects for each text storage object.
// Writing Tools returns updates for each context object separately, making it
// easier for you to incorporate changes into the correct text storage object.
// Don’t create multiple context objects if your view has only one text
// storage object.
// 
// The `scope` parameter tells you what content Writing Tools expects you to
// provide in your context object. For example, Writing Tools expects you to
// provide the selected text when the parameter contains the
// [NSWritingToolsCoordinator.ContextScope.userSelection] option. When Writing
// Tools requests a subset of your overall text, include some of the
// surrounding text in your context object too. Writing Tools can use the
// extra text you provide to improve the results of its evaluation. For
// example, it might use an entire paragraph, instead of only the selected
// sentence, to evaluate ways to rewrite that sentence. It’s best to include
// the text up to the nearest paragraph boundary before and after the
// selection. If you include extra text in your context object, set the
// [range] property to the range of the selected text.
// 
// Pass the context objects you create to the provided completion handler
// before your method returns. Writing Tools waits for you to call the
// completion handler before proceeding with the operation.
//
// [NSWritingToolsCoordinator.ContextScope.userSelection]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/ContextScope/userSelection
// [NSWritingToolsCoordinator.Context]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Context
// [range]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Context/range
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
// use the provided [NSWritingToolsCoordinator.AnimationParameters] object to
// create additional animations to run at the same time as the system-provided
// animations.
//
// [NSWritingToolsCoordinator.AnimationParameters]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/AnimationParameters
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
// textAnimation: The type of animation Writing Tools is preparing.
//
// range: The range of text that requires a preview image. This range is relative to
// the text in your context object, and it’s your responsibility to match
// that location to the correct location in your text storage. If you
// initialized the context object with the entire contents of your view’s
// text storage, you can use `range` as-is to access that text storage.
// However, if you initialized the context object with only a portion of your
// view’s text, add the starting location of your context object’s text to
// this value to get the correct range for that text storage.
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
// To build a preview of your content in macOS, render the specified range of
// text into an image with a transparent background and use that image to
// create your [NSTextPreview] object directly. Set the [PresentationFrame]
// property to the rectangle in your view’s coordinate space that contains
// the text you captured. Set the [CandidateRects] property to the selection
// rectangles for the text, which you get from your view’s layout manager.
// Writing Tools uses this information to place your image directly above the
// text in your view.
// 
// For a single animation type, the system calls the
// `writingToolsCoordinator(_:)` method, followed sequentially by this method
// and then the
// [WritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion]
// method. Each method executes asynchronously, but the system calls the next
// method in the sequence only after you call the completion handler of the
// previous method. However, multiple animations can run simultaneously, so
// check the `textAnimation` parameter to differentiate sequences.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol/writingToolsCoordinator(_:requestsPreviewFor:of:in:completion:)

func (o NSWritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion(writingToolsCoordinator INSWritingToolsCoordinator, textAnimation NSWritingToolsCoordinatorTextAnimation, range_ foundation.NSRange, context objectivec.IObject, completion ArrayHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("writingToolsCoordinator:requestsPreviewForTextAnimation:ofRange:inContext:completion:"), writingToolsCoordinator, textAnimation, range_, context, completion)
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
// `writingToolsCoordinator(_:)` method, followed sequentially by this method
// and then the
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
// example, it creates an [NSWritingToolsCoordinator.TextAnimation.anticipate]
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
// sequentially by the `writingToolsCoordinator(_:)` and
// [WritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion]
// methods. Each method executes asynchronously, but the system calls the next
// method in the sequence only after you call the completion handler of the
// previous method. However, multiple animations can run simultaneously, so
// check the `textAnimation` and `range` parameters to differentiate
// sequences.
//
// [NSWritingToolsCoordinator.TextAnimation.anticipate]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/TextAnimation/anticipate
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
// `writingToolsCoordinator(_:)` and `writingToolsCoordinator(_:)` methods for
// the same animation type. However, Writing Tools can interleave calls to
// this method with calls to prepare an animation of a different type. In your
// implementation of this method, make sure the actions you take don’t
// interfere with other in-flight animations.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol/writingToolsCoordinator(_:finish:for:in:completion:)

func (o NSWritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion(writingToolsCoordinator INSWritingToolsCoordinator, textAnimation NSWritingToolsCoordinatorTextAnimation, range_ foundation.NSRange, context objectivec.IObject, completion VoidHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("writingToolsCoordinator:finishTextAnimation:forRange:inContext:completion:"), writingToolsCoordinator, textAnimation, range_, context, completion)
	}

// Asks the delegate to provide the bounding paths for the specified text in
// your view.
//
// writingToolsCoordinator: The coordinator object requesting information from your custom view.
//
// range: The range of text to evaluate. This range is relative to the text in your
// context object, and it’s your responsibility to match that location to
// the correct location in your text storage. If you initialized the context
// object with the entire contents of your view’s text storage, you can use
// `range` as-is to access that text storage. However, if you initialized the
// context object with only a portion of your view’s text, add the starting
// location of your context object’s text to this value to get the correct
// range for that text storage.
//
// context: The context object with the target text. Use this object to find the text
// in your view’s text storage.
//
// completion: A handler to execute with the required information. The handler has no
// return value and takes an array of Bezier paths as a parameter. You must
// call this handler at some point during your method’s implementation.
//
// # Discussion
// 
// After applying proofreading marks to your view’s text, Writing Tools lets
// the person accept or reject individual suggestions. To facilitate
// interactions, the coordinator asks your delegate to provide one or more
// Bezier paths that surround those proofreading suggestions. For each
// distinct range of text with a suggestion, it calls this method to get the
// Bezier paths that surround the corresponding text.
// 
// After you determine the location of the specified range of text in your
// view’s text storage, find the rectangle around that text. If you’re
// using TextKit, call the [EnumerateTextSegmentsInRangeTypeOptionsUsingBlock]
// method of your view’s [NSTextLayoutManager] to compute the selection
// rectangles for that text. That method finds the text segments that contain
// the text and returns the frame rectangle for each one. Create a Bezier path
// for each rectangle, and convert the coordinates of each path to the
// coordinate space of the view in your coordinator’s
// [DecorationContainerView] property. Pass the resulting paths to the
// completion handler.
//
// [NSTextLayoutManager]: https://developer.apple.com/documentation/UIKit/NSTextLayoutManager
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol/writingToolsCoordinator(_:requestsBoundingBezierPathsFor:in:completion:)

func (o NSWritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion(writingToolsCoordinator INSWritingToolsCoordinator, range_ foundation.NSRange, context objectivec.IObject, completion ArrayHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("writingToolsCoordinator:requestsBoundingBezierPathsForRange:inContext:completion:"), writingToolsCoordinator, range_, context, completion)
	}

// Asks the delegate to provide an underline shape for the specified text
// during a proofreading session.
//
// writingToolsCoordinator: The coordinator object requesting information from your custom view.
//
// range: The range of text to evaluate. This range is relative to the text in your
// context object, and it’s your responsibility to match that location to
// the correct location in your text storage. If you initialized the context
// object with the entire contents of your view’s text storage, you can use
// `range` as-is to access that text storage. However, if you initialized the
// context object with only a portion of your view’s text, add the starting
// location of your context object’s text to this value to get the correct
// range for that text storage.
//
// context: The context object with the target text. Use this object to find the text
// in your view’s text storage.
//
// completion: A handler to execute with the required information. The handler has no
// return value and takes an array of Bezier paths as a parameter. You must
// call this handler at some point during your method’s implementation.
//
// # Discussion
// 
// When applying proofreading marks to your view’s content, the coordinator
// calls this method to retrieve a shape to draw under the specified text. You
// provide the shape using one or more Bezier paths, and the coordinator draws
// and animates that shape during the proofreading session.
// 
// After you determine the location of the specified range of text in your
// view’s text storage, find the rectangle around that text. If you’re
// using TextKit, you can call the
// [EnumerateTextSegmentsInRangeTypeOptionsUsingBlock] method of your view’s
// [NSTextLayoutManager] to get the rectangles for a range of text. Convert
// the coordinates of each rectangle to the coordinate space of the view in
// your coordinator’s [DecorationContainerView] property. Use those
// rectangles to create the Bezier paths for your text. For example, you might
// create a path with a straight or wavy line at the bottom of the rectangle.
//
// [NSTextLayoutManager]: https://developer.apple.com/documentation/UIKit/NSTextLayoutManager
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol/writingToolsCoordinator(_:requestsUnderlinePathsFor:in:completion:)

func (o NSWritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion(writingToolsCoordinator INSWritingToolsCoordinator, range_ foundation.NSRange, context objectivec.IObject, completion ArrayHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("writingToolsCoordinator:requestsUnderlinePathsForRange:inContext:completion:"), writingToolsCoordinator, range_, context, completion)
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
// [NSWritingToolsCoordinator.State.inactive] state and moves to other states
// as it presents UI and starts interacting with your view’s content. For
// example, it moves to the
// `NSWritingToolsCoordinator/State/interactiveUpdating` state when it’s
// making changes to your view’s text storage.
//
// [NSWritingToolsCoordinator.State.inactive]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/State-swift.enum/inactive
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol/writingToolsCoordinator(_:willChangeTo:completion:)

func (o NSWritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorWillChangeToStateCompletion(writingToolsCoordinator INSWritingToolsCoordinator, newState NSWritingToolsCoordinatorState, completion VoidHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("writingToolsCoordinator:willChangeToState:completion:"), writingToolsCoordinator, newState, completion)
	}

// Asks the delegate to divide the specified range of text into the separate
// containers that render that text.
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
// return value and takes an array of [NSValue] types, each of which contains
// an [NSRange]. The union of the ranges you pass to this handler must equal
// all of the text in `range`. The order of the ranges in the array must be
// sequential, with each new range’s starting location coming after the
// previous one. There must also not be any gaps or overlap between ranges.
// You must call this handler at some point during your implementation.
// //
// [NSRange]: https://developer.apple.com/documentation/Foundation/NSRange-c.struct
// [NSValue]: https://developer.apple.com/documentation/Foundation/NSValue
//
// # Discussion
// 
// If your view uses multiple [NSTextContainer] objects to draw text in
// different regions, use this method to tell Writing Tools about the
// containers that display the specified text. In your implementation,
// subdivide `range` to create one new range for each portion of text that
// resides in a different container object. For example, if the text in
// `range` is split between two containers, provide two new [NSRange] types
// that reflect the portion of the total text in each container. If `range`
// resides completely within one container, call the completion handler with
// `range` as the only value in the array.
// 
// When configuring animations for your view, Writing Tools asks your delegate
// to provide separate previews for each of your view’s container object.
// Specifically, it calls your delegate’s `writingToolsCoordinator(_:)`
// method separately for each range of text you return in the completion
// handler. Your implementation of that method must create a preview suitable
// for animating the content from the underlying text container.
//
// [NSRange]: https://developer.apple.com/documentation/Foundation/NSRange-c.struct
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol/writingToolsCoordinator(_:requestsSingleContainerSubrangesOf:in:completion:)

func (o NSWritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion(writingToolsCoordinator INSWritingToolsCoordinator, range_ foundation.NSRange, context objectivec.IObject, completion ArrayHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("writingToolsCoordinator:requestsSingleContainerSubrangesOfRange:inContext:completion:"), writingToolsCoordinator, range_, context, completion)
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
// `writingToolsCoordinator(_:)` method, Writing Tools calls this method for
// each subrange of text you provided. Find or provide a view situated visibly
// below the specified text in your text view. It’s also satisfactory to
// provide a view that’s visually in front of the text. Writing Tools uses
// the provided view to host any proofreading marks for the specified range of
// text.
// 
// If your view has only one text container, use the coordinator’s
// [DecorationContainerView] property to specify the view to use for
// proofreading marks.
//
// See: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Delegate-swift.protocol/writingToolsCoordinator(_:requestsDecorationContainerViewFor:in:completion:)

func (o NSWritingToolsCoordinatorDelegateObject) WritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion(writingToolsCoordinator INSWritingToolsCoordinator, range_ foundation.NSRange, context objectivec.IObject, completion ViewHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("writingToolsCoordinator:requestsDecorationContainerViewForRange:inContext:completion:"), writingToolsCoordinator, range_, context, completion)
	}





// NSWritingToolsCoordinatorDelegateConfig holds optional typed callbacks for [NSWritingToolsCoordinatorDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nswritingtoolscoordinatordelegate
type NSWritingToolsCoordinatorDelegateConfig struct {

	// Other Methods
	// WritingToolsCoordinatorRequestsContextsForScopeCompletion — Asks your delegate to provide the text to evaluate during the Writing Tools operation.
	WritingToolsCoordinatorRequestsContextsForScopeCompletion func(writingToolsCoordinator NSWritingToolsCoordinator, scope NSWritingToolsCoordinatorContextScope, completion objc.ID)
	// WritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion — Tells the delegate that there are text changes to incorporate into the view.
	WritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion func(writingToolsCoordinator NSWritingToolsCoordinator, range_ foundation.NSRange, context objc.ID, replacementText foundation.NSAttributedString, reason NSWritingToolsCoordinatorTextReplacementReason, animationParameters objc.ID, completion objc.ID)
	// WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion — Asks the delegate for a preview image and layout information for the specified text.
	WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion func(writingToolsCoordinator NSWritingToolsCoordinator, textAnimation NSWritingToolsCoordinatorTextAnimation, range_ foundation.NSRange, context objc.ID, completion objc.ID)
	// WritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion — Asks the delegate to provide the bounding paths for the specified text in your view.
	WritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion func(writingToolsCoordinator NSWritingToolsCoordinator, range_ foundation.NSRange, context objc.ID, completion objc.ID)
	// WritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion — Asks the delegate to provide an underline shape for the specified text during a proofreading session.
	WritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion func(writingToolsCoordinator NSWritingToolsCoordinator, range_ foundation.NSRange, context objc.ID, completion objc.ID)
	// WritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion — Asks the delegate to divide the specified range of text into the separate containers that render that text.
	WritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion func(writingToolsCoordinator NSWritingToolsCoordinator, range_ foundation.NSRange, context objc.ID, completion objc.ID)
	// WritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion — Asks the delegate to provide a decoration view for the specified range of text.
	WritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion func(writingToolsCoordinator NSWritingToolsCoordinator, range_ foundation.NSRange, context objc.ID, completion objc.ID)
}

// NewNSWritingToolsCoordinatorDelegate creates an Objective-C object implementing the [NSWritingToolsCoordinatorDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSWritingToolsCoordinatorDelegateObject] satisfies the [NSWritingToolsCoordinatorDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nswritingtoolscoordinatordelegate
func NewNSWritingToolsCoordinatorDelegate(config NSWritingToolsCoordinatorDelegateConfig) NSWritingToolsCoordinatorDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSWritingToolsCoordinatorDelegate_%d", n)

	var methods []objc.MethodDef

	if config.WritingToolsCoordinatorRequestsContextsForScopeCompletion != nil {
		fn := config.WritingToolsCoordinatorRequestsContextsForScopeCompletion
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("writingToolsCoordinator:requestsContextsForScope:completion:"),
			Fn: func(self objc.ID, _cmd objc.SEL, writingToolsCoordinatorID objc.ID, scope NSWritingToolsCoordinatorContextScope, completionID objc.ID) {
				writingToolsCoordinator := NSWritingToolsCoordinatorFromID(writingToolsCoordinatorID)
				completion := completionID
				fn(writingToolsCoordinator, scope, completion)
			},
		})
	}

	if config.WritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion != nil {
		fn := config.WritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("writingToolsCoordinator:replaceRange:inContext:proposedText:reason:animationParameters:completion:"),
			Fn: func(self objc.ID, _cmd objc.SEL, writingToolsCoordinatorID objc.ID, range_ foundation.NSRange, contextID objc.ID, replacementTextID objc.ID, reason NSWritingToolsCoordinatorTextReplacementReason, animationParametersID objc.ID, completionID objc.ID) {
				writingToolsCoordinator := NSWritingToolsCoordinatorFromID(writingToolsCoordinatorID)
				context := contextID
				replacementText := foundation.NSAttributedStringFromID(replacementTextID)
				animationParameters := animationParametersID
				completion := completionID
				fn(writingToolsCoordinator, range_, context, replacementText, reason, animationParameters, completion)
			},
		})
	}

	if config.WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion != nil {
		fn := config.WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("writingToolsCoordinator:requestsPreviewForTextAnimation:ofRange:inContext:completion:"),
			Fn: func(self objc.ID, _cmd objc.SEL, writingToolsCoordinatorID objc.ID, textAnimation NSWritingToolsCoordinatorTextAnimation, range_ foundation.NSRange, contextID objc.ID, completionID objc.ID) {
				writingToolsCoordinator := NSWritingToolsCoordinatorFromID(writingToolsCoordinatorID)
				context := contextID
				completion := completionID
				fn(writingToolsCoordinator, textAnimation, range_, context, completion)
			},
		})
	}

	if config.WritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion != nil {
		fn := config.WritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("writingToolsCoordinator:requestsBoundingBezierPathsForRange:inContext:completion:"),
			Fn: func(self objc.ID, _cmd objc.SEL, writingToolsCoordinatorID objc.ID, range_ foundation.NSRange, contextID objc.ID, completionID objc.ID) {
				writingToolsCoordinator := NSWritingToolsCoordinatorFromID(writingToolsCoordinatorID)
				context := contextID
				completion := completionID
				fn(writingToolsCoordinator, range_, context, completion)
			},
		})
	}

	if config.WritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion != nil {
		fn := config.WritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("writingToolsCoordinator:requestsUnderlinePathsForRange:inContext:completion:"),
			Fn: func(self objc.ID, _cmd objc.SEL, writingToolsCoordinatorID objc.ID, range_ foundation.NSRange, contextID objc.ID, completionID objc.ID) {
				writingToolsCoordinator := NSWritingToolsCoordinatorFromID(writingToolsCoordinatorID)
				context := contextID
				completion := completionID
				fn(writingToolsCoordinator, range_, context, completion)
			},
		})
	}

	if config.WritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion != nil {
		fn := config.WritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("writingToolsCoordinator:requestsSingleContainerSubrangesOfRange:inContext:completion:"),
			Fn: func(self objc.ID, _cmd objc.SEL, writingToolsCoordinatorID objc.ID, range_ foundation.NSRange, contextID objc.ID, completionID objc.ID) {
				writingToolsCoordinator := NSWritingToolsCoordinatorFromID(writingToolsCoordinatorID)
				context := contextID
				completion := completionID
				fn(writingToolsCoordinator, range_, context, completion)
			},
		})
	}

	if config.WritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion != nil {
		fn := config.WritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("writingToolsCoordinator:requestsDecorationContainerViewForRange:inContext:completion:"),
			Fn: func(self objc.ID, _cmd objc.SEL, writingToolsCoordinatorID objc.ID, range_ foundation.NSRange, contextID objc.ID, completionID objc.ID) {
				writingToolsCoordinator := NSWritingToolsCoordinatorFromID(writingToolsCoordinatorID)
				context := contextID
				completion := completionID
				fn(writingToolsCoordinator, range_, context, completion)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSWritingToolsCoordinatorDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSWritingToolsCoordinatorDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSWritingToolsCoordinatorDelegateObjectFromID(instance)
}





