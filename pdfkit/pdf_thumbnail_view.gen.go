// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [PDFThumbnailView] class.
var (
	_PDFThumbnailViewClass     PDFThumbnailViewClass
	_PDFThumbnailViewClassOnce sync.Once
)

func getPDFThumbnailViewClass() PDFThumbnailViewClass {
	_PDFThumbnailViewClassOnce.Do(func() {
		_PDFThumbnailViewClass = PDFThumbnailViewClass{class: objc.GetClass("PDFThumbnailView")}
	})
	return _PDFThumbnailViewClass
}

// GetPDFThumbnailViewClass returns the class object for PDFThumbnailView.
func GetPDFThumbnailViewClass() PDFThumbnailViewClass {
	return getPDFThumbnailViewClass()
}

type PDFThumbnailViewClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PDFThumbnailViewClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PDFThumbnailViewClass) Alloc() PDFThumbnailView {
	rv := objc.Send[PDFThumbnailView](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// An object that contains a set of thumbnails, each of which represents a
// page in a PDF document.
//
// # Accessing the Associated PDF View
//
//   - [PDFThumbnailView.PDFView]: Returns the [PDFView] object associated with the thumbnail view.
//   - [PDFThumbnailView.SetPDFView]
//
// # Managing the Size of a Thumbnail View
//
//   - [PDFThumbnailView.ThumbnailSize]: Returns the maximum width and height of the thumbnails in the thumbnail view.
//   - [PDFThumbnailView.SetThumbnailSize]
//
// # Working with Thumbnail View Display Characteristics
//
//   - [PDFThumbnailView.MaximumNumberOfColumns]: Returns the maximum number of columns of thumbnails the thumbnail view can display.
//   - [PDFThumbnailView.SetMaximumNumberOfColumns]
//   - [PDFThumbnailView.LabelFont]: Returns the font used to label the thumbnails.
//   - [PDFThumbnailView.SetLabelFont]
//   - [PDFThumbnailView.BackgroundColor]: Returns the color used in the background of the thumbnail view.
//   - [PDFThumbnailView.SetBackgroundColor]
//
// # Managing the Behavior of a Thumbnail View
//
//   - [PDFThumbnailView.AllowsDragging]: Returns a Boolean value indicating whether users can drag thumbnails (that is, re-order pages in the document) within the thumbnail view.
//   - [PDFThumbnailView.SetAllowsDragging]
//   - [PDFThumbnailView.AllowsMultipleSelection]: Returns a Boolean value indicating whether users can select multiple thumbnails in the thumbnail view at one time.
//   - [PDFThumbnailView.SetAllowsMultipleSelection]
//   - [PDFThumbnailView.SelectedPages]: Returns an array of PDF pages that correspond to the selected thumbnails in the thumbnail view.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView
type PDFThumbnailView struct {
	appkit.NSView
}

// PDFThumbnailViewFromID constructs a [PDFThumbnailView] from an objc.ID.
//
// An object that contains a set of thumbnails, each of which represents a
// page in a PDF document.
func PDFThumbnailViewFromID(id objc.ID) PDFThumbnailView {
	return PDFThumbnailView{NSView: appkit.NSViewFromID(id)}
}

// NOTE: PDFThumbnailView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [PDFThumbnailView] class.
//
// # Accessing the Associated PDF View
//
//   - [IPDFThumbnailView.PDFView]: Returns the [PDFView] object associated with the thumbnail view.
//   - [IPDFThumbnailView.SetPDFView]
//
// # Managing the Size of a Thumbnail View
//
//   - [IPDFThumbnailView.ThumbnailSize]: Returns the maximum width and height of the thumbnails in the thumbnail view.
//   - [IPDFThumbnailView.SetThumbnailSize]
//
// # Working with Thumbnail View Display Characteristics
//
//   - [IPDFThumbnailView.MaximumNumberOfColumns]: Returns the maximum number of columns of thumbnails the thumbnail view can display.
//   - [IPDFThumbnailView.SetMaximumNumberOfColumns]
//   - [IPDFThumbnailView.LabelFont]: Returns the font used to label the thumbnails.
//   - [IPDFThumbnailView.SetLabelFont]
//   - [IPDFThumbnailView.BackgroundColor]: Returns the color used in the background of the thumbnail view.
//   - [IPDFThumbnailView.SetBackgroundColor]
//
// # Managing the Behavior of a Thumbnail View
//
//   - [IPDFThumbnailView.AllowsDragging]: Returns a Boolean value indicating whether users can drag thumbnails (that is, re-order pages in the document) within the thumbnail view.
//   - [IPDFThumbnailView.SetAllowsDragging]
//   - [IPDFThumbnailView.AllowsMultipleSelection]: Returns a Boolean value indicating whether users can select multiple thumbnails in the thumbnail view at one time.
//   - [IPDFThumbnailView.SetAllowsMultipleSelection]
//   - [IPDFThumbnailView.SelectedPages]: Returns an array of PDF pages that correspond to the selected thumbnails in the thumbnail view.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView
type IPDFThumbnailView interface {
	appkit.INSView

	// Topic: Accessing the Associated PDF View

	// Returns the [PDFView] object associated with the thumbnail view.
	PDFView() IPDFView
	SetPDFView(value IPDFView)

	// Topic: Managing the Size of a Thumbnail View

	// Returns the maximum width and height of the thumbnails in the thumbnail view.
	ThumbnailSize() corefoundation.CGSize
	SetThumbnailSize(value corefoundation.CGSize)

	// Topic: Working with Thumbnail View Display Characteristics

	// Returns the maximum number of columns of thumbnails the thumbnail view can display.
	MaximumNumberOfColumns() uint
	SetMaximumNumberOfColumns(value uint)
	// Returns the font used to label the thumbnails.
	LabelFont() appkit.NSFont
	SetLabelFont(value appkit.NSFont)
	// Returns the color used in the background of the thumbnail view.
	BackgroundColor() objc.ID
	SetBackgroundColor(value objc.ID)

	// Topic: Managing the Behavior of a Thumbnail View

	// Returns a Boolean value indicating whether users can drag thumbnails (that is, re-order pages in the document) within the thumbnail view.
	AllowsDragging() bool
	SetAllowsDragging(value bool)
	// Returns a Boolean value indicating whether users can select multiple thumbnails in the thumbnail view at one time.
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	// Returns an array of PDF pages that correspond to the selected thumbnails in the thumbnail view.
	SelectedPages() []PDFPage
}

// Init initializes the instance.
func (p PDFThumbnailView) Init() PDFThumbnailView {
	rv := objc.Send[PDFThumbnailView](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PDFThumbnailView) Autorelease() PDFThumbnailView {
	rv := objc.Send[PDFThumbnailView](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFThumbnailView creates a new PDFThumbnailView instance.
func NewPDFThumbnailView() PDFThumbnailView {
	class := getPDFThumbnailViewClass()
	rv := objc.Send[PDFThumbnailView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the [PDFView] object associated with the thumbnail view.
//
// # Return Value
//
// The PDF view object associated with the thumbnail view.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/pdfView
func (p PDFThumbnailView) PDFView() IPDFView {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("PDFView"))
	return PDFViewFromID(objc.ID(rv))
}
func (p PDFThumbnailView) SetPDFView(value IPDFView) {
	objc.Send[struct{}](p.ID, objc.Sel("setPDFView:"), value)
}

// Returns the maximum width and height of the thumbnails in the thumbnail
// view.
//
// # Return Value
//
// The maximum width and height of the thumbnails in the thumbnail view.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/thumbnailSize
func (p PDFThumbnailView) ThumbnailSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](p.ID, objc.Sel("thumbnailSize"))
	return corefoundation.CGSize(rv)
}
func (p PDFThumbnailView) SetThumbnailSize(value corefoundation.CGSize) {
	objc.Send[struct{}](p.ID, objc.Sel("setThumbnailSize:"), value)
}

// Returns the maximum number of columns of thumbnails the thumbnail view can
// display.
//
// # Return Value
//
// The maximum number of columns of thumbnails the thumbnail view can display.
// If `0`, the thumbnail displays as many columns of thumbnails as fit in its
// size.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/maximumNumberOfColumns
func (p PDFThumbnailView) MaximumNumberOfColumns() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("maximumNumberOfColumns"))
	return rv
}
func (p PDFThumbnailView) SetMaximumNumberOfColumns(value uint) {
	objc.Send[struct{}](p.ID, objc.Sel("setMaximumNumberOfColumns:"), value)
}

// Returns the font used to label the thumbnails.
//
// # Return Value
//
// The font used in the thumbnail labels.
//
// # Discussion
//
// Typically, the label of a thumbnail is the page number of the page it
// represents.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/labelFont
func (p PDFThumbnailView) LabelFont() appkit.NSFont {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("labelFont"))
	return appkit.NSFontFromID(objc.ID(rv))
}
func (p PDFThumbnailView) SetLabelFont(value appkit.NSFont) {
	objc.Send[struct{}](p.ID, objc.Sel("setLabelFont:"), value)
}

// Returns the color used in the background of the thumbnail view.
//
// # Return Value
//
// The color of the background in the thumbnail view.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/backgroundColor
func (p PDFThumbnailView) BackgroundColor() objc.ID {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("backgroundColor"))
	return rv
}
func (p PDFThumbnailView) SetBackgroundColor(value objc.ID) {
	objc.Send[struct{}](p.ID, objc.Sel("setBackgroundColor:"), value)
}

