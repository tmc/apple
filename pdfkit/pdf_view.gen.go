// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PDFView] class.
var (
	_PDFViewClass     PDFViewClass
	_PDFViewClassOnce sync.Once
)

func getPDFViewClass() PDFViewClass {
	_PDFViewClassOnce.Do(func() {
		_PDFViewClass = PDFViewClass{class: objc.GetClass("PDFView")}
	})
	return _PDFViewClass
}

// GetPDFViewClass returns the class object for PDFView.
func GetPDFViewClass() PDFViewClass {
	return getPDFViewClass()
}

type PDFViewClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PDFViewClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PDFViewClass) Alloc() PDFView {
	rv := objc.Send[PDFView](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// An object that encapsulates the functionality of PDF Kit into a single
// widget that you can add to your application using Interface Builder.
//
// # Overview
//
// [PDFView] may be the only class you need to deal with for adding PDF
// functionality to your application. It lets you display PDF data and allows
// users to select content, navigate through a document, set zoom level, and
// copy textual content to the Pasteboard. [PDFView] also keeps track of page
// history.
//
// You can subclass [PDFView] to create a custom PDF viewer.
//
// You can also create a custom PDF viewer by using the PDF Kit utility
// classes directly and not using [PDFView] at all.
//
// # Associating a Document with a View
//
//   - [PDFView.Document]: Returns the document associated with a [PDFView] object.
//   - [PDFView.SetDocument]
//
// # Navigating Within a Document
//
//   - [PDFView.CurrentPage]: Returns the current page.
//   - [PDFView.CurrentDestination]: Returns a [PDFDestination] object representing the current page and the current point in the view specified in page space.
//   - [PDFView.VisiblePages]: Returns an array of [PDFPage] objects that represent the currently visible pages.
//
// # Setting the Delegate
//
//   - [PDFView.Delegate]: Returns the view’s delegate.
//   - [PDFView.SetDelegate]
//
// # Instance Properties
//
//   - [PDFView.InMarkupMode]
//   - [PDFView.SetInMarkupMode]
//   - [PDFView.PageOverlayViewProvider]
//   - [PDFView.SetPageOverlayViewProvider]
//   - [PDFView.PageShadowsEnabled]
//   - [PDFView.SetPageShadowsEnabled]
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView
type PDFView struct {
	appkit.NSView
}

// PDFViewFromID constructs a [PDFView] from an objc.ID.
//
// An object that encapsulates the functionality of PDF Kit into a single
// widget that you can add to your application using Interface Builder.
func PDFViewFromID(id objc.ID) PDFView {
	return PDFView{NSView: appkit.NSViewFromID(id)}
}

// NOTE: PDFView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [PDFView] class.
//
// # Associating a Document with a View
//
//   - [IPDFView.Document]: Returns the document associated with a [PDFView] object.
//   - [IPDFView.SetDocument]
//
// # Navigating Within a Document
//
//   - [IPDFView.CurrentPage]: Returns the current page.
//   - [IPDFView.CurrentDestination]: Returns a [PDFDestination] object representing the current page and the current point in the view specified in page space.
//   - [IPDFView.VisiblePages]: Returns an array of [PDFPage] objects that represent the currently visible pages.
//
// # Setting the Delegate
//
//   - [IPDFView.Delegate]: Returns the view’s delegate.
//   - [IPDFView.SetDelegate]
//
// # Instance Properties
//
//   - [IPDFView.InMarkupMode]
//   - [IPDFView.SetInMarkupMode]
//   - [IPDFView.PageOverlayViewProvider]
//   - [IPDFView.SetPageOverlayViewProvider]
//   - [IPDFView.PageShadowsEnabled]
//   - [IPDFView.SetPageShadowsEnabled]
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView
type IPDFView interface {
	appkit.INSView

	// Topic: Associating a Document with a View

	// Returns the document associated with a [PDFView] object.
	Document() IPDFDocument
	SetDocument(value IPDFDocument)

	// Topic: Navigating Within a Document

	// Returns the current page.
	CurrentPage() IPDFPage
	// Returns a [PDFDestination] object representing the current page and the current point in the view specified in page space.
	CurrentDestination() IPDFDestination
	// Returns an array of [PDFPage] objects that represent the currently visible pages.
	VisiblePages() []PDFPage

	// Topic: Setting the Delegate

	// Returns the view’s delegate.
	Delegate() PDFViewDelegate
	SetDelegate(value PDFViewDelegate)

	// Topic: Instance Properties

	InMarkupMode() bool
	SetInMarkupMode(value bool)
	PageOverlayViewProvider() PDFPageOverlayViewProvider
	SetPageOverlayViewProvider(value PDFPageOverlayViewProvider)
	PageShadowsEnabled() bool
	SetPageShadowsEnabled(value bool)

	// A Boolean value indicating whether you can drag a file into the view.
	AcceptsDraggedFiles() bool
	SetAcceptsDraggedFiles(value bool)
	// A Boolean value indicating whether autoscaling is set.
	AutoScales() bool
	SetAutoScales(value bool)
	// The view’s background color.
	BackgroundColor() objc.ID
	SetBackgroundColor(value objc.ID)
	// Returns a Boolean value indicating whether the user can navigate to the previous page in the page history.
	CanGoBack() bool
	// Returns a Boolean value indicating whether the user can navigate to the next page in the page history.
	CanGoForward() bool
	// Returns a Boolean value indicating whether the user can navigate to the first page of the document.
	CanGoToFirstPage() bool
	// Returns a Boolean value indicating whether the user can navigate to the last page of the document.
	CanGoToLastPage() bool
	// Returns a Boolean value indicating whether the user can navigate to the next page of the document.
	CanGoToNextPage() bool
	// Returns a Boolean value indicating whether the user can navigate to the previous page of the document.
	CanGoToPreviousPage() bool
	// Returns a Boolean value indicating whether the user can magnify the view and zoom in.
	CanZoomIn() bool
	// Returns a Boolean value indicating whether the user can view an expanded area and zoom out.
	CanZoomOut() bool
	// The current selection.
	CurrentSelection() IPDFSelection
	SetCurrentSelection(value IPDFSelection)
	// The current style of display box.
	DisplayBox() PDFDisplayBox
	SetDisplayBox(value PDFDisplayBox)
	// The layout direction, either vertical or horizontal, for the given display mode.
	DisplayDirection() PDFDisplayDirection
	SetDisplayDirection(value PDFDisplayDirection)
	// The current display mode.
	DisplayMode() PDFDisplayMode
	SetDisplayMode(value PDFDisplayMode)
	// A Boolean value indicating whether the view will display the first page as a book cover (meaningful only when the document is in two-up or two-up continuous display mode).
	DisplaysAsBook() bool
	SetDisplaysAsBook(value bool)
	// A Boolean value indicating whether the view is displaying page breaks.
	DisplaysPageBreaks() bool
	SetDisplaysPageBreaks(value bool)
	// The presentation of pages from right-to-left.
	DisplaysRTL() bool
	SetDisplaysRTL(value bool)
	// The innermost view used by [PDFView] or by your [PDFView] subclass.
	DocumentView() objc.ID
	// A Boolean value indicating whether to turns on or off data detection, which adds annotations for detected URLs in a page.
	EnableDataDetectors() bool
	SetEnableDataDetectors(value bool)
	// Returns the array of selections that are highlighted using `setHighlightedSelections`.
	HighlightedSelections() []PDFSelection
	SetHighlightedSelections(value []PDFSelection)
	// The interpolation quality for images drawn into the [PDFView] context.
	InterpolationQuality() PDFInterpolationQuality
	SetInterpolationQuality(value PDFInterpolationQuality)
	// The maximum scaling factor for the PDF document.
	MaxScaleFactor() float64
	SetMaxScaleFactor(value float64)
	// The minimum scaling factor for the PDF document.
	MinScaleFactor() float64
	SetMinScaleFactor(value float64)
	// The spacing between pages as defined by the top, bottom, left, and right margins.
	PageBreakMargins() objectivec.IObject
	SetPageBreakMargins(value objectivec.IObject)
	// The current scale factor for the view.
	ScaleFactor() float64
	SetScaleFactor(value float64)
	// The “size to fit” scale factor that `autoScales` would use for scaling the current document and layout.
	ScaleFactorForSizeToFit() float64
	// Tells the PDF view that an annotation on the specified page has changed.
	AnnotationsChangedOnPage(page IPDFPage)
	// Returns the type of area the mouse cursor is over.
	AreaOfInterestForMouse(event objectivec.IObject) PDFAreaOfInterest
	// Returns the type of area for a specific cursor location point.
	AreaOfInterestForPoint(cursorLocation corefoundation.CGPoint) PDFAreaOfInterest
	// Clears the selection.
	ClearSelection()
	// Converts a point from page space to view space.
	ConvertPointFromPage(point corefoundation.CGPoint, page IPDFPage) corefoundation.CGPoint
	// Converts a point from view space to page space.
	ConvertPointToPage(point corefoundation.CGPoint, page IPDFPage) corefoundation.CGPoint
	// Converts a rectangle from page space to view space.
	ConvertRectFromPage(rect corefoundation.CGRect, page IPDFPage) corefoundation.CGRect
	// Converts a rectangle from view space to page space.
	ConvertRectToPage(rect corefoundation.CGRect, page IPDFPage) corefoundation.CGRect
	// Copies the text in the selection, if any, to the Pasteboard.
	Copy(sender objectivec.IObject)
	// Perform post-page rendering for a page rendered to a context.
	DrawPagePostToContext(page IPDFPage, context coregraphics.CGContextRef)
	// Draw and render a visible page to a context.
	DrawPageToContext(page IPDFPage, context coregraphics.CGContextRef)
	// Navigates back one step in the page history.
	GoBack(sender objectivec.IObject)
	// Navigates forward one step in the page history.
	GoForward(sender objectivec.IObject)
	// Navigates to the specified destination.
	GoToDestination(destination IPDFDestination)
	// Navigates to the first page of the document.
	GoToFirstPage(sender objectivec.IObject)
	// Navigates to the last page of the document.
	GoToLastPage(sender objectivec.IObject)
	// Navigates to the next page of the document.
	GoToNextPage(sender objectivec.IObject)
	// Scrolls to the specified page.
	GoToPage(page IPDFPage)
	// Navigates to the previous page of the document.
	GoToPreviousPage(sender objectivec.IObject)
	// Navigates to the specified rectangle on the specified page.
	GoToRectOnPage(rect corefoundation.CGRect, page IPDFPage)
	// Scrolls to the first character of the specified selection.
	GoToSelection(selection IPDFSelection)
	// Performs layout of the inner views.
	LayoutDocumentView()
	// Returns the page containing a point specified in view coordinates.
	PageForPointNearest(point corefoundation.CGPoint, nearest bool) IPDFPage
	// Performs the specified action.
	PerformAction(action IPDFAction)
	// Prints the document with the specified printer information.
	PrintWithInfoAutoRotate(printInfo appkit.NSPrintInfo, doRotate bool)
	// Prints the document with the specified printer and page-scaling information.
	PrintWithInfoAutoRotatePageScaling(printInfo appkit.NSPrintInfo, doRotate bool, scale PDFPrintScalingMode)
	// Returns the size needed to display a row of the current document page.
	RowSizeForPage(page IPDFPage) corefoundation.CGSize
	// Scrolls the view until the selection is visible.
	ScrollSelectionToVisible(sender objectivec.IObject)
	// Sets the current selection, in an animated way, if desired.
	SetCurrentSelectionAnimate(selection IPDFSelection, animate bool)
	// Sets the type of mouse cursor according to the type of area the mouse cursor is over.
	SetCursorForAreaOfInterest(area PDFAreaOfInterest)
	// Zooms in by increasing the scaling factor.
	ZoomIn(sender objectivec.IObject)
	// Zooms out by decreasing the scaling factor.
	ZoomOut(sender objectivec.IObject)
}

// Init initializes the instance.
func (p PDFView) Init() PDFView {
	rv := objc.Send[PDFView](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PDFView) Autorelease() PDFView {
	rv := objc.Send[PDFView](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFView creates a new PDFView instance.
func NewPDFView() PDFView {
	class := getPDFViewClass()
	rv := objc.Send[PDFView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Tells the PDF view that an annotation on the specified page has changed.
//
// # Discussion
//
// When the [PDFView] object receives this message, it rescans for tool tips
// and pop-ups and informs the [PDFThumbailView] objects so the thumbnail
// images can be redrawn.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/annotationsChanged(on:)
func (p PDFView) AnnotationsChangedOnPage(page IPDFPage) {
	objc.Send[objc.ID](p.ID, objc.Sel("annotationsChangedOnPage:"), page)
}

// Returns the type of area the mouse cursor is over.
//
// # Discussion
//
// The [PDFAreaOfInterest] enumeration defines the various area types. This
// method is for custom subclasses of the [PDFView] class. Use it if you
// override the [NSResponder] class’s [mouseMoved(with:)] method or related
// methods.
//
// Refer to [Constants] for the various values of the area-of-interest
// constants. Each of these constants contributes to the value of the
// [PDFAreaOfInterest] bit field.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/areaOfInterest(forMouse:)
//
// [mouseMoved(with:)]: https://developer.apple.com/documentation/AppKit/NSResponder/mouseMoved(with:)
func (p PDFView) AreaOfInterestForMouse(event objectivec.IObject) PDFAreaOfInterest {
	rv := objc.Send[PDFAreaOfInterest](p.ID, objc.Sel("areaOfInterestForMouse:"), event)
	return PDFAreaOfInterest(rv)
}

// Returns the type of area for a specific cursor location point.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/areaOfInterest(for:)
func (p PDFView) AreaOfInterestForPoint(cursorLocation corefoundation.CGPoint) PDFAreaOfInterest {
	rv := objc.Send[PDFAreaOfInterest](p.ID, objc.Sel("areaOfInterestForPoint:"), cursorLocation)
	return PDFAreaOfInterest(rv)
}

// Clears the selection.
//
// # Discussion
//
// The view redraws as necessary but does not scroll. This call is equivalent
// to calling `[PDFView NULL].`
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/clearSelection()
func (p PDFView) ClearSelection() {
	objc.Send[objc.ID](p.ID, objc.Sel("clearSelection"))
}

// Converts a point from page space to view space.
//
// # Discussion
//
// Page space is a 72 dpi coordinate system with the origin at the lower-left
// corner of the current page. View space is a coordinate system with the
// origin at the lower-left corner of the current PDF view.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/convert(_:from:)-4evlx
func (p PDFView) ConvertPointFromPage(point corefoundation.CGPoint, page IPDFPage) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](p.ID, objc.Sel("convertPoint:fromPage:"), point, page)
	return corefoundation.CGPoint(rv)
}

// Converts a point from view space to page space.
//
// # Discussion
//
// Page space is a 72 dpi coordinate system with the origin at the lower-left
// corner of the current page. View space is a coordinate system with the
// origin at the lower-left corner of the current PDF view.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/convert(_:to:)-9twqk
func (p PDFView) ConvertPointToPage(point corefoundation.CGPoint, page IPDFPage) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](p.ID, objc.Sel("convertPoint:toPage:"), point, page)
	return corefoundation.CGPoint(rv)
}

// Converts a rectangle from page space to view space.
//
// # Discussion
//
// Page space is a 72 dpi coordinate system with the origin at the lower-left
// corner of the current page. View space is a coordinate system with the
// origin at the lower-left corner of the current PDF view.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/convert(_:from:)-9xv1z
func (p PDFView) ConvertRectFromPage(rect corefoundation.CGRect, page IPDFPage) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p.ID, objc.Sel("convertRect:fromPage:"), rect, page)
	return corefoundation.CGRect(rv)
}

// Converts a rectangle from view space to page space.
//
// # Discussion
//
// Page space is a 72 dpi coordinate system with the origin at the lower-left
// corner of the current page. View space is a coordinate system with the
// origin at the lower-left corner of the current PDF view.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/convert(_:to:)-8cp0c
func (p PDFView) ConvertRectToPage(rect corefoundation.CGRect, page IPDFPage) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p.ID, objc.Sel("convertRect:toPage:"), rect, page)
	return corefoundation.CGRect(rv)
}

