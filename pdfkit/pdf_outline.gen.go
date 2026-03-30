// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PDFOutline] class.
var (
	_PDFOutlineClass     PDFOutlineClass
	_PDFOutlineClassOnce sync.Once
)

func getPDFOutlineClass() PDFOutlineClass {
	_PDFOutlineClassOnce.Do(func() {
		_PDFOutlineClass = PDFOutlineClass{class: objc.GetClass("PDFOutline")}
	})
	return _PDFOutlineClass
}

// GetPDFOutlineClass returns the class object for PDFOutline.
func GetPDFOutlineClass() PDFOutlineClass {
	return getPDFOutlineClass()
}

type PDFOutlineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PDFOutlineClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PDFOutlineClass) Alloc() PDFOutline {
	rv := objc.Send[PDFOutline](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// A [PDFOutline] object is an element in a tree-structured hierarchy that can
// represent the structure of a PDF document.
//
// # Overview
//
// An outline is an optional component of a PDF document, useful for viewing
// the structure of the document and for navigating within it.
//
// Outlines are created by the document’s author. If you represent a PDF
// document outline using outline objects, the root of the hierarchy is
// obtained from the PDF document itself. This root outline is not visible and
// serves merely as a container for the visible outlines.
//
// # Getting Information About an Outline
//
//   - [PDFOutline.Document]: Returns the document with which the outline is associated.
//   - [PDFOutline.NumberOfChildren]: Returns the number of child outline objects in the outline.
//   - [PDFOutline.Parent]: Returns the parent outline object of the outline (returns [NULL] if called on the root outline object).
//   - [PDFOutline.ChildAtIndex]: Returns the child outline object at the specified index.
//   - [PDFOutline.Index]: Returns the index of the outline.
//
// # Managing Outline Labels
//
//   - [PDFOutline.Label]: Returns the label for the outline.
//   - [PDFOutline.SetLabel]
//
// # Managing Actions and Destinations
//
//   - [PDFOutline.Destination]: Returns the destination associated with the outline.
//   - [PDFOutline.SetDestination]
//   - [PDFOutline.Action]: Returns the action performed when users click the outline.
//   - [PDFOutline.SetAction]
//
// # Changing an Outline Hierarchy
//
//   - [PDFOutline.InsertChildAtIndex]: Inserts the specified outline object at the specified index.
//   - [PDFOutline.RemoveFromParent]: Removes the outline object from its parent (does nothing if outline object is the root outline object).
//
// # Managing the Disclosure of an Outline Object
//
//   - [PDFOutline.IsOpen]: Returns a Boolean value that indicates whether the outline object is initially disclosed.
//   - [PDFOutline.SetIsOpen]
//
// See: https://developer.apple.com/documentation/PDFKit/PDFOutline
type PDFOutline struct {
	objectivec.Object
}

// PDFOutlineFromID constructs a [PDFOutline] from an objc.ID.
//
// A [PDFOutline] object is an element in a tree-structured hierarchy that can
// represent the structure of a PDF document.
func PDFOutlineFromID(id objc.ID) PDFOutline {
	return PDFOutline{objectivec.Object{ID: id}}
}

// NOTE: PDFOutline adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [PDFOutline] class.
//
// # Getting Information About an Outline
//
//   - [IPDFOutline.Document]: Returns the document with which the outline is associated.
//   - [IPDFOutline.NumberOfChildren]: Returns the number of child outline objects in the outline.
//   - [IPDFOutline.Parent]: Returns the parent outline object of the outline (returns [NULL] if called on the root outline object).
//   - [IPDFOutline.ChildAtIndex]: Returns the child outline object at the specified index.
//   - [IPDFOutline.Index]: Returns the index of the outline.
//
// # Managing Outline Labels
//
//   - [IPDFOutline.Label]: Returns the label for the outline.
//   - [IPDFOutline.SetLabel]
//
// # Managing Actions and Destinations
//
//   - [IPDFOutline.Destination]: Returns the destination associated with the outline.
//   - [IPDFOutline.SetDestination]
//   - [IPDFOutline.Action]: Returns the action performed when users click the outline.
//   - [IPDFOutline.SetAction]
//
// # Changing an Outline Hierarchy
//
//   - [IPDFOutline.InsertChildAtIndex]: Inserts the specified outline object at the specified index.
//   - [IPDFOutline.RemoveFromParent]: Removes the outline object from its parent (does nothing if outline object is the root outline object).
//
// # Managing the Disclosure of an Outline Object
//
//   - [IPDFOutline.IsOpen]: Returns a Boolean value that indicates whether the outline object is initially disclosed.
//   - [IPDFOutline.SetIsOpen]
//
// See: https://developer.apple.com/documentation/PDFKit/PDFOutline
type IPDFOutline interface {
	objectivec.IObject

	// Topic: Getting Information About an Outline

	// Returns the document with which the outline is associated.
	Document() IPDFDocument
	// Returns the number of child outline objects in the outline.
	NumberOfChildren() uint
	// Returns the parent outline object of the outline (returns [NULL] if called on the root outline object).
	Parent() IPDFOutline
	// Returns the child outline object at the specified index.
	ChildAtIndex(index uint) IPDFOutline
	// Returns the index of the outline.
	Index() uint

	// Topic: Managing Outline Labels

	// Returns the label for the outline.
	Label() string
	SetLabel(value string)

	// Topic: Managing Actions and Destinations

	// Returns the destination associated with the outline.
	Destination() IPDFDestination
	SetDestination(value IPDFDestination)
	// Returns the action performed when users click the outline.
	Action() IPDFAction
	SetAction(value IPDFAction)

	// Topic: Changing an Outline Hierarchy

	// Inserts the specified outline object at the specified index.
	InsertChildAtIndex(child IPDFOutline, index uint)
	// Removes the outline object from its parent (does nothing if outline object is the root outline object).
	RemoveFromParent()

	// Topic: Managing the Disclosure of an Outline Object

	// Returns a Boolean value that indicates whether the outline object is initially disclosed.
	IsOpen() bool
	SetIsOpen(value bool)
}

// Init initializes the instance.
func (p PDFOutline) Init() PDFOutline {
	rv := objc.Send[PDFOutline](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PDFOutline) Autorelease() PDFOutline {
	rv := objc.Send[PDFOutline](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFOutline creates a new PDFOutline instance.
func NewPDFOutline() PDFOutline {
	class := getPDFOutlineClass()
	rv := objc.Send[PDFOutline](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the child outline object at the specified index.
//
// # Discussion
//
// The index is zero-based. This method throws an exception if `index` is out
// of range.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFOutline/child(at:)
func (p PDFOutline) ChildAtIndex(index uint) IPDFOutline {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("childAtIndex:"), index)
	return PDFOutlineFromID(rv)
}

// Inserts the specified outline object at the specified index.
//
// # Discussion
//
// To build a PDF outline hierarchy, use this method to add child outline
// objects. Before you call this method on a [PDFOutline] object that already
// has a parent, you should retain the object and call [RemoveFromParent] on
// it first.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFOutline/insertChild(_:at:)
func (p PDFOutline) InsertChildAtIndex(child IPDFOutline, index uint) {
	objc.Send[objc.ID](p.ID, objc.Sel("insertChild:atIndex:"), child, index)
}

// Removes the outline object from its parent (does nothing if outline object
// is the root outline object).
//
// See: https://developer.apple.com/documentation/PDFKit/PDFOutline/removeFromParent()
func (p PDFOutline) RemoveFromParent() {
	objc.Send[objc.ID](p.ID, objc.Sel("removeFromParent"))
}

// Returns the document with which the outline is associated.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFOutline/document
func (p PDFOutline) Document() IPDFDocument {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("document"))
	return PDFDocumentFromID(objc.ID(rv))
}

// Returns the number of child outline objects in the outline.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFOutline/numberOfChildren
func (p PDFOutline) NumberOfChildren() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("numberOfChildren"))
	return rv
}

// Returns the parent outline object of the outline (returns [NULL] if called
// on the root outline object).
//
// See: https://developer.apple.com/documentation/PDFKit/PDFOutline/parent
func (p PDFOutline) Parent() IPDFOutline {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("parent"))
	return PDFOutlineFromID(objc.ID(rv))
}

// Returns the index of the outline.
//
// # Discussion
//
// The index of the outline object is relative to its siblings and from the
// perspective of the parent of the outline object. The root outline object,
// and any outline object without a parent, has an index value of `0`.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFOutline/index
func (p PDFOutline) Index() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("index"))
	return rv
}

