// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Formatter] class.
var (
	_FormatterClass     FormatterClass
	_FormatterClassOnce sync.Once
)

func getFormatterClass() FormatterClass {
	_FormatterClassOnce.Do(func() {
		_FormatterClass = FormatterClass{class: objc.GetClass("NSFormatter")}
	})
	return _FormatterClass
}

// GetFormatterClass returns the class object for NSFormatter.
func GetFormatterClass() FormatterClass {
	return getFormatterClass()
}

type FormatterClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (fc FormatterClass) Alloc() Formatter {
	rv := objc.Send[Formatter](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}







// An abstract class that declares an interface for objects that create,
// interpret, and validate the textual representation of values.
//
// # Overview
// 
// The Foundation framework provides several concrete subclasses of
// [NSFormatter], including [NSByteCountFormatter], [NSDateFormatter],
// [NSDateComponentsFormatter], [NSDateIntervalFormatter],
// [NSMeasurementFormatter], [NSNumberFormatter], and
// [NSPersonNameComponentsFormatter].
// 
// # Subclassing Notes
// 
// [NSFormatter] is intended for subclassing. A custom formatter can restrict
// the input and enhance the display of data in novel ways. For example, you
// could have a custom formatter that ensures that serial numbers entered by a
// user conform to predefined formats. Before you decide to create a custom
// formatter, make sure that you cannot configure the public subclasses to
// satisfy your requirements.
// 
// For instructions on how to create your own custom formatter, see [Creating
// a Custom Formatter].
//
// [Creating a Custom Formatter]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DataFormatting/Articles/CreatingACustomFormatter.html#//apple_ref/doc/uid/20000196
//
// # Getting Textual Representations of Object Values
//
//   - [Formatter.StringForObjectValue]: The default implementation of this method raises an exception.
//   - [Formatter.AttributedStringForObjectValueWithDefaultAttributes]: The default implementation returns `nil` to indicate that the formatter object does not provide an attributed string.
//   - [Formatter.EditingStringForObjectValue]: The default implementation of this method invokes [string(for:)](<doc://com.apple.foundation/documentation/Foundation/Formatter/string(for:)>).
//
// # Getting Object Values for Textual Representations
//
//   - [Formatter.GetObjectValueForStringErrorDescription]: The default implementation of this method raises an exception.
//
// # Validating Partial Strings
//
//   - [Formatter.IsPartialStringValidNewEditingStringErrorDescription]: Returns a Boolean value that indicates whether a partial string is valid.
//   - [Formatter.IsPartialStringValidProposedSelectedRangeOriginalStringOriginalSelectedRangeErrorDescription]: This method should be implemented in subclasses that want to validate user changes to a string in a field, where the user changes are not necessarily at the end of the string, and preserve the selection (or set a different one, such as selecting the erroneous part of the string the user has typed).
//
// See: https://developer.apple.com/documentation/Foundation/Formatter
type Formatter struct {
	objectivec.Object
}

// FormatterFromID constructs a [Formatter] from an objc.ID.
//
// An abstract class that declares an interface for objects that create,
// interpret, and validate the textual representation of values.
func FormatterFromID(id objc.ID) Formatter {
	return NSFormatter{objectivec.Object{id}}
}

// NSFormatterFromID is an alias for [FormatterFromID] for cross-framework compatibility.
func NSFormatterFromID(id objc.ID) Formatter { return FormatterFromID(id) }
// NOTE: Formatter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [Formatter] class.
//
// # Getting Textual Representations of Object Values
//
//   - [IFormatter.StringForObjectValue]: The default implementation of this method raises an exception.
//   - [IFormatter.AttributedStringForObjectValueWithDefaultAttributes]: The default implementation returns `nil` to indicate that the formatter object does not provide an attributed string.
//   - [IFormatter.EditingStringForObjectValue]: The default implementation of this method invokes [string(for:)](<doc://com.apple.foundation/documentation/Foundation/Formatter/string(for:)>).
//
// # Getting Object Values for Textual Representations
//
//   - [IFormatter.GetObjectValueForStringErrorDescription]: The default implementation of this method raises an exception.
//
// # Validating Partial Strings
//
//   - [IFormatter.IsPartialStringValidNewEditingStringErrorDescription]: Returns a Boolean value that indicates whether a partial string is valid.
//   - [IFormatter.IsPartialStringValidProposedSelectedRangeOriginalStringOriginalSelectedRangeErrorDescription]: This method should be implemented in subclasses that want to validate user changes to a string in a field, where the user changes are not necessarily at the end of the string, and preserve the selection (or set a different one, such as selecting the erroneous part of the string the user has typed).
//
// See: https://developer.apple.com/documentation/Foundation/Formatter
type IFormatter interface {
	objectivec.IObject
	NSCopying

	// Topic: Getting Textual Representations of Object Values

	// The default implementation of this method raises an exception.
	StringForObjectValue(obj objectivec.IObject) string
	// The default implementation returns `nil` to indicate that the formatter object does not provide an attributed string.
	AttributedStringForObjectValueWithDefaultAttributes(obj objectivec.IObject, attrs INSDictionary) INSAttributedString
	// The default implementation of this method invokes [string(for:)](<doc://com.apple.foundation/documentation/Foundation/Formatter/string(for:)>).
	EditingStringForObjectValue(obj objectivec.IObject) string

	// Topic: Getting Object Values for Textual Representations

	// The default implementation of this method raises an exception.
	GetObjectValueForStringErrorDescription(obj []objectivec.IObject, string_ string, error_ string) bool

	// Topic: Validating Partial Strings

	// Returns a Boolean value that indicates whether a partial string is valid.
	IsPartialStringValidNewEditingStringErrorDescription(partialString string, newString string, error_ string) bool
	// This method should be implemented in subclasses that want to validate user changes to a string in a field, where the user changes are not necessarily at the end of the string, and preserve the selection (or set a different one, such as selecting the erroneous part of the string the user has typed).
	IsPartialStringValidProposedSelectedRangeOriginalStringOriginalSelectedRangeErrorDescription(partialStringPtr string, proposedSelRangePtr NSRangePointer, origString string, origSelRange NSRange, error_ string) bool

	InitWithCoder(coder INSCoder) Formatter
	// Encodes the receiver using a given archiver.
	EncodeWithCoder(coder INSCoder)
}





// Init initializes the instance.
func (f Formatter) Init() Formatter {
	rv := objc.Send[Formatter](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f Formatter) Autorelease() Formatter {
	rv := objc.Send[Formatter](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewFormatter creates a new Formatter instance.
func NewFormatter() Formatter {
	class := getFormatterClass()
	rv := objc.Send[Formatter](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewFormatterWithCoder(coder INSCoder) Formatter {
	instance := getFormatterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return FormatterFromID(rv)
}







// The default implementation of this method raises an exception.
//
// obj: The object for which a textual representation is returned.
//
// # Return Value
// 
// An [NSString] object that textually represents `object` for display.
// Returns `nil` if `object` is not of the correct class.
//
// # Discussion
// 
// When implementing a subclass, return the [NSString] object that textually
// represents the cell’s object for display and—if
// [EditingStringForObjectValue] is unimplemented—for editing. First test
// the passed-in object to see if it’s of the correct class. If it isn’t,
// return `nil`; but if it is of the right class, return a properly formatted
// and, if necessary, localized string. (See the specification of the
// [NSString] class for formatting and localizing details.)
// 
// The following implementation (which is paired with the
// [GetObjectValueForStringErrorDescription] example above) prefixes a
// two-digit float representation with a dollar sign:
//
// See: https://developer.apple.com/documentation/Foundation/Formatter/string(for:)
func (f Formatter) StringForObjectValue(obj objectivec.IObject) string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("stringForObjectValue:"), obj)
	return NSStringFromID(rv).String()
}

// The default implementation returns `nil` to indicate that the formatter
// object does not provide an attributed string.
//
// obj: The object for which a textual representation is returned.
//
// attrs: The default attributes to use for the returned attributed string.
//
// # Return Value
// 
// An attributed string that represents `anObject`.
//
// # Discussion
// 
// When implementing a subclass, return an [NSAttributedString] object if the
// string for display should have some attributes. For instance, you might
// want negative values in a financial application to appear in red text.
// Invoke your implementation of [StringForObjectValue] to get the
// non-attributed string, then create an [NSAttributedString] object with it
// (see [InitWithString]). Use the `attributes` default dictionary to reset
// the attributes of the string when a change in value warrants it (for
// example, a negative value becomes positive) For information on creating
// attributed strings, see [Attributed String Programming Guide].
//
// [Attributed String Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/AttributedStrings/AttributedStrings.html#//apple_ref/doc/uid/10000036i
//
// See: https://developer.apple.com/documentation/Foundation/Formatter/attributedString(for:withDefaultAttributes:)
func (f Formatter) AttributedStringForObjectValueWithDefaultAttributes(obj objectivec.IObject, attrs INSDictionary) INSAttributedString {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("attributedStringForObjectValue:withDefaultAttributes:"), obj, attrs)
	return NSAttributedStringFromID(rv)
}

// The default implementation of this method invokes [StringForObjectValue].
//
// obj: The object for which to return an editing string.
//
// # Return Value
// 
// An [NSString] object that is used for editing the textual representation of
// `anObject`.
//
// # Discussion
// 
// When implementing a subclass, override this method only when the string
// that users see and the string that they edit are different. In your
// implementation, return an [NSString] object that is used for editing,
// following the logic recommended for implementing [StringForObjectValue]. As
// an example, you would implement this method if you want the dollar signs in
// displayed strings removed for editing.
//
// See: https://developer.apple.com/documentation/Foundation/Formatter/editingString(for:)
func (f Formatter) EditingStringForObjectValue(obj objectivec.IObject) string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("editingStringForObjectValue:"), obj)
	return NSStringFromID(rv).String()
}

// The default implementation of this method raises an exception.
//
// obj: If conversion is successful, upon return contains the object created from
// `string`.
//
// string: The string to parse.
//
// error: If non-`nil`, if there is a error during the conversion, upon return
// contains an [NSString] object that describes the problem.
//
// # Return Value
// 
// [true] if the conversion from string to cell content object was successful,
// otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// When implementing this method in a subclass, return by reference the object
// `anObject` created from `string`. If `string` is equal to the value of the
// converted object, such as for formatters whose converted value type is
// [NSString], it can be returned by reference without creating a new object.
// 
// Return [true] if the conversion is successful. If you return [false], also
// return by indirection (in `error`) a localized user-presentable [NSString]
// object that explains the reason why the conversion failed; the delegate (if
// any) of the [NSControl] object managing the cell can then respond to the
// failure in control:didFailToFormatString:errorDescription:. However, if
// `error` is `nil`, the sender is not interested in the error description,
// and you should not attempt to assign one.
// 
// The following example (which is paired with the example given in
// [StringForObjectValue]) converts a string representation of a dollar amount
// that includes the dollar sign; it uses an [NSScanner] instance to convert
// this amount to a float after stripping out the initial dollar sign.
// 
// # Special Considerations
// 
// Prior to OS X v10.6, the implementation of this method in both
// [NSNumberFormatter] and [NSDateFormatter] would return [true] and an object
// value even if only part of the string could be parsed. This is problematic
// because you cannot be sure what portion of the string was parsed. For
// applications linked on or after OS X v10.6, this method instead returns an
// error if part of the string cannot be parsed. You can use `` to get the old
// behavior—it returns the range of the substring that was successfully
// parsed.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/Formatter/getObjectValue(_:for:errorDescription:)
func (f Formatter) GetObjectValueForStringErrorDescription(obj []objectivec.IObject, string_ string, error_ string) bool {
	rv := objc.Send[bool](f.ID, objc.Sel("getObjectValue:forString:errorDescription:"), objectivec.IObjectSliceToNSArray(obj), objc.String(string_), objc.String(error_))
	return rv
}

// Returns a Boolean value that indicates whether a partial string is valid.
//
// partialString: The text currently in a cell.
//
// newString: If `partialString` needs to be modified, upon return contains the
// replacement string.
//
// error: If non-`nil`, if validation fails contains an [NSString] object that
// describes the problem.
//
// # Return Value
// 
// [true] if `partialString` is an acceptable value, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is invoked each time the user presses a key while the cell has
// the keyboard focus—it lets you verify and edit the cell text as the user
// types it.
// 
// In a subclass implementation, evaluate `partialString` according to the
// context, edit the text if necessary, and return by reference any edited
// string in `newString`. Return [true] if `partialString` is acceptable and
// [false] if `partialString` is unacceptable. If you return [false] and
// `newString` is `nil`, the cell displays `partialString` minus the last
// character typed. If you return [false], you can also return by indirection
// an [NSString] object (in `error`) that explains the reason why the
// validation failed; the delegate (if any) of the [NSControl] object managing
// the cell can then respond to the failure in
// control:didFailToValidatePartialString:errorDescription:. The selection
// range will always be set to the end of the text if replacement occurs.
// 
// This method is a compatibility method. If a subclass overrides this method
// and does not override
// [IsPartialStringValidProposedSelectedRangeOriginalStringOriginalSelectedRangeErrorDescription],
// this method will be called as before
// ([IsPartialStringValidProposedSelectedRangeOriginalStringOriginalSelectedRangeErrorDescription]
// just calls this one by default).
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/Formatter/isPartialStringValid(_:newEditingString:errorDescription:)
func (f Formatter) IsPartialStringValidNewEditingStringErrorDescription(partialString string, newString string, error_ string) bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isPartialStringValid:newEditingString:errorDescription:"), objc.String(partialString), objc.String(newString), objc.String(error_))
	return rv
}