// Copies the text in the selection, if any, to the Pasteboard.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/copy(_:)
func (p PDFView) Copy(sender objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("copy:"), sender)
}

// Perform post-page rendering for a page rendered to a context.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/drawPagePost(_:to:)
func (p PDFView) DrawPagePostToContext(page IPDFPage, context coregraphics.CGContextRef) {
	objc.Send[objc.ID](p.ID, objc.Sel("drawPagePost:toContext:"), page, context)
}

// Draw and render a visible page to a context.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/draw(_:to:)
func (p PDFView) DrawPageToContext(page IPDFPage, context coregraphics.CGContextRef) {
	objc.Send[objc.ID](p.ID, objc.Sel("drawPage:toContext:"), page, context)
}

// Navigates back one step in the page history.
//
// # Discussion
//
// The page history gets built as your application calls navigation methods
// such as [GoToDestination] and [GoToLastPage].
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/goBack(_:)
func (p PDFView) GoBack(sender objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("goBack:"), sender)
}

// Navigates forward one step in the page history.
//
// # Discussion
//
// The page history gets built as your application calls navigation methods
// such as [GoToDestination] and [GoToLastPage].
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/goForward(_:)
func (p PDFView) GoForward(sender objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("goForward:"), sender)
}

// Navigates to the specified destination.
//
// # Discussion
//
// Destinations include a page and a point on the page specified in page
// space.
//
// Page space is a 72 dpi coordinate system with the origin at the lower-left
// corner of the current page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/go(to:)-5lh5d
func (p PDFView) GoToDestination(destination IPDFDestination) {
	objc.Send[objc.ID](p.ID, objc.Sel("goToDestination:"), destination)
}

