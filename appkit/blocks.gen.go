// Code generated from Apple documentation. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// AnimationContextHandler handles A block object containing animations for this transaction group.
//
// Used by:
//   - [NSAnimationContext.RunAnimationGroupCompletionHandler]
//   - [NSAnimationContext.RunAnimationGroup]
type AnimationContextHandler = func(*NSAnimationContext)

// NewAnimationContextBlock wraps a Go [AnimationContextHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSAnimationContext.RunAnimationGroupCompletionHandler]
//   - [NSAnimationContext.RunAnimationGroup]
func NewAnimationContextBlock(handler AnimationContextHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSAnimationContext
		if resultID != 0 {
			v := NSAnimationContextFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}



// AppearanceHandler is the signature for a completion handler block.
//
// Used by:
//   - [NSColor.ColorWithNameDynamicProvider]
type AppearanceHandler = func(*NSAppearance)

// NewAppearanceBlock wraps a Go [AppearanceHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSColor.ColorWithNameDynamicProvider]
func NewAppearanceBlock(handler AppearanceHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSAppearance
		if resultID != 0 {
			v := NSAppearanceFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}



// ArrayHandler handles The completion handler that is called when the user clicks the OK or Cancel button in the open panel.
//   - items: The items to add to the results array. The `handleMatchedItems` block can be invoked from any thread desired.  If it is called more than once the additional results will be appended after previous items until the maximum is reached.
//
// Used by:
//   - [NSDocumentController.BeginOpenPanelWithCompletionHandler]
//   - [NSSpellChecker.RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler]
//   - [NSUserInterfaceItemSearching.SearchForItemsWithSearchStringResultLimitMatchedItemHandler]
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion]
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion]
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion]
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion]
type ArrayHandler = func(*[]foundation.NSTextCheckingResult)

// NewArrayBlock wraps a Go [ArrayHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSDocumentController.BeginOpenPanelWithCompletionHandler]
//   - [NSSpellChecker.RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler]
//   - [NSUserInterfaceItemSearching.SearchForItemsWithSearchStringResultLimitMatchedItemHandler]
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion]
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion]
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion]
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion]
func NewArrayBlock(handler ArrayHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *[]foundation.NSTextCheckingResult
		if resultID != 0 {
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]foundation.NSTextCheckingResult, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = foundation.NSTextCheckingResultFromID(item.GetID())
			}
			result = &res
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}



// AttributedStringHandler handles A completion handler to execute with the results of the operation.
//
// Used by:
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion]
type AttributedStringHandler = func(*foundation.NSAttributedString)

