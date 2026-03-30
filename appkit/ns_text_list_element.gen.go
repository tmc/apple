// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextListElement] class.
var (
	_NSTextListElementClass     NSTextListElementClass
	_NSTextListElementClassOnce sync.Once
)

func getNSTextListElementClass() NSTextListElementClass {
	_NSTextListElementClassOnce.Do(func() {
		_NSTextListElementClass = NSTextListElementClass{class: objc.GetClass("NSTextListElement")}
	})
	return _NSTextListElementClass
}

// GetNSTextListElementClass returns the class object for NSTextListElement.
func GetNSTextListElementClass() NSTextListElementClass {
	return getNSTextListElementClass()
}

type NSTextListElementClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTextListElementClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextListElementClass) Alloc() NSTextListElement {
	rv := objc.Send[NSTextListElement](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A class that represents a text list node.
//
// # Create a text list element
//
//   - [NSTextListElement.InitWithParentElementTextListContentsMarkerAttributesChildElements]: Creates a text list element with the parent, list elements, nesting level, and marker attributes you provide.
//
// # Accessing the text elements
//
//   - [NSTextListElement.TextList]: The value that represents the text list.
//
// # Accessing the text list’s attributes
//
//   - [NSTextListElement.MarkerAttributes]: A dictionary of attributed string keys and IDs that represent the list’s marker attributes.
//
// # Accessing the formatted string data
//
//   - [NSTextListElement.Contents]: The text list element contents without markers and formatting.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextListElement
type NSTextListElement struct {
	NSTextParagraph
}

// NSTextListElementFromID constructs a [NSTextListElement] from an objc.ID.
//
// A class that represents a text list node.
func NSTextListElementFromID(id objc.ID) NSTextListElement {
	return NSTextListElement{NSTextParagraph: NSTextParagraphFromID(id)}
}

// NOTE: NSTextListElement adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextListElement] class.
//
// # Create a text list element
//
//   - [INSTextListElement.InitWithParentElementTextListContentsMarkerAttributesChildElements]: Creates a text list element with the parent, list elements, nesting level, and marker attributes you provide.
//
// # Accessing the text elements
//
//   - [INSTextListElement.TextList]: The value that represents the text list.
//
// # Accessing the text list’s attributes
//
//   - [INSTextListElement.MarkerAttributes]: A dictionary of attributed string keys and IDs that represent the list’s marker attributes.
//
// # Accessing the formatted string data
//
//   - [INSTextListElement.Contents]: The text list element contents without markers and formatting.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextListElement
type INSTextListElement interface {
	INSTextParagraph

	// Topic: Create a text list element

	// Creates a text list element with the parent, list elements, nesting level, and marker attributes you provide.
	InitWithParentElementTextListContentsMarkerAttributesChildElements(parent INSTextListElement, textList INSTextList, contents foundation.NSAttributedString, markerAttributes foundation.INSDictionary, children []NSTextListElement) NSTextListElement

	// Topic: Accessing the text elements

	// The value that represents the text list.
	TextList() INSTextList

	// Topic: Accessing the text list’s attributes

	// A dictionary of attributed string keys and IDs that represent the list’s marker attributes.
	MarkerAttributes() foundation.INSDictionary

	// Topic: Accessing the formatted string data

	// The text list element contents without markers and formatting.
	Contents() foundation.NSAttributedString
}

// Init initializes the instance.
func (t NSTextListElement) Init() NSTextListElement {
	rv := objc.Send[NSTextListElement](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextListElement) Autorelease() NSTextListElement {
	rv := objc.Send[NSTextListElement](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextListElement creates a new NSTextListElement instance.
func NewNSTextListElement() NSTextListElement {
	class := getNSTextListElementClass()
	rv := objc.Send[NSTextListElement](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new paragraph with the attributed string you provide.
//
// attributedString: An [NSAttributedString].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextParagraph/init(attributedString:)
//
// [NSAttributedString]: https://developer.apple.com/documentation/Foundation/NSAttributedString
func NewTextListElementWithAttributedString(attributedString foundation.NSAttributedString) NSTextListElement {
	instance := getNSTextListElementClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAttributedString:"), attributedString)
	return NSTextListElementFromID(rv)
}

// Creates a text list element with the list elements and nesting level you
// provide.
//
// children: An array of [NSTextListElement] elements.
//
// textList: The [NSTextList] to add elements to.
//
// nestingLevel: An integer value that describes the level of nesting of these elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextListElement/init(children:textList:nestingLevel:)
func NewTextListElementWithChildElementsTextListNestingLevel(children []NSTextListElement, textList INSTextList, nestingLevel int) NSTextListElement {
	rv := objc.Send[objc.ID](objc.ID(getNSTextListElementClass().class), objc.Sel("textListElementWithChildElements:textList:nestingLevel:"), objectivec.IObjectSliceToNSArray(children), textList, nestingLevel)
	return NSTextListElementFromID(rv)
}

// Creates a text list element with the list elements, nesting level, and
// marker attributes you provide.
//
// contents: An [NSAttributedString] that contains the contents of the text list
// element.
//
// markerAttributes: A dictionary of [NSAttributedString.Key] keys and IDs that describe the
// marker attributes.
//
// textList: The [NSTextList] to add elements to.
//
// children: An array of [NSTextListElement] elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextListElement/init(contents:markerAttributes:textList:children:)
//
// [NSAttributedString]: https://developer.apple.com/documentation/Foundation/NSAttributedString
// [NSAttributedString.Key]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key
func NewTextListElementWithContentsMarkerAttributesTextListChildElements(contents foundation.NSAttributedString, markerAttributes foundation.INSDictionary, textList INSTextList, children []NSTextListElement) NSTextListElement {
	rv := objc.Send[objc.ID](objc.ID(getNSTextListElementClass().class), objc.Sel("textListElementWithContents:markerAttributes:textList:childElements:"), contents, markerAttributes, textList, objectivec.IObjectSliceToNSArray(children))
	return NSTextListElementFromID(rv)
}

// Creates a text list element with the parent, list elements, nesting level,
// and marker attributes you provide.
//
// parent: The parent [NSTextListElement] of this element, if any.
//
// textList: The [NSTextList] to add elements to.
//
// contents: An [NSAttributedString] that contains the contents of the text list
// element.
//
// markerAttributes: A dictionary of [NSAttributedString.Key] keys and IDs that describe the
// marker attributes.
//
// children: An array of [NSTextListElement] elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextListElement/init(parent:textList:contents:markerAttributes:children:)
//
// [NSAttributedString]: https://developer.apple.com/documentation/Foundation/NSAttributedString
// [NSAttributedString.Key]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key
func NewTextListElementWithParentElementTextListContentsMarkerAttributesChildElements(parent INSTextListElement, textList INSTextList, contents foundation.NSAttributedString, markerAttributes foundation.INSDictionary, children []NSTextListElement) NSTextListElement {
	instance := getNSTextListElementClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParentElement:textList:contents:markerAttributes:childElements:"), parent, textList, contents, markerAttributes, objectivec.IObjectSliceToNSArray(children))
	return NSTextListElementFromID(rv)
}

// Creates a new text element with the content manager you provide.
//
// textContentManager: The [NSTextContentManager] to apply to this text element.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElement/init(textContentManager:)
func NewTextListElementWithTextContentManager(textContentManager INSTextContentManager) NSTextListElement {
	instance := getNSTextListElementClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTextContentManager:"), textContentManager)
	return NSTextListElementFromID(rv)
}

// Creates a text list element with the parent, list elements, nesting level,
// and marker attributes you provide.
//
// parent: The parent [NSTextListElement] of this element, if any.
//
// textList: The [NSTextList] to add elements to.
//
// contents: An [NSAttributedString] that contains the contents of the text list
// element.
//
// markerAttributes: A dictionary of [NSAttributedString.Key] keys and IDs that describe the
// marker attributes.
//
// children: An array of [NSTextListElement] elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextListElement/init(parent:textList:contents:markerAttributes:children:)
//
// [NSAttributedString]: https://developer.apple.com/documentation/Foundation/NSAttributedString
// [NSAttributedString.Key]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key
func (t NSTextListElement) InitWithParentElementTextListContentsMarkerAttributesChildElements(parent INSTextListElement, textList INSTextList, contents foundation.NSAttributedString, markerAttributes foundation.INSDictionary, children []NSTextListElement) NSTextListElement {
	rv := objc.Send[NSTextListElement](t.ID, objc.Sel("initWithParentElement:textList:contents:markerAttributes:childElements:"), parent, textList, contents, markerAttributes, objectivec.IObjectSliceToNSArray(children))
	return rv
}

// The value that represents the text list.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextListElement/textList
func (t NSTextListElement) TextList() INSTextList {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textList"))
	return NSTextListFromID(objc.ID(rv))
}

// A dictionary of attributed string keys and IDs that represent the list’s
// marker attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextListElement/markerAttributes
func (t NSTextListElement) MarkerAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("markerAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The text list element contents without markers and formatting.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextListElement/contents
func (t NSTextListElement) Contents() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("contents"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
