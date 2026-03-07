// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf


// A set of optional methods implemented by delegates of [NSTokenField](<doc://com.apple.appkit/documentation/AppKit/NSTokenField>) objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldDelegate
type NSTokenFieldDelegate interface {
	objectivec.IObject
	NSControlTextEditingDelegate
	NSTextFieldDelegate
}



// NSTokenFieldDelegateObject wraps an existing Objective-C object that conforms to the NSTokenFieldDelegate protocol.
type NSTokenFieldDelegateObject struct {
	objectivec.Object
}
func (o NSTokenFieldDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSTokenFieldDelegateObjectFromID constructs a [NSTokenFieldDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTokenFieldDelegateObjectFromID(id objc.ID) NSTokenFieldDelegateObject {
	return NSTokenFieldDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Allows the delegate to provide a string to be displayed as a proxy for the
// given represented object.
//
// tokenField: The token field that sent the message.
//
// representedObject: A represented object of the token field.
//
// # Return Value
// 
// The string to be used as a proxy for `representedObject`. If you return
// `nil` or do not implement this method, then `representedObject` is
// displayed as the string.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldDelegate/tokenField(_:displayStringForRepresentedObject:)

func (o NSTokenFieldDelegateObject) TokenFieldDisplayStringForRepresentedObject(tokenField INSTokenField, representedObject objectivec.IObject) string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("tokenField:displayStringForRepresentedObject:"), tokenField, representedObject)
	return foundation.NSStringFromID(rv).String()
	}

// Allows the delegate to return the token style for editing the specified
// represented object.
//
// tokenField: The token field that sent the message.
//
// representedObject: A represented object of the token field.
//
// # Return Value
// 
// The style that should be used to display the representedObject. Possible
// values are shown in NSTokenStyle Values.
//
// # Discussion
// 
// If the delegate implements this method and returns an
// [NSTokenField.TokenStyle] that differs from the style set by [TokenStyle],
// the value the delegate returns is preferred.
// 
// If the delegate does not implement this method, the token field’s
// [TokenStyle] is used.
//
// [NSTokenField.TokenStyle]: https://developer.apple.com/documentation/AppKit/NSTokenField/TokenStyle-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldDelegate/tokenField(_:styleForRepresentedObject:)

func (o NSTokenFieldDelegateObject) TokenFieldStyleForRepresentedObject(tokenField INSTokenField, representedObject objectivec.IObject) NSTokenStyle {
	
	rv := objc.Send[NSTokenStyle](o.ID, objc.Sel("tokenField:styleForRepresentedObject:"), tokenField, representedObject)
	return rv
	}

// Allows the delegate to provide an array of appropriate completions for the
// contents of the receiver.
//
// tokenField: The token field where editing is occurring.
//
// substring: The partial string that is to be completed.
//
// tokenIndex: The index of the token being edited.
//
// selectedIndex: Optionally, you can return by-reference an index into the returned array
// that specifies which of the completions should be initially selected. If
// none are to be selected, return by reference `-1`.
//
// # Return Value
// 
// An array of strings that are possible completions.
//
// # Discussion
// 
// If the delegate does not implement this method, no completions are
// provided.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldDelegate/tokenField(_:completionsForSubstring:indexOfToken:indexOfSelectedItem:)

func (o NSTokenFieldDelegateObject) TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem(tokenField INSTokenField, substring string, tokenIndex int, selectedIndex int) foundation.INSArray {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("tokenField:completionsForSubstring:indexOfToken:indexOfSelectedItem:"), tokenField, objc.String(substring), tokenIndex, selectedIndex)
	return foundation.NSArrayFromID(rv)
	}

// Allows the delegate to provide a string to be edited as a proxy for a
// represented object.
//
// tokenField: The token field that sent the message.
//
// representedObject: A represented object of the token field.
//
// # Return Value
// 
// A string that’s an editable proxy of the represented object, or `nil` if
// the token should not be editable.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldDelegate/tokenField(_:editingStringForRepresentedObject:)

func (o NSTokenFieldDelegateObject) TokenFieldEditingStringForRepresentedObject(tokenField INSTokenField, representedObject objectivec.IObject) string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("tokenField:editingStringForRepresentedObject:"), tokenField, representedObject)
	return foundation.NSStringFromID(rv).String()
	}