// Returns a Boolean value indicating whether users can drag thumbnails (that
// is, re-order pages in the document) within the thumbnail view.
//
// # Return Value
//
// true if users can re-order pages by dragging thumbnails, false otherwise.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/allowsDragging
func (p PDFThumbnailView) AllowsDragging() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("allowsDragging"))
	return rv
}
func (p PDFThumbnailView) SetAllowsDragging(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAllowsDragging:"), value)
}

// Returns a Boolean value indicating whether users can select multiple
// thumbnails in the thumbnail view at one time.
//
// # Return Value
//
// true if users can select multiple thumbnails simultaneously, false
// otherwise.
//
// # Discussion
//
// By default, [PDFThumbnailView] allows only a single thumbnail to be
// selected at one time. When this is the case, you can get the PDF page that
// corresponds to the selected thumbnail using the [PDFView] method
// [CurrentPage].
//
// When multiple selections are enabled, however, you must use [SelectedPages]
// to get the pages that correspond to the set of selected thumbnails.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/allowsMultipleSelection
func (p PDFThumbnailView) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}
func (p PDFThumbnailView) SetAllowsMultipleSelection(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAllowsMultipleSelection:"), value)
}

// Returns an array of PDF pages that correspond to the selected thumbnails in
// the thumbnail view.
//
// # Return Value
//
// An array of PDF pages that correspond to the thumbnails selected in the
// thumbnail view.
//
// # Discussion
//
// If the thumbnail view allows multiple selections (if
// [AllowsMultipleSelection] returns true), you can use this method to get the
// PDF pages that correspond to the selected thumbnails.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/selectedPages
func (p PDFThumbnailView) SelectedPages() []PDFPage {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("selectedPages"))
	return objc.ConvertSlice(rv, func(id objc.ID) PDFPage {
		return PDFPageFromID(id)
	})
}