// NewAttributedStringBlock wraps a Go [AttributedStringHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorReplaceRangeInContextProposedTextReasonAnimationParametersCompletion]
func NewAttributedStringBlock(handler AttributedStringHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *foundation.NSAttributedString
		if resultID != 0 {
			v := foundation.NSAttributedStringFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}



// CGRectTextContainerHandler handles A closure you provide to determine if the enumeration finishes early.
//
// Used by:
//   - [NSTextLayoutManager.EnumerateTextSegmentsInRangeTypeOptionsUsingBlock]
type CGRectTextContainerHandler = func(*NSTextRange, *NSTextContainer)

// NewCGRectTextContainerBlock wraps a Go [CGRectTextContainerHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSTextLayoutManager.EnumerateTextSegmentsInRangeTypeOptionsUsingBlock]
func NewCGRectTextContainerBlock(handler CGRectTextContainerHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID) {
		var result *NSTextRange
		if resultID != 0 {
			v := NSTextRangeFromID(resultID)
			result = &v
		}
		var extra0 *NSTextContainer
		if extra0ID != 0 {
			v := NSTextContainerFromID(extra0ID)
			extra0 = &v
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}



// ColorHandler handles The handler block for processing the user-selected color.
//   - selectedColor: The selected color.
//
// Used by:
//   - [NSColorSampler.ShowSamplerWithSelectionHandler]
type ColorHandler = func(*NSColor)

// NewColorBlock wraps a Go [ColorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSColorSampler.ShowSamplerWithSelectionHandler]
func NewColorBlock(handler ColorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSColor
		if resultID != 0 {
			v := NSColorFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}



// DictionaryErrorHandler handles The completion handler block object to call when the operation completes.
//   - newURLs: A dictionary parameter whose keys and values are [NSURL](<doc://com.apple.documentation/documentation/Foundation/NSURL>) objects. Each key is a URL from the [URLs] parameter. The value of each key is a URL representing the location of the duplicated file. If this method could not duplicate a file, the corresponding URL is not included in the dictionary.
//   - error: If the operation succeeded for every file, this parameter is `nil`. If the operation failed for one or more files, the parameter contains an error object describing the overall result of the operation in a manner suitable for presentation to the user.
//
// Used by:
//   - [NSWorkspace.DuplicateURLsCompletionHandler]
//   - [NSWorkspace.RecycleURLsCompletionHandler]
type DictionaryErrorHandler = func(*foundation.INSDictionary, error)



// DocumentErrorHandler handles The completion handler block object passed in to be called at some point in the future, perhaps after the method invocation has returned.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSDocumentController.OpenDocumentWithContentsOfURLDisplayCompletionHandler]
//   - [NSDocumentController.ReopenDocumentForURLWithContentsOfURLDisplayCompletionHandler]
type DocumentErrorHandler = func(*NSDocument, error)

// NewDocumentErrorBlock wraps a Go [DocumentErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSDocumentController.OpenDocumentWithContentsOfURLDisplayCompletionHandler]
//   - [NSDocumentController.ReopenDocumentForURLWithContentsOfURLDisplayCompletionHandler]
func NewDocumentErrorBlock(handler DocumentErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *NSDocument
		if resultID != 0 {
			v := NSDocumentFromID(resultID)
			result = &v
		}
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}



// ErrorHandler handles The completion handler block object passed in to be invoked at some point in the future, perhaps after the method invocation has returned.
//   - success: A Boolean value indicating whether the sharing attempt was successful. The value is [true](<doc://com.apple.documentation/documentation/Swift/true>) if sharing succeeded, or [false](<doc://com.apple.documentation/documentation/Swift/false>) if it didn’t.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSAccessibilityCustomAction.InitWithNameHandler]
//   - [NSAlert.BeginSheetModalForWindowCompletionHandler]
//   - [NSApplication.RegisterUserInterfaceItemSearchHandler]
//   - [NSApplication.UnregisterUserInterfaceItemSearchHandler]
//   - [NSCollectionView.PerformBatchUpdatesCompletionHandler]
//   - [NSCustomImageRep.InitWithSizeFlippedDrawingHandler]
//   - [NSDocument.AccommodatePresentedItemDeletionWithCompletionHandler]
//   - [NSDocument.AutosaveWithImplicitCancellabilityCompletionHandler]
//   - [NSDocument.LockDocumentWithCompletionHandler]
//   - [NSDocument.LockWithCompletionHandler]
//   - [NSDocument.MoveDocumentWithCompletionHandler]
//   - [NSDocument.MoveToURLCompletionHandler]
//   - [NSDocument.SavePresentedItemChangesWithCompletionHandler]
//   - [NSDocument.SaveToURLOfTypeForSaveOperationCompletionHandler]
//   - [NSDocument.ShareDocumentWithSharingServiceCompletionHandler]
//   - [NSDocument.UnlockDocumentWithCompletionHandler]
//   - [NSDocument.UnlockWithCompletionHandler]
//   - [NSDocumentController.BeginOpenPanelForTypesCompletionHandler]
//   - [NSEvent.TrackSwipeEventWithOptionsDampenAmountThresholdMinMaxUsingHandler]
//   - [NSFilePromiseProviderDelegate.FilePromiseProviderWritePromiseToURLCompletionHandler]
//   - [NSFontAssetRequest.DownloadFontAssetsWithCompletionHandler]
//   - [NSImage.ImageWithSizeFlippedDrawingHandler]
//   - [NSPDFPanel.BeginSheetWithPDFInfoModalForWindowCompletionHandler]
//   - [NSPageLayout.BeginSheetUsingPrintInfoOnWindowCompletionHandler]
//   - [NSPrintPanel.BeginSheetUsingPrintInfoOnWindowCompletionHandler]
//   - [NSSavePanel.BeginSheetModalForWindowCompletionHandler]
//   - [NSSavePanel.BeginWithCompletionHandler]
//   - [NSSpellChecker.RequestCheckingOfStringRangeTypesOptionsInSpellDocumentWithTagCompletionHandler]
//   - [NSTableView.EnumerateAvailableRowViewsUsingBlock]
//   - [NSTableViewRowAction.RowActionWithStyleTitleHandler]
//   - [NSTextContentManager.SynchronizeTextLayoutManagers]
//   - [NSTextContentManager.SynchronizeToBackingStore]
//   - [NSTextContentStorage.SynchronizeToBackingStore]
//   - [NSTextElementProvider.SynchronizeToBackingStore]
//   - [NSWindow.BeginCriticalSheetCompletionHandler]
//   - [NSWindow.BeginSheetCompletionHandler]
//   - [NSWindow.RequestSharingOfWindowCompletionHandler]
//   - [NSWindow.RequestSharingOfWindowUsingPreviewTitleCompletionHandler]
//   - [NSWindow.TrackEventsMatchingMaskTimeoutModeHandler]
//   - [NSWindow.TransferWindowSharingToWindowCompletionHandler]
//   - [NSWorkspace.SetDefaultApplicationAtURLToOpenContentTypeCompletionHandler]
//   - [NSWorkspace.SetDefaultApplicationAtURLToOpenContentTypeOfFileAtURLCompletionHandler]
//   - [NSWorkspace.SetDefaultApplicationAtURLToOpenFileAtURLCompletionHandler]
//   - [NSWorkspace.SetDefaultApplicationAtURLToOpenURLsWithSchemeCompletionHandler]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSAccessibilityCustomAction.InitWithNameHandler]
//   - [NSAlert.BeginSheetModalForWindowCompletionHandler]
//   - [NSApplication.RegisterUserInterfaceItemSearchHandler]
//   - [NSApplication.UnregisterUserInterfaceItemSearchHandler]
//   - [NSCollectionView.PerformBatchUpdatesCompletionHandler]
//   - [NSCustomImageRep.InitWithSizeFlippedDrawingHandler]
//   - [NSDocument.AccommodatePresentedItemDeletionWithCompletionHandler]
//   - [NSDocument.AutosaveWithImplicitCancellabilityCompletionHandler]
//   - [NSDocument.LockDocumentWithCompletionHandler]
//   - [NSDocument.LockWithCompletionHandler]
//   - [NSDocument.MoveDocumentWithCompletionHandler]
//   - [NSDocument.MoveToURLCompletionHandler]
//   - [NSDocument.SavePresentedItemChangesWithCompletionHandler]
//   - [NSDocument.SaveToURLOfTypeForSaveOperationCompletionHandler]
//   - [NSDocument.ShareDocumentWithSharingServiceCompletionHandler]
//   - [NSDocument.UnlockDocumentWithCompletionHandler]
//   - [NSDocument.UnlockWithCompletionHandler]
//   - [NSDocumentController.BeginOpenPanelForTypesCompletionHandler]
//   - [NSEvent.TrackSwipeEventWithOptionsDampenAmountThresholdMinMaxUsingHandler]
//   - [NSFilePromiseProviderDelegate.FilePromiseProviderWritePromiseToURLCompletionHandler]
//   - [NSFontAssetRequest.DownloadFontAssetsWithCompletionHandler]
//   - [NSImage.ImageWithSizeFlippedDrawingHandler]
//   - [NSPDFPanel.BeginSheetWithPDFInfoModalForWindowCompletionHandler]
//   - [NSPageLayout.BeginSheetUsingPrintInfoOnWindowCompletionHandler]
//   - [NSPrintPanel.BeginSheetUsingPrintInfoOnWindowCompletionHandler]
//   - [NSSavePanel.BeginSheetModalForWindowCompletionHandler]
//   - [NSSavePanel.BeginWithCompletionHandler]
//   - [NSSpellChecker.RequestCheckingOfStringRangeTypesOptionsInSpellDocumentWithTagCompletionHandler]
//   - [NSTableView.EnumerateAvailableRowViewsUsingBlock]
//   - [NSTableViewRowAction.RowActionWithStyleTitleHandler]
//   - [NSTextContentManager.SynchronizeTextLayoutManagers]
//   - [NSTextContentManager.SynchronizeToBackingStore]
//   - [NSTextContentStorage.SynchronizeToBackingStore]
//   - [NSTextElementProvider.SynchronizeToBackingStore]
//   - [NSWindow.BeginCriticalSheetCompletionHandler]
//   - [NSWindow.BeginSheetCompletionHandler]
//   - [NSWindow.RequestSharingOfWindowCompletionHandler]
//   - [NSWindow.RequestSharingOfWindowUsingPreviewTitleCompletionHandler]
//   - [NSWindow.TrackEventsMatchingMaskTimeoutModeHandler]
//   - [NSWindow.TransferWindowSharingToWindowCompletionHandler]
//   - [NSWorkspace.SetDefaultApplicationAtURLToOpenContentTypeCompletionHandler]
//   - [NSWorkspace.SetDefaultApplicationAtURLToOpenContentTypeOfFileAtURLCompletionHandler]
//   - [NSWorkspace.SetDefaultApplicationAtURLToOpenFileAtURLCompletionHandler]
//   - [NSWorkspace.SetDefaultApplicationAtURLToOpenURLsWithSchemeCompletionHandler]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}



// EventHandler handles The event handler block object.
//
// Used by:
//   - [NSEvent.AddGlobalMonitorForEventsMatchingMaskHandler]
//   - [NSEvent.AddLocalMonitorForEventsMatchingMaskHandler]
type EventHandler = func(*NSEvent)

// NewEventBlock wraps a Go [EventHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSEvent.AddGlobalMonitorForEventsMatchingMaskHandler]
//   - [NSEvent.AddLocalMonitorForEventsMatchingMaskHandler]
func NewEventBlock(handler EventHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSEvent
		if resultID != 0 {
			v := NSEventFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}



// ObjectHandler handles completion with a primitive value.
type ObjectHandler = func(objectivec.IObject)

// NewObjectBlock wraps a Go [ObjectHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewObjectBlock(handler ObjectHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, valID objc.ID) {
		var val objectivec.IObject
		if valID != 0 {
			obj := objectivec.ObjectFromID(valID)
			val = &obj
		}
		handler(val)
	})
	return objc.ID(block), func() { block.Release() }
}



// RangeHandler handles The originProvider block object should return the baseline origin for the first character at the adjusted range.
//   - adjustedRange: The adjusted range.
//
// Used by:
//   - [NSView.ShowDefinitionForAttributedStringRangeOptionsBaselineOriginProvider]
type RangeHandler = func(foundation.NSRange)

// NewRangeBlock wraps a Go [RangeHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSView.ShowDefinitionForAttributedStringRangeOptionsBaselineOriginProvider]
func NewRangeBlock(handler RangeHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal foundation.NSRange) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}



