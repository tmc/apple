// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

// Package pdfkit provides Go bindings for the PDFKit framework.
//
// Display and manipulate PDF documents in your apps.
//
// # Views
//
//   - PDFView: An object that encapsulates the functionality of PDF Kit into a single widget that you can add to your application using Interface Builder. ([PDFViewDelegate])
//   - PDFThumbnailView: An object that contains a set of thumbnails, each of which represents a page in a PDF document.
//
// # Content Model
//
//   - PDFDocument: An object that represents PDF data or a PDF file and defines methods for writing, searching, and selecting PDF data. ([PDFDocumentDelegate], [PDFDocumentPermissions], [PDFDocumentAttribute], [PDFDocumentWriteOption])
//   - PDFPage: , a subclass of , defines methods used to render PDF pages and work with annotations, text, and selections. ([PDFDisplayBox], [PDFDisplayDirection])
//   - PDFOutline: A  object is an element in a tree-structured hierarchy that can represent the structure of a PDF document.
//   - PDFSelection: A  object identifies a contiguous or noncontiguous selection of text in a PDF document.
//
// # Annotations
//
//   - Adding Widgets to a PDF Document: Add text, button, and choice widgets to a PDF document.
//   - Adding Custom Graphics to a PDF: Create and add custom annotation and page graphics to your PDF document.
//   - Custom Graphics: Demonstrates adding a watermark to a PDF page.
//   - PDF Widgets: Demonstrates adding widgets—interactive form elements—to a PDF document.
//   - PDFAnnotation: An annotation in a PDF document. ([PDFAnnotationSubtype], [PDFAction], [PDFDestination], [PDFAnnotationKey], [PDFBorder])
//
// # Protocols
//
//   - PDFPageOverlayViewProvider
//
// # Key Types
//
//   - [PDFView] - An object that encapsulates the functionality of PDF Kit into a single widget that you can add to your application using Interface Builder.
//   - [PDFAnnotation] - An annotation in a PDF document.
//   - [PDFDocument] - An object that represents PDF data or a PDF file and defines methods for writing, searching, and selecting PDF data.
//   - [PDFPage] - [PDFPage], a subclass of [NSObject], defines methods used to render PDF pages and work with annotations, text, and selections.
//   - [PDFSelection] - A [PDFSelection] object identifies a contiguous or noncontiguous selection of text in a PDF document.
//   - [PDFBorder] - An optional border for an annotation that lies completely within the annotation rectangle.
//   - [PDFAppearanceCharacteristics] - An object that represents appearance characteristics of a widget annotation.
//   - [PDFOutline] - A [PDFOutline] object is an element in a tree-structured hierarchy that can represent the structure of a PDF document.
//   - [PDFDestination] - A [PDFDestination] object describes a point on a PDF page.
//   - [PDFThumbnailView] - An object that contains a set of thumbnails, each of which represents a page in a PDF document.
//
// [PDFKit Documentation]: https://developer.apple.com/documentation/PDFKit
package pdfkit

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the PDFKit library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/PDFKit.framework/PDFKit",
	"/usr/lib/libPDFKit.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	fmt.Fprintf(os.Stderr, "warning: PDFKit: failed to load framework from any known path\n")
}
