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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSAnimationContext
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSAppearance
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSAppearanceFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// ArrayHandler handles The Block used to add results to the search.
//   - items: The items to add to the results array. The `handleMatchedItems` block can be invoked from any thread desired.  If it is called more than once the additional results will be appended after previous items until the maximum is reached.
//
// Used by:
//   - [NSUserInterfaceItemSearching.SearchForItemsWithSearchStringResultLimitMatchedItemHandler]
type ArrayHandler = func(*foundation.INSArray)

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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *foundation.NSAttributedString
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := foundation.NSAttributedStringFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// BoolCGRectHandler handles A block that draws the image representation content in the provided graphics context.
//   - dstRect: The destination rectangle in which to draw. The coordinates of this rectangle are specified in points.
//
// Used by:
//   - [NSCustomImageRep.InitWithSizeFlippedDrawingHandler]
//   - [NSImage.ImageWithSizeFlippedDrawingHandler]
type BoolCGRectHandler = func(corefoundation.CGRect) bool

// NewBoolCGRectBlock wraps a Go [BoolCGRectHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSCustomImageRep.InitWithSizeFlippedDrawingHandler]
//   - [NSImage.ImageWithSizeFlippedDrawingHandler]
func NewBoolCGRectBlock(handler BoolCGRectHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal corefoundation.CGRect) bool {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// BoolHandler handles A completion handler block to execute when the changes made in the `updates` block have finished animating.
//   - finished: A Boolean value indicating whether the animations completed successfully. The value of this parameter is [true](<doc://com.apple.documentation/documentation/Swift/true>) if the animations ran to completion or [false](<doc://com.apple.documentation/documentation/Swift/false>) if they were interrupted.
//
// Used by:
//   - [NSCollectionView.PerformBatchUpdatesCompletionHandler]
//   - [NSDocument.LockDocumentWithCompletionHandler]
//   - [NSDocument.MoveDocumentWithCompletionHandler]
//   - [NSDocument.ShareDocumentWithSharingServiceCompletionHandler]
//   - [NSDocument.UnlockDocumentWithCompletionHandler]
type BoolHandler = func(bool)

// NewBoolBlock wraps a Go [BoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSCollectionView.PerformBatchUpdatesCompletionHandler]
//   - [NSDocument.LockDocumentWithCompletionHandler]
//   - [NSDocument.MoveDocumentWithCompletionHandler]
//   - [NSDocument.ShareDocumentWithSharingServiceCompletionHandler]
//   - [NSDocument.UnlockDocumentWithCompletionHandler]
func NewBoolBlock(handler BoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal bool) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// CGPointNSRangeHandler handles The originProvider block object should return the baseline origin for the first character at the adjusted range.
//   - adjustedRange: The adjusted range.
//
// Used by:
//   - [NSView.ShowDefinitionForAttributedStringRangeOptionsBaselineOriginProvider]
type CGPointNSRangeHandler = func(foundation.NSRange) corefoundation.CGPoint

// NewCGPointNSRangeBlock wraps a Go [CGPointNSRangeHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSView.ShowDefinitionForAttributedStringRangeOptionsBaselineOriginProvider]
func NewCGPointNSRangeBlock(handler CGPointNSRangeHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal foundation.NSRange) corefoundation.CGPoint {
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// CGRectBoolHandler handles The block to apply to the glyph range.
//   - rect: The current enclosing rectangle.
//   - stop: A reference to a Boolean value. The block can set the value to [true](<doc://com.apple.documentation/documentation/Swift/true>) to stop further processing of the array. The stop argument is an out-only argument. You should only set this Boolean to [true](<doc://com.apple.documentation/documentation/Swift/true>) within the block.
//
// Used by:
//   - [NSLayoutManager.EnumerateEnclosingRectsForGlyphRangeWithinSelectedGlyphRangeInTextContainerUsingBlock]
type CGRectBoolHandler = func(corefoundation.CGRect, bool)

// NewCGRectBoolBlock wraps a Go [CGRectBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSLayoutManager.EnumerateEnclosingRectsForGlyphRangeWithinSelectedGlyphRangeInTextContainerUsingBlock]
func NewCGRectBoolBlock(handler CGRectBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive corefoundation.CGRect, extra0 bool) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// CGRectCGRectTextContainerNSRangeBoolHandler handles The block to apply to the glyph range.
//   - rect: The current line fragment rectangle.
//   - usedRect: The portion of the line fragment rectangle that actually contains glyphs or other marks that are drawn (including the text container’s line fragment padding).
//   - textContainer: The text container in which the glyphs are laid out.
//   - glyphRange: The range of glyphs laid out in the current line fragment.
//   - stop: A reference to a Boolean value. The block can set the value to [true](<doc://com.apple.documentation/documentation/Swift/true>) to stop further processing of the array. The stop argument is an out-only argument. You should only set this Boolean to [true](<doc://com.apple.documentation/documentation/Swift/true>) within the block.
//
// Used by:
//   - [NSLayoutManager.EnumerateLineFragmentsForGlyphRangeUsingBlock]
type CGRectCGRectTextContainerNSRangeBoolHandler = func(corefoundation.CGRect, corefoundation.CGRect, *NSTextContainer, foundation.NSRange, bool)

// NewCGRectCGRectTextContainerNSRangeBoolBlock wraps a Go [CGRectCGRectTextContainerNSRangeBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSLayoutManager.EnumerateLineFragmentsForGlyphRangeUsingBlock]
func NewCGRectCGRectTextContainerNSRangeBoolBlock(handler CGRectCGRectTextContainerNSRangeBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive corefoundation.CGRect, extra0 corefoundation.CGRect, extra1ID objc.ID, extra2 foundation.NSRange, extra3 bool) {
		var extra1 *NSTextContainer
		if extra1ID != 0 {
			objc.Send[objc.ID](extra1ID, objc.Sel("retain"))
			v := NSTextContainerFromID(extra1ID)
			extra1 = &v
		}
		handler(primitive, extra0, extra1, extra2, extra3)
	})
	return objc.ID(block), func() { block.Release() }
}