// Navigates to the first page of the document.
//
// # Discussion
//
// PDF Kit records the move in its page history.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/goToFirstPage(_:)
func (p PDFView) GoToFirstPage(sender objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("goToFirstPage:"), sender)
}

// Navigates to the last page of the document.
//
// # Discussion
//
// PDF Kit records the move in its page history.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/goToLastPage(_:)
func (p PDFView) GoToLastPage(sender objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("goToLastPage:"), sender)
}

// Navigates to the next page of the document.
//
// # Discussion
//
// PDF Kit records the move in its page history.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/goToNextPage(_:)
func (p PDFView) GoToNextPage(sender objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("goToNextPage:"), sender)
}

// Scrolls to the specified page.
//
// # Discussion
//
// PDF Kit records the move in its page history.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/go(to:)-6x8y2
func (p PDFView) GoToPage(page IPDFPage) {
	objc.Send[objc.ID](p.ID, objc.Sel("goToPage:"), page)
}

// Navigates to the previous page of the document.
//
// # Discussion
//
// PDF Kit records the move in its page history.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/goToPreviousPage(_:)
func (p PDFView) GoToPreviousPage(sender objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("goToPreviousPage:"), sender)
}

// Navigates to the specified rectangle on the specified page.
//
// # Discussion
//
// If the specified rectangle is already visible, this method does nothing.
// This allows you to scroll the [PDFView] object to a specific
// [PDFAnnotation] or [PDFSelection] object, because both of these objects
// have bounds methods that return an annotation or selection position in page
// space.
//
// Note that `rect` is specified in page-space coordinates. Page space is a 72
// dpi coordinate system with the origin at the lower-left corner of the
// current page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/go(to:on:)
func (p PDFView) GoToRectOnPage(rect corefoundation.CGRect, page IPDFPage) {
	objc.Send[objc.ID](p.ID, objc.Sel("goToRect:onPage:"), rect, page)
}

