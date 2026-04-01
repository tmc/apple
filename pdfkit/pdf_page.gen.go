// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PDFPage] class.
var (
	_PDFPageClass     PDFPageClass
	_PDFPageClassOnce sync.Once
)

func getPDFPageClass() PDFPageClass {
	_PDFPageClassOnce.Do(func() {
		_PDFPageClass = PDFPageClass{class: objc.GetClass("PDFPage")}
	})
	return _PDFPageClass
}

// GetPDFPageClass returns the class object for PDFPage.
func GetPDFPageClass() PDFPageClass {
	return getPDFPageClass()
}

type PDFPageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PDFPageClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PDFPageClass) Alloc() PDFPage {
	rv := objc.Send[PDFPage](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// [PDFPage], a subclass of [NSObject], defines methods used to render PDF
// pages and work with annotations, text, and selections.
//
// # Overview
//
// [PDFPage] objects are flexible and powerful. With them you can render PDF
// content onscreen or to a printer, add annotations, count characters, define
// selections, and get the textual content of a page as an [NSString] object.
//
// Your application instantiates a [PDFPage] object by asking for one from a
// [PDFDocument] object.
//
// For simple display and navigation of PDF documents within your application,
// you don’t need to use [PDFPage]. You need only use [PDFView].
//
// # Initializing a Page
//
//   - [PDFPage.InitWithImage]: Creates a new [PDFPage] object and initializes it with the specified [NSImage] object.
//
// # Getting Information About a Page
//
//   - [PDFPage.Document]: Returns the [PDFDocument] object with which the page is associated.
//   - [PDFPage.Label]: Returns the label for the page.
//   - [PDFPage.BoundsForBox]: Returns the bounds for the specified PDF display box.
//   - [PDFPage.SetBoundsForBox]: Sets the bounds for the specified box.
//   - [PDFPage.Rotation]: Sets the rotation angle for the page in degrees.
//   - [PDFPage.SetRotation]
//
// # Working with Annotations
//
//   - [PDFPage.Annotations]: Returns an array containing the page’s annotations.
//   - [PDFPage.DisplaysAnnotations]: Returns a Boolean value indicating whether annotations are displayed for the page.
//   - [PDFPage.SetDisplaysAnnotations]
//   - [PDFPage.AddAnnotation]: Adds the specified annotation object to the page.
//   - [PDFPage.RemoveAnnotation]: Removes the specified annotation from the page.
//   - [PDFPage.AnnotationAtPoint]: Returns the annotation, if there is one, at the specified point.
//
// # Working with Textual Content
//
//   - [PDFPage.NumberOfCharacters]: Returns the number of characters on the page, including whitespace characters.
//   - [PDFPage.String]: Returns an [NSString] object representing the text on the page.
//   - [PDFPage.AttributedString]: Returns an [NSAttributedString] object representing the text on the page.
//   - [PDFPage.CharacterBoundsAtIndex]: Returns the bounds, in page space, of the character at the specified index.
//   - [PDFPage.CharacterIndexAtPoint]: Returns the character index value for the specified point in page space.
//
// # Working with Selections
//
//   - [PDFPage.SelectionForRect]: Returns the text enclosed within the specified rectangle, expressed in page (user) coordinates.
//   - [PDFPage.SelectionForWordAtPoint]: Returns the whole word that includes the specified point.
//   - [PDFPage.SelectionForLineAtPoint]: Returns the whole line of text that includes the specified point.
//   - [PDFPage.SelectionFromPointToPoint]: Returns the text between the two specified points in page space.
//   - [PDFPage.SelectionForRange]: Returns the text contained within the specified range.
//
// # Initializers
//
//   - [PDFPage.InitWithImageOptions]
//
// # Instance Properties
//
//   - [PDFPage.DataRepresentation]: Returns the PDF data (that is, a PDF document) representing this page. This method does not preserve external page links.
//   - [PDFPage.PageRef]
//
// # Instance Methods
//
//   - [PDFPage.DrawWithBoxToContext]
//   - [PDFPage.ThumbnailOfSizeForBox]
//   - [PDFPage.TransformContextForBox]
//   - [PDFPage.TransformForBox]
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage
type PDFPage struct {
	objectivec.Object
}

// PDFPageFromID constructs a [PDFPage] from an objc.ID.
//
// [PDFPage], a subclass of [NSObject], defines methods used to render PDF
// pages and work with annotations, text, and selections.
func PDFPageFromID(id objc.ID) PDFPage {
	return PDFPage{objectivec.Object{ID: id}}
}

// NOTE: PDFPage adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [PDFPage] class.
//
// # Initializing a Page
//
//   - [IPDFPage.InitWithImage]: Creates a new [PDFPage] object and initializes it with the specified [NSImage] object.
//
// # Getting Information About a Page
//
//   - [IPDFPage.Document]: Returns the [PDFDocument] object with which the page is associated.
//   - [IPDFPage.Label]: Returns the label for the page.
//   - [IPDFPage.BoundsForBox]: Returns the bounds for the specified PDF display box.
//   - [IPDFPage.SetBoundsForBox]: Sets the bounds for the specified box.
//   - [IPDFPage.Rotation]: Sets the rotation angle for the page in degrees.
//   - [IPDFPage.SetRotation]
//
// # Working with Annotations
//
//   - [IPDFPage.Annotations]: Returns an array containing the page’s annotations.
//   - [IPDFPage.DisplaysAnnotations]: Returns a Boolean value indicating whether annotations are displayed for the page.
//   - [IPDFPage.SetDisplaysAnnotations]
//   - [IPDFPage.AddAnnotation]: Adds the specified annotation object to the page.
//   - [IPDFPage.RemoveAnnotation]: Removes the specified annotation from the page.
//   - [IPDFPage.AnnotationAtPoint]: Returns the annotation, if there is one, at the specified point.
//
// # Working with Textual Content
//
//   - [IPDFPage.NumberOfCharacters]: Returns the number of characters on the page, including whitespace characters.
//   - [IPDFPage.String]: Returns an [NSString] object representing the text on the page.
//   - [IPDFPage.AttributedString]: Returns an [NSAttributedString] object representing the text on the page.
//   - [IPDFPage.CharacterBoundsAtIndex]: Returns the bounds, in page space, of the character at the specified index.
//   - [IPDFPage.CharacterIndexAtPoint]: Returns the character index value for the specified point in page space.
//
// # Working with Selections
//
//   - [IPDFPage.SelectionForRect]: Returns the text enclosed within the specified rectangle, expressed in page (user) coordinates.
//   - [IPDFPage.SelectionForWordAtPoint]: Returns the whole word that includes the specified point.
//   - [IPDFPage.SelectionForLineAtPoint]: Returns the whole line of text that includes the specified point.
//   - [IPDFPage.SelectionFromPointToPoint]: Returns the text between the two specified points in page space.
//   - [IPDFPage.SelectionForRange]: Returns the text contained within the specified range.
//
// # Initializers
//
//   - [IPDFPage.InitWithImageOptions]
//
// # Instance Properties
//
//   - [IPDFPage.DataRepresentation]: Returns the PDF data (that is, a PDF document) representing this page. This method does not preserve external page links.
//   - [IPDFPage.PageRef]
//
// # Instance Methods
//
//   - [IPDFPage.DrawWithBoxToContext]
//   - [IPDFPage.ThumbnailOfSizeForBox]
//   - [IPDFPage.TransformContextForBox]
//   - [IPDFPage.TransformForBox]
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage
type IPDFPage interface {
	objectivec.IObject

	// Topic: Initializing a Page

	// Creates a new [PDFPage] object and initializes it with the specified [NSImage] object.
	InitWithImage(image appkit.NSImage) PDFPage

	// Topic: Getting Information About a Page

	// Returns the [PDFDocument] object with which the page is associated.
	Document() IPDFDocument
	// Returns the label for the page.
	Label() string
	// Returns the bounds for the specified PDF display box.
	BoundsForBox(box PDFDisplayBox) corefoundation.CGRect
	// Sets the bounds for the specified box.
	SetBoundsForBox(bounds corefoundation.CGRect, box PDFDisplayBox)
	// Sets the rotation angle for the page in degrees.
	Rotation() int
	SetRotation(value int)

	// Topic: Working with Annotations

	// Returns an array containing the page’s annotations.
	Annotations() []PDFAnnotation
	// Returns a Boolean value indicating whether annotations are displayed for the page.
	DisplaysAnnotations() bool
	SetDisplaysAnnotations(value bool)
	// Adds the specified annotation object to the page.
	AddAnnotation(annotation IPDFAnnotation)
	// Removes the specified annotation from the page.
	RemoveAnnotation(annotation IPDFAnnotation)
	// Returns the annotation, if there is one, at the specified point.
	AnnotationAtPoint(point corefoundation.CGPoint) IPDFAnnotation

	// Topic: Working with Textual Content

	// Returns the number of characters on the page, including whitespace characters.
	NumberOfCharacters() uint
	// Returns an [NSString] object representing the text on the page.
	String() string
	// Returns an [NSAttributedString] object representing the text on the page.
	AttributedString() foundation.NSAttributedString
	// Returns the bounds, in page space, of the character at the specified index.
	CharacterBoundsAtIndex(index int) corefoundation.CGRect
	// Returns the character index value for the specified point in page space.
	CharacterIndexAtPoint(point corefoundation.CGPoint) int

	// Topic: Working with Selections

	// Returns the text enclosed within the specified rectangle, expressed in page (user) coordinates.
	SelectionForRect(rect corefoundation.CGRect) IPDFSelection
	// Returns the whole word that includes the specified point.
	SelectionForWordAtPoint(point corefoundation.CGPoint) IPDFSelection
	// Returns the whole line of text that includes the specified point.
	SelectionForLineAtPoint(point corefoundation.CGPoint) IPDFSelection
	// Returns the text between the two specified points in page space.
	SelectionFromPointToPoint(startPoint corefoundation.CGPoint, endPoint corefoundation.CGPoint) IPDFSelection
	// Returns the text contained within the specified range.
	SelectionForRange(range_ foundation.NSRange) IPDFSelection

	// Topic: Initializers

	InitWithImageOptions(image appkit.NSImage, options foundation.INSDictionary) PDFPage

	// Topic: Instance Properties

	// Returns the PDF data (that is, a PDF document) representing this page. This method does not preserve external page links.
	DataRepresentation() foundation.INSData
	PageRef() coregraphics.CGPDFPageRef

	// Topic: Instance Methods

	DrawWithBoxToContext(box PDFDisplayBox, context coregraphics.CGContextRef)
	ThumbnailOfSizeForBox(size corefoundation.CGSize, box PDFDisplayBox) appkit.NSImage
	TransformContextForBox(context coregraphics.CGContextRef, box PDFDisplayBox)
	TransformForBox(box PDFDisplayBox) corefoundation.CGAffineTransform
}

// Init initializes the instance.
func (p PDFPage) Init() PDFPage {
	rv := objc.Send[PDFPage](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PDFPage) Autorelease() PDFPage {
	rv := objc.Send[PDFPage](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFPage creates a new PDFPage instance.
func NewPDFPage() PDFPage {
	class := getPDFPageClass()
	rv := objc.Send[PDFPage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new [PDFPage] object and initializes it with the specified
// [NSImage] object.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage/init(image:)
func NewPDFPageWithImage(image appkit.NSImage) PDFPage {
	instance := getPDFPageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImage:"), image)
	return PDFPageFromID(rv)
}

// See: https://developer.apple.com/documentation/PDFKit/PDFPage/init(image:options:)
func NewPDFPageWithImageOptions(image appkit.NSImage, options foundation.INSDictionary) PDFPage {
	instance := getPDFPageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImage:options:"), image, options)
	return PDFPageFromID(rv)
}

// Creates a new [PDFPage] object and initializes it with the specified
// [NSImage] object.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage/init(image:)
func (p PDFPage) InitWithImage(image appkit.NSImage) PDFPage {
	rv := objc.Send[PDFPage](p.ID, objc.Sel("initWithImage:"), image)
	return rv
}

// Returns the bounds for the specified PDF display box.
//
// # Discussion
//
// The [PDFDisplayBox] enumeration defines the various box types.
//
// Note that only the media box is required for a PDF. If you request the
// bounds for the crop box, but the PDF does not include a crop box, the
// bounds for the media box are returned instead. If you request the bounds
// for other box types, and the PDF does not includes these types, the bounds
// for the crop box are returned instead.
//
// The coordinates for the box are in page space, so you might need to
// transform the points if the page has a rotation on it. Also, note that the
// bounds `boundsForBox` returns are intersected with the page’s media box.
//
// `boundsForBox` throws a range exception if `box` is not in range.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage/bounds(for:)
//
// [PDFDisplayBox]: https://developer.apple.com/documentation/PDFKit/PDFDisplayBox
func (p PDFPage) BoundsForBox(box PDFDisplayBox) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p.ID, objc.Sel("boundsForBox:"), box)
	return corefoundation.CGRect(rv)
}

// Sets the bounds for the specified box.
//
// # Discussion
//
// If the box does not exist, this method creates it for you.
//
// To remove a box, pass [NSZeroRect] for the bounds (note that you cannot
// remove the media box). If the box bounds are not in range, this method
// throws a range exception.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage/setBounds(_:for:)
func (p PDFPage) SetBoundsForBox(bounds corefoundation.CGRect, box PDFDisplayBox) {
	objc.Send[objc.ID](p.ID, objc.Sel("setBounds:forBox:"), bounds, box)
}

// Adds the specified annotation object to the page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage/addAnnotation(_:)
func (p PDFPage) AddAnnotation(annotation IPDFAnnotation) {
	objc.Send[objc.ID](p.ID, objc.Sel("addAnnotation:"), annotation)
}

// Removes the specified annotation from the page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage/removeAnnotation(_:)
func (p PDFPage) RemoveAnnotation(annotation IPDFAnnotation) {
	objc.Send[objc.ID](p.ID, objc.Sel("removeAnnotation:"), annotation)
}

// Returns the annotation, if there is one, at the specified point.
//
// # Discussion
//
// Use this method for hit-testing based on the current cursor position. If
// more than one annotation shares the specified point, the frontmost (or
// topmost) one is returned (the annotations are searched in reverse order of
// their appearance in the PDF data file). Returns [NULL] if there is no
// annotation at `point`.
//
// Specify the point in page space. Page space is a 72 dpi coordinate system
// with the origin at the lower-left corner of the current page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage/annotation(at:)
func (p PDFPage) AnnotationAtPoint(point corefoundation.CGPoint) IPDFAnnotation {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("annotationAtPoint:"), point)
	return PDFAnnotationFromID(rv)
}

// Returns the bounds, in page space, of the character at the specified index.
//
// # Discussion
//
// In the unlikely event that there is more than one character at the
// specified index point, only the bounds of the first character is returned.
//
// Page space is a 72 dpi coordinate system with the origin at the lower-left
// corner of the current page. Note that the bounds returned are not
// guaranteed to have integer coordinates.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage/characterBounds(at:)
func (p PDFPage) CharacterBoundsAtIndex(index int) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p.ID, objc.Sel("characterBoundsAtIndex:"), index)
	return corefoundation.CGRect(rv)
}

// Returns the character index value for the specified point in page space.
//
// # Discussion
//
// If there is no character at the specified point, the method returns `-1`.
//
// Page space is a 72 dpi coordinate system with the origin at the lower-left
// corner of the current page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage/characterIndex(at:)
func (p PDFPage) CharacterIndexAtPoint(point corefoundation.CGPoint) int {
	rv := objc.Send[int](p.ID, objc.Sel("characterIndexAtPoint:"), point)
	return rv
}

// Returns the text enclosed within the specified rectangle, expressed in page
// (user) coordinates.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage/selection(for:)-2ckpi
func (p PDFPage) SelectionForRect(rect corefoundation.CGRect) IPDFSelection {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("selectionForRect:"), rect)
	return PDFSelectionFromID(rv)
}

// Returns the whole word that includes the specified point.
//
// # Discussion
//
// Returns [NULL] if no word contains `point`.
//
// Use this method to respond to a double-click.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage/selectionForWord(at:)
func (p PDFPage) SelectionForWordAtPoint(point corefoundation.CGPoint) IPDFSelection {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("selectionForWordAtPoint:"), point)
	return PDFSelectionFromID(rv)
}

