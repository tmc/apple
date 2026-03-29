// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSTextParagraph] class.
var (
	_NSTextParagraphClass     NSTextParagraphClass
	_NSTextParagraphClassOnce sync.Once
)

func getNSTextParagraphClass() NSTextParagraphClass {
	_NSTextParagraphClassOnce.Do(func() {
		_NSTextParagraphClass = NSTextParagraphClass{class: objc.GetClass("NSTextParagraph")}
	})
	return _NSTextParagraphClass
}

// GetNSTextParagraphClass returns the class object for NSTextParagraph.
func GetNSTextParagraphClass() NSTextParagraphClass {
	return getNSTextParagraphClass()
}

type NSTextParagraphClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTextParagraphClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextParagraphClass) Alloc() NSTextParagraph {
	rv := objc.Send[NSTextParagraph](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A class that represents a single paragraph backed by an attributed string
// as the contents.
//
// # Creating a paragraph
//
//   - [NSTextParagraph.InitWithAttributedString]: Creates a new paragraph with the attributed string you provide.
//
// # Getting paragraph characteristics
//
//   - [NSTextParagraph.AttributedString]: Returns the source attributed string.
//   - [NSTextParagraph.ParagraphContentRange]: Returns the range of the paragraph in the containing text’s attributed string.
//   - [NSTextParagraph.ParagraphSeparatorRange]: Returns the range of the paragraph separator in the containing text’s attributed string.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextParagraph
type NSTextParagraph struct {
	NSTextElement
}

// NSTextParagraphFromID constructs a [NSTextParagraph] from an objc.ID.
//
// A class that represents a single paragraph backed by an attributed string
// as the contents.
func NSTextParagraphFromID(id objc.ID) NSTextParagraph {
	return NSTextParagraph{NSTextElement: NSTextElementFromID(id)}
}
// NOTE: NSTextParagraph adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextParagraph] class.
//
// # Creating a paragraph
//
//   - [INSTextParagraph.InitWithAttributedString]: Creates a new paragraph with the attributed string you provide.
//
// # Getting paragraph characteristics
//
//   - [INSTextParagraph.AttributedString]: Returns the source attributed string.
//   - [INSTextParagraph.ParagraphContentRange]: Returns the range of the paragraph in the containing text’s attributed string.
//   - [INSTextParagraph.ParagraphSeparatorRange]: Returns the range of the paragraph separator in the containing text’s attributed string.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextParagraph
type INSTextParagraph interface {
	INSTextElement

	// Topic: Creating a paragraph

	// Creates a new paragraph with the attributed string you provide.
	InitWithAttributedString(attributedString foundation.NSAttributedString) NSTextParagraph

	// Topic: Getting paragraph characteristics

	// Returns the source attributed string.
	AttributedString() foundation.NSAttributedString
	// Returns the range of the paragraph in the containing text’s attributed string.
	ParagraphContentRange() INSTextRange
	// Returns the range of the paragraph separator in the containing text’s attributed string.
	ParagraphSeparatorRange() INSTextRange
}

// Init initializes the instance.
func (t NSTextParagraph) Init() NSTextParagraph {
	rv := objc.Send[NSTextParagraph](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextParagraph) Autorelease() NSTextParagraph {
	rv := objc.Send[NSTextParagraph](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextParagraph creates a new NSTextParagraph instance.
func NewNSTextParagraph() NSTextParagraph {
	class := getNSTextParagraphClass()
	rv := objc.Send[NSTextParagraph](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new paragraph with the attributed string you provide.
//
// attributedString: An [NSAttributedString].
// //
// [NSAttributedString]: https://developer.apple.com/documentation/Foundation/NSAttributedString
//
// See: https://developer.apple.com/documentation/AppKit/NSTextParagraph/init(attributedString:)
func NewTextParagraphWithAttributedString(attributedString foundation.NSAttributedString) NSTextParagraph {
	instance := getNSTextParagraphClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAttributedString:"), attributedString)
	return NSTextParagraphFromID(rv)
}

// Creates a new text element with the content manager you provide.
//
// textContentManager: The [NSTextContentManager] to apply to this text element.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElement/init(textContentManager:)
func NewTextParagraphWithTextContentManager(textContentManager INSTextContentManager) NSTextParagraph {
	instance := getNSTextParagraphClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTextContentManager:"), textContentManager)
	return NSTextParagraphFromID(rv)
}

// Creates a new paragraph with the attributed string you provide.
//
// attributedString: An [NSAttributedString].
// //
// [NSAttributedString]: https://developer.apple.com/documentation/Foundation/NSAttributedString
//
// See: https://developer.apple.com/documentation/AppKit/NSTextParagraph/init(attributedString:)
func (t NSTextParagraph) InitWithAttributedString(attributedString foundation.NSAttributedString) NSTextParagraph {
	rv := objc.Send[NSTextParagraph](t.ID, objc.Sel("initWithAttributedString:"), attributedString)
	return rv
}

// Returns the source attributed string.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextParagraph/attributedString
func (t NSTextParagraph) AttributedString() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("attributedString"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
// Returns the range of the paragraph in the containing text’s attributed
// string.
//
// # Discussion
// 
// The containing text is [NSTextContentStorage]’s [AttributedString].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextParagraph/paragraphContentRange
func (t NSTextParagraph) ParagraphContentRange() INSTextRange {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("paragraphContentRange"))
	return NSTextRangeFromID(objc.ID(rv))
}
// Returns the range of the paragraph separator in the containing text’s
// attributed string.
//
// # Discussion
// 
// The containing text is [NSTextContentStorage]’s [AttributedString].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextParagraph/paragraphSeparatorRange
func (t NSTextParagraph) ParagraphSeparatorRange() INSTextRange {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("paragraphSeparatorRange"))
	return NSTextRangeFromID(objc.ID(rv))
}