// CGRectFloat64Handler handles A block that draws a graphical representation of the stepper’s value in the specified rectangle.
//
// Used by:
//   - [NSStepperTouchBarItem.StepperTouchBarItemWithIdentifierDrawingHandler]
type CGRectFloat64Handler = func(corefoundation.CGRect, float64)

// NewCGRectFloat64Block wraps a Go [CGRectFloat64Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSStepperTouchBarItem.StepperTouchBarItemWithIdentifierDrawingHandler]
func NewCGRectFloat64Block(handler CGRectFloat64Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive corefoundation.CGRect, extra0 float64) {
		handler(primitive, extra0)
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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSColor
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSColorFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// DocumentBoolErrorHandler handles The completion handler block object passed in to be called at some point in the future, perhaps after the method invocation has returned.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSDocumentController.OpenDocumentWithContentsOfURLDisplayCompletionHandler]
//   - [NSDocumentController.ReopenDocumentForURLWithContentsOfURLDisplayCompletionHandler]
type DocumentBoolErrorHandler = func(*NSDocument, bool, error)

// NewDocumentBoolErrorBlock wraps a Go [DocumentBoolErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSDocumentController.OpenDocumentWithContentsOfURLDisplayCompletionHandler]
//   - [NSDocumentController.ReopenDocumentForURLWithContentsOfURLDisplayCompletionHandler]
func NewDocumentBoolErrorBlock(handler DocumentBoolErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 bool, errID objc.ID) {
		var result *NSDocument
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSDocumentFromID(resultID)
			result = &v
		}
		handler(result, extra0, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// DraggingItemIntBoolHandler handles The block to execute for the enumeration.
//
// Used by:
//   - [NSDraggingInfo.EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock]
//   - [NSDraggingSession.EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock]
type DraggingItemIntBoolHandler = func(*NSDraggingItem, int, bool)

// NewDraggingItemIntBoolBlock wraps a Go [DraggingItemIntBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSDraggingInfo.EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock]
//   - [NSDraggingSession.EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock]
func NewDraggingItemIntBoolBlock(handler DraggingItemIntBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 int, extra1 bool) {
		var result *NSDraggingItem
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSDraggingItemFromID(resultID)
			result = &v
		}
		handler(result, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// ErrorHandler handles The completion handler block object passed in to be invoked at some point in the future, perhaps after the method invocation has returned.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSDocument.AccommodatePresentedItemDeletionWithCompletionHandler]
//   - [NSDocument.AutosaveWithImplicitCancellabilityCompletionHandler]
//   - [NSDocument.LockWithCompletionHandler]
//   - [NSDocument.MoveToURLCompletionHandler]
//   - [NSDocument.PerformActivityWithSynchronousWaitingUsingBlock]
//   - [NSDocument.PerformAsynchronousFileAccessUsingBlock]
//   - [NSDocument.RelinquishPresentedItemToReader]
//   - [NSDocument.RelinquishPresentedItemToWriter]
//   - [NSDocument.SavePresentedItemChangesWithCompletionHandler]
//   - [NSDocument.SaveToURLOfTypeForSaveOperationCompletionHandler]
//   - [NSDocument.UnlockWithCompletionHandler]
//   - [NSFilePromiseProviderDelegate.FilePromiseProviderWritePromiseToURLCompletionHandler]
//   - [NSFontAssetRequest.DownloadFontAssetsWithCompletionHandler]
//   - [NSTextContentManager.SynchronizeTextLayoutManagers]
//   - [NSTextContentManager.SynchronizeToBackingStore]
//   - [NSTextContentStorage.SynchronizeToBackingStore]
//   - [NSTextElementProvider.SynchronizeToBackingStore]
//   - [NSWindow.RequestSharingOfWindowCompletionHandler]
//   - [NSWindow.RequestSharingOfWindowUsingPreviewTitleCompletionHandler]
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
//   - [NSDocument.AccommodatePresentedItemDeletionWithCompletionHandler]
//   - [NSDocument.AutosaveWithImplicitCancellabilityCompletionHandler]
//   - [NSDocument.LockWithCompletionHandler]
//   - [NSDocument.MoveToURLCompletionHandler]
//   - [NSDocument.PerformActivityWithSynchronousWaitingUsingBlock]
//   - [NSDocument.PerformAsynchronousFileAccessUsingBlock]
//   - [NSDocument.RelinquishPresentedItemToReader]
//   - [NSDocument.RelinquishPresentedItemToWriter]
//   - [NSDocument.SavePresentedItemChangesWithCompletionHandler]
//   - [NSDocument.SaveToURLOfTypeForSaveOperationCompletionHandler]
//   - [NSDocument.UnlockWithCompletionHandler]
//   - [NSFilePromiseProviderDelegate.FilePromiseProviderWritePromiseToURLCompletionHandler]
//   - [NSFontAssetRequest.DownloadFontAssetsWithCompletionHandler]
//   - [NSTextContentManager.SynchronizeTextLayoutManagers]
//   - [NSTextContentManager.SynchronizeToBackingStore]
//   - [NSTextContentStorage.SynchronizeToBackingStore]
//   - [NSTextElementProvider.SynchronizeToBackingStore]
//   - [NSWindow.RequestSharingOfWindowCompletionHandler]
//   - [NSWindow.RequestSharingOfWindowUsingPreviewTitleCompletionHandler]
//   - [NSWindow.TransferWindowSharingToWindowCompletionHandler]
//   - [NSWorkspace.SetDefaultApplicationAtURLToOpenContentTypeCompletionHandler]
//   - [NSWorkspace.SetDefaultApplicationAtURLToOpenContentTypeOfFileAtURLCompletionHandler]
//   - [NSWorkspace.SetDefaultApplicationAtURLToOpenFileAtURLCompletionHandler]
//   - [NSWorkspace.SetDefaultApplicationAtURLToOpenURLsWithSchemeCompletionHandler]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		handler(foundation.SafeErrorFrom(errID))
	})
	objc.SetNSErrorBlockSignature(block)
	return objc.ID(block), func() { block.Release() }
}