// Scrolls to the first character of the specified selection.
//
// # Discussion
//
// PDF Kit records the move in its page history.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/go(to:)-3t5go
func (p PDFView) GoToSelection(selection IPDFSelection) {
	objc.Send[objc.ID](p.ID, objc.Sel("goToSelection:"), selection)
}

// Performs layout of the inner views.
//
// # Discussion
//
// The [PDFView] actually contains several subviews, such as the document view
// (where the PDF is actually drawn) and a “matte view” (which may appear
// as a gray area around the PDF content, depending on the scaling). Changes
// to the PDF content may require changes to these inner views, so you must
// call this method explicitly if you use PDF Kit utility classes to add or
// remove a page, rotate a page, or perform other operations affecting visible
// layout.
//
// This method is called automatically from [PDFView] methods that affect the
// visible layout (such as `setDocument(_:)`, `setDisplayBox(_:)` or
// [ZoomIn]).
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/layoutDocumentView()
func (p PDFView) LayoutDocumentView() {
	objc.Send[objc.ID](p.ID, objc.Sel("layoutDocumentView"))
}

// Returns the page containing a point specified in view coordinates.
//
// # Discussion
//
// Returns [NULL] if there’s no page at the specified point and `nearest` is
// set to false.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/page(for:nearest:)
func (p PDFView) PageForPointNearest(point corefoundation.CGPoint, nearest bool) IPDFPage {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("pageForPoint:nearest:"), point, nearest)
	return PDFPageFromID(rv)
}

