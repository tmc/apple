// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSTextContentStorage] class.
var (
	_NSTextContentStorageClass     NSTextContentStorageClass
	_NSTextContentStorageClassOnce sync.Once
)

func getNSTextContentStorageClass() NSTextContentStorageClass {
	_NSTextContentStorageClassOnce.Do(func() {
		_NSTextContentStorageClass = NSTextContentStorageClass{class: objc.GetClass("NSTextContentStorage")}
	})
	return _NSTextContentStorageClass
}

// GetNSTextContentStorageClass returns the class object for NSTextContentStorage.
func GetNSTextContentStorageClass() NSTextContentStorageClass {
	return getNSTextContentStorageClass()
}

type NSTextContentStorageClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextContentStorageClass) Alloc() NSTextContentStorage {
	rv := objc.Send[NSTextContentStorage](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A concrete object for managing your view’s text content and generating
// the text elements necessary for layout.
//
// # Overview
// 
// An [NSTextContentStorage] object provides the backing store for a view that
// contains text. This object stores the text in an attributed string object,
// and defaults to using an [NSTextStorage] object. It also maps portions of
// the text to [NSTextElement] objects to organize the text into paragraphs,
// lists, and other common element types found in text content. During layout,
// TextKit uses these elements to lay out and render the text in your view.
// 
// The standard system views use an [NSTextContentStorage] object to manage
// their text content. When building a custom text view, use this type to
// store the text for your view. [NSTextContentStorage] works with an
// associated [NSTextLayoutManager] to lay out your view’s text. When
// someone inserts new text or edits the existing text, call the
// [PerformEditingTransactionUsingBlock] method and use a block to modify the
// contents of the [NSTextContentStorage.AttributedString] property. Wrapping your edits in an edit
// transaction lets the rest of the text system respond to those changes.
// 
// TextKit uses the abstract [NSTextLocation] protocol to identify locations
// within text. [NSTextContentStorage] manager provides its own implementation
// of this protocol to represent locations within its storage object. To get
// the start and end locations, access the object’s [NSTextContentStorage.DocumentRange] property
// and use them to create new location objects. If you provide your own
// implementation of the [NSTextLocation] protocol to manage locations in your
// content, subclass [NSTextContentManager] and implement your own storage
// object to support those locations.
//
// # Managing the stored text
//
//   - [NSTextContentStorage.AttributedString]: An attributed string that contains the contents of the document.
//   - [NSTextContentStorage.SetAttributedString]
//
// # Managing text elements
//
//   - [NSTextContentStorage.TextElementForAttributedString]: Returns the text element corresponding to object’s attributed string.
//   - [NSTextContentStorage.AttributedStringForTextElement]: Returns a new attributed string for the text element.
//
// # Instance Properties
//
//   - [NSTextContentStorage.IncludesTextListMarkers]
//   - [NSTextContentStorage.SetIncludesTextListMarkers]
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentStorage
type NSTextContentStorage struct {
	NSTextContentManager
}

// NSTextContentStorageFromID constructs a [NSTextContentStorage] from an objc.ID.
//
// A concrete object for managing your view’s text content and generating
// the text elements necessary for layout.
func NSTextContentStorageFromID(id objc.ID) NSTextContentStorage {
	return NSTextContentStorage{NSTextContentManager: NSTextContentManagerFromID(id)}
}
// NOTE: NSTextContentStorage adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSTextContentStorage] class.
//
// # Managing the stored text
//
//   - [INSTextContentStorage.AttributedString]: An attributed string that contains the contents of the document.
//   - [INSTextContentStorage.SetAttributedString]
//
// # Managing text elements
//
//   - [INSTextContentStorage.TextElementForAttributedString]: Returns the text element corresponding to object’s attributed string.
//   - [INSTextContentStorage.AttributedStringForTextElement]: Returns a new attributed string for the text element.
//
// # Instance Properties
//
//   - [INSTextContentStorage.IncludesTextListMarkers]
//   - [INSTextContentStorage.SetIncludesTextListMarkers]
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentStorage
type INSTextContentStorage interface {
	INSTextContentManager
	NSTextStorageObserving

	// Topic: Managing the stored text

	// An attributed string that contains the contents of the document.
	AttributedString() foundation.NSAttributedString
	SetAttributedString(value foundation.NSAttributedString)

	// Topic: Managing text elements

	// Returns the text element corresponding to object’s attributed string.
	TextElementForAttributedString(attributedString foundation.NSAttributedString) INSTextElement
	// Returns a new attributed string for the text element.
	AttributedStringForTextElement(textElement INSTextElement) foundation.NSAttributedString

	// Topic: Instance Properties

	IncludesTextListMarkers() bool
	SetIncludesTextListMarkers(value bool)

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (t NSTextContentStorage) Init() NSTextContentStorage {
	rv := objc.Send[NSTextContentStorage](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextContentStorage) Autorelease() NSTextContentStorage {
	rv := objc.Send[NSTextContentStorage](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextContentStorage creates a new NSTextContentStorage instance.
func NewNSTextContentStorage() NSTextContentStorage {
	class := getNSTextContentStorageClass()
	rv := objc.Send[NSTextContentStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a new content manager object from data in an unarchiver.
//
// coder: An unachiver that conforms to the [NSCoder] class.
// //
// [NSCoder]: https://developer.apple.com/documentation/Foundation/NSCoder
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentManager/init(coder:)
func NewTextContentStorageWithCoder(coder foundation.INSCoder) NSTextContentStorage {
	instance := getNSTextContentStorageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTextContentStorageFromID(rv)
}







// Returns the text element corresponding to object’s attributed string.
//
// attributedString: The attributed string to map into an [NSTextElement].
//
// # Return Value
// 
// An [NSTextElement], or `nil`.
//
// # Discussion
// 
// Returns `nil` when `attributedString` contains attributes not mappable to
// [NSTextElement].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentStorage/textElement(for:)
func (t NSTextContentStorage) TextElementForAttributedString(attributedString foundation.NSAttributedString) INSTextElement {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textElementForAttributedString:"), attributedString)
	return NSTextElementFromID(rv)
}

// Returns a new attributed string for the text element.
//
// textElement: The [NSTextElement] to map into an attributed string.
//
// # Return Value
// 
// An [NSAttributedString], or `nil`.
//
// [NSAttributedString]: https://developer.apple.com/documentation/Foundation/NSAttributedString
//
// # Discussion
// 
// Returns `nil` if the method can’t map `textElement` to an
// [NSAttributedString].
//
// [NSAttributedString]: https://developer.apple.com/documentation/Foundation/NSAttributedString
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentStorage/attributedString(for:)
func (t NSTextContentStorage) AttributedStringForTextElement(textElement INSTextElement) foundation.NSAttributedString {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("attributedStringForTextElement:"), textElement)
	return foundation.NSAttributedStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorageObserving/performEditingTransaction(for:using:)
func (t NSTextContentStorage) PerformEditingTransactionForTextStorageUsingBlock(textStorage NSTextStorage, transaction VoidHandler) {
		_block1, _cleanup1 := NewVoidBlock(transaction)
	defer _cleanup1()
		objc.Send[objc.ID](t.ID, objc.Sel("performEditingTransactionForTextStorage:usingBlock:"), textStorage, _block1)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorageObserving/processEditing(for:edited:range:changeInLength:invalidatedRange:)
func (t NSTextContentStorage) ProcessEditingForTextStorageEditedRangeChangeInLengthInvalidatedRange(textStorage NSTextStorage, editMask NSTextStorageEditActions, newCharRange foundation.NSRange, delta int, invalidatedCharRange foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("processEditingForTextStorage:edited:range:changeInLength:invalidatedRange:"), textStorage, editMask, newCharRange, delta, invalidatedCharRange)
}
func (t NSTextContentStorage) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}












// An attributed string that contains the contents of the document.
//
// # Discussion
// 
// The default value of this property is an [NSTextStorage] object. When you
// need to change the text in your view, fetch this string and make your
// changes to it. When making changes, place them in a block and pass them to
// the [PerformEditingTransactionUsingBlock] method. Wrapping changes in an
// edit transaction gives the rest of the text system an opportunity to
// respond to those changes. For example, the layout manager uses edit
// transactions to update the text layout for any content in the visible
// portion of your view.
// 
// If you assign a new value to this property, the object replaces the current
// string with the one you provide. Don’t set the value of this property to
// `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentStorage/attributedString
func (t NSTextContentStorage) AttributedString() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("attributedString"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (t NSTextContentStorage) SetAttributedString(value foundation.NSAttributedString) {
	objc.Send[struct{}](t.ID, objc.Sel("setAttributedString:"), value)
}



// See: https://developer.apple.com/documentation/AppKit/NSTextContentStorage/includesTextListMarkers
func (t NSTextContentStorage) IncludesTextListMarkers() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("includesTextListMarkers"))
	return rv
}
func (t NSTextContentStorage) SetIncludesTextListMarkers(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setIncludesTextListMarkers:"), value)
}



// See: https://developer.apple.com/documentation/AppKit/NSTextStorageObserving/textStorage
func (t NSTextContentStorage) TextStorage() NSTextStorage {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textStorage"))
	return NSTextStorageFromID(objc.ID(rv))
}
func (t NSTextContentStorage) SetTextStorage(value NSTextStorage) {
	objc.Send[struct{}](t.ID, objc.Sel("setTextStorage:"), value)
}

















			// Protocol methods for NSTextStorageObserving
			










// PerformEditingTransactionForTextStorageUsingBlockSync is a synchronous wrapper around [NSTextContentStorage.PerformEditingTransactionForTextStorageUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (t NSTextContentStorage) PerformEditingTransactionForTextStorageUsingBlockSync(ctx context.Context, textStorage NSTextStorage) error {
	done := make(chan struct{}, 1)
	t.PerformEditingTransactionForTextStorageUsingBlock(textStorage, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}






