// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextStorage] class.
var (
	_NSTextStorageClass     NSTextStorageClass
	_NSTextStorageClassOnce sync.Once
)

func getNSTextStorageClass() NSTextStorageClass {
	_NSTextStorageClassOnce.Do(func() {
		_NSTextStorageClass = NSTextStorageClass{class: objc.GetClass("NSTextStorage")}
	})
	return _NSTextStorageClass
}

// GetNSTextStorageClass returns the class object for NSTextStorage.
func GetNSTextStorageClass() NSTextStorageClass {
	return getNSTextStorageClass()
}

type NSTextStorageClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextStorageClass) Alloc() NSTextStorage {
	rv := objc.Send[NSTextStorage](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// The fundamental storage mechanism of TextKit that contains the text managed
// by the system.
//
// # Overview
// 
// [NSTextStorage] is a semi-concrete subclass of [NSMutableAttributedString]
// that adds behavior for managing a set of client [NSLayoutManager] objects.
// A text storage object notifies its layout managers of changes to its
// characters or attributes, which lets the layout managers redisplay the text
// as needed.
// 
// You can access a text storage object from any thread of your app, but your
// app must guarantee access from only one thread at a time.
// 
// In macOS, this class also defines properties for getting and setting
// scriptable attributes of [NSTextStorage] objects. Unless you’re dealing
// with scriptability, you shouldn’t access these properties directly. In
// particular, using the [NSTextStorage.Characters], [NSTextStorage.Words], or [NSTextStorage.Paragraphs] properties is
// an inefficient way to manipulate the text storage, since accessing these
// properties involves the creation of many objects. Instead, use the text
// access methods defined by [NSMutableAttributedString],
// [NSAttributedString], [NSMutableString], and [NSString] to perform
// character-level manipulation.
// 
// # Subclassing Notes
// 
// The [NSTextStorage] class implements change management through the
// [beginEditing()] and [endEditing()] methods, as well as verification of
// attributes, delegate handling, and layout management notification. The one
// aspect it doesn’t implement is managing the actual attributed string
// storage, which subclasses manage by overriding the two [NSAttributedString]
// primitives:
// 
// - [NSTextStorage.String]
// - [attributes(at:effectiveRange:)]
// 
// Subclasses must also override two [NSMutableAttributedString] primitives:
// 
// - [replaceCharacters(in:with:)]
// - [setAttributes(_:range:)]
// 
// These primitives should perform the change, then call
// [NSTextStorage.EditedRangeChangeInLength] to let the parent class know there are changes.
//
// [NSAttributedString]: https://developer.apple.com/documentation/Foundation/NSAttributedString
// [NSMutableAttributedString]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString
// [NSMutableString]: https://developer.apple.com/documentation/Foundation/NSMutableString
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [attributes(at:effectiveRange:)]: https://developer.apple.com/documentation/Foundation/NSAttributedString/attributes(at:effectiveRange:)
// [beginEditing()]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/beginEditing()
// [endEditing()]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/endEditing()
// [replaceCharacters(in:with:)]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/replaceCharacters(in:with:)-6oq9r
// [setAttributes(_:range:)]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/setAttributes(_:range:)
//
// # Processing the editing actions
//
//   - [NSTextStorage.Delegate]: The delegate for the text storage object.
//   - [NSTextStorage.SetDelegate]
//
// # Accessing the layout managers
//
//   - [NSTextStorage.LayoutManagers]: The layout managers for the text storage object.
//   - [NSTextStorage.AddLayoutManager]: Adds a layout manager to the text storage object’s set of layout managers.
//   - [NSTextStorage.RemoveLayoutManager]: Removes a layout manager from the text storage object’s set of layout managers.
//
// # Managing edits
//
//   - [NSTextStorage.EditedRangeChangeInLength]: Tracks changes made to the text storage object, allowing the text storage to record the full extent of changes.
//   - [NSTextStorage.ProcessEditing]: Cleans up changes to the text storage object and notifies its delegate and layout managers of changes.
//
// # Fixing the string attributes
//
//   - [NSTextStorage.InvalidateAttributesInRange]: Invalidates attributes in the specified range.
//   - [NSTextStorage.EnsureAttributesAreFixedInRange]: Ensures that attribute fixing occurs in the specified range.
//   - [NSTextStorage.FixesAttributesLazily]: A Boolean value that indicates whether the text storage object fixes attributes lazily.
//
// # Determining the nature of changes
//
//   - [NSTextStorage.EditedMask]: A mask that describes the kinds of edits pending for the text storage object.
//   - [NSTextStorage.EditedRange]: The range of text that contains changes.
//   - [NSTextStorage.ChangeInLength]: The difference between the current length of the edited range and its length before editing.
//
// # Accessing scriptable properties
//
//   - [NSTextStorage.AttributeRuns]: The text storage contents as an array of attribute runs.
//   - [NSTextStorage.SetAttributeRuns]
//   - [NSTextStorage.Paragraphs]: The text storage contents as an array of paragraphs.
//   - [NSTextStorage.SetParagraphs]
//   - [NSTextStorage.Words]: The text storage contents as an array of words.
//   - [NSTextStorage.SetWords]
//   - [NSTextStorage.Characters]: The text storage contents as an array of characters.
//   - [NSTextStorage.SetCharacters]
//   - [NSTextStorage.Font]: The font for the text storage.
//   - [NSTextStorage.SetFont]
//   - [NSTextStorage.ForegroundColor]: The color for the text.
//   - [NSTextStorage.SetForegroundColor]
//
// # Accessing the storage controller
//
//   - [NSTextStorage.TextStorageObserver]: The observer for the text storage object.
//   - [NSTextStorage.SetTextStorageObserver]
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorage
type NSTextStorage struct {
	foundation.NSMutableAttributedString
}

// NSTextStorageFromID constructs a [NSTextStorage] from an objc.ID.
//
// The fundamental storage mechanism of TextKit that contains the text managed
// by the system.
func NSTextStorageFromID(id objc.ID) NSTextStorage {
	return NSTextStorage{NSMutableAttributedString: foundation.NSMutableAttributedStringFromID(id)}
}
// NOTE: NSTextStorage adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSTextStorage] class.
//
// # Processing the editing actions
//
//   - [INSTextStorage.Delegate]: The delegate for the text storage object.
//   - [INSTextStorage.SetDelegate]
//
// # Accessing the layout managers
//
//   - [INSTextStorage.LayoutManagers]: The layout managers for the text storage object.
//   - [INSTextStorage.AddLayoutManager]: Adds a layout manager to the text storage object’s set of layout managers.
//   - [INSTextStorage.RemoveLayoutManager]: Removes a layout manager from the text storage object’s set of layout managers.
//
// # Managing edits
//
//   - [INSTextStorage.EditedRangeChangeInLength]: Tracks changes made to the text storage object, allowing the text storage to record the full extent of changes.
//   - [INSTextStorage.ProcessEditing]: Cleans up changes to the text storage object and notifies its delegate and layout managers of changes.
//
// # Fixing the string attributes
//
//   - [INSTextStorage.InvalidateAttributesInRange]: Invalidates attributes in the specified range.
//   - [INSTextStorage.EnsureAttributesAreFixedInRange]: Ensures that attribute fixing occurs in the specified range.
//   - [INSTextStorage.FixesAttributesLazily]: A Boolean value that indicates whether the text storage object fixes attributes lazily.
//
// # Determining the nature of changes
//
//   - [INSTextStorage.EditedMask]: A mask that describes the kinds of edits pending for the text storage object.
//   - [INSTextStorage.EditedRange]: The range of text that contains changes.
//   - [INSTextStorage.ChangeInLength]: The difference between the current length of the edited range and its length before editing.
//
// # Accessing scriptable properties
//
//   - [INSTextStorage.AttributeRuns]: The text storage contents as an array of attribute runs.
//   - [INSTextStorage.SetAttributeRuns]
//   - [INSTextStorage.Paragraphs]: The text storage contents as an array of paragraphs.
//   - [INSTextStorage.SetParagraphs]
//   - [INSTextStorage.Words]: The text storage contents as an array of words.
//   - [INSTextStorage.SetWords]
//   - [INSTextStorage.Characters]: The text storage contents as an array of characters.
//   - [INSTextStorage.SetCharacters]
//   - [INSTextStorage.Font]: The font for the text storage.
//   - [INSTextStorage.SetFont]
//   - [INSTextStorage.ForegroundColor]: The color for the text.
//   - [INSTextStorage.SetForegroundColor]
//
// # Accessing the storage controller
//
//   - [INSTextStorage.TextStorageObserver]: The observer for the text storage object.
//   - [INSTextStorage.SetTextStorageObserver]
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorage
type INSTextStorage interface {
	foundation.INSMutableAttributedString
	NSPasteboardWriting

	// Topic: Processing the editing actions

	// The delegate for the text storage object.
	Delegate() NSTextStorageDelegate
	SetDelegate(value NSTextStorageDelegate)

	// Topic: Accessing the layout managers

	// The layout managers for the text storage object.
	LayoutManagers() []NSLayoutManager
	// Adds a layout manager to the text storage object’s set of layout managers.
	AddLayoutManager(aLayoutManager INSLayoutManager)
	// Removes a layout manager from the text storage object’s set of layout managers.
	RemoveLayoutManager(aLayoutManager INSLayoutManager)

	// Topic: Managing edits

	// Tracks changes made to the text storage object, allowing the text storage to record the full extent of changes.
	EditedRangeChangeInLength(editedMask NSTextStorageEditActions, editedRange foundation.NSRange, delta int)
	// Cleans up changes to the text storage object and notifies its delegate and layout managers of changes.
	ProcessEditing()

	// Topic: Fixing the string attributes

	// Invalidates attributes in the specified range.
	InvalidateAttributesInRange(range_ foundation.NSRange)
	// Ensures that attribute fixing occurs in the specified range.
	EnsureAttributesAreFixedInRange(range_ foundation.NSRange)
	// A Boolean value that indicates whether the text storage object fixes attributes lazily.
	FixesAttributesLazily() bool

	// Topic: Determining the nature of changes

	// A mask that describes the kinds of edits pending for the text storage object.
	EditedMask() NSTextStorageEditActions
	// The range of text that contains changes.
	EditedRange() foundation.NSRange
	// The difference between the current length of the edited range and its length before editing.
	ChangeInLength() int

	// Topic: Accessing scriptable properties

	// The text storage contents as an array of attribute runs.
	AttributeRuns() []NSTextStorage
	SetAttributeRuns(value []NSTextStorage)
	// The text storage contents as an array of paragraphs.
	Paragraphs() []NSTextStorage
	SetParagraphs(value []NSTextStorage)
	// The text storage contents as an array of words.
	Words() []NSTextStorage
	SetWords(value []NSTextStorage)
	// The text storage contents as an array of characters.
	Characters() []NSTextStorage
	SetCharacters(value []NSTextStorage)
	// The font for the text storage.
	Font() NSFont
	SetFont(value NSFont)
	// The color for the text.
	ForegroundColor() INSColor
	SetForegroundColor(value INSColor)

	// Topic: Accessing the storage controller

	// The observer for the text storage object.
	TextStorageObserver() NSTextStorageObserving
	SetTextStorageObserver(value NSTextStorageObserving)

	// Initializes an instance with a property list object and a type string.
	InitWithPasteboardPropertyListOfType(propertyList objectivec.IObject, type_ NSPasteboardType) NSTextStorage
}





// Init initializes the instance.
func (t NSTextStorage) Init() NSTextStorage {
	rv := objc.Send[NSTextStorage](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextStorage) Autorelease() NSTextStorage {
	rv := objc.Send[NSTextStorage](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextStorage creates a new NSTextStorage instance.
func NewNSTextStorage() NSTextStorage {
	class := getNSTextStorageClass()
	rv := objc.Send[NSTextStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Initializes an instance with a property list object and a type string.
//
// propertyList: A property list containing data to initialize the receiver.
// 
// By default, the property list object is an instance of [NSData]. If you
// implement [ReadingOptionsForTypePasteboard] and specify an option other
// than [PasteboardReadingAsData], the `propertyList` may be any other
// property list object.
//
// type: A UTI supported by the receiver for reading (one of the types returned by
// [ReadableTypesForPasteboard]).
//
// # Return Value
// 
// An object initialized using the data in `propertyList`.
//
// # Discussion
// 
// This method is considered optional because, if [ReadableTypesForPasteboard]
// returns just a single type, and that type uses the
// [PasteboardReadingAsKeyedArchive] reading option, then instances are
// initialized using [init(coder:)] instead of this method.
//
// [init(coder:)]: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/init(pasteboardPropertyList:ofType:)
func NewTextStorageWithPasteboardPropertyListOfType(propertyList objectivec.IObject, type_ NSPasteboardType) NSTextStorage {
	instance := getNSTextStorageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPasteboardPropertyList:ofType:"), propertyList, objc.String(string(type_)))
	return NSTextStorageFromID(rv)
}







// Adds a layout manager to the text storage object’s set of layout
// managers.
//
// aLayoutManager: The layout manager to add.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorage/addLayoutManager(_:)
func (t NSTextStorage) AddLayoutManager(aLayoutManager INSLayoutManager) {
	objc.Send[objc.ID](t.ID, objc.Sel("addLayoutManager:"), aLayoutManager)
}

// Removes a layout manager from the text storage object’s set of layout
// managers.
//
// aLayoutManager: The layout manager to remove.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorage/removeLayoutManager(_:)
func (t NSTextStorage) RemoveLayoutManager(aLayoutManager INSLayoutManager) {
	objc.Send[objc.ID](t.ID, objc.Sel("removeLayoutManager:"), aLayoutManager)
}

// Tracks changes made to the text storage object, allowing the text storage
// to record the full extent of changes.
//
// editedMask: A mask specifying the nature of the changes. You make the value by
// combining with the C bitwise OR operator the options described in
// [NSTextStorageEditActions].
// //
// [NSTextStorageEditActions]: https://developer.apple.com/documentation/AppKit/NSTextStorageEditActions
//
// editedRange: The extent of characters affected before the change took place.
//
// delta: The number of characters added to or removed from `oldRange`. If no
// characters where edited as noted by `mask`, its value is irrelevant and
// undefined. For example, when replacing “The” with “Several” in the
// string “The files couldn’t be saved”, `oldRange` is {0, 3} and
// `lengthChange` is 4.
//
// # Discussion
// 
// This method invokes [ProcessEditing] if there are no outstanding
// [beginEditing()] calls. [NSTextStorage] invokes this method automatically
// each time it makes a change to its attributed string. Subclasses that
// override or add methods that alter their attributed strings directly should
// invoke this method after making those changes; otherwise you shouldn’t
// invoke this method. The information accumulated with this method is then
// used in an invocation of [ProcessEditing] to report the affected portion of
// the receiver.
// 
// The methods for querying changes, [EditedRange] and [ChangeInLength],
// indicate the extent of characters affected after the change. This method
// expects the characters before the change because that information is
// readily available as the argument to whatever method performs the change
// (such as [replaceCharacters(in:with:)]).
//
// [beginEditing()]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/beginEditing()
// [replaceCharacters(in:with:)]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/replaceCharacters(in:with:)-6oq9r
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorage/edited(_:range:changeInLength:)
func (t NSTextStorage) EditedRangeChangeInLength(editedMask NSTextStorageEditActions, editedRange foundation.NSRange, delta int) {
	objc.Send[objc.ID](t.ID, objc.Sel("edited:range:changeInLength:"), editedMask, editedRange, delta)
}

// Cleans up changes to the text storage object and notifies its delegate and
// layout managers of changes.
//
// # Discussion
// 
// This method is automatically invoked in response to an
// [EditedRangeChangeInLength] message or an [endEditing()] message if edits
// were made within the scope of a [beginEditing()] block. You should never
// need to invoke it directly.
// 
// This method begins by posting an [willProcessEditingNotification] to the
// default notification center (which results in the delegate receiving a
// [TextStorageWillProcessEditingRangeChangeInLength] message). Then it fixes
// attributes. After this, it posts an [didProcessEditingNotification] to the
// default notification center (which results in the delegate receiving a
// [TextStorageDidProcessEditingRangeChangeInLength] message). Finally, it
// sends a [textStorage(_:edited:range:changeInLength:invalidatedRange:)]
// message to each of the receiver’s layout managers using the argument
// values provided.
//
// [beginEditing()]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/beginEditing()
// [didProcessEditingNotification]: https://developer.apple.com/documentation/AppKit/NSTextStorage/didProcessEditingNotification
// [endEditing()]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/endEditing()
// [textStorage(_:edited:range:changeInLength:invalidatedRange:)]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/textStorage(_:edited:range:changeInLength:invalidatedRange:)
// [willProcessEditingNotification]: https://developer.apple.com/documentation/AppKit/NSTextStorage/willProcessEditingNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorage/processEditing()
func (t NSTextStorage) ProcessEditing() {
	objc.Send[objc.ID](t.ID, objc.Sel("processEditing"))
}

// Invalidates attributes in the specified range.
//
// range: The range of characters whose attributes the method should invalidate.
//
// # Discussion
// 
// Called from [ProcessEditing] to invalidate attributes when the text storage
// changes. If the receiver isn’t lazy, this method calls
// [fixAttributes(in:)]. If lazy attribute fixing is in effect, this method
// instead records the range needing fixing.
//
// [fixAttributes(in:)]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/fixAttributes(in:)
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorage/invalidateAttributes(in:)
func (t NSTextStorage) InvalidateAttributesInRange(range_ foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("invalidateAttributesInRange:"), range_)
}

// Ensures that attribute fixing occurs in the specified range.
//
// range: The range of characters to examine.
//
// # Discussion
// 
// An [NSTextStorage] object using lazy attribute fixing is required to call
// this method before accessing any attributes within `range`. This method
// gives attribute fixing a chance to occur if necessary. [NSTextStorage]
// subclasses wishing to support laziness must call this method from all
// attribute accessors they implement.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorage/ensureAttributesAreFixed(in:)
func (t NSTextStorage) EnsureAttributesAreFixedInRange(range_ foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("ensureAttributesAreFixedInRange:"), range_)
}

// Initializes an instance with a property list object and a type string.
//
// propertyList: A property list containing data to initialize the receiver.
// 
// By default, the property list object is an instance of [NSData]. If you
// implement [ReadingOptionsForTypePasteboard] and specify an option other
// than [PasteboardReadingAsData], the `propertyList` may be any other
// property list object.
//
// type: A UTI supported by the receiver for reading (one of the types returned by
// [ReadableTypesForPasteboard]).
//
// # Return Value
// 
// An object initialized using the data in `propertyList`.
//
// # Discussion
// 
// This method is considered optional because, if [ReadableTypesForPasteboard]
// returns just a single type, and that type uses the
// [PasteboardReadingAsKeyedArchive] reading option, then instances are
// initialized using [init(coder:)] instead of this method.
//
// [init(coder:)]: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/init(pasteboardPropertyList:ofType:)
func (t NSTextStorage) InitWithPasteboardPropertyListOfType(propertyList objectivec.IObject, type_ NSPasteboardType) NSTextStorage {
	rv := objc.Send[NSTextStorage](t.ID, objc.Sel("initWithPasteboardPropertyList:ofType:"), propertyList, objc.String(string(type_)))
	return rv
}

// Returns a property list object to represent the receiver on a pasteboard as
// an object of a specified type.
//
// type: One of the types the receiver supports for writing (one of the UTIs
// returned by its implementation of [WritableTypesForPasteboard]).
//
// # Return Value
// 
// A property list object to represent the receiver on a pasteboard as an
// object of type `type`.
//
// # Discussion
// 
// The returned value will commonly be the [NSData] object for the specified
// data type. However, if this method returns either a string, or any other
// property-list type, the pasteboard will automatically convert these items
// to the correct data format required for the pasteboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardWriting/pasteboardPropertyList(forType:)
func (t NSTextStorage) PasteboardPropertyListForType(type_ NSPasteboardType) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("pasteboardPropertyListForType:"), objc.String(string(type_)))
	return objectivec.Object{ID: rv}
}

// Returns an array of UTI strings of data types the receiver can write to a
// given pasteboard.
//
// pasteboard: A pasteboard.
// 
// You can use this argument to provide different options based on the
// pasteboard name, if you need to.
//
// # Return Value
// 
// An array of UTI strings of data types the receiver can write to
// `pasteboard`.
//
// # Discussion
// 
// By default, data for the first returned type is put onto the pasteboard
// immediately, with the remaining types being promised.
// 
// To change the default behavior, implement
// -writingOptionsForType:pasteboard: and return [PasteboardWritingPromised]
// to lazily provide data for types, return no option to provide the data for
// that type immediately. Use the pasteboard argument to provide different
// types based on the pasteboard name, if desired. Do not perform other
// pasteboard operations in the method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardWriting/writableTypes(for:)
func (t NSTextStorage) WritableTypesForPasteboard(pasteboard INSPasteboard) []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("writableTypesForPasteboard:"), pasteboard)
	return objc.ConvertSliceToStrings(rv)
}

// Returns options for writing data of a specified type to a given pasteboard.
//
// type: One of the types the receiver supports for writing (one of the UTIs
// returned by its implementation of [WritableTypesForPasteboard]).
//
// pasteboard: A pasteboard.
// 
// You can use this argument to provide different options based on the
// pasteboard name, if you need to.
//
// # Return Value
// 
// Options for writing data of type type to `pasteboard`. Return `0` for no
// options, or a value given in [Pasteboard Writing Options].
//
// [Pasteboard Writing Options]: https://developer.apple.com/documentation/AppKit/pasteboard-writing-options
//
// # Discussion
// 
// Do not perform other pasteboard operations in the method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardWriting/writingOptions(forType:pasteboard:)
func (t NSTextStorage) WritingOptionsForTypePasteboard(type_ NSPasteboardType, pasteboard INSPasteboard) NSPasteboardWritingOptions {
	rv := objc.Send[NSPasteboardWritingOptions](t.ID, objc.Sel("writingOptionsForType:pasteboard:"), objc.String(string(type_)), pasteboard)
	return NSPasteboardWritingOptions(rv)
}





// Returns an array of uniform type identifier strings of data types the
// receiver can read from the pasteboard and initialize from.
//
// pasteboard: A pasteboard. You can use the pasteboard argument to provide different
// types based on the pasteboard name, if you need to.
//
// # Return Value
// 
// An array of uniform type identifier strings of data types instances that
// the receiver can read from the pasteboard and initialize from.
//
// # Discussion
// 
// By default, the system provides the data for a type to
// [InitWithPasteboardPropertyListOfType] as an instance of [NSData]. If you
// implement [ReadingOptionsForTypePasteboard] and specify a different option,
// the system converts the [NSData] object for a type to an [NSString] object
// or any other property list object.
// 
// # Special Considerations
// 
// Don’t perform other pasteboard operations in the method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/readableTypes(for:)
func (_NSTextStorageClass NSTextStorageClass) ReadableTypesForPasteboard(pasteboard INSPasteboard) []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSTextStorageClass.class), objc.Sel("readableTypesForPasteboard:"), pasteboard)
	return objc.ConvertSliceToStrings(rv)
}