// Allows the delegate to provide a represented object for the given editing
// string.
//
// tokenField: The token field that sent the message.
//
// editingString: The edited string representation of a represented object.
//
// # Return Value
// 
// A represented object that is displayed rather than the editing string.
//
// # Discussion
// 
// If your application uses some object other than an [NSString] for their
// represented objects, you should return a new, autoreleased instance of that
// object from this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldDelegate/tokenField(_:representedObjectForEditing:)

func (o NSTokenFieldDelegateObject) TokenFieldRepresentedObjectForEditingString(tokenField INSTokenField, editingString string) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("tokenField:representedObjectForEditingString:"), tokenField, objc.String(editingString))
	return objectivec.Object{ID: rv}
	}

// Allows the delegate to validate the tokens to be added to the receiver at a
// particular location.
//
// tokenField: The token field that sent the message.
//
// tokens: An array of tokens to be inserted in the receiver at `index`.
//
// index: The index of the receiver in which the array of tokens to be validated
// (`tokens`) will be inserted.
//
// # Return Value
// 
// An array of validated tokens.
//
// # Discussion
// 
// The delegate can return the array unchanged or return a modified array of
// tokens. To reject the add completely, return an empty array. Returning
// `nil` causes an error.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldDelegate/tokenField(_:shouldAdd:at:)

func (o NSTokenFieldDelegateObject) TokenFieldShouldAddObjectsAtIndex(tokenField INSTokenField, tokens foundation.INSArray, index uint) foundation.INSArray {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("tokenField:shouldAddObjects:atIndex:"), tokenField, tokens, index)
	return foundation.NSArrayFromID(rv)
	}

// Allows the delegate to return an array of objects representing the data
// read from the specified pasteboard.
//
// tokenField: The token field that sent the message.
//
// pboard: The pasteboard from which to read the represented objects.
//
// # Return Value
// 
// An array of represented objects created from the pasteboard data.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldDelegate/tokenField(_:readFrom:)

func (o NSTokenFieldDelegateObject) TokenFieldReadFromPasteboard(tokenField INSTokenField, pboard INSPasteboard) foundation.INSArray {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("tokenField:readFromPasteboard:"), tokenField, pboard)
	return foundation.NSArrayFromID(rv)
	}

// Sent so the delegate can write represented objects to the pasteboard
// corresponding to a given array of display strings.
//
// tokenField: The token field that sent the message.
//
// objects: An array of represented objects associated with the token field.
//
// pboard: The pasteboard to which to write the represented objects.
//
// # Return Value
// 
// [true] if the delegate writes the represented objects to the pasteboard,
// [false] otherwise. If [false], the token field writes the display strings
// to the [NSStringPboardType] pasteboard.
//
// [NSStringPboardType]: https://developer.apple.com/documentation/AppKit/NSStringPboardType
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldDelegate/tokenField(_:writeRepresentedObjects:to:)

func (o NSTokenFieldDelegateObject) TokenFieldWriteRepresentedObjectsToPasteboard(tokenField INSTokenField, objects foundation.INSArray, pboard INSPasteboard) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("tokenField:writeRepresentedObjects:toPasteboard:"), tokenField, objects, pboard)
	return rv
	}

// Allows the delegate to specify whether the given represented object
// provides a menu.
//
// tokenField: The token field that sent the message.
//
// representedObject: A represented object of the token field.
//
// # Return Value
// 
// [true] if the represented object has a menu, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// By default tokens in a token field have no menus.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldDelegate/tokenField(_:hasMenuForRepresentedObject:)

func (o NSTokenFieldDelegateObject) TokenFieldHasMenuForRepresentedObject(tokenField INSTokenField, representedObject objectivec.IObject) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("tokenField:hasMenuForRepresentedObject:"), tokenField, representedObject)
	return rv
	}

// Allows the delegate to provide a menu for the specified represented object.
//
// tokenField: The token field that sent the message.
//
// representedObject: A represented object of the token field.
//
// # Return Value
// 
// The menu associated with the represented object.
//
// # Discussion
// 
// The returned menu should be autoreleased. By default tokens in a token
// field do not return menus.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldDelegate/tokenField(_:menuForRepresentedObject:)

func (o NSTokenFieldDelegateObject) TokenFieldMenuForRepresentedObject(tokenField INSTokenField, representedObject objectivec.IObject) INSMenu {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("tokenField:menuForRepresentedObject:"), tokenField, representedObject)
	return NSMenuFromID(rv)
	}

