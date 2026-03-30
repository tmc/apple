// Code generated from Apple documentation for PDFKit. DO NOT EDIT.
//go:build ios
// +build ios

package pdfkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// See: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/contentInset
func (p PDFThumbnailView) ContentInset() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("contentInset"))
	return objectivec.Object{ID: rv}
}
func (p PDFThumbnailView) SetContentInset(value objectivec.IObject) {
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