// Returns options for reading data of a specified type from a given
// pasteboard.
//
// type: A UTI supported by instances of the receiver for reading (one of the types
// returned by [ReadableTypesForPasteboard]).
//
// pasteboard: A pasteboard.
// 
// You can use the pasteboard argument to provide return different based on
// the pasteboard name, should you need to do so.
//
// # Return Value
// 
// Options for reading data of `type` from `pasteboard`. For a list of valid
// values, see [NSPasteboard.ReadingOptions].
//
// [NSPasteboard.ReadingOptions]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptions
//
// # Discussion
// 
// Do not perform other pasteboard operations in this method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/readingOptions(forType:pasteboard:)
func (_NSTextStorageClass NSTextStorageClass) ReadingOptionsForTypePasteboard(type_ NSPasteboardType, pasteboard INSPasteboard) NSPasteboardReadingOptions {
	rv := objc.Send[NSPasteboardReadingOptions](objc.ID(_NSTextStorageClass.class), objc.Sel("readingOptionsForType:pasteboard:"), objc.String(string(type_)), pasteboard)
	return NSPasteboardReadingOptions(rv)
}








// The delegate for the text storage object.
//
// # Discussion
// 
// Use a delegate object to monitor edits occurring to the text contents. Your
// delegate object must conform to the [NSTextStorageDelegate] protocol.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorage/delegate
func (t NSTextStorage) Delegate() NSTextStorageDelegate {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("delegate"))
	return NSTextStorageDelegateObjectFromID(rv)
}
func (t NSTextStorage) SetDelegate(value NSTextStorageDelegate) {
	objc.Send[struct{}](t.ID, objc.Sel("setDelegate:"), value)
}