// RectHandler handles completion with a primitive value.
//
// Used by:
//   - [NSStepperTouchBarItem.StepperTouchBarItemWithIdentifierDrawingHandler]
type RectHandler = func(corefoundation.CGRect)

// NewRectBlock wraps a Go [RectHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSStepperTouchBarItem.StepperTouchBarItemWithIdentifierDrawingHandler]
func NewRectBlock(handler RectHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal corefoundation.CGRect) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}



// RunningApplicationErrorHandler handles The completion handler block to call asynchronously with the results.
//   - app: On success, this parameter contains a reference to the app that opened the URL. If the app didn’t open the URL successfully, this parameter is `nil`.
//   - error: On failure, this parameter contains an [NSError](<doc://com.apple.documentation/documentation/Foundation/NSError>) object indicating the reason for the failure. If the method opened the URL successfully, this parameter is `nil`.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSWorkspace.OpenApplicationAtURLConfigurationCompletionHandler]
//   - [NSWorkspace.OpenURLConfigurationCompletionHandler]
//   - [NSWorkspace.OpenURLsWithApplicationAtURLConfigurationCompletionHandler]
type RunningApplicationErrorHandler = func(*NSRunningApplication, error)

// NewRunningApplicationErrorBlock wraps a Go [RunningApplicationErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSWorkspace.OpenApplicationAtURLConfigurationCompletionHandler]
//   - [NSWorkspace.OpenURLConfigurationCompletionHandler]
//   - [NSWorkspace.OpenURLsWithApplicationAtURLConfigurationCompletionHandler]
func NewRunningApplicationErrorBlock(handler RunningApplicationErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *NSRunningApplication
		if resultID != 0 {
			v := NSRunningApplicationFromID(resultID)
			result = &v
		}
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}