// This method should be implemented in subclasses that want to validate user
// changes to a string in a field, where the user changes are not necessarily
// at the end of the string, and preserve the selection (or set a different
// one, such as selecting the erroneous part of the string the user has
// typed).
//
// partialStringPtr: The new string to validate.
//
// proposedSelRangePtr: The selection range that will be used if the string is accepted or
// replaced.
//
// origString: The original string, before the proposed change.
//
// origSelRange: The selection range over which the change is to take place.
// 
// If the user change is a deletion, `origSelRange` contains the range of the
// deleted characters.
//
// error: If non-`nil`, if validation fails contains an [NSString] object that
// describes the problem.
//
// # Return Value
// 
// [true] if `partialStringPtr` is acceptable, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// In a subclass implementation, evaluate `partialString` according to the
// context. Return [true] if `partialStringPtr` is acceptable and [false] if
// `partialStringPtr` is unacceptable. If you return [false] and assign a new
// string to `partialStringPtr` and a new range to `proposedSelRangePtr`, the
// string and selection range are changed, otherwise, if no values are
// assigned to `partialStringPtr` or `proposedSelRangePtr`, the change is
// rejected. If you return [false], you can also return by indirection an
// [NSString] object (in `error`) that explains the reason why the validation
// failed; the delegate (if any) of the [NSControl] object managing the cell
// can then respond to the failure in
// control:didFailToValidatePartialString:errorDescription:.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/Formatter/isPartialStringValid(_:proposedSelectedRange:originalString:originalSelectedRange:errorDescription:)
func (f Formatter) IsPartialStringValidProposedSelectedRangeOriginalStringOriginalSelectedRangeErrorDescription(partialStringPtr string, proposedSelRangePtr NSRangePointer, origString string, origSelRange NSRange, error_ string) bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isPartialStringValid:proposedSelectedRange:originalString:originalSelectedRange:errorDescription:"), objc.String(partialStringPtr), proposedSelRangePtr, objc.String(origString), origSelRange, objc.String(error_))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (f Formatter) InitWithCoder(coder INSCoder) Formatter {
	rv := objc.Send[Formatter](f.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (f Formatter) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}























			// Protocol methods for NSCopying
			