// Returns the label for the outline.
//
// # Discussion
//
// The root outline serves only as a container for the outlines it owns; it
// does not have a label.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFOutline/label
func (p PDFOutline) Label() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (p PDFOutline) SetLabel(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setLabel:"), objc.String(value))
}

// Returns the destination associated with the outline.
//
// # Discussion
//
// The root outline serves only as a container for the outlines it owns; it
// does not have a destination. Note that a [PDFOutline] object can have
// either a destination or an action, not both.
//
// This method may return [NULL] if the outline has an associated action
// instead of a destination. Note that if the associated action is a
// [PDFActionGoTo], this method returns the destination from the
// [PDFActionGoTo] object. However, it is better to use the [Action] method
// for this purpose.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFOutline/destination
func (p PDFOutline) Destination() IPDFDestination {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("destination"))
	return PDFDestinationFromID(objc.ID(rv))
}
func (p PDFOutline) SetDestination(value IPDFDestination) {
	objc.Send[struct{}](p.ID, objc.Sel("setDestination:"), value)
}

// Returns the action performed when users click the outline.
//
// # Discussion
//
// The root outline serves only as a container for the outlines it owns; it
// does not have an action. Note that a [PDFOutline] object can have either an
// action or a destination, not both.
//
// If the [PDFOutline] object has a destination, instead of an action,
// `action` returns a [PDFActionGoTo] object (this is equivalent to calling
// [Destination] on the [PDFOutline] object). For other action types, `action`
// returns the appropriate PDF Kit action type object, such as [PDFActionURL].
//
// See: https://developer.apple.com/documentation/PDFKit/PDFOutline/action
func (p PDFOutline) Action() IPDFAction {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("action"))
	return PDFActionFromID(objc.ID(rv))
}
func (p PDFOutline) SetAction(value IPDFAction) {
	objc.Send[struct{}](p.ID, objc.Sel("setAction:"), value)
}

// Returns a Boolean value that indicates whether the outline object is
// initially disclosed.
//
// # Discussion
//
// Calling `isOpen` on an outline object that has no children always returns
// false. Calling `isOpen` on the root outline object always returns true.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFOutline/isOpen
func (p PDFOutline) IsOpen() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isOpen"))
	return rv
}
func (p PDFOutline) SetIsOpen(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setIsOpen:"), value)
}