// Returns the whole line of text that includes the specified point.
//
// # Discussion
//
// Returns [NULL] if no line of text contains `point`.
//
// Use this method to respond to a triple-click.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage/selectionForLine(at:)
func (p PDFPage) SelectionForLineAtPoint(point corefoundation.CGPoint) IPDFSelection {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("selectionForLineAtPoint:"), point)
	return PDFSelectionFromID(rv)
}

// Returns the text between the two specified points in page space.
//
// # Discussion
//
// Either point may be the one closer to the start of the page. In determining
// the selection, the points are sorted first top to bottom and then left to
// right.
//
// Page space is a 72 dpi coordinate system with the origin at the lower-left
// corner of the current page.
//
// To visualize the selection, picture the rectangle defined by `startPoint`
// and `endPoint`. The selection begins at the first character fully within
// the defined rectangle and closest to its upper-left corner. The selection
// ends at the last character fully within the defined rectangle and closest
// to its lower-right corner.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage/selection(from:to:)
func (p PDFPage) SelectionFromPointToPoint(startPoint corefoundation.CGPoint, endPoint corefoundation.CGPoint) IPDFSelection {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("selectionFromPoint:toPoint:"), startPoint, endPoint)
	return PDFSelectionFromID(rv)
}

// Returns the text contained within the specified range.
//
// # Discussion
//
// This method raises an exception if the range length is `0` or if either end
// of the range is outside the range of characters on the page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage/selection(for:)-20y9d
func (p PDFPage) SelectionForRange(range_ foundation.NSRange) IPDFSelection {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("selectionForRange:"), range_)
	return PDFSelectionFromID(rv)
}