// Performs the specified action.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/perform(_:)
func (p PDFView) PerformAction(action IPDFAction) {
	objc.Send[objc.ID](p.ID, objc.Sel("performAction:"), action)
}

// Prints the document with the specified printer information.
//
// # Discussion
//
// If `autoRotate` is set to true, then ths method ignores the orientation
// attribute in the [NSPrintInfo] object and instead chooses the orientation
// that best fits the page to the paper size. This orientation occurs on a
// page-by-page basis.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/print(with:autoRotate:)
func (p PDFView) PrintWithInfoAutoRotate(printInfo appkit.NSPrintInfo, doRotate bool) {
	objc.Send[objc.ID](p.ID, objc.Sel("printWithInfo:autoRotate:"), printInfo, doRotate)
}

// Prints the document with the specified printer and page-scaling
// information.
//
// # Discussion
//
// If `pageScaling` is set to `kPDFPrintPageScaleToFit`, each page is scaled
// up or down to best fit the paper size. If `pageScaling` is set to
// `kPDFPrintPageScaleDownToFit`, only large pages are scaled down to fit;
// small pages are not scaled up to fit. Specifying `kPDFPrintPageScaleNone`
// for `pageScaling` is equivalent to calling [PrintWithInfoAutoRotate]. See
// PDFDocument for more information on page-scaling types.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/print(with:autoRotate:pageScaling:)
func (p PDFView) PrintWithInfoAutoRotatePageScaling(printInfo appkit.NSPrintInfo, doRotate bool, scale PDFPrintScalingMode) {
	objc.Send[objc.ID](p.ID, objc.Sel("printWithInfo:autoRotate:pageScaling:"), printInfo, doRotate, scale)
}