// The layout managers for the text storage object.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorage/layoutManagers
func (t NSTextStorage) LayoutManagers() []NSLayoutManager {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("layoutManagers"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSLayoutManager {
		return NSLayoutManagerFromID(id)
	})
}



// A Boolean value that indicates whether the text storage object fixes
// attributes lazily.
//
// # Discussion
// 
// When subclassing, the default value of this property is [false], meaning
// that your subclass fixes attributes immediately when they change. The
// system’s concrete subclass overrides this property and sets it to [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorage/fixesAttributesLazily
func (t NSTextStorage) FixesAttributesLazily() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("fixesAttributesLazily"))
	return rv
}



// A mask that describes the kinds of edits pending for the text storage
// object.
//
// # Discussion
// 
// This property indicates pending changes for attributes, characters, or
// both. Use the C bitwise AND operator to test the value against the
// [TextStorageEditedAttributes] or [TextStorageEditedCharacters] constants;
// testing for equality fails if you add additional mask flags later. The text
// storage object’s associated delegate and layout managers can use this
// information to determine the nature of edits in their respective
// notification methods.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorage/editedMask
func (t NSTextStorage) EditedMask() NSTextStorageEditActions {
	rv := objc.Send[NSTextStorageEditActions](t.ID, objc.Sel("editedMask"))
	return NSTextStorageEditActions(rv)
}



