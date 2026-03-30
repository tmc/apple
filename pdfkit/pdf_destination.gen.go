// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PDFDestination] class.
var (
	_PDFDestinationClass     PDFDestinationClass
	_PDFDestinationClassOnce sync.Once
)

func getPDFDestinationClass() PDFDestinationClass {
	_PDFDestinationClassOnce.Do(func() {
		_PDFDestinationClass = PDFDestinationClass{class: objc.GetClass("PDFDestination")}
	})
	return _PDFDestinationClass
}

// GetPDFDestinationClass returns the class object for PDFDestination.
func GetPDFDestinationClass() PDFDestinationClass {
	return getPDFDestinationClass()
}

type PDFDestinationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PDFDestinationClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PDFDestinationClass) Alloc() PDFDestination {
	rv := objc.Send[PDFDestination](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// A [PDFDestination] object describes a point on a PDF page.
//
// # Overview
//
// In typical usage, you do not initialize [PDFDestination] objects but rather
// get them as either attributes of [PDFAnnotationLink] or [PDFOutline]
// objects, or in response to the [PDFView] method [PDFDestination.CurrentDestination].
//
// # Initializing a Destination
//
//   - [PDFDestination.InitWithPageAtPoint]: Initializes the destination.
//
// # Getting Pages and Points
//
//   - [PDFDestination.Page]: Returns the page that the destination refers to.
//   - [PDFDestination.Point]: Returns the point, in page space, that the destination refers to.
//   - [PDFDestination.KPDFDestinationUnspecifiedValue]
//
// # Getting a Relative Location
//
//   - [PDFDestination.Compare]: Returns a comparison result that indicates the location of the destination in the document, relative to the current position.
//
// # Instance Properties
//
//   - [PDFDestination.Zoom]
//   - [PDFDestination.SetZoom]
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDestination
//
// [PDFAnnotationLink]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationLink
type PDFDestination struct {
	objectivec.Object
}

// PDFDestinationFromID constructs a [PDFDestination] from an objc.ID.
//
// A [PDFDestination] object describes a point on a PDF page.
func PDFDestinationFromID(id objc.ID) PDFDestination {
	return PDFDestination{objectivec.Object{ID: id}}
}

// NOTE: PDFDestination adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [PDFDestination] class.
//
// # Initializing a Destination
//
//   - [IPDFDestination.InitWithPageAtPoint]: Initializes the destination.
//
// # Getting Pages and Points
//
//   - [IPDFDestination.Page]: Returns the page that the destination refers to.
//   - [IPDFDestination.Point]: Returns the point, in page space, that the destination refers to.
//   - [IPDFDestination.KPDFDestinationUnspecifiedValue]
//
// # Getting a Relative Location
//
//   - [IPDFDestination.Compare]: Returns a comparison result that indicates the location of the destination in the document, relative to the current position.
//
// # Instance Properties
//
//   - [IPDFDestination.Zoom]
//   - [IPDFDestination.SetZoom]
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDestination
type IPDFDestination interface {
	objectivec.IObject

	// Topic: Initializing a Destination

	// Initializes the destination.
	InitWithPageAtPoint(page IPDFPage, point corefoundation.CGPoint) PDFDestination

	// Topic: Getting Pages and Points

	// Returns the page that the destination refers to.
	Page() IPDFPage
	// Returns the point, in page space, that the destination refers to.
	Point() corefoundation.CGPoint
	KPDFDestinationUnspecifiedValue() float64

	// Topic: Getting a Relative Location

	// Returns a comparison result that indicates the location of the destination in the document, relative to the current position.
	Compare(destination IPDFDestination) foundation.NSComparisonResult

	// Topic: Instance Properties

	Zoom() float64
	SetZoom(value float64)

	// An object that represents an action for a PDF element, such as a link annotation.
	Action() IPDFAction
	SetAction(value IPDFAction)
	// Returns a
	CurrentDestination() IPDFDestination
	SetCurrentDestination(value IPDFDestination)
	// Returns the modification date of the annotation.
	ModificationDate() foundation.INSDate
	SetModificationDate(value foundation.INSDate)
	// Returns the type of the annotation.
	Type() string
	SetType(value string)
	// Returns the name of the user who created the annotation.
	UserName() string
	SetUserName(value string)
}

// Init initializes the instance.
func (p PDFDestination) Init() PDFDestination {
	rv := objc.Send[PDFDestination](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PDFDestination) Autorelease() PDFDestination {
	rv := objc.Send[PDFDestination](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFDestination creates a new PDFDestination instance.
func NewPDFDestination() PDFDestination {
	class := getPDFDestinationClass()
	rv := objc.Send[PDFDestination](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes the destination.
//
// page: The page of the destination.
//
// point: The point of the destination, in page space.
//
// # Return Value
//
// An initialized [PDFDestination] instance, or [NULL] if the object could not
// be initialized.
//
// # Discussion
//
// Specify `point` in page space. Typically, there’s no need to initialize
// destinations. Instead, you get them from [PDFAnnotationLink], [PDFOutline],
// or [PDFView] objects.
//
// Page space is a 72-dpi coordinate system with the origin at the lower-left
// corner of the current page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDestination/init(page:at:)
//
// [PDFAnnotationLink]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationLink
func NewPDFDestinationWithPageAtPoint(page IPDFPage, point corefoundation.CGPoint) PDFDestination {
	instance := getPDFDestinationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPage:atPoint:"), page, point)
	return PDFDestinationFromID(rv)
}

// Initializes the destination.
//
// page: The page of the destination.
//
// point: The point of the destination, in page space.
//
// # Return Value
//
// An initialized [PDFDestination] instance, or [NULL] if the object could not
// be initialized.
//
// # Discussion
//
// Specify `point` in page space. Typically, there’s no need to initialize
// destinations. Instead, you get them from [PDFAnnotationLink], [PDFOutline],
// or [PDFView] objects.
//
// Page space is a 72-dpi coordinate system with the origin at the lower-left
// corner of the current page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDestination/init(page:at:)
//
// [PDFAnnotationLink]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationLink
func (p PDFDestination) InitWithPageAtPoint(page IPDFPage, point corefoundation.CGPoint) PDFDestination {
	rv := objc.Send[PDFDestination](p.ID, objc.Sel("initWithPage:atPoint:"), page, point)
	return rv
}

// Returns a comparison result that indicates the location of the destination
// in the document, relative to the current position.
//
// destination: The destination in the document to be located.
//
// # Return Value
//
// A comparison result, indicating the position of the passed-in destination
// relative to the current position.
//
// # Discussion
//
// If `destination` is between the receiver’s position and the end of the
// document, `compare` returns [NSOrderedAscending]; if it is between the
// receiver’s position and the beginning of the document, `compare` returns
// [NSOrderedDescending]. Otherwise, if `destination` matches the receiver’s
// position, `compare` returns [NSOrderedSame].
//
// This method ignores the horizontal component of the destination point (the
// x value). If the destination’s vertical component (or y value) is
// [PDFDestination], `compare` treats the destination as if its y value is the
// top point on the destination page.
//
// An exception is raised if `destination` does not have a page associated
// with it or if its page is associated with a document other than the
// receiver’s document.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDestination/compare(_:)
func (p PDFDestination) Compare(destination IPDFDestination) foundation.NSComparisonResult {
	rv := objc.Send[foundation.NSComparisonResult](p.ID, objc.Sel("compare:"), destination)
	return foundation.NSComparisonResult(rv)
}

// Returns the page that the destination refers to.
//
// # Return Value
//
// The page referred to by the destination.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDestination/page
func (p PDFDestination) Page() IPDFPage {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("page"))
	return PDFPageFromID(objc.ID(rv))
}

// Returns the point, in page space, that the destination refers to.
//
// # Return Value
//
// The point, in page space, referred to by the destination.
//
// # Discussion
//
// Page space is a 72 dpi coordinate system with the origin at the lower-left
// corner of the current page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFDestination/point
func (p PDFDestination) Point() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](p.ID, objc.Sel("point"))
	return corefoundation.CGPoint(rv)
}

// See: https://developer.apple.com/documentation/pdfkit/kpdfdestinationunspecifiedvalue
func (p PDFDestination) KPDFDestinationUnspecifiedValue() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("kPDFDestinationUnspecifiedValue"))
	return rv
}

// See: https://developer.apple.com/documentation/PDFKit/PDFDestination/zoom
func (p PDFDestination) Zoom() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("zoom"))
	return rv
}
func (p PDFDestination) SetZoom(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setZoom:"), value)
}

