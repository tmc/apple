// Code generated from Apple documentation. DO NOT EDIT.

package pdfkit

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

var ()

var ()

var ()

var ()

var ()

var ()

var ()

var ()

var (
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentWriteOption/accessPermissionsOption
	PDFDocumentAccessPermissionsOption PDFDocumentWriteOption
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentWriteOption/burnInAnnotationsOption
	PDFDocumentBurnInAnnotationsOption PDFDocumentWriteOption
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentWriteOption/optimizeImagesForScreenOption
	PDFDocumentOptimizeImagesForScreenOption PDFDocumentWriteOption
	// PDFDocumentOwnerPasswordOption is an [NSString] object for the owner’s password which is required for encryption.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentWriteOption/ownerPasswordOption
	PDFDocumentOwnerPasswordOption PDFDocumentWriteOption
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentWriteOption/saveImagesAsJPEGOption
	PDFDocumentSaveImagesAsJPEGOption PDFDocumentWriteOption
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentWriteOption/saveTextFromOCROption
	PDFDocumentSaveTextFromOCROption PDFDocumentWriteOption
	// PDFDocumentUserPasswordOption is an [NSString] object for the user’s password which is optional for encryption.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentWriteOption/userPasswordOption
	PDFDocumentUserPasswordOption PDFDocumentWriteOption
)

var (
	// PDFDocumentAuthorAttribute is an optional text string containing the name of the author of the document.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentAttribute/authorAttribute
	PDFDocumentAuthorAttribute PDFDocumentAttribute
	// PDFDocumentCreationDateAttribute is an optional text string containing the document’s creation date.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentAttribute/creationDateAttribute
	PDFDocumentCreationDateAttribute PDFDocumentAttribute
	// PDFDocumentCreatorAttribute is an optional text string containing the name of the application that created the document content.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentAttribute/creatorAttribute
	PDFDocumentCreatorAttribute PDFDocumentAttribute
	// PDFDocumentKeywordsAttribute is an optional array of text strings containing keywords for the document.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentAttribute/keywordsAttribute
	PDFDocumentKeywordsAttribute PDFDocumentAttribute
	// PDFDocumentModificationDateAttribute is an optional text string containing the document’s last-modified date.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentAttribute/modificationDateAttribute
	PDFDocumentModificationDateAttribute PDFDocumentAttribute
	// PDFDocumentProducerAttribute is an optional text string containing the name of the application that produced the PDF data for the document.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentAttribute/producerAttribute
	PDFDocumentProducerAttribute PDFDocumentAttribute
	// PDFDocumentSubjectAttribute is an optional text string containing a description of the subject of the document.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentAttribute/subjectAttribute
	PDFDocumentSubjectAttribute PDFDocumentAttribute
	// PDFDocumentTitleAttribute is an optional text string containing the title of the document.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentAttribute/titleAttribute
	PDFDocumentTitleAttribute PDFDocumentAttribute
)

