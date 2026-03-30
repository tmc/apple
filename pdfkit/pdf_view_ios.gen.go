// Code generated from Apple documentation for PDFKit. DO NOT EDIT.
//go:build ios
// +build ios

package pdfkit

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Changes the scroll view to use a [UIPageViewController] to layout and
// navigate pages.
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/usePageViewController(_:withViewOptions:)
func (p PDFView) UsePageViewControllerWithViewOptions(enable bool, viewOptions foundation.INSDictionary) {
	objc.Send[objc.ID](p.ID, objc.Sel("usePageViewController:withViewOptions:"), enable, viewOptions)
}

// See: https://developer.apple.com/documentation/PDFKit/PDFView/findInteraction
func (p PDFView) FindInteraction() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("findInteraction"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/PDFKit/PDFView/isFindInteractionEnabled
func (p PDFView) FindInteractionEnabled() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isFindInteractionEnabled"))
	return rv
}
func (p PDFView) SetFindInteractionEnabled(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setFindInteractionEnabled:"), value)
}

// A Boolean value indicating whether the scroll view is using a
// [UIPageViewController].
//
// See: https://developer.apple.com/documentation/PDFKit/PDFView/isUsingPageViewController
func (p PDFView) IsUsingPageViewController() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isUsingPageViewController"))
	return rv
}