// Invoked when the insertion point leaves a cell belonging to the specified
// control, but before the value of the cell’s object is displayed.
//
// control: The control whose object value needs to be validated.
//
// obj: The object value to validate.
//
// # Return Value
// 
// [true] if you want to allow the control to display the specified value;
// otherwise, [false] to reject the value and return the cursor to the
// control’s cell.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method gives the delegate the opportunity to validate the contents of
// the control’s cell (or selected cell). In validating, the delegate should
// check the value in the `object` parameter and determine if it falls within
// a permissible range, has required attributes, accords with a given context,
// and so on. Examples of objects subject to such evaluations are an [NSDate]
// object that should not represent a future date or a monetary amount
// (represented by an [NSNumber] object) that exceeds a predetermined limit.
//
// See: https://developer.apple.com/documentation/AppKit/NSControlTextEditingDelegate/control(_:isValidObject:)

func (o NSTokenFieldDelegateObject) ControlIsValidObject(control INSControl, obj objectivec.IObject) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("control:isValidObject:"), control, obj)
	return rv
	}

// Invoked when the formatter for the cell belonging to `control` (or selected
// cell) rejects a partial string a user is typing into the cell.
//
// control: The control whose cell rejected the string.
//
// string: The string that includes the character that caused the rejection.
//
// error: A localized, user-presentable string that explains why the string was
// rejected.
//
// # Discussion
// 
// You can implement this method to display a warning message or perform a
// similar action when the user enters improperly formatted text.
//
// See: https://developer.apple.com/documentation/AppKit/NSControlTextEditingDelegate/control(_:didFailToValidatePartialString:errorDescription:)

func (o NSTokenFieldDelegateObject) ControlDidFailToValidatePartialStringErrorDescription(control INSControl, string_ string, error_ string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("control:didFailToValidatePartialString:errorDescription:"), control, objc.String(string_), objc.String(error_))
	}

// Invoked when the formatter for the cell belonging to the specified control
// cannot convert a string to an underlying object.
//
// control: The control whose cell could not convert the string.
//
// string: The string that could not be converted.
//
// error: A localized, user-presentable string that explains why the conversion
// failed.
//
// # Return Value
// 
// [true] if the value in the string parameter should be accepted as is;
// otherwise, [false] if the value in the parameter should be rejected.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Your implementation of this method should evaluate the error or query the
// user an appropriate value indicating whether the string should be accepted
// or rejected.
//
// See: https://developer.apple.com/documentation/AppKit/NSControlTextEditingDelegate/control(_:didFailToFormatString:errorDescription:)

func (o NSTokenFieldDelegateObject) ControlDidFailToFormatStringErrorDescription(control INSControl, string_ string, error_ string) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("control:didFailToFormatString:errorDescription:"), control, objc.String(string_), objc.String(error_))
	return rv
	}

// Invoked when the user tries to enter a character in a cell of a control
// that allows editing of text (such as a text field or form field).
//
// control: The control whose content is about to be edited.
//
// fieldEditor: The field editor of the control.
//
// # Return Value
// 
// [true] if the control’s field editor should be allowed to start editing
// the text; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// You can use this method to allow or disallow editing in a control. This
// message is sent by the control directly to its delegate object.
//
// See: https://developer.apple.com/documentation/AppKit/NSControlTextEditingDelegate/control(_:textShouldBeginEditing:)

func (o NSTokenFieldDelegateObject) ControlTextShouldBeginEditing(control INSControl, fieldEditor INSText) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("control:textShouldBeginEditing:"), control, fieldEditor)
	return rv
	}

// Invoked when the insertion point tries to leave a cell of the control that
// has been edited.
//
// control: The control for which editing is about to end.
//
// fieldEditor: The field editor of the control. You can use this parameter to get the
// edited text.
//
// # Return Value
// 
// [true] if the insertion point should be allowed to end the editing session;
// otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This message is sent only by controls that allow editing of text (such as a
// text field or a form field). This message is sent by the control directly
// to its delegate object.
//
// See: https://developer.apple.com/documentation/AppKit/NSControlTextEditingDelegate/control(_:textShouldEndEditing:)

func (o NSTokenFieldDelegateObject) ControlTextShouldEndEditing(control INSControl, fieldEditor INSText) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("control:textShouldEndEditing:"), control, fieldEditor)
	return rv
	}

