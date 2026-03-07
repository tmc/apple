// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The [NSMatrixDelegate] protocol defines the optional methods implemented by delegates of [NSMatrix] objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrixDelegate
type NSMatrixDelegate interface {
	objectivec.IObject
	NSControlTextEditingDelegate
}



// NSMatrixDelegateObject wraps an existing Objective-C object that conforms to the NSMatrixDelegate protocol.
type NSMatrixDelegateObject struct {
	objectivec.Object
}
func (o NSMatrixDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSMatrixDelegateObjectFromID constructs a [NSMatrixDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSMatrixDelegateObjectFromID(id objc.ID) NSMatrixDelegateObject {
	return NSMatrixDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
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

func (o NSMatrixDelegateObject) ControlIsValidObject(control INSControl, obj objectivec.IObject) bool {
	
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

func (o NSMatrixDelegateObject) ControlDidFailToValidatePartialStringErrorDescription(control INSControl, string_ string, error_ string) {
	
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

func (o NSMatrixDelegateObject) ControlDidFailToFormatStringErrorDescription(control INSControl, string_ string, error_ string) bool {
	
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

func (o NSMatrixDelegateObject) ControlTextShouldBeginEditing(control INSControl, fieldEditor INSText) bool {
	
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

func (o NSMatrixDelegateObject) ControlTextShouldEndEditing(control INSControl, fieldEditor INSText) bool {
	
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

func (o NSMatrixDelegateObject) ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem(control INSControl, textView INSTextView, words []string, charRange foundation.NSRange, index int) []string {
	
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

func (o NSMatrixDelegateObject) ControlTextViewDoCommandBySelector(control INSControl, textView INSTextView, commandSelector objc.SEL) bool {
	
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

func (o NSMatrixDelegateObject) ControlTextDidBeginEditing(obj foundation.NSNotification) {
	
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

func (o NSMatrixDelegateObject) ControlTextDidChange(obj foundation.NSNotification) {
	
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

func (o NSMatrixDelegateObject) ControlTextDidEndEditing(obj foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("controlTextDidEndEditing:"), obj)
	}







