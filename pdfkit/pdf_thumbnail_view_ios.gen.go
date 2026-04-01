// Code generated from Apple documentation for PDFKit. DO NOT EDIT.
//go:build ios
// +build ios

package pdfkit

import (
	"github.com/tmc/apple/objc"
)

// See: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/contentInset
func (p PDFThumbnailView) ContentInset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p.ID, objc.Sel("contentInset"))
	return rv
}
func (p PDFThumbnailView) SetContentInset(value unsafe.Pointer) {
	objc.Send[struct{}](p.ID, objc.Sel("setContentInset:"), value)
}

// See: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/layoutMode
func (p PDFThumbnailView) LayoutMode() PDFThumbnailLayoutMode {
	rv := objc.Send[PDFThumbnailLayoutMode](p.ID, objc.Sel("layoutMode"))
	return PDFThumbnailLayoutMode(rv)
}
func (p PDFThumbnailView) SetLayoutMode(value PDFThumbnailLayoutMode) {
	objc.Send[struct{}](p.ID, objc.Sel("setLayoutMode:"), value)
}