// Invoked to allow you to control the list of proposed text completions
// generated by text fields and other controls.
//
// control: The control whose cell initiated the message. If the control contains
// multiple cells, the one that initiated the message is usually the selected
// cell.
//
// textView: The field editor of the control. You can use this parameter to get the
// typed text.
//
// words: An array of [NSString] objects containing the potential completions. The
// completion strings are listed in their order of preference in the array.
//
// charRange: The range of characters the user has already typed.
//
// index: On input, an integer variable with the default value of `0`. On output, you
// can set this value to an index in the returned array indicating which item
// should be selected initially. Set the value to `-1` to indicate there
// should not be an initial selection.
//
// # Return Value
// 
// An array of [NSString] objects containing the list of completions to use in
// place of the array in the `words` parameter. The returned array should list
// the completions in their preferred order
//
// # Discussion
// 
// Each string you return should be a complete word that the user might be
// trying to type. The strings must be complete words rather than just the
// remainder of the word, in case completion requires some slight modification
// of what the user has already typed—for example, the addition of an
// accent, or a change in capitalization. You can also use this method to
// support abbreviations that complete into words that do not start with the
// characters of the abbreviation. The `index` argument allows you to return
// by reference an index specifying which of the completions should be
// selected initially.
// 
// The actual means of presentation of the potential completions is determined
// by the [Complete] method of [NSTextView].
//
// See: https://developer.apple.com/documentation/AppKit/NSControlTextEditingDelegate/control(_:textView:completions:forPartialWordRange:indexOfSelectedItem:)

func (o NSTokenFieldDelegateObject) ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem(control INSControl, textView INSTextView, words []string, charRange foundation.NSRange, index int) []string {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("control:textView:completions:forPartialWordRange:indexOfSelectedItem:"), control, textView, objectivec.StringSliceToNSArray(words), charRange, index)
	return objc.ConvertSliceToStrings(rv)
	}

// Invoked when users press keys with predefined bindings in a cell of the
// specified control.
//
// control: The control whose cell initiated the message. If the control contains
// multiple cells, the one that initiated the message is usually the selected
// cell.
//
// textView: The field editor of the control.
//
// commandSelector: The selector that was associated with the binding.
//
// # Return Value
// 
// [true] if the delegate object handles the key binding; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// These bindings are usually implemented as methods (`command`) defined in
// the [NSResponder] class; examples of such key bindings are arrow keys (for
// directional movement) and the Escape key (for name completion). By
// implementing this method, the delegate can override the default
// implementation of `command` and supply its own behavior.
// 
// For example, the default method for completing partially typed pathnames or
// symbols (usually when users press the Escape key) is `complete(_:)`. The
// default implementation of the `complete(_:)` method (in [NSResponder]) does
// nothing. The delegate could evaluate the method in the `command` parameter
// and, if it’s `complete(_:)`, get the current string from the `textView`
// parameter and then expand it, or display a list of potential completions,
// or do whatever else is appropriate.
//
// See: https://developer.apple.com/documentation/AppKit/NSControlTextEditingDelegate/control(_:textView:doCommandBy:)

func (o NSTokenFieldDelegateObject) ControlTextViewDoCommandBySelector(control INSControl, textView INSTextView, commandSelector objc.SEL) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("control:textView:doCommandBySelector:"), control, textView, commandSelector)
	return rv
	}

// Tells the delegate that the control started editing its text content.
//
// obj: A notification object that contains details about the editing
// configuration.
//
// # Discussion
// 
// Use the key `“NSFieldEditor”` to obtain the field editor from the
// notification object’s `userInfo` dictionary.
//
// See: https://developer.apple.com/documentation/AppKit/NSControlTextEditingDelegate/controlTextDidBeginEditing(_:)

func (o NSTokenFieldDelegateObject) ControlTextDidBeginEditing(obj foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("controlTextDidBeginEditing:"), obj)
	}

// Tells the delegate that the control made changes to its text content.
//
// obj: A notification object that contains details about the editing
// configuration.
//
// # Discussion
// 
// Use the key `“NSFieldEditor”` to obtain the field editor from the
// notification object’s `userInfo` dictionary.
//
// See: https://developer.apple.com/documentation/AppKit/NSControlTextEditingDelegate/controlTextDidChange(_:)

