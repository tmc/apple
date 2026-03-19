// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// The optional methods that delegates of text storage objects implement to handle text-edit processing.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorageDelegate
type NSTextStorageDelegate interface {
	objectivec.IObject
}

// NSTextStorageDelegateObject wraps an existing Objective-C object that conforms to the NSTextStorageDelegate protocol.
type NSTextStorageDelegateObject struct {
	objectivec.Object
}
func (o NSTextStorageDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSTextStorageDelegateObjectFromID constructs a [NSTextStorageDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTextStorageDelegateObjectFromID(id objc.ID) NSTextStorageDelegateObject {
	return NSTextStorageDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The method the framework calls when a text storage object is about to
// process edits.
//
// textStorage: The text storage object processing edits.
//
// editedMask: The types of edits to do: [TextStorageEditedAttributes]
// [TextStorageEditedCharacters], or both.
//
// editedRange: The range in the original string (before the edit).
//
// delta: The length delta for the editing changes.
//
// # Discussion
// 
// Sent inside [ProcessEditing] right before fixing attributes. Delegates can
// change the characters or attributes.
// 
// The delegate can verify the changed state of the text storage object and
// make changes to the text storage object’s characters or attributes to
// enforce whatever constraints it establishes. Programmatic changes don’t
// cause the object to send this message.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorageDelegate/textStorage(_:willProcessEditing:range:changeInLength:)
func (o NSTextStorageDelegateObject) TextStorageWillProcessEditingRangeChangeInLength(textStorage NSTextStorage, editedMask NSTextStorageEditActions, editedRange foundation.NSRange, delta int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("textStorage:willProcessEditing:range:changeInLength:"), textStorage, editedMask, editedRange, delta)
	}
// The method the framework calls when a text storage object has finished
// processing edits.
//
// textStorage: The text storage object processing edits.
//
// editedMask: The types of edits done: [TextStorageEditedAttributes],
// [TextStorageEditedCharacters], or both.
//
// editedRange: The range in the original string (before the edit).
//
// delta: The length delta for the editing changes.
//
// # Discussion
// 
// Sent inside [ProcessEditing] right before notifying layout managers.
// Delegates can change the attributes.
// 
// The delegate can verify the final state of the text storage object; it
// can’t change the text storage object’s characters without leaving it in
// an inconsistent state, but if necessary it can change attributes. Note that
// even in this case it’s possible to put a text storage object into an
// inconsistent state—for example, by changing the font of a range to one
// that doesn’t support the characters in that range, such as using a Latin
// font for Kanji text.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorageDelegate/textStorage(_:didProcessEditing:range:changeInLength:)
func (o NSTextStorageDelegateObject) TextStorageDidProcessEditingRangeChangeInLength(textStorage NSTextStorage, editedMask NSTextStorageEditActions, editedRange foundation.NSRange, delta int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("textStorage:didProcessEditing:range:changeInLength:"), textStorage, editedMask, editedRange, delta)
	}

// NSTextStorageDelegateConfig holds optional typed callbacks for [NSTextStorageDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nstextstoragedelegate
type NSTextStorageDelegateConfig struct {

	// Processing edit actions
	// TextStorageWillProcessEditingRangeChangeInLength — The method the framework calls when a text storage object is about to process edits.
	TextStorageWillProcessEditingRangeChangeInLength func(textStorage NSTextStorage, editedMask NSTextStorageEditActions, editedRange foundation.NSRange, delta int)
	// TextStorageDidProcessEditingRangeChangeInLength — The method the framework calls when a text storage object has finished processing edits.
	TextStorageDidProcessEditingRangeChangeInLength func(textStorage NSTextStorage, editedMask NSTextStorageEditActions, editedRange foundation.NSRange, delta int)
}

// NewNSTextStorageDelegate creates an Objective-C object implementing the [NSTextStorageDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSTextStorageDelegateObject] satisfies the [NSTextStorageDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nstextstoragedelegate
func NewNSTextStorageDelegate(config NSTextStorageDelegateConfig) NSTextStorageDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSTextStorageDelegate_%d", n)

	var methods []objc.MethodDef

	if config.TextStorageWillProcessEditingRangeChangeInLength != nil {
		fn := config.TextStorageWillProcessEditingRangeChangeInLength
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textStorage:willProcessEditing:range:changeInLength:"),
			Fn: func(self objc.ID, _cmd objc.SEL, textStorageID objc.ID, editedMask NSTextStorageEditActions, editedRange foundation.NSRange, delta int) {
				textStorage := NSTextStorageFromID(textStorageID)
				fn(textStorage, editedMask, editedRange, delta)
			},
		})
	}

	if config.TextStorageDidProcessEditingRangeChangeInLength != nil {
		fn := config.TextStorageDidProcessEditingRangeChangeInLength
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textStorage:didProcessEditing:range:changeInLength:"),
			Fn: func(self objc.ID, _cmd objc.SEL, textStorageID objc.ID, editedMask NSTextStorageEditActions, editedRange foundation.NSRange, delta int) {
				textStorage := NSTextStorageFromID(textStorageID)
				fn(textStorage, editedMask, editedRange, delta)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSTextStorageDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSTextStorageDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSTextStorageDelegateObjectFromID(instance)
}