// SliderAccessoryHandler is the signature for a completion handler block.
//
// Used by:
//   - [NSSliderAccessoryBehavior.BehaviorWithHandler]
type SliderAccessoryHandler = func(*NSSliderAccessory)

// NewSliderAccessoryBlock wraps a Go [SliderAccessoryHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSSliderAccessoryBehavior.BehaviorWithHandler]
func NewSliderAccessoryBlock(handler SliderAccessoryHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSSliderAccessory
		if resultID != 0 {
			v := NSSliderAccessoryFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}



// StringHandler handles The Block called when a the correction indicator is dismissed.
//   - acceptedString: The correction string the user excepted. If the user does not select a correction string nil is returned.
//
// Used by:
//   - [NSSpellChecker.ShowCorrectionIndicatorOfTypePrimaryStringAlternativeStringsForStringInRectViewCompletionHandler]
type StringHandler = func(*string)



// TextElementHandler handles A block you use to evaluate whether to continue the enumeration or tell the method to stop.
//
// Used by:
//   - [NSTextContentManager.EnumerateTextElementsFromLocationOptionsUsingBlock]
//   - [NSTextContentStorage.EnumerateTextElementsFromLocationOptionsUsingBlock]
//   - [NSTextElementProvider.EnumerateTextElementsFromLocationOptionsUsingBlock]
type TextElementHandler = func(*NSTextElement)

// NewTextElementBlock wraps a Go [TextElementHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSTextContentManager.EnumerateTextElementsFromLocationOptionsUsingBlock]
//   - [NSTextContentStorage.EnumerateTextElementsFromLocationOptionsUsingBlock]
//   - [NSTextElementProvider.EnumerateTextElementsFromLocationOptionsUsingBlock]
func NewTextElementBlock(handler TextElementHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSTextElement
		if resultID != 0 {
			v := NSTextElementFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}



// TextLayoutFragmentHandler handles A closure you provide that determines if the enumeration finishes early.
//
// Used by:
//   - [NSTextLayoutManager.EnumerateTextLayoutFragmentsFromLocationOptionsUsingBlock]
type TextLayoutFragmentHandler = func(*NSTextLayoutFragment)

// NewTextLayoutFragmentBlock wraps a Go [TextLayoutFragmentHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSTextLayoutManager.EnumerateTextLayoutFragmentsFromLocationOptionsUsingBlock]
func NewTextLayoutFragmentBlock(handler TextLayoutFragmentHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSTextLayoutFragment
		if resultID != 0 {
			v := NSTextLayoutFragmentFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}



// TextLayoutManagerTextLayoutFragmentHandler is the signature for a completion handler block.
type TextLayoutManagerTextLayoutFragmentHandler = func(*NSTextLayoutManager, *NSTextLayoutFragment)

// NewTextLayoutManagerTextLayoutFragmentBlock wraps a Go [TextLayoutManagerTextLayoutFragmentHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewTextLayoutManagerTextLayoutFragmentBlock(handler TextLayoutManagerTextLayoutFragmentHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID) {
		var result *NSTextLayoutManager
		if resultID != 0 {
			v := NSTextLayoutManagerFromID(resultID)
			result = &v
		}
		var extra0 *NSTextLayoutFragment
		if extra0ID != 0 {
			v := NSTextLayoutFragmentFromID(extra0ID)
			extra0 = &v
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}



// TextPreviewHandler handles A completion handler to execute when you are done.
//
// Used by:
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsPreviewForRectInContextCompletion]
type TextPreviewHandler = func(*NSTextPreview)

// NewTextPreviewBlock wraps a Go [TextPreviewHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsPreviewForRectInContextCompletion]
func NewTextPreviewBlock(handler TextPreviewHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSTextPreview
		if resultID != 0 {
			v := NSTextPreviewFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}



// URLErrorHandler handles A block to be called on the supplied operationQueue when the promised file is ready to be read.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSFilePromiseReceiver.ReceivePromisedFilesAtDestinationOptionsOperationQueueReader]
type URLErrorHandler = func(*foundation.NSURL, error)

// NewURLErrorBlock wraps a Go [URLErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSFilePromiseReceiver.ReceivePromisedFilesAtDestinationOptionsOperationQueueReader]
func NewURLErrorBlock(handler URLErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *foundation.NSURL
		if resultID != 0 {
			v := foundation.NSURLFromID(resultID)
			result = &v
		}
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}



// UUIDHandler handles A handler to execute with the required information.
//
// Used by:
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsRangeInContextWithIdentifierForPointCompletion]
type UUIDHandler = func(*foundation.NSUUID)

// NewUUIDBlock wraps a Go [UUIDHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsRangeInContextWithIdentifierForPointCompletion]
func NewUUIDBlock(handler UUIDHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *foundation.NSUUID
		if resultID != 0 {
			v := foundation.NSUUIDFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}



// ViewHandler handles A completion handler to execute when you are done.
//
// Used by:
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion]
type ViewHandler = func(*NSView)

// NewViewBlock wraps a Go [ViewHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsDecorationContainerViewForRangeInContextCompletion]
func NewViewBlock(handler ViewHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSView
		if resultID != 0 {
			v := NSViewFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}



// VoidHandler handles The completion handler block to call when the Versions browser is fully dismissed.
//
// Used by:
//   - [NSAnimationContext.RunAnimationGroupCompletionHandler]
//   - [NSAppearance.PerformAsCurrentDrawingAppearance]
//   - [NSApplicationDelegate.ApplicationContinueUserActivityRestorationHandler]
//   - [NSCollectionView.PerformBatchUpdatesCompletionHandler]
//   - [NSDocument.ContinueActivityUsingBlock]
//   - [NSDocument.ContinueAsynchronousWorkOnMainThreadUsingBlock]
//   - [NSDocument.PerformActivityWithSynchronousWaitingUsingBlock]
//   - [NSDocument.PerformAsynchronousFileAccessUsingBlock]
//   - [NSDocument.PerformSynchronousFileAccessUsingBlock]
//   - [NSDocument.RelinquishPresentedItemToReader]
//   - [NSDocument.RelinquishPresentedItemToWriter]
//   - [NSDocument.StopBrowsingVersionsWithCompletionHandler]
//   - [NSScrubber.PerformSequentialBatchUpdates]
//   - [NSSharingService.InitWithTitleImageAlternateImageHandler]
//   - [NSStoryboardSegue.SegueWithIdentifierSourceDestinationPerformHandler]
//   - [NSTableViewDiffableDataSource.ApplySnapshotAnimatingDifferencesCompletion]
//   - [NSTextContentManager.PerformEditingTransactionUsingBlock]
//   - [NSTextContentStorage.PerformEditingTransactionForTextStorageUsingBlock]
//   - [NSTextLayoutManager.EnumerateRenderingAttributesFromLocationReverseUsingBlock]
//   - [NSTextStorageObserving.PerformEditingTransactionForTextStorageUsingBlock]
//   - [NSViewController.TransitionFromViewControllerToViewControllerOptionsCompletionHandler]
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion]
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorPrepareForTextAnimationForRangeInContextCompletion]
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsContextsForScopeCompletion]
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorSelectRangesInContextCompletion]
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorWillChangeToStateCompletion]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSAnimationContext.RunAnimationGroupCompletionHandler]
//   - [NSAppearance.PerformAsCurrentDrawingAppearance]
//   - [NSApplicationDelegate.ApplicationContinueUserActivityRestorationHandler]
//   - [NSCollectionView.PerformBatchUpdatesCompletionHandler]
//   - [NSDocument.ContinueActivityUsingBlock]
//   - [NSDocument.ContinueAsynchronousWorkOnMainThreadUsingBlock]
//   - [NSDocument.PerformActivityWithSynchronousWaitingUsingBlock]
//   - [NSDocument.PerformAsynchronousFileAccessUsingBlock]
//   - [NSDocument.PerformSynchronousFileAccessUsingBlock]
//   - [NSDocument.RelinquishPresentedItemToReader]
//   - [NSDocument.RelinquishPresentedItemToWriter]
//   - [NSDocument.StopBrowsingVersionsWithCompletionHandler]
//   - [NSScrubber.PerformSequentialBatchUpdates]
//   - [NSSharingService.InitWithTitleImageAlternateImageHandler]
//   - [NSStoryboardSegue.SegueWithIdentifierSourceDestinationPerformHandler]
//   - [NSTableViewDiffableDataSource.ApplySnapshotAnimatingDifferencesCompletion]
//   - [NSTextContentManager.PerformEditingTransactionUsingBlock]
//   - [NSTextContentStorage.PerformEditingTransactionForTextStorageUsingBlock]
//   - [NSTextLayoutManager.EnumerateRenderingAttributesFromLocationReverseUsingBlock]
//   - [NSTextStorageObserving.PerformEditingTransactionForTextStorageUsingBlock]
//   - [NSViewController.TransitionFromViewControllerToViewControllerOptionsCompletionHandler]
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorFinishTextAnimationForRangeInContextCompletion]
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorPrepareForTextAnimationForRangeInContextCompletion]
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsContextsForScopeCompletion]
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorSelectRangesInContextCompletion]
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorWillChangeToStateCompletion]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}



