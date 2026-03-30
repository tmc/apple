// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/PDFKit/PDFAccessPermissions
type PDFAccessPermissions uint

const (
	PDFAllowsCommenting           PDFAccessPermissions = 64
	PDFAllowsContentAccessibility PDFAccessPermissions = 32
	PDFAllowsContentCopying       PDFAccessPermissions = 16
	PDFAllowsDocumentAssembly     PDFAccessPermissions = 8
	PDFAllowsDocumentChanges      PDFAccessPermissions = 4
	PDFAllowsFormFieldEntry       PDFAccessPermissions = 128
	PDFAllowsHighQualityPrinting  PDFAccessPermissions = 2
	PDFAllowsLowQualityPrinting   PDFAccessPermissions = 1
)

func (e PDFAccessPermissions) String() string {
	switch e {
	case PDFAllowsCommenting:
		return "PDFAllowsCommenting"
	case PDFAllowsContentAccessibility:
		return "PDFAllowsContentAccessibility"
	case PDFAllowsContentCopying:
		return "PDFAllowsContentCopying"
	case PDFAllowsDocumentAssembly:
		return "PDFAllowsDocumentAssembly"
	case PDFAllowsDocumentChanges:
		return "PDFAllowsDocumentChanges"
	case PDFAllowsFormFieldEntry:
		return "PDFAllowsFormFieldEntry"
	case PDFAllowsHighQualityPrinting:
		return "PDFAllowsHighQualityPrinting"
	case PDFAllowsLowQualityPrinting:
		return "PDFAllowsLowQualityPrinting"
	default:
		return fmt.Sprintf("PDFAccessPermissions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/PDFKit/PDFActionNamedName
type PDFActionNamedName int

const (
	// KPDFActionNamedFind: The Find action.
	KPDFActionNamedFind PDFActionNamedName = 8
	// KPDFActionNamedFirstPage: The First Page action.
	KPDFActionNamedFirstPage PDFActionNamedName = 3
	// KPDFActionNamedGoBack: The Go Back action.
	KPDFActionNamedGoBack PDFActionNamedName = 5
	// KPDFActionNamedGoForward: The Go Forward action.
	KPDFActionNamedGoForward PDFActionNamedName = 6
	// KPDFActionNamedGoToPage: The Go to Page action.
	KPDFActionNamedGoToPage PDFActionNamedName = 7
	// KPDFActionNamedLastPage: The Last Page action.
	KPDFActionNamedLastPage PDFActionNamedName = 4
	// KPDFActionNamedNextPage: The Next Page action.
	KPDFActionNamedNextPage PDFActionNamedName = 1
	// KPDFActionNamedNone: The action has no name.
	KPDFActionNamedNone PDFActionNamedName = 0
	// KPDFActionNamedPreviousPage: The Previous Page action.
	KPDFActionNamedPreviousPage PDFActionNamedName = 2
	// KPDFActionNamedPrint: The Print action.
	KPDFActionNamedPrint PDFActionNamedName = 9
	// KPDFActionNamedZoomIn: The Zoom In action.
	KPDFActionNamedZoomIn PDFActionNamedName = 10
	// KPDFActionNamedZoomOut: The Zoom Out action.
	KPDFActionNamedZoomOut PDFActionNamedName = 11
)

func (e PDFActionNamedName) String() string {
	switch e {
	case KPDFActionNamedFind:
		return "KPDFActionNamedFind"
	case KPDFActionNamedFirstPage:
		return "KPDFActionNamedFirstPage"
	case KPDFActionNamedGoBack:
		return "KPDFActionNamedGoBack"
	case KPDFActionNamedGoForward:
		return "KPDFActionNamedGoForward"
	case KPDFActionNamedGoToPage:
		return "KPDFActionNamedGoToPage"
	case KPDFActionNamedLastPage:
		return "KPDFActionNamedLastPage"
	case KPDFActionNamedNextPage:
		return "KPDFActionNamedNextPage"
	case KPDFActionNamedNone:
		return "KPDFActionNamedNone"
	case KPDFActionNamedPreviousPage:
		return "KPDFActionNamedPreviousPage"
	case KPDFActionNamedPrint:
		return "KPDFActionNamedPrint"
	case KPDFActionNamedZoomIn:
		return "KPDFActionNamedZoomIn"
	case KPDFActionNamedZoomOut:
		return "KPDFActionNamedZoomOut"
	default:
		return fmt.Sprintf("PDFActionNamedName(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/PDFKit/PDFAreaOfInterest
type PDFAreaOfInterest int

const (
	// KPDFAnnotationArea: The mouse is over an annotation.
	KPDFAnnotationArea PDFAreaOfInterest = 4
	KPDFAnyArea        PDFAreaOfInterest = 9223372036854775807
	// KPDFControlArea: The mouse is over a control.
	KPDFControlArea PDFAreaOfInterest = 16
	// KPDFIconArea: The mouse is over an icon.
	KPDFIconArea PDFAreaOfInterest = 64
	// KPDFImageArea: The mouse is over an image.
	KPDFImageArea PDFAreaOfInterest = 256
	// KPDFLinkArea: The mouse is over a link.
	KPDFLinkArea PDFAreaOfInterest = 8
	// KPDFNoArea: The mouse is over an undefined area.
	KPDFNoArea PDFAreaOfInterest = 0
	// KPDFPageArea: The mouse is over a page.
	KPDFPageArea PDFAreaOfInterest = 1
	// KPDFPopupArea: The mouse is over a popup menu.
	KPDFPopupArea PDFAreaOfInterest = 128
	// KPDFTextArea: The mouse is over text.
	KPDFTextArea PDFAreaOfInterest = 2
	// KPDFTextFieldArea: The mouse is over a text field.
	KPDFTextFieldArea PDFAreaOfInterest = 32
)

func (e PDFAreaOfInterest) String() string {
	switch e {
	case KPDFAnnotationArea:
		return "KPDFAnnotationArea"
	case KPDFAnyArea:
		return "KPDFAnyArea"
	case KPDFControlArea:
		return "KPDFControlArea"
	case KPDFIconArea:
		return "KPDFIconArea"
	case KPDFImageArea:
		return "KPDFImageArea"
	case KPDFLinkArea:
		return "KPDFLinkArea"
	case KPDFNoArea:
		return "KPDFNoArea"
	case KPDFPageArea:
		return "KPDFPageArea"
	case KPDFPopupArea:
		return "KPDFPopupArea"
	case KPDFTextArea:
		return "KPDFTextArea"
	case KPDFTextFieldArea:
		return "KPDFTextFieldArea"
	default:
		return fmt.Sprintf("PDFAreaOfInterest(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/PDFKit/PDFBorderStyle
type PDFBorderStyle int

const (
	// KPDFBorderStyleBeveled: Beveled border.
	KPDFBorderStyleBeveled PDFBorderStyle = 2
	// KPDFBorderStyleDashed: Dashed border.
	KPDFBorderStyleDashed PDFBorderStyle = 1
	// KPDFBorderStyleInset: Inset border.
	KPDFBorderStyleInset PDFBorderStyle = 3
	// KPDFBorderStyleSolid: Solid border.
	KPDFBorderStyleSolid PDFBorderStyle = 0
	// KPDFBorderStyleUnderline: Underline border.
	KPDFBorderStyleUnderline PDFBorderStyle = 4
)

func (e PDFBorderStyle) String() string {
	switch e {
	case KPDFBorderStyleBeveled:
		return "KPDFBorderStyleBeveled"
	case KPDFBorderStyleDashed:
		return "KPDFBorderStyleDashed"
	case KPDFBorderStyleInset:
		return "KPDFBorderStyleInset"
	case KPDFBorderStyleSolid:
		return "KPDFBorderStyleSolid"
	case KPDFBorderStyleUnderline:
		return "KPDFBorderStyleUnderline"
	default:
		return fmt.Sprintf("PDFBorderStyle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/PDFKit/PDFDisplayBox
type PDFDisplayBox int

const (
	// KPDFDisplayBoxArtBox: A rectangle defining the boundaries of the page’s meaningful content including surrounding white space intended for display.
	KPDFDisplayBoxArtBox PDFDisplayBox = 4
	// KPDFDisplayBoxBleedBox: A rectangle defining the boundaries of the clip region for the page contents in a production environment.
	KPDFDisplayBoxBleedBox PDFDisplayBox = 2
	// KPDFDisplayBoxCropBox: A rectangle defining the boundaries of the visible region , expressed in default user-space units.
	KPDFDisplayBoxCropBox PDFDisplayBox = 1
	// KPDFDisplayBoxMediaBox: A rectangle defining the boundaries of the physical medium for display or printing, expressed in default user-space units.
	KPDFDisplayBoxMediaBox PDFDisplayBox = 0
	// KPDFDisplayBoxTrimBox: A rectangle defining the intended boundaries of the finished page.
	KPDFDisplayBoxTrimBox PDFDisplayBox = 3
)

func (e PDFDisplayBox) String() string {
	switch e {
	case KPDFDisplayBoxArtBox:
		return "KPDFDisplayBoxArtBox"
	case KPDFDisplayBoxBleedBox:
		return "KPDFDisplayBoxBleedBox"
	case KPDFDisplayBoxCropBox:
		return "KPDFDisplayBoxCropBox"
	case KPDFDisplayBoxMediaBox:
		return "KPDFDisplayBoxMediaBox"
	case KPDFDisplayBoxTrimBox:
		return "KPDFDisplayBoxTrimBox"
	default:
		return fmt.Sprintf("PDFDisplayBox(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/PDFKit/PDFDisplayDirection
type PDFDisplayDirection int

const (
	KPDFDisplayDirectionHorizontal PDFDisplayDirection = 1
	KPDFDisplayDirectionVertical   PDFDisplayDirection = 0
)

func (e PDFDisplayDirection) String() string {
	switch e {
	case KPDFDisplayDirectionHorizontal:
		return "KPDFDisplayDirectionHorizontal"
	case KPDFDisplayDirectionVertical:
		return "KPDFDisplayDirectionVertical"
	default:
		return fmt.Sprintf("PDFDisplayDirection(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/PDFKit/PDFDisplayMode
type PDFDisplayMode int

const (
	// KPDFDisplaySinglePage: A display mode where the document displays one page at a time horizontally and vertically.
	KPDFDisplaySinglePage PDFDisplayMode = 0
	// KPDFDisplaySinglePageContinuous: A display mode where the document displays in continuous mode vertically, with single-page width horizontally.
	KPDFDisplaySinglePageContinuous PDFDisplayMode = 1
	// KPDFDisplayTwoUp: A display mode where the document displays two pages side-by-side.
	KPDFDisplayTwoUp PDFDisplayMode = 2
	// KPDFDisplayTwoUpContinuous: A display mode where the document displays in continuous mode vertically and displays two pages side-by-side horizontally.
	KPDFDisplayTwoUpContinuous PDFDisplayMode = 3
)

func (e PDFDisplayMode) String() string {
	switch e {
	case KPDFDisplaySinglePage:
		return "KPDFDisplaySinglePage"
	case KPDFDisplaySinglePageContinuous:
		return "KPDFDisplaySinglePageContinuous"
	case KPDFDisplayTwoUp:
		return "KPDFDisplayTwoUp"
	case KPDFDisplayTwoUpContinuous:
		return "KPDFDisplayTwoUpContinuous"
	default:
		return fmt.Sprintf("PDFDisplayMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentPermissions
type PDFDocumentPermissions int

const (
	// KPDFDocumentPermissionsNone: The status that indicates no document permissions.
	KPDFDocumentPermissionsNone PDFDocumentPermissions = 0
	// KPDFDocumentPermissionsOwner: The status that indicates owner document permissions.
	KPDFDocumentPermissionsOwner PDFDocumentPermissions = 2
	// KPDFDocumentPermissionsUser: The status that indicates user document permissions.
	KPDFDocumentPermissionsUser PDFDocumentPermissions = 1
)

func (e PDFDocumentPermissions) String() string {
	switch e {
	case KPDFDocumentPermissionsNone:
		return "KPDFDocumentPermissionsNone"
	case KPDFDocumentPermissionsOwner:
		return "KPDFDocumentPermissionsOwner"
	case KPDFDocumentPermissionsUser:
		return "KPDFDocumentPermissionsUser"
	default:
		return fmt.Sprintf("PDFDocumentPermissions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/PDFKit/PDFInterpolationQuality
type PDFInterpolationQuality int

const (
	// KPDFInterpolationQualityHigh: The case specifying high interpolation quality.
	KPDFInterpolationQualityHigh PDFInterpolationQuality = 2
	// KPDFInterpolationQualityLow: The case specifying low interpolation quality.
	KPDFInterpolationQualityLow PDFInterpolationQuality = 1
	// KPDFInterpolationQualityNone: The case where no interpolation quality is specified.
	KPDFInterpolationQualityNone PDFInterpolationQuality = 0
)

func (e PDFInterpolationQuality) String() string {
	switch e {
	case KPDFInterpolationQualityHigh:
		return "KPDFInterpolationQualityHigh"
	case KPDFInterpolationQualityLow:
		return "KPDFInterpolationQualityLow"
	case KPDFInterpolationQualityNone:
		return "KPDFInterpolationQualityNone"
	default:
		return fmt.Sprintf("PDFInterpolationQuality(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/PDFKit/PDFLineStyle
type PDFLineStyle int

const (
	// KPDFLineStyleCircle: A circular line ending filled with the annotation’s interior color, if any.
	KPDFLineStyleCircle PDFLineStyle = 2
	// KPDFLineStyleClosedArrow: A closed arrowhead line ending, consisting of a triangle with the acute vertex at the line end and filled with the annotation’s interior color, if any.
	KPDFLineStyleClosedArrow PDFLineStyle = 5
	// KPDFLineStyleDiamond: A diamond-shaped line ending filled with the annotation’s interior color, if any.
	KPDFLineStyleDiamond PDFLineStyle = 3
	// KPDFLineStyleNone: No line ending.
	KPDFLineStyleNone PDFLineStyle = 0
	// KPDFLineStyleOpenArrow: An open arrowhead line ending, composed from two short lines meeting in an acute angle at the line end.
	KPDFLineStyleOpenArrow PDFLineStyle = 4
	// KPDFLineStyleSquare: A square line ending filled with the annotation’s interior color, if any.
	KPDFLineStyleSquare PDFLineStyle = 1
)

func (e PDFLineStyle) String() string {
	switch e {
	case KPDFLineStyleCircle:
		return "KPDFLineStyleCircle"
	case KPDFLineStyleClosedArrow:
		return "KPDFLineStyleClosedArrow"
	case KPDFLineStyleDiamond:
		return "KPDFLineStyleDiamond"
	case KPDFLineStyleNone:
		return "KPDFLineStyleNone"
	case KPDFLineStyleOpenArrow:
		return "KPDFLineStyleOpenArrow"
	case KPDFLineStyleSquare:
		return "KPDFLineStyleSquare"
	default:
		return fmt.Sprintf("PDFLineStyle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/PDFKit/PDFMarkupType
type PDFMarkupType int

const (
	// KPDFMarkupTypeHighlight: Highlight style for the markup.
	KPDFMarkupTypeHighlight PDFMarkupType = 0
	// KPDFMarkupTypeRedact: The redaction style for markup.
	KPDFMarkupTypeRedact PDFMarkupType = 3
	// KPDFMarkupTypeStrikeOut: Strikethrough style for the markup.
	KPDFMarkupTypeStrikeOut PDFMarkupType = 1
	// KPDFMarkupTypeUnderline: Underline style for the markup.
	KPDFMarkupTypeUnderline PDFMarkupType = 2
)

func (e PDFMarkupType) String() string {
	switch e {
	case KPDFMarkupTypeHighlight:
		return "KPDFMarkupTypeHighlight"
	case KPDFMarkupTypeRedact:
		return "KPDFMarkupTypeRedact"
	case KPDFMarkupTypeStrikeOut:
		return "KPDFMarkupTypeStrikeOut"
	case KPDFMarkupTypeUnderline:
		return "KPDFMarkupTypeUnderline"
	default:
		return fmt.Sprintf("PDFMarkupType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/PDFKit/PDFPrintScalingMode
type PDFPrintScalingMode int

const (
	// KPDFPrintPageScaleDownToFit: Scale large pages down to fit the paper size (smaller pages do not get scaled up).
	KPDFPrintPageScaleDownToFit PDFPrintScalingMode = 2
	// KPDFPrintPageScaleNone: Do not apply scaling to the page when printing.
	KPDFPrintPageScaleNone PDFPrintScalingMode = 0
	// KPDFPrintPageScaleToFit: Scale each page up or down to best fit the paper size.
	KPDFPrintPageScaleToFit PDFPrintScalingMode = 1
)

func (e PDFPrintScalingMode) String() string {
	switch e {
	case KPDFPrintPageScaleDownToFit:
		return "KPDFPrintPageScaleDownToFit"
	case KPDFPrintPageScaleNone:
		return "KPDFPrintPageScaleNone"
	case KPDFPrintPageScaleToFit:
		return "KPDFPrintPageScaleToFit"
	default:
		return fmt.Sprintf("PDFPrintScalingMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/PDFKit/PDFSelectionGranularity
type PDFSelectionGranularity uint

const (
	PDFSelectionGranularityCharacter PDFSelectionGranularity = 0
	PDFSelectionGranularityLine      PDFSelectionGranularity = 2
	PDFSelectionGranularityWord      PDFSelectionGranularity = 1
)

func (e PDFSelectionGranularity) String() string {
	switch e {
	case PDFSelectionGranularityCharacter:
		return "PDFSelectionGranularityCharacter"
	case PDFSelectionGranularityLine:
		return "PDFSelectionGranularityLine"
	case PDFSelectionGranularityWord:
		return "PDFSelectionGranularityWord"
	default:
		return fmt.Sprintf("PDFSelectionGranularity(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/PDFKit/PDFTextAnnotationIconType
type PDFTextAnnotationIconType int

const (
	// KPDFTextAnnotationIconComment: Comment annotation icon.
	KPDFTextAnnotationIconComment PDFTextAnnotationIconType = 0
	// KPDFTextAnnotationIconHelp: Help annotation icon.
	KPDFTextAnnotationIconHelp PDFTextAnnotationIconType = 3
	// KPDFTextAnnotationIconInsert: Insert annotation icon.
	KPDFTextAnnotationIconInsert PDFTextAnnotationIconType = 6
	// KPDFTextAnnotationIconKey: Key annotation icon.
	KPDFTextAnnotationIconKey PDFTextAnnotationIconType = 1
	// KPDFTextAnnotationIconNewParagraph: New Paragraph annotation icon.
	KPDFTextAnnotationIconNewParagraph PDFTextAnnotationIconType = 4
	// KPDFTextAnnotationIconNote: Note annotation icon.
	KPDFTextAnnotationIconNote PDFTextAnnotationIconType = 2
	// KPDFTextAnnotationIconParagraph: Paragraph annotation icon.
	KPDFTextAnnotationIconParagraph PDFTextAnnotationIconType = 5
)

func (e PDFTextAnnotationIconType) String() string {
	switch e {
	case KPDFTextAnnotationIconComment:
		return "KPDFTextAnnotationIconComment"
	case KPDFTextAnnotationIconHelp:
		return "KPDFTextAnnotationIconHelp"
	case KPDFTextAnnotationIconInsert:
		return "KPDFTextAnnotationIconInsert"
	case KPDFTextAnnotationIconKey:
		return "KPDFTextAnnotationIconKey"
	case KPDFTextAnnotationIconNewParagraph:
		return "KPDFTextAnnotationIconNewParagraph"
	case KPDFTextAnnotationIconNote:
		return "KPDFTextAnnotationIconNote"
	case KPDFTextAnnotationIconParagraph:
		return "KPDFTextAnnotationIconParagraph"
	default:
		return fmt.Sprintf("PDFTextAnnotationIconType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/PDFKit/PDFThumbnailLayoutMode
type PDFThumbnailLayoutMode int

const (
	PDFThumbnailLayoutModeHorizontal PDFThumbnailLayoutMode = 0
	PDFThumbnailLayoutModeVertical   PDFThumbnailLayoutMode = 0
)

func (e PDFThumbnailLayoutMode) String() string {
	switch e {
	case PDFThumbnailLayoutModeHorizontal:
		return "PDFThumbnailLayoutModeHorizontal"
	default:
		return fmt.Sprintf("PDFThumbnailLayoutMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/PDFKit/PDFWidgetCellState
type PDFWidgetCellState int

const (
	// KPDFWidgetMixedState: The button widget is in a mixed state, neither on nor off.
	KPDFWidgetMixedState PDFWidgetCellState = -1
	// KPDFWidgetOffState: The button widget is in an unselected state.
	KPDFWidgetOffState PDFWidgetCellState = 0
	// KPDFWidgetOnState: The button widget is in a selected state.
	KPDFWidgetOnState PDFWidgetCellState = 1
)

func (e PDFWidgetCellState) String() string {
	switch e {
	case KPDFWidgetMixedState:
		return "KPDFWidgetMixedState"
	case KPDFWidgetOffState:
		return "KPDFWidgetOffState"
	case KPDFWidgetOnState:
		return "KPDFWidgetOnState"
	default:
		return fmt.Sprintf("PDFWidgetCellState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/PDFKit/PDFWidgetControlType
type PDFWidgetControlType int

const (
	// KPDFWidgetCheckBoxControl: Check box control.
	KPDFWidgetCheckBoxControl PDFWidgetControlType = 2
	// KPDFWidgetPushButtonControl: Push button control.
	KPDFWidgetPushButtonControl PDFWidgetControlType = 0
	// KPDFWidgetRadioButtonControl: Radio button control.
	KPDFWidgetRadioButtonControl PDFWidgetControlType = 1
	// KPDFWidgetUnknownControl: Unknown control type.
	KPDFWidgetUnknownControl PDFWidgetControlType = -1
)

func (e PDFWidgetControlType) String() string {
	switch e {
	case KPDFWidgetCheckBoxControl:
		return "KPDFWidgetCheckBoxControl"
	case KPDFWidgetPushButtonControl:
		return "KPDFWidgetPushButtonControl"
	case KPDFWidgetRadioButtonControl:
		return "KPDFWidgetRadioButtonControl"
	case KPDFWidgetUnknownControl:
		return "KPDFWidgetUnknownControl"
	default:
		return fmt.Sprintf("PDFWidgetControlType(%d)", e)
	}
}