// An object that represents an action for a PDF element, such as a link
// annotation.
//
// See: https://developer.apple.com/documentation/pdfkit/pdfannotation/action
func (p PDFDestination) Action() IPDFAction {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("action"))
	return PDFActionFromID(objc.ID(rv))
}
func (p PDFDestination) SetAction(value IPDFAction) {
	objc.Send[struct{}](p.ID, objc.Sel("setAction:"), value)
}

// Returns a
//
// See: https://developer.apple.com/documentation/pdfkit/pdfview/currentdestination
func (p PDFDestination) CurrentDestination() IPDFDestination {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("currentDestination"))
	return PDFDestinationFromID(objc.ID(rv))
}
func (p PDFDestination) SetCurrentDestination(value IPDFDestination) {
	objc.Send[struct{}](p.ID, objc.Sel("setCurrentDestination:"), value)
}

// Returns the modification date of the annotation.
//
// See: https://developer.apple.com/documentation/pdfkit/pdfannotation/modificationdate
func (p PDFDestination) ModificationDate() foundation.INSDate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("modificationDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (p PDFDestination) SetModificationDate(value foundation.INSDate) {
	objc.Send[struct{}](p.ID, objc.Sel("setModificationDate:"), value)
}

// Returns the type of the annotation.
//
// See: https://developer.apple.com/documentation/pdfkit/pdfannotation/type
func (p PDFDestination) Type() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("type"))
	return foundation.NSStringFromID(rv).String()
}
func (p PDFDestination) SetType(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setType:"), objc.String(value))
}

// Returns the name of the user who created the annotation.
//
// See: https://developer.apple.com/documentation/pdfkit/pdfannotation/username
func (p PDFDestination) UserName() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("userName"))
	return foundation.NSStringFromID(rv).String()
}
func (p PDFDestination) SetUserName(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setUserName:"), objc.String(value))
}
