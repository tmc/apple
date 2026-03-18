// Code generated from Apple documentation. DO NOT EDIT.

package appkit

import (
	"unsafe"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NSAccessibilityActionName is constants that describe types of actions.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Action
type NSAccessibilityActionName = string

// NSAccessibilityAnnotationAttributeKey is keys for annotation attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/AnnotationAttributeKey
type NSAccessibilityAnnotationAttributeKey = string

// NSAccessibilityAttributeName is constants that describe attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute
type NSAccessibilityAttributeName = string

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityDateTimeComponentsFlags
type NSAccessibilityDateTimeComponentsFlags = uint

// NSAccessibilityFontAttributeKey is keys for font attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/FontAttributeKey
type NSAccessibilityFontAttributeKey = string

// NSAccessibilityLoadingToken is a token type for loading accessibility elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityLoadingToken
type NSAccessibilityLoadingToken = objc.ID

// NSAccessibilityNotificationName is the name of the notification.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification
type NSAccessibilityNotificationName = string

// NSAccessibilityNotificationUserInfoKey is the key in the user info dictionary for a notification.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/NotificationUserInfoKey
type NSAccessibilityNotificationUserInfoKey = string

// NSAccessibilityOrientationValue is values that indicate the orientation of user interface elements, such as scroll bars and split views.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/OrientationValue
type NSAccessibilityOrientationValue = string

// NSAccessibilityParameterizedAttributeName is values that describe parameterized attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/ParameterizedAttribute
type NSAccessibilityParameterizedAttributeName = string

// NSAccessibilityRole is values that describe types of objects that accessibility elements represent.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role
type NSAccessibilityRole = string

// NSAccessibilityRulerMarkerTypeValue is values that describe ruler marker types.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/RulerMarkerTypeValue
type NSAccessibilityRulerMarkerTypeValue = string

// NSAccessibilityRulerUnitValue is values that indicate the unit values of a ruler or layout area.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/RulerUnitValue
type NSAccessibilityRulerUnitValue = string

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySearchKey
type NSAccessibilitySearchKey = string

// NSAccessibilitySortDirectionValue is values that indicate the sort direction of a column.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/SortDirectionValue
type NSAccessibilitySortDirectionValue = string

// NSAccessibilitySubrole is values that describe specialized object subtypes that accessibility elements represent.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Subrole
type NSAccessibilitySubrole = string

// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyKey
type NSAnimatablePropertyKey = string

// NSAnimationProgress is the animation progress, as a floating-point number between `0.0` and `1.0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/Progress
type NSAnimationProgress = float32

// NSAppKitVersion is constants for determining which version of AppKit is available.
//
// See: https://developer.apple.com/documentation/AppKit/NSAppKitVersion
type NSAppKitVersion = float64

// See: https://developer.apple.com/documentation/AppKit/NSAppearance/Name-swift.struct
type NSAppearanceName = string

// NSAttributedStringDocumentAttributeKey is the attributes you apply to an entire document.
//
// See: https://developer.apple.com/documentation/AppKit/NSAttributedStringDocumentAttributeKey
type NSAttributedStringDocumentAttributeKey = string

// NSAttributedStringDocumentReadingOptionKey is options for constructing an attributed string from data you read from disk.
//
// See: https://developer.apple.com/documentation/AppKit/NSAttributedStringDocumentReadingOptionKey
type NSAttributedStringDocumentReadingOptionKey = string

// NSAttributedStringDocumentType is constants for the document type document attribute key.
//
// See: https://developer.apple.com/documentation/AppKit/NSAttributedStringDocumentType
type NSAttributedStringDocumentType = string

// See: https://developer.apple.com/documentation/AppKit/NSBindingInfoKey
type NSBindingInfoKey = string

// NSBindingName is values that specify a binding for certain methods.
//
// See: https://developer.apple.com/documentation/AppKit/NSBindingName
type NSBindingName = string

// See: https://developer.apple.com/documentation/AppKit/NSBindingOption
type NSBindingOption = string

// NSBitmapImageRepPropertyKey is constants that identify bitmap image representation properties.
//
// See: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/PropertyKey
type NSBitmapImageRepPropertyKey = string

// See: https://developer.apple.com/documentation/AppKit/NSBrowser/ColumnsAutosaveName-swift.typealias
type NSBrowserColumnsAutosaveName = string

// NSCollectionLayoutGroupCustomItemProvider is a closure that creates and returns each of the custom group’s items.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutGroupCustomItemProvider
type NSCollectionLayoutGroupCustomItemProvider = func(NSCollectionLayoutEnvironment) []NSCollectionLayoutGroupCustomItem

// NSCollectionLayoutSectionVisibleItemsInvalidationHandler is a closure called before each layout cycle to allow modification of items in a section immediately before they’re displayed.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSectionVisibleItemsInvalidationHandler
type NSCollectionLayoutSectionVisibleItemsInvalidationHandler = func([]objectivec.IObject, corefoundation.CGPoint, NSCollectionLayoutEnvironment)

// NSCollectionViewCompositionalLayoutSectionProvider is a closure that creates and returns each of the layout’s sections.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewCompositionalLayoutSectionProvider
type NSCollectionViewCompositionalLayoutSectionProvider = func(int, NSCollectionLayoutEnvironment) NSCollectionLayoutSection

// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/DecorationElementKind
type NSCollectionViewDecorationElementKind = string

// NSCollectionViewDiffableDataSourceItemProvider is a closure that configures and returns an item for a collection view from its diffable data source.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDiffableDataSourceReferenceItemProvider
type NSCollectionViewDiffableDataSourceItemProvider = func(NSCollectionView, objc.ID, objectivec.IObject) NSCollectionViewItem

// NSCollectionViewDiffableDataSourceSupplementaryViewProvider is a closure that configures and returns a collection view’s supplementary view, such as a header or footer, from a diffable data source.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDiffableDataSourceReferenceSupplementaryViewProvider
type NSCollectionViewDiffableDataSourceSupplementaryViewProvider = func(NSCollectionView, string, objc.ID) NSView

// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/SupplementaryElementKind
type NSCollectionViewSupplementaryElementKind = string

// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewTransitionLayout/AnimatedKey
type NSCollectionViewTransitionLayoutAnimatedKey = string

// NSColorListName is the name assigned to a color list.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorList/Name-swift.typealias
type NSColorListName = string

// NSColorSpaceName is constants that specify color space names.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpaceName
type NSColorSpaceName = string

// NSControlStateValue is a constant that indicates whether a control is on, off, or in a mixed state.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/StateValue
type NSControlStateValue = int

// NSDataAssetName is the name of a data asset.
//
// See: https://developer.apple.com/documentation/AppKit/NSDataAsset/Name-swift.typealias
type NSDataAssetName = string

// NSDeviceDescriptionKey is these constants are the keys for device description dictionaries.
//
// See: https://developer.apple.com/documentation/AppKit/NSDeviceDescriptionKey
type NSDeviceDescriptionKey = string

// NSDraggingImageComponentKey is keys that identify components of a dragging image.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingItem/ImageComponentKey
type NSDraggingImageComponentKey = string

// NSFontCollectionActionTypeKey is the following actions are possible values of the [actionUserInfoKey] in the [didChangeNotification] `userInfo` method.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/ActionTypeKey
type NSFontCollectionActionTypeKey = string

// NSFontCollectionMatchingOptionKey is these constants are used by the [matchingDescriptors(options:)] and [matchingDescriptors(forFamily:options:)] options dictionary parameters.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontCollectionMatchingOptionKey
type NSFontCollectionMatchingOptionKey = string

// NSFontCollectionName is the constants represent the standard mutable collection names—these names are included in the list of [allFontCollectionNames]–they have special meaning to the Cocoa font system and should not be hidden or renamed.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/Name
type NSFontCollectionName = string

// NSFontCollectionUserInfoKey is these constants are used as keys in the [didChangeNotification] `userInfo` dictionary to indicate the changes that have taken place.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontCollection/UserInfoKey
type NSFontCollectionUserInfoKey = string

// NSFontDescriptorAttributeName is constants for the names of font attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/AttributeName
type NSFontDescriptorAttributeName = string

// NSFontDescriptorFeatureKey is constants to use as keys to retrieve information about a font descriptor from its feature dictionary.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/FeatureKey
type NSFontDescriptorFeatureKey = string

// NSFontDescriptorSystemDesign is constants for font designs, such as monospace, rounded, and serif.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SystemDesign
type NSFontDescriptorSystemDesign = string

// NSFontDescriptorTraitKey is constants that can be used as keys to retrieve information about a font descriptor from its trait dictionary.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/TraitKey
type NSFontDescriptorTraitKey = string

// NSFontDescriptorVariationKey is constants that can be used as keys to retrieve information about a font descriptor from its variation axis dictionary.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/VariationKey
type NSFontDescriptorVariationKey = string

// NSFontFamilyClass is constants that classify certain stylistic qualities of the font.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontFamilyClass
type NSFontFamilyClass = uint32

// NSFontSymbolicTraits is a symbolic description of stylistic aspects of a font.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontSymbolicTraits
type NSFontSymbolicTraits = uint32

// NSFontTextStyle is constants that specify the preferred text styles you use with fonts.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/TextStyle
type NSFontTextStyle = string

// NSFontTextStyleOptionKey is the options that you apply when requesting the font or font descriptor of a preferred text style.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/TextStyleOptionKey
type NSFontTextStyleOptionKey = string

// NSFontWeight is system-defined font-weight values.
//
// See: https://developer.apple.com/documentation/AppKit/NSFont/Weight
type NSFontWeight = float64

// See: https://developer.apple.com/documentation/AppKit/NSFont/Width
type NSFontWidth = float64

// NSGlyph is the type used to specify glyphs.
//
// See: https://developer.apple.com/documentation/AppKit/NSGlyph
type NSGlyph = uint32

// NSGraphicsContextAttributeKey is constants that specify the dictionary keys for the attributes of the graphics context.
//
// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/AttributeKey
type NSGraphicsContextAttributeKey = string

// NSGraphicsContextRepresentationFormatName is constants that specify values for the representation format name key in a graphic context’s attributes dictionary.
//
// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/RepresentationFormatName
type NSGraphicsContextRepresentationFormatName = string

// See: https://developer.apple.com/documentation/AppKit/NSHelpManager/AnchorName
type NSHelpAnchorName = string

// See: https://developer.apple.com/documentation/AppKit/NSHelpManager/BookName
type NSHelpBookName = string

// See: https://developer.apple.com/documentation/AppKit/NSHelpManager/ContextHelpKey
type NSHelpManagerContextHelpKey = string

// NSImageHintKey is constants for the keys to include in a hints dictionary when drawing the image.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/HintKey
type NSImageHintKey = string

// NSImageName is named images, defined by the system or you, for use in your app.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/Name-swift.typealias
type NSImageName = string

// NSLayoutPriority is layout priority used to indicate the relative importance of constraints, allowing Auto Layout to make appropriate tradeoffs when satisfying the constraints of the system as a whole.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Priority-swift.struct
type NSLayoutPriority = float32

// NSModalResponse is a set of button return values for modal dialogs.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/ModalResponse
type NSModalResponse = int

// See: https://developer.apple.com/documentation/AppKit/NSNib/Name
type NSNibName = string

// See: https://developer.apple.com/documentation/AppKit/NSPageController/ObjectIdentifier
type NSPageControllerObjectIdentifier = string

// NSPasteboardDetectionPattern is a pattern to detect on the pasteboard, such as a URL, text, or a number.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardDetectionPattern
type NSPasteboardDetectionPattern = string

// NSPasteboardMetadataType is a metadata type to detect on the pasteboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardMetadataType
type NSPasteboardMetadataType = string

// NSPasteboardName is constants that represent the standard pasteboard names.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/Name-swift.struct
type NSPasteboardName = string

// NSPasteboardReadingOptionKey is options for reading pasteboard data.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptionKey
type NSPasteboardReadingOptionKey = string

// NSPasteboardType is the supported pasteboard types.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType
type NSPasteboardType = string

// NSPasteboardTypeFindPanelSearchOptionKey is search options for the find panel.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/FindPanelSearchOptionKey
type NSPasteboardTypeFindPanelSearchOptionKey = string

// NSPasteboardTypeTextFinderOptionKey is search options for text in Finder.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/TextFinderOptionKey
type NSPasteboardTypeTextFinderOptionKey = string

// NSPopoverCloseReasonValue is values that specify the reason for the [willCloseNotification] notification.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/CloseReason
type NSPopoverCloseReasonValue = string

// NSPrintInfoAttributeKey is constants that specify print job attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey
type NSPrintInfoAttributeKey = string

// NSPrintInfoSettingKey is the type you use to specify a print info setting key.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/SettingKey
type NSPrintInfoSettingKey = string

// NSPrintJobDispositionValue is constants that specify values for the print job disposition.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintInfo/JobDisposition-swift.struct
type NSPrintJobDispositionValue = string

// NSPrintPanelAccessorySummaryKey is constants that specify the accessory panel keys.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintPanel/AccessorySummaryKey
type NSPrintPanelAccessorySummaryKey = string

// NSPrintPanelJobStyleHint is constants that specify job style hints for activating the simplified Print panel interface and setting the options to display.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrintPanel/JobStyleHint-swift.struct
type NSPrintPanelJobStyleHint = string

// NSPrinterPaperName is the type you use to specify the name of a type of paper.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrinter/PaperName
type NSPrinterPaperName = string

// NSPrinterTypeName is the type you use to describe a printer’s make and model.
//
// See: https://developer.apple.com/documentation/AppKit/NSPrinter/TypeName
type NSPrinterTypeName = string

// NSRuleEditorPredicatePartKey is these strings are used as keys to the dictionary returned from the delegate’s [ruleEditor(_:predicatePartsForCriterion:withDisplayValue:inRow:)] optional method. To construct a valid predicate, the union of the dictionaries for each item in the row must contain the required parts.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/PredicatePartKey
type NSRuleEditorPredicatePartKey = string

// See: https://developer.apple.com/documentation/AppKit/NSRulerView/UnitName
type NSRulerViewUnitName = string

// NSSearchFieldRecentsAutosaveName is the string that stores the name under which a search field automatically archives a list of recent search strings.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchField/RecentsAutosaveName-swift.typealias
type NSSearchFieldRecentsAutosaveName = string

// See: https://developer.apple.com/documentation/AppKit/NSServiceProviderName
type NSServiceProviderName = string

// NSSharingServiceName is constants that describe the sharing services that macOS supports.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingService/Name
type NSSharingServiceName = string

// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessory/Width
type NSSliderAccessoryWidth = float64

// See: https://developer.apple.com/documentation/AppKit/NSSound/Name-swift.typealias
type NSSoundName = string

// See: https://developer.apple.com/documentation/AppKit/NSSound/PlaybackDeviceIdentifier-swift.typealias
type NSSoundPlaybackDeviceIdentifier = string

// NSSpeechCommandDelimiterKey is keys for the command delimiters.
//
// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/CommandDelimiterKey
type NSSpeechCommandDelimiterKey = string

// NSSpeechDictionaryKey is these constants identify key-value pairs used to add vocabulary to the dictionary using [addSpeechDictionary(_:)].
//
// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/DictionaryKey
type NSSpeechDictionaryKey = string

// NSSpeechErrorKey is keys that identify errors that may occur during speech synthesis.
//
// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/ErrorKey
type NSSpeechErrorKey = string

// NSSpeechMode is keys for the speaking mode.
//
// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/Mode
type NSSpeechMode = string

// NSSpeechPhonemeInfoKey is keys for the speech phoneme information.
//
// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/PhonemeInfoKey
type NSSpeechPhonemeInfoKey = string

// NSSpeechPropertyKey is these constants are used with [setObject(_:forProperty:)] and [object(forProperty:)] to get or set the characteristics of a synthesizer.
//
// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey
type NSSpeechPropertyKey = string

// NSSpeechStatusKey is keys for the speech synthesizier status.
//
// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/StatusKey
type NSSpeechStatusKey = string

// NSSpeechSynthesizerInfoKey is keys for the speech synthesizier information.
//
// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/SpeechPropertyKey/SynthesizerInfoKey
type NSSpeechSynthesizerInfoKey = string

//
// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/VoiceName
type NSSpeechSynthesizerVoiceName = string

// NSSplitViewAutosaveName is the type that specifies the split view’s autosave name.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitView/AutosaveName-swift.typealias
type NSSplitViewAutosaveName = string

// NSStackViewVisibilityPriority is the various Auto Layout priorities for a view in the stack view to remain attached.
//
// See: https://developer.apple.com/documentation/AppKit/NSStackView/VisibilityPriority
type NSStackViewVisibilityPriority = float32

// See: https://developer.apple.com/documentation/AppKit/NSStatusItem/AutosaveName-swift.typealias
type NSStatusItemAutosaveName = string

// NSStoryboardControllerCreator is a block that you use to handle the custom creation of controller objects from your storyboard file.
//
// See: https://developer.apple.com/documentation/AppKit/NSStoryboardControllerCreator
type NSStoryboardControllerCreator = func(foundation.NSCoder) objectivec.IObject

// NSStoryboardName is the name of the storyboard file.
//
// See: https://developer.apple.com/documentation/AppKit/NSStoryboard/Name
type NSStoryboardName = string

// NSStoryboardSceneIdentifier is a string that uniquely identifies a view controller or window controller in your storyboard file.
//
// See: https://developer.apple.com/documentation/AppKit/NSStoryboard/SceneIdentifier
type NSStoryboardSceneIdentifier = string

// See: https://developer.apple.com/documentation/AppKit/NSStoryboardSegue/Identifier-swift.typealias
type NSStoryboardSegueIdentifier = string

// See: https://developer.apple.com/documentation/AppKit/NSTableView/AutosaveName-swift.typealias
type NSTableViewAutosaveName = string

// NSTableViewDiffableDataSourceCellProvider is a closure that configures and returns a cell for a table view from its diffable data source.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReferenceCellProvider
type NSTableViewDiffableDataSourceCellProvider = func(NSTableView, NSTableColumn, int, objectivec.IObject) NSView

// NSTableViewDiffableDataSourceRowProvider is a closure that configures and returns a row view for a table view from its diffable data source.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReferenceRowProvider
type NSTableViewDiffableDataSourceRowProvider = func(NSTableView, int, objectivec.IObject) NSTableRowView

// NSTableViewDiffableDataSourceSectionHeaderViewProvider is a closure that configures and returns a section header view for a table view from its diffable data source.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReferenceSectionHeaderViewProvider
type NSTableViewDiffableDataSourceSectionHeaderViewProvider = func(NSTableView, int, objectivec.IObject) NSView

// NSTextCheckingOptionKey is constants that define options for text checking.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/OptionKey
type NSTextCheckingOptionKey = string

// NSTextContentType is constants that identify the semantic meaning for a text-entry area.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentType
type NSTextContentType = string

// NSTextEffectStyle is constants for the type of effect to apply to the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextEffectStyle
type NSTextEffectStyle = string

// NSTextHighlightColorScheme is constants that specify the highlight color to use with the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextHighlightColorScheme
type NSTextHighlightColorScheme = string

// NSTextHighlightStyle is constants that specify the type of highlight to apply to text.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextHighlightStyle
type NSTextHighlightStyle = string

// See: https://developer.apple.com/documentation/AppKit/NSTextInputSourceIdentifier
type NSTextInputSourceIdentifier = string

// NSTextLayoutSectionKey is constants for the text layout sections document attribute key.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutSectionKey
type NSTextLayoutSectionKey = string

// NSTextListMarkerFormat is constants that describe marker symbols you can apply to list elements in text lists.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextList/MarkerFormat-swift.struct
type NSTextListMarkerFormat = string

// NSTextStorageEditedOptions is ** Deprecations ***
//
// See: https://developer.apple.com/documentation/AppKit/NSTextStorageEditedOptions
type NSTextStorageEditedOptions = uint

// NSTextTabOptionKey is the terminating character for a tab column.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTab/OptionKey
type NSTextTabOptionKey = string

// NSToolbarIdentifier is a string value that you use to differentiate your app’s toolbars.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/Identifier-swift.typealias
type NSToolbarIdentifier = string

// NSToolbarItemIdentifier is constants for the standard toolbar items that the system provides.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/Identifier
type NSToolbarItemIdentifier = string

// NSToolbarItemVisibilityPriority is constants that indicate which toolbar items to keep in the toolbar when space is limited.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/VisibilityPriority-swift.struct
type NSToolbarItemVisibilityPriority = int

// NSToolbarUserInfoKey is constants for specifying toolbar-related information in notifications.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarUserInfoKey
type NSToolbarUserInfoKey = string

// NSTouchBarCustomizationIdentifier is the default type for a Touch Bar customization identifier.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBar/CustomizationIdentifier-swift.typealias
type NSTouchBarCustomizationIdentifier = string

// NSTouchBarItemIdentifier is an identifier for an item in the Touch Bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/Identifier-swift.struct
type NSTouchBarItemIdentifier = string

// NSTouchBarItemPriority is priorities for the visibility of a Touch Bar item.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/Priority
type NSTouchBarItemPriority = float32

// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceItemIdentifier
type NSUserInterfaceItemIdentifier = string

// NSViewAnimationEffectName is the following constants specify the animation effect to apply and are used as values for the animation effect property of the animation view. See the description of [effect] for usage details.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewAnimation/EffectName
type NSViewAnimationEffectName = string

// NSViewAnimationKey is the following string constants are keys for the dictionaries in the array passed into [init(viewAnimations:)] and [viewAnimations].
//
// See: https://developer.apple.com/documentation/AppKit/NSViewAnimation/Key
type NSViewAnimationKey = string

// NSVoiceAttributeKey is the following constants are keys for the dictionary returned by [attributes(forVoice:)].
//
// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/VoiceAttributeKey
type NSVoiceAttributeKey = string

// NSVoiceGenderName is the following constants define voice gender attributes, which are the allowable values of the [gender] key returned by [attributes(forVoice:)].
//
// Deprecated: Deprecated since macOS 14.0. Use AVSpeechSynthesizer in AVFoundation instead
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/VoiceGender
type NSVoiceGenderName = string

// NSWindowFrameAutosaveName is the type of a window’s frame autosave name.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/FrameAutosaveName-swift.typealias
type NSWindowFrameAutosaveName = string

// NSWindowLevel is the standard window levels in macOS.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/Level-swift.struct
type NSWindowLevel = int

// NSWindowPersistableFrameDescriptor is the type of a window’s frame descriptor.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/PersistableFrameDescriptor
type NSWindowPersistableFrameDescriptor = string

// NSWindowTabbingIdentifier is a value that allows a group of related windows.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/TabbingIdentifier-swift.typealias
type NSWindowTabbingIdentifier = string

// NSWorkspaceDesktopImageOptionKey is keys that indicate how to display a new desktop image.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/DesktopImageOptionKey
type NSWorkspaceDesktopImageOptionKey = string

// WindowRef is an opaque Carbon window reference.

type WindowRef = unsafe.Pointer