// Returns the size needed to display a row of the current document page.
//
// # Discussion
//
// The size is dependent on the current scale factor and display attributes.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/rowSize(for:)
func (p PDFView) RowSizeForPage(page IPDFPage) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](p.ID, objc.Sel("rowSizeForPage:"), page)
	return corefoundation.CGSize(rv)
}

// Scrolls the view until the selection is visible.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/scrollSelectionToVisible(_:)
func (p PDFView) ScrollSelectionToVisible(sender objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("scrollSelectionToVisible:"), sender)
}

// Sets the current selection, in an animated way, if desired.
//
// # Discussion
//
// This method behaves as `setCurrentSelection(_:)`, but with the addition of
// animation, if `animate` is true. The animation serves to draw the user’s
// attention to the new selection, which can be useful when implementing
// search.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/setCurrentSelection(_:animate:)
func (p PDFView) SetCurrentSelectionAnimate(selection IPDFSelection, animate bool) {
	objc.Send[objc.ID](p.ID, objc.Sel("setCurrentSelection:animate:"), selection, animate)
}

// Sets the type of mouse cursor according to the type of area the mouse
// cursor is over.
//
// # Discussion
//
// This method is especially useful for custom subclasses of the [PDFView]
// class.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/setCursorFor(_:)
func (p PDFView) SetCursorForAreaOfInterest(area PDFAreaOfInterest) {
	objc.Send[objc.ID](p.ID, objc.Sel("setCursorForAreaOfInterest:"), area)
}

// Zooms in by increasing the scaling factor.
//
// # Discussion
//
// Each invocation of `zoomIn` muliplies the scaling factor by the square root
// of 2.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/zoomIn(_:)
func (p PDFView) ZoomIn(sender objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("zoomIn:"), sender)
}

// Zooms out by decreasing the scaling factor.
//
// # Discussion
//
// Each invocation of `zoomOut` divides the scaling factor by the square root
// of 2.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/zoomOut(_:)
func (p PDFView) ZoomOut(sender objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("zoomOut:"), sender)
}

// Returns the document associated with a [PDFView] object.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/document
func (p PDFView) Document() IPDFDocument {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("document"))
	return PDFDocumentFromID(objc.ID(rv))
}
func (p PDFView) SetDocument(value IPDFDocument) {
	objc.Send[struct{}](p.ID, objc.Sel("setDocument:"), value)
}

// Returns the current page.
//
// # Discussion
//
// When there are two pages in the view in a two-up mode, “current page”
// is the left page. For continuous modes, returns the page crossing a
// horizontal line halfway between the view’s top and bottom bounds.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/currentPage
func (p PDFView) CurrentPage() IPDFPage {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("currentPage"))
	return PDFPageFromID(objc.ID(rv))
}

// Returns a [PDFDestination] object representing the current page and the
// current point in the view specified in page space.
//
// # Discussion
//
// Page space is a 72 dpi coordinate system with the origin at the lower-left
// corner of the current page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/currentDestination
func (p PDFView) CurrentDestination() IPDFDestination {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("currentDestination"))
	return PDFDestinationFromID(objc.ID(rv))
}

// Returns an array of [PDFPage] objects that represent the currently visible
// pages.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/visiblePages
func (p PDFView) VisiblePages() []PDFPage {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("visiblePages"))
	return objc.ConvertSlice(rv, func(id objc.ID) PDFPage {
		return PDFPageFromID(id)
	})
}

// Returns the view’s delegate.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/delegate
func (p PDFView) Delegate() PDFViewDelegate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("delegate"))
	return PDFViewDelegateObjectFromID(rv)
}
func (p PDFView) SetDelegate(value PDFViewDelegate) {
	objc.Send[struct{}](p.ID, objc.Sel("setDelegate:"), value)
}

// See: https://developer.apple.com/documentation/PDFKit/PDFView/isInMarkupMode
func (p PDFView) InMarkupMode() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isInMarkupMode"))
	return rv
}
func (p PDFView) SetInMarkupMode(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setInMarkupMode:"), value)
}

// See: https://developer.apple.com/documentation/PDFKit/PDFView/pageOverlayViewProvider
func (p PDFView) PageOverlayViewProvider() PDFPageOverlayViewProvider {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("pageOverlayViewProvider"))
	return PDFPageOverlayViewProviderObjectFromID(rv)
}
func (p PDFView) SetPageOverlayViewProvider(value PDFPageOverlayViewProvider) {
	objc.Send[struct{}](p.ID, objc.Sel("setPageOverlayViewProvider:"), value)
}