// EventBoolHandler handles A block that is called to track the events.
//   - event: The event to examine.
//   - stop: A Boolean value that indicates when tracking should stop.
//
// Used by:
//   - [NSWindow.TrackEventsMatchingMaskTimeoutModeHandler]
type EventBoolHandler = func(*NSEvent, bool)

// NewEventBoolBlock wraps a Go [EventBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSWindow.TrackEventsMatchingMaskTimeoutModeHandler]
func NewEventBoolBlock(handler EventBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 bool) {
		var result *NSEvent
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSEventFromID(resultID)
			result = &v
		}
		handler(result, extra0)
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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSEvent
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSEventFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// Float32Handler handles completion with a primitive value.
type Float32Handler = func(float32)

// NewFloat32Block wraps a Go [Float32Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewFloat32Block(handler Float32Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal float32) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// Float64NSEventPhaseBoolBoolHandler handles The Block used as the tracking handler.
//   - gestureAmount: The amount of gesture that you should display in the user interface. This may be a fractional amount.
//   - phase: The phase of the physical gesture as performed by the user. See [NSEvent.Phase](<doc://com.apple.appkit/documentation/AppKit/NSEvent/Phase-swift.struct>) for possible values. When the phase is either [ended](<doc://com.apple.appkit/documentation/AppKit/NSEvent/Phase-swift.struct/ended>), or [mayBegin](<doc://com.apple.appkit/documentation/AppKit/NSEvent/Phase-swift.struct/mayBegin>), the user has physically ended the gesture successfully or un-successfully, respectively.
//   - isComplete: Signifies the swipe and animation are complete and you should release any temporary animation objects.
//   - stop: A reference to a Boolean value. The Block can set the value to [true](<doc://com.apple.documentation/documentation/Swift/true>) to stop further processing of the array. The `stop` argument is an out-only argument. You should only ever set this Boolean to [true](<doc://com.apple.documentation/documentation/Swift/true>) within the Block
//
// Used by:
//   - [NSEvent.TrackSwipeEventWithOptionsDampenAmountThresholdMinMaxUsingHandler]
type Float64NSEventPhaseBoolBoolHandler = func(float64, NSEventPhase, bool, bool)

