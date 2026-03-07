// Code generated from Apple documentation. DO NOT EDIT.

package appkit

import (
	"unsafe"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

type NSAccessibilityActionName = string







type NSAccessibilityAnnotationAttributeKey = string







type NSAccessibilityAttributeName = string







type NSAccessibilityDateTimeComponentsFlags = uint







type NSAccessibilityFontAttributeKey = string







type NSAccessibilityLoadingToken = objc.ID







type NSAccessibilityNotificationName = string







type NSAccessibilityNotificationUserInfoKey = string







type NSAccessibilityOrientationValue = string







type NSAccessibilityParameterizedAttributeName = string







type NSAccessibilityRole = string







type NSAccessibilityRulerMarkerTypeValue = string







type NSAccessibilityRulerUnitValue = string







type NSAccessibilitySearchKey = string







type NSAccessibilitySortDirectionValue = string







type NSAccessibilitySubrole = string







type NSAnimatablePropertyKey = string







type NSAnimationProgress = float32







type NSAppKitVersion = float64







type NSAppearanceName = string







type NSAttributedStringDocumentAttributeKey = string







type NSAttributedStringDocumentReadingOptionKey = string







type NSAttributedStringDocumentType = string







type NSBindingInfoKey = string







type NSBindingName = string







type NSBindingOption = string







type NSBitmapImageRepPropertyKey = string







type NSBrowserColumnsAutosaveName = string







type NSCollectionLayoutGroupCustomItemProvider = func(NSCollectionLayoutEnvironment) []NSCollectionLayoutGroupCustomItem







type NSCollectionLayoutSectionVisibleItemsInvalidationHandler = func([]objectivec.IObject, corefoundation.CGPoint, NSCollectionLayoutEnvironment)







type NSCollectionViewCompositionalLayoutSectionProvider = func(int, NSCollectionLayoutEnvironment) NSCollectionLayoutSection







type NSCollectionViewDecorationElementKind = string







type NSCollectionViewDiffableDataSourceItemProvider = func(NSCollectionView, objc.ID, objectivec.IObject) NSCollectionViewItem







type NSCollectionViewDiffableDataSourceSupplementaryViewProvider = func(NSCollectionView, string, objc.ID) NSView







type NSCollectionViewSupplementaryElementKind = string







type NSCollectionViewTransitionLayoutAnimatedKey = string







type NSColorListName = string







type NSColorSpaceName = string







type NSControlStateValue = int







type NSDataAssetName = string







type NSDeviceDescriptionKey = string







type NSDraggingImageComponentKey = string







type NSFontCollectionActionTypeKey = string







type NSFontCollectionMatchingOptionKey = string







type NSFontCollectionName = string







type NSFontCollectionUserInfoKey = string







type NSFontDescriptorAttributeName = string







type NSFontDescriptorFeatureKey = string







type NSFontDescriptorSystemDesign = string







type NSFontDescriptorTraitKey = string







type NSFontDescriptorVariationKey = string







type NSFontFamilyClass = uint32







type NSFontSymbolicTraits = uint32







type NSFontTextStyle = string







type NSFontTextStyleOptionKey = string







type NSFontWeight = float64







type NSFontWidth = float64







type NSGlyph = uint32







type NSGraphicsContextAttributeKey = string







type NSGraphicsContextRepresentationFormatName = string







type NSHelpAnchorName = string







type NSHelpBookName = string







type NSHelpManagerContextHelpKey = string







type NSImageHintKey = string







type NSImageName = string







type NSLayoutPriority = float32







type NSModalResponse = int







type NSNibName = string







type NSPageControllerObjectIdentifier = string







type NSPasteboardDetectionPattern = string







type NSPasteboardMetadataType = string







type NSPasteboardName = string







type NSPasteboardReadingOptionKey = string







type NSPasteboardType = string







type NSPasteboardTypeFindPanelSearchOptionKey = string







type NSPasteboardTypeTextFinderOptionKey = string







type NSPopoverCloseReasonValue = string







type NSPrintInfoAttributeKey = string







type NSPrintInfoSettingKey = string







type NSPrintJobDispositionValue = string







type NSPrintPanelAccessorySummaryKey = string







type NSPrintPanelJobStyleHint = string







type NSPrinterPaperName = string







type NSPrinterTypeName = string







type NSRuleEditorPredicatePartKey = string







type NSRulerViewUnitName = string







type NSSearchFieldRecentsAutosaveName = string







type NSServiceProviderName = string







type NSSharingServiceName = string







type NSSliderAccessoryWidth = float64







type NSSoundName = string







type NSSoundPlaybackDeviceIdentifier = string







// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
type NSSpeechCommandDelimiterKey = string







// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
type NSSpeechDictionaryKey = string







// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
type NSSpeechErrorKey = string







// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
type NSSpeechMode = string







// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
type NSSpeechPhonemeInfoKey = string







// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
type NSSpeechPropertyKey = string







// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
type NSSpeechStatusKey = string







// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
type NSSpeechSynthesizerInfoKey = string







// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
type NSSpeechSynthesizerVoiceName = string







type NSSplitViewAutosaveName = string







type NSStackViewVisibilityPriority = float32







type NSStatusItemAutosaveName = string







type NSStoryboardControllerCreator = func(foundation.NSCoder) objectivec.IObject







type NSStoryboardName = string







type NSStoryboardSceneIdentifier = string







type NSStoryboardSegueIdentifier = string







type NSTableViewAutosaveName = string







type NSTableViewDiffableDataSourceCellProvider = func(NSTableView, NSTableColumn, int, objectivec.IObject) NSView







type NSTableViewDiffableDataSourceRowProvider = func(NSTableView, int, objectivec.IObject) NSTableRowView







type NSTableViewDiffableDataSourceSectionHeaderViewProvider = func(NSTableView, int, objectivec.IObject) NSView







type NSTextCheckingOptionKey = string







type NSTextContentType = string







type NSTextEffectStyle = string







type NSTextHighlightColorScheme = string







type NSTextHighlightStyle = string







type NSTextInputSourceIdentifier = string







type NSTextLayoutSectionKey = string







type NSTextListMarkerFormat = string







type NSTextStorageEditedOptions = uint







type NSTextTabOptionKey = string







type NSToolbarIdentifier = string







type NSToolbarItemIdentifier = string







type NSToolbarItemVisibilityPriority = int







type NSToolbarUserInfoKey = string







type NSTouchBarCustomizationIdentifier = string







type NSTouchBarItemIdentifier = string







type NSTouchBarItemPriority = float32







type NSUserInterfaceItemIdentifier = string







type NSViewAnimationEffectName = string







type NSViewAnimationKey = string







// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
type NSVoiceAttributeKey = string







// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
type NSVoiceGenderName = string







type NSWindowFrameAutosaveName = string







type NSWindowLevel = int







type NSWindowPersistableFrameDescriptor = string







type NSWindowTabbingIdentifier = string







type NSWorkspaceDesktopImageOptionKey = string





// WindowRef is an opaque Carbon window reference.

type WindowRef = unsafe.Pointer