var (
	// PDFDocumentDidBeginFindNotification is a notification that the [beginFindString(_:withOptions:)] or [findString(_:withOptions:)] method begins finding.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentDidBeginFindNotification
	PDFDocumentDidBeginFindNotification foundation.NSNotificationName
	// PDFDocumentDidBeginPageFindNotification is a notification that a find operation begins working on a new page of a document.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentDidBeginPageFindNotification
	PDFDocumentDidBeginPageFindNotification foundation.NSNotificationName
	// PDFDocumentDidBeginPageWriteNotification is a notification that a write operation begins working on a page in a document.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentDidBeginPageWriteNotification
	PDFDocumentDidBeginPageWriteNotification foundation.NSNotificationName
	// PDFDocumentDidBeginWriteNotification is a notification that a write operation begins working on a document.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentDidBeginWriteNotification
	PDFDocumentDidBeginWriteNotification foundation.NSNotificationName
	// PDFDocumentDidEndFindNotification is a notification that the [beginFindString(_:withOptions:)] or [findString(_:withOptions:)] method returns.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentDidEndFindNotification
	PDFDocumentDidEndFindNotification foundation.NSNotificationName
	// PDFDocumentDidEndPageFindNotification is a notification that a find operation finishes working on a page in a document.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentDidEndPageFindNotification
	PDFDocumentDidEndPageFindNotification foundation.NSNotificationName
	// PDFDocumentDidEndPageWriteNotification is a notification that a write operation finishes working on a page in a document.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentDidEndPageWriteNotification
	PDFDocumentDidEndPageWriteNotification foundation.NSNotificationName
	// PDFDocumentDidEndWriteNotification is a notification that a write operation finishes working on a document.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentDidEndWriteNotification
	PDFDocumentDidEndWriteNotification foundation.NSNotificationName
	// PDFDocumentDidFindMatchNotification is a notification that a string match is found in a document.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentDidFindMatchNotification
	PDFDocumentDidFindMatchNotification foundation.NSNotificationName
	// PDFDocumentDidUnlockNotification is a notification that a document unlocks after a [unlock(withPassword:)] message.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentDidUnlockNotification
	PDFDocumentDidUnlockNotification foundation.NSNotificationName
	// PDFViewAnnotationHitNotification is a notification posted when the user clicks on an annotation.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFViewAnnotationHitNotification
	PDFViewAnnotationHitNotification foundation.NSNotificationName
	// PDFViewAnnotationWillHitNotification is a notification posted before the user clicks an annotation.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFViewAnnotationWillHitNotification
	PDFViewAnnotationWillHitNotification foundation.NSNotificationName
	// PDFViewChangedHistoryNotification is a notification posted when the page history changes.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFViewChangedHistoryNotification
	PDFViewChangedHistoryNotification foundation.NSNotificationName
	// PDFViewCopyPermissionNotification is a notification posted when the user attempts to copy to the pasteboard without the appropriate permissions.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFViewCopyPermissionNotification
	PDFViewCopyPermissionNotification foundation.NSNotificationName
	// PDFViewDisplayBoxChangedNotification is a notification posted when the display box has changed.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFViewDisplayBoxChangedNotification
	PDFViewDisplayBoxChangedNotification foundation.NSNotificationName
	// PDFViewDisplayModeChangedNotification is a notification posted when the display mode has changed.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFViewDisplayModeChangedNotification
	PDFViewDisplayModeChangedNotification foundation.NSNotificationName
	// PDFViewDocumentChangedNotification is a notification posted when a new document is associated with the view.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFViewDocumentChangedNotification
	PDFViewDocumentChangedNotification foundation.NSNotificationName
	// PDFViewPageChangedNotification is a notification posted when a new page becomes the current page.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFViewPageChangedNotification
	PDFViewPageChangedNotification foundation.NSNotificationName
	// PDFViewPrintPermissionNotification is a notification posted when the user attempts to print without the appropriate permissions.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFViewPrintPermissionNotification
	PDFViewPrintPermissionNotification foundation.NSNotificationName
	// PDFViewScaleChangedNotification is a notification posted when the scale factor changes.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFViewScaleChangedNotification
	PDFViewScaleChangedNotification foundation.NSNotificationName
	// PDFViewSelectionChangedNotification is a notification posted when the current selection has changed.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFViewSelectionChangedNotification
	PDFViewSelectionChangedNotification foundation.NSNotificationName
	// PDFViewVisiblePagesChangedNotification is a notification posted when the visible pages have changed.
	//
	// See: https://developer.apple.com/documentation/PDFKit/PDFViewVisiblePagesChangedNotification
	PDFViewVisiblePagesChangedNotification foundation.NSNotificationName
)

var (
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentFoundSelectionKey
	PDFDocumentFoundSelectionKey string
	// See: https://developer.apple.com/documentation/PDFKit/PDFDocumentPageIndexKey
	PDFDocumentPageIndexKey string
	// See: https://developer.apple.com/documentation/PDFKit/PDFThumbnailViewDocumentEditedNotification
	PDFThumbnailViewDocumentEditedNotification string
)