// See: https://developer.apple.com/documentation/PDFKit/PDFPage/init(image:options:)
func (p PDFPage) InitWithImageOptions(image appkit.NSImage, options foundation.INSDictionary) PDFPage {
	rv := objc.Send[PDFPage](p.ID, objc.Sel("initWithImage:options:"), image, options)
	return rv
}

// See: https://developer.apple.com/documentation/PDFKit/PDFPage/draw(with:to:)
func (p PDFPage) DrawWithBoxToContext(box PDFDisplayBox, context coregraphics.CGContextRef) {
	objc.Send[objc.ID](p.ID, objc.Sel("drawWithBox:toContext:"), box, context)
}

// See: https://developer.apple.com/documentation/PDFKit/PDFPage/thumbnail(of:for:)
func (p PDFPage) ThumbnailOfSizeForBox(size corefoundation.CGSize, box PDFDisplayBox) appkit.NSImage {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("thumbnailOfSize:forBox:"), size, box)
	return appkit.NSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/PDFKit/PDFPage/transform(_:for:)
func (p PDFPage) TransformContextForBox(context coregraphics.CGContextRef, box PDFDisplayBox) {
	objc.Send[objc.ID](p.ID, objc.Sel("transformContext:forBox:"), context, box)
}

// See: https://developer.apple.com/documentation/PDFKit/PDFPage/transform(for:)
func (p PDFPage) TransformForBox(box PDFDisplayBox) corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](p.ID, objc.Sel("transformForBox:"), box)
	return corefoundation.CGAffineTransform(rv)
}