// The range of text that contains changes.
//
// # Discussion
// 
// The specified range can reflect changes to characters or attributes. The
// text storage object’s delegate and layout managers can use this
// information to determine the nature of edits in their respective
// notification methods.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorage/editedRange
func (t NSTextStorage) EditedRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("editedRange"))
	return foundation.NSRange(rv)
}



// The difference between the current length of the edited range and its
// length before editing.
//
// # Discussion
// 
// This property reflects difference between the current length of the edited
// range and its length before editing began—that is, before the first call
// to the [beginEditing()] method or a single call to
// the[EditedRangeChangeInLength] method. This difference is accumulated with
// each call to the [EditedRangeChangeInLength] method, until the changes are
// finally processed.
// 
// The text storage object’s delegate and layout managers can use this
// information to determine the nature of edits in their respective
// notification methods.
//
// [beginEditing()]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/beginEditing()
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorage/changeInLength
func (t NSTextStorage) ChangeInLength() int {
	rv := objc.Send[int](t.ID, objc.Sel("changeInLength"))
	return rv
}



// The text storage contents as an array of attribute runs.
//
// # Discussion
// 
// Unless you’re dealing with scriptability, you shouldn’t use or modify
// this property directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorage/attributeRuns
func (t NSTextStorage) AttributeRuns() []NSTextStorage {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("attributeRuns"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSTextStorage {
		return NSTextStorageFromID(id)
	})
}
func (t NSTextStorage) SetAttributeRuns(value []NSTextStorage) {
	objc.Send[struct{}](t.ID, objc.Sel("setAttributeRuns:"), objectivec.IObjectSliceToNSArray(value))
}