var ()

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationHighlightingModeInvert"); err == nil && ptr != 0 {
		PDFAnnotationHighlightingModes.Invert = *(*PDFAnnotationHighlightingMode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationHighlightingModeNone"); err == nil && ptr != 0 {
		PDFAnnotationHighlightingModes.None = *(*PDFAnnotationHighlightingMode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationHighlightingModeOutline"); err == nil && ptr != 0 {
		PDFAnnotationHighlightingModes.Outline = *(*PDFAnnotationHighlightingMode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationHighlightingModePush"); err == nil && ptr != 0 {
		PDFAnnotationHighlightingModes.Push = *(*PDFAnnotationHighlightingMode)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyAction"); err == nil && ptr != 0 {
		PDFAnnotationKeys.Action = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyAdditionalActions"); err == nil && ptr != 0 {
		PDFAnnotationKeys.AdditionalActions = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyAppearanceDictionary"); err == nil && ptr != 0 {
		PDFAnnotationKeys.AppearanceDictionary = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyAppearanceState"); err == nil && ptr != 0 {
		PDFAnnotationKeys.AppearanceState = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyBorder"); err == nil && ptr != 0 {
		PDFAnnotationKeys.Border = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyBorderStyle"); err == nil && ptr != 0 {
		PDFAnnotationKeys.BorderStyle = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyColor"); err == nil && ptr != 0 {
		PDFAnnotationKeys.Color = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyContents"); err == nil && ptr != 0 {
		PDFAnnotationKeys.Contents = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyDate"); err == nil && ptr != 0 {
		PDFAnnotationKeys.Date = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyDefaultAppearance"); err == nil && ptr != 0 {
		PDFAnnotationKeys.DefaultAppearance = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyDestination"); err == nil && ptr != 0 {
		PDFAnnotationKeys.Destination = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyFlags"); err == nil && ptr != 0 {
		PDFAnnotationKeys.Flags = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyHighlightingMode"); err == nil && ptr != 0 {
		PDFAnnotationKeys.HighlightingMode = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyIconName"); err == nil && ptr != 0 {
		PDFAnnotationKeys.IconName = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyInklist"); err == nil && ptr != 0 {
		PDFAnnotationKeys.Inklist = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyInteriorColor"); err == nil && ptr != 0 {
		PDFAnnotationKeys.InteriorColor = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyLineEndingStyles"); err == nil && ptr != 0 {
		PDFAnnotationKeys.LineEndingStyles = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyLinePoints"); err == nil && ptr != 0 {
		PDFAnnotationKeys.LinePoints = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyName"); err == nil && ptr != 0 {
		PDFAnnotationKeys.Name = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyOpen"); err == nil && ptr != 0 {
		PDFAnnotationKeys.Open = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyPage"); err == nil && ptr != 0 {
		PDFAnnotationKeys.Page = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyParent"); err == nil && ptr != 0 {
		PDFAnnotationKeys.Parent = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyPopup"); err == nil && ptr != 0 {
		PDFAnnotationKeys.Popup = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyQuadPoints"); err == nil && ptr != 0 {
		PDFAnnotationKeys.QuadPoints = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyQuadding"); err == nil && ptr != 0 {
		PDFAnnotationKeys.Quadding = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyRect"); err == nil && ptr != 0 {
		PDFAnnotationKeys.Rect = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeySubtype"); err == nil && ptr != 0 {
		PDFAnnotationKeys.Subtype = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyTextLabel"); err == nil && ptr != 0 {
		PDFAnnotationKeys.TextLabel = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyWidgetAppearanceDictionary"); err == nil && ptr != 0 {
		PDFAnnotationKeys.WidgetAppearanceDictionary = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyWidgetBackgroundColor"); err == nil && ptr != 0 {
		PDFAnnotationKeys.WidgetBackgroundColor = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyWidgetBorderColor"); err == nil && ptr != 0 {
		PDFAnnotationKeys.WidgetBorderColor = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyWidgetCaption"); err == nil && ptr != 0 {
		PDFAnnotationKeys.WidgetCaption = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyWidgetDefaultValue"); err == nil && ptr != 0 {
		PDFAnnotationKeys.WidgetDefaultValue = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyWidgetDownCaption"); err == nil && ptr != 0 {
		PDFAnnotationKeys.WidgetDownCaption = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyWidgetFieldFlags"); err == nil && ptr != 0 {
		PDFAnnotationKeys.WidgetFieldFlags = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyWidgetFieldType"); err == nil && ptr != 0 {
		PDFAnnotationKeys.WidgetFieldType = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyWidgetMaxLen"); err == nil && ptr != 0 {
		PDFAnnotationKeys.WidgetMaxLen = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyWidgetOptions"); err == nil && ptr != 0 {
		PDFAnnotationKeys.WidgetOptions = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyWidgetRolloverCaption"); err == nil && ptr != 0 {
		PDFAnnotationKeys.WidgetRolloverCaption = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyWidgetRotation"); err == nil && ptr != 0 {
		PDFAnnotationKeys.WidgetRotation = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyWidgetTextLabelUI"); err == nil && ptr != 0 {
		PDFAnnotationKeys.WidgetTextLabelUI = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationKeyWidgetValue"); err == nil && ptr != 0 {
		PDFAnnotationKeys.WidgetValue = *(*PDFAnnotationKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationLineEndingStyleCircle"); err == nil && ptr != 0 {
		PDFAnnotationLineEndingStyles.Circle = *(*PDFAnnotationLineEndingStyle)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationLineEndingStyleClosedArrow"); err == nil && ptr != 0 {
		PDFAnnotationLineEndingStyles.ClosedArrow = *(*PDFAnnotationLineEndingStyle)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationLineEndingStyleDiamond"); err == nil && ptr != 0 {
		PDFAnnotationLineEndingStyles.Diamond = *(*PDFAnnotationLineEndingStyle)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationLineEndingStyleNone"); err == nil && ptr != 0 {
		PDFAnnotationLineEndingStyles.None = *(*PDFAnnotationLineEndingStyle)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationLineEndingStyleOpenArrow"); err == nil && ptr != 0 {
		PDFAnnotationLineEndingStyles.OpenArrow = *(*PDFAnnotationLineEndingStyle)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationLineEndingStyleSquare"); err == nil && ptr != 0 {
		PDFAnnotationLineEndingStyles.Square = *(*PDFAnnotationLineEndingStyle)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationSubtypeCircle"); err == nil && ptr != 0 {
		PDFAnnotationSubtypes.Circle = *(*PDFAnnotationSubtype)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationSubtypeFreeText"); err == nil && ptr != 0 {
		PDFAnnotationSubtypes.FreeText = *(*PDFAnnotationSubtype)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationSubtypeHighlight"); err == nil && ptr != 0 {
		PDFAnnotationSubtypes.Highlight = *(*PDFAnnotationSubtype)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationSubtypeInk"); err == nil && ptr != 0 {
		PDFAnnotationSubtypes.Ink = *(*PDFAnnotationSubtype)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationSubtypeLine"); err == nil && ptr != 0 {
		PDFAnnotationSubtypes.Line = *(*PDFAnnotationSubtype)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationSubtypeLink"); err == nil && ptr != 0 {
		PDFAnnotationSubtypes.Link = *(*PDFAnnotationSubtype)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationSubtypePopup"); err == nil && ptr != 0 {
		PDFAnnotationSubtypes.Popup = *(*PDFAnnotationSubtype)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationSubtypeSquare"); err == nil && ptr != 0 {
		PDFAnnotationSubtypes.Square = *(*PDFAnnotationSubtype)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationSubtypeStamp"); err == nil && ptr != 0 {
		PDFAnnotationSubtypes.Stamp = *(*PDFAnnotationSubtype)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationSubtypeStrikeOut"); err == nil && ptr != 0 {
		PDFAnnotationSubtypes.StrikeOut = *(*PDFAnnotationSubtype)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationSubtypeText"); err == nil && ptr != 0 {
		PDFAnnotationSubtypes.Text = *(*PDFAnnotationSubtype)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationSubtypeUnderline"); err == nil && ptr != 0 {
		PDFAnnotationSubtypes.Underline = *(*PDFAnnotationSubtype)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationSubtypeWidget"); err == nil && ptr != 0 {
		PDFAnnotationSubtypes.Widget = *(*PDFAnnotationSubtype)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationTextIconTypeComment"); err == nil && ptr != 0 {
		PDFAnnotationTextIconTypes.Comment = *(*PDFAnnotationTextIconType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationTextIconTypeHelp"); err == nil && ptr != 0 {
		PDFAnnotationTextIconTypes.Help = *(*PDFAnnotationTextIconType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationTextIconTypeInsert"); err == nil && ptr != 0 {
		PDFAnnotationTextIconTypes.Insert = *(*PDFAnnotationTextIconType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationTextIconTypeKey"); err == nil && ptr != 0 {
		PDFAnnotationTextIconTypes.Key = *(*PDFAnnotationTextIconType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationTextIconTypeNewParagraph"); err == nil && ptr != 0 {
		PDFAnnotationTextIconTypes.NewParagraph = *(*PDFAnnotationTextIconType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationTextIconTypeNote"); err == nil && ptr != 0 {
		PDFAnnotationTextIconTypes.Note = *(*PDFAnnotationTextIconType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationTextIconTypeParagraph"); err == nil && ptr != 0 {
		PDFAnnotationTextIconTypes.Paragraph = *(*PDFAnnotationTextIconType)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationWidgetSubtypeButton"); err == nil && ptr != 0 {
		PDFAnnotationWidgetSubtypes.Button = *(*PDFAnnotationWidgetSubtype)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationWidgetSubtypeChoice"); err == nil && ptr != 0 {
		PDFAnnotationWidgetSubtypes.Choice = *(*PDFAnnotationWidgetSubtype)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationWidgetSubtypeSignature"); err == nil && ptr != 0 {
		PDFAnnotationWidgetSubtypes.Signature = *(*PDFAnnotationWidgetSubtype)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAnnotationWidgetSubtypeText"); err == nil && ptr != 0 {
		PDFAnnotationWidgetSubtypes.Text = *(*PDFAnnotationWidgetSubtype)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAppearanceCharacteristicsKeyBackgroundColor"); err == nil && ptr != 0 {
		PDFAppearanceCharacteristicsKeys.BackgroundColor = *(*PDFAppearanceCharacteristicsKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAppearanceCharacteristicsKeyBorderColor"); err == nil && ptr != 0 {
		PDFAppearanceCharacteristicsKeys.BorderColor = *(*PDFAppearanceCharacteristicsKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAppearanceCharacteristicsKeyCaption"); err == nil && ptr != 0 {
		PDFAppearanceCharacteristicsKeys.Caption = *(*PDFAppearanceCharacteristicsKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAppearanceCharacteristicsKeyDownCaption"); err == nil && ptr != 0 {
		PDFAppearanceCharacteristicsKeys.DownCaption = *(*PDFAppearanceCharacteristicsKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAppearanceCharacteristicsKeyRolloverCaption"); err == nil && ptr != 0 {
		PDFAppearanceCharacteristicsKeys.RolloverCaption = *(*PDFAppearanceCharacteristicsKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFAppearanceCharacteristicsKeyRotation"); err == nil && ptr != 0 {
		PDFAppearanceCharacteristicsKeys.Rotation = *(*PDFAppearanceCharacteristicsKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFBorderKeyDashPattern"); err == nil && ptr != 0 {
		PDFBorderKeys.DashPattern = *(*PDFBorderKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFBorderKeyLineWidth"); err == nil && ptr != 0 {
		PDFBorderKeys.LineWidth = *(*PDFBorderKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFBorderKeyStyle"); err == nil && ptr != 0 {
		PDFBorderKeys.Style = *(*PDFBorderKey)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentAccessPermissionsOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentAccessPermissionsOption = PDFDocumentWriteOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentAuthorAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentAuthorAttribute = PDFDocumentAttribute(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentBurnInAnnotationsOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentBurnInAnnotationsOption = PDFDocumentWriteOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentCreationDateAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentCreationDateAttribute = PDFDocumentAttribute(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentCreatorAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentCreatorAttribute = PDFDocumentAttribute(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentDidBeginFindNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentDidBeginFindNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentDidBeginPageFindNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentDidBeginPageFindNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentDidBeginPageWriteNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentDidBeginPageWriteNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentDidBeginWriteNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentDidBeginWriteNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentDidEndFindNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentDidEndFindNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentDidEndPageFindNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentDidEndPageFindNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentDidEndPageWriteNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentDidEndPageWriteNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentDidEndWriteNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentDidEndWriteNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentDidFindMatchNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentDidFindMatchNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentDidUnlockNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentDidUnlockNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentFoundSelectionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentFoundSelectionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentKeywordsAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentKeywordsAttribute = PDFDocumentAttribute(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentModificationDateAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentModificationDateAttribute = PDFDocumentAttribute(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentOptimizeImagesForScreenOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentOptimizeImagesForScreenOption = PDFDocumentWriteOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentOwnerPasswordOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentOwnerPasswordOption = PDFDocumentWriteOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentPageIndexKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentPageIndexKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentProducerAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentProducerAttribute = PDFDocumentAttribute(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentSaveImagesAsJPEGOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentSaveImagesAsJPEGOption = PDFDocumentWriteOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentSaveTextFromOCROption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentSaveTextFromOCROption = PDFDocumentWriteOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentSubjectAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentSubjectAttribute = PDFDocumentAttribute(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentTitleAttribute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentTitleAttribute = PDFDocumentAttribute(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFDocumentUserPasswordOption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFDocumentUserPasswordOption = PDFDocumentWriteOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFPageImageInitializationOptionCompressionQuality"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFPageImageInitializationOptions.CompressionQuality = PDFPageImageInitializationOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFPageImageInitializationOptionMediaBox"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFPageImageInitializationOptions.MediaBox = PDFPageImageInitializationOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFPageImageInitializationOptionRotation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFPageImageInitializationOptions.Rotation = PDFPageImageInitializationOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFPageImageInitializationOptionUpscaleIfSmaller"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFPageImageInitializationOptions.UpscaleIfSmaller = PDFPageImageInitializationOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFThumbnailViewDocumentEditedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFThumbnailViewDocumentEditedNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFViewAnnotationHitNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFViewAnnotationHitNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFViewAnnotationWillHitNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFViewAnnotationWillHitNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFViewChangedHistoryNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFViewChangedHistoryNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFViewCopyPermissionNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFViewCopyPermissionNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFViewDisplayBoxChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFViewDisplayBoxChangedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFViewDisplayModeChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFViewDisplayModeChangedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFViewDocumentChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFViewDocumentChangedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFViewPageChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFViewPageChangedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFViewPrintPermissionNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFViewPrintPermissionNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFViewScaleChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFViewScaleChangedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFViewSelectionChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFViewSelectionChangedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "PDFViewVisiblePagesChangedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				PDFViewVisiblePagesChangedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

}

// PDFAnnotationHighlightingModes provides typed accessors for [PDFAnnotationHighlightingMode] constants.
var PDFAnnotationHighlightingModes struct {
	// Invert: A highlight mode that inverts the content of the annotation.
	Invert PDFAnnotationHighlightingMode
	// None: A highlight mode that doesn’t change the appearance of the annotation.
	None PDFAnnotationHighlightingMode
	// Outline: A highlight mode that inverts the annotation’s border.
	Outline PDFAnnotationHighlightingMode
	// Push: A highlight mode that renders a pressed appearance for the annotation.
	Push PDFAnnotationHighlightingMode
}

// PDFAnnotationKeys provides typed accessors for [PDFAnnotationKey] constants.
var PDFAnnotationKeys struct {
	// Action: A dictionary or PDF action object that represents an action to take, such as when the user clicks or taps a button.
	Action PDFAnnotationKey
	// AdditionalActions: A dictionary or PDF action object that represents additional actions an annotation can perform, such as when it receives input focus.
	AdditionalActions PDFAnnotationKey
	// AppearanceDictionary: A dictionary that contains properties for controlling the annotation’s visual appearance.
	AppearanceDictionary PDFAnnotationKey
	// AppearanceState: A string that specifies the appearance stream for the annotation.
	AppearanceState PDFAnnotationKey
	// Border: An array of integers or border objects that describes the border of the annotation.
	Border PDFAnnotationKey
	// BorderStyle: A dictionary that contains the properties of the annotation’s border.
	BorderStyle PDFAnnotationKey
	// Color: An array of floats or a color object that specifies the annotation’s color.
	Color PDFAnnotationKey
	// Contents: The text that the annotation displays or represents.
	Contents PDFAnnotationKey
	// Date: The date, or string representation of a date, of the annotation’s most recent modification.
	Date PDFAnnotationKey
	// DefaultAppearance: A string value a free text annotation uses to format the text.
	DefaultAppearance PDFAnnotationKey
	// Destination: An array, name, or string that represents the destination of an action.
	Destination PDFAnnotationKey
	// Flags: An integer value that specifies flags for the annotation.
	Flags PDFAnnotationKey
	// HighlightingMode: A string value that defines the way an annotation highlights when the user activates it, such as when clicking or tapping a link.
	HighlightingMode PDFAnnotationKey
	// IconName: A string value that specifies the name of an icon for a text or stamp annotation.
	IconName PDFAnnotationKey
	// Inklist: An array of arrays that represents stroked paths.
	Inklist PDFAnnotationKey
	// InteriorColor: An array of floating point values or a PDF color object that annotations use to fill interior space, such as line endings, squares, or circles.
	InteriorColor PDFAnnotationKey
	// LineEndingStyles: An array of string values that specifies the styles to use for the ends of lines.
	LineEndingStyles PDFAnnotationKey
	// LinePoints: An array of floating point values that specifies the starting and ending points, in page-space coordinates, of a line.
	LinePoints PDFAnnotationKey
	// Name: A string that uniquely identifies the annotation among all annotations on the same page.
	Name PDFAnnotationKey
	// Open: A Boolean value that specifies whether the pop-up is in an opened state, showing its text content, or in a closed state and showing an icon.
	Open PDFAnnotationKey
	// Page: A dictionary or PDF page object that includes the annotation.
	Page PDFAnnotationKey
	// Parent: A dictionary or annotation object that a pop-up or widget belongs to.
	Parent PDFAnnotationKey
	// Popup: A dictionary or annotation object that specifies the annotation to pop up for text entry or editing.
	Popup PDFAnnotationKey
	// QuadPoints: An array of floating point values that specifies a rectangular region of a page.
	QuadPoints PDFAnnotationKey
	// Quadding: An integer value that specifies left, right, or center justification.
	Quadding PDFAnnotationKey
	// Rect: The rectangle that the annotation occupies on the page, in page-space coordinates.
	Rect PDFAnnotationKey
	// Subtype: The type of annotation that the entries in a dictionary describe.
	Subtype PDFAnnotationKey
	// TextLabel: A string that represents the title of the annotation.
	TextLabel PDFAnnotationKey
	// WidgetAppearanceDictionary: A dictionary or appearance characteristic object that contains properties for controlling the widget’s visual appearance.
	WidgetAppearanceDictionary PDFAnnotationKey
	// WidgetBackgroundColor: An array of floating point values or a PDF color object that specifies the widget’s background color.
	WidgetBackgroundColor PDFAnnotationKey
	// WidgetBorderColor: An array of floating point values or a PDF color object that specifies the widget’s border color.
	WidgetBorderColor PDFAnnotationKey
	// WidgetCaption: A string that a push button widget displays when it isn’t in a pressed state.
	WidgetCaption PDFAnnotationKey
	// WidgetDefaultValue: A default value for the widget.
	WidgetDefaultValue PDFAnnotationKey
	// WidgetDownCaption: A string that a push button widgets displays when it’s in a pressed state.
	WidgetDownCaption PDFAnnotationKey
	// WidgetFieldFlags: An integer value that specifies flags for a widget.
	WidgetFieldFlags PDFAnnotationKey
	// WidgetFieldType: A string that specifies the type of widget, such as button, checkbox, or signature field.
	WidgetFieldType PDFAnnotationKey
	// WidgetMaxLen: An integer value that specifies the maximum length of a text field, in characters.
	WidgetMaxLen PDFAnnotationKey
	// WidgetOptions: An array that specifies the options to present in radio buttons or choice lists.
	WidgetOptions PDFAnnotationKey
	// WidgetRolloverCaption: A string that push button widgets display when the pointer is over the button, but not clicking it.
	WidgetRolloverCaption PDFAnnotationKey
	// WidgetRotation: An integer value that specifies the rotation of the widget.
	WidgetRotation PDFAnnotationKey
	// WidgetTextLabelUI: A user-visible alternative field name that identifies the widget, typically for accessibility purposes.
	WidgetTextLabelUI PDFAnnotationKey
	// WidgetValue: The widget’s value, typically for text and choice widgets.
	WidgetValue PDFAnnotationKey
}

// PDFAnnotationLineEndingStyles provides typed accessors for [PDFAnnotationLineEndingStyle] constants.
var PDFAnnotationLineEndingStyles struct {
	// Circle: A style that displays a circle line ending and fills it with the annotation’s interior color.
	Circle PDFAnnotationLineEndingStyle
	// ClosedArrow: A style that displays a closed arrowhead line ending and fills it with the annotation’s interior color.
	ClosedArrow PDFAnnotationLineEndingStyle
	Diamond     PDFAnnotationLineEndingStyle
	None        PDFAnnotationLineEndingStyle
	// OpenArrow: A style that displays an open arrowhead line ending.
	OpenArrow PDFAnnotationLineEndingStyle
	// Square: A style that displays a square line ending and fills it with the annotation’s interior color.
	Square PDFAnnotationLineEndingStyle
}

// PDFAnnotationSubtypes provides typed accessors for [PDFAnnotationSubtype] constants.
var PDFAnnotationSubtypes struct {
	// Circle: An annotation that renders a circle shape.
	Circle PDFAnnotationSubtype
	// FreeText: An annotation that displays an editable text field.
	FreeText PDFAnnotationSubtype
	// Highlight: An annotation that highlights text.
	Highlight PDFAnnotationSubtype
	// Ink: An annotation that represents a freehand scribble.
	Ink PDFAnnotationSubtype
	// Line: An annotation that displays a single straight line.
	Line PDFAnnotationSubtype
	// Link: An annotation that provides a hyperlink to a location in the document, or an action to perform when the user clicks or taps it.
	Link PDFAnnotationSubtype
	// Popup: An annotation that displays text in a pop-up window for entry and editing.
	Popup PDFAnnotationSubtype
	// Square: An annotation that displays a square shape.
	Square PDFAnnotationSubtype
	// Stamp: An annotation that displays text or a graphic as if a rubber stamp imprints it on the page.
	Stamp PDFAnnotationSubtype
	// StrikeOut: An annotation that strikes out text.
	StrikeOut PDFAnnotationSubtype
	// Text: An annotation that displays a collapsible note that contains text.
	Text PDFAnnotationSubtype
	// Underline: An annotation that underlines text.
	Underline PDFAnnotationSubtype
	// Widget: An annotation that displays interactive form elements, including text or signature fields, radio buttons, checkboxes, push buttons, pop-ups, and tables.
	Widget PDFAnnotationSubtype
}

// PDFAnnotationTextIconTypes provides typed accessors for [PDFAnnotationTextIconType] constants.
var PDFAnnotationTextIconTypes struct {
	// Comment: An icon that represents a comment.
	Comment PDFAnnotationTextIconType
	// Help: An icon that indicates help information is available.
	Help PDFAnnotationTextIconType
	// Insert: An icon that represents an insertion command.
	Insert PDFAnnotationTextIconType
	// Key: An icon that represents a key.
	Key PDFAnnotationTextIconType
	// NewParagraph: An icon that represents a new paragraph.
	NewParagraph PDFAnnotationTextIconType
	// Note: An icon that represents a note.
	Note PDFAnnotationTextIconType
	// Paragraph: An icon that represents a paragraph.
	Paragraph PDFAnnotationTextIconType
}

// PDFAnnotationWidgetSubtypes provides typed accessors for [PDFAnnotationWidgetSubtype] constants.
var PDFAnnotationWidgetSubtypes struct {
	// Button: A button widget type, including push buttons, checkboxes, and radio buttons.
	Button PDFAnnotationWidgetSubtype
	// Choice: A type that presents a list of choices the user can choose from.
	Choice PDFAnnotationWidgetSubtype
	// Signature: A digital signature widget type.
	Signature PDFAnnotationWidgetSubtype
	// Text: A text field the user can type text in.
	Text PDFAnnotationWidgetSubtype
}

// PDFAppearanceCharacteristicsKeys provides typed accessors for [PDFAppearanceCharacteristicsKey] constants.
var PDFAppearanceCharacteristicsKeys struct {
	// BackgroundColor: The background color of the widget annotation.
	BackgroundColor PDFAppearanceCharacteristicsKey
	// BorderColor: The border color of the widget annotation.
	BorderColor PDFAppearanceCharacteristicsKey
	// Caption: The text that the button widget annotation displays when the user isn’t interacting with it.
	Caption PDFAppearanceCharacteristicsKey
	// DownCaption: The text that the button widget annotation displays when the user holds down on it.
	DownCaption PDFAppearanceCharacteristicsKey
	// RolloverCaption: The text that the button widget annotation displays when the user hovers the pointer over it.
	RolloverCaption PDFAppearanceCharacteristicsKey
	// Rotation: The number of degrees, in multiples of 90, that the widget annotation rotates counterclockwise relative to the page.
	Rotation PDFAppearanceCharacteristicsKey
}

// PDFBorderKeys provides typed accessors for [PDFBorderKey] constants.
var PDFBorderKeys struct {
	DashPattern PDFBorderKey
	LineWidth   PDFBorderKey
	Style       PDFBorderKey
}

// PDFPageImageInitializationOptions provides typed accessors for [PDFPageImageInitializationOption] constants.
var PDFPageImageInitializationOptions struct {
	CompressionQuality PDFPageImageInitializationOption
	MediaBox           PDFPageImageInitializationOption
	Rotation           PDFPageImageInitializationOption
	UpscaleIfSmaller   PDFPageImageInitializationOption
}