// Returns the [PDFDocument] object with which the page is associated.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage/document
func (p PDFPage) Document() IPDFDocument {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("document"))
	return PDFDocumentFromID(objc.ID(rv))
}

// Returns the label for the page.
//
// # Discussion
//
// Typically, the label is “1” for the first page, “2” for the second
// page, and so on, but nonnumerical labels are also possible (such as
// “xxi”, “4-1” and so on).
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage/label
func (p PDFPage) Label() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the rotation angle for the page in degrees.
//
// # Discussion
//
// The rotation must be a positive or negative multiple of 90 (negative angles
// are converted to their positive equivalents; for example, -90 is changed to
// 270); otherwise this method throws an exception.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage/rotation
func (p PDFPage) Rotation() int {
	rv := objc.Send[int](p.ID, objc.Sel("rotation"))
	return rv
}
func (p PDFPage) SetRotation(value int) {
	objc.Send[struct{}](p.ID, objc.Sel("setRotation:"), value)
}

// Returns an array containing the page’s annotations.
//
// # Discussion
//
// The elements of the array will most likely be typed to subclasses of the
// [PDFAnnotation] class.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage/annotations
func (p PDFPage) Annotations() []PDFAnnotation {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("annotations"))
	return objc.ConvertSlice(rv, func(id objc.ID) PDFAnnotation {
		return PDFAnnotationFromID(id)
	})
}

