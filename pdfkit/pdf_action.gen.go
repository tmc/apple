// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PDFAction] class.
var (
	_PDFActionClass     PDFActionClass
	_PDFActionClassOnce sync.Once
)

func getPDFActionClass() PDFActionClass {
	_PDFActionClassOnce.Do(func() {
		_PDFActionClass = PDFActionClass{class: objc.GetClass("PDFAction")}
	})
	return _PDFActionClass
}

// GetPDFActionClass returns the class object for PDFAction.
func GetPDFActionClass() PDFActionClass {
	return getPDFActionClass()
}

type PDFActionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PDFActionClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PDFActionClass) Alloc() PDFAction {
	rv := objc.Send[PDFAction](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// An action that is performed when, for example, a PDF annotation is
// activated or an outline item is clicked.
//
// # Overview
//
// A [PDFAction] object represents an action associated with a PDF element,
// such as an annotation or a link, that the viewer application can perform.
// See the Adobe PDF Specification for more about actions and action types.
//
// [PDFAction] is an abstract superclass of the following concrete classes:
//
// - [PDFActionGoTo] - [PDFActionNamed] - [PDFActionRemoteGoTo] -
// [PDFActionResetForm] - [PDFActionURL]
//
// # Getting the Action Type
//
//   - [PDFAction.Type]: Returns the type of the action.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAction
type PDFAction struct {
	objectivec.Object
}

// PDFActionFromID constructs a [PDFAction] from an objc.ID.
//
// An action that is performed when, for example, a PDF annotation is
// activated or an outline item is clicked.
func PDFActionFromID(id objc.ID) PDFAction {
	return PDFAction{objectivec.Object{ID: id}}
}

// NOTE: PDFAction adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [PDFAction] class.
//
// # Getting the Action Type
//
//   - [IPDFAction.Type]: Returns the type of the action.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAction
type IPDFAction interface {
	objectivec.IObject

	// Topic: Getting the Action Type

	// Returns the type of the action.
	Type() string

	// An object that represents an action for a PDF element, such as a link annotation.
	Action() IPDFAction
	SetAction(value IPDFAction)
	// Returns the modification date of the annotation.
	ModificationDate() foundation.INSDate
	SetModificationDate(value foundation.INSDate)
	// Returns the page that the annotation is associated with.
	Page() IPDFPage
	SetPage(value IPDFPage)
	// Returns the name of the user who created the annotation.
	UserName() string
	SetUserName(value string)
}

// Init initializes the instance.
func (p PDFAction) Init() PDFAction {
	rv := objc.Send[PDFAction](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PDFAction) Autorelease() PDFAction {
	rv := objc.Send[PDFAction](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAction creates a new PDFAction instance.
func NewPDFAction() PDFAction {
	class := getPDFActionClass()
	rv := objc.Send[PDFAction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the type of the action.
//
// # Return Value
//
// The type of the PDF action.
//
// # Discussion
//
// The PDF action type returned by this method may not correspond precisely to
// the name of a [PDFAction] subclass. For example, a [PDFActionURL] object
// might return “URI” or “Launch,” depending on the original action as
// defined by the Adobe PDF Specification. In the PDF Kit, these two actions
// are handled in the single [PDFActionURL] subclass, and the more familiar
// term “URL” is used instead.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFAction/type
func (p PDFAction) Type() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("type"))
	return foundation.NSStringFromID(rv).String()
}

// An object that represents an action for a PDF element, such as a link
// annotation.
//
// See: https://developer.apple.com/documentation/pdfkit/pdfannotation/action
func (p PDFAction) Action() IPDFAction {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("action"))
	return PDFActionFromID(objc.ID(rv))
}
func (p PDFAction) SetAction(value IPDFAction) {
	objc.Send[struct{}](p.ID, objc.Sel("setAction:"), value)
}

// Returns the modification date of the annotation.
//
// See: https://developer.apple.com/documentation/pdfkit/pdfannotation/modificationdate
func (p PDFAction) ModificationDate() foundation.INSDate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("modificationDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (p PDFAction) SetModificationDate(value foundation.INSDate) {
	objc.Send[struct{}](p.ID, objc.Sel("setModificationDate:"), value)
}

// Returns the page that the annotation is associated with.
//
// See: https://developer.apple.com/documentation/pdfkit/pdfannotation/page
func (p PDFAction) Page() IPDFPage {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("page"))
	return PDFPageFromID(objc.ID(rv))
}
func (p PDFAction) SetPage(value IPDFPage) {
	objc.Send[struct{}](p.ID, objc.Sel("setPage:"), value)
}

// Returns the name of the user who created the annotation.
//
// See: https://developer.apple.com/documentation/pdfkit/pdfannotation/username
func (p PDFAction) UserName() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("userName"))
	return foundation.NSStringFromID(rv).String()
}
func (p PDFAction) SetUserName(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setUserName:"), objc.String(value))
}
