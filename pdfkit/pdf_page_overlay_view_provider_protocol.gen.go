// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// PDFPageOverlayViewProvider protocol.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFPageOverlayViewProvider
type PDFPageOverlayViewProvider interface {
	objectivec.IObject

	// PdfViewOverlayViewForPage protocol.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFPageOverlayViewProvider/pdfView(_:overlayViewFor:)
	PdfViewOverlayViewForPage(view IPDFView, page IPDFPage) appkit.NSView
}

// PDFPageOverlayViewProviderObject wraps an existing Objective-C object that conforms to the PDFPageOverlayViewProvider protocol.
type PDFPageOverlayViewProviderObject struct {
	objectivec.Object
}

func (o PDFPageOverlayViewProviderObject) BaseObject() objectivec.Object {
	return o.Object
}

// PDFPageOverlayViewProviderObjectFromID constructs a [PDFPageOverlayViewProviderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func PDFPageOverlayViewProviderObjectFromID(id objc.ID) PDFPageOverlayViewProviderObject {
	return PDFPageOverlayViewProviderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/PDFKit/PDFPageOverlayViewProvider/pdfView(_:overlayViewFor:)
func (o PDFPageOverlayViewProviderObject) PdfViewOverlayViewForPage(view IPDFView, page IPDFPage) appkit.NSView {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("pdfView:overlayViewForPage:"), view, page)
	return appkit.NSViewFromID(rv)
}

// See: https://developer.apple.com/documentation/PDFKit/PDFPageOverlayViewProvider/pdfView(_:willDisplayOverlayView:for:)
func (o PDFPageOverlayViewProviderObject) PdfViewWillDisplayOverlayViewForPage(pdfView IPDFView, overlayView appkit.NSView, page IPDFPage) {
	objc.Send[struct{}](o.ID, objc.Sel("pdfView:willDisplayOverlayView:forPage:"), pdfView, overlayView, page)
}

// See: https://developer.apple.com/documentation/PDFKit/PDFPageOverlayViewProvider/pdfView(_:willEndDisplayingOverlayView:for:)
func (o PDFPageOverlayViewProviderObject) PdfViewWillEndDisplayingOverlayViewForPage(pdfView IPDFView, overlayView appkit.NSView, page IPDFPage) {
	objc.Send[struct{}](o.ID, objc.Sel("pdfView:willEndDisplayingOverlayView:forPage:"), pdfView, overlayView, page)
}
