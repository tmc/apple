// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PDFSelection] class.
var (
	_PDFSelectionClass     PDFSelectionClass
	_PDFSelectionClassOnce sync.Once
)

func getPDFSelectionClass() PDFSelectionClass {
	_PDFSelectionClassOnce.Do(func() {
		_PDFSelectionClass = PDFSelectionClass{class: objc.GetClass("PDFSelection")}
	})
	return _PDFSelectionClass
}

// GetPDFSelectionClass returns the class object for PDFSelection.
func GetPDFSelectionClass() PDFSelectionClass {
	return getPDFSelectionClass()
}

type PDFSelectionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PDFSelectionClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PDFSelectionClass) Alloc() PDFSelection {
	rv := objc.Send[PDFSelection](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// A [PDFSelection] object identifies a contiguous or noncontiguous selection
// of text in a PDF document.
//
// # Initializing a Selection
//
//   - [PDFSelection.InitWithDocument]: Returns an empty [PDFSelection] object.
//
// # Getting Information About a Selection
//
//   - [PDFSelection.Pages]: Returns the array of pages contained in the selection.
//   - [PDFSelection.String]: Returns an [NSString] object representing the text contained in the selection (may contain linefeed characters).
//   - [PDFSelection.AttributedString]: Returns an [NSAttributedString] object representing the text contained in the selection (may contain linefeed characters).
//   - [PDFSelection.BoundsForPage]: Returns the bounds of the selection on the specified page.
//   - [PDFSelection.SelectionsByLine]: Returns an array of selections, one for each line of text covered by the receiver.
//   - [PDFSelection.Color]: Sets the color used for the drawing of a selection in both active and inactive states.
//   - [PDFSelection.SetColor]
//
// # Modifying a Selection
//
//   - [PDFSelection.AddSelection]: Adds the specified selection to the receiving selection.
//   - [PDFSelection.AddSelections]: Adds the specified array of selections to the receiving selection.
//   - [PDFSelection.ExtendSelectionAtEnd]: Extends the selection from its end toward the end of the document.
//   - [PDFSelection.ExtendSelectionAtStart]: Extends the selection from its start toward the beginning of the document.
//
// # Managing Selection Drawing
//
//   - [PDFSelection.DrawForPageActive]: Calls [draw(for:with:active:)](<doc://com.apple.pdfkit/documentation/PDFKit/PDFSelection/draw(for:with:active:)>) with a default value for box parameter.
//   - [PDFSelection.DrawForPageWithBoxActive]: Draws the selection relative to the origin of the specified box in page space.
//
// # Instance Methods
//
//   - [PDFSelection.ExtendSelectionForLineBoundaries]
//   - [PDFSelection.NumberOfTextRangesOnPage]
//   - [PDFSelection.RangeAtIndexOnPage]
//
// See: https://developer.apple.com/documentation/PDFKit/PDFSelection
type PDFSelection struct {
	objectivec.Object
}

// PDFSelectionFromID constructs a [PDFSelection] from an objc.ID.
//
// A [PDFSelection] object identifies a contiguous or noncontiguous selection
// of text in a PDF document.
func PDFSelectionFromID(id objc.ID) PDFSelection {
	return PDFSelection{objectivec.Object{ID: id}}
}

// NOTE: PDFSelection adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [PDFSelection] class.
//
// # Initializing a Selection
//
//   - [IPDFSelection.InitWithDocument]: Returns an empty [PDFSelection] object.
//
// # Getting Information About a Selection
//
//   - [IPDFSelection.Pages]: Returns the array of pages contained in the selection.
//   - [IPDFSelection.String]: Returns an [NSString] object representing the text contained in the selection (may contain linefeed characters).
//   - [IPDFSelection.AttributedString]: Returns an [NSAttributedString] object representing the text contained in the selection (may contain linefeed characters).
//   - [IPDFSelection.BoundsForPage]: Returns the bounds of the selection on the specified page.
//   - [IPDFSelection.SelectionsByLine]: Returns an array of selections, one for each line of text covered by the receiver.
//   - [IPDFSelection.Color]: Sets the color used for the drawing of a selection in both active and inactive states.
//   - [IPDFSelection.SetColor]
//
// # Modifying a Selection
//
//   - [IPDFSelection.AddSelection]: Adds the specified selection to the receiving selection.
//   - [IPDFSelection.AddSelections]: Adds the specified array of selections to the receiving selection.
//   - [IPDFSelection.ExtendSelectionAtEnd]: Extends the selection from its end toward the end of the document.
//   - [IPDFSelection.ExtendSelectionAtStart]: Extends the selection from its start toward the beginning of the document.
//
// # Managing Selection Drawing
//
//   - [IPDFSelection.DrawForPageActive]: Calls [draw(for:with:active:)](<doc://com.apple.pdfkit/documentation/PDFKit/PDFSelection/draw(for:with:active:)>) with a default value for box parameter.
//   - [IPDFSelection.DrawForPageWithBoxActive]: Draws the selection relative to the origin of the specified box in page space.
//
// # Instance Methods
//
//   - [IPDFSelection.ExtendSelectionForLineBoundaries]
//   - [IPDFSelection.NumberOfTextRangesOnPage]
//   - [IPDFSelection.RangeAtIndexOnPage]
//
// See: https://developer.apple.com/documentation/PDFKit/PDFSelection
type IPDFSelection interface {
	objectivec.IObject

	// Topic: Initializing a Selection

	// Returns an empty [PDFSelection] object.
	InitWithDocument(document IPDFDocument) PDFSelection

	// Topic: Getting Information About a Selection

	// Returns the array of pages contained in the selection.
	Pages() []PDFPage
	// Returns an [NSString] object representing the text contained in the selection (may contain linefeed characters).
	String() string
	// Returns an [NSAttributedString] object representing the text contained in the selection (may contain linefeed characters).
	AttributedString() foundation.NSAttributedString
	// Returns the bounds of the selection on the specified page.
	BoundsForPage(page IPDFPage) corefoundation.CGRect
	// Returns an array of selections, one for each line of text covered by the receiver.
	SelectionsByLine() []PDFSelection
	// Sets the color used for the drawing of a selection in both active and inactive states.
	Color() appkit.NSColor
	SetColor(value appkit.NSColor)

	// Topic: Modifying a Selection

	// Adds the specified selection to the receiving selection.
	AddSelection(selection IPDFSelection)
	// Adds the specified array of selections to the receiving selection.
	AddSelections(selections []PDFSelection)
	// Extends the selection from its end toward the end of the document.
	ExtendSelectionAtEnd(succeed int)
	// Extends the selection from its start toward the beginning of the document.
	ExtendSelectionAtStart(precede int)

	// Topic: Managing Selection Drawing

	// Calls [draw(for:with:active:)](<doc://com.apple.pdfkit/documentation/PDFKit/PDFSelection/draw(for:with:active:)>) with a default value for box parameter.
	DrawForPageActive(page IPDFPage, active bool)
	// Draws the selection relative to the origin of the specified box in page space.
	DrawForPageWithBoxActive(page IPDFPage, box PDFDisplayBox, active bool)

	// Topic: Instance Methods

	ExtendSelectionForLineBoundaries()
	NumberOfTextRangesOnPage(page IPDFPage) uint
	RangeAtIndexOnPage(index uint, page IPDFPage) foundation.NSRange
}

// Init initializes the instance.
func (p PDFSelection) Init() PDFSelection {
	rv := objc.Send[PDFSelection](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PDFSelection) Autorelease() PDFSelection {
	rv := objc.Send[PDFSelection](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFSelection creates a new PDFSelection instance.
func NewPDFSelection() PDFSelection {
	class := getPDFSelectionClass()
	rv := objc.Send[PDFSelection](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an empty [PDFSelection] object.
//
// # Discussion
//
// Typically, you don’t need to create a [PDFSelection] object, but you can
// use an empty [PDFSelection] object as a container into which you can place
// selections, using [PDFSelection] and addSelections.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFSelection/init(document:)
func NewPDFSelectionWithDocument(document IPDFDocument) PDFSelection {
	instance := getPDFSelectionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDocument:"), document)
	return PDFSelectionFromID(rv)
}

// Returns an empty [PDFSelection] object.
//
// # Discussion
//
// Typically, you don’t need to create a [PDFSelection] object, but you can
// use an empty [PDFSelection] object as a container into which you can place
// selections, using [PDFSelection] and addSelections.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFSelection/init(document:)
func (p PDFSelection) InitWithDocument(document IPDFDocument) PDFSelection {
	rv := objc.Send[PDFSelection](p.ID, objc.Sel("initWithDocument:"), document)
	return rv
}

// Returns the bounds of the selection on the specified page.
//
// # Discussion
//
// The selection rectangle is given in page space.
//
// Page space is a 72 dpi coordinate system with the origin at the lower-left
// corner of the current page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFSelection/bounds(for:)
func (p PDFSelection) BoundsForPage(page IPDFPage) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p.ID, objc.Sel("boundsForPage:"), page)
	return corefoundation.CGRect(rv)
}

// Returns an array of selections, one for each line of text covered by the
// receiver.
//
// # Discussion
//
// If you call this method on a [PDFSelection] object that represents a
// paragraph, for example, `selectionsByLine` returns an array that contains
// one [PDFSelection] object for each line of text in the paragraph.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFSelection/selectionsByLine()
func (p PDFSelection) SelectionsByLine() []PDFSelection {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("selectionsByLine"))
	return objc.ConvertSlice(rv, func(id objc.ID) PDFSelection {
		return PDFSelectionFromID(id)
	})
}

// Adds the specified selection to the receiving selection.
//
// # Discussion
//
// Selections do not have to be contiguous. If the selection to be added
// overlaps with the receiving selection, the overlap is removed in a process
// called normalization.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFSelection/add(_:)-8c2r
func (p PDFSelection) AddSelection(selection IPDFSelection) {
	objc.Send[objc.ID](p.ID, objc.Sel("addSelection:"), selection)
}

// Adds the specified array of selections to the receiving selection.
//
// # Discussion
//
// This method provides better performance than multiple calls to
// `addSelection` if you need to add several selections to an existing
// selection. This is because the normalization of the selection (the removal
// of any overlaps between selections) occurs only once, after all selections
// have been added.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFSelection/add(_:)-3fyld
func (p PDFSelection) AddSelections(selections []PDFSelection) {
	objc.Send[objc.ID](p.ID, objc.Sel("addSelections:"), objectivec.IObjectSliceToNSArray(selections))
}

// Extends the selection from its end toward the end of the document.
//
// # Discussion
//
// The selection may be extended by any amount, up to and including the end of
// the document.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFSelection/extend(atEnd:)
func (p PDFSelection) ExtendSelectionAtEnd(succeed int) {
	objc.Send[objc.ID](p.ID, objc.Sel("extendSelectionAtEnd:"), succeed)
}

// Extends the selection from its start toward the beginning of the document.
//
// # Discussion
//
// The selection may be extended by any amount, up to and including the
// beginning of the document.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFSelection/extend(atStart:)
func (p PDFSelection) ExtendSelectionAtStart(precede int) {
	objc.Send[objc.ID](p.ID, objc.Sel("extendSelectionAtStart:"), precede)
}

// Calls [DrawForPageWithBoxActive] with a default value for box parameter.
//
// # Discussion
//
// The default value is `kPDFDisplayBoxCropBox`. If active is true, drawing
// uses `selectedTextBackgroundColor`. If false, it uses
// `secondarySelectedControlColor`.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFSelection/draw(for:active:)
func (p PDFSelection) DrawForPageActive(page IPDFPage, active bool) {
	objc.Send[objc.ID](p.ID, objc.Sel("drawForPage:active:"), page, active)
}

// Draws the selection relative to the origin of the specified box in page
// space.
//
// # Discussion
//
// The selection is drawn using the current highlight color. If active is
// true, drawing uses `selectedTextBackgroundColor`. If false, it uses
// `secondarySelectedControlColor`. Refer to the [PDFPage] class for the list
// of available box types.
//
// Page space is a 72 dpi coordinate system with the origin at the lower-left
// corner of the current page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFSelection/draw(for:with:active:)
func (p PDFSelection) DrawForPageWithBoxActive(page IPDFPage, box PDFDisplayBox, active bool) {
	objc.Send[objc.ID](p.ID, objc.Sel("drawForPage:withBox:active:"), page, box, active)
}

// See: https://developer.apple.com/documentation/PDFKit/PDFSelection/extendForLineBoundaries()
func (p PDFSelection) ExtendSelectionForLineBoundaries() {
	objc.Send[objc.ID](p.ID, objc.Sel("extendSelectionForLineBoundaries"))
}

// See: https://developer.apple.com/documentation/PDFKit/PDFSelection/numberOfTextRanges(on:)
func (p PDFSelection) NumberOfTextRangesOnPage(page IPDFPage) uint {
	rv := objc.Send[uint](p.ID, objc.Sel("numberOfTextRangesOnPage:"), page)
	return rv
}

// See: https://developer.apple.com/documentation/PDFKit/PDFSelection/range(at:on:)
func (p PDFSelection) RangeAtIndexOnPage(index uint, page IPDFPage) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](p.ID, objc.Sel("rangeAtIndex:onPage:"), index, page)
	return foundation.NSRange(rv)
}

// Returns the array of pages contained in the selection.
//
// # Discussion
//
// Pages are sorted by index number.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFSelection/pages
func (p PDFSelection) Pages() []PDFPage {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("pages"))
	return objc.ConvertSlice(rv, func(id objc.ID) PDFPage {
		return PDFPageFromID(id)
	})
}

// Returns an [NSString] object representing the text contained in the
// selection (may contain linefeed characters).
//
// See: https://developer.apple.com/documentation/PDFKit/PDFSelection/string
func (p PDFSelection) String() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("string"))
	return foundation.NSStringFromID(rv).String()
}

// Returns an [NSAttributedString] object representing the text contained in
// the selection (may contain linefeed characters).
//
// See: https://developer.apple.com/documentation/PDFKit/PDFSelection/attributedString
func (p PDFSelection) AttributedString() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("attributedString"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}

// Sets the color used for the drawing of a selection in both active and
// inactive states.
//
// # Discussion
//
// When no color has been specified for the [PDFSelection] objects in a
// document, the selections are drawn using `[NSColor
// selectedTextBackgroundColor]` for the active state and `[NSColor
// secondarySelectedControlColor]` for the inactive state. Use the `setColor`
// method to supply a color you want to be used for the drawing of both active
// and inactive selections.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFSelection/color
func (p PDFSelection) Color() appkit.NSColor {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("color"))
	return appkit.NSColorFromID(objc.ID(rv))
}
func (p PDFSelection) SetColor(value appkit.NSColor) {
	objc.Send[struct{}](p.ID, objc.Sel("setColor:"), value)
}