// NewFloat64NSEventPhaseBoolBoolBlock wraps a Go [Float64NSEventPhaseBoolBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSEvent.TrackSwipeEventWithOptionsDampenAmountThresholdMinMaxUsingHandler]
func NewFloat64NSEventPhaseBoolBoolBlock(handler Float64NSEventPhaseBoolBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive float64, extra0 NSEventPhase, extra1 bool, extra2 bool) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// Float64NSTextLocationBoolBoolHandler handles The closure to invoke once for each logical caret edge in the line fragment, in left-to-right visual order.
//
// Used by:
//   - [NSTextLayoutManager.EnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock]
//   - [NSTextSelectionDataSource.EnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock]
type Float64NSTextLocationBoolBoolHandler = func(float64, NSTextLocation, bool, bool)

// NewFloat64NSTextLocationBoolBoolBlock wraps a Go [Float64NSTextLocationBoolBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSTextLayoutManager.EnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock]
//   - [NSTextSelectionDataSource.EnumerateCaretOffsetsInLineFragmentAtLocationUsingBlock]
func NewFloat64NSTextLocationBoolBoolBlock(handler Float64NSTextLocationBoolBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive float64, extra0ID objc.ID, extra1 bool, extra2 bool) {
		var extra0 NSTextLocation
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			extra0 = NSTextLocationObjectFromID(extra0ID)
		}
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// IntHandler handles The completion handler that runs when the user clicks the OK or Cancel button in the Open dialog.
//
// Used by:
//   - [NSDocumentController.BeginOpenPanelForTypesCompletionHandler]
//   - [NSPDFPanel.BeginSheetWithPDFInfoModalForWindowCompletionHandler]
type IntHandler = func(int)

// NewIntBlock wraps a Go [IntHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSDocumentController.BeginOpenPanelForTypesCompletionHandler]
//   - [NSPDFPanel.BeginSheetWithPDFInfoModalForWindowCompletionHandler]
func NewIntBlock(handler IntHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal int) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// IntTextCheckingResultArrayHandler handles completion with primitive and object results.
//
// Used by:
//   - [NSSpellChecker.RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler]
type IntTextCheckingResultArrayHandler = func(int, *foundation.NSArray)