// WindowErrorHandler handles A block object to execute with the results of creating the window.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSApplication.RestoreWindowWithIdentifierStateCompletionHandler]
//   - [NSDocument.RestoreDocumentWindowWithIdentifierStateCompletionHandler]
//   - [NSDocumentController.RestoreWindowWithIdentifierStateCompletionHandler]
//   - [NSWindowRestoration.RestoreWindowWithIdentifierStateCompletionHandler]
type WindowErrorHandler = func(*NSWindow, error)

// NewWindowErrorBlock wraps a Go [WindowErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSApplication.RestoreWindowWithIdentifierStateCompletionHandler]
//   - [NSDocument.RestoreDocumentWindowWithIdentifierStateCompletionHandler]
//   - [NSDocumentController.RestoreWindowWithIdentifierStateCompletionHandler]
//   - [NSWindowRestoration.RestoreWindowWithIdentifierStateCompletionHandler]
func NewWindowErrorBlock(handler WindowErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *NSWindow
		if resultID != 0 {
			v := NSWindowFromID(resultID)
			result = &v
		}
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}



// WorkspaceAuthorizationErrorHandler handles The completion handler to call when the authorization request is completed.
//   - authorization: The authorization granted for this app. Use it when creating a new [FileManager](<doc://com.apple.documentation/documentation/Foundation/FileManager>) with [init(authorization:)](<doc://com.apple.documentation/documentation/Foundation/FileManager/init(authorization:)>).
//   - error: `nil` if the app is authorized; otherwise, a pointer to the authorization error.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSWorkspace.RequestAuthorizationOfTypeCompletionHandler]
type WorkspaceAuthorizationErrorHandler = func(*NSWorkspaceAuthorization, error)

// NewWorkspaceAuthorizationErrorBlock wraps a Go [WorkspaceAuthorizationErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSWorkspace.RequestAuthorizationOfTypeCompletionHandler]
func NewWorkspaceAuthorizationErrorBlock(handler WorkspaceAuthorizationErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *NSWorkspaceAuthorization
		if resultID != 0 {
			v := NSWorkspaceAuthorizationFromID(resultID)
			result = &v
		}
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}



// structCGRectHandler handles completion with a primitive value.
type structCGRectHandler = func(corefoundation.CGRect)

// NewstructCGRectBlock wraps a Go [structCGRectHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewstructCGRectBlock(handler structCGRectHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal corefoundation.CGRect) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}