func (o NSTokenFieldDelegateObject) ControlTextDidChange(obj foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("controlTextDidChange:"), obj)
	}

// Tells the delegate that the control finished editing its text content and
// committed the changes.
//
// obj: A notification object that contains details about the editing
// configuration.
//
// # Discussion
// 
// Use the key `“NSFieldEditor”` to obtain the field editor from the
// notification object’s `userInfo` dictionary.
//
// See: https://developer.apple.com/documentation/AppKit/NSControlTextEditingDelegate/controlTextDidEndEditing(_:)

func (o NSTokenFieldDelegateObject) ControlTextDidEndEditing(obj foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("controlTextDidEndEditing:"), obj)
	}

//
// See: https://developer.apple.com/documentation/AppKit/NSTextFieldDelegate/textField(_:textView:candidatesForSelectedRange:)

func (o NSTokenFieldDelegateObject) TextFieldTextViewCandidatesForSelectedRange(textField INSTextField, textView INSTextView, selectedRange foundation.NSRange) foundation.INSArray {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("textField:textView:candidatesForSelectedRange:"), textField, textView, selectedRange)
	return foundation.NSArrayFromID(rv)
	}

//
// See: https://developer.apple.com/documentation/AppKit/NSTextFieldDelegate/textField(_:textView:shouldSelectCandidateAt:)

func (o NSTokenFieldDelegateObject) TextFieldTextViewShouldSelectCandidateAtIndex(textField INSTextField, textView INSTextView, index uint) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("textField:textView:shouldSelectCandidateAtIndex:"), textField, textView, index)
	return rv
	}





// NSTokenFieldDelegateConfig holds optional typed callbacks for [NSTokenFieldDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nstokenfielddelegate
type NSTokenFieldDelegateConfig struct {

	// Other Methods
	// ShouldAddObjectsAtIndex — Allows the delegate to validate the tokens to be added to the receiver at a particular location.
	ShouldAddObjectsAtIndex func(tokenField NSTokenField, tokens foundation.INSArray, index uint) foundation.INSArray
	// ReadFromPasteboard — Allows the delegate to return an array of objects representing the data read from the specified pasteboard.
	ReadFromPasteboard func(tokenField NSTokenField, pboard NSPasteboard) foundation.INSArray
	// WriteRepresentedObjectsToPasteboard — Sent so the delegate can write represented objects to the pasteboard corresponding to a given array of display strings.
	WriteRepresentedObjectsToPasteboard func(tokenField NSTokenField, objects foundation.INSArray, pboard NSPasteboard) bool
}

// NewNSTokenFieldDelegate creates an Objective-C object implementing the [NSTokenFieldDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSTokenFieldDelegateObject] satisfies the [NSTokenFieldDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nstokenfielddelegate
func NewNSTokenFieldDelegate(config NSTokenFieldDelegateConfig) NSTokenFieldDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSTokenFieldDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ShouldAddObjectsAtIndex != nil {
		fn := config.ShouldAddObjectsAtIndex
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tokenField:shouldAddObjects:atIndex:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tokenFieldID objc.ID, tokensID objc.ID, index uint) objc.ID {
				tokenField := NSTokenFieldFromID(tokenFieldID)
				tokens := foundation.NSArrayFromID(tokensID)
				return fn(tokenField, tokens, index).GetID()
			},
		})
	}

	if config.ReadFromPasteboard != nil {
		fn := config.ReadFromPasteboard
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tokenField:readFromPasteboard:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tokenFieldID objc.ID, pboardID objc.ID) objc.ID {
				tokenField := NSTokenFieldFromID(tokenFieldID)
				pboard := NSPasteboardFromID(pboardID)
				return fn(tokenField, pboard).GetID()
			},
		})
	}

	if config.WriteRepresentedObjectsToPasteboard != nil {
		fn := config.WriteRepresentedObjectsToPasteboard
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tokenField:writeRepresentedObjects:toPasteboard:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tokenFieldID objc.ID, objectsID objc.ID, pboardID objc.ID) bool {
				tokenField := NSTokenFieldFromID(tokenFieldID)
				objects := foundation.NSArrayFromID(objectsID)
				pboard := NSPasteboardFromID(pboardID)
				return fn(tokenField, objects, pboard)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSTokenFieldDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSTokenFieldDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSTokenFieldDelegateObjectFromID(instance)
}





