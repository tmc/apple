// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextElement] class.
var (
	_NSTextElementClass     NSTextElementClass
	_NSTextElementClassOnce sync.Once
)

func getNSTextElementClass() NSTextElementClass {
	_NSTextElementClassOnce.Do(func() {
		_NSTextElementClass = NSTextElementClass{class: objc.GetClass("NSTextElement")}
	})
	return _NSTextElementClass
}

// GetNSTextElementClass returns the class object for NSTextElement.
func GetNSTextElementClass() NSTextElementClass {
	return getNSTextElementClass()
}

type NSTextElementClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTextElementClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextElementClass) Alloc() NSTextElement {
	rv := objc.Send[NSTextElement](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An abstract base class that represents the smallest units of text layout
// such as paragraphs or attachments.
//
// # Creating a text element
//
//   - [NSTextElement.InitWithTextContentManager]: Creates a new text element with the content manager you provide.
//
// # Accessing the content manager
//
//   - [NSTextElement.TextContentManager]: The value that represents the current content manager.
//   - [NSTextElement.SetTextContentManager]
//
// # Accessing the text element range
//
//   - [NSTextElement.ElementRange]: A range value that represents the range of the element inside the document.
//   - [NSTextElement.SetElementRange]
//
// # Accessing text elements
//
//   - [NSTextElement.IsRepresentedElement]: A Boolean value that indicates whether this element is in the text layout.
//   - [NSTextElement.ParentElement]: A value that represents the parent element if this text element is a child of an enclosing element.
//   - [NSTextElement.ChildElements]: An array of zero or more child text elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElement
type NSTextElement struct {
	objectivec.Object
}

// NSTextElementFromID constructs a [NSTextElement] from an objc.ID.
//
// An abstract base class that represents the smallest units of text layout
// such as paragraphs or attachments.
func NSTextElementFromID(id objc.ID) NSTextElement {
	return NSTextElement{objectivec.Object{ID: id}}
}

// NOTE: NSTextElement adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextElement] class.
//
// # Creating a text element
//
//   - [INSTextElement.InitWithTextContentManager]: Creates a new text element with the content manager you provide.
//
// # Accessing the content manager
//
//   - [INSTextElement.TextContentManager]: The value that represents the current content manager.
//   - [INSTextElement.SetTextContentManager]
//
// # Accessing the text element range
//
//   - [INSTextElement.ElementRange]: A range value that represents the range of the element inside the document.
//   - [INSTextElement.SetElementRange]
//
// # Accessing text elements
//
//   - [INSTextElement.IsRepresentedElement]: A Boolean value that indicates whether this element is in the text layout.
//   - [INSTextElement.ParentElement]: A value that represents the parent element if this text element is a child of an enclosing element.
//   - [INSTextElement.ChildElements]: An array of zero or more child text elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElement
type INSTextElement interface {
	objectivec.IObject

	// Topic: Creating a text element

	// Creates a new text element with the content manager you provide.
	InitWithTextContentManager(textContentManager INSTextContentManager) NSTextElement

	// Topic: Accessing the content manager

	// The value that represents the current content manager.
	TextContentManager() INSTextContentManager
	SetTextContentManager(value INSTextContentManager)

	// Topic: Accessing the text element range

	// A range value that represents the range of the element inside the document.
	ElementRange() INSTextRange
	SetElementRange(value INSTextRange)

	// Topic: Accessing text elements

	// A Boolean value that indicates whether this element is in the text layout.
	IsRepresentedElement() bool
	// A value that represents the parent element if this text element is a child of an enclosing element.
	ParentElement() INSTextElement
	// An array of zero or more child text elements.
	ChildElements() []NSTextElement
}

// Init initializes the instance.
func (t NSTextElement) Init() NSTextElement {
	rv := objc.Send[NSTextElement](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextElement) Autorelease() NSTextElement {
	rv := objc.Send[NSTextElement](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextElement creates a new NSTextElement instance.
func NewNSTextElement() NSTextElement {
	class := getNSTextElementClass()
	rv := objc.Send[NSTextElement](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new text element with the content manager you provide.
//
// textContentManager: The [NSTextContentManager] to apply to this text element.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElement/init(textContentManager:)
func NewTextElementWithTextContentManager(textContentManager INSTextContentManager) NSTextElement {
	instance := getNSTextElementClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTextContentManager:"), textContentManager)
	return NSTextElementFromID(rv)
}

// Creates a new text element with the content manager you provide.
//
// textContentManager: The [NSTextContentManager] to apply to this text element.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElement/init(textContentManager:)
func (t NSTextElement) InitWithTextContentManager(textContentManager INSTextContentManager) NSTextElement {
	rv := objc.Send[NSTextElement](t.ID, objc.Sel("initWithTextContentManager:"), textContentManager)
	return rv
}

// The value that represents the current content manager.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElement/textContentManager
func (t NSTextElement) TextContentManager() INSTextContentManager {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textContentManager"))
	return NSTextContentManagerFromID(objc.ID(rv))
}
func (t NSTextElement) SetTextContentManager(value INSTextContentManager) {
	objc.Send[struct{}](t.ID, objc.Sel("setTextContentManager:"), value)
}

// A range value that represents the range of the element inside the document.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElement/elementRange
func (t NSTextElement) ElementRange() INSTextRange {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("elementRange"))
	return NSTextRangeFromID(objc.ID(rv))
}
func (t NSTextElement) SetElementRange(value INSTextRange) {
	objc.Send[struct{}](t.ID, objc.Sel("setElementRange:"), value)
}

// A Boolean value that indicates whether this element is in the text layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElement/isRepresentedElement
func (t NSTextElement) IsRepresentedElement() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isRepresentedElement"))
	return rv
}

// A value that represents the parent element if this text element is a child
// of an enclosing element.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElement/parent
func (t NSTextElement) ParentElement() INSTextElement {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("parentElement"))
	return NSTextElementFromID(objc.ID(rv))
}

// An array of zero or more child text elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElement/childElements
func (t NSTextElement) ChildElements() []NSTextElement {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("childElements"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSTextElement {
		return NSTextElementFromID(id)
	})
}