// The text storage contents as an array of paragraphs.
//
// # Discussion
// 
// Unless you’re dealing with scriptability, you shouldn’t use or modify
// this property directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorage/paragraphs
func (t NSTextStorage) Paragraphs() []NSTextStorage {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("paragraphs"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSTextStorage {
		return NSTextStorageFromID(id)
	})
}
func (t NSTextStorage) SetParagraphs(value []NSTextStorage) {
	objc.Send[struct{}](t.ID, objc.Sel("setParagraphs:"), objectivec.IObjectSliceToNSArray(value))
}



// The text storage contents as an array of words.
//
// # Discussion
// 
// Unless you’re dealing with scriptability, you shouldn’t use or modify
// this property directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorage/words
func (t NSTextStorage) Words() []NSTextStorage {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("words"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSTextStorage {
		return NSTextStorageFromID(id)
	})
}
func (t NSTextStorage) SetWords(value []NSTextStorage) {
	objc.Send[struct{}](t.ID, objc.Sel("setWords:"), objectivec.IObjectSliceToNSArray(value))
}



// The text storage contents as an array of characters.
//
// # Discussion
// 
// Unless you’re dealing with scriptability, you shouldn’t use or modify
// this property directly. For indexed access to characters, use
// [NSAttributedString]’s [length] method to access the string, and
// [NSString]’s [character(at:)] method to access the individual characters.
//
// [character(at:)]: https://developer.apple.com/documentation/Foundation/NSString/character(at:)
// [length]: https://developer.apple.com/documentation/Foundation/NSAttributedString/length
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorage/characters
func (t NSTextStorage) Characters() []NSTextStorage {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("characters"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSTextStorage {
		return NSTextStorageFromID(id)
	})
}
func (t NSTextStorage) SetCharacters(value []NSTextStorage) {
	objc.Send[struct{}](t.ID, objc.Sel("setCharacters:"), objectivec.IObjectSliceToNSArray(value))
}