// See: https://developer.apple.com/documentation/PDFKit/PDFView/pageShadowsEnabled
func (p PDFView) PageShadowsEnabled() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("pageShadowsEnabled"))
	return rv
}
func (p PDFView) SetPageShadowsEnabled(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("enablePageShadows:"), value)
}

// A Boolean value indicating whether you can drag a file into the view.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/acceptsDraggedFiles
func (p PDFView) AcceptsDraggedFiles() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("acceptsDraggedFiles"))
	return rv
}
func (p PDFView) SetAcceptsDraggedFiles(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAcceptsDraggedFiles:"), value)
}

// A Boolean value indicating whether autoscaling is set.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/autoScales
func (p PDFView) AutoScales() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("autoScales"))
	return rv
}
func (p PDFView) SetAutoScales(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAutoScales:"), value)
}

// The view’s background color.
//
// # Discussion
//
// A view’s background is the area displayed to either side of a PDF
// document’s pages. The background also appears between pages when page
// breaks are enabled. The default color is a 50% gray.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/backgroundColor
func (p PDFView) BackgroundColor() objc.ID {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("backgroundColor"))
	return rv
}
func (p PDFView) SetBackgroundColor(value objc.ID) {
	objc.Send[struct{}](p.ID, objc.Sel("setBackgroundColor:"), value)
}

// Returns a Boolean value indicating whether the user can navigate to the
// previous page in the page history.
//
// # Discussion
//
// The page history gets built as your application calls navigation methods
// such as [GoToDestination] and [GoToLastPage].
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/canGoBack
func (p PDFView) CanGoBack() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canGoBack"))
	return rv
}

// Returns a Boolean value indicating whether the user can navigate to the
// next page in the page history.
//
// # Discussion
//
// The page history gets built as your application calls navigation methods
// such as [GoToDestination] and [GoToLastPage].
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/canGoForward
func (p PDFView) CanGoForward() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canGoForward"))
	return rv
}

// Returns a Boolean value indicating whether the user can navigate to the
// first page of the document.
//
// # Discussion
//
// The return value will be true unless the view is already displaying the
// first page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/canGoToFirstPage
func (p PDFView) CanGoToFirstPage() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canGoToFirstPage"))
	return rv
}

// Returns a Boolean value indicating whether the user can navigate to the
// last page of the document.
//
// # Discussion
//
// The return value will be true unless the view is already displaying the
// last page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/canGoToLastPage
func (p PDFView) CanGoToLastPage() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canGoToLastPage"))
	return rv
}

// Returns a Boolean value indicating whether the user can navigate to the
// next page of the document.
//
// # Discussion
//
// The return value will be true unless the view is displaying the last page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/canGoToNextPage
func (p PDFView) CanGoToNextPage() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canGoToNextPage"))
	return rv
}

// Returns a Boolean value indicating whether the user can navigate to the
// previous page of the document.
//
// # Discussion
//
// The return value will be true unless the view is displaying the first page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/canGoToPreviousPage
func (p PDFView) CanGoToPreviousPage() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canGoToPreviousPage"))
	return rv
}

// Returns a Boolean value indicating whether the user can magnify the view
// and zoom in.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/canZoomIn
func (p PDFView) CanZoomIn() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canZoomIn"))
	return rv
}

// Returns a Boolean value indicating whether the user can view an expanded
// area and zoom out.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/canZoomOut
func (p PDFView) CanZoomOut() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canZoomOut"))
	return rv
}

// The current selection.
//
// # Discussion
//
// Returns [NULL] if no selection exists.
//
// Note that this method returns the actual instance of the current
// [PDFSelection] object. Therefore, if you want to modify it, you should make
// a copy of the returned selection and modify that, instead.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/currentSelection
func (p PDFView) CurrentSelection() IPDFSelection {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("currentSelection"))
	return PDFSelectionFromID(objc.ID(rv))
}
func (p PDFView) SetCurrentSelection(value IPDFSelection) {
	objc.Send[struct{}](p.ID, objc.Sel("setCurrentSelection:"), value)
}

// The current style of display box.
//
// # Discussion
//
// The available values for display boxes are defined in the Constants section
// in the [PDFPage] class.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/displayBox
func (p PDFView) DisplayBox() PDFDisplayBox {
	rv := objc.Send[PDFDisplayBox](p.ID, objc.Sel("displayBox"))
	return PDFDisplayBox(rv)
}
func (p PDFView) SetDisplayBox(value PDFDisplayBox) {
	objc.Send[struct{}](p.ID, objc.Sel("setDisplayBox:"), value)
}