// NewIntTextCheckingResultArrayBlock wraps a Go [IntTextCheckingResultArrayHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSSpellChecker.RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler]
func NewIntTextCheckingResultArrayBlock(handler IntTextCheckingResultArrayHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive int, extra0ID objc.ID) {
		var extra0 *foundation.NSArray
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := foundation.NSArrayFromID(extra0ID)
			extra0 = &v
		}
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// IntTextCheckingResultArrayOrthographyIntHandler handles The completion handler block object will be called (in an arbitrary context) when results are available, with the sequence number and results.
//   - sequenceNumber: A monotonically increasing sequence number.
//   - results: An array of [NSTextCheckingResult](<doc://com.apple.documentation/documentation/Foundation/NSTextCheckingResult>) objects describing particular items found during checking and their individual ranges, sorted by range origin, then range end, then result type.
//   - orthography: The orthography of the string.
//   - wordCount: The number of words in the range of the string.
//
// Used by:
//   - [NSSpellChecker.RequestCheckingOfStringRangeTypesOptionsInSpellDocumentWithTagCompletionHandler]
type IntTextCheckingResultArrayOrthographyIntHandler = func(int, *foundation.NSArray, *foundation.NSOrthography, int)

// NewIntTextCheckingResultArrayOrthographyIntBlock wraps a Go [IntTextCheckingResultArrayOrthographyIntHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSSpellChecker.RequestCheckingOfStringRangeTypesOptionsInSpellDocumentWithTagCompletionHandler]
func NewIntTextCheckingResultArrayOrthographyIntBlock(handler IntTextCheckingResultArrayOrthographyIntHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive int, extra0ID objc.ID, extra1ID objc.ID, extra2 int) {
		var extra0 *foundation.NSArray
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := foundation.NSArrayFromID(extra0ID)
			extra0 = &v
		}
		var extra1 *foundation.NSOrthography
		if extra1ID != 0 {
			objc.Send[objc.ID](extra1ID, objc.Sel("retain"))
			v := foundation.NSOrthographyFromID(extra1ID)
			extra1 = &v
		}
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// MenuHandler handles The closure to invoke when someone selects the menu item.
//
// Used by:
//   - [NSMenu.PaletteMenuWithColorsTitlesSelectionHandler]
//   - [NSMenu.PaletteMenuWithColorsTitlesTemplateImageSelectionHandler]
type MenuHandler = func(*NSMenu)

// NewMenuBlock wraps a Go [MenuHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSMenu.PaletteMenuWithColorsTitlesSelectionHandler]
//   - [NSMenu.PaletteMenuWithColorsTitlesTemplateImageSelectionHandler]
func NewMenuBlock(handler MenuHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSMenu
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSMenuFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// ModalResponseHandler handles The completion handler that gets called when the sheet’s modal session ends.
//   - result: The action taken by the user. The value of this parameter is [NSFileHandlingPanelOKButton](<doc://com.apple.appkit/documentation/AppKit/NSFileHandlingPanelOKButton>) if the user chose the OK button or [NSFileHandlingPanelCancelButton](<doc://com.apple.appkit/documentation/AppKit/NSFileHandlingPanelCancelButton>) if the user chose the Cancel button.
//
// Used by:
//   - [NSAlert.BeginSheetModalForWindowCompletionHandler]
//   - [NSSavePanel.BeginSheetModalForWindowCompletionHandler]
//   - [NSSavePanel.BeginWithCompletionHandler]
//   - [NSWindow.BeginCriticalSheetCompletionHandler]
//   - [NSWindow.BeginSheetCompletionHandler]
type ModalResponseHandler = func(NSModalResponse)

// NewModalResponseBlock wraps a Go [ModalResponseHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSAlert.BeginSheetModalForWindowCompletionHandler]
//   - [NSSavePanel.BeginSheetModalForWindowCompletionHandler]
//   - [NSSavePanel.BeginWithCompletionHandler]
//   - [NSWindow.BeginCriticalSheetCompletionHandler]
//   - [NSWindow.BeginSheetCompletionHandler]
func NewModalResponseBlock(handler ModalResponseHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal NSModalResponse) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// NSAttributedStringIObjectHandler handles completion with primitive and object results.
type NSAttributedStringIObjectHandler = func(objectivec.IObject, int64)

// NewNSAttributedStringIObjectBlock wraps a Go [NSAttributedStringIObjectHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNSAttributedStringIObjectBlock(handler NSAttributedStringIObjectHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, extra0 int64) {
		var primitive objectivec.IObject
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			obj := objectivec.ObjectFromID(primitiveID)
			primitive = &obj
		}
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// NSBezierPathArrayHandler handles A handler to execute with the required information.
//
// Used by:
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion]
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion]
type NSBezierPathArrayHandler = func(*[]NSBezierPath)