// Returns a Boolean value indicating whether annotations are displayed for
// the page.
//
// # Discussion
//
// If true, the page will draw annotations when a drawing method is called.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage/displaysAnnotations
func (p PDFPage) DisplaysAnnotations() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("displaysAnnotations"))
	return rv
}
func (p PDFPage) SetDisplaysAnnotations(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setDisplaysAnnotations:"), value)
}

// Returns the number of characters on the page, including whitespace
// characters.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage/numberOfCharacters
func (p PDFPage) NumberOfCharacters() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("numberOfCharacters"))
	return rv
}

// Returns an [NSString] object representing the text on the page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage/string
func (p PDFPage) String() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("string"))
	return foundation.NSStringFromID(rv).String()
}

// Returns an [NSAttributedString] object representing the text on the page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage/attributedString
func (p PDFPage) AttributedString() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("attributedString"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}

// Returns the PDF data (that is, a PDF document) representing this page. This
// method does not preserve external page links.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPage/dataRepresentation
func (p PDFPage) DataRepresentation() foundation.INSData {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("dataRepresentation"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/PDFKit/PDFPage/pageRef
func (p PDFPage) PageRef() coregraphics.CGPDFPageRef {
	rv := objc.Send[coregraphics.CGPDFPageRef](p.ID, objc.Sel("pageRef"))
	return coregraphics.CGPDFPageRef(rv)
}