// The layout direction, either vertical or horizontal, for the given display
// mode.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/displayDirection
func (p PDFView) DisplayDirection() PDFDisplayDirection {
	rv := objc.Send[PDFDisplayDirection](p.ID, objc.Sel("displayDirection"))
	return PDFDisplayDirection(rv)
}
func (p PDFView) SetDisplayDirection(value PDFDisplayDirection) {
	objc.Send[struct{}](p.ID, objc.Sel("setDisplayDirection:"), value)
}

// The current display mode.
//
// # Discussion
//
// See [Constants] for possible values.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/displayMode
func (p PDFView) DisplayMode() PDFDisplayMode {
	rv := objc.Send[PDFDisplayMode](p.ID, objc.Sel("displayMode"))
	return PDFDisplayMode(rv)
}
func (p PDFView) SetDisplayMode(value PDFDisplayMode) {
	objc.Send[struct{}](p.ID, objc.Sel("setDisplayMode:"), value)
}

// A Boolean value indicating whether the view will display the first page as
// a book cover (meaningful only when the document is in two-up or two-up
// continuous display mode).
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/displaysAsBook
func (p PDFView) DisplaysAsBook() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("displaysAsBook"))
	return rv
}
func (p PDFView) SetDisplaysAsBook(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setDisplaysAsBook:"), value)
}

// A Boolean value indicating whether the view is displaying page breaks.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/displaysPageBreaks
func (p PDFView) DisplaysPageBreaks() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("displaysPageBreaks"))
	return rv
}
func (p PDFView) SetDisplaysPageBreaks(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setDisplaysPageBreaks:"), value)
}

// The presentation of pages from right-to-left.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/displaysRTL
func (p PDFView) DisplaysRTL() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("displaysRTL"))
	return rv
}
func (p PDFView) SetDisplaysRTL(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setDisplaysRTL:"), value)
}

// The innermost view used by [PDFView] or by your [PDFView] subclass.
//
// # Discussion
//
// The innermost view is the one displaying the visible document pages. This
// method is useful when converting coordinates from one view to another.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/documentView
func (p PDFView) DocumentView() objc.ID {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("documentView"))
	return rv
}

// A Boolean value indicating whether to turns on or off data detection, which
// adds annotations for detected URLs in a page.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/enableDataDetectors
func (p PDFView) EnableDataDetectors() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("enableDataDetectors"))
	return rv
}
func (p PDFView) SetEnableDataDetectors(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setEnableDataDetectors:"), value)
}

// Returns the array of selections that are highlighted using
// `setHighlightedSelections`.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/highlightedSelections
func (p PDFView) HighlightedSelections() []PDFSelection {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("highlightedSelections"))
	return objc.ConvertSlice(rv, func(id objc.ID) PDFSelection {
		return PDFSelectionFromID(id)
	})
}
func (p PDFView) SetHighlightedSelections(value []PDFSelection) {
	objc.Send[struct{}](p.ID, objc.Sel("setHighlightedSelections:"), objectivec.IObjectSliceToNSArray(value))
}

// The interpolation quality for images drawn into the [PDFView] context.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/interpolationQuality
func (p PDFView) InterpolationQuality() PDFInterpolationQuality {
	rv := objc.Send[PDFInterpolationQuality](p.ID, objc.Sel("interpolationQuality"))
	return PDFInterpolationQuality(rv)
}
func (p PDFView) SetInterpolationQuality(value PDFInterpolationQuality) {
	objc.Send[struct{}](p.ID, objc.Sel("setInterpolationQuality:"), value)
}

// The maximum scaling factor for the PDF document.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/maxScaleFactor
func (p PDFView) MaxScaleFactor() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("maxScaleFactor"))
	return rv
}
func (p PDFView) SetMaxScaleFactor(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setMaxScaleFactor:"), value)
}

// The minimum scaling factor for the PDF document.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/minScaleFactor
func (p PDFView) MinScaleFactor() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("minScaleFactor"))
	return rv
}
func (p PDFView) SetMinScaleFactor(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setMinScaleFactor:"), value)
}

// The spacing between pages as defined by the top, bottom, left, and right
// margins.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/pageBreakMargins
func (p PDFView) PageBreakMargins() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("pageBreakMargins"))
	return objectivec.Object{ID: rv}
}
func (p PDFView) SetPageBreakMargins(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setPageBreakMargins:"), value)
}

// The current scale factor for the view.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/scaleFactor
func (p PDFView) ScaleFactor() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("scaleFactor"))
	return rv
}
func (p PDFView) SetScaleFactor(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setScaleFactor:"), value)
}

// The “size to fit” scale factor that `autoScales` would use for scaling
// the current document and layout.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/scaleFactorForSizeToFit
func (p PDFView) ScaleFactorForSizeToFit() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("scaleFactorForSizeToFit"))
	return rv
}