// NewNSBezierPathArrayBlock wraps a Go [NSBezierPathArrayHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsBoundingBezierPathsForRangeInContextCompletion]
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsUnderlinePathsForRangeInContextCompletion]
func NewNSBezierPathArrayBlock(handler NSBezierPathArrayHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *[]NSBezierPath
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]NSBezierPath, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = NSBezierPathFromID(item.GetID())
			}
			result = &res
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// NSRangeUUIDHandler handles A handler to execute with the required information.
//
// Used by:
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsRangeInContextWithIdentifierForPointCompletion]
type NSRangeUUIDHandler = func(foundation.NSRange, *foundation.NSUUID)

// NewNSRangeUUIDBlock wraps a Go [NSRangeUUIDHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsRangeInContextWithIdentifierForPointCompletion]
func NewNSRangeUUIDBlock(handler NSRangeUUIDHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive foundation.NSRange, extra0ID objc.ID) {
		var extra0 *foundation.NSUUID
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := foundation.NSUUIDFromID(extra0ID)
			extra0 = &v
		}
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// NSTextLocationBoolHandler handles A closure to invoke to evaluate the container boundaries; end the enumeration early by returning `false`.
//
// Used by:
//   - [NSTextLayoutManager.EnumerateContainerBoundariesFromLocationReverseUsingBlock]
//   - [NSTextSelectionDataSource.EnumerateContainerBoundariesFromLocationReverseUsingBlock]
type NSTextLocationBoolHandler = func(NSTextLocation, bool)

// NewNSTextLocationBoolBlock wraps a Go [NSTextLocationBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSTextLayoutManager.EnumerateContainerBoundariesFromLocationReverseUsingBlock]
//   - [NSTextSelectionDataSource.EnumerateContainerBoundariesFromLocationReverseUsingBlock]
func NewNSTextLocationBoolBlock(handler NSTextLocationBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 bool) {
		var result NSTextLocation
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = NSTextLocationObjectFromID(resultID)
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// NSTextPreviewArrayHandler handles A completion handler to execute when you are done.
//
// Used by:
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion]
type NSTextPreviewArrayHandler = func(*[]NSTextPreview)

// NewNSTextPreviewArrayBlock wraps a Go [NSTextPreviewArrayHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsPreviewForTextAnimationOfRangeInContextCompletion]
func NewNSTextPreviewArrayBlock(handler NSTextPreviewArrayHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *[]NSTextPreview
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]NSTextPreview, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = NSTextPreviewFromID(item.GetID())
			}
			result = &res
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// NSURLArrayHandler handles The completion handler that is called when the user clicks the OK or Cancel button in the open panel.
//
// Used by:
//   - [NSDocumentController.BeginOpenPanelWithCompletionHandler]
type NSURLArrayHandler = func(*[]foundation.NSURL)

// NewNSURLArrayBlock wraps a Go [NSURLArrayHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSDocumentController.BeginOpenPanelWithCompletionHandler]
func NewNSURLArrayBlock(handler NSURLArrayHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *[]foundation.NSURL
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]foundation.NSURL, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = foundation.NSURLFromID(item.GetID())
			}
			result = &res
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// NSValueArrayHandler handles A completion handler to execute when you are done.
//
// Used by:
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion]
type NSValueArrayHandler = func(*[]foundation.NSValue)

// NewNSValueArrayBlock wraps a Go [NSValueArrayHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSWritingToolsCoordinatorDelegate.WritingToolsCoordinatorRequestsSingleContainerSubrangesOfRangeInContextCompletion]
func NewNSValueArrayBlock(handler NSValueArrayHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *[]foundation.NSValue
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]foundation.NSValue, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = foundation.NSValueFromID(item.GetID())
			}
			result = &res
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// PageLayoutResultHandler handles completion with a primitive value.
//
// Used by:
//   - [NSPageLayout.BeginSheetUsingPrintInfoOnWindowCompletionHandler]
type PageLayoutResultHandler = func(NSPageLayoutResult)