// The font for the text storage.
//
// # Discussion
// 
// Unless you’re dealing with scriptability, you shouldn’t use or modify
// this property directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorage/font
func (t NSTextStorage) Font() NSFont {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("font"))
	return NSFontFromID(objc.ID(rv))
}
func (t NSTextStorage) SetFont(value NSFont) {
	objc.Send[struct{}](t.ID, objc.Sel("setFont:"), value)
}



// The color for the text.
//
// # Discussion
// 
// Unless you’re dealing with scriptability, you shouldn’t use or modify
// this property directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorage/foregroundColor
func (t NSTextStorage) ForegroundColor() INSColor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("foregroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (t NSTextStorage) SetForegroundColor(value INSColor) {
	objc.Send[struct{}](t.ID, objc.Sel("setForegroundColor:"), value)
}



// The observer for the text storage object.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorage/textStorageObserver
func (t NSTextStorage) TextStorageObserver() NSTextStorageObserving {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textStorageObserver"))
	return NSTextStorageObservingObjectFromID(rv)
}
func (t NSTextStorage) SetTextStorageObserver(value NSTextStorageObserving) {
	objc.Send[struct{}](t.ID, objc.Sel("setTextStorageObserver:"), value)
}


















			// Protocol methods for NSPasteboardWriting
			
















