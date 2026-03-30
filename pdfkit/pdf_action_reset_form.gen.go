// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PDFActionResetForm] class.
var (
	_PDFActionResetFormClass     PDFActionResetFormClass
	_PDFActionResetFormClassOnce sync.Once
)

func getPDFActionResetFormClass() PDFActionResetFormClass {
	_PDFActionResetFormClassOnce.Do(func() {
		_PDFActionResetFormClass = PDFActionResetFormClass{class: objc.GetClass("PDFActionResetForm")}
	})
	return _PDFActionResetFormClass
}

// GetPDFActionResetFormClass returns the class object for PDFActionResetForm.
func GetPDFActionResetFormClass() PDFActionResetFormClass {
	return getPDFActionResetFormClass()
}

type PDFActionResetFormClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PDFActionResetFormClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PDFActionResetFormClass) Alloc() PDFActionResetForm {
	rv := objc.Send[PDFActionResetForm](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// [PDFActionResetForm], a subclass of [PDFAction], defines methods for
// getting and clearing fields in a PDF form.
//
// # Overview
//
// A [PDFActionResetForm] object represents an action associated with a PDF
// form.
//
// # Accessing and Changing Fields
//
//   - [PDFActionResetForm.Fields]: Returns an array of fields associated with the reset action.
//   - [PDFActionResetForm.SetFields]
//
// # Determining Whether Fields are Cleared When the Action is Performed
//
//   - [PDFActionResetForm.FieldsIncludedAreCleared]: Sets whether the fields associated with the reset action are cleared when the action is performed.
//   - [PDFActionResetForm.SetFieldsIncludedAreCleared]
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionResetForm
type PDFActionResetForm struct {
	PDFAction
}

// PDFActionResetFormFromID constructs a [PDFActionResetForm] from an objc.ID.
//
// [PDFActionResetForm], a subclass of [PDFAction], defines methods for
// getting and clearing fields in a PDF form.
func PDFActionResetFormFromID(id objc.ID) PDFActionResetForm {
	return PDFActionResetForm{PDFAction: PDFActionFromID(id)}
}

// NOTE: PDFActionResetForm adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [PDFActionResetForm] class.
//
// # Accessing and Changing Fields
//
//   - [IPDFActionResetForm.Fields]: Returns an array of fields associated with the reset action.
//   - [IPDFActionResetForm.SetFields]
//
// # Determining Whether Fields are Cleared When the Action is Performed
//
//   - [IPDFActionResetForm.FieldsIncludedAreCleared]: Sets whether the fields associated with the reset action are cleared when the action is performed.
//   - [IPDFActionResetForm.SetFieldsIncludedAreCleared]
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionResetForm
type IPDFActionResetForm interface {
	IPDFAction

	// Topic: Accessing and Changing Fields

	// Returns an array of fields associated with the reset action.
	Fields() []string
	SetFields(value []string)

	// Topic: Determining Whether Fields are Cleared When the Action is Performed

	// Sets whether the fields associated with the reset action are cleared when the action is performed.
	FieldsIncludedAreCleared() bool
	SetFieldsIncludedAreCleared(value bool)
}

// Init initializes the instance.
func (p PDFActionResetForm) Init() PDFActionResetForm {
	rv := objc.Send[PDFActionResetForm](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PDFActionResetForm) Autorelease() PDFActionResetForm {
	rv := objc.Send[PDFActionResetForm](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFActionResetForm creates a new PDFActionResetForm instance.
func NewPDFActionResetForm() PDFActionResetForm {
	class := getPDFActionResetFormClass()
	rv := objc.Send[PDFActionResetForm](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an array of fields associated with the reset action.
//
// # Return Value
//
// An array of [NSString] objects that corresponds to the `fieldNames`
// property of widget annotations (such as [PDFAnnotationButtonWidget]) on the
// PDF page. This method can return [NULL].
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionResetForm/fields
//
// [PDFAnnotationButtonWidget]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationButtonWidget
func (p PDFActionResetForm) Fields() []string {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("fields"))
	return objc.ConvertSliceToStrings(rv)
}
func (p PDFActionResetForm) SetFields(value []string) {
	objc.Send[struct{}](p.ID, objc.Sel("setFields:"), objectivec.StringSliceToNSArray(value))
}

// Sets whether the fields associated with the reset action are cleared when
// the action is performed.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFActionResetForm/fieldsIncludedAreCleared
func (p PDFActionResetForm) FieldsIncludedAreCleared() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("fieldsIncludedAreCleared"))
	return rv
}
func (p PDFActionResetForm) SetFieldsIncludedAreCleared(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setFieldsIncludedAreCleared:"), value)
}