// NewPageLayoutResultBlock wraps a Go [PageLayoutResultHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSPageLayout.BeginSheetUsingPrintInfoOnWindowCompletionHandler]
func NewPageLayoutResultBlock(handler PageLayoutResultHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal NSPageLayoutResult) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// PrintPanelResultHandler handles completion with a primitive value.
//
// Used by:
//   - [NSPrintPanel.BeginSheetUsingPrintInfoOnWindowCompletionHandler]
type PrintPanelResultHandler = func(NSPrintPanelResult)

// NewPrintPanelResultBlock wraps a Go [PrintPanelResultHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSPrintPanel.BeginSheetUsingPrintInfoOnWindowCompletionHandler]
func NewPrintPanelResultBlock(handler PrintPanelResultHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal NSPrintPanelResult) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// RunningApplicationErrorHandler handles The completion handler block to call asynchronously with the results.
//   - app: On success, this parameter contains a reference to the app that opened the URL. If the app didn’t open the URL successfully, this parameter is `nil`.
//   - error: On failure, this parameter contains an [NSError](<doc://com.apple.documentation/documentation/Foundation/NSError>) object indicating the reason for the failure. If the method opened the URL successfully, this parameter is `nil`.
//
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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *NSRunningApplication
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSRunningApplicationFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSSliderAccessory
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
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

// StringSetErrorHandler handles A block the system invokes after detecting patterns on the pasteboard.
//
// Used by:
//   - [NSPasteboard.DetectPatternsForPatternsCompletionHandler]
//   - [NSPasteboardItem.DetectPatternsForPatternsCompletionHandler]
type StringSetErrorHandler = func(*foundation.INSSet, error)

// StringTextRangeTextRangeBoolHandler handles A closure to invoke to evaluate the substrings; end the enumeration early by returning `false`.
//
// Used by:
//   - [NSTextLayoutManager.EnumerateSubstringsFromLocationOptionsUsingBlock]
//   - [NSTextSelectionDataSource.EnumerateSubstringsFromLocationOptionsUsingBlock]
type StringTextRangeTextRangeBoolHandler = func(*string, *NSTextRange, *NSTextRange, *bool)

// StringidDictionaryErrorHandler handles A block the system invokes after detecting metadata on the pasteboard.
//
// Used by:
//   - [NSPasteboard.DetectMetadataForTypesCompletionHandler]
//   - [NSPasteboard.DetectValuesForPatternsCompletionHandler]
//   - [NSPasteboardItem.DetectMetadataForTypesCompletionHandler]
//   - [NSPasteboardItem.DetectValuesForPatternsCompletionHandler]
type StringidDictionaryErrorHandler = func(*foundation.INSDictionary, error)

// TableRowViewIntHandler handles The [Block] to apply to elements in the set.
//   - rowView: The view for the row.
//   - row: The index of the row.
//
// Used by:
//   - [NSTableView.EnumerateAvailableRowViewsUsingBlock]
type TableRowViewIntHandler = func(*NSTableRowView, int)

// NewTableRowViewIntBlock wraps a Go [TableRowViewIntHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSTableView.EnumerateAvailableRowViewsUsingBlock]
func NewTableRowViewIntBlock(handler TableRowViewIntHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 int) {
		var result *NSTableRowView
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSTableRowViewFromID(resultID)
			result = &v
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// TableViewRowActionIntHandler handles The block to execute when the user clicks the button associated with this action.
//   - action: The action object representing the action that the user selected.
//   - indexPath: The table row that the user acted on.
//
// Used by:
//   - [NSTableViewRowAction.RowActionWithStyleTitleHandler]
type TableViewRowActionIntHandler = func(*NSTableViewRowAction, int)

// NewTableViewRowActionIntBlock wraps a Go [TableViewRowActionIntHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSTableViewRowAction.RowActionWithStyleTitleHandler]
func NewTableViewRowActionIntBlock(handler TableViewRowActionIntHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 int) {
		var result *NSTableViewRowAction
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSTableViewRowActionFromID(resultID)
			result = &v
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSTextElement
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSTextLayoutFragment
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID) {
		var result *NSTextLayoutManager
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSTextLayoutManagerFromID(resultID)
			result = &v
		}
		var extra0 *NSTextLayoutFragment
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSTextPreview
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSTextPreviewFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// TextRangeCGRectFloat64TextContainerHandler handles A closure you provide to determine if the enumeration finishes early.
//
// Used by:
//   - [NSTextLayoutManager.EnumerateTextSegmentsInRangeTypeOptionsUsingBlock]
type TextRangeCGRectFloat64TextContainerHandler = func(*NSTextRange, corefoundation.CGRect, float64, *NSTextContainer)

// NewTextRangeCGRectFloat64TextContainerBlock wraps a Go [TextRangeCGRectFloat64TextContainerHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSTextLayoutManager.EnumerateTextSegmentsInRangeTypeOptionsUsingBlock]
func NewTextRangeCGRectFloat64TextContainerBlock(handler TextRangeCGRectFloat64TextContainerHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 corefoundation.CGRect, extra1 float64, extra2ID objc.ID) {
		var result *NSTextRange
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSTextRangeFromID(resultID)
			result = &v
		}
		var extra2 *NSTextContainer
		if extra2ID != 0 {
			objc.Send[objc.ID](extra2ID, objc.Sel("retain"))
			v := NSTextContainerFromID(extra2ID)
			extra2 = &v
		}
		handler(result, extra0, extra1, extra2)
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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *foundation.NSURL
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := foundation.NSURLFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// URLNSURLDictionaryErrorHandler handles The completion handler block object to call when the operation completes.
//   - newURLs: A dictionary parameter whose keys and values are [NSURL](<doc://com.apple.documentation/documentation/Foundation/NSURL>) objects. Each key is a URL from the [URLs] parameter. The value of each key is a URL representing the location of the duplicated file. If this method could not duplicate a file, the corresponding URL is not included in the dictionary.
//   - error: If the operation succeeded for every file, this parameter is `nil`. If the operation failed for one or more files, the parameter contains an error object describing the overall result of the operation in a manner suitable for presentation to the user.
//
// Used by:
//   - [NSWorkspace.DuplicateURLsCompletionHandler]
//   - [NSWorkspace.RecycleURLsCompletionHandler]
type URLNSURLDictionaryErrorHandler = func(*foundation.INSDictionary, error)

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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSView
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSViewFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler handles A Block object called when animations for this transaction group are completed.
//
// Used by:
//   - [NSAnimationContext.RunAnimationGroupCompletionHandler]
//   - [NSAppearance.PerformAsCurrentDrawingAppearance]
//   - [NSApplicationDelegate.ApplicationContinueUserActivityRestorationHandler]
//   - [NSCollectionView.PerformBatchUpdatesCompletionHandler]
//   - [NSDocument.ContinueActivityUsingBlock]
//   - [NSDocument.ContinueAsynchronousWorkOnMainThreadUsingBlock]
//   - [NSDocument.PerformSynchronousFileAccessUsingBlock]
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
//   - [NSDocument.PerformSynchronousFileAccessUsingBlock]
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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// WindowBoolHandler handles The block to execute for each window.
//   - window: The window for which to execute the block.
//   - stop: A Boolean value that stops the enumeration early when set to [true](<doc://com.apple.documentation/documentation/Swift/true>) (the default value is [false](<doc://com.apple.documentation/documentation/Swift/false>)).
//
// Used by:
//   - [NSApplication.EnumerateWindowsWithOptionsUsingBlock]
type WindowBoolHandler = func(*NSWindow, bool)

// NewWindowBoolBlock wraps a Go [WindowBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSApplication.EnumerateWindowsWithOptionsUsingBlock]
func NewWindowBoolBlock(handler WindowBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 bool) {
		var result *NSWindow
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSWindowFromID(resultID)
			result = &v
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// WindowErrorHandler handles A Block object to execute with the results of creating the window.
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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *NSWindow
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSWindowFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// WorkspaceAuthorizationErrorHandler handles The completion handler to call when the authorization request is completed.
//   - authorization: The authorization granted for this app. Use it when creating a new [FileManager](<doc://com.apple.documentation/documentation/Foundation/FileManager>) with [init(authorization:)](<doc://com.apple.documentation/documentation/Foundation/FileManager/init(authorization:)>).
//   - error: `nil` if the app is authorized; otherwise, a pointer to the authorization error.
//
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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *NSWorkspaceAuthorization
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSWorkspaceAuthorizationFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}
